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
	Field           types.String `tfsdk:"field"`
	Search          types.String `tfsdk:"search"`
	Values          types.List   `tfsdk:"values"`
	FreeFormFilters types.List   `tfsdk:"free_form_filters"`
	StaticFilters   types.List   `tfsdk:"static_filters"`
	Tfid            types.String `tfsdk:"tfid"`
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
				Description: "Filter values when a specific field is queried.",
				Computed:    true,
				ElementType: types.StringType,
			},
			"free_form_filters": schema.ListAttribute{
				Description: "Free-form filter field names (returned when no field is specified).",
				Computed:    true,
				ElementType: types.StringType,
			},
			"static_filters": schema.ListAttribute{
				Description: "Static filter field names (returned when no field is specified).",
				Computed:    true,
				ElementType: types.StringType,
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
		// When no field param is set the API returns a categorised object instead of a string
		// array. Detect this by trying to unmarshal the raw error body as the object format.
		if data.Field.IsNull() {
			if bodyErr, ok := err.(interface{ Body() []byte }); ok {
				var filterObj struct {
					FreeFormFilters []string `json:"free form filters"`
					StaticFilters   []string `json:"static filters"`
				}
				if json.Unmarshal(bodyErr.Body(), &filterObj) == nil {
					freeFormList, diags := types.ListValueFrom(ctx, types.StringType, filterObj.FreeFormFilters)
					resp.Diagnostics.Append(diags...)
					staticList, diags := types.ListValueFrom(ctx, types.StringType, filterObj.StaticFilters)
					resp.Diagnostics.Append(diags...)
					emptyList, diags := types.ListValueFrom(ctx, types.StringType, []string{})
					resp.Diagnostics.Append(diags...)
					if resp.Diagnostics.HasError() {
						return
					}
					data.Values = emptyList
					data.FreeFormFilters = freeFormList
					data.StaticFilters = staticList
					data.Tfid = types.StringValue("discovered_application_filters")
					resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
					return
				}
			}
		}
		resp.Diagnostics.AddError("Error calling ListDiscoveredApplicationFilters", fmt.Sprintf("API error: %s", err.Error()))
		resp.Diagnostics.AddError("API Request Failed", utils.PrintScmError(err))
		return
	}

	// field was specified: response is a string array of matching values
	valueList, diags := types.ListValueFrom(ctx, types.StringType, values)
	resp.Diagnostics.Append(diags...)
	emptyList, diags := types.ListValueFrom(ctx, types.StringType, []string{})
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Values = valueList
	data.FreeFormFilters = emptyList
	data.StaticFilters = emptyList
	data.Tfid = types.StringValue("discovered_application_filters")
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
