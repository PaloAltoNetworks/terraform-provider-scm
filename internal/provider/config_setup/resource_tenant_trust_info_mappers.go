package provider

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/config_setup"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/config_setup"
)

// --- Unpacker for TenantTrustInfo ---
func unpackTenantTrustInfoToSdk(ctx context.Context, obj types.Object) (*config_setup.TenantTrustInfo, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.TenantTrustInfo", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.TenantTrustInfo
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk config_setup.TenantTrustInfo
	var d diag.Diagnostics

	// Handling Primitives
	if !model.Created.IsNull() && !model.Created.IsUnknown() {
		if t, parseErr := time.Parse(time.RFC3339, model.Created.ValueString()); parseErr == nil {
			sdk.Created = &t
		}
	}

	// Handling Primitives
	if !model.CreatedBy.IsNull() && !model.CreatedBy.IsUnknown() {
		sdk.CreatedBy = model.CreatedBy.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "CreatedBy", "value": *sdk.CreatedBy})
	}

	// Handling Primitives
	if !model.CurrentStatus.IsNull() && !model.CurrentStatus.IsUnknown() {
		sdk.CurrentStatus = model.CurrentStatus.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "CurrentStatus", "value": *sdk.CurrentStatus})
	}

	// Handling Primitives
	if !model.DonorCluster.IsNull() && !model.DonorCluster.IsUnknown() {
		sdk.DonorCluster = model.DonorCluster.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DonorCluster", "value": *sdk.DonorCluster})
	}

	// Handling Primitives
	if !model.DonorMsgUuid.IsNull() && !model.DonorMsgUuid.IsUnknown() {
		sdk.DonorMsgUuid = model.DonorMsgUuid.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DonorMsgUuid", "value": *sdk.DonorMsgUuid})
	}

	// Handling Primitives
	if !model.DonorProject.IsNull() && !model.DonorProject.IsUnknown() {
		sdk.DonorProject = model.DonorProject.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DonorProject", "value": *sdk.DonorProject})
	}

	// Handling Primitives
	if !model.DonorRegion.IsNull() && !model.DonorRegion.IsUnknown() {
		sdk.DonorRegion = model.DonorRegion.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DonorRegion", "value": *sdk.DonorRegion})
	}

	// Handling Primitives
	if !model.DonorTenantId.IsNull() && !model.DonorTenantId.IsUnknown() {
		sdk.DonorTenantId = model.DonorTenantId.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DonorTenantId", "value": *sdk.DonorTenantId})
	}

	// Handling Primitives
	if !model.DonorTenantName.IsNull() && !model.DonorTenantName.IsUnknown() {
		sdk.DonorTenantName = model.DonorTenantName.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DonorTenantName", "value": *sdk.DonorTenantName})
	}

	// Handling Primitives
	if !model.DonorTrustInfoId.IsNull() && !model.DonorTrustInfoId.IsUnknown() {
		val := int32(model.DonorTrustInfoId.ValueInt64())
		sdk.DonorTrustInfoId = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DonorTrustInfoId", "value": *sdk.DonorTrustInfoId})
	}

	// Handling Primitives
	if !model.DonorTsg.IsNull() && !model.DonorTsg.IsUnknown() {
		sdk.DonorTsg = model.DonorTsg.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DonorTsg", "value": *sdk.DonorTsg})
	}

	// Handling Primitives
	if !model.ErrorDetails.IsNull() && !model.ErrorDetails.IsUnknown() {
		sdk.ErrorDetails = model.ErrorDetails.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "ErrorDetails", "value": *sdk.ErrorDetails})
	}

	// Handling Primitives
	if !model.LastUpdated.IsNull() && !model.LastUpdated.IsUnknown() {
		if t, parseErr := time.Parse(time.RFC3339, model.LastUpdated.ValueString()); parseErr == nil {
			sdk.LastUpdated = &t
		}
	}

	// Handling Primitives
	if !model.Psk.IsNull() && !model.Psk.IsUnknown() {
		sdk.Psk = model.Psk.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Psk", "value": *sdk.Psk})
	}

	// Handling Primitives
	if !model.RecipientCluster.IsNull() && !model.RecipientCluster.IsUnknown() {
		sdk.RecipientCluster = model.RecipientCluster.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "RecipientCluster", "value": *sdk.RecipientCluster})
	}

	// Handling Primitives
	if !model.RecipientMsgUuid.IsNull() && !model.RecipientMsgUuid.IsUnknown() {
		sdk.RecipientMsgUuid = model.RecipientMsgUuid.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "RecipientMsgUuid", "value": *sdk.RecipientMsgUuid})
	}

	// Handling Primitives
	if !model.RecipientProject.IsNull() && !model.RecipientProject.IsUnknown() {
		sdk.RecipientProject = model.RecipientProject.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "RecipientProject", "value": *sdk.RecipientProject})
	}

	// Handling Primitives
	if !model.RecipientRegion.IsNull() && !model.RecipientRegion.IsUnknown() {
		sdk.RecipientRegion = model.RecipientRegion.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "RecipientRegion", "value": *sdk.RecipientRegion})
	}

	// Handling Primitives
	if !model.RecipientTenantId.IsNull() && !model.RecipientTenantId.IsUnknown() {
		sdk.RecipientTenantId = model.RecipientTenantId.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "RecipientTenantId", "value": *sdk.RecipientTenantId})
	}

	// Handling Primitives
	if !model.RecipientTenantName.IsNull() && !model.RecipientTenantName.IsUnknown() {
		sdk.RecipientTenantName = model.RecipientTenantName.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "RecipientTenantName", "value": *sdk.RecipientTenantName})
	}

	// Handling Primitives
	if !model.RecipientTrustInfoId.IsNull() && !model.RecipientTrustInfoId.IsUnknown() {
		val := int32(model.RecipientTrustInfoId.ValueInt64())
		sdk.RecipientTrustInfoId = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "RecipientTrustInfoId", "value": *sdk.RecipientTrustInfoId})
	}

	// Handling Primitives
	if !model.RecipientTsg.IsNull() && !model.RecipientTsg.IsUnknown() {
		sdk.RecipientTsg = model.RecipientTsg.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "RecipientTsg", "value": *sdk.RecipientTsg})
	}

	// Handling Primitives
	if !model.TrustId.IsNull() && !model.TrustId.IsUnknown() {
		val := int32(model.TrustId.ValueInt64())
		sdk.TrustId = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "TrustId", "value": *sdk.TrustId})
	}

	// Handling Primitives
	if !model.UpdatedBy.IsNull() && !model.UpdatedBy.IsUnknown() {
		sdk.UpdatedBy = model.UpdatedBy.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "UpdatedBy", "value": *sdk.UpdatedBy})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.TenantTrustInfo", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for TenantTrustInfo ---
