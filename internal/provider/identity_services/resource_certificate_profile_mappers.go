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

// --- Unpacker for CertificateProfiles ---
func unpackCertificateProfilesToSdk(ctx context.Context, obj types.Object) (*identity_services.CertificateProfiles, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.CertificateProfiles", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.CertificateProfiles
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk identity_services.CertificateProfiles
	var d diag.Diagnostics
	// Handling Primitives
	if !model.BlockExpiredCert.IsNull() && !model.BlockExpiredCert.IsUnknown() {
		sdk.BlockExpiredCert = model.BlockExpiredCert.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "BlockExpiredCert", "value": *sdk.BlockExpiredCert})
	}

	// Handling Primitives
	if !model.BlockTimeoutCert.IsNull() && !model.BlockTimeoutCert.IsUnknown() {
		sdk.BlockTimeoutCert = model.BlockTimeoutCert.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "BlockTimeoutCert", "value": *sdk.BlockTimeoutCert})
	}

	// Handling Primitives
	if !model.BlockUnauthenticatedCert.IsNull() && !model.BlockUnauthenticatedCert.IsUnknown() {
		sdk.BlockUnauthenticatedCert = model.BlockUnauthenticatedCert.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "BlockUnauthenticatedCert", "value": *sdk.BlockUnauthenticatedCert})
	}

	// Handling Primitives
	if !model.BlockUnknownCert.IsNull() && !model.BlockUnknownCert.IsUnknown() {
		sdk.BlockUnknownCert = model.BlockUnknownCert.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "BlockUnknownCert", "value": *sdk.BlockUnknownCert})
	}

	// Handling Lists
	if !model.CaCertificates.IsNull() && !model.CaCertificates.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field CaCertificates")
		unpacked, d := unpackCertificateProfilesCaCertificatesInnerListToSdk(ctx, model.CaCertificates)
		diags.Append(d...)
		sdk.CaCertificates = unpacked
	}

	// Handling Primitives
	if !model.CertStatusTimeout.IsNull() && !model.CertStatusTimeout.IsUnknown() {
		sdk.CertStatusTimeout = model.CertStatusTimeout.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "CertStatusTimeout", "value": *sdk.CertStatusTimeout})
	}

	// Handling Primitives
	if !model.CrlReceiveTimeout.IsNull() && !model.CrlReceiveTimeout.IsUnknown() {
		sdk.CrlReceiveTimeout = model.CrlReceiveTimeout.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "CrlReceiveTimeout", "value": *sdk.CrlReceiveTimeout})
	}

	// Handling Primitives
	if !model.Device.IsNull() && !model.Device.IsUnknown() {
		sdk.Device = model.Device.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Device", "value": *sdk.Device})
	}

	// Handling Primitives
	if !model.Domain.IsNull() && !model.Domain.IsUnknown() {
		sdk.Domain = model.Domain.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Domain", "value": *sdk.Domain})
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
	if !model.OcspReceiveTimeout.IsNull() && !model.OcspReceiveTimeout.IsUnknown() {
		sdk.OcspReceiveTimeout = model.OcspReceiveTimeout.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "OcspReceiveTimeout", "value": *sdk.OcspReceiveTimeout})
	}

	// Handling Primitives
	if !model.Snippet.IsNull() && !model.Snippet.IsUnknown() {
		sdk.Snippet = model.Snippet.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	}

	// Handling Primitives
	if !model.UseCrl.IsNull() && !model.UseCrl.IsUnknown() {
		sdk.UseCrl = model.UseCrl.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "UseCrl", "value": *sdk.UseCrl})
	}

	// Handling Primitives
	if !model.UseOcsp.IsNull() && !model.UseOcsp.IsUnknown() {
		sdk.UseOcsp = model.UseOcsp.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "UseOcsp", "value": *sdk.UseOcsp})
	}

	// Handling Objects
	if !model.UsernameField.IsNull() && !model.UsernameField.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field UsernameField")
		unpacked, d := unpackCertificateProfilesUsernameFieldToSdk(ctx, model.UsernameField)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "UsernameField"})
		}
		if unpacked != nil {
			sdk.UsernameField = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.CertificateProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for CertificateProfiles ---
func packCertificateProfilesFromSdk(ctx context.Context, sdk identity_services.CertificateProfiles) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.CertificateProfiles", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.CertificateProfiles
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.BlockExpiredCert != nil {
		model.BlockExpiredCert = basetypes.NewBoolValue(*sdk.BlockExpiredCert)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "BlockExpiredCert", "value": *sdk.BlockExpiredCert})
	} else {
		model.BlockExpiredCert = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.BlockTimeoutCert != nil {
		model.BlockTimeoutCert = basetypes.NewBoolValue(*sdk.BlockTimeoutCert)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "BlockTimeoutCert", "value": *sdk.BlockTimeoutCert})
	} else {
		model.BlockTimeoutCert = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.BlockUnauthenticatedCert != nil {
		model.BlockUnauthenticatedCert = basetypes.NewBoolValue(*sdk.BlockUnauthenticatedCert)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "BlockUnauthenticatedCert", "value": *sdk.BlockUnauthenticatedCert})
	} else {
		model.BlockUnauthenticatedCert = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.BlockUnknownCert != nil {
		model.BlockUnknownCert = basetypes.NewBoolValue(*sdk.BlockUnknownCert)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "BlockUnknownCert", "value": *sdk.BlockUnknownCert})
	} else {
		model.BlockUnknownCert = basetypes.NewBoolNull()
	}
	// Handling Lists
	if sdk.CaCertificates != nil {
		tflog.Debug(ctx, "Packing list of objects for field CaCertificates")
		packed, d := packCertificateProfilesCaCertificatesInnerListFromSdk(ctx, sdk.CaCertificates)
		diags.Append(d...)
		model.CaCertificates = packed
	} else {
		model.CaCertificates = basetypes.NewListNull(models.CertificateProfilesCaCertificatesInner{}.AttrType())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.CertStatusTimeout != nil {
		model.CertStatusTimeout = basetypes.NewStringValue(*sdk.CertStatusTimeout)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "CertStatusTimeout", "value": *sdk.CertStatusTimeout})
	} else {
		model.CertStatusTimeout = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.CrlReceiveTimeout != nil {
		model.CrlReceiveTimeout = basetypes.NewStringValue(*sdk.CrlReceiveTimeout)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "CrlReceiveTimeout", "value": *sdk.CrlReceiveTimeout})
	} else {
		model.CrlReceiveTimeout = basetypes.NewStringNull()
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
	if sdk.Domain != nil {
		model.Domain = basetypes.NewStringValue(*sdk.Domain)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Domain", "value": *sdk.Domain})
	} else {
		model.Domain = basetypes.NewStringNull()
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
	if sdk.OcspReceiveTimeout != nil {
		model.OcspReceiveTimeout = basetypes.NewStringValue(*sdk.OcspReceiveTimeout)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "OcspReceiveTimeout", "value": *sdk.OcspReceiveTimeout})
	} else {
		model.OcspReceiveTimeout = basetypes.NewStringNull()
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
	if sdk.UseCrl != nil {
		model.UseCrl = basetypes.NewBoolValue(*sdk.UseCrl)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "UseCrl", "value": *sdk.UseCrl})
	} else {
		model.UseCrl = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.UseOcsp != nil {
		model.UseOcsp = basetypes.NewBoolValue(*sdk.UseOcsp)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "UseOcsp", "value": *sdk.UseOcsp})
	} else {
		model.UseOcsp = basetypes.NewBoolNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.UsernameField != nil {
		tflog.Debug(ctx, "Packing nested object for field UsernameField")
		packed, d := packCertificateProfilesUsernameFieldFromSdk(ctx, *sdk.UsernameField)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "UsernameField"})
		}
		model.UsernameField = packed
	} else {
		model.UsernameField = basetypes.NewObjectNull(models.CertificateProfilesUsernameField{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.CertificateProfiles{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.CertificateProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for CertificateProfiles ---
func unpackCertificateProfilesListToSdk(ctx context.Context, list types.List) ([]identity_services.CertificateProfiles, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.CertificateProfiles")
	diags := diag.Diagnostics{}
	var data []models.CertificateProfiles
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]identity_services.CertificateProfiles, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.CertificateProfiles{}.AttrTypes(), &item)
		unpacked, d := unpackCertificateProfilesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.CertificateProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for CertificateProfiles ---
func packCertificateProfilesListFromSdk(ctx context.Context, sdks []identity_services.CertificateProfiles) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.CertificateProfiles")
	diags := diag.Diagnostics{}
	var data []models.CertificateProfiles

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.CertificateProfiles
		obj, d := packCertificateProfilesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.CertificateProfiles{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.CertificateProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.CertificateProfiles{}.AttrType(), data)
}

// --- Unpacker for CertificateProfilesCaCertificatesInner ---
func unpackCertificateProfilesCaCertificatesInnerToSdk(ctx context.Context, obj types.Object) (*identity_services.CertificateProfilesCaCertificatesInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.CertificateProfilesCaCertificatesInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.CertificateProfilesCaCertificatesInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk identity_services.CertificateProfilesCaCertificatesInner
	var d diag.Diagnostics
	// Handling Primitives
	if !model.DefaultOcspUrl.IsNull() && !model.DefaultOcspUrl.IsUnknown() {
		sdk.DefaultOcspUrl = model.DefaultOcspUrl.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DefaultOcspUrl", "value": *sdk.DefaultOcspUrl})
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	// Handling Primitives
	if !model.OcspVerifyCert.IsNull() && !model.OcspVerifyCert.IsUnknown() {
		sdk.OcspVerifyCert = model.OcspVerifyCert.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "OcspVerifyCert", "value": *sdk.OcspVerifyCert})
	}

	// Handling Primitives
	if !model.TemplateName.IsNull() && !model.TemplateName.IsUnknown() {
		sdk.TemplateName = model.TemplateName.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "TemplateName", "value": *sdk.TemplateName})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.CertificateProfilesCaCertificatesInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for CertificateProfilesCaCertificatesInner ---
func packCertificateProfilesCaCertificatesInnerFromSdk(ctx context.Context, sdk identity_services.CertificateProfilesCaCertificatesInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.CertificateProfilesCaCertificatesInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.CertificateProfilesCaCertificatesInner
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.DefaultOcspUrl != nil {
		model.DefaultOcspUrl = basetypes.NewStringValue(*sdk.DefaultOcspUrl)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DefaultOcspUrl", "value": *sdk.DefaultOcspUrl})
	} else {
		model.DefaultOcspUrl = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	// Handling Primitives
	// Standard primitive packing
	if sdk.OcspVerifyCert != nil {
		model.OcspVerifyCert = basetypes.NewStringValue(*sdk.OcspVerifyCert)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "OcspVerifyCert", "value": *sdk.OcspVerifyCert})
	} else {
		model.OcspVerifyCert = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.TemplateName != nil {
		model.TemplateName = basetypes.NewStringValue(*sdk.TemplateName)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "TemplateName", "value": *sdk.TemplateName})
	} else {
		model.TemplateName = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.CertificateProfilesCaCertificatesInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.CertificateProfilesCaCertificatesInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for CertificateProfilesCaCertificatesInner ---
func unpackCertificateProfilesCaCertificatesInnerListToSdk(ctx context.Context, list types.List) ([]identity_services.CertificateProfilesCaCertificatesInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.CertificateProfilesCaCertificatesInner")
	diags := diag.Diagnostics{}
	var data []models.CertificateProfilesCaCertificatesInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]identity_services.CertificateProfilesCaCertificatesInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.CertificateProfilesCaCertificatesInner{}.AttrTypes(), &item)
		unpacked, d := unpackCertificateProfilesCaCertificatesInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.CertificateProfilesCaCertificatesInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for CertificateProfilesCaCertificatesInner ---
