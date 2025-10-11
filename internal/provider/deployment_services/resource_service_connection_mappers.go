package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/deployment_services"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/deployment_services"
)

// serviceConnectionsSensitiveValuePatcher is an in-memory struct to temporarily store plaintext
// and encrypted values for sensitive fields during the Create/Update/Read workflows.
type serviceConnectionsSensitiveValuePatcher struct {
	bgp_peer_secret_plaintext     basetypes.StringValue
	bgp_peer_secret_encrypted     basetypes.StringValue
	protocol_bgp_secret_plaintext basetypes.StringValue
	protocol_bgp_secret_encrypted basetypes.StringValue
}

// populatePatcherFromState populates the patcher struct from the resource's state.
func (p *serviceConnectionsSensitiveValuePatcher) populatePatcherFromState(ctx context.Context, state models.ServiceConnections) diag.Diagnostics {
	var diags diag.Diagnostics
	if state.EncryptedValues.IsNull() || state.EncryptedValues.IsUnknown() {
		return diags
	}

	var ev map[string]string
	diags.Append(state.EncryptedValues.ElementsAs(ctx, &ev, false)...)
	if diags.HasError() {
		return diags
	}
	if val, ok := ev["bgp_peer_secret_plaintext"]; ok {
		p.bgp_peer_secret_plaintext = basetypes.NewStringValue(val)
	}
	if val, ok := ev["bgp_peer_secret_encrypted"]; ok {
		p.bgp_peer_secret_encrypted = basetypes.NewStringValue(val)
	}
	if val, ok := ev["protocol_bgp_secret_plaintext"]; ok {
		p.protocol_bgp_secret_plaintext = basetypes.NewStringValue(val)
	}
	if val, ok := ev["protocol_bgp_secret_encrypted"]; ok {
		p.protocol_bgp_secret_encrypted = basetypes.NewStringValue(val)
	}

	return diags
}

// populateEncryptedValuesMap returns a map of the patcher's values for saving to state.
func (p *serviceConnectionsSensitiveValuePatcher) populateEncryptedValuesMap() map[string]string {
	ev := make(map[string]string)
	if !p.bgp_peer_secret_plaintext.IsNull() {
		ev["bgp_peer_secret_plaintext"] = p.bgp_peer_secret_plaintext.ValueString()
	}
	if !p.bgp_peer_secret_encrypted.IsNull() {
		ev["bgp_peer_secret_encrypted"] = p.bgp_peer_secret_encrypted.ValueString()
	}
	if !p.protocol_bgp_secret_plaintext.IsNull() {
		ev["protocol_bgp_secret_plaintext"] = p.protocol_bgp_secret_plaintext.ValueString()
	}
	if !p.protocol_bgp_secret_encrypted.IsNull() {
		ev["protocol_bgp_secret_encrypted"] = p.protocol_bgp_secret_encrypted.ValueString()
	}
	return ev
}

