package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/objects"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/objects"
)

// --- Unpacker for ApplicationFilters ---
func unpackApplicationFiltersToSdk(ctx context.Context, obj types.Object) (*objects.ApplicationFilters, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ApplicationFilters", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ApplicationFilters
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.ApplicationFilters
	var d diag.Diagnostics
	// Handling Lists
	if !model.Category.IsNull() && !model.Category.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Category")
		diags.Append(model.Category.ElementsAs(ctx, &sdk.Category, false)...)
	}

	// Handling Primitives
	if !model.Device.IsNull() && !model.Device.IsUnknown() {
		sdk.Device = model.Device.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Device", "value": *sdk.Device})
	}

	// Handling Primitives
	if !model.Evasive.IsNull() && !model.Evasive.IsUnknown() {
		sdk.Evasive = model.Evasive.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Evasive", "value": *sdk.Evasive})
	}

	// Handling Primitives
	if !model.ExcessiveBandwidthUse.IsNull() && !model.ExcessiveBandwidthUse.IsUnknown() {
		sdk.ExcessiveBandwidthUse = model.ExcessiveBandwidthUse.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "ExcessiveBandwidthUse", "value": *sdk.ExcessiveBandwidthUse})
	}

	// Handling Lists
	if !model.Exclude.IsNull() && !model.Exclude.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Exclude")
		diags.Append(model.Exclude.ElementsAs(ctx, &sdk.Exclude, false)...)
	}

	// Handling Primitives
	if !model.Folder.IsNull() && !model.Folder.IsUnknown() {
		sdk.Folder = model.Folder.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Folder", "value": *sdk.Folder})
	}

	// Handling Primitives
	if !model.HasKnownVulnerabilities.IsNull() && !model.HasKnownVulnerabilities.IsUnknown() {
		sdk.HasKnownVulnerabilities = model.HasKnownVulnerabilities.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "HasKnownVulnerabilities", "value": *sdk.HasKnownVulnerabilities})
	}

	// Handling Primitives
	if !model.Id.IsNull() && !model.Id.IsUnknown() {
		sdk.Id = model.Id.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Id", "value": *sdk.Id})
	}

	// Handling Primitives
	if !model.IsSaas.IsNull() && !model.IsSaas.IsUnknown() {
		sdk.IsSaas = model.IsSaas.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "IsSaas", "value": *sdk.IsSaas})
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	// Handling Primitives
	if !model.NewAppid.IsNull() && !model.NewAppid.IsUnknown() {
		sdk.NewAppid = model.NewAppid.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "NewAppid", "value": *sdk.NewAppid})
	}

	// Handling Primitives
	if !model.Pervasive.IsNull() && !model.Pervasive.IsUnknown() {
		sdk.Pervasive = model.Pervasive.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Pervasive", "value": *sdk.Pervasive})
	}

	// Handling Primitives
	if !model.ProneToMisuse.IsNull() && !model.ProneToMisuse.IsUnknown() {
		sdk.ProneToMisuse = model.ProneToMisuse.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "ProneToMisuse", "value": *sdk.ProneToMisuse})
	}

	// Handling Lists
	if !model.Risk.IsNull() && !model.Risk.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Risk")
		diags.Append(model.Risk.ElementsAs(ctx, &sdk.Risk, false)...)
	}

	// Handling Lists
	if !model.SaasCertifications.IsNull() && !model.SaasCertifications.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field SaasCertifications")
		diags.Append(model.SaasCertifications.ElementsAs(ctx, &sdk.SaasCertifications, false)...)
	}

	// Handling Lists
	if !model.SaasRisk.IsNull() && !model.SaasRisk.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field SaasRisk")
		diags.Append(model.SaasRisk.ElementsAs(ctx, &sdk.SaasRisk, false)...)
	}

	// Handling Primitives
	if !model.Snippet.IsNull() && !model.Snippet.IsUnknown() {
		sdk.Snippet = model.Snippet.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	}

	// Handling Lists
	if !model.Subcategory.IsNull() && !model.Subcategory.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Subcategory")
		diags.Append(model.Subcategory.ElementsAs(ctx, &sdk.Subcategory, false)...)
	}

	// Handling Objects
	if !model.Tagging.IsNull() && !model.Tagging.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Tagging")
		unpacked, d := unpackApplicationFiltersTaggingToSdk(ctx, model.Tagging)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Tagging"})
		}
		if unpacked != nil {
			sdk.Tagging = unpacked
		}
	}

	// Handling Lists
	if !model.Technology.IsNull() && !model.Technology.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Technology")
		diags.Append(model.Technology.ElementsAs(ctx, &sdk.Technology, false)...)
	}

	// Handling Primitives
	if !model.TransfersFiles.IsNull() && !model.TransfersFiles.IsUnknown() {
		sdk.TransfersFiles = model.TransfersFiles.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "TransfersFiles", "value": *sdk.TransfersFiles})
	}

	// Handling Primitives
	if !model.TunnelsOtherApps.IsNull() && !model.TunnelsOtherApps.IsUnknown() {
		sdk.TunnelsOtherApps = model.TunnelsOtherApps.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "TunnelsOtherApps", "value": *sdk.TunnelsOtherApps})
	}

	// Handling Primitives
	if !model.UsedByMalware.IsNull() && !model.UsedByMalware.IsUnknown() {
		sdk.UsedByMalware = model.UsedByMalware.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "UsedByMalware", "value": *sdk.UsedByMalware})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ApplicationFilters", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ApplicationFilters ---
