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

// --- Unpacker for BgpRedistributionProfiles ---
func unpackBgpRedistributionProfilesToSdk(ctx context.Context, obj types.Object) (*network_services.BgpRedistributionProfiles, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpRedistributionProfiles", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpRedistributionProfiles
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpRedistributionProfiles
	var d diag.Diagnostics
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

	// Handling Objects
	if !model.Ipv4.IsNull() && !model.Ipv4.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Ipv4")
		unpacked, d := unpackBgpRedistributionProfilesIpv4ToSdk(ctx, model.Ipv4)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Ipv4"})
		}
		if unpacked != nil {
			sdk.Ipv4 = *unpacked
		}
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

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpRedistributionProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpRedistributionProfiles ---
func packBgpRedistributionProfilesFromSdk(ctx context.Context, sdk network_services.BgpRedistributionProfiles) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpRedistributionProfiles", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpRedistributionProfiles
	var d diag.Diagnostics
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
	// Handling Objects
	// This is a regular nested object that has its own packer.
	// Logic for non-pointer / value-type nested objects
	if !reflect.ValueOf(sdk.Ipv4).IsZero() {
		tflog.Debug(ctx, "Packing nested object for field Ipv4")
		packed, d := packBgpRedistributionProfilesIpv4FromSdk(ctx, sdk.Ipv4)
		diags.Append(d...)
		model.Ipv4 = packed
	} else {
		model.Ipv4 = basetypes.NewObjectNull(models.BgpRedistributionProfilesIpv4{}.AttrTypes())
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
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BgpRedistributionProfiles{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpRedistributionProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpRedistributionProfiles ---
func unpackBgpRedistributionProfilesListToSdk(ctx context.Context, list types.List) ([]network_services.BgpRedistributionProfiles, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpRedistributionProfiles")
	diags := diag.Diagnostics{}
	var data []models.BgpRedistributionProfiles
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpRedistributionProfiles, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpRedistributionProfiles{}.AttrTypes(), &item)
		unpacked, d := unpackBgpRedistributionProfilesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpRedistributionProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpRedistributionProfiles ---
func packBgpRedistributionProfilesListFromSdk(ctx context.Context, sdks []network_services.BgpRedistributionProfiles) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpRedistributionProfiles")
	diags := diag.Diagnostics{}
	var data []models.BgpRedistributionProfiles

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpRedistributionProfiles
		obj, d := packBgpRedistributionProfilesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpRedistributionProfiles{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpRedistributionProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpRedistributionProfiles{}.AttrType(), data)
}

// --- Unpacker for BgpRedistributionProfilesIpv4 ---
func unpackBgpRedistributionProfilesIpv4ToSdk(ctx context.Context, obj types.Object) (*network_services.BgpRedistributionProfilesIpv4, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpRedistributionProfilesIpv4", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpRedistributionProfilesIpv4
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpRedistributionProfilesIpv4
	var d diag.Diagnostics
	// Handling Objects
	if !model.Unicast.IsNull() && !model.Unicast.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Unicast")
		unpacked, d := unpackBgpRedistributionProfilesIpv4UnicastToSdk(ctx, model.Unicast)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Unicast"})
		}
		if unpacked != nil {
			sdk.Unicast = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpRedistributionProfilesIpv4", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpRedistributionProfilesIpv4 ---
func packBgpRedistributionProfilesIpv4FromSdk(ctx context.Context, sdk network_services.BgpRedistributionProfilesIpv4) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpRedistributionProfilesIpv4", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpRedistributionProfilesIpv4
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Unicast != nil {
		tflog.Debug(ctx, "Packing nested object for field Unicast")
		packed, d := packBgpRedistributionProfilesIpv4UnicastFromSdk(ctx, *sdk.Unicast)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Unicast"})
		}
		model.Unicast = packed
	} else {
		model.Unicast = basetypes.NewObjectNull(models.BgpRedistributionProfilesIpv4Unicast{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BgpRedistributionProfilesIpv4{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpRedistributionProfilesIpv4", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpRedistributionProfilesIpv4 ---
func unpackBgpRedistributionProfilesIpv4ListToSdk(ctx context.Context, list types.List) ([]network_services.BgpRedistributionProfilesIpv4, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpRedistributionProfilesIpv4")
	diags := diag.Diagnostics{}
	var data []models.BgpRedistributionProfilesIpv4
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpRedistributionProfilesIpv4, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpRedistributionProfilesIpv4{}.AttrTypes(), &item)
		unpacked, d := unpackBgpRedistributionProfilesIpv4ToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpRedistributionProfilesIpv4", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpRedistributionProfilesIpv4 ---
func packBgpRedistributionProfilesIpv4ListFromSdk(ctx context.Context, sdks []network_services.BgpRedistributionProfilesIpv4) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpRedistributionProfilesIpv4")
	diags := diag.Diagnostics{}
	var data []models.BgpRedistributionProfilesIpv4

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpRedistributionProfilesIpv4
		obj, d := packBgpRedistributionProfilesIpv4FromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpRedistributionProfilesIpv4{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpRedistributionProfilesIpv4", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpRedistributionProfilesIpv4{}.AttrType(), data)
}

// --- Unpacker for BgpRedistributionProfilesIpv4Unicast ---
func unpackBgpRedistributionProfilesIpv4UnicastToSdk(ctx context.Context, obj types.Object) (*network_services.BgpRedistributionProfilesIpv4Unicast, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpRedistributionProfilesIpv4Unicast", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpRedistributionProfilesIpv4Unicast
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpRedistributionProfilesIpv4Unicast
	var d diag.Diagnostics
	// Handling Objects
	if !model.Connected.IsNull() && !model.Connected.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Connected")
		unpacked, d := unpackBgpRedistributionProfilesIpv4UnicastConnectedToSdk(ctx, model.Connected)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Connected"})
		}
		if unpacked != nil {
			sdk.Connected = unpacked
		}
	}

	// Handling Objects
	if !model.Ospf.IsNull() && !model.Ospf.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Ospf")
		unpacked, d := unpackBgpRedistributionProfilesIpv4UnicastOspfToSdk(ctx, model.Ospf)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Ospf"})
		}
		if unpacked != nil {
			sdk.Ospf = unpacked
		}
	}

	// Handling Objects
	if !model.Static.IsNull() && !model.Static.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Static")
		unpacked, d := unpackBgpRedistributionProfilesIpv4UnicastStaticToSdk(ctx, model.Static)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Static"})
		}
		if unpacked != nil {
			sdk.Static = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpRedistributionProfilesIpv4Unicast", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpRedistributionProfilesIpv4Unicast ---
func packBgpRedistributionProfilesIpv4UnicastFromSdk(ctx context.Context, sdk network_services.BgpRedistributionProfilesIpv4Unicast) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpRedistributionProfilesIpv4Unicast", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpRedistributionProfilesIpv4Unicast
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Connected != nil {
		tflog.Debug(ctx, "Packing nested object for field Connected")
		packed, d := packBgpRedistributionProfilesIpv4UnicastConnectedFromSdk(ctx, *sdk.Connected)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Connected"})
		}
		model.Connected = packed
	} else {
		model.Connected = basetypes.NewObjectNull(models.BgpRedistributionProfilesIpv4UnicastConnected{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Ospf != nil {
		tflog.Debug(ctx, "Packing nested object for field Ospf")
		packed, d := packBgpRedistributionProfilesIpv4UnicastOspfFromSdk(ctx, *sdk.Ospf)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Ospf"})
		}
		model.Ospf = packed
	} else {
		model.Ospf = basetypes.NewObjectNull(models.BgpRedistributionProfilesIpv4UnicastOspf{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Static != nil {
		tflog.Debug(ctx, "Packing nested object for field Static")
		packed, d := packBgpRedistributionProfilesIpv4UnicastStaticFromSdk(ctx, *sdk.Static)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Static"})
		}
		model.Static = packed
	} else {
		model.Static = basetypes.NewObjectNull(models.BgpRedistributionProfilesIpv4UnicastStatic{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BgpRedistributionProfilesIpv4Unicast{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpRedistributionProfilesIpv4Unicast", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpRedistributionProfilesIpv4Unicast ---
func unpackBgpRedistributionProfilesIpv4UnicastListToSdk(ctx context.Context, list types.List) ([]network_services.BgpRedistributionProfilesIpv4Unicast, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpRedistributionProfilesIpv4Unicast")
	diags := diag.Diagnostics{}
	var data []models.BgpRedistributionProfilesIpv4Unicast
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpRedistributionProfilesIpv4Unicast, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpRedistributionProfilesIpv4Unicast{}.AttrTypes(), &item)
		unpacked, d := unpackBgpRedistributionProfilesIpv4UnicastToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpRedistributionProfilesIpv4Unicast", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpRedistributionProfilesIpv4Unicast ---
func packBgpRedistributionProfilesIpv4UnicastListFromSdk(ctx context.Context, sdks []network_services.BgpRedistributionProfilesIpv4Unicast) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpRedistributionProfilesIpv4Unicast")
	diags := diag.Diagnostics{}
	var data []models.BgpRedistributionProfilesIpv4Unicast

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpRedistributionProfilesIpv4Unicast
		obj, d := packBgpRedistributionProfilesIpv4UnicastFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpRedistributionProfilesIpv4Unicast{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpRedistributionProfilesIpv4Unicast", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpRedistributionProfilesIpv4Unicast{}.AttrType(), data)
}

// --- Unpacker for BgpRedistributionProfilesIpv4UnicastConnected ---
func unpackBgpRedistributionProfilesIpv4UnicastConnectedToSdk(ctx context.Context, obj types.Object) (*network_services.BgpRedistributionProfilesIpv4UnicastConnected, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpRedistributionProfilesIpv4UnicastConnected", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpRedistributionProfilesIpv4UnicastConnected
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpRedistributionProfilesIpv4UnicastConnected
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Enable.IsNull() && !model.Enable.IsUnknown() {
		sdk.Enable = model.Enable.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Enable", "value": *sdk.Enable})
	}

	// Handling Primitives
	if !model.Metric.IsNull() && !model.Metric.IsUnknown() {
		val := int32(model.Metric.ValueInt64())
		sdk.Metric = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Metric", "value": *sdk.Metric})
	}

	// Handling Primitives
	if !model.RouteMap.IsNull() && !model.RouteMap.IsUnknown() {
		sdk.RouteMap = model.RouteMap.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "RouteMap", "value": *sdk.RouteMap})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpRedistributionProfilesIpv4UnicastConnected", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpRedistributionProfilesIpv4UnicastConnected ---
func packBgpRedistributionProfilesIpv4UnicastConnectedFromSdk(ctx context.Context, sdk network_services.BgpRedistributionProfilesIpv4UnicastConnected) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpRedistributionProfilesIpv4UnicastConnected", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpRedistributionProfilesIpv4UnicastConnected
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Enable != nil {
		model.Enable = basetypes.NewBoolValue(*sdk.Enable)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Enable", "value": *sdk.Enable})
	} else {
		model.Enable = basetypes.NewBoolNull()
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
	if sdk.RouteMap != nil {
		model.RouteMap = basetypes.NewStringValue(*sdk.RouteMap)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "RouteMap", "value": *sdk.RouteMap})
	} else {
		model.RouteMap = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BgpRedistributionProfilesIpv4UnicastConnected{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpRedistributionProfilesIpv4UnicastConnected", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpRedistributionProfilesIpv4UnicastConnected ---
func unpackBgpRedistributionProfilesIpv4UnicastConnectedListToSdk(ctx context.Context, list types.List) ([]network_services.BgpRedistributionProfilesIpv4UnicastConnected, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpRedistributionProfilesIpv4UnicastConnected")
	diags := diag.Diagnostics{}
	var data []models.BgpRedistributionProfilesIpv4UnicastConnected
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpRedistributionProfilesIpv4UnicastConnected, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpRedistributionProfilesIpv4UnicastConnected{}.AttrTypes(), &item)
		unpacked, d := unpackBgpRedistributionProfilesIpv4UnicastConnectedToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpRedistributionProfilesIpv4UnicastConnected", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpRedistributionProfilesIpv4UnicastConnected ---
func packBgpRedistributionProfilesIpv4UnicastConnectedListFromSdk(ctx context.Context, sdks []network_services.BgpRedistributionProfilesIpv4UnicastConnected) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpRedistributionProfilesIpv4UnicastConnected")
	diags := diag.Diagnostics{}
	var data []models.BgpRedistributionProfilesIpv4UnicastConnected

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpRedistributionProfilesIpv4UnicastConnected
		obj, d := packBgpRedistributionProfilesIpv4UnicastConnectedFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpRedistributionProfilesIpv4UnicastConnected{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpRedistributionProfilesIpv4UnicastConnected", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpRedistributionProfilesIpv4UnicastConnected{}.AttrType(), data)
}

// --- Unpacker for BgpRedistributionProfilesIpv4UnicastOspf ---
func unpackBgpRedistributionProfilesIpv4UnicastOspfToSdk(ctx context.Context, obj types.Object) (*network_services.BgpRedistributionProfilesIpv4UnicastOspf, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpRedistributionProfilesIpv4UnicastOspf", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpRedistributionProfilesIpv4UnicastOspf
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpRedistributionProfilesIpv4UnicastOspf
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Enable.IsNull() && !model.Enable.IsUnknown() {
		sdk.Enable = model.Enable.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Enable", "value": *sdk.Enable})
	}

	// Handling Primitives
	if !model.Metric.IsNull() && !model.Metric.IsUnknown() {
		val := int32(model.Metric.ValueInt64())
		sdk.Metric = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Metric", "value": *sdk.Metric})
	}

	// Handling Primitives
	if !model.RouteMap.IsNull() && !model.RouteMap.IsUnknown() {
		sdk.RouteMap = model.RouteMap.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "RouteMap", "value": *sdk.RouteMap})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpRedistributionProfilesIpv4UnicastOspf", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpRedistributionProfilesIpv4UnicastOspf ---
func packBgpRedistributionProfilesIpv4UnicastOspfFromSdk(ctx context.Context, sdk network_services.BgpRedistributionProfilesIpv4UnicastOspf) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpRedistributionProfilesIpv4UnicastOspf", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpRedistributionProfilesIpv4UnicastOspf
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Enable != nil {
		model.Enable = basetypes.NewBoolValue(*sdk.Enable)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Enable", "value": *sdk.Enable})
	} else {
		model.Enable = basetypes.NewBoolNull()
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
	if sdk.RouteMap != nil {
		model.RouteMap = basetypes.NewStringValue(*sdk.RouteMap)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "RouteMap", "value": *sdk.RouteMap})
	} else {
		model.RouteMap = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BgpRedistributionProfilesIpv4UnicastOspf{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpRedistributionProfilesIpv4UnicastOspf", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpRedistributionProfilesIpv4UnicastOspf ---
func unpackBgpRedistributionProfilesIpv4UnicastOspfListToSdk(ctx context.Context, list types.List) ([]network_services.BgpRedistributionProfilesIpv4UnicastOspf, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpRedistributionProfilesIpv4UnicastOspf")
	diags := diag.Diagnostics{}
	var data []models.BgpRedistributionProfilesIpv4UnicastOspf
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpRedistributionProfilesIpv4UnicastOspf, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpRedistributionProfilesIpv4UnicastOspf{}.AttrTypes(), &item)
		unpacked, d := unpackBgpRedistributionProfilesIpv4UnicastOspfToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpRedistributionProfilesIpv4UnicastOspf", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpRedistributionProfilesIpv4UnicastOspf ---
func packBgpRedistributionProfilesIpv4UnicastOspfListFromSdk(ctx context.Context, sdks []network_services.BgpRedistributionProfilesIpv4UnicastOspf) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpRedistributionProfilesIpv4UnicastOspf")
	diags := diag.Diagnostics{}
	var data []models.BgpRedistributionProfilesIpv4UnicastOspf

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpRedistributionProfilesIpv4UnicastOspf
		obj, d := packBgpRedistributionProfilesIpv4UnicastOspfFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpRedistributionProfilesIpv4UnicastOspf{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpRedistributionProfilesIpv4UnicastOspf", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpRedistributionProfilesIpv4UnicastOspf{}.AttrType(), data)
}

// --- Unpacker for BgpRedistributionProfilesIpv4UnicastStatic ---
func unpackBgpRedistributionProfilesIpv4UnicastStaticToSdk(ctx context.Context, obj types.Object) (*network_services.BgpRedistributionProfilesIpv4UnicastStatic, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpRedistributionProfilesIpv4UnicastStatic", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpRedistributionProfilesIpv4UnicastStatic
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpRedistributionProfilesIpv4UnicastStatic
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Enable.IsNull() && !model.Enable.IsUnknown() {
		sdk.Enable = model.Enable.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Enable", "value": *sdk.Enable})
	}

	// Handling Primitives
	if !model.Metric.IsNull() && !model.Metric.IsUnknown() {
		val := int32(model.Metric.ValueInt64())
		sdk.Metric = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Metric", "value": *sdk.Metric})
	}

	// Handling Primitives
	if !model.RouteMap.IsNull() && !model.RouteMap.IsUnknown() {
		sdk.RouteMap = model.RouteMap.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "RouteMap", "value": *sdk.RouteMap})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpRedistributionProfilesIpv4UnicastStatic", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpRedistributionProfilesIpv4UnicastStatic ---
func packBgpRedistributionProfilesIpv4UnicastStaticFromSdk(ctx context.Context, sdk network_services.BgpRedistributionProfilesIpv4UnicastStatic) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpRedistributionProfilesIpv4UnicastStatic", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpRedistributionProfilesIpv4UnicastStatic
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Enable != nil {
		model.Enable = basetypes.NewBoolValue(*sdk.Enable)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Enable", "value": *sdk.Enable})
	} else {
		model.Enable = basetypes.NewBoolNull()
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
	if sdk.RouteMap != nil {
		model.RouteMap = basetypes.NewStringValue(*sdk.RouteMap)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "RouteMap", "value": *sdk.RouteMap})
	} else {
		model.RouteMap = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BgpRedistributionProfilesIpv4UnicastStatic{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpRedistributionProfilesIpv4UnicastStatic", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpRedistributionProfilesIpv4UnicastStatic ---
func unpackBgpRedistributionProfilesIpv4UnicastStaticListToSdk(ctx context.Context, list types.List) ([]network_services.BgpRedistributionProfilesIpv4UnicastStatic, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpRedistributionProfilesIpv4UnicastStatic")
	diags := diag.Diagnostics{}
	var data []models.BgpRedistributionProfilesIpv4UnicastStatic
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpRedistributionProfilesIpv4UnicastStatic, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpRedistributionProfilesIpv4UnicastStatic{}.AttrTypes(), &item)
		unpacked, d := unpackBgpRedistributionProfilesIpv4UnicastStaticToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpRedistributionProfilesIpv4UnicastStatic", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpRedistributionProfilesIpv4UnicastStatic ---
func packBgpRedistributionProfilesIpv4UnicastStaticListFromSdk(ctx context.Context, sdks []network_services.BgpRedistributionProfilesIpv4UnicastStatic) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpRedistributionProfilesIpv4UnicastStatic")
	diags := diag.Diagnostics{}
	var data []models.BgpRedistributionProfilesIpv4UnicastStatic

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpRedistributionProfilesIpv4UnicastStatic
		obj, d := packBgpRedistributionProfilesIpv4UnicastStaticFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpRedistributionProfilesIpv4UnicastStatic{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpRedistributionProfilesIpv4UnicastStatic", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpRedistributionProfilesIpv4UnicastStatic{}.AttrType(), data)
}
