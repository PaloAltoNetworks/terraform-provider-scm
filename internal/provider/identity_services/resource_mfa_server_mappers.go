package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/identity_services"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/identity_services"
)

// mfaServersSensitiveValuePatcher is an in-memory struct to temporarily store plaintext
// and encrypted values for sensitive fields during the Create/Update/Read workflows.
type mfaServersSensitiveValuePatcher struct {
	mfa_vendor_type_duo_security_v2_duo_secret_key_plaintext       basetypes.StringValue
	mfa_vendor_type_duo_security_v2_duo_secret_key_encrypted       basetypes.StringValue
	mfa_vendor_type_okta_adaptive_v1_okta_token_plaintext          basetypes.StringValue
	mfa_vendor_type_okta_adaptive_v1_okta_token_encrypted          basetypes.StringValue
	mfa_vendor_type_ping_identity_v1_ping_use_base64_key_plaintext basetypes.StringValue
	mfa_vendor_type_ping_identity_v1_ping_use_base64_key_encrypted basetypes.StringValue
	mfa_vendor_type_rsa_securid_access_v1_rsa_accesskey_plaintext  basetypes.StringValue
	mfa_vendor_type_rsa_securid_access_v1_rsa_accesskey_encrypted  basetypes.StringValue
}

// populatePatcherFromState populates the patcher struct from the resource's state.
func (p *mfaServersSensitiveValuePatcher) populatePatcherFromState(ctx context.Context, state models.MfaServers) diag.Diagnostics {
	var diags diag.Diagnostics
	if state.EncryptedValues.IsNull() || state.EncryptedValues.IsUnknown() {
		return diags
	}

	var ev map[string]string
	diags.Append(state.EncryptedValues.ElementsAs(ctx, &ev, false)...)
	if diags.HasError() {
		return diags
	}
	if val, ok := ev["mfa_vendor_type_duo_security_v2_duo_secret_key_plaintext"]; ok {
		p.mfa_vendor_type_duo_security_v2_duo_secret_key_plaintext = basetypes.NewStringValue(val)
	}
	if val, ok := ev["mfa_vendor_type_duo_security_v2_duo_secret_key_encrypted"]; ok {
		p.mfa_vendor_type_duo_security_v2_duo_secret_key_encrypted = basetypes.NewStringValue(val)
	}
	if val, ok := ev["mfa_vendor_type_okta_adaptive_v1_okta_token_plaintext"]; ok {
		p.mfa_vendor_type_okta_adaptive_v1_okta_token_plaintext = basetypes.NewStringValue(val)
	}
	if val, ok := ev["mfa_vendor_type_okta_adaptive_v1_okta_token_encrypted"]; ok {
		p.mfa_vendor_type_okta_adaptive_v1_okta_token_encrypted = basetypes.NewStringValue(val)
	}
	if val, ok := ev["mfa_vendor_type_ping_identity_v1_ping_use_base64_key_plaintext"]; ok {
		p.mfa_vendor_type_ping_identity_v1_ping_use_base64_key_plaintext = basetypes.NewStringValue(val)
	}
	if val, ok := ev["mfa_vendor_type_ping_identity_v1_ping_use_base64_key_encrypted"]; ok {
		p.mfa_vendor_type_ping_identity_v1_ping_use_base64_key_encrypted = basetypes.NewStringValue(val)
	}
	if val, ok := ev["mfa_vendor_type_rsa_securid_access_v1_rsa_accesskey_plaintext"]; ok {
		p.mfa_vendor_type_rsa_securid_access_v1_rsa_accesskey_plaintext = basetypes.NewStringValue(val)
	}
	if val, ok := ev["mfa_vendor_type_rsa_securid_access_v1_rsa_accesskey_encrypted"]; ok {
		p.mfa_vendor_type_rsa_securid_access_v1_rsa_accesskey_encrypted = basetypes.NewStringValue(val)
	}

	return diags
}