func packApplicationFiltersFromSdk(ctx context.Context, sdk objects.ApplicationFilters) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ApplicationFilters", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ApplicationFilters
	var d diag.Diagnostics
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
	if sdk.Device != nil {
		model.Device = basetypes.NewStringValue(*sdk.Device)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Device", "value": *sdk.Device})
	} else {
		model.Device = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Evasive != nil {
		model.Evasive = basetypes.NewBoolValue(*sdk.Evasive)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Evasive", "value": *sdk.Evasive})
	} else {
		model.Evasive = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.ExcessiveBandwidthUse != nil {
		model.ExcessiveBandwidthUse = basetypes.NewBoolValue(*sdk.ExcessiveBandwidthUse)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "ExcessiveBandwidthUse", "value": *sdk.ExcessiveBandwidthUse})
	} else {
		model.ExcessiveBandwidthUse = basetypes.NewBoolNull()
	}
	// Handling Lists
	if sdk.Exclude != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Exclude")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Exclude, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Exclude)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Exclude = basetypes.NewListNull(elemType)
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
	if sdk.HasKnownVulnerabilities != nil {
		model.HasKnownVulnerabilities = basetypes.NewBoolValue(*sdk.HasKnownVulnerabilities)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "HasKnownVulnerabilities", "value": *sdk.HasKnownVulnerabilities})
	} else {
		model.HasKnownVulnerabilities = basetypes.NewBoolNull()
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
	if sdk.IsSaas != nil {
		model.IsSaas = basetypes.NewBoolValue(*sdk.IsSaas)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "IsSaas", "value": *sdk.IsSaas})
	} else {
		model.IsSaas = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	// Handling Primitives
	// Standard primitive packing
	if sdk.NewAppid != nil {
		model.NewAppid = basetypes.NewBoolValue(*sdk.NewAppid)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "NewAppid", "value": *sdk.NewAppid})
	} else {
		model.NewAppid = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Pervasive != nil {
		model.Pervasive = basetypes.NewBoolValue(*sdk.Pervasive)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Pervasive", "value": *sdk.Pervasive})
	} else {
		model.Pervasive = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.ProneToMisuse != nil {
		model.ProneToMisuse = basetypes.NewBoolValue(*sdk.ProneToMisuse)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "ProneToMisuse", "value": *sdk.ProneToMisuse})
	} else {
		model.ProneToMisuse = basetypes.NewBoolNull()
	}
	// Handling Lists
	if sdk.Risk != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Risk")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		elemType = basetypes.Int64Type{}
		model.Risk, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Risk)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		elemType = basetypes.Int64Type{}
		model.Risk = basetypes.NewListNull(elemType)
	}
	// Handling Lists
	if sdk.SaasCertifications != nil {
		tflog.Debug(ctx, "Packing list of primitives for field SaasCertifications")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.SaasCertifications, d = basetypes.NewListValueFrom(ctx, elemType, sdk.SaasCertifications)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.SaasCertifications = basetypes.NewListNull(elemType)
	}
	// Handling Lists
	if sdk.SaasRisk != nil {
		tflog.Debug(ctx, "Packing list of primitives for field SaasRisk")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.SaasRisk, d = basetypes.NewListValueFrom(ctx, elemType, sdk.SaasRisk)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.SaasRisk = basetypes.NewListNull(elemType)
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
	if sdk.Subcategory != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Subcategory")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Subcategory, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Subcategory)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Subcategory = basetypes.NewListNull(elemType)
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Tagging != nil {
		tflog.Debug(ctx, "Packing nested object for field Tagging")
		packed, d := packApplicationFiltersTaggingFromSdk(ctx, *sdk.Tagging)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Tagging"})
		}
		model.Tagging = packed
	} else {
		model.Tagging = basetypes.NewObjectNull(models.ApplicationFiltersTagging{}.AttrTypes())
	}
	// Handling Lists
	if sdk.Technology != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Technology")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Technology, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Technology)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Technology = basetypes.NewListNull(elemType)
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.TransfersFiles != nil {
		model.TransfersFiles = basetypes.NewBoolValue(*sdk.TransfersFiles)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "TransfersFiles", "value": *sdk.TransfersFiles})
	} else {
		model.TransfersFiles = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.TunnelsOtherApps != nil {
		model.TunnelsOtherApps = basetypes.NewBoolValue(*sdk.TunnelsOtherApps)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "TunnelsOtherApps", "value": *sdk.TunnelsOtherApps})
	} else {
		model.TunnelsOtherApps = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.UsedByMalware != nil {
		model.UsedByMalware = basetypes.NewBoolValue(*sdk.UsedByMalware)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "UsedByMalware", "value": *sdk.UsedByMalware})
	} else {
		model.UsedByMalware = basetypes.NewBoolNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ApplicationFilters{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ApplicationFilters", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ApplicationFilters ---
func unpackApplicationFiltersListToSdk(ctx context.Context, list types.List) ([]objects.ApplicationFilters, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ApplicationFilters")
	diags := diag.Diagnostics{}
	var data []models.ApplicationFilters
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.ApplicationFilters, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ApplicationFilters{}.AttrTypes(), &item)
		unpacked, d := unpackApplicationFiltersToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ApplicationFilters", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ApplicationFilters ---
func packApplicationFiltersListFromSdk(ctx context.Context, sdks []objects.ApplicationFilters) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ApplicationFilters")
	diags := diag.Diagnostics{}
	var data []models.ApplicationFilters

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ApplicationFilters
		obj, d := packApplicationFiltersFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ApplicationFilters{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ApplicationFilters", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ApplicationFilters{}.AttrType(), data)
}

// --- Unpacker for ApplicationFiltersTagging ---
func unpackApplicationFiltersTaggingToSdk(ctx context.Context, obj types.Object) (*objects.ApplicationFiltersTagging, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ApplicationFiltersTagging", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ApplicationFiltersTagging
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.ApplicationFiltersTagging
	var d diag.Diagnostics
	// Handling Primitives
	if !model.NoTag.IsNull() && !model.NoTag.IsUnknown() {
		sdk.NoTag = model.NoTag.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "NoTag", "value": *sdk.NoTag})
	}

	// Handling Lists
	if !model.Tag.IsNull() && !model.Tag.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Tag")
		diags.Append(model.Tag.ElementsAs(ctx, &sdk.Tag, false)...)
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ApplicationFiltersTagging", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ApplicationFiltersTagging ---
func packApplicationFiltersTaggingFromSdk(ctx context.Context, sdk objects.ApplicationFiltersTagging) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ApplicationFiltersTagging", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ApplicationFiltersTagging
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.NoTag != nil {
		model.NoTag = basetypes.NewBoolValue(*sdk.NoTag)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "NoTag", "value": *sdk.NoTag})
	} else {
		model.NoTag = basetypes.NewBoolNull()
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

	obj, d := types.ObjectValueFrom(ctx, models.ApplicationFiltersTagging{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ApplicationFiltersTagging", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ApplicationFiltersTagging ---
func unpackApplicationFiltersTaggingListToSdk(ctx context.Context, list types.List) ([]objects.ApplicationFiltersTagging, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ApplicationFiltersTagging")
	diags := diag.Diagnostics{}
	var data []models.ApplicationFiltersTagging
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.ApplicationFiltersTagging, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ApplicationFiltersTagging{}.AttrTypes(), &item)
		unpacked, d := unpackApplicationFiltersTaggingToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ApplicationFiltersTagging", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ApplicationFiltersTagging ---
func packApplicationFiltersTaggingListFromSdk(ctx context.Context, sdks []objects.ApplicationFiltersTagging) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ApplicationFiltersTagging")
	diags := diag.Diagnostics{}
	var data []models.ApplicationFiltersTagging

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ApplicationFiltersTagging
		obj, d := packApplicationFiltersTaggingFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ApplicationFiltersTagging{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ApplicationFiltersTagging", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ApplicationFiltersTagging{}.AttrType(), data)
}
