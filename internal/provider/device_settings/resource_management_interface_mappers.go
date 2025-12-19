package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/device_settings"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/device_settings"
)

// --- Unpacker for ManagementInterface ---
func unpackManagementInterfaceToSdk(ctx context.Context, obj types.Object) (*device_settings.ManagementInterface, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ManagementInterface", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ManagementInterface
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk device_settings.ManagementInterface
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
		sdk.Id = model.Id.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Id", "value": *sdk.Id})
	}

	// Handling Objects
	if !model.ManagementInterface.IsNull() && !model.ManagementInterface.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field ManagementInterface")
		unpacked, d := unpackManagementInterfaceManagementInterfaceToSdk(ctx, model.ManagementInterface)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "ManagementInterface"})
		}
		if unpacked != nil {
			sdk.ManagementInterface = unpacked
		}
	}

	// Handling Primitives
	if !model.Snippet.IsNull() && !model.Snippet.IsUnknown() {
		sdk.Snippet = model.Snippet.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ManagementInterface", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ManagementInterface ---
func packManagementInterfaceFromSdk(ctx context.Context, sdk device_settings.ManagementInterface) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ManagementInterface", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ManagementInterface
	var d diag.Diagnostics
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
	if sdk.ManagementInterface != nil {
		tflog.Debug(ctx, "Packing nested object for field ManagementInterface")
		packed, d := packManagementInterfaceManagementInterfaceFromSdk(ctx, *sdk.ManagementInterface)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "ManagementInterface"})
		}
		model.ManagementInterface = packed
	} else {
		model.ManagementInterface = basetypes.NewObjectNull(models.ManagementInterfaceManagementInterface{}.AttrTypes())
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

	obj, d := types.ObjectValueFrom(ctx, models.ManagementInterface{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ManagementInterface", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ManagementInterface ---
func unpackManagementInterfaceListToSdk(ctx context.Context, list types.List) ([]device_settings.ManagementInterface, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ManagementInterface")
	diags := diag.Diagnostics{}
	var data []models.ManagementInterface
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]device_settings.ManagementInterface, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ManagementInterface{}.AttrTypes(), &item)
		unpacked, d := unpackManagementInterfaceToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ManagementInterface", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ManagementInterface ---
func packManagementInterfaceListFromSdk(ctx context.Context, sdks []device_settings.ManagementInterface) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ManagementInterface")
	diags := diag.Diagnostics{}
	var data []models.ManagementInterface

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ManagementInterface
		obj, d := packManagementInterfaceFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ManagementInterface{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ManagementInterface", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ManagementInterface{}.AttrType(), data)
}

// --- Unpacker for ManagementInterfaceManagementInterface ---
func unpackManagementInterfaceManagementInterfaceToSdk(ctx context.Context, obj types.Object) (*device_settings.ManagementInterfaceManagementInterface, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ManagementInterfaceManagementInterface", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ManagementInterfaceManagementInterface
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk device_settings.ManagementInterfaceManagementInterface
	var d diag.Diagnostics
	// Handling Objects
	if !model.MgmtType.IsNull() && !model.MgmtType.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field MgmtType")
		unpacked, d := unpackManagementInterfaceManagementInterfaceMgmtTypeToSdk(ctx, model.MgmtType)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "MgmtType"})
		}
		if unpacked != nil {
			sdk.MgmtType = unpacked
		}
	}

	// Handling Primitives
	if !model.Mtu.IsNull() && !model.Mtu.IsUnknown() {
		val := int32(model.Mtu.ValueInt64())
		sdk.Mtu = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Mtu", "value": *sdk.Mtu})
	}

	// Handling Lists
	if !model.PermittedIp.IsNull() && !model.PermittedIp.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field PermittedIp")
		unpacked, d := unpackManagementInterfaceManagementInterfacePermittedIpInnerListToSdk(ctx, model.PermittedIp)
		diags.Append(d...)
		sdk.PermittedIp = unpacked
	}

	// Handling Objects
	if !model.Service.IsNull() && !model.Service.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Service")
		unpacked, d := unpackManagementInterfaceManagementInterfaceServiceToSdk(ctx, model.Service)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Service"})
		}
		if unpacked != nil {
			sdk.Service = unpacked
		}
	}

	// Handling Primitives
	if !model.SpeedDuplex.IsNull() && !model.SpeedDuplex.IsUnknown() {
		sdk.SpeedDuplex = model.SpeedDuplex.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "SpeedDuplex", "value": *sdk.SpeedDuplex})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ManagementInterfaceManagementInterface", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ManagementInterfaceManagementInterface ---
