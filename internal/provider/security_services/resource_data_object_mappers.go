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

// --- Unpacker for DataObjects ---
func unpackDataObjectsToSdk(ctx context.Context, obj types.Object) (*security_services.DataObjects, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.DataObjects", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.DataObjects
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.DataObjects
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
	if !model.DisableOverride.IsNull() && !model.DisableOverride.IsUnknown() {
		sdk.DisableOverride = model.DisableOverride.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DisableOverride", "value": *sdk.DisableOverride})
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

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
	}

	// Handling Objects
	if !model.PatternType.IsNull() && !model.PatternType.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field PatternType")
		unpacked, d := unpackDataObjectsPatternTypeToSdk(ctx, model.PatternType)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "PatternType"})
		}
		if unpacked != nil {
			sdk.PatternType = unpacked
		}
	}

	// Handling Primitives
	if !model.Snippet.IsNull() && !model.Snippet.IsUnknown() {
		sdk.Snippet = model.Snippet.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.DataObjects", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for DataObjects ---
func packDataObjectsFromSdk(ctx context.Context, sdk security_services.DataObjects) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.DataObjects", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.DataObjects
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
	if sdk.DisableOverride != nil {
		model.DisableOverride = basetypes.NewStringValue(*sdk.DisableOverride)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DisableOverride", "value": *sdk.DisableOverride})
	} else {
		model.DisableOverride = basetypes.NewStringNull()
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
	// Handling Primitives
	// Standard primitive packing
	if sdk.Name != nil {
		model.Name = basetypes.NewStringValue(*sdk.Name)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
	} else {
		model.Name = basetypes.NewStringNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.PatternType != nil {
		tflog.Debug(ctx, "Packing nested object for field PatternType")
		packed, d := packDataObjectsPatternTypeFromSdk(ctx, *sdk.PatternType)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "PatternType"})
		}
		model.PatternType = packed
	} else {
		model.PatternType = basetypes.NewObjectNull(models.DataObjectsPatternType{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Snippet != nil {
		model.Snippet = basetypes.NewStringValue(*sdk.Snippet)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	} else {
		model.Snippet = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.DataObjects{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.DataObjects", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for DataObjects ---
func unpackDataObjectsListToSdk(ctx context.Context, list types.List) ([]security_services.DataObjects, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.DataObjects")
	diags := diag.Diagnostics{}
	var data []models.DataObjects
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.DataObjects, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.DataObjects{}.AttrTypes(), &item)
		unpacked, d := unpackDataObjectsToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.DataObjects", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for DataObjects ---
func packDataObjectsListFromSdk(ctx context.Context, sdks []security_services.DataObjects) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.DataObjects")
	diags := diag.Diagnostics{}
	var data []models.DataObjects

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.DataObjects
		obj, d := packDataObjectsFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.DataObjects{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.DataObjects", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.DataObjects{}.AttrType(), data)
}

// --- Unpacker for DataObjectsPatternType ---
func unpackDataObjectsPatternTypeToSdk(ctx context.Context, obj types.Object) (*security_services.DataObjectsPatternType, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.DataObjectsPatternType", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.DataObjectsPatternType
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.DataObjectsPatternType
	var d diag.Diagnostics
	// Handling Objects
	if !model.FileProperties.IsNull() && !model.FileProperties.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field FileProperties")
		unpacked, d := unpackDataObjectsPatternTypeFilePropertiesToSdk(ctx, model.FileProperties)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "FileProperties"})
		}
		if unpacked != nil {
			sdk.FileProperties = unpacked
		}
	}

	// Handling Objects
	if !model.Predefined.IsNull() && !model.Predefined.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Predefined")
		unpacked, d := unpackDataObjectsPatternTypePredefinedToSdk(ctx, model.Predefined)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Predefined"})
		}
		if unpacked != nil {
			sdk.Predefined = unpacked
		}
	}

	// Handling Objects
	if !model.Regex.IsNull() && !model.Regex.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Regex")
		unpacked, d := unpackDataObjectsPatternTypeRegexToSdk(ctx, model.Regex)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Regex"})
		}
		if unpacked != nil {
			sdk.Regex = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.DataObjectsPatternType", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for DataObjectsPatternType ---
func packDataObjectsPatternTypeFromSdk(ctx context.Context, sdk security_services.DataObjectsPatternType) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.DataObjectsPatternType", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.DataObjectsPatternType
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.FileProperties != nil {
		tflog.Debug(ctx, "Packing nested object for field FileProperties")
		packed, d := packDataObjectsPatternTypeFilePropertiesFromSdk(ctx, *sdk.FileProperties)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "FileProperties"})
		}
		model.FileProperties = packed
	} else {
		model.FileProperties = basetypes.NewObjectNull(models.DataObjectsPatternTypeFileProperties{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Predefined != nil {
		tflog.Debug(ctx, "Packing nested object for field Predefined")
		packed, d := packDataObjectsPatternTypePredefinedFromSdk(ctx, *sdk.Predefined)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Predefined"})
		}
		model.Predefined = packed
	} else {
		model.Predefined = basetypes.NewObjectNull(models.DataObjectsPatternTypePredefined{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Regex != nil {
		tflog.Debug(ctx, "Packing nested object for field Regex")
		packed, d := packDataObjectsPatternTypeRegexFromSdk(ctx, *sdk.Regex)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Regex"})
		}
		model.Regex = packed
	} else {
		model.Regex = basetypes.NewObjectNull(models.DataObjectsPatternTypeRegex{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.DataObjectsPatternType{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.DataObjectsPatternType", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for DataObjectsPatternType ---
func unpackDataObjectsPatternTypeListToSdk(ctx context.Context, list types.List) ([]security_services.DataObjectsPatternType, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.DataObjectsPatternType")
	diags := diag.Diagnostics{}
	var data []models.DataObjectsPatternType
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.DataObjectsPatternType, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.DataObjectsPatternType{}.AttrTypes(), &item)
		unpacked, d := unpackDataObjectsPatternTypeToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.DataObjectsPatternType", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for DataObjectsPatternType ---
func packDataObjectsPatternTypeListFromSdk(ctx context.Context, sdks []security_services.DataObjectsPatternType) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.DataObjectsPatternType")
	diags := diag.Diagnostics{}
	var data []models.DataObjectsPatternType

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.DataObjectsPatternType
		obj, d := packDataObjectsPatternTypeFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.DataObjectsPatternType{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.DataObjectsPatternType", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.DataObjectsPatternType{}.AttrType(), data)
}

// --- Unpacker for DataObjectsPatternTypeFileProperties ---
func unpackDataObjectsPatternTypeFilePropertiesToSdk(ctx context.Context, obj types.Object) (*security_services.DataObjectsPatternTypeFileProperties, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.DataObjectsPatternTypeFileProperties", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.DataObjectsPatternTypeFileProperties
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.DataObjectsPatternTypeFileProperties
	var d diag.Diagnostics
	// Handling Lists
	if !model.Pattern.IsNull() && !model.Pattern.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field Pattern")
		unpacked, d := unpackDataObjectsPatternTypeFilePropertiesPatternInnerListToSdk(ctx, model.Pattern)
		diags.Append(d...)
		sdk.Pattern = unpacked
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.DataObjectsPatternTypeFileProperties", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for DataObjectsPatternTypeFileProperties ---
func packDataObjectsPatternTypeFilePropertiesFromSdk(ctx context.Context, sdk security_services.DataObjectsPatternTypeFileProperties) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.DataObjectsPatternTypeFileProperties", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.DataObjectsPatternTypeFileProperties
	var d diag.Diagnostics
	// Handling Lists
	if sdk.Pattern != nil {
		tflog.Debug(ctx, "Packing list of objects for field Pattern")
		packed, d := packDataObjectsPatternTypeFilePropertiesPatternInnerListFromSdk(ctx, sdk.Pattern)
		diags.Append(d...)
		model.Pattern = packed
	} else {
		model.Pattern = basetypes.NewListNull(models.DataObjectsPatternTypeFilePropertiesPatternInner{}.AttrType())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.DataObjectsPatternTypeFileProperties{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.DataObjectsPatternTypeFileProperties", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for DataObjectsPatternTypeFileProperties ---
func unpackDataObjectsPatternTypeFilePropertiesListToSdk(ctx context.Context, list types.List) ([]security_services.DataObjectsPatternTypeFileProperties, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.DataObjectsPatternTypeFileProperties")
	diags := diag.Diagnostics{}
	var data []models.DataObjectsPatternTypeFileProperties
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.DataObjectsPatternTypeFileProperties, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.DataObjectsPatternTypeFileProperties{}.AttrTypes(), &item)
		unpacked, d := unpackDataObjectsPatternTypeFilePropertiesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.DataObjectsPatternTypeFileProperties", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for DataObjectsPatternTypeFileProperties ---
func packDataObjectsPatternTypeFilePropertiesListFromSdk(ctx context.Context, sdks []security_services.DataObjectsPatternTypeFileProperties) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.DataObjectsPatternTypeFileProperties")
	diags := diag.Diagnostics{}
	var data []models.DataObjectsPatternTypeFileProperties

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.DataObjectsPatternTypeFileProperties
		obj, d := packDataObjectsPatternTypeFilePropertiesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.DataObjectsPatternTypeFileProperties{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.DataObjectsPatternTypeFileProperties", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.DataObjectsPatternTypeFileProperties{}.AttrType(), data)
}

// --- Unpacker for DataObjectsPatternTypeFilePropertiesPatternInner ---
func unpackDataObjectsPatternTypeFilePropertiesPatternInnerToSdk(ctx context.Context, obj types.Object) (*security_services.DataObjectsPatternTypeFilePropertiesPatternInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.DataObjectsPatternTypeFilePropertiesPatternInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.DataObjectsPatternTypeFilePropertiesPatternInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.DataObjectsPatternTypeFilePropertiesPatternInner
	var d diag.Diagnostics
	// Handling Primitives
	if !model.FileProperty.IsNull() && !model.FileProperty.IsUnknown() {
		sdk.FileProperty = model.FileProperty.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "FileProperty", "value": *sdk.FileProperty})
	}

	// Handling Primitives
	if !model.FileType.IsNull() && !model.FileType.IsUnknown() {
		sdk.FileType = model.FileType.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "FileType", "value": *sdk.FileType})
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
	}

	// Handling Primitives
	if !model.PropertyValue.IsNull() && !model.PropertyValue.IsUnknown() {
		sdk.PropertyValue = model.PropertyValue.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "PropertyValue", "value": *sdk.PropertyValue})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.DataObjectsPatternTypeFilePropertiesPatternInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for DataObjectsPatternTypeFilePropertiesPatternInner ---
func packDataObjectsPatternTypeFilePropertiesPatternInnerFromSdk(ctx context.Context, sdk security_services.DataObjectsPatternTypeFilePropertiesPatternInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.DataObjectsPatternTypeFilePropertiesPatternInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.DataObjectsPatternTypeFilePropertiesPatternInner
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.FileProperty != nil {
		model.FileProperty = basetypes.NewStringValue(*sdk.FileProperty)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "FileProperty", "value": *sdk.FileProperty})
	} else {
		model.FileProperty = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.FileType != nil {
		model.FileType = basetypes.NewStringValue(*sdk.FileType)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "FileType", "value": *sdk.FileType})
	} else {
		model.FileType = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Name != nil {
		model.Name = basetypes.NewStringValue(*sdk.Name)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
	} else {
		model.Name = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.PropertyValue != nil {
		model.PropertyValue = basetypes.NewStringValue(*sdk.PropertyValue)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "PropertyValue", "value": *sdk.PropertyValue})
	} else {
		model.PropertyValue = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.DataObjectsPatternTypeFilePropertiesPatternInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.DataObjectsPatternTypeFilePropertiesPatternInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for DataObjectsPatternTypeFilePropertiesPatternInner ---
func unpackDataObjectsPatternTypeFilePropertiesPatternInnerListToSdk(ctx context.Context, list types.List) ([]security_services.DataObjectsPatternTypeFilePropertiesPatternInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.DataObjectsPatternTypeFilePropertiesPatternInner")
	diags := diag.Diagnostics{}
	var data []models.DataObjectsPatternTypeFilePropertiesPatternInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.DataObjectsPatternTypeFilePropertiesPatternInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.DataObjectsPatternTypeFilePropertiesPatternInner{}.AttrTypes(), &item)
		unpacked, d := unpackDataObjectsPatternTypeFilePropertiesPatternInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.DataObjectsPatternTypeFilePropertiesPatternInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for DataObjectsPatternTypeFilePropertiesPatternInner ---
func packDataObjectsPatternTypeFilePropertiesPatternInnerListFromSdk(ctx context.Context, sdks []security_services.DataObjectsPatternTypeFilePropertiesPatternInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.DataObjectsPatternTypeFilePropertiesPatternInner")
	diags := diag.Diagnostics{}
	var data []models.DataObjectsPatternTypeFilePropertiesPatternInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.DataObjectsPatternTypeFilePropertiesPatternInner
		obj, d := packDataObjectsPatternTypeFilePropertiesPatternInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.DataObjectsPatternTypeFilePropertiesPatternInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.DataObjectsPatternTypeFilePropertiesPatternInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.DataObjectsPatternTypeFilePropertiesPatternInner{}.AttrType(), data)
}

// --- Unpacker for DataObjectsPatternTypePredefined ---
func unpackDataObjectsPatternTypePredefinedToSdk(ctx context.Context, obj types.Object) (*security_services.DataObjectsPatternTypePredefined, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.DataObjectsPatternTypePredefined", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.DataObjectsPatternTypePredefined
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.DataObjectsPatternTypePredefined
	var d diag.Diagnostics
	// Handling Lists
	if !model.Pattern.IsNull() && !model.Pattern.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field Pattern")
		unpacked, d := unpackDataObjectsPatternTypePredefinedPatternInnerListToSdk(ctx, model.Pattern)
		diags.Append(d...)
		sdk.Pattern = unpacked
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.DataObjectsPatternTypePredefined", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for DataObjectsPatternTypePredefined ---
func packDataObjectsPatternTypePredefinedFromSdk(ctx context.Context, sdk security_services.DataObjectsPatternTypePredefined) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.DataObjectsPatternTypePredefined", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.DataObjectsPatternTypePredefined
	var d diag.Diagnostics
	// Handling Lists
	if sdk.Pattern != nil {
		tflog.Debug(ctx, "Packing list of objects for field Pattern")
		packed, d := packDataObjectsPatternTypePredefinedPatternInnerListFromSdk(ctx, sdk.Pattern)
		diags.Append(d...)
		model.Pattern = packed
	} else {
		model.Pattern = basetypes.NewListNull(models.DataObjectsPatternTypePredefinedPatternInner{}.AttrType())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.DataObjectsPatternTypePredefined{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.DataObjectsPatternTypePredefined", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for DataObjectsPatternTypePredefined ---
func unpackDataObjectsPatternTypePredefinedListToSdk(ctx context.Context, list types.List) ([]security_services.DataObjectsPatternTypePredefined, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.DataObjectsPatternTypePredefined")
	diags := diag.Diagnostics{}
	var data []models.DataObjectsPatternTypePredefined
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.DataObjectsPatternTypePredefined, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.DataObjectsPatternTypePredefined{}.AttrTypes(), &item)
		unpacked, d := unpackDataObjectsPatternTypePredefinedToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.DataObjectsPatternTypePredefined", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for DataObjectsPatternTypePredefined ---
func packDataObjectsPatternTypePredefinedListFromSdk(ctx context.Context, sdks []security_services.DataObjectsPatternTypePredefined) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.DataObjectsPatternTypePredefined")
	diags := diag.Diagnostics{}
	var data []models.DataObjectsPatternTypePredefined

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.DataObjectsPatternTypePredefined
		obj, d := packDataObjectsPatternTypePredefinedFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.DataObjectsPatternTypePredefined{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.DataObjectsPatternTypePredefined", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.DataObjectsPatternTypePredefined{}.AttrType(), data)
}

// --- Unpacker for DataObjectsPatternTypePredefinedPatternInner ---
func unpackDataObjectsPatternTypePredefinedPatternInnerToSdk(ctx context.Context, obj types.Object) (*security_services.DataObjectsPatternTypePredefinedPatternInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.DataObjectsPatternTypePredefinedPatternInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.DataObjectsPatternTypePredefinedPatternInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.DataObjectsPatternTypePredefinedPatternInner
	var d diag.Diagnostics
	// Handling Lists
	if !model.FileType.IsNull() && !model.FileType.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field FileType")
		diags.Append(model.FileType.ElementsAs(ctx, &sdk.FileType, false)...)
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.DataObjectsPatternTypePredefinedPatternInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for DataObjectsPatternTypePredefinedPatternInner ---
func packDataObjectsPatternTypePredefinedPatternInnerFromSdk(ctx context.Context, sdk security_services.DataObjectsPatternTypePredefinedPatternInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.DataObjectsPatternTypePredefinedPatternInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.DataObjectsPatternTypePredefinedPatternInner
	var d diag.Diagnostics
	// Handling Lists
	if sdk.FileType != nil {
		tflog.Debug(ctx, "Packing list of primitives for field FileType")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.FileType, d = basetypes.NewListValueFrom(ctx, elemType, sdk.FileType)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.FileType = basetypes.NewListNull(elemType)
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Name != nil {
		model.Name = basetypes.NewStringValue(*sdk.Name)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
	} else {
		model.Name = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.DataObjectsPatternTypePredefinedPatternInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.DataObjectsPatternTypePredefinedPatternInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for DataObjectsPatternTypePredefinedPatternInner ---
func unpackDataObjectsPatternTypePredefinedPatternInnerListToSdk(ctx context.Context, list types.List) ([]security_services.DataObjectsPatternTypePredefinedPatternInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.DataObjectsPatternTypePredefinedPatternInner")
	diags := diag.Diagnostics{}
	var data []models.DataObjectsPatternTypePredefinedPatternInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.DataObjectsPatternTypePredefinedPatternInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.DataObjectsPatternTypePredefinedPatternInner{}.AttrTypes(), &item)
		unpacked, d := unpackDataObjectsPatternTypePredefinedPatternInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.DataObjectsPatternTypePredefinedPatternInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for DataObjectsPatternTypePredefinedPatternInner ---
func packDataObjectsPatternTypePredefinedPatternInnerListFromSdk(ctx context.Context, sdks []security_services.DataObjectsPatternTypePredefinedPatternInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.DataObjectsPatternTypePredefinedPatternInner")
	diags := diag.Diagnostics{}
	var data []models.DataObjectsPatternTypePredefinedPatternInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.DataObjectsPatternTypePredefinedPatternInner
		obj, d := packDataObjectsPatternTypePredefinedPatternInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.DataObjectsPatternTypePredefinedPatternInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.DataObjectsPatternTypePredefinedPatternInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.DataObjectsPatternTypePredefinedPatternInner{}.AttrType(), data)
}

// --- Unpacker for DataObjectsPatternTypeRegex ---
func unpackDataObjectsPatternTypeRegexToSdk(ctx context.Context, obj types.Object) (*security_services.DataObjectsPatternTypeRegex, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.DataObjectsPatternTypeRegex", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.DataObjectsPatternTypeRegex
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.DataObjectsPatternTypeRegex
	var d diag.Diagnostics
	// Handling Lists
	if !model.Pattern.IsNull() && !model.Pattern.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field Pattern")
		unpacked, d := unpackDataObjectsPatternTypeRegexPatternInnerListToSdk(ctx, model.Pattern)
		diags.Append(d...)
		sdk.Pattern = unpacked
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.DataObjectsPatternTypeRegex", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for DataObjectsPatternTypeRegex ---
func packDataObjectsPatternTypeRegexFromSdk(ctx context.Context, sdk security_services.DataObjectsPatternTypeRegex) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.DataObjectsPatternTypeRegex", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.DataObjectsPatternTypeRegex
	var d diag.Diagnostics
	// Handling Lists
	if sdk.Pattern != nil {
		tflog.Debug(ctx, "Packing list of objects for field Pattern")
		packed, d := packDataObjectsPatternTypeRegexPatternInnerListFromSdk(ctx, sdk.Pattern)
		diags.Append(d...)
		model.Pattern = packed
	} else {
		model.Pattern = basetypes.NewListNull(models.DataObjectsPatternTypeRegexPatternInner{}.AttrType())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.DataObjectsPatternTypeRegex{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.DataObjectsPatternTypeRegex", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for DataObjectsPatternTypeRegex ---
func unpackDataObjectsPatternTypeRegexListToSdk(ctx context.Context, list types.List) ([]security_services.DataObjectsPatternTypeRegex, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.DataObjectsPatternTypeRegex")
	diags := diag.Diagnostics{}
	var data []models.DataObjectsPatternTypeRegex
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.DataObjectsPatternTypeRegex, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.DataObjectsPatternTypeRegex{}.AttrTypes(), &item)
		unpacked, d := unpackDataObjectsPatternTypeRegexToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.DataObjectsPatternTypeRegex", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for DataObjectsPatternTypeRegex ---
func packDataObjectsPatternTypeRegexListFromSdk(ctx context.Context, sdks []security_services.DataObjectsPatternTypeRegex) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.DataObjectsPatternTypeRegex")
	diags := diag.Diagnostics{}
	var data []models.DataObjectsPatternTypeRegex

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.DataObjectsPatternTypeRegex
		obj, d := packDataObjectsPatternTypeRegexFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.DataObjectsPatternTypeRegex{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.DataObjectsPatternTypeRegex", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.DataObjectsPatternTypeRegex{}.AttrType(), data)
}

// --- Unpacker for DataObjectsPatternTypeRegexPatternInner ---
func unpackDataObjectsPatternTypeRegexPatternInnerToSdk(ctx context.Context, obj types.Object) (*security_services.DataObjectsPatternTypeRegexPatternInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.DataObjectsPatternTypeRegexPatternInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.DataObjectsPatternTypeRegexPatternInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.DataObjectsPatternTypeRegexPatternInner
	var d diag.Diagnostics
	// Handling Lists
	if !model.FileType.IsNull() && !model.FileType.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field FileType")
		diags.Append(model.FileType.ElementsAs(ctx, &sdk.FileType, false)...)
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
	}

	// Handling Primitives
	if !model.Regex.IsNull() && !model.Regex.IsUnknown() {
		sdk.Regex = model.Regex.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Regex", "value": *sdk.Regex})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.DataObjectsPatternTypeRegexPatternInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for DataObjectsPatternTypeRegexPatternInner ---
func packDataObjectsPatternTypeRegexPatternInnerFromSdk(ctx context.Context, sdk security_services.DataObjectsPatternTypeRegexPatternInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.DataObjectsPatternTypeRegexPatternInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.DataObjectsPatternTypeRegexPatternInner
	var d diag.Diagnostics
	// Handling Lists
	if sdk.FileType != nil {
		tflog.Debug(ctx, "Packing list of primitives for field FileType")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.FileType, d = basetypes.NewListValueFrom(ctx, elemType, sdk.FileType)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.FileType = basetypes.NewListNull(elemType)
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Name != nil {
		model.Name = basetypes.NewStringValue(*sdk.Name)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
	} else {
		model.Name = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Regex != nil {
		model.Regex = basetypes.NewStringValue(*sdk.Regex)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Regex", "value": *sdk.Regex})
	} else {
		model.Regex = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.DataObjectsPatternTypeRegexPatternInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.DataObjectsPatternTypeRegexPatternInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for DataObjectsPatternTypeRegexPatternInner ---
func unpackDataObjectsPatternTypeRegexPatternInnerListToSdk(ctx context.Context, list types.List) ([]security_services.DataObjectsPatternTypeRegexPatternInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.DataObjectsPatternTypeRegexPatternInner")
	diags := diag.Diagnostics{}
	var data []models.DataObjectsPatternTypeRegexPatternInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.DataObjectsPatternTypeRegexPatternInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.DataObjectsPatternTypeRegexPatternInner{}.AttrTypes(), &item)
		unpacked, d := unpackDataObjectsPatternTypeRegexPatternInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.DataObjectsPatternTypeRegexPatternInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for DataObjectsPatternTypeRegexPatternInner ---
func packDataObjectsPatternTypeRegexPatternInnerListFromSdk(ctx context.Context, sdks []security_services.DataObjectsPatternTypeRegexPatternInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.DataObjectsPatternTypeRegexPatternInner")
	diags := diag.Diagnostics{}
	var data []models.DataObjectsPatternTypeRegexPatternInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.DataObjectsPatternTypeRegexPatternInner
		obj, d := packDataObjectsPatternTypeRegexPatternInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.DataObjectsPatternTypeRegexPatternInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.DataObjectsPatternTypeRegexPatternInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.DataObjectsPatternTypeRegexPatternInner{}.AttrType(), data)
}
