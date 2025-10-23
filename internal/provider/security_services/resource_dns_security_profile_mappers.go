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

// --- Unpacker for DnsSecurityProfiles ---
func unpackDnsSecurityProfilesToSdk(ctx context.Context, obj types.Object) (*security_services.DnsSecurityProfiles, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.DnsSecurityProfiles", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.DnsSecurityProfiles
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.DnsSecurityProfiles
	var d diag.Diagnostics

	// Handling Objects
	if !model.BotnetDomains.IsNull() && !model.BotnetDomains.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field BotnetDomains")
		unpacked, d := unpackDnsSecurityProfilesBotnetDomainsToSdk(ctx, model.BotnetDomains)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "BotnetDomains"})
		}
		if unpacked != nil {
			sdk.BotnetDomains = unpacked
		}
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
		sdk.Id = model.Id.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Id", "value": *sdk.Id})
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
	}

	// Handling Primitives
	if !model.Snippet.IsNull() && !model.Snippet.IsUnknown() {
		sdk.Snippet = model.Snippet.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.DnsSecurityProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for DnsSecurityProfiles ---
func packDnsSecurityProfilesFromSdk(ctx context.Context, sdk security_services.DnsSecurityProfiles) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.DnsSecurityProfiles", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.DnsSecurityProfiles
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.BotnetDomains != nil {
		tflog.Debug(ctx, "Packing nested object for field BotnetDomains")
		packed, d := packDnsSecurityProfilesBotnetDomainsFromSdk(ctx, *sdk.BotnetDomains)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "BotnetDomains"})
		}
		model.BotnetDomains = packed
	} else {
		model.BotnetDomains = basetypes.NewObjectNull(models.DnsSecurityProfilesBotnetDomains{}.AttrTypes())
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
	if sdk.Id != nil {
		model.Id = basetypes.NewStringValue(*sdk.Id)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Id", "value": *sdk.Id})
	} else {
		model.Id = basetypes.NewStringNull()
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
	if sdk.Snippet != nil {
		model.Snippet = basetypes.NewStringValue(*sdk.Snippet)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	} else {
		model.Snippet = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.DnsSecurityProfiles{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.DnsSecurityProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for DnsSecurityProfiles ---
func unpackDnsSecurityProfilesListToSdk(ctx context.Context, list types.List) ([]security_services.DnsSecurityProfiles, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.DnsSecurityProfiles")
	diags := diag.Diagnostics{}
	var data []models.DnsSecurityProfiles
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.DnsSecurityProfiles, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.DnsSecurityProfiles{}.AttrTypes(), &item)
		unpacked, d := unpackDnsSecurityProfilesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.DnsSecurityProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for DnsSecurityProfiles ---
func packDnsSecurityProfilesListFromSdk(ctx context.Context, sdks []security_services.DnsSecurityProfiles) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.DnsSecurityProfiles")
	diags := diag.Diagnostics{}
	var data []models.DnsSecurityProfiles

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.DnsSecurityProfiles
		obj, d := packDnsSecurityProfilesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.DnsSecurityProfiles{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.DnsSecurityProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.DnsSecurityProfiles{}.AttrType(), data)
}

// --- Unpacker for DnsSecurityProfilesBotnetDomains ---
func unpackDnsSecurityProfilesBotnetDomainsToSdk(ctx context.Context, obj types.Object) (*security_services.DnsSecurityProfilesBotnetDomains, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.DnsSecurityProfilesBotnetDomains", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.DnsSecurityProfilesBotnetDomains
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.DnsSecurityProfilesBotnetDomains
	var d diag.Diagnostics
	// Handling Lists
	if !model.DnsSecurityCategories.IsNull() && !model.DnsSecurityCategories.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field DnsSecurityCategories")
		unpacked, d := unpackDnsSecurityProfilesBotnetDomainsDnsSecurityCategoriesInnerListToSdk(ctx, model.DnsSecurityCategories)
		diags.Append(d...)
		sdk.DnsSecurityCategories = unpacked
	}

	// Handling Lists
	if !model.Lists.IsNull() && !model.Lists.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field Lists")
		unpacked, d := unpackDnsSecurityProfilesBotnetDomainsListsInnerListToSdk(ctx, model.Lists)
		diags.Append(d...)
		sdk.Lists = unpacked
	}

	// Handling Objects
	if !model.Sinkhole.IsNull() && !model.Sinkhole.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Sinkhole")
		unpacked, d := unpackDnsSecurityProfilesBotnetDomainsSinkholeToSdk(ctx, model.Sinkhole)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Sinkhole"})
		}
		if unpacked != nil {
			sdk.Sinkhole = unpacked
		}
	}

	// Handling Lists
	if !model.Whitelist.IsNull() && !model.Whitelist.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field Whitelist")
		unpacked, d := unpackDnsSecurityProfilesBotnetDomainsWhitelistInnerListToSdk(ctx, model.Whitelist)
		diags.Append(d...)
		sdk.Whitelist = unpacked
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.DnsSecurityProfilesBotnetDomains", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for DnsSecurityProfilesBotnetDomains ---
func packDnsSecurityProfilesBotnetDomainsFromSdk(ctx context.Context, sdk security_services.DnsSecurityProfilesBotnetDomains) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.DnsSecurityProfilesBotnetDomains", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.DnsSecurityProfilesBotnetDomains
	var d diag.Diagnostics
	// Handling Lists
	if sdk.DnsSecurityCategories != nil {
		tflog.Debug(ctx, "Packing list of objects for field DnsSecurityCategories")
		packed, d := packDnsSecurityProfilesBotnetDomainsDnsSecurityCategoriesInnerListFromSdk(ctx, sdk.DnsSecurityCategories)
		diags.Append(d...)
		model.DnsSecurityCategories = packed
	} else {
		model.DnsSecurityCategories = basetypes.NewListNull(models.DnsSecurityProfilesBotnetDomainsDnsSecurityCategoriesInner{}.AttrType())
	}
	// Handling Lists
	if sdk.Lists != nil {
		tflog.Debug(ctx, "Packing list of objects for field Lists")
		packed, d := packDnsSecurityProfilesBotnetDomainsListsInnerListFromSdk(ctx, sdk.Lists)
		diags.Append(d...)
		model.Lists = packed
	} else {
		model.Lists = basetypes.NewListNull(models.DnsSecurityProfilesBotnetDomainsListsInner{}.AttrType())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Sinkhole != nil {
		tflog.Debug(ctx, "Packing nested object for field Sinkhole")
		packed, d := packDnsSecurityProfilesBotnetDomainsSinkholeFromSdk(ctx, *sdk.Sinkhole)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Sinkhole"})
		}
		model.Sinkhole = packed
	} else {
		model.Sinkhole = basetypes.NewObjectNull(models.DnsSecurityProfilesBotnetDomainsSinkhole{}.AttrTypes())
	}
	// Handling Lists
	if sdk.Whitelist != nil {
		tflog.Debug(ctx, "Packing list of objects for field Whitelist")
		packed, d := packDnsSecurityProfilesBotnetDomainsWhitelistInnerListFromSdk(ctx, sdk.Whitelist)
		diags.Append(d...)
		model.Whitelist = packed
	} else {
		model.Whitelist = basetypes.NewListNull(models.DnsSecurityProfilesBotnetDomainsWhitelistInner{}.AttrType())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.DnsSecurityProfilesBotnetDomains{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.DnsSecurityProfilesBotnetDomains", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for DnsSecurityProfilesBotnetDomains ---
func unpackDnsSecurityProfilesBotnetDomainsListToSdk(ctx context.Context, list types.List) ([]security_services.DnsSecurityProfilesBotnetDomains, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.DnsSecurityProfilesBotnetDomains")
	diags := diag.Diagnostics{}
	var data []models.DnsSecurityProfilesBotnetDomains
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.DnsSecurityProfilesBotnetDomains, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.DnsSecurityProfilesBotnetDomains{}.AttrTypes(), &item)
		unpacked, d := unpackDnsSecurityProfilesBotnetDomainsToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.DnsSecurityProfilesBotnetDomains", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for DnsSecurityProfilesBotnetDomains ---
func packDnsSecurityProfilesBotnetDomainsListFromSdk(ctx context.Context, sdks []security_services.DnsSecurityProfilesBotnetDomains) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.DnsSecurityProfilesBotnetDomains")
	diags := diag.Diagnostics{}
	var data []models.DnsSecurityProfilesBotnetDomains

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.DnsSecurityProfilesBotnetDomains
		obj, d := packDnsSecurityProfilesBotnetDomainsFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.DnsSecurityProfilesBotnetDomains{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.DnsSecurityProfilesBotnetDomains", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.DnsSecurityProfilesBotnetDomains{}.AttrType(), data)
}

// --- Unpacker for DnsSecurityProfilesBotnetDomainsDnsSecurityCategoriesInner ---
func unpackDnsSecurityProfilesBotnetDomainsDnsSecurityCategoriesInnerToSdk(ctx context.Context, obj types.Object) (*security_services.DnsSecurityProfilesBotnetDomainsDnsSecurityCategoriesInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.DnsSecurityProfilesBotnetDomainsDnsSecurityCategoriesInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.DnsSecurityProfilesBotnetDomainsDnsSecurityCategoriesInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.DnsSecurityProfilesBotnetDomainsDnsSecurityCategoriesInner
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Action.IsNull() && !model.Action.IsUnknown() {
		sdk.Action = model.Action.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Action", "value": *sdk.Action})
	}

	// Handling Primitives
	if !model.LogLevel.IsNull() && !model.LogLevel.IsUnknown() {
		sdk.LogLevel = model.LogLevel.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "LogLevel", "value": *sdk.LogLevel})
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

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.DnsSecurityProfilesBotnetDomainsDnsSecurityCategoriesInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for DnsSecurityProfilesBotnetDomainsDnsSecurityCategoriesInner ---
func packDnsSecurityProfilesBotnetDomainsDnsSecurityCategoriesInnerFromSdk(ctx context.Context, sdk security_services.DnsSecurityProfilesBotnetDomainsDnsSecurityCategoriesInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.DnsSecurityProfilesBotnetDomainsDnsSecurityCategoriesInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.DnsSecurityProfilesBotnetDomainsDnsSecurityCategoriesInner
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Action != nil {
		model.Action = basetypes.NewStringValue(*sdk.Action)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Action", "value": *sdk.Action})
	} else {
		model.Action = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.LogLevel != nil {
		model.LogLevel = basetypes.NewStringValue(*sdk.LogLevel)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "LogLevel", "value": *sdk.LogLevel})
	} else {
		model.LogLevel = basetypes.NewStringNull()
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
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.DnsSecurityProfilesBotnetDomainsDnsSecurityCategoriesInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.DnsSecurityProfilesBotnetDomainsDnsSecurityCategoriesInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for DnsSecurityProfilesBotnetDomainsDnsSecurityCategoriesInner ---
func unpackDnsSecurityProfilesBotnetDomainsDnsSecurityCategoriesInnerListToSdk(ctx context.Context, list types.List) ([]security_services.DnsSecurityProfilesBotnetDomainsDnsSecurityCategoriesInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.DnsSecurityProfilesBotnetDomainsDnsSecurityCategoriesInner")
	diags := diag.Diagnostics{}
	var data []models.DnsSecurityProfilesBotnetDomainsDnsSecurityCategoriesInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.DnsSecurityProfilesBotnetDomainsDnsSecurityCategoriesInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.DnsSecurityProfilesBotnetDomainsDnsSecurityCategoriesInner{}.AttrTypes(), &item)
		unpacked, d := unpackDnsSecurityProfilesBotnetDomainsDnsSecurityCategoriesInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.DnsSecurityProfilesBotnetDomainsDnsSecurityCategoriesInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for DnsSecurityProfilesBotnetDomainsDnsSecurityCategoriesInner ---
func packDnsSecurityProfilesBotnetDomainsDnsSecurityCategoriesInnerListFromSdk(ctx context.Context, sdks []security_services.DnsSecurityProfilesBotnetDomainsDnsSecurityCategoriesInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.DnsSecurityProfilesBotnetDomainsDnsSecurityCategoriesInner")
	diags := diag.Diagnostics{}
	var data []models.DnsSecurityProfilesBotnetDomainsDnsSecurityCategoriesInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.DnsSecurityProfilesBotnetDomainsDnsSecurityCategoriesInner
		obj, d := packDnsSecurityProfilesBotnetDomainsDnsSecurityCategoriesInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.DnsSecurityProfilesBotnetDomainsDnsSecurityCategoriesInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.DnsSecurityProfilesBotnetDomainsDnsSecurityCategoriesInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.DnsSecurityProfilesBotnetDomainsDnsSecurityCategoriesInner{}.AttrType(), data)
}

// --- Unpacker for DnsSecurityProfilesBotnetDomainsListsInner ---
func unpackDnsSecurityProfilesBotnetDomainsListsInnerToSdk(ctx context.Context, obj types.Object) (*security_services.DnsSecurityProfilesBotnetDomainsListsInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.DnsSecurityProfilesBotnetDomainsListsInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.DnsSecurityProfilesBotnetDomainsListsInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.DnsSecurityProfilesBotnetDomainsListsInner
	var d diag.Diagnostics
	// Handling Objects
	if !model.Action.IsNull() && !model.Action.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Action")
		unpacked, d := unpackDnsSecurityProfilesBotnetDomainsListsInnerActionToSdk(ctx, model.Action)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Action"})
		}
		if unpacked != nil {
			sdk.Action = unpacked
		}
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	// Handling Primitives
	if !model.PacketCapture.IsNull() && !model.PacketCapture.IsUnknown() {
		sdk.PacketCapture = model.PacketCapture.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "PacketCapture", "value": *sdk.PacketCapture})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.DnsSecurityProfilesBotnetDomainsListsInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for DnsSecurityProfilesBotnetDomainsListsInner ---
func packDnsSecurityProfilesBotnetDomainsListsInnerFromSdk(ctx context.Context, sdk security_services.DnsSecurityProfilesBotnetDomainsListsInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.DnsSecurityProfilesBotnetDomainsListsInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.DnsSecurityProfilesBotnetDomainsListsInner
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Action != nil {
		tflog.Debug(ctx, "Packing nested object for field Action")
		packed, d := packDnsSecurityProfilesBotnetDomainsListsInnerActionFromSdk(ctx, *sdk.Action)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Action"})
		}
		model.Action = packed
	} else {
		model.Action = basetypes.NewObjectNull(models.DnsSecurityProfilesBotnetDomainsListsInnerAction{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	// Handling Primitives
	// Standard primitive packing
	if sdk.PacketCapture != nil {
		model.PacketCapture = basetypes.NewStringValue(*sdk.PacketCapture)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "PacketCapture", "value": *sdk.PacketCapture})
	} else {
		model.PacketCapture = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.DnsSecurityProfilesBotnetDomainsListsInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.DnsSecurityProfilesBotnetDomainsListsInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for DnsSecurityProfilesBotnetDomainsListsInner ---
func unpackDnsSecurityProfilesBotnetDomainsListsInnerListToSdk(ctx context.Context, list types.List) ([]security_services.DnsSecurityProfilesBotnetDomainsListsInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.DnsSecurityProfilesBotnetDomainsListsInner")
	diags := diag.Diagnostics{}
	var data []models.DnsSecurityProfilesBotnetDomainsListsInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.DnsSecurityProfilesBotnetDomainsListsInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.DnsSecurityProfilesBotnetDomainsListsInner{}.AttrTypes(), &item)
		unpacked, d := unpackDnsSecurityProfilesBotnetDomainsListsInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.DnsSecurityProfilesBotnetDomainsListsInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for DnsSecurityProfilesBotnetDomainsListsInner ---
func packDnsSecurityProfilesBotnetDomainsListsInnerListFromSdk(ctx context.Context, sdks []security_services.DnsSecurityProfilesBotnetDomainsListsInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.DnsSecurityProfilesBotnetDomainsListsInner")
	diags := diag.Diagnostics{}
	var data []models.DnsSecurityProfilesBotnetDomainsListsInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.DnsSecurityProfilesBotnetDomainsListsInner
		obj, d := packDnsSecurityProfilesBotnetDomainsListsInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.DnsSecurityProfilesBotnetDomainsListsInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.DnsSecurityProfilesBotnetDomainsListsInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.DnsSecurityProfilesBotnetDomainsListsInner{}.AttrType(), data)
}

// --- Unpacker for DnsSecurityProfilesBotnetDomainsListsInnerAction ---
func unpackDnsSecurityProfilesBotnetDomainsListsInnerActionToSdk(ctx context.Context, obj types.Object) (*security_services.DnsSecurityProfilesBotnetDomainsListsInnerAction, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.DnsSecurityProfilesBotnetDomainsListsInnerAction", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.DnsSecurityProfilesBotnetDomainsListsInnerAction
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.DnsSecurityProfilesBotnetDomainsListsInnerAction
	var d diag.Diagnostics
	// Handling Typeless Objects
	if !model.Alert.IsNull() && !model.Alert.IsUnknown() {
		tflog.Debug(ctx, "Unpacking typeless object for field Alert")
		sdk.Alert = make(map[string]interface{})
	}

	// Handling Typeless Objects
	if !model.Allow.IsNull() && !model.Allow.IsUnknown() {
		tflog.Debug(ctx, "Unpacking typeless object for field Allow")
		sdk.Allow = make(map[string]interface{})
	}

	// Handling Typeless Objects
	if !model.Block.IsNull() && !model.Block.IsUnknown() {
		tflog.Debug(ctx, "Unpacking typeless object for field Block")
		sdk.Block = make(map[string]interface{})
	}

	// Handling Typeless Objects
	if !model.Sinkhole.IsNull() && !model.Sinkhole.IsUnknown() {
		tflog.Debug(ctx, "Unpacking typeless object for field Sinkhole")
		sdk.Sinkhole = make(map[string]interface{})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.DnsSecurityProfilesBotnetDomainsListsInnerAction", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for DnsSecurityProfilesBotnetDomainsListsInnerAction ---
func packDnsSecurityProfilesBotnetDomainsListsInnerActionFromSdk(ctx context.Context, sdk security_services.DnsSecurityProfilesBotnetDomainsListsInnerAction) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.DnsSecurityProfilesBotnetDomainsListsInnerAction", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.DnsSecurityProfilesBotnetDomainsListsInnerAction
	var d diag.Diagnostics
	// Handling Objects
	// This is a marker object (e.g. CHAP: {}). We just need to create an empty, non-null object.
	if sdk.Alert != nil && !reflect.ValueOf(sdk.Alert).IsNil() {
		tflog.Debug(ctx, "Packing typeless object for field Alert")
		var d diag.Diagnostics
		// Create an empty object with no attributes, which signifies its presence.
		model.Alert, d = basetypes.NewObjectValue(map[string]attr.Type{}, map[string]attr.Value{})
		diags.Append(d...)
	} else {
		// Since this field is part of a oneOf, being nil means it's not selected.
		// We make the object null with an empty attribute map.
		model.Alert = basetypes.NewObjectNull(map[string]attr.Type{})
	}
	// Handling Objects
	// This is a marker object (e.g. CHAP: {}). We just need to create an empty, non-null object.
	if sdk.Allow != nil && !reflect.ValueOf(sdk.Allow).IsNil() {
		tflog.Debug(ctx, "Packing typeless object for field Allow")
		var d diag.Diagnostics
		// Create an empty object with no attributes, which signifies its presence.
		model.Allow, d = basetypes.NewObjectValue(map[string]attr.Type{}, map[string]attr.Value{})
		diags.Append(d...)
	} else {
		// Since this field is part of a oneOf, being nil means it's not selected.
		// We make the object null with an empty attribute map.
		model.Allow = basetypes.NewObjectNull(map[string]attr.Type{})
	}
	// Handling Objects
	// This is a marker object (e.g. CHAP: {}). We just need to create an empty, non-null object.
	if sdk.Block != nil && !reflect.ValueOf(sdk.Block).IsNil() {
		tflog.Debug(ctx, "Packing typeless object for field Block")
		var d diag.Diagnostics
		// Create an empty object with no attributes, which signifies its presence.
		model.Block, d = basetypes.NewObjectValue(map[string]attr.Type{}, map[string]attr.Value{})
		diags.Append(d...)
	} else {
		// Since this field is part of a oneOf, being nil means it's not selected.
		// We make the object null with an empty attribute map.
		model.Block = basetypes.NewObjectNull(map[string]attr.Type{})
	}
	// Handling Objects
	// This is a marker object (e.g. CHAP: {}). We just need to create an empty, non-null object.
	if sdk.Sinkhole != nil && !reflect.ValueOf(sdk.Sinkhole).IsNil() {
		tflog.Debug(ctx, "Packing typeless object for field Sinkhole")
		var d diag.Diagnostics
		// Create an empty object with no attributes, which signifies its presence.
		model.Sinkhole, d = basetypes.NewObjectValue(map[string]attr.Type{}, map[string]attr.Value{})
		diags.Append(d...)
	} else {
		// Since this field is part of a oneOf, being nil means it's not selected.
		// We make the object null with an empty attribute map.
		model.Sinkhole = basetypes.NewObjectNull(map[string]attr.Type{})
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.DnsSecurityProfilesBotnetDomainsListsInnerAction{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.DnsSecurityProfilesBotnetDomainsListsInnerAction", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for DnsSecurityProfilesBotnetDomainsListsInnerAction ---
func unpackDnsSecurityProfilesBotnetDomainsListsInnerActionListToSdk(ctx context.Context, list types.List) ([]security_services.DnsSecurityProfilesBotnetDomainsListsInnerAction, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.DnsSecurityProfilesBotnetDomainsListsInnerAction")
	diags := diag.Diagnostics{}
	var data []models.DnsSecurityProfilesBotnetDomainsListsInnerAction
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.DnsSecurityProfilesBotnetDomainsListsInnerAction, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.DnsSecurityProfilesBotnetDomainsListsInnerAction{}.AttrTypes(), &item)
		unpacked, d := unpackDnsSecurityProfilesBotnetDomainsListsInnerActionToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.DnsSecurityProfilesBotnetDomainsListsInnerAction", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for DnsSecurityProfilesBotnetDomainsListsInnerAction ---
func packDnsSecurityProfilesBotnetDomainsListsInnerActionListFromSdk(ctx context.Context, sdks []security_services.DnsSecurityProfilesBotnetDomainsListsInnerAction) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.DnsSecurityProfilesBotnetDomainsListsInnerAction")
	diags := diag.Diagnostics{}
	var data []models.DnsSecurityProfilesBotnetDomainsListsInnerAction

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.DnsSecurityProfilesBotnetDomainsListsInnerAction
		obj, d := packDnsSecurityProfilesBotnetDomainsListsInnerActionFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.DnsSecurityProfilesBotnetDomainsListsInnerAction{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.DnsSecurityProfilesBotnetDomainsListsInnerAction", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.DnsSecurityProfilesBotnetDomainsListsInnerAction{}.AttrType(), data)
}

// --- Unpacker for DnsSecurityProfilesBotnetDomainsSinkhole ---
func unpackDnsSecurityProfilesBotnetDomainsSinkholeToSdk(ctx context.Context, obj types.Object) (*security_services.DnsSecurityProfilesBotnetDomainsSinkhole, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.DnsSecurityProfilesBotnetDomainsSinkhole", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.DnsSecurityProfilesBotnetDomainsSinkhole
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.DnsSecurityProfilesBotnetDomainsSinkhole
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Ipv4Address.IsNull() && !model.Ipv4Address.IsUnknown() {
		sdk.Ipv4Address = model.Ipv4Address.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Ipv4Address", "value": *sdk.Ipv4Address})
	}

	// Handling Primitives
	if !model.Ipv6Address.IsNull() && !model.Ipv6Address.IsUnknown() {
		sdk.Ipv6Address = model.Ipv6Address.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Ipv6Address", "value": *sdk.Ipv6Address})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.DnsSecurityProfilesBotnetDomainsSinkhole", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for DnsSecurityProfilesBotnetDomainsSinkhole ---
func packDnsSecurityProfilesBotnetDomainsSinkholeFromSdk(ctx context.Context, sdk security_services.DnsSecurityProfilesBotnetDomainsSinkhole) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.DnsSecurityProfilesBotnetDomainsSinkhole", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.DnsSecurityProfilesBotnetDomainsSinkhole
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Ipv4Address != nil {
		model.Ipv4Address = basetypes.NewStringValue(*sdk.Ipv4Address)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Ipv4Address", "value": *sdk.Ipv4Address})
	} else {
		model.Ipv4Address = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Ipv6Address != nil {
		model.Ipv6Address = basetypes.NewStringValue(*sdk.Ipv6Address)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Ipv6Address", "value": *sdk.Ipv6Address})
	} else {
		model.Ipv6Address = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.DnsSecurityProfilesBotnetDomainsSinkhole{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.DnsSecurityProfilesBotnetDomainsSinkhole", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for DnsSecurityProfilesBotnetDomainsSinkhole ---
func unpackDnsSecurityProfilesBotnetDomainsSinkholeListToSdk(ctx context.Context, list types.List) ([]security_services.DnsSecurityProfilesBotnetDomainsSinkhole, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.DnsSecurityProfilesBotnetDomainsSinkhole")
	diags := diag.Diagnostics{}
	var data []models.DnsSecurityProfilesBotnetDomainsSinkhole
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.DnsSecurityProfilesBotnetDomainsSinkhole, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.DnsSecurityProfilesBotnetDomainsSinkhole{}.AttrTypes(), &item)
		unpacked, d := unpackDnsSecurityProfilesBotnetDomainsSinkholeToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.DnsSecurityProfilesBotnetDomainsSinkhole", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for DnsSecurityProfilesBotnetDomainsSinkhole ---
func packDnsSecurityProfilesBotnetDomainsSinkholeListFromSdk(ctx context.Context, sdks []security_services.DnsSecurityProfilesBotnetDomainsSinkhole) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.DnsSecurityProfilesBotnetDomainsSinkhole")
	diags := diag.Diagnostics{}
	var data []models.DnsSecurityProfilesBotnetDomainsSinkhole

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.DnsSecurityProfilesBotnetDomainsSinkhole
		obj, d := packDnsSecurityProfilesBotnetDomainsSinkholeFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.DnsSecurityProfilesBotnetDomainsSinkhole{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.DnsSecurityProfilesBotnetDomainsSinkhole", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.DnsSecurityProfilesBotnetDomainsSinkhole{}.AttrType(), data)
}

// --- Unpacker for DnsSecurityProfilesBotnetDomainsWhitelistInner ---
func unpackDnsSecurityProfilesBotnetDomainsWhitelistInnerToSdk(ctx context.Context, obj types.Object) (*security_services.DnsSecurityProfilesBotnetDomainsWhitelistInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.DnsSecurityProfilesBotnetDomainsWhitelistInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.DnsSecurityProfilesBotnetDomainsWhitelistInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.DnsSecurityProfilesBotnetDomainsWhitelistInner
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Description.IsNull() && !model.Description.IsUnknown() {
		sdk.Description = model.Description.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Description", "value": *sdk.Description})
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.DnsSecurityProfilesBotnetDomainsWhitelistInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for DnsSecurityProfilesBotnetDomainsWhitelistInner ---
func packDnsSecurityProfilesBotnetDomainsWhitelistInnerFromSdk(ctx context.Context, sdk security_services.DnsSecurityProfilesBotnetDomainsWhitelistInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.DnsSecurityProfilesBotnetDomainsWhitelistInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.DnsSecurityProfilesBotnetDomainsWhitelistInner
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
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.DnsSecurityProfilesBotnetDomainsWhitelistInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.DnsSecurityProfilesBotnetDomainsWhitelistInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for DnsSecurityProfilesBotnetDomainsWhitelistInner ---
func unpackDnsSecurityProfilesBotnetDomainsWhitelistInnerListToSdk(ctx context.Context, list types.List) ([]security_services.DnsSecurityProfilesBotnetDomainsWhitelistInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.DnsSecurityProfilesBotnetDomainsWhitelistInner")
	diags := diag.Diagnostics{}
	var data []models.DnsSecurityProfilesBotnetDomainsWhitelistInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.DnsSecurityProfilesBotnetDomainsWhitelistInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.DnsSecurityProfilesBotnetDomainsWhitelistInner{}.AttrTypes(), &item)
		unpacked, d := unpackDnsSecurityProfilesBotnetDomainsWhitelistInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.DnsSecurityProfilesBotnetDomainsWhitelistInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for DnsSecurityProfilesBotnetDomainsWhitelistInner ---
func packDnsSecurityProfilesBotnetDomainsWhitelistInnerListFromSdk(ctx context.Context, sdks []security_services.DnsSecurityProfilesBotnetDomainsWhitelistInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.DnsSecurityProfilesBotnetDomainsWhitelistInner")
	diags := diag.Diagnostics{}
	var data []models.DnsSecurityProfilesBotnetDomainsWhitelistInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.DnsSecurityProfilesBotnetDomainsWhitelistInner
		obj, d := packDnsSecurityProfilesBotnetDomainsWhitelistInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.DnsSecurityProfilesBotnetDomainsWhitelistInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.DnsSecurityProfilesBotnetDomainsWhitelistInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.DnsSecurityProfilesBotnetDomainsWhitelistInner{}.AttrType(), data)
}
