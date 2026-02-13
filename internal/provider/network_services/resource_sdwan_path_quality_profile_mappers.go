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

// --- Unpacker for SdwanPathQualityProfiles ---
func unpackSdwanPathQualityProfilesToSdk(ctx context.Context, obj types.Object) (*network_services.SdwanPathQualityProfiles, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.SdwanPathQualityProfiles", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.SdwanPathQualityProfiles
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.SdwanPathQualityProfiles
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
	if !model.Metric.IsNull() && !model.Metric.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Metric")
		unpacked, d := unpackSdwanPathQualityProfilesMetricToSdk(ctx, model.Metric)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Metric"})
		}
		if unpacked != nil {
			sdk.Metric = *unpacked
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

	tflog.Debug(ctx, "Exiting unpack helper for models.SdwanPathQualityProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for SdwanPathQualityProfiles ---
func packSdwanPathQualityProfilesFromSdk(ctx context.Context, sdk network_services.SdwanPathQualityProfiles) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.SdwanPathQualityProfiles", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.SdwanPathQualityProfiles
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
	// Logic for non-pointer / value-type nested objects
	if !reflect.ValueOf(sdk.Metric).IsZero() {
		tflog.Debug(ctx, "Packing nested object for field Metric")
		packed, d := packSdwanPathQualityProfilesMetricFromSdk(ctx, sdk.Metric)
		diags.Append(d...)
		model.Metric = packed
	} else {
		model.Metric = basetypes.NewObjectNull(models.SdwanPathQualityProfilesMetric{}.AttrTypes())
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

	obj, d := types.ObjectValueFrom(ctx, models.SdwanPathQualityProfiles{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.SdwanPathQualityProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for SdwanPathQualityProfiles ---
func unpackSdwanPathQualityProfilesListToSdk(ctx context.Context, list types.List) ([]network_services.SdwanPathQualityProfiles, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.SdwanPathQualityProfiles")
	diags := diag.Diagnostics{}
	var data []models.SdwanPathQualityProfiles
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.SdwanPathQualityProfiles, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.SdwanPathQualityProfiles{}.AttrTypes(), &item)
		unpacked, d := unpackSdwanPathQualityProfilesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.SdwanPathQualityProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for SdwanPathQualityProfiles ---
func packSdwanPathQualityProfilesListFromSdk(ctx context.Context, sdks []network_services.SdwanPathQualityProfiles) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.SdwanPathQualityProfiles")
	diags := diag.Diagnostics{}
	var data []models.SdwanPathQualityProfiles

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.SdwanPathQualityProfiles
		obj, d := packSdwanPathQualityProfilesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.SdwanPathQualityProfiles{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.SdwanPathQualityProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.SdwanPathQualityProfiles{}.AttrType(), data)
}

// --- Unpacker for SdwanPathQualityProfilesMetric ---
func unpackSdwanPathQualityProfilesMetricToSdk(ctx context.Context, obj types.Object) (*network_services.SdwanPathQualityProfilesMetric, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.SdwanPathQualityProfilesMetric", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.SdwanPathQualityProfilesMetric
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.SdwanPathQualityProfilesMetric
	var d diag.Diagnostics
	// Handling Objects
	if !model.Jitter.IsNull() && !model.Jitter.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Jitter")
		unpacked, d := unpackSdwanPathQualityProfilesMetricJitterToSdk(ctx, model.Jitter)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Jitter"})
		}
		if unpacked != nil {
			sdk.Jitter = *unpacked
		}
	}

	// Handling Objects
	if !model.Latency.IsNull() && !model.Latency.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Latency")
		unpacked, d := unpackSdwanPathQualityProfilesMetricLatencyToSdk(ctx, model.Latency)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Latency"})
		}
		if unpacked != nil {
			sdk.Latency = *unpacked
		}
	}

	// Handling Objects
	if !model.PktLoss.IsNull() && !model.PktLoss.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field PktLoss")
		unpacked, d := unpackSdwanPathQualityProfilesMetricPktLossToSdk(ctx, model.PktLoss)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "PktLoss"})
		}
		if unpacked != nil {
			sdk.PktLoss = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.SdwanPathQualityProfilesMetric", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for SdwanPathQualityProfilesMetric ---
func packSdwanPathQualityProfilesMetricFromSdk(ctx context.Context, sdk network_services.SdwanPathQualityProfilesMetric) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.SdwanPathQualityProfilesMetric", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.SdwanPathQualityProfilesMetric
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	// Logic for non-pointer / value-type nested objects
	if !reflect.ValueOf(sdk.Jitter).IsZero() {
		tflog.Debug(ctx, "Packing nested object for field Jitter")
		packed, d := packSdwanPathQualityProfilesMetricJitterFromSdk(ctx, sdk.Jitter)
		diags.Append(d...)
		model.Jitter = packed
	} else {
		model.Jitter = basetypes.NewObjectNull(models.SdwanPathQualityProfilesMetricJitter{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	// Logic for non-pointer / value-type nested objects
	if !reflect.ValueOf(sdk.Latency).IsZero() {
		tflog.Debug(ctx, "Packing nested object for field Latency")
		packed, d := packSdwanPathQualityProfilesMetricLatencyFromSdk(ctx, sdk.Latency)
		diags.Append(d...)
		model.Latency = packed
	} else {
		model.Latency = basetypes.NewObjectNull(models.SdwanPathQualityProfilesMetricLatency{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.PktLoss != nil {
		tflog.Debug(ctx, "Packing nested object for field PktLoss")
		packed, d := packSdwanPathQualityProfilesMetricPktLossFromSdk(ctx, *sdk.PktLoss)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "PktLoss"})
		}
		model.PktLoss = packed
	} else {
		model.PktLoss = basetypes.NewObjectNull(models.SdwanPathQualityProfilesMetricPktLoss{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.SdwanPathQualityProfilesMetric{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.SdwanPathQualityProfilesMetric", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for SdwanPathQualityProfilesMetric ---
func unpackSdwanPathQualityProfilesMetricListToSdk(ctx context.Context, list types.List) ([]network_services.SdwanPathQualityProfilesMetric, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.SdwanPathQualityProfilesMetric")
	diags := diag.Diagnostics{}
	var data []models.SdwanPathQualityProfilesMetric
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.SdwanPathQualityProfilesMetric, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.SdwanPathQualityProfilesMetric{}.AttrTypes(), &item)
		unpacked, d := unpackSdwanPathQualityProfilesMetricToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.SdwanPathQualityProfilesMetric", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for SdwanPathQualityProfilesMetric ---
func packSdwanPathQualityProfilesMetricListFromSdk(ctx context.Context, sdks []network_services.SdwanPathQualityProfilesMetric) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.SdwanPathQualityProfilesMetric")
	diags := diag.Diagnostics{}
	var data []models.SdwanPathQualityProfilesMetric

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.SdwanPathQualityProfilesMetric
		obj, d := packSdwanPathQualityProfilesMetricFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.SdwanPathQualityProfilesMetric{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.SdwanPathQualityProfilesMetric", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.SdwanPathQualityProfilesMetric{}.AttrType(), data)
}

// --- Unpacker for SdwanPathQualityProfilesMetricJitter ---
func unpackSdwanPathQualityProfilesMetricJitterToSdk(ctx context.Context, obj types.Object) (*network_services.SdwanPathQualityProfilesMetricJitter, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.SdwanPathQualityProfilesMetricJitter", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.SdwanPathQualityProfilesMetricJitter
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.SdwanPathQualityProfilesMetricJitter
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Sensitivity.IsNull() && !model.Sensitivity.IsUnknown() {
		sdk.Sensitivity = model.Sensitivity.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Sensitivity", "value": sdk.Sensitivity})
	}

	// Handling Primitives
	if !model.Threshold.IsNull() && !model.Threshold.IsUnknown() {
		sdk.Threshold = int32(model.Threshold.ValueInt64())
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Threshold", "value": sdk.Threshold})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.SdwanPathQualityProfilesMetricJitter", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for SdwanPathQualityProfilesMetricJitter ---
func packSdwanPathQualityProfilesMetricJitterFromSdk(ctx context.Context, sdk network_services.SdwanPathQualityProfilesMetricJitter) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.SdwanPathQualityProfilesMetricJitter", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.SdwanPathQualityProfilesMetricJitter
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.Sensitivity = basetypes.NewStringValue(sdk.Sensitivity)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Sensitivity", "value": sdk.Sensitivity})
	// Handling Primitives
	// Standard primitive packing
	model.Threshold = basetypes.NewInt64Value(int64(sdk.Threshold))
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.SdwanPathQualityProfilesMetricJitter{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.SdwanPathQualityProfilesMetricJitter", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for SdwanPathQualityProfilesMetricJitter ---
func unpackSdwanPathQualityProfilesMetricJitterListToSdk(ctx context.Context, list types.List) ([]network_services.SdwanPathQualityProfilesMetricJitter, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.SdwanPathQualityProfilesMetricJitter")
	diags := diag.Diagnostics{}
	var data []models.SdwanPathQualityProfilesMetricJitter
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.SdwanPathQualityProfilesMetricJitter, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.SdwanPathQualityProfilesMetricJitter{}.AttrTypes(), &item)
		unpacked, d := unpackSdwanPathQualityProfilesMetricJitterToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.SdwanPathQualityProfilesMetricJitter", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for SdwanPathQualityProfilesMetricJitter ---
func packSdwanPathQualityProfilesMetricJitterListFromSdk(ctx context.Context, sdks []network_services.SdwanPathQualityProfilesMetricJitter) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.SdwanPathQualityProfilesMetricJitter")
	diags := diag.Diagnostics{}
	var data []models.SdwanPathQualityProfilesMetricJitter

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.SdwanPathQualityProfilesMetricJitter
		obj, d := packSdwanPathQualityProfilesMetricJitterFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.SdwanPathQualityProfilesMetricJitter{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.SdwanPathQualityProfilesMetricJitter", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.SdwanPathQualityProfilesMetricJitter{}.AttrType(), data)
}

// --- Unpacker for SdwanPathQualityProfilesMetricLatency ---
func unpackSdwanPathQualityProfilesMetricLatencyToSdk(ctx context.Context, obj types.Object) (*network_services.SdwanPathQualityProfilesMetricLatency, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.SdwanPathQualityProfilesMetricLatency", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.SdwanPathQualityProfilesMetricLatency
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.SdwanPathQualityProfilesMetricLatency
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Sensitivity.IsNull() && !model.Sensitivity.IsUnknown() {
		sdk.Sensitivity = model.Sensitivity.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Sensitivity", "value": sdk.Sensitivity})
	}

	// Handling Primitives
	if !model.Threshold.IsNull() && !model.Threshold.IsUnknown() {
		sdk.Threshold = int32(model.Threshold.ValueInt64())
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Threshold", "value": sdk.Threshold})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.SdwanPathQualityProfilesMetricLatency", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for SdwanPathQualityProfilesMetricLatency ---
func packSdwanPathQualityProfilesMetricLatencyFromSdk(ctx context.Context, sdk network_services.SdwanPathQualityProfilesMetricLatency) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.SdwanPathQualityProfilesMetricLatency", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.SdwanPathQualityProfilesMetricLatency
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.Sensitivity = basetypes.NewStringValue(sdk.Sensitivity)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Sensitivity", "value": sdk.Sensitivity})
	// Handling Primitives
	// Standard primitive packing
	model.Threshold = basetypes.NewInt64Value(int64(sdk.Threshold))
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.SdwanPathQualityProfilesMetricLatency{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.SdwanPathQualityProfilesMetricLatency", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for SdwanPathQualityProfilesMetricLatency ---
func unpackSdwanPathQualityProfilesMetricLatencyListToSdk(ctx context.Context, list types.List) ([]network_services.SdwanPathQualityProfilesMetricLatency, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.SdwanPathQualityProfilesMetricLatency")
	diags := diag.Diagnostics{}
	var data []models.SdwanPathQualityProfilesMetricLatency
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.SdwanPathQualityProfilesMetricLatency, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.SdwanPathQualityProfilesMetricLatency{}.AttrTypes(), &item)
		unpacked, d := unpackSdwanPathQualityProfilesMetricLatencyToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.SdwanPathQualityProfilesMetricLatency", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for SdwanPathQualityProfilesMetricLatency ---
func packSdwanPathQualityProfilesMetricLatencyListFromSdk(ctx context.Context, sdks []network_services.SdwanPathQualityProfilesMetricLatency) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.SdwanPathQualityProfilesMetricLatency")
	diags := diag.Diagnostics{}
	var data []models.SdwanPathQualityProfilesMetricLatency

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.SdwanPathQualityProfilesMetricLatency
		obj, d := packSdwanPathQualityProfilesMetricLatencyFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.SdwanPathQualityProfilesMetricLatency{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.SdwanPathQualityProfilesMetricLatency", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.SdwanPathQualityProfilesMetricLatency{}.AttrType(), data)
}

// --- Unpacker for SdwanPathQualityProfilesMetricPktLoss ---
func unpackSdwanPathQualityProfilesMetricPktLossToSdk(ctx context.Context, obj types.Object) (*network_services.SdwanPathQualityProfilesMetricPktLoss, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.SdwanPathQualityProfilesMetricPktLoss", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.SdwanPathQualityProfilesMetricPktLoss
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.SdwanPathQualityProfilesMetricPktLoss
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Sensitivity.IsNull() && !model.Sensitivity.IsUnknown() {
		sdk.Sensitivity = model.Sensitivity.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Sensitivity", "value": sdk.Sensitivity})
	}

	// Handling Primitives
	if !model.Threshold.IsNull() && !model.Threshold.IsUnknown() {
		sdk.Threshold = int32(model.Threshold.ValueInt64())
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Threshold", "value": sdk.Threshold})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.SdwanPathQualityProfilesMetricPktLoss", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for SdwanPathQualityProfilesMetricPktLoss ---
func packSdwanPathQualityProfilesMetricPktLossFromSdk(ctx context.Context, sdk network_services.SdwanPathQualityProfilesMetricPktLoss) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.SdwanPathQualityProfilesMetricPktLoss", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.SdwanPathQualityProfilesMetricPktLoss
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.Sensitivity = basetypes.NewStringValue(sdk.Sensitivity)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Sensitivity", "value": sdk.Sensitivity})
	// Handling Primitives
	// Standard primitive packing
	model.Threshold = basetypes.NewInt64Value(int64(sdk.Threshold))
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.SdwanPathQualityProfilesMetricPktLoss{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.SdwanPathQualityProfilesMetricPktLoss", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for SdwanPathQualityProfilesMetricPktLoss ---
func unpackSdwanPathQualityProfilesMetricPktLossListToSdk(ctx context.Context, list types.List) ([]network_services.SdwanPathQualityProfilesMetricPktLoss, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.SdwanPathQualityProfilesMetricPktLoss")
	diags := diag.Diagnostics{}
	var data []models.SdwanPathQualityProfilesMetricPktLoss
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.SdwanPathQualityProfilesMetricPktLoss, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.SdwanPathQualityProfilesMetricPktLoss{}.AttrTypes(), &item)
		unpacked, d := unpackSdwanPathQualityProfilesMetricPktLossToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.SdwanPathQualityProfilesMetricPktLoss", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for SdwanPathQualityProfilesMetricPktLoss ---
func packSdwanPathQualityProfilesMetricPktLossListFromSdk(ctx context.Context, sdks []network_services.SdwanPathQualityProfilesMetricPktLoss) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.SdwanPathQualityProfilesMetricPktLoss")
	diags := diag.Diagnostics{}
	var data []models.SdwanPathQualityProfilesMetricPktLoss

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.SdwanPathQualityProfilesMetricPktLoss
		obj, d := packSdwanPathQualityProfilesMetricPktLossFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.SdwanPathQualityProfilesMetricPktLoss{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.SdwanPathQualityProfilesMetricPktLoss", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.SdwanPathQualityProfilesMetricPktLoss{}.AttrType(), data)
}
