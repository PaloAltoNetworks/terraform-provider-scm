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

	tfTypes "github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/paloaltonetworks/scm-go/generated/network_services"

	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/network_services"
	"github.com/paloaltonetworks/terraform-provider-scm/internal/utils"
)

// DATA SOURCE for SCM ZoneProtectionProfile (Package: network_services)
var (
	_ datasource.DataSource              = &ZoneProtectionProfileDataSource{}
	_ datasource.DataSourceWithConfigure = &ZoneProtectionProfileDataSource{}
)

func NewZoneProtectionProfileDataSource() datasource.DataSource {
	return &ZoneProtectionProfileDataSource{}
}

// ZoneProtectionProfileDataSource defines the data source implementation.
type ZoneProtectionProfileDataSource struct {
	client *network_services.APIClient
}

func (d *ZoneProtectionProfileDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	tflog.Debug(ctx, "--- ENTER: ZoneProtectionProfileDataSource.Metadata ---")
	resp.TypeName = req.ProviderTypeName + "_zone_protection_profile"
}

func (d *ZoneProtectionProfileDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	tflog.Debug(ctx, "--- ENTER: ZoneProtectionProfileDataSource.Schema ---")
	// Use the pre-generated schema from the model file.
	resp.Schema = models.ZoneProtectionProfilesDataSourceSchema
}

func (d *ZoneProtectionProfileDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	tflog.Debug(ctx, "--- ENTER: ZoneProtectionProfilesDataSource.Configure ---")
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

func (d *ZoneProtectionProfileDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	tflog.Debug(ctx, "--- ENTER: ZoneProtectionProfileDataSource.Read ---")

	var data models.ZoneProtectionProfiles

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
		tflog.Debug(ctx, "Reading ZoneProtectionProfiles data source by ID", map[string]interface{}{"id": objectId})

		getReq := d.client.ZoneProtectionProfilesAPI.GetZoneProtectionProfilesByID(ctx, objectId)
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
			resp.Diagnostics.AddError("Error Reading ZoneProtectionProfiles", fmt.Sprintf("Could not read ZoneProtectionProfiles with ID %s: %s", objectId, err.Error()))
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
		packedObject, diags := packZoneProtectionProfilesFromSdk(ctx, *scmObject)
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
		tflog.Debug(ctx, "Reading ZoneProtectionProfiles data source by Name", map[string]interface{}{"name": objectName})

		listReq := d.client.ZoneProtectionProfilesAPI.ListZoneProtectionProfiles(ctx)

		// Use reflection to dynamically check for and apply scope filters.

		v := reflect.ValueOf(data)

		if f := v.FieldByName("Folder"); f.IsValid() {
			if folder, ok := f.Interface().(tfTypes.String); ok && !folder.IsNull() {
				listReq = listReq.Folder(folder.ValueString())
			}
		}

		if f := v.FieldByName("Snippet"); f.IsValid() {
			if snippet, ok := f.Interface().(tfTypes.String); ok && !snippet.IsNull() {
				listReq = listReq.Snippet(snippet.ValueString())
			}
		}

		if f := v.FieldByName("Device"); f.IsValid() {
			if device, ok := f.Interface().(tfTypes.String); ok && !device.IsNull() {
				listReq = listReq.Device(device.ValueString())
			}
		}

		listResponse, httpRes, err := listReq.Execute()
		if err != nil {
			resp.Diagnostics.AddError("Error Listing ZoneProtectionProfiless", fmt.Sprintf("Could not list ZoneProtectionProfiless: %s", err.Error()))
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
		var foundObject *network_services.ZoneProtectionProfiles
		for i := range listResponse.GetData() {
			item := listResponse.GetData()[i]
			if item.GetName() == objectName {
				foundObject = &item
				break
			}
		}

		if foundObject == nil {
			resp.Diagnostics.AddError("ZoneProtectionProfiles Not Found", fmt.Sprintf("No ZoneProtectionProfiles found with name: %s", objectName))
			return
		}

		// Create a packed object from the SCM response.
		packedObject, diags := packZoneProtectionProfilesFromSdk(ctx, *foundObject)
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
		resp.Diagnostics.AddError("Missing Identifier", "Either 'id' or 'name' must be provided for the ZoneProtectionProfiles data source.")
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
