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
	"github.com/paloaltonetworks/scm-go/generated/objects"

	"github.com/paloaltonetworks/terraform-provider-scm/internal/models/objects"
)

// DATA SOURCE for SCM QuarantinedDevice (Package: objects)
var (
	_ datasource.DataSource              = &QuarantinedDeviceDataSource{}
	_ datasource.DataSourceWithConfigure = &QuarantinedDeviceDataSource{}
)

func NewQuarantinedDeviceDataSource() datasource.DataSource {
	return &QuarantinedDeviceDataSource{}
}

// QuarantinedDeviceDataSource defines the data source implementation.
type QuarantinedDeviceDataSource struct {
	client *objects.APIClient
}

func (d *QuarantinedDeviceDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	tflog.Debug(ctx, "--- ENTER: QuarantinedDeviceDataSource.Metadata ---")
	resp.TypeName = req.ProviderTypeName + "_quarantined_device"
}

func (d *QuarantinedDeviceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	tflog.Debug(ctx, "--- ENTER: QuarantinedDeviceDataSource.Schema ---")
	// Use the pre-generated schema from the model file.
	resp.Schema = models.QuarantinedDevicesDataSourceSchema
}

func (d *QuarantinedDeviceDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	tflog.Debug(ctx, "--- ENTER: QuarantinedDevicesDataSource.Configure ---")
	if req.ProviderData == nil {
		return
	}
	clients, ok := req.ProviderData.(map[string]interface{})
	if !ok {
		resp.Diagnostics.AddError("Unexpected Data Source Configure Type", fmt.Sprintf("Expected map[string]interface{}, got: %T.", req.ProviderData))
		return
	}
	client, ok := clients["objects"].(*objects.APIClient)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Client Type", fmt.Sprintf("Expected *objects.APIClient for 'objects' client."))
		return
	}
	d.client = client
}

func (d *QuarantinedDeviceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	tflog.Debug(ctx, "--- ENTER: QuarantinedDeviceDataSource.Read ---")

	var data models.QuarantinedDevices

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
	tflog.Debug(ctx, "Reading Singleton QuarantinedDevice")

	// 1. Perform the API call (no ID argument)
	readReq := d.client.QuarantinedDevicesAPI.ListQuarantinedDevices(ctx)

	// 2. Add query parameters if any
	if !data.HostId.IsNull() {
		readReq = readReq.HostId(data.HostId.ValueString())
	}
	if !data.SerialNumber.IsNull() {
		readReq = readReq.SerialNumber(data.SerialNumber.ValueString())
	}

	// 3. Execute using interface{} to capture any response type
	var scmObjectInterface interface{}
	scmObjectInterface, _, err := readReq.Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error Reading QuarantinedDevice", err.Error())
		return
	}

	// 4. Dynamic Response Handling (Reflection + JSON)
	var scmObject *objects.QuarantinedDevices
	val := reflect.ValueOf(scmObjectInterface)
	if val.Kind() == reflect.Ptr && !val.IsNil() { val = val.Elem() }

	if val.Kind() == reflect.Struct {
		dataField := val.FieldByName("Data")
		if dataField.IsValid() && dataField.Kind() == reflect.Slice {
			if dataField.Len() > 0 {
				firstItem := dataField.Index(0).Interface()
				jsonBytes, _ := json.Marshal(firstItem)
				var targetStruct objects.QuarantinedDevices
				if err := json.Unmarshal(jsonBytes, &targetStruct); err == nil {
					scmObject = &targetStruct
				}
			} else {
				resp.Diagnostics.AddError("Not Found", "The singleton resource was not found (empty list returned).")
				return
			}
		} else {
            jsonBytes, _ := json.Marshal(scmObjectInterface)
            var targetStruct objects.QuarantinedDevices
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
	packedObject, diags := packQuarantinedDevicesFromSdk(ctx, *scmObject)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() { return }

	resp.Diagnostics.Append(packedObject.As(ctx, &data, basetypes.ObjectAsOptions{})...)
	if resp.Diagnostics.HasError() { return }

	// Force synthetic ID
	data.Tfid = types.StringValue("singleton_quarantined_devices")



	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

*/
