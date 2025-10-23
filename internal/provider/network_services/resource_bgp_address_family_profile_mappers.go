package provider

import (
	"context"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/network_services"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/network_services"
)

// --- Unpacker for BgpAddressFamilyProfiles ---
func unpackBgpAddressFamilyProfilesToSdk(ctx context.Context, obj types.Object) (*network_services.BgpAddressFamilyProfiles, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpAddressFamilyProfiles", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpAddressFamilyProfiles
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpAddressFamilyProfiles
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
		unpacked, d := unpackBgpAddressFamilyProfilesIpv4ToSdk(ctx, model.Ipv4)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Ipv4"})
		}
		if unpacked != nil {
			sdk.Ipv4 = unpacked
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

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpAddressFamilyProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpAddressFamilyProfiles ---
func packBgpAddressFamilyProfilesFromSdk(ctx context.Context, sdk network_services.BgpAddressFamilyProfiles) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpAddressFamilyProfiles", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpAddressFamilyProfiles
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
	if sdk.Ipv4 != nil {
		tflog.Debug(ctx, "Packing nested object for field Ipv4")
		packed, d := packBgpAddressFamilyProfilesIpv4FromSdk(ctx, *sdk.Ipv4)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Ipv4"})
		}
		model.Ipv4 = packed
	} else {
		model.Ipv4 = basetypes.NewObjectNull(models.BgpAddressFamilyProfilesIpv4{}.AttrTypes())
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

	obj, d := types.ObjectValueFrom(ctx, models.BgpAddressFamilyProfiles{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpAddressFamilyProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpAddressFamilyProfiles ---
func unpackBgpAddressFamilyProfilesListToSdk(ctx context.Context, list types.List) ([]network_services.BgpAddressFamilyProfiles, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpAddressFamilyProfiles")
	diags := diag.Diagnostics{}
	var data []models.BgpAddressFamilyProfiles
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpAddressFamilyProfiles, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpAddressFamilyProfiles{}.AttrTypes(), &item)
		unpacked, d := unpackBgpAddressFamilyProfilesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpAddressFamilyProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpAddressFamilyProfiles ---
func packBgpAddressFamilyProfilesListFromSdk(ctx context.Context, sdks []network_services.BgpAddressFamilyProfiles) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpAddressFamilyProfiles")
	diags := diag.Diagnostics{}
	var data []models.BgpAddressFamilyProfiles

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpAddressFamilyProfiles
		obj, d := packBgpAddressFamilyProfilesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpAddressFamilyProfiles{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpAddressFamilyProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpAddressFamilyProfiles{}.AttrType(), data)
}

// --- Unpacker for BgpAddressFamilyProfilesIpv4 ---
func unpackBgpAddressFamilyProfilesIpv4ToSdk(ctx context.Context, obj types.Object) (*network_services.BgpAddressFamilyProfilesIpv4, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpAddressFamilyProfilesIpv4", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpAddressFamilyProfilesIpv4
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpAddressFamilyProfilesIpv4
	var d diag.Diagnostics
	// Handling Objects
	if !model.Ipv4.IsNull() && !model.Ipv4.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Ipv4")
		unpacked, d := unpackBgpAddressFamilyProfilesIpv4Ipv4ToSdk(ctx, model.Ipv4)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Ipv4"})
		}
		if unpacked != nil {
			sdk.Ipv4 = *unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpAddressFamilyProfilesIpv4", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpAddressFamilyProfilesIpv4 ---
func packBgpAddressFamilyProfilesIpv4FromSdk(ctx context.Context, sdk network_services.BgpAddressFamilyProfilesIpv4) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpAddressFamilyProfilesIpv4", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpAddressFamilyProfilesIpv4
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	// Logic for non-pointer / value-type nested objects
	if !reflect.ValueOf(sdk.Ipv4).IsZero() {
		tflog.Debug(ctx, "Packing nested object for field Ipv4")
		packed, d := packBgpAddressFamilyProfilesIpv4Ipv4FromSdk(ctx, sdk.Ipv4)
		diags.Append(d...)
		model.Ipv4 = packed
	} else {
		model.Ipv4 = basetypes.NewObjectNull(models.BgpAddressFamilyProfilesIpv4Ipv4{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BgpAddressFamilyProfilesIpv4{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpAddressFamilyProfilesIpv4", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpAddressFamilyProfilesIpv4 ---
func unpackBgpAddressFamilyProfilesIpv4ListToSdk(ctx context.Context, list types.List) ([]network_services.BgpAddressFamilyProfilesIpv4, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpAddressFamilyProfilesIpv4")
	diags := diag.Diagnostics{}
	var data []models.BgpAddressFamilyProfilesIpv4
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpAddressFamilyProfilesIpv4, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpAddressFamilyProfilesIpv4{}.AttrTypes(), &item)
		unpacked, d := unpackBgpAddressFamilyProfilesIpv4ToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpAddressFamilyProfilesIpv4", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpAddressFamilyProfilesIpv4 ---
func packBgpAddressFamilyProfilesIpv4ListFromSdk(ctx context.Context, sdks []network_services.BgpAddressFamilyProfilesIpv4) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpAddressFamilyProfilesIpv4")
	diags := diag.Diagnostics{}
	var data []models.BgpAddressFamilyProfilesIpv4

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpAddressFamilyProfilesIpv4
		obj, d := packBgpAddressFamilyProfilesIpv4FromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpAddressFamilyProfilesIpv4{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpAddressFamilyProfilesIpv4", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpAddressFamilyProfilesIpv4{}.AttrType(), data)
}

// --- Unpacker for BgpAddressFamilyProfilesIpv4Ipv4 ---
func unpackBgpAddressFamilyProfilesIpv4Ipv4ToSdk(ctx context.Context, obj types.Object) (*network_services.BgpAddressFamilyProfilesIpv4Ipv4, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpAddressFamilyProfilesIpv4Ipv4", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpAddressFamilyProfilesIpv4Ipv4
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpAddressFamilyProfilesIpv4Ipv4
	var d diag.Diagnostics
	// Handling Objects
	if !model.Multicast.IsNull() && !model.Multicast.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Multicast")
		unpacked, d := unpackBgpAddressFamilyToSdk(ctx, model.Multicast)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Multicast"})
		}
		if unpacked != nil {
			sdk.Multicast = unpacked
		}
	}

	// Handling Objects
	if !model.Unicast.IsNull() && !model.Unicast.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Unicast")
		unpacked, d := unpackBgpAddressFamilyToSdk(ctx, model.Unicast)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Unicast"})
		}
		if unpacked != nil {
			sdk.Unicast = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpAddressFamilyProfilesIpv4Ipv4", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpAddressFamilyProfilesIpv4Ipv4 ---
func packBgpAddressFamilyProfilesIpv4Ipv4FromSdk(ctx context.Context, sdk network_services.BgpAddressFamilyProfilesIpv4Ipv4) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpAddressFamilyProfilesIpv4Ipv4", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpAddressFamilyProfilesIpv4Ipv4
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Multicast != nil {
		tflog.Debug(ctx, "Packing nested object for field Multicast")
		packed, d := packBgpAddressFamilyFromSdk(ctx, *sdk.Multicast)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Multicast"})
		}
		model.Multicast = packed
	} else {
		model.Multicast = basetypes.NewObjectNull(models.BgpAddressFamily{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Unicast != nil {
		tflog.Debug(ctx, "Packing nested object for field Unicast")
		packed, d := packBgpAddressFamilyFromSdk(ctx, *sdk.Unicast)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Unicast"})
		}
		model.Unicast = packed
	} else {
		model.Unicast = basetypes.NewObjectNull(models.BgpAddressFamily{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BgpAddressFamilyProfilesIpv4Ipv4{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpAddressFamilyProfilesIpv4Ipv4", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpAddressFamilyProfilesIpv4Ipv4 ---
func unpackBgpAddressFamilyProfilesIpv4Ipv4ListToSdk(ctx context.Context, list types.List) ([]network_services.BgpAddressFamilyProfilesIpv4Ipv4, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpAddressFamilyProfilesIpv4Ipv4")
	diags := diag.Diagnostics{}
	var data []models.BgpAddressFamilyProfilesIpv4Ipv4
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpAddressFamilyProfilesIpv4Ipv4, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpAddressFamilyProfilesIpv4Ipv4{}.AttrTypes(), &item)
		unpacked, d := unpackBgpAddressFamilyProfilesIpv4Ipv4ToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpAddressFamilyProfilesIpv4Ipv4", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpAddressFamilyProfilesIpv4Ipv4 ---
func packBgpAddressFamilyProfilesIpv4Ipv4ListFromSdk(ctx context.Context, sdks []network_services.BgpAddressFamilyProfilesIpv4Ipv4) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpAddressFamilyProfilesIpv4Ipv4")
	diags := diag.Diagnostics{}
	var data []models.BgpAddressFamilyProfilesIpv4Ipv4

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpAddressFamilyProfilesIpv4Ipv4
		obj, d := packBgpAddressFamilyProfilesIpv4Ipv4FromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpAddressFamilyProfilesIpv4Ipv4{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpAddressFamilyProfilesIpv4Ipv4", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpAddressFamilyProfilesIpv4Ipv4{}.AttrType(), data)
}

// --- Unpacker for BgpAddressFamily ---
func unpackBgpAddressFamilyToSdk(ctx context.Context, obj types.Object) (*network_services.BgpAddressFamily, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpAddressFamily", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpAddressFamily
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpAddressFamily
	var d diag.Diagnostics
	// Handling Objects
	if !model.AddPath.IsNull() && !model.AddPath.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field AddPath")
		unpacked, d := unpackBgpAddressFamilyAddPathToSdk(ctx, model.AddPath)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "AddPath"})
		}
		if unpacked != nil {
			sdk.AddPath = unpacked
		}
	}

	// Handling Objects
	if !model.AllowasIn.IsNull() && !model.AllowasIn.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field AllowasIn")
		unpacked, d := unpackBgpAddressFamilyAllowasInToSdk(ctx, model.AllowasIn)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "AllowasIn"})
		}
		if unpacked != nil {
			sdk.AllowasIn = unpacked
		}
	}

	// Handling Primitives
	if !model.AsOverride.IsNull() && !model.AsOverride.IsUnknown() {
		sdk.AsOverride = model.AsOverride.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "AsOverride", "value": *sdk.AsOverride})
	}

	// Handling Primitives
	if !model.DefaultOriginate.IsNull() && !model.DefaultOriginate.IsUnknown() {
		sdk.DefaultOriginate = model.DefaultOriginate.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DefaultOriginate", "value": *sdk.DefaultOriginate})
	}

	// Handling Primitives
	if !model.DefaultOriginateMap.IsNull() && !model.DefaultOriginateMap.IsUnknown() {
		sdk.DefaultOriginateMap = model.DefaultOriginateMap.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DefaultOriginateMap", "value": *sdk.DefaultOriginateMap})
	}

	// Handling Primitives
	if !model.Enable.IsNull() && !model.Enable.IsUnknown() {
		sdk.Enable = model.Enable.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Enable", "value": *sdk.Enable})
	}

	// Handling Objects
	if !model.MaximumPrefix.IsNull() && !model.MaximumPrefix.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field MaximumPrefix")
		unpacked, d := unpackBgpAddressFamilyMaximumPrefixToSdk(ctx, model.MaximumPrefix)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "MaximumPrefix"})
		}
		if unpacked != nil {
			sdk.MaximumPrefix = unpacked
		}
	}

	// Handling Objects
	if !model.NextHop.IsNull() && !model.NextHop.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field NextHop")
		unpacked, d := unpackBgpAddressFamilyNextHopToSdk(ctx, model.NextHop)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "NextHop"})
		}
		if unpacked != nil {
			sdk.NextHop = unpacked
		}
	}

	// Handling Objects
	if !model.Orf.IsNull() && !model.Orf.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Orf")
		unpacked, d := unpackBgpAddressFamilyOrfToSdk(ctx, model.Orf)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Orf"})
		}
		if unpacked != nil {
			sdk.Orf = unpacked
		}
	}

	// Handling Objects
	if !model.RemovePrivateAS.IsNull() && !model.RemovePrivateAS.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field RemovePrivateAS")
		unpacked, d := unpackBgpAddressFamilyRemovePrivateASToSdk(ctx, model.RemovePrivateAS)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "RemovePrivateAS"})
		}
		if unpacked != nil {
			sdk.RemovePrivateAS = unpacked
		}
	}

	// Handling Primitives
	if !model.RouteReflectorClient.IsNull() && !model.RouteReflectorClient.IsUnknown() {
		sdk.RouteReflectorClient = model.RouteReflectorClient.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "RouteReflectorClient", "value": *sdk.RouteReflectorClient})
	}

	// Handling Objects
	if !model.SendCommunity.IsNull() && !model.SendCommunity.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field SendCommunity")
		unpacked, d := unpackBgpAddressFamilySendCommunityToSdk(ctx, model.SendCommunity)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "SendCommunity"})
		}
		if unpacked != nil {
			sdk.SendCommunity = unpacked
		}
	}

	// Handling Primitives
	if !model.SoftReconfigWithStoredInfo.IsNull() && !model.SoftReconfigWithStoredInfo.IsUnknown() {
		sdk.SoftReconfigWithStoredInfo = model.SoftReconfigWithStoredInfo.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "SoftReconfigWithStoredInfo", "value": *sdk.SoftReconfigWithStoredInfo})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpAddressFamily", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpAddressFamily ---
