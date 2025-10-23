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

// ikeGatewaysSensitiveValuePatcher is an in-memory struct to temporarily store plaintext
// and encrypted values for sensitive fields during the Create/Update/Read workflows.
type ikeGatewaysSensitiveValuePatcher struct {
	authentication_pre_shared_key_key_plaintext basetypes.StringValue
	authentication_pre_shared_key_key_encrypted basetypes.StringValue
}

// populatePatcherFromState populates the patcher struct from the resource's state.
func (p *ikeGatewaysSensitiveValuePatcher) populatePatcherFromState(ctx context.Context, state models.IkeGateways) diag.Diagnostics {
	var diags diag.Diagnostics
	if state.EncryptedValues.IsNull() || state.EncryptedValues.IsUnknown() {
		return diags
	}

	var ev map[string]string
	diags.Append(state.EncryptedValues.ElementsAs(ctx, &ev, false)...)
	if diags.HasError() {
		return diags
	}
	if val, ok := ev["authentication_pre_shared_key_key_plaintext"]; ok {
		p.authentication_pre_shared_key_key_plaintext = basetypes.NewStringValue(val)
	}
	if val, ok := ev["authentication_pre_shared_key_key_encrypted"]; ok {
		p.authentication_pre_shared_key_key_encrypted = basetypes.NewStringValue(val)
	}

	return diags
}

// populateEncryptedValuesMap returns a map of the patcher's values for saving to state.
func (p *ikeGatewaysSensitiveValuePatcher) populateEncryptedValuesMap() map[string]string {
	ev := make(map[string]string)
	if !p.authentication_pre_shared_key_key_plaintext.IsNull() {
		ev["authentication_pre_shared_key_key_plaintext"] = p.authentication_pre_shared_key_key_plaintext.ValueString()
	}
	if !p.authentication_pre_shared_key_key_encrypted.IsNull() {
		ev["authentication_pre_shared_key_key_encrypted"] = p.authentication_pre_shared_key_key_encrypted.ValueString()
	}
	return ev
}

