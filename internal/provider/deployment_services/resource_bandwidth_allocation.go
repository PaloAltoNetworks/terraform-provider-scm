package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/paloaltonetworks/scm-go/generated/deployment_services"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/deployment_services"
)

// RESOURCE: BandwidthAllocation (Specialized: List-Managed / No-ID)
var (
	_ resource.Resource                = &BandwidthAllocationResource{}
	_ resource.ResourceWithConfigure   = &BandwidthAllocationResource{}
	_ resource.ResourceWithImportState = &BandwidthAllocationResource{}
)

func NewBandwidthAllocationResource() resource.Resource {
	return &BandwidthAllocationResource{}
}

type BandwidthAllocationResource struct {
	client *deployment_services.APIClient
}

func (r *BandwidthAllocationResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_bandwidth_allocation"
}

func (r *BandwidthAllocationResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = models.BandwidthAllocationsResourceSchema
}

func (r *BandwidthAllocationResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *BandwidthAllocationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Starting Create function for BandwidthAllocation")
	var data models.BandwidthAllocations

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Unpack the plan to an SCM SDK object
	planObject, diags := types.ObjectValueFrom(ctx, models.BandwidthAllocations{}.AttrTypes(), &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	unpackedScmObject, diags := unpackBandwidthAllocationsToSdk(ctx, planObject)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Call the API (POST)
	createReq := r.client.BandwidthAllocationsAPI.CreateBandwidthAllocations(ctx).BandwidthAllocations(*unpackedScmObject)

	// Apply params

	createdObject, _, err := createReq.Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error creating BandwidthAllocation", err.Error())
		return
	}

	// The 'Name' is the identity for this resource.
	if !data.Name.IsNull() {
		data.Tfid = data.Name
	} else {
		data.Tfid = types.StringValue("unknown")
	}

	// Re-pack the response if available
	if createdObject != nil {
		packedObject, diags := packBandwidthAllocationsFromSdk(ctx, *createdObject)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

		var apiData models.BandwidthAllocations
		resp.Diagnostics.Append(packedObject.As(ctx, &apiData, basetypes.ObjectAsOptions{})...)
		if resp.Diagnostics.HasError() {
			return
		}

		// 🟢 FIX: MASKING + VISIBILITY
		// If the user provided a list, we force the state to match the plan to avoid errors.
		if !data.SpnNameList.IsNull() && !data.SpnNameList.IsUnknown() {
			// Check if the API returned DIFFERENT data than what we sent.
			// If so, warn the user so they know what the real values are.
			if !apiData.SpnNameList.Equal(data.SpnNameList) {
				var realSpns []string
				apiData.SpnNameList.ElementsAs(ctx, &realSpns, false)
				resp.Diagnostics.AddWarning(
					"API Response Modified Input",
					fmt.Sprintf(
						"The API modified the 'spn_name_list' (e.g., added suffixes or items).\n"+
							"To avoid consistency errors, we have saved your original input to the state.\n"+
							"Real values returned by API: %v\n"+
							"Please update your configuration to match these values to avoid future drift.",
						realSpns,
					),
				)
			}

			// Overwrite with user input to satisfy Terraform consistency check
			apiData.SpnNameList = data.SpnNameList
		}

		data = apiData
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BandwidthAllocationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	tflog.Debug(ctx, "Starting Read function for BandwidthAllocation")
	var state models.BandwidthAllocations
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Identify the target name
	targetName := state.Name.ValueString()
	if state.Tfid.ValueString() != "" && state.Tfid.ValueString() != "unknown" {
		targetName = state.Tfid.ValueString()
	}

	// 1. Call the List operation to get all items
	listReq := r.client.BandwidthAllocationsAPI.ListBandwidthAllocations(ctx)
	listResponse, _, err := listReq.Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error reading BandwidthAllocation", err.Error())
		return
	}

	// 2. Iterate through the list to find the matching item
	var foundItem *deployment_services.BandwidthAllocations

	if listResponse != nil && listResponse.Data != nil {
		for _, item := range listResponse.Data {
			if item.Name != "" && item.Name == targetName {
				foundItem = &item
				break
			}
		}
	}

	if foundItem == nil {
		tflog.Warn(ctx, "BandwidthAllocation not found in list, removing from state", map[string]interface{}{"name": targetName})
		resp.State.RemoveResource(ctx)
		return
	}

	// 3. Pack the found item
	packedObject, diags := packBandwidthAllocationsFromSdk(ctx, *foundItem)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var apiData models.BandwidthAllocations
	resp.Diagnostics.Append(packedObject.As(ctx, &apiData, basetypes.ObjectAsOptions{})...)
	if resp.Diagnostics.HasError() {
		return
	}

	// 🟢 NO MASKING: Let the Read operation reveal the true state (drift detection).

	apiData.Tfid = state.Name
	resp.Diagnostics.Append(resp.State.Set(ctx, &apiData)...)
}

