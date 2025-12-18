package provider

import (
	"context"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/device_settings"

	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/device_settings"
	"github.com/paloaltonetworks/terraform-provider-scm/internal/utils"
)

var (
	_ datasource.DataSource              = &SessionSettingListDataSource{}
	_ datasource.DataSourceWithConfigure = &SessionSettingListDataSource{}
)

func NewSessionSettingListDataSource() datasource.DataSource {
	return &SessionSettingListDataSource{}
}

// SessionSettingListDataSource defines the data source implementation.
type SessionSettingListDataSource struct {
	client *device_settings.APIClient
}

func (d *SessionSettingListDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_session_setting_list"
}

func (d *SessionSettingListDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = models.SessionSettingsListDataSourceSchema
}

func (d *SessionSettingListDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	clients, ok := req.ProviderData.(map[string]interface{})
	if !ok {
		resp.Diagnostics.AddError("Unexpected Data Source Configure Type", fmt.Sprintf("Expected map[string]interface{}, got: %T.", req.ProviderData))
		return
	}
	client, ok := clients["device_settings"].(*device_settings.APIClient)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Client Type", fmt.Sprintf("Expected *device_settings.APIClient for 'device_settings' client."))
		return
	}
	d.client = client
}

func (d *SessionSettingListDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data models.SessionSettingsListModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Create the API request.
	listReq := d.client.SessionSettingsAPI.ListSessionSettings(ctx)
	if !data.Folder.IsNull() {
		// START: Add dynamic query parameter handling
		tflog.Debug(ctx, "Applying filter", map[string]interface{}{"param": "folder", "value": data.Folder})
		listReq = listReq.Folder(data.Folder.ValueString())
		// END: Add dynamic query parameter handling
	}
	if !data.Snippet.IsNull() {
		// START: Add dynamic query parameter handling
		tflog.Debug(ctx, "Applying filter", map[string]interface{}{"param": "snippet", "value": data.Snippet})
		listReq = listReq.Snippet(data.Snippet.ValueString())
		// END: Add dynamic query parameter handling
	}
	if !data.Device.IsNull() {
		// START: Add dynamic query parameter handling
		tflog.Debug(ctx, "Applying filter", map[string]interface{}{"param": "device", "value": data.Device})
		listReq = listReq.Device(data.Device.ValueString())
		// END: Add dynamic query parameter handling
	}

	// Execute the request.
	listResponse, _, err := listReq.Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error Listing SessionSettingss", fmt.Sprintf("Could not list SessionSettingss: %s", err.Error()))
		detailedMessage := utils.PrintScmError(err)
		resp.Diagnostics.AddError(
			"Resource Listing Failed: API Request Failed",
			detailedMessage,
		)
		return
	}

	// RAW LIST LOGIC:
	// The API returned a JSON array [...] directly.
	if listResponse == nil {
		return // Nothing to do.
	}

	// 1. TOTAL: We calculate it manually based on the slice length.
	total := int64(len(listResponse))
	data.Total = types.Int64PointerValue(&total)

	// 2. LIMIT/OFFSET: We do not update them from the server (metadata missing).
	//    We keep the values requested by the user in the config.

	// 3. PACKING: Pass the listResponse directly (it IS the slice).
	packedList, diags := packSessionSettingsListFromSdk(ctx, listResponse)

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

	v := reflect.ValueOf(data)

	if f := v.FieldByName("Folder"); f.IsValid() {
		if val, ok := f.Interface().(types.String); ok && !val.IsNull() {
			idBuilder.WriteString(val.ValueString())
		}
	}

	idBuilder.WriteString(":")

	if f := v.FieldByName("Snippet"); f.IsValid() {
		if val, ok := f.Interface().(types.String); ok && !val.IsNull() {
			idBuilder.WriteString(val.ValueString())
		}
	}

	idBuilder.WriteString(":")

	if f := v.FieldByName("Device"); f.IsValid() {
		if val, ok := f.Interface().(types.String); ok && !val.IsNull() {
			idBuilder.WriteString(val.ValueString())
		}
	}

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
