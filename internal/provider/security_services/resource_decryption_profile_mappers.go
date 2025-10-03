package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/security_services"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/security_services"
)

// --- Unpacker for DecryptionProfiles ---
func unpackDecryptionProfilesToSdk(ctx context.Context, obj types.Object) (*security_services.DecryptionProfiles, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.DecryptionProfiles", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.DecryptionProfiles
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.DecryptionProfiles
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
		sdk.Id = model.Id.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Id", "value": sdk.Id})
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

	// Handling Objects
	if !model.SslForwardProxy.IsNull() && !model.SslForwardProxy.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field SslForwardProxy")
		unpacked, d := unpackDecryptionProfilesSslForwardProxyToSdk(ctx, model.SslForwardProxy)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "SslForwardProxy"})
		}
		if unpacked != nil {
			sdk.SslForwardProxy = unpacked
		}
	}

	// Handling Objects
	if !model.SslInboundProxy.IsNull() && !model.SslInboundProxy.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field SslInboundProxy")
		unpacked, d := unpackDecryptionProfilesSslInboundProxyToSdk(ctx, model.SslInboundProxy)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "SslInboundProxy"})
		}
		if unpacked != nil {
			sdk.SslInboundProxy = unpacked
		}
	}

	// Handling Objects
	if !model.SslNoProxy.IsNull() && !model.SslNoProxy.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field SslNoProxy")
		unpacked, d := unpackDecryptionProfilesSslNoProxyToSdk(ctx, model.SslNoProxy)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "SslNoProxy"})
		}
		if unpacked != nil {
			sdk.SslNoProxy = unpacked
		}
	}

	// Handling Objects
	if !model.SslProtocolSettings.IsNull() && !model.SslProtocolSettings.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field SslProtocolSettings")
		unpacked, d := unpackDecryptionProfilesSslProtocolSettingsToSdk(ctx, model.SslProtocolSettings)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "SslProtocolSettings"})
		}
		if unpacked != nil {
			sdk.SslProtocolSettings = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.DecryptionProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for DecryptionProfiles ---
func packDecryptionProfilesFromSdk(ctx context.Context, sdk security_services.DecryptionProfiles) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.DecryptionProfiles", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.DecryptionProfiles
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
	model.Id = basetypes.NewStringValue(sdk.Id)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Id", "value": sdk.Id})
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
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.SslForwardProxy != nil {
		tflog.Debug(ctx, "Packing nested object for field SslForwardProxy")
		packed, d := packDecryptionProfilesSslForwardProxyFromSdk(ctx, *sdk.SslForwardProxy)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "SslForwardProxy"})
		}
		model.SslForwardProxy = packed
	} else {
		model.SslForwardProxy = basetypes.NewObjectNull(models.DecryptionProfilesSslForwardProxy{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.SslInboundProxy != nil {
		tflog.Debug(ctx, "Packing nested object for field SslInboundProxy")
		packed, d := packDecryptionProfilesSslInboundProxyFromSdk(ctx, *sdk.SslInboundProxy)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "SslInboundProxy"})
		}
		model.SslInboundProxy = packed
	} else {
		model.SslInboundProxy = basetypes.NewObjectNull(models.DecryptionProfilesSslInboundProxy{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.SslNoProxy != nil {
		tflog.Debug(ctx, "Packing nested object for field SslNoProxy")
		packed, d := packDecryptionProfilesSslNoProxyFromSdk(ctx, *sdk.SslNoProxy)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "SslNoProxy"})
		}
		model.SslNoProxy = packed
	} else {
		model.SslNoProxy = basetypes.NewObjectNull(models.DecryptionProfilesSslNoProxy{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.SslProtocolSettings != nil {
		tflog.Debug(ctx, "Packing nested object for field SslProtocolSettings")
		packed, d := packDecryptionProfilesSslProtocolSettingsFromSdk(ctx, *sdk.SslProtocolSettings)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "SslProtocolSettings"})
		}
		model.SslProtocolSettings = packed
	} else {
		model.SslProtocolSettings = basetypes.NewObjectNull(models.DecryptionProfilesSslProtocolSettings{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.DecryptionProfiles{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.DecryptionProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for DecryptionProfiles ---
func unpackDecryptionProfilesListToSdk(ctx context.Context, list types.List) ([]security_services.DecryptionProfiles, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.DecryptionProfiles")
	diags := diag.Diagnostics{}
	var data []models.DecryptionProfiles
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.DecryptionProfiles, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.DecryptionProfiles{}.AttrTypes(), &item)
		unpacked, d := unpackDecryptionProfilesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.DecryptionProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for DecryptionProfiles ---
func packDecryptionProfilesListFromSdk(ctx context.Context, sdks []security_services.DecryptionProfiles) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.DecryptionProfiles")
	diags := diag.Diagnostics{}
	var data []models.DecryptionProfiles

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.DecryptionProfiles
		obj, d := packDecryptionProfilesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.DecryptionProfiles{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.DecryptionProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.DecryptionProfiles{}.AttrType(), data)
}

// --- Unpacker for DecryptionProfilesSslForwardProxy ---
func unpackDecryptionProfilesSslForwardProxyToSdk(ctx context.Context, obj types.Object) (*security_services.DecryptionProfilesSslForwardProxy, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.DecryptionProfilesSslForwardProxy", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.DecryptionProfilesSslForwardProxy
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.DecryptionProfilesSslForwardProxy
	var d diag.Diagnostics
	// Handling Primitives
	if !model.AutoIncludeAltname.IsNull() && !model.AutoIncludeAltname.IsUnknown() {
		sdk.AutoIncludeAltname = model.AutoIncludeAltname.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "AutoIncludeAltname", "value": *sdk.AutoIncludeAltname})
	}

	// Handling Primitives
	if !model.BlockClientCert.IsNull() && !model.BlockClientCert.IsUnknown() {
		sdk.BlockClientCert = model.BlockClientCert.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "BlockClientCert", "value": *sdk.BlockClientCert})
	}

	// Handling Primitives
	if !model.BlockExpiredCertificate.IsNull() && !model.BlockExpiredCertificate.IsUnknown() {
		sdk.BlockExpiredCertificate = model.BlockExpiredCertificate.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "BlockExpiredCertificate", "value": *sdk.BlockExpiredCertificate})
	}

	// Handling Primitives
	if !model.BlockTimeoutCert.IsNull() && !model.BlockTimeoutCert.IsUnknown() {
		sdk.BlockTimeoutCert = model.BlockTimeoutCert.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "BlockTimeoutCert", "value": *sdk.BlockTimeoutCert})
	}

	// Handling Primitives
	if !model.BlockTls13DowngradeNoResource.IsNull() && !model.BlockTls13DowngradeNoResource.IsUnknown() {
		sdk.BlockTls13DowngradeNoResource = model.BlockTls13DowngradeNoResource.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "BlockTls13DowngradeNoResource", "value": *sdk.BlockTls13DowngradeNoResource})
	}

	// Handling Primitives
	if !model.BlockUnknownCert.IsNull() && !model.BlockUnknownCert.IsUnknown() {
		sdk.BlockUnknownCert = model.BlockUnknownCert.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "BlockUnknownCert", "value": *sdk.BlockUnknownCert})
	}

	// Handling Primitives
	if !model.BlockUnsupportedCipher.IsNull() && !model.BlockUnsupportedCipher.IsUnknown() {
		sdk.BlockUnsupportedCipher = model.BlockUnsupportedCipher.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "BlockUnsupportedCipher", "value": *sdk.BlockUnsupportedCipher})
	}

	// Handling Primitives
	if !model.BlockUnsupportedVersion.IsNull() && !model.BlockUnsupportedVersion.IsUnknown() {
		sdk.BlockUnsupportedVersion = model.BlockUnsupportedVersion.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "BlockUnsupportedVersion", "value": *sdk.BlockUnsupportedVersion})
	}

	// Handling Primitives
	if !model.BlockUntrustedIssuer.IsNull() && !model.BlockUntrustedIssuer.IsUnknown() {
		sdk.BlockUntrustedIssuer = model.BlockUntrustedIssuer.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "BlockUntrustedIssuer", "value": *sdk.BlockUntrustedIssuer})
	}

	// Handling Primitives
	if !model.RestrictCertExts.IsNull() && !model.RestrictCertExts.IsUnknown() {
		sdk.RestrictCertExts = model.RestrictCertExts.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "RestrictCertExts", "value": *sdk.RestrictCertExts})
	}

	// Handling Primitives
	if !model.StripAlpn.IsNull() && !model.StripAlpn.IsUnknown() {
		sdk.StripAlpn = model.StripAlpn.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "StripAlpn", "value": *sdk.StripAlpn})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.DecryptionProfilesSslForwardProxy", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for DecryptionProfilesSslForwardProxy ---
