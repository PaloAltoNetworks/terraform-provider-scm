package provider

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/network_services"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/network_services"
)

// RESOURCE for SCM LoopbackInterface (Package: network_services)
var (
	_ resource.Resource                = &LoopbackInterfaceResource{}
	_ resource.ResourceWithConfigure   = &LoopbackInterfaceResource{}
	_ resource.ResourceWithImportState = &LoopbackInterfaceResource{}
)

func NewLoopbackInterfaceResource() resource.Resource {
	return &LoopbackInterfaceResource{}
}

// LoopbackInterfaceResource defines the resource implementation.
type LoopbackInterfaceResource struct {
	client *network_services.APIClient
}

func (r *LoopbackInterfaceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_loopback_interface"
}

func (r *LoopbackInterfaceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = models.LoopbackInterfacesResourceSchema
}

func (r *LoopbackInterfaceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *LoopbackInterfaceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Starting Create function for LoopbackInterface")
	var data models.LoopbackInterfaces

	// 1. Get the plan from Terraform into the data model.
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Unpack the plan to an SCM SDK object.
	planObject, diags := types.ObjectValueFrom(ctx, models.LoopbackInterfaces{}.AttrTypes(), &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// 2. Unpack the request BODY from data into an SDK object.
	unpackedScmObject, diags := unpackLoopbackInterfacesToSdk(ctx, planObject)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating loopback_interfaces on SCM API")

	// 3. Initiate the API request with the body.
	createReq := r.client.LoopbackInterfacesAPI.CreateLoopbackInterfaces(ctx).LoopbackInterfaces(*unpackedScmObject)

	// 4. BLOCK 1: Add the request PARAMETERS to the API call.

	// 5. Execute the API call.
	createdObject, _, err := createReq.Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error creating loopback_interfaces", err.Error())
		return
	}

	// 6. Pack the API response back into a Terraform model data.
	packedObject, diags := packLoopbackInterfacesFromSdk(ctx, *createdObject)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(packedObject.As(ctx, &data, basetypes.ObjectAsOptions{})...)
	if resp.Diagnostics.HasError() {
		return
	}

	// 7. BLOCK 2: Restore the PARAMETER values from the original plan.
	//    This is necessary for parameters that are sent to the API but not returned in the response.

	// Set the Terraform ID and save the final state.
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

	tflog.Debug(ctx, "Created loopback_interfaces", map[string]interface{}{"tfid": data.Tfid.ValueString()})
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LoopbackInterfaceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Step 1 - Initialize a data and savestate of type models.LoopbackInterfaces - which is the TF schema struct
	tflog.Debug(ctx, "Starting Read function for LoopbackInterface")
	var data, savestate models.LoopbackInterfaces

	// Step 2 - Fetch the state into savestate
	resp.Diagnostics.Append(req.State.Get(ctx, &savestate)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tokens := strings.Split(savestate.Tfid.ValueString(), ":")
	if len(tokens) != 4 {
		resp.Diagnostics.AddError("Error parsing TFID", fmt.Sprintf("Expected a TFID with 4 parts separated by ':', but got %d parts for TFID %s", len(tokens), savestate.Tfid.ValueString()))
		return
	}
	objectId := tokens[3]

	// Step 3 - Make read api call with id = id from state tfid
	tflog.Debug(ctx, "Reading loopback_interfaces from SCM API", map[string]interface{}{"id": objectId})
	getReq := r.client.LoopbackInterfacesAPI.GetLoopbackInterfacesByID(ctx, objectId)
	scmObject, _, err := getReq.Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error reading loopback_interfaces", err.Error())
		return
	}

	// Step 5 - Pack the scm object into a terraform model and put it in data we initialized in step 1
	packedObject, diags := packLoopbackInterfacesFromSdk(ctx, *scmObject)
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

	// Step 9 - Set folder, snippet, device from params back into data if present

	// --- FOLDER RESTORATION (tokens[0]) ---

	// Use reflection to safely restore the Folder field from the TFID token 0.
	vFolder := reflect.ValueOf(&data).Elem() // Unique variable: vFolder
	fFolder := vFolder.FieldByName("Folder") // Unique variable: fFolder

	if fFolder.IsValid() && fFolder.CanSet() {
		tokenValue := tokens[0]

		if tokenValue != "" {
			newStringValue := basetypes.NewStringValue(tokenValue)
			fFolder.Set(reflect.ValueOf(newStringValue))
		} else {
			newNullValue := basetypes.NewStringNull()
			fFolder.Set(reflect.ValueOf(newNullValue))
		}
	}

	// --- SNIPPET RESTORATION (tokens[1]) ---

	// Use reflection to safely restore the Snippet field from the TFID token 1.
	vSnippet := reflect.ValueOf(&data).Elem()   // Unique variable: vSnippet
	fSnippet := vSnippet.FieldByName("Snippet") // Unique variable: fSnippet

	if fSnippet.IsValid() && fSnippet.CanSet() {
		tokenValue := tokens[1]

		if tokenValue != "" {
			newStringValue := basetypes.NewStringValue(tokenValue)
			fSnippet.Set(reflect.ValueOf(newStringValue))
		} else {
			newNullValue := basetypes.NewStringNull()
			fSnippet.Set(reflect.ValueOf(newNullValue))
		}
	}

	// --- DEVICE RESTORATION (tokens[2]) ---

	// Use reflection to safely restore the Device field from the TFID token 2.
	vDevice := reflect.ValueOf(&data).Elem() // Unique variable: vDevice
	fDevice := vDevice.FieldByName("Device") // Unique variable: fDevice

	if fDevice.IsValid() && fDevice.CanSet() {
		tokenValue := tokens[2]

		if tokenValue != "" {
			newStringValue := basetypes.NewStringValue(tokenValue)
			fDevice.Set(reflect.ValueOf(newStringValue))
		} else {
			newNullValue := basetypes.NewStringNull()
			fDevice.Set(reflect.ValueOf(newNullValue))
		}
	}

	// Step 10 - Set data back into tf state and done
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LoopbackInterfaceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	// Step 1: Initialize a plan and state of type models.LoopbackInterfaces which is the terraform schema struct
	tflog.Debug(ctx, "Starting Update function for LoopbackInterface")
	var plan, state models.LoopbackInterfaces

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
	planObject, diags := types.ObjectValueFrom(ctx, models.LoopbackInterfaces{}.AttrTypes(), &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Step 4: Unpack the plan object to an SCM Object
	unpackedScmObject, diags := unpackLoopbackInterfacesToSdk(ctx, planObject)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Step 5: Update calls cannot have id sent in payload, so remove it
	// ID is a pointer, so we nil it out to omit it from the update payload.
	unpackedScmObject.Id = nil

	// Step 6: Get id from token and make update call
	tokens := strings.Split(state.Tfid.ValueString(), ":")
	if len(tokens) != 4 {
		resp.Diagnostics.AddError("Error parsing TFID", fmt.Sprintf("Expected a TFID with 4 parts separated by ':', but got %d parts for TFID %s", len(tokens), state.Tfid.ValueString()))
		return
	}
	objectId := tokens[3]

	tflog.Debug(ctx, "Updating loopback_interfaces on SCM API", map[string]interface{}{"id": objectId})
	updateReq := r.client.LoopbackInterfacesAPI.UpdateLoopbackInterfacesByID(ctx, objectId).LoopbackInterfaces(*unpackedScmObject)

	// Step 7: Retain update parameters so we dont lose them
	// ======================== START: ADD THIS BLOCK ========================
	// Apply any operation parameters from the plan.
	// ========================= END: ADD THIS BLOCK =========================

	// Step 8: Make the update call and get an SCM updatedObject
	updatedObject, _, err := updateReq.Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error updating loopback_interfaces", err.Error())
		return
	}

	// Step 9: Pack the SCM updatedObject into a TF object
	packedObject, diags := packLoopbackInterfacesFromSdk(ctx, *updatedObject)
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
	// _ = req.Plan.GetAttribute(ctx, path.Root("id"), &plan.Id)
	//

	// Step 10: Carry over tfid from state into plan
	plan.Tfid = state.Tfid

	// Step 11: Copy write-only attributes from the prior state to the plan for things like position in security rule

	tflog.Debug(ctx, "Updated loopback_interfaces", map[string]interface{}{"tfid": plan.Tfid.ValueString()})
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *LoopbackInterfaceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data models.LoopbackInterfaces
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tokens := strings.Split(data.Tfid.ValueString(), ":")
	if len(tokens) != 4 {
		resp.Diagnostics.AddError("Error parsing TFID", fmt.Sprintf("Expected a TFID with 4 parts separated by ':', but got %d parts for TFID %s", len(tokens), data.Tfid.ValueString()))
		return
	}
	objectId := tokens[3]

	tflog.Debug(ctx, "Deleting loopback_interfaces", map[string]interface{}{"id": objectId})
	deleteReq := r.client.LoopbackInterfacesAPI.DeleteLoopbackInterfacesByID(ctx, objectId)
	_, err := deleteReq.Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error deleting loopback_interfaces", err.Error())
	}
}

func (r *LoopbackInterfaceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("tfid"), req, resp)
}