// --- Unpacker for IkeGateways ---
func unpackIkeGatewaysToSdk(ctx context.Context, obj types.Object) (*network_services.IkeGateways, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.IkeGateways", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.IkeGateways
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.IkeGateways
	var d diag.Diagnostics

	// Handling Objects
	if !model.Authentication.IsNull() && !model.Authentication.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Authentication")
		unpacked, d := unpackIkeGatewaysAuthenticationToSdk(ctx, model.Authentication)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Authentication"})
		}
		if unpacked != nil {
			sdk.Authentication = *unpacked
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
		sdk.Id = model.Id.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Id", "value": *sdk.Id})
	}

	// Handling Objects
	if !model.LocalAddress.IsNull() && !model.LocalAddress.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field LocalAddress")
		unpacked, d := unpackIkeGatewaysLocalAddressToSdk(ctx, model.LocalAddress)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "LocalAddress"})
		}
		if unpacked != nil {
			sdk.LocalAddress = unpacked
		}
	}

	// Handling Objects
	if !model.LocalId.IsNull() && !model.LocalId.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field LocalId")
		unpacked, d := unpackIkeGatewaysLocalIdToSdk(ctx, model.LocalId)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "LocalId"})
		}
		if unpacked != nil {
			sdk.LocalId = unpacked
		}
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	// Handling Objects
	if !model.PeerAddress.IsNull() && !model.PeerAddress.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field PeerAddress")
		unpacked, d := unpackIkeGatewaysPeerAddressToSdk(ctx, model.PeerAddress)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "PeerAddress"})
		}
		if unpacked != nil {
			sdk.PeerAddress = *unpacked
		}
	}

	// Handling Objects
	if !model.PeerId.IsNull() && !model.PeerId.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field PeerId")
		unpacked, d := unpackIkeGatewaysPeerIdToSdk(ctx, model.PeerId)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "PeerId"})
		}
		if unpacked != nil {
			sdk.PeerId = unpacked
		}
	}

	// Handling Objects
	if !model.Protocol.IsNull() && !model.Protocol.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Protocol")
		unpacked, d := unpackIkeGatewaysProtocolToSdk(ctx, model.Protocol)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Protocol"})
		}
		if unpacked != nil {
			sdk.Protocol = *unpacked
		}
	}

	// Handling Objects
	if !model.ProtocolCommon.IsNull() && !model.ProtocolCommon.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field ProtocolCommon")
		unpacked, d := unpackIkeGatewaysProtocolCommonToSdk(ctx, model.ProtocolCommon)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "ProtocolCommon"})
		}
		if unpacked != nil {
			sdk.ProtocolCommon = unpacked
		}
	}

	// Handling Primitives
	if !model.Snippet.IsNull() && !model.Snippet.IsUnknown() {
		sdk.Snippet = model.Snippet.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.IkeGateways", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for IkeGateways ---
func packIkeGatewaysFromSdk(ctx context.Context, sdk network_services.IkeGateways) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.IkeGateways", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.IkeGateways
	var d diag.Diagnostics
	// MEGA FIX FOR MAP TYPE MISMATCH (NOT ALL MODELS MAY HAVE EncryptedValues)
	model.EncryptedValues = basetypes.NewMapNull(basetypes.StringType{})
	// Handling Objects
	// This is a regular nested object that has its own packer.
	// Logic for non-pointer / value-type nested objects
	if !reflect.ValueOf(sdk.Authentication).IsZero() {
		tflog.Debug(ctx, "Packing nested object for field Authentication")
		packed, d := packIkeGatewaysAuthenticationFromSdk(ctx, sdk.Authentication)
		diags.Append(d...)
		model.Authentication = packed
	} else {
		model.Authentication = basetypes.NewObjectNull(models.IkeGatewaysAuthentication{}.AttrTypes())
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
	if sdk.Id != nil {
		model.Id = basetypes.NewStringValue(*sdk.Id)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Id", "value": *sdk.Id})
	} else {
		model.Id = basetypes.NewStringNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.LocalAddress != nil {
		tflog.Debug(ctx, "Packing nested object for field LocalAddress")
		packed, d := packIkeGatewaysLocalAddressFromSdk(ctx, *sdk.LocalAddress)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "LocalAddress"})
		}
		model.LocalAddress = packed
	} else {
		model.LocalAddress = basetypes.NewObjectNull(models.IkeGatewaysLocalAddress{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.LocalId != nil {
		tflog.Debug(ctx, "Packing nested object for field LocalId")
		packed, d := packIkeGatewaysLocalIdFromSdk(ctx, *sdk.LocalId)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "LocalId"})
		}
		model.LocalId = packed
	} else {
		model.LocalId = basetypes.NewObjectNull(models.IkeGatewaysLocalId{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	// Handling Objects
	// This is a regular nested object that has its own packer.
	// Logic for non-pointer / value-type nested objects
	if !reflect.ValueOf(sdk.PeerAddress).IsZero() {
		tflog.Debug(ctx, "Packing nested object for field PeerAddress")
		packed, d := packIkeGatewaysPeerAddressFromSdk(ctx, sdk.PeerAddress)
		diags.Append(d...)
		model.PeerAddress = packed
	} else {
		model.PeerAddress = basetypes.NewObjectNull(models.IkeGatewaysPeerAddress{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.PeerId != nil {
		tflog.Debug(ctx, "Packing nested object for field PeerId")
		packed, d := packIkeGatewaysPeerIdFromSdk(ctx, *sdk.PeerId)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "PeerId"})
		}
		model.PeerId = packed
	} else {
		model.PeerId = basetypes.NewObjectNull(models.IkeGatewaysPeerId{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	// Logic for non-pointer / value-type nested objects
	if !reflect.ValueOf(sdk.Protocol).IsZero() {
		tflog.Debug(ctx, "Packing nested object for field Protocol")
		packed, d := packIkeGatewaysProtocolFromSdk(ctx, sdk.Protocol)
		diags.Append(d...)
		model.Protocol = packed
	} else {
		model.Protocol = basetypes.NewObjectNull(models.IkeGatewaysProtocol{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.ProtocolCommon != nil {
		tflog.Debug(ctx, "Packing nested object for field ProtocolCommon")
		packed, d := packIkeGatewaysProtocolCommonFromSdk(ctx, *sdk.ProtocolCommon)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "ProtocolCommon"})
		}
		model.ProtocolCommon = packed
	} else {
		model.ProtocolCommon = basetypes.NewObjectNull(models.IkeGatewaysProtocolCommon{}.AttrTypes())
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

	obj, d := types.ObjectValueFrom(ctx, models.IkeGateways{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.IkeGateways", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for IkeGateways ---
func unpackIkeGatewaysListToSdk(ctx context.Context, list types.List) ([]network_services.IkeGateways, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.IkeGateways")
	diags := diag.Diagnostics{}
	var data []models.IkeGateways
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.IkeGateways, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.IkeGateways{}.AttrTypes(), &item)
		unpacked, d := unpackIkeGatewaysToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.IkeGateways", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for IkeGateways ---
func packIkeGatewaysListFromSdk(ctx context.Context, sdks []network_services.IkeGateways) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.IkeGateways")
	diags := diag.Diagnostics{}
	var data []models.IkeGateways

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.IkeGateways
		obj, d := packIkeGatewaysFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.IkeGateways{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.IkeGateways", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.IkeGateways{}.AttrType(), data)
}

// --- Unpacker for IkeGatewaysAuthentication ---
func unpackIkeGatewaysAuthenticationToSdk(ctx context.Context, obj types.Object) (*network_services.IkeGatewaysAuthentication, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.IkeGatewaysAuthentication", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.IkeGatewaysAuthentication
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.IkeGatewaysAuthentication
	var d diag.Diagnostics
	// Handling Objects
	if !model.Certificate.IsNull() && !model.Certificate.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Certificate")
		unpacked, d := unpackIkeGatewaysAuthenticationCertificateToSdk(ctx, model.Certificate)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Certificate"})
		}
		if unpacked != nil {
			sdk.Certificate = unpacked
		}
	}

	// Handling Objects
	if !model.PreSharedKey.IsNull() && !model.PreSharedKey.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field PreSharedKey")
		unpacked, d := unpackIkeGatewaysAuthenticationPreSharedKeyToSdk(ctx, model.PreSharedKey)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "PreSharedKey"})
		}
		if unpacked != nil {
			sdk.PreSharedKey = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.IkeGatewaysAuthentication", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for IkeGatewaysAuthentication ---
func packIkeGatewaysAuthenticationFromSdk(ctx context.Context, sdk network_services.IkeGatewaysAuthentication) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.IkeGatewaysAuthentication", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.IkeGatewaysAuthentication
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Certificate != nil {
		tflog.Debug(ctx, "Packing nested object for field Certificate")
		packed, d := packIkeGatewaysAuthenticationCertificateFromSdk(ctx, *sdk.Certificate)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Certificate"})
		}
		model.Certificate = packed
	} else {
		model.Certificate = basetypes.NewObjectNull(models.IkeGatewaysAuthenticationCertificate{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.PreSharedKey != nil {
		tflog.Debug(ctx, "Packing nested object for field PreSharedKey")
		packed, d := packIkeGatewaysAuthenticationPreSharedKeyFromSdk(ctx, *sdk.PreSharedKey)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "PreSharedKey"})
		}
		model.PreSharedKey = packed
	} else {
		model.PreSharedKey = basetypes.NewObjectNull(models.IkeGatewaysAuthenticationPreSharedKey{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.IkeGatewaysAuthentication{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.IkeGatewaysAuthentication", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for IkeGatewaysAuthentication ---
func unpackIkeGatewaysAuthenticationListToSdk(ctx context.Context, list types.List) ([]network_services.IkeGatewaysAuthentication, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.IkeGatewaysAuthentication")
	diags := diag.Diagnostics{}
	var data []models.IkeGatewaysAuthentication
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.IkeGatewaysAuthentication, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.IkeGatewaysAuthentication{}.AttrTypes(), &item)
		unpacked, d := unpackIkeGatewaysAuthenticationToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.IkeGatewaysAuthentication", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for IkeGatewaysAuthentication ---
func packIkeGatewaysAuthenticationListFromSdk(ctx context.Context, sdks []network_services.IkeGatewaysAuthentication) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.IkeGatewaysAuthentication")
	diags := diag.Diagnostics{}
	var data []models.IkeGatewaysAuthentication

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.IkeGatewaysAuthentication
		obj, d := packIkeGatewaysAuthenticationFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.IkeGatewaysAuthentication{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.IkeGatewaysAuthentication", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.IkeGatewaysAuthentication{}.AttrType(), data)
}

// --- Unpacker for IkeGatewaysAuthenticationCertificate ---
func unpackIkeGatewaysAuthenticationCertificateToSdk(ctx context.Context, obj types.Object) (*network_services.IkeGatewaysAuthenticationCertificate, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.IkeGatewaysAuthenticationCertificate", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.IkeGatewaysAuthenticationCertificate
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.IkeGatewaysAuthenticationCertificate
	var d diag.Diagnostics
	// Handling Primitives
	if !model.AllowIdPayloadMismatch.IsNull() && !model.AllowIdPayloadMismatch.IsUnknown() {
		sdk.AllowIdPayloadMismatch = model.AllowIdPayloadMismatch.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "AllowIdPayloadMismatch", "value": *sdk.AllowIdPayloadMismatch})
	}

	// Handling Primitives
	if !model.CertificateProfile.IsNull() && !model.CertificateProfile.IsUnknown() {
		sdk.CertificateProfile = model.CertificateProfile.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "CertificateProfile", "value": *sdk.CertificateProfile})
	}

	// Handling Objects
	if !model.LocalCertificate.IsNull() && !model.LocalCertificate.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field LocalCertificate")
		unpacked, d := unpackIkeGatewaysAuthenticationCertificateLocalCertificateToSdk(ctx, model.LocalCertificate)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "LocalCertificate"})
		}
		if unpacked != nil {
			sdk.LocalCertificate = unpacked
		}
	}

	// Handling Primitives
	if !model.StrictValidationRevocation.IsNull() && !model.StrictValidationRevocation.IsUnknown() {
		sdk.StrictValidationRevocation = model.StrictValidationRevocation.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "StrictValidationRevocation", "value": *sdk.StrictValidationRevocation})
	}

	// Handling Primitives
	if !model.UseManagementAsSource.IsNull() && !model.UseManagementAsSource.IsUnknown() {
		sdk.UseManagementAsSource = model.UseManagementAsSource.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "UseManagementAsSource", "value": *sdk.UseManagementAsSource})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.IkeGatewaysAuthenticationCertificate", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for IkeGatewaysAuthenticationCertificate ---