func packDecryptionProfilesSslForwardProxyFromSdk(ctx context.Context, sdk security_services.DecryptionProfilesSslForwardProxy) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.DecryptionProfilesSslForwardProxy", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.DecryptionProfilesSslForwardProxy
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.AutoIncludeAltname != nil {
		model.AutoIncludeAltname = basetypes.NewBoolValue(*sdk.AutoIncludeAltname)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "AutoIncludeAltname", "value": *sdk.AutoIncludeAltname})
	} else {
		model.AutoIncludeAltname = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.BlockClientCert != nil {
		model.BlockClientCert = basetypes.NewBoolValue(*sdk.BlockClientCert)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "BlockClientCert", "value": *sdk.BlockClientCert})
	} else {
		model.BlockClientCert = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.BlockExpiredCertificate != nil {
		model.BlockExpiredCertificate = basetypes.NewBoolValue(*sdk.BlockExpiredCertificate)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "BlockExpiredCertificate", "value": *sdk.BlockExpiredCertificate})
	} else {
		model.BlockExpiredCertificate = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.BlockTimeoutCert != nil {
		model.BlockTimeoutCert = basetypes.NewBoolValue(*sdk.BlockTimeoutCert)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "BlockTimeoutCert", "value": *sdk.BlockTimeoutCert})
	} else {
		model.BlockTimeoutCert = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.BlockTls13DowngradeNoResource != nil {
		model.BlockTls13DowngradeNoResource = basetypes.NewBoolValue(*sdk.BlockTls13DowngradeNoResource)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "BlockTls13DowngradeNoResource", "value": *sdk.BlockTls13DowngradeNoResource})
	} else {
		model.BlockTls13DowngradeNoResource = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.BlockUnknownCert != nil {
		model.BlockUnknownCert = basetypes.NewBoolValue(*sdk.BlockUnknownCert)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "BlockUnknownCert", "value": *sdk.BlockUnknownCert})
	} else {
		model.BlockUnknownCert = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.BlockUnsupportedCipher != nil {
		model.BlockUnsupportedCipher = basetypes.NewBoolValue(*sdk.BlockUnsupportedCipher)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "BlockUnsupportedCipher", "value": *sdk.BlockUnsupportedCipher})
	} else {
		model.BlockUnsupportedCipher = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.BlockUnsupportedVersion != nil {
		model.BlockUnsupportedVersion = basetypes.NewBoolValue(*sdk.BlockUnsupportedVersion)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "BlockUnsupportedVersion", "value": *sdk.BlockUnsupportedVersion})
	} else {
		model.BlockUnsupportedVersion = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.BlockUntrustedIssuer != nil {
		model.BlockUntrustedIssuer = basetypes.NewBoolValue(*sdk.BlockUntrustedIssuer)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "BlockUntrustedIssuer", "value": *sdk.BlockUntrustedIssuer})
	} else {
		model.BlockUntrustedIssuer = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.RestrictCertExts != nil {
		model.RestrictCertExts = basetypes.NewBoolValue(*sdk.RestrictCertExts)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "RestrictCertExts", "value": *sdk.RestrictCertExts})
	} else {
		model.RestrictCertExts = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.StripAlpn != nil {
		model.StripAlpn = basetypes.NewBoolValue(*sdk.StripAlpn)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "StripAlpn", "value": *sdk.StripAlpn})
	} else {
		model.StripAlpn = basetypes.NewBoolNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.DecryptionProfilesSslForwardProxy{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.DecryptionProfilesSslForwardProxy", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for DecryptionProfilesSslForwardProxy ---
func unpackDecryptionProfilesSslForwardProxyListToSdk(ctx context.Context, list types.List) ([]security_services.DecryptionProfilesSslForwardProxy, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.DecryptionProfilesSslForwardProxy")
	diags := diag.Diagnostics{}
	var data []models.DecryptionProfilesSslForwardProxy
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.DecryptionProfilesSslForwardProxy, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.DecryptionProfilesSslForwardProxy{}.AttrTypes(), &item)
		unpacked, d := unpackDecryptionProfilesSslForwardProxyToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.DecryptionProfilesSslForwardProxy", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for DecryptionProfilesSslForwardProxy ---
func packDecryptionProfilesSslForwardProxyListFromSdk(ctx context.Context, sdks []security_services.DecryptionProfilesSslForwardProxy) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.DecryptionProfilesSslForwardProxy")
	diags := diag.Diagnostics{}
	var data []models.DecryptionProfilesSslForwardProxy

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.DecryptionProfilesSslForwardProxy
		obj, d := packDecryptionProfilesSslForwardProxyFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.DecryptionProfilesSslForwardProxy{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.DecryptionProfilesSslForwardProxy", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.DecryptionProfilesSslForwardProxy{}.AttrType(), data)
}

// --- Unpacker for DecryptionProfilesSslInboundProxy ---
func unpackDecryptionProfilesSslInboundProxyToSdk(ctx context.Context, obj types.Object) (*security_services.DecryptionProfilesSslInboundProxy, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.DecryptionProfilesSslInboundProxy", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.DecryptionProfilesSslInboundProxy
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.DecryptionProfilesSslInboundProxy
	var d diag.Diagnostics
	// Handling Primitives
	if !model.BlockIfHsmUnavailable.IsNull() && !model.BlockIfHsmUnavailable.IsUnknown() {
		sdk.BlockIfHsmUnavailable = model.BlockIfHsmUnavailable.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "BlockIfHsmUnavailable", "value": *sdk.BlockIfHsmUnavailable})
	}

	// Handling Primitives
	if !model.BlockIfNoResource.IsNull() && !model.BlockIfNoResource.IsUnknown() {
		sdk.BlockIfNoResource = model.BlockIfNoResource.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "BlockIfNoResource", "value": *sdk.BlockIfNoResource})
	}

	// Handling Primitives
	if !model.BlockUnsupportedCipher.IsNull() && !model.BlockUnsupportedCipher.IsUnknown() {
		sdk.BlockUnsupportedCipher = model.BlockUnsupportedCipher.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "BlockUnsupportedCipher", "value": *sdk.BlockUnsupportedCipher})
	}

	// Handling Primitives
	if !model.BlockUnsupportedVersion.IsNull() && !model.BlockUnsupportedVersion.IsUnknown() {
		sdk.BlockUnsupportedVersion = model.BlockUnsupportedVersion.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "BlockUnsupportedVersion", "value": *sdk.BlockUnsupportedVersion})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.DecryptionProfilesSslInboundProxy", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for DecryptionProfilesSslInboundProxy ---
