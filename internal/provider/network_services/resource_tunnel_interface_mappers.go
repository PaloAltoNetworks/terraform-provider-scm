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

// --- Unpacker for TunnelInterfaces ---
func unpackTunnelInterfacesToSdk(ctx context.Context, obj types.Object) (*network_services.TunnelInterfaces, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.TunnelInterfaces", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.TunnelInterfaces
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.TunnelInterfaces
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
		tflog.Debug(ctx, "Unpacking list of objects for field Ip")
		unpacked, d := unpackTunnelInterfacesIpInnerListToSdk(ctx, model.Ip)
		diags.Append(d...)
		sdk.Ip = unpacked
	}

	// Handling Objects
	if !model.Ipv6.IsNull() && !model.Ipv6.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Ipv6")
		unpacked, d := unpackTunnelInterfacesIpv6ToSdk(ctx, model.Ipv6)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Ipv6"})
		}
		if unpacked != nil {
			sdk.Ipv6 = unpacked
		}
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
	if !model.Snippet.IsNull() && !model.Snippet.IsUnknown() {
		sdk.Snippet = model.Snippet.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.TunnelInterfaces", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for TunnelInterfaces ---
func packTunnelInterfacesFromSdk(ctx context.Context, sdk network_services.TunnelInterfaces) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.TunnelInterfaces", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.TunnelInterfaces
	var d diag.Diagnostics
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
		tflog.Debug(ctx, "Packing list of objects for field Ip")
		packed, d := packTunnelInterfacesIpInnerListFromSdk(ctx, sdk.Ip)
		diags.Append(d...)
		model.Ip = packed
	} else {
		model.Ip = basetypes.NewListNull(models.TunnelInterfacesIpInner{}.AttrType())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Ipv6 != nil {
		tflog.Debug(ctx, "Packing nested object for field Ipv6")
		packed, d := packTunnelInterfacesIpv6FromSdk(ctx, *sdk.Ipv6)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Ipv6"})
		}
		model.Ipv6 = packed
	} else {
		model.Ipv6 = basetypes.NewObjectNull(models.TunnelInterfacesIpv6{}.AttrTypes())
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
	if sdk.Snippet != nil {
		model.Snippet = basetypes.NewStringValue(*sdk.Snippet)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	} else {
		model.Snippet = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.TunnelInterfaces{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.TunnelInterfaces", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for TunnelInterfaces ---
func unpackTunnelInterfacesListToSdk(ctx context.Context, list types.List) ([]network_services.TunnelInterfaces, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.TunnelInterfaces")
	diags := diag.Diagnostics{}
	var data []models.TunnelInterfaces
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.TunnelInterfaces, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.TunnelInterfaces{}.AttrTypes(), &item)
		unpacked, d := unpackTunnelInterfacesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.TunnelInterfaces", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for TunnelInterfaces ---
func packTunnelInterfacesListFromSdk(ctx context.Context, sdks []network_services.TunnelInterfaces) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.TunnelInterfaces")
	diags := diag.Diagnostics{}
	var data []models.TunnelInterfaces

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.TunnelInterfaces
		obj, d := packTunnelInterfacesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.TunnelInterfaces{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.TunnelInterfaces", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.TunnelInterfaces{}.AttrType(), data)
}

// --- Unpacker for TunnelInterfacesIpInner ---
func unpackTunnelInterfacesIpInnerToSdk(ctx context.Context, obj types.Object) (*network_services.TunnelInterfacesIpInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.TunnelInterfacesIpInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.TunnelInterfacesIpInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.TunnelInterfacesIpInner
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.TunnelInterfacesIpInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for TunnelInterfacesIpInner ---
func packTunnelInterfacesIpInnerFromSdk(ctx context.Context, sdk network_services.TunnelInterfacesIpInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.TunnelInterfacesIpInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.TunnelInterfacesIpInner
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.TunnelInterfacesIpInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.TunnelInterfacesIpInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for TunnelInterfacesIpInner ---
func unpackTunnelInterfacesIpInnerListToSdk(ctx context.Context, list types.List) ([]network_services.TunnelInterfacesIpInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.TunnelInterfacesIpInner")
	diags := diag.Diagnostics{}
	var data []models.TunnelInterfacesIpInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.TunnelInterfacesIpInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.TunnelInterfacesIpInner{}.AttrTypes(), &item)
		unpacked, d := unpackTunnelInterfacesIpInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.TunnelInterfacesIpInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for TunnelInterfacesIpInner ---
func packTunnelInterfacesIpInnerListFromSdk(ctx context.Context, sdks []network_services.TunnelInterfacesIpInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.TunnelInterfacesIpInner")
	diags := diag.Diagnostics{}
	var data []models.TunnelInterfacesIpInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.TunnelInterfacesIpInner
		obj, d := packTunnelInterfacesIpInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.TunnelInterfacesIpInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.TunnelInterfacesIpInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.TunnelInterfacesIpInner{}.AttrType(), data)
}

// --- Unpacker for TunnelInterfacesIpv6 ---
func unpackTunnelInterfacesIpv6ToSdk(ctx context.Context, obj types.Object) (*network_services.TunnelInterfacesIpv6, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.TunnelInterfacesIpv6", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.TunnelInterfacesIpv6
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.TunnelInterfacesIpv6
	var d diag.Diagnostics
	// Handling Lists
	if !model.Address.IsNull() && !model.Address.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field Address")
		unpacked, d := unpackLoopbackInterfacesIpv6AddressInnerListToSdk(ctx, model.Address)
		diags.Append(d...)
		sdk.Address = unpacked
	}

	// Handling Primitives
	if !model.Enabled.IsNull() && !model.Enabled.IsUnknown() {
		sdk.Enabled = model.Enabled.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Enabled", "value": *sdk.Enabled})
	}

	// Handling Primitives
	if !model.InterfaceId.IsNull() && !model.InterfaceId.IsUnknown() {
		sdk.InterfaceId = model.InterfaceId.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "InterfaceId", "value": *sdk.InterfaceId})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.TunnelInterfacesIpv6", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for TunnelInterfacesIpv6 ---
func packTunnelInterfacesIpv6FromSdk(ctx context.Context, sdk network_services.TunnelInterfacesIpv6) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.TunnelInterfacesIpv6", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.TunnelInterfacesIpv6
	var d diag.Diagnostics
	// Handling Lists
	if sdk.Address != nil {
		tflog.Debug(ctx, "Packing list of objects for field Address")
		packed, d := packLoopbackInterfacesIpv6AddressInnerListFromSdk(ctx, sdk.Address)
		diags.Append(d...)
		model.Address = packed
	} else {
		model.Address = basetypes.NewListNull(models.LoopbackInterfacesIpv6AddressInner{}.AttrType())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Enabled != nil {
		model.Enabled = basetypes.NewBoolValue(*sdk.Enabled)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Enabled", "value": *sdk.Enabled})
	} else {
		model.Enabled = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.InterfaceId != nil {
		model.InterfaceId = basetypes.NewStringValue(*sdk.InterfaceId)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "InterfaceId", "value": *sdk.InterfaceId})
	} else {
		model.InterfaceId = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.TunnelInterfacesIpv6{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.TunnelInterfacesIpv6", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for TunnelInterfacesIpv6 ---
func unpackTunnelInterfacesIpv6ListToSdk(ctx context.Context, list types.List) ([]network_services.TunnelInterfacesIpv6, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.TunnelInterfacesIpv6")
	diags := diag.Diagnostics{}
	var data []models.TunnelInterfacesIpv6
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.TunnelInterfacesIpv6, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.TunnelInterfacesIpv6{}.AttrTypes(), &item)
		unpacked, d := unpackTunnelInterfacesIpv6ToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.TunnelInterfacesIpv6", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for TunnelInterfacesIpv6 ---
func packTunnelInterfacesIpv6ListFromSdk(ctx context.Context, sdks []network_services.TunnelInterfacesIpv6) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.TunnelInterfacesIpv6")
	diags := diag.Diagnostics{}
	var data []models.TunnelInterfacesIpv6

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.TunnelInterfacesIpv6
		obj, d := packTunnelInterfacesIpv6FromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.TunnelInterfacesIpv6{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.TunnelInterfacesIpv6", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.TunnelInterfacesIpv6{}.AttrType(), data)
}

// --- Unpacker for LoopbackInterfacesIpv6AddressInner ---
func unpackLoopbackInterfacesIpv6AddressInnerToSdk(ctx context.Context, obj types.Object) (*network_services.LoopbackInterfacesIpv6AddressInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.LoopbackInterfacesIpv6AddressInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.LoopbackInterfacesIpv6AddressInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.LoopbackInterfacesIpv6AddressInner
	var d diag.Diagnostics
	// Handling Typeless Objects
	if !model.Anycast.IsNull() && !model.Anycast.IsUnknown() {
		tflog.Debug(ctx, "Unpacking typeless object for field Anycast")
		sdk.Anycast = make(map[string]interface{})
	}

	// Handling Primitives
	if !model.EnableOnInterface.IsNull() && !model.EnableOnInterface.IsUnknown() {
		sdk.EnableOnInterface = model.EnableOnInterface.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "EnableOnInterface", "value": *sdk.EnableOnInterface})
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
	}

	// Handling Typeless Objects
	if !model.Prefix.IsNull() && !model.Prefix.IsUnknown() {
		tflog.Debug(ctx, "Unpacking typeless object for field Prefix")
		sdk.Prefix = make(map[string]interface{})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.LoopbackInterfacesIpv6AddressInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for LoopbackInterfacesIpv6AddressInner ---
func packLoopbackInterfacesIpv6AddressInnerFromSdk(ctx context.Context, sdk network_services.LoopbackInterfacesIpv6AddressInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.LoopbackInterfacesIpv6AddressInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.LoopbackInterfacesIpv6AddressInner
	var d diag.Diagnostics
	// Handling Objects
	// This is a marker object (e.g. CHAP: {}). We just need to create an empty, non-null object.
	if sdk.Anycast != nil && !reflect.ValueOf(sdk.Anycast).IsNil() {
		tflog.Debug(ctx, "Packing typeless object for field Anycast")
		var d diag.Diagnostics
		// Create an empty object with no attributes, which signifies its presence.
		model.Anycast, d = basetypes.NewObjectValue(map[string]attr.Type{}, map[string]attr.Value{})
		diags.Append(d...)
	} else {
		// Since this field is part of a oneOf, being nil means it's not selected.
		// We make the object null with an empty attribute map.
		model.Anycast = basetypes.NewObjectNull(map[string]attr.Type{})
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.EnableOnInterface != nil {
		model.EnableOnInterface = basetypes.NewBoolValue(*sdk.EnableOnInterface)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "EnableOnInterface", "value": *sdk.EnableOnInterface})
	} else {
		model.EnableOnInterface = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Name != nil {
		model.Name = basetypes.NewStringValue(*sdk.Name)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
	} else {
		model.Name = basetypes.NewStringNull()
	}
	// Handling Objects
	// This is a marker object (e.g. CHAP: {}). We just need to create an empty, non-null object.
	if sdk.Prefix != nil && !reflect.ValueOf(sdk.Prefix).IsNil() {
		tflog.Debug(ctx, "Packing typeless object for field Prefix")
		var d diag.Diagnostics
		// Create an empty object with no attributes, which signifies its presence.
		model.Prefix, d = basetypes.NewObjectValue(map[string]attr.Type{}, map[string]attr.Value{})
		diags.Append(d...)
	} else {
		// Since this field is part of a oneOf, being nil means it's not selected.
		// We make the object null with an empty attribute map.
		model.Prefix = basetypes.NewObjectNull(map[string]attr.Type{})
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.LoopbackInterfacesIpv6AddressInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.LoopbackInterfacesIpv6AddressInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for LoopbackInterfacesIpv6AddressInner ---
func unpackLoopbackInterfacesIpv6AddressInnerListToSdk(ctx context.Context, list types.List) ([]network_services.LoopbackInterfacesIpv6AddressInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.LoopbackInterfacesIpv6AddressInner")
	diags := diag.Diagnostics{}
	var data []models.LoopbackInterfacesIpv6AddressInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.LoopbackInterfacesIpv6AddressInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.LoopbackInterfacesIpv6AddressInner{}.AttrTypes(), &item)
		unpacked, d := unpackLoopbackInterfacesIpv6AddressInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.LoopbackInterfacesIpv6AddressInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for LoopbackInterfacesIpv6AddressInner ---
func packLoopbackInterfacesIpv6AddressInnerListFromSdk(ctx context.Context, sdks []network_services.LoopbackInterfacesIpv6AddressInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.LoopbackInterfacesIpv6AddressInner")
	diags := diag.Diagnostics{}
	var data []models.LoopbackInterfacesIpv6AddressInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.LoopbackInterfacesIpv6AddressInner
		obj, d := packLoopbackInterfacesIpv6AddressInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.LoopbackInterfacesIpv6AddressInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.LoopbackInterfacesIpv6AddressInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.LoopbackInterfacesIpv6AddressInner{}.AttrType(), data)
}
