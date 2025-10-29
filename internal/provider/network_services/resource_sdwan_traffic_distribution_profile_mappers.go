package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/network_services"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/network_services"
)

// --- Unpacker for SdwanTrafficDistributionProfiles ---
func unpackSdwanTrafficDistributionProfilesToSdk(ctx context.Context, obj types.Object) (*network_services.SdwanTrafficDistributionProfiles, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.SdwanTrafficDistributionProfiles", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.SdwanTrafficDistributionProfiles
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.SdwanTrafficDistributionProfiles
	var d diag.Diagnostics

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
	if !model.LinkTags.IsNull() && !model.LinkTags.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field LinkTags")
		unpacked, d := unpackSdwanTrafficDistributionProfilesLinkTagsInnerListToSdk(ctx, model.LinkTags)
		diags.Append(d...)
		sdk.LinkTags = unpacked
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	// Handling Primitives
	if !model.Snippet.IsNull() && !model.Snippet.IsUnknown() {
		sdk.Snippet = model.Snippet.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	}

	// Handling Primitives
	if !model.TrafficDistribution.IsNull() && !model.TrafficDistribution.IsUnknown() {
		sdk.TrafficDistribution = model.TrafficDistribution.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "TrafficDistribution", "value": *sdk.TrafficDistribution})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.SdwanTrafficDistributionProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for SdwanTrafficDistributionProfiles ---
func packSdwanTrafficDistributionProfilesFromSdk(ctx context.Context, sdk network_services.SdwanTrafficDistributionProfiles) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.SdwanTrafficDistributionProfiles", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.SdwanTrafficDistributionProfiles
	var d diag.Diagnostics
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
	if sdk.LinkTags != nil {
		tflog.Debug(ctx, "Packing list of objects for field LinkTags")
		packed, d := packSdwanTrafficDistributionProfilesLinkTagsInnerListFromSdk(ctx, sdk.LinkTags)
		diags.Append(d...)
		model.LinkTags = packed
	} else {
		model.LinkTags = basetypes.NewListNull(models.SdwanTrafficDistributionProfilesLinkTagsInner{}.AttrType())
	}
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	// Handling Primitives
	// Standard primitive packing
	if sdk.Snippet != nil {
		model.Snippet = basetypes.NewStringValue(*sdk.Snippet)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	} else {
		model.Snippet = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.TrafficDistribution != nil {
		model.TrafficDistribution = basetypes.NewStringValue(*sdk.TrafficDistribution)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "TrafficDistribution", "value": *sdk.TrafficDistribution})
	} else {
		model.TrafficDistribution = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.SdwanTrafficDistributionProfiles{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.SdwanTrafficDistributionProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for SdwanTrafficDistributionProfiles ---
func unpackSdwanTrafficDistributionProfilesListToSdk(ctx context.Context, list types.List) ([]network_services.SdwanTrafficDistributionProfiles, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.SdwanTrafficDistributionProfiles")
	diags := diag.Diagnostics{}
	var data []models.SdwanTrafficDistributionProfiles
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.SdwanTrafficDistributionProfiles, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.SdwanTrafficDistributionProfiles{}.AttrTypes(), &item)
		unpacked, d := unpackSdwanTrafficDistributionProfilesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.SdwanTrafficDistributionProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for SdwanTrafficDistributionProfiles ---
func packSdwanTrafficDistributionProfilesListFromSdk(ctx context.Context, sdks []network_services.SdwanTrafficDistributionProfiles) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.SdwanTrafficDistributionProfiles")
	diags := diag.Diagnostics{}
	var data []models.SdwanTrafficDistributionProfiles

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.SdwanTrafficDistributionProfiles
		obj, d := packSdwanTrafficDistributionProfilesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.SdwanTrafficDistributionProfiles{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.SdwanTrafficDistributionProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.SdwanTrafficDistributionProfiles{}.AttrType(), data)
}

// --- Unpacker for SdwanTrafficDistributionProfilesLinkTagsInner ---
func unpackSdwanTrafficDistributionProfilesLinkTagsInnerToSdk(ctx context.Context, obj types.Object) (*network_services.SdwanTrafficDistributionProfilesLinkTagsInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.SdwanTrafficDistributionProfilesLinkTagsInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.SdwanTrafficDistributionProfilesLinkTagsInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.SdwanTrafficDistributionProfilesLinkTagsInner
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	// Handling Primitives
	if !model.Weight.IsNull() && !model.Weight.IsUnknown() {
		val := int32(model.Weight.ValueInt64())
		sdk.Weight = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Weight", "value": *sdk.Weight})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.SdwanTrafficDistributionProfilesLinkTagsInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for SdwanTrafficDistributionProfilesLinkTagsInner ---
func packSdwanTrafficDistributionProfilesLinkTagsInnerFromSdk(ctx context.Context, sdk network_services.SdwanTrafficDistributionProfilesLinkTagsInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.SdwanTrafficDistributionProfilesLinkTagsInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.SdwanTrafficDistributionProfilesLinkTagsInner
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	// Handling Primitives
	// Standard primitive packing
	if sdk.Weight != nil {
		model.Weight = basetypes.NewInt64Value(int64(*sdk.Weight))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Weight", "value": *sdk.Weight})
	} else {
		model.Weight = basetypes.NewInt64Null()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.SdwanTrafficDistributionProfilesLinkTagsInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.SdwanTrafficDistributionProfilesLinkTagsInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for SdwanTrafficDistributionProfilesLinkTagsInner ---
func unpackSdwanTrafficDistributionProfilesLinkTagsInnerListToSdk(ctx context.Context, list types.List) ([]network_services.SdwanTrafficDistributionProfilesLinkTagsInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.SdwanTrafficDistributionProfilesLinkTagsInner")
	diags := diag.Diagnostics{}
	var data []models.SdwanTrafficDistributionProfilesLinkTagsInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.SdwanTrafficDistributionProfilesLinkTagsInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.SdwanTrafficDistributionProfilesLinkTagsInner{}.AttrTypes(), &item)
		unpacked, d := unpackSdwanTrafficDistributionProfilesLinkTagsInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.SdwanTrafficDistributionProfilesLinkTagsInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for SdwanTrafficDistributionProfilesLinkTagsInner ---
func packSdwanTrafficDistributionProfilesLinkTagsInnerListFromSdk(ctx context.Context, sdks []network_services.SdwanTrafficDistributionProfilesLinkTagsInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.SdwanTrafficDistributionProfilesLinkTagsInner")
	diags := diag.Diagnostics{}
	var data []models.SdwanTrafficDistributionProfilesLinkTagsInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.SdwanTrafficDistributionProfilesLinkTagsInner
		obj, d := packSdwanTrafficDistributionProfilesLinkTagsInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.SdwanTrafficDistributionProfilesLinkTagsInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.SdwanTrafficDistributionProfilesLinkTagsInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.SdwanTrafficDistributionProfilesLinkTagsInner{}.AttrType(), data)
}
