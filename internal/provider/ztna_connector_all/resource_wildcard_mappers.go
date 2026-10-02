package provider

import (
	"context"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/ztna_connector_all"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/ztna_connector_all"
)

// --- Unpacker for Wildcards ---
func unpackWildcardsToSdk(ctx context.Context, obj types.Object) (*ztna_connector_all.Wildcards, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.Wildcards", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.Wildcards
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk ztna_connector_all.Wildcards
	var d diag.Diagnostics

	// Handling Primitives
	if !model.AppEnabled.IsNull() && !model.AppEnabled.IsUnknown() {
		sdk.AppEnabled = model.AppEnabled.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "AppEnabled", "value": *sdk.AppEnabled})
	}

	// Handling Typeless Objects
	if !model.Applications.IsNull() && !model.Applications.IsUnknown() {
		tflog.Debug(ctx, "Unpacking typeless object for field Applications")
		sdk.Applications = make(map[string]interface{})
	}

	// Handling Primitives
	if !model.CreatedTime.IsNull() && !model.CreatedTime.IsUnknown() {
		sdk.CreatedTime = model.CreatedTime.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "CreatedTime", "value": *sdk.CreatedTime})
	}

	// Handling Primitives
	if !model.Description.IsNull() && !model.Description.IsUnknown() {
		sdk.Description = model.Description.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Description", "value": *sdk.Description})
	}

	// Handling Primitives
	if !model.EnablePolicy.IsNull() && !model.EnablePolicy.IsUnknown() {
		sdk.EnablePolicy = model.EnablePolicy.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "EnablePolicy", "value": *sdk.EnablePolicy})
	}

	// Handling Primitives
	if !model.Fqdn.IsNull() && !model.Fqdn.IsUnknown() {
		sdk.Fqdn = model.Fqdn.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Fqdn", "value": sdk.Fqdn})
	}

	// Handling Primitives
	if !model.Group.IsNull() && !model.Group.IsUnknown() {
		sdk.Group = model.Group.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Group", "value": sdk.Group})
	}

	// Handling Primitives
	if !model.IcmpAllowed.IsNull() && !model.IcmpAllowed.IsUnknown() {
		sdk.IcmpAllowed = model.IcmpAllowed.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "IcmpAllowed", "value": *sdk.IcmpAllowed})
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
	if !model.Oid.IsNull() && !model.Oid.IsUnknown() {
		sdk.Oid = model.Oid.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Oid", "value": *sdk.Oid})
	}

	// Handling Primitives
	if !model.ProbePort.IsNull() && !model.ProbePort.IsUnknown() {
		sdk.ProbePort = model.ProbePort.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "ProbePort", "value": *sdk.ProbePort})
	}

	// Handling Primitives
	if !model.ProbeType.IsNull() && !model.ProbeType.IsUnknown() {
		sdk.ProbeType = model.ProbeType.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "ProbeType", "value": *sdk.ProbeType})
	}

	// Handling Primitives
	if !model.TcpPort.IsNull() && !model.TcpPort.IsUnknown() {
		sdk.TcpPort = model.TcpPort.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "TcpPort", "value": *sdk.TcpPort})
	}

	// Handling Primitives
	if !model.UdpPort.IsNull() && !model.UdpPort.IsUnknown() {
		sdk.UdpPort = model.UdpPort.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "UdpPort", "value": *sdk.UdpPort})
	}

	// Handling Primitives
	if !model.UpdatedTime.IsNull() && !model.UpdatedTime.IsUnknown() {
		sdk.UpdatedTime = model.UpdatedTime.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "UpdatedTime", "value": *sdk.UpdatedTime})
	}

	// Handling Primitives
	if !model.UseDcIp.IsNull() && !model.UseDcIp.IsUnknown() {
		sdk.UseDcIp = model.UseDcIp.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "UseDcIp", "value": *sdk.UseDcIp})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.Wildcards", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for Wildcards ---
