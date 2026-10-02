package provider

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/ztna_connector_all"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/ztna_connector_all"
	"github.com/paloaltonetworks/terraform-provider-scm/internal/utils"
)

// RESOURCE for SCM Wildcard (Package: ztna_connector_all)
var (
	_ resource.Resource                = &WildcardResource{}
	_ resource.ResourceWithConfigure   = &WildcardResource{}
	_ resource.ResourceWithImportState = &WildcardResource{}
)

func NewWildcardResource() resource.Resource {
	return &WildcardResource{}
}

// WildcardResource defines the resource implementation.
type WildcardResource struct {
	client *ztna_connector_all.APIClient
}

func (r *WildcardResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "ztna_wildcard"
}

func (r *WildcardResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = models.WildcardsResourceSchema
}

func (r *WildcardResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	clients, ok := req.ProviderData.(map[string]interface{})
	if !ok {
		resp.Diagnostics.AddError("Unexpected Resource Configure Type", fmt.Sprintf("Expected map[string]interface{}, got: %T.", req.ProviderData))
		return
	}
	client, ok := clients["ztna_connector_all"].(*ztna_connector_all.APIClient)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Client Type", fmt.Sprintf("Expected *ztna_connector_all.APIClient for 'ztna_connector_all' client."))
		return
	}
	r.client = client
}

