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

	// Handling Objects
	if !model.DestinationTranslation.IsNull() && !model.DestinationTranslation.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field DestinationTranslation")
		unpacked, d := unpackNatRulesDestinationTranslationToSdk(ctx, model.DestinationTranslation)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "DestinationTranslation"})
		}
		if unpacked != nil {
			sdk.DestinationTranslation = unpacked
		}
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

	// Handling Objects
	if !model.DynamicDestinationTranslation.IsNull() && !model.DynamicDestinationTranslation.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field DynamicDestinationTranslation")
		unpacked, d := unpackNatRulesDynamicDestinationTranslationToSdk(ctx, model.DynamicDestinationTranslation)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "DynamicDestinationTranslation"})
		}
		if unpacked != nil {
			sdk.DynamicDestinationTranslation = unpacked
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
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.DestinationTranslation != nil {
		tflog.Debug(ctx, "Packing nested object for field DestinationTranslation")
		packed, d := packNatRulesDestinationTranslationFromSdk(ctx, *sdk.DestinationTranslation)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "DestinationTranslation"})
		}
		model.DestinationTranslation = packed
	} else {
		model.DestinationTranslation = basetypes.NewObjectNull(models.NatRulesDestinationTranslation{}.AttrTypes())
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
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.DynamicDestinationTranslation != nil {
		tflog.Debug(ctx, "Packing nested object for field DynamicDestinationTranslation")
		packed, d := packNatRulesDynamicDestinationTranslationFromSdk(ctx, *sdk.DynamicDestinationTranslation)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "DynamicDestinationTranslation"})
		}
		model.DynamicDestinationTranslation = packed
	} else {
		model.DynamicDestinationTranslation = basetypes.NewObjectNull(models.NatRulesDynamicDestinationTranslation{}.AttrTypes())
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

// --- Unpacker for NatRulesDestinationTranslation ---
func unpackNatRulesDestinationTranslationToSdk(ctx context.Context, obj types.Object) (*network_services.NatRulesDestinationTranslation, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.NatRulesDestinationTranslation", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.NatRulesDestinationTranslation
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.NatRulesDestinationTranslation
	var d diag.Diagnostics
	// Handling Objects
	if !model.DnsRewrite.IsNull() && !model.DnsRewrite.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field DnsRewrite")
		unpacked, d := unpackNatRulesDestinationTranslationDnsRewriteToSdk(ctx, model.DnsRewrite)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "DnsRewrite"})
		}
		if unpacked != nil {
			sdk.DnsRewrite = unpacked
		}
	}

	// Handling Primitives
	if !model.TranslatedAddress.IsNull() && !model.TranslatedAddress.IsUnknown() {
		sdk.TranslatedAddress = model.TranslatedAddress.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "TranslatedAddress", "value": *sdk.TranslatedAddress})
	}

	// Handling Primitives
	if !model.TranslatedPort.IsNull() && !model.TranslatedPort.IsUnknown() {
		val := int32(model.TranslatedPort.ValueInt64())
		sdk.TranslatedPort = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "TranslatedPort", "value": *sdk.TranslatedPort})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.NatRulesDestinationTranslation", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for NatRulesDestinationTranslation ---
