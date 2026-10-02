package provider

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/paloaltonetworks/scm-go/generated/ztna_connector_all"

	"github.com/paloaltonetworks/terraform-provider-scm/internal/utils"
)

var (
	_ datasource.DataSource              = &ConnectorGroupUpgradeStatusDataSource{}
	_ datasource.DataSourceWithConfigure = &ConnectorGroupUpgradeStatusDataSource{}
)

func NewConnectorGroupUpgradeStatusDataSource() datasource.DataSource {
	return &ConnectorGroupUpgradeStatusDataSource{}
}

type ConnectorGroupUpgradeStatusDataSource struct {
	client *ztna_connector_all.APIClient
}

type connectorGroupUpgradeStatusModel struct {
	Oid            types.String `tfsdk:"oid"`
	Data           types.String `tfsdk:"data"`
	Name           types.String `tfsdk:"name"`
	RollingUpgrade types.Bool   `tfsdk:"rolling_upgrade"`
	UpgradeStatus  types.String `tfsdk:"upgrade_status"`
	Tfid           types.String `tfsdk:"tfid"`
}

func (d *ConnectorGroupUpgradeStatusDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "ztna_connector_group_upgrade_status"
}

func (d *ConnectorGroupUpgradeStatusDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Get Connector Group Scheduled Upgrade Status",
		Attributes: map[string]schema.Attribute{
			"oid": schema.StringAttribute{
				Description: "Parent connector OID.",
				Required:    true,
			},
			"data": schema.StringAttribute{
				Description: "List of connector upgrade statuses within this group",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "Connector group name",
				Computed:    true,
			},
			"rolling_upgrade": schema.BoolAttribute{
				Description: "Whether rolling upgrade is enabled for this connector group",
				Computed:    true,
			},
			"upgrade_status": schema.StringAttribute{
				Description: "Overall upgrade status for the connector group",
				Computed:    true,
			},
			"tfid": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (d *ConnectorGroupUpgradeStatusDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *ConnectorGroupUpgradeStatusDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data connectorGroupUpgradeStatusModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	result, _, err := d.client.ConnectorGroupAPI.GetConnectorGroupScheduledUpgradeStatus(ctx, data.Oid.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error reading ConnectorGroupUpgradeStatus", fmt.Sprintf("API error: %s", err.Error()))
		resp.Diagnostics.AddError("API Request Failed", utils.PrintScmError(err))
		return
	}
	if result == nil {
		resp.Diagnostics.AddError("Empty Response", "API returned nil result.")
		return
	}
	if jsonBytesData, jsonErrData := json.Marshal(result.GetData()); jsonErrData == nil {
		data.Data = types.StringValue(string(jsonBytesData))
	} else {
		data.Data = types.StringValue("")
	}
	data.Name = types.StringValue(result.GetName())
	data.RollingUpgrade = types.BoolValue(result.GetRollingUpgrade())
	data.UpgradeStatus = types.StringValue(result.GetUpgradeStatus())
	data.Tfid = types.StringValue(":::" + data.Oid.ValueString())
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
