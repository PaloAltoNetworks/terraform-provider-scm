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

// --- Unpacker for GeneralSettings ---
func unpackGeneralSettingsToSdk(ctx context.Context, obj types.Object) (*device_settings.GeneralSettings, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.GeneralSettings", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.GeneralSettings
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk device_settings.GeneralSettings
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

	// Handling Objects
	if !model.General.IsNull() && !model.General.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field General")
		unpacked, d := unpackGeneralSettingsGeneralToSdk(ctx, model.General)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "General"})
		}
		if unpacked != nil {
			sdk.General = unpacked
		}
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

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.GeneralSettings", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for GeneralSettings ---
func packGeneralSettingsFromSdk(ctx context.Context, sdk device_settings.GeneralSettings) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.GeneralSettings", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.GeneralSettings
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
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.General != nil {
		tflog.Debug(ctx, "Packing nested object for field General")
		packed, d := packGeneralSettingsGeneralFromSdk(ctx, *sdk.General)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "General"})
		}
		model.General = packed
	} else {
		model.General = basetypes.NewObjectNull(models.GeneralSettingsGeneral{}.AttrTypes())
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
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.GeneralSettings{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.GeneralSettings", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for GeneralSettings ---
func unpackGeneralSettingsListToSdk(ctx context.Context, list types.List) ([]device_settings.GeneralSettings, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.GeneralSettings")
	diags := diag.Diagnostics{}
	var data []models.GeneralSettings
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]device_settings.GeneralSettings, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.GeneralSettings{}.AttrTypes(), &item)
		unpacked, d := unpackGeneralSettingsToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.GeneralSettings", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for GeneralSettings ---
func packGeneralSettingsListFromSdk(ctx context.Context, sdks []device_settings.GeneralSettings) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.GeneralSettings")
	diags := diag.Diagnostics{}
	var data []models.GeneralSettings

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.GeneralSettings
		obj, d := packGeneralSettingsFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.GeneralSettings{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.GeneralSettings", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.GeneralSettings{}.AttrType(), data)
}

// --- Unpacker for GeneralSettingsGeneral ---
func unpackGeneralSettingsGeneralToSdk(ctx context.Context, obj types.Object) (*device_settings.GeneralSettingsGeneral, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.GeneralSettingsGeneral", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.GeneralSettingsGeneral
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk device_settings.GeneralSettingsGeneral
	var d diag.Diagnostics
	// Handling Primitives
	if !model.AckLoginBanner.IsNull() && !model.AckLoginBanner.IsUnknown() {
		sdk.AckLoginBanner = model.AckLoginBanner.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "AckLoginBanner", "value": *sdk.AckLoginBanner})
	}

	// Handling Primitives
	if !model.Domain.IsNull() && !model.Domain.IsUnknown() {
		sdk.Domain = model.Domain.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Domain", "value": *sdk.Domain})
	}

	// Handling Objects
	if !model.GeoLocation.IsNull() && !model.GeoLocation.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field GeoLocation")
		unpacked, d := unpackGeneralSettingsGeneralGeoLocationToSdk(ctx, model.GeoLocation)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "GeoLocation"})
		}
		if unpacked != nil {
			sdk.GeoLocation = unpacked
		}
	}

	// Handling Primitives
	if !model.Locale.IsNull() && !model.Locale.IsUnknown() {
		sdk.Locale = model.Locale.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Locale", "value": *sdk.Locale})
	}

	// Handling Primitives
	if !model.LoginBanner.IsNull() && !model.LoginBanner.IsUnknown() {
		sdk.LoginBanner = model.LoginBanner.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "LoginBanner", "value": *sdk.LoginBanner})
	}

	// Handling Objects
	if !model.Setting.IsNull() && !model.Setting.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Setting")
		unpacked, d := unpackGeneralSettingsGeneralSettingToSdk(ctx, model.Setting)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Setting"})
		}
		if unpacked != nil {
			sdk.Setting = unpacked
		}
	}

	// Handling Primitives
	if !model.SslTlsServiceProfile.IsNull() && !model.SslTlsServiceProfile.IsUnknown() {
		sdk.SslTlsServiceProfile = model.SslTlsServiceProfile.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "SslTlsServiceProfile", "value": *sdk.SslTlsServiceProfile})
	}

	// Handling Primitives
	if !model.Timezone.IsNull() && !model.Timezone.IsUnknown() {
		sdk.Timezone = model.Timezone.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Timezone", "value": *sdk.Timezone})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.GeneralSettingsGeneral", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for GeneralSettingsGeneral ---
func packGeneralSettingsGeneralFromSdk(ctx context.Context, sdk device_settings.GeneralSettingsGeneral) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.GeneralSettingsGeneral", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.GeneralSettingsGeneral
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.AckLoginBanner != nil {
		model.AckLoginBanner = basetypes.NewBoolValue(*sdk.AckLoginBanner)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "AckLoginBanner", "value": *sdk.AckLoginBanner})
	} else {
		model.AckLoginBanner = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Domain != nil {
		model.Domain = basetypes.NewStringValue(*sdk.Domain)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Domain", "value": *sdk.Domain})
	} else {
		model.Domain = basetypes.NewStringNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.GeoLocation != nil {
		tflog.Debug(ctx, "Packing nested object for field GeoLocation")
		packed, d := packGeneralSettingsGeneralGeoLocationFromSdk(ctx, *sdk.GeoLocation)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "GeoLocation"})
		}
		model.GeoLocation = packed
	} else {
		model.GeoLocation = basetypes.NewObjectNull(models.GeneralSettingsGeneralGeoLocation{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Locale != nil {
		model.Locale = basetypes.NewStringValue(*sdk.Locale)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Locale", "value": *sdk.Locale})
	} else {
		model.Locale = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.LoginBanner != nil {
		model.LoginBanner = basetypes.NewStringValue(*sdk.LoginBanner)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "LoginBanner", "value": *sdk.LoginBanner})
	} else {
		model.LoginBanner = basetypes.NewStringNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Setting != nil {
		tflog.Debug(ctx, "Packing nested object for field Setting")
		packed, d := packGeneralSettingsGeneralSettingFromSdk(ctx, *sdk.Setting)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Setting"})
		}
		model.Setting = packed
	} else {
		model.Setting = basetypes.NewObjectNull(models.GeneralSettingsGeneralSetting{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.SslTlsServiceProfile != nil {
		model.SslTlsServiceProfile = basetypes.NewStringValue(*sdk.SslTlsServiceProfile)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "SslTlsServiceProfile", "value": *sdk.SslTlsServiceProfile})
	} else {
		model.SslTlsServiceProfile = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Timezone != nil {
		model.Timezone = basetypes.NewStringValue(*sdk.Timezone)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Timezone", "value": *sdk.Timezone})
	} else {
		model.Timezone = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.GeneralSettingsGeneral{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.GeneralSettingsGeneral", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for GeneralSettingsGeneral ---
func unpackGeneralSettingsGeneralListToSdk(ctx context.Context, list types.List) ([]device_settings.GeneralSettingsGeneral, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.GeneralSettingsGeneral")
	diags := diag.Diagnostics{}
	var data []models.GeneralSettingsGeneral
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]device_settings.GeneralSettingsGeneral, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.GeneralSettingsGeneral{}.AttrTypes(), &item)
		unpacked, d := unpackGeneralSettingsGeneralToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.GeneralSettingsGeneral", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for GeneralSettingsGeneral ---
func packGeneralSettingsGeneralListFromSdk(ctx context.Context, sdks []device_settings.GeneralSettingsGeneral) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.GeneralSettingsGeneral")
	diags := diag.Diagnostics{}
	var data []models.GeneralSettingsGeneral

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.GeneralSettingsGeneral
		obj, d := packGeneralSettingsGeneralFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.GeneralSettingsGeneral{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.GeneralSettingsGeneral", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.GeneralSettingsGeneral{}.AttrType(), data)
}

// --- Unpacker for GeneralSettingsGeneralGeoLocation ---
func unpackGeneralSettingsGeneralGeoLocationToSdk(ctx context.Context, obj types.Object) (*device_settings.GeneralSettingsGeneralGeoLocation, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.GeneralSettingsGeneralGeoLocation", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.GeneralSettingsGeneralGeoLocation
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk device_settings.GeneralSettingsGeneralGeoLocation
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Latitude.IsNull() && !model.Latitude.IsUnknown() {
		sdk.Latitude = model.Latitude.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Latitude", "value": *sdk.Latitude})
	}

	// Handling Primitives
	if !model.Longitude.IsNull() && !model.Longitude.IsUnknown() {
		sdk.Longitude = model.Longitude.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Longitude", "value": *sdk.Longitude})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.GeneralSettingsGeneralGeoLocation", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for GeneralSettingsGeneralGeoLocation ---
func packGeneralSettingsGeneralGeoLocationFromSdk(ctx context.Context, sdk device_settings.GeneralSettingsGeneralGeoLocation) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.GeneralSettingsGeneralGeoLocation", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.GeneralSettingsGeneralGeoLocation
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Latitude != nil {
		model.Latitude = basetypes.NewStringValue(*sdk.Latitude)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Latitude", "value": *sdk.Latitude})
	} else {
		model.Latitude = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Longitude != nil {
		model.Longitude = basetypes.NewStringValue(*sdk.Longitude)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Longitude", "value": *sdk.Longitude})
	} else {
		model.Longitude = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.GeneralSettingsGeneralGeoLocation{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.GeneralSettingsGeneralGeoLocation", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for GeneralSettingsGeneralGeoLocation ---
func unpackGeneralSettingsGeneralGeoLocationListToSdk(ctx context.Context, list types.List) ([]device_settings.GeneralSettingsGeneralGeoLocation, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.GeneralSettingsGeneralGeoLocation")
	diags := diag.Diagnostics{}
	var data []models.GeneralSettingsGeneralGeoLocation
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]device_settings.GeneralSettingsGeneralGeoLocation, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.GeneralSettingsGeneralGeoLocation{}.AttrTypes(), &item)
		unpacked, d := unpackGeneralSettingsGeneralGeoLocationToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.GeneralSettingsGeneralGeoLocation", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for GeneralSettingsGeneralGeoLocation ---
func packGeneralSettingsGeneralGeoLocationListFromSdk(ctx context.Context, sdks []device_settings.GeneralSettingsGeneralGeoLocation) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.GeneralSettingsGeneralGeoLocation")
	diags := diag.Diagnostics{}
	var data []models.GeneralSettingsGeneralGeoLocation

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.GeneralSettingsGeneralGeoLocation
		obj, d := packGeneralSettingsGeneralGeoLocationFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.GeneralSettingsGeneralGeoLocation{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.GeneralSettingsGeneralGeoLocation", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.GeneralSettingsGeneralGeoLocation{}.AttrType(), data)
}

// --- Unpacker for GeneralSettingsGeneralSetting ---
func unpackGeneralSettingsGeneralSettingToSdk(ctx context.Context, obj types.Object) (*device_settings.GeneralSettingsGeneralSetting, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.GeneralSettingsGeneralSetting", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.GeneralSettingsGeneralSetting
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk device_settings.GeneralSettingsGeneralSetting
	var d diag.Diagnostics
	// Handling Primitives
	if !model.AutoMacDetect.IsNull() && !model.AutoMacDetect.IsUnknown() {
		sdk.AutoMacDetect = model.AutoMacDetect.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "AutoMacDetect", "value": *sdk.AutoMacDetect})
	}

	// Handling Primitives
	if !model.FailOpen.IsNull() && !model.FailOpen.IsUnknown() {
		sdk.FailOpen = model.FailOpen.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "FailOpen", "value": *sdk.FailOpen})
	}

	// Handling Objects
	if !model.Management.IsNull() && !model.Management.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Management")
		unpacked, d := unpackGeneralSettingsGeneralSettingManagementToSdk(ctx, model.Management)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Management"})
		}
		if unpacked != nil {
			sdk.Management = unpacked
		}
	}

	// Handling Primitives
	if !model.TunnelAcceleration.IsNull() && !model.TunnelAcceleration.IsUnknown() {
		sdk.TunnelAcceleration = model.TunnelAcceleration.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "TunnelAcceleration", "value": *sdk.TunnelAcceleration})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.GeneralSettingsGeneralSetting", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for GeneralSettingsGeneralSetting ---
func packGeneralSettingsGeneralSettingFromSdk(ctx context.Context, sdk device_settings.GeneralSettingsGeneralSetting) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.GeneralSettingsGeneralSetting", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.GeneralSettingsGeneralSetting
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.AutoMacDetect != nil {
		model.AutoMacDetect = basetypes.NewBoolValue(*sdk.AutoMacDetect)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "AutoMacDetect", "value": *sdk.AutoMacDetect})
	} else {
		model.AutoMacDetect = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.FailOpen != nil {
		model.FailOpen = basetypes.NewBoolValue(*sdk.FailOpen)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "FailOpen", "value": *sdk.FailOpen})
	} else {
		model.FailOpen = basetypes.NewBoolNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Management != nil {
		tflog.Debug(ctx, "Packing nested object for field Management")
		packed, d := packGeneralSettingsGeneralSettingManagementFromSdk(ctx, *sdk.Management)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Management"})
		}
		model.Management = packed
	} else {
		model.Management = basetypes.NewObjectNull(models.GeneralSettingsGeneralSettingManagement{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.TunnelAcceleration != nil {
		model.TunnelAcceleration = basetypes.NewBoolValue(*sdk.TunnelAcceleration)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "TunnelAcceleration", "value": *sdk.TunnelAcceleration})
	} else {
		model.TunnelAcceleration = basetypes.NewBoolNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.GeneralSettingsGeneralSetting{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.GeneralSettingsGeneralSetting", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for GeneralSettingsGeneralSetting ---
func unpackGeneralSettingsGeneralSettingListToSdk(ctx context.Context, list types.List) ([]device_settings.GeneralSettingsGeneralSetting, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.GeneralSettingsGeneralSetting")
	diags := diag.Diagnostics{}
	var data []models.GeneralSettingsGeneralSetting
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]device_settings.GeneralSettingsGeneralSetting, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.GeneralSettingsGeneralSetting{}.AttrTypes(), &item)
		unpacked, d := unpackGeneralSettingsGeneralSettingToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.GeneralSettingsGeneralSetting", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for GeneralSettingsGeneralSetting ---
func packGeneralSettingsGeneralSettingListFromSdk(ctx context.Context, sdks []device_settings.GeneralSettingsGeneralSetting) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.GeneralSettingsGeneralSetting")
	diags := diag.Diagnostics{}
	var data []models.GeneralSettingsGeneralSetting

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.GeneralSettingsGeneralSetting
		obj, d := packGeneralSettingsGeneralSettingFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.GeneralSettingsGeneralSetting{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.GeneralSettingsGeneralSetting", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.GeneralSettingsGeneralSetting{}.AttrType(), data)
}

// --- Unpacker for GeneralSettingsGeneralSettingManagement ---
func unpackGeneralSettingsGeneralSettingManagementToSdk(ctx context.Context, obj types.Object) (*device_settings.GeneralSettingsGeneralSettingManagement, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.GeneralSettingsGeneralSettingManagement", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.GeneralSettingsGeneralSettingManagement
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk device_settings.GeneralSettingsGeneralSettingManagement
	var d diag.Diagnostics
	// Handling Primitives
	if !model.AutoAcquireCommitLock.IsNull() && !model.AutoAcquireCommitLock.IsUnknown() {
		sdk.AutoAcquireCommitLock = model.AutoAcquireCommitLock.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "AutoAcquireCommitLock", "value": *sdk.AutoAcquireCommitLock})
	}

	// Handling Primitives
	if !model.EnableCertificateExpirationCheck.IsNull() && !model.EnableCertificateExpirationCheck.IsUnknown() {
		sdk.EnableCertificateExpirationCheck = model.EnableCertificateExpirationCheck.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "EnableCertificateExpirationCheck", "value": *sdk.EnableCertificateExpirationCheck})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.GeneralSettingsGeneralSettingManagement", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for GeneralSettingsGeneralSettingManagement ---
