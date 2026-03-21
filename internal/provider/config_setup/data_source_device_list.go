package provider

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/config_setup"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/config_setup"
	"github.com/paloaltonetworks/terraform-provider-scm/internal/utils"
)

var (
	_ datasource.DataSource              = &DeviceListDataSource{}
	_ datasource.DataSourceWithConfigure = &DeviceListDataSource{}
)

func NewDeviceListDataSource() datasource.DataSource {
	return &DeviceListDataSource{}
}

// DeviceListDataSource defines the data source implementation.
type DeviceListDataSource struct {
	client *config_setup.APIClient
}

func (d *DeviceListDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_list"
}

func (d *DeviceListDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = models.DevicesListDataSourceSchema
}

func (d *DeviceListDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	clients, ok := req.ProviderData.(map[string]interface{})
	if !ok {
		resp.Diagnostics.AddError("Unexpected Data Source Configure Type", fmt.Sprintf("Expected map[string]interface{}, got: %T.", req.ProviderData))
		return
	}
	client, ok := clients["config_setup"].(*config_setup.APIClient)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Client Type", "Expected *config_setup.APIClient for 'config_setup' client.")
		return
	}
	d.client = client
}

func (d *DeviceListDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data models.DevicesListModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Create the API request.
	listReq := d.client.DevicesAPI.ListDevices(ctx)

	if !data.Limit.IsNull() {
		tflog.Debug(ctx, "Applying filter", map[string]interface{}{"param": "limit", "value": data.Limit})
		listReq = listReq.Limit(int32(data.Limit.ValueInt64()))
	}
	if !data.Offset.IsNull() {
		tflog.Debug(ctx, "Applying filter", map[string]interface{}{"param": "offset", "value": data.Offset})
		listReq = listReq.Offset(int32(data.Offset.ValueInt64()))
	}
	if !data.Folder.IsNull() {
		tflog.Debug(ctx, "Applying filter", map[string]interface{}{"param": "folder", "value": data.Folder})
		listReq = listReq.Folder(data.Folder.ValueString())
	}

	// Execute the request.
	listResponse, _, err := listReq.Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error Listing Devices", fmt.Sprintf("Could not list Devices: %s", err.Error()))
		detailedMessage := utils.PrintScmError(err)
		resp.Diagnostics.AddError("Resource Listing Failed: API Request Failed", detailedMessage)
		return
	}

	if listResponse == nil || listResponse.GetData() == nil {
		return
	}

	total := int64(listResponse.GetTotal())
	data.Total = types.Int64PointerValue(&total)
	data.Limit = types.Int64Value(int64(listResponse.GetLimit()))
	data.Offset = types.Int64Value(int64(listResponse.GetOffset()))

	// Pack the list response.
	packedList, diags := packDevicesListFromSdk(ctx, listResponse.GetData())
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(packedList.ElementsAs(ctx, &data.Data, false)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Create a unique TFID for this data source.
	var idBuilder strings.Builder
	idBuilder.WriteString(":::")
	if !data.Folder.IsNull() {
		idBuilder.WriteString(data.Folder.ValueString())
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
