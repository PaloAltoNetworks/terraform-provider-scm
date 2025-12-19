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

// --- Unpacker for SessionTimeouts ---
func unpackSessionTimeoutsToSdk(ctx context.Context, obj types.Object) (*device_settings.SessionTimeouts, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.SessionTimeouts", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.SessionTimeouts
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk device_settings.SessionTimeouts
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
	if !model.SessionTimeouts.IsNull() && !model.SessionTimeouts.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field SessionTimeouts")
		unpacked, d := unpackSessionTimeoutsSessionTimeoutsToSdk(ctx, model.SessionTimeouts)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "SessionTimeouts"})
		}
		if unpacked != nil {
			sdk.SessionTimeouts = unpacked
		}
	}

	// Handling Primitives
	if !model.Snippet.IsNull() && !model.Snippet.IsUnknown() {
		sdk.Snippet = model.Snippet.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.SessionTimeouts", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for SessionTimeouts ---
func packSessionTimeoutsFromSdk(ctx context.Context, sdk device_settings.SessionTimeouts) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.SessionTimeouts", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.SessionTimeouts
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
	if sdk.SessionTimeouts != nil {
		tflog.Debug(ctx, "Packing nested object for field SessionTimeouts")
		packed, d := packSessionTimeoutsSessionTimeoutsFromSdk(ctx, *sdk.SessionTimeouts)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "SessionTimeouts"})
		}
		model.SessionTimeouts = packed
	} else {
		model.SessionTimeouts = basetypes.NewObjectNull(models.SessionTimeoutsSessionTimeouts{}.AttrTypes())
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

	obj, d := types.ObjectValueFrom(ctx, models.SessionTimeouts{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.SessionTimeouts", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for SessionTimeouts ---
func unpackSessionTimeoutsListToSdk(ctx context.Context, list types.List) ([]device_settings.SessionTimeouts, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.SessionTimeouts")
	diags := diag.Diagnostics{}
	var data []models.SessionTimeouts
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]device_settings.SessionTimeouts, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.SessionTimeouts{}.AttrTypes(), &item)
		unpacked, d := unpackSessionTimeoutsToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.SessionTimeouts", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for SessionTimeouts ---
func packSessionTimeoutsListFromSdk(ctx context.Context, sdks []device_settings.SessionTimeouts) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.SessionTimeouts")
	diags := diag.Diagnostics{}
	var data []models.SessionTimeouts

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.SessionTimeouts
		obj, d := packSessionTimeoutsFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.SessionTimeouts{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.SessionTimeouts", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.SessionTimeouts{}.AttrType(), data)
}

// --- Unpacker for SessionTimeoutsSessionTimeouts ---
func unpackSessionTimeoutsSessionTimeoutsToSdk(ctx context.Context, obj types.Object) (*device_settings.SessionTimeoutsSessionTimeouts, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.SessionTimeoutsSessionTimeouts", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.SessionTimeoutsSessionTimeouts
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk device_settings.SessionTimeoutsSessionTimeouts
	var d diag.Diagnostics
	// Handling Primitives
	if !model.TimeoutCaptivePortal.IsNull() && !model.TimeoutCaptivePortal.IsUnknown() {
		val := int32(model.TimeoutCaptivePortal.ValueInt64())
		sdk.TimeoutCaptivePortal = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "TimeoutCaptivePortal", "value": *sdk.TimeoutCaptivePortal})
	}

	// Handling Primitives
	if !model.TimeoutDefault.IsNull() && !model.TimeoutDefault.IsUnknown() {
		val := int32(model.TimeoutDefault.ValueInt64())
		sdk.TimeoutDefault = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "TimeoutDefault", "value": *sdk.TimeoutDefault})
	}

	// Handling Primitives
	if !model.TimeoutDiscardDefault.IsNull() && !model.TimeoutDiscardDefault.IsUnknown() {
		val := int32(model.TimeoutDiscardDefault.ValueInt64())
		sdk.TimeoutDiscardDefault = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "TimeoutDiscardDefault", "value": *sdk.TimeoutDiscardDefault})
	}

	// Handling Primitives
	if !model.TimeoutDiscardTcp.IsNull() && !model.TimeoutDiscardTcp.IsUnknown() {
		val := int32(model.TimeoutDiscardTcp.ValueInt64())
		sdk.TimeoutDiscardTcp = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "TimeoutDiscardTcp", "value": *sdk.TimeoutDiscardTcp})
	}

	// Handling Primitives
	if !model.TimeoutDiscardUdp.IsNull() && !model.TimeoutDiscardUdp.IsUnknown() {
		val := int32(model.TimeoutDiscardUdp.ValueInt64())
		sdk.TimeoutDiscardUdp = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "TimeoutDiscardUdp", "value": *sdk.TimeoutDiscardUdp})
	}

	// Handling Primitives
	if !model.TimeoutIcmp.IsNull() && !model.TimeoutIcmp.IsUnknown() {
		val := int32(model.TimeoutIcmp.ValueInt64())
		sdk.TimeoutIcmp = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "TimeoutIcmp", "value": *sdk.TimeoutIcmp})
	}

	// Handling Primitives
	if !model.TimeoutScan.IsNull() && !model.TimeoutScan.IsUnknown() {
		val := int32(model.TimeoutScan.ValueInt64())
		sdk.TimeoutScan = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "TimeoutScan", "value": *sdk.TimeoutScan})
	}

	// Handling Primitives
	if !model.TimeoutTcp.IsNull() && !model.TimeoutTcp.IsUnknown() {
		val := int32(model.TimeoutTcp.ValueInt64())
		sdk.TimeoutTcp = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "TimeoutTcp", "value": *sdk.TimeoutTcp})
	}

	// Handling Primitives
	if !model.TimeoutTcpHalfClosed.IsNull() && !model.TimeoutTcpHalfClosed.IsUnknown() {
		val := int32(model.TimeoutTcpHalfClosed.ValueInt64())
		sdk.TimeoutTcpHalfClosed = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "TimeoutTcpHalfClosed", "value": *sdk.TimeoutTcpHalfClosed})
	}

	// Handling Primitives
	if !model.TimeoutTcpTimeWait.IsNull() && !model.TimeoutTcpTimeWait.IsUnknown() {
		val := int32(model.TimeoutTcpTimeWait.ValueInt64())
		sdk.TimeoutTcpTimeWait = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "TimeoutTcpTimeWait", "value": *sdk.TimeoutTcpTimeWait})
	}

	// Handling Primitives
	if !model.TimeoutTcpUnverifiedRst.IsNull() && !model.TimeoutTcpUnverifiedRst.IsUnknown() {
		val := int32(model.TimeoutTcpUnverifiedRst.ValueInt64())
		sdk.TimeoutTcpUnverifiedRst = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "TimeoutTcpUnverifiedRst", "value": *sdk.TimeoutTcpUnverifiedRst})
	}

	// Handling Primitives
	if !model.TimeoutTcphandshake.IsNull() && !model.TimeoutTcphandshake.IsUnknown() {
		val := int32(model.TimeoutTcphandshake.ValueInt64())
		sdk.TimeoutTcphandshake = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "TimeoutTcphandshake", "value": *sdk.TimeoutTcphandshake})
	}

	// Handling Primitives
	if !model.TimeoutTcpinit.IsNull() && !model.TimeoutTcpinit.IsUnknown() {
		val := int32(model.TimeoutTcpinit.ValueInt64())
		sdk.TimeoutTcpinit = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "TimeoutTcpinit", "value": *sdk.TimeoutTcpinit})
	}

	// Handling Primitives
	if !model.TimeoutUdp.IsNull() && !model.TimeoutUdp.IsUnknown() {
		val := int32(model.TimeoutUdp.ValueInt64())
		sdk.TimeoutUdp = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "TimeoutUdp", "value": *sdk.TimeoutUdp})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.SessionTimeoutsSessionTimeouts", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for SessionTimeoutsSessionTimeouts ---