func packIkeGatewaysAuthenticationCertificateFromSdk(ctx context.Context, sdk network_services.IkeGatewaysAuthenticationCertificate) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.IkeGatewaysAuthenticationCertificate", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.IkeGatewaysAuthenticationCertificate
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.AllowIdPayloadMismatch != nil {
		model.AllowIdPayloadMismatch = basetypes.NewBoolValue(*sdk.AllowIdPayloadMismatch)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "AllowIdPayloadMismatch", "value": *sdk.AllowIdPayloadMismatch})
	} else {
		model.AllowIdPayloadMismatch = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.CertificateProfile != nil {
		model.CertificateProfile = basetypes.NewStringValue(*sdk.CertificateProfile)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "CertificateProfile", "value": *sdk.CertificateProfile})
	} else {
		model.CertificateProfile = basetypes.NewStringNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.LocalCertificate != nil {
		tflog.Debug(ctx, "Packing nested object for field LocalCertificate")
		packed, d := packIkeGatewaysAuthenticationCertificateLocalCertificateFromSdk(ctx, *sdk.LocalCertificate)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "LocalCertificate"})
		}
		model.LocalCertificate = packed
	} else {
		model.LocalCertificate = basetypes.NewObjectNull(models.IkeGatewaysAuthenticationCertificateLocalCertificate{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.StrictValidationRevocation != nil {
		model.StrictValidationRevocation = basetypes.NewBoolValue(*sdk.StrictValidationRevocation)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "StrictValidationRevocation", "value": *sdk.StrictValidationRevocation})
	} else {
		model.StrictValidationRevocation = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.UseManagementAsSource != nil {
		model.UseManagementAsSource = basetypes.NewBoolValue(*sdk.UseManagementAsSource)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "UseManagementAsSource", "value": *sdk.UseManagementAsSource})
	} else {
		model.UseManagementAsSource = basetypes.NewBoolNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.IkeGatewaysAuthenticationCertificate{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.IkeGatewaysAuthenticationCertificate", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for IkeGatewaysAuthenticationCertificate ---
func unpackIkeGatewaysAuthenticationCertificateListToSdk(ctx context.Context, list types.List) ([]network_services.IkeGatewaysAuthenticationCertificate, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.IkeGatewaysAuthenticationCertificate")
	diags := diag.Diagnostics{}
	var data []models.IkeGatewaysAuthenticationCertificate
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.IkeGatewaysAuthenticationCertificate, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.IkeGatewaysAuthenticationCertificate{}.AttrTypes(), &item)
		unpacked, d := unpackIkeGatewaysAuthenticationCertificateToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.IkeGatewaysAuthenticationCertificate", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for IkeGatewaysAuthenticationCertificate ---
func packIkeGatewaysAuthenticationCertificateListFromSdk(ctx context.Context, sdks []network_services.IkeGatewaysAuthenticationCertificate) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.IkeGatewaysAuthenticationCertificate")
	diags := diag.Diagnostics{}
	var data []models.IkeGatewaysAuthenticationCertificate

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.IkeGatewaysAuthenticationCertificate
		obj, d := packIkeGatewaysAuthenticationCertificateFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.IkeGatewaysAuthenticationCertificate{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.IkeGatewaysAuthenticationCertificate", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.IkeGatewaysAuthenticationCertificate{}.AttrType(), data)
}

// --- Unpacker for IkeGatewaysAuthenticationCertificateLocalCertificate ---
func unpackIkeGatewaysAuthenticationCertificateLocalCertificateToSdk(ctx context.Context, obj types.Object) (*network_services.IkeGatewaysAuthenticationCertificateLocalCertificate, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.IkeGatewaysAuthenticationCertificateLocalCertificate", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.IkeGatewaysAuthenticationCertificateLocalCertificate
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.IkeGatewaysAuthenticationCertificateLocalCertificate
	var d diag.Diagnostics
	// Handling Primitives
	if !model.LocalCertificateName.IsNull() && !model.LocalCertificateName.IsUnknown() {
		sdk.LocalCertificateName = model.LocalCertificateName.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "LocalCertificateName", "value": *sdk.LocalCertificateName})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.IkeGatewaysAuthenticationCertificateLocalCertificate", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for IkeGatewaysAuthenticationCertificateLocalCertificate ---
func packIkeGatewaysAuthenticationCertificateLocalCertificateFromSdk(ctx context.Context, sdk network_services.IkeGatewaysAuthenticationCertificateLocalCertificate) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.IkeGatewaysAuthenticationCertificateLocalCertificate", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.IkeGatewaysAuthenticationCertificateLocalCertificate
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.LocalCertificateName != nil {
		model.LocalCertificateName = basetypes.NewStringValue(*sdk.LocalCertificateName)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "LocalCertificateName", "value": *sdk.LocalCertificateName})
	} else {
		model.LocalCertificateName = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.IkeGatewaysAuthenticationCertificateLocalCertificate{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.IkeGatewaysAuthenticationCertificateLocalCertificate", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for IkeGatewaysAuthenticationCertificateLocalCertificate ---
func unpackIkeGatewaysAuthenticationCertificateLocalCertificateListToSdk(ctx context.Context, list types.List) ([]network_services.IkeGatewaysAuthenticationCertificateLocalCertificate, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.IkeGatewaysAuthenticationCertificateLocalCertificate")
	diags := diag.Diagnostics{}
	var data []models.IkeGatewaysAuthenticationCertificateLocalCertificate
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.IkeGatewaysAuthenticationCertificateLocalCertificate, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.IkeGatewaysAuthenticationCertificateLocalCertificate{}.AttrTypes(), &item)
		unpacked, d := unpackIkeGatewaysAuthenticationCertificateLocalCertificateToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.IkeGatewaysAuthenticationCertificateLocalCertificate", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for IkeGatewaysAuthenticationCertificateLocalCertificate ---
func packIkeGatewaysAuthenticationCertificateLocalCertificateListFromSdk(ctx context.Context, sdks []network_services.IkeGatewaysAuthenticationCertificateLocalCertificate) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.IkeGatewaysAuthenticationCertificateLocalCertificate")
	diags := diag.Diagnostics{}
	var data []models.IkeGatewaysAuthenticationCertificateLocalCertificate

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.IkeGatewaysAuthenticationCertificateLocalCertificate
		obj, d := packIkeGatewaysAuthenticationCertificateLocalCertificateFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.IkeGatewaysAuthenticationCertificateLocalCertificate{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.IkeGatewaysAuthenticationCertificateLocalCertificate", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.IkeGatewaysAuthenticationCertificateLocalCertificate{}.AttrType(), data)
}

// --- Unpacker for IkeGatewaysAuthenticationPreSharedKey ---
func unpackIkeGatewaysAuthenticationPreSharedKeyToSdk(ctx context.Context, obj types.Object) (*network_services.IkeGatewaysAuthenticationPreSharedKey, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.IkeGatewaysAuthenticationPreSharedKey", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.IkeGatewaysAuthenticationPreSharedKey
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.IkeGatewaysAuthenticationPreSharedKey
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Key.IsNull() && !model.Key.IsUnknown() {
		sdk.Key = model.Key.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Key", "value": *sdk.Key})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.IkeGatewaysAuthenticationPreSharedKey", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for IkeGatewaysAuthenticationPreSharedKey ---
func packIkeGatewaysAuthenticationPreSharedKeyFromSdk(ctx context.Context, sdk network_services.IkeGatewaysAuthenticationPreSharedKey) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.IkeGatewaysAuthenticationPreSharedKey", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.IkeGatewaysAuthenticationPreSharedKey
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Key != nil {
		model.Key = basetypes.NewStringValue(*sdk.Key)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Key", "value": *sdk.Key})
	} else {
		model.Key = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.IkeGatewaysAuthenticationPreSharedKey{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.IkeGatewaysAuthenticationPreSharedKey", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for IkeGatewaysAuthenticationPreSharedKey ---
func unpackIkeGatewaysAuthenticationPreSharedKeyListToSdk(ctx context.Context, list types.List) ([]network_services.IkeGatewaysAuthenticationPreSharedKey, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.IkeGatewaysAuthenticationPreSharedKey")
	diags := diag.Diagnostics{}
	var data []models.IkeGatewaysAuthenticationPreSharedKey
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.IkeGatewaysAuthenticationPreSharedKey, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.IkeGatewaysAuthenticationPreSharedKey{}.AttrTypes(), &item)
		unpacked, d := unpackIkeGatewaysAuthenticationPreSharedKeyToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.IkeGatewaysAuthenticationPreSharedKey", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for IkeGatewaysAuthenticationPreSharedKey ---
func packIkeGatewaysAuthenticationPreSharedKeyListFromSdk(ctx context.Context, sdks []network_services.IkeGatewaysAuthenticationPreSharedKey) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.IkeGatewaysAuthenticationPreSharedKey")
	diags := diag.Diagnostics{}
	var data []models.IkeGatewaysAuthenticationPreSharedKey

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.IkeGatewaysAuthenticationPreSharedKey
		obj, d := packIkeGatewaysAuthenticationPreSharedKeyFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.IkeGatewaysAuthenticationPreSharedKey{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.IkeGatewaysAuthenticationPreSharedKey", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.IkeGatewaysAuthenticationPreSharedKey{}.AttrType(), data)
}

// --- Unpacker for IkeGatewaysLocalAddress ---
func unpackIkeGatewaysLocalAddressToSdk(ctx context.Context, obj types.Object) (*network_services.IkeGatewaysLocalAddress, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.IkeGatewaysLocalAddress", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.IkeGatewaysLocalAddress
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.IkeGatewaysLocalAddress
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Interface.IsNull() && !model.Interface.IsUnknown() {
		sdk.Interface = model.Interface.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Interface", "value": *sdk.Interface})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.IkeGatewaysLocalAddress", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for IkeGatewaysLocalAddress ---
func packIkeGatewaysLocalAddressFromSdk(ctx context.Context, sdk network_services.IkeGatewaysLocalAddress) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.IkeGatewaysLocalAddress", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.IkeGatewaysLocalAddress
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Interface != nil {
		model.Interface = basetypes.NewStringValue(*sdk.Interface)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Interface", "value": *sdk.Interface})
	} else {
		model.Interface = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.IkeGatewaysLocalAddress{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.IkeGatewaysLocalAddress", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for IkeGatewaysLocalAddress ---
func unpackIkeGatewaysLocalAddressListToSdk(ctx context.Context, list types.List) ([]network_services.IkeGatewaysLocalAddress, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.IkeGatewaysLocalAddress")
	diags := diag.Diagnostics{}
	var data []models.IkeGatewaysLocalAddress
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.IkeGatewaysLocalAddress, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.IkeGatewaysLocalAddress{}.AttrTypes(), &item)
		unpacked, d := unpackIkeGatewaysLocalAddressToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.IkeGatewaysLocalAddress", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for IkeGatewaysLocalAddress ---
func packIkeGatewaysLocalAddressListFromSdk(ctx context.Context, sdks []network_services.IkeGatewaysLocalAddress) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.IkeGatewaysLocalAddress")
	diags := diag.Diagnostics{}
	var data []models.IkeGatewaysLocalAddress

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.IkeGatewaysLocalAddress
		obj, d := packIkeGatewaysLocalAddressFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.IkeGatewaysLocalAddress{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.IkeGatewaysLocalAddress", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.IkeGatewaysLocalAddress{}.AttrType(), data)
}

// --- Unpacker for IkeGatewaysLocalId ---
func unpackIkeGatewaysLocalIdToSdk(ctx context.Context, obj types.Object) (*network_services.IkeGatewaysLocalId, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.IkeGatewaysLocalId", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.IkeGatewaysLocalId
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.IkeGatewaysLocalId
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Id.IsNull() && !model.Id.IsUnknown() {
		sdk.Id = model.Id.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Id", "value": *sdk.Id})
	}

	// Handling Primitives
	if !model.Type.IsNull() && !model.Type.IsUnknown() {
		sdk.Type = model.Type.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Type", "value": *sdk.Type})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.IkeGatewaysLocalId", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for IkeGatewaysLocalId ---
func packIkeGatewaysLocalIdFromSdk(ctx context.Context, sdk network_services.IkeGatewaysLocalId) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.IkeGatewaysLocalId", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.IkeGatewaysLocalId
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Id != nil {
		model.Id = basetypes.NewStringValue(*sdk.Id)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Id", "value": *sdk.Id})
	} else {
		model.Id = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Type != nil {
		model.Type = basetypes.NewStringValue(*sdk.Type)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Type", "value": *sdk.Type})
	} else {
		model.Type = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.IkeGatewaysLocalId{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.IkeGatewaysLocalId", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for IkeGatewaysLocalId ---
func unpackIkeGatewaysLocalIdListToSdk(ctx context.Context, list types.List) ([]network_services.IkeGatewaysLocalId, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.IkeGatewaysLocalId")
	diags := diag.Diagnostics{}
	var data []models.IkeGatewaysLocalId
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.IkeGatewaysLocalId, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.IkeGatewaysLocalId{}.AttrTypes(), &item)
		unpacked, d := unpackIkeGatewaysLocalIdToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.IkeGatewaysLocalId", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for IkeGatewaysLocalId ---
func packIkeGatewaysLocalIdListFromSdk(ctx context.Context, sdks []network_services.IkeGatewaysLocalId) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.IkeGatewaysLocalId")
	diags := diag.Diagnostics{}
	var data []models.IkeGatewaysLocalId

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.IkeGatewaysLocalId
		obj, d := packIkeGatewaysLocalIdFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.IkeGatewaysLocalId{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.IkeGatewaysLocalId", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.IkeGatewaysLocalId{}.AttrType(), data)
}

// --- Unpacker for IkeGatewaysPeerAddress ---
func unpackIkeGatewaysPeerAddressToSdk(ctx context.Context, obj types.Object) (*network_services.IkeGatewaysPeerAddress, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.IkeGatewaysPeerAddress", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.IkeGatewaysPeerAddress
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.IkeGatewaysPeerAddress
	var d diag.Diagnostics
	// Handling Typeless Objects
	if !model.Dynamic.IsNull() && !model.Dynamic.IsUnknown() {
		tflog.Debug(ctx, "Unpacking typeless object for field Dynamic")
		sdk.Dynamic = make(map[string]interface{})
	}

	// Handling Primitives
	if !model.Fqdn.IsNull() && !model.Fqdn.IsUnknown() {
		sdk.Fqdn = model.Fqdn.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Fqdn", "value": *sdk.Fqdn})
	}

	// Handling Primitives
	if !model.Ip.IsNull() && !model.Ip.IsUnknown() {
		sdk.Ip = model.Ip.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Ip", "value": *sdk.Ip})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.IkeGatewaysPeerAddress", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for IkeGatewaysPeerAddress ---
func packIkeGatewaysPeerAddressFromSdk(ctx context.Context, sdk network_services.IkeGatewaysPeerAddress) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.IkeGatewaysPeerAddress", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.IkeGatewaysPeerAddress
	var d diag.Diagnostics
	// Handling Objects
	// This is a marker object (e.g. CHAP: {}). We just need to create an empty, non-null object.
	if sdk.Dynamic != nil && !reflect.ValueOf(sdk.Dynamic).IsNil() {
		tflog.Debug(ctx, "Packing typeless object for field Dynamic")
		var d diag.Diagnostics
		// Create an empty object with no attributes, which signifies its presence.
		model.Dynamic, d = basetypes.NewObjectValue(map[string]attr.Type{}, map[string]attr.Value{})
		diags.Append(d...)
	} else {
		// Since this field is part of a oneOf, being nil means it's not selected.
		// We make the object null with an empty attribute map.
		model.Dynamic = basetypes.NewObjectNull(map[string]attr.Type{})
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Fqdn != nil {
		model.Fqdn = basetypes.NewStringValue(*sdk.Fqdn)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Fqdn", "value": *sdk.Fqdn})
	} else {
		model.Fqdn = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Ip != nil {
		model.Ip = basetypes.NewStringValue(*sdk.Ip)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Ip", "value": *sdk.Ip})
	} else {
		model.Ip = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.IkeGatewaysPeerAddress{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.IkeGatewaysPeerAddress", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for IkeGatewaysPeerAddress ---
func unpackIkeGatewaysPeerAddressListToSdk(ctx context.Context, list types.List) ([]network_services.IkeGatewaysPeerAddress, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.IkeGatewaysPeerAddress")
	diags := diag.Diagnostics{}
	var data []models.IkeGatewaysPeerAddress
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.IkeGatewaysPeerAddress, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.IkeGatewaysPeerAddress{}.AttrTypes(), &item)
		unpacked, d := unpackIkeGatewaysPeerAddressToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.IkeGatewaysPeerAddress", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for IkeGatewaysPeerAddress ---
func packIkeGatewaysPeerAddressListFromSdk(ctx context.Context, sdks []network_services.IkeGatewaysPeerAddress) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.IkeGatewaysPeerAddress")
	diags := diag.Diagnostics{}
	var data []models.IkeGatewaysPeerAddress

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.IkeGatewaysPeerAddress
		obj, d := packIkeGatewaysPeerAddressFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.IkeGatewaysPeerAddress{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.IkeGatewaysPeerAddress", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.IkeGatewaysPeerAddress{}.AttrType(), data)
}

// --- Unpacker for IkeGatewaysPeerId ---
func unpackIkeGatewaysPeerIdToSdk(ctx context.Context, obj types.Object) (*network_services.IkeGatewaysPeerId, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.IkeGatewaysPeerId", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.IkeGatewaysPeerId
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.IkeGatewaysPeerId
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Id.IsNull() && !model.Id.IsUnknown() {
		sdk.Id = model.Id.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Id", "value": *sdk.Id})
	}

	// Handling Primitives
	if !model.Type.IsNull() && !model.Type.IsUnknown() {
		sdk.Type = model.Type.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Type", "value": *sdk.Type})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.IkeGatewaysPeerId", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for IkeGatewaysPeerId ---
func packIkeGatewaysPeerIdFromSdk(ctx context.Context, sdk network_services.IkeGatewaysPeerId) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.IkeGatewaysPeerId", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.IkeGatewaysPeerId
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Id != nil {
		model.Id = basetypes.NewStringValue(*sdk.Id)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Id", "value": *sdk.Id})
	} else {
		model.Id = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Type != nil {
		model.Type = basetypes.NewStringValue(*sdk.Type)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Type", "value": *sdk.Type})
	} else {
		model.Type = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.IkeGatewaysPeerId{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.IkeGatewaysPeerId", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for IkeGatewaysPeerId ---
func unpackIkeGatewaysPeerIdListToSdk(ctx context.Context, list types.List) ([]network_services.IkeGatewaysPeerId, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.IkeGatewaysPeerId")
	diags := diag.Diagnostics{}
	var data []models.IkeGatewaysPeerId
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.IkeGatewaysPeerId, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.IkeGatewaysPeerId{}.AttrTypes(), &item)
		unpacked, d := unpackIkeGatewaysPeerIdToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.IkeGatewaysPeerId", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for IkeGatewaysPeerId ---
func packIkeGatewaysPeerIdListFromSdk(ctx context.Context, sdks []network_services.IkeGatewaysPeerId) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.IkeGatewaysPeerId")
	diags := diag.Diagnostics{}
	var data []models.IkeGatewaysPeerId

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.IkeGatewaysPeerId
		obj, d := packIkeGatewaysPeerIdFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.IkeGatewaysPeerId{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.IkeGatewaysPeerId", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.IkeGatewaysPeerId{}.AttrType(), data)
}

// --- Unpacker for IkeGatewaysProtocol ---
func unpackIkeGatewaysProtocolToSdk(ctx context.Context, obj types.Object) (*network_services.IkeGatewaysProtocol, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.IkeGatewaysProtocol", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.IkeGatewaysProtocol
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.IkeGatewaysProtocol
	var d diag.Diagnostics
	// Handling Objects
	if !model.Ikev1.IsNull() && !model.Ikev1.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Ikev1")
		unpacked, d := unpackIkeGatewaysProtocolIkev1ToSdk(ctx, model.Ikev1)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Ikev1"})
		}
		if unpacked != nil {
			sdk.Ikev1 = unpacked
		}
	}

	// Handling Objects
	if !model.Ikev2.IsNull() && !model.Ikev2.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Ikev2")
		unpacked, d := unpackIkeGatewaysProtocolIkev1ToSdk(ctx, model.Ikev2)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Ikev2"})
		}
		if unpacked != nil {
			sdk.Ikev2 = unpacked
		}
	}

	// Handling Primitives
	if !model.Version.IsNull() && !model.Version.IsUnknown() {
		sdk.Version = model.Version.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Version", "value": *sdk.Version})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.IkeGatewaysProtocol", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for IkeGatewaysProtocol ---
func packIkeGatewaysProtocolFromSdk(ctx context.Context, sdk network_services.IkeGatewaysProtocol) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.IkeGatewaysProtocol", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.IkeGatewaysProtocol
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Ikev1 != nil {
		tflog.Debug(ctx, "Packing nested object for field Ikev1")
		packed, d := packIkeGatewaysProtocolIkev1FromSdk(ctx, *sdk.Ikev1)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Ikev1"})
		}
		model.Ikev1 = packed
	} else {
		model.Ikev1 = basetypes.NewObjectNull(models.IkeGatewaysProtocolIkev1{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Ikev2 != nil {
		tflog.Debug(ctx, "Packing nested object for field Ikev2")
		packed, d := packIkeGatewaysProtocolIkev1FromSdk(ctx, *sdk.Ikev2)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Ikev2"})
		}
		model.Ikev2 = packed
	} else {
		model.Ikev2 = basetypes.NewObjectNull(models.IkeGatewaysProtocolIkev1{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Version != nil {
		model.Version = basetypes.NewStringValue(*sdk.Version)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Version", "value": *sdk.Version})
	} else {
		model.Version = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.IkeGatewaysProtocol{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.IkeGatewaysProtocol", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for IkeGatewaysProtocol ---
func unpackIkeGatewaysProtocolListToSdk(ctx context.Context, list types.List) ([]network_services.IkeGatewaysProtocol, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.IkeGatewaysProtocol")
	diags := diag.Diagnostics{}
	var data []models.IkeGatewaysProtocol
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.IkeGatewaysProtocol, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.IkeGatewaysProtocol{}.AttrTypes(), &item)
		unpacked, d := unpackIkeGatewaysProtocolToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.IkeGatewaysProtocol", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for IkeGatewaysProtocol ---
func packIkeGatewaysProtocolListFromSdk(ctx context.Context, sdks []network_services.IkeGatewaysProtocol) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.IkeGatewaysProtocol")
	diags := diag.Diagnostics{}
	var data []models.IkeGatewaysProtocol

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.IkeGatewaysProtocol
		obj, d := packIkeGatewaysProtocolFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.IkeGatewaysProtocol{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.IkeGatewaysProtocol", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.IkeGatewaysProtocol{}.AttrType(), data)
}

// --- Unpacker for IkeGatewaysProtocolIkev1 ---
func unpackIkeGatewaysProtocolIkev1ToSdk(ctx context.Context, obj types.Object) (*network_services.IkeGatewaysProtocolIkev1, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.IkeGatewaysProtocolIkev1", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.IkeGatewaysProtocolIkev1
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.IkeGatewaysProtocolIkev1
	var d diag.Diagnostics
	// Handling Objects
	if !model.Dpd.IsNull() && !model.Dpd.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Dpd")
		unpacked, d := unpackIkeGatewaysProtocolIkev1DpdToSdk(ctx, model.Dpd)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Dpd"})
		}
		if unpacked != nil {
			sdk.Dpd = unpacked
		}
	}

	// Handling Primitives
	if !model.IkeCryptoProfile.IsNull() && !model.IkeCryptoProfile.IsUnknown() {
		sdk.IkeCryptoProfile = model.IkeCryptoProfile.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "IkeCryptoProfile", "value": *sdk.IkeCryptoProfile})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.IkeGatewaysProtocolIkev1", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for IkeGatewaysProtocolIkev1 ---
func packIkeGatewaysProtocolIkev1FromSdk(ctx context.Context, sdk network_services.IkeGatewaysProtocolIkev1) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.IkeGatewaysProtocolIkev1", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.IkeGatewaysProtocolIkev1
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Dpd != nil {
		tflog.Debug(ctx, "Packing nested object for field Dpd")
		packed, d := packIkeGatewaysProtocolIkev1DpdFromSdk(ctx, *sdk.Dpd)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Dpd"})
		}
		model.Dpd = packed
	} else {
		model.Dpd = basetypes.NewObjectNull(models.IkeGatewaysProtocolIkev1Dpd{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.IkeCryptoProfile != nil {
		model.IkeCryptoProfile = basetypes.NewStringValue(*sdk.IkeCryptoProfile)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "IkeCryptoProfile", "value": *sdk.IkeCryptoProfile})
	} else {
		model.IkeCryptoProfile = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.IkeGatewaysProtocolIkev1{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.IkeGatewaysProtocolIkev1", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for IkeGatewaysProtocolIkev1 ---
func unpackIkeGatewaysProtocolIkev1ListToSdk(ctx context.Context, list types.List) ([]network_services.IkeGatewaysProtocolIkev1, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.IkeGatewaysProtocolIkev1")
	diags := diag.Diagnostics{}
	var data []models.IkeGatewaysProtocolIkev1
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.IkeGatewaysProtocolIkev1, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.IkeGatewaysProtocolIkev1{}.AttrTypes(), &item)
		unpacked, d := unpackIkeGatewaysProtocolIkev1ToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.IkeGatewaysProtocolIkev1", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for IkeGatewaysProtocolIkev1 ---
func packIkeGatewaysProtocolIkev1ListFromSdk(ctx context.Context, sdks []network_services.IkeGatewaysProtocolIkev1) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.IkeGatewaysProtocolIkev1")
	diags := diag.Diagnostics{}
	var data []models.IkeGatewaysProtocolIkev1

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.IkeGatewaysProtocolIkev1
		obj, d := packIkeGatewaysProtocolIkev1FromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.IkeGatewaysProtocolIkev1{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.IkeGatewaysProtocolIkev1", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.IkeGatewaysProtocolIkev1{}.AttrType(), data)
}

// --- Unpacker for IkeGatewaysProtocolIkev1Dpd ---
func unpackIkeGatewaysProtocolIkev1DpdToSdk(ctx context.Context, obj types.Object) (*network_services.IkeGatewaysProtocolIkev1Dpd, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.IkeGatewaysProtocolIkev1Dpd", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.IkeGatewaysProtocolIkev1Dpd
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.IkeGatewaysProtocolIkev1Dpd
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Enable.IsNull() && !model.Enable.IsUnknown() {
		sdk.Enable = model.Enable.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Enable", "value": *sdk.Enable})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.IkeGatewaysProtocolIkev1Dpd", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for IkeGatewaysProtocolIkev1Dpd ---
func packIkeGatewaysProtocolIkev1DpdFromSdk(ctx context.Context, sdk network_services.IkeGatewaysProtocolIkev1Dpd) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.IkeGatewaysProtocolIkev1Dpd", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.IkeGatewaysProtocolIkev1Dpd
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Enable != nil {
		model.Enable = basetypes.NewBoolValue(*sdk.Enable)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Enable", "value": *sdk.Enable})
	} else {
		model.Enable = basetypes.NewBoolNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.IkeGatewaysProtocolIkev1Dpd{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.IkeGatewaysProtocolIkev1Dpd", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for IkeGatewaysProtocolIkev1Dpd ---
func unpackIkeGatewaysProtocolIkev1DpdListToSdk(ctx context.Context, list types.List) ([]network_services.IkeGatewaysProtocolIkev1Dpd, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.IkeGatewaysProtocolIkev1Dpd")
	diags := diag.Diagnostics{}
	var data []models.IkeGatewaysProtocolIkev1Dpd
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.IkeGatewaysProtocolIkev1Dpd, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.IkeGatewaysProtocolIkev1Dpd{}.AttrTypes(), &item)
		unpacked, d := unpackIkeGatewaysProtocolIkev1DpdToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.IkeGatewaysProtocolIkev1Dpd", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for IkeGatewaysProtocolIkev1Dpd ---
func packIkeGatewaysProtocolIkev1DpdListFromSdk(ctx context.Context, sdks []network_services.IkeGatewaysProtocolIkev1Dpd) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.IkeGatewaysProtocolIkev1Dpd")
	diags := diag.Diagnostics{}
	var data []models.IkeGatewaysProtocolIkev1Dpd

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.IkeGatewaysProtocolIkev1Dpd
		obj, d := packIkeGatewaysProtocolIkev1DpdFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.IkeGatewaysProtocolIkev1Dpd{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.IkeGatewaysProtocolIkev1Dpd", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.IkeGatewaysProtocolIkev1Dpd{}.AttrType(), data)
}

// --- Unpacker for IkeGatewaysProtocolCommon ---
func unpackIkeGatewaysProtocolCommonToSdk(ctx context.Context, obj types.Object) (*network_services.IkeGatewaysProtocolCommon, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.IkeGatewaysProtocolCommon", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.IkeGatewaysProtocolCommon
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.IkeGatewaysProtocolCommon
	var d diag.Diagnostics
	// Handling Objects
	if !model.Fragmentation.IsNull() && !model.Fragmentation.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Fragmentation")
		unpacked, d := unpackIkeGatewaysProtocolCommonFragmentationToSdk(ctx, model.Fragmentation)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Fragmentation"})
		}
		if unpacked != nil {
			sdk.Fragmentation = unpacked
		}
	}

	// Handling Objects
	if !model.NatTraversal.IsNull() && !model.NatTraversal.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field NatTraversal")
		unpacked, d := unpackIkeGatewaysProtocolCommonNatTraversalToSdk(ctx, model.NatTraversal)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "NatTraversal"})
		}
		if unpacked != nil {
			sdk.NatTraversal = unpacked
		}
	}

	// Handling Primitives
	if !model.PassiveMode.IsNull() && !model.PassiveMode.IsUnknown() {
		sdk.PassiveMode = model.PassiveMode.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "PassiveMode", "value": *sdk.PassiveMode})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.IkeGatewaysProtocolCommon", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for IkeGatewaysProtocolCommon ---
func packIkeGatewaysProtocolCommonFromSdk(ctx context.Context, sdk network_services.IkeGatewaysProtocolCommon) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.IkeGatewaysProtocolCommon", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.IkeGatewaysProtocolCommon
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Fragmentation != nil {
		tflog.Debug(ctx, "Packing nested object for field Fragmentation")
		packed, d := packIkeGatewaysProtocolCommonFragmentationFromSdk(ctx, *sdk.Fragmentation)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Fragmentation"})
		}
		model.Fragmentation = packed
	} else {
		model.Fragmentation = basetypes.NewObjectNull(models.IkeGatewaysProtocolCommonFragmentation{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.NatTraversal != nil {
		tflog.Debug(ctx, "Packing nested object for field NatTraversal")
		packed, d := packIkeGatewaysProtocolCommonNatTraversalFromSdk(ctx, *sdk.NatTraversal)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "NatTraversal"})
		}
		model.NatTraversal = packed
	} else {
		model.NatTraversal = basetypes.NewObjectNull(models.IkeGatewaysProtocolCommonNatTraversal{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.PassiveMode != nil {
		model.PassiveMode = basetypes.NewBoolValue(*sdk.PassiveMode)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "PassiveMode", "value": *sdk.PassiveMode})
	} else {
		model.PassiveMode = basetypes.NewBoolNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.IkeGatewaysProtocolCommon{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.IkeGatewaysProtocolCommon", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for IkeGatewaysProtocolCommon ---
func unpackIkeGatewaysProtocolCommonListToSdk(ctx context.Context, list types.List) ([]network_services.IkeGatewaysProtocolCommon, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.IkeGatewaysProtocolCommon")
	diags := diag.Diagnostics{}
	var data []models.IkeGatewaysProtocolCommon
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.IkeGatewaysProtocolCommon, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.IkeGatewaysProtocolCommon{}.AttrTypes(), &item)
		unpacked, d := unpackIkeGatewaysProtocolCommonToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.IkeGatewaysProtocolCommon", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for IkeGatewaysProtocolCommon ---
func packIkeGatewaysProtocolCommonListFromSdk(ctx context.Context, sdks []network_services.IkeGatewaysProtocolCommon) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.IkeGatewaysProtocolCommon")
	diags := diag.Diagnostics{}
	var data []models.IkeGatewaysProtocolCommon

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.IkeGatewaysProtocolCommon
		obj, d := packIkeGatewaysProtocolCommonFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.IkeGatewaysProtocolCommon{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.IkeGatewaysProtocolCommon", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.IkeGatewaysProtocolCommon{}.AttrType(), data)
}

// --- Unpacker for IkeGatewaysProtocolCommonFragmentation ---
func unpackIkeGatewaysProtocolCommonFragmentationToSdk(ctx context.Context, obj types.Object) (*network_services.IkeGatewaysProtocolCommonFragmentation, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.IkeGatewaysProtocolCommonFragmentation", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.IkeGatewaysProtocolCommonFragmentation
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.IkeGatewaysProtocolCommonFragmentation
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Enable.IsNull() && !model.Enable.IsUnknown() {
		sdk.Enable = model.Enable.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Enable", "value": *sdk.Enable})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.IkeGatewaysProtocolCommonFragmentation", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for IkeGatewaysProtocolCommonFragmentation ---
func packIkeGatewaysProtocolCommonFragmentationFromSdk(ctx context.Context, sdk network_services.IkeGatewaysProtocolCommonFragmentation) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.IkeGatewaysProtocolCommonFragmentation", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.IkeGatewaysProtocolCommonFragmentation
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Enable != nil {
		model.Enable = basetypes.NewBoolValue(*sdk.Enable)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Enable", "value": *sdk.Enable})
	} else {
		model.Enable = basetypes.NewBoolNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.IkeGatewaysProtocolCommonFragmentation{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.IkeGatewaysProtocolCommonFragmentation", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for IkeGatewaysProtocolCommonFragmentation ---
func unpackIkeGatewaysProtocolCommonFragmentationListToSdk(ctx context.Context, list types.List) ([]network_services.IkeGatewaysProtocolCommonFragmentation, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.IkeGatewaysProtocolCommonFragmentation")
	diags := diag.Diagnostics{}
	var data []models.IkeGatewaysProtocolCommonFragmentation
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.IkeGatewaysProtocolCommonFragmentation, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.IkeGatewaysProtocolCommonFragmentation{}.AttrTypes(), &item)
		unpacked, d := unpackIkeGatewaysProtocolCommonFragmentationToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.IkeGatewaysProtocolCommonFragmentation", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for IkeGatewaysProtocolCommonFragmentation ---
func packIkeGatewaysProtocolCommonFragmentationListFromSdk(ctx context.Context, sdks []network_services.IkeGatewaysProtocolCommonFragmentation) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.IkeGatewaysProtocolCommonFragmentation")
	diags := diag.Diagnostics{}
	var data []models.IkeGatewaysProtocolCommonFragmentation

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.IkeGatewaysProtocolCommonFragmentation
		obj, d := packIkeGatewaysProtocolCommonFragmentationFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.IkeGatewaysProtocolCommonFragmentation{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.IkeGatewaysProtocolCommonFragmentation", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.IkeGatewaysProtocolCommonFragmentation{}.AttrType(), data)
}

// --- Unpacker for IkeGatewaysProtocolCommonNatTraversal ---
func unpackIkeGatewaysProtocolCommonNatTraversalToSdk(ctx context.Context, obj types.Object) (*network_services.IkeGatewaysProtocolCommonNatTraversal, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.IkeGatewaysProtocolCommonNatTraversal", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.IkeGatewaysProtocolCommonNatTraversal
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.IkeGatewaysProtocolCommonNatTraversal
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Enable.IsNull() && !model.Enable.IsUnknown() {
		sdk.Enable = model.Enable.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Enable", "value": *sdk.Enable})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.IkeGatewaysProtocolCommonNatTraversal", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for IkeGatewaysProtocolCommonNatTraversal ---
func packIkeGatewaysProtocolCommonNatTraversalFromSdk(ctx context.Context, sdk network_services.IkeGatewaysProtocolCommonNatTraversal) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.IkeGatewaysProtocolCommonNatTraversal", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.IkeGatewaysProtocolCommonNatTraversal
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Enable != nil {
		model.Enable = basetypes.NewBoolValue(*sdk.Enable)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Enable", "value": *sdk.Enable})
	} else {
		model.Enable = basetypes.NewBoolNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.IkeGatewaysProtocolCommonNatTraversal{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.IkeGatewaysProtocolCommonNatTraversal", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for IkeGatewaysProtocolCommonNatTraversal ---
func unpackIkeGatewaysProtocolCommonNatTraversalListToSdk(ctx context.Context, list types.List) ([]network_services.IkeGatewaysProtocolCommonNatTraversal, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.IkeGatewaysProtocolCommonNatTraversal")
	diags := diag.Diagnostics{}
	var data []models.IkeGatewaysProtocolCommonNatTraversal
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.IkeGatewaysProtocolCommonNatTraversal, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.IkeGatewaysProtocolCommonNatTraversal{}.AttrTypes(), &item)
		unpacked, d := unpackIkeGatewaysProtocolCommonNatTraversalToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.IkeGatewaysProtocolCommonNatTraversal", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for IkeGatewaysProtocolCommonNatTraversal ---
func packIkeGatewaysProtocolCommonNatTraversalListFromSdk(ctx context.Context, sdks []network_services.IkeGatewaysProtocolCommonNatTraversal) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.IkeGatewaysProtocolCommonNatTraversal")
	diags := diag.Diagnostics{}
	var data []models.IkeGatewaysProtocolCommonNatTraversal

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.IkeGatewaysProtocolCommonNatTraversal
		obj, d := packIkeGatewaysProtocolCommonNatTraversalFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.IkeGatewaysProtocolCommonNatTraversal{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.IkeGatewaysProtocolCommonNatTraversal", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.IkeGatewaysProtocolCommonNatTraversal{}.AttrType(), data)
}
