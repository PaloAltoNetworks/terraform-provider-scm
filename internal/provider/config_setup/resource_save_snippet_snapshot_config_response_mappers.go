package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/config_setup"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/config_setup"
)

// --- Unpacker for SaveSnippetSnapshotConfigResponse ---
func unpackSaveSnippetSnapshotConfigResponseToSdk(ctx context.Context, obj types.Object) (*config_setup.SaveSnippetSnapshotConfigResponse, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.SaveSnippetSnapshotConfigResponse", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.SaveSnippetSnapshotConfigResponse
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk config_setup.SaveSnippetSnapshotConfigResponse
	var d diag.Diagnostics

	// Handling Objects
	if !model.Result.IsNull() && !model.Result.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Result")
		unpacked, d := unpackSaveSnippetSnapshotConfigResponseResultToSdk(ctx, model.Result)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Result"})
		}
		if unpacked != nil {
			sdk.Result = unpacked
		}
	}

	// Handling Primitives
	if !model.Status.IsNull() && !model.Status.IsUnknown() {
		sdk.Status = model.Status.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Status", "value": *sdk.Status})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.SaveSnippetSnapshotConfigResponse", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for SaveSnippetSnapshotConfigResponse ---
func packSaveSnippetSnapshotConfigResponseFromSdk(ctx context.Context, sdk config_setup.SaveSnippetSnapshotConfigResponse) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.SaveSnippetSnapshotConfigResponse", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.SaveSnippetSnapshotConfigResponse
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Result != nil {
		tflog.Debug(ctx, "Packing nested object for field Result")
		packed, d := packSaveSnippetSnapshotConfigResponseResultFromSdk(ctx, *sdk.Result)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Result"})
		}
		model.Result = packed
	} else {
		model.Result = basetypes.NewObjectNull(models.SaveSnippetSnapshotConfigResponseResult{}.AttrTypes())
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

	obj, d := types.ObjectValueFrom(ctx, models.SaveSnippetSnapshotConfigResponse{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.SaveSnippetSnapshotConfigResponse", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for SaveSnippetSnapshotConfigResponse ---
func unpackSaveSnippetSnapshotConfigResponseListToSdk(ctx context.Context, list types.List) ([]config_setup.SaveSnippetSnapshotConfigResponse, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.SaveSnippetSnapshotConfigResponse")
	diags := diag.Diagnostics{}
	var data []models.SaveSnippetSnapshotConfigResponse
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]config_setup.SaveSnippetSnapshotConfigResponse, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.SaveSnippetSnapshotConfigResponse{}.AttrTypes(), &item)
		unpacked, d := unpackSaveSnippetSnapshotConfigResponseToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.SaveSnippetSnapshotConfigResponse", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for SaveSnippetSnapshotConfigResponse ---
func packSaveSnippetSnapshotConfigResponseListFromSdk(ctx context.Context, sdks []config_setup.SaveSnippetSnapshotConfigResponse) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.SaveSnippetSnapshotConfigResponse")
	diags := diag.Diagnostics{}
	var data []models.SaveSnippetSnapshotConfigResponse

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.SaveSnippetSnapshotConfigResponse
		obj, d := packSaveSnippetSnapshotConfigResponseFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.SaveSnippetSnapshotConfigResponse{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.SaveSnippetSnapshotConfigResponse", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.SaveSnippetSnapshotConfigResponse{}.AttrType(), data)
}

// --- Unpacker for SaveSnippetSnapshotConfigResponseResult ---
func unpackSaveSnippetSnapshotConfigResponseResultToSdk(ctx context.Context, obj types.Object) (*config_setup.SaveSnippetSnapshotConfigResponseResult, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.SaveSnippetSnapshotConfigResponseResult", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.SaveSnippetSnapshotConfigResponseResult
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk config_setup.SaveSnippetSnapshotConfigResponseResult
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Version.IsNull() && !model.Version.IsUnknown() {
		sdk.Version = model.Version.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Version", "value": *sdk.Version})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.SaveSnippetSnapshotConfigResponseResult", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for SaveSnippetSnapshotConfigResponseResult ---
func packSaveSnippetSnapshotConfigResponseResultFromSdk(ctx context.Context, sdk config_setup.SaveSnippetSnapshotConfigResponseResult) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.SaveSnippetSnapshotConfigResponseResult", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.SaveSnippetSnapshotConfigResponseResult
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Version != nil {
		model.Version = basetypes.NewStringValue(*sdk.Version)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Version", "value": *sdk.Version})
	} else {
		model.Version = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.SaveSnippetSnapshotConfigResponseResult{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.SaveSnippetSnapshotConfigResponseResult", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for SaveSnippetSnapshotConfigResponseResult ---
func unpackSaveSnippetSnapshotConfigResponseResultListToSdk(ctx context.Context, list types.List) ([]config_setup.SaveSnippetSnapshotConfigResponseResult, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.SaveSnippetSnapshotConfigResponseResult")
	diags := diag.Diagnostics{}
	var data []models.SaveSnippetSnapshotConfigResponseResult
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]config_setup.SaveSnippetSnapshotConfigResponseResult, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.SaveSnippetSnapshotConfigResponseResult{}.AttrTypes(), &item)
		unpacked, d := unpackSaveSnippetSnapshotConfigResponseResultToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.SaveSnippetSnapshotConfigResponseResult", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for SaveSnippetSnapshotConfigResponseResult ---
func packSaveSnippetSnapshotConfigResponseResultListFromSdk(ctx context.Context, sdks []config_setup.SaveSnippetSnapshotConfigResponseResult) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.SaveSnippetSnapshotConfigResponseResult")
	diags := diag.Diagnostics{}
	var data []models.SaveSnippetSnapshotConfigResponseResult

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.SaveSnippetSnapshotConfigResponseResult
		obj, d := packSaveSnippetSnapshotConfigResponseResultFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.SaveSnippetSnapshotConfigResponseResult{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.SaveSnippetSnapshotConfigResponseResult", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.SaveSnippetSnapshotConfigResponseResult{}.AttrType(), data)
}
