package provider

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/paloaltonetworks/scm-go/generated/ztna_connector_all"

	"github.com/paloaltonetworks/terraform-provider-scm/internal/utils"
)

var (
	_ resource.Resource              = &ConnectorGroupScheduledUpgradeResource{}
	_ resource.ResourceWithConfigure = &ConnectorGroupScheduledUpgradeResource{}
)

func NewConnectorGroupScheduledUpgradeResource() resource.Resource {
	return &ConnectorGroupScheduledUpgradeResource{}
}

type ConnectorGroupScheduledUpgradeResource struct {
	client *ztna_connector_all.APIClient
}

type connectorGroupScheduledUpgradeResourceModel struct {
	Oid               types.String `tfsdk:"oid"`
	DrainTimeout      types.Int64  `tfsdk:"drain_timeout"`
	ImageId           types.String `tfsdk:"image_id"`
	RollingUpgrade    types.Bool   `tfsdk:"rolling_upgrade"`
	ScheduledDownload types.String `tfsdk:"scheduled_download"`
	ScheduledUpgrade  types.String `tfsdk:"scheduled_upgrade"`
	Tfid              types.String `tfsdk:"tfid"`
}

func (r *ConnectorGroupScheduledUpgradeResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "ztna_connector_group_scheduled_upgrade"
}

