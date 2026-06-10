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
