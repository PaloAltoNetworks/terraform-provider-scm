package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/deployment_services"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/deployment_services"
)

// --- Unpacker for Locations ---
func unpackLocationsToSdk(ctx context.Context, obj types.Object) (*deployment_services.Locations, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.Locations", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.Locations
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk deployment_services.Locations
	var d diag.Diagnostics

	// Handling Primitives
	if !model.AggregateRegion.IsNull() && !model.AggregateRegion.IsUnknown() {
		sdk.AggregateRegion = model.AggregateRegion.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "AggregateRegion", "value": *sdk.AggregateRegion})
	}

	// Handling Primitives
	if !model.Continent.IsNull() && !model.Continent.IsUnknown() {
		sdk.Continent = model.Continent.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Continent", "value": *sdk.Continent})
	}

	// Handling Primitives
	if !model.Display.IsNull() && !model.Display.IsUnknown() {
		sdk.Display = model.Display.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Display", "value": *sdk.Display})
	}

	// Handling Primitives
	if !model.Latitude.IsNull() && !model.Latitude.IsUnknown() {
		val := float32(model.Latitude.ValueFloat64())
		sdk.Latitude = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Latitude", "value": *sdk.Latitude})
	}

	// Handling Primitives
	if !model.Longitude.IsNull() && !model.Longitude.IsUnknown() {
		val := float32(model.Longitude.ValueFloat64())
		sdk.Longitude = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Longitude", "value": *sdk.Longitude})
	}

	// Handling Primitives
	if !model.Region.IsNull() && !model.Region.IsUnknown() {
		sdk.Region = model.Region.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Region", "value": *sdk.Region})
	}

	// Handling Primitives
	if !model.Value.IsNull() && !model.Value.IsUnknown() {
		sdk.Value = model.Value.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Value", "value": *sdk.Value})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.Locations", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for Locations ---
func packLocationsFromSdk(ctx context.Context, sdk deployment_services.Locations) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.Locations", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.Locations
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.AggregateRegion != nil {
		model.AggregateRegion = basetypes.NewStringValue(*sdk.AggregateRegion)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "AggregateRegion", "value": *sdk.AggregateRegion})
	} else {
		model.AggregateRegion = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Continent != nil {
		model.Continent = basetypes.NewStringValue(*sdk.Continent)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Continent", "value": *sdk.Continent})
	} else {
		model.Continent = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Display != nil {
		model.Display = basetypes.NewStringValue(*sdk.Display)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Display", "value": *sdk.Display})
	} else {
		model.Display = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Latitude != nil {
		model.Latitude = basetypes.NewFloat64Value(float64(*sdk.Latitude))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Latitude", "value": *sdk.Latitude})
	} else {
		model.Latitude = basetypes.NewFloat64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Longitude != nil {
		model.Longitude = basetypes.NewFloat64Value(float64(*sdk.Longitude))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Longitude", "value": *sdk.Longitude})
	} else {
		model.Longitude = basetypes.NewFloat64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Region != nil {
		model.Region = basetypes.NewStringValue(*sdk.Region)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Region", "value": *sdk.Region})
	} else {
		model.Region = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Value != nil {
		model.Value = basetypes.NewStringValue(*sdk.Value)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Value", "value": *sdk.Value})
	} else {
		model.Value = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.Locations{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.Locations", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for Locations ---
func unpackLocationsListToSdk(ctx context.Context, list types.List) ([]deployment_services.Locations, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.Locations")
	diags := diag.Diagnostics{}
	var data []models.Locations
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]deployment_services.Locations, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.Locations{}.AttrTypes(), &item)
		unpacked, d := unpackLocationsToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.Locations", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for Locations ---
func packLocationsListFromSdk(ctx context.Context, sdks []deployment_services.Locations) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.Locations")
	diags := diag.Diagnostics{}
	var data []models.Locations

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.Locations
		obj, d := packLocationsFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.Locations{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.Locations", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.Locations{}.AttrType(), data)
}
