package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/security_services"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/security_services"
)

// --- Unpacker for DataFilteringProfiles ---
func unpackDataFilteringProfilesToSdk(ctx context.Context, obj types.Object) (*security_services.DataFilteringProfiles, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.DataFilteringProfiles", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.DataFilteringProfiles
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.DataFilteringProfiles
	var d diag.Diagnostics

	// Handling Primitives
	if !model.DataCapture.IsNull() && !model.DataCapture.IsUnknown() {
		sdk.DataCapture = model.DataCapture.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DataCapture", "value": *sdk.DataCapture})
	}

	// Handling Primitives
	if !model.Description.IsNull() && !model.Description.IsUnknown() {
		sdk.Description = model.Description.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Description", "value": *sdk.Description})
	}

	// Handling Primitives
	if !model.Device.IsNull() && !model.Device.IsUnknown() {
		sdk.Device = model.Device.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Device", "value": *sdk.Device})
	}

	// Handling Primitives
	if !model.DisableOverride.IsNull() && !model.DisableOverride.IsUnknown() {
		sdk.DisableOverride = model.DisableOverride.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DisableOverride", "value": *sdk.DisableOverride})
	}

	// Handling Primitives
	if !model.Folder.IsNull() && !model.Folder.IsUnknown() {
		sdk.Folder = model.Folder.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Folder", "value": *sdk.Folder})
	}

	// Handling Primitives
	if !model.Id.IsNull() && !model.Id.IsUnknown() {
		sdk.Id = model.Id.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Id", "value": *sdk.Id})
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
	}

	// Handling Lists
	if !model.Rules.IsNull() && !model.Rules.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field Rules")
		unpacked, d := unpackDataFilteringProfilesRulesInnerListToSdk(ctx, model.Rules)
		diags.Append(d...)
		sdk.Rules = unpacked
	}

	// Handling Primitives
	if !model.Snippet.IsNull() && !model.Snippet.IsUnknown() {
		sdk.Snippet = model.Snippet.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.DataFilteringProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for DataFilteringProfiles ---
func packDataFilteringProfilesFromSdk(ctx context.Context, sdk security_services.DataFilteringProfiles) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.DataFilteringProfiles", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.DataFilteringProfiles
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.DataCapture != nil {
		model.DataCapture = basetypes.NewBoolValue(*sdk.DataCapture)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DataCapture", "value": *sdk.DataCapture})
	} else {
		model.DataCapture = basetypes.NewBoolNull()
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
	if sdk.Device != nil {
		model.Device = basetypes.NewStringValue(*sdk.Device)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Device", "value": *sdk.Device})
	} else {
		model.Device = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.DisableOverride != nil {
		model.DisableOverride = basetypes.NewStringValue(*sdk.DisableOverride)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DisableOverride", "value": *sdk.DisableOverride})
	} else {
		model.DisableOverride = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Folder != nil {
		model.Folder = basetypes.NewStringValue(*sdk.Folder)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Folder", "value": *sdk.Folder})
	} else {
		model.Folder = basetypes.NewStringNull()
	}
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
	if sdk.Name != nil {
		model.Name = basetypes.NewStringValue(*sdk.Name)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
	} else {
		model.Name = basetypes.NewStringNull()
	}
	// Handling Lists
	if sdk.Rules != nil {
		tflog.Debug(ctx, "Packing list of objects for field Rules")
		packed, d := packDataFilteringProfilesRulesInnerListFromSdk(ctx, sdk.Rules)
		diags.Append(d...)
		model.Rules = packed
	} else {
		model.Rules = basetypes.NewListNull(models.DataFilteringProfilesRulesInner{}.AttrType())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Snippet != nil {
		model.Snippet = basetypes.NewStringValue(*sdk.Snippet)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	} else {
		model.Snippet = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.DataFilteringProfiles{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.DataFilteringProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for DataFilteringProfiles ---
func unpackDataFilteringProfilesListToSdk(ctx context.Context, list types.List) ([]security_services.DataFilteringProfiles, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.DataFilteringProfiles")
	diags := diag.Diagnostics{}
	var data []models.DataFilteringProfiles
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.DataFilteringProfiles, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.DataFilteringProfiles{}.AttrTypes(), &item)
		unpacked, d := unpackDataFilteringProfilesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.DataFilteringProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for DataFilteringProfiles ---
func packDataFilteringProfilesListFromSdk(ctx context.Context, sdks []security_services.DataFilteringProfiles) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.DataFilteringProfiles")
	diags := diag.Diagnostics{}
	var data []models.DataFilteringProfiles

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.DataFilteringProfiles
		obj, d := packDataFilteringProfilesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.DataFilteringProfiles{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.DataFilteringProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.DataFilteringProfiles{}.AttrType(), data)
}

// --- Unpacker for DataFilteringProfilesRulesInner ---
func unpackDataFilteringProfilesRulesInnerToSdk(ctx context.Context, obj types.Object) (*security_services.DataFilteringProfilesRulesInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.DataFilteringProfilesRulesInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.DataFilteringProfilesRulesInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.DataFilteringProfilesRulesInner
	var d diag.Diagnostics
	// Handling Primitives
	if !model.AlertThreshold.IsNull() && !model.AlertThreshold.IsUnknown() {
		val := int32(model.AlertThreshold.ValueInt64())
		sdk.AlertThreshold = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "AlertThreshold", "value": *sdk.AlertThreshold})
	}

	// Handling Lists
	if !model.Application.IsNull() && !model.Application.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Application")
		diags.Append(model.Application.ElementsAs(ctx, &sdk.Application, false)...)
	}

	// Handling Primitives
	if !model.BlockThreshold.IsNull() && !model.BlockThreshold.IsUnknown() {
		val := int32(model.BlockThreshold.ValueInt64())
		sdk.BlockThreshold = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "BlockThreshold", "value": *sdk.BlockThreshold})
	}

	// Handling Primitives
	if !model.DataObject.IsNull() && !model.DataObject.IsUnknown() {
		sdk.DataObject = model.DataObject.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DataObject", "value": *sdk.DataObject})
	}

	// Handling Primitives
	if !model.Direction.IsNull() && !model.Direction.IsUnknown() {
		sdk.Direction = model.Direction.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Direction", "value": *sdk.Direction})
	}

	// Handling Lists
	if !model.FileType.IsNull() && !model.FileType.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field FileType")
		diags.Append(model.FileType.ElementsAs(ctx, &sdk.FileType, false)...)
	}

	// Handling Primitives
	if !model.LogSeverity.IsNull() && !model.LogSeverity.IsUnknown() {
		sdk.LogSeverity = model.LogSeverity.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "LogSeverity", "value": *sdk.LogSeverity})
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.DataFilteringProfilesRulesInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for DataFilteringProfilesRulesInner ---
func packDataFilteringProfilesRulesInnerFromSdk(ctx context.Context, sdk security_services.DataFilteringProfilesRulesInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.DataFilteringProfilesRulesInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.DataFilteringProfilesRulesInner
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.AlertThreshold != nil {
		model.AlertThreshold = basetypes.NewInt64Value(int64(*sdk.AlertThreshold))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "AlertThreshold", "value": *sdk.AlertThreshold})
	} else {
		model.AlertThreshold = basetypes.NewInt64Null()
	}
	// Handling Lists
	if sdk.Application != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Application")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Application, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Application)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Application = basetypes.NewListNull(elemType)
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.BlockThreshold != nil {
		model.BlockThreshold = basetypes.NewInt64Value(int64(*sdk.BlockThreshold))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "BlockThreshold", "value": *sdk.BlockThreshold})
	} else {
		model.BlockThreshold = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.DataObject != nil {
		model.DataObject = basetypes.NewStringValue(*sdk.DataObject)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DataObject", "value": *sdk.DataObject})
	} else {
		model.DataObject = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Direction != nil {
		model.Direction = basetypes.NewStringValue(*sdk.Direction)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Direction", "value": *sdk.Direction})
	} else {
		model.Direction = basetypes.NewStringNull()
	}
	// Handling Lists
	if sdk.FileType != nil {
		tflog.Debug(ctx, "Packing list of primitives for field FileType")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.FileType, d = basetypes.NewListValueFrom(ctx, elemType, sdk.FileType)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.FileType = basetypes.NewListNull(elemType)
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.LogSeverity != nil {
		model.LogSeverity = basetypes.NewStringValue(*sdk.LogSeverity)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "LogSeverity", "value": *sdk.LogSeverity})
	} else {
		model.LogSeverity = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Name != nil {
		model.Name = basetypes.NewStringValue(*sdk.Name)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
	} else {
		model.Name = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.DataFilteringProfilesRulesInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.DataFilteringProfilesRulesInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for DataFilteringProfilesRulesInner ---
func unpackDataFilteringProfilesRulesInnerListToSdk(ctx context.Context, list types.List) ([]security_services.DataFilteringProfilesRulesInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.DataFilteringProfilesRulesInner")
	diags := diag.Diagnostics{}
	var data []models.DataFilteringProfilesRulesInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.DataFilteringProfilesRulesInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.DataFilteringProfilesRulesInner{}.AttrTypes(), &item)
		unpacked, d := unpackDataFilteringProfilesRulesInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.DataFilteringProfilesRulesInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for DataFilteringProfilesRulesInner ---
func packDataFilteringProfilesRulesInnerListFromSdk(ctx context.Context, sdks []security_services.DataFilteringProfilesRulesInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.DataFilteringProfilesRulesInner")
	diags := diag.Diagnostics{}
	var data []models.DataFilteringProfilesRulesInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.DataFilteringProfilesRulesInner
		obj, d := packDataFilteringProfilesRulesInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.DataFilteringProfilesRulesInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.DataFilteringProfilesRulesInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.DataFilteringProfilesRulesInner{}.AttrType(), data)
}
