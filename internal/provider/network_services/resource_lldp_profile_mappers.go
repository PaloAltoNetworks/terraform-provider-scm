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

// --- Unpacker for LldpProfiles ---
func unpackLldpProfilesToSdk(ctx context.Context, obj types.Object) (*network_services.LldpProfiles, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.LldpProfiles", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.LldpProfiles
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.LldpProfiles
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
	if !model.Mode.IsNull() && !model.Mode.IsUnknown() {
		sdk.Mode = model.Mode.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Mode", "value": *sdk.Mode})
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	// Handling Objects
	if !model.OptionTlvs.IsNull() && !model.OptionTlvs.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field OptionTlvs")
		unpacked, d := unpackLldpProfilesOptionTlvsToSdk(ctx, model.OptionTlvs)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "OptionTlvs"})
		}
		if unpacked != nil {
			sdk.OptionTlvs = unpacked
		}
	}

	// Handling Primitives
	if !model.Snippet.IsNull() && !model.Snippet.IsUnknown() {
		sdk.Snippet = model.Snippet.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	}

	// Handling Primitives
	if !model.SnmpSyslogNotification.IsNull() && !model.SnmpSyslogNotification.IsUnknown() {
		sdk.SnmpSyslogNotification = model.SnmpSyslogNotification.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "SnmpSyslogNotification", "value": *sdk.SnmpSyslogNotification})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.LldpProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for LldpProfiles ---
func packLldpProfilesFromSdk(ctx context.Context, sdk network_services.LldpProfiles) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.LldpProfiles", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.LldpProfiles
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
	if sdk.Mode != nil {
		model.Mode = basetypes.NewStringValue(*sdk.Mode)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Mode", "value": *sdk.Mode})
	} else {
		model.Mode = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.OptionTlvs != nil {
		tflog.Debug(ctx, "Packing nested object for field OptionTlvs")
		packed, d := packLldpProfilesOptionTlvsFromSdk(ctx, *sdk.OptionTlvs)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "OptionTlvs"})
		}
		model.OptionTlvs = packed
	} else {
		model.OptionTlvs = basetypes.NewObjectNull(models.LldpProfilesOptionTlvs{}.AttrTypes())
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
	if sdk.SnmpSyslogNotification != nil {
		model.SnmpSyslogNotification = basetypes.NewBoolValue(*sdk.SnmpSyslogNotification)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "SnmpSyslogNotification", "value": *sdk.SnmpSyslogNotification})
	} else {
		model.SnmpSyslogNotification = basetypes.NewBoolNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.LldpProfiles{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.LldpProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for LldpProfiles ---
func unpackLldpProfilesListToSdk(ctx context.Context, list types.List) ([]network_services.LldpProfiles, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.LldpProfiles")
	diags := diag.Diagnostics{}
	var data []models.LldpProfiles
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.LldpProfiles, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.LldpProfiles{}.AttrTypes(), &item)
		unpacked, d := unpackLldpProfilesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.LldpProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for LldpProfiles ---
func packLldpProfilesListFromSdk(ctx context.Context, sdks []network_services.LldpProfiles) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.LldpProfiles")
	diags := diag.Diagnostics{}
	var data []models.LldpProfiles

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.LldpProfiles
		obj, d := packLldpProfilesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.LldpProfiles{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.LldpProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.LldpProfiles{}.AttrType(), data)
}

// --- Unpacker for LldpProfilesOptionTlvs ---
func unpackLldpProfilesOptionTlvsToSdk(ctx context.Context, obj types.Object) (*network_services.LldpProfilesOptionTlvs, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.LldpProfilesOptionTlvs", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.LldpProfilesOptionTlvs
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.LldpProfilesOptionTlvs
	var d diag.Diagnostics
	// Handling Objects
	if !model.ManagementAddress.IsNull() && !model.ManagementAddress.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field ManagementAddress")
		unpacked, d := unpackLldpProfilesOptionTlvsManagementAddressToSdk(ctx, model.ManagementAddress)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "ManagementAddress"})
		}
		if unpacked != nil {
			sdk.ManagementAddress = unpacked
		}
	}

	// Handling Primitives
	if !model.PortDescription.IsNull() && !model.PortDescription.IsUnknown() {
		sdk.PortDescription = model.PortDescription.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "PortDescription", "value": *sdk.PortDescription})
	}

	// Handling Primitives
	if !model.SystemCapabilities.IsNull() && !model.SystemCapabilities.IsUnknown() {
		sdk.SystemCapabilities = model.SystemCapabilities.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "SystemCapabilities", "value": *sdk.SystemCapabilities})
	}

	// Handling Primitives
	if !model.SystemDescription.IsNull() && !model.SystemDescription.IsUnknown() {
		sdk.SystemDescription = model.SystemDescription.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "SystemDescription", "value": *sdk.SystemDescription})
	}

	// Handling Primitives
	if !model.SystemName.IsNull() && !model.SystemName.IsUnknown() {
		sdk.SystemName = model.SystemName.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "SystemName", "value": *sdk.SystemName})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.LldpProfilesOptionTlvs", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for LldpProfilesOptionTlvs ---
