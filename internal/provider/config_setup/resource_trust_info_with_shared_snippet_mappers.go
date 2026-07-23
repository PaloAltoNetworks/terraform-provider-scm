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

// --- Unpacker for TrustInfoWithSharedSnippets ---
func unpackTrustInfoWithSharedSnippetsToSdk(ctx context.Context, obj types.Object) (*config_setup.TrustInfoWithSharedSnippets, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.TrustInfoWithSharedSnippets", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.TrustInfoWithSharedSnippets
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk config_setup.TrustInfoWithSharedSnippets
	var d diag.Diagnostics

	// Handling Primitives
	if !model.Created.IsNull() && !model.Created.IsUnknown() {
		if t, parseErr := time.Parse(time.RFC3339, model.Created.ValueString()); parseErr == nil {
			sdk.Created = &t
		}
	}

	// Handling Primitives
	if !model.DonorCreated.IsNull() && !model.DonorCreated.IsUnknown() {
		val := int32(model.DonorCreated.ValueInt64())
		sdk.DonorCreated = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DonorCreated", "value": *sdk.DonorCreated})
	}

	// Handling Primitives
	if !model.DonorSnippetFileId.IsNull() && !model.DonorSnippetFileId.IsUnknown() {
		val := int32(model.DonorSnippetFileId.ValueInt64())
		sdk.DonorSnippetFileId = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DonorSnippetFileId", "value": *sdk.DonorSnippetFileId})
	}

	// Handling Primitives
	if !model.DonorSnippetVersion.IsNull() && !model.DonorSnippetVersion.IsUnknown() {
		val := int32(model.DonorSnippetVersion.ValueInt64())
		sdk.DonorSnippetVersion = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DonorSnippetVersion", "value": *sdk.DonorSnippetVersion})
	}

	// Handling Primitives
	if !model.DonorTsg.IsNull() && !model.DonorTsg.IsUnknown() {
		sdk.DonorTsg = model.DonorTsg.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DonorTsg", "value": *sdk.DonorTsg})
	}

	// Handling Primitives
	if !model.Error.IsNull() && !model.Error.IsUnknown() {
		sdk.Error = model.Error.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Error", "value": *sdk.Error})
	}

	// Handling Primitives
	if !model.Id.IsNull() && !model.Id.IsUnknown() {
		val := int32(model.Id.ValueInt64())
		sdk.Id = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Id", "value": *sdk.Id})
	}

	// Handling Primitives
	if !model.LastUpdated.IsNull() && !model.LastUpdated.IsUnknown() {
		if t, parseErr := time.Parse(time.RFC3339, model.LastUpdated.ValueString()); parseErr == nil {
			sdk.LastUpdated = &t
		}
	}

	// Handling Primitives
	if !model.MsgUuid.IsNull() && !model.MsgUuid.IsUnknown() {
		sdk.MsgUuid = model.MsgUuid.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "MsgUuid", "value": *sdk.MsgUuid})
	}

	// Handling Primitives
	if !model.RecipientPausedUpdate.IsNull() && !model.RecipientPausedUpdate.IsUnknown() {
		val := int32(model.RecipientPausedUpdate.ValueInt64())
		sdk.RecipientPausedUpdate = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "RecipientPausedUpdate", "value": *sdk.RecipientPausedUpdate})
	}

	// Handling Primitives
	if !model.RecipientSnippetFileId.IsNull() && !model.RecipientSnippetFileId.IsUnknown() {
		val := int32(model.RecipientSnippetFileId.ValueInt64())
		sdk.RecipientSnippetFileId = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "RecipientSnippetFileId", "value": *sdk.RecipientSnippetFileId})
	}

	// Handling Primitives
	if !model.RecipientSnippetVersion.IsNull() && !model.RecipientSnippetVersion.IsUnknown() {
		val := int32(model.RecipientSnippetVersion.ValueInt64())
		sdk.RecipientSnippetVersion = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "RecipientSnippetVersion", "value": *sdk.RecipientSnippetVersion})
	}

	// Handling Primitives
	if !model.RecipientTsg.IsNull() && !model.RecipientTsg.IsUnknown() {
		sdk.RecipientTsg = model.RecipientTsg.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "RecipientTsg", "value": *sdk.RecipientTsg})
	}

	// Handling Primitives
	if !model.RecipientValidateBeforeUpdate.IsNull() && !model.RecipientValidateBeforeUpdate.IsUnknown() {
		val := int32(model.RecipientValidateBeforeUpdate.ValueInt64())
		sdk.RecipientValidateBeforeUpdate = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "RecipientValidateBeforeUpdate", "value": *sdk.RecipientValidateBeforeUpdate})
	}

	// Handling Lists
	if !model.SharedSnippets.IsNull() && !model.SharedSnippets.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field SharedSnippets")
		unpacked, d := unpackSnippetShareInfoListToSdk(ctx, model.SharedSnippets)
		diags.Append(d...)
		sdk.SharedSnippets = unpacked
	}

	// Handling Primitives
	if !model.SnippetName.IsNull() && !model.SnippetName.IsUnknown() {
		sdk.SnippetName = model.SnippetName.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "SnippetName", "value": *sdk.SnippetName})
	}

	// Handling Primitives
	if !model.SnippetUuid.IsNull() && !model.SnippetUuid.IsUnknown() {
		sdk.SnippetUuid = model.SnippetUuid.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "SnippetUuid", "value": *sdk.SnippetUuid})
	}

	// Handling Primitives
	if !model.Status.IsNull() && !model.Status.IsUnknown() {
		sdk.Status = model.Status.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Status", "value": *sdk.Status})
	}

	// Handling Primitives
	if !model.UpdatedBy.IsNull() && !model.UpdatedBy.IsUnknown() {
		sdk.UpdatedBy = model.UpdatedBy.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "UpdatedBy", "value": *sdk.UpdatedBy})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.TrustInfoWithSharedSnippets", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for TrustInfoWithSharedSnippets ---
