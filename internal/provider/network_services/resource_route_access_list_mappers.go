package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/network_services"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/network_services"
)

// --- Unpacker for RouteAccessLists ---
func unpackRouteAccessListsToSdk(ctx context.Context, obj types.Object) (*network_services.RouteAccessLists, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.RouteAccessLists", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.RouteAccessLists
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.RouteAccessLists
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
		unpacked, d := unpackRouteAccessListsTypeToSdk(ctx, model.Type)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Type"})
		}
		if unpacked != nil {
			sdk.Type = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.RouteAccessLists", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for RouteAccessLists ---
func packRouteAccessListsFromSdk(ctx context.Context, sdk network_services.RouteAccessLists) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.RouteAccessLists", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.RouteAccessLists
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
		packed, d := packRouteAccessListsTypeFromSdk(ctx, *sdk.Type)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Type"})
		}
		model.Type = packed
	} else {
		model.Type = basetypes.NewObjectNull(models.RouteAccessListsType{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.RouteAccessLists{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.RouteAccessLists", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for RouteAccessLists ---
func unpackRouteAccessListsListToSdk(ctx context.Context, list types.List) ([]network_services.RouteAccessLists, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.RouteAccessLists")
	diags := diag.Diagnostics{}
	var data []models.RouteAccessLists
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.RouteAccessLists, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.RouteAccessLists{}.AttrTypes(), &item)
		unpacked, d := unpackRouteAccessListsToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.RouteAccessLists", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for RouteAccessLists ---
func packRouteAccessListsListFromSdk(ctx context.Context, sdks []network_services.RouteAccessLists) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.RouteAccessLists")
	diags := diag.Diagnostics{}
	var data []models.RouteAccessLists

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.RouteAccessLists
		obj, d := packRouteAccessListsFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.RouteAccessLists{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.RouteAccessLists", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.RouteAccessLists{}.AttrType(), data)
}

// --- Unpacker for RouteAccessListsType ---
func unpackRouteAccessListsTypeToSdk(ctx context.Context, obj types.Object) (*network_services.RouteAccessListsType, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.RouteAccessListsType", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.RouteAccessListsType
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.RouteAccessListsType
	var d diag.Diagnostics
	// Handling Objects
	if !model.Ipv4.IsNull() && !model.Ipv4.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Ipv4")
		unpacked, d := unpackRouteAccessListsTypeIpv4ToSdk(ctx, model.Ipv4)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Ipv4"})
		}
		if unpacked != nil {
			sdk.Ipv4 = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.RouteAccessListsType", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for RouteAccessListsType ---
func packRouteAccessListsTypeFromSdk(ctx context.Context, sdk network_services.RouteAccessListsType) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.RouteAccessListsType", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.RouteAccessListsType
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Ipv4 != nil {
		tflog.Debug(ctx, "Packing nested object for field Ipv4")
		packed, d := packRouteAccessListsTypeIpv4FromSdk(ctx, *sdk.Ipv4)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Ipv4"})
		}
		model.Ipv4 = packed
	} else {
		model.Ipv4 = basetypes.NewObjectNull(models.RouteAccessListsTypeIpv4{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.RouteAccessListsType{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.RouteAccessListsType", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for RouteAccessListsType ---
func unpackRouteAccessListsTypeListToSdk(ctx context.Context, list types.List) ([]network_services.RouteAccessListsType, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.RouteAccessListsType")
	diags := diag.Diagnostics{}
	var data []models.RouteAccessListsType
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.RouteAccessListsType, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.RouteAccessListsType{}.AttrTypes(), &item)
		unpacked, d := unpackRouteAccessListsTypeToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.RouteAccessListsType", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for RouteAccessListsType ---
func packRouteAccessListsTypeListFromSdk(ctx context.Context, sdks []network_services.RouteAccessListsType) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.RouteAccessListsType")
	diags := diag.Diagnostics{}
	var data []models.RouteAccessListsType

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.RouteAccessListsType
		obj, d := packRouteAccessListsTypeFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.RouteAccessListsType{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.RouteAccessListsType", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.RouteAccessListsType{}.AttrType(), data)
}

// --- Unpacker for RouteAccessListsTypeIpv4 ---
func unpackRouteAccessListsTypeIpv4ToSdk(ctx context.Context, obj types.Object) (*network_services.RouteAccessListsTypeIpv4, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.RouteAccessListsTypeIpv4", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.RouteAccessListsTypeIpv4
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.RouteAccessListsTypeIpv4
	var d diag.Diagnostics
	// Handling Lists
	if !model.Ipv4Entry.IsNull() && !model.Ipv4Entry.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field Ipv4Entry")
		unpacked, d := unpackRouteAccessListsTypeIpv4Ipv4EntryInnerListToSdk(ctx, model.Ipv4Entry)
		diags.Append(d...)
		sdk.Ipv4Entry = unpacked
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.RouteAccessListsTypeIpv4", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for RouteAccessListsTypeIpv4 ---
func packRouteAccessListsTypeIpv4FromSdk(ctx context.Context, sdk network_services.RouteAccessListsTypeIpv4) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.RouteAccessListsTypeIpv4", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.RouteAccessListsTypeIpv4
	var d diag.Diagnostics
	// Handling Lists
	if sdk.Ipv4Entry != nil {
		tflog.Debug(ctx, "Packing list of objects for field Ipv4Entry")
		packed, d := packRouteAccessListsTypeIpv4Ipv4EntryInnerListFromSdk(ctx, sdk.Ipv4Entry)
		diags.Append(d...)
		model.Ipv4Entry = packed
	} else {
		model.Ipv4Entry = basetypes.NewListNull(models.RouteAccessListsTypeIpv4Ipv4EntryInner{}.AttrType())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.RouteAccessListsTypeIpv4{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.RouteAccessListsTypeIpv4", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for RouteAccessListsTypeIpv4 ---
func unpackRouteAccessListsTypeIpv4ListToSdk(ctx context.Context, list types.List) ([]network_services.RouteAccessListsTypeIpv4, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.RouteAccessListsTypeIpv4")
	diags := diag.Diagnostics{}
	var data []models.RouteAccessListsTypeIpv4
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.RouteAccessListsTypeIpv4, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.RouteAccessListsTypeIpv4{}.AttrTypes(), &item)
		unpacked, d := unpackRouteAccessListsTypeIpv4ToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.RouteAccessListsTypeIpv4", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for RouteAccessListsTypeIpv4 ---
func packRouteAccessListsTypeIpv4ListFromSdk(ctx context.Context, sdks []network_services.RouteAccessListsTypeIpv4) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.RouteAccessListsTypeIpv4")
	diags := diag.Diagnostics{}
	var data []models.RouteAccessListsTypeIpv4

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.RouteAccessListsTypeIpv4
		obj, d := packRouteAccessListsTypeIpv4FromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.RouteAccessListsTypeIpv4{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.RouteAccessListsTypeIpv4", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.RouteAccessListsTypeIpv4{}.AttrType(), data)
}

// --- Unpacker for RouteAccessListsTypeIpv4Ipv4EntryInner ---
func unpackRouteAccessListsTypeIpv4Ipv4EntryInnerToSdk(ctx context.Context, obj types.Object) (*network_services.RouteAccessListsTypeIpv4Ipv4EntryInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.RouteAccessListsTypeIpv4Ipv4EntryInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.RouteAccessListsTypeIpv4Ipv4EntryInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.RouteAccessListsTypeIpv4Ipv4EntryInner
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Action.IsNull() && !model.Action.IsUnknown() {
		sdk.Action = model.Action.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Action", "value": *sdk.Action})
	}

	// Handling Objects
	if !model.DestinationAddress.IsNull() && !model.DestinationAddress.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field DestinationAddress")
		unpacked, d := unpackRouteAccessListsTypeIpv4Ipv4EntryInnerDestinationAddressToSdk(ctx, model.DestinationAddress)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "DestinationAddress"})
		}
		if unpacked != nil {
			sdk.DestinationAddress = unpacked
		}
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		val := int32(model.Name.ValueInt64())
		sdk.Name = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
	}

	// Handling Objects
	if !model.SourceAddress.IsNull() && !model.SourceAddress.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field SourceAddress")
		unpacked, d := unpackRouteAccessListsTypeIpv4Ipv4EntryInnerSourceAddressToSdk(ctx, model.SourceAddress)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "SourceAddress"})
		}
		if unpacked != nil {
			sdk.SourceAddress = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.RouteAccessListsTypeIpv4Ipv4EntryInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for RouteAccessListsTypeIpv4Ipv4EntryInner ---
func packRouteAccessListsTypeIpv4Ipv4EntryInnerFromSdk(ctx context.Context, sdk network_services.RouteAccessListsTypeIpv4Ipv4EntryInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.RouteAccessListsTypeIpv4Ipv4EntryInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.RouteAccessListsTypeIpv4Ipv4EntryInner
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Action != nil {
		model.Action = basetypes.NewStringValue(*sdk.Action)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Action", "value": *sdk.Action})
	} else {
		model.Action = basetypes.NewStringNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.DestinationAddress != nil {
		tflog.Debug(ctx, "Packing nested object for field DestinationAddress")
		packed, d := packRouteAccessListsTypeIpv4Ipv4EntryInnerDestinationAddressFromSdk(ctx, *sdk.DestinationAddress)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "DestinationAddress"})
		}
		model.DestinationAddress = packed
	} else {
		model.DestinationAddress = basetypes.NewObjectNull(models.RouteAccessListsTypeIpv4Ipv4EntryInnerDestinationAddress{}.AttrTypes())
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
	if sdk.SourceAddress != nil {
		tflog.Debug(ctx, "Packing nested object for field SourceAddress")
		packed, d := packRouteAccessListsTypeIpv4Ipv4EntryInnerSourceAddressFromSdk(ctx, *sdk.SourceAddress)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "SourceAddress"})
		}
		model.SourceAddress = packed
	} else {
		model.SourceAddress = basetypes.NewObjectNull(models.RouteAccessListsTypeIpv4Ipv4EntryInnerSourceAddress{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.RouteAccessListsTypeIpv4Ipv4EntryInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.RouteAccessListsTypeIpv4Ipv4EntryInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for RouteAccessListsTypeIpv4Ipv4EntryInner ---
func unpackRouteAccessListsTypeIpv4Ipv4EntryInnerListToSdk(ctx context.Context, list types.List) ([]network_services.RouteAccessListsTypeIpv4Ipv4EntryInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.RouteAccessListsTypeIpv4Ipv4EntryInner")
	diags := diag.Diagnostics{}
	var data []models.RouteAccessListsTypeIpv4Ipv4EntryInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.RouteAccessListsTypeIpv4Ipv4EntryInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.RouteAccessListsTypeIpv4Ipv4EntryInner{}.AttrTypes(), &item)
		unpacked, d := unpackRouteAccessListsTypeIpv4Ipv4EntryInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.RouteAccessListsTypeIpv4Ipv4EntryInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for RouteAccessListsTypeIpv4Ipv4EntryInner ---
func packRouteAccessListsTypeIpv4Ipv4EntryInnerListFromSdk(ctx context.Context, sdks []network_services.RouteAccessListsTypeIpv4Ipv4EntryInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.RouteAccessListsTypeIpv4Ipv4EntryInner")
	diags := diag.Diagnostics{}
	var data []models.RouteAccessListsTypeIpv4Ipv4EntryInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.RouteAccessListsTypeIpv4Ipv4EntryInner
		obj, d := packRouteAccessListsTypeIpv4Ipv4EntryInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.RouteAccessListsTypeIpv4Ipv4EntryInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.RouteAccessListsTypeIpv4Ipv4EntryInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.RouteAccessListsTypeIpv4Ipv4EntryInner{}.AttrType(), data)
}

// --- Unpacker for RouteAccessListsTypeIpv4Ipv4EntryInnerDestinationAddress ---
func unpackRouteAccessListsTypeIpv4Ipv4EntryInnerDestinationAddressToSdk(ctx context.Context, obj types.Object) (*network_services.RouteAccessListsTypeIpv4Ipv4EntryInnerDestinationAddress, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.RouteAccessListsTypeIpv4Ipv4EntryInnerDestinationAddress", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.RouteAccessListsTypeIpv4Ipv4EntryInnerDestinationAddress
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.RouteAccessListsTypeIpv4Ipv4EntryInnerDestinationAddress
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Address.IsNull() && !model.Address.IsUnknown() {
		sdk.Address = model.Address.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Address", "value": *sdk.Address})
	}

	// Handling Primitives
	if !model.Wildcard.IsNull() && !model.Wildcard.IsUnknown() {
		sdk.Wildcard = model.Wildcard.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Wildcard", "value": *sdk.Wildcard})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.RouteAccessListsTypeIpv4Ipv4EntryInnerDestinationAddress", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for RouteAccessListsTypeIpv4Ipv4EntryInnerDestinationAddress ---
func packRouteAccessListsTypeIpv4Ipv4EntryInnerDestinationAddressFromSdk(ctx context.Context, sdk network_services.RouteAccessListsTypeIpv4Ipv4EntryInnerDestinationAddress) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.RouteAccessListsTypeIpv4Ipv4EntryInnerDestinationAddress", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.RouteAccessListsTypeIpv4Ipv4EntryInnerDestinationAddress
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Address != nil {
		model.Address = basetypes.NewStringValue(*sdk.Address)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Address", "value": *sdk.Address})
	} else {
		model.Address = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Wildcard != nil {
		model.Wildcard = basetypes.NewStringValue(*sdk.Wildcard)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Wildcard", "value": *sdk.Wildcard})
	} else {
		model.Wildcard = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.RouteAccessListsTypeIpv4Ipv4EntryInnerDestinationAddress{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.RouteAccessListsTypeIpv4Ipv4EntryInnerDestinationAddress", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for RouteAccessListsTypeIpv4Ipv4EntryInnerDestinationAddress ---
func unpackRouteAccessListsTypeIpv4Ipv4EntryInnerDestinationAddressListToSdk(ctx context.Context, list types.List) ([]network_services.RouteAccessListsTypeIpv4Ipv4EntryInnerDestinationAddress, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.RouteAccessListsTypeIpv4Ipv4EntryInnerDestinationAddress")
	diags := diag.Diagnostics{}
	var data []models.RouteAccessListsTypeIpv4Ipv4EntryInnerDestinationAddress
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.RouteAccessListsTypeIpv4Ipv4EntryInnerDestinationAddress, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.RouteAccessListsTypeIpv4Ipv4EntryInnerDestinationAddress{}.AttrTypes(), &item)
		unpacked, d := unpackRouteAccessListsTypeIpv4Ipv4EntryInnerDestinationAddressToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.RouteAccessListsTypeIpv4Ipv4EntryInnerDestinationAddress", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for RouteAccessListsTypeIpv4Ipv4EntryInnerDestinationAddress ---
func packRouteAccessListsTypeIpv4Ipv4EntryInnerDestinationAddressListFromSdk(ctx context.Context, sdks []network_services.RouteAccessListsTypeIpv4Ipv4EntryInnerDestinationAddress) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.RouteAccessListsTypeIpv4Ipv4EntryInnerDestinationAddress")
	diags := diag.Diagnostics{}
	var data []models.RouteAccessListsTypeIpv4Ipv4EntryInnerDestinationAddress

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.RouteAccessListsTypeIpv4Ipv4EntryInnerDestinationAddress
		obj, d := packRouteAccessListsTypeIpv4Ipv4EntryInnerDestinationAddressFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.RouteAccessListsTypeIpv4Ipv4EntryInnerDestinationAddress{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.RouteAccessListsTypeIpv4Ipv4EntryInnerDestinationAddress", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.RouteAccessListsTypeIpv4Ipv4EntryInnerDestinationAddress{}.AttrType(), data)
}

// --- Unpacker for RouteAccessListsTypeIpv4Ipv4EntryInnerSourceAddress ---
func unpackRouteAccessListsTypeIpv4Ipv4EntryInnerSourceAddressToSdk(ctx context.Context, obj types.Object) (*network_services.RouteAccessListsTypeIpv4Ipv4EntryInnerSourceAddress, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.RouteAccessListsTypeIpv4Ipv4EntryInnerSourceAddress", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.RouteAccessListsTypeIpv4Ipv4EntryInnerSourceAddress
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.RouteAccessListsTypeIpv4Ipv4EntryInnerSourceAddress
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Address.IsNull() && !model.Address.IsUnknown() {
		sdk.Address = model.Address.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Address", "value": *sdk.Address})
	}

	// Handling Primitives
	if !model.Wildcard.IsNull() && !model.Wildcard.IsUnknown() {
		sdk.Wildcard = model.Wildcard.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Wildcard", "value": *sdk.Wildcard})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.RouteAccessListsTypeIpv4Ipv4EntryInnerSourceAddress", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for RouteAccessListsTypeIpv4Ipv4EntryInnerSourceAddress ---
func packRouteAccessListsTypeIpv4Ipv4EntryInnerSourceAddressFromSdk(ctx context.Context, sdk network_services.RouteAccessListsTypeIpv4Ipv4EntryInnerSourceAddress) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.RouteAccessListsTypeIpv4Ipv4EntryInnerSourceAddress", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.RouteAccessListsTypeIpv4Ipv4EntryInnerSourceAddress
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Address != nil {
		model.Address = basetypes.NewStringValue(*sdk.Address)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Address", "value": *sdk.Address})
	} else {
		model.Address = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Wildcard != nil {
		model.Wildcard = basetypes.NewStringValue(*sdk.Wildcard)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Wildcard", "value": *sdk.Wildcard})
	} else {
		model.Wildcard = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.RouteAccessListsTypeIpv4Ipv4EntryInnerSourceAddress{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.RouteAccessListsTypeIpv4Ipv4EntryInnerSourceAddress", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for RouteAccessListsTypeIpv4Ipv4EntryInnerSourceAddress ---
func unpackRouteAccessListsTypeIpv4Ipv4EntryInnerSourceAddressListToSdk(ctx context.Context, list types.List) ([]network_services.RouteAccessListsTypeIpv4Ipv4EntryInnerSourceAddress, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.RouteAccessListsTypeIpv4Ipv4EntryInnerSourceAddress")
	diags := diag.Diagnostics{}
	var data []models.RouteAccessListsTypeIpv4Ipv4EntryInnerSourceAddress
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.RouteAccessListsTypeIpv4Ipv4EntryInnerSourceAddress, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.RouteAccessListsTypeIpv4Ipv4EntryInnerSourceAddress{}.AttrTypes(), &item)
		unpacked, d := unpackRouteAccessListsTypeIpv4Ipv4EntryInnerSourceAddressToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.RouteAccessListsTypeIpv4Ipv4EntryInnerSourceAddress", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for RouteAccessListsTypeIpv4Ipv4EntryInnerSourceAddress ---
func packRouteAccessListsTypeIpv4Ipv4EntryInnerSourceAddressListFromSdk(ctx context.Context, sdks []network_services.RouteAccessListsTypeIpv4Ipv4EntryInnerSourceAddress) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.RouteAccessListsTypeIpv4Ipv4EntryInnerSourceAddress")
	diags := diag.Diagnostics{}
	var data []models.RouteAccessListsTypeIpv4Ipv4EntryInnerSourceAddress

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.RouteAccessListsTypeIpv4Ipv4EntryInnerSourceAddress
		obj, d := packRouteAccessListsTypeIpv4Ipv4EntryInnerSourceAddressFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.RouteAccessListsTypeIpv4Ipv4EntryInnerSourceAddress{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.RouteAccessListsTypeIpv4Ipv4EntryInnerSourceAddress", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.RouteAccessListsTypeIpv4Ipv4EntryInnerSourceAddress{}.AttrType(), data)
}
