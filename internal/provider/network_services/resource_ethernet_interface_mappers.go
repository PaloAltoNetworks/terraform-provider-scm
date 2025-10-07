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

// ethernetInterfacesSensitiveValuePatcher is an in-memory struct to temporarily store plaintext
// and encrypted values for sensitive fields during the Create/Update/Read workflows.
type ethernetInterfacesSensitiveValuePatcher struct {
	layer3_pppoe_password_plaintext basetypes.StringValue
	layer3_pppoe_password_encrypted basetypes.StringValue
}

// populatePatcherFromState populates the patcher struct from the resource's state.
func (p *ethernetInterfacesSensitiveValuePatcher) populatePatcherFromState(ctx context.Context, state models.EthernetInterfaces) diag.Diagnostics {
	var diags diag.Diagnostics
	if state.EncryptedValues.IsNull() || state.EncryptedValues.IsUnknown() {
		return diags
	}

	var ev map[string]string
	diags.Append(state.EncryptedValues.ElementsAs(ctx, &ev, false)...)
	if diags.HasError() {
		return diags
	}
	if val, ok := ev["layer3_pppoe_password_plaintext"]; ok {
		p.layer3_pppoe_password_plaintext = basetypes.NewStringValue(val)
	}
	if val, ok := ev["layer3_pppoe_password_encrypted"]; ok {
		p.layer3_pppoe_password_encrypted = basetypes.NewStringValue(val)
	}

	return diags
}

// populateEncryptedValuesMap returns a map of the patcher's values for saving to state.
func (p *ethernetInterfacesSensitiveValuePatcher) populateEncryptedValuesMap() map[string]string {
	ev := make(map[string]string)
	if !p.layer3_pppoe_password_plaintext.IsNull() {
		ev["layer3_pppoe_password_plaintext"] = p.layer3_pppoe_password_plaintext.ValueString()
	}
	if !p.layer3_pppoe_password_encrypted.IsNull() {
		ev["layer3_pppoe_password_encrypted"] = p.layer3_pppoe_password_encrypted.ValueString()
	}
	return ev
}

