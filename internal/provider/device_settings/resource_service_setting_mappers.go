package provider

import (
	"context"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/device_settings"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/device_settings"
)

// serviceSettingsSensitiveValuePatcher is an in-memory struct to temporarily store plaintext
// and encrypted values for sensitive fields during the Create/Update/Read workflows.
type serviceSettingsSensitiveValuePatcher struct {
	services_ntp_servers_primary_ntp_server_authentication_type_symmetric_key_algorithm_md5_authentication_key_plaintext    basetypes.StringValue
	services_ntp_servers_primary_ntp_server_authentication_type_symmetric_key_algorithm_md5_authentication_key_encrypted    basetypes.StringValue
	services_ntp_servers_primary_ntp_server_authentication_type_symmetric_key_algorithm_sha1_authentication_key_plaintext   basetypes.StringValue
	services_ntp_servers_primary_ntp_server_authentication_type_symmetric_key_algorithm_sha1_authentication_key_encrypted   basetypes.StringValue
	services_ntp_servers_secondary_ntp_server_authentication_type_symmetric_key_algorithm_md5_authentication_key_plaintext  basetypes.StringValue
	services_ntp_servers_secondary_ntp_server_authentication_type_symmetric_key_algorithm_md5_authentication_key_encrypted  basetypes.StringValue
	services_ntp_servers_secondary_ntp_server_authentication_type_symmetric_key_algorithm_sha1_authentication_key_plaintext basetypes.StringValue
	services_ntp_servers_secondary_ntp_server_authentication_type_symmetric_key_algorithm_sha1_authentication_key_encrypted basetypes.StringValue
	services_secure_proxy_password_plaintext                                                                                basetypes.StringValue
	services_secure_proxy_password_encrypted                                                                                basetypes.StringValue
}

// populatePatcherFromState populates the patcher struct from the resource's state.
func (p *serviceSettingsSensitiveValuePatcher) populatePatcherFromState(ctx context.Context, state models.ServiceSettings) diag.Diagnostics {
	var diags diag.Diagnostics
	if state.EncryptedValues.IsNull() || state.EncryptedValues.IsUnknown() {
		return diags
	}

	var ev map[string]string
	diags.Append(state.EncryptedValues.ElementsAs(ctx, &ev, false)...)
	if diags.HasError() {
		return diags
	}
	if val, ok := ev["services_ntp_servers_primary_ntp_server_authentication_type_symmetric_key_algorithm_md5_authentication_key_plaintext"]; ok {
		p.services_ntp_servers_primary_ntp_server_authentication_type_symmetric_key_algorithm_md5_authentication_key_plaintext = basetypes.NewStringValue(val)
	}
	if val, ok := ev["services_ntp_servers_primary_ntp_server_authentication_type_symmetric_key_algorithm_md5_authentication_key_encrypted"]; ok {
		p.services_ntp_servers_primary_ntp_server_authentication_type_symmetric_key_algorithm_md5_authentication_key_encrypted = basetypes.NewStringValue(val)
	}
	if val, ok := ev["services_ntp_servers_primary_ntp_server_authentication_type_symmetric_key_algorithm_sha1_authentication_key_plaintext"]; ok {
		p.services_ntp_servers_primary_ntp_server_authentication_type_symmetric_key_algorithm_sha1_authentication_key_plaintext = basetypes.NewStringValue(val)
	}
	if val, ok := ev["services_ntp_servers_primary_ntp_server_authentication_type_symmetric_key_algorithm_sha1_authentication_key_encrypted"]; ok {
		p.services_ntp_servers_primary_ntp_server_authentication_type_symmetric_key_algorithm_sha1_authentication_key_encrypted = basetypes.NewStringValue(val)
	}
	if val, ok := ev["services_ntp_servers_secondary_ntp_server_authentication_type_symmetric_key_algorithm_md5_authentication_key_plaintext"]; ok {
		p.services_ntp_servers_secondary_ntp_server_authentication_type_symmetric_key_algorithm_md5_authentication_key_plaintext = basetypes.NewStringValue(val)
	}
	if val, ok := ev["services_ntp_servers_secondary_ntp_server_authentication_type_symmetric_key_algorithm_md5_authentication_key_encrypted"]; ok {
		p.services_ntp_servers_secondary_ntp_server_authentication_type_symmetric_key_algorithm_md5_authentication_key_encrypted = basetypes.NewStringValue(val)
	}
	if val, ok := ev["services_ntp_servers_secondary_ntp_server_authentication_type_symmetric_key_algorithm_sha1_authentication_key_plaintext"]; ok {
		p.services_ntp_servers_secondary_ntp_server_authentication_type_symmetric_key_algorithm_sha1_authentication_key_plaintext = basetypes.NewStringValue(val)
	}
	if val, ok := ev["services_ntp_servers_secondary_ntp_server_authentication_type_symmetric_key_algorithm_sha1_authentication_key_encrypted"]; ok {
		p.services_ntp_servers_secondary_ntp_server_authentication_type_symmetric_key_algorithm_sha1_authentication_key_encrypted = basetypes.NewStringValue(val)
	}
	if val, ok := ev["services_secure_proxy_password_plaintext"]; ok {
		p.services_secure_proxy_password_plaintext = basetypes.NewStringValue(val)
	}
	if val, ok := ev["services_secure_proxy_password_encrypted"]; ok {
		p.services_secure_proxy_password_encrypted = basetypes.NewStringValue(val)
	}

	return diags
}

