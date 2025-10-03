package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/network_services"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/network_services"
)

// --- Unpacker for RouteCommunityLists ---
func unpackRouteCommunityListsToSdk(ctx context.Context, obj types.Object) (*network_services.RouteCommunityLists, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.RouteCommunityLists", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.RouteCommunityLists
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.RouteCommunityLists
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
		sdk.Id = model.Id.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Id", "value": *sdk.Id})
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

	// Handling Objects
	if !model.Type.IsNull() && !model.Type.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Type")
		unpacked, d := unpackRouteCommunityListsTypeToSdk(ctx, model.Type)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Type"})
		}
		if unpacked != nil {
			sdk.Type = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.RouteCommunityLists", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for RouteCommunityLists ---
func packRouteCommunityListsFromSdk(ctx context.Context, sdk network_services.RouteCommunityLists) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.RouteCommunityLists", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.RouteCommunityLists
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
	// Handling Primitives
	// Standard primitive packing
	if sdk.Snippet != nil {
		model.Snippet = basetypes.NewStringValue(*sdk.Snippet)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	} else {
		model.Snippet = basetypes.NewStringNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Type != nil {
		tflog.Debug(ctx, "Packing nested object for field Type")
		packed, d := packRouteCommunityListsTypeFromSdk(ctx, *sdk.Type)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Type"})
		}
		model.Type = packed
	} else {
		model.Type = basetypes.NewObjectNull(models.RouteCommunityListsType{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.RouteCommunityLists{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.RouteCommunityLists", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for RouteCommunityLists ---
func unpackRouteCommunityListsListToSdk(ctx context.Context, list types.List) ([]network_services.RouteCommunityLists, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.RouteCommunityLists")
	diags := diag.Diagnostics{}
	var data []models.RouteCommunityLists
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.RouteCommunityLists, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.RouteCommunityLists{}.AttrTypes(), &item)
		unpacked, d := unpackRouteCommunityListsToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.RouteCommunityLists", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for RouteCommunityLists ---
func packRouteCommunityListsListFromSdk(ctx context.Context, sdks []network_services.RouteCommunityLists) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.RouteCommunityLists")
	diags := diag.Diagnostics{}
	var data []models.RouteCommunityLists

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.RouteCommunityLists
		obj, d := packRouteCommunityListsFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.RouteCommunityLists{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.RouteCommunityLists", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.RouteCommunityLists{}.AttrType(), data)
}

// --- Unpacker for RouteCommunityListsType ---
func unpackRouteCommunityListsTypeToSdk(ctx context.Context, obj types.Object) (*network_services.RouteCommunityListsType, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.RouteCommunityListsType", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.RouteCommunityListsType
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.RouteCommunityListsType
	var d diag.Diagnostics
	// Handling Objects
	if !model.Extended.IsNull() && !model.Extended.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Extended")
		unpacked, d := unpackRouteCommunityListsTypeExtendedToSdk(ctx, model.Extended)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Extended"})
		}
		if unpacked != nil {
			sdk.Extended = unpacked
		}
	}

	// Handling Objects
	if !model.Large.IsNull() && !model.Large.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Large")
		unpacked, d := unpackRouteCommunityListsTypeLargeToSdk(ctx, model.Large)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Large"})
		}
		if unpacked != nil {
			sdk.Large = unpacked
		}
	}

	// Handling Objects
	if !model.Regular.IsNull() && !model.Regular.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Regular")
		unpacked, d := unpackRouteCommunityListsTypeRegularToSdk(ctx, model.Regular)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Regular"})
		}
		if unpacked != nil {
			sdk.Regular = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.RouteCommunityListsType", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for RouteCommunityListsType ---
func packRouteCommunityListsTypeFromSdk(ctx context.Context, sdk network_services.RouteCommunityListsType) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.RouteCommunityListsType", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.RouteCommunityListsType
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Extended != nil {
		tflog.Debug(ctx, "Packing nested object for field Extended")
		packed, d := packRouteCommunityListsTypeExtendedFromSdk(ctx, *sdk.Extended)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Extended"})
		}
		model.Extended = packed
	} else {
		model.Extended = basetypes.NewObjectNull(models.RouteCommunityListsTypeExtended{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Large != nil {
		tflog.Debug(ctx, "Packing nested object for field Large")
		packed, d := packRouteCommunityListsTypeLargeFromSdk(ctx, *sdk.Large)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Large"})
		}
		model.Large = packed
	} else {
		model.Large = basetypes.NewObjectNull(models.RouteCommunityListsTypeLarge{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Regular != nil {
		tflog.Debug(ctx, "Packing nested object for field Regular")
		packed, d := packRouteCommunityListsTypeRegularFromSdk(ctx, *sdk.Regular)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Regular"})
		}
		model.Regular = packed
	} else {
		model.Regular = basetypes.NewObjectNull(models.RouteCommunityListsTypeRegular{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.RouteCommunityListsType{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.RouteCommunityListsType", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for RouteCommunityListsType ---
func unpackRouteCommunityListsTypeListToSdk(ctx context.Context, list types.List) ([]network_services.RouteCommunityListsType, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.RouteCommunityListsType")
	diags := diag.Diagnostics{}
	var data []models.RouteCommunityListsType
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.RouteCommunityListsType, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.RouteCommunityListsType{}.AttrTypes(), &item)
		unpacked, d := unpackRouteCommunityListsTypeToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.RouteCommunityListsType", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for RouteCommunityListsType ---
func packRouteCommunityListsTypeListFromSdk(ctx context.Context, sdks []network_services.RouteCommunityListsType) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.RouteCommunityListsType")
	diags := diag.Diagnostics{}
	var data []models.RouteCommunityListsType

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.RouteCommunityListsType
		obj, d := packRouteCommunityListsTypeFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.RouteCommunityListsType{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.RouteCommunityListsType", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.RouteCommunityListsType{}.AttrType(), data)
}

// --- Unpacker for RouteCommunityListsTypeExtended ---
func unpackRouteCommunityListsTypeExtendedToSdk(ctx context.Context, obj types.Object) (*network_services.RouteCommunityListsTypeExtended, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.RouteCommunityListsTypeExtended", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.RouteCommunityListsTypeExtended
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.RouteCommunityListsTypeExtended
	var d diag.Diagnostics
	// Handling Lists
	if !model.ExtendedEntry.IsNull() && !model.ExtendedEntry.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field ExtendedEntry")
		unpacked, d := unpackRouteCommunityListsTypeExtendedExtendedEntryInnerListToSdk(ctx, model.ExtendedEntry)
		diags.Append(d...)
		sdk.ExtendedEntry = unpacked
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.RouteCommunityListsTypeExtended", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for RouteCommunityListsTypeExtended ---
func packRouteCommunityListsTypeExtendedFromSdk(ctx context.Context, sdk network_services.RouteCommunityListsTypeExtended) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.RouteCommunityListsTypeExtended", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.RouteCommunityListsTypeExtended
	var d diag.Diagnostics
	// Handling Lists
	if sdk.ExtendedEntry != nil {
		tflog.Debug(ctx, "Packing list of objects for field ExtendedEntry")
		packed, d := packRouteCommunityListsTypeExtendedExtendedEntryInnerListFromSdk(ctx, sdk.ExtendedEntry)
		diags.Append(d...)
		model.ExtendedEntry = packed
	} else {
		model.ExtendedEntry = basetypes.NewListNull(models.RouteCommunityListsTypeExtendedExtendedEntryInner{}.AttrType())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.RouteCommunityListsTypeExtended{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.RouteCommunityListsTypeExtended", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for RouteCommunityListsTypeExtended ---
func unpackRouteCommunityListsTypeExtendedListToSdk(ctx context.Context, list types.List) ([]network_services.RouteCommunityListsTypeExtended, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.RouteCommunityListsTypeExtended")
	diags := diag.Diagnostics{}
	var data []models.RouteCommunityListsTypeExtended
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.RouteCommunityListsTypeExtended, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.RouteCommunityListsTypeExtended{}.AttrTypes(), &item)
		unpacked, d := unpackRouteCommunityListsTypeExtendedToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.RouteCommunityListsTypeExtended", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for RouteCommunityListsTypeExtended ---
func packRouteCommunityListsTypeExtendedListFromSdk(ctx context.Context, sdks []network_services.RouteCommunityListsTypeExtended) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.RouteCommunityListsTypeExtended")
	diags := diag.Diagnostics{}
	var data []models.RouteCommunityListsTypeExtended

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.RouteCommunityListsTypeExtended
		obj, d := packRouteCommunityListsTypeExtendedFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.RouteCommunityListsTypeExtended{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.RouteCommunityListsTypeExtended", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.RouteCommunityListsTypeExtended{}.AttrType(), data)
}

// --- Unpacker for RouteCommunityListsTypeExtendedExtendedEntryInner ---
func unpackRouteCommunityListsTypeExtendedExtendedEntryInnerToSdk(ctx context.Context, obj types.Object) (*network_services.RouteCommunityListsTypeExtendedExtendedEntryInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.RouteCommunityListsTypeExtendedExtendedEntryInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.RouteCommunityListsTypeExtendedExtendedEntryInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.RouteCommunityListsTypeExtendedExtendedEntryInner
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Action.IsNull() && !model.Action.IsUnknown() {
		sdk.Action = model.Action.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Action", "value": *sdk.Action})
	}

	// Handling Lists
	if !model.LcRegex.IsNull() && !model.LcRegex.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field LcRegex")
		diags.Append(model.LcRegex.ElementsAs(ctx, &sdk.LcRegex, false)...)
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		val := int32(model.Name.ValueInt64())
		sdk.Name = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.RouteCommunityListsTypeExtendedExtendedEntryInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for RouteCommunityListsTypeExtendedExtendedEntryInner ---
func packRouteCommunityListsTypeExtendedExtendedEntryInnerFromSdk(ctx context.Context, sdk network_services.RouteCommunityListsTypeExtendedExtendedEntryInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.RouteCommunityListsTypeExtendedExtendedEntryInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.RouteCommunityListsTypeExtendedExtendedEntryInner
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Action != nil {
		model.Action = basetypes.NewStringValue(*sdk.Action)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Action", "value": *sdk.Action})
	} else {
		model.Action = basetypes.NewStringNull()
	}
	// Handling Lists
	if sdk.LcRegex != nil {
		tflog.Debug(ctx, "Packing list of primitives for field LcRegex")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.LcRegex, d = basetypes.NewListValueFrom(ctx, elemType, sdk.LcRegex)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.LcRegex = basetypes.NewListNull(elemType)
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Name != nil {
		model.Name = basetypes.NewInt64Value(int64(*sdk.Name))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
	} else {
		model.Name = basetypes.NewInt64Null()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.RouteCommunityListsTypeExtendedExtendedEntryInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.RouteCommunityListsTypeExtendedExtendedEntryInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for RouteCommunityListsTypeExtendedExtendedEntryInner ---
func unpackRouteCommunityListsTypeExtendedExtendedEntryInnerListToSdk(ctx context.Context, list types.List) ([]network_services.RouteCommunityListsTypeExtendedExtendedEntryInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.RouteCommunityListsTypeExtendedExtendedEntryInner")
	diags := diag.Diagnostics{}
	var data []models.RouteCommunityListsTypeExtendedExtendedEntryInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.RouteCommunityListsTypeExtendedExtendedEntryInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.RouteCommunityListsTypeExtendedExtendedEntryInner{}.AttrTypes(), &item)
		unpacked, d := unpackRouteCommunityListsTypeExtendedExtendedEntryInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.RouteCommunityListsTypeExtendedExtendedEntryInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for RouteCommunityListsTypeExtendedExtendedEntryInner ---
func packRouteCommunityListsTypeExtendedExtendedEntryInnerListFromSdk(ctx context.Context, sdks []network_services.RouteCommunityListsTypeExtendedExtendedEntryInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.RouteCommunityListsTypeExtendedExtendedEntryInner")
	diags := diag.Diagnostics{}
	var data []models.RouteCommunityListsTypeExtendedExtendedEntryInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.RouteCommunityListsTypeExtendedExtendedEntryInner
		obj, d := packRouteCommunityListsTypeExtendedExtendedEntryInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.RouteCommunityListsTypeExtendedExtendedEntryInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.RouteCommunityListsTypeExtendedExtendedEntryInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.RouteCommunityListsTypeExtendedExtendedEntryInner{}.AttrType(), data)
}

// --- Unpacker for RouteCommunityListsTypeLarge ---
func unpackRouteCommunityListsTypeLargeToSdk(ctx context.Context, obj types.Object) (*network_services.RouteCommunityListsTypeLarge, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.RouteCommunityListsTypeLarge", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.RouteCommunityListsTypeLarge
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.RouteCommunityListsTypeLarge
	var d diag.Diagnostics
	// Handling Lists
	if !model.LargeEntry.IsNull() && !model.LargeEntry.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field LargeEntry")
		unpacked, d := unpackRouteCommunityListsTypeLargeLargeEntryInnerListToSdk(ctx, model.LargeEntry)
		diags.Append(d...)
		sdk.LargeEntry = unpacked
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.RouteCommunityListsTypeLarge", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for RouteCommunityListsTypeLarge ---
func packRouteCommunityListsTypeLargeFromSdk(ctx context.Context, sdk network_services.RouteCommunityListsTypeLarge) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.RouteCommunityListsTypeLarge", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.RouteCommunityListsTypeLarge
	var d diag.Diagnostics
	// Handling Lists
	if sdk.LargeEntry != nil {
		tflog.Debug(ctx, "Packing list of objects for field LargeEntry")
		packed, d := packRouteCommunityListsTypeLargeLargeEntryInnerListFromSdk(ctx, sdk.LargeEntry)
		diags.Append(d...)
		model.LargeEntry = packed
	} else {
		model.LargeEntry = basetypes.NewListNull(models.RouteCommunityListsTypeLargeLargeEntryInner{}.AttrType())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.RouteCommunityListsTypeLarge{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.RouteCommunityListsTypeLarge", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for RouteCommunityListsTypeLarge ---
func unpackRouteCommunityListsTypeLargeListToSdk(ctx context.Context, list types.List) ([]network_services.RouteCommunityListsTypeLarge, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.RouteCommunityListsTypeLarge")
	diags := diag.Diagnostics{}
	var data []models.RouteCommunityListsTypeLarge
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.RouteCommunityListsTypeLarge, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.RouteCommunityListsTypeLarge{}.AttrTypes(), &item)
		unpacked, d := unpackRouteCommunityListsTypeLargeToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.RouteCommunityListsTypeLarge", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for RouteCommunityListsTypeLarge ---
func packRouteCommunityListsTypeLargeListFromSdk(ctx context.Context, sdks []network_services.RouteCommunityListsTypeLarge) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.RouteCommunityListsTypeLarge")
	diags := diag.Diagnostics{}
	var data []models.RouteCommunityListsTypeLarge

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.RouteCommunityListsTypeLarge
		obj, d := packRouteCommunityListsTypeLargeFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.RouteCommunityListsTypeLarge{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.RouteCommunityListsTypeLarge", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.RouteCommunityListsTypeLarge{}.AttrType(), data)
}

// --- Unpacker for RouteCommunityListsTypeLargeLargeEntryInner ---
func unpackRouteCommunityListsTypeLargeLargeEntryInnerToSdk(ctx context.Context, obj types.Object) (*network_services.RouteCommunityListsTypeLargeLargeEntryInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.RouteCommunityListsTypeLargeLargeEntryInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.RouteCommunityListsTypeLargeLargeEntryInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.RouteCommunityListsTypeLargeLargeEntryInner
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Action.IsNull() && !model.Action.IsUnknown() {
		sdk.Action = model.Action.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Action", "value": *sdk.Action})
	}

	// Handling Lists
	if !model.LcRegex.IsNull() && !model.LcRegex.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field LcRegex")
		diags.Append(model.LcRegex.ElementsAs(ctx, &sdk.LcRegex, false)...)
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		val := int32(model.Name.ValueInt64())
		sdk.Name = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.RouteCommunityListsTypeLargeLargeEntryInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for RouteCommunityListsTypeLargeLargeEntryInner ---
func packRouteCommunityListsTypeLargeLargeEntryInnerFromSdk(ctx context.Context, sdk network_services.RouteCommunityListsTypeLargeLargeEntryInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.RouteCommunityListsTypeLargeLargeEntryInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.RouteCommunityListsTypeLargeLargeEntryInner
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Action != nil {
		model.Action = basetypes.NewStringValue(*sdk.Action)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Action", "value": *sdk.Action})
	} else {
		model.Action = basetypes.NewStringNull()
	}
	// Handling Lists
	if sdk.LcRegex != nil {
		tflog.Debug(ctx, "Packing list of primitives for field LcRegex")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.LcRegex, d = basetypes.NewListValueFrom(ctx, elemType, sdk.LcRegex)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.LcRegex = basetypes.NewListNull(elemType)
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Name != nil {
		model.Name = basetypes.NewInt64Value(int64(*sdk.Name))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
	} else {
		model.Name = basetypes.NewInt64Null()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.RouteCommunityListsTypeLargeLargeEntryInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.RouteCommunityListsTypeLargeLargeEntryInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for RouteCommunityListsTypeLargeLargeEntryInner ---
func unpackRouteCommunityListsTypeLargeLargeEntryInnerListToSdk(ctx context.Context, list types.List) ([]network_services.RouteCommunityListsTypeLargeLargeEntryInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.RouteCommunityListsTypeLargeLargeEntryInner")
	diags := diag.Diagnostics{}
	var data []models.RouteCommunityListsTypeLargeLargeEntryInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.RouteCommunityListsTypeLargeLargeEntryInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.RouteCommunityListsTypeLargeLargeEntryInner{}.AttrTypes(), &item)
		unpacked, d := unpackRouteCommunityListsTypeLargeLargeEntryInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.RouteCommunityListsTypeLargeLargeEntryInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for RouteCommunityListsTypeLargeLargeEntryInner ---
func packRouteCommunityListsTypeLargeLargeEntryInnerListFromSdk(ctx context.Context, sdks []network_services.RouteCommunityListsTypeLargeLargeEntryInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.RouteCommunityListsTypeLargeLargeEntryInner")
	diags := diag.Diagnostics{}
	var data []models.RouteCommunityListsTypeLargeLargeEntryInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.RouteCommunityListsTypeLargeLargeEntryInner
		obj, d := packRouteCommunityListsTypeLargeLargeEntryInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.RouteCommunityListsTypeLargeLargeEntryInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.RouteCommunityListsTypeLargeLargeEntryInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.RouteCommunityListsTypeLargeLargeEntryInner{}.AttrType(), data)
}

// --- Unpacker for RouteCommunityListsTypeRegular ---
func unpackRouteCommunityListsTypeRegularToSdk(ctx context.Context, obj types.Object) (*network_services.RouteCommunityListsTypeRegular, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.RouteCommunityListsTypeRegular", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.RouteCommunityListsTypeRegular
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.RouteCommunityListsTypeRegular
	var d diag.Diagnostics
	// Handling Lists
	if !model.RegularEntry.IsNull() && !model.RegularEntry.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field RegularEntry")
		unpacked, d := unpackRouteCommunityListsTypeRegularRegularEntryInnerListToSdk(ctx, model.RegularEntry)
		diags.Append(d...)
		sdk.RegularEntry = unpacked
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.RouteCommunityListsTypeRegular", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for RouteCommunityListsTypeRegular ---
func packRouteCommunityListsTypeRegularFromSdk(ctx context.Context, sdk network_services.RouteCommunityListsTypeRegular) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.RouteCommunityListsTypeRegular", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.RouteCommunityListsTypeRegular
	var d diag.Diagnostics
	// Handling Lists
	if sdk.RegularEntry != nil {
		tflog.Debug(ctx, "Packing list of objects for field RegularEntry")
		packed, d := packRouteCommunityListsTypeRegularRegularEntryInnerListFromSdk(ctx, sdk.RegularEntry)
		diags.Append(d...)
		model.RegularEntry = packed
	} else {
		model.RegularEntry = basetypes.NewListNull(models.RouteCommunityListsTypeRegularRegularEntryInner{}.AttrType())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.RouteCommunityListsTypeRegular{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.RouteCommunityListsTypeRegular", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for RouteCommunityListsTypeRegular ---
func unpackRouteCommunityListsTypeRegularListToSdk(ctx context.Context, list types.List) ([]network_services.RouteCommunityListsTypeRegular, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.RouteCommunityListsTypeRegular")
	diags := diag.Diagnostics{}
	var data []models.RouteCommunityListsTypeRegular
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.RouteCommunityListsTypeRegular, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.RouteCommunityListsTypeRegular{}.AttrTypes(), &item)
		unpacked, d := unpackRouteCommunityListsTypeRegularToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.RouteCommunityListsTypeRegular", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for RouteCommunityListsTypeRegular ---
func packRouteCommunityListsTypeRegularListFromSdk(ctx context.Context, sdks []network_services.RouteCommunityListsTypeRegular) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.RouteCommunityListsTypeRegular")
	diags := diag.Diagnostics{}
	var data []models.RouteCommunityListsTypeRegular

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.RouteCommunityListsTypeRegular
		obj, d := packRouteCommunityListsTypeRegularFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.RouteCommunityListsTypeRegular{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.RouteCommunityListsTypeRegular", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.RouteCommunityListsTypeRegular{}.AttrType(), data)
}

// --- Unpacker for RouteCommunityListsTypeRegularRegularEntryInner ---
func unpackRouteCommunityListsTypeRegularRegularEntryInnerToSdk(ctx context.Context, obj types.Object) (*network_services.RouteCommunityListsTypeRegularRegularEntryInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.RouteCommunityListsTypeRegularRegularEntryInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.RouteCommunityListsTypeRegularRegularEntryInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.RouteCommunityListsTypeRegularRegularEntryInner
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Action.IsNull() && !model.Action.IsUnknown() {
		sdk.Action = model.Action.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Action", "value": *sdk.Action})
	}

	// Handling Lists
	if !model.Community.IsNull() && !model.Community.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Community")
		diags.Append(model.Community.ElementsAs(ctx, &sdk.Community, false)...)
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		val := int32(model.Name.ValueInt64())
		sdk.Name = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.RouteCommunityListsTypeRegularRegularEntryInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for RouteCommunityListsTypeRegularRegularEntryInner ---
func packRouteCommunityListsTypeRegularRegularEntryInnerFromSdk(ctx context.Context, sdk network_services.RouteCommunityListsTypeRegularRegularEntryInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.RouteCommunityListsTypeRegularRegularEntryInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.RouteCommunityListsTypeRegularRegularEntryInner
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Action != nil {
		model.Action = basetypes.NewStringValue(*sdk.Action)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Action", "value": *sdk.Action})
	} else {
		model.Action = basetypes.NewStringNull()
	}
	// Handling Lists
	if sdk.Community != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Community")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Community, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Community)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Community = basetypes.NewListNull(elemType)
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Name != nil {
		model.Name = basetypes.NewInt64Value(int64(*sdk.Name))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
	} else {
		model.Name = basetypes.NewInt64Null()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.RouteCommunityListsTypeRegularRegularEntryInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.RouteCommunityListsTypeRegularRegularEntryInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for RouteCommunityListsTypeRegularRegularEntryInner ---
func unpackRouteCommunityListsTypeRegularRegularEntryInnerListToSdk(ctx context.Context, list types.List) ([]network_services.RouteCommunityListsTypeRegularRegularEntryInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.RouteCommunityListsTypeRegularRegularEntryInner")
	diags := diag.Diagnostics{}
	var data []models.RouteCommunityListsTypeRegularRegularEntryInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.RouteCommunityListsTypeRegularRegularEntryInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.RouteCommunityListsTypeRegularRegularEntryInner{}.AttrTypes(), &item)
		unpacked, d := unpackRouteCommunityListsTypeRegularRegularEntryInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.RouteCommunityListsTypeRegularRegularEntryInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for RouteCommunityListsTypeRegularRegularEntryInner ---
func packRouteCommunityListsTypeRegularRegularEntryInnerListFromSdk(ctx context.Context, sdks []network_services.RouteCommunityListsTypeRegularRegularEntryInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.RouteCommunityListsTypeRegularRegularEntryInner")
	diags := diag.Diagnostics{}
	var data []models.RouteCommunityListsTypeRegularRegularEntryInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.RouteCommunityListsTypeRegularRegularEntryInner
		obj, d := packRouteCommunityListsTypeRegularRegularEntryInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.RouteCommunityListsTypeRegularRegularEntryInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.RouteCommunityListsTypeRegularRegularEntryInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.RouteCommunityListsTypeRegularRegularEntryInner{}.AttrType(), data)
}
