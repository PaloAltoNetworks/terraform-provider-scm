package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/identity_services"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/identity_services"
)

// --- Unpacker for KerberosServerProfiles ---
func unpackKerberosServerProfilesToSdk(ctx context.Context, obj types.Object) (*identity_services.KerberosServerProfiles, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.KerberosServerProfiles", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.KerberosServerProfiles
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk identity_services.KerberosServerProfiles
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

	// Handling Lists
	if !model.Server.IsNull() && !model.Server.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field Server")
		unpacked, d := unpackKerberosServerProfilesServerInnerListToSdk(ctx, model.Server)
		diags.Append(d...)
		sdk.Server = unpacked
	}

	// Handling Primitives
	if !model.Snippet.IsNull() && !model.Snippet.IsUnknown() {
		sdk.Snippet = model.Snippet.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.KerberosServerProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for KerberosServerProfiles ---
func packKerberosServerProfilesFromSdk(ctx context.Context, sdk identity_services.KerberosServerProfiles) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.KerberosServerProfiles", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.KerberosServerProfiles
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
	// Handling Lists
	if sdk.Server != nil {
		tflog.Debug(ctx, "Packing list of objects for field Server")
		packed, d := packKerberosServerProfilesServerInnerListFromSdk(ctx, sdk.Server)
		diags.Append(d...)
		model.Server = packed
	} else {
		model.Server = basetypes.NewListNull(models.KerberosServerProfilesServerInner{}.AttrType())
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

	obj, d := types.ObjectValueFrom(ctx, models.KerberosServerProfiles{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.KerberosServerProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for KerberosServerProfiles ---
func unpackKerberosServerProfilesListToSdk(ctx context.Context, list types.List) ([]identity_services.KerberosServerProfiles, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.KerberosServerProfiles")
	diags := diag.Diagnostics{}
	var data []models.KerberosServerProfiles
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]identity_services.KerberosServerProfiles, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.KerberosServerProfiles{}.AttrTypes(), &item)
		unpacked, d := unpackKerberosServerProfilesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.KerberosServerProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for KerberosServerProfiles ---
func packKerberosServerProfilesListFromSdk(ctx context.Context, sdks []identity_services.KerberosServerProfiles) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.KerberosServerProfiles")
	diags := diag.Diagnostics{}
	var data []models.KerberosServerProfiles

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.KerberosServerProfiles
		obj, d := packKerberosServerProfilesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.KerberosServerProfiles{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.KerberosServerProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.KerberosServerProfiles{}.AttrType(), data)
}

// --- Unpacker for KerberosServerProfilesServerInner ---
func unpackKerberosServerProfilesServerInnerToSdk(ctx context.Context, obj types.Object) (*identity_services.KerberosServerProfilesServerInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.KerberosServerProfilesServerInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.KerberosServerProfilesServerInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk identity_services.KerberosServerProfilesServerInner
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Host.IsNull() && !model.Host.IsUnknown() {
		sdk.Host = model.Host.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Host", "value": sdk.Host})
	}

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

	tflog.Debug(ctx, "Exiting unpack helper for models.KerberosServerProfilesServerInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for KerberosServerProfilesServerInner ---
func packKerberosServerProfilesServerInnerFromSdk(ctx context.Context, sdk identity_services.KerberosServerProfilesServerInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.KerberosServerProfilesServerInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.KerberosServerProfilesServerInner
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.Host = basetypes.NewStringValue(sdk.Host)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Host", "value": sdk.Host})
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

	obj, d := types.ObjectValueFrom(ctx, models.KerberosServerProfilesServerInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.KerberosServerProfilesServerInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for KerberosServerProfilesServerInner ---
func unpackKerberosServerProfilesServerInnerListToSdk(ctx context.Context, list types.List) ([]identity_services.KerberosServerProfilesServerInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.KerberosServerProfilesServerInner")
	diags := diag.Diagnostics{}
	var data []models.KerberosServerProfilesServerInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]identity_services.KerberosServerProfilesServerInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.KerberosServerProfilesServerInner{}.AttrTypes(), &item)
		unpacked, d := unpackKerberosServerProfilesServerInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.KerberosServerProfilesServerInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for KerberosServerProfilesServerInner ---
func packKerberosServerProfilesServerInnerListFromSdk(ctx context.Context, sdks []identity_services.KerberosServerProfilesServerInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.KerberosServerProfilesServerInner")
	diags := diag.Diagnostics{}
	var data []models.KerberosServerProfilesServerInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.KerberosServerProfilesServerInner
		obj, d := packKerberosServerProfilesServerInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.KerberosServerProfilesServerInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.KerberosServerProfilesServerInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.KerberosServerProfilesServerInner{}.AttrType(), data)
}
