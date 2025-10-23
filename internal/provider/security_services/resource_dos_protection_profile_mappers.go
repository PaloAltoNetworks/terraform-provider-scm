package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/security_services"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/security_services"
)

// --- Unpacker for DosProtectionProfiles ---
func unpackDosProtectionProfilesToSdk(ctx context.Context, obj types.Object) (*security_services.DosProtectionProfiles, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.DosProtectionProfiles", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.DosProtectionProfiles
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.DosProtectionProfiles
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

	// Handling Objects
	if !model.Flood.IsNull() && !model.Flood.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Flood")
		unpacked, d := unpackDosProtectionProfilesFloodToSdk(ctx, model.Flood)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Flood"})
		}
		if unpacked != nil {
			sdk.Flood = unpacked
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
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	// Handling Objects
	if !model.Resource.IsNull() && !model.Resource.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Resource")
		unpacked, d := unpackDosProtectionProfilesResourceToSdk(ctx, model.Resource)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Resource"})
		}
		if unpacked != nil {
			sdk.Resource = unpacked
		}
	}

	// Handling Primitives
	if !model.Snippet.IsNull() && !model.Snippet.IsUnknown() {
		sdk.Snippet = model.Snippet.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	}

	// Handling Primitives
	if !model.Type.IsNull() && !model.Type.IsUnknown() {
		sdk.Type = model.Type.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Type", "value": sdk.Type})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.DosProtectionProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for DosProtectionProfiles ---
func packDosProtectionProfilesFromSdk(ctx context.Context, sdk security_services.DosProtectionProfiles) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.DosProtectionProfiles", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.DosProtectionProfiles
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
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Flood != nil {
		tflog.Debug(ctx, "Packing nested object for field Flood")
		packed, d := packDosProtectionProfilesFloodFromSdk(ctx, *sdk.Flood)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Flood"})
		}
		model.Flood = packed
	} else {
		model.Flood = basetypes.NewObjectNull(models.DosProtectionProfilesFlood{}.AttrTypes())
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
	if sdk.Resource != nil {
		tflog.Debug(ctx, "Packing nested object for field Resource")
		packed, d := packDosProtectionProfilesResourceFromSdk(ctx, *sdk.Resource)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Resource"})
		}
		model.Resource = packed
	} else {
		model.Resource = basetypes.NewObjectNull(models.DosProtectionProfilesResource{}.AttrTypes())
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
	model.Type = basetypes.NewStringValue(sdk.Type)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Type", "value": sdk.Type})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.DosProtectionProfiles{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.DosProtectionProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for DosProtectionProfiles ---
func unpackDosProtectionProfilesListToSdk(ctx context.Context, list types.List) ([]security_services.DosProtectionProfiles, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.DosProtectionProfiles")
	diags := diag.Diagnostics{}
	var data []models.DosProtectionProfiles
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.DosProtectionProfiles, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.DosProtectionProfiles{}.AttrTypes(), &item)
		unpacked, d := unpackDosProtectionProfilesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.DosProtectionProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for DosProtectionProfiles ---
func packDosProtectionProfilesListFromSdk(ctx context.Context, sdks []security_services.DosProtectionProfiles) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.DosProtectionProfiles")
	diags := diag.Diagnostics{}
	var data []models.DosProtectionProfiles

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.DosProtectionProfiles
		obj, d := packDosProtectionProfilesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.DosProtectionProfiles{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.DosProtectionProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.DosProtectionProfiles{}.AttrType(), data)
}

