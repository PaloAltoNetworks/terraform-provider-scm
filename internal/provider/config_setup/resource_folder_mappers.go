package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/config_setup"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/config_setup"
)

// --- Unpacker for Folders ---
func unpackFoldersToSdk(ctx context.Context, obj types.Object) (*config_setup.Folders, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.Folders", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.Folders
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk config_setup.Folders
	var d diag.Diagnostics

	// Handling Primitives
	if !model.Description.IsNull() && !model.Description.IsUnknown() {
		sdk.Description = model.Description.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Description", "value": *sdk.Description})
	}

	// Handling Primitives
	if !model.Id.IsNull() && !model.Id.IsUnknown() {
		sdk.Id = model.Id.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Id", "value": sdk.Id})
	}

	// Handling Lists
	if !model.Labels.IsNull() && !model.Labels.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Labels")
		diags.Append(model.Labels.ElementsAs(ctx, &sdk.Labels, false)...)
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	// Handling Primitives
	if !model.Parent.IsNull() && !model.Parent.IsUnknown() {
		sdk.Parent = model.Parent.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Parent", "value": sdk.Parent})
	}

	// Handling Lists
	if !model.Snippets.IsNull() && !model.Snippets.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Snippets")
		diags.Append(model.Snippets.ElementsAs(ctx, &sdk.Snippets, false)...)
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.Folders", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for Folders ---
func packFoldersFromSdk(ctx context.Context, sdk config_setup.Folders) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.Folders", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.Folders
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
	model.Id = basetypes.NewStringValue(sdk.Id)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Id", "value": sdk.Id})
	// Handling Lists
	if sdk.Labels != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Labels")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Labels, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Labels)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Labels = basetypes.NewListNull(elemType)
	}
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	// Handling Primitives
	// Standard primitive packing
	model.Parent = basetypes.NewStringValue(sdk.Parent)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Parent", "value": sdk.Parent})
	// Handling Lists
	if sdk.Snippets != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Snippets")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Snippets, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Snippets)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Snippets = basetypes.NewListNull(elemType)
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.Folders{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.Folders", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for Folders ---
func unpackFoldersListToSdk(ctx context.Context, list types.List) ([]config_setup.Folders, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.Folders")
	diags := diag.Diagnostics{}
	var data []models.Folders
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]config_setup.Folders, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.Folders{}.AttrTypes(), &item)
		unpacked, d := unpackFoldersToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.Folders", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for Folders ---
func packFoldersListFromSdk(ctx context.Context, sdks []config_setup.Folders) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.Folders")
	diags := diag.Diagnostics{}
	var data []models.Folders

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.Folders
		obj, d := packFoldersFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.Folders{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.Folders", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.Folders{}.AttrType(), data)
}