// populateEncryptedValuesMap returns a map of the patcher's values for saving to state.
func (p *mfaServersSensitiveValuePatcher) populateEncryptedValuesMap() map[string]string {
	ev := make(map[string]string)
	if !p.mfa_vendor_type_duo_security_v2_duo_secret_key_plaintext.IsNull() {
		ev["mfa_vendor_type_duo_security_v2_duo_secret_key_plaintext"] = p.mfa_vendor_type_duo_security_v2_duo_secret_key_plaintext.ValueString()
	}
	if !p.mfa_vendor_type_duo_security_v2_duo_secret_key_encrypted.IsNull() {
		ev["mfa_vendor_type_duo_security_v2_duo_secret_key_encrypted"] = p.mfa_vendor_type_duo_security_v2_duo_secret_key_encrypted.ValueString()
	}
	if !p.mfa_vendor_type_okta_adaptive_v1_okta_token_plaintext.IsNull() {
		ev["mfa_vendor_type_okta_adaptive_v1_okta_token_plaintext"] = p.mfa_vendor_type_okta_adaptive_v1_okta_token_plaintext.ValueString()
	}
	if !p.mfa_vendor_type_okta_adaptive_v1_okta_token_encrypted.IsNull() {
		ev["mfa_vendor_type_okta_adaptive_v1_okta_token_encrypted"] = p.mfa_vendor_type_okta_adaptive_v1_okta_token_encrypted.ValueString()
	}
	if !p.mfa_vendor_type_ping_identity_v1_ping_use_base64_key_plaintext.IsNull() {
		ev["mfa_vendor_type_ping_identity_v1_ping_use_base64_key_plaintext"] = p.mfa_vendor_type_ping_identity_v1_ping_use_base64_key_plaintext.ValueString()
	}
	if !p.mfa_vendor_type_ping_identity_v1_ping_use_base64_key_encrypted.IsNull() {
		ev["mfa_vendor_type_ping_identity_v1_ping_use_base64_key_encrypted"] = p.mfa_vendor_type_ping_identity_v1_ping_use_base64_key_encrypted.ValueString()
	}
	if !p.mfa_vendor_type_rsa_securid_access_v1_rsa_accesskey_plaintext.IsNull() {
		ev["mfa_vendor_type_rsa_securid_access_v1_rsa_accesskey_plaintext"] = p.mfa_vendor_type_rsa_securid_access_v1_rsa_accesskey_plaintext.ValueString()
	}
	if !p.mfa_vendor_type_rsa_securid_access_v1_rsa_accesskey_encrypted.IsNull() {
		ev["mfa_vendor_type_rsa_securid_access_v1_rsa_accesskey_encrypted"] = p.mfa_vendor_type_rsa_securid_access_v1_rsa_accesskey_encrypted.ValueString()
	}
	return ev
}