func packDecryptionProfilesSslInboundProxyFromSdk(ctx context.Context, sdk security_services.DecryptionProfilesSslInboundProxy) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.DecryptionProfilesSslInboundProxy", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.DecryptionProfilesSslInboundProxy
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.BlockIfHsmUnavailable != nil {
		model.BlockIfHsmUnavailable = basetypes.NewBoolValue(*sdk.BlockIfHsmUnavailable)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "BlockIfHsmUnavailable", "value": *sdk.BlockIfHsmUnavailable})
	} else {
		model.BlockIfHsmUnavailable = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.BlockIfNoResource != nil {
		model.BlockIfNoResource = basetypes.NewBoolValue(*sdk.BlockIfNoResource)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "BlockIfNoResource", "value": *sdk.BlockIfNoResource})
	} else {
		model.BlockIfNoResource = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.BlockUnsupportedCipher != nil {
		model.BlockUnsupportedCipher = basetypes.NewBoolValue(*sdk.BlockUnsupportedCipher)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "BlockUnsupportedCipher", "value": *sdk.BlockUnsupportedCipher})
	} else {
		model.BlockUnsupportedCipher = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.BlockUnsupportedVersion != nil {
		model.BlockUnsupportedVersion = basetypes.NewBoolValue(*sdk.BlockUnsupportedVersion)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "BlockUnsupportedVersion", "value": *sdk.BlockUnsupportedVersion})
	} else {
		model.BlockUnsupportedVersion = basetypes.NewBoolNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.DecryptionProfilesSslInboundProxy{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.DecryptionProfilesSslInboundProxy", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for DecryptionProfilesSslInboundProxy ---
func unpackDecryptionProfilesSslInboundProxyListToSdk(ctx context.Context, list types.List) ([]security_services.DecryptionProfilesSslInboundProxy, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.DecryptionProfilesSslInboundProxy")
	diags := diag.Diagnostics{}
	var data []models.DecryptionProfilesSslInboundProxy
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.DecryptionProfilesSslInboundProxy, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.DecryptionProfilesSslInboundProxy{}.AttrTypes(), &item)
		unpacked, d := unpackDecryptionProfilesSslInboundProxyToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.DecryptionProfilesSslInboundProxy", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for DecryptionProfilesSslInboundProxy ---
func packDecryptionProfilesSslInboundProxyListFromSdk(ctx context.Context, sdks []security_services.DecryptionProfilesSslInboundProxy) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.DecryptionProfilesSslInboundProxy")
	diags := diag.Diagnostics{}
	var data []models.DecryptionProfilesSslInboundProxy

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.DecryptionProfilesSslInboundProxy
		obj, d := packDecryptionProfilesSslInboundProxyFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.DecryptionProfilesSslInboundProxy{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.DecryptionProfilesSslInboundProxy", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.DecryptionProfilesSslInboundProxy{}.AttrType(), data)
}

// --- Unpacker for DecryptionProfilesSslNoProxy ---
func unpackDecryptionProfilesSslNoProxyToSdk(ctx context.Context, obj types.Object) (*security_services.DecryptionProfilesSslNoProxy, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.DecryptionProfilesSslNoProxy", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.DecryptionProfilesSslNoProxy
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.DecryptionProfilesSslNoProxy
	var d diag.Diagnostics
	// Handling Primitives
	if !model.BlockExpiredCertificate.IsNull() && !model.BlockExpiredCertificate.IsUnknown() {
		sdk.BlockExpiredCertificate = model.BlockExpiredCertificate.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "BlockExpiredCertificate", "value": *sdk.BlockExpiredCertificate})
	}

	// Handling Primitives
	if !model.BlockUntrustedIssuer.IsNull() && !model.BlockUntrustedIssuer.IsUnknown() {
		sdk.BlockUntrustedIssuer = model.BlockUntrustedIssuer.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "BlockUntrustedIssuer", "value": *sdk.BlockUntrustedIssuer})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.DecryptionProfilesSslNoProxy", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for DecryptionProfilesSslNoProxy ---