func packLldpProfilesOptionTlvsFromSdk(ctx context.Context, sdk network_services.LldpProfilesOptionTlvs) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.LldpProfilesOptionTlvs", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.LldpProfilesOptionTlvs
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.ManagementAddress != nil {
		tflog.Debug(ctx, "Packing nested object for field ManagementAddress")
		packed, d := packLldpProfilesOptionTlvsManagementAddressFromSdk(ctx, *sdk.ManagementAddress)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "ManagementAddress"})
		}
		model.ManagementAddress = packed
	} else {
		model.ManagementAddress = basetypes.NewObjectNull(models.LldpProfilesOptionTlvsManagementAddress{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.PortDescription != nil {
		model.PortDescription = basetypes.NewBoolValue(*sdk.PortDescription)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "PortDescription", "value": *sdk.PortDescription})
	} else {
		model.PortDescription = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.SystemCapabilities != nil {
		model.SystemCapabilities = basetypes.NewBoolValue(*sdk.SystemCapabilities)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "SystemCapabilities", "value": *sdk.SystemCapabilities})
	} else {
		model.SystemCapabilities = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.SystemDescription != nil {
		model.SystemDescription = basetypes.NewBoolValue(*sdk.SystemDescription)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "SystemDescription", "value": *sdk.SystemDescription})
	} else {
		model.SystemDescription = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.SystemName != nil {
		model.SystemName = basetypes.NewBoolValue(*sdk.SystemName)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "SystemName", "value": *sdk.SystemName})
	} else {
		model.SystemName = basetypes.NewBoolNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.LldpProfilesOptionTlvs{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.LldpProfilesOptionTlvs", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for LldpProfilesOptionTlvs ---
func unpackLldpProfilesOptionTlvsListToSdk(ctx context.Context, list types.List) ([]network_services.LldpProfilesOptionTlvs, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.LldpProfilesOptionTlvs")
	diags := diag.Diagnostics{}
	var data []models.LldpProfilesOptionTlvs
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.LldpProfilesOptionTlvs, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.LldpProfilesOptionTlvs{}.AttrTypes(), &item)
		unpacked, d := unpackLldpProfilesOptionTlvsToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.LldpProfilesOptionTlvs", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for LldpProfilesOptionTlvs ---
func packLldpProfilesOptionTlvsListFromSdk(ctx context.Context, sdks []network_services.LldpProfilesOptionTlvs) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.LldpProfilesOptionTlvs")
	diags := diag.Diagnostics{}
	var data []models.LldpProfilesOptionTlvs

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.LldpProfilesOptionTlvs
		obj, d := packLldpProfilesOptionTlvsFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.LldpProfilesOptionTlvs{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.LldpProfilesOptionTlvs", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.LldpProfilesOptionTlvs{}.AttrType(), data)
}

// --- Unpacker for LldpProfilesOptionTlvsManagementAddress ---
func unpackLldpProfilesOptionTlvsManagementAddressToSdk(ctx context.Context, obj types.Object) (*network_services.LldpProfilesOptionTlvsManagementAddress, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.LldpProfilesOptionTlvsManagementAddress", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.LldpProfilesOptionTlvsManagementAddress
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.LldpProfilesOptionTlvsManagementAddress
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Enabled.IsNull() && !model.Enabled.IsUnknown() {
		sdk.Enabled = model.Enabled.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Enabled", "value": *sdk.Enabled})
	}

	// Handling Lists
	if !model.Iplist.IsNull() && !model.Iplist.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field Iplist")
		unpacked, d := unpackLldpProfilesOptionTlvsManagementAddressIplistInnerListToSdk(ctx, model.Iplist)
		diags.Append(d...)
		sdk.Iplist = unpacked
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.LldpProfilesOptionTlvsManagementAddress", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for LldpProfilesOptionTlvsManagementAddress ---
func packLldpProfilesOptionTlvsManagementAddressFromSdk(ctx context.Context, sdk network_services.LldpProfilesOptionTlvsManagementAddress) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.LldpProfilesOptionTlvsManagementAddress", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.LldpProfilesOptionTlvsManagementAddress
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Enabled != nil {
		model.Enabled = basetypes.NewBoolValue(*sdk.Enabled)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Enabled", "value": *sdk.Enabled})
	} else {
		model.Enabled = basetypes.NewBoolNull()
	}
	// Handling Lists
	if sdk.Iplist != nil {
		tflog.Debug(ctx, "Packing list of objects for field Iplist")
		packed, d := packLldpProfilesOptionTlvsManagementAddressIplistInnerListFromSdk(ctx, sdk.Iplist)
		diags.Append(d...)
		model.Iplist = packed
	} else {
		model.Iplist = basetypes.NewListNull(models.LldpProfilesOptionTlvsManagementAddressIplistInner{}.AttrType())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.LldpProfilesOptionTlvsManagementAddress{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.LldpProfilesOptionTlvsManagementAddress", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for LldpProfilesOptionTlvsManagementAddress ---
func unpackLldpProfilesOptionTlvsManagementAddressListToSdk(ctx context.Context, list types.List) ([]network_services.LldpProfilesOptionTlvsManagementAddress, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.LldpProfilesOptionTlvsManagementAddress")
	diags := diag.Diagnostics{}
	var data []models.LldpProfilesOptionTlvsManagementAddress
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.LldpProfilesOptionTlvsManagementAddress, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.LldpProfilesOptionTlvsManagementAddress{}.AttrTypes(), &item)
		unpacked, d := unpackLldpProfilesOptionTlvsManagementAddressToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.LldpProfilesOptionTlvsManagementAddress", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for LldpProfilesOptionTlvsManagementAddress ---
func packLldpProfilesOptionTlvsManagementAddressListFromSdk(ctx context.Context, sdks []network_services.LldpProfilesOptionTlvsManagementAddress) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.LldpProfilesOptionTlvsManagementAddress")
	diags := diag.Diagnostics{}
	var data []models.LldpProfilesOptionTlvsManagementAddress

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.LldpProfilesOptionTlvsManagementAddress
		obj, d := packLldpProfilesOptionTlvsManagementAddressFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.LldpProfilesOptionTlvsManagementAddress{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.LldpProfilesOptionTlvsManagementAddress", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.LldpProfilesOptionTlvsManagementAddress{}.AttrType(), data)
}

// --- Unpacker for LldpProfilesOptionTlvsManagementAddressIplistInner ---
func unpackLldpProfilesOptionTlvsManagementAddressIplistInnerToSdk(ctx context.Context, obj types.Object) (*network_services.LldpProfilesOptionTlvsManagementAddressIplistInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.LldpProfilesOptionTlvsManagementAddressIplistInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.LldpProfilesOptionTlvsManagementAddressIplistInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.LldpProfilesOptionTlvsManagementAddressIplistInner
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Interface.IsNull() && !model.Interface.IsUnknown() {
		sdk.Interface = model.Interface.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Interface", "value": *sdk.Interface})
	}

	// Handling Primitives
	if !model.Ipv4.IsNull() && !model.Ipv4.IsUnknown() {
		sdk.Ipv4 = model.Ipv4.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Ipv4", "value": *sdk.Ipv4})
	}

	// Handling Primitives
	if !model.Ipv6.IsNull() && !model.Ipv6.IsUnknown() {
		sdk.Ipv6 = model.Ipv6.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Ipv6", "value": *sdk.Ipv6})
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.LldpProfilesOptionTlvsManagementAddressIplistInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for LldpProfilesOptionTlvsManagementAddressIplistInner ---
func packLldpProfilesOptionTlvsManagementAddressIplistInnerFromSdk(ctx context.Context, sdk network_services.LldpProfilesOptionTlvsManagementAddressIplistInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.LldpProfilesOptionTlvsManagementAddressIplistInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.LldpProfilesOptionTlvsManagementAddressIplistInner
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Interface != nil {
		model.Interface = basetypes.NewStringValue(*sdk.Interface)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Interface", "value": *sdk.Interface})
	} else {
		model.Interface = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Ipv4 != nil {
		model.Ipv4 = basetypes.NewStringValue(*sdk.Ipv4)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Ipv4", "value": *sdk.Ipv4})
	} else {
		model.Ipv4 = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Ipv6 != nil {
		model.Ipv6 = basetypes.NewStringValue(*sdk.Ipv6)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Ipv6", "value": *sdk.Ipv6})
	} else {
		model.Ipv6 = basetypes.NewStringNull()
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

	obj, d := types.ObjectValueFrom(ctx, models.LldpProfilesOptionTlvsManagementAddressIplistInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.LldpProfilesOptionTlvsManagementAddressIplistInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for LldpProfilesOptionTlvsManagementAddressIplistInner ---
func unpackLldpProfilesOptionTlvsManagementAddressIplistInnerListToSdk(ctx context.Context, list types.List) ([]network_services.LldpProfilesOptionTlvsManagementAddressIplistInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.LldpProfilesOptionTlvsManagementAddressIplistInner")
	diags := diag.Diagnostics{}
	var data []models.LldpProfilesOptionTlvsManagementAddressIplistInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.LldpProfilesOptionTlvsManagementAddressIplistInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.LldpProfilesOptionTlvsManagementAddressIplistInner{}.AttrTypes(), &item)
		unpacked, d := unpackLldpProfilesOptionTlvsManagementAddressIplistInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.LldpProfilesOptionTlvsManagementAddressIplistInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for LldpProfilesOptionTlvsManagementAddressIplistInner ---
func packLldpProfilesOptionTlvsManagementAddressIplistInnerListFromSdk(ctx context.Context, sdks []network_services.LldpProfilesOptionTlvsManagementAddressIplistInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.LldpProfilesOptionTlvsManagementAddressIplistInner")
	diags := diag.Diagnostics{}
	var data []models.LldpProfilesOptionTlvsManagementAddressIplistInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.LldpProfilesOptionTlvsManagementAddressIplistInner
		obj, d := packLldpProfilesOptionTlvsManagementAddressIplistInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.LldpProfilesOptionTlvsManagementAddressIplistInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.LldpProfilesOptionTlvsManagementAddressIplistInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.LldpProfilesOptionTlvsManagementAddressIplistInner{}.AttrType(), data)
}