// populateEncryptedValuesMap returns a map of the patcher's values for saving to state.
func (p *serviceSettingsSensitiveValuePatcher) populateEncryptedValuesMap() map[string]string {
	ev := make(map[string]string)
	if !p.services_ntp_servers_primary_ntp_server_authentication_type_symmetric_key_algorithm_md5_authentication_key_plaintext.IsNull() {
		ev["services_ntp_servers_primary_ntp_server_authentication_type_symmetric_key_algorithm_md5_authentication_key_plaintext"] = p.services_ntp_servers_primary_ntp_server_authentication_type_symmetric_key_algorithm_md5_authentication_key_plaintext.ValueString()
	}
	if !p.services_ntp_servers_primary_ntp_server_authentication_type_symmetric_key_algorithm_md5_authentication_key_encrypted.IsNull() {
		ev["services_ntp_servers_primary_ntp_server_authentication_type_symmetric_key_algorithm_md5_authentication_key_encrypted"] = p.services_ntp_servers_primary_ntp_server_authentication_type_symmetric_key_algorithm_md5_authentication_key_encrypted.ValueString()
	}
	if !p.services_ntp_servers_primary_ntp_server_authentication_type_symmetric_key_algorithm_sha1_authentication_key_plaintext.IsNull() {
		ev["services_ntp_servers_primary_ntp_server_authentication_type_symmetric_key_algorithm_sha1_authentication_key_plaintext"] = p.services_ntp_servers_primary_ntp_server_authentication_type_symmetric_key_algorithm_sha1_authentication_key_plaintext.ValueString()
	}
	if !p.services_ntp_servers_primary_ntp_server_authentication_type_symmetric_key_algorithm_sha1_authentication_key_encrypted.IsNull() {
		ev["services_ntp_servers_primary_ntp_server_authentication_type_symmetric_key_algorithm_sha1_authentication_key_encrypted"] = p.services_ntp_servers_primary_ntp_server_authentication_type_symmetric_key_algorithm_sha1_authentication_key_encrypted.ValueString()
	}
	if !p.services_ntp_servers_secondary_ntp_server_authentication_type_symmetric_key_algorithm_md5_authentication_key_plaintext.IsNull() {
		ev["services_ntp_servers_secondary_ntp_server_authentication_type_symmetric_key_algorithm_md5_authentication_key_plaintext"] = p.services_ntp_servers_secondary_ntp_server_authentication_type_symmetric_key_algorithm_md5_authentication_key_plaintext.ValueString()
	}
	if !p.services_ntp_servers_secondary_ntp_server_authentication_type_symmetric_key_algorithm_md5_authentication_key_encrypted.IsNull() {
		ev["services_ntp_servers_secondary_ntp_server_authentication_type_symmetric_key_algorithm_md5_authentication_key_encrypted"] = p.services_ntp_servers_secondary_ntp_server_authentication_type_symmetric_key_algorithm_md5_authentication_key_encrypted.ValueString()
	}
	if !p.services_ntp_servers_secondary_ntp_server_authentication_type_symmetric_key_algorithm_sha1_authentication_key_plaintext.IsNull() {
		ev["services_ntp_servers_secondary_ntp_server_authentication_type_symmetric_key_algorithm_sha1_authentication_key_plaintext"] = p.services_ntp_servers_secondary_ntp_server_authentication_type_symmetric_key_algorithm_sha1_authentication_key_plaintext.ValueString()
	}
	if !p.services_ntp_servers_secondary_ntp_server_authentication_type_symmetric_key_algorithm_sha1_authentication_key_encrypted.IsNull() {
		ev["services_ntp_servers_secondary_ntp_server_authentication_type_symmetric_key_algorithm_sha1_authentication_key_encrypted"] = p.services_ntp_servers_secondary_ntp_server_authentication_type_symmetric_key_algorithm_sha1_authentication_key_encrypted.ValueString()
	}
	if !p.services_secure_proxy_password_plaintext.IsNull() {
		ev["services_secure_proxy_password_plaintext"] = p.services_secure_proxy_password_plaintext.ValueString()
	}
	if !p.services_secure_proxy_password_encrypted.IsNull() {
		ev["services_secure_proxy_password_encrypted"] = p.services_secure_proxy_password_encrypted.ValueString()
	}
	return ev
}

