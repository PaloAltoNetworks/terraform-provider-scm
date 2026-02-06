package provider

import (
	"context"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/objects"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/objects"
)

// --- Unpacker for HipObjects ---
func unpackHipObjectsToSdk(ctx context.Context, obj types.Object) (*objects.HipObjects, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.HipObjects", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.HipObjects
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.HipObjects
	var d diag.Diagnostics

	// Handling Objects
	if !model.AntiMalware.IsNull() && !model.AntiMalware.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field AntiMalware")
		unpacked, d := unpackHipObjectsAntiMalwareToSdk(ctx, model.AntiMalware)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "AntiMalware"})
		}
		if unpacked != nil {
			sdk.AntiMalware = unpacked
		}
	}

	// Handling Objects
	if !model.Certificate.IsNull() && !model.Certificate.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Certificate")
		unpacked, d := unpackHipObjectsCertificateToSdk(ctx, model.Certificate)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Certificate"})
		}
		if unpacked != nil {
			sdk.Certificate = unpacked
		}
	}

	// Handling Objects
	if !model.CustomChecks.IsNull() && !model.CustomChecks.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field CustomChecks")
		unpacked, d := unpackHipObjectsCustomChecksToSdk(ctx, model.CustomChecks)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "CustomChecks"})
		}
		if unpacked != nil {
			sdk.CustomChecks = unpacked
		}
	}

	// Handling Objects
	if !model.DataLossPrevention.IsNull() && !model.DataLossPrevention.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field DataLossPrevention")
		unpacked, d := unpackHipObjectsDataLossPreventionToSdk(ctx, model.DataLossPrevention)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "DataLossPrevention"})
		}
		if unpacked != nil {
			sdk.DataLossPrevention = unpacked
		}
	}

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
	if !model.DiskBackup.IsNull() && !model.DiskBackup.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field DiskBackup")
		unpacked, d := unpackHipObjectsDiskBackupToSdk(ctx, model.DiskBackup)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "DiskBackup"})
		}
		if unpacked != nil {
			sdk.DiskBackup = unpacked
		}
	}

	// Handling Objects
	if !model.DiskEncryption.IsNull() && !model.DiskEncryption.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field DiskEncryption")
		unpacked, d := unpackHipObjectsDiskEncryptionToSdk(ctx, model.DiskEncryption)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "DiskEncryption"})
		}
		if unpacked != nil {
			sdk.DiskEncryption = unpacked
		}
	}

	// Handling Objects
	if !model.Firewall.IsNull() && !model.Firewall.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Firewall")
		unpacked, d := unpackHipObjectsFirewallToSdk(ctx, model.Firewall)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Firewall"})
		}
		if unpacked != nil {
			sdk.Firewall = unpacked
		}
	}

	// Handling Primitives
	if !model.Folder.IsNull() && !model.Folder.IsUnknown() {
		sdk.Folder = model.Folder.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Folder", "value": *sdk.Folder})
	}

	// Handling Objects
	if !model.HostInfo.IsNull() && !model.HostInfo.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field HostInfo")
		unpacked, d := unpackHipObjectsHostInfoToSdk(ctx, model.HostInfo)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "HostInfo"})
		}
		if unpacked != nil {
			sdk.HostInfo = unpacked
		}
	}

	// Handling Primitives
	if !model.Id.IsNull() && !model.Id.IsUnknown() {
		sdk.Id = model.Id.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Id", "value": sdk.Id})
	}

	// Handling Objects
	if !model.MobileDevice.IsNull() && !model.MobileDevice.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field MobileDevice")
		unpacked, d := unpackHipObjectsMobileDeviceToSdk(ctx, model.MobileDevice)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "MobileDevice"})
		}
		if unpacked != nil {
			sdk.MobileDevice = unpacked
		}
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	// Handling Objects
	if !model.NetworkInfo.IsNull() && !model.NetworkInfo.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field NetworkInfo")
		unpacked, d := unpackHipObjectsNetworkInfoToSdk(ctx, model.NetworkInfo)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "NetworkInfo"})
		}
		if unpacked != nil {
			sdk.NetworkInfo = unpacked
		}
	}

	// Handling Objects
	if !model.PatchManagement.IsNull() && !model.PatchManagement.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field PatchManagement")
		unpacked, d := unpackHipObjectsPatchManagementToSdk(ctx, model.PatchManagement)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "PatchManagement"})
		}
		if unpacked != nil {
			sdk.PatchManagement = unpacked
		}
	}

	// Handling Primitives
	if !model.Snippet.IsNull() && !model.Snippet.IsUnknown() {
		sdk.Snippet = model.Snippet.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.HipObjects", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for HipObjects ---
func packHipObjectsFromSdk(ctx context.Context, sdk objects.HipObjects) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.HipObjects", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.HipObjects
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.AntiMalware != nil {
		tflog.Debug(ctx, "Packing nested object for field AntiMalware")
		packed, d := packHipObjectsAntiMalwareFromSdk(ctx, *sdk.AntiMalware)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "AntiMalware"})
		}
		model.AntiMalware = packed
	} else {
		model.AntiMalware = basetypes.NewObjectNull(models.HipObjectsAntiMalware{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Certificate != nil {
		tflog.Debug(ctx, "Packing nested object for field Certificate")
		packed, d := packHipObjectsCertificateFromSdk(ctx, *sdk.Certificate)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Certificate"})
		}
		model.Certificate = packed
	} else {
		model.Certificate = basetypes.NewObjectNull(models.HipObjectsCertificate{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.CustomChecks != nil {
		tflog.Debug(ctx, "Packing nested object for field CustomChecks")
		packed, d := packHipObjectsCustomChecksFromSdk(ctx, *sdk.CustomChecks)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "CustomChecks"})
		}
		model.CustomChecks = packed
	} else {
		model.CustomChecks = basetypes.NewObjectNull(models.HipObjectsCustomChecks{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.DataLossPrevention != nil {
		tflog.Debug(ctx, "Packing nested object for field DataLossPrevention")
		packed, d := packHipObjectsDataLossPreventionFromSdk(ctx, *sdk.DataLossPrevention)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "DataLossPrevention"})
		}
		model.DataLossPrevention = packed
	} else {
		model.DataLossPrevention = basetypes.NewObjectNull(models.HipObjectsDataLossPrevention{}.AttrTypes())
	}
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
	if sdk.DiskBackup != nil {
		tflog.Debug(ctx, "Packing nested object for field DiskBackup")
		packed, d := packHipObjectsDiskBackupFromSdk(ctx, *sdk.DiskBackup)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "DiskBackup"})
		}
		model.DiskBackup = packed
	} else {
		model.DiskBackup = basetypes.NewObjectNull(models.HipObjectsDiskBackup{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.DiskEncryption != nil {
		tflog.Debug(ctx, "Packing nested object for field DiskEncryption")
		packed, d := packHipObjectsDiskEncryptionFromSdk(ctx, *sdk.DiskEncryption)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "DiskEncryption"})
		}
		model.DiskEncryption = packed
	} else {
		model.DiskEncryption = basetypes.NewObjectNull(models.HipObjectsDiskEncryption{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Firewall != nil {
		tflog.Debug(ctx, "Packing nested object for field Firewall")
		packed, d := packHipObjectsFirewallFromSdk(ctx, *sdk.Firewall)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Firewall"})
		}
		model.Firewall = packed
	} else {
		model.Firewall = basetypes.NewObjectNull(models.HipObjectsFirewall{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Folder != nil {
		model.Folder = basetypes.NewStringValue(*sdk.Folder)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Folder", "value": *sdk.Folder})
	} else {
		model.Folder = basetypes.NewStringNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.HostInfo != nil {
		tflog.Debug(ctx, "Packing nested object for field HostInfo")
		packed, d := packHipObjectsHostInfoFromSdk(ctx, *sdk.HostInfo)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "HostInfo"})
		}
		model.HostInfo = packed
	} else {
		model.HostInfo = basetypes.NewObjectNull(models.HipObjectsHostInfo{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	model.Id = basetypes.NewStringValue(sdk.Id)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Id", "value": sdk.Id})
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.MobileDevice != nil {
		tflog.Debug(ctx, "Packing nested object for field MobileDevice")
		packed, d := packHipObjectsMobileDeviceFromSdk(ctx, *sdk.MobileDevice)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "MobileDevice"})
		}
		model.MobileDevice = packed
	} else {
		model.MobileDevice = basetypes.NewObjectNull(models.HipObjectsMobileDevice{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.NetworkInfo != nil {
		tflog.Debug(ctx, "Packing nested object for field NetworkInfo")
		packed, d := packHipObjectsNetworkInfoFromSdk(ctx, *sdk.NetworkInfo)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "NetworkInfo"})
		}
		model.NetworkInfo = packed
	} else {
		model.NetworkInfo = basetypes.NewObjectNull(models.HipObjectsNetworkInfo{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.PatchManagement != nil {
		tflog.Debug(ctx, "Packing nested object for field PatchManagement")
		packed, d := packHipObjectsPatchManagementFromSdk(ctx, *sdk.PatchManagement)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "PatchManagement"})
		}
		model.PatchManagement = packed
	} else {
		model.PatchManagement = basetypes.NewObjectNull(models.HipObjectsPatchManagement{}.AttrTypes())
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

	obj, d := types.ObjectValueFrom(ctx, models.HipObjects{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.HipObjects", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for HipObjects ---
func unpackHipObjectsListToSdk(ctx context.Context, list types.List) ([]objects.HipObjects, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.HipObjects")
	diags := diag.Diagnostics{}
	var data []models.HipObjects
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.HipObjects, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.HipObjects{}.AttrTypes(), &item)
		unpacked, d := unpackHipObjectsToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.HipObjects", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for HipObjects ---
func packHipObjectsListFromSdk(ctx context.Context, sdks []objects.HipObjects) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.HipObjects")
	diags := diag.Diagnostics{}
	var data []models.HipObjects

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.HipObjects
		obj, d := packHipObjectsFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.HipObjects{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.HipObjects", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.HipObjects{}.AttrType(), data)
}

// --- Unpacker for HipObjectsAntiMalware ---
func unpackHipObjectsAntiMalwareToSdk(ctx context.Context, obj types.Object) (*objects.HipObjectsAntiMalware, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.HipObjectsAntiMalware", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.HipObjectsAntiMalware
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.HipObjectsAntiMalware
	var d diag.Diagnostics
	// Handling Objects
	if !model.Criteria.IsNull() && !model.Criteria.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Criteria")
		unpacked, d := unpackHipObjectsAntiMalwareCriteriaToSdk(ctx, model.Criteria)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Criteria"})
		}
		if unpacked != nil {
			sdk.Criteria = unpacked
		}
	}

	// Handling Primitives
	if !model.ExcludeVendor.IsNull() && !model.ExcludeVendor.IsUnknown() {
		sdk.ExcludeVendor = model.ExcludeVendor.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "ExcludeVendor", "value": *sdk.ExcludeVendor})
	}

	// Handling Lists
	if !model.Vendor.IsNull() && !model.Vendor.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field Vendor")
		unpacked, d := unpackHipObjectsAntiMalwareVendorInnerListToSdk(ctx, model.Vendor)
		diags.Append(d...)
		sdk.Vendor = unpacked
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.HipObjectsAntiMalware", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for HipObjectsAntiMalware ---
func packHipObjectsAntiMalwareFromSdk(ctx context.Context, sdk objects.HipObjectsAntiMalware) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.HipObjectsAntiMalware", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.HipObjectsAntiMalware
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Criteria != nil {
		tflog.Debug(ctx, "Packing nested object for field Criteria")
		packed, d := packHipObjectsAntiMalwareCriteriaFromSdk(ctx, *sdk.Criteria)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Criteria"})
		}
		model.Criteria = packed
	} else {
		model.Criteria = basetypes.NewObjectNull(models.HipObjectsAntiMalwareCriteria{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.ExcludeVendor != nil {
		model.ExcludeVendor = basetypes.NewBoolValue(*sdk.ExcludeVendor)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "ExcludeVendor", "value": *sdk.ExcludeVendor})
	} else {
		model.ExcludeVendor = basetypes.NewBoolNull()
	}
	// Handling Lists
	if sdk.Vendor != nil {
		tflog.Debug(ctx, "Packing list of objects for field Vendor")
		packed, d := packHipObjectsAntiMalwareVendorInnerListFromSdk(ctx, sdk.Vendor)
		diags.Append(d...)
		model.Vendor = packed
	} else {
		model.Vendor = basetypes.NewListNull(models.HipObjectsAntiMalwareVendorInner{}.AttrType())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.HipObjectsAntiMalware{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.HipObjectsAntiMalware", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for HipObjectsAntiMalware ---
func unpackHipObjectsAntiMalwareListToSdk(ctx context.Context, list types.List) ([]objects.HipObjectsAntiMalware, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.HipObjectsAntiMalware")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsAntiMalware
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.HipObjectsAntiMalware, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.HipObjectsAntiMalware{}.AttrTypes(), &item)
		unpacked, d := unpackHipObjectsAntiMalwareToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.HipObjectsAntiMalware", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for HipObjectsAntiMalware ---
func packHipObjectsAntiMalwareListFromSdk(ctx context.Context, sdks []objects.HipObjectsAntiMalware) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.HipObjectsAntiMalware")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsAntiMalware

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.HipObjectsAntiMalware
		obj, d := packHipObjectsAntiMalwareFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.HipObjectsAntiMalware{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.HipObjectsAntiMalware", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.HipObjectsAntiMalware{}.AttrType(), data)
}

// --- Unpacker for HipObjectsAntiMalwareCriteria ---
func unpackHipObjectsAntiMalwareCriteriaToSdk(ctx context.Context, obj types.Object) (*objects.HipObjectsAntiMalwareCriteria, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.HipObjectsAntiMalwareCriteria", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.HipObjectsAntiMalwareCriteria
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.HipObjectsAntiMalwareCriteria
	var d diag.Diagnostics
	// Handling Primitives
	if !model.IsInstalled.IsNull() && !model.IsInstalled.IsUnknown() {
		sdk.IsInstalled = model.IsInstalled.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "IsInstalled", "value": *sdk.IsInstalled})
	}

	// Handling Objects
	if !model.LastScanTime.IsNull() && !model.LastScanTime.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field LastScanTime")
		unpacked, d := unpackHipObjectsAntiMalwareCriteriaLastScanTimeToSdk(ctx, model.LastScanTime)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "LastScanTime"})
		}
		if unpacked != nil {
			sdk.LastScanTime = unpacked
		}
	}

	// Handling Objects
	if !model.ProductVersion.IsNull() && !model.ProductVersion.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field ProductVersion")
		unpacked, d := unpackHipObjectsAntiMalwareCriteriaProductVersionToSdk(ctx, model.ProductVersion)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "ProductVersion"})
		}
		if unpacked != nil {
			sdk.ProductVersion = unpacked
		}
	}

	// Handling Primitives
	if !model.RealTimeProtection.IsNull() && !model.RealTimeProtection.IsUnknown() {
		sdk.RealTimeProtection = model.RealTimeProtection.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "RealTimeProtection", "value": *sdk.RealTimeProtection})
	}

	// Handling Objects
	if !model.VirdefVersion.IsNull() && !model.VirdefVersion.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field VirdefVersion")
		unpacked, d := unpackHipObjectsAntiMalwareCriteriaVirdefVersionToSdk(ctx, model.VirdefVersion)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "VirdefVersion"})
		}
		if unpacked != nil {
			sdk.VirdefVersion = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.HipObjectsAntiMalwareCriteria", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for HipObjectsAntiMalwareCriteria ---
func packHipObjectsAntiMalwareCriteriaFromSdk(ctx context.Context, sdk objects.HipObjectsAntiMalwareCriteria) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.HipObjectsAntiMalwareCriteria", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.HipObjectsAntiMalwareCriteria
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.IsInstalled != nil {
		model.IsInstalled = basetypes.NewBoolValue(*sdk.IsInstalled)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "IsInstalled", "value": *sdk.IsInstalled})
	} else {
		model.IsInstalled = basetypes.NewBoolNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.LastScanTime != nil {
		tflog.Debug(ctx, "Packing nested object for field LastScanTime")
		packed, d := packHipObjectsAntiMalwareCriteriaLastScanTimeFromSdk(ctx, *sdk.LastScanTime)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "LastScanTime"})
		}
		model.LastScanTime = packed
	} else {
		model.LastScanTime = basetypes.NewObjectNull(models.HipObjectsAntiMalwareCriteriaLastScanTime{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.ProductVersion != nil {
		tflog.Debug(ctx, "Packing nested object for field ProductVersion")
		packed, d := packHipObjectsAntiMalwareCriteriaProductVersionFromSdk(ctx, *sdk.ProductVersion)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "ProductVersion"})
		}
		model.ProductVersion = packed
	} else {
		model.ProductVersion = basetypes.NewObjectNull(models.HipObjectsAntiMalwareCriteriaProductVersion{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.RealTimeProtection != nil {
		model.RealTimeProtection = basetypes.NewStringValue(*sdk.RealTimeProtection)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "RealTimeProtection", "value": *sdk.RealTimeProtection})
	} else {
		model.RealTimeProtection = basetypes.NewStringNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.VirdefVersion != nil {
		tflog.Debug(ctx, "Packing nested object for field VirdefVersion")
		packed, d := packHipObjectsAntiMalwareCriteriaVirdefVersionFromSdk(ctx, *sdk.VirdefVersion)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "VirdefVersion"})
		}
		model.VirdefVersion = packed
	} else {
		model.VirdefVersion = basetypes.NewObjectNull(models.HipObjectsAntiMalwareCriteriaVirdefVersion{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.HipObjectsAntiMalwareCriteria{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.HipObjectsAntiMalwareCriteria", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for HipObjectsAntiMalwareCriteria ---
func unpackHipObjectsAntiMalwareCriteriaListToSdk(ctx context.Context, list types.List) ([]objects.HipObjectsAntiMalwareCriteria, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.HipObjectsAntiMalwareCriteria")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsAntiMalwareCriteria
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.HipObjectsAntiMalwareCriteria, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.HipObjectsAntiMalwareCriteria{}.AttrTypes(), &item)
		unpacked, d := unpackHipObjectsAntiMalwareCriteriaToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.HipObjectsAntiMalwareCriteria", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for HipObjectsAntiMalwareCriteria ---
func packHipObjectsAntiMalwareCriteriaListFromSdk(ctx context.Context, sdks []objects.HipObjectsAntiMalwareCriteria) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.HipObjectsAntiMalwareCriteria")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsAntiMalwareCriteria

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.HipObjectsAntiMalwareCriteria
		obj, d := packHipObjectsAntiMalwareCriteriaFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.HipObjectsAntiMalwareCriteria{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.HipObjectsAntiMalwareCriteria", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.HipObjectsAntiMalwareCriteria{}.AttrType(), data)
}

// --- Unpacker for HipObjectsAntiMalwareCriteriaLastScanTime ---
func unpackHipObjectsAntiMalwareCriteriaLastScanTimeToSdk(ctx context.Context, obj types.Object) (*objects.HipObjectsAntiMalwareCriteriaLastScanTime, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.HipObjectsAntiMalwareCriteriaLastScanTime", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.HipObjectsAntiMalwareCriteriaLastScanTime
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.HipObjectsAntiMalwareCriteriaLastScanTime
	var d diag.Diagnostics
	// Handling Typeless Objects
	if !model.NotAvailable.IsNull() && !model.NotAvailable.IsUnknown() {
		tflog.Debug(ctx, "Unpacking typeless object for field NotAvailable")
		sdk.NotAvailable = make(map[string]interface{})
	}

	// Handling Objects
	if !model.NotWithin.IsNull() && !model.NotWithin.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field NotWithin")
		unpacked, d := unpackHipObjectsAntiMalwareCriteriaLastScanTimeNotWithinToSdk(ctx, model.NotWithin)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "NotWithin"})
		}
		if unpacked != nil {
			sdk.NotWithin = unpacked
		}
	}

	// Handling Objects
	if !model.Within.IsNull() && !model.Within.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Within")
		unpacked, d := unpackHipObjectsAntiMalwareCriteriaLastScanTimeNotWithinToSdk(ctx, model.Within)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Within"})
		}
		if unpacked != nil {
			sdk.Within = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.HipObjectsAntiMalwareCriteriaLastScanTime", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for HipObjectsAntiMalwareCriteriaLastScanTime ---
func packHipObjectsAntiMalwareCriteriaLastScanTimeFromSdk(ctx context.Context, sdk objects.HipObjectsAntiMalwareCriteriaLastScanTime) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.HipObjectsAntiMalwareCriteriaLastScanTime", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.HipObjectsAntiMalwareCriteriaLastScanTime
	var d diag.Diagnostics
	// Handling Objects
	// This is a marker object (e.g. CHAP: {}). We just need to create an empty, non-null object.
	if sdk.NotAvailable != nil && !reflect.ValueOf(sdk.NotAvailable).IsNil() {
		tflog.Debug(ctx, "Packing typeless object for field NotAvailable")
		var d diag.Diagnostics
		// Create an empty object with no attributes, which signifies its presence.
		model.NotAvailable, d = basetypes.NewObjectValue(map[string]attr.Type{}, map[string]attr.Value{})
		diags.Append(d...)
	} else {
		// Since this field is part of a oneOf, being nil means it's not selected.
		// We make the object null with an empty attribute map.
		model.NotAvailable = basetypes.NewObjectNull(map[string]attr.Type{})
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.NotWithin != nil {
		tflog.Debug(ctx, "Packing nested object for field NotWithin")
		packed, d := packHipObjectsAntiMalwareCriteriaLastScanTimeNotWithinFromSdk(ctx, *sdk.NotWithin)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "NotWithin"})
		}
		model.NotWithin = packed
	} else {
		model.NotWithin = basetypes.NewObjectNull(models.HipObjectsAntiMalwareCriteriaLastScanTimeNotWithin{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Within != nil {
		tflog.Debug(ctx, "Packing nested object for field Within")
		packed, d := packHipObjectsAntiMalwareCriteriaLastScanTimeNotWithinFromSdk(ctx, *sdk.Within)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Within"})
		}
		model.Within = packed
	} else {
		model.Within = basetypes.NewObjectNull(models.HipObjectsAntiMalwareCriteriaLastScanTimeNotWithin{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.HipObjectsAntiMalwareCriteriaLastScanTime{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.HipObjectsAntiMalwareCriteriaLastScanTime", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for HipObjectsAntiMalwareCriteriaLastScanTime ---
func unpackHipObjectsAntiMalwareCriteriaLastScanTimeListToSdk(ctx context.Context, list types.List) ([]objects.HipObjectsAntiMalwareCriteriaLastScanTime, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.HipObjectsAntiMalwareCriteriaLastScanTime")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsAntiMalwareCriteriaLastScanTime
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.HipObjectsAntiMalwareCriteriaLastScanTime, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.HipObjectsAntiMalwareCriteriaLastScanTime{}.AttrTypes(), &item)
		unpacked, d := unpackHipObjectsAntiMalwareCriteriaLastScanTimeToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.HipObjectsAntiMalwareCriteriaLastScanTime", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for HipObjectsAntiMalwareCriteriaLastScanTime ---
func packHipObjectsAntiMalwareCriteriaLastScanTimeListFromSdk(ctx context.Context, sdks []objects.HipObjectsAntiMalwareCriteriaLastScanTime) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.HipObjectsAntiMalwareCriteriaLastScanTime")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsAntiMalwareCriteriaLastScanTime

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.HipObjectsAntiMalwareCriteriaLastScanTime
		obj, d := packHipObjectsAntiMalwareCriteriaLastScanTimeFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.HipObjectsAntiMalwareCriteriaLastScanTime{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.HipObjectsAntiMalwareCriteriaLastScanTime", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.HipObjectsAntiMalwareCriteriaLastScanTime{}.AttrType(), data)
}

// --- Unpacker for HipObjectsAntiMalwareCriteriaLastScanTimeNotWithin ---
func unpackHipObjectsAntiMalwareCriteriaLastScanTimeNotWithinToSdk(ctx context.Context, obj types.Object) (*objects.HipObjectsAntiMalwareCriteriaLastScanTimeNotWithin, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.HipObjectsAntiMalwareCriteriaLastScanTimeNotWithin", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.HipObjectsAntiMalwareCriteriaLastScanTimeNotWithin
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.HipObjectsAntiMalwareCriteriaLastScanTimeNotWithin
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Days.IsNull() && !model.Days.IsUnknown() {
		val := int32(model.Days.ValueInt64())
		sdk.Days = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Days", "value": *sdk.Days})
	}

	// Handling Primitives
	if !model.Hours.IsNull() && !model.Hours.IsUnknown() {
		val := int32(model.Hours.ValueInt64())
		sdk.Hours = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Hours", "value": *sdk.Hours})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.HipObjectsAntiMalwareCriteriaLastScanTimeNotWithin", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for HipObjectsAntiMalwareCriteriaLastScanTimeNotWithin ---
func packHipObjectsAntiMalwareCriteriaLastScanTimeNotWithinFromSdk(ctx context.Context, sdk objects.HipObjectsAntiMalwareCriteriaLastScanTimeNotWithin) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.HipObjectsAntiMalwareCriteriaLastScanTimeNotWithin", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.HipObjectsAntiMalwareCriteriaLastScanTimeNotWithin
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Days != nil {
		model.Days = basetypes.NewInt64Value(int64(*sdk.Days))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Days", "value": *sdk.Days})
	} else {
		model.Days = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Hours != nil {
		model.Hours = basetypes.NewInt64Value(int64(*sdk.Hours))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Hours", "value": *sdk.Hours})
	} else {
		model.Hours = basetypes.NewInt64Null()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.HipObjectsAntiMalwareCriteriaLastScanTimeNotWithin{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.HipObjectsAntiMalwareCriteriaLastScanTimeNotWithin", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for HipObjectsAntiMalwareCriteriaLastScanTimeNotWithin ---
func unpackHipObjectsAntiMalwareCriteriaLastScanTimeNotWithinListToSdk(ctx context.Context, list types.List) ([]objects.HipObjectsAntiMalwareCriteriaLastScanTimeNotWithin, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.HipObjectsAntiMalwareCriteriaLastScanTimeNotWithin")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsAntiMalwareCriteriaLastScanTimeNotWithin
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.HipObjectsAntiMalwareCriteriaLastScanTimeNotWithin, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.HipObjectsAntiMalwareCriteriaLastScanTimeNotWithin{}.AttrTypes(), &item)
		unpacked, d := unpackHipObjectsAntiMalwareCriteriaLastScanTimeNotWithinToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.HipObjectsAntiMalwareCriteriaLastScanTimeNotWithin", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for HipObjectsAntiMalwareCriteriaLastScanTimeNotWithin ---
func packHipObjectsAntiMalwareCriteriaLastScanTimeNotWithinListFromSdk(ctx context.Context, sdks []objects.HipObjectsAntiMalwareCriteriaLastScanTimeNotWithin) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.HipObjectsAntiMalwareCriteriaLastScanTimeNotWithin")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsAntiMalwareCriteriaLastScanTimeNotWithin

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.HipObjectsAntiMalwareCriteriaLastScanTimeNotWithin
		obj, d := packHipObjectsAntiMalwareCriteriaLastScanTimeNotWithinFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.HipObjectsAntiMalwareCriteriaLastScanTimeNotWithin{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.HipObjectsAntiMalwareCriteriaLastScanTimeNotWithin", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.HipObjectsAntiMalwareCriteriaLastScanTimeNotWithin{}.AttrType(), data)
}

// --- Unpacker for HipObjectsAntiMalwareCriteriaProductVersion ---
func unpackHipObjectsAntiMalwareCriteriaProductVersionToSdk(ctx context.Context, obj types.Object) (*objects.HipObjectsAntiMalwareCriteriaProductVersion, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.HipObjectsAntiMalwareCriteriaProductVersion", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.HipObjectsAntiMalwareCriteriaProductVersion
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.HipObjectsAntiMalwareCriteriaProductVersion
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Contains.IsNull() && !model.Contains.IsUnknown() {
		sdk.Contains = model.Contains.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Contains", "value": *sdk.Contains})
	}

	// Handling Primitives
	if !model.GreaterEqual.IsNull() && !model.GreaterEqual.IsUnknown() {
		sdk.GreaterEqual = model.GreaterEqual.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "GreaterEqual", "value": *sdk.GreaterEqual})
	}

	// Handling Primitives
	if !model.GreaterThan.IsNull() && !model.GreaterThan.IsUnknown() {
		sdk.GreaterThan = model.GreaterThan.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "GreaterThan", "value": *sdk.GreaterThan})
	}

	// Handling Primitives
	if !model.Is.IsNull() && !model.Is.IsUnknown() {
		sdk.Is = model.Is.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Is", "value": *sdk.Is})
	}

	// Handling Primitives
	if !model.IsNot.IsNull() && !model.IsNot.IsUnknown() {
		sdk.IsNot = model.IsNot.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "IsNot", "value": *sdk.IsNot})
	}

	// Handling Primitives
	if !model.LessEqual.IsNull() && !model.LessEqual.IsUnknown() {
		sdk.LessEqual = model.LessEqual.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "LessEqual", "value": *sdk.LessEqual})
	}

	// Handling Primitives
	if !model.LessThan.IsNull() && !model.LessThan.IsUnknown() {
		sdk.LessThan = model.LessThan.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "LessThan", "value": *sdk.LessThan})
	}

	// Handling Objects
	if !model.NotWithin.IsNull() && !model.NotWithin.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field NotWithin")
		unpacked, d := unpackHipObjectsAntiMalwareCriteriaProductVersionNotWithinToSdk(ctx, model.NotWithin)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "NotWithin"})
		}
		if unpacked != nil {
			sdk.NotWithin = unpacked
		}
	}

	// Handling Objects
	if !model.Within.IsNull() && !model.Within.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Within")
		unpacked, d := unpackHipObjectsAntiMalwareCriteriaProductVersionNotWithinToSdk(ctx, model.Within)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Within"})
		}
		if unpacked != nil {
			sdk.Within = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.HipObjectsAntiMalwareCriteriaProductVersion", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for HipObjectsAntiMalwareCriteriaProductVersion ---
func packHipObjectsAntiMalwareCriteriaProductVersionFromSdk(ctx context.Context, sdk objects.HipObjectsAntiMalwareCriteriaProductVersion) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.HipObjectsAntiMalwareCriteriaProductVersion", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.HipObjectsAntiMalwareCriteriaProductVersion
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Contains != nil {
		model.Contains = basetypes.NewStringValue(*sdk.Contains)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Contains", "value": *sdk.Contains})
	} else {
		model.Contains = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.GreaterEqual != nil {
		model.GreaterEqual = basetypes.NewStringValue(*sdk.GreaterEqual)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "GreaterEqual", "value": *sdk.GreaterEqual})
	} else {
		model.GreaterEqual = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.GreaterThan != nil {
		model.GreaterThan = basetypes.NewStringValue(*sdk.GreaterThan)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "GreaterThan", "value": *sdk.GreaterThan})
	} else {
		model.GreaterThan = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Is != nil {
		model.Is = basetypes.NewStringValue(*sdk.Is)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Is", "value": *sdk.Is})
	} else {
		model.Is = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.IsNot != nil {
		model.IsNot = basetypes.NewStringValue(*sdk.IsNot)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "IsNot", "value": *sdk.IsNot})
	} else {
		model.IsNot = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.LessEqual != nil {
		model.LessEqual = basetypes.NewStringValue(*sdk.LessEqual)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "LessEqual", "value": *sdk.LessEqual})
	} else {
		model.LessEqual = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.LessThan != nil {
		model.LessThan = basetypes.NewStringValue(*sdk.LessThan)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "LessThan", "value": *sdk.LessThan})
	} else {
		model.LessThan = basetypes.NewStringNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.NotWithin != nil {
		tflog.Debug(ctx, "Packing nested object for field NotWithin")
		packed, d := packHipObjectsAntiMalwareCriteriaProductVersionNotWithinFromSdk(ctx, *sdk.NotWithin)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "NotWithin"})
		}
		model.NotWithin = packed
	} else {
		model.NotWithin = basetypes.NewObjectNull(models.HipObjectsAntiMalwareCriteriaProductVersionNotWithin{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Within != nil {
		tflog.Debug(ctx, "Packing nested object for field Within")
		packed, d := packHipObjectsAntiMalwareCriteriaProductVersionNotWithinFromSdk(ctx, *sdk.Within)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Within"})
		}
		model.Within = packed
	} else {
		model.Within = basetypes.NewObjectNull(models.HipObjectsAntiMalwareCriteriaProductVersionNotWithin{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.HipObjectsAntiMalwareCriteriaProductVersion{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.HipObjectsAntiMalwareCriteriaProductVersion", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for HipObjectsAntiMalwareCriteriaProductVersion ---
func unpackHipObjectsAntiMalwareCriteriaProductVersionListToSdk(ctx context.Context, list types.List) ([]objects.HipObjectsAntiMalwareCriteriaProductVersion, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.HipObjectsAntiMalwareCriteriaProductVersion")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsAntiMalwareCriteriaProductVersion
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.HipObjectsAntiMalwareCriteriaProductVersion, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.HipObjectsAntiMalwareCriteriaProductVersion{}.AttrTypes(), &item)
		unpacked, d := unpackHipObjectsAntiMalwareCriteriaProductVersionToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.HipObjectsAntiMalwareCriteriaProductVersion", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for HipObjectsAntiMalwareCriteriaProductVersion ---
func packHipObjectsAntiMalwareCriteriaProductVersionListFromSdk(ctx context.Context, sdks []objects.HipObjectsAntiMalwareCriteriaProductVersion) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.HipObjectsAntiMalwareCriteriaProductVersion")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsAntiMalwareCriteriaProductVersion

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.HipObjectsAntiMalwareCriteriaProductVersion
		obj, d := packHipObjectsAntiMalwareCriteriaProductVersionFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.HipObjectsAntiMalwareCriteriaProductVersion{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.HipObjectsAntiMalwareCriteriaProductVersion", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.HipObjectsAntiMalwareCriteriaProductVersion{}.AttrType(), data)
}

// --- Unpacker for HipObjectsAntiMalwareCriteriaProductVersionNotWithin ---
func unpackHipObjectsAntiMalwareCriteriaProductVersionNotWithinToSdk(ctx context.Context, obj types.Object) (*objects.HipObjectsAntiMalwareCriteriaProductVersionNotWithin, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.HipObjectsAntiMalwareCriteriaProductVersionNotWithin", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.HipObjectsAntiMalwareCriteriaProductVersionNotWithin
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.HipObjectsAntiMalwareCriteriaProductVersionNotWithin
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Versions.IsNull() && !model.Versions.IsUnknown() {
		sdk.Versions = int32(model.Versions.ValueInt64())
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Versions", "value": sdk.Versions})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.HipObjectsAntiMalwareCriteriaProductVersionNotWithin", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for HipObjectsAntiMalwareCriteriaProductVersionNotWithin ---
func packHipObjectsAntiMalwareCriteriaProductVersionNotWithinFromSdk(ctx context.Context, sdk objects.HipObjectsAntiMalwareCriteriaProductVersionNotWithin) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.HipObjectsAntiMalwareCriteriaProductVersionNotWithin", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.HipObjectsAntiMalwareCriteriaProductVersionNotWithin
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.Versions = basetypes.NewInt64Value(int64(sdk.Versions))
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.HipObjectsAntiMalwareCriteriaProductVersionNotWithin{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.HipObjectsAntiMalwareCriteriaProductVersionNotWithin", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for HipObjectsAntiMalwareCriteriaProductVersionNotWithin ---
func unpackHipObjectsAntiMalwareCriteriaProductVersionNotWithinListToSdk(ctx context.Context, list types.List) ([]objects.HipObjectsAntiMalwareCriteriaProductVersionNotWithin, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.HipObjectsAntiMalwareCriteriaProductVersionNotWithin")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsAntiMalwareCriteriaProductVersionNotWithin
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.HipObjectsAntiMalwareCriteriaProductVersionNotWithin, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.HipObjectsAntiMalwareCriteriaProductVersionNotWithin{}.AttrTypes(), &item)
		unpacked, d := unpackHipObjectsAntiMalwareCriteriaProductVersionNotWithinToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.HipObjectsAntiMalwareCriteriaProductVersionNotWithin", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for HipObjectsAntiMalwareCriteriaProductVersionNotWithin ---
func packHipObjectsAntiMalwareCriteriaProductVersionNotWithinListFromSdk(ctx context.Context, sdks []objects.HipObjectsAntiMalwareCriteriaProductVersionNotWithin) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.HipObjectsAntiMalwareCriteriaProductVersionNotWithin")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsAntiMalwareCriteriaProductVersionNotWithin

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.HipObjectsAntiMalwareCriteriaProductVersionNotWithin
		obj, d := packHipObjectsAntiMalwareCriteriaProductVersionNotWithinFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.HipObjectsAntiMalwareCriteriaProductVersionNotWithin{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.HipObjectsAntiMalwareCriteriaProductVersionNotWithin", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.HipObjectsAntiMalwareCriteriaProductVersionNotWithin{}.AttrType(), data)
}

// --- Unpacker for HipObjectsAntiMalwareCriteriaVirdefVersion ---
func unpackHipObjectsAntiMalwareCriteriaVirdefVersionToSdk(ctx context.Context, obj types.Object) (*objects.HipObjectsAntiMalwareCriteriaVirdefVersion, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.HipObjectsAntiMalwareCriteriaVirdefVersion", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.HipObjectsAntiMalwareCriteriaVirdefVersion
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.HipObjectsAntiMalwareCriteriaVirdefVersion
	var d diag.Diagnostics
	// Handling Objects
	if !model.NotWithin.IsNull() && !model.NotWithin.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field NotWithin")
		unpacked, d := unpackHipObjectsAntiMalwareCriteriaVirdefVersionNotWithinToSdk(ctx, model.NotWithin)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "NotWithin"})
		}
		if unpacked != nil {
			sdk.NotWithin = unpacked
		}
	}

	// Handling Objects
	if !model.Within.IsNull() && !model.Within.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Within")
		unpacked, d := unpackHipObjectsAntiMalwareCriteriaVirdefVersionNotWithinToSdk(ctx, model.Within)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Within"})
		}
		if unpacked != nil {
			sdk.Within = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.HipObjectsAntiMalwareCriteriaVirdefVersion", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for HipObjectsAntiMalwareCriteriaVirdefVersion ---
func packHipObjectsAntiMalwareCriteriaVirdefVersionFromSdk(ctx context.Context, sdk objects.HipObjectsAntiMalwareCriteriaVirdefVersion) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.HipObjectsAntiMalwareCriteriaVirdefVersion", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.HipObjectsAntiMalwareCriteriaVirdefVersion
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.NotWithin != nil {
		tflog.Debug(ctx, "Packing nested object for field NotWithin")
		packed, d := packHipObjectsAntiMalwareCriteriaVirdefVersionNotWithinFromSdk(ctx, *sdk.NotWithin)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "NotWithin"})
		}
		model.NotWithin = packed
	} else {
		model.NotWithin = basetypes.NewObjectNull(models.HipObjectsAntiMalwareCriteriaVirdefVersionNotWithin{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Within != nil {
		tflog.Debug(ctx, "Packing nested object for field Within")
		packed, d := packHipObjectsAntiMalwareCriteriaVirdefVersionNotWithinFromSdk(ctx, *sdk.Within)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Within"})
		}
		model.Within = packed
	} else {
		model.Within = basetypes.NewObjectNull(models.HipObjectsAntiMalwareCriteriaVirdefVersionNotWithin{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.HipObjectsAntiMalwareCriteriaVirdefVersion{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.HipObjectsAntiMalwareCriteriaVirdefVersion", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for HipObjectsAntiMalwareCriteriaVirdefVersion ---
func unpackHipObjectsAntiMalwareCriteriaVirdefVersionListToSdk(ctx context.Context, list types.List) ([]objects.HipObjectsAntiMalwareCriteriaVirdefVersion, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.HipObjectsAntiMalwareCriteriaVirdefVersion")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsAntiMalwareCriteriaVirdefVersion
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.HipObjectsAntiMalwareCriteriaVirdefVersion, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.HipObjectsAntiMalwareCriteriaVirdefVersion{}.AttrTypes(), &item)
		unpacked, d := unpackHipObjectsAntiMalwareCriteriaVirdefVersionToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.HipObjectsAntiMalwareCriteriaVirdefVersion", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for HipObjectsAntiMalwareCriteriaVirdefVersion ---
func packHipObjectsAntiMalwareCriteriaVirdefVersionListFromSdk(ctx context.Context, sdks []objects.HipObjectsAntiMalwareCriteriaVirdefVersion) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.HipObjectsAntiMalwareCriteriaVirdefVersion")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsAntiMalwareCriteriaVirdefVersion

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.HipObjectsAntiMalwareCriteriaVirdefVersion
		obj, d := packHipObjectsAntiMalwareCriteriaVirdefVersionFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.HipObjectsAntiMalwareCriteriaVirdefVersion{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.HipObjectsAntiMalwareCriteriaVirdefVersion", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.HipObjectsAntiMalwareCriteriaVirdefVersion{}.AttrType(), data)
}

// --- Unpacker for HipObjectsAntiMalwareCriteriaVirdefVersionNotWithin ---
func unpackHipObjectsAntiMalwareCriteriaVirdefVersionNotWithinToSdk(ctx context.Context, obj types.Object) (*objects.HipObjectsAntiMalwareCriteriaVirdefVersionNotWithin, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.HipObjectsAntiMalwareCriteriaVirdefVersionNotWithin", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.HipObjectsAntiMalwareCriteriaVirdefVersionNotWithin
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.HipObjectsAntiMalwareCriteriaVirdefVersionNotWithin
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Days.IsNull() && !model.Days.IsUnknown() {
		val := int32(model.Days.ValueInt64())
		sdk.Days = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Days", "value": *sdk.Days})
	}

	// Handling Primitives
	if !model.Versions.IsNull() && !model.Versions.IsUnknown() {
		val := int32(model.Versions.ValueInt64())
		sdk.Versions = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Versions", "value": *sdk.Versions})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.HipObjectsAntiMalwareCriteriaVirdefVersionNotWithin", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for HipObjectsAntiMalwareCriteriaVirdefVersionNotWithin ---
func packHipObjectsAntiMalwareCriteriaVirdefVersionNotWithinFromSdk(ctx context.Context, sdk objects.HipObjectsAntiMalwareCriteriaVirdefVersionNotWithin) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.HipObjectsAntiMalwareCriteriaVirdefVersionNotWithin", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.HipObjectsAntiMalwareCriteriaVirdefVersionNotWithin
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Days != nil {
		model.Days = basetypes.NewInt64Value(int64(*sdk.Days))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Days", "value": *sdk.Days})
	} else {
		model.Days = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Versions != nil {
		model.Versions = basetypes.NewInt64Value(int64(*sdk.Versions))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Versions", "value": *sdk.Versions})
	} else {
		model.Versions = basetypes.NewInt64Null()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.HipObjectsAntiMalwareCriteriaVirdefVersionNotWithin{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.HipObjectsAntiMalwareCriteriaVirdefVersionNotWithin", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for HipObjectsAntiMalwareCriteriaVirdefVersionNotWithin ---
func unpackHipObjectsAntiMalwareCriteriaVirdefVersionNotWithinListToSdk(ctx context.Context, list types.List) ([]objects.HipObjectsAntiMalwareCriteriaVirdefVersionNotWithin, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.HipObjectsAntiMalwareCriteriaVirdefVersionNotWithin")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsAntiMalwareCriteriaVirdefVersionNotWithin
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.HipObjectsAntiMalwareCriteriaVirdefVersionNotWithin, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.HipObjectsAntiMalwareCriteriaVirdefVersionNotWithin{}.AttrTypes(), &item)
		unpacked, d := unpackHipObjectsAntiMalwareCriteriaVirdefVersionNotWithinToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.HipObjectsAntiMalwareCriteriaVirdefVersionNotWithin", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for HipObjectsAntiMalwareCriteriaVirdefVersionNotWithin ---
func packHipObjectsAntiMalwareCriteriaVirdefVersionNotWithinListFromSdk(ctx context.Context, sdks []objects.HipObjectsAntiMalwareCriteriaVirdefVersionNotWithin) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.HipObjectsAntiMalwareCriteriaVirdefVersionNotWithin")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsAntiMalwareCriteriaVirdefVersionNotWithin

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.HipObjectsAntiMalwareCriteriaVirdefVersionNotWithin
		obj, d := packHipObjectsAntiMalwareCriteriaVirdefVersionNotWithinFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.HipObjectsAntiMalwareCriteriaVirdefVersionNotWithin{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.HipObjectsAntiMalwareCriteriaVirdefVersionNotWithin", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.HipObjectsAntiMalwareCriteriaVirdefVersionNotWithin{}.AttrType(), data)
}

// --- Unpacker for HipObjectsAntiMalwareVendorInner ---
func unpackHipObjectsAntiMalwareVendorInnerToSdk(ctx context.Context, obj types.Object) (*objects.HipObjectsAntiMalwareVendorInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.HipObjectsAntiMalwareVendorInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.HipObjectsAntiMalwareVendorInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.HipObjectsAntiMalwareVendorInner
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	// Handling Lists
	if !model.Product.IsNull() && !model.Product.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Product")
		diags.Append(model.Product.ElementsAs(ctx, &sdk.Product, false)...)
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.HipObjectsAntiMalwareVendorInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for HipObjectsAntiMalwareVendorInner ---
func packHipObjectsAntiMalwareVendorInnerFromSdk(ctx context.Context, sdk objects.HipObjectsAntiMalwareVendorInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.HipObjectsAntiMalwareVendorInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.HipObjectsAntiMalwareVendorInner
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	// Handling Lists
	if sdk.Product != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Product")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Product, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Product)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Product = basetypes.NewListNull(elemType)
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.HipObjectsAntiMalwareVendorInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.HipObjectsAntiMalwareVendorInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for HipObjectsAntiMalwareVendorInner ---
func unpackHipObjectsAntiMalwareVendorInnerListToSdk(ctx context.Context, list types.List) ([]objects.HipObjectsAntiMalwareVendorInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.HipObjectsAntiMalwareVendorInner")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsAntiMalwareVendorInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.HipObjectsAntiMalwareVendorInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.HipObjectsAntiMalwareVendorInner{}.AttrTypes(), &item)
		unpacked, d := unpackHipObjectsAntiMalwareVendorInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.HipObjectsAntiMalwareVendorInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for HipObjectsAntiMalwareVendorInner ---
func packHipObjectsAntiMalwareVendorInnerListFromSdk(ctx context.Context, sdks []objects.HipObjectsAntiMalwareVendorInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.HipObjectsAntiMalwareVendorInner")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsAntiMalwareVendorInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.HipObjectsAntiMalwareVendorInner
		obj, d := packHipObjectsAntiMalwareVendorInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.HipObjectsAntiMalwareVendorInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.HipObjectsAntiMalwareVendorInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.HipObjectsAntiMalwareVendorInner{}.AttrType(), data)
}

// --- Unpacker for HipObjectsCertificate ---
func unpackHipObjectsCertificateToSdk(ctx context.Context, obj types.Object) (*objects.HipObjectsCertificate, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.HipObjectsCertificate", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.HipObjectsCertificate
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.HipObjectsCertificate
	var d diag.Diagnostics
	// Handling Objects
	if !model.Criteria.IsNull() && !model.Criteria.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Criteria")
		unpacked, d := unpackHipObjectsCertificateCriteriaToSdk(ctx, model.Criteria)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Criteria"})
		}
		if unpacked != nil {
			sdk.Criteria = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.HipObjectsCertificate", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for HipObjectsCertificate ---
func packHipObjectsCertificateFromSdk(ctx context.Context, sdk objects.HipObjectsCertificate) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.HipObjectsCertificate", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.HipObjectsCertificate
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Criteria != nil {
		tflog.Debug(ctx, "Packing nested object for field Criteria")
		packed, d := packHipObjectsCertificateCriteriaFromSdk(ctx, *sdk.Criteria)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Criteria"})
		}
		model.Criteria = packed
	} else {
		model.Criteria = basetypes.NewObjectNull(models.HipObjectsCertificateCriteria{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.HipObjectsCertificate{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.HipObjectsCertificate", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for HipObjectsCertificate ---
func unpackHipObjectsCertificateListToSdk(ctx context.Context, list types.List) ([]objects.HipObjectsCertificate, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.HipObjectsCertificate")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsCertificate
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.HipObjectsCertificate, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.HipObjectsCertificate{}.AttrTypes(), &item)
		unpacked, d := unpackHipObjectsCertificateToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.HipObjectsCertificate", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for HipObjectsCertificate ---
func packHipObjectsCertificateListFromSdk(ctx context.Context, sdks []objects.HipObjectsCertificate) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.HipObjectsCertificate")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsCertificate

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.HipObjectsCertificate
		obj, d := packHipObjectsCertificateFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.HipObjectsCertificate{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.HipObjectsCertificate", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.HipObjectsCertificate{}.AttrType(), data)
}

// --- Unpacker for HipObjectsCertificateCriteria ---
func unpackHipObjectsCertificateCriteriaToSdk(ctx context.Context, obj types.Object) (*objects.HipObjectsCertificateCriteria, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.HipObjectsCertificateCriteria", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.HipObjectsCertificateCriteria
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.HipObjectsCertificateCriteria
	var d diag.Diagnostics
	// Handling Lists
	if !model.CertificateAttributes.IsNull() && !model.CertificateAttributes.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field CertificateAttributes")
		unpacked, d := unpackHipObjectsCertificateCriteriaCertificateAttributesInnerListToSdk(ctx, model.CertificateAttributes)
		diags.Append(d...)
		sdk.CertificateAttributes = unpacked
	}

	// Handling Primitives
	if !model.CertificateProfile.IsNull() && !model.CertificateProfile.IsUnknown() {
		sdk.CertificateProfile = model.CertificateProfile.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "CertificateProfile", "value": *sdk.CertificateProfile})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.HipObjectsCertificateCriteria", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for HipObjectsCertificateCriteria ---
func packHipObjectsCertificateCriteriaFromSdk(ctx context.Context, sdk objects.HipObjectsCertificateCriteria) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.HipObjectsCertificateCriteria", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.HipObjectsCertificateCriteria
	var d diag.Diagnostics
	// Handling Lists
	if sdk.CertificateAttributes != nil {
		tflog.Debug(ctx, "Packing list of objects for field CertificateAttributes")
		packed, d := packHipObjectsCertificateCriteriaCertificateAttributesInnerListFromSdk(ctx, sdk.CertificateAttributes)
		diags.Append(d...)
		model.CertificateAttributes = packed
	} else {
		model.CertificateAttributes = basetypes.NewListNull(models.HipObjectsCertificateCriteriaCertificateAttributesInner{}.AttrType())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.CertificateProfile != nil {
		model.CertificateProfile = basetypes.NewStringValue(*sdk.CertificateProfile)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "CertificateProfile", "value": *sdk.CertificateProfile})
	} else {
		model.CertificateProfile = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.HipObjectsCertificateCriteria{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.HipObjectsCertificateCriteria", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for HipObjectsCertificateCriteria ---
func unpackHipObjectsCertificateCriteriaListToSdk(ctx context.Context, list types.List) ([]objects.HipObjectsCertificateCriteria, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.HipObjectsCertificateCriteria")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsCertificateCriteria
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.HipObjectsCertificateCriteria, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.HipObjectsCertificateCriteria{}.AttrTypes(), &item)
		unpacked, d := unpackHipObjectsCertificateCriteriaToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.HipObjectsCertificateCriteria", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for HipObjectsCertificateCriteria ---
func packHipObjectsCertificateCriteriaListFromSdk(ctx context.Context, sdks []objects.HipObjectsCertificateCriteria) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.HipObjectsCertificateCriteria")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsCertificateCriteria

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.HipObjectsCertificateCriteria
		obj, d := packHipObjectsCertificateCriteriaFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.HipObjectsCertificateCriteria{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.HipObjectsCertificateCriteria", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.HipObjectsCertificateCriteria{}.AttrType(), data)
}

// --- Unpacker for HipObjectsCertificateCriteriaCertificateAttributesInner ---
func unpackHipObjectsCertificateCriteriaCertificateAttributesInnerToSdk(ctx context.Context, obj types.Object) (*objects.HipObjectsCertificateCriteriaCertificateAttributesInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.HipObjectsCertificateCriteriaCertificateAttributesInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.HipObjectsCertificateCriteriaCertificateAttributesInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.HipObjectsCertificateCriteriaCertificateAttributesInner
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	// Handling Primitives
	if !model.Value.IsNull() && !model.Value.IsUnknown() {
		sdk.Value = model.Value.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Value", "value": *sdk.Value})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.HipObjectsCertificateCriteriaCertificateAttributesInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for HipObjectsCertificateCriteriaCertificateAttributesInner ---
func packHipObjectsCertificateCriteriaCertificateAttributesInnerFromSdk(ctx context.Context, sdk objects.HipObjectsCertificateCriteriaCertificateAttributesInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.HipObjectsCertificateCriteriaCertificateAttributesInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.HipObjectsCertificateCriteriaCertificateAttributesInner
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	// Handling Primitives
	// Standard primitive packing
	if sdk.Value != nil {
		model.Value = basetypes.NewStringValue(*sdk.Value)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Value", "value": *sdk.Value})
	} else {
		model.Value = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.HipObjectsCertificateCriteriaCertificateAttributesInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.HipObjectsCertificateCriteriaCertificateAttributesInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for HipObjectsCertificateCriteriaCertificateAttributesInner ---
func unpackHipObjectsCertificateCriteriaCertificateAttributesInnerListToSdk(ctx context.Context, list types.List) ([]objects.HipObjectsCertificateCriteriaCertificateAttributesInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.HipObjectsCertificateCriteriaCertificateAttributesInner")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsCertificateCriteriaCertificateAttributesInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.HipObjectsCertificateCriteriaCertificateAttributesInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.HipObjectsCertificateCriteriaCertificateAttributesInner{}.AttrTypes(), &item)
		unpacked, d := unpackHipObjectsCertificateCriteriaCertificateAttributesInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.HipObjectsCertificateCriteriaCertificateAttributesInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for HipObjectsCertificateCriteriaCertificateAttributesInner ---
func packHipObjectsCertificateCriteriaCertificateAttributesInnerListFromSdk(ctx context.Context, sdks []objects.HipObjectsCertificateCriteriaCertificateAttributesInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.HipObjectsCertificateCriteriaCertificateAttributesInner")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsCertificateCriteriaCertificateAttributesInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.HipObjectsCertificateCriteriaCertificateAttributesInner
		obj, d := packHipObjectsCertificateCriteriaCertificateAttributesInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.HipObjectsCertificateCriteriaCertificateAttributesInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.HipObjectsCertificateCriteriaCertificateAttributesInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.HipObjectsCertificateCriteriaCertificateAttributesInner{}.AttrType(), data)
}

// --- Unpacker for HipObjectsCustomChecks ---
func unpackHipObjectsCustomChecksToSdk(ctx context.Context, obj types.Object) (*objects.HipObjectsCustomChecks, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.HipObjectsCustomChecks", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.HipObjectsCustomChecks
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.HipObjectsCustomChecks
	var d diag.Diagnostics
	// Handling Objects
	if !model.Criteria.IsNull() && !model.Criteria.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Criteria")
		unpacked, d := unpackHipObjectsCustomChecksCriteriaToSdk(ctx, model.Criteria)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Criteria"})
		}
		if unpacked != nil {
			sdk.Criteria = *unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.HipObjectsCustomChecks", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for HipObjectsCustomChecks ---
func packHipObjectsCustomChecksFromSdk(ctx context.Context, sdk objects.HipObjectsCustomChecks) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.HipObjectsCustomChecks", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.HipObjectsCustomChecks
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	// Logic for non-pointer / value-type nested objects
	if !reflect.ValueOf(sdk.Criteria).IsZero() {
		tflog.Debug(ctx, "Packing nested object for field Criteria")
		packed, d := packHipObjectsCustomChecksCriteriaFromSdk(ctx, sdk.Criteria)
		diags.Append(d...)
		model.Criteria = packed
	} else {
		model.Criteria = basetypes.NewObjectNull(models.HipObjectsCustomChecksCriteria{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.HipObjectsCustomChecks{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.HipObjectsCustomChecks", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for HipObjectsCustomChecks ---
func unpackHipObjectsCustomChecksListToSdk(ctx context.Context, list types.List) ([]objects.HipObjectsCustomChecks, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.HipObjectsCustomChecks")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsCustomChecks
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.HipObjectsCustomChecks, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.HipObjectsCustomChecks{}.AttrTypes(), &item)
		unpacked, d := unpackHipObjectsCustomChecksToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.HipObjectsCustomChecks", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for HipObjectsCustomChecks ---
func packHipObjectsCustomChecksListFromSdk(ctx context.Context, sdks []objects.HipObjectsCustomChecks) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.HipObjectsCustomChecks")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsCustomChecks

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.HipObjectsCustomChecks
		obj, d := packHipObjectsCustomChecksFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.HipObjectsCustomChecks{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.HipObjectsCustomChecks", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.HipObjectsCustomChecks{}.AttrType(), data)
}

// --- Unpacker for HipObjectsCustomChecksCriteria ---
func unpackHipObjectsCustomChecksCriteriaToSdk(ctx context.Context, obj types.Object) (*objects.HipObjectsCustomChecksCriteria, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.HipObjectsCustomChecksCriteria", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.HipObjectsCustomChecksCriteria
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.HipObjectsCustomChecksCriteria
	var d diag.Diagnostics
	// Handling Lists
	if !model.Plist.IsNull() && !model.Plist.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field Plist")
		unpacked, d := unpackHipObjectsCustomChecksCriteriaPlistInnerListToSdk(ctx, model.Plist)
		diags.Append(d...)
		sdk.Plist = unpacked
	}

	// Handling Lists
	if !model.ProcessList.IsNull() && !model.ProcessList.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field ProcessList")
		unpacked, d := unpackHipObjectsCustomChecksCriteriaProcessListInnerListToSdk(ctx, model.ProcessList)
		diags.Append(d...)
		sdk.ProcessList = unpacked
	}

	// Handling Lists
	if !model.RegistryKey.IsNull() && !model.RegistryKey.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field RegistryKey")
		unpacked, d := unpackHipObjectsCustomChecksCriteriaRegistryKeyInnerListToSdk(ctx, model.RegistryKey)
		diags.Append(d...)
		sdk.RegistryKey = unpacked
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.HipObjectsCustomChecksCriteria", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for HipObjectsCustomChecksCriteria ---
func packHipObjectsCustomChecksCriteriaFromSdk(ctx context.Context, sdk objects.HipObjectsCustomChecksCriteria) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.HipObjectsCustomChecksCriteria", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.HipObjectsCustomChecksCriteria
	var d diag.Diagnostics
	// Handling Lists
	if sdk.Plist != nil {
		tflog.Debug(ctx, "Packing list of objects for field Plist")
		packed, d := packHipObjectsCustomChecksCriteriaPlistInnerListFromSdk(ctx, sdk.Plist)
		diags.Append(d...)
		model.Plist = packed
	} else {
		model.Plist = basetypes.NewListNull(models.HipObjectsCustomChecksCriteriaPlistInner{}.AttrType())
	}
	// Handling Lists
	if sdk.ProcessList != nil {
		tflog.Debug(ctx, "Packing list of objects for field ProcessList")
		packed, d := packHipObjectsCustomChecksCriteriaProcessListInnerListFromSdk(ctx, sdk.ProcessList)
		diags.Append(d...)
		model.ProcessList = packed
	} else {
		model.ProcessList = basetypes.NewListNull(models.HipObjectsCustomChecksCriteriaProcessListInner{}.AttrType())
	}
	// Handling Lists
	if sdk.RegistryKey != nil {
		tflog.Debug(ctx, "Packing list of objects for field RegistryKey")
		packed, d := packHipObjectsCustomChecksCriteriaRegistryKeyInnerListFromSdk(ctx, sdk.RegistryKey)
		diags.Append(d...)
		model.RegistryKey = packed
	} else {
		model.RegistryKey = basetypes.NewListNull(models.HipObjectsCustomChecksCriteriaRegistryKeyInner{}.AttrType())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.HipObjectsCustomChecksCriteria{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.HipObjectsCustomChecksCriteria", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for HipObjectsCustomChecksCriteria ---
func unpackHipObjectsCustomChecksCriteriaListToSdk(ctx context.Context, list types.List) ([]objects.HipObjectsCustomChecksCriteria, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.HipObjectsCustomChecksCriteria")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsCustomChecksCriteria
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.HipObjectsCustomChecksCriteria, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.HipObjectsCustomChecksCriteria{}.AttrTypes(), &item)
		unpacked, d := unpackHipObjectsCustomChecksCriteriaToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.HipObjectsCustomChecksCriteria", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for HipObjectsCustomChecksCriteria ---
func packHipObjectsCustomChecksCriteriaListFromSdk(ctx context.Context, sdks []objects.HipObjectsCustomChecksCriteria) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.HipObjectsCustomChecksCriteria")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsCustomChecksCriteria

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.HipObjectsCustomChecksCriteria
		obj, d := packHipObjectsCustomChecksCriteriaFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.HipObjectsCustomChecksCriteria{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.HipObjectsCustomChecksCriteria", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.HipObjectsCustomChecksCriteria{}.AttrType(), data)
}

// --- Unpacker for HipObjectsCustomChecksCriteriaPlistInner ---
func unpackHipObjectsCustomChecksCriteriaPlistInnerToSdk(ctx context.Context, obj types.Object) (*objects.HipObjectsCustomChecksCriteriaPlistInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.HipObjectsCustomChecksCriteriaPlistInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.HipObjectsCustomChecksCriteriaPlistInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.HipObjectsCustomChecksCriteriaPlistInner
	var d diag.Diagnostics
	// Handling Lists
	if !model.Key.IsNull() && !model.Key.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field Key")
		unpacked, d := unpackHipObjectsCustomChecksCriteriaPlistInnerKeyInnerListToSdk(ctx, model.Key)
		diags.Append(d...)
		sdk.Key = unpacked
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	// Handling Primitives
	if !model.Negate.IsNull() && !model.Negate.IsUnknown() {
		sdk.Negate = model.Negate.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Negate", "value": *sdk.Negate})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.HipObjectsCustomChecksCriteriaPlistInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for HipObjectsCustomChecksCriteriaPlistInner ---
func packHipObjectsCustomChecksCriteriaPlistInnerFromSdk(ctx context.Context, sdk objects.HipObjectsCustomChecksCriteriaPlistInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.HipObjectsCustomChecksCriteriaPlistInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.HipObjectsCustomChecksCriteriaPlistInner
	var d diag.Diagnostics
	// Handling Lists
	if sdk.Key != nil {
		tflog.Debug(ctx, "Packing list of objects for field Key")
		packed, d := packHipObjectsCustomChecksCriteriaPlistInnerKeyInnerListFromSdk(ctx, sdk.Key)
		diags.Append(d...)
		model.Key = packed
	} else {
		model.Key = basetypes.NewListNull(models.HipObjectsCustomChecksCriteriaPlistInnerKeyInner{}.AttrType())
	}
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	// Handling Primitives
	// Standard primitive packing
	if sdk.Negate != nil {
		model.Negate = basetypes.NewBoolValue(*sdk.Negate)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Negate", "value": *sdk.Negate})
	} else {
		model.Negate = basetypes.NewBoolNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.HipObjectsCustomChecksCriteriaPlistInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.HipObjectsCustomChecksCriteriaPlistInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for HipObjectsCustomChecksCriteriaPlistInner ---
func unpackHipObjectsCustomChecksCriteriaPlistInnerListToSdk(ctx context.Context, list types.List) ([]objects.HipObjectsCustomChecksCriteriaPlistInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.HipObjectsCustomChecksCriteriaPlistInner")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsCustomChecksCriteriaPlistInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.HipObjectsCustomChecksCriteriaPlistInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.HipObjectsCustomChecksCriteriaPlistInner{}.AttrTypes(), &item)
		unpacked, d := unpackHipObjectsCustomChecksCriteriaPlistInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.HipObjectsCustomChecksCriteriaPlistInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for HipObjectsCustomChecksCriteriaPlistInner ---
func packHipObjectsCustomChecksCriteriaPlistInnerListFromSdk(ctx context.Context, sdks []objects.HipObjectsCustomChecksCriteriaPlistInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.HipObjectsCustomChecksCriteriaPlistInner")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsCustomChecksCriteriaPlistInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.HipObjectsCustomChecksCriteriaPlistInner
		obj, d := packHipObjectsCustomChecksCriteriaPlistInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.HipObjectsCustomChecksCriteriaPlistInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.HipObjectsCustomChecksCriteriaPlistInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.HipObjectsCustomChecksCriteriaPlistInner{}.AttrType(), data)
}

// --- Unpacker for HipObjectsCustomChecksCriteriaPlistInnerKeyInner ---
func unpackHipObjectsCustomChecksCriteriaPlistInnerKeyInnerToSdk(ctx context.Context, obj types.Object) (*objects.HipObjectsCustomChecksCriteriaPlistInnerKeyInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.HipObjectsCustomChecksCriteriaPlistInnerKeyInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.HipObjectsCustomChecksCriteriaPlistInnerKeyInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.HipObjectsCustomChecksCriteriaPlistInnerKeyInner
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	// Handling Primitives
	if !model.Negate.IsNull() && !model.Negate.IsUnknown() {
		sdk.Negate = model.Negate.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Negate", "value": *sdk.Negate})
	}

	// Handling Primitives
	if !model.Value.IsNull() && !model.Value.IsUnknown() {
		sdk.Value = model.Value.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Value", "value": *sdk.Value})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.HipObjectsCustomChecksCriteriaPlistInnerKeyInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for HipObjectsCustomChecksCriteriaPlistInnerKeyInner ---
func packHipObjectsCustomChecksCriteriaPlistInnerKeyInnerFromSdk(ctx context.Context, sdk objects.HipObjectsCustomChecksCriteriaPlistInnerKeyInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.HipObjectsCustomChecksCriteriaPlistInnerKeyInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.HipObjectsCustomChecksCriteriaPlistInnerKeyInner
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	// Handling Primitives
	// Standard primitive packing
	if sdk.Negate != nil {
		model.Negate = basetypes.NewBoolValue(*sdk.Negate)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Negate", "value": *sdk.Negate})
	} else {
		model.Negate = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Value != nil {
		model.Value = basetypes.NewStringValue(*sdk.Value)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Value", "value": *sdk.Value})
	} else {
		model.Value = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.HipObjectsCustomChecksCriteriaPlistInnerKeyInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.HipObjectsCustomChecksCriteriaPlistInnerKeyInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for HipObjectsCustomChecksCriteriaPlistInnerKeyInner ---
func unpackHipObjectsCustomChecksCriteriaPlistInnerKeyInnerListToSdk(ctx context.Context, list types.List) ([]objects.HipObjectsCustomChecksCriteriaPlistInnerKeyInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.HipObjectsCustomChecksCriteriaPlistInnerKeyInner")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsCustomChecksCriteriaPlistInnerKeyInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.HipObjectsCustomChecksCriteriaPlistInnerKeyInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.HipObjectsCustomChecksCriteriaPlistInnerKeyInner{}.AttrTypes(), &item)
		unpacked, d := unpackHipObjectsCustomChecksCriteriaPlistInnerKeyInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.HipObjectsCustomChecksCriteriaPlistInnerKeyInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for HipObjectsCustomChecksCriteriaPlistInnerKeyInner ---
func packHipObjectsCustomChecksCriteriaPlistInnerKeyInnerListFromSdk(ctx context.Context, sdks []objects.HipObjectsCustomChecksCriteriaPlistInnerKeyInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.HipObjectsCustomChecksCriteriaPlistInnerKeyInner")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsCustomChecksCriteriaPlistInnerKeyInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.HipObjectsCustomChecksCriteriaPlistInnerKeyInner
		obj, d := packHipObjectsCustomChecksCriteriaPlistInnerKeyInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.HipObjectsCustomChecksCriteriaPlistInnerKeyInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.HipObjectsCustomChecksCriteriaPlistInnerKeyInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.HipObjectsCustomChecksCriteriaPlistInnerKeyInner{}.AttrType(), data)
}

// --- Unpacker for HipObjectsCustomChecksCriteriaProcessListInner ---
func unpackHipObjectsCustomChecksCriteriaProcessListInnerToSdk(ctx context.Context, obj types.Object) (*objects.HipObjectsCustomChecksCriteriaProcessListInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.HipObjectsCustomChecksCriteriaProcessListInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.HipObjectsCustomChecksCriteriaProcessListInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.HipObjectsCustomChecksCriteriaProcessListInner
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	// Handling Primitives
	if !model.Running.IsNull() && !model.Running.IsUnknown() {
		sdk.Running = model.Running.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Running", "value": *sdk.Running})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.HipObjectsCustomChecksCriteriaProcessListInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for HipObjectsCustomChecksCriteriaProcessListInner ---
func packHipObjectsCustomChecksCriteriaProcessListInnerFromSdk(ctx context.Context, sdk objects.HipObjectsCustomChecksCriteriaProcessListInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.HipObjectsCustomChecksCriteriaProcessListInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.HipObjectsCustomChecksCriteriaProcessListInner
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	// Handling Primitives
	// Standard primitive packing
	if sdk.Running != nil {
		model.Running = basetypes.NewBoolValue(*sdk.Running)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Running", "value": *sdk.Running})
	} else {
		model.Running = basetypes.NewBoolNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.HipObjectsCustomChecksCriteriaProcessListInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.HipObjectsCustomChecksCriteriaProcessListInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for HipObjectsCustomChecksCriteriaProcessListInner ---
func unpackHipObjectsCustomChecksCriteriaProcessListInnerListToSdk(ctx context.Context, list types.List) ([]objects.HipObjectsCustomChecksCriteriaProcessListInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.HipObjectsCustomChecksCriteriaProcessListInner")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsCustomChecksCriteriaProcessListInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.HipObjectsCustomChecksCriteriaProcessListInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.HipObjectsCustomChecksCriteriaProcessListInner{}.AttrTypes(), &item)
		unpacked, d := unpackHipObjectsCustomChecksCriteriaProcessListInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.HipObjectsCustomChecksCriteriaProcessListInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for HipObjectsCustomChecksCriteriaProcessListInner ---
func packHipObjectsCustomChecksCriteriaProcessListInnerListFromSdk(ctx context.Context, sdks []objects.HipObjectsCustomChecksCriteriaProcessListInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.HipObjectsCustomChecksCriteriaProcessListInner")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsCustomChecksCriteriaProcessListInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.HipObjectsCustomChecksCriteriaProcessListInner
		obj, d := packHipObjectsCustomChecksCriteriaProcessListInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.HipObjectsCustomChecksCriteriaProcessListInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.HipObjectsCustomChecksCriteriaProcessListInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.HipObjectsCustomChecksCriteriaProcessListInner{}.AttrType(), data)
}

// --- Unpacker for HipObjectsCustomChecksCriteriaRegistryKeyInner ---
func unpackHipObjectsCustomChecksCriteriaRegistryKeyInnerToSdk(ctx context.Context, obj types.Object) (*objects.HipObjectsCustomChecksCriteriaRegistryKeyInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.HipObjectsCustomChecksCriteriaRegistryKeyInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.HipObjectsCustomChecksCriteriaRegistryKeyInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.HipObjectsCustomChecksCriteriaRegistryKeyInner
	var d diag.Diagnostics
	// Handling Primitives
	if !model.DefaultValueData.IsNull() && !model.DefaultValueData.IsUnknown() {
		sdk.DefaultValueData = model.DefaultValueData.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DefaultValueData", "value": *sdk.DefaultValueData})
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	// Handling Primitives
	if !model.Negate.IsNull() && !model.Negate.IsUnknown() {
		sdk.Negate = model.Negate.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Negate", "value": *sdk.Negate})
	}

	// Handling Lists
	if !model.RegistryValue.IsNull() && !model.RegistryValue.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field RegistryValue")
		unpacked, d := unpackHipObjectsCustomChecksCriteriaRegistryKeyInnerRegistryValueInnerListToSdk(ctx, model.RegistryValue)
		diags.Append(d...)
		sdk.RegistryValue = unpacked
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.HipObjectsCustomChecksCriteriaRegistryKeyInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for HipObjectsCustomChecksCriteriaRegistryKeyInner ---
func packHipObjectsCustomChecksCriteriaRegistryKeyInnerFromSdk(ctx context.Context, sdk objects.HipObjectsCustomChecksCriteriaRegistryKeyInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.HipObjectsCustomChecksCriteriaRegistryKeyInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.HipObjectsCustomChecksCriteriaRegistryKeyInner
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.DefaultValueData != nil {
		model.DefaultValueData = basetypes.NewStringValue(*sdk.DefaultValueData)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DefaultValueData", "value": *sdk.DefaultValueData})
	} else {
		model.DefaultValueData = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	// Handling Primitives
	// Standard primitive packing
	if sdk.Negate != nil {
		model.Negate = basetypes.NewBoolValue(*sdk.Negate)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Negate", "value": *sdk.Negate})
	} else {
		model.Negate = basetypes.NewBoolNull()
	}
	// Handling Lists
	if sdk.RegistryValue != nil {
		tflog.Debug(ctx, "Packing list of objects for field RegistryValue")
		packed, d := packHipObjectsCustomChecksCriteriaRegistryKeyInnerRegistryValueInnerListFromSdk(ctx, sdk.RegistryValue)
		diags.Append(d...)
		model.RegistryValue = packed
	} else {
		model.RegistryValue = basetypes.NewListNull(models.HipObjectsCustomChecksCriteriaRegistryKeyInnerRegistryValueInner{}.AttrType())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.HipObjectsCustomChecksCriteriaRegistryKeyInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.HipObjectsCustomChecksCriteriaRegistryKeyInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for HipObjectsCustomChecksCriteriaRegistryKeyInner ---
func unpackHipObjectsCustomChecksCriteriaRegistryKeyInnerListToSdk(ctx context.Context, list types.List) ([]objects.HipObjectsCustomChecksCriteriaRegistryKeyInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.HipObjectsCustomChecksCriteriaRegistryKeyInner")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsCustomChecksCriteriaRegistryKeyInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.HipObjectsCustomChecksCriteriaRegistryKeyInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.HipObjectsCustomChecksCriteriaRegistryKeyInner{}.AttrTypes(), &item)
		unpacked, d := unpackHipObjectsCustomChecksCriteriaRegistryKeyInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.HipObjectsCustomChecksCriteriaRegistryKeyInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for HipObjectsCustomChecksCriteriaRegistryKeyInner ---
func packHipObjectsCustomChecksCriteriaRegistryKeyInnerListFromSdk(ctx context.Context, sdks []objects.HipObjectsCustomChecksCriteriaRegistryKeyInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.HipObjectsCustomChecksCriteriaRegistryKeyInner")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsCustomChecksCriteriaRegistryKeyInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.HipObjectsCustomChecksCriteriaRegistryKeyInner
		obj, d := packHipObjectsCustomChecksCriteriaRegistryKeyInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.HipObjectsCustomChecksCriteriaRegistryKeyInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.HipObjectsCustomChecksCriteriaRegistryKeyInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.HipObjectsCustomChecksCriteriaRegistryKeyInner{}.AttrType(), data)
}

// --- Unpacker for HipObjectsCustomChecksCriteriaRegistryKeyInnerRegistryValueInner ---
func unpackHipObjectsCustomChecksCriteriaRegistryKeyInnerRegistryValueInnerToSdk(ctx context.Context, obj types.Object) (*objects.HipObjectsCustomChecksCriteriaRegistryKeyInnerRegistryValueInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.HipObjectsCustomChecksCriteriaRegistryKeyInnerRegistryValueInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.HipObjectsCustomChecksCriteriaRegistryKeyInnerRegistryValueInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.HipObjectsCustomChecksCriteriaRegistryKeyInnerRegistryValueInner
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	// Handling Primitives
	if !model.Negate.IsNull() && !model.Negate.IsUnknown() {
		sdk.Negate = model.Negate.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Negate", "value": *sdk.Negate})
	}

	// Handling Primitives
	if !model.ValueData.IsNull() && !model.ValueData.IsUnknown() {
		sdk.ValueData = model.ValueData.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "ValueData", "value": *sdk.ValueData})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.HipObjectsCustomChecksCriteriaRegistryKeyInnerRegistryValueInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for HipObjectsCustomChecksCriteriaRegistryKeyInnerRegistryValueInner ---
func packHipObjectsCustomChecksCriteriaRegistryKeyInnerRegistryValueInnerFromSdk(ctx context.Context, sdk objects.HipObjectsCustomChecksCriteriaRegistryKeyInnerRegistryValueInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.HipObjectsCustomChecksCriteriaRegistryKeyInnerRegistryValueInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.HipObjectsCustomChecksCriteriaRegistryKeyInnerRegistryValueInner
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	// Handling Primitives
	// Standard primitive packing
	if sdk.Negate != nil {
		model.Negate = basetypes.NewBoolValue(*sdk.Negate)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Negate", "value": *sdk.Negate})
	} else {
		model.Negate = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.ValueData != nil {
		model.ValueData = basetypes.NewStringValue(*sdk.ValueData)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "ValueData", "value": *sdk.ValueData})
	} else {
		model.ValueData = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.HipObjectsCustomChecksCriteriaRegistryKeyInnerRegistryValueInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.HipObjectsCustomChecksCriteriaRegistryKeyInnerRegistryValueInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for HipObjectsCustomChecksCriteriaRegistryKeyInnerRegistryValueInner ---
func unpackHipObjectsCustomChecksCriteriaRegistryKeyInnerRegistryValueInnerListToSdk(ctx context.Context, list types.List) ([]objects.HipObjectsCustomChecksCriteriaRegistryKeyInnerRegistryValueInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.HipObjectsCustomChecksCriteriaRegistryKeyInnerRegistryValueInner")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsCustomChecksCriteriaRegistryKeyInnerRegistryValueInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.HipObjectsCustomChecksCriteriaRegistryKeyInnerRegistryValueInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.HipObjectsCustomChecksCriteriaRegistryKeyInnerRegistryValueInner{}.AttrTypes(), &item)
		unpacked, d := unpackHipObjectsCustomChecksCriteriaRegistryKeyInnerRegistryValueInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.HipObjectsCustomChecksCriteriaRegistryKeyInnerRegistryValueInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for HipObjectsCustomChecksCriteriaRegistryKeyInnerRegistryValueInner ---
func packHipObjectsCustomChecksCriteriaRegistryKeyInnerRegistryValueInnerListFromSdk(ctx context.Context, sdks []objects.HipObjectsCustomChecksCriteriaRegistryKeyInnerRegistryValueInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.HipObjectsCustomChecksCriteriaRegistryKeyInnerRegistryValueInner")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsCustomChecksCriteriaRegistryKeyInnerRegistryValueInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.HipObjectsCustomChecksCriteriaRegistryKeyInnerRegistryValueInner
		obj, d := packHipObjectsCustomChecksCriteriaRegistryKeyInnerRegistryValueInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.HipObjectsCustomChecksCriteriaRegistryKeyInnerRegistryValueInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.HipObjectsCustomChecksCriteriaRegistryKeyInnerRegistryValueInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.HipObjectsCustomChecksCriteriaRegistryKeyInnerRegistryValueInner{}.AttrType(), data)
}

// --- Unpacker for HipObjectsDataLossPrevention ---
func unpackHipObjectsDataLossPreventionToSdk(ctx context.Context, obj types.Object) (*objects.HipObjectsDataLossPrevention, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.HipObjectsDataLossPrevention", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.HipObjectsDataLossPrevention
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.HipObjectsDataLossPrevention
	var d diag.Diagnostics
	// Handling Objects
	if !model.Criteria.IsNull() && !model.Criteria.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Criteria")
		unpacked, d := unpackHipObjectsDataLossPreventionCriteriaToSdk(ctx, model.Criteria)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Criteria"})
		}
		if unpacked != nil {
			sdk.Criteria = unpacked
		}
	}

	// Handling Primitives
	if !model.ExcludeVendor.IsNull() && !model.ExcludeVendor.IsUnknown() {
		sdk.ExcludeVendor = model.ExcludeVendor.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "ExcludeVendor", "value": *sdk.ExcludeVendor})
	}

	// Handling Lists
	if !model.Vendor.IsNull() && !model.Vendor.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field Vendor")
		unpacked, d := unpackHipObjectsDataLossPreventionVendorInnerListToSdk(ctx, model.Vendor)
		diags.Append(d...)
		sdk.Vendor = unpacked
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.HipObjectsDataLossPrevention", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for HipObjectsDataLossPrevention ---
func packHipObjectsDataLossPreventionFromSdk(ctx context.Context, sdk objects.HipObjectsDataLossPrevention) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.HipObjectsDataLossPrevention", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.HipObjectsDataLossPrevention
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Criteria != nil {
		tflog.Debug(ctx, "Packing nested object for field Criteria")
		packed, d := packHipObjectsDataLossPreventionCriteriaFromSdk(ctx, *sdk.Criteria)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Criteria"})
		}
		model.Criteria = packed
	} else {
		model.Criteria = basetypes.NewObjectNull(models.HipObjectsDataLossPreventionCriteria{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.ExcludeVendor != nil {
		model.ExcludeVendor = basetypes.NewBoolValue(*sdk.ExcludeVendor)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "ExcludeVendor", "value": *sdk.ExcludeVendor})
	} else {
		model.ExcludeVendor = basetypes.NewBoolNull()
	}
	// Handling Lists
	if sdk.Vendor != nil {
		tflog.Debug(ctx, "Packing list of objects for field Vendor")
		packed, d := packHipObjectsDataLossPreventionVendorInnerListFromSdk(ctx, sdk.Vendor)
		diags.Append(d...)
		model.Vendor = packed
	} else {
		model.Vendor = basetypes.NewListNull(models.HipObjectsDataLossPreventionVendorInner{}.AttrType())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.HipObjectsDataLossPrevention{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.HipObjectsDataLossPrevention", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for HipObjectsDataLossPrevention ---
func unpackHipObjectsDataLossPreventionListToSdk(ctx context.Context, list types.List) ([]objects.HipObjectsDataLossPrevention, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.HipObjectsDataLossPrevention")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsDataLossPrevention
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.HipObjectsDataLossPrevention, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.HipObjectsDataLossPrevention{}.AttrTypes(), &item)
		unpacked, d := unpackHipObjectsDataLossPreventionToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.HipObjectsDataLossPrevention", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for HipObjectsDataLossPrevention ---
func packHipObjectsDataLossPreventionListFromSdk(ctx context.Context, sdks []objects.HipObjectsDataLossPrevention) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.HipObjectsDataLossPrevention")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsDataLossPrevention

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.HipObjectsDataLossPrevention
		obj, d := packHipObjectsDataLossPreventionFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.HipObjectsDataLossPrevention{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.HipObjectsDataLossPrevention", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.HipObjectsDataLossPrevention{}.AttrType(), data)
}

// --- Unpacker for HipObjectsDataLossPreventionCriteria ---
func unpackHipObjectsDataLossPreventionCriteriaToSdk(ctx context.Context, obj types.Object) (*objects.HipObjectsDataLossPreventionCriteria, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.HipObjectsDataLossPreventionCriteria", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.HipObjectsDataLossPreventionCriteria
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.HipObjectsDataLossPreventionCriteria
	var d diag.Diagnostics
	// Handling Primitives
	if !model.IsEnabled.IsNull() && !model.IsEnabled.IsUnknown() {
		sdk.IsEnabled = model.IsEnabled.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "IsEnabled", "value": *sdk.IsEnabled})
	}

	// Handling Primitives
	if !model.IsInstalled.IsNull() && !model.IsInstalled.IsUnknown() {
		sdk.IsInstalled = model.IsInstalled.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "IsInstalled", "value": *sdk.IsInstalled})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.HipObjectsDataLossPreventionCriteria", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for HipObjectsDataLossPreventionCriteria ---
func packHipObjectsDataLossPreventionCriteriaFromSdk(ctx context.Context, sdk objects.HipObjectsDataLossPreventionCriteria) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.HipObjectsDataLossPreventionCriteria", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.HipObjectsDataLossPreventionCriteria
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.IsEnabled != nil {
		model.IsEnabled = basetypes.NewStringValue(*sdk.IsEnabled)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "IsEnabled", "value": *sdk.IsEnabled})
	} else {
		model.IsEnabled = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.IsInstalled != nil {
		model.IsInstalled = basetypes.NewBoolValue(*sdk.IsInstalled)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "IsInstalled", "value": *sdk.IsInstalled})
	} else {
		model.IsInstalled = basetypes.NewBoolNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.HipObjectsDataLossPreventionCriteria{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.HipObjectsDataLossPreventionCriteria", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for HipObjectsDataLossPreventionCriteria ---
func unpackHipObjectsDataLossPreventionCriteriaListToSdk(ctx context.Context, list types.List) ([]objects.HipObjectsDataLossPreventionCriteria, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.HipObjectsDataLossPreventionCriteria")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsDataLossPreventionCriteria
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.HipObjectsDataLossPreventionCriteria, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.HipObjectsDataLossPreventionCriteria{}.AttrTypes(), &item)
		unpacked, d := unpackHipObjectsDataLossPreventionCriteriaToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.HipObjectsDataLossPreventionCriteria", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for HipObjectsDataLossPreventionCriteria ---
func packHipObjectsDataLossPreventionCriteriaListFromSdk(ctx context.Context, sdks []objects.HipObjectsDataLossPreventionCriteria) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.HipObjectsDataLossPreventionCriteria")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsDataLossPreventionCriteria

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.HipObjectsDataLossPreventionCriteria
		obj, d := packHipObjectsDataLossPreventionCriteriaFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.HipObjectsDataLossPreventionCriteria{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.HipObjectsDataLossPreventionCriteria", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.HipObjectsDataLossPreventionCriteria{}.AttrType(), data)
}

// --- Unpacker for HipObjectsDataLossPreventionVendorInner ---
func unpackHipObjectsDataLossPreventionVendorInnerToSdk(ctx context.Context, obj types.Object) (*objects.HipObjectsDataLossPreventionVendorInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.HipObjectsDataLossPreventionVendorInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.HipObjectsDataLossPreventionVendorInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.HipObjectsDataLossPreventionVendorInner
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	// Handling Lists
	if !model.Product.IsNull() && !model.Product.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Product")
		diags.Append(model.Product.ElementsAs(ctx, &sdk.Product, false)...)
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.HipObjectsDataLossPreventionVendorInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for HipObjectsDataLossPreventionVendorInner ---
func packHipObjectsDataLossPreventionVendorInnerFromSdk(ctx context.Context, sdk objects.HipObjectsDataLossPreventionVendorInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.HipObjectsDataLossPreventionVendorInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.HipObjectsDataLossPreventionVendorInner
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	// Handling Lists
	if sdk.Product != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Product")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Product, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Product)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Product = basetypes.NewListNull(elemType)
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.HipObjectsDataLossPreventionVendorInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.HipObjectsDataLossPreventionVendorInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for HipObjectsDataLossPreventionVendorInner ---
func unpackHipObjectsDataLossPreventionVendorInnerListToSdk(ctx context.Context, list types.List) ([]objects.HipObjectsDataLossPreventionVendorInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.HipObjectsDataLossPreventionVendorInner")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsDataLossPreventionVendorInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.HipObjectsDataLossPreventionVendorInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.HipObjectsDataLossPreventionVendorInner{}.AttrTypes(), &item)
		unpacked, d := unpackHipObjectsDataLossPreventionVendorInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.HipObjectsDataLossPreventionVendorInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for HipObjectsDataLossPreventionVendorInner ---
func packHipObjectsDataLossPreventionVendorInnerListFromSdk(ctx context.Context, sdks []objects.HipObjectsDataLossPreventionVendorInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.HipObjectsDataLossPreventionVendorInner")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsDataLossPreventionVendorInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.HipObjectsDataLossPreventionVendorInner
		obj, d := packHipObjectsDataLossPreventionVendorInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.HipObjectsDataLossPreventionVendorInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.HipObjectsDataLossPreventionVendorInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.HipObjectsDataLossPreventionVendorInner{}.AttrType(), data)
}

// --- Unpacker for HipObjectsDiskBackup ---
func unpackHipObjectsDiskBackupToSdk(ctx context.Context, obj types.Object) (*objects.HipObjectsDiskBackup, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.HipObjectsDiskBackup", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.HipObjectsDiskBackup
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.HipObjectsDiskBackup
	var d diag.Diagnostics
	// Handling Objects
	if !model.Criteria.IsNull() && !model.Criteria.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Criteria")
		unpacked, d := unpackHipObjectsDiskBackupCriteriaToSdk(ctx, model.Criteria)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Criteria"})
		}
		if unpacked != nil {
			sdk.Criteria = unpacked
		}
	}

	// Handling Primitives
	if !model.ExcludeVendor.IsNull() && !model.ExcludeVendor.IsUnknown() {
		sdk.ExcludeVendor = model.ExcludeVendor.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "ExcludeVendor", "value": *sdk.ExcludeVendor})
	}

	// Handling Lists
	if !model.Vendor.IsNull() && !model.Vendor.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field Vendor")
		unpacked, d := unpackHipObjectsAntiMalwareVendorInnerListToSdk(ctx, model.Vendor)
		diags.Append(d...)
		sdk.Vendor = unpacked
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.HipObjectsDiskBackup", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for HipObjectsDiskBackup ---
func packHipObjectsDiskBackupFromSdk(ctx context.Context, sdk objects.HipObjectsDiskBackup) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.HipObjectsDiskBackup", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.HipObjectsDiskBackup
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Criteria != nil {
		tflog.Debug(ctx, "Packing nested object for field Criteria")
		packed, d := packHipObjectsDiskBackupCriteriaFromSdk(ctx, *sdk.Criteria)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Criteria"})
		}
		model.Criteria = packed
	} else {
		model.Criteria = basetypes.NewObjectNull(models.HipObjectsDiskBackupCriteria{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.ExcludeVendor != nil {
		model.ExcludeVendor = basetypes.NewBoolValue(*sdk.ExcludeVendor)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "ExcludeVendor", "value": *sdk.ExcludeVendor})
	} else {
		model.ExcludeVendor = basetypes.NewBoolNull()
	}
	// Handling Lists
	if sdk.Vendor != nil {
		tflog.Debug(ctx, "Packing list of objects for field Vendor")
		packed, d := packHipObjectsAntiMalwareVendorInnerListFromSdk(ctx, sdk.Vendor)
		diags.Append(d...)
		model.Vendor = packed
	} else {
		model.Vendor = basetypes.NewListNull(models.HipObjectsAntiMalwareVendorInner{}.AttrType())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.HipObjectsDiskBackup{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.HipObjectsDiskBackup", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for HipObjectsDiskBackup ---
func unpackHipObjectsDiskBackupListToSdk(ctx context.Context, list types.List) ([]objects.HipObjectsDiskBackup, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.HipObjectsDiskBackup")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsDiskBackup
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.HipObjectsDiskBackup, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.HipObjectsDiskBackup{}.AttrTypes(), &item)
		unpacked, d := unpackHipObjectsDiskBackupToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.HipObjectsDiskBackup", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for HipObjectsDiskBackup ---
func packHipObjectsDiskBackupListFromSdk(ctx context.Context, sdks []objects.HipObjectsDiskBackup) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.HipObjectsDiskBackup")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsDiskBackup

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.HipObjectsDiskBackup
		obj, d := packHipObjectsDiskBackupFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.HipObjectsDiskBackup{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.HipObjectsDiskBackup", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.HipObjectsDiskBackup{}.AttrType(), data)
}

// --- Unpacker for HipObjectsDiskBackupCriteria ---
func unpackHipObjectsDiskBackupCriteriaToSdk(ctx context.Context, obj types.Object) (*objects.HipObjectsDiskBackupCriteria, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.HipObjectsDiskBackupCriteria", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.HipObjectsDiskBackupCriteria
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.HipObjectsDiskBackupCriteria
	var d diag.Diagnostics
	// Handling Primitives
	if !model.IsInstalled.IsNull() && !model.IsInstalled.IsUnknown() {
		sdk.IsInstalled = model.IsInstalled.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "IsInstalled", "value": *sdk.IsInstalled})
	}

	// Handling Objects
	if !model.LastBackupTime.IsNull() && !model.LastBackupTime.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field LastBackupTime")
		unpacked, d := unpackHipObjectsAntiMalwareCriteriaLastScanTimeToSdk(ctx, model.LastBackupTime)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "LastBackupTime"})
		}
		if unpacked != nil {
			sdk.LastBackupTime = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.HipObjectsDiskBackupCriteria", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for HipObjectsDiskBackupCriteria ---
func packHipObjectsDiskBackupCriteriaFromSdk(ctx context.Context, sdk objects.HipObjectsDiskBackupCriteria) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.HipObjectsDiskBackupCriteria", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.HipObjectsDiskBackupCriteria
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.IsInstalled != nil {
		model.IsInstalled = basetypes.NewBoolValue(*sdk.IsInstalled)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "IsInstalled", "value": *sdk.IsInstalled})
	} else {
		model.IsInstalled = basetypes.NewBoolNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.LastBackupTime != nil {
		tflog.Debug(ctx, "Packing nested object for field LastBackupTime")
		packed, d := packHipObjectsAntiMalwareCriteriaLastScanTimeFromSdk(ctx, *sdk.LastBackupTime)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "LastBackupTime"})
		}
		model.LastBackupTime = packed
	} else {
		model.LastBackupTime = basetypes.NewObjectNull(models.HipObjectsAntiMalwareCriteriaLastScanTime{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.HipObjectsDiskBackupCriteria{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.HipObjectsDiskBackupCriteria", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for HipObjectsDiskBackupCriteria ---
func unpackHipObjectsDiskBackupCriteriaListToSdk(ctx context.Context, list types.List) ([]objects.HipObjectsDiskBackupCriteria, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.HipObjectsDiskBackupCriteria")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsDiskBackupCriteria
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.HipObjectsDiskBackupCriteria, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.HipObjectsDiskBackupCriteria{}.AttrTypes(), &item)
		unpacked, d := unpackHipObjectsDiskBackupCriteriaToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.HipObjectsDiskBackupCriteria", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for HipObjectsDiskBackupCriteria ---
func packHipObjectsDiskBackupCriteriaListFromSdk(ctx context.Context, sdks []objects.HipObjectsDiskBackupCriteria) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.HipObjectsDiskBackupCriteria")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsDiskBackupCriteria

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.HipObjectsDiskBackupCriteria
		obj, d := packHipObjectsDiskBackupCriteriaFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.HipObjectsDiskBackupCriteria{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.HipObjectsDiskBackupCriteria", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.HipObjectsDiskBackupCriteria{}.AttrType(), data)
}

// --- Unpacker for HipObjectsDiskEncryption ---
func unpackHipObjectsDiskEncryptionToSdk(ctx context.Context, obj types.Object) (*objects.HipObjectsDiskEncryption, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.HipObjectsDiskEncryption", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.HipObjectsDiskEncryption
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.HipObjectsDiskEncryption
	var d diag.Diagnostics
	// Handling Objects
	if !model.Criteria.IsNull() && !model.Criteria.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Criteria")
		unpacked, d := unpackHipObjectsDiskEncryptionCriteriaToSdk(ctx, model.Criteria)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Criteria"})
		}
		if unpacked != nil {
			sdk.Criteria = unpacked
		}
	}

	// Handling Primitives
	if !model.ExcludeVendor.IsNull() && !model.ExcludeVendor.IsUnknown() {
		sdk.ExcludeVendor = model.ExcludeVendor.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "ExcludeVendor", "value": *sdk.ExcludeVendor})
	}

	// Handling Lists
	if !model.Vendor.IsNull() && !model.Vendor.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field Vendor")
		unpacked, d := unpackHipObjectsAntiMalwareVendorInnerListToSdk(ctx, model.Vendor)
		diags.Append(d...)
		sdk.Vendor = unpacked
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.HipObjectsDiskEncryption", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for HipObjectsDiskEncryption ---
func packHipObjectsDiskEncryptionFromSdk(ctx context.Context, sdk objects.HipObjectsDiskEncryption) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.HipObjectsDiskEncryption", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.HipObjectsDiskEncryption
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Criteria != nil {
		tflog.Debug(ctx, "Packing nested object for field Criteria")
		packed, d := packHipObjectsDiskEncryptionCriteriaFromSdk(ctx, *sdk.Criteria)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Criteria"})
		}
		model.Criteria = packed
	} else {
		model.Criteria = basetypes.NewObjectNull(models.HipObjectsDiskEncryptionCriteria{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.ExcludeVendor != nil {
		model.ExcludeVendor = basetypes.NewBoolValue(*sdk.ExcludeVendor)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "ExcludeVendor", "value": *sdk.ExcludeVendor})
	} else {
		model.ExcludeVendor = basetypes.NewBoolNull()
	}
	// Handling Lists
	if sdk.Vendor != nil {
		tflog.Debug(ctx, "Packing list of objects for field Vendor")
		packed, d := packHipObjectsAntiMalwareVendorInnerListFromSdk(ctx, sdk.Vendor)
		diags.Append(d...)
		model.Vendor = packed
	} else {
		model.Vendor = basetypes.NewListNull(models.HipObjectsAntiMalwareVendorInner{}.AttrType())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.HipObjectsDiskEncryption{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.HipObjectsDiskEncryption", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for HipObjectsDiskEncryption ---
func unpackHipObjectsDiskEncryptionListToSdk(ctx context.Context, list types.List) ([]objects.HipObjectsDiskEncryption, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.HipObjectsDiskEncryption")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsDiskEncryption
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.HipObjectsDiskEncryption, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.HipObjectsDiskEncryption{}.AttrTypes(), &item)
		unpacked, d := unpackHipObjectsDiskEncryptionToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.HipObjectsDiskEncryption", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for HipObjectsDiskEncryption ---
func packHipObjectsDiskEncryptionListFromSdk(ctx context.Context, sdks []objects.HipObjectsDiskEncryption) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.HipObjectsDiskEncryption")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsDiskEncryption

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.HipObjectsDiskEncryption
		obj, d := packHipObjectsDiskEncryptionFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.HipObjectsDiskEncryption{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.HipObjectsDiskEncryption", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.HipObjectsDiskEncryption{}.AttrType(), data)
}

// --- Unpacker for HipObjectsDiskEncryptionCriteria ---
func unpackHipObjectsDiskEncryptionCriteriaToSdk(ctx context.Context, obj types.Object) (*objects.HipObjectsDiskEncryptionCriteria, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.HipObjectsDiskEncryptionCriteria", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.HipObjectsDiskEncryptionCriteria
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.HipObjectsDiskEncryptionCriteria
	var d diag.Diagnostics
	// Handling Lists
	if !model.EncryptedLocations.IsNull() && !model.EncryptedLocations.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field EncryptedLocations")
		unpacked, d := unpackHipObjectsDiskEncryptionCriteriaEncryptedLocationsInnerListToSdk(ctx, model.EncryptedLocations)
		diags.Append(d...)
		sdk.EncryptedLocations = unpacked
	}

	// Handling Primitives
	if !model.IsInstalled.IsNull() && !model.IsInstalled.IsUnknown() {
		sdk.IsInstalled = model.IsInstalled.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "IsInstalled", "value": *sdk.IsInstalled})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.HipObjectsDiskEncryptionCriteria", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for HipObjectsDiskEncryptionCriteria ---
func packHipObjectsDiskEncryptionCriteriaFromSdk(ctx context.Context, sdk objects.HipObjectsDiskEncryptionCriteria) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.HipObjectsDiskEncryptionCriteria", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.HipObjectsDiskEncryptionCriteria
	var d diag.Diagnostics
	// Handling Lists
	if sdk.EncryptedLocations != nil {
		tflog.Debug(ctx, "Packing list of objects for field EncryptedLocations")
		packed, d := packHipObjectsDiskEncryptionCriteriaEncryptedLocationsInnerListFromSdk(ctx, sdk.EncryptedLocations)
		diags.Append(d...)
		model.EncryptedLocations = packed
	} else {
		model.EncryptedLocations = basetypes.NewListNull(models.HipObjectsDiskEncryptionCriteriaEncryptedLocationsInner{}.AttrType())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.IsInstalled != nil {
		model.IsInstalled = basetypes.NewBoolValue(*sdk.IsInstalled)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "IsInstalled", "value": *sdk.IsInstalled})
	} else {
		model.IsInstalled = basetypes.NewBoolNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.HipObjectsDiskEncryptionCriteria{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.HipObjectsDiskEncryptionCriteria", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for HipObjectsDiskEncryptionCriteria ---
func unpackHipObjectsDiskEncryptionCriteriaListToSdk(ctx context.Context, list types.List) ([]objects.HipObjectsDiskEncryptionCriteria, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.HipObjectsDiskEncryptionCriteria")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsDiskEncryptionCriteria
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.HipObjectsDiskEncryptionCriteria, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.HipObjectsDiskEncryptionCriteria{}.AttrTypes(), &item)
		unpacked, d := unpackHipObjectsDiskEncryptionCriteriaToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.HipObjectsDiskEncryptionCriteria", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for HipObjectsDiskEncryptionCriteria ---
func packHipObjectsDiskEncryptionCriteriaListFromSdk(ctx context.Context, sdks []objects.HipObjectsDiskEncryptionCriteria) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.HipObjectsDiskEncryptionCriteria")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsDiskEncryptionCriteria

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.HipObjectsDiskEncryptionCriteria
		obj, d := packHipObjectsDiskEncryptionCriteriaFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.HipObjectsDiskEncryptionCriteria{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.HipObjectsDiskEncryptionCriteria", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.HipObjectsDiskEncryptionCriteria{}.AttrType(), data)
}

// --- Unpacker for HipObjectsDiskEncryptionCriteriaEncryptedLocationsInner ---
func unpackHipObjectsDiskEncryptionCriteriaEncryptedLocationsInnerToSdk(ctx context.Context, obj types.Object) (*objects.HipObjectsDiskEncryptionCriteriaEncryptedLocationsInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.HipObjectsDiskEncryptionCriteriaEncryptedLocationsInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.HipObjectsDiskEncryptionCriteriaEncryptedLocationsInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.HipObjectsDiskEncryptionCriteriaEncryptedLocationsInner
	var d diag.Diagnostics
	// Handling Objects
	if !model.EncryptionState.IsNull() && !model.EncryptionState.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field EncryptionState")
		unpacked, d := unpackHipObjectsDiskEncryptionCriteriaEncryptedLocationsInnerEncryptionStateToSdk(ctx, model.EncryptionState)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "EncryptionState"})
		}
		if unpacked != nil {
			sdk.EncryptionState = unpacked
		}
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.HipObjectsDiskEncryptionCriteriaEncryptedLocationsInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for HipObjectsDiskEncryptionCriteriaEncryptedLocationsInner ---
func packHipObjectsDiskEncryptionCriteriaEncryptedLocationsInnerFromSdk(ctx context.Context, sdk objects.HipObjectsDiskEncryptionCriteriaEncryptedLocationsInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.HipObjectsDiskEncryptionCriteriaEncryptedLocationsInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.HipObjectsDiskEncryptionCriteriaEncryptedLocationsInner
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.EncryptionState != nil {
		tflog.Debug(ctx, "Packing nested object for field EncryptionState")
		packed, d := packHipObjectsDiskEncryptionCriteriaEncryptedLocationsInnerEncryptionStateFromSdk(ctx, *sdk.EncryptionState)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "EncryptionState"})
		}
		model.EncryptionState = packed
	} else {
		model.EncryptionState = basetypes.NewObjectNull(models.HipObjectsDiskEncryptionCriteriaEncryptedLocationsInnerEncryptionState{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.HipObjectsDiskEncryptionCriteriaEncryptedLocationsInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.HipObjectsDiskEncryptionCriteriaEncryptedLocationsInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for HipObjectsDiskEncryptionCriteriaEncryptedLocationsInner ---
func unpackHipObjectsDiskEncryptionCriteriaEncryptedLocationsInnerListToSdk(ctx context.Context, list types.List) ([]objects.HipObjectsDiskEncryptionCriteriaEncryptedLocationsInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.HipObjectsDiskEncryptionCriteriaEncryptedLocationsInner")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsDiskEncryptionCriteriaEncryptedLocationsInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.HipObjectsDiskEncryptionCriteriaEncryptedLocationsInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.HipObjectsDiskEncryptionCriteriaEncryptedLocationsInner{}.AttrTypes(), &item)
		unpacked, d := unpackHipObjectsDiskEncryptionCriteriaEncryptedLocationsInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.HipObjectsDiskEncryptionCriteriaEncryptedLocationsInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for HipObjectsDiskEncryptionCriteriaEncryptedLocationsInner ---
func packHipObjectsDiskEncryptionCriteriaEncryptedLocationsInnerListFromSdk(ctx context.Context, sdks []objects.HipObjectsDiskEncryptionCriteriaEncryptedLocationsInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.HipObjectsDiskEncryptionCriteriaEncryptedLocationsInner")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsDiskEncryptionCriteriaEncryptedLocationsInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.HipObjectsDiskEncryptionCriteriaEncryptedLocationsInner
		obj, d := packHipObjectsDiskEncryptionCriteriaEncryptedLocationsInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.HipObjectsDiskEncryptionCriteriaEncryptedLocationsInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.HipObjectsDiskEncryptionCriteriaEncryptedLocationsInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.HipObjectsDiskEncryptionCriteriaEncryptedLocationsInner{}.AttrType(), data)
}

// --- Unpacker for HipObjectsDiskEncryptionCriteriaEncryptedLocationsInnerEncryptionState ---
func unpackHipObjectsDiskEncryptionCriteriaEncryptedLocationsInnerEncryptionStateToSdk(ctx context.Context, obj types.Object) (*objects.HipObjectsDiskEncryptionCriteriaEncryptedLocationsInnerEncryptionState, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.HipObjectsDiskEncryptionCriteriaEncryptedLocationsInnerEncryptionState", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.HipObjectsDiskEncryptionCriteriaEncryptedLocationsInnerEncryptionState
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.HipObjectsDiskEncryptionCriteriaEncryptedLocationsInnerEncryptionState
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Is.IsNull() && !model.Is.IsUnknown() {
		sdk.Is = model.Is.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Is", "value": *sdk.Is})
	}

	// Handling Primitives
	if !model.IsNot.IsNull() && !model.IsNot.IsUnknown() {
		sdk.IsNot = model.IsNot.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "IsNot", "value": *sdk.IsNot})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.HipObjectsDiskEncryptionCriteriaEncryptedLocationsInnerEncryptionState", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for HipObjectsDiskEncryptionCriteriaEncryptedLocationsInnerEncryptionState ---
func packHipObjectsDiskEncryptionCriteriaEncryptedLocationsInnerEncryptionStateFromSdk(ctx context.Context, sdk objects.HipObjectsDiskEncryptionCriteriaEncryptedLocationsInnerEncryptionState) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.HipObjectsDiskEncryptionCriteriaEncryptedLocationsInnerEncryptionState", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.HipObjectsDiskEncryptionCriteriaEncryptedLocationsInnerEncryptionState
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Is != nil {
		model.Is = basetypes.NewStringValue(*sdk.Is)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Is", "value": *sdk.Is})
	} else {
		model.Is = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.IsNot != nil {
		model.IsNot = basetypes.NewStringValue(*sdk.IsNot)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "IsNot", "value": *sdk.IsNot})
	} else {
		model.IsNot = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.HipObjectsDiskEncryptionCriteriaEncryptedLocationsInnerEncryptionState{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.HipObjectsDiskEncryptionCriteriaEncryptedLocationsInnerEncryptionState", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for HipObjectsDiskEncryptionCriteriaEncryptedLocationsInnerEncryptionState ---
func unpackHipObjectsDiskEncryptionCriteriaEncryptedLocationsInnerEncryptionStateListToSdk(ctx context.Context, list types.List) ([]objects.HipObjectsDiskEncryptionCriteriaEncryptedLocationsInnerEncryptionState, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.HipObjectsDiskEncryptionCriteriaEncryptedLocationsInnerEncryptionState")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsDiskEncryptionCriteriaEncryptedLocationsInnerEncryptionState
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.HipObjectsDiskEncryptionCriteriaEncryptedLocationsInnerEncryptionState, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.HipObjectsDiskEncryptionCriteriaEncryptedLocationsInnerEncryptionState{}.AttrTypes(), &item)
		unpacked, d := unpackHipObjectsDiskEncryptionCriteriaEncryptedLocationsInnerEncryptionStateToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.HipObjectsDiskEncryptionCriteriaEncryptedLocationsInnerEncryptionState", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for HipObjectsDiskEncryptionCriteriaEncryptedLocationsInnerEncryptionState ---
func packHipObjectsDiskEncryptionCriteriaEncryptedLocationsInnerEncryptionStateListFromSdk(ctx context.Context, sdks []objects.HipObjectsDiskEncryptionCriteriaEncryptedLocationsInnerEncryptionState) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.HipObjectsDiskEncryptionCriteriaEncryptedLocationsInnerEncryptionState")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsDiskEncryptionCriteriaEncryptedLocationsInnerEncryptionState

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.HipObjectsDiskEncryptionCriteriaEncryptedLocationsInnerEncryptionState
		obj, d := packHipObjectsDiskEncryptionCriteriaEncryptedLocationsInnerEncryptionStateFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.HipObjectsDiskEncryptionCriteriaEncryptedLocationsInnerEncryptionState{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.HipObjectsDiskEncryptionCriteriaEncryptedLocationsInnerEncryptionState", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.HipObjectsDiskEncryptionCriteriaEncryptedLocationsInnerEncryptionState{}.AttrType(), data)
}

// --- Unpacker for HipObjectsFirewall ---
func unpackHipObjectsFirewallToSdk(ctx context.Context, obj types.Object) (*objects.HipObjectsFirewall, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.HipObjectsFirewall", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.HipObjectsFirewall
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.HipObjectsFirewall
	var d diag.Diagnostics
	// Handling Objects
	if !model.Criteria.IsNull() && !model.Criteria.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Criteria")
		unpacked, d := unpackHipObjectsDataLossPreventionCriteriaToSdk(ctx, model.Criteria)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Criteria"})
		}
		if unpacked != nil {
			sdk.Criteria = unpacked
		}
	}

	// Handling Primitives
	if !model.ExcludeVendor.IsNull() && !model.ExcludeVendor.IsUnknown() {
		sdk.ExcludeVendor = model.ExcludeVendor.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "ExcludeVendor", "value": *sdk.ExcludeVendor})
	}

	// Handling Lists
	if !model.Vendor.IsNull() && !model.Vendor.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field Vendor")
		unpacked, d := unpackHipObjectsAntiMalwareVendorInnerListToSdk(ctx, model.Vendor)
		diags.Append(d...)
		sdk.Vendor = unpacked
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.HipObjectsFirewall", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for HipObjectsFirewall ---
func packHipObjectsFirewallFromSdk(ctx context.Context, sdk objects.HipObjectsFirewall) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.HipObjectsFirewall", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.HipObjectsFirewall
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Criteria != nil {
		tflog.Debug(ctx, "Packing nested object for field Criteria")
		packed, d := packHipObjectsDataLossPreventionCriteriaFromSdk(ctx, *sdk.Criteria)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Criteria"})
		}
		model.Criteria = packed
	} else {
		model.Criteria = basetypes.NewObjectNull(models.HipObjectsDataLossPreventionCriteria{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.ExcludeVendor != nil {
		model.ExcludeVendor = basetypes.NewBoolValue(*sdk.ExcludeVendor)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "ExcludeVendor", "value": *sdk.ExcludeVendor})
	} else {
		model.ExcludeVendor = basetypes.NewBoolNull()
	}
	// Handling Lists
	if sdk.Vendor != nil {
		tflog.Debug(ctx, "Packing list of objects for field Vendor")
		packed, d := packHipObjectsAntiMalwareVendorInnerListFromSdk(ctx, sdk.Vendor)
		diags.Append(d...)
		model.Vendor = packed
	} else {
		model.Vendor = basetypes.NewListNull(models.HipObjectsAntiMalwareVendorInner{}.AttrType())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.HipObjectsFirewall{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.HipObjectsFirewall", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for HipObjectsFirewall ---
func unpackHipObjectsFirewallListToSdk(ctx context.Context, list types.List) ([]objects.HipObjectsFirewall, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.HipObjectsFirewall")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsFirewall
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.HipObjectsFirewall, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.HipObjectsFirewall{}.AttrTypes(), &item)
		unpacked, d := unpackHipObjectsFirewallToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.HipObjectsFirewall", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for HipObjectsFirewall ---
func packHipObjectsFirewallListFromSdk(ctx context.Context, sdks []objects.HipObjectsFirewall) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.HipObjectsFirewall")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsFirewall

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.HipObjectsFirewall
		obj, d := packHipObjectsFirewallFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.HipObjectsFirewall{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.HipObjectsFirewall", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.HipObjectsFirewall{}.AttrType(), data)
}

// --- Unpacker for HipObjectsHostInfo ---
func unpackHipObjectsHostInfoToSdk(ctx context.Context, obj types.Object) (*objects.HipObjectsHostInfo, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.HipObjectsHostInfo", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.HipObjectsHostInfo
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.HipObjectsHostInfo
	var d diag.Diagnostics
	// Handling Objects
	if !model.Criteria.IsNull() && !model.Criteria.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Criteria")
		unpacked, d := unpackHipObjectsHostInfoCriteriaToSdk(ctx, model.Criteria)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Criteria"})
		}
		if unpacked != nil {
			sdk.Criteria = *unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.HipObjectsHostInfo", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for HipObjectsHostInfo ---
func packHipObjectsHostInfoFromSdk(ctx context.Context, sdk objects.HipObjectsHostInfo) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.HipObjectsHostInfo", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.HipObjectsHostInfo
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	// Logic for non-pointer / value-type nested objects
	if !reflect.ValueOf(sdk.Criteria).IsZero() {
		tflog.Debug(ctx, "Packing nested object for field Criteria")
		packed, d := packHipObjectsHostInfoCriteriaFromSdk(ctx, sdk.Criteria)
		diags.Append(d...)
		model.Criteria = packed
	} else {
		model.Criteria = basetypes.NewObjectNull(models.HipObjectsHostInfoCriteria{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.HipObjectsHostInfo{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.HipObjectsHostInfo", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for HipObjectsHostInfo ---
func unpackHipObjectsHostInfoListToSdk(ctx context.Context, list types.List) ([]objects.HipObjectsHostInfo, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.HipObjectsHostInfo")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsHostInfo
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.HipObjectsHostInfo, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.HipObjectsHostInfo{}.AttrTypes(), &item)
		unpacked, d := unpackHipObjectsHostInfoToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.HipObjectsHostInfo", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for HipObjectsHostInfo ---
func packHipObjectsHostInfoListFromSdk(ctx context.Context, sdks []objects.HipObjectsHostInfo) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.HipObjectsHostInfo")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsHostInfo

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.HipObjectsHostInfo
		obj, d := packHipObjectsHostInfoFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.HipObjectsHostInfo{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.HipObjectsHostInfo", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.HipObjectsHostInfo{}.AttrType(), data)
}

// --- Unpacker for HipObjectsHostInfoCriteria ---
func unpackHipObjectsHostInfoCriteriaToSdk(ctx context.Context, obj types.Object) (*objects.HipObjectsHostInfoCriteria, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.HipObjectsHostInfoCriteria", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.HipObjectsHostInfoCriteria
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.HipObjectsHostInfoCriteria
	var d diag.Diagnostics
	// Handling Objects
	if !model.ClientVersion.IsNull() && !model.ClientVersion.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field ClientVersion")
		unpacked, d := unpackHipObjectsHostInfoCriteriaClientVersionToSdk(ctx, model.ClientVersion)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "ClientVersion"})
		}
		if unpacked != nil {
			sdk.ClientVersion = unpacked
		}
	}

	// Handling Objects
	if !model.Domain.IsNull() && !model.Domain.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Domain")
		unpacked, d := unpackHipObjectsHostInfoCriteriaClientVersionToSdk(ctx, model.Domain)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Domain"})
		}
		if unpacked != nil {
			sdk.Domain = unpacked
		}
	}

	// Handling Objects
	if !model.HostId.IsNull() && !model.HostId.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field HostId")
		unpacked, d := unpackHipObjectsHostInfoCriteriaClientVersionToSdk(ctx, model.HostId)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "HostId"})
		}
		if unpacked != nil {
			sdk.HostId = unpacked
		}
	}

	// Handling Objects
	if !model.HostName.IsNull() && !model.HostName.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field HostName")
		unpacked, d := unpackHipObjectsHostInfoCriteriaClientVersionToSdk(ctx, model.HostName)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "HostName"})
		}
		if unpacked != nil {
			sdk.HostName = unpacked
		}
	}

	// Handling Primitives
	if !model.Managed.IsNull() && !model.Managed.IsUnknown() {
		sdk.Managed = model.Managed.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Managed", "value": *sdk.Managed})
	}

	// Handling Objects
	if !model.Os.IsNull() && !model.Os.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Os")
		unpacked, d := unpackHipObjectsHostInfoCriteriaOsToSdk(ctx, model.Os)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Os"})
		}
		if unpacked != nil {
			sdk.Os = unpacked
		}
	}

	// Handling Objects
	if !model.SerialNumber.IsNull() && !model.SerialNumber.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field SerialNumber")
		unpacked, d := unpackHipObjectsHostInfoCriteriaClientVersionToSdk(ctx, model.SerialNumber)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "SerialNumber"})
		}
		if unpacked != nil {
			sdk.SerialNumber = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.HipObjectsHostInfoCriteria", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for HipObjectsHostInfoCriteria ---
func packHipObjectsHostInfoCriteriaFromSdk(ctx context.Context, sdk objects.HipObjectsHostInfoCriteria) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.HipObjectsHostInfoCriteria", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.HipObjectsHostInfoCriteria
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.ClientVersion != nil {
		tflog.Debug(ctx, "Packing nested object for field ClientVersion")
		packed, d := packHipObjectsHostInfoCriteriaClientVersionFromSdk(ctx, *sdk.ClientVersion)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "ClientVersion"})
		}
		model.ClientVersion = packed
	} else {
		model.ClientVersion = basetypes.NewObjectNull(models.HipObjectsHostInfoCriteriaClientVersion{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Domain != nil {
		tflog.Debug(ctx, "Packing nested object for field Domain")
		packed, d := packHipObjectsHostInfoCriteriaClientVersionFromSdk(ctx, *sdk.Domain)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Domain"})
		}
		model.Domain = packed
	} else {
		model.Domain = basetypes.NewObjectNull(models.HipObjectsHostInfoCriteriaClientVersion{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.HostId != nil {
		tflog.Debug(ctx, "Packing nested object for field HostId")
		packed, d := packHipObjectsHostInfoCriteriaClientVersionFromSdk(ctx, *sdk.HostId)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "HostId"})
		}
		model.HostId = packed
	} else {
		model.HostId = basetypes.NewObjectNull(models.HipObjectsHostInfoCriteriaClientVersion{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.HostName != nil {
		tflog.Debug(ctx, "Packing nested object for field HostName")
		packed, d := packHipObjectsHostInfoCriteriaClientVersionFromSdk(ctx, *sdk.HostName)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "HostName"})
		}
		model.HostName = packed
	} else {
		model.HostName = basetypes.NewObjectNull(models.HipObjectsHostInfoCriteriaClientVersion{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Managed != nil {
		model.Managed = basetypes.NewBoolValue(*sdk.Managed)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Managed", "value": *sdk.Managed})
	} else {
		model.Managed = basetypes.NewBoolNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Os != nil {
		tflog.Debug(ctx, "Packing nested object for field Os")
		packed, d := packHipObjectsHostInfoCriteriaOsFromSdk(ctx, *sdk.Os)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Os"})
		}
		model.Os = packed
	} else {
		model.Os = basetypes.NewObjectNull(models.HipObjectsHostInfoCriteriaOs{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.SerialNumber != nil {
		tflog.Debug(ctx, "Packing nested object for field SerialNumber")
		packed, d := packHipObjectsHostInfoCriteriaClientVersionFromSdk(ctx, *sdk.SerialNumber)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "SerialNumber"})
		}
		model.SerialNumber = packed
	} else {
		model.SerialNumber = basetypes.NewObjectNull(models.HipObjectsHostInfoCriteriaClientVersion{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.HipObjectsHostInfoCriteria{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.HipObjectsHostInfoCriteria", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for HipObjectsHostInfoCriteria ---
func unpackHipObjectsHostInfoCriteriaListToSdk(ctx context.Context, list types.List) ([]objects.HipObjectsHostInfoCriteria, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.HipObjectsHostInfoCriteria")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsHostInfoCriteria
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.HipObjectsHostInfoCriteria, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.HipObjectsHostInfoCriteria{}.AttrTypes(), &item)
		unpacked, d := unpackHipObjectsHostInfoCriteriaToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.HipObjectsHostInfoCriteria", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for HipObjectsHostInfoCriteria ---
func packHipObjectsHostInfoCriteriaListFromSdk(ctx context.Context, sdks []objects.HipObjectsHostInfoCriteria) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.HipObjectsHostInfoCriteria")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsHostInfoCriteria

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.HipObjectsHostInfoCriteria
		obj, d := packHipObjectsHostInfoCriteriaFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.HipObjectsHostInfoCriteria{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.HipObjectsHostInfoCriteria", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.HipObjectsHostInfoCriteria{}.AttrType(), data)
}

// --- Unpacker for HipObjectsHostInfoCriteriaClientVersion ---
func unpackHipObjectsHostInfoCriteriaClientVersionToSdk(ctx context.Context, obj types.Object) (*objects.HipObjectsHostInfoCriteriaClientVersion, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.HipObjectsHostInfoCriteriaClientVersion", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.HipObjectsHostInfoCriteriaClientVersion
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.HipObjectsHostInfoCriteriaClientVersion
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Contains.IsNull() && !model.Contains.IsUnknown() {
		sdk.Contains = model.Contains.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Contains", "value": *sdk.Contains})
	}

	// Handling Primitives
	if !model.Is.IsNull() && !model.Is.IsUnknown() {
		sdk.Is = model.Is.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Is", "value": *sdk.Is})
	}

	// Handling Primitives
	if !model.IsNot.IsNull() && !model.IsNot.IsUnknown() {
		sdk.IsNot = model.IsNot.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "IsNot", "value": *sdk.IsNot})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.HipObjectsHostInfoCriteriaClientVersion", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for HipObjectsHostInfoCriteriaClientVersion ---
func packHipObjectsHostInfoCriteriaClientVersionFromSdk(ctx context.Context, sdk objects.HipObjectsHostInfoCriteriaClientVersion) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.HipObjectsHostInfoCriteriaClientVersion", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.HipObjectsHostInfoCriteriaClientVersion
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Contains != nil {
		model.Contains = basetypes.NewStringValue(*sdk.Contains)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Contains", "value": *sdk.Contains})
	} else {
		model.Contains = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Is != nil {
		model.Is = basetypes.NewStringValue(*sdk.Is)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Is", "value": *sdk.Is})
	} else {
		model.Is = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.IsNot != nil {
		model.IsNot = basetypes.NewStringValue(*sdk.IsNot)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "IsNot", "value": *sdk.IsNot})
	} else {
		model.IsNot = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.HipObjectsHostInfoCriteriaClientVersion{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.HipObjectsHostInfoCriteriaClientVersion", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for HipObjectsHostInfoCriteriaClientVersion ---
func unpackHipObjectsHostInfoCriteriaClientVersionListToSdk(ctx context.Context, list types.List) ([]objects.HipObjectsHostInfoCriteriaClientVersion, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.HipObjectsHostInfoCriteriaClientVersion")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsHostInfoCriteriaClientVersion
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.HipObjectsHostInfoCriteriaClientVersion, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.HipObjectsHostInfoCriteriaClientVersion{}.AttrTypes(), &item)
		unpacked, d := unpackHipObjectsHostInfoCriteriaClientVersionToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.HipObjectsHostInfoCriteriaClientVersion", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for HipObjectsHostInfoCriteriaClientVersion ---
func packHipObjectsHostInfoCriteriaClientVersionListFromSdk(ctx context.Context, sdks []objects.HipObjectsHostInfoCriteriaClientVersion) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.HipObjectsHostInfoCriteriaClientVersion")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsHostInfoCriteriaClientVersion

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.HipObjectsHostInfoCriteriaClientVersion
		obj, d := packHipObjectsHostInfoCriteriaClientVersionFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.HipObjectsHostInfoCriteriaClientVersion{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.HipObjectsHostInfoCriteriaClientVersion", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.HipObjectsHostInfoCriteriaClientVersion{}.AttrType(), data)
}

// --- Unpacker for HipObjectsHostInfoCriteriaOs ---
func unpackHipObjectsHostInfoCriteriaOsToSdk(ctx context.Context, obj types.Object) (*objects.HipObjectsHostInfoCriteriaOs, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.HipObjectsHostInfoCriteriaOs", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.HipObjectsHostInfoCriteriaOs
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.HipObjectsHostInfoCriteriaOs
	var d diag.Diagnostics
	// Handling Objects
	if !model.Contains.IsNull() && !model.Contains.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Contains")
		unpacked, d := unpackHipObjectsHostInfoCriteriaOsContainsToSdk(ctx, model.Contains)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Contains"})
		}
		if unpacked != nil {
			sdk.Contains = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.HipObjectsHostInfoCriteriaOs", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for HipObjectsHostInfoCriteriaOs ---
func packHipObjectsHostInfoCriteriaOsFromSdk(ctx context.Context, sdk objects.HipObjectsHostInfoCriteriaOs) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.HipObjectsHostInfoCriteriaOs", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.HipObjectsHostInfoCriteriaOs
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Contains != nil {
		tflog.Debug(ctx, "Packing nested object for field Contains")
		packed, d := packHipObjectsHostInfoCriteriaOsContainsFromSdk(ctx, *sdk.Contains)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Contains"})
		}
		model.Contains = packed
	} else {
		model.Contains = basetypes.NewObjectNull(models.HipObjectsHostInfoCriteriaOsContains{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.HipObjectsHostInfoCriteriaOs{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.HipObjectsHostInfoCriteriaOs", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for HipObjectsHostInfoCriteriaOs ---
func unpackHipObjectsHostInfoCriteriaOsListToSdk(ctx context.Context, list types.List) ([]objects.HipObjectsHostInfoCriteriaOs, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.HipObjectsHostInfoCriteriaOs")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsHostInfoCriteriaOs
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.HipObjectsHostInfoCriteriaOs, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.HipObjectsHostInfoCriteriaOs{}.AttrTypes(), &item)
		unpacked, d := unpackHipObjectsHostInfoCriteriaOsToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.HipObjectsHostInfoCriteriaOs", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for HipObjectsHostInfoCriteriaOs ---
func packHipObjectsHostInfoCriteriaOsListFromSdk(ctx context.Context, sdks []objects.HipObjectsHostInfoCriteriaOs) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.HipObjectsHostInfoCriteriaOs")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsHostInfoCriteriaOs

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.HipObjectsHostInfoCriteriaOs
		obj, d := packHipObjectsHostInfoCriteriaOsFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.HipObjectsHostInfoCriteriaOs{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.HipObjectsHostInfoCriteriaOs", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.HipObjectsHostInfoCriteriaOs{}.AttrType(), data)
}

// --- Unpacker for HipObjectsHostInfoCriteriaOsContains ---
func unpackHipObjectsHostInfoCriteriaOsContainsToSdk(ctx context.Context, obj types.Object) (*objects.HipObjectsHostInfoCriteriaOsContains, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.HipObjectsHostInfoCriteriaOsContains", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.HipObjectsHostInfoCriteriaOsContains
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.HipObjectsHostInfoCriteriaOsContains
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Apple.IsNull() && !model.Apple.IsUnknown() {
		sdk.Apple = model.Apple.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Apple", "value": *sdk.Apple})
	}

	// Handling Primitives
	if !model.Google.IsNull() && !model.Google.IsUnknown() {
		sdk.Google = model.Google.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Google", "value": *sdk.Google})
	}

	// Handling Primitives
	if !model.Linux.IsNull() && !model.Linux.IsUnknown() {
		sdk.Linux = model.Linux.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Linux", "value": *sdk.Linux})
	}

	// Handling Primitives
	if !model.Microsoft.IsNull() && !model.Microsoft.IsUnknown() {
		sdk.Microsoft = model.Microsoft.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Microsoft", "value": *sdk.Microsoft})
	}

	// Handling Primitives
	if !model.Other.IsNull() && !model.Other.IsUnknown() {
		sdk.Other = model.Other.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Other", "value": *sdk.Other})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.HipObjectsHostInfoCriteriaOsContains", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for HipObjectsHostInfoCriteriaOsContains ---
func packHipObjectsHostInfoCriteriaOsContainsFromSdk(ctx context.Context, sdk objects.HipObjectsHostInfoCriteriaOsContains) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.HipObjectsHostInfoCriteriaOsContains", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.HipObjectsHostInfoCriteriaOsContains
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Apple != nil {
		model.Apple = basetypes.NewStringValue(*sdk.Apple)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Apple", "value": *sdk.Apple})
	} else {
		model.Apple = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Google != nil {
		model.Google = basetypes.NewStringValue(*sdk.Google)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Google", "value": *sdk.Google})
	} else {
		model.Google = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Linux != nil {
		model.Linux = basetypes.NewStringValue(*sdk.Linux)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Linux", "value": *sdk.Linux})
	} else {
		model.Linux = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Microsoft != nil {
		model.Microsoft = basetypes.NewStringValue(*sdk.Microsoft)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Microsoft", "value": *sdk.Microsoft})
	} else {
		model.Microsoft = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Other != nil {
		model.Other = basetypes.NewStringValue(*sdk.Other)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Other", "value": *sdk.Other})
	} else {
		model.Other = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.HipObjectsHostInfoCriteriaOsContains{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.HipObjectsHostInfoCriteriaOsContains", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for HipObjectsHostInfoCriteriaOsContains ---
func unpackHipObjectsHostInfoCriteriaOsContainsListToSdk(ctx context.Context, list types.List) ([]objects.HipObjectsHostInfoCriteriaOsContains, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.HipObjectsHostInfoCriteriaOsContains")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsHostInfoCriteriaOsContains
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.HipObjectsHostInfoCriteriaOsContains, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.HipObjectsHostInfoCriteriaOsContains{}.AttrTypes(), &item)
		unpacked, d := unpackHipObjectsHostInfoCriteriaOsContainsToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.HipObjectsHostInfoCriteriaOsContains", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for HipObjectsHostInfoCriteriaOsContains ---
func packHipObjectsHostInfoCriteriaOsContainsListFromSdk(ctx context.Context, sdks []objects.HipObjectsHostInfoCriteriaOsContains) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.HipObjectsHostInfoCriteriaOsContains")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsHostInfoCriteriaOsContains

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.HipObjectsHostInfoCriteriaOsContains
		obj, d := packHipObjectsHostInfoCriteriaOsContainsFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.HipObjectsHostInfoCriteriaOsContains{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.HipObjectsHostInfoCriteriaOsContains", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.HipObjectsHostInfoCriteriaOsContains{}.AttrType(), data)
}

// --- Unpacker for HipObjectsMobileDevice ---
func unpackHipObjectsMobileDeviceToSdk(ctx context.Context, obj types.Object) (*objects.HipObjectsMobileDevice, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.HipObjectsMobileDevice", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.HipObjectsMobileDevice
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.HipObjectsMobileDevice
	var d diag.Diagnostics
	// Handling Objects
	if !model.Criteria.IsNull() && !model.Criteria.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Criteria")
		unpacked, d := unpackHipObjectsMobileDeviceCriteriaToSdk(ctx, model.Criteria)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Criteria"})
		}
		if unpacked != nil {
			sdk.Criteria = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.HipObjectsMobileDevice", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for HipObjectsMobileDevice ---
func packHipObjectsMobileDeviceFromSdk(ctx context.Context, sdk objects.HipObjectsMobileDevice) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.HipObjectsMobileDevice", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.HipObjectsMobileDevice
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Criteria != nil {
		tflog.Debug(ctx, "Packing nested object for field Criteria")
		packed, d := packHipObjectsMobileDeviceCriteriaFromSdk(ctx, *sdk.Criteria)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Criteria"})
		}
		model.Criteria = packed
	} else {
		model.Criteria = basetypes.NewObjectNull(models.HipObjectsMobileDeviceCriteria{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.HipObjectsMobileDevice{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.HipObjectsMobileDevice", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for HipObjectsMobileDevice ---
func unpackHipObjectsMobileDeviceListToSdk(ctx context.Context, list types.List) ([]objects.HipObjectsMobileDevice, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.HipObjectsMobileDevice")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsMobileDevice
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.HipObjectsMobileDevice, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.HipObjectsMobileDevice{}.AttrTypes(), &item)
		unpacked, d := unpackHipObjectsMobileDeviceToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.HipObjectsMobileDevice", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for HipObjectsMobileDevice ---
func packHipObjectsMobileDeviceListFromSdk(ctx context.Context, sdks []objects.HipObjectsMobileDevice) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.HipObjectsMobileDevice")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsMobileDevice

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.HipObjectsMobileDevice
		obj, d := packHipObjectsMobileDeviceFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.HipObjectsMobileDevice{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.HipObjectsMobileDevice", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.HipObjectsMobileDevice{}.AttrType(), data)
}

// --- Unpacker for HipObjectsMobileDeviceCriteria ---
func unpackHipObjectsMobileDeviceCriteriaToSdk(ctx context.Context, obj types.Object) (*objects.HipObjectsMobileDeviceCriteria, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.HipObjectsMobileDeviceCriteria", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.HipObjectsMobileDeviceCriteria
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.HipObjectsMobileDeviceCriteria
	var d diag.Diagnostics
	// Handling Objects
	if !model.Applications.IsNull() && !model.Applications.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Applications")
		unpacked, d := unpackHipObjectsMobileDeviceCriteriaApplicationsToSdk(ctx, model.Applications)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Applications"})
		}
		if unpacked != nil {
			sdk.Applications = unpacked
		}
	}

	// Handling Primitives
	if !model.DiskEncrypted.IsNull() && !model.DiskEncrypted.IsUnknown() {
		sdk.DiskEncrypted = model.DiskEncrypted.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DiskEncrypted", "value": *sdk.DiskEncrypted})
	}

	// Handling Objects
	if !model.Imei.IsNull() && !model.Imei.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Imei")
		unpacked, d := unpackHipObjectsHostInfoCriteriaClientVersionToSdk(ctx, model.Imei)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Imei"})
		}
		if unpacked != nil {
			sdk.Imei = unpacked
		}
	}

	// Handling Primitives
	if !model.Jailbroken.IsNull() && !model.Jailbroken.IsUnknown() {
		sdk.Jailbroken = model.Jailbroken.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Jailbroken", "value": *sdk.Jailbroken})
	}

	// Handling Objects
	if !model.LastCheckinTime.IsNull() && !model.LastCheckinTime.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field LastCheckinTime")
		unpacked, d := unpackHipObjectsMobileDeviceCriteriaLastCheckinTimeToSdk(ctx, model.LastCheckinTime)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "LastCheckinTime"})
		}
		if unpacked != nil {
			sdk.LastCheckinTime = unpacked
		}
	}

	// Handling Objects
	if !model.Model.IsNull() && !model.Model.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Model")
		unpacked, d := unpackHipObjectsHostInfoCriteriaClientVersionToSdk(ctx, model.Model)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Model"})
		}
		if unpacked != nil {
			sdk.Model = unpacked
		}
	}

	// Handling Primitives
	if !model.PasscodeSet.IsNull() && !model.PasscodeSet.IsUnknown() {
		sdk.PasscodeSet = model.PasscodeSet.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "PasscodeSet", "value": *sdk.PasscodeSet})
	}

	// Handling Objects
	if !model.PhoneNumber.IsNull() && !model.PhoneNumber.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field PhoneNumber")
		unpacked, d := unpackHipObjectsHostInfoCriteriaClientVersionToSdk(ctx, model.PhoneNumber)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "PhoneNumber"})
		}
		if unpacked != nil {
			sdk.PhoneNumber = unpacked
		}
	}

	// Handling Objects
	if !model.Tag.IsNull() && !model.Tag.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Tag")
		unpacked, d := unpackHipObjectsHostInfoCriteriaClientVersionToSdk(ctx, model.Tag)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Tag"})
		}
		if unpacked != nil {
			sdk.Tag = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.HipObjectsMobileDeviceCriteria", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for HipObjectsMobileDeviceCriteria ---
func packHipObjectsMobileDeviceCriteriaFromSdk(ctx context.Context, sdk objects.HipObjectsMobileDeviceCriteria) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.HipObjectsMobileDeviceCriteria", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.HipObjectsMobileDeviceCriteria
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Applications != nil {
		tflog.Debug(ctx, "Packing nested object for field Applications")
		packed, d := packHipObjectsMobileDeviceCriteriaApplicationsFromSdk(ctx, *sdk.Applications)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Applications"})
		}
		model.Applications = packed
	} else {
		model.Applications = basetypes.NewObjectNull(models.HipObjectsMobileDeviceCriteriaApplications{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.DiskEncrypted != nil {
		model.DiskEncrypted = basetypes.NewBoolValue(*sdk.DiskEncrypted)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DiskEncrypted", "value": *sdk.DiskEncrypted})
	} else {
		model.DiskEncrypted = basetypes.NewBoolNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Imei != nil {
		tflog.Debug(ctx, "Packing nested object for field Imei")
		packed, d := packHipObjectsHostInfoCriteriaClientVersionFromSdk(ctx, *sdk.Imei)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Imei"})
		}
		model.Imei = packed
	} else {
		model.Imei = basetypes.NewObjectNull(models.HipObjectsHostInfoCriteriaClientVersion{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Jailbroken != nil {
		model.Jailbroken = basetypes.NewBoolValue(*sdk.Jailbroken)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Jailbroken", "value": *sdk.Jailbroken})
	} else {
		model.Jailbroken = basetypes.NewBoolNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.LastCheckinTime != nil {
		tflog.Debug(ctx, "Packing nested object for field LastCheckinTime")
		packed, d := packHipObjectsMobileDeviceCriteriaLastCheckinTimeFromSdk(ctx, *sdk.LastCheckinTime)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "LastCheckinTime"})
		}
		model.LastCheckinTime = packed
	} else {
		model.LastCheckinTime = basetypes.NewObjectNull(models.HipObjectsMobileDeviceCriteriaLastCheckinTime{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Model != nil {
		tflog.Debug(ctx, "Packing nested object for field Model")
		packed, d := packHipObjectsHostInfoCriteriaClientVersionFromSdk(ctx, *sdk.Model)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Model"})
		}
		model.Model = packed
	} else {
		model.Model = basetypes.NewObjectNull(models.HipObjectsHostInfoCriteriaClientVersion{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.PasscodeSet != nil {
		model.PasscodeSet = basetypes.NewBoolValue(*sdk.PasscodeSet)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "PasscodeSet", "value": *sdk.PasscodeSet})
	} else {
		model.PasscodeSet = basetypes.NewBoolNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.PhoneNumber != nil {
		tflog.Debug(ctx, "Packing nested object for field PhoneNumber")
		packed, d := packHipObjectsHostInfoCriteriaClientVersionFromSdk(ctx, *sdk.PhoneNumber)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "PhoneNumber"})
		}
		model.PhoneNumber = packed
	} else {
		model.PhoneNumber = basetypes.NewObjectNull(models.HipObjectsHostInfoCriteriaClientVersion{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Tag != nil {
		tflog.Debug(ctx, "Packing nested object for field Tag")
		packed, d := packHipObjectsHostInfoCriteriaClientVersionFromSdk(ctx, *sdk.Tag)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Tag"})
		}
		model.Tag = packed
	} else {
		model.Tag = basetypes.NewObjectNull(models.HipObjectsHostInfoCriteriaClientVersion{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.HipObjectsMobileDeviceCriteria{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.HipObjectsMobileDeviceCriteria", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for HipObjectsMobileDeviceCriteria ---
func unpackHipObjectsMobileDeviceCriteriaListToSdk(ctx context.Context, list types.List) ([]objects.HipObjectsMobileDeviceCriteria, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.HipObjectsMobileDeviceCriteria")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsMobileDeviceCriteria
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.HipObjectsMobileDeviceCriteria, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.HipObjectsMobileDeviceCriteria{}.AttrTypes(), &item)
		unpacked, d := unpackHipObjectsMobileDeviceCriteriaToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.HipObjectsMobileDeviceCriteria", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for HipObjectsMobileDeviceCriteria ---
func packHipObjectsMobileDeviceCriteriaListFromSdk(ctx context.Context, sdks []objects.HipObjectsMobileDeviceCriteria) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.HipObjectsMobileDeviceCriteria")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsMobileDeviceCriteria

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.HipObjectsMobileDeviceCriteria
		obj, d := packHipObjectsMobileDeviceCriteriaFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.HipObjectsMobileDeviceCriteria{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.HipObjectsMobileDeviceCriteria", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.HipObjectsMobileDeviceCriteria{}.AttrType(), data)
}

// --- Unpacker for HipObjectsMobileDeviceCriteriaApplications ---
func unpackHipObjectsMobileDeviceCriteriaApplicationsToSdk(ctx context.Context, obj types.Object) (*objects.HipObjectsMobileDeviceCriteriaApplications, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.HipObjectsMobileDeviceCriteriaApplications", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.HipObjectsMobileDeviceCriteriaApplications
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.HipObjectsMobileDeviceCriteriaApplications
	var d diag.Diagnostics
	// Handling Objects
	if !model.HasMalware.IsNull() && !model.HasMalware.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field HasMalware")
		unpacked, d := unpackHipObjectsMobileDeviceCriteriaApplicationsHasMalwareToSdk(ctx, model.HasMalware)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "HasMalware"})
		}
		if unpacked != nil {
			sdk.HasMalware = unpacked
		}
	}

	// Handling Primitives
	if !model.HasUnmanagedApp.IsNull() && !model.HasUnmanagedApp.IsUnknown() {
		sdk.HasUnmanagedApp = model.HasUnmanagedApp.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "HasUnmanagedApp", "value": *sdk.HasUnmanagedApp})
	}

	// Handling Lists
	if !model.Includes.IsNull() && !model.Includes.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field Includes")
		unpacked, d := unpackHipObjectsMobileDeviceCriteriaApplicationsHasMalwareYesExcludesInnerListToSdk(ctx, model.Includes)
		diags.Append(d...)
		sdk.Includes = unpacked
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.HipObjectsMobileDeviceCriteriaApplications", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for HipObjectsMobileDeviceCriteriaApplications ---
func packHipObjectsMobileDeviceCriteriaApplicationsFromSdk(ctx context.Context, sdk objects.HipObjectsMobileDeviceCriteriaApplications) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.HipObjectsMobileDeviceCriteriaApplications", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.HipObjectsMobileDeviceCriteriaApplications
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.HasMalware != nil {
		tflog.Debug(ctx, "Packing nested object for field HasMalware")
		packed, d := packHipObjectsMobileDeviceCriteriaApplicationsHasMalwareFromSdk(ctx, *sdk.HasMalware)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "HasMalware"})
		}
		model.HasMalware = packed
	} else {
		model.HasMalware = basetypes.NewObjectNull(models.HipObjectsMobileDeviceCriteriaApplicationsHasMalware{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.HasUnmanagedApp != nil {
		model.HasUnmanagedApp = basetypes.NewBoolValue(*sdk.HasUnmanagedApp)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "HasUnmanagedApp", "value": *sdk.HasUnmanagedApp})
	} else {
		model.HasUnmanagedApp = basetypes.NewBoolNull()
	}
	// Handling Lists
	if sdk.Includes != nil {
		tflog.Debug(ctx, "Packing list of objects for field Includes")
		packed, d := packHipObjectsMobileDeviceCriteriaApplicationsHasMalwareYesExcludesInnerListFromSdk(ctx, sdk.Includes)
		diags.Append(d...)
		model.Includes = packed
	} else {
		model.Includes = basetypes.NewListNull(models.HipObjectsMobileDeviceCriteriaApplicationsHasMalwareYesExcludesInner{}.AttrType())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.HipObjectsMobileDeviceCriteriaApplications{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.HipObjectsMobileDeviceCriteriaApplications", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for HipObjectsMobileDeviceCriteriaApplications ---
func unpackHipObjectsMobileDeviceCriteriaApplicationsListToSdk(ctx context.Context, list types.List) ([]objects.HipObjectsMobileDeviceCriteriaApplications, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.HipObjectsMobileDeviceCriteriaApplications")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsMobileDeviceCriteriaApplications
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.HipObjectsMobileDeviceCriteriaApplications, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.HipObjectsMobileDeviceCriteriaApplications{}.AttrTypes(), &item)
		unpacked, d := unpackHipObjectsMobileDeviceCriteriaApplicationsToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.HipObjectsMobileDeviceCriteriaApplications", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for HipObjectsMobileDeviceCriteriaApplications ---
func packHipObjectsMobileDeviceCriteriaApplicationsListFromSdk(ctx context.Context, sdks []objects.HipObjectsMobileDeviceCriteriaApplications) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.HipObjectsMobileDeviceCriteriaApplications")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsMobileDeviceCriteriaApplications

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.HipObjectsMobileDeviceCriteriaApplications
		obj, d := packHipObjectsMobileDeviceCriteriaApplicationsFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.HipObjectsMobileDeviceCriteriaApplications{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.HipObjectsMobileDeviceCriteriaApplications", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.HipObjectsMobileDeviceCriteriaApplications{}.AttrType(), data)
}

// --- Unpacker for HipObjectsMobileDeviceCriteriaApplicationsHasMalware ---
func unpackHipObjectsMobileDeviceCriteriaApplicationsHasMalwareToSdk(ctx context.Context, obj types.Object) (*objects.HipObjectsMobileDeviceCriteriaApplicationsHasMalware, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.HipObjectsMobileDeviceCriteriaApplicationsHasMalware", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.HipObjectsMobileDeviceCriteriaApplicationsHasMalware
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.HipObjectsMobileDeviceCriteriaApplicationsHasMalware
	var d diag.Diagnostics
	// Handling Typeless Objects
	if !model.No.IsNull() && !model.No.IsUnknown() {
		tflog.Debug(ctx, "Unpacking typeless object for field No")
		sdk.No = make(map[string]interface{})
	}

	// Handling Objects
	if !model.Yes.IsNull() && !model.Yes.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Yes")
		unpacked, d := unpackHipObjectsMobileDeviceCriteriaApplicationsHasMalwareYesToSdk(ctx, model.Yes)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Yes"})
		}
		if unpacked != nil {
			sdk.Yes = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.HipObjectsMobileDeviceCriteriaApplicationsHasMalware", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for HipObjectsMobileDeviceCriteriaApplicationsHasMalware ---
func packHipObjectsMobileDeviceCriteriaApplicationsHasMalwareFromSdk(ctx context.Context, sdk objects.HipObjectsMobileDeviceCriteriaApplicationsHasMalware) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.HipObjectsMobileDeviceCriteriaApplicationsHasMalware", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.HipObjectsMobileDeviceCriteriaApplicationsHasMalware
	var d diag.Diagnostics
	// Handling Objects
	// This is a marker object (e.g. CHAP: {}). We just need to create an empty, non-null object.
	if sdk.No != nil && !reflect.ValueOf(sdk.No).IsNil() {
		tflog.Debug(ctx, "Packing typeless object for field No")
		var d diag.Diagnostics
		// Create an empty object with no attributes, which signifies its presence.
		model.No, d = basetypes.NewObjectValue(map[string]attr.Type{}, map[string]attr.Value{})
		diags.Append(d...)
	} else {
		// Since this field is part of a oneOf, being nil means it's not selected.
		// We make the object null with an empty attribute map.
		model.No = basetypes.NewObjectNull(map[string]attr.Type{})
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Yes != nil {
		tflog.Debug(ctx, "Packing nested object for field Yes")
		packed, d := packHipObjectsMobileDeviceCriteriaApplicationsHasMalwareYesFromSdk(ctx, *sdk.Yes)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Yes"})
		}
		model.Yes = packed
	} else {
		model.Yes = basetypes.NewObjectNull(models.HipObjectsMobileDeviceCriteriaApplicationsHasMalwareYes{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.HipObjectsMobileDeviceCriteriaApplicationsHasMalware{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.HipObjectsMobileDeviceCriteriaApplicationsHasMalware", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for HipObjectsMobileDeviceCriteriaApplicationsHasMalware ---
func unpackHipObjectsMobileDeviceCriteriaApplicationsHasMalwareListToSdk(ctx context.Context, list types.List) ([]objects.HipObjectsMobileDeviceCriteriaApplicationsHasMalware, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.HipObjectsMobileDeviceCriteriaApplicationsHasMalware")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsMobileDeviceCriteriaApplicationsHasMalware
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.HipObjectsMobileDeviceCriteriaApplicationsHasMalware, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.HipObjectsMobileDeviceCriteriaApplicationsHasMalware{}.AttrTypes(), &item)
		unpacked, d := unpackHipObjectsMobileDeviceCriteriaApplicationsHasMalwareToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.HipObjectsMobileDeviceCriteriaApplicationsHasMalware", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for HipObjectsMobileDeviceCriteriaApplicationsHasMalware ---
func packHipObjectsMobileDeviceCriteriaApplicationsHasMalwareListFromSdk(ctx context.Context, sdks []objects.HipObjectsMobileDeviceCriteriaApplicationsHasMalware) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.HipObjectsMobileDeviceCriteriaApplicationsHasMalware")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsMobileDeviceCriteriaApplicationsHasMalware

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.HipObjectsMobileDeviceCriteriaApplicationsHasMalware
		obj, d := packHipObjectsMobileDeviceCriteriaApplicationsHasMalwareFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.HipObjectsMobileDeviceCriteriaApplicationsHasMalware{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.HipObjectsMobileDeviceCriteriaApplicationsHasMalware", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.HipObjectsMobileDeviceCriteriaApplicationsHasMalware{}.AttrType(), data)
}

// --- Unpacker for HipObjectsMobileDeviceCriteriaApplicationsHasMalwareYes ---
func unpackHipObjectsMobileDeviceCriteriaApplicationsHasMalwareYesToSdk(ctx context.Context, obj types.Object) (*objects.HipObjectsMobileDeviceCriteriaApplicationsHasMalwareYes, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.HipObjectsMobileDeviceCriteriaApplicationsHasMalwareYes", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.HipObjectsMobileDeviceCriteriaApplicationsHasMalwareYes
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.HipObjectsMobileDeviceCriteriaApplicationsHasMalwareYes
	var d diag.Diagnostics
	// Handling Lists
	if !model.Excludes.IsNull() && !model.Excludes.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field Excludes")
		unpacked, d := unpackHipObjectsMobileDeviceCriteriaApplicationsHasMalwareYesExcludesInnerListToSdk(ctx, model.Excludes)
		diags.Append(d...)
		sdk.Excludes = unpacked
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.HipObjectsMobileDeviceCriteriaApplicationsHasMalwareYes", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for HipObjectsMobileDeviceCriteriaApplicationsHasMalwareYes ---
func packHipObjectsMobileDeviceCriteriaApplicationsHasMalwareYesFromSdk(ctx context.Context, sdk objects.HipObjectsMobileDeviceCriteriaApplicationsHasMalwareYes) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.HipObjectsMobileDeviceCriteriaApplicationsHasMalwareYes", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.HipObjectsMobileDeviceCriteriaApplicationsHasMalwareYes
	var d diag.Diagnostics
	// Handling Lists
	if sdk.Excludes != nil {
		tflog.Debug(ctx, "Packing list of objects for field Excludes")
		packed, d := packHipObjectsMobileDeviceCriteriaApplicationsHasMalwareYesExcludesInnerListFromSdk(ctx, sdk.Excludes)
		diags.Append(d...)
		model.Excludes = packed
	} else {
		model.Excludes = basetypes.NewListNull(models.HipObjectsMobileDeviceCriteriaApplicationsHasMalwareYesExcludesInner{}.AttrType())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.HipObjectsMobileDeviceCriteriaApplicationsHasMalwareYes{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.HipObjectsMobileDeviceCriteriaApplicationsHasMalwareYes", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for HipObjectsMobileDeviceCriteriaApplicationsHasMalwareYes ---
func unpackHipObjectsMobileDeviceCriteriaApplicationsHasMalwareYesListToSdk(ctx context.Context, list types.List) ([]objects.HipObjectsMobileDeviceCriteriaApplicationsHasMalwareYes, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.HipObjectsMobileDeviceCriteriaApplicationsHasMalwareYes")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsMobileDeviceCriteriaApplicationsHasMalwareYes
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.HipObjectsMobileDeviceCriteriaApplicationsHasMalwareYes, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.HipObjectsMobileDeviceCriteriaApplicationsHasMalwareYes{}.AttrTypes(), &item)
		unpacked, d := unpackHipObjectsMobileDeviceCriteriaApplicationsHasMalwareYesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.HipObjectsMobileDeviceCriteriaApplicationsHasMalwareYes", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for HipObjectsMobileDeviceCriteriaApplicationsHasMalwareYes ---
func packHipObjectsMobileDeviceCriteriaApplicationsHasMalwareYesListFromSdk(ctx context.Context, sdks []objects.HipObjectsMobileDeviceCriteriaApplicationsHasMalwareYes) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.HipObjectsMobileDeviceCriteriaApplicationsHasMalwareYes")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsMobileDeviceCriteriaApplicationsHasMalwareYes

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.HipObjectsMobileDeviceCriteriaApplicationsHasMalwareYes
		obj, d := packHipObjectsMobileDeviceCriteriaApplicationsHasMalwareYesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.HipObjectsMobileDeviceCriteriaApplicationsHasMalwareYes{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.HipObjectsMobileDeviceCriteriaApplicationsHasMalwareYes", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.HipObjectsMobileDeviceCriteriaApplicationsHasMalwareYes{}.AttrType(), data)
}

// --- Unpacker for HipObjectsMobileDeviceCriteriaApplicationsHasMalwareYesExcludesInner ---
func unpackHipObjectsMobileDeviceCriteriaApplicationsHasMalwareYesExcludesInnerToSdk(ctx context.Context, obj types.Object) (*objects.HipObjectsMobileDeviceCriteriaApplicationsHasMalwareYesExcludesInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.HipObjectsMobileDeviceCriteriaApplicationsHasMalwareYesExcludesInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.HipObjectsMobileDeviceCriteriaApplicationsHasMalwareYesExcludesInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.HipObjectsMobileDeviceCriteriaApplicationsHasMalwareYesExcludesInner
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Hash.IsNull() && !model.Hash.IsUnknown() {
		sdk.Hash = model.Hash.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Hash", "value": *sdk.Hash})
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	// Handling Primitives
	if !model.Package.IsNull() && !model.Package.IsUnknown() {
		sdk.Package = model.Package.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Package", "value": *sdk.Package})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.HipObjectsMobileDeviceCriteriaApplicationsHasMalwareYesExcludesInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for HipObjectsMobileDeviceCriteriaApplicationsHasMalwareYesExcludesInner ---
func packHipObjectsMobileDeviceCriteriaApplicationsHasMalwareYesExcludesInnerFromSdk(ctx context.Context, sdk objects.HipObjectsMobileDeviceCriteriaApplicationsHasMalwareYesExcludesInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.HipObjectsMobileDeviceCriteriaApplicationsHasMalwareYesExcludesInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.HipObjectsMobileDeviceCriteriaApplicationsHasMalwareYesExcludesInner
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Hash != nil {
		model.Hash = basetypes.NewStringValue(*sdk.Hash)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Hash", "value": *sdk.Hash})
	} else {
		model.Hash = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	// Handling Primitives
	// Standard primitive packing
	if sdk.Package != nil {
		model.Package = basetypes.NewStringValue(*sdk.Package)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Package", "value": *sdk.Package})
	} else {
		model.Package = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.HipObjectsMobileDeviceCriteriaApplicationsHasMalwareYesExcludesInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.HipObjectsMobileDeviceCriteriaApplicationsHasMalwareYesExcludesInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for HipObjectsMobileDeviceCriteriaApplicationsHasMalwareYesExcludesInner ---
func unpackHipObjectsMobileDeviceCriteriaApplicationsHasMalwareYesExcludesInnerListToSdk(ctx context.Context, list types.List) ([]objects.HipObjectsMobileDeviceCriteriaApplicationsHasMalwareYesExcludesInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.HipObjectsMobileDeviceCriteriaApplicationsHasMalwareYesExcludesInner")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsMobileDeviceCriteriaApplicationsHasMalwareYesExcludesInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.HipObjectsMobileDeviceCriteriaApplicationsHasMalwareYesExcludesInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.HipObjectsMobileDeviceCriteriaApplicationsHasMalwareYesExcludesInner{}.AttrTypes(), &item)
		unpacked, d := unpackHipObjectsMobileDeviceCriteriaApplicationsHasMalwareYesExcludesInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.HipObjectsMobileDeviceCriteriaApplicationsHasMalwareYesExcludesInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for HipObjectsMobileDeviceCriteriaApplicationsHasMalwareYesExcludesInner ---
func packHipObjectsMobileDeviceCriteriaApplicationsHasMalwareYesExcludesInnerListFromSdk(ctx context.Context, sdks []objects.HipObjectsMobileDeviceCriteriaApplicationsHasMalwareYesExcludesInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.HipObjectsMobileDeviceCriteriaApplicationsHasMalwareYesExcludesInner")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsMobileDeviceCriteriaApplicationsHasMalwareYesExcludesInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.HipObjectsMobileDeviceCriteriaApplicationsHasMalwareYesExcludesInner
		obj, d := packHipObjectsMobileDeviceCriteriaApplicationsHasMalwareYesExcludesInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.HipObjectsMobileDeviceCriteriaApplicationsHasMalwareYesExcludesInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.HipObjectsMobileDeviceCriteriaApplicationsHasMalwareYesExcludesInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.HipObjectsMobileDeviceCriteriaApplicationsHasMalwareYesExcludesInner{}.AttrType(), data)
}

// --- Unpacker for HipObjectsMobileDeviceCriteriaLastCheckinTime ---
func unpackHipObjectsMobileDeviceCriteriaLastCheckinTimeToSdk(ctx context.Context, obj types.Object) (*objects.HipObjectsMobileDeviceCriteriaLastCheckinTime, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.HipObjectsMobileDeviceCriteriaLastCheckinTime", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.HipObjectsMobileDeviceCriteriaLastCheckinTime
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.HipObjectsMobileDeviceCriteriaLastCheckinTime
	var d diag.Diagnostics
	// Handling Objects
	if !model.NotWithin.IsNull() && !model.NotWithin.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field NotWithin")
		unpacked, d := unpackHipObjectsMobileDeviceCriteriaLastCheckinTimeNotWithinToSdk(ctx, model.NotWithin)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "NotWithin"})
		}
		if unpacked != nil {
			sdk.NotWithin = unpacked
		}
	}

	// Handling Objects
	if !model.Within.IsNull() && !model.Within.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Within")
		unpacked, d := unpackHipObjectsMobileDeviceCriteriaLastCheckinTimeNotWithinToSdk(ctx, model.Within)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Within"})
		}
		if unpacked != nil {
			sdk.Within = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.HipObjectsMobileDeviceCriteriaLastCheckinTime", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for HipObjectsMobileDeviceCriteriaLastCheckinTime ---
func packHipObjectsMobileDeviceCriteriaLastCheckinTimeFromSdk(ctx context.Context, sdk objects.HipObjectsMobileDeviceCriteriaLastCheckinTime) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.HipObjectsMobileDeviceCriteriaLastCheckinTime", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.HipObjectsMobileDeviceCriteriaLastCheckinTime
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.NotWithin != nil {
		tflog.Debug(ctx, "Packing nested object for field NotWithin")
		packed, d := packHipObjectsMobileDeviceCriteriaLastCheckinTimeNotWithinFromSdk(ctx, *sdk.NotWithin)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "NotWithin"})
		}
		model.NotWithin = packed
	} else {
		model.NotWithin = basetypes.NewObjectNull(models.HipObjectsMobileDeviceCriteriaLastCheckinTimeNotWithin{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Within != nil {
		tflog.Debug(ctx, "Packing nested object for field Within")
		packed, d := packHipObjectsMobileDeviceCriteriaLastCheckinTimeNotWithinFromSdk(ctx, *sdk.Within)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Within"})
		}
		model.Within = packed
	} else {
		model.Within = basetypes.NewObjectNull(models.HipObjectsMobileDeviceCriteriaLastCheckinTimeNotWithin{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.HipObjectsMobileDeviceCriteriaLastCheckinTime{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.HipObjectsMobileDeviceCriteriaLastCheckinTime", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for HipObjectsMobileDeviceCriteriaLastCheckinTime ---
func unpackHipObjectsMobileDeviceCriteriaLastCheckinTimeListToSdk(ctx context.Context, list types.List) ([]objects.HipObjectsMobileDeviceCriteriaLastCheckinTime, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.HipObjectsMobileDeviceCriteriaLastCheckinTime")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsMobileDeviceCriteriaLastCheckinTime
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.HipObjectsMobileDeviceCriteriaLastCheckinTime, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.HipObjectsMobileDeviceCriteriaLastCheckinTime{}.AttrTypes(), &item)
		unpacked, d := unpackHipObjectsMobileDeviceCriteriaLastCheckinTimeToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.HipObjectsMobileDeviceCriteriaLastCheckinTime", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for HipObjectsMobileDeviceCriteriaLastCheckinTime ---
func packHipObjectsMobileDeviceCriteriaLastCheckinTimeListFromSdk(ctx context.Context, sdks []objects.HipObjectsMobileDeviceCriteriaLastCheckinTime) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.HipObjectsMobileDeviceCriteriaLastCheckinTime")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsMobileDeviceCriteriaLastCheckinTime

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.HipObjectsMobileDeviceCriteriaLastCheckinTime
		obj, d := packHipObjectsMobileDeviceCriteriaLastCheckinTimeFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.HipObjectsMobileDeviceCriteriaLastCheckinTime{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.HipObjectsMobileDeviceCriteriaLastCheckinTime", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.HipObjectsMobileDeviceCriteriaLastCheckinTime{}.AttrType(), data)
}

// --- Unpacker for HipObjectsMobileDeviceCriteriaLastCheckinTimeNotWithin ---
func unpackHipObjectsMobileDeviceCriteriaLastCheckinTimeNotWithinToSdk(ctx context.Context, obj types.Object) (*objects.HipObjectsMobileDeviceCriteriaLastCheckinTimeNotWithin, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.HipObjectsMobileDeviceCriteriaLastCheckinTimeNotWithin", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.HipObjectsMobileDeviceCriteriaLastCheckinTimeNotWithin
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.HipObjectsMobileDeviceCriteriaLastCheckinTimeNotWithin
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Days.IsNull() && !model.Days.IsUnknown() {
		sdk.Days = int32(model.Days.ValueInt64())
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Days", "value": sdk.Days})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.HipObjectsMobileDeviceCriteriaLastCheckinTimeNotWithin", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for HipObjectsMobileDeviceCriteriaLastCheckinTimeNotWithin ---
func packHipObjectsMobileDeviceCriteriaLastCheckinTimeNotWithinFromSdk(ctx context.Context, sdk objects.HipObjectsMobileDeviceCriteriaLastCheckinTimeNotWithin) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.HipObjectsMobileDeviceCriteriaLastCheckinTimeNotWithin", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.HipObjectsMobileDeviceCriteriaLastCheckinTimeNotWithin
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.Days = basetypes.NewInt64Value(int64(sdk.Days))
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.HipObjectsMobileDeviceCriteriaLastCheckinTimeNotWithin{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.HipObjectsMobileDeviceCriteriaLastCheckinTimeNotWithin", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for HipObjectsMobileDeviceCriteriaLastCheckinTimeNotWithin ---
func unpackHipObjectsMobileDeviceCriteriaLastCheckinTimeNotWithinListToSdk(ctx context.Context, list types.List) ([]objects.HipObjectsMobileDeviceCriteriaLastCheckinTimeNotWithin, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.HipObjectsMobileDeviceCriteriaLastCheckinTimeNotWithin")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsMobileDeviceCriteriaLastCheckinTimeNotWithin
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.HipObjectsMobileDeviceCriteriaLastCheckinTimeNotWithin, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.HipObjectsMobileDeviceCriteriaLastCheckinTimeNotWithin{}.AttrTypes(), &item)
		unpacked, d := unpackHipObjectsMobileDeviceCriteriaLastCheckinTimeNotWithinToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.HipObjectsMobileDeviceCriteriaLastCheckinTimeNotWithin", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for HipObjectsMobileDeviceCriteriaLastCheckinTimeNotWithin ---
func packHipObjectsMobileDeviceCriteriaLastCheckinTimeNotWithinListFromSdk(ctx context.Context, sdks []objects.HipObjectsMobileDeviceCriteriaLastCheckinTimeNotWithin) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.HipObjectsMobileDeviceCriteriaLastCheckinTimeNotWithin")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsMobileDeviceCriteriaLastCheckinTimeNotWithin

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.HipObjectsMobileDeviceCriteriaLastCheckinTimeNotWithin
		obj, d := packHipObjectsMobileDeviceCriteriaLastCheckinTimeNotWithinFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.HipObjectsMobileDeviceCriteriaLastCheckinTimeNotWithin{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.HipObjectsMobileDeviceCriteriaLastCheckinTimeNotWithin", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.HipObjectsMobileDeviceCriteriaLastCheckinTimeNotWithin{}.AttrType(), data)
}

// --- Unpacker for HipObjectsNetworkInfo ---
func unpackHipObjectsNetworkInfoToSdk(ctx context.Context, obj types.Object) (*objects.HipObjectsNetworkInfo, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.HipObjectsNetworkInfo", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.HipObjectsNetworkInfo
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.HipObjectsNetworkInfo
	var d diag.Diagnostics
	// Handling Objects
	if !model.Criteria.IsNull() && !model.Criteria.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Criteria")
		unpacked, d := unpackHipObjectsNetworkInfoCriteriaToSdk(ctx, model.Criteria)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Criteria"})
		}
		if unpacked != nil {
			sdk.Criteria = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.HipObjectsNetworkInfo", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for HipObjectsNetworkInfo ---
func packHipObjectsNetworkInfoFromSdk(ctx context.Context, sdk objects.HipObjectsNetworkInfo) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.HipObjectsNetworkInfo", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.HipObjectsNetworkInfo
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Criteria != nil {
		tflog.Debug(ctx, "Packing nested object for field Criteria")
		packed, d := packHipObjectsNetworkInfoCriteriaFromSdk(ctx, *sdk.Criteria)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Criteria"})
		}
		model.Criteria = packed
	} else {
		model.Criteria = basetypes.NewObjectNull(models.HipObjectsNetworkInfoCriteria{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.HipObjectsNetworkInfo{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.HipObjectsNetworkInfo", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for HipObjectsNetworkInfo ---
func unpackHipObjectsNetworkInfoListToSdk(ctx context.Context, list types.List) ([]objects.HipObjectsNetworkInfo, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.HipObjectsNetworkInfo")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsNetworkInfo
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.HipObjectsNetworkInfo, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.HipObjectsNetworkInfo{}.AttrTypes(), &item)
		unpacked, d := unpackHipObjectsNetworkInfoToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.HipObjectsNetworkInfo", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for HipObjectsNetworkInfo ---
func packHipObjectsNetworkInfoListFromSdk(ctx context.Context, sdks []objects.HipObjectsNetworkInfo) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.HipObjectsNetworkInfo")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsNetworkInfo

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.HipObjectsNetworkInfo
		obj, d := packHipObjectsNetworkInfoFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.HipObjectsNetworkInfo{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.HipObjectsNetworkInfo", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.HipObjectsNetworkInfo{}.AttrType(), data)
}

// --- Unpacker for HipObjectsNetworkInfoCriteria ---
func unpackHipObjectsNetworkInfoCriteriaToSdk(ctx context.Context, obj types.Object) (*objects.HipObjectsNetworkInfoCriteria, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.HipObjectsNetworkInfoCriteria", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.HipObjectsNetworkInfoCriteria
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.HipObjectsNetworkInfoCriteria
	var d diag.Diagnostics
	// Handling Objects
	if !model.Network.IsNull() && !model.Network.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Network")
		unpacked, d := unpackHipObjectsNetworkInfoCriteriaNetworkToSdk(ctx, model.Network)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Network"})
		}
		if unpacked != nil {
			sdk.Network = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.HipObjectsNetworkInfoCriteria", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for HipObjectsNetworkInfoCriteria ---
func packHipObjectsNetworkInfoCriteriaFromSdk(ctx context.Context, sdk objects.HipObjectsNetworkInfoCriteria) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.HipObjectsNetworkInfoCriteria", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.HipObjectsNetworkInfoCriteria
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Network != nil {
		tflog.Debug(ctx, "Packing nested object for field Network")
		packed, d := packHipObjectsNetworkInfoCriteriaNetworkFromSdk(ctx, *sdk.Network)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Network"})
		}
		model.Network = packed
	} else {
		model.Network = basetypes.NewObjectNull(models.HipObjectsNetworkInfoCriteriaNetwork{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.HipObjectsNetworkInfoCriteria{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.HipObjectsNetworkInfoCriteria", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for HipObjectsNetworkInfoCriteria ---
func unpackHipObjectsNetworkInfoCriteriaListToSdk(ctx context.Context, list types.List) ([]objects.HipObjectsNetworkInfoCriteria, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.HipObjectsNetworkInfoCriteria")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsNetworkInfoCriteria
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.HipObjectsNetworkInfoCriteria, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.HipObjectsNetworkInfoCriteria{}.AttrTypes(), &item)
		unpacked, d := unpackHipObjectsNetworkInfoCriteriaToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.HipObjectsNetworkInfoCriteria", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for HipObjectsNetworkInfoCriteria ---
func packHipObjectsNetworkInfoCriteriaListFromSdk(ctx context.Context, sdks []objects.HipObjectsNetworkInfoCriteria) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.HipObjectsNetworkInfoCriteria")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsNetworkInfoCriteria

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.HipObjectsNetworkInfoCriteria
		obj, d := packHipObjectsNetworkInfoCriteriaFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.HipObjectsNetworkInfoCriteria{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.HipObjectsNetworkInfoCriteria", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.HipObjectsNetworkInfoCriteria{}.AttrType(), data)
}

// --- Unpacker for HipObjectsNetworkInfoCriteriaNetwork ---
func unpackHipObjectsNetworkInfoCriteriaNetworkToSdk(ctx context.Context, obj types.Object) (*objects.HipObjectsNetworkInfoCriteriaNetwork, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.HipObjectsNetworkInfoCriteriaNetwork", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.HipObjectsNetworkInfoCriteriaNetwork
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.HipObjectsNetworkInfoCriteriaNetwork
	var d diag.Diagnostics
	// Handling Objects
	if !model.Is.IsNull() && !model.Is.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Is")
		unpacked, d := unpackHipObjectsNetworkInfoCriteriaNetworkIsToSdk(ctx, model.Is)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Is"})
		}
		if unpacked != nil {
			sdk.Is = unpacked
		}
	}

	// Handling Objects
	if !model.IsNot.IsNull() && !model.IsNot.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field IsNot")
		unpacked, d := unpackHipObjectsNetworkInfoCriteriaNetworkIsNotToSdk(ctx, model.IsNot)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "IsNot"})
		}
		if unpacked != nil {
			sdk.IsNot = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.HipObjectsNetworkInfoCriteriaNetwork", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for HipObjectsNetworkInfoCriteriaNetwork ---
func packHipObjectsNetworkInfoCriteriaNetworkFromSdk(ctx context.Context, sdk objects.HipObjectsNetworkInfoCriteriaNetwork) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.HipObjectsNetworkInfoCriteriaNetwork", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.HipObjectsNetworkInfoCriteriaNetwork
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Is != nil {
		tflog.Debug(ctx, "Packing nested object for field Is")
		packed, d := packHipObjectsNetworkInfoCriteriaNetworkIsFromSdk(ctx, *sdk.Is)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Is"})
		}
		model.Is = packed
	} else {
		model.Is = basetypes.NewObjectNull(models.HipObjectsNetworkInfoCriteriaNetworkIs{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.IsNot != nil {
		tflog.Debug(ctx, "Packing nested object for field IsNot")
		packed, d := packHipObjectsNetworkInfoCriteriaNetworkIsNotFromSdk(ctx, *sdk.IsNot)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "IsNot"})
		}
		model.IsNot = packed
	} else {
		model.IsNot = basetypes.NewObjectNull(models.HipObjectsNetworkInfoCriteriaNetworkIsNot{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.HipObjectsNetworkInfoCriteriaNetwork{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.HipObjectsNetworkInfoCriteriaNetwork", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for HipObjectsNetworkInfoCriteriaNetwork ---
func unpackHipObjectsNetworkInfoCriteriaNetworkListToSdk(ctx context.Context, list types.List) ([]objects.HipObjectsNetworkInfoCriteriaNetwork, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.HipObjectsNetworkInfoCriteriaNetwork")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsNetworkInfoCriteriaNetwork
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.HipObjectsNetworkInfoCriteriaNetwork, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.HipObjectsNetworkInfoCriteriaNetwork{}.AttrTypes(), &item)
		unpacked, d := unpackHipObjectsNetworkInfoCriteriaNetworkToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.HipObjectsNetworkInfoCriteriaNetwork", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for HipObjectsNetworkInfoCriteriaNetwork ---
func packHipObjectsNetworkInfoCriteriaNetworkListFromSdk(ctx context.Context, sdks []objects.HipObjectsNetworkInfoCriteriaNetwork) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.HipObjectsNetworkInfoCriteriaNetwork")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsNetworkInfoCriteriaNetwork

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.HipObjectsNetworkInfoCriteriaNetwork
		obj, d := packHipObjectsNetworkInfoCriteriaNetworkFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.HipObjectsNetworkInfoCriteriaNetwork{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.HipObjectsNetworkInfoCriteriaNetwork", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.HipObjectsNetworkInfoCriteriaNetwork{}.AttrType(), data)
}

// --- Unpacker for HipObjectsNetworkInfoCriteriaNetworkIs ---
func unpackHipObjectsNetworkInfoCriteriaNetworkIsToSdk(ctx context.Context, obj types.Object) (*objects.HipObjectsNetworkInfoCriteriaNetworkIs, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.HipObjectsNetworkInfoCriteriaNetworkIs", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.HipObjectsNetworkInfoCriteriaNetworkIs
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.HipObjectsNetworkInfoCriteriaNetworkIs
	var d diag.Diagnostics
	// Handling Objects
	if !model.Mobile.IsNull() && !model.Mobile.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Mobile")
		unpacked, d := unpackHipObjectsNetworkInfoCriteriaNetworkIsMobileToSdk(ctx, model.Mobile)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Mobile"})
		}
		if unpacked != nil {
			sdk.Mobile = unpacked
		}
	}

	// Handling Typeless Objects
	if !model.Unknown.IsNull() && !model.Unknown.IsUnknown() {
		tflog.Debug(ctx, "Unpacking typeless object for field Unknown")
		sdk.Unknown = make(map[string]interface{})
	}

	// Handling Objects
	if !model.Wifi.IsNull() && !model.Wifi.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Wifi")
		unpacked, d := unpackHipObjectsNetworkInfoCriteriaNetworkIsWifiToSdk(ctx, model.Wifi)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Wifi"})
		}
		if unpacked != nil {
			sdk.Wifi = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.HipObjectsNetworkInfoCriteriaNetworkIs", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for HipObjectsNetworkInfoCriteriaNetworkIs ---
func packHipObjectsNetworkInfoCriteriaNetworkIsFromSdk(ctx context.Context, sdk objects.HipObjectsNetworkInfoCriteriaNetworkIs) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.HipObjectsNetworkInfoCriteriaNetworkIs", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.HipObjectsNetworkInfoCriteriaNetworkIs
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Mobile != nil {
		tflog.Debug(ctx, "Packing nested object for field Mobile")
		packed, d := packHipObjectsNetworkInfoCriteriaNetworkIsMobileFromSdk(ctx, *sdk.Mobile)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Mobile"})
		}
		model.Mobile = packed
	} else {
		model.Mobile = basetypes.NewObjectNull(models.HipObjectsNetworkInfoCriteriaNetworkIsMobile{}.AttrTypes())
	}
	// Handling Objects
	// This is a marker object (e.g. CHAP: {}). We just need to create an empty, non-null object.
	if sdk.Unknown != nil && !reflect.ValueOf(sdk.Unknown).IsNil() {
		tflog.Debug(ctx, "Packing typeless object for field Unknown")
		var d diag.Diagnostics
		// Create an empty object with no attributes, which signifies its presence.
		model.Unknown, d = basetypes.NewObjectValue(map[string]attr.Type{}, map[string]attr.Value{})
		diags.Append(d...)
	} else {
		// Since this field is part of a oneOf, being nil means it's not selected.
		// We make the object null with an empty attribute map.
		model.Unknown = basetypes.NewObjectNull(map[string]attr.Type{})
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Wifi != nil {
		tflog.Debug(ctx, "Packing nested object for field Wifi")
		packed, d := packHipObjectsNetworkInfoCriteriaNetworkIsWifiFromSdk(ctx, *sdk.Wifi)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Wifi"})
		}
		model.Wifi = packed
	} else {
		model.Wifi = basetypes.NewObjectNull(models.HipObjectsNetworkInfoCriteriaNetworkIsWifi{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.HipObjectsNetworkInfoCriteriaNetworkIs{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.HipObjectsNetworkInfoCriteriaNetworkIs", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for HipObjectsNetworkInfoCriteriaNetworkIs ---
func unpackHipObjectsNetworkInfoCriteriaNetworkIsListToSdk(ctx context.Context, list types.List) ([]objects.HipObjectsNetworkInfoCriteriaNetworkIs, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.HipObjectsNetworkInfoCriteriaNetworkIs")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsNetworkInfoCriteriaNetworkIs
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.HipObjectsNetworkInfoCriteriaNetworkIs, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.HipObjectsNetworkInfoCriteriaNetworkIs{}.AttrTypes(), &item)
		unpacked, d := unpackHipObjectsNetworkInfoCriteriaNetworkIsToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.HipObjectsNetworkInfoCriteriaNetworkIs", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for HipObjectsNetworkInfoCriteriaNetworkIs ---
func packHipObjectsNetworkInfoCriteriaNetworkIsListFromSdk(ctx context.Context, sdks []objects.HipObjectsNetworkInfoCriteriaNetworkIs) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.HipObjectsNetworkInfoCriteriaNetworkIs")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsNetworkInfoCriteriaNetworkIs

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.HipObjectsNetworkInfoCriteriaNetworkIs
		obj, d := packHipObjectsNetworkInfoCriteriaNetworkIsFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.HipObjectsNetworkInfoCriteriaNetworkIs{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.HipObjectsNetworkInfoCriteriaNetworkIs", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.HipObjectsNetworkInfoCriteriaNetworkIs{}.AttrType(), data)
}

// --- Unpacker for HipObjectsNetworkInfoCriteriaNetworkIsMobile ---
func unpackHipObjectsNetworkInfoCriteriaNetworkIsMobileToSdk(ctx context.Context, obj types.Object) (*objects.HipObjectsNetworkInfoCriteriaNetworkIsMobile, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.HipObjectsNetworkInfoCriteriaNetworkIsMobile", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.HipObjectsNetworkInfoCriteriaNetworkIsMobile
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.HipObjectsNetworkInfoCriteriaNetworkIsMobile
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Carrier.IsNull() && !model.Carrier.IsUnknown() {
		sdk.Carrier = model.Carrier.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Carrier", "value": *sdk.Carrier})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.HipObjectsNetworkInfoCriteriaNetworkIsMobile", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for HipObjectsNetworkInfoCriteriaNetworkIsMobile ---
func packHipObjectsNetworkInfoCriteriaNetworkIsMobileFromSdk(ctx context.Context, sdk objects.HipObjectsNetworkInfoCriteriaNetworkIsMobile) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.HipObjectsNetworkInfoCriteriaNetworkIsMobile", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.HipObjectsNetworkInfoCriteriaNetworkIsMobile
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Carrier != nil {
		model.Carrier = basetypes.NewStringValue(*sdk.Carrier)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Carrier", "value": *sdk.Carrier})
	} else {
		model.Carrier = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.HipObjectsNetworkInfoCriteriaNetworkIsMobile{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.HipObjectsNetworkInfoCriteriaNetworkIsMobile", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for HipObjectsNetworkInfoCriteriaNetworkIsMobile ---
func unpackHipObjectsNetworkInfoCriteriaNetworkIsMobileListToSdk(ctx context.Context, list types.List) ([]objects.HipObjectsNetworkInfoCriteriaNetworkIsMobile, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.HipObjectsNetworkInfoCriteriaNetworkIsMobile")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsNetworkInfoCriteriaNetworkIsMobile
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.HipObjectsNetworkInfoCriteriaNetworkIsMobile, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.HipObjectsNetworkInfoCriteriaNetworkIsMobile{}.AttrTypes(), &item)
		unpacked, d := unpackHipObjectsNetworkInfoCriteriaNetworkIsMobileToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.HipObjectsNetworkInfoCriteriaNetworkIsMobile", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for HipObjectsNetworkInfoCriteriaNetworkIsMobile ---
func packHipObjectsNetworkInfoCriteriaNetworkIsMobileListFromSdk(ctx context.Context, sdks []objects.HipObjectsNetworkInfoCriteriaNetworkIsMobile) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.HipObjectsNetworkInfoCriteriaNetworkIsMobile")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsNetworkInfoCriteriaNetworkIsMobile

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.HipObjectsNetworkInfoCriteriaNetworkIsMobile
		obj, d := packHipObjectsNetworkInfoCriteriaNetworkIsMobileFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.HipObjectsNetworkInfoCriteriaNetworkIsMobile{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.HipObjectsNetworkInfoCriteriaNetworkIsMobile", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.HipObjectsNetworkInfoCriteriaNetworkIsMobile{}.AttrType(), data)
}

// --- Unpacker for HipObjectsNetworkInfoCriteriaNetworkIsWifi ---
func unpackHipObjectsNetworkInfoCriteriaNetworkIsWifiToSdk(ctx context.Context, obj types.Object) (*objects.HipObjectsNetworkInfoCriteriaNetworkIsWifi, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.HipObjectsNetworkInfoCriteriaNetworkIsWifi", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.HipObjectsNetworkInfoCriteriaNetworkIsWifi
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.HipObjectsNetworkInfoCriteriaNetworkIsWifi
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Ssid.IsNull() && !model.Ssid.IsUnknown() {
		sdk.Ssid = model.Ssid.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Ssid", "value": *sdk.Ssid})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.HipObjectsNetworkInfoCriteriaNetworkIsWifi", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for HipObjectsNetworkInfoCriteriaNetworkIsWifi ---
func packHipObjectsNetworkInfoCriteriaNetworkIsWifiFromSdk(ctx context.Context, sdk objects.HipObjectsNetworkInfoCriteriaNetworkIsWifi) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.HipObjectsNetworkInfoCriteriaNetworkIsWifi", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.HipObjectsNetworkInfoCriteriaNetworkIsWifi
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Ssid != nil {
		model.Ssid = basetypes.NewStringValue(*sdk.Ssid)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Ssid", "value": *sdk.Ssid})
	} else {
		model.Ssid = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.HipObjectsNetworkInfoCriteriaNetworkIsWifi{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.HipObjectsNetworkInfoCriteriaNetworkIsWifi", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for HipObjectsNetworkInfoCriteriaNetworkIsWifi ---
func unpackHipObjectsNetworkInfoCriteriaNetworkIsWifiListToSdk(ctx context.Context, list types.List) ([]objects.HipObjectsNetworkInfoCriteriaNetworkIsWifi, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.HipObjectsNetworkInfoCriteriaNetworkIsWifi")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsNetworkInfoCriteriaNetworkIsWifi
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.HipObjectsNetworkInfoCriteriaNetworkIsWifi, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.HipObjectsNetworkInfoCriteriaNetworkIsWifi{}.AttrTypes(), &item)
		unpacked, d := unpackHipObjectsNetworkInfoCriteriaNetworkIsWifiToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.HipObjectsNetworkInfoCriteriaNetworkIsWifi", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for HipObjectsNetworkInfoCriteriaNetworkIsWifi ---
func packHipObjectsNetworkInfoCriteriaNetworkIsWifiListFromSdk(ctx context.Context, sdks []objects.HipObjectsNetworkInfoCriteriaNetworkIsWifi) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.HipObjectsNetworkInfoCriteriaNetworkIsWifi")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsNetworkInfoCriteriaNetworkIsWifi

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.HipObjectsNetworkInfoCriteriaNetworkIsWifi
		obj, d := packHipObjectsNetworkInfoCriteriaNetworkIsWifiFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.HipObjectsNetworkInfoCriteriaNetworkIsWifi{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.HipObjectsNetworkInfoCriteriaNetworkIsWifi", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.HipObjectsNetworkInfoCriteriaNetworkIsWifi{}.AttrType(), data)
}

// --- Unpacker for HipObjectsNetworkInfoCriteriaNetworkIsNot ---
func unpackHipObjectsNetworkInfoCriteriaNetworkIsNotToSdk(ctx context.Context, obj types.Object) (*objects.HipObjectsNetworkInfoCriteriaNetworkIsNot, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.HipObjectsNetworkInfoCriteriaNetworkIsNot", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.HipObjectsNetworkInfoCriteriaNetworkIsNot
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.HipObjectsNetworkInfoCriteriaNetworkIsNot
	var d diag.Diagnostics
	// Handling Typeless Objects
	if !model.Ethernet.IsNull() && !model.Ethernet.IsUnknown() {
		tflog.Debug(ctx, "Unpacking typeless object for field Ethernet")
		sdk.Ethernet = make(map[string]interface{})
	}

	// Handling Objects
	if !model.Mobile.IsNull() && !model.Mobile.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Mobile")
		unpacked, d := unpackHipObjectsNetworkInfoCriteriaNetworkIsMobileToSdk(ctx, model.Mobile)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Mobile"})
		}
		if unpacked != nil {
			sdk.Mobile = unpacked
		}
	}

	// Handling Typeless Objects
	if !model.Unknown.IsNull() && !model.Unknown.IsUnknown() {
		tflog.Debug(ctx, "Unpacking typeless object for field Unknown")
		sdk.Unknown = make(map[string]interface{})
	}

	// Handling Objects
	if !model.Wifi.IsNull() && !model.Wifi.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Wifi")
		unpacked, d := unpackHipObjectsNetworkInfoCriteriaNetworkIsWifiToSdk(ctx, model.Wifi)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Wifi"})
		}
		if unpacked != nil {
			sdk.Wifi = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.HipObjectsNetworkInfoCriteriaNetworkIsNot", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for HipObjectsNetworkInfoCriteriaNetworkIsNot ---
func packHipObjectsNetworkInfoCriteriaNetworkIsNotFromSdk(ctx context.Context, sdk objects.HipObjectsNetworkInfoCriteriaNetworkIsNot) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.HipObjectsNetworkInfoCriteriaNetworkIsNot", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.HipObjectsNetworkInfoCriteriaNetworkIsNot
	var d diag.Diagnostics
	// Handling Objects
	// This is a marker object (e.g. CHAP: {}). We just need to create an empty, non-null object.
	if sdk.Ethernet != nil && !reflect.ValueOf(sdk.Ethernet).IsNil() {
		tflog.Debug(ctx, "Packing typeless object for field Ethernet")
		var d diag.Diagnostics
		// Create an empty object with no attributes, which signifies its presence.
		model.Ethernet, d = basetypes.NewObjectValue(map[string]attr.Type{}, map[string]attr.Value{})
		diags.Append(d...)
	} else {
		// Since this field is part of a oneOf, being nil means it's not selected.
		// We make the object null with an empty attribute map.
		model.Ethernet = basetypes.NewObjectNull(map[string]attr.Type{})
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Mobile != nil {
		tflog.Debug(ctx, "Packing nested object for field Mobile")
		packed, d := packHipObjectsNetworkInfoCriteriaNetworkIsMobileFromSdk(ctx, *sdk.Mobile)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Mobile"})
		}
		model.Mobile = packed
	} else {
		model.Mobile = basetypes.NewObjectNull(models.HipObjectsNetworkInfoCriteriaNetworkIsMobile{}.AttrTypes())
	}
	// Handling Objects
	// This is a marker object (e.g. CHAP: {}). We just need to create an empty, non-null object.
	if sdk.Unknown != nil && !reflect.ValueOf(sdk.Unknown).IsNil() {
		tflog.Debug(ctx, "Packing typeless object for field Unknown")
		var d diag.Diagnostics
		// Create an empty object with no attributes, which signifies its presence.
		model.Unknown, d = basetypes.NewObjectValue(map[string]attr.Type{}, map[string]attr.Value{})
		diags.Append(d...)
	} else {
		// Since this field is part of a oneOf, being nil means it's not selected.
		// We make the object null with an empty attribute map.
		model.Unknown = basetypes.NewObjectNull(map[string]attr.Type{})
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Wifi != nil {
		tflog.Debug(ctx, "Packing nested object for field Wifi")
		packed, d := packHipObjectsNetworkInfoCriteriaNetworkIsWifiFromSdk(ctx, *sdk.Wifi)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Wifi"})
		}
		model.Wifi = packed
	} else {
		model.Wifi = basetypes.NewObjectNull(models.HipObjectsNetworkInfoCriteriaNetworkIsWifi{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.HipObjectsNetworkInfoCriteriaNetworkIsNot{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.HipObjectsNetworkInfoCriteriaNetworkIsNot", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for HipObjectsNetworkInfoCriteriaNetworkIsNot ---
func unpackHipObjectsNetworkInfoCriteriaNetworkIsNotListToSdk(ctx context.Context, list types.List) ([]objects.HipObjectsNetworkInfoCriteriaNetworkIsNot, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.HipObjectsNetworkInfoCriteriaNetworkIsNot")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsNetworkInfoCriteriaNetworkIsNot
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.HipObjectsNetworkInfoCriteriaNetworkIsNot, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.HipObjectsNetworkInfoCriteriaNetworkIsNot{}.AttrTypes(), &item)
		unpacked, d := unpackHipObjectsNetworkInfoCriteriaNetworkIsNotToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.HipObjectsNetworkInfoCriteriaNetworkIsNot", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for HipObjectsNetworkInfoCriteriaNetworkIsNot ---
func packHipObjectsNetworkInfoCriteriaNetworkIsNotListFromSdk(ctx context.Context, sdks []objects.HipObjectsNetworkInfoCriteriaNetworkIsNot) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.HipObjectsNetworkInfoCriteriaNetworkIsNot")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsNetworkInfoCriteriaNetworkIsNot

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.HipObjectsNetworkInfoCriteriaNetworkIsNot
		obj, d := packHipObjectsNetworkInfoCriteriaNetworkIsNotFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.HipObjectsNetworkInfoCriteriaNetworkIsNot{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.HipObjectsNetworkInfoCriteriaNetworkIsNot", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.HipObjectsNetworkInfoCriteriaNetworkIsNot{}.AttrType(), data)
}

// --- Unpacker for HipObjectsPatchManagement ---
func unpackHipObjectsPatchManagementToSdk(ctx context.Context, obj types.Object) (*objects.HipObjectsPatchManagement, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.HipObjectsPatchManagement", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.HipObjectsPatchManagement
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.HipObjectsPatchManagement
	var d diag.Diagnostics
	// Handling Objects
	if !model.Criteria.IsNull() && !model.Criteria.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Criteria")
		unpacked, d := unpackHipObjectsPatchManagementCriteriaToSdk(ctx, model.Criteria)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Criteria"})
		}
		if unpacked != nil {
			sdk.Criteria = unpacked
		}
	}

	// Handling Primitives
	if !model.ExcludeVendor.IsNull() && !model.ExcludeVendor.IsUnknown() {
		sdk.ExcludeVendor = model.ExcludeVendor.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "ExcludeVendor", "value": *sdk.ExcludeVendor})
	}

	// Handling Lists
	if !model.Vendor.IsNull() && !model.Vendor.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field Vendor")
		unpacked, d := unpackHipObjectsDataLossPreventionVendorInnerListToSdk(ctx, model.Vendor)
		diags.Append(d...)
		sdk.Vendor = unpacked
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.HipObjectsPatchManagement", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for HipObjectsPatchManagement ---
func packHipObjectsPatchManagementFromSdk(ctx context.Context, sdk objects.HipObjectsPatchManagement) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.HipObjectsPatchManagement", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.HipObjectsPatchManagement
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Criteria != nil {
		tflog.Debug(ctx, "Packing nested object for field Criteria")
		packed, d := packHipObjectsPatchManagementCriteriaFromSdk(ctx, *sdk.Criteria)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Criteria"})
		}
		model.Criteria = packed
	} else {
		model.Criteria = basetypes.NewObjectNull(models.HipObjectsPatchManagementCriteria{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.ExcludeVendor != nil {
		model.ExcludeVendor = basetypes.NewBoolValue(*sdk.ExcludeVendor)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "ExcludeVendor", "value": *sdk.ExcludeVendor})
	} else {
		model.ExcludeVendor = basetypes.NewBoolNull()
	}
	// Handling Lists
	if sdk.Vendor != nil {
		tflog.Debug(ctx, "Packing list of objects for field Vendor")
		packed, d := packHipObjectsDataLossPreventionVendorInnerListFromSdk(ctx, sdk.Vendor)
		diags.Append(d...)
		model.Vendor = packed
	} else {
		model.Vendor = basetypes.NewListNull(models.HipObjectsDataLossPreventionVendorInner{}.AttrType())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.HipObjectsPatchManagement{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.HipObjectsPatchManagement", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for HipObjectsPatchManagement ---
func unpackHipObjectsPatchManagementListToSdk(ctx context.Context, list types.List) ([]objects.HipObjectsPatchManagement, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.HipObjectsPatchManagement")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsPatchManagement
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.HipObjectsPatchManagement, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.HipObjectsPatchManagement{}.AttrTypes(), &item)
		unpacked, d := unpackHipObjectsPatchManagementToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.HipObjectsPatchManagement", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for HipObjectsPatchManagement ---
func packHipObjectsPatchManagementListFromSdk(ctx context.Context, sdks []objects.HipObjectsPatchManagement) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.HipObjectsPatchManagement")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsPatchManagement

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.HipObjectsPatchManagement
		obj, d := packHipObjectsPatchManagementFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.HipObjectsPatchManagement{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.HipObjectsPatchManagement", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.HipObjectsPatchManagement{}.AttrType(), data)
}

// --- Unpacker for HipObjectsPatchManagementCriteria ---
func unpackHipObjectsPatchManagementCriteriaToSdk(ctx context.Context, obj types.Object) (*objects.HipObjectsPatchManagementCriteria, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.HipObjectsPatchManagementCriteria", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.HipObjectsPatchManagementCriteria
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.HipObjectsPatchManagementCriteria
	var d diag.Diagnostics
	// Handling Primitives
	if !model.IsEnabled.IsNull() && !model.IsEnabled.IsUnknown() {
		sdk.IsEnabled = model.IsEnabled.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "IsEnabled", "value": *sdk.IsEnabled})
	}

	// Handling Primitives
	if !model.IsInstalled.IsNull() && !model.IsInstalled.IsUnknown() {
		sdk.IsInstalled = model.IsInstalled.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "IsInstalled", "value": *sdk.IsInstalled})
	}

	// Handling Objects
	if !model.MissingPatches.IsNull() && !model.MissingPatches.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field MissingPatches")
		unpacked, d := unpackHipObjectsPatchManagementCriteriaMissingPatchesToSdk(ctx, model.MissingPatches)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "MissingPatches"})
		}
		if unpacked != nil {
			sdk.MissingPatches = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.HipObjectsPatchManagementCriteria", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for HipObjectsPatchManagementCriteria ---
func packHipObjectsPatchManagementCriteriaFromSdk(ctx context.Context, sdk objects.HipObjectsPatchManagementCriteria) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.HipObjectsPatchManagementCriteria", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.HipObjectsPatchManagementCriteria
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.IsEnabled != nil {
		model.IsEnabled = basetypes.NewStringValue(*sdk.IsEnabled)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "IsEnabled", "value": *sdk.IsEnabled})
	} else {
		model.IsEnabled = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.IsInstalled != nil {
		model.IsInstalled = basetypes.NewBoolValue(*sdk.IsInstalled)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "IsInstalled", "value": *sdk.IsInstalled})
	} else {
		model.IsInstalled = basetypes.NewBoolNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.MissingPatches != nil {
		tflog.Debug(ctx, "Packing nested object for field MissingPatches")
		packed, d := packHipObjectsPatchManagementCriteriaMissingPatchesFromSdk(ctx, *sdk.MissingPatches)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "MissingPatches"})
		}
		model.MissingPatches = packed
	} else {
		model.MissingPatches = basetypes.NewObjectNull(models.HipObjectsPatchManagementCriteriaMissingPatches{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.HipObjectsPatchManagementCriteria{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.HipObjectsPatchManagementCriteria", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for HipObjectsPatchManagementCriteria ---
func unpackHipObjectsPatchManagementCriteriaListToSdk(ctx context.Context, list types.List) ([]objects.HipObjectsPatchManagementCriteria, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.HipObjectsPatchManagementCriteria")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsPatchManagementCriteria
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.HipObjectsPatchManagementCriteria, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.HipObjectsPatchManagementCriteria{}.AttrTypes(), &item)
		unpacked, d := unpackHipObjectsPatchManagementCriteriaToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.HipObjectsPatchManagementCriteria", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for HipObjectsPatchManagementCriteria ---
func packHipObjectsPatchManagementCriteriaListFromSdk(ctx context.Context, sdks []objects.HipObjectsPatchManagementCriteria) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.HipObjectsPatchManagementCriteria")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsPatchManagementCriteria

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.HipObjectsPatchManagementCriteria
		obj, d := packHipObjectsPatchManagementCriteriaFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.HipObjectsPatchManagementCriteria{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.HipObjectsPatchManagementCriteria", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.HipObjectsPatchManagementCriteria{}.AttrType(), data)
}

// --- Unpacker for HipObjectsPatchManagementCriteriaMissingPatches ---
func unpackHipObjectsPatchManagementCriteriaMissingPatchesToSdk(ctx context.Context, obj types.Object) (*objects.HipObjectsPatchManagementCriteriaMissingPatches, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.HipObjectsPatchManagementCriteriaMissingPatches", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.HipObjectsPatchManagementCriteriaMissingPatches
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.HipObjectsPatchManagementCriteriaMissingPatches
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Check.IsNull() && !model.Check.IsUnknown() {
		sdk.Check = model.Check.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Check", "value": sdk.Check})
	}

	// Handling Lists
	if !model.Patches.IsNull() && !model.Patches.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Patches")
		diags.Append(model.Patches.ElementsAs(ctx, &sdk.Patches, false)...)
	}

	// Handling Objects
	if !model.Severity.IsNull() && !model.Severity.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Severity")
		unpacked, d := unpackHipObjectsPatchManagementCriteriaMissingPatchesSeverityToSdk(ctx, model.Severity)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Severity"})
		}
		if unpacked != nil {
			sdk.Severity = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.HipObjectsPatchManagementCriteriaMissingPatches", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for HipObjectsPatchManagementCriteriaMissingPatches ---
func packHipObjectsPatchManagementCriteriaMissingPatchesFromSdk(ctx context.Context, sdk objects.HipObjectsPatchManagementCriteriaMissingPatches) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.HipObjectsPatchManagementCriteriaMissingPatches", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.HipObjectsPatchManagementCriteriaMissingPatches
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.Check = basetypes.NewStringValue(sdk.Check)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Check", "value": sdk.Check})
	// Handling Lists
	if sdk.Patches != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Patches")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Patches, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Patches)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Patches = basetypes.NewListNull(elemType)
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Severity != nil {
		tflog.Debug(ctx, "Packing nested object for field Severity")
		packed, d := packHipObjectsPatchManagementCriteriaMissingPatchesSeverityFromSdk(ctx, *sdk.Severity)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Severity"})
		}
		model.Severity = packed
	} else {
		model.Severity = basetypes.NewObjectNull(models.HipObjectsPatchManagementCriteriaMissingPatchesSeverity{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.HipObjectsPatchManagementCriteriaMissingPatches{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.HipObjectsPatchManagementCriteriaMissingPatches", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for HipObjectsPatchManagementCriteriaMissingPatches ---
func unpackHipObjectsPatchManagementCriteriaMissingPatchesListToSdk(ctx context.Context, list types.List) ([]objects.HipObjectsPatchManagementCriteriaMissingPatches, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.HipObjectsPatchManagementCriteriaMissingPatches")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsPatchManagementCriteriaMissingPatches
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.HipObjectsPatchManagementCriteriaMissingPatches, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.HipObjectsPatchManagementCriteriaMissingPatches{}.AttrTypes(), &item)
		unpacked, d := unpackHipObjectsPatchManagementCriteriaMissingPatchesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.HipObjectsPatchManagementCriteriaMissingPatches", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for HipObjectsPatchManagementCriteriaMissingPatches ---
func packHipObjectsPatchManagementCriteriaMissingPatchesListFromSdk(ctx context.Context, sdks []objects.HipObjectsPatchManagementCriteriaMissingPatches) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.HipObjectsPatchManagementCriteriaMissingPatches")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsPatchManagementCriteriaMissingPatches

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.HipObjectsPatchManagementCriteriaMissingPatches
		obj, d := packHipObjectsPatchManagementCriteriaMissingPatchesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.HipObjectsPatchManagementCriteriaMissingPatches{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.HipObjectsPatchManagementCriteriaMissingPatches", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.HipObjectsPatchManagementCriteriaMissingPatches{}.AttrType(), data)
}

// --- Unpacker for HipObjectsPatchManagementCriteriaMissingPatchesSeverity ---
func unpackHipObjectsPatchManagementCriteriaMissingPatchesSeverityToSdk(ctx context.Context, obj types.Object) (*objects.HipObjectsPatchManagementCriteriaMissingPatchesSeverity, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.HipObjectsPatchManagementCriteriaMissingPatchesSeverity", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.HipObjectsPatchManagementCriteriaMissingPatchesSeverity
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.HipObjectsPatchManagementCriteriaMissingPatchesSeverity
	var d diag.Diagnostics
	// Handling Primitives
	if !model.GreaterEqual.IsNull() && !model.GreaterEqual.IsUnknown() {
		val := int32(model.GreaterEqual.ValueInt64())
		sdk.GreaterEqual = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "GreaterEqual", "value": *sdk.GreaterEqual})
	}

	// Handling Primitives
	if !model.GreaterThan.IsNull() && !model.GreaterThan.IsUnknown() {
		val := int32(model.GreaterThan.ValueInt64())
		sdk.GreaterThan = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "GreaterThan", "value": *sdk.GreaterThan})
	}

	// Handling Primitives
	if !model.Is.IsNull() && !model.Is.IsUnknown() {
		val := int32(model.Is.ValueInt64())
		sdk.Is = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Is", "value": *sdk.Is})
	}

	// Handling Primitives
	if !model.IsNot.IsNull() && !model.IsNot.IsUnknown() {
		val := int32(model.IsNot.ValueInt64())
		sdk.IsNot = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "IsNot", "value": *sdk.IsNot})
	}

	// Handling Primitives
	if !model.LessEqual.IsNull() && !model.LessEqual.IsUnknown() {
		val := int32(model.LessEqual.ValueInt64())
		sdk.LessEqual = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "LessEqual", "value": *sdk.LessEqual})
	}

	// Handling Primitives
	if !model.LessThan.IsNull() && !model.LessThan.IsUnknown() {
		val := int32(model.LessThan.ValueInt64())
		sdk.LessThan = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "LessThan", "value": *sdk.LessThan})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.HipObjectsPatchManagementCriteriaMissingPatchesSeverity", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for HipObjectsPatchManagementCriteriaMissingPatchesSeverity ---
func packHipObjectsPatchManagementCriteriaMissingPatchesSeverityFromSdk(ctx context.Context, sdk objects.HipObjectsPatchManagementCriteriaMissingPatchesSeverity) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.HipObjectsPatchManagementCriteriaMissingPatchesSeverity", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.HipObjectsPatchManagementCriteriaMissingPatchesSeverity
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.GreaterEqual != nil {
		model.GreaterEqual = basetypes.NewInt64Value(int64(*sdk.GreaterEqual))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "GreaterEqual", "value": *sdk.GreaterEqual})
	} else {
		model.GreaterEqual = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.GreaterThan != nil {
		model.GreaterThan = basetypes.NewInt64Value(int64(*sdk.GreaterThan))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "GreaterThan", "value": *sdk.GreaterThan})
	} else {
		model.GreaterThan = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Is != nil {
		model.Is = basetypes.NewInt64Value(int64(*sdk.Is))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Is", "value": *sdk.Is})
	} else {
		model.Is = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.IsNot != nil {
		model.IsNot = basetypes.NewInt64Value(int64(*sdk.IsNot))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "IsNot", "value": *sdk.IsNot})
	} else {
		model.IsNot = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.LessEqual != nil {
		model.LessEqual = basetypes.NewInt64Value(int64(*sdk.LessEqual))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "LessEqual", "value": *sdk.LessEqual})
	} else {
		model.LessEqual = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.LessThan != nil {
		model.LessThan = basetypes.NewInt64Value(int64(*sdk.LessThan))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "LessThan", "value": *sdk.LessThan})
	} else {
		model.LessThan = basetypes.NewInt64Null()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.HipObjectsPatchManagementCriteriaMissingPatchesSeverity{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.HipObjectsPatchManagementCriteriaMissingPatchesSeverity", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for HipObjectsPatchManagementCriteriaMissingPatchesSeverity ---
func unpackHipObjectsPatchManagementCriteriaMissingPatchesSeverityListToSdk(ctx context.Context, list types.List) ([]objects.HipObjectsPatchManagementCriteriaMissingPatchesSeverity, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.HipObjectsPatchManagementCriteriaMissingPatchesSeverity")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsPatchManagementCriteriaMissingPatchesSeverity
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.HipObjectsPatchManagementCriteriaMissingPatchesSeverity, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.HipObjectsPatchManagementCriteriaMissingPatchesSeverity{}.AttrTypes(), &item)
		unpacked, d := unpackHipObjectsPatchManagementCriteriaMissingPatchesSeverityToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.HipObjectsPatchManagementCriteriaMissingPatchesSeverity", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for HipObjectsPatchManagementCriteriaMissingPatchesSeverity ---
func packHipObjectsPatchManagementCriteriaMissingPatchesSeverityListFromSdk(ctx context.Context, sdks []objects.HipObjectsPatchManagementCriteriaMissingPatchesSeverity) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.HipObjectsPatchManagementCriteriaMissingPatchesSeverity")
	diags := diag.Diagnostics{}
	var data []models.HipObjectsPatchManagementCriteriaMissingPatchesSeverity

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.HipObjectsPatchManagementCriteriaMissingPatchesSeverity
		obj, d := packHipObjectsPatchManagementCriteriaMissingPatchesSeverityFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.HipObjectsPatchManagementCriteriaMissingPatchesSeverity{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.HipObjectsPatchManagementCriteriaMissingPatchesSeverity", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.HipObjectsPatchManagementCriteriaMissingPatchesSeverity{}.AttrType(), data)
}