func packManagementInterfaceManagementInterfaceFromSdk(ctx context.Context, sdk device_settings.ManagementInterfaceManagementInterface) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ManagementInterfaceManagementInterface", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ManagementInterfaceManagementInterface
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.MgmtType != nil {
		tflog.Debug(ctx, "Packing nested object for field MgmtType")
		packed, d := packManagementInterfaceManagementInterfaceMgmtTypeFromSdk(ctx, *sdk.MgmtType)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "MgmtType"})
		}
		model.MgmtType = packed
	} else {
		model.MgmtType = basetypes.NewObjectNull(models.ManagementInterfaceManagementInterfaceMgmtType{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Mtu != nil {
		model.Mtu = basetypes.NewInt64Value(int64(*sdk.Mtu))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Mtu", "value": *sdk.Mtu})
	} else {
		model.Mtu = basetypes.NewInt64Null()
	}
	// Handling Lists
	if sdk.PermittedIp != nil {
		tflog.Debug(ctx, "Packing list of objects for field PermittedIp")
		packed, d := packManagementInterfaceManagementInterfacePermittedIpInnerListFromSdk(ctx, sdk.PermittedIp)
		diags.Append(d...)
		model.PermittedIp = packed
	} else {
		model.PermittedIp = basetypes.NewListNull(models.ManagementInterfaceManagementInterfacePermittedIpInner{}.AttrType())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Service != nil {
		tflog.Debug(ctx, "Packing nested object for field Service")
		packed, d := packManagementInterfaceManagementInterfaceServiceFromSdk(ctx, *sdk.Service)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Service"})
		}
		model.Service = packed
	} else {
		model.Service = basetypes.NewObjectNull(models.ManagementInterfaceManagementInterfaceService{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.SpeedDuplex != nil {
		model.SpeedDuplex = basetypes.NewStringValue(*sdk.SpeedDuplex)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "SpeedDuplex", "value": *sdk.SpeedDuplex})
	} else {
		model.SpeedDuplex = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ManagementInterfaceManagementInterface{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ManagementInterfaceManagementInterface", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ManagementInterfaceManagementInterface ---
func unpackManagementInterfaceManagementInterfaceListToSdk(ctx context.Context, list types.List) ([]device_settings.ManagementInterfaceManagementInterface, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ManagementInterfaceManagementInterface")
	diags := diag.Diagnostics{}
	var data []models.ManagementInterfaceManagementInterface
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]device_settings.ManagementInterfaceManagementInterface, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ManagementInterfaceManagementInterface{}.AttrTypes(), &item)
		unpacked, d := unpackManagementInterfaceManagementInterfaceToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ManagementInterfaceManagementInterface", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ManagementInterfaceManagementInterface ---
func packManagementInterfaceManagementInterfaceListFromSdk(ctx context.Context, sdks []device_settings.ManagementInterfaceManagementInterface) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ManagementInterfaceManagementInterface")
	diags := diag.Diagnostics{}
	var data []models.ManagementInterfaceManagementInterface

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ManagementInterfaceManagementInterface
		obj, d := packManagementInterfaceManagementInterfaceFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ManagementInterfaceManagementInterface{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ManagementInterfaceManagementInterface", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ManagementInterfaceManagementInterface{}.AttrType(), data)
}

// --- Unpacker for ManagementInterfaceManagementInterfaceMgmtType ---
func unpackManagementInterfaceManagementInterfaceMgmtTypeToSdk(ctx context.Context, obj types.Object) (*device_settings.ManagementInterfaceManagementInterfaceMgmtType, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ManagementInterfaceManagementInterfaceMgmtType", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ManagementInterfaceManagementInterfaceMgmtType
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk device_settings.ManagementInterfaceManagementInterfaceMgmtType
	var d diag.Diagnostics
	// Handling Objects
	if !model.DhcpClient.IsNull() && !model.DhcpClient.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field DhcpClient")
		unpacked, d := unpackManagementInterfaceManagementInterfaceMgmtTypeDhcpClientToSdk(ctx, model.DhcpClient)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "DhcpClient"})
		}
		if unpacked != nil {
			sdk.DhcpClient = unpacked
		}
	}

	// Handling Objects
	if !model.Static.IsNull() && !model.Static.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Static")
		unpacked, d := unpackManagementInterfaceManagementInterfaceMgmtTypeStaticToSdk(ctx, model.Static)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Static"})
		}
		if unpacked != nil {
			sdk.Static = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ManagementInterfaceManagementInterfaceMgmtType", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ManagementInterfaceManagementInterfaceMgmtType ---
func packManagementInterfaceManagementInterfaceMgmtTypeFromSdk(ctx context.Context, sdk device_settings.ManagementInterfaceManagementInterfaceMgmtType) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ManagementInterfaceManagementInterfaceMgmtType", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ManagementInterfaceManagementInterfaceMgmtType
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.DhcpClient != nil {
		tflog.Debug(ctx, "Packing nested object for field DhcpClient")
		packed, d := packManagementInterfaceManagementInterfaceMgmtTypeDhcpClientFromSdk(ctx, *sdk.DhcpClient)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "DhcpClient"})
		}
		model.DhcpClient = packed
	} else {
		model.DhcpClient = basetypes.NewObjectNull(models.ManagementInterfaceManagementInterfaceMgmtTypeDhcpClient{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Static != nil {
		tflog.Debug(ctx, "Packing nested object for field Static")
		packed, d := packManagementInterfaceManagementInterfaceMgmtTypeStaticFromSdk(ctx, *sdk.Static)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Static"})
		}
		model.Static = packed
	} else {
		model.Static = basetypes.NewObjectNull(models.ManagementInterfaceManagementInterfaceMgmtTypeStatic{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ManagementInterfaceManagementInterfaceMgmtType{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ManagementInterfaceManagementInterfaceMgmtType", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ManagementInterfaceManagementInterfaceMgmtType ---
func unpackManagementInterfaceManagementInterfaceMgmtTypeListToSdk(ctx context.Context, list types.List) ([]device_settings.ManagementInterfaceManagementInterfaceMgmtType, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ManagementInterfaceManagementInterfaceMgmtType")
	diags := diag.Diagnostics{}
	var data []models.ManagementInterfaceManagementInterfaceMgmtType
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]device_settings.ManagementInterfaceManagementInterfaceMgmtType, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ManagementInterfaceManagementInterfaceMgmtType{}.AttrTypes(), &item)
		unpacked, d := unpackManagementInterfaceManagementInterfaceMgmtTypeToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ManagementInterfaceManagementInterfaceMgmtType", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ManagementInterfaceManagementInterfaceMgmtType ---
func packManagementInterfaceManagementInterfaceMgmtTypeListFromSdk(ctx context.Context, sdks []device_settings.ManagementInterfaceManagementInterfaceMgmtType) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ManagementInterfaceManagementInterfaceMgmtType")
	diags := diag.Diagnostics{}
	var data []models.ManagementInterfaceManagementInterfaceMgmtType

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ManagementInterfaceManagementInterfaceMgmtType
		obj, d := packManagementInterfaceManagementInterfaceMgmtTypeFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ManagementInterfaceManagementInterfaceMgmtType{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ManagementInterfaceManagementInterfaceMgmtType", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ManagementInterfaceManagementInterfaceMgmtType{}.AttrType(), data)
}

// --- Unpacker for ManagementInterfaceManagementInterfaceMgmtTypeDhcpClient ---
func unpackManagementInterfaceManagementInterfaceMgmtTypeDhcpClientToSdk(ctx context.Context, obj types.Object) (*device_settings.ManagementInterfaceManagementInterfaceMgmtTypeDhcpClient, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ManagementInterfaceManagementInterfaceMgmtTypeDhcpClient", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ManagementInterfaceManagementInterfaceMgmtTypeDhcpClient
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk device_settings.ManagementInterfaceManagementInterfaceMgmtTypeDhcpClient
	var d diag.Diagnostics
	// Handling Primitives
	if !model.AcceptDhcpDomain.IsNull() && !model.AcceptDhcpDomain.IsUnknown() {
		sdk.AcceptDhcpDomain = model.AcceptDhcpDomain.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "AcceptDhcpDomain", "value": *sdk.AcceptDhcpDomain})
	}

	// Handling Primitives
	if !model.AcceptDhcpHostname.IsNull() && !model.AcceptDhcpHostname.IsUnknown() {
		sdk.AcceptDhcpHostname = model.AcceptDhcpHostname.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "AcceptDhcpHostname", "value": *sdk.AcceptDhcpHostname})
	}

	// Handling Primitives
	if !model.SendClientId.IsNull() && !model.SendClientId.IsUnknown() {
		sdk.SendClientId = model.SendClientId.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "SendClientId", "value": *sdk.SendClientId})
	}

	// Handling Primitives
	if !model.SendHostname.IsNull() && !model.SendHostname.IsUnknown() {
		sdk.SendHostname = model.SendHostname.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "SendHostname", "value": *sdk.SendHostname})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ManagementInterfaceManagementInterfaceMgmtTypeDhcpClient", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ManagementInterfaceManagementInterfaceMgmtTypeDhcpClient ---
