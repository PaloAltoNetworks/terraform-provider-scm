package provider

/*

import (
	"context"
	"fmt"
	"reflect"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/objects"
	"github.com/paloaltonetworks/scm-go/generated/objects"
)










// --- Unpacker for AutoTagActions ---
func unpackAutoTagActionsToSdk(ctx context.Context, obj types.Object) (*objects.AutoTagActions, diag.Diagnostics) {
    tflog.Debug(ctx, "Entering unpack helper for models.AutoTagActions", map[string]interface{}{"tf_object": obj})
    diags := diag.Diagnostics{}
    var model models.AutoTagActions
    diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
    if diags.HasError() {
        tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
        return nil, diags
    }
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

    var sdk objects.AutoTagActions
    var d diag.Diagnostics

    // Handling Lists
    if !model.Actions.IsNull() && !model.Actions.IsUnknown() {
        tflog.Debug(ctx, "Unpacking list of objects for field Actions")
        unpacked, d := unpackAutoTagActionsActionsInnerListToSdk(ctx, model.Actions)
        diags.Append(d...)
        sdk.Actions = unpacked
    }

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
    if !model.Filter.IsNull() && !model.Filter.IsUnknown() {
        sdk.Filter = model.Filter.ValueString()
        tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Filter", "value": sdk.Filter})
    }

    // Handling Primitives
    if !model.Folder.IsNull() && !model.Folder.IsUnknown() {
        sdk.Folder = model.Folder.ValueStringPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Folder", "value": *sdk.Folder})
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

    // Handling Primitives
    if !model.Quarantine.IsNull() && !model.Quarantine.IsUnknown() {
        sdk.Quarantine = model.Quarantine.ValueBoolPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Quarantine", "value": *sdk.Quarantine})
    }

    // Handling Primitives
    if !model.SendToPanorama.IsNull() && !model.SendToPanorama.IsUnknown() {
        sdk.SendToPanorama = model.SendToPanorama.ValueBoolPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "SendToPanorama", "value": *sdk.SendToPanorama})
    }

    // Handling Primitives
    if !model.Snippet.IsNull() && !model.Snippet.IsUnknown() {
        sdk.Snippet = model.Snippet.ValueStringPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
    }

diags.Append(d...)

    tflog.Debug(ctx, "Exiting unpack helper for models.AutoTagActions", map[string]interface{}{"has_errors": diags.HasError()})
    return &sdk, diags

}

// --- Packer for AutoTagActions ---
func packAutoTagActionsFromSdk(ctx context.Context, sdk objects.AutoTagActions) (types.Object, diag.Diagnostics) {
    tflog.Debug(ctx, "Entering pack helper for models.AutoTagActions", map[string]interface{}{"sdk_struct": sdk})
    diags := diag.Diagnostics{}
    var model models.AutoTagActions
    var d diag.Diagnostics
    // Handling Lists
    if sdk.Actions != nil {
        tflog.Debug(ctx, "Packing list of objects for field Actions")
        packed, d := packAutoTagActionsActionsInnerListFromSdk(ctx, sdk.Actions)
        diags.Append(d...)
        model.Actions = packed
    } else {
        model.Actions = basetypes.NewListNull(models.AutoTagActionsActionsInner{}.AttrType())
    }
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
    model.Filter = basetypes.NewStringValue(sdk.Filter)
    tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Filter", "value": sdk.Filter})
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
    model.LogType = basetypes.NewStringValue(sdk.LogType)
    tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "LogType", "value": sdk.LogType})
    // Handling Primitives
    // Standard primitive packing
    model.Name = basetypes.NewStringValue(sdk.Name)
    tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
    // Handling Primitives
    // Standard primitive packing
    if sdk.Quarantine != nil {
        model.Quarantine = basetypes.NewBoolValue(*sdk.Quarantine)
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Quarantine", "value": *sdk.Quarantine})
    } else {
        model.Quarantine = basetypes.NewBoolNull()
    }
    // Handling Primitives
    // Standard primitive packing
    if sdk.SendToPanorama != nil {
        model.SendToPanorama = basetypes.NewBoolValue(*sdk.SendToPanorama)
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "SendToPanorama", "value": *sdk.SendToPanorama})
    } else {
        model.SendToPanorama = basetypes.NewBoolNull()
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

    obj, d := types.ObjectValueFrom(ctx, models.AutoTagActions{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
    diags.Append(d...)
    tflog.Debug(ctx, "Exiting pack helper for models.AutoTagActions", map[string]interface{}{"has_errors": diags.HasError()})
    return obj, diags

}

// --- List Unpacker for AutoTagActions ---
func unpackAutoTagActionsListToSdk(ctx context.Context, list types.List) ([]objects.AutoTagActions, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.AutoTagActions")
	diags := diag.Diagnostics{}
	var data []models.AutoTagActions
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.AutoTagActions, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.AutoTagActions{}.AttrTypes(), &item)
		unpacked, d := unpackAutoTagActionsToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.AutoTagActions", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for AutoTagActions ---
func packAutoTagActionsListFromSdk(ctx context.Context, sdks []objects.AutoTagActions) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.AutoTagActions")
	diags := diag.Diagnostics{}
	var data []models.AutoTagActions

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.AutoTagActions
		obj, d := packAutoTagActionsFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.AutoTagActions{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.AutoTagActions", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.AutoTagActions{}.AttrType(), data)
}

// --- Unpacker for AutoTagActionsActionsInner ---
func unpackAutoTagActionsActionsInnerToSdk(ctx context.Context, obj types.Object) (*objects.AutoTagActionsActionsInner, diag.Diagnostics) {
    tflog.Debug(ctx, "Entering unpack helper for models.AutoTagActionsActionsInner", map[string]interface{}{"tf_object": obj})
    diags := diag.Diagnostics{}
    var model models.AutoTagActionsActionsInner
    diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
    if diags.HasError() {
        tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
        return nil, diags
    }
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

    var sdk objects.AutoTagActionsActionsInner
    var d diag.Diagnostics
    // Handling Primitives
    if !model.Name.IsNull() && !model.Name.IsUnknown() {
        sdk.Name = model.Name.ValueString()
        tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
    }

    // Handling Objects
    if !model.Type.IsNull() && !model.Type.IsUnknown() {
        tflog.Debug(ctx, "Unpacking nested object for field Type")
        unpacked, d := unpackAutoTagActionsActionsInnerTypeToSdk(ctx, model.Type)
        diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Type"})
		}
        if unpacked != nil {
            sdk.Type = *unpacked
        }
    }

diags.Append(d...)

    tflog.Debug(ctx, "Exiting unpack helper for models.AutoTagActionsActionsInner", map[string]interface{}{"has_errors": diags.HasError()})
    return &sdk, diags

}

// --- Packer for AutoTagActionsActionsInner ---
func packAutoTagActionsActionsInnerFromSdk(ctx context.Context, sdk objects.AutoTagActionsActionsInner) (types.Object, diag.Diagnostics) {
    tflog.Debug(ctx, "Entering pack helper for models.AutoTagActionsActionsInner", map[string]interface{}{"sdk_struct": sdk})
    diags := diag.Diagnostics{}
    var model models.AutoTagActionsActionsInner
    var d diag.Diagnostics
    // Handling Primitives
    // Standard primitive packing
    model.Name = basetypes.NewStringValue(sdk.Name)
    tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
    // Handling Objects
    // This is a regular nested object that has its own packer.
    // Logic for non-pointer / value-type nested objects
    if !reflect.ValueOf(sdk.Type).IsZero() {
        tflog.Debug(ctx, "Packing nested object for field Type")
        packed, d := packAutoTagActionsActionsInnerTypeFromSdk(ctx, sdk.Type)
        diags.Append(d...)
        model.Type = packed
    } else {
        model.Type = basetypes.NewObjectNull(models.AutoTagActionsActionsInnerType{}.AttrTypes())
    }
diags.Append(d...)

    obj, d := types.ObjectValueFrom(ctx, models.AutoTagActionsActionsInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
    diags.Append(d...)
    tflog.Debug(ctx, "Exiting pack helper for models.AutoTagActionsActionsInner", map[string]interface{}{"has_errors": diags.HasError()})
    return obj, diags

}

// --- List Unpacker for AutoTagActionsActionsInner ---
func unpackAutoTagActionsActionsInnerListToSdk(ctx context.Context, list types.List) ([]objects.AutoTagActionsActionsInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.AutoTagActionsActionsInner")
	diags := diag.Diagnostics{}
	var data []models.AutoTagActionsActionsInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.AutoTagActionsActionsInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.AutoTagActionsActionsInner{}.AttrTypes(), &item)
		unpacked, d := unpackAutoTagActionsActionsInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.AutoTagActionsActionsInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for AutoTagActionsActionsInner ---
func packAutoTagActionsActionsInnerListFromSdk(ctx context.Context, sdks []objects.AutoTagActionsActionsInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.AutoTagActionsActionsInner")
	diags := diag.Diagnostics{}
	var data []models.AutoTagActionsActionsInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.AutoTagActionsActionsInner
		obj, d := packAutoTagActionsActionsInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.AutoTagActionsActionsInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.AutoTagActionsActionsInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.AutoTagActionsActionsInner{}.AttrType(), data)
}

// --- Unpacker for AutoTagActionsActionsInnerType ---
func unpackAutoTagActionsActionsInnerTypeToSdk(ctx context.Context, obj types.Object) (*objects.AutoTagActionsActionsInnerType, diag.Diagnostics) {
    tflog.Debug(ctx, "Entering unpack helper for models.AutoTagActionsActionsInnerType", map[string]interface{}{"tf_object": obj})
    diags := diag.Diagnostics{}
    var model models.AutoTagActionsActionsInnerType
    diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
    if diags.HasError() {
        tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
        return nil, diags
    }
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

    var sdk objects.AutoTagActionsActionsInnerType
    var d diag.Diagnostics
    // Handling Objects
    if !model.Tagging.IsNull() && !model.Tagging.IsUnknown() {
        tflog.Debug(ctx, "Unpacking nested object for field Tagging")
        unpacked, d := unpackAutoTagActionsActionsInnerTypeTaggingToSdk(ctx, model.Tagging)
        diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Tagging"})
		}
        if unpacked != nil {
            sdk.Tagging = *unpacked
        }
    }

diags.Append(d...)

    tflog.Debug(ctx, "Exiting unpack helper for models.AutoTagActionsActionsInnerType", map[string]interface{}{"has_errors": diags.HasError()})
    return &sdk, diags

}

// --- Packer for AutoTagActionsActionsInnerType ---
func packAutoTagActionsActionsInnerTypeFromSdk(ctx context.Context, sdk objects.AutoTagActionsActionsInnerType) (types.Object, diag.Diagnostics) {
    tflog.Debug(ctx, "Entering pack helper for models.AutoTagActionsActionsInnerType", map[string]interface{}{"sdk_struct": sdk})
    diags := diag.Diagnostics{}
    var model models.AutoTagActionsActionsInnerType
    var d diag.Diagnostics
    // Handling Objects
    // This is a regular nested object that has its own packer.
    // Logic for non-pointer / value-type nested objects
    if !reflect.ValueOf(sdk.Tagging).IsZero() {
        tflog.Debug(ctx, "Packing nested object for field Tagging")
        packed, d := packAutoTagActionsActionsInnerTypeTaggingFromSdk(ctx, sdk.Tagging)
        diags.Append(d...)
        model.Tagging = packed
    } else {
        model.Tagging = basetypes.NewObjectNull(models.AutoTagActionsActionsInnerTypeTagging{}.AttrTypes())
    }
diags.Append(d...)

    obj, d := types.ObjectValueFrom(ctx, models.AutoTagActionsActionsInnerType{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
    diags.Append(d...)
    tflog.Debug(ctx, "Exiting pack helper for models.AutoTagActionsActionsInnerType", map[string]interface{}{"has_errors": diags.HasError()})
    return obj, diags

}

// --- List Unpacker for AutoTagActionsActionsInnerType ---
func unpackAutoTagActionsActionsInnerTypeListToSdk(ctx context.Context, list types.List) ([]objects.AutoTagActionsActionsInnerType, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.AutoTagActionsActionsInnerType")
	diags := diag.Diagnostics{}
	var data []models.AutoTagActionsActionsInnerType
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.AutoTagActionsActionsInnerType, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.AutoTagActionsActionsInnerType{}.AttrTypes(), &item)
		unpacked, d := unpackAutoTagActionsActionsInnerTypeToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.AutoTagActionsActionsInnerType", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for AutoTagActionsActionsInnerType ---
func packAutoTagActionsActionsInnerTypeListFromSdk(ctx context.Context, sdks []objects.AutoTagActionsActionsInnerType) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.AutoTagActionsActionsInnerType")
	diags := diag.Diagnostics{}
	var data []models.AutoTagActionsActionsInnerType

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.AutoTagActionsActionsInnerType
		obj, d := packAutoTagActionsActionsInnerTypeFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.AutoTagActionsActionsInnerType{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.AutoTagActionsActionsInnerType", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.AutoTagActionsActionsInnerType{}.AttrType(), data)
}

// --- Unpacker for AutoTagActionsActionsInnerTypeTagging ---
func unpackAutoTagActionsActionsInnerTypeTaggingToSdk(ctx context.Context, obj types.Object) (*objects.AutoTagActionsActionsInnerTypeTagging, diag.Diagnostics) {
    tflog.Debug(ctx, "Entering unpack helper for models.AutoTagActionsActionsInnerTypeTagging", map[string]interface{}{"tf_object": obj})
    diags := diag.Diagnostics{}
    var model models.AutoTagActionsActionsInnerTypeTagging
    diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
    if diags.HasError() {
        tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
        return nil, diags
    }
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

    var sdk objects.AutoTagActionsActionsInnerTypeTagging
    var d diag.Diagnostics
    // Handling Primitives
    if !model.Action.IsNull() && !model.Action.IsUnknown() {
        sdk.Action = model.Action.ValueString()
        tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Action", "value": sdk.Action})
    }

    // Handling Lists
    if !model.Tags.IsNull() && !model.Tags.IsUnknown() {
        tflog.Debug(ctx, "Unpacking list of primitives for field Tags")
        diags.Append(model.Tags.ElementsAs(ctx, &sdk.Tags, false)...)
    }

    // Handling Primitives
    if !model.Target.IsNull() && !model.Target.IsUnknown() {
        sdk.Target = model.Target.ValueString()
        tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Target", "value": sdk.Target})
    }

    // Handling Primitives
    if !model.Timeout.IsNull() && !model.Timeout.IsUnknown() {
        val := float32(model.Timeout.ValueFloat64())
        sdk.Timeout = &val
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Timeout", "value": *sdk.Timeout})
    }

diags.Append(d...)

    tflog.Debug(ctx, "Exiting unpack helper for models.AutoTagActionsActionsInnerTypeTagging", map[string]interface{}{"has_errors": diags.HasError()})
    return &sdk, diags

}

// --- Packer for AutoTagActionsActionsInnerTypeTagging ---
func packAutoTagActionsActionsInnerTypeTaggingFromSdk(ctx context.Context, sdk objects.AutoTagActionsActionsInnerTypeTagging) (types.Object, diag.Diagnostics) {
    tflog.Debug(ctx, "Entering pack helper for models.AutoTagActionsActionsInnerTypeTagging", map[string]interface{}{"sdk_struct": sdk})
    diags := diag.Diagnostics{}
    var model models.AutoTagActionsActionsInnerTypeTagging
    var d diag.Diagnostics
    // Handling Primitives
    // Standard primitive packing
    model.Action = basetypes.NewStringValue(sdk.Action)
    tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Action", "value": sdk.Action})
    // Handling Lists
    if sdk.Tags != nil {
        tflog.Debug(ctx, "Packing list of primitives for field Tags")
        var d diag.Diagnostics
        // This logic now dynamically determines the element type based on the SDK's Go type.
        var elemType attr.Type = basetypes.StringType{} // Default to string
        model.Tags, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Tags)
        diags.Append(d...)
    } else {
        // This logic now creates a correctly typed null list.
        var elemType attr.Type = basetypes.StringType{} // Default to string
        model.Tags = basetypes.NewListNull(elemType)
    }
    // Handling Primitives
    // Standard primitive packing
    model.Target = basetypes.NewStringValue(sdk.Target)
    tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Target", "value": sdk.Target})
    // Handling Primitives
    // Standard primitive packing
    if sdk.Timeout != nil {
        model.Timeout = basetypes.NewFloat64Value(float64(*sdk.Timeout))
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Timeout", "value": *sdk.Timeout})
    } else {
        model.Timeout = basetypes.NewFloat64Null()
    }
diags.Append(d...)

    obj, d := types.ObjectValueFrom(ctx, models.AutoTagActionsActionsInnerTypeTagging{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
    diags.Append(d...)
    tflog.Debug(ctx, "Exiting pack helper for models.AutoTagActionsActionsInnerTypeTagging", map[string]interface{}{"has_errors": diags.HasError()})
    return obj, diags

}

// --- List Unpacker for AutoTagActionsActionsInnerTypeTagging ---
func unpackAutoTagActionsActionsInnerTypeTaggingListToSdk(ctx context.Context, list types.List) ([]objects.AutoTagActionsActionsInnerTypeTagging, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.AutoTagActionsActionsInnerTypeTagging")
	diags := diag.Diagnostics{}
	var data []models.AutoTagActionsActionsInnerTypeTagging
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.AutoTagActionsActionsInnerTypeTagging, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.AutoTagActionsActionsInnerTypeTagging{}.AttrTypes(), &item)
		unpacked, d := unpackAutoTagActionsActionsInnerTypeTaggingToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.AutoTagActionsActionsInnerTypeTagging", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for AutoTagActionsActionsInnerTypeTagging ---
func packAutoTagActionsActionsInnerTypeTaggingListFromSdk(ctx context.Context, sdks []objects.AutoTagActionsActionsInnerTypeTagging) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.AutoTagActionsActionsInnerTypeTagging")
	diags := diag.Diagnostics{}
	var data []models.AutoTagActionsActionsInnerTypeTagging

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.AutoTagActionsActionsInnerTypeTagging
		obj, d := packAutoTagActionsActionsInnerTypeTaggingFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.AutoTagActionsActionsInnerTypeTagging{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.AutoTagActionsActionsInnerTypeTagging", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.AutoTagActionsActionsInnerTypeTagging{}.AttrType(), data)
}


*/
