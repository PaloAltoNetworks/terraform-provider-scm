package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/objects"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/objects"
)

// --- Unpacker for SyslogServerProfiles ---
func unpackSyslogServerProfilesToSdk(ctx context.Context, obj types.Object) (*objects.SyslogServerProfiles, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.SyslogServerProfiles", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.SyslogServerProfiles
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.SyslogServerProfiles
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

	// Handling Objects
	if !model.Format.IsNull() && !model.Format.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Format")
		unpacked, d := unpackSyslogServerProfilesFormatToSdk(ctx, model.Format)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Format"})
		}
		if unpacked != nil {
			sdk.Format = unpacked
		}
	}

	// Handling Primitives
	if !model.Id.IsNull() && !model.Id.IsUnknown() {
		sdk.Id = model.Id.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Id", "value": sdk.Id})
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	// Handling Lists
	if !model.Server.IsNull() && !model.Server.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field Server")
		unpacked, d := unpackSyslogServerProfilesServerInnerListToSdk(ctx, model.Server)
		diags.Append(d...)
		sdk.Server = unpacked
	}

	// Handling Primitives
	if !model.Snippet.IsNull() && !model.Snippet.IsUnknown() {
		sdk.Snippet = model.Snippet.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.SyslogServerProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for SyslogServerProfiles ---
func packSyslogServerProfilesFromSdk(ctx context.Context, sdk objects.SyslogServerProfiles) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.SyslogServerProfiles", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.SyslogServerProfiles
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
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Format != nil {
		tflog.Debug(ctx, "Packing nested object for field Format")
		packed, d := packSyslogServerProfilesFormatFromSdk(ctx, *sdk.Format)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Format"})
		}
		model.Format = packed
	} else {
		model.Format = basetypes.NewObjectNull(models.SyslogServerProfilesFormat{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	model.Id = basetypes.NewStringValue(sdk.Id)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Id", "value": sdk.Id})
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	// Handling Lists
	if sdk.Server != nil {
		tflog.Debug(ctx, "Packing list of objects for field Server")
		packed, d := packSyslogServerProfilesServerInnerListFromSdk(ctx, sdk.Server)
		diags.Append(d...)
		model.Server = packed
	} else {
		model.Server = basetypes.NewListNull(models.SyslogServerProfilesServerInner{}.AttrType())
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

	obj, d := types.ObjectValueFrom(ctx, models.SyslogServerProfiles{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.SyslogServerProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for SyslogServerProfiles ---
func unpackSyslogServerProfilesListToSdk(ctx context.Context, list types.List) ([]objects.SyslogServerProfiles, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.SyslogServerProfiles")
	diags := diag.Diagnostics{}
	var data []models.SyslogServerProfiles
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.SyslogServerProfiles, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.SyslogServerProfiles{}.AttrTypes(), &item)
		unpacked, d := unpackSyslogServerProfilesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.SyslogServerProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for SyslogServerProfiles ---
func packSyslogServerProfilesListFromSdk(ctx context.Context, sdks []objects.SyslogServerProfiles) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.SyslogServerProfiles")
	diags := diag.Diagnostics{}
	var data []models.SyslogServerProfiles

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.SyslogServerProfiles
		obj, d := packSyslogServerProfilesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.SyslogServerProfiles{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.SyslogServerProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.SyslogServerProfiles{}.AttrType(), data)
}

// --- Unpacker for SyslogServerProfilesFormat ---
func unpackSyslogServerProfilesFormatToSdk(ctx context.Context, obj types.Object) (*objects.SyslogServerProfilesFormat, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.SyslogServerProfilesFormat", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.SyslogServerProfilesFormat
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.SyslogServerProfilesFormat
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Auth.IsNull() && !model.Auth.IsUnknown() {
		sdk.Auth = model.Auth.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Auth", "value": *sdk.Auth})
	}

	// Handling Primitives
	if !model.Config.IsNull() && !model.Config.IsUnknown() {
		sdk.Config = model.Config.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Config", "value": *sdk.Config})
	}

	// Handling Primitives
	if !model.Correlation.IsNull() && !model.Correlation.IsUnknown() {
		sdk.Correlation = model.Correlation.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Correlation", "value": *sdk.Correlation})
	}

	// Handling Primitives
	if !model.Data.IsNull() && !model.Data.IsUnknown() {
		sdk.Data = model.Data.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Data", "value": *sdk.Data})
	}

	// Handling Primitives
	if !model.Decryption.IsNull() && !model.Decryption.IsUnknown() {
		sdk.Decryption = model.Decryption.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Decryption", "value": *sdk.Decryption})
	}

	// Handling Objects
	if !model.Escaping.IsNull() && !model.Escaping.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Escaping")
		unpacked, d := unpackSyslogServerProfilesFormatEscapingToSdk(ctx, model.Escaping)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Escaping"})
		}
		if unpacked != nil {
			sdk.Escaping = unpacked
		}
	}

	// Handling Primitives
	if !model.Globalprotect.IsNull() && !model.Globalprotect.IsUnknown() {
		sdk.Globalprotect = model.Globalprotect.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Globalprotect", "value": *sdk.Globalprotect})
	}

	// Handling Primitives
	if !model.Gtp.IsNull() && !model.Gtp.IsUnknown() {
		sdk.Gtp = model.Gtp.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Gtp", "value": *sdk.Gtp})
	}

	// Handling Primitives
	if !model.HipMatch.IsNull() && !model.HipMatch.IsUnknown() {
		sdk.HipMatch = model.HipMatch.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "HipMatch", "value": *sdk.HipMatch})
	}

	// Handling Primitives
	if !model.Iptag.IsNull() && !model.Iptag.IsUnknown() {
		sdk.Iptag = model.Iptag.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Iptag", "value": *sdk.Iptag})
	}

	// Handling Primitives
	if !model.Sctp.IsNull() && !model.Sctp.IsUnknown() {
		sdk.Sctp = model.Sctp.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Sctp", "value": *sdk.Sctp})
	}

	// Handling Primitives
	if !model.System.IsNull() && !model.System.IsUnknown() {
		sdk.System = model.System.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "System", "value": *sdk.System})
	}

	// Handling Primitives
	if !model.Threat.IsNull() && !model.Threat.IsUnknown() {
		sdk.Threat = model.Threat.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Threat", "value": *sdk.Threat})
	}

	// Handling Primitives
	if !model.Traffic.IsNull() && !model.Traffic.IsUnknown() {
		sdk.Traffic = model.Traffic.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Traffic", "value": *sdk.Traffic})
	}

	// Handling Primitives
	if !model.Tunnel.IsNull() && !model.Tunnel.IsUnknown() {
		sdk.Tunnel = model.Tunnel.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Tunnel", "value": *sdk.Tunnel})
	}

	// Handling Primitives
	if !model.Url.IsNull() && !model.Url.IsUnknown() {
		sdk.Url = model.Url.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Url", "value": *sdk.Url})
	}

	// Handling Primitives
	if !model.Userid.IsNull() && !model.Userid.IsUnknown() {
		sdk.Userid = model.Userid.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Userid", "value": *sdk.Userid})
	}

	// Handling Primitives
	if !model.Wildfire.IsNull() && !model.Wildfire.IsUnknown() {
		sdk.Wildfire = model.Wildfire.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Wildfire", "value": *sdk.Wildfire})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.SyslogServerProfilesFormat", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for SyslogServerProfilesFormat ---
func packSyslogServerProfilesFormatFromSdk(ctx context.Context, sdk objects.SyslogServerProfilesFormat) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.SyslogServerProfilesFormat", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.SyslogServerProfilesFormat
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Auth != nil {
		model.Auth = basetypes.NewStringValue(*sdk.Auth)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Auth", "value": *sdk.Auth})
	} else {
		model.Auth = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Config != nil {
		model.Config = basetypes.NewStringValue(*sdk.Config)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Config", "value": *sdk.Config})
	} else {
		model.Config = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Correlation != nil {
		model.Correlation = basetypes.NewStringValue(*sdk.Correlation)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Correlation", "value": *sdk.Correlation})
	} else {
		model.Correlation = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Data != nil {
		model.Data = basetypes.NewStringValue(*sdk.Data)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Data", "value": *sdk.Data})
	} else {
		model.Data = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Decryption != nil {
		model.Decryption = basetypes.NewStringValue(*sdk.Decryption)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Decryption", "value": *sdk.Decryption})
	} else {
		model.Decryption = basetypes.NewStringNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Escaping != nil {
		tflog.Debug(ctx, "Packing nested object for field Escaping")
		packed, d := packSyslogServerProfilesFormatEscapingFromSdk(ctx, *sdk.Escaping)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Escaping"})
		}
		model.Escaping = packed
	} else {
		model.Escaping = basetypes.NewObjectNull(models.SyslogServerProfilesFormatEscaping{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Globalprotect != nil {
		model.Globalprotect = basetypes.NewStringValue(*sdk.Globalprotect)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Globalprotect", "value": *sdk.Globalprotect})
	} else {
		model.Globalprotect = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Gtp != nil {
		model.Gtp = basetypes.NewStringValue(*sdk.Gtp)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Gtp", "value": *sdk.Gtp})
	} else {
		model.Gtp = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.HipMatch != nil {
		model.HipMatch = basetypes.NewStringValue(*sdk.HipMatch)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "HipMatch", "value": *sdk.HipMatch})
	} else {
		model.HipMatch = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Iptag != nil {
		model.Iptag = basetypes.NewStringValue(*sdk.Iptag)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Iptag", "value": *sdk.Iptag})
	} else {
		model.Iptag = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Sctp != nil {
		model.Sctp = basetypes.NewStringValue(*sdk.Sctp)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Sctp", "value": *sdk.Sctp})
	} else {
		model.Sctp = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.System != nil {
		model.System = basetypes.NewStringValue(*sdk.System)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "System", "value": *sdk.System})
	} else {
		model.System = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Threat != nil {
		model.Threat = basetypes.NewStringValue(*sdk.Threat)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Threat", "value": *sdk.Threat})
	} else {
		model.Threat = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Traffic != nil {
		model.Traffic = basetypes.NewStringValue(*sdk.Traffic)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Traffic", "value": *sdk.Traffic})
	} else {
		model.Traffic = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Tunnel != nil {
		model.Tunnel = basetypes.NewStringValue(*sdk.Tunnel)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Tunnel", "value": *sdk.Tunnel})
	} else {
		model.Tunnel = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Url != nil {
		model.Url = basetypes.NewStringValue(*sdk.Url)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Url", "value": *sdk.Url})
	} else {
		model.Url = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Userid != nil {
		model.Userid = basetypes.NewStringValue(*sdk.Userid)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Userid", "value": *sdk.Userid})
	} else {
		model.Userid = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Wildfire != nil {
		model.Wildfire = basetypes.NewStringValue(*sdk.Wildfire)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Wildfire", "value": *sdk.Wildfire})
	} else {
		model.Wildfire = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.SyslogServerProfilesFormat{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.SyslogServerProfilesFormat", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for SyslogServerProfilesFormat ---
func unpackSyslogServerProfilesFormatListToSdk(ctx context.Context, list types.List) ([]objects.SyslogServerProfilesFormat, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.SyslogServerProfilesFormat")
	diags := diag.Diagnostics{}
	var data []models.SyslogServerProfilesFormat
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.SyslogServerProfilesFormat, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.SyslogServerProfilesFormat{}.AttrTypes(), &item)
		unpacked, d := unpackSyslogServerProfilesFormatToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.SyslogServerProfilesFormat", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for SyslogServerProfilesFormat ---
func packSyslogServerProfilesFormatListFromSdk(ctx context.Context, sdks []objects.SyslogServerProfilesFormat) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.SyslogServerProfilesFormat")
	diags := diag.Diagnostics{}
	var data []models.SyslogServerProfilesFormat

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.SyslogServerProfilesFormat
		obj, d := packSyslogServerProfilesFormatFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.SyslogServerProfilesFormat{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.SyslogServerProfilesFormat", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.SyslogServerProfilesFormat{}.AttrType(), data)
}

// --- Unpacker for SyslogServerProfilesFormatEscaping ---
func unpackSyslogServerProfilesFormatEscapingToSdk(ctx context.Context, obj types.Object) (*objects.SyslogServerProfilesFormatEscaping, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.SyslogServerProfilesFormatEscaping", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.SyslogServerProfilesFormatEscaping
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.SyslogServerProfilesFormatEscaping
	var d diag.Diagnostics
	// Handling Primitives
	if !model.EscapeCharacter.IsNull() && !model.EscapeCharacter.IsUnknown() {
		sdk.EscapeCharacter = model.EscapeCharacter.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "EscapeCharacter", "value": *sdk.EscapeCharacter})
	}

	// Handling Primitives
	if !model.EscapedCharacters.IsNull() && !model.EscapedCharacters.IsUnknown() {
		sdk.EscapedCharacters = model.EscapedCharacters.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "EscapedCharacters", "value": *sdk.EscapedCharacters})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.SyslogServerProfilesFormatEscaping", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for SyslogServerProfilesFormatEscaping ---
func packSyslogServerProfilesFormatEscapingFromSdk(ctx context.Context, sdk objects.SyslogServerProfilesFormatEscaping) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.SyslogServerProfilesFormatEscaping", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.SyslogServerProfilesFormatEscaping
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.EscapeCharacter != nil {
		model.EscapeCharacter = basetypes.NewStringValue(*sdk.EscapeCharacter)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "EscapeCharacter", "value": *sdk.EscapeCharacter})
	} else {
		model.EscapeCharacter = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.EscapedCharacters != nil {
		model.EscapedCharacters = basetypes.NewStringValue(*sdk.EscapedCharacters)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "EscapedCharacters", "value": *sdk.EscapedCharacters})
	} else {
		model.EscapedCharacters = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.SyslogServerProfilesFormatEscaping{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.SyslogServerProfilesFormatEscaping", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for SyslogServerProfilesFormatEscaping ---
func unpackSyslogServerProfilesFormatEscapingListToSdk(ctx context.Context, list types.List) ([]objects.SyslogServerProfilesFormatEscaping, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.SyslogServerProfilesFormatEscaping")
	diags := diag.Diagnostics{}
	var data []models.SyslogServerProfilesFormatEscaping
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.SyslogServerProfilesFormatEscaping, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.SyslogServerProfilesFormatEscaping{}.AttrTypes(), &item)
		unpacked, d := unpackSyslogServerProfilesFormatEscapingToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.SyslogServerProfilesFormatEscaping", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for SyslogServerProfilesFormatEscaping ---
func packSyslogServerProfilesFormatEscapingListFromSdk(ctx context.Context, sdks []objects.SyslogServerProfilesFormatEscaping) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.SyslogServerProfilesFormatEscaping")
	diags := diag.Diagnostics{}
	var data []models.SyslogServerProfilesFormatEscaping

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.SyslogServerProfilesFormatEscaping
		obj, d := packSyslogServerProfilesFormatEscapingFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.SyslogServerProfilesFormatEscaping{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.SyslogServerProfilesFormatEscaping", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.SyslogServerProfilesFormatEscaping{}.AttrType(), data)
}

// --- Unpacker for SyslogServerProfilesServerInner ---
func unpackSyslogServerProfilesServerInnerToSdk(ctx context.Context, obj types.Object) (*objects.SyslogServerProfilesServerInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.SyslogServerProfilesServerInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.SyslogServerProfilesServerInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.SyslogServerProfilesServerInner
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Facility.IsNull() && !model.Facility.IsUnknown() {
		sdk.Facility = model.Facility.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Facility", "value": *sdk.Facility})
	}

	// Handling Primitives
	if !model.Format.IsNull() && !model.Format.IsUnknown() {
		sdk.Format = model.Format.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Format", "value": *sdk.Format})
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
	}

	// Handling Primitives
	if !model.Port.IsNull() && !model.Port.IsUnknown() {
		val := int32(model.Port.ValueInt64())
		sdk.Port = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Port", "value": *sdk.Port})
	}

	// Handling Primitives
	if !model.Server.IsNull() && !model.Server.IsUnknown() {
		sdk.Server = model.Server.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Server", "value": *sdk.Server})
	}

	// Handling Primitives
	if !model.Transport.IsNull() && !model.Transport.IsUnknown() {
		sdk.Transport = model.Transport.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Transport", "value": *sdk.Transport})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.SyslogServerProfilesServerInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for SyslogServerProfilesServerInner ---
func packSyslogServerProfilesServerInnerFromSdk(ctx context.Context, sdk objects.SyslogServerProfilesServerInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.SyslogServerProfilesServerInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.SyslogServerProfilesServerInner
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Facility != nil {
		model.Facility = basetypes.NewStringValue(*sdk.Facility)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Facility", "value": *sdk.Facility})
	} else {
		model.Facility = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Format != nil {
		model.Format = basetypes.NewStringValue(*sdk.Format)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Format", "value": *sdk.Format})
	} else {
		model.Format = basetypes.NewStringNull()
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
	if sdk.Port != nil {
		model.Port = basetypes.NewInt64Value(int64(*sdk.Port))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Port", "value": *sdk.Port})
	} else {
		model.Port = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Server != nil {
		model.Server = basetypes.NewStringValue(*sdk.Server)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Server", "value": *sdk.Server})
	} else {
		model.Server = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Transport != nil {
		model.Transport = basetypes.NewStringValue(*sdk.Transport)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Transport", "value": *sdk.Transport})
	} else {
		model.Transport = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.SyslogServerProfilesServerInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.SyslogServerProfilesServerInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for SyslogServerProfilesServerInner ---
func unpackSyslogServerProfilesServerInnerListToSdk(ctx context.Context, list types.List) ([]objects.SyslogServerProfilesServerInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.SyslogServerProfilesServerInner")
	diags := diag.Diagnostics{}
	var data []models.SyslogServerProfilesServerInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.SyslogServerProfilesServerInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.SyslogServerProfilesServerInner{}.AttrTypes(), &item)
		unpacked, d := unpackSyslogServerProfilesServerInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.SyslogServerProfilesServerInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for SyslogServerProfilesServerInner ---
func packSyslogServerProfilesServerInnerListFromSdk(ctx context.Context, sdks []objects.SyslogServerProfilesServerInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.SyslogServerProfilesServerInner")
	diags := diag.Diagnostics{}
	var data []models.SyslogServerProfilesServerInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.SyslogServerProfilesServerInner
		obj, d := packSyslogServerProfilesServerInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.SyslogServerProfilesServerInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.SyslogServerProfilesServerInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.SyslogServerProfilesServerInner{}.AttrType(), data)
}
