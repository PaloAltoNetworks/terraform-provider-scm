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

// --- Unpacker for DeviceRedistributionCollector ---
func unpackDeviceRedistributionCollectorToSdk(ctx context.Context, obj types.Object) (*device_settings.DeviceRedistributionCollector, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.DeviceRedistributionCollector", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.DeviceRedistributionCollector
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk device_settings.DeviceRedistributionCollector
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
	if !model.RedistributionCollector.IsNull() && !model.RedistributionCollector.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field RedistributionCollector")
		unpacked, d := unpackDeviceRedistributionCollectorRedistributionCollectorToSdk(ctx, model.RedistributionCollector)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "RedistributionCollector"})
		}
		if unpacked != nil {
			sdk.RedistributionCollector = unpacked
		}
	}

	// Handling Primitives
	if !model.Snippet.IsNull() && !model.Snippet.IsUnknown() {
		sdk.Snippet = model.Snippet.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.DeviceRedistributionCollector", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for DeviceRedistributionCollector ---
func packDeviceRedistributionCollectorFromSdk(ctx context.Context, sdk device_settings.DeviceRedistributionCollector) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.DeviceRedistributionCollector", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.DeviceRedistributionCollector
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
	if sdk.RedistributionCollector != nil {
		tflog.Debug(ctx, "Packing nested object for field RedistributionCollector")
		packed, d := packDeviceRedistributionCollectorRedistributionCollectorFromSdk(ctx, *sdk.RedistributionCollector)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "RedistributionCollector"})
		}
		model.RedistributionCollector = packed
	} else {
		model.RedistributionCollector = basetypes.NewObjectNull(models.DeviceRedistributionCollectorRedistributionCollector{}.AttrTypes())
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

	obj, d := types.ObjectValueFrom(ctx, models.DeviceRedistributionCollector{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.DeviceRedistributionCollector", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for DeviceRedistributionCollector ---
func unpackDeviceRedistributionCollectorListToSdk(ctx context.Context, list types.List) ([]device_settings.DeviceRedistributionCollector, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.DeviceRedistributionCollector")
	diags := diag.Diagnostics{}
	var data []models.DeviceRedistributionCollector
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]device_settings.DeviceRedistributionCollector, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.DeviceRedistributionCollector{}.AttrTypes(), &item)
		unpacked, d := unpackDeviceRedistributionCollectorToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.DeviceRedistributionCollector", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for DeviceRedistributionCollector ---
func packDeviceRedistributionCollectorListFromSdk(ctx context.Context, sdks []device_settings.DeviceRedistributionCollector) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.DeviceRedistributionCollector")
	diags := diag.Diagnostics{}
	var data []models.DeviceRedistributionCollector

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.DeviceRedistributionCollector
		obj, d := packDeviceRedistributionCollectorFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.DeviceRedistributionCollector{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.DeviceRedistributionCollector", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.DeviceRedistributionCollector{}.AttrType(), data)
}

// --- Unpacker for DeviceRedistributionCollectorRedistributionCollector ---
func unpackDeviceRedistributionCollectorRedistributionCollectorToSdk(ctx context.Context, obj types.Object) (*device_settings.DeviceRedistributionCollectorRedistributionCollector, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.DeviceRedistributionCollectorRedistributionCollector", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.DeviceRedistributionCollectorRedistributionCollector
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk device_settings.DeviceRedistributionCollectorRedistributionCollector
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Interface.IsNull() && !model.Interface.IsUnknown() {
		sdk.Interface = model.Interface.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Interface", "value": *sdk.Interface})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.DeviceRedistributionCollectorRedistributionCollector", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for DeviceRedistributionCollectorRedistributionCollector ---
func packDeviceRedistributionCollectorRedistributionCollectorFromSdk(ctx context.Context, sdk device_settings.DeviceRedistributionCollectorRedistributionCollector) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.DeviceRedistributionCollectorRedistributionCollector", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.DeviceRedistributionCollectorRedistributionCollector
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Interface != nil {
		model.Interface = basetypes.NewStringValue(*sdk.Interface)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Interface", "value": *sdk.Interface})
	} else {
		model.Interface = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.DeviceRedistributionCollectorRedistributionCollector{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.DeviceRedistributionCollectorRedistributionCollector", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for DeviceRedistributionCollectorRedistributionCollector ---
func unpackDeviceRedistributionCollectorRedistributionCollectorListToSdk(ctx context.Context, list types.List) ([]device_settings.DeviceRedistributionCollectorRedistributionCollector, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.DeviceRedistributionCollectorRedistributionCollector")
	diags := diag.Diagnostics{}
	var data []models.DeviceRedistributionCollectorRedistributionCollector
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]device_settings.DeviceRedistributionCollectorRedistributionCollector, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.DeviceRedistributionCollectorRedistributionCollector{}.AttrTypes(), &item)
		unpacked, d := unpackDeviceRedistributionCollectorRedistributionCollectorToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.DeviceRedistributionCollectorRedistributionCollector", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for DeviceRedistributionCollectorRedistributionCollector ---
func packDeviceRedistributionCollectorRedistributionCollectorListFromSdk(ctx context.Context, sdks []device_settings.DeviceRedistributionCollectorRedistributionCollector) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.DeviceRedistributionCollectorRedistributionCollector")
	diags := diag.Diagnostics{}
	var data []models.DeviceRedistributionCollectorRedistributionCollector

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.DeviceRedistributionCollectorRedistributionCollector
		obj, d := packDeviceRedistributionCollectorRedistributionCollectorFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.DeviceRedistributionCollectorRedistributionCollector{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.DeviceRedistributionCollectorRedistributionCollector", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.DeviceRedistributionCollectorRedistributionCollector{}.AttrType(), data)
}
