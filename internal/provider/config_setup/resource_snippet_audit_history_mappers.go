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

// --- Unpacker for SnippetAuditHistory ---
func unpackSnippetAuditHistoryToSdk(ctx context.Context, obj types.Object) (*config_setup.SnippetAuditHistory, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.SnippetAuditHistory", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.SnippetAuditHistory
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk config_setup.SnippetAuditHistory
	var d diag.Diagnostics

	// Handling Primitives
	if !model.Action.IsNull() && !model.Action.IsUnknown() {
		sdk.Action = model.Action.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Action", "value": *sdk.Action})
	}

	// Handling Primitives
	if !model.Created.IsNull() && !model.Created.IsUnknown() {
		if t, parseErr := time.Parse(time.RFC3339, model.Created.ValueString()); parseErr == nil {
			sdk.Created = &t
		}
	}

	// Handling Primitives
	if !model.Deleted.IsNull() && !model.Deleted.IsUnknown() {
		val := int32(model.Deleted.ValueInt64())
		sdk.Deleted = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Deleted", "value": *sdk.Deleted})
	}

	// Handling Primitives
	if !model.Details.IsNull() && !model.Details.IsUnknown() {
		sdk.Details = model.Details.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Details", "value": *sdk.Details})
	}

	// Handling Primitives
	if !model.Display.IsNull() && !model.Display.IsUnknown() {
		val := int32(model.Display.ValueInt64())
		sdk.Display = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Display", "value": *sdk.Display})
	}

	// Handling Primitives
	if !model.DonorCreated.IsNull() && !model.DonorCreated.IsUnknown() {
		val := int32(model.DonorCreated.ValueInt64())
		sdk.DonorCreated = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DonorCreated", "value": *sdk.DonorCreated})
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
	if !model.Id.IsNull() && !model.Id.IsUnknown() {
		val := int32(model.Id.ValueInt64())
		sdk.Id = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Id", "value": *sdk.Id})
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
	if !model.SnippetUuid.IsNull() && !model.SnippetUuid.IsUnknown() {
		sdk.SnippetUuid = model.SnippetUuid.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "SnippetUuid", "value": *sdk.SnippetUuid})
	}

	// Handling Primitives
	if !model.User.IsNull() && !model.User.IsUnknown() {
		sdk.User = model.User.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "User", "value": *sdk.User})
	}

	// Handling Primitives
	if !model.Version.IsNull() && !model.Version.IsUnknown() {
		sdk.Version = model.Version.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Version", "value": *sdk.Version})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.SnippetAuditHistory", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for SnippetAuditHistory ---
func packSnippetAuditHistoryFromSdk(ctx context.Context, sdk config_setup.SnippetAuditHistory) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.SnippetAuditHistory", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.SnippetAuditHistory
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Action != nil {
		model.Action = basetypes.NewStringValue(*sdk.Action)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Action", "value": *sdk.Action})
	} else {
		model.Action = basetypes.NewStringNull()
	}
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
	if sdk.Deleted != nil {
		model.Deleted = basetypes.NewInt64Value(int64(*sdk.Deleted))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Deleted", "value": *sdk.Deleted})
	} else {
		model.Deleted = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Details != nil {
		model.Details = basetypes.NewStringValue(*sdk.Details)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Details", "value": *sdk.Details})
	} else {
		model.Details = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Display != nil {
		model.Display = basetypes.NewInt64Value(int64(*sdk.Display))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Display", "value": *sdk.Display})
	} else {
		model.Display = basetypes.NewInt64Null()
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
	if sdk.Id != nil {
		model.Id = basetypes.NewInt64Value(int64(*sdk.Id))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Id", "value": *sdk.Id})
	} else {
		model.Id = basetypes.NewInt64Null()
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
	if sdk.SnippetUuid != nil {
		model.SnippetUuid = basetypes.NewStringValue(*sdk.SnippetUuid)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "SnippetUuid", "value": *sdk.SnippetUuid})
	} else {
		model.SnippetUuid = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.User != nil {
		model.User = basetypes.NewStringValue(*sdk.User)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "User", "value": *sdk.User})
	} else {
		model.User = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Version != nil {
		model.Version = basetypes.NewStringValue(*sdk.Version)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Version", "value": *sdk.Version})
	} else {
		model.Version = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.SnippetAuditHistory{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.SnippetAuditHistory", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for SnippetAuditHistory ---
func unpackSnippetAuditHistoryListToSdk(ctx context.Context, list types.List) ([]config_setup.SnippetAuditHistory, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.SnippetAuditHistory")
	diags := diag.Diagnostics{}
	var data []models.SnippetAuditHistory
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]config_setup.SnippetAuditHistory, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.SnippetAuditHistory{}.AttrTypes(), &item)
		unpacked, d := unpackSnippetAuditHistoryToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.SnippetAuditHistory", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for SnippetAuditHistory ---
func packSnippetAuditHistoryListFromSdk(ctx context.Context, sdks []config_setup.SnippetAuditHistory) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.SnippetAuditHistory")
	diags := diag.Diagnostics{}
	var data []models.SnippetAuditHistory

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.SnippetAuditHistory
		obj, d := packSnippetAuditHistoryFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.SnippetAuditHistory{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.SnippetAuditHistory", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.SnippetAuditHistory{}.AttrType(), data)
}