func packManagementInterfaceManagementInterfaceMgmtTypeDhcpClientFromSdk(ctx context.Context, sdk device_settings.ManagementInterfaceManagementInterfaceMgmtTypeDhcpClient) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ManagementInterfaceManagementInterfaceMgmtTypeDhcpClient", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ManagementInterfaceManagementInterfaceMgmtTypeDhcpClient
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.AcceptDhcpDomain != nil {
		model.AcceptDhcpDomain = basetypes.NewBoolValue(*sdk.AcceptDhcpDomain)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "AcceptDhcpDomain", "value": *sdk.AcceptDhcpDomain})
	} else {
		model.AcceptDhcpDomain = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.AcceptDhcpHostname != nil {
		model.AcceptDhcpHostname = basetypes.NewBoolValue(*sdk.AcceptDhcpHostname)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "AcceptDhcpHostname", "value": *sdk.AcceptDhcpHostname})
	} else {
		model.AcceptDhcpHostname = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.SendClientId != nil {
		model.SendClientId = basetypes.NewBoolValue(*sdk.SendClientId)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "SendClientId", "value": *sdk.SendClientId})
	} else {
		model.SendClientId = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.SendHostname != nil {
		model.SendHostname = basetypes.NewBoolValue(*sdk.SendHostname)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "SendHostname", "value": *sdk.SendHostname})
	} else {
		model.SendHostname = basetypes.NewBoolNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ManagementInterfaceManagementInterfaceMgmtTypeDhcpClient{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ManagementInterfaceManagementInterfaceMgmtTypeDhcpClient", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ManagementInterfaceManagementInterfaceMgmtTypeDhcpClient ---
func unpackManagementInterfaceManagementInterfaceMgmtTypeDhcpClientListToSdk(ctx context.Context, list types.List) ([]device_settings.ManagementInterfaceManagementInterfaceMgmtTypeDhcpClient, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ManagementInterfaceManagementInterfaceMgmtTypeDhcpClient")
	diags := diag.Diagnostics{}
	var data []models.ManagementInterfaceManagementInterfaceMgmtTypeDhcpClient
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]device_settings.ManagementInterfaceManagementInterfaceMgmtTypeDhcpClient, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ManagementInterfaceManagementInterfaceMgmtTypeDhcpClient{}.AttrTypes(), &item)
		unpacked, d := unpackManagementInterfaceManagementInterfaceMgmtTypeDhcpClientToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ManagementInterfaceManagementInterfaceMgmtTypeDhcpClient", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ManagementInterfaceManagementInterfaceMgmtTypeDhcpClient ---
func packManagementInterfaceManagementInterfaceMgmtTypeDhcpClientListFromSdk(ctx context.Context, sdks []device_settings.ManagementInterfaceManagementInterfaceMgmtTypeDhcpClient) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ManagementInterfaceManagementInterfaceMgmtTypeDhcpClient")
	diags := diag.Diagnostics{}
	var data []models.ManagementInterfaceManagementInterfaceMgmtTypeDhcpClient

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ManagementInterfaceManagementInterfaceMgmtTypeDhcpClient
		obj, d := packManagementInterfaceManagementInterfaceMgmtTypeDhcpClientFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ManagementInterfaceManagementInterfaceMgmtTypeDhcpClient{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ManagementInterfaceManagementInterfaceMgmtTypeDhcpClient", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ManagementInterfaceManagementInterfaceMgmtTypeDhcpClient{}.AttrType(), data)
}

// --- Unpacker for ManagementInterfaceManagementInterfaceMgmtTypeStatic ---
func unpackManagementInterfaceManagementInterfaceMgmtTypeStaticToSdk(ctx context.Context, obj types.Object) (*device_settings.ManagementInterfaceManagementInterfaceMgmtTypeStatic, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ManagementInterfaceManagementInterfaceMgmtTypeStatic", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ManagementInterfaceManagementInterfaceMgmtTypeStatic
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk device_settings.ManagementInterfaceManagementInterfaceMgmtTypeStatic
	var d diag.Diagnostics
	// Handling Primitives
	if !model.DefaultGateway.IsNull() && !model.DefaultGateway.IsUnknown() {
		sdk.DefaultGateway = model.DefaultGateway.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "DefaultGateway", "value": sdk.DefaultGateway})
	}

	// Handling Primitives
	if !model.IpAddress.IsNull() && !model.IpAddress.IsUnknown() {
		sdk.IpAddress = model.IpAddress.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "IpAddress", "value": sdk.IpAddress})
	}

	// Handling Primitives
	if !model.Netmask.IsNull() && !model.Netmask.IsUnknown() {
		sdk.Netmask = model.Netmask.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Netmask", "value": sdk.Netmask})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ManagementInterfaceManagementInterfaceMgmtTypeStatic", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ManagementInterfaceManagementInterfaceMgmtTypeStatic ---
func packManagementInterfaceManagementInterfaceMgmtTypeStaticFromSdk(ctx context.Context, sdk device_settings.ManagementInterfaceManagementInterfaceMgmtTypeStatic) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ManagementInterfaceManagementInterfaceMgmtTypeStatic", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ManagementInterfaceManagementInterfaceMgmtTypeStatic
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.DefaultGateway = basetypes.NewStringValue(sdk.DefaultGateway)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "DefaultGateway", "value": sdk.DefaultGateway})
	// Handling Primitives
	// Standard primitive packing
	model.IpAddress = basetypes.NewStringValue(sdk.IpAddress)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "IpAddress", "value": sdk.IpAddress})
	// Handling Primitives
	// Standard primitive packing
	model.Netmask = basetypes.NewStringValue(sdk.Netmask)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Netmask", "value": sdk.Netmask})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ManagementInterfaceManagementInterfaceMgmtTypeStatic{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ManagementInterfaceManagementInterfaceMgmtTypeStatic", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ManagementInterfaceManagementInterfaceMgmtTypeStatic ---
