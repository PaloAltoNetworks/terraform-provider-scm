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

// --- Unpacker for ContentIdSettings ---
func unpackContentIdSettingsToSdk(ctx context.Context, obj types.Object) (*device_settings.ContentIdSettings, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ContentIdSettings", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ContentIdSettings
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk device_settings.ContentIdSettings
	var d diag.Diagnostics

	// Handling Objects
	if !model.ContentId.IsNull() && !model.ContentId.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field ContentId")
		unpacked, d := unpackContentIdSettingsContentIdToSdk(ctx, model.ContentId)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "ContentId"})
		}
		if unpacked != nil {
			sdk.ContentId = unpacked
		}
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
	if !model.Snippet.IsNull() && !model.Snippet.IsUnknown() {
		sdk.Snippet = model.Snippet.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ContentIdSettings", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ContentIdSettings ---
func packContentIdSettingsFromSdk(ctx context.Context, sdk device_settings.ContentIdSettings) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ContentIdSettings", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ContentIdSettings
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.ContentId != nil {
		tflog.Debug(ctx, "Packing nested object for field ContentId")
		packed, d := packContentIdSettingsContentIdFromSdk(ctx, *sdk.ContentId)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "ContentId"})
		}
		model.ContentId = packed
	} else {
		model.ContentId = basetypes.NewObjectNull(models.ContentIdSettingsContentId{}.AttrTypes())
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
	if sdk.Snippet != nil {
		model.Snippet = basetypes.NewStringValue(*sdk.Snippet)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	} else {
		model.Snippet = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ContentIdSettings{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ContentIdSettings", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ContentIdSettings ---
func unpackContentIdSettingsListToSdk(ctx context.Context, list types.List) ([]device_settings.ContentIdSettings, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ContentIdSettings")
	diags := diag.Diagnostics{}
	var data []models.ContentIdSettings
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]device_settings.ContentIdSettings, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ContentIdSettings{}.AttrTypes(), &item)
		unpacked, d := unpackContentIdSettingsToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ContentIdSettings", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ContentIdSettings ---
func packContentIdSettingsListFromSdk(ctx context.Context, sdks []device_settings.ContentIdSettings) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ContentIdSettings")
	diags := diag.Diagnostics{}
	var data []models.ContentIdSettings

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ContentIdSettings
		obj, d := packContentIdSettingsFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ContentIdSettings{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ContentIdSettings", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ContentIdSettings{}.AttrType(), data)
}

// --- Unpacker for ContentIdSettingsContentId ---
func unpackContentIdSettingsContentIdToSdk(ctx context.Context, obj types.Object) (*device_settings.ContentIdSettingsContentId, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ContentIdSettingsContentId", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ContentIdSettingsContentId
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk device_settings.ContentIdSettingsContentId
	var d diag.Diagnostics
	// Handling Primitives
	if !model.AllowForwardDecryptedContent.IsNull() && !model.AllowForwardDecryptedContent.IsUnknown() {
		sdk.AllowForwardDecryptedContent = model.AllowForwardDecryptedContent.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "AllowForwardDecryptedContent", "value": *sdk.AllowForwardDecryptedContent})
	}

	// Handling Primitives
	if !model.AllowHttpRange.IsNull() && !model.AllowHttpRange.IsUnknown() {
		sdk.AllowHttpRange = model.AllowHttpRange.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "AllowHttpRange", "value": *sdk.AllowHttpRange})
	}

	// Handling Objects
	if !model.Application.IsNull() && !model.Application.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Application")
		unpacked, d := unpackContentIdSettingsContentIdApplicationToSdk(ctx, model.Application)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Application"})
		}
		if unpacked != nil {
			sdk.Application = unpacked
		}
	}

	// Handling Primitives
	if !model.ExtendedCaptureSegment.IsNull() && !model.ExtendedCaptureSegment.IsUnknown() {
		val := int32(model.ExtendedCaptureSegment.ValueInt64())
		sdk.ExtendedCaptureSegment = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "ExtendedCaptureSegment", "value": *sdk.ExtendedCaptureSegment})
	}

	// Handling Primitives
	if !model.StripXFwdFor.IsNull() && !model.StripXFwdFor.IsUnknown() {
		sdk.StripXFwdFor = model.StripXFwdFor.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "StripXFwdFor", "value": *sdk.StripXFwdFor})
	}

	// Handling Primitives
	if !model.TcpBypassExceedQueue.IsNull() && !model.TcpBypassExceedQueue.IsUnknown() {
		sdk.TcpBypassExceedQueue = model.TcpBypassExceedQueue.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "TcpBypassExceedQueue", "value": *sdk.TcpBypassExceedQueue})
	}

	// Handling Primitives
	if !model.UdpBypassExceedQueue.IsNull() && !model.UdpBypassExceedQueue.IsUnknown() {
		sdk.UdpBypassExceedQueue = model.UdpBypassExceedQueue.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "UdpBypassExceedQueue", "value": *sdk.UdpBypassExceedQueue})
	}

	// Handling Primitives
	if !model.XForwardedFor.IsNull() && !model.XForwardedFor.IsUnknown() {
		sdk.XForwardedFor = model.XForwardedFor.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "XForwardedFor", "value": *sdk.XForwardedFor})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ContentIdSettingsContentId", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ContentIdSettingsContentId ---