func (r *BandwidthAllocationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Starting Update function for BandwidthAllocation")
	var plan models.BandwidthAllocations
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	planObject, diags := types.ObjectValueFrom(ctx, models.BandwidthAllocations{}.AttrTypes(), &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	unpackedScmObject, diags := unpackBandwidthAllocationsToSdk(ctx, planObject)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Call Update (PUT)
	updateReq := r.client.BandwidthAllocationsAPI.UpdateBandwidthAllocations(ctx).BandwidthAllocations(*unpackedScmObject)

	updatedObject, _, err := updateReq.Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error updating BandwidthAllocation", err.Error())
		return
	}

	if updatedObject != nil {
		packedObject, diags := packBandwidthAllocationsFromSdk(ctx, *updatedObject)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

		var apiData models.BandwidthAllocations
		resp.Diagnostics.Append(packedObject.As(ctx, &apiData, basetypes.ObjectAsOptions{})...)
		if resp.Diagnostics.HasError() {
			return
		}

		// 🟢 FIX: MASKING + VISIBILITY for Update
		if !plan.SpnNameList.IsNull() && !plan.SpnNameList.IsUnknown() {
			if !apiData.SpnNameList.Equal(plan.SpnNameList) {
				var realSpns []string
				apiData.SpnNameList.ElementsAs(ctx, &realSpns, false)
				resp.Diagnostics.AddWarning(
					"API Response Modified Input",
					fmt.Sprintf(
						"The API modified the 'spn_name_list' during update.\n"+
							"Real values returned by API: %v\n"+
							"We have saved your input to state to avoid errors, but drift will be detected on next plan.",
						realSpns,
					),
				)
			}
			apiData.SpnNameList = plan.SpnNameList
		}
		plan = apiData
	}

	plan.Tfid = plan.Name

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *BandwidthAllocationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Starting Delete function for BandwidthAllocation")
	var state models.BandwidthAllocations
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// 🟢 FIX: Fetch the LIVE object first for robust deletion.
	targetName := state.Name.ValueString()
	if state.Tfid.ValueString() != "" && state.Tfid.ValueString() != "unknown" {
		targetName = state.Tfid.ValueString()
	}

	listReq := r.client.BandwidthAllocationsAPI.ListBandwidthAllocations(ctx)
	listResponse, _, err := listReq.Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error reading BandwidthAllocation for deletion", err.Error())
		return
	}

	var foundItem *deployment_services.BandwidthAllocations
	if listResponse != nil && listResponse.Data != nil {
		for _, item := range listResponse.Data {
			if item.Name != "" && item.Name == targetName {
				foundItem = &item
				break
			}
		}
	}

	if foundItem == nil {
		tflog.Warn(ctx, "Resource not found on server during delete, assuming already deleted", map[string]interface{}{"name": targetName})
		return
	}

	// Call Delete using LIVE data
	deleteReq := r.client.BandwidthAllocationsAPI.DeleteBandwidthAllocations(ctx)
	deleteReq = deleteReq.Name(foundItem.Name)
	if len(foundItem.SpnNameList) > 0 {
		deleteReq = deleteReq.SpnNameList(strings.Join(foundItem.SpnNameList, ","))
	}

	_, err = deleteReq.Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error deleting BandwidthAllocation", err.Error())
		return
	}
}

func (r *BandwidthAllocationResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("name"), req, resp)
}
