package provider

/*

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/objects"
	"github.com/paloaltonetworks/terraform-provider-scm/internal/models/objects"
)

// RESOURCE for SCM QuarantinedDevice (Package: objects)
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
	tflog.Debug(ctx, "Starting Create function for QuarantinedDevice")
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

	// 3. Initiate the API request with the body.
	createReq := r.client.QuarantinedDevicesAPI.CreateQuarantinedDevices(ctx).QuarantinedDevices(*unpackedScmObject)

	// 4. BLOCK 1: Add the request PARAMETERS to the API call.

	// 5. Execute the API call.
	createdObject, _, err := createReq.Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error creating quarantined_devices", err.Error())
		return
	}

	// 6. Pack the API response back into a Terraform model data.
	packedObject, diags := packQuarantinedDevicesFromSdk(ctx, *createdObject)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() { return }
	resp.Diagnostics.Append(packedObject.As(ctx, &data, basetypes.ObjectAsOptions{})...)
	if resp.Diagnostics.HasError() { return }

	// 7. BLOCK 2: Restore the PARAMETER values from the original plan.
    //    This is necessary for parameters that are sent to the API but not returned in the response.



	// Set the Terraform ID and save the final state.
	var idBuilder strings.Builder


	idBuilder.WriteString(":")

	idBuilder.WriteString(":")

	idBuilder.WriteString(":")
	idBuilder.WriteString(data.Id.ValueString())
	data.Tfid = types.StringValue(idBuilder.String())

	tflog.Debug(ctx, "Created quarantined_devices", map[string]interface{}{"tfid": data.Tfid.ValueString()})
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *QuarantinedDeviceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Step 1 - Initialize a data and savestate of type models.QuarantinedDevices - which is the TF schema struct
	tflog.Debug(ctx, "Starting Read function for QuarantinedDevice")
	var data, savestate models.QuarantinedDevices

	// Step 2 - Fetch the state into savestate
	resp.Diagnostics.Append(req.State.Get(ctx, &savestate)...)
	if resp.Diagnostics.HasError() { return }

	tokens := strings.Split(savestate.Tfid.ValueString(), ":")
	if len(tokens) != 4 {
		resp.Diagnostics.AddError("Error parsing TFID", fmt.Sprintf("Expected a TFID with 4 parts separated by ':', but got %d parts for TFID %s", len(tokens), savestate.Tfid.ValueString()))
		return
	}
	objectId := tokens[3]

	// Step 3 - Make read api call with id = id from state tfid
	tflog.Debug(ctx, "Reading quarantined_devices from SCM API", map[string]interface{}{"id": objectId})
	getReq := r.client.QuarantinedDevicesAPI.(ctx, objectId)
	scmObject, _, err := getReq.Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error reading quarantined_devices", err.Error())
		return
	}



	// Step 5 - Pack the scm object into a terraform model and put it in data we initialized in step 1
	packedObject, diags := packQuarantinedDevicesFromSdk(ctx, *scmObject)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() { return }
	resp.Diagnostics.Append(packedObject.As(ctx, &data, basetypes.ObjectAsOptions{})...)
	if resp.Diagnostics.HasError() { return }



	// Step 7 - Carry over tfid from state back into data
	data.Tfid = savestate.Tfid

	// Step 8 - Set things in params back into data object from the savestate - things like position of security rule

	// Step 9 - Set folder, snippet, device from params back into data if present

	// --- FOLDER RESTORATION (tokens[0]) ---


	// --- SNIPPET RESTORATION (tokens[1]) ---


	// --- DEVICE RESTORATION (tokens[2]) ---


	// Step 10 - Set data back into tf state and done
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *QuarantinedDeviceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	// Step 1: Initialize a plan and state of type models.QuarantinedDevices which is the terraform schema struct
	tflog.Debug(ctx, "Starting Update function for QuarantinedDevice")
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

	// Step 5: Update calls cannot have id sent in payload, so remove it
	// ID is a string, so we set it to its zero value ("") to omit it from the update payload.
	unpackedScmObject.Id = ""

	// Step 6: Get id from token and make update call
	tokens := strings.Split(state.Tfid.ValueString(), ":")
	if len(tokens) != 4 {
        resp.Diagnostics.AddError("Error parsing TFID", fmt.Sprintf("Expected a TFID with 4 parts separated by ':', but got %d parts for TFID %s", len(tokens), state.Tfid.ValueString()))
		return
    }
	objectId := tokens[3]

	tflog.Debug(ctx, "Updating quarantined_devices on SCM API", map[string]interface{}{"id": objectId})
	updateReq := r.client.QuarantinedDevicesAPI.(ctx, objectId).QuarantinedDevices(*unpackedScmObject)

	// Step 7: Retain update parameters so we dont lose them
	// ======================== START: ADD THIS BLOCK ========================
	// Apply any operation parameters from the plan.
	// ========================= END: ADD THIS BLOCK =========================

	// Step 8: Make the update call and get an SCM updatedObject
	updatedObject, _, err := updateReq.Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error updating quarantined_devices", err.Error())
		return
	}

	// Step 9: Pack the SCM updatedObject into a TF object
	packedObject, diags := packQuarantinedDevicesFromSdk(ctx, *updatedObject)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() { return }
	resp.Diagnostics.Append(packedObject.As(ctx, &plan, basetypes.ObjectAsOptions{})...)
	if resp.Diagnostics.HasError() { return }

	// // Preserve any write-only parameter values from the plan.
	//



	// Step 10: Carry over tfid from state into plan
	plan.Tfid = state.Tfid

    // Step 11: Copy write-only attributes from the prior state to the plan for things like position in security rule

	tflog.Debug(ctx, "Updated quarantined_devices", map[string]interface{}{"tfid": plan.Tfid.ValueString()})
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *QuarantinedDeviceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data models.QuarantinedDevices
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() { return }

	tokens := strings.Split(data.Tfid.ValueString(), ":")
	if len(tokens) != 4 {
        resp.Diagnostics.AddError("Error parsing TFID", fmt.Sprintf("Expected a TFID with 4 parts separated by ':', but got %d parts for TFID %s", len(tokens), data.Tfid.ValueString()))
		return
	}
	objectId := tokens[3]

	tflog.Debug(ctx, "Deleting quarantined_devices", map[string]interface{}{"id": objectId})
	deleteReq := r.client.QuarantinedDevicesAPI.DeleteQuarantinedDevices(ctx, objectId)
	_, err := deleteReq.Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error deleting quarantined_devices", err.Error())
	}
}

func (r *QuarantinedDeviceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("tfid"), req, resp)
}















*/