func packTrustInfoWithSharedSnippetsFromSdk(ctx context.Context, sdk config_setup.TrustInfoWithSharedSnippets) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.TrustInfoWithSharedSnippets", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.TrustInfoWithSharedSnippets
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
	if sdk.DonorCreated != nil {
		model.DonorCreated = basetypes.NewInt64Value(int64(*sdk.DonorCreated))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DonorCreated", "value": *sdk.DonorCreated})
	} else {
		model.DonorCreated = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.DonorSnippetFileId != nil {
		model.DonorSnippetFileId = basetypes.NewInt64Value(int64(*sdk.DonorSnippetFileId))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DonorSnippetFileId", "value": *sdk.DonorSnippetFileId})
	} else {
		model.DonorSnippetFileId = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.DonorSnippetVersion != nil {
		model.DonorSnippetVersion = basetypes.NewInt64Value(int64(*sdk.DonorSnippetVersion))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DonorSnippetVersion", "value": *sdk.DonorSnippetVersion})
	} else {
		model.DonorSnippetVersion = basetypes.NewInt64Null()
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
	if sdk.Error != nil {
		model.Error = basetypes.NewStringValue(*sdk.Error)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Error", "value": *sdk.Error})
	} else {
		model.Error = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Id != nil {
		model.Id = basetypes.NewInt64Value(int64(*sdk.Id))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Id", "value": *sdk.Id})
	} else {
		model.Id = basetypes.NewInt64Null()
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
	if sdk.MsgUuid != nil {
		model.MsgUuid = basetypes.NewStringValue(*sdk.MsgUuid)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "MsgUuid", "value": *sdk.MsgUuid})
	} else {
		model.MsgUuid = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.RecipientPausedUpdate != nil {
		model.RecipientPausedUpdate = basetypes.NewInt64Value(int64(*sdk.RecipientPausedUpdate))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "RecipientPausedUpdate", "value": *sdk.RecipientPausedUpdate})
	} else {
		model.RecipientPausedUpdate = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.RecipientSnippetFileId != nil {
		model.RecipientSnippetFileId = basetypes.NewInt64Value(int64(*sdk.RecipientSnippetFileId))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "RecipientSnippetFileId", "value": *sdk.RecipientSnippetFileId})
	} else {
		model.RecipientSnippetFileId = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.RecipientSnippetVersion != nil {
		model.RecipientSnippetVersion = basetypes.NewInt64Value(int64(*sdk.RecipientSnippetVersion))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "RecipientSnippetVersion", "value": *sdk.RecipientSnippetVersion})
	} else {
		model.RecipientSnippetVersion = basetypes.NewInt64Null()
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
	if sdk.RecipientValidateBeforeUpdate != nil {
		model.RecipientValidateBeforeUpdate = basetypes.NewInt64Value(int64(*sdk.RecipientValidateBeforeUpdate))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "RecipientValidateBeforeUpdate", "value": *sdk.RecipientValidateBeforeUpdate})
	} else {
		model.RecipientValidateBeforeUpdate = basetypes.NewInt64Null()
	}
	// Handling Lists
	if sdk.SharedSnippets != nil {
		tflog.Debug(ctx, "Packing list of objects for field SharedSnippets")
		packed, d := packSnippetShareInfoListFromSdk(ctx, sdk.SharedSnippets)
		diags.Append(d...)
		model.SharedSnippets = packed
	} else {
		model.SharedSnippets = basetypes.NewListNull(models.SnippetShareInfo{}.AttrType())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.SnippetName != nil {
		model.SnippetName = basetypes.NewStringValue(*sdk.SnippetName)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "SnippetName", "value": *sdk.SnippetName})
	} else {
		model.SnippetName = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.SnippetUuid != nil {
		model.SnippetUuid = basetypes.NewStringValue(*sdk.SnippetUuid)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "SnippetUuid", "value": *sdk.SnippetUuid})
	} else {
		model.SnippetUuid = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Status != nil {
		model.Status = basetypes.NewStringValue(*sdk.Status)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Status", "value": *sdk.Status})
	} else {
		model.Status = basetypes.NewStringNull()
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

	obj, d := types.ObjectValueFrom(ctx, models.TrustInfoWithSharedSnippets{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.TrustInfoWithSharedSnippets", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for TrustInfoWithSharedSnippets ---
func unpackTrustInfoWithSharedSnippetsListToSdk(ctx context.Context, list types.List) ([]config_setup.TrustInfoWithSharedSnippets, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.TrustInfoWithSharedSnippets")
	diags := diag.Diagnostics{}
	var data []models.TrustInfoWithSharedSnippets
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]config_setup.TrustInfoWithSharedSnippets, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.TrustInfoWithSharedSnippets{}.AttrTypes(), &item)
		unpacked, d := unpackTrustInfoWithSharedSnippetsToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.TrustInfoWithSharedSnippets", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for TrustInfoWithSharedSnippets ---
func packTrustInfoWithSharedSnippetsListFromSdk(ctx context.Context, sdks []config_setup.TrustInfoWithSharedSnippets) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.TrustInfoWithSharedSnippets")
	diags := diag.Diagnostics{}
	var data []models.TrustInfoWithSharedSnippets

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.TrustInfoWithSharedSnippets
		obj, d := packTrustInfoWithSharedSnippetsFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.TrustInfoWithSharedSnippets{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.TrustInfoWithSharedSnippets", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.TrustInfoWithSharedSnippets{}.AttrType(), data)
}
