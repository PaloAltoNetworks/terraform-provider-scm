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

// --- Unpacker for DnsProxies ---
func unpackDnsProxiesToSdk(ctx context.Context, obj types.Object) (*network_services.DnsProxies, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.DnsProxies", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.DnsProxies
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.DnsProxies
	var d diag.Diagnostics

	// Handling Objects
	if !model.Cache.IsNull() && !model.Cache.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Cache")
		unpacked, d := unpackDnsProxiesCacheToSdk(ctx, model.Cache)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Cache"})
		}
		if unpacked != nil {
			sdk.Cache = unpacked
		}
	}

	// Handling Objects
	if !model.Default.IsNull() && !model.Default.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Default")
		unpacked, d := unpackDnsProxiesDefaultToSdk(ctx, model.Default)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Default"})
		}
		if unpacked != nil {
			sdk.Default = *unpacked
		}
	}

	// Handling Primitives
	if !model.Device.IsNull() && !model.Device.IsUnknown() {
		sdk.Device = model.Device.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Device", "value": *sdk.Device})
	}

	// Handling Lists
	if !model.DomainServers.IsNull() && !model.DomainServers.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field DomainServers")
		unpacked, d := unpackDnsProxiesDomainServersInnerListToSdk(ctx, model.DomainServers)
		diags.Append(d...)
		sdk.DomainServers = unpacked
	}

	// Handling Primitives
	if !model.Enabled.IsNull() && !model.Enabled.IsUnknown() {
		sdk.Enabled = model.Enabled.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Enabled", "value": *sdk.Enabled})
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
	if !model.Interface.IsNull() && !model.Interface.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Interface")
		diags.Append(model.Interface.ElementsAs(ctx, &sdk.Interface, false)...)
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

	// Handling Lists
	if !model.StaticEntries.IsNull() && !model.StaticEntries.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field StaticEntries")
		unpacked, d := unpackDnsProxiesStaticEntriesInnerListToSdk(ctx, model.StaticEntries)
		diags.Append(d...)
		sdk.StaticEntries = unpacked
	}

	// Handling Objects
	if !model.TcpQueries.IsNull() && !model.TcpQueries.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field TcpQueries")
		unpacked, d := unpackDnsProxiesTcpQueriesToSdk(ctx, model.TcpQueries)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "TcpQueries"})
		}
		if unpacked != nil {
			sdk.TcpQueries = unpacked
		}
	}

	// Handling Objects
	if !model.UdpQueries.IsNull() && !model.UdpQueries.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field UdpQueries")
		unpacked, d := unpackDnsProxiesUdpQueriesToSdk(ctx, model.UdpQueries)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "UdpQueries"})
		}
		if unpacked != nil {
			sdk.UdpQueries = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.DnsProxies", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for DnsProxies ---
func packDnsProxiesFromSdk(ctx context.Context, sdk network_services.DnsProxies) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.DnsProxies", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.DnsProxies
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Cache != nil {
		tflog.Debug(ctx, "Packing nested object for field Cache")
		packed, d := packDnsProxiesCacheFromSdk(ctx, *sdk.Cache)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Cache"})
		}
		model.Cache = packed
	} else {
		model.Cache = basetypes.NewObjectNull(models.DnsProxiesCache{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	// Logic for non-pointer / value-type nested objects
	if !reflect.ValueOf(sdk.Default).IsZero() {
		tflog.Debug(ctx, "Packing nested object for field Default")
		packed, d := packDnsProxiesDefaultFromSdk(ctx, sdk.Default)
		diags.Append(d...)
		model.Default = packed
	} else {
		model.Default = basetypes.NewObjectNull(models.DnsProxiesDefault{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Device != nil {
		model.Device = basetypes.NewStringValue(*sdk.Device)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Device", "value": *sdk.Device})
	} else {
		model.Device = basetypes.NewStringNull()
	}
	// Handling Lists
	if sdk.DomainServers != nil {
		tflog.Debug(ctx, "Packing list of objects for field DomainServers")
		packed, d := packDnsProxiesDomainServersInnerListFromSdk(ctx, sdk.DomainServers)
		diags.Append(d...)
		model.DomainServers = packed
	} else {
		model.DomainServers = basetypes.NewListNull(models.DnsProxiesDomainServersInner{}.AttrType())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Enabled != nil {
		model.Enabled = basetypes.NewBoolValue(*sdk.Enabled)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Enabled", "value": *sdk.Enabled})
	} else {
		model.Enabled = basetypes.NewBoolNull()
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
	if sdk.Interface != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Interface")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Interface, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Interface)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Interface = basetypes.NewListNull(elemType)
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
	// Handling Lists
	if sdk.StaticEntries != nil {
		tflog.Debug(ctx, "Packing list of objects for field StaticEntries")
		packed, d := packDnsProxiesStaticEntriesInnerListFromSdk(ctx, sdk.StaticEntries)
		diags.Append(d...)
		model.StaticEntries = packed
	} else {
		model.StaticEntries = basetypes.NewListNull(models.DnsProxiesStaticEntriesInner{}.AttrType())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.TcpQueries != nil {
		tflog.Debug(ctx, "Packing nested object for field TcpQueries")
		packed, d := packDnsProxiesTcpQueriesFromSdk(ctx, *sdk.TcpQueries)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "TcpQueries"})
		}
		model.TcpQueries = packed
	} else {
		model.TcpQueries = basetypes.NewObjectNull(models.DnsProxiesTcpQueries{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.UdpQueries != nil {
		tflog.Debug(ctx, "Packing nested object for field UdpQueries")
		packed, d := packDnsProxiesUdpQueriesFromSdk(ctx, *sdk.UdpQueries)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "UdpQueries"})
		}
		model.UdpQueries = packed
	} else {
		model.UdpQueries = basetypes.NewObjectNull(models.DnsProxiesUdpQueries{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.DnsProxies{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.DnsProxies", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for DnsProxies ---
func unpackDnsProxiesListToSdk(ctx context.Context, list types.List) ([]network_services.DnsProxies, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.DnsProxies")
	diags := diag.Diagnostics{}
	var data []models.DnsProxies
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.DnsProxies, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.DnsProxies{}.AttrTypes(), &item)
		unpacked, d := unpackDnsProxiesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.DnsProxies", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for DnsProxies ---
func packDnsProxiesListFromSdk(ctx context.Context, sdks []network_services.DnsProxies) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.DnsProxies")
	diags := diag.Diagnostics{}
	var data []models.DnsProxies

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.DnsProxies
		obj, d := packDnsProxiesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.DnsProxies{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.DnsProxies", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.DnsProxies{}.AttrType(), data)
}

// --- Unpacker for DnsProxiesCache ---
func unpackDnsProxiesCacheToSdk(ctx context.Context, obj types.Object) (*network_services.DnsProxiesCache, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.DnsProxiesCache", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.DnsProxiesCache
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.DnsProxiesCache
	var d diag.Diagnostics
	// Handling Primitives
	if !model.CacheEdns.IsNull() && !model.CacheEdns.IsUnknown() {
		sdk.CacheEdns = model.CacheEdns.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "CacheEdns", "value": *sdk.CacheEdns})
	}

	// Handling Primitives
	if !model.Enabled.IsNull() && !model.Enabled.IsUnknown() {
		sdk.Enabled = model.Enabled.ValueBool()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Enabled", "value": sdk.Enabled})
	}

	// Handling Objects
	if !model.MaxTtl.IsNull() && !model.MaxTtl.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field MaxTtl")
		unpacked, d := unpackDnsProxiesCacheMaxTtlToSdk(ctx, model.MaxTtl)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "MaxTtl"})
		}
		if unpacked != nil {
			sdk.MaxTtl = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.DnsProxiesCache", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for DnsProxiesCache ---
func packDnsProxiesCacheFromSdk(ctx context.Context, sdk network_services.DnsProxiesCache) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.DnsProxiesCache", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.DnsProxiesCache
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.CacheEdns != nil {
		model.CacheEdns = basetypes.NewBoolValue(*sdk.CacheEdns)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "CacheEdns", "value": *sdk.CacheEdns})
	} else {
		model.CacheEdns = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	model.Enabled = basetypes.NewBoolValue(sdk.Enabled)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Enabled", "value": sdk.Enabled})
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.MaxTtl != nil {
		tflog.Debug(ctx, "Packing nested object for field MaxTtl")
		packed, d := packDnsProxiesCacheMaxTtlFromSdk(ctx, *sdk.MaxTtl)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "MaxTtl"})
		}
		model.MaxTtl = packed
	} else {
		model.MaxTtl = basetypes.NewObjectNull(models.DnsProxiesCacheMaxTtl{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.DnsProxiesCache{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.DnsProxiesCache", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for DnsProxiesCache ---
func unpackDnsProxiesCacheListToSdk(ctx context.Context, list types.List) ([]network_services.DnsProxiesCache, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.DnsProxiesCache")
	diags := diag.Diagnostics{}
	var data []models.DnsProxiesCache
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.DnsProxiesCache, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.DnsProxiesCache{}.AttrTypes(), &item)
		unpacked, d := unpackDnsProxiesCacheToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.DnsProxiesCache", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for DnsProxiesCache ---
func packDnsProxiesCacheListFromSdk(ctx context.Context, sdks []network_services.DnsProxiesCache) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.DnsProxiesCache")
	diags := diag.Diagnostics{}
	var data []models.DnsProxiesCache

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.DnsProxiesCache
		obj, d := packDnsProxiesCacheFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.DnsProxiesCache{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.DnsProxiesCache", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.DnsProxiesCache{}.AttrType(), data)
}

// --- Unpacker for DnsProxiesCacheMaxTtl ---
func unpackDnsProxiesCacheMaxTtlToSdk(ctx context.Context, obj types.Object) (*network_services.DnsProxiesCacheMaxTtl, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.DnsProxiesCacheMaxTtl", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.DnsProxiesCacheMaxTtl
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.DnsProxiesCacheMaxTtl
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Enabled.IsNull() && !model.Enabled.IsUnknown() {
		sdk.Enabled = model.Enabled.ValueBool()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Enabled", "value": sdk.Enabled})
	}

	// Handling Primitives
	if !model.TimeToLive.IsNull() && !model.TimeToLive.IsUnknown() {
		val := int32(model.TimeToLive.ValueInt64())
		sdk.TimeToLive = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "TimeToLive", "value": *sdk.TimeToLive})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.DnsProxiesCacheMaxTtl", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for DnsProxiesCacheMaxTtl ---
func packDnsProxiesCacheMaxTtlFromSdk(ctx context.Context, sdk network_services.DnsProxiesCacheMaxTtl) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.DnsProxiesCacheMaxTtl", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.DnsProxiesCacheMaxTtl
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.Enabled = basetypes.NewBoolValue(sdk.Enabled)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Enabled", "value": sdk.Enabled})
	// Handling Primitives
	// Standard primitive packing
	if sdk.TimeToLive != nil {
		model.TimeToLive = basetypes.NewInt64Value(int64(*sdk.TimeToLive))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "TimeToLive", "value": *sdk.TimeToLive})
	} else {
		model.TimeToLive = basetypes.NewInt64Null()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.DnsProxiesCacheMaxTtl{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.DnsProxiesCacheMaxTtl", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for DnsProxiesCacheMaxTtl ---
func unpackDnsProxiesCacheMaxTtlListToSdk(ctx context.Context, list types.List) ([]network_services.DnsProxiesCacheMaxTtl, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.DnsProxiesCacheMaxTtl")
	diags := diag.Diagnostics{}
	var data []models.DnsProxiesCacheMaxTtl
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.DnsProxiesCacheMaxTtl, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.DnsProxiesCacheMaxTtl{}.AttrTypes(), &item)
		unpacked, d := unpackDnsProxiesCacheMaxTtlToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.DnsProxiesCacheMaxTtl", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for DnsProxiesCacheMaxTtl ---
func packDnsProxiesCacheMaxTtlListFromSdk(ctx context.Context, sdks []network_services.DnsProxiesCacheMaxTtl) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.DnsProxiesCacheMaxTtl")
	diags := diag.Diagnostics{}
	var data []models.DnsProxiesCacheMaxTtl

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.DnsProxiesCacheMaxTtl
		obj, d := packDnsProxiesCacheMaxTtlFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.DnsProxiesCacheMaxTtl{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.DnsProxiesCacheMaxTtl", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.DnsProxiesCacheMaxTtl{}.AttrType(), data)
}

// --- Unpacker for DnsProxiesDefault ---
func unpackDnsProxiesDefaultToSdk(ctx context.Context, obj types.Object) (*network_services.DnsProxiesDefault, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.DnsProxiesDefault", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.DnsProxiesDefault
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.DnsProxiesDefault
	var d diag.Diagnostics
	// Handling Objects
	if !model.Inheritance.IsNull() && !model.Inheritance.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Inheritance")
		unpacked, d := unpackDnsProxiesDefaultInheritanceToSdk(ctx, model.Inheritance)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Inheritance"})
		}
		if unpacked != nil {
			sdk.Inheritance = unpacked
		}
	}

	// Handling Primitives
	if !model.Primary.IsNull() && !model.Primary.IsUnknown() {
		sdk.Primary = model.Primary.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Primary", "value": sdk.Primary})
	}

	// Handling Primitives
	if !model.Secondary.IsNull() && !model.Secondary.IsUnknown() {
		sdk.Secondary = model.Secondary.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Secondary", "value": *sdk.Secondary})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.DnsProxiesDefault", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for DnsProxiesDefault ---
func packDnsProxiesDefaultFromSdk(ctx context.Context, sdk network_services.DnsProxiesDefault) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.DnsProxiesDefault", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.DnsProxiesDefault
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Inheritance != nil {
		tflog.Debug(ctx, "Packing nested object for field Inheritance")
		packed, d := packDnsProxiesDefaultInheritanceFromSdk(ctx, *sdk.Inheritance)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Inheritance"})
		}
		model.Inheritance = packed
	} else {
		model.Inheritance = basetypes.NewObjectNull(models.DnsProxiesDefaultInheritance{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	model.Primary = basetypes.NewStringValue(sdk.Primary)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Primary", "value": sdk.Primary})
	// Handling Primitives
	// Standard primitive packing
	if sdk.Secondary != nil {
		model.Secondary = basetypes.NewStringValue(*sdk.Secondary)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Secondary", "value": *sdk.Secondary})
	} else {
		model.Secondary = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.DnsProxiesDefault{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.DnsProxiesDefault", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for DnsProxiesDefault ---
func unpackDnsProxiesDefaultListToSdk(ctx context.Context, list types.List) ([]network_services.DnsProxiesDefault, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.DnsProxiesDefault")
	diags := diag.Diagnostics{}
	var data []models.DnsProxiesDefault
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.DnsProxiesDefault, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.DnsProxiesDefault{}.AttrTypes(), &item)
		unpacked, d := unpackDnsProxiesDefaultToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.DnsProxiesDefault", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for DnsProxiesDefault ---
func packDnsProxiesDefaultListFromSdk(ctx context.Context, sdks []network_services.DnsProxiesDefault) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.DnsProxiesDefault")
	diags := diag.Diagnostics{}
	var data []models.DnsProxiesDefault

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.DnsProxiesDefault
		obj, d := packDnsProxiesDefaultFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.DnsProxiesDefault{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.DnsProxiesDefault", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.DnsProxiesDefault{}.AttrType(), data)
}

// --- Unpacker for DnsProxiesDefaultInheritance ---
func unpackDnsProxiesDefaultInheritanceToSdk(ctx context.Context, obj types.Object) (*network_services.DnsProxiesDefaultInheritance, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.DnsProxiesDefaultInheritance", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.DnsProxiesDefaultInheritance
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.DnsProxiesDefaultInheritance
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Source.IsNull() && !model.Source.IsUnknown() {
		sdk.Source = model.Source.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Source", "value": *sdk.Source})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.DnsProxiesDefaultInheritance", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for DnsProxiesDefaultInheritance ---
func packDnsProxiesDefaultInheritanceFromSdk(ctx context.Context, sdk network_services.DnsProxiesDefaultInheritance) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.DnsProxiesDefaultInheritance", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.DnsProxiesDefaultInheritance
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Source != nil {
		model.Source = basetypes.NewStringValue(*sdk.Source)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Source", "value": *sdk.Source})
	} else {
		model.Source = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.DnsProxiesDefaultInheritance{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.DnsProxiesDefaultInheritance", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for DnsProxiesDefaultInheritance ---
func unpackDnsProxiesDefaultInheritanceListToSdk(ctx context.Context, list types.List) ([]network_services.DnsProxiesDefaultInheritance, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.DnsProxiesDefaultInheritance")
	diags := diag.Diagnostics{}
	var data []models.DnsProxiesDefaultInheritance
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.DnsProxiesDefaultInheritance, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.DnsProxiesDefaultInheritance{}.AttrTypes(), &item)
		unpacked, d := unpackDnsProxiesDefaultInheritanceToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.DnsProxiesDefaultInheritance", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for DnsProxiesDefaultInheritance ---
func packDnsProxiesDefaultInheritanceListFromSdk(ctx context.Context, sdks []network_services.DnsProxiesDefaultInheritance) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.DnsProxiesDefaultInheritance")
	diags := diag.Diagnostics{}
	var data []models.DnsProxiesDefaultInheritance

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.DnsProxiesDefaultInheritance
		obj, d := packDnsProxiesDefaultInheritanceFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.DnsProxiesDefaultInheritance{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.DnsProxiesDefaultInheritance", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.DnsProxiesDefaultInheritance{}.AttrType(), data)
}

// --- Unpacker for DnsProxiesDomainServersInner ---
func unpackDnsProxiesDomainServersInnerToSdk(ctx context.Context, obj types.Object) (*network_services.DnsProxiesDomainServersInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.DnsProxiesDomainServersInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.DnsProxiesDomainServersInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.DnsProxiesDomainServersInner
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Cacheable.IsNull() && !model.Cacheable.IsUnknown() {
		sdk.Cacheable = model.Cacheable.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Cacheable", "value": *sdk.Cacheable})
	}

	// Handling Lists
	if !model.DomainName.IsNull() && !model.DomainName.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field DomainName")
		diags.Append(model.DomainName.ElementsAs(ctx, &sdk.DomainName, false)...)
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	// Handling Primitives
	if !model.Primary.IsNull() && !model.Primary.IsUnknown() {
		sdk.Primary = model.Primary.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Primary", "value": sdk.Primary})
	}

	// Handling Primitives
	if !model.Secondary.IsNull() && !model.Secondary.IsUnknown() {
		sdk.Secondary = model.Secondary.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Secondary", "value": *sdk.Secondary})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.DnsProxiesDomainServersInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for DnsProxiesDomainServersInner ---
func packDnsProxiesDomainServersInnerFromSdk(ctx context.Context, sdk network_services.DnsProxiesDomainServersInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.DnsProxiesDomainServersInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.DnsProxiesDomainServersInner
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Cacheable != nil {
		model.Cacheable = basetypes.NewBoolValue(*sdk.Cacheable)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Cacheable", "value": *sdk.Cacheable})
	} else {
		model.Cacheable = basetypes.NewBoolNull()
	}
	// Handling Lists
	if sdk.DomainName != nil {
		tflog.Debug(ctx, "Packing list of primitives for field DomainName")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.DomainName, d = basetypes.NewListValueFrom(ctx, elemType, sdk.DomainName)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.DomainName = basetypes.NewListNull(elemType)
	}
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	// Handling Primitives
	// Standard primitive packing
	model.Primary = basetypes.NewStringValue(sdk.Primary)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Primary", "value": sdk.Primary})
	// Handling Primitives
	// Standard primitive packing
	if sdk.Secondary != nil {
		model.Secondary = basetypes.NewStringValue(*sdk.Secondary)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Secondary", "value": *sdk.Secondary})
	} else {
		model.Secondary = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.DnsProxiesDomainServersInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.DnsProxiesDomainServersInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for DnsProxiesDomainServersInner ---
func unpackDnsProxiesDomainServersInnerListToSdk(ctx context.Context, list types.List) ([]network_services.DnsProxiesDomainServersInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.DnsProxiesDomainServersInner")
	diags := diag.Diagnostics{}
	var data []models.DnsProxiesDomainServersInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.DnsProxiesDomainServersInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.DnsProxiesDomainServersInner{}.AttrTypes(), &item)
		unpacked, d := unpackDnsProxiesDomainServersInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.DnsProxiesDomainServersInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for DnsProxiesDomainServersInner ---
func packDnsProxiesDomainServersInnerListFromSdk(ctx context.Context, sdks []network_services.DnsProxiesDomainServersInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.DnsProxiesDomainServersInner")
	diags := diag.Diagnostics{}
	var data []models.DnsProxiesDomainServersInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.DnsProxiesDomainServersInner
		obj, d := packDnsProxiesDomainServersInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.DnsProxiesDomainServersInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.DnsProxiesDomainServersInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.DnsProxiesDomainServersInner{}.AttrType(), data)
}

// --- Unpacker for DnsProxiesStaticEntriesInner ---
func unpackDnsProxiesStaticEntriesInnerToSdk(ctx context.Context, obj types.Object) (*network_services.DnsProxiesStaticEntriesInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.DnsProxiesStaticEntriesInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.DnsProxiesStaticEntriesInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.DnsProxiesStaticEntriesInner
	var d diag.Diagnostics
	// Handling Lists
	if !model.Address.IsNull() && !model.Address.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Address")
		diags.Append(model.Address.ElementsAs(ctx, &sdk.Address, false)...)
	}

	// Handling Primitives
	if !model.Domain.IsNull() && !model.Domain.IsUnknown() {
		sdk.Domain = model.Domain.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Domain", "value": sdk.Domain})
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.DnsProxiesStaticEntriesInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for DnsProxiesStaticEntriesInner ---
func packDnsProxiesStaticEntriesInnerFromSdk(ctx context.Context, sdk network_services.DnsProxiesStaticEntriesInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.DnsProxiesStaticEntriesInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.DnsProxiesStaticEntriesInner
	var d diag.Diagnostics
	// Handling Lists
	if sdk.Address != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Address")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Address, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Address)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Address = basetypes.NewListNull(elemType)
	}
	// Handling Primitives
	// Standard primitive packing
	model.Domain = basetypes.NewStringValue(sdk.Domain)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Domain", "value": sdk.Domain})
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.DnsProxiesStaticEntriesInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.DnsProxiesStaticEntriesInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for DnsProxiesStaticEntriesInner ---
func unpackDnsProxiesStaticEntriesInnerListToSdk(ctx context.Context, list types.List) ([]network_services.DnsProxiesStaticEntriesInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.DnsProxiesStaticEntriesInner")
	diags := diag.Diagnostics{}
	var data []models.DnsProxiesStaticEntriesInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.DnsProxiesStaticEntriesInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.DnsProxiesStaticEntriesInner{}.AttrTypes(), &item)
		unpacked, d := unpackDnsProxiesStaticEntriesInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.DnsProxiesStaticEntriesInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for DnsProxiesStaticEntriesInner ---
func packDnsProxiesStaticEntriesInnerListFromSdk(ctx context.Context, sdks []network_services.DnsProxiesStaticEntriesInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.DnsProxiesStaticEntriesInner")
	diags := diag.Diagnostics{}
	var data []models.DnsProxiesStaticEntriesInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.DnsProxiesStaticEntriesInner
		obj, d := packDnsProxiesStaticEntriesInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.DnsProxiesStaticEntriesInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.DnsProxiesStaticEntriesInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.DnsProxiesStaticEntriesInner{}.AttrType(), data)
}

// --- Unpacker for DnsProxiesTcpQueries ---
func unpackDnsProxiesTcpQueriesToSdk(ctx context.Context, obj types.Object) (*network_services.DnsProxiesTcpQueries, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.DnsProxiesTcpQueries", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.DnsProxiesTcpQueries
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.DnsProxiesTcpQueries
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Enabled.IsNull() && !model.Enabled.IsUnknown() {
		sdk.Enabled = model.Enabled.ValueBool()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Enabled", "value": sdk.Enabled})
	}

	// Handling Primitives
	if !model.MaxPendingRequests.IsNull() && !model.MaxPendingRequests.IsUnknown() {
		val := int32(model.MaxPendingRequests.ValueInt64())
		sdk.MaxPendingRequests = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "MaxPendingRequests", "value": *sdk.MaxPendingRequests})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.DnsProxiesTcpQueries", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for DnsProxiesTcpQueries ---
func packDnsProxiesTcpQueriesFromSdk(ctx context.Context, sdk network_services.DnsProxiesTcpQueries) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.DnsProxiesTcpQueries", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.DnsProxiesTcpQueries
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.Enabled = basetypes.NewBoolValue(sdk.Enabled)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Enabled", "value": sdk.Enabled})
	// Handling Primitives
	// Standard primitive packing
	if sdk.MaxPendingRequests != nil {
		model.MaxPendingRequests = basetypes.NewInt64Value(int64(*sdk.MaxPendingRequests))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "MaxPendingRequests", "value": *sdk.MaxPendingRequests})
	} else {
		model.MaxPendingRequests = basetypes.NewInt64Null()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.DnsProxiesTcpQueries{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.DnsProxiesTcpQueries", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for DnsProxiesTcpQueries ---
func unpackDnsProxiesTcpQueriesListToSdk(ctx context.Context, list types.List) ([]network_services.DnsProxiesTcpQueries, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.DnsProxiesTcpQueries")
	diags := diag.Diagnostics{}
	var data []models.DnsProxiesTcpQueries
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.DnsProxiesTcpQueries, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.DnsProxiesTcpQueries{}.AttrTypes(), &item)
		unpacked, d := unpackDnsProxiesTcpQueriesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.DnsProxiesTcpQueries", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for DnsProxiesTcpQueries ---
func packDnsProxiesTcpQueriesListFromSdk(ctx context.Context, sdks []network_services.DnsProxiesTcpQueries) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.DnsProxiesTcpQueries")
	diags := diag.Diagnostics{}
	var data []models.DnsProxiesTcpQueries

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.DnsProxiesTcpQueries
		obj, d := packDnsProxiesTcpQueriesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.DnsProxiesTcpQueries{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.DnsProxiesTcpQueries", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.DnsProxiesTcpQueries{}.AttrType(), data)
}

// --- Unpacker for DnsProxiesUdpQueries ---
func unpackDnsProxiesUdpQueriesToSdk(ctx context.Context, obj types.Object) (*network_services.DnsProxiesUdpQueries, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.DnsProxiesUdpQueries", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.DnsProxiesUdpQueries
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.DnsProxiesUdpQueries
	var d diag.Diagnostics
	// Handling Objects
	if !model.Retries.IsNull() && !model.Retries.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Retries")
		unpacked, d := unpackDnsProxiesUdpQueriesRetriesToSdk(ctx, model.Retries)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Retries"})
		}
		if unpacked != nil {
			sdk.Retries = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.DnsProxiesUdpQueries", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for DnsProxiesUdpQueries ---
func packDnsProxiesUdpQueriesFromSdk(ctx context.Context, sdk network_services.DnsProxiesUdpQueries) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.DnsProxiesUdpQueries", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.DnsProxiesUdpQueries
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Retries != nil {
		tflog.Debug(ctx, "Packing nested object for field Retries")
		packed, d := packDnsProxiesUdpQueriesRetriesFromSdk(ctx, *sdk.Retries)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Retries"})
		}
		model.Retries = packed
	} else {
		model.Retries = basetypes.NewObjectNull(models.DnsProxiesUdpQueriesRetries{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.DnsProxiesUdpQueries{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.DnsProxiesUdpQueries", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for DnsProxiesUdpQueries ---
func unpackDnsProxiesUdpQueriesListToSdk(ctx context.Context, list types.List) ([]network_services.DnsProxiesUdpQueries, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.DnsProxiesUdpQueries")
	diags := diag.Diagnostics{}
	var data []models.DnsProxiesUdpQueries
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.DnsProxiesUdpQueries, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.DnsProxiesUdpQueries{}.AttrTypes(), &item)
		unpacked, d := unpackDnsProxiesUdpQueriesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.DnsProxiesUdpQueries", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for DnsProxiesUdpQueries ---
func packDnsProxiesUdpQueriesListFromSdk(ctx context.Context, sdks []network_services.DnsProxiesUdpQueries) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.DnsProxiesUdpQueries")
	diags := diag.Diagnostics{}
	var data []models.DnsProxiesUdpQueries

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.DnsProxiesUdpQueries
		obj, d := packDnsProxiesUdpQueriesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.DnsProxiesUdpQueries{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.DnsProxiesUdpQueries", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.DnsProxiesUdpQueries{}.AttrType(), data)
}

// --- Unpacker for DnsProxiesUdpQueriesRetries ---
func unpackDnsProxiesUdpQueriesRetriesToSdk(ctx context.Context, obj types.Object) (*network_services.DnsProxiesUdpQueriesRetries, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.DnsProxiesUdpQueriesRetries", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.DnsProxiesUdpQueriesRetries
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.DnsProxiesUdpQueriesRetries
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Attempts.IsNull() && !model.Attempts.IsUnknown() {
		val := int32(model.Attempts.ValueInt64())
		sdk.Attempts = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Attempts", "value": *sdk.Attempts})
	}

	// Handling Primitives
	if !model.Interval.IsNull() && !model.Interval.IsUnknown() {
		val := int32(model.Interval.ValueInt64())
		sdk.Interval = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Interval", "value": *sdk.Interval})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.DnsProxiesUdpQueriesRetries", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for DnsProxiesUdpQueriesRetries ---
func packDnsProxiesUdpQueriesRetriesFromSdk(ctx context.Context, sdk network_services.DnsProxiesUdpQueriesRetries) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.DnsProxiesUdpQueriesRetries", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.DnsProxiesUdpQueriesRetries
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Attempts != nil {
		model.Attempts = basetypes.NewInt64Value(int64(*sdk.Attempts))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Attempts", "value": *sdk.Attempts})
	} else {
		model.Attempts = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Interval != nil {
		model.Interval = basetypes.NewInt64Value(int64(*sdk.Interval))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Interval", "value": *sdk.Interval})
	} else {
		model.Interval = basetypes.NewInt64Null()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.DnsProxiesUdpQueriesRetries{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.DnsProxiesUdpQueriesRetries", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for DnsProxiesUdpQueriesRetries ---
func unpackDnsProxiesUdpQueriesRetriesListToSdk(ctx context.Context, list types.List) ([]network_services.DnsProxiesUdpQueriesRetries, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.DnsProxiesUdpQueriesRetries")
	diags := diag.Diagnostics{}
	var data []models.DnsProxiesUdpQueriesRetries
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.DnsProxiesUdpQueriesRetries, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.DnsProxiesUdpQueriesRetries{}.AttrTypes(), &item)
		unpacked, d := unpackDnsProxiesUdpQueriesRetriesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.DnsProxiesUdpQueriesRetries", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for DnsProxiesUdpQueriesRetries ---
func packDnsProxiesUdpQueriesRetriesListFromSdk(ctx context.Context, sdks []network_services.DnsProxiesUdpQueriesRetries) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.DnsProxiesUdpQueriesRetries")
	diags := diag.Diagnostics{}
	var data []models.DnsProxiesUdpQueriesRetries

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.DnsProxiesUdpQueriesRetries
		obj, d := packDnsProxiesUdpQueriesRetriesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.DnsProxiesUdpQueriesRetries{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.DnsProxiesUdpQueriesRetries", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.DnsProxiesUdpQueriesRetries{}.AttrType(), data)
}
