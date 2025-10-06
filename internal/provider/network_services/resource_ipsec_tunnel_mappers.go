package provider

import (
	"context"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/network_services"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/network_services"
)

// --- Unpacker for IpsecTunnels ---
func unpackIpsecTunnelsToSdk(ctx context.Context, obj types.Object) (*network_services.IpsecTunnels, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.IpsecTunnels", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.IpsecTunnels
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.IpsecTunnels
	var d diag.Diagnostics
	// Handling Primitives
	if !model.AntiReplay.IsNull() && !model.AntiReplay.IsUnknown() {
		sdk.AntiReplay = model.AntiReplay.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "AntiReplay", "value": *sdk.AntiReplay})
	}

	// Handling Objects
	if !model.AutoKey.IsNull() && !model.AutoKey.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field AutoKey")
		unpacked, d := unpackIpsecTunnelsAutoKeyToSdk(ctx, model.AutoKey)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "AutoKey"})
		}
		if unpacked != nil {
			sdk.AutoKey = *unpacked
		}
	}

	// Handling Primitives
	if !model.CopyTos.IsNull() && !model.CopyTos.IsUnknown() {
		sdk.CopyTos = model.CopyTos.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "CopyTos", "value": *sdk.CopyTos})
	}

	// Handling Primitives
	if !model.Device.IsNull() && !model.Device.IsUnknown() {
		sdk.Device = model.Device.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Device", "value": *sdk.Device})
	}

	// Handling Primitives
	if !model.EnableGreEncapsulation.IsNull() && !model.EnableGreEncapsulation.IsUnknown() {
		sdk.EnableGreEncapsulation = model.EnableGreEncapsulation.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "EnableGreEncapsulation", "value": *sdk.EnableGreEncapsulation})
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

	// Handling Primitives
	if !model.TunnelInterface.IsNull() && !model.TunnelInterface.IsUnknown() {
		sdk.TunnelInterface = model.TunnelInterface.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "TunnelInterface", "value": *sdk.TunnelInterface})
	}

	// Handling Objects
	if !model.TunnelMonitor.IsNull() && !model.TunnelMonitor.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field TunnelMonitor")
		unpacked, d := unpackIpsecTunnelsTunnelMonitorToSdk(ctx, model.TunnelMonitor)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "TunnelMonitor"})
		}
		if unpacked != nil {
			sdk.TunnelMonitor = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.IpsecTunnels", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for IpsecTunnels ---
func packIpsecTunnelsFromSdk(ctx context.Context, sdk network_services.IpsecTunnels) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.IpsecTunnels", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.IpsecTunnels
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.AntiReplay != nil {
		model.AntiReplay = basetypes.NewBoolValue(*sdk.AntiReplay)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "AntiReplay", "value": *sdk.AntiReplay})
	} else {
		model.AntiReplay = basetypes.NewBoolNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	// Logic for non-pointer / value-type nested objects
	if !reflect.ValueOf(sdk.AutoKey).IsZero() {
		tflog.Debug(ctx, "Packing nested object for field AutoKey")
		packed, d := packIpsecTunnelsAutoKeyFromSdk(ctx, sdk.AutoKey)
		diags.Append(d...)
		model.AutoKey = packed
	} else {
		model.AutoKey = basetypes.NewObjectNull(models.IpsecTunnelsAutoKey{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.CopyTos != nil {
		model.CopyTos = basetypes.NewBoolValue(*sdk.CopyTos)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "CopyTos", "value": *sdk.CopyTos})
	} else {
		model.CopyTos = basetypes.NewBoolNull()
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
	if sdk.EnableGreEncapsulation != nil {
		model.EnableGreEncapsulation = basetypes.NewBoolValue(*sdk.EnableGreEncapsulation)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "EnableGreEncapsulation", "value": *sdk.EnableGreEncapsulation})
	} else {
		model.EnableGreEncapsulation = basetypes.NewBoolNull()
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
	// Handling Primitives
	// Standard primitive packing
	if sdk.TunnelInterface != nil {
		model.TunnelInterface = basetypes.NewStringValue(*sdk.TunnelInterface)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "TunnelInterface", "value": *sdk.TunnelInterface})
	} else {
		model.TunnelInterface = basetypes.NewStringNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.TunnelMonitor != nil {
		tflog.Debug(ctx, "Packing nested object for field TunnelMonitor")
		packed, d := packIpsecTunnelsTunnelMonitorFromSdk(ctx, *sdk.TunnelMonitor)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "TunnelMonitor"})
		}
		model.TunnelMonitor = packed
	} else {
		model.TunnelMonitor = basetypes.NewObjectNull(models.IpsecTunnelsTunnelMonitor{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.IpsecTunnels{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.IpsecTunnels", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for IpsecTunnels ---
func unpackIpsecTunnelsListToSdk(ctx context.Context, list types.List) ([]network_services.IpsecTunnels, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.IpsecTunnels")
	diags := diag.Diagnostics{}
	var data []models.IpsecTunnels
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.IpsecTunnels, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.IpsecTunnels{}.AttrTypes(), &item)
		unpacked, d := unpackIpsecTunnelsToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.IpsecTunnels", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for IpsecTunnels ---
func packIpsecTunnelsListFromSdk(ctx context.Context, sdks []network_services.IpsecTunnels) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.IpsecTunnels")
	diags := diag.Diagnostics{}
	var data []models.IpsecTunnels

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.IpsecTunnels
		obj, d := packIpsecTunnelsFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.IpsecTunnels{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.IpsecTunnels", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.IpsecTunnels{}.AttrType(), data)
}

// --- Unpacker for IpsecTunnelsAutoKey ---
func unpackIpsecTunnelsAutoKeyToSdk(ctx context.Context, obj types.Object) (*network_services.IpsecTunnelsAutoKey, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.IpsecTunnelsAutoKey", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.IpsecTunnelsAutoKey
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.IpsecTunnelsAutoKey
	var d diag.Diagnostics
	// Handling Lists
	if !model.IkeGateway.IsNull() && !model.IkeGateway.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field IkeGateway")
		unpacked, d := unpackIpsecTunnelsAutoKeyIkeGatewayInnerListToSdk(ctx, model.IkeGateway)
		diags.Append(d...)
		sdk.IkeGateway = unpacked
	}

	// Handling Primitives
	if !model.IpsecCryptoProfile.IsNull() && !model.IpsecCryptoProfile.IsUnknown() {
		sdk.IpsecCryptoProfile = model.IpsecCryptoProfile.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "IpsecCryptoProfile", "value": sdk.IpsecCryptoProfile})
	}

	// Handling Lists
	if !model.ProxyId.IsNull() && !model.ProxyId.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field ProxyId")
		unpacked, d := unpackIpsecTunnelsAutoKeyProxyIdInnerListToSdk(ctx, model.ProxyId)
		diags.Append(d...)
		sdk.ProxyId = unpacked
	}

	// Handling Lists
	if !model.ProxyIdV6.IsNull() && !model.ProxyIdV6.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field ProxyIdV6")
		unpacked, d := unpackIpsecTunnelsAutoKeyProxyIdInnerListToSdk(ctx, model.ProxyIdV6)
		diags.Append(d...)
		sdk.ProxyIdV6 = unpacked
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.IpsecTunnelsAutoKey", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for IpsecTunnelsAutoKey ---
func packIpsecTunnelsAutoKeyFromSdk(ctx context.Context, sdk network_services.IpsecTunnelsAutoKey) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.IpsecTunnelsAutoKey", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.IpsecTunnelsAutoKey
	var d diag.Diagnostics
	// Handling Lists
	if sdk.IkeGateway != nil {
		tflog.Debug(ctx, "Packing list of objects for field IkeGateway")
		packed, d := packIpsecTunnelsAutoKeyIkeGatewayInnerListFromSdk(ctx, sdk.IkeGateway)
		diags.Append(d...)
		model.IkeGateway = packed
	} else {
		model.IkeGateway = basetypes.NewListNull(models.IpsecTunnelsAutoKeyIkeGatewayInner{}.AttrType())
	}
	// Handling Primitives
	// Standard primitive packing
	model.IpsecCryptoProfile = basetypes.NewStringValue(sdk.IpsecCryptoProfile)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "IpsecCryptoProfile", "value": sdk.IpsecCryptoProfile})
	// Handling Lists
	if sdk.ProxyId != nil {
		tflog.Debug(ctx, "Packing list of objects for field ProxyId")
		packed, d := packIpsecTunnelsAutoKeyProxyIdInnerListFromSdk(ctx, sdk.ProxyId)
		diags.Append(d...)
		model.ProxyId = packed
	} else {
		model.ProxyId = basetypes.NewListNull(models.IpsecTunnelsAutoKeyProxyIdInner{}.AttrType())
	}
	// Handling Lists
	if sdk.ProxyIdV6 != nil {
		tflog.Debug(ctx, "Packing list of objects for field ProxyIdV6")
		packed, d := packIpsecTunnelsAutoKeyProxyIdInnerListFromSdk(ctx, sdk.ProxyIdV6)
		diags.Append(d...)
		model.ProxyIdV6 = packed
	} else {
		model.ProxyIdV6 = basetypes.NewListNull(models.IpsecTunnelsAutoKeyProxyIdInner{}.AttrType())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.IpsecTunnelsAutoKey{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.IpsecTunnelsAutoKey", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for IpsecTunnelsAutoKey ---
func unpackIpsecTunnelsAutoKeyListToSdk(ctx context.Context, list types.List) ([]network_services.IpsecTunnelsAutoKey, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.IpsecTunnelsAutoKey")
	diags := diag.Diagnostics{}
	var data []models.IpsecTunnelsAutoKey
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.IpsecTunnelsAutoKey, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.IpsecTunnelsAutoKey{}.AttrTypes(), &item)
		unpacked, d := unpackIpsecTunnelsAutoKeyToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.IpsecTunnelsAutoKey", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for IpsecTunnelsAutoKey ---
func packIpsecTunnelsAutoKeyListFromSdk(ctx context.Context, sdks []network_services.IpsecTunnelsAutoKey) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.IpsecTunnelsAutoKey")
	diags := diag.Diagnostics{}
	var data []models.IpsecTunnelsAutoKey

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.IpsecTunnelsAutoKey
		obj, d := packIpsecTunnelsAutoKeyFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.IpsecTunnelsAutoKey{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.IpsecTunnelsAutoKey", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.IpsecTunnelsAutoKey{}.AttrType(), data)
}

// --- Unpacker for IpsecTunnelsAutoKeyIkeGatewayInner ---
func unpackIpsecTunnelsAutoKeyIkeGatewayInnerToSdk(ctx context.Context, obj types.Object) (*network_services.IpsecTunnelsAutoKeyIkeGatewayInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.IpsecTunnelsAutoKeyIkeGatewayInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.IpsecTunnelsAutoKeyIkeGatewayInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.IpsecTunnelsAutoKeyIkeGatewayInner
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.IpsecTunnelsAutoKeyIkeGatewayInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for IpsecTunnelsAutoKeyIkeGatewayInner ---
func packIpsecTunnelsAutoKeyIkeGatewayInnerFromSdk(ctx context.Context, sdk network_services.IpsecTunnelsAutoKeyIkeGatewayInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.IpsecTunnelsAutoKeyIkeGatewayInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.IpsecTunnelsAutoKeyIkeGatewayInner
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Name != nil {
		model.Name = basetypes.NewStringValue(*sdk.Name)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
	} else {
		model.Name = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.IpsecTunnelsAutoKeyIkeGatewayInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.IpsecTunnelsAutoKeyIkeGatewayInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for IpsecTunnelsAutoKeyIkeGatewayInner ---
func unpackIpsecTunnelsAutoKeyIkeGatewayInnerListToSdk(ctx context.Context, list types.List) ([]network_services.IpsecTunnelsAutoKeyIkeGatewayInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.IpsecTunnelsAutoKeyIkeGatewayInner")
	diags := diag.Diagnostics{}
	var data []models.IpsecTunnelsAutoKeyIkeGatewayInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.IpsecTunnelsAutoKeyIkeGatewayInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.IpsecTunnelsAutoKeyIkeGatewayInner{}.AttrTypes(), &item)
		unpacked, d := unpackIpsecTunnelsAutoKeyIkeGatewayInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.IpsecTunnelsAutoKeyIkeGatewayInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for IpsecTunnelsAutoKeyIkeGatewayInner ---
func packIpsecTunnelsAutoKeyIkeGatewayInnerListFromSdk(ctx context.Context, sdks []network_services.IpsecTunnelsAutoKeyIkeGatewayInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.IpsecTunnelsAutoKeyIkeGatewayInner")
	diags := diag.Diagnostics{}
	var data []models.IpsecTunnelsAutoKeyIkeGatewayInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.IpsecTunnelsAutoKeyIkeGatewayInner
		obj, d := packIpsecTunnelsAutoKeyIkeGatewayInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.IpsecTunnelsAutoKeyIkeGatewayInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.IpsecTunnelsAutoKeyIkeGatewayInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.IpsecTunnelsAutoKeyIkeGatewayInner{}.AttrType(), data)
}

// --- Unpacker for IpsecTunnelsAutoKeyProxyIdInner ---
func unpackIpsecTunnelsAutoKeyProxyIdInnerToSdk(ctx context.Context, obj types.Object) (*network_services.IpsecTunnelsAutoKeyProxyIdInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.IpsecTunnelsAutoKeyProxyIdInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.IpsecTunnelsAutoKeyProxyIdInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.IpsecTunnelsAutoKeyProxyIdInner
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Local.IsNull() && !model.Local.IsUnknown() {
		sdk.Local = model.Local.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Local", "value": *sdk.Local})
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	// Handling Objects
	if !model.Protocol.IsNull() && !model.Protocol.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Protocol")
		unpacked, d := unpackIpsecTunnelsAutoKeyProxyIdInnerProtocolToSdk(ctx, model.Protocol)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Protocol"})
		}
		if unpacked != nil {
			sdk.Protocol = unpacked
		}
	}

	// Handling Primitives
	if !model.Remote.IsNull() && !model.Remote.IsUnknown() {
		sdk.Remote = model.Remote.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Remote", "value": *sdk.Remote})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.IpsecTunnelsAutoKeyProxyIdInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for IpsecTunnelsAutoKeyProxyIdInner ---
func packIpsecTunnelsAutoKeyProxyIdInnerFromSdk(ctx context.Context, sdk network_services.IpsecTunnelsAutoKeyProxyIdInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.IpsecTunnelsAutoKeyProxyIdInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.IpsecTunnelsAutoKeyProxyIdInner
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Local != nil {
		model.Local = basetypes.NewStringValue(*sdk.Local)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Local", "value": *sdk.Local})
	} else {
		model.Local = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Protocol != nil {
		tflog.Debug(ctx, "Packing nested object for field Protocol")
		packed, d := packIpsecTunnelsAutoKeyProxyIdInnerProtocolFromSdk(ctx, *sdk.Protocol)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Protocol"})
		}
		model.Protocol = packed
	} else {
		model.Protocol = basetypes.NewObjectNull(models.IpsecTunnelsAutoKeyProxyIdInnerProtocol{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Remote != nil {
		model.Remote = basetypes.NewStringValue(*sdk.Remote)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Remote", "value": *sdk.Remote})
	} else {
		model.Remote = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.IpsecTunnelsAutoKeyProxyIdInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.IpsecTunnelsAutoKeyProxyIdInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for IpsecTunnelsAutoKeyProxyIdInner ---
func unpackIpsecTunnelsAutoKeyProxyIdInnerListToSdk(ctx context.Context, list types.List) ([]network_services.IpsecTunnelsAutoKeyProxyIdInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.IpsecTunnelsAutoKeyProxyIdInner")
	diags := diag.Diagnostics{}
	var data []models.IpsecTunnelsAutoKeyProxyIdInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.IpsecTunnelsAutoKeyProxyIdInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.IpsecTunnelsAutoKeyProxyIdInner{}.AttrTypes(), &item)
		unpacked, d := unpackIpsecTunnelsAutoKeyProxyIdInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.IpsecTunnelsAutoKeyProxyIdInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for IpsecTunnelsAutoKeyProxyIdInner ---
func packIpsecTunnelsAutoKeyProxyIdInnerListFromSdk(ctx context.Context, sdks []network_services.IpsecTunnelsAutoKeyProxyIdInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.IpsecTunnelsAutoKeyProxyIdInner")
	diags := diag.Diagnostics{}
	var data []models.IpsecTunnelsAutoKeyProxyIdInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.IpsecTunnelsAutoKeyProxyIdInner
		obj, d := packIpsecTunnelsAutoKeyProxyIdInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.IpsecTunnelsAutoKeyProxyIdInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.IpsecTunnelsAutoKeyProxyIdInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.IpsecTunnelsAutoKeyProxyIdInner{}.AttrType(), data)
}

// --- Unpacker for IpsecTunnelsAutoKeyProxyIdInnerProtocol ---
func unpackIpsecTunnelsAutoKeyProxyIdInnerProtocolToSdk(ctx context.Context, obj types.Object) (*network_services.IpsecTunnelsAutoKeyProxyIdInnerProtocol, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.IpsecTunnelsAutoKeyProxyIdInnerProtocol", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.IpsecTunnelsAutoKeyProxyIdInnerProtocol
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.IpsecTunnelsAutoKeyProxyIdInnerProtocol
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Number.IsNull() && !model.Number.IsUnknown() {
		val := int32(model.Number.ValueInt64())
		sdk.Number = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Number", "value": *sdk.Number})
	}

	// Handling Objects
	if !model.Tcp.IsNull() && !model.Tcp.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Tcp")
		unpacked, d := unpackIpsecTunnelsAutoKeyProxyIdInnerProtocolTcpToSdk(ctx, model.Tcp)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Tcp"})
		}
		if unpacked != nil {
			sdk.Tcp = unpacked
		}
	}

	// Handling Objects
	if !model.Udp.IsNull() && !model.Udp.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Udp")
		unpacked, d := unpackIpsecTunnelsAutoKeyProxyIdInnerProtocolUdpToSdk(ctx, model.Udp)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Udp"})
		}
		if unpacked != nil {
			sdk.Udp = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.IpsecTunnelsAutoKeyProxyIdInnerProtocol", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for IpsecTunnelsAutoKeyProxyIdInnerProtocol ---
func packIpsecTunnelsAutoKeyProxyIdInnerProtocolFromSdk(ctx context.Context, sdk network_services.IpsecTunnelsAutoKeyProxyIdInnerProtocol) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.IpsecTunnelsAutoKeyProxyIdInnerProtocol", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.IpsecTunnelsAutoKeyProxyIdInnerProtocol
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Number != nil {
		model.Number = basetypes.NewInt64Value(int64(*sdk.Number))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Number", "value": *sdk.Number})
	} else {
		model.Number = basetypes.NewInt64Null()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Tcp != nil {
		tflog.Debug(ctx, "Packing nested object for field Tcp")
		packed, d := packIpsecTunnelsAutoKeyProxyIdInnerProtocolTcpFromSdk(ctx, *sdk.Tcp)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Tcp"})
		}
		model.Tcp = packed
	} else {
		model.Tcp = basetypes.NewObjectNull(models.IpsecTunnelsAutoKeyProxyIdInnerProtocolTcp{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Udp != nil {
		tflog.Debug(ctx, "Packing nested object for field Udp")
		packed, d := packIpsecTunnelsAutoKeyProxyIdInnerProtocolUdpFromSdk(ctx, *sdk.Udp)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Udp"})
		}
		model.Udp = packed
	} else {
		model.Udp = basetypes.NewObjectNull(models.IpsecTunnelsAutoKeyProxyIdInnerProtocolUdp{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.IpsecTunnelsAutoKeyProxyIdInnerProtocol{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.IpsecTunnelsAutoKeyProxyIdInnerProtocol", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for IpsecTunnelsAutoKeyProxyIdInnerProtocol ---
func unpackIpsecTunnelsAutoKeyProxyIdInnerProtocolListToSdk(ctx context.Context, list types.List) ([]network_services.IpsecTunnelsAutoKeyProxyIdInnerProtocol, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.IpsecTunnelsAutoKeyProxyIdInnerProtocol")
	diags := diag.Diagnostics{}
	var data []models.IpsecTunnelsAutoKeyProxyIdInnerProtocol
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.IpsecTunnelsAutoKeyProxyIdInnerProtocol, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.IpsecTunnelsAutoKeyProxyIdInnerProtocol{}.AttrTypes(), &item)
		unpacked, d := unpackIpsecTunnelsAutoKeyProxyIdInnerProtocolToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.IpsecTunnelsAutoKeyProxyIdInnerProtocol", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for IpsecTunnelsAutoKeyProxyIdInnerProtocol ---
func packIpsecTunnelsAutoKeyProxyIdInnerProtocolListFromSdk(ctx context.Context, sdks []network_services.IpsecTunnelsAutoKeyProxyIdInnerProtocol) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.IpsecTunnelsAutoKeyProxyIdInnerProtocol")
	diags := diag.Diagnostics{}
	var data []models.IpsecTunnelsAutoKeyProxyIdInnerProtocol

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.IpsecTunnelsAutoKeyProxyIdInnerProtocol
		obj, d := packIpsecTunnelsAutoKeyProxyIdInnerProtocolFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.IpsecTunnelsAutoKeyProxyIdInnerProtocol{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.IpsecTunnelsAutoKeyProxyIdInnerProtocol", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.IpsecTunnelsAutoKeyProxyIdInnerProtocol{}.AttrType(), data)
}

// --- Unpacker for IpsecTunnelsAutoKeyProxyIdInnerProtocolTcp ---
func unpackIpsecTunnelsAutoKeyProxyIdInnerProtocolTcpToSdk(ctx context.Context, obj types.Object) (*network_services.IpsecTunnelsAutoKeyProxyIdInnerProtocolTcp, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.IpsecTunnelsAutoKeyProxyIdInnerProtocolTcp", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.IpsecTunnelsAutoKeyProxyIdInnerProtocolTcp
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.IpsecTunnelsAutoKeyProxyIdInnerProtocolTcp
	var d diag.Diagnostics
	// Handling Primitives
	if !model.LocalPort.IsNull() && !model.LocalPort.IsUnknown() {
		val := int32(model.LocalPort.ValueInt64())
		sdk.LocalPort = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "LocalPort", "value": *sdk.LocalPort})
	}

	// Handling Primitives
	if !model.RemotePort.IsNull() && !model.RemotePort.IsUnknown() {
		val := int32(model.RemotePort.ValueInt64())
		sdk.RemotePort = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "RemotePort", "value": *sdk.RemotePort})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.IpsecTunnelsAutoKeyProxyIdInnerProtocolTcp", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for IpsecTunnelsAutoKeyProxyIdInnerProtocolTcp ---
func packIpsecTunnelsAutoKeyProxyIdInnerProtocolTcpFromSdk(ctx context.Context, sdk network_services.IpsecTunnelsAutoKeyProxyIdInnerProtocolTcp) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.IpsecTunnelsAutoKeyProxyIdInnerProtocolTcp", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.IpsecTunnelsAutoKeyProxyIdInnerProtocolTcp
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.LocalPort != nil {
		model.LocalPort = basetypes.NewInt64Value(int64(*sdk.LocalPort))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "LocalPort", "value": *sdk.LocalPort})
	} else {
		model.LocalPort = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.RemotePort != nil {
		model.RemotePort = basetypes.NewInt64Value(int64(*sdk.RemotePort))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "RemotePort", "value": *sdk.RemotePort})
	} else {
		model.RemotePort = basetypes.NewInt64Null()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.IpsecTunnelsAutoKeyProxyIdInnerProtocolTcp{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.IpsecTunnelsAutoKeyProxyIdInnerProtocolTcp", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for IpsecTunnelsAutoKeyProxyIdInnerProtocolTcp ---
func unpackIpsecTunnelsAutoKeyProxyIdInnerProtocolTcpListToSdk(ctx context.Context, list types.List) ([]network_services.IpsecTunnelsAutoKeyProxyIdInnerProtocolTcp, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.IpsecTunnelsAutoKeyProxyIdInnerProtocolTcp")
	diags := diag.Diagnostics{}
	var data []models.IpsecTunnelsAutoKeyProxyIdInnerProtocolTcp
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.IpsecTunnelsAutoKeyProxyIdInnerProtocolTcp, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.IpsecTunnelsAutoKeyProxyIdInnerProtocolTcp{}.AttrTypes(), &item)
		unpacked, d := unpackIpsecTunnelsAutoKeyProxyIdInnerProtocolTcpToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.IpsecTunnelsAutoKeyProxyIdInnerProtocolTcp", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for IpsecTunnelsAutoKeyProxyIdInnerProtocolTcp ---
func packIpsecTunnelsAutoKeyProxyIdInnerProtocolTcpListFromSdk(ctx context.Context, sdks []network_services.IpsecTunnelsAutoKeyProxyIdInnerProtocolTcp) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.IpsecTunnelsAutoKeyProxyIdInnerProtocolTcp")
	diags := diag.Diagnostics{}
	var data []models.IpsecTunnelsAutoKeyProxyIdInnerProtocolTcp

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.IpsecTunnelsAutoKeyProxyIdInnerProtocolTcp
		obj, d := packIpsecTunnelsAutoKeyProxyIdInnerProtocolTcpFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.IpsecTunnelsAutoKeyProxyIdInnerProtocolTcp{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.IpsecTunnelsAutoKeyProxyIdInnerProtocolTcp", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.IpsecTunnelsAutoKeyProxyIdInnerProtocolTcp{}.AttrType(), data)
}

// --- Unpacker for IpsecTunnelsAutoKeyProxyIdInnerProtocolUdp ---
func unpackIpsecTunnelsAutoKeyProxyIdInnerProtocolUdpToSdk(ctx context.Context, obj types.Object) (*network_services.IpsecTunnelsAutoKeyProxyIdInnerProtocolUdp, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.IpsecTunnelsAutoKeyProxyIdInnerProtocolUdp", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.IpsecTunnelsAutoKeyProxyIdInnerProtocolUdp
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.IpsecTunnelsAutoKeyProxyIdInnerProtocolUdp
	var d diag.Diagnostics
	// Handling Primitives
	if !model.LocalPort.IsNull() && !model.LocalPort.IsUnknown() {
		val := int32(model.LocalPort.ValueInt64())
		sdk.LocalPort = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "LocalPort", "value": *sdk.LocalPort})
	}

	// Handling Primitives
	if !model.RemotePort.IsNull() && !model.RemotePort.IsUnknown() {
		val := int32(model.RemotePort.ValueInt64())
		sdk.RemotePort = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "RemotePort", "value": *sdk.RemotePort})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.IpsecTunnelsAutoKeyProxyIdInnerProtocolUdp", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for IpsecTunnelsAutoKeyProxyIdInnerProtocolUdp ---
func packIpsecTunnelsAutoKeyProxyIdInnerProtocolUdpFromSdk(ctx context.Context, sdk network_services.IpsecTunnelsAutoKeyProxyIdInnerProtocolUdp) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.IpsecTunnelsAutoKeyProxyIdInnerProtocolUdp", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.IpsecTunnelsAutoKeyProxyIdInnerProtocolUdp
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.LocalPort != nil {
		model.LocalPort = basetypes.NewInt64Value(int64(*sdk.LocalPort))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "LocalPort", "value": *sdk.LocalPort})
	} else {
		model.LocalPort = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.RemotePort != nil {
		model.RemotePort = basetypes.NewInt64Value(int64(*sdk.RemotePort))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "RemotePort", "value": *sdk.RemotePort})
	} else {
		model.RemotePort = basetypes.NewInt64Null()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.IpsecTunnelsAutoKeyProxyIdInnerProtocolUdp{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.IpsecTunnelsAutoKeyProxyIdInnerProtocolUdp", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for IpsecTunnelsAutoKeyProxyIdInnerProtocolUdp ---
func unpackIpsecTunnelsAutoKeyProxyIdInnerProtocolUdpListToSdk(ctx context.Context, list types.List) ([]network_services.IpsecTunnelsAutoKeyProxyIdInnerProtocolUdp, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.IpsecTunnelsAutoKeyProxyIdInnerProtocolUdp")
	diags := diag.Diagnostics{}
	var data []models.IpsecTunnelsAutoKeyProxyIdInnerProtocolUdp
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.IpsecTunnelsAutoKeyProxyIdInnerProtocolUdp, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.IpsecTunnelsAutoKeyProxyIdInnerProtocolUdp{}.AttrTypes(), &item)
		unpacked, d := unpackIpsecTunnelsAutoKeyProxyIdInnerProtocolUdpToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.IpsecTunnelsAutoKeyProxyIdInnerProtocolUdp", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for IpsecTunnelsAutoKeyProxyIdInnerProtocolUdp ---
func packIpsecTunnelsAutoKeyProxyIdInnerProtocolUdpListFromSdk(ctx context.Context, sdks []network_services.IpsecTunnelsAutoKeyProxyIdInnerProtocolUdp) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.IpsecTunnelsAutoKeyProxyIdInnerProtocolUdp")
	diags := diag.Diagnostics{}
	var data []models.IpsecTunnelsAutoKeyProxyIdInnerProtocolUdp

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.IpsecTunnelsAutoKeyProxyIdInnerProtocolUdp
		obj, d := packIpsecTunnelsAutoKeyProxyIdInnerProtocolUdpFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.IpsecTunnelsAutoKeyProxyIdInnerProtocolUdp{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.IpsecTunnelsAutoKeyProxyIdInnerProtocolUdp", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.IpsecTunnelsAutoKeyProxyIdInnerProtocolUdp{}.AttrType(), data)
}

// --- Unpacker for IpsecTunnelsTunnelMonitor ---
func unpackIpsecTunnelsTunnelMonitorToSdk(ctx context.Context, obj types.Object) (*network_services.IpsecTunnelsTunnelMonitor, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.IpsecTunnelsTunnelMonitor", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.IpsecTunnelsTunnelMonitor
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.IpsecTunnelsTunnelMonitor
	var d diag.Diagnostics
	// Handling Primitives
	if !model.DestinationIp.IsNull() && !model.DestinationIp.IsUnknown() {
		sdk.DestinationIp = model.DestinationIp.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "DestinationIp", "value": sdk.DestinationIp})
	}

	// Handling Primitives
	if !model.Enable.IsNull() && !model.Enable.IsUnknown() {
		sdk.Enable = model.Enable.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Enable", "value": *sdk.Enable})
	}

	// Handling Primitives
	if !model.ProxyId.IsNull() && !model.ProxyId.IsUnknown() {
		sdk.ProxyId = model.ProxyId.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "ProxyId", "value": *sdk.ProxyId})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.IpsecTunnelsTunnelMonitor", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for IpsecTunnelsTunnelMonitor ---
func packIpsecTunnelsTunnelMonitorFromSdk(ctx context.Context, sdk network_services.IpsecTunnelsTunnelMonitor) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.IpsecTunnelsTunnelMonitor", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.IpsecTunnelsTunnelMonitor
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.DestinationIp = basetypes.NewStringValue(sdk.DestinationIp)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "DestinationIp", "value": sdk.DestinationIp})
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
	if sdk.ProxyId != nil {
		model.ProxyId = basetypes.NewStringValue(*sdk.ProxyId)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "ProxyId", "value": *sdk.ProxyId})
	} else {
		model.ProxyId = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.IpsecTunnelsTunnelMonitor{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.IpsecTunnelsTunnelMonitor", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for IpsecTunnelsTunnelMonitor ---
func unpackIpsecTunnelsTunnelMonitorListToSdk(ctx context.Context, list types.List) ([]network_services.IpsecTunnelsTunnelMonitor, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.IpsecTunnelsTunnelMonitor")
	diags := diag.Diagnostics{}
	var data []models.IpsecTunnelsTunnelMonitor
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.IpsecTunnelsTunnelMonitor, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.IpsecTunnelsTunnelMonitor{}.AttrTypes(), &item)
		unpacked, d := unpackIpsecTunnelsTunnelMonitorToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.IpsecTunnelsTunnelMonitor", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for IpsecTunnelsTunnelMonitor ---
func packIpsecTunnelsTunnelMonitorListFromSdk(ctx context.Context, sdks []network_services.IpsecTunnelsTunnelMonitor) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.IpsecTunnelsTunnelMonitor")
	diags := diag.Diagnostics{}
	var data []models.IpsecTunnelsTunnelMonitor

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.IpsecTunnelsTunnelMonitor
		obj, d := packIpsecTunnelsTunnelMonitorFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.IpsecTunnelsTunnelMonitor{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.IpsecTunnelsTunnelMonitor", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.IpsecTunnelsTunnelMonitor{}.AttrType(), data)
}
