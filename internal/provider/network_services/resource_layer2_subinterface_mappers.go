package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/network_services"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/network_services"
)

// --- Unpacker for Layer2Subinterfaces ---
func unpackLayer2SubinterfacesToSdk(ctx context.Context, obj types.Object) (*network_services.Layer2Subinterfaces, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.Layer2Subinterfaces", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.Layer2Subinterfaces
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.Layer2Subinterfaces
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Comment.IsNull() && !model.Comment.IsUnknown() {
		sdk.Comment = model.Comment.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Comment", "value": *sdk.Comment})
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

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	// Handling Primitives
	if !model.ParentInterface.IsNull() && !model.ParentInterface.IsUnknown() {
		sdk.ParentInterface = model.ParentInterface.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "ParentInterface", "value": *sdk.ParentInterface})
	}

	// Handling Primitives
	if !model.Snippet.IsNull() && !model.Snippet.IsUnknown() {
		sdk.Snippet = model.Snippet.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	}

	// Handling Primitives
	if !model.VlanTag.IsNull() && !model.VlanTag.IsUnknown() {
		val := float32(model.VlanTag.ValueFloat64())
		sdk.VlanTag = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "VlanTag", "value": *sdk.VlanTag})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.Layer2Subinterfaces", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for Layer2Subinterfaces ---
func packLayer2SubinterfacesFromSdk(ctx context.Context, sdk network_services.Layer2Subinterfaces) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.Layer2Subinterfaces", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.Layer2Subinterfaces
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Comment != nil {
		model.Comment = basetypes.NewStringValue(*sdk.Comment)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Comment", "value": *sdk.Comment})
	} else {
		model.Comment = basetypes.NewStringNull()
	}
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
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	// Handling Primitives
	// Standard primitive packing
	if sdk.ParentInterface != nil {
		model.ParentInterface = basetypes.NewStringValue(*sdk.ParentInterface)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "ParentInterface", "value": *sdk.ParentInterface})
	} else {
		model.ParentInterface = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Snippet != nil {
		model.Snippet = basetypes.NewStringValue(*sdk.Snippet)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	} else {
		model.Snippet = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.VlanTag != nil {
		model.VlanTag = basetypes.NewFloat64Value(float64(*sdk.VlanTag))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "VlanTag", "value": *sdk.VlanTag})
	} else {
		model.VlanTag = basetypes.NewFloat64Null()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.Layer2Subinterfaces{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.Layer2Subinterfaces", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for Layer2Subinterfaces ---
func unpackLayer2SubinterfacesListToSdk(ctx context.Context, list types.List) ([]network_services.Layer2Subinterfaces, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.Layer2Subinterfaces")
	diags := diag.Diagnostics{}
	var data []models.Layer2Subinterfaces
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.Layer2Subinterfaces, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.Layer2Subinterfaces{}.AttrTypes(), &item)
		unpacked, d := unpackLayer2SubinterfacesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.Layer2Subinterfaces", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for Layer2Subinterfaces ---
func packLayer2SubinterfacesListFromSdk(ctx context.Context, sdks []network_services.Layer2Subinterfaces) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.Layer2Subinterfaces")
	diags := diag.Diagnostics{}
	var data []models.Layer2Subinterfaces

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.Layer2Subinterfaces
		obj, d := packLayer2SubinterfacesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.Layer2Subinterfaces{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.Layer2Subinterfaces", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.Layer2Subinterfaces{}.AttrType(), data)
}