// --- Unpacker for ServiceSettings ---
func unpackServiceSettingsToSdk(ctx context.Context, obj types.Object) (*device_settings.ServiceSettings, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ServiceSettings", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ServiceSettings
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk device_settings.ServiceSettings
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
	if !model.Services.IsNull() && !model.Services.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Services")
		unpacked, d := unpackServiceSettingsServicesToSdk(ctx, model.Services)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Services"})
		}
		if unpacked != nil {
			sdk.Services = unpacked
		}
	}

	// Handling Primitives
	if !model.Snippet.IsNull() && !model.Snippet.IsUnknown() {
		sdk.Snippet = model.Snippet.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ServiceSettings", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ServiceSettings ---
func packServiceSettingsFromSdk(ctx context.Context, sdk device_settings.ServiceSettings) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ServiceSettings", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ServiceSettings
	var d diag.Diagnostics
	// MEGA FIX FOR MAP TYPE MISMATCH (NOT ALL MODELS MAY HAVE EncryptedValues)
	model.EncryptedValues = basetypes.NewMapNull(basetypes.StringType{})
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
	if sdk.Services != nil {
		tflog.Debug(ctx, "Packing nested object for field Services")
		packed, d := packServiceSettingsServicesFromSdk(ctx, *sdk.Services)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Services"})
		}
		model.Services = packed
	} else {
		model.Services = basetypes.NewObjectNull(models.ServiceSettingsServices{}.AttrTypes())
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

	obj, d := types.ObjectValueFrom(ctx, models.ServiceSettings{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ServiceSettings", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ServiceSettings ---
func unpackServiceSettingsListToSdk(ctx context.Context, list types.List) ([]device_settings.ServiceSettings, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ServiceSettings")
	diags := diag.Diagnostics{}
	var data []models.ServiceSettings
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]device_settings.ServiceSettings, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ServiceSettings{}.AttrTypes(), &item)
		unpacked, d := unpackServiceSettingsToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ServiceSettings", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ServiceSettings ---
func packServiceSettingsListFromSdk(ctx context.Context, sdks []device_settings.ServiceSettings) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ServiceSettings")
	diags := diag.Diagnostics{}
	var data []models.ServiceSettings

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ServiceSettings
		obj, d := packServiceSettingsFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ServiceSettings{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ServiceSettings", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ServiceSettings{}.AttrType(), data)
}

// --- Unpacker for ServiceSettingsServices ---
func unpackServiceSettingsServicesToSdk(ctx context.Context, obj types.Object) (*device_settings.ServiceSettingsServices, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ServiceSettingsServices", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ServiceSettingsServices
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk device_settings.ServiceSettingsServices
	var d diag.Diagnostics
	// Handling Objects
	if !model.DnsSetting.IsNull() && !model.DnsSetting.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field DnsSetting")
		unpacked, d := unpackServiceSettingsServicesDnsSettingToSdk(ctx, model.DnsSetting)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "DnsSetting"})
		}
		if unpacked != nil {
			sdk.DnsSetting = unpacked
		}
	}

	// Handling Primitives
	if !model.FqdnRefreshTime.IsNull() && !model.FqdnRefreshTime.IsUnknown() {
		val := float32(model.FqdnRefreshTime.ValueFloat64())
		sdk.FqdnRefreshTime = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "FqdnRefreshTime", "value": *sdk.FqdnRefreshTime})
	}

	// Handling Primitives
	if !model.FqdnStaleEntryTimeout.IsNull() && !model.FqdnStaleEntryTimeout.IsUnknown() {
		val := float32(model.FqdnStaleEntryTimeout.ValueFloat64())
		sdk.FqdnStaleEntryTimeout = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "FqdnStaleEntryTimeout", "value": *sdk.FqdnStaleEntryTimeout})
	}

	// Handling Primitives
	if !model.InlineCloudProxy.IsNull() && !model.InlineCloudProxy.IsUnknown() {
		sdk.InlineCloudProxy = model.InlineCloudProxy.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "InlineCloudProxy", "value": *sdk.InlineCloudProxy})
	}

	// Handling Primitives
	if !model.LcaasUseProxy.IsNull() && !model.LcaasUseProxy.IsUnknown() {
		sdk.LcaasUseProxy = model.LcaasUseProxy.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "LcaasUseProxy", "value": *sdk.LcaasUseProxy})
	}

	// Handling Objects
	if !model.NtpServers.IsNull() && !model.NtpServers.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field NtpServers")
		unpacked, d := unpackServiceSettingsServicesNtpServersToSdk(ctx, model.NtpServers)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "NtpServers"})
		}
		if unpacked != nil {
			sdk.NtpServers = unpacked
		}
	}

	// Handling Primitives
	if !model.SecureProxyPassword.IsNull() && !model.SecureProxyPassword.IsUnknown() {
		sdk.SecureProxyPassword = model.SecureProxyPassword.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "SecureProxyPassword", "value": *sdk.SecureProxyPassword})
	}

	// Handling Primitives
	if !model.SecureProxyPort.IsNull() && !model.SecureProxyPort.IsUnknown() {
		val := float32(model.SecureProxyPort.ValueFloat64())
		sdk.SecureProxyPort = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "SecureProxyPort", "value": *sdk.SecureProxyPort})
	}

	// Handling Primitives
	if !model.SecureProxyServer.IsNull() && !model.SecureProxyServer.IsUnknown() {
		sdk.SecureProxyServer = model.SecureProxyServer.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "SecureProxyServer", "value": *sdk.SecureProxyServer})
	}

	// Handling Primitives
	if !model.SecureProxyUser.IsNull() && !model.SecureProxyUser.IsUnknown() {
		sdk.SecureProxyUser = model.SecureProxyUser.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "SecureProxyUser", "value": *sdk.SecureProxyUser})
	}

	// Handling Primitives
	if !model.ServerVerification.IsNull() && !model.ServerVerification.IsUnknown() {
		sdk.ServerVerification = model.ServerVerification.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "ServerVerification", "value": *sdk.ServerVerification})
	}

	// Handling Primitives
	if !model.UpdateServer.IsNull() && !model.UpdateServer.IsUnknown() {
		sdk.UpdateServer = model.UpdateServer.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "UpdateServer", "value": *sdk.UpdateServer})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ServiceSettingsServices", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ServiceSettingsServices ---
