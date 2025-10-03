package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/deployment_services"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/deployment_services"
)

// --- Unpacker for ServiceConnectionGroups ---
func unpackServiceConnectionGroupsToSdk(ctx context.Context, obj types.Object) (*deployment_services.ServiceConnectionGroups, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ServiceConnectionGroups", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ServiceConnectionGroups
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk deployment_services.ServiceConnectionGroups
	var d diag.Diagnostics
	// Handling Primitives
	if !model.DisableSnat.IsNull() && !model.DisableSnat.IsUnknown() {
		sdk.DisableSnat = model.DisableSnat.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DisableSnat", "value": *sdk.DisableSnat})
	}

	// Handling Primitives
	if !model.Folder.IsNull() && !model.Folder.IsUnknown() {
		sdk.Folder = model.Folder.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Folder", "value": sdk.Folder})
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
	if !model.PbfOnly.IsNull() && !model.PbfOnly.IsUnknown() {
		sdk.PbfOnly = model.PbfOnly.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "PbfOnly", "value": *sdk.PbfOnly})
	}

	// Handling Lists
	if !model.Target.IsNull() && !model.Target.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Target")
		diags.Append(model.Target.ElementsAs(ctx, &sdk.Target, false)...)
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ServiceConnectionGroups", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ServiceConnectionGroups ---
func packServiceConnectionGroupsFromSdk(ctx context.Context, sdk deployment_services.ServiceConnectionGroups) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ServiceConnectionGroups", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ServiceConnectionGroups
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.DisableSnat != nil {
		model.DisableSnat = basetypes.NewBoolValue(*sdk.DisableSnat)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DisableSnat", "value": *sdk.DisableSnat})
	} else {
		model.DisableSnat = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	model.Folder = basetypes.NewStringValue(sdk.Folder)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Folder", "value": sdk.Folder})
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
	if sdk.PbfOnly != nil {
		model.PbfOnly = basetypes.NewBoolValue(*sdk.PbfOnly)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "PbfOnly", "value": *sdk.PbfOnly})
	} else {
		model.PbfOnly = basetypes.NewBoolNull()
	}
	// Handling Lists
	if sdk.Target != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Target")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Target, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Target)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Target = basetypes.NewListNull(elemType)
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ServiceConnectionGroups{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ServiceConnectionGroups", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ServiceConnectionGroups ---
func unpackServiceConnectionGroupsListToSdk(ctx context.Context, list types.List) ([]deployment_services.ServiceConnectionGroups, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ServiceConnectionGroups")
	diags := diag.Diagnostics{}
	var data []models.ServiceConnectionGroups
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]deployment_services.ServiceConnectionGroups, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ServiceConnectionGroups{}.AttrTypes(), &item)
		unpacked, d := unpackServiceConnectionGroupsToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ServiceConnectionGroups", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ServiceConnectionGroups ---
func packServiceConnectionGroupsListFromSdk(ctx context.Context, sdks []deployment_services.ServiceConnectionGroups) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ServiceConnectionGroups")
	diags := diag.Diagnostics{}
	var data []models.ServiceConnectionGroups

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ServiceConnectionGroups
		obj, d := packServiceConnectionGroupsFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ServiceConnectionGroups{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ServiceConnectionGroups", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ServiceConnectionGroups{}.AttrType(), data)
}
