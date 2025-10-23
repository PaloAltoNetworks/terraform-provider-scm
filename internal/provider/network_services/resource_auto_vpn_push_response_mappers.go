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

	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/network_services"
	"github.com/paloaltonetworks/scm-go/generated/network_services"
)










// --- Unpacker for AutoVpnPushResponse ---
func unpackAutoVpnPushResponseToSdk(ctx context.Context, obj types.Object) (*network_services.AutoVpnPushResponse, diag.Diagnostics) {
    tflog.Debug(ctx, "Entering unpack helper for models.AutoVpnPushResponse", map[string]interface{}{"tf_object": obj})
    diags := diag.Diagnostics{}
    var model models.AutoVpnPushResponse
    diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
    if diags.HasError() {
        tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
        return nil, diags
    }
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

    var sdk network_services.AutoVpnPushResponse
    var d diag.Diagnostics

    // Handling Primitives
    if !model.Job.IsNull() && !model.Job.IsUnknown() {
        sdk.Job = model.Job.ValueStringPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Job", "value": *sdk.Job})
    }

    // Handling Primitives
    if !model.Message.IsNull() && !model.Message.IsUnknown() {
        sdk.Message = model.Message.ValueStringPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Message", "value": *sdk.Message})
    }

    // Handling Primitives
    if !model.Success.IsNull() && !model.Success.IsUnknown() {
        sdk.Success = model.Success.ValueBoolPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Success", "value": *sdk.Success})
    }

diags.Append(d...)

    tflog.Debug(ctx, "Exiting unpack helper for models.AutoVpnPushResponse", map[string]interface{}{"has_errors": diags.HasError()})
    return &sdk, diags

}

