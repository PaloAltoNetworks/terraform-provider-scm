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

	"github.com/paloaltonetworks/scm-go/generated/ztna_connector_all"

	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/ztna_connector_all"
)

// DATA SOURCE for SCM TenantStatus (Package: ztna_connector_all)
var (
	_ datasource.DataSource              = &TenantStatusDataSource{}
	_ datasource.DataSourceWithConfigure = &TenantStatusDataSource{}
)

func NewTenantStatusDataSource() datasource.DataSource {
	return &TenantStatusDataSource{}
}

// TenantStatusDataSource defines the data source implementation.
type TenantStatusDataSource struct {
	client *ztna_connector_all.APIClient
}

func (d *TenantStatusDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	tflog.Debug(ctx, "--- ENTER: TenantStatusDataSource.Metadata ---")
	resp.TypeName = "ztna_tenant_status"
}

func (d *TenantStatusDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	tflog.Debug(ctx, "--- ENTER: TenantStatusDataSource.Schema ---")
	// Use the pre-generated schema from the model file.
	resp.Schema = models.TenantStatusDataSourceSchema
}

func (d *TenantStatusDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	tflog.Debug(ctx, "--- ENTER: TenantStatusDataSource.Configure ---")
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

func (d *TenantStatusDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	tflog.Debug(ctx, "--- ENTER: TenantStatusDataSource.Read ---")

	var data models.TenantStatus

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
	tflog.Debug(ctx, "Reading Singleton TenantStatus")

	// 1. Perform the API call (no ID argument)
	readReq := d.client.TenantAPI.GetTenantStatus(ctx)

	// 2. Add query parameters if any

	// 3. Execute using interface{} to capture any response type
	var scmObjectInterface interface{}
	scmObjectInterface, _, err := readReq.Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error Reading TenantStatus", err.Error())
		return
	}

	// 4. Dynamic Response Handling (Reflection + JSON)
	var scmObject *ztna_connector_all.TenantStatus
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
				var targetStruct ztna_connector_all.TenantStatus
				if err := json.Unmarshal(jsonBytes, &targetStruct); err == nil {
					scmObject = &targetStruct
				}
			} else {
				resp.Diagnostics.AddError("Not Found", "The singleton resource was not found (empty list returned).")
				return
			}
		} else {
			jsonBytes, _ := json.Marshal(scmObjectInterface)
			var targetStruct ztna_connector_all.TenantStatus
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
	packedObject, diags := packTenantStatusFromSdk(ctx, *scmObject)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(packedObject.As(ctx, &data, basetypes.ObjectAsOptions{})...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Force synthetic ID
	data.Tfid = types.StringValue("singleton_tenant_status")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