// --- Unpacker for EthernetInterfaces ---
func unpackEthernetInterfacesToSdk(ctx context.Context, obj types.Object) (*network_services.EthernetInterfaces, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.EthernetInterfaces", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.EthernetInterfaces
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.EthernetInterfaces
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Comment.IsNull() && !model.Comment.IsUnknown() {
		sdk.Comment = model.Comment.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Comment", "value": *sdk.Comment})
	}

	// Handling Primitives
	if !model.DefaultValue.IsNull() && !model.DefaultValue.IsUnknown() {
		sdk.DefaultValue = model.DefaultValue.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DefaultValue", "value": *sdk.DefaultValue})
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

	// Handling Objects
	if !model.Layer2.IsNull() && !model.Layer2.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Layer2")
		unpacked, d := unpackEthernetInterfacesLayer2ToSdk(ctx, model.Layer2)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Layer2"})
		}
		if unpacked != nil {
			sdk.Layer2 = unpacked
		}
	}

	// Handling Objects
	if !model.Layer3.IsNull() && !model.Layer3.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Layer3")
		unpacked, d := unpackEthernetInterfacesLayer3ToSdk(ctx, model.Layer3)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Layer3"})
		}
		if unpacked != nil {
			sdk.Layer3 = unpacked
		}
	}

	// Handling Primitives
	if !model.LinkDuplex.IsNull() && !model.LinkDuplex.IsUnknown() {
		sdk.LinkDuplex = model.LinkDuplex.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "LinkDuplex", "value": *sdk.LinkDuplex})
	}

	// Handling Primitives
	if !model.LinkSpeed.IsNull() && !model.LinkSpeed.IsUnknown() {
		sdk.LinkSpeed = model.LinkSpeed.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "LinkSpeed", "value": *sdk.LinkSpeed})
	}

	// Handling Primitives
	if !model.LinkState.IsNull() && !model.LinkState.IsUnknown() {
		sdk.LinkState = model.LinkState.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "LinkState", "value": *sdk.LinkState})
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	// Handling Objects
	if !model.Poe.IsNull() && !model.Poe.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Poe")
		unpacked, d := unpackPoeToSdk(ctx, model.Poe)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Poe"})
		}
		if unpacked != nil {
			sdk.Poe = unpacked
		}
	}

	// Handling Primitives
	if !model.Snippet.IsNull() && !model.Snippet.IsUnknown() {
		sdk.Snippet = model.Snippet.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	}

	// Handling Typeless Objects
	if !model.Tap.IsNull() && !model.Tap.IsUnknown() {
		tflog.Debug(ctx, "Unpacking typeless object for field Tap")
		sdk.Tap = make(map[string]interface{})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.EthernetInterfaces", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for EthernetInterfaces ---
func packEthernetInterfacesFromSdk(ctx context.Context, sdk network_services.EthernetInterfaces) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.EthernetInterfaces", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.EthernetInterfaces
	var d diag.Diagnostics
	// MEGA FIX FOR MAP TYPE MISMATCH (NOT ALL MODELS MAY HAVE EncryptedValues)
	model.EncryptedValues = basetypes.NewMapNull(basetypes.StringType{})
	// Handling Primitives
	// Standard primitive packing
	if sdk.Comment != nil {
		model.Comment = basetypes.NewStringValue(*sdk.Comment)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Comment", "value": *sdk.Comment})
	} else {
		model.Comment = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.DefaultValue != nil {
		model.DefaultValue = basetypes.NewStringValue(*sdk.DefaultValue)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DefaultValue", "value": *sdk.DefaultValue})
	} else {
		model.DefaultValue = basetypes.NewStringNull()
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
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Layer2 != nil {
		tflog.Debug(ctx, "Packing nested object for field Layer2")
		packed, d := packEthernetInterfacesLayer2FromSdk(ctx, *sdk.Layer2)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Layer2"})
		}
		model.Layer2 = packed
	} else {
		model.Layer2 = basetypes.NewObjectNull(models.EthernetInterfacesLayer2{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Layer3 != nil {
		tflog.Debug(ctx, "Packing nested object for field Layer3")
		packed, d := packEthernetInterfacesLayer3FromSdk(ctx, *sdk.Layer3)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Layer3"})
		}
		model.Layer3 = packed
	} else {
		model.Layer3 = basetypes.NewObjectNull(models.EthernetInterfacesLayer3{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.LinkDuplex != nil {
		model.LinkDuplex = basetypes.NewStringValue(*sdk.LinkDuplex)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "LinkDuplex", "value": *sdk.LinkDuplex})
	} else {
		model.LinkDuplex = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.LinkSpeed != nil {
		model.LinkSpeed = basetypes.NewStringValue(*sdk.LinkSpeed)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "LinkSpeed", "value": *sdk.LinkSpeed})
	} else {
		model.LinkSpeed = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.LinkState != nil {
		model.LinkState = basetypes.NewStringValue(*sdk.LinkState)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "LinkState", "value": *sdk.LinkState})
	} else {
		model.LinkState = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Poe != nil {
		tflog.Debug(ctx, "Packing nested object for field Poe")
		packed, d := packPoeFromSdk(ctx, *sdk.Poe)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Poe"})
		}
		model.Poe = packed
	} else {
		model.Poe = basetypes.NewObjectNull(models.Poe{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Snippet != nil {
		model.Snippet = basetypes.NewStringValue(*sdk.Snippet)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	} else {
		model.Snippet = basetypes.NewStringNull()
	}
	// Handling Objects
	// This is a marker object (e.g. CHAP: {}). We just need to create an empty, non-null object.
	if sdk.Tap != nil && !reflect.ValueOf(sdk.Tap).IsNil() {
		tflog.Debug(ctx, "Packing typeless object for field Tap")
		var d diag.Diagnostics
		// Create an empty object with no attributes, which signifies its presence.
		model.Tap, d = basetypes.NewObjectValue(map[string]attr.Type{}, map[string]attr.Value{})
		diags.Append(d...)
	} else {
		// Since this field is part of a oneOf, being nil means it's not selected.
		// We make the object null with an empty attribute map.
		model.Tap = basetypes.NewObjectNull(map[string]attr.Type{})
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.EthernetInterfaces{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.EthernetInterfaces", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for EthernetInterfaces ---
func unpackEthernetInterfacesListToSdk(ctx context.Context, list types.List) ([]network_services.EthernetInterfaces, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.EthernetInterfaces")
	diags := diag.Diagnostics{}
	var data []models.EthernetInterfaces
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.EthernetInterfaces, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.EthernetInterfaces{}.AttrTypes(), &item)
		unpacked, d := unpackEthernetInterfacesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.EthernetInterfaces", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for EthernetInterfaces ---
func packEthernetInterfacesListFromSdk(ctx context.Context, sdks []network_services.EthernetInterfaces) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.EthernetInterfaces")
	diags := diag.Diagnostics{}
	var data []models.EthernetInterfaces

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.EthernetInterfaces
		obj, d := packEthernetInterfacesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.EthernetInterfaces{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.EthernetInterfaces", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.EthernetInterfaces{}.AttrType(), data)
}

// --- Unpacker for EthernetInterfacesLayer2 ---
func unpackEthernetInterfacesLayer2ToSdk(ctx context.Context, obj types.Object) (*network_services.EthernetInterfacesLayer2, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.EthernetInterfacesLayer2", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.EthernetInterfacesLayer2
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.EthernetInterfacesLayer2
	var d diag.Diagnostics
	// Handling Primitives
	if !model.VlanTag.IsNull() && !model.VlanTag.IsUnknown() {
		val := int32(model.VlanTag.ValueInt64())
		sdk.VlanTag = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "VlanTag", "value": *sdk.VlanTag})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.EthernetInterfacesLayer2", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for EthernetInterfacesLayer2 ---
func packEthernetInterfacesLayer2FromSdk(ctx context.Context, sdk network_services.EthernetInterfacesLayer2) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.EthernetInterfacesLayer2", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.EthernetInterfacesLayer2
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.VlanTag != nil {
		model.VlanTag = basetypes.NewInt64Value(int64(*sdk.VlanTag))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "VlanTag", "value": *sdk.VlanTag})
	} else {
		model.VlanTag = basetypes.NewInt64Null()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.EthernetInterfacesLayer2{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.EthernetInterfacesLayer2", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for EthernetInterfacesLayer2 ---
func unpackEthernetInterfacesLayer2ListToSdk(ctx context.Context, list types.List) ([]network_services.EthernetInterfacesLayer2, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.EthernetInterfacesLayer2")
	diags := diag.Diagnostics{}
	var data []models.EthernetInterfacesLayer2
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.EthernetInterfacesLayer2, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.EthernetInterfacesLayer2{}.AttrTypes(), &item)
		unpacked, d := unpackEthernetInterfacesLayer2ToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.EthernetInterfacesLayer2", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for EthernetInterfacesLayer2 ---
func packEthernetInterfacesLayer2ListFromSdk(ctx context.Context, sdks []network_services.EthernetInterfacesLayer2) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.EthernetInterfacesLayer2")
	diags := diag.Diagnostics{}
	var data []models.EthernetInterfacesLayer2

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.EthernetInterfacesLayer2
		obj, d := packEthernetInterfacesLayer2FromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.EthernetInterfacesLayer2{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.EthernetInterfacesLayer2", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.EthernetInterfacesLayer2{}.AttrType(), data)
}

// --- Unpacker for EthernetInterfacesLayer3 ---
func unpackEthernetInterfacesLayer3ToSdk(ctx context.Context, obj types.Object) (*network_services.EthernetInterfacesLayer3, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.EthernetInterfacesLayer3", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.EthernetInterfacesLayer3
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.EthernetInterfacesLayer3
	var d diag.Diagnostics
	// Handling Lists
	if !model.Arp.IsNull() && !model.Arp.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field Arp")
		unpacked, d := unpackArpInnerListToSdk(ctx, model.Arp)
		diags.Append(d...)
		sdk.Arp = unpacked
	}

	// Handling Objects
	if !model.DdnsConfig.IsNull() && !model.DdnsConfig.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field DdnsConfig")
		unpacked, d := unpackEthernetInterfacesLayer3DdnsConfigToSdk(ctx, model.DdnsConfig)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "DdnsConfig"})
		}
		if unpacked != nil {
			sdk.DdnsConfig = unpacked
		}
	}

	// Handling Objects
	if !model.DhcpClient.IsNull() && !model.DhcpClient.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field DhcpClient")
		unpacked, d := unpackDhcpClientToSdk(ctx, model.DhcpClient)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "DhcpClient"})
		}
		if unpacked != nil {
			sdk.DhcpClient = unpacked
		}
	}

	// Handling Primitives
	if !model.InterfaceManagementProfile.IsNull() && !model.InterfaceManagementProfile.IsUnknown() {
		sdk.InterfaceManagementProfile = model.InterfaceManagementProfile.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "InterfaceManagementProfile", "value": *sdk.InterfaceManagementProfile})
	}

	// Handling Lists
	if !model.Ip.IsNull() && !model.Ip.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Ip")
		diags.Append(model.Ip.ElementsAs(ctx, &sdk.Ip, false)...)
	}

	// Handling Primitives
	if !model.Mtu.IsNull() && !model.Mtu.IsUnknown() {
		val := int32(model.Mtu.ValueInt64())
		sdk.Mtu = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Mtu", "value": *sdk.Mtu})
	}

	// Handling Objects
	if !model.Pppoe.IsNull() && !model.Pppoe.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Pppoe")
		unpacked, d := unpackEthernetInterfacesLayer3PppoeToSdk(ctx, model.Pppoe)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Pppoe"})
		}
		if unpacked != nil {
			sdk.Pppoe = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.EthernetInterfacesLayer3", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for EthernetInterfacesLayer3 ---
func packEthernetInterfacesLayer3FromSdk(ctx context.Context, sdk network_services.EthernetInterfacesLayer3) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.EthernetInterfacesLayer3", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.EthernetInterfacesLayer3
	var d diag.Diagnostics
	// Handling Lists
	if sdk.Arp != nil {
		tflog.Debug(ctx, "Packing list of objects for field Arp")
		packed, d := packArpInnerListFromSdk(ctx, sdk.Arp)
		diags.Append(d...)
		model.Arp = packed
	} else {
		model.Arp = basetypes.NewListNull(models.ArpInner{}.AttrType())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.DdnsConfig != nil {
		tflog.Debug(ctx, "Packing nested object for field DdnsConfig")
		packed, d := packEthernetInterfacesLayer3DdnsConfigFromSdk(ctx, *sdk.DdnsConfig)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "DdnsConfig"})
		}
		model.DdnsConfig = packed
	} else {
		model.DdnsConfig = basetypes.NewObjectNull(models.EthernetInterfacesLayer3DdnsConfig{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.DhcpClient != nil {
		tflog.Debug(ctx, "Packing nested object for field DhcpClient")
		packed, d := packDhcpClientFromSdk(ctx, *sdk.DhcpClient)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "DhcpClient"})
		}
		model.DhcpClient = packed
	} else {
		model.DhcpClient = basetypes.NewObjectNull(models.DhcpClient{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.InterfaceManagementProfile != nil {
		model.InterfaceManagementProfile = basetypes.NewStringValue(*sdk.InterfaceManagementProfile)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "InterfaceManagementProfile", "value": *sdk.InterfaceManagementProfile})
	} else {
		model.InterfaceManagementProfile = basetypes.NewStringNull()
	}
	// Handling Lists
	if sdk.Ip != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Ip")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Ip, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Ip)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Ip = basetypes.NewListNull(elemType)
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Mtu != nil {
		model.Mtu = basetypes.NewInt64Value(int64(*sdk.Mtu))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Mtu", "value": *sdk.Mtu})
	} else {
		model.Mtu = basetypes.NewInt64Null()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Pppoe != nil {
		tflog.Debug(ctx, "Packing nested object for field Pppoe")
		packed, d := packEthernetInterfacesLayer3PppoeFromSdk(ctx, *sdk.Pppoe)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Pppoe"})
		}
		model.Pppoe = packed
	} else {
		model.Pppoe = basetypes.NewObjectNull(models.EthernetInterfacesLayer3Pppoe{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.EthernetInterfacesLayer3{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.EthernetInterfacesLayer3", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for EthernetInterfacesLayer3 ---
func unpackEthernetInterfacesLayer3ListToSdk(ctx context.Context, list types.List) ([]network_services.EthernetInterfacesLayer3, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.EthernetInterfacesLayer3")
	diags := diag.Diagnostics{}
	var data []models.EthernetInterfacesLayer3
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.EthernetInterfacesLayer3, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.EthernetInterfacesLayer3{}.AttrTypes(), &item)
		unpacked, d := unpackEthernetInterfacesLayer3ToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.EthernetInterfacesLayer3", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for EthernetInterfacesLayer3 ---
func packEthernetInterfacesLayer3ListFromSdk(ctx context.Context, sdks []network_services.EthernetInterfacesLayer3) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.EthernetInterfacesLayer3")
	diags := diag.Diagnostics{}
	var data []models.EthernetInterfacesLayer3

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.EthernetInterfacesLayer3
		obj, d := packEthernetInterfacesLayer3FromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.EthernetInterfacesLayer3{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.EthernetInterfacesLayer3", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.EthernetInterfacesLayer3{}.AttrType(), data)
}

// --- Unpacker for ArpInner ---
func unpackArpInnerToSdk(ctx context.Context, obj types.Object) (*network_services.ArpInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ArpInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ArpInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.ArpInner
	var d diag.Diagnostics
	// Handling Primitives
	if !model.HwAddress.IsNull() && !model.HwAddress.IsUnknown() {
		sdk.HwAddress = model.HwAddress.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "HwAddress", "value": *sdk.HwAddress})
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ArpInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ArpInner ---
func packArpInnerFromSdk(ctx context.Context, sdk network_services.ArpInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ArpInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ArpInner
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.HwAddress != nil {
		model.HwAddress = basetypes.NewStringValue(*sdk.HwAddress)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "HwAddress", "value": *sdk.HwAddress})
	} else {
		model.HwAddress = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Name != nil {
		model.Name = basetypes.NewStringValue(*sdk.Name)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
	} else {
		model.Name = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ArpInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ArpInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ArpInner ---
func unpackArpInnerListToSdk(ctx context.Context, list types.List) ([]network_services.ArpInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ArpInner")
	diags := diag.Diagnostics{}
	var data []models.ArpInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.ArpInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ArpInner{}.AttrTypes(), &item)
		unpacked, d := unpackArpInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ArpInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ArpInner ---
func packArpInnerListFromSdk(ctx context.Context, sdks []network_services.ArpInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ArpInner")
	diags := diag.Diagnostics{}
	var data []models.ArpInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ArpInner
		obj, d := packArpInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ArpInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ArpInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ArpInner{}.AttrType(), data)
}

// --- Unpacker for EthernetInterfacesLayer3DdnsConfig ---
func unpackEthernetInterfacesLayer3DdnsConfigToSdk(ctx context.Context, obj types.Object) (*network_services.EthernetInterfacesLayer3DdnsConfig, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.EthernetInterfacesLayer3DdnsConfig", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.EthernetInterfacesLayer3DdnsConfig
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.EthernetInterfacesLayer3DdnsConfig
	var d diag.Diagnostics
	// Handling Primitives
	if !model.DdnsCertProfile.IsNull() && !model.DdnsCertProfile.IsUnknown() {
		sdk.DdnsCertProfile = model.DdnsCertProfile.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "DdnsCertProfile", "value": sdk.DdnsCertProfile})
	}

	// Handling Primitives
	if !model.DdnsEnabled.IsNull() && !model.DdnsEnabled.IsUnknown() {
		sdk.DdnsEnabled = model.DdnsEnabled.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DdnsEnabled", "value": *sdk.DdnsEnabled})
	}

	// Handling Primitives
	if !model.DdnsHostname.IsNull() && !model.DdnsHostname.IsUnknown() {
		sdk.DdnsHostname = model.DdnsHostname.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "DdnsHostname", "value": sdk.DdnsHostname})
	}

	// Handling Primitives
	if !model.DdnsIp.IsNull() && !model.DdnsIp.IsUnknown() {
		sdk.DdnsIp = model.DdnsIp.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DdnsIp", "value": *sdk.DdnsIp})
	}

	// Handling Primitives
	if !model.DdnsUpdateInterval.IsNull() && !model.DdnsUpdateInterval.IsUnknown() {
		val := int32(model.DdnsUpdateInterval.ValueInt64())
		sdk.DdnsUpdateInterval = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DdnsUpdateInterval", "value": *sdk.DdnsUpdateInterval})
	}

	// Handling Primitives
	if !model.DdnsVendor.IsNull() && !model.DdnsVendor.IsUnknown() {
		sdk.DdnsVendor = model.DdnsVendor.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "DdnsVendor", "value": sdk.DdnsVendor})
	}

	// Handling Primitives
	if !model.DdnsVendorConfig.IsNull() && !model.DdnsVendorConfig.IsUnknown() {
		sdk.DdnsVendorConfig = model.DdnsVendorConfig.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "DdnsVendorConfig", "value": sdk.DdnsVendorConfig})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.EthernetInterfacesLayer3DdnsConfig", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for EthernetInterfacesLayer3DdnsConfig ---