func packDecryptionProfilesSslNoProxyFromSdk(ctx context.Context, sdk security_services.DecryptionProfilesSslNoProxy) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.DecryptionProfilesSslNoProxy", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.DecryptionProfilesSslNoProxy
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.BlockExpiredCertificate != nil {
		model.BlockExpiredCertificate = basetypes.NewBoolValue(*sdk.BlockExpiredCertificate)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "BlockExpiredCertificate", "value": *sdk.BlockExpiredCertificate})
	} else {
		model.BlockExpiredCertificate = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.BlockUntrustedIssuer != nil {
		model.BlockUntrustedIssuer = basetypes.NewBoolValue(*sdk.BlockUntrustedIssuer)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "BlockUntrustedIssuer", "value": *sdk.BlockUntrustedIssuer})
	} else {
		model.BlockUntrustedIssuer = basetypes.NewBoolNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.DecryptionProfilesSslNoProxy{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.DecryptionProfilesSslNoProxy", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for DecryptionProfilesSslNoProxy ---
func unpackDecryptionProfilesSslNoProxyListToSdk(ctx context.Context, list types.List) ([]security_services.DecryptionProfilesSslNoProxy, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.DecryptionProfilesSslNoProxy")
	diags := diag.Diagnostics{}
	var data []models.DecryptionProfilesSslNoProxy
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.DecryptionProfilesSslNoProxy, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.DecryptionProfilesSslNoProxy{}.AttrTypes(), &item)
		unpacked, d := unpackDecryptionProfilesSslNoProxyToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.DecryptionProfilesSslNoProxy", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for DecryptionProfilesSslNoProxy ---
func packDecryptionProfilesSslNoProxyListFromSdk(ctx context.Context, sdks []security_services.DecryptionProfilesSslNoProxy) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.DecryptionProfilesSslNoProxy")
	diags := diag.Diagnostics{}
	var data []models.DecryptionProfilesSslNoProxy

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.DecryptionProfilesSslNoProxy
		obj, d := packDecryptionProfilesSslNoProxyFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.DecryptionProfilesSslNoProxy{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.DecryptionProfilesSslNoProxy", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.DecryptionProfilesSslNoProxy{}.AttrType(), data)
}

// --- Unpacker for DecryptionProfilesSslProtocolSettings ---
func unpackDecryptionProfilesSslProtocolSettingsToSdk(ctx context.Context, obj types.Object) (*security_services.DecryptionProfilesSslProtocolSettings, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.DecryptionProfilesSslProtocolSettings", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.DecryptionProfilesSslProtocolSettings
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.DecryptionProfilesSslProtocolSettings
	var d diag.Diagnostics
	// Handling Primitives
	if !model.AuthAlgoMd5.IsNull() && !model.AuthAlgoMd5.IsUnknown() {
		sdk.AuthAlgoMd5 = model.AuthAlgoMd5.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "AuthAlgoMd5", "value": *sdk.AuthAlgoMd5})
	}

	// Handling Primitives
	if !model.AuthAlgoSha1.IsNull() && !model.AuthAlgoSha1.IsUnknown() {
		sdk.AuthAlgoSha1 = model.AuthAlgoSha1.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "AuthAlgoSha1", "value": *sdk.AuthAlgoSha1})
	}

	// Handling Primitives
	if !model.AuthAlgoSha256.IsNull() && !model.AuthAlgoSha256.IsUnknown() {
		sdk.AuthAlgoSha256 = model.AuthAlgoSha256.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "AuthAlgoSha256", "value": *sdk.AuthAlgoSha256})
	}

	// Handling Primitives
	if !model.AuthAlgoSha384.IsNull() && !model.AuthAlgoSha384.IsUnknown() {
		sdk.AuthAlgoSha384 = model.AuthAlgoSha384.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "AuthAlgoSha384", "value": *sdk.AuthAlgoSha384})
	}

	// Handling Primitives
	if !model.EncAlgo3des.IsNull() && !model.EncAlgo3des.IsUnknown() {
		sdk.EncAlgo3des = model.EncAlgo3des.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "EncAlgo3des", "value": *sdk.EncAlgo3des})
	}

	// Handling Primitives
	if !model.EncAlgoAes128Cbc.IsNull() && !model.EncAlgoAes128Cbc.IsUnknown() {
		sdk.EncAlgoAes128Cbc = model.EncAlgoAes128Cbc.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "EncAlgoAes128Cbc", "value": *sdk.EncAlgoAes128Cbc})
	}

	// Handling Primitives
	if !model.EncAlgoAes128Gcm.IsNull() && !model.EncAlgoAes128Gcm.IsUnknown() {
		sdk.EncAlgoAes128Gcm = model.EncAlgoAes128Gcm.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "EncAlgoAes128Gcm", "value": *sdk.EncAlgoAes128Gcm})
	}

	// Handling Primitives
	if !model.EncAlgoAes256Cbc.IsNull() && !model.EncAlgoAes256Cbc.IsUnknown() {
		sdk.EncAlgoAes256Cbc = model.EncAlgoAes256Cbc.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "EncAlgoAes256Cbc", "value": *sdk.EncAlgoAes256Cbc})
	}

	// Handling Primitives
	if !model.EncAlgoAes256Gcm.IsNull() && !model.EncAlgoAes256Gcm.IsUnknown() {
		sdk.EncAlgoAes256Gcm = model.EncAlgoAes256Gcm.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "EncAlgoAes256Gcm", "value": *sdk.EncAlgoAes256Gcm})
	}

	// Handling Primitives
	if !model.EncAlgoChacha20Poly1305.IsNull() && !model.EncAlgoChacha20Poly1305.IsUnknown() {
		sdk.EncAlgoChacha20Poly1305 = model.EncAlgoChacha20Poly1305.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "EncAlgoChacha20Poly1305", "value": *sdk.EncAlgoChacha20Poly1305})
	}

	// Handling Primitives
	if !model.EncAlgoRc4.IsNull() && !model.EncAlgoRc4.IsUnknown() {
		sdk.EncAlgoRc4 = model.EncAlgoRc4.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "EncAlgoRc4", "value": *sdk.EncAlgoRc4})
	}

	// Handling Primitives
	if !model.KeyxchgAlgoDhe.IsNull() && !model.KeyxchgAlgoDhe.IsUnknown() {
		sdk.KeyxchgAlgoDhe = model.KeyxchgAlgoDhe.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "KeyxchgAlgoDhe", "value": *sdk.KeyxchgAlgoDhe})
	}

	// Handling Primitives
	if !model.KeyxchgAlgoEcdhe.IsNull() && !model.KeyxchgAlgoEcdhe.IsUnknown() {
		sdk.KeyxchgAlgoEcdhe = model.KeyxchgAlgoEcdhe.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "KeyxchgAlgoEcdhe", "value": *sdk.KeyxchgAlgoEcdhe})
	}

	// Handling Primitives
	if !model.KeyxchgAlgoRsa.IsNull() && !model.KeyxchgAlgoRsa.IsUnknown() {
		sdk.KeyxchgAlgoRsa = model.KeyxchgAlgoRsa.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "KeyxchgAlgoRsa", "value": *sdk.KeyxchgAlgoRsa})
	}

	// Handling Primitives
	if !model.MaxVersion.IsNull() && !model.MaxVersion.IsUnknown() {
		sdk.MaxVersion = model.MaxVersion.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "MaxVersion", "value": *sdk.MaxVersion})
	}

	// Handling Primitives
	if !model.MinVersion.IsNull() && !model.MinVersion.IsUnknown() {
		sdk.MinVersion = model.MinVersion.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "MinVersion", "value": *sdk.MinVersion})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.DecryptionProfilesSslProtocolSettings", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for DecryptionProfilesSslProtocolSettings ---
