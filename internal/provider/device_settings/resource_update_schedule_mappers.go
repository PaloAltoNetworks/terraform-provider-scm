package provider

import (
	"context"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/device_settings"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/device_settings"
)

// --- Unpacker for UpdateSchedule ---
func unpackUpdateScheduleToSdk(ctx context.Context, obj types.Object) (*device_settings.UpdateSchedule, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.UpdateSchedule", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.UpdateSchedule
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk device_settings.UpdateSchedule
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
	if !model.UpdateSchedule.IsNull() && !model.UpdateSchedule.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field UpdateSchedule")
		unpacked, d := unpackUpdateScheduleUpdateScheduleToSdk(ctx, model.UpdateSchedule)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "UpdateSchedule"})
		}
		if unpacked != nil {
			sdk.UpdateSchedule = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.UpdateSchedule", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for UpdateSchedule ---
func packUpdateScheduleFromSdk(ctx context.Context, sdk device_settings.UpdateSchedule) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.UpdateSchedule", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.UpdateSchedule
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
	if sdk.UpdateSchedule != nil {
		tflog.Debug(ctx, "Packing nested object for field UpdateSchedule")
		packed, d := packUpdateScheduleUpdateScheduleFromSdk(ctx, *sdk.UpdateSchedule)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "UpdateSchedule"})
		}
		model.UpdateSchedule = packed
	} else {
		model.UpdateSchedule = basetypes.NewObjectNull(models.UpdateScheduleUpdateSchedule{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.UpdateSchedule{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.UpdateSchedule", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for UpdateSchedule ---
func unpackUpdateScheduleListToSdk(ctx context.Context, list types.List) ([]device_settings.UpdateSchedule, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.UpdateSchedule")
	diags := diag.Diagnostics{}
	var data []models.UpdateSchedule
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]device_settings.UpdateSchedule, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.UpdateSchedule{}.AttrTypes(), &item)
		unpacked, d := unpackUpdateScheduleToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.UpdateSchedule", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for UpdateSchedule ---
func packUpdateScheduleListFromSdk(ctx context.Context, sdks []device_settings.UpdateSchedule) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.UpdateSchedule")
	diags := diag.Diagnostics{}
	var data []models.UpdateSchedule

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.UpdateSchedule
		obj, d := packUpdateScheduleFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.UpdateSchedule{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.UpdateSchedule", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.UpdateSchedule{}.AttrType(), data)
}

// --- Unpacker for UpdateScheduleUpdateSchedule ---
func unpackUpdateScheduleUpdateScheduleToSdk(ctx context.Context, obj types.Object) (*device_settings.UpdateScheduleUpdateSchedule, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.UpdateScheduleUpdateSchedule", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.UpdateScheduleUpdateSchedule
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk device_settings.UpdateScheduleUpdateSchedule
	var d diag.Diagnostics
	// Handling Objects
	if !model.AntiVirus.IsNull() && !model.AntiVirus.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field AntiVirus")
		unpacked, d := unpackUpdateScheduleUpdateScheduleAntiVirusToSdk(ctx, model.AntiVirus)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "AntiVirus"})
		}
		if unpacked != nil {
			sdk.AntiVirus = *unpacked
		}
	}

	// Handling Objects
	if !model.Threats.IsNull() && !model.Threats.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Threats")
		unpacked, d := unpackUpdateScheduleUpdateScheduleThreatsToSdk(ctx, model.Threats)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Threats"})
		}
		if unpacked != nil {
			sdk.Threats = *unpacked
		}
	}

	// Handling Objects
	if !model.Wildfire.IsNull() && !model.Wildfire.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Wildfire")
		unpacked, d := unpackUpdateScheduleUpdateScheduleWildfireToSdk(ctx, model.Wildfire)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Wildfire"})
		}
		if unpacked != nil {
			sdk.Wildfire = *unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.UpdateScheduleUpdateSchedule", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for UpdateScheduleUpdateSchedule ---
func packUpdateScheduleUpdateScheduleFromSdk(ctx context.Context, sdk device_settings.UpdateScheduleUpdateSchedule) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.UpdateScheduleUpdateSchedule", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.UpdateScheduleUpdateSchedule
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	// Logic for non-pointer / value-type nested objects
	if !reflect.ValueOf(sdk.AntiVirus).IsZero() {
		tflog.Debug(ctx, "Packing nested object for field AntiVirus")
		packed, d := packUpdateScheduleUpdateScheduleAntiVirusFromSdk(ctx, sdk.AntiVirus)
		diags.Append(d...)
		model.AntiVirus = packed
	} else {
		model.AntiVirus = basetypes.NewObjectNull(models.UpdateScheduleUpdateScheduleAntiVirus{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	// Logic for non-pointer / value-type nested objects
	if !reflect.ValueOf(sdk.Threats).IsZero() {
		tflog.Debug(ctx, "Packing nested object for field Threats")
		packed, d := packUpdateScheduleUpdateScheduleThreatsFromSdk(ctx, sdk.Threats)
		diags.Append(d...)
		model.Threats = packed
	} else {
		model.Threats = basetypes.NewObjectNull(models.UpdateScheduleUpdateScheduleThreats{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	// Logic for non-pointer / value-type nested objects
	if !reflect.ValueOf(sdk.Wildfire).IsZero() {
		tflog.Debug(ctx, "Packing nested object for field Wildfire")
		packed, d := packUpdateScheduleUpdateScheduleWildfireFromSdk(ctx, sdk.Wildfire)
		diags.Append(d...)
		model.Wildfire = packed
	} else {
		model.Wildfire = basetypes.NewObjectNull(models.UpdateScheduleUpdateScheduleWildfire{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.UpdateScheduleUpdateSchedule{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.UpdateScheduleUpdateSchedule", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for UpdateScheduleUpdateSchedule ---
func unpackUpdateScheduleUpdateScheduleListToSdk(ctx context.Context, list types.List) ([]device_settings.UpdateScheduleUpdateSchedule, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.UpdateScheduleUpdateSchedule")
	diags := diag.Diagnostics{}
	var data []models.UpdateScheduleUpdateSchedule
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]device_settings.UpdateScheduleUpdateSchedule, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.UpdateScheduleUpdateSchedule{}.AttrTypes(), &item)
		unpacked, d := unpackUpdateScheduleUpdateScheduleToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.UpdateScheduleUpdateSchedule", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for UpdateScheduleUpdateSchedule ---
func packUpdateScheduleUpdateScheduleListFromSdk(ctx context.Context, sdks []device_settings.UpdateScheduleUpdateSchedule) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.UpdateScheduleUpdateSchedule")
	diags := diag.Diagnostics{}
	var data []models.UpdateScheduleUpdateSchedule

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.UpdateScheduleUpdateSchedule
		obj, d := packUpdateScheduleUpdateScheduleFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.UpdateScheduleUpdateSchedule{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.UpdateScheduleUpdateSchedule", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.UpdateScheduleUpdateSchedule{}.AttrType(), data)
}

// --- Unpacker for UpdateScheduleUpdateScheduleAntiVirus ---
func unpackUpdateScheduleUpdateScheduleAntiVirusToSdk(ctx context.Context, obj types.Object) (*device_settings.UpdateScheduleUpdateScheduleAntiVirus, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.UpdateScheduleUpdateScheduleAntiVirus", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.UpdateScheduleUpdateScheduleAntiVirus
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk device_settings.UpdateScheduleUpdateScheduleAntiVirus
	var d diag.Diagnostics
	// Handling Objects
	if !model.Recurring.IsNull() && !model.Recurring.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Recurring")
		unpacked, d := unpackUpdateScheduleUpdateScheduleAntiVirusRecurringToSdk(ctx, model.Recurring)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Recurring"})
		}
		if unpacked != nil {
			sdk.Recurring = *unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.UpdateScheduleUpdateScheduleAntiVirus", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for UpdateScheduleUpdateScheduleAntiVirus ---
func packUpdateScheduleUpdateScheduleAntiVirusFromSdk(ctx context.Context, sdk device_settings.UpdateScheduleUpdateScheduleAntiVirus) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.UpdateScheduleUpdateScheduleAntiVirus", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.UpdateScheduleUpdateScheduleAntiVirus
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	// Logic for non-pointer / value-type nested objects
	if !reflect.ValueOf(sdk.Recurring).IsZero() {
		tflog.Debug(ctx, "Packing nested object for field Recurring")
		packed, d := packUpdateScheduleUpdateScheduleAntiVirusRecurringFromSdk(ctx, sdk.Recurring)
		diags.Append(d...)
		model.Recurring = packed
	} else {
		model.Recurring = basetypes.NewObjectNull(models.UpdateScheduleUpdateScheduleAntiVirusRecurring{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.UpdateScheduleUpdateScheduleAntiVirus{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.UpdateScheduleUpdateScheduleAntiVirus", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for UpdateScheduleUpdateScheduleAntiVirus ---
func unpackUpdateScheduleUpdateScheduleAntiVirusListToSdk(ctx context.Context, list types.List) ([]device_settings.UpdateScheduleUpdateScheduleAntiVirus, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.UpdateScheduleUpdateScheduleAntiVirus")
	diags := diag.Diagnostics{}
	var data []models.UpdateScheduleUpdateScheduleAntiVirus
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]device_settings.UpdateScheduleUpdateScheduleAntiVirus, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.UpdateScheduleUpdateScheduleAntiVirus{}.AttrTypes(), &item)
		unpacked, d := unpackUpdateScheduleUpdateScheduleAntiVirusToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.UpdateScheduleUpdateScheduleAntiVirus", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for UpdateScheduleUpdateScheduleAntiVirus ---
func packUpdateScheduleUpdateScheduleAntiVirusListFromSdk(ctx context.Context, sdks []device_settings.UpdateScheduleUpdateScheduleAntiVirus) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.UpdateScheduleUpdateScheduleAntiVirus")
	diags := diag.Diagnostics{}
	var data []models.UpdateScheduleUpdateScheduleAntiVirus

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.UpdateScheduleUpdateScheduleAntiVirus
		obj, d := packUpdateScheduleUpdateScheduleAntiVirusFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.UpdateScheduleUpdateScheduleAntiVirus{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.UpdateScheduleUpdateScheduleAntiVirus", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.UpdateScheduleUpdateScheduleAntiVirus{}.AttrType(), data)
}

// --- Unpacker for UpdateScheduleUpdateScheduleAntiVirusRecurring ---
func unpackUpdateScheduleUpdateScheduleAntiVirusRecurringToSdk(ctx context.Context, obj types.Object) (*device_settings.UpdateScheduleUpdateScheduleAntiVirusRecurring, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.UpdateScheduleUpdateScheduleAntiVirusRecurring", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.UpdateScheduleUpdateScheduleAntiVirusRecurring
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk device_settings.UpdateScheduleUpdateScheduleAntiVirusRecurring
	var d diag.Diagnostics
	// Handling Objects
	if !model.Daily.IsNull() && !model.Daily.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Daily")
		unpacked, d := unpackUpdateScheduleUpdateScheduleAntiVirusRecurringDailyToSdk(ctx, model.Daily)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Daily"})
		}
		if unpacked != nil {
			sdk.Daily = unpacked
		}
	}

	// Handling Objects
	if !model.Hourly.IsNull() && !model.Hourly.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Hourly")
		unpacked, d := unpackUpdateScheduleUpdateScheduleAntiVirusRecurringHourlyToSdk(ctx, model.Hourly)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Hourly"})
		}
		if unpacked != nil {
			sdk.Hourly = unpacked
		}
	}

	// Handling Typeless Objects
	if !model.None.IsNull() && !model.None.IsUnknown() {
		tflog.Debug(ctx, "Unpacking typeless object for field None")
		sdk.None = make(map[string]interface{})
	}

	// Handling Primitives
	if !model.SyncToPeer.IsNull() && !model.SyncToPeer.IsUnknown() {
		sdk.SyncToPeer = model.SyncToPeer.ValueBool()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "SyncToPeer", "value": sdk.SyncToPeer})
	}

	// Handling Primitives
	if !model.Threshold.IsNull() && !model.Threshold.IsUnknown() {
		val := int32(model.Threshold.ValueInt64())
		sdk.Threshold = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Threshold", "value": *sdk.Threshold})
	}

	// Handling Objects
	if !model.Weekly.IsNull() && !model.Weekly.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Weekly")
		unpacked, d := unpackUpdateScheduleUpdateScheduleAntiVirusRecurringWeeklyToSdk(ctx, model.Weekly)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Weekly"})
		}
		if unpacked != nil {
			sdk.Weekly = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.UpdateScheduleUpdateScheduleAntiVirusRecurring", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for UpdateScheduleUpdateScheduleAntiVirusRecurring ---
func packUpdateScheduleUpdateScheduleAntiVirusRecurringFromSdk(ctx context.Context, sdk device_settings.UpdateScheduleUpdateScheduleAntiVirusRecurring) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.UpdateScheduleUpdateScheduleAntiVirusRecurring", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.UpdateScheduleUpdateScheduleAntiVirusRecurring
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Daily != nil {
		tflog.Debug(ctx, "Packing nested object for field Daily")
		packed, d := packUpdateScheduleUpdateScheduleAntiVirusRecurringDailyFromSdk(ctx, *sdk.Daily)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Daily"})
		}
		model.Daily = packed
	} else {
		model.Daily = basetypes.NewObjectNull(models.UpdateScheduleUpdateScheduleAntiVirusRecurringDaily{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Hourly != nil {
		tflog.Debug(ctx, "Packing nested object for field Hourly")
		packed, d := packUpdateScheduleUpdateScheduleAntiVirusRecurringHourlyFromSdk(ctx, *sdk.Hourly)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Hourly"})
		}
		model.Hourly = packed
	} else {
		model.Hourly = basetypes.NewObjectNull(models.UpdateScheduleUpdateScheduleAntiVirusRecurringHourly{}.AttrTypes())
	}
	// Handling Objects
	// This is a marker object (e.g. CHAP: {}). We just need to create an empty, non-null object.
	if sdk.None != nil && !reflect.ValueOf(sdk.None).IsNil() {
		tflog.Debug(ctx, "Packing typeless object for field None")
		var d diag.Diagnostics
		// Create an empty object with no attributes, which signifies its presence.
		model.None, d = basetypes.NewObjectValue(map[string]attr.Type{}, map[string]attr.Value{})
		diags.Append(d...)
	} else {
		// Since this field is part of a oneOf, being nil means it's not selected.
		// We make the object null with an empty attribute map.
		model.None = basetypes.NewObjectNull(map[string]attr.Type{})
	}
	// Handling Primitives
	// Standard primitive packing
	model.SyncToPeer = basetypes.NewBoolValue(sdk.SyncToPeer)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "SyncToPeer", "value": sdk.SyncToPeer})
	// Handling Primitives
	// Standard primitive packing
	if sdk.Threshold != nil {
		model.Threshold = basetypes.NewInt64Value(int64(*sdk.Threshold))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Threshold", "value": *sdk.Threshold})
	} else {
		model.Threshold = basetypes.NewInt64Null()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Weekly != nil {
		tflog.Debug(ctx, "Packing nested object for field Weekly")
		packed, d := packUpdateScheduleUpdateScheduleAntiVirusRecurringWeeklyFromSdk(ctx, *sdk.Weekly)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Weekly"})
		}
		model.Weekly = packed
	} else {
		model.Weekly = basetypes.NewObjectNull(models.UpdateScheduleUpdateScheduleAntiVirusRecurringWeekly{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.UpdateScheduleUpdateScheduleAntiVirusRecurring{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.UpdateScheduleUpdateScheduleAntiVirusRecurring", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for UpdateScheduleUpdateScheduleAntiVirusRecurring ---
func unpackUpdateScheduleUpdateScheduleAntiVirusRecurringListToSdk(ctx context.Context, list types.List) ([]device_settings.UpdateScheduleUpdateScheduleAntiVirusRecurring, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.UpdateScheduleUpdateScheduleAntiVirusRecurring")
	diags := diag.Diagnostics{}
	var data []models.UpdateScheduleUpdateScheduleAntiVirusRecurring
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]device_settings.UpdateScheduleUpdateScheduleAntiVirusRecurring, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.UpdateScheduleUpdateScheduleAntiVirusRecurring{}.AttrTypes(), &item)
		unpacked, d := unpackUpdateScheduleUpdateScheduleAntiVirusRecurringToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.UpdateScheduleUpdateScheduleAntiVirusRecurring", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for UpdateScheduleUpdateScheduleAntiVirusRecurring ---
func packUpdateScheduleUpdateScheduleAntiVirusRecurringListFromSdk(ctx context.Context, sdks []device_settings.UpdateScheduleUpdateScheduleAntiVirusRecurring) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.UpdateScheduleUpdateScheduleAntiVirusRecurring")
	diags := diag.Diagnostics{}
	var data []models.UpdateScheduleUpdateScheduleAntiVirusRecurring

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.UpdateScheduleUpdateScheduleAntiVirusRecurring
		obj, d := packUpdateScheduleUpdateScheduleAntiVirusRecurringFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.UpdateScheduleUpdateScheduleAntiVirusRecurring{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.UpdateScheduleUpdateScheduleAntiVirusRecurring", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.UpdateScheduleUpdateScheduleAntiVirusRecurring{}.AttrType(), data)
}

// --- Unpacker for UpdateScheduleUpdateScheduleAntiVirusRecurringDaily ---
func unpackUpdateScheduleUpdateScheduleAntiVirusRecurringDailyToSdk(ctx context.Context, obj types.Object) (*device_settings.UpdateScheduleUpdateScheduleAntiVirusRecurringDaily, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.UpdateScheduleUpdateScheduleAntiVirusRecurringDaily", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.UpdateScheduleUpdateScheduleAntiVirusRecurringDaily
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk device_settings.UpdateScheduleUpdateScheduleAntiVirusRecurringDaily
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Action.IsNull() && !model.Action.IsUnknown() {
		sdk.Action = model.Action.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Action", "value": *sdk.Action})
	}

	// Handling Primitives
	if !model.At.IsNull() && !model.At.IsUnknown() {
		sdk.At = model.At.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "At", "value": sdk.At})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.UpdateScheduleUpdateScheduleAntiVirusRecurringDaily", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for UpdateScheduleUpdateScheduleAntiVirusRecurringDaily ---
func packUpdateScheduleUpdateScheduleAntiVirusRecurringDailyFromSdk(ctx context.Context, sdk device_settings.UpdateScheduleUpdateScheduleAntiVirusRecurringDaily) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.UpdateScheduleUpdateScheduleAntiVirusRecurringDaily", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.UpdateScheduleUpdateScheduleAntiVirusRecurringDaily
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Action != nil {
		model.Action = basetypes.NewStringValue(*sdk.Action)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Action", "value": *sdk.Action})
	} else {
		model.Action = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	model.At = basetypes.NewStringValue(sdk.At)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "At", "value": sdk.At})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.UpdateScheduleUpdateScheduleAntiVirusRecurringDaily{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.UpdateScheduleUpdateScheduleAntiVirusRecurringDaily", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for UpdateScheduleUpdateScheduleAntiVirusRecurringDaily ---
func unpackUpdateScheduleUpdateScheduleAntiVirusRecurringDailyListToSdk(ctx context.Context, list types.List) ([]device_settings.UpdateScheduleUpdateScheduleAntiVirusRecurringDaily, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.UpdateScheduleUpdateScheduleAntiVirusRecurringDaily")
	diags := diag.Diagnostics{}
	var data []models.UpdateScheduleUpdateScheduleAntiVirusRecurringDaily
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]device_settings.UpdateScheduleUpdateScheduleAntiVirusRecurringDaily, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.UpdateScheduleUpdateScheduleAntiVirusRecurringDaily{}.AttrTypes(), &item)
		unpacked, d := unpackUpdateScheduleUpdateScheduleAntiVirusRecurringDailyToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.UpdateScheduleUpdateScheduleAntiVirusRecurringDaily", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for UpdateScheduleUpdateScheduleAntiVirusRecurringDaily ---
func packUpdateScheduleUpdateScheduleAntiVirusRecurringDailyListFromSdk(ctx context.Context, sdks []device_settings.UpdateScheduleUpdateScheduleAntiVirusRecurringDaily) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.UpdateScheduleUpdateScheduleAntiVirusRecurringDaily")
	diags := diag.Diagnostics{}
	var data []models.UpdateScheduleUpdateScheduleAntiVirusRecurringDaily

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.UpdateScheduleUpdateScheduleAntiVirusRecurringDaily
		obj, d := packUpdateScheduleUpdateScheduleAntiVirusRecurringDailyFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.UpdateScheduleUpdateScheduleAntiVirusRecurringDaily{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.UpdateScheduleUpdateScheduleAntiVirusRecurringDaily", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.UpdateScheduleUpdateScheduleAntiVirusRecurringDaily{}.AttrType(), data)
}

// --- Unpacker for UpdateScheduleUpdateScheduleAntiVirusRecurringHourly ---
func unpackUpdateScheduleUpdateScheduleAntiVirusRecurringHourlyToSdk(ctx context.Context, obj types.Object) (*device_settings.UpdateScheduleUpdateScheduleAntiVirusRecurringHourly, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.UpdateScheduleUpdateScheduleAntiVirusRecurringHourly", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.UpdateScheduleUpdateScheduleAntiVirusRecurringHourly
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk device_settings.UpdateScheduleUpdateScheduleAntiVirusRecurringHourly
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Action.IsNull() && !model.Action.IsUnknown() {
		sdk.Action = model.Action.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Action", "value": *sdk.Action})
	}

	// Handling Primitives
	if !model.At.IsNull() && !model.At.IsUnknown() {
		sdk.At = int32(model.At.ValueInt64())
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "At", "value": sdk.At})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.UpdateScheduleUpdateScheduleAntiVirusRecurringHourly", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for UpdateScheduleUpdateScheduleAntiVirusRecurringHourly ---
func packUpdateScheduleUpdateScheduleAntiVirusRecurringHourlyFromSdk(ctx context.Context, sdk device_settings.UpdateScheduleUpdateScheduleAntiVirusRecurringHourly) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.UpdateScheduleUpdateScheduleAntiVirusRecurringHourly", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.UpdateScheduleUpdateScheduleAntiVirusRecurringHourly
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Action != nil {
		model.Action = basetypes.NewStringValue(*sdk.Action)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Action", "value": *sdk.Action})
	} else {
		model.Action = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	model.At = basetypes.NewInt64Value(int64(sdk.At))
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "At", "value": sdk.At})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.UpdateScheduleUpdateScheduleAntiVirusRecurringHourly{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.UpdateScheduleUpdateScheduleAntiVirusRecurringHourly", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for UpdateScheduleUpdateScheduleAntiVirusRecurringHourly ---
func unpackUpdateScheduleUpdateScheduleAntiVirusRecurringHourlyListToSdk(ctx context.Context, list types.List) ([]device_settings.UpdateScheduleUpdateScheduleAntiVirusRecurringHourly, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.UpdateScheduleUpdateScheduleAntiVirusRecurringHourly")
	diags := diag.Diagnostics{}
	var data []models.UpdateScheduleUpdateScheduleAntiVirusRecurringHourly
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]device_settings.UpdateScheduleUpdateScheduleAntiVirusRecurringHourly, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.UpdateScheduleUpdateScheduleAntiVirusRecurringHourly{}.AttrTypes(), &item)
		unpacked, d := unpackUpdateScheduleUpdateScheduleAntiVirusRecurringHourlyToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.UpdateScheduleUpdateScheduleAntiVirusRecurringHourly", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for UpdateScheduleUpdateScheduleAntiVirusRecurringHourly ---
func packUpdateScheduleUpdateScheduleAntiVirusRecurringHourlyListFromSdk(ctx context.Context, sdks []device_settings.UpdateScheduleUpdateScheduleAntiVirusRecurringHourly) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.UpdateScheduleUpdateScheduleAntiVirusRecurringHourly")
	diags := diag.Diagnostics{}
	var data []models.UpdateScheduleUpdateScheduleAntiVirusRecurringHourly

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.UpdateScheduleUpdateScheduleAntiVirusRecurringHourly
		obj, d := packUpdateScheduleUpdateScheduleAntiVirusRecurringHourlyFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.UpdateScheduleUpdateScheduleAntiVirusRecurringHourly{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.UpdateScheduleUpdateScheduleAntiVirusRecurringHourly", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.UpdateScheduleUpdateScheduleAntiVirusRecurringHourly{}.AttrType(), data)
}

// --- Unpacker for UpdateScheduleUpdateScheduleAntiVirusRecurringWeekly ---
func unpackUpdateScheduleUpdateScheduleAntiVirusRecurringWeeklyToSdk(ctx context.Context, obj types.Object) (*device_settings.UpdateScheduleUpdateScheduleAntiVirusRecurringWeekly, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.UpdateScheduleUpdateScheduleAntiVirusRecurringWeekly", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.UpdateScheduleUpdateScheduleAntiVirusRecurringWeekly
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk device_settings.UpdateScheduleUpdateScheduleAntiVirusRecurringWeekly
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Action.IsNull() && !model.Action.IsUnknown() {
		sdk.Action = model.Action.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Action", "value": *sdk.Action})
	}

	// Handling Primitives
	if !model.At.IsNull() && !model.At.IsUnknown() {
		sdk.At = model.At.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "At", "value": *sdk.At})
	}

	// Handling Primitives
	if !model.DayOfWeek.IsNull() && !model.DayOfWeek.IsUnknown() {
		sdk.DayOfWeek = model.DayOfWeek.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DayOfWeek", "value": *sdk.DayOfWeek})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.UpdateScheduleUpdateScheduleAntiVirusRecurringWeekly", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for UpdateScheduleUpdateScheduleAntiVirusRecurringWeekly ---
func packUpdateScheduleUpdateScheduleAntiVirusRecurringWeeklyFromSdk(ctx context.Context, sdk device_settings.UpdateScheduleUpdateScheduleAntiVirusRecurringWeekly) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.UpdateScheduleUpdateScheduleAntiVirusRecurringWeekly", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.UpdateScheduleUpdateScheduleAntiVirusRecurringWeekly
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Action != nil {
		model.Action = basetypes.NewStringValue(*sdk.Action)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Action", "value": *sdk.Action})
	} else {
		model.Action = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.At != nil {
		model.At = basetypes.NewStringValue(*sdk.At)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "At", "value": *sdk.At})
	} else {
		model.At = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.DayOfWeek != nil {
		model.DayOfWeek = basetypes.NewStringValue(*sdk.DayOfWeek)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DayOfWeek", "value": *sdk.DayOfWeek})
	} else {
		model.DayOfWeek = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.UpdateScheduleUpdateScheduleAntiVirusRecurringWeekly{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.UpdateScheduleUpdateScheduleAntiVirusRecurringWeekly", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for UpdateScheduleUpdateScheduleAntiVirusRecurringWeekly ---
func unpackUpdateScheduleUpdateScheduleAntiVirusRecurringWeeklyListToSdk(ctx context.Context, list types.List) ([]device_settings.UpdateScheduleUpdateScheduleAntiVirusRecurringWeekly, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.UpdateScheduleUpdateScheduleAntiVirusRecurringWeekly")
	diags := diag.Diagnostics{}
	var data []models.UpdateScheduleUpdateScheduleAntiVirusRecurringWeekly
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]device_settings.UpdateScheduleUpdateScheduleAntiVirusRecurringWeekly, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.UpdateScheduleUpdateScheduleAntiVirusRecurringWeekly{}.AttrTypes(), &item)
		unpacked, d := unpackUpdateScheduleUpdateScheduleAntiVirusRecurringWeeklyToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.UpdateScheduleUpdateScheduleAntiVirusRecurringWeekly", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for UpdateScheduleUpdateScheduleAntiVirusRecurringWeekly ---
func packUpdateScheduleUpdateScheduleAntiVirusRecurringWeeklyListFromSdk(ctx context.Context, sdks []device_settings.UpdateScheduleUpdateScheduleAntiVirusRecurringWeekly) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.UpdateScheduleUpdateScheduleAntiVirusRecurringWeekly")
	diags := diag.Diagnostics{}
	var data []models.UpdateScheduleUpdateScheduleAntiVirusRecurringWeekly

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.UpdateScheduleUpdateScheduleAntiVirusRecurringWeekly
		obj, d := packUpdateScheduleUpdateScheduleAntiVirusRecurringWeeklyFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.UpdateScheduleUpdateScheduleAntiVirusRecurringWeekly{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.UpdateScheduleUpdateScheduleAntiVirusRecurringWeekly", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.UpdateScheduleUpdateScheduleAntiVirusRecurringWeekly{}.AttrType(), data)
}

// --- Unpacker for UpdateScheduleUpdateScheduleThreats ---
func unpackUpdateScheduleUpdateScheduleThreatsToSdk(ctx context.Context, obj types.Object) (*device_settings.UpdateScheduleUpdateScheduleThreats, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.UpdateScheduleUpdateScheduleThreats", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.UpdateScheduleUpdateScheduleThreats
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk device_settings.UpdateScheduleUpdateScheduleThreats
	var d diag.Diagnostics
	// Handling Objects
	if !model.Recurring.IsNull() && !model.Recurring.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Recurring")
		unpacked, d := unpackUpdateScheduleUpdateScheduleThreatsRecurringToSdk(ctx, model.Recurring)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Recurring"})
		}
		if unpacked != nil {
			sdk.Recurring = *unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.UpdateScheduleUpdateScheduleThreats", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for UpdateScheduleUpdateScheduleThreats ---
func packUpdateScheduleUpdateScheduleThreatsFromSdk(ctx context.Context, sdk device_settings.UpdateScheduleUpdateScheduleThreats) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.UpdateScheduleUpdateScheduleThreats", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.UpdateScheduleUpdateScheduleThreats
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	// Logic for non-pointer / value-type nested objects
	if !reflect.ValueOf(sdk.Recurring).IsZero() {
		tflog.Debug(ctx, "Packing nested object for field Recurring")
		packed, d := packUpdateScheduleUpdateScheduleThreatsRecurringFromSdk(ctx, sdk.Recurring)
		diags.Append(d...)
		model.Recurring = packed
	} else {
		model.Recurring = basetypes.NewObjectNull(models.UpdateScheduleUpdateScheduleThreatsRecurring{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.UpdateScheduleUpdateScheduleThreats{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.UpdateScheduleUpdateScheduleThreats", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for UpdateScheduleUpdateScheduleThreats ---
func unpackUpdateScheduleUpdateScheduleThreatsListToSdk(ctx context.Context, list types.List) ([]device_settings.UpdateScheduleUpdateScheduleThreats, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.UpdateScheduleUpdateScheduleThreats")
	diags := diag.Diagnostics{}
	var data []models.UpdateScheduleUpdateScheduleThreats
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]device_settings.UpdateScheduleUpdateScheduleThreats, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.UpdateScheduleUpdateScheduleThreats{}.AttrTypes(), &item)
		unpacked, d := unpackUpdateScheduleUpdateScheduleThreatsToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.UpdateScheduleUpdateScheduleThreats", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for UpdateScheduleUpdateScheduleThreats ---
func packUpdateScheduleUpdateScheduleThreatsListFromSdk(ctx context.Context, sdks []device_settings.UpdateScheduleUpdateScheduleThreats) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.UpdateScheduleUpdateScheduleThreats")
	diags := diag.Diagnostics{}
	var data []models.UpdateScheduleUpdateScheduleThreats

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.UpdateScheduleUpdateScheduleThreats
		obj, d := packUpdateScheduleUpdateScheduleThreatsFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.UpdateScheduleUpdateScheduleThreats{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.UpdateScheduleUpdateScheduleThreats", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.UpdateScheduleUpdateScheduleThreats{}.AttrType(), data)
}

// --- Unpacker for UpdateScheduleUpdateScheduleThreatsRecurring ---
func unpackUpdateScheduleUpdateScheduleThreatsRecurringToSdk(ctx context.Context, obj types.Object) (*device_settings.UpdateScheduleUpdateScheduleThreatsRecurring, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.UpdateScheduleUpdateScheduleThreatsRecurring", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.UpdateScheduleUpdateScheduleThreatsRecurring
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk device_settings.UpdateScheduleUpdateScheduleThreatsRecurring
	var d diag.Diagnostics
	// Handling Objects
	if !model.Daily.IsNull() && !model.Daily.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Daily")
		unpacked, d := unpackUpdateScheduleUpdateScheduleThreatsRecurringDailyToSdk(ctx, model.Daily)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Daily"})
		}
		if unpacked != nil {
			sdk.Daily = unpacked
		}
	}

	// Handling Objects
	if !model.Every30Mins.IsNull() && !model.Every30Mins.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Every30Mins")
		unpacked, d := unpackUpdateScheduleUpdateScheduleThreatsRecurringEvery30MinsToSdk(ctx, model.Every30Mins)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Every30Mins"})
		}
		if unpacked != nil {
			sdk.Every30Mins = unpacked
		}
	}

	// Handling Objects
	if !model.Hourly.IsNull() && !model.Hourly.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Hourly")
		unpacked, d := unpackUpdateScheduleUpdateScheduleThreatsRecurringHourlyToSdk(ctx, model.Hourly)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Hourly"})
		}
		if unpacked != nil {
			sdk.Hourly = unpacked
		}
	}

	// Handling Primitives
	if !model.NewAppThreshold.IsNull() && !model.NewAppThreshold.IsUnknown() {
		val := int32(model.NewAppThreshold.ValueInt64())
		sdk.NewAppThreshold = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "NewAppThreshold", "value": *sdk.NewAppThreshold})
	}

	// Handling Typeless Objects
	if !model.None.IsNull() && !model.None.IsUnknown() {
		tflog.Debug(ctx, "Unpacking typeless object for field None")
		sdk.None = make(map[string]interface{})
	}

	// Handling Primitives
	if !model.SyncToPeer.IsNull() && !model.SyncToPeer.IsUnknown() {
		sdk.SyncToPeer = model.SyncToPeer.ValueBool()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "SyncToPeer", "value": sdk.SyncToPeer})
	}

	// Handling Primitives
	if !model.Threshold.IsNull() && !model.Threshold.IsUnknown() {
		val := int32(model.Threshold.ValueInt64())
		sdk.Threshold = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Threshold", "value": *sdk.Threshold})
	}

	// Handling Objects
	if !model.Weekly.IsNull() && !model.Weekly.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Weekly")
		unpacked, d := unpackUpdateScheduleUpdateScheduleThreatsRecurringWeeklyToSdk(ctx, model.Weekly)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Weekly"})
		}
		if unpacked != nil {
			sdk.Weekly = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.UpdateScheduleUpdateScheduleThreatsRecurring", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for UpdateScheduleUpdateScheduleThreatsRecurring ---
func packUpdateScheduleUpdateScheduleThreatsRecurringFromSdk(ctx context.Context, sdk device_settings.UpdateScheduleUpdateScheduleThreatsRecurring) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.UpdateScheduleUpdateScheduleThreatsRecurring", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.UpdateScheduleUpdateScheduleThreatsRecurring
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Daily != nil {
		tflog.Debug(ctx, "Packing nested object for field Daily")
		packed, d := packUpdateScheduleUpdateScheduleThreatsRecurringDailyFromSdk(ctx, *sdk.Daily)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Daily"})
		}
		model.Daily = packed
	} else {
		model.Daily = basetypes.NewObjectNull(models.UpdateScheduleUpdateScheduleThreatsRecurringDaily{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Every30Mins != nil {
		tflog.Debug(ctx, "Packing nested object for field Every30Mins")
		packed, d := packUpdateScheduleUpdateScheduleThreatsRecurringEvery30MinsFromSdk(ctx, *sdk.Every30Mins)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Every30Mins"})
		}
		model.Every30Mins = packed
	} else {
		model.Every30Mins = basetypes.NewObjectNull(models.UpdateScheduleUpdateScheduleThreatsRecurringEvery30Mins{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Hourly != nil {
		tflog.Debug(ctx, "Packing nested object for field Hourly")
		packed, d := packUpdateScheduleUpdateScheduleThreatsRecurringHourlyFromSdk(ctx, *sdk.Hourly)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Hourly"})
		}
		model.Hourly = packed
	} else {
		model.Hourly = basetypes.NewObjectNull(models.UpdateScheduleUpdateScheduleThreatsRecurringHourly{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.NewAppThreshold != nil {
		model.NewAppThreshold = basetypes.NewInt64Value(int64(*sdk.NewAppThreshold))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "NewAppThreshold", "value": *sdk.NewAppThreshold})
	} else {
		model.NewAppThreshold = basetypes.NewInt64Null()
	}
	// Handling Objects
	// This is a marker object (e.g. CHAP: {}). We just need to create an empty, non-null object.
	if sdk.None != nil && !reflect.ValueOf(sdk.None).IsNil() {
		tflog.Debug(ctx, "Packing typeless object for field None")
		var d diag.Diagnostics
		// Create an empty object with no attributes, which signifies its presence.
		model.None, d = basetypes.NewObjectValue(map[string]attr.Type{}, map[string]attr.Value{})
		diags.Append(d...)
	} else {
		// Since this field is part of a oneOf, being nil means it's not selected.
		// We make the object null with an empty attribute map.
		model.None = basetypes.NewObjectNull(map[string]attr.Type{})
	}
	// Handling Primitives
	// Standard primitive packing
	model.SyncToPeer = basetypes.NewBoolValue(sdk.SyncToPeer)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "SyncToPeer", "value": sdk.SyncToPeer})
	// Handling Primitives
	// Standard primitive packing
	if sdk.Threshold != nil {
		model.Threshold = basetypes.NewInt64Value(int64(*sdk.Threshold))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Threshold", "value": *sdk.Threshold})
	} else {
		model.Threshold = basetypes.NewInt64Null()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Weekly != nil {
		tflog.Debug(ctx, "Packing nested object for field Weekly")
		packed, d := packUpdateScheduleUpdateScheduleThreatsRecurringWeeklyFromSdk(ctx, *sdk.Weekly)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Weekly"})
		}
		model.Weekly = packed
	} else {
		model.Weekly = basetypes.NewObjectNull(models.UpdateScheduleUpdateScheduleThreatsRecurringWeekly{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.UpdateScheduleUpdateScheduleThreatsRecurring{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.UpdateScheduleUpdateScheduleThreatsRecurring", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for UpdateScheduleUpdateScheduleThreatsRecurring ---
func unpackUpdateScheduleUpdateScheduleThreatsRecurringListToSdk(ctx context.Context, list types.List) ([]device_settings.UpdateScheduleUpdateScheduleThreatsRecurring, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.UpdateScheduleUpdateScheduleThreatsRecurring")
	diags := diag.Diagnostics{}
	var data []models.UpdateScheduleUpdateScheduleThreatsRecurring
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]device_settings.UpdateScheduleUpdateScheduleThreatsRecurring, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.UpdateScheduleUpdateScheduleThreatsRecurring{}.AttrTypes(), &item)
		unpacked, d := unpackUpdateScheduleUpdateScheduleThreatsRecurringToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.UpdateScheduleUpdateScheduleThreatsRecurring", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for UpdateScheduleUpdateScheduleThreatsRecurring ---
func packUpdateScheduleUpdateScheduleThreatsRecurringListFromSdk(ctx context.Context, sdks []device_settings.UpdateScheduleUpdateScheduleThreatsRecurring) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.UpdateScheduleUpdateScheduleThreatsRecurring")
	diags := diag.Diagnostics{}
	var data []models.UpdateScheduleUpdateScheduleThreatsRecurring

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.UpdateScheduleUpdateScheduleThreatsRecurring
		obj, d := packUpdateScheduleUpdateScheduleThreatsRecurringFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.UpdateScheduleUpdateScheduleThreatsRecurring{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.UpdateScheduleUpdateScheduleThreatsRecurring", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.UpdateScheduleUpdateScheduleThreatsRecurring{}.AttrType(), data)
}

// --- Unpacker for UpdateScheduleUpdateScheduleThreatsRecurringDaily ---
func unpackUpdateScheduleUpdateScheduleThreatsRecurringDailyToSdk(ctx context.Context, obj types.Object) (*device_settings.UpdateScheduleUpdateScheduleThreatsRecurringDaily, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.UpdateScheduleUpdateScheduleThreatsRecurringDaily", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.UpdateScheduleUpdateScheduleThreatsRecurringDaily
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk device_settings.UpdateScheduleUpdateScheduleThreatsRecurringDaily
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Action.IsNull() && !model.Action.IsUnknown() {
		sdk.Action = model.Action.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Action", "value": *sdk.Action})
	}

	// Handling Primitives
	if !model.At.IsNull() && !model.At.IsUnknown() {
		sdk.At = model.At.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "At", "value": sdk.At})
	}

	// Handling Primitives
	if !model.DisableNewContent.IsNull() && !model.DisableNewContent.IsUnknown() {
		sdk.DisableNewContent = model.DisableNewContent.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DisableNewContent", "value": *sdk.DisableNewContent})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.UpdateScheduleUpdateScheduleThreatsRecurringDaily", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for UpdateScheduleUpdateScheduleThreatsRecurringDaily ---
func packUpdateScheduleUpdateScheduleThreatsRecurringDailyFromSdk(ctx context.Context, sdk device_settings.UpdateScheduleUpdateScheduleThreatsRecurringDaily) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.UpdateScheduleUpdateScheduleThreatsRecurringDaily", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.UpdateScheduleUpdateScheduleThreatsRecurringDaily
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Action != nil {
		model.Action = basetypes.NewStringValue(*sdk.Action)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Action", "value": *sdk.Action})
	} else {
		model.Action = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	model.At = basetypes.NewStringValue(sdk.At)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "At", "value": sdk.At})
	// Handling Primitives
	// Standard primitive packing
	if sdk.DisableNewContent != nil {
		model.DisableNewContent = basetypes.NewBoolValue(*sdk.DisableNewContent)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DisableNewContent", "value": *sdk.DisableNewContent})
	} else {
		model.DisableNewContent = basetypes.NewBoolNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.UpdateScheduleUpdateScheduleThreatsRecurringDaily{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.UpdateScheduleUpdateScheduleThreatsRecurringDaily", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for UpdateScheduleUpdateScheduleThreatsRecurringDaily ---
func unpackUpdateScheduleUpdateScheduleThreatsRecurringDailyListToSdk(ctx context.Context, list types.List) ([]device_settings.UpdateScheduleUpdateScheduleThreatsRecurringDaily, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.UpdateScheduleUpdateScheduleThreatsRecurringDaily")
	diags := diag.Diagnostics{}
	var data []models.UpdateScheduleUpdateScheduleThreatsRecurringDaily
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]device_settings.UpdateScheduleUpdateScheduleThreatsRecurringDaily, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.UpdateScheduleUpdateScheduleThreatsRecurringDaily{}.AttrTypes(), &item)
		unpacked, d := unpackUpdateScheduleUpdateScheduleThreatsRecurringDailyToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.UpdateScheduleUpdateScheduleThreatsRecurringDaily", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for UpdateScheduleUpdateScheduleThreatsRecurringDaily ---
func packUpdateScheduleUpdateScheduleThreatsRecurringDailyListFromSdk(ctx context.Context, sdks []device_settings.UpdateScheduleUpdateScheduleThreatsRecurringDaily) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.UpdateScheduleUpdateScheduleThreatsRecurringDaily")
	diags := diag.Diagnostics{}
	var data []models.UpdateScheduleUpdateScheduleThreatsRecurringDaily

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.UpdateScheduleUpdateScheduleThreatsRecurringDaily
		obj, d := packUpdateScheduleUpdateScheduleThreatsRecurringDailyFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.UpdateScheduleUpdateScheduleThreatsRecurringDaily{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.UpdateScheduleUpdateScheduleThreatsRecurringDaily", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.UpdateScheduleUpdateScheduleThreatsRecurringDaily{}.AttrType(), data)
}

// --- Unpacker for UpdateScheduleUpdateScheduleThreatsRecurringEvery30Mins ---
func unpackUpdateScheduleUpdateScheduleThreatsRecurringEvery30MinsToSdk(ctx context.Context, obj types.Object) (*device_settings.UpdateScheduleUpdateScheduleThreatsRecurringEvery30Mins, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.UpdateScheduleUpdateScheduleThreatsRecurringEvery30Mins", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.UpdateScheduleUpdateScheduleThreatsRecurringEvery30Mins
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk device_settings.UpdateScheduleUpdateScheduleThreatsRecurringEvery30Mins
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Action.IsNull() && !model.Action.IsUnknown() {
		sdk.Action = model.Action.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Action", "value": *sdk.Action})
	}

	// Handling Primitives
	if !model.At.IsNull() && !model.At.IsUnknown() {
		val := int32(model.At.ValueInt64())
		sdk.At = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "At", "value": *sdk.At})
	}

	// Handling Primitives
	if !model.DisableNewContent.IsNull() && !model.DisableNewContent.IsUnknown() {
		sdk.DisableNewContent = model.DisableNewContent.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DisableNewContent", "value": *sdk.DisableNewContent})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.UpdateScheduleUpdateScheduleThreatsRecurringEvery30Mins", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for UpdateScheduleUpdateScheduleThreatsRecurringEvery30Mins ---
func packUpdateScheduleUpdateScheduleThreatsRecurringEvery30MinsFromSdk(ctx context.Context, sdk device_settings.UpdateScheduleUpdateScheduleThreatsRecurringEvery30Mins) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.UpdateScheduleUpdateScheduleThreatsRecurringEvery30Mins", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.UpdateScheduleUpdateScheduleThreatsRecurringEvery30Mins
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Action != nil {
		model.Action = basetypes.NewStringValue(*sdk.Action)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Action", "value": *sdk.Action})
	} else {
		model.Action = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.At != nil {
		model.At = basetypes.NewInt64Value(int64(*sdk.At))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "At", "value": *sdk.At})
	} else {
		model.At = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.DisableNewContent != nil {
		model.DisableNewContent = basetypes.NewBoolValue(*sdk.DisableNewContent)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DisableNewContent", "value": *sdk.DisableNewContent})
	} else {
		model.DisableNewContent = basetypes.NewBoolNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.UpdateScheduleUpdateScheduleThreatsRecurringEvery30Mins{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.UpdateScheduleUpdateScheduleThreatsRecurringEvery30Mins", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for UpdateScheduleUpdateScheduleThreatsRecurringEvery30Mins ---
func unpackUpdateScheduleUpdateScheduleThreatsRecurringEvery30MinsListToSdk(ctx context.Context, list types.List) ([]device_settings.UpdateScheduleUpdateScheduleThreatsRecurringEvery30Mins, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.UpdateScheduleUpdateScheduleThreatsRecurringEvery30Mins")
	diags := diag.Diagnostics{}
	var data []models.UpdateScheduleUpdateScheduleThreatsRecurringEvery30Mins
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]device_settings.UpdateScheduleUpdateScheduleThreatsRecurringEvery30Mins, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.UpdateScheduleUpdateScheduleThreatsRecurringEvery30Mins{}.AttrTypes(), &item)
		unpacked, d := unpackUpdateScheduleUpdateScheduleThreatsRecurringEvery30MinsToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.UpdateScheduleUpdateScheduleThreatsRecurringEvery30Mins", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for UpdateScheduleUpdateScheduleThreatsRecurringEvery30Mins ---
func packUpdateScheduleUpdateScheduleThreatsRecurringEvery30MinsListFromSdk(ctx context.Context, sdks []device_settings.UpdateScheduleUpdateScheduleThreatsRecurringEvery30Mins) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.UpdateScheduleUpdateScheduleThreatsRecurringEvery30Mins")
	diags := diag.Diagnostics{}
	var data []models.UpdateScheduleUpdateScheduleThreatsRecurringEvery30Mins

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.UpdateScheduleUpdateScheduleThreatsRecurringEvery30Mins
		obj, d := packUpdateScheduleUpdateScheduleThreatsRecurringEvery30MinsFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.UpdateScheduleUpdateScheduleThreatsRecurringEvery30Mins{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.UpdateScheduleUpdateScheduleThreatsRecurringEvery30Mins", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.UpdateScheduleUpdateScheduleThreatsRecurringEvery30Mins{}.AttrType(), data)
}

// --- Unpacker for UpdateScheduleUpdateScheduleThreatsRecurringHourly ---
func unpackUpdateScheduleUpdateScheduleThreatsRecurringHourlyToSdk(ctx context.Context, obj types.Object) (*device_settings.UpdateScheduleUpdateScheduleThreatsRecurringHourly, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.UpdateScheduleUpdateScheduleThreatsRecurringHourly", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.UpdateScheduleUpdateScheduleThreatsRecurringHourly
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk device_settings.UpdateScheduleUpdateScheduleThreatsRecurringHourly
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Action.IsNull() && !model.Action.IsUnknown() {
		sdk.Action = model.Action.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Action", "value": *sdk.Action})
	}

	// Handling Primitives
	if !model.At.IsNull() && !model.At.IsUnknown() {
		sdk.At = float32(model.At.ValueFloat64())
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "At", "value": sdk.At})
	}

	// Handling Primitives
	if !model.DisableNewContent.IsNull() && !model.DisableNewContent.IsUnknown() {
		sdk.DisableNewContent = model.DisableNewContent.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DisableNewContent", "value": *sdk.DisableNewContent})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.UpdateScheduleUpdateScheduleThreatsRecurringHourly", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for UpdateScheduleUpdateScheduleThreatsRecurringHourly ---
func packUpdateScheduleUpdateScheduleThreatsRecurringHourlyFromSdk(ctx context.Context, sdk device_settings.UpdateScheduleUpdateScheduleThreatsRecurringHourly) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.UpdateScheduleUpdateScheduleThreatsRecurringHourly", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.UpdateScheduleUpdateScheduleThreatsRecurringHourly
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Action != nil {
		model.Action = basetypes.NewStringValue(*sdk.Action)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Action", "value": *sdk.Action})
	} else {
		model.Action = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	model.At = basetypes.NewFloat64Value(float64(sdk.At))
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "At", "value": sdk.At})
	// Handling Primitives
	// Standard primitive packing
	if sdk.DisableNewContent != nil {
		model.DisableNewContent = basetypes.NewBoolValue(*sdk.DisableNewContent)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DisableNewContent", "value": *sdk.DisableNewContent})
	} else {
		model.DisableNewContent = basetypes.NewBoolNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.UpdateScheduleUpdateScheduleThreatsRecurringHourly{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.UpdateScheduleUpdateScheduleThreatsRecurringHourly", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for UpdateScheduleUpdateScheduleThreatsRecurringHourly ---
func unpackUpdateScheduleUpdateScheduleThreatsRecurringHourlyListToSdk(ctx context.Context, list types.List) ([]device_settings.UpdateScheduleUpdateScheduleThreatsRecurringHourly, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.UpdateScheduleUpdateScheduleThreatsRecurringHourly")
	diags := diag.Diagnostics{}
	var data []models.UpdateScheduleUpdateScheduleThreatsRecurringHourly
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]device_settings.UpdateScheduleUpdateScheduleThreatsRecurringHourly, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.UpdateScheduleUpdateScheduleThreatsRecurringHourly{}.AttrTypes(), &item)
		unpacked, d := unpackUpdateScheduleUpdateScheduleThreatsRecurringHourlyToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.UpdateScheduleUpdateScheduleThreatsRecurringHourly", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for UpdateScheduleUpdateScheduleThreatsRecurringHourly ---
func packUpdateScheduleUpdateScheduleThreatsRecurringHourlyListFromSdk(ctx context.Context, sdks []device_settings.UpdateScheduleUpdateScheduleThreatsRecurringHourly) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.UpdateScheduleUpdateScheduleThreatsRecurringHourly")
	diags := diag.Diagnostics{}
	var data []models.UpdateScheduleUpdateScheduleThreatsRecurringHourly

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.UpdateScheduleUpdateScheduleThreatsRecurringHourly
		obj, d := packUpdateScheduleUpdateScheduleThreatsRecurringHourlyFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.UpdateScheduleUpdateScheduleThreatsRecurringHourly{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.UpdateScheduleUpdateScheduleThreatsRecurringHourly", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.UpdateScheduleUpdateScheduleThreatsRecurringHourly{}.AttrType(), data)
}

// --- Unpacker for UpdateScheduleUpdateScheduleThreatsRecurringWeekly ---
func unpackUpdateScheduleUpdateScheduleThreatsRecurringWeeklyToSdk(ctx context.Context, obj types.Object) (*device_settings.UpdateScheduleUpdateScheduleThreatsRecurringWeekly, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.UpdateScheduleUpdateScheduleThreatsRecurringWeekly", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.UpdateScheduleUpdateScheduleThreatsRecurringWeekly
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk device_settings.UpdateScheduleUpdateScheduleThreatsRecurringWeekly
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Action.IsNull() && !model.Action.IsUnknown() {
		sdk.Action = model.Action.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Action", "value": *sdk.Action})
	}

	// Handling Primitives
	if !model.At.IsNull() && !model.At.IsUnknown() {
		sdk.At = model.At.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "At", "value": sdk.At})
	}

	// Handling Primitives
	if !model.DayOfWeek.IsNull() && !model.DayOfWeek.IsUnknown() {
		sdk.DayOfWeek = model.DayOfWeek.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "DayOfWeek", "value": sdk.DayOfWeek})
	}

	// Handling Primitives
	if !model.DisableNewContent.IsNull() && !model.DisableNewContent.IsUnknown() {
		sdk.DisableNewContent = model.DisableNewContent.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DisableNewContent", "value": *sdk.DisableNewContent})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.UpdateScheduleUpdateScheduleThreatsRecurringWeekly", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for UpdateScheduleUpdateScheduleThreatsRecurringWeekly ---
func packUpdateScheduleUpdateScheduleThreatsRecurringWeeklyFromSdk(ctx context.Context, sdk device_settings.UpdateScheduleUpdateScheduleThreatsRecurringWeekly) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.UpdateScheduleUpdateScheduleThreatsRecurringWeekly", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.UpdateScheduleUpdateScheduleThreatsRecurringWeekly
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Action != nil {
		model.Action = basetypes.NewStringValue(*sdk.Action)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Action", "value": *sdk.Action})
	} else {
		model.Action = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	model.At = basetypes.NewStringValue(sdk.At)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "At", "value": sdk.At})
	// Handling Primitives
	// Standard primitive packing
	model.DayOfWeek = basetypes.NewStringValue(sdk.DayOfWeek)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "DayOfWeek", "value": sdk.DayOfWeek})
	// Handling Primitives
	// Standard primitive packing
	if sdk.DisableNewContent != nil {
		model.DisableNewContent = basetypes.NewBoolValue(*sdk.DisableNewContent)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DisableNewContent", "value": *sdk.DisableNewContent})
	} else {
		model.DisableNewContent = basetypes.NewBoolNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.UpdateScheduleUpdateScheduleThreatsRecurringWeekly{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.UpdateScheduleUpdateScheduleThreatsRecurringWeekly", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for UpdateScheduleUpdateScheduleThreatsRecurringWeekly ---
func unpackUpdateScheduleUpdateScheduleThreatsRecurringWeeklyListToSdk(ctx context.Context, list types.List) ([]device_settings.UpdateScheduleUpdateScheduleThreatsRecurringWeekly, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.UpdateScheduleUpdateScheduleThreatsRecurringWeekly")
	diags := diag.Diagnostics{}
	var data []models.UpdateScheduleUpdateScheduleThreatsRecurringWeekly
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]device_settings.UpdateScheduleUpdateScheduleThreatsRecurringWeekly, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.UpdateScheduleUpdateScheduleThreatsRecurringWeekly{}.AttrTypes(), &item)
		unpacked, d := unpackUpdateScheduleUpdateScheduleThreatsRecurringWeeklyToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.UpdateScheduleUpdateScheduleThreatsRecurringWeekly", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for UpdateScheduleUpdateScheduleThreatsRecurringWeekly ---
func packUpdateScheduleUpdateScheduleThreatsRecurringWeeklyListFromSdk(ctx context.Context, sdks []device_settings.UpdateScheduleUpdateScheduleThreatsRecurringWeekly) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.UpdateScheduleUpdateScheduleThreatsRecurringWeekly")
	diags := diag.Diagnostics{}
	var data []models.UpdateScheduleUpdateScheduleThreatsRecurringWeekly

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.UpdateScheduleUpdateScheduleThreatsRecurringWeekly
		obj, d := packUpdateScheduleUpdateScheduleThreatsRecurringWeeklyFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.UpdateScheduleUpdateScheduleThreatsRecurringWeekly{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.UpdateScheduleUpdateScheduleThreatsRecurringWeekly", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.UpdateScheduleUpdateScheduleThreatsRecurringWeekly{}.AttrType(), data)
}

// --- Unpacker for UpdateScheduleUpdateScheduleWildfire ---
func unpackUpdateScheduleUpdateScheduleWildfireToSdk(ctx context.Context, obj types.Object) (*device_settings.UpdateScheduleUpdateScheduleWildfire, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.UpdateScheduleUpdateScheduleWildfire", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.UpdateScheduleUpdateScheduleWildfire
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk device_settings.UpdateScheduleUpdateScheduleWildfire
	var d diag.Diagnostics
	// Handling Objects
	if !model.Recurring.IsNull() && !model.Recurring.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Recurring")
		unpacked, d := unpackUpdateScheduleUpdateScheduleWildfireRecurringToSdk(ctx, model.Recurring)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Recurring"})
		}
		if unpacked != nil {
			sdk.Recurring = *unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.UpdateScheduleUpdateScheduleWildfire", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for UpdateScheduleUpdateScheduleWildfire ---
func packUpdateScheduleUpdateScheduleWildfireFromSdk(ctx context.Context, sdk device_settings.UpdateScheduleUpdateScheduleWildfire) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.UpdateScheduleUpdateScheduleWildfire", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.UpdateScheduleUpdateScheduleWildfire
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	// Logic for non-pointer / value-type nested objects
	if !reflect.ValueOf(sdk.Recurring).IsZero() {
		tflog.Debug(ctx, "Packing nested object for field Recurring")
		packed, d := packUpdateScheduleUpdateScheduleWildfireRecurringFromSdk(ctx, sdk.Recurring)
		diags.Append(d...)
		model.Recurring = packed
	} else {
		model.Recurring = basetypes.NewObjectNull(models.UpdateScheduleUpdateScheduleWildfireRecurring{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.UpdateScheduleUpdateScheduleWildfire{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.UpdateScheduleUpdateScheduleWildfire", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for UpdateScheduleUpdateScheduleWildfire ---
func unpackUpdateScheduleUpdateScheduleWildfireListToSdk(ctx context.Context, list types.List) ([]device_settings.UpdateScheduleUpdateScheduleWildfire, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.UpdateScheduleUpdateScheduleWildfire")
	diags := diag.Diagnostics{}
	var data []models.UpdateScheduleUpdateScheduleWildfire
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]device_settings.UpdateScheduleUpdateScheduleWildfire, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.UpdateScheduleUpdateScheduleWildfire{}.AttrTypes(), &item)
		unpacked, d := unpackUpdateScheduleUpdateScheduleWildfireToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.UpdateScheduleUpdateScheduleWildfire", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for UpdateScheduleUpdateScheduleWildfire ---
func packUpdateScheduleUpdateScheduleWildfireListFromSdk(ctx context.Context, sdks []device_settings.UpdateScheduleUpdateScheduleWildfire) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.UpdateScheduleUpdateScheduleWildfire")
	diags := diag.Diagnostics{}
	var data []models.UpdateScheduleUpdateScheduleWildfire

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.UpdateScheduleUpdateScheduleWildfire
		obj, d := packUpdateScheduleUpdateScheduleWildfireFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.UpdateScheduleUpdateScheduleWildfire{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.UpdateScheduleUpdateScheduleWildfire", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.UpdateScheduleUpdateScheduleWildfire{}.AttrType(), data)
}

// --- Unpacker for UpdateScheduleUpdateScheduleWildfireRecurring ---
func unpackUpdateScheduleUpdateScheduleWildfireRecurringToSdk(ctx context.Context, obj types.Object) (*device_settings.UpdateScheduleUpdateScheduleWildfireRecurring, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.UpdateScheduleUpdateScheduleWildfireRecurring", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.UpdateScheduleUpdateScheduleWildfireRecurring
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk device_settings.UpdateScheduleUpdateScheduleWildfireRecurring
	var d diag.Diagnostics
	// Handling Objects
	if !model.Every15Mins.IsNull() && !model.Every15Mins.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Every15Mins")
		unpacked, d := unpackUpdateScheduleUpdateScheduleWildfireRecurringEvery15MinsToSdk(ctx, model.Every15Mins)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Every15Mins"})
		}
		if unpacked != nil {
			sdk.Every15Mins = unpacked
		}
	}

	// Handling Objects
	if !model.Every30Mins.IsNull() && !model.Every30Mins.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Every30Mins")
		unpacked, d := unpackUpdateScheduleUpdateScheduleWildfireRecurringEvery30MinsToSdk(ctx, model.Every30Mins)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Every30Mins"})
		}
		if unpacked != nil {
			sdk.Every30Mins = unpacked
		}
	}

	// Handling Objects
	if !model.EveryHour.IsNull() && !model.EveryHour.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field EveryHour")
		unpacked, d := unpackUpdateScheduleUpdateScheduleWildfireRecurringEveryHourToSdk(ctx, model.EveryHour)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "EveryHour"})
		}
		if unpacked != nil {
			sdk.EveryHour = unpacked
		}
	}

	// Handling Objects
	if !model.EveryMin.IsNull() && !model.EveryMin.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field EveryMin")
		unpacked, d := unpackUpdateScheduleUpdateScheduleWildfireRecurringEveryMinToSdk(ctx, model.EveryMin)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "EveryMin"})
		}
		if unpacked != nil {
			sdk.EveryMin = unpacked
		}
	}

	// Handling Typeless Objects
	if !model.None.IsNull() && !model.None.IsUnknown() {
		tflog.Debug(ctx, "Unpacking typeless object for field None")
		sdk.None = make(map[string]interface{})
	}

	// Handling Typeless Objects
	if !model.RealTime.IsNull() && !model.RealTime.IsUnknown() {
		tflog.Debug(ctx, "Unpacking typeless object for field RealTime")
		sdk.RealTime = make(map[string]interface{})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.UpdateScheduleUpdateScheduleWildfireRecurring", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for UpdateScheduleUpdateScheduleWildfireRecurring ---
func packUpdateScheduleUpdateScheduleWildfireRecurringFromSdk(ctx context.Context, sdk device_settings.UpdateScheduleUpdateScheduleWildfireRecurring) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.UpdateScheduleUpdateScheduleWildfireRecurring", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.UpdateScheduleUpdateScheduleWildfireRecurring
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Every15Mins != nil {
		tflog.Debug(ctx, "Packing nested object for field Every15Mins")
		packed, d := packUpdateScheduleUpdateScheduleWildfireRecurringEvery15MinsFromSdk(ctx, *sdk.Every15Mins)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Every15Mins"})
		}
		model.Every15Mins = packed
	} else {
		model.Every15Mins = basetypes.NewObjectNull(models.UpdateScheduleUpdateScheduleWildfireRecurringEvery15Mins{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Every30Mins != nil {
		tflog.Debug(ctx, "Packing nested object for field Every30Mins")
		packed, d := packUpdateScheduleUpdateScheduleWildfireRecurringEvery30MinsFromSdk(ctx, *sdk.Every30Mins)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Every30Mins"})
		}
		model.Every30Mins = packed
	} else {
		model.Every30Mins = basetypes.NewObjectNull(models.UpdateScheduleUpdateScheduleWildfireRecurringEvery30Mins{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.EveryHour != nil {
		tflog.Debug(ctx, "Packing nested object for field EveryHour")
		packed, d := packUpdateScheduleUpdateScheduleWildfireRecurringEveryHourFromSdk(ctx, *sdk.EveryHour)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "EveryHour"})
		}
		model.EveryHour = packed
	} else {
		model.EveryHour = basetypes.NewObjectNull(models.UpdateScheduleUpdateScheduleWildfireRecurringEveryHour{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.EveryMin != nil {
		tflog.Debug(ctx, "Packing nested object for field EveryMin")
		packed, d := packUpdateScheduleUpdateScheduleWildfireRecurringEveryMinFromSdk(ctx, *sdk.EveryMin)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "EveryMin"})
		}
		model.EveryMin = packed
	} else {
		model.EveryMin = basetypes.NewObjectNull(models.UpdateScheduleUpdateScheduleWildfireRecurringEveryMin{}.AttrTypes())
	}
	// Handling Objects
	// This is a marker object (e.g. CHAP: {}). We just need to create an empty, non-null object.
	if sdk.None != nil && !reflect.ValueOf(sdk.None).IsNil() {
		tflog.Debug(ctx, "Packing typeless object for field None")
		var d diag.Diagnostics
		// Create an empty object with no attributes, which signifies its presence.
		model.None, d = basetypes.NewObjectValue(map[string]attr.Type{}, map[string]attr.Value{})
		diags.Append(d...)
	} else {
		// Since this field is part of a oneOf, being nil means it's not selected.
		// We make the object null with an empty attribute map.
		model.None = basetypes.NewObjectNull(map[string]attr.Type{})
	}
	// Handling Objects
	// This is a marker object (e.g. CHAP: {}). We just need to create an empty, non-null object.
	if sdk.RealTime != nil && !reflect.ValueOf(sdk.RealTime).IsNil() {
		tflog.Debug(ctx, "Packing typeless object for field RealTime")
		var d diag.Diagnostics
		// Create an empty object with no attributes, which signifies its presence.
		model.RealTime, d = basetypes.NewObjectValue(map[string]attr.Type{}, map[string]attr.Value{})
		diags.Append(d...)
	} else {
		// Since this field is part of a oneOf, being nil means it's not selected.
		// We make the object null with an empty attribute map.
		model.RealTime = basetypes.NewObjectNull(map[string]attr.Type{})
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.UpdateScheduleUpdateScheduleWildfireRecurring{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.UpdateScheduleUpdateScheduleWildfireRecurring", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for UpdateScheduleUpdateScheduleWildfireRecurring ---
func unpackUpdateScheduleUpdateScheduleWildfireRecurringListToSdk(ctx context.Context, list types.List) ([]device_settings.UpdateScheduleUpdateScheduleWildfireRecurring, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.UpdateScheduleUpdateScheduleWildfireRecurring")
	diags := diag.Diagnostics{}
	var data []models.UpdateScheduleUpdateScheduleWildfireRecurring
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]device_settings.UpdateScheduleUpdateScheduleWildfireRecurring, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.UpdateScheduleUpdateScheduleWildfireRecurring{}.AttrTypes(), &item)
		unpacked, d := unpackUpdateScheduleUpdateScheduleWildfireRecurringToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.UpdateScheduleUpdateScheduleWildfireRecurring", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for UpdateScheduleUpdateScheduleWildfireRecurring ---
func packUpdateScheduleUpdateScheduleWildfireRecurringListFromSdk(ctx context.Context, sdks []device_settings.UpdateScheduleUpdateScheduleWildfireRecurring) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.UpdateScheduleUpdateScheduleWildfireRecurring")
	diags := diag.Diagnostics{}
	var data []models.UpdateScheduleUpdateScheduleWildfireRecurring

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.UpdateScheduleUpdateScheduleWildfireRecurring
		obj, d := packUpdateScheduleUpdateScheduleWildfireRecurringFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.UpdateScheduleUpdateScheduleWildfireRecurring{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.UpdateScheduleUpdateScheduleWildfireRecurring", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.UpdateScheduleUpdateScheduleWildfireRecurring{}.AttrType(), data)
}

// --- Unpacker for UpdateScheduleUpdateScheduleWildfireRecurringEvery15Mins ---
func unpackUpdateScheduleUpdateScheduleWildfireRecurringEvery15MinsToSdk(ctx context.Context, obj types.Object) (*device_settings.UpdateScheduleUpdateScheduleWildfireRecurringEvery15Mins, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.UpdateScheduleUpdateScheduleWildfireRecurringEvery15Mins", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.UpdateScheduleUpdateScheduleWildfireRecurringEvery15Mins
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk device_settings.UpdateScheduleUpdateScheduleWildfireRecurringEvery15Mins
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Action.IsNull() && !model.Action.IsUnknown() {
		sdk.Action = model.Action.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Action", "value": *sdk.Action})
	}

	// Handling Primitives
	if !model.At.IsNull() && !model.At.IsUnknown() {
		val := int32(model.At.ValueInt64())
		sdk.At = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "At", "value": *sdk.At})
	}

	// Handling Primitives
	if !model.SyncToPeer.IsNull() && !model.SyncToPeer.IsUnknown() {
		sdk.SyncToPeer = model.SyncToPeer.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "SyncToPeer", "value": *sdk.SyncToPeer})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.UpdateScheduleUpdateScheduleWildfireRecurringEvery15Mins", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for UpdateScheduleUpdateScheduleWildfireRecurringEvery15Mins ---
func packUpdateScheduleUpdateScheduleWildfireRecurringEvery15MinsFromSdk(ctx context.Context, sdk device_settings.UpdateScheduleUpdateScheduleWildfireRecurringEvery15Mins) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.UpdateScheduleUpdateScheduleWildfireRecurringEvery15Mins", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.UpdateScheduleUpdateScheduleWildfireRecurringEvery15Mins
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Action != nil {
		model.Action = basetypes.NewStringValue(*sdk.Action)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Action", "value": *sdk.Action})
	} else {
		model.Action = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.At != nil {
		model.At = basetypes.NewInt64Value(int64(*sdk.At))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "At", "value": *sdk.At})
	} else {
		model.At = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.SyncToPeer != nil {
		model.SyncToPeer = basetypes.NewBoolValue(*sdk.SyncToPeer)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "SyncToPeer", "value": *sdk.SyncToPeer})
	} else {
		model.SyncToPeer = basetypes.NewBoolNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.UpdateScheduleUpdateScheduleWildfireRecurringEvery15Mins{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.UpdateScheduleUpdateScheduleWildfireRecurringEvery15Mins", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for UpdateScheduleUpdateScheduleWildfireRecurringEvery15Mins ---
func unpackUpdateScheduleUpdateScheduleWildfireRecurringEvery15MinsListToSdk(ctx context.Context, list types.List) ([]device_settings.UpdateScheduleUpdateScheduleWildfireRecurringEvery15Mins, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.UpdateScheduleUpdateScheduleWildfireRecurringEvery15Mins")
	diags := diag.Diagnostics{}
	var data []models.UpdateScheduleUpdateScheduleWildfireRecurringEvery15Mins
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]device_settings.UpdateScheduleUpdateScheduleWildfireRecurringEvery15Mins, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.UpdateScheduleUpdateScheduleWildfireRecurringEvery15Mins{}.AttrTypes(), &item)
		unpacked, d := unpackUpdateScheduleUpdateScheduleWildfireRecurringEvery15MinsToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.UpdateScheduleUpdateScheduleWildfireRecurringEvery15Mins", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for UpdateScheduleUpdateScheduleWildfireRecurringEvery15Mins ---
func packUpdateScheduleUpdateScheduleWildfireRecurringEvery15MinsListFromSdk(ctx context.Context, sdks []device_settings.UpdateScheduleUpdateScheduleWildfireRecurringEvery15Mins) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.UpdateScheduleUpdateScheduleWildfireRecurringEvery15Mins")
	diags := diag.Diagnostics{}
	var data []models.UpdateScheduleUpdateScheduleWildfireRecurringEvery15Mins

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.UpdateScheduleUpdateScheduleWildfireRecurringEvery15Mins
		obj, d := packUpdateScheduleUpdateScheduleWildfireRecurringEvery15MinsFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.UpdateScheduleUpdateScheduleWildfireRecurringEvery15Mins{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.UpdateScheduleUpdateScheduleWildfireRecurringEvery15Mins", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.UpdateScheduleUpdateScheduleWildfireRecurringEvery15Mins{}.AttrType(), data)
}

// --- Unpacker for UpdateScheduleUpdateScheduleWildfireRecurringEvery30Mins ---
func unpackUpdateScheduleUpdateScheduleWildfireRecurringEvery30MinsToSdk(ctx context.Context, obj types.Object) (*device_settings.UpdateScheduleUpdateScheduleWildfireRecurringEvery30Mins, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.UpdateScheduleUpdateScheduleWildfireRecurringEvery30Mins", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.UpdateScheduleUpdateScheduleWildfireRecurringEvery30Mins
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk device_settings.UpdateScheduleUpdateScheduleWildfireRecurringEvery30Mins
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Action.IsNull() && !model.Action.IsUnknown() {
		sdk.Action = model.Action.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Action", "value": *sdk.Action})
	}

	// Handling Primitives
	if !model.At.IsNull() && !model.At.IsUnknown() {
		val := int32(model.At.ValueInt64())
		sdk.At = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "At", "value": *sdk.At})
	}

	// Handling Primitives
	if !model.SyncToPeer.IsNull() && !model.SyncToPeer.IsUnknown() {
		sdk.SyncToPeer = model.SyncToPeer.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "SyncToPeer", "value": *sdk.SyncToPeer})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.UpdateScheduleUpdateScheduleWildfireRecurringEvery30Mins", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for UpdateScheduleUpdateScheduleWildfireRecurringEvery30Mins ---
