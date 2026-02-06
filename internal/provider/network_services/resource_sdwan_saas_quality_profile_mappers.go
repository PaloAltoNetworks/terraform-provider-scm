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

// --- Unpacker for SdwanSaasQualityProfiles ---
func unpackSdwanSaasQualityProfilesToSdk(ctx context.Context, obj types.Object) (*network_services.SdwanSaasQualityProfiles, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.SdwanSaasQualityProfiles", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.SdwanSaasQualityProfiles
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.SdwanSaasQualityProfiles
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

	// Handling Objects
	if !model.MonitorMode.IsNull() && !model.MonitorMode.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field MonitorMode")
		unpacked, d := unpackSdwanSaasQualityProfilesMonitorModeToSdk(ctx, model.MonitorMode)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "MonitorMode"})
		}
		if unpacked != nil {
			sdk.MonitorMode = *unpacked
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

	tflog.Debug(ctx, "Exiting unpack helper for models.SdwanSaasQualityProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for SdwanSaasQualityProfiles ---
func packSdwanSaasQualityProfilesFromSdk(ctx context.Context, sdk network_services.SdwanSaasQualityProfiles) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.SdwanSaasQualityProfiles", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.SdwanSaasQualityProfiles
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
	// Handling Objects
	// This is a regular nested object that has its own packer.
	// Logic for non-pointer / value-type nested objects
	if !reflect.ValueOf(sdk.MonitorMode).IsZero() {
		tflog.Debug(ctx, "Packing nested object for field MonitorMode")
		packed, d := packSdwanSaasQualityProfilesMonitorModeFromSdk(ctx, sdk.MonitorMode)
		diags.Append(d...)
		model.MonitorMode = packed
	} else {
		model.MonitorMode = basetypes.NewObjectNull(models.SdwanSaasQualityProfilesMonitorMode{}.AttrTypes())
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

	obj, d := types.ObjectValueFrom(ctx, models.SdwanSaasQualityProfiles{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.SdwanSaasQualityProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for SdwanSaasQualityProfiles ---
func unpackSdwanSaasQualityProfilesListToSdk(ctx context.Context, list types.List) ([]network_services.SdwanSaasQualityProfiles, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.SdwanSaasQualityProfiles")
	diags := diag.Diagnostics{}
	var data []models.SdwanSaasQualityProfiles
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.SdwanSaasQualityProfiles, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.SdwanSaasQualityProfiles{}.AttrTypes(), &item)
		unpacked, d := unpackSdwanSaasQualityProfilesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.SdwanSaasQualityProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for SdwanSaasQualityProfiles ---
func packSdwanSaasQualityProfilesListFromSdk(ctx context.Context, sdks []network_services.SdwanSaasQualityProfiles) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.SdwanSaasQualityProfiles")
	diags := diag.Diagnostics{}
	var data []models.SdwanSaasQualityProfiles

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.SdwanSaasQualityProfiles
		obj, d := packSdwanSaasQualityProfilesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.SdwanSaasQualityProfiles{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.SdwanSaasQualityProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.SdwanSaasQualityProfiles{}.AttrType(), data)
}

// --- Unpacker for SdwanSaasQualityProfilesMonitorMode ---
func unpackSdwanSaasQualityProfilesMonitorModeToSdk(ctx context.Context, obj types.Object) (*network_services.SdwanSaasQualityProfilesMonitorMode, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.SdwanSaasQualityProfilesMonitorMode", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.SdwanSaasQualityProfilesMonitorMode
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.SdwanSaasQualityProfilesMonitorMode
	var d diag.Diagnostics
	// Handling Typeless Objects
	if !model.Adaptive.IsNull() && !model.Adaptive.IsUnknown() {
		tflog.Debug(ctx, "Unpacking typeless object for field Adaptive")
		sdk.Adaptive = make(map[string]interface{})
	}

	// Handling Objects
	if !model.HttpHttps.IsNull() && !model.HttpHttps.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field HttpHttps")
		unpacked, d := unpackSdwanSaasQualityProfilesMonitorModeHttpHttpsToSdk(ctx, model.HttpHttps)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "HttpHttps"})
		}
		if unpacked != nil {
			sdk.HttpHttps = unpacked
		}
	}

	// Handling Objects
	if !model.StaticIp.IsNull() && !model.StaticIp.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field StaticIp")
		unpacked, d := unpackSdwanSaasQualityProfilesMonitorModeStaticIpToSdk(ctx, model.StaticIp)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "StaticIp"})
		}
		if unpacked != nil {
			sdk.StaticIp = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.SdwanSaasQualityProfilesMonitorMode", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for SdwanSaasQualityProfilesMonitorMode ---
func packSdwanSaasQualityProfilesMonitorModeFromSdk(ctx context.Context, sdk network_services.SdwanSaasQualityProfilesMonitorMode) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.SdwanSaasQualityProfilesMonitorMode", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.SdwanSaasQualityProfilesMonitorMode
	var d diag.Diagnostics
	// Handling Objects
	// This is a marker object (e.g. CHAP: {}). We just need to create an empty, non-null object.
	if sdk.Adaptive != nil && !reflect.ValueOf(sdk.Adaptive).IsNil() {
		tflog.Debug(ctx, "Packing typeless object for field Adaptive")
		var d diag.Diagnostics
		// Create an empty object with no attributes, which signifies its presence.
		model.Adaptive, d = basetypes.NewObjectValue(map[string]attr.Type{}, map[string]attr.Value{})
		diags.Append(d...)
	} else {
		// Since this field is part of a oneOf, being nil means it's not selected.
		// We make the object null with an empty attribute map.
		model.Adaptive = basetypes.NewObjectNull(map[string]attr.Type{})
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.HttpHttps != nil {
		tflog.Debug(ctx, "Packing nested object for field HttpHttps")
		packed, d := packSdwanSaasQualityProfilesMonitorModeHttpHttpsFromSdk(ctx, *sdk.HttpHttps)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "HttpHttps"})
		}
		model.HttpHttps = packed
	} else {
		model.HttpHttps = basetypes.NewObjectNull(models.SdwanSaasQualityProfilesMonitorModeHttpHttps{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.StaticIp != nil {
		tflog.Debug(ctx, "Packing nested object for field StaticIp")
		packed, d := packSdwanSaasQualityProfilesMonitorModeStaticIpFromSdk(ctx, *sdk.StaticIp)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "StaticIp"})
		}
		model.StaticIp = packed
	} else {
		model.StaticIp = basetypes.NewObjectNull(models.SdwanSaasQualityProfilesMonitorModeStaticIp{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.SdwanSaasQualityProfilesMonitorMode{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.SdwanSaasQualityProfilesMonitorMode", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for SdwanSaasQualityProfilesMonitorMode ---
func unpackSdwanSaasQualityProfilesMonitorModeListToSdk(ctx context.Context, list types.List) ([]network_services.SdwanSaasQualityProfilesMonitorMode, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.SdwanSaasQualityProfilesMonitorMode")
	diags := diag.Diagnostics{}
	var data []models.SdwanSaasQualityProfilesMonitorMode
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.SdwanSaasQualityProfilesMonitorMode, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.SdwanSaasQualityProfilesMonitorMode{}.AttrTypes(), &item)
		unpacked, d := unpackSdwanSaasQualityProfilesMonitorModeToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.SdwanSaasQualityProfilesMonitorMode", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for SdwanSaasQualityProfilesMonitorMode ---
func packSdwanSaasQualityProfilesMonitorModeListFromSdk(ctx context.Context, sdks []network_services.SdwanSaasQualityProfilesMonitorMode) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.SdwanSaasQualityProfilesMonitorMode")
	diags := diag.Diagnostics{}
	var data []models.SdwanSaasQualityProfilesMonitorMode

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.SdwanSaasQualityProfilesMonitorMode
		obj, d := packSdwanSaasQualityProfilesMonitorModeFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.SdwanSaasQualityProfilesMonitorMode{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.SdwanSaasQualityProfilesMonitorMode", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.SdwanSaasQualityProfilesMonitorMode{}.AttrType(), data)
}

// --- Unpacker for SdwanSaasQualityProfilesMonitorModeHttpHttps ---
func unpackSdwanSaasQualityProfilesMonitorModeHttpHttpsToSdk(ctx context.Context, obj types.Object) (*network_services.SdwanSaasQualityProfilesMonitorModeHttpHttps, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.SdwanSaasQualityProfilesMonitorModeHttpHttps", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.SdwanSaasQualityProfilesMonitorModeHttpHttps
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.SdwanSaasQualityProfilesMonitorModeHttpHttps
	var d diag.Diagnostics
	// Handling Primitives
	if !model.MonitoredUrl.IsNull() && !model.MonitoredUrl.IsUnknown() {
		sdk.MonitoredUrl = model.MonitoredUrl.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "MonitoredUrl", "value": sdk.MonitoredUrl})
	}

	// Handling Primitives
	if !model.ProbeInterval.IsNull() && !model.ProbeInterval.IsUnknown() {
		sdk.ProbeInterval = int32(model.ProbeInterval.ValueInt64())
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "ProbeInterval", "value": sdk.ProbeInterval})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.SdwanSaasQualityProfilesMonitorModeHttpHttps", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for SdwanSaasQualityProfilesMonitorModeHttpHttps ---
func packSdwanSaasQualityProfilesMonitorModeHttpHttpsFromSdk(ctx context.Context, sdk network_services.SdwanSaasQualityProfilesMonitorModeHttpHttps) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.SdwanSaasQualityProfilesMonitorModeHttpHttps", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.SdwanSaasQualityProfilesMonitorModeHttpHttps
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.MonitoredUrl = basetypes.NewStringValue(sdk.MonitoredUrl)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "MonitoredUrl", "value": sdk.MonitoredUrl})
	// Handling Primitives
	// Standard primitive packing
	model.ProbeInterval = basetypes.NewInt64Value(int64(sdk.ProbeInterval))
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.SdwanSaasQualityProfilesMonitorModeHttpHttps{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.SdwanSaasQualityProfilesMonitorModeHttpHttps", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for SdwanSaasQualityProfilesMonitorModeHttpHttps ---
func unpackSdwanSaasQualityProfilesMonitorModeHttpHttpsListToSdk(ctx context.Context, list types.List) ([]network_services.SdwanSaasQualityProfilesMonitorModeHttpHttps, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.SdwanSaasQualityProfilesMonitorModeHttpHttps")
	diags := diag.Diagnostics{}
	var data []models.SdwanSaasQualityProfilesMonitorModeHttpHttps
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.SdwanSaasQualityProfilesMonitorModeHttpHttps, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.SdwanSaasQualityProfilesMonitorModeHttpHttps{}.AttrTypes(), &item)
		unpacked, d := unpackSdwanSaasQualityProfilesMonitorModeHttpHttpsToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.SdwanSaasQualityProfilesMonitorModeHttpHttps", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for SdwanSaasQualityProfilesMonitorModeHttpHttps ---
func packSdwanSaasQualityProfilesMonitorModeHttpHttpsListFromSdk(ctx context.Context, sdks []network_services.SdwanSaasQualityProfilesMonitorModeHttpHttps) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.SdwanSaasQualityProfilesMonitorModeHttpHttps")
	diags := diag.Diagnostics{}
	var data []models.SdwanSaasQualityProfilesMonitorModeHttpHttps

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.SdwanSaasQualityProfilesMonitorModeHttpHttps
		obj, d := packSdwanSaasQualityProfilesMonitorModeHttpHttpsFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.SdwanSaasQualityProfilesMonitorModeHttpHttps{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.SdwanSaasQualityProfilesMonitorModeHttpHttps", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.SdwanSaasQualityProfilesMonitorModeHttpHttps{}.AttrType(), data)
}

// --- Unpacker for SdwanSaasQualityProfilesMonitorModeStaticIp ---
func unpackSdwanSaasQualityProfilesMonitorModeStaticIpToSdk(ctx context.Context, obj types.Object) (*network_services.SdwanSaasQualityProfilesMonitorModeStaticIp, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.SdwanSaasQualityProfilesMonitorModeStaticIp", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.SdwanSaasQualityProfilesMonitorModeStaticIp
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.SdwanSaasQualityProfilesMonitorModeStaticIp
	var d diag.Diagnostics
	// Handling Objects
	if !model.Fqdn.IsNull() && !model.Fqdn.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Fqdn")
		unpacked, d := unpackSdwanSaasQualityProfilesMonitorModeStaticIpFqdnToSdk(ctx, model.Fqdn)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Fqdn"})
		}
		if unpacked != nil {
			sdk.Fqdn = unpacked
		}
	}

	// Handling Lists
	if !model.IpAddress.IsNull() && !model.IpAddress.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field IpAddress")
		unpacked, d := unpackSdwanSaasQualityProfilesMonitorModeStaticIpIpAddressInnerListToSdk(ctx, model.IpAddress)
		diags.Append(d...)
		sdk.IpAddress = unpacked
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.SdwanSaasQualityProfilesMonitorModeStaticIp", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for SdwanSaasQualityProfilesMonitorModeStaticIp ---
func packSdwanSaasQualityProfilesMonitorModeStaticIpFromSdk(ctx context.Context, sdk network_services.SdwanSaasQualityProfilesMonitorModeStaticIp) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.SdwanSaasQualityProfilesMonitorModeStaticIp", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.SdwanSaasQualityProfilesMonitorModeStaticIp
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Fqdn != nil {
		tflog.Debug(ctx, "Packing nested object for field Fqdn")
		packed, d := packSdwanSaasQualityProfilesMonitorModeStaticIpFqdnFromSdk(ctx, *sdk.Fqdn)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Fqdn"})
		}
		model.Fqdn = packed
	} else {
		model.Fqdn = basetypes.NewObjectNull(models.SdwanSaasQualityProfilesMonitorModeStaticIpFqdn{}.AttrTypes())
	}
	// Handling Lists
	if sdk.IpAddress != nil {
		tflog.Debug(ctx, "Packing list of objects for field IpAddress")
		packed, d := packSdwanSaasQualityProfilesMonitorModeStaticIpIpAddressInnerListFromSdk(ctx, sdk.IpAddress)
		diags.Append(d...)
		model.IpAddress = packed
	} else {
		model.IpAddress = basetypes.NewListNull(models.SdwanSaasQualityProfilesMonitorModeStaticIpIpAddressInner{}.AttrType())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.SdwanSaasQualityProfilesMonitorModeStaticIp{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.SdwanSaasQualityProfilesMonitorModeStaticIp", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for SdwanSaasQualityProfilesMonitorModeStaticIp ---
func unpackSdwanSaasQualityProfilesMonitorModeStaticIpListToSdk(ctx context.Context, list types.List) ([]network_services.SdwanSaasQualityProfilesMonitorModeStaticIp, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.SdwanSaasQualityProfilesMonitorModeStaticIp")
	diags := diag.Diagnostics{}
	var data []models.SdwanSaasQualityProfilesMonitorModeStaticIp
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.SdwanSaasQualityProfilesMonitorModeStaticIp, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.SdwanSaasQualityProfilesMonitorModeStaticIp{}.AttrTypes(), &item)
		unpacked, d := unpackSdwanSaasQualityProfilesMonitorModeStaticIpToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.SdwanSaasQualityProfilesMonitorModeStaticIp", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for SdwanSaasQualityProfilesMonitorModeStaticIp ---
func packSdwanSaasQualityProfilesMonitorModeStaticIpListFromSdk(ctx context.Context, sdks []network_services.SdwanSaasQualityProfilesMonitorModeStaticIp) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.SdwanSaasQualityProfilesMonitorModeStaticIp")
	diags := diag.Diagnostics{}
	var data []models.SdwanSaasQualityProfilesMonitorModeStaticIp

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.SdwanSaasQualityProfilesMonitorModeStaticIp
		obj, d := packSdwanSaasQualityProfilesMonitorModeStaticIpFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.SdwanSaasQualityProfilesMonitorModeStaticIp{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.SdwanSaasQualityProfilesMonitorModeStaticIp", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.SdwanSaasQualityProfilesMonitorModeStaticIp{}.AttrType(), data)
}

// --- Unpacker for SdwanSaasQualityProfilesMonitorModeStaticIpFqdn ---
func unpackSdwanSaasQualityProfilesMonitorModeStaticIpFqdnToSdk(ctx context.Context, obj types.Object) (*network_services.SdwanSaasQualityProfilesMonitorModeStaticIpFqdn, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.SdwanSaasQualityProfilesMonitorModeStaticIpFqdn", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.SdwanSaasQualityProfilesMonitorModeStaticIpFqdn
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.SdwanSaasQualityProfilesMonitorModeStaticIpFqdn
	var d diag.Diagnostics
	// Handling Primitives
	if !model.FqdnName.IsNull() && !model.FqdnName.IsUnknown() {
		sdk.FqdnName = model.FqdnName.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "FqdnName", "value": sdk.FqdnName})
	}

	// Handling Primitives
	if !model.ProbeInterval.IsNull() && !model.ProbeInterval.IsUnknown() {
		sdk.ProbeInterval = int32(model.ProbeInterval.ValueInt64())
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "ProbeInterval", "value": sdk.ProbeInterval})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.SdwanSaasQualityProfilesMonitorModeStaticIpFqdn", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for SdwanSaasQualityProfilesMonitorModeStaticIpFqdn ---
func packSdwanSaasQualityProfilesMonitorModeStaticIpFqdnFromSdk(ctx context.Context, sdk network_services.SdwanSaasQualityProfilesMonitorModeStaticIpFqdn) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.SdwanSaasQualityProfilesMonitorModeStaticIpFqdn", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.SdwanSaasQualityProfilesMonitorModeStaticIpFqdn
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.FqdnName = basetypes.NewStringValue(sdk.FqdnName)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "FqdnName", "value": sdk.FqdnName})
	// Handling Primitives
	// Standard primitive packing
	model.ProbeInterval = basetypes.NewInt64Value(int64(sdk.ProbeInterval))
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.SdwanSaasQualityProfilesMonitorModeStaticIpFqdn{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.SdwanSaasQualityProfilesMonitorModeStaticIpFqdn", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for SdwanSaasQualityProfilesMonitorModeStaticIpFqdn ---
func unpackSdwanSaasQualityProfilesMonitorModeStaticIpFqdnListToSdk(ctx context.Context, list types.List) ([]network_services.SdwanSaasQualityProfilesMonitorModeStaticIpFqdn, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.SdwanSaasQualityProfilesMonitorModeStaticIpFqdn")
	diags := diag.Diagnostics{}
	var data []models.SdwanSaasQualityProfilesMonitorModeStaticIpFqdn
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.SdwanSaasQualityProfilesMonitorModeStaticIpFqdn, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.SdwanSaasQualityProfilesMonitorModeStaticIpFqdn{}.AttrTypes(), &item)
		unpacked, d := unpackSdwanSaasQualityProfilesMonitorModeStaticIpFqdnToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.SdwanSaasQualityProfilesMonitorModeStaticIpFqdn", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for SdwanSaasQualityProfilesMonitorModeStaticIpFqdn ---
func packSdwanSaasQualityProfilesMonitorModeStaticIpFqdnListFromSdk(ctx context.Context, sdks []network_services.SdwanSaasQualityProfilesMonitorModeStaticIpFqdn) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.SdwanSaasQualityProfilesMonitorModeStaticIpFqdn")
	diags := diag.Diagnostics{}
	var data []models.SdwanSaasQualityProfilesMonitorModeStaticIpFqdn

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.SdwanSaasQualityProfilesMonitorModeStaticIpFqdn
		obj, d := packSdwanSaasQualityProfilesMonitorModeStaticIpFqdnFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.SdwanSaasQualityProfilesMonitorModeStaticIpFqdn{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.SdwanSaasQualityProfilesMonitorModeStaticIpFqdn", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.SdwanSaasQualityProfilesMonitorModeStaticIpFqdn{}.AttrType(), data)
}

// --- Unpacker for SdwanSaasQualityProfilesMonitorModeStaticIpIpAddressInner ---
func unpackSdwanSaasQualityProfilesMonitorModeStaticIpIpAddressInnerToSdk(ctx context.Context, obj types.Object) (*network_services.SdwanSaasQualityProfilesMonitorModeStaticIpIpAddressInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.SdwanSaasQualityProfilesMonitorModeStaticIpIpAddressInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.SdwanSaasQualityProfilesMonitorModeStaticIpIpAddressInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.SdwanSaasQualityProfilesMonitorModeStaticIpIpAddressInner
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	// Handling Primitives
	if !model.ProbeInterval.IsNull() && !model.ProbeInterval.IsUnknown() {
		sdk.ProbeInterval = int32(model.ProbeInterval.ValueInt64())
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "ProbeInterval", "value": sdk.ProbeInterval})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.SdwanSaasQualityProfilesMonitorModeStaticIpIpAddressInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for SdwanSaasQualityProfilesMonitorModeStaticIpIpAddressInner ---
func packSdwanSaasQualityProfilesMonitorModeStaticIpIpAddressInnerFromSdk(ctx context.Context, sdk network_services.SdwanSaasQualityProfilesMonitorModeStaticIpIpAddressInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.SdwanSaasQualityProfilesMonitorModeStaticIpIpAddressInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.SdwanSaasQualityProfilesMonitorModeStaticIpIpAddressInner
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	// Handling Primitives
	// Standard primitive packing
	model.ProbeInterval = basetypes.NewInt64Value(int64(sdk.ProbeInterval))
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.SdwanSaasQualityProfilesMonitorModeStaticIpIpAddressInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.SdwanSaasQualityProfilesMonitorModeStaticIpIpAddressInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for SdwanSaasQualityProfilesMonitorModeStaticIpIpAddressInner ---
func unpackSdwanSaasQualityProfilesMonitorModeStaticIpIpAddressInnerListToSdk(ctx context.Context, list types.List) ([]network_services.SdwanSaasQualityProfilesMonitorModeStaticIpIpAddressInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.SdwanSaasQualityProfilesMonitorModeStaticIpIpAddressInner")
	diags := diag.Diagnostics{}
	var data []models.SdwanSaasQualityProfilesMonitorModeStaticIpIpAddressInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.SdwanSaasQualityProfilesMonitorModeStaticIpIpAddressInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.SdwanSaasQualityProfilesMonitorModeStaticIpIpAddressInner{}.AttrTypes(), &item)
		unpacked, d := unpackSdwanSaasQualityProfilesMonitorModeStaticIpIpAddressInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.SdwanSaasQualityProfilesMonitorModeStaticIpIpAddressInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for SdwanSaasQualityProfilesMonitorModeStaticIpIpAddressInner ---
func packSdwanSaasQualityProfilesMonitorModeStaticIpIpAddressInnerListFromSdk(ctx context.Context, sdks []network_services.SdwanSaasQualityProfilesMonitorModeStaticIpIpAddressInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.SdwanSaasQualityProfilesMonitorModeStaticIpIpAddressInner")
	diags := diag.Diagnostics{}
	var data []models.SdwanSaasQualityProfilesMonitorModeStaticIpIpAddressInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.SdwanSaasQualityProfilesMonitorModeStaticIpIpAddressInner
		obj, d := packSdwanSaasQualityProfilesMonitorModeStaticIpIpAddressInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.SdwanSaasQualityProfilesMonitorModeStaticIpIpAddressInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.SdwanSaasQualityProfilesMonitorModeStaticIpIpAddressInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.SdwanSaasQualityProfilesMonitorModeStaticIpIpAddressInner{}.AttrType(), data)
}
