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

// --- Unpacker for AntiSpywareProfiles ---
func unpackAntiSpywareProfilesToSdk(ctx context.Context, obj types.Object) (*security_services.AntiSpywareProfiles, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.AntiSpywareProfiles", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.AntiSpywareProfiles
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.AntiSpywareProfiles
	var d diag.Diagnostics
	// Handling Primitives
	if !model.CloudInlineAnalysis.IsNull() && !model.CloudInlineAnalysis.IsUnknown() {
		sdk.CloudInlineAnalysis = model.CloudInlineAnalysis.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "CloudInlineAnalysis", "value": *sdk.CloudInlineAnalysis})
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

	// Handling Lists
	if !model.InlineExceptionEdlUrl.IsNull() && !model.InlineExceptionEdlUrl.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field InlineExceptionEdlUrl")
		diags.Append(model.InlineExceptionEdlUrl.ElementsAs(ctx, &sdk.InlineExceptionEdlUrl, false)...)
	}

	// Handling Lists
	if !model.InlineExceptionIpAddress.IsNull() && !model.InlineExceptionIpAddress.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field InlineExceptionIpAddress")
		diags.Append(model.InlineExceptionIpAddress.ElementsAs(ctx, &sdk.InlineExceptionIpAddress, false)...)
	}

	// Handling Lists
	if !model.MicaEngineSpywareEnabled.IsNull() && !model.MicaEngineSpywareEnabled.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field MicaEngineSpywareEnabled")
		unpacked, d := unpackAntiSpywareProfilesMicaEngineSpywareEnabledInnerListToSdk(ctx, model.MicaEngineSpywareEnabled)
		diags.Append(d...)
		sdk.MicaEngineSpywareEnabled = unpacked
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	// Handling Lists
	if !model.Rules.IsNull() && !model.Rules.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field Rules")
		unpacked, d := unpackAntiSpywareProfilesRulesInnerListToSdk(ctx, model.Rules)
		diags.Append(d...)
		sdk.Rules = unpacked
	}

	// Handling Primitives
	if !model.Snippet.IsNull() && !model.Snippet.IsUnknown() {
		sdk.Snippet = model.Snippet.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	}

	// Handling Lists
	if !model.ThreatException.IsNull() && !model.ThreatException.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field ThreatException")
		unpacked, d := unpackAntiSpywareProfilesThreatExceptionInnerListToSdk(ctx, model.ThreatException)
		diags.Append(d...)
		sdk.ThreatException = unpacked
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.AntiSpywareProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for AntiSpywareProfiles ---
func packAntiSpywareProfilesFromSdk(ctx context.Context, sdk security_services.AntiSpywareProfiles) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.AntiSpywareProfiles", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.AntiSpywareProfiles
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.CloudInlineAnalysis != nil {
		model.CloudInlineAnalysis = basetypes.NewBoolValue(*sdk.CloudInlineAnalysis)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "CloudInlineAnalysis", "value": *sdk.CloudInlineAnalysis})
	} else {
		model.CloudInlineAnalysis = basetypes.NewBoolNull()
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
	// Handling Lists
	if sdk.InlineExceptionEdlUrl != nil {
		tflog.Debug(ctx, "Packing list of primitives for field InlineExceptionEdlUrl")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.InlineExceptionEdlUrl, d = basetypes.NewListValueFrom(ctx, elemType, sdk.InlineExceptionEdlUrl)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.InlineExceptionEdlUrl = basetypes.NewListNull(elemType)
	}
	// Handling Lists
	if sdk.InlineExceptionIpAddress != nil {
		tflog.Debug(ctx, "Packing list of primitives for field InlineExceptionIpAddress")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.InlineExceptionIpAddress, d = basetypes.NewListValueFrom(ctx, elemType, sdk.InlineExceptionIpAddress)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.InlineExceptionIpAddress = basetypes.NewListNull(elemType)
	}
	// Handling Lists
	if sdk.MicaEngineSpywareEnabled != nil {
		tflog.Debug(ctx, "Packing list of objects for field MicaEngineSpywareEnabled")
		packed, d := packAntiSpywareProfilesMicaEngineSpywareEnabledInnerListFromSdk(ctx, sdk.MicaEngineSpywareEnabled)
		diags.Append(d...)
		model.MicaEngineSpywareEnabled = packed
	} else {
		model.MicaEngineSpywareEnabled = basetypes.NewListNull(models.AntiSpywareProfilesMicaEngineSpywareEnabledInner{}.AttrType())
	}
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	// Handling Lists
	if sdk.Rules != nil {
		tflog.Debug(ctx, "Packing list of objects for field Rules")
		packed, d := packAntiSpywareProfilesRulesInnerListFromSdk(ctx, sdk.Rules)
		diags.Append(d...)
		model.Rules = packed
	} else {
		model.Rules = basetypes.NewListNull(models.AntiSpywareProfilesRulesInner{}.AttrType())
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
	if sdk.ThreatException != nil {
		tflog.Debug(ctx, "Packing list of objects for field ThreatException")
		packed, d := packAntiSpywareProfilesThreatExceptionInnerListFromSdk(ctx, sdk.ThreatException)
		diags.Append(d...)
		model.ThreatException = packed
	} else {
		model.ThreatException = basetypes.NewListNull(models.AntiSpywareProfilesThreatExceptionInner{}.AttrType())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.AntiSpywareProfiles{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.AntiSpywareProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for AntiSpywareProfiles ---
func unpackAntiSpywareProfilesListToSdk(ctx context.Context, list types.List) ([]security_services.AntiSpywareProfiles, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.AntiSpywareProfiles")
	diags := diag.Diagnostics{}
	var data []models.AntiSpywareProfiles
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.AntiSpywareProfiles, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.AntiSpywareProfiles{}.AttrTypes(), &item)
		unpacked, d := unpackAntiSpywareProfilesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.AntiSpywareProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for AntiSpywareProfiles ---
func packAntiSpywareProfilesListFromSdk(ctx context.Context, sdks []security_services.AntiSpywareProfiles) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.AntiSpywareProfiles")
	diags := diag.Diagnostics{}
	var data []models.AntiSpywareProfiles

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.AntiSpywareProfiles
		obj, d := packAntiSpywareProfilesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.AntiSpywareProfiles{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.AntiSpywareProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.AntiSpywareProfiles{}.AttrType(), data)
}

// --- Unpacker for AntiSpywareProfilesMicaEngineSpywareEnabledInner ---
func unpackAntiSpywareProfilesMicaEngineSpywareEnabledInnerToSdk(ctx context.Context, obj types.Object) (*security_services.AntiSpywareProfilesMicaEngineSpywareEnabledInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.AntiSpywareProfilesMicaEngineSpywareEnabledInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.AntiSpywareProfilesMicaEngineSpywareEnabledInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.AntiSpywareProfilesMicaEngineSpywareEnabledInner
	var d diag.Diagnostics
	// Handling Primitives
	if !model.InlinePolicyAction.IsNull() && !model.InlinePolicyAction.IsUnknown() {
		sdk.InlinePolicyAction = model.InlinePolicyAction.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "InlinePolicyAction", "value": *sdk.InlinePolicyAction})
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.AntiSpywareProfilesMicaEngineSpywareEnabledInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for AntiSpywareProfilesMicaEngineSpywareEnabledInner ---
func packAntiSpywareProfilesMicaEngineSpywareEnabledInnerFromSdk(ctx context.Context, sdk security_services.AntiSpywareProfilesMicaEngineSpywareEnabledInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.AntiSpywareProfilesMicaEngineSpywareEnabledInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.AntiSpywareProfilesMicaEngineSpywareEnabledInner
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.InlinePolicyAction != nil {
		model.InlinePolicyAction = basetypes.NewStringValue(*sdk.InlinePolicyAction)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "InlinePolicyAction", "value": *sdk.InlinePolicyAction})
	} else {
		model.InlinePolicyAction = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Name != nil {
		model.Name = basetypes.NewStringValue(*sdk.Name)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
	} else {
		model.Name = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.AntiSpywareProfilesMicaEngineSpywareEnabledInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.AntiSpywareProfilesMicaEngineSpywareEnabledInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for AntiSpywareProfilesMicaEngineSpywareEnabledInner ---
func unpackAntiSpywareProfilesMicaEngineSpywareEnabledInnerListToSdk(ctx context.Context, list types.List) ([]security_services.AntiSpywareProfilesMicaEngineSpywareEnabledInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.AntiSpywareProfilesMicaEngineSpywareEnabledInner")
	diags := diag.Diagnostics{}
	var data []models.AntiSpywareProfilesMicaEngineSpywareEnabledInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.AntiSpywareProfilesMicaEngineSpywareEnabledInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.AntiSpywareProfilesMicaEngineSpywareEnabledInner{}.AttrTypes(), &item)
		unpacked, d := unpackAntiSpywareProfilesMicaEngineSpywareEnabledInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.AntiSpywareProfilesMicaEngineSpywareEnabledInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for AntiSpywareProfilesMicaEngineSpywareEnabledInner ---
func packAntiSpywareProfilesMicaEngineSpywareEnabledInnerListFromSdk(ctx context.Context, sdks []security_services.AntiSpywareProfilesMicaEngineSpywareEnabledInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.AntiSpywareProfilesMicaEngineSpywareEnabledInner")
	diags := diag.Diagnostics{}
	var data []models.AntiSpywareProfilesMicaEngineSpywareEnabledInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.AntiSpywareProfilesMicaEngineSpywareEnabledInner
		obj, d := packAntiSpywareProfilesMicaEngineSpywareEnabledInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.AntiSpywareProfilesMicaEngineSpywareEnabledInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.AntiSpywareProfilesMicaEngineSpywareEnabledInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.AntiSpywareProfilesMicaEngineSpywareEnabledInner{}.AttrType(), data)
}

// --- Unpacker for AntiSpywareProfilesRulesInner ---
func unpackAntiSpywareProfilesRulesInnerToSdk(ctx context.Context, obj types.Object) (*security_services.AntiSpywareProfilesRulesInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.AntiSpywareProfilesRulesInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.AntiSpywareProfilesRulesInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.AntiSpywareProfilesRulesInner
	var d diag.Diagnostics
	// Handling Objects
	if !model.Action.IsNull() && !model.Action.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Action")
		unpacked, d := unpackAntiSpywareProfilesRulesInnerActionToSdk(ctx, model.Action)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Action"})
		}
		if unpacked != nil {
			sdk.Action = unpacked
		}
	}

	// Handling Primitives
	if !model.Category.IsNull() && !model.Category.IsUnknown() {
		sdk.Category = model.Category.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Category", "value": *sdk.Category})
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
	}

	// Handling Primitives
	if !model.PacketCapture.IsNull() && !model.PacketCapture.IsUnknown() {
		sdk.PacketCapture = model.PacketCapture.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "PacketCapture", "value": *sdk.PacketCapture})
	}

	// Handling Lists
	if !model.Severity.IsNull() && !model.Severity.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Severity")
		diags.Append(model.Severity.ElementsAs(ctx, &sdk.Severity, false)...)
	}

	// Handling Primitives
	if !model.ThreatName.IsNull() && !model.ThreatName.IsUnknown() {
		sdk.ThreatName = model.ThreatName.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "ThreatName", "value": *sdk.ThreatName})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.AntiSpywareProfilesRulesInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for AntiSpywareProfilesRulesInner ---
