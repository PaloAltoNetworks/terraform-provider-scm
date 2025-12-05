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

// --- Unpacker for MotdBannerSettings ---
func unpackMotdBannerSettingsToSdk(ctx context.Context, obj types.Object) (*device_settings.MotdBannerSettings, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.MotdBannerSettings", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.MotdBannerSettings
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk device_settings.MotdBannerSettings
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
	if !model.MotdAndBanner.IsNull() && !model.MotdAndBanner.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field MotdAndBanner")
		unpacked, d := unpackMotdBannerSettingsMotdAndBannerToSdk(ctx, model.MotdAndBanner)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "MotdAndBanner"})
		}
		if unpacked != nil {
			sdk.MotdAndBanner = unpacked
		}
	}

	// Handling Primitives
	if !model.Snippet.IsNull() && !model.Snippet.IsUnknown() {
		sdk.Snippet = model.Snippet.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.MotdBannerSettings", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for MotdBannerSettings ---
func packMotdBannerSettingsFromSdk(ctx context.Context, sdk device_settings.MotdBannerSettings) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.MotdBannerSettings", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.MotdBannerSettings
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
	if sdk.MotdAndBanner != nil {
		tflog.Debug(ctx, "Packing nested object for field MotdAndBanner")
		packed, d := packMotdBannerSettingsMotdAndBannerFromSdk(ctx, *sdk.MotdAndBanner)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "MotdAndBanner"})
		}
		model.MotdAndBanner = packed
	} else {
		model.MotdAndBanner = basetypes.NewObjectNull(models.MotdBannerSettingsMotdAndBanner{}.AttrTypes())
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

	obj, d := types.ObjectValueFrom(ctx, models.MotdBannerSettings{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.MotdBannerSettings", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for MotdBannerSettings ---
func unpackMotdBannerSettingsListToSdk(ctx context.Context, list types.List) ([]device_settings.MotdBannerSettings, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.MotdBannerSettings")
	diags := diag.Diagnostics{}
	var data []models.MotdBannerSettings
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]device_settings.MotdBannerSettings, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.MotdBannerSettings{}.AttrTypes(), &item)
		unpacked, d := unpackMotdBannerSettingsToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.MotdBannerSettings", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for MotdBannerSettings ---
func packMotdBannerSettingsListFromSdk(ctx context.Context, sdks []device_settings.MotdBannerSettings) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.MotdBannerSettings")
	diags := diag.Diagnostics{}
	var data []models.MotdBannerSettings

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.MotdBannerSettings
		obj, d := packMotdBannerSettingsFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.MotdBannerSettings{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.MotdBannerSettings", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.MotdBannerSettings{}.AttrType(), data)
}

// --- Unpacker for MotdBannerSettingsMotdAndBanner ---
func unpackMotdBannerSettingsMotdAndBannerToSdk(ctx context.Context, obj types.Object) (*device_settings.MotdBannerSettingsMotdAndBanner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.MotdBannerSettingsMotdAndBanner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.MotdBannerSettingsMotdAndBanner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk device_settings.MotdBannerSettingsMotdAndBanner
	var d diag.Diagnostics
	// Handling Primitives
	if !model.BannerFooter.IsNull() && !model.BannerFooter.IsUnknown() {
		sdk.BannerFooter = model.BannerFooter.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "BannerFooter", "value": *sdk.BannerFooter})
	}

	// Handling Primitives
	if !model.BannerFooterColor.IsNull() && !model.BannerFooterColor.IsUnknown() {
		// Enum Handling
		val := device_settings.MotdColor(model.BannerFooterColor.ValueString())
		sdk.BannerFooterColor = &val
	}

	// Handling Primitives
	if !model.BannerFooterTextColor.IsNull() && !model.BannerFooterTextColor.IsUnknown() {
		// Enum Handling
		val := device_settings.MotdColor(model.BannerFooterTextColor.ValueString())
		sdk.BannerFooterTextColor = &val
	}

	// Handling Primitives
	if !model.BannerHeader.IsNull() && !model.BannerHeader.IsUnknown() {
		sdk.BannerHeader = model.BannerHeader.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "BannerHeader", "value": *sdk.BannerHeader})
	}

	// Handling Primitives
	if !model.BannerHeaderColor.IsNull() && !model.BannerHeaderColor.IsUnknown() {
		// Enum Handling
		val := device_settings.MotdColor(model.BannerHeaderColor.ValueString())
		sdk.BannerHeaderColor = &val
	}

	// Handling Primitives
	if !model.BannerHeaderFooterMatch.IsNull() && !model.BannerHeaderFooterMatch.IsUnknown() {
		sdk.BannerHeaderFooterMatch = model.BannerHeaderFooterMatch.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "BannerHeaderFooterMatch", "value": *sdk.BannerHeaderFooterMatch})
	}

	// Handling Primitives
	if !model.BannerHeaderTextColor.IsNull() && !model.BannerHeaderTextColor.IsUnknown() {
		// Enum Handling
		val := device_settings.MotdColor(model.BannerHeaderTextColor.ValueString())
		sdk.BannerHeaderTextColor = &val
	}

	// Handling Primitives
	if !model.Message.IsNull() && !model.Message.IsUnknown() {
		sdk.Message = model.Message.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Message", "value": *sdk.Message})
	}

	// Handling Primitives
	if !model.MotdColor.IsNull() && !model.MotdColor.IsUnknown() {
		// Enum Handling
		val := device_settings.MotdColor(model.MotdColor.ValueString())
		sdk.MotdColor = &val
	}

	// Handling Primitives
	if !model.MotdDoNotDisplayAgain.IsNull() && !model.MotdDoNotDisplayAgain.IsUnknown() {
		sdk.MotdDoNotDisplayAgain = model.MotdDoNotDisplayAgain.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "MotdDoNotDisplayAgain", "value": *sdk.MotdDoNotDisplayAgain})
	}

	// Handling Primitives
	if !model.MotdEnable.IsNull() && !model.MotdEnable.IsUnknown() {
		sdk.MotdEnable = model.MotdEnable.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "MotdEnable", "value": *sdk.MotdEnable})
	}

	// Handling Primitives
	if !model.MotdTitle.IsNull() && !model.MotdTitle.IsUnknown() {
		sdk.MotdTitle = model.MotdTitle.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "MotdTitle", "value": *sdk.MotdTitle})
	}

	// Handling Primitives
	if !model.Severity.IsNull() && !model.Severity.IsUnknown() {
		sdk.Severity = model.Severity.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Severity", "value": *sdk.Severity})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.MotdBannerSettingsMotdAndBanner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for MotdBannerSettingsMotdAndBanner ---
func packMotdBannerSettingsMotdAndBannerFromSdk(ctx context.Context, sdk device_settings.MotdBannerSettingsMotdAndBanner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.MotdBannerSettingsMotdAndBanner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.MotdBannerSettingsMotdAndBanner
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.BannerFooter != nil {
		model.BannerFooter = basetypes.NewStringValue(*sdk.BannerFooter)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "BannerFooter", "value": *sdk.BannerFooter})
	} else {
		model.BannerFooter = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.BannerFooterColor != nil {
		model.BannerFooterColor = basetypes.NewStringValue(string(*sdk.BannerFooterColor))
	} else {
		model.BannerFooterColor = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.BannerFooterTextColor != nil {
		model.BannerFooterTextColor = basetypes.NewStringValue(string(*sdk.BannerFooterTextColor))
	} else {
		model.BannerFooterTextColor = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.BannerHeader != nil {
		model.BannerHeader = basetypes.NewStringValue(*sdk.BannerHeader)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "BannerHeader", "value": *sdk.BannerHeader})
	} else {
		model.BannerHeader = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.BannerHeaderColor != nil {
		model.BannerHeaderColor = basetypes.NewStringValue(string(*sdk.BannerHeaderColor))
	} else {
		model.BannerHeaderColor = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.BannerHeaderFooterMatch != nil {
		model.BannerHeaderFooterMatch = basetypes.NewBoolValue(*sdk.BannerHeaderFooterMatch)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "BannerHeaderFooterMatch", "value": *sdk.BannerHeaderFooterMatch})
	} else {
		model.BannerHeaderFooterMatch = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.BannerHeaderTextColor != nil {
		model.BannerHeaderTextColor = basetypes.NewStringValue(string(*sdk.BannerHeaderTextColor))
	} else {
		model.BannerHeaderTextColor = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Message != nil {
		model.Message = basetypes.NewStringValue(*sdk.Message)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Message", "value": *sdk.Message})
	} else {
		model.Message = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.MotdColor != nil {
		model.MotdColor = basetypes.NewStringValue(string(*sdk.MotdColor))
	} else {
		model.MotdColor = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.MotdDoNotDisplayAgain != nil {
		model.MotdDoNotDisplayAgain = basetypes.NewBoolValue(*sdk.MotdDoNotDisplayAgain)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "MotdDoNotDisplayAgain", "value": *sdk.MotdDoNotDisplayAgain})
	} else {
		model.MotdDoNotDisplayAgain = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.MotdEnable != nil {
		model.MotdEnable = basetypes.NewBoolValue(*sdk.MotdEnable)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "MotdEnable", "value": *sdk.MotdEnable})
	} else {
		model.MotdEnable = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.MotdTitle != nil {
		model.MotdTitle = basetypes.NewStringValue(*sdk.MotdTitle)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "MotdTitle", "value": *sdk.MotdTitle})
	} else {
		model.MotdTitle = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Severity != nil {
		model.Severity = basetypes.NewStringValue(*sdk.Severity)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Severity", "value": *sdk.Severity})
	} else {
		model.Severity = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.MotdBannerSettingsMotdAndBanner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.MotdBannerSettingsMotdAndBanner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for MotdBannerSettingsMotdAndBanner ---
func unpackMotdBannerSettingsMotdAndBannerListToSdk(ctx context.Context, list types.List) ([]device_settings.MotdBannerSettingsMotdAndBanner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.MotdBannerSettingsMotdAndBanner")
	diags := diag.Diagnostics{}
	var data []models.MotdBannerSettingsMotdAndBanner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]device_settings.MotdBannerSettingsMotdAndBanner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.MotdBannerSettingsMotdAndBanner{}.AttrTypes(), &item)
		unpacked, d := unpackMotdBannerSettingsMotdAndBannerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.MotdBannerSettingsMotdAndBanner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for MotdBannerSettingsMotdAndBanner ---
func packMotdBannerSettingsMotdAndBannerListFromSdk(ctx context.Context, sdks []device_settings.MotdBannerSettingsMotdAndBanner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.MotdBannerSettingsMotdAndBanner")
	diags := diag.Diagnostics{}
	var data []models.MotdBannerSettingsMotdAndBanner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.MotdBannerSettingsMotdAndBanner
		obj, d := packMotdBannerSettingsMotdAndBannerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.MotdBannerSettingsMotdAndBanner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.MotdBannerSettingsMotdAndBanner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.MotdBannerSettingsMotdAndBanner{}.AttrType(), data)
}
