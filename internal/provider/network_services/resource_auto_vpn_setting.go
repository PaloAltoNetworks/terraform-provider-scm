package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/network_services"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/network_services"
	"github.com/paloaltonetworks/terraform-provider-scm/internal/utils"
)

// SINGLETON RESOURCE for SCM AutoVpnSetting (Package: network_services)
var (
	_ resource.Resource                = &AutoVpnSettingResource{}
	_ resource.ResourceWithConfigure   = &AutoVpnSettingResource{}
	_ resource.ResourceWithImportState = &AutoVpnSettingResource{}
)

func NewAutoVpnSettingResource() resource.Resource {
	return &AutoVpnSettingResource{}
}

// AutoVpnSettingResource defines the resource implementation.
type AutoVpnSettingResource struct {
	client *network_services.APIClient
}

func (r *AutoVpnSettingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_auto_vpn_setting"
}

func (r *AutoVpnSettingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = models.AutoVpnSettingsResourceSchema
}

func (r *AutoVpnSettingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	clients, ok := req.ProviderData.(map[string]interface{})
	if !ok {
		resp.Diagnostics.AddError("Unexpected Resource Configure Type", fmt.Sprintf("Expected map[string]interface{}, got: %T.", req.ProviderData))
		return
	}
	client, ok := clients["network_services"].(*network_services.APIClient)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Client Type", fmt.Sprintf("Expected *network_services.APIClient for 'network_services' client."))
		return
	}
	r.client = client
}