// --- Unpacker for MfaServers ---
func unpackMfaServersToSdk(ctx context.Context, obj types.Object) (*identity_services.MfaServers, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.MfaServers", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.MfaServers
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk identity_services.MfaServers
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
	if !model.MfaCertProfile.IsNull() && !model.MfaCertProfile.IsUnknown() {
		sdk.MfaCertProfile = model.MfaCertProfile.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "MfaCertProfile", "value": sdk.MfaCertProfile})
	}

	// Handling Objects
	if !model.MfaVendorType.IsNull() && !model.MfaVendorType.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field MfaVendorType")
		unpacked, d := unpackMfaServersMfaVendorTypeToSdk(ctx, model.MfaVendorType)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "MfaVendorType"})
		}
		if unpacked != nil {
			sdk.MfaVendorType = unpacked
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

	tflog.Debug(ctx, "Exiting unpack helper for models.MfaServers", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for MfaServers ---
func packMfaServersFromSdk(ctx context.Context, sdk identity_services.MfaServers) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.MfaServers", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.MfaServers
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
	model.Id = basetypes.NewStringValue(sdk.Id)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Id", "value": sdk.Id})
	// Handling Primitives
	// Standard primitive packing
	model.MfaCertProfile = basetypes.NewStringValue(sdk.MfaCertProfile)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "MfaCertProfile", "value": sdk.MfaCertProfile})
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.MfaVendorType != nil {
		tflog.Debug(ctx, "Packing nested object for field MfaVendorType")
		packed, d := packMfaServersMfaVendorTypeFromSdk(ctx, *sdk.MfaVendorType)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "MfaVendorType"})
		}
		model.MfaVendorType = packed
	} else {
		model.MfaVendorType = basetypes.NewObjectNull(models.MfaServersMfaVendorType{}.AttrTypes())
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

	obj, d := types.ObjectValueFrom(ctx, models.MfaServers{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.MfaServers", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for MfaServers ---
func unpackMfaServersListToSdk(ctx context.Context, list types.List) ([]identity_services.MfaServers, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.MfaServers")
	diags := diag.Diagnostics{}
	var data []models.MfaServers
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]identity_services.MfaServers, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.MfaServers{}.AttrTypes(), &item)
		unpacked, d := unpackMfaServersToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.MfaServers", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for MfaServers ---
func packMfaServersListFromSdk(ctx context.Context, sdks []identity_services.MfaServers) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.MfaServers")
	diags := diag.Diagnostics{}
	var data []models.MfaServers

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.MfaServers
		obj, d := packMfaServersFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.MfaServers{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.MfaServers", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.MfaServers{}.AttrType(), data)
}

// --- Unpacker for MfaServersMfaVendorType ---
func unpackMfaServersMfaVendorTypeToSdk(ctx context.Context, obj types.Object) (*identity_services.MfaServersMfaVendorType, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.MfaServersMfaVendorType", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.MfaServersMfaVendorType
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk identity_services.MfaServersMfaVendorType
	var d diag.Diagnostics
	// Handling Objects
	if !model.DuoSecurityV2.IsNull() && !model.DuoSecurityV2.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field DuoSecurityV2")
		unpacked, d := unpackMfaServersMfaVendorTypeDuoSecurityV2ToSdk(ctx, model.DuoSecurityV2)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "DuoSecurityV2"})
		}
		if unpacked != nil {
			sdk.DuoSecurityV2 = unpacked
		}
	}

	// Handling Objects
	if !model.OktaAdaptiveV1.IsNull() && !model.OktaAdaptiveV1.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field OktaAdaptiveV1")
		unpacked, d := unpackMfaServersMfaVendorTypeOktaAdaptiveV1ToSdk(ctx, model.OktaAdaptiveV1)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "OktaAdaptiveV1"})
		}
		if unpacked != nil {
			sdk.OktaAdaptiveV1 = unpacked
		}
	}

	// Handling Objects
	if !model.PingIdentityV1.IsNull() && !model.PingIdentityV1.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field PingIdentityV1")
		unpacked, d := unpackMfaServersMfaVendorTypePingIdentityV1ToSdk(ctx, model.PingIdentityV1)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "PingIdentityV1"})
		}
		if unpacked != nil {
			sdk.PingIdentityV1 = unpacked
		}
	}

	// Handling Objects
	if !model.RsaSecuridAccessV1.IsNull() && !model.RsaSecuridAccessV1.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field RsaSecuridAccessV1")
		unpacked, d := unpackMfaServersMfaVendorTypeRsaSecuridAccessV1ToSdk(ctx, model.RsaSecuridAccessV1)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "RsaSecuridAccessV1"})
		}
		if unpacked != nil {
			sdk.RsaSecuridAccessV1 = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.MfaServersMfaVendorType", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for MfaServersMfaVendorType ---
func packMfaServersMfaVendorTypeFromSdk(ctx context.Context, sdk identity_services.MfaServersMfaVendorType) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.MfaServersMfaVendorType", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.MfaServersMfaVendorType
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.DuoSecurityV2 != nil {
		tflog.Debug(ctx, "Packing nested object for field DuoSecurityV2")
		packed, d := packMfaServersMfaVendorTypeDuoSecurityV2FromSdk(ctx, *sdk.DuoSecurityV2)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "DuoSecurityV2"})
		}
		model.DuoSecurityV2 = packed
	} else {
		model.DuoSecurityV2 = basetypes.NewObjectNull(models.MfaServersMfaVendorTypeDuoSecurityV2{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.OktaAdaptiveV1 != nil {
		tflog.Debug(ctx, "Packing nested object for field OktaAdaptiveV1")
		packed, d := packMfaServersMfaVendorTypeOktaAdaptiveV1FromSdk(ctx, *sdk.OktaAdaptiveV1)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "OktaAdaptiveV1"})
		}
		model.OktaAdaptiveV1 = packed
	} else {
		model.OktaAdaptiveV1 = basetypes.NewObjectNull(models.MfaServersMfaVendorTypeOktaAdaptiveV1{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.PingIdentityV1 != nil {
		tflog.Debug(ctx, "Packing nested object for field PingIdentityV1")
		packed, d := packMfaServersMfaVendorTypePingIdentityV1FromSdk(ctx, *sdk.PingIdentityV1)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "PingIdentityV1"})
		}
		model.PingIdentityV1 = packed
	} else {
		model.PingIdentityV1 = basetypes.NewObjectNull(models.MfaServersMfaVendorTypePingIdentityV1{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.RsaSecuridAccessV1 != nil {
		tflog.Debug(ctx, "Packing nested object for field RsaSecuridAccessV1")
		packed, d := packMfaServersMfaVendorTypeRsaSecuridAccessV1FromSdk(ctx, *sdk.RsaSecuridAccessV1)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "RsaSecuridAccessV1"})
		}
		model.RsaSecuridAccessV1 = packed
	} else {
		model.RsaSecuridAccessV1 = basetypes.NewObjectNull(models.MfaServersMfaVendorTypeRsaSecuridAccessV1{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.MfaServersMfaVendorType{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.MfaServersMfaVendorType", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for MfaServersMfaVendorType ---
func unpackMfaServersMfaVendorTypeListToSdk(ctx context.Context, list types.List) ([]identity_services.MfaServersMfaVendorType, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.MfaServersMfaVendorType")
	diags := diag.Diagnostics{}
	var data []models.MfaServersMfaVendorType
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]identity_services.MfaServersMfaVendorType, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.MfaServersMfaVendorType{}.AttrTypes(), &item)
		unpacked, d := unpackMfaServersMfaVendorTypeToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.MfaServersMfaVendorType", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for MfaServersMfaVendorType ---
func packMfaServersMfaVendorTypeListFromSdk(ctx context.Context, sdks []identity_services.MfaServersMfaVendorType) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.MfaServersMfaVendorType")
	diags := diag.Diagnostics{}
	var data []models.MfaServersMfaVendorType

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.MfaServersMfaVendorType
		obj, d := packMfaServersMfaVendorTypeFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.MfaServersMfaVendorType{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.MfaServersMfaVendorType", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.MfaServersMfaVendorType{}.AttrType(), data)
}

// --- Unpacker for MfaServersMfaVendorTypeDuoSecurityV2 ---
func unpackMfaServersMfaVendorTypeDuoSecurityV2ToSdk(ctx context.Context, obj types.Object) (*identity_services.MfaServersMfaVendorTypeDuoSecurityV2, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.MfaServersMfaVendorTypeDuoSecurityV2", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.MfaServersMfaVendorTypeDuoSecurityV2
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk identity_services.MfaServersMfaVendorTypeDuoSecurityV2
	var d diag.Diagnostics
	// Handling Primitives
	if !model.DuoApiHost.IsNull() && !model.DuoApiHost.IsUnknown() {
		sdk.DuoApiHost = model.DuoApiHost.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "DuoApiHost", "value": sdk.DuoApiHost})
	}

	// Handling Primitives
	if !model.DuoBaseuri.IsNull() && !model.DuoBaseuri.IsUnknown() {
		sdk.DuoBaseuri = model.DuoBaseuri.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "DuoBaseuri", "value": sdk.DuoBaseuri})
	}

	// Handling Primitives
	if !model.DuoIntegrationKey.IsNull() && !model.DuoIntegrationKey.IsUnknown() {
		sdk.DuoIntegrationKey = model.DuoIntegrationKey.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "DuoIntegrationKey", "value": sdk.DuoIntegrationKey})
	}

	// Handling Primitives
	if !model.DuoSecretKey.IsNull() && !model.DuoSecretKey.IsUnknown() {
		sdk.DuoSecretKey = model.DuoSecretKey.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "DuoSecretKey", "value": sdk.DuoSecretKey})
	}

	// Handling Primitives
	if !model.DuoTimeout.IsNull() && !model.DuoTimeout.IsUnknown() {
		sdk.DuoTimeout = int32(model.DuoTimeout.ValueInt64())
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "DuoTimeout", "value": sdk.DuoTimeout})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.MfaServersMfaVendorTypeDuoSecurityV2", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for MfaServersMfaVendorTypeDuoSecurityV2 ---
func packMfaServersMfaVendorTypeDuoSecurityV2FromSdk(ctx context.Context, sdk identity_services.MfaServersMfaVendorTypeDuoSecurityV2) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.MfaServersMfaVendorTypeDuoSecurityV2", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.MfaServersMfaVendorTypeDuoSecurityV2
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.DuoApiHost = basetypes.NewStringValue(sdk.DuoApiHost)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "DuoApiHost", "value": sdk.DuoApiHost})
	// Handling Primitives
	// Standard primitive packing
	model.DuoBaseuri = basetypes.NewStringValue(sdk.DuoBaseuri)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "DuoBaseuri", "value": sdk.DuoBaseuri})
	// Handling Primitives
	// Standard primitive packing
	model.DuoIntegrationKey = basetypes.NewStringValue(sdk.DuoIntegrationKey)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "DuoIntegrationKey", "value": sdk.DuoIntegrationKey})
	// Handling Primitives
	// Standard primitive packing
	model.DuoSecretKey = basetypes.NewStringValue(sdk.DuoSecretKey)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "DuoSecretKey", "value": sdk.DuoSecretKey})
	// Handling Primitives
	// Standard primitive packing
	model.DuoTimeout = basetypes.NewInt64Value(int64(sdk.DuoTimeout))
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "DuoTimeout", "value": sdk.DuoTimeout})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.MfaServersMfaVendorTypeDuoSecurityV2{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.MfaServersMfaVendorTypeDuoSecurityV2", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for MfaServersMfaVendorTypeDuoSecurityV2 ---
func unpackMfaServersMfaVendorTypeDuoSecurityV2ListToSdk(ctx context.Context, list types.List) ([]identity_services.MfaServersMfaVendorTypeDuoSecurityV2, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.MfaServersMfaVendorTypeDuoSecurityV2")
	diags := diag.Diagnostics{}
	var data []models.MfaServersMfaVendorTypeDuoSecurityV2
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]identity_services.MfaServersMfaVendorTypeDuoSecurityV2, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.MfaServersMfaVendorTypeDuoSecurityV2{}.AttrTypes(), &item)
		unpacked, d := unpackMfaServersMfaVendorTypeDuoSecurityV2ToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.MfaServersMfaVendorTypeDuoSecurityV2", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for MfaServersMfaVendorTypeDuoSecurityV2 ---
