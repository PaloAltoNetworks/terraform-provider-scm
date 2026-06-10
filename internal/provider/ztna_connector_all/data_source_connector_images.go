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
	_ datasource.DataSource              = &ConnectorImagesDataSource{}
	_ datasource.DataSourceWithConfigure = &ConnectorImagesDataSource{}
)

func NewConnectorImagesDataSource() datasource.DataSource {
	return &ConnectorImagesDataSource{}
}

type ConnectorImagesDataSource struct {
	client *ztna_connector_all.APIClient
}

type connectorImagesModel struct {
	Values types.List   `tfsdk:"values"`
	Total  types.Int64  `tfsdk:"total"`
	Tfid   types.String `tfsdk:"tfid"`
}

func (d *ConnectorImagesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "ztna_connector_images"
}

func (d *ConnectorImagesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "List Connector Image Versions",
		Attributes: map[string]schema.Attribute{
			"values": schema.ListAttribute{
				Description: "List of values returned by the API.",
				Computed:    true,
				ElementType: types.StringType,
			},
			"total": schema.Int64Attribute{
				Description: "Number of values returned.",
				Computed:    true,
			},
			"tfid": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (d *ConnectorImagesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *ConnectorImagesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data connectorImagesModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	apiReq := d.client.ConnectorAPI.ListConnectorImages(ctx)

	values, _, err := apiReq.Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error calling ListConnectorImages", fmt.Sprintf("API error: %s", err.Error()))
		resp.Diagnostics.AddError("API Request Failed", utils.PrintScmError(err))
		return
	}

	listVal, diags := types.ListValueFrom(ctx, types.StringType, values)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Values = listVal
	data.Total = types.Int64Value(int64(len(values)))
	data.Tfid = types.StringValue("connector_images")
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
