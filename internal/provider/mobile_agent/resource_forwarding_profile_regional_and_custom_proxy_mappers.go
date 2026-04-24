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

// --- Unpacker for ForwardingProfileRegionalAndCustomProxies ---
func unpackForwardingProfileRegionalAndCustomProxiesToSdk(ctx context.Context, obj types.Object) (*mobile_agent.ForwardingProfileRegionalAndCustomProxies, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ForwardingProfileRegionalAndCustomProxies", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ForwardingProfileRegionalAndCustomProxies
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk mobile_agent.ForwardingProfileRegionalAndCustomProxies
	var d diag.Diagnostics

	// Handling Lists
	if !model.ConnectivityPreference.IsNull() && !model.ConnectivityPreference.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field ConnectivityPreference")
		unpacked, d := unpackForwardingProfileRegionalAndCustomProxiesConnectivityPreferenceInnerListToSdk(ctx, model.ConnectivityPreference)
		diags.Append(d...)
		sdk.ConnectivityPreference = unpacked
	}

	// Handling Primitives
	if !model.Description.IsNull() && !model.Description.IsUnknown() {
		sdk.Description = model.Description.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Description", "value": *sdk.Description})
	}

	// Handling Primitives
	if !model.FallbackOption.IsNull() && !model.FallbackOption.IsUnknown() {
		sdk.FallbackOption = model.FallbackOption.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "FallbackOption", "value": *sdk.FallbackOption})
	}

	// Handling Primitives
	if !model.Id.IsNull() && !model.Id.IsUnknown() {
		sdk.Id = model.Id.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Id", "value": *sdk.Id})
	}

	// Handling Primitives
	if !model.LocationPreference.IsNull() && !model.LocationPreference.IsUnknown() {
		sdk.LocationPreference = model.LocationPreference.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "LocationPreference", "value": *sdk.LocationPreference})
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	// Handling Lists
	if !model.PrismaAccessLocations.IsNull() && !model.PrismaAccessLocations.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field PrismaAccessLocations")
		unpacked, d := unpackForwardingProfileRegionalAndCustomProxiesPrismaAccessLocationsInnerListToSdk(ctx, model.PrismaAccessLocations)
		diags.Append(d...)
		sdk.PrismaAccessLocations = unpacked
	}

	// Handling Objects
	if !model.Proxy1.IsNull() && !model.Proxy1.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Proxy1")
		unpacked, d := unpackForwardingProfileRegionalAndCustomProxiesProxy1ToSdk(ctx, model.Proxy1)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Proxy1"})
		}
		if unpacked != nil {
			sdk.Proxy1 = unpacked
		}
	}

	// Handling Objects
	if !model.Proxy2.IsNull() && !model.Proxy2.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Proxy2")
		unpacked, d := unpackForwardingProfileRegionalAndCustomProxiesProxy2ToSdk(ctx, model.Proxy2)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Proxy2"})
		}
		if unpacked != nil {
			sdk.Proxy2 = unpacked
		}
	}

	// Handling Primitives
	if !model.Type.IsNull() && !model.Type.IsUnknown() {
		sdk.Type = model.Type.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Type", "value": *sdk.Type})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ForwardingProfileRegionalAndCustomProxies", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ForwardingProfileRegionalAndCustomProxies ---
func packForwardingProfileRegionalAndCustomProxiesFromSdk(ctx context.Context, sdk mobile_agent.ForwardingProfileRegionalAndCustomProxies) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ForwardingProfileRegionalAndCustomProxies", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ForwardingProfileRegionalAndCustomProxies
	var d diag.Diagnostics
	// Handling Lists
	if sdk.ConnectivityPreference != nil {
		tflog.Debug(ctx, "Packing list of objects for field ConnectivityPreference")
		packed, d := packForwardingProfileRegionalAndCustomProxiesConnectivityPreferenceInnerListFromSdk(ctx, sdk.ConnectivityPreference)
		diags.Append(d...)
		model.ConnectivityPreference = packed
	} else {
		model.ConnectivityPreference = basetypes.NewListNull(models.ForwardingProfileRegionalAndCustomProxiesConnectivityPreferenceInner{}.AttrType())
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
	if sdk.FallbackOption != nil {
		model.FallbackOption = basetypes.NewStringValue(*sdk.FallbackOption)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "FallbackOption", "value": *sdk.FallbackOption})
	} else {
		model.FallbackOption = basetypes.NewStringNull()
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
	if sdk.LocationPreference != nil {
		model.LocationPreference = basetypes.NewStringValue(*sdk.LocationPreference)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "LocationPreference", "value": *sdk.LocationPreference})
	} else {
		model.LocationPreference = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	// Handling Lists
	if sdk.PrismaAccessLocations != nil {
		tflog.Debug(ctx, "Packing list of objects for field PrismaAccessLocations")
		packed, d := packForwardingProfileRegionalAndCustomProxiesPrismaAccessLocationsInnerListFromSdk(ctx, sdk.PrismaAccessLocations)
		diags.Append(d...)
		model.PrismaAccessLocations = packed
	} else {
		model.PrismaAccessLocations = basetypes.NewListNull(models.ForwardingProfileRegionalAndCustomProxiesPrismaAccessLocationsInner{}.AttrType())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Proxy1 != nil {
		tflog.Debug(ctx, "Packing nested object for field Proxy1")
		packed, d := packForwardingProfileRegionalAndCustomProxiesProxy1FromSdk(ctx, *sdk.Proxy1)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Proxy1"})
		}
		model.Proxy1 = packed
	} else {
		model.Proxy1 = basetypes.NewObjectNull(models.ForwardingProfileRegionalAndCustomProxiesProxy1{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Proxy2 != nil {
		tflog.Debug(ctx, "Packing nested object for field Proxy2")
		packed, d := packForwardingProfileRegionalAndCustomProxiesProxy2FromSdk(ctx, *sdk.Proxy2)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Proxy2"})
		}
		model.Proxy2 = packed
	} else {
		model.Proxy2 = basetypes.NewObjectNull(models.ForwardingProfileRegionalAndCustomProxiesProxy2{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Type != nil {
		model.Type = basetypes.NewStringValue(*sdk.Type)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Type", "value": *sdk.Type})
	} else {
		model.Type = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ForwardingProfileRegionalAndCustomProxies{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ForwardingProfileRegionalAndCustomProxies", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ForwardingProfileRegionalAndCustomProxies ---
func unpackForwardingProfileRegionalAndCustomProxiesListToSdk(ctx context.Context, list types.List) ([]mobile_agent.ForwardingProfileRegionalAndCustomProxies, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ForwardingProfileRegionalAndCustomProxies")
	diags := diag.Diagnostics{}
	var data []models.ForwardingProfileRegionalAndCustomProxies
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]mobile_agent.ForwardingProfileRegionalAndCustomProxies, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ForwardingProfileRegionalAndCustomProxies{}.AttrTypes(), &item)
		unpacked, d := unpackForwardingProfileRegionalAndCustomProxiesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ForwardingProfileRegionalAndCustomProxies", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ForwardingProfileRegionalAndCustomProxies ---
func packForwardingProfileRegionalAndCustomProxiesListFromSdk(ctx context.Context, sdks []mobile_agent.ForwardingProfileRegionalAndCustomProxies) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ForwardingProfileRegionalAndCustomProxies")
	diags := diag.Diagnostics{}
	var data []models.ForwardingProfileRegionalAndCustomProxies

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ForwardingProfileRegionalAndCustomProxies
		obj, d := packForwardingProfileRegionalAndCustomProxiesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ForwardingProfileRegionalAndCustomProxies{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ForwardingProfileRegionalAndCustomProxies", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ForwardingProfileRegionalAndCustomProxies{}.AttrType(), data)
}

// --- Unpacker for ForwardingProfileRegionalAndCustomProxiesConnectivityPreferenceInner ---
func unpackForwardingProfileRegionalAndCustomProxiesConnectivityPreferenceInnerToSdk(ctx context.Context, obj types.Object) (*mobile_agent.ForwardingProfileRegionalAndCustomProxiesConnectivityPreferenceInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ForwardingProfileRegionalAndCustomProxiesConnectivityPreferenceInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ForwardingProfileRegionalAndCustomProxiesConnectivityPreferenceInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk mobile_agent.ForwardingProfileRegionalAndCustomProxiesConnectivityPreferenceInner
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Enabled.IsNull() && !model.Enabled.IsUnknown() {
		sdk.Enabled = model.Enabled.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Enabled", "value": *sdk.Enabled})
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ForwardingProfileRegionalAndCustomProxiesConnectivityPreferenceInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ForwardingProfileRegionalAndCustomProxiesConnectivityPreferenceInner ---
func packForwardingProfileRegionalAndCustomProxiesConnectivityPreferenceInnerFromSdk(ctx context.Context, sdk mobile_agent.ForwardingProfileRegionalAndCustomProxiesConnectivityPreferenceInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ForwardingProfileRegionalAndCustomProxiesConnectivityPreferenceInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ForwardingProfileRegionalAndCustomProxiesConnectivityPreferenceInner
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Enabled != nil {
		model.Enabled = basetypes.NewBoolValue(*sdk.Enabled)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Enabled", "value": *sdk.Enabled})
	} else {
		model.Enabled = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ForwardingProfileRegionalAndCustomProxiesConnectivityPreferenceInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ForwardingProfileRegionalAndCustomProxiesConnectivityPreferenceInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ForwardingProfileRegionalAndCustomProxiesConnectivityPreferenceInner ---
func unpackForwardingProfileRegionalAndCustomProxiesConnectivityPreferenceInnerListToSdk(ctx context.Context, list types.List) ([]mobile_agent.ForwardingProfileRegionalAndCustomProxiesConnectivityPreferenceInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ForwardingProfileRegionalAndCustomProxiesConnectivityPreferenceInner")
	diags := diag.Diagnostics{}
	var data []models.ForwardingProfileRegionalAndCustomProxiesConnectivityPreferenceInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]mobile_agent.ForwardingProfileRegionalAndCustomProxiesConnectivityPreferenceInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ForwardingProfileRegionalAndCustomProxiesConnectivityPreferenceInner{}.AttrTypes(), &item)
		unpacked, d := unpackForwardingProfileRegionalAndCustomProxiesConnectivityPreferenceInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ForwardingProfileRegionalAndCustomProxiesConnectivityPreferenceInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ForwardingProfileRegionalAndCustomProxiesConnectivityPreferenceInner ---
func packForwardingProfileRegionalAndCustomProxiesConnectivityPreferenceInnerListFromSdk(ctx context.Context, sdks []mobile_agent.ForwardingProfileRegionalAndCustomProxiesConnectivityPreferenceInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ForwardingProfileRegionalAndCustomProxiesConnectivityPreferenceInner")
	diags := diag.Diagnostics{}
	var data []models.ForwardingProfileRegionalAndCustomProxiesConnectivityPreferenceInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ForwardingProfileRegionalAndCustomProxiesConnectivityPreferenceInner
		obj, d := packForwardingProfileRegionalAndCustomProxiesConnectivityPreferenceInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ForwardingProfileRegionalAndCustomProxiesConnectivityPreferenceInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ForwardingProfileRegionalAndCustomProxiesConnectivityPreferenceInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ForwardingProfileRegionalAndCustomProxiesConnectivityPreferenceInner{}.AttrType(), data)
}

// --- Unpacker for ForwardingProfileRegionalAndCustomProxiesPrismaAccessLocationsInner ---
func unpackForwardingProfileRegionalAndCustomProxiesPrismaAccessLocationsInnerToSdk(ctx context.Context, obj types.Object) (*mobile_agent.ForwardingProfileRegionalAndCustomProxiesPrismaAccessLocationsInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ForwardingProfileRegionalAndCustomProxiesPrismaAccessLocationsInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ForwardingProfileRegionalAndCustomProxiesPrismaAccessLocationsInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk mobile_agent.ForwardingProfileRegionalAndCustomProxiesPrismaAccessLocationsInner
	var d diag.Diagnostics
	// Handling Lists
	if !model.Locations.IsNull() && !model.Locations.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Locations")
		diags.Append(model.Locations.ElementsAs(ctx, &sdk.Locations, false)...)
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ForwardingProfileRegionalAndCustomProxiesPrismaAccessLocationsInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ForwardingProfileRegionalAndCustomProxiesPrismaAccessLocationsInner ---
func packForwardingProfileRegionalAndCustomProxiesPrismaAccessLocationsInnerFromSdk(ctx context.Context, sdk mobile_agent.ForwardingProfileRegionalAndCustomProxiesPrismaAccessLocationsInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ForwardingProfileRegionalAndCustomProxiesPrismaAccessLocationsInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ForwardingProfileRegionalAndCustomProxiesPrismaAccessLocationsInner
	var d diag.Diagnostics
	// Handling Lists
	if sdk.Locations != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Locations")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Locations, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Locations)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Locations = basetypes.NewListNull(elemType)
	}
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ForwardingProfileRegionalAndCustomProxiesPrismaAccessLocationsInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ForwardingProfileRegionalAndCustomProxiesPrismaAccessLocationsInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ForwardingProfileRegionalAndCustomProxiesPrismaAccessLocationsInner ---
func unpackForwardingProfileRegionalAndCustomProxiesPrismaAccessLocationsInnerListToSdk(ctx context.Context, list types.List) ([]mobile_agent.ForwardingProfileRegionalAndCustomProxiesPrismaAccessLocationsInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ForwardingProfileRegionalAndCustomProxiesPrismaAccessLocationsInner")
	diags := diag.Diagnostics{}
	var data []models.ForwardingProfileRegionalAndCustomProxiesPrismaAccessLocationsInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]mobile_agent.ForwardingProfileRegionalAndCustomProxiesPrismaAccessLocationsInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ForwardingProfileRegionalAndCustomProxiesPrismaAccessLocationsInner{}.AttrTypes(), &item)
		unpacked, d := unpackForwardingProfileRegionalAndCustomProxiesPrismaAccessLocationsInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ForwardingProfileRegionalAndCustomProxiesPrismaAccessLocationsInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ForwardingProfileRegionalAndCustomProxiesPrismaAccessLocationsInner ---
func packForwardingProfileRegionalAndCustomProxiesPrismaAccessLocationsInnerListFromSdk(ctx context.Context, sdks []mobile_agent.ForwardingProfileRegionalAndCustomProxiesPrismaAccessLocationsInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ForwardingProfileRegionalAndCustomProxiesPrismaAccessLocationsInner")
	diags := diag.Diagnostics{}
	var data []models.ForwardingProfileRegionalAndCustomProxiesPrismaAccessLocationsInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ForwardingProfileRegionalAndCustomProxiesPrismaAccessLocationsInner
		obj, d := packForwardingProfileRegionalAndCustomProxiesPrismaAccessLocationsInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ForwardingProfileRegionalAndCustomProxiesPrismaAccessLocationsInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ForwardingProfileRegionalAndCustomProxiesPrismaAccessLocationsInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ForwardingProfileRegionalAndCustomProxiesPrismaAccessLocationsInner{}.AttrType(), data)
}

// --- Unpacker for ForwardingProfileRegionalAndCustomProxiesProxy1 ---
func unpackForwardingProfileRegionalAndCustomProxiesProxy1ToSdk(ctx context.Context, obj types.Object) (*mobile_agent.ForwardingProfileRegionalAndCustomProxiesProxy1, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ForwardingProfileRegionalAndCustomProxiesProxy1", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ForwardingProfileRegionalAndCustomProxiesProxy1
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk mobile_agent.ForwardingProfileRegionalAndCustomProxiesProxy1
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Fqdn.IsNull() && !model.Fqdn.IsUnknown() {
		sdk.Fqdn = model.Fqdn.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Fqdn", "value": *sdk.Fqdn})
	}

	// Handling Primitives
	if !model.Location.IsNull() && !model.Location.IsUnknown() {
		sdk.Location = model.Location.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Location", "value": *sdk.Location})
	}

	// Handling Primitives
	if !model.Port.IsNull() && !model.Port.IsUnknown() {
		val := int32(model.Port.ValueInt64())
		sdk.Port = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Port", "value": *sdk.Port})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ForwardingProfileRegionalAndCustomProxiesProxy1", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ForwardingProfileRegionalAndCustomProxiesProxy1 ---
func packForwardingProfileRegionalAndCustomProxiesProxy1FromSdk(ctx context.Context, sdk mobile_agent.ForwardingProfileRegionalAndCustomProxiesProxy1) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ForwardingProfileRegionalAndCustomProxiesProxy1", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ForwardingProfileRegionalAndCustomProxiesProxy1
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Fqdn != nil {
		model.Fqdn = basetypes.NewStringValue(*sdk.Fqdn)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Fqdn", "value": *sdk.Fqdn})
	} else {
		model.Fqdn = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Location != nil {
		model.Location = basetypes.NewStringValue(*sdk.Location)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Location", "value": *sdk.Location})
	} else {
		model.Location = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Port != nil {
		model.Port = basetypes.NewInt64Value(int64(*sdk.Port))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Port", "value": *sdk.Port})
	} else {
		model.Port = basetypes.NewInt64Null()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ForwardingProfileRegionalAndCustomProxiesProxy1{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ForwardingProfileRegionalAndCustomProxiesProxy1", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ForwardingProfileRegionalAndCustomProxiesProxy1 ---
func unpackForwardingProfileRegionalAndCustomProxiesProxy1ListToSdk(ctx context.Context, list types.List) ([]mobile_agent.ForwardingProfileRegionalAndCustomProxiesProxy1, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ForwardingProfileRegionalAndCustomProxiesProxy1")
	diags := diag.Diagnostics{}
	var data []models.ForwardingProfileRegionalAndCustomProxiesProxy1
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]mobile_agent.ForwardingProfileRegionalAndCustomProxiesProxy1, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ForwardingProfileRegionalAndCustomProxiesProxy1{}.AttrTypes(), &item)
		unpacked, d := unpackForwardingProfileRegionalAndCustomProxiesProxy1ToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ForwardingProfileRegionalAndCustomProxiesProxy1", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ForwardingProfileRegionalAndCustomProxiesProxy1 ---
func packForwardingProfileRegionalAndCustomProxiesProxy1ListFromSdk(ctx context.Context, sdks []mobile_agent.ForwardingProfileRegionalAndCustomProxiesProxy1) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ForwardingProfileRegionalAndCustomProxiesProxy1")
	diags := diag.Diagnostics{}
	var data []models.ForwardingProfileRegionalAndCustomProxiesProxy1

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ForwardingProfileRegionalAndCustomProxiesProxy1
		obj, d := packForwardingProfileRegionalAndCustomProxiesProxy1FromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ForwardingProfileRegionalAndCustomProxiesProxy1{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ForwardingProfileRegionalAndCustomProxiesProxy1", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ForwardingProfileRegionalAndCustomProxiesProxy1{}.AttrType(), data)
}

// --- Unpacker for ForwardingProfileRegionalAndCustomProxiesProxy2 ---
func unpackForwardingProfileRegionalAndCustomProxiesProxy2ToSdk(ctx context.Context, obj types.Object) (*mobile_agent.ForwardingProfileRegionalAndCustomProxiesProxy2, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ForwardingProfileRegionalAndCustomProxiesProxy2", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ForwardingProfileRegionalAndCustomProxiesProxy2
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk mobile_agent.ForwardingProfileRegionalAndCustomProxiesProxy2
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Fqdn.IsNull() && !model.Fqdn.IsUnknown() {
		sdk.Fqdn = model.Fqdn.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Fqdn", "value": *sdk.Fqdn})
	}

	// Handling Primitives
	if !model.Location.IsNull() && !model.Location.IsUnknown() {
		sdk.Location = model.Location.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Location", "value": *sdk.Location})
	}

	// Handling Primitives
	if !model.Port.IsNull() && !model.Port.IsUnknown() {
		val := int32(model.Port.ValueInt64())
		sdk.Port = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Port", "value": *sdk.Port})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ForwardingProfileRegionalAndCustomProxiesProxy2", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ForwardingProfileRegionalAndCustomProxiesProxy2 ---
func packForwardingProfileRegionalAndCustomProxiesProxy2FromSdk(ctx context.Context, sdk mobile_agent.ForwardingProfileRegionalAndCustomProxiesProxy2) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ForwardingProfileRegionalAndCustomProxiesProxy2", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ForwardingProfileRegionalAndCustomProxiesProxy2
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Fqdn != nil {
		model.Fqdn = basetypes.NewStringValue(*sdk.Fqdn)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Fqdn", "value": *sdk.Fqdn})
	} else {
		model.Fqdn = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Location != nil {
		model.Location = basetypes.NewStringValue(*sdk.Location)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Location", "value": *sdk.Location})
	} else {
		model.Location = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Port != nil {
		model.Port = basetypes.NewInt64Value(int64(*sdk.Port))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Port", "value": *sdk.Port})
	} else {
		model.Port = basetypes.NewInt64Null()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ForwardingProfileRegionalAndCustomProxiesProxy2{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ForwardingProfileRegionalAndCustomProxiesProxy2", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ForwardingProfileRegionalAndCustomProxiesProxy2 ---
func unpackForwardingProfileRegionalAndCustomProxiesProxy2ListToSdk(ctx context.Context, list types.List) ([]mobile_agent.ForwardingProfileRegionalAndCustomProxiesProxy2, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ForwardingProfileRegionalAndCustomProxiesProxy2")
	diags := diag.Diagnostics{}
	var data []models.ForwardingProfileRegionalAndCustomProxiesProxy2
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]mobile_agent.ForwardingProfileRegionalAndCustomProxiesProxy2, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ForwardingProfileRegionalAndCustomProxiesProxy2{}.AttrTypes(), &item)
		unpacked, d := unpackForwardingProfileRegionalAndCustomProxiesProxy2ToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ForwardingProfileRegionalAndCustomProxiesProxy2", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ForwardingProfileRegionalAndCustomProxiesProxy2 ---
func packForwardingProfileRegionalAndCustomProxiesProxy2ListFromSdk(ctx context.Context, sdks []mobile_agent.ForwardingProfileRegionalAndCustomProxiesProxy2) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ForwardingProfileRegionalAndCustomProxiesProxy2")
	diags := diag.Diagnostics{}
	var data []models.ForwardingProfileRegionalAndCustomProxiesProxy2

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ForwardingProfileRegionalAndCustomProxiesProxy2
		obj, d := packForwardingProfileRegionalAndCustomProxiesProxy2FromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ForwardingProfileRegionalAndCustomProxiesProxy2{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ForwardingProfileRegionalAndCustomProxiesProxy2", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ForwardingProfileRegionalAndCustomProxiesProxy2{}.AttrType(), data)
}