func packMfaServersMfaVendorTypeDuoSecurityV2ListFromSdk(ctx context.Context, sdks []identity_services.MfaServersMfaVendorTypeDuoSecurityV2) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.MfaServersMfaVendorTypeDuoSecurityV2")
	diags := diag.Diagnostics{}
	var data []models.MfaServersMfaVendorTypeDuoSecurityV2

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.MfaServersMfaVendorTypeDuoSecurityV2
		obj, d := packMfaServersMfaVendorTypeDuoSecurityV2FromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.MfaServersMfaVendorTypeDuoSecurityV2{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.MfaServersMfaVendorTypeDuoSecurityV2", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.MfaServersMfaVendorTypeDuoSecurityV2{}.AttrType(), data)
}

// --- Unpacker for MfaServersMfaVendorTypeOktaAdaptiveV1 ---
func unpackMfaServersMfaVendorTypeOktaAdaptiveV1ToSdk(ctx context.Context, obj types.Object) (*identity_services.MfaServersMfaVendorTypeOktaAdaptiveV1, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.MfaServersMfaVendorTypeOktaAdaptiveV1", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.MfaServersMfaVendorTypeOktaAdaptiveV1
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk identity_services.MfaServersMfaVendorTypeOktaAdaptiveV1
	var d diag.Diagnostics
	// Handling Primitives
	if !model.OktaApiHost.IsNull() && !model.OktaApiHost.IsUnknown() {
		sdk.OktaApiHost = model.OktaApiHost.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "OktaApiHost", "value": sdk.OktaApiHost})
	}

	// Handling Primitives
	if !model.OktaBaseuri.IsNull() && !model.OktaBaseuri.IsUnknown() {
		sdk.OktaBaseuri = model.OktaBaseuri.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "OktaBaseuri", "value": sdk.OktaBaseuri})
	}

	// Handling Primitives
	if !model.OktaOrg.IsNull() && !model.OktaOrg.IsUnknown() {
		sdk.OktaOrg = model.OktaOrg.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "OktaOrg", "value": sdk.OktaOrg})
	}

	// Handling Primitives
	if !model.OktaTimeout.IsNull() && !model.OktaTimeout.IsUnknown() {
		sdk.OktaTimeout = int32(model.OktaTimeout.ValueInt64())
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "OktaTimeout", "value": sdk.OktaTimeout})
	}

	// Handling Primitives
	if !model.OktaToken.IsNull() && !model.OktaToken.IsUnknown() {
		sdk.OktaToken = model.OktaToken.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "OktaToken", "value": sdk.OktaToken})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.MfaServersMfaVendorTypeOktaAdaptiveV1", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for MfaServersMfaVendorTypeOktaAdaptiveV1 ---