func packGeneralSettingsGeneralSettingManagementFromSdk(ctx context.Context, sdk device_settings.GeneralSettingsGeneralSettingManagement) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.GeneralSettingsGeneralSettingManagement", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.GeneralSettingsGeneralSettingManagement
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.AutoAcquireCommitLock != nil {
		model.AutoAcquireCommitLock = basetypes.NewBoolValue(*sdk.AutoAcquireCommitLock)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "AutoAcquireCommitLock", "value": *sdk.AutoAcquireCommitLock})
	} else {
		model.AutoAcquireCommitLock = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.EnableCertificateExpirationCheck != nil {
		model.EnableCertificateExpirationCheck = basetypes.NewBoolValue(*sdk.EnableCertificateExpirationCheck)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "EnableCertificateExpirationCheck", "value": *sdk.EnableCertificateExpirationCheck})
	} else {
		model.EnableCertificateExpirationCheck = basetypes.NewBoolNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.GeneralSettingsGeneralSettingManagement{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.GeneralSettingsGeneralSettingManagement", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for GeneralSettingsGeneralSettingManagement ---
func unpackGeneralSettingsGeneralSettingManagementListToSdk(ctx context.Context, list types.List) ([]device_settings.GeneralSettingsGeneralSettingManagement, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.GeneralSettingsGeneralSettingManagement")
	diags := diag.Diagnostics{}
	var data []models.GeneralSettingsGeneralSettingManagement
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]device_settings.GeneralSettingsGeneralSettingManagement, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.GeneralSettingsGeneralSettingManagement{}.AttrTypes(), &item)
		unpacked, d := unpackGeneralSettingsGeneralSettingManagementToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.GeneralSettingsGeneralSettingManagement", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for GeneralSettingsGeneralSettingManagement ---
func packGeneralSettingsGeneralSettingManagementListFromSdk(ctx context.Context, sdks []device_settings.GeneralSettingsGeneralSettingManagement) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.GeneralSettingsGeneralSettingManagement")
	diags := diag.Diagnostics{}
	var data []models.GeneralSettingsGeneralSettingManagement

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.GeneralSettingsGeneralSettingManagement
		obj, d := packGeneralSettingsGeneralSettingManagementFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.GeneralSettingsGeneralSettingManagement{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.GeneralSettingsGeneralSettingManagement", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.GeneralSettingsGeneralSettingManagement{}.AttrType(), data)
}
