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

// --- Unpacker for SnippetShareInfo ---
func unpackSnippetShareInfoToSdk(ctx context.Context, obj types.Object) (*config_setup.SnippetShareInfo, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.SnippetShareInfo", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.SnippetShareInfo
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk config_setup.SnippetShareInfo
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

	// Handling Lists
	if !model.Properties.IsNull() && !model.Properties.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field Properties")
		unpacked, d := unpackSnippetSharePropertyListToSdk(ctx, model.Properties)
		diags.Append(d...)
		sdk.Properties = unpacked
	}

	// Handling Primitives
	if !model.RecipientPausedUpdate.IsNull() && !model.RecipientPausedUpdate.IsUnknown() {
		sdk.RecipientPausedUpdate = model.RecipientPausedUpdate.ValueBoolPointer()
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
	if !model.RecipientTsg.IsNull() && !model.RecipientTsg.IsUnknown() {
		sdk.RecipientTsg = model.RecipientTsg.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "RecipientTsg", "value": *sdk.RecipientTsg})
	}

	// Handling Primitives
	if !model.RecipientValidateBeforeUpdate.IsNull() && !model.RecipientValidateBeforeUpdate.IsUnknown() {
		sdk.RecipientValidateBeforeUpdate = model.RecipientValidateBeforeUpdate.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "RecipientValidateBeforeUpdate", "value": *sdk.RecipientValidateBeforeUpdate})
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

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.SnippetShareInfo", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for SnippetShareInfo ---
func packSnippetShareInfoFromSdk(ctx context.Context, sdk config_setup.SnippetShareInfo) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.SnippetShareInfo", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.SnippetShareInfo
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
	// Handling Lists
	if sdk.Properties != nil {
		tflog.Debug(ctx, "Packing list of objects for field Properties")
		packed, d := packSnippetSharePropertyListFromSdk(ctx, sdk.Properties)
		diags.Append(d...)
		model.Properties = packed
	} else {
		model.Properties = basetypes.NewListNull(models.SnippetShareProperty{}.AttrType())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.RecipientPausedUpdate != nil {
		model.RecipientPausedUpdate = basetypes.NewBoolValue(*sdk.RecipientPausedUpdate)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "RecipientPausedUpdate", "value": *sdk.RecipientPausedUpdate})
	} else {
		model.RecipientPausedUpdate = basetypes.NewBoolNull()
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
	if sdk.RecipientTsg != nil {
		model.RecipientTsg = basetypes.NewStringValue(*sdk.RecipientTsg)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "RecipientTsg", "value": *sdk.RecipientTsg})
	} else {
		model.RecipientTsg = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.RecipientValidateBeforeUpdate != nil {
		model.RecipientValidateBeforeUpdate = basetypes.NewBoolValue(*sdk.RecipientValidateBeforeUpdate)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "RecipientValidateBeforeUpdate", "value": *sdk.RecipientValidateBeforeUpdate})
	} else {
		model.RecipientValidateBeforeUpdate = basetypes.NewBoolNull()
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
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.SnippetShareInfo{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.SnippetShareInfo", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for SnippetShareInfo ---
func unpackSnippetShareInfoListToSdk(ctx context.Context, list types.List) ([]config_setup.SnippetShareInfo, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.SnippetShareInfo")
	diags := diag.Diagnostics{}
	var data []models.SnippetShareInfo
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]config_setup.SnippetShareInfo, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.SnippetShareInfo{}.AttrTypes(), &item)
		unpacked, d := unpackSnippetShareInfoToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.SnippetShareInfo", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for SnippetShareInfo ---
func packSnippetShareInfoListFromSdk(ctx context.Context, sdks []config_setup.SnippetShareInfo) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.SnippetShareInfo")
	diags := diag.Diagnostics{}
	var data []models.SnippetShareInfo

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.SnippetShareInfo
		obj, d := packSnippetShareInfoFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.SnippetShareInfo{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.SnippetShareInfo", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.SnippetShareInfo{}.AttrType(), data)
}

// --- Unpacker for SnippetShareProperty ---
func unpackSnippetSharePropertyToSdk(ctx context.Context, obj types.Object) (*config_setup.SnippetShareProperty, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.SnippetShareProperty", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.SnippetShareProperty
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk config_setup.SnippetShareProperty
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
	if !model.DonorTenant.IsNull() && !model.DonorTenant.IsUnknown() {
		sdk.DonorTenant = model.DonorTenant.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DonorTenant", "value": *sdk.DonorTenant})
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
	if !model.MsgUuid.IsNull() && !model.MsgUuid.IsUnknown() {
		sdk.MsgUuid = model.MsgUuid.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "MsgUuid", "value": *sdk.MsgUuid})
	}

	// Handling Primitives
	if !model.PropertyName.IsNull() && !model.PropertyName.IsUnknown() {
		sdk.PropertyName = model.PropertyName.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "PropertyName", "value": *sdk.PropertyName})
	}

	// Handling Primitives
	if !model.PropertyValue.IsNull() && !model.PropertyValue.IsUnknown() {
		sdk.PropertyValue = model.PropertyValue.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "PropertyValue", "value": *sdk.PropertyValue})
	}

	// Handling Primitives
	if !model.RecipientTenant.IsNull() && !model.RecipientTenant.IsUnknown() {
		sdk.RecipientTenant = model.RecipientTenant.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "RecipientTenant", "value": *sdk.RecipientTenant})
	}

	// Handling Primitives
	if !model.RecipientTsg.IsNull() && !model.RecipientTsg.IsUnknown() {
		sdk.RecipientTsg = model.RecipientTsg.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "RecipientTsg", "value": *sdk.RecipientTsg})
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
	if !model.Updated.IsNull() && !model.Updated.IsUnknown() {
		if t, parseErr := time.Parse(time.RFC3339, model.Updated.ValueString()); parseErr == nil {
			sdk.Updated = &t
		}
	}

	// Handling Primitives
	if !model.UpdatedBy.IsNull() && !model.UpdatedBy.IsUnknown() {
		sdk.UpdatedBy = model.UpdatedBy.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "UpdatedBy", "value": *sdk.UpdatedBy})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.SnippetShareProperty", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for SnippetShareProperty ---
