package provider

/*

import (
	"context"
	"fmt"
	"reflect"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/objects"
	"github.com/paloaltonetworks/scm-go/generated/objects"
)










// --- Unpacker for QuarantinedDevices ---
func unpackQuarantinedDevicesToSdk(ctx context.Context, obj types.Object) (*objects.QuarantinedDevices, diag.Diagnostics) {
    tflog.Debug(ctx, "Entering unpack helper for models.QuarantinedDevices", map[string]interface{}{"tf_object": obj})
    diags := diag.Diagnostics{}
    var model models.QuarantinedDevices
    diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
    if diags.HasError() {
        tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
        return nil, diags
    }
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

    var sdk objects.QuarantinedDevices
    var d diag.Diagnostics

    // Handling Primitives
    if !model.HostId.IsNull() && !model.HostId.IsUnknown() {
        sdk.HostId = model.HostId.ValueString()
        tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "HostId", "value": sdk.HostId})
    }

    // Handling Primitives
    if !model.SerialNumber.IsNull() && !model.SerialNumber.IsUnknown() {
        sdk.SerialNumber = model.SerialNumber.ValueStringPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "SerialNumber", "value": *sdk.SerialNumber})
    }

diags.Append(d...)

    tflog.Debug(ctx, "Exiting unpack helper for models.QuarantinedDevices", map[string]interface{}{"has_errors": diags.HasError()})
    return &sdk, diags

}

// --- Packer for QuarantinedDevices ---
func packQuarantinedDevicesFromSdk(ctx context.Context, sdk objects.QuarantinedDevices) (types.Object, diag.Diagnostics) {
    tflog.Debug(ctx, "Entering pack helper for models.QuarantinedDevices", map[string]interface{}{"sdk_struct": sdk})
    diags := diag.Diagnostics{}
    var model models.QuarantinedDevices
    var d diag.Diagnostics
    // Handling Primitives
    // Standard primitive packing
    model.HostId = basetypes.NewStringValue(sdk.HostId)
    tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "HostId", "value": sdk.HostId})
    // Handling Primitives
    // Standard primitive packing
    if sdk.SerialNumber != nil {
        model.SerialNumber = basetypes.NewStringValue(*sdk.SerialNumber)
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "SerialNumber", "value": *sdk.SerialNumber})
    } else {
        model.SerialNumber = basetypes.NewStringNull()
    }
diags.Append(d...)

    obj, d := types.ObjectValueFrom(ctx, models.QuarantinedDevices{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
    diags.Append(d...)
    tflog.Debug(ctx, "Exiting pack helper for models.QuarantinedDevices", map[string]interface{}{"has_errors": diags.HasError()})
    return obj, diags

}

// --- List Unpacker for QuarantinedDevices ---
func unpackQuarantinedDevicesListToSdk(ctx context.Context, list types.List) ([]objects.QuarantinedDevices, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.QuarantinedDevices")
	diags := diag.Diagnostics{}
	var data []models.QuarantinedDevices
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.QuarantinedDevices, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.QuarantinedDevices{}.AttrTypes(), &item)
		unpacked, d := unpackQuarantinedDevicesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.QuarantinedDevices", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for QuarantinedDevices ---
func packQuarantinedDevicesListFromSdk(ctx context.Context, sdks []objects.QuarantinedDevices) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.QuarantinedDevices")
	diags := diag.Diagnostics{}
	var data []models.QuarantinedDevices

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.QuarantinedDevices
		obj, d := packQuarantinedDevicesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.QuarantinedDevices{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.QuarantinedDevices", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.QuarantinedDevices{}.AttrType(), data)
}


*/
