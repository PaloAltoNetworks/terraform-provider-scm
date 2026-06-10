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
	_ datasource.DataSource              = &SubnetFiltersDataSource{}
	_ datasource.DataSourceWithConfigure = &SubnetFiltersDataSource{}
)

func NewSubnetFiltersDataSource() datasource.DataSource {
	return &SubnetFiltersDataSource{}
}

type SubnetFiltersDataSource struct {
	client *ztna_connector_all.APIClient
}

type subnetFiltersModel struct {
	Field  types.String `tfsdk:"field"`
	Search types.String `tfsdk:"search"`
	Values types.List   `tfsdk:"values"`
	Total  types.Int64  `tfsdk:"total"`
	Tfid   types.String `tfsdk:"tfid"`
}

func (d *SubnetFiltersDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "ztna_subnet_filters"
}

func (d *SubnetFiltersDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "List Subnet Filters",
		Attributes: map[string]schema.Attribute{
			"field": schema.StringAttribute{
				Description: "String that represents a static filter field. Call any of the /filters endpoints without specifying a field to get a list of all available fields.",
				Optional:    true,
			},
			"search": schema.StringAttribute{
				Description: "String to filter list results by. Is searched over multiple fields in each object. Multiple searches can be specified.",
				Optional:    true,
			},
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

func (d *SubnetFiltersDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *SubnetFiltersDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data subnetFiltersModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	apiReq := d.client.SubnetAPI.ListSubnetFilters(ctx)
	if !data.Field.IsNull() {
		apiReq = apiReq.Field(data.Field.ValueString())
	}
	if !data.Search.IsNull() {
		apiReq = apiReq.Search(data.Search.ValueString())
	}

	values, _, err := apiReq.Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error calling ListSubnetFilters", fmt.Sprintf("API error: %s", err.Error()))
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
	data.Tfid = types.StringValue("subnet_filters")
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
