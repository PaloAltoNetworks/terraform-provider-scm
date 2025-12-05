package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/device_settings"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/device_settings"
)

// --- Unpacker for VpnSettings ---
func unpackVpnSettingsToSdk(ctx context.Context, obj types.Object) (*device_settings.VpnSettings, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.VpnSettings", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.VpnSettings
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk device_settings.VpnSettings
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

	// Handling Primitives
	if !model.Snippet.IsNull() && !model.Snippet.IsUnknown() {
		sdk.Snippet = model.Snippet.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	}

	// Handling Objects
	if !model.Vpn.IsNull() && !model.Vpn.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Vpn")
		unpacked, d := unpackVpnSettingsVpnToSdk(ctx, model.Vpn)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Vpn"})
		}
		if unpacked != nil {
			sdk.Vpn = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.VpnSettings", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for VpnSettings ---
func packVpnSettingsFromSdk(ctx context.Context, sdk device_settings.VpnSettings) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.VpnSettings", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.VpnSettings
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
	if sdk.Vpn != nil {
		tflog.Debug(ctx, "Packing nested object for field Vpn")
		packed, d := packVpnSettingsVpnFromSdk(ctx, *sdk.Vpn)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Vpn"})
		}
		model.Vpn = packed
	} else {
		model.Vpn = basetypes.NewObjectNull(models.VpnSettingsVpn{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.VpnSettings{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.VpnSettings", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for VpnSettings ---
func unpackVpnSettingsListToSdk(ctx context.Context, list types.List) ([]device_settings.VpnSettings, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.VpnSettings")
	diags := diag.Diagnostics{}
	var data []models.VpnSettings
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]device_settings.VpnSettings, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.VpnSettings{}.AttrTypes(), &item)
		unpacked, d := unpackVpnSettingsToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.VpnSettings", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for VpnSettings ---
func packVpnSettingsListFromSdk(ctx context.Context, sdks []device_settings.VpnSettings) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.VpnSettings")
	diags := diag.Diagnostics{}
	var data []models.VpnSettings

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.VpnSettings
		obj, d := packVpnSettingsFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.VpnSettings{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.VpnSettings", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.VpnSettings{}.AttrType(), data)
}

// --- Unpacker for VpnSettingsVpn ---
func unpackVpnSettingsVpnToSdk(ctx context.Context, obj types.Object) (*device_settings.VpnSettingsVpn, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.VpnSettingsVpn", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.VpnSettingsVpn
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk device_settings.VpnSettingsVpn
	var d diag.Diagnostics
	// Handling Objects
	if !model.Ikev2.IsNull() && !model.Ikev2.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Ikev2")
		unpacked, d := unpackVpnSettingsVpnIkev2ToSdk(ctx, model.Ikev2)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Ikev2"})
		}
		if unpacked != nil {
			sdk.Ikev2 = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.VpnSettingsVpn", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for VpnSettingsVpn ---
func packVpnSettingsVpnFromSdk(ctx context.Context, sdk device_settings.VpnSettingsVpn) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.VpnSettingsVpn", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.VpnSettingsVpn
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Ikev2 != nil {
		tflog.Debug(ctx, "Packing nested object for field Ikev2")
		packed, d := packVpnSettingsVpnIkev2FromSdk(ctx, *sdk.Ikev2)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Ikev2"})
		}
		model.Ikev2 = packed
	} else {
		model.Ikev2 = basetypes.NewObjectNull(models.VpnSettingsVpnIkev2{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.VpnSettingsVpn{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.VpnSettingsVpn", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for VpnSettingsVpn ---
func unpackVpnSettingsVpnListToSdk(ctx context.Context, list types.List) ([]device_settings.VpnSettingsVpn, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.VpnSettingsVpn")
	diags := diag.Diagnostics{}
	var data []models.VpnSettingsVpn
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]device_settings.VpnSettingsVpn, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.VpnSettingsVpn{}.AttrTypes(), &item)
		unpacked, d := unpackVpnSettingsVpnToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.VpnSettingsVpn", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for VpnSettingsVpn ---
func packVpnSettingsVpnListFromSdk(ctx context.Context, sdks []device_settings.VpnSettingsVpn) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.VpnSettingsVpn")
	diags := diag.Diagnostics{}
	var data []models.VpnSettingsVpn

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.VpnSettingsVpn
		obj, d := packVpnSettingsVpnFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.VpnSettingsVpn{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.VpnSettingsVpn", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.VpnSettingsVpn{}.AttrType(), data)
}

// --- Unpacker for VpnSettingsVpnIkev2 ---
func unpackVpnSettingsVpnIkev2ToSdk(ctx context.Context, obj types.Object) (*device_settings.VpnSettingsVpnIkev2, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.VpnSettingsVpnIkev2", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.VpnSettingsVpnIkev2
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk device_settings.VpnSettingsVpnIkev2
	var d diag.Diagnostics
	// Handling Primitives
	if !model.CertificateCacheSize.IsNull() && !model.CertificateCacheSize.IsUnknown() {
		val := int32(model.CertificateCacheSize.ValueInt64())
		sdk.CertificateCacheSize = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "CertificateCacheSize", "value": *sdk.CertificateCacheSize})
	}

	// Handling Primitives
	if !model.CookieThreshold.IsNull() && !model.CookieThreshold.IsUnknown() {
		val := int32(model.CookieThreshold.ValueInt64())
		sdk.CookieThreshold = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "CookieThreshold", "value": *sdk.CookieThreshold})
	}

	// Handling Primitives
	if !model.MaxHalfOpenedSa.IsNull() && !model.MaxHalfOpenedSa.IsUnknown() {
		val := int32(model.MaxHalfOpenedSa.ValueInt64())
		sdk.MaxHalfOpenedSa = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "MaxHalfOpenedSa", "value": *sdk.MaxHalfOpenedSa})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.VpnSettingsVpnIkev2", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for VpnSettingsVpnIkev2 ---
func packVpnSettingsVpnIkev2FromSdk(ctx context.Context, sdk device_settings.VpnSettingsVpnIkev2) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.VpnSettingsVpnIkev2", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.VpnSettingsVpnIkev2
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.CertificateCacheSize != nil {
		model.CertificateCacheSize = basetypes.NewInt64Value(int64(*sdk.CertificateCacheSize))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "CertificateCacheSize", "value": *sdk.CertificateCacheSize})
	} else {
		model.CertificateCacheSize = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.CookieThreshold != nil {
		model.CookieThreshold = basetypes.NewInt64Value(int64(*sdk.CookieThreshold))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "CookieThreshold", "value": *sdk.CookieThreshold})
	} else {
		model.CookieThreshold = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.MaxHalfOpenedSa != nil {
		model.MaxHalfOpenedSa = basetypes.NewInt64Value(int64(*sdk.MaxHalfOpenedSa))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "MaxHalfOpenedSa", "value": *sdk.MaxHalfOpenedSa})
	} else {
		model.MaxHalfOpenedSa = basetypes.NewInt64Null()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.VpnSettingsVpnIkev2{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.VpnSettingsVpnIkev2", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for VpnSettingsVpnIkev2 ---
func unpackVpnSettingsVpnIkev2ListToSdk(ctx context.Context, list types.List) ([]device_settings.VpnSettingsVpnIkev2, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.VpnSettingsVpnIkev2")
	diags := diag.Diagnostics{}
	var data []models.VpnSettingsVpnIkev2
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]device_settings.VpnSettingsVpnIkev2, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.VpnSettingsVpnIkev2{}.AttrTypes(), &item)
		unpacked, d := unpackVpnSettingsVpnIkev2ToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.VpnSettingsVpnIkev2", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for VpnSettingsVpnIkev2 ---
func packVpnSettingsVpnIkev2ListFromSdk(ctx context.Context, sdks []device_settings.VpnSettingsVpnIkev2) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.VpnSettingsVpnIkev2")
	diags := diag.Diagnostics{}
	var data []models.VpnSettingsVpnIkev2

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.VpnSettingsVpnIkev2
		obj, d := packVpnSettingsVpnIkev2FromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.VpnSettingsVpnIkev2{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.VpnSettingsVpnIkev2", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.VpnSettingsVpnIkev2{}.AttrType(), data)
}