func packUpdateScheduleUpdateScheduleWildfireRecurringEvery30MinsFromSdk(ctx context.Context, sdk device_settings.UpdateScheduleUpdateScheduleWildfireRecurringEvery30Mins) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.UpdateScheduleUpdateScheduleWildfireRecurringEvery30Mins", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.UpdateScheduleUpdateScheduleWildfireRecurringEvery30Mins
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Action != nil {
		model.Action = basetypes.NewStringValue(*sdk.Action)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Action", "value": *sdk.Action})
	} else {
		model.Action = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.At != nil {
		model.At = basetypes.NewInt64Value(int64(*sdk.At))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "At", "value": *sdk.At})
	} else {
		model.At = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.SyncToPeer != nil {
		model.SyncToPeer = basetypes.NewBoolValue(*sdk.SyncToPeer)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "SyncToPeer", "value": *sdk.SyncToPeer})
	} else {
		model.SyncToPeer = basetypes.NewBoolNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.UpdateScheduleUpdateScheduleWildfireRecurringEvery30Mins{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.UpdateScheduleUpdateScheduleWildfireRecurringEvery30Mins", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for UpdateScheduleUpdateScheduleWildfireRecurringEvery30Mins ---
func unpackUpdateScheduleUpdateScheduleWildfireRecurringEvery30MinsListToSdk(ctx context.Context, list types.List) ([]device_settings.UpdateScheduleUpdateScheduleWildfireRecurringEvery30Mins, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.UpdateScheduleUpdateScheduleWildfireRecurringEvery30Mins")
	diags := diag.Diagnostics{}
	var data []models.UpdateScheduleUpdateScheduleWildfireRecurringEvery30Mins
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]device_settings.UpdateScheduleUpdateScheduleWildfireRecurringEvery30Mins, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.UpdateScheduleUpdateScheduleWildfireRecurringEvery30Mins{}.AttrTypes(), &item)
		unpacked, d := unpackUpdateScheduleUpdateScheduleWildfireRecurringEvery30MinsToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.UpdateScheduleUpdateScheduleWildfireRecurringEvery30Mins", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for UpdateScheduleUpdateScheduleWildfireRecurringEvery30Mins ---
func packUpdateScheduleUpdateScheduleWildfireRecurringEvery30MinsListFromSdk(ctx context.Context, sdks []device_settings.UpdateScheduleUpdateScheduleWildfireRecurringEvery30Mins) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.UpdateScheduleUpdateScheduleWildfireRecurringEvery30Mins")
	diags := diag.Diagnostics{}
	var data []models.UpdateScheduleUpdateScheduleWildfireRecurringEvery30Mins

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.UpdateScheduleUpdateScheduleWildfireRecurringEvery30Mins
		obj, d := packUpdateScheduleUpdateScheduleWildfireRecurringEvery30MinsFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.UpdateScheduleUpdateScheduleWildfireRecurringEvery30Mins{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.UpdateScheduleUpdateScheduleWildfireRecurringEvery30Mins", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.UpdateScheduleUpdateScheduleWildfireRecurringEvery30Mins{}.AttrType(), data)
}

// --- Unpacker for UpdateScheduleUpdateScheduleWildfireRecurringEveryHour ---
func unpackUpdateScheduleUpdateScheduleWildfireRecurringEveryHourToSdk(ctx context.Context, obj types.Object) (*device_settings.UpdateScheduleUpdateScheduleWildfireRecurringEveryHour, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.UpdateScheduleUpdateScheduleWildfireRecurringEveryHour", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.UpdateScheduleUpdateScheduleWildfireRecurringEveryHour
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk device_settings.UpdateScheduleUpdateScheduleWildfireRecurringEveryHour
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Action.IsNull() && !model.Action.IsUnknown() {
		sdk.Action = model.Action.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Action", "value": *sdk.Action})
	}

	// Handling Primitives
	if !model.At.IsNull() && !model.At.IsUnknown() {
		val := int32(model.At.ValueInt64())
		sdk.At = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "At", "value": *sdk.At})
	}

	// Handling Primitives
	if !model.SyncToPeer.IsNull() && !model.SyncToPeer.IsUnknown() {
		sdk.SyncToPeer = model.SyncToPeer.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "SyncToPeer", "value": *sdk.SyncToPeer})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.UpdateScheduleUpdateScheduleWildfireRecurringEveryHour", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for UpdateScheduleUpdateScheduleWildfireRecurringEveryHour ---
