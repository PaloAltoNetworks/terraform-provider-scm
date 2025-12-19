package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/network_services"

	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/network_services"
)

// DATA SOURCE for SCM AutoVpnSetting (Package: network_services)
var (
	_ datasource.DataSource              = &AutoVpnSettingDataSource{}
	_ datasource.DataSourceWithConfigure = &AutoVpnSettingDataSource{}
)

func NewAutoVpnSettingDataSource() datasource.DataSource {
	return &AutoVpnSettingDataSource{}
}

// AutoVpnSettingDataSource defines the data source implementation.
type AutoVpnSettingDataSource struct {
	client *network_services.APIClient
}

func (d *AutoVpnSettingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	tflog.Debug(ctx, "--- ENTER: AutoVpnSettingDataSource.Metadata ---")
	resp.TypeName = req.ProviderTypeName + "_auto_vpn_setting"
}

func (d *AutoVpnSettingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	tflog.Debug(ctx, "--- ENTER: AutoVpnSettingDataSource.Schema ---")
	// Use the pre-generated schema from the model file.
	resp.Schema = models.AutoVpnSettingsDataSourceSchema
}

func (d *AutoVpnSettingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	tflog.Debug(ctx, "--- ENTER: AutoVpnSettingsDataSource.Configure ---")
	if req.ProviderData == nil {
		return
	}
	clients, ok := req.ProviderData.(map[string]interface{})
	if !ok {
		resp.Diagnostics.AddError("Unexpected Data Source Configure Type", fmt.Sprintf("Expected map[string]interface{}, got: %T.", req.ProviderData))
		return
	}
	client, ok := clients["network_services"].(*network_services.APIClient)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Client Type", fmt.Sprintf("Expected *network_services.APIClient for 'network_services' client."))
		return
	}
	d.client = client
}

func (d *AutoVpnSettingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	tflog.Debug(ctx, "--- ENTER: AutoVpnSettingDataSource.Read ---")

	var data models.AutoVpnSettings

	// TF LOGGING ADDED: Log before the potentially crashing line.
	tflog.Debug(ctx, "--- VERIFICATION LOG: About to call req.Config.Get() ---")

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	// TF LOGGING ADDED: This log will ONLY print if the line above succeeds.
	tflog.Debug(ctx, "--- VERIFICATION LOG: Call to req.Config.Get() succeeded. ---")

	if resp.Diagnostics.HasError() {
		// TF LOGGING ADDED: Log if diagnostics has an error after Get().
		tflog.Debug(ctx, "--- VERIFICATION LOG: req.Config.Get() resulted in a diagnostic error.")
		return
	}
	// --- SINGLETON DATA SOURCE LOGIC ---
	tflog.Debug(ctx, "Reading Singleton AutoVpnSetting")

	// 1. Perform the API call (no ID argument)
	readReq := d.client.AutoVPNSettingsAPI.GetAutoVPNSettings(ctx)

	// 2. Add query parameters if any

	// 3. Execute using interface{} to capture any response type
	var scmObjectInterface interface{}
	scmObjectInterface, _, err := readReq.Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error Reading AutoVpnSetting", err.Error())
		return
	}

	// 4. Dynamic Response Handling (Reflection + JSON)
	var scmObject *network_services.AutoVpnSettings
	val := reflect.ValueOf(scmObjectInterface)
	if val.Kind() == reflect.Ptr && !val.IsNil() {
		val = val.Elem()
	}

	if val.Kind() == reflect.Struct {
		dataField := val.FieldByName("Data")
		if dataField.IsValid() && dataField.Kind() == reflect.Slice {
			if dataField.Len() > 0 {
				firstItem := dataField.Index(0).Interface()
				jsonBytes, _ := json.Marshal(firstItem)
				var targetStruct network_services.AutoVpnSettings
				if err := json.Unmarshal(jsonBytes, &targetStruct); err == nil {
					scmObject = &targetStruct
				}
			} else {
				resp.Diagnostics.AddError("Not Found", "The singleton resource was not found (empty list returned).")
				return
			}
		} else {
			jsonBytes, _ := json.Marshal(scmObjectInterface)
			var targetStruct network_services.AutoVpnSettings
			if err := json.Unmarshal(jsonBytes, &targetStruct); err == nil {
				scmObject = &targetStruct
			}
		}
	}

	if scmObject == nil {
		resp.Diagnostics.AddError("Error Processing Response", "Could not convert API response to expected model.")
		return
	}

	// 5. Pack and Set State
	packedObject, diags := packAutoVpnSettingsFromSdk(ctx, *scmObject)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(packedObject.As(ctx, &data, basetypes.ObjectAsOptions{})...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Force synthetic ID
	data.Tfid = types.StringValue("singleton_auto_vpn_settings")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