func packServiceSettingsServicesFromSdk(ctx context.Context, sdk device_settings.ServiceSettingsServices) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ServiceSettingsServices", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ServiceSettingsServices
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.DnsSetting != nil {
		tflog.Debug(ctx, "Packing nested object for field DnsSetting")
		packed, d := packServiceSettingsServicesDnsSettingFromSdk(ctx, *sdk.DnsSetting)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "DnsSetting"})
		}
		model.DnsSetting = packed
	} else {
		model.DnsSetting = basetypes.NewObjectNull(models.ServiceSettingsServicesDnsSetting{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.FqdnRefreshTime != nil {
		model.FqdnRefreshTime = basetypes.NewFloat64Value(float64(*sdk.FqdnRefreshTime))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "FqdnRefreshTime", "value": *sdk.FqdnRefreshTime})
	} else {
		model.FqdnRefreshTime = basetypes.NewFloat64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.FqdnStaleEntryTimeout != nil {
		model.FqdnStaleEntryTimeout = basetypes.NewFloat64Value(float64(*sdk.FqdnStaleEntryTimeout))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "FqdnStaleEntryTimeout", "value": *sdk.FqdnStaleEntryTimeout})
	} else {
		model.FqdnStaleEntryTimeout = basetypes.NewFloat64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.InlineCloudProxy != nil {
		model.InlineCloudProxy = basetypes.NewBoolValue(*sdk.InlineCloudProxy)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "InlineCloudProxy", "value": *sdk.InlineCloudProxy})
	} else {
		model.InlineCloudProxy = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.LcaasUseProxy != nil {
		model.LcaasUseProxy = basetypes.NewBoolValue(*sdk.LcaasUseProxy)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "LcaasUseProxy", "value": *sdk.LcaasUseProxy})
	} else {
		model.LcaasUseProxy = basetypes.NewBoolNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.NtpServers != nil {
		tflog.Debug(ctx, "Packing nested object for field NtpServers")
		packed, d := packServiceSettingsServicesNtpServersFromSdk(ctx, *sdk.NtpServers)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "NtpServers"})
		}
		model.NtpServers = packed
	} else {
		model.NtpServers = basetypes.NewObjectNull(models.ServiceSettingsServicesNtpServers{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.SecureProxyPassword != nil {
		model.SecureProxyPassword = basetypes.NewStringValue(*sdk.SecureProxyPassword)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "SecureProxyPassword", "value": *sdk.SecureProxyPassword})
	} else {
		model.SecureProxyPassword = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.SecureProxyPort != nil {
		model.SecureProxyPort = basetypes.NewFloat64Value(float64(*sdk.SecureProxyPort))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "SecureProxyPort", "value": *sdk.SecureProxyPort})
	} else {
		model.SecureProxyPort = basetypes.NewFloat64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.SecureProxyServer != nil {
		model.SecureProxyServer = basetypes.NewStringValue(*sdk.SecureProxyServer)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "SecureProxyServer", "value": *sdk.SecureProxyServer})
	} else {
		model.SecureProxyServer = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.SecureProxyUser != nil {
		model.SecureProxyUser = basetypes.NewStringValue(*sdk.SecureProxyUser)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "SecureProxyUser", "value": *sdk.SecureProxyUser})
	} else {
		model.SecureProxyUser = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.ServerVerification != nil {
		model.ServerVerification = basetypes.NewBoolValue(*sdk.ServerVerification)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "ServerVerification", "value": *sdk.ServerVerification})
	} else {
		model.ServerVerification = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.UpdateServer != nil {
		model.UpdateServer = basetypes.NewStringValue(*sdk.UpdateServer)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "UpdateServer", "value": *sdk.UpdateServer})
	} else {
		model.UpdateServer = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ServiceSettingsServices{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ServiceSettingsServices", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ServiceSettingsServices ---
func unpackServiceSettingsServicesListToSdk(ctx context.Context, list types.List) ([]device_settings.ServiceSettingsServices, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ServiceSettingsServices")
	diags := diag.Diagnostics{}
	var data []models.ServiceSettingsServices
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]device_settings.ServiceSettingsServices, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ServiceSettingsServices{}.AttrTypes(), &item)
		unpacked, d := unpackServiceSettingsServicesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ServiceSettingsServices", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ServiceSettingsServices ---
func packServiceSettingsServicesListFromSdk(ctx context.Context, sdks []device_settings.ServiceSettingsServices) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ServiceSettingsServices")
	diags := diag.Diagnostics{}
	var data []models.ServiceSettingsServices

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ServiceSettingsServices
		obj, d := packServiceSettingsServicesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ServiceSettingsServices{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ServiceSettingsServices", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ServiceSettingsServices{}.AttrType(), data)
}

// --- Unpacker for ServiceSettingsServicesDnsSetting ---
func unpackServiceSettingsServicesDnsSettingToSdk(ctx context.Context, obj types.Object) (*device_settings.ServiceSettingsServicesDnsSetting, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ServiceSettingsServicesDnsSetting", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ServiceSettingsServicesDnsSetting
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk device_settings.ServiceSettingsServicesDnsSetting
	var d diag.Diagnostics
	// Handling Primitives
	if !model.DnsProxyObject.IsNull() && !model.DnsProxyObject.IsUnknown() {
		sdk.DnsProxyObject = model.DnsProxyObject.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DnsProxyObject", "value": *sdk.DnsProxyObject})
	}

	// Handling Objects
	if !model.Servers.IsNull() && !model.Servers.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Servers")
		unpacked, d := unpackServiceSettingsServicesDnsSettingServersToSdk(ctx, model.Servers)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Servers"})
		}
		if unpacked != nil {
			sdk.Servers = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ServiceSettingsServicesDnsSetting", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ServiceSettingsServicesDnsSetting ---
func packServiceSettingsServicesDnsSettingFromSdk(ctx context.Context, sdk device_settings.ServiceSettingsServicesDnsSetting) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ServiceSettingsServicesDnsSetting", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ServiceSettingsServicesDnsSetting
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.DnsProxyObject != nil {
		model.DnsProxyObject = basetypes.NewStringValue(*sdk.DnsProxyObject)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DnsProxyObject", "value": *sdk.DnsProxyObject})
	} else {
		model.DnsProxyObject = basetypes.NewStringNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Servers != nil {
		tflog.Debug(ctx, "Packing nested object for field Servers")
		packed, d := packServiceSettingsServicesDnsSettingServersFromSdk(ctx, *sdk.Servers)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Servers"})
		}
		model.Servers = packed
	} else {
		model.Servers = basetypes.NewObjectNull(models.ServiceSettingsServicesDnsSettingServers{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ServiceSettingsServicesDnsSetting{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ServiceSettingsServicesDnsSetting", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ServiceSettingsServicesDnsSetting ---
func unpackServiceSettingsServicesDnsSettingListToSdk(ctx context.Context, list types.List) ([]device_settings.ServiceSettingsServicesDnsSetting, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ServiceSettingsServicesDnsSetting")
	diags := diag.Diagnostics{}
	var data []models.ServiceSettingsServicesDnsSetting
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]device_settings.ServiceSettingsServicesDnsSetting, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ServiceSettingsServicesDnsSetting{}.AttrTypes(), &item)
		unpacked, d := unpackServiceSettingsServicesDnsSettingToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ServiceSettingsServicesDnsSetting", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ServiceSettingsServicesDnsSetting ---
func packServiceSettingsServicesDnsSettingListFromSdk(ctx context.Context, sdks []device_settings.ServiceSettingsServicesDnsSetting) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ServiceSettingsServicesDnsSetting")
	diags := diag.Diagnostics{}
	var data []models.ServiceSettingsServicesDnsSetting

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ServiceSettingsServicesDnsSetting
		obj, d := packServiceSettingsServicesDnsSettingFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ServiceSettingsServicesDnsSetting{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ServiceSettingsServicesDnsSetting", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ServiceSettingsServicesDnsSetting{}.AttrType(), data)
}

// --- Unpacker for ServiceSettingsServicesDnsSettingServers ---
func unpackServiceSettingsServicesDnsSettingServersToSdk(ctx context.Context, obj types.Object) (*device_settings.ServiceSettingsServicesDnsSettingServers, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ServiceSettingsServicesDnsSettingServers", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ServiceSettingsServicesDnsSettingServers
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk device_settings.ServiceSettingsServicesDnsSettingServers
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Primary.IsNull() && !model.Primary.IsUnknown() {
		sdk.Primary = model.Primary.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Primary", "value": *sdk.Primary})
	}

	// Handling Primitives
	if !model.Secondary.IsNull() && !model.Secondary.IsUnknown() {
		sdk.Secondary = model.Secondary.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Secondary", "value": *sdk.Secondary})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ServiceSettingsServicesDnsSettingServers", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ServiceSettingsServicesDnsSettingServers ---
func packServiceSettingsServicesDnsSettingServersFromSdk(ctx context.Context, sdk device_settings.ServiceSettingsServicesDnsSettingServers) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ServiceSettingsServicesDnsSettingServers", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ServiceSettingsServicesDnsSettingServers
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Primary != nil {
		model.Primary = basetypes.NewStringValue(*sdk.Primary)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Primary", "value": *sdk.Primary})
	} else {
		model.Primary = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Secondary != nil {
		model.Secondary = basetypes.NewStringValue(*sdk.Secondary)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Secondary", "value": *sdk.Secondary})
	} else {
		model.Secondary = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ServiceSettingsServicesDnsSettingServers{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ServiceSettingsServicesDnsSettingServers", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ServiceSettingsServicesDnsSettingServers ---
func unpackServiceSettingsServicesDnsSettingServersListToSdk(ctx context.Context, list types.List) ([]device_settings.ServiceSettingsServicesDnsSettingServers, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ServiceSettingsServicesDnsSettingServers")
	diags := diag.Diagnostics{}
	var data []models.ServiceSettingsServicesDnsSettingServers
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]device_settings.ServiceSettingsServicesDnsSettingServers, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ServiceSettingsServicesDnsSettingServers{}.AttrTypes(), &item)
		unpacked, d := unpackServiceSettingsServicesDnsSettingServersToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ServiceSettingsServicesDnsSettingServers", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ServiceSettingsServicesDnsSettingServers ---
func packServiceSettingsServicesDnsSettingServersListFromSdk(ctx context.Context, sdks []device_settings.ServiceSettingsServicesDnsSettingServers) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ServiceSettingsServicesDnsSettingServers")
	diags := diag.Diagnostics{}
	var data []models.ServiceSettingsServicesDnsSettingServers

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ServiceSettingsServicesDnsSettingServers
		obj, d := packServiceSettingsServicesDnsSettingServersFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ServiceSettingsServicesDnsSettingServers{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ServiceSettingsServicesDnsSettingServers", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ServiceSettingsServicesDnsSettingServers{}.AttrType(), data)
}

// --- Unpacker for ServiceSettingsServicesNtpServers ---
func unpackServiceSettingsServicesNtpServersToSdk(ctx context.Context, obj types.Object) (*device_settings.ServiceSettingsServicesNtpServers, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ServiceSettingsServicesNtpServers", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ServiceSettingsServicesNtpServers
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk device_settings.ServiceSettingsServicesNtpServers
	var d diag.Diagnostics
	// Handling Objects
	if !model.PrimaryNtpServer.IsNull() && !model.PrimaryNtpServer.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field PrimaryNtpServer")
		unpacked, d := unpackServiceSettingsServicesNtpServersPrimaryNtpServerToSdk(ctx, model.PrimaryNtpServer)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "PrimaryNtpServer"})
		}
		if unpacked != nil {
			sdk.PrimaryNtpServer = unpacked
		}
	}

	// Handling Objects
	if !model.SecondaryNtpServer.IsNull() && !model.SecondaryNtpServer.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field SecondaryNtpServer")
		unpacked, d := unpackServiceSettingsServicesNtpServersPrimaryNtpServerToSdk(ctx, model.SecondaryNtpServer)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "SecondaryNtpServer"})
		}
		if unpacked != nil {
			sdk.SecondaryNtpServer = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ServiceSettingsServicesNtpServers", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ServiceSettingsServicesNtpServers ---
func packServiceSettingsServicesNtpServersFromSdk(ctx context.Context, sdk device_settings.ServiceSettingsServicesNtpServers) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ServiceSettingsServicesNtpServers", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ServiceSettingsServicesNtpServers
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.PrimaryNtpServer != nil {
		tflog.Debug(ctx, "Packing nested object for field PrimaryNtpServer")
		packed, d := packServiceSettingsServicesNtpServersPrimaryNtpServerFromSdk(ctx, *sdk.PrimaryNtpServer)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "PrimaryNtpServer"})
		}
		model.PrimaryNtpServer = packed
	} else {
		model.PrimaryNtpServer = basetypes.NewObjectNull(models.ServiceSettingsServicesNtpServersPrimaryNtpServer{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.SecondaryNtpServer != nil {
		tflog.Debug(ctx, "Packing nested object for field SecondaryNtpServer")
		packed, d := packServiceSettingsServicesNtpServersPrimaryNtpServerFromSdk(ctx, *sdk.SecondaryNtpServer)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "SecondaryNtpServer"})
		}
		model.SecondaryNtpServer = packed
	} else {
		model.SecondaryNtpServer = basetypes.NewObjectNull(models.ServiceSettingsServicesNtpServersPrimaryNtpServer{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServers{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ServiceSettingsServicesNtpServers", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ServiceSettingsServicesNtpServers ---
func unpackServiceSettingsServicesNtpServersListToSdk(ctx context.Context, list types.List) ([]device_settings.ServiceSettingsServicesNtpServers, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ServiceSettingsServicesNtpServers")
	diags := diag.Diagnostics{}
	var data []models.ServiceSettingsServicesNtpServers
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]device_settings.ServiceSettingsServicesNtpServers, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServers{}.AttrTypes(), &item)
		unpacked, d := unpackServiceSettingsServicesNtpServersToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ServiceSettingsServicesNtpServers", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ServiceSettingsServicesNtpServers ---
func packServiceSettingsServicesNtpServersListFromSdk(ctx context.Context, sdks []device_settings.ServiceSettingsServicesNtpServers) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ServiceSettingsServicesNtpServers")
	diags := diag.Diagnostics{}
	var data []models.ServiceSettingsServicesNtpServers

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ServiceSettingsServicesNtpServers
		obj, d := packServiceSettingsServicesNtpServersFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ServiceSettingsServicesNtpServers{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ServiceSettingsServicesNtpServers", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ServiceSettingsServicesNtpServers{}.AttrType(), data)
}

// --- Unpacker for ServiceSettingsServicesNtpServersPrimaryNtpServer ---
func unpackServiceSettingsServicesNtpServersPrimaryNtpServerToSdk(ctx context.Context, obj types.Object) (*device_settings.ServiceSettingsServicesNtpServersPrimaryNtpServer, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ServiceSettingsServicesNtpServersPrimaryNtpServer", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ServiceSettingsServicesNtpServersPrimaryNtpServer
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk device_settings.ServiceSettingsServicesNtpServersPrimaryNtpServer
	var d diag.Diagnostics
	// Handling Objects
	if !model.AuthenticationType.IsNull() && !model.AuthenticationType.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field AuthenticationType")
		unpacked, d := unpackServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeToSdk(ctx, model.AuthenticationType)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "AuthenticationType"})
		}
		if unpacked != nil {
			sdk.AuthenticationType = unpacked
		}
	}

	// Handling Primitives
	if !model.NtpServerAddress.IsNull() && !model.NtpServerAddress.IsUnknown() {
		sdk.NtpServerAddress = model.NtpServerAddress.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "NtpServerAddress", "value": *sdk.NtpServerAddress})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ServiceSettingsServicesNtpServersPrimaryNtpServer", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ServiceSettingsServicesNtpServersPrimaryNtpServer ---
