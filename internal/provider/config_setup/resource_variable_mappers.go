package provider

import (
	"context"
	"reflect"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/config_setup"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/config_setup"
)

// --- Unpacker for Variables ---
func unpackVariablesToSdk(ctx context.Context, obj types.Object) (*config_setup.Variables, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.Variables", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.Variables
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk config_setup.Variables
	var d diag.Diagnostics

	// Handling Primitives
	if !model.Description.IsNull() && !model.Description.IsUnknown() {
		sdk.Description = model.Description.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Description", "value": *sdk.Description})
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
		sdk.Id = model.Id.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Id", "value": sdk.Id})
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	// Handling Primitives
	if !model.Overridden.IsNull() && !model.Overridden.IsUnknown() {
		sdk.Overridden = model.Overridden.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Overridden", "value": *sdk.Overridden})
	}

	// Handling Primitives
	if !model.Snippet.IsNull() && !model.Snippet.IsUnknown() {
		sdk.Snippet = model.Snippet.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	}

	// Handling Primitives
	if !model.Type.IsNull() && !model.Type.IsUnknown() {
		sdk.Type = model.Type.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Type", "value": sdk.Type})
	}

	// Handling Primitives
	if !model.Value.IsNull() && !model.Value.IsUnknown() {
		sdk.Value = model.Value.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Value", "value": sdk.Value})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.Variables", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for Variables ---
func packVariablesFromSdk(ctx context.Context, sdk config_setup.Variables) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.Variables", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.Variables
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Description != nil {
		model.Description = basetypes.NewStringValue(*sdk.Description)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Description", "value": *sdk.Description})
	} else {
		model.Description = basetypes.NewStringNull()
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
	model.Id = basetypes.NewStringValue(sdk.Id)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Id", "value": sdk.Id})
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	// Handling Primitives
	// Standard primitive packing
	if sdk.Overridden != nil {
		model.Overridden = basetypes.NewBoolValue(*sdk.Overridden)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Overridden", "value": *sdk.Overridden})
	} else {
		model.Overridden = basetypes.NewBoolNull()
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
	model.Type = basetypes.NewStringValue(sdk.Type)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Type", "value": sdk.Type})
	// Handling Primitives
	// Universal packer for interface{} types that are mapped to a StringValue in the model.
	// All underlying primitive types will be converted to their string representation.
	if sdk.Value != nil {
		tflog.Debug(ctx, "Packing interface value for field Value")
		switch v := sdk.Value.(type) {
		case string:
			model.Value = basetypes.NewStringValue(v)
		case int:
			model.Value = basetypes.NewStringValue(strconv.Itoa(v))
		case int64:
			model.Value = basetypes.NewStringValue(strconv.FormatInt(v, 10))
		case bool:
			model.Value = basetypes.NewStringValue(strconv.FormatBool(v))
		case float64:
			model.Value = basetypes.NewStringValue(strconv.FormatFloat(v, 'f', -1, 64))
		default:
			tflog.Warn(ctx, "Unexpected type for interface field", map[string]interface{}{"field": "Value", "type": reflect.TypeOf(v).String()})
			model.Value = basetypes.NewStringNull()
		}
	} else {
		model.Value = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.Variables{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.Variables", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for Variables ---
func unpackVariablesListToSdk(ctx context.Context, list types.List) ([]config_setup.Variables, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.Variables")
	diags := diag.Diagnostics{}
	var data []models.Variables
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]config_setup.Variables, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.Variables{}.AttrTypes(), &item)
		unpacked, d := unpackVariablesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.Variables", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for Variables ---
func packVariablesListFromSdk(ctx context.Context, sdks []config_setup.Variables) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.Variables")
	diags := diag.Diagnostics{}
	var data []models.Variables

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.Variables
		obj, d := packVariablesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.Variables{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.Variables", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.Variables{}.AttrType(), data)
}