// --- Unpacker for ServiceConnections ---
func unpackServiceConnectionsToSdk(ctx context.Context, obj types.Object) (*deployment_services.ServiceConnections, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ServiceConnections", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ServiceConnections
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk deployment_services.ServiceConnections
	var d diag.Diagnostics
	// Handling Primitives
	if !model.BackupSC.IsNull() && !model.BackupSC.IsUnknown() {
		sdk.BackupSC = model.BackupSC.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "BackupSC", "value": *sdk.BackupSC})
	}

	// Handling Objects
	if !model.BgpPeer.IsNull() && !model.BgpPeer.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field BgpPeer")
		unpacked, d := unpackServiceConnectionsBgpPeerToSdk(ctx, model.BgpPeer)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "BgpPeer"})
		}
		if unpacked != nil {
			sdk.BgpPeer = unpacked
		}
	}

	// Handling Primitives
	if !model.Id.IsNull() && !model.Id.IsUnknown() {
		sdk.Id = model.Id.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Id", "value": sdk.Id})
	}

	// Handling Primitives
	if !model.IpsecTunnel.IsNull() && !model.IpsecTunnel.IsUnknown() {
		sdk.IpsecTunnel = model.IpsecTunnel.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "IpsecTunnel", "value": sdk.IpsecTunnel})
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	// Handling Primitives
	if !model.NatPool.IsNull() && !model.NatPool.IsUnknown() {
		sdk.NatPool = model.NatPool.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "NatPool", "value": *sdk.NatPool})
	}

	// Handling Primitives
	if !model.NoExportCommunity.IsNull() && !model.NoExportCommunity.IsUnknown() {
		sdk.NoExportCommunity = model.NoExportCommunity.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "NoExportCommunity", "value": *sdk.NoExportCommunity})
	}

	// Handling Primitives
	if !model.OnboardingType.IsNull() && !model.OnboardingType.IsUnknown() {
		sdk.OnboardingType = model.OnboardingType.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "OnboardingType", "value": *sdk.OnboardingType})
	}

	// Handling Objects
	if !model.Protocol.IsNull() && !model.Protocol.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Protocol")
		unpacked, d := unpackServiceConnectionsProtocolToSdk(ctx, model.Protocol)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Protocol"})
		}
		if unpacked != nil {
			sdk.Protocol = unpacked
		}
	}

	// Handling Objects
	if !model.Qos.IsNull() && !model.Qos.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Qos")
		unpacked, d := unpackServiceConnectionsQosToSdk(ctx, model.Qos)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Qos"})
		}
		if unpacked != nil {
			sdk.Qos = unpacked
		}
	}

	// Handling Primitives
	if !model.Region.IsNull() && !model.Region.IsUnknown() {
		sdk.Region = model.Region.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Region", "value": sdk.Region})
	}

	// Handling Primitives
	if !model.SecondaryIpsecTunnel.IsNull() && !model.SecondaryIpsecTunnel.IsUnknown() {
		sdk.SecondaryIpsecTunnel = model.SecondaryIpsecTunnel.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "SecondaryIpsecTunnel", "value": *sdk.SecondaryIpsecTunnel})
	}

	// Handling Primitives
	if !model.SourceNat.IsNull() && !model.SourceNat.IsUnknown() {
		sdk.SourceNat = model.SourceNat.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "SourceNat", "value": *sdk.SourceNat})
	}

	// Handling Lists
	if !model.Subnets.IsNull() && !model.Subnets.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Subnets")
		diags.Append(model.Subnets.ElementsAs(ctx, &sdk.Subnets, false)...)
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ServiceConnections", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ServiceConnections ---
func packServiceConnectionsFromSdk(ctx context.Context, sdk deployment_services.ServiceConnections) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ServiceConnections", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ServiceConnections
	var d diag.Diagnostics
	// MEGA FIX FOR MAP TYPE MISMATCH (NOT ALL MODELS MAY HAVE EncryptedValues)
	model.EncryptedValues = basetypes.NewMapNull(basetypes.StringType{})
	// Handling Primitives
	// Standard primitive packing
	if sdk.BackupSC != nil {
		model.BackupSC = basetypes.NewStringValue(*sdk.BackupSC)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "BackupSC", "value": *sdk.BackupSC})
	} else {
		model.BackupSC = basetypes.NewStringNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.BgpPeer != nil {
		tflog.Debug(ctx, "Packing nested object for field BgpPeer")
		packed, d := packServiceConnectionsBgpPeerFromSdk(ctx, *sdk.BgpPeer)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "BgpPeer"})
		}
		model.BgpPeer = packed
	} else {
		model.BgpPeer = basetypes.NewObjectNull(models.ServiceConnectionsBgpPeer{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	model.Id = basetypes.NewStringValue(sdk.Id)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Id", "value": sdk.Id})
	// Handling Primitives
	// Standard primitive packing
	model.IpsecTunnel = basetypes.NewStringValue(sdk.IpsecTunnel)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "IpsecTunnel", "value": sdk.IpsecTunnel})
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	// Handling Primitives
	// Standard primitive packing
	if sdk.NatPool != nil {
		model.NatPool = basetypes.NewStringValue(*sdk.NatPool)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "NatPool", "value": *sdk.NatPool})
	} else {
		model.NatPool = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.NoExportCommunity != nil {
		model.NoExportCommunity = basetypes.NewStringValue(*sdk.NoExportCommunity)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "NoExportCommunity", "value": *sdk.NoExportCommunity})
	} else {
		model.NoExportCommunity = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.OnboardingType != nil {
		model.OnboardingType = basetypes.NewStringValue(*sdk.OnboardingType)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "OnboardingType", "value": *sdk.OnboardingType})
	} else {
		model.OnboardingType = basetypes.NewStringNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Protocol != nil {
		tflog.Debug(ctx, "Packing nested object for field Protocol")
		packed, d := packServiceConnectionsProtocolFromSdk(ctx, *sdk.Protocol)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Protocol"})
		}
		model.Protocol = packed
	} else {
		model.Protocol = basetypes.NewObjectNull(models.ServiceConnectionsProtocol{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Qos != nil {
		tflog.Debug(ctx, "Packing nested object for field Qos")
		packed, d := packServiceConnectionsQosFromSdk(ctx, *sdk.Qos)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Qos"})
		}
		model.Qos = packed
	} else {
		model.Qos = basetypes.NewObjectNull(models.ServiceConnectionsQos{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	model.Region = basetypes.NewStringValue(sdk.Region)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Region", "value": sdk.Region})
	// Handling Primitives
	// Standard primitive packing
	if sdk.SecondaryIpsecTunnel != nil {
		model.SecondaryIpsecTunnel = basetypes.NewStringValue(*sdk.SecondaryIpsecTunnel)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "SecondaryIpsecTunnel", "value": *sdk.SecondaryIpsecTunnel})
	} else {
		model.SecondaryIpsecTunnel = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.SourceNat != nil {
		model.SourceNat = basetypes.NewBoolValue(*sdk.SourceNat)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "SourceNat", "value": *sdk.SourceNat})
	} else {
		model.SourceNat = basetypes.NewBoolNull()
	}
	// Handling Lists
	if sdk.Subnets != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Subnets")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Subnets, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Subnets)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Subnets = basetypes.NewListNull(elemType)
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ServiceConnections{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ServiceConnections", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ServiceConnections ---
func unpackServiceConnectionsListToSdk(ctx context.Context, list types.List) ([]deployment_services.ServiceConnections, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ServiceConnections")
	diags := diag.Diagnostics{}
	var data []models.ServiceConnections
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]deployment_services.ServiceConnections, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ServiceConnections{}.AttrTypes(), &item)
		unpacked, d := unpackServiceConnectionsToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ServiceConnections", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ServiceConnections ---
func packServiceConnectionsListFromSdk(ctx context.Context, sdks []deployment_services.ServiceConnections) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ServiceConnections")
	diags := diag.Diagnostics{}
	var data []models.ServiceConnections

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ServiceConnections
		obj, d := packServiceConnectionsFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ServiceConnections{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ServiceConnections", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ServiceConnections{}.AttrType(), data)
}

// --- Unpacker for ServiceConnectionsBgpPeer ---
func unpackServiceConnectionsBgpPeerToSdk(ctx context.Context, obj types.Object) (*deployment_services.ServiceConnectionsBgpPeer, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ServiceConnectionsBgpPeer", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ServiceConnectionsBgpPeer
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk deployment_services.ServiceConnectionsBgpPeer
	var d diag.Diagnostics
	// Handling Primitives
	if !model.LocalIpAddress.IsNull() && !model.LocalIpAddress.IsUnknown() {
		sdk.LocalIpAddress = model.LocalIpAddress.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "LocalIpAddress", "value": *sdk.LocalIpAddress})
	}

	// Handling Primitives
	if !model.LocalIpv6Address.IsNull() && !model.LocalIpv6Address.IsUnknown() {
		sdk.LocalIpv6Address = model.LocalIpv6Address.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "LocalIpv6Address", "value": *sdk.LocalIpv6Address})
	}

	// Handling Primitives
	if !model.PeerIpAddress.IsNull() && !model.PeerIpAddress.IsUnknown() {
		sdk.PeerIpAddress = model.PeerIpAddress.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "PeerIpAddress", "value": *sdk.PeerIpAddress})
	}

	// Handling Primitives
	if !model.PeerIpv6Address.IsNull() && !model.PeerIpv6Address.IsUnknown() {
		sdk.PeerIpv6Address = model.PeerIpv6Address.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "PeerIpv6Address", "value": *sdk.PeerIpv6Address})
	}

	// Handling Primitives
	if !model.Secret.IsNull() && !model.Secret.IsUnknown() {
		sdk.Secret = model.Secret.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Secret", "value": *sdk.Secret})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ServiceConnectionsBgpPeer", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ServiceConnectionsBgpPeer ---
func packServiceConnectionsBgpPeerFromSdk(ctx context.Context, sdk deployment_services.ServiceConnectionsBgpPeer) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ServiceConnectionsBgpPeer", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ServiceConnectionsBgpPeer
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.LocalIpAddress != nil {
		model.LocalIpAddress = basetypes.NewStringValue(*sdk.LocalIpAddress)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "LocalIpAddress", "value": *sdk.LocalIpAddress})
	} else {
		model.LocalIpAddress = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.LocalIpv6Address != nil {
		model.LocalIpv6Address = basetypes.NewStringValue(*sdk.LocalIpv6Address)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "LocalIpv6Address", "value": *sdk.LocalIpv6Address})
	} else {
		model.LocalIpv6Address = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.PeerIpAddress != nil {
		model.PeerIpAddress = basetypes.NewStringValue(*sdk.PeerIpAddress)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "PeerIpAddress", "value": *sdk.PeerIpAddress})
	} else {
		model.PeerIpAddress = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.PeerIpv6Address != nil {
		model.PeerIpv6Address = basetypes.NewStringValue(*sdk.PeerIpv6Address)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "PeerIpv6Address", "value": *sdk.PeerIpv6Address})
	} else {
		model.PeerIpv6Address = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Secret != nil {
		model.Secret = basetypes.NewStringValue(*sdk.Secret)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Secret", "value": *sdk.Secret})
	} else {
		model.Secret = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ServiceConnectionsBgpPeer{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ServiceConnectionsBgpPeer", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ServiceConnectionsBgpPeer ---
func unpackServiceConnectionsBgpPeerListToSdk(ctx context.Context, list types.List) ([]deployment_services.ServiceConnectionsBgpPeer, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ServiceConnectionsBgpPeer")
	diags := diag.Diagnostics{}
	var data []models.ServiceConnectionsBgpPeer
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]deployment_services.ServiceConnectionsBgpPeer, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ServiceConnectionsBgpPeer{}.AttrTypes(), &item)
		unpacked, d := unpackServiceConnectionsBgpPeerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ServiceConnectionsBgpPeer", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ServiceConnectionsBgpPeer ---
func packServiceConnectionsBgpPeerListFromSdk(ctx context.Context, sdks []deployment_services.ServiceConnectionsBgpPeer) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ServiceConnectionsBgpPeer")
	diags := diag.Diagnostics{}
	var data []models.ServiceConnectionsBgpPeer

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ServiceConnectionsBgpPeer
		obj, d := packServiceConnectionsBgpPeerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ServiceConnectionsBgpPeer{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ServiceConnectionsBgpPeer", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ServiceConnectionsBgpPeer{}.AttrType(), data)
}

// --- Unpacker for ServiceConnectionsProtocol ---
func unpackServiceConnectionsProtocolToSdk(ctx context.Context, obj types.Object) (*deployment_services.ServiceConnectionsProtocol, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ServiceConnectionsProtocol", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ServiceConnectionsProtocol
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk deployment_services.ServiceConnectionsProtocol
	var d diag.Diagnostics
	// Handling Objects
	if !model.Bgp.IsNull() && !model.Bgp.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Bgp")
		unpacked, d := unpackServiceConnectionsProtocolBgpToSdk(ctx, model.Bgp)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Bgp"})
		}
		if unpacked != nil {
			sdk.Bgp = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ServiceConnectionsProtocol", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ServiceConnectionsProtocol ---
func packServiceConnectionsProtocolFromSdk(ctx context.Context, sdk deployment_services.ServiceConnectionsProtocol) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ServiceConnectionsProtocol", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ServiceConnectionsProtocol
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Bgp != nil {
		tflog.Debug(ctx, "Packing nested object for field Bgp")
		packed, d := packServiceConnectionsProtocolBgpFromSdk(ctx, *sdk.Bgp)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Bgp"})
		}
		model.Bgp = packed
	} else {
		model.Bgp = basetypes.NewObjectNull(models.ServiceConnectionsProtocolBgp{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ServiceConnectionsProtocol{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ServiceConnectionsProtocol", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ServiceConnectionsProtocol ---
func unpackServiceConnectionsProtocolListToSdk(ctx context.Context, list types.List) ([]deployment_services.ServiceConnectionsProtocol, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ServiceConnectionsProtocol")
	diags := diag.Diagnostics{}
	var data []models.ServiceConnectionsProtocol
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]deployment_services.ServiceConnectionsProtocol, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ServiceConnectionsProtocol{}.AttrTypes(), &item)
		unpacked, d := unpackServiceConnectionsProtocolToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ServiceConnectionsProtocol", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ServiceConnectionsProtocol ---
func packServiceConnectionsProtocolListFromSdk(ctx context.Context, sdks []deployment_services.ServiceConnectionsProtocol) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ServiceConnectionsProtocol")
	diags := diag.Diagnostics{}
	var data []models.ServiceConnectionsProtocol

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ServiceConnectionsProtocol
		obj, d := packServiceConnectionsProtocolFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ServiceConnectionsProtocol{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ServiceConnectionsProtocol", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ServiceConnectionsProtocol{}.AttrType(), data)
}

// --- Unpacker for ServiceConnectionsProtocolBgp ---
func unpackServiceConnectionsProtocolBgpToSdk(ctx context.Context, obj types.Object) (*deployment_services.ServiceConnectionsProtocolBgp, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ServiceConnectionsProtocolBgp", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ServiceConnectionsProtocolBgp
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk deployment_services.ServiceConnectionsProtocolBgp
	var d diag.Diagnostics
	// Handling Primitives
	if !model.DoNotExportRoutes.IsNull() && !model.DoNotExportRoutes.IsUnknown() {
		sdk.DoNotExportRoutes = model.DoNotExportRoutes.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DoNotExportRoutes", "value": *sdk.DoNotExportRoutes})
	}

	// Handling Primitives
	if !model.Enable.IsNull() && !model.Enable.IsUnknown() {
		sdk.Enable = model.Enable.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Enable", "value": *sdk.Enable})
	}

	// Handling Primitives
	if !model.FastFailover.IsNull() && !model.FastFailover.IsUnknown() {
		sdk.FastFailover = model.FastFailover.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "FastFailover", "value": *sdk.FastFailover})
	}

	// Handling Primitives
	if !model.LocalIpAddress.IsNull() && !model.LocalIpAddress.IsUnknown() {
		sdk.LocalIpAddress = model.LocalIpAddress.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "LocalIpAddress", "value": *sdk.LocalIpAddress})
	}

	// Handling Primitives
	if !model.OriginateDefaultRoute.IsNull() && !model.OriginateDefaultRoute.IsUnknown() {
		sdk.OriginateDefaultRoute = model.OriginateDefaultRoute.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "OriginateDefaultRoute", "value": *sdk.OriginateDefaultRoute})
	}

	// Handling Primitives
	if !model.PeerAs.IsNull() && !model.PeerAs.IsUnknown() {
		sdk.PeerAs = model.PeerAs.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "PeerAs", "value": *sdk.PeerAs})
	}

	// Handling Primitives
	if !model.PeerIpAddress.IsNull() && !model.PeerIpAddress.IsUnknown() {
		sdk.PeerIpAddress = model.PeerIpAddress.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "PeerIpAddress", "value": *sdk.PeerIpAddress})
	}

	// Handling Primitives
	if !model.Secret.IsNull() && !model.Secret.IsUnknown() {
		sdk.Secret = model.Secret.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Secret", "value": *sdk.Secret})
	}

	// Handling Primitives
	if !model.SummarizeMobileUserRoutes.IsNull() && !model.SummarizeMobileUserRoutes.IsUnknown() {
		sdk.SummarizeMobileUserRoutes = model.SummarizeMobileUserRoutes.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "SummarizeMobileUserRoutes", "value": *sdk.SummarizeMobileUserRoutes})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ServiceConnectionsProtocolBgp", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ServiceConnectionsProtocolBgp ---
