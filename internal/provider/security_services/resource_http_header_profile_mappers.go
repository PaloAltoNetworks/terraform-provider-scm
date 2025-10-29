package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/security_services"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/security_services"
)

// --- Unpacker for HttpHeaderProfiles ---
func unpackHttpHeaderProfilesToSdk(ctx context.Context, obj types.Object) (*security_services.HttpHeaderProfiles, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.HttpHeaderProfiles", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.HttpHeaderProfiles
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.HttpHeaderProfiles
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

	// Handling Lists
	if !model.HttpHeaderInsertion.IsNull() && !model.HttpHeaderInsertion.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field HttpHeaderInsertion")
		unpacked, d := unpackHttpHeaderProfilesHttpHeaderInsertionInnerListToSdk(ctx, model.HttpHeaderInsertion)
		diags.Append(d...)
		sdk.HttpHeaderInsertion = unpacked
	}

	// Handling Primitives
	if !model.Id.IsNull() && !model.Id.IsUnknown() {
		sdk.Id = model.Id.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Id", "value": *sdk.Id})
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

	tflog.Debug(ctx, "Exiting unpack helper for models.HttpHeaderProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for HttpHeaderProfiles ---
func packHttpHeaderProfilesFromSdk(ctx context.Context, sdk security_services.HttpHeaderProfiles) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.HttpHeaderProfiles", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.HttpHeaderProfiles
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
	// Handling Lists
	if sdk.HttpHeaderInsertion != nil {
		tflog.Debug(ctx, "Packing list of objects for field HttpHeaderInsertion")
		packed, d := packHttpHeaderProfilesHttpHeaderInsertionInnerListFromSdk(ctx, sdk.HttpHeaderInsertion)
		diags.Append(d...)
		model.HttpHeaderInsertion = packed
	} else {
		model.HttpHeaderInsertion = basetypes.NewListNull(models.HttpHeaderProfilesHttpHeaderInsertionInner{}.AttrType())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Id != nil {
		model.Id = basetypes.NewStringValue(*sdk.Id)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Id", "value": *sdk.Id})
	} else {
		model.Id = basetypes.NewStringNull()
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

	obj, d := types.ObjectValueFrom(ctx, models.HttpHeaderProfiles{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.HttpHeaderProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for HttpHeaderProfiles ---
func unpackHttpHeaderProfilesListToSdk(ctx context.Context, list types.List) ([]security_services.HttpHeaderProfiles, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.HttpHeaderProfiles")
	diags := diag.Diagnostics{}
	var data []models.HttpHeaderProfiles
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.HttpHeaderProfiles, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.HttpHeaderProfiles{}.AttrTypes(), &item)
		unpacked, d := unpackHttpHeaderProfilesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.HttpHeaderProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for HttpHeaderProfiles ---
func packHttpHeaderProfilesListFromSdk(ctx context.Context, sdks []security_services.HttpHeaderProfiles) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.HttpHeaderProfiles")
	diags := diag.Diagnostics{}
	var data []models.HttpHeaderProfiles

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.HttpHeaderProfiles
		obj, d := packHttpHeaderProfilesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.HttpHeaderProfiles{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.HttpHeaderProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.HttpHeaderProfiles{}.AttrType(), data)
}

// --- Unpacker for HttpHeaderProfilesHttpHeaderInsertionInner ---
func unpackHttpHeaderProfilesHttpHeaderInsertionInnerToSdk(ctx context.Context, obj types.Object) (*security_services.HttpHeaderProfilesHttpHeaderInsertionInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.HttpHeaderProfilesHttpHeaderInsertionInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.HttpHeaderProfilesHttpHeaderInsertionInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.HttpHeaderProfilesHttpHeaderInsertionInner
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	// Handling Lists
	if !model.Type.IsNull() && !model.Type.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field Type")
		unpacked, d := unpackHttpHeaderProfilesHttpHeaderInsertionInnerTypeInnerListToSdk(ctx, model.Type)
		diags.Append(d...)
		sdk.Type = unpacked
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.HttpHeaderProfilesHttpHeaderInsertionInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for HttpHeaderProfilesHttpHeaderInsertionInner ---
func packHttpHeaderProfilesHttpHeaderInsertionInnerFromSdk(ctx context.Context, sdk security_services.HttpHeaderProfilesHttpHeaderInsertionInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.HttpHeaderProfilesHttpHeaderInsertionInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.HttpHeaderProfilesHttpHeaderInsertionInner
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	// Handling Lists
	if sdk.Type != nil {
		tflog.Debug(ctx, "Packing list of objects for field Type")
		packed, d := packHttpHeaderProfilesHttpHeaderInsertionInnerTypeInnerListFromSdk(ctx, sdk.Type)
		diags.Append(d...)
		model.Type = packed
	} else {
		model.Type = basetypes.NewListNull(models.HttpHeaderProfilesHttpHeaderInsertionInnerTypeInner{}.AttrType())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.HttpHeaderProfilesHttpHeaderInsertionInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.HttpHeaderProfilesHttpHeaderInsertionInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for HttpHeaderProfilesHttpHeaderInsertionInner ---
func unpackHttpHeaderProfilesHttpHeaderInsertionInnerListToSdk(ctx context.Context, list types.List) ([]security_services.HttpHeaderProfilesHttpHeaderInsertionInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.HttpHeaderProfilesHttpHeaderInsertionInner")
	diags := diag.Diagnostics{}
	var data []models.HttpHeaderProfilesHttpHeaderInsertionInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.HttpHeaderProfilesHttpHeaderInsertionInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.HttpHeaderProfilesHttpHeaderInsertionInner{}.AttrTypes(), &item)
		unpacked, d := unpackHttpHeaderProfilesHttpHeaderInsertionInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.HttpHeaderProfilesHttpHeaderInsertionInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for HttpHeaderProfilesHttpHeaderInsertionInner ---
func packHttpHeaderProfilesHttpHeaderInsertionInnerListFromSdk(ctx context.Context, sdks []security_services.HttpHeaderProfilesHttpHeaderInsertionInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.HttpHeaderProfilesHttpHeaderInsertionInner")
	diags := diag.Diagnostics{}
	var data []models.HttpHeaderProfilesHttpHeaderInsertionInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.HttpHeaderProfilesHttpHeaderInsertionInner
		obj, d := packHttpHeaderProfilesHttpHeaderInsertionInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.HttpHeaderProfilesHttpHeaderInsertionInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.HttpHeaderProfilesHttpHeaderInsertionInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.HttpHeaderProfilesHttpHeaderInsertionInner{}.AttrType(), data)
}

// --- Unpacker for HttpHeaderProfilesHttpHeaderInsertionInnerTypeInner ---
func unpackHttpHeaderProfilesHttpHeaderInsertionInnerTypeInnerToSdk(ctx context.Context, obj types.Object) (*security_services.HttpHeaderProfilesHttpHeaderInsertionInnerTypeInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.HttpHeaderProfilesHttpHeaderInsertionInnerTypeInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.HttpHeaderProfilesHttpHeaderInsertionInnerTypeInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.HttpHeaderProfilesHttpHeaderInsertionInnerTypeInner
	var d diag.Diagnostics
	// Handling Lists
	if !model.Domains.IsNull() && !model.Domains.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Domains")
		diags.Append(model.Domains.ElementsAs(ctx, &sdk.Domains, false)...)
	}

	// Handling Lists
	if !model.Headers.IsNull() && !model.Headers.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field Headers")
		unpacked, d := unpackHttpHeaderProfilesHttpHeaderInsertionInnerTypeInnerHeadersInnerListToSdk(ctx, model.Headers)
		diags.Append(d...)
		sdk.Headers = unpacked
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.HttpHeaderProfilesHttpHeaderInsertionInnerTypeInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for HttpHeaderProfilesHttpHeaderInsertionInnerTypeInner ---
func packHttpHeaderProfilesHttpHeaderInsertionInnerTypeInnerFromSdk(ctx context.Context, sdk security_services.HttpHeaderProfilesHttpHeaderInsertionInnerTypeInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.HttpHeaderProfilesHttpHeaderInsertionInnerTypeInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.HttpHeaderProfilesHttpHeaderInsertionInnerTypeInner
	var d diag.Diagnostics
	// Handling Lists
	if sdk.Domains != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Domains")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Domains, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Domains)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Domains = basetypes.NewListNull(elemType)
	}
	// Handling Lists
	if sdk.Headers != nil {
		tflog.Debug(ctx, "Packing list of objects for field Headers")
		packed, d := packHttpHeaderProfilesHttpHeaderInsertionInnerTypeInnerHeadersInnerListFromSdk(ctx, sdk.Headers)
		diags.Append(d...)
		model.Headers = packed
	} else {
		model.Headers = basetypes.NewListNull(models.HttpHeaderProfilesHttpHeaderInsertionInnerTypeInnerHeadersInner{}.AttrType())
	}
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.HttpHeaderProfilesHttpHeaderInsertionInnerTypeInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.HttpHeaderProfilesHttpHeaderInsertionInnerTypeInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for HttpHeaderProfilesHttpHeaderInsertionInnerTypeInner ---
func unpackHttpHeaderProfilesHttpHeaderInsertionInnerTypeInnerListToSdk(ctx context.Context, list types.List) ([]security_services.HttpHeaderProfilesHttpHeaderInsertionInnerTypeInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.HttpHeaderProfilesHttpHeaderInsertionInnerTypeInner")
	diags := diag.Diagnostics{}
	var data []models.HttpHeaderProfilesHttpHeaderInsertionInnerTypeInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.HttpHeaderProfilesHttpHeaderInsertionInnerTypeInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.HttpHeaderProfilesHttpHeaderInsertionInnerTypeInner{}.AttrTypes(), &item)
		unpacked, d := unpackHttpHeaderProfilesHttpHeaderInsertionInnerTypeInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.HttpHeaderProfilesHttpHeaderInsertionInnerTypeInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for HttpHeaderProfilesHttpHeaderInsertionInnerTypeInner ---
func packHttpHeaderProfilesHttpHeaderInsertionInnerTypeInnerListFromSdk(ctx context.Context, sdks []security_services.HttpHeaderProfilesHttpHeaderInsertionInnerTypeInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.HttpHeaderProfilesHttpHeaderInsertionInnerTypeInner")
	diags := diag.Diagnostics{}
	var data []models.HttpHeaderProfilesHttpHeaderInsertionInnerTypeInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.HttpHeaderProfilesHttpHeaderInsertionInnerTypeInner
		obj, d := packHttpHeaderProfilesHttpHeaderInsertionInnerTypeInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.HttpHeaderProfilesHttpHeaderInsertionInnerTypeInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.HttpHeaderProfilesHttpHeaderInsertionInnerTypeInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.HttpHeaderProfilesHttpHeaderInsertionInnerTypeInner{}.AttrType(), data)
}

// --- Unpacker for HttpHeaderProfilesHttpHeaderInsertionInnerTypeInnerHeadersInner ---
func unpackHttpHeaderProfilesHttpHeaderInsertionInnerTypeInnerHeadersInnerToSdk(ctx context.Context, obj types.Object) (*security_services.HttpHeaderProfilesHttpHeaderInsertionInnerTypeInnerHeadersInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.HttpHeaderProfilesHttpHeaderInsertionInnerTypeInnerHeadersInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.HttpHeaderProfilesHttpHeaderInsertionInnerTypeInnerHeadersInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.HttpHeaderProfilesHttpHeaderInsertionInnerTypeInnerHeadersInner
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Header.IsNull() && !model.Header.IsUnknown() {
		sdk.Header = model.Header.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Header", "value": sdk.Header})
	}

	// Handling Primitives
	if !model.Log.IsNull() && !model.Log.IsUnknown() {
		sdk.Log = model.Log.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Log", "value": *sdk.Log})
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	// Handling Primitives
	if !model.Value.IsNull() && !model.Value.IsUnknown() {
		sdk.Value = model.Value.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Value", "value": sdk.Value})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.HttpHeaderProfilesHttpHeaderInsertionInnerTypeInnerHeadersInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for HttpHeaderProfilesHttpHeaderInsertionInnerTypeInnerHeadersInner ---
func packHttpHeaderProfilesHttpHeaderInsertionInnerTypeInnerHeadersInnerFromSdk(ctx context.Context, sdk security_services.HttpHeaderProfilesHttpHeaderInsertionInnerTypeInnerHeadersInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.HttpHeaderProfilesHttpHeaderInsertionInnerTypeInnerHeadersInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.HttpHeaderProfilesHttpHeaderInsertionInnerTypeInnerHeadersInner
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.Header = basetypes.NewStringValue(sdk.Header)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Header", "value": sdk.Header})
	// Handling Primitives
	// Standard primitive packing
	if sdk.Log != nil {
		model.Log = basetypes.NewBoolValue(*sdk.Log)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Log", "value": *sdk.Log})
	} else {
		model.Log = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	// Handling Primitives
	// Standard primitive packing
	model.Value = basetypes.NewStringValue(sdk.Value)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Value", "value": sdk.Value})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.HttpHeaderProfilesHttpHeaderInsertionInnerTypeInnerHeadersInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.HttpHeaderProfilesHttpHeaderInsertionInnerTypeInnerHeadersInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for HttpHeaderProfilesHttpHeaderInsertionInnerTypeInnerHeadersInner ---
func unpackHttpHeaderProfilesHttpHeaderInsertionInnerTypeInnerHeadersInnerListToSdk(ctx context.Context, list types.List) ([]security_services.HttpHeaderProfilesHttpHeaderInsertionInnerTypeInnerHeadersInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.HttpHeaderProfilesHttpHeaderInsertionInnerTypeInnerHeadersInner")
	diags := diag.Diagnostics{}
	var data []models.HttpHeaderProfilesHttpHeaderInsertionInnerTypeInnerHeadersInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.HttpHeaderProfilesHttpHeaderInsertionInnerTypeInnerHeadersInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.HttpHeaderProfilesHttpHeaderInsertionInnerTypeInnerHeadersInner{}.AttrTypes(), &item)
		unpacked, d := unpackHttpHeaderProfilesHttpHeaderInsertionInnerTypeInnerHeadersInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.HttpHeaderProfilesHttpHeaderInsertionInnerTypeInnerHeadersInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for HttpHeaderProfilesHttpHeaderInsertionInnerTypeInnerHeadersInner ---
func packHttpHeaderProfilesHttpHeaderInsertionInnerTypeInnerHeadersInnerListFromSdk(ctx context.Context, sdks []security_services.HttpHeaderProfilesHttpHeaderInsertionInnerTypeInnerHeadersInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.HttpHeaderProfilesHttpHeaderInsertionInnerTypeInnerHeadersInner")
	diags := diag.Diagnostics{}
	var data []models.HttpHeaderProfilesHttpHeaderInsertionInnerTypeInnerHeadersInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.HttpHeaderProfilesHttpHeaderInsertionInnerTypeInnerHeadersInner
		obj, d := packHttpHeaderProfilesHttpHeaderInsertionInnerTypeInnerHeadersInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.HttpHeaderProfilesHttpHeaderInsertionInnerTypeInnerHeadersInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.HttpHeaderProfilesHttpHeaderInsertionInnerTypeInnerHeadersInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.HttpHeaderProfilesHttpHeaderInsertionInnerTypeInnerHeadersInner{}.AttrType(), data)
}
