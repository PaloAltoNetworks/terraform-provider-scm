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

// --- Unpacker for QosProfiles ---
func unpackQosProfilesToSdk(ctx context.Context, obj types.Object) (*network_services.QosProfiles, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.QosProfiles", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.QosProfiles
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.QosProfiles
	var d diag.Diagnostics

	// Handling Objects
	if !model.AggregateBandwidth.IsNull() && !model.AggregateBandwidth.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field AggregateBandwidth")
		unpacked, d := unpackQosProfilesAggregateBandwidthToSdk(ctx, model.AggregateBandwidth)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "AggregateBandwidth"})
		}
		if unpacked != nil {
			sdk.AggregateBandwidth = unpacked
		}
	}

	// Handling Objects
	if !model.ClassBandwidthType.IsNull() && !model.ClassBandwidthType.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field ClassBandwidthType")
		unpacked, d := unpackQosProfilesClassBandwidthTypeToSdk(ctx, model.ClassBandwidthType)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "ClassBandwidthType"})
		}
		if unpacked != nil {
			sdk.ClassBandwidthType = unpacked
		}
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

	// Handling Primitives
	if !model.Snippet.IsNull() && !model.Snippet.IsUnknown() {
		sdk.Snippet = model.Snippet.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.QosProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for QosProfiles ---
func packQosProfilesFromSdk(ctx context.Context, sdk network_services.QosProfiles) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.QosProfiles", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.QosProfiles
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.AggregateBandwidth != nil {
		tflog.Debug(ctx, "Packing nested object for field AggregateBandwidth")
		packed, d := packQosProfilesAggregateBandwidthFromSdk(ctx, *sdk.AggregateBandwidth)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "AggregateBandwidth"})
		}
		model.AggregateBandwidth = packed
	} else {
		model.AggregateBandwidth = basetypes.NewObjectNull(models.QosProfilesAggregateBandwidth{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.ClassBandwidthType != nil {
		tflog.Debug(ctx, "Packing nested object for field ClassBandwidthType")
		packed, d := packQosProfilesClassBandwidthTypeFromSdk(ctx, *sdk.ClassBandwidthType)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "ClassBandwidthType"})
		}
		model.ClassBandwidthType = packed
	} else {
		model.ClassBandwidthType = basetypes.NewObjectNull(models.QosProfilesClassBandwidthType{}.AttrTypes())
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
	// Handling Primitives
	// Standard primitive packing
	if sdk.Snippet != nil {
		model.Snippet = basetypes.NewStringValue(*sdk.Snippet)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	} else {
		model.Snippet = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.QosProfiles{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.QosProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for QosProfiles ---
func unpackQosProfilesListToSdk(ctx context.Context, list types.List) ([]network_services.QosProfiles, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.QosProfiles")
	diags := diag.Diagnostics{}
	var data []models.QosProfiles
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.QosProfiles, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.QosProfiles{}.AttrTypes(), &item)
		unpacked, d := unpackQosProfilesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.QosProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for QosProfiles ---
func packQosProfilesListFromSdk(ctx context.Context, sdks []network_services.QosProfiles) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.QosProfiles")
	diags := diag.Diagnostics{}
	var data []models.QosProfiles

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.QosProfiles
		obj, d := packQosProfilesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.QosProfiles{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.QosProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.QosProfiles{}.AttrType(), data)
}

// --- Unpacker for QosProfilesAggregateBandwidth ---
func unpackQosProfilesAggregateBandwidthToSdk(ctx context.Context, obj types.Object) (*network_services.QosProfilesAggregateBandwidth, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.QosProfilesAggregateBandwidth", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.QosProfilesAggregateBandwidth
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.QosProfilesAggregateBandwidth
	var d diag.Diagnostics
	// Handling Primitives
	if !model.EgressGuaranteed.IsNull() && !model.EgressGuaranteed.IsUnknown() {
		val := int32(model.EgressGuaranteed.ValueInt64())
		sdk.EgressGuaranteed = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "EgressGuaranteed", "value": *sdk.EgressGuaranteed})
	}

	// Handling Primitives
	if !model.EgressMax.IsNull() && !model.EgressMax.IsUnknown() {
		val := int32(model.EgressMax.ValueInt64())
		sdk.EgressMax = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "EgressMax", "value": *sdk.EgressMax})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.QosProfilesAggregateBandwidth", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for QosProfilesAggregateBandwidth ---
