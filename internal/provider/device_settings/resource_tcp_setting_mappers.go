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

// --- Unpacker for TcpSettings ---
func unpackTcpSettingsToSdk(ctx context.Context, obj types.Object) (*device_settings.TcpSettings, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.TcpSettings", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.TcpSettings
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk device_settings.TcpSettings
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
	if !model.Tcp.IsNull() && !model.Tcp.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Tcp")
		unpacked, d := unpackTcpSettingsTcpToSdk(ctx, model.Tcp)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Tcp"})
		}
		if unpacked != nil {
			sdk.Tcp = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.TcpSettings", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for TcpSettings ---
func packTcpSettingsFromSdk(ctx context.Context, sdk device_settings.TcpSettings) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.TcpSettings", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.TcpSettings
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
	if sdk.Tcp != nil {
		tflog.Debug(ctx, "Packing nested object for field Tcp")
		packed, d := packTcpSettingsTcpFromSdk(ctx, *sdk.Tcp)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Tcp"})
		}
		model.Tcp = packed
	} else {
		model.Tcp = basetypes.NewObjectNull(models.TcpSettingsTcp{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.TcpSettings{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.TcpSettings", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for TcpSettings ---
func unpackTcpSettingsListToSdk(ctx context.Context, list types.List) ([]device_settings.TcpSettings, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.TcpSettings")
	diags := diag.Diagnostics{}
	var data []models.TcpSettings
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]device_settings.TcpSettings, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.TcpSettings{}.AttrTypes(), &item)
		unpacked, d := unpackTcpSettingsToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.TcpSettings", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for TcpSettings ---
func packTcpSettingsListFromSdk(ctx context.Context, sdks []device_settings.TcpSettings) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.TcpSettings")
	diags := diag.Diagnostics{}
	var data []models.TcpSettings

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.TcpSettings
		obj, d := packTcpSettingsFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.TcpSettings{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.TcpSettings", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.TcpSettings{}.AttrType(), data)
}

// --- Unpacker for TcpSettingsTcp ---
func unpackTcpSettingsTcpToSdk(ctx context.Context, obj types.Object) (*device_settings.TcpSettingsTcp, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.TcpSettingsTcp", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.TcpSettingsTcp
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk device_settings.TcpSettingsTcp
	var d diag.Diagnostics
	// Handling Primitives
	if !model.AllowChallengeAck.IsNull() && !model.AllowChallengeAck.IsUnknown() {
		sdk.AllowChallengeAck = model.AllowChallengeAck.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "AllowChallengeAck", "value": *sdk.AllowChallengeAck})
	}

	// Handling Primitives
	if !model.AsymmetricPath.IsNull() && !model.AsymmetricPath.IsUnknown() {
		sdk.AsymmetricPath = model.AsymmetricPath.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "AsymmetricPath", "value": *sdk.AsymmetricPath})
	}

	// Handling Primitives
	if !model.BypassExceedOoQueue.IsNull() && !model.BypassExceedOoQueue.IsUnknown() {
		sdk.BypassExceedOoQueue = model.BypassExceedOoQueue.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "BypassExceedOoQueue", "value": *sdk.BypassExceedOoQueue})
	}

	// Handling Primitives
	if !model.CheckTimestampOption.IsNull() && !model.CheckTimestampOption.IsUnknown() {
		sdk.CheckTimestampOption = model.CheckTimestampOption.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "CheckTimestampOption", "value": *sdk.CheckTimestampOption})
	}

	// Handling Primitives
	if !model.DropZeroFlag.IsNull() && !model.DropZeroFlag.IsUnknown() {
		sdk.DropZeroFlag = model.DropZeroFlag.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DropZeroFlag", "value": *sdk.DropZeroFlag})
	}

	// Handling Primitives
	if !model.SiptcpCleartextProxy.IsNull() && !model.SiptcpCleartextProxy.IsUnknown() {
		sdk.SiptcpCleartextProxy = model.SiptcpCleartextProxy.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "SiptcpCleartextProxy", "value": *sdk.SiptcpCleartextProxy})
	}

	// Handling Primitives
	if !model.StripMptcpOption.IsNull() && !model.StripMptcpOption.IsUnknown() {
		sdk.StripMptcpOption = model.StripMptcpOption.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "StripMptcpOption", "value": *sdk.StripMptcpOption})
	}

	// Handling Primitives
	if !model.TcpRetransmitScan.IsNull() && !model.TcpRetransmitScan.IsUnknown() {
		sdk.TcpRetransmitScan = model.TcpRetransmitScan.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "TcpRetransmitScan", "value": *sdk.TcpRetransmitScan})
	}

	// Handling Primitives
	if !model.UrgentData.IsNull() && !model.UrgentData.IsUnknown() {
		sdk.UrgentData = model.UrgentData.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "UrgentData", "value": *sdk.UrgentData})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.TcpSettingsTcp", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for TcpSettingsTcp ---
