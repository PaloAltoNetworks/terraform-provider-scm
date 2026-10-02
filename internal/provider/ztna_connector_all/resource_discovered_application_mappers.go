package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/ztna_connector_all"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/ztna_connector_all"
)

// --- Unpacker for DiscoveredApplications ---
func unpackDiscoveredApplicationsToSdk(ctx context.Context, obj types.Object) (*ztna_connector_all.DiscoveredApplications, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.DiscoveredApplications", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.DiscoveredApplications
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk ztna_connector_all.DiscoveredApplications
	var d diag.Diagnostics

	// Handling Lists
	if !model.Applications.IsNull() && !model.Applications.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field Applications")
		unpacked, d := unpackDiscoveredApplicationsApplicationsInnerListToSdk(ctx, model.Applications)
		diags.Append(d...)
		sdk.Applications = unpacked
	}

	// Handling Primitives
	if !model.CieTenantId.IsNull() && !model.CieTenantId.IsUnknown() {
		sdk.CieTenantId = model.CieTenantId.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "CieTenantId", "value": *sdk.CieTenantId})
	}

	// Handling Primitives
	if !model.Count.IsNull() && !model.Count.IsUnknown() {
		val := float32(model.Count.ValueFloat64())
		sdk.Count = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Count", "value": *sdk.Count})
	}

	// Handling Primitives
	if !model.TenantId.IsNull() && !model.TenantId.IsUnknown() {
		sdk.TenantId = model.TenantId.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "TenantId", "value": *sdk.TenantId})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.DiscoveredApplications", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for DiscoveredApplications ---
