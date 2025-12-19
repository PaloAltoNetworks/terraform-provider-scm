package provider

/*

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/objects"
	"github.com/paloaltonetworks/terraform-provider-scm/internal/models/objects"
	"github.com/paloaltonetworks/terraform-provider-scm/internal/utils"
)

// SINGLETON RESOURCE for SCM QuarantinedDevice (Package: objects)
var (
	_ resource.Resource                = &QuarantinedDeviceResource{}
	_ resource.ResourceWithConfigure   = &QuarantinedDeviceResource{}
	_ resource.ResourceWithImportState = &QuarantinedDeviceResource{}
)

func NewQuarantinedDeviceResource() resource.Resource {
	return &QuarantinedDeviceResource{}
}

// QuarantinedDeviceResource defines the resource implementation.
type QuarantinedDeviceResource struct {
	client *objects.APIClient
}

func (r *QuarantinedDeviceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_quarantined_device"
}

func (r *QuarantinedDeviceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = models.QuarantinedDevicesResourceSchema

	resp.Schema.MarkdownDescription = "**Singleton Resource.** " + resp.Schema.MarkdownDescription +
		"\n\nThis resource is a singleton, meaning only one instance can exist. " +
		"If the resource typically exists (e.g. bgp_routing), you should import it before managing it."
}

func (r *QuarantinedDeviceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	clients, ok := req.ProviderData.(map[string]interface{})
	if !ok {
		resp.Diagnostics.AddError("Unexpected Resource Configure Type", fmt.Sprintf("Expected map[string]interface{}, got: %T.", req.ProviderData))
		return
	}
	client, ok := clients["objects"].(*objects.APIClient)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Client Type", fmt.Sprintf("Expected *objects.APIClient for 'objects' client."))
		return
	}
	r.client = client
}

func (r *QuarantinedDeviceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Starting Create function for Singleton QuarantinedDevice")
	var data models.QuarantinedDevices

	// 1. Get the plan from Terraform into the data model.
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() { return }



	// Unpack the plan to an SCM SDK object.
	planObject, diags := types.ObjectValueFrom(ctx, models.QuarantinedDevices{}.AttrTypes(), &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() { return }

	// 2. Unpack the request BODY from data into an SDK object.
	unpackedScmObject, diags := unpackQuarantinedDevicesToSdk(ctx, planObject)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() { return }

	tflog.Debug(ctx, "Creating quarantined_devices on SCM API")

	// --- START: DYNAMIC CREATE LOGIC ---
	var scmObjectInterface interface{}
	var err error


		// This singleton has a POST for Create (e.g., ssl-decryption-settings)
		tflog.Debug(ctx, "Using POST operation: CreateQuarantinedDevices")


		createReq := r.client.QuarantinedDevicesAPI.CreateQuarantinedDevices(ctx).QuarantinedDevices(*unpackedScmObject)

		scmObjectInterface, _, err = createReq.Execute()


	// --- END: DYNAMIC CREATE LOGIC ---

	if err != nil {
		resp.Diagnostics.AddError("Error creating quarantined_devices", err.Error())
		detailedMessage := utils.PrintScmError(err)

		resp.Diagnostics.AddError(
			"SCM Resource Creation Failed: API Request Failed",
			detailedMessage,
		)
		return
	}

	// --- START MODIFICATION: DYNAMIC RESPONSE HANDLING (Applied to Create) ---
	var createdObject *objects.QuarantinedDevices

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
				var targetStruct objects.QuarantinedDevices
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
            var targetStruct objects.QuarantinedDevices
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

	packedObject, diags := packQuarantinedDevicesFromSdk(ctx, *createdObject)
	// --- END MODIFICATION ---

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() { return }
	resp.Diagnostics.Append(packedObject.As(ctx, &data, basetypes.ObjectAsOptions{})...)
	if resp.Diagnostics.HasError() { return }



	data.Tfid = types.StringValue("singleton_quarantined_devices")

	tflog.Debug(ctx, "Created quarantined_devices", map[string]interface{}{"tfid": data.Tfid.ValueString()})
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *QuarantinedDeviceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Step 1 - Initialize a data and savestate of type models.QuarantinedDevices - which is the TF schema struct
	tflog.Debug(ctx, "Starting Read function for Singleton QuarantinedDevice")
	var data, savestate models.QuarantinedDevices

	// Step 2 - Fetch the state into savestate
	resp.Diagnostics.Append(req.State.Get(ctx, &savestate)...)
	if resp.Diagnostics.HasError() { return }

	tflog.Debug(ctx, "Reading quarantined_devices from SCM API")
	getReq := r.client.QuarantinedDevicesAPI.ListQuarantinedDevices(ctx)

	// --- START: DYNAMIC PARAMETERS FOR READ ---
	// Use reflection to safely check for and use query parameters if they exist in the model
	v := reflect.ValueOf(savestate)
	f := v.FieldByName("HostId")
	if f.IsValid() && !f.IsZero() && !f.Interface().(types.String).IsNull() {
		getReq = getReq.HostId(f.Interface().(types.String).ValueString())
	}
	// Use reflection to safely check for and use query parameters if they exist in the model
	v := reflect.ValueOf(savestate)
	f := v.FieldByName("SerialNumber")
	if f.IsValid() && !f.IsZero() && !f.Interface().(types.String).IsNull() {
		getReq = getReq.SerialNumber(f.Interface().(types.String).ValueString())
	}
	// --- END: DYNAMIC PARAMETERS FOR READ ---

	// Use interface{} to handle flexible return types (List vs Object)
	var scmObjectInterface interface{}
	var httpErr *http.Response
	var err error
	scmObjectInterface, httpErr, err = getReq.Execute()

	if err != nil {
		if httpErr != nil && httpErr.StatusCode == http.StatusNotFound {
			// FIX: Remove undefined variable objectId from log
			tflog.Debug(ctx, "Got no quarantined_devices on read SCM API. Remove from state to let terraform create", map[string]interface{}{"tfid": savestate.Tfid.ValueString()})
			resp.State.RemoveResource(ctx)
		} else {
			// FIX: Remove undefined variable objectId from log
			tflog.Debug(ctx, "Got an exception on read SCM API.", map[string]interface{}{"tfid": savestate.Tfid.ValueString()})
			resp.Diagnostics.AddError("Error reading quarantined_devices", err.Error())
			detailedMessage := utils.PrintScmError(err)
			resp.Diagnostics.AddError("SCM Resource Read Failed: API Request Failed", detailedMessage)
		}
		return
	}

	// --- START MODIFICATION: DYNAMIC RESPONSE HANDLING (Applied to Read) ---
	var scmObject *objects.QuarantinedDevices

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
				var targetStruct objects.QuarantinedDevices
				if err := json.Unmarshal(jsonBytes, &targetStruct); err == nil {
					scmObject = &targetStruct
				}
			} else {
                // List is empty, treat as Not Found
                tflog.Debug(ctx, "Got no quarantined_devices on read SCM API (empty list). Remove from state", map[string]interface{}{"tfid": savestate.Tfid.ValueString()})
				resp.State.RemoveResource(ctx)
				return
			}
		} else {
            // 2. Not a list wrapper, assume it's the direct object (e.g., bgp-routing)
            // Use the same JSON trick to be safe against pointer mismatches
            jsonBytes, _ := json.Marshal(scmObjectInterface)
            var targetStruct objects.QuarantinedDevices
            if err := json.Unmarshal(jsonBytes, &targetStruct); err == nil {
                scmObject = &targetStruct
            }
        }
	}

	// If after all checks, scmObject is still nil, then the API returned 200 OK but no object was found.
	if scmObject == nil {
		tflog.Debug(ctx, "Got no quarantined_devices on read SCM API (nil object). Remove from state to let terraform create", map[string]interface{}{"tfid": savestate.Tfid.ValueString()})
		resp.State.RemoveResource(ctx)
		return
	}

	// Step 5 - Pack the scm object into a terraform model and put it in data we initialized in step 1
	packedObject, diags := packQuarantinedDevicesFromSdk(ctx, *scmObject)
	// --- END MODIFICATION ---



	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() { return }
	resp.Diagnostics.Append(packedObject.As(ctx, &data, basetypes.ObjectAsOptions{})...)
	if resp.Diagnostics.HasError() { return }



	// Step 7 - Carry over tfid from state back into data
	data.Tfid = savestate.Tfid

	// Step 8 - Set things in params back into data object from the savestate - things like position of security rule

	// --- START: RESTORE READ PARAMETERS ---
	// Use reflection to safely copy query parameter back to data model
	vData := reflect.ValueOf(&data).Elem()
	vSave := reflect.ValueOf(savestate)
	fData := vData.FieldByName("HostId")
	fSave := vSave.FieldByName("HostId")
	if fData.IsValid() && fData.CanSet() && fSave.IsValid() {
		fData.Set(fSave)
	}
	// Use reflection to safely copy query parameter back to data model
	vData := reflect.ValueOf(&data).Elem()
	vSave := reflect.ValueOf(savestate)
	fData := vData.FieldByName("SerialNumber")
	fSave := vSave.FieldByName("SerialNumber")
	if fData.IsValid() && fData.CanSet() && fSave.IsValid() {
		fData.Set(fSave)
	}
	// --- END: RESTORE READ PARAMETERS ---

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *QuarantinedDeviceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	// Step 1: Initialize a plan and state of type models.QuarantinedDevices which is the terraform schema struct
	tflog.Debug(ctx, "Starting Update function for Singleton QuarantinedDevice")
	var plan, state models.QuarantinedDevices

	// Step 2: Get the plan from plan file (resource.tf) into plan and state from tfstate into state
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() { return }
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() { return }



	// Step 3: Creates a plan object from the plan
	planObject, diags := types.ObjectValueFrom(ctx, models.QuarantinedDevices{}.AttrTypes(), &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() { return }

	// Step 4: Unpack the plan object to an SCM Object
	unpackedScmObject, diags := unpackQuarantinedDevicesToSdk(ctx, planObject)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() { return }

	// --- START: MODIFIED API CALL (no path param) ---
	var scmObjectInterface interface{}
	var httpErr *http.Response
	var err error


		resp.Diagnostics.AddError("Update Not Supported", "This singleton resource does not support Update (PUT) operations.")
		return

	// --- END: MODIFIED API CALL ---

	if err != nil {
		if httpErr != nil && httpErr.StatusCode == http.StatusNotFound {
			// FIX: Remove undefined variable objectId from log
			tflog.Debug(ctx, "Got no quarantined_devices on update SCM API. Remove from state to let terraform create", map[string]interface{}{"tfid": state.Tfid.ValueString()})
			resp.State.RemoveResource(ctx)
		} else {
			// FIX: Remove undefined variable objectId from log
			tflog.Debug(ctx, "Got an exception on update SCM API.", map[string]interface{}{"tfid": state.Tfid.ValueString()})
			resp.Diagnostics.AddError("Error updating quarantined_devices", err.Error())
			detailedMessage := utils.PrintScmError(err)
			resp.Diagnostics.AddError("SCM Resource Update Failed: API Request Failed", detailedMessage)
		}
		return
	}

	// --- START MODIFICATION: DYNAMIC RESPONSE HANDLING (Applied to Update) ---
	var updatedObject *objects.QuarantinedDevices

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
				var targetStruct objects.QuarantinedDevices
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
            var targetStruct objects.QuarantinedDevices
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

	packedObject, diags := packQuarantinedDevicesFromSdk(ctx, *updatedObject)
	// --- END MODIFICATION ---

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() { return }
	resp.Diagnostics.Append(packedObject.As(ctx, &plan, basetypes.ObjectAsOptions{})...)
	if resp.Diagnostics.HasError() { return }

	// // Preserve any write-only parameter values from the plan.
	//



	// Step 10: Carry over tfid from state into plan
	plan.Tfid = state.Tfid

    // Step 11: Copy write-only attributes from the prior state to the plan for things like position in security rule

	// --- START: RESTORE READ PARAMETERS ---
	// Use reflection to safely copy query parameter back to data model
	vPlan := reflect.ValueOf(&plan).Elem()
	vState := reflect.ValueOf(state)
	fPlan := vPlan.FieldByName("HostId")
	fState := vState.FieldByName("HostId")
	if fPlan.IsValid() && fPlan.CanSet() && fState.IsValid() {
		fPlan.Set(fState)
	}
	// Use reflection to safely copy query parameter back to data model
	vPlan := reflect.ValueOf(&plan).Elem()
	vState := reflect.ValueOf(state)
	fPlan := vPlan.FieldByName("SerialNumber")
	fState := vState.FieldByName("SerialNumber")
	if fPlan.IsValid() && fPlan.CanSet() && fState.IsValid() {
		fPlan.Set(fState)
	}
	// --- END: RESTORE READ PARAMETERS ---

	tflog.Debug(ctx, "Updated quarantined_devices", map[string]interface{}{"tfid": plan.Tfid.ValueString()})
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *QuarantinedDeviceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// --- START: DYNAMIC DELETE LOGIC ---

		tflog.Debug(ctx, "Deleting singleton quarantined_devices from SCM API")
		deleteReq := r.client.QuarantinedDevicesAPI.DeleteQuarantinedDevices(ctx)
		_, _, err := deleteReq.Execute()
		if err != nil {
			resp.Diagnostics.AddError("Error deleting quarantined_devices", err.Error())
			detailedMessage := utils.PrintScmError(err)
			resp.Diagnostics.AddError("SCM Resource Deletion Failed: API Request Failed", detailedMessage)
		}

	// --- END: DYNAMIC DELETE LOGIC ---
}

func (r *QuarantinedDeviceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// --- START: MODIFIED IMPORT ---
	// We expect the ID to be "singleton" or the resource name.
	if req.ID != "singleton" && req.ID != "quarantined_devices" {
		resp.Diagnostics.AddError("Unexpected Import Identifier", fmt.Sprintf("Expected 'singleton' or 'quarantined_devices', got: %s", req.ID))
		return
	}
	// All singleton imports map to a static "singleton" tfid.
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("tfid"), types.StringValue("singleton_quarantined_devices"))...)
	// --- END: MODIFIED IMPORT ---
}

















*/
