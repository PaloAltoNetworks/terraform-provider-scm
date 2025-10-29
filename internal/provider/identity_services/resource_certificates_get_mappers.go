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

	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/identity_services"
	"github.com/paloaltonetworks/scm-go/generated/identity_services"
)










// --- Unpacker for CertificatesGet ---
func unpackCertificatesGetToSdk(ctx context.Context, obj types.Object) (*identity_services.CertificatesGet, diag.Diagnostics) {
    tflog.Debug(ctx, "Entering unpack helper for models.CertificatesGet", map[string]interface{}{"tf_object": obj})
    diags := diag.Diagnostics{}
    var model models.CertificatesGet
    diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
    if diags.HasError() {
        tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
        return nil, diags
    }
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

    var sdk identity_services.CertificatesGet
    var d diag.Diagnostics

    // Handling Primitives
    if !model.Algorithm.IsNull() && !model.Algorithm.IsUnknown() {
        sdk.Algorithm = model.Algorithm.ValueStringPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Algorithm", "value": *sdk.Algorithm})
    }

    // Handling Primitives
    if !model.Ca.IsNull() && !model.Ca.IsUnknown() {
        sdk.Ca = model.Ca.ValueBoolPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Ca", "value": *sdk.Ca})
    }

    // Handling Primitives
    if !model.CommonName.IsNull() && !model.CommonName.IsUnknown() {
        sdk.CommonName = model.CommonName.ValueStringPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "CommonName", "value": *sdk.CommonName})
    }

    // Handling Primitives
    if !model.CommonNameInt.IsNull() && !model.CommonNameInt.IsUnknown() {
        sdk.CommonNameInt = model.CommonNameInt.ValueStringPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "CommonNameInt", "value": *sdk.CommonNameInt})
    }

    // Handling Primitives
    if !model.Device.IsNull() && !model.Device.IsUnknown() {
        sdk.Device = model.Device.ValueStringPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Device", "value": *sdk.Device})
    }

    // Handling Primitives
    if !model.ExpiryEpoch.IsNull() && !model.ExpiryEpoch.IsUnknown() {
        sdk.ExpiryEpoch = model.ExpiryEpoch.ValueStringPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "ExpiryEpoch", "value": *sdk.ExpiryEpoch})
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
    if !model.Issuer.IsNull() && !model.Issuer.IsUnknown() {
        sdk.Issuer = model.Issuer.ValueStringPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Issuer", "value": *sdk.Issuer})
    }

    // Handling Primitives
    if !model.IssuerHash.IsNull() && !model.IssuerHash.IsUnknown() {
        sdk.IssuerHash = model.IssuerHash.ValueStringPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "IssuerHash", "value": *sdk.IssuerHash})
    }

    // Handling Primitives
    if !model.Name.IsNull() && !model.Name.IsUnknown() {
        sdk.Name = model.Name.ValueStringPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
    }

    // Handling Primitives
    if !model.NotValidAfter.IsNull() && !model.NotValidAfter.IsUnknown() {
        sdk.NotValidAfter = model.NotValidAfter.ValueStringPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "NotValidAfter", "value": *sdk.NotValidAfter})
    }

    // Handling Primitives
    if !model.NotValidBefore.IsNull() && !model.NotValidBefore.IsUnknown() {
        sdk.NotValidBefore = model.NotValidBefore.ValueStringPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "NotValidBefore", "value": *sdk.NotValidBefore})
    }

    // Handling Primitives
    if !model.PublicKey.IsNull() && !model.PublicKey.IsUnknown() {
        sdk.PublicKey = model.PublicKey.ValueStringPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "PublicKey", "value": *sdk.PublicKey})
    }

    // Handling Primitives
    if !model.Snippet.IsNull() && !model.Snippet.IsUnknown() {
        sdk.Snippet = model.Snippet.ValueStringPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
    }

    // Handling Primitives
    if !model.Subject.IsNull() && !model.Subject.IsUnknown() {
        sdk.Subject = model.Subject.ValueStringPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Subject", "value": *sdk.Subject})
    }

    // Handling Primitives
    if !model.SubjectHash.IsNull() && !model.SubjectHash.IsUnknown() {
        sdk.SubjectHash = model.SubjectHash.ValueStringPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "SubjectHash", "value": *sdk.SubjectHash})
    }

    // Handling Primitives
    if !model.SubjectInt.IsNull() && !model.SubjectInt.IsUnknown() {
        sdk.SubjectInt = model.SubjectInt.ValueStringPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "SubjectInt", "value": *sdk.SubjectInt})
    }

diags.Append(d...)

    tflog.Debug(ctx, "Exiting unpack helper for models.CertificatesGet", map[string]interface{}{"has_errors": diags.HasError()})
    return &sdk, diags

}

