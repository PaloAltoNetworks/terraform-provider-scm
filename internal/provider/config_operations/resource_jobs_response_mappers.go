package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/config_operations"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/config_operations"
)

// --- Unpacker for JobsResponse ---
func unpackJobsResponseToSdk(ctx context.Context, obj types.Object) (*config_operations.JobsResponse, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.JobsResponse", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.JobsResponse
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk config_operations.JobsResponse
	var d diag.Diagnostics

	// Handling Lists
	if !model.Data.IsNull() && !model.Data.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field Data")
		unpacked, d := unpackJobsListToSdk(ctx, model.Data)
		diags.Append(d...)
		sdk.Data = unpacked
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.JobsResponse", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for JobsResponse ---
func packJobsResponseFromSdk(ctx context.Context, sdk config_operations.JobsResponse) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.JobsResponse", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.JobsResponse
	var d diag.Diagnostics
	// Handling Lists
	if sdk.Data != nil {
		tflog.Debug(ctx, "Packing list of objects for field Data")
		packed, d := packJobsListFromSdk(ctx, sdk.Data)
		diags.Append(d...)
		model.Data = packed
	} else {
		model.Data = basetypes.NewListNull(models.Jobs{}.AttrType())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.JobsResponse{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.JobsResponse", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for JobsResponse ---
func unpackJobsResponseListToSdk(ctx context.Context, list types.List) ([]config_operations.JobsResponse, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.JobsResponse")
	diags := diag.Diagnostics{}
	var data []models.JobsResponse
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]config_operations.JobsResponse, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.JobsResponse{}.AttrTypes(), &item)
		unpacked, d := unpackJobsResponseToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.JobsResponse", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for JobsResponse ---
func packJobsResponseListFromSdk(ctx context.Context, sdks []config_operations.JobsResponse) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.JobsResponse")
	diags := diag.Diagnostics{}
	var data []models.JobsResponse

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.JobsResponse
		obj, d := packJobsResponseFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.JobsResponse{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.JobsResponse", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.JobsResponse{}.AttrType(), data)
}

// --- Unpacker for Jobs ---
func unpackJobsToSdk(ctx context.Context, obj types.Object) (*config_operations.Jobs, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.Jobs", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.Jobs
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk config_operations.Jobs
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Description.IsNull() && !model.Description.IsUnknown() {
		sdk.Description = model.Description.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Description", "value": *sdk.Description})
	}

	// Handling Primitives
	if !model.Details.IsNull() && !model.Details.IsUnknown() {
		sdk.Details = model.Details.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Details", "value": *sdk.Details})
	}

	// Handling Primitives
	if !model.DeviceName.IsNull() && !model.DeviceName.IsUnknown() {
		sdk.DeviceName = model.DeviceName.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "DeviceName", "value": sdk.DeviceName})
	}

	// Handling Primitives
	if !model.EndTs.IsNull() && !model.EndTs.IsUnknown() {
		sdk.EndTs = model.EndTs.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "EndTs", "value": sdk.EndTs})
	}

	// Handling Primitives
	if !model.Id.IsNull() && !model.Id.IsUnknown() {
		sdk.Id = model.Id.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Id", "value": sdk.Id})
	}

	// Handling Primitives
	if !model.JobResult.IsNull() && !model.JobResult.IsUnknown() {
		sdk.JobResult = model.JobResult.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "JobResult", "value": sdk.JobResult})
	}

	// Handling Primitives
	if !model.JobStatus.IsNull() && !model.JobStatus.IsUnknown() {
		sdk.JobStatus = model.JobStatus.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "JobStatus", "value": sdk.JobStatus})
	}

	// Handling Primitives
	if !model.JobType.IsNull() && !model.JobType.IsUnknown() {
		sdk.JobType = model.JobType.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "JobType", "value": sdk.JobType})
	}

	// Handling Primitives
	if !model.ParentId.IsNull() && !model.ParentId.IsUnknown() {
		sdk.ParentId = model.ParentId.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "ParentId", "value": sdk.ParentId})
	}

	// Handling Primitives
	if !model.Percent.IsNull() && !model.Percent.IsUnknown() {
		sdk.Percent = model.Percent.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Percent", "value": sdk.Percent})
	}

	// Handling Primitives
	if !model.ResultStr.IsNull() && !model.ResultStr.IsUnknown() {
		sdk.ResultStr = model.ResultStr.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "ResultStr", "value": sdk.ResultStr})
	}

	// Handling Primitives
	if !model.StartTs.IsNull() && !model.StartTs.IsUnknown() {
		sdk.StartTs = model.StartTs.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "StartTs", "value": sdk.StartTs})
	}

	// Handling Primitives
	if !model.StatusStr.IsNull() && !model.StatusStr.IsUnknown() {
		sdk.StatusStr = model.StatusStr.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "StatusStr", "value": sdk.StatusStr})
	}

	// Handling Primitives
	if !model.Summary.IsNull() && !model.Summary.IsUnknown() {
		sdk.Summary = model.Summary.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Summary", "value": sdk.Summary})
	}

	// Handling Primitives
	if !model.TypeStr.IsNull() && !model.TypeStr.IsUnknown() {
		sdk.TypeStr = model.TypeStr.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "TypeStr", "value": sdk.TypeStr})
	}

	// Handling Primitives
	if !model.Uname.IsNull() && !model.Uname.IsUnknown() {
		sdk.Uname = model.Uname.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Uname", "value": sdk.Uname})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.Jobs", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for Jobs ---
func packJobsFromSdk(ctx context.Context, sdk config_operations.Jobs) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.Jobs", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.Jobs
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Description != nil {
		model.Description = basetypes.NewStringValue(*sdk.Description)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Description", "value": *sdk.Description})
	} else {
		model.Description = basetypes.NewStringNull()
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
	model.DeviceName = basetypes.NewStringValue(sdk.DeviceName)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "DeviceName", "value": sdk.DeviceName})
	// Handling Primitives
	// Standard primitive packing
	model.EndTs = basetypes.NewStringValue(sdk.EndTs)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "EndTs", "value": sdk.EndTs})
	// Handling Primitives
	// Standard primitive packing
	model.Id = basetypes.NewStringValue(sdk.Id)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Id", "value": sdk.Id})
	// Handling Primitives
	// Standard primitive packing
	model.JobResult = basetypes.NewStringValue(sdk.JobResult)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "JobResult", "value": sdk.JobResult})
	// Handling Primitives
	// Standard primitive packing
	model.JobStatus = basetypes.NewStringValue(sdk.JobStatus)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "JobStatus", "value": sdk.JobStatus})
	// Handling Primitives
	// Standard primitive packing
	model.JobType = basetypes.NewStringValue(sdk.JobType)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "JobType", "value": sdk.JobType})
	// Handling Primitives
	// Standard primitive packing
	model.ParentId = basetypes.NewStringValue(sdk.ParentId)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "ParentId", "value": sdk.ParentId})
	// Handling Primitives
	// Standard primitive packing
	model.Percent = basetypes.NewStringValue(sdk.Percent)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Percent", "value": sdk.Percent})
	// Handling Primitives
	// Standard primitive packing
	model.ResultStr = basetypes.NewStringValue(sdk.ResultStr)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "ResultStr", "value": sdk.ResultStr})
	// Handling Primitives
	// Standard primitive packing
	model.StartTs = basetypes.NewStringValue(sdk.StartTs)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "StartTs", "value": sdk.StartTs})
	// Handling Primitives
	// Standard primitive packing
	model.StatusStr = basetypes.NewStringValue(sdk.StatusStr)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "StatusStr", "value": sdk.StatusStr})
	// Handling Primitives
	// Standard primitive packing
	model.Summary = basetypes.NewStringValue(sdk.Summary)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Summary", "value": sdk.Summary})
	// Handling Primitives
	// Standard primitive packing
	model.TypeStr = basetypes.NewStringValue(sdk.TypeStr)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "TypeStr", "value": sdk.TypeStr})
	// Handling Primitives
	// Standard primitive packing
	model.Uname = basetypes.NewStringValue(sdk.Uname)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Uname", "value": sdk.Uname})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.Jobs{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.Jobs", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for Jobs ---
func unpackJobsListToSdk(ctx context.Context, list types.List) ([]config_operations.Jobs, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.Jobs")
	diags := diag.Diagnostics{}
	var data []models.Jobs
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]config_operations.Jobs, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.Jobs{}.AttrTypes(), &item)
		unpacked, d := unpackJobsToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.Jobs", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for Jobs ---
func packJobsListFromSdk(ctx context.Context, sdks []config_operations.Jobs) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.Jobs")
	diags := diag.Diagnostics{}
	var data []models.Jobs

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.Jobs
		obj, d := packJobsFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.Jobs{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.Jobs", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.Jobs{}.AttrType(), data)
}
