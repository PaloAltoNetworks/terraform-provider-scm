package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/mobile_agent"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/mobile_agent"
)

// --- Unpacker for ForwardingProfileSourceApplications ---
func unpackForwardingProfileSourceApplicationsToSdk(ctx context.Context, obj types.Object) (*mobile_agent.ForwardingProfileSourceApplications, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ForwardingProfileSourceApplications", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ForwardingProfileSourceApplications
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk mobile_agent.ForwardingProfileSourceApplications
	var d diag.Diagnostics

	// Handling Lists
	if !model.Applications.IsNull() && !model.Applications.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Applications")
		diags.Append(model.Applications.ElementsAs(ctx, &sdk.Applications, false)...)
	}

	// Handling Primitives
	if !model.Description.IsNull() && !model.Description.IsUnknown() {
		sdk.Description = model.Description.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Description", "value": *sdk.Description})
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

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ForwardingProfileSourceApplications", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ForwardingProfileSourceApplications ---
func packForwardingProfileSourceApplicationsFromSdk(ctx context.Context, sdk mobile_agent.ForwardingProfileSourceApplications) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ForwardingProfileSourceApplications", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ForwardingProfileSourceApplications
	var d diag.Diagnostics
	// Handling Lists
	if sdk.Applications != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Applications")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Applications, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Applications)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Applications = basetypes.NewListNull(elemType)
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
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ForwardingProfileSourceApplications{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ForwardingProfileSourceApplications", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ForwardingProfileSourceApplications ---
func unpackForwardingProfileSourceApplicationsListToSdk(ctx context.Context, list types.List) ([]mobile_agent.ForwardingProfileSourceApplications, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ForwardingProfileSourceApplications")
	diags := diag.Diagnostics{}
	var data []models.ForwardingProfileSourceApplications
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]mobile_agent.ForwardingProfileSourceApplications, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ForwardingProfileSourceApplications{}.AttrTypes(), &item)
		unpacked, d := unpackForwardingProfileSourceApplicationsToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ForwardingProfileSourceApplications", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ForwardingProfileSourceApplications ---
func packForwardingProfileSourceApplicationsListFromSdk(ctx context.Context, sdks []mobile_agent.ForwardingProfileSourceApplications) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ForwardingProfileSourceApplications")
	diags := diag.Diagnostics{}
	var data []models.ForwardingProfileSourceApplications

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ForwardingProfileSourceApplications
		obj, d := packForwardingProfileSourceApplicationsFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ForwardingProfileSourceApplications{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ForwardingProfileSourceApplications", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ForwardingProfileSourceApplications{}.AttrType(), data)
}
