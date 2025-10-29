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

	// Logic to handle read by ID or by name/list.
	// We prioritize reading by ID if it is provided.
	if !data.Id.IsNull() {
		objectId := data.Id.ValueString()
		tflog.Debug(ctx, "Reading QuarantinedDevices data source by ID", map[string]interface{}{"id": objectId})

		getReq := d.client.QuarantinedDevicesAPI.(ctx, objectId)
		scmObject, httpRes, err := getReq.Execute()

		var statusCode int
		if httpRes != nil {
			statusCode = httpRes.StatusCode
		}
		tflog.Debug(ctx, "--- SCM API CALL: Execution complete.", map[string]interface{}{
			"error":      err,
			"statusCode": statusCode,
			"scmObjectIsNil": scmObject == nil,
		})

		if err != nil {
			resp.Diagnostics.AddError("Error Reading QuarantinedDevices", fmt.Sprintf("Could not read QuarantinedDevices with ID %s: %s", objectId, err.Error()))
			detailedMessage := utils.PrintScmError(err)
			resp.Diagnostics.AddError(
				"Tag Listing Failed: API Request Failed",
				detailedMessage,
			)
			return
		}
		if httpRes.StatusCode != 200 {
			resp.Diagnostics.AddError("Unexpected HTTP status code", fmt.Sprintf("Expected 200, got %d", httpRes.StatusCode))
			return
		}

		tflog.Debug(ctx, "--- DATA SOURCE READ: API call successful. About to call the packer.")

		// Create a packed object from the SCM response.
		packedObject, diags := packQuarantinedDevicesFromSdk(ctx, *scmObject)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

		// Load the packed object into the data model.
		resp.Diagnostics.Append(packedObject.As(ctx, &data, basetypes.ObjectAsOptions{})...)
		if resp.Diagnostics.HasError() {
			return
		}

	}  else {
		resp.Diagnostics.AddError("Missing Identifier", "Either 'id' or 'name' must be provided for the QuarantinedDevices data source.")
		return
	}

	// Create the composite Tfid for consistency.
	var idBuilder strings.Builder


	idBuilder.WriteString(":")

	idBuilder.WriteString(":")

	idBuilder.WriteString(":")
	idBuilder.WriteString(data.Id.ValueString())
	data.Tfid = types.StringValue(idBuilder.String())



	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

*/