func unpackManagementInterfaceManagementInterfaceMgmtTypeStaticListToSdk(ctx context.Context, list types.List) ([]device_settings.ManagementInterfaceManagementInterfaceMgmtTypeStatic, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ManagementInterfaceManagementInterfaceMgmtTypeStatic")
	diags := diag.Diagnostics{}
	var data []models.ManagementInterfaceManagementInterfaceMgmtTypeStatic
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]device_settings.ManagementInterfaceManagementInterfaceMgmtTypeStatic, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ManagementInterfaceManagementInterfaceMgmtTypeStatic{}.AttrTypes(), &item)
		unpacked, d := unpackManagementInterfaceManagementInterfaceMgmtTypeStaticToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ManagementInterfaceManagementInterfaceMgmtTypeStatic", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ManagementInterfaceManagementInterfaceMgmtTypeStatic ---
func packManagementInterfaceManagementInterfaceMgmtTypeStaticListFromSdk(ctx context.Context, sdks []device_settings.ManagementInterfaceManagementInterfaceMgmtTypeStatic) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ManagementInterfaceManagementInterfaceMgmtTypeStatic")
	diags := diag.Diagnostics{}
	var data []models.ManagementInterfaceManagementInterfaceMgmtTypeStatic

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ManagementInterfaceManagementInterfaceMgmtTypeStatic
		obj, d := packManagementInterfaceManagementInterfaceMgmtTypeStaticFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ManagementInterfaceManagementInterfaceMgmtTypeStatic{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ManagementInterfaceManagementInterfaceMgmtTypeStatic", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ManagementInterfaceManagementInterfaceMgmtTypeStatic{}.AttrType(), data)
}

