package provider

/*

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	tfTypes "github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/paloaltonetworks/scm-go/generated/network_services"

	"github.com/paloaltonetworks/terraform-provider-scm/internal/models/network_services"
)

// DATA SOURCE for SCM SslDecryptionSetting (Package: network_services)
var (
	_ datasource.DataSource              = &SslDecryptionSettingDataSource{}
	_ datasource.DataSourceWithConfigure = &SslDecryptionSettingDataSource{}
)

func NewSslDecryptionSettingDataSource() datasource.DataSource {
	return &SslDecryptionSettingDataSource{}
}

// SslDecryptionSettingDataSource defines the data source implementation.
type SslDecryptionSettingDataSource struct {
	client *network_services.APIClient
}

func (d *SslDecryptionSettingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	tflog.Debug(ctx, "--- ENTER: SslDecryptionSettingDataSource.Metadata ---")
	resp.TypeName = req.ProviderTypeName + "_ssl_decryption_setting"
}

func (d *SslDecryptionSettingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	tflog.Debug(ctx, "--- ENTER: SslDecryptionSettingDataSource.Schema ---")
	// Use the pre-generated schema from the model file.
	resp.Schema = models.SslDecryptionSettingsDataSourceSchema
}

func (d *SslDecryptionSettingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	tflog.Debug(ctx, "--- ENTER: SslDecryptionSettingsDataSource.Configure ---")
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

func (d *SslDecryptionSettingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	tflog.Debug(ctx, "--- ENTER: SslDecryptionSettingDataSource.Read ---")

	var data models.SslDecryptionSettings

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
	tflog.Debug(ctx, "Reading Singleton SslDecryptionSetting")

	// 1. Perform the API call (no ID argument)
	readReq := d.client.SslDecryptionSettingsAPI.getSslDecryptionSettings(ctx)

	// 2. Add query parameters if any
	if !data.Folder.IsNull() {
		readReq = readReq.Folder(data.Folder.ValueString())
	}
	if !data.Snippet.IsNull() {
		readReq = readReq.Snippet(data.Snippet.ValueString())
	}
	if !data.Device.IsNull() {
		readReq = readReq.Device(data.Device.ValueString())
	}
	if !data.Offset.IsNull() {
		readReq = readReq.Offset(int32(data.Offset.ValueInt64()))
	}
	if !data.Limit.IsNull() {
		readReq = readReq.Limit(int32(data.Limit.ValueInt64()))
	}

	// 3. Execute using interface{} to capture any response type
	var scmObjectInterface interface{}
	scmObjectInterface, _, err := readReq.Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error Reading SslDecryptionSetting", err.Error())
		return
	}

	// 4. Dynamic Response Handling (Reflection + JSON)
	var scmObject *network_services.SslDecryptionSettings
	val := reflect.ValueOf(scmObjectInterface)
	if val.Kind() == reflect.Ptr && !val.IsNil() { val = val.Elem() }

	if val.Kind() == reflect.Struct {
		dataField := val.FieldByName("Data")
		if dataField.IsValid() && dataField.Kind() == reflect.Slice {
			if dataField.Len() > 0 {
				firstItem := dataField.Index(0).Interface()
				jsonBytes, _ := json.Marshal(firstItem)
				var targetStruct network_services.SslDecryptionSettings
				if err := json.Unmarshal(jsonBytes, &targetStruct); err == nil {
					scmObject = &targetStruct
				}
			} else {
				resp.Diagnostics.AddError("Not Found", "The singleton resource was not found (empty list returned).")
				return
			}
		} else {
            jsonBytes, _ := json.Marshal(scmObjectInterface)
            var targetStruct network_services.SslDecryptionSettings
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
	packedObject, diags := packSslDecryptionSettingsFromSdk(ctx, *scmObject)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() { return }

	resp.Diagnostics.Append(packedObject.As(ctx, &data, basetypes.ObjectAsOptions{})...)
	if resp.Diagnostics.HasError() { return }

	// Force synthetic ID
	data.Tfid = types.StringValue("singleton_ssl_decryption_settings")



	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

*/