func packSessionTimeoutsSessionTimeoutsFromSdk(ctx context.Context, sdk device_settings.SessionTimeoutsSessionTimeouts) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.SessionTimeoutsSessionTimeouts", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.SessionTimeoutsSessionTimeouts
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.TimeoutCaptivePortal != nil {
		model.TimeoutCaptivePortal = basetypes.NewInt64Value(int64(*sdk.TimeoutCaptivePortal))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "TimeoutCaptivePortal", "value": *sdk.TimeoutCaptivePortal})
	} else {
		model.TimeoutCaptivePortal = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.TimeoutDefault != nil {
		model.TimeoutDefault = basetypes.NewInt64Value(int64(*sdk.TimeoutDefault))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "TimeoutDefault", "value": *sdk.TimeoutDefault})
	} else {
		model.TimeoutDefault = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.TimeoutDiscardDefault != nil {
		model.TimeoutDiscardDefault = basetypes.NewInt64Value(int64(*sdk.TimeoutDiscardDefault))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "TimeoutDiscardDefault", "value": *sdk.TimeoutDiscardDefault})
	} else {
		model.TimeoutDiscardDefault = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.TimeoutDiscardTcp != nil {
		model.TimeoutDiscardTcp = basetypes.NewInt64Value(int64(*sdk.TimeoutDiscardTcp))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "TimeoutDiscardTcp", "value": *sdk.TimeoutDiscardTcp})
	} else {
		model.TimeoutDiscardTcp = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.TimeoutDiscardUdp != nil {
		model.TimeoutDiscardUdp = basetypes.NewInt64Value(int64(*sdk.TimeoutDiscardUdp))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "TimeoutDiscardUdp", "value": *sdk.TimeoutDiscardUdp})
	} else {
		model.TimeoutDiscardUdp = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.TimeoutIcmp != nil {
		model.TimeoutIcmp = basetypes.NewInt64Value(int64(*sdk.TimeoutIcmp))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "TimeoutIcmp", "value": *sdk.TimeoutIcmp})
	} else {
		model.TimeoutIcmp = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.TimeoutScan != nil {
		model.TimeoutScan = basetypes.NewInt64Value(int64(*sdk.TimeoutScan))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "TimeoutScan", "value": *sdk.TimeoutScan})
	} else {
		model.TimeoutScan = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.TimeoutTcp != nil {
		model.TimeoutTcp = basetypes.NewInt64Value(int64(*sdk.TimeoutTcp))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "TimeoutTcp", "value": *sdk.TimeoutTcp})
	} else {
		model.TimeoutTcp = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.TimeoutTcpHalfClosed != nil {
		model.TimeoutTcpHalfClosed = basetypes.NewInt64Value(int64(*sdk.TimeoutTcpHalfClosed))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "TimeoutTcpHalfClosed", "value": *sdk.TimeoutTcpHalfClosed})
	} else {
		model.TimeoutTcpHalfClosed = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.TimeoutTcpTimeWait != nil {
		model.TimeoutTcpTimeWait = basetypes.NewInt64Value(int64(*sdk.TimeoutTcpTimeWait))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "TimeoutTcpTimeWait", "value": *sdk.TimeoutTcpTimeWait})
	} else {
		model.TimeoutTcpTimeWait = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.TimeoutTcpUnverifiedRst != nil {
		model.TimeoutTcpUnverifiedRst = basetypes.NewInt64Value(int64(*sdk.TimeoutTcpUnverifiedRst))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "TimeoutTcpUnverifiedRst", "value": *sdk.TimeoutTcpUnverifiedRst})
	} else {
		model.TimeoutTcpUnverifiedRst = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.TimeoutTcphandshake != nil {
		model.TimeoutTcphandshake = basetypes.NewInt64Value(int64(*sdk.TimeoutTcphandshake))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "TimeoutTcphandshake", "value": *sdk.TimeoutTcphandshake})
	} else {
		model.TimeoutTcphandshake = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.TimeoutTcpinit != nil {
		model.TimeoutTcpinit = basetypes.NewInt64Value(int64(*sdk.TimeoutTcpinit))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "TimeoutTcpinit", "value": *sdk.TimeoutTcpinit})
	} else {
		model.TimeoutTcpinit = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.TimeoutUdp != nil {
		model.TimeoutUdp = basetypes.NewInt64Value(int64(*sdk.TimeoutUdp))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "TimeoutUdp", "value": *sdk.TimeoutUdp})
	} else {
		model.TimeoutUdp = basetypes.NewInt64Null()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.SessionTimeoutsSessionTimeouts{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.SessionTimeoutsSessionTimeouts", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for SessionTimeoutsSessionTimeouts ---
func unpackSessionTimeoutsSessionTimeoutsListToSdk(ctx context.Context, list types.List) ([]device_settings.SessionTimeoutsSessionTimeouts, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.SessionTimeoutsSessionTimeouts")
	diags := diag.Diagnostics{}
	var data []models.SessionTimeoutsSessionTimeouts
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]device_settings.SessionTimeoutsSessionTimeouts, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.SessionTimeoutsSessionTimeouts{}.AttrTypes(), &item)
		unpacked, d := unpackSessionTimeoutsSessionTimeoutsToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.SessionTimeoutsSessionTimeouts", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for SessionTimeoutsSessionTimeouts ---
func packSessionTimeoutsSessionTimeoutsListFromSdk(ctx context.Context, sdks []device_settings.SessionTimeoutsSessionTimeouts) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.SessionTimeoutsSessionTimeouts")
	diags := diag.Diagnostics{}
	var data []models.SessionTimeoutsSessionTimeouts

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.SessionTimeoutsSessionTimeouts
		obj, d := packSessionTimeoutsSessionTimeoutsFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.SessionTimeoutsSessionTimeouts{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.SessionTimeoutsSessionTimeouts", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.SessionTimeoutsSessionTimeouts{}.AttrType(), data)
}
