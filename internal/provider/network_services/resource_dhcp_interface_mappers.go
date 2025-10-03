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

// --- Unpacker for DhcpInterfaces ---
func unpackDhcpInterfacesToSdk(ctx context.Context, obj types.Object) (*network_services.DhcpInterfaces, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.DhcpInterfaces", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.DhcpInterfaces
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.DhcpInterfaces
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

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	// Handling Objects
	if !model.Relay.IsNull() && !model.Relay.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Relay")
		unpacked, d := unpackDhcpInterfacesRelayToSdk(ctx, model.Relay)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Relay"})
		}
		if unpacked != nil {
			sdk.Relay = unpacked
		}
	}

	// Handling Objects
	if !model.Server.IsNull() && !model.Server.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Server")
		unpacked, d := unpackDhcpInterfacesServerToSdk(ctx, model.Server)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Server"})
		}
		if unpacked != nil {
			sdk.Server = unpacked
		}
	}

	// Handling Primitives
	if !model.Snippet.IsNull() && !model.Snippet.IsUnknown() {
		sdk.Snippet = model.Snippet.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.DhcpInterfaces", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for DhcpInterfaces ---
func packDhcpInterfacesFromSdk(ctx context.Context, sdk network_services.DhcpInterfaces) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.DhcpInterfaces", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.DhcpInterfaces
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
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Relay != nil {
		tflog.Debug(ctx, "Packing nested object for field Relay")
		packed, d := packDhcpInterfacesRelayFromSdk(ctx, *sdk.Relay)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Relay"})
		}
		model.Relay = packed
	} else {
		model.Relay = basetypes.NewObjectNull(models.DhcpInterfacesRelay{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Server != nil {
		tflog.Debug(ctx, "Packing nested object for field Server")
		packed, d := packDhcpInterfacesServerFromSdk(ctx, *sdk.Server)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Server"})
		}
		model.Server = packed
	} else {
		model.Server = basetypes.NewObjectNull(models.DhcpInterfacesServer{}.AttrTypes())
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

	obj, d := types.ObjectValueFrom(ctx, models.DhcpInterfaces{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.DhcpInterfaces", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for DhcpInterfaces ---
func unpackDhcpInterfacesListToSdk(ctx context.Context, list types.List) ([]network_services.DhcpInterfaces, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.DhcpInterfaces")
	diags := diag.Diagnostics{}
	var data []models.DhcpInterfaces
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.DhcpInterfaces, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.DhcpInterfaces{}.AttrTypes(), &item)
		unpacked, d := unpackDhcpInterfacesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.DhcpInterfaces", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for DhcpInterfaces ---
func packDhcpInterfacesListFromSdk(ctx context.Context, sdks []network_services.DhcpInterfaces) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.DhcpInterfaces")
	diags := diag.Diagnostics{}
	var data []models.DhcpInterfaces

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.DhcpInterfaces
		obj, d := packDhcpInterfacesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.DhcpInterfaces{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.DhcpInterfaces", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.DhcpInterfaces{}.AttrType(), data)
}

// --- Unpacker for DhcpInterfacesRelay ---
func unpackDhcpInterfacesRelayToSdk(ctx context.Context, obj types.Object) (*network_services.DhcpInterfacesRelay, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.DhcpInterfacesRelay", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.DhcpInterfacesRelay
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.DhcpInterfacesRelay
	var d diag.Diagnostics
	// Handling Objects
	if !model.Ip.IsNull() && !model.Ip.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Ip")
		unpacked, d := unpackDhcpInterfacesRelayIpToSdk(ctx, model.Ip)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Ip"})
		}
		if unpacked != nil {
			sdk.Ip = *unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.DhcpInterfacesRelay", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for DhcpInterfacesRelay ---
func packDhcpInterfacesRelayFromSdk(ctx context.Context, sdk network_services.DhcpInterfacesRelay) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.DhcpInterfacesRelay", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.DhcpInterfacesRelay
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	// Logic for non-pointer / value-type nested objects
	if !reflect.ValueOf(sdk.Ip).IsZero() {
		tflog.Debug(ctx, "Packing nested object for field Ip")
		packed, d := packDhcpInterfacesRelayIpFromSdk(ctx, sdk.Ip)
		diags.Append(d...)
		model.Ip = packed
	} else {
		model.Ip = basetypes.NewObjectNull(models.DhcpInterfacesRelayIp{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.DhcpInterfacesRelay{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.DhcpInterfacesRelay", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for DhcpInterfacesRelay ---
func unpackDhcpInterfacesRelayListToSdk(ctx context.Context, list types.List) ([]network_services.DhcpInterfacesRelay, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.DhcpInterfacesRelay")
	diags := diag.Diagnostics{}
	var data []models.DhcpInterfacesRelay
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.DhcpInterfacesRelay, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.DhcpInterfacesRelay{}.AttrTypes(), &item)
		unpacked, d := unpackDhcpInterfacesRelayToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.DhcpInterfacesRelay", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for DhcpInterfacesRelay ---
func packDhcpInterfacesRelayListFromSdk(ctx context.Context, sdks []network_services.DhcpInterfacesRelay) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.DhcpInterfacesRelay")
	diags := diag.Diagnostics{}
	var data []models.DhcpInterfacesRelay

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.DhcpInterfacesRelay
		obj, d := packDhcpInterfacesRelayFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.DhcpInterfacesRelay{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.DhcpInterfacesRelay", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.DhcpInterfacesRelay{}.AttrType(), data)
}

// --- Unpacker for DhcpInterfacesRelayIp ---
func unpackDhcpInterfacesRelayIpToSdk(ctx context.Context, obj types.Object) (*network_services.DhcpInterfacesRelayIp, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.DhcpInterfacesRelayIp", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.DhcpInterfacesRelayIp
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.DhcpInterfacesRelayIp
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Enabled.IsNull() && !model.Enabled.IsUnknown() {
		sdk.Enabled = model.Enabled.ValueBool()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Enabled", "value": sdk.Enabled})
	}

	// Handling Lists
	if !model.Server.IsNull() && !model.Server.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Server")
		diags.Append(model.Server.ElementsAs(ctx, &sdk.Server, false)...)
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.DhcpInterfacesRelayIp", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for DhcpInterfacesRelayIp ---
func packDhcpInterfacesRelayIpFromSdk(ctx context.Context, sdk network_services.DhcpInterfacesRelayIp) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.DhcpInterfacesRelayIp", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.DhcpInterfacesRelayIp
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.Enabled = basetypes.NewBoolValue(sdk.Enabled)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Enabled", "value": sdk.Enabled})
	// Handling Lists
	if sdk.Server != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Server")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Server, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Server)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Server = basetypes.NewListNull(elemType)
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.DhcpInterfacesRelayIp{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.DhcpInterfacesRelayIp", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for DhcpInterfacesRelayIp ---
func unpackDhcpInterfacesRelayIpListToSdk(ctx context.Context, list types.List) ([]network_services.DhcpInterfacesRelayIp, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.DhcpInterfacesRelayIp")
	diags := diag.Diagnostics{}
	var data []models.DhcpInterfacesRelayIp
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.DhcpInterfacesRelayIp, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.DhcpInterfacesRelayIp{}.AttrTypes(), &item)
		unpacked, d := unpackDhcpInterfacesRelayIpToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.DhcpInterfacesRelayIp", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for DhcpInterfacesRelayIp ---
func packDhcpInterfacesRelayIpListFromSdk(ctx context.Context, sdks []network_services.DhcpInterfacesRelayIp) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.DhcpInterfacesRelayIp")
	diags := diag.Diagnostics{}
	var data []models.DhcpInterfacesRelayIp

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.DhcpInterfacesRelayIp
		obj, d := packDhcpInterfacesRelayIpFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.DhcpInterfacesRelayIp{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.DhcpInterfacesRelayIp", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.DhcpInterfacesRelayIp{}.AttrType(), data)
}

// --- Unpacker for DhcpInterfacesServer ---
func unpackDhcpInterfacesServerToSdk(ctx context.Context, obj types.Object) (*network_services.DhcpInterfacesServer, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.DhcpInterfacesServer", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.DhcpInterfacesServer
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.DhcpInterfacesServer
	var d diag.Diagnostics
	// Handling Lists
	if !model.IpPool.IsNull() && !model.IpPool.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field IpPool")
		diags.Append(model.IpPool.ElementsAs(ctx, &sdk.IpPool, false)...)
	}

	// Handling Primitives
	if !model.Mode.IsNull() && !model.Mode.IsUnknown() {
		sdk.Mode = model.Mode.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Mode", "value": *sdk.Mode})
	}

	// Handling Objects
	if !model.Option.IsNull() && !model.Option.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Option")
		unpacked, d := unpackDhcpInterfacesServerOptionToSdk(ctx, model.Option)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Option"})
		}
		if unpacked != nil {
			sdk.Option = unpacked
		}
	}

	// Handling Primitives
	if !model.ProbeIp.IsNull() && !model.ProbeIp.IsUnknown() {
		sdk.ProbeIp = model.ProbeIp.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "ProbeIp", "value": *sdk.ProbeIp})
	}

	// Handling Lists
	if !model.Reserved.IsNull() && !model.Reserved.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field Reserved")
		unpacked, d := unpackDhcpInterfacesServerReservedInnerListToSdk(ctx, model.Reserved)
		diags.Append(d...)
		sdk.Reserved = unpacked
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.DhcpInterfacesServer", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for DhcpInterfacesServer ---
func packDhcpInterfacesServerFromSdk(ctx context.Context, sdk network_services.DhcpInterfacesServer) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.DhcpInterfacesServer", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.DhcpInterfacesServer
	var d diag.Diagnostics
	// Handling Lists
	if sdk.IpPool != nil {
		tflog.Debug(ctx, "Packing list of primitives for field IpPool")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.IpPool, d = basetypes.NewListValueFrom(ctx, elemType, sdk.IpPool)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.IpPool = basetypes.NewListNull(elemType)
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Mode != nil {
		model.Mode = basetypes.NewStringValue(*sdk.Mode)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Mode", "value": *sdk.Mode})
	} else {
		model.Mode = basetypes.NewStringNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Option != nil {
		tflog.Debug(ctx, "Packing nested object for field Option")
		packed, d := packDhcpInterfacesServerOptionFromSdk(ctx, *sdk.Option)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Option"})
		}
		model.Option = packed
	} else {
		model.Option = basetypes.NewObjectNull(models.DhcpInterfacesServerOption{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.ProbeIp != nil {
		model.ProbeIp = basetypes.NewBoolValue(*sdk.ProbeIp)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "ProbeIp", "value": *sdk.ProbeIp})
	} else {
		model.ProbeIp = basetypes.NewBoolNull()
	}
	// Handling Lists
	if sdk.Reserved != nil {
		tflog.Debug(ctx, "Packing list of objects for field Reserved")
		packed, d := packDhcpInterfacesServerReservedInnerListFromSdk(ctx, sdk.Reserved)
		diags.Append(d...)
		model.Reserved = packed
	} else {
		model.Reserved = basetypes.NewListNull(models.DhcpInterfacesServerReservedInner{}.AttrType())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.DhcpInterfacesServer{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.DhcpInterfacesServer", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for DhcpInterfacesServer ---
func unpackDhcpInterfacesServerListToSdk(ctx context.Context, list types.List) ([]network_services.DhcpInterfacesServer, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.DhcpInterfacesServer")
	diags := diag.Diagnostics{}
	var data []models.DhcpInterfacesServer
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.DhcpInterfacesServer, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.DhcpInterfacesServer{}.AttrTypes(), &item)
		unpacked, d := unpackDhcpInterfacesServerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.DhcpInterfacesServer", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for DhcpInterfacesServer ---
func packDhcpInterfacesServerListFromSdk(ctx context.Context, sdks []network_services.DhcpInterfacesServer) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.DhcpInterfacesServer")
	diags := diag.Diagnostics{}
	var data []models.DhcpInterfacesServer

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.DhcpInterfacesServer
		obj, d := packDhcpInterfacesServerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.DhcpInterfacesServer{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.DhcpInterfacesServer", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.DhcpInterfacesServer{}.AttrType(), data)
}

// --- Unpacker for DhcpInterfacesServerOption ---
func unpackDhcpInterfacesServerOptionToSdk(ctx context.Context, obj types.Object) (*network_services.DhcpInterfacesServerOption, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.DhcpInterfacesServerOption", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.DhcpInterfacesServerOption
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.DhcpInterfacesServerOption
	var d diag.Diagnostics
	// Handling Objects
	if !model.Dns.IsNull() && !model.Dns.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Dns")
		unpacked, d := unpackDhcpInterfacesServerOptionDnsToSdk(ctx, model.Dns)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Dns"})
		}
		if unpacked != nil {
			sdk.Dns = unpacked
		}
	}

	// Handling Primitives
	if !model.DnsSuffix.IsNull() && !model.DnsSuffix.IsUnknown() {
		sdk.DnsSuffix = model.DnsSuffix.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DnsSuffix", "value": *sdk.DnsSuffix})
	}

	// Handling Primitives
	if !model.Gateway.IsNull() && !model.Gateway.IsUnknown() {
		sdk.Gateway = model.Gateway.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Gateway", "value": *sdk.Gateway})
	}

	// Handling Objects
	if !model.Inheritance.IsNull() && !model.Inheritance.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Inheritance")
		unpacked, d := unpackDhcpInterfacesServerOptionInheritanceToSdk(ctx, model.Inheritance)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Inheritance"})
		}
		if unpacked != nil {
			sdk.Inheritance = unpacked
		}
	}

	// Handling Objects
	if !model.Lease.IsNull() && !model.Lease.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Lease")
		unpacked, d := unpackDhcpInterfacesServerOptionLeaseToSdk(ctx, model.Lease)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Lease"})
		}
		if unpacked != nil {
			sdk.Lease = unpacked
		}
	}

	// Handling Objects
	if !model.Nis.IsNull() && !model.Nis.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Nis")
		unpacked, d := unpackDhcpInterfacesServerOptionNisToSdk(ctx, model.Nis)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Nis"})
		}
		if unpacked != nil {
			sdk.Nis = unpacked
		}
	}

	// Handling Objects
	if !model.Ntp.IsNull() && !model.Ntp.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Ntp")
		unpacked, d := unpackDhcpInterfacesServerOptionNtpToSdk(ctx, model.Ntp)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Ntp"})
		}
		if unpacked != nil {
			sdk.Ntp = unpacked
		}
	}

	// Handling Primitives
	if !model.Pop3Server.IsNull() && !model.Pop3Server.IsUnknown() {
		sdk.Pop3Server = model.Pop3Server.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Pop3Server", "value": *sdk.Pop3Server})
	}

	// Handling Primitives
	if !model.SmtpServer.IsNull() && !model.SmtpServer.IsUnknown() {
		sdk.SmtpServer = model.SmtpServer.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "SmtpServer", "value": *sdk.SmtpServer})
	}

	// Handling Primitives
	if !model.SubnetMask.IsNull() && !model.SubnetMask.IsUnknown() {
		sdk.SubnetMask = model.SubnetMask.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "SubnetMask", "value": *sdk.SubnetMask})
	}

	// Handling Lists
	if !model.UserDefined.IsNull() && !model.UserDefined.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field UserDefined")
		unpacked, d := unpackDhcpInterfacesServerOptionUserDefinedInnerListToSdk(ctx, model.UserDefined)
		diags.Append(d...)
		sdk.UserDefined = unpacked
	}

	// Handling Objects
	if !model.Wins.IsNull() && !model.Wins.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Wins")
		unpacked, d := unpackDhcpInterfacesServerOptionWinsToSdk(ctx, model.Wins)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Wins"})
		}
		if unpacked != nil {
			sdk.Wins = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.DhcpInterfacesServerOption", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for DhcpInterfacesServerOption ---
func packDhcpInterfacesServerOptionFromSdk(ctx context.Context, sdk network_services.DhcpInterfacesServerOption) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.DhcpInterfacesServerOption", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.DhcpInterfacesServerOption
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Dns != nil {
		tflog.Debug(ctx, "Packing nested object for field Dns")
		packed, d := packDhcpInterfacesServerOptionDnsFromSdk(ctx, *sdk.Dns)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Dns"})
		}
		model.Dns = packed
	} else {
		model.Dns = basetypes.NewObjectNull(models.DhcpInterfacesServerOptionDns{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.DnsSuffix != nil {
		model.DnsSuffix = basetypes.NewStringValue(*sdk.DnsSuffix)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DnsSuffix", "value": *sdk.DnsSuffix})
	} else {
		model.DnsSuffix = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Gateway != nil {
		model.Gateway = basetypes.NewStringValue(*sdk.Gateway)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Gateway", "value": *sdk.Gateway})
	} else {
		model.Gateway = basetypes.NewStringNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Inheritance != nil {
		tflog.Debug(ctx, "Packing nested object for field Inheritance")
		packed, d := packDhcpInterfacesServerOptionInheritanceFromSdk(ctx, *sdk.Inheritance)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Inheritance"})
		}
		model.Inheritance = packed
	} else {
		model.Inheritance = basetypes.NewObjectNull(models.DhcpInterfacesServerOptionInheritance{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Lease != nil {
		tflog.Debug(ctx, "Packing nested object for field Lease")
		packed, d := packDhcpInterfacesServerOptionLeaseFromSdk(ctx, *sdk.Lease)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Lease"})
		}
		model.Lease = packed
	} else {
		model.Lease = basetypes.NewObjectNull(models.DhcpInterfacesServerOptionLease{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Nis != nil {
		tflog.Debug(ctx, "Packing nested object for field Nis")
		packed, d := packDhcpInterfacesServerOptionNisFromSdk(ctx, *sdk.Nis)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Nis"})
		}
		model.Nis = packed
	} else {
		model.Nis = basetypes.NewObjectNull(models.DhcpInterfacesServerOptionNis{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Ntp != nil {
		tflog.Debug(ctx, "Packing nested object for field Ntp")
		packed, d := packDhcpInterfacesServerOptionNtpFromSdk(ctx, *sdk.Ntp)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Ntp"})
		}
		model.Ntp = packed
	} else {
		model.Ntp = basetypes.NewObjectNull(models.DhcpInterfacesServerOptionNtp{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Pop3Server != nil {
		model.Pop3Server = basetypes.NewStringValue(*sdk.Pop3Server)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Pop3Server", "value": *sdk.Pop3Server})
	} else {
		model.Pop3Server = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.SmtpServer != nil {
		model.SmtpServer = basetypes.NewStringValue(*sdk.SmtpServer)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "SmtpServer", "value": *sdk.SmtpServer})
	} else {
		model.SmtpServer = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.SubnetMask != nil {
		model.SubnetMask = basetypes.NewStringValue(*sdk.SubnetMask)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "SubnetMask", "value": *sdk.SubnetMask})
	} else {
		model.SubnetMask = basetypes.NewStringNull()
	}
	// Handling Lists
	if sdk.UserDefined != nil {
		tflog.Debug(ctx, "Packing list of objects for field UserDefined")
		packed, d := packDhcpInterfacesServerOptionUserDefinedInnerListFromSdk(ctx, sdk.UserDefined)
		diags.Append(d...)
		model.UserDefined = packed
	} else {
		model.UserDefined = basetypes.NewListNull(models.DhcpInterfacesServerOptionUserDefinedInner{}.AttrType())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Wins != nil {
		tflog.Debug(ctx, "Packing nested object for field Wins")
		packed, d := packDhcpInterfacesServerOptionWinsFromSdk(ctx, *sdk.Wins)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Wins"})
		}
		model.Wins = packed
	} else {
		model.Wins = basetypes.NewObjectNull(models.DhcpInterfacesServerOptionWins{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.DhcpInterfacesServerOption{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.DhcpInterfacesServerOption", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for DhcpInterfacesServerOption ---
func unpackDhcpInterfacesServerOptionListToSdk(ctx context.Context, list types.List) ([]network_services.DhcpInterfacesServerOption, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.DhcpInterfacesServerOption")
	diags := diag.Diagnostics{}
	var data []models.DhcpInterfacesServerOption
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.DhcpInterfacesServerOption, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.DhcpInterfacesServerOption{}.AttrTypes(), &item)
		unpacked, d := unpackDhcpInterfacesServerOptionToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.DhcpInterfacesServerOption", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for DhcpInterfacesServerOption ---
func packDhcpInterfacesServerOptionListFromSdk(ctx context.Context, sdks []network_services.DhcpInterfacesServerOption) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.DhcpInterfacesServerOption")
	diags := diag.Diagnostics{}
	var data []models.DhcpInterfacesServerOption

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.DhcpInterfacesServerOption
		obj, d := packDhcpInterfacesServerOptionFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.DhcpInterfacesServerOption{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.DhcpInterfacesServerOption", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.DhcpInterfacesServerOption{}.AttrType(), data)
}

// --- Unpacker for DhcpInterfacesServerOptionDns ---
func unpackDhcpInterfacesServerOptionDnsToSdk(ctx context.Context, obj types.Object) (*network_services.DhcpInterfacesServerOptionDns, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.DhcpInterfacesServerOptionDns", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.DhcpInterfacesServerOptionDns
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.DhcpInterfacesServerOptionDns
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Primary.IsNull() && !model.Primary.IsUnknown() {
		sdk.Primary = model.Primary.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Primary", "value": *sdk.Primary})
	}

	// Handling Primitives
	if !model.Secondary.IsNull() && !model.Secondary.IsUnknown() {
		sdk.Secondary = model.Secondary.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Secondary", "value": *sdk.Secondary})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.DhcpInterfacesServerOptionDns", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for DhcpInterfacesServerOptionDns ---
func packDhcpInterfacesServerOptionDnsFromSdk(ctx context.Context, sdk network_services.DhcpInterfacesServerOptionDns) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.DhcpInterfacesServerOptionDns", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.DhcpInterfacesServerOptionDns
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Primary != nil {
		model.Primary = basetypes.NewStringValue(*sdk.Primary)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Primary", "value": *sdk.Primary})
	} else {
		model.Primary = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Secondary != nil {
		model.Secondary = basetypes.NewStringValue(*sdk.Secondary)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Secondary", "value": *sdk.Secondary})
	} else {
		model.Secondary = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.DhcpInterfacesServerOptionDns{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.DhcpInterfacesServerOptionDns", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for DhcpInterfacesServerOptionDns ---
func unpackDhcpInterfacesServerOptionDnsListToSdk(ctx context.Context, list types.List) ([]network_services.DhcpInterfacesServerOptionDns, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.DhcpInterfacesServerOptionDns")
	diags := diag.Diagnostics{}
	var data []models.DhcpInterfacesServerOptionDns
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.DhcpInterfacesServerOptionDns, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.DhcpInterfacesServerOptionDns{}.AttrTypes(), &item)
		unpacked, d := unpackDhcpInterfacesServerOptionDnsToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.DhcpInterfacesServerOptionDns", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for DhcpInterfacesServerOptionDns ---
func packDhcpInterfacesServerOptionDnsListFromSdk(ctx context.Context, sdks []network_services.DhcpInterfacesServerOptionDns) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.DhcpInterfacesServerOptionDns")
	diags := diag.Diagnostics{}
	var data []models.DhcpInterfacesServerOptionDns

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.DhcpInterfacesServerOptionDns
		obj, d := packDhcpInterfacesServerOptionDnsFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.DhcpInterfacesServerOptionDns{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.DhcpInterfacesServerOptionDns", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.DhcpInterfacesServerOptionDns{}.AttrType(), data)
}

// --- Unpacker for DhcpInterfacesServerOptionInheritance ---
func unpackDhcpInterfacesServerOptionInheritanceToSdk(ctx context.Context, obj types.Object) (*network_services.DhcpInterfacesServerOptionInheritance, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.DhcpInterfacesServerOptionInheritance", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.DhcpInterfacesServerOptionInheritance
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.DhcpInterfacesServerOptionInheritance
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Source.IsNull() && !model.Source.IsUnknown() {
		sdk.Source = model.Source.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Source", "value": *sdk.Source})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.DhcpInterfacesServerOptionInheritance", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for DhcpInterfacesServerOptionInheritance ---
func packDhcpInterfacesServerOptionInheritanceFromSdk(ctx context.Context, sdk network_services.DhcpInterfacesServerOptionInheritance) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.DhcpInterfacesServerOptionInheritance", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.DhcpInterfacesServerOptionInheritance
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Source != nil {
		model.Source = basetypes.NewStringValue(*sdk.Source)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Source", "value": *sdk.Source})
	} else {
		model.Source = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.DhcpInterfacesServerOptionInheritance{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.DhcpInterfacesServerOptionInheritance", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for DhcpInterfacesServerOptionInheritance ---
func unpackDhcpInterfacesServerOptionInheritanceListToSdk(ctx context.Context, list types.List) ([]network_services.DhcpInterfacesServerOptionInheritance, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.DhcpInterfacesServerOptionInheritance")
	diags := diag.Diagnostics{}
	var data []models.DhcpInterfacesServerOptionInheritance
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.DhcpInterfacesServerOptionInheritance, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.DhcpInterfacesServerOptionInheritance{}.AttrTypes(), &item)
		unpacked, d := unpackDhcpInterfacesServerOptionInheritanceToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.DhcpInterfacesServerOptionInheritance", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for DhcpInterfacesServerOptionInheritance ---
func packDhcpInterfacesServerOptionInheritanceListFromSdk(ctx context.Context, sdks []network_services.DhcpInterfacesServerOptionInheritance) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.DhcpInterfacesServerOptionInheritance")
	diags := diag.Diagnostics{}
	var data []models.DhcpInterfacesServerOptionInheritance

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.DhcpInterfacesServerOptionInheritance
		obj, d := packDhcpInterfacesServerOptionInheritanceFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.DhcpInterfacesServerOptionInheritance{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.DhcpInterfacesServerOptionInheritance", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.DhcpInterfacesServerOptionInheritance{}.AttrType(), data)
}

// --- Unpacker for DhcpInterfacesServerOptionLease ---
func unpackDhcpInterfacesServerOptionLeaseToSdk(ctx context.Context, obj types.Object) (*network_services.DhcpInterfacesServerOptionLease, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.DhcpInterfacesServerOptionLease", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.DhcpInterfacesServerOptionLease
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.DhcpInterfacesServerOptionLease
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Timeout.IsNull() && !model.Timeout.IsUnknown() {
		val := int32(model.Timeout.ValueInt64())
		sdk.Timeout = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Timeout", "value": *sdk.Timeout})
	}

	// Handling Typeless Objects
	if !model.Unlimited.IsNull() && !model.Unlimited.IsUnknown() {
		tflog.Debug(ctx, "Unpacking typeless object for field Unlimited")
		sdk.Unlimited = make(map[string]interface{})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.DhcpInterfacesServerOptionLease", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for DhcpInterfacesServerOptionLease ---
func packDhcpInterfacesServerOptionLeaseFromSdk(ctx context.Context, sdk network_services.DhcpInterfacesServerOptionLease) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.DhcpInterfacesServerOptionLease", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.DhcpInterfacesServerOptionLease
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Timeout != nil {
		model.Timeout = basetypes.NewInt64Value(int64(*sdk.Timeout))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Timeout", "value": *sdk.Timeout})
	} else {
		model.Timeout = basetypes.NewInt64Null()
	}
	// Handling Objects
	// This is a marker object (e.g. CHAP: {}). We just need to create an empty, non-null object.
	if sdk.Unlimited != nil && !reflect.ValueOf(sdk.Unlimited).IsNil() {
		tflog.Debug(ctx, "Packing typeless object for field Unlimited")
		var d diag.Diagnostics
		// Create an empty object with no attributes, which signifies its presence.
		model.Unlimited, d = basetypes.NewObjectValue(map[string]attr.Type{}, map[string]attr.Value{})
		diags.Append(d...)
	} else {
		// Since this field is part of a oneOf, being nil means it's not selected.
		// We make the object null with an empty attribute map.
		model.Unlimited = basetypes.NewObjectNull(map[string]attr.Type{})
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.DhcpInterfacesServerOptionLease{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.DhcpInterfacesServerOptionLease", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for DhcpInterfacesServerOptionLease ---
func unpackDhcpInterfacesServerOptionLeaseListToSdk(ctx context.Context, list types.List) ([]network_services.DhcpInterfacesServerOptionLease, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.DhcpInterfacesServerOptionLease")
	diags := diag.Diagnostics{}
	var data []models.DhcpInterfacesServerOptionLease
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.DhcpInterfacesServerOptionLease, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.DhcpInterfacesServerOptionLease{}.AttrTypes(), &item)
		unpacked, d := unpackDhcpInterfacesServerOptionLeaseToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.DhcpInterfacesServerOptionLease", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for DhcpInterfacesServerOptionLease ---
func packDhcpInterfacesServerOptionLeaseListFromSdk(ctx context.Context, sdks []network_services.DhcpInterfacesServerOptionLease) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.DhcpInterfacesServerOptionLease")
	diags := diag.Diagnostics{}
	var data []models.DhcpInterfacesServerOptionLease

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.DhcpInterfacesServerOptionLease
		obj, d := packDhcpInterfacesServerOptionLeaseFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.DhcpInterfacesServerOptionLease{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.DhcpInterfacesServerOptionLease", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.DhcpInterfacesServerOptionLease{}.AttrType(), data)
}

// --- Unpacker for DhcpInterfacesServerOptionNis ---
func unpackDhcpInterfacesServerOptionNisToSdk(ctx context.Context, obj types.Object) (*network_services.DhcpInterfacesServerOptionNis, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.DhcpInterfacesServerOptionNis", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.DhcpInterfacesServerOptionNis
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.DhcpInterfacesServerOptionNis
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Primary.IsNull() && !model.Primary.IsUnknown() {
		sdk.Primary = model.Primary.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Primary", "value": *sdk.Primary})
	}

	// Handling Primitives
	if !model.Secondary.IsNull() && !model.Secondary.IsUnknown() {
		sdk.Secondary = model.Secondary.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Secondary", "value": *sdk.Secondary})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.DhcpInterfacesServerOptionNis", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for DhcpInterfacesServerOptionNis ---
func packDhcpInterfacesServerOptionNisFromSdk(ctx context.Context, sdk network_services.DhcpInterfacesServerOptionNis) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.DhcpInterfacesServerOptionNis", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.DhcpInterfacesServerOptionNis
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Primary != nil {
		model.Primary = basetypes.NewStringValue(*sdk.Primary)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Primary", "value": *sdk.Primary})
	} else {
		model.Primary = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Secondary != nil {
		model.Secondary = basetypes.NewStringValue(*sdk.Secondary)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Secondary", "value": *sdk.Secondary})
	} else {
		model.Secondary = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.DhcpInterfacesServerOptionNis{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.DhcpInterfacesServerOptionNis", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for DhcpInterfacesServerOptionNis ---
func unpackDhcpInterfacesServerOptionNisListToSdk(ctx context.Context, list types.List) ([]network_services.DhcpInterfacesServerOptionNis, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.DhcpInterfacesServerOptionNis")
	diags := diag.Diagnostics{}
	var data []models.DhcpInterfacesServerOptionNis
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.DhcpInterfacesServerOptionNis, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.DhcpInterfacesServerOptionNis{}.AttrTypes(), &item)
		unpacked, d := unpackDhcpInterfacesServerOptionNisToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.DhcpInterfacesServerOptionNis", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for DhcpInterfacesServerOptionNis ---
func packDhcpInterfacesServerOptionNisListFromSdk(ctx context.Context, sdks []network_services.DhcpInterfacesServerOptionNis) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.DhcpInterfacesServerOptionNis")
	diags := diag.Diagnostics{}
	var data []models.DhcpInterfacesServerOptionNis

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.DhcpInterfacesServerOptionNis
		obj, d := packDhcpInterfacesServerOptionNisFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.DhcpInterfacesServerOptionNis{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.DhcpInterfacesServerOptionNis", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.DhcpInterfacesServerOptionNis{}.AttrType(), data)
}

// --- Unpacker for DhcpInterfacesServerOptionNtp ---
func unpackDhcpInterfacesServerOptionNtpToSdk(ctx context.Context, obj types.Object) (*network_services.DhcpInterfacesServerOptionNtp, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.DhcpInterfacesServerOptionNtp", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.DhcpInterfacesServerOptionNtp
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.DhcpInterfacesServerOptionNtp
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Primary.IsNull() && !model.Primary.IsUnknown() {
		sdk.Primary = model.Primary.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Primary", "value": *sdk.Primary})
	}

	// Handling Primitives
	if !model.Secondary.IsNull() && !model.Secondary.IsUnknown() {
		sdk.Secondary = model.Secondary.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Secondary", "value": *sdk.Secondary})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.DhcpInterfacesServerOptionNtp", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for DhcpInterfacesServerOptionNtp ---
func packDhcpInterfacesServerOptionNtpFromSdk(ctx context.Context, sdk network_services.DhcpInterfacesServerOptionNtp) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.DhcpInterfacesServerOptionNtp", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.DhcpInterfacesServerOptionNtp
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Primary != nil {
		model.Primary = basetypes.NewStringValue(*sdk.Primary)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Primary", "value": *sdk.Primary})
	} else {
		model.Primary = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Secondary != nil {
		model.Secondary = basetypes.NewStringValue(*sdk.Secondary)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Secondary", "value": *sdk.Secondary})
	} else {
		model.Secondary = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.DhcpInterfacesServerOptionNtp{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.DhcpInterfacesServerOptionNtp", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for DhcpInterfacesServerOptionNtp ---
func unpackDhcpInterfacesServerOptionNtpListToSdk(ctx context.Context, list types.List) ([]network_services.DhcpInterfacesServerOptionNtp, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.DhcpInterfacesServerOptionNtp")
	diags := diag.Diagnostics{}
	var data []models.DhcpInterfacesServerOptionNtp
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.DhcpInterfacesServerOptionNtp, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.DhcpInterfacesServerOptionNtp{}.AttrTypes(), &item)
		unpacked, d := unpackDhcpInterfacesServerOptionNtpToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.DhcpInterfacesServerOptionNtp", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for DhcpInterfacesServerOptionNtp ---
func packDhcpInterfacesServerOptionNtpListFromSdk(ctx context.Context, sdks []network_services.DhcpInterfacesServerOptionNtp) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.DhcpInterfacesServerOptionNtp")
	diags := diag.Diagnostics{}
	var data []models.DhcpInterfacesServerOptionNtp

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.DhcpInterfacesServerOptionNtp
		obj, d := packDhcpInterfacesServerOptionNtpFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.DhcpInterfacesServerOptionNtp{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.DhcpInterfacesServerOptionNtp", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.DhcpInterfacesServerOptionNtp{}.AttrType(), data)
}

// --- Unpacker for DhcpInterfacesServerOptionUserDefinedInner ---
func unpackDhcpInterfacesServerOptionUserDefinedInnerToSdk(ctx context.Context, obj types.Object) (*network_services.DhcpInterfacesServerOptionUserDefinedInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.DhcpInterfacesServerOptionUserDefinedInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.DhcpInterfacesServerOptionUserDefinedInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.DhcpInterfacesServerOptionUserDefinedInner
	var d diag.Diagnostics
	// Handling Lists
	if !model.Ascii.IsNull() && !model.Ascii.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Ascii")
		diags.Append(model.Ascii.ElementsAs(ctx, &sdk.Ascii, false)...)
	}

	// Handling Primitives
	if !model.Code.IsNull() && !model.Code.IsUnknown() {
		val := int32(model.Code.ValueInt64())
		sdk.Code = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Code", "value": *sdk.Code})
	}

	// Handling Lists
	if !model.Hex.IsNull() && !model.Hex.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Hex")
		diags.Append(model.Hex.ElementsAs(ctx, &sdk.Hex, false)...)
	}

	// Handling Primitives
	if !model.Inherited.IsNull() && !model.Inherited.IsUnknown() {
		sdk.Inherited = model.Inherited.ValueBool()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Inherited", "value": sdk.Inherited})
	}

	// Handling Lists
	if !model.Ip.IsNull() && !model.Ip.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Ip")
		diags.Append(model.Ip.ElementsAs(ctx, &sdk.Ip, false)...)
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.DhcpInterfacesServerOptionUserDefinedInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for DhcpInterfacesServerOptionUserDefinedInner ---
func packDhcpInterfacesServerOptionUserDefinedInnerFromSdk(ctx context.Context, sdk network_services.DhcpInterfacesServerOptionUserDefinedInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.DhcpInterfacesServerOptionUserDefinedInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.DhcpInterfacesServerOptionUserDefinedInner
	var d diag.Diagnostics
	// Handling Lists
	if sdk.Ascii != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Ascii")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Ascii, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Ascii)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Ascii = basetypes.NewListNull(elemType)
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Code != nil {
		model.Code = basetypes.NewInt64Value(int64(*sdk.Code))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Code", "value": *sdk.Code})
	} else {
		model.Code = basetypes.NewInt64Null()
	}
	// Handling Lists
	if sdk.Hex != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Hex")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Hex, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Hex)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Hex = basetypes.NewListNull(elemType)
	}
	// Handling Primitives
	// Standard primitive packing
	model.Inherited = basetypes.NewBoolValue(sdk.Inherited)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Inherited", "value": sdk.Inherited})
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
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.DhcpInterfacesServerOptionUserDefinedInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.DhcpInterfacesServerOptionUserDefinedInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for DhcpInterfacesServerOptionUserDefinedInner ---
func unpackDhcpInterfacesServerOptionUserDefinedInnerListToSdk(ctx context.Context, list types.List) ([]network_services.DhcpInterfacesServerOptionUserDefinedInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.DhcpInterfacesServerOptionUserDefinedInner")
	diags := diag.Diagnostics{}
	var data []models.DhcpInterfacesServerOptionUserDefinedInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.DhcpInterfacesServerOptionUserDefinedInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.DhcpInterfacesServerOptionUserDefinedInner{}.AttrTypes(), &item)
		unpacked, d := unpackDhcpInterfacesServerOptionUserDefinedInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.DhcpInterfacesServerOptionUserDefinedInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for DhcpInterfacesServerOptionUserDefinedInner ---
func packDhcpInterfacesServerOptionUserDefinedInnerListFromSdk(ctx context.Context, sdks []network_services.DhcpInterfacesServerOptionUserDefinedInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.DhcpInterfacesServerOptionUserDefinedInner")
	diags := diag.Diagnostics{}
	var data []models.DhcpInterfacesServerOptionUserDefinedInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.DhcpInterfacesServerOptionUserDefinedInner
		obj, d := packDhcpInterfacesServerOptionUserDefinedInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.DhcpInterfacesServerOptionUserDefinedInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.DhcpInterfacesServerOptionUserDefinedInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.DhcpInterfacesServerOptionUserDefinedInner{}.AttrType(), data)
}

// --- Unpacker for DhcpInterfacesServerOptionWins ---
func unpackDhcpInterfacesServerOptionWinsToSdk(ctx context.Context, obj types.Object) (*network_services.DhcpInterfacesServerOptionWins, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.DhcpInterfacesServerOptionWins", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.DhcpInterfacesServerOptionWins
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.DhcpInterfacesServerOptionWins
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Primary.IsNull() && !model.Primary.IsUnknown() {
		sdk.Primary = model.Primary.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Primary", "value": *sdk.Primary})
	}

	// Handling Primitives
	if !model.Secondary.IsNull() && !model.Secondary.IsUnknown() {
		sdk.Secondary = model.Secondary.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Secondary", "value": *sdk.Secondary})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.DhcpInterfacesServerOptionWins", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for DhcpInterfacesServerOptionWins ---
func packDhcpInterfacesServerOptionWinsFromSdk(ctx context.Context, sdk network_services.DhcpInterfacesServerOptionWins) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.DhcpInterfacesServerOptionWins", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.DhcpInterfacesServerOptionWins
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Primary != nil {
		model.Primary = basetypes.NewStringValue(*sdk.Primary)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Primary", "value": *sdk.Primary})
	} else {
		model.Primary = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Secondary != nil {
		model.Secondary = basetypes.NewStringValue(*sdk.Secondary)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Secondary", "value": *sdk.Secondary})
	} else {
		model.Secondary = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.DhcpInterfacesServerOptionWins{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.DhcpInterfacesServerOptionWins", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for DhcpInterfacesServerOptionWins ---
func unpackDhcpInterfacesServerOptionWinsListToSdk(ctx context.Context, list types.List) ([]network_services.DhcpInterfacesServerOptionWins, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.DhcpInterfacesServerOptionWins")
	diags := diag.Diagnostics{}
	var data []models.DhcpInterfacesServerOptionWins
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.DhcpInterfacesServerOptionWins, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.DhcpInterfacesServerOptionWins{}.AttrTypes(), &item)
		unpacked, d := unpackDhcpInterfacesServerOptionWinsToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.DhcpInterfacesServerOptionWins", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for DhcpInterfacesServerOptionWins ---
func packDhcpInterfacesServerOptionWinsListFromSdk(ctx context.Context, sdks []network_services.DhcpInterfacesServerOptionWins) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.DhcpInterfacesServerOptionWins")
	diags := diag.Diagnostics{}
	var data []models.DhcpInterfacesServerOptionWins

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.DhcpInterfacesServerOptionWins
		obj, d := packDhcpInterfacesServerOptionWinsFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.DhcpInterfacesServerOptionWins{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.DhcpInterfacesServerOptionWins", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.DhcpInterfacesServerOptionWins{}.AttrType(), data)
}

// --- Unpacker for DhcpInterfacesServerReservedInner ---
func unpackDhcpInterfacesServerReservedInnerToSdk(ctx context.Context, obj types.Object) (*network_services.DhcpInterfacesServerReservedInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.DhcpInterfacesServerReservedInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.DhcpInterfacesServerReservedInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.DhcpInterfacesServerReservedInner
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Description.IsNull() && !model.Description.IsUnknown() {
		sdk.Description = model.Description.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Description", "value": *sdk.Description})
	}

	// Handling Primitives
	if !model.Mac.IsNull() && !model.Mac.IsUnknown() {
		sdk.Mac = model.Mac.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Mac", "value": *sdk.Mac})
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.DhcpInterfacesServerReservedInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for DhcpInterfacesServerReservedInner ---
func packDhcpInterfacesServerReservedInnerFromSdk(ctx context.Context, sdk network_services.DhcpInterfacesServerReservedInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.DhcpInterfacesServerReservedInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.DhcpInterfacesServerReservedInner
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
	if sdk.Mac != nil {
		model.Mac = basetypes.NewStringValue(*sdk.Mac)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Mac", "value": *sdk.Mac})
	} else {
		model.Mac = basetypes.NewStringNull()
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

	obj, d := types.ObjectValueFrom(ctx, models.DhcpInterfacesServerReservedInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.DhcpInterfacesServerReservedInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for DhcpInterfacesServerReservedInner ---
func unpackDhcpInterfacesServerReservedInnerListToSdk(ctx context.Context, list types.List) ([]network_services.DhcpInterfacesServerReservedInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.DhcpInterfacesServerReservedInner")
	diags := diag.Diagnostics{}
	var data []models.DhcpInterfacesServerReservedInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.DhcpInterfacesServerReservedInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.DhcpInterfacesServerReservedInner{}.AttrTypes(), &item)
		unpacked, d := unpackDhcpInterfacesServerReservedInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.DhcpInterfacesServerReservedInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for DhcpInterfacesServerReservedInner ---
func packDhcpInterfacesServerReservedInnerListFromSdk(ctx context.Context, sdks []network_services.DhcpInterfacesServerReservedInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.DhcpInterfacesServerReservedInner")
	diags := diag.Diagnostics{}
	var data []models.DhcpInterfacesServerReservedInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.DhcpInterfacesServerReservedInner
		obj, d := packDhcpInterfacesServerReservedInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.DhcpInterfacesServerReservedInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.DhcpInterfacesServerReservedInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.DhcpInterfacesServerReservedInner{}.AttrType(), data)
}
