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

// scepProfilesSensitiveValuePatcher is an in-memory struct to temporarily store plaintext
// and encrypted values for sensitive fields during the Create/Update/Read workflows.
type scepProfilesSensitiveValuePatcher struct {
	scep_challenge_dynamic_password_plaintext basetypes.StringValue
	scep_challenge_dynamic_password_encrypted basetypes.StringValue
}

// populatePatcherFromState populates the patcher struct from the resource's state.
func (p *scepProfilesSensitiveValuePatcher) populatePatcherFromState(ctx context.Context, state models.ScepProfiles) diag.Diagnostics {
	var diags diag.Diagnostics
	if state.EncryptedValues.IsNull() || state.EncryptedValues.IsUnknown() {
		return diags
	}

	var ev map[string]string
	diags.Append(state.EncryptedValues.ElementsAs(ctx, &ev, false)...)
	if diags.HasError() {
		return diags
	}
	if val, ok := ev["scep_challenge_dynamic_password_plaintext"]; ok {
		p.scep_challenge_dynamic_password_plaintext = basetypes.NewStringValue(val)
	}
	if val, ok := ev["scep_challenge_dynamic_password_encrypted"]; ok {
		p.scep_challenge_dynamic_password_encrypted = basetypes.NewStringValue(val)
	}

	return diags
}

// populateEncryptedValuesMap returns a map of the patcher's values for saving to state.
func (p *scepProfilesSensitiveValuePatcher) populateEncryptedValuesMap() map[string]string {
	ev := make(map[string]string)
	if !p.scep_challenge_dynamic_password_plaintext.IsNull() {
		ev["scep_challenge_dynamic_password_plaintext"] = p.scep_challenge_dynamic_password_plaintext.ValueString()
	}
	if !p.scep_challenge_dynamic_password_encrypted.IsNull() {
		ev["scep_challenge_dynamic_password_encrypted"] = p.scep_challenge_dynamic_password_encrypted.ValueString()
	}
	return ev
}

// --- Unpacker for ScepProfiles ---
func unpackScepProfilesToSdk(ctx context.Context, obj types.Object) (*identity_services.ScepProfiles, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ScepProfiles", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ScepProfiles
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk identity_services.ScepProfiles
	var d diag.Diagnostics
	// Handling Objects
	if !model.Algorithm.IsNull() && !model.Algorithm.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Algorithm")
		unpacked, d := unpackScepProfilesAlgorithmToSdk(ctx, model.Algorithm)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Algorithm"})
		}
		if unpacked != nil {
			sdk.Algorithm = *unpacked
		}
	}

	// Handling Primitives
	if !model.CaIdentityName.IsNull() && !model.CaIdentityName.IsUnknown() {
		sdk.CaIdentityName = model.CaIdentityName.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "CaIdentityName", "value": sdk.CaIdentityName})
	}

	// Handling Objects
	if !model.CertificateAttributes.IsNull() && !model.CertificateAttributes.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field CertificateAttributes")
		unpacked, d := unpackScepProfilesCertificateAttributesToSdk(ctx, model.CertificateAttributes)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "CertificateAttributes"})
		}
		if unpacked != nil {
			sdk.CertificateAttributes = unpacked
		}
	}

	// Handling Primitives
	if !model.Device.IsNull() && !model.Device.IsUnknown() {
		sdk.Device = model.Device.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Device", "value": *sdk.Device})
	}

	// Handling Primitives
	if !model.Digest.IsNull() && !model.Digest.IsUnknown() {
		sdk.Digest = model.Digest.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Digest", "value": sdk.Digest})
	}

	// Handling Primitives
	if !model.Fingerprint.IsNull() && !model.Fingerprint.IsUnknown() {
		sdk.Fingerprint = model.Fingerprint.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Fingerprint", "value": *sdk.Fingerprint})
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
	if !model.ScepCaCert.IsNull() && !model.ScepCaCert.IsUnknown() {
		sdk.ScepCaCert = model.ScepCaCert.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "ScepCaCert", "value": *sdk.ScepCaCert})
	}

	// Handling Objects
	if !model.ScepChallenge.IsNull() && !model.ScepChallenge.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field ScepChallenge")
		unpacked, d := unpackScepProfilesScepChallengeToSdk(ctx, model.ScepChallenge)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "ScepChallenge"})
		}
		if unpacked != nil {
			sdk.ScepChallenge = *unpacked
		}
	}

	// Handling Primitives
	if !model.ScepClientCert.IsNull() && !model.ScepClientCert.IsUnknown() {
		sdk.ScepClientCert = model.ScepClientCert.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "ScepClientCert", "value": *sdk.ScepClientCert})
	}

	// Handling Primitives
	if !model.ScepUrl.IsNull() && !model.ScepUrl.IsUnknown() {
		sdk.ScepUrl = model.ScepUrl.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "ScepUrl", "value": sdk.ScepUrl})
	}

	// Handling Primitives
	if !model.Snippet.IsNull() && !model.Snippet.IsUnknown() {
		sdk.Snippet = model.Snippet.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	}

	// Handling Primitives
	if !model.Subject.IsNull() && !model.Subject.IsUnknown() {
		sdk.Subject = model.Subject.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Subject", "value": sdk.Subject})
	}

	// Handling Primitives
	if !model.UseAsDigitalSignature.IsNull() && !model.UseAsDigitalSignature.IsUnknown() {
		sdk.UseAsDigitalSignature = model.UseAsDigitalSignature.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "UseAsDigitalSignature", "value": *sdk.UseAsDigitalSignature})
	}

	// Handling Primitives
	if !model.UseForKeyEncipherment.IsNull() && !model.UseForKeyEncipherment.IsUnknown() {
		sdk.UseForKeyEncipherment = model.UseForKeyEncipherment.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "UseForKeyEncipherment", "value": *sdk.UseForKeyEncipherment})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ScepProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ScepProfiles ---
