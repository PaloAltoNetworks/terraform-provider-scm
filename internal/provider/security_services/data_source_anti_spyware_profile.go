package provider

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/security_services"

	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/security_services"
	"github.com/paloaltonetworks/terraform-provider-scm/internal/utils"
)

// DATA SOURCE for SCM AntiSpywareProfile (Package: security_services)
var (
	_ datasource.DataSource              = &AntiSpywareProfileDataSource{}
	_ datasource.DataSourceWithConfigure = &AntiSpywareProfileDataSource{}
)

func NewAntiSpywareProfileDataSource() datasource.DataSource {
	return &AntiSpywareProfileDataSource{}
}

// AntiSpywareProfileDataSource defines the data source implementation.
type AntiSpywareProfileDataSource struct {
	client *security_services.APIClient
}

func (d *AntiSpywareProfileDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	tflog.Debug(ctx, "--- ENTER: AntiSpywareProfileDataSource.Metadata ---")
	resp.TypeName = req.ProviderTypeName + "_anti_spyware_profile"
}

func (d *AntiSpywareProfileDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	tflog.Debug(ctx, "--- ENTER: AntiSpywareProfileDataSource.Schema ---")
	// Use the pre-generated schema from the model file.
	resp.Schema = models.AntiSpywareProfilesDataSourceSchema
}

func (d *AntiSpywareProfileDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	tflog.Debug(ctx, "--- ENTER: AntiSpywareProfilesDataSource.Configure ---")
	if req.ProviderData == nil {
		return
	}
	clients, ok := req.ProviderData.(map[string]interface{})
	if !ok {
		resp.Diagnostics.AddError("Unexpected Data Source Configure Type", fmt.Sprintf("Expected map[string]interface{}, got: %T.", req.ProviderData))
		return
	}
	client, ok := clients["security_services"].(*security_services.APIClient)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Client Type", fmt.Sprintf("Expected *security_services.APIClient for 'security_services' client."))
		return
	}
	d.client = client
}

func (d *AntiSpywareProfileDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	tflog.Debug(ctx, "--- ENTER: AntiSpywareProfileDataSource.Read ---")

	var data models.AntiSpywareProfiles

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
		tflog.Debug(ctx, "Reading AntiSpywareProfiles data source by ID", map[string]interface{}{"id": objectId})

		getReq := d.client.AntiSpywareProfilesAPI.GetAntiSpywareProfilesByID(ctx, objectId)
		scmObject, httpRes, err := getReq.Execute()

		var statusCode int
		if httpRes != nil {
			statusCode = httpRes.StatusCode
		}
		tflog.Debug(ctx, "--- SCM API CALL: Execution complete.", map[string]interface{}{
			"error":          err,
			"statusCode":     statusCode,
			"scmObjectIsNil": scmObject == nil,
		})

		if err != nil {
			resp.Diagnostics.AddError("Error Reading AntiSpywareProfiles", fmt.Sprintf("Could not read AntiSpywareProfiles with ID %s: %s", objectId, err.Error()))
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
		packedObject, diags := packAntiSpywareProfilesFromSdk(ctx, *scmObject)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

		// Load the packed object into the data model.
		resp.Diagnostics.Append(packedObject.As(ctx, &data, basetypes.ObjectAsOptions{})...)
		if resp.Diagnostics.HasError() {
			return
		}

	} else if !data.Name.IsNull() {
		objectName := data.Name.ValueString()
		tflog.Debug(ctx, "Reading AntiSpywareProfiles data source by Name", map[string]interface{}{"name": objectName})

		listReq := d.client.AntiSpywareProfilesAPI.ListAntiSpywareProfiles(ctx)
		if !data.Folder.IsNull() {
			tflog.Debug(ctx, "Applying filter", map[string]interface{}{"param": "folder", "value": data.Folder})
			listReq = listReq.Folder(data.Folder.ValueString())
		}
		if !data.Snippet.IsNull() {
			tflog.Debug(ctx, "Applying filter", map[string]interface{}{"param": "snippet", "value": data.Snippet})
			listReq = listReq.Snippet(data.Snippet.ValueString())
		}
		if !data.Device.IsNull() {
			tflog.Debug(ctx, "Applying filter", map[string]interface{}{"param": "device", "value": data.Device})
			listReq = listReq.Device(data.Device.ValueString())
		}

		listResponse, httpRes, err := listReq.Execute()
		if err != nil {
			resp.Diagnostics.AddError("Error Listing AntiSpywareProfiless", fmt.Sprintf("Could not list AntiSpywareProfiless: %s", err.Error()))
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

		// Find the specific object from the list.
		var foundObject *security_services.AntiSpywareProfiles
		for i := range listResponse.GetData() {
			item := listResponse.GetData()[i]
			if item.GetName() == objectName {
				foundObject = &item
				break
			}
		}

		if foundObject == nil {
			resp.Diagnostics.AddError("AntiSpywareProfiles Not Found", fmt.Sprintf("No AntiSpywareProfiles found with name: %s", objectName))
			return
		}

		// Create a packed object from the SCM response.
		packedObject, diags := packAntiSpywareProfilesFromSdk(ctx, *foundObject)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

		// Load the packed object into the data model.
		resp.Diagnostics.Append(packedObject.As(ctx, &data, basetypes.ObjectAsOptions{})...)
		if resp.Diagnostics.HasError() {
			return
		}

	} else {
		resp.Diagnostics.AddError("Missing Identifier", "Either 'id' or 'name' must be provided for the AntiSpywareProfiles data source.")
		return
	}

	// Create the composite Tfid for consistency.
	var idBuilder strings.Builder

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
	idBuilder.WriteString(data.Id.ValueString())
	data.Tfid = types.StringValue(idBuilder.String())

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