func packDiscoveredApplicationsFromSdk(ctx context.Context, sdk ztna_connector_all.DiscoveredApplications) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.DiscoveredApplications", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.DiscoveredApplications
	var d diag.Diagnostics
	// Handling Lists
	if sdk.Applications != nil {
		tflog.Debug(ctx, "Packing list of objects for field Applications")
		packed, d := packDiscoveredApplicationsApplicationsInnerListFromSdk(ctx, sdk.Applications)
		diags.Append(d...)
		model.Applications = packed
	} else {
		model.Applications = basetypes.NewListNull(models.DiscoveredApplicationsApplicationsInner{}.AttrType())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.CieTenantId != nil {
		model.CieTenantId = basetypes.NewStringValue(*sdk.CieTenantId)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "CieTenantId", "value": *sdk.CieTenantId})
	} else {
		model.CieTenantId = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Count != nil {
		model.Count = basetypes.NewFloat64Value(float64(*sdk.Count))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Count", "value": *sdk.Count})
	} else {
		model.Count = basetypes.NewFloat64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.TenantId != nil {
		model.TenantId = basetypes.NewStringValue(*sdk.TenantId)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "TenantId", "value": *sdk.TenantId})
	} else {
		model.TenantId = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.DiscoveredApplications{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.DiscoveredApplications", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for DiscoveredApplications ---
func unpackDiscoveredApplicationsListToSdk(ctx context.Context, list types.List) ([]ztna_connector_all.DiscoveredApplications, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.DiscoveredApplications")
	diags := diag.Diagnostics{}
	var data []models.DiscoveredApplications
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]ztna_connector_all.DiscoveredApplications, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.DiscoveredApplications{}.AttrTypes(), &item)
		unpacked, d := unpackDiscoveredApplicationsToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.DiscoveredApplications", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for DiscoveredApplications ---
func packDiscoveredApplicationsListFromSdk(ctx context.Context, sdks []ztna_connector_all.DiscoveredApplications) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.DiscoveredApplications")
	diags := diag.Diagnostics{}
	var data []models.DiscoveredApplications

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.DiscoveredApplications
		obj, d := packDiscoveredApplicationsFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.DiscoveredApplications{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.DiscoveredApplications", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.DiscoveredApplications{}.AttrType(), data)
}

// --- Unpacker for DiscoveredApplicationsApplicationsInner ---
func unpackDiscoveredApplicationsApplicationsInnerToSdk(ctx context.Context, obj types.Object) (*ztna_connector_all.DiscoveredApplicationsApplicationsInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.DiscoveredApplicationsApplicationsInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.DiscoveredApplicationsApplicationsInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk ztna_connector_all.DiscoveredApplicationsApplicationsInner
	var d diag.Diagnostics
	// Handling Lists
	if !model.AppSpec.IsNull() && !model.AppSpec.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field AppSpec")
		unpacked, d := unpackDiscoveredApplicationsApplicationsInnerAppSpecInnerListToSdk(ctx, model.AppSpec)
		diags.Append(d...)
		sdk.AppSpec = unpacked
	}

	// Handling Primitives
	if !model.Fqdn.IsNull() && !model.Fqdn.IsUnknown() {
		sdk.Fqdn = model.Fqdn.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Fqdn", "value": *sdk.Fqdn})
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

	// Handling Primitives
	if !model.Port.IsNull() && !model.Port.IsUnknown() {
		sdk.Port = model.Port.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Port", "value": *sdk.Port})
	}

	// Handling Primitives
	if !model.Protocol.IsNull() && !model.Protocol.IsUnknown() {
		sdk.Protocol = model.Protocol.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Protocol", "value": *sdk.Protocol})
	}

	// Handling Primitives
	if !model.Provider.IsNull() && !model.Provider.IsUnknown() {
		sdk.Provider = model.Provider.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Provider", "value": *sdk.Provider})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.DiscoveredApplicationsApplicationsInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for DiscoveredApplicationsApplicationsInner ---
func packDiscoveredApplicationsApplicationsInnerFromSdk(ctx context.Context, sdk ztna_connector_all.DiscoveredApplicationsApplicationsInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.DiscoveredApplicationsApplicationsInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.DiscoveredApplicationsApplicationsInner
	var d diag.Diagnostics
	// Handling Lists
	if sdk.AppSpec != nil {
		tflog.Debug(ctx, "Packing list of objects for field AppSpec")
		packed, d := packDiscoveredApplicationsApplicationsInnerAppSpecInnerListFromSdk(ctx, sdk.AppSpec)
		diags.Append(d...)
		model.AppSpec = packed
	} else {
		model.AppSpec = basetypes.NewListNull(models.DiscoveredApplicationsApplicationsInnerAppSpecInner{}.AttrType())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Fqdn != nil {
		model.Fqdn = basetypes.NewStringValue(*sdk.Fqdn)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Fqdn", "value": *sdk.Fqdn})
	} else {
		model.Fqdn = basetypes.NewStringNull()
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
	// Handling Primitives
	// Standard primitive packing
	if sdk.Port != nil {
		model.Port = basetypes.NewStringValue(*sdk.Port)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Port", "value": *sdk.Port})
	} else {
		model.Port = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Protocol != nil {
		model.Protocol = basetypes.NewStringValue(*sdk.Protocol)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Protocol", "value": *sdk.Protocol})
	} else {
		model.Protocol = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Provider != nil {
		model.Provider = basetypes.NewStringValue(*sdk.Provider)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Provider", "value": *sdk.Provider})
	} else {
		model.Provider = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.DiscoveredApplicationsApplicationsInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.DiscoveredApplicationsApplicationsInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for DiscoveredApplicationsApplicationsInner ---
func unpackDiscoveredApplicationsApplicationsInnerListToSdk(ctx context.Context, list types.List) ([]ztna_connector_all.DiscoveredApplicationsApplicationsInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.DiscoveredApplicationsApplicationsInner")
	diags := diag.Diagnostics{}
	var data []models.DiscoveredApplicationsApplicationsInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]ztna_connector_all.DiscoveredApplicationsApplicationsInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.DiscoveredApplicationsApplicationsInner{}.AttrTypes(), &item)
		unpacked, d := unpackDiscoveredApplicationsApplicationsInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.DiscoveredApplicationsApplicationsInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for DiscoveredApplicationsApplicationsInner ---
func packDiscoveredApplicationsApplicationsInnerListFromSdk(ctx context.Context, sdks []ztna_connector_all.DiscoveredApplicationsApplicationsInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.DiscoveredApplicationsApplicationsInner")
	diags := diag.Diagnostics{}
	var data []models.DiscoveredApplicationsApplicationsInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.DiscoveredApplicationsApplicationsInner
		obj, d := packDiscoveredApplicationsApplicationsInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.DiscoveredApplicationsApplicationsInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.DiscoveredApplicationsApplicationsInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.DiscoveredApplicationsApplicationsInner{}.AttrType(), data)
}

// --- Unpacker for DiscoveredApplicationsApplicationsInnerAppSpecInner ---
func unpackDiscoveredApplicationsApplicationsInnerAppSpecInnerToSdk(ctx context.Context, obj types.Object) (*ztna_connector_all.DiscoveredApplicationsApplicationsInnerAppSpecInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.DiscoveredApplicationsApplicationsInnerAppSpecInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.DiscoveredApplicationsApplicationsInnerAppSpecInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk ztna_connector_all.DiscoveredApplicationsApplicationsInnerAppSpecInner
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Fqdn.IsNull() && !model.Fqdn.IsUnknown() {
		sdk.Fqdn = model.Fqdn.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Fqdn", "value": *sdk.Fqdn})
	}

	// Handling Primitives
	if !model.Port.IsNull() && !model.Port.IsUnknown() {
		sdk.Port = model.Port.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Port", "value": *sdk.Port})
	}

	// Handling Primitives
	if !model.Protocol.IsNull() && !model.Protocol.IsUnknown() {
		sdk.Protocol = model.Protocol.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Protocol", "value": *sdk.Protocol})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.DiscoveredApplicationsApplicationsInnerAppSpecInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for DiscoveredApplicationsApplicationsInnerAppSpecInner ---
func packDiscoveredApplicationsApplicationsInnerAppSpecInnerFromSdk(ctx context.Context, sdk ztna_connector_all.DiscoveredApplicationsApplicationsInnerAppSpecInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.DiscoveredApplicationsApplicationsInnerAppSpecInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.DiscoveredApplicationsApplicationsInnerAppSpecInner
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Fqdn != nil {
		model.Fqdn = basetypes.NewStringValue(*sdk.Fqdn)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Fqdn", "value": *sdk.Fqdn})
	} else {
		model.Fqdn = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Port != nil {
		model.Port = basetypes.NewStringValue(*sdk.Port)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Port", "value": *sdk.Port})
	} else {
		model.Port = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Protocol != nil {
		model.Protocol = basetypes.NewStringValue(*sdk.Protocol)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Protocol", "value": *sdk.Protocol})
	} else {
		model.Protocol = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.DiscoveredApplicationsApplicationsInnerAppSpecInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.DiscoveredApplicationsApplicationsInnerAppSpecInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for DiscoveredApplicationsApplicationsInnerAppSpecInner ---
func unpackDiscoveredApplicationsApplicationsInnerAppSpecInnerListToSdk(ctx context.Context, list types.List) ([]ztna_connector_all.DiscoveredApplicationsApplicationsInnerAppSpecInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.DiscoveredApplicationsApplicationsInnerAppSpecInner")
	diags := diag.Diagnostics{}
	var data []models.DiscoveredApplicationsApplicationsInnerAppSpecInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]ztna_connector_all.DiscoveredApplicationsApplicationsInnerAppSpecInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.DiscoveredApplicationsApplicationsInnerAppSpecInner{}.AttrTypes(), &item)
		unpacked, d := unpackDiscoveredApplicationsApplicationsInnerAppSpecInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.DiscoveredApplicationsApplicationsInnerAppSpecInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for DiscoveredApplicationsApplicationsInnerAppSpecInner ---
func packDiscoveredApplicationsApplicationsInnerAppSpecInnerListFromSdk(ctx context.Context, sdks []ztna_connector_all.DiscoveredApplicationsApplicationsInnerAppSpecInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.DiscoveredApplicationsApplicationsInnerAppSpecInner")
	diags := diag.Diagnostics{}
	var data []models.DiscoveredApplicationsApplicationsInnerAppSpecInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.DiscoveredApplicationsApplicationsInnerAppSpecInner
		obj, d := packDiscoveredApplicationsApplicationsInnerAppSpecInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.DiscoveredApplicationsApplicationsInnerAppSpecInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.DiscoveredApplicationsApplicationsInnerAppSpecInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.DiscoveredApplicationsApplicationsInnerAppSpecInner{}.AttrType(), data)
}
