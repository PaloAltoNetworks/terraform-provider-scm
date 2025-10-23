package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/network_services"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/network_services"
)

// --- Unpacker for Layer3Subinterfaces ---
func unpackLayer3SubinterfacesToSdk(ctx context.Context, obj types.Object) (*network_services.Layer3Subinterfaces, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.Layer3Subinterfaces", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.Layer3Subinterfaces
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.Layer3Subinterfaces
	var d diag.Diagnostics

	// Handling Lists
	if !model.Arp.IsNull() && !model.Arp.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field Arp")
		unpacked, d := unpackLayer3SubinterfacesArpInnerListToSdk(ctx, model.Arp)
		diags.Append(d...)
		sdk.Arp = unpacked
	}

	// Handling Primitives
	if !model.Comment.IsNull() && !model.Comment.IsUnknown() {
		sdk.Comment = model.Comment.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Comment", "value": *sdk.Comment})
	}

	// Handling Objects
	if !model.DdnsConfig.IsNull() && !model.DdnsConfig.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field DdnsConfig")
		unpacked, d := unpackLayer3SubinterfacesDdnsConfigToSdk(ctx, model.DdnsConfig)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "DdnsConfig"})
		}
		if unpacked != nil {
			sdk.DdnsConfig = unpacked
		}
	}

	// Handling Primitives
	if !model.Device.IsNull() && !model.Device.IsUnknown() {
		sdk.Device = model.Device.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Device", "value": *sdk.Device})
	}

	// Handling Objects
	if !model.DhcpClient.IsNull() && !model.DhcpClient.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field DhcpClient")
		unpacked, d := unpackLayer3SubInterfacesDhcpClientDhcpClientToSdk(ctx, model.DhcpClient)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "DhcpClient"})
		}
		if unpacked != nil {
			sdk.DhcpClient = unpacked
		}
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

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	// Handling Primitives
	if !model.ParentInterface.IsNull() && !model.ParentInterface.IsUnknown() {
		sdk.ParentInterface = model.ParentInterface.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "ParentInterface", "value": *sdk.ParentInterface})
	}

	// Handling Primitives
	if !model.Snippet.IsNull() && !model.Snippet.IsUnknown() {
		sdk.Snippet = model.Snippet.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	}

	// Handling Primitives
	if !model.Tag.IsNull() && !model.Tag.IsUnknown() {
		val := float32(model.Tag.ValueFloat64())
		sdk.Tag = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Tag", "value": *sdk.Tag})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.Layer3Subinterfaces", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for Layer3Subinterfaces ---
func packLayer3SubinterfacesFromSdk(ctx context.Context, sdk network_services.Layer3Subinterfaces) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.Layer3Subinterfaces", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.Layer3Subinterfaces
	var d diag.Diagnostics
	// Handling Lists
	if sdk.Arp != nil {
		tflog.Debug(ctx, "Packing list of objects for field Arp")
		packed, d := packLayer3SubinterfacesArpInnerListFromSdk(ctx, sdk.Arp)
		diags.Append(d...)
		model.Arp = packed
	} else {
		model.Arp = basetypes.NewListNull(models.Layer3SubinterfacesArpInner{}.AttrType())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Comment != nil {
		model.Comment = basetypes.NewStringValue(*sdk.Comment)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Comment", "value": *sdk.Comment})
	} else {
		model.Comment = basetypes.NewStringNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.DdnsConfig != nil {
		tflog.Debug(ctx, "Packing nested object for field DdnsConfig")
		packed, d := packLayer3SubinterfacesDdnsConfigFromSdk(ctx, *sdk.DdnsConfig)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "DdnsConfig"})
		}
		model.DdnsConfig = packed
	} else {
		model.DdnsConfig = basetypes.NewObjectNull(models.Layer3SubinterfacesDdnsConfig{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Device != nil {
		model.Device = basetypes.NewStringValue(*sdk.Device)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Device", "value": *sdk.Device})
	} else {
		model.Device = basetypes.NewStringNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.DhcpClient != nil {
		tflog.Debug(ctx, "Packing nested object for field DhcpClient")
		packed, d := packLayer3SubInterfacesDhcpClientDhcpClientFromSdk(ctx, *sdk.DhcpClient)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "DhcpClient"})
		}
		model.DhcpClient = packed
	} else {
		model.DhcpClient = basetypes.NewObjectNull(models.Layer3SubInterfacesDhcpClientDhcpClient{}.AttrTypes())
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
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	// Handling Primitives
	// Standard primitive packing
	if sdk.ParentInterface != nil {
		model.ParentInterface = basetypes.NewStringValue(*sdk.ParentInterface)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "ParentInterface", "value": *sdk.ParentInterface})
	} else {
		model.ParentInterface = basetypes.NewStringNull()
	}
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
	if sdk.Tag != nil {
		model.Tag = basetypes.NewFloat64Value(float64(*sdk.Tag))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Tag", "value": *sdk.Tag})
	} else {
		model.Tag = basetypes.NewFloat64Null()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.Layer3Subinterfaces{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.Layer3Subinterfaces", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for Layer3Subinterfaces ---
func unpackLayer3SubinterfacesListToSdk(ctx context.Context, list types.List) ([]network_services.Layer3Subinterfaces, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.Layer3Subinterfaces")
	diags := diag.Diagnostics{}
	var data []models.Layer3Subinterfaces
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.Layer3Subinterfaces, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.Layer3Subinterfaces{}.AttrTypes(), &item)
		unpacked, d := unpackLayer3SubinterfacesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.Layer3Subinterfaces", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for Layer3Subinterfaces ---
func packLayer3SubinterfacesListFromSdk(ctx context.Context, sdks []network_services.Layer3Subinterfaces) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.Layer3Subinterfaces")
	diags := diag.Diagnostics{}
	var data []models.Layer3Subinterfaces

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.Layer3Subinterfaces
		obj, d := packLayer3SubinterfacesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.Layer3Subinterfaces{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.Layer3Subinterfaces", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.Layer3Subinterfaces{}.AttrType(), data)
}

// --- Unpacker for Layer3SubinterfacesArpInner ---
func unpackLayer3SubinterfacesArpInnerToSdk(ctx context.Context, obj types.Object) (*network_services.Layer3SubinterfacesArpInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.Layer3SubinterfacesArpInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.Layer3SubinterfacesArpInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.Layer3SubinterfacesArpInner
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

	tflog.Debug(ctx, "Exiting unpack helper for models.Layer3SubinterfacesArpInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for Layer3SubinterfacesArpInner ---
func packLayer3SubinterfacesArpInnerFromSdk(ctx context.Context, sdk network_services.Layer3SubinterfacesArpInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.Layer3SubinterfacesArpInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.Layer3SubinterfacesArpInner
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

	obj, d := types.ObjectValueFrom(ctx, models.Layer3SubinterfacesArpInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.Layer3SubinterfacesArpInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for Layer3SubinterfacesArpInner ---
func unpackLayer3SubinterfacesArpInnerListToSdk(ctx context.Context, list types.List) ([]network_services.Layer3SubinterfacesArpInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.Layer3SubinterfacesArpInner")
	diags := diag.Diagnostics{}
	var data []models.Layer3SubinterfacesArpInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.Layer3SubinterfacesArpInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.Layer3SubinterfacesArpInner{}.AttrTypes(), &item)
		unpacked, d := unpackLayer3SubinterfacesArpInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.Layer3SubinterfacesArpInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for Layer3SubinterfacesArpInner ---
func packLayer3SubinterfacesArpInnerListFromSdk(ctx context.Context, sdks []network_services.Layer3SubinterfacesArpInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.Layer3SubinterfacesArpInner")
	diags := diag.Diagnostics{}
	var data []models.Layer3SubinterfacesArpInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.Layer3SubinterfacesArpInner
		obj, d := packLayer3SubinterfacesArpInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.Layer3SubinterfacesArpInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.Layer3SubinterfacesArpInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.Layer3SubinterfacesArpInner{}.AttrType(), data)
}

// --- Unpacker for Layer3SubinterfacesDdnsConfig ---
func unpackLayer3SubinterfacesDdnsConfigToSdk(ctx context.Context, obj types.Object) (*network_services.Layer3SubinterfacesDdnsConfig, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.Layer3SubinterfacesDdnsConfig", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.Layer3SubinterfacesDdnsConfig
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.Layer3SubinterfacesDdnsConfig
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

	tflog.Debug(ctx, "Exiting unpack helper for models.Layer3SubinterfacesDdnsConfig", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for Layer3SubinterfacesDdnsConfig ---
func packLayer3SubinterfacesDdnsConfigFromSdk(ctx context.Context, sdk network_services.Layer3SubinterfacesDdnsConfig) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.Layer3SubinterfacesDdnsConfig", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.Layer3SubinterfacesDdnsConfig
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

	obj, d := types.ObjectValueFrom(ctx, models.Layer3SubinterfacesDdnsConfig{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.Layer3SubinterfacesDdnsConfig", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for Layer3SubinterfacesDdnsConfig ---
func unpackLayer3SubinterfacesDdnsConfigListToSdk(ctx context.Context, list types.List) ([]network_services.Layer3SubinterfacesDdnsConfig, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.Layer3SubinterfacesDdnsConfig")
	diags := diag.Diagnostics{}
	var data []models.Layer3SubinterfacesDdnsConfig
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.Layer3SubinterfacesDdnsConfig, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.Layer3SubinterfacesDdnsConfig{}.AttrTypes(), &item)
		unpacked, d := unpackLayer3SubinterfacesDdnsConfigToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.Layer3SubinterfacesDdnsConfig", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for Layer3SubinterfacesDdnsConfig ---
func packLayer3SubinterfacesDdnsConfigListFromSdk(ctx context.Context, sdks []network_services.Layer3SubinterfacesDdnsConfig) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.Layer3SubinterfacesDdnsConfig")
	diags := diag.Diagnostics{}
	var data []models.Layer3SubinterfacesDdnsConfig

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.Layer3SubinterfacesDdnsConfig
		obj, d := packLayer3SubinterfacesDdnsConfigFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.Layer3SubinterfacesDdnsConfig{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.Layer3SubinterfacesDdnsConfig", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.Layer3SubinterfacesDdnsConfig{}.AttrType(), data)
}

// --- Unpacker for Layer3SubInterfacesDhcpClientDhcpClient ---
func unpackLayer3SubInterfacesDhcpClientDhcpClientToSdk(ctx context.Context, obj types.Object) (*network_services.Layer3SubInterfacesDhcpClientDhcpClient, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.Layer3SubInterfacesDhcpClientDhcpClient", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.Layer3SubInterfacesDhcpClientDhcpClient
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.Layer3SubInterfacesDhcpClientDhcpClient
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
		unpacked, d := unpackLayer3SubInterfacesDhcpClientDhcpClientSendHostnameToSdk(ctx, model.SendHostname)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "SendHostname"})
		}
		if unpacked != nil {
			sdk.SendHostname = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.Layer3SubInterfacesDhcpClientDhcpClient", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for Layer3SubInterfacesDhcpClientDhcpClient ---
func packLayer3SubInterfacesDhcpClientDhcpClientFromSdk(ctx context.Context, sdk network_services.Layer3SubInterfacesDhcpClientDhcpClient) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.Layer3SubInterfacesDhcpClientDhcpClient", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.Layer3SubInterfacesDhcpClientDhcpClient
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
		packed, d := packLayer3SubInterfacesDhcpClientDhcpClientSendHostnameFromSdk(ctx, *sdk.SendHostname)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "SendHostname"})
		}
		model.SendHostname = packed
	} else {
		model.SendHostname = basetypes.NewObjectNull(models.Layer3SubInterfacesDhcpClientDhcpClientSendHostname{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.Layer3SubInterfacesDhcpClientDhcpClient{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.Layer3SubInterfacesDhcpClientDhcpClient", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for Layer3SubInterfacesDhcpClientDhcpClient ---
func unpackLayer3SubInterfacesDhcpClientDhcpClientListToSdk(ctx context.Context, list types.List) ([]network_services.Layer3SubInterfacesDhcpClientDhcpClient, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.Layer3SubInterfacesDhcpClientDhcpClient")
	diags := diag.Diagnostics{}
	var data []models.Layer3SubInterfacesDhcpClientDhcpClient
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.Layer3SubInterfacesDhcpClientDhcpClient, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.Layer3SubInterfacesDhcpClientDhcpClient{}.AttrTypes(), &item)
		unpacked, d := unpackLayer3SubInterfacesDhcpClientDhcpClientToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.Layer3SubInterfacesDhcpClientDhcpClient", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for Layer3SubInterfacesDhcpClientDhcpClient ---
func packLayer3SubInterfacesDhcpClientDhcpClientListFromSdk(ctx context.Context, sdks []network_services.Layer3SubInterfacesDhcpClientDhcpClient) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.Layer3SubInterfacesDhcpClientDhcpClient")
	diags := diag.Diagnostics{}
	var data []models.Layer3SubInterfacesDhcpClientDhcpClient

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.Layer3SubInterfacesDhcpClientDhcpClient
		obj, d := packLayer3SubInterfacesDhcpClientDhcpClientFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.Layer3SubInterfacesDhcpClientDhcpClient{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.Layer3SubInterfacesDhcpClientDhcpClient", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.Layer3SubInterfacesDhcpClientDhcpClient{}.AttrType(), data)
}

// --- Unpacker for Layer3SubInterfacesDhcpClientDhcpClientSendHostname ---
func unpackLayer3SubInterfacesDhcpClientDhcpClientSendHostnameToSdk(ctx context.Context, obj types.Object) (*network_services.Layer3SubInterfacesDhcpClientDhcpClientSendHostname, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.Layer3SubInterfacesDhcpClientDhcpClientSendHostname", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.Layer3SubInterfacesDhcpClientDhcpClientSendHostname
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.Layer3SubInterfacesDhcpClientDhcpClientSendHostname
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

	tflog.Debug(ctx, "Exiting unpack helper for models.Layer3SubInterfacesDhcpClientDhcpClientSendHostname", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for Layer3SubInterfacesDhcpClientDhcpClientSendHostname ---
func packLayer3SubInterfacesDhcpClientDhcpClientSendHostnameFromSdk(ctx context.Context, sdk network_services.Layer3SubInterfacesDhcpClientDhcpClientSendHostname) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.Layer3SubInterfacesDhcpClientDhcpClientSendHostname", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.Layer3SubInterfacesDhcpClientDhcpClientSendHostname
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

	obj, d := types.ObjectValueFrom(ctx, models.Layer3SubInterfacesDhcpClientDhcpClientSendHostname{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.Layer3SubInterfacesDhcpClientDhcpClientSendHostname", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for Layer3SubInterfacesDhcpClientDhcpClientSendHostname ---
func unpackLayer3SubInterfacesDhcpClientDhcpClientSendHostnameListToSdk(ctx context.Context, list types.List) ([]network_services.Layer3SubInterfacesDhcpClientDhcpClientSendHostname, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.Layer3SubInterfacesDhcpClientDhcpClientSendHostname")
	diags := diag.Diagnostics{}
	var data []models.Layer3SubInterfacesDhcpClientDhcpClientSendHostname
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.Layer3SubInterfacesDhcpClientDhcpClientSendHostname, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.Layer3SubInterfacesDhcpClientDhcpClientSendHostname{}.AttrTypes(), &item)
		unpacked, d := unpackLayer3SubInterfacesDhcpClientDhcpClientSendHostnameToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.Layer3SubInterfacesDhcpClientDhcpClientSendHostname", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for Layer3SubInterfacesDhcpClientDhcpClientSendHostname ---
func packLayer3SubInterfacesDhcpClientDhcpClientSendHostnameListFromSdk(ctx context.Context, sdks []network_services.Layer3SubInterfacesDhcpClientDhcpClientSendHostname) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.Layer3SubInterfacesDhcpClientDhcpClientSendHostname")
	diags := diag.Diagnostics{}
	var data []models.Layer3SubInterfacesDhcpClientDhcpClientSendHostname

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.Layer3SubInterfacesDhcpClientDhcpClientSendHostname
		obj, d := packLayer3SubInterfacesDhcpClientDhcpClientSendHostnameFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.Layer3SubInterfacesDhcpClientDhcpClientSendHostname{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.Layer3SubInterfacesDhcpClientDhcpClientSendHostname", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.Layer3SubInterfacesDhcpClientDhcpClientSendHostname{}.AttrType(), data)
}