func packServiceSettingsServicesNtpServersPrimaryNtpServerFromSdk(ctx context.Context, sdk device_settings.ServiceSettingsServicesNtpServersPrimaryNtpServer) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ServiceSettingsServicesNtpServersPrimaryNtpServer", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ServiceSettingsServicesNtpServersPrimaryNtpServer
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.AuthenticationType != nil {
		tflog.Debug(ctx, "Packing nested object for field AuthenticationType")
		packed, d := packServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeFromSdk(ctx, *sdk.AuthenticationType)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "AuthenticationType"})
		}
		model.AuthenticationType = packed
	} else {
		model.AuthenticationType = basetypes.NewObjectNull(models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationType{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.NtpServerAddress != nil {
		model.NtpServerAddress = basetypes.NewStringValue(*sdk.NtpServerAddress)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "NtpServerAddress", "value": *sdk.NtpServerAddress})
	} else {
		model.NtpServerAddress = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServer{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ServiceSettingsServicesNtpServersPrimaryNtpServer", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ServiceSettingsServicesNtpServersPrimaryNtpServer ---
func unpackServiceSettingsServicesNtpServersPrimaryNtpServerListToSdk(ctx context.Context, list types.List) ([]device_settings.ServiceSettingsServicesNtpServersPrimaryNtpServer, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ServiceSettingsServicesNtpServersPrimaryNtpServer")
	diags := diag.Diagnostics{}
	var data []models.ServiceSettingsServicesNtpServersPrimaryNtpServer
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]device_settings.ServiceSettingsServicesNtpServersPrimaryNtpServer, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServer{}.AttrTypes(), &item)
		unpacked, d := unpackServiceSettingsServicesNtpServersPrimaryNtpServerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ServiceSettingsServicesNtpServersPrimaryNtpServer", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ServiceSettingsServicesNtpServersPrimaryNtpServer ---
func packServiceSettingsServicesNtpServersPrimaryNtpServerListFromSdk(ctx context.Context, sdks []device_settings.ServiceSettingsServicesNtpServersPrimaryNtpServer) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ServiceSettingsServicesNtpServersPrimaryNtpServer")
	diags := diag.Diagnostics{}
	var data []models.ServiceSettingsServicesNtpServersPrimaryNtpServer

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ServiceSettingsServicesNtpServersPrimaryNtpServer
		obj, d := packServiceSettingsServicesNtpServersPrimaryNtpServerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ServiceSettingsServicesNtpServersPrimaryNtpServer{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ServiceSettingsServicesNtpServersPrimaryNtpServer", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServer{}.AttrType(), data)
}

// --- Unpacker for ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationType ---
func unpackServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeToSdk(ctx context.Context, obj types.Object) (*device_settings.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationType, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationType", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationType
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk device_settings.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationType
	var d diag.Diagnostics
	// Handling Typeless Objects
	if !model.Autokey.IsNull() && !model.Autokey.IsUnknown() {
		tflog.Debug(ctx, "Unpacking typeless object for field Autokey")
		sdk.Autokey = make(map[string]interface{})
	}

	// Handling Typeless Objects
	if !model.None.IsNull() && !model.None.IsUnknown() {
		tflog.Debug(ctx, "Unpacking typeless object for field None")
		sdk.None = make(map[string]interface{})
	}

	// Handling Objects
	if !model.SymmetricKey.IsNull() && !model.SymmetricKey.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field SymmetricKey")
		unpacked, d := unpackServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyToSdk(ctx, model.SymmetricKey)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "SymmetricKey"})
		}
		if unpacked != nil {
			sdk.SymmetricKey = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationType", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationType ---
func packServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeFromSdk(ctx context.Context, sdk device_settings.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationType) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationType", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationType
	var d diag.Diagnostics
	// Handling Objects
	// This is a marker object (e.g. CHAP: {}). We just need to create an empty, non-null object.
	if sdk.Autokey != nil && !reflect.ValueOf(sdk.Autokey).IsNil() {
		tflog.Debug(ctx, "Packing typeless object for field Autokey")
		var d diag.Diagnostics
		// Create an empty object with no attributes, which signifies its presence.
		model.Autokey, d = basetypes.NewObjectValue(map[string]attr.Type{}, map[string]attr.Value{})
		diags.Append(d...)
	} else {
		// Since this field is part of a oneOf, being nil means it's not selected.
		// We make the object null with an empty attribute map.
		model.Autokey = basetypes.NewObjectNull(map[string]attr.Type{})
	}
	// Handling Objects
	// This is a marker object (e.g. CHAP: {}). We just need to create an empty, non-null object.
	if sdk.None != nil && !reflect.ValueOf(sdk.None).IsNil() {
		tflog.Debug(ctx, "Packing typeless object for field None")
		var d diag.Diagnostics
		// Create an empty object with no attributes, which signifies its presence.
		model.None, d = basetypes.NewObjectValue(map[string]attr.Type{}, map[string]attr.Value{})
		diags.Append(d...)
	} else {
		// Since this field is part of a oneOf, being nil means it's not selected.
		// We make the object null with an empty attribute map.
		model.None = basetypes.NewObjectNull(map[string]attr.Type{})
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.SymmetricKey != nil {
		tflog.Debug(ctx, "Packing nested object for field SymmetricKey")
		packed, d := packServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyFromSdk(ctx, *sdk.SymmetricKey)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "SymmetricKey"})
		}
		model.SymmetricKey = packed
	} else {
		model.SymmetricKey = basetypes.NewObjectNull(models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationType{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationType", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationType ---
func unpackServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeListToSdk(ctx context.Context, list types.List) ([]device_settings.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationType, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationType")
	diags := diag.Diagnostics{}
	var data []models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationType
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]device_settings.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationType, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationType{}.AttrTypes(), &item)
		unpacked, d := unpackServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationType", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationType ---
func packServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeListFromSdk(ctx context.Context, sdks []device_settings.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationType) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationType")
	diags := diag.Diagnostics{}
	var data []models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationType

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationType
		obj, d := packServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationType{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationType", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationType{}.AttrType(), data)
}

// --- Unpacker for ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey ---
func unpackServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyToSdk(ctx context.Context, obj types.Object) (*device_settings.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk device_settings.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey
	var d diag.Diagnostics
	// Handling Objects
	if !model.Algorithm.IsNull() && !model.Algorithm.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Algorithm")
		unpacked, d := unpackServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmToSdk(ctx, model.Algorithm)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Algorithm"})
		}
		if unpacked != nil {
			sdk.Algorithm = unpacked
		}
	}

	// Handling Primitives
	if !model.KeyId.IsNull() && !model.KeyId.IsUnknown() {
		val := float32(model.KeyId.ValueFloat64())
		sdk.KeyId = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "KeyId", "value": *sdk.KeyId})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey ---
func packServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyFromSdk(ctx context.Context, sdk device_settings.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Algorithm != nil {
		tflog.Debug(ctx, "Packing nested object for field Algorithm")
		packed, d := packServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmFromSdk(ctx, *sdk.Algorithm)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Algorithm"})
		}
		model.Algorithm = packed
	} else {
		model.Algorithm = basetypes.NewObjectNull(models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.KeyId != nil {
		model.KeyId = basetypes.NewFloat64Value(float64(*sdk.KeyId))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "KeyId", "value": *sdk.KeyId})
	} else {
		model.KeyId = basetypes.NewFloat64Null()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey ---
func unpackServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyListToSdk(ctx context.Context, list types.List) ([]device_settings.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey")
	diags := diag.Diagnostics{}
	var data []models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]device_settings.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey{}.AttrTypes(), &item)
		unpacked, d := unpackServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey ---
func packServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyListFromSdk(ctx context.Context, sdks []device_settings.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey")
	diags := diag.Diagnostics{}
	var data []models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey
		obj, d := packServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey{}.AttrType(), data)
}

// --- Unpacker for ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm ---
func unpackServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmToSdk(ctx context.Context, obj types.Object) (*device_settings.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk device_settings.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm
	var d diag.Diagnostics
	// Handling Objects
	if !model.Md5.IsNull() && !model.Md5.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Md5")
		unpacked, d := unpackServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5ToSdk(ctx, model.Md5)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Md5"})
		}
		if unpacked != nil {
			sdk.Md5 = unpacked
		}
	}

	// Handling Objects
	if !model.Sha1.IsNull() && !model.Sha1.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Sha1")
		unpacked, d := unpackServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5ToSdk(ctx, model.Sha1)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Sha1"})
		}
		if unpacked != nil {
			sdk.Sha1 = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm ---
func packServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmFromSdk(ctx context.Context, sdk device_settings.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Md5 != nil {
		tflog.Debug(ctx, "Packing nested object for field Md5")
		packed, d := packServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5FromSdk(ctx, *sdk.Md5)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Md5"})
		}
		model.Md5 = packed
	} else {
		model.Md5 = basetypes.NewObjectNull(models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Sha1 != nil {
		tflog.Debug(ctx, "Packing nested object for field Sha1")
		packed, d := packServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5FromSdk(ctx, *sdk.Sha1)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Sha1"})
		}
		model.Sha1 = packed
	} else {
		model.Sha1 = basetypes.NewObjectNull(models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm ---
func unpackServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmListToSdk(ctx context.Context, list types.List) ([]device_settings.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm")
	diags := diag.Diagnostics{}
	var data []models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]device_settings.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm{}.AttrTypes(), &item)
		unpacked, d := unpackServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm ---
func packServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmListFromSdk(ctx context.Context, sdks []device_settings.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm")
	diags := diag.Diagnostics{}
	var data []models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm
		obj, d := packServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm{}.AttrType(), data)
}

// --- Unpacker for ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5 ---
func unpackServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5ToSdk(ctx context.Context, obj types.Object) (*device_settings.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk device_settings.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5
	var d diag.Diagnostics
	// Handling Primitives
	if !model.AuthenticationKey.IsNull() && !model.AuthenticationKey.IsUnknown() {
		sdk.AuthenticationKey = model.AuthenticationKey.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "AuthenticationKey", "value": *sdk.AuthenticationKey})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5 ---
func packServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5FromSdk(ctx context.Context, sdk device_settings.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.AuthenticationKey != nil {
		model.AuthenticationKey = basetypes.NewStringValue(*sdk.AuthenticationKey)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "AuthenticationKey", "value": *sdk.AuthenticationKey})
	} else {
		model.AuthenticationKey = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5 ---
func unpackServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5ListToSdk(ctx context.Context, list types.List) ([]device_settings.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5")
	diags := diag.Diagnostics{}
	var data []models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]device_settings.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5{}.AttrTypes(), &item)
		unpacked, d := unpackServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5ToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5 ---
func packServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5ListFromSdk(ctx context.Context, sdks []device_settings.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5")
	diags := diag.Diagnostics{}
	var data []models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5
		obj, d := packServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5FromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5{}.AttrType(), data)
}