func packCertificateProfilesCaCertificatesInnerListFromSdk(ctx context.Context, sdks []identity_services.CertificateProfilesCaCertificatesInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.CertificateProfilesCaCertificatesInner")
	diags := diag.Diagnostics{}
	var data []models.CertificateProfilesCaCertificatesInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.CertificateProfilesCaCertificatesInner
		obj, d := packCertificateProfilesCaCertificatesInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.CertificateProfilesCaCertificatesInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.CertificateProfilesCaCertificatesInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.CertificateProfilesCaCertificatesInner{}.AttrType(), data)
}

// --- Unpacker for CertificateProfilesUsernameField ---
func unpackCertificateProfilesUsernameFieldToSdk(ctx context.Context, obj types.Object) (*identity_services.CertificateProfilesUsernameField, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.CertificateProfilesUsernameField", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.CertificateProfilesUsernameField
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk identity_services.CertificateProfilesUsernameField
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Subject.IsNull() && !model.Subject.IsUnknown() {
		sdk.Subject = model.Subject.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Subject", "value": *sdk.Subject})
	}

	// Handling Primitives
	if !model.SubjectAlt.IsNull() && !model.SubjectAlt.IsUnknown() {
		sdk.SubjectAlt = model.SubjectAlt.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "SubjectAlt", "value": *sdk.SubjectAlt})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.CertificateProfilesUsernameField", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for CertificateProfilesUsernameField ---
