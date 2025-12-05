package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/device_settings"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/device_settings"
)

// --- Unpacker for AuthenticationSettings ---
func unpackAuthenticationSettingsToSdk(ctx context.Context, obj types.Object) (*device_settings.AuthenticationSettings, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.AuthenticationSettings", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.AuthenticationSettings
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk device_settings.AuthenticationSettings
	var d diag.Diagnostics

	// Handling Objects
	if !model.Authentication.IsNull() && !model.Authentication.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Authentication")
		unpacked, d := unpackAuthenticationSettingsAuthenticationToSdk(ctx, model.Authentication)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Authentication"})
		}
		if unpacked != nil {
			sdk.Authentication = unpacked
		}
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
	if !model.Snippet.IsNull() && !model.Snippet.IsUnknown() {
		sdk.Snippet = model.Snippet.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.AuthenticationSettings", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for AuthenticationSettings ---
func packAuthenticationSettingsFromSdk(ctx context.Context, sdk device_settings.AuthenticationSettings) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.AuthenticationSettings", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.AuthenticationSettings
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Authentication != nil {
		tflog.Debug(ctx, "Packing nested object for field Authentication")
		packed, d := packAuthenticationSettingsAuthenticationFromSdk(ctx, *sdk.Authentication)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Authentication"})
		}
		model.Authentication = packed
	} else {
		model.Authentication = basetypes.NewObjectNull(models.AuthenticationSettingsAuthentication{}.AttrTypes())
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
	model.Id = basetypes.NewStringValue(sdk.Id)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Id", "value": sdk.Id})
	// Handling Primitives
	// Standard primitive packing
	if sdk.Snippet != nil {
		model.Snippet = basetypes.NewStringValue(*sdk.Snippet)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	} else {
		model.Snippet = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.AuthenticationSettings{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.AuthenticationSettings", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for AuthenticationSettings ---
func unpackAuthenticationSettingsListToSdk(ctx context.Context, list types.List) ([]device_settings.AuthenticationSettings, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.AuthenticationSettings")
	diags := diag.Diagnostics{}
	var data []models.AuthenticationSettings
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]device_settings.AuthenticationSettings, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.AuthenticationSettings{}.AttrTypes(), &item)
		unpacked, d := unpackAuthenticationSettingsToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.AuthenticationSettings", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for AuthenticationSettings ---
func packAuthenticationSettingsListFromSdk(ctx context.Context, sdks []device_settings.AuthenticationSettings) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.AuthenticationSettings")
	diags := diag.Diagnostics{}
	var data []models.AuthenticationSettings

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.AuthenticationSettings
		obj, d := packAuthenticationSettingsFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.AuthenticationSettings{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.AuthenticationSettings", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.AuthenticationSettings{}.AttrType(), data)
}

// --- Unpacker for AuthenticationSettingsAuthentication ---
func unpackAuthenticationSettingsAuthenticationToSdk(ctx context.Context, obj types.Object) (*device_settings.AuthenticationSettingsAuthentication, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.AuthenticationSettingsAuthentication", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.AuthenticationSettingsAuthentication
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk device_settings.AuthenticationSettingsAuthentication
	var d diag.Diagnostics
	// Handling Primitives
	if !model.AccountingServerProfile.IsNull() && !model.AccountingServerProfile.IsUnknown() {
		sdk.AccountingServerProfile = model.AccountingServerProfile.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "AccountingServerProfile", "value": *sdk.AccountingServerProfile})
	}

	// Handling Primitives
	if !model.AuthenticationProfile.IsNull() && !model.AuthenticationProfile.IsUnknown() {
		sdk.AuthenticationProfile = model.AuthenticationProfile.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "AuthenticationProfile", "value": *sdk.AuthenticationProfile})
	}

	// Handling Primitives
	if !model.CertificateProfile.IsNull() && !model.CertificateProfile.IsUnknown() {
		sdk.CertificateProfile = model.CertificateProfile.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "CertificateProfile", "value": *sdk.CertificateProfile})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.AuthenticationSettingsAuthentication", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for AuthenticationSettingsAuthentication ---
func packAuthenticationSettingsAuthenticationFromSdk(ctx context.Context, sdk device_settings.AuthenticationSettingsAuthentication) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.AuthenticationSettingsAuthentication", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.AuthenticationSettingsAuthentication
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.AccountingServerProfile != nil {
		model.AccountingServerProfile = basetypes.NewStringValue(*sdk.AccountingServerProfile)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "AccountingServerProfile", "value": *sdk.AccountingServerProfile})
	} else {
		model.AccountingServerProfile = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.AuthenticationProfile != nil {
		model.AuthenticationProfile = basetypes.NewStringValue(*sdk.AuthenticationProfile)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "AuthenticationProfile", "value": *sdk.AuthenticationProfile})
	} else {
		model.AuthenticationProfile = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.CertificateProfile != nil {
		model.CertificateProfile = basetypes.NewStringValue(*sdk.CertificateProfile)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "CertificateProfile", "value": *sdk.CertificateProfile})
	} else {
		model.CertificateProfile = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.AuthenticationSettingsAuthentication{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.AuthenticationSettingsAuthentication", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for AuthenticationSettingsAuthentication ---
func unpackAuthenticationSettingsAuthenticationListToSdk(ctx context.Context, list types.List) ([]device_settings.AuthenticationSettingsAuthentication, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.AuthenticationSettingsAuthentication")
	diags := diag.Diagnostics{}
	var data []models.AuthenticationSettingsAuthentication
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]device_settings.AuthenticationSettingsAuthentication, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.AuthenticationSettingsAuthentication{}.AttrTypes(), &item)
		unpacked, d := unpackAuthenticationSettingsAuthenticationToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.AuthenticationSettingsAuthentication", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for AuthenticationSettingsAuthentication ---
func packAuthenticationSettingsAuthenticationListFromSdk(ctx context.Context, sdks []device_settings.AuthenticationSettingsAuthentication) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.AuthenticationSettingsAuthentication")
	diags := diag.Diagnostics{}
	var data []models.AuthenticationSettingsAuthentication

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.AuthenticationSettingsAuthentication
		obj, d := packAuthenticationSettingsAuthenticationFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.AuthenticationSettingsAuthentication{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.AuthenticationSettingsAuthentication", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.AuthenticationSettingsAuthentication{}.AttrType(), data)
}