func packUpdateScheduleUpdateScheduleWildfireRecurringEveryHourFromSdk(ctx context.Context, sdk device_settings.UpdateScheduleUpdateScheduleWildfireRecurringEveryHour) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.UpdateScheduleUpdateScheduleWildfireRecurringEveryHour", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.UpdateScheduleUpdateScheduleWildfireRecurringEveryHour
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Action != nil {
		model.Action = basetypes.NewStringValue(*sdk.Action)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Action", "value": *sdk.Action})
	} else {
		model.Action = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.At != nil {
		model.At = basetypes.NewInt64Value(int64(*sdk.At))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "At", "value": *sdk.At})
	} else {
		model.At = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.SyncToPeer != nil {
		model.SyncToPeer = basetypes.NewBoolValue(*sdk.SyncToPeer)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "SyncToPeer", "value": *sdk.SyncToPeer})
	} else {
		model.SyncToPeer = basetypes.NewBoolNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.UpdateScheduleUpdateScheduleWildfireRecurringEveryHour{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.UpdateScheduleUpdateScheduleWildfireRecurringEveryHour", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for UpdateScheduleUpdateScheduleWildfireRecurringEveryHour ---
func unpackUpdateScheduleUpdateScheduleWildfireRecurringEveryHourListToSdk(ctx context.Context, list types.List) ([]device_settings.UpdateScheduleUpdateScheduleWildfireRecurringEveryHour, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.UpdateScheduleUpdateScheduleWildfireRecurringEveryHour")
	diags := diag.Diagnostics{}
	var data []models.UpdateScheduleUpdateScheduleWildfireRecurringEveryHour
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]device_settings.UpdateScheduleUpdateScheduleWildfireRecurringEveryHour, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.UpdateScheduleUpdateScheduleWildfireRecurringEveryHour{}.AttrTypes(), &item)
		unpacked, d := unpackUpdateScheduleUpdateScheduleWildfireRecurringEveryHourToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.UpdateScheduleUpdateScheduleWildfireRecurringEveryHour", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for UpdateScheduleUpdateScheduleWildfireRecurringEveryHour ---
func packUpdateScheduleUpdateScheduleWildfireRecurringEveryHourListFromSdk(ctx context.Context, sdks []device_settings.UpdateScheduleUpdateScheduleWildfireRecurringEveryHour) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.UpdateScheduleUpdateScheduleWildfireRecurringEveryHour")
	diags := diag.Diagnostics{}
	var data []models.UpdateScheduleUpdateScheduleWildfireRecurringEveryHour

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.UpdateScheduleUpdateScheduleWildfireRecurringEveryHour
		obj, d := packUpdateScheduleUpdateScheduleWildfireRecurringEveryHourFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.UpdateScheduleUpdateScheduleWildfireRecurringEveryHour{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.UpdateScheduleUpdateScheduleWildfireRecurringEveryHour", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.UpdateScheduleUpdateScheduleWildfireRecurringEveryHour{}.AttrType(), data)
}

// --- Unpacker for UpdateScheduleUpdateScheduleWildfireRecurringEveryMin ---
func unpackUpdateScheduleUpdateScheduleWildfireRecurringEveryMinToSdk(ctx context.Context, obj types.Object) (*device_settings.UpdateScheduleUpdateScheduleWildfireRecurringEveryMin, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.UpdateScheduleUpdateScheduleWildfireRecurringEveryMin", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.UpdateScheduleUpdateScheduleWildfireRecurringEveryMin
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk device_settings.UpdateScheduleUpdateScheduleWildfireRecurringEveryMin
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Action.IsNull() && !model.Action.IsUnknown() {
		sdk.Action = model.Action.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Action", "value": *sdk.Action})
	}

	// Handling Primitives
	if !model.SyncToPeer.IsNull() && !model.SyncToPeer.IsUnknown() {
		sdk.SyncToPeer = model.SyncToPeer.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "SyncToPeer", "value": *sdk.SyncToPeer})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.UpdateScheduleUpdateScheduleWildfireRecurringEveryMin", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for UpdateScheduleUpdateScheduleWildfireRecurringEveryMin ---
func packUpdateScheduleUpdateScheduleWildfireRecurringEveryMinFromSdk(ctx context.Context, sdk device_settings.UpdateScheduleUpdateScheduleWildfireRecurringEveryMin) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.UpdateScheduleUpdateScheduleWildfireRecurringEveryMin", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.UpdateScheduleUpdateScheduleWildfireRecurringEveryMin
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Action != nil {
		model.Action = basetypes.NewStringValue(*sdk.Action)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Action", "value": *sdk.Action})
	} else {
		model.Action = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.SyncToPeer != nil {
		model.SyncToPeer = basetypes.NewBoolValue(*sdk.SyncToPeer)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "SyncToPeer", "value": *sdk.SyncToPeer})
	} else {
		model.SyncToPeer = basetypes.NewBoolNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.UpdateScheduleUpdateScheduleWildfireRecurringEveryMin{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.UpdateScheduleUpdateScheduleWildfireRecurringEveryMin", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for UpdateScheduleUpdateScheduleWildfireRecurringEveryMin ---
func unpackUpdateScheduleUpdateScheduleWildfireRecurringEveryMinListToSdk(ctx context.Context, list types.List) ([]device_settings.UpdateScheduleUpdateScheduleWildfireRecurringEveryMin, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.UpdateScheduleUpdateScheduleWildfireRecurringEveryMin")
	diags := diag.Diagnostics{}
	var data []models.UpdateScheduleUpdateScheduleWildfireRecurringEveryMin
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]device_settings.UpdateScheduleUpdateScheduleWildfireRecurringEveryMin, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.UpdateScheduleUpdateScheduleWildfireRecurringEveryMin{}.AttrTypes(), &item)
		unpacked, d := unpackUpdateScheduleUpdateScheduleWildfireRecurringEveryMinToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.UpdateScheduleUpdateScheduleWildfireRecurringEveryMin", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for UpdateScheduleUpdateScheduleWildfireRecurringEveryMin ---
func packUpdateScheduleUpdateScheduleWildfireRecurringEveryMinListFromSdk(ctx context.Context, sdks []device_settings.UpdateScheduleUpdateScheduleWildfireRecurringEveryMin) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.UpdateScheduleUpdateScheduleWildfireRecurringEveryMin")
	diags := diag.Diagnostics{}
	var data []models.UpdateScheduleUpdateScheduleWildfireRecurringEveryMin

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.UpdateScheduleUpdateScheduleWildfireRecurringEveryMin
		obj, d := packUpdateScheduleUpdateScheduleWildfireRecurringEveryMinFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.UpdateScheduleUpdateScheduleWildfireRecurringEveryMin{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.UpdateScheduleUpdateScheduleWildfireRecurringEveryMin", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.UpdateScheduleUpdateScheduleWildfireRecurringEveryMin{}.AttrType(), data)
}
