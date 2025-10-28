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

// --- Unpacker for BgpRouteMapRedistributions ---
func unpackBgpRouteMapRedistributionsToSdk(ctx context.Context, obj types.Object) (*network_services.BgpRouteMapRedistributions, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpRouteMapRedistributions", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributions
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpRouteMapRedistributions
	var d diag.Diagnostics

	// Handling Objects
	if !model.Bgp.IsNull() && !model.Bgp.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Bgp")
		unpacked, d := unpackBgpRouteMapRedistributionsBgpToSdk(ctx, model.Bgp)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Bgp"})
		}
		if unpacked != nil {
			sdk.Bgp = unpacked
		}
	}

	// Handling Objects
	if !model.ConnectedStatic.IsNull() && !model.ConnectedStatic.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field ConnectedStatic")
		unpacked, d := unpackBgpRouteMapRedistributionsConnectedStaticToSdk(ctx, model.ConnectedStatic)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "ConnectedStatic"})
		}
		if unpacked != nil {
			sdk.ConnectedStatic = unpacked
		}
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

	// Handling Objects
	if !model.Ospf.IsNull() && !model.Ospf.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Ospf")
		unpacked, d := unpackBgpRouteMapRedistributionsOspfToSdk(ctx, model.Ospf)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Ospf"})
		}
		if unpacked != nil {
			sdk.Ospf = unpacked
		}
	}

	// Handling Primitives
	if !model.Snippet.IsNull() && !model.Snippet.IsUnknown() {
		sdk.Snippet = model.Snippet.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpRouteMapRedistributions", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpRouteMapRedistributions ---
func packBgpRouteMapRedistributionsFromSdk(ctx context.Context, sdk network_services.BgpRouteMapRedistributions) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpRouteMapRedistributions", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributions
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Bgp != nil {
		tflog.Debug(ctx, "Packing nested object for field Bgp")
		packed, d := packBgpRouteMapRedistributionsBgpFromSdk(ctx, *sdk.Bgp)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Bgp"})
		}
		model.Bgp = packed
	} else {
		model.Bgp = basetypes.NewObjectNull(models.BgpRouteMapRedistributionsBgp{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.ConnectedStatic != nil {
		tflog.Debug(ctx, "Packing nested object for field ConnectedStatic")
		packed, d := packBgpRouteMapRedistributionsConnectedStaticFromSdk(ctx, *sdk.ConnectedStatic)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "ConnectedStatic"})
		}
		model.ConnectedStatic = packed
	} else {
		model.ConnectedStatic = basetypes.NewObjectNull(models.BgpRouteMapRedistributionsConnectedStatic{}.AttrTypes())
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
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Ospf != nil {
		tflog.Debug(ctx, "Packing nested object for field Ospf")
		packed, d := packBgpRouteMapRedistributionsOspfFromSdk(ctx, *sdk.Ospf)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Ospf"})
		}
		model.Ospf = packed
	} else {
		model.Ospf = basetypes.NewObjectNull(models.BgpRouteMapRedistributionsOspf{}.AttrTypes())
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

	obj, d := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributions{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpRouteMapRedistributions", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpRouteMapRedistributions ---
func unpackBgpRouteMapRedistributionsListToSdk(ctx context.Context, list types.List) ([]network_services.BgpRouteMapRedistributions, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpRouteMapRedistributions")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributions
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpRouteMapRedistributions, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributions{}.AttrTypes(), &item)
		unpacked, d := unpackBgpRouteMapRedistributionsToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpRouteMapRedistributions", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpRouteMapRedistributions ---
func packBgpRouteMapRedistributionsListFromSdk(ctx context.Context, sdks []network_services.BgpRouteMapRedistributions) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpRouteMapRedistributions")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributions

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpRouteMapRedistributions
		obj, d := packBgpRouteMapRedistributionsFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpRouteMapRedistributions{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpRouteMapRedistributions", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpRouteMapRedistributions{}.AttrType(), data)
}

// --- Unpacker for BgpRouteMapRedistributionsBgp ---
func unpackBgpRouteMapRedistributionsBgpToSdk(ctx context.Context, obj types.Object) (*network_services.BgpRouteMapRedistributionsBgp, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpRouteMapRedistributionsBgp", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsBgp
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpRouteMapRedistributionsBgp
	var d diag.Diagnostics
	// Handling Objects
	if !model.Ospf.IsNull() && !model.Ospf.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Ospf")
		unpacked, d := unpackBgpRouteMapRedistributionsBgpOspfToSdk(ctx, model.Ospf)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Ospf"})
		}
		if unpacked != nil {
			sdk.Ospf = unpacked
		}
	}

	// Handling Objects
	if !model.Rib.IsNull() && !model.Rib.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Rib")
		unpacked, d := unpackBgpRouteMapRedistributionsBgpRibToSdk(ctx, model.Rib)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Rib"})
		}
		if unpacked != nil {
			sdk.Rib = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpRouteMapRedistributionsBgp", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpRouteMapRedistributionsBgp ---
func packBgpRouteMapRedistributionsBgpFromSdk(ctx context.Context, sdk network_services.BgpRouteMapRedistributionsBgp) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpRouteMapRedistributionsBgp", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsBgp
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Ospf != nil {
		tflog.Debug(ctx, "Packing nested object for field Ospf")
		packed, d := packBgpRouteMapRedistributionsBgpOspfFromSdk(ctx, *sdk.Ospf)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Ospf"})
		}
		model.Ospf = packed
	} else {
		model.Ospf = basetypes.NewObjectNull(models.BgpRouteMapRedistributionsBgpOspf{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Rib != nil {
		tflog.Debug(ctx, "Packing nested object for field Rib")
		packed, d := packBgpRouteMapRedistributionsBgpRibFromSdk(ctx, *sdk.Rib)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Rib"})
		}
		model.Rib = packed
	} else {
		model.Rib = basetypes.NewObjectNull(models.BgpRouteMapRedistributionsBgpRib{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsBgp{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpRouteMapRedistributionsBgp", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpRouteMapRedistributionsBgp ---
func unpackBgpRouteMapRedistributionsBgpListToSdk(ctx context.Context, list types.List) ([]network_services.BgpRouteMapRedistributionsBgp, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpRouteMapRedistributionsBgp")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsBgp
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpRouteMapRedistributionsBgp, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsBgp{}.AttrTypes(), &item)
		unpacked, d := unpackBgpRouteMapRedistributionsBgpToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpRouteMapRedistributionsBgp", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpRouteMapRedistributionsBgp ---
func packBgpRouteMapRedistributionsBgpListFromSdk(ctx context.Context, sdks []network_services.BgpRouteMapRedistributionsBgp) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpRouteMapRedistributionsBgp")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsBgp

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpRouteMapRedistributionsBgp
		obj, d := packBgpRouteMapRedistributionsBgpFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpRouteMapRedistributionsBgp{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpRouteMapRedistributionsBgp", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpRouteMapRedistributionsBgp{}.AttrType(), data)
}

// --- Unpacker for BgpRouteMapRedistributionsBgpOspf ---
func unpackBgpRouteMapRedistributionsBgpOspfToSdk(ctx context.Context, obj types.Object) (*network_services.BgpRouteMapRedistributionsBgpOspf, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpRouteMapRedistributionsBgpOspf", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsBgpOspf
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpRouteMapRedistributionsBgpOspf
	var d diag.Diagnostics
	// Handling Lists
	if !model.RouteMap.IsNull() && !model.RouteMap.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field RouteMap")
		unpacked, d := unpackBgpRouteMapRedistributionsBgpOspfRouteMapInnerListToSdk(ctx, model.RouteMap)
		diags.Append(d...)
		sdk.RouteMap = unpacked
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpRouteMapRedistributionsBgpOspf", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpRouteMapRedistributionsBgpOspf ---
func packBgpRouteMapRedistributionsBgpOspfFromSdk(ctx context.Context, sdk network_services.BgpRouteMapRedistributionsBgpOspf) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpRouteMapRedistributionsBgpOspf", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsBgpOspf
	var d diag.Diagnostics
	// Handling Lists
	if sdk.RouteMap != nil {
		tflog.Debug(ctx, "Packing list of objects for field RouteMap")
		packed, d := packBgpRouteMapRedistributionsBgpOspfRouteMapInnerListFromSdk(ctx, sdk.RouteMap)
		diags.Append(d...)
		model.RouteMap = packed
	} else {
		model.RouteMap = basetypes.NewListNull(models.BgpRouteMapRedistributionsBgpOspfRouteMapInner{}.AttrType())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsBgpOspf{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpRouteMapRedistributionsBgpOspf", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpRouteMapRedistributionsBgpOspf ---
func unpackBgpRouteMapRedistributionsBgpOspfListToSdk(ctx context.Context, list types.List) ([]network_services.BgpRouteMapRedistributionsBgpOspf, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpRouteMapRedistributionsBgpOspf")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsBgpOspf
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpRouteMapRedistributionsBgpOspf, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsBgpOspf{}.AttrTypes(), &item)
		unpacked, d := unpackBgpRouteMapRedistributionsBgpOspfToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpRouteMapRedistributionsBgpOspf", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpRouteMapRedistributionsBgpOspf ---
func packBgpRouteMapRedistributionsBgpOspfListFromSdk(ctx context.Context, sdks []network_services.BgpRouteMapRedistributionsBgpOspf) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpRouteMapRedistributionsBgpOspf")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsBgpOspf

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpRouteMapRedistributionsBgpOspf
		obj, d := packBgpRouteMapRedistributionsBgpOspfFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpRouteMapRedistributionsBgpOspf{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpRouteMapRedistributionsBgpOspf", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpRouteMapRedistributionsBgpOspf{}.AttrType(), data)
}

// --- Unpacker for BgpRouteMapRedistributionsBgpOspfRouteMapInner ---
func unpackBgpRouteMapRedistributionsBgpOspfRouteMapInnerToSdk(ctx context.Context, obj types.Object) (*network_services.BgpRouteMapRedistributionsBgpOspfRouteMapInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpRouteMapRedistributionsBgpOspfRouteMapInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsBgpOspfRouteMapInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpRouteMapRedistributionsBgpOspfRouteMapInner
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
		unpacked, d := unpackBgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchToSdk(ctx, model.Match)
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
		unpacked, d := unpackBgpRouteMapRedistributionsBgpOspfRouteMapInnerSetToSdk(ctx, model.Set)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Set"})
		}
		if unpacked != nil {
			sdk.Set = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpRouteMapRedistributionsBgpOspfRouteMapInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpRouteMapRedistributionsBgpOspfRouteMapInner ---
func packBgpRouteMapRedistributionsBgpOspfRouteMapInnerFromSdk(ctx context.Context, sdk network_services.BgpRouteMapRedistributionsBgpOspfRouteMapInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpRouteMapRedistributionsBgpOspfRouteMapInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsBgpOspfRouteMapInner
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
		packed, d := packBgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchFromSdk(ctx, *sdk.Match)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Match"})
		}
		model.Match = packed
	} else {
		model.Match = basetypes.NewObjectNull(models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatch{}.AttrTypes())
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
		packed, d := packBgpRouteMapRedistributionsBgpOspfRouteMapInnerSetFromSdk(ctx, *sdk.Set)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Set"})
		}
		model.Set = packed
	} else {
		model.Set = basetypes.NewObjectNull(models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerSet{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsBgpOspfRouteMapInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpRouteMapRedistributionsBgpOspfRouteMapInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpRouteMapRedistributionsBgpOspfRouteMapInner ---
func unpackBgpRouteMapRedistributionsBgpOspfRouteMapInnerListToSdk(ctx context.Context, list types.List) ([]network_services.BgpRouteMapRedistributionsBgpOspfRouteMapInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpRouteMapRedistributionsBgpOspfRouteMapInner")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsBgpOspfRouteMapInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpRouteMapRedistributionsBgpOspfRouteMapInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsBgpOspfRouteMapInner{}.AttrTypes(), &item)
		unpacked, d := unpackBgpRouteMapRedistributionsBgpOspfRouteMapInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpRouteMapRedistributionsBgpOspfRouteMapInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpRouteMapRedistributionsBgpOspfRouteMapInner ---
func packBgpRouteMapRedistributionsBgpOspfRouteMapInnerListFromSdk(ctx context.Context, sdks []network_services.BgpRouteMapRedistributionsBgpOspfRouteMapInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpRouteMapRedistributionsBgpOspfRouteMapInner")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsBgpOspfRouteMapInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpRouteMapRedistributionsBgpOspfRouteMapInner
		obj, d := packBgpRouteMapRedistributionsBgpOspfRouteMapInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpRouteMapRedistributionsBgpOspfRouteMapInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpRouteMapRedistributionsBgpOspfRouteMapInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpRouteMapRedistributionsBgpOspfRouteMapInner{}.AttrType(), data)
}

// --- Unpacker for BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatch ---
func unpackBgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchToSdk(ctx context.Context, obj types.Object) (*network_services.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatch, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatch", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatch
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatch
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
		unpacked, d := unpackBgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4ToSdk(ctx, model.Ipv4)
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

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatch", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatch ---
func packBgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchFromSdk(ctx context.Context, sdk network_services.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatch) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatch", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatch
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
		packed, d := packBgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4FromSdk(ctx, *sdk.Ipv4)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Ipv4"})
		}
		model.Ipv4 = packed
	} else {
		model.Ipv4 = basetypes.NewObjectNull(models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4{}.AttrTypes())
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

	obj, d := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatch{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatch", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatch ---
func unpackBgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchListToSdk(ctx context.Context, list types.List) ([]network_services.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatch, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatch")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatch
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatch, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatch{}.AttrTypes(), &item)
		unpacked, d := unpackBgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatch", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatch ---
func packBgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchListFromSdk(ctx context.Context, sdks []network_services.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatch) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatch")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatch

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatch
		obj, d := packBgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatch{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatch", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatch{}.AttrType(), data)
}

// --- Unpacker for BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4 ---
func unpackBgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4ToSdk(ctx context.Context, obj types.Object) (*network_services.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4
	var d diag.Diagnostics
	// Handling Objects
	if !model.Address.IsNull() && !model.Address.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Address")
		unpacked, d := unpackBgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4AddressToSdk(ctx, model.Address)
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
		unpacked, d := unpackBgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4NextHopToSdk(ctx, model.NextHop)
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
		unpacked, d := unpackBgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4RouteSourceToSdk(ctx, model.RouteSource)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "RouteSource"})
		}
		if unpacked != nil {
			sdk.RouteSource = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4 ---
func packBgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4FromSdk(ctx context.Context, sdk network_services.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Address != nil {
		tflog.Debug(ctx, "Packing nested object for field Address")
		packed, d := packBgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4AddressFromSdk(ctx, *sdk.Address)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Address"})
		}
		model.Address = packed
	} else {
		model.Address = basetypes.NewObjectNull(models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4Address{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.NextHop != nil {
		tflog.Debug(ctx, "Packing nested object for field NextHop")
		packed, d := packBgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4NextHopFromSdk(ctx, *sdk.NextHop)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "NextHop"})
		}
		model.NextHop = packed
	} else {
		model.NextHop = basetypes.NewObjectNull(models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4NextHop{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.RouteSource != nil {
		tflog.Debug(ctx, "Packing nested object for field RouteSource")
		packed, d := packBgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4RouteSourceFromSdk(ctx, *sdk.RouteSource)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "RouteSource"})
		}
		model.RouteSource = packed
	} else {
		model.RouteSource = basetypes.NewObjectNull(models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4RouteSource{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4 ---
func unpackBgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4ListToSdk(ctx context.Context, list types.List) ([]network_services.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4{}.AttrTypes(), &item)
		unpacked, d := unpackBgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4ToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4 ---
func packBgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4ListFromSdk(ctx context.Context, sdks []network_services.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4
		obj, d := packBgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4FromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4{}.AttrType(), data)
}

// --- Unpacker for BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4Address ---
func unpackBgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4AddressToSdk(ctx context.Context, obj types.Object) (*network_services.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4Address, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4Address", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4Address
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4Address
	var d diag.Diagnostics
	// Handling Primitives
	if !model.AccessList.IsNull() && !model.AccessList.IsUnknown() {
		sdk.AccessList = model.AccessList.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "AccessList", "value": *sdk.AccessList})
	}

	// Handling Primitives
	if !model.PrefixList.IsNull() && !model.PrefixList.IsUnknown() {
		sdk.PrefixList = model.PrefixList.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "PrefixList", "value": *sdk.PrefixList})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4Address", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4Address ---
func packBgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4AddressFromSdk(ctx context.Context, sdk network_services.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4Address) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4Address", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4Address
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.AccessList != nil {
		model.AccessList = basetypes.NewStringValue(*sdk.AccessList)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "AccessList", "value": *sdk.AccessList})
	} else {
		model.AccessList = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.PrefixList != nil {
		model.PrefixList = basetypes.NewStringValue(*sdk.PrefixList)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "PrefixList", "value": *sdk.PrefixList})
	} else {
		model.PrefixList = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4Address{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4Address", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4Address ---
func unpackBgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4AddressListToSdk(ctx context.Context, list types.List) ([]network_services.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4Address, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4Address")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4Address
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4Address, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4Address{}.AttrTypes(), &item)
		unpacked, d := unpackBgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4AddressToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4Address", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4Address ---
func packBgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4AddressListFromSdk(ctx context.Context, sdks []network_services.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4Address) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4Address")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4Address

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4Address
		obj, d := packBgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4AddressFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4Address{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4Address", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4Address{}.AttrType(), data)
}

// --- Unpacker for BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4NextHop ---
func unpackBgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4NextHopToSdk(ctx context.Context, obj types.Object) (*network_services.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4NextHop, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4NextHop", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4NextHop
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4NextHop
	var d diag.Diagnostics
	// Handling Primitives
	if !model.AccessList.IsNull() && !model.AccessList.IsUnknown() {
		sdk.AccessList = model.AccessList.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "AccessList", "value": *sdk.AccessList})
	}

	// Handling Primitives
	if !model.PrefixList.IsNull() && !model.PrefixList.IsUnknown() {
		sdk.PrefixList = model.PrefixList.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "PrefixList", "value": *sdk.PrefixList})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4NextHop", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4NextHop ---
func packBgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4NextHopFromSdk(ctx context.Context, sdk network_services.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4NextHop) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4NextHop", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4NextHop
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.AccessList != nil {
		model.AccessList = basetypes.NewStringValue(*sdk.AccessList)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "AccessList", "value": *sdk.AccessList})
	} else {
		model.AccessList = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.PrefixList != nil {
		model.PrefixList = basetypes.NewStringValue(*sdk.PrefixList)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "PrefixList", "value": *sdk.PrefixList})
	} else {
		model.PrefixList = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4NextHop{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4NextHop", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4NextHop ---
func unpackBgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4NextHopListToSdk(ctx context.Context, list types.List) ([]network_services.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4NextHop, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4NextHop")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4NextHop
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4NextHop, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4NextHop{}.AttrTypes(), &item)
		unpacked, d := unpackBgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4NextHopToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4NextHop", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4NextHop ---
func packBgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4NextHopListFromSdk(ctx context.Context, sdks []network_services.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4NextHop) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4NextHop")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4NextHop

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4NextHop
		obj, d := packBgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4NextHopFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4NextHop{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4NextHop", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4NextHop{}.AttrType(), data)
}

// --- Unpacker for BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4RouteSource ---
func unpackBgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4RouteSourceToSdk(ctx context.Context, obj types.Object) (*network_services.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4RouteSource, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4RouteSource", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4RouteSource
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4RouteSource
	var d diag.Diagnostics
	// Handling Primitives
	if !model.AccessList.IsNull() && !model.AccessList.IsUnknown() {
		sdk.AccessList = model.AccessList.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "AccessList", "value": *sdk.AccessList})
	}

	// Handling Primitives
	if !model.PrefixList.IsNull() && !model.PrefixList.IsUnknown() {
		sdk.PrefixList = model.PrefixList.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "PrefixList", "value": *sdk.PrefixList})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4RouteSource", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4RouteSource ---
func packBgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4RouteSourceFromSdk(ctx context.Context, sdk network_services.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4RouteSource) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4RouteSource", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4RouteSource
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.AccessList != nil {
		model.AccessList = basetypes.NewStringValue(*sdk.AccessList)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "AccessList", "value": *sdk.AccessList})
	} else {
		model.AccessList = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.PrefixList != nil {
		model.PrefixList = basetypes.NewStringValue(*sdk.PrefixList)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "PrefixList", "value": *sdk.PrefixList})
	} else {
		model.PrefixList = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4RouteSource{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4RouteSource", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4RouteSource ---
func unpackBgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4RouteSourceListToSdk(ctx context.Context, list types.List) ([]network_services.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4RouteSource, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4RouteSource")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4RouteSource
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4RouteSource, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4RouteSource{}.AttrTypes(), &item)
		unpacked, d := unpackBgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4RouteSourceToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4RouteSource", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4RouteSource ---
func packBgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4RouteSourceListFromSdk(ctx context.Context, sdks []network_services.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4RouteSource) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4RouteSource")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4RouteSource

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4RouteSource
		obj, d := packBgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4RouteSourceFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4RouteSource{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4RouteSource", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4RouteSource{}.AttrType(), data)
}

// --- Unpacker for BgpRouteMapRedistributionsBgpOspfRouteMapInnerSet ---
func unpackBgpRouteMapRedistributionsBgpOspfRouteMapInnerSetToSdk(ctx context.Context, obj types.Object) (*network_services.BgpRouteMapRedistributionsBgpOspfRouteMapInnerSet, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerSet", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerSet
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpRouteMapRedistributionsBgpOspfRouteMapInnerSet
	var d diag.Diagnostics
	// Handling Objects
	if !model.Metric.IsNull() && !model.Metric.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Metric")
		unpacked, d := unpackBgpRouteMapRedistributionsBgpOspfRouteMapInnerSetMetricToSdk(ctx, model.Metric)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Metric"})
		}
		if unpacked != nil {
			sdk.Metric = unpacked
		}
	}

	// Handling Primitives
	if !model.MetricType.IsNull() && !model.MetricType.IsUnknown() {
		sdk.MetricType = model.MetricType.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "MetricType", "value": *sdk.MetricType})
	}

	// Handling Primitives
	if !model.Tag.IsNull() && !model.Tag.IsUnknown() {
		val := int32(model.Tag.ValueInt64())
		sdk.Tag = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Tag", "value": *sdk.Tag})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerSet", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpRouteMapRedistributionsBgpOspfRouteMapInnerSet ---
func packBgpRouteMapRedistributionsBgpOspfRouteMapInnerSetFromSdk(ctx context.Context, sdk network_services.BgpRouteMapRedistributionsBgpOspfRouteMapInnerSet) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerSet", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerSet
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Metric != nil {
		tflog.Debug(ctx, "Packing nested object for field Metric")
		packed, d := packBgpRouteMapRedistributionsBgpOspfRouteMapInnerSetMetricFromSdk(ctx, *sdk.Metric)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Metric"})
		}
		model.Metric = packed
	} else {
		model.Metric = basetypes.NewObjectNull(models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerSetMetric{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.MetricType != nil {
		model.MetricType = basetypes.NewStringValue(*sdk.MetricType)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "MetricType", "value": *sdk.MetricType})
	} else {
		model.MetricType = basetypes.NewStringNull()
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

	obj, d := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerSet{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerSet", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpRouteMapRedistributionsBgpOspfRouteMapInnerSet ---
func unpackBgpRouteMapRedistributionsBgpOspfRouteMapInnerSetListToSdk(ctx context.Context, list types.List) ([]network_services.BgpRouteMapRedistributionsBgpOspfRouteMapInnerSet, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerSet")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerSet
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpRouteMapRedistributionsBgpOspfRouteMapInnerSet, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerSet{}.AttrTypes(), &item)
		unpacked, d := unpackBgpRouteMapRedistributionsBgpOspfRouteMapInnerSetToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerSet", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpRouteMapRedistributionsBgpOspfRouteMapInnerSet ---
func packBgpRouteMapRedistributionsBgpOspfRouteMapInnerSetListFromSdk(ctx context.Context, sdks []network_services.BgpRouteMapRedistributionsBgpOspfRouteMapInnerSet) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerSet")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerSet

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerSet
		obj, d := packBgpRouteMapRedistributionsBgpOspfRouteMapInnerSetFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerSet{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerSet", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerSet{}.AttrType(), data)
}

// --- Unpacker for BgpRouteMapRedistributionsBgpOspfRouteMapInnerSetMetric ---
func unpackBgpRouteMapRedistributionsBgpOspfRouteMapInnerSetMetricToSdk(ctx context.Context, obj types.Object) (*network_services.BgpRouteMapRedistributionsBgpOspfRouteMapInnerSetMetric, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerSetMetric", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerSetMetric
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpRouteMapRedistributionsBgpOspfRouteMapInnerSetMetric
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Action.IsNull() && !model.Action.IsUnknown() {
		sdk.Action = model.Action.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Action", "value": *sdk.Action})
	}

	// Handling Primitives
	if !model.Value.IsNull() && !model.Value.IsUnknown() {
		val := int32(model.Value.ValueInt64())
		sdk.Value = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Value", "value": *sdk.Value})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerSetMetric", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpRouteMapRedistributionsBgpOspfRouteMapInnerSetMetric ---
func packBgpRouteMapRedistributionsBgpOspfRouteMapInnerSetMetricFromSdk(ctx context.Context, sdk network_services.BgpRouteMapRedistributionsBgpOspfRouteMapInnerSetMetric) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerSetMetric", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerSetMetric
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
	if sdk.Value != nil {
		model.Value = basetypes.NewInt64Value(int64(*sdk.Value))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Value", "value": *sdk.Value})
	} else {
		model.Value = basetypes.NewInt64Null()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerSetMetric{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerSetMetric", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpRouteMapRedistributionsBgpOspfRouteMapInnerSetMetric ---
func unpackBgpRouteMapRedistributionsBgpOspfRouteMapInnerSetMetricListToSdk(ctx context.Context, list types.List) ([]network_services.BgpRouteMapRedistributionsBgpOspfRouteMapInnerSetMetric, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerSetMetric")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerSetMetric
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpRouteMapRedistributionsBgpOspfRouteMapInnerSetMetric, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerSetMetric{}.AttrTypes(), &item)
		unpacked, d := unpackBgpRouteMapRedistributionsBgpOspfRouteMapInnerSetMetricToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerSetMetric", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpRouteMapRedistributionsBgpOspfRouteMapInnerSetMetric ---
func packBgpRouteMapRedistributionsBgpOspfRouteMapInnerSetMetricListFromSdk(ctx context.Context, sdks []network_services.BgpRouteMapRedistributionsBgpOspfRouteMapInnerSetMetric) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerSetMetric")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerSetMetric

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerSetMetric
		obj, d := packBgpRouteMapRedistributionsBgpOspfRouteMapInnerSetMetricFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerSetMetric{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerSetMetric", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpRouteMapRedistributionsBgpOspfRouteMapInnerSetMetric{}.AttrType(), data)
}

// --- Unpacker for BgpRouteMapRedistributionsBgpRib ---
func unpackBgpRouteMapRedistributionsBgpRibToSdk(ctx context.Context, obj types.Object) (*network_services.BgpRouteMapRedistributionsBgpRib, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpRouteMapRedistributionsBgpRib", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsBgpRib
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpRouteMapRedistributionsBgpRib
	var d diag.Diagnostics
	// Handling Lists
	if !model.RouteMap.IsNull() && !model.RouteMap.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field RouteMap")
		unpacked, d := unpackBgpRouteMapRedistributionsBgpRibRouteMapInnerListToSdk(ctx, model.RouteMap)
		diags.Append(d...)
		sdk.RouteMap = unpacked
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpRouteMapRedistributionsBgpRib", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpRouteMapRedistributionsBgpRib ---
func packBgpRouteMapRedistributionsBgpRibFromSdk(ctx context.Context, sdk network_services.BgpRouteMapRedistributionsBgpRib) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpRouteMapRedistributionsBgpRib", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsBgpRib
	var d diag.Diagnostics
	// Handling Lists
	if sdk.RouteMap != nil {
		tflog.Debug(ctx, "Packing list of objects for field RouteMap")
		packed, d := packBgpRouteMapRedistributionsBgpRibRouteMapInnerListFromSdk(ctx, sdk.RouteMap)
		diags.Append(d...)
		model.RouteMap = packed
	} else {
		model.RouteMap = basetypes.NewListNull(models.BgpRouteMapRedistributionsBgpRibRouteMapInner{}.AttrType())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsBgpRib{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpRouteMapRedistributionsBgpRib", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpRouteMapRedistributionsBgpRib ---
func unpackBgpRouteMapRedistributionsBgpRibListToSdk(ctx context.Context, list types.List) ([]network_services.BgpRouteMapRedistributionsBgpRib, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpRouteMapRedistributionsBgpRib")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsBgpRib
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpRouteMapRedistributionsBgpRib, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsBgpRib{}.AttrTypes(), &item)
		unpacked, d := unpackBgpRouteMapRedistributionsBgpRibToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpRouteMapRedistributionsBgpRib", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpRouteMapRedistributionsBgpRib ---
func packBgpRouteMapRedistributionsBgpRibListFromSdk(ctx context.Context, sdks []network_services.BgpRouteMapRedistributionsBgpRib) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpRouteMapRedistributionsBgpRib")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsBgpRib

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpRouteMapRedistributionsBgpRib
		obj, d := packBgpRouteMapRedistributionsBgpRibFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpRouteMapRedistributionsBgpRib{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpRouteMapRedistributionsBgpRib", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpRouteMapRedistributionsBgpRib{}.AttrType(), data)
}

// --- Unpacker for BgpRouteMapRedistributionsBgpRibRouteMapInner ---
func unpackBgpRouteMapRedistributionsBgpRibRouteMapInnerToSdk(ctx context.Context, obj types.Object) (*network_services.BgpRouteMapRedistributionsBgpRibRouteMapInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpRouteMapRedistributionsBgpRibRouteMapInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsBgpRibRouteMapInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpRouteMapRedistributionsBgpRibRouteMapInner
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
		unpacked, d := unpackBgpRouteMapRedistributionsBgpRibRouteMapInnerMatchToSdk(ctx, model.Match)
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
		unpacked, d := unpackBgpRouteMapRedistributionsBgpRibRouteMapInnerSetToSdk(ctx, model.Set)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Set"})
		}
		if unpacked != nil {
			sdk.Set = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpRouteMapRedistributionsBgpRibRouteMapInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpRouteMapRedistributionsBgpRibRouteMapInner ---
func packBgpRouteMapRedistributionsBgpRibRouteMapInnerFromSdk(ctx context.Context, sdk network_services.BgpRouteMapRedistributionsBgpRibRouteMapInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpRouteMapRedistributionsBgpRibRouteMapInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsBgpRibRouteMapInner
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
		packed, d := packBgpRouteMapRedistributionsBgpRibRouteMapInnerMatchFromSdk(ctx, *sdk.Match)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Match"})
		}
		model.Match = packed
	} else {
		model.Match = basetypes.NewObjectNull(models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatch{}.AttrTypes())
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
		packed, d := packBgpRouteMapRedistributionsBgpRibRouteMapInnerSetFromSdk(ctx, *sdk.Set)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Set"})
		}
		model.Set = packed
	} else {
		model.Set = basetypes.NewObjectNull(models.BgpRouteMapRedistributionsBgpRibRouteMapInnerSet{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsBgpRibRouteMapInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpRouteMapRedistributionsBgpRibRouteMapInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpRouteMapRedistributionsBgpRibRouteMapInner ---
func unpackBgpRouteMapRedistributionsBgpRibRouteMapInnerListToSdk(ctx context.Context, list types.List) ([]network_services.BgpRouteMapRedistributionsBgpRibRouteMapInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpRouteMapRedistributionsBgpRibRouteMapInner")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsBgpRibRouteMapInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpRouteMapRedistributionsBgpRibRouteMapInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsBgpRibRouteMapInner{}.AttrTypes(), &item)
		unpacked, d := unpackBgpRouteMapRedistributionsBgpRibRouteMapInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpRouteMapRedistributionsBgpRibRouteMapInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpRouteMapRedistributionsBgpRibRouteMapInner ---
func packBgpRouteMapRedistributionsBgpRibRouteMapInnerListFromSdk(ctx context.Context, sdks []network_services.BgpRouteMapRedistributionsBgpRibRouteMapInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpRouteMapRedistributionsBgpRibRouteMapInner")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsBgpRibRouteMapInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpRouteMapRedistributionsBgpRibRouteMapInner
		obj, d := packBgpRouteMapRedistributionsBgpRibRouteMapInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpRouteMapRedistributionsBgpRibRouteMapInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpRouteMapRedistributionsBgpRibRouteMapInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpRouteMapRedistributionsBgpRibRouteMapInner{}.AttrType(), data)
}

// --- Unpacker for BgpRouteMapRedistributionsBgpRibRouteMapInnerMatch ---
func unpackBgpRouteMapRedistributionsBgpRibRouteMapInnerMatchToSdk(ctx context.Context, obj types.Object) (*network_services.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatch, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatch", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatch
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatch
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
		unpacked, d := unpackBgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4ToSdk(ctx, model.Ipv4)
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

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatch", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpRouteMapRedistributionsBgpRibRouteMapInnerMatch ---
func packBgpRouteMapRedistributionsBgpRibRouteMapInnerMatchFromSdk(ctx context.Context, sdk network_services.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatch) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatch", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatch
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
		packed, d := packBgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4FromSdk(ctx, *sdk.Ipv4)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Ipv4"})
		}
		model.Ipv4 = packed
	} else {
		model.Ipv4 = basetypes.NewObjectNull(models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4{}.AttrTypes())
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

	obj, d := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatch{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatch", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpRouteMapRedistributionsBgpRibRouteMapInnerMatch ---
func unpackBgpRouteMapRedistributionsBgpRibRouteMapInnerMatchListToSdk(ctx context.Context, list types.List) ([]network_services.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatch, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatch")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatch
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatch, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatch{}.AttrTypes(), &item)
		unpacked, d := unpackBgpRouteMapRedistributionsBgpRibRouteMapInnerMatchToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatch", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpRouteMapRedistributionsBgpRibRouteMapInnerMatch ---
func packBgpRouteMapRedistributionsBgpRibRouteMapInnerMatchListFromSdk(ctx context.Context, sdks []network_services.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatch) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatch")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatch

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatch
		obj, d := packBgpRouteMapRedistributionsBgpRibRouteMapInnerMatchFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatch{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatch", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatch{}.AttrType(), data)
}

// --- Unpacker for BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4 ---
func unpackBgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4ToSdk(ctx context.Context, obj types.Object) (*network_services.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4
	var d diag.Diagnostics
	// Handling Objects
	if !model.Address.IsNull() && !model.Address.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Address")
		unpacked, d := unpackBgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4AddressToSdk(ctx, model.Address)
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
		unpacked, d := unpackBgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4NextHopToSdk(ctx, model.NextHop)
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

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4 ---
func packBgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4FromSdk(ctx context.Context, sdk network_services.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Address != nil {
		tflog.Debug(ctx, "Packing nested object for field Address")
		packed, d := packBgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4AddressFromSdk(ctx, *sdk.Address)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Address"})
		}
		model.Address = packed
	} else {
		model.Address = basetypes.NewObjectNull(models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4Address{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.NextHop != nil {
		tflog.Debug(ctx, "Packing nested object for field NextHop")
		packed, d := packBgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4NextHopFromSdk(ctx, *sdk.NextHop)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "NextHop"})
		}
		model.NextHop = packed
	} else {
		model.NextHop = basetypes.NewObjectNull(models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4NextHop{}.AttrTypes())
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

	obj, d := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4 ---
func unpackBgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4ListToSdk(ctx context.Context, list types.List) ([]network_services.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4{}.AttrTypes(), &item)
		unpacked, d := unpackBgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4ToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4 ---
func packBgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4ListFromSdk(ctx context.Context, sdks []network_services.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4
		obj, d := packBgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4FromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4{}.AttrType(), data)
}

// --- Unpacker for BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4Address ---
func unpackBgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4AddressToSdk(ctx context.Context, obj types.Object) (*network_services.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4Address, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4Address", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4Address
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4Address
	var d diag.Diagnostics
	// Handling Primitives
	if !model.AccessList.IsNull() && !model.AccessList.IsUnknown() {
		sdk.AccessList = model.AccessList.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "AccessList", "value": *sdk.AccessList})
	}

	// Handling Primitives
	if !model.PrefixList.IsNull() && !model.PrefixList.IsUnknown() {
		sdk.PrefixList = model.PrefixList.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "PrefixList", "value": *sdk.PrefixList})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4Address", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4Address ---
func packBgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4AddressFromSdk(ctx context.Context, sdk network_services.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4Address) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4Address", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4Address
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.AccessList != nil {
		model.AccessList = basetypes.NewStringValue(*sdk.AccessList)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "AccessList", "value": *sdk.AccessList})
	} else {
		model.AccessList = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.PrefixList != nil {
		model.PrefixList = basetypes.NewStringValue(*sdk.PrefixList)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "PrefixList", "value": *sdk.PrefixList})
	} else {
		model.PrefixList = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4Address{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4Address", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4Address ---
func unpackBgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4AddressListToSdk(ctx context.Context, list types.List) ([]network_services.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4Address, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4Address")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4Address
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4Address, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4Address{}.AttrTypes(), &item)
		unpacked, d := unpackBgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4AddressToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4Address", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4Address ---
func packBgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4AddressListFromSdk(ctx context.Context, sdks []network_services.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4Address) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4Address")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4Address

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4Address
		obj, d := packBgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4AddressFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4Address{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4Address", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4Address{}.AttrType(), data)
}

// --- Unpacker for BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4NextHop ---
func unpackBgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4NextHopToSdk(ctx context.Context, obj types.Object) (*network_services.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4NextHop, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4NextHop", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4NextHop
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4NextHop
	var d diag.Diagnostics
	// Handling Primitives
	if !model.AccessList.IsNull() && !model.AccessList.IsUnknown() {
		sdk.AccessList = model.AccessList.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "AccessList", "value": *sdk.AccessList})
	}

	// Handling Primitives
	if !model.PrefixList.IsNull() && !model.PrefixList.IsUnknown() {
		sdk.PrefixList = model.PrefixList.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "PrefixList", "value": *sdk.PrefixList})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4NextHop", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4NextHop ---
func packBgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4NextHopFromSdk(ctx context.Context, sdk network_services.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4NextHop) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4NextHop", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4NextHop
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.AccessList != nil {
		model.AccessList = basetypes.NewStringValue(*sdk.AccessList)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "AccessList", "value": *sdk.AccessList})
	} else {
		model.AccessList = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.PrefixList != nil {
		model.PrefixList = basetypes.NewStringValue(*sdk.PrefixList)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "PrefixList", "value": *sdk.PrefixList})
	} else {
		model.PrefixList = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4NextHop{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4NextHop", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4NextHop ---
func unpackBgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4NextHopListToSdk(ctx context.Context, list types.List) ([]network_services.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4NextHop, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4NextHop")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4NextHop
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4NextHop, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4NextHop{}.AttrTypes(), &item)
		unpacked, d := unpackBgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4NextHopToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4NextHop", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4NextHop ---
func packBgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4NextHopListFromSdk(ctx context.Context, sdks []network_services.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4NextHop) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4NextHop")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4NextHop

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4NextHop
		obj, d := packBgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4NextHopFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4NextHop{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4NextHop", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4NextHop{}.AttrType(), data)
}

// --- Unpacker for BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4RouteSource ---
func unpackBgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4RouteSourceToSdk(ctx context.Context, obj types.Object) (*network_services.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4RouteSource, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4RouteSource", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4RouteSource
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4RouteSource
	var d diag.Diagnostics
	// Handling Primitives
	if !model.AccessList.IsNull() && !model.AccessList.IsUnknown() {
		sdk.AccessList = model.AccessList.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "AccessList", "value": *sdk.AccessList})
	}

	// Handling Primitives
	if !model.PrefixList.IsNull() && !model.PrefixList.IsUnknown() {
		sdk.PrefixList = model.PrefixList.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "PrefixList", "value": *sdk.PrefixList})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4RouteSource", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4RouteSource ---
func packBgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4RouteSourceFromSdk(ctx context.Context, sdk network_services.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4RouteSource) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4RouteSource", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4RouteSource
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.AccessList != nil {
		model.AccessList = basetypes.NewStringValue(*sdk.AccessList)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "AccessList", "value": *sdk.AccessList})
	} else {
		model.AccessList = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.PrefixList != nil {
		model.PrefixList = basetypes.NewStringValue(*sdk.PrefixList)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "PrefixList", "value": *sdk.PrefixList})
	} else {
		model.PrefixList = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4RouteSource{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4RouteSource", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4RouteSource ---
func unpackBgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4RouteSourceListToSdk(ctx context.Context, list types.List) ([]network_services.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4RouteSource, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4RouteSource")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4RouteSource
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4RouteSource, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4RouteSource{}.AttrTypes(), &item)
		unpacked, d := unpackBgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4RouteSourceToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4RouteSource", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4RouteSource ---
func packBgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4RouteSourceListFromSdk(ctx context.Context, sdks []network_services.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4RouteSource) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4RouteSource")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4RouteSource

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4RouteSource
		obj, d := packBgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4RouteSourceFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4RouteSource{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4RouteSource", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4RouteSource{}.AttrType(), data)
}

// --- Unpacker for BgpRouteMapRedistributionsBgpRibRouteMapInnerSet ---
func unpackBgpRouteMapRedistributionsBgpRibRouteMapInnerSetToSdk(ctx context.Context, obj types.Object) (*network_services.BgpRouteMapRedistributionsBgpRibRouteMapInnerSet, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpRouteMapRedistributionsBgpRibRouteMapInnerSet", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsBgpRibRouteMapInnerSet
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpRouteMapRedistributionsBgpRibRouteMapInnerSet
	var d diag.Diagnostics
	// Handling Primitives
	if !model.SourceAddress.IsNull() && !model.SourceAddress.IsUnknown() {
		sdk.SourceAddress = model.SourceAddress.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "SourceAddress", "value": *sdk.SourceAddress})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpRouteMapRedistributionsBgpRibRouteMapInnerSet", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpRouteMapRedistributionsBgpRibRouteMapInnerSet ---
func packBgpRouteMapRedistributionsBgpRibRouteMapInnerSetFromSdk(ctx context.Context, sdk network_services.BgpRouteMapRedistributionsBgpRibRouteMapInnerSet) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpRouteMapRedistributionsBgpRibRouteMapInnerSet", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsBgpRibRouteMapInnerSet
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.SourceAddress != nil {
		model.SourceAddress = basetypes.NewStringValue(*sdk.SourceAddress)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "SourceAddress", "value": *sdk.SourceAddress})
	} else {
		model.SourceAddress = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsBgpRibRouteMapInnerSet{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpRouteMapRedistributionsBgpRibRouteMapInnerSet", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpRouteMapRedistributionsBgpRibRouteMapInnerSet ---
func unpackBgpRouteMapRedistributionsBgpRibRouteMapInnerSetListToSdk(ctx context.Context, list types.List) ([]network_services.BgpRouteMapRedistributionsBgpRibRouteMapInnerSet, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpRouteMapRedistributionsBgpRibRouteMapInnerSet")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsBgpRibRouteMapInnerSet
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpRouteMapRedistributionsBgpRibRouteMapInnerSet, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsBgpRibRouteMapInnerSet{}.AttrTypes(), &item)
		unpacked, d := unpackBgpRouteMapRedistributionsBgpRibRouteMapInnerSetToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpRouteMapRedistributionsBgpRibRouteMapInnerSet", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpRouteMapRedistributionsBgpRibRouteMapInnerSet ---
func packBgpRouteMapRedistributionsBgpRibRouteMapInnerSetListFromSdk(ctx context.Context, sdks []network_services.BgpRouteMapRedistributionsBgpRibRouteMapInnerSet) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpRouteMapRedistributionsBgpRibRouteMapInnerSet")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsBgpRibRouteMapInnerSet

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpRouteMapRedistributionsBgpRibRouteMapInnerSet
		obj, d := packBgpRouteMapRedistributionsBgpRibRouteMapInnerSetFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpRouteMapRedistributionsBgpRibRouteMapInnerSet{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpRouteMapRedistributionsBgpRibRouteMapInnerSet", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpRouteMapRedistributionsBgpRibRouteMapInnerSet{}.AttrType(), data)
}

// --- Unpacker for BgpRouteMapRedistributionsConnectedStatic ---
func unpackBgpRouteMapRedistributionsConnectedStaticToSdk(ctx context.Context, obj types.Object) (*network_services.BgpRouteMapRedistributionsConnectedStatic, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpRouteMapRedistributionsConnectedStatic", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsConnectedStatic
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpRouteMapRedistributionsConnectedStatic
	var d diag.Diagnostics
	// Handling Objects
	if !model.Bgp.IsNull() && !model.Bgp.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Bgp")
		unpacked, d := unpackBgpRouteMapRedistributionsConnectedStaticBgpToSdk(ctx, model.Bgp)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Bgp"})
		}
		if unpacked != nil {
			sdk.Bgp = unpacked
		}
	}

	// Handling Objects
	if !model.Ospf.IsNull() && !model.Ospf.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Ospf")
		unpacked, d := unpackBgpRouteMapRedistributionsConnectedStaticOspfToSdk(ctx, model.Ospf)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Ospf"})
		}
		if unpacked != nil {
			sdk.Ospf = unpacked
		}
	}

	// Handling Objects
	if !model.Rib.IsNull() && !model.Rib.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Rib")
		unpacked, d := unpackBgpRouteMapRedistributionsConnectedStaticRibToSdk(ctx, model.Rib)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Rib"})
		}
		if unpacked != nil {
			sdk.Rib = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpRouteMapRedistributionsConnectedStatic", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpRouteMapRedistributionsConnectedStatic ---
func packBgpRouteMapRedistributionsConnectedStaticFromSdk(ctx context.Context, sdk network_services.BgpRouteMapRedistributionsConnectedStatic) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpRouteMapRedistributionsConnectedStatic", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsConnectedStatic
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Bgp != nil {
		tflog.Debug(ctx, "Packing nested object for field Bgp")
		packed, d := packBgpRouteMapRedistributionsConnectedStaticBgpFromSdk(ctx, *sdk.Bgp)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Bgp"})
		}
		model.Bgp = packed
	} else {
		model.Bgp = basetypes.NewObjectNull(models.BgpRouteMapRedistributionsConnectedStaticBgp{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Ospf != nil {
		tflog.Debug(ctx, "Packing nested object for field Ospf")
		packed, d := packBgpRouteMapRedistributionsConnectedStaticOspfFromSdk(ctx, *sdk.Ospf)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Ospf"})
		}
		model.Ospf = packed
	} else {
		model.Ospf = basetypes.NewObjectNull(models.BgpRouteMapRedistributionsConnectedStaticOspf{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Rib != nil {
		tflog.Debug(ctx, "Packing nested object for field Rib")
		packed, d := packBgpRouteMapRedistributionsConnectedStaticRibFromSdk(ctx, *sdk.Rib)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Rib"})
		}
		model.Rib = packed
	} else {
		model.Rib = basetypes.NewObjectNull(models.BgpRouteMapRedistributionsConnectedStaticRib{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsConnectedStatic{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpRouteMapRedistributionsConnectedStatic", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpRouteMapRedistributionsConnectedStatic ---
func unpackBgpRouteMapRedistributionsConnectedStaticListToSdk(ctx context.Context, list types.List) ([]network_services.BgpRouteMapRedistributionsConnectedStatic, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpRouteMapRedistributionsConnectedStatic")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsConnectedStatic
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpRouteMapRedistributionsConnectedStatic, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsConnectedStatic{}.AttrTypes(), &item)
		unpacked, d := unpackBgpRouteMapRedistributionsConnectedStaticToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpRouteMapRedistributionsConnectedStatic", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpRouteMapRedistributionsConnectedStatic ---
func packBgpRouteMapRedistributionsConnectedStaticListFromSdk(ctx context.Context, sdks []network_services.BgpRouteMapRedistributionsConnectedStatic) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpRouteMapRedistributionsConnectedStatic")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsConnectedStatic

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpRouteMapRedistributionsConnectedStatic
		obj, d := packBgpRouteMapRedistributionsConnectedStaticFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpRouteMapRedistributionsConnectedStatic{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpRouteMapRedistributionsConnectedStatic", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpRouteMapRedistributionsConnectedStatic{}.AttrType(), data)
}

// --- Unpacker for BgpRouteMapRedistributionsConnectedStaticBgp ---
func unpackBgpRouteMapRedistributionsConnectedStaticBgpToSdk(ctx context.Context, obj types.Object) (*network_services.BgpRouteMapRedistributionsConnectedStaticBgp, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpRouteMapRedistributionsConnectedStaticBgp", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsConnectedStaticBgp
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpRouteMapRedistributionsConnectedStaticBgp
	var d diag.Diagnostics
	// Handling Lists
	if !model.RouteMap.IsNull() && !model.RouteMap.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field RouteMap")
		unpacked, d := unpackBgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerListToSdk(ctx, model.RouteMap)
		diags.Append(d...)
		sdk.RouteMap = unpacked
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpRouteMapRedistributionsConnectedStaticBgp", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpRouteMapRedistributionsConnectedStaticBgp ---
func packBgpRouteMapRedistributionsConnectedStaticBgpFromSdk(ctx context.Context, sdk network_services.BgpRouteMapRedistributionsConnectedStaticBgp) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpRouteMapRedistributionsConnectedStaticBgp", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsConnectedStaticBgp
	var d diag.Diagnostics
	// Handling Lists
	if sdk.RouteMap != nil {
		tflog.Debug(ctx, "Packing list of objects for field RouteMap")
		packed, d := packBgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerListFromSdk(ctx, sdk.RouteMap)
		diags.Append(d...)
		model.RouteMap = packed
	} else {
		model.RouteMap = basetypes.NewListNull(models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInner{}.AttrType())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsConnectedStaticBgp{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpRouteMapRedistributionsConnectedStaticBgp", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpRouteMapRedistributionsConnectedStaticBgp ---
func unpackBgpRouteMapRedistributionsConnectedStaticBgpListToSdk(ctx context.Context, list types.List) ([]network_services.BgpRouteMapRedistributionsConnectedStaticBgp, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpRouteMapRedistributionsConnectedStaticBgp")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsConnectedStaticBgp
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpRouteMapRedistributionsConnectedStaticBgp, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsConnectedStaticBgp{}.AttrTypes(), &item)
		unpacked, d := unpackBgpRouteMapRedistributionsConnectedStaticBgpToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpRouteMapRedistributionsConnectedStaticBgp", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpRouteMapRedistributionsConnectedStaticBgp ---
func packBgpRouteMapRedistributionsConnectedStaticBgpListFromSdk(ctx context.Context, sdks []network_services.BgpRouteMapRedistributionsConnectedStaticBgp) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpRouteMapRedistributionsConnectedStaticBgp")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsConnectedStaticBgp

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpRouteMapRedistributionsConnectedStaticBgp
		obj, d := packBgpRouteMapRedistributionsConnectedStaticBgpFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpRouteMapRedistributionsConnectedStaticBgp{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpRouteMapRedistributionsConnectedStaticBgp", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpRouteMapRedistributionsConnectedStaticBgp{}.AttrType(), data)
}

// --- Unpacker for BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInner ---
func unpackBgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerToSdk(ctx context.Context, obj types.Object) (*network_services.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInner
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
		unpacked, d := unpackBgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatchToSdk(ctx, model.Match)
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
		unpacked, d := unpackBgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetToSdk(ctx, model.Set)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Set"})
		}
		if unpacked != nil {
			sdk.Set = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInner ---
func packBgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerFromSdk(ctx context.Context, sdk network_services.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInner
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
		packed, d := packBgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatchFromSdk(ctx, *sdk.Match)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Match"})
		}
		model.Match = packed
	} else {
		model.Match = basetypes.NewObjectNull(models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatch{}.AttrTypes())
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
		packed, d := packBgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetFromSdk(ctx, *sdk.Set)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Set"})
		}
		model.Set = packed
	} else {
		model.Set = basetypes.NewObjectNull(models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSet{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInner ---
func unpackBgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerListToSdk(ctx context.Context, list types.List) ([]network_services.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInner")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInner{}.AttrTypes(), &item)
		unpacked, d := unpackBgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInner ---
func packBgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerListFromSdk(ctx context.Context, sdks []network_services.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInner")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInner
		obj, d := packBgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInner{}.AttrType(), data)
}

// --- Unpacker for BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatch ---
func unpackBgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatchToSdk(ctx context.Context, obj types.Object) (*network_services.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatch, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatch", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatch
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatch
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Interface.IsNull() && !model.Interface.IsUnknown() {
		sdk.Interface = model.Interface.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Interface", "value": *sdk.Interface})
	}

	// Handling Objects
	if !model.Ipv4.IsNull() && !model.Ipv4.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Ipv4")
		unpacked, d := unpackBgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatchIpv4ToSdk(ctx, model.Ipv4)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Ipv4"})
		}
		if unpacked != nil {
			sdk.Ipv4 = unpacked
		}
	}

	// Handling Primitives
	if !model.Metric.IsNull() && !model.Metric.IsUnknown() {
		val := int32(model.Metric.ValueInt64())
		sdk.Metric = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Metric", "value": *sdk.Metric})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatch", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatch ---
func packBgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatchFromSdk(ctx context.Context, sdk network_services.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatch) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatch", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatch
	var d diag.Diagnostics
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
		packed, d := packBgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatchIpv4FromSdk(ctx, *sdk.Ipv4)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Ipv4"})
		}
		model.Ipv4 = packed
	} else {
		model.Ipv4 = basetypes.NewObjectNull(models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatchIpv4{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Metric != nil {
		model.Metric = basetypes.NewInt64Value(int64(*sdk.Metric))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Metric", "value": *sdk.Metric})
	} else {
		model.Metric = basetypes.NewInt64Null()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatch{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatch", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatch ---
func unpackBgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatchListToSdk(ctx context.Context, list types.List) ([]network_services.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatch, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatch")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatch
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatch, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatch{}.AttrTypes(), &item)
		unpacked, d := unpackBgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatchToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatch", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatch ---
func packBgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatchListFromSdk(ctx context.Context, sdks []network_services.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatch) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatch")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatch

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatch
		obj, d := packBgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatchFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatch{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatch", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatch{}.AttrType(), data)
}

// --- Unpacker for BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatchIpv4 ---
func unpackBgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatchIpv4ToSdk(ctx context.Context, obj types.Object) (*network_services.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatchIpv4, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatchIpv4", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatchIpv4
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatchIpv4
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

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatchIpv4", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatchIpv4 ---
func packBgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatchIpv4FromSdk(ctx context.Context, sdk network_services.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatchIpv4) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatchIpv4", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatchIpv4
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
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatchIpv4{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatchIpv4", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatchIpv4 ---
func unpackBgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatchIpv4ListToSdk(ctx context.Context, list types.List) ([]network_services.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatchIpv4, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatchIpv4")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatchIpv4
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatchIpv4, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatchIpv4{}.AttrTypes(), &item)
		unpacked, d := unpackBgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatchIpv4ToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatchIpv4", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatchIpv4 ---
func packBgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatchIpv4ListFromSdk(ctx context.Context, sdks []network_services.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatchIpv4) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatchIpv4")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatchIpv4

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatchIpv4
		obj, d := packBgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatchIpv4FromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatchIpv4{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatchIpv4", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatchIpv4{}.AttrType(), data)
}

// --- Unpacker for BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSet ---
func unpackBgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetToSdk(ctx context.Context, obj types.Object) (*network_services.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSet, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSet", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSet
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSet
	var d diag.Diagnostics
	// Handling Objects
	if !model.Aggregator.IsNull() && !model.Aggregator.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Aggregator")
		unpacked, d := unpackBgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetAggregatorToSdk(ctx, model.Aggregator)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Aggregator"})
		}
		if unpacked != nil {
			sdk.Aggregator = unpacked
		}
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

	// Handling Lists
	if !model.RegularCommunity.IsNull() && !model.RegularCommunity.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field RegularCommunity")
		diags.Append(model.RegularCommunity.ElementsAs(ctx, &sdk.RegularCommunity, false)...)
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

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSet", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSet ---
func packBgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetFromSdk(ctx context.Context, sdk network_services.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSet) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSet", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSet
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Aggregator != nil {
		tflog.Debug(ctx, "Packing nested object for field Aggregator")
		packed, d := packBgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetAggregatorFromSdk(ctx, *sdk.Aggregator)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Aggregator"})
		}
		model.Aggregator = packed
	} else {
		model.Aggregator = basetypes.NewObjectNull(models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetAggregator{}.AttrTypes())
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

	obj, d := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSet{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSet", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSet ---
func unpackBgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetListToSdk(ctx context.Context, list types.List) ([]network_services.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSet, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSet")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSet
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSet, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSet{}.AttrTypes(), &item)
		unpacked, d := unpackBgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSet", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSet ---
func packBgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetListFromSdk(ctx context.Context, sdks []network_services.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSet) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSet")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSet

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSet
		obj, d := packBgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSet{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSet", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSet{}.AttrType(), data)
}

// --- Unpacker for BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetAggregator ---
func unpackBgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetAggregatorToSdk(ctx context.Context, obj types.Object) (*network_services.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetAggregator, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetAggregator", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetAggregator
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetAggregator
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

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetAggregator", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetAggregator ---
func packBgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetAggregatorFromSdk(ctx context.Context, sdk network_services.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetAggregator) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetAggregator", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetAggregator
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

	obj, d := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetAggregator{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetAggregator", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetAggregator ---
func unpackBgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetAggregatorListToSdk(ctx context.Context, list types.List) ([]network_services.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetAggregator, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetAggregator")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetAggregator
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetAggregator, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetAggregator{}.AttrTypes(), &item)
		unpacked, d := unpackBgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetAggregatorToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetAggregator", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetAggregator ---
func packBgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetAggregatorListFromSdk(ctx context.Context, sdks []network_services.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetAggregator) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetAggregator")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetAggregator

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetAggregator
		obj, d := packBgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetAggregatorFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetAggregator{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetAggregator", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetAggregator{}.AttrType(), data)
}

// --- Unpacker for BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetIpv4 ---
func unpackBgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetIpv4ToSdk(ctx context.Context, obj types.Object) (*network_services.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetIpv4, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetIpv4", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetIpv4
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetIpv4
	var d diag.Diagnostics
	// Handling Primitives
	if !model.NextHop.IsNull() && !model.NextHop.IsUnknown() {
		sdk.NextHop = model.NextHop.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "NextHop", "value": *sdk.NextHop})
	}

	// Handling Primitives
	if !model.SourceAddress.IsNull() && !model.SourceAddress.IsUnknown() {
		sdk.SourceAddress = model.SourceAddress.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "SourceAddress", "value": *sdk.SourceAddress})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetIpv4", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetIpv4 ---
func packBgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetIpv4FromSdk(ctx context.Context, sdk network_services.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetIpv4) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetIpv4", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetIpv4
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.NextHop != nil {
		model.NextHop = basetypes.NewStringValue(*sdk.NextHop)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "NextHop", "value": *sdk.NextHop})
	} else {
		model.NextHop = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.SourceAddress != nil {
		model.SourceAddress = basetypes.NewStringValue(*sdk.SourceAddress)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "SourceAddress", "value": *sdk.SourceAddress})
	} else {
		model.SourceAddress = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetIpv4{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetIpv4", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetIpv4 ---
func unpackBgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetIpv4ListToSdk(ctx context.Context, list types.List) ([]network_services.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetIpv4, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetIpv4")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetIpv4
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetIpv4, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetIpv4{}.AttrTypes(), &item)
		unpacked, d := unpackBgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetIpv4ToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetIpv4", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetIpv4 ---
func packBgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetIpv4ListFromSdk(ctx context.Context, sdks []network_services.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetIpv4) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetIpv4")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetIpv4

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetIpv4
		obj, d := packBgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetIpv4FromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetIpv4{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetIpv4", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetIpv4{}.AttrType(), data)
}

// --- Unpacker for BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetMetric ---
func unpackBgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetMetricToSdk(ctx context.Context, obj types.Object) (*network_services.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetMetric, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetMetric", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetMetric
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetMetric
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Action.IsNull() && !model.Action.IsUnknown() {
		sdk.Action = model.Action.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Action", "value": *sdk.Action})
	}

	// Handling Primitives
	if !model.Value.IsNull() && !model.Value.IsUnknown() {
		val := int32(model.Value.ValueInt64())
		sdk.Value = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Value", "value": *sdk.Value})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetMetric", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetMetric ---
func packBgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetMetricFromSdk(ctx context.Context, sdk network_services.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetMetric) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetMetric", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetMetric
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
	if sdk.Value != nil {
		model.Value = basetypes.NewInt64Value(int64(*sdk.Value))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Value", "value": *sdk.Value})
	} else {
		model.Value = basetypes.NewInt64Null()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetMetric{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetMetric", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetMetric ---
func unpackBgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetMetricListToSdk(ctx context.Context, list types.List) ([]network_services.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetMetric, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetMetric")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetMetric
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetMetric, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetMetric{}.AttrTypes(), &item)
		unpacked, d := unpackBgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetMetricToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetMetric", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetMetric ---
func packBgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetMetricListFromSdk(ctx context.Context, sdks []network_services.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetMetric) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetMetric")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetMetric

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetMetric
		obj, d := packBgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetMetricFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetMetric{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetMetric", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetMetric{}.AttrType(), data)
}

// --- Unpacker for BgpRouteMapRedistributionsConnectedStaticOspf ---
func unpackBgpRouteMapRedistributionsConnectedStaticOspfToSdk(ctx context.Context, obj types.Object) (*network_services.BgpRouteMapRedistributionsConnectedStaticOspf, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpRouteMapRedistributionsConnectedStaticOspf", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsConnectedStaticOspf
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpRouteMapRedistributionsConnectedStaticOspf
	var d diag.Diagnostics
	// Handling Lists
	if !model.RouteMap.IsNull() && !model.RouteMap.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field RouteMap")
		unpacked, d := unpackBgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerListToSdk(ctx, model.RouteMap)
		diags.Append(d...)
		sdk.RouteMap = unpacked
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpRouteMapRedistributionsConnectedStaticOspf", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpRouteMapRedistributionsConnectedStaticOspf ---
func packBgpRouteMapRedistributionsConnectedStaticOspfFromSdk(ctx context.Context, sdk network_services.BgpRouteMapRedistributionsConnectedStaticOspf) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpRouteMapRedistributionsConnectedStaticOspf", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsConnectedStaticOspf
	var d diag.Diagnostics
	// Handling Lists
	if sdk.RouteMap != nil {
		tflog.Debug(ctx, "Packing list of objects for field RouteMap")
		packed, d := packBgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerListFromSdk(ctx, sdk.RouteMap)
		diags.Append(d...)
		model.RouteMap = packed
	} else {
		model.RouteMap = basetypes.NewListNull(models.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInner{}.AttrType())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsConnectedStaticOspf{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpRouteMapRedistributionsConnectedStaticOspf", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpRouteMapRedistributionsConnectedStaticOspf ---
func unpackBgpRouteMapRedistributionsConnectedStaticOspfListToSdk(ctx context.Context, list types.List) ([]network_services.BgpRouteMapRedistributionsConnectedStaticOspf, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpRouteMapRedistributionsConnectedStaticOspf")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsConnectedStaticOspf
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpRouteMapRedistributionsConnectedStaticOspf, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsConnectedStaticOspf{}.AttrTypes(), &item)
		unpacked, d := unpackBgpRouteMapRedistributionsConnectedStaticOspfToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpRouteMapRedistributionsConnectedStaticOspf", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpRouteMapRedistributionsConnectedStaticOspf ---
func packBgpRouteMapRedistributionsConnectedStaticOspfListFromSdk(ctx context.Context, sdks []network_services.BgpRouteMapRedistributionsConnectedStaticOspf) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpRouteMapRedistributionsConnectedStaticOspf")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsConnectedStaticOspf

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpRouteMapRedistributionsConnectedStaticOspf
		obj, d := packBgpRouteMapRedistributionsConnectedStaticOspfFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpRouteMapRedistributionsConnectedStaticOspf{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpRouteMapRedistributionsConnectedStaticOspf", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpRouteMapRedistributionsConnectedStaticOspf{}.AttrType(), data)
}

// --- Unpacker for BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInner ---
func unpackBgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerToSdk(ctx context.Context, obj types.Object) (*network_services.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInner
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
		unpacked, d := unpackBgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatchToSdk(ctx, model.Match)
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
		unpacked, d := unpackBgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerSetToSdk(ctx, model.Set)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Set"})
		}
		if unpacked != nil {
			sdk.Set = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInner ---
func packBgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerFromSdk(ctx context.Context, sdk network_services.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInner
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
		packed, d := packBgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatchFromSdk(ctx, *sdk.Match)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Match"})
		}
		model.Match = packed
	} else {
		model.Match = basetypes.NewObjectNull(models.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatch{}.AttrTypes())
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
		packed, d := packBgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerSetFromSdk(ctx, *sdk.Set)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Set"})
		}
		model.Set = packed
	} else {
		model.Set = basetypes.NewObjectNull(models.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerSet{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInner ---
func unpackBgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerListToSdk(ctx context.Context, list types.List) ([]network_services.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInner")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInner{}.AttrTypes(), &item)
		unpacked, d := unpackBgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInner ---
func packBgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerListFromSdk(ctx context.Context, sdks []network_services.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInner")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInner
		obj, d := packBgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInner{}.AttrType(), data)
}

// --- Unpacker for BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatch ---
func unpackBgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatchToSdk(ctx context.Context, obj types.Object) (*network_services.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatch, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatch", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatch
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatch
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Interface.IsNull() && !model.Interface.IsUnknown() {
		sdk.Interface = model.Interface.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Interface", "value": *sdk.Interface})
	}

	// Handling Objects
	if !model.Ipv4.IsNull() && !model.Ipv4.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Ipv4")
		unpacked, d := unpackBgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatchIpv4ToSdk(ctx, model.Ipv4)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Ipv4"})
		}
		if unpacked != nil {
			sdk.Ipv4 = unpacked
		}
	}

	// Handling Primitives
	if !model.Metric.IsNull() && !model.Metric.IsUnknown() {
		val := int32(model.Metric.ValueInt64())
		sdk.Metric = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Metric", "value": *sdk.Metric})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatch", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatch ---
func packBgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatchFromSdk(ctx context.Context, sdk network_services.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatch) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatch", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatch
	var d diag.Diagnostics
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
		packed, d := packBgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatchIpv4FromSdk(ctx, *sdk.Ipv4)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Ipv4"})
		}
		model.Ipv4 = packed
	} else {
		model.Ipv4 = basetypes.NewObjectNull(models.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatchIpv4{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Metric != nil {
		model.Metric = basetypes.NewInt64Value(int64(*sdk.Metric))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Metric", "value": *sdk.Metric})
	} else {
		model.Metric = basetypes.NewInt64Null()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatch{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatch", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatch ---
func unpackBgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatchListToSdk(ctx context.Context, list types.List) ([]network_services.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatch, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatch")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatch
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatch, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatch{}.AttrTypes(), &item)
		unpacked, d := unpackBgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatchToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatch", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatch ---
func packBgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatchListFromSdk(ctx context.Context, sdks []network_services.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatch) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatch")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatch

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatch
		obj, d := packBgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatchFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatch{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatch", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatch{}.AttrType(), data)
}

// --- Unpacker for BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatchIpv4 ---
func unpackBgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatchIpv4ToSdk(ctx context.Context, obj types.Object) (*network_services.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatchIpv4, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatchIpv4", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatchIpv4
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatchIpv4
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

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatchIpv4", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatchIpv4 ---
func packBgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatchIpv4FromSdk(ctx context.Context, sdk network_services.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatchIpv4) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatchIpv4", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatchIpv4
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
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatchIpv4{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatchIpv4", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatchIpv4 ---
func unpackBgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatchIpv4ListToSdk(ctx context.Context, list types.List) ([]network_services.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatchIpv4, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatchIpv4")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatchIpv4
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatchIpv4, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatchIpv4{}.AttrTypes(), &item)
		unpacked, d := unpackBgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatchIpv4ToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatchIpv4", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatchIpv4 ---
func packBgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatchIpv4ListFromSdk(ctx context.Context, sdks []network_services.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatchIpv4) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatchIpv4")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatchIpv4

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatchIpv4
		obj, d := packBgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatchIpv4FromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatchIpv4{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatchIpv4", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatchIpv4{}.AttrType(), data)
}

// --- Unpacker for BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerSet ---
func unpackBgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerSetToSdk(ctx context.Context, obj types.Object) (*network_services.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerSet, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerSet", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerSet
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerSet
	var d diag.Diagnostics
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
	if !model.MetricType.IsNull() && !model.MetricType.IsUnknown() {
		sdk.MetricType = model.MetricType.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "MetricType", "value": *sdk.MetricType})
	}

	// Handling Primitives
	if !model.Tag.IsNull() && !model.Tag.IsUnknown() {
		val := int32(model.Tag.ValueInt64())
		sdk.Tag = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Tag", "value": *sdk.Tag})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerSet", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerSet ---
func packBgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerSetFromSdk(ctx context.Context, sdk network_services.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerSet) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerSet", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerSet
	var d diag.Diagnostics
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
	if sdk.MetricType != nil {
		model.MetricType = basetypes.NewStringValue(*sdk.MetricType)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "MetricType", "value": *sdk.MetricType})
	} else {
		model.MetricType = basetypes.NewStringNull()
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

	obj, d := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerSet{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerSet", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerSet ---
func unpackBgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerSetListToSdk(ctx context.Context, list types.List) ([]network_services.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerSet, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerSet")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerSet
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerSet, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerSet{}.AttrTypes(), &item)
		unpacked, d := unpackBgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerSetToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerSet", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerSet ---
func packBgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerSetListFromSdk(ctx context.Context, sdks []network_services.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerSet) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerSet")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerSet

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerSet
		obj, d := packBgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerSetFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerSet{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerSet", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerSet{}.AttrType(), data)
}

// --- Unpacker for BgpRouteMapRedistributionsConnectedStaticRib ---
func unpackBgpRouteMapRedistributionsConnectedStaticRibToSdk(ctx context.Context, obj types.Object) (*network_services.BgpRouteMapRedistributionsConnectedStaticRib, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpRouteMapRedistributionsConnectedStaticRib", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsConnectedStaticRib
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpRouteMapRedistributionsConnectedStaticRib
	var d diag.Diagnostics
	// Handling Lists
	if !model.RouteMap.IsNull() && !model.RouteMap.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field RouteMap")
		unpacked, d := unpackBgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerListToSdk(ctx, model.RouteMap)
		diags.Append(d...)
		sdk.RouteMap = unpacked
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpRouteMapRedistributionsConnectedStaticRib", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpRouteMapRedistributionsConnectedStaticRib ---
func packBgpRouteMapRedistributionsConnectedStaticRibFromSdk(ctx context.Context, sdk network_services.BgpRouteMapRedistributionsConnectedStaticRib) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpRouteMapRedistributionsConnectedStaticRib", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsConnectedStaticRib
	var d diag.Diagnostics
	// Handling Lists
	if sdk.RouteMap != nil {
		tflog.Debug(ctx, "Packing list of objects for field RouteMap")
		packed, d := packBgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerListFromSdk(ctx, sdk.RouteMap)
		diags.Append(d...)
		model.RouteMap = packed
	} else {
		model.RouteMap = basetypes.NewListNull(models.BgpRouteMapRedistributionsConnectedStaticRibRouteMapInner{}.AttrType())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsConnectedStaticRib{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpRouteMapRedistributionsConnectedStaticRib", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpRouteMapRedistributionsConnectedStaticRib ---
func unpackBgpRouteMapRedistributionsConnectedStaticRibListToSdk(ctx context.Context, list types.List) ([]network_services.BgpRouteMapRedistributionsConnectedStaticRib, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpRouteMapRedistributionsConnectedStaticRib")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsConnectedStaticRib
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpRouteMapRedistributionsConnectedStaticRib, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsConnectedStaticRib{}.AttrTypes(), &item)
		unpacked, d := unpackBgpRouteMapRedistributionsConnectedStaticRibToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpRouteMapRedistributionsConnectedStaticRib", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpRouteMapRedistributionsConnectedStaticRib ---
func packBgpRouteMapRedistributionsConnectedStaticRibListFromSdk(ctx context.Context, sdks []network_services.BgpRouteMapRedistributionsConnectedStaticRib) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpRouteMapRedistributionsConnectedStaticRib")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsConnectedStaticRib

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpRouteMapRedistributionsConnectedStaticRib
		obj, d := packBgpRouteMapRedistributionsConnectedStaticRibFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpRouteMapRedistributionsConnectedStaticRib{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpRouteMapRedistributionsConnectedStaticRib", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpRouteMapRedistributionsConnectedStaticRib{}.AttrType(), data)
}

// --- Unpacker for BgpRouteMapRedistributionsConnectedStaticRibRouteMapInner ---
func unpackBgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerToSdk(ctx context.Context, obj types.Object) (*network_services.BgpRouteMapRedistributionsConnectedStaticRibRouteMapInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpRouteMapRedistributionsConnectedStaticRibRouteMapInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsConnectedStaticRibRouteMapInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpRouteMapRedistributionsConnectedStaticRibRouteMapInner
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
		unpacked, d := unpackBgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatchToSdk(ctx, model.Match)
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
		unpacked, d := unpackBgpRouteMapRedistributionsBgpRibRouteMapInnerSetToSdk(ctx, model.Set)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Set"})
		}
		if unpacked != nil {
			sdk.Set = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpRouteMapRedistributionsConnectedStaticRibRouteMapInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpRouteMapRedistributionsConnectedStaticRibRouteMapInner ---
func packBgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerFromSdk(ctx context.Context, sdk network_services.BgpRouteMapRedistributionsConnectedStaticRibRouteMapInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpRouteMapRedistributionsConnectedStaticRibRouteMapInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsConnectedStaticRibRouteMapInner
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
		packed, d := packBgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatchFromSdk(ctx, *sdk.Match)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Match"})
		}
		model.Match = packed
	} else {
		model.Match = basetypes.NewObjectNull(models.BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatch{}.AttrTypes())
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
		packed, d := packBgpRouteMapRedistributionsBgpRibRouteMapInnerSetFromSdk(ctx, *sdk.Set)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Set"})
		}
		model.Set = packed
	} else {
		model.Set = basetypes.NewObjectNull(models.BgpRouteMapRedistributionsBgpRibRouteMapInnerSet{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsConnectedStaticRibRouteMapInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpRouteMapRedistributionsConnectedStaticRibRouteMapInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpRouteMapRedistributionsConnectedStaticRibRouteMapInner ---
func unpackBgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerListToSdk(ctx context.Context, list types.List) ([]network_services.BgpRouteMapRedistributionsConnectedStaticRibRouteMapInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpRouteMapRedistributionsConnectedStaticRibRouteMapInner")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsConnectedStaticRibRouteMapInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpRouteMapRedistributionsConnectedStaticRibRouteMapInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsConnectedStaticRibRouteMapInner{}.AttrTypes(), &item)
		unpacked, d := unpackBgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpRouteMapRedistributionsConnectedStaticRibRouteMapInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpRouteMapRedistributionsConnectedStaticRibRouteMapInner ---
func packBgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerListFromSdk(ctx context.Context, sdks []network_services.BgpRouteMapRedistributionsConnectedStaticRibRouteMapInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpRouteMapRedistributionsConnectedStaticRibRouteMapInner")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsConnectedStaticRibRouteMapInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpRouteMapRedistributionsConnectedStaticRibRouteMapInner
		obj, d := packBgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpRouteMapRedistributionsConnectedStaticRibRouteMapInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpRouteMapRedistributionsConnectedStaticRibRouteMapInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpRouteMapRedistributionsConnectedStaticRibRouteMapInner{}.AttrType(), data)
}

// --- Unpacker for BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatch ---
func unpackBgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatchToSdk(ctx context.Context, obj types.Object) (*network_services.BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatch, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatch", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatch
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatch
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Interface.IsNull() && !model.Interface.IsUnknown() {
		sdk.Interface = model.Interface.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Interface", "value": *sdk.Interface})
	}

	// Handling Objects
	if !model.Ipv4.IsNull() && !model.Ipv4.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Ipv4")
		unpacked, d := unpackBgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatchIpv4ToSdk(ctx, model.Ipv4)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Ipv4"})
		}
		if unpacked != nil {
			sdk.Ipv4 = unpacked
		}
	}

	// Handling Primitives
	if !model.Metric.IsNull() && !model.Metric.IsUnknown() {
		val := int32(model.Metric.ValueInt64())
		sdk.Metric = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Metric", "value": *sdk.Metric})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatch", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatch ---
func packBgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatchFromSdk(ctx context.Context, sdk network_services.BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatch) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatch", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatch
	var d diag.Diagnostics
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
		packed, d := packBgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatchIpv4FromSdk(ctx, *sdk.Ipv4)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Ipv4"})
		}
		model.Ipv4 = packed
	} else {
		model.Ipv4 = basetypes.NewObjectNull(models.BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatchIpv4{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Metric != nil {
		model.Metric = basetypes.NewInt64Value(int64(*sdk.Metric))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Metric", "value": *sdk.Metric})
	} else {
		model.Metric = basetypes.NewInt64Null()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatch{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatch", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatch ---
func unpackBgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatchListToSdk(ctx context.Context, list types.List) ([]network_services.BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatch, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatch")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatch
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatch, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatch{}.AttrTypes(), &item)
		unpacked, d := unpackBgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatchToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatch", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatch ---
func packBgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatchListFromSdk(ctx context.Context, sdks []network_services.BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatch) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatch")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatch

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatch
		obj, d := packBgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatchFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatch{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatch", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatch{}.AttrType(), data)
}

// --- Unpacker for BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatchIpv4 ---
func unpackBgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatchIpv4ToSdk(ctx context.Context, obj types.Object) (*network_services.BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatchIpv4, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatchIpv4", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatchIpv4
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatchIpv4
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

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatchIpv4", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatchIpv4 ---
func packBgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatchIpv4FromSdk(ctx context.Context, sdk network_services.BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatchIpv4) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatchIpv4", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatchIpv4
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
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatchIpv4{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatchIpv4", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatchIpv4 ---
func unpackBgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatchIpv4ListToSdk(ctx context.Context, list types.List) ([]network_services.BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatchIpv4, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatchIpv4")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatchIpv4
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatchIpv4, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatchIpv4{}.AttrTypes(), &item)
		unpacked, d := unpackBgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatchIpv4ToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatchIpv4", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatchIpv4 ---
func packBgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatchIpv4ListFromSdk(ctx context.Context, sdks []network_services.BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatchIpv4) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatchIpv4")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatchIpv4

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatchIpv4
		obj, d := packBgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatchIpv4FromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatchIpv4{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatchIpv4", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatchIpv4{}.AttrType(), data)
}

// --- Unpacker for BgpRouteMapRedistributionsOspf ---
func unpackBgpRouteMapRedistributionsOspfToSdk(ctx context.Context, obj types.Object) (*network_services.BgpRouteMapRedistributionsOspf, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpRouteMapRedistributionsOspf", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsOspf
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpRouteMapRedistributionsOspf
	var d diag.Diagnostics
	// Handling Objects
	if !model.Bgp.IsNull() && !model.Bgp.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Bgp")
		unpacked, d := unpackBgpRouteMapRedistributionsOspfBgpToSdk(ctx, model.Bgp)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Bgp"})
		}
		if unpacked != nil {
			sdk.Bgp = unpacked
		}
	}

	// Handling Objects
	if !model.Rib.IsNull() && !model.Rib.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Rib")
		unpacked, d := unpackBgpRouteMapRedistributionsOspfRibToSdk(ctx, model.Rib)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Rib"})
		}
		if unpacked != nil {
			sdk.Rib = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpRouteMapRedistributionsOspf", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpRouteMapRedistributionsOspf ---
func packBgpRouteMapRedistributionsOspfFromSdk(ctx context.Context, sdk network_services.BgpRouteMapRedistributionsOspf) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpRouteMapRedistributionsOspf", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsOspf
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Bgp != nil {
		tflog.Debug(ctx, "Packing nested object for field Bgp")
		packed, d := packBgpRouteMapRedistributionsOspfBgpFromSdk(ctx, *sdk.Bgp)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Bgp"})
		}
		model.Bgp = packed
	} else {
		model.Bgp = basetypes.NewObjectNull(models.BgpRouteMapRedistributionsOspfBgp{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Rib != nil {
		tflog.Debug(ctx, "Packing nested object for field Rib")
		packed, d := packBgpRouteMapRedistributionsOspfRibFromSdk(ctx, *sdk.Rib)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Rib"})
		}
		model.Rib = packed
	} else {
		model.Rib = basetypes.NewObjectNull(models.BgpRouteMapRedistributionsOspfRib{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsOspf{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpRouteMapRedistributionsOspf", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpRouteMapRedistributionsOspf ---
func unpackBgpRouteMapRedistributionsOspfListToSdk(ctx context.Context, list types.List) ([]network_services.BgpRouteMapRedistributionsOspf, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpRouteMapRedistributionsOspf")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsOspf
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpRouteMapRedistributionsOspf, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsOspf{}.AttrTypes(), &item)
		unpacked, d := unpackBgpRouteMapRedistributionsOspfToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpRouteMapRedistributionsOspf", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpRouteMapRedistributionsOspf ---
func packBgpRouteMapRedistributionsOspfListFromSdk(ctx context.Context, sdks []network_services.BgpRouteMapRedistributionsOspf) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpRouteMapRedistributionsOspf")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsOspf

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpRouteMapRedistributionsOspf
		obj, d := packBgpRouteMapRedistributionsOspfFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpRouteMapRedistributionsOspf{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpRouteMapRedistributionsOspf", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpRouteMapRedistributionsOspf{}.AttrType(), data)
}

// --- Unpacker for BgpRouteMapRedistributionsOspfBgp ---
func unpackBgpRouteMapRedistributionsOspfBgpToSdk(ctx context.Context, obj types.Object) (*network_services.BgpRouteMapRedistributionsOspfBgp, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpRouteMapRedistributionsOspfBgp", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsOspfBgp
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpRouteMapRedistributionsOspfBgp
	var d diag.Diagnostics
	// Handling Lists
	if !model.RouteMap.IsNull() && !model.RouteMap.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field RouteMap")
		unpacked, d := unpackBgpRouteMapRedistributionsOspfBgpRouteMapInnerListToSdk(ctx, model.RouteMap)
		diags.Append(d...)
		sdk.RouteMap = unpacked
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpRouteMapRedistributionsOspfBgp", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpRouteMapRedistributionsOspfBgp ---
func packBgpRouteMapRedistributionsOspfBgpFromSdk(ctx context.Context, sdk network_services.BgpRouteMapRedistributionsOspfBgp) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpRouteMapRedistributionsOspfBgp", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsOspfBgp
	var d diag.Diagnostics
	// Handling Lists
	if sdk.RouteMap != nil {
		tflog.Debug(ctx, "Packing list of objects for field RouteMap")
		packed, d := packBgpRouteMapRedistributionsOspfBgpRouteMapInnerListFromSdk(ctx, sdk.RouteMap)
		diags.Append(d...)
		model.RouteMap = packed
	} else {
		model.RouteMap = basetypes.NewListNull(models.BgpRouteMapRedistributionsOspfBgpRouteMapInner{}.AttrType())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsOspfBgp{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpRouteMapRedistributionsOspfBgp", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpRouteMapRedistributionsOspfBgp ---
func unpackBgpRouteMapRedistributionsOspfBgpListToSdk(ctx context.Context, list types.List) ([]network_services.BgpRouteMapRedistributionsOspfBgp, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpRouteMapRedistributionsOspfBgp")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsOspfBgp
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpRouteMapRedistributionsOspfBgp, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsOspfBgp{}.AttrTypes(), &item)
		unpacked, d := unpackBgpRouteMapRedistributionsOspfBgpToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpRouteMapRedistributionsOspfBgp", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpRouteMapRedistributionsOspfBgp ---
func packBgpRouteMapRedistributionsOspfBgpListFromSdk(ctx context.Context, sdks []network_services.BgpRouteMapRedistributionsOspfBgp) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpRouteMapRedistributionsOspfBgp")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsOspfBgp

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpRouteMapRedistributionsOspfBgp
		obj, d := packBgpRouteMapRedistributionsOspfBgpFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpRouteMapRedistributionsOspfBgp{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpRouteMapRedistributionsOspfBgp", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpRouteMapRedistributionsOspfBgp{}.AttrType(), data)
}

// --- Unpacker for BgpRouteMapRedistributionsOspfBgpRouteMapInner ---
func unpackBgpRouteMapRedistributionsOspfBgpRouteMapInnerToSdk(ctx context.Context, obj types.Object) (*network_services.BgpRouteMapRedistributionsOspfBgpRouteMapInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpRouteMapRedistributionsOspfBgpRouteMapInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsOspfBgpRouteMapInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpRouteMapRedistributionsOspfBgpRouteMapInner
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
		unpacked, d := unpackBgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchToSdk(ctx, model.Match)
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
		unpacked, d := unpackBgpRouteMapRedistributionsOspfBgpRouteMapInnerSetToSdk(ctx, model.Set)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Set"})
		}
		if unpacked != nil {
			sdk.Set = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpRouteMapRedistributionsOspfBgpRouteMapInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpRouteMapRedistributionsOspfBgpRouteMapInner ---
func packBgpRouteMapRedistributionsOspfBgpRouteMapInnerFromSdk(ctx context.Context, sdk network_services.BgpRouteMapRedistributionsOspfBgpRouteMapInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpRouteMapRedistributionsOspfBgpRouteMapInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsOspfBgpRouteMapInner
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
		packed, d := packBgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchFromSdk(ctx, *sdk.Match)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Match"})
		}
		model.Match = packed
	} else {
		model.Match = basetypes.NewObjectNull(models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatch{}.AttrTypes())
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
		packed, d := packBgpRouteMapRedistributionsOspfBgpRouteMapInnerSetFromSdk(ctx, *sdk.Set)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Set"})
		}
		model.Set = packed
	} else {
		model.Set = basetypes.NewObjectNull(models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerSet{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsOspfBgpRouteMapInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpRouteMapRedistributionsOspfBgpRouteMapInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpRouteMapRedistributionsOspfBgpRouteMapInner ---
func unpackBgpRouteMapRedistributionsOspfBgpRouteMapInnerListToSdk(ctx context.Context, list types.List) ([]network_services.BgpRouteMapRedistributionsOspfBgpRouteMapInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpRouteMapRedistributionsOspfBgpRouteMapInner")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsOspfBgpRouteMapInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpRouteMapRedistributionsOspfBgpRouteMapInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsOspfBgpRouteMapInner{}.AttrTypes(), &item)
		unpacked, d := unpackBgpRouteMapRedistributionsOspfBgpRouteMapInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpRouteMapRedistributionsOspfBgpRouteMapInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpRouteMapRedistributionsOspfBgpRouteMapInner ---
func packBgpRouteMapRedistributionsOspfBgpRouteMapInnerListFromSdk(ctx context.Context, sdks []network_services.BgpRouteMapRedistributionsOspfBgpRouteMapInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpRouteMapRedistributionsOspfBgpRouteMapInner")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsOspfBgpRouteMapInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpRouteMapRedistributionsOspfBgpRouteMapInner
		obj, d := packBgpRouteMapRedistributionsOspfBgpRouteMapInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpRouteMapRedistributionsOspfBgpRouteMapInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpRouteMapRedistributionsOspfBgpRouteMapInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpRouteMapRedistributionsOspfBgpRouteMapInner{}.AttrType(), data)
}

// --- Unpacker for BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatch ---
func unpackBgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchToSdk(ctx context.Context, obj types.Object) (*network_services.BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatch, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatch", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatch
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatch
	var d diag.Diagnostics
	// Handling Objects
	if !model.Address.IsNull() && !model.Address.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Address")
		unpacked, d := unpackBgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchAddressToSdk(ctx, model.Address)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Address"})
		}
		if unpacked != nil {
			sdk.Address = unpacked
		}
	}

	// Handling Primitives
	if !model.Interface.IsNull() && !model.Interface.IsUnknown() {
		sdk.Interface = model.Interface.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Interface", "value": *sdk.Interface})
	}

	// Handling Primitives
	if !model.Metric.IsNull() && !model.Metric.IsUnknown() {
		val := int32(model.Metric.ValueInt64())
		sdk.Metric = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Metric", "value": *sdk.Metric})
	}

	// Handling Objects
	if !model.NextHop.IsNull() && !model.NextHop.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field NextHop")
		unpacked, d := unpackBgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchNextHopToSdk(ctx, model.NextHop)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "NextHop"})
		}
		if unpacked != nil {
			sdk.NextHop = unpacked
		}
	}

	// Handling Primitives
	if !model.Tag.IsNull() && !model.Tag.IsUnknown() {
		val := int32(model.Tag.ValueInt64())
		sdk.Tag = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Tag", "value": *sdk.Tag})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatch", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatch ---
func packBgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchFromSdk(ctx context.Context, sdk network_services.BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatch) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatch", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatch
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Address != nil {
		tflog.Debug(ctx, "Packing nested object for field Address")
		packed, d := packBgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchAddressFromSdk(ctx, *sdk.Address)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Address"})
		}
		model.Address = packed
	} else {
		model.Address = basetypes.NewObjectNull(models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchAddress{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Interface != nil {
		model.Interface = basetypes.NewStringValue(*sdk.Interface)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Interface", "value": *sdk.Interface})
	} else {
		model.Interface = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Metric != nil {
		model.Metric = basetypes.NewInt64Value(int64(*sdk.Metric))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Metric", "value": *sdk.Metric})
	} else {
		model.Metric = basetypes.NewInt64Null()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.NextHop != nil {
		tflog.Debug(ctx, "Packing nested object for field NextHop")
		packed, d := packBgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchNextHopFromSdk(ctx, *sdk.NextHop)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "NextHop"})
		}
		model.NextHop = packed
	} else {
		model.NextHop = basetypes.NewObjectNull(models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchNextHop{}.AttrTypes())
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

	obj, d := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatch{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatch", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatch ---
func unpackBgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchListToSdk(ctx context.Context, list types.List) ([]network_services.BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatch, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatch")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatch
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatch, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatch{}.AttrTypes(), &item)
		unpacked, d := unpackBgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatch", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatch ---
func packBgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchListFromSdk(ctx context.Context, sdks []network_services.BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatch) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatch")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatch

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatch
		obj, d := packBgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatch{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatch", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatch{}.AttrType(), data)
}

// --- Unpacker for BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchAddress ---
func unpackBgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchAddressToSdk(ctx context.Context, obj types.Object) (*network_services.BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchAddress, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchAddress", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchAddress
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchAddress
	var d diag.Diagnostics
	// Handling Primitives
	if !model.AccessList.IsNull() && !model.AccessList.IsUnknown() {
		sdk.AccessList = model.AccessList.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "AccessList", "value": *sdk.AccessList})
	}

	// Handling Primitives
	if !model.PrefixList.IsNull() && !model.PrefixList.IsUnknown() {
		sdk.PrefixList = model.PrefixList.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "PrefixList", "value": *sdk.PrefixList})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchAddress", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchAddress ---
func packBgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchAddressFromSdk(ctx context.Context, sdk network_services.BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchAddress) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchAddress", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchAddress
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.AccessList != nil {
		model.AccessList = basetypes.NewStringValue(*sdk.AccessList)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "AccessList", "value": *sdk.AccessList})
	} else {
		model.AccessList = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.PrefixList != nil {
		model.PrefixList = basetypes.NewStringValue(*sdk.PrefixList)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "PrefixList", "value": *sdk.PrefixList})
	} else {
		model.PrefixList = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchAddress{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchAddress", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchAddress ---
func unpackBgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchAddressListToSdk(ctx context.Context, list types.List) ([]network_services.BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchAddress, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchAddress")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchAddress
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchAddress, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchAddress{}.AttrTypes(), &item)
		unpacked, d := unpackBgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchAddressToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchAddress", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchAddress ---
func packBgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchAddressListFromSdk(ctx context.Context, sdks []network_services.BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchAddress) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchAddress")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchAddress

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchAddress
		obj, d := packBgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchAddressFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchAddress{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchAddress", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchAddress{}.AttrType(), data)
}

// --- Unpacker for BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchNextHop ---
func unpackBgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchNextHopToSdk(ctx context.Context, obj types.Object) (*network_services.BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchNextHop, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchNextHop", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchNextHop
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchNextHop
	var d diag.Diagnostics
	// Handling Primitives
	if !model.AccessList.IsNull() && !model.AccessList.IsUnknown() {
		sdk.AccessList = model.AccessList.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "AccessList", "value": *sdk.AccessList})
	}

	// Handling Primitives
	if !model.PrefixList.IsNull() && !model.PrefixList.IsUnknown() {
		sdk.PrefixList = model.PrefixList.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "PrefixList", "value": *sdk.PrefixList})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchNextHop", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchNextHop ---
func packBgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchNextHopFromSdk(ctx context.Context, sdk network_services.BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchNextHop) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchNextHop", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchNextHop
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.AccessList != nil {
		model.AccessList = basetypes.NewStringValue(*sdk.AccessList)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "AccessList", "value": *sdk.AccessList})
	} else {
		model.AccessList = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.PrefixList != nil {
		model.PrefixList = basetypes.NewStringValue(*sdk.PrefixList)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "PrefixList", "value": *sdk.PrefixList})
	} else {
		model.PrefixList = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchNextHop{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchNextHop", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchNextHop ---
func unpackBgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchNextHopListToSdk(ctx context.Context, list types.List) ([]network_services.BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchNextHop, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchNextHop")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchNextHop
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchNextHop, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchNextHop{}.AttrTypes(), &item)
		unpacked, d := unpackBgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchNextHopToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchNextHop", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchNextHop ---
func packBgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchNextHopListFromSdk(ctx context.Context, sdks []network_services.BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchNextHop) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchNextHop")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchNextHop

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchNextHop
		obj, d := packBgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchNextHopFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchNextHop{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchNextHop", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchNextHop{}.AttrType(), data)
}

// --- Unpacker for BgpRouteMapRedistributionsOspfBgpRouteMapInnerSet ---
func unpackBgpRouteMapRedistributionsOspfBgpRouteMapInnerSetToSdk(ctx context.Context, obj types.Object) (*network_services.BgpRouteMapRedistributionsOspfBgpRouteMapInnerSet, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerSet", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerSet
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpRouteMapRedistributionsOspfBgpRouteMapInnerSet
	var d diag.Diagnostics
	// Handling Objects
	if !model.Aggregator.IsNull() && !model.Aggregator.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Aggregator")
		unpacked, d := unpackBgpRouteMapRedistributionsOspfBgpRouteMapInnerSetAggregatorToSdk(ctx, model.Aggregator)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Aggregator"})
		}
		if unpacked != nil {
			sdk.Aggregator = unpacked
		}
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

	// Handling Lists
	if !model.RegularCommunity.IsNull() && !model.RegularCommunity.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field RegularCommunity")
		diags.Append(model.RegularCommunity.ElementsAs(ctx, &sdk.RegularCommunity, false)...)
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

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerSet", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpRouteMapRedistributionsOspfBgpRouteMapInnerSet ---
func packBgpRouteMapRedistributionsOspfBgpRouteMapInnerSetFromSdk(ctx context.Context, sdk network_services.BgpRouteMapRedistributionsOspfBgpRouteMapInnerSet) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerSet", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerSet
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Aggregator != nil {
		tflog.Debug(ctx, "Packing nested object for field Aggregator")
		packed, d := packBgpRouteMapRedistributionsOspfBgpRouteMapInnerSetAggregatorFromSdk(ctx, *sdk.Aggregator)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Aggregator"})
		}
		model.Aggregator = packed
	} else {
		model.Aggregator = basetypes.NewObjectNull(models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerSetAggregator{}.AttrTypes())
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

	obj, d := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerSet{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerSet", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpRouteMapRedistributionsOspfBgpRouteMapInnerSet ---
func unpackBgpRouteMapRedistributionsOspfBgpRouteMapInnerSetListToSdk(ctx context.Context, list types.List) ([]network_services.BgpRouteMapRedistributionsOspfBgpRouteMapInnerSet, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerSet")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerSet
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpRouteMapRedistributionsOspfBgpRouteMapInnerSet, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerSet{}.AttrTypes(), &item)
		unpacked, d := unpackBgpRouteMapRedistributionsOspfBgpRouteMapInnerSetToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerSet", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpRouteMapRedistributionsOspfBgpRouteMapInnerSet ---
func packBgpRouteMapRedistributionsOspfBgpRouteMapInnerSetListFromSdk(ctx context.Context, sdks []network_services.BgpRouteMapRedistributionsOspfBgpRouteMapInnerSet) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerSet")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerSet

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerSet
		obj, d := packBgpRouteMapRedistributionsOspfBgpRouteMapInnerSetFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerSet{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerSet", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerSet{}.AttrType(), data)
}

// --- Unpacker for BgpRouteMapRedistributionsOspfBgpRouteMapInnerSetAggregator ---
func unpackBgpRouteMapRedistributionsOspfBgpRouteMapInnerSetAggregatorToSdk(ctx context.Context, obj types.Object) (*network_services.BgpRouteMapRedistributionsOspfBgpRouteMapInnerSetAggregator, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerSetAggregator", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerSetAggregator
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpRouteMapRedistributionsOspfBgpRouteMapInnerSetAggregator
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

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerSetAggregator", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpRouteMapRedistributionsOspfBgpRouteMapInnerSetAggregator ---
func packBgpRouteMapRedistributionsOspfBgpRouteMapInnerSetAggregatorFromSdk(ctx context.Context, sdk network_services.BgpRouteMapRedistributionsOspfBgpRouteMapInnerSetAggregator) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerSetAggregator", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerSetAggregator
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

	obj, d := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerSetAggregator{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerSetAggregator", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpRouteMapRedistributionsOspfBgpRouteMapInnerSetAggregator ---
func unpackBgpRouteMapRedistributionsOspfBgpRouteMapInnerSetAggregatorListToSdk(ctx context.Context, list types.List) ([]network_services.BgpRouteMapRedistributionsOspfBgpRouteMapInnerSetAggregator, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerSetAggregator")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerSetAggregator
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpRouteMapRedistributionsOspfBgpRouteMapInnerSetAggregator, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerSetAggregator{}.AttrTypes(), &item)
		unpacked, d := unpackBgpRouteMapRedistributionsOspfBgpRouteMapInnerSetAggregatorToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerSetAggregator", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpRouteMapRedistributionsOspfBgpRouteMapInnerSetAggregator ---
func packBgpRouteMapRedistributionsOspfBgpRouteMapInnerSetAggregatorListFromSdk(ctx context.Context, sdks []network_services.BgpRouteMapRedistributionsOspfBgpRouteMapInnerSetAggregator) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerSetAggregator")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerSetAggregator

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerSetAggregator
		obj, d := packBgpRouteMapRedistributionsOspfBgpRouteMapInnerSetAggregatorFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerSetAggregator{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerSetAggregator", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpRouteMapRedistributionsOspfBgpRouteMapInnerSetAggregator{}.AttrType(), data)
}

// --- Unpacker for BgpRouteMapRedistributionsOspfRib ---
func unpackBgpRouteMapRedistributionsOspfRibToSdk(ctx context.Context, obj types.Object) (*network_services.BgpRouteMapRedistributionsOspfRib, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpRouteMapRedistributionsOspfRib", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsOspfRib
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpRouteMapRedistributionsOspfRib
	var d diag.Diagnostics
	// Handling Lists
	if !model.RouteMap.IsNull() && !model.RouteMap.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field RouteMap")
		unpacked, d := unpackBgpRouteMapRedistributionsOspfRibRouteMapInnerListToSdk(ctx, model.RouteMap)
		diags.Append(d...)
		sdk.RouteMap = unpacked
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpRouteMapRedistributionsOspfRib", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpRouteMapRedistributionsOspfRib ---
func packBgpRouteMapRedistributionsOspfRibFromSdk(ctx context.Context, sdk network_services.BgpRouteMapRedistributionsOspfRib) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpRouteMapRedistributionsOspfRib", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsOspfRib
	var d diag.Diagnostics
	// Handling Lists
	if sdk.RouteMap != nil {
		tflog.Debug(ctx, "Packing list of objects for field RouteMap")
		packed, d := packBgpRouteMapRedistributionsOspfRibRouteMapInnerListFromSdk(ctx, sdk.RouteMap)
		diags.Append(d...)
		model.RouteMap = packed
	} else {
		model.RouteMap = basetypes.NewListNull(models.BgpRouteMapRedistributionsOspfRibRouteMapInner{}.AttrType())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsOspfRib{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpRouteMapRedistributionsOspfRib", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpRouteMapRedistributionsOspfRib ---
func unpackBgpRouteMapRedistributionsOspfRibListToSdk(ctx context.Context, list types.List) ([]network_services.BgpRouteMapRedistributionsOspfRib, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpRouteMapRedistributionsOspfRib")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsOspfRib
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpRouteMapRedistributionsOspfRib, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsOspfRib{}.AttrTypes(), &item)
		unpacked, d := unpackBgpRouteMapRedistributionsOspfRibToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpRouteMapRedistributionsOspfRib", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpRouteMapRedistributionsOspfRib ---
func packBgpRouteMapRedistributionsOspfRibListFromSdk(ctx context.Context, sdks []network_services.BgpRouteMapRedistributionsOspfRib) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpRouteMapRedistributionsOspfRib")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsOspfRib

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpRouteMapRedistributionsOspfRib
		obj, d := packBgpRouteMapRedistributionsOspfRibFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpRouteMapRedistributionsOspfRib{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpRouteMapRedistributionsOspfRib", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpRouteMapRedistributionsOspfRib{}.AttrType(), data)
}

// --- Unpacker for BgpRouteMapRedistributionsOspfRibRouteMapInner ---
func unpackBgpRouteMapRedistributionsOspfRibRouteMapInnerToSdk(ctx context.Context, obj types.Object) (*network_services.BgpRouteMapRedistributionsOspfRibRouteMapInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpRouteMapRedistributionsOspfRibRouteMapInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsOspfRibRouteMapInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpRouteMapRedistributionsOspfRibRouteMapInner
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
		unpacked, d := unpackBgpRouteMapRedistributionsOspfRibRouteMapInnerMatchToSdk(ctx, model.Match)
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
		unpacked, d := unpackBgpRouteMapRedistributionsBgpRibRouteMapInnerSetToSdk(ctx, model.Set)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Set"})
		}
		if unpacked != nil {
			sdk.Set = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpRouteMapRedistributionsOspfRibRouteMapInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpRouteMapRedistributionsOspfRibRouteMapInner ---
func packBgpRouteMapRedistributionsOspfRibRouteMapInnerFromSdk(ctx context.Context, sdk network_services.BgpRouteMapRedistributionsOspfRibRouteMapInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpRouteMapRedistributionsOspfRibRouteMapInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsOspfRibRouteMapInner
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
		packed, d := packBgpRouteMapRedistributionsOspfRibRouteMapInnerMatchFromSdk(ctx, *sdk.Match)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Match"})
		}
		model.Match = packed
	} else {
		model.Match = basetypes.NewObjectNull(models.BgpRouteMapRedistributionsOspfRibRouteMapInnerMatch{}.AttrTypes())
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
		packed, d := packBgpRouteMapRedistributionsBgpRibRouteMapInnerSetFromSdk(ctx, *sdk.Set)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Set"})
		}
		model.Set = packed
	} else {
		model.Set = basetypes.NewObjectNull(models.BgpRouteMapRedistributionsBgpRibRouteMapInnerSet{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsOspfRibRouteMapInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpRouteMapRedistributionsOspfRibRouteMapInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpRouteMapRedistributionsOspfRibRouteMapInner ---
func unpackBgpRouteMapRedistributionsOspfRibRouteMapInnerListToSdk(ctx context.Context, list types.List) ([]network_services.BgpRouteMapRedistributionsOspfRibRouteMapInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpRouteMapRedistributionsOspfRibRouteMapInner")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsOspfRibRouteMapInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpRouteMapRedistributionsOspfRibRouteMapInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsOspfRibRouteMapInner{}.AttrTypes(), &item)
		unpacked, d := unpackBgpRouteMapRedistributionsOspfRibRouteMapInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpRouteMapRedistributionsOspfRibRouteMapInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpRouteMapRedistributionsOspfRibRouteMapInner ---
func packBgpRouteMapRedistributionsOspfRibRouteMapInnerListFromSdk(ctx context.Context, sdks []network_services.BgpRouteMapRedistributionsOspfRibRouteMapInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpRouteMapRedistributionsOspfRibRouteMapInner")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsOspfRibRouteMapInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpRouteMapRedistributionsOspfRibRouteMapInner
		obj, d := packBgpRouteMapRedistributionsOspfRibRouteMapInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpRouteMapRedistributionsOspfRibRouteMapInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpRouteMapRedistributionsOspfRibRouteMapInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpRouteMapRedistributionsOspfRibRouteMapInner{}.AttrType(), data)
}

// --- Unpacker for BgpRouteMapRedistributionsOspfRibRouteMapInnerMatch ---
func unpackBgpRouteMapRedistributionsOspfRibRouteMapInnerMatchToSdk(ctx context.Context, obj types.Object) (*network_services.BgpRouteMapRedistributionsOspfRibRouteMapInnerMatch, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpRouteMapRedistributionsOspfRibRouteMapInnerMatch", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsOspfRibRouteMapInnerMatch
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpRouteMapRedistributionsOspfRibRouteMapInnerMatch
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

	// Handling Primitives
	if !model.Interface.IsNull() && !model.Interface.IsUnknown() {
		sdk.Interface = model.Interface.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Interface", "value": *sdk.Interface})
	}

	// Handling Primitives
	if !model.Metric.IsNull() && !model.Metric.IsUnknown() {
		val := int32(model.Metric.ValueInt64())
		sdk.Metric = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Metric", "value": *sdk.Metric})
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

	// Handling Primitives
	if !model.Tag.IsNull() && !model.Tag.IsUnknown() {
		val := int32(model.Tag.ValueInt64())
		sdk.Tag = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Tag", "value": *sdk.Tag})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpRouteMapRedistributionsOspfRibRouteMapInnerMatch", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpRouteMapRedistributionsOspfRibRouteMapInnerMatch ---
func packBgpRouteMapRedistributionsOspfRibRouteMapInnerMatchFromSdk(ctx context.Context, sdk network_services.BgpRouteMapRedistributionsOspfRibRouteMapInnerMatch) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpRouteMapRedistributionsOspfRibRouteMapInnerMatch", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpRouteMapRedistributionsOspfRibRouteMapInnerMatch
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
	// Handling Primitives
	// Standard primitive packing
	if sdk.Interface != nil {
		model.Interface = basetypes.NewStringValue(*sdk.Interface)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Interface", "value": *sdk.Interface})
	} else {
		model.Interface = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Metric != nil {
		model.Metric = basetypes.NewInt64Value(int64(*sdk.Metric))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Metric", "value": *sdk.Metric})
	} else {
		model.Metric = basetypes.NewInt64Null()
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
	// Handling Primitives
	// Standard primitive packing
	if sdk.Tag != nil {
		model.Tag = basetypes.NewInt64Value(int64(*sdk.Tag))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Tag", "value": *sdk.Tag})
	} else {
		model.Tag = basetypes.NewInt64Null()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsOspfRibRouteMapInnerMatch{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpRouteMapRedistributionsOspfRibRouteMapInnerMatch", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpRouteMapRedistributionsOspfRibRouteMapInnerMatch ---
func unpackBgpRouteMapRedistributionsOspfRibRouteMapInnerMatchListToSdk(ctx context.Context, list types.List) ([]network_services.BgpRouteMapRedistributionsOspfRibRouteMapInnerMatch, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpRouteMapRedistributionsOspfRibRouteMapInnerMatch")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsOspfRibRouteMapInnerMatch
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpRouteMapRedistributionsOspfRibRouteMapInnerMatch, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpRouteMapRedistributionsOspfRibRouteMapInnerMatch{}.AttrTypes(), &item)
		unpacked, d := unpackBgpRouteMapRedistributionsOspfRibRouteMapInnerMatchToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpRouteMapRedistributionsOspfRibRouteMapInnerMatch", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpRouteMapRedistributionsOspfRibRouteMapInnerMatch ---
func packBgpRouteMapRedistributionsOspfRibRouteMapInnerMatchListFromSdk(ctx context.Context, sdks []network_services.BgpRouteMapRedistributionsOspfRibRouteMapInnerMatch) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpRouteMapRedistributionsOspfRibRouteMapInnerMatch")
	diags := diag.Diagnostics{}
	var data []models.BgpRouteMapRedistributionsOspfRibRouteMapInnerMatch

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpRouteMapRedistributionsOspfRibRouteMapInnerMatch
		obj, d := packBgpRouteMapRedistributionsOspfRibRouteMapInnerMatchFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpRouteMapRedistributionsOspfRibRouteMapInnerMatch{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpRouteMapRedistributionsOspfRibRouteMapInnerMatch", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpRouteMapRedistributionsOspfRibRouteMapInnerMatch{}.AttrType(), data)
}
