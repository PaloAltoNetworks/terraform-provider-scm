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

	"github.com/paloaltonetworks/scm-go/generated/deployment_services"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/deployment_services"
	"github.com/paloaltonetworks/terraform-provider-scm/internal/utils"
)

// SINGLETON RESOURCE for SCM BgpRouting (Package: deployment_services)
var (
	_ resource.Resource                = &BgpRoutingResource{}
	_ resource.ResourceWithConfigure   = &BgpRoutingResource{}
	_ resource.ResourceWithImportState = &BgpRoutingResource{}
)

func NewBgpRoutingResource() resource.Resource {
	return &BgpRoutingResource{}
}

// BgpRoutingResource defines the resource implementation.
type BgpRoutingResource struct {
	client *deployment_services.APIClient
}

func (r *BgpRoutingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_bgp_routing"
}

func (r *BgpRoutingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = models.BgpRoutingResourceSchema
}

func (r *BgpRoutingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	clients, ok := req.ProviderData.(map[string]interface{})
	if !ok {
		resp.Diagnostics.AddError("Unexpected Resource Configure Type", fmt.Sprintf("Expected map[string]interface{}, got: %T.", req.ProviderData))
		return
	}
	client, ok := clients["deployment_services"].(*deployment_services.APIClient)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Client Type", fmt.Sprintf("Expected *deployment_services.APIClient for 'deployment_services' client."))
		return
	}
	r.client = client
}

func (r *BgpRoutingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Starting Create function for Singleton BgpRouting")
	var data models.BgpRouting

	// 1. Get the plan from Terraform into the data model.
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Unpack the plan to an SCM SDK object.
	planObject, diags := types.ObjectValueFrom(ctx, models.BgpRouting{}.AttrTypes(), &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// 2. Unpack the request BODY from data into an SDK object.
	unpackedScmObject, diags := unpackBgpRoutingToSdk(ctx, planObject)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating bgp_routing on SCM API")

	// --- START: DYNAMIC CREATE LOGIC ---
	var scmObjectInterface interface{}
	var err error

	// This singleton uses PUT for Create (e.g., bgp-routing)
	tflog.Debug(ctx, "Using PUT operation for Create: UpdateBGPRouting")

	// --- START AUTOMATED CONVERSION (JSON Marshaling) ---
	// Asymmetric Schema Detected: The API expects bgp-routing, but we have BgpRouting.
	var updateObj deployment_services.BgpRouting

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

	createReq := r.client.BGPRoutingAPI.UpdateBGPRouting(ctx).BgpRouting(updateObj)
	// --- END AUTOMATED CONVERSION ---

	scmObjectInterface, _, err = createReq.Execute()

	// --- END: DYNAMIC CREATE LOGIC ---

	if err != nil {
		resp.Diagnostics.AddError("Error creating bgp_routing", err.Error())
		detailedMessage := utils.PrintScmError(err)

		resp.Diagnostics.AddError(
			"SCM Resource Creation Failed: API Request Failed",
			detailedMessage,
		)
		return
	}

	// --- START MODIFICATION: DYNAMIC RESPONSE HANDLING (Applied to Create) ---
	var createdObject *deployment_services.BgpRouting

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
				var targetStruct deployment_services.BgpRouting
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
			var targetStruct deployment_services.BgpRouting
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

	packedObject, diags := packBgpRoutingFromSdk(ctx, *createdObject)
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

	tflog.Debug(ctx, "Created bgp_routing", map[string]interface{}{"tfid": data.Tfid.ValueString()})
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BgpRoutingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Step 1 - Initialize a data and savestate of type models.BgpRouting - which is the TF schema struct
	tflog.Debug(ctx, "Starting Read function for Singleton BgpRouting")
	var data, savestate models.BgpRouting

	// Step 2 - Fetch the state into savestate
	resp.Diagnostics.Append(req.State.Get(ctx, &savestate)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading bgp_routing from SCM API")
	getReq := r.client.BGPRoutingAPI.GetBGPRouting(ctx)

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
			tflog.Debug(ctx, "Got no bgp_routing on read SCM API. Remove from state to let terraform create", map[string]interface{}{"tfid": savestate.Tfid.ValueString()})
			resp.State.RemoveResource(ctx)
		} else {
			// FIX: Remove undefined variable objectId from log
			tflog.Debug(ctx, "Got an exception on read SCM API.", map[string]interface{}{"tfid": savestate.Tfid.ValueString()})
			resp.Diagnostics.AddError("Error reading bgp_routing", err.Error())
			detailedMessage := utils.PrintScmError(err)
			resp.Diagnostics.AddError("SCM Resource Read Failed: API Request Failed", detailedMessage)
		}
		return
	}

	// --- START MODIFICATION: DYNAMIC RESPONSE HANDLING (Applied to Read) ---
	var scmObject *deployment_services.BgpRouting

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
				var targetStruct deployment_services.BgpRouting
				if err := json.Unmarshal(jsonBytes, &targetStruct); err == nil {
					scmObject = &targetStruct
				}
			} else {
				// List is empty, treat as Not Found
				tflog.Debug(ctx, "Got no bgp_routing on read SCM API (empty list). Remove from state", map[string]interface{}{"tfid": savestate.Tfid.ValueString()})
				resp.State.RemoveResource(ctx)
				return
			}
		} else {
			// 2. Not a list wrapper, assume it's the direct object (e.g., bgp-routing)
			// Use the same JSON trick to be safe against pointer mismatches
			jsonBytes, _ := json.Marshal(scmObjectInterface)
			var targetStruct deployment_services.BgpRouting
			if err := json.Unmarshal(jsonBytes, &targetStruct); err == nil {
				scmObject = &targetStruct
			}
		}
	}

	// If after all checks, scmObject is still nil, then the API returned 200 OK but no object was found.
	if scmObject == nil {
		tflog.Debug(ctx, "Got no bgp_routing on read SCM API (nil object). Remove from state to let terraform create", map[string]interface{}{"tfid": savestate.Tfid.ValueString()})
		resp.State.RemoveResource(ctx)
		return
	}

	// Step 5 - Pack the scm object into a terraform model and put it in data we initialized in step 1
	packedObject, diags := packBgpRoutingFromSdk(ctx, *scmObject)
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

