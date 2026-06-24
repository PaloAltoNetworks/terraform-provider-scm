package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/ztna_connector_all"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/ztna_connector_all"
)

// --- Unpacker for Connectors ---
func unpackConnectorsToSdk(ctx context.Context, obj types.Object) (*ztna_connector_all.Connectors, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.Connectors", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.Connectors
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk ztna_connector_all.Connectors
	var d diag.Diagnostics

	// Handling Primitives
	if !model.CreatedTime.IsNull() && !model.CreatedTime.IsUnknown() {
		sdk.CreatedTime = model.CreatedTime.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "CreatedTime", "value": *sdk.CreatedTime})
	}

	// Handling Primitives
	if !model.Description.IsNull() && !model.Description.IsUnknown() {
		sdk.Description = model.Description.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Description", "value": *sdk.Description})
	}

	// Handling Primitives
	if !model.Group.IsNull() && !model.Group.IsUnknown() {
		sdk.Group = model.Group.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Group", "value": sdk.Group})
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	// Handling Primitives
	if !model.Oid.IsNull() && !model.Oid.IsUnknown() {
		sdk.Oid = model.Oid.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Oid", "value": *sdk.Oid})
	}

	// Handling Primitives
	if !model.UpdatedTime.IsNull() && !model.UpdatedTime.IsUnknown() {
		sdk.UpdatedTime = model.UpdatedTime.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "UpdatedTime", "value": *sdk.UpdatedTime})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.Connectors", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for Connectors ---
func packConnectorsFromSdk(ctx context.Context, sdk ztna_connector_all.Connectors) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.Connectors", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.Connectors
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.CreatedTime != nil {
		model.CreatedTime = basetypes.NewStringValue(*sdk.CreatedTime)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "CreatedTime", "value": *sdk.CreatedTime})
	} else {
		model.CreatedTime = basetypes.NewStringNull()
	}
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
	model.Group = basetypes.NewStringValue(sdk.Group)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Group", "value": sdk.Group})
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	// Handling Primitives
	// Standard primitive packing
	if sdk.Oid != nil {
		model.Oid = basetypes.NewStringValue(*sdk.Oid)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Oid", "value": *sdk.Oid})
	} else {
		model.Oid = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.UpdatedTime != nil {
		model.UpdatedTime = basetypes.NewStringValue(*sdk.UpdatedTime)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "UpdatedTime", "value": *sdk.UpdatedTime})
	} else {
		model.UpdatedTime = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.Connectors{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.Connectors", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for Connectors ---
func unpackConnectorsListToSdk(ctx context.Context, list types.List) ([]ztna_connector_all.Connectors, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.Connectors")
	diags := diag.Diagnostics{}
	var data []models.Connectors
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]ztna_connector_all.Connectors, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.Connectors{}.AttrTypes(), &item)
		unpacked, d := unpackConnectorsToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.Connectors", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for Connectors ---
func packConnectorsListFromSdk(ctx context.Context, sdks []ztna_connector_all.Connectors) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.Connectors")
	diags := diag.Diagnostics{}
	var data []models.Connectors

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.Connectors
		obj, d := packConnectorsFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.Connectors{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.Connectors", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.Connectors{}.AttrType(), data)
}
