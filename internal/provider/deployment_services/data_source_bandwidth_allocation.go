package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/paloaltonetworks/scm-go/generated/deployment_services"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/deployment_services"
	"github.com/paloaltonetworks/terraform-provider-scm/internal/utils"
)

// DATA SOURCE: BandwidthAllocation (Specialized: List-Managed / No-ID)
var (
	_ datasource.DataSource              = &BandwidthAllocationDataSource{}
	_ datasource.DataSourceWithConfigure = &BandwidthAllocationDataSource{}
)

func NewBandwidthAllocationDataSource() datasource.DataSource {
	return &BandwidthAllocationDataSource{}
}

type BandwidthAllocationDataSource struct {
	client *deployment_services.APIClient
}

func (d *BandwidthAllocationDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_bandwidth_allocation"
}

func (d *BandwidthAllocationDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = models.BandwidthAllocationsDataSourceSchema
}

func (d *BandwidthAllocationDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	clients, ok := req.ProviderData.(map[string]interface{})
	if !ok {
		resp.Diagnostics.AddError("Unexpected Data Source Configure Type", fmt.Sprintf("Expected map[string]interface{}, got: %T.", req.ProviderData))
		return
	}
	client, ok := clients["deployment_services"].(*deployment_services.APIClient)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Client Type", fmt.Sprintf("Expected *deployment_services.APIClient for 'deployment_services' client."))
		return
	}
	d.client = client
}

func (d *BandwidthAllocationDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	tflog.Debug(ctx, "Starting Read function for No-ID Data Source BandwidthAllocation")
	var config models.BandwidthAllocations
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// We require the 'name' to find the specific item
	if config.Name.IsNull() {
		resp.Diagnostics.AddError("Missing Attribute", "The 'name' attribute is required for this data source.")
		return
	}
	targetName := config.Name.ValueString()

	// 1. Call the List operation to get all items
	listReq := d.client.BandwidthAllocationsAPI.ListBandwidthAllocations(ctx)

	listResponse, _, err := listReq.Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error reading BandwidthAllocation", err.Error())
		detailedMessage := utils.PrintScmError(err)
		resp.Diagnostics.AddError(
			"Resource Listing Failed: API Request Failed",
			detailedMessage,
		)
		return
	}

	// 2. Iterate through the list to find the matching item
	var foundItem *deployment_services.BandwidthAllocations

	if listResponse != nil && listResponse.Data != nil {
		for _, item := range listResponse.Data {
			if item.Name != "" && item.Name == targetName {
				foundItem = &item
				break
			}
		}
	}

	if foundItem == nil {
		resp.Diagnostics.AddError("Not Found", fmt.Sprintf("BandwidthAllocation with name '%s' was not found.", targetName))
		return
	}

	// 3. Pack the found item into the model
	packedObject, diags := packBandwidthAllocationsFromSdk(ctx, *foundItem)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state models.BandwidthAllocations
	resp.Diagnostics.Append(packedObject.As(ctx, &state, basetypes.ObjectAsOptions{})...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set the Tfid to the name
	state.Tfid = state.Name

	// Initialize sensitive fields if present (required by framework)

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}
