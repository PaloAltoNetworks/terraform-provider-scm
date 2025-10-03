package provider

import (
	"context"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/deployment_services"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/deployment_services"
)

// remoteNetworksSensitiveValuePatcher is an in-memory struct to temporarily store plaintext
// and encrypted values for sensitive fields during the Create/Update/Read workflows.
type remoteNetworksSensitiveValuePatcher struct {
	protocol_bgp_secret_plaintext      basetypes.StringValue
	protocol_bgp_secret_encrypted      basetypes.StringValue
	protocol_bgp_peer_secret_plaintext basetypes.StringValue
	protocol_bgp_peer_secret_encrypted basetypes.StringValue
}

// populatePatcherFromState populates the patcher struct from the resource's state.
func (p *remoteNetworksSensitiveValuePatcher) populatePatcherFromState(ctx context.Context, state models.RemoteNetworks) diag.Diagnostics {
	var diags diag.Diagnostics
	if state.EncryptedValues.IsNull() || state.EncryptedValues.IsUnknown() {
		return diags
	}

	var ev map[string]string
	diags.Append(state.EncryptedValues.ElementsAs(ctx, &ev, false)...)
	if diags.HasError() {
		return diags
	}
	if val, ok := ev["protocol_bgp_secret_plaintext"]; ok {
		p.protocol_bgp_secret_plaintext = basetypes.NewStringValue(val)
	}
	if val, ok := ev["protocol_bgp_secret_encrypted"]; ok {
		p.protocol_bgp_secret_encrypted = basetypes.NewStringValue(val)
	}
	if val, ok := ev["protocol_bgp_peer_secret_plaintext"]; ok {
		p.protocol_bgp_peer_secret_plaintext = basetypes.NewStringValue(val)
	}
	if val, ok := ev["protocol_bgp_peer_secret_encrypted"]; ok {
		p.protocol_bgp_peer_secret_encrypted = basetypes.NewStringValue(val)
	}

	return diags
}

// populateEncryptedValuesMap returns a map of the patcher's values for saving to state.
func (p *remoteNetworksSensitiveValuePatcher) populateEncryptedValuesMap() map[string]string {
	ev := make(map[string]string)
	if !p.protocol_bgp_secret_plaintext.IsNull() {
		ev["protocol_bgp_secret_plaintext"] = p.protocol_bgp_secret_plaintext.ValueString()
	}
	if !p.protocol_bgp_secret_encrypted.IsNull() {
		ev["protocol_bgp_secret_encrypted"] = p.protocol_bgp_secret_encrypted.ValueString()
	}
	if !p.protocol_bgp_peer_secret_plaintext.IsNull() {
		ev["protocol_bgp_peer_secret_plaintext"] = p.protocol_bgp_peer_secret_plaintext.ValueString()
	}
	if !p.protocol_bgp_peer_secret_encrypted.IsNull() {
		ev["protocol_bgp_peer_secret_encrypted"] = p.protocol_bgp_peer_secret_encrypted.ValueString()
	}
	return ev
}

