package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/security_services"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/security_services"
)

// --- Unpacker for AppOverrideRules ---
func unpackAppOverrideRulesToSdk(ctx context.Context, obj types.Object) (*security_services.AppOverrideRules, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.AppOverrideRules", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.AppOverrideRules
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.AppOverrideRules
	var d diag.Diagnostics

	// Handling Primitives
	if !model.Application.IsNull() && !model.Application.IsUnknown() {
		sdk.Application = model.Application.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Application", "value": sdk.Application})
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
	if !model.GroupTag.IsNull() && !model.GroupTag.IsUnknown() {
		sdk.GroupTag = model.GroupTag.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "GroupTag", "value": *sdk.GroupTag})
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
	if !model.NegateDestination.IsNull() && !model.NegateDestination.IsUnknown() {
		sdk.NegateDestination = model.NegateDestination.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "NegateDestination", "value": *sdk.NegateDestination})
	}

	// Handling Primitives
	if !model.NegateSource.IsNull() && !model.NegateSource.IsUnknown() {
		sdk.NegateSource = model.NegateSource.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "NegateSource", "value": *sdk.NegateSource})
	}

	// Handling Primitives
	if !model.Port.IsNull() && !model.Port.IsUnknown() {
		sdk.Port = model.Port.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Port", "value": sdk.Port})
	}

	// Handling Primitives
	if !model.Protocol.IsNull() && !model.Protocol.IsUnknown() {
		sdk.Protocol = model.Protocol.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Protocol", "value": sdk.Protocol})
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

	tflog.Debug(ctx, "Exiting unpack helper for models.AppOverrideRules", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for AppOverrideRules ---
func packAppOverrideRulesFromSdk(ctx context.Context, sdk security_services.AppOverrideRules) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.AppOverrideRules", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.AppOverrideRules
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.Application = basetypes.NewStringValue(sdk.Application)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Application", "value": sdk.Application})
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
	if sdk.GroupTag != nil {
		model.GroupTag = basetypes.NewStringValue(*sdk.GroupTag)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "GroupTag", "value": *sdk.GroupTag})
	} else {
		model.GroupTag = basetypes.NewStringNull()
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
	if sdk.NegateDestination != nil {
		model.NegateDestination = basetypes.NewBoolValue(*sdk.NegateDestination)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "NegateDestination", "value": *sdk.NegateDestination})
	} else {
		model.NegateDestination = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.NegateSource != nil {
		model.NegateSource = basetypes.NewBoolValue(*sdk.NegateSource)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "NegateSource", "value": *sdk.NegateSource})
	} else {
		model.NegateSource = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	model.Port = basetypes.NewStringValue(sdk.Port)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Port", "value": sdk.Port})
	// Handling Primitives
	// Standard primitive packing
	model.Protocol = basetypes.NewStringValue(sdk.Protocol)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Protocol", "value": sdk.Protocol})
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

	obj, d := types.ObjectValueFrom(ctx, models.AppOverrideRules{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.AppOverrideRules", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for AppOverrideRules ---
func unpackAppOverrideRulesListToSdk(ctx context.Context, list types.List) ([]security_services.AppOverrideRules, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.AppOverrideRules")
	diags := diag.Diagnostics{}
	var data []models.AppOverrideRules
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.AppOverrideRules, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.AppOverrideRules{}.AttrTypes(), &item)
		unpacked, d := unpackAppOverrideRulesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.AppOverrideRules", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for AppOverrideRules ---
func packAppOverrideRulesListFromSdk(ctx context.Context, sdks []security_services.AppOverrideRules) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.AppOverrideRules")
	diags := diag.Diagnostics{}
	var data []models.AppOverrideRules

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.AppOverrideRules
		obj, d := packAppOverrideRulesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.AppOverrideRules{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.AppOverrideRules", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.AppOverrideRules{}.AttrType(), data)
}
