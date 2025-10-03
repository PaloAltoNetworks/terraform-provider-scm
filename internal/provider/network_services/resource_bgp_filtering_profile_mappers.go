package provider

import (
	"context"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/network_services"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/network_services"
)

// --- Unpacker for BgpFilteringProfiles ---
func unpackBgpFilteringProfilesToSdk(ctx context.Context, obj types.Object) (*network_services.BgpFilteringProfiles, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpFilteringProfiles", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpFilteringProfiles
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpFilteringProfiles
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

	// Handling Objects
	if !model.Ipv4.IsNull() && !model.Ipv4.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Ipv4")
		unpacked, d := unpackBgpFilteringProfilesIpv4ToSdk(ctx, model.Ipv4)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Ipv4"})
		}
		if unpacked != nil {
			sdk.Ipv4 = unpacked
		}
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

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpFilteringProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpFilteringProfiles ---
func packBgpFilteringProfilesFromSdk(ctx context.Context, sdk network_services.BgpFilteringProfiles) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpFilteringProfiles", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpFilteringProfiles
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
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Ipv4 != nil {
		tflog.Debug(ctx, "Packing nested object for field Ipv4")
		packed, d := packBgpFilteringProfilesIpv4FromSdk(ctx, *sdk.Ipv4)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Ipv4"})
		}
		model.Ipv4 = packed
	} else {
		model.Ipv4 = basetypes.NewObjectNull(models.BgpFilteringProfilesIpv4{}.AttrTypes())
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
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BgpFilteringProfiles{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpFilteringProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpFilteringProfiles ---
func unpackBgpFilteringProfilesListToSdk(ctx context.Context, list types.List) ([]network_services.BgpFilteringProfiles, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpFilteringProfiles")
	diags := diag.Diagnostics{}
	var data []models.BgpFilteringProfiles
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpFilteringProfiles, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpFilteringProfiles{}.AttrTypes(), &item)
		unpacked, d := unpackBgpFilteringProfilesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpFilteringProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpFilteringProfiles ---
func packBgpFilteringProfilesListFromSdk(ctx context.Context, sdks []network_services.BgpFilteringProfiles) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpFilteringProfiles")
	diags := diag.Diagnostics{}
	var data []models.BgpFilteringProfiles

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpFilteringProfiles
		obj, d := packBgpFilteringProfilesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpFilteringProfiles{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpFilteringProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpFilteringProfiles{}.AttrType(), data)
}

// --- Unpacker for BgpFilteringProfilesIpv4 ---
func unpackBgpFilteringProfilesIpv4ToSdk(ctx context.Context, obj types.Object) (*network_services.BgpFilteringProfilesIpv4, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpFilteringProfilesIpv4", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpFilteringProfilesIpv4
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpFilteringProfilesIpv4
	var d diag.Diagnostics
	// Handling Objects
	if !model.Ipv4.IsNull() && !model.Ipv4.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Ipv4")
		unpacked, d := unpackBgpFilteringProfilesIpv4Ipv4ToSdk(ctx, model.Ipv4)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Ipv4"})
		}
		if unpacked != nil {
			sdk.Ipv4 = *unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpFilteringProfilesIpv4", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpFilteringProfilesIpv4 ---
func packBgpFilteringProfilesIpv4FromSdk(ctx context.Context, sdk network_services.BgpFilteringProfilesIpv4) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpFilteringProfilesIpv4", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpFilteringProfilesIpv4
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	// Logic for non-pointer / value-type nested objects
	if !reflect.ValueOf(sdk.Ipv4).IsZero() {
		tflog.Debug(ctx, "Packing nested object for field Ipv4")
		packed, d := packBgpFilteringProfilesIpv4Ipv4FromSdk(ctx, sdk.Ipv4)
		diags.Append(d...)
		model.Ipv4 = packed
	} else {
		model.Ipv4 = basetypes.NewObjectNull(models.BgpFilteringProfilesIpv4Ipv4{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BgpFilteringProfilesIpv4{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpFilteringProfilesIpv4", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpFilteringProfilesIpv4 ---
func unpackBgpFilteringProfilesIpv4ListToSdk(ctx context.Context, list types.List) ([]network_services.BgpFilteringProfilesIpv4, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpFilteringProfilesIpv4")
	diags := diag.Diagnostics{}
	var data []models.BgpFilteringProfilesIpv4
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpFilteringProfilesIpv4, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpFilteringProfilesIpv4{}.AttrTypes(), &item)
		unpacked, d := unpackBgpFilteringProfilesIpv4ToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpFilteringProfilesIpv4", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpFilteringProfilesIpv4 ---
func packBgpFilteringProfilesIpv4ListFromSdk(ctx context.Context, sdks []network_services.BgpFilteringProfilesIpv4) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpFilteringProfilesIpv4")
	diags := diag.Diagnostics{}
	var data []models.BgpFilteringProfilesIpv4

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpFilteringProfilesIpv4
		obj, d := packBgpFilteringProfilesIpv4FromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpFilteringProfilesIpv4{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpFilteringProfilesIpv4", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpFilteringProfilesIpv4{}.AttrType(), data)
}

// --- Unpacker for BgpFilteringProfilesIpv4Ipv4 ---
func unpackBgpFilteringProfilesIpv4Ipv4ToSdk(ctx context.Context, obj types.Object) (*network_services.BgpFilteringProfilesIpv4Ipv4, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpFilteringProfilesIpv4Ipv4", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpFilteringProfilesIpv4Ipv4
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpFilteringProfilesIpv4Ipv4
	var d diag.Diagnostics
	// Handling Objects
	if !model.Multicast.IsNull() && !model.Multicast.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Multicast")
		unpacked, d := unpackBgpFilteringProfilesIpv4Ipv4MulticastToSdk(ctx, model.Multicast)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Multicast"})
		}
		if unpacked != nil {
			sdk.Multicast = unpacked
		}
	}

	// Handling Objects
	if !model.Unicast.IsNull() && !model.Unicast.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Unicast")
		unpacked, d := unpackBgpFilterToSdk(ctx, model.Unicast)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Unicast"})
		}
		if unpacked != nil {
			sdk.Unicast = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpFilteringProfilesIpv4Ipv4", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpFilteringProfilesIpv4Ipv4 ---
func packBgpFilteringProfilesIpv4Ipv4FromSdk(ctx context.Context, sdk network_services.BgpFilteringProfilesIpv4Ipv4) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpFilteringProfilesIpv4Ipv4", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpFilteringProfilesIpv4Ipv4
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Multicast != nil {
		tflog.Debug(ctx, "Packing nested object for field Multicast")
		packed, d := packBgpFilteringProfilesIpv4Ipv4MulticastFromSdk(ctx, *sdk.Multicast)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Multicast"})
		}
		model.Multicast = packed
	} else {
		model.Multicast = basetypes.NewObjectNull(models.BgpFilteringProfilesIpv4Ipv4Multicast{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Unicast != nil {
		tflog.Debug(ctx, "Packing nested object for field Unicast")
		packed, d := packBgpFilterFromSdk(ctx, *sdk.Unicast)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Unicast"})
		}
		model.Unicast = packed
	} else {
		model.Unicast = basetypes.NewObjectNull(models.BgpFilter{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BgpFilteringProfilesIpv4Ipv4{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpFilteringProfilesIpv4Ipv4", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpFilteringProfilesIpv4Ipv4 ---
func unpackBgpFilteringProfilesIpv4Ipv4ListToSdk(ctx context.Context, list types.List) ([]network_services.BgpFilteringProfilesIpv4Ipv4, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpFilteringProfilesIpv4Ipv4")
	diags := diag.Diagnostics{}
	var data []models.BgpFilteringProfilesIpv4Ipv4
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpFilteringProfilesIpv4Ipv4, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpFilteringProfilesIpv4Ipv4{}.AttrTypes(), &item)
		unpacked, d := unpackBgpFilteringProfilesIpv4Ipv4ToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpFilteringProfilesIpv4Ipv4", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpFilteringProfilesIpv4Ipv4 ---
func packBgpFilteringProfilesIpv4Ipv4ListFromSdk(ctx context.Context, sdks []network_services.BgpFilteringProfilesIpv4Ipv4) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpFilteringProfilesIpv4Ipv4")
	diags := diag.Diagnostics{}
	var data []models.BgpFilteringProfilesIpv4Ipv4

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpFilteringProfilesIpv4Ipv4
		obj, d := packBgpFilteringProfilesIpv4Ipv4FromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpFilteringProfilesIpv4Ipv4{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpFilteringProfilesIpv4Ipv4", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpFilteringProfilesIpv4Ipv4{}.AttrType(), data)
}

// --- Unpacker for BgpFilteringProfilesIpv4Ipv4Multicast ---
func unpackBgpFilteringProfilesIpv4Ipv4MulticastToSdk(ctx context.Context, obj types.Object) (*network_services.BgpFilteringProfilesIpv4Ipv4Multicast, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpFilteringProfilesIpv4Ipv4Multicast", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpFilteringProfilesIpv4Ipv4Multicast
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpFilteringProfilesIpv4Ipv4Multicast
	var d diag.Diagnostics
	// Handling Objects
	if !model.ConditionalAdvertisement.IsNull() && !model.ConditionalAdvertisement.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field ConditionalAdvertisement")
		unpacked, d := unpackBgpFilterConditionalAdvertisementToSdk(ctx, model.ConditionalAdvertisement)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "ConditionalAdvertisement"})
		}
		if unpacked != nil {
			sdk.ConditionalAdvertisement = unpacked
		}
	}

	// Handling Objects
	if !model.FilterList.IsNull() && !model.FilterList.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field FilterList")
		unpacked, d := unpackBgpFilterFilterListToSdk(ctx, model.FilterList)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "FilterList"})
		}
		if unpacked != nil {
			sdk.FilterList = unpacked
		}
	}

	// Handling Objects
	if !model.InboundNetworkFilters.IsNull() && !model.InboundNetworkFilters.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field InboundNetworkFilters")
		unpacked, d := unpackBgpFilterInboundNetworkFiltersToSdk(ctx, model.InboundNetworkFilters)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "InboundNetworkFilters"})
		}
		if unpacked != nil {
			sdk.InboundNetworkFilters = unpacked
		}
	}

	// Handling Primitives
	if !model.Inherit.IsNull() && !model.Inherit.IsUnknown() {
		sdk.Inherit = model.Inherit.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Inherit", "value": *sdk.Inherit})
	}

	// Handling Objects
	if !model.OutboundNetworkFilters.IsNull() && !model.OutboundNetworkFilters.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field OutboundNetworkFilters")
		unpacked, d := unpackBgpFilterInboundNetworkFiltersToSdk(ctx, model.OutboundNetworkFilters)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "OutboundNetworkFilters"})
		}
		if unpacked != nil {
			sdk.OutboundNetworkFilters = unpacked
		}
	}

	// Handling Objects
	if !model.RouteMaps.IsNull() && !model.RouteMaps.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field RouteMaps")
		unpacked, d := unpackBgpFilterFilterListToSdk(ctx, model.RouteMaps)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "RouteMaps"})
		}
		if unpacked != nil {
			sdk.RouteMaps = unpacked
		}
	}

	// Handling Primitives
	if !model.UnsuppressMap.IsNull() && !model.UnsuppressMap.IsUnknown() {
		sdk.UnsuppressMap = model.UnsuppressMap.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "UnsuppressMap", "value": *sdk.UnsuppressMap})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpFilteringProfilesIpv4Ipv4Multicast", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpFilteringProfilesIpv4Ipv4Multicast ---