func packDecryptionProfilesSslProtocolSettingsFromSdk(ctx context.Context, sdk security_services.DecryptionProfilesSslProtocolSettings) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.DecryptionProfilesSslProtocolSettings", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.DecryptionProfilesSslProtocolSettings
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.AuthAlgoMd5 != nil {
		model.AuthAlgoMd5 = basetypes.NewBoolValue(*sdk.AuthAlgoMd5)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "AuthAlgoMd5", "value": *sdk.AuthAlgoMd5})
	} else {
		model.AuthAlgoMd5 = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.AuthAlgoSha1 != nil {
		model.AuthAlgoSha1 = basetypes.NewBoolValue(*sdk.AuthAlgoSha1)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "AuthAlgoSha1", "value": *sdk.AuthAlgoSha1})
	} else {
		model.AuthAlgoSha1 = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.AuthAlgoSha256 != nil {
		model.AuthAlgoSha256 = basetypes.NewBoolValue(*sdk.AuthAlgoSha256)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "AuthAlgoSha256", "value": *sdk.AuthAlgoSha256})
	} else {
		model.AuthAlgoSha256 = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.AuthAlgoSha384 != nil {
		model.AuthAlgoSha384 = basetypes.NewBoolValue(*sdk.AuthAlgoSha384)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "AuthAlgoSha384", "value": *sdk.AuthAlgoSha384})
	} else {
		model.AuthAlgoSha384 = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.EncAlgo3des != nil {
		model.EncAlgo3des = basetypes.NewBoolValue(*sdk.EncAlgo3des)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "EncAlgo3des", "value": *sdk.EncAlgo3des})
	} else {
		model.EncAlgo3des = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.EncAlgoAes128Cbc != nil {
		model.EncAlgoAes128Cbc = basetypes.NewBoolValue(*sdk.EncAlgoAes128Cbc)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "EncAlgoAes128Cbc", "value": *sdk.EncAlgoAes128Cbc})
	} else {
		model.EncAlgoAes128Cbc = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.EncAlgoAes128Gcm != nil {
		model.EncAlgoAes128Gcm = basetypes.NewBoolValue(*sdk.EncAlgoAes128Gcm)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "EncAlgoAes128Gcm", "value": *sdk.EncAlgoAes128Gcm})
	} else {
		model.EncAlgoAes128Gcm = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.EncAlgoAes256Cbc != nil {
		model.EncAlgoAes256Cbc = basetypes.NewBoolValue(*sdk.EncAlgoAes256Cbc)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "EncAlgoAes256Cbc", "value": *sdk.EncAlgoAes256Cbc})
	} else {
		model.EncAlgoAes256Cbc = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.EncAlgoAes256Gcm != nil {
		model.EncAlgoAes256Gcm = basetypes.NewBoolValue(*sdk.EncAlgoAes256Gcm)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "EncAlgoAes256Gcm", "value": *sdk.EncAlgoAes256Gcm})
	} else {
		model.EncAlgoAes256Gcm = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.EncAlgoChacha20Poly1305 != nil {
		model.EncAlgoChacha20Poly1305 = basetypes.NewBoolValue(*sdk.EncAlgoChacha20Poly1305)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "EncAlgoChacha20Poly1305", "value": *sdk.EncAlgoChacha20Poly1305})
	} else {
		model.EncAlgoChacha20Poly1305 = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.EncAlgoRc4 != nil {
		model.EncAlgoRc4 = basetypes.NewBoolValue(*sdk.EncAlgoRc4)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "EncAlgoRc4", "value": *sdk.EncAlgoRc4})
	} else {
		model.EncAlgoRc4 = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.KeyxchgAlgoDhe != nil {
		model.KeyxchgAlgoDhe = basetypes.NewBoolValue(*sdk.KeyxchgAlgoDhe)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "KeyxchgAlgoDhe", "value": *sdk.KeyxchgAlgoDhe})
	} else {
		model.KeyxchgAlgoDhe = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.KeyxchgAlgoEcdhe != nil {
		model.KeyxchgAlgoEcdhe = basetypes.NewBoolValue(*sdk.KeyxchgAlgoEcdhe)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "KeyxchgAlgoEcdhe", "value": *sdk.KeyxchgAlgoEcdhe})
	} else {
		model.KeyxchgAlgoEcdhe = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.KeyxchgAlgoRsa != nil {
		model.KeyxchgAlgoRsa = basetypes.NewBoolValue(*sdk.KeyxchgAlgoRsa)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "KeyxchgAlgoRsa", "value": *sdk.KeyxchgAlgoRsa})
	} else {
		model.KeyxchgAlgoRsa = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.MaxVersion != nil {
		model.MaxVersion = basetypes.NewStringValue(*sdk.MaxVersion)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "MaxVersion", "value": *sdk.MaxVersion})
	} else {
		model.MaxVersion = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.MinVersion != nil {
		model.MinVersion = basetypes.NewStringValue(*sdk.MinVersion)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "MinVersion", "value": *sdk.MinVersion})
	} else {
		model.MinVersion = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.DecryptionProfilesSslProtocolSettings{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.DecryptionProfilesSslProtocolSettings", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for DecryptionProfilesSslProtocolSettings ---
func unpackDecryptionProfilesSslProtocolSettingsListToSdk(ctx context.Context, list types.List) ([]security_services.DecryptionProfilesSslProtocolSettings, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.DecryptionProfilesSslProtocolSettings")
	diags := diag.Diagnostics{}
	var data []models.DecryptionProfilesSslProtocolSettings
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.DecryptionProfilesSslProtocolSettings, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.DecryptionProfilesSslProtocolSettings{}.AttrTypes(), &item)
		unpacked, d := unpackDecryptionProfilesSslProtocolSettingsToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.DecryptionProfilesSslProtocolSettings", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for DecryptionProfilesSslProtocolSettings ---
func packDecryptionProfilesSslProtocolSettingsListFromSdk(ctx context.Context, sdks []security_services.DecryptionProfilesSslProtocolSettings) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.DecryptionProfilesSslProtocolSettings")
	diags := diag.Diagnostics{}
	var data []models.DecryptionProfilesSslProtocolSettings

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.DecryptionProfilesSslProtocolSettings
		obj, d := packDecryptionProfilesSslProtocolSettingsFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.DecryptionProfilesSslProtocolSettings{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.DecryptionProfilesSslProtocolSettings", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.DecryptionProfilesSslProtocolSettings{}.AttrType(), data)
}
