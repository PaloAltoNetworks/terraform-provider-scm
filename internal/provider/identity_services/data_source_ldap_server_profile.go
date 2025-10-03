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
	"github.com/paloaltonetworks/scm-go/generated/identity_services"

	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/identity_services"
)

// DATA SOURCE for SCM LdapServerProfile (Package: identity_services)
var (
	_ datasource.DataSource              = &LdapServerProfileDataSource{}
	_ datasource.DataSourceWithConfigure = &LdapServerProfileDataSource{}
)

func NewLdapServerProfileDataSource() datasource.DataSource {
	return &LdapServerProfileDataSource{}
}

// LdapServerProfileDataSource defines the data source implementation.
type LdapServerProfileDataSource struct {
	client *identity_services.APIClient
}

func (d *LdapServerProfileDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	tflog.Debug(ctx, "--- ENTER: LdapServerProfileDataSource.Metadata ---")
	resp.TypeName = req.ProviderTypeName + "_ldap_server_profile"
}

func (d *LdapServerProfileDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	tflog.Debug(ctx, "--- ENTER: LdapServerProfileDataSource.Schema ---")
	// Use the pre-generated schema from the model file.
	resp.Schema = models.LdapServerProfilesDataSourceSchema
}

func (d *LdapServerProfileDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	tflog.Debug(ctx, "--- ENTER: LdapServerProfilesDataSource.Configure ---")
	if req.ProviderData == nil {
		return
	}
	clients, ok := req.ProviderData.(map[string]interface{})
	if !ok {
		resp.Diagnostics.AddError("Unexpected Data Source Configure Type", fmt.Sprintf("Expected map[string]interface{}, got: %T.", req.ProviderData))
		return
	}
	client, ok := clients["identity_services"].(*identity_services.APIClient)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Client Type", fmt.Sprintf("Expected *identity_services.APIClient for 'identity_services' client."))
		return
	}
	d.client = client
}

func (d *LdapServerProfileDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	tflog.Debug(ctx, "--- ENTER: LdapServerProfileDataSource.Read ---")

	var data models.LdapServerProfiles

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
		tflog.Debug(ctx, "Reading LdapServerProfiles data source by ID", map[string]interface{}{"id": objectId})

		getReq := d.client.LDAPServerProfilesAPI.GetLDAPServerProfilesByID(ctx, objectId)
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
			resp.Diagnostics.AddError("Error Reading LdapServerProfiles", fmt.Sprintf("Could not read LdapServerProfiles with ID %s: %s", objectId, err.Error()))
			return
		}
		if httpRes.StatusCode != 200 {
			resp.Diagnostics.AddError("Unexpected HTTP status code", fmt.Sprintf("Expected 200, got %d", httpRes.StatusCode))
			return
		}

		tflog.Debug(ctx, "--- DATA SOURCE READ: API call successful. About to call the packer.")

		// Create a packed object from the SCM response.
		packedObject, diags := packLdapServerProfilesFromSdk(ctx, *scmObject)
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
		tflog.Debug(ctx, "Reading LdapServerProfiles data source by Name", map[string]interface{}{"name": objectName})

		listReq := d.client.LDAPServerProfilesAPI.ListLDAPServerProfiles(ctx)

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
			resp.Diagnostics.AddError("Error Listing LdapServerProfiless", fmt.Sprintf("Could not list LdapServerProfiless: %s", err.Error()))
			return
		}
		if httpRes.StatusCode != 200 {
			resp.Diagnostics.AddError("Unexpected HTTP status code", fmt.Sprintf("Expected 200, got %d", httpRes.StatusCode))
			return
		}

		// Find the specific object from the list.
		var foundObject *identity_services.LdapServerProfiles
		for i := range listResponse.GetData() {
			item := listResponse.GetData()[i]
			if item.GetName() == objectName {
				foundObject = &item
				break
			}
		}

		if foundObject == nil {
			resp.Diagnostics.AddError("LdapServerProfiles Not Found", fmt.Sprintf("No LdapServerProfiles found with name: %s", objectName))
			return
		}

		// Create a packed object from the SCM response.
		packedObject, diags := packLdapServerProfilesFromSdk(ctx, *foundObject)
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
		resp.Diagnostics.AddError("Missing Identifier", "Either 'id' or 'name' must be provided for the LdapServerProfiles data source.")
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

	// Data sources don't store sensitive values, but the model requires the field to be properly initialized.
	data.EncryptedValues = types.MapNull(types.StringType)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
