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

// --- Unpacker for AggregateEthernetInterfaces ---
func unpackAggregateEthernetInterfacesToSdk(ctx context.Context, obj types.Object) (*network_services.AggregateEthernetInterfaces, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.AggregateEthernetInterfaces", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.AggregateEthernetInterfaces
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.AggregateEthernetInterfaces
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

	// Handling Objects
	if !model.Layer2.IsNull() && !model.Layer2.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Layer2")
		unpacked, d := unpackAggregateEthernetInterfacesLayer2ToSdk(ctx, model.Layer2)
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
		unpacked, d := unpackAggregateEthernetInterfacesLayer3ToSdk(ctx, model.Layer3)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Layer3"})
		}
		if unpacked != nil {
			sdk.Layer3 = unpacked
		}
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

	tflog.Debug(ctx, "Exiting unpack helper for models.AggregateEthernetInterfaces", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for AggregateEthernetInterfaces ---
func packAggregateEthernetInterfacesFromSdk(ctx context.Context, sdk network_services.AggregateEthernetInterfaces) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.AggregateEthernetInterfaces", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.AggregateEthernetInterfaces
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
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Layer2 != nil {
		tflog.Debug(ctx, "Packing nested object for field Layer2")
		packed, d := packAggregateEthernetInterfacesLayer2FromSdk(ctx, *sdk.Layer2)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Layer2"})
		}
		model.Layer2 = packed
	} else {
		model.Layer2 = basetypes.NewObjectNull(models.AggregateEthernetInterfacesLayer2{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Layer3 != nil {
		tflog.Debug(ctx, "Packing nested object for field Layer3")
		packed, d := packAggregateEthernetInterfacesLayer3FromSdk(ctx, *sdk.Layer3)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Layer3"})
		}
		model.Layer3 = packed
	} else {
		model.Layer3 = basetypes.NewObjectNull(models.AggregateEthernetInterfacesLayer3{}.AttrTypes())
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

	obj, d := types.ObjectValueFrom(ctx, models.AggregateEthernetInterfaces{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.AggregateEthernetInterfaces", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for AggregateEthernetInterfaces ---
func unpackAggregateEthernetInterfacesListToSdk(ctx context.Context, list types.List) ([]network_services.AggregateEthernetInterfaces, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.AggregateEthernetInterfaces")
	diags := diag.Diagnostics{}
	var data []models.AggregateEthernetInterfaces
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.AggregateEthernetInterfaces, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.AggregateEthernetInterfaces{}.AttrTypes(), &item)
		unpacked, d := unpackAggregateEthernetInterfacesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.AggregateEthernetInterfaces", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for AggregateEthernetInterfaces ---
func packAggregateEthernetInterfacesListFromSdk(ctx context.Context, sdks []network_services.AggregateEthernetInterfaces) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.AggregateEthernetInterfaces")
	diags := diag.Diagnostics{}
	var data []models.AggregateEthernetInterfaces

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.AggregateEthernetInterfaces
		obj, d := packAggregateEthernetInterfacesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.AggregateEthernetInterfaces{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.AggregateEthernetInterfaces", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.AggregateEthernetInterfaces{}.AttrType(), data)
}

// --- Unpacker for AggregateEthernetInterfacesLayer2 ---
func unpackAggregateEthernetInterfacesLayer2ToSdk(ctx context.Context, obj types.Object) (*network_services.AggregateEthernetInterfacesLayer2, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.AggregateEthernetInterfacesLayer2", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.AggregateEthernetInterfacesLayer2
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.AggregateEthernetInterfacesLayer2
	var d diag.Diagnostics
	// Handling Objects
	if !model.Lacp.IsNull() && !model.Lacp.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Lacp")
		unpacked, d := unpackLacpToSdk(ctx, model.Lacp)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Lacp"})
		}
		if unpacked != nil {
			sdk.Lacp = unpacked
		}
	}

	// Handling Primitives
	if !model.VlanTag.IsNull() && !model.VlanTag.IsUnknown() {
		val := int32(model.VlanTag.ValueInt64())
		sdk.VlanTag = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "VlanTag", "value": *sdk.VlanTag})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.AggregateEthernetInterfacesLayer2", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for AggregateEthernetInterfacesLayer2 ---
func packAggregateEthernetInterfacesLayer2FromSdk(ctx context.Context, sdk network_services.AggregateEthernetInterfacesLayer2) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.AggregateEthernetInterfacesLayer2", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.AggregateEthernetInterfacesLayer2
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Lacp != nil {
		tflog.Debug(ctx, "Packing nested object for field Lacp")
		packed, d := packLacpFromSdk(ctx, *sdk.Lacp)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Lacp"})
		}
		model.Lacp = packed
	} else {
		model.Lacp = basetypes.NewObjectNull(models.Lacp{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.VlanTag != nil {
		model.VlanTag = basetypes.NewInt64Value(int64(*sdk.VlanTag))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "VlanTag", "value": *sdk.VlanTag})
	} else {
		model.VlanTag = basetypes.NewInt64Null()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.AggregateEthernetInterfacesLayer2{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.AggregateEthernetInterfacesLayer2", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for AggregateEthernetInterfacesLayer2 ---
func unpackAggregateEthernetInterfacesLayer2ListToSdk(ctx context.Context, list types.List) ([]network_services.AggregateEthernetInterfacesLayer2, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.AggregateEthernetInterfacesLayer2")
	diags := diag.Diagnostics{}
	var data []models.AggregateEthernetInterfacesLayer2
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.AggregateEthernetInterfacesLayer2, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.AggregateEthernetInterfacesLayer2{}.AttrTypes(), &item)
		unpacked, d := unpackAggregateEthernetInterfacesLayer2ToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.AggregateEthernetInterfacesLayer2", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for AggregateEthernetInterfacesLayer2 ---
func packAggregateEthernetInterfacesLayer2ListFromSdk(ctx context.Context, sdks []network_services.AggregateEthernetInterfacesLayer2) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.AggregateEthernetInterfacesLayer2")
	diags := diag.Diagnostics{}
	var data []models.AggregateEthernetInterfacesLayer2

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.AggregateEthernetInterfacesLayer2
		obj, d := packAggregateEthernetInterfacesLayer2FromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.AggregateEthernetInterfacesLayer2{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.AggregateEthernetInterfacesLayer2", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.AggregateEthernetInterfacesLayer2{}.AttrType(), data)
}

// --- Unpacker for Lacp ---
func unpackLacpToSdk(ctx context.Context, obj types.Object) (*network_services.Lacp, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.Lacp", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.Lacp
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.Lacp
	var d diag.Diagnostics
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
	if !model.MaxPorts.IsNull() && !model.MaxPorts.IsUnknown() {
		val := int32(model.MaxPorts.ValueInt64())
		sdk.MaxPorts = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "MaxPorts", "value": *sdk.MaxPorts})
	}

	// Handling Primitives
	if !model.Mode.IsNull() && !model.Mode.IsUnknown() {
		sdk.Mode = model.Mode.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Mode", "value": *sdk.Mode})
	}

	// Handling Primitives
	if !model.SystemPriority.IsNull() && !model.SystemPriority.IsUnknown() {
		val := int32(model.SystemPriority.ValueInt64())
		sdk.SystemPriority = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "SystemPriority", "value": *sdk.SystemPriority})
	}

	// Handling Primitives
	if !model.TransmissionRate.IsNull() && !model.TransmissionRate.IsUnknown() {
		sdk.TransmissionRate = model.TransmissionRate.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "TransmissionRate", "value": *sdk.TransmissionRate})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.Lacp", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for Lacp ---
func packLacpFromSdk(ctx context.Context, sdk network_services.Lacp) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.Lacp", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.Lacp
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
	if sdk.FastFailover != nil {
		model.FastFailover = basetypes.NewBoolValue(*sdk.FastFailover)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "FastFailover", "value": *sdk.FastFailover})
	} else {
		model.FastFailover = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.MaxPorts != nil {
		model.MaxPorts = basetypes.NewInt64Value(int64(*sdk.MaxPorts))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "MaxPorts", "value": *sdk.MaxPorts})
	} else {
		model.MaxPorts = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Mode != nil {
		model.Mode = basetypes.NewStringValue(*sdk.Mode)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Mode", "value": *sdk.Mode})
	} else {
		model.Mode = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.SystemPriority != nil {
		model.SystemPriority = basetypes.NewInt64Value(int64(*sdk.SystemPriority))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "SystemPriority", "value": *sdk.SystemPriority})
	} else {
		model.SystemPriority = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.TransmissionRate != nil {
		model.TransmissionRate = basetypes.NewStringValue(*sdk.TransmissionRate)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "TransmissionRate", "value": *sdk.TransmissionRate})
	} else {
		model.TransmissionRate = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.Lacp{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.Lacp", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for Lacp ---
func unpackLacpListToSdk(ctx context.Context, list types.List) ([]network_services.Lacp, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.Lacp")
	diags := diag.Diagnostics{}
	var data []models.Lacp
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.Lacp, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.Lacp{}.AttrTypes(), &item)
		unpacked, d := unpackLacpToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.Lacp", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for Lacp ---
func packLacpListFromSdk(ctx context.Context, sdks []network_services.Lacp) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.Lacp")
	diags := diag.Diagnostics{}
	var data []models.Lacp

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.Lacp
		obj, d := packLacpFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.Lacp{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.Lacp", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.Lacp{}.AttrType(), data)
}

// --- Unpacker for AggregateEthernetInterfacesLayer3 ---
func unpackAggregateEthernetInterfacesLayer3ToSdk(ctx context.Context, obj types.Object) (*network_services.AggregateEthernetInterfacesLayer3, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.AggregateEthernetInterfacesLayer3", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.AggregateEthernetInterfacesLayer3
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.AggregateEthernetInterfacesLayer3
	var d diag.Diagnostics
	// Handling Lists
	if !model.Arp.IsNull() && !model.Arp.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field Arp")
		unpacked, d := unpackAggEthernetArpInnerListToSdk(ctx, model.Arp)
		diags.Append(d...)
		sdk.Arp = unpacked
	}

	// Handling Objects
	if !model.DdnsConfig.IsNull() && !model.DdnsConfig.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field DdnsConfig")
		unpacked, d := unpackAggregateEthernetInterfacesLayer3DdnsConfigToSdk(ctx, model.DdnsConfig)
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
		unpacked, d := unpackAggEthernetDhcpClientDhcpClientToSdk(ctx, model.DhcpClient)
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

	// Handling Objects
	if !model.Lacp.IsNull() && !model.Lacp.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Lacp")
		unpacked, d := unpackLacpToSdk(ctx, model.Lacp)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Lacp"})
		}
		if unpacked != nil {
			sdk.Lacp = unpacked
		}
	}

	// Handling Primitives
	if !model.Mtu.IsNull() && !model.Mtu.IsUnknown() {
		val := int32(model.Mtu.ValueInt64())
		sdk.Mtu = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Mtu", "value": *sdk.Mtu})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.AggregateEthernetInterfacesLayer3", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for AggregateEthernetInterfacesLayer3 ---
func packAggregateEthernetInterfacesLayer3FromSdk(ctx context.Context, sdk network_services.AggregateEthernetInterfacesLayer3) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.AggregateEthernetInterfacesLayer3", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.AggregateEthernetInterfacesLayer3
	var d diag.Diagnostics
	// Handling Lists
	if sdk.Arp != nil {
		tflog.Debug(ctx, "Packing list of objects for field Arp")
		packed, d := packAggEthernetArpInnerListFromSdk(ctx, sdk.Arp)
		diags.Append(d...)
		model.Arp = packed
	} else {
		model.Arp = basetypes.NewListNull(models.AggEthernetArpInner{}.AttrType())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.DdnsConfig != nil {
		tflog.Debug(ctx, "Packing nested object for field DdnsConfig")
		packed, d := packAggregateEthernetInterfacesLayer3DdnsConfigFromSdk(ctx, *sdk.DdnsConfig)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "DdnsConfig"})
		}
		model.DdnsConfig = packed
	} else {
		model.DdnsConfig = basetypes.NewObjectNull(models.AggregateEthernetInterfacesLayer3DdnsConfig{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.DhcpClient != nil {
		tflog.Debug(ctx, "Packing nested object for field DhcpClient")
		packed, d := packAggEthernetDhcpClientDhcpClientFromSdk(ctx, *sdk.DhcpClient)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "DhcpClient"})
		}
		model.DhcpClient = packed
	} else {
		model.DhcpClient = basetypes.NewObjectNull(models.AggEthernetDhcpClientDhcpClient{}.AttrTypes())
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
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Lacp != nil {
		tflog.Debug(ctx, "Packing nested object for field Lacp")
		packed, d := packLacpFromSdk(ctx, *sdk.Lacp)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Lacp"})
		}
		model.Lacp = packed
	} else {
		model.Lacp = basetypes.NewObjectNull(models.Lacp{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Mtu != nil {
		model.Mtu = basetypes.NewInt64Value(int64(*sdk.Mtu))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Mtu", "value": *sdk.Mtu})
	} else {
		model.Mtu = basetypes.NewInt64Null()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.AggregateEthernetInterfacesLayer3{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.AggregateEthernetInterfacesLayer3", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for AggregateEthernetInterfacesLayer3 ---
func unpackAggregateEthernetInterfacesLayer3ListToSdk(ctx context.Context, list types.List) ([]network_services.AggregateEthernetInterfacesLayer3, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.AggregateEthernetInterfacesLayer3")
	diags := diag.Diagnostics{}
	var data []models.AggregateEthernetInterfacesLayer3
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.AggregateEthernetInterfacesLayer3, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.AggregateEthernetInterfacesLayer3{}.AttrTypes(), &item)
		unpacked, d := unpackAggregateEthernetInterfacesLayer3ToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.AggregateEthernetInterfacesLayer3", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for AggregateEthernetInterfacesLayer3 ---
func packAggregateEthernetInterfacesLayer3ListFromSdk(ctx context.Context, sdks []network_services.AggregateEthernetInterfacesLayer3) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.AggregateEthernetInterfacesLayer3")
	diags := diag.Diagnostics{}
	var data []models.AggregateEthernetInterfacesLayer3

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.AggregateEthernetInterfacesLayer3
		obj, d := packAggregateEthernetInterfacesLayer3FromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.AggregateEthernetInterfacesLayer3{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.AggregateEthernetInterfacesLayer3", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.AggregateEthernetInterfacesLayer3{}.AttrType(), data)
}

// --- Unpacker for AggEthernetArpInner ---
func unpackAggEthernetArpInnerToSdk(ctx context.Context, obj types.Object) (*network_services.AggEthernetArpInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.AggEthernetArpInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.AggEthernetArpInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.AggEthernetArpInner
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

	tflog.Debug(ctx, "Exiting unpack helper for models.AggEthernetArpInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for AggEthernetArpInner ---
func packAggEthernetArpInnerFromSdk(ctx context.Context, sdk network_services.AggEthernetArpInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.AggEthernetArpInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.AggEthernetArpInner
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

	obj, d := types.ObjectValueFrom(ctx, models.AggEthernetArpInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.AggEthernetArpInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for AggEthernetArpInner ---
func unpackAggEthernetArpInnerListToSdk(ctx context.Context, list types.List) ([]network_services.AggEthernetArpInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.AggEthernetArpInner")
	diags := diag.Diagnostics{}
	var data []models.AggEthernetArpInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.AggEthernetArpInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.AggEthernetArpInner{}.AttrTypes(), &item)
		unpacked, d := unpackAggEthernetArpInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.AggEthernetArpInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for AggEthernetArpInner ---
func packAggEthernetArpInnerListFromSdk(ctx context.Context, sdks []network_services.AggEthernetArpInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.AggEthernetArpInner")
	diags := diag.Diagnostics{}
	var data []models.AggEthernetArpInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.AggEthernetArpInner
		obj, d := packAggEthernetArpInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.AggEthernetArpInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.AggEthernetArpInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.AggEthernetArpInner{}.AttrType(), data)
}

// --- Unpacker for AggregateEthernetInterfacesLayer3DdnsConfig ---
func unpackAggregateEthernetInterfacesLayer3DdnsConfigToSdk(ctx context.Context, obj types.Object) (*network_services.AggregateEthernetInterfacesLayer3DdnsConfig, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.AggregateEthernetInterfacesLayer3DdnsConfig", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.AggregateEthernetInterfacesLayer3DdnsConfig
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.AggregateEthernetInterfacesLayer3DdnsConfig
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

	tflog.Debug(ctx, "Exiting unpack helper for models.AggregateEthernetInterfacesLayer3DdnsConfig", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for AggregateEthernetInterfacesLayer3DdnsConfig ---
func packAggregateEthernetInterfacesLayer3DdnsConfigFromSdk(ctx context.Context, sdk network_services.AggregateEthernetInterfacesLayer3DdnsConfig) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.AggregateEthernetInterfacesLayer3DdnsConfig", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.AggregateEthernetInterfacesLayer3DdnsConfig
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

	obj, d := types.ObjectValueFrom(ctx, models.AggregateEthernetInterfacesLayer3DdnsConfig{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.AggregateEthernetInterfacesLayer3DdnsConfig", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for AggregateEthernetInterfacesLayer3DdnsConfig ---
func unpackAggregateEthernetInterfacesLayer3DdnsConfigListToSdk(ctx context.Context, list types.List) ([]network_services.AggregateEthernetInterfacesLayer3DdnsConfig, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.AggregateEthernetInterfacesLayer3DdnsConfig")
	diags := diag.Diagnostics{}
	var data []models.AggregateEthernetInterfacesLayer3DdnsConfig
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.AggregateEthernetInterfacesLayer3DdnsConfig, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.AggregateEthernetInterfacesLayer3DdnsConfig{}.AttrTypes(), &item)
		unpacked, d := unpackAggregateEthernetInterfacesLayer3DdnsConfigToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.AggregateEthernetInterfacesLayer3DdnsConfig", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for AggregateEthernetInterfacesLayer3DdnsConfig ---
func packAggregateEthernetInterfacesLayer3DdnsConfigListFromSdk(ctx context.Context, sdks []network_services.AggregateEthernetInterfacesLayer3DdnsConfig) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.AggregateEthernetInterfacesLayer3DdnsConfig")
	diags := diag.Diagnostics{}
	var data []models.AggregateEthernetInterfacesLayer3DdnsConfig

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.AggregateEthernetInterfacesLayer3DdnsConfig
		obj, d := packAggregateEthernetInterfacesLayer3DdnsConfigFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.AggregateEthernetInterfacesLayer3DdnsConfig{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.AggregateEthernetInterfacesLayer3DdnsConfig", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.AggregateEthernetInterfacesLayer3DdnsConfig{}.AttrType(), data)
}

// --- Unpacker for AggEthernetDhcpClientDhcpClient ---
func unpackAggEthernetDhcpClientDhcpClientToSdk(ctx context.Context, obj types.Object) (*network_services.AggEthernetDhcpClientDhcpClient, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.AggEthernetDhcpClientDhcpClient", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.AggEthernetDhcpClientDhcpClient
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.AggEthernetDhcpClientDhcpClient
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
		unpacked, d := unpackAggEthernetDhcpClientDhcpClientSendHostnameToSdk(ctx, model.SendHostname)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "SendHostname"})
		}
		if unpacked != nil {
			sdk.SendHostname = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.AggEthernetDhcpClientDhcpClient", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for AggEthernetDhcpClientDhcpClient ---
func packAggEthernetDhcpClientDhcpClientFromSdk(ctx context.Context, sdk network_services.AggEthernetDhcpClientDhcpClient) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.AggEthernetDhcpClientDhcpClient", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.AggEthernetDhcpClientDhcpClient
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
		packed, d := packAggEthernetDhcpClientDhcpClientSendHostnameFromSdk(ctx, *sdk.SendHostname)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "SendHostname"})
		}
		model.SendHostname = packed
	} else {
		model.SendHostname = basetypes.NewObjectNull(models.AggEthernetDhcpClientDhcpClientSendHostname{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.AggEthernetDhcpClientDhcpClient{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.AggEthernetDhcpClientDhcpClient", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for AggEthernetDhcpClientDhcpClient ---
func unpackAggEthernetDhcpClientDhcpClientListToSdk(ctx context.Context, list types.List) ([]network_services.AggEthernetDhcpClientDhcpClient, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.AggEthernetDhcpClientDhcpClient")
	diags := diag.Diagnostics{}
	var data []models.AggEthernetDhcpClientDhcpClient
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.AggEthernetDhcpClientDhcpClient, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.AggEthernetDhcpClientDhcpClient{}.AttrTypes(), &item)
		unpacked, d := unpackAggEthernetDhcpClientDhcpClientToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.AggEthernetDhcpClientDhcpClient", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for AggEthernetDhcpClientDhcpClient ---
func packAggEthernetDhcpClientDhcpClientListFromSdk(ctx context.Context, sdks []network_services.AggEthernetDhcpClientDhcpClient) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.AggEthernetDhcpClientDhcpClient")
	diags := diag.Diagnostics{}
	var data []models.AggEthernetDhcpClientDhcpClient

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.AggEthernetDhcpClientDhcpClient
		obj, d := packAggEthernetDhcpClientDhcpClientFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.AggEthernetDhcpClientDhcpClient{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.AggEthernetDhcpClientDhcpClient", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.AggEthernetDhcpClientDhcpClient{}.AttrType(), data)
}

// --- Unpacker for AggEthernetDhcpClientDhcpClientSendHostname ---
func unpackAggEthernetDhcpClientDhcpClientSendHostnameToSdk(ctx context.Context, obj types.Object) (*network_services.AggEthernetDhcpClientDhcpClientSendHostname, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.AggEthernetDhcpClientDhcpClientSendHostname", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.AggEthernetDhcpClientDhcpClientSendHostname
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.AggEthernetDhcpClientDhcpClientSendHostname
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

	tflog.Debug(ctx, "Exiting unpack helper for models.AggEthernetDhcpClientDhcpClientSendHostname", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for AggEthernetDhcpClientDhcpClientSendHostname ---
func packAggEthernetDhcpClientDhcpClientSendHostnameFromSdk(ctx context.Context, sdk network_services.AggEthernetDhcpClientDhcpClientSendHostname) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.AggEthernetDhcpClientDhcpClientSendHostname", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.AggEthernetDhcpClientDhcpClientSendHostname
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

	obj, d := types.ObjectValueFrom(ctx, models.AggEthernetDhcpClientDhcpClientSendHostname{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.AggEthernetDhcpClientDhcpClientSendHostname", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for AggEthernetDhcpClientDhcpClientSendHostname ---
func unpackAggEthernetDhcpClientDhcpClientSendHostnameListToSdk(ctx context.Context, list types.List) ([]network_services.AggEthernetDhcpClientDhcpClientSendHostname, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.AggEthernetDhcpClientDhcpClientSendHostname")
	diags := diag.Diagnostics{}
	var data []models.AggEthernetDhcpClientDhcpClientSendHostname
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.AggEthernetDhcpClientDhcpClientSendHostname, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.AggEthernetDhcpClientDhcpClientSendHostname{}.AttrTypes(), &item)
		unpacked, d := unpackAggEthernetDhcpClientDhcpClientSendHostnameToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.AggEthernetDhcpClientDhcpClientSendHostname", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for AggEthernetDhcpClientDhcpClientSendHostname ---
func packAggEthernetDhcpClientDhcpClientSendHostnameListFromSdk(ctx context.Context, sdks []network_services.AggEthernetDhcpClientDhcpClientSendHostname) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.AggEthernetDhcpClientDhcpClientSendHostname")
	diags := diag.Diagnostics{}
	var data []models.AggEthernetDhcpClientDhcpClientSendHostname

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.AggEthernetDhcpClientDhcpClientSendHostname
		obj, d := packAggEthernetDhcpClientDhcpClientSendHostnameFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.AggEthernetDhcpClientDhcpClientSendHostname{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.AggEthernetDhcpClientDhcpClientSendHostname", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.AggEthernetDhcpClientDhcpClientSendHostname{}.AttrType(), data)
}
