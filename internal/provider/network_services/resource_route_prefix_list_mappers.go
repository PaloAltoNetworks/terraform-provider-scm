package provider

import (
	"context"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/network_services"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/network_services"
)

// --- Unpacker for RoutePrefixLists ---
func unpackRoutePrefixListsToSdk(ctx context.Context, obj types.Object) (*network_services.RoutePrefixLists, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.RoutePrefixLists", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.RoutePrefixLists
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.RoutePrefixLists
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
		unpacked, d := unpackRoutePrefixListsTypeToSdk(ctx, model.Type)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Type"})
		}
		if unpacked != nil {
			sdk.Type = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.RoutePrefixLists", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for RoutePrefixLists ---
func packRoutePrefixListsFromSdk(ctx context.Context, sdk network_services.RoutePrefixLists) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.RoutePrefixLists", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.RoutePrefixLists
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
		packed, d := packRoutePrefixListsTypeFromSdk(ctx, *sdk.Type)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Type"})
		}
		model.Type = packed
	} else {
		model.Type = basetypes.NewObjectNull(models.RoutePrefixListsType{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.RoutePrefixLists{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.RoutePrefixLists", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for RoutePrefixLists ---
func unpackRoutePrefixListsListToSdk(ctx context.Context, list types.List) ([]network_services.RoutePrefixLists, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.RoutePrefixLists")
	diags := diag.Diagnostics{}
	var data []models.RoutePrefixLists
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.RoutePrefixLists, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.RoutePrefixLists{}.AttrTypes(), &item)
		unpacked, d := unpackRoutePrefixListsToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.RoutePrefixLists", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for RoutePrefixLists ---
func packRoutePrefixListsListFromSdk(ctx context.Context, sdks []network_services.RoutePrefixLists) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.RoutePrefixLists")
	diags := diag.Diagnostics{}
	var data []models.RoutePrefixLists

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.RoutePrefixLists
		obj, d := packRoutePrefixListsFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.RoutePrefixLists{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.RoutePrefixLists", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.RoutePrefixLists{}.AttrType(), data)
}

// --- Unpacker for RoutePrefixListsType ---
func unpackRoutePrefixListsTypeToSdk(ctx context.Context, obj types.Object) (*network_services.RoutePrefixListsType, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.RoutePrefixListsType", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.RoutePrefixListsType
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.RoutePrefixListsType
	var d diag.Diagnostics
	// Handling Objects
	if !model.Ipv4.IsNull() && !model.Ipv4.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Ipv4")
		unpacked, d := unpackRoutePrefixListsTypeIpv4ToSdk(ctx, model.Ipv4)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Ipv4"})
		}
		if unpacked != nil {
			sdk.Ipv4 = *unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.RoutePrefixListsType", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for RoutePrefixListsType ---
func packRoutePrefixListsTypeFromSdk(ctx context.Context, sdk network_services.RoutePrefixListsType) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.RoutePrefixListsType", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.RoutePrefixListsType
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	// Logic for non-pointer / value-type nested objects
	if !reflect.ValueOf(sdk.Ipv4).IsZero() {
		tflog.Debug(ctx, "Packing nested object for field Ipv4")
		packed, d := packRoutePrefixListsTypeIpv4FromSdk(ctx, sdk.Ipv4)
		diags.Append(d...)
		model.Ipv4 = packed
	} else {
		model.Ipv4 = basetypes.NewObjectNull(models.RoutePrefixListsTypeIpv4{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.RoutePrefixListsType{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.RoutePrefixListsType", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for RoutePrefixListsType ---
func unpackRoutePrefixListsTypeListToSdk(ctx context.Context, list types.List) ([]network_services.RoutePrefixListsType, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.RoutePrefixListsType")
	diags := diag.Diagnostics{}
	var data []models.RoutePrefixListsType
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.RoutePrefixListsType, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.RoutePrefixListsType{}.AttrTypes(), &item)
		unpacked, d := unpackRoutePrefixListsTypeToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.RoutePrefixListsType", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for RoutePrefixListsType ---
func packRoutePrefixListsTypeListFromSdk(ctx context.Context, sdks []network_services.RoutePrefixListsType) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.RoutePrefixListsType")
	diags := diag.Diagnostics{}
	var data []models.RoutePrefixListsType

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.RoutePrefixListsType
		obj, d := packRoutePrefixListsTypeFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.RoutePrefixListsType{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.RoutePrefixListsType", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.RoutePrefixListsType{}.AttrType(), data)
}

// --- Unpacker for RoutePrefixListsTypeIpv4 ---
func unpackRoutePrefixListsTypeIpv4ToSdk(ctx context.Context, obj types.Object) (*network_services.RoutePrefixListsTypeIpv4, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.RoutePrefixListsTypeIpv4", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.RoutePrefixListsTypeIpv4
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.RoutePrefixListsTypeIpv4
	var d diag.Diagnostics
	// Handling Lists
	if !model.Ipv4Entry.IsNull() && !model.Ipv4Entry.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field Ipv4Entry")
		unpacked, d := unpackRoutePrefixListsTypeIpv4Ipv4EntryInnerListToSdk(ctx, model.Ipv4Entry)
		diags.Append(d...)
		sdk.Ipv4Entry = unpacked
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.RoutePrefixListsTypeIpv4", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for RoutePrefixListsTypeIpv4 ---
func packRoutePrefixListsTypeIpv4FromSdk(ctx context.Context, sdk network_services.RoutePrefixListsTypeIpv4) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.RoutePrefixListsTypeIpv4", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.RoutePrefixListsTypeIpv4
	var d diag.Diagnostics
	// Handling Lists
	if sdk.Ipv4Entry != nil {
		tflog.Debug(ctx, "Packing list of objects for field Ipv4Entry")
		packed, d := packRoutePrefixListsTypeIpv4Ipv4EntryInnerListFromSdk(ctx, sdk.Ipv4Entry)
		diags.Append(d...)
		model.Ipv4Entry = packed
	} else {
		model.Ipv4Entry = basetypes.NewListNull(models.RoutePrefixListsTypeIpv4Ipv4EntryInner{}.AttrType())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.RoutePrefixListsTypeIpv4{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.RoutePrefixListsTypeIpv4", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for RoutePrefixListsTypeIpv4 ---
func unpackRoutePrefixListsTypeIpv4ListToSdk(ctx context.Context, list types.List) ([]network_services.RoutePrefixListsTypeIpv4, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.RoutePrefixListsTypeIpv4")
	diags := diag.Diagnostics{}
	var data []models.RoutePrefixListsTypeIpv4
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.RoutePrefixListsTypeIpv4, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.RoutePrefixListsTypeIpv4{}.AttrTypes(), &item)
		unpacked, d := unpackRoutePrefixListsTypeIpv4ToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.RoutePrefixListsTypeIpv4", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for RoutePrefixListsTypeIpv4 ---
func packRoutePrefixListsTypeIpv4ListFromSdk(ctx context.Context, sdks []network_services.RoutePrefixListsTypeIpv4) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.RoutePrefixListsTypeIpv4")
	diags := diag.Diagnostics{}
	var data []models.RoutePrefixListsTypeIpv4

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.RoutePrefixListsTypeIpv4
		obj, d := packRoutePrefixListsTypeIpv4FromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.RoutePrefixListsTypeIpv4{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.RoutePrefixListsTypeIpv4", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.RoutePrefixListsTypeIpv4{}.AttrType(), data)
}

// --- Unpacker for RoutePrefixListsTypeIpv4Ipv4EntryInner ---
func unpackRoutePrefixListsTypeIpv4Ipv4EntryInnerToSdk(ctx context.Context, obj types.Object) (*network_services.RoutePrefixListsTypeIpv4Ipv4EntryInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.RoutePrefixListsTypeIpv4Ipv4EntryInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.RoutePrefixListsTypeIpv4Ipv4EntryInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.RoutePrefixListsTypeIpv4Ipv4EntryInner
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Action.IsNull() && !model.Action.IsUnknown() {
		sdk.Action = model.Action.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Action", "value": *sdk.Action})
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		val := int32(model.Name.ValueInt64())
		sdk.Name = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
	}

	// Handling Objects
	if !model.Prefix.IsNull() && !model.Prefix.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Prefix")
		unpacked, d := unpackRoutePrefixListsTypeIpv4Ipv4EntryInnerPrefixToSdk(ctx, model.Prefix)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Prefix"})
		}
		if unpacked != nil {
			sdk.Prefix = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.RoutePrefixListsTypeIpv4Ipv4EntryInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for RoutePrefixListsTypeIpv4Ipv4EntryInner ---
func packRoutePrefixListsTypeIpv4Ipv4EntryInnerFromSdk(ctx context.Context, sdk network_services.RoutePrefixListsTypeIpv4Ipv4EntryInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.RoutePrefixListsTypeIpv4Ipv4EntryInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.RoutePrefixListsTypeIpv4Ipv4EntryInner
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Action != nil {
		model.Action = basetypes.NewStringValue(*sdk.Action)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Action", "value": *sdk.Action})
	} else {
		model.Action = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Name != nil {
		model.Name = basetypes.NewInt64Value(int64(*sdk.Name))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
	} else {
		model.Name = basetypes.NewInt64Null()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Prefix != nil {
		tflog.Debug(ctx, "Packing nested object for field Prefix")
		packed, d := packRoutePrefixListsTypeIpv4Ipv4EntryInnerPrefixFromSdk(ctx, *sdk.Prefix)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Prefix"})
		}
		model.Prefix = packed
	} else {
		model.Prefix = basetypes.NewObjectNull(models.RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefix{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.RoutePrefixListsTypeIpv4Ipv4EntryInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.RoutePrefixListsTypeIpv4Ipv4EntryInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for RoutePrefixListsTypeIpv4Ipv4EntryInner ---
func unpackRoutePrefixListsTypeIpv4Ipv4EntryInnerListToSdk(ctx context.Context, list types.List) ([]network_services.RoutePrefixListsTypeIpv4Ipv4EntryInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.RoutePrefixListsTypeIpv4Ipv4EntryInner")
	diags := diag.Diagnostics{}
	var data []models.RoutePrefixListsTypeIpv4Ipv4EntryInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.RoutePrefixListsTypeIpv4Ipv4EntryInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.RoutePrefixListsTypeIpv4Ipv4EntryInner{}.AttrTypes(), &item)
		unpacked, d := unpackRoutePrefixListsTypeIpv4Ipv4EntryInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.RoutePrefixListsTypeIpv4Ipv4EntryInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for RoutePrefixListsTypeIpv4Ipv4EntryInner ---
func packRoutePrefixListsTypeIpv4Ipv4EntryInnerListFromSdk(ctx context.Context, sdks []network_services.RoutePrefixListsTypeIpv4Ipv4EntryInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.RoutePrefixListsTypeIpv4Ipv4EntryInner")
	diags := diag.Diagnostics{}
	var data []models.RoutePrefixListsTypeIpv4Ipv4EntryInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.RoutePrefixListsTypeIpv4Ipv4EntryInner
		obj, d := packRoutePrefixListsTypeIpv4Ipv4EntryInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.RoutePrefixListsTypeIpv4Ipv4EntryInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.RoutePrefixListsTypeIpv4Ipv4EntryInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.RoutePrefixListsTypeIpv4Ipv4EntryInner{}.AttrType(), data)
}

// --- Unpacker for RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefix ---
func unpackRoutePrefixListsTypeIpv4Ipv4EntryInnerPrefixToSdk(ctx context.Context, obj types.Object) (*network_services.RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefix, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefix", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefix
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefix
	var d diag.Diagnostics
	// Handling Objects
	if !model.Entry.IsNull() && !model.Entry.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Entry")
		unpacked, d := unpackRoutePrefixListsTypeIpv4Ipv4EntryInnerPrefixEntryToSdk(ctx, model.Entry)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Entry"})
		}
		if unpacked != nil {
			sdk.Entry = unpacked
		}
	}

	// Handling Primitives
	if !model.Network.IsNull() && !model.Network.IsUnknown() {
		sdk.Network = model.Network.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Network", "value": *sdk.Network})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefix", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefix ---
func packRoutePrefixListsTypeIpv4Ipv4EntryInnerPrefixFromSdk(ctx context.Context, sdk network_services.RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefix) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefix", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefix
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Entry != nil {
		tflog.Debug(ctx, "Packing nested object for field Entry")
		packed, d := packRoutePrefixListsTypeIpv4Ipv4EntryInnerPrefixEntryFromSdk(ctx, *sdk.Entry)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Entry"})
		}
		model.Entry = packed
	} else {
		model.Entry = basetypes.NewObjectNull(models.RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefixEntry{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Network != nil {
		model.Network = basetypes.NewStringValue(*sdk.Network)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Network", "value": *sdk.Network})
	} else {
		model.Network = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefix{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefix", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefix ---
func unpackRoutePrefixListsTypeIpv4Ipv4EntryInnerPrefixListToSdk(ctx context.Context, list types.List) ([]network_services.RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefix, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefix")
	diags := diag.Diagnostics{}
	var data []models.RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefix
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefix, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefix{}.AttrTypes(), &item)
		unpacked, d := unpackRoutePrefixListsTypeIpv4Ipv4EntryInnerPrefixToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefix", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefix ---
func packRoutePrefixListsTypeIpv4Ipv4EntryInnerPrefixListFromSdk(ctx context.Context, sdks []network_services.RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefix) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefix")
	diags := diag.Diagnostics{}
	var data []models.RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefix

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefix
		obj, d := packRoutePrefixListsTypeIpv4Ipv4EntryInnerPrefixFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefix{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefix", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefix{}.AttrType(), data)
}

// --- Unpacker for RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefixEntry ---
func unpackRoutePrefixListsTypeIpv4Ipv4EntryInnerPrefixEntryToSdk(ctx context.Context, obj types.Object) (*network_services.RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefixEntry, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefixEntry", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefixEntry
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefixEntry
	var d diag.Diagnostics
	// Handling Primitives
	if !model.GreaterThanOrEqual.IsNull() && !model.GreaterThanOrEqual.IsUnknown() {
		val := int32(model.GreaterThanOrEqual.ValueInt64())
		sdk.GreaterThanOrEqual = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "GreaterThanOrEqual", "value": *sdk.GreaterThanOrEqual})
	}

	// Handling Primitives
	if !model.LessThanOrEqual.IsNull() && !model.LessThanOrEqual.IsUnknown() {
		val := int32(model.LessThanOrEqual.ValueInt64())
		sdk.LessThanOrEqual = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "LessThanOrEqual", "value": *sdk.LessThanOrEqual})
	}

	// Handling Primitives
	if !model.Network.IsNull() && !model.Network.IsUnknown() {
		sdk.Network = model.Network.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Network", "value": *sdk.Network})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefixEntry", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefixEntry ---
func packRoutePrefixListsTypeIpv4Ipv4EntryInnerPrefixEntryFromSdk(ctx context.Context, sdk network_services.RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefixEntry) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefixEntry", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefixEntry
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.GreaterThanOrEqual != nil {
		model.GreaterThanOrEqual = basetypes.NewInt64Value(int64(*sdk.GreaterThanOrEqual))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "GreaterThanOrEqual", "value": *sdk.GreaterThanOrEqual})
	} else {
		model.GreaterThanOrEqual = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.LessThanOrEqual != nil {
		model.LessThanOrEqual = basetypes.NewInt64Value(int64(*sdk.LessThanOrEqual))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "LessThanOrEqual", "value": *sdk.LessThanOrEqual})
	} else {
		model.LessThanOrEqual = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Network != nil {
		model.Network = basetypes.NewStringValue(*sdk.Network)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Network", "value": *sdk.Network})
	} else {
		model.Network = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefixEntry{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefixEntry", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefixEntry ---
func unpackRoutePrefixListsTypeIpv4Ipv4EntryInnerPrefixEntryListToSdk(ctx context.Context, list types.List) ([]network_services.RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefixEntry, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefixEntry")
	diags := diag.Diagnostics{}
	var data []models.RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefixEntry
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefixEntry, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefixEntry{}.AttrTypes(), &item)
		unpacked, d := unpackRoutePrefixListsTypeIpv4Ipv4EntryInnerPrefixEntryToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefixEntry", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefixEntry ---
func packRoutePrefixListsTypeIpv4Ipv4EntryInnerPrefixEntryListFromSdk(ctx context.Context, sdks []network_services.RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefixEntry) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefixEntry")
	diags := diag.Diagnostics{}
	var data []models.RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefixEntry

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefixEntry
		obj, d := packRoutePrefixListsTypeIpv4Ipv4EntryInnerPrefixEntryFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefixEntry{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefixEntry", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefixEntry{}.AttrType(), data)
}