func packAntiSpywareProfilesRulesInnerFromSdk(ctx context.Context, sdk security_services.AntiSpywareProfilesRulesInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.AntiSpywareProfilesRulesInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.AntiSpywareProfilesRulesInner
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Action != nil {
		tflog.Debug(ctx, "Packing nested object for field Action")
		packed, d := packAntiSpywareProfilesRulesInnerActionFromSdk(ctx, *sdk.Action)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Action"})
		}
		model.Action = packed
	} else {
		model.Action = basetypes.NewObjectNull(models.AntiSpywareProfilesRulesInnerAction{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Category != nil {
		model.Category = basetypes.NewStringValue(*sdk.Category)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Category", "value": *sdk.Category})
	} else {
		model.Category = basetypes.NewStringNull()
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
	if sdk.PacketCapture != nil {
		model.PacketCapture = basetypes.NewStringValue(*sdk.PacketCapture)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "PacketCapture", "value": *sdk.PacketCapture})
	} else {
		model.PacketCapture = basetypes.NewStringNull()
	}
	// Handling Lists
	if sdk.Severity != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Severity")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Severity, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Severity)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Severity = basetypes.NewListNull(elemType)
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.ThreatName != nil {
		model.ThreatName = basetypes.NewStringValue(*sdk.ThreatName)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "ThreatName", "value": *sdk.ThreatName})
	} else {
		model.ThreatName = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.AntiSpywareProfilesRulesInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.AntiSpywareProfilesRulesInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for AntiSpywareProfilesRulesInner ---
func unpackAntiSpywareProfilesRulesInnerListToSdk(ctx context.Context, list types.List) ([]security_services.AntiSpywareProfilesRulesInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.AntiSpywareProfilesRulesInner")
	diags := diag.Diagnostics{}
	var data []models.AntiSpywareProfilesRulesInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.AntiSpywareProfilesRulesInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.AntiSpywareProfilesRulesInner{}.AttrTypes(), &item)
		unpacked, d := unpackAntiSpywareProfilesRulesInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.AntiSpywareProfilesRulesInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for AntiSpywareProfilesRulesInner ---
func packAntiSpywareProfilesRulesInnerListFromSdk(ctx context.Context, sdks []security_services.AntiSpywareProfilesRulesInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.AntiSpywareProfilesRulesInner")
	diags := diag.Diagnostics{}
	var data []models.AntiSpywareProfilesRulesInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.AntiSpywareProfilesRulesInner
		obj, d := packAntiSpywareProfilesRulesInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.AntiSpywareProfilesRulesInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.AntiSpywareProfilesRulesInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.AntiSpywareProfilesRulesInner{}.AttrType(), data)
}

