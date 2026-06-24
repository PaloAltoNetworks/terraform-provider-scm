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
	_ datasource.DataSource              = &DiscoveredApplicationFiltersDataSource{}
	_ datasource.DataSourceWithConfigure = &DiscoveredApplicationFiltersDataSource{}
)

func NewDiscoveredApplicationFiltersDataSource() datasource.DataSource {
	return &DiscoveredApplicationFiltersDataSource{}
}

type DiscoveredApplicationFiltersDataSource struct {
	client *ztna_connector_all.APIClient
}

type discoveredApplicationFiltersModel struct {
	Field  types.String `tfsdk:"field"`
	Search types.String `tfsdk:"search"`
	Values types.List   `tfsdk:"values"`
	Total  types.Int64  `tfsdk:"total"`
	Tfid   types.String `tfsdk:"tfid"`
}

func (d *DiscoveredApplicationFiltersDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "ztna_discovered_application_filters"
}

func (d *DiscoveredApplicationFiltersDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "List Discovered Application Filters",
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

func (d *DiscoveredApplicationFiltersDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *DiscoveredApplicationFiltersDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data discoveredApplicationFiltersModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	apiReq := d.client.DiscoveredApplicationAPI.ListDiscoveredApplicationFilters(ctx)
	if !data.Field.IsNull() {
		apiReq = apiReq.Field(data.Field.ValueString())
	}
	if !data.Search.IsNull() {
		apiReq = apiReq.Search(data.Search.ValueString())
	}

	values, _, err := apiReq.Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error calling ListDiscoveredApplicationFilters", fmt.Sprintf("API error: %s", err.Error()))
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
	data.Tfid = types.StringValue("discovered_application_filters")
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
