package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/objects"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/objects"
)

// --- Unpacker for LogForwardingProfiles ---
func unpackLogForwardingProfilesToSdk(ctx context.Context, obj types.Object) (*objects.LogForwardingProfiles, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.LogForwardingProfiles", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.LogForwardingProfiles
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.LogForwardingProfiles
	var d diag.Diagnostics

	// Handling Primitives
	if !model.Description.IsNull() && !model.Description.IsUnknown() {
		sdk.Description = model.Description.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Description", "value": *sdk.Description})
	}

	// Handling Primitives
	if !model.Device.IsNull() && !model.Device.IsUnknown() {
		sdk.Device = model.Device.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Device", "value": *sdk.Device})
	}

	// Handling Primitives
	if !model.Folder.IsNull() && !model.Folder.IsUnknown() {
		sdk.Folder = model.Folder.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Folder", "value": *sdk.Folder})
	}

	// Handling Primitives
	if !model.Id.IsNull() && !model.Id.IsUnknown() {
		sdk.Id = model.Id.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Id", "value": *sdk.Id})
	}

	// Handling Lists
	if !model.MatchList.IsNull() && !model.MatchList.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field MatchList")
		unpacked, d := unpackLogForwardingProfilesMatchListInnerListToSdk(ctx, model.MatchList)
		diags.Append(d...)
		sdk.MatchList = unpacked
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	// Handling Primitives
	if !model.Snippet.IsNull() && !model.Snippet.IsUnknown() {
		sdk.Snippet = model.Snippet.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.LogForwardingProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for LogForwardingProfiles ---
func packLogForwardingProfilesFromSdk(ctx context.Context, sdk objects.LogForwardingProfiles) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.LogForwardingProfiles", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.LogForwardingProfiles
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
	if sdk.Device != nil {
		model.Device = basetypes.NewStringValue(*sdk.Device)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Device", "value": *sdk.Device})
	} else {
		model.Device = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Folder != nil {
		model.Folder = basetypes.NewStringValue(*sdk.Folder)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Folder", "value": *sdk.Folder})
	} else {
		model.Folder = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Id != nil {
		model.Id = basetypes.NewStringValue(*sdk.Id)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Id", "value": *sdk.Id})
	} else {
		model.Id = basetypes.NewStringNull()
	}
	// Handling Lists
	if sdk.MatchList != nil {
		tflog.Debug(ctx, "Packing list of objects for field MatchList")
		packed, d := packLogForwardingProfilesMatchListInnerListFromSdk(ctx, sdk.MatchList)
		diags.Append(d...)
		model.MatchList = packed
	} else {
		model.MatchList = basetypes.NewListNull(models.LogForwardingProfilesMatchListInner{}.AttrType())
	}
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	// Handling Primitives
	// Standard primitive packing
	if sdk.Snippet != nil {
		model.Snippet = basetypes.NewStringValue(*sdk.Snippet)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	} else {
		model.Snippet = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.LogForwardingProfiles{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.LogForwardingProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for LogForwardingProfiles ---
func unpackLogForwardingProfilesListToSdk(ctx context.Context, list types.List) ([]objects.LogForwardingProfiles, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.LogForwardingProfiles")
	diags := diag.Diagnostics{}
	var data []models.LogForwardingProfiles
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.LogForwardingProfiles, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.LogForwardingProfiles{}.AttrTypes(), &item)
		unpacked, d := unpackLogForwardingProfilesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.LogForwardingProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for LogForwardingProfiles ---
func packLogForwardingProfilesListFromSdk(ctx context.Context, sdks []objects.LogForwardingProfiles) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.LogForwardingProfiles")
	diags := diag.Diagnostics{}
	var data []models.LogForwardingProfiles

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.LogForwardingProfiles
		obj, d := packLogForwardingProfilesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.LogForwardingProfiles{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.LogForwardingProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.LogForwardingProfiles{}.AttrType(), data)
}

// --- Unpacker for LogForwardingProfilesMatchListInner ---
func unpackLogForwardingProfilesMatchListInnerToSdk(ctx context.Context, obj types.Object) (*objects.LogForwardingProfilesMatchListInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.LogForwardingProfilesMatchListInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.LogForwardingProfilesMatchListInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.LogForwardingProfilesMatchListInner
	var d diag.Diagnostics
	// Handling Primitives
	if !model.ActionDesc.IsNull() && !model.ActionDesc.IsUnknown() {
		sdk.ActionDesc = model.ActionDesc.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "ActionDesc", "value": *sdk.ActionDesc})
	}

	// Handling Primitives
	if !model.Filter.IsNull() && !model.Filter.IsUnknown() {
		sdk.Filter = model.Filter.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Filter", "value": sdk.Filter})
	}

	// Handling Primitives
	if !model.LogType.IsNull() && !model.LogType.IsUnknown() {
		sdk.LogType = model.LogType.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "LogType", "value": sdk.LogType})
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	// Handling Lists
	if !model.SendEmail.IsNull() && !model.SendEmail.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field SendEmail")
		diags.Append(model.SendEmail.ElementsAs(ctx, &sdk.SendEmail, false)...)
	}

	// Handling Lists
	if !model.SendHttp.IsNull() && !model.SendHttp.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field SendHttp")
		diags.Append(model.SendHttp.ElementsAs(ctx, &sdk.SendHttp, false)...)
	}

	// Handling Lists
	if !model.SendSnmptrap.IsNull() && !model.SendSnmptrap.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field SendSnmptrap")
		diags.Append(model.SendSnmptrap.ElementsAs(ctx, &sdk.SendSnmptrap, false)...)
	}

	// Handling Lists
	if !model.SendSyslog.IsNull() && !model.SendSyslog.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field SendSyslog")
		diags.Append(model.SendSyslog.ElementsAs(ctx, &sdk.SendSyslog, false)...)
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.LogForwardingProfilesMatchListInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for LogForwardingProfilesMatchListInner ---
func packLogForwardingProfilesMatchListInnerFromSdk(ctx context.Context, sdk objects.LogForwardingProfilesMatchListInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.LogForwardingProfilesMatchListInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.LogForwardingProfilesMatchListInner
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.ActionDesc != nil {
		model.ActionDesc = basetypes.NewStringValue(*sdk.ActionDesc)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "ActionDesc", "value": *sdk.ActionDesc})
	} else {
		model.ActionDesc = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	model.Filter = basetypes.NewStringValue(sdk.Filter)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Filter", "value": sdk.Filter})
	// Handling Primitives
	// Standard primitive packing
	model.LogType = basetypes.NewStringValue(sdk.LogType)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "LogType", "value": sdk.LogType})
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	// Handling Lists
	if sdk.SendEmail != nil {
		tflog.Debug(ctx, "Packing list of primitives for field SendEmail")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.SendEmail, d = basetypes.NewListValueFrom(ctx, elemType, sdk.SendEmail)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.SendEmail = basetypes.NewListNull(elemType)
	}
	// Handling Lists
	if sdk.SendHttp != nil {
		tflog.Debug(ctx, "Packing list of primitives for field SendHttp")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.SendHttp, d = basetypes.NewListValueFrom(ctx, elemType, sdk.SendHttp)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.SendHttp = basetypes.NewListNull(elemType)
	}
	// Handling Lists
	if sdk.SendSnmptrap != nil {
		tflog.Debug(ctx, "Packing list of primitives for field SendSnmptrap")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.SendSnmptrap, d = basetypes.NewListValueFrom(ctx, elemType, sdk.SendSnmptrap)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.SendSnmptrap = basetypes.NewListNull(elemType)
	}
	// Handling Lists
	if sdk.SendSyslog != nil {
		tflog.Debug(ctx, "Packing list of primitives for field SendSyslog")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.SendSyslog, d = basetypes.NewListValueFrom(ctx, elemType, sdk.SendSyslog)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.SendSyslog = basetypes.NewListNull(elemType)
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.LogForwardingProfilesMatchListInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.LogForwardingProfilesMatchListInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for LogForwardingProfilesMatchListInner ---
func unpackLogForwardingProfilesMatchListInnerListToSdk(ctx context.Context, list types.List) ([]objects.LogForwardingProfilesMatchListInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.LogForwardingProfilesMatchListInner")
	diags := diag.Diagnostics{}
	var data []models.LogForwardingProfilesMatchListInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.LogForwardingProfilesMatchListInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.LogForwardingProfilesMatchListInner{}.AttrTypes(), &item)
		unpacked, d := unpackLogForwardingProfilesMatchListInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.LogForwardingProfilesMatchListInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for LogForwardingProfilesMatchListInner ---
func packLogForwardingProfilesMatchListInnerListFromSdk(ctx context.Context, sdks []objects.LogForwardingProfilesMatchListInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.LogForwardingProfilesMatchListInner")
	diags := diag.Diagnostics{}
	var data []models.LogForwardingProfilesMatchListInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.LogForwardingProfilesMatchListInner
		obj, d := packLogForwardingProfilesMatchListInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.LogForwardingProfilesMatchListInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.LogForwardingProfilesMatchListInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.LogForwardingProfilesMatchListInner{}.AttrType(), data)
}
