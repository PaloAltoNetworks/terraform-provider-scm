package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/paloaltonetworks/scm-go/generated/ztna_connector_all"

	"github.com/paloaltonetworks/terraform-provider-scm/internal/utils"
)

var (
	_ resource.Resource              = &ConnectorQuiesceResource{}
	_ resource.ResourceWithConfigure = &ConnectorQuiesceResource{}
)

func NewConnectorQuiesceResource() resource.Resource {
	return &ConnectorQuiesceResource{}
}

type ConnectorQuiesceResource struct {
	client *ztna_connector_all.APIClient
}

type connectorQuiesceResourceModel struct {
	Oid  types.String `tfsdk:"oid"`
	Mode types.String `tfsdk:"mode"`
	Tfid types.String `tfsdk:"tfid"`
}

func (r *ConnectorQuiesceResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "ztna_connector_quiesce"
}

func (r *ConnectorQuiesceResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "",
		Attributes: map[string]schema.Attribute{
			"oid": schema.StringAttribute{
				Description: "Parent connector OID. Changing this forces a new resource.",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"mode": schema.StringAttribute{
				Description: "",
				Required:    true,
			},
			"tfid": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (r *ConnectorQuiesceResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	clients, ok := req.ProviderData.(map[string]interface{})
	if !ok {
		resp.Diagnostics.AddError("Unexpected Configure Type", fmt.Sprintf("Expected map[string]interface{}, got: %T.", req.ProviderData))
		return
	}
	client, ok := clients["ztna_connector_all"].(*ztna_connector_all.APIClient)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Client Type", "Expected *ztna_connector_all.APIClient.")
		return
	}
	r.client = client
}

func (r *ConnectorQuiesceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan connectorQuiesceResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	body := ztna_connector_all.ConnectorQuiesce{}
	body.SetMode(plan.Mode.ValueString())
	_, err := r.client.ConnectorAPI.UpdateConnectorsQuiesceByID(ctx, plan.Oid.ValueString()).ConnectorQuiesce(body).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error creating ConnectorQuiesce", fmt.Sprintf("API error: %s", err.Error()))
		resp.Diagnostics.AddError("API Request Failed", utils.PrintScmError(err))
		return
	}
	result, _, err2 := r.client.ConnectorAPI.GetConnectorsQuiesceByID(ctx, plan.Oid.ValueString()).Execute()
	if err2 != nil {
		resp.Diagnostics.AddError("Error reading ConnectorQuiesce after create", fmt.Sprintf("API error: %s", err2.Error()))
		return
	}
	if result != nil {
		plan.Mode = types.StringValue(result.GetMode())
	}
	plan.Tfid = types.StringValue(":::" + plan.Oid.ValueString())
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *ConnectorQuiesceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state connectorQuiesceResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	result, _, err := r.client.ConnectorAPI.GetConnectorsQuiesceByID(ctx, state.Oid.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error reading ConnectorQuiesce", fmt.Sprintf("API error: %s", err.Error()))
		resp.Diagnostics.AddError("API Request Failed", utils.PrintScmError(err))
		return
	}
	if result != nil {
		state.Mode = types.StringValue(result.GetMode())
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *ConnectorQuiesceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan connectorQuiesceResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state connectorQuiesceResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	body := ztna_connector_all.ConnectorQuiesce{}
	body.SetMode(plan.Mode.ValueString())
	_, err := r.client.ConnectorAPI.UpdateConnectorsQuiesceByID(ctx, state.Oid.ValueString()).ConnectorQuiesce(body).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error updating ConnectorQuiesce", fmt.Sprintf("API error: %s", err.Error()))
		resp.Diagnostics.AddError("API Request Failed", utils.PrintScmError(err))
		return
	}
	result, _, err2 := r.client.ConnectorAPI.GetConnectorsQuiesceByID(ctx, state.Oid.ValueString()).Execute()
	if err2 == nil && result != nil {
		plan.Mode = types.StringValue(result.GetMode())
	}
	plan.Tfid = state.Tfid
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *ConnectorQuiesceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state connectorQuiesceResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
