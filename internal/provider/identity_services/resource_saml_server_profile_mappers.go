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

// --- Unpacker for SamlServerProfiles ---
func unpackSamlServerProfilesToSdk(ctx context.Context, obj types.Object) (*identity_services.SamlServerProfiles, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.SamlServerProfiles", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.SamlServerProfiles
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk identity_services.SamlServerProfiles
	var d diag.Diagnostics

	// Handling Primitives
	if !model.Certificate.IsNull() && !model.Certificate.IsUnknown() {
		sdk.Certificate = model.Certificate.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Certificate", "value": sdk.Certificate})
	}

	// Handling Primitives
	if !model.Device.IsNull() && !model.Device.IsUnknown() {
		sdk.Device = model.Device.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Device", "value": *sdk.Device})
	}

	// Handling Primitives
	if !model.EntityId.IsNull() && !model.EntityId.IsUnknown() {
		sdk.EntityId = model.EntityId.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "EntityId", "value": sdk.EntityId})
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
	if !model.MaxClockSkew.IsNull() && !model.MaxClockSkew.IsUnknown() {
		val := int32(model.MaxClockSkew.ValueInt64())
		sdk.MaxClockSkew = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "MaxClockSkew", "value": *sdk.MaxClockSkew})
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	// Handling Primitives
	if !model.SloBindings.IsNull() && !model.SloBindings.IsUnknown() {
		sdk.SloBindings = model.SloBindings.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "SloBindings", "value": *sdk.SloBindings})
	}

	// Handling Primitives
	if !model.Snippet.IsNull() && !model.Snippet.IsUnknown() {
		sdk.Snippet = model.Snippet.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	}

	// Handling Primitives
	if !model.SsoBindings.IsNull() && !model.SsoBindings.IsUnknown() {
		sdk.SsoBindings = model.SsoBindings.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "SsoBindings", "value": sdk.SsoBindings})
	}

	// Handling Primitives
	if !model.SsoUrl.IsNull() && !model.SsoUrl.IsUnknown() {
		sdk.SsoUrl = model.SsoUrl.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "SsoUrl", "value": sdk.SsoUrl})
	}

	// Handling Primitives
	if !model.ValidateIdpCertificate.IsNull() && !model.ValidateIdpCertificate.IsUnknown() {
		sdk.ValidateIdpCertificate = model.ValidateIdpCertificate.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "ValidateIdpCertificate", "value": *sdk.ValidateIdpCertificate})
	}

	// Handling Primitives
	if !model.WantAuthRequestsSigned.IsNull() && !model.WantAuthRequestsSigned.IsUnknown() {
		sdk.WantAuthRequestsSigned = model.WantAuthRequestsSigned.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "WantAuthRequestsSigned", "value": *sdk.WantAuthRequestsSigned})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.SamlServerProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for SamlServerProfiles ---
func packSamlServerProfilesFromSdk(ctx context.Context, sdk identity_services.SamlServerProfiles) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.SamlServerProfiles", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.SamlServerProfiles
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.Certificate = basetypes.NewStringValue(sdk.Certificate)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Certificate", "value": sdk.Certificate})
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
	model.EntityId = basetypes.NewStringValue(sdk.EntityId)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "EntityId", "value": sdk.EntityId})
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
	if sdk.MaxClockSkew != nil {
		model.MaxClockSkew = basetypes.NewInt64Value(int64(*sdk.MaxClockSkew))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "MaxClockSkew", "value": *sdk.MaxClockSkew})
	} else {
		model.MaxClockSkew = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	// Handling Primitives
	// Standard primitive packing
	if sdk.SloBindings != nil {
		model.SloBindings = basetypes.NewStringValue(*sdk.SloBindings)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "SloBindings", "value": *sdk.SloBindings})
	} else {
		model.SloBindings = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Snippet != nil {
		model.Snippet = basetypes.NewStringValue(*sdk.Snippet)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	} else {
		model.Snippet = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	model.SsoBindings = basetypes.NewStringValue(sdk.SsoBindings)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "SsoBindings", "value": sdk.SsoBindings})
	// Handling Primitives
	// Standard primitive packing
	model.SsoUrl = basetypes.NewStringValue(sdk.SsoUrl)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "SsoUrl", "value": sdk.SsoUrl})
	// Handling Primitives
	// Standard primitive packing
	if sdk.ValidateIdpCertificate != nil {
		model.ValidateIdpCertificate = basetypes.NewBoolValue(*sdk.ValidateIdpCertificate)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "ValidateIdpCertificate", "value": *sdk.ValidateIdpCertificate})
	} else {
		model.ValidateIdpCertificate = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.WantAuthRequestsSigned != nil {
		model.WantAuthRequestsSigned = basetypes.NewBoolValue(*sdk.WantAuthRequestsSigned)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "WantAuthRequestsSigned", "value": *sdk.WantAuthRequestsSigned})
	} else {
		model.WantAuthRequestsSigned = basetypes.NewBoolNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.SamlServerProfiles{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.SamlServerProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for SamlServerProfiles ---
func unpackSamlServerProfilesListToSdk(ctx context.Context, list types.List) ([]identity_services.SamlServerProfiles, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.SamlServerProfiles")
	diags := diag.Diagnostics{}
	var data []models.SamlServerProfiles
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]identity_services.SamlServerProfiles, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.SamlServerProfiles{}.AttrTypes(), &item)
		unpacked, d := unpackSamlServerProfilesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.SamlServerProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for SamlServerProfiles ---
func packSamlServerProfilesListFromSdk(ctx context.Context, sdks []identity_services.SamlServerProfiles) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.SamlServerProfiles")
	diags := diag.Diagnostics{}
	var data []models.SamlServerProfiles

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.SamlServerProfiles
		obj, d := packSamlServerProfilesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.SamlServerProfiles{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.SamlServerProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.SamlServerProfiles{}.AttrType(), data)
}