func packScepProfilesFromSdk(ctx context.Context, sdk identity_services.ScepProfiles) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ScepProfiles", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ScepProfiles
	var d diag.Diagnostics
	// MEGA FIX FOR MAP TYPE MISMATCH (NOT ALL MODELS MAY HAVE EncryptedValues)
	model.EncryptedValues = basetypes.NewMapNull(basetypes.StringType{})
	// Handling Objects
	// This is a regular nested object that has its own packer.
	// Logic for non-pointer / value-type nested objects
	if !reflect.ValueOf(sdk.Algorithm).IsZero() {
		tflog.Debug(ctx, "Packing nested object for field Algorithm")
		packed, d := packScepProfilesAlgorithmFromSdk(ctx, sdk.Algorithm)
		diags.Append(d...)
		model.Algorithm = packed
	} else {
		model.Algorithm = basetypes.NewObjectNull(models.ScepProfilesAlgorithm{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	model.CaIdentityName = basetypes.NewStringValue(sdk.CaIdentityName)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "CaIdentityName", "value": sdk.CaIdentityName})
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.CertificateAttributes != nil {
		tflog.Debug(ctx, "Packing nested object for field CertificateAttributes")
		packed, d := packScepProfilesCertificateAttributesFromSdk(ctx, *sdk.CertificateAttributes)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "CertificateAttributes"})
		}
		model.CertificateAttributes = packed
	} else {
		model.CertificateAttributes = basetypes.NewObjectNull(models.ScepProfilesCertificateAttributes{}.AttrTypes())
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
	model.Digest = basetypes.NewStringValue(sdk.Digest)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Digest", "value": sdk.Digest})
	// Handling Primitives
	// Standard primitive packing
	if sdk.Fingerprint != nil {
		model.Fingerprint = basetypes.NewStringValue(*sdk.Fingerprint)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Fingerprint", "value": *sdk.Fingerprint})
	} else {
		model.Fingerprint = basetypes.NewStringNull()
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
	if sdk.ScepCaCert != nil {
		model.ScepCaCert = basetypes.NewStringValue(*sdk.ScepCaCert)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "ScepCaCert", "value": *sdk.ScepCaCert})
	} else {
		model.ScepCaCert = basetypes.NewStringNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	// Logic for non-pointer / value-type nested objects
	if !reflect.ValueOf(sdk.ScepChallenge).IsZero() {
		tflog.Debug(ctx, "Packing nested object for field ScepChallenge")
		packed, d := packScepProfilesScepChallengeFromSdk(ctx, sdk.ScepChallenge)
		diags.Append(d...)
		model.ScepChallenge = packed
	} else {
		model.ScepChallenge = basetypes.NewObjectNull(models.ScepProfilesScepChallenge{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.ScepClientCert != nil {
		model.ScepClientCert = basetypes.NewStringValue(*sdk.ScepClientCert)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "ScepClientCert", "value": *sdk.ScepClientCert})
	} else {
		model.ScepClientCert = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	model.ScepUrl = basetypes.NewStringValue(sdk.ScepUrl)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "ScepUrl", "value": sdk.ScepUrl})
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
	model.Subject = basetypes.NewStringValue(sdk.Subject)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Subject", "value": sdk.Subject})
	// Handling Primitives
	// Standard primitive packing
	if sdk.UseAsDigitalSignature != nil {
		model.UseAsDigitalSignature = basetypes.NewBoolValue(*sdk.UseAsDigitalSignature)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "UseAsDigitalSignature", "value": *sdk.UseAsDigitalSignature})
	} else {
		model.UseAsDigitalSignature = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.UseForKeyEncipherment != nil {
		model.UseForKeyEncipherment = basetypes.NewBoolValue(*sdk.UseForKeyEncipherment)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "UseForKeyEncipherment", "value": *sdk.UseForKeyEncipherment})
	} else {
		model.UseForKeyEncipherment = basetypes.NewBoolNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ScepProfiles{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ScepProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ScepProfiles ---
func unpackScepProfilesListToSdk(ctx context.Context, list types.List) ([]identity_services.ScepProfiles, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ScepProfiles")
	diags := diag.Diagnostics{}
	var data []models.ScepProfiles
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]identity_services.ScepProfiles, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ScepProfiles{}.AttrTypes(), &item)
		unpacked, d := unpackScepProfilesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ScepProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ScepProfiles ---
func packScepProfilesListFromSdk(ctx context.Context, sdks []identity_services.ScepProfiles) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ScepProfiles")
	diags := diag.Diagnostics{}
	var data []models.ScepProfiles

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ScepProfiles
		obj, d := packScepProfilesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ScepProfiles{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ScepProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ScepProfiles{}.AttrType(), data)
}

