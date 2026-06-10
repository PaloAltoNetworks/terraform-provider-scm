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
	_ datasource.DataSource              = &ConnectorQuiesceDataSource{}
	_ datasource.DataSourceWithConfigure = &ConnectorQuiesceDataSource{}
)

func NewConnectorQuiesceDataSource() datasource.DataSource {
	return &ConnectorQuiesceDataSource{}
}

type ConnectorQuiesceDataSource struct {
	client *ztna_connector_all.APIClient
}

type connectorQuiesceModel struct {
	Oid  types.String `tfsdk:"oid"`
	Mode types.String `tfsdk:"mode"`
	Tfid types.String `tfsdk:"tfid"`
}

func (d *ConnectorQuiesceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "ztna_connector_quiesce"
}

func (d *ConnectorQuiesceDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Get Connector Quiesce State",
		Attributes: map[string]schema.Attribute{
			"oid": schema.StringAttribute{
				Description: "Parent connector OID.",
				Required:    true,
			},
			"mode": schema.StringAttribute{
				Description: "",
				Computed:    true,
			},
			"tfid": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (d *ConnectorQuiesceDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *ConnectorQuiesceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data connectorQuiesceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	result, _, err := d.client.ConnectorAPI.GetConnectorsQuiesceByID(ctx, data.Oid.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error reading ConnectorQuiesce", fmt.Sprintf("API error: %s", err.Error()))
		resp.Diagnostics.AddError("API Request Failed", utils.PrintScmError(err))
		return
	}
	if result == nil {
		resp.Diagnostics.AddError("Empty Response", "API returned nil result.")
		return
	}
	data.Mode = types.StringValue(result.GetMode())
	data.Tfid = types.StringValue(":::" + data.Oid.ValueString())
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