func packMfaServersMfaVendorTypeOktaAdaptiveV1FromSdk(ctx context.Context, sdk identity_services.MfaServersMfaVendorTypeOktaAdaptiveV1) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.MfaServersMfaVendorTypeOktaAdaptiveV1", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.MfaServersMfaVendorTypeOktaAdaptiveV1
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.OktaApiHost = basetypes.NewStringValue(sdk.OktaApiHost)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "OktaApiHost", "value": sdk.OktaApiHost})
	// Handling Primitives
	// Standard primitive packing
	model.OktaBaseuri = basetypes.NewStringValue(sdk.OktaBaseuri)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "OktaBaseuri", "value": sdk.OktaBaseuri})
	// Handling Primitives
	// Standard primitive packing
	model.OktaOrg = basetypes.NewStringValue(sdk.OktaOrg)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "OktaOrg", "value": sdk.OktaOrg})
	// Handling Primitives
	// Standard primitive packing
	model.OktaTimeout = basetypes.NewInt64Value(int64(sdk.OktaTimeout))
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "OktaTimeout", "value": sdk.OktaTimeout})
	// Handling Primitives
	// Standard primitive packing
	model.OktaToken = basetypes.NewStringValue(sdk.OktaToken)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "OktaToken", "value": sdk.OktaToken})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.MfaServersMfaVendorTypeOktaAdaptiveV1{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.MfaServersMfaVendorTypeOktaAdaptiveV1", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for MfaServersMfaVendorTypeOktaAdaptiveV1 ---
func unpackMfaServersMfaVendorTypeOktaAdaptiveV1ListToSdk(ctx context.Context, list types.List) ([]identity_services.MfaServersMfaVendorTypeOktaAdaptiveV1, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.MfaServersMfaVendorTypeOktaAdaptiveV1")
	diags := diag.Diagnostics{}
	var data []models.MfaServersMfaVendorTypeOktaAdaptiveV1
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]identity_services.MfaServersMfaVendorTypeOktaAdaptiveV1, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.MfaServersMfaVendorTypeOktaAdaptiveV1{}.AttrTypes(), &item)
		unpacked, d := unpackMfaServersMfaVendorTypeOktaAdaptiveV1ToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.MfaServersMfaVendorTypeOktaAdaptiveV1", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for MfaServersMfaVendorTypeOktaAdaptiveV1 ---
