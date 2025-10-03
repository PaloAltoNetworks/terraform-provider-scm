package provider

import (
	"context"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/security_services"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/security_services"
)

// --- Unpacker for DosProtectionRules ---
func unpackDosProtectionRulesToSdk(ctx context.Context, obj types.Object) (*security_services.DosProtectionRules, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.DosProtectionRules", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.DosProtectionRules
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.DosProtectionRules
	var d diag.Diagnostics
	// Handling Objects
	if !model.Action.IsNull() && !model.Action.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Action")
		unpacked, d := unpackDosProtectionRulesActionToSdk(ctx, model.Action)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Action"})
		}
		if unpacked != nil {
			sdk.Action = unpacked
		}
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

	// Handling Primitives
	if !model.Disabled.IsNull() && !model.Disabled.IsUnknown() {
		sdk.Disabled = model.Disabled.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Disabled", "value": *sdk.Disabled})
	}

	// Handling Primitives
	if !model.Folder.IsNull() && !model.Folder.IsUnknown() {
		sdk.Folder = model.Folder.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Folder", "value": *sdk.Folder})
	}

	// Handling Lists
	if !model.From.IsNull() && !model.From.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field From")
		diags.Append(model.From.ElementsAs(ctx, &sdk.From, false)...)
	}

	// Handling Primitives
	if !model.Id.IsNull() && !model.Id.IsUnknown() {
		sdk.Id = model.Id.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Id", "value": *sdk.Id})
	}

	// Handling Primitives
	if !model.LogSetting.IsNull() && !model.LogSetting.IsUnknown() {
		sdk.LogSetting = model.LogSetting.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "LogSetting", "value": *sdk.LogSetting})
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	// Handling Primitives
	if !model.Position.IsNull() && !model.Position.IsUnknown() {
		sdk.Position = model.Position.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Position", "value": *sdk.Position})
	}

	// Handling Objects
	if !model.Protection.IsNull() && !model.Protection.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Protection")
		unpacked, d := unpackDosProtectionRulesProtectionToSdk(ctx, model.Protection)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Protection"})
		}
		if unpacked != nil {
			sdk.Protection = unpacked
		}
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

	// Handling Lists
	if !model.To.IsNull() && !model.To.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field To")
		diags.Append(model.To.ElementsAs(ctx, &sdk.To, false)...)
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.DosProtectionRules", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for DosProtectionRules ---
func packDosProtectionRulesFromSdk(ctx context.Context, sdk security_services.DosProtectionRules) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.DosProtectionRules", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.DosProtectionRules
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Action != nil {
		tflog.Debug(ctx, "Packing nested object for field Action")
		packed, d := packDosProtectionRulesActionFromSdk(ctx, *sdk.Action)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Action"})
		}
		model.Action = packed
	} else {
		model.Action = basetypes.NewObjectNull(models.DosProtectionRulesAction{}.AttrTypes())
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
	// Handling Primitives
	// Standard primitive packing
	if sdk.Disabled != nil {
		model.Disabled = basetypes.NewBoolValue(*sdk.Disabled)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Disabled", "value": *sdk.Disabled})
	} else {
		model.Disabled = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Folder != nil {
		model.Folder = basetypes.NewStringValue(*sdk.Folder)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Folder", "value": *sdk.Folder})
	} else {
		model.Folder = basetypes.NewStringNull()
	}
	// Handling Lists
	if sdk.From != nil {
		tflog.Debug(ctx, "Packing list of primitives for field From")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.From, d = basetypes.NewListValueFrom(ctx, elemType, sdk.From)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.From = basetypes.NewListNull(elemType)
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
	if sdk.LogSetting != nil {
		model.LogSetting = basetypes.NewStringValue(*sdk.LogSetting)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "LogSetting", "value": *sdk.LogSetting})
	} else {
		model.LogSetting = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	// Handling Primitives
	// Standard primitive packing
	if sdk.Position != nil {
		model.Position = basetypes.NewStringValue(*sdk.Position)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Position", "value": *sdk.Position})
	} else {
		model.Position = basetypes.NewStringNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Protection != nil {
		tflog.Debug(ctx, "Packing nested object for field Protection")
		packed, d := packDosProtectionRulesProtectionFromSdk(ctx, *sdk.Protection)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Protection"})
		}
		model.Protection = packed
	} else {
		model.Protection = basetypes.NewObjectNull(models.DosProtectionRulesProtection{}.AttrTypes())
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
	// Handling Lists
	if sdk.To != nil {
		tflog.Debug(ctx, "Packing list of primitives for field To")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.To, d = basetypes.NewListValueFrom(ctx, elemType, sdk.To)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.To = basetypes.NewListNull(elemType)
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.DosProtectionRules{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.DosProtectionRules", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for DosProtectionRules ---
func unpackDosProtectionRulesListToSdk(ctx context.Context, list types.List) ([]security_services.DosProtectionRules, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.DosProtectionRules")
	diags := diag.Diagnostics{}
	var data []models.DosProtectionRules
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.DosProtectionRules, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.DosProtectionRules{}.AttrTypes(), &item)
		unpacked, d := unpackDosProtectionRulesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.DosProtectionRules", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for DosProtectionRules ---
func packDosProtectionRulesListFromSdk(ctx context.Context, sdks []security_services.DosProtectionRules) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.DosProtectionRules")
	diags := diag.Diagnostics{}
	var data []models.DosProtectionRules

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.DosProtectionRules
		obj, d := packDosProtectionRulesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.DosProtectionRules{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.DosProtectionRules", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.DosProtectionRules{}.AttrType(), data)
}

// --- Unpacker for DosProtectionRulesAction ---
func unpackDosProtectionRulesActionToSdk(ctx context.Context, obj types.Object) (*security_services.DosProtectionRulesAction, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.DosProtectionRulesAction", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.DosProtectionRulesAction
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.DosProtectionRulesAction
	var d diag.Diagnostics
	// Handling Typeless Objects
	if !model.Allow.IsNull() && !model.Allow.IsUnknown() {
		tflog.Debug(ctx, "Unpacking typeless object for field Allow")
		sdk.Allow = make(map[string]interface{})
	}

	// Handling Typeless Objects
	if !model.Deny.IsNull() && !model.Deny.IsUnknown() {
		tflog.Debug(ctx, "Unpacking typeless object for field Deny")
		sdk.Deny = make(map[string]interface{})
	}

	// Handling Typeless Objects
	if !model.Protect.IsNull() && !model.Protect.IsUnknown() {
		tflog.Debug(ctx, "Unpacking typeless object for field Protect")
		sdk.Protect = make(map[string]interface{})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.DosProtectionRulesAction", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for DosProtectionRulesAction ---
func packDosProtectionRulesActionFromSdk(ctx context.Context, sdk security_services.DosProtectionRulesAction) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.DosProtectionRulesAction", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.DosProtectionRulesAction
	var d diag.Diagnostics
	// Handling Objects
	// This is a marker object (e.g. CHAP: {}). We just need to create an empty, non-null object.
	if sdk.Allow != nil && !reflect.ValueOf(sdk.Allow).IsNil() {
		tflog.Debug(ctx, "Packing typeless object for field Allow")
		var d diag.Diagnostics
		// Create an empty object with no attributes, which signifies its presence.
		model.Allow, d = basetypes.NewObjectValue(map[string]attr.Type{}, map[string]attr.Value{})
		diags.Append(d...)
	} else {
		// Since this field is part of a oneOf, being nil means it's not selected.
		// We make the object null with an empty attribute map.
		model.Allow = basetypes.NewObjectNull(map[string]attr.Type{})
	}
	// Handling Objects
	// This is a marker object (e.g. CHAP: {}). We just need to create an empty, non-null object.
	if sdk.Deny != nil && !reflect.ValueOf(sdk.Deny).IsNil() {
		tflog.Debug(ctx, "Packing typeless object for field Deny")
		var d diag.Diagnostics
		// Create an empty object with no attributes, which signifies its presence.
		model.Deny, d = basetypes.NewObjectValue(map[string]attr.Type{}, map[string]attr.Value{})
		diags.Append(d...)
	} else {
		// Since this field is part of a oneOf, being nil means it's not selected.
		// We make the object null with an empty attribute map.
		model.Deny = basetypes.NewObjectNull(map[string]attr.Type{})
	}
	// Handling Objects
	// This is a marker object (e.g. CHAP: {}). We just need to create an empty, non-null object.
	if sdk.Protect != nil && !reflect.ValueOf(sdk.Protect).IsNil() {
		tflog.Debug(ctx, "Packing typeless object for field Protect")
		var d diag.Diagnostics
		// Create an empty object with no attributes, which signifies its presence.
		model.Protect, d = basetypes.NewObjectValue(map[string]attr.Type{}, map[string]attr.Value{})
		diags.Append(d...)
	} else {
		// Since this field is part of a oneOf, being nil means it's not selected.
		// We make the object null with an empty attribute map.
		model.Protect = basetypes.NewObjectNull(map[string]attr.Type{})
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.DosProtectionRulesAction{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.DosProtectionRulesAction", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for DosProtectionRulesAction ---
func unpackDosProtectionRulesActionListToSdk(ctx context.Context, list types.List) ([]security_services.DosProtectionRulesAction, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.DosProtectionRulesAction")
	diags := diag.Diagnostics{}
	var data []models.DosProtectionRulesAction
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.DosProtectionRulesAction, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.DosProtectionRulesAction{}.AttrTypes(), &item)
		unpacked, d := unpackDosProtectionRulesActionToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.DosProtectionRulesAction", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for DosProtectionRulesAction ---
func packDosProtectionRulesActionListFromSdk(ctx context.Context, sdks []security_services.DosProtectionRulesAction) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.DosProtectionRulesAction")
	diags := diag.Diagnostics{}
	var data []models.DosProtectionRulesAction

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.DosProtectionRulesAction
		obj, d := packDosProtectionRulesActionFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.DosProtectionRulesAction{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.DosProtectionRulesAction", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.DosProtectionRulesAction{}.AttrType(), data)
}

// --- Unpacker for DosProtectionRulesProtection ---
func unpackDosProtectionRulesProtectionToSdk(ctx context.Context, obj types.Object) (*security_services.DosProtectionRulesProtection, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.DosProtectionRulesProtection", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.DosProtectionRulesProtection
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.DosProtectionRulesProtection
	var d diag.Diagnostics
	// Handling Objects
	if !model.Aggregate.IsNull() && !model.Aggregate.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Aggregate")
		unpacked, d := unpackDosProtectionRulesProtectionAggregateToSdk(ctx, model.Aggregate)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Aggregate"})
		}
		if unpacked != nil {
			sdk.Aggregate = unpacked
		}
	}

	// Handling Objects
	if !model.Classified.IsNull() && !model.Classified.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Classified")
		unpacked, d := unpackDosProtectionRulesProtectionClassifiedToSdk(ctx, model.Classified)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Classified"})
		}
		if unpacked != nil {
			sdk.Classified = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.DosProtectionRulesProtection", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for DosProtectionRulesProtection ---
func packDosProtectionRulesProtectionFromSdk(ctx context.Context, sdk security_services.DosProtectionRulesProtection) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.DosProtectionRulesProtection", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.DosProtectionRulesProtection
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Aggregate != nil {
		tflog.Debug(ctx, "Packing nested object for field Aggregate")
		packed, d := packDosProtectionRulesProtectionAggregateFromSdk(ctx, *sdk.Aggregate)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Aggregate"})
		}
		model.Aggregate = packed
	} else {
		model.Aggregate = basetypes.NewObjectNull(models.DosProtectionRulesProtectionAggregate{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Classified != nil {
		tflog.Debug(ctx, "Packing nested object for field Classified")
		packed, d := packDosProtectionRulesProtectionClassifiedFromSdk(ctx, *sdk.Classified)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Classified"})
		}
		model.Classified = packed
	} else {
		model.Classified = basetypes.NewObjectNull(models.DosProtectionRulesProtectionClassified{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.DosProtectionRulesProtection{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.DosProtectionRulesProtection", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for DosProtectionRulesProtection ---
func unpackDosProtectionRulesProtectionListToSdk(ctx context.Context, list types.List) ([]security_services.DosProtectionRulesProtection, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.DosProtectionRulesProtection")
	diags := diag.Diagnostics{}
	var data []models.DosProtectionRulesProtection
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.DosProtectionRulesProtection, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.DosProtectionRulesProtection{}.AttrTypes(), &item)
		unpacked, d := unpackDosProtectionRulesProtectionToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.DosProtectionRulesProtection", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for DosProtectionRulesProtection ---
func packDosProtectionRulesProtectionListFromSdk(ctx context.Context, sdks []security_services.DosProtectionRulesProtection) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.DosProtectionRulesProtection")
	diags := diag.Diagnostics{}
	var data []models.DosProtectionRulesProtection

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.DosProtectionRulesProtection
		obj, d := packDosProtectionRulesProtectionFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.DosProtectionRulesProtection{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.DosProtectionRulesProtection", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.DosProtectionRulesProtection{}.AttrType(), data)
}

// --- Unpacker for DosProtectionRulesProtectionAggregate ---
func unpackDosProtectionRulesProtectionAggregateToSdk(ctx context.Context, obj types.Object) (*security_services.DosProtectionRulesProtectionAggregate, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.DosProtectionRulesProtectionAggregate", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.DosProtectionRulesProtectionAggregate
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.DosProtectionRulesProtectionAggregate
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Profile.IsNull() && !model.Profile.IsUnknown() {
		sdk.Profile = model.Profile.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Profile", "value": sdk.Profile})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.DosProtectionRulesProtectionAggregate", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for DosProtectionRulesProtectionAggregate ---
func packDosProtectionRulesProtectionAggregateFromSdk(ctx context.Context, sdk security_services.DosProtectionRulesProtectionAggregate) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.DosProtectionRulesProtectionAggregate", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.DosProtectionRulesProtectionAggregate
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.Profile = basetypes.NewStringValue(sdk.Profile)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Profile", "value": sdk.Profile})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.DosProtectionRulesProtectionAggregate{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.DosProtectionRulesProtectionAggregate", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for DosProtectionRulesProtectionAggregate ---
func unpackDosProtectionRulesProtectionAggregateListToSdk(ctx context.Context, list types.List) ([]security_services.DosProtectionRulesProtectionAggregate, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.DosProtectionRulesProtectionAggregate")
	diags := diag.Diagnostics{}
	var data []models.DosProtectionRulesProtectionAggregate
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.DosProtectionRulesProtectionAggregate, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.DosProtectionRulesProtectionAggregate{}.AttrTypes(), &item)
		unpacked, d := unpackDosProtectionRulesProtectionAggregateToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.DosProtectionRulesProtectionAggregate", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for DosProtectionRulesProtectionAggregate ---
func packDosProtectionRulesProtectionAggregateListFromSdk(ctx context.Context, sdks []security_services.DosProtectionRulesProtectionAggregate) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.DosProtectionRulesProtectionAggregate")
	diags := diag.Diagnostics{}
	var data []models.DosProtectionRulesProtectionAggregate

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.DosProtectionRulesProtectionAggregate
		obj, d := packDosProtectionRulesProtectionAggregateFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.DosProtectionRulesProtectionAggregate{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.DosProtectionRulesProtectionAggregate", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.DosProtectionRulesProtectionAggregate{}.AttrType(), data)
}

// --- Unpacker for DosProtectionRulesProtectionClassified ---
func unpackDosProtectionRulesProtectionClassifiedToSdk(ctx context.Context, obj types.Object) (*security_services.DosProtectionRulesProtectionClassified, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.DosProtectionRulesProtectionClassified", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.DosProtectionRulesProtectionClassified
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.DosProtectionRulesProtectionClassified
	var d diag.Diagnostics
	// Handling Objects
	if !model.ClassificationCriteria.IsNull() && !model.ClassificationCriteria.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field ClassificationCriteria")
		unpacked, d := unpackDosProtectionRulesProtectionClassifiedClassificationCriteriaToSdk(ctx, model.ClassificationCriteria)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "ClassificationCriteria"})
		}
		if unpacked != nil {
			sdk.ClassificationCriteria = *unpacked
		}
	}

	// Handling Primitives
	if !model.Profile.IsNull() && !model.Profile.IsUnknown() {
		sdk.Profile = model.Profile.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Profile", "value": sdk.Profile})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.DosProtectionRulesProtectionClassified", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for DosProtectionRulesProtectionClassified ---
func packDosProtectionRulesProtectionClassifiedFromSdk(ctx context.Context, sdk security_services.DosProtectionRulesProtectionClassified) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.DosProtectionRulesProtectionClassified", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.DosProtectionRulesProtectionClassified
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	// Logic for non-pointer / value-type nested objects
	if !reflect.ValueOf(sdk.ClassificationCriteria).IsZero() {
		tflog.Debug(ctx, "Packing nested object for field ClassificationCriteria")
		packed, d := packDosProtectionRulesProtectionClassifiedClassificationCriteriaFromSdk(ctx, sdk.ClassificationCriteria)
		diags.Append(d...)
		model.ClassificationCriteria = packed
	} else {
		model.ClassificationCriteria = basetypes.NewObjectNull(models.DosProtectionRulesProtectionClassifiedClassificationCriteria{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	model.Profile = basetypes.NewStringValue(sdk.Profile)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Profile", "value": sdk.Profile})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.DosProtectionRulesProtectionClassified{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.DosProtectionRulesProtectionClassified", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for DosProtectionRulesProtectionClassified ---
func unpackDosProtectionRulesProtectionClassifiedListToSdk(ctx context.Context, list types.List) ([]security_services.DosProtectionRulesProtectionClassified, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.DosProtectionRulesProtectionClassified")
	diags := diag.Diagnostics{}
	var data []models.DosProtectionRulesProtectionClassified
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.DosProtectionRulesProtectionClassified, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.DosProtectionRulesProtectionClassified{}.AttrTypes(), &item)
		unpacked, d := unpackDosProtectionRulesProtectionClassifiedToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.DosProtectionRulesProtectionClassified", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for DosProtectionRulesProtectionClassified ---
func packDosProtectionRulesProtectionClassifiedListFromSdk(ctx context.Context, sdks []security_services.DosProtectionRulesProtectionClassified) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.DosProtectionRulesProtectionClassified")
	diags := diag.Diagnostics{}
	var data []models.DosProtectionRulesProtectionClassified

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.DosProtectionRulesProtectionClassified
		obj, d := packDosProtectionRulesProtectionClassifiedFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.DosProtectionRulesProtectionClassified{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.DosProtectionRulesProtectionClassified", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.DosProtectionRulesProtectionClassified{}.AttrType(), data)
}

// --- Unpacker for DosProtectionRulesProtectionClassifiedClassificationCriteria ---
func unpackDosProtectionRulesProtectionClassifiedClassificationCriteriaToSdk(ctx context.Context, obj types.Object) (*security_services.DosProtectionRulesProtectionClassifiedClassificationCriteria, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.DosProtectionRulesProtectionClassifiedClassificationCriteria", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.DosProtectionRulesProtectionClassifiedClassificationCriteria
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.DosProtectionRulesProtectionClassifiedClassificationCriteria
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Address.IsNull() && !model.Address.IsUnknown() {
		sdk.Address = model.Address.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Address", "value": sdk.Address})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.DosProtectionRulesProtectionClassifiedClassificationCriteria", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for DosProtectionRulesProtectionClassifiedClassificationCriteria ---
func packDosProtectionRulesProtectionClassifiedClassificationCriteriaFromSdk(ctx context.Context, sdk security_services.DosProtectionRulesProtectionClassifiedClassificationCriteria) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.DosProtectionRulesProtectionClassifiedClassificationCriteria", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.DosProtectionRulesProtectionClassifiedClassificationCriteria
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.Address = basetypes.NewStringValue(sdk.Address)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Address", "value": sdk.Address})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.DosProtectionRulesProtectionClassifiedClassificationCriteria{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.DosProtectionRulesProtectionClassifiedClassificationCriteria", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for DosProtectionRulesProtectionClassifiedClassificationCriteria ---
func unpackDosProtectionRulesProtectionClassifiedClassificationCriteriaListToSdk(ctx context.Context, list types.List) ([]security_services.DosProtectionRulesProtectionClassifiedClassificationCriteria, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.DosProtectionRulesProtectionClassifiedClassificationCriteria")
	diags := diag.Diagnostics{}
	var data []models.DosProtectionRulesProtectionClassifiedClassificationCriteria
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.DosProtectionRulesProtectionClassifiedClassificationCriteria, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.DosProtectionRulesProtectionClassifiedClassificationCriteria{}.AttrTypes(), &item)
		unpacked, d := unpackDosProtectionRulesProtectionClassifiedClassificationCriteriaToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.DosProtectionRulesProtectionClassifiedClassificationCriteria", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for DosProtectionRulesProtectionClassifiedClassificationCriteria ---
func packDosProtectionRulesProtectionClassifiedClassificationCriteriaListFromSdk(ctx context.Context, sdks []security_services.DosProtectionRulesProtectionClassifiedClassificationCriteria) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.DosProtectionRulesProtectionClassifiedClassificationCriteria")
	diags := diag.Diagnostics{}
	var data []models.DosProtectionRulesProtectionClassifiedClassificationCriteria

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.DosProtectionRulesProtectionClassifiedClassificationCriteria
		obj, d := packDosProtectionRulesProtectionClassifiedClassificationCriteriaFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.DosProtectionRulesProtectionClassifiedClassificationCriteria{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.DosProtectionRulesProtectionClassifiedClassificationCriteria", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.DosProtectionRulesProtectionClassifiedClassificationCriteria{}.AttrType(), data)
}