func (r *ConnectorGroupScheduledUpgradeResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
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
			"drain_timeout": schema.Int64Attribute{
				Description: "Drain timeout in seconds for rolling upgrades.  If omitted, defaults to 0.",
				Optional:    true,
				Computed:    true,
			},
			"image_id": schema.StringAttribute{
				Description: "The connector image version ID to upgrade to.",
				Required:    true,
			},
			"rolling_upgrade": schema.BoolAttribute{
				Description: "Whether to perform a rolling upgrade.  If omitted, defaults to false. Requires SaasAgent version 6.1.0 or later.",
				Optional:    true,
				Computed:    true,
			},
			"scheduled_download": schema.StringAttribute{
				Description: "The scheduled download time in RFC3339 format (UTC).  If omitted, defaults to current UTC time.",
				Optional:    true,
				Computed:    true,
			},
			"scheduled_upgrade": schema.StringAttribute{
				Description: "The scheduled upgrade time in RFC3339 format (UTC).  Must be after the scheduled_download time. If omitted, defaults to current UTC time.",
				Optional:    true,
				Computed:    true,
			},
			"tfid": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (r *ConnectorGroupScheduledUpgradeResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *ConnectorGroupScheduledUpgradeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan connectorGroupScheduledUpgradeResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	body := ztna_connector_all.ConnectorGroupScheduledUpgrade{}
	if !plan.DrainTimeout.IsNull() {
		body.SetDrainTimeout(int32(plan.DrainTimeout.ValueInt64()))
	}
	body.SetImageId(plan.ImageId.ValueString())
	if !plan.RollingUpgrade.IsNull() {
		body.SetRollingUpgrade(plan.RollingUpgrade.ValueBool())
	}
	if !plan.ScheduledDownload.IsNull() {
		if t, parseErr := time.Parse(time.RFC3339, plan.ScheduledDownload.ValueString()); parseErr == nil {
			body.SetScheduledDownload(t)
		}
	}
	if !plan.ScheduledUpgrade.IsNull() {
		if t, parseErr := time.Parse(time.RFC3339, plan.ScheduledUpgrade.ValueString()); parseErr == nil {
			body.SetScheduledUpgrade(t)
		}
	}
	_, err := r.client.ConnectorGroupAPI.CreateConnectorGroupScheduledUpgrade(ctx, plan.Oid.ValueString()).ConnectorGroupScheduledUpgrade(body).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error creating ConnectorGroupScheduledUpgrade", fmt.Sprintf("API error: %s", err.Error()))
		resp.Diagnostics.AddError("API Request Failed", utils.PrintScmError(err))
		return
	}
	result, _, err2 := r.client.ConnectorGroupAPI.GetConnectorGroupScheduledUpgrade(ctx, plan.Oid.ValueString()).Execute()
	if err2 != nil {
		resp.Diagnostics.AddError("Error reading ConnectorGroupScheduledUpgrade after create", fmt.Sprintf("API error: %s", err2.Error()))
		return
	}
	if result != nil {
		plan.DrainTimeout = types.Int64Value(int64(result.GetDrainTimeout()))
		plan.ImageId = types.StringValue(result.GetImageId())
		plan.RollingUpgrade = types.BoolValue(result.GetRollingUpgrade())
		plan.ScheduledDownload = types.StringValue(result.GetScheduledDownload().Format(time.RFC3339))
		plan.ScheduledUpgrade = types.StringValue(result.GetScheduledUpgrade().Format(time.RFC3339))
	}
	plan.Tfid = types.StringValue(":::" + plan.Oid.ValueString())
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *ConnectorGroupScheduledUpgradeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state connectorGroupScheduledUpgradeResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	result, _, err := r.client.ConnectorGroupAPI.GetConnectorGroupScheduledUpgrade(ctx, state.Oid.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error reading ConnectorGroupScheduledUpgrade", fmt.Sprintf("API error: %s", err.Error()))
		resp.Diagnostics.AddError("API Request Failed", utils.PrintScmError(err))
		return
	}
	if result != nil {
		state.DrainTimeout = types.Int64Value(int64(result.GetDrainTimeout()))
		state.ImageId = types.StringValue(result.GetImageId())
		state.RollingUpgrade = types.BoolValue(result.GetRollingUpgrade())
		state.ScheduledDownload = types.StringValue(result.GetScheduledDownload().Format(time.RFC3339))
		state.ScheduledUpgrade = types.StringValue(result.GetScheduledUpgrade().Format(time.RFC3339))
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *ConnectorGroupScheduledUpgradeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan connectorGroupScheduledUpgradeResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state connectorGroupScheduledUpgradeResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	body := ztna_connector_all.ConnectorGroupScheduledUpgrade{}
	if !plan.DrainTimeout.IsNull() {
		body.SetDrainTimeout(int32(plan.DrainTimeout.ValueInt64()))
	}
	body.SetImageId(plan.ImageId.ValueString())
	if !plan.RollingUpgrade.IsNull() {
		body.SetRollingUpgrade(plan.RollingUpgrade.ValueBool())
	}
	if !plan.ScheduledDownload.IsNull() {
		if t, parseErr := time.Parse(time.RFC3339, plan.ScheduledDownload.ValueString()); parseErr == nil {
			body.SetScheduledDownload(t)
		}
	}
	if !plan.ScheduledUpgrade.IsNull() {
		if t, parseErr := time.Parse(time.RFC3339, plan.ScheduledUpgrade.ValueString()); parseErr == nil {
			body.SetScheduledUpgrade(t)
		}
	}
	_, err := r.client.ConnectorGroupAPI.UpdateConnectorGroupScheduledUpgrade(ctx, state.Oid.ValueString()).ConnectorGroupScheduledUpgrade(body).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error updating ConnectorGroupScheduledUpgrade", fmt.Sprintf("API error: %s", err.Error()))
		resp.Diagnostics.AddError("API Request Failed", utils.PrintScmError(err))
		return
	}
	result, _, err2 := r.client.ConnectorGroupAPI.GetConnectorGroupScheduledUpgrade(ctx, state.Oid.ValueString()).Execute()
	if err2 == nil && result != nil {
		plan.DrainTimeout = types.Int64Value(int64(result.GetDrainTimeout()))
		plan.ImageId = types.StringValue(result.GetImageId())
		plan.RollingUpgrade = types.BoolValue(result.GetRollingUpgrade())
		plan.ScheduledDownload = types.StringValue(result.GetScheduledDownload().Format(time.RFC3339))
		plan.ScheduledUpgrade = types.StringValue(result.GetScheduledUpgrade().Format(time.RFC3339))
	}
	plan.Tfid = state.Tfid
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *ConnectorGroupScheduledUpgradeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state connectorGroupScheduledUpgradeResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.client.ConnectorGroupAPI.DeleteConnectorGroupScheduledUpgrade(ctx, state.Oid.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error deleting ConnectorGroupScheduledUpgrade", fmt.Sprintf("API error: %s", err.Error()))
		resp.Diagnostics.AddError("API Request Failed", utils.PrintScmError(err))
	}
}
