package provider

import (
	"context"
	"fmt"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource              = &LinkResource{}
	_ resource.ResourceWithConfigure = &LinkResource{}
)

// These mirror the server-side checks in the link service so mistakes fail at plan time.
var (
	gcpTenantIDRegex = regexp.MustCompile(
		`^(?:organizations/[0-9]{1,32}|projects/(?:[a-z][a-z0-9-]{4,28}[a-z0-9]|[0-9]{1,32})|[0-9]{1,32}|[a-z][a-z0-9-]{4,28}[a-z0-9])$`)
)

func NewLinkResource() resource.Resource {
	return &LinkResource{}
}

// LinkResource implements scm_cngfw_gcp_link. Create generates a link code for a GCP
// tenant; Delete unlinks the tenant. There is no API to read a link back, so Read keeps
// the stored state.
type LinkResource struct {
	client *Client
}

type linkResourceModel struct {
	Id          types.String `tfsdk:"id"`
	GcpTenantId types.String `tfsdk:"gcp_tenant_id"`
	TsgId       types.String `tfsdk:"tsg_id"`
	CspId       types.String `tfsdk:"csp_id"`
	Region      types.String `tfsdk:"region"`
	LinkCode    types.String `tfsdk:"link_code"`
	ExpiresAt   types.Int64  `tfsdk:"expires_at"`
}

func (r *LinkResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "scm_cngfw_gcp_link"
}

func (r *LinkResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Links a Google Cloud tenant to a Strata Cloud Manager TSG for Cloud NGFW for GCP. " +
			"Creating this resource generates a one-time link code, which the customer enters in the Google Cloud " +
			"console when creating a firewall endpoint. The code expires after 48 hours. Destroying this resource " +
			"unlinks the firewall endpoint created with the code. A code that was never used is not revoked: it " +
			"stays valid until it expires. Changing any argument forces a new resource.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "The GCP tenant ID.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"gcp_tenant_id": schema.StringAttribute{
				Description: "The GCP tenant to link: `organizations/{number}`, `projects/{project ID or number}`, a bare number, " +
					"or a bare project ID (6-30 characters: lowercase letters, digits and hyphens, starting with a letter). " +
					"It must match the tenant the customer selects in the Google Cloud console exactly.",
				Required: true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(gcpTenantIDRegex, "must be organizations/{number}, projects/{id or number}, a number, or a project ID"),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"tsg_id": schema.StringAttribute{
				Description: "The Tenant Service Group the GCP tenant is linked to: the TSG the provider credentials are scoped to.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"csp_id": schema.StringAttribute{
				Description: "The Customer Support Portal account ID of the TSG, as looked up and returned by the link service.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"region": schema.StringAttribute{
				Description: "The Cloud NGFW region the link is stored in, sent as the `x-panw-region` header. " +
					"If unset, the API gateway default is used.",
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"link_code": schema.StringAttribute{
				Description: "The one-time link code to enter in the Google Cloud console.",
				Computed:    true,
				Sensitive:   true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"expires_at": schema.Int64Attribute{
				Description: "When the link code expires, in Unix epoch seconds.",
				Computed:    true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (r *LinkResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	clients, ok := req.ProviderData.(map[string]interface{})
	if !ok {
		resp.Diagnostics.AddError("Unexpected Configure Type", fmt.Sprintf("Expected map[string]interface{}, got: %T.", req.ProviderData))
		return
	}
	client, ok := clients["cngfw_gcp"].(*Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Client Type", "Expected *Client for cngfw_gcp.")
		return
	}
	r.client = client
}

func (r *LinkResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan linkResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// The service takes the TSG from the provider credentials (the API gateway passes it on)
	// and looks up the TSG's CSP account and cloud management tenant itself.
	tflog.Info(ctx, "Generating Cloud NGFW for GCP link code", map[string]interface{}{
		"gcp_tenant_id": plan.GcpTenantId.ValueString(),
	})
	result, err := r.client.GenerateLinkCode(ctx, plan.Region.ValueString(), plan.GcpTenantId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error generating Cloud NGFW for GCP link code", err.Error())
		return
	}

	plan.Id = plan.GcpTenantId
	plan.TsgId = types.StringValue(result.GetTSGID())
	plan.CspId = types.StringValue(result.GetCSPID())
	plan.LinkCode = types.StringValue(result.GetCode())
	plan.ExpiresAt = types.Int64Value(result.GetExpiresAt())
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *LinkResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// The link API has no read route, so the stored state is the source of truth.
	var state linkResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *LinkResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Every argument forces replacement, so there is nothing to send. Keep the planned values.
	var plan linkResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *LinkResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state linkResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Unlink is keyed by the link code: it retires the firewall endpoint bound with this code.
	tflog.Info(ctx, "Unlinking Cloud NGFW for GCP link code", map[string]interface{}{
		"gcp_tenant_id": state.GcpTenantId.ValueString(),
	})
	if err := r.client.Unlink(ctx, state.Region.ValueString(), state.GcpTenantId.ValueString(), state.LinkCode.ValueString()); err != nil {
		resp.Diagnostics.AddError("Error unlinking Cloud NGFW for GCP link code", err.Error())
	}
}
