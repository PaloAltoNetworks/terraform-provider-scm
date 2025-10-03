package provider

import (
	"context"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/deployment_services"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/deployment_services"
)

// --- Unpacker for BgpRouting ---
func unpackBgpRoutingToSdk(ctx context.Context, obj types.Object) (*deployment_services.BgpRouting, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpRouting", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpRouting
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk deployment_services.BgpRouting
	var d diag.Diagnostics
	// Handling Primitives
	if !model.AcceptRouteOverSC.IsNull() && !model.AcceptRouteOverSC.IsUnknown() {
		sdk.AcceptRouteOverSC = model.AcceptRouteOverSC.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "AcceptRouteOverSC", "value": *sdk.AcceptRouteOverSC})
	}

	// Handling Primitives
	if !model.AddHostRouteToIkePeer.IsNull() && !model.AddHostRouteToIkePeer.IsUnknown() {
		sdk.AddHostRouteToIkePeer = model.AddHostRouteToIkePeer.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "AddHostRouteToIkePeer", "value": *sdk.AddHostRouteToIkePeer})
	}

	// Handling Primitives
	if !model.BackboneRouting.IsNull() && !model.BackboneRouting.IsUnknown() {
		sdk.BackboneRouting = model.BackboneRouting.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "BackboneRouting", "value": *sdk.BackboneRouting})
	}

	// Handling Lists
	if !model.OutboundRoutesForServices.IsNull() && !model.OutboundRoutesForServices.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field OutboundRoutesForServices")
		diags.Append(model.OutboundRoutesForServices.ElementsAs(ctx, &sdk.OutboundRoutesForServices, false)...)
	}

	// Handling Objects
	if !model.RoutingPreference.IsNull() && !model.RoutingPreference.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field RoutingPreference")
		unpacked, d := unpackBgpRoutingRoutingPreferenceToSdk(ctx, model.RoutingPreference)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "RoutingPreference"})
		}
		if unpacked != nil {
			sdk.RoutingPreference = unpacked
		}
	}

	// Handling Primitives
	if !model.WithdrawStaticRoute.IsNull() && !model.WithdrawStaticRoute.IsUnknown() {
		sdk.WithdrawStaticRoute = model.WithdrawStaticRoute.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "WithdrawStaticRoute", "value": *sdk.WithdrawStaticRoute})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpRouting", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpRouting ---