// --- Packer for CertificatesGet ---
func packCertificatesGetFromSdk(ctx context.Context, sdk identity_services.CertificatesGet) (types.Object, diag.Diagnostics) {
    tflog.Debug(ctx, "Entering pack helper for models.CertificatesGet", map[string]interface{}{"sdk_struct": sdk})
    diags := diag.Diagnostics{}
    var model models.CertificatesGet
    var d diag.Diagnostics
    // Handling Primitives
    // Standard primitive packing
    if sdk.Algorithm != nil {
        model.Algorithm = basetypes.NewStringValue(*sdk.Algorithm)
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Algorithm", "value": *sdk.Algorithm})
    } else {
        model.Algorithm = basetypes.NewStringNull()
    }
    // Handling Primitives
    // Standard primitive packing
    if sdk.Ca != nil {
        model.Ca = basetypes.NewBoolValue(*sdk.Ca)
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Ca", "value": *sdk.Ca})
    } else {
        model.Ca = basetypes.NewBoolNull()
    }
    // Handling Primitives
    // Standard primitive packing
    if sdk.CommonName != nil {
        model.CommonName = basetypes.NewStringValue(*sdk.CommonName)
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "CommonName", "value": *sdk.CommonName})
    } else {
        model.CommonName = basetypes.NewStringNull()
    }
    // Handling Primitives
    // Standard primitive packing
    if sdk.CommonNameInt != nil {
        model.CommonNameInt = basetypes.NewStringValue(*sdk.CommonNameInt)
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "CommonNameInt", "value": *sdk.CommonNameInt})
    } else {
        model.CommonNameInt = basetypes.NewStringNull()
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
    if sdk.ExpiryEpoch != nil {
        model.ExpiryEpoch = basetypes.NewStringValue(*sdk.ExpiryEpoch)
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "ExpiryEpoch", "value": *sdk.ExpiryEpoch})
    } else {
        model.ExpiryEpoch = basetypes.NewStringNull()
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
    if sdk.Issuer != nil {
        model.Issuer = basetypes.NewStringValue(*sdk.Issuer)
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Issuer", "value": *sdk.Issuer})
    } else {
        model.Issuer = basetypes.NewStringNull()
    }
    // Handling Primitives
    // Standard primitive packing
    if sdk.IssuerHash != nil {
        model.IssuerHash = basetypes.NewStringValue(*sdk.IssuerHash)
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "IssuerHash", "value": *sdk.IssuerHash})
    } else {
        model.IssuerHash = basetypes.NewStringNull()
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
    if sdk.NotValidAfter != nil {
        model.NotValidAfter = basetypes.NewStringValue(*sdk.NotValidAfter)
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "NotValidAfter", "value": *sdk.NotValidAfter})
    } else {
        model.NotValidAfter = basetypes.NewStringNull()
    }
    // Handling Primitives
    // Standard primitive packing
    if sdk.NotValidBefore != nil {
        model.NotValidBefore = basetypes.NewStringValue(*sdk.NotValidBefore)
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "NotValidBefore", "value": *sdk.NotValidBefore})
    } else {
        model.NotValidBefore = basetypes.NewStringNull()
    }
    // Handling Primitives
    // Standard primitive packing
    if sdk.PublicKey != nil {
        model.PublicKey = basetypes.NewStringValue(*sdk.PublicKey)
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "PublicKey", "value": *sdk.PublicKey})
    } else {
        model.PublicKey = basetypes.NewStringNull()
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
    if sdk.Subject != nil {
        model.Subject = basetypes.NewStringValue(*sdk.Subject)
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Subject", "value": *sdk.Subject})
    } else {
        model.Subject = basetypes.NewStringNull()
    }
    // Handling Primitives
    // Standard primitive packing
    if sdk.SubjectHash != nil {
        model.SubjectHash = basetypes.NewStringValue(*sdk.SubjectHash)
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "SubjectHash", "value": *sdk.SubjectHash})
    } else {
        model.SubjectHash = basetypes.NewStringNull()
    }
    // Handling Primitives
    // Standard primitive packing
    if sdk.SubjectInt != nil {
        model.SubjectInt = basetypes.NewStringValue(*sdk.SubjectInt)
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "SubjectInt", "value": *sdk.SubjectInt})
    } else {
        model.SubjectInt = basetypes.NewStringNull()
    }
diags.Append(d...)

    obj, d := types.ObjectValueFrom(ctx, models.CertificatesGet{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
    diags.Append(d...)
    tflog.Debug(ctx, "Exiting pack helper for models.CertificatesGet", map[string]interface{}{"has_errors": diags.HasError()})
    return obj, diags

}

// --- List Unpacker for CertificatesGet ---
func unpackCertificatesGetListToSdk(ctx context.Context, list types.List) ([]identity_services.CertificatesGet, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.CertificatesGet")
	diags := diag.Diagnostics{}
	var data []models.CertificatesGet
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]identity_services.CertificatesGet, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.CertificatesGet{}.AttrTypes(), &item)
		unpacked, d := unpackCertificatesGetToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.CertificatesGet", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for CertificatesGet ---