func packTenantTrustInfoFromSdk(ctx context.Context, sdk config_setup.TenantTrustInfo) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.TenantTrustInfo", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.TenantTrustInfo
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Created != nil {
		model.Created = basetypes.NewStringValue(sdk.Created.Format(time.RFC3339))
		tflog.Debug(ctx, "Packed time pointer", map[string]interface{}{"field": "Created", "value": sdk.Created.Format(time.RFC3339)})
	} else {
		model.Created = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.CreatedBy != nil {
		model.CreatedBy = basetypes.NewStringValue(*sdk.CreatedBy)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "CreatedBy", "value": *sdk.CreatedBy})
	} else {
		model.CreatedBy = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.CurrentStatus != nil {
		model.CurrentStatus = basetypes.NewStringValue(*sdk.CurrentStatus)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "CurrentStatus", "value": *sdk.CurrentStatus})
	} else {
		model.CurrentStatus = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.DonorCluster != nil {
		model.DonorCluster = basetypes.NewStringValue(*sdk.DonorCluster)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DonorCluster", "value": *sdk.DonorCluster})
	} else {
		model.DonorCluster = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.DonorMsgUuid != nil {
		model.DonorMsgUuid = basetypes.NewStringValue(*sdk.DonorMsgUuid)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DonorMsgUuid", "value": *sdk.DonorMsgUuid})
	} else {
		model.DonorMsgUuid = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.DonorProject != nil {
		model.DonorProject = basetypes.NewStringValue(*sdk.DonorProject)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DonorProject", "value": *sdk.DonorProject})
	} else {
		model.DonorProject = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.DonorRegion != nil {
		model.DonorRegion = basetypes.NewStringValue(*sdk.DonorRegion)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DonorRegion", "value": *sdk.DonorRegion})
	} else {
		model.DonorRegion = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.DonorTenantId != nil {
		model.DonorTenantId = basetypes.NewStringValue(*sdk.DonorTenantId)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DonorTenantId", "value": *sdk.DonorTenantId})
	} else {
		model.DonorTenantId = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.DonorTenantName != nil {
		model.DonorTenantName = basetypes.NewStringValue(*sdk.DonorTenantName)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DonorTenantName", "value": *sdk.DonorTenantName})
	} else {
		model.DonorTenantName = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.DonorTrustInfoId != nil {
		model.DonorTrustInfoId = basetypes.NewInt64Value(int64(*sdk.DonorTrustInfoId))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DonorTrustInfoId", "value": *sdk.DonorTrustInfoId})
	} else {
		model.DonorTrustInfoId = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.DonorTsg != nil {
		model.DonorTsg = basetypes.NewStringValue(*sdk.DonorTsg)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DonorTsg", "value": *sdk.DonorTsg})
	} else {
		model.DonorTsg = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.ErrorDetails != nil {
		model.ErrorDetails = basetypes.NewStringValue(*sdk.ErrorDetails)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "ErrorDetails", "value": *sdk.ErrorDetails})
	} else {
		model.ErrorDetails = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.LastUpdated != nil {
		model.LastUpdated = basetypes.NewStringValue(sdk.LastUpdated.Format(time.RFC3339))
		tflog.Debug(ctx, "Packed time pointer", map[string]interface{}{"field": "LastUpdated", "value": sdk.LastUpdated.Format(time.RFC3339)})
	} else {
		model.LastUpdated = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Psk != nil {
		model.Psk = basetypes.NewStringValue(*sdk.Psk)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Psk", "value": *sdk.Psk})
	} else {
		model.Psk = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.RecipientCluster != nil {
		model.RecipientCluster = basetypes.NewStringValue(*sdk.RecipientCluster)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "RecipientCluster", "value": *sdk.RecipientCluster})
	} else {
		model.RecipientCluster = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.RecipientMsgUuid != nil {
		model.RecipientMsgUuid = basetypes.NewStringValue(*sdk.RecipientMsgUuid)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "RecipientMsgUuid", "value": *sdk.RecipientMsgUuid})
	} else {
		model.RecipientMsgUuid = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.RecipientProject != nil {
		model.RecipientProject = basetypes.NewStringValue(*sdk.RecipientProject)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "RecipientProject", "value": *sdk.RecipientProject})
	} else {
		model.RecipientProject = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.RecipientRegion != nil {
		model.RecipientRegion = basetypes.NewStringValue(*sdk.RecipientRegion)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "RecipientRegion", "value": *sdk.RecipientRegion})
	} else {
		model.RecipientRegion = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.RecipientTenantId != nil {
		model.RecipientTenantId = basetypes.NewStringValue(*sdk.RecipientTenantId)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "RecipientTenantId", "value": *sdk.RecipientTenantId})
	} else {
		model.RecipientTenantId = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.RecipientTenantName != nil {
		model.RecipientTenantName = basetypes.NewStringValue(*sdk.RecipientTenantName)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "RecipientTenantName", "value": *sdk.RecipientTenantName})
	} else {
		model.RecipientTenantName = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.RecipientTrustInfoId != nil {
		model.RecipientTrustInfoId = basetypes.NewInt64Value(int64(*sdk.RecipientTrustInfoId))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "RecipientTrustInfoId", "value": *sdk.RecipientTrustInfoId})
	} else {
		model.RecipientTrustInfoId = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.RecipientTsg != nil {
		model.RecipientTsg = basetypes.NewStringValue(*sdk.RecipientTsg)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "RecipientTsg", "value": *sdk.RecipientTsg})
	} else {
		model.RecipientTsg = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.TrustId != nil {
		model.TrustId = basetypes.NewInt64Value(int64(*sdk.TrustId))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "TrustId", "value": *sdk.TrustId})
	} else {
		model.TrustId = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.UpdatedBy != nil {
		model.UpdatedBy = basetypes.NewStringValue(*sdk.UpdatedBy)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "UpdatedBy", "value": *sdk.UpdatedBy})
	} else {
		model.UpdatedBy = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.TenantTrustInfo{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.TenantTrustInfo", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for TenantTrustInfo ---
func unpackTenantTrustInfoListToSdk(ctx context.Context, list types.List) ([]config_setup.TenantTrustInfo, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.TenantTrustInfo")
	diags := diag.Diagnostics{}
	var data []models.TenantTrustInfo
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]config_setup.TenantTrustInfo, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.TenantTrustInfo{}.AttrTypes(), &item)
		unpacked, d := unpackTenantTrustInfoToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.TenantTrustInfo", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for TenantTrustInfo ---
func packTenantTrustInfoListFromSdk(ctx context.Context, sdks []config_setup.TenantTrustInfo) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.TenantTrustInfo")
	diags := diag.Diagnostics{}
	var data []models.TenantTrustInfo

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.TenantTrustInfo
		obj, d := packTenantTrustInfoFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.TenantTrustInfo{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.TenantTrustInfo", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.TenantTrustInfo{}.AttrType(), data)
}