// --- Packer for AutoVpnPushResponse ---
func packAutoVpnPushResponseFromSdk(ctx context.Context, sdk network_services.AutoVpnPushResponse) (types.Object, diag.Diagnostics) {
    tflog.Debug(ctx, "Entering pack helper for models.AutoVpnPushResponse", map[string]interface{}{"sdk_struct": sdk})
    diags := diag.Diagnostics{}
    var model models.AutoVpnPushResponse
    var d diag.Diagnostics
    // Handling Primitives
    // Standard primitive packing
    if sdk.Job != nil {
        model.Job = basetypes.NewStringValue(*sdk.Job)
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Job", "value": *sdk.Job})
    } else {
        model.Job = basetypes.NewStringNull()
    }
    // Handling Primitives
    // Standard primitive packing
    if sdk.Message != nil {
        model.Message = basetypes.NewStringValue(*sdk.Message)
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Message", "value": *sdk.Message})
    } else {
        model.Message = basetypes.NewStringNull()
    }
    // Handling Primitives
    // Standard primitive packing
    if sdk.Success != nil {
        model.Success = basetypes.NewBoolValue(*sdk.Success)
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Success", "value": *sdk.Success})
    } else {
        model.Success = basetypes.NewBoolNull()
    }
diags.Append(d...)

    obj, d := types.ObjectValueFrom(ctx, models.AutoVpnPushResponse{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
    diags.Append(d...)
    tflog.Debug(ctx, "Exiting pack helper for models.AutoVpnPushResponse", map[string]interface{}{"has_errors": diags.HasError()})
    return obj, diags

}

// --- List Unpacker for AutoVpnPushResponse ---
func unpackAutoVpnPushResponseListToSdk(ctx context.Context, list types.List) ([]network_services.AutoVpnPushResponse, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.AutoVpnPushResponse")
	diags := diag.Diagnostics{}
	var data []models.AutoVpnPushResponse
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.AutoVpnPushResponse, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.AutoVpnPushResponse{}.AttrTypes(), &item)
		unpacked, d := unpackAutoVpnPushResponseToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.AutoVpnPushResponse", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for AutoVpnPushResponse ---
func packAutoVpnPushResponseListFromSdk(ctx context.Context, sdks []network_services.AutoVpnPushResponse) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.AutoVpnPushResponse")
	diags := diag.Diagnostics{}
	var data []models.AutoVpnPushResponse

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.AutoVpnPushResponse
		obj, d := packAutoVpnPushResponseFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.AutoVpnPushResponse{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.AutoVpnPushResponse", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.AutoVpnPushResponse{}.AttrType(), data)
}

// --- Unpacker for AutoVpnPushConfig ---
func unpackAutoVpnPushConfigToSdk(ctx context.Context, obj types.Object) (*network_services.AutoVpnPushConfig, diag.Diagnostics) {
    tflog.Debug(ctx, "Entering unpack helper for models.AutoVpnPushConfig", map[string]interface{}{"tf_object": obj})
    diags := diag.Diagnostics{}
    var model models.AutoVpnPushConfig
    diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
    if diags.HasError() {
        tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
        return nil, diags
    }
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

    var sdk network_services.AutoVpnPushConfig
    var d diag.Diagnostics

    // Handling Lists
    if !model.AutoVpnDevices.IsNull() && !model.AutoVpnDevices.IsUnknown() {
        tflog.Debug(ctx, "Unpacking list of objects for field AutoVpnDevices")
        unpacked, d := unpackAutoVpnPushConfigAutoVpnDevicesInnerListToSdk(ctx, model.AutoVpnDevices)
        diags.Append(d...)
        sdk.AutoVpnDevices = unpacked
    }

diags.Append(d...)

    tflog.Debug(ctx, "Exiting unpack helper for models.AutoVpnPushConfig", map[string]interface{}{"has_errors": diags.HasError()})
    return &sdk, diags

}

// --- Packer for AutoVpnPushConfig ---
func packAutoVpnPushConfigFromSdk(ctx context.Context, sdk network_services.AutoVpnPushConfig) (types.Object, diag.Diagnostics) {
    tflog.Debug(ctx, "Entering pack helper for models.AutoVpnPushConfig", map[string]interface{}{"sdk_struct": sdk})
    diags := diag.Diagnostics{}
    var model models.AutoVpnPushConfig
    var d diag.Diagnostics
    // Handling Lists
    if sdk.AutoVpnDevices != nil {
        tflog.Debug(ctx, "Packing list of objects for field AutoVpnDevices")
        packed, d := packAutoVpnPushConfigAutoVpnDevicesInnerListFromSdk(ctx, sdk.AutoVpnDevices)
        diags.Append(d...)
        model.AutoVpnDevices = packed
    } else {
        model.AutoVpnDevices = basetypes.NewListNull(models.AutoVpnPushConfigAutoVpnDevicesInner{}.AttrType())
    }
diags.Append(d...)

    obj, d := types.ObjectValueFrom(ctx, models.AutoVpnPushConfig{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
    diags.Append(d...)
    tflog.Debug(ctx, "Exiting pack helper for models.AutoVpnPushConfig", map[string]interface{}{"has_errors": diags.HasError()})
    return obj, diags

}

// --- List Unpacker for AutoVpnPushConfig ---
func unpackAutoVpnPushConfigListToSdk(ctx context.Context, list types.List) ([]network_services.AutoVpnPushConfig, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.AutoVpnPushConfig")
	diags := diag.Diagnostics{}
	var data []models.AutoVpnPushConfig
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.AutoVpnPushConfig, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.AutoVpnPushConfig{}.AttrTypes(), &item)
		unpacked, d := unpackAutoVpnPushConfigToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.AutoVpnPushConfig", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for AutoVpnPushConfig ---
func packAutoVpnPushConfigListFromSdk(ctx context.Context, sdks []network_services.AutoVpnPushConfig) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.AutoVpnPushConfig")
	diags := diag.Diagnostics{}
	var data []models.AutoVpnPushConfig

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.AutoVpnPushConfig
		obj, d := packAutoVpnPushConfigFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.AutoVpnPushConfig{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.AutoVpnPushConfig", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.AutoVpnPushConfig{}.AttrType(), data)
}

// --- Unpacker for AutoVpnPushConfigAutoVpnDevicesInner ---
func unpackAutoVpnPushConfigAutoVpnDevicesInnerToSdk(ctx context.Context, obj types.Object) (*network_services.AutoVpnPushConfigAutoVpnDevicesInner, diag.Diagnostics) {
    tflog.Debug(ctx, "Entering unpack helper for models.AutoVpnPushConfigAutoVpnDevicesInner", map[string]interface{}{"tf_object": obj})
    diags := diag.Diagnostics{}
    var model models.AutoVpnPushConfigAutoVpnDevicesInner
    diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
    if diags.HasError() {
        tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
        return nil, diags
    }
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

    var sdk network_services.AutoVpnPushConfigAutoVpnDevicesInner
    var d diag.Diagnostics
    // Handling Primitives
    if !model.Name.IsNull() && !model.Name.IsUnknown() {
        sdk.Name = model.Name.ValueStringPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
    }

    // Handling Primitives
    if !model.RefreshPsk.IsNull() && !model.RefreshPsk.IsUnknown() {
        sdk.RefreshPsk = model.RefreshPsk.ValueBoolPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "RefreshPsk", "value": *sdk.RefreshPsk})
    }

diags.Append(d...)

    tflog.Debug(ctx, "Exiting unpack helper for models.AutoVpnPushConfigAutoVpnDevicesInner", map[string]interface{}{"has_errors": diags.HasError()})
    return &sdk, diags

}

// --- Packer for AutoVpnPushConfigAutoVpnDevicesInner ---
func packAutoVpnPushConfigAutoVpnDevicesInnerFromSdk(ctx context.Context, sdk network_services.AutoVpnPushConfigAutoVpnDevicesInner) (types.Object, diag.Diagnostics) {
    tflog.Debug(ctx, "Entering pack helper for models.AutoVpnPushConfigAutoVpnDevicesInner", map[string]interface{}{"sdk_struct": sdk})
    diags := diag.Diagnostics{}
    var model models.AutoVpnPushConfigAutoVpnDevicesInner
    var d diag.Diagnostics
    // Handling Primitives
    // Standard primitive packing
    if sdk.Name != nil {
        model.Name = basetypes.NewStringValue(*sdk.Name)
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
    } else {
        model.Name = basetypes.NewStringNull()
    }
    // Handling Primitives
    // Standard primitive packing
    if sdk.RefreshPsk != nil {
        model.RefreshPsk = basetypes.NewBoolValue(*sdk.RefreshPsk)
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "RefreshPsk", "value": *sdk.RefreshPsk})
    } else {
        model.RefreshPsk = basetypes.NewBoolNull()
    }
diags.Append(d...)

    obj, d := types.ObjectValueFrom(ctx, models.AutoVpnPushConfigAutoVpnDevicesInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
    diags.Append(d...)
    tflog.Debug(ctx, "Exiting pack helper for models.AutoVpnPushConfigAutoVpnDevicesInner", map[string]interface{}{"has_errors": diags.HasError()})
    return obj, diags

}

// --- List Unpacker for AutoVpnPushConfigAutoVpnDevicesInner ---
func unpackAutoVpnPushConfigAutoVpnDevicesInnerListToSdk(ctx context.Context, list types.List) ([]network_services.AutoVpnPushConfigAutoVpnDevicesInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.AutoVpnPushConfigAutoVpnDevicesInner")
	diags := diag.Diagnostics{}
	var data []models.AutoVpnPushConfigAutoVpnDevicesInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.AutoVpnPushConfigAutoVpnDevicesInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.AutoVpnPushConfigAutoVpnDevicesInner{}.AttrTypes(), &item)
		unpacked, d := unpackAutoVpnPushConfigAutoVpnDevicesInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.AutoVpnPushConfigAutoVpnDevicesInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for AutoVpnPushConfigAutoVpnDevicesInner ---
func packAutoVpnPushConfigAutoVpnDevicesInnerListFromSdk(ctx context.Context, sdks []network_services.AutoVpnPushConfigAutoVpnDevicesInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.AutoVpnPushConfigAutoVpnDevicesInner")
	diags := diag.Diagnostics{}
	var data []models.AutoVpnPushConfigAutoVpnDevicesInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.AutoVpnPushConfigAutoVpnDevicesInner
		obj, d := packAutoVpnPushConfigAutoVpnDevicesInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.AutoVpnPushConfigAutoVpnDevicesInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.AutoVpnPushConfigAutoVpnDevicesInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.AutoVpnPushConfigAutoVpnDevicesInner{}.AttrType(), data)
}


*/