func (r *WildcardResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Starting Create function for Wildcard")
	var data models.Wildcards

	// 1. Get the plan from Terraform into the data model.
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Unpack the plan to an SCM SDK object.
	planObject, diags := types.ObjectValueFrom(ctx, models.Wildcards{}.AttrTypes(), &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// 2. Unpack the request BODY from data into an SDK object.
	unpackedScmObject, diags := unpackWildcardsToSdk(ctx, planObject)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating wildcards on SCM API")

	// 3. Initiate the API request with the body.
	createReq := r.client.WildcardAPI.CreateWildcard(ctx).Wildcards(*unpackedScmObject)

	// 4. BLOCK 1: Add the request PARAMETERS to the API call.

	// 5. Execute the API call.
	_, err := createReq.Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error creating wildcards", err.Error())
		detailedMessage := utils.PrintScmError(err)

		resp.Diagnostics.AddError(
			"SCM Resource Creation Failed: API Request Failed",
			detailedMessage,
		)
		return
	}

	// API returns no body on create — fetch the object by name to get the assigned ID.
	createdObject, err := r.client.WildcardAPI.FetchWildcard(ctx, data.Name.ValueString(), nil, nil, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error fetching created wildcards after creation", err.Error())
		return
	}
	if createdObject == nil {
		resp.Diagnostics.AddError("Error fetching created wildcards", fmt.Sprintf("Could not find wildcards with name %s after creation", data.Name.ValueString()))
		return
	}
	// 6. Pack the API response back into a Terraform model data.
	packedObject, diags := packWildcardsFromSdk(ctx, *createdObject)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// 6a. Normalize null lists: when the API returns null for a list field that
	// the plan had as [] (empty list), coerce to empty list to satisfy Terraform's
	// consistency requirement. This recurses into nested objects and list elements.
	packedObject, diags = utils.NormalizeNullLists(ctx, planObject, packedObject)
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
	// NOTE: Skip the path parameter (e.g. "id", "oid") — its value comes from the API, not the plan.

	// Set the Terraform ID and save the final state.
	var idBuilder strings.Builder

	idBuilder.WriteString(":")

	idBuilder.WriteString(":")

	idBuilder.WriteString(":")
	idBuilder.WriteString(data.Oid.ValueString())
	data.Tfid = types.StringValue(idBuilder.String())

	tflog.Debug(ctx, "Created wildcards", map[string]interface{}{"tfid": data.Tfid.ValueString()})
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WildcardResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Step 1 - Initialize a data and savestate of type models.Wildcards - which is the TF schema struct
	tflog.Debug(ctx, "Starting Read function for Wildcard")
	var data, savestate models.Wildcards

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
	tflog.Debug(ctx, "Reading wildcards from SCM API", map[string]interface{}{"id": objectId})
	getReq := r.client.WildcardAPI.GetWildcardByID(ctx, objectId)
	scmObject, httpErr, err := getReq.Execute()
	if err != nil {
		if httpErr != nil && httpErr.StatusCode == http.StatusNotFound {
			tflog.Debug(ctx, "Got no wildcards on read SCM API. Remove from state to let terraform create", map[string]interface{}{"id": objectId})
			resp.State.RemoveResource(ctx)
		} else {
			tflog.Debug(ctx, "Got an exception on read SCM API. ", map[string]interface{}{"id": objectId})
			resp.Diagnostics.AddError("Error reading wildcards", err.Error())
			detailedMessage := utils.PrintScmError(err)

			resp.Diagnostics.AddError(
				"SCM Resource Read Failed: API Request Failed",
				detailedMessage,
			)
		}
		return
	}

	// Step 5 - Pack the scm object into a terraform model and put it in data we initialized in step 1
	packedObject, diags := packWildcardsFromSdk(ctx, *scmObject)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Step 5a - Normalize null lists against prior state. Read has no plan, so we
	// compare the API response against savestate. When the API returns [] for a
	// list that was null in prior state, coerce it back to null to avoid a
	// perpetual diff on refresh. Real changes (populated lists) are preserved.
	savestateObject, diags := types.ObjectValueFrom(ctx, models.Wildcards{}.AttrTypes(), &savestate)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	packedObject, diags = utils.NormalizeNullLists(ctx, savestateObject, packedObject)
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

	// --- SNIPPET RESTORATION (tokens[1]) ---

	// --- DEVICE RESTORATION (tokens[2]) ---

	// Step 10 - Set data back into tf state and done
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WildcardResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	// Step 1: Initialize a plan and state of type models.Wildcards which is the terraform schema struct
	tflog.Debug(ctx, "Starting Update function for Wildcard")
	var plan, state models.Wildcards

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
	planObject, diags := types.ObjectValueFrom(ctx, models.Wildcards{}.AttrTypes(), &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Step 4: Unpack the plan object to an SCM Object
	unpackedScmObject, diags := unpackWildcardsToSdk(ctx, planObject)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Step 5: Update calls cannot have id sent in payload, so remove it
	// ID is a pointer, so we nil it out to omit it from the update payload.
	unpackedScmObject.Oid = nil

	// Step 6: Get id from token and make update call
	tokens := strings.Split(state.Tfid.ValueString(), ":")
	if len(tokens) != 4 {
		resp.Diagnostics.AddError("Error parsing TFID", fmt.Sprintf("Expected a TFID with 4 parts separated by ':', but got %d parts for TFID %s", len(tokens), state.Tfid.ValueString()))
		return
	}
	objectId := tokens[3]

	tflog.Debug(ctx, "Updating wildcards on SCM API", map[string]interface{}{"id": objectId})
	updateReq := r.client.WildcardAPI.UpdateWildcardByID(ctx, objectId).Wildcards(*unpackedScmObject)

	// Step 7: Retain update parameters so we dont lose them
	// ======================== START: ADD THIS BLOCK ========================
	// Apply any operation parameters from the plan.
	// ========================= END: ADD THIS BLOCK =========================

	// Step 8: Make the update call and get an SCM updatedObject
	httpErr, err := updateReq.Execute()
	if err != nil {
		if httpErr != nil && httpErr.StatusCode == http.StatusNotFound {
			tflog.Debug(ctx, "Got no wildcards on update SCM API. Remove from state to let terraform create", map[string]interface{}{"id": objectId})
			resp.State.RemoveResource(ctx)
		} else {
			tflog.Debug(ctx, "Got an exception on update SCM API. ", map[string]interface{}{"id": objectId})
			resp.Diagnostics.AddError("Error updating wildcards", err.Error())
			detailedMessage := utils.PrintScmError(err)

			resp.Diagnostics.AddError(
				"SCM Resource Update Failed: API Request Failed",
				detailedMessage,
			)
		}
		return
	}

	// API returns no body on update — re-fetch by ID.
	updatedObject, _, err := r.client.WildcardAPI.GetWildcardByID(ctx, objectId).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error fetching updated wildcards after update", err.Error())
		return
	}

	// Step 9: Pack the SCM updatedObject into a TF object
	packedObject, diags := packWildcardsFromSdk(ctx, *updatedObject)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Step 9a: Normalize null lists: when the API returns null for a list field that
	// the plan had as [] (empty list), coerce to empty list to satisfy Terraform's
	// consistency requirement. This recurses into nested objects and list elements.
	packedObject, diags = utils.NormalizeNullLists(ctx, planObject, packedObject)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(packedObject.As(ctx, &plan, basetypes.ObjectAsOptions{})...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve any operation parameter values from the plan (folder, snippet, device).
	// This ensures the user's configured value is preserved regardless of what the API returns.
	// NOTE: Skip the path parameter (e.g. "id", "oid") — its value comes from the API re-fetch, not the plan.

	// Step 10: Carry over tfid from state into plan
	plan.Tfid = state.Tfid

	// Step 11: Copy write-only attributes from the prior state to the plan for things like position in security rule

	tflog.Debug(ctx, "Updated wildcards", map[string]interface{}{"tfid": plan.Tfid.ValueString()})
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *WildcardResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data models.Wildcards
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

	tflog.Debug(ctx, "Deleting wildcards", map[string]interface{}{"id": objectId})
	deleteReq := r.client.WildcardAPI.DeleteWildcardByID(ctx, objectId)
	httpResp, err := deleteReq.Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error deleting wildcards", err.Error())
		detailedMessage := utils.PrintScmError(err)
		resp.Diagnostics.AddError(
			"SCM Resource Deleteion Failed: API Request Failed",
			detailedMessage,
		)
		return
	}
	// For 202 Accepted responses the delete is asynchronous. Poll the GET endpoint
	// until the resource is gone (404) or a timeout is reached, so that dependent
	// resources (e.g. a connector group) are not destroyed before this one is fully
	// removed on the backend.
	if httpResp != nil && httpResp.StatusCode == http.StatusAccepted {
		deadline := time.Now().Add(2 * time.Minute)
		for time.Now().Before(deadline) {
			time.Sleep(3 * time.Second)
			_, getResp, getErr := r.client.WildcardAPI.GetWildcardByID(ctx, objectId).Execute()
			if getErr != nil {
				// If the SDK returns an error check whether it is a 404 — that means deletion is complete.
				if getResp != nil && getResp.StatusCode == http.StatusNotFound {
					break
				}
				// Any other error: stop polling and surface it.
				resp.Diagnostics.AddWarning("Delete poll error", getErr.Error())
				break
			}
			if getResp != nil && getResp.StatusCode == http.StatusNotFound {
				break
			}
		}
	}
}

func (r *WildcardResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	// If validation passes, store the import ID in tfid
	resource.ImportStatePassthroughID(ctx, path.Root("tfid"), req, resp)
}
