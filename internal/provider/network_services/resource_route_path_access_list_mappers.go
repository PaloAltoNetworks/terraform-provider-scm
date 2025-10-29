package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/network_services"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/network_services"
)

// --- Unpacker for RoutePathAccessLists ---
func unpackRoutePathAccessListsToSdk(ctx context.Context, obj types.Object) (*network_services.RoutePathAccessLists, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.RoutePathAccessLists", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.RoutePathAccessLists
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.RoutePathAccessLists
	var d diag.Diagnostics

	// Handling Lists
	if !model.AspathEntry.IsNull() && !model.AspathEntry.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field AspathEntry")
		unpacked, d := unpackRoutePathAccessListsAspathEntryInnerListToSdk(ctx, model.AspathEntry)
		diags.Append(d...)
		sdk.AspathEntry = unpacked
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
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	// Handling Primitives
	if !model.Snippet.IsNull() && !model.Snippet.IsUnknown() {
		sdk.Snippet = model.Snippet.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.RoutePathAccessLists", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for RoutePathAccessLists ---
func packRoutePathAccessListsFromSdk(ctx context.Context, sdk network_services.RoutePathAccessLists) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.RoutePathAccessLists", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.RoutePathAccessLists
	var d diag.Diagnostics
	// Handling Lists
	if sdk.AspathEntry != nil {
		tflog.Debug(ctx, "Packing list of objects for field AspathEntry")
		packed, d := packRoutePathAccessListsAspathEntryInnerListFromSdk(ctx, sdk.AspathEntry)
		diags.Append(d...)
		model.AspathEntry = packed
	} else {
		model.AspathEntry = basetypes.NewListNull(models.RoutePathAccessListsAspathEntryInner{}.AttrType())
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

	obj, d := types.ObjectValueFrom(ctx, models.RoutePathAccessLists{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.RoutePathAccessLists", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for RoutePathAccessLists ---
func unpackRoutePathAccessListsListToSdk(ctx context.Context, list types.List) ([]network_services.RoutePathAccessLists, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.RoutePathAccessLists")
	diags := diag.Diagnostics{}
	var data []models.RoutePathAccessLists
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.RoutePathAccessLists, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.RoutePathAccessLists{}.AttrTypes(), &item)
		unpacked, d := unpackRoutePathAccessListsToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.RoutePathAccessLists", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for RoutePathAccessLists ---
func packRoutePathAccessListsListFromSdk(ctx context.Context, sdks []network_services.RoutePathAccessLists) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.RoutePathAccessLists")
	diags := diag.Diagnostics{}
	var data []models.RoutePathAccessLists

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.RoutePathAccessLists
		obj, d := packRoutePathAccessListsFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.RoutePathAccessLists{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.RoutePathAccessLists", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.RoutePathAccessLists{}.AttrType(), data)
}

// --- Unpacker for RoutePathAccessListsAspathEntryInner ---
func unpackRoutePathAccessListsAspathEntryInnerToSdk(ctx context.Context, obj types.Object) (*network_services.RoutePathAccessListsAspathEntryInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.RoutePathAccessListsAspathEntryInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.RoutePathAccessListsAspathEntryInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.RoutePathAccessListsAspathEntryInner
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Action.IsNull() && !model.Action.IsUnknown() {
		sdk.Action = model.Action.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Action", "value": *sdk.Action})
	}

	// Handling Primitives
	if !model.AspathRegex.IsNull() && !model.AspathRegex.IsUnknown() {
		sdk.AspathRegex = model.AspathRegex.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "AspathRegex", "value": *sdk.AspathRegex})
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		val := int32(model.Name.ValueInt64())
		sdk.Name = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.RoutePathAccessListsAspathEntryInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for RoutePathAccessListsAspathEntryInner ---
func packRoutePathAccessListsAspathEntryInnerFromSdk(ctx context.Context, sdk network_services.RoutePathAccessListsAspathEntryInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.RoutePathAccessListsAspathEntryInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.RoutePathAccessListsAspathEntryInner
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
	if sdk.AspathRegex != nil {
		model.AspathRegex = basetypes.NewStringValue(*sdk.AspathRegex)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "AspathRegex", "value": *sdk.AspathRegex})
	} else {
		model.AspathRegex = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Name != nil {
		model.Name = basetypes.NewInt64Value(int64(*sdk.Name))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
	} else {
		model.Name = basetypes.NewInt64Null()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.RoutePathAccessListsAspathEntryInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.RoutePathAccessListsAspathEntryInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for RoutePathAccessListsAspathEntryInner ---
func unpackRoutePathAccessListsAspathEntryInnerListToSdk(ctx context.Context, list types.List) ([]network_services.RoutePathAccessListsAspathEntryInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.RoutePathAccessListsAspathEntryInner")
	diags := diag.Diagnostics{}
	var data []models.RoutePathAccessListsAspathEntryInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.RoutePathAccessListsAspathEntryInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.RoutePathAccessListsAspathEntryInner{}.AttrTypes(), &item)
		unpacked, d := unpackRoutePathAccessListsAspathEntryInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.RoutePathAccessListsAspathEntryInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for RoutePathAccessListsAspathEntryInner ---
func packRoutePathAccessListsAspathEntryInnerListFromSdk(ctx context.Context, sdks []network_services.RoutePathAccessListsAspathEntryInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.RoutePathAccessListsAspathEntryInner")
	diags := diag.Diagnostics{}
	var data []models.RoutePathAccessListsAspathEntryInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.RoutePathAccessListsAspathEntryInner
		obj, d := packRoutePathAccessListsAspathEntryInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.RoutePathAccessListsAspathEntryInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.RoutePathAccessListsAspathEntryInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.RoutePathAccessListsAspathEntryInner{}.AttrType(), data)
}