func packNatRulesDestinationTranslationFromSdk(ctx context.Context, sdk network_services.NatRulesDestinationTranslation) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.NatRulesDestinationTranslation", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.NatRulesDestinationTranslation
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.DnsRewrite != nil {
		tflog.Debug(ctx, "Packing nested object for field DnsRewrite")
		packed, d := packNatRulesDestinationTranslationDnsRewriteFromSdk(ctx, *sdk.DnsRewrite)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "DnsRewrite"})
		}
		model.DnsRewrite = packed
	} else {
		model.DnsRewrite = basetypes.NewObjectNull(models.NatRulesDestinationTranslationDnsRewrite{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.TranslatedAddress != nil {
		model.TranslatedAddress = basetypes.NewStringValue(*sdk.TranslatedAddress)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "TranslatedAddress", "value": *sdk.TranslatedAddress})
	} else {
		model.TranslatedAddress = basetypes.NewStringNull()
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

	obj, d := types.ObjectValueFrom(ctx, models.NatRulesDestinationTranslation{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.NatRulesDestinationTranslation", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for NatRulesDestinationTranslation ---
func unpackNatRulesDestinationTranslationListToSdk(ctx context.Context, list types.List) ([]network_services.NatRulesDestinationTranslation, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.NatRulesDestinationTranslation")
	diags := diag.Diagnostics{}
	var data []models.NatRulesDestinationTranslation
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.NatRulesDestinationTranslation, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.NatRulesDestinationTranslation{}.AttrTypes(), &item)
		unpacked, d := unpackNatRulesDestinationTranslationToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.NatRulesDestinationTranslation", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for NatRulesDestinationTranslation ---
func packNatRulesDestinationTranslationListFromSdk(ctx context.Context, sdks []network_services.NatRulesDestinationTranslation) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.NatRulesDestinationTranslation")
	diags := diag.Diagnostics{}
	var data []models.NatRulesDestinationTranslation

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.NatRulesDestinationTranslation
		obj, d := packNatRulesDestinationTranslationFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.NatRulesDestinationTranslation{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.NatRulesDestinationTranslation", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.NatRulesDestinationTranslation{}.AttrType(), data)
}

// --- Unpacker for NatRulesDestinationTranslationDnsRewrite ---
func unpackNatRulesDestinationTranslationDnsRewriteToSdk(ctx context.Context, obj types.Object) (*network_services.NatRulesDestinationTranslationDnsRewrite, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.NatRulesDestinationTranslationDnsRewrite", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.NatRulesDestinationTranslationDnsRewrite
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.NatRulesDestinationTranslationDnsRewrite
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Direction.IsNull() && !model.Direction.IsUnknown() {
		sdk.Direction = model.Direction.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Direction", "value": *sdk.Direction})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.NatRulesDestinationTranslationDnsRewrite", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for NatRulesDestinationTranslationDnsRewrite ---
func packNatRulesDestinationTranslationDnsRewriteFromSdk(ctx context.Context, sdk network_services.NatRulesDestinationTranslationDnsRewrite) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.NatRulesDestinationTranslationDnsRewrite", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.NatRulesDestinationTranslationDnsRewrite
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

	obj, d := types.ObjectValueFrom(ctx, models.NatRulesDestinationTranslationDnsRewrite{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.NatRulesDestinationTranslationDnsRewrite", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for NatRulesDestinationTranslationDnsRewrite ---
func unpackNatRulesDestinationTranslationDnsRewriteListToSdk(ctx context.Context, list types.List) ([]network_services.NatRulesDestinationTranslationDnsRewrite, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.NatRulesDestinationTranslationDnsRewrite")
	diags := diag.Diagnostics{}
	var data []models.NatRulesDestinationTranslationDnsRewrite
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.NatRulesDestinationTranslationDnsRewrite, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.NatRulesDestinationTranslationDnsRewrite{}.AttrTypes(), &item)
		unpacked, d := unpackNatRulesDestinationTranslationDnsRewriteToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.NatRulesDestinationTranslationDnsRewrite", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for NatRulesDestinationTranslationDnsRewrite ---
func packNatRulesDestinationTranslationDnsRewriteListFromSdk(ctx context.Context, sdks []network_services.NatRulesDestinationTranslationDnsRewrite) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.NatRulesDestinationTranslationDnsRewrite")
	diags := diag.Diagnostics{}
	var data []models.NatRulesDestinationTranslationDnsRewrite

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.NatRulesDestinationTranslationDnsRewrite
		obj, d := packNatRulesDestinationTranslationDnsRewriteFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.NatRulesDestinationTranslationDnsRewrite{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.NatRulesDestinationTranslationDnsRewrite", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.NatRulesDestinationTranslationDnsRewrite{}.AttrType(), data)
}

// --- Unpacker for NatRulesDynamicDestinationTranslation ---
func unpackNatRulesDynamicDestinationTranslationToSdk(ctx context.Context, obj types.Object) (*network_services.NatRulesDynamicDestinationTranslation, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.NatRulesDynamicDestinationTranslation", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.NatRulesDynamicDestinationTranslation
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.NatRulesDynamicDestinationTranslation
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Distribution.IsNull() && !model.Distribution.IsUnknown() {
		sdk.Distribution = model.Distribution.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Distribution", "value": *sdk.Distribution})
	}

	// Handling Primitives
	if !model.TranslatedAddress.IsNull() && !model.TranslatedAddress.IsUnknown() {
		sdk.TranslatedAddress = model.TranslatedAddress.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "TranslatedAddress", "value": *sdk.TranslatedAddress})
	}

	// Handling Primitives
	if !model.TranslatedPort.IsNull() && !model.TranslatedPort.IsUnknown() {
		val := int32(model.TranslatedPort.ValueInt64())
		sdk.TranslatedPort = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "TranslatedPort", "value": *sdk.TranslatedPort})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.NatRulesDynamicDestinationTranslation", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for NatRulesDynamicDestinationTranslation ---
func packNatRulesDynamicDestinationTranslationFromSdk(ctx context.Context, sdk network_services.NatRulesDynamicDestinationTranslation) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.NatRulesDynamicDestinationTranslation", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.NatRulesDynamicDestinationTranslation
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Distribution != nil {
		model.Distribution = basetypes.NewStringValue(*sdk.Distribution)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Distribution", "value": *sdk.Distribution})
	} else {
		model.Distribution = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.TranslatedAddress != nil {
		model.TranslatedAddress = basetypes.NewStringValue(*sdk.TranslatedAddress)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "TranslatedAddress", "value": *sdk.TranslatedAddress})
	} else {
		model.TranslatedAddress = basetypes.NewStringNull()
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

	obj, d := types.ObjectValueFrom(ctx, models.NatRulesDynamicDestinationTranslation{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.NatRulesDynamicDestinationTranslation", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for NatRulesDynamicDestinationTranslation ---
func unpackNatRulesDynamicDestinationTranslationListToSdk(ctx context.Context, list types.List) ([]network_services.NatRulesDynamicDestinationTranslation, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.NatRulesDynamicDestinationTranslation")
	diags := diag.Diagnostics{}
	var data []models.NatRulesDynamicDestinationTranslation
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.NatRulesDynamicDestinationTranslation, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.NatRulesDynamicDestinationTranslation{}.AttrTypes(), &item)
		unpacked, d := unpackNatRulesDynamicDestinationTranslationToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.NatRulesDynamicDestinationTranslation", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for NatRulesDynamicDestinationTranslation ---
func packNatRulesDynamicDestinationTranslationListFromSdk(ctx context.Context, sdks []network_services.NatRulesDynamicDestinationTranslation) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.NatRulesDynamicDestinationTranslation")
	diags := diag.Diagnostics{}
	var data []models.NatRulesDynamicDestinationTranslation

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.NatRulesDynamicDestinationTranslation
		obj, d := packNatRulesDynamicDestinationTranslationFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.NatRulesDynamicDestinationTranslation{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.NatRulesDynamicDestinationTranslation", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.NatRulesDynamicDestinationTranslation{}.AttrType(), data)
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
	// Handling Objects
	if !model.DynamicIp.IsNull() && !model.DynamicIp.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field DynamicIp")
		unpacked, d := unpackNatRulesSourceTranslationDynamicIpToSdk(ctx, model.DynamicIp)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "DynamicIp"})
		}
		if unpacked != nil {
			sdk.DynamicIp = unpacked
		}
	}

	// Handling Objects
	if !model.DynamicIpAndPort.IsNull() && !model.DynamicIpAndPort.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field DynamicIpAndPort")
		unpacked, d := unpackNatRulesSourceTranslationDynamicIpAndPortToSdk(ctx, model.DynamicIpAndPort)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "DynamicIpAndPort"})
		}
		if unpacked != nil {
			sdk.DynamicIpAndPort = unpacked
		}
	}

	// Handling Objects
	if !model.StaticIp.IsNull() && !model.StaticIp.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field StaticIp")
		unpacked, d := unpackNatRulesSourceTranslationStaticIpToSdk(ctx, model.StaticIp)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "StaticIp"})
		}
		if unpacked != nil {
			sdk.StaticIp = unpacked
		}
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
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.DynamicIp != nil {
		tflog.Debug(ctx, "Packing nested object for field DynamicIp")
		packed, d := packNatRulesSourceTranslationDynamicIpFromSdk(ctx, *sdk.DynamicIp)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "DynamicIp"})
		}
		model.DynamicIp = packed
	} else {
		model.DynamicIp = basetypes.NewObjectNull(models.NatRulesSourceTranslationDynamicIp{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.DynamicIpAndPort != nil {
		tflog.Debug(ctx, "Packing nested object for field DynamicIpAndPort")
		packed, d := packNatRulesSourceTranslationDynamicIpAndPortFromSdk(ctx, *sdk.DynamicIpAndPort)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "DynamicIpAndPort"})
		}
		model.DynamicIpAndPort = packed
	} else {
		model.DynamicIpAndPort = basetypes.NewObjectNull(models.NatRulesSourceTranslationDynamicIpAndPort{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.StaticIp != nil {
		tflog.Debug(ctx, "Packing nested object for field StaticIp")
		packed, d := packNatRulesSourceTranslationStaticIpFromSdk(ctx, *sdk.StaticIp)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "StaticIp"})
		}
		model.StaticIp = packed
	} else {
		model.StaticIp = basetypes.NewObjectNull(models.NatRulesSourceTranslationStaticIp{}.AttrTypes())
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

// --- Unpacker for NatRulesSourceTranslationDynamicIp ---
func unpackNatRulesSourceTranslationDynamicIpToSdk(ctx context.Context, obj types.Object) (*network_services.NatRulesSourceTranslationDynamicIp, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.NatRulesSourceTranslationDynamicIp", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.NatRulesSourceTranslationDynamicIp
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.NatRulesSourceTranslationDynamicIp
	var d diag.Diagnostics
	// Handling Objects
	if !model.Fallback.IsNull() && !model.Fallback.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Fallback")
		unpacked, d := unpackNatRulesSourceTranslationDynamicIpFallbackToSdk(ctx, model.Fallback)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Fallback"})
		}
		if unpacked != nil {
			sdk.Fallback = unpacked
		}
	}

	// Handling Lists
	if !model.TranslatedAddress.IsNull() && !model.TranslatedAddress.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field TranslatedAddress")
		diags.Append(model.TranslatedAddress.ElementsAs(ctx, &sdk.TranslatedAddress, false)...)
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.NatRulesSourceTranslationDynamicIp", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for NatRulesSourceTranslationDynamicIp ---
func packNatRulesSourceTranslationDynamicIpFromSdk(ctx context.Context, sdk network_services.NatRulesSourceTranslationDynamicIp) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.NatRulesSourceTranslationDynamicIp", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.NatRulesSourceTranslationDynamicIp
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Fallback != nil {
		tflog.Debug(ctx, "Packing nested object for field Fallback")
		packed, d := packNatRulesSourceTranslationDynamicIpFallbackFromSdk(ctx, *sdk.Fallback)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Fallback"})
		}
		model.Fallback = packed
	} else {
		model.Fallback = basetypes.NewObjectNull(models.NatRulesSourceTranslationDynamicIpFallback{}.AttrTypes())
	}
	// Handling Lists
	if sdk.TranslatedAddress != nil {
		tflog.Debug(ctx, "Packing list of primitives for field TranslatedAddress")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.TranslatedAddress, d = basetypes.NewListValueFrom(ctx, elemType, sdk.TranslatedAddress)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.TranslatedAddress = basetypes.NewListNull(elemType)
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.NatRulesSourceTranslationDynamicIp{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.NatRulesSourceTranslationDynamicIp", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for NatRulesSourceTranslationDynamicIp ---
func unpackNatRulesSourceTranslationDynamicIpListToSdk(ctx context.Context, list types.List) ([]network_services.NatRulesSourceTranslationDynamicIp, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.NatRulesSourceTranslationDynamicIp")
	diags := diag.Diagnostics{}
	var data []models.NatRulesSourceTranslationDynamicIp
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.NatRulesSourceTranslationDynamicIp, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.NatRulesSourceTranslationDynamicIp{}.AttrTypes(), &item)
		unpacked, d := unpackNatRulesSourceTranslationDynamicIpToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.NatRulesSourceTranslationDynamicIp", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for NatRulesSourceTranslationDynamicIp ---
func packNatRulesSourceTranslationDynamicIpListFromSdk(ctx context.Context, sdks []network_services.NatRulesSourceTranslationDynamicIp) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.NatRulesSourceTranslationDynamicIp")
	diags := diag.Diagnostics{}
	var data []models.NatRulesSourceTranslationDynamicIp

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.NatRulesSourceTranslationDynamicIp
		obj, d := packNatRulesSourceTranslationDynamicIpFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.NatRulesSourceTranslationDynamicIp{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.NatRulesSourceTranslationDynamicIp", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.NatRulesSourceTranslationDynamicIp{}.AttrType(), data)
}

// --- Unpacker for NatRulesSourceTranslationDynamicIpFallback ---
func unpackNatRulesSourceTranslationDynamicIpFallbackToSdk(ctx context.Context, obj types.Object) (*network_services.NatRulesSourceTranslationDynamicIpFallback, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.NatRulesSourceTranslationDynamicIpFallback", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.NatRulesSourceTranslationDynamicIpFallback
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.NatRulesSourceTranslationDynamicIpFallback
	var d diag.Diagnostics
	// Handling Objects
	if !model.InterfaceAddress.IsNull() && !model.InterfaceAddress.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field InterfaceAddress")
		unpacked, d := unpackNatRulesSourceTranslationDynamicIpFallbackInterfaceAddressToSdk(ctx, model.InterfaceAddress)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "InterfaceAddress"})
		}
		if unpacked != nil {
			sdk.InterfaceAddress = unpacked
		}
	}

	// Handling Lists
	if !model.TranslatedAddress.IsNull() && !model.TranslatedAddress.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field TranslatedAddress")
		diags.Append(model.TranslatedAddress.ElementsAs(ctx, &sdk.TranslatedAddress, false)...)
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.NatRulesSourceTranslationDynamicIpFallback", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for NatRulesSourceTranslationDynamicIpFallback ---
func packNatRulesSourceTranslationDynamicIpFallbackFromSdk(ctx context.Context, sdk network_services.NatRulesSourceTranslationDynamicIpFallback) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.NatRulesSourceTranslationDynamicIpFallback", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.NatRulesSourceTranslationDynamicIpFallback
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.InterfaceAddress != nil {
		tflog.Debug(ctx, "Packing nested object for field InterfaceAddress")
		packed, d := packNatRulesSourceTranslationDynamicIpFallbackInterfaceAddressFromSdk(ctx, *sdk.InterfaceAddress)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "InterfaceAddress"})
		}
		model.InterfaceAddress = packed
	} else {
		model.InterfaceAddress = basetypes.NewObjectNull(models.NatRulesSourceTranslationDynamicIpFallbackInterfaceAddress{}.AttrTypes())
	}
	// Handling Lists
	if sdk.TranslatedAddress != nil {
		tflog.Debug(ctx, "Packing list of primitives for field TranslatedAddress")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.TranslatedAddress, d = basetypes.NewListValueFrom(ctx, elemType, sdk.TranslatedAddress)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.TranslatedAddress = basetypes.NewListNull(elemType)
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.NatRulesSourceTranslationDynamicIpFallback{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.NatRulesSourceTranslationDynamicIpFallback", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for NatRulesSourceTranslationDynamicIpFallback ---
func unpackNatRulesSourceTranslationDynamicIpFallbackListToSdk(ctx context.Context, list types.List) ([]network_services.NatRulesSourceTranslationDynamicIpFallback, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.NatRulesSourceTranslationDynamicIpFallback")
	diags := diag.Diagnostics{}
	var data []models.NatRulesSourceTranslationDynamicIpFallback
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.NatRulesSourceTranslationDynamicIpFallback, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.NatRulesSourceTranslationDynamicIpFallback{}.AttrTypes(), &item)
		unpacked, d := unpackNatRulesSourceTranslationDynamicIpFallbackToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.NatRulesSourceTranslationDynamicIpFallback", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for NatRulesSourceTranslationDynamicIpFallback ---
func packNatRulesSourceTranslationDynamicIpFallbackListFromSdk(ctx context.Context, sdks []network_services.NatRulesSourceTranslationDynamicIpFallback) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.NatRulesSourceTranslationDynamicIpFallback")
	diags := diag.Diagnostics{}
	var data []models.NatRulesSourceTranslationDynamicIpFallback

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.NatRulesSourceTranslationDynamicIpFallback
		obj, d := packNatRulesSourceTranslationDynamicIpFallbackFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.NatRulesSourceTranslationDynamicIpFallback{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.NatRulesSourceTranslationDynamicIpFallback", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.NatRulesSourceTranslationDynamicIpFallback{}.AttrType(), data)
}

// --- Unpacker for NatRulesSourceTranslationDynamicIpFallbackInterfaceAddress ---
func unpackNatRulesSourceTranslationDynamicIpFallbackInterfaceAddressToSdk(ctx context.Context, obj types.Object) (*network_services.NatRulesSourceTranslationDynamicIpFallbackInterfaceAddress, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.NatRulesSourceTranslationDynamicIpFallbackInterfaceAddress", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.NatRulesSourceTranslationDynamicIpFallbackInterfaceAddress
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.NatRulesSourceTranslationDynamicIpFallbackInterfaceAddress
	var d diag.Diagnostics
	// Handling Primitives
	if !model.FloatingIp.IsNull() && !model.FloatingIp.IsUnknown() {
		sdk.FloatingIp = model.FloatingIp.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "FloatingIp", "value": *sdk.FloatingIp})
	}

	// Handling Primitives
	if !model.Interface.IsNull() && !model.Interface.IsUnknown() {
		sdk.Interface = model.Interface.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Interface", "value": *sdk.Interface})
	}

	// Handling Primitives
	if !model.Ip.IsNull() && !model.Ip.IsUnknown() {
		sdk.Ip = model.Ip.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Ip", "value": *sdk.Ip})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.NatRulesSourceTranslationDynamicIpFallbackInterfaceAddress", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for NatRulesSourceTranslationDynamicIpFallbackInterfaceAddress ---
func packNatRulesSourceTranslationDynamicIpFallbackInterfaceAddressFromSdk(ctx context.Context, sdk network_services.NatRulesSourceTranslationDynamicIpFallbackInterfaceAddress) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.NatRulesSourceTranslationDynamicIpFallbackInterfaceAddress", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.NatRulesSourceTranslationDynamicIpFallbackInterfaceAddress
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.FloatingIp != nil {
		model.FloatingIp = basetypes.NewStringValue(*sdk.FloatingIp)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "FloatingIp", "value": *sdk.FloatingIp})
	} else {
		model.FloatingIp = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Interface != nil {
		model.Interface = basetypes.NewStringValue(*sdk.Interface)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Interface", "value": *sdk.Interface})
	} else {
		model.Interface = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Ip != nil {
		model.Ip = basetypes.NewStringValue(*sdk.Ip)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Ip", "value": *sdk.Ip})
	} else {
		model.Ip = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.NatRulesSourceTranslationDynamicIpFallbackInterfaceAddress{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.NatRulesSourceTranslationDynamicIpFallbackInterfaceAddress", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for NatRulesSourceTranslationDynamicIpFallbackInterfaceAddress ---
func unpackNatRulesSourceTranslationDynamicIpFallbackInterfaceAddressListToSdk(ctx context.Context, list types.List) ([]network_services.NatRulesSourceTranslationDynamicIpFallbackInterfaceAddress, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.NatRulesSourceTranslationDynamicIpFallbackInterfaceAddress")
	diags := diag.Diagnostics{}
	var data []models.NatRulesSourceTranslationDynamicIpFallbackInterfaceAddress
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.NatRulesSourceTranslationDynamicIpFallbackInterfaceAddress, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.NatRulesSourceTranslationDynamicIpFallbackInterfaceAddress{}.AttrTypes(), &item)
		unpacked, d := unpackNatRulesSourceTranslationDynamicIpFallbackInterfaceAddressToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.NatRulesSourceTranslationDynamicIpFallbackInterfaceAddress", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for NatRulesSourceTranslationDynamicIpFallbackInterfaceAddress ---
func packNatRulesSourceTranslationDynamicIpFallbackInterfaceAddressListFromSdk(ctx context.Context, sdks []network_services.NatRulesSourceTranslationDynamicIpFallbackInterfaceAddress) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.NatRulesSourceTranslationDynamicIpFallbackInterfaceAddress")
	diags := diag.Diagnostics{}
	var data []models.NatRulesSourceTranslationDynamicIpFallbackInterfaceAddress

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.NatRulesSourceTranslationDynamicIpFallbackInterfaceAddress
		obj, d := packNatRulesSourceTranslationDynamicIpFallbackInterfaceAddressFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.NatRulesSourceTranslationDynamicIpFallbackInterfaceAddress{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.NatRulesSourceTranslationDynamicIpFallbackInterfaceAddress", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.NatRulesSourceTranslationDynamicIpFallbackInterfaceAddress{}.AttrType(), data)
}

// --- Unpacker for NatRulesSourceTranslationDynamicIpAndPort ---
func unpackNatRulesSourceTranslationDynamicIpAndPortToSdk(ctx context.Context, obj types.Object) (*network_services.NatRulesSourceTranslationDynamicIpAndPort, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.NatRulesSourceTranslationDynamicIpAndPort", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.NatRulesSourceTranslationDynamicIpAndPort
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.NatRulesSourceTranslationDynamicIpAndPort
	var d diag.Diagnostics
	// Handling Objects
	if !model.InterfaceAddress.IsNull() && !model.InterfaceAddress.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field InterfaceAddress")
		unpacked, d := unpackNatRulesSourceTranslationDynamicIpAndPortInterfaceAddressToSdk(ctx, model.InterfaceAddress)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "InterfaceAddress"})
		}
		if unpacked != nil {
			sdk.InterfaceAddress = unpacked
		}
	}

	// Handling Lists
	if !model.TranslatedAddress.IsNull() && !model.TranslatedAddress.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field TranslatedAddress")
		diags.Append(model.TranslatedAddress.ElementsAs(ctx, &sdk.TranslatedAddress, false)...)
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.NatRulesSourceTranslationDynamicIpAndPort", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for NatRulesSourceTranslationDynamicIpAndPort ---
func packNatRulesSourceTranslationDynamicIpAndPortFromSdk(ctx context.Context, sdk network_services.NatRulesSourceTranslationDynamicIpAndPort) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.NatRulesSourceTranslationDynamicIpAndPort", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.NatRulesSourceTranslationDynamicIpAndPort
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.InterfaceAddress != nil {
		tflog.Debug(ctx, "Packing nested object for field InterfaceAddress")
		packed, d := packNatRulesSourceTranslationDynamicIpAndPortInterfaceAddressFromSdk(ctx, *sdk.InterfaceAddress)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "InterfaceAddress"})
		}
		model.InterfaceAddress = packed
	} else {
		model.InterfaceAddress = basetypes.NewObjectNull(models.NatRulesSourceTranslationDynamicIpAndPortInterfaceAddress{}.AttrTypes())
	}
	// Handling Lists
	if sdk.TranslatedAddress != nil {
		tflog.Debug(ctx, "Packing list of primitives for field TranslatedAddress")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.TranslatedAddress, d = basetypes.NewListValueFrom(ctx, elemType, sdk.TranslatedAddress)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.TranslatedAddress = basetypes.NewListNull(elemType)
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.NatRulesSourceTranslationDynamicIpAndPort{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.NatRulesSourceTranslationDynamicIpAndPort", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for NatRulesSourceTranslationDynamicIpAndPort ---
func unpackNatRulesSourceTranslationDynamicIpAndPortListToSdk(ctx context.Context, list types.List) ([]network_services.NatRulesSourceTranslationDynamicIpAndPort, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.NatRulesSourceTranslationDynamicIpAndPort")
	diags := diag.Diagnostics{}
	var data []models.NatRulesSourceTranslationDynamicIpAndPort
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.NatRulesSourceTranslationDynamicIpAndPort, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.NatRulesSourceTranslationDynamicIpAndPort{}.AttrTypes(), &item)
		unpacked, d := unpackNatRulesSourceTranslationDynamicIpAndPortToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.NatRulesSourceTranslationDynamicIpAndPort", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for NatRulesSourceTranslationDynamicIpAndPort ---
func packNatRulesSourceTranslationDynamicIpAndPortListFromSdk(ctx context.Context, sdks []network_services.NatRulesSourceTranslationDynamicIpAndPort) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.NatRulesSourceTranslationDynamicIpAndPort")
	diags := diag.Diagnostics{}
	var data []models.NatRulesSourceTranslationDynamicIpAndPort

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.NatRulesSourceTranslationDynamicIpAndPort
		obj, d := packNatRulesSourceTranslationDynamicIpAndPortFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.NatRulesSourceTranslationDynamicIpAndPort{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.NatRulesSourceTranslationDynamicIpAndPort", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.NatRulesSourceTranslationDynamicIpAndPort{}.AttrType(), data)
}

// --- Unpacker for NatRulesSourceTranslationDynamicIpAndPortInterfaceAddress ---
func unpackNatRulesSourceTranslationDynamicIpAndPortInterfaceAddressToSdk(ctx context.Context, obj types.Object) (*network_services.NatRulesSourceTranslationDynamicIpAndPortInterfaceAddress, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.NatRulesSourceTranslationDynamicIpAndPortInterfaceAddress", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.NatRulesSourceTranslationDynamicIpAndPortInterfaceAddress
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.NatRulesSourceTranslationDynamicIpAndPortInterfaceAddress
	var d diag.Diagnostics
	// Handling Primitives
	if !model.FloatingIp.IsNull() && !model.FloatingIp.IsUnknown() {
		sdk.FloatingIp = model.FloatingIp.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "FloatingIp", "value": *sdk.FloatingIp})
	}

	// Handling Primitives
	if !model.Interface.IsNull() && !model.Interface.IsUnknown() {
		sdk.Interface = model.Interface.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Interface", "value": *sdk.Interface})
	}

	// Handling Primitives
	if !model.Ip.IsNull() && !model.Ip.IsUnknown() {
		sdk.Ip = model.Ip.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Ip", "value": *sdk.Ip})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.NatRulesSourceTranslationDynamicIpAndPortInterfaceAddress", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for NatRulesSourceTranslationDynamicIpAndPortInterfaceAddress ---
func packNatRulesSourceTranslationDynamicIpAndPortInterfaceAddressFromSdk(ctx context.Context, sdk network_services.NatRulesSourceTranslationDynamicIpAndPortInterfaceAddress) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.NatRulesSourceTranslationDynamicIpAndPortInterfaceAddress", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.NatRulesSourceTranslationDynamicIpAndPortInterfaceAddress
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.FloatingIp != nil {
		model.FloatingIp = basetypes.NewStringValue(*sdk.FloatingIp)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "FloatingIp", "value": *sdk.FloatingIp})
	} else {
		model.FloatingIp = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Interface != nil {
		model.Interface = basetypes.NewStringValue(*sdk.Interface)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Interface", "value": *sdk.Interface})
	} else {
		model.Interface = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Ip != nil {
		model.Ip = basetypes.NewStringValue(*sdk.Ip)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Ip", "value": *sdk.Ip})
	} else {
		model.Ip = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.NatRulesSourceTranslationDynamicIpAndPortInterfaceAddress{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.NatRulesSourceTranslationDynamicIpAndPortInterfaceAddress", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for NatRulesSourceTranslationDynamicIpAndPortInterfaceAddress ---
func unpackNatRulesSourceTranslationDynamicIpAndPortInterfaceAddressListToSdk(ctx context.Context, list types.List) ([]network_services.NatRulesSourceTranslationDynamicIpAndPortInterfaceAddress, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.NatRulesSourceTranslationDynamicIpAndPortInterfaceAddress")
	diags := diag.Diagnostics{}
	var data []models.NatRulesSourceTranslationDynamicIpAndPortInterfaceAddress
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.NatRulesSourceTranslationDynamicIpAndPortInterfaceAddress, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.NatRulesSourceTranslationDynamicIpAndPortInterfaceAddress{}.AttrTypes(), &item)
		unpacked, d := unpackNatRulesSourceTranslationDynamicIpAndPortInterfaceAddressToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.NatRulesSourceTranslationDynamicIpAndPortInterfaceAddress", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for NatRulesSourceTranslationDynamicIpAndPortInterfaceAddress ---
func packNatRulesSourceTranslationDynamicIpAndPortInterfaceAddressListFromSdk(ctx context.Context, sdks []network_services.NatRulesSourceTranslationDynamicIpAndPortInterfaceAddress) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.NatRulesSourceTranslationDynamicIpAndPortInterfaceAddress")
	diags := diag.Diagnostics{}
	var data []models.NatRulesSourceTranslationDynamicIpAndPortInterfaceAddress

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.NatRulesSourceTranslationDynamicIpAndPortInterfaceAddress
		obj, d := packNatRulesSourceTranslationDynamicIpAndPortInterfaceAddressFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.NatRulesSourceTranslationDynamicIpAndPortInterfaceAddress{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.NatRulesSourceTranslationDynamicIpAndPortInterfaceAddress", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.NatRulesSourceTranslationDynamicIpAndPortInterfaceAddress{}.AttrType(), data)
}

// --- Unpacker for NatRulesSourceTranslationStaticIp ---
func unpackNatRulesSourceTranslationStaticIpToSdk(ctx context.Context, obj types.Object) (*network_services.NatRulesSourceTranslationStaticIp, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.NatRulesSourceTranslationStaticIp", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.NatRulesSourceTranslationStaticIp
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.NatRulesSourceTranslationStaticIp
	var d diag.Diagnostics
	// Handling Primitives
	if !model.BiDirectional.IsNull() && !model.BiDirectional.IsUnknown() {
		sdk.BiDirectional = model.BiDirectional.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "BiDirectional", "value": *sdk.BiDirectional})
	}

	// Handling Primitives
	if !model.TranslatedAddress.IsNull() && !model.TranslatedAddress.IsUnknown() {
		sdk.TranslatedAddress = model.TranslatedAddress.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "TranslatedAddress", "value": *sdk.TranslatedAddress})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.NatRulesSourceTranslationStaticIp", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for NatRulesSourceTranslationStaticIp ---
func packNatRulesSourceTranslationStaticIpFromSdk(ctx context.Context, sdk network_services.NatRulesSourceTranslationStaticIp) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.NatRulesSourceTranslationStaticIp", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.NatRulesSourceTranslationStaticIp
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.BiDirectional != nil {
		model.BiDirectional = basetypes.NewStringValue(*sdk.BiDirectional)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "BiDirectional", "value": *sdk.BiDirectional})
	} else {
		model.BiDirectional = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.TranslatedAddress != nil {
		model.TranslatedAddress = basetypes.NewStringValue(*sdk.TranslatedAddress)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "TranslatedAddress", "value": *sdk.TranslatedAddress})
	} else {
		model.TranslatedAddress = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.NatRulesSourceTranslationStaticIp{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.NatRulesSourceTranslationStaticIp", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for NatRulesSourceTranslationStaticIp ---
func unpackNatRulesSourceTranslationStaticIpListToSdk(ctx context.Context, list types.List) ([]network_services.NatRulesSourceTranslationStaticIp, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.NatRulesSourceTranslationStaticIp")
	diags := diag.Diagnostics{}
	var data []models.NatRulesSourceTranslationStaticIp
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.NatRulesSourceTranslationStaticIp, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.NatRulesSourceTranslationStaticIp{}.AttrTypes(), &item)
		unpacked, d := unpackNatRulesSourceTranslationStaticIpToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.NatRulesSourceTranslationStaticIp", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for NatRulesSourceTranslationStaticIp ---
func packNatRulesSourceTranslationStaticIpListFromSdk(ctx context.Context, sdks []network_services.NatRulesSourceTranslationStaticIp) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.NatRulesSourceTranslationStaticIp")
	diags := diag.Diagnostics{}
	var data []models.NatRulesSourceTranslationStaticIp

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.NatRulesSourceTranslationStaticIp
		obj, d := packNatRulesSourceTranslationStaticIpFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.NatRulesSourceTranslationStaticIp{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.NatRulesSourceTranslationStaticIp", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.NatRulesSourceTranslationStaticIp{}.AttrType(), data)
}
