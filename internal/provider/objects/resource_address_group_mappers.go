package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/objects"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/objects"
)

// --- Unpacker for AddressGroups ---
func unpackAddressGroupsToSdk(ctx context.Context, obj types.Object) (*objects.AddressGroups, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.AddressGroups", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.AddressGroups
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.AddressGroups
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

	// Handling Objects
	if !model.Dynamic.IsNull() && !model.Dynamic.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Dynamic")
		unpacked, d := unpackAddressGroupsDynamicToSdk(ctx, model.Dynamic)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Dynamic"})
		}
		if unpacked != nil {
			sdk.Dynamic = unpacked
		}
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

	// Handling Primitives
	if !model.Snippet.IsNull() && !model.Snippet.IsUnknown() {
		sdk.Snippet = model.Snippet.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	}

	// Handling Lists
	if !model.Static.IsNull() && !model.Static.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Static")
		diags.Append(model.Static.ElementsAs(ctx, &sdk.Static, false)...)
	}

	// Handling Lists
	if !model.Tag.IsNull() && !model.Tag.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Tag")
		diags.Append(model.Tag.ElementsAs(ctx, &sdk.Tag, false)...)
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.AddressGroups", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for AddressGroups ---
func packAddressGroupsFromSdk(ctx context.Context, sdk objects.AddressGroups) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.AddressGroups", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.AddressGroups
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
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Dynamic != nil {
		tflog.Debug(ctx, "Packing nested object for field Dynamic")
		packed, d := packAddressGroupsDynamicFromSdk(ctx, *sdk.Dynamic)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Dynamic"})
		}
		model.Dynamic = packed
	} else {
		model.Dynamic = basetypes.NewObjectNull(models.AddressGroupsDynamic{}.AttrTypes())
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
	// Handling Primitives
	// Standard primitive packing
	if sdk.Snippet != nil {
		model.Snippet = basetypes.NewStringValue(*sdk.Snippet)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	} else {
		model.Snippet = basetypes.NewStringNull()
	}
	// Handling Lists
	if sdk.Static != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Static")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Static, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Static)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Static = basetypes.NewListNull(elemType)
	}
	// Handling Lists
	if sdk.Tag != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Tag")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Tag, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Tag)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Tag = basetypes.NewListNull(elemType)
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.AddressGroups{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.AddressGroups", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for AddressGroups ---
func unpackAddressGroupsListToSdk(ctx context.Context, list types.List) ([]objects.AddressGroups, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.AddressGroups")
	diags := diag.Diagnostics{}
	var data []models.AddressGroups
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.AddressGroups, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.AddressGroups{}.AttrTypes(), &item)
		unpacked, d := unpackAddressGroupsToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.AddressGroups", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for AddressGroups ---
func packAddressGroupsListFromSdk(ctx context.Context, sdks []objects.AddressGroups) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.AddressGroups")
	diags := diag.Diagnostics{}
	var data []models.AddressGroups

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.AddressGroups
		obj, d := packAddressGroupsFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.AddressGroups{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.AddressGroups", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.AddressGroups{}.AttrType(), data)
}

// --- Unpacker for AddressGroupsDynamic ---
func unpackAddressGroupsDynamicToSdk(ctx context.Context, obj types.Object) (*objects.AddressGroupsDynamic, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.AddressGroupsDynamic", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.AddressGroupsDynamic
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.AddressGroupsDynamic
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Filter.IsNull() && !model.Filter.IsUnknown() {
		sdk.Filter = model.Filter.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Filter", "value": sdk.Filter})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.AddressGroupsDynamic", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for AddressGroupsDynamic ---
func packAddressGroupsDynamicFromSdk(ctx context.Context, sdk objects.AddressGroupsDynamic) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.AddressGroupsDynamic", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.AddressGroupsDynamic
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.Filter = basetypes.NewStringValue(sdk.Filter)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Filter", "value": sdk.Filter})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.AddressGroupsDynamic{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.AddressGroupsDynamic", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for AddressGroupsDynamic ---
func unpackAddressGroupsDynamicListToSdk(ctx context.Context, list types.List) ([]objects.AddressGroupsDynamic, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.AddressGroupsDynamic")
	diags := diag.Diagnostics{}
	var data []models.AddressGroupsDynamic
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.AddressGroupsDynamic, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.AddressGroupsDynamic{}.AttrTypes(), &item)
		unpacked, d := unpackAddressGroupsDynamicToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.AddressGroupsDynamic", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for AddressGroupsDynamic ---
func packAddressGroupsDynamicListFromSdk(ctx context.Context, sdks []objects.AddressGroupsDynamic) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.AddressGroupsDynamic")
	diags := diag.Diagnostics{}
	var data []models.AddressGroupsDynamic

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.AddressGroupsDynamic
		obj, d := packAddressGroupsDynamicFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.AddressGroupsDynamic{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.AddressGroupsDynamic", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.AddressGroupsDynamic{}.AttrType(), data)
}
