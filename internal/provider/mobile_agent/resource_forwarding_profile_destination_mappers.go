package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/mobile_agent"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/mobile_agent"
)

// --- Unpacker for ForwardingProfileDestinations ---
func unpackForwardingProfileDestinationsToSdk(ctx context.Context, obj types.Object) (*mobile_agent.ForwardingProfileDestinations, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ForwardingProfileDestinations", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ForwardingProfileDestinations
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk mobile_agent.ForwardingProfileDestinations
	var d diag.Diagnostics

	// Handling Primitives
	if !model.Description.IsNull() && !model.Description.IsUnknown() {
		sdk.Description = model.Description.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Description", "value": *sdk.Description})
	}

	// Handling Lists
	if !model.Fqdn.IsNull() && !model.Fqdn.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field Fqdn")
		unpacked, d := unpackForwardingProfileDestinationFqdnEntryListToSdk(ctx, model.Fqdn)
		diags.Append(d...)
		sdk.Fqdn = unpacked
	}

	// Handling Primitives
	if !model.Id.IsNull() && !model.Id.IsUnknown() {
		sdk.Id = model.Id.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Id", "value": *sdk.Id})
	}

	// Handling Lists
	if !model.IpAddresses.IsNull() && !model.IpAddresses.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field IpAddresses")
		unpacked, d := unpackForwardingProfileDestinationIpEntryListToSdk(ctx, model.IpAddresses)
		diags.Append(d...)
		sdk.IpAddresses = unpacked
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ForwardingProfileDestinations", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ForwardingProfileDestinations ---
func packForwardingProfileDestinationsFromSdk(ctx context.Context, sdk mobile_agent.ForwardingProfileDestinations) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ForwardingProfileDestinations", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ForwardingProfileDestinations
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Description != nil {
		model.Description = basetypes.NewStringValue(*sdk.Description)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Description", "value": *sdk.Description})
	} else {
		model.Description = basetypes.NewStringNull()
	}
	// Handling Lists
	if sdk.Fqdn != nil {
		tflog.Debug(ctx, "Packing list of objects for field Fqdn")
		packed, d := packForwardingProfileDestinationFqdnEntryListFromSdk(ctx, sdk.Fqdn)
		diags.Append(d...)
		model.Fqdn = packed
	} else {
		model.Fqdn = basetypes.NewListNull(models.ForwardingProfileDestinationFqdnEntry{}.AttrType())
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
	if sdk.IpAddresses != nil {
		tflog.Debug(ctx, "Packing list of objects for field IpAddresses")
		packed, d := packForwardingProfileDestinationIpEntryListFromSdk(ctx, sdk.IpAddresses)
		diags.Append(d...)
		model.IpAddresses = packed
	} else {
		model.IpAddresses = basetypes.NewListNull(models.ForwardingProfileDestinationIpEntry{}.AttrType())
	}
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ForwardingProfileDestinations{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ForwardingProfileDestinations", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ForwardingProfileDestinations ---
func unpackForwardingProfileDestinationsListToSdk(ctx context.Context, list types.List) ([]mobile_agent.ForwardingProfileDestinations, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ForwardingProfileDestinations")
	diags := diag.Diagnostics{}
	var data []models.ForwardingProfileDestinations
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]mobile_agent.ForwardingProfileDestinations, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ForwardingProfileDestinations{}.AttrTypes(), &item)
		unpacked, d := unpackForwardingProfileDestinationsToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ForwardingProfileDestinations", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ForwardingProfileDestinations ---
func packForwardingProfileDestinationsListFromSdk(ctx context.Context, sdks []mobile_agent.ForwardingProfileDestinations) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ForwardingProfileDestinations")
	diags := diag.Diagnostics{}
	var data []models.ForwardingProfileDestinations

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ForwardingProfileDestinations
		obj, d := packForwardingProfileDestinationsFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ForwardingProfileDestinations{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ForwardingProfileDestinations", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ForwardingProfileDestinations{}.AttrType(), data)
}

// --- Unpacker for ForwardingProfileDestinationFqdnEntry ---
func unpackForwardingProfileDestinationFqdnEntryToSdk(ctx context.Context, obj types.Object) (*mobile_agent.ForwardingProfileDestinationFqdnEntry, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ForwardingProfileDestinationFqdnEntry", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ForwardingProfileDestinationFqdnEntry
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk mobile_agent.ForwardingProfileDestinationFqdnEntry
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	// Handling Primitives
	if !model.Port.IsNull() && !model.Port.IsUnknown() {
		val := int32(model.Port.ValueInt64())
		sdk.Port = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Port", "value": *sdk.Port})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ForwardingProfileDestinationFqdnEntry", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ForwardingProfileDestinationFqdnEntry ---
func packForwardingProfileDestinationFqdnEntryFromSdk(ctx context.Context, sdk mobile_agent.ForwardingProfileDestinationFqdnEntry) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ForwardingProfileDestinationFqdnEntry", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ForwardingProfileDestinationFqdnEntry
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	// Handling Primitives
	// Standard primitive packing
	if sdk.Port != nil {
		model.Port = basetypes.NewInt64Value(int64(*sdk.Port))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Port", "value": *sdk.Port})
	} else {
		model.Port = basetypes.NewInt64Null()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ForwardingProfileDestinationFqdnEntry{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ForwardingProfileDestinationFqdnEntry", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ForwardingProfileDestinationFqdnEntry ---
func unpackForwardingProfileDestinationFqdnEntryListToSdk(ctx context.Context, list types.List) ([]mobile_agent.ForwardingProfileDestinationFqdnEntry, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ForwardingProfileDestinationFqdnEntry")
	diags := diag.Diagnostics{}
	var data []models.ForwardingProfileDestinationFqdnEntry
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]mobile_agent.ForwardingProfileDestinationFqdnEntry, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ForwardingProfileDestinationFqdnEntry{}.AttrTypes(), &item)
		unpacked, d := unpackForwardingProfileDestinationFqdnEntryToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ForwardingProfileDestinationFqdnEntry", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ForwardingProfileDestinationFqdnEntry ---
func packForwardingProfileDestinationFqdnEntryListFromSdk(ctx context.Context, sdks []mobile_agent.ForwardingProfileDestinationFqdnEntry) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ForwardingProfileDestinationFqdnEntry")
	diags := diag.Diagnostics{}
	var data []models.ForwardingProfileDestinationFqdnEntry

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ForwardingProfileDestinationFqdnEntry
		obj, d := packForwardingProfileDestinationFqdnEntryFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ForwardingProfileDestinationFqdnEntry{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ForwardingProfileDestinationFqdnEntry", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ForwardingProfileDestinationFqdnEntry{}.AttrType(), data)
}

// --- Unpacker for ForwardingProfileDestinationIpEntry ---
func unpackForwardingProfileDestinationIpEntryToSdk(ctx context.Context, obj types.Object) (*mobile_agent.ForwardingProfileDestinationIpEntry, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ForwardingProfileDestinationIpEntry", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ForwardingProfileDestinationIpEntry
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk mobile_agent.ForwardingProfileDestinationIpEntry
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	// Handling Primitives
	if !model.Port.IsNull() && !model.Port.IsUnknown() {
		val := int32(model.Port.ValueInt64())
		sdk.Port = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Port", "value": *sdk.Port})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ForwardingProfileDestinationIpEntry", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ForwardingProfileDestinationIpEntry ---
func packForwardingProfileDestinationIpEntryFromSdk(ctx context.Context, sdk mobile_agent.ForwardingProfileDestinationIpEntry) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ForwardingProfileDestinationIpEntry", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ForwardingProfileDestinationIpEntry
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	// Handling Primitives
	// Standard primitive packing
	if sdk.Port != nil {
		model.Port = basetypes.NewInt64Value(int64(*sdk.Port))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Port", "value": *sdk.Port})
	} else {
		model.Port = basetypes.NewInt64Null()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ForwardingProfileDestinationIpEntry{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ForwardingProfileDestinationIpEntry", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ForwardingProfileDestinationIpEntry ---
func unpackForwardingProfileDestinationIpEntryListToSdk(ctx context.Context, list types.List) ([]mobile_agent.ForwardingProfileDestinationIpEntry, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ForwardingProfileDestinationIpEntry")
	diags := diag.Diagnostics{}
	var data []models.ForwardingProfileDestinationIpEntry
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]mobile_agent.ForwardingProfileDestinationIpEntry, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ForwardingProfileDestinationIpEntry{}.AttrTypes(), &item)
		unpacked, d := unpackForwardingProfileDestinationIpEntryToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ForwardingProfileDestinationIpEntry", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ForwardingProfileDestinationIpEntry ---
func packForwardingProfileDestinationIpEntryListFromSdk(ctx context.Context, sdks []mobile_agent.ForwardingProfileDestinationIpEntry) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ForwardingProfileDestinationIpEntry")
	diags := diag.Diagnostics{}
	var data []models.ForwardingProfileDestinationIpEntry

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ForwardingProfileDestinationIpEntry
		obj, d := packForwardingProfileDestinationIpEntryFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ForwardingProfileDestinationIpEntry{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ForwardingProfileDestinationIpEntry", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ForwardingProfileDestinationIpEntry{}.AttrType(), data)
}