func packServiceConnectionsProtocolBgpFromSdk(ctx context.Context, sdk deployment_services.ServiceConnectionsProtocolBgp) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ServiceConnectionsProtocolBgp", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ServiceConnectionsProtocolBgp
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.DoNotExportRoutes != nil {
		model.DoNotExportRoutes = basetypes.NewBoolValue(*sdk.DoNotExportRoutes)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DoNotExportRoutes", "value": *sdk.DoNotExportRoutes})
	} else {
		model.DoNotExportRoutes = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Enable != nil {
		model.Enable = basetypes.NewBoolValue(*sdk.Enable)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Enable", "value": *sdk.Enable})
	} else {
		model.Enable = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.FastFailover != nil {
		model.FastFailover = basetypes.NewBoolValue(*sdk.FastFailover)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "FastFailover", "value": *sdk.FastFailover})
	} else {
		model.FastFailover = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.LocalIpAddress != nil {
		model.LocalIpAddress = basetypes.NewStringValue(*sdk.LocalIpAddress)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "LocalIpAddress", "value": *sdk.LocalIpAddress})
	} else {
		model.LocalIpAddress = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.OriginateDefaultRoute != nil {
		model.OriginateDefaultRoute = basetypes.NewBoolValue(*sdk.OriginateDefaultRoute)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "OriginateDefaultRoute", "value": *sdk.OriginateDefaultRoute})
	} else {
		model.OriginateDefaultRoute = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.PeerAs != nil {
		model.PeerAs = basetypes.NewStringValue(*sdk.PeerAs)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "PeerAs", "value": *sdk.PeerAs})
	} else {
		model.PeerAs = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.PeerIpAddress != nil {
		model.PeerIpAddress = basetypes.NewStringValue(*sdk.PeerIpAddress)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "PeerIpAddress", "value": *sdk.PeerIpAddress})
	} else {
		model.PeerIpAddress = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Secret != nil {
		model.Secret = basetypes.NewStringValue(*sdk.Secret)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Secret", "value": *sdk.Secret})
	} else {
		model.Secret = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.SummarizeMobileUserRoutes != nil {
		model.SummarizeMobileUserRoutes = basetypes.NewBoolValue(*sdk.SummarizeMobileUserRoutes)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "SummarizeMobileUserRoutes", "value": *sdk.SummarizeMobileUserRoutes})
	} else {
		model.SummarizeMobileUserRoutes = basetypes.NewBoolNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ServiceConnectionsProtocolBgp{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ServiceConnectionsProtocolBgp", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ServiceConnectionsProtocolBgp ---
func unpackServiceConnectionsProtocolBgpListToSdk(ctx context.Context, list types.List) ([]deployment_services.ServiceConnectionsProtocolBgp, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ServiceConnectionsProtocolBgp")
	diags := diag.Diagnostics{}
	var data []models.ServiceConnectionsProtocolBgp
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]deployment_services.ServiceConnectionsProtocolBgp, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ServiceConnectionsProtocolBgp{}.AttrTypes(), &item)
		unpacked, d := unpackServiceConnectionsProtocolBgpToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ServiceConnectionsProtocolBgp", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ServiceConnectionsProtocolBgp ---
func packServiceConnectionsProtocolBgpListFromSdk(ctx context.Context, sdks []deployment_services.ServiceConnectionsProtocolBgp) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ServiceConnectionsProtocolBgp")
	diags := diag.Diagnostics{}
	var data []models.ServiceConnectionsProtocolBgp

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ServiceConnectionsProtocolBgp
		obj, d := packServiceConnectionsProtocolBgpFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ServiceConnectionsProtocolBgp{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ServiceConnectionsProtocolBgp", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ServiceConnectionsProtocolBgp{}.AttrType(), data)
}

// --- Unpacker for ServiceConnectionsQos ---
func unpackServiceConnectionsQosToSdk(ctx context.Context, obj types.Object) (*deployment_services.ServiceConnectionsQos, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ServiceConnectionsQos", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ServiceConnectionsQos
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk deployment_services.ServiceConnectionsQos
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Enable.IsNull() && !model.Enable.IsUnknown() {
		sdk.Enable = model.Enable.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Enable", "value": *sdk.Enable})
	}

	// Handling Primitives
	if !model.QosProfile.IsNull() && !model.QosProfile.IsUnknown() {
		sdk.QosProfile = model.QosProfile.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "QosProfile", "value": *sdk.QosProfile})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ServiceConnectionsQos", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ServiceConnectionsQos ---
func packServiceConnectionsQosFromSdk(ctx context.Context, sdk deployment_services.ServiceConnectionsQos) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ServiceConnectionsQos", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ServiceConnectionsQos
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Enable != nil {
		model.Enable = basetypes.NewBoolValue(*sdk.Enable)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Enable", "value": *sdk.Enable})
	} else {
		model.Enable = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.QosProfile != nil {
		model.QosProfile = basetypes.NewStringValue(*sdk.QosProfile)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "QosProfile", "value": *sdk.QosProfile})
	} else {
		model.QosProfile = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ServiceConnectionsQos{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ServiceConnectionsQos", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ServiceConnectionsQos ---
func unpackServiceConnectionsQosListToSdk(ctx context.Context, list types.List) ([]deployment_services.ServiceConnectionsQos, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ServiceConnectionsQos")
	diags := diag.Diagnostics{}
	var data []models.ServiceConnectionsQos
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]deployment_services.ServiceConnectionsQos, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ServiceConnectionsQos{}.AttrTypes(), &item)
		unpacked, d := unpackServiceConnectionsQosToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ServiceConnectionsQos", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ServiceConnectionsQos ---
func packServiceConnectionsQosListFromSdk(ctx context.Context, sdks []deployment_services.ServiceConnectionsQos) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ServiceConnectionsQos")
	diags := diag.Diagnostics{}
	var data []models.ServiceConnectionsQos

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ServiceConnectionsQos
		obj, d := packServiceConnectionsQosFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ServiceConnectionsQos{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ServiceConnectionsQos", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ServiceConnectionsQos{}.AttrType(), data)
}
