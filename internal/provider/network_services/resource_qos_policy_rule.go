package provider

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/network_services"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/network_services"
	"github.com/paloaltonetworks/terraform-provider-scm/internal/utils"
)

// RESOURCE for SCM QosPolicyRule (Package: network_services)
var (
	_ resource.Resource                = &QosPolicyRuleResource{}
	_ resource.ResourceWithConfigure   = &QosPolicyRuleResource{}
	_ resource.ResourceWithImportState = &QosPolicyRuleResource{}
)

func NewQosPolicyRuleResource() resource.Resource {
	return &QosPolicyRuleResource{}
}

// QosPolicyRuleResource defines the resource implementation.
type QosPolicyRuleResource struct {
	client *network_services.APIClient
}

func (r *QosPolicyRuleResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_qos_policy_rule"
}

func (r *QosPolicyRuleResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = models.QosPolicyRulesResourceSchema
}

func (r *QosPolicyRuleResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *QosPolicyRuleResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Starting Create function for QosPolicyRule")
	var data models.QosPolicyRules

	// 1. Get the plan from Terraform into the data model.
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// --- START: Save original synthetic values ---
	var originalRelativePosition basetypes.StringValue
	var originalTargetRule basetypes.StringValue
	if !data.RelativePosition.IsNull() {
		originalRelativePosition = data.RelativePosition
	}
	if !data.TargetRule.IsNull() {
		originalTargetRule = data.TargetRule
	}
	// --- END: Save original synthetic values ---

	// Unpack the plan to an SCM SDK object.
	planObject, diags := types.ObjectValueFrom(ctx, models.QosPolicyRules{}.AttrTypes(), &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// 2. Unpack the request BODY from data into an SDK object.
	unpackedScmObject, diags := unpackQosPolicyRulesToSdk(ctx, planObject)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating qos_policy_rules on SCM API")

	// 3. Initiate the API request with the body.
	createReq := r.client.QoSRulesAPI.CreateQoSPolicyRules(ctx).QosPolicyRules(*unpackedScmObject)

	// 4. BLOCK 1: Add the request PARAMETERS to the API call.
	if !data.Position.IsNull() {
		createReq = createReq.Position(data.Position.ValueString())
	}

	// 5. Execute the API call.
	createdObject, _, err := createReq.Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error creating qos_policy_rules", err.Error())
		detailedMessage := utils.PrintScmError(err)

		resp.Diagnostics.AddError(
			"SCM Resource Creation Failed: API Request Failed",
			detailedMessage,
		)
		return
	}

	// If positioning attributes are set, move the rule after creation.
	if !data.RelativePosition.IsNull() && createdObject != nil && createdObject.Id != nil {
		relPos := data.RelativePosition.ValueString()
		targetRule := data.TargetRule.ValueString()
		ruleID := *createdObject.Id
		rulebase := data.Position.ValueString()
		if data.Position.IsNull() {
			rulebase = "pre"
		} // Default rulebase

		tflog.Debug(ctx, "Attempting to move newly created qos_policy_rule rule", map[string]any{"rule_id": ruleID, "relative_position": relPos, "target_rule": targetRule, "rulebase": rulebase})
		if err := r.moveRule(ctx, ruleID, relPos, targetRule, rulebase); err != nil {
			resp.Diagnostics.AddWarning("Failed to position qos_policy_rule rule after create", fmt.Sprintf("Rule created but move failed: %s.", err.Error()))
		}
	}

	// 6. Pack the API response back into a Terraform model data.
	packedObject, diags := packQosPolicyRulesFromSdk(ctx, *createdObject)
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

	// Explicitly restore synthetic Terraform attributes using the saved original values.
	data.RelativePosition = originalRelativePosition
	data.TargetRule = originalTargetRule

	// 7. BLOCK 2: Restore the PARAMETER values from the original plan.
	//    This is necessary for parameters that are sent to the API but not returned in the response.
	_ = req.Plan.GetAttribute(ctx, path.Root("position"), &data.Position)

	// FOLDER NORMALIZATION: Handle folder value translation and normalization.
	// This handles both deprecated value translation and Shared/Prisma Access normalization.
	// Get the user's configured folder value from the request plan.
	var userConfiguredFolder basetypes.StringValue
	_ = req.Plan.GetAttribute(ctx, path.Root("folder"), &userConfiguredFolder)

	if !userConfiguredFolder.IsNull() && !userConfiguredFolder.IsUnknown() && !data.Folder.IsNull() {
		userFolderStr := userConfiguredFolder.ValueString()
		apiReturnedFolder := data.Folder.ValueString()

		// First, translate any deprecated values from the API
		translatedFolder := utils.FolderAPIToState(apiReturnedFolder)

		// Then, if user configured "Shared" or "Prisma Access", normalize to match user's choice
		if (userFolderStr == "Shared" || userFolderStr == "Prisma Access") &&
			(translatedFolder == "Shared" || translatedFolder == "Prisma Access") {
			translatedFolder = userFolderStr
		}

		// Set the final normalized value
		data.Folder = basetypes.NewStringValue(translatedFolder)
		tflog.Debug(ctx, "Normalized folder value", map[string]interface{}{
			"configured":   userFolderStr,
			"api_returned": apiReturnedFolder,
			"normalized":   translatedFolder,
		})
	}

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

	tflog.Debug(ctx, "Created qos_policy_rules", map[string]interface{}{"tfid": data.Tfid.ValueString()})
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *QosPolicyRuleResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Step 1 - Initialize a data and savestate of type models.QosPolicyRules - which is the TF schema struct
	tflog.Debug(ctx, "Starting Read function for QosPolicyRule")
	var data, savestate models.QosPolicyRules

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
	tflog.Debug(ctx, "Reading qos_policy_rules from SCM API", map[string]interface{}{"id": objectId})
	getReq := r.client.QoSRulesAPI.GetQoSPolicyRulesByID(ctx, objectId)
	scmObject, httpErr, err := getReq.Execute()
	if err != nil {
		if httpErr != nil && httpErr.StatusCode == http.StatusNotFound {
			tflog.Debug(ctx, "Got no qos_policy_rules on read SCM API. Remove from state to let terraform create", map[string]interface{}{"id": objectId})
			resp.State.RemoveResource(ctx)
		} else {
			tflog.Debug(ctx, "Got an exception on read SCM API. ", map[string]interface{}{"id": objectId})
			resp.Diagnostics.AddError("Error reading qos_policy_rules", err.Error())
			detailedMessage := utils.PrintScmError(err)

			resp.Diagnostics.AddError(
				"SCM Resource Read Failed: API Request Failed",
				detailedMessage,
			)
		}
		return
	}

	// Step 5 - Pack the scm object into a terraform model and put it in data we initialized in step 1
	packedObject, diags := packQosPolicyRulesFromSdk(ctx, *scmObject)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(packedObject.As(ctx, &data, basetypes.ObjectAsOptions{})...)
	if resp.Diagnostics.HasError() {
		return
	}

	// FOLDER NORMALIZATION: Handle folder value translation and normalization.
	// This handles both deprecated value translation and Shared/Prisma Access normalization.
	if !savestate.Folder.IsNull() && !savestate.Folder.IsUnknown() && !data.Folder.IsNull() {
		savedFolderValue := savestate.Folder.ValueString()
		apiReturnedFolder := data.Folder.ValueString()

		// First, translate any deprecated values from the API
		translatedFolder := utils.FolderAPIToState(apiReturnedFolder)

		// Then, if saved state had "Shared" or "Prisma Access", normalize to match saved state
		if (savedFolderValue == "Shared" || savedFolderValue == "Prisma Access") &&
			(translatedFolder == "Shared" || translatedFolder == "Prisma Access") {
			translatedFolder = savedFolderValue
		}

		// Set the final normalized value
		data.Folder = basetypes.NewStringValue(translatedFolder)
		tflog.Debug(ctx, "Normalized folder value", map[string]interface{}{
			"saved_state":  savedFolderValue,
			"api_returned": apiReturnedFolder,
			"normalized":   translatedFolder,
		})
	}

	// Step 7 - Carry over tfid from state back into data
	data.Tfid = savestate.Tfid

	// Step 8 - Set things in params back into data object from the savestate - things like position of security rule
	data.Position = savestate.Position

	// Explicitly restore synthetic Terraform attributes from prior state.
	data.RelativePosition = savestate.RelativePosition
	data.TargetRule = savestate.TargetRule

	// Step 9 - Set folder, snippet, device from params back into data if present

	// --- FOLDER RESTORATION (tokens[0]) ---

	// Use reflection to safely restore the Folder field from the TFID token 0.
	vFolder := reflect.ValueOf(&data).Elem() // Unique variable: vFolder
	fFolder := vFolder.FieldByName("Folder") // Unique variable: fFolder

	if fFolder.IsValid() && fFolder.CanSet() {
		tokenValue := tokens[0]

		// No validation here - just restore the value from TFID
		// Validation happens during ImportState for actual import operations
		// This allows backward compatibility for existing state with old folder values

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

func (r *QosPolicyRuleResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	// Step 1: Initialize a plan and state of type models.QosPolicyRules which is the terraform schema struct
	tflog.Debug(ctx, "Starting Update function for QosPolicyRule")
	var plan, state models.QosPolicyRules

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
	planObject, diags := types.ObjectValueFrom(ctx, models.QosPolicyRules{}.AttrTypes(), &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Step 4: Unpack the plan object to an SCM Object
	unpackedScmObject, diags := unpackQosPolicyRulesToSdk(ctx, planObject)
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

	tflog.Debug(ctx, "Updating qos_policy_rules on SCM API", map[string]interface{}{"id": objectId})
	updateReq := r.client.QoSRulesAPI.UpdateQoSPolicyRulesByID(ctx, objectId).QosPolicyRules(*unpackedScmObject)

	// Step 7: Retain update parameters so we dont lose them
	// ======================== START: ADD THIS BLOCK ========================
	// Apply any operation parameters from the plan.
	// ========================= END: ADD THIS BLOCK =========================

	// Step 8: Make the update call and get an SCM updatedObject
	updatedObject, httpErr, err := updateReq.Execute()
	if err != nil {
		if httpErr != nil && httpErr.StatusCode == http.StatusNotFound {
			tflog.Debug(ctx, "Got no qos_policy_rules on update SCM API. Remove from state to let terraform create", map[string]interface{}{"id": objectId})
			resp.State.RemoveResource(ctx)
		} else {
			tflog.Debug(ctx, "Got an exception on update SCM API. ", map[string]interface{}{"id": objectId})
			resp.Diagnostics.AddError("Error updating qos_policy_rules", err.Error())
			detailedMessage := utils.PrintScmError(err)

			resp.Diagnostics.AddError(
				"SCM Resource Update Failed: API Request Failed",
				detailedMessage,
			)
		}
		return
	}

	// If positioning attributes changed in the plan, move the rule after update.
	if updatedObject != nil && updatedObject.Id != nil {
		needsMove := false
		if plan.RelativePosition.IsNull() && !state.RelativePosition.IsNull() {
			// If relative_position is removed, we don't need to call move API.
			needsMove = false
			tflog.Debug(ctx, "Relative position removed, skipping move for rule", map[string]any{"rule_id": *updatedObject.Id})
		} else if !plan.RelativePosition.IsNull() {
			// Check if position or target actually changed
			positionChanged := state.RelativePosition.IsNull() || !state.RelativePosition.Equal(plan.RelativePosition)
			targetChanged := (state.TargetRule.IsNull() != plan.TargetRule.IsNull()) || (!state.TargetRule.IsNull() && !state.TargetRule.Equal(plan.TargetRule))
			if positionChanged || targetChanged {
				needsMove = true
			}
		}
		if needsMove {
			relPos := plan.RelativePosition.ValueString()
			targetRule := plan.TargetRule.ValueString()
			ruleID := *updatedObject.Id
			rulebase := plan.Position.ValueString()
			if plan.Position.IsNull() {
				rulebase = "pre"
			} // Default rulebase

			tflog.Debug(ctx, "Attempting to move updated qos_policy_rule rule", map[string]any{"rule_id": ruleID, "relative_position": relPos, "target_rule": targetRule, "rulebase": rulebase})
			if err := r.moveRule(ctx, ruleID, relPos, targetRule, rulebase); err != nil {
				resp.Diagnostics.AddWarning("Failed to position qos_policy_rule rule after update", fmt.Sprintf("Rule updated but move failed: %s.", err.Error()))
			}
		} else if !plan.RelativePosition.IsNull() {
			// Log only if a position was specified but didn't change
			tflog.Debug(ctx, "Positioning attributes unchanged, skipping move for rule", map[string]any{"rule_id": *updatedObject.Id})
		}
	}

	// Step 9: Pack the SCM updatedObject into a TF object
	packedObject, diags := packQosPolicyRulesFromSdk(ctx, *updatedObject)
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
	_ = req.Plan.GetAttribute(ctx, path.Root("id"), &plan.Id)

	// FOLDER NORMALIZATION: Handle folder value translation and normalization.
	// This handles both deprecated value translation and Shared/Prisma Access normalization.
	// We need to get the user's configured folder value from the request plan.
	var userConfiguredFolder basetypes.StringValue
	_ = req.Plan.GetAttribute(ctx, path.Root("folder"), &userConfiguredFolder)

	if !userConfiguredFolder.IsNull() && !userConfiguredFolder.IsUnknown() && !plan.Folder.IsNull() {
		userFolderStr := userConfiguredFolder.ValueString()
		apiReturnedFolder := plan.Folder.ValueString()

		// First, translate any deprecated values from the API
		translatedFolder := utils.FolderAPIToState(apiReturnedFolder)

		// Then, if user configured "Shared" or "Prisma Access", normalize to match user's choice
		if (userFolderStr == "Shared" || userFolderStr == "Prisma Access") &&
			(translatedFolder == "Shared" || translatedFolder == "Prisma Access") {
			translatedFolder = userFolderStr
		}

		// Set the final normalized value
		plan.Folder = basetypes.NewStringValue(translatedFolder)
		tflog.Debug(ctx, "Normalized folder value", map[string]interface{}{
			"configured":   userFolderStr,
			"api_returned": apiReturnedFolder,
			"normalized":   translatedFolder,
		})
	}

	// Step 10: Carry over tfid from state into plan
	plan.Tfid = state.Tfid

	// Step 11: Copy write-only attributes from the prior state to the plan for things like position in security rule
	// Restore parameter position from plan
	_ = req.Plan.GetAttribute(ctx, path.Root("position"), &plan.Position)

	// Explicitly restore synthetic move attributes from the PLAN.
	_ = req.Plan.GetAttribute(ctx, path.Root("relative_position"), &plan.RelativePosition)
	_ = req.Plan.GetAttribute(ctx, path.Root("target_rule"), &plan.TargetRule)

	tflog.Debug(ctx, "Updated qos_policy_rules", map[string]interface{}{"tfid": plan.Tfid.ValueString()})
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *QosPolicyRuleResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data models.QosPolicyRules
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

	tflog.Debug(ctx, "Deleting qos_policy_rules", map[string]interface{}{"id": objectId})
	deleteReq := r.client.QoSRulesAPI.DeleteQoSPolicyRulesByID(ctx, objectId)
	_, err := deleteReq.Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error deleting qos_policy_rules", err.Error())
		detailedMessage := utils.PrintScmError(err)

		resp.Diagnostics.AddError(
			"SCM Resource Deleteion Failed: API Request Failed",
			detailedMessage,
		)
	}
}

func (r *QosPolicyRuleResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	// Parse and validate the import ID before storing it - block deprecated folder values
	// Note: Both "Shared" and "Prisma Access" are accepted as valid values
	importID := req.ID
	tokens := strings.Split(importID, ":::")

	if len(tokens) > 0 && tokens[0] != "" {
		folderValue := tokens[0]
		if folderValue == "All Firewalls" {
			resp.Diagnostics.AddError(
				"Invalid Folder Value in Import ID",
				"The folder value 'All Firewalls' is not allowed in import IDs. Please use 'ngfw-shared' instead.\nExample: terraform import scm_qos_policy_rule.example \"ngfw-shared\":::\"<id>\"",
			)
			return
		}
		if folderValue == "Global" {
			resp.Diagnostics.AddError(
				"Invalid Folder Value in Import ID",
				"The folder value 'Global' is not allowed in import IDs. Please use 'All' instead.\nExample: terraform import scm_qos_policy_rule.example \"All\":::\"<id>\"",
			)
			return
		}
		if folderValue == "Explicit Proxy" {
			resp.Diagnostics.AddError(
				"Invalid Folder Value in Import ID",
				"The folder value 'Explicit Proxy' is not allowed in import IDs. Please use 'Mobile Users Explicit Proxy' instead.\nExample: terraform import scm_qos_policy_rule.example \"Mobile Users Explicit Proxy\":::\"<id>\"",
			)
			return
		}
		if folderValue == "Access Agent" {
			resp.Diagnostics.AddError(
				"Invalid Folder Value in Import ID",
				"The folder value 'Access Agent' is not allowed in import IDs. Please use 'Mobile Users' instead.\nExample: terraform import scm_qos_policy_rule.example \"Mobile Users\":::\"<id>\"",
			)
			return
		}
	}

	// If validation passes, store the import ID in tfid
	resource.ImportStatePassthroughID(ctx, path.Root("tfid"), req, resp)
}

// moveRule moves a rule using the SCM Go SDK.
func (r *QosPolicyRuleResource) moveRule(ctx context.Context, ruleID string, relativePosition string, targetRule string, rulebase string) error {
	if relativePosition == "" {
		tflog.Debug(ctx, "No relative_position specified, skipping move.", map[string]any{"rule_id": ruleID})
		return nil
	}
	if rulebase == "" {
		// Default rulebase if not provided, adjust if necessary for your API
		rulebase = "pre"
		tflog.Debug(ctx, "Rulebase not specified, defaulting to 'pre'.", map[string]any{"rule_id": ruleID})
	}

	// Validate that target_rule is provided when needed
	if (relativePosition == "before" || relativePosition == "after") && targetRule == "" {
		// Allow empty target for specific resources if needed, otherwise enforce.
		// Example: return fmt.Errorf("target_rule must be specified when relative_position is 'before' or 'after' for rule %s", ruleID)
		tflog.Warn(ctx, "target_rule is empty for relative position '"+relativePosition+"', proceeding but API might require it.", map[string]any{"rule_id": ruleID})
	}

	// Ensure the client is configured
	if r.client == nil {
		return fmt.Errorf("SCM client is not configured, cannot move rule %s", ruleID)
	}

	// Prepare the SDK request payload for the move operation.
	sdkMovePayload := network_services.RuleBasedMove{
		Destination: relativePosition,
		Rulebase:    rulebase,
		// Only set DestinationRule if targetRule is provided
	}
	if targetRule != "" {
		sdkMovePayload.DestinationRule = &targetRule
	}

	// Construct and execute the SDK request.
	moveReq := r.client.QoSRulesAPI.MoveQoSPolicyRulesByID(ctx, ruleID).RuleBasedMove(sdkMovePayload) // Use MoveOperationID here
	httpResp, err := moveReq.Execute()

	// Handle potential errors
	if err != nil {
		tflog.Error(ctx, "Move request failed", map[string]any{"rule_id": ruleID, "error": err.Error()})
		// Attempt to get more detailed error message
		var statusCode int
		if httpResp != nil {
			statusCode = httpResp.StatusCode
		}
		tflog.Error(ctx, "Move request failed", map[string]any{"rule_id": ruleID, "error": err.Error(), "status_code": statusCode})
		detailedMessage := utils.PrintScmError(err) // Assuming utils.PrintScmError handles nil httpResp gracefully
		return fmt.Errorf("move request for rule %s failed: %s. API response details: %s", ruleID, err.Error(), detailedMessage)
	}

	// Check HTTP status code for success (typically 200 OK for move)
	if httpResp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(httpResp.Body)
		tflog.Error(ctx, "Move request returned non-OK status", map[string]any{"rule_id": ruleID, "status_code": httpResp.StatusCode, "response_body": string(bodyBytes)})
		return fmt.Errorf("move request for rule %s returned status %d: %s", ruleID, httpResp.StatusCode, string(bodyBytes))
	}

	tflog.Info(ctx, "Successfully moved rule", map[string]any{"rule_id": ruleID, "position": relativePosition, "target": targetRule, "rulebase": rulebase})
	return nil
}