func packCertificateProfilesUsernameFieldFromSdk(ctx context.Context, sdk identity_services.CertificateProfilesUsernameField) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.CertificateProfilesUsernameField", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.CertificateProfilesUsernameField
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Subject != nil {
		model.Subject = basetypes.NewStringValue(*sdk.Subject)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Subject", "value": *sdk.Subject})
	} else {
		model.Subject = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.SubjectAlt != nil {
		model.SubjectAlt = basetypes.NewStringValue(*sdk.SubjectAlt)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "SubjectAlt", "value": *sdk.SubjectAlt})
	} else {
		model.SubjectAlt = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.CertificateProfilesUsernameField{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.CertificateProfilesUsernameField", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for CertificateProfilesUsernameField ---
func unpackCertificateProfilesUsernameFieldListToSdk(ctx context.Context, list types.List) ([]identity_services.CertificateProfilesUsernameField, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.CertificateProfilesUsernameField")
	diags := diag.Diagnostics{}
	var data []models.CertificateProfilesUsernameField
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]identity_services.CertificateProfilesUsernameField, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.CertificateProfilesUsernameField{}.AttrTypes(), &item)
		unpacked, d := unpackCertificateProfilesUsernameFieldToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.CertificateProfilesUsernameField", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for CertificateProfilesUsernameField ---
func packCertificateProfilesUsernameFieldListFromSdk(ctx context.Context, sdks []identity_services.CertificateProfilesUsernameField) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.CertificateProfilesUsernameField")
	diags := diag.Diagnostics{}
	var data []models.CertificateProfilesUsernameField

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.CertificateProfilesUsernameField
		obj, d := packCertificateProfilesUsernameFieldFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.CertificateProfilesUsernameField{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.CertificateProfilesUsernameField", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.CertificateProfilesUsernameField{}.AttrType(), data)
}
