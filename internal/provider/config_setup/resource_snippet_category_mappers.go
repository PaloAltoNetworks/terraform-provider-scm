package provider

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/config_setup"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/config_setup"
)

// --- Unpacker for SnippetCategories ---
func unpackSnippetCategoriesToSdk(ctx context.Context, obj types.Object) (*config_setup.SnippetCategories, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.SnippetCategories", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.SnippetCategories
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk config_setup.SnippetCategories
	var d diag.Diagnostics

	// Handling Primitives
	if !model.CreatedIn.IsNull() && !model.CreatedIn.IsUnknown() {
		if t, parseErr := time.Parse(time.RFC3339, model.CreatedIn.ValueString()); parseErr == nil {
			sdk.CreatedIn = &t
		}
	}

	// Handling Primitives
	if !model.Description.IsNull() && !model.Description.IsUnknown() {
		sdk.Description = model.Description.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Description", "value": *sdk.Description})
	}

	// Handling Primitives
	if !model.DisplayName.IsNull() && !model.DisplayName.IsUnknown() {
		sdk.DisplayName = model.DisplayName.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DisplayName", "value": *sdk.DisplayName})
	}

	// Handling Primitives
	if !model.DonorCreated.IsNull() && !model.DonorCreated.IsUnknown() {
		val := int32(model.DonorCreated.ValueInt64())
		sdk.DonorCreated = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DonorCreated", "value": *sdk.DonorCreated})
	}

	// Handling Primitives
	if !model.DonorSnippetFileId.IsNull() && !model.DonorSnippetFileId.IsUnknown() {
		val := int32(model.DonorSnippetFileId.ValueInt64())
		sdk.DonorSnippetFileId = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DonorSnippetFileId", "value": *sdk.DonorSnippetFileId})
	}

	// Handling Primitives
	if !model.DonorSnippetVersion.IsNull() && !model.DonorSnippetVersion.IsUnknown() {
		val := int32(model.DonorSnippetVersion.ValueInt64())
		sdk.DonorSnippetVersion = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DonorSnippetVersion", "value": *sdk.DonorSnippetVersion})
	}

	// Handling Primitives
	if !model.DonorTenantId.IsNull() && !model.DonorTenantId.IsUnknown() {
		sdk.DonorTenantId = model.DonorTenantId.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DonorTenantId", "value": *sdk.DonorTenantId})
	}

	// Handling Primitives
	if !model.DonorTenantName.IsNull() && !model.DonorTenantName.IsUnknown() {
		sdk.DonorTenantName = model.DonorTenantName.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DonorTenantName", "value": *sdk.DonorTenantName})
	}

	// Handling Primitives
	if !model.DonorTsg.IsNull() && !model.DonorTsg.IsUnknown() {
		sdk.DonorTsg = model.DonorTsg.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DonorTsg", "value": *sdk.DonorTsg})
	}

	// Handling Primitives
	if !model.EnablePrefix.IsNull() && !model.EnablePrefix.IsUnknown() {
		sdk.EnablePrefix = model.EnablePrefix.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "EnablePrefix", "value": *sdk.EnablePrefix})
	}

	// Handling Primitives
	if !model.Error.IsNull() && !model.Error.IsUnknown() {
		sdk.Error = model.Error.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Error", "value": *sdk.Error})
	}

	// Handling Lists
	if !model.Folders.IsNull() && !model.Folders.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field Folders")
		unpacked, d := unpackUsedFoldersListToSdk(ctx, model.Folders)
		diags.Append(d...)
		sdk.Folders = unpacked
	}

	// Handling Primitives
	if !model.Id.IsNull() && !model.Id.IsUnknown() {
		sdk.Id = model.Id.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Id", "value": sdk.Id})
	}

	// Handling Lists
	if !model.Labels.IsNull() && !model.Labels.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Labels")
		diags.Append(model.Labels.ElementsAs(ctx, &sdk.Labels, false)...)
	}

	// Handling Primitives
	if !model.LastUpdate.IsNull() && !model.LastUpdate.IsUnknown() {
		if t, parseErr := time.Parse(time.RFC3339, model.LastUpdate.ValueString()); parseErr == nil {
			sdk.LastUpdate = &t
		}
	}

	// Handling Primitives
	if !model.MsgUuid.IsNull() && !model.MsgUuid.IsUnknown() {
		sdk.MsgUuid = model.MsgUuid.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "MsgUuid", "value": *sdk.MsgUuid})
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	// Handling Primitives
	if !model.Prefix.IsNull() && !model.Prefix.IsUnknown() {
		sdk.Prefix = model.Prefix.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Prefix", "value": *sdk.Prefix})
	}

	// Handling Primitives
	if !model.RecipientPausedUpdate.IsNull() && !model.RecipientPausedUpdate.IsUnknown() {
		sdk.RecipientPausedUpdate = model.RecipientPausedUpdate.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "RecipientPausedUpdate", "value": *sdk.RecipientPausedUpdate})
	}

	// Handling Primitives
	if !model.RecipientTenantId.IsNull() && !model.RecipientTenantId.IsUnknown() {
		sdk.RecipientTenantId = model.RecipientTenantId.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "RecipientTenantId", "value": *sdk.RecipientTenantId})
	}

	// Handling Primitives
	if !model.RecipientTenantName.IsNull() && !model.RecipientTenantName.IsUnknown() {
		sdk.RecipientTenantName = model.RecipientTenantName.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "RecipientTenantName", "value": *sdk.RecipientTenantName})
	}

	// Handling Primitives
	if !model.RecipientTsg.IsNull() && !model.RecipientTsg.IsUnknown() {
		sdk.RecipientTsg = model.RecipientTsg.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "RecipientTsg", "value": *sdk.RecipientTsg})
	}

	// Handling Primitives
	if !model.RecipientValidateBeforeUpdate.IsNull() && !model.RecipientValidateBeforeUpdate.IsUnknown() {
		sdk.RecipientValidateBeforeUpdate = model.RecipientValidateBeforeUpdate.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "RecipientValidateBeforeUpdate", "value": *sdk.RecipientValidateBeforeUpdate})
	}

	// Handling Primitives
	if !model.SharedIn.IsNull() && !model.SharedIn.IsUnknown() {
		sdk.SharedIn = model.SharedIn.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "SharedIn", "value": *sdk.SharedIn})
	}

	// Handling Primitives
	if !model.SnippetUuid.IsNull() && !model.SnippetUuid.IsUnknown() {
		sdk.SnippetUuid = model.SnippetUuid.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "SnippetUuid", "value": *sdk.SnippetUuid})
	}

	// Handling Primitives
	if !model.Status.IsNull() && !model.Status.IsUnknown() {
		sdk.Status = model.Status.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Status", "value": *sdk.Status})
	}

	// Handling Primitives
	if !model.Type.IsNull() && !model.Type.IsUnknown() {
		sdk.Type = model.Type.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Type", "value": *sdk.Type})
	}

	// Handling Primitives
	if !model.Version.IsNull() && !model.Version.IsUnknown() {
		val := int32(model.Version.ValueInt64())
		sdk.Version = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Version", "value": *sdk.Version})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.SnippetCategories", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for SnippetCategories ---
func packSnippetCategoriesFromSdk(ctx context.Context, sdk config_setup.SnippetCategories) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.SnippetCategories", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.SnippetCategories
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.CreatedIn != nil {
		model.CreatedIn = basetypes.NewStringValue(sdk.CreatedIn.Format(time.RFC3339))
		tflog.Debug(ctx, "Packed time pointer", map[string]interface{}{"field": "CreatedIn", "value": sdk.CreatedIn.Format(time.RFC3339)})
	} else {
		model.CreatedIn = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Description != nil {
		model.Description = basetypes.NewStringValue(*sdk.Description)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Description", "value": *sdk.Description})
	} else {
		model.Description = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.DisplayName != nil {
		model.DisplayName = basetypes.NewStringValue(*sdk.DisplayName)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DisplayName", "value": *sdk.DisplayName})
	} else {
		model.DisplayName = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.DonorCreated != nil {
		model.DonorCreated = basetypes.NewInt64Value(int64(*sdk.DonorCreated))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DonorCreated", "value": *sdk.DonorCreated})
	} else {
		model.DonorCreated = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.DonorSnippetFileId != nil {
		model.DonorSnippetFileId = basetypes.NewInt64Value(int64(*sdk.DonorSnippetFileId))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DonorSnippetFileId", "value": *sdk.DonorSnippetFileId})
	} else {
		model.DonorSnippetFileId = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.DonorSnippetVersion != nil {
		model.DonorSnippetVersion = basetypes.NewInt64Value(int64(*sdk.DonorSnippetVersion))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DonorSnippetVersion", "value": *sdk.DonorSnippetVersion})
	} else {
		model.DonorSnippetVersion = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.DonorTenantId != nil {
		model.DonorTenantId = basetypes.NewStringValue(*sdk.DonorTenantId)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DonorTenantId", "value": *sdk.DonorTenantId})
	} else {
		model.DonorTenantId = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.DonorTenantName != nil {
		model.DonorTenantName = basetypes.NewStringValue(*sdk.DonorTenantName)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DonorTenantName", "value": *sdk.DonorTenantName})
	} else {
		model.DonorTenantName = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.DonorTsg != nil {
		model.DonorTsg = basetypes.NewStringValue(*sdk.DonorTsg)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DonorTsg", "value": *sdk.DonorTsg})
	} else {
		model.DonorTsg = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.EnablePrefix != nil {
		model.EnablePrefix = basetypes.NewBoolValue(*sdk.EnablePrefix)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "EnablePrefix", "value": *sdk.EnablePrefix})
	} else {
		model.EnablePrefix = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Error != nil {
		model.Error = basetypes.NewStringValue(*sdk.Error)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Error", "value": *sdk.Error})
	} else {
		model.Error = basetypes.NewStringNull()
	}
	// Handling Lists
	if sdk.Folders != nil {
		tflog.Debug(ctx, "Packing list of objects for field Folders")
		packed, d := packUsedFoldersListFromSdk(ctx, sdk.Folders)
		diags.Append(d...)
		model.Folders = packed
	} else {
		model.Folders = basetypes.NewListNull(models.UsedFolders{}.AttrType())
	}
	// Handling Primitives
	// Standard primitive packing
	model.Id = basetypes.NewStringValue(sdk.Id)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Id", "value": sdk.Id})
	// Handling Lists
	if sdk.Labels != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Labels")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Labels, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Labels)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Labels = basetypes.NewListNull(elemType)
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.LastUpdate != nil {
		model.LastUpdate = basetypes.NewStringValue(sdk.LastUpdate.Format(time.RFC3339))
		tflog.Debug(ctx, "Packed time pointer", map[string]interface{}{"field": "LastUpdate", "value": sdk.LastUpdate.Format(time.RFC3339)})
	} else {
		model.LastUpdate = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.MsgUuid != nil {
		model.MsgUuid = basetypes.NewStringValue(*sdk.MsgUuid)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "MsgUuid", "value": *sdk.MsgUuid})
	} else {
		model.MsgUuid = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	// Handling Primitives
	// Standard primitive packing
	if sdk.Prefix != nil {
		model.Prefix = basetypes.NewStringValue(*sdk.Prefix)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Prefix", "value": *sdk.Prefix})
	} else {
		model.Prefix = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.RecipientPausedUpdate != nil {
		model.RecipientPausedUpdate = basetypes.NewBoolValue(*sdk.RecipientPausedUpdate)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "RecipientPausedUpdate", "value": *sdk.RecipientPausedUpdate})
	} else {
		model.RecipientPausedUpdate = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.RecipientTenantId != nil {
		model.RecipientTenantId = basetypes.NewStringValue(*sdk.RecipientTenantId)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "RecipientTenantId", "value": *sdk.RecipientTenantId})
	} else {
		model.RecipientTenantId = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.RecipientTenantName != nil {
		model.RecipientTenantName = basetypes.NewStringValue(*sdk.RecipientTenantName)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "RecipientTenantName", "value": *sdk.RecipientTenantName})
	} else {
		model.RecipientTenantName = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.RecipientTsg != nil {
		model.RecipientTsg = basetypes.NewStringValue(*sdk.RecipientTsg)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "RecipientTsg", "value": *sdk.RecipientTsg})
	} else {
		model.RecipientTsg = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.RecipientValidateBeforeUpdate != nil {
		model.RecipientValidateBeforeUpdate = basetypes.NewBoolValue(*sdk.RecipientValidateBeforeUpdate)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "RecipientValidateBeforeUpdate", "value": *sdk.RecipientValidateBeforeUpdate})
	} else {
		model.RecipientValidateBeforeUpdate = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.SharedIn != nil {
		model.SharedIn = basetypes.NewStringValue(*sdk.SharedIn)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "SharedIn", "value": *sdk.SharedIn})
	} else {
		model.SharedIn = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.SnippetUuid != nil {
		model.SnippetUuid = basetypes.NewStringValue(*sdk.SnippetUuid)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "SnippetUuid", "value": *sdk.SnippetUuid})
	} else {
		model.SnippetUuid = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Status != nil {
		model.Status = basetypes.NewStringValue(*sdk.Status)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Status", "value": *sdk.Status})
	} else {
		model.Status = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Type != nil {
		model.Type = basetypes.NewStringValue(*sdk.Type)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Type", "value": *sdk.Type})
	} else {
		model.Type = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Version != nil {
		model.Version = basetypes.NewInt64Value(int64(*sdk.Version))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Version", "value": *sdk.Version})
	} else {
		model.Version = basetypes.NewInt64Null()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.SnippetCategories{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.SnippetCategories", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for SnippetCategories ---
func unpackSnippetCategoriesListToSdk(ctx context.Context, list types.List) ([]config_setup.SnippetCategories, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.SnippetCategories")
	diags := diag.Diagnostics{}
	var data []models.SnippetCategories
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]config_setup.SnippetCategories, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.SnippetCategories{}.AttrTypes(), &item)
		unpacked, d := unpackSnippetCategoriesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.SnippetCategories", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for SnippetCategories ---
func packSnippetCategoriesListFromSdk(ctx context.Context, sdks []config_setup.SnippetCategories) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.SnippetCategories")
	diags := diag.Diagnostics{}
	var data []models.SnippetCategories

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.SnippetCategories
		obj, d := packSnippetCategoriesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.SnippetCategories{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.SnippetCategories", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.SnippetCategories{}.AttrType(), data)
}

// --- Unpacker for UsedFolders ---
func unpackUsedFoldersToSdk(ctx context.Context, obj types.Object) (*config_setup.UsedFolders, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.UsedFolders", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.UsedFolders
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk config_setup.UsedFolders
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Id.IsNull() && !model.Id.IsUnknown() {
		sdk.Id = model.Id.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Id", "value": *sdk.Id})
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.UsedFolders", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for UsedFolders ---
func packUsedFoldersFromSdk(ctx context.Context, sdk config_setup.UsedFolders) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.UsedFolders", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.UsedFolders
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Id != nil {
		model.Id = basetypes.NewStringValue(*sdk.Id)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Id", "value": *sdk.Id})
	} else {
		model.Id = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.UsedFolders{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.UsedFolders", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for UsedFolders ---
func unpackUsedFoldersListToSdk(ctx context.Context, list types.List) ([]config_setup.UsedFolders, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.UsedFolders")
	diags := diag.Diagnostics{}
	var data []models.UsedFolders
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]config_setup.UsedFolders, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.UsedFolders{}.AttrTypes(), &item)
		unpacked, d := unpackUsedFoldersToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.UsedFolders", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for UsedFolders ---
func packUsedFoldersListFromSdk(ctx context.Context, sdks []config_setup.UsedFolders) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.UsedFolders")
	diags := diag.Diagnostics{}
	var data []models.UsedFolders

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.UsedFolders
		obj, d := packUsedFoldersFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.UsedFolders{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.UsedFolders", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.UsedFolders{}.AttrType(), data)
}