// --- Unpacker for RemoteNetworks ---
func unpackRemoteNetworksToSdk(ctx context.Context, obj types.Object) (*deployment_services.RemoteNetworks, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.RemoteNetworks", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.RemoteNetworks
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk deployment_services.RemoteNetworks
	var d diag.Diagnostics
	// Handling Primitives
	if !model.EcmpLoadBalancing.IsNull() && !model.EcmpLoadBalancing.IsUnknown() {
		sdk.EcmpLoadBalancing = model.EcmpLoadBalancing.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "EcmpLoadBalancing", "value": *sdk.EcmpLoadBalancing})
	}

	// Handling Lists
	if !model.EcmpTunnels.IsNull() && !model.EcmpTunnels.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field EcmpTunnels")
		unpacked, d := unpackRemoteNetworksEcmpTunnelsInnerListToSdk(ctx, model.EcmpTunnels)
		diags.Append(d...)
		sdk.EcmpTunnels = unpacked
	}

	// Handling Primitives
	if !model.Folder.IsNull() && !model.Folder.IsUnknown() {
		sdk.Folder = model.Folder.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Folder", "value": sdk.Folder})
	}

	// Handling Primitives
	if !model.Id.IsNull() && !model.Id.IsUnknown() {
		sdk.Id = model.Id.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Id", "value": sdk.Id})
	}

	// Handling Primitives
	if !model.IpsecTunnel.IsNull() && !model.IpsecTunnel.IsUnknown() {
		sdk.IpsecTunnel = model.IpsecTunnel.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "IpsecTunnel", "value": *sdk.IpsecTunnel})
	}

	// Handling Primitives
	if !model.LicenseType.IsNull() && !model.LicenseType.IsUnknown() {
		sdk.LicenseType = model.LicenseType.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "LicenseType", "value": sdk.LicenseType})
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	// Handling Objects
	if !model.Protocol.IsNull() && !model.Protocol.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Protocol")
		unpacked, d := unpackRemoteNetworksProtocolToSdk(ctx, model.Protocol)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Protocol"})
		}
		if unpacked != nil {
			sdk.Protocol = unpacked
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
	if !model.SpnName.IsNull() && !model.SpnName.IsUnknown() {
		sdk.SpnName = model.SpnName.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "SpnName", "value": *sdk.SpnName})
	}

	// Handling Lists
	if !model.Subnets.IsNull() && !model.Subnets.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Subnets")
		diags.Append(model.Subnets.ElementsAs(ctx, &sdk.Subnets, false)...)
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.RemoteNetworks", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for RemoteNetworks ---
func packRemoteNetworksFromSdk(ctx context.Context, sdk deployment_services.RemoteNetworks) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.RemoteNetworks", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.RemoteNetworks
	var d diag.Diagnostics
	// MEGA FIX FOR MAP TYPE MISMATCH (NOT ALL MODELS MAY HAVE EncryptedValues)
	model.EncryptedValues = basetypes.NewMapNull(basetypes.StringType{})
	// Handling Primitives
	// Standard primitive packing
	if sdk.EcmpLoadBalancing != nil {
		model.EcmpLoadBalancing = basetypes.NewStringValue(*sdk.EcmpLoadBalancing)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "EcmpLoadBalancing", "value": *sdk.EcmpLoadBalancing})
	} else {
		model.EcmpLoadBalancing = basetypes.NewStringNull()
	}
	// Handling Lists
	if sdk.EcmpTunnels != nil {
		tflog.Debug(ctx, "Packing list of objects for field EcmpTunnels")
		packed, d := packRemoteNetworksEcmpTunnelsInnerListFromSdk(ctx, sdk.EcmpTunnels)
		diags.Append(d...)
		model.EcmpTunnels = packed
	} else {
		model.EcmpTunnels = basetypes.NewListNull(models.RemoteNetworksEcmpTunnelsInner{}.AttrType())
	}
	// Handling Primitives
	// Standard primitive packing
	model.Folder = basetypes.NewStringValue(sdk.Folder)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Folder", "value": sdk.Folder})
	// Handling Primitives
	// Standard primitive packing
	model.Id = basetypes.NewStringValue(sdk.Id)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Id", "value": sdk.Id})
	// Handling Primitives
	// Standard primitive packing
	if sdk.IpsecTunnel != nil {
		model.IpsecTunnel = basetypes.NewStringValue(*sdk.IpsecTunnel)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "IpsecTunnel", "value": *sdk.IpsecTunnel})
	} else {
		model.IpsecTunnel = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	model.LicenseType = basetypes.NewStringValue(sdk.LicenseType)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "LicenseType", "value": sdk.LicenseType})
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Protocol != nil {
		tflog.Debug(ctx, "Packing nested object for field Protocol")
		packed, d := packRemoteNetworksProtocolFromSdk(ctx, *sdk.Protocol)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Protocol"})
		}
		model.Protocol = packed
	} else {
		model.Protocol = basetypes.NewObjectNull(models.RemoteNetworksProtocol{}.AttrTypes())
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
	if sdk.SpnName != nil {
		model.SpnName = basetypes.NewStringValue(*sdk.SpnName)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "SpnName", "value": *sdk.SpnName})
	} else {
		model.SpnName = basetypes.NewStringNull()
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

	obj, d := types.ObjectValueFrom(ctx, models.RemoteNetworks{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.RemoteNetworks", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for RemoteNetworks ---
func unpackRemoteNetworksListToSdk(ctx context.Context, list types.List) ([]deployment_services.RemoteNetworks, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.RemoteNetworks")
	diags := diag.Diagnostics{}
	var data []models.RemoteNetworks
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]deployment_services.RemoteNetworks, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.RemoteNetworks{}.AttrTypes(), &item)
		unpacked, d := unpackRemoteNetworksToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.RemoteNetworks", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for RemoteNetworks ---
func packRemoteNetworksListFromSdk(ctx context.Context, sdks []deployment_services.RemoteNetworks) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.RemoteNetworks")
	diags := diag.Diagnostics{}
	var data []models.RemoteNetworks

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.RemoteNetworks
		obj, d := packRemoteNetworksFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.RemoteNetworks{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.RemoteNetworks", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.RemoteNetworks{}.AttrType(), data)
}

// --- Unpacker for RemoteNetworksEcmpTunnelsInner ---
func unpackRemoteNetworksEcmpTunnelsInnerToSdk(ctx context.Context, obj types.Object) (*deployment_services.RemoteNetworksEcmpTunnelsInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.RemoteNetworksEcmpTunnelsInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.RemoteNetworksEcmpTunnelsInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk deployment_services.RemoteNetworksEcmpTunnelsInner
	var d diag.Diagnostics
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

	// Handling Objects
	if !model.Protocol.IsNull() && !model.Protocol.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Protocol")
		unpacked, d := unpackRemoteNetworksEcmpTunnelsInnerProtocolToSdk(ctx, model.Protocol)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Protocol"})
		}
		if unpacked != nil {
			sdk.Protocol = *unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.RemoteNetworksEcmpTunnelsInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for RemoteNetworksEcmpTunnelsInner ---
func packRemoteNetworksEcmpTunnelsInnerFromSdk(ctx context.Context, sdk deployment_services.RemoteNetworksEcmpTunnelsInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.RemoteNetworksEcmpTunnelsInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.RemoteNetworksEcmpTunnelsInner
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.IpsecTunnel = basetypes.NewStringValue(sdk.IpsecTunnel)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "IpsecTunnel", "value": sdk.IpsecTunnel})
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	// Handling Objects
	// This is a regular nested object that has its own packer.
	// Logic for non-pointer / value-type nested objects
	if !reflect.ValueOf(sdk.Protocol).IsZero() {
		tflog.Debug(ctx, "Packing nested object for field Protocol")
		packed, d := packRemoteNetworksEcmpTunnelsInnerProtocolFromSdk(ctx, sdk.Protocol)
		diags.Append(d...)
		model.Protocol = packed
	} else {
		model.Protocol = basetypes.NewObjectNull(models.RemoteNetworksEcmpTunnelsInnerProtocol{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.RemoteNetworksEcmpTunnelsInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.RemoteNetworksEcmpTunnelsInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for RemoteNetworksEcmpTunnelsInner ---
func unpackRemoteNetworksEcmpTunnelsInnerListToSdk(ctx context.Context, list types.List) ([]deployment_services.RemoteNetworksEcmpTunnelsInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.RemoteNetworksEcmpTunnelsInner")
	diags := diag.Diagnostics{}
	var data []models.RemoteNetworksEcmpTunnelsInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]deployment_services.RemoteNetworksEcmpTunnelsInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.RemoteNetworksEcmpTunnelsInner{}.AttrTypes(), &item)
		unpacked, d := unpackRemoteNetworksEcmpTunnelsInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.RemoteNetworksEcmpTunnelsInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for RemoteNetworksEcmpTunnelsInner ---
func packRemoteNetworksEcmpTunnelsInnerListFromSdk(ctx context.Context, sdks []deployment_services.RemoteNetworksEcmpTunnelsInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.RemoteNetworksEcmpTunnelsInner")
	diags := diag.Diagnostics{}
	var data []models.RemoteNetworksEcmpTunnelsInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.RemoteNetworksEcmpTunnelsInner
		obj, d := packRemoteNetworksEcmpTunnelsInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.RemoteNetworksEcmpTunnelsInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.RemoteNetworksEcmpTunnelsInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.RemoteNetworksEcmpTunnelsInner{}.AttrType(), data)
}

// --- Unpacker for RemoteNetworksEcmpTunnelsInnerProtocol ---
func unpackRemoteNetworksEcmpTunnelsInnerProtocolToSdk(ctx context.Context, obj types.Object) (*deployment_services.RemoteNetworksEcmpTunnelsInnerProtocol, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.RemoteNetworksEcmpTunnelsInnerProtocol", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.RemoteNetworksEcmpTunnelsInnerProtocol
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk deployment_services.RemoteNetworksEcmpTunnelsInnerProtocol
	var d diag.Diagnostics
	// Handling Objects
	if !model.Bgp.IsNull() && !model.Bgp.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Bgp")
		unpacked, d := unpackRemoteNetworksProtocolBgpToSdk(ctx, model.Bgp)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Bgp"})
		}
		if unpacked != nil {
			sdk.Bgp = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.RemoteNetworksEcmpTunnelsInnerProtocol", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for RemoteNetworksEcmpTunnelsInnerProtocol ---
func packRemoteNetworksEcmpTunnelsInnerProtocolFromSdk(ctx context.Context, sdk deployment_services.RemoteNetworksEcmpTunnelsInnerProtocol) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.RemoteNetworksEcmpTunnelsInnerProtocol", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.RemoteNetworksEcmpTunnelsInnerProtocol
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Bgp != nil {
		tflog.Debug(ctx, "Packing nested object for field Bgp")
		packed, d := packRemoteNetworksProtocolBgpFromSdk(ctx, *sdk.Bgp)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Bgp"})
		}
		model.Bgp = packed
	} else {
		model.Bgp = basetypes.NewObjectNull(models.RemoteNetworksProtocolBgp{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.RemoteNetworksEcmpTunnelsInnerProtocol{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.RemoteNetworksEcmpTunnelsInnerProtocol", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for RemoteNetworksEcmpTunnelsInnerProtocol ---
func unpackRemoteNetworksEcmpTunnelsInnerProtocolListToSdk(ctx context.Context, list types.List) ([]deployment_services.RemoteNetworksEcmpTunnelsInnerProtocol, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.RemoteNetworksEcmpTunnelsInnerProtocol")
	diags := diag.Diagnostics{}
	var data []models.RemoteNetworksEcmpTunnelsInnerProtocol
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]deployment_services.RemoteNetworksEcmpTunnelsInnerProtocol, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.RemoteNetworksEcmpTunnelsInnerProtocol{}.AttrTypes(), &item)
		unpacked, d := unpackRemoteNetworksEcmpTunnelsInnerProtocolToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.RemoteNetworksEcmpTunnelsInnerProtocol", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for RemoteNetworksEcmpTunnelsInnerProtocol ---
func packRemoteNetworksEcmpTunnelsInnerProtocolListFromSdk(ctx context.Context, sdks []deployment_services.RemoteNetworksEcmpTunnelsInnerProtocol) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.RemoteNetworksEcmpTunnelsInnerProtocol")
	diags := diag.Diagnostics{}
	var data []models.RemoteNetworksEcmpTunnelsInnerProtocol

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.RemoteNetworksEcmpTunnelsInnerProtocol
		obj, d := packRemoteNetworksEcmpTunnelsInnerProtocolFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.RemoteNetworksEcmpTunnelsInnerProtocol{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.RemoteNetworksEcmpTunnelsInnerProtocol", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.RemoteNetworksEcmpTunnelsInnerProtocol{}.AttrType(), data)
}

// --- Unpacker for RemoteNetworksProtocolBgp ---
func unpackRemoteNetworksProtocolBgpToSdk(ctx context.Context, obj types.Object) (*deployment_services.RemoteNetworksProtocolBgp, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.RemoteNetworksProtocolBgp", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.RemoteNetworksProtocolBgp
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk deployment_services.RemoteNetworksProtocolBgp
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
	if !model.PeeringType.IsNull() && !model.PeeringType.IsUnknown() {
		sdk.PeeringType = model.PeeringType.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "PeeringType", "value": *sdk.PeeringType})
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

	tflog.Debug(ctx, "Exiting unpack helper for models.RemoteNetworksProtocolBgp", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for RemoteNetworksProtocolBgp ---
func packRemoteNetworksProtocolBgpFromSdk(ctx context.Context, sdk deployment_services.RemoteNetworksProtocolBgp) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.RemoteNetworksProtocolBgp", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.RemoteNetworksProtocolBgp
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
	if sdk.PeeringType != nil {
		model.PeeringType = basetypes.NewStringValue(*sdk.PeeringType)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "PeeringType", "value": *sdk.PeeringType})
	} else {
		model.PeeringType = basetypes.NewStringNull()
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

	obj, d := types.ObjectValueFrom(ctx, models.RemoteNetworksProtocolBgp{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.RemoteNetworksProtocolBgp", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for RemoteNetworksProtocolBgp ---
func unpackRemoteNetworksProtocolBgpListToSdk(ctx context.Context, list types.List) ([]deployment_services.RemoteNetworksProtocolBgp, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.RemoteNetworksProtocolBgp")
	diags := diag.Diagnostics{}
	var data []models.RemoteNetworksProtocolBgp
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]deployment_services.RemoteNetworksProtocolBgp, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.RemoteNetworksProtocolBgp{}.AttrTypes(), &item)
		unpacked, d := unpackRemoteNetworksProtocolBgpToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.RemoteNetworksProtocolBgp", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for RemoteNetworksProtocolBgp ---
func packRemoteNetworksProtocolBgpListFromSdk(ctx context.Context, sdks []deployment_services.RemoteNetworksProtocolBgp) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.RemoteNetworksProtocolBgp")
	diags := diag.Diagnostics{}
	var data []models.RemoteNetworksProtocolBgp

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.RemoteNetworksProtocolBgp
		obj, d := packRemoteNetworksProtocolBgpFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.RemoteNetworksProtocolBgp{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.RemoteNetworksProtocolBgp", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.RemoteNetworksProtocolBgp{}.AttrType(), data)
}

// --- Unpacker for RemoteNetworksProtocol ---
func unpackRemoteNetworksProtocolToSdk(ctx context.Context, obj types.Object) (*deployment_services.RemoteNetworksProtocol, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.RemoteNetworksProtocol", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.RemoteNetworksProtocol
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk deployment_services.RemoteNetworksProtocol
	var d diag.Diagnostics
	// Handling Objects
	if !model.Bgp.IsNull() && !model.Bgp.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Bgp")
		unpacked, d := unpackRemoteNetworksProtocolBgpToSdk(ctx, model.Bgp)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Bgp"})
		}
		if unpacked != nil {
			sdk.Bgp = unpacked
		}
	}

	// Handling Objects
	if !model.BgpPeer.IsNull() && !model.BgpPeer.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field BgpPeer")
		unpacked, d := unpackRemoteNetworksProtocolBgpPeerToSdk(ctx, model.BgpPeer)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "BgpPeer"})
		}
		if unpacked != nil {
			sdk.BgpPeer = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.RemoteNetworksProtocol", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for RemoteNetworksProtocol ---
func packRemoteNetworksProtocolFromSdk(ctx context.Context, sdk deployment_services.RemoteNetworksProtocol) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.RemoteNetworksProtocol", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.RemoteNetworksProtocol
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Bgp != nil {
		tflog.Debug(ctx, "Packing nested object for field Bgp")
		packed, d := packRemoteNetworksProtocolBgpFromSdk(ctx, *sdk.Bgp)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Bgp"})
		}
		model.Bgp = packed
	} else {
		model.Bgp = basetypes.NewObjectNull(models.RemoteNetworksProtocolBgp{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.BgpPeer != nil {
		tflog.Debug(ctx, "Packing nested object for field BgpPeer")
		packed, d := packRemoteNetworksProtocolBgpPeerFromSdk(ctx, *sdk.BgpPeer)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "BgpPeer"})
		}
		model.BgpPeer = packed
	} else {
		model.BgpPeer = basetypes.NewObjectNull(models.RemoteNetworksProtocolBgpPeer{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.RemoteNetworksProtocol{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.RemoteNetworksProtocol", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for RemoteNetworksProtocol ---
func unpackRemoteNetworksProtocolListToSdk(ctx context.Context, list types.List) ([]deployment_services.RemoteNetworksProtocol, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.RemoteNetworksProtocol")
	diags := diag.Diagnostics{}
	var data []models.RemoteNetworksProtocol
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]deployment_services.RemoteNetworksProtocol, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.RemoteNetworksProtocol{}.AttrTypes(), &item)
		unpacked, d := unpackRemoteNetworksProtocolToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.RemoteNetworksProtocol", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for RemoteNetworksProtocol ---
func packRemoteNetworksProtocolListFromSdk(ctx context.Context, sdks []deployment_services.RemoteNetworksProtocol) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.RemoteNetworksProtocol")
	diags := diag.Diagnostics{}
	var data []models.RemoteNetworksProtocol

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.RemoteNetworksProtocol
		obj, d := packRemoteNetworksProtocolFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.RemoteNetworksProtocol{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.RemoteNetworksProtocol", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.RemoteNetworksProtocol{}.AttrType(), data)
}

// --- Unpacker for RemoteNetworksProtocolBgpPeer ---
func unpackRemoteNetworksProtocolBgpPeerToSdk(ctx context.Context, obj types.Object) (*deployment_services.RemoteNetworksProtocolBgpPeer, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.RemoteNetworksProtocolBgpPeer", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.RemoteNetworksProtocolBgpPeer
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk deployment_services.RemoteNetworksProtocolBgpPeer
	var d diag.Diagnostics
	// Handling Primitives
	if !model.LocalIpAddress.IsNull() && !model.LocalIpAddress.IsUnknown() {
		sdk.LocalIpAddress = model.LocalIpAddress.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "LocalIpAddress", "value": *sdk.LocalIpAddress})
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

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.RemoteNetworksProtocolBgpPeer", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for RemoteNetworksProtocolBgpPeer ---
func packRemoteNetworksProtocolBgpPeerFromSdk(ctx context.Context, sdk deployment_services.RemoteNetworksProtocolBgpPeer) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.RemoteNetworksProtocolBgpPeer", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.RemoteNetworksProtocolBgpPeer
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
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.RemoteNetworksProtocolBgpPeer{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.RemoteNetworksProtocolBgpPeer", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for RemoteNetworksProtocolBgpPeer ---
func unpackRemoteNetworksProtocolBgpPeerListToSdk(ctx context.Context, list types.List) ([]deployment_services.RemoteNetworksProtocolBgpPeer, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.RemoteNetworksProtocolBgpPeer")
	diags := diag.Diagnostics{}
	var data []models.RemoteNetworksProtocolBgpPeer
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]deployment_services.RemoteNetworksProtocolBgpPeer, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.RemoteNetworksProtocolBgpPeer{}.AttrTypes(), &item)
		unpacked, d := unpackRemoteNetworksProtocolBgpPeerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.RemoteNetworksProtocolBgpPeer", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for RemoteNetworksProtocolBgpPeer ---
func packRemoteNetworksProtocolBgpPeerListFromSdk(ctx context.Context, sdks []deployment_services.RemoteNetworksProtocolBgpPeer) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.RemoteNetworksProtocolBgpPeer")
	diags := diag.Diagnostics{}
	var data []models.RemoteNetworksProtocolBgpPeer

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.RemoteNetworksProtocolBgpPeer
		obj, d := packRemoteNetworksProtocolBgpPeerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.RemoteNetworksProtocolBgpPeer{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.RemoteNetworksProtocolBgpPeer", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.RemoteNetworksProtocolBgpPeer{}.AttrType(), data)
}