// --- Unpacker for AntiSpywareProfilesThreatExceptionInner ---
func unpackAntiSpywareProfilesThreatExceptionInnerToSdk(ctx context.Context, obj types.Object) (*security_services.AntiSpywareProfilesThreatExceptionInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.AntiSpywareProfilesThreatExceptionInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.AntiSpywareProfilesThreatExceptionInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.AntiSpywareProfilesThreatExceptionInner
	var d diag.Diagnostics
	// Handling Objects
	if !model.Action.IsNull() && !model.Action.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Action")
		unpacked, d := unpackAntiSpywareProfilesThreatExceptionInnerActionToSdk(ctx, model.Action)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Action"})
		}
		if unpacked != nil {
			sdk.Action = unpacked
		}
	}

	// Handling Lists
	if !model.ExemptIp.IsNull() && !model.ExemptIp.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field ExemptIp")
		unpacked, d := unpackAntiSpywareProfilesThreatExceptionInnerExemptIpInnerListToSdk(ctx, model.ExemptIp)
		diags.Append(d...)
		sdk.ExemptIp = unpacked
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
	}

	// Handling Primitives
	if !model.Notes.IsNull() && !model.Notes.IsUnknown() {
		sdk.Notes = model.Notes.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Notes", "value": *sdk.Notes})
	}

	// Handling Primitives
	if !model.PacketCapture.IsNull() && !model.PacketCapture.IsUnknown() {
		sdk.PacketCapture = model.PacketCapture.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "PacketCapture", "value": *sdk.PacketCapture})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.AntiSpywareProfilesThreatExceptionInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for AntiSpywareProfilesThreatExceptionInner ---