func packWildcardsFromSdk(ctx context.Context, sdk ztna_connector_all.Wildcards) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.Wildcards", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.Wildcards
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.AppEnabled != nil {
		model.AppEnabled = basetypes.NewBoolValue(*sdk.AppEnabled)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "AppEnabled", "value": *sdk.AppEnabled})
	} else {
		model.AppEnabled = basetypes.NewBoolNull()
	}
	// Handling Objects
	// This is a marker object (e.g. CHAP: {}). We just need to create an empty, non-null object.
	if sdk.Applications != nil && !reflect.ValueOf(sdk.Applications).IsNil() {
		tflog.Debug(ctx, "Packing typeless object for field Applications")
		var d diag.Diagnostics
		// Create an empty object with no attributes, which signifies its presence.
		model.Applications, d = basetypes.NewObjectValue(map[string]attr.Type{}, map[string]attr.Value{})
		diags.Append(d...)
	} else {
		// Since this field is part of a oneOf, being nil means it's not selected.
		// We make the object null with an empty attribute map.
		model.Applications = basetypes.NewObjectNull(map[string]attr.Type{})
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.CreatedTime != nil {
		model.CreatedTime = basetypes.NewStringValue(*sdk.CreatedTime)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "CreatedTime", "value": *sdk.CreatedTime})
	} else {
		model.CreatedTime = basetypes.NewStringNull()
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
	if sdk.EnablePolicy != nil {
		model.EnablePolicy = basetypes.NewBoolValue(*sdk.EnablePolicy)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "EnablePolicy", "value": *sdk.EnablePolicy})
	} else {
		model.EnablePolicy = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	model.Fqdn = basetypes.NewStringValue(sdk.Fqdn)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Fqdn", "value": sdk.Fqdn})
	// Handling Primitives
	// Standard primitive packing
	model.Group = basetypes.NewStringValue(sdk.Group)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Group", "value": sdk.Group})
	// Handling Primitives
	// Standard primitive packing
	if sdk.IcmpAllowed != nil {
		model.IcmpAllowed = basetypes.NewBoolValue(*sdk.IcmpAllowed)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "IcmpAllowed", "value": *sdk.IcmpAllowed})
	} else {
		model.IcmpAllowed = basetypes.NewBoolNull()
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
	if sdk.Oid != nil {
		model.Oid = basetypes.NewStringValue(*sdk.Oid)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Oid", "value": *sdk.Oid})
	} else {
		if sdk.Id != nil {
			model.Oid = basetypes.NewStringValue(*sdk.Id)
			tflog.Debug(ctx, "Packed primitive pointer from fallback field", map[string]interface{}{"field": "Oid", "fallback": "Id"})
		} else {
			model.Oid = basetypes.NewStringNull()
		}
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.ProbePort != nil {
		model.ProbePort = basetypes.NewStringValue(*sdk.ProbePort)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "ProbePort", "value": *sdk.ProbePort})
	} else {
		model.ProbePort = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.ProbeType != nil {
		model.ProbeType = basetypes.NewStringValue(*sdk.ProbeType)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "ProbeType", "value": *sdk.ProbeType})
	} else {
		model.ProbeType = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.TcpPort != nil {
		model.TcpPort = basetypes.NewStringValue(*sdk.TcpPort)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "TcpPort", "value": *sdk.TcpPort})
	} else {
		model.TcpPort = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.UdpPort != nil {
		model.UdpPort = basetypes.NewStringValue(*sdk.UdpPort)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "UdpPort", "value": *sdk.UdpPort})
	} else {
		model.UdpPort = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.UpdatedTime != nil {
		model.UpdatedTime = basetypes.NewStringValue(*sdk.UpdatedTime)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "UpdatedTime", "value": *sdk.UpdatedTime})
	} else {
		model.UpdatedTime = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.UseDcIp != nil {
		model.UseDcIp = basetypes.NewBoolValue(*sdk.UseDcIp)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "UseDcIp", "value": *sdk.UseDcIp})
	} else {
		model.UseDcIp = basetypes.NewBoolNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.Wildcards{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.Wildcards", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for Wildcards ---
func unpackWildcardsListToSdk(ctx context.Context, list types.List) ([]ztna_connector_all.Wildcards, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.Wildcards")
	diags := diag.Diagnostics{}
	var data []models.Wildcards
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]ztna_connector_all.Wildcards, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.Wildcards{}.AttrTypes(), &item)
		unpacked, d := unpackWildcardsToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.Wildcards", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for Wildcards ---
func packWildcardsListFromSdk(ctx context.Context, sdks []ztna_connector_all.Wildcards) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.Wildcards")
	diags := diag.Diagnostics{}
	var data []models.Wildcards

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.Wildcards
		obj, d := packWildcardsFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.Wildcards{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.Wildcards", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.Wildcards{}.AttrType(), data)
}
