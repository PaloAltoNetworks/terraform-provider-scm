package provider

import (
	"context"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/objects"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/objects"
)

// --- Unpacker for Schedules ---
func unpackSchedulesToSdk(ctx context.Context, obj types.Object) (*objects.Schedules, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.Schedules", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.Schedules
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.Schedules
	var d diag.Diagnostics
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
		sdk.Id = model.Id.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Id", "value": sdk.Id})
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	// Handling Objects
	if !model.ScheduleType.IsNull() && !model.ScheduleType.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field ScheduleType")
		unpacked, d := unpackSchedulesScheduleTypeToSdk(ctx, model.ScheduleType)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "ScheduleType"})
		}
		if unpacked != nil {
			sdk.ScheduleType = *unpacked
		}
	}

	// Handling Primitives
	if !model.Snippet.IsNull() && !model.Snippet.IsUnknown() {
		sdk.Snippet = model.Snippet.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.Schedules", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for Schedules ---
func packSchedulesFromSdk(ctx context.Context, sdk objects.Schedules) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.Schedules", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.Schedules
	var d diag.Diagnostics
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
	model.Id = basetypes.NewStringValue(sdk.Id)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Id", "value": sdk.Id})
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	// Handling Objects
	// This is a regular nested object that has its own packer.
	// Logic for non-pointer / value-type nested objects
	if !reflect.ValueOf(sdk.ScheduleType).IsZero() {
		tflog.Debug(ctx, "Packing nested object for field ScheduleType")
		packed, d := packSchedulesScheduleTypeFromSdk(ctx, sdk.ScheduleType)
		diags.Append(d...)
		model.ScheduleType = packed
	} else {
		model.ScheduleType = basetypes.NewObjectNull(models.SchedulesScheduleType{}.AttrTypes())
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

	obj, d := types.ObjectValueFrom(ctx, models.Schedules{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.Schedules", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for Schedules ---
func unpackSchedulesListToSdk(ctx context.Context, list types.List) ([]objects.Schedules, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.Schedules")
	diags := diag.Diagnostics{}
	var data []models.Schedules
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.Schedules, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.Schedules{}.AttrTypes(), &item)
		unpacked, d := unpackSchedulesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.Schedules", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for Schedules ---
func packSchedulesListFromSdk(ctx context.Context, sdks []objects.Schedules) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.Schedules")
	diags := diag.Diagnostics{}
	var data []models.Schedules

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.Schedules
		obj, d := packSchedulesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.Schedules{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.Schedules", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.Schedules{}.AttrType(), data)
}

// --- Unpacker for SchedulesScheduleType ---
func unpackSchedulesScheduleTypeToSdk(ctx context.Context, obj types.Object) (*objects.SchedulesScheduleType, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.SchedulesScheduleType", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.SchedulesScheduleType
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.SchedulesScheduleType
	var d diag.Diagnostics
	// Handling Lists
	if !model.NonRecurring.IsNull() && !model.NonRecurring.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field NonRecurring")
		diags.Append(model.NonRecurring.ElementsAs(ctx, &sdk.NonRecurring, false)...)
	}

	// Handling Objects
	if !model.Recurring.IsNull() && !model.Recurring.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Recurring")
		unpacked, d := unpackSchedulesScheduleTypeRecurringToSdk(ctx, model.Recurring)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Recurring"})
		}
		if unpacked != nil {
			sdk.Recurring = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.SchedulesScheduleType", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for SchedulesScheduleType ---
func packSchedulesScheduleTypeFromSdk(ctx context.Context, sdk objects.SchedulesScheduleType) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.SchedulesScheduleType", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.SchedulesScheduleType
	var d diag.Diagnostics
	// Handling Lists
	if sdk.NonRecurring != nil {
		tflog.Debug(ctx, "Packing list of primitives for field NonRecurring")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.NonRecurring, d = basetypes.NewListValueFrom(ctx, elemType, sdk.NonRecurring)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.NonRecurring = basetypes.NewListNull(elemType)
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Recurring != nil {
		tflog.Debug(ctx, "Packing nested object for field Recurring")
		packed, d := packSchedulesScheduleTypeRecurringFromSdk(ctx, *sdk.Recurring)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Recurring"})
		}
		model.Recurring = packed
	} else {
		model.Recurring = basetypes.NewObjectNull(models.SchedulesScheduleTypeRecurring{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.SchedulesScheduleType{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.SchedulesScheduleType", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for SchedulesScheduleType ---
func unpackSchedulesScheduleTypeListToSdk(ctx context.Context, list types.List) ([]objects.SchedulesScheduleType, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.SchedulesScheduleType")
	diags := diag.Diagnostics{}
	var data []models.SchedulesScheduleType
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.SchedulesScheduleType, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.SchedulesScheduleType{}.AttrTypes(), &item)
		unpacked, d := unpackSchedulesScheduleTypeToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.SchedulesScheduleType", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for SchedulesScheduleType ---
func packSchedulesScheduleTypeListFromSdk(ctx context.Context, sdks []objects.SchedulesScheduleType) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.SchedulesScheduleType")
	diags := diag.Diagnostics{}
	var data []models.SchedulesScheduleType

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.SchedulesScheduleType
		obj, d := packSchedulesScheduleTypeFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.SchedulesScheduleType{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.SchedulesScheduleType", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.SchedulesScheduleType{}.AttrType(), data)
}

// --- Unpacker for SchedulesScheduleTypeRecurring ---
func unpackSchedulesScheduleTypeRecurringToSdk(ctx context.Context, obj types.Object) (*objects.SchedulesScheduleTypeRecurring, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.SchedulesScheduleTypeRecurring", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.SchedulesScheduleTypeRecurring
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.SchedulesScheduleTypeRecurring
	var d diag.Diagnostics
	// Handling Lists
	if !model.Daily.IsNull() && !model.Daily.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Daily")
		diags.Append(model.Daily.ElementsAs(ctx, &sdk.Daily, false)...)
	}

	// Handling Objects
	if !model.Weekly.IsNull() && !model.Weekly.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Weekly")
		unpacked, d := unpackSchedulesScheduleTypeRecurringWeeklyToSdk(ctx, model.Weekly)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Weekly"})
		}
		if unpacked != nil {
			sdk.Weekly = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.SchedulesScheduleTypeRecurring", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for SchedulesScheduleTypeRecurring ---
func packSchedulesScheduleTypeRecurringFromSdk(ctx context.Context, sdk objects.SchedulesScheduleTypeRecurring) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.SchedulesScheduleTypeRecurring", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.SchedulesScheduleTypeRecurring
	var d diag.Diagnostics
	// Handling Lists
	if sdk.Daily != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Daily")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Daily, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Daily)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Daily = basetypes.NewListNull(elemType)
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Weekly != nil {
		tflog.Debug(ctx, "Packing nested object for field Weekly")
		packed, d := packSchedulesScheduleTypeRecurringWeeklyFromSdk(ctx, *sdk.Weekly)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Weekly"})
		}
		model.Weekly = packed
	} else {
		model.Weekly = basetypes.NewObjectNull(models.SchedulesScheduleTypeRecurringWeekly{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.SchedulesScheduleTypeRecurring{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.SchedulesScheduleTypeRecurring", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for SchedulesScheduleTypeRecurring ---
func unpackSchedulesScheduleTypeRecurringListToSdk(ctx context.Context, list types.List) ([]objects.SchedulesScheduleTypeRecurring, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.SchedulesScheduleTypeRecurring")
	diags := diag.Diagnostics{}
	var data []models.SchedulesScheduleTypeRecurring
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.SchedulesScheduleTypeRecurring, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.SchedulesScheduleTypeRecurring{}.AttrTypes(), &item)
		unpacked, d := unpackSchedulesScheduleTypeRecurringToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.SchedulesScheduleTypeRecurring", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for SchedulesScheduleTypeRecurring ---
func packSchedulesScheduleTypeRecurringListFromSdk(ctx context.Context, sdks []objects.SchedulesScheduleTypeRecurring) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.SchedulesScheduleTypeRecurring")
	diags := diag.Diagnostics{}
	var data []models.SchedulesScheduleTypeRecurring

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.SchedulesScheduleTypeRecurring
		obj, d := packSchedulesScheduleTypeRecurringFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.SchedulesScheduleTypeRecurring{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.SchedulesScheduleTypeRecurring", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.SchedulesScheduleTypeRecurring{}.AttrType(), data)
}

// --- Unpacker for SchedulesScheduleTypeRecurringWeekly ---
func unpackSchedulesScheduleTypeRecurringWeeklyToSdk(ctx context.Context, obj types.Object) (*objects.SchedulesScheduleTypeRecurringWeekly, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.SchedulesScheduleTypeRecurringWeekly", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.SchedulesScheduleTypeRecurringWeekly
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.SchedulesScheduleTypeRecurringWeekly
	var d diag.Diagnostics
	// Handling Lists
	if !model.Friday.IsNull() && !model.Friday.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Friday")
		diags.Append(model.Friday.ElementsAs(ctx, &sdk.Friday, false)...)
	}

	// Handling Lists
	if !model.Monday.IsNull() && !model.Monday.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Monday")
		diags.Append(model.Monday.ElementsAs(ctx, &sdk.Monday, false)...)
	}

	// Handling Lists
	if !model.Saturday.IsNull() && !model.Saturday.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Saturday")
		diags.Append(model.Saturday.ElementsAs(ctx, &sdk.Saturday, false)...)
	}

	// Handling Lists
	if !model.Sunday.IsNull() && !model.Sunday.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Sunday")
		diags.Append(model.Sunday.ElementsAs(ctx, &sdk.Sunday, false)...)
	}

	// Handling Lists
	if !model.Thursday.IsNull() && !model.Thursday.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Thursday")
		diags.Append(model.Thursday.ElementsAs(ctx, &sdk.Thursday, false)...)
	}

	// Handling Lists
	if !model.Tuesday.IsNull() && !model.Tuesday.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Tuesday")
		diags.Append(model.Tuesday.ElementsAs(ctx, &sdk.Tuesday, false)...)
	}

	// Handling Lists
	if !model.Wednesday.IsNull() && !model.Wednesday.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Wednesday")
		diags.Append(model.Wednesday.ElementsAs(ctx, &sdk.Wednesday, false)...)
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.SchedulesScheduleTypeRecurringWeekly", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for SchedulesScheduleTypeRecurringWeekly ---
func packSchedulesScheduleTypeRecurringWeeklyFromSdk(ctx context.Context, sdk objects.SchedulesScheduleTypeRecurringWeekly) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.SchedulesScheduleTypeRecurringWeekly", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.SchedulesScheduleTypeRecurringWeekly
	var d diag.Diagnostics
	// Handling Lists
	if sdk.Friday != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Friday")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Friday, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Friday)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Friday = basetypes.NewListNull(elemType)
	}
	// Handling Lists
	if sdk.Monday != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Monday")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Monday, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Monday)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Monday = basetypes.NewListNull(elemType)
	}
	// Handling Lists
	if sdk.Saturday != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Saturday")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Saturday, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Saturday)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Saturday = basetypes.NewListNull(elemType)
	}
	// Handling Lists
	if sdk.Sunday != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Sunday")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Sunday, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Sunday)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Sunday = basetypes.NewListNull(elemType)
	}
	// Handling Lists
	if sdk.Thursday != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Thursday")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Thursday, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Thursday)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Thursday = basetypes.NewListNull(elemType)
	}
	// Handling Lists
	if sdk.Tuesday != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Tuesday")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Tuesday, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Tuesday)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Tuesday = basetypes.NewListNull(elemType)
	}
	// Handling Lists
	if sdk.Wednesday != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Wednesday")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Wednesday, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Wednesday)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Wednesday = basetypes.NewListNull(elemType)
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.SchedulesScheduleTypeRecurringWeekly{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.SchedulesScheduleTypeRecurringWeekly", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for SchedulesScheduleTypeRecurringWeekly ---
func unpackSchedulesScheduleTypeRecurringWeeklyListToSdk(ctx context.Context, list types.List) ([]objects.SchedulesScheduleTypeRecurringWeekly, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.SchedulesScheduleTypeRecurringWeekly")
	diags := diag.Diagnostics{}
	var data []models.SchedulesScheduleTypeRecurringWeekly
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.SchedulesScheduleTypeRecurringWeekly, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.SchedulesScheduleTypeRecurringWeekly{}.AttrTypes(), &item)
		unpacked, d := unpackSchedulesScheduleTypeRecurringWeeklyToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.SchedulesScheduleTypeRecurringWeekly", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for SchedulesScheduleTypeRecurringWeekly ---
func packSchedulesScheduleTypeRecurringWeeklyListFromSdk(ctx context.Context, sdks []objects.SchedulesScheduleTypeRecurringWeekly) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.SchedulesScheduleTypeRecurringWeekly")
	diags := diag.Diagnostics{}
	var data []models.SchedulesScheduleTypeRecurringWeekly

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.SchedulesScheduleTypeRecurringWeekly
		obj, d := packSchedulesScheduleTypeRecurringWeeklyFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.SchedulesScheduleTypeRecurringWeekly{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.SchedulesScheduleTypeRecurringWeekly", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.SchedulesScheduleTypeRecurringWeekly{}.AttrType(), data)
}
