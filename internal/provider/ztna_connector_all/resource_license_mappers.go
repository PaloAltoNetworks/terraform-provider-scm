package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/ztna_connector_all"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/ztna_connector_all"
)

// --- Unpacker for License ---
func unpackLicenseToSdk(ctx context.Context, obj types.Object) (*ztna_connector_all.License, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.License", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.License
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk ztna_connector_all.License
	var d diag.Diagnostics

	// Handling Primitives
	if !model.Applications.IsNull() && !model.Applications.IsUnknown() {
		val := float32(model.Applications.ValueFloat64())
		sdk.Applications = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Applications", "value": *sdk.Applications})
	}

	// Handling Primitives
	if !model.Connectors.IsNull() && !model.Connectors.IsUnknown() {
		val := float32(model.Connectors.ValueFloat64())
		sdk.Connectors = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Connectors", "value": *sdk.Connectors})
	}

	// Handling Primitives
	if !model.Expiry.IsNull() && !model.Expiry.IsUnknown() {
		sdk.Expiry = model.Expiry.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Expiry", "value": *sdk.Expiry})
	}

	// Handling Primitives
	if !model.LicenseName.IsNull() && !model.LicenseName.IsUnknown() {
		sdk.LicenseName = model.LicenseName.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "LicenseName", "value": *sdk.LicenseName})
	}

	// Handling Primitives
	if !model.MaxApplications.IsNull() && !model.MaxApplications.IsUnknown() {
		val := float32(model.MaxApplications.ValueFloat64())
		sdk.MaxApplications = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "MaxApplications", "value": *sdk.MaxApplications})
	}

	// Handling Primitives
	if !model.MaxConnectors.IsNull() && !model.MaxConnectors.IsUnknown() {
		val := float32(model.MaxConnectors.ValueFloat64())
		sdk.MaxConnectors = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "MaxConnectors", "value": *sdk.MaxConnectors})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.License", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for License ---
func packLicenseFromSdk(ctx context.Context, sdk ztna_connector_all.License) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.License", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.License
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Applications != nil {
		model.Applications = basetypes.NewFloat64Value(float64(*sdk.Applications))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Applications", "value": *sdk.Applications})
	} else {
		model.Applications = basetypes.NewFloat64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Connectors != nil {
		model.Connectors = basetypes.NewFloat64Value(float64(*sdk.Connectors))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Connectors", "value": *sdk.Connectors})
	} else {
		model.Connectors = basetypes.NewFloat64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Expiry != nil {
		model.Expiry = basetypes.NewStringValue(*sdk.Expiry)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Expiry", "value": *sdk.Expiry})
	} else {
		model.Expiry = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.LicenseName != nil {
		model.LicenseName = basetypes.NewStringValue(*sdk.LicenseName)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "LicenseName", "value": *sdk.LicenseName})
	} else {
		model.LicenseName = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.MaxApplications != nil {
		model.MaxApplications = basetypes.NewFloat64Value(float64(*sdk.MaxApplications))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "MaxApplications", "value": *sdk.MaxApplications})
	} else {
		model.MaxApplications = basetypes.NewFloat64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.MaxConnectors != nil {
		model.MaxConnectors = basetypes.NewFloat64Value(float64(*sdk.MaxConnectors))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "MaxConnectors", "value": *sdk.MaxConnectors})
	} else {
		model.MaxConnectors = basetypes.NewFloat64Null()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.License{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.License", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for License ---
func unpackLicenseListToSdk(ctx context.Context, list types.List) ([]ztna_connector_all.License, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.License")
	diags := diag.Diagnostics{}
	var data []models.License
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]ztna_connector_all.License, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.License{}.AttrTypes(), &item)
		unpacked, d := unpackLicenseToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.License", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for License ---
func packLicenseListFromSdk(ctx context.Context, sdks []ztna_connector_all.License) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.License")
	diags := diag.Diagnostics{}
	var data []models.License

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.License
		obj, d := packLicenseFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.License{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.License", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.License{}.AttrType(), data)
}
