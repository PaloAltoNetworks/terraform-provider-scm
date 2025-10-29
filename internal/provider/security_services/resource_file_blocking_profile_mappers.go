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

// --- Unpacker for FileBlockingProfiles ---
func unpackFileBlockingProfilesToSdk(ctx context.Context, obj types.Object) (*security_services.FileBlockingProfiles, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.FileBlockingProfiles", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.FileBlockingProfiles
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.FileBlockingProfiles
	var d diag.Diagnostics

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
	if !model.Folder.IsNull() && !model.Folder.IsUnknown() {
		sdk.Folder = model.Folder.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Folder", "value": *sdk.Folder})
	}

	// Handling Primitives
	if !model.Id.IsNull() && !model.Id.IsUnknown() {
		sdk.Id = model.Id.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Id", "value": sdk.Id})
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	// Handling Lists
	if !model.Rules.IsNull() && !model.Rules.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field Rules")
		unpacked, d := unpackFileBlockingProfilesRulesInnerListToSdk(ctx, model.Rules)
		diags.Append(d...)
		sdk.Rules = unpacked
	}

	// Handling Primitives
	if !model.Snippet.IsNull() && !model.Snippet.IsUnknown() {
		sdk.Snippet = model.Snippet.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.FileBlockingProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for FileBlockingProfiles ---
func packFileBlockingProfilesFromSdk(ctx context.Context, sdk security_services.FileBlockingProfiles) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.FileBlockingProfiles", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.FileBlockingProfiles
	var d diag.Diagnostics
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
	if sdk.Folder != nil {
		model.Folder = basetypes.NewStringValue(*sdk.Folder)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Folder", "value": *sdk.Folder})
	} else {
		model.Folder = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	model.Id = basetypes.NewStringValue(sdk.Id)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Id", "value": sdk.Id})
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	// Handling Lists
	if sdk.Rules != nil {
		tflog.Debug(ctx, "Packing list of objects for field Rules")
		packed, d := packFileBlockingProfilesRulesInnerListFromSdk(ctx, sdk.Rules)
		diags.Append(d...)
		model.Rules = packed
	} else {
		model.Rules = basetypes.NewListNull(models.FileBlockingProfilesRulesInner{}.AttrType())
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

	obj, d := types.ObjectValueFrom(ctx, models.FileBlockingProfiles{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.FileBlockingProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for FileBlockingProfiles ---
func unpackFileBlockingProfilesListToSdk(ctx context.Context, list types.List) ([]security_services.FileBlockingProfiles, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.FileBlockingProfiles")
	diags := diag.Diagnostics{}
	var data []models.FileBlockingProfiles
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.FileBlockingProfiles, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.FileBlockingProfiles{}.AttrTypes(), &item)
		unpacked, d := unpackFileBlockingProfilesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.FileBlockingProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for FileBlockingProfiles ---
func packFileBlockingProfilesListFromSdk(ctx context.Context, sdks []security_services.FileBlockingProfiles) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.FileBlockingProfiles")
	diags := diag.Diagnostics{}
	var data []models.FileBlockingProfiles

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.FileBlockingProfiles
		obj, d := packFileBlockingProfilesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.FileBlockingProfiles{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.FileBlockingProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.FileBlockingProfiles{}.AttrType(), data)
}

// --- Unpacker for FileBlockingProfilesRulesInner ---
func unpackFileBlockingProfilesRulesInnerToSdk(ctx context.Context, obj types.Object) (*security_services.FileBlockingProfilesRulesInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.FileBlockingProfilesRulesInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.FileBlockingProfilesRulesInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.FileBlockingProfilesRulesInner
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Action.IsNull() && !model.Action.IsUnknown() {
		sdk.Action = model.Action.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Action", "value": sdk.Action})
	}

	// Handling Lists
	if !model.Application.IsNull() && !model.Application.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Application")
		diags.Append(model.Application.ElementsAs(ctx, &sdk.Application, false)...)
	}

	// Handling Primitives
	if !model.Direction.IsNull() && !model.Direction.IsUnknown() {
		sdk.Direction = model.Direction.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Direction", "value": sdk.Direction})
	}

	// Handling Lists
	if !model.FileType.IsNull() && !model.FileType.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field FileType")
		diags.Append(model.FileType.ElementsAs(ctx, &sdk.FileType, false)...)
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.FileBlockingProfilesRulesInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for FileBlockingProfilesRulesInner ---
func packFileBlockingProfilesRulesInnerFromSdk(ctx context.Context, sdk security_services.FileBlockingProfilesRulesInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.FileBlockingProfilesRulesInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.FileBlockingProfilesRulesInner
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.Action = basetypes.NewStringValue(sdk.Action)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Action", "value": sdk.Action})
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
	model.Direction = basetypes.NewStringValue(sdk.Direction)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Direction", "value": sdk.Direction})
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
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.FileBlockingProfilesRulesInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.FileBlockingProfilesRulesInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for FileBlockingProfilesRulesInner ---
func unpackFileBlockingProfilesRulesInnerListToSdk(ctx context.Context, list types.List) ([]security_services.FileBlockingProfilesRulesInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.FileBlockingProfilesRulesInner")
	diags := diag.Diagnostics{}
	var data []models.FileBlockingProfilesRulesInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.FileBlockingProfilesRulesInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.FileBlockingProfilesRulesInner{}.AttrTypes(), &item)
		unpacked, d := unpackFileBlockingProfilesRulesInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.FileBlockingProfilesRulesInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for FileBlockingProfilesRulesInner ---
func packFileBlockingProfilesRulesInnerListFromSdk(ctx context.Context, sdks []security_services.FileBlockingProfilesRulesInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.FileBlockingProfilesRulesInner")
	diags := diag.Diagnostics{}
	var data []models.FileBlockingProfilesRulesInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.FileBlockingProfilesRulesInner
		obj, d := packFileBlockingProfilesRulesInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.FileBlockingProfilesRulesInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.FileBlockingProfilesRulesInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.FileBlockingProfilesRulesInner{}.AttrType(), data)
}