func packTcpSettingsTcpFromSdk(ctx context.Context, sdk device_settings.TcpSettingsTcp) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.TcpSettingsTcp", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.TcpSettingsTcp
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.AllowChallengeAck != nil {
		model.AllowChallengeAck = basetypes.NewBoolValue(*sdk.AllowChallengeAck)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "AllowChallengeAck", "value": *sdk.AllowChallengeAck})
	} else {
		model.AllowChallengeAck = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.AsymmetricPath != nil {
		model.AsymmetricPath = basetypes.NewStringValue(*sdk.AsymmetricPath)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "AsymmetricPath", "value": *sdk.AsymmetricPath})
	} else {
		model.AsymmetricPath = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.BypassExceedOoQueue != nil {
		model.BypassExceedOoQueue = basetypes.NewBoolValue(*sdk.BypassExceedOoQueue)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "BypassExceedOoQueue", "value": *sdk.BypassExceedOoQueue})
	} else {
		model.BypassExceedOoQueue = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.CheckTimestampOption != nil {
		model.CheckTimestampOption = basetypes.NewBoolValue(*sdk.CheckTimestampOption)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "CheckTimestampOption", "value": *sdk.CheckTimestampOption})
	} else {
		model.CheckTimestampOption = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.DropZeroFlag != nil {
		model.DropZeroFlag = basetypes.NewBoolValue(*sdk.DropZeroFlag)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DropZeroFlag", "value": *sdk.DropZeroFlag})
	} else {
		model.DropZeroFlag = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.SiptcpCleartextProxy != nil {
		model.SiptcpCleartextProxy = basetypes.NewStringValue(*sdk.SiptcpCleartextProxy)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "SiptcpCleartextProxy", "value": *sdk.SiptcpCleartextProxy})
	} else {
		model.SiptcpCleartextProxy = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.StripMptcpOption != nil {
		model.StripMptcpOption = basetypes.NewBoolValue(*sdk.StripMptcpOption)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "StripMptcpOption", "value": *sdk.StripMptcpOption})
	} else {
		model.StripMptcpOption = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.TcpRetransmitScan != nil {
		model.TcpRetransmitScan = basetypes.NewBoolValue(*sdk.TcpRetransmitScan)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "TcpRetransmitScan", "value": *sdk.TcpRetransmitScan})
	} else {
		model.TcpRetransmitScan = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.UrgentData != nil {
		model.UrgentData = basetypes.NewStringValue(*sdk.UrgentData)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "UrgentData", "value": *sdk.UrgentData})
	} else {
		model.UrgentData = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.TcpSettingsTcp{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.TcpSettingsTcp", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for TcpSettingsTcp ---
func unpackTcpSettingsTcpListToSdk(ctx context.Context, list types.List) ([]device_settings.TcpSettingsTcp, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.TcpSettingsTcp")
	diags := diag.Diagnostics{}
	var data []models.TcpSettingsTcp
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]device_settings.TcpSettingsTcp, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.TcpSettingsTcp{}.AttrTypes(), &item)
		unpacked, d := unpackTcpSettingsTcpToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.TcpSettingsTcp", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for TcpSettingsTcp ---
func packTcpSettingsTcpListFromSdk(ctx context.Context, sdks []device_settings.TcpSettingsTcp) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.TcpSettingsTcp")
	diags := diag.Diagnostics{}
	var data []models.TcpSettingsTcp

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.TcpSettingsTcp
		obj, d := packTcpSettingsTcpFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.TcpSettingsTcp{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.TcpSettingsTcp", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.TcpSettingsTcp{}.AttrType(), data)
}
