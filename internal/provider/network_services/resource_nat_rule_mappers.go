package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/network_services"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/network_services"
)

// --- Unpacker for NatRules ---
func unpackNatRulesToSdk(ctx context.Context, obj types.Object) (*network_services.NatRules, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.NatRules", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.NatRules
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.NatRules
	var d diag.Diagnostics

	// Handling Primitives
	if !model.ActiveActiveDeviceBinding.IsNull() && !model.ActiveActiveDeviceBinding.IsUnknown() {
		sdk.ActiveActiveDeviceBinding = model.ActiveActiveDeviceBinding.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "ActiveActiveDeviceBinding", "value": *sdk.ActiveActiveDeviceBinding})
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
	if !model.Distribution.IsNull() && !model.Distribution.IsUnknown() {
		sdk.Distribution = model.Distribution.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Distribution", "value": *sdk.Distribution})
	}

	// Handling Objects
	if !model.DnsRewrite.IsNull() && !model.DnsRewrite.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field DnsRewrite")
		unpacked, d := unpackNatRulesDnsRewriteToSdk(ctx, model.DnsRewrite)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "DnsRewrite"})
		}
		if unpacked != nil {
			sdk.DnsRewrite = unpacked
		}
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
		sdk.Id = model.Id.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Id", "value": sdk.Id})
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	// Handling Primitives
	if !model.NatType.IsNull() && !model.NatType.IsUnknown() {
		sdk.NatType = model.NatType.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "NatType", "value": *sdk.NatType})
	}

	// Handling Primitives
	if !model.Service.IsNull() && !model.Service.IsUnknown() {
		sdk.Service = model.Service.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Service", "value": sdk.Service})
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

	// Handling Objects
	if !model.SourceTranslation.IsNull() && !model.SourceTranslation.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field SourceTranslation")
		unpacked, d := unpackNatRulesSourceTranslationToSdk(ctx, model.SourceTranslation)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "SourceTranslation"})
		}
		if unpacked != nil {
			sdk.SourceTranslation = unpacked
		}
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

	// Handling Primitives
	if !model.ToInterface.IsNull() && !model.ToInterface.IsUnknown() {
		sdk.ToInterface = model.ToInterface.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "ToInterface", "value": *sdk.ToInterface})
	}

	// Handling Primitives
	if !model.TranslatedAddressSingle.IsNull() && !model.TranslatedAddressSingle.IsUnknown() {
		sdk.TranslatedAddressSingle = model.TranslatedAddressSingle.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "TranslatedAddressSingle", "value": *sdk.TranslatedAddressSingle})
	}

	// Handling Primitives
	if !model.TranslatedPort.IsNull() && !model.TranslatedPort.IsUnknown() {
		val := int32(model.TranslatedPort.ValueInt64())
		sdk.TranslatedPort = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "TranslatedPort", "value": *sdk.TranslatedPort})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.NatRules", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for NatRules ---