func packBgpRoutingFromSdk(ctx context.Context, sdk deployment_services.BgpRouting) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpRouting", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpRouting
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.AcceptRouteOverSC != nil {
		model.AcceptRouteOverSC = basetypes.NewBoolValue(*sdk.AcceptRouteOverSC)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "AcceptRouteOverSC", "value": *sdk.AcceptRouteOverSC})
	} else {
		model.AcceptRouteOverSC = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.AddHostRouteToIkePeer != nil {
		model.AddHostRouteToIkePeer = basetypes.NewBoolValue(*sdk.AddHostRouteToIkePeer)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "AddHostRouteToIkePeer", "value": *sdk.AddHostRouteToIkePeer})
	} else {
		model.AddHostRouteToIkePeer = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.BackboneRouting != nil {
		model.BackboneRouting = basetypes.NewStringValue(*sdk.BackboneRouting)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "BackboneRouting", "value": *sdk.BackboneRouting})
	} else {
		model.BackboneRouting = basetypes.NewStringNull()
	}
	// Handling Lists
	if sdk.OutboundRoutesForServices != nil {
		tflog.Debug(ctx, "Packing list of primitives for field OutboundRoutesForServices")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.OutboundRoutesForServices, d = basetypes.NewListValueFrom(ctx, elemType, sdk.OutboundRoutesForServices)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.OutboundRoutesForServices = basetypes.NewListNull(elemType)
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.RoutingPreference != nil {
		tflog.Debug(ctx, "Packing nested object for field RoutingPreference")
		packed, d := packBgpRoutingRoutingPreferenceFromSdk(ctx, *sdk.RoutingPreference)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "RoutingPreference"})
		}
		model.RoutingPreference = packed
	} else {
		model.RoutingPreference = basetypes.NewObjectNull(models.BgpRoutingRoutingPreference{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.WithdrawStaticRoute != nil {
		model.WithdrawStaticRoute = basetypes.NewBoolValue(*sdk.WithdrawStaticRoute)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "WithdrawStaticRoute", "value": *sdk.WithdrawStaticRoute})
	} else {
		model.WithdrawStaticRoute = basetypes.NewBoolNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BgpRouting{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpRouting", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpRouting ---
func unpackBgpRoutingListToSdk(ctx context.Context, list types.List) ([]deployment_services.BgpRouting, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpRouting")
	diags := diag.Diagnostics{}
	var data []models.BgpRouting
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]deployment_services.BgpRouting, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpRouting{}.AttrTypes(), &item)
		unpacked, d := unpackBgpRoutingToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpRouting", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpRouting ---
func packBgpRoutingListFromSdk(ctx context.Context, sdks []deployment_services.BgpRouting) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpRouting")
	diags := diag.Diagnostics{}
	var data []models.BgpRouting

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpRouting
		obj, d := packBgpRoutingFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpRouting{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpRouting", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpRouting{}.AttrType(), data)
}

// --- Unpacker for BgpRoutingRoutingPreference ---
func unpackBgpRoutingRoutingPreferenceToSdk(ctx context.Context, obj types.Object) (*deployment_services.BgpRoutingRoutingPreference, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpRoutingRoutingPreference", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpRoutingRoutingPreference
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk deployment_services.BgpRoutingRoutingPreference
	var d diag.Diagnostics
	// Handling Typeless Objects
	if !model.Default.IsNull() && !model.Default.IsUnknown() {
		tflog.Debug(ctx, "Unpacking typeless object for field Default")
		sdk.Default = make(map[string]interface{})
	}

	// Handling Typeless Objects
	if !model.HotPotatoRouting.IsNull() && !model.HotPotatoRouting.IsUnknown() {
		tflog.Debug(ctx, "Unpacking typeless object for field HotPotatoRouting")
		sdk.HotPotatoRouting = make(map[string]interface{})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpRoutingRoutingPreference", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpRoutingRoutingPreference ---
func packBgpRoutingRoutingPreferenceFromSdk(ctx context.Context, sdk deployment_services.BgpRoutingRoutingPreference) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpRoutingRoutingPreference", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpRoutingRoutingPreference
	var d diag.Diagnostics
	// Handling Objects
	// This is a marker object (e.g. CHAP: {}). We just need to create an empty, non-null object.
	if sdk.Default != nil && !reflect.ValueOf(sdk.Default).IsNil() {
		tflog.Debug(ctx, "Packing typeless object for field Default")
		var d diag.Diagnostics
		// Create an empty object with no attributes, which signifies its presence.
		model.Default, d = basetypes.NewObjectValue(map[string]attr.Type{}, map[string]attr.Value{})
		diags.Append(d...)
	} else {
		// Since this field is part of a oneOf, being nil means it's not selected.
		// We make the object null with an empty attribute map.
		model.Default = basetypes.NewObjectNull(map[string]attr.Type{})
	}
	// Handling Objects
	// This is a marker object (e.g. CHAP: {}). We just need to create an empty, non-null object.
	if sdk.HotPotatoRouting != nil && !reflect.ValueOf(sdk.HotPotatoRouting).IsNil() {
		tflog.Debug(ctx, "Packing typeless object for field HotPotatoRouting")
		var d diag.Diagnostics
		// Create an empty object with no attributes, which signifies its presence.
		model.HotPotatoRouting, d = basetypes.NewObjectValue(map[string]attr.Type{}, map[string]attr.Value{})
		diags.Append(d...)
	} else {
		// Since this field is part of a oneOf, being nil means it's not selected.
		// We make the object null with an empty attribute map.
		model.HotPotatoRouting = basetypes.NewObjectNull(map[string]attr.Type{})
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BgpRoutingRoutingPreference{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpRoutingRoutingPreference", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpRoutingRoutingPreference ---
func unpackBgpRoutingRoutingPreferenceListToSdk(ctx context.Context, list types.List) ([]deployment_services.BgpRoutingRoutingPreference, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpRoutingRoutingPreference")
	diags := diag.Diagnostics{}
	var data []models.BgpRoutingRoutingPreference
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]deployment_services.BgpRoutingRoutingPreference, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpRoutingRoutingPreference{}.AttrTypes(), &item)
		unpacked, d := unpackBgpRoutingRoutingPreferenceToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpRoutingRoutingPreference", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpRoutingRoutingPreference ---
func packBgpRoutingRoutingPreferenceListFromSdk(ctx context.Context, sdks []deployment_services.BgpRoutingRoutingPreference) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpRoutingRoutingPreference")
	diags := diag.Diagnostics{}
	var data []models.BgpRoutingRoutingPreference

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpRoutingRoutingPreference
		obj, d := packBgpRoutingRoutingPreferenceFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpRoutingRoutingPreference{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpRoutingRoutingPreference", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpRoutingRoutingPreference{}.AttrType(), data)
}
