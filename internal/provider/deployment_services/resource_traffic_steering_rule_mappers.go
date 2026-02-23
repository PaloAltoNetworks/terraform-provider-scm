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

// --- Unpacker for TrafficSteeringRules ---
func unpackTrafficSteeringRulesToSdk(ctx context.Context, obj types.Object) (*deployment_services.TrafficSteeringRules, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.TrafficSteeringRules", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.TrafficSteeringRules
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk deployment_services.TrafficSteeringRules
	var d diag.Diagnostics

	// Handling Objects
	if !model.Action.IsNull() && !model.Action.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Action")
		unpacked, d := unpackTrafficSteeringRulesActionToSdk(ctx, model.Action)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Action"})
		}
		if unpacked != nil {
			sdk.Action = unpacked
		}
	}

	// Handling Lists
	if !model.Category.IsNull() && !model.Category.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Category")
		diags.Append(model.Category.ElementsAs(ctx, &sdk.Category, false)...)
	}

	// Handling Lists
	if !model.Destination.IsNull() && !model.Destination.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Destination")
		diags.Append(model.Destination.ElementsAs(ctx, &sdk.Destination, false)...)
	}

	// Handling Primitives
	if !model.Folder.IsNull() && !model.Folder.IsUnknown() {
		sdk.Folder = model.Folder.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Folder", "value": *sdk.Folder})
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

	// Handling Lists
	if !model.Service.IsNull() && !model.Service.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Service")
		diags.Append(model.Service.ElementsAs(ctx, &sdk.Service, false)...)
	}

	// Handling Lists
	if !model.Source.IsNull() && !model.Source.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Source")
		diags.Append(model.Source.ElementsAs(ctx, &sdk.Source, false)...)
	}

	// Handling Lists
	if !model.SourceUser.IsNull() && !model.SourceUser.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field SourceUser")
		diags.Append(model.SourceUser.ElementsAs(ctx, &sdk.SourceUser, false)...)
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.TrafficSteeringRules", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for TrafficSteeringRules ---
func packTrafficSteeringRulesFromSdk(ctx context.Context, sdk deployment_services.TrafficSteeringRules) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.TrafficSteeringRules", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.TrafficSteeringRules
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Action != nil {
		tflog.Debug(ctx, "Packing nested object for field Action")
		packed, d := packTrafficSteeringRulesActionFromSdk(ctx, *sdk.Action)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Action"})
		}
		model.Action = packed
	} else {
		model.Action = basetypes.NewObjectNull(models.TrafficSteeringRulesAction{}.AttrTypes())
	}
	// Handling Lists
	if sdk.Category != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Category")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Category, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Category)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Category = basetypes.NewListNull(elemType)
	}
	// Handling Lists
	if sdk.Destination != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Destination")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Destination, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Destination)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Destination = basetypes.NewListNull(elemType)
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Folder != nil {
		model.Folder = basetypes.NewStringValue(*sdk.Folder)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Folder", "value": *sdk.Folder})
	} else {
		model.Folder = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	model.Id = basetypes.NewStringValue(sdk.Id)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Id", "value": sdk.Id})
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	// Handling Lists
	if sdk.Service != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Service")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Service, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Service)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Service = basetypes.NewListNull(elemType)
	}
	// Handling Lists
	if sdk.Source != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Source")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Source, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Source)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Source = basetypes.NewListNull(elemType)
	}
	// Handling Lists
	if sdk.SourceUser != nil {
		tflog.Debug(ctx, "Packing list of primitives for field SourceUser")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.SourceUser, d = basetypes.NewListValueFrom(ctx, elemType, sdk.SourceUser)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.SourceUser = basetypes.NewListNull(elemType)
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.TrafficSteeringRules{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.TrafficSteeringRules", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for TrafficSteeringRules ---
func unpackTrafficSteeringRulesListToSdk(ctx context.Context, list types.List) ([]deployment_services.TrafficSteeringRules, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.TrafficSteeringRules")
	diags := diag.Diagnostics{}
	var data []models.TrafficSteeringRules
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]deployment_services.TrafficSteeringRules, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.TrafficSteeringRules{}.AttrTypes(), &item)
		unpacked, d := unpackTrafficSteeringRulesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.TrafficSteeringRules", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for TrafficSteeringRules ---
func packTrafficSteeringRulesListFromSdk(ctx context.Context, sdks []deployment_services.TrafficSteeringRules) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.TrafficSteeringRules")
	diags := diag.Diagnostics{}
	var data []models.TrafficSteeringRules

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.TrafficSteeringRules
		obj, d := packTrafficSteeringRulesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.TrafficSteeringRules{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.TrafficSteeringRules", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.TrafficSteeringRules{}.AttrType(), data)
}

// --- Unpacker for TrafficSteeringRulesAction ---
func unpackTrafficSteeringRulesActionToSdk(ctx context.Context, obj types.Object) (*deployment_services.TrafficSteeringRulesAction, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.TrafficSteeringRulesAction", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.TrafficSteeringRulesAction
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk deployment_services.TrafficSteeringRulesAction
	var d diag.Diagnostics
	// Handling Objects
	if !model.Forward.IsNull() && !model.Forward.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Forward")
		unpacked, d := unpackTrafficSteeringRulesActionForwardToSdk(ctx, model.Forward)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Forward"})
		}
		if unpacked != nil {
			sdk.Forward = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.TrafficSteeringRulesAction", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for TrafficSteeringRulesAction ---
func packTrafficSteeringRulesActionFromSdk(ctx context.Context, sdk deployment_services.TrafficSteeringRulesAction) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.TrafficSteeringRulesAction", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.TrafficSteeringRulesAction
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Forward != nil {
		tflog.Debug(ctx, "Packing nested object for field Forward")
		packed, d := packTrafficSteeringRulesActionForwardFromSdk(ctx, *sdk.Forward)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Forward"})
		}
		model.Forward = packed
	} else {
		model.Forward = basetypes.NewObjectNull(models.TrafficSteeringRulesActionForward{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.TrafficSteeringRulesAction{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.TrafficSteeringRulesAction", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for TrafficSteeringRulesAction ---
func unpackTrafficSteeringRulesActionListToSdk(ctx context.Context, list types.List) ([]deployment_services.TrafficSteeringRulesAction, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.TrafficSteeringRulesAction")
	diags := diag.Diagnostics{}
	var data []models.TrafficSteeringRulesAction
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]deployment_services.TrafficSteeringRulesAction, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.TrafficSteeringRulesAction{}.AttrTypes(), &item)
		unpacked, d := unpackTrafficSteeringRulesActionToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.TrafficSteeringRulesAction", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for TrafficSteeringRulesAction ---
func packTrafficSteeringRulesActionListFromSdk(ctx context.Context, sdks []deployment_services.TrafficSteeringRulesAction) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.TrafficSteeringRulesAction")
	diags := diag.Diagnostics{}
	var data []models.TrafficSteeringRulesAction

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.TrafficSteeringRulesAction
		obj, d := packTrafficSteeringRulesActionFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.TrafficSteeringRulesAction{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.TrafficSteeringRulesAction", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.TrafficSteeringRulesAction{}.AttrType(), data)
}

// --- Unpacker for TrafficSteeringRulesActionForward ---
func unpackTrafficSteeringRulesActionForwardToSdk(ctx context.Context, obj types.Object) (*deployment_services.TrafficSteeringRulesActionForward, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.TrafficSteeringRulesActionForward", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.TrafficSteeringRulesActionForward
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk deployment_services.TrafficSteeringRulesActionForward
	var d diag.Diagnostics
	// Handling Objects
	if !model.Forward.IsNull() && !model.Forward.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Forward")
		unpacked, d := unpackTrafficSteeringRulesActionForwardForwardToSdk(ctx, model.Forward)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Forward"})
		}
		if unpacked != nil {
			sdk.Forward = unpacked
		}
	}

	// Handling Typeless Objects
	if !model.NoPbf.IsNull() && !model.NoPbf.IsUnknown() {
		tflog.Debug(ctx, "Unpacking typeless object for field NoPbf")
		sdk.NoPbf = make(map[string]interface{})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.TrafficSteeringRulesActionForward", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for TrafficSteeringRulesActionForward ---
func packTrafficSteeringRulesActionForwardFromSdk(ctx context.Context, sdk deployment_services.TrafficSteeringRulesActionForward) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.TrafficSteeringRulesActionForward", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.TrafficSteeringRulesActionForward
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Forward != nil {
		tflog.Debug(ctx, "Packing nested object for field Forward")
		packed, d := packTrafficSteeringRulesActionForwardForwardFromSdk(ctx, *sdk.Forward)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Forward"})
		}
		model.Forward = packed
	} else {
		model.Forward = basetypes.NewObjectNull(models.TrafficSteeringRulesActionForwardForward{}.AttrTypes())
	}
	// Handling Objects
	// This is a marker object (e.g. CHAP: {}). We just need to create an empty, non-null object.
	if sdk.NoPbf != nil && !reflect.ValueOf(sdk.NoPbf).IsNil() {
		tflog.Debug(ctx, "Packing typeless object for field NoPbf")
		var d diag.Diagnostics
		// Create an empty object with no attributes, which signifies its presence.
		model.NoPbf, d = basetypes.NewObjectValue(map[string]attr.Type{}, map[string]attr.Value{})
		diags.Append(d...)
	} else {
		// Since this field is part of a oneOf, being nil means it's not selected.
		// We make the object null with an empty attribute map.
		model.NoPbf = basetypes.NewObjectNull(map[string]attr.Type{})
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.TrafficSteeringRulesActionForward{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.TrafficSteeringRulesActionForward", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for TrafficSteeringRulesActionForward ---
func unpackTrafficSteeringRulesActionForwardListToSdk(ctx context.Context, list types.List) ([]deployment_services.TrafficSteeringRulesActionForward, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.TrafficSteeringRulesActionForward")
	diags := diag.Diagnostics{}
	var data []models.TrafficSteeringRulesActionForward
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]deployment_services.TrafficSteeringRulesActionForward, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.TrafficSteeringRulesActionForward{}.AttrTypes(), &item)
		unpacked, d := unpackTrafficSteeringRulesActionForwardToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.TrafficSteeringRulesActionForward", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for TrafficSteeringRulesActionForward ---
func packTrafficSteeringRulesActionForwardListFromSdk(ctx context.Context, sdks []deployment_services.TrafficSteeringRulesActionForward) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.TrafficSteeringRulesActionForward")
	diags := diag.Diagnostics{}
	var data []models.TrafficSteeringRulesActionForward

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.TrafficSteeringRulesActionForward
		obj, d := packTrafficSteeringRulesActionForwardFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.TrafficSteeringRulesActionForward{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.TrafficSteeringRulesActionForward", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.TrafficSteeringRulesActionForward{}.AttrType(), data)
}

// --- Unpacker for TrafficSteeringRulesActionForwardForward ---
func unpackTrafficSteeringRulesActionForwardForwardToSdk(ctx context.Context, obj types.Object) (*deployment_services.TrafficSteeringRulesActionForwardForward, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.TrafficSteeringRulesActionForwardForward", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.TrafficSteeringRulesActionForwardForward
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk deployment_services.TrafficSteeringRulesActionForwardForward
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Target.IsNull() && !model.Target.IsUnknown() {
		sdk.Target = model.Target.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Target", "value": *sdk.Target})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.TrafficSteeringRulesActionForwardForward", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for TrafficSteeringRulesActionForwardForward ---
func packTrafficSteeringRulesActionForwardForwardFromSdk(ctx context.Context, sdk deployment_services.TrafficSteeringRulesActionForwardForward) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.TrafficSteeringRulesActionForwardForward", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.TrafficSteeringRulesActionForwardForward
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Target != nil {
		model.Target = basetypes.NewStringValue(*sdk.Target)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Target", "value": *sdk.Target})
	} else {
		model.Target = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.TrafficSteeringRulesActionForwardForward{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.TrafficSteeringRulesActionForwardForward", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for TrafficSteeringRulesActionForwardForward ---
func unpackTrafficSteeringRulesActionForwardForwardListToSdk(ctx context.Context, list types.List) ([]deployment_services.TrafficSteeringRulesActionForwardForward, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.TrafficSteeringRulesActionForwardForward")
	diags := diag.Diagnostics{}
	var data []models.TrafficSteeringRulesActionForwardForward
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]deployment_services.TrafficSteeringRulesActionForwardForward, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.TrafficSteeringRulesActionForwardForward{}.AttrTypes(), &item)
		unpacked, d := unpackTrafficSteeringRulesActionForwardForwardToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.TrafficSteeringRulesActionForwardForward", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for TrafficSteeringRulesActionForwardForward ---
func packTrafficSteeringRulesActionForwardForwardListFromSdk(ctx context.Context, sdks []deployment_services.TrafficSteeringRulesActionForwardForward) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.TrafficSteeringRulesActionForwardForward")
	diags := diag.Diagnostics{}
	var data []models.TrafficSteeringRulesActionForwardForward

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.TrafficSteeringRulesActionForwardForward
		obj, d := packTrafficSteeringRulesActionForwardForwardFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.TrafficSteeringRulesActionForwardForward{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.TrafficSteeringRulesActionForwardForward", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.TrafficSteeringRulesActionForwardForward{}.AttrType(), data)
}