func packBgpFilteringProfilesIpv4Ipv4MulticastFromSdk(ctx context.Context, sdk network_services.BgpFilteringProfilesIpv4Ipv4Multicast) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpFilteringProfilesIpv4Ipv4Multicast", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpFilteringProfilesIpv4Ipv4Multicast
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.ConditionalAdvertisement != nil {
		tflog.Debug(ctx, "Packing nested object for field ConditionalAdvertisement")
		packed, d := packBgpFilterConditionalAdvertisementFromSdk(ctx, *sdk.ConditionalAdvertisement)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "ConditionalAdvertisement"})
		}
		model.ConditionalAdvertisement = packed
	} else {
		model.ConditionalAdvertisement = basetypes.NewObjectNull(models.BgpFilterConditionalAdvertisement{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.FilterList != nil {
		tflog.Debug(ctx, "Packing nested object for field FilterList")
		packed, d := packBgpFilterFilterListFromSdk(ctx, *sdk.FilterList)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "FilterList"})
		}
		model.FilterList = packed
	} else {
		model.FilterList = basetypes.NewObjectNull(models.BgpFilterFilterList{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.InboundNetworkFilters != nil {
		tflog.Debug(ctx, "Packing nested object for field InboundNetworkFilters")
		packed, d := packBgpFilterInboundNetworkFiltersFromSdk(ctx, *sdk.InboundNetworkFilters)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "InboundNetworkFilters"})
		}
		model.InboundNetworkFilters = packed
	} else {
		model.InboundNetworkFilters = basetypes.NewObjectNull(models.BgpFilterInboundNetworkFilters{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Inherit != nil {
		model.Inherit = basetypes.NewBoolValue(*sdk.Inherit)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Inherit", "value": *sdk.Inherit})
	} else {
		model.Inherit = basetypes.NewBoolNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.OutboundNetworkFilters != nil {
		tflog.Debug(ctx, "Packing nested object for field OutboundNetworkFilters")
		packed, d := packBgpFilterInboundNetworkFiltersFromSdk(ctx, *sdk.OutboundNetworkFilters)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "OutboundNetworkFilters"})
		}
		model.OutboundNetworkFilters = packed
	} else {
		model.OutboundNetworkFilters = basetypes.NewObjectNull(models.BgpFilterInboundNetworkFilters{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.RouteMaps != nil {
		tflog.Debug(ctx, "Packing nested object for field RouteMaps")
		packed, d := packBgpFilterFilterListFromSdk(ctx, *sdk.RouteMaps)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "RouteMaps"})
		}
		model.RouteMaps = packed
	} else {
		model.RouteMaps = basetypes.NewObjectNull(models.BgpFilterFilterList{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.UnsuppressMap != nil {
		model.UnsuppressMap = basetypes.NewStringValue(*sdk.UnsuppressMap)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "UnsuppressMap", "value": *sdk.UnsuppressMap})
	} else {
		model.UnsuppressMap = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BgpFilteringProfilesIpv4Ipv4Multicast{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpFilteringProfilesIpv4Ipv4Multicast", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpFilteringProfilesIpv4Ipv4Multicast ---
func unpackBgpFilteringProfilesIpv4Ipv4MulticastListToSdk(ctx context.Context, list types.List) ([]network_services.BgpFilteringProfilesIpv4Ipv4Multicast, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpFilteringProfilesIpv4Ipv4Multicast")
	diags := diag.Diagnostics{}
	var data []models.BgpFilteringProfilesIpv4Ipv4Multicast
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpFilteringProfilesIpv4Ipv4Multicast, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpFilteringProfilesIpv4Ipv4Multicast{}.AttrTypes(), &item)
		unpacked, d := unpackBgpFilteringProfilesIpv4Ipv4MulticastToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpFilteringProfilesIpv4Ipv4Multicast", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpFilteringProfilesIpv4Ipv4Multicast ---
func packBgpFilteringProfilesIpv4Ipv4MulticastListFromSdk(ctx context.Context, sdks []network_services.BgpFilteringProfilesIpv4Ipv4Multicast) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpFilteringProfilesIpv4Ipv4Multicast")
	diags := diag.Diagnostics{}
	var data []models.BgpFilteringProfilesIpv4Ipv4Multicast

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpFilteringProfilesIpv4Ipv4Multicast
		obj, d := packBgpFilteringProfilesIpv4Ipv4MulticastFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpFilteringProfilesIpv4Ipv4Multicast{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpFilteringProfilesIpv4Ipv4Multicast", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpFilteringProfilesIpv4Ipv4Multicast{}.AttrType(), data)
}

// --- Unpacker for BgpFilterConditionalAdvertisement ---
func unpackBgpFilterConditionalAdvertisementToSdk(ctx context.Context, obj types.Object) (*network_services.BgpFilterConditionalAdvertisement, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpFilterConditionalAdvertisement", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpFilterConditionalAdvertisement
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpFilterConditionalAdvertisement
	var d diag.Diagnostics
	// Handling Objects
	if !model.Exist.IsNull() && !model.Exist.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Exist")
		unpacked, d := unpackBgpFilterConditionalAdvertisementExistToSdk(ctx, model.Exist)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Exist"})
		}
		if unpacked != nil {
			sdk.Exist = unpacked
		}
	}

	// Handling Objects
	if !model.NonExist.IsNull() && !model.NonExist.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field NonExist")
		unpacked, d := unpackBgpFilterConditionalAdvertisementNonExistToSdk(ctx, model.NonExist)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "NonExist"})
		}
		if unpacked != nil {
			sdk.NonExist = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpFilterConditionalAdvertisement", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpFilterConditionalAdvertisement ---
func packBgpFilterConditionalAdvertisementFromSdk(ctx context.Context, sdk network_services.BgpFilterConditionalAdvertisement) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpFilterConditionalAdvertisement", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpFilterConditionalAdvertisement
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Exist != nil {
		tflog.Debug(ctx, "Packing nested object for field Exist")
		packed, d := packBgpFilterConditionalAdvertisementExistFromSdk(ctx, *sdk.Exist)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Exist"})
		}
		model.Exist = packed
	} else {
		model.Exist = basetypes.NewObjectNull(models.BgpFilterConditionalAdvertisementExist{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.NonExist != nil {
		tflog.Debug(ctx, "Packing nested object for field NonExist")
		packed, d := packBgpFilterConditionalAdvertisementNonExistFromSdk(ctx, *sdk.NonExist)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "NonExist"})
		}
		model.NonExist = packed
	} else {
		model.NonExist = basetypes.NewObjectNull(models.BgpFilterConditionalAdvertisementNonExist{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BgpFilterConditionalAdvertisement{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpFilterConditionalAdvertisement", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpFilterConditionalAdvertisement ---
func unpackBgpFilterConditionalAdvertisementListToSdk(ctx context.Context, list types.List) ([]network_services.BgpFilterConditionalAdvertisement, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpFilterConditionalAdvertisement")
	diags := diag.Diagnostics{}
	var data []models.BgpFilterConditionalAdvertisement
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpFilterConditionalAdvertisement, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpFilterConditionalAdvertisement{}.AttrTypes(), &item)
		unpacked, d := unpackBgpFilterConditionalAdvertisementToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpFilterConditionalAdvertisement", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpFilterConditionalAdvertisement ---
func packBgpFilterConditionalAdvertisementListFromSdk(ctx context.Context, sdks []network_services.BgpFilterConditionalAdvertisement) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpFilterConditionalAdvertisement")
	diags := diag.Diagnostics{}
	var data []models.BgpFilterConditionalAdvertisement

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpFilterConditionalAdvertisement
		obj, d := packBgpFilterConditionalAdvertisementFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpFilterConditionalAdvertisement{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpFilterConditionalAdvertisement", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpFilterConditionalAdvertisement{}.AttrType(), data)
}

// --- Unpacker for BgpFilterConditionalAdvertisementExist ---
func unpackBgpFilterConditionalAdvertisementExistToSdk(ctx context.Context, obj types.Object) (*network_services.BgpFilterConditionalAdvertisementExist, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpFilterConditionalAdvertisementExist", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpFilterConditionalAdvertisementExist
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpFilterConditionalAdvertisementExist
	var d diag.Diagnostics
	// Handling Primitives
	if !model.AdvertiseMap.IsNull() && !model.AdvertiseMap.IsUnknown() {
		sdk.AdvertiseMap = model.AdvertiseMap.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "AdvertiseMap", "value": *sdk.AdvertiseMap})
	}

	// Handling Primitives
	if !model.ExistMap.IsNull() && !model.ExistMap.IsUnknown() {
		sdk.ExistMap = model.ExistMap.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "ExistMap", "value": *sdk.ExistMap})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpFilterConditionalAdvertisementExist", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpFilterConditionalAdvertisementExist ---
func packBgpFilterConditionalAdvertisementExistFromSdk(ctx context.Context, sdk network_services.BgpFilterConditionalAdvertisementExist) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpFilterConditionalAdvertisementExist", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpFilterConditionalAdvertisementExist
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.AdvertiseMap != nil {
		model.AdvertiseMap = basetypes.NewStringValue(*sdk.AdvertiseMap)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "AdvertiseMap", "value": *sdk.AdvertiseMap})
	} else {
		model.AdvertiseMap = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.ExistMap != nil {
		model.ExistMap = basetypes.NewStringValue(*sdk.ExistMap)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "ExistMap", "value": *sdk.ExistMap})
	} else {
		model.ExistMap = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BgpFilterConditionalAdvertisementExist{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpFilterConditionalAdvertisementExist", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpFilterConditionalAdvertisementExist ---
func unpackBgpFilterConditionalAdvertisementExistListToSdk(ctx context.Context, list types.List) ([]network_services.BgpFilterConditionalAdvertisementExist, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpFilterConditionalAdvertisementExist")
	diags := diag.Diagnostics{}
	var data []models.BgpFilterConditionalAdvertisementExist
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpFilterConditionalAdvertisementExist, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpFilterConditionalAdvertisementExist{}.AttrTypes(), &item)
		unpacked, d := unpackBgpFilterConditionalAdvertisementExistToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpFilterConditionalAdvertisementExist", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpFilterConditionalAdvertisementExist ---
func packBgpFilterConditionalAdvertisementExistListFromSdk(ctx context.Context, sdks []network_services.BgpFilterConditionalAdvertisementExist) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpFilterConditionalAdvertisementExist")
	diags := diag.Diagnostics{}
	var data []models.BgpFilterConditionalAdvertisementExist

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpFilterConditionalAdvertisementExist
		obj, d := packBgpFilterConditionalAdvertisementExistFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpFilterConditionalAdvertisementExist{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpFilterConditionalAdvertisementExist", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpFilterConditionalAdvertisementExist{}.AttrType(), data)
}

// --- Unpacker for BgpFilterConditionalAdvertisementNonExist ---
func unpackBgpFilterConditionalAdvertisementNonExistToSdk(ctx context.Context, obj types.Object) (*network_services.BgpFilterConditionalAdvertisementNonExist, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpFilterConditionalAdvertisementNonExist", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpFilterConditionalAdvertisementNonExist
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpFilterConditionalAdvertisementNonExist
	var d diag.Diagnostics
	// Handling Primitives
	if !model.AdvertiseMap.IsNull() && !model.AdvertiseMap.IsUnknown() {
		sdk.AdvertiseMap = model.AdvertiseMap.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "AdvertiseMap", "value": *sdk.AdvertiseMap})
	}

	// Handling Primitives
	if !model.NonExistMap.IsNull() && !model.NonExistMap.IsUnknown() {
		sdk.NonExistMap = model.NonExistMap.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "NonExistMap", "value": *sdk.NonExistMap})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpFilterConditionalAdvertisementNonExist", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpFilterConditionalAdvertisementNonExist ---
func packBgpFilterConditionalAdvertisementNonExistFromSdk(ctx context.Context, sdk network_services.BgpFilterConditionalAdvertisementNonExist) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpFilterConditionalAdvertisementNonExist", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpFilterConditionalAdvertisementNonExist
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.AdvertiseMap != nil {
		model.AdvertiseMap = basetypes.NewStringValue(*sdk.AdvertiseMap)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "AdvertiseMap", "value": *sdk.AdvertiseMap})
	} else {
		model.AdvertiseMap = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.NonExistMap != nil {
		model.NonExistMap = basetypes.NewStringValue(*sdk.NonExistMap)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "NonExistMap", "value": *sdk.NonExistMap})
	} else {
		model.NonExistMap = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BgpFilterConditionalAdvertisementNonExist{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpFilterConditionalAdvertisementNonExist", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpFilterConditionalAdvertisementNonExist ---
func unpackBgpFilterConditionalAdvertisementNonExistListToSdk(ctx context.Context, list types.List) ([]network_services.BgpFilterConditionalAdvertisementNonExist, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpFilterConditionalAdvertisementNonExist")
	diags := diag.Diagnostics{}
	var data []models.BgpFilterConditionalAdvertisementNonExist
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpFilterConditionalAdvertisementNonExist, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpFilterConditionalAdvertisementNonExist{}.AttrTypes(), &item)
		unpacked, d := unpackBgpFilterConditionalAdvertisementNonExistToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpFilterConditionalAdvertisementNonExist", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpFilterConditionalAdvertisementNonExist ---
func packBgpFilterConditionalAdvertisementNonExistListFromSdk(ctx context.Context, sdks []network_services.BgpFilterConditionalAdvertisementNonExist) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpFilterConditionalAdvertisementNonExist")
	diags := diag.Diagnostics{}
	var data []models.BgpFilterConditionalAdvertisementNonExist

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpFilterConditionalAdvertisementNonExist
		obj, d := packBgpFilterConditionalAdvertisementNonExistFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpFilterConditionalAdvertisementNonExist{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpFilterConditionalAdvertisementNonExist", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpFilterConditionalAdvertisementNonExist{}.AttrType(), data)
}

// --- Unpacker for BgpFilterFilterList ---
func unpackBgpFilterFilterListToSdk(ctx context.Context, obj types.Object) (*network_services.BgpFilterFilterList, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpFilterFilterList", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpFilterFilterList
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpFilterFilterList
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Inbound.IsNull() && !model.Inbound.IsUnknown() {
		sdk.Inbound = model.Inbound.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Inbound", "value": *sdk.Inbound})
	}

	// Handling Primitives
	if !model.Outbound.IsNull() && !model.Outbound.IsUnknown() {
		sdk.Outbound = model.Outbound.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Outbound", "value": *sdk.Outbound})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpFilterFilterList", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpFilterFilterList ---
func packBgpFilterFilterListFromSdk(ctx context.Context, sdk network_services.BgpFilterFilterList) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpFilterFilterList", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpFilterFilterList
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Inbound != nil {
		model.Inbound = basetypes.NewStringValue(*sdk.Inbound)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Inbound", "value": *sdk.Inbound})
	} else {
		model.Inbound = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Outbound != nil {
		model.Outbound = basetypes.NewStringValue(*sdk.Outbound)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Outbound", "value": *sdk.Outbound})
	} else {
		model.Outbound = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BgpFilterFilterList{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpFilterFilterList", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpFilterFilterList ---
func unpackBgpFilterFilterListListToSdk(ctx context.Context, list types.List) ([]network_services.BgpFilterFilterList, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpFilterFilterList")
	diags := diag.Diagnostics{}
	var data []models.BgpFilterFilterList
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpFilterFilterList, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpFilterFilterList{}.AttrTypes(), &item)
		unpacked, d := unpackBgpFilterFilterListToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpFilterFilterList", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpFilterFilterList ---
func packBgpFilterFilterListListFromSdk(ctx context.Context, sdks []network_services.BgpFilterFilterList) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpFilterFilterList")
	diags := diag.Diagnostics{}
	var data []models.BgpFilterFilterList

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpFilterFilterList
		obj, d := packBgpFilterFilterListFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpFilterFilterList{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpFilterFilterList", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpFilterFilterList{}.AttrType(), data)
}

// --- Unpacker for BgpFilterInboundNetworkFilters ---
func unpackBgpFilterInboundNetworkFiltersToSdk(ctx context.Context, obj types.Object) (*network_services.BgpFilterInboundNetworkFilters, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpFilterInboundNetworkFilters", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpFilterInboundNetworkFilters
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpFilterInboundNetworkFilters
	var d diag.Diagnostics
	// Handling Primitives
	if !model.DistributeList.IsNull() && !model.DistributeList.IsUnknown() {
		sdk.DistributeList = model.DistributeList.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DistributeList", "value": *sdk.DistributeList})
	}

	// Handling Primitives
	if !model.PrefixList.IsNull() && !model.PrefixList.IsUnknown() {
		sdk.PrefixList = model.PrefixList.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "PrefixList", "value": *sdk.PrefixList})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpFilterInboundNetworkFilters", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpFilterInboundNetworkFilters ---
func packBgpFilterInboundNetworkFiltersFromSdk(ctx context.Context, sdk network_services.BgpFilterInboundNetworkFilters) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpFilterInboundNetworkFilters", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpFilterInboundNetworkFilters
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.DistributeList != nil {
		model.DistributeList = basetypes.NewStringValue(*sdk.DistributeList)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DistributeList", "value": *sdk.DistributeList})
	} else {
		model.DistributeList = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.PrefixList != nil {
		model.PrefixList = basetypes.NewStringValue(*sdk.PrefixList)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "PrefixList", "value": *sdk.PrefixList})
	} else {
		model.PrefixList = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BgpFilterInboundNetworkFilters{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpFilterInboundNetworkFilters", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpFilterInboundNetworkFilters ---
func unpackBgpFilterInboundNetworkFiltersListToSdk(ctx context.Context, list types.List) ([]network_services.BgpFilterInboundNetworkFilters, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpFilterInboundNetworkFilters")
	diags := diag.Diagnostics{}
	var data []models.BgpFilterInboundNetworkFilters
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpFilterInboundNetworkFilters, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpFilterInboundNetworkFilters{}.AttrTypes(), &item)
		unpacked, d := unpackBgpFilterInboundNetworkFiltersToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpFilterInboundNetworkFilters", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpFilterInboundNetworkFilters ---
func packBgpFilterInboundNetworkFiltersListFromSdk(ctx context.Context, sdks []network_services.BgpFilterInboundNetworkFilters) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpFilterInboundNetworkFilters")
	diags := diag.Diagnostics{}
	var data []models.BgpFilterInboundNetworkFilters

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpFilterInboundNetworkFilters
		obj, d := packBgpFilterInboundNetworkFiltersFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpFilterInboundNetworkFilters{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpFilterInboundNetworkFilters", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpFilterInboundNetworkFilters{}.AttrType(), data)
}

// --- Unpacker for BgpFilter ---
func unpackBgpFilterToSdk(ctx context.Context, obj types.Object) (*network_services.BgpFilter, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BgpFilter", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BgpFilter
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.BgpFilter
	var d diag.Diagnostics
	// Handling Objects
	if !model.ConditionalAdvertisement.IsNull() && !model.ConditionalAdvertisement.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field ConditionalAdvertisement")
		unpacked, d := unpackBgpFilterConditionalAdvertisementToSdk(ctx, model.ConditionalAdvertisement)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "ConditionalAdvertisement"})
		}
		if unpacked != nil {
			sdk.ConditionalAdvertisement = unpacked
		}
	}

	// Handling Objects
	if !model.FilterList.IsNull() && !model.FilterList.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field FilterList")
		unpacked, d := unpackBgpFilterFilterListToSdk(ctx, model.FilterList)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "FilterList"})
		}
		if unpacked != nil {
			sdk.FilterList = unpacked
		}
	}

	// Handling Objects
	if !model.InboundNetworkFilters.IsNull() && !model.InboundNetworkFilters.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field InboundNetworkFilters")
		unpacked, d := unpackBgpFilterInboundNetworkFiltersToSdk(ctx, model.InboundNetworkFilters)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "InboundNetworkFilters"})
		}
		if unpacked != nil {
			sdk.InboundNetworkFilters = unpacked
		}
	}

	// Handling Objects
	if !model.OutboundNetworkFilters.IsNull() && !model.OutboundNetworkFilters.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field OutboundNetworkFilters")
		unpacked, d := unpackBgpFilterInboundNetworkFiltersToSdk(ctx, model.OutboundNetworkFilters)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "OutboundNetworkFilters"})
		}
		if unpacked != nil {
			sdk.OutboundNetworkFilters = unpacked
		}
	}

	// Handling Objects
	if !model.RouteMaps.IsNull() && !model.RouteMaps.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field RouteMaps")
		unpacked, d := unpackBgpFilterFilterListToSdk(ctx, model.RouteMaps)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "RouteMaps"})
		}
		if unpacked != nil {
			sdk.RouteMaps = unpacked
		}
	}

	// Handling Primitives
	if !model.UnsuppressMap.IsNull() && !model.UnsuppressMap.IsUnknown() {
		sdk.UnsuppressMap = model.UnsuppressMap.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "UnsuppressMap", "value": *sdk.UnsuppressMap})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BgpFilter", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BgpFilter ---