func packContentIdSettingsContentIdFromSdk(ctx context.Context, sdk device_settings.ContentIdSettingsContentId) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ContentIdSettingsContentId", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ContentIdSettingsContentId
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.AllowForwardDecryptedContent != nil {
		model.AllowForwardDecryptedContent = basetypes.NewBoolValue(*sdk.AllowForwardDecryptedContent)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "AllowForwardDecryptedContent", "value": *sdk.AllowForwardDecryptedContent})
	} else {
		model.AllowForwardDecryptedContent = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.AllowHttpRange != nil {
		model.AllowHttpRange = basetypes.NewBoolValue(*sdk.AllowHttpRange)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "AllowHttpRange", "value": *sdk.AllowHttpRange})
	} else {
		model.AllowHttpRange = basetypes.NewBoolNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Application != nil {
		tflog.Debug(ctx, "Packing nested object for field Application")
		packed, d := packContentIdSettingsContentIdApplicationFromSdk(ctx, *sdk.Application)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Application"})
		}
		model.Application = packed
	} else {
		model.Application = basetypes.NewObjectNull(models.ContentIdSettingsContentIdApplication{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.ExtendedCaptureSegment != nil {
		model.ExtendedCaptureSegment = basetypes.NewInt64Value(int64(*sdk.ExtendedCaptureSegment))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "ExtendedCaptureSegment", "value": *sdk.ExtendedCaptureSegment})
	} else {
		model.ExtendedCaptureSegment = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.StripXFwdFor != nil {
		model.StripXFwdFor = basetypes.NewBoolValue(*sdk.StripXFwdFor)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "StripXFwdFor", "value": *sdk.StripXFwdFor})
	} else {
		model.StripXFwdFor = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.TcpBypassExceedQueue != nil {
		model.TcpBypassExceedQueue = basetypes.NewBoolValue(*sdk.TcpBypassExceedQueue)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "TcpBypassExceedQueue", "value": *sdk.TcpBypassExceedQueue})
	} else {
		model.TcpBypassExceedQueue = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.UdpBypassExceedQueue != nil {
		model.UdpBypassExceedQueue = basetypes.NewBoolValue(*sdk.UdpBypassExceedQueue)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "UdpBypassExceedQueue", "value": *sdk.UdpBypassExceedQueue})
	} else {
		model.UdpBypassExceedQueue = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.XForwardedFor != nil {
		model.XForwardedFor = basetypes.NewStringValue(*sdk.XForwardedFor)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "XForwardedFor", "value": *sdk.XForwardedFor})
	} else {
		model.XForwardedFor = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ContentIdSettingsContentId{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ContentIdSettingsContentId", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ContentIdSettingsContentId ---
func unpackContentIdSettingsContentIdListToSdk(ctx context.Context, list types.List) ([]device_settings.ContentIdSettingsContentId, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ContentIdSettingsContentId")
	diags := diag.Diagnostics{}
	var data []models.ContentIdSettingsContentId
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]device_settings.ContentIdSettingsContentId, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ContentIdSettingsContentId{}.AttrTypes(), &item)
		unpacked, d := unpackContentIdSettingsContentIdToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ContentIdSettingsContentId", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ContentIdSettingsContentId ---
func packContentIdSettingsContentIdListFromSdk(ctx context.Context, sdks []device_settings.ContentIdSettingsContentId) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ContentIdSettingsContentId")
	diags := diag.Diagnostics{}
	var data []models.ContentIdSettingsContentId

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ContentIdSettingsContentId
		obj, d := packContentIdSettingsContentIdFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ContentIdSettingsContentId{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ContentIdSettingsContentId", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ContentIdSettingsContentId{}.AttrType(), data)
}

// --- Unpacker for ContentIdSettingsContentIdApplication ---
func unpackContentIdSettingsContentIdApplicationToSdk(ctx context.Context, obj types.Object) (*device_settings.ContentIdSettingsContentIdApplication, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ContentIdSettingsContentIdApplication", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ContentIdSettingsContentIdApplication
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk device_settings.ContentIdSettingsContentIdApplication
	var d diag.Diagnostics
	// Handling Primitives
	if !model.BypassExceedQueue.IsNull() && !model.BypassExceedQueue.IsUnknown() {
		sdk.BypassExceedQueue = model.BypassExceedQueue.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "BypassExceedQueue", "value": *sdk.BypassExceedQueue})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ContentIdSettingsContentIdApplication", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ContentIdSettingsContentIdApplication ---
func packContentIdSettingsContentIdApplicationFromSdk(ctx context.Context, sdk device_settings.ContentIdSettingsContentIdApplication) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ContentIdSettingsContentIdApplication", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ContentIdSettingsContentIdApplication
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.BypassExceedQueue != nil {
		model.BypassExceedQueue = basetypes.NewBoolValue(*sdk.BypassExceedQueue)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "BypassExceedQueue", "value": *sdk.BypassExceedQueue})
	} else {
		model.BypassExceedQueue = basetypes.NewBoolNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ContentIdSettingsContentIdApplication{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ContentIdSettingsContentIdApplication", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ContentIdSettingsContentIdApplication ---
func unpackContentIdSettingsContentIdApplicationListToSdk(ctx context.Context, list types.List) ([]device_settings.ContentIdSettingsContentIdApplication, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ContentIdSettingsContentIdApplication")
	diags := diag.Diagnostics{}
	var data []models.ContentIdSettingsContentIdApplication
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]device_settings.ContentIdSettingsContentIdApplication, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ContentIdSettingsContentIdApplication{}.AttrTypes(), &item)
		unpacked, d := unpackContentIdSettingsContentIdApplicationToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ContentIdSettingsContentIdApplication", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ContentIdSettingsContentIdApplication ---
func packContentIdSettingsContentIdApplicationListFromSdk(ctx context.Context, sdks []device_settings.ContentIdSettingsContentIdApplication) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ContentIdSettingsContentIdApplication")
	diags := diag.Diagnostics{}
	var data []models.ContentIdSettingsContentIdApplication

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ContentIdSettingsContentIdApplication
		obj, d := packContentIdSettingsContentIdApplicationFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ContentIdSettingsContentIdApplication{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ContentIdSettingsContentIdApplication", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ContentIdSettingsContentIdApplication{}.AttrType(), data)
}
