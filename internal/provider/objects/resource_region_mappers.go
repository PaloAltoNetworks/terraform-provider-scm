package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/objects"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/objects"
)

// --- Unpacker for Regions ---
func unpackRegionsToSdk(ctx context.Context, obj types.Object) (*objects.Regions, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.Regions", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.Regions
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.Regions
	var d diag.Diagnostics

	// Handling Lists
	if !model.Address.IsNull() && !model.Address.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Address")
		diags.Append(model.Address.ElementsAs(ctx, &sdk.Address, false)...)
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

	// Handling Objects
	if !model.GeoLocation.IsNull() && !model.GeoLocation.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field GeoLocation")
		unpacked, d := unpackRegionsGeoLocationToSdk(ctx, model.GeoLocation)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "GeoLocation"})
		}
		if unpacked != nil {
			sdk.GeoLocation = unpacked
		}
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

	tflog.Debug(ctx, "Exiting unpack helper for models.Regions", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for Regions ---
func packRegionsFromSdk(ctx context.Context, sdk objects.Regions) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.Regions", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.Regions
	var d diag.Diagnostics
	// Handling Lists
	if sdk.Address != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Address")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Address, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Address)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Address = basetypes.NewListNull(elemType)
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
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.GeoLocation != nil {
		tflog.Debug(ctx, "Packing nested object for field GeoLocation")
		packed, d := packRegionsGeoLocationFromSdk(ctx, *sdk.GeoLocation)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "GeoLocation"})
		}
		model.GeoLocation = packed
	} else {
		model.GeoLocation = basetypes.NewObjectNull(models.RegionsGeoLocation{}.AttrTypes())
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
	if sdk.Snippet != nil {
		model.Snippet = basetypes.NewStringValue(*sdk.Snippet)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	} else {
		model.Snippet = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.Regions{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.Regions", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for Regions ---
func unpackRegionsListToSdk(ctx context.Context, list types.List) ([]objects.Regions, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.Regions")
	diags := diag.Diagnostics{}
	var data []models.Regions
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.Regions, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.Regions{}.AttrTypes(), &item)
		unpacked, d := unpackRegionsToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.Regions", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for Regions ---
func packRegionsListFromSdk(ctx context.Context, sdks []objects.Regions) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.Regions")
	diags := diag.Diagnostics{}
	var data []models.Regions

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.Regions
		obj, d := packRegionsFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.Regions{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.Regions", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.Regions{}.AttrType(), data)
}

// --- Unpacker for RegionsGeoLocation ---
func unpackRegionsGeoLocationToSdk(ctx context.Context, obj types.Object) (*objects.RegionsGeoLocation, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.RegionsGeoLocation", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.RegionsGeoLocation
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.RegionsGeoLocation
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Latitude.IsNull() && !model.Latitude.IsUnknown() {
		sdk.Latitude = float32(model.Latitude.ValueFloat64())
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Latitude", "value": sdk.Latitude})
	}

	// Handling Primitives
	if !model.Longitude.IsNull() && !model.Longitude.IsUnknown() {
		sdk.Longitude = float32(model.Longitude.ValueFloat64())
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Longitude", "value": sdk.Longitude})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.RegionsGeoLocation", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for RegionsGeoLocation ---
func packRegionsGeoLocationFromSdk(ctx context.Context, sdk objects.RegionsGeoLocation) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.RegionsGeoLocation", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.RegionsGeoLocation
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.Latitude = basetypes.NewFloat64Value(float64(sdk.Latitude))
	// Handling Primitives
	// Standard primitive packing
	model.Longitude = basetypes.NewFloat64Value(float64(sdk.Longitude))
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.RegionsGeoLocation{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.RegionsGeoLocation", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for RegionsGeoLocation ---
func unpackRegionsGeoLocationListToSdk(ctx context.Context, list types.List) ([]objects.RegionsGeoLocation, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.RegionsGeoLocation")
	diags := diag.Diagnostics{}
	var data []models.RegionsGeoLocation
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.RegionsGeoLocation, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.RegionsGeoLocation{}.AttrTypes(), &item)
		unpacked, d := unpackRegionsGeoLocationToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.RegionsGeoLocation", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for RegionsGeoLocation ---
func packRegionsGeoLocationListFromSdk(ctx context.Context, sdks []objects.RegionsGeoLocation) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.RegionsGeoLocation")
	diags := diag.Diagnostics{}
	var data []models.RegionsGeoLocation

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.RegionsGeoLocation
		obj, d := packRegionsGeoLocationFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.RegionsGeoLocation{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.RegionsGeoLocation", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.RegionsGeoLocation{}.AttrType(), data)
}