func packMfaServersMfaVendorTypeOktaAdaptiveV1ListFromSdk(ctx context.Context, sdks []identity_services.MfaServersMfaVendorTypeOktaAdaptiveV1) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.MfaServersMfaVendorTypeOktaAdaptiveV1")
	diags := diag.Diagnostics{}
	var data []models.MfaServersMfaVendorTypeOktaAdaptiveV1

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.MfaServersMfaVendorTypeOktaAdaptiveV1
		obj, d := packMfaServersMfaVendorTypeOktaAdaptiveV1FromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.MfaServersMfaVendorTypeOktaAdaptiveV1{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.MfaServersMfaVendorTypeOktaAdaptiveV1", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.MfaServersMfaVendorTypeOktaAdaptiveV1{}.AttrType(), data)
}

// --- Unpacker for MfaServersMfaVendorTypePingIdentityV1 ---
func unpackMfaServersMfaVendorTypePingIdentityV1ToSdk(ctx context.Context, obj types.Object) (*identity_services.MfaServersMfaVendorTypePingIdentityV1, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.MfaServersMfaVendorTypePingIdentityV1", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.MfaServersMfaVendorTypePingIdentityV1
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk identity_services.MfaServersMfaVendorTypePingIdentityV1
	var d diag.Diagnostics
	// Handling Primitives
	if !model.PingApiHost.IsNull() && !model.PingApiHost.IsUnknown() {
		sdk.PingApiHost = model.PingApiHost.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "PingApiHost", "value": sdk.PingApiHost})
	}

	// Handling Primitives
	if !model.PingBaseuri.IsNull() && !model.PingBaseuri.IsUnknown() {
		sdk.PingBaseuri = model.PingBaseuri.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "PingBaseuri", "value": sdk.PingBaseuri})
	}

	// Handling Primitives
	if !model.PingOrgAlias.IsNull() && !model.PingOrgAlias.IsUnknown() {
		sdk.PingOrgAlias = model.PingOrgAlias.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "PingOrgAlias", "value": *sdk.PingOrgAlias})
	}

	// Handling Primitives
	if !model.PingTimeout.IsNull() && !model.PingTimeout.IsUnknown() {
		sdk.PingTimeout = int32(model.PingTimeout.ValueInt64())
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "PingTimeout", "value": sdk.PingTimeout})
	}

	// Handling Primitives
	if !model.PingToken.IsNull() && !model.PingToken.IsUnknown() {
		sdk.PingToken = model.PingToken.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "PingToken", "value": sdk.PingToken})
	}

	// Handling Primitives
	if !model.PingUseBase64Key.IsNull() && !model.PingUseBase64Key.IsUnknown() {
		sdk.PingUseBase64Key = model.PingUseBase64Key.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "PingUseBase64Key", "value": sdk.PingUseBase64Key})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.MfaServersMfaVendorTypePingIdentityV1", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for MfaServersMfaVendorTypePingIdentityV1 ---
