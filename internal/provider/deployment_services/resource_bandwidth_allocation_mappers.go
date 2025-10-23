package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/deployment_services"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/deployment_services"
)

// --- Unpacker for BandwidthAllocations ---
func unpackBandwidthAllocationsToSdk(ctx context.Context, obj types.Object) (*deployment_services.BandwidthAllocations, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BandwidthAllocations", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BandwidthAllocations
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk deployment_services.BandwidthAllocations
	var d diag.Diagnostics

	// Handling Primitives
	if !model.AllocatedBandwidth.IsNull() && !model.AllocatedBandwidth.IsUnknown() {
		sdk.AllocatedBandwidth = float32(model.AllocatedBandwidth.ValueFloat64())
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "AllocatedBandwidth", "value": sdk.AllocatedBandwidth})
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	// Handling Objects
	if !model.Qos.IsNull() && !model.Qos.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Qos")
		unpacked, d := unpackBandwidthAllocationsQosToSdk(ctx, model.Qos)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Qos"})
		}
		if unpacked != nil {
			sdk.Qos = unpacked
		}
	}

	// Handling Lists
	if !model.SpnNameList.IsNull() && !model.SpnNameList.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field SpnNameList")
		diags.Append(model.SpnNameList.ElementsAs(ctx, &sdk.SpnNameList, false)...)
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BandwidthAllocations", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BandwidthAllocations ---
func packBandwidthAllocationsFromSdk(ctx context.Context, sdk deployment_services.BandwidthAllocations) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BandwidthAllocations", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BandwidthAllocations
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.AllocatedBandwidth = basetypes.NewFloat64Value(float64(sdk.AllocatedBandwidth))
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "AllocatedBandwidth", "value": sdk.AllocatedBandwidth})
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Qos != nil {
		tflog.Debug(ctx, "Packing nested object for field Qos")
		packed, d := packBandwidthAllocationsQosFromSdk(ctx, *sdk.Qos)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Qos"})
		}
		model.Qos = packed
	} else {
		model.Qos = basetypes.NewObjectNull(models.BandwidthAllocationsQos{}.AttrTypes())
	}
	// Handling Lists
	if sdk.SpnNameList != nil {
		tflog.Debug(ctx, "Packing list of primitives for field SpnNameList")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.SpnNameList, d = basetypes.NewListValueFrom(ctx, elemType, sdk.SpnNameList)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.SpnNameList = basetypes.NewListNull(elemType)
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BandwidthAllocations{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BandwidthAllocations", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BandwidthAllocations ---
func unpackBandwidthAllocationsListToSdk(ctx context.Context, list types.List) ([]deployment_services.BandwidthAllocations, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BandwidthAllocations")
	diags := diag.Diagnostics{}
	var data []models.BandwidthAllocations
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]deployment_services.BandwidthAllocations, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BandwidthAllocations{}.AttrTypes(), &item)
		unpacked, d := unpackBandwidthAllocationsToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BandwidthAllocations", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BandwidthAllocations ---
func packBandwidthAllocationsListFromSdk(ctx context.Context, sdks []deployment_services.BandwidthAllocations) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BandwidthAllocations")
	diags := diag.Diagnostics{}
	var data []models.BandwidthAllocations

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BandwidthAllocations
		obj, d := packBandwidthAllocationsFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BandwidthAllocations{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BandwidthAllocations", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BandwidthAllocations{}.AttrType(), data)
}

// --- Unpacker for BandwidthAllocationsQos ---
func unpackBandwidthAllocationsQosToSdk(ctx context.Context, obj types.Object) (*deployment_services.BandwidthAllocationsQos, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BandwidthAllocationsQos", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BandwidthAllocationsQos
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk deployment_services.BandwidthAllocationsQos
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Customized.IsNull() && !model.Customized.IsUnknown() {
		sdk.Customized = model.Customized.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Customized", "value": *sdk.Customized})
	}

	// Handling Primitives
	if !model.Enabled.IsNull() && !model.Enabled.IsUnknown() {
		sdk.Enabled = model.Enabled.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Enabled", "value": *sdk.Enabled})
	}

	// Handling Primitives
	if !model.GuaranteedRatio.IsNull() && !model.GuaranteedRatio.IsUnknown() {
		val := float32(model.GuaranteedRatio.ValueFloat64())
		sdk.GuaranteedRatio = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "GuaranteedRatio", "value": *sdk.GuaranteedRatio})
	}

	// Handling Primitives
	if !model.Profile.IsNull() && !model.Profile.IsUnknown() {
		sdk.Profile = model.Profile.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Profile", "value": *sdk.Profile})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BandwidthAllocationsQos", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BandwidthAllocationsQos ---
func packBandwidthAllocationsQosFromSdk(ctx context.Context, sdk deployment_services.BandwidthAllocationsQos) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BandwidthAllocationsQos", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BandwidthAllocationsQos
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Customized != nil {
		model.Customized = basetypes.NewBoolValue(*sdk.Customized)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Customized", "value": *sdk.Customized})
	} else {
		model.Customized = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Enabled != nil {
		model.Enabled = basetypes.NewBoolValue(*sdk.Enabled)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Enabled", "value": *sdk.Enabled})
	} else {
		model.Enabled = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.GuaranteedRatio != nil {
		model.GuaranteedRatio = basetypes.NewFloat64Value(float64(*sdk.GuaranteedRatio))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "GuaranteedRatio", "value": *sdk.GuaranteedRatio})
	} else {
		model.GuaranteedRatio = basetypes.NewFloat64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Profile != nil {
		model.Profile = basetypes.NewStringValue(*sdk.Profile)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Profile", "value": *sdk.Profile})
	} else {
		model.Profile = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BandwidthAllocationsQos{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BandwidthAllocationsQos", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BandwidthAllocationsQos ---
func unpackBandwidthAllocationsQosListToSdk(ctx context.Context, list types.List) ([]deployment_services.BandwidthAllocationsQos, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BandwidthAllocationsQos")
	diags := diag.Diagnostics{}
	var data []models.BandwidthAllocationsQos
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]deployment_services.BandwidthAllocationsQos, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BandwidthAllocationsQos{}.AttrTypes(), &item)
		unpacked, d := unpackBandwidthAllocationsQosToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BandwidthAllocationsQos", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BandwidthAllocationsQos ---
func packBandwidthAllocationsQosListFromSdk(ctx context.Context, sdks []deployment_services.BandwidthAllocationsQos) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BandwidthAllocationsQos")
	diags := diag.Diagnostics{}
	var data []models.BandwidthAllocationsQos

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BandwidthAllocationsQos
		obj, d := packBandwidthAllocationsQosFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BandwidthAllocationsQos{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BandwidthAllocationsQos", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BandwidthAllocationsQos{}.AttrType(), data)
}
