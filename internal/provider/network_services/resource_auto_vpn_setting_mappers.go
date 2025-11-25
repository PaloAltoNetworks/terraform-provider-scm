package provider

import (
	"context"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/network_services"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/network_services"
)

// --- Unpacker for AutoVpnSettings ---
func unpackAutoVpnSettingsToSdk(ctx context.Context, obj types.Object) (*network_services.AutoVpnSettings, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.AutoVpnSettings", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.AutoVpnSettings
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.AutoVpnSettings
	var d diag.Diagnostics

	// Handling Objects
	if !model.AsRange.IsNull() && !model.AsRange.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field AsRange")
		unpacked, d := unpackAutoVpnSettingsAsRangeToSdk(ctx, model.AsRange)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "AsRange"})
		}
		if unpacked != nil {
			sdk.AsRange = *unpacked
		}
	}

	// Handling Primitives
	if !model.EnableMeshBetweenHubs.IsNull() && !model.EnableMeshBetweenHubs.IsUnknown() {
		sdk.EnableMeshBetweenHubs = model.EnableMeshBetweenHubs.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "EnableMeshBetweenHubs", "value": *sdk.EnableMeshBetweenHubs})
	}

	// Handling Lists
	if !model.VpnAddressPool.IsNull() && !model.VpnAddressPool.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field VpnAddressPool")
		diags.Append(model.VpnAddressPool.ElementsAs(ctx, &sdk.VpnAddressPool, false)...)
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.AutoVpnSettings", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for AutoVpnSettings ---
func packAutoVpnSettingsFromSdk(ctx context.Context, sdk network_services.AutoVpnSettings) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.AutoVpnSettings", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.AutoVpnSettings
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	// Logic for non-pointer / value-type nested objects
	if !reflect.ValueOf(sdk.AsRange).IsZero() {
		tflog.Debug(ctx, "Packing nested object for field AsRange")
		packed, d := packAutoVpnSettingsAsRangeFromSdk(ctx, sdk.AsRange)
		diags.Append(d...)
		model.AsRange = packed
	} else {
		model.AsRange = basetypes.NewObjectNull(models.AutoVpnSettingsAsRange{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.EnableMeshBetweenHubs != nil {
		model.EnableMeshBetweenHubs = basetypes.NewBoolValue(*sdk.EnableMeshBetweenHubs)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "EnableMeshBetweenHubs", "value": *sdk.EnableMeshBetweenHubs})
	} else {
		model.EnableMeshBetweenHubs = basetypes.NewBoolNull()
	}
	// Handling Lists
	if sdk.VpnAddressPool != nil {
		tflog.Debug(ctx, "Packing list of primitives for field VpnAddressPool")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.VpnAddressPool, d = basetypes.NewListValueFrom(ctx, elemType, sdk.VpnAddressPool)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.VpnAddressPool = basetypes.NewListNull(elemType)
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.AutoVpnSettings{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.AutoVpnSettings", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for AutoVpnSettings ---
func unpackAutoVpnSettingsListToSdk(ctx context.Context, list types.List) ([]network_services.AutoVpnSettings, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.AutoVpnSettings")
	diags := diag.Diagnostics{}
	var data []models.AutoVpnSettings
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.AutoVpnSettings, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.AutoVpnSettings{}.AttrTypes(), &item)
		unpacked, d := unpackAutoVpnSettingsToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.AutoVpnSettings", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for AutoVpnSettings ---
func packAutoVpnSettingsListFromSdk(ctx context.Context, sdks []network_services.AutoVpnSettings) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.AutoVpnSettings")
	diags := diag.Diagnostics{}
	var data []models.AutoVpnSettings

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.AutoVpnSettings
		obj, d := packAutoVpnSettingsFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.AutoVpnSettings{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.AutoVpnSettings", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.AutoVpnSettings{}.AttrType(), data)
}

// --- Unpacker for AutoVpnSettingsAsRange ---
func unpackAutoVpnSettingsAsRangeToSdk(ctx context.Context, obj types.Object) (*network_services.AutoVpnSettingsAsRange, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.AutoVpnSettingsAsRange", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.AutoVpnSettingsAsRange
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.AutoVpnSettingsAsRange
	var d diag.Diagnostics
	// Handling Primitives
	if !model.End.IsNull() && !model.End.IsUnknown() {
		val := int32(model.End.ValueInt64())
		sdk.End = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "End", "value": *sdk.End})
	}

	// Handling Primitives
	if !model.Start.IsNull() && !model.Start.IsUnknown() {
		val := int32(model.Start.ValueInt64())
		sdk.Start = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Start", "value": *sdk.Start})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.AutoVpnSettingsAsRange", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for AutoVpnSettingsAsRange ---
func packAutoVpnSettingsAsRangeFromSdk(ctx context.Context, sdk network_services.AutoVpnSettingsAsRange) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.AutoVpnSettingsAsRange", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.AutoVpnSettingsAsRange
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.End != nil {
		model.End = basetypes.NewInt64Value(int64(*sdk.End))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "End", "value": *sdk.End})
	} else {
		model.End = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Start != nil {
		model.Start = basetypes.NewInt64Value(int64(*sdk.Start))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Start", "value": *sdk.Start})
	} else {
		model.Start = basetypes.NewInt64Null()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.AutoVpnSettingsAsRange{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.AutoVpnSettingsAsRange", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for AutoVpnSettingsAsRange ---
func unpackAutoVpnSettingsAsRangeListToSdk(ctx context.Context, list types.List) ([]network_services.AutoVpnSettingsAsRange, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.AutoVpnSettingsAsRange")
	diags := diag.Diagnostics{}
	var data []models.AutoVpnSettingsAsRange
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.AutoVpnSettingsAsRange, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.AutoVpnSettingsAsRange{}.AttrTypes(), &item)
		unpacked, d := unpackAutoVpnSettingsAsRangeToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.AutoVpnSettingsAsRange", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for AutoVpnSettingsAsRange ---
func packAutoVpnSettingsAsRangeListFromSdk(ctx context.Context, sdks []network_services.AutoVpnSettingsAsRange) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.AutoVpnSettingsAsRange")
	diags := diag.Diagnostics{}
	var data []models.AutoVpnSettingsAsRange

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.AutoVpnSettingsAsRange
		obj, d := packAutoVpnSettingsAsRangeFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.AutoVpnSettingsAsRange{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.AutoVpnSettingsAsRange", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.AutoVpnSettingsAsRange{}.AttrType(), data)
}
