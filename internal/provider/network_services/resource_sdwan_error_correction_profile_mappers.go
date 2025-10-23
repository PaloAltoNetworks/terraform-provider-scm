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

// --- Unpacker for SdwanErrorCorrectionProfiles ---
func unpackSdwanErrorCorrectionProfilesToSdk(ctx context.Context, obj types.Object) (*network_services.SdwanErrorCorrectionProfiles, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.SdwanErrorCorrectionProfiles", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.SdwanErrorCorrectionProfiles
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.SdwanErrorCorrectionProfiles
	var d diag.Diagnostics

	// Handling Primitives
	if !model.ActivationThreshold.IsNull() && !model.ActivationThreshold.IsUnknown() {
		sdk.ActivationThreshold = float32(model.ActivationThreshold.ValueFloat64())
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "ActivationThreshold", "value": sdk.ActivationThreshold})
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

	// Handling Objects
	if !model.Mode.IsNull() && !model.Mode.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Mode")
		unpacked, d := unpackSdwanErrorCorrectionProfilesModeToSdk(ctx, model.Mode)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Mode"})
		}
		if unpacked != nil {
			sdk.Mode = *unpacked
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

	tflog.Debug(ctx, "Exiting unpack helper for models.SdwanErrorCorrectionProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for SdwanErrorCorrectionProfiles ---
func packSdwanErrorCorrectionProfilesFromSdk(ctx context.Context, sdk network_services.SdwanErrorCorrectionProfiles) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.SdwanErrorCorrectionProfiles", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.SdwanErrorCorrectionProfiles
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.ActivationThreshold = basetypes.NewFloat64Value(float64(sdk.ActivationThreshold))
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "ActivationThreshold", "value": sdk.ActivationThreshold})
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
	if !reflect.ValueOf(sdk.Mode).IsZero() {
		tflog.Debug(ctx, "Packing nested object for field Mode")
		packed, d := packSdwanErrorCorrectionProfilesModeFromSdk(ctx, sdk.Mode)
		diags.Append(d...)
		model.Mode = packed
	} else {
		model.Mode = basetypes.NewObjectNull(models.SdwanErrorCorrectionProfilesMode{}.AttrTypes())
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

	obj, d := types.ObjectValueFrom(ctx, models.SdwanErrorCorrectionProfiles{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.SdwanErrorCorrectionProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for SdwanErrorCorrectionProfiles ---
func unpackSdwanErrorCorrectionProfilesListToSdk(ctx context.Context, list types.List) ([]network_services.SdwanErrorCorrectionProfiles, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.SdwanErrorCorrectionProfiles")
	diags := diag.Diagnostics{}
	var data []models.SdwanErrorCorrectionProfiles
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.SdwanErrorCorrectionProfiles, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.SdwanErrorCorrectionProfiles{}.AttrTypes(), &item)
		unpacked, d := unpackSdwanErrorCorrectionProfilesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.SdwanErrorCorrectionProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for SdwanErrorCorrectionProfiles ---
func packSdwanErrorCorrectionProfilesListFromSdk(ctx context.Context, sdks []network_services.SdwanErrorCorrectionProfiles) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.SdwanErrorCorrectionProfiles")
	diags := diag.Diagnostics{}
	var data []models.SdwanErrorCorrectionProfiles

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.SdwanErrorCorrectionProfiles
		obj, d := packSdwanErrorCorrectionProfilesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.SdwanErrorCorrectionProfiles{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.SdwanErrorCorrectionProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.SdwanErrorCorrectionProfiles{}.AttrType(), data)
}

// --- Unpacker for SdwanErrorCorrectionProfilesMode ---
func unpackSdwanErrorCorrectionProfilesModeToSdk(ctx context.Context, obj types.Object) (*network_services.SdwanErrorCorrectionProfilesMode, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.SdwanErrorCorrectionProfilesMode", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.SdwanErrorCorrectionProfilesMode
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.SdwanErrorCorrectionProfilesMode
	var d diag.Diagnostics
	// Handling Objects
	if !model.ForwardErrorCorrection.IsNull() && !model.ForwardErrorCorrection.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field ForwardErrorCorrection")
		unpacked, d := unpackSdwanErrorCorrectionProfilesModeForwardErrorCorrectionToSdk(ctx, model.ForwardErrorCorrection)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "ForwardErrorCorrection"})
		}
		if unpacked != nil {
			sdk.ForwardErrorCorrection = unpacked
		}
	}

	// Handling Objects
	if !model.PacketDuplication.IsNull() && !model.PacketDuplication.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field PacketDuplication")
		unpacked, d := unpackSdwanErrorCorrectionProfilesModePacketDuplicationToSdk(ctx, model.PacketDuplication)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "PacketDuplication"})
		}
		if unpacked != nil {
			sdk.PacketDuplication = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.SdwanErrorCorrectionProfilesMode", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for SdwanErrorCorrectionProfilesMode ---
func packSdwanErrorCorrectionProfilesModeFromSdk(ctx context.Context, sdk network_services.SdwanErrorCorrectionProfilesMode) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.SdwanErrorCorrectionProfilesMode", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.SdwanErrorCorrectionProfilesMode
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.ForwardErrorCorrection != nil {
		tflog.Debug(ctx, "Packing nested object for field ForwardErrorCorrection")
		packed, d := packSdwanErrorCorrectionProfilesModeForwardErrorCorrectionFromSdk(ctx, *sdk.ForwardErrorCorrection)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "ForwardErrorCorrection"})
		}
		model.ForwardErrorCorrection = packed
	} else {
		model.ForwardErrorCorrection = basetypes.NewObjectNull(models.SdwanErrorCorrectionProfilesModeForwardErrorCorrection{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.PacketDuplication != nil {
		tflog.Debug(ctx, "Packing nested object for field PacketDuplication")
		packed, d := packSdwanErrorCorrectionProfilesModePacketDuplicationFromSdk(ctx, *sdk.PacketDuplication)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "PacketDuplication"})
		}
		model.PacketDuplication = packed
	} else {
		model.PacketDuplication = basetypes.NewObjectNull(models.SdwanErrorCorrectionProfilesModePacketDuplication{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.SdwanErrorCorrectionProfilesMode{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.SdwanErrorCorrectionProfilesMode", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for SdwanErrorCorrectionProfilesMode ---
func unpackSdwanErrorCorrectionProfilesModeListToSdk(ctx context.Context, list types.List) ([]network_services.SdwanErrorCorrectionProfilesMode, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.SdwanErrorCorrectionProfilesMode")
	diags := diag.Diagnostics{}
	var data []models.SdwanErrorCorrectionProfilesMode
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.SdwanErrorCorrectionProfilesMode, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.SdwanErrorCorrectionProfilesMode{}.AttrTypes(), &item)
		unpacked, d := unpackSdwanErrorCorrectionProfilesModeToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.SdwanErrorCorrectionProfilesMode", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for SdwanErrorCorrectionProfilesMode ---
func packSdwanErrorCorrectionProfilesModeListFromSdk(ctx context.Context, sdks []network_services.SdwanErrorCorrectionProfilesMode) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.SdwanErrorCorrectionProfilesMode")
	diags := diag.Diagnostics{}
	var data []models.SdwanErrorCorrectionProfilesMode

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.SdwanErrorCorrectionProfilesMode
		obj, d := packSdwanErrorCorrectionProfilesModeFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.SdwanErrorCorrectionProfilesMode{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.SdwanErrorCorrectionProfilesMode", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.SdwanErrorCorrectionProfilesMode{}.AttrType(), data)
}

// --- Unpacker for SdwanErrorCorrectionProfilesModeForwardErrorCorrection ---
func unpackSdwanErrorCorrectionProfilesModeForwardErrorCorrectionToSdk(ctx context.Context, obj types.Object) (*network_services.SdwanErrorCorrectionProfilesModeForwardErrorCorrection, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.SdwanErrorCorrectionProfilesModeForwardErrorCorrection", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.SdwanErrorCorrectionProfilesModeForwardErrorCorrection
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.SdwanErrorCorrectionProfilesModeForwardErrorCorrection
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Ratio.IsNull() && !model.Ratio.IsUnknown() {
		sdk.Ratio = model.Ratio.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Ratio", "value": sdk.Ratio})
	}

	// Handling Primitives
	if !model.RecoveryDuration.IsNull() && !model.RecoveryDuration.IsUnknown() {
		sdk.RecoveryDuration = float32(model.RecoveryDuration.ValueFloat64())
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "RecoveryDuration", "value": sdk.RecoveryDuration})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.SdwanErrorCorrectionProfilesModeForwardErrorCorrection", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for SdwanErrorCorrectionProfilesModeForwardErrorCorrection ---
func packSdwanErrorCorrectionProfilesModeForwardErrorCorrectionFromSdk(ctx context.Context, sdk network_services.SdwanErrorCorrectionProfilesModeForwardErrorCorrection) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.SdwanErrorCorrectionProfilesModeForwardErrorCorrection", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.SdwanErrorCorrectionProfilesModeForwardErrorCorrection
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.Ratio = basetypes.NewStringValue(sdk.Ratio)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Ratio", "value": sdk.Ratio})
	// Handling Primitives
	// Standard primitive packing
	model.RecoveryDuration = basetypes.NewFloat64Value(float64(sdk.RecoveryDuration))
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "RecoveryDuration", "value": sdk.RecoveryDuration})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.SdwanErrorCorrectionProfilesModeForwardErrorCorrection{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.SdwanErrorCorrectionProfilesModeForwardErrorCorrection", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for SdwanErrorCorrectionProfilesModeForwardErrorCorrection ---
func unpackSdwanErrorCorrectionProfilesModeForwardErrorCorrectionListToSdk(ctx context.Context, list types.List) ([]network_services.SdwanErrorCorrectionProfilesModeForwardErrorCorrection, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.SdwanErrorCorrectionProfilesModeForwardErrorCorrection")
	diags := diag.Diagnostics{}
	var data []models.SdwanErrorCorrectionProfilesModeForwardErrorCorrection
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.SdwanErrorCorrectionProfilesModeForwardErrorCorrection, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.SdwanErrorCorrectionProfilesModeForwardErrorCorrection{}.AttrTypes(), &item)
		unpacked, d := unpackSdwanErrorCorrectionProfilesModeForwardErrorCorrectionToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.SdwanErrorCorrectionProfilesModeForwardErrorCorrection", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for SdwanErrorCorrectionProfilesModeForwardErrorCorrection ---
func packSdwanErrorCorrectionProfilesModeForwardErrorCorrectionListFromSdk(ctx context.Context, sdks []network_services.SdwanErrorCorrectionProfilesModeForwardErrorCorrection) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.SdwanErrorCorrectionProfilesModeForwardErrorCorrection")
	diags := diag.Diagnostics{}
	var data []models.SdwanErrorCorrectionProfilesModeForwardErrorCorrection

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.SdwanErrorCorrectionProfilesModeForwardErrorCorrection
		obj, d := packSdwanErrorCorrectionProfilesModeForwardErrorCorrectionFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.SdwanErrorCorrectionProfilesModeForwardErrorCorrection{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.SdwanErrorCorrectionProfilesModeForwardErrorCorrection", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.SdwanErrorCorrectionProfilesModeForwardErrorCorrection{}.AttrType(), data)
}

// --- Unpacker for SdwanErrorCorrectionProfilesModePacketDuplication ---
func unpackSdwanErrorCorrectionProfilesModePacketDuplicationToSdk(ctx context.Context, obj types.Object) (*network_services.SdwanErrorCorrectionProfilesModePacketDuplication, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.SdwanErrorCorrectionProfilesModePacketDuplication", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.SdwanErrorCorrectionProfilesModePacketDuplication
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.SdwanErrorCorrectionProfilesModePacketDuplication
	var d diag.Diagnostics
	// Handling Primitives
	if !model.RecoveryDurationPd.IsNull() && !model.RecoveryDurationPd.IsUnknown() {
		sdk.RecoveryDurationPd = float32(model.RecoveryDurationPd.ValueFloat64())
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "RecoveryDurationPd", "value": sdk.RecoveryDurationPd})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.SdwanErrorCorrectionProfilesModePacketDuplication", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for SdwanErrorCorrectionProfilesModePacketDuplication ---
func packSdwanErrorCorrectionProfilesModePacketDuplicationFromSdk(ctx context.Context, sdk network_services.SdwanErrorCorrectionProfilesModePacketDuplication) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.SdwanErrorCorrectionProfilesModePacketDuplication", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.SdwanErrorCorrectionProfilesModePacketDuplication
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.RecoveryDurationPd = basetypes.NewFloat64Value(float64(sdk.RecoveryDurationPd))
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "RecoveryDurationPd", "value": sdk.RecoveryDurationPd})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.SdwanErrorCorrectionProfilesModePacketDuplication{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.SdwanErrorCorrectionProfilesModePacketDuplication", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for SdwanErrorCorrectionProfilesModePacketDuplication ---
func unpackSdwanErrorCorrectionProfilesModePacketDuplicationListToSdk(ctx context.Context, list types.List) ([]network_services.SdwanErrorCorrectionProfilesModePacketDuplication, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.SdwanErrorCorrectionProfilesModePacketDuplication")
	diags := diag.Diagnostics{}
	var data []models.SdwanErrorCorrectionProfilesModePacketDuplication
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.SdwanErrorCorrectionProfilesModePacketDuplication, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.SdwanErrorCorrectionProfilesModePacketDuplication{}.AttrTypes(), &item)
		unpacked, d := unpackSdwanErrorCorrectionProfilesModePacketDuplicationToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.SdwanErrorCorrectionProfilesModePacketDuplication", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for SdwanErrorCorrectionProfilesModePacketDuplication ---
func packSdwanErrorCorrectionProfilesModePacketDuplicationListFromSdk(ctx context.Context, sdks []network_services.SdwanErrorCorrectionProfilesModePacketDuplication) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.SdwanErrorCorrectionProfilesModePacketDuplication")
	diags := diag.Diagnostics{}
	var data []models.SdwanErrorCorrectionProfilesModePacketDuplication

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.SdwanErrorCorrectionProfilesModePacketDuplication
		obj, d := packSdwanErrorCorrectionProfilesModePacketDuplicationFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.SdwanErrorCorrectionProfilesModePacketDuplication{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.SdwanErrorCorrectionProfilesModePacketDuplication", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.SdwanErrorCorrectionProfilesModePacketDuplication{}.AttrType(), data)
}