func packSnippetSharePropertyFromSdk(ctx context.Context, sdk config_setup.SnippetShareProperty) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.SnippetShareProperty", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.SnippetShareProperty
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
	if sdk.DonorTenant != nil {
		model.DonorTenant = basetypes.NewStringValue(*sdk.DonorTenant)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DonorTenant", "value": *sdk.DonorTenant})
	} else {
		model.DonorTenant = basetypes.NewStringNull()
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
	if sdk.MsgUuid != nil {
		model.MsgUuid = basetypes.NewStringValue(*sdk.MsgUuid)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "MsgUuid", "value": *sdk.MsgUuid})
	} else {
		model.MsgUuid = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.PropertyName != nil {
		model.PropertyName = basetypes.NewStringValue(*sdk.PropertyName)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "PropertyName", "value": *sdk.PropertyName})
	} else {
		model.PropertyName = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.PropertyValue != nil {
		model.PropertyValue = basetypes.NewStringValue(*sdk.PropertyValue)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "PropertyValue", "value": *sdk.PropertyValue})
	} else {
		model.PropertyValue = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.RecipientTenant != nil {
		model.RecipientTenant = basetypes.NewStringValue(*sdk.RecipientTenant)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "RecipientTenant", "value": *sdk.RecipientTenant})
	} else {
		model.RecipientTenant = basetypes.NewStringNull()
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
	if sdk.Updated != nil {
		model.Updated = basetypes.NewStringValue(sdk.Updated.Format(time.RFC3339))
		tflog.Debug(ctx, "Packed time pointer", map[string]interface{}{"field": "Updated", "value": sdk.Updated.Format(time.RFC3339)})
	} else {
		model.Updated = basetypes.NewStringNull()
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

	obj, d := types.ObjectValueFrom(ctx, models.SnippetShareProperty{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.SnippetShareProperty", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for SnippetShareProperty ---
func unpackSnippetSharePropertyListToSdk(ctx context.Context, list types.List) ([]config_setup.SnippetShareProperty, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.SnippetShareProperty")
	diags := diag.Diagnostics{}
	var data []models.SnippetShareProperty
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]config_setup.SnippetShareProperty, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.SnippetShareProperty{}.AttrTypes(), &item)
		unpacked, d := unpackSnippetSharePropertyToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.SnippetShareProperty", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for SnippetShareProperty ---
func packSnippetSharePropertyListFromSdk(ctx context.Context, sdks []config_setup.SnippetShareProperty) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.SnippetShareProperty")
	diags := diag.Diagnostics{}
	var data []models.SnippetShareProperty

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.SnippetShareProperty
		obj, d := packSnippetSharePropertyFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.SnippetShareProperty{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.SnippetShareProperty", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.SnippetShareProperty{}.AttrType(), data)
}