func packMfaServersMfaVendorTypePingIdentityV1FromSdk(ctx context.Context, sdk identity_services.MfaServersMfaVendorTypePingIdentityV1) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.MfaServersMfaVendorTypePingIdentityV1", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.MfaServersMfaVendorTypePingIdentityV1
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.PingApiHost = basetypes.NewStringValue(sdk.PingApiHost)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "PingApiHost", "value": sdk.PingApiHost})
	// Handling Primitives
	// Standard primitive packing
	model.PingBaseuri = basetypes.NewStringValue(sdk.PingBaseuri)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "PingBaseuri", "value": sdk.PingBaseuri})
	// Handling Primitives
	// Standard primitive packing
	if sdk.PingOrgAlias != nil {
		model.PingOrgAlias = basetypes.NewStringValue(*sdk.PingOrgAlias)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "PingOrgAlias", "value": *sdk.PingOrgAlias})
	} else {
		model.PingOrgAlias = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	model.PingTimeout = basetypes.NewInt64Value(int64(sdk.PingTimeout))
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "PingTimeout", "value": sdk.PingTimeout})
	// Handling Primitives
	// Standard primitive packing
	model.PingToken = basetypes.NewStringValue(sdk.PingToken)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "PingToken", "value": sdk.PingToken})
	// Handling Primitives
	// Standard primitive packing
	model.PingUseBase64Key = basetypes.NewStringValue(sdk.PingUseBase64Key)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "PingUseBase64Key", "value": sdk.PingUseBase64Key})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.MfaServersMfaVendorTypePingIdentityV1{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.MfaServersMfaVendorTypePingIdentityV1", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for MfaServersMfaVendorTypePingIdentityV1 ---
func unpackMfaServersMfaVendorTypePingIdentityV1ListToSdk(ctx context.Context, list types.List) ([]identity_services.MfaServersMfaVendorTypePingIdentityV1, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.MfaServersMfaVendorTypePingIdentityV1")
	diags := diag.Diagnostics{}
	var data []models.MfaServersMfaVendorTypePingIdentityV1
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]identity_services.MfaServersMfaVendorTypePingIdentityV1, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.MfaServersMfaVendorTypePingIdentityV1{}.AttrTypes(), &item)
		unpacked, d := unpackMfaServersMfaVendorTypePingIdentityV1ToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.MfaServersMfaVendorTypePingIdentityV1", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for MfaServersMfaVendorTypePingIdentityV1 ---
func packMfaServersMfaVendorTypePingIdentityV1ListFromSdk(ctx context.Context, sdks []identity_services.MfaServersMfaVendorTypePingIdentityV1) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.MfaServersMfaVendorTypePingIdentityV1")
	diags := diag.Diagnostics{}
	var data []models.MfaServersMfaVendorTypePingIdentityV1

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.MfaServersMfaVendorTypePingIdentityV1
		obj, d := packMfaServersMfaVendorTypePingIdentityV1FromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.MfaServersMfaVendorTypePingIdentityV1{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.MfaServersMfaVendorTypePingIdentityV1", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.MfaServersMfaVendorTypePingIdentityV1{}.AttrType(), data)
}

// --- Unpacker for MfaServersMfaVendorTypeRsaSecuridAccessV1 ---
func unpackMfaServersMfaVendorTypeRsaSecuridAccessV1ToSdk(ctx context.Context, obj types.Object) (*identity_services.MfaServersMfaVendorTypeRsaSecuridAccessV1, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.MfaServersMfaVendorTypeRsaSecuridAccessV1", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.MfaServersMfaVendorTypeRsaSecuridAccessV1
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk identity_services.MfaServersMfaVendorTypeRsaSecuridAccessV1
	var d diag.Diagnostics
	// Handling Primitives
	if !model.RsaAccessid.IsNull() && !model.RsaAccessid.IsUnknown() {
		sdk.RsaAccessid = model.RsaAccessid.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "RsaAccessid", "value": *sdk.RsaAccessid})
	}

	// Handling Primitives
	if !model.RsaAccesskey.IsNull() && !model.RsaAccesskey.IsUnknown() {
		sdk.RsaAccesskey = model.RsaAccesskey.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "RsaAccesskey", "value": *sdk.RsaAccesskey})
	}

	// Handling Primitives
	if !model.RsaApiHost.IsNull() && !model.RsaApiHost.IsUnknown() {
		sdk.RsaApiHost = model.RsaApiHost.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "RsaApiHost", "value": *sdk.RsaApiHost})
	}

	// Handling Primitives
	if !model.RsaAssurancepolicyid.IsNull() && !model.RsaAssurancepolicyid.IsUnknown() {
		sdk.RsaAssurancepolicyid = model.RsaAssurancepolicyid.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "RsaAssurancepolicyid", "value": *sdk.RsaAssurancepolicyid})
	}

	// Handling Primitives
	if !model.RsaBaseuri.IsNull() && !model.RsaBaseuri.IsUnknown() {
		sdk.RsaBaseuri = model.RsaBaseuri.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "RsaBaseuri", "value": *sdk.RsaBaseuri})
	}

	// Handling Primitives
	if !model.RsaTimeout.IsNull() && !model.RsaTimeout.IsUnknown() {
		val := int32(model.RsaTimeout.ValueInt64())
		sdk.RsaTimeout = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "RsaTimeout", "value": *sdk.RsaTimeout})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.MfaServersMfaVendorTypeRsaSecuridAccessV1", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for MfaServersMfaVendorTypeRsaSecuridAccessV1 ---
