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

// --- Unpacker for WildfireAntiVirusProfiles ---
func unpackWildfireAntiVirusProfilesToSdk(ctx context.Context, obj types.Object) (*security_services.WildfireAntiVirusProfiles, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.WildfireAntiVirusProfiles", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.WildfireAntiVirusProfiles
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.WildfireAntiVirusProfiles
	var d diag.Diagnostics

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
		sdk.Id = model.Id.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Id", "value": *sdk.Id})
	}

	// Handling Lists
	if !model.MlavException.IsNull() && !model.MlavException.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field MlavException")
		unpacked, d := unpackWildfireAntiVirusProfilesMlavExceptionInnerListToSdk(ctx, model.MlavException)
		diags.Append(d...)
		sdk.MlavException = unpacked
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	// Handling Primitives
	if !model.PacketCapture.IsNull() && !model.PacketCapture.IsUnknown() {
		sdk.PacketCapture = model.PacketCapture.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "PacketCapture", "value": *sdk.PacketCapture})
	}

	// Handling Lists
	if !model.Rules.IsNull() && !model.Rules.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field Rules")
		unpacked, d := unpackWildfireAntiVirusProfilesRulesInnerListToSdk(ctx, model.Rules)
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
		unpacked, d := unpackWildfireAntiVirusProfilesThreatExceptionInnerListToSdk(ctx, model.ThreatException)
		diags.Append(d...)
		sdk.ThreatException = unpacked
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.WildfireAntiVirusProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for WildfireAntiVirusProfiles ---
func packWildfireAntiVirusProfilesFromSdk(ctx context.Context, sdk security_services.WildfireAntiVirusProfiles) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.WildfireAntiVirusProfiles", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.WildfireAntiVirusProfiles
	var d diag.Diagnostics
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
	if sdk.Id != nil {
		model.Id = basetypes.NewStringValue(*sdk.Id)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Id", "value": *sdk.Id})
	} else {
		model.Id = basetypes.NewStringNull()
	}
	// Handling Lists
	if sdk.MlavException != nil {
		tflog.Debug(ctx, "Packing list of objects for field MlavException")
		packed, d := packWildfireAntiVirusProfilesMlavExceptionInnerListFromSdk(ctx, sdk.MlavException)
		diags.Append(d...)
		model.MlavException = packed
	} else {
		model.MlavException = basetypes.NewListNull(models.WildfireAntiVirusProfilesMlavExceptionInner{}.AttrType())
	}
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	// Handling Primitives
	// Standard primitive packing
	if sdk.PacketCapture != nil {
		model.PacketCapture = basetypes.NewBoolValue(*sdk.PacketCapture)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "PacketCapture", "value": *sdk.PacketCapture})
	} else {
		model.PacketCapture = basetypes.NewBoolNull()
	}
	// Handling Lists
	if sdk.Rules != nil {
		tflog.Debug(ctx, "Packing list of objects for field Rules")
		packed, d := packWildfireAntiVirusProfilesRulesInnerListFromSdk(ctx, sdk.Rules)
		diags.Append(d...)
		model.Rules = packed
	} else {
		model.Rules = basetypes.NewListNull(models.WildfireAntiVirusProfilesRulesInner{}.AttrType())
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
		packed, d := packWildfireAntiVirusProfilesThreatExceptionInnerListFromSdk(ctx, sdk.ThreatException)
		diags.Append(d...)
		model.ThreatException = packed
	} else {
		model.ThreatException = basetypes.NewListNull(models.WildfireAntiVirusProfilesThreatExceptionInner{}.AttrType())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.WildfireAntiVirusProfiles{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.WildfireAntiVirusProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for WildfireAntiVirusProfiles ---
func unpackWildfireAntiVirusProfilesListToSdk(ctx context.Context, list types.List) ([]security_services.WildfireAntiVirusProfiles, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.WildfireAntiVirusProfiles")
	diags := diag.Diagnostics{}
	var data []models.WildfireAntiVirusProfiles
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.WildfireAntiVirusProfiles, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.WildfireAntiVirusProfiles{}.AttrTypes(), &item)
		unpacked, d := unpackWildfireAntiVirusProfilesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.WildfireAntiVirusProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for WildfireAntiVirusProfiles ---
func packWildfireAntiVirusProfilesListFromSdk(ctx context.Context, sdks []security_services.WildfireAntiVirusProfiles) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.WildfireAntiVirusProfiles")
	diags := diag.Diagnostics{}
	var data []models.WildfireAntiVirusProfiles

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.WildfireAntiVirusProfiles
		obj, d := packWildfireAntiVirusProfilesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.WildfireAntiVirusProfiles{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.WildfireAntiVirusProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.WildfireAntiVirusProfiles{}.AttrType(), data)
}

// --- Unpacker for WildfireAntiVirusProfilesMlavExceptionInner ---
func unpackWildfireAntiVirusProfilesMlavExceptionInnerToSdk(ctx context.Context, obj types.Object) (*security_services.WildfireAntiVirusProfilesMlavExceptionInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.WildfireAntiVirusProfilesMlavExceptionInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.WildfireAntiVirusProfilesMlavExceptionInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.WildfireAntiVirusProfilesMlavExceptionInner
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Description.IsNull() && !model.Description.IsUnknown() {
		sdk.Description = model.Description.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Description", "value": *sdk.Description})
	}

	// Handling Primitives
	if !model.Filename.IsNull() && !model.Filename.IsUnknown() {
		sdk.Filename = model.Filename.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Filename", "value": *sdk.Filename})
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.WildfireAntiVirusProfilesMlavExceptionInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for WildfireAntiVirusProfilesMlavExceptionInner ---
func packWildfireAntiVirusProfilesMlavExceptionInnerFromSdk(ctx context.Context, sdk security_services.WildfireAntiVirusProfilesMlavExceptionInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.WildfireAntiVirusProfilesMlavExceptionInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.WildfireAntiVirusProfilesMlavExceptionInner
	var d diag.Diagnostics
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
	if sdk.Filename != nil {
		model.Filename = basetypes.NewStringValue(*sdk.Filename)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Filename", "value": *sdk.Filename})
	} else {
		model.Filename = basetypes.NewStringNull()
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

	obj, d := types.ObjectValueFrom(ctx, models.WildfireAntiVirusProfilesMlavExceptionInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.WildfireAntiVirusProfilesMlavExceptionInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for WildfireAntiVirusProfilesMlavExceptionInner ---
func unpackWildfireAntiVirusProfilesMlavExceptionInnerListToSdk(ctx context.Context, list types.List) ([]security_services.WildfireAntiVirusProfilesMlavExceptionInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.WildfireAntiVirusProfilesMlavExceptionInner")
	diags := diag.Diagnostics{}
	var data []models.WildfireAntiVirusProfilesMlavExceptionInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.WildfireAntiVirusProfilesMlavExceptionInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.WildfireAntiVirusProfilesMlavExceptionInner{}.AttrTypes(), &item)
		unpacked, d := unpackWildfireAntiVirusProfilesMlavExceptionInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.WildfireAntiVirusProfilesMlavExceptionInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for WildfireAntiVirusProfilesMlavExceptionInner ---
func packWildfireAntiVirusProfilesMlavExceptionInnerListFromSdk(ctx context.Context, sdks []security_services.WildfireAntiVirusProfilesMlavExceptionInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.WildfireAntiVirusProfilesMlavExceptionInner")
	diags := diag.Diagnostics{}
	var data []models.WildfireAntiVirusProfilesMlavExceptionInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.WildfireAntiVirusProfilesMlavExceptionInner
		obj, d := packWildfireAntiVirusProfilesMlavExceptionInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.WildfireAntiVirusProfilesMlavExceptionInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.WildfireAntiVirusProfilesMlavExceptionInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.WildfireAntiVirusProfilesMlavExceptionInner{}.AttrType(), data)
}

// --- Unpacker for WildfireAntiVirusProfilesRulesInner ---
func unpackWildfireAntiVirusProfilesRulesInnerToSdk(ctx context.Context, obj types.Object) (*security_services.WildfireAntiVirusProfilesRulesInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.WildfireAntiVirusProfilesRulesInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.WildfireAntiVirusProfilesRulesInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.WildfireAntiVirusProfilesRulesInner
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Analysis.IsNull() && !model.Analysis.IsUnknown() {
		sdk.Analysis = model.Analysis.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Analysis", "value": *sdk.Analysis})
	}

	// Handling Lists
	if !model.Application.IsNull() && !model.Application.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Application")
		diags.Append(model.Application.ElementsAs(ctx, &sdk.Application, false)...)
	}

	// Handling Primitives
	if !model.Direction.IsNull() && !model.Direction.IsUnknown() {
		sdk.Direction = model.Direction.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Direction", "value": *sdk.Direction})
	}

	// Handling Lists
	if !model.FileType.IsNull() && !model.FileType.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field FileType")
		diags.Append(model.FileType.ElementsAs(ctx, &sdk.FileType, false)...)
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.WildfireAntiVirusProfilesRulesInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for WildfireAntiVirusProfilesRulesInner ---
func packWildfireAntiVirusProfilesRulesInnerFromSdk(ctx context.Context, sdk security_services.WildfireAntiVirusProfilesRulesInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.WildfireAntiVirusProfilesRulesInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.WildfireAntiVirusProfilesRulesInner
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Analysis != nil {
		model.Analysis = basetypes.NewStringValue(*sdk.Analysis)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Analysis", "value": *sdk.Analysis})
	} else {
		model.Analysis = basetypes.NewStringNull()
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
	if sdk.Direction != nil {
		model.Direction = basetypes.NewStringValue(*sdk.Direction)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Direction", "value": *sdk.Direction})
	} else {
		model.Direction = basetypes.NewStringNull()
	}
	// Handling Lists
	if sdk.FileType != nil {
		tflog.Debug(ctx, "Packing list of primitives for field FileType")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.FileType, d = basetypes.NewListValueFrom(ctx, elemType, sdk.FileType)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.FileType = basetypes.NewListNull(elemType)
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

	obj, d := types.ObjectValueFrom(ctx, models.WildfireAntiVirusProfilesRulesInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.WildfireAntiVirusProfilesRulesInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for WildfireAntiVirusProfilesRulesInner ---
func unpackWildfireAntiVirusProfilesRulesInnerListToSdk(ctx context.Context, list types.List) ([]security_services.WildfireAntiVirusProfilesRulesInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.WildfireAntiVirusProfilesRulesInner")
	diags := diag.Diagnostics{}
	var data []models.WildfireAntiVirusProfilesRulesInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.WildfireAntiVirusProfilesRulesInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.WildfireAntiVirusProfilesRulesInner{}.AttrTypes(), &item)
		unpacked, d := unpackWildfireAntiVirusProfilesRulesInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.WildfireAntiVirusProfilesRulesInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for WildfireAntiVirusProfilesRulesInner ---
func packWildfireAntiVirusProfilesRulesInnerListFromSdk(ctx context.Context, sdks []security_services.WildfireAntiVirusProfilesRulesInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.WildfireAntiVirusProfilesRulesInner")
	diags := diag.Diagnostics{}
	var data []models.WildfireAntiVirusProfilesRulesInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.WildfireAntiVirusProfilesRulesInner
		obj, d := packWildfireAntiVirusProfilesRulesInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.WildfireAntiVirusProfilesRulesInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.WildfireAntiVirusProfilesRulesInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.WildfireAntiVirusProfilesRulesInner{}.AttrType(), data)
}

// --- Unpacker for WildfireAntiVirusProfilesThreatExceptionInner ---
func unpackWildfireAntiVirusProfilesThreatExceptionInnerToSdk(ctx context.Context, obj types.Object) (*security_services.WildfireAntiVirusProfilesThreatExceptionInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.WildfireAntiVirusProfilesThreatExceptionInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.WildfireAntiVirusProfilesThreatExceptionInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.WildfireAntiVirusProfilesThreatExceptionInner
	var d diag.Diagnostics
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

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.WildfireAntiVirusProfilesThreatExceptionInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for WildfireAntiVirusProfilesThreatExceptionInner ---
func packWildfireAntiVirusProfilesThreatExceptionInnerFromSdk(ctx context.Context, sdk security_services.WildfireAntiVirusProfilesThreatExceptionInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.WildfireAntiVirusProfilesThreatExceptionInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.WildfireAntiVirusProfilesThreatExceptionInner
	var d diag.Diagnostics
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
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.WildfireAntiVirusProfilesThreatExceptionInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.WildfireAntiVirusProfilesThreatExceptionInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for WildfireAntiVirusProfilesThreatExceptionInner ---
func unpackWildfireAntiVirusProfilesThreatExceptionInnerListToSdk(ctx context.Context, list types.List) ([]security_services.WildfireAntiVirusProfilesThreatExceptionInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.WildfireAntiVirusProfilesThreatExceptionInner")
	diags := diag.Diagnostics{}
	var data []models.WildfireAntiVirusProfilesThreatExceptionInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.WildfireAntiVirusProfilesThreatExceptionInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.WildfireAntiVirusProfilesThreatExceptionInner{}.AttrTypes(), &item)
		unpacked, d := unpackWildfireAntiVirusProfilesThreatExceptionInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.WildfireAntiVirusProfilesThreatExceptionInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for WildfireAntiVirusProfilesThreatExceptionInner ---
func packWildfireAntiVirusProfilesThreatExceptionInnerListFromSdk(ctx context.Context, sdks []security_services.WildfireAntiVirusProfilesThreatExceptionInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.WildfireAntiVirusProfilesThreatExceptionInner")
	diags := diag.Diagnostics{}
	var data []models.WildfireAntiVirusProfilesThreatExceptionInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.WildfireAntiVirusProfilesThreatExceptionInner
		obj, d := packWildfireAntiVirusProfilesThreatExceptionInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.WildfireAntiVirusProfilesThreatExceptionInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.WildfireAntiVirusProfilesThreatExceptionInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.WildfireAntiVirusProfilesThreatExceptionInner{}.AttrType(), data)
}