func packAntiSpywareProfilesThreatExceptionInnerFromSdk(ctx context.Context, sdk security_services.AntiSpywareProfilesThreatExceptionInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.AntiSpywareProfilesThreatExceptionInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.AntiSpywareProfilesThreatExceptionInner
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Action != nil {
		tflog.Debug(ctx, "Packing nested object for field Action")
		packed, d := packAntiSpywareProfilesThreatExceptionInnerActionFromSdk(ctx, *sdk.Action)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Action"})
		}
		model.Action = packed
	} else {
		model.Action = basetypes.NewObjectNull(models.AntiSpywareProfilesThreatExceptionInnerAction{}.AttrTypes())
	}
	// Handling Lists
	if sdk.ExemptIp != nil {
		tflog.Debug(ctx, "Packing list of objects for field ExemptIp")
		packed, d := packAntiSpywareProfilesThreatExceptionInnerExemptIpInnerListFromSdk(ctx, sdk.ExemptIp)
		diags.Append(d...)
		model.ExemptIp = packed
	} else {
		model.ExemptIp = basetypes.NewListNull(models.AntiSpywareProfilesThreatExceptionInnerExemptIpInner{}.AttrType())
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
	if sdk.Notes != nil {
		model.Notes = basetypes.NewStringValue(*sdk.Notes)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Notes", "value": *sdk.Notes})
	} else {
		model.Notes = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.PacketCapture != nil {
		model.PacketCapture = basetypes.NewStringValue(*sdk.PacketCapture)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "PacketCapture", "value": *sdk.PacketCapture})
	} else {
		model.PacketCapture = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.AntiSpywareProfilesThreatExceptionInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.AntiSpywareProfilesThreatExceptionInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for AntiSpywareProfilesThreatExceptionInner ---
func unpackAntiSpywareProfilesThreatExceptionInnerListToSdk(ctx context.Context, list types.List) ([]security_services.AntiSpywareProfilesThreatExceptionInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.AntiSpywareProfilesThreatExceptionInner")
	diags := diag.Diagnostics{}
	var data []models.AntiSpywareProfilesThreatExceptionInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.AntiSpywareProfilesThreatExceptionInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.AntiSpywareProfilesThreatExceptionInner{}.AttrTypes(), &item)
		unpacked, d := unpackAntiSpywareProfilesThreatExceptionInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.AntiSpywareProfilesThreatExceptionInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for AntiSpywareProfilesThreatExceptionInner ---
func packAntiSpywareProfilesThreatExceptionInnerListFromSdk(ctx context.Context, sdks []security_services.AntiSpywareProfilesThreatExceptionInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.AntiSpywareProfilesThreatExceptionInner")
	diags := diag.Diagnostics{}
	var data []models.AntiSpywareProfilesThreatExceptionInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.AntiSpywareProfilesThreatExceptionInner
		obj, d := packAntiSpywareProfilesThreatExceptionInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.AntiSpywareProfilesThreatExceptionInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.AntiSpywareProfilesThreatExceptionInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.AntiSpywareProfilesThreatExceptionInner{}.AttrType(), data)
}
