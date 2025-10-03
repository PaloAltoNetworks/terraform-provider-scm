package provider

import (
	"context"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/identity_services"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/identity_services"
)

// --- Unpacker for TlsServiceProfiles ---
func unpackTlsServiceProfilesToSdk(ctx context.Context, obj types.Object) (*identity_services.TlsServiceProfiles, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.TlsServiceProfiles", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.TlsServiceProfiles
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk identity_services.TlsServiceProfiles
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Certificate.IsNull() && !model.Certificate.IsUnknown() {
		sdk.Certificate = model.Certificate.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Certificate", "value": sdk.Certificate})
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
		sdk.Id = model.Id.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Id", "value": sdk.Id})
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	// Handling Objects
	if !model.ProtocolSettings.IsNull() && !model.ProtocolSettings.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field ProtocolSettings")
		unpacked, d := unpackTlsServiceProfilesProtocolSettingsToSdk(ctx, model.ProtocolSettings)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "ProtocolSettings"})
		}
		if unpacked != nil {
			sdk.ProtocolSettings = *unpacked
		}
	}

	// Handling Primitives
	if !model.Snippet.IsNull() && !model.Snippet.IsUnknown() {
		sdk.Snippet = model.Snippet.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.TlsServiceProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for TlsServiceProfiles ---
func packTlsServiceProfilesFromSdk(ctx context.Context, sdk identity_services.TlsServiceProfiles) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.TlsServiceProfiles", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.TlsServiceProfiles
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.Certificate = basetypes.NewStringValue(sdk.Certificate)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Certificate", "value": sdk.Certificate})
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
	// Handling Objects
	// This is a regular nested object that has its own packer.
	// Logic for non-pointer / value-type nested objects
	if !reflect.ValueOf(sdk.ProtocolSettings).IsZero() {
		tflog.Debug(ctx, "Packing nested object for field ProtocolSettings")
		packed, d := packTlsServiceProfilesProtocolSettingsFromSdk(ctx, sdk.ProtocolSettings)
		diags.Append(d...)
		model.ProtocolSettings = packed
	} else {
		model.ProtocolSettings = basetypes.NewObjectNull(models.TlsServiceProfilesProtocolSettings{}.AttrTypes())
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

	obj, d := types.ObjectValueFrom(ctx, models.TlsServiceProfiles{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.TlsServiceProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for TlsServiceProfiles ---
func unpackTlsServiceProfilesListToSdk(ctx context.Context, list types.List) ([]identity_services.TlsServiceProfiles, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.TlsServiceProfiles")
	diags := diag.Diagnostics{}
	var data []models.TlsServiceProfiles
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]identity_services.TlsServiceProfiles, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.TlsServiceProfiles{}.AttrTypes(), &item)
		unpacked, d := unpackTlsServiceProfilesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.TlsServiceProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for TlsServiceProfiles ---
func packTlsServiceProfilesListFromSdk(ctx context.Context, sdks []identity_services.TlsServiceProfiles) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.TlsServiceProfiles")
	diags := diag.Diagnostics{}
	var data []models.TlsServiceProfiles

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.TlsServiceProfiles
		obj, d := packTlsServiceProfilesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.TlsServiceProfiles{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.TlsServiceProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.TlsServiceProfiles{}.AttrType(), data)
}

// --- Unpacker for TlsServiceProfilesProtocolSettings ---
func unpackTlsServiceProfilesProtocolSettingsToSdk(ctx context.Context, obj types.Object) (*identity_services.TlsServiceProfilesProtocolSettings, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.TlsServiceProfilesProtocolSettings", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.TlsServiceProfilesProtocolSettings
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk identity_services.TlsServiceProfilesProtocolSettings
	var d diag.Diagnostics
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

	tflog.Debug(ctx, "Exiting unpack helper for models.TlsServiceProfilesProtocolSettings", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for TlsServiceProfilesProtocolSettings ---
func packTlsServiceProfilesProtocolSettingsFromSdk(ctx context.Context, sdk identity_services.TlsServiceProfilesProtocolSettings) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.TlsServiceProfilesProtocolSettings", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.TlsServiceProfilesProtocolSettings
	var d diag.Diagnostics
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

	obj, d := types.ObjectValueFrom(ctx, models.TlsServiceProfilesProtocolSettings{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.TlsServiceProfilesProtocolSettings", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for TlsServiceProfilesProtocolSettings ---
func unpackTlsServiceProfilesProtocolSettingsListToSdk(ctx context.Context, list types.List) ([]identity_services.TlsServiceProfilesProtocolSettings, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.TlsServiceProfilesProtocolSettings")
	diags := diag.Diagnostics{}
	var data []models.TlsServiceProfilesProtocolSettings
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]identity_services.TlsServiceProfilesProtocolSettings, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.TlsServiceProfilesProtocolSettings{}.AttrTypes(), &item)
		unpacked, d := unpackTlsServiceProfilesProtocolSettingsToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.TlsServiceProfilesProtocolSettings", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for TlsServiceProfilesProtocolSettings ---
func packTlsServiceProfilesProtocolSettingsListFromSdk(ctx context.Context, sdks []identity_services.TlsServiceProfilesProtocolSettings) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.TlsServiceProfilesProtocolSettings")
	diags := diag.Diagnostics{}
	var data []models.TlsServiceProfilesProtocolSettings

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.TlsServiceProfilesProtocolSettings
		obj, d := packTlsServiceProfilesProtocolSettingsFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.TlsServiceProfilesProtocolSettings{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.TlsServiceProfilesProtocolSettings", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.TlsServiceProfilesProtocolSettings{}.AttrType(), data)
}