func packMfaServersMfaVendorTypeRsaSecuridAccessV1FromSdk(ctx context.Context, sdk identity_services.MfaServersMfaVendorTypeRsaSecuridAccessV1) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.MfaServersMfaVendorTypeRsaSecuridAccessV1", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.MfaServersMfaVendorTypeRsaSecuridAccessV1
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.RsaAccessid != nil {
		model.RsaAccessid = basetypes.NewStringValue(*sdk.RsaAccessid)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "RsaAccessid", "value": *sdk.RsaAccessid})
	} else {
		model.RsaAccessid = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.RsaAccesskey != nil {
		model.RsaAccesskey = basetypes.NewStringValue(*sdk.RsaAccesskey)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "RsaAccesskey", "value": *sdk.RsaAccesskey})
	} else {
		model.RsaAccesskey = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.RsaApiHost != nil {
		model.RsaApiHost = basetypes.NewStringValue(*sdk.RsaApiHost)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "RsaApiHost", "value": *sdk.RsaApiHost})
	} else {
		model.RsaApiHost = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.RsaAssurancepolicyid != nil {
		model.RsaAssurancepolicyid = basetypes.NewStringValue(*sdk.RsaAssurancepolicyid)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "RsaAssurancepolicyid", "value": *sdk.RsaAssurancepolicyid})
	} else {
		model.RsaAssurancepolicyid = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.RsaBaseuri != nil {
		model.RsaBaseuri = basetypes.NewStringValue(*sdk.RsaBaseuri)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "RsaBaseuri", "value": *sdk.RsaBaseuri})
	} else {
		model.RsaBaseuri = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.RsaTimeout != nil {
		model.RsaTimeout = basetypes.NewInt64Value(int64(*sdk.RsaTimeout))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "RsaTimeout", "value": *sdk.RsaTimeout})
	} else {
		model.RsaTimeout = basetypes.NewInt64Null()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.MfaServersMfaVendorTypeRsaSecuridAccessV1{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.MfaServersMfaVendorTypeRsaSecuridAccessV1", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for MfaServersMfaVendorTypeRsaSecuridAccessV1 ---
func unpackMfaServersMfaVendorTypeRsaSecuridAccessV1ListToSdk(ctx context.Context, list types.List) ([]identity_services.MfaServersMfaVendorTypeRsaSecuridAccessV1, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.MfaServersMfaVendorTypeRsaSecuridAccessV1")
	diags := diag.Diagnostics{}
	var data []models.MfaServersMfaVendorTypeRsaSecuridAccessV1
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]identity_services.MfaServersMfaVendorTypeRsaSecuridAccessV1, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.MfaServersMfaVendorTypeRsaSecuridAccessV1{}.AttrTypes(), &item)
		unpacked, d := unpackMfaServersMfaVendorTypeRsaSecuridAccessV1ToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.MfaServersMfaVendorTypeRsaSecuridAccessV1", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for MfaServersMfaVendorTypeRsaSecuridAccessV1 ---
func packMfaServersMfaVendorTypeRsaSecuridAccessV1ListFromSdk(ctx context.Context, sdks []identity_services.MfaServersMfaVendorTypeRsaSecuridAccessV1) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.MfaServersMfaVendorTypeRsaSecuridAccessV1")
	diags := diag.Diagnostics{}
	var data []models.MfaServersMfaVendorTypeRsaSecuridAccessV1

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.MfaServersMfaVendorTypeRsaSecuridAccessV1
		obj, d := packMfaServersMfaVendorTypeRsaSecuridAccessV1FromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.MfaServersMfaVendorTypeRsaSecuridAccessV1{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.MfaServersMfaVendorTypeRsaSecuridAccessV1", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.MfaServersMfaVendorTypeRsaSecuridAccessV1{}.AttrType(), data)
}