func packCertificatesGetListFromSdk(ctx context.Context, sdks []identity_services.CertificatesGet) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.CertificatesGet")
	diags := diag.Diagnostics{}
	var data []models.CertificatesGet

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.CertificatesGet
		obj, d := packCertificatesGetFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.CertificatesGet{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.CertificatesGet", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.CertificatesGet{}.AttrType(), data)
}

// --- Unpacker for CertificatesImport ---
func unpackCertificatesImportToSdk(ctx context.Context, obj types.Object) (*identity_services.CertificatesImport, diag.Diagnostics) {
    tflog.Debug(ctx, "Entering unpack helper for models.CertificatesImport", map[string]interface{}{"tf_object": obj})
    diags := diag.Diagnostics{}
    var model models.CertificatesImport
    diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
    if diags.HasError() {
        tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
        return nil, diags
    }
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

    var sdk identity_services.CertificatesImport
    var d diag.Diagnostics


    // Handling Primitives
    if !model.CertificateFile.IsNull() && !model.CertificateFile.IsUnknown() {
        sdk.CertificateFile = model.CertificateFile.ValueString()
        tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "CertificateFile", "value": sdk.CertificateFile})
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
    if !model.Format.IsNull() && !model.Format.IsUnknown() {
        sdk.Format = model.Format.ValueString()
        tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Format", "value": sdk.Format})
    }

    // Handling Primitives
    if !model.KeyFile.IsNull() && !model.KeyFile.IsUnknown() {
        sdk.KeyFile = model.KeyFile.ValueStringPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "KeyFile", "value": *sdk.KeyFile})
    }

    // Handling Primitives
    if !model.Name.IsNull() && !model.Name.IsUnknown() {
        sdk.Name = model.Name.ValueString()
        tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
    }

    // Handling Primitives
    if !model.Passphrase.IsNull() && !model.Passphrase.IsUnknown() {
        sdk.Passphrase = model.Passphrase.ValueStringPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Passphrase", "value": *sdk.Passphrase})
    }

    // Handling Primitives
    if !model.Snippet.IsNull() && !model.Snippet.IsUnknown() {
        sdk.Snippet = model.Snippet.ValueStringPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
    }

diags.Append(d...)

    tflog.Debug(ctx, "Exiting unpack helper for models.CertificatesImport", map[string]interface{}{"has_errors": diags.HasError()})
    return &sdk, diags

}

// --- Packer for CertificatesImport ---
func packCertificatesImportFromSdk(ctx context.Context, sdk identity_services.CertificatesImport) (types.Object, diag.Diagnostics) {
    tflog.Debug(ctx, "Entering pack helper for models.CertificatesImport", map[string]interface{}{"sdk_struct": sdk})
    diags := diag.Diagnostics{}
    var model models.CertificatesImport
    var d diag.Diagnostics
    // Handling Primitives
    // Standard primitive packing
    model.CertificateFile = basetypes.NewStringValue(sdk.CertificateFile)
    tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "CertificateFile", "value": sdk.CertificateFile})
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
    model.Format = basetypes.NewStringValue(sdk.Format)
    tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Format", "value": sdk.Format})
    // Handling Primitives
    // Standard primitive packing
    if sdk.KeyFile != nil {
        model.KeyFile = basetypes.NewStringValue(*sdk.KeyFile)
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "KeyFile", "value": *sdk.KeyFile})
    } else {
        model.KeyFile = basetypes.NewStringNull()
    }
    // Handling Primitives
    // Standard primitive packing
    model.Name = basetypes.NewStringValue(sdk.Name)
    tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
    // Handling Primitives
    // Standard primitive packing
    if sdk.Passphrase != nil {
        model.Passphrase = basetypes.NewStringValue(*sdk.Passphrase)
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Passphrase", "value": *sdk.Passphrase})
    } else {
        model.Passphrase = basetypes.NewStringNull()
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

    obj, d := types.ObjectValueFrom(ctx, models.CertificatesImport{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
    diags.Append(d...)
    tflog.Debug(ctx, "Exiting pack helper for models.CertificatesImport", map[string]interface{}{"has_errors": diags.HasError()})
    return obj, diags

}

// --- List Unpacker for CertificatesImport ---
func unpackCertificatesImportListToSdk(ctx context.Context, list types.List) ([]identity_services.CertificatesImport, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.CertificatesImport")
	diags := diag.Diagnostics{}
	var data []models.CertificatesImport
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]identity_services.CertificatesImport, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.CertificatesImport{}.AttrTypes(), &item)
		unpacked, d := unpackCertificatesImportToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.CertificatesImport", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for CertificatesImport ---
func packCertificatesImportListFromSdk(ctx context.Context, sdks []identity_services.CertificatesImport) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.CertificatesImport")
	diags := diag.Diagnostics{}
	var data []models.CertificatesImport

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.CertificatesImport
		obj, d := packCertificatesImportFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.CertificatesImport{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.CertificatesImport", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.CertificatesImport{}.AttrType(), data)
}


*/
