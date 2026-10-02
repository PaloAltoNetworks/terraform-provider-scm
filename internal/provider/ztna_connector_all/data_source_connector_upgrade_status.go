package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/paloaltonetworks/scm-go/generated/ztna_connector_all"

	"github.com/paloaltonetworks/terraform-provider-scm/internal/utils"
)

var (
	_ datasource.DataSource              = &ConnectorUpgradeStatusDataSource{}
	_ datasource.DataSourceWithConfigure = &ConnectorUpgradeStatusDataSource{}
)

func NewConnectorUpgradeStatusDataSource() datasource.DataSource {
	return &ConnectorUpgradeStatusDataSource{}
}

type ConnectorUpgradeStatusDataSource struct {
	client *ztna_connector_all.APIClient
}

type connectorUpgradeStatusModel struct {
	Oid                types.String `tfsdk:"oid"`
	ActiveImageId      types.String `tfsdk:"active_image_id"`
	ActiveVersion      types.String `tfsdk:"active_version"`
	DownloadPercent    types.Int64  `tfsdk:"download_percent"`
	FailureInfo        types.String `tfsdk:"failure_info"`
	PreviousImageId    types.String `tfsdk:"previous_image_id"`
	ScheduledDownload  types.String `tfsdk:"scheduled_download"`
	ScheduledUpgrade   types.String `tfsdk:"scheduled_upgrade"`
	UpgradeDescription types.String `tfsdk:"upgrade_description"`
	UpgradeImageId     types.String `tfsdk:"upgrade_image_id"`
	UpgradeState       types.String `tfsdk:"upgrade_state"`
	Tfid               types.String `tfsdk:"tfid"`
}

func (d *ConnectorUpgradeStatusDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "ztna_connector_upgrade_status"
}

func (d *ConnectorUpgradeStatusDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Get Connector Scheduled Upgrade Status",
		Attributes: map[string]schema.Attribute{
			"oid": schema.StringAttribute{
				Description: "Parent connector OID.",
				Required:    true,
			},
			"active_image_id": schema.StringAttribute{
				Description: "Currently active image ID",
				Computed:    true,
			},
			"active_version": schema.StringAttribute{
				Description: "Currently active software version",
				Computed:    true,
			},
			"download_percent": schema.Int64Attribute{
				Description: "Download progress percentage",
				Computed:    true,
			},
			"failure_info": schema.StringAttribute{
				Description: "Failure information if upgrade failed",
				Computed:    true,
			},
			"previous_image_id": schema.StringAttribute{
				Description: "Previous image ID",
				Computed:    true,
			},
			"scheduled_download": schema.StringAttribute{
				Description: "Scheduled download time",
				Computed:    true,
			},
			"scheduled_upgrade": schema.StringAttribute{
				Description: "Scheduled upgrade time",
				Computed:    true,
			},
			"upgrade_description": schema.StringAttribute{
				Description: "Description of the upgrade",
				Computed:    true,
			},
			"upgrade_image_id": schema.StringAttribute{
				Description: "Target upgrade image ID",
				Computed:    true,
			},
			"upgrade_state": schema.StringAttribute{
				Description: "Current state of the upgrade process",
				Computed:    true,
			},
			"tfid": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (d *ConnectorUpgradeStatusDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	d.client = client
}

func (d *ConnectorUpgradeStatusDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data connectorUpgradeStatusModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	result, _, err := d.client.ConnectorAPI.GetConnectorsScheduledUpgradeStatusByID(ctx, data.Oid.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error reading ConnectorUpgradeStatus", fmt.Sprintf("API error: %s", err.Error()))
		resp.Diagnostics.AddError("API Request Failed", utils.PrintScmError(err))
		return
	}
	if result == nil {
		resp.Diagnostics.AddError("Empty Response", "API returned nil result.")
		return
	}
	data.ActiveImageId = types.StringValue(result.GetActiveImageId())
	data.ActiveVersion = types.StringValue(result.GetActiveVersion())
	data.DownloadPercent = types.Int64Value(int64(result.GetDownloadPercent()))
	data.FailureInfo = types.StringValue(result.GetFailureInfo())
	data.PreviousImageId = types.StringValue(result.GetPreviousImageId())
	data.ScheduledDownload = types.StringValue(result.GetScheduledDownload())
	data.ScheduledUpgrade = types.StringValue(result.GetScheduledUpgrade())
	data.UpgradeDescription = types.StringValue(result.GetUpgradeDescription())
	data.UpgradeImageId = types.StringValue(result.GetUpgradeImageId())
	data.UpgradeState = types.StringValue(result.GetUpgradeState())
	data.Tfid = types.StringValue(":::" + data.Oid.ValueString())
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