// --- Unpacker for DosProtectionProfilesFlood ---
func unpackDosProtectionProfilesFloodToSdk(ctx context.Context, obj types.Object) (*security_services.DosProtectionProfilesFlood, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.DosProtectionProfilesFlood", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.DosProtectionProfilesFlood
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.DosProtectionProfilesFlood
	var d diag.Diagnostics
	// Handling Objects
	if !model.Icmp.IsNull() && !model.Icmp.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Icmp")
		unpacked, d := unpackDosProtectionProfilesFloodIcmpToSdk(ctx, model.Icmp)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Icmp"})
		}
		if unpacked != nil {
			sdk.Icmp = unpacked
		}
	}

	// Handling Objects
	if !model.Icmpv6.IsNull() && !model.Icmpv6.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Icmpv6")
		unpacked, d := unpackDosProtectionProfilesFloodIcmpToSdk(ctx, model.Icmpv6)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Icmpv6"})
		}
		if unpacked != nil {
			sdk.Icmpv6 = unpacked
		}
	}

	// Handling Objects
	if !model.OtherIp.IsNull() && !model.OtherIp.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field OtherIp")
		unpacked, d := unpackDosProtectionProfilesFloodIcmpToSdk(ctx, model.OtherIp)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "OtherIp"})
		}
		if unpacked != nil {
			sdk.OtherIp = unpacked
		}
	}

	// Handling Objects
	if !model.TcpSyn.IsNull() && !model.TcpSyn.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field TcpSyn")
		unpacked, d := unpackDosProtectionProfilesFloodTcpSynToSdk(ctx, model.TcpSyn)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "TcpSyn"})
		}
		if unpacked != nil {
			sdk.TcpSyn = unpacked
		}
	}

	// Handling Objects
	if !model.Udp.IsNull() && !model.Udp.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Udp")
		unpacked, d := unpackDosProtectionProfilesFloodIcmpToSdk(ctx, model.Udp)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Udp"})
		}
		if unpacked != nil {
			sdk.Udp = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.DosProtectionProfilesFlood", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for DosProtectionProfilesFlood ---
func packDosProtectionProfilesFloodFromSdk(ctx context.Context, sdk security_services.DosProtectionProfilesFlood) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.DosProtectionProfilesFlood", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.DosProtectionProfilesFlood
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Icmp != nil {
		tflog.Debug(ctx, "Packing nested object for field Icmp")
		packed, d := packDosProtectionProfilesFloodIcmpFromSdk(ctx, *sdk.Icmp)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Icmp"})
		}
		model.Icmp = packed
	} else {
		model.Icmp = basetypes.NewObjectNull(models.DosProtectionProfilesFloodIcmp{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Icmpv6 != nil {
		tflog.Debug(ctx, "Packing nested object for field Icmpv6")
		packed, d := packDosProtectionProfilesFloodIcmpFromSdk(ctx, *sdk.Icmpv6)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Icmpv6"})
		}
		model.Icmpv6 = packed
	} else {
		model.Icmpv6 = basetypes.NewObjectNull(models.DosProtectionProfilesFloodIcmp{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.OtherIp != nil {
		tflog.Debug(ctx, "Packing nested object for field OtherIp")
		packed, d := packDosProtectionProfilesFloodIcmpFromSdk(ctx, *sdk.OtherIp)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "OtherIp"})
		}
		model.OtherIp = packed
	} else {
		model.OtherIp = basetypes.NewObjectNull(models.DosProtectionProfilesFloodIcmp{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.TcpSyn != nil {
		tflog.Debug(ctx, "Packing nested object for field TcpSyn")
		packed, d := packDosProtectionProfilesFloodTcpSynFromSdk(ctx, *sdk.TcpSyn)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "TcpSyn"})
		}
		model.TcpSyn = packed
	} else {
		model.TcpSyn = basetypes.NewObjectNull(models.DosProtectionProfilesFloodTcpSyn{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Udp != nil {
		tflog.Debug(ctx, "Packing nested object for field Udp")
		packed, d := packDosProtectionProfilesFloodIcmpFromSdk(ctx, *sdk.Udp)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Udp"})
		}
		model.Udp = packed
	} else {
		model.Udp = basetypes.NewObjectNull(models.DosProtectionProfilesFloodIcmp{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.DosProtectionProfilesFlood{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.DosProtectionProfilesFlood", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for DosProtectionProfilesFlood ---
func unpackDosProtectionProfilesFloodListToSdk(ctx context.Context, list types.List) ([]security_services.DosProtectionProfilesFlood, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.DosProtectionProfilesFlood")
	diags := diag.Diagnostics{}
	var data []models.DosProtectionProfilesFlood
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.DosProtectionProfilesFlood, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.DosProtectionProfilesFlood{}.AttrTypes(), &item)
		unpacked, d := unpackDosProtectionProfilesFloodToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.DosProtectionProfilesFlood", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for DosProtectionProfilesFlood ---
func packDosProtectionProfilesFloodListFromSdk(ctx context.Context, sdks []security_services.DosProtectionProfilesFlood) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.DosProtectionProfilesFlood")
	diags := diag.Diagnostics{}
	var data []models.DosProtectionProfilesFlood

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.DosProtectionProfilesFlood
		obj, d := packDosProtectionProfilesFloodFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.DosProtectionProfilesFlood{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.DosProtectionProfilesFlood", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.DosProtectionProfilesFlood{}.AttrType(), data)
}

// --- Unpacker for DosProtectionProfilesFloodIcmp ---
func unpackDosProtectionProfilesFloodIcmpToSdk(ctx context.Context, obj types.Object) (*security_services.DosProtectionProfilesFloodIcmp, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.DosProtectionProfilesFloodIcmp", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.DosProtectionProfilesFloodIcmp
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.DosProtectionProfilesFloodIcmp
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Enable.IsNull() && !model.Enable.IsUnknown() {
		sdk.Enable = model.Enable.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Enable", "value": *sdk.Enable})
	}

	// Handling Objects
	if !model.Red.IsNull() && !model.Red.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Red")
		unpacked, d := unpackDosProtectionProfilesFloodIcmpRedToSdk(ctx, model.Red)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Red"})
		}
		if unpacked != nil {
			sdk.Red = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.DosProtectionProfilesFloodIcmp", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for DosProtectionProfilesFloodIcmp ---
func packDosProtectionProfilesFloodIcmpFromSdk(ctx context.Context, sdk security_services.DosProtectionProfilesFloodIcmp) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.DosProtectionProfilesFloodIcmp", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.DosProtectionProfilesFloodIcmp
	var d diag.Diagnostics
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
	if sdk.Red != nil {
		tflog.Debug(ctx, "Packing nested object for field Red")
		packed, d := packDosProtectionProfilesFloodIcmpRedFromSdk(ctx, *sdk.Red)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Red"})
		}
		model.Red = packed
	} else {
		model.Red = basetypes.NewObjectNull(models.DosProtectionProfilesFloodIcmpRed{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.DosProtectionProfilesFloodIcmp{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.DosProtectionProfilesFloodIcmp", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for DosProtectionProfilesFloodIcmp ---
func unpackDosProtectionProfilesFloodIcmpListToSdk(ctx context.Context, list types.List) ([]security_services.DosProtectionProfilesFloodIcmp, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.DosProtectionProfilesFloodIcmp")
	diags := diag.Diagnostics{}
	var data []models.DosProtectionProfilesFloodIcmp
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.DosProtectionProfilesFloodIcmp, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.DosProtectionProfilesFloodIcmp{}.AttrTypes(), &item)
		unpacked, d := unpackDosProtectionProfilesFloodIcmpToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.DosProtectionProfilesFloodIcmp", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for DosProtectionProfilesFloodIcmp ---
func packDosProtectionProfilesFloodIcmpListFromSdk(ctx context.Context, sdks []security_services.DosProtectionProfilesFloodIcmp) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.DosProtectionProfilesFloodIcmp")
	diags := diag.Diagnostics{}
	var data []models.DosProtectionProfilesFloodIcmp

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.DosProtectionProfilesFloodIcmp
		obj, d := packDosProtectionProfilesFloodIcmpFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.DosProtectionProfilesFloodIcmp{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.DosProtectionProfilesFloodIcmp", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.DosProtectionProfilesFloodIcmp{}.AttrType(), data)
}

// --- Unpacker for DosProtectionProfilesFloodIcmpRed ---
func unpackDosProtectionProfilesFloodIcmpRedToSdk(ctx context.Context, obj types.Object) (*security_services.DosProtectionProfilesFloodIcmpRed, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.DosProtectionProfilesFloodIcmpRed", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.DosProtectionProfilesFloodIcmpRed
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.DosProtectionProfilesFloodIcmpRed
	var d diag.Diagnostics
	// Handling Primitives
	if !model.ActivateRate.IsNull() && !model.ActivateRate.IsUnknown() {
		sdk.ActivateRate = int32(model.ActivateRate.ValueInt64())
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "ActivateRate", "value": sdk.ActivateRate})
	}

	// Handling Primitives
	if !model.AlarmRate.IsNull() && !model.AlarmRate.IsUnknown() {
		sdk.AlarmRate = int32(model.AlarmRate.ValueInt64())
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "AlarmRate", "value": sdk.AlarmRate})
	}

	// Handling Objects
	if !model.Block.IsNull() && !model.Block.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Block")
		unpacked, d := unpackDosProtectionProfilesFloodIcmpRedBlockToSdk(ctx, model.Block)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Block"})
		}
		if unpacked != nil {
			sdk.Block = unpacked
		}
	}

	// Handling Primitives
	if !model.MaximalRate.IsNull() && !model.MaximalRate.IsUnknown() {
		sdk.MaximalRate = int32(model.MaximalRate.ValueInt64())
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "MaximalRate", "value": sdk.MaximalRate})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.DosProtectionProfilesFloodIcmpRed", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for DosProtectionProfilesFloodIcmpRed ---
func packDosProtectionProfilesFloodIcmpRedFromSdk(ctx context.Context, sdk security_services.DosProtectionProfilesFloodIcmpRed) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.DosProtectionProfilesFloodIcmpRed", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.DosProtectionProfilesFloodIcmpRed
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.ActivateRate = basetypes.NewInt64Value(int64(sdk.ActivateRate))
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "ActivateRate", "value": sdk.ActivateRate})
	// Handling Primitives
	// Standard primitive packing
	model.AlarmRate = basetypes.NewInt64Value(int64(sdk.AlarmRate))
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "AlarmRate", "value": sdk.AlarmRate})
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Block != nil {
		tflog.Debug(ctx, "Packing nested object for field Block")
		packed, d := packDosProtectionProfilesFloodIcmpRedBlockFromSdk(ctx, *sdk.Block)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Block"})
		}
		model.Block = packed
	} else {
		model.Block = basetypes.NewObjectNull(models.DosProtectionProfilesFloodIcmpRedBlock{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	model.MaximalRate = basetypes.NewInt64Value(int64(sdk.MaximalRate))
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "MaximalRate", "value": sdk.MaximalRate})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.DosProtectionProfilesFloodIcmpRed{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.DosProtectionProfilesFloodIcmpRed", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for DosProtectionProfilesFloodIcmpRed ---
func unpackDosProtectionProfilesFloodIcmpRedListToSdk(ctx context.Context, list types.List) ([]security_services.DosProtectionProfilesFloodIcmpRed, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.DosProtectionProfilesFloodIcmpRed")
	diags := diag.Diagnostics{}
	var data []models.DosProtectionProfilesFloodIcmpRed
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.DosProtectionProfilesFloodIcmpRed, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.DosProtectionProfilesFloodIcmpRed{}.AttrTypes(), &item)
		unpacked, d := unpackDosProtectionProfilesFloodIcmpRedToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.DosProtectionProfilesFloodIcmpRed", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for DosProtectionProfilesFloodIcmpRed ---
func packDosProtectionProfilesFloodIcmpRedListFromSdk(ctx context.Context, sdks []security_services.DosProtectionProfilesFloodIcmpRed) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.DosProtectionProfilesFloodIcmpRed")
	diags := diag.Diagnostics{}
	var data []models.DosProtectionProfilesFloodIcmpRed

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.DosProtectionProfilesFloodIcmpRed
		obj, d := packDosProtectionProfilesFloodIcmpRedFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.DosProtectionProfilesFloodIcmpRed{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.DosProtectionProfilesFloodIcmpRed", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.DosProtectionProfilesFloodIcmpRed{}.AttrType(), data)
}

// --- Unpacker for DosProtectionProfilesFloodIcmpRedBlock ---
func unpackDosProtectionProfilesFloodIcmpRedBlockToSdk(ctx context.Context, obj types.Object) (*security_services.DosProtectionProfilesFloodIcmpRedBlock, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.DosProtectionProfilesFloodIcmpRedBlock", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.DosProtectionProfilesFloodIcmpRedBlock
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.DosProtectionProfilesFloodIcmpRedBlock
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Duration.IsNull() && !model.Duration.IsUnknown() {
		val := int32(model.Duration.ValueInt64())
		sdk.Duration = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Duration", "value": *sdk.Duration})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.DosProtectionProfilesFloodIcmpRedBlock", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for DosProtectionProfilesFloodIcmpRedBlock ---
func packDosProtectionProfilesFloodIcmpRedBlockFromSdk(ctx context.Context, sdk security_services.DosProtectionProfilesFloodIcmpRedBlock) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.DosProtectionProfilesFloodIcmpRedBlock", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.DosProtectionProfilesFloodIcmpRedBlock
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Duration != nil {
		model.Duration = basetypes.NewInt64Value(int64(*sdk.Duration))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Duration", "value": *sdk.Duration})
	} else {
		model.Duration = basetypes.NewInt64Null()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.DosProtectionProfilesFloodIcmpRedBlock{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.DosProtectionProfilesFloodIcmpRedBlock", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for DosProtectionProfilesFloodIcmpRedBlock ---
func unpackDosProtectionProfilesFloodIcmpRedBlockListToSdk(ctx context.Context, list types.List) ([]security_services.DosProtectionProfilesFloodIcmpRedBlock, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.DosProtectionProfilesFloodIcmpRedBlock")
	diags := diag.Diagnostics{}
	var data []models.DosProtectionProfilesFloodIcmpRedBlock
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.DosProtectionProfilesFloodIcmpRedBlock, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.DosProtectionProfilesFloodIcmpRedBlock{}.AttrTypes(), &item)
		unpacked, d := unpackDosProtectionProfilesFloodIcmpRedBlockToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.DosProtectionProfilesFloodIcmpRedBlock", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for DosProtectionProfilesFloodIcmpRedBlock ---
func packDosProtectionProfilesFloodIcmpRedBlockListFromSdk(ctx context.Context, sdks []security_services.DosProtectionProfilesFloodIcmpRedBlock) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.DosProtectionProfilesFloodIcmpRedBlock")
	diags := diag.Diagnostics{}
	var data []models.DosProtectionProfilesFloodIcmpRedBlock

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.DosProtectionProfilesFloodIcmpRedBlock
		obj, d := packDosProtectionProfilesFloodIcmpRedBlockFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.DosProtectionProfilesFloodIcmpRedBlock{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.DosProtectionProfilesFloodIcmpRedBlock", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.DosProtectionProfilesFloodIcmpRedBlock{}.AttrType(), data)
}

// --- Unpacker for DosProtectionProfilesFloodTcpSyn ---
func unpackDosProtectionProfilesFloodTcpSynToSdk(ctx context.Context, obj types.Object) (*security_services.DosProtectionProfilesFloodTcpSyn, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.DosProtectionProfilesFloodTcpSyn", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.DosProtectionProfilesFloodTcpSyn
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.DosProtectionProfilesFloodTcpSyn
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Enable.IsNull() && !model.Enable.IsUnknown() {
		sdk.Enable = model.Enable.ValueBool()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Enable", "value": sdk.Enable})
	}

	// Handling Objects
	if !model.Red.IsNull() && !model.Red.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Red")
		unpacked, d := unpackDosProtectionProfilesFloodIcmpRedToSdk(ctx, model.Red)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Red"})
		}
		if unpacked != nil {
			sdk.Red = unpacked
		}
	}

	// Handling Objects
	if !model.SynCookies.IsNull() && !model.SynCookies.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field SynCookies")
		unpacked, d := unpackDosProtectionProfilesFloodTcpSynSynCookiesToSdk(ctx, model.SynCookies)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "SynCookies"})
		}
		if unpacked != nil {
			sdk.SynCookies = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.DosProtectionProfilesFloodTcpSyn", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for DosProtectionProfilesFloodTcpSyn ---
func packDosProtectionProfilesFloodTcpSynFromSdk(ctx context.Context, sdk security_services.DosProtectionProfilesFloodTcpSyn) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.DosProtectionProfilesFloodTcpSyn", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.DosProtectionProfilesFloodTcpSyn
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.Enable = basetypes.NewBoolValue(sdk.Enable)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Enable", "value": sdk.Enable})
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Red != nil {
		tflog.Debug(ctx, "Packing nested object for field Red")
		packed, d := packDosProtectionProfilesFloodIcmpRedFromSdk(ctx, *sdk.Red)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Red"})
		}
		model.Red = packed
	} else {
		model.Red = basetypes.NewObjectNull(models.DosProtectionProfilesFloodIcmpRed{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.SynCookies != nil {
		tflog.Debug(ctx, "Packing nested object for field SynCookies")
		packed, d := packDosProtectionProfilesFloodTcpSynSynCookiesFromSdk(ctx, *sdk.SynCookies)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "SynCookies"})
		}
		model.SynCookies = packed
	} else {
		model.SynCookies = basetypes.NewObjectNull(models.DosProtectionProfilesFloodTcpSynSynCookies{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.DosProtectionProfilesFloodTcpSyn{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.DosProtectionProfilesFloodTcpSyn", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for DosProtectionProfilesFloodTcpSyn ---
func unpackDosProtectionProfilesFloodTcpSynListToSdk(ctx context.Context, list types.List) ([]security_services.DosProtectionProfilesFloodTcpSyn, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.DosProtectionProfilesFloodTcpSyn")
	diags := diag.Diagnostics{}
	var data []models.DosProtectionProfilesFloodTcpSyn
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.DosProtectionProfilesFloodTcpSyn, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.DosProtectionProfilesFloodTcpSyn{}.AttrTypes(), &item)
		unpacked, d := unpackDosProtectionProfilesFloodTcpSynToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.DosProtectionProfilesFloodTcpSyn", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for DosProtectionProfilesFloodTcpSyn ---
func packDosProtectionProfilesFloodTcpSynListFromSdk(ctx context.Context, sdks []security_services.DosProtectionProfilesFloodTcpSyn) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.DosProtectionProfilesFloodTcpSyn")
	diags := diag.Diagnostics{}
	var data []models.DosProtectionProfilesFloodTcpSyn

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.DosProtectionProfilesFloodTcpSyn
		obj, d := packDosProtectionProfilesFloodTcpSynFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.DosProtectionProfilesFloodTcpSyn{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.DosProtectionProfilesFloodTcpSyn", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.DosProtectionProfilesFloodTcpSyn{}.AttrType(), data)
}

// --- Unpacker for DosProtectionProfilesFloodTcpSynSynCookies ---
func unpackDosProtectionProfilesFloodTcpSynSynCookiesToSdk(ctx context.Context, obj types.Object) (*security_services.DosProtectionProfilesFloodTcpSynSynCookies, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.DosProtectionProfilesFloodTcpSynSynCookies", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.DosProtectionProfilesFloodTcpSynSynCookies
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.DosProtectionProfilesFloodTcpSynSynCookies
	var d diag.Diagnostics
	// Handling Primitives
	if !model.ActivateRate.IsNull() && !model.ActivateRate.IsUnknown() {
		sdk.ActivateRate = int32(model.ActivateRate.ValueInt64())
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "ActivateRate", "value": sdk.ActivateRate})
	}

	// Handling Primitives
	if !model.AlarmRate.IsNull() && !model.AlarmRate.IsUnknown() {
		sdk.AlarmRate = int32(model.AlarmRate.ValueInt64())
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "AlarmRate", "value": sdk.AlarmRate})
	}

	// Handling Objects
	if !model.Block.IsNull() && !model.Block.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Block")
		unpacked, d := unpackDosProtectionProfilesFloodTcpSynSynCookiesBlockToSdk(ctx, model.Block)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Block"})
		}
		if unpacked != nil {
			sdk.Block = unpacked
		}
	}

	// Handling Primitives
	if !model.MaximalRate.IsNull() && !model.MaximalRate.IsUnknown() {
		sdk.MaximalRate = int32(model.MaximalRate.ValueInt64())
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "MaximalRate", "value": sdk.MaximalRate})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.DosProtectionProfilesFloodTcpSynSynCookies", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for DosProtectionProfilesFloodTcpSynSynCookies ---
func packDosProtectionProfilesFloodTcpSynSynCookiesFromSdk(ctx context.Context, sdk security_services.DosProtectionProfilesFloodTcpSynSynCookies) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.DosProtectionProfilesFloodTcpSynSynCookies", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.DosProtectionProfilesFloodTcpSynSynCookies
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.ActivateRate = basetypes.NewInt64Value(int64(sdk.ActivateRate))
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "ActivateRate", "value": sdk.ActivateRate})
	// Handling Primitives
	// Standard primitive packing
	model.AlarmRate = basetypes.NewInt64Value(int64(sdk.AlarmRate))
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "AlarmRate", "value": sdk.AlarmRate})
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Block != nil {
		tflog.Debug(ctx, "Packing nested object for field Block")
		packed, d := packDosProtectionProfilesFloodTcpSynSynCookiesBlockFromSdk(ctx, *sdk.Block)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Block"})
		}
		model.Block = packed
	} else {
		model.Block = basetypes.NewObjectNull(models.DosProtectionProfilesFloodTcpSynSynCookiesBlock{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	model.MaximalRate = basetypes.NewInt64Value(int64(sdk.MaximalRate))
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "MaximalRate", "value": sdk.MaximalRate})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.DosProtectionProfilesFloodTcpSynSynCookies{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.DosProtectionProfilesFloodTcpSynSynCookies", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for DosProtectionProfilesFloodTcpSynSynCookies ---
func unpackDosProtectionProfilesFloodTcpSynSynCookiesListToSdk(ctx context.Context, list types.List) ([]security_services.DosProtectionProfilesFloodTcpSynSynCookies, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.DosProtectionProfilesFloodTcpSynSynCookies")
	diags := diag.Diagnostics{}
	var data []models.DosProtectionProfilesFloodTcpSynSynCookies
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.DosProtectionProfilesFloodTcpSynSynCookies, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.DosProtectionProfilesFloodTcpSynSynCookies{}.AttrTypes(), &item)
		unpacked, d := unpackDosProtectionProfilesFloodTcpSynSynCookiesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.DosProtectionProfilesFloodTcpSynSynCookies", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for DosProtectionProfilesFloodTcpSynSynCookies ---
func packDosProtectionProfilesFloodTcpSynSynCookiesListFromSdk(ctx context.Context, sdks []security_services.DosProtectionProfilesFloodTcpSynSynCookies) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.DosProtectionProfilesFloodTcpSynSynCookies")
	diags := diag.Diagnostics{}
	var data []models.DosProtectionProfilesFloodTcpSynSynCookies

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.DosProtectionProfilesFloodTcpSynSynCookies
		obj, d := packDosProtectionProfilesFloodTcpSynSynCookiesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.DosProtectionProfilesFloodTcpSynSynCookies{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.DosProtectionProfilesFloodTcpSynSynCookies", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.DosProtectionProfilesFloodTcpSynSynCookies{}.AttrType(), data)
}

// --- Unpacker for DosProtectionProfilesFloodTcpSynSynCookiesBlock ---
func unpackDosProtectionProfilesFloodTcpSynSynCookiesBlockToSdk(ctx context.Context, obj types.Object) (*security_services.DosProtectionProfilesFloodTcpSynSynCookiesBlock, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.DosProtectionProfilesFloodTcpSynSynCookiesBlock", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.DosProtectionProfilesFloodTcpSynSynCookiesBlock
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.DosProtectionProfilesFloodTcpSynSynCookiesBlock
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Duration.IsNull() && !model.Duration.IsUnknown() {
		val := int32(model.Duration.ValueInt64())
		sdk.Duration = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Duration", "value": *sdk.Duration})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.DosProtectionProfilesFloodTcpSynSynCookiesBlock", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for DosProtectionProfilesFloodTcpSynSynCookiesBlock ---
func packDosProtectionProfilesFloodTcpSynSynCookiesBlockFromSdk(ctx context.Context, sdk security_services.DosProtectionProfilesFloodTcpSynSynCookiesBlock) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.DosProtectionProfilesFloodTcpSynSynCookiesBlock", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.DosProtectionProfilesFloodTcpSynSynCookiesBlock
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Duration != nil {
		model.Duration = basetypes.NewInt64Value(int64(*sdk.Duration))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Duration", "value": *sdk.Duration})
	} else {
		model.Duration = basetypes.NewInt64Null()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.DosProtectionProfilesFloodTcpSynSynCookiesBlock{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.DosProtectionProfilesFloodTcpSynSynCookiesBlock", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for DosProtectionProfilesFloodTcpSynSynCookiesBlock ---
func unpackDosProtectionProfilesFloodTcpSynSynCookiesBlockListToSdk(ctx context.Context, list types.List) ([]security_services.DosProtectionProfilesFloodTcpSynSynCookiesBlock, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.DosProtectionProfilesFloodTcpSynSynCookiesBlock")
	diags := diag.Diagnostics{}
	var data []models.DosProtectionProfilesFloodTcpSynSynCookiesBlock
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.DosProtectionProfilesFloodTcpSynSynCookiesBlock, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.DosProtectionProfilesFloodTcpSynSynCookiesBlock{}.AttrTypes(), &item)
		unpacked, d := unpackDosProtectionProfilesFloodTcpSynSynCookiesBlockToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.DosProtectionProfilesFloodTcpSynSynCookiesBlock", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for DosProtectionProfilesFloodTcpSynSynCookiesBlock ---
func packDosProtectionProfilesFloodTcpSynSynCookiesBlockListFromSdk(ctx context.Context, sdks []security_services.DosProtectionProfilesFloodTcpSynSynCookiesBlock) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.DosProtectionProfilesFloodTcpSynSynCookiesBlock")
	diags := diag.Diagnostics{}
	var data []models.DosProtectionProfilesFloodTcpSynSynCookiesBlock

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.DosProtectionProfilesFloodTcpSynSynCookiesBlock
		obj, d := packDosProtectionProfilesFloodTcpSynSynCookiesBlockFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.DosProtectionProfilesFloodTcpSynSynCookiesBlock{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.DosProtectionProfilesFloodTcpSynSynCookiesBlock", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.DosProtectionProfilesFloodTcpSynSynCookiesBlock{}.AttrType(), data)
}

// --- Unpacker for DosProtectionProfilesResource ---
func unpackDosProtectionProfilesResourceToSdk(ctx context.Context, obj types.Object) (*security_services.DosProtectionProfilesResource, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.DosProtectionProfilesResource", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.DosProtectionProfilesResource
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.DosProtectionProfilesResource
	var d diag.Diagnostics
	// Handling Objects
	if !model.Sessions.IsNull() && !model.Sessions.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Sessions")
		unpacked, d := unpackDosProtectionProfilesResourceSessionsToSdk(ctx, model.Sessions)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Sessions"})
		}
		if unpacked != nil {
			sdk.Sessions = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.DosProtectionProfilesResource", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for DosProtectionProfilesResource ---
func packDosProtectionProfilesResourceFromSdk(ctx context.Context, sdk security_services.DosProtectionProfilesResource) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.DosProtectionProfilesResource", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.DosProtectionProfilesResource
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Sessions != nil {
		tflog.Debug(ctx, "Packing nested object for field Sessions")
		packed, d := packDosProtectionProfilesResourceSessionsFromSdk(ctx, *sdk.Sessions)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Sessions"})
		}
		model.Sessions = packed
	} else {
		model.Sessions = basetypes.NewObjectNull(models.DosProtectionProfilesResourceSessions{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.DosProtectionProfilesResource{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.DosProtectionProfilesResource", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for DosProtectionProfilesResource ---
func unpackDosProtectionProfilesResourceListToSdk(ctx context.Context, list types.List) ([]security_services.DosProtectionProfilesResource, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.DosProtectionProfilesResource")
	diags := diag.Diagnostics{}
	var data []models.DosProtectionProfilesResource
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.DosProtectionProfilesResource, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.DosProtectionProfilesResource{}.AttrTypes(), &item)
		unpacked, d := unpackDosProtectionProfilesResourceToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.DosProtectionProfilesResource", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for DosProtectionProfilesResource ---
func packDosProtectionProfilesResourceListFromSdk(ctx context.Context, sdks []security_services.DosProtectionProfilesResource) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.DosProtectionProfilesResource")
	diags := diag.Diagnostics{}
	var data []models.DosProtectionProfilesResource

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.DosProtectionProfilesResource
		obj, d := packDosProtectionProfilesResourceFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.DosProtectionProfilesResource{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.DosProtectionProfilesResource", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.DosProtectionProfilesResource{}.AttrType(), data)
}

// --- Unpacker for DosProtectionProfilesResourceSessions ---
func unpackDosProtectionProfilesResourceSessionsToSdk(ctx context.Context, obj types.Object) (*security_services.DosProtectionProfilesResourceSessions, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.DosProtectionProfilesResourceSessions", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.DosProtectionProfilesResourceSessions
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.DosProtectionProfilesResourceSessions
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Enabled.IsNull() && !model.Enabled.IsUnknown() {
		sdk.Enabled = model.Enabled.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Enabled", "value": *sdk.Enabled})
	}

	// Handling Primitives
	if !model.MaxConcurrentLimit.IsNull() && !model.MaxConcurrentLimit.IsUnknown() {
		val := int32(model.MaxConcurrentLimit.ValueInt64())
		sdk.MaxConcurrentLimit = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "MaxConcurrentLimit", "value": *sdk.MaxConcurrentLimit})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.DosProtectionProfilesResourceSessions", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for DosProtectionProfilesResourceSessions ---
func packDosProtectionProfilesResourceSessionsFromSdk(ctx context.Context, sdk security_services.DosProtectionProfilesResourceSessions) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.DosProtectionProfilesResourceSessions", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.DosProtectionProfilesResourceSessions
	var d diag.Diagnostics
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
	if sdk.MaxConcurrentLimit != nil {
		model.MaxConcurrentLimit = basetypes.NewInt64Value(int64(*sdk.MaxConcurrentLimit))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "MaxConcurrentLimit", "value": *sdk.MaxConcurrentLimit})
	} else {
		model.MaxConcurrentLimit = basetypes.NewInt64Null()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.DosProtectionProfilesResourceSessions{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.DosProtectionProfilesResourceSessions", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for DosProtectionProfilesResourceSessions ---
func unpackDosProtectionProfilesResourceSessionsListToSdk(ctx context.Context, list types.List) ([]security_services.DosProtectionProfilesResourceSessions, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.DosProtectionProfilesResourceSessions")
	diags := diag.Diagnostics{}
	var data []models.DosProtectionProfilesResourceSessions
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.DosProtectionProfilesResourceSessions, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.DosProtectionProfilesResourceSessions{}.AttrTypes(), &item)
		unpacked, d := unpackDosProtectionProfilesResourceSessionsToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.DosProtectionProfilesResourceSessions", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for DosProtectionProfilesResourceSessions ---
func packDosProtectionProfilesResourceSessionsListFromSdk(ctx context.Context, sdks []security_services.DosProtectionProfilesResourceSessions) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.DosProtectionProfilesResourceSessions")
	diags := diag.Diagnostics{}
	var data []models.DosProtectionProfilesResourceSessions

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.DosProtectionProfilesResourceSessions
		obj, d := packDosProtectionProfilesResourceSessionsFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.DosProtectionProfilesResourceSessions{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.DosProtectionProfilesResourceSessions", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.DosProtectionProfilesResourceSessions{}.AttrType(), data)
}
