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

// --- Unpacker for SdwanRules ---
func unpackSdwanRulesToSdk(ctx context.Context, obj types.Object) (*network_services.SdwanRules, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.SdwanRules", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.SdwanRules
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.SdwanRules
	var d diag.Diagnostics

	// Handling Objects
	if !model.Action.IsNull() && !model.Action.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Action")
		unpacked, d := unpackSdwanRulesActionToSdk(ctx, model.Action)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Action"})
		}
		if unpacked != nil {
			sdk.Action = *unpacked
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

	// Handling Primitives
	if !model.Disabled.IsNull() && !model.Disabled.IsUnknown() {
		sdk.Disabled = model.Disabled.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Disabled", "value": *sdk.Disabled})
	}

	// Handling Primitives
	if !model.ErrorCorrectionProfile.IsNull() && !model.ErrorCorrectionProfile.IsUnknown() {
		sdk.ErrorCorrectionProfile = model.ErrorCorrectionProfile.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "ErrorCorrectionProfile", "value": *sdk.ErrorCorrectionProfile})
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
	if !model.PathQualityProfile.IsNull() && !model.PathQualityProfile.IsUnknown() {
		sdk.PathQualityProfile = model.PathQualityProfile.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "PathQualityProfile", "value": sdk.PathQualityProfile})
	}

	// Handling Primitives
	if !model.Position.IsNull() && !model.Position.IsUnknown() {
		sdk.Position = model.Position.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Position", "value": sdk.Position})
	}

	// Handling Primitives
	if !model.SaasQualityProfile.IsNull() && !model.SaasQualityProfile.IsUnknown() {
		sdk.SaasQualityProfile = model.SaasQualityProfile.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "SaasQualityProfile", "value": *sdk.SaasQualityProfile})
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

	tflog.Debug(ctx, "Exiting unpack helper for models.SdwanRules", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for SdwanRules ---
func packSdwanRulesFromSdk(ctx context.Context, sdk network_services.SdwanRules) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.SdwanRules", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.SdwanRules
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	// Logic for non-pointer / value-type nested objects
	if !reflect.ValueOf(sdk.Action).IsZero() {
		tflog.Debug(ctx, "Packing nested object for field Action")
		packed, d := packSdwanRulesActionFromSdk(ctx, sdk.Action)
		diags.Append(d...)
		model.Action = packed
	} else {
		model.Action = basetypes.NewObjectNull(models.SdwanRulesAction{}.AttrTypes())
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
	if sdk.ErrorCorrectionProfile != nil {
		model.ErrorCorrectionProfile = basetypes.NewStringValue(*sdk.ErrorCorrectionProfile)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "ErrorCorrectionProfile", "value": *sdk.ErrorCorrectionProfile})
	} else {
		model.ErrorCorrectionProfile = basetypes.NewStringNull()
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
	model.PathQualityProfile = basetypes.NewStringValue(sdk.PathQualityProfile)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "PathQualityProfile", "value": sdk.PathQualityProfile})
	// Handling Primitives
	// Standard primitive packing
	model.Position = basetypes.NewStringValue(sdk.Position)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Position", "value": sdk.Position})
	// Handling Primitives
	// Standard primitive packing
	if sdk.SaasQualityProfile != nil {
		model.SaasQualityProfile = basetypes.NewStringValue(*sdk.SaasQualityProfile)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "SaasQualityProfile", "value": *sdk.SaasQualityProfile})
	} else {
		model.SaasQualityProfile = basetypes.NewStringNull()
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

	obj, d := types.ObjectValueFrom(ctx, models.SdwanRules{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.SdwanRules", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for SdwanRules ---
func unpackSdwanRulesListToSdk(ctx context.Context, list types.List) ([]network_services.SdwanRules, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.SdwanRules")
	diags := diag.Diagnostics{}
	var data []models.SdwanRules
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.SdwanRules, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.SdwanRules{}.AttrTypes(), &item)
		unpacked, d := unpackSdwanRulesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.SdwanRules", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for SdwanRules ---
func packSdwanRulesListFromSdk(ctx context.Context, sdks []network_services.SdwanRules) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.SdwanRules")
	diags := diag.Diagnostics{}
	var data []models.SdwanRules

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.SdwanRules
		obj, d := packSdwanRulesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.SdwanRules{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.SdwanRules", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.SdwanRules{}.AttrType(), data)
}

// --- Unpacker for SdwanRulesAction ---
func unpackSdwanRulesActionToSdk(ctx context.Context, obj types.Object) (*network_services.SdwanRulesAction, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.SdwanRulesAction", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.SdwanRulesAction
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.SdwanRulesAction
	var d diag.Diagnostics
	// Handling Primitives
	if !model.TrafficDistributionProfile.IsNull() && !model.TrafficDistributionProfile.IsUnknown() {
		sdk.TrafficDistributionProfile = model.TrafficDistributionProfile.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "TrafficDistributionProfile", "value": sdk.TrafficDistributionProfile})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.SdwanRulesAction", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for SdwanRulesAction ---
func packSdwanRulesActionFromSdk(ctx context.Context, sdk network_services.SdwanRulesAction) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.SdwanRulesAction", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.SdwanRulesAction
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.TrafficDistributionProfile = basetypes.NewStringValue(sdk.TrafficDistributionProfile)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "TrafficDistributionProfile", "value": sdk.TrafficDistributionProfile})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.SdwanRulesAction{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.SdwanRulesAction", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for SdwanRulesAction ---
func unpackSdwanRulesActionListToSdk(ctx context.Context, list types.List) ([]network_services.SdwanRulesAction, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.SdwanRulesAction")
	diags := diag.Diagnostics{}
	var data []models.SdwanRulesAction
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.SdwanRulesAction, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.SdwanRulesAction{}.AttrTypes(), &item)
		unpacked, d := unpackSdwanRulesActionToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.SdwanRulesAction", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for SdwanRulesAction ---
func packSdwanRulesActionListFromSdk(ctx context.Context, sdks []network_services.SdwanRulesAction) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.SdwanRulesAction")
	diags := diag.Diagnostics{}
	var data []models.SdwanRulesAction

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.SdwanRulesAction
		obj, d := packSdwanRulesActionFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.SdwanRulesAction{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.SdwanRulesAction", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.SdwanRulesAction{}.AttrType(), data)
}
