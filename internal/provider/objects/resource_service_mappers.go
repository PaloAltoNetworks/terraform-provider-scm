package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/objects"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/objects"
)

// --- Unpacker for Services ---
func unpackServicesToSdk(ctx context.Context, obj types.Object) (*objects.Services, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.Services", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.Services
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.Services
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Description.IsNull() && !model.Description.IsUnknown() {
		sdk.Description = model.Description.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Description", "value": *sdk.Description})
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
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	// Handling Objects
	if !model.Protocol.IsNull() && !model.Protocol.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Protocol")
		unpacked, d := unpackServicesProtocolToSdk(ctx, model.Protocol)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Protocol"})
		}
		if unpacked != nil {
			sdk.Protocol = unpacked
		}
	}

	// Handling Primitives
	if !model.Snippet.IsNull() && !model.Snippet.IsUnknown() {
		sdk.Snippet = model.Snippet.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	}

	// Handling Lists
	if !model.Tag.IsNull() && !model.Tag.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Tag")
		diags.Append(model.Tag.ElementsAs(ctx, &sdk.Tag, false)...)
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.Services", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for Services ---
func packServicesFromSdk(ctx context.Context, sdk objects.Services) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.Services", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.Services
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
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Protocol != nil {
		tflog.Debug(ctx, "Packing nested object for field Protocol")
		packed, d := packServicesProtocolFromSdk(ctx, *sdk.Protocol)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Protocol"})
		}
		model.Protocol = packed
	} else {
		model.Protocol = basetypes.NewObjectNull(models.ServicesProtocol{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Snippet != nil {
		model.Snippet = basetypes.NewStringValue(*sdk.Snippet)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	} else {
		model.Snippet = basetypes.NewStringNull()
	}
	// Handling Lists
	if sdk.Tag != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Tag")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Tag, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Tag)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Tag = basetypes.NewListNull(elemType)
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.Services{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.Services", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for Services ---
func unpackServicesListToSdk(ctx context.Context, list types.List) ([]objects.Services, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.Services")
	diags := diag.Diagnostics{}
	var data []models.Services
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.Services, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.Services{}.AttrTypes(), &item)
		unpacked, d := unpackServicesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.Services", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for Services ---
func packServicesListFromSdk(ctx context.Context, sdks []objects.Services) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.Services")
	diags := diag.Diagnostics{}
	var data []models.Services

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.Services
		obj, d := packServicesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.Services{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.Services", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.Services{}.AttrType(), data)
}

// --- Unpacker for ServicesProtocol ---
func unpackServicesProtocolToSdk(ctx context.Context, obj types.Object) (*objects.ServicesProtocol, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ServicesProtocol", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ServicesProtocol
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.ServicesProtocol
	var d diag.Diagnostics
	// Handling Objects
	if !model.Tcp.IsNull() && !model.Tcp.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Tcp")
		unpacked, d := unpackServicesProtocolTcpToSdk(ctx, model.Tcp)
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
		unpacked, d := unpackServicesProtocolUdpToSdk(ctx, model.Udp)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Udp"})
		}
		if unpacked != nil {
			sdk.Udp = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ServicesProtocol", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ServicesProtocol ---
func packServicesProtocolFromSdk(ctx context.Context, sdk objects.ServicesProtocol) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ServicesProtocol", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ServicesProtocol
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Tcp != nil {
		tflog.Debug(ctx, "Packing nested object for field Tcp")
		packed, d := packServicesProtocolTcpFromSdk(ctx, *sdk.Tcp)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Tcp"})
		}
		model.Tcp = packed
	} else {
		model.Tcp = basetypes.NewObjectNull(models.ServicesProtocolTcp{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Udp != nil {
		tflog.Debug(ctx, "Packing nested object for field Udp")
		packed, d := packServicesProtocolUdpFromSdk(ctx, *sdk.Udp)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Udp"})
		}
		model.Udp = packed
	} else {
		model.Udp = basetypes.NewObjectNull(models.ServicesProtocolUdp{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ServicesProtocol{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ServicesProtocol", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ServicesProtocol ---
func unpackServicesProtocolListToSdk(ctx context.Context, list types.List) ([]objects.ServicesProtocol, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ServicesProtocol")
	diags := diag.Diagnostics{}
	var data []models.ServicesProtocol
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.ServicesProtocol, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ServicesProtocol{}.AttrTypes(), &item)
		unpacked, d := unpackServicesProtocolToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ServicesProtocol", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ServicesProtocol ---
func packServicesProtocolListFromSdk(ctx context.Context, sdks []objects.ServicesProtocol) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ServicesProtocol")
	diags := diag.Diagnostics{}
	var data []models.ServicesProtocol

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ServicesProtocol
		obj, d := packServicesProtocolFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ServicesProtocol{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ServicesProtocol", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ServicesProtocol{}.AttrType(), data)
}

// --- Unpacker for ServicesProtocolTcp ---
func unpackServicesProtocolTcpToSdk(ctx context.Context, obj types.Object) (*objects.ServicesProtocolTcp, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ServicesProtocolTcp", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ServicesProtocolTcp
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.ServicesProtocolTcp
	var d diag.Diagnostics
	// Handling Objects
	if !model.Override.IsNull() && !model.Override.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Override")
		unpacked, d := unpackServicesProtocolTcpOverrideToSdk(ctx, model.Override)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Override"})
		}
		if unpacked != nil {
			sdk.Override = unpacked
		}
	}

	// Handling Primitives
	if !model.Port.IsNull() && !model.Port.IsUnknown() {
		sdk.Port = model.Port.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Port", "value": sdk.Port})
	}

	// Handling Primitives
	if !model.SourcePort.IsNull() && !model.SourcePort.IsUnknown() {
		sdk.SourcePort = model.SourcePort.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "SourcePort", "value": *sdk.SourcePort})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ServicesProtocolTcp", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ServicesProtocolTcp ---
func packServicesProtocolTcpFromSdk(ctx context.Context, sdk objects.ServicesProtocolTcp) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ServicesProtocolTcp", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ServicesProtocolTcp
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Override != nil {
		tflog.Debug(ctx, "Packing nested object for field Override")
		packed, d := packServicesProtocolTcpOverrideFromSdk(ctx, *sdk.Override)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Override"})
		}
		model.Override = packed
	} else {
		model.Override = basetypes.NewObjectNull(models.ServicesProtocolTcpOverride{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	model.Port = basetypes.NewStringValue(sdk.Port)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Port", "value": sdk.Port})
	// Handling Primitives
	// Standard primitive packing
	if sdk.SourcePort != nil {
		model.SourcePort = basetypes.NewStringValue(*sdk.SourcePort)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "SourcePort", "value": *sdk.SourcePort})
	} else {
		model.SourcePort = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ServicesProtocolTcp{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ServicesProtocolTcp", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ServicesProtocolTcp ---
func unpackServicesProtocolTcpListToSdk(ctx context.Context, list types.List) ([]objects.ServicesProtocolTcp, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ServicesProtocolTcp")
	diags := diag.Diagnostics{}
	var data []models.ServicesProtocolTcp
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.ServicesProtocolTcp, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ServicesProtocolTcp{}.AttrTypes(), &item)
		unpacked, d := unpackServicesProtocolTcpToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ServicesProtocolTcp", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ServicesProtocolTcp ---
func packServicesProtocolTcpListFromSdk(ctx context.Context, sdks []objects.ServicesProtocolTcp) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ServicesProtocolTcp")
	diags := diag.Diagnostics{}
	var data []models.ServicesProtocolTcp

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ServicesProtocolTcp
		obj, d := packServicesProtocolTcpFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ServicesProtocolTcp{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ServicesProtocolTcp", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ServicesProtocolTcp{}.AttrType(), data)
}

// --- Unpacker for ServicesProtocolTcpOverride ---
func unpackServicesProtocolTcpOverrideToSdk(ctx context.Context, obj types.Object) (*objects.ServicesProtocolTcpOverride, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ServicesProtocolTcpOverride", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ServicesProtocolTcpOverride
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.ServicesProtocolTcpOverride
	var d diag.Diagnostics
	// Handling Primitives
	if !model.HalfcloseTimeout.IsNull() && !model.HalfcloseTimeout.IsUnknown() {
		val := int32(model.HalfcloseTimeout.ValueInt64())
		sdk.HalfcloseTimeout = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "HalfcloseTimeout", "value": *sdk.HalfcloseTimeout})
	}

	// Handling Primitives
	if !model.Timeout.IsNull() && !model.Timeout.IsUnknown() {
		val := int32(model.Timeout.ValueInt64())
		sdk.Timeout = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Timeout", "value": *sdk.Timeout})
	}

	// Handling Primitives
	if !model.TimewaitTimeout.IsNull() && !model.TimewaitTimeout.IsUnknown() {
		val := int32(model.TimewaitTimeout.ValueInt64())
		sdk.TimewaitTimeout = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "TimewaitTimeout", "value": *sdk.TimewaitTimeout})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ServicesProtocolTcpOverride", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ServicesProtocolTcpOverride ---
func packServicesProtocolTcpOverrideFromSdk(ctx context.Context, sdk objects.ServicesProtocolTcpOverride) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ServicesProtocolTcpOverride", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ServicesProtocolTcpOverride
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.HalfcloseTimeout != nil {
		model.HalfcloseTimeout = basetypes.NewInt64Value(int64(*sdk.HalfcloseTimeout))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "HalfcloseTimeout", "value": *sdk.HalfcloseTimeout})
	} else {
		model.HalfcloseTimeout = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Timeout != nil {
		model.Timeout = basetypes.NewInt64Value(int64(*sdk.Timeout))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Timeout", "value": *sdk.Timeout})
	} else {
		model.Timeout = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.TimewaitTimeout != nil {
		model.TimewaitTimeout = basetypes.NewInt64Value(int64(*sdk.TimewaitTimeout))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "TimewaitTimeout", "value": *sdk.TimewaitTimeout})
	} else {
		model.TimewaitTimeout = basetypes.NewInt64Null()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ServicesProtocolTcpOverride{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ServicesProtocolTcpOverride", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ServicesProtocolTcpOverride ---
func unpackServicesProtocolTcpOverrideListToSdk(ctx context.Context, list types.List) ([]objects.ServicesProtocolTcpOverride, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ServicesProtocolTcpOverride")
	diags := diag.Diagnostics{}
	var data []models.ServicesProtocolTcpOverride
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.ServicesProtocolTcpOverride, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ServicesProtocolTcpOverride{}.AttrTypes(), &item)
		unpacked, d := unpackServicesProtocolTcpOverrideToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ServicesProtocolTcpOverride", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ServicesProtocolTcpOverride ---
func packServicesProtocolTcpOverrideListFromSdk(ctx context.Context, sdks []objects.ServicesProtocolTcpOverride) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ServicesProtocolTcpOverride")
	diags := diag.Diagnostics{}
	var data []models.ServicesProtocolTcpOverride

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ServicesProtocolTcpOverride
		obj, d := packServicesProtocolTcpOverrideFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ServicesProtocolTcpOverride{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ServicesProtocolTcpOverride", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ServicesProtocolTcpOverride{}.AttrType(), data)
}

// --- Unpacker for ServicesProtocolUdp ---
func unpackServicesProtocolUdpToSdk(ctx context.Context, obj types.Object) (*objects.ServicesProtocolUdp, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ServicesProtocolUdp", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ServicesProtocolUdp
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.ServicesProtocolUdp
	var d diag.Diagnostics
	// Handling Objects
	if !model.Override.IsNull() && !model.Override.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Override")
		unpacked, d := unpackServicesProtocolUdpOverrideToSdk(ctx, model.Override)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Override"})
		}
		if unpacked != nil {
			sdk.Override = unpacked
		}
	}

	// Handling Primitives
	if !model.Port.IsNull() && !model.Port.IsUnknown() {
		sdk.Port = model.Port.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Port", "value": sdk.Port})
	}

	// Handling Primitives
	if !model.SourcePort.IsNull() && !model.SourcePort.IsUnknown() {
		sdk.SourcePort = model.SourcePort.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "SourcePort", "value": *sdk.SourcePort})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ServicesProtocolUdp", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ServicesProtocolUdp ---
func packServicesProtocolUdpFromSdk(ctx context.Context, sdk objects.ServicesProtocolUdp) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ServicesProtocolUdp", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ServicesProtocolUdp
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Override != nil {
		tflog.Debug(ctx, "Packing nested object for field Override")
		packed, d := packServicesProtocolUdpOverrideFromSdk(ctx, *sdk.Override)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Override"})
		}
		model.Override = packed
	} else {
		model.Override = basetypes.NewObjectNull(models.ServicesProtocolUdpOverride{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	model.Port = basetypes.NewStringValue(sdk.Port)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Port", "value": sdk.Port})
	// Handling Primitives
	// Standard primitive packing
	if sdk.SourcePort != nil {
		model.SourcePort = basetypes.NewStringValue(*sdk.SourcePort)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "SourcePort", "value": *sdk.SourcePort})
	} else {
		model.SourcePort = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ServicesProtocolUdp{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ServicesProtocolUdp", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ServicesProtocolUdp ---
func unpackServicesProtocolUdpListToSdk(ctx context.Context, list types.List) ([]objects.ServicesProtocolUdp, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ServicesProtocolUdp")
	diags := diag.Diagnostics{}
	var data []models.ServicesProtocolUdp
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.ServicesProtocolUdp, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ServicesProtocolUdp{}.AttrTypes(), &item)
		unpacked, d := unpackServicesProtocolUdpToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ServicesProtocolUdp", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ServicesProtocolUdp ---
func packServicesProtocolUdpListFromSdk(ctx context.Context, sdks []objects.ServicesProtocolUdp) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ServicesProtocolUdp")
	diags := diag.Diagnostics{}
	var data []models.ServicesProtocolUdp

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ServicesProtocolUdp
		obj, d := packServicesProtocolUdpFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ServicesProtocolUdp{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ServicesProtocolUdp", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ServicesProtocolUdp{}.AttrType(), data)
}

// --- Unpacker for ServicesProtocolUdpOverride ---
func unpackServicesProtocolUdpOverrideToSdk(ctx context.Context, obj types.Object) (*objects.ServicesProtocolUdpOverride, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ServicesProtocolUdpOverride", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ServicesProtocolUdpOverride
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.ServicesProtocolUdpOverride
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Timeout.IsNull() && !model.Timeout.IsUnknown() {
		val := int32(model.Timeout.ValueInt64())
		sdk.Timeout = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Timeout", "value": *sdk.Timeout})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ServicesProtocolUdpOverride", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ServicesProtocolUdpOverride ---
func packServicesProtocolUdpOverrideFromSdk(ctx context.Context, sdk objects.ServicesProtocolUdpOverride) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ServicesProtocolUdpOverride", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ServicesProtocolUdpOverride
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Timeout != nil {
		model.Timeout = basetypes.NewInt64Value(int64(*sdk.Timeout))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Timeout", "value": *sdk.Timeout})
	} else {
		model.Timeout = basetypes.NewInt64Null()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ServicesProtocolUdpOverride{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ServicesProtocolUdpOverride", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ServicesProtocolUdpOverride ---
func unpackServicesProtocolUdpOverrideListToSdk(ctx context.Context, list types.List) ([]objects.ServicesProtocolUdpOverride, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ServicesProtocolUdpOverride")
	diags := diag.Diagnostics{}
	var data []models.ServicesProtocolUdpOverride
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.ServicesProtocolUdpOverride, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ServicesProtocolUdpOverride{}.AttrTypes(), &item)
		unpacked, d := unpackServicesProtocolUdpOverrideToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ServicesProtocolUdpOverride", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ServicesProtocolUdpOverride ---
func packServicesProtocolUdpOverrideListFromSdk(ctx context.Context, sdks []objects.ServicesProtocolUdpOverride) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ServicesProtocolUdpOverride")
	diags := diag.Diagnostics{}
	var data []models.ServicesProtocolUdpOverride

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ServicesProtocolUdpOverride
		obj, d := packServicesProtocolUdpOverrideFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ServicesProtocolUdpOverride{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ServicesProtocolUdpOverride", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ServicesProtocolUdpOverride{}.AttrType(), data)
}
