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

// --- Unpacker for LicenseResult ---
func unpackLicenseResultToSdk(ctx context.Context, obj types.Object) (*network_services.LicenseResult, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.LicenseResult", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.LicenseResult
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.LicenseResult
	var d diag.Diagnostics

	// Handling Lists
	if !model.ConfiguredLicenses.IsNull() && !model.ConfiguredLicenses.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field ConfiguredLicenses")
		unpacked, d := unpackLicenseInfoListToSdk(ctx, model.ConfiguredLicenses)
		diags.Append(d...)
		sdk.ConfiguredLicenses = unpacked
	}

	// Handling Lists
	if !model.LicenseModel.IsNull() && !model.LicenseModel.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field LicenseModel")
		diags.Append(model.LicenseModel.ElementsAs(ctx, &sdk.LicenseModel, false)...)
	}

	// Handling Primitives
	if !model.OperationalLicense.IsNull() && !model.OperationalLicense.IsUnknown() {
		sdk.OperationalLicense = model.OperationalLicense.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "OperationalLicense", "value": *sdk.OperationalLicense})
	}

	// Handling Lists
	if !model.PurchasedLicenses.IsNull() && !model.PurchasedLicenses.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field PurchasedLicenses")
		unpacked, d := unpackLicenseInfoListToSdk(ctx, model.PurchasedLicenses)
		diags.Append(d...)
		sdk.PurchasedLicenses = unpacked
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.LicenseResult", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for LicenseResult ---
func packLicenseResultFromSdk(ctx context.Context, sdk network_services.LicenseResult) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.LicenseResult", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.LicenseResult
	var d diag.Diagnostics
	// Handling Lists
	if sdk.ConfiguredLicenses != nil {
		tflog.Debug(ctx, "Packing list of objects for field ConfiguredLicenses")
		packed, d := packLicenseInfoListFromSdk(ctx, sdk.ConfiguredLicenses)
		diags.Append(d...)
		model.ConfiguredLicenses = packed
	} else {
		model.ConfiguredLicenses = basetypes.NewListNull(models.LicenseInfo{}.AttrType())
	}
	// Handling Lists
	if sdk.LicenseModel != nil {
		tflog.Debug(ctx, "Packing list of primitives for field LicenseModel")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.LicenseModel, d = basetypes.NewListValueFrom(ctx, elemType, sdk.LicenseModel)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.LicenseModel = basetypes.NewListNull(elemType)
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.OperationalLicense != nil {
		model.OperationalLicense = basetypes.NewStringValue(*sdk.OperationalLicense)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "OperationalLicense", "value": *sdk.OperationalLicense})
	} else {
		model.OperationalLicense = basetypes.NewStringNull()
	}
	// Handling Lists
	if sdk.PurchasedLicenses != nil {
		tflog.Debug(ctx, "Packing list of objects for field PurchasedLicenses")
		packed, d := packLicenseInfoListFromSdk(ctx, sdk.PurchasedLicenses)
		diags.Append(d...)
		model.PurchasedLicenses = packed
	} else {
		model.PurchasedLicenses = basetypes.NewListNull(models.LicenseInfo{}.AttrType())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.LicenseResult{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.LicenseResult", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for LicenseResult ---
func unpackLicenseResultListToSdk(ctx context.Context, list types.List) ([]network_services.LicenseResult, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.LicenseResult")
	diags := diag.Diagnostics{}
	var data []models.LicenseResult
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.LicenseResult, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.LicenseResult{}.AttrTypes(), &item)
		unpacked, d := unpackLicenseResultToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.LicenseResult", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for LicenseResult ---
func packLicenseResultListFromSdk(ctx context.Context, sdks []network_services.LicenseResult) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.LicenseResult")
	diags := diag.Diagnostics{}
	var data []models.LicenseResult

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.LicenseResult
		obj, d := packLicenseResultFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.LicenseResult{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.LicenseResult", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.LicenseResult{}.AttrType(), data)
}

// --- Unpacker for LicenseInfo ---
func unpackLicenseInfoToSdk(ctx context.Context, obj types.Object) (*network_services.LicenseInfo, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.LicenseInfo", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.LicenseInfo
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.LicenseInfo
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Count.IsNull() && !model.Count.IsUnknown() {
		val := int32(model.Count.ValueInt64())
		sdk.Count = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Count", "value": *sdk.Count})
	}

	// Handling Primitives
	if !model.LicenseType.IsNull() && !model.LicenseType.IsUnknown() {
		sdk.LicenseType = model.LicenseType.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "LicenseType", "value": *sdk.LicenseType})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.LicenseInfo", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for LicenseInfo ---
func packLicenseInfoFromSdk(ctx context.Context, sdk network_services.LicenseInfo) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.LicenseInfo", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.LicenseInfo
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Count != nil {
		model.Count = basetypes.NewInt64Value(int64(*sdk.Count))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Count", "value": *sdk.Count})
	} else {
		model.Count = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.LicenseType != nil {
		model.LicenseType = basetypes.NewStringValue(*sdk.LicenseType)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "LicenseType", "value": *sdk.LicenseType})
	} else {
		model.LicenseType = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.LicenseInfo{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.LicenseInfo", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for LicenseInfo ---
func unpackLicenseInfoListToSdk(ctx context.Context, list types.List) ([]network_services.LicenseInfo, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.LicenseInfo")
	diags := diag.Diagnostics{}
	var data []models.LicenseInfo
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.LicenseInfo, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.LicenseInfo{}.AttrTypes(), &item)
		unpacked, d := unpackLicenseInfoToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.LicenseInfo", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for LicenseInfo ---
func packLicenseInfoListFromSdk(ctx context.Context, sdks []network_services.LicenseInfo) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.LicenseInfo")
	diags := diag.Diagnostics{}
	var data []models.LicenseInfo

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.LicenseInfo
		obj, d := packLicenseInfoFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.LicenseInfo{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.LicenseInfo", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.LicenseInfo{}.AttrType(), data)
}
