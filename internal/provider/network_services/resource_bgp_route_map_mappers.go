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

// --- Unpacker for BgpRouteMaps ---
func unpackBgpRouteMapsToSdk(ctx context.Context, obj types.Object) (*network_services.BgpRouteMaps, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpRouteMaps", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMaps
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpRouteMaps
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

	// Handling Lists
	if !model.RouteMap.IsNull() && !model.RouteMap.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field RouteMap")
		unpacked, d := unpackBgpRouteMapsRouteMapInnerListToSdk(ctx, model.RouteMap)
		diags.Append(d...)
		sdk.RouteMap = unpacked
	}

	// Handling Primitives
	if !model.Snippet.IsNull() && !model.Snippet.IsUnknown() {
		sdk.Snippet = model.Snippet.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpRouteMaps", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpRouteMaps ---
func packBgpRouteMapsFromSdk(ctx context.Context, sdk network_services.BgpRouteMaps) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpRouteMaps", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMaps
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
	// Handling Lists
	if sdk.RouteMap != nil {
		tflog.Debug(ctx, "Packing list of objects for field RouteMap")
		packed, d := packBgpRouteMapsRouteMapInnerListFromSdk(ctx, sdk.RouteMap)
		diags.Append(d...)
		model.RouteMap = packed
	} else {
		model.RouteMap = basetypes.NewListNull(models.BgpRouteMapsRouteMapInner{}.AttrType())
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

	obj, d := types.ObjectValueFrom(ctx, models.BgpRouteMaps{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpRouteMaps", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpRouteMaps ---
func unpackBgpRouteMapsListToSdk(ctx context.Context, list types.List) ([]network_services.BgpRouteMaps, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpRouteMaps")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMaps
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpRouteMaps, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpRouteMaps{}.AttrTypes(), &item)
		unpacked, d := unpackBgpRouteMapsToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpRouteMaps", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpRouteMaps ---
func packBgpRouteMapsListFromSdk(ctx context.Context, sdks []network_services.BgpRouteMaps) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpRouteMaps")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMaps

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpRouteMaps
		obj, d := packBgpRouteMapsFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpRouteMaps{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpRouteMaps", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpRouteMaps{}.AttrType(), data)
}

// --- Unpacker for BgpRouteMapsRouteMapInner ---
func unpackBgpRouteMapsRouteMapInnerToSdk(ctx context.Context, obj types.Object) (*network_services.BgpRouteMapsRouteMapInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpRouteMapsRouteMapInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapsRouteMapInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpRouteMapsRouteMapInner
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Action.IsNull() && !model.Action.IsUnknown() {
		sdk.Action = model.Action.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Action", "value": *sdk.Action})
	}

	// Handling Primitives
	if !model.Description.IsNull() && !model.Description.IsUnknown() {
		sdk.Description = model.Description.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Description", "value": *sdk.Description})
	}

	// Handling Objects
	if !model.Match.IsNull() && !model.Match.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Match")
		unpacked, d := unpackBgpRouteMapsRouteMapInnerMatchToSdk(ctx, model.Match)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Match"})
		}
		if unpacked != nil {
			sdk.Match = unpacked
		}
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		val := int32(model.Name.ValueInt64())
		sdk.Name = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
	}

	// Handling Objects
	if !model.Set.IsNull() && !model.Set.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Set")
		unpacked, d := unpackBgpRouteMapsRouteMapInnerSetToSdk(ctx, model.Set)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Set"})
		}
		if unpacked != nil {
			sdk.Set = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpRouteMapsRouteMapInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpRouteMapsRouteMapInner ---
func packBgpRouteMapsRouteMapInnerFromSdk(ctx context.Context, sdk network_services.BgpRouteMapsRouteMapInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpRouteMapsRouteMapInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapsRouteMapInner
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
	if sdk.Description != nil {
		model.Description = basetypes.NewStringValue(*sdk.Description)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Description", "value": *sdk.Description})
	} else {
		model.Description = basetypes.NewStringNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Match != nil {
		tflog.Debug(ctx, "Packing nested object for field Match")
		packed, d := packBgpRouteMapsRouteMapInnerMatchFromSdk(ctx, *sdk.Match)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Match"})
		}
		model.Match = packed
	} else {
		model.Match = basetypes.NewObjectNull(models.BgpRouteMapsRouteMapInnerMatch{}.AttrTypes())
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
	if sdk.Set != nil {
		tflog.Debug(ctx, "Packing nested object for field Set")
		packed, d := packBgpRouteMapsRouteMapInnerSetFromSdk(ctx, *sdk.Set)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Set"})
		}
		model.Set = packed
	} else {
		model.Set = basetypes.NewObjectNull(models.BgpRouteMapsRouteMapInnerSet{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BgpRouteMapsRouteMapInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpRouteMapsRouteMapInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpRouteMapsRouteMapInner ---
func unpackBgpRouteMapsRouteMapInnerListToSdk(ctx context.Context, list types.List) ([]network_services.BgpRouteMapsRouteMapInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpRouteMapsRouteMapInner")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapsRouteMapInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpRouteMapsRouteMapInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpRouteMapsRouteMapInner{}.AttrTypes(), &item)
		unpacked, d := unpackBgpRouteMapsRouteMapInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpRouteMapsRouteMapInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpRouteMapsRouteMapInner ---
func packBgpRouteMapsRouteMapInnerListFromSdk(ctx context.Context, sdks []network_services.BgpRouteMapsRouteMapInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpRouteMapsRouteMapInner")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapsRouteMapInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpRouteMapsRouteMapInner
		obj, d := packBgpRouteMapsRouteMapInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpRouteMapsRouteMapInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpRouteMapsRouteMapInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpRouteMapsRouteMapInner{}.AttrType(), data)
}

// --- Unpacker for BgpRouteMapsRouteMapInnerMatch ---
func unpackBgpRouteMapsRouteMapInnerMatchToSdk(ctx context.Context, obj types.Object) (*network_services.BgpRouteMapsRouteMapInnerMatch, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpRouteMapsRouteMapInnerMatch", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapsRouteMapInnerMatch
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpRouteMapsRouteMapInnerMatch
	var d diag.Diagnostics
	// Handling Primitives
	if !model.AsPathAccessList.IsNull() && !model.AsPathAccessList.IsUnknown() {
		sdk.AsPathAccessList = model.AsPathAccessList.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "AsPathAccessList", "value": *sdk.AsPathAccessList})
	}

	// Handling Primitives
	if !model.ExtendedCommunity.IsNull() && !model.ExtendedCommunity.IsUnknown() {
		sdk.ExtendedCommunity = model.ExtendedCommunity.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "ExtendedCommunity", "value": *sdk.ExtendedCommunity})
	}

	// Handling Primitives
	if !model.Interface.IsNull() && !model.Interface.IsUnknown() {
		sdk.Interface = model.Interface.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Interface", "value": *sdk.Interface})
	}

	// Handling Objects
	if !model.Ipv4.IsNull() && !model.Ipv4.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Ipv4")
		unpacked, d := unpackBgpRouteMapsRouteMapInnerMatchIpv4ToSdk(ctx, model.Ipv4)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Ipv4"})
		}
		if unpacked != nil {
			sdk.Ipv4 = unpacked
		}
	}

	// Handling Primitives
	if !model.LargeCommunity.IsNull() && !model.LargeCommunity.IsUnknown() {
		sdk.LargeCommunity = model.LargeCommunity.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "LargeCommunity", "value": *sdk.LargeCommunity})
	}

	// Handling Primitives
	if !model.LocalPreference.IsNull() && !model.LocalPreference.IsUnknown() {
		val := int32(model.LocalPreference.ValueInt64())
		sdk.LocalPreference = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "LocalPreference", "value": *sdk.LocalPreference})
	}

	// Handling Primitives
	if !model.Metric.IsNull() && !model.Metric.IsUnknown() {
		val := int32(model.Metric.ValueInt64())
		sdk.Metric = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Metric", "value": *sdk.Metric})
	}

	// Handling Primitives
	if !model.Origin.IsNull() && !model.Origin.IsUnknown() {
		sdk.Origin = model.Origin.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Origin", "value": *sdk.Origin})
	}

	// Handling Primitives
	if !model.Peer.IsNull() && !model.Peer.IsUnknown() {
		sdk.Peer = model.Peer.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Peer", "value": *sdk.Peer})
	}

	// Handling Primitives
	if !model.RegularCommunity.IsNull() && !model.RegularCommunity.IsUnknown() {
		sdk.RegularCommunity = model.RegularCommunity.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "RegularCommunity", "value": *sdk.RegularCommunity})
	}

	// Handling Primitives
	if !model.Tag.IsNull() && !model.Tag.IsUnknown() {
		val := int32(model.Tag.ValueInt64())
		sdk.Tag = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Tag", "value": *sdk.Tag})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpRouteMapsRouteMapInnerMatch", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpRouteMapsRouteMapInnerMatch ---
func packBgpRouteMapsRouteMapInnerMatchFromSdk(ctx context.Context, sdk network_services.BgpRouteMapsRouteMapInnerMatch) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpRouteMapsRouteMapInnerMatch", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapsRouteMapInnerMatch
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.AsPathAccessList != nil {
		model.AsPathAccessList = basetypes.NewStringValue(*sdk.AsPathAccessList)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "AsPathAccessList", "value": *sdk.AsPathAccessList})
	} else {
		model.AsPathAccessList = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.ExtendedCommunity != nil {
		model.ExtendedCommunity = basetypes.NewStringValue(*sdk.ExtendedCommunity)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "ExtendedCommunity", "value": *sdk.ExtendedCommunity})
	} else {
		model.ExtendedCommunity = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Interface != nil {
		model.Interface = basetypes.NewStringValue(*sdk.Interface)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Interface", "value": *sdk.Interface})
	} else {
		model.Interface = basetypes.NewStringNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Ipv4 != nil {
		tflog.Debug(ctx, "Packing nested object for field Ipv4")
		packed, d := packBgpRouteMapsRouteMapInnerMatchIpv4FromSdk(ctx, *sdk.Ipv4)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Ipv4"})
		}
		model.Ipv4 = packed
	} else {
		model.Ipv4 = basetypes.NewObjectNull(models.BgpRouteMapsRouteMapInnerMatchIpv4{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.LargeCommunity != nil {
		model.LargeCommunity = basetypes.NewStringValue(*sdk.LargeCommunity)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "LargeCommunity", "value": *sdk.LargeCommunity})
	} else {
		model.LargeCommunity = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.LocalPreference != nil {
		model.LocalPreference = basetypes.NewInt64Value(int64(*sdk.LocalPreference))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "LocalPreference", "value": *sdk.LocalPreference})
	} else {
		model.LocalPreference = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Metric != nil {
		model.Metric = basetypes.NewInt64Value(int64(*sdk.Metric))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Metric", "value": *sdk.Metric})
	} else {
		model.Metric = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Origin != nil {
		model.Origin = basetypes.NewStringValue(*sdk.Origin)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Origin", "value": *sdk.Origin})
	} else {
		model.Origin = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Peer != nil {
		model.Peer = basetypes.NewStringValue(*sdk.Peer)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Peer", "value": *sdk.Peer})
	} else {
		model.Peer = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.RegularCommunity != nil {
		model.RegularCommunity = basetypes.NewStringValue(*sdk.RegularCommunity)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "RegularCommunity", "value": *sdk.RegularCommunity})
	} else {
		model.RegularCommunity = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Tag != nil {
		model.Tag = basetypes.NewInt64Value(int64(*sdk.Tag))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Tag", "value": *sdk.Tag})
	} else {
		model.Tag = basetypes.NewInt64Null()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BgpRouteMapsRouteMapInnerMatch{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpRouteMapsRouteMapInnerMatch", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpRouteMapsRouteMapInnerMatch ---
func unpackBgpRouteMapsRouteMapInnerMatchListToSdk(ctx context.Context, list types.List) ([]network_services.BgpRouteMapsRouteMapInnerMatch, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpRouteMapsRouteMapInnerMatch")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapsRouteMapInnerMatch
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpRouteMapsRouteMapInnerMatch, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpRouteMapsRouteMapInnerMatch{}.AttrTypes(), &item)
		unpacked, d := unpackBgpRouteMapsRouteMapInnerMatchToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpRouteMapsRouteMapInnerMatch", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpRouteMapsRouteMapInnerMatch ---
func packBgpRouteMapsRouteMapInnerMatchListFromSdk(ctx context.Context, sdks []network_services.BgpRouteMapsRouteMapInnerMatch) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpRouteMapsRouteMapInnerMatch")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapsRouteMapInnerMatch

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpRouteMapsRouteMapInnerMatch
		obj, d := packBgpRouteMapsRouteMapInnerMatchFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpRouteMapsRouteMapInnerMatch{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpRouteMapsRouteMapInnerMatch", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpRouteMapsRouteMapInnerMatch{}.AttrType(), data)
}

// --- Unpacker for BgpRouteMapsRouteMapInnerMatchIpv4 ---
func unpackBgpRouteMapsRouteMapInnerMatchIpv4ToSdk(ctx context.Context, obj types.Object) (*network_services.BgpRouteMapsRouteMapInnerMatchIpv4, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpRouteMapsRouteMapInnerMatchIpv4", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapsRouteMapInnerMatchIpv4
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpRouteMapsRouteMapInnerMatchIpv4
	var d diag.Diagnostics
	// Handling Objects
	if !model.Address.IsNull() && !model.Address.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Address")
		unpacked, d := unpackBgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4RouteSourceToSdk(ctx, model.Address)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Address"})
		}
		if unpacked != nil {
			sdk.Address = unpacked
		}
	}

	// Handling Objects
	if !model.NextHop.IsNull() && !model.NextHop.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field NextHop")
		unpacked, d := unpackBgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4RouteSourceToSdk(ctx, model.NextHop)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "NextHop"})
		}
		if unpacked != nil {
			sdk.NextHop = unpacked
		}
	}

	// Handling Objects
	if !model.RouteSource.IsNull() && !model.RouteSource.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field RouteSource")
		unpacked, d := unpackBgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4RouteSourceToSdk(ctx, model.RouteSource)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "RouteSource"})
		}
		if unpacked != nil {
			sdk.RouteSource = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpRouteMapsRouteMapInnerMatchIpv4", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpRouteMapsRouteMapInnerMatchIpv4 ---
func packBgpRouteMapsRouteMapInnerMatchIpv4FromSdk(ctx context.Context, sdk network_services.BgpRouteMapsRouteMapInnerMatchIpv4) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpRouteMapsRouteMapInnerMatchIpv4", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapsRouteMapInnerMatchIpv4
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Address != nil {
		tflog.Debug(ctx, "Packing nested object for field Address")
		packed, d := packBgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4RouteSourceFromSdk(ctx, *sdk.Address)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Address"})
		}
		model.Address = packed
	} else {
		model.Address = basetypes.NewObjectNull(models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4RouteSource{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.NextHop != nil {
		tflog.Debug(ctx, "Packing nested object for field NextHop")
		packed, d := packBgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4RouteSourceFromSdk(ctx, *sdk.NextHop)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "NextHop"})
		}
		model.NextHop = packed
	} else {
		model.NextHop = basetypes.NewObjectNull(models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4RouteSource{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.RouteSource != nil {
		tflog.Debug(ctx, "Packing nested object for field RouteSource")
		packed, d := packBgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4RouteSourceFromSdk(ctx, *sdk.RouteSource)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "RouteSource"})
		}
		model.RouteSource = packed
	} else {
		model.RouteSource = basetypes.NewObjectNull(models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4RouteSource{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BgpRouteMapsRouteMapInnerMatchIpv4{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpRouteMapsRouteMapInnerMatchIpv4", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpRouteMapsRouteMapInnerMatchIpv4 ---
func unpackBgpRouteMapsRouteMapInnerMatchIpv4ListToSdk(ctx context.Context, list types.List) ([]network_services.BgpRouteMapsRouteMapInnerMatchIpv4, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpRouteMapsRouteMapInnerMatchIpv4")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapsRouteMapInnerMatchIpv4
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpRouteMapsRouteMapInnerMatchIpv4, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpRouteMapsRouteMapInnerMatchIpv4{}.AttrTypes(), &item)
		unpacked, d := unpackBgpRouteMapsRouteMapInnerMatchIpv4ToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpRouteMapsRouteMapInnerMatchIpv4", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpRouteMapsRouteMapInnerMatchIpv4 ---
func packBgpRouteMapsRouteMapInnerMatchIpv4ListFromSdk(ctx context.Context, sdks []network_services.BgpRouteMapsRouteMapInnerMatchIpv4) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpRouteMapsRouteMapInnerMatchIpv4")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapsRouteMapInnerMatchIpv4

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpRouteMapsRouteMapInnerMatchIpv4
		obj, d := packBgpRouteMapsRouteMapInnerMatchIpv4FromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpRouteMapsRouteMapInnerMatchIpv4{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpRouteMapsRouteMapInnerMatchIpv4", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpRouteMapsRouteMapInnerMatchIpv4{}.AttrType(), data)
}

// --- Unpacker for BgpRouteMapsRouteMapInnerSet ---
func unpackBgpRouteMapsRouteMapInnerSetToSdk(ctx context.Context, obj types.Object) (*network_services.BgpRouteMapsRouteMapInnerSet, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpRouteMapsRouteMapInnerSet", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapsRouteMapInnerSet
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpRouteMapsRouteMapInnerSet
	var d diag.Diagnostics
	// Handling Objects
	if !model.Aggregator.IsNull() && !model.Aggregator.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Aggregator")
		unpacked, d := unpackBgpRouteMapsRouteMapInnerSetAggregatorToSdk(ctx, model.Aggregator)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Aggregator"})
		}
		if unpacked != nil {
			sdk.Aggregator = unpacked
		}
	}

	// Handling Lists
	if !model.AspathExclude.IsNull() && !model.AspathExclude.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field AspathExclude")
		diags.Append(model.AspathExclude.ElementsAs(ctx, &sdk.AspathExclude, false)...)
	}

	// Handling Lists
	if !model.AspathPrepend.IsNull() && !model.AspathPrepend.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field AspathPrepend")
		diags.Append(model.AspathPrepend.ElementsAs(ctx, &sdk.AspathPrepend, false)...)
	}

	// Handling Primitives
	if !model.AtomicAggregate.IsNull() && !model.AtomicAggregate.IsUnknown() {
		sdk.AtomicAggregate = model.AtomicAggregate.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "AtomicAggregate", "value": *sdk.AtomicAggregate})
	}

	// Handling Objects
	if !model.Ipv4.IsNull() && !model.Ipv4.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Ipv4")
		unpacked, d := unpackBgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetIpv4ToSdk(ctx, model.Ipv4)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Ipv4"})
		}
		if unpacked != nil {
			sdk.Ipv4 = unpacked
		}
	}

	// Handling Lists
	if !model.LargeCommunity.IsNull() && !model.LargeCommunity.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field LargeCommunity")
		diags.Append(model.LargeCommunity.ElementsAs(ctx, &sdk.LargeCommunity, false)...)
	}

	// Handling Primitives
	if !model.LocalPreference.IsNull() && !model.LocalPreference.IsUnknown() {
		val := int32(model.LocalPreference.ValueInt64())
		sdk.LocalPreference = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "LocalPreference", "value": *sdk.LocalPreference})
	}

	// Handling Objects
	if !model.Metric.IsNull() && !model.Metric.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Metric")
		unpacked, d := unpackBgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetMetricToSdk(ctx, model.Metric)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Metric"})
		}
		if unpacked != nil {
			sdk.Metric = unpacked
		}
	}

	// Handling Primitives
	if !model.Origin.IsNull() && !model.Origin.IsUnknown() {
		sdk.Origin = model.Origin.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Origin", "value": *sdk.Origin})
	}

	// Handling Primitives
	if !model.OriginatorId.IsNull() && !model.OriginatorId.IsUnknown() {
		sdk.OriginatorId = model.OriginatorId.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "OriginatorId", "value": *sdk.OriginatorId})
	}

	// Handling Primitives
	if !model.OverwriteLargeCommunity.IsNull() && !model.OverwriteLargeCommunity.IsUnknown() {
		sdk.OverwriteLargeCommunity = model.OverwriteLargeCommunity.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "OverwriteLargeCommunity", "value": *sdk.OverwriteLargeCommunity})
	}

	// Handling Primitives
	if !model.OverwriteRegularCommunity.IsNull() && !model.OverwriteRegularCommunity.IsUnknown() {
		sdk.OverwriteRegularCommunity = model.OverwriteRegularCommunity.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "OverwriteRegularCommunity", "value": *sdk.OverwriteRegularCommunity})
	}

	// Handling Lists
	if !model.RegularCommunity.IsNull() && !model.RegularCommunity.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field RegularCommunity")
		diags.Append(model.RegularCommunity.ElementsAs(ctx, &sdk.RegularCommunity, false)...)
	}

	// Handling Primitives
	if !model.RemoveLargeCommunity.IsNull() && !model.RemoveLargeCommunity.IsUnknown() {
		sdk.RemoveLargeCommunity = model.RemoveLargeCommunity.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "RemoveLargeCommunity", "value": *sdk.RemoveLargeCommunity})
	}

	// Handling Primitives
	if !model.RemoveRegularCommunity.IsNull() && !model.RemoveRegularCommunity.IsUnknown() {
		sdk.RemoveRegularCommunity = model.RemoveRegularCommunity.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "RemoveRegularCommunity", "value": *sdk.RemoveRegularCommunity})
	}

	// Handling Primitives
	if !model.Tag.IsNull() && !model.Tag.IsUnknown() {
		val := int32(model.Tag.ValueInt64())
		sdk.Tag = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Tag", "value": *sdk.Tag})
	}

	// Handling Primitives
	if !model.Weight.IsNull() && !model.Weight.IsUnknown() {
		val := int32(model.Weight.ValueInt64())
		sdk.Weight = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Weight", "value": *sdk.Weight})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpRouteMapsRouteMapInnerSet", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpRouteMapsRouteMapInnerSet ---
func packBgpRouteMapsRouteMapInnerSetFromSdk(ctx context.Context, sdk network_services.BgpRouteMapsRouteMapInnerSet) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpRouteMapsRouteMapInnerSet", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapsRouteMapInnerSet
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Aggregator != nil {
		tflog.Debug(ctx, "Packing nested object for field Aggregator")
		packed, d := packBgpRouteMapsRouteMapInnerSetAggregatorFromSdk(ctx, *sdk.Aggregator)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Aggregator"})
		}
		model.Aggregator = packed
	} else {
		model.Aggregator = basetypes.NewObjectNull(models.BgpRouteMapsRouteMapInnerSetAggregator{}.AttrTypes())
	}
	// Handling Lists
	if sdk.AspathExclude != nil {
		tflog.Debug(ctx, "Packing list of primitives for field AspathExclude")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		elemType = basetypes.Int64Type{}
		model.AspathExclude, d = basetypes.NewListValueFrom(ctx, elemType, sdk.AspathExclude)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		elemType = basetypes.Int64Type{}
		model.AspathExclude = basetypes.NewListNull(elemType)
	}
	// Handling Lists
	if sdk.AspathPrepend != nil {
		tflog.Debug(ctx, "Packing list of primitives for field AspathPrepend")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		elemType = basetypes.Int64Type{}
		model.AspathPrepend, d = basetypes.NewListValueFrom(ctx, elemType, sdk.AspathPrepend)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		elemType = basetypes.Int64Type{}
		model.AspathPrepend = basetypes.NewListNull(elemType)
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.AtomicAggregate != nil {
		model.AtomicAggregate = basetypes.NewBoolValue(*sdk.AtomicAggregate)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "AtomicAggregate", "value": *sdk.AtomicAggregate})
	} else {
		model.AtomicAggregate = basetypes.NewBoolNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Ipv4 != nil {
		tflog.Debug(ctx, "Packing nested object for field Ipv4")
		packed, d := packBgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetIpv4FromSdk(ctx, *sdk.Ipv4)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Ipv4"})
		}
		model.Ipv4 = packed
	} else {
		model.Ipv4 = basetypes.NewObjectNull(models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetIpv4{}.AttrTypes())
	}
	// Handling Lists
	if sdk.LargeCommunity != nil {
		tflog.Debug(ctx, "Packing list of primitives for field LargeCommunity")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.LargeCommunity, d = basetypes.NewListValueFrom(ctx, elemType, sdk.LargeCommunity)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.LargeCommunity = basetypes.NewListNull(elemType)
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.LocalPreference != nil {
		model.LocalPreference = basetypes.NewInt64Value(int64(*sdk.LocalPreference))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "LocalPreference", "value": *sdk.LocalPreference})
	} else {
		model.LocalPreference = basetypes.NewInt64Null()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Metric != nil {
		tflog.Debug(ctx, "Packing nested object for field Metric")
		packed, d := packBgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetMetricFromSdk(ctx, *sdk.Metric)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Metric"})
		}
		model.Metric = packed
	} else {
		model.Metric = basetypes.NewObjectNull(models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetMetric{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Origin != nil {
		model.Origin = basetypes.NewStringValue(*sdk.Origin)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Origin", "value": *sdk.Origin})
	} else {
		model.Origin = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.OriginatorId != nil {
		model.OriginatorId = basetypes.NewStringValue(*sdk.OriginatorId)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "OriginatorId", "value": *sdk.OriginatorId})
	} else {
		model.OriginatorId = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.OverwriteLargeCommunity != nil {
		model.OverwriteLargeCommunity = basetypes.NewBoolValue(*sdk.OverwriteLargeCommunity)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "OverwriteLargeCommunity", "value": *sdk.OverwriteLargeCommunity})
	} else {
		model.OverwriteLargeCommunity = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.OverwriteRegularCommunity != nil {
		model.OverwriteRegularCommunity = basetypes.NewBoolValue(*sdk.OverwriteRegularCommunity)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "OverwriteRegularCommunity", "value": *sdk.OverwriteRegularCommunity})
	} else {
		model.OverwriteRegularCommunity = basetypes.NewBoolNull()
	}
	// Handling Lists
	if sdk.RegularCommunity != nil {
		tflog.Debug(ctx, "Packing list of primitives for field RegularCommunity")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.RegularCommunity, d = basetypes.NewListValueFrom(ctx, elemType, sdk.RegularCommunity)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.RegularCommunity = basetypes.NewListNull(elemType)
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.RemoveLargeCommunity != nil {
		model.RemoveLargeCommunity = basetypes.NewStringValue(*sdk.RemoveLargeCommunity)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "RemoveLargeCommunity", "value": *sdk.RemoveLargeCommunity})
	} else {
		model.RemoveLargeCommunity = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.RemoveRegularCommunity != nil {
		model.RemoveRegularCommunity = basetypes.NewStringValue(*sdk.RemoveRegularCommunity)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "RemoveRegularCommunity", "value": *sdk.RemoveRegularCommunity})
	} else {
		model.RemoveRegularCommunity = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Tag != nil {
		model.Tag = basetypes.NewInt64Value(int64(*sdk.Tag))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Tag", "value": *sdk.Tag})
	} else {
		model.Tag = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Weight != nil {
		model.Weight = basetypes.NewInt64Value(int64(*sdk.Weight))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Weight", "value": *sdk.Weight})
	} else {
		model.Weight = basetypes.NewInt64Null()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BgpRouteMapsRouteMapInnerSet{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpRouteMapsRouteMapInnerSet", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpRouteMapsRouteMapInnerSet ---
func unpackBgpRouteMapsRouteMapInnerSetListToSdk(ctx context.Context, list types.List) ([]network_services.BgpRouteMapsRouteMapInnerSet, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpRouteMapsRouteMapInnerSet")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapsRouteMapInnerSet
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpRouteMapsRouteMapInnerSet, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpRouteMapsRouteMapInnerSet{}.AttrTypes(), &item)
		unpacked, d := unpackBgpRouteMapsRouteMapInnerSetToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpRouteMapsRouteMapInnerSet", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpRouteMapsRouteMapInnerSet ---
func packBgpRouteMapsRouteMapInnerSetListFromSdk(ctx context.Context, sdks []network_services.BgpRouteMapsRouteMapInnerSet) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpRouteMapsRouteMapInnerSet")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapsRouteMapInnerSet

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpRouteMapsRouteMapInnerSet
		obj, d := packBgpRouteMapsRouteMapInnerSetFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpRouteMapsRouteMapInnerSet{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpRouteMapsRouteMapInnerSet", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpRouteMapsRouteMapInnerSet{}.AttrType(), data)
}

// --- Unpacker for BgpRouteMapsRouteMapInnerSetAggregator ---
func unpackBgpRouteMapsRouteMapInnerSetAggregatorToSdk(ctx context.Context, obj types.Object) (*network_services.BgpRouteMapsRouteMapInnerSetAggregator, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpRouteMapsRouteMapInnerSetAggregator", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapsRouteMapInnerSetAggregator
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpRouteMapsRouteMapInnerSetAggregator
	var d diag.Diagnostics
	// Handling Primitives
	if !model.As.IsNull() && !model.As.IsUnknown() {
		val := int32(model.As.ValueInt64())
		sdk.As = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "As", "value": *sdk.As})
	}

	// Handling Primitives
	if !model.RouterId.IsNull() && !model.RouterId.IsUnknown() {
		sdk.RouterId = model.RouterId.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "RouterId", "value": *sdk.RouterId})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpRouteMapsRouteMapInnerSetAggregator", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpRouteMapsRouteMapInnerSetAggregator ---
func packBgpRouteMapsRouteMapInnerSetAggregatorFromSdk(ctx context.Context, sdk network_services.BgpRouteMapsRouteMapInnerSetAggregator) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpRouteMapsRouteMapInnerSetAggregator", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapsRouteMapInnerSetAggregator
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.As != nil {
		model.As = basetypes.NewInt64Value(int64(*sdk.As))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "As", "value": *sdk.As})
	} else {
		model.As = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.RouterId != nil {
		model.RouterId = basetypes.NewStringValue(*sdk.RouterId)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "RouterId", "value": *sdk.RouterId})
	} else {
		model.RouterId = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BgpRouteMapsRouteMapInnerSetAggregator{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpRouteMapsRouteMapInnerSetAggregator", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpRouteMapsRouteMapInnerSetAggregator ---
func unpackBgpRouteMapsRouteMapInnerSetAggregatorListToSdk(ctx context.Context, list types.List) ([]network_services.BgpRouteMapsRouteMapInnerSetAggregator, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpRouteMapsRouteMapInnerSetAggregator")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapsRouteMapInnerSetAggregator
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpRouteMapsRouteMapInnerSetAggregator, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpRouteMapsRouteMapInnerSetAggregator{}.AttrTypes(), &item)
		unpacked, d := unpackBgpRouteMapsRouteMapInnerSetAggregatorToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpRouteMapsRouteMapInnerSetAggregator", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpRouteMapsRouteMapInnerSetAggregator ---
func packBgpRouteMapsRouteMapInnerSetAggregatorListFromSdk(ctx context.Context, sdks []network_services.BgpRouteMapsRouteMapInnerSetAggregator) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpRouteMapsRouteMapInnerSetAggregator")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapsRouteMapInnerSetAggregator

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpRouteMapsRouteMapInnerSetAggregator
		obj, d := packBgpRouteMapsRouteMapInnerSetAggregatorFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpRouteMapsRouteMapInnerSetAggregator{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpRouteMapsRouteMapInnerSetAggregator", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpRouteMapsRouteMapInnerSetAggregator{}.AttrType(), data)
}
