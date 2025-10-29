package provider

import (
	"context"

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
