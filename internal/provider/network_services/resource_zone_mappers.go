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

// --- Unpacker for Zones ---
func unpackZonesToSdk(ctx context.Context, obj types.Object) (*network_services.Zones, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.Zones", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.Zones
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.Zones
	var d diag.Diagnostics

	// Handling Primitives
	if !model.Device.IsNull() && !model.Device.IsUnknown() {
		sdk.Device = model.Device.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Device", "value": *sdk.Device})
	}

	// Handling Objects
	if !model.DeviceAcl.IsNull() && !model.DeviceAcl.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field DeviceAcl")
		unpacked, d := unpackZonesDeviceAclToSdk(ctx, model.DeviceAcl)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "DeviceAcl"})
		}
		if unpacked != nil {
			sdk.DeviceAcl = unpacked
		}
	}

	// Handling Primitives
	if !model.DosLogSetting.IsNull() && !model.DosLogSetting.IsUnknown() {
		sdk.DosLogSetting = model.DosLogSetting.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DosLogSetting", "value": *sdk.DosLogSetting})
	}

	// Handling Primitives
	if !model.DosProfile.IsNull() && !model.DosProfile.IsUnknown() {
		sdk.DosProfile = model.DosProfile.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DosProfile", "value": *sdk.DosProfile})
	}

	// Handling Primitives
	if !model.EnableDeviceIdentification.IsNull() && !model.EnableDeviceIdentification.IsUnknown() {
		sdk.EnableDeviceIdentification = model.EnableDeviceIdentification.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "EnableDeviceIdentification", "value": *sdk.EnableDeviceIdentification})
	}

	// Handling Primitives
	if !model.EnableUserIdentification.IsNull() && !model.EnableUserIdentification.IsUnknown() {
		sdk.EnableUserIdentification = model.EnableUserIdentification.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "EnableUserIdentification", "value": *sdk.EnableUserIdentification})
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
	if !model.Network.IsNull() && !model.Network.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Network")
		unpacked, d := unpackZonesNetworkToSdk(ctx, model.Network)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Network"})
		}
		if unpacked != nil {
			sdk.Network = unpacked
		}
	}

	// Handling Primitives
	if !model.Snippet.IsNull() && !model.Snippet.IsUnknown() {
		sdk.Snippet = model.Snippet.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	}

	// Handling Objects
	if !model.UserAcl.IsNull() && !model.UserAcl.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field UserAcl")
		unpacked, d := unpackZonesDeviceAclToSdk(ctx, model.UserAcl)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "UserAcl"})
		}
		if unpacked != nil {
			sdk.UserAcl = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.Zones", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for Zones ---
func packZonesFromSdk(ctx context.Context, sdk network_services.Zones) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.Zones", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.Zones
	var d diag.Diagnostics
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
	if sdk.DeviceAcl != nil {
		tflog.Debug(ctx, "Packing nested object for field DeviceAcl")
		packed, d := packZonesDeviceAclFromSdk(ctx, *sdk.DeviceAcl)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "DeviceAcl"})
		}
		model.DeviceAcl = packed
	} else {
		model.DeviceAcl = basetypes.NewObjectNull(models.ZonesDeviceAcl{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.DosLogSetting != nil {
		model.DosLogSetting = basetypes.NewStringValue(*sdk.DosLogSetting)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DosLogSetting", "value": *sdk.DosLogSetting})
	} else {
		model.DosLogSetting = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.DosProfile != nil {
		model.DosProfile = basetypes.NewStringValue(*sdk.DosProfile)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DosProfile", "value": *sdk.DosProfile})
	} else {
		model.DosProfile = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.EnableDeviceIdentification != nil {
		model.EnableDeviceIdentification = basetypes.NewBoolValue(*sdk.EnableDeviceIdentification)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "EnableDeviceIdentification", "value": *sdk.EnableDeviceIdentification})
	} else {
		model.EnableDeviceIdentification = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.EnableUserIdentification != nil {
		model.EnableUserIdentification = basetypes.NewBoolValue(*sdk.EnableUserIdentification)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "EnableUserIdentification", "value": *sdk.EnableUserIdentification})
	} else {
		model.EnableUserIdentification = basetypes.NewBoolNull()
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
	if sdk.Network != nil {
		tflog.Debug(ctx, "Packing nested object for field Network")
		packed, d := packZonesNetworkFromSdk(ctx, *sdk.Network)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Network"})
		}
		model.Network = packed
	} else {
		model.Network = basetypes.NewObjectNull(models.ZonesNetwork{}.AttrTypes())
	}
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
	if sdk.UserAcl != nil {
		tflog.Debug(ctx, "Packing nested object for field UserAcl")
		packed, d := packZonesDeviceAclFromSdk(ctx, *sdk.UserAcl)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "UserAcl"})
		}
		model.UserAcl = packed
	} else {
		model.UserAcl = basetypes.NewObjectNull(models.ZonesDeviceAcl{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.Zones{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.Zones", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for Zones ---
func unpackZonesListToSdk(ctx context.Context, list types.List) ([]network_services.Zones, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.Zones")
	diags := diag.Diagnostics{}
	var data []models.Zones
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.Zones, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.Zones{}.AttrTypes(), &item)
		unpacked, d := unpackZonesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.Zones", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for Zones ---
func packZonesListFromSdk(ctx context.Context, sdks []network_services.Zones) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.Zones")
	diags := diag.Diagnostics{}
	var data []models.Zones

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.Zones
		obj, d := packZonesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.Zones{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.Zones", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.Zones{}.AttrType(), data)
}

// --- Unpacker for ZonesDeviceAcl ---
func unpackZonesDeviceAclToSdk(ctx context.Context, obj types.Object) (*network_services.ZonesDeviceAcl, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ZonesDeviceAcl", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ZonesDeviceAcl
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.ZonesDeviceAcl
	var d diag.Diagnostics
	// Handling Lists
	if !model.ExcludeList.IsNull() && !model.ExcludeList.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field ExcludeList")
		diags.Append(model.ExcludeList.ElementsAs(ctx, &sdk.ExcludeList, false)...)
	}

	// Handling Lists
	if !model.IncludeList.IsNull() && !model.IncludeList.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field IncludeList")
		diags.Append(model.IncludeList.ElementsAs(ctx, &sdk.IncludeList, false)...)
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ZonesDeviceAcl", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ZonesDeviceAcl ---
func packZonesDeviceAclFromSdk(ctx context.Context, sdk network_services.ZonesDeviceAcl) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ZonesDeviceAcl", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ZonesDeviceAcl
	var d diag.Diagnostics
	// Handling Lists
	if sdk.ExcludeList != nil {
		tflog.Debug(ctx, "Packing list of primitives for field ExcludeList")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.ExcludeList, d = basetypes.NewListValueFrom(ctx, elemType, sdk.ExcludeList)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.ExcludeList = basetypes.NewListNull(elemType)
	}
	// Handling Lists
	if sdk.IncludeList != nil {
		tflog.Debug(ctx, "Packing list of primitives for field IncludeList")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.IncludeList, d = basetypes.NewListValueFrom(ctx, elemType, sdk.IncludeList)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.IncludeList = basetypes.NewListNull(elemType)
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ZonesDeviceAcl{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ZonesDeviceAcl", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ZonesDeviceAcl ---
func unpackZonesDeviceAclListToSdk(ctx context.Context, list types.List) ([]network_services.ZonesDeviceAcl, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ZonesDeviceAcl")
	diags := diag.Diagnostics{}
	var data []models.ZonesDeviceAcl
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.ZonesDeviceAcl, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ZonesDeviceAcl{}.AttrTypes(), &item)
		unpacked, d := unpackZonesDeviceAclToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ZonesDeviceAcl", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ZonesDeviceAcl ---
func packZonesDeviceAclListFromSdk(ctx context.Context, sdks []network_services.ZonesDeviceAcl) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ZonesDeviceAcl")
	diags := diag.Diagnostics{}
	var data []models.ZonesDeviceAcl

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ZonesDeviceAcl
		obj, d := packZonesDeviceAclFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ZonesDeviceAcl{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ZonesDeviceAcl", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ZonesDeviceAcl{}.AttrType(), data)
}

// --- Unpacker for ZonesNetwork ---
func unpackZonesNetworkToSdk(ctx context.Context, obj types.Object) (*network_services.ZonesNetwork, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ZonesNetwork", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ZonesNetwork
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.ZonesNetwork
	var d diag.Diagnostics
	// Handling Primitives
	if !model.EnablePacketBufferProtection.IsNull() && !model.EnablePacketBufferProtection.IsUnknown() {
		sdk.EnablePacketBufferProtection = model.EnablePacketBufferProtection.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "EnablePacketBufferProtection", "value": *sdk.EnablePacketBufferProtection})
	}

	// Handling Primitives
	if !model.LogSetting.IsNull() && !model.LogSetting.IsUnknown() {
		sdk.LogSetting = model.LogSetting.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "LogSetting", "value": *sdk.LogSetting})
	}

	// Handling Primitives
	if !model.ZoneProtectionProfile.IsNull() && !model.ZoneProtectionProfile.IsUnknown() {
		sdk.ZoneProtectionProfile = model.ZoneProtectionProfile.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "ZoneProtectionProfile", "value": *sdk.ZoneProtectionProfile})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ZonesNetwork", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ZonesNetwork ---
func packZonesNetworkFromSdk(ctx context.Context, sdk network_services.ZonesNetwork) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ZonesNetwork", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ZonesNetwork
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.EnablePacketBufferProtection != nil {
		model.EnablePacketBufferProtection = basetypes.NewBoolValue(*sdk.EnablePacketBufferProtection)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "EnablePacketBufferProtection", "value": *sdk.EnablePacketBufferProtection})
	} else {
		model.EnablePacketBufferProtection = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.LogSetting != nil {
		model.LogSetting = basetypes.NewStringValue(*sdk.LogSetting)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "LogSetting", "value": *sdk.LogSetting})
	} else {
		model.LogSetting = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.ZoneProtectionProfile != nil {
		model.ZoneProtectionProfile = basetypes.NewStringValue(*sdk.ZoneProtectionProfile)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "ZoneProtectionProfile", "value": *sdk.ZoneProtectionProfile})
	} else {
		model.ZoneProtectionProfile = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ZonesNetwork{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ZonesNetwork", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ZonesNetwork ---
func unpackZonesNetworkListToSdk(ctx context.Context, list types.List) ([]network_services.ZonesNetwork, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ZonesNetwork")
	diags := diag.Diagnostics{}
	var data []models.ZonesNetwork
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.ZonesNetwork, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ZonesNetwork{}.AttrTypes(), &item)
		unpacked, d := unpackZonesNetworkToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ZonesNetwork", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ZonesNetwork ---
func packZonesNetworkListFromSdk(ctx context.Context, sdks []network_services.ZonesNetwork) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ZonesNetwork")
	diags := diag.Diagnostics{}
	var data []models.ZonesNetwork

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ZonesNetwork
		obj, d := packZonesNetworkFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ZonesNetwork{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ZonesNetwork", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ZonesNetwork{}.AttrType(), data)
}
