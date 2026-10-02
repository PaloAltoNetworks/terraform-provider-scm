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

// --- Unpacker for Subnets ---
func unpackSubnetsToSdk(ctx context.Context, obj types.Object) (*ztna_connector_all.Subnets, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.Subnets", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.Subnets
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk ztna_connector_all.Subnets
	var d diag.Diagnostics

	// Handling Primitives
	if !model.AppEnabled.IsNull() && !model.AppEnabled.IsUnknown() {
		sdk.AppEnabled = model.AppEnabled.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "AppEnabled", "value": *sdk.AppEnabled})
	}

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
	if !model.IcmpAllowed.IsNull() && !model.IcmpAllowed.IsUnknown() {
		sdk.IcmpAllowed = model.IcmpAllowed.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "IcmpAllowed", "value": *sdk.IcmpAllowed})
	}

	// Handling Primitives
	if !model.IpSubnets.IsNull() && !model.IpSubnets.IsUnknown() {
		sdk.IpSubnets = model.IpSubnets.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "IpSubnets", "value": sdk.IpSubnets})
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

	tflog.Debug(ctx, "Exiting unpack helper for models.Subnets", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for Subnets ---
func packSubnetsFromSdk(ctx context.Context, sdk ztna_connector_all.Subnets) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.Subnets", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.Subnets
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.AppEnabled != nil {
		model.AppEnabled = basetypes.NewBoolValue(*sdk.AppEnabled)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "AppEnabled", "value": *sdk.AppEnabled})
	} else {
		model.AppEnabled = basetypes.NewBoolNull()
	}
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
	if sdk.IcmpAllowed != nil {
		model.IcmpAllowed = basetypes.NewBoolValue(*sdk.IcmpAllowed)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "IcmpAllowed", "value": *sdk.IcmpAllowed})
	} else {
		model.IcmpAllowed = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	model.IpSubnets = basetypes.NewStringValue(sdk.IpSubnets)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "IpSubnets", "value": sdk.IpSubnets})
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

	obj, d := types.ObjectValueFrom(ctx, models.Subnets{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.Subnets", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for Subnets ---
func unpackSubnetsListToSdk(ctx context.Context, list types.List) ([]ztna_connector_all.Subnets, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.Subnets")
	diags := diag.Diagnostics{}
	var data []models.Subnets
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]ztna_connector_all.Subnets, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.Subnets{}.AttrTypes(), &item)
		unpacked, d := unpackSubnetsToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.Subnets", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for Subnets ---
func packSubnetsListFromSdk(ctx context.Context, sdks []ztna_connector_all.Subnets) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.Subnets")
	diags := diag.Diagnostics{}
	var data []models.Subnets

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.Subnets
		obj, d := packSubnetsFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.Subnets{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.Subnets", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.Subnets{}.AttrType(), data)
}
