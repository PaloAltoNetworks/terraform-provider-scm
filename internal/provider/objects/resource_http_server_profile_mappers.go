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

// --- Unpacker for HttpServerProfiles ---
func unpackHttpServerProfilesToSdk(ctx context.Context, obj types.Object) (*objects.HttpServerProfiles, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.HttpServerProfiles", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.HttpServerProfiles
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.HttpServerProfiles
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
		unpacked, d := unpackHttpServerProfilesFormatToSdk(ctx, model.Format)
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
		unpacked, d := unpackHttpServerProfilesServerInnerListToSdk(ctx, model.Server)
		diags.Append(d...)
		sdk.Server = unpacked
	}

	// Handling Primitives
	if !model.Snippet.IsNull() && !model.Snippet.IsUnknown() {
		sdk.Snippet = model.Snippet.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	}

	// Handling Primitives
	if !model.TagRegistration.IsNull() && !model.TagRegistration.IsUnknown() {
		sdk.TagRegistration = model.TagRegistration.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "TagRegistration", "value": *sdk.TagRegistration})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.HttpServerProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for HttpServerProfiles ---
func packHttpServerProfilesFromSdk(ctx context.Context, sdk objects.HttpServerProfiles) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.HttpServerProfiles", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.HttpServerProfiles
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
		packed, d := packHttpServerProfilesFormatFromSdk(ctx, *sdk.Format)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Format"})
		}
		model.Format = packed
	} else {
		model.Format = basetypes.NewObjectNull(models.HttpServerProfilesFormat{}.AttrTypes())
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
		packed, d := packHttpServerProfilesServerInnerListFromSdk(ctx, sdk.Server)
		diags.Append(d...)
		model.Server = packed
	} else {
		model.Server = basetypes.NewListNull(models.HttpServerProfilesServerInner{}.AttrType())
	}
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
	if sdk.TagRegistration != nil {
		model.TagRegistration = basetypes.NewBoolValue(*sdk.TagRegistration)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "TagRegistration", "value": *sdk.TagRegistration})
	} else {
		model.TagRegistration = basetypes.NewBoolNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.HttpServerProfiles{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.HttpServerProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for HttpServerProfiles ---
func unpackHttpServerProfilesListToSdk(ctx context.Context, list types.List) ([]objects.HttpServerProfiles, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.HttpServerProfiles")
	diags := diag.Diagnostics{}
	var data []models.HttpServerProfiles
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.HttpServerProfiles, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.HttpServerProfiles{}.AttrTypes(), &item)
		unpacked, d := unpackHttpServerProfilesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.HttpServerProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for HttpServerProfiles ---
func packHttpServerProfilesListFromSdk(ctx context.Context, sdks []objects.HttpServerProfiles) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.HttpServerProfiles")
	diags := diag.Diagnostics{}
	var data []models.HttpServerProfiles

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.HttpServerProfiles
		obj, d := packHttpServerProfilesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.HttpServerProfiles{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.HttpServerProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.HttpServerProfiles{}.AttrType(), data)
}

// --- Unpacker for HttpServerProfilesFormat ---
func unpackHttpServerProfilesFormatToSdk(ctx context.Context, obj types.Object) (*objects.HttpServerProfilesFormat, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.HttpServerProfilesFormat", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.HttpServerProfilesFormat
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.HttpServerProfilesFormat
	var d diag.Diagnostics
	// Handling Objects
	if !model.Auth.IsNull() && !model.Auth.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Auth")
		unpacked, d := unpackPayloadFormatToSdk(ctx, model.Auth)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Auth"})
		}
		if unpacked != nil {
			sdk.Auth = unpacked
		}
	}

	// Handling Objects
	if !model.Config.IsNull() && !model.Config.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Config")
		unpacked, d := unpackPayloadFormatToSdk(ctx, model.Config)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Config"})
		}
		if unpacked != nil {
			sdk.Config = unpacked
		}
	}

	// Handling Objects
	if !model.Correlation.IsNull() && !model.Correlation.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Correlation")
		unpacked, d := unpackPayloadFormatToSdk(ctx, model.Correlation)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Correlation"})
		}
		if unpacked != nil {
			sdk.Correlation = unpacked
		}
	}

	// Handling Objects
	if !model.Data.IsNull() && !model.Data.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Data")
		unpacked, d := unpackPayloadFormatToSdk(ctx, model.Data)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Data"})
		}
		if unpacked != nil {
			sdk.Data = unpacked
		}
	}

	// Handling Objects
	if !model.Decryption.IsNull() && !model.Decryption.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Decryption")
		unpacked, d := unpackPayloadFormatToSdk(ctx, model.Decryption)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Decryption"})
		}
		if unpacked != nil {
			sdk.Decryption = unpacked
		}
	}

	// Handling Objects
	if !model.Globalprotect.IsNull() && !model.Globalprotect.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Globalprotect")
		unpacked, d := unpackPayloadFormatToSdk(ctx, model.Globalprotect)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Globalprotect"})
		}
		if unpacked != nil {
			sdk.Globalprotect = unpacked
		}
	}

	// Handling Objects
	if !model.Gtp.IsNull() && !model.Gtp.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Gtp")
		unpacked, d := unpackPayloadFormatToSdk(ctx, model.Gtp)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Gtp"})
		}
		if unpacked != nil {
			sdk.Gtp = unpacked
		}
	}

	// Handling Objects
	if !model.HipMatch.IsNull() && !model.HipMatch.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field HipMatch")
		unpacked, d := unpackPayloadFormatToSdk(ctx, model.HipMatch)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "HipMatch"})
		}
		if unpacked != nil {
			sdk.HipMatch = unpacked
		}
	}

	// Handling Objects
	if !model.Iptag.IsNull() && !model.Iptag.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Iptag")
		unpacked, d := unpackPayloadFormatToSdk(ctx, model.Iptag)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Iptag"})
		}
		if unpacked != nil {
			sdk.Iptag = unpacked
		}
	}

	// Handling Objects
	if !model.Sctp.IsNull() && !model.Sctp.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Sctp")
		unpacked, d := unpackPayloadFormatToSdk(ctx, model.Sctp)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Sctp"})
		}
		if unpacked != nil {
			sdk.Sctp = unpacked
		}
	}

	// Handling Objects
	if !model.System.IsNull() && !model.System.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field System")
		unpacked, d := unpackPayloadFormatToSdk(ctx, model.System)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "System"})
		}
		if unpacked != nil {
			sdk.System = unpacked
		}
	}

	// Handling Objects
	if !model.Threat.IsNull() && !model.Threat.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Threat")
		unpacked, d := unpackPayloadFormatToSdk(ctx, model.Threat)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Threat"})
		}
		if unpacked != nil {
			sdk.Threat = unpacked
		}
	}

	// Handling Objects
	if !model.Traffic.IsNull() && !model.Traffic.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Traffic")
		unpacked, d := unpackPayloadFormatToSdk(ctx, model.Traffic)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Traffic"})
		}
		if unpacked != nil {
			sdk.Traffic = unpacked
		}
	}

	// Handling Objects
	if !model.Tunnel.IsNull() && !model.Tunnel.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Tunnel")
		unpacked, d := unpackPayloadFormatToSdk(ctx, model.Tunnel)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Tunnel"})
		}
		if unpacked != nil {
			sdk.Tunnel = unpacked
		}
	}

	// Handling Objects
	if !model.Url.IsNull() && !model.Url.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Url")
		unpacked, d := unpackPayloadFormatToSdk(ctx, model.Url)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Url"})
		}
		if unpacked != nil {
			sdk.Url = unpacked
		}
	}

	// Handling Objects
	if !model.Userid.IsNull() && !model.Userid.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Userid")
		unpacked, d := unpackPayloadFormatToSdk(ctx, model.Userid)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Userid"})
		}
		if unpacked != nil {
			sdk.Userid = unpacked
		}
	}

	// Handling Objects
	if !model.Wildfire.IsNull() && !model.Wildfire.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Wildfire")
		unpacked, d := unpackPayloadFormatToSdk(ctx, model.Wildfire)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Wildfire"})
		}
		if unpacked != nil {
			sdk.Wildfire = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.HttpServerProfilesFormat", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for HttpServerProfilesFormat ---
func packHttpServerProfilesFormatFromSdk(ctx context.Context, sdk objects.HttpServerProfilesFormat) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.HttpServerProfilesFormat", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.HttpServerProfilesFormat
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Auth != nil {
		tflog.Debug(ctx, "Packing nested object for field Auth")
		packed, d := packPayloadFormatFromSdk(ctx, *sdk.Auth)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Auth"})
		}
		model.Auth = packed
	} else {
		model.Auth = basetypes.NewObjectNull(models.PayloadFormat{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Config != nil {
		tflog.Debug(ctx, "Packing nested object for field Config")
		packed, d := packPayloadFormatFromSdk(ctx, *sdk.Config)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Config"})
		}
		model.Config = packed
	} else {
		model.Config = basetypes.NewObjectNull(models.PayloadFormat{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Correlation != nil {
		tflog.Debug(ctx, "Packing nested object for field Correlation")
		packed, d := packPayloadFormatFromSdk(ctx, *sdk.Correlation)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Correlation"})
		}
		model.Correlation = packed
	} else {
		model.Correlation = basetypes.NewObjectNull(models.PayloadFormat{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Data != nil {
		tflog.Debug(ctx, "Packing nested object for field Data")
		packed, d := packPayloadFormatFromSdk(ctx, *sdk.Data)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Data"})
		}
		model.Data = packed
	} else {
		model.Data = basetypes.NewObjectNull(models.PayloadFormat{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Decryption != nil {
		tflog.Debug(ctx, "Packing nested object for field Decryption")
		packed, d := packPayloadFormatFromSdk(ctx, *sdk.Decryption)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Decryption"})
		}
		model.Decryption = packed
	} else {
		model.Decryption = basetypes.NewObjectNull(models.PayloadFormat{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Globalprotect != nil {
		tflog.Debug(ctx, "Packing nested object for field Globalprotect")
		packed, d := packPayloadFormatFromSdk(ctx, *sdk.Globalprotect)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Globalprotect"})
		}
		model.Globalprotect = packed
	} else {
		model.Globalprotect = basetypes.NewObjectNull(models.PayloadFormat{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Gtp != nil {
		tflog.Debug(ctx, "Packing nested object for field Gtp")
		packed, d := packPayloadFormatFromSdk(ctx, *sdk.Gtp)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Gtp"})
		}
		model.Gtp = packed
	} else {
		model.Gtp = basetypes.NewObjectNull(models.PayloadFormat{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.HipMatch != nil {
		tflog.Debug(ctx, "Packing nested object for field HipMatch")
		packed, d := packPayloadFormatFromSdk(ctx, *sdk.HipMatch)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "HipMatch"})
		}
		model.HipMatch = packed
	} else {
		model.HipMatch = basetypes.NewObjectNull(models.PayloadFormat{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Iptag != nil {
		tflog.Debug(ctx, "Packing nested object for field Iptag")
		packed, d := packPayloadFormatFromSdk(ctx, *sdk.Iptag)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Iptag"})
		}
		model.Iptag = packed
	} else {
		model.Iptag = basetypes.NewObjectNull(models.PayloadFormat{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Sctp != nil {
		tflog.Debug(ctx, "Packing nested object for field Sctp")
		packed, d := packPayloadFormatFromSdk(ctx, *sdk.Sctp)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Sctp"})
		}
		model.Sctp = packed
	} else {
		model.Sctp = basetypes.NewObjectNull(models.PayloadFormat{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.System != nil {
		tflog.Debug(ctx, "Packing nested object for field System")
		packed, d := packPayloadFormatFromSdk(ctx, *sdk.System)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "System"})
		}
		model.System = packed
	} else {
		model.System = basetypes.NewObjectNull(models.PayloadFormat{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Threat != nil {
		tflog.Debug(ctx, "Packing nested object for field Threat")
		packed, d := packPayloadFormatFromSdk(ctx, *sdk.Threat)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Threat"})
		}
		model.Threat = packed
	} else {
		model.Threat = basetypes.NewObjectNull(models.PayloadFormat{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Traffic != nil {
		tflog.Debug(ctx, "Packing nested object for field Traffic")
		packed, d := packPayloadFormatFromSdk(ctx, *sdk.Traffic)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Traffic"})
		}
		model.Traffic = packed
	} else {
		model.Traffic = basetypes.NewObjectNull(models.PayloadFormat{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Tunnel != nil {
		tflog.Debug(ctx, "Packing nested object for field Tunnel")
		packed, d := packPayloadFormatFromSdk(ctx, *sdk.Tunnel)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Tunnel"})
		}
		model.Tunnel = packed
	} else {
		model.Tunnel = basetypes.NewObjectNull(models.PayloadFormat{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Url != nil {
		tflog.Debug(ctx, "Packing nested object for field Url")
		packed, d := packPayloadFormatFromSdk(ctx, *sdk.Url)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Url"})
		}
		model.Url = packed
	} else {
		model.Url = basetypes.NewObjectNull(models.PayloadFormat{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Userid != nil {
		tflog.Debug(ctx, "Packing nested object for field Userid")
		packed, d := packPayloadFormatFromSdk(ctx, *sdk.Userid)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Userid"})
		}
		model.Userid = packed
	} else {
		model.Userid = basetypes.NewObjectNull(models.PayloadFormat{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Wildfire != nil {
		tflog.Debug(ctx, "Packing nested object for field Wildfire")
		packed, d := packPayloadFormatFromSdk(ctx, *sdk.Wildfire)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Wildfire"})
		}
		model.Wildfire = packed
	} else {
		model.Wildfire = basetypes.NewObjectNull(models.PayloadFormat{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.HttpServerProfilesFormat{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.HttpServerProfilesFormat", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for HttpServerProfilesFormat ---
func unpackHttpServerProfilesFormatListToSdk(ctx context.Context, list types.List) ([]objects.HttpServerProfilesFormat, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.HttpServerProfilesFormat")
	diags := diag.Diagnostics{}
	var data []models.HttpServerProfilesFormat
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.HttpServerProfilesFormat, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.HttpServerProfilesFormat{}.AttrTypes(), &item)
		unpacked, d := unpackHttpServerProfilesFormatToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.HttpServerProfilesFormat", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for HttpServerProfilesFormat ---
func packHttpServerProfilesFormatListFromSdk(ctx context.Context, sdks []objects.HttpServerProfilesFormat) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.HttpServerProfilesFormat")
	diags := diag.Diagnostics{}
	var data []models.HttpServerProfilesFormat

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.HttpServerProfilesFormat
		obj, d := packHttpServerProfilesFormatFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.HttpServerProfilesFormat{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.HttpServerProfilesFormat", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.HttpServerProfilesFormat{}.AttrType(), data)
}

// --- Unpacker for PayloadFormat ---
func unpackPayloadFormatToSdk(ctx context.Context, obj types.Object) (*objects.PayloadFormat, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.PayloadFormat", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.PayloadFormat
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.PayloadFormat
	var d diag.Diagnostics
	// Handling Lists
	if !model.Headers.IsNull() && !model.Headers.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field Headers")
		unpacked, d := unpackPayloadFormatHeadersInnerListToSdk(ctx, model.Headers)
		diags.Append(d...)
		sdk.Headers = unpacked
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
	}

	// Handling Lists
	if !model.Params.IsNull() && !model.Params.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field Params")
		unpacked, d := unpackPayloadFormatParamsInnerListToSdk(ctx, model.Params)
		diags.Append(d...)
		sdk.Params = unpacked
	}

	// Handling Primitives
	if !model.Payload.IsNull() && !model.Payload.IsUnknown() {
		sdk.Payload = model.Payload.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Payload", "value": *sdk.Payload})
	}

	// Handling Primitives
	if !model.UrlFormat.IsNull() && !model.UrlFormat.IsUnknown() {
		sdk.UrlFormat = model.UrlFormat.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "UrlFormat", "value": *sdk.UrlFormat})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.PayloadFormat", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for PayloadFormat ---
func packPayloadFormatFromSdk(ctx context.Context, sdk objects.PayloadFormat) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.PayloadFormat", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.PayloadFormat
	var d diag.Diagnostics
	// Handling Lists
	if sdk.Headers != nil {
		tflog.Debug(ctx, "Packing list of objects for field Headers")
		packed, d := packPayloadFormatHeadersInnerListFromSdk(ctx, sdk.Headers)
		diags.Append(d...)
		model.Headers = packed
	} else {
		model.Headers = basetypes.NewListNull(models.PayloadFormatHeadersInner{}.AttrType())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Name != nil {
		model.Name = basetypes.NewStringValue(*sdk.Name)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
	} else {
		model.Name = basetypes.NewStringNull()
	}
	// Handling Lists
	if sdk.Params != nil {
		tflog.Debug(ctx, "Packing list of objects for field Params")
		packed, d := packPayloadFormatParamsInnerListFromSdk(ctx, sdk.Params)
		diags.Append(d...)
		model.Params = packed
	} else {
		model.Params = basetypes.NewListNull(models.PayloadFormatParamsInner{}.AttrType())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Payload != nil {
		model.Payload = basetypes.NewStringValue(*sdk.Payload)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Payload", "value": *sdk.Payload})
	} else {
		model.Payload = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.UrlFormat != nil {
		model.UrlFormat = basetypes.NewStringValue(*sdk.UrlFormat)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "UrlFormat", "value": *sdk.UrlFormat})
	} else {
		model.UrlFormat = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.PayloadFormat{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.PayloadFormat", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for PayloadFormat ---
func unpackPayloadFormatListToSdk(ctx context.Context, list types.List) ([]objects.PayloadFormat, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.PayloadFormat")
	diags := diag.Diagnostics{}
	var data []models.PayloadFormat
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.PayloadFormat, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.PayloadFormat{}.AttrTypes(), &item)
		unpacked, d := unpackPayloadFormatToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.PayloadFormat", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for PayloadFormat ---
func packPayloadFormatListFromSdk(ctx context.Context, sdks []objects.PayloadFormat) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.PayloadFormat")
	diags := diag.Diagnostics{}
	var data []models.PayloadFormat

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.PayloadFormat
		obj, d := packPayloadFormatFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.PayloadFormat{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.PayloadFormat", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.PayloadFormat{}.AttrType(), data)
}

// --- Unpacker for PayloadFormatHeadersInner ---
func unpackPayloadFormatHeadersInnerToSdk(ctx context.Context, obj types.Object) (*objects.PayloadFormatHeadersInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.PayloadFormatHeadersInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.PayloadFormatHeadersInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.PayloadFormatHeadersInner
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
	}

	// Handling Primitives
	if !model.Value.IsNull() && !model.Value.IsUnknown() {
		sdk.Value = model.Value.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Value", "value": *sdk.Value})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.PayloadFormatHeadersInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for PayloadFormatHeadersInner ---
func packPayloadFormatHeadersInnerFromSdk(ctx context.Context, sdk objects.PayloadFormatHeadersInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.PayloadFormatHeadersInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.PayloadFormatHeadersInner
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
	if sdk.Value != nil {
		model.Value = basetypes.NewStringValue(*sdk.Value)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Value", "value": *sdk.Value})
	} else {
		model.Value = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.PayloadFormatHeadersInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.PayloadFormatHeadersInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for PayloadFormatHeadersInner ---
func unpackPayloadFormatHeadersInnerListToSdk(ctx context.Context, list types.List) ([]objects.PayloadFormatHeadersInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.PayloadFormatHeadersInner")
	diags := diag.Diagnostics{}
	var data []models.PayloadFormatHeadersInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.PayloadFormatHeadersInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.PayloadFormatHeadersInner{}.AttrTypes(), &item)
		unpacked, d := unpackPayloadFormatHeadersInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.PayloadFormatHeadersInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for PayloadFormatHeadersInner ---
func packPayloadFormatHeadersInnerListFromSdk(ctx context.Context, sdks []objects.PayloadFormatHeadersInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.PayloadFormatHeadersInner")
	diags := diag.Diagnostics{}
	var data []models.PayloadFormatHeadersInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.PayloadFormatHeadersInner
		obj, d := packPayloadFormatHeadersInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.PayloadFormatHeadersInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.PayloadFormatHeadersInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.PayloadFormatHeadersInner{}.AttrType(), data)
}

// --- Unpacker for PayloadFormatParamsInner ---
func unpackPayloadFormatParamsInnerToSdk(ctx context.Context, obj types.Object) (*objects.PayloadFormatParamsInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.PayloadFormatParamsInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.PayloadFormatParamsInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.PayloadFormatParamsInner
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
	}

	// Handling Primitives
	if !model.Value.IsNull() && !model.Value.IsUnknown() {
		sdk.Value = model.Value.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Value", "value": *sdk.Value})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.PayloadFormatParamsInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for PayloadFormatParamsInner ---
func packPayloadFormatParamsInnerFromSdk(ctx context.Context, sdk objects.PayloadFormatParamsInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.PayloadFormatParamsInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.PayloadFormatParamsInner
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
	if sdk.Value != nil {
		model.Value = basetypes.NewStringValue(*sdk.Value)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Value", "value": *sdk.Value})
	} else {
		model.Value = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.PayloadFormatParamsInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.PayloadFormatParamsInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for PayloadFormatParamsInner ---
func unpackPayloadFormatParamsInnerListToSdk(ctx context.Context, list types.List) ([]objects.PayloadFormatParamsInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.PayloadFormatParamsInner")
	diags := diag.Diagnostics{}
	var data []models.PayloadFormatParamsInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.PayloadFormatParamsInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.PayloadFormatParamsInner{}.AttrTypes(), &item)
		unpacked, d := unpackPayloadFormatParamsInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.PayloadFormatParamsInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for PayloadFormatParamsInner ---
func packPayloadFormatParamsInnerListFromSdk(ctx context.Context, sdks []objects.PayloadFormatParamsInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.PayloadFormatParamsInner")
	diags := diag.Diagnostics{}
	var data []models.PayloadFormatParamsInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.PayloadFormatParamsInner
		obj, d := packPayloadFormatParamsInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.PayloadFormatParamsInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.PayloadFormatParamsInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.PayloadFormatParamsInner{}.AttrType(), data)
}

// --- Unpacker for HttpServerProfilesServerInner ---
func unpackHttpServerProfilesServerInnerToSdk(ctx context.Context, obj types.Object) (*objects.HttpServerProfilesServerInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.HttpServerProfilesServerInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.HttpServerProfilesServerInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.HttpServerProfilesServerInner
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Address.IsNull() && !model.Address.IsUnknown() {
		sdk.Address = model.Address.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Address", "value": *sdk.Address})
	}

	// Handling Primitives
	if !model.CertificateProfile.IsNull() && !model.CertificateProfile.IsUnknown() {
		sdk.CertificateProfile = model.CertificateProfile.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "CertificateProfile", "value": *sdk.CertificateProfile})
	}

	// Handling Primitives
	if !model.HttpMethod.IsNull() && !model.HttpMethod.IsUnknown() {
		sdk.HttpMethod = model.HttpMethod.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "HttpMethod", "value": *sdk.HttpMethod})
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
	if !model.Protocol.IsNull() && !model.Protocol.IsUnknown() {
		sdk.Protocol = model.Protocol.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Protocol", "value": *sdk.Protocol})
	}

	// Handling Primitives
	if !model.TlsVersion.IsNull() && !model.TlsVersion.IsUnknown() {
		sdk.TlsVersion = model.TlsVersion.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "TlsVersion", "value": *sdk.TlsVersion})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.HttpServerProfilesServerInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for HttpServerProfilesServerInner ---
func packHttpServerProfilesServerInnerFromSdk(ctx context.Context, sdk objects.HttpServerProfilesServerInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.HttpServerProfilesServerInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.HttpServerProfilesServerInner
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Address != nil {
		model.Address = basetypes.NewStringValue(*sdk.Address)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Address", "value": *sdk.Address})
	} else {
		model.Address = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.CertificateProfile != nil {
		model.CertificateProfile = basetypes.NewStringValue(*sdk.CertificateProfile)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "CertificateProfile", "value": *sdk.CertificateProfile})
	} else {
		model.CertificateProfile = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.HttpMethod != nil {
		model.HttpMethod = basetypes.NewStringValue(*sdk.HttpMethod)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "HttpMethod", "value": *sdk.HttpMethod})
	} else {
		model.HttpMethod = basetypes.NewStringNull()
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
	if sdk.Protocol != nil {
		model.Protocol = basetypes.NewStringValue(*sdk.Protocol)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Protocol", "value": *sdk.Protocol})
	} else {
		model.Protocol = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.TlsVersion != nil {
		model.TlsVersion = basetypes.NewStringValue(*sdk.TlsVersion)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "TlsVersion", "value": *sdk.TlsVersion})
	} else {
		model.TlsVersion = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.HttpServerProfilesServerInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.HttpServerProfilesServerInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for HttpServerProfilesServerInner ---
func unpackHttpServerProfilesServerInnerListToSdk(ctx context.Context, list types.List) ([]objects.HttpServerProfilesServerInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.HttpServerProfilesServerInner")
	diags := diag.Diagnostics{}
	var data []models.HttpServerProfilesServerInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.HttpServerProfilesServerInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.HttpServerProfilesServerInner{}.AttrTypes(), &item)
		unpacked, d := unpackHttpServerProfilesServerInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.HttpServerProfilesServerInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for HttpServerProfilesServerInner ---
func packHttpServerProfilesServerInnerListFromSdk(ctx context.Context, sdks []objects.HttpServerProfilesServerInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.HttpServerProfilesServerInner")
	diags := diag.Diagnostics{}
	var data []models.HttpServerProfilesServerInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.HttpServerProfilesServerInner
		obj, d := packHttpServerProfilesServerInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.HttpServerProfilesServerInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.HttpServerProfilesServerInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.HttpServerProfilesServerInner{}.AttrType(), data)
}