func packNatRulesFromSdk(ctx context.Context, sdk network_services.NatRules) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.NatRules", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.NatRules
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.ActiveActiveDeviceBinding != nil {
		model.ActiveActiveDeviceBinding = basetypes.NewStringValue(*sdk.ActiveActiveDeviceBinding)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "ActiveActiveDeviceBinding", "value": *sdk.ActiveActiveDeviceBinding})
	} else {
		model.ActiveActiveDeviceBinding = basetypes.NewStringNull()
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
	if sdk.Distribution != nil {
		model.Distribution = basetypes.NewStringValue(*sdk.Distribution)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Distribution", "value": *sdk.Distribution})
	} else {
		model.Distribution = basetypes.NewStringNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.DnsRewrite != nil {
		tflog.Debug(ctx, "Packing nested object for field DnsRewrite")
		packed, d := packNatRulesDnsRewriteFromSdk(ctx, *sdk.DnsRewrite)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "DnsRewrite"})
		}
		model.DnsRewrite = packed
	} else {
		model.DnsRewrite = basetypes.NewObjectNull(models.NatRulesDnsRewrite{}.AttrTypes())
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
	model.Id = basetypes.NewStringValue(sdk.Id)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Id", "value": sdk.Id})
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	// Handling Primitives
	// Standard primitive packing
	if sdk.NatType != nil {
		model.NatType = basetypes.NewStringValue(*sdk.NatType)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "NatType", "value": *sdk.NatType})
	} else {
		model.NatType = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	model.Service = basetypes.NewStringValue(sdk.Service)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Service", "value": sdk.Service})
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
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.SourceTranslation != nil {
		tflog.Debug(ctx, "Packing nested object for field SourceTranslation")
		packed, d := packNatRulesSourceTranslationFromSdk(ctx, *sdk.SourceTranslation)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "SourceTranslation"})
		}
		model.SourceTranslation = packed
	} else {
		model.SourceTranslation = basetypes.NewObjectNull(models.NatRulesSourceTranslation{}.AttrTypes())
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
	// Handling Primitives
	// Standard primitive packing
	if sdk.ToInterface != nil {
		model.ToInterface = basetypes.NewStringValue(*sdk.ToInterface)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "ToInterface", "value": *sdk.ToInterface})
	} else {
		model.ToInterface = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.TranslatedAddressSingle != nil {
		model.TranslatedAddressSingle = basetypes.NewStringValue(*sdk.TranslatedAddressSingle)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "TranslatedAddressSingle", "value": *sdk.TranslatedAddressSingle})
	} else {
		model.TranslatedAddressSingle = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.TranslatedPort != nil {
		model.TranslatedPort = basetypes.NewInt64Value(int64(*sdk.TranslatedPort))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "TranslatedPort", "value": *sdk.TranslatedPort})
	} else {
		model.TranslatedPort = basetypes.NewInt64Null()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.NatRules{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.NatRules", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for NatRules ---
func unpackNatRulesListToSdk(ctx context.Context, list types.List) ([]network_services.NatRules, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.NatRules")
	diags := diag.Diagnostics{}
	var data []models.NatRules
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.NatRules, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.NatRules{}.AttrTypes(), &item)
		unpacked, d := unpackNatRulesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.NatRules", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for NatRules ---
func packNatRulesListFromSdk(ctx context.Context, sdks []network_services.NatRules) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.NatRules")
	diags := diag.Diagnostics{}
	var data []models.NatRules

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.NatRules
		obj, d := packNatRulesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.NatRules{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.NatRules", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.NatRules{}.AttrType(), data)
}

// --- Unpacker for NatRulesDnsRewrite ---
func unpackNatRulesDnsRewriteToSdk(ctx context.Context, obj types.Object) (*network_services.NatRulesDnsRewrite, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.NatRulesDnsRewrite", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.NatRulesDnsRewrite
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.NatRulesDnsRewrite
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Direction.IsNull() && !model.Direction.IsUnknown() {
		sdk.Direction = model.Direction.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Direction", "value": *sdk.Direction})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.NatRulesDnsRewrite", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for NatRulesDnsRewrite ---
func packNatRulesDnsRewriteFromSdk(ctx context.Context, sdk network_services.NatRulesDnsRewrite) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.NatRulesDnsRewrite", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.NatRulesDnsRewrite
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Direction != nil {
		model.Direction = basetypes.NewStringValue(*sdk.Direction)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Direction", "value": *sdk.Direction})
	} else {
		model.Direction = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.NatRulesDnsRewrite{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.NatRulesDnsRewrite", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for NatRulesDnsRewrite ---
func unpackNatRulesDnsRewriteListToSdk(ctx context.Context, list types.List) ([]network_services.NatRulesDnsRewrite, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.NatRulesDnsRewrite")
	diags := diag.Diagnostics{}
	var data []models.NatRulesDnsRewrite
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.NatRulesDnsRewrite, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.NatRulesDnsRewrite{}.AttrTypes(), &item)
		unpacked, d := unpackNatRulesDnsRewriteToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.NatRulesDnsRewrite", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for NatRulesDnsRewrite ---
func packNatRulesDnsRewriteListFromSdk(ctx context.Context, sdks []network_services.NatRulesDnsRewrite) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.NatRulesDnsRewrite")
	diags := diag.Diagnostics{}
	var data []models.NatRulesDnsRewrite

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.NatRulesDnsRewrite
		obj, d := packNatRulesDnsRewriteFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.NatRulesDnsRewrite{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.NatRulesDnsRewrite", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.NatRulesDnsRewrite{}.AttrType(), data)
}

// --- Unpacker for NatRulesSourceTranslation ---
func unpackNatRulesSourceTranslationToSdk(ctx context.Context, obj types.Object) (*network_services.NatRulesSourceTranslation, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.NatRulesSourceTranslation", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.NatRulesSourceTranslation
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.NatRulesSourceTranslation
	var d diag.Diagnostics
	// Handling Primitives
	if !model.BiDirectional.IsNull() && !model.BiDirectional.IsUnknown() {
		sdk.BiDirectional = model.BiDirectional.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "BiDirectional", "value": *sdk.BiDirectional})
	}

	// Handling Objects
	if !model.Fallback.IsNull() && !model.Fallback.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Fallback")
		unpacked, d := unpackNatRulesSourceTranslationFallbackToSdk(ctx, model.Fallback)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Fallback"})
		}
		if unpacked != nil {
			sdk.Fallback = unpacked
		}
	}

	// Handling Primitives
	if !model.Interface.IsNull() && !model.Interface.IsUnknown() {
		sdk.Interface = model.Interface.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Interface", "value": *sdk.Interface})
	}

	// Handling Lists
	if !model.TranslatedAddressArray.IsNull() && !model.TranslatedAddressArray.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field TranslatedAddressArray")
		diags.Append(model.TranslatedAddressArray.ElementsAs(ctx, &sdk.TranslatedAddressArray, false)...)
	}

	// Handling Primitives
	if !model.TranslatedAddressSingle.IsNull() && !model.TranslatedAddressSingle.IsUnknown() {
		sdk.TranslatedAddressSingle = model.TranslatedAddressSingle.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "TranslatedAddressSingle", "value": *sdk.TranslatedAddressSingle})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.NatRulesSourceTranslation", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for NatRulesSourceTranslation ---
func packNatRulesSourceTranslationFromSdk(ctx context.Context, sdk network_services.NatRulesSourceTranslation) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.NatRulesSourceTranslation", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.NatRulesSourceTranslation
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.BiDirectional != nil {
		model.BiDirectional = basetypes.NewBoolValue(*sdk.BiDirectional)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "BiDirectional", "value": *sdk.BiDirectional})
	} else {
		model.BiDirectional = basetypes.NewBoolNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Fallback != nil {
		tflog.Debug(ctx, "Packing nested object for field Fallback")
		packed, d := packNatRulesSourceTranslationFallbackFromSdk(ctx, *sdk.Fallback)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Fallback"})
		}
		model.Fallback = packed
	} else {
		model.Fallback = basetypes.NewObjectNull(models.NatRulesSourceTranslationFallback{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Interface != nil {
		model.Interface = basetypes.NewStringValue(*sdk.Interface)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Interface", "value": *sdk.Interface})
	} else {
		model.Interface = basetypes.NewStringNull()
	}
	// Handling Lists
	if sdk.TranslatedAddressArray != nil {
		tflog.Debug(ctx, "Packing list of primitives for field TranslatedAddressArray")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.TranslatedAddressArray, d = basetypes.NewListValueFrom(ctx, elemType, sdk.TranslatedAddressArray)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.TranslatedAddressArray = basetypes.NewListNull(elemType)
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.TranslatedAddressSingle != nil {
		model.TranslatedAddressSingle = basetypes.NewStringValue(*sdk.TranslatedAddressSingle)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "TranslatedAddressSingle", "value": *sdk.TranslatedAddressSingle})
	} else {
		model.TranslatedAddressSingle = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.NatRulesSourceTranslation{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.NatRulesSourceTranslation", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for NatRulesSourceTranslation ---
func unpackNatRulesSourceTranslationListToSdk(ctx context.Context, list types.List) ([]network_services.NatRulesSourceTranslation, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.NatRulesSourceTranslation")
	diags := diag.Diagnostics{}
	var data []models.NatRulesSourceTranslation
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.NatRulesSourceTranslation, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.NatRulesSourceTranslation{}.AttrTypes(), &item)
		unpacked, d := unpackNatRulesSourceTranslationToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.NatRulesSourceTranslation", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for NatRulesSourceTranslation ---
func packNatRulesSourceTranslationListFromSdk(ctx context.Context, sdks []network_services.NatRulesSourceTranslation) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.NatRulesSourceTranslation")
	diags := diag.Diagnostics{}
	var data []models.NatRulesSourceTranslation

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.NatRulesSourceTranslation
		obj, d := packNatRulesSourceTranslationFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.NatRulesSourceTranslation{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.NatRulesSourceTranslation", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.NatRulesSourceTranslation{}.AttrType(), data)
}

// --- Unpacker for NatRulesSourceTranslationFallback ---
func unpackNatRulesSourceTranslationFallbackToSdk(ctx context.Context, obj types.Object) (*network_services.NatRulesSourceTranslationFallback, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.NatRulesSourceTranslationFallback", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.NatRulesSourceTranslationFallback
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.NatRulesSourceTranslationFallback
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Interface.IsNull() && !model.Interface.IsUnknown() {
		sdk.Interface = model.Interface.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Interface", "value": *sdk.Interface})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.NatRulesSourceTranslationFallback", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for NatRulesSourceTranslationFallback ---
func packNatRulesSourceTranslationFallbackFromSdk(ctx context.Context, sdk network_services.NatRulesSourceTranslationFallback) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.NatRulesSourceTranslationFallback", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.NatRulesSourceTranslationFallback
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Interface != nil {
		model.Interface = basetypes.NewStringValue(*sdk.Interface)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Interface", "value": *sdk.Interface})
	} else {
		model.Interface = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.NatRulesSourceTranslationFallback{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.NatRulesSourceTranslationFallback", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for NatRulesSourceTranslationFallback ---
func unpackNatRulesSourceTranslationFallbackListToSdk(ctx context.Context, list types.List) ([]network_services.NatRulesSourceTranslationFallback, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.NatRulesSourceTranslationFallback")
	diags := diag.Diagnostics{}
	var data []models.NatRulesSourceTranslationFallback
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.NatRulesSourceTranslationFallback, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.NatRulesSourceTranslationFallback{}.AttrTypes(), &item)
		unpacked, d := unpackNatRulesSourceTranslationFallbackToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.NatRulesSourceTranslationFallback", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for NatRulesSourceTranslationFallback ---
func packNatRulesSourceTranslationFallbackListFromSdk(ctx context.Context, sdks []network_services.NatRulesSourceTranslationFallback) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.NatRulesSourceTranslationFallback")
	diags := diag.Diagnostics{}
	var data []models.NatRulesSourceTranslationFallback

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.NatRulesSourceTranslationFallback
		obj, d := packNatRulesSourceTranslationFallbackFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.NatRulesSourceTranslationFallback{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.NatRulesSourceTranslationFallback", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.NatRulesSourceTranslationFallback{}.AttrType(), data)
}
