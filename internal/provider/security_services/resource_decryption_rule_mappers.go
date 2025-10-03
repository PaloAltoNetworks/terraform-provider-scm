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

// --- Unpacker for DecryptionRules ---
func unpackDecryptionRulesToSdk(ctx context.Context, obj types.Object) (*security_services.DecryptionRules, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.DecryptionRules", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.DecryptionRules
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.DecryptionRules
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Action.IsNull() && !model.Action.IsUnknown() {
		sdk.Action = model.Action.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Action", "value": sdk.Action})
	}

	// Handling Lists
	if !model.Category.IsNull() && !model.Category.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Category")
		diags.Append(model.Category.ElementsAs(ctx, &sdk.Category, false)...)
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

	// Handling Lists
	if !model.DestinationHip.IsNull() && !model.DestinationHip.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field DestinationHip")
		diags.Append(model.DestinationHip.ElementsAs(ctx, &sdk.DestinationHip, false)...)
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
		sdk.Id = model.Id.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Id", "value": sdk.Id})
	}

	// Handling Primitives
	if !model.LogFail.IsNull() && !model.LogFail.IsUnknown() {
		sdk.LogFail = model.LogFail.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "LogFail", "value": *sdk.LogFail})
	}

	// Handling Primitives
	if !model.LogSetting.IsNull() && !model.LogSetting.IsUnknown() {
		sdk.LogSetting = model.LogSetting.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "LogSetting", "value": *sdk.LogSetting})
	}

	// Handling Primitives
	if !model.LogSuccess.IsNull() && !model.LogSuccess.IsUnknown() {
		sdk.LogSuccess = model.LogSuccess.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "LogSuccess", "value": *sdk.LogSuccess})
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
	if !model.Profile.IsNull() && !model.Profile.IsUnknown() {
		sdk.Profile = model.Profile.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Profile", "value": *sdk.Profile})
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
	if !model.SourceHip.IsNull() && !model.SourceHip.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field SourceHip")
		diags.Append(model.SourceHip.ElementsAs(ctx, &sdk.SourceHip, false)...)
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

	// Handling Objects
	if !model.Type.IsNull() && !model.Type.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Type")
		unpacked, d := unpackDecryptionRulesTypeToSdk(ctx, model.Type)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Type"})
		}
		if unpacked != nil {
			sdk.Type = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.DecryptionRules", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for DecryptionRules ---
func packDecryptionRulesFromSdk(ctx context.Context, sdk security_services.DecryptionRules) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.DecryptionRules", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.DecryptionRules
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.Action = basetypes.NewStringValue(sdk.Action)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Action", "value": sdk.Action})
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
	// Handling Lists
	if sdk.DestinationHip != nil {
		tflog.Debug(ctx, "Packing list of primitives for field DestinationHip")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.DestinationHip, d = basetypes.NewListValueFrom(ctx, elemType, sdk.DestinationHip)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.DestinationHip = basetypes.NewListNull(elemType)
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
	model.Id = basetypes.NewStringValue(sdk.Id)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Id", "value": sdk.Id})
	// Handling Primitives
	// Standard primitive packing
	if sdk.LogFail != nil {
		model.LogFail = basetypes.NewBoolValue(*sdk.LogFail)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "LogFail", "value": *sdk.LogFail})
	} else {
		model.LogFail = basetypes.NewBoolNull()
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
	if sdk.LogSuccess != nil {
		model.LogSuccess = basetypes.NewBoolValue(*sdk.LogSuccess)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "LogSuccess", "value": *sdk.LogSuccess})
	} else {
		model.LogSuccess = basetypes.NewBoolNull()
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
	if sdk.Profile != nil {
		model.Profile = basetypes.NewStringValue(*sdk.Profile)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Profile", "value": *sdk.Profile})
	} else {
		model.Profile = basetypes.NewStringNull()
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
	if sdk.SourceHip != nil {
		tflog.Debug(ctx, "Packing list of primitives for field SourceHip")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.SourceHip, d = basetypes.NewListValueFrom(ctx, elemType, sdk.SourceHip)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.SourceHip = basetypes.NewListNull(elemType)
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
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Type != nil {
		tflog.Debug(ctx, "Packing nested object for field Type")
		packed, d := packDecryptionRulesTypeFromSdk(ctx, *sdk.Type)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Type"})
		}
		model.Type = packed
	} else {
		model.Type = basetypes.NewObjectNull(models.DecryptionRulesType{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.DecryptionRules{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.DecryptionRules", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for DecryptionRules ---
func unpackDecryptionRulesListToSdk(ctx context.Context, list types.List) ([]security_services.DecryptionRules, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.DecryptionRules")
	diags := diag.Diagnostics{}
	var data []models.DecryptionRules
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.DecryptionRules, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.DecryptionRules{}.AttrTypes(), &item)
		unpacked, d := unpackDecryptionRulesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.DecryptionRules", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for DecryptionRules ---
func packDecryptionRulesListFromSdk(ctx context.Context, sdks []security_services.DecryptionRules) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.DecryptionRules")
	diags := diag.Diagnostics{}
	var data []models.DecryptionRules

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.DecryptionRules
		obj, d := packDecryptionRulesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.DecryptionRules{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.DecryptionRules", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.DecryptionRules{}.AttrType(), data)
}

// --- Unpacker for DecryptionRulesType ---
func unpackDecryptionRulesTypeToSdk(ctx context.Context, obj types.Object) (*security_services.DecryptionRulesType, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.DecryptionRulesType", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.DecryptionRulesType
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.DecryptionRulesType
	var d diag.Diagnostics
	// Handling Typeless Objects
	if !model.SslForwardProxy.IsNull() && !model.SslForwardProxy.IsUnknown() {
		tflog.Debug(ctx, "Unpacking typeless object for field SslForwardProxy")
		sdk.SslForwardProxy = make(map[string]interface{})
	}

	// Handling Primitives
	if !model.SslInboundInspection.IsNull() && !model.SslInboundInspection.IsUnknown() {
		sdk.SslInboundInspection = model.SslInboundInspection.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "SslInboundInspection", "value": *sdk.SslInboundInspection})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.DecryptionRulesType", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for DecryptionRulesType ---
func packDecryptionRulesTypeFromSdk(ctx context.Context, sdk security_services.DecryptionRulesType) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.DecryptionRulesType", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.DecryptionRulesType
	var d diag.Diagnostics
	// Handling Objects
	// This is a marker object (e.g. CHAP: {}). We just need to create an empty, non-null object.
	if sdk.SslForwardProxy != nil && !reflect.ValueOf(sdk.SslForwardProxy).IsNil() {
		tflog.Debug(ctx, "Packing typeless object for field SslForwardProxy")
		var d diag.Diagnostics
		// Create an empty object with no attributes, which signifies its presence.
		model.SslForwardProxy, d = basetypes.NewObjectValue(map[string]attr.Type{}, map[string]attr.Value{})
		diags.Append(d...)
	} else {
		// Since this field is part of a oneOf, being nil means it's not selected.
		// We make the object null with an empty attribute map.
		model.SslForwardProxy = basetypes.NewObjectNull(map[string]attr.Type{})
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.SslInboundInspection != nil {
		model.SslInboundInspection = basetypes.NewStringValue(*sdk.SslInboundInspection)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "SslInboundInspection", "value": *sdk.SslInboundInspection})
	} else {
		model.SslInboundInspection = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.DecryptionRulesType{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.DecryptionRulesType", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for DecryptionRulesType ---
func unpackDecryptionRulesTypeListToSdk(ctx context.Context, list types.List) ([]security_services.DecryptionRulesType, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.DecryptionRulesType")
	diags := diag.Diagnostics{}
	var data []models.DecryptionRulesType
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.DecryptionRulesType, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.DecryptionRulesType{}.AttrTypes(), &item)
		unpacked, d := unpackDecryptionRulesTypeToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.DecryptionRulesType", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for DecryptionRulesType ---
func packDecryptionRulesTypeListFromSdk(ctx context.Context, sdks []security_services.DecryptionRulesType) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.DecryptionRulesType")
	diags := diag.Diagnostics{}
	var data []models.DecryptionRulesType

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.DecryptionRulesType
		obj, d := packDecryptionRulesTypeFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.DecryptionRulesType{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.DecryptionRulesType", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.DecryptionRulesType{}.AttrType(), data)
}
