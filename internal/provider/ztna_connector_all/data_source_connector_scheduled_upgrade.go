package provider

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/paloaltonetworks/scm-go/generated/ztna_connector_all"

	"github.com/paloaltonetworks/terraform-provider-scm/internal/utils"
)

var (
	_ datasource.DataSource              = &ConnectorScheduledUpgradeDataSource{}
	_ datasource.DataSourceWithConfigure = &ConnectorScheduledUpgradeDataSource{}
)

func NewConnectorScheduledUpgradeDataSource() datasource.DataSource {
	return &ConnectorScheduledUpgradeDataSource{}
}

type ConnectorScheduledUpgradeDataSource struct {
	client *ztna_connector_all.APIClient
}

type connectorScheduledUpgradeModel struct {
	Oid               types.String `tfsdk:"oid"`
	ImageId           types.String `tfsdk:"image_id"`
	ScheduledDownload types.String `tfsdk:"scheduled_download"`
	ScheduledUpgrade  types.String `tfsdk:"scheduled_upgrade"`
	Tfid              types.String `tfsdk:"tfid"`
}

func (d *ConnectorScheduledUpgradeDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "ztna_connector_scheduled_upgrade"
}

func (d *ConnectorScheduledUpgradeDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Get Connector Scheduled Upgrade",
		Attributes: map[string]schema.Attribute{
			"oid": schema.StringAttribute{
				Description: "Parent connector OID.",
				Required:    true,
			},
			"image_id": schema.StringAttribute{
				Description: "The connector image version ID to upgrade to.",
				Computed:    true,
			},
			"scheduled_download": schema.StringAttribute{
				Description: "The scheduled download time in RFC3339 format (UTC).  If omitted, defaults to current UTC time.",
				Computed:    true,
			},
			"scheduled_upgrade": schema.StringAttribute{
				Description: "The scheduled upgrade time in RFC3339 format (UTC).  Must be after the scheduled_download time. If omitted, defaults to current UTC time.",
				Computed:    true,
			},
			"tfid": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (d *ConnectorScheduledUpgradeDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *ConnectorScheduledUpgradeDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data connectorScheduledUpgradeModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	result, _, err := d.client.ConnectorAPI.GetConnectorsScheduledUpgradeByID(ctx, data.Oid.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error reading ConnectorScheduledUpgrade", fmt.Sprintf("API error: %s", err.Error()))
		resp.Diagnostics.AddError("API Request Failed", utils.PrintScmError(err))
		return
	}
	if result == nil {
		resp.Diagnostics.AddError("Empty Response", "API returned nil result.")
		return
	}
	data.ImageId = types.StringValue(result.GetImageId())
	data.ScheduledDownload = types.StringValue(result.GetScheduledDownload().Format(time.RFC3339))
	data.ScheduledUpgrade = types.StringValue(result.GetScheduledUpgrade().Format(time.RFC3339))
	data.Tfid = types.StringValue(":::" + data.Oid.ValueString())
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