func packBgpFilterFromSdk(ctx context.Context, sdk network_services.BgpFilter) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BgpFilter", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BgpFilter
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.ConditionalAdvertisement != nil {
		tflog.Debug(ctx, "Packing nested object for field ConditionalAdvertisement")
		packed, d := packBgpFilterConditionalAdvertisementFromSdk(ctx, *sdk.ConditionalAdvertisement)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "ConditionalAdvertisement"})
		}
		model.ConditionalAdvertisement = packed
	} else {
		model.ConditionalAdvertisement = basetypes.NewObjectNull(models.BgpFilterConditionalAdvertisement{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.FilterList != nil {
		tflog.Debug(ctx, "Packing nested object for field FilterList")
		packed, d := packBgpFilterFilterListFromSdk(ctx, *sdk.FilterList)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "FilterList"})
		}
		model.FilterList = packed
	} else {
		model.FilterList = basetypes.NewObjectNull(models.BgpFilterFilterList{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.InboundNetworkFilters != nil {
		tflog.Debug(ctx, "Packing nested object for field InboundNetworkFilters")
		packed, d := packBgpFilterInboundNetworkFiltersFromSdk(ctx, *sdk.InboundNetworkFilters)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "InboundNetworkFilters"})
		}
		model.InboundNetworkFilters = packed
	} else {
		model.InboundNetworkFilters = basetypes.NewObjectNull(models.BgpFilterInboundNetworkFilters{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.OutboundNetworkFilters != nil {
		tflog.Debug(ctx, "Packing nested object for field OutboundNetworkFilters")
		packed, d := packBgpFilterInboundNetworkFiltersFromSdk(ctx, *sdk.OutboundNetworkFilters)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "OutboundNetworkFilters"})
		}
		model.OutboundNetworkFilters = packed
	} else {
		model.OutboundNetworkFilters = basetypes.NewObjectNull(models.BgpFilterInboundNetworkFilters{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.RouteMaps != nil {
		tflog.Debug(ctx, "Packing nested object for field RouteMaps")
		packed, d := packBgpFilterFilterListFromSdk(ctx, *sdk.RouteMaps)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "RouteMaps"})
		}
		model.RouteMaps = packed
	} else {
		model.RouteMaps = basetypes.NewObjectNull(models.BgpFilterFilterList{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.UnsuppressMap != nil {
		model.UnsuppressMap = basetypes.NewStringValue(*sdk.UnsuppressMap)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "UnsuppressMap", "value": *sdk.UnsuppressMap})
	} else {
		model.UnsuppressMap = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BgpFilter{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BgpFilter", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BgpFilter ---
func unpackBgpFilterListToSdk(ctx context.Context, list types.List) ([]network_services.BgpFilter, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BgpFilter")
	diags := diag.Diagnostics{}
	var data []models.BgpFilter
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.BgpFilter, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BgpFilter{}.AttrTypes(), &item)
		unpacked, d := unpackBgpFilterToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BgpFilter", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BgpFilter ---
func packBgpFilterListFromSdk(ctx context.Context, sdks []network_services.BgpFilter) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BgpFilter")
	diags := diag.Diagnostics{}
	var data []models.BgpFilter

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BgpFilter
		obj, d := packBgpFilterFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BgpFilter{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BgpFilter", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BgpFilter{}.AttrType(), data)
}
