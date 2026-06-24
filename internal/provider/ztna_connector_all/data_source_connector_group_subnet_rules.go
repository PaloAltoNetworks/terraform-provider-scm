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
	_ datasource.DataSource              = &ConnectorGroupSubnetRulesDataSource{}
	_ datasource.DataSourceWithConfigure = &ConnectorGroupSubnetRulesDataSource{}
)

func NewConnectorGroupSubnetRulesDataSource() datasource.DataSource {
	return &ConnectorGroupSubnetRulesDataSource{}
}

type ConnectorGroupSubnetRulesDataSource struct {
	client *ztna_connector_all.APIClient
}

type connectorGroupSubnetRulesModel struct {
	Oid     types.String `tfsdk:"oid"`
	Group   types.String `tfsdk:"group"`
	Subnets types.String `tfsdk:"subnets"`
	Tfid    types.String `tfsdk:"tfid"`
}

func (d *ConnectorGroupSubnetRulesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "ztna_connector_group_subnet_rules"
}

func (d *ConnectorGroupSubnetRulesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "List Subnets per Connector Group",
		Attributes: map[string]schema.Attribute{
			"oid": schema.StringAttribute{
				Description: "Parent connector OID.",
				Required:    true,
			},
			"group": schema.StringAttribute{
				Description: "",
				Computed:    true,
			},
			"subnets": schema.StringAttribute{
				Description: "",
				Computed:    true,
			},
			"tfid": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (d *ConnectorGroupSubnetRulesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *ConnectorGroupSubnetRulesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data connectorGroupSubnetRulesModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	result, _, err := d.client.ConnectorGroupAPI.ListConnectorGroupSubnets(ctx, data.Oid.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error reading ConnectorGroupSubnetRules", fmt.Sprintf("API error: %s", err.Error()))
		resp.Diagnostics.AddError("API Request Failed", utils.PrintScmError(err))
		return
	}
	if result == nil {
		resp.Diagnostics.AddError("Empty Response", "API returned nil result.")
		return
	}
	if jsonBytesGroup, jsonErrGroup := json.Marshal(result.GetGroup()); jsonErrGroup == nil {
		data.Group = types.StringValue(string(jsonBytesGroup))
	} else {
		data.Group = types.StringValue("")
	}
	if jsonBytesSubnets, jsonErrSubnets := json.Marshal(result.GetSubnets()); jsonErrSubnets == nil {
		data.Subnets = types.StringValue(string(jsonBytesSubnets))
	} else {
		data.Subnets = types.StringValue("")
	}
	data.Tfid = types.StringValue(":::" + data.Oid.ValueString())
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
