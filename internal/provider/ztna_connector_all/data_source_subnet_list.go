package provider

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/ztna_connector_all"

	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/ztna_connector_all"
	"github.com/paloaltonetworks/terraform-provider-scm/internal/utils"
)

var (
	_ datasource.DataSource              = &SubnetListDataSource{}
	_ datasource.DataSourceWithConfigure = &SubnetListDataSource{}
)

func NewSubnetListDataSource() datasource.DataSource {
	return &SubnetListDataSource{}
}

// SubnetListDataSource defines the data source implementation.
type SubnetListDataSource struct {
	client *ztna_connector_all.APIClient
}

func (d *SubnetListDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "ztna_subnet_list"
}

func (d *SubnetListDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = models.SubnetsListDataSourceSchema
}

func (d *SubnetListDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	clients, ok := req.ProviderData.(map[string]interface{})
	if !ok {
		resp.Diagnostics.AddError("Unexpected Data Source Configure Type", fmt.Sprintf("Expected map[string]interface{}, got: %T.", req.ProviderData))
		return
	}
	client, ok := clients["ztna_connector_all"].(*ztna_connector_all.APIClient)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Client Type", fmt.Sprintf("Expected *ztna_connector_all.APIClient for 'ztna_connector_all' client."))
		return
	}
	d.client = client
}

func (d *SubnetListDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data models.SubnetsListModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Create the API request.
	listReq := d.client.SubnetAPI.ListSubnets(ctx)
	if !data.Offset.IsNull() {
		// START: Add dynamic query parameter handling
		tflog.Debug(ctx, "Applying filter", map[string]interface{}{"param": "offset", "value": data.Offset})
		listReq = listReq.Offset(int32(data.Offset.ValueInt64()))
		// END: Add dynamic query parameter handling
	}
	if !data.Limit.IsNull() {
		// START: Add dynamic query parameter handling
		tflog.Debug(ctx, "Applying filter", map[string]interface{}{"param": "limit", "value": data.Limit})
		listReq = listReq.Limit(int32(data.Limit.ValueInt64()))
		// END: Add dynamic query parameter handling
	}
	if !data.Sort.IsNull() {
		// START: Add dynamic query parameter handling
		tflog.Debug(ctx, "Applying filter", map[string]interface{}{"param": "sort", "value": data.Sort})
		listReq = listReq.Sort(data.Sort.ValueString())
		// END: Add dynamic query parameter handling
	}
	if !data.Search.IsNull() {
		// START: Add dynamic query parameter handling
		tflog.Debug(ctx, "Applying filter", map[string]interface{}{"param": "search", "value": data.Search})
		listReq = listReq.Search(data.Search.ValueString())
		// END: Add dynamic query parameter handling
	}
	if !data.Filters.IsNull() {
		// START: Add dynamic query parameter handling
		tflog.Debug(ctx, "Applying filter", map[string]interface{}{"param": "filters", "value": data.Filters})
		listReq = listReq.Filters(data.Filters.ValueString())
		// END: Add dynamic query parameter handling
	}

	// Execute the request.
	listResponse, _, err := listReq.Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error Listing Subnetss", fmt.Sprintf("Could not list Subnetss: %s", err.Error()))
		detailedMessage := utils.PrintScmError(err)
		resp.Diagnostics.AddError(
			"Resource Listing Failed: API Request Failed",
			detailedMessage,
		)
		return
	}

	// Convert the response to the Terraform model.
	if listResponse == nil || listResponse.GetData() == nil {
		return // Nothing to do.
	}

	total := int64(listResponse.GetTotal())
	data.Total = types.Int64PointerValue(&total)
	data.Limit = types.Int64Value(int64(listResponse.GetLimit()))
	data.Offset = types.Int64Value(int64(listResponse.GetOffset()))

	// =================== START: THE IMPROVEMENT ===================
	// Use the generated list packer to pack the SCM items into a TF list.
	packedList, diags := packSubnetsListFromSdk(ctx, listResponse.GetData())
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Convert the TF list to a Go slice of the model and set it to the data.
	resp.Diagnostics.Append(packedList.ElementsAs(ctx, &data.Data, false)...)
	if resp.Diagnostics.HasError() {
		return
	}
	// ===================  END: THE IMPROVEMENT  ===================

	// Create a unique TFID for this data source.
	var idBuilder strings.Builder

	// Use reflection again for Tfid creation to ensure safety

	idBuilder.WriteString(":")

	idBuilder.WriteString(":")

	idBuilder.WriteString(":")
	if !data.Name.IsNull() {
		idBuilder.WriteString(data.Name.ValueString())
	}
	idBuilder.WriteString(":")
	if !data.Limit.IsNull() {
		idBuilder.WriteString(strconv.FormatInt(data.Limit.ValueInt64(), 10))
	}
	idBuilder.WriteString(":")
	if !data.Offset.IsNull() {
		idBuilder.WriteString(strconv.FormatInt(data.Offset.ValueInt64(), 10))
	}
	data.Tfid = types.StringValue(idBuilder.String())

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
