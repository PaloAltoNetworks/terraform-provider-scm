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

// --- Unpacker for TenantStatus ---
func unpackTenantStatusToSdk(ctx context.Context, obj types.Object) (*ztna_connector_all.TenantStatus, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.TenantStatus", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.TenantStatus
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk ztna_connector_all.TenantStatus
	var d diag.Diagnostics

	// Handling Primitives
	if !model.Status.IsNull() && !model.Status.IsUnknown() {
		sdk.Status = model.Status.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Status", "value": sdk.Status})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.TenantStatus", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for TenantStatus ---
func packTenantStatusFromSdk(ctx context.Context, sdk ztna_connector_all.TenantStatus) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.TenantStatus", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.TenantStatus
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.Status = basetypes.NewStringValue(sdk.Status)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Status", "value": sdk.Status})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.TenantStatus{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.TenantStatus", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for TenantStatus ---
func unpackTenantStatusListToSdk(ctx context.Context, list types.List) ([]ztna_connector_all.TenantStatus, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.TenantStatus")
	diags := diag.Diagnostics{}
	var data []models.TenantStatus
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]ztna_connector_all.TenantStatus, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.TenantStatus{}.AttrTypes(), &item)
		unpacked, d := unpackTenantStatusToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.TenantStatus", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for TenantStatus ---
func packTenantStatusListFromSdk(ctx context.Context, sdks []ztna_connector_all.TenantStatus) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.TenantStatus")
	diags := diag.Diagnostics{}
	var data []models.TenantStatus

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.TenantStatus
		obj, d := packTenantStatusFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.TenantStatus{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.TenantStatus", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.TenantStatus{}.AttrType(), data)
}
