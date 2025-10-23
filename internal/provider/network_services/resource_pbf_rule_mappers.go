package provider

import (
	"context"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/network_services"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/network_services"
)

// --- Unpacker for PbfRules ---
func unpackPbfRulesToSdk(ctx context.Context, obj types.Object) (*network_services.PbfRules, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.PbfRules", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.PbfRules
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.PbfRules
	var d diag.Diagnostics

	// Handling Objects
	if !model.Action.IsNull() && !model.Action.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Action")
		unpacked, d := unpackPbfRulesActionToSdk(ctx, model.Action)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Action"})
		}
		if unpacked != nil {
			sdk.Action = unpacked
		}
	}

	// Handling Lists
	if !model.Application.IsNull() && !model.Application.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Application")
		diags.Append(model.Application.ElementsAs(ctx, &sdk.Application, false)...)
	}

	// Handling Primitives
	if !model.Description.IsNull() && !model.Description.IsUnknown() {
		sdk.Description = model.Description.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Description", "value": *sdk.Description})
	}

	// Handling Lists
	if !model.Destination.IsNull() && !model.Destination.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Destination")
		diags.Append(model.Destination.ElementsAs(ctx, &sdk.Destination, false)...)
	}

	// Handling Primitives
	if !model.Device.IsNull() && !model.Device.IsUnknown() {
		sdk.Device = model.Device.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Device", "value": *sdk.Device})
	}

	// Handling Objects
	if !model.EnforceSymmetricReturn.IsNull() && !model.EnforceSymmetricReturn.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field EnforceSymmetricReturn")
		unpacked, d := unpackPbfRulesEnforceSymmetricReturnToSdk(ctx, model.EnforceSymmetricReturn)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "EnforceSymmetricReturn"})
		}
		if unpacked != nil {
			sdk.EnforceSymmetricReturn = unpacked
		}
	}

	// Handling Primitives
	if !model.Folder.IsNull() && !model.Folder.IsUnknown() {
		sdk.Folder = model.Folder.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Folder", "value": *sdk.Folder})
	}

	// Handling Objects
	if !model.From.IsNull() && !model.From.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field From")
		unpacked, d := unpackPbfRulesFromToSdk(ctx, model.From)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "From"})
		}
		if unpacked != nil {
			sdk.From = unpacked
		}
	}

	// Handling Primitives
	if !model.Id.IsNull() && !model.Id.IsUnknown() {
		sdk.Id = model.Id.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Id", "value": *sdk.Id})
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
	}

	// Handling Primitives
	if !model.Schedule.IsNull() && !model.Schedule.IsUnknown() {
		sdk.Schedule = model.Schedule.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Schedule", "value": *sdk.Schedule})
	}

	// Handling Lists
	if !model.Service.IsNull() && !model.Service.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Service")
		diags.Append(model.Service.ElementsAs(ctx, &sdk.Service, false)...)
	}

	// Handling Primitives
	if !model.Snippet.IsNull() && !model.Snippet.IsUnknown() {
		sdk.Snippet = model.Snippet.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
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

	// Handling Lists
	if !model.Tag.IsNull() && !model.Tag.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Tag")
		diags.Append(model.Tag.ElementsAs(ctx, &sdk.Tag, false)...)
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.PbfRules", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for PbfRules ---
func packPbfRulesFromSdk(ctx context.Context, sdk network_services.PbfRules) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.PbfRules", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.PbfRules
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Action != nil {
		tflog.Debug(ctx, "Packing nested object for field Action")
		packed, d := packPbfRulesActionFromSdk(ctx, *sdk.Action)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Action"})
		}
		model.Action = packed
	} else {
		model.Action = basetypes.NewObjectNull(models.PbfRulesAction{}.AttrTypes())
	}
	// Handling Lists
	if sdk.Application != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Application")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Application, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Application)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Application = basetypes.NewListNull(elemType)
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Description != nil {
		model.Description = basetypes.NewStringValue(*sdk.Description)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Description", "value": *sdk.Description})
	} else {
		model.Description = basetypes.NewStringNull()
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
	if sdk.Device != nil {
		model.Device = basetypes.NewStringValue(*sdk.Device)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Device", "value": *sdk.Device})
	} else {
		model.Device = basetypes.NewStringNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.EnforceSymmetricReturn != nil {
		tflog.Debug(ctx, "Packing nested object for field EnforceSymmetricReturn")
		packed, d := packPbfRulesEnforceSymmetricReturnFromSdk(ctx, *sdk.EnforceSymmetricReturn)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "EnforceSymmetricReturn"})
		}
		model.EnforceSymmetricReturn = packed
	} else {
		model.EnforceSymmetricReturn = basetypes.NewObjectNull(models.PbfRulesEnforceSymmetricReturn{}.AttrTypes())
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
	if sdk.From != nil {
		tflog.Debug(ctx, "Packing nested object for field From")
		packed, d := packPbfRulesFromFromSdk(ctx, *sdk.From)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "From"})
		}
		model.From = packed
	} else {
		model.From = basetypes.NewObjectNull(models.PbfRulesFrom{}.AttrTypes())
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
	if sdk.Name != nil {
		model.Name = basetypes.NewStringValue(*sdk.Name)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
	} else {
		model.Name = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Schedule != nil {
		model.Schedule = basetypes.NewStringValue(*sdk.Schedule)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Schedule", "value": *sdk.Schedule})
	} else {
		model.Schedule = basetypes.NewStringNull()
	}
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
	// Handling Primitives
	// Standard primitive packing
	if sdk.Snippet != nil {
		model.Snippet = basetypes.NewStringValue(*sdk.Snippet)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	} else {
		model.Snippet = basetypes.NewStringNull()
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
	// Handling Lists
	if sdk.Tag != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Tag")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Tag, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Tag)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Tag = basetypes.NewListNull(elemType)
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.PbfRules{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.PbfRules", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for PbfRules ---
func unpackPbfRulesListToSdk(ctx context.Context, list types.List) ([]network_services.PbfRules, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.PbfRules")
	diags := diag.Diagnostics{}
	var data []models.PbfRules
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.PbfRules, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.PbfRules{}.AttrTypes(), &item)
		unpacked, d := unpackPbfRulesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.PbfRules", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for PbfRules ---
func packPbfRulesListFromSdk(ctx context.Context, sdks []network_services.PbfRules) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.PbfRules")
	diags := diag.Diagnostics{}
	var data []models.PbfRules

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.PbfRules
		obj, d := packPbfRulesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.PbfRules{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.PbfRules", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.PbfRules{}.AttrType(), data)
}

// --- Unpacker for PbfRulesAction ---
func unpackPbfRulesActionToSdk(ctx context.Context, obj types.Object) (*network_services.PbfRulesAction, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.PbfRulesAction", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.PbfRulesAction
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.PbfRulesAction
	var d diag.Diagnostics
	// Handling Typeless Objects
	if !model.Discard.IsNull() && !model.Discard.IsUnknown() {
		tflog.Debug(ctx, "Unpacking typeless object for field Discard")
		sdk.Discard = make(map[string]interface{})
	}

	// Handling Objects
	if !model.Forward.IsNull() && !model.Forward.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Forward")
		unpacked, d := unpackPbfRulesActionForwardToSdk(ctx, model.Forward)
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

	tflog.Debug(ctx, "Exiting unpack helper for models.PbfRulesAction", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for PbfRulesAction ---
func packPbfRulesActionFromSdk(ctx context.Context, sdk network_services.PbfRulesAction) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.PbfRulesAction", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.PbfRulesAction
	var d diag.Diagnostics
	// Handling Objects
	// This is a marker object (e.g. CHAP: {}). We just need to create an empty, non-null object.
	if sdk.Discard != nil && !reflect.ValueOf(sdk.Discard).IsNil() {
		tflog.Debug(ctx, "Packing typeless object for field Discard")
		var d diag.Diagnostics
		// Create an empty object with no attributes, which signifies its presence.
		model.Discard, d = basetypes.NewObjectValue(map[string]attr.Type{}, map[string]attr.Value{})
		diags.Append(d...)
	} else {
		// Since this field is part of a oneOf, being nil means it's not selected.
		// We make the object null with an empty attribute map.
		model.Discard = basetypes.NewObjectNull(map[string]attr.Type{})
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Forward != nil {
		tflog.Debug(ctx, "Packing nested object for field Forward")
		packed, d := packPbfRulesActionForwardFromSdk(ctx, *sdk.Forward)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Forward"})
		}
		model.Forward = packed
	} else {
		model.Forward = basetypes.NewObjectNull(models.PbfRulesActionForward{}.AttrTypes())
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

	obj, d := types.ObjectValueFrom(ctx, models.PbfRulesAction{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.PbfRulesAction", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for PbfRulesAction ---
func unpackPbfRulesActionListToSdk(ctx context.Context, list types.List) ([]network_services.PbfRulesAction, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.PbfRulesAction")
	diags := diag.Diagnostics{}
	var data []models.PbfRulesAction
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.PbfRulesAction, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.PbfRulesAction{}.AttrTypes(), &item)
		unpacked, d := unpackPbfRulesActionToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.PbfRulesAction", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for PbfRulesAction ---
func packPbfRulesActionListFromSdk(ctx context.Context, sdks []network_services.PbfRulesAction) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.PbfRulesAction")
	diags := diag.Diagnostics{}
	var data []models.PbfRulesAction

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.PbfRulesAction
		obj, d := packPbfRulesActionFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.PbfRulesAction{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.PbfRulesAction", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.PbfRulesAction{}.AttrType(), data)
}

// --- Unpacker for PbfRulesActionForward ---
func unpackPbfRulesActionForwardToSdk(ctx context.Context, obj types.Object) (*network_services.PbfRulesActionForward, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.PbfRulesActionForward", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.PbfRulesActionForward
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.PbfRulesActionForward
	var d diag.Diagnostics
	// Handling Primitives
	if !model.EgressInterface.IsNull() && !model.EgressInterface.IsUnknown() {
		sdk.EgressInterface = model.EgressInterface.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "EgressInterface", "value": *sdk.EgressInterface})
	}

	// Handling Objects
	if !model.Monitor.IsNull() && !model.Monitor.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Monitor")
		unpacked, d := unpackPbfRulesActionForwardMonitorToSdk(ctx, model.Monitor)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Monitor"})
		}
		if unpacked != nil {
			sdk.Monitor = unpacked
		}
	}

	// Handling Objects
	if !model.Nexthop.IsNull() && !model.Nexthop.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Nexthop")
		unpacked, d := unpackPbfRulesActionForwardNexthopToSdk(ctx, model.Nexthop)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Nexthop"})
		}
		if unpacked != nil {
			sdk.Nexthop = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.PbfRulesActionForward", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for PbfRulesActionForward ---
func packPbfRulesActionForwardFromSdk(ctx context.Context, sdk network_services.PbfRulesActionForward) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.PbfRulesActionForward", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.PbfRulesActionForward
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.EgressInterface != nil {
		model.EgressInterface = basetypes.NewStringValue(*sdk.EgressInterface)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "EgressInterface", "value": *sdk.EgressInterface})
	} else {
		model.EgressInterface = basetypes.NewStringNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Monitor != nil {
		tflog.Debug(ctx, "Packing nested object for field Monitor")
		packed, d := packPbfRulesActionForwardMonitorFromSdk(ctx, *sdk.Monitor)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Monitor"})
		}
		model.Monitor = packed
	} else {
		model.Monitor = basetypes.NewObjectNull(models.PbfRulesActionForwardMonitor{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Nexthop != nil {
		tflog.Debug(ctx, "Packing nested object for field Nexthop")
		packed, d := packPbfRulesActionForwardNexthopFromSdk(ctx, *sdk.Nexthop)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Nexthop"})
		}
		model.Nexthop = packed
	} else {
		model.Nexthop = basetypes.NewObjectNull(models.PbfRulesActionForwardNexthop{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.PbfRulesActionForward{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.PbfRulesActionForward", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for PbfRulesActionForward ---
func unpackPbfRulesActionForwardListToSdk(ctx context.Context, list types.List) ([]network_services.PbfRulesActionForward, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.PbfRulesActionForward")
	diags := diag.Diagnostics{}
	var data []models.PbfRulesActionForward
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.PbfRulesActionForward, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.PbfRulesActionForward{}.AttrTypes(), &item)
		unpacked, d := unpackPbfRulesActionForwardToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.PbfRulesActionForward", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for PbfRulesActionForward ---
func packPbfRulesActionForwardListFromSdk(ctx context.Context, sdks []network_services.PbfRulesActionForward) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.PbfRulesActionForward")
	diags := diag.Diagnostics{}
	var data []models.PbfRulesActionForward

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.PbfRulesActionForward
		obj, d := packPbfRulesActionForwardFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.PbfRulesActionForward{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.PbfRulesActionForward", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.PbfRulesActionForward{}.AttrType(), data)
}

// --- Unpacker for PbfRulesActionForwardMonitor ---
func unpackPbfRulesActionForwardMonitorToSdk(ctx context.Context, obj types.Object) (*network_services.PbfRulesActionForwardMonitor, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.PbfRulesActionForwardMonitor", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.PbfRulesActionForwardMonitor
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.PbfRulesActionForwardMonitor
	var d diag.Diagnostics
	// Handling Primitives
	if !model.DisableIfUnreachable.IsNull() && !model.DisableIfUnreachable.IsUnknown() {
		sdk.DisableIfUnreachable = model.DisableIfUnreachable.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DisableIfUnreachable", "value": *sdk.DisableIfUnreachable})
	}

	// Handling Primitives
	if !model.IpAddress.IsNull() && !model.IpAddress.IsUnknown() {
		sdk.IpAddress = model.IpAddress.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "IpAddress", "value": *sdk.IpAddress})
	}

	// Handling Primitives
	if !model.Profile.IsNull() && !model.Profile.IsUnknown() {
		sdk.Profile = model.Profile.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Profile", "value": *sdk.Profile})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.PbfRulesActionForwardMonitor", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for PbfRulesActionForwardMonitor ---
func packPbfRulesActionForwardMonitorFromSdk(ctx context.Context, sdk network_services.PbfRulesActionForwardMonitor) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.PbfRulesActionForwardMonitor", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.PbfRulesActionForwardMonitor
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.DisableIfUnreachable != nil {
		model.DisableIfUnreachable = basetypes.NewBoolValue(*sdk.DisableIfUnreachable)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DisableIfUnreachable", "value": *sdk.DisableIfUnreachable})
	} else {
		model.DisableIfUnreachable = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.IpAddress != nil {
		model.IpAddress = basetypes.NewStringValue(*sdk.IpAddress)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "IpAddress", "value": *sdk.IpAddress})
	} else {
		model.IpAddress = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Profile != nil {
		model.Profile = basetypes.NewStringValue(*sdk.Profile)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Profile", "value": *sdk.Profile})
	} else {
		model.Profile = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.PbfRulesActionForwardMonitor{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.PbfRulesActionForwardMonitor", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for PbfRulesActionForwardMonitor ---
func unpackPbfRulesActionForwardMonitorListToSdk(ctx context.Context, list types.List) ([]network_services.PbfRulesActionForwardMonitor, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.PbfRulesActionForwardMonitor")
	diags := diag.Diagnostics{}
	var data []models.PbfRulesActionForwardMonitor
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.PbfRulesActionForwardMonitor, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.PbfRulesActionForwardMonitor{}.AttrTypes(), &item)
		unpacked, d := unpackPbfRulesActionForwardMonitorToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.PbfRulesActionForwardMonitor", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for PbfRulesActionForwardMonitor ---
func packPbfRulesActionForwardMonitorListFromSdk(ctx context.Context, sdks []network_services.PbfRulesActionForwardMonitor) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.PbfRulesActionForwardMonitor")
	diags := diag.Diagnostics{}
	var data []models.PbfRulesActionForwardMonitor

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.PbfRulesActionForwardMonitor
		obj, d := packPbfRulesActionForwardMonitorFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.PbfRulesActionForwardMonitor{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.PbfRulesActionForwardMonitor", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.PbfRulesActionForwardMonitor{}.AttrType(), data)
}

// --- Unpacker for PbfRulesActionForwardNexthop ---
func unpackPbfRulesActionForwardNexthopToSdk(ctx context.Context, obj types.Object) (*network_services.PbfRulesActionForwardNexthop, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.PbfRulesActionForwardNexthop", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.PbfRulesActionForwardNexthop
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.PbfRulesActionForwardNexthop
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Fqdn.IsNull() && !model.Fqdn.IsUnknown() {
		sdk.Fqdn = model.Fqdn.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Fqdn", "value": *sdk.Fqdn})
	}

	// Handling Primitives
	if !model.IpAddress.IsNull() && !model.IpAddress.IsUnknown() {
		sdk.IpAddress = model.IpAddress.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "IpAddress", "value": *sdk.IpAddress})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.PbfRulesActionForwardNexthop", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for PbfRulesActionForwardNexthop ---
func packPbfRulesActionForwardNexthopFromSdk(ctx context.Context, sdk network_services.PbfRulesActionForwardNexthop) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.PbfRulesActionForwardNexthop", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.PbfRulesActionForwardNexthop
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
	if sdk.IpAddress != nil {
		model.IpAddress = basetypes.NewStringValue(*sdk.IpAddress)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "IpAddress", "value": *sdk.IpAddress})
	} else {
		model.IpAddress = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.PbfRulesActionForwardNexthop{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.PbfRulesActionForwardNexthop", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for PbfRulesActionForwardNexthop ---
func unpackPbfRulesActionForwardNexthopListToSdk(ctx context.Context, list types.List) ([]network_services.PbfRulesActionForwardNexthop, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.PbfRulesActionForwardNexthop")
	diags := diag.Diagnostics{}
	var data []models.PbfRulesActionForwardNexthop
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.PbfRulesActionForwardNexthop, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.PbfRulesActionForwardNexthop{}.AttrTypes(), &item)
		unpacked, d := unpackPbfRulesActionForwardNexthopToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.PbfRulesActionForwardNexthop", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for PbfRulesActionForwardNexthop ---
func packPbfRulesActionForwardNexthopListFromSdk(ctx context.Context, sdks []network_services.PbfRulesActionForwardNexthop) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.PbfRulesActionForwardNexthop")
	diags := diag.Diagnostics{}
	var data []models.PbfRulesActionForwardNexthop

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.PbfRulesActionForwardNexthop
		obj, d := packPbfRulesActionForwardNexthopFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.PbfRulesActionForwardNexthop{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.PbfRulesActionForwardNexthop", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.PbfRulesActionForwardNexthop{}.AttrType(), data)
}

// --- Unpacker for PbfRulesEnforceSymmetricReturn ---
func unpackPbfRulesEnforceSymmetricReturnToSdk(ctx context.Context, obj types.Object) (*network_services.PbfRulesEnforceSymmetricReturn, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.PbfRulesEnforceSymmetricReturn", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.PbfRulesEnforceSymmetricReturn
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.PbfRulesEnforceSymmetricReturn
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Enabled.IsNull() && !model.Enabled.IsUnknown() {
		sdk.Enabled = model.Enabled.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Enabled", "value": *sdk.Enabled})
	}

	// Handling Lists
	if !model.NexthopAddressList.IsNull() && !model.NexthopAddressList.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field NexthopAddressList")
		unpacked, d := unpackPbfRulesEnforceSymmetricReturnNexthopAddressListInnerListToSdk(ctx, model.NexthopAddressList)
		diags.Append(d...)
		sdk.NexthopAddressList = unpacked
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.PbfRulesEnforceSymmetricReturn", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for PbfRulesEnforceSymmetricReturn ---
func packPbfRulesEnforceSymmetricReturnFromSdk(ctx context.Context, sdk network_services.PbfRulesEnforceSymmetricReturn) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.PbfRulesEnforceSymmetricReturn", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.PbfRulesEnforceSymmetricReturn
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Enabled != nil {
		model.Enabled = basetypes.NewBoolValue(*sdk.Enabled)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Enabled", "value": *sdk.Enabled})
	} else {
		model.Enabled = basetypes.NewBoolNull()
	}
	// Handling Lists
	if sdk.NexthopAddressList != nil {
		tflog.Debug(ctx, "Packing list of objects for field NexthopAddressList")
		packed, d := packPbfRulesEnforceSymmetricReturnNexthopAddressListInnerListFromSdk(ctx, sdk.NexthopAddressList)
		diags.Append(d...)
		model.NexthopAddressList = packed
	} else {
		model.NexthopAddressList = basetypes.NewListNull(models.PbfRulesEnforceSymmetricReturnNexthopAddressListInner{}.AttrType())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.PbfRulesEnforceSymmetricReturn{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.PbfRulesEnforceSymmetricReturn", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for PbfRulesEnforceSymmetricReturn ---
func unpackPbfRulesEnforceSymmetricReturnListToSdk(ctx context.Context, list types.List) ([]network_services.PbfRulesEnforceSymmetricReturn, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.PbfRulesEnforceSymmetricReturn")
	diags := diag.Diagnostics{}
	var data []models.PbfRulesEnforceSymmetricReturn
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.PbfRulesEnforceSymmetricReturn, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.PbfRulesEnforceSymmetricReturn{}.AttrTypes(), &item)
		unpacked, d := unpackPbfRulesEnforceSymmetricReturnToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.PbfRulesEnforceSymmetricReturn", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for PbfRulesEnforceSymmetricReturn ---
func packPbfRulesEnforceSymmetricReturnListFromSdk(ctx context.Context, sdks []network_services.PbfRulesEnforceSymmetricReturn) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.PbfRulesEnforceSymmetricReturn")
	diags := diag.Diagnostics{}
	var data []models.PbfRulesEnforceSymmetricReturn

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.PbfRulesEnforceSymmetricReturn
		obj, d := packPbfRulesEnforceSymmetricReturnFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.PbfRulesEnforceSymmetricReturn{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.PbfRulesEnforceSymmetricReturn", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.PbfRulesEnforceSymmetricReturn{}.AttrType(), data)
}

// --- Unpacker for PbfRulesEnforceSymmetricReturnNexthopAddressListInner ---
func unpackPbfRulesEnforceSymmetricReturnNexthopAddressListInnerToSdk(ctx context.Context, obj types.Object) (*network_services.PbfRulesEnforceSymmetricReturnNexthopAddressListInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.PbfRulesEnforceSymmetricReturnNexthopAddressListInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.PbfRulesEnforceSymmetricReturnNexthopAddressListInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.PbfRulesEnforceSymmetricReturnNexthopAddressListInner
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.PbfRulesEnforceSymmetricReturnNexthopAddressListInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for PbfRulesEnforceSymmetricReturnNexthopAddressListInner ---
func packPbfRulesEnforceSymmetricReturnNexthopAddressListInnerFromSdk(ctx context.Context, sdk network_services.PbfRulesEnforceSymmetricReturnNexthopAddressListInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.PbfRulesEnforceSymmetricReturnNexthopAddressListInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.PbfRulesEnforceSymmetricReturnNexthopAddressListInner
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Name != nil {
		model.Name = basetypes.NewStringValue(*sdk.Name)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
	} else {
		model.Name = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.PbfRulesEnforceSymmetricReturnNexthopAddressListInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.PbfRulesEnforceSymmetricReturnNexthopAddressListInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for PbfRulesEnforceSymmetricReturnNexthopAddressListInner ---
func unpackPbfRulesEnforceSymmetricReturnNexthopAddressListInnerListToSdk(ctx context.Context, list types.List) ([]network_services.PbfRulesEnforceSymmetricReturnNexthopAddressListInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.PbfRulesEnforceSymmetricReturnNexthopAddressListInner")
	diags := diag.Diagnostics{}
	var data []models.PbfRulesEnforceSymmetricReturnNexthopAddressListInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.PbfRulesEnforceSymmetricReturnNexthopAddressListInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.PbfRulesEnforceSymmetricReturnNexthopAddressListInner{}.AttrTypes(), &item)
		unpacked, d := unpackPbfRulesEnforceSymmetricReturnNexthopAddressListInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.PbfRulesEnforceSymmetricReturnNexthopAddressListInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for PbfRulesEnforceSymmetricReturnNexthopAddressListInner ---
func packPbfRulesEnforceSymmetricReturnNexthopAddressListInnerListFromSdk(ctx context.Context, sdks []network_services.PbfRulesEnforceSymmetricReturnNexthopAddressListInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.PbfRulesEnforceSymmetricReturnNexthopAddressListInner")
	diags := diag.Diagnostics{}
	var data []models.PbfRulesEnforceSymmetricReturnNexthopAddressListInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.PbfRulesEnforceSymmetricReturnNexthopAddressListInner
		obj, d := packPbfRulesEnforceSymmetricReturnNexthopAddressListInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.PbfRulesEnforceSymmetricReturnNexthopAddressListInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.PbfRulesEnforceSymmetricReturnNexthopAddressListInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.PbfRulesEnforceSymmetricReturnNexthopAddressListInner{}.AttrType(), data)
}

// --- Unpacker for PbfRulesFrom ---
func unpackPbfRulesFromToSdk(ctx context.Context, obj types.Object) (*network_services.PbfRulesFrom, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.PbfRulesFrom", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.PbfRulesFrom
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.PbfRulesFrom
	var d diag.Diagnostics
	// Handling Lists
	if !model.Interface.IsNull() && !model.Interface.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Interface")
		diags.Append(model.Interface.ElementsAs(ctx, &sdk.Interface, false)...)
	}

	// Handling Lists
	if !model.Zone.IsNull() && !model.Zone.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Zone")
		diags.Append(model.Zone.ElementsAs(ctx, &sdk.Zone, false)...)
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.PbfRulesFrom", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for PbfRulesFrom ---
func packPbfRulesFromFromSdk(ctx context.Context, sdk network_services.PbfRulesFrom) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.PbfRulesFrom", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.PbfRulesFrom
	var d diag.Diagnostics
	// Handling Lists
	if sdk.Interface != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Interface")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Interface, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Interface)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Interface = basetypes.NewListNull(elemType)
	}
	// Handling Lists
	if sdk.Zone != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Zone")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Zone, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Zone)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Zone = basetypes.NewListNull(elemType)
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.PbfRulesFrom{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.PbfRulesFrom", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for PbfRulesFrom ---
func unpackPbfRulesFromListToSdk(ctx context.Context, list types.List) ([]network_services.PbfRulesFrom, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.PbfRulesFrom")
	diags := diag.Diagnostics{}
	var data []models.PbfRulesFrom
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.PbfRulesFrom, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.PbfRulesFrom{}.AttrTypes(), &item)
		unpacked, d := unpackPbfRulesFromToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.PbfRulesFrom", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for PbfRulesFrom ---
func packPbfRulesFromListFromSdk(ctx context.Context, sdks []network_services.PbfRulesFrom) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.PbfRulesFrom")
	diags := diag.Diagnostics{}
	var data []models.PbfRulesFrom

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.PbfRulesFrom
		obj, d := packPbfRulesFromFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.PbfRulesFrom{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.PbfRulesFrom", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.PbfRulesFrom{}.AttrType(), data)
}