func (r *BgpRoutingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	// Step 1: Initialize a plan and state of type models.BgpRouting which is the terraform schema struct
	tflog.Debug(ctx, "Starting Update function for Singleton BgpRouting")
	var plan, state models.BgpRouting

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
	planObject, diags := types.ObjectValueFrom(ctx, models.BgpRouting{}.AttrTypes(), &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Step 4: Unpack the plan object to an SCM Object
	unpackedScmObject, diags := unpackBgpRoutingToSdk(ctx, planObject)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// --- START: MODIFIED API CALL (no path param) ---
	var scmObjectInterface interface{}
	var httpErr *http.Response
	var err error

	tflog.Debug(ctx, "Updating bgp_routing on SCM API")

	// --- START AUTOMATED CONVERSION (JSON Marshaling) ---
	// Asymmetric Schema Detected: The API expects bgp-routing, but we have BgpRouting.
	var updateObj deployment_services.BgpRouting

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

	updateReq := r.client.BGPRoutingAPI.UpdateBGPRouting(ctx).BgpRouting(updateObj)
	// --- END AUTOMATED CONVERSION ---

	scmObjectInterface, httpErr, err = updateReq.Execute()

	// --- END: MODIFIED API CALL ---

	if err != nil {
		if httpErr != nil && httpErr.StatusCode == http.StatusNotFound {
			// FIX: Remove undefined variable objectId from log
			tflog.Debug(ctx, "Got no bgp_routing on update SCM API. Remove from state to let terraform create", map[string]interface{}{"tfid": state.Tfid.ValueString()})
			resp.State.RemoveResource(ctx)
		} else {
			// FIX: Remove undefined variable objectId from log
			tflog.Debug(ctx, "Got an exception on update SCM API.", map[string]interface{}{"tfid": state.Tfid.ValueString()})
			resp.Diagnostics.AddError("Error updating bgp_routing", err.Error())
			detailedMessage := utils.PrintScmError(err)
			resp.Diagnostics.AddError("SCM Resource Update Failed: API Request Failed", detailedMessage)
		}
		return
	}

	// --- START MODIFICATION: DYNAMIC RESPONSE HANDLING (Applied to Update) ---
	var updatedObject *deployment_services.BgpRouting

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
				var targetStruct deployment_services.BgpRouting
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
			var targetStruct deployment_services.BgpRouting
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

	packedObject, diags := packBgpRoutingFromSdk(ctx, *updatedObject)
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

	tflog.Debug(ctx, "Updated bgp_routing", map[string]interface{}{"tfid": plan.Tfid.ValueString()})
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *BgpRoutingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// --- START: DYNAMIC DELETE LOGIC ---

	// If no delete operation, we just remove from state.
	tflog.Warn(ctx, "bgp_routing is a singleton and cannot be deleted from the API. Removing from state.")

	// --- END: DYNAMIC DELETE LOGIC ---
}

func (r *BgpRoutingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// --- START: MODIFIED IMPORT ---
	// We expect the ID to be "singleton" or the resource name.
	if req.ID != "singleton" && req.ID != "bgp_routing" {
		resp.Diagnostics.AddError("Unexpected Import Identifier", fmt.Sprintf("Expected 'singleton' or 'bgp_routing', got: %s", req.ID))
		return
	}
	// All singleton imports map to a static "singleton" tfid.
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("tfid"), types.StringValue("singleton"))...)
	// --- END: MODIFIED IMPORT ---
}