func packBgpAddressFamilyFromSdk(ctx context.Context, sdk network_services.BgpAddressFamily) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpAddressFamily", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpAddressFamily
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.AddPath != nil {
		tflog.Debug(ctx, "Packing nested object for field AddPath")
		packed, d := packBgpAddressFamilyAddPathFromSdk(ctx, *sdk.AddPath)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "AddPath"})
		}
		model.AddPath = packed
	} else {
		model.AddPath = basetypes.NewObjectNull(models.BgpAddressFamilyAddPath{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.AllowasIn != nil {
		tflog.Debug(ctx, "Packing nested object for field AllowasIn")
		packed, d := packBgpAddressFamilyAllowasInFromSdk(ctx, *sdk.AllowasIn)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "AllowasIn"})
		}
		model.AllowasIn = packed
	} else {
		model.AllowasIn = basetypes.NewObjectNull(models.BgpAddressFamilyAllowasIn{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.AsOverride != nil {
		model.AsOverride = basetypes.NewBoolValue(*sdk.AsOverride)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "AsOverride", "value": *sdk.AsOverride})
	} else {
		model.AsOverride = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.DefaultOriginate != nil {
		model.DefaultOriginate = basetypes.NewBoolValue(*sdk.DefaultOriginate)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DefaultOriginate", "value": *sdk.DefaultOriginate})
	} else {
		model.DefaultOriginate = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.DefaultOriginateMap != nil {
		model.DefaultOriginateMap = basetypes.NewStringValue(*sdk.DefaultOriginateMap)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DefaultOriginateMap", "value": *sdk.DefaultOriginateMap})
	} else {
		model.DefaultOriginateMap = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Enable != nil {
		model.Enable = basetypes.NewBoolValue(*sdk.Enable)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Enable", "value": *sdk.Enable})
	} else {
		model.Enable = basetypes.NewBoolNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.MaximumPrefix != nil {
		tflog.Debug(ctx, "Packing nested object for field MaximumPrefix")
		packed, d := packBgpAddressFamilyMaximumPrefixFromSdk(ctx, *sdk.MaximumPrefix)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "MaximumPrefix"})
		}
		model.MaximumPrefix = packed
	} else {
		model.MaximumPrefix = basetypes.NewObjectNull(models.BgpAddressFamilyMaximumPrefix{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.NextHop != nil {
		tflog.Debug(ctx, "Packing nested object for field NextHop")
		packed, d := packBgpAddressFamilyNextHopFromSdk(ctx, *sdk.NextHop)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "NextHop"})
		}
		model.NextHop = packed
	} else {
		model.NextHop = basetypes.NewObjectNull(models.BgpAddressFamilyNextHop{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Orf != nil {
		tflog.Debug(ctx, "Packing nested object for field Orf")
		packed, d := packBgpAddressFamilyOrfFromSdk(ctx, *sdk.Orf)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Orf"})
		}
		model.Orf = packed
	} else {
		model.Orf = basetypes.NewObjectNull(models.BgpAddressFamilyOrf{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.RemovePrivateAS != nil {
		tflog.Debug(ctx, "Packing nested object for field RemovePrivateAS")
		packed, d := packBgpAddressFamilyRemovePrivateASFromSdk(ctx, *sdk.RemovePrivateAS)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "RemovePrivateAS"})
		}
		model.RemovePrivateAS = packed
	} else {
		model.RemovePrivateAS = basetypes.NewObjectNull(models.BgpAddressFamilyRemovePrivateAS{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.RouteReflectorClient != nil {
		model.RouteReflectorClient = basetypes.NewBoolValue(*sdk.RouteReflectorClient)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "RouteReflectorClient", "value": *sdk.RouteReflectorClient})
	} else {
		model.RouteReflectorClient = basetypes.NewBoolNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.SendCommunity != nil {
		tflog.Debug(ctx, "Packing nested object for field SendCommunity")
		packed, d := packBgpAddressFamilySendCommunityFromSdk(ctx, *sdk.SendCommunity)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "SendCommunity"})
		}
		model.SendCommunity = packed
	} else {
		model.SendCommunity = basetypes.NewObjectNull(models.BgpAddressFamilySendCommunity{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.SoftReconfigWithStoredInfo != nil {
		model.SoftReconfigWithStoredInfo = basetypes.NewBoolValue(*sdk.SoftReconfigWithStoredInfo)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "SoftReconfigWithStoredInfo", "value": *sdk.SoftReconfigWithStoredInfo})
	} else {
		model.SoftReconfigWithStoredInfo = basetypes.NewBoolNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BgpAddressFamily{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpAddressFamily", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpAddressFamily ---
func unpackBgpAddressFamilyListToSdk(ctx context.Context, list types.List) ([]network_services.BgpAddressFamily, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpAddressFamily")
	diags := diag.Diagnostics{}
	var data []models.BgpAddressFamily
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpAddressFamily, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpAddressFamily{}.AttrTypes(), &item)
		unpacked, d := unpackBgpAddressFamilyToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpAddressFamily", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpAddressFamily ---
func packBgpAddressFamilyListFromSdk(ctx context.Context, sdks []network_services.BgpAddressFamily) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpAddressFamily")
	diags := diag.Diagnostics{}
	var data []models.BgpAddressFamily

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpAddressFamily
		obj, d := packBgpAddressFamilyFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpAddressFamily{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpAddressFamily", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpAddressFamily{}.AttrType(), data)
}

// --- Unpacker for BgpAddressFamilyAddPath ---
func unpackBgpAddressFamilyAddPathToSdk(ctx context.Context, obj types.Object) (*network_services.BgpAddressFamilyAddPath, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpAddressFamilyAddPath", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpAddressFamilyAddPath
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpAddressFamilyAddPath
	var d diag.Diagnostics
	// Handling Primitives
	if !model.TxAllPaths.IsNull() && !model.TxAllPaths.IsUnknown() {
		sdk.TxAllPaths = model.TxAllPaths.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "TxAllPaths", "value": *sdk.TxAllPaths})
	}

	// Handling Primitives
	if !model.TxBestpathPerAS.IsNull() && !model.TxBestpathPerAS.IsUnknown() {
		sdk.TxBestpathPerAS = model.TxBestpathPerAS.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "TxBestpathPerAS", "value": *sdk.TxBestpathPerAS})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpAddressFamilyAddPath", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpAddressFamilyAddPath ---
func packBgpAddressFamilyAddPathFromSdk(ctx context.Context, sdk network_services.BgpAddressFamilyAddPath) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpAddressFamilyAddPath", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpAddressFamilyAddPath
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.TxAllPaths != nil {
		model.TxAllPaths = basetypes.NewBoolValue(*sdk.TxAllPaths)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "TxAllPaths", "value": *sdk.TxAllPaths})
	} else {
		model.TxAllPaths = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.TxBestpathPerAS != nil {
		model.TxBestpathPerAS = basetypes.NewBoolValue(*sdk.TxBestpathPerAS)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "TxBestpathPerAS", "value": *sdk.TxBestpathPerAS})
	} else {
		model.TxBestpathPerAS = basetypes.NewBoolNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BgpAddressFamilyAddPath{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpAddressFamilyAddPath", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpAddressFamilyAddPath ---
func unpackBgpAddressFamilyAddPathListToSdk(ctx context.Context, list types.List) ([]network_services.BgpAddressFamilyAddPath, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpAddressFamilyAddPath")
	diags := diag.Diagnostics{}
	var data []models.BgpAddressFamilyAddPath
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpAddressFamilyAddPath, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpAddressFamilyAddPath{}.AttrTypes(), &item)
		unpacked, d := unpackBgpAddressFamilyAddPathToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpAddressFamilyAddPath", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpAddressFamilyAddPath ---
func packBgpAddressFamilyAddPathListFromSdk(ctx context.Context, sdks []network_services.BgpAddressFamilyAddPath) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpAddressFamilyAddPath")
	diags := diag.Diagnostics{}
	var data []models.BgpAddressFamilyAddPath

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpAddressFamilyAddPath
		obj, d := packBgpAddressFamilyAddPathFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpAddressFamilyAddPath{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpAddressFamilyAddPath", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpAddressFamilyAddPath{}.AttrType(), data)
}

// --- Unpacker for BgpAddressFamilyAllowasIn ---
func unpackBgpAddressFamilyAllowasInToSdk(ctx context.Context, obj types.Object) (*network_services.BgpAddressFamilyAllowasIn, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpAddressFamilyAllowasIn", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpAddressFamilyAllowasIn
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpAddressFamilyAllowasIn
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Occurrence.IsNull() && !model.Occurrence.IsUnknown() {
		val := int32(model.Occurrence.ValueInt64())
		sdk.Occurrence = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Occurrence", "value": *sdk.Occurrence})
	}

	// Handling Typeless Objects
	if !model.Origin.IsNull() && !model.Origin.IsUnknown() {
		tflog.Debug(ctx, "Unpacking typeless object for field Origin")
		sdk.Origin = make(map[string]interface{})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpAddressFamilyAllowasIn", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpAddressFamilyAllowasIn ---
func packBgpAddressFamilyAllowasInFromSdk(ctx context.Context, sdk network_services.BgpAddressFamilyAllowasIn) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpAddressFamilyAllowasIn", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpAddressFamilyAllowasIn
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Occurrence != nil {
		model.Occurrence = basetypes.NewInt64Value(int64(*sdk.Occurrence))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Occurrence", "value": *sdk.Occurrence})
	} else {
		model.Occurrence = basetypes.NewInt64Null()
	}
	// Handling Objects
	// This is a marker object (e.g. CHAP: {}). We just need to create an empty, non-null object.
	if sdk.Origin != nil && !reflect.ValueOf(sdk.Origin).IsNil() {
		tflog.Debug(ctx, "Packing typeless object for field Origin")
		var d diag.Diagnostics
		// Create an empty object with no attributes, which signifies its presence.
		model.Origin, d = basetypes.NewObjectValue(map[string]attr.Type{}, map[string]attr.Value{})
		diags.Append(d...)
	} else {
		// Since this field is part of a oneOf, being nil means it's not selected.
		// We make the object null with an empty attribute map.
		model.Origin = basetypes.NewObjectNull(map[string]attr.Type{})
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BgpAddressFamilyAllowasIn{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpAddressFamilyAllowasIn", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpAddressFamilyAllowasIn ---
func unpackBgpAddressFamilyAllowasInListToSdk(ctx context.Context, list types.List) ([]network_services.BgpAddressFamilyAllowasIn, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpAddressFamilyAllowasIn")
	diags := diag.Diagnostics{}
	var data []models.BgpAddressFamilyAllowasIn
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpAddressFamilyAllowasIn, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpAddressFamilyAllowasIn{}.AttrTypes(), &item)
		unpacked, d := unpackBgpAddressFamilyAllowasInToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpAddressFamilyAllowasIn", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpAddressFamilyAllowasIn ---
func packBgpAddressFamilyAllowasInListFromSdk(ctx context.Context, sdks []network_services.BgpAddressFamilyAllowasIn) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpAddressFamilyAllowasIn")
	diags := diag.Diagnostics{}
	var data []models.BgpAddressFamilyAllowasIn

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpAddressFamilyAllowasIn
		obj, d := packBgpAddressFamilyAllowasInFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpAddressFamilyAllowasIn{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpAddressFamilyAllowasIn", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpAddressFamilyAllowasIn{}.AttrType(), data)
}

// --- Unpacker for BgpAddressFamilyMaximumPrefix ---
func unpackBgpAddressFamilyMaximumPrefixToSdk(ctx context.Context, obj types.Object) (*network_services.BgpAddressFamilyMaximumPrefix, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpAddressFamilyMaximumPrefix", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpAddressFamilyMaximumPrefix
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpAddressFamilyMaximumPrefix
	var d diag.Diagnostics
	// Handling Objects
	if !model.Action.IsNull() && !model.Action.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Action")
		unpacked, d := unpackBgpAddressFamilyMaximumPrefixActionToSdk(ctx, model.Action)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Action"})
		}
		if unpacked != nil {
			sdk.Action = unpacked
		}
	}

	// Handling Primitives
	if !model.NumPrefixes.IsNull() && !model.NumPrefixes.IsUnknown() {
		val := int32(model.NumPrefixes.ValueInt64())
		sdk.NumPrefixes = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "NumPrefixes", "value": *sdk.NumPrefixes})
	}

	// Handling Primitives
	if !model.Threshold.IsNull() && !model.Threshold.IsUnknown() {
		val := int32(model.Threshold.ValueInt64())
		sdk.Threshold = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Threshold", "value": *sdk.Threshold})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpAddressFamilyMaximumPrefix", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpAddressFamilyMaximumPrefix ---
func packBgpAddressFamilyMaximumPrefixFromSdk(ctx context.Context, sdk network_services.BgpAddressFamilyMaximumPrefix) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpAddressFamilyMaximumPrefix", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpAddressFamilyMaximumPrefix
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Action != nil {
		tflog.Debug(ctx, "Packing nested object for field Action")
		packed, d := packBgpAddressFamilyMaximumPrefixActionFromSdk(ctx, *sdk.Action)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Action"})
		}
		model.Action = packed
	} else {
		model.Action = basetypes.NewObjectNull(models.BgpAddressFamilyMaximumPrefixAction{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.NumPrefixes != nil {
		model.NumPrefixes = basetypes.NewInt64Value(int64(*sdk.NumPrefixes))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "NumPrefixes", "value": *sdk.NumPrefixes})
	} else {
		model.NumPrefixes = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Threshold != nil {
		model.Threshold = basetypes.NewInt64Value(int64(*sdk.Threshold))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Threshold", "value": *sdk.Threshold})
	} else {
		model.Threshold = basetypes.NewInt64Null()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BgpAddressFamilyMaximumPrefix{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpAddressFamilyMaximumPrefix", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpAddressFamilyMaximumPrefix ---
func unpackBgpAddressFamilyMaximumPrefixListToSdk(ctx context.Context, list types.List) ([]network_services.BgpAddressFamilyMaximumPrefix, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpAddressFamilyMaximumPrefix")
	diags := diag.Diagnostics{}
	var data []models.BgpAddressFamilyMaximumPrefix
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpAddressFamilyMaximumPrefix, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpAddressFamilyMaximumPrefix{}.AttrTypes(), &item)
		unpacked, d := unpackBgpAddressFamilyMaximumPrefixToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpAddressFamilyMaximumPrefix", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpAddressFamilyMaximumPrefix ---
func packBgpAddressFamilyMaximumPrefixListFromSdk(ctx context.Context, sdks []network_services.BgpAddressFamilyMaximumPrefix) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpAddressFamilyMaximumPrefix")
	diags := diag.Diagnostics{}
	var data []models.BgpAddressFamilyMaximumPrefix

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpAddressFamilyMaximumPrefix
		obj, d := packBgpAddressFamilyMaximumPrefixFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpAddressFamilyMaximumPrefix{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpAddressFamilyMaximumPrefix", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpAddressFamilyMaximumPrefix{}.AttrType(), data)
}

// --- Unpacker for BgpAddressFamilyMaximumPrefixAction ---
func unpackBgpAddressFamilyMaximumPrefixActionToSdk(ctx context.Context, obj types.Object) (*network_services.BgpAddressFamilyMaximumPrefixAction, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpAddressFamilyMaximumPrefixAction", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpAddressFamilyMaximumPrefixAction
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpAddressFamilyMaximumPrefixAction
	var d diag.Diagnostics
	// Handling Objects
	if !model.Restart.IsNull() && !model.Restart.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Restart")
		unpacked, d := unpackBgpAddressFamilyMaximumPrefixActionRestartToSdk(ctx, model.Restart)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Restart"})
		}
		if unpacked != nil {
			sdk.Restart = unpacked
		}
	}

	// Handling Typeless Objects
	if !model.WarningOnly.IsNull() && !model.WarningOnly.IsUnknown() {
		tflog.Debug(ctx, "Unpacking typeless object for field WarningOnly")
		sdk.WarningOnly = make(map[string]interface{})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpAddressFamilyMaximumPrefixAction", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpAddressFamilyMaximumPrefixAction ---
func packBgpAddressFamilyMaximumPrefixActionFromSdk(ctx context.Context, sdk network_services.BgpAddressFamilyMaximumPrefixAction) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpAddressFamilyMaximumPrefixAction", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpAddressFamilyMaximumPrefixAction
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Restart != nil {
		tflog.Debug(ctx, "Packing nested object for field Restart")
		packed, d := packBgpAddressFamilyMaximumPrefixActionRestartFromSdk(ctx, *sdk.Restart)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Restart"})
		}
		model.Restart = packed
	} else {
		model.Restart = basetypes.NewObjectNull(models.BgpAddressFamilyMaximumPrefixActionRestart{}.AttrTypes())
	}
	// Handling Objects
	// This is a marker object (e.g. CHAP: {}). We just need to create an empty, non-null object.
	if sdk.WarningOnly != nil && !reflect.ValueOf(sdk.WarningOnly).IsNil() {
		tflog.Debug(ctx, "Packing typeless object for field WarningOnly")
		var d diag.Diagnostics
		// Create an empty object with no attributes, which signifies its presence.
		model.WarningOnly, d = basetypes.NewObjectValue(map[string]attr.Type{}, map[string]attr.Value{})
		diags.Append(d...)
	} else {
		// Since this field is part of a oneOf, being nil means it's not selected.
		// We make the object null with an empty attribute map.
		model.WarningOnly = basetypes.NewObjectNull(map[string]attr.Type{})
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BgpAddressFamilyMaximumPrefixAction{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpAddressFamilyMaximumPrefixAction", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpAddressFamilyMaximumPrefixAction ---
func unpackBgpAddressFamilyMaximumPrefixActionListToSdk(ctx context.Context, list types.List) ([]network_services.BgpAddressFamilyMaximumPrefixAction, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpAddressFamilyMaximumPrefixAction")
	diags := diag.Diagnostics{}
	var data []models.BgpAddressFamilyMaximumPrefixAction
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpAddressFamilyMaximumPrefixAction, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpAddressFamilyMaximumPrefixAction{}.AttrTypes(), &item)
		unpacked, d := unpackBgpAddressFamilyMaximumPrefixActionToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpAddressFamilyMaximumPrefixAction", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpAddressFamilyMaximumPrefixAction ---
func packBgpAddressFamilyMaximumPrefixActionListFromSdk(ctx context.Context, sdks []network_services.BgpAddressFamilyMaximumPrefixAction) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpAddressFamilyMaximumPrefixAction")
	diags := diag.Diagnostics{}
	var data []models.BgpAddressFamilyMaximumPrefixAction

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpAddressFamilyMaximumPrefixAction
		obj, d := packBgpAddressFamilyMaximumPrefixActionFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpAddressFamilyMaximumPrefixAction{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpAddressFamilyMaximumPrefixAction", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpAddressFamilyMaximumPrefixAction{}.AttrType(), data)
}

// --- Unpacker for BgpAddressFamilyMaximumPrefixActionRestart ---
func unpackBgpAddressFamilyMaximumPrefixActionRestartToSdk(ctx context.Context, obj types.Object) (*network_services.BgpAddressFamilyMaximumPrefixActionRestart, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpAddressFamilyMaximumPrefixActionRestart", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpAddressFamilyMaximumPrefixActionRestart
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpAddressFamilyMaximumPrefixActionRestart
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Interval.IsNull() && !model.Interval.IsUnknown() {
		val := int32(model.Interval.ValueInt64())
		sdk.Interval = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Interval", "value": *sdk.Interval})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpAddressFamilyMaximumPrefixActionRestart", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpAddressFamilyMaximumPrefixActionRestart ---
func packBgpAddressFamilyMaximumPrefixActionRestartFromSdk(ctx context.Context, sdk network_services.BgpAddressFamilyMaximumPrefixActionRestart) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpAddressFamilyMaximumPrefixActionRestart", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpAddressFamilyMaximumPrefixActionRestart
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Interval != nil {
		model.Interval = basetypes.NewInt64Value(int64(*sdk.Interval))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Interval", "value": *sdk.Interval})
	} else {
		model.Interval = basetypes.NewInt64Null()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BgpAddressFamilyMaximumPrefixActionRestart{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpAddressFamilyMaximumPrefixActionRestart", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpAddressFamilyMaximumPrefixActionRestart ---
func unpackBgpAddressFamilyMaximumPrefixActionRestartListToSdk(ctx context.Context, list types.List) ([]network_services.BgpAddressFamilyMaximumPrefixActionRestart, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpAddressFamilyMaximumPrefixActionRestart")
	diags := diag.Diagnostics{}
	var data []models.BgpAddressFamilyMaximumPrefixActionRestart
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpAddressFamilyMaximumPrefixActionRestart, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpAddressFamilyMaximumPrefixActionRestart{}.AttrTypes(), &item)
		unpacked, d := unpackBgpAddressFamilyMaximumPrefixActionRestartToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpAddressFamilyMaximumPrefixActionRestart", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpAddressFamilyMaximumPrefixActionRestart ---
func packBgpAddressFamilyMaximumPrefixActionRestartListFromSdk(ctx context.Context, sdks []network_services.BgpAddressFamilyMaximumPrefixActionRestart) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpAddressFamilyMaximumPrefixActionRestart")
	diags := diag.Diagnostics{}
	var data []models.BgpAddressFamilyMaximumPrefixActionRestart

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpAddressFamilyMaximumPrefixActionRestart
		obj, d := packBgpAddressFamilyMaximumPrefixActionRestartFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpAddressFamilyMaximumPrefixActionRestart{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpAddressFamilyMaximumPrefixActionRestart", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpAddressFamilyMaximumPrefixActionRestart{}.AttrType(), data)
}

// --- Unpacker for BgpAddressFamilyNextHop ---
func unpackBgpAddressFamilyNextHopToSdk(ctx context.Context, obj types.Object) (*network_services.BgpAddressFamilyNextHop, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpAddressFamilyNextHop", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpAddressFamilyNextHop
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpAddressFamilyNextHop
	var d diag.Diagnostics
	// Handling Typeless Objects
	if !model.Self.IsNull() && !model.Self.IsUnknown() {
		tflog.Debug(ctx, "Unpacking typeless object for field Self")
		sdk.Self = make(map[string]interface{})
	}

	// Handling Typeless Objects
	if !model.SelfForce.IsNull() && !model.SelfForce.IsUnknown() {
		tflog.Debug(ctx, "Unpacking typeless object for field SelfForce")
		sdk.SelfForce = make(map[string]interface{})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpAddressFamilyNextHop", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpAddressFamilyNextHop ---
func packBgpAddressFamilyNextHopFromSdk(ctx context.Context, sdk network_services.BgpAddressFamilyNextHop) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpAddressFamilyNextHop", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpAddressFamilyNextHop
	var d diag.Diagnostics
	// Handling Objects
	// This is a marker object (e.g. CHAP: {}). We just need to create an empty, non-null object.
	if sdk.Self != nil && !reflect.ValueOf(sdk.Self).IsNil() {
		tflog.Debug(ctx, "Packing typeless object for field Self")
		var d diag.Diagnostics
		// Create an empty object with no attributes, which signifies its presence.
		model.Self, d = basetypes.NewObjectValue(map[string]attr.Type{}, map[string]attr.Value{})
		diags.Append(d...)
	} else {
		// Since this field is part of a oneOf, being nil means it's not selected.
		// We make the object null with an empty attribute map.
		model.Self = basetypes.NewObjectNull(map[string]attr.Type{})
	}
	// Handling Objects
	// This is a marker object (e.g. CHAP: {}). We just need to create an empty, non-null object.
	if sdk.SelfForce != nil && !reflect.ValueOf(sdk.SelfForce).IsNil() {
		tflog.Debug(ctx, "Packing typeless object for field SelfForce")
		var d diag.Diagnostics
		// Create an empty object with no attributes, which signifies its presence.
		model.SelfForce, d = basetypes.NewObjectValue(map[string]attr.Type{}, map[string]attr.Value{})
		diags.Append(d...)
	} else {
		// Since this field is part of a oneOf, being nil means it's not selected.
		// We make the object null with an empty attribute map.
		model.SelfForce = basetypes.NewObjectNull(map[string]attr.Type{})
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BgpAddressFamilyNextHop{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpAddressFamilyNextHop", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpAddressFamilyNextHop ---
func unpackBgpAddressFamilyNextHopListToSdk(ctx context.Context, list types.List) ([]network_services.BgpAddressFamilyNextHop, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpAddressFamilyNextHop")
	diags := diag.Diagnostics{}
	var data []models.BgpAddressFamilyNextHop
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpAddressFamilyNextHop, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpAddressFamilyNextHop{}.AttrTypes(), &item)
		unpacked, d := unpackBgpAddressFamilyNextHopToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpAddressFamilyNextHop", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpAddressFamilyNextHop ---
func packBgpAddressFamilyNextHopListFromSdk(ctx context.Context, sdks []network_services.BgpAddressFamilyNextHop) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpAddressFamilyNextHop")
	diags := diag.Diagnostics{}
	var data []models.BgpAddressFamilyNextHop

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpAddressFamilyNextHop
		obj, d := packBgpAddressFamilyNextHopFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpAddressFamilyNextHop{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpAddressFamilyNextHop", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpAddressFamilyNextHop{}.AttrType(), data)
}

// --- Unpacker for BgpAddressFamilyOrf ---
func unpackBgpAddressFamilyOrfToSdk(ctx context.Context, obj types.Object) (*network_services.BgpAddressFamilyOrf, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpAddressFamilyOrf", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpAddressFamilyOrf
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpAddressFamilyOrf
	var d diag.Diagnostics
	// Handling Primitives
	if !model.OrfPrefixList.IsNull() && !model.OrfPrefixList.IsUnknown() {
		sdk.OrfPrefixList = model.OrfPrefixList.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "OrfPrefixList", "value": *sdk.OrfPrefixList})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpAddressFamilyOrf", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpAddressFamilyOrf ---
func packBgpAddressFamilyOrfFromSdk(ctx context.Context, sdk network_services.BgpAddressFamilyOrf) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpAddressFamilyOrf", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpAddressFamilyOrf
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.OrfPrefixList != nil {
		model.OrfPrefixList = basetypes.NewStringValue(*sdk.OrfPrefixList)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "OrfPrefixList", "value": *sdk.OrfPrefixList})
	} else {
		model.OrfPrefixList = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BgpAddressFamilyOrf{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpAddressFamilyOrf", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpAddressFamilyOrf ---
func unpackBgpAddressFamilyOrfListToSdk(ctx context.Context, list types.List) ([]network_services.BgpAddressFamilyOrf, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpAddressFamilyOrf")
	diags := diag.Diagnostics{}
	var data []models.BgpAddressFamilyOrf
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpAddressFamilyOrf, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpAddressFamilyOrf{}.AttrTypes(), &item)
		unpacked, d := unpackBgpAddressFamilyOrfToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpAddressFamilyOrf", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpAddressFamilyOrf ---
func packBgpAddressFamilyOrfListFromSdk(ctx context.Context, sdks []network_services.BgpAddressFamilyOrf) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpAddressFamilyOrf")
	diags := diag.Diagnostics{}
	var data []models.BgpAddressFamilyOrf

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpAddressFamilyOrf
		obj, d := packBgpAddressFamilyOrfFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpAddressFamilyOrf{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpAddressFamilyOrf", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpAddressFamilyOrf{}.AttrType(), data)
}

// --- Unpacker for BgpAddressFamilyRemovePrivateAS ---
func unpackBgpAddressFamilyRemovePrivateASToSdk(ctx context.Context, obj types.Object) (*network_services.BgpAddressFamilyRemovePrivateAS, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpAddressFamilyRemovePrivateAS", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpAddressFamilyRemovePrivateAS
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpAddressFamilyRemovePrivateAS
	var d diag.Diagnostics
	// Handling Typeless Objects
	if !model.All.IsNull() && !model.All.IsUnknown() {
		tflog.Debug(ctx, "Unpacking typeless object for field All")
		sdk.All = make(map[string]interface{})
	}

	// Handling Typeless Objects
	if !model.ReplaceAS.IsNull() && !model.ReplaceAS.IsUnknown() {
		tflog.Debug(ctx, "Unpacking typeless object for field ReplaceAS")
		sdk.ReplaceAS = make(map[string]interface{})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpAddressFamilyRemovePrivateAS", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpAddressFamilyRemovePrivateAS ---
func packBgpAddressFamilyRemovePrivateASFromSdk(ctx context.Context, sdk network_services.BgpAddressFamilyRemovePrivateAS) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpAddressFamilyRemovePrivateAS", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpAddressFamilyRemovePrivateAS
	var d diag.Diagnostics
	// Handling Objects
	// This is a marker object (e.g. CHAP: {}). We just need to create an empty, non-null object.
	if sdk.All != nil && !reflect.ValueOf(sdk.All).IsNil() {
		tflog.Debug(ctx, "Packing typeless object for field All")
		var d diag.Diagnostics
		// Create an empty object with no attributes, which signifies its presence.
		model.All, d = basetypes.NewObjectValue(map[string]attr.Type{}, map[string]attr.Value{})
		diags.Append(d...)
	} else {
		// Since this field is part of a oneOf, being nil means it's not selected.
		// We make the object null with an empty attribute map.
		model.All = basetypes.NewObjectNull(map[string]attr.Type{})
	}
	// Handling Objects
	// This is a marker object (e.g. CHAP: {}). We just need to create an empty, non-null object.
	if sdk.ReplaceAS != nil && !reflect.ValueOf(sdk.ReplaceAS).IsNil() {
		tflog.Debug(ctx, "Packing typeless object for field ReplaceAS")
		var d diag.Diagnostics
		// Create an empty object with no attributes, which signifies its presence.
		model.ReplaceAS, d = basetypes.NewObjectValue(map[string]attr.Type{}, map[string]attr.Value{})
		diags.Append(d...)
	} else {
		// Since this field is part of a oneOf, being nil means it's not selected.
		// We make the object null with an empty attribute map.
		model.ReplaceAS = basetypes.NewObjectNull(map[string]attr.Type{})
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BgpAddressFamilyRemovePrivateAS{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpAddressFamilyRemovePrivateAS", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpAddressFamilyRemovePrivateAS ---
func unpackBgpAddressFamilyRemovePrivateASListToSdk(ctx context.Context, list types.List) ([]network_services.BgpAddressFamilyRemovePrivateAS, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpAddressFamilyRemovePrivateAS")
	diags := diag.Diagnostics{}
	var data []models.BgpAddressFamilyRemovePrivateAS
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpAddressFamilyRemovePrivateAS, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpAddressFamilyRemovePrivateAS{}.AttrTypes(), &item)
		unpacked, d := unpackBgpAddressFamilyRemovePrivateASToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpAddressFamilyRemovePrivateAS", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpAddressFamilyRemovePrivateAS ---
func packBgpAddressFamilyRemovePrivateASListFromSdk(ctx context.Context, sdks []network_services.BgpAddressFamilyRemovePrivateAS) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpAddressFamilyRemovePrivateAS")
	diags := diag.Diagnostics{}
	var data []models.BgpAddressFamilyRemovePrivateAS

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpAddressFamilyRemovePrivateAS
		obj, d := packBgpAddressFamilyRemovePrivateASFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpAddressFamilyRemovePrivateAS{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpAddressFamilyRemovePrivateAS", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpAddressFamilyRemovePrivateAS{}.AttrType(), data)
}

// --- Unpacker for BgpAddressFamilySendCommunity ---
func unpackBgpAddressFamilySendCommunityToSdk(ctx context.Context, obj types.Object) (*network_services.BgpAddressFamilySendCommunity, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpAddressFamilySendCommunity", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpAddressFamilySendCommunity
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpAddressFamilySendCommunity
	var d diag.Diagnostics
	// Handling Typeless Objects
	if !model.All.IsNull() && !model.All.IsUnknown() {
		tflog.Debug(ctx, "Unpacking typeless object for field All")
		sdk.All = make(map[string]interface{})
	}

	// Handling Typeless Objects
	if !model.Both.IsNull() && !model.Both.IsUnknown() {
		tflog.Debug(ctx, "Unpacking typeless object for field Both")
		sdk.Both = make(map[string]interface{})
	}

	// Handling Typeless Objects
	if !model.Extended.IsNull() && !model.Extended.IsUnknown() {
		tflog.Debug(ctx, "Unpacking typeless object for field Extended")
		sdk.Extended = make(map[string]interface{})
	}

	// Handling Typeless Objects
	if !model.Large.IsNull() && !model.Large.IsUnknown() {
		tflog.Debug(ctx, "Unpacking typeless object for field Large")
		sdk.Large = make(map[string]interface{})
	}

	// Handling Typeless Objects
	if !model.Standard.IsNull() && !model.Standard.IsUnknown() {
		tflog.Debug(ctx, "Unpacking typeless object for field Standard")
		sdk.Standard = make(map[string]interface{})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpAddressFamilySendCommunity", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpAddressFamilySendCommunity ---
func packBgpAddressFamilySendCommunityFromSdk(ctx context.Context, sdk network_services.BgpAddressFamilySendCommunity) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpAddressFamilySendCommunity", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpAddressFamilySendCommunity
	var d diag.Diagnostics
	// Handling Objects
	// This is a marker object (e.g. CHAP: {}). We just need to create an empty, non-null object.
	if sdk.All != nil && !reflect.ValueOf(sdk.All).IsNil() {
		tflog.Debug(ctx, "Packing typeless object for field All")
		var d diag.Diagnostics
		// Create an empty object with no attributes, which signifies its presence.
		model.All, d = basetypes.NewObjectValue(map[string]attr.Type{}, map[string]attr.Value{})
		diags.Append(d...)
	} else {
		// Since this field is part of a oneOf, being nil means it's not selected.
		// We make the object null with an empty attribute map.
		model.All = basetypes.NewObjectNull(map[string]attr.Type{})
	}
	// Handling Objects
	// This is a marker object (e.g. CHAP: {}). We just need to create an empty, non-null object.
	if sdk.Both != nil && !reflect.ValueOf(sdk.Both).IsNil() {
		tflog.Debug(ctx, "Packing typeless object for field Both")
		var d diag.Diagnostics
		// Create an empty object with no attributes, which signifies its presence.
		model.Both, d = basetypes.NewObjectValue(map[string]attr.Type{}, map[string]attr.Value{})
		diags.Append(d...)
	} else {
		// Since this field is part of a oneOf, being nil means it's not selected.
		// We make the object null with an empty attribute map.
		model.Both = basetypes.NewObjectNull(map[string]attr.Type{})
	}
	// Handling Objects
	// This is a marker object (e.g. CHAP: {}). We just need to create an empty, non-null object.
	if sdk.Extended != nil && !reflect.ValueOf(sdk.Extended).IsNil() {
		tflog.Debug(ctx, "Packing typeless object for field Extended")
		var d diag.Diagnostics
		// Create an empty object with no attributes, which signifies its presence.
		model.Extended, d = basetypes.NewObjectValue(map[string]attr.Type{}, map[string]attr.Value{})
		diags.Append(d...)
	} else {
		// Since this field is part of a oneOf, being nil means it's not selected.
		// We make the object null with an empty attribute map.
		model.Extended = basetypes.NewObjectNull(map[string]attr.Type{})
	}
	// Handling Objects
	// This is a marker object (e.g. CHAP: {}). We just need to create an empty, non-null object.
	if sdk.Large != nil && !reflect.ValueOf(sdk.Large).IsNil() {
		tflog.Debug(ctx, "Packing typeless object for field Large")
		var d diag.Diagnostics
		// Create an empty object with no attributes, which signifies its presence.
		model.Large, d = basetypes.NewObjectValue(map[string]attr.Type{}, map[string]attr.Value{})
		diags.Append(d...)
	} else {
		// Since this field is part of a oneOf, being nil means it's not selected.
		// We make the object null with an empty attribute map.
		model.Large = basetypes.NewObjectNull(map[string]attr.Type{})
	}
	// Handling Objects
	// This is a marker object (e.g. CHAP: {}). We just need to create an empty, non-null object.
	if sdk.Standard != nil && !reflect.ValueOf(sdk.Standard).IsNil() {
		tflog.Debug(ctx, "Packing typeless object for field Standard")
		var d diag.Diagnostics
		// Create an empty object with no attributes, which signifies its presence.
		model.Standard, d = basetypes.NewObjectValue(map[string]attr.Type{}, map[string]attr.Value{})
		diags.Append(d...)
	} else {
		// Since this field is part of a oneOf, being nil means it's not selected.
		// We make the object null with an empty attribute map.
		model.Standard = basetypes.NewObjectNull(map[string]attr.Type{})
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BgpAddressFamilySendCommunity{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpAddressFamilySendCommunity", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpAddressFamilySendCommunity ---
func unpackBgpAddressFamilySendCommunityListToSdk(ctx context.Context, list types.List) ([]network_services.BgpAddressFamilySendCommunity, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpAddressFamilySendCommunity")
	diags := diag.Diagnostics{}
	var data []models.BgpAddressFamilySendCommunity
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpAddressFamilySendCommunity, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpAddressFamilySendCommunity{}.AttrTypes(), &item)
		unpacked, d := unpackBgpAddressFamilySendCommunityToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpAddressFamilySendCommunity", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpAddressFamilySendCommunity ---
func packBgpAddressFamilySendCommunityListFromSdk(ctx context.Context, sdks []network_services.BgpAddressFamilySendCommunity) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpAddressFamilySendCommunity")
	diags := diag.Diagnostics{}
	var data []models.BgpAddressFamilySendCommunity

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpAddressFamilySendCommunity
		obj, d := packBgpAddressFamilySendCommunityFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpAddressFamilySendCommunity{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpAddressFamilySendCommunity", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpAddressFamilySendCommunity{}.AttrType(), data)
}
