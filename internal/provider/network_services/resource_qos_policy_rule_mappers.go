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

// --- Unpacker for QosPolicyRules ---
func unpackQosPolicyRulesToSdk(ctx context.Context, obj types.Object) (*network_services.QosPolicyRules, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.QosPolicyRules", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.QosPolicyRules
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.QosPolicyRules
	var d diag.Diagnostics
	// Handling Objects
	if !model.Action.IsNull() && !model.Action.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Action")
		unpacked, d := unpackQosPolicyRulesActionToSdk(ctx, model.Action)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Action"})
		}
		if unpacked != nil {
			sdk.Action = *unpacked
		}
	}

	// Handling Primitives
	if !model.Description.IsNull() && !model.Description.IsUnknown() {
		sdk.Description = model.Description.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Description", "value": *sdk.Description})
	}

	// Handling Primitives
	if !model.Device.IsNull() && !model.Device.IsUnknown() {
		sdk.Device = model.Device.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Device", "value": *sdk.Device})
	}

	// Handling Objects
	if !model.DscpTos.IsNull() && !model.DscpTos.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field DscpTos")
		unpacked, d := unpackQosPolicyRulesDscpTosToSdk(ctx, model.DscpTos)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "DscpTos"})
		}
		if unpacked != nil {
			sdk.DscpTos = unpacked
		}
	}

	// Handling Primitives
	if !model.Folder.IsNull() && !model.Folder.IsUnknown() {
		sdk.Folder = model.Folder.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Folder", "value": *sdk.Folder})
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

	// Handling Primitives
	if !model.Schedule.IsNull() && !model.Schedule.IsUnknown() {
		sdk.Schedule = model.Schedule.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Schedule", "value": *sdk.Schedule})
	}

	// Handling Primitives
	if !model.Snippet.IsNull() && !model.Snippet.IsUnknown() {
		sdk.Snippet = model.Snippet.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.QosPolicyRules", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for QosPolicyRules ---
func packQosPolicyRulesFromSdk(ctx context.Context, sdk network_services.QosPolicyRules) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.QosPolicyRules", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.QosPolicyRules
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	// Logic for non-pointer / value-type nested objects
	if !reflect.ValueOf(sdk.Action).IsZero() {
		tflog.Debug(ctx, "Packing nested object for field Action")
		packed, d := packQosPolicyRulesActionFromSdk(ctx, sdk.Action)
		diags.Append(d...)
		model.Action = packed
	} else {
		model.Action = basetypes.NewObjectNull(models.QosPolicyRulesAction{}.AttrTypes())
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
	if sdk.Device != nil {
		model.Device = basetypes.NewStringValue(*sdk.Device)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Device", "value": *sdk.Device})
	} else {
		model.Device = basetypes.NewStringNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.DscpTos != nil {
		tflog.Debug(ctx, "Packing nested object for field DscpTos")
		packed, d := packQosPolicyRulesDscpTosFromSdk(ctx, *sdk.DscpTos)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "DscpTos"})
		}
		model.DscpTos = packed
	} else {
		model.DscpTos = basetypes.NewObjectNull(models.QosPolicyRulesDscpTos{}.AttrTypes())
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
	// Handling Primitives
	// Standard primitive packing
	if sdk.Schedule != nil {
		model.Schedule = basetypes.NewStringValue(*sdk.Schedule)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Schedule", "value": *sdk.Schedule})
	} else {
		model.Schedule = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Snippet != nil {
		model.Snippet = basetypes.NewStringValue(*sdk.Snippet)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	} else {
		model.Snippet = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.QosPolicyRules{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.QosPolicyRules", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for QosPolicyRules ---
func unpackQosPolicyRulesListToSdk(ctx context.Context, list types.List) ([]network_services.QosPolicyRules, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.QosPolicyRules")
	diags := diag.Diagnostics{}
	var data []models.QosPolicyRules
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.QosPolicyRules, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.QosPolicyRules{}.AttrTypes(), &item)
		unpacked, d := unpackQosPolicyRulesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.QosPolicyRules", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for QosPolicyRules ---
func packQosPolicyRulesListFromSdk(ctx context.Context, sdks []network_services.QosPolicyRules) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.QosPolicyRules")
	diags := diag.Diagnostics{}
	var data []models.QosPolicyRules

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.QosPolicyRules
		obj, d := packQosPolicyRulesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.QosPolicyRules{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.QosPolicyRules", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.QosPolicyRules{}.AttrType(), data)
}

// --- Unpacker for QosPolicyRulesAction ---
func unpackQosPolicyRulesActionToSdk(ctx context.Context, obj types.Object) (*network_services.QosPolicyRulesAction, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.QosPolicyRulesAction", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.QosPolicyRulesAction
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.QosPolicyRulesAction
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Class.IsNull() && !model.Class.IsUnknown() {
		sdk.Class = model.Class.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Class", "value": *sdk.Class})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.QosPolicyRulesAction", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for QosPolicyRulesAction ---
func packQosPolicyRulesActionFromSdk(ctx context.Context, sdk network_services.QosPolicyRulesAction) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.QosPolicyRulesAction", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.QosPolicyRulesAction
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Class != nil {
		model.Class = basetypes.NewStringValue(*sdk.Class)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Class", "value": *sdk.Class})
	} else {
		model.Class = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.QosPolicyRulesAction{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.QosPolicyRulesAction", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for QosPolicyRulesAction ---
func unpackQosPolicyRulesActionListToSdk(ctx context.Context, list types.List) ([]network_services.QosPolicyRulesAction, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.QosPolicyRulesAction")
	diags := diag.Diagnostics{}
	var data []models.QosPolicyRulesAction
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.QosPolicyRulesAction, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.QosPolicyRulesAction{}.AttrTypes(), &item)
		unpacked, d := unpackQosPolicyRulesActionToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.QosPolicyRulesAction", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for QosPolicyRulesAction ---
func packQosPolicyRulesActionListFromSdk(ctx context.Context, sdks []network_services.QosPolicyRulesAction) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.QosPolicyRulesAction")
	diags := diag.Diagnostics{}
	var data []models.QosPolicyRulesAction

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.QosPolicyRulesAction
		obj, d := packQosPolicyRulesActionFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.QosPolicyRulesAction{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.QosPolicyRulesAction", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.QosPolicyRulesAction{}.AttrType(), data)
}

// --- Unpacker for QosPolicyRulesDscpTos ---
func unpackQosPolicyRulesDscpTosToSdk(ctx context.Context, obj types.Object) (*network_services.QosPolicyRulesDscpTos, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.QosPolicyRulesDscpTos", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.QosPolicyRulesDscpTos
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.QosPolicyRulesDscpTos
	var d diag.Diagnostics
	// Handling Lists
	if !model.Codepoints.IsNull() && !model.Codepoints.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field Codepoints")
		unpacked, d := unpackQosPolicyRulesDscpTosCodepointsInnerListToSdk(ctx, model.Codepoints)
		diags.Append(d...)
		sdk.Codepoints = unpacked
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.QosPolicyRulesDscpTos", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for QosPolicyRulesDscpTos ---
func packQosPolicyRulesDscpTosFromSdk(ctx context.Context, sdk network_services.QosPolicyRulesDscpTos) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.QosPolicyRulesDscpTos", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.QosPolicyRulesDscpTos
	var d diag.Diagnostics
	// Handling Lists
	if sdk.Codepoints != nil {
		tflog.Debug(ctx, "Packing list of objects for field Codepoints")
		packed, d := packQosPolicyRulesDscpTosCodepointsInnerListFromSdk(ctx, sdk.Codepoints)
		diags.Append(d...)
		model.Codepoints = packed
	} else {
		model.Codepoints = basetypes.NewListNull(models.QosPolicyRulesDscpTosCodepointsInner{}.AttrType())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.QosPolicyRulesDscpTos{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.QosPolicyRulesDscpTos", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for QosPolicyRulesDscpTos ---
func unpackQosPolicyRulesDscpTosListToSdk(ctx context.Context, list types.List) ([]network_services.QosPolicyRulesDscpTos, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.QosPolicyRulesDscpTos")
	diags := diag.Diagnostics{}
	var data []models.QosPolicyRulesDscpTos
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.QosPolicyRulesDscpTos, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.QosPolicyRulesDscpTos{}.AttrTypes(), &item)
		unpacked, d := unpackQosPolicyRulesDscpTosToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.QosPolicyRulesDscpTos", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for QosPolicyRulesDscpTos ---
func packQosPolicyRulesDscpTosListFromSdk(ctx context.Context, sdks []network_services.QosPolicyRulesDscpTos) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.QosPolicyRulesDscpTos")
	diags := diag.Diagnostics{}
	var data []models.QosPolicyRulesDscpTos

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.QosPolicyRulesDscpTos
		obj, d := packQosPolicyRulesDscpTosFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.QosPolicyRulesDscpTos{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.QosPolicyRulesDscpTos", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.QosPolicyRulesDscpTos{}.AttrType(), data)
}

// --- Unpacker for QosPolicyRulesDscpTosCodepointsInner ---
func unpackQosPolicyRulesDscpTosCodepointsInnerToSdk(ctx context.Context, obj types.Object) (*network_services.QosPolicyRulesDscpTosCodepointsInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.QosPolicyRulesDscpTosCodepointsInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.QosPolicyRulesDscpTosCodepointsInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.QosPolicyRulesDscpTosCodepointsInner
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
	}

	// Handling Objects
	if !model.Type.IsNull() && !model.Type.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Type")
		unpacked, d := unpackQosPolicyRulesDscpTosCodepointsInnerTypeToSdk(ctx, model.Type)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Type"})
		}
		if unpacked != nil {
			sdk.Type = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.QosPolicyRulesDscpTosCodepointsInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for QosPolicyRulesDscpTosCodepointsInner ---
func packQosPolicyRulesDscpTosCodepointsInnerFromSdk(ctx context.Context, sdk network_services.QosPolicyRulesDscpTosCodepointsInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.QosPolicyRulesDscpTosCodepointsInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.QosPolicyRulesDscpTosCodepointsInner
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Name != nil {
		model.Name = basetypes.NewStringValue(*sdk.Name)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
	} else {
		model.Name = basetypes.NewStringNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Type != nil {
		tflog.Debug(ctx, "Packing nested object for field Type")
		packed, d := packQosPolicyRulesDscpTosCodepointsInnerTypeFromSdk(ctx, *sdk.Type)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Type"})
		}
		model.Type = packed
	} else {
		model.Type = basetypes.NewObjectNull(models.QosPolicyRulesDscpTosCodepointsInnerType{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.QosPolicyRulesDscpTosCodepointsInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.QosPolicyRulesDscpTosCodepointsInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for QosPolicyRulesDscpTosCodepointsInner ---
func unpackQosPolicyRulesDscpTosCodepointsInnerListToSdk(ctx context.Context, list types.List) ([]network_services.QosPolicyRulesDscpTosCodepointsInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.QosPolicyRulesDscpTosCodepointsInner")
	diags := diag.Diagnostics{}
	var data []models.QosPolicyRulesDscpTosCodepointsInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.QosPolicyRulesDscpTosCodepointsInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.QosPolicyRulesDscpTosCodepointsInner{}.AttrTypes(), &item)
		unpacked, d := unpackQosPolicyRulesDscpTosCodepointsInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.QosPolicyRulesDscpTosCodepointsInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for QosPolicyRulesDscpTosCodepointsInner ---
func packQosPolicyRulesDscpTosCodepointsInnerListFromSdk(ctx context.Context, sdks []network_services.QosPolicyRulesDscpTosCodepointsInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.QosPolicyRulesDscpTosCodepointsInner")
	diags := diag.Diagnostics{}
	var data []models.QosPolicyRulesDscpTosCodepointsInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.QosPolicyRulesDscpTosCodepointsInner
		obj, d := packQosPolicyRulesDscpTosCodepointsInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.QosPolicyRulesDscpTosCodepointsInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.QosPolicyRulesDscpTosCodepointsInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.QosPolicyRulesDscpTosCodepointsInner{}.AttrType(), data)
}

// --- Unpacker for QosPolicyRulesDscpTosCodepointsInnerType ---
func unpackQosPolicyRulesDscpTosCodepointsInnerTypeToSdk(ctx context.Context, obj types.Object) (*network_services.QosPolicyRulesDscpTosCodepointsInnerType, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.QosPolicyRulesDscpTosCodepointsInnerType", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.QosPolicyRulesDscpTosCodepointsInnerType
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.QosPolicyRulesDscpTosCodepointsInnerType
	var d diag.Diagnostics
	// Handling Objects
	if !model.Af.IsNull() && !model.Af.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Af")
		unpacked, d := unpackQosPolicyRulesDscpTosCodepointsInnerTypeAfToSdk(ctx, model.Af)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Af"})
		}
		if unpacked != nil {
			sdk.Af = unpacked
		}
	}

	// Handling Objects
	if !model.Cs.IsNull() && !model.Cs.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Cs")
		unpacked, d := unpackQosPolicyRulesDscpTosCodepointsInnerTypeAfToSdk(ctx, model.Cs)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Cs"})
		}
		if unpacked != nil {
			sdk.Cs = unpacked
		}
	}

	// Handling Objects
	if !model.Custom.IsNull() && !model.Custom.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Custom")
		unpacked, d := unpackQosPolicyRulesDscpTosCodepointsInnerTypeCustomToSdk(ctx, model.Custom)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Custom"})
		}
		if unpacked != nil {
			sdk.Custom = unpacked
		}
	}

	// Handling Typeless Objects
	if !model.Ef.IsNull() && !model.Ef.IsUnknown() {
		tflog.Debug(ctx, "Unpacking typeless object for field Ef")
		sdk.Ef = make(map[string]interface{})
	}

	// Handling Objects
	if !model.Tos.IsNull() && !model.Tos.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Tos")
		unpacked, d := unpackQosPolicyRulesDscpTosCodepointsInnerTypeAfToSdk(ctx, model.Tos)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Tos"})
		}
		if unpacked != nil {
			sdk.Tos = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.QosPolicyRulesDscpTosCodepointsInnerType", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for QosPolicyRulesDscpTosCodepointsInnerType ---
func packQosPolicyRulesDscpTosCodepointsInnerTypeFromSdk(ctx context.Context, sdk network_services.QosPolicyRulesDscpTosCodepointsInnerType) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.QosPolicyRulesDscpTosCodepointsInnerType", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.QosPolicyRulesDscpTosCodepointsInnerType
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Af != nil {
		tflog.Debug(ctx, "Packing nested object for field Af")
		packed, d := packQosPolicyRulesDscpTosCodepointsInnerTypeAfFromSdk(ctx, *sdk.Af)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Af"})
		}
		model.Af = packed
	} else {
		model.Af = basetypes.NewObjectNull(models.QosPolicyRulesDscpTosCodepointsInnerTypeAf{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Cs != nil {
		tflog.Debug(ctx, "Packing nested object for field Cs")
		packed, d := packQosPolicyRulesDscpTosCodepointsInnerTypeAfFromSdk(ctx, *sdk.Cs)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Cs"})
		}
		model.Cs = packed
	} else {
		model.Cs = basetypes.NewObjectNull(models.QosPolicyRulesDscpTosCodepointsInnerTypeAf{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Custom != nil {
		tflog.Debug(ctx, "Packing nested object for field Custom")
		packed, d := packQosPolicyRulesDscpTosCodepointsInnerTypeCustomFromSdk(ctx, *sdk.Custom)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Custom"})
		}
		model.Custom = packed
	} else {
		model.Custom = basetypes.NewObjectNull(models.QosPolicyRulesDscpTosCodepointsInnerTypeCustom{}.AttrTypes())
	}
	// Handling Objects
	// This is a marker object (e.g. CHAP: {}). We just need to create an empty, non-null object.
	if sdk.Ef != nil && !reflect.ValueOf(sdk.Ef).IsNil() {
		tflog.Debug(ctx, "Packing typeless object for field Ef")
		var d diag.Diagnostics
		// Create an empty object with no attributes, which signifies its presence.
		model.Ef, d = basetypes.NewObjectValue(map[string]attr.Type{}, map[string]attr.Value{})
		diags.Append(d...)
	} else {
		// Since this field is part of a oneOf, being nil means it's not selected.
		// We make the object null with an empty attribute map.
		model.Ef = basetypes.NewObjectNull(map[string]attr.Type{})
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Tos != nil {
		tflog.Debug(ctx, "Packing nested object for field Tos")
		packed, d := packQosPolicyRulesDscpTosCodepointsInnerTypeAfFromSdk(ctx, *sdk.Tos)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Tos"})
		}
		model.Tos = packed
	} else {
		model.Tos = basetypes.NewObjectNull(models.QosPolicyRulesDscpTosCodepointsInnerTypeAf{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.QosPolicyRulesDscpTosCodepointsInnerType{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.QosPolicyRulesDscpTosCodepointsInnerType", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for QosPolicyRulesDscpTosCodepointsInnerType ---
func unpackQosPolicyRulesDscpTosCodepointsInnerTypeListToSdk(ctx context.Context, list types.List) ([]network_services.QosPolicyRulesDscpTosCodepointsInnerType, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.QosPolicyRulesDscpTosCodepointsInnerType")
	diags := diag.Diagnostics{}
	var data []models.QosPolicyRulesDscpTosCodepointsInnerType
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.QosPolicyRulesDscpTosCodepointsInnerType, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.QosPolicyRulesDscpTosCodepointsInnerType{}.AttrTypes(), &item)
		unpacked, d := unpackQosPolicyRulesDscpTosCodepointsInnerTypeToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.QosPolicyRulesDscpTosCodepointsInnerType", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for QosPolicyRulesDscpTosCodepointsInnerType ---
func packQosPolicyRulesDscpTosCodepointsInnerTypeListFromSdk(ctx context.Context, sdks []network_services.QosPolicyRulesDscpTosCodepointsInnerType) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.QosPolicyRulesDscpTosCodepointsInnerType")
	diags := diag.Diagnostics{}
	var data []models.QosPolicyRulesDscpTosCodepointsInnerType

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.QosPolicyRulesDscpTosCodepointsInnerType
		obj, d := packQosPolicyRulesDscpTosCodepointsInnerTypeFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.QosPolicyRulesDscpTosCodepointsInnerType{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.QosPolicyRulesDscpTosCodepointsInnerType", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.QosPolicyRulesDscpTosCodepointsInnerType{}.AttrType(), data)
}

// --- Unpacker for QosPolicyRulesDscpTosCodepointsInnerTypeAf ---
func unpackQosPolicyRulesDscpTosCodepointsInnerTypeAfToSdk(ctx context.Context, obj types.Object) (*network_services.QosPolicyRulesDscpTosCodepointsInnerTypeAf, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.QosPolicyRulesDscpTosCodepointsInnerTypeAf", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.QosPolicyRulesDscpTosCodepointsInnerTypeAf
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.QosPolicyRulesDscpTosCodepointsInnerTypeAf
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Codepoint.IsNull() && !model.Codepoint.IsUnknown() {
		sdk.Codepoint = model.Codepoint.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Codepoint", "value": *sdk.Codepoint})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.QosPolicyRulesDscpTosCodepointsInnerTypeAf", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for QosPolicyRulesDscpTosCodepointsInnerTypeAf ---
func packQosPolicyRulesDscpTosCodepointsInnerTypeAfFromSdk(ctx context.Context, sdk network_services.QosPolicyRulesDscpTosCodepointsInnerTypeAf) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.QosPolicyRulesDscpTosCodepointsInnerTypeAf", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.QosPolicyRulesDscpTosCodepointsInnerTypeAf
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Codepoint != nil {
		model.Codepoint = basetypes.NewStringValue(*sdk.Codepoint)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Codepoint", "value": *sdk.Codepoint})
	} else {
		model.Codepoint = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.QosPolicyRulesDscpTosCodepointsInnerTypeAf{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.QosPolicyRulesDscpTosCodepointsInnerTypeAf", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for QosPolicyRulesDscpTosCodepointsInnerTypeAf ---
func unpackQosPolicyRulesDscpTosCodepointsInnerTypeAfListToSdk(ctx context.Context, list types.List) ([]network_services.QosPolicyRulesDscpTosCodepointsInnerTypeAf, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.QosPolicyRulesDscpTosCodepointsInnerTypeAf")
	diags := diag.Diagnostics{}
	var data []models.QosPolicyRulesDscpTosCodepointsInnerTypeAf
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.QosPolicyRulesDscpTosCodepointsInnerTypeAf, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.QosPolicyRulesDscpTosCodepointsInnerTypeAf{}.AttrTypes(), &item)
		unpacked, d := unpackQosPolicyRulesDscpTosCodepointsInnerTypeAfToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.QosPolicyRulesDscpTosCodepointsInnerTypeAf", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for QosPolicyRulesDscpTosCodepointsInnerTypeAf ---
func packQosPolicyRulesDscpTosCodepointsInnerTypeAfListFromSdk(ctx context.Context, sdks []network_services.QosPolicyRulesDscpTosCodepointsInnerTypeAf) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.QosPolicyRulesDscpTosCodepointsInnerTypeAf")
	diags := diag.Diagnostics{}
	var data []models.QosPolicyRulesDscpTosCodepointsInnerTypeAf

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.QosPolicyRulesDscpTosCodepointsInnerTypeAf
		obj, d := packQosPolicyRulesDscpTosCodepointsInnerTypeAfFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.QosPolicyRulesDscpTosCodepointsInnerTypeAf{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.QosPolicyRulesDscpTosCodepointsInnerTypeAf", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.QosPolicyRulesDscpTosCodepointsInnerTypeAf{}.AttrType(), data)
}

// --- Unpacker for QosPolicyRulesDscpTosCodepointsInnerTypeCustom ---
func unpackQosPolicyRulesDscpTosCodepointsInnerTypeCustomToSdk(ctx context.Context, obj types.Object) (*network_services.QosPolicyRulesDscpTosCodepointsInnerTypeCustom, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.QosPolicyRulesDscpTosCodepointsInnerTypeCustom", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.QosPolicyRulesDscpTosCodepointsInnerTypeCustom
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.QosPolicyRulesDscpTosCodepointsInnerTypeCustom
	var d diag.Diagnostics
	// Handling Objects
	if !model.Codepoint.IsNull() && !model.Codepoint.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Codepoint")
		unpacked, d := unpackQosPolicyRulesDscpTosCodepointsInnerTypeCustomCodepointToSdk(ctx, model.Codepoint)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Codepoint"})
		}
		if unpacked != nil {
			sdk.Codepoint = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.QosPolicyRulesDscpTosCodepointsInnerTypeCustom", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for QosPolicyRulesDscpTosCodepointsInnerTypeCustom ---
func packQosPolicyRulesDscpTosCodepointsInnerTypeCustomFromSdk(ctx context.Context, sdk network_services.QosPolicyRulesDscpTosCodepointsInnerTypeCustom) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.QosPolicyRulesDscpTosCodepointsInnerTypeCustom", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.QosPolicyRulesDscpTosCodepointsInnerTypeCustom
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Codepoint != nil {
		tflog.Debug(ctx, "Packing nested object for field Codepoint")
		packed, d := packQosPolicyRulesDscpTosCodepointsInnerTypeCustomCodepointFromSdk(ctx, *sdk.Codepoint)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Codepoint"})
		}
		model.Codepoint = packed
	} else {
		model.Codepoint = basetypes.NewObjectNull(models.QosPolicyRulesDscpTosCodepointsInnerTypeCustomCodepoint{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.QosPolicyRulesDscpTosCodepointsInnerTypeCustom{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.QosPolicyRulesDscpTosCodepointsInnerTypeCustom", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for QosPolicyRulesDscpTosCodepointsInnerTypeCustom ---
func unpackQosPolicyRulesDscpTosCodepointsInnerTypeCustomListToSdk(ctx context.Context, list types.List) ([]network_services.QosPolicyRulesDscpTosCodepointsInnerTypeCustom, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.QosPolicyRulesDscpTosCodepointsInnerTypeCustom")
	diags := diag.Diagnostics{}
	var data []models.QosPolicyRulesDscpTosCodepointsInnerTypeCustom
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.QosPolicyRulesDscpTosCodepointsInnerTypeCustom, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.QosPolicyRulesDscpTosCodepointsInnerTypeCustom{}.AttrTypes(), &item)
		unpacked, d := unpackQosPolicyRulesDscpTosCodepointsInnerTypeCustomToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.QosPolicyRulesDscpTosCodepointsInnerTypeCustom", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for QosPolicyRulesDscpTosCodepointsInnerTypeCustom ---
func packQosPolicyRulesDscpTosCodepointsInnerTypeCustomListFromSdk(ctx context.Context, sdks []network_services.QosPolicyRulesDscpTosCodepointsInnerTypeCustom) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.QosPolicyRulesDscpTosCodepointsInnerTypeCustom")
	diags := diag.Diagnostics{}
	var data []models.QosPolicyRulesDscpTosCodepointsInnerTypeCustom

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.QosPolicyRulesDscpTosCodepointsInnerTypeCustom
		obj, d := packQosPolicyRulesDscpTosCodepointsInnerTypeCustomFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.QosPolicyRulesDscpTosCodepointsInnerTypeCustom{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.QosPolicyRulesDscpTosCodepointsInnerTypeCustom", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.QosPolicyRulesDscpTosCodepointsInnerTypeCustom{}.AttrType(), data)
}

// --- Unpacker for QosPolicyRulesDscpTosCodepointsInnerTypeCustomCodepoint ---
func unpackQosPolicyRulesDscpTosCodepointsInnerTypeCustomCodepointToSdk(ctx context.Context, obj types.Object) (*network_services.QosPolicyRulesDscpTosCodepointsInnerTypeCustomCodepoint, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.QosPolicyRulesDscpTosCodepointsInnerTypeCustomCodepoint", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.QosPolicyRulesDscpTosCodepointsInnerTypeCustomCodepoint
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.QosPolicyRulesDscpTosCodepointsInnerTypeCustomCodepoint
	var d diag.Diagnostics
	// Handling Primitives
	if !model.BinaryValue.IsNull() && !model.BinaryValue.IsUnknown() {
		sdk.BinaryValue = model.BinaryValue.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "BinaryValue", "value": *sdk.BinaryValue})
	}

	// Handling Primitives
	if !model.CodepointName.IsNull() && !model.CodepointName.IsUnknown() {
		sdk.CodepointName = model.CodepointName.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "CodepointName", "value": *sdk.CodepointName})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.QosPolicyRulesDscpTosCodepointsInnerTypeCustomCodepoint", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for QosPolicyRulesDscpTosCodepointsInnerTypeCustomCodepoint ---
func packQosPolicyRulesDscpTosCodepointsInnerTypeCustomCodepointFromSdk(ctx context.Context, sdk network_services.QosPolicyRulesDscpTosCodepointsInnerTypeCustomCodepoint) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.QosPolicyRulesDscpTosCodepointsInnerTypeCustomCodepoint", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.QosPolicyRulesDscpTosCodepointsInnerTypeCustomCodepoint
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.BinaryValue != nil {
		model.BinaryValue = basetypes.NewStringValue(*sdk.BinaryValue)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "BinaryValue", "value": *sdk.BinaryValue})
	} else {
		model.BinaryValue = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.CodepointName != nil {
		model.CodepointName = basetypes.NewStringValue(*sdk.CodepointName)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "CodepointName", "value": *sdk.CodepointName})
	} else {
		model.CodepointName = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.QosPolicyRulesDscpTosCodepointsInnerTypeCustomCodepoint{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.QosPolicyRulesDscpTosCodepointsInnerTypeCustomCodepoint", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for QosPolicyRulesDscpTosCodepointsInnerTypeCustomCodepoint ---
func unpackQosPolicyRulesDscpTosCodepointsInnerTypeCustomCodepointListToSdk(ctx context.Context, list types.List) ([]network_services.QosPolicyRulesDscpTosCodepointsInnerTypeCustomCodepoint, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.QosPolicyRulesDscpTosCodepointsInnerTypeCustomCodepoint")
	diags := diag.Diagnostics{}
	var data []models.QosPolicyRulesDscpTosCodepointsInnerTypeCustomCodepoint
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.QosPolicyRulesDscpTosCodepointsInnerTypeCustomCodepoint, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.QosPolicyRulesDscpTosCodepointsInnerTypeCustomCodepoint{}.AttrTypes(), &item)
		unpacked, d := unpackQosPolicyRulesDscpTosCodepointsInnerTypeCustomCodepointToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.QosPolicyRulesDscpTosCodepointsInnerTypeCustomCodepoint", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for QosPolicyRulesDscpTosCodepointsInnerTypeCustomCodepoint ---
func packQosPolicyRulesDscpTosCodepointsInnerTypeCustomCodepointListFromSdk(ctx context.Context, sdks []network_services.QosPolicyRulesDscpTosCodepointsInnerTypeCustomCodepoint) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.QosPolicyRulesDscpTosCodepointsInnerTypeCustomCodepoint")
	diags := diag.Diagnostics{}
	var data []models.QosPolicyRulesDscpTosCodepointsInnerTypeCustomCodepoint

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.QosPolicyRulesDscpTosCodepointsInnerTypeCustomCodepoint
		obj, d := packQosPolicyRulesDscpTosCodepointsInnerTypeCustomCodepointFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.QosPolicyRulesDscpTosCodepointsInnerTypeCustomCodepoint{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.QosPolicyRulesDscpTosCodepointsInnerTypeCustomCodepoint", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.QosPolicyRulesDscpTosCodepointsInnerTypeCustomCodepoint{}.AttrType(), data)
}