func (r *AutoVpnSettingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Starting Create function for Singleton AutoVpnSetting")
	var data models.AutoVpnSettings

	// 1. Get the plan from Terraform into the data model.
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Unpack the plan to an SCM SDK object.
	planObject, diags := types.ObjectValueFrom(ctx, models.AutoVpnSettings{}.AttrTypes(), &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// 2. Unpack the request BODY from data into an SDK object.
	unpackedScmObject, diags := unpackAutoVpnSettingsToSdk(ctx, planObject)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating auto_vpn_settings on SCM API")

	// --- START: DYNAMIC CREATE LOGIC ---
	var scmObjectInterface interface{}
	var err error

	// This singleton uses PUT for Create (e.g., bgp-routing)
	tflog.Debug(ctx, "Using PUT operation for Create: UpdateAutoVPNSettings")

	// --- START AUTOMATED CONVERSION (JSON Marshaling) ---
	// Asymmetric Schema Detected: The API expects auto-vpn-settings, but we have AutoVpnSettings.
	var updateObj network_services.AutoVpnSettings

	// 1. Marshal the read object to JSON
	jsonBytes, err := json.Marshal(unpackedScmObject)
	if err != nil {
		resp.Diagnostics.AddError("Error Converting Model", "Could not marshal read model to JSON: "+err.Error())
		return
	}

	// 2. Unmarshal into the update object. This ignores extra fields.
	err = json.Unmarshal(jsonBytes, &updateObj)
	if err != nil {
		resp.Diagnostics.AddError("Error Converting Model", "Could not unmarshal JSON to update model: "+err.Error())
		return
	}

	createReq := r.client.AutoVPNSettingsAPI.UpdateAutoVPNSettings(ctx).AutoVpnSettings(updateObj)
	// --- END AUTOMATED CONVERSION ---

	scmObjectInterface, _, err = createReq.Execute()

	// --- END: DYNAMIC CREATE LOGIC ---

	if err != nil {
		resp.Diagnostics.AddError("Error creating auto_vpn_settings", err.Error())
		detailedMessage := utils.PrintScmError(err)

		resp.Diagnostics.AddError(
			"SCM Resource Creation Failed: API Request Failed",
			detailedMessage,
		)
		return
	}

	// --- START MODIFICATION: DYNAMIC RESPONSE HANDLING (Applied to Create) ---
	var createdObject *network_services.AutoVpnSettings

	// Use reflection to inspect the return type at runtime.
	val := reflect.ValueOf(scmObjectInterface)
	if val.Kind() == reflect.Ptr && !val.IsNil() {
		val = val.Elem()
	}

	if val.Kind() == reflect.Struct {
		// 1. Check for "Data" field (List Wrapper pattern)
		dataField := val.FieldByName("Data")
		if dataField.IsValid() && dataField.Kind() == reflect.Slice {
			// It is a list response (e.g., saas-tenant-restrictions)
			if dataField.Len() > 0 {
				// Extract the first item
				firstItem := dataField.Index(0).Interface()

				// We need to convert this interface{} back to the concrete SDK struct type.
				jsonBytes, _ := json.Marshal(firstItem)
				var targetStruct network_services.AutoVpnSettings
				if err := json.Unmarshal(jsonBytes, &targetStruct); err == nil {
					createdObject = &targetStruct
				}
			} else {
				// List is empty
			}
		} else {
			// 2. Not a list wrapper, assume it's the direct object (e.g., bgp-routing)
			// Use the same JSON trick to be safe against pointer mismatches
			jsonBytes, _ := json.Marshal(scmObjectInterface)
			var targetStruct network_services.AutoVpnSettings
			if err := json.Unmarshal(jsonBytes, &targetStruct); err == nil {
				createdObject = &targetStruct
			}
		}
	}

	if createdObject == nil {
		// If API returned 200/201 but no object, we cannot set the state.
		resp.Diagnostics.AddError("API Response Missing Object", "API call successful but returned no object to set state with. Check SCM logs.")
		return
	}

	packedObject, diags := packAutoVpnSettingsFromSdk(ctx, *createdObject)
	// --- END MODIFICATION ---

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(packedObject.As(ctx, &data, basetypes.ObjectAsOptions{})...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Tfid = types.StringValue("singleton")

	tflog.Debug(ctx, "Created auto_vpn_settings", map[string]interface{}{"tfid": data.Tfid.ValueString()})
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AutoVpnSettingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Step 1 - Initialize a data and savestate of type models.AutoVpnSettings - which is the TF schema struct
	tflog.Debug(ctx, "Starting Read function for Singleton AutoVpnSetting")
	var data, savestate models.AutoVpnSettings

	// Step 2 - Fetch the state into savestate
	resp.Diagnostics.Append(req.State.Get(ctx, &savestate)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading auto_vpn_settings from SCM API")
	getReq := r.client.AutoVPNSettingsAPI.GetAutoVPNSettings(ctx)

	// --- START: DYNAMIC PARAMETERS FOR READ ---
	// --- END: DYNAMIC PARAMETERS FOR READ ---

	// Use interface{} to handle flexible return types (List vs Object)
	var scmObjectInterface interface{}
	var httpErr *http.Response
	var err error
	scmObjectInterface, httpErr, err = getReq.Execute()

	if err != nil {
		if httpErr != nil && httpErr.StatusCode == http.StatusNotFound {
			// FIX: Remove undefined variable objectId from log
			tflog.Debug(ctx, "Got no auto_vpn_settings on read SCM API. Remove from state to let terraform create", map[string]interface{}{"tfid": savestate.Tfid.ValueString()})
			resp.State.RemoveResource(ctx)
		} else {
			// FIX: Remove undefined variable objectId from log
			tflog.Debug(ctx, "Got an exception on read SCM API.", map[string]interface{}{"tfid": savestate.Tfid.ValueString()})
			resp.Diagnostics.AddError("Error reading auto_vpn_settings", err.Error())
			detailedMessage := utils.PrintScmError(err)
			resp.Diagnostics.AddError("SCM Resource Read Failed: API Request Failed", detailedMessage)
		}
		return
	}

	// --- START MODIFICATION: DYNAMIC RESPONSE HANDLING (Applied to Read) ---
	var scmObject *network_services.AutoVpnSettings

	// Use reflection to inspect the return type at runtime.
	val := reflect.ValueOf(scmObjectInterface)
	if val.Kind() == reflect.Ptr && !val.IsNil() {
		val = val.Elem()
	}

	if val.Kind() == reflect.Struct {
		// 1. Check for "Data" field (List Wrapper pattern)
		dataField := val.FieldByName("Data")
		if dataField.IsValid() && dataField.Kind() == reflect.Slice {
			// It is a list response (e.g., saas-tenant-restrictions)
			if dataField.Len() > 0 {
				// Extract the first item
				firstItem := dataField.Index(0).Interface()

				// We need to convert this interface{} back to the concrete SDK struct type.
				jsonBytes, _ := json.Marshal(firstItem)
				var targetStruct network_services.AutoVpnSettings
				if err := json.Unmarshal(jsonBytes, &targetStruct); err == nil {
					scmObject = &targetStruct
				}
			} else {
				// List is empty, treat as Not Found
				tflog.Debug(ctx, "Got no auto_vpn_settings on read SCM API (empty list). Remove from state", map[string]interface{}{"tfid": savestate.Tfid.ValueString()})
				resp.State.RemoveResource(ctx)
				return
			}
		} else {
			// 2. Not a list wrapper, assume it's the direct object (e.g., bgp-routing)
			// Use the same JSON trick to be safe against pointer mismatches
			jsonBytes, _ := json.Marshal(scmObjectInterface)
			var targetStruct network_services.AutoVpnSettings
			if err := json.Unmarshal(jsonBytes, &targetStruct); err == nil {
				scmObject = &targetStruct
			}
		}
	}

	// If after all checks, scmObject is still nil, then the API returned 200 OK but no object was found.
	if scmObject == nil {
		tflog.Debug(ctx, "Got no auto_vpn_settings on read SCM API (nil object). Remove from state to let terraform create", map[string]interface{}{"tfid": savestate.Tfid.ValueString()})
		resp.State.RemoveResource(ctx)
		return
	}

	// Step 5 - Pack the scm object into a terraform model and put it in data we initialized in step 1
	packedObject, diags := packAutoVpnSettingsFromSdk(ctx, *scmObject)
	// --- END MODIFICATION ---

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(packedObject.As(ctx, &data, basetypes.ObjectAsOptions{})...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Step 7 - Carry over tfid from state back into data
	data.Tfid = savestate.Tfid

	// Step 8 - Set things in params back into data object from the savestate - things like position of security rule

	// --- START: RESTORE READ PARAMETERS ---
	// --- END: RESTORE READ PARAMETERS ---

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AutoVpnSettingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	// Step 1: Initialize a plan and state of type models.AutoVpnSettings which is the terraform schema struct
	tflog.Debug(ctx, "Starting Update function for Singleton AutoVpnSetting")
	var plan, state models.AutoVpnSettings

	// Step 2: Get the plan from plan file (resource.tf) into plan and state from tfstate into state
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Step 3: Creates a plan object from the plan
	planObject, diags := types.ObjectValueFrom(ctx, models.AutoVpnSettings{}.AttrTypes(), &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Step 4: Unpack the plan object to an SCM Object
	unpackedScmObject, diags := unpackAutoVpnSettingsToSdk(ctx, planObject)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// --- START: MODIFIED API CALL (no path param) ---
	var scmObjectInterface interface{}
	var httpErr *http.Response
	var err error

	tflog.Debug(ctx, "Updating auto_vpn_settings on SCM API")

	// --- START AUTOMATED CONVERSION (JSON Marshaling) ---
	// Asymmetric Schema Detected: The API expects auto-vpn-settings, but we have AutoVpnSettings.
	var updateObj network_services.AutoVpnSettings

	// 1. Marshal the read object to JSON
	jsonBytes, err := json.Marshal(unpackedScmObject)
	if err != nil {
		resp.Diagnostics.AddError("Error Converting Model", "Could not marshal read model to JSON: "+err.Error())
		return
	}

	// 2. Unmarshal into the update object.
	err = json.Unmarshal(jsonBytes, &updateObj)
	if err != nil {
		resp.Diagnostics.AddError("Error Converting Model", "Could not unmarshal JSON to update model: "+err.Error())
		return
	}

	updateReq := r.client.AutoVPNSettingsAPI.UpdateAutoVPNSettings(ctx).AutoVpnSettings(updateObj)
	// --- END AUTOMATED CONVERSION ---

	scmObjectInterface, httpErr, err = updateReq.Execute()

	// --- END: MODIFIED API CALL ---

	if err != nil {
		if httpErr != nil && httpErr.StatusCode == http.StatusNotFound {
			// FIX: Remove undefined variable objectId from log
			tflog.Debug(ctx, "Got no auto_vpn_settings on update SCM API. Remove from state to let terraform create", map[string]interface{}{"tfid": state.Tfid.ValueString()})
			resp.State.RemoveResource(ctx)
		} else {
			// FIX: Remove undefined variable objectId from log
			tflog.Debug(ctx, "Got an exception on update SCM API.", map[string]interface{}{"tfid": state.Tfid.ValueString()})
			resp.Diagnostics.AddError("Error updating auto_vpn_settings", err.Error())
			detailedMessage := utils.PrintScmError(err)
			resp.Diagnostics.AddError("SCM Resource Update Failed: API Request Failed", detailedMessage)
		}
		return
	}

	// --- START MODIFICATION: DYNAMIC RESPONSE HANDLING (Applied to Update) ---
	var updatedObject *network_services.AutoVpnSettings

	// Use reflection to inspect the return type at runtime.
	val := reflect.ValueOf(scmObjectInterface)
	if val.Kind() == reflect.Ptr && !val.IsNil() {
		val = val.Elem()
	}

	if val.Kind() == reflect.Struct {
		// 1. Check for "Data" field (List Wrapper pattern)
		dataField := val.FieldByName("Data")
		if dataField.IsValid() && dataField.Kind() == reflect.Slice {
			// It is a list response (e.g., saas-tenant-restrictions)
			if dataField.Len() > 0 {
				// Extract the first item
				firstItem := dataField.Index(0).Interface()

				// We need to convert this interface{} back to the concrete SDK struct type.
				jsonBytes, _ := json.Marshal(firstItem)
				var targetStruct network_services.AutoVpnSettings
				if err := json.Unmarshal(jsonBytes, &targetStruct); err == nil {
					updatedObject = &targetStruct
				}
			} else {
				// List is empty
			}
		} else {
			// 2. Not a list wrapper, assume it's the direct object (e.g., bgp-routing)
			// Use the same JSON trick to be safe against pointer mismatches
			jsonBytes, _ := json.Marshal(scmObjectInterface)
			var targetStruct network_services.AutoVpnSettings
			if err := json.Unmarshal(jsonBytes, &targetStruct); err == nil {
				updatedObject = &targetStruct
			}
		}
	}

	if updatedObject == nil {
		// If API returned 200/201 but no object, something is wrong, fail gracefully.
		resp.Diagnostics.AddError("API Response Missing Object", "API call successful but returned no object after update. Check SCM logs.")
		return
	}

	packedObject, diags := packAutoVpnSettingsFromSdk(ctx, *updatedObject)
	// --- END MODIFICATION ---

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(packedObject.As(ctx, &plan, basetypes.ObjectAsOptions{})...)
	if resp.Diagnostics.HasError() {
		return
	}

	// // Preserve any write-only parameter values from the plan.
	//

	// Step 10: Carry over tfid from state into plan
	plan.Tfid = state.Tfid

	// Step 11: Copy write-only attributes from the prior state to the plan for things like position in security rule

	// --- START: RESTORE READ PARAMETERS ---
	// --- END: RESTORE READ PARAMETERS ---

	tflog.Debug(ctx, "Updated auto_vpn_settings", map[string]interface{}{"tfid": plan.Tfid.ValueString()})
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *AutoVpnSettingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// --- START: DYNAMIC DELETE LOGIC ---

	// If no delete operation, we just remove from state.
	tflog.Warn(ctx, "auto_vpn_settings is a singleton and cannot be deleted from the API. Removing from state.")

	// --- END: DYNAMIC DELETE LOGIC ---
}

func (r *AutoVpnSettingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// --- START: MODIFIED IMPORT ---
	// We expect the ID to be "singleton" or the resource name.
	if req.ID != "singleton" && req.ID != "auto_vpn_setting" {
		resp.Diagnostics.AddError("Unexpected Import Identifier", fmt.Sprintf("Expected 'singleton' or 'auto_vpn_setting', got: %s", req.ID))
		return
	}
	// All singleton imports map to a static "singleton" tfid.
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("tfid"), types.StringValue("singleton"))...)
	// --- END: MODIFIED IMPORT ---
}
