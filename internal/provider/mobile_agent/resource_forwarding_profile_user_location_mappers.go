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

// --- Unpacker for ForwardingProfileUserLocations ---
func unpackForwardingProfileUserLocationsToSdk(ctx context.Context, obj types.Object) (*mobile_agent.ForwardingProfileUserLocations, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ForwardingProfileUserLocations", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ForwardingProfileUserLocations
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk mobile_agent.ForwardingProfileUserLocations
	var d diag.Diagnostics

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

	// Handling Objects
	if !model.InternalHostDetection.IsNull() && !model.InternalHostDetection.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field InternalHostDetection")
		unpacked, d := unpackForwardingProfileUserLocationsInternalHostDetectionToSdk(ctx, model.InternalHostDetection)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "InternalHostDetection"})
		}
		if unpacked != nil {
			sdk.InternalHostDetection = unpacked
		}
	}

	// Handling Lists
	if !model.IpAddresses.IsNull() && !model.IpAddresses.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field IpAddresses")
		diags.Append(model.IpAddresses.ElementsAs(ctx, &sdk.IpAddresses, false)...)
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ForwardingProfileUserLocations", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ForwardingProfileUserLocations ---
func packForwardingProfileUserLocationsFromSdk(ctx context.Context, sdk mobile_agent.ForwardingProfileUserLocations) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ForwardingProfileUserLocations", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ForwardingProfileUserLocations
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
	if sdk.Id != nil {
		model.Id = basetypes.NewStringValue(*sdk.Id)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Id", "value": *sdk.Id})
	} else {
		model.Id = basetypes.NewStringNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.InternalHostDetection != nil {
		tflog.Debug(ctx, "Packing nested object for field InternalHostDetection")
		packed, d := packForwardingProfileUserLocationsInternalHostDetectionFromSdk(ctx, *sdk.InternalHostDetection)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "InternalHostDetection"})
		}
		model.InternalHostDetection = packed
	} else {
		model.InternalHostDetection = basetypes.NewObjectNull(models.ForwardingProfileUserLocationsInternalHostDetection{}.AttrTypes())
	}
	// Handling Lists
	if sdk.IpAddresses != nil {
		tflog.Debug(ctx, "Packing list of primitives for field IpAddresses")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.IpAddresses, d = basetypes.NewListValueFrom(ctx, elemType, sdk.IpAddresses)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.IpAddresses = basetypes.NewListNull(elemType)
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Name != nil {
		model.Name = basetypes.NewStringValue(*sdk.Name)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
	} else {
		model.Name = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ForwardingProfileUserLocations{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ForwardingProfileUserLocations", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ForwardingProfileUserLocations ---
func unpackForwardingProfileUserLocationsListToSdk(ctx context.Context, list types.List) ([]mobile_agent.ForwardingProfileUserLocations, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ForwardingProfileUserLocations")
	diags := diag.Diagnostics{}
	var data []models.ForwardingProfileUserLocations
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]mobile_agent.ForwardingProfileUserLocations, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ForwardingProfileUserLocations{}.AttrTypes(), &item)
		unpacked, d := unpackForwardingProfileUserLocationsToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ForwardingProfileUserLocations", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ForwardingProfileUserLocations ---
func packForwardingProfileUserLocationsListFromSdk(ctx context.Context, sdks []mobile_agent.ForwardingProfileUserLocations) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ForwardingProfileUserLocations")
	diags := diag.Diagnostics{}
	var data []models.ForwardingProfileUserLocations

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ForwardingProfileUserLocations
		obj, d := packForwardingProfileUserLocationsFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ForwardingProfileUserLocations{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ForwardingProfileUserLocations", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ForwardingProfileUserLocations{}.AttrType(), data)
}

// --- Unpacker for ForwardingProfileUserLocationsInternalHostDetection ---
func unpackForwardingProfileUserLocationsInternalHostDetectionToSdk(ctx context.Context, obj types.Object) (*mobile_agent.ForwardingProfileUserLocationsInternalHostDetection, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ForwardingProfileUserLocationsInternalHostDetection", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ForwardingProfileUserLocationsInternalHostDetection
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk mobile_agent.ForwardingProfileUserLocationsInternalHostDetection
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Fqdn.IsNull() && !model.Fqdn.IsUnknown() {
		sdk.Fqdn = model.Fqdn.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Fqdn", "value": sdk.Fqdn})
	}

	// Handling Primitives
	if !model.IpAddress.IsNull() && !model.IpAddress.IsUnknown() {
		sdk.IpAddress = model.IpAddress.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "IpAddress", "value": sdk.IpAddress})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ForwardingProfileUserLocationsInternalHostDetection", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ForwardingProfileUserLocationsInternalHostDetection ---
func packForwardingProfileUserLocationsInternalHostDetectionFromSdk(ctx context.Context, sdk mobile_agent.ForwardingProfileUserLocationsInternalHostDetection) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ForwardingProfileUserLocationsInternalHostDetection", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ForwardingProfileUserLocationsInternalHostDetection
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.Fqdn = basetypes.NewStringValue(sdk.Fqdn)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Fqdn", "value": sdk.Fqdn})
	// Handling Primitives
	// Standard primitive packing
	model.IpAddress = basetypes.NewStringValue(sdk.IpAddress)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "IpAddress", "value": sdk.IpAddress})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ForwardingProfileUserLocationsInternalHostDetection{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ForwardingProfileUserLocationsInternalHostDetection", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ForwardingProfileUserLocationsInternalHostDetection ---
func unpackForwardingProfileUserLocationsInternalHostDetectionListToSdk(ctx context.Context, list types.List) ([]mobile_agent.ForwardingProfileUserLocationsInternalHostDetection, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ForwardingProfileUserLocationsInternalHostDetection")
	diags := diag.Diagnostics{}
	var data []models.ForwardingProfileUserLocationsInternalHostDetection
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]mobile_agent.ForwardingProfileUserLocationsInternalHostDetection, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ForwardingProfileUserLocationsInternalHostDetection{}.AttrTypes(), &item)
		unpacked, d := unpackForwardingProfileUserLocationsInternalHostDetectionToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ForwardingProfileUserLocationsInternalHostDetection", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ForwardingProfileUserLocationsInternalHostDetection ---
func packForwardingProfileUserLocationsInternalHostDetectionListFromSdk(ctx context.Context, sdks []mobile_agent.ForwardingProfileUserLocationsInternalHostDetection) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ForwardingProfileUserLocationsInternalHostDetection")
	diags := diag.Diagnostics{}
	var data []models.ForwardingProfileUserLocationsInternalHostDetection

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ForwardingProfileUserLocationsInternalHostDetection
		obj, d := packForwardingProfileUserLocationsInternalHostDetectionFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ForwardingProfileUserLocationsInternalHostDetection{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ForwardingProfileUserLocationsInternalHostDetection", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ForwardingProfileUserLocationsInternalHostDetection{}.AttrType(), data)
}