// --- Unpacker for ScepProfilesAlgorithm ---
func unpackScepProfilesAlgorithmToSdk(ctx context.Context, obj types.Object) (*identity_services.ScepProfilesAlgorithm, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ScepProfilesAlgorithm", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ScepProfilesAlgorithm
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk identity_services.ScepProfilesAlgorithm
	var d diag.Diagnostics
	// Handling Objects
	if !model.Rsa.IsNull() && !model.Rsa.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Rsa")
		unpacked, d := unpackScepProfilesAlgorithmRsaToSdk(ctx, model.Rsa)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Rsa"})
		}
		if unpacked != nil {
			sdk.Rsa = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ScepProfilesAlgorithm", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ScepProfilesAlgorithm ---
func packScepProfilesAlgorithmFromSdk(ctx context.Context, sdk identity_services.ScepProfilesAlgorithm) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ScepProfilesAlgorithm", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ScepProfilesAlgorithm
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Rsa != nil {
		tflog.Debug(ctx, "Packing nested object for field Rsa")
		packed, d := packScepProfilesAlgorithmRsaFromSdk(ctx, *sdk.Rsa)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Rsa"})
		}
		model.Rsa = packed
	} else {
		model.Rsa = basetypes.NewObjectNull(models.ScepProfilesAlgorithmRsa{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ScepProfilesAlgorithm{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ScepProfilesAlgorithm", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ScepProfilesAlgorithm ---
func unpackScepProfilesAlgorithmListToSdk(ctx context.Context, list types.List) ([]identity_services.ScepProfilesAlgorithm, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ScepProfilesAlgorithm")
	diags := diag.Diagnostics{}
	var data []models.ScepProfilesAlgorithm
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]identity_services.ScepProfilesAlgorithm, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ScepProfilesAlgorithm{}.AttrTypes(), &item)
		unpacked, d := unpackScepProfilesAlgorithmToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ScepProfilesAlgorithm", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ScepProfilesAlgorithm ---
func packScepProfilesAlgorithmListFromSdk(ctx context.Context, sdks []identity_services.ScepProfilesAlgorithm) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ScepProfilesAlgorithm")
	diags := diag.Diagnostics{}
	var data []models.ScepProfilesAlgorithm

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ScepProfilesAlgorithm
		obj, d := packScepProfilesAlgorithmFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ScepProfilesAlgorithm{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ScepProfilesAlgorithm", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ScepProfilesAlgorithm{}.AttrType(), data)
}

// --- Unpacker for ScepProfilesAlgorithmRsa ---
func unpackScepProfilesAlgorithmRsaToSdk(ctx context.Context, obj types.Object) (*identity_services.ScepProfilesAlgorithmRsa, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ScepProfilesAlgorithmRsa", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ScepProfilesAlgorithmRsa
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk identity_services.ScepProfilesAlgorithmRsa
	var d diag.Diagnostics
	// Handling Primitives
	if !model.RsaNbits.IsNull() && !model.RsaNbits.IsUnknown() {
		val := int32(model.RsaNbits.ValueInt64())
		sdk.RsaNbits = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "RsaNbits", "value": *sdk.RsaNbits})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ScepProfilesAlgorithmRsa", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ScepProfilesAlgorithmRsa ---
func packScepProfilesAlgorithmRsaFromSdk(ctx context.Context, sdk identity_services.ScepProfilesAlgorithmRsa) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ScepProfilesAlgorithmRsa", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ScepProfilesAlgorithmRsa
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.RsaNbits != nil {
		model.RsaNbits = basetypes.NewInt64Value(int64(*sdk.RsaNbits))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "RsaNbits", "value": *sdk.RsaNbits})
	} else {
		model.RsaNbits = basetypes.NewInt64Null()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ScepProfilesAlgorithmRsa{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ScepProfilesAlgorithmRsa", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ScepProfilesAlgorithmRsa ---
func unpackScepProfilesAlgorithmRsaListToSdk(ctx context.Context, list types.List) ([]identity_services.ScepProfilesAlgorithmRsa, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ScepProfilesAlgorithmRsa")
	diags := diag.Diagnostics{}
	var data []models.ScepProfilesAlgorithmRsa
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]identity_services.ScepProfilesAlgorithmRsa, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ScepProfilesAlgorithmRsa{}.AttrTypes(), &item)
		unpacked, d := unpackScepProfilesAlgorithmRsaToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ScepProfilesAlgorithmRsa", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ScepProfilesAlgorithmRsa ---
func packScepProfilesAlgorithmRsaListFromSdk(ctx context.Context, sdks []identity_services.ScepProfilesAlgorithmRsa) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ScepProfilesAlgorithmRsa")
	diags := diag.Diagnostics{}
	var data []models.ScepProfilesAlgorithmRsa

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ScepProfilesAlgorithmRsa
		obj, d := packScepProfilesAlgorithmRsaFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ScepProfilesAlgorithmRsa{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ScepProfilesAlgorithmRsa", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ScepProfilesAlgorithmRsa{}.AttrType(), data)
}

// --- Unpacker for ScepProfilesCertificateAttributes ---
func unpackScepProfilesCertificateAttributesToSdk(ctx context.Context, obj types.Object) (*identity_services.ScepProfilesCertificateAttributes, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ScepProfilesCertificateAttributes", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ScepProfilesCertificateAttributes
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk identity_services.ScepProfilesCertificateAttributes
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Dnsname.IsNull() && !model.Dnsname.IsUnknown() {
		sdk.Dnsname = model.Dnsname.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Dnsname", "value": *sdk.Dnsname})
	}

	// Handling Primitives
	if !model.Rfc822name.IsNull() && !model.Rfc822name.IsUnknown() {
		sdk.Rfc822name = model.Rfc822name.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Rfc822name", "value": *sdk.Rfc822name})
	}

	// Handling Primitives
	if !model.UniformResourceIdentifier.IsNull() && !model.UniformResourceIdentifier.IsUnknown() {
		sdk.UniformResourceIdentifier = model.UniformResourceIdentifier.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "UniformResourceIdentifier", "value": *sdk.UniformResourceIdentifier})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ScepProfilesCertificateAttributes", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ScepProfilesCertificateAttributes ---
func packScepProfilesCertificateAttributesFromSdk(ctx context.Context, sdk identity_services.ScepProfilesCertificateAttributes) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ScepProfilesCertificateAttributes", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ScepProfilesCertificateAttributes
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Dnsname != nil {
		model.Dnsname = basetypes.NewStringValue(*sdk.Dnsname)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Dnsname", "value": *sdk.Dnsname})
	} else {
		model.Dnsname = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Rfc822name != nil {
		model.Rfc822name = basetypes.NewStringValue(*sdk.Rfc822name)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Rfc822name", "value": *sdk.Rfc822name})
	} else {
		model.Rfc822name = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.UniformResourceIdentifier != nil {
		model.UniformResourceIdentifier = basetypes.NewStringValue(*sdk.UniformResourceIdentifier)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "UniformResourceIdentifier", "value": *sdk.UniformResourceIdentifier})
	} else {
		model.UniformResourceIdentifier = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ScepProfilesCertificateAttributes{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ScepProfilesCertificateAttributes", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ScepProfilesCertificateAttributes ---
func unpackScepProfilesCertificateAttributesListToSdk(ctx context.Context, list types.List) ([]identity_services.ScepProfilesCertificateAttributes, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ScepProfilesCertificateAttributes")
	diags := diag.Diagnostics{}
	var data []models.ScepProfilesCertificateAttributes
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]identity_services.ScepProfilesCertificateAttributes, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ScepProfilesCertificateAttributes{}.AttrTypes(), &item)
		unpacked, d := unpackScepProfilesCertificateAttributesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ScepProfilesCertificateAttributes", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ScepProfilesCertificateAttributes ---
func packScepProfilesCertificateAttributesListFromSdk(ctx context.Context, sdks []identity_services.ScepProfilesCertificateAttributes) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ScepProfilesCertificateAttributes")
	diags := diag.Diagnostics{}
	var data []models.ScepProfilesCertificateAttributes

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ScepProfilesCertificateAttributes
		obj, d := packScepProfilesCertificateAttributesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ScepProfilesCertificateAttributes{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ScepProfilesCertificateAttributes", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ScepProfilesCertificateAttributes{}.AttrType(), data)
}

// --- Unpacker for ScepProfilesScepChallenge ---
func unpackScepProfilesScepChallengeToSdk(ctx context.Context, obj types.Object) (*identity_services.ScepProfilesScepChallenge, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ScepProfilesScepChallenge", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ScepProfilesScepChallenge
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk identity_services.ScepProfilesScepChallenge
	var d diag.Diagnostics
	// Handling Objects
	if !model.Dynamic.IsNull() && !model.Dynamic.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Dynamic")
		unpacked, d := unpackScepProfilesScepChallengeDynamicToSdk(ctx, model.Dynamic)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Dynamic"})
		}
		if unpacked != nil {
			sdk.Dynamic = unpacked
		}
	}

	// Handling Primitives
	if !model.Fixed.IsNull() && !model.Fixed.IsUnknown() {
		sdk.Fixed = model.Fixed.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Fixed", "value": *sdk.Fixed})
	}

	// Handling Primitives
	if !model.None.IsNull() && !model.None.IsUnknown() {
		sdk.None = model.None.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "None", "value": *sdk.None})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ScepProfilesScepChallenge", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ScepProfilesScepChallenge ---
func packScepProfilesScepChallengeFromSdk(ctx context.Context, sdk identity_services.ScepProfilesScepChallenge) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ScepProfilesScepChallenge", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ScepProfilesScepChallenge
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Dynamic != nil {
		tflog.Debug(ctx, "Packing nested object for field Dynamic")
		packed, d := packScepProfilesScepChallengeDynamicFromSdk(ctx, *sdk.Dynamic)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Dynamic"})
		}
		model.Dynamic = packed
	} else {
		model.Dynamic = basetypes.NewObjectNull(models.ScepProfilesScepChallengeDynamic{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Fixed != nil {
		model.Fixed = basetypes.NewStringValue(*sdk.Fixed)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Fixed", "value": *sdk.Fixed})
	} else {
		model.Fixed = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.None != nil {
		model.None = basetypes.NewStringValue(*sdk.None)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "None", "value": *sdk.None})
	} else {
		model.None = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ScepProfilesScepChallenge{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ScepProfilesScepChallenge", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ScepProfilesScepChallenge ---
func unpackScepProfilesScepChallengeListToSdk(ctx context.Context, list types.List) ([]identity_services.ScepProfilesScepChallenge, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ScepProfilesScepChallenge")
	diags := diag.Diagnostics{}
	var data []models.ScepProfilesScepChallenge
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]identity_services.ScepProfilesScepChallenge, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ScepProfilesScepChallenge{}.AttrTypes(), &item)
		unpacked, d := unpackScepProfilesScepChallengeToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ScepProfilesScepChallenge", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ScepProfilesScepChallenge ---
func packScepProfilesScepChallengeListFromSdk(ctx context.Context, sdks []identity_services.ScepProfilesScepChallenge) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ScepProfilesScepChallenge")
	diags := diag.Diagnostics{}
	var data []models.ScepProfilesScepChallenge

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ScepProfilesScepChallenge
		obj, d := packScepProfilesScepChallengeFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ScepProfilesScepChallenge{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ScepProfilesScepChallenge", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ScepProfilesScepChallenge{}.AttrType(), data)
}

// --- Unpacker for ScepProfilesScepChallengeDynamic ---
func unpackScepProfilesScepChallengeDynamicToSdk(ctx context.Context, obj types.Object) (*identity_services.ScepProfilesScepChallengeDynamic, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ScepProfilesScepChallengeDynamic", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ScepProfilesScepChallengeDynamic
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk identity_services.ScepProfilesScepChallengeDynamic
	var d diag.Diagnostics
	// Handling Primitives
	if !model.OtpServerUrl.IsNull() && !model.OtpServerUrl.IsUnknown() {
		sdk.OtpServerUrl = model.OtpServerUrl.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "OtpServerUrl", "value": *sdk.OtpServerUrl})
	}

	// Handling Primitives
	if !model.Password.IsNull() && !model.Password.IsUnknown() {
		sdk.Password = model.Password.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Password", "value": *sdk.Password})
	}

	// Handling Primitives
	if !model.Username.IsNull() && !model.Username.IsUnknown() {
		sdk.Username = model.Username.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Username", "value": *sdk.Username})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ScepProfilesScepChallengeDynamic", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ScepProfilesScepChallengeDynamic ---
func packScepProfilesScepChallengeDynamicFromSdk(ctx context.Context, sdk identity_services.ScepProfilesScepChallengeDynamic) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ScepProfilesScepChallengeDynamic", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ScepProfilesScepChallengeDynamic
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.OtpServerUrl != nil {
		model.OtpServerUrl = basetypes.NewStringValue(*sdk.OtpServerUrl)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "OtpServerUrl", "value": *sdk.OtpServerUrl})
	} else {
		model.OtpServerUrl = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Password != nil {
		model.Password = basetypes.NewStringValue(*sdk.Password)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Password", "value": *sdk.Password})
	} else {
		model.Password = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Username != nil {
		model.Username = basetypes.NewStringValue(*sdk.Username)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Username", "value": *sdk.Username})
	} else {
		model.Username = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ScepProfilesScepChallengeDynamic{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ScepProfilesScepChallengeDynamic", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ScepProfilesScepChallengeDynamic ---
func unpackScepProfilesScepChallengeDynamicListToSdk(ctx context.Context, list types.List) ([]identity_services.ScepProfilesScepChallengeDynamic, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ScepProfilesScepChallengeDynamic")
	diags := diag.Diagnostics{}
	var data []models.ScepProfilesScepChallengeDynamic
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]identity_services.ScepProfilesScepChallengeDynamic, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ScepProfilesScepChallengeDynamic{}.AttrTypes(), &item)
		unpacked, d := unpackScepProfilesScepChallengeDynamicToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ScepProfilesScepChallengeDynamic", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ScepProfilesScepChallengeDynamic ---
func packScepProfilesScepChallengeDynamicListFromSdk(ctx context.Context, sdks []identity_services.ScepProfilesScepChallengeDynamic) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ScepProfilesScepChallengeDynamic")
	diags := diag.Diagnostics{}
	var data []models.ScepProfilesScepChallengeDynamic

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ScepProfilesScepChallengeDynamic
		obj, d := packScepProfilesScepChallengeDynamicFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ScepProfilesScepChallengeDynamic{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ScepProfilesScepChallengeDynamic", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ScepProfilesScepChallengeDynamic{}.AttrType(), data)
}