func packEthernetInterfacesLayer3DdnsConfigFromSdk(ctx context.Context, sdk network_services.EthernetInterfacesLayer3DdnsConfig) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.EthernetInterfacesLayer3DdnsConfig", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.EthernetInterfacesLayer3DdnsConfig
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.DdnsCertProfile = basetypes.NewStringValue(sdk.DdnsCertProfile)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "DdnsCertProfile", "value": sdk.DdnsCertProfile})
	// Handling Primitives
	// Standard primitive packing
	if sdk.DdnsEnabled != nil {
		model.DdnsEnabled = basetypes.NewBoolValue(*sdk.DdnsEnabled)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DdnsEnabled", "value": *sdk.DdnsEnabled})
	} else {
		model.DdnsEnabled = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	model.DdnsHostname = basetypes.NewStringValue(sdk.DdnsHostname)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "DdnsHostname", "value": sdk.DdnsHostname})
	// Handling Primitives
	// Standard primitive packing
	if sdk.DdnsIp != nil {
		model.DdnsIp = basetypes.NewStringValue(*sdk.DdnsIp)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DdnsIp", "value": *sdk.DdnsIp})
	} else {
		model.DdnsIp = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.DdnsUpdateInterval != nil {
		model.DdnsUpdateInterval = basetypes.NewInt64Value(int64(*sdk.DdnsUpdateInterval))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DdnsUpdateInterval", "value": *sdk.DdnsUpdateInterval})
	} else {
		model.DdnsUpdateInterval = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	model.DdnsVendor = basetypes.NewStringValue(sdk.DdnsVendor)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "DdnsVendor", "value": sdk.DdnsVendor})
	// Handling Primitives
	// Standard primitive packing
	model.DdnsVendorConfig = basetypes.NewStringValue(sdk.DdnsVendorConfig)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "DdnsVendorConfig", "value": sdk.DdnsVendorConfig})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.EthernetInterfacesLayer3DdnsConfig{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.EthernetInterfacesLayer3DdnsConfig", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for EthernetInterfacesLayer3DdnsConfig ---
func unpackEthernetInterfacesLayer3DdnsConfigListToSdk(ctx context.Context, list types.List) ([]network_services.EthernetInterfacesLayer3DdnsConfig, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.EthernetInterfacesLayer3DdnsConfig")
	diags := diag.Diagnostics{}
	var data []models.EthernetInterfacesLayer3DdnsConfig
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.EthernetInterfacesLayer3DdnsConfig, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.EthernetInterfacesLayer3DdnsConfig{}.AttrTypes(), &item)
		unpacked, d := unpackEthernetInterfacesLayer3DdnsConfigToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.EthernetInterfacesLayer3DdnsConfig", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for EthernetInterfacesLayer3DdnsConfig ---
func packEthernetInterfacesLayer3DdnsConfigListFromSdk(ctx context.Context, sdks []network_services.EthernetInterfacesLayer3DdnsConfig) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.EthernetInterfacesLayer3DdnsConfig")
	diags := diag.Diagnostics{}
	var data []models.EthernetInterfacesLayer3DdnsConfig

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.EthernetInterfacesLayer3DdnsConfig
		obj, d := packEthernetInterfacesLayer3DdnsConfigFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.EthernetInterfacesLayer3DdnsConfig{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.EthernetInterfacesLayer3DdnsConfig", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.EthernetInterfacesLayer3DdnsConfig{}.AttrType(), data)
}

// --- Unpacker for DhcpClient ---
func unpackDhcpClientToSdk(ctx context.Context, obj types.Object) (*network_services.DhcpClient, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.DhcpClient", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.DhcpClient
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.DhcpClient
	var d diag.Diagnostics
	// Handling Objects
	if !model.DhcpClient.IsNull() && !model.DhcpClient.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field DhcpClient")
		unpacked, d := unpackAggregateEthernetInterfacesLayer3DhcpClientToSdk(ctx, model.DhcpClient)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "DhcpClient"})
		}
		if unpacked != nil {
			sdk.DhcpClient = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.DhcpClient", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for DhcpClient ---
func packDhcpClientFromSdk(ctx context.Context, sdk network_services.DhcpClient) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.DhcpClient", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.DhcpClient
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.DhcpClient != nil {
		tflog.Debug(ctx, "Packing nested object for field DhcpClient")
		packed, d := packAggregateEthernetInterfacesLayer3DhcpClientFromSdk(ctx, *sdk.DhcpClient)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "DhcpClient"})
		}
		model.DhcpClient = packed
	} else {
		model.DhcpClient = basetypes.NewObjectNull(models.AggregateEthernetInterfacesLayer3DhcpClient{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.DhcpClient{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.DhcpClient", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for DhcpClient ---
func unpackDhcpClientListToSdk(ctx context.Context, list types.List) ([]network_services.DhcpClient, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.DhcpClient")
	diags := diag.Diagnostics{}
	var data []models.DhcpClient
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.DhcpClient, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.DhcpClient{}.AttrTypes(), &item)
		unpacked, d := unpackDhcpClientToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.DhcpClient", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for DhcpClient ---
func packDhcpClientListFromSdk(ctx context.Context, sdks []network_services.DhcpClient) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.DhcpClient")
	diags := diag.Diagnostics{}
	var data []models.DhcpClient

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.DhcpClient
		obj, d := packDhcpClientFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.DhcpClient{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.DhcpClient", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.DhcpClient{}.AttrType(), data)
}

// --- Unpacker for AggregateEthernetInterfacesLayer3DhcpClient ---
func unpackAggregateEthernetInterfacesLayer3DhcpClientToSdk(ctx context.Context, obj types.Object) (*network_services.AggregateEthernetInterfacesLayer3DhcpClient, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.AggregateEthernetInterfacesLayer3DhcpClient", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.AggregateEthernetInterfacesLayer3DhcpClient
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.AggregateEthernetInterfacesLayer3DhcpClient
	var d diag.Diagnostics
	// Handling Primitives
	if !model.CreateDefaultRoute.IsNull() && !model.CreateDefaultRoute.IsUnknown() {
		sdk.CreateDefaultRoute = model.CreateDefaultRoute.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "CreateDefaultRoute", "value": *sdk.CreateDefaultRoute})
	}

	// Handling Primitives
	if !model.DefaultRouteMetric.IsNull() && !model.DefaultRouteMetric.IsUnknown() {
		val := int32(model.DefaultRouteMetric.ValueInt64())
		sdk.DefaultRouteMetric = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DefaultRouteMetric", "value": *sdk.DefaultRouteMetric})
	}

	// Handling Primitives
	if !model.Enable.IsNull() && !model.Enable.IsUnknown() {
		sdk.Enable = model.Enable.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Enable", "value": *sdk.Enable})
	}

	// Handling Objects
	if !model.SendHostname.IsNull() && !model.SendHostname.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field SendHostname")
		unpacked, d := unpackAggregateEthernetInterfacesLayer3DhcpClientSendHostnameToSdk(ctx, model.SendHostname)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "SendHostname"})
		}
		if unpacked != nil {
			sdk.SendHostname = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.AggregateEthernetInterfacesLayer3DhcpClient", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for AggregateEthernetInterfacesLayer3DhcpClient ---
func packAggregateEthernetInterfacesLayer3DhcpClientFromSdk(ctx context.Context, sdk network_services.AggregateEthernetInterfacesLayer3DhcpClient) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.AggregateEthernetInterfacesLayer3DhcpClient", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.AggregateEthernetInterfacesLayer3DhcpClient
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.CreateDefaultRoute != nil {
		model.CreateDefaultRoute = basetypes.NewBoolValue(*sdk.CreateDefaultRoute)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "CreateDefaultRoute", "value": *sdk.CreateDefaultRoute})
	} else {
		model.CreateDefaultRoute = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.DefaultRouteMetric != nil {
		model.DefaultRouteMetric = basetypes.NewInt64Value(int64(*sdk.DefaultRouteMetric))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DefaultRouteMetric", "value": *sdk.DefaultRouteMetric})
	} else {
		model.DefaultRouteMetric = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Enable != nil {
		model.Enable = basetypes.NewBoolValue(*sdk.Enable)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Enable", "value": *sdk.Enable})
	} else {
		model.Enable = basetypes.NewBoolNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.SendHostname != nil {
		tflog.Debug(ctx, "Packing nested object for field SendHostname")
		packed, d := packAggregateEthernetInterfacesLayer3DhcpClientSendHostnameFromSdk(ctx, *sdk.SendHostname)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "SendHostname"})
		}
		model.SendHostname = packed
	} else {
		model.SendHostname = basetypes.NewObjectNull(models.AggregateEthernetInterfacesLayer3DhcpClientSendHostname{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.AggregateEthernetInterfacesLayer3DhcpClient{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.AggregateEthernetInterfacesLayer3DhcpClient", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for AggregateEthernetInterfacesLayer3DhcpClient ---
func unpackAggregateEthernetInterfacesLayer3DhcpClientListToSdk(ctx context.Context, list types.List) ([]network_services.AggregateEthernetInterfacesLayer3DhcpClient, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.AggregateEthernetInterfacesLayer3DhcpClient")
	diags := diag.Diagnostics{}
	var data []models.AggregateEthernetInterfacesLayer3DhcpClient
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.AggregateEthernetInterfacesLayer3DhcpClient, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.AggregateEthernetInterfacesLayer3DhcpClient{}.AttrTypes(), &item)
		unpacked, d := unpackAggregateEthernetInterfacesLayer3DhcpClientToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.AggregateEthernetInterfacesLayer3DhcpClient", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for AggregateEthernetInterfacesLayer3DhcpClient ---
func packAggregateEthernetInterfacesLayer3DhcpClientListFromSdk(ctx context.Context, sdks []network_services.AggregateEthernetInterfacesLayer3DhcpClient) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.AggregateEthernetInterfacesLayer3DhcpClient")
	diags := diag.Diagnostics{}
	var data []models.AggregateEthernetInterfacesLayer3DhcpClient

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.AggregateEthernetInterfacesLayer3DhcpClient
		obj, d := packAggregateEthernetInterfacesLayer3DhcpClientFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.AggregateEthernetInterfacesLayer3DhcpClient{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.AggregateEthernetInterfacesLayer3DhcpClient", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.AggregateEthernetInterfacesLayer3DhcpClient{}.AttrType(), data)
}

// --- Unpacker for AggregateEthernetInterfacesLayer3DhcpClientSendHostname ---
func unpackAggregateEthernetInterfacesLayer3DhcpClientSendHostnameToSdk(ctx context.Context, obj types.Object) (*network_services.AggregateEthernetInterfacesLayer3DhcpClientSendHostname, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.AggregateEthernetInterfacesLayer3DhcpClientSendHostname", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.AggregateEthernetInterfacesLayer3DhcpClientSendHostname
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.AggregateEthernetInterfacesLayer3DhcpClientSendHostname
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Enable.IsNull() && !model.Enable.IsUnknown() {
		sdk.Enable = model.Enable.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Enable", "value": *sdk.Enable})
	}

	// Handling Primitives
	if !model.Hostname.IsNull() && !model.Hostname.IsUnknown() {
		sdk.Hostname = model.Hostname.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Hostname", "value": *sdk.Hostname})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.AggregateEthernetInterfacesLayer3DhcpClientSendHostname", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for AggregateEthernetInterfacesLayer3DhcpClientSendHostname ---
func packAggregateEthernetInterfacesLayer3DhcpClientSendHostnameFromSdk(ctx context.Context, sdk network_services.AggregateEthernetInterfacesLayer3DhcpClientSendHostname) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.AggregateEthernetInterfacesLayer3DhcpClientSendHostname", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.AggregateEthernetInterfacesLayer3DhcpClientSendHostname
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
	if sdk.Hostname != nil {
		model.Hostname = basetypes.NewStringValue(*sdk.Hostname)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Hostname", "value": *sdk.Hostname})
	} else {
		model.Hostname = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.AggregateEthernetInterfacesLayer3DhcpClientSendHostname{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.AggregateEthernetInterfacesLayer3DhcpClientSendHostname", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for AggregateEthernetInterfacesLayer3DhcpClientSendHostname ---
func unpackAggregateEthernetInterfacesLayer3DhcpClientSendHostnameListToSdk(ctx context.Context, list types.List) ([]network_services.AggregateEthernetInterfacesLayer3DhcpClientSendHostname, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.AggregateEthernetInterfacesLayer3DhcpClientSendHostname")
	diags := diag.Diagnostics{}
	var data []models.AggregateEthernetInterfacesLayer3DhcpClientSendHostname
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.AggregateEthernetInterfacesLayer3DhcpClientSendHostname, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.AggregateEthernetInterfacesLayer3DhcpClientSendHostname{}.AttrTypes(), &item)
		unpacked, d := unpackAggregateEthernetInterfacesLayer3DhcpClientSendHostnameToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.AggregateEthernetInterfacesLayer3DhcpClientSendHostname", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for AggregateEthernetInterfacesLayer3DhcpClientSendHostname ---
func packAggregateEthernetInterfacesLayer3DhcpClientSendHostnameListFromSdk(ctx context.Context, sdks []network_services.AggregateEthernetInterfacesLayer3DhcpClientSendHostname) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.AggregateEthernetInterfacesLayer3DhcpClientSendHostname")
	diags := diag.Diagnostics{}
	var data []models.AggregateEthernetInterfacesLayer3DhcpClientSendHostname

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.AggregateEthernetInterfacesLayer3DhcpClientSendHostname
		obj, d := packAggregateEthernetInterfacesLayer3DhcpClientSendHostnameFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.AggregateEthernetInterfacesLayer3DhcpClientSendHostname{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.AggregateEthernetInterfacesLayer3DhcpClientSendHostname", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.AggregateEthernetInterfacesLayer3DhcpClientSendHostname{}.AttrType(), data)
}

// --- Unpacker for EthernetInterfacesLayer3Pppoe ---
func unpackEthernetInterfacesLayer3PppoeToSdk(ctx context.Context, obj types.Object) (*network_services.EthernetInterfacesLayer3Pppoe, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.EthernetInterfacesLayer3Pppoe", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.EthernetInterfacesLayer3Pppoe
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.EthernetInterfacesLayer3Pppoe
	var d diag.Diagnostics
	// Handling Primitives
	if !model.AccessConcentrator.IsNull() && !model.AccessConcentrator.IsUnknown() {
		sdk.AccessConcentrator = model.AccessConcentrator.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "AccessConcentrator", "value": *sdk.AccessConcentrator})
	}

	// Handling Primitives
	if !model.Authentication.IsNull() && !model.Authentication.IsUnknown() {
		sdk.Authentication = model.Authentication.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Authentication", "value": *sdk.Authentication})
	}

	// Handling Primitives
	if !model.DefaultRouteMetric.IsNull() && !model.DefaultRouteMetric.IsUnknown() {
		val := int32(model.DefaultRouteMetric.ValueInt64())
		sdk.DefaultRouteMetric = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DefaultRouteMetric", "value": *sdk.DefaultRouteMetric})
	}

	// Handling Primitives
	if !model.Enable.IsNull() && !model.Enable.IsUnknown() {
		sdk.Enable = model.Enable.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Enable", "value": *sdk.Enable})
	}

	// Handling Primitives
	if !model.Passive.IsNull() && !model.Passive.IsUnknown() {
		sdk.Passive = model.Passive.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Passive", "value": *sdk.Passive})
	}

	// Handling Primitives
	if !model.Password.IsNull() && !model.Password.IsUnknown() {
		sdk.Password = model.Password.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Password", "value": sdk.Password})
	}

	// Handling Primitives
	if !model.Service.IsNull() && !model.Service.IsUnknown() {
		sdk.Service = model.Service.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Service", "value": *sdk.Service})
	}

	// Handling Objects
	if !model.StaticAddress.IsNull() && !model.StaticAddress.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field StaticAddress")
		unpacked, d := unpackEthernetInterfacesLayer3PppoeStaticAddressToSdk(ctx, model.StaticAddress)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "StaticAddress"})
		}
		if unpacked != nil {
			sdk.StaticAddress = unpacked
		}
	}

	// Handling Primitives
	if !model.Username.IsNull() && !model.Username.IsUnknown() {
		sdk.Username = model.Username.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Username", "value": sdk.Username})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.EthernetInterfacesLayer3Pppoe", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for EthernetInterfacesLayer3Pppoe ---
func packEthernetInterfacesLayer3PppoeFromSdk(ctx context.Context, sdk network_services.EthernetInterfacesLayer3Pppoe) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.EthernetInterfacesLayer3Pppoe", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.EthernetInterfacesLayer3Pppoe
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.AccessConcentrator != nil {
		model.AccessConcentrator = basetypes.NewStringValue(*sdk.AccessConcentrator)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "AccessConcentrator", "value": *sdk.AccessConcentrator})
	} else {
		model.AccessConcentrator = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Authentication != nil {
		model.Authentication = basetypes.NewStringValue(*sdk.Authentication)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Authentication", "value": *sdk.Authentication})
	} else {
		model.Authentication = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.DefaultRouteMetric != nil {
		model.DefaultRouteMetric = basetypes.NewInt64Value(int64(*sdk.DefaultRouteMetric))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DefaultRouteMetric", "value": *sdk.DefaultRouteMetric})
	} else {
		model.DefaultRouteMetric = basetypes.NewInt64Null()
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
	if sdk.Passive != nil {
		model.Passive = basetypes.NewBoolValue(*sdk.Passive)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Passive", "value": *sdk.Passive})
	} else {
		model.Passive = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	model.Password = basetypes.NewStringValue(sdk.Password)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Password", "value": sdk.Password})
	// Handling Primitives
	// Standard primitive packing
	if sdk.Service != nil {
		model.Service = basetypes.NewStringValue(*sdk.Service)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Service", "value": *sdk.Service})
	} else {
		model.Service = basetypes.NewStringNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.StaticAddress != nil {
		tflog.Debug(ctx, "Packing nested object for field StaticAddress")
		packed, d := packEthernetInterfacesLayer3PppoeStaticAddressFromSdk(ctx, *sdk.StaticAddress)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "StaticAddress"})
		}
		model.StaticAddress = packed
	} else {
		model.StaticAddress = basetypes.NewObjectNull(models.EthernetInterfacesLayer3PppoeStaticAddress{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	model.Username = basetypes.NewStringValue(sdk.Username)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Username", "value": sdk.Username})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.EthernetInterfacesLayer3Pppoe{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.EthernetInterfacesLayer3Pppoe", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for EthernetInterfacesLayer3Pppoe ---
func unpackEthernetInterfacesLayer3PppoeListToSdk(ctx context.Context, list types.List) ([]network_services.EthernetInterfacesLayer3Pppoe, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.EthernetInterfacesLayer3Pppoe")
	diags := diag.Diagnostics{}
	var data []models.EthernetInterfacesLayer3Pppoe
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.EthernetInterfacesLayer3Pppoe, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.EthernetInterfacesLayer3Pppoe{}.AttrTypes(), &item)
		unpacked, d := unpackEthernetInterfacesLayer3PppoeToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.EthernetInterfacesLayer3Pppoe", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for EthernetInterfacesLayer3Pppoe ---
func packEthernetInterfacesLayer3PppoeListFromSdk(ctx context.Context, sdks []network_services.EthernetInterfacesLayer3Pppoe) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.EthernetInterfacesLayer3Pppoe")
	diags := diag.Diagnostics{}
	var data []models.EthernetInterfacesLayer3Pppoe

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.EthernetInterfacesLayer3Pppoe
		obj, d := packEthernetInterfacesLayer3PppoeFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.EthernetInterfacesLayer3Pppoe{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.EthernetInterfacesLayer3Pppoe", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.EthernetInterfacesLayer3Pppoe{}.AttrType(), data)
}

// --- Unpacker for EthernetInterfacesLayer3PppoeStaticAddress ---
func unpackEthernetInterfacesLayer3PppoeStaticAddressToSdk(ctx context.Context, obj types.Object) (*network_services.EthernetInterfacesLayer3PppoeStaticAddress, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.EthernetInterfacesLayer3PppoeStaticAddress", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.EthernetInterfacesLayer3PppoeStaticAddress
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.EthernetInterfacesLayer3PppoeStaticAddress
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Ip.IsNull() && !model.Ip.IsUnknown() {
		sdk.Ip = model.Ip.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Ip", "value": sdk.Ip})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.EthernetInterfacesLayer3PppoeStaticAddress", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for EthernetInterfacesLayer3PppoeStaticAddress ---
func packEthernetInterfacesLayer3PppoeStaticAddressFromSdk(ctx context.Context, sdk network_services.EthernetInterfacesLayer3PppoeStaticAddress) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.EthernetInterfacesLayer3PppoeStaticAddress", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.EthernetInterfacesLayer3PppoeStaticAddress
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.Ip = basetypes.NewStringValue(sdk.Ip)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Ip", "value": sdk.Ip})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.EthernetInterfacesLayer3PppoeStaticAddress{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.EthernetInterfacesLayer3PppoeStaticAddress", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for EthernetInterfacesLayer3PppoeStaticAddress ---
func unpackEthernetInterfacesLayer3PppoeStaticAddressListToSdk(ctx context.Context, list types.List) ([]network_services.EthernetInterfacesLayer3PppoeStaticAddress, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.EthernetInterfacesLayer3PppoeStaticAddress")
	diags := diag.Diagnostics{}
	var data []models.EthernetInterfacesLayer3PppoeStaticAddress
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.EthernetInterfacesLayer3PppoeStaticAddress, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.EthernetInterfacesLayer3PppoeStaticAddress{}.AttrTypes(), &item)
		unpacked, d := unpackEthernetInterfacesLayer3PppoeStaticAddressToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.EthernetInterfacesLayer3PppoeStaticAddress", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for EthernetInterfacesLayer3PppoeStaticAddress ---
func packEthernetInterfacesLayer3PppoeStaticAddressListFromSdk(ctx context.Context, sdks []network_services.EthernetInterfacesLayer3PppoeStaticAddress) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.EthernetInterfacesLayer3PppoeStaticAddress")
	diags := diag.Diagnostics{}
	var data []models.EthernetInterfacesLayer3PppoeStaticAddress

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.EthernetInterfacesLayer3PppoeStaticAddress
		obj, d := packEthernetInterfacesLayer3PppoeStaticAddressFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.EthernetInterfacesLayer3PppoeStaticAddress{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.EthernetInterfacesLayer3PppoeStaticAddress", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.EthernetInterfacesLayer3PppoeStaticAddress{}.AttrType(), data)
}

// --- Unpacker for Poe ---
func unpackPoeToSdk(ctx context.Context, obj types.Object) (*network_services.Poe, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.Poe", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.Poe
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.Poe
	var d diag.Diagnostics
	// Handling Primitives
	if !model.PoeEnabled.IsNull() && !model.PoeEnabled.IsUnknown() {
		sdk.PoeEnabled = model.PoeEnabled.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "PoeEnabled", "value": *sdk.PoeEnabled})
	}

	// Handling Primitives
	if !model.PoeRsvdPwr.IsNull() && !model.PoeRsvdPwr.IsUnknown() {
		val := int32(model.PoeRsvdPwr.ValueInt64())
		sdk.PoeRsvdPwr = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "PoeRsvdPwr", "value": *sdk.PoeRsvdPwr})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.Poe", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for Poe ---
func packPoeFromSdk(ctx context.Context, sdk network_services.Poe) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.Poe", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.Poe
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.PoeEnabled != nil {
		model.PoeEnabled = basetypes.NewBoolValue(*sdk.PoeEnabled)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "PoeEnabled", "value": *sdk.PoeEnabled})
	} else {
		model.PoeEnabled = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.PoeRsvdPwr != nil {
		model.PoeRsvdPwr = basetypes.NewInt64Value(int64(*sdk.PoeRsvdPwr))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "PoeRsvdPwr", "value": *sdk.PoeRsvdPwr})
	} else {
		model.PoeRsvdPwr = basetypes.NewInt64Null()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.Poe{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.Poe", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for Poe ---
func unpackPoeListToSdk(ctx context.Context, list types.List) ([]network_services.Poe, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.Poe")
	diags := diag.Diagnostics{}
	var data []models.Poe
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.Poe, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.Poe{}.AttrTypes(), &item)
		unpacked, d := unpackPoeToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.Poe", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for Poe ---
func packPoeListFromSdk(ctx context.Context, sdks []network_services.Poe) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.Poe")
	diags := diag.Diagnostics{}
	var data []models.Poe

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.Poe
		obj, d := packPoeFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.Poe{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.Poe", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.Poe{}.AttrType(), data)
}