// --- Unpacker for ManagementInterfaceManagementInterfacePermittedIpInner ---
func unpackManagementInterfaceManagementInterfacePermittedIpInnerToSdk(ctx context.Context, obj types.Object) (*device_settings.ManagementInterfaceManagementInterfacePermittedIpInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ManagementInterfaceManagementInterfacePermittedIpInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ManagementInterfaceManagementInterfacePermittedIpInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk device_settings.ManagementInterfaceManagementInterfacePermittedIpInner
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Description.IsNull() && !model.Description.IsUnknown() {
		sdk.Description = model.Description.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Description", "value": *sdk.Description})
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ManagementInterfaceManagementInterfacePermittedIpInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ManagementInterfaceManagementInterfacePermittedIpInner ---
func packManagementInterfaceManagementInterfacePermittedIpInnerFromSdk(ctx context.Context, sdk device_settings.ManagementInterfaceManagementInterfacePermittedIpInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ManagementInterfaceManagementInterfacePermittedIpInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ManagementInterfaceManagementInterfacePermittedIpInner
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Description != nil {
		model.Description = basetypes.NewStringValue(*sdk.Description)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Description", "value": *sdk.Description})
	} else {
		model.Description = basetypes.NewStringNull()
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

	obj, d := types.ObjectValueFrom(ctx, models.ManagementInterfaceManagementInterfacePermittedIpInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ManagementInterfaceManagementInterfacePermittedIpInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ManagementInterfaceManagementInterfacePermittedIpInner ---
func unpackManagementInterfaceManagementInterfacePermittedIpInnerListToSdk(ctx context.Context, list types.List) ([]device_settings.ManagementInterfaceManagementInterfacePermittedIpInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ManagementInterfaceManagementInterfacePermittedIpInner")
	diags := diag.Diagnostics{}
	var data []models.ManagementInterfaceManagementInterfacePermittedIpInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]device_settings.ManagementInterfaceManagementInterfacePermittedIpInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ManagementInterfaceManagementInterfacePermittedIpInner{}.AttrTypes(), &item)
		unpacked, d := unpackManagementInterfaceManagementInterfacePermittedIpInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ManagementInterfaceManagementInterfacePermittedIpInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ManagementInterfaceManagementInterfacePermittedIpInner ---
func packManagementInterfaceManagementInterfacePermittedIpInnerListFromSdk(ctx context.Context, sdks []device_settings.ManagementInterfaceManagementInterfacePermittedIpInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ManagementInterfaceManagementInterfacePermittedIpInner")
	diags := diag.Diagnostics{}
	var data []models.ManagementInterfaceManagementInterfacePermittedIpInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ManagementInterfaceManagementInterfacePermittedIpInner
		obj, d := packManagementInterfaceManagementInterfacePermittedIpInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ManagementInterfaceManagementInterfacePermittedIpInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ManagementInterfaceManagementInterfacePermittedIpInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ManagementInterfaceManagementInterfacePermittedIpInner{}.AttrType(), data)
}

// --- Unpacker for ManagementInterfaceManagementInterfaceService ---
func unpackManagementInterfaceManagementInterfaceServiceToSdk(ctx context.Context, obj types.Object) (*device_settings.ManagementInterfaceManagementInterfaceService, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ManagementInterfaceManagementInterfaceService", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ManagementInterfaceManagementInterfaceService
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk device_settings.ManagementInterfaceManagementInterfaceService
	var d diag.Diagnostics
	// Handling Primitives
	if !model.DisableHttp.IsNull() && !model.DisableHttp.IsUnknown() {
		sdk.DisableHttp = model.DisableHttp.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DisableHttp", "value": *sdk.DisableHttp})
	}

	// Handling Primitives
	if !model.DisableHttpOcsp.IsNull() && !model.DisableHttpOcsp.IsUnknown() {
		sdk.DisableHttpOcsp = model.DisableHttpOcsp.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DisableHttpOcsp", "value": *sdk.DisableHttpOcsp})
	}

	// Handling Primitives
	if !model.DisableHttps.IsNull() && !model.DisableHttps.IsUnknown() {
		sdk.DisableHttps = model.DisableHttps.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DisableHttps", "value": *sdk.DisableHttps})
	}

	// Handling Primitives
	if !model.DisableIcmp.IsNull() && !model.DisableIcmp.IsUnknown() {
		sdk.DisableIcmp = model.DisableIcmp.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DisableIcmp", "value": *sdk.DisableIcmp})
	}

	// Handling Primitives
	if !model.DisableSnmp.IsNull() && !model.DisableSnmp.IsUnknown() {
		sdk.DisableSnmp = model.DisableSnmp.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DisableSnmp", "value": *sdk.DisableSnmp})
	}

	// Handling Primitives
	if !model.DisableSsh.IsNull() && !model.DisableSsh.IsUnknown() {
		sdk.DisableSsh = model.DisableSsh.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DisableSsh", "value": *sdk.DisableSsh})
	}

	// Handling Primitives
	if !model.DisableTelnet.IsNull() && !model.DisableTelnet.IsUnknown() {
		sdk.DisableTelnet = model.DisableTelnet.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DisableTelnet", "value": *sdk.DisableTelnet})
	}

	// Handling Primitives
	if !model.DisableUseridService.IsNull() && !model.DisableUseridService.IsUnknown() {
		sdk.DisableUseridService = model.DisableUseridService.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DisableUseridService", "value": *sdk.DisableUseridService})
	}

	// Handling Primitives
	if !model.DisableUseridSyslogListenerSsl.IsNull() && !model.DisableUseridSyslogListenerSsl.IsUnknown() {
		sdk.DisableUseridSyslogListenerSsl = model.DisableUseridSyslogListenerSsl.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DisableUseridSyslogListenerSsl", "value": *sdk.DisableUseridSyslogListenerSsl})
	}

	// Handling Primitives
	if !model.DisableUseridSyslogListenerUdp.IsNull() && !model.DisableUseridSyslogListenerUdp.IsUnknown() {
		sdk.DisableUseridSyslogListenerUdp = model.DisableUseridSyslogListenerUdp.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DisableUseridSyslogListenerUdp", "value": *sdk.DisableUseridSyslogListenerUdp})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ManagementInterfaceManagementInterfaceService", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ManagementInterfaceManagementInterfaceService ---
func packManagementInterfaceManagementInterfaceServiceFromSdk(ctx context.Context, sdk device_settings.ManagementInterfaceManagementInterfaceService) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ManagementInterfaceManagementInterfaceService", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ManagementInterfaceManagementInterfaceService
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.DisableHttp != nil {
		model.DisableHttp = basetypes.NewBoolValue(*sdk.DisableHttp)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DisableHttp", "value": *sdk.DisableHttp})
	} else {
		model.DisableHttp = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.DisableHttpOcsp != nil {
		model.DisableHttpOcsp = basetypes.NewBoolValue(*sdk.DisableHttpOcsp)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DisableHttpOcsp", "value": *sdk.DisableHttpOcsp})
	} else {
		model.DisableHttpOcsp = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.DisableHttps != nil {
		model.DisableHttps = basetypes.NewBoolValue(*sdk.DisableHttps)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DisableHttps", "value": *sdk.DisableHttps})
	} else {
		model.DisableHttps = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.DisableIcmp != nil {
		model.DisableIcmp = basetypes.NewBoolValue(*sdk.DisableIcmp)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DisableIcmp", "value": *sdk.DisableIcmp})
	} else {
		model.DisableIcmp = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.DisableSnmp != nil {
		model.DisableSnmp = basetypes.NewBoolValue(*sdk.DisableSnmp)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DisableSnmp", "value": *sdk.DisableSnmp})
	} else {
		model.DisableSnmp = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.DisableSsh != nil {
		model.DisableSsh = basetypes.NewBoolValue(*sdk.DisableSsh)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DisableSsh", "value": *sdk.DisableSsh})
	} else {
		model.DisableSsh = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.DisableTelnet != nil {
		model.DisableTelnet = basetypes.NewBoolValue(*sdk.DisableTelnet)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DisableTelnet", "value": *sdk.DisableTelnet})
	} else {
		model.DisableTelnet = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.DisableUseridService != nil {
		model.DisableUseridService = basetypes.NewBoolValue(*sdk.DisableUseridService)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DisableUseridService", "value": *sdk.DisableUseridService})
	} else {
		model.DisableUseridService = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.DisableUseridSyslogListenerSsl != nil {
		model.DisableUseridSyslogListenerSsl = basetypes.NewBoolValue(*sdk.DisableUseridSyslogListenerSsl)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DisableUseridSyslogListenerSsl", "value": *sdk.DisableUseridSyslogListenerSsl})
	} else {
		model.DisableUseridSyslogListenerSsl = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.DisableUseridSyslogListenerUdp != nil {
		model.DisableUseridSyslogListenerUdp = basetypes.NewBoolValue(*sdk.DisableUseridSyslogListenerUdp)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DisableUseridSyslogListenerUdp", "value": *sdk.DisableUseridSyslogListenerUdp})
	} else {
		model.DisableUseridSyslogListenerUdp = basetypes.NewBoolNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ManagementInterfaceManagementInterfaceService{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ManagementInterfaceManagementInterfaceService", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ManagementInterfaceManagementInterfaceService ---
func unpackManagementInterfaceManagementInterfaceServiceListToSdk(ctx context.Context, list types.List) ([]device_settings.ManagementInterfaceManagementInterfaceService, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ManagementInterfaceManagementInterfaceService")
	diags := diag.Diagnostics{}
	var data []models.ManagementInterfaceManagementInterfaceService
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]device_settings.ManagementInterfaceManagementInterfaceService, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ManagementInterfaceManagementInterfaceService{}.AttrTypes(), &item)
		unpacked, d := unpackManagementInterfaceManagementInterfaceServiceToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ManagementInterfaceManagementInterfaceService", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ManagementInterfaceManagementInterfaceService ---
func packManagementInterfaceManagementInterfaceServiceListFromSdk(ctx context.Context, sdks []device_settings.ManagementInterfaceManagementInterfaceService) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ManagementInterfaceManagementInterfaceService")
	diags := diag.Diagnostics{}
	var data []models.ManagementInterfaceManagementInterfaceService

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ManagementInterfaceManagementInterfaceService
		obj, d := packManagementInterfaceManagementInterfaceServiceFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ManagementInterfaceManagementInterfaceService{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ManagementInterfaceManagementInterfaceService", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ManagementInterfaceManagementInterfaceService{}.AttrType(), data)
}
