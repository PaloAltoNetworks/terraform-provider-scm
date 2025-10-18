package provider

import (
	"context"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/paloaltonetworks/scm-go/generated/deployment_services"

	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/deployment_services"
	"github.com/paloaltonetworks/terraform-provider-scm/internal/utils"
)

var (
	_ datasource.DataSource              = &ServiceConnectionGroupListDataSource{}
	_ datasource.DataSourceWithConfigure = &ServiceConnectionGroupListDataSource{}
)

func NewServiceConnectionGroupListDataSource() datasource.DataSource {
	return &ServiceConnectionGroupListDataSource{}
}

// ServiceConnectionGroupListDataSource defines the data source implementation.
type ServiceConnectionGroupListDataSource struct {
	client *deployment_services.APIClient
}

func (d *ServiceConnectionGroupListDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service_connection_group_list"
}

func (d *ServiceConnectionGroupListDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = models.ServiceConnectionGroupsListDataSourceSchema
}

func (d *ServiceConnectionGroupListDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *ServiceConnectionGroupListDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data models.ServiceConnectionGroupsListModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Create the API request.
	listReq := d.client.ServiceConnectionGroupsAPI.ListServiceConnectionGroups(ctx)

	// Apply filters from the configuration.

	v := reflect.ValueOf(data)

	if f := v.FieldByName("Folder"); f.IsValid() {
		if val, ok := f.Interface().(types.String); ok && !val.IsNull() {
			listReq = listReq.Folder(val.ValueString())
		}
	}

	if !data.Limit.IsNull() {
		listReq = listReq.Limit(int32(data.Limit.ValueInt64()))
	}
	if !data.Offset.IsNull() {
		listReq = listReq.Offset(int32(data.Offset.ValueInt64()))
	}

	// Execute the request.
	listResponse, _, err := listReq.Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error Listing ServiceConnectionGroupss", fmt.Sprintf("Could not list ServiceConnectionGroupss: %s", err.Error()))
		detailedMessage := utils.PrintScmError(err)
		resp.Diagnostics.AddError(
			"Tag Listing Failed: API Request Failed",
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
	packedList, diags := packServiceConnectionGroupsListFromSdk(ctx, listResponse.GetData())
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

	v = reflect.ValueOf(data)

	if f := v.FieldByName("Folder"); f.IsValid() {
		if val, ok := f.Interface().(types.String); ok && !val.IsNull() {
			idBuilder.WriteString(val.ValueString())
		}
	}

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