func packQosProfilesAggregateBandwidthFromSdk(ctx context.Context, sdk network_services.QosProfilesAggregateBandwidth) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.QosProfilesAggregateBandwidth", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.QosProfilesAggregateBandwidth
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.EgressGuaranteed != nil {
		model.EgressGuaranteed = basetypes.NewInt64Value(int64(*sdk.EgressGuaranteed))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "EgressGuaranteed", "value": *sdk.EgressGuaranteed})
	} else {
		model.EgressGuaranteed = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.EgressMax != nil {
		model.EgressMax = basetypes.NewInt64Value(int64(*sdk.EgressMax))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "EgressMax", "value": *sdk.EgressMax})
	} else {
		model.EgressMax = basetypes.NewInt64Null()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.QosProfilesAggregateBandwidth{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.QosProfilesAggregateBandwidth", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for QosProfilesAggregateBandwidth ---
func unpackQosProfilesAggregateBandwidthListToSdk(ctx context.Context, list types.List) ([]network_services.QosProfilesAggregateBandwidth, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.QosProfilesAggregateBandwidth")
	diags := diag.Diagnostics{}
	var data []models.QosProfilesAggregateBandwidth
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.QosProfilesAggregateBandwidth, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.QosProfilesAggregateBandwidth{}.AttrTypes(), &item)
		unpacked, d := unpackQosProfilesAggregateBandwidthToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.QosProfilesAggregateBandwidth", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for QosProfilesAggregateBandwidth ---
func packQosProfilesAggregateBandwidthListFromSdk(ctx context.Context, sdks []network_services.QosProfilesAggregateBandwidth) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.QosProfilesAggregateBandwidth")
	diags := diag.Diagnostics{}
	var data []models.QosProfilesAggregateBandwidth

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.QosProfilesAggregateBandwidth
		obj, d := packQosProfilesAggregateBandwidthFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.QosProfilesAggregateBandwidth{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.QosProfilesAggregateBandwidth", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.QosProfilesAggregateBandwidth{}.AttrType(), data)
}

// --- Unpacker for QosProfilesClassBandwidthType ---
func unpackQosProfilesClassBandwidthTypeToSdk(ctx context.Context, obj types.Object) (*network_services.QosProfilesClassBandwidthType, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.QosProfilesClassBandwidthType", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.QosProfilesClassBandwidthType
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.QosProfilesClassBandwidthType
	var d diag.Diagnostics
	// Handling Objects
	if !model.Mbps.IsNull() && !model.Mbps.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Mbps")
		unpacked, d := unpackQosProfilesClassBandwidthTypeMbpsToSdk(ctx, model.Mbps)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Mbps"})
		}
		if unpacked != nil {
			sdk.Mbps = unpacked
		}
	}

	// Handling Objects
	if !model.Percentage.IsNull() && !model.Percentage.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Percentage")
		unpacked, d := unpackQosProfilesClassBandwidthTypePercentageToSdk(ctx, model.Percentage)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Percentage"})
		}
		if unpacked != nil {
			sdk.Percentage = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.QosProfilesClassBandwidthType", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for QosProfilesClassBandwidthType ---
func packQosProfilesClassBandwidthTypeFromSdk(ctx context.Context, sdk network_services.QosProfilesClassBandwidthType) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.QosProfilesClassBandwidthType", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.QosProfilesClassBandwidthType
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Mbps != nil {
		tflog.Debug(ctx, "Packing nested object for field Mbps")
		packed, d := packQosProfilesClassBandwidthTypeMbpsFromSdk(ctx, *sdk.Mbps)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Mbps"})
		}
		model.Mbps = packed
	} else {
		model.Mbps = basetypes.NewObjectNull(models.QosProfilesClassBandwidthTypeMbps{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Percentage != nil {
		tflog.Debug(ctx, "Packing nested object for field Percentage")
		packed, d := packQosProfilesClassBandwidthTypePercentageFromSdk(ctx, *sdk.Percentage)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Percentage"})
		}
		model.Percentage = packed
	} else {
		model.Percentage = basetypes.NewObjectNull(models.QosProfilesClassBandwidthTypePercentage{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.QosProfilesClassBandwidthType{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.QosProfilesClassBandwidthType", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for QosProfilesClassBandwidthType ---
func unpackQosProfilesClassBandwidthTypeListToSdk(ctx context.Context, list types.List) ([]network_services.QosProfilesClassBandwidthType, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.QosProfilesClassBandwidthType")
	diags := diag.Diagnostics{}
	var data []models.QosProfilesClassBandwidthType
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.QosProfilesClassBandwidthType, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.QosProfilesClassBandwidthType{}.AttrTypes(), &item)
		unpacked, d := unpackQosProfilesClassBandwidthTypeToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.QosProfilesClassBandwidthType", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for QosProfilesClassBandwidthType ---
func packQosProfilesClassBandwidthTypeListFromSdk(ctx context.Context, sdks []network_services.QosProfilesClassBandwidthType) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.QosProfilesClassBandwidthType")
	diags := diag.Diagnostics{}
	var data []models.QosProfilesClassBandwidthType

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.QosProfilesClassBandwidthType
		obj, d := packQosProfilesClassBandwidthTypeFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.QosProfilesClassBandwidthType{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.QosProfilesClassBandwidthType", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.QosProfilesClassBandwidthType{}.AttrType(), data)
}

// --- Unpacker for QosProfilesClassBandwidthTypeMbps ---
func unpackQosProfilesClassBandwidthTypeMbpsToSdk(ctx context.Context, obj types.Object) (*network_services.QosProfilesClassBandwidthTypeMbps, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.QosProfilesClassBandwidthTypeMbps", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.QosProfilesClassBandwidthTypeMbps
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.QosProfilesClassBandwidthTypeMbps
	var d diag.Diagnostics
	// Handling Lists
	if !model.Class.IsNull() && !model.Class.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field Class")
		unpacked, d := unpackQosProfilesClassBandwidthTypeMbpsClassInnerListToSdk(ctx, model.Class)
		diags.Append(d...)
		sdk.Class = unpacked
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.QosProfilesClassBandwidthTypeMbps", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for QosProfilesClassBandwidthTypeMbps ---
func packQosProfilesClassBandwidthTypeMbpsFromSdk(ctx context.Context, sdk network_services.QosProfilesClassBandwidthTypeMbps) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.QosProfilesClassBandwidthTypeMbps", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.QosProfilesClassBandwidthTypeMbps
	var d diag.Diagnostics
	// Handling Lists
	if sdk.Class != nil {
		tflog.Debug(ctx, "Packing list of objects for field Class")
		packed, d := packQosProfilesClassBandwidthTypeMbpsClassInnerListFromSdk(ctx, sdk.Class)
		diags.Append(d...)
		model.Class = packed
	} else {
		model.Class = basetypes.NewListNull(models.QosProfilesClassBandwidthTypeMbpsClassInner{}.AttrType())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.QosProfilesClassBandwidthTypeMbps{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.QosProfilesClassBandwidthTypeMbps", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for QosProfilesClassBandwidthTypeMbps ---
func unpackQosProfilesClassBandwidthTypeMbpsListToSdk(ctx context.Context, list types.List) ([]network_services.QosProfilesClassBandwidthTypeMbps, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.QosProfilesClassBandwidthTypeMbps")
	diags := diag.Diagnostics{}
	var data []models.QosProfilesClassBandwidthTypeMbps
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.QosProfilesClassBandwidthTypeMbps, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.QosProfilesClassBandwidthTypeMbps{}.AttrTypes(), &item)
		unpacked, d := unpackQosProfilesClassBandwidthTypeMbpsToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.QosProfilesClassBandwidthTypeMbps", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for QosProfilesClassBandwidthTypeMbps ---
func packQosProfilesClassBandwidthTypeMbpsListFromSdk(ctx context.Context, sdks []network_services.QosProfilesClassBandwidthTypeMbps) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.QosProfilesClassBandwidthTypeMbps")
	diags := diag.Diagnostics{}
	var data []models.QosProfilesClassBandwidthTypeMbps

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.QosProfilesClassBandwidthTypeMbps
		obj, d := packQosProfilesClassBandwidthTypeMbpsFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.QosProfilesClassBandwidthTypeMbps{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.QosProfilesClassBandwidthTypeMbps", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.QosProfilesClassBandwidthTypeMbps{}.AttrType(), data)
}

// --- Unpacker for QosProfilesClassBandwidthTypeMbpsClassInner ---
func unpackQosProfilesClassBandwidthTypeMbpsClassInnerToSdk(ctx context.Context, obj types.Object) (*network_services.QosProfilesClassBandwidthTypeMbpsClassInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.QosProfilesClassBandwidthTypeMbpsClassInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.QosProfilesClassBandwidthTypeMbpsClassInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.QosProfilesClassBandwidthTypeMbpsClassInner
	var d diag.Diagnostics
	// Handling Objects
	if !model.ClassBandwidth.IsNull() && !model.ClassBandwidth.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field ClassBandwidth")
		unpacked, d := unpackQosProfilesClassBandwidthTypeMbpsClassInnerClassBandwidthToSdk(ctx, model.ClassBandwidth)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "ClassBandwidth"})
		}
		if unpacked != nil {
			sdk.ClassBandwidth = unpacked
		}
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
	}

	// Handling Primitives
	if !model.Priority.IsNull() && !model.Priority.IsUnknown() {
		sdk.Priority = model.Priority.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Priority", "value": *sdk.Priority})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.QosProfilesClassBandwidthTypeMbpsClassInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for QosProfilesClassBandwidthTypeMbpsClassInner ---
func packQosProfilesClassBandwidthTypeMbpsClassInnerFromSdk(ctx context.Context, sdk network_services.QosProfilesClassBandwidthTypeMbpsClassInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.QosProfilesClassBandwidthTypeMbpsClassInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.QosProfilesClassBandwidthTypeMbpsClassInner
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.ClassBandwidth != nil {
		tflog.Debug(ctx, "Packing nested object for field ClassBandwidth")
		packed, d := packQosProfilesClassBandwidthTypeMbpsClassInnerClassBandwidthFromSdk(ctx, *sdk.ClassBandwidth)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "ClassBandwidth"})
		}
		model.ClassBandwidth = packed
	} else {
		model.ClassBandwidth = basetypes.NewObjectNull(models.QosProfilesClassBandwidthTypeMbpsClassInnerClassBandwidth{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Name != nil {
		model.Name = basetypes.NewStringValue(*sdk.Name)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
	} else {
		model.Name = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Priority != nil {
		model.Priority = basetypes.NewStringValue(*sdk.Priority)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Priority", "value": *sdk.Priority})
	} else {
		model.Priority = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.QosProfilesClassBandwidthTypeMbpsClassInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.QosProfilesClassBandwidthTypeMbpsClassInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for QosProfilesClassBandwidthTypeMbpsClassInner ---
func unpackQosProfilesClassBandwidthTypeMbpsClassInnerListToSdk(ctx context.Context, list types.List) ([]network_services.QosProfilesClassBandwidthTypeMbpsClassInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.QosProfilesClassBandwidthTypeMbpsClassInner")
	diags := diag.Diagnostics{}
	var data []models.QosProfilesClassBandwidthTypeMbpsClassInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.QosProfilesClassBandwidthTypeMbpsClassInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.QosProfilesClassBandwidthTypeMbpsClassInner{}.AttrTypes(), &item)
		unpacked, d := unpackQosProfilesClassBandwidthTypeMbpsClassInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.QosProfilesClassBandwidthTypeMbpsClassInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for QosProfilesClassBandwidthTypeMbpsClassInner ---
func packQosProfilesClassBandwidthTypeMbpsClassInnerListFromSdk(ctx context.Context, sdks []network_services.QosProfilesClassBandwidthTypeMbpsClassInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.QosProfilesClassBandwidthTypeMbpsClassInner")
	diags := diag.Diagnostics{}
	var data []models.QosProfilesClassBandwidthTypeMbpsClassInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.QosProfilesClassBandwidthTypeMbpsClassInner
		obj, d := packQosProfilesClassBandwidthTypeMbpsClassInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.QosProfilesClassBandwidthTypeMbpsClassInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.QosProfilesClassBandwidthTypeMbpsClassInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.QosProfilesClassBandwidthTypeMbpsClassInner{}.AttrType(), data)
}

// --- Unpacker for QosProfilesClassBandwidthTypeMbpsClassInnerClassBandwidth ---
func unpackQosProfilesClassBandwidthTypeMbpsClassInnerClassBandwidthToSdk(ctx context.Context, obj types.Object) (*network_services.QosProfilesClassBandwidthTypeMbpsClassInnerClassBandwidth, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.QosProfilesClassBandwidthTypeMbpsClassInnerClassBandwidth", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.QosProfilesClassBandwidthTypeMbpsClassInnerClassBandwidth
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.QosProfilesClassBandwidthTypeMbpsClassInnerClassBandwidth
	var d diag.Diagnostics
	// Handling Primitives
	if !model.EgressGuaranteed.IsNull() && !model.EgressGuaranteed.IsUnknown() {
		val := int32(model.EgressGuaranteed.ValueInt64())
		sdk.EgressGuaranteed = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "EgressGuaranteed", "value": *sdk.EgressGuaranteed})
	}

	// Handling Primitives
	if !model.EgressMax.IsNull() && !model.EgressMax.IsUnknown() {
		val := int32(model.EgressMax.ValueInt64())
		sdk.EgressMax = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "EgressMax", "value": *sdk.EgressMax})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.QosProfilesClassBandwidthTypeMbpsClassInnerClassBandwidth", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for QosProfilesClassBandwidthTypeMbpsClassInnerClassBandwidth ---
func packQosProfilesClassBandwidthTypeMbpsClassInnerClassBandwidthFromSdk(ctx context.Context, sdk network_services.QosProfilesClassBandwidthTypeMbpsClassInnerClassBandwidth) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.QosProfilesClassBandwidthTypeMbpsClassInnerClassBandwidth", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.QosProfilesClassBandwidthTypeMbpsClassInnerClassBandwidth
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.EgressGuaranteed != nil {
		model.EgressGuaranteed = basetypes.NewInt64Value(int64(*sdk.EgressGuaranteed))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "EgressGuaranteed", "value": *sdk.EgressGuaranteed})
	} else {
		model.EgressGuaranteed = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.EgressMax != nil {
		model.EgressMax = basetypes.NewInt64Value(int64(*sdk.EgressMax))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "EgressMax", "value": *sdk.EgressMax})
	} else {
		model.EgressMax = basetypes.NewInt64Null()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.QosProfilesClassBandwidthTypeMbpsClassInnerClassBandwidth{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.QosProfilesClassBandwidthTypeMbpsClassInnerClassBandwidth", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for QosProfilesClassBandwidthTypeMbpsClassInnerClassBandwidth ---
func unpackQosProfilesClassBandwidthTypeMbpsClassInnerClassBandwidthListToSdk(ctx context.Context, list types.List) ([]network_services.QosProfilesClassBandwidthTypeMbpsClassInnerClassBandwidth, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.QosProfilesClassBandwidthTypeMbpsClassInnerClassBandwidth")
	diags := diag.Diagnostics{}
	var data []models.QosProfilesClassBandwidthTypeMbpsClassInnerClassBandwidth
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.QosProfilesClassBandwidthTypeMbpsClassInnerClassBandwidth, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.QosProfilesClassBandwidthTypeMbpsClassInnerClassBandwidth{}.AttrTypes(), &item)
		unpacked, d := unpackQosProfilesClassBandwidthTypeMbpsClassInnerClassBandwidthToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.QosProfilesClassBandwidthTypeMbpsClassInnerClassBandwidth", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for QosProfilesClassBandwidthTypeMbpsClassInnerClassBandwidth ---
func packQosProfilesClassBandwidthTypeMbpsClassInnerClassBandwidthListFromSdk(ctx context.Context, sdks []network_services.QosProfilesClassBandwidthTypeMbpsClassInnerClassBandwidth) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.QosProfilesClassBandwidthTypeMbpsClassInnerClassBandwidth")
	diags := diag.Diagnostics{}
	var data []models.QosProfilesClassBandwidthTypeMbpsClassInnerClassBandwidth

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.QosProfilesClassBandwidthTypeMbpsClassInnerClassBandwidth
		obj, d := packQosProfilesClassBandwidthTypeMbpsClassInnerClassBandwidthFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.QosProfilesClassBandwidthTypeMbpsClassInnerClassBandwidth{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.QosProfilesClassBandwidthTypeMbpsClassInnerClassBandwidth", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.QosProfilesClassBandwidthTypeMbpsClassInnerClassBandwidth{}.AttrType(), data)
}

// --- Unpacker for QosProfilesClassBandwidthTypePercentage ---
func unpackQosProfilesClassBandwidthTypePercentageToSdk(ctx context.Context, obj types.Object) (*network_services.QosProfilesClassBandwidthTypePercentage, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.QosProfilesClassBandwidthTypePercentage", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.QosProfilesClassBandwidthTypePercentage
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.QosProfilesClassBandwidthTypePercentage
	var d diag.Diagnostics
	// Handling Lists
	if !model.Class.IsNull() && !model.Class.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field Class")
		unpacked, d := unpackQosProfilesClassBandwidthTypePercentageClassInnerListToSdk(ctx, model.Class)
		diags.Append(d...)
		sdk.Class = unpacked
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.QosProfilesClassBandwidthTypePercentage", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for QosProfilesClassBandwidthTypePercentage ---
func packQosProfilesClassBandwidthTypePercentageFromSdk(ctx context.Context, sdk network_services.QosProfilesClassBandwidthTypePercentage) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.QosProfilesClassBandwidthTypePercentage", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.QosProfilesClassBandwidthTypePercentage
	var d diag.Diagnostics
	// Handling Lists
	if sdk.Class != nil {
		tflog.Debug(ctx, "Packing list of objects for field Class")
		packed, d := packQosProfilesClassBandwidthTypePercentageClassInnerListFromSdk(ctx, sdk.Class)
		diags.Append(d...)
		model.Class = packed
	} else {
		model.Class = basetypes.NewListNull(models.QosProfilesClassBandwidthTypePercentageClassInner{}.AttrType())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.QosProfilesClassBandwidthTypePercentage{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.QosProfilesClassBandwidthTypePercentage", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for QosProfilesClassBandwidthTypePercentage ---
func unpackQosProfilesClassBandwidthTypePercentageListToSdk(ctx context.Context, list types.List) ([]network_services.QosProfilesClassBandwidthTypePercentage, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.QosProfilesClassBandwidthTypePercentage")
	diags := diag.Diagnostics{}
	var data []models.QosProfilesClassBandwidthTypePercentage
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.QosProfilesClassBandwidthTypePercentage, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.QosProfilesClassBandwidthTypePercentage{}.AttrTypes(), &item)
		unpacked, d := unpackQosProfilesClassBandwidthTypePercentageToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.QosProfilesClassBandwidthTypePercentage", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for QosProfilesClassBandwidthTypePercentage ---
func packQosProfilesClassBandwidthTypePercentageListFromSdk(ctx context.Context, sdks []network_services.QosProfilesClassBandwidthTypePercentage) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.QosProfilesClassBandwidthTypePercentage")
	diags := diag.Diagnostics{}
	var data []models.QosProfilesClassBandwidthTypePercentage

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.QosProfilesClassBandwidthTypePercentage
		obj, d := packQosProfilesClassBandwidthTypePercentageFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.QosProfilesClassBandwidthTypePercentage{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.QosProfilesClassBandwidthTypePercentage", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.QosProfilesClassBandwidthTypePercentage{}.AttrType(), data)
}

// --- Unpacker for QosProfilesClassBandwidthTypePercentageClassInner ---
func unpackQosProfilesClassBandwidthTypePercentageClassInnerToSdk(ctx context.Context, obj types.Object) (*network_services.QosProfilesClassBandwidthTypePercentageClassInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.QosProfilesClassBandwidthTypePercentageClassInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.QosProfilesClassBandwidthTypePercentageClassInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.QosProfilesClassBandwidthTypePercentageClassInner
	var d diag.Diagnostics
	// Handling Objects
	if !model.ClassBandwidth.IsNull() && !model.ClassBandwidth.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field ClassBandwidth")
		unpacked, d := unpackQosProfilesClassBandwidthTypePercentageClassInnerClassBandwidthToSdk(ctx, model.ClassBandwidth)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "ClassBandwidth"})
		}
		if unpacked != nil {
			sdk.ClassBandwidth = unpacked
		}
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
	}

	// Handling Primitives
	if !model.Priority.IsNull() && !model.Priority.IsUnknown() {
		sdk.Priority = model.Priority.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Priority", "value": *sdk.Priority})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.QosProfilesClassBandwidthTypePercentageClassInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for QosProfilesClassBandwidthTypePercentageClassInner ---
func packQosProfilesClassBandwidthTypePercentageClassInnerFromSdk(ctx context.Context, sdk network_services.QosProfilesClassBandwidthTypePercentageClassInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.QosProfilesClassBandwidthTypePercentageClassInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.QosProfilesClassBandwidthTypePercentageClassInner
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.ClassBandwidth != nil {
		tflog.Debug(ctx, "Packing nested object for field ClassBandwidth")
		packed, d := packQosProfilesClassBandwidthTypePercentageClassInnerClassBandwidthFromSdk(ctx, *sdk.ClassBandwidth)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "ClassBandwidth"})
		}
		model.ClassBandwidth = packed
	} else {
		model.ClassBandwidth = basetypes.NewObjectNull(models.QosProfilesClassBandwidthTypePercentageClassInnerClassBandwidth{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Name != nil {
		model.Name = basetypes.NewStringValue(*sdk.Name)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
	} else {
		model.Name = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Priority != nil {
		model.Priority = basetypes.NewStringValue(*sdk.Priority)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Priority", "value": *sdk.Priority})
	} else {
		model.Priority = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.QosProfilesClassBandwidthTypePercentageClassInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.QosProfilesClassBandwidthTypePercentageClassInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for QosProfilesClassBandwidthTypePercentageClassInner ---
func unpackQosProfilesClassBandwidthTypePercentageClassInnerListToSdk(ctx context.Context, list types.List) ([]network_services.QosProfilesClassBandwidthTypePercentageClassInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.QosProfilesClassBandwidthTypePercentageClassInner")
	diags := diag.Diagnostics{}
	var data []models.QosProfilesClassBandwidthTypePercentageClassInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.QosProfilesClassBandwidthTypePercentageClassInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.QosProfilesClassBandwidthTypePercentageClassInner{}.AttrTypes(), &item)
		unpacked, d := unpackQosProfilesClassBandwidthTypePercentageClassInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.QosProfilesClassBandwidthTypePercentageClassInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for QosProfilesClassBandwidthTypePercentageClassInner ---
func packQosProfilesClassBandwidthTypePercentageClassInnerListFromSdk(ctx context.Context, sdks []network_services.QosProfilesClassBandwidthTypePercentageClassInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.QosProfilesClassBandwidthTypePercentageClassInner")
	diags := diag.Diagnostics{}
	var data []models.QosProfilesClassBandwidthTypePercentageClassInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.QosProfilesClassBandwidthTypePercentageClassInner
		obj, d := packQosProfilesClassBandwidthTypePercentageClassInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.QosProfilesClassBandwidthTypePercentageClassInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.QosProfilesClassBandwidthTypePercentageClassInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.QosProfilesClassBandwidthTypePercentageClassInner{}.AttrType(), data)
}

// --- Unpacker for QosProfilesClassBandwidthTypePercentageClassInnerClassBandwidth ---
func unpackQosProfilesClassBandwidthTypePercentageClassInnerClassBandwidthToSdk(ctx context.Context, obj types.Object) (*network_services.QosProfilesClassBandwidthTypePercentageClassInnerClassBandwidth, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.QosProfilesClassBandwidthTypePercentageClassInnerClassBandwidth", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.QosProfilesClassBandwidthTypePercentageClassInnerClassBandwidth
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.QosProfilesClassBandwidthTypePercentageClassInnerClassBandwidth
	var d diag.Diagnostics
	// Handling Primitives
	if !model.EgressGuaranteed.IsNull() && !model.EgressGuaranteed.IsUnknown() {
		val := int32(model.EgressGuaranteed.ValueInt64())
		sdk.EgressGuaranteed = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "EgressGuaranteed", "value": *sdk.EgressGuaranteed})
	}

	// Handling Primitives
	if !model.EgressMax.IsNull() && !model.EgressMax.IsUnknown() {
		val := int32(model.EgressMax.ValueInt64())
		sdk.EgressMax = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "EgressMax", "value": *sdk.EgressMax})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.QosProfilesClassBandwidthTypePercentageClassInnerClassBandwidth", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for QosProfilesClassBandwidthTypePercentageClassInnerClassBandwidth ---
func packQosProfilesClassBandwidthTypePercentageClassInnerClassBandwidthFromSdk(ctx context.Context, sdk network_services.QosProfilesClassBandwidthTypePercentageClassInnerClassBandwidth) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.QosProfilesClassBandwidthTypePercentageClassInnerClassBandwidth", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.QosProfilesClassBandwidthTypePercentageClassInnerClassBandwidth
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.EgressGuaranteed != nil {
		model.EgressGuaranteed = basetypes.NewInt64Value(int64(*sdk.EgressGuaranteed))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "EgressGuaranteed", "value": *sdk.EgressGuaranteed})
	} else {
		model.EgressGuaranteed = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.EgressMax != nil {
		model.EgressMax = basetypes.NewInt64Value(int64(*sdk.EgressMax))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "EgressMax", "value": *sdk.EgressMax})
	} else {
		model.EgressMax = basetypes.NewInt64Null()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.QosProfilesClassBandwidthTypePercentageClassInnerClassBandwidth{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.QosProfilesClassBandwidthTypePercentageClassInnerClassBandwidth", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for QosProfilesClassBandwidthTypePercentageClassInnerClassBandwidth ---
func unpackQosProfilesClassBandwidthTypePercentageClassInnerClassBandwidthListToSdk(ctx context.Context, list types.List) ([]network_services.QosProfilesClassBandwidthTypePercentageClassInnerClassBandwidth, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.QosProfilesClassBandwidthTypePercentageClassInnerClassBandwidth")
	diags := diag.Diagnostics{}
	var data []models.QosProfilesClassBandwidthTypePercentageClassInnerClassBandwidth
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.QosProfilesClassBandwidthTypePercentageClassInnerClassBandwidth, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.QosProfilesClassBandwidthTypePercentageClassInnerClassBandwidth{}.AttrTypes(), &item)
		unpacked, d := unpackQosProfilesClassBandwidthTypePercentageClassInnerClassBandwidthToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.QosProfilesClassBandwidthTypePercentageClassInnerClassBandwidth", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for QosProfilesClassBandwidthTypePercentageClassInnerClassBandwidth ---
func packQosProfilesClassBandwidthTypePercentageClassInnerClassBandwidthListFromSdk(ctx context.Context, sdks []network_services.QosProfilesClassBandwidthTypePercentageClassInnerClassBandwidth) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.QosProfilesClassBandwidthTypePercentageClassInnerClassBandwidth")
	diags := diag.Diagnostics{}
	var data []models.QosProfilesClassBandwidthTypePercentageClassInnerClassBandwidth

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.QosProfilesClassBandwidthTypePercentageClassInnerClassBandwidth
		obj, d := packQosProfilesClassBandwidthTypePercentageClassInnerClassBandwidthFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.QosProfilesClassBandwidthTypePercentageClassInnerClassBandwidth{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.QosProfilesClassBandwidthTypePercentageClassInnerClassBandwidth", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.QosProfilesClassBandwidthTypePercentageClassInnerClassBandwidth{}.AttrType(), data)
}
