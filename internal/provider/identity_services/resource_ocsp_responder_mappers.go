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

	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/identity_services"
	"github.com/paloaltonetworks/scm-go/generated/identity_services"
)










// --- Unpacker for OcspResponders ---
func unpackOcspRespondersToSdk(ctx context.Context, obj types.Object) (*identity_services.OcspResponders, diag.Diagnostics) {
    tflog.Debug(ctx, "Entering unpack helper for models.OcspResponders", map[string]interface{}{"tf_object": obj})
    diags := diag.Diagnostics{}
    var model models.OcspResponders
    diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
    if diags.HasError() {
        tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
        return nil, diags
    }
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

    var sdk identity_services.OcspResponders
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
    if !model.HostName.IsNull() && !model.HostName.IsUnknown() {
        sdk.HostName = model.HostName.ValueString()
        tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "HostName", "value": sdk.HostName})
    }

    // Handling Primitives
    if !model.Id.IsNull() && !model.Id.IsUnknown() {
        sdk.Id = model.Id.ValueString()
        tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Id", "value": sdk.Id})
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

    tflog.Debug(ctx, "Exiting unpack helper for models.OcspResponders", map[string]interface{}{"has_errors": diags.HasError()})
    return &sdk, diags

}

// --- Packer for OcspResponders ---
func packOcspRespondersFromSdk(ctx context.Context, sdk identity_services.OcspResponders) (types.Object, diag.Diagnostics) {
    tflog.Debug(ctx, "Entering pack helper for models.OcspResponders", map[string]interface{}{"sdk_struct": sdk})
    diags := diag.Diagnostics{}
    var model models.OcspResponders
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
    model.HostName = basetypes.NewStringValue(sdk.HostName)
    tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "HostName", "value": sdk.HostName})
    // Handling Primitives
    // Standard primitive packing
    model.Id = basetypes.NewStringValue(sdk.Id)
    tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Id", "value": sdk.Id})
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

    obj, d := types.ObjectValueFrom(ctx, models.OcspResponders{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
    diags.Append(d...)
    tflog.Debug(ctx, "Exiting pack helper for models.OcspResponders", map[string]interface{}{"has_errors": diags.HasError()})
    return obj, diags

}

// --- List Unpacker for OcspResponders ---
func unpackOcspRespondersListToSdk(ctx context.Context, list types.List) ([]identity_services.OcspResponders, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.OcspResponders")
	diags := diag.Diagnostics{}
	var data []models.OcspResponders
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]identity_services.OcspResponders, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.OcspResponders{}.AttrTypes(), &item)
		unpacked, d := unpackOcspRespondersToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.OcspResponders", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for OcspResponders ---
func packOcspRespondersListFromSdk(ctx context.Context, sdks []identity_services.OcspResponders) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.OcspResponders")
	diags := diag.Diagnostics{}
	var data []models.OcspResponders

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.OcspResponders
		obj, d := packOcspRespondersFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.OcspResponders{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.OcspResponders", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.OcspResponders{}.AttrType(), data)
}


*/
