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

// --- Unpacker for CertificatesPost ---
func unpackCertificatesPostToSdk(ctx context.Context, obj types.Object) (*identity_services.CertificatesPost, diag.Diagnostics) {
    tflog.Debug(ctx, "Entering unpack helper for models.CertificatesPost", map[string]interface{}{"tf_object": obj})
    diags := diag.Diagnostics{}
    var model models.CertificatesPost
    diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
    if diags.HasError() {
        tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
        return nil, diags
    }
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

    var sdk identity_services.CertificatesPost
    var d diag.Diagnostics

    // Handling Objects
    if !model.Algorithm.IsNull() && !model.Algorithm.IsUnknown() {
        tflog.Debug(ctx, "Unpacking nested object for field Algorithm")
        unpacked, d := unpackCertificatesPostAlgorithmToSdk(ctx, model.Algorithm)
        diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Algorithm"})
		}
        if unpacked != nil {
            sdk.Algorithm = *unpacked
        }
    }

    // Handling Lists
    if !model.AlternateEmail.IsNull() && !model.AlternateEmail.IsUnknown() {
        tflog.Debug(ctx, "Unpacking list of primitives for field AlternateEmail")
        diags.Append(model.AlternateEmail.ElementsAs(ctx, &sdk.AlternateEmail, false)...)
    }

    // Handling Primitives
    if !model.CertificateName.IsNull() && !model.CertificateName.IsUnknown() {
        sdk.CertificateName = model.CertificateName.ValueString()
        tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "CertificateName", "value": sdk.CertificateName})
    }

    // Handling Primitives
    if !model.CommonName.IsNull() && !model.CommonName.IsUnknown() {
        sdk.CommonName = model.CommonName.ValueString()
        tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "CommonName", "value": sdk.CommonName})
    }

    // Handling Primitives
    if !model.CountryCode.IsNull() && !model.CountryCode.IsUnknown() {
        sdk.CountryCode = model.CountryCode.ValueStringPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "CountryCode", "value": *sdk.CountryCode})
    }

    // Handling Primitives
    if !model.DayTillExpiration.IsNull() && !model.DayTillExpiration.IsUnknown() {
        val := int32(model.DayTillExpiration.ValueInt64())
        sdk.DayTillExpiration = &val
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DayTillExpiration", "value": *sdk.DayTillExpiration})
    }

    // Handling Lists
    if !model.Department.IsNull() && !model.Department.IsUnknown() {
        tflog.Debug(ctx, "Unpacking list of primitives for field Department")
        diags.Append(model.Department.ElementsAs(ctx, &sdk.Department, false)...)
    }

    // Handling Primitives
    if !model.Device.IsNull() && !model.Device.IsUnknown() {
        sdk.Device = model.Device.ValueStringPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Device", "value": *sdk.Device})
    }

    // Handling Primitives
    if !model.Digest.IsNull() && !model.Digest.IsUnknown() {
        sdk.Digest = model.Digest.ValueString()
        tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Digest", "value": sdk.Digest})
    }

    // Handling Primitives
    if !model.Email.IsNull() && !model.Email.IsUnknown() {
        sdk.Email = model.Email.ValueStringPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Email", "value": *sdk.Email})
    }

    // Handling Primitives
    if !model.Folder.IsNull() && !model.Folder.IsUnknown() {
        sdk.Folder = model.Folder.ValueStringPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Folder", "value": *sdk.Folder})
    }

    // Handling Lists
    if !model.Hostname.IsNull() && !model.Hostname.IsUnknown() {
        tflog.Debug(ctx, "Unpacking list of primitives for field Hostname")
        diags.Append(model.Hostname.ElementsAs(ctx, &sdk.Hostname, false)...)
    }

    // Handling Lists
    if !model.Ip.IsNull() && !model.Ip.IsUnknown() {
        tflog.Debug(ctx, "Unpacking list of primitives for field Ip")
        diags.Append(model.Ip.ElementsAs(ctx, &sdk.Ip, false)...)
    }

    // Handling Primitives
    if !model.IsBlockPrivateKey.IsNull() && !model.IsBlockPrivateKey.IsUnknown() {
        sdk.IsBlockPrivateKey = model.IsBlockPrivateKey.ValueBoolPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "IsBlockPrivateKey", "value": *sdk.IsBlockPrivateKey})
    }

    // Handling Primitives
    if !model.IsCertificateAuthority.IsNull() && !model.IsCertificateAuthority.IsUnknown() {
        sdk.IsCertificateAuthority = model.IsCertificateAuthority.ValueBoolPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "IsCertificateAuthority", "value": *sdk.IsCertificateAuthority})
    }

    // Handling Primitives
    if !model.Locality.IsNull() && !model.Locality.IsUnknown() {
        sdk.Locality = model.Locality.ValueStringPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Locality", "value": *sdk.Locality})
    }

    // Handling Primitives
    if !model.OcspResponderUrl.IsNull() && !model.OcspResponderUrl.IsUnknown() {
        sdk.OcspResponderUrl = model.OcspResponderUrl.ValueStringPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "OcspResponderUrl", "value": *sdk.OcspResponderUrl})
    }

    // Handling Primitives
    if !model.SignedBy.IsNull() && !model.SignedBy.IsUnknown() {
        sdk.SignedBy = model.SignedBy.ValueString()
        tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "SignedBy", "value": sdk.SignedBy})
    }

    // Handling Primitives
    if !model.Snippet.IsNull() && !model.Snippet.IsUnknown() {
        sdk.Snippet = model.Snippet.ValueStringPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
    }

    // Handling Primitives
    if !model.State.IsNull() && !model.State.IsUnknown() {
        sdk.State = model.State.ValueStringPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "State", "value": *sdk.State})
    }

diags.Append(d...)

    tflog.Debug(ctx, "Exiting unpack helper for models.CertificatesPost", map[string]interface{}{"has_errors": diags.HasError()})
    return &sdk, diags

}

// --- Packer for CertificatesPost ---
func packCertificatesPostFromSdk(ctx context.Context, sdk identity_services.CertificatesPost) (types.Object, diag.Diagnostics) {
    tflog.Debug(ctx, "Entering pack helper for models.CertificatesPost", map[string]interface{}{"sdk_struct": sdk})
    diags := diag.Diagnostics{}
    var model models.CertificatesPost
    var d diag.Diagnostics
    // Handling Objects
    // This is a regular nested object that has its own packer.
    // Logic for non-pointer / value-type nested objects
    if !reflect.ValueOf(sdk.Algorithm).IsZero() {
        tflog.Debug(ctx, "Packing nested object for field Algorithm")
        packed, d := packCertificatesPostAlgorithmFromSdk(ctx, sdk.Algorithm)
        diags.Append(d...)
        model.Algorithm = packed
    } else {
        model.Algorithm = basetypes.NewObjectNull(models.CertificatesPostAlgorithm{}.AttrTypes())
    }
    // Handling Lists
    if sdk.AlternateEmail != nil {
        tflog.Debug(ctx, "Packing list of primitives for field AlternateEmail")
        var d diag.Diagnostics
        // This logic now dynamically determines the element type based on the SDK's Go type.
        var elemType attr.Type = basetypes.StringType{} // Default to string
        model.AlternateEmail, d = basetypes.NewListValueFrom(ctx, elemType, sdk.AlternateEmail)
        diags.Append(d...)
    } else {
        // This logic now creates a correctly typed null list.
        var elemType attr.Type = basetypes.StringType{} // Default to string
        model.AlternateEmail = basetypes.NewListNull(elemType)
    }
    // Handling Primitives
    // Standard primitive packing
    model.CertificateName = basetypes.NewStringValue(sdk.CertificateName)
    tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "CertificateName", "value": sdk.CertificateName})
    // Handling Primitives
    // Standard primitive packing
    model.CommonName = basetypes.NewStringValue(sdk.CommonName)
    tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "CommonName", "value": sdk.CommonName})
    // Handling Primitives
    // Standard primitive packing
    if sdk.CountryCode != nil {
        model.CountryCode = basetypes.NewStringValue(*sdk.CountryCode)
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "CountryCode", "value": *sdk.CountryCode})
    } else {
        model.CountryCode = basetypes.NewStringNull()
    }
    // Handling Primitives
    // Standard primitive packing
    if sdk.DayTillExpiration != nil {
        model.DayTillExpiration = basetypes.NewInt64Value(int64(*sdk.DayTillExpiration))
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DayTillExpiration", "value": *sdk.DayTillExpiration})
    } else {
        model.DayTillExpiration = basetypes.NewInt64Null()
    }
    // Handling Lists
    if sdk.Department != nil {
        tflog.Debug(ctx, "Packing list of primitives for field Department")
        var d diag.Diagnostics
        // This logic now dynamically determines the element type based on the SDK's Go type.
        var elemType attr.Type = basetypes.StringType{} // Default to string
        model.Department, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Department)
        diags.Append(d...)
    } else {
        // This logic now creates a correctly typed null list.
        var elemType attr.Type = basetypes.StringType{} // Default to string
        model.Department = basetypes.NewListNull(elemType)
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
    model.Digest = basetypes.NewStringValue(sdk.Digest)
    tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Digest", "value": sdk.Digest})
    // Handling Primitives
    // Standard primitive packing
    if sdk.Email != nil {
        model.Email = basetypes.NewStringValue(*sdk.Email)
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Email", "value": *sdk.Email})
    } else {
        model.Email = basetypes.NewStringNull()
    }
    // Handling Primitives
    // Standard primitive packing
    if sdk.Folder != nil {
        model.Folder = basetypes.NewStringValue(*sdk.Folder)
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Folder", "value": *sdk.Folder})
    } else {
        model.Folder = basetypes.NewStringNull()
    }
    // Handling Lists
    if sdk.Hostname != nil {
        tflog.Debug(ctx, "Packing list of primitives for field Hostname")
        var d diag.Diagnostics
        // This logic now dynamically determines the element type based on the SDK's Go type.
        var elemType attr.Type = basetypes.StringType{} // Default to string
        model.Hostname, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Hostname)
        diags.Append(d...)
    } else {
        // This logic now creates a correctly typed null list.
        var elemType attr.Type = basetypes.StringType{} // Default to string
        model.Hostname = basetypes.NewListNull(elemType)
    }
    // Handling Lists
    if sdk.Ip != nil {
        tflog.Debug(ctx, "Packing list of primitives for field Ip")
        var d diag.Diagnostics
        // This logic now dynamically determines the element type based on the SDK's Go type.
        var elemType attr.Type = basetypes.StringType{} // Default to string
        model.Ip, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Ip)
        diags.Append(d...)
    } else {
        // This logic now creates a correctly typed null list.
        var elemType attr.Type = basetypes.StringType{} // Default to string
        model.Ip = basetypes.NewListNull(elemType)
    }
    // Handling Primitives
    // Standard primitive packing
    if sdk.IsBlockPrivateKey != nil {
        model.IsBlockPrivateKey = basetypes.NewBoolValue(*sdk.IsBlockPrivateKey)
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "IsBlockPrivateKey", "value": *sdk.IsBlockPrivateKey})
    } else {
        model.IsBlockPrivateKey = basetypes.NewBoolNull()
    }
    // Handling Primitives
    // Standard primitive packing
    if sdk.IsCertificateAuthority != nil {
        model.IsCertificateAuthority = basetypes.NewBoolValue(*sdk.IsCertificateAuthority)
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "IsCertificateAuthority", "value": *sdk.IsCertificateAuthority})
    } else {
        model.IsCertificateAuthority = basetypes.NewBoolNull()
    }
    // Handling Primitives
    // Standard primitive packing
    if sdk.Locality != nil {
        model.Locality = basetypes.NewStringValue(*sdk.Locality)
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Locality", "value": *sdk.Locality})
    } else {
        model.Locality = basetypes.NewStringNull()
    }
    // Handling Primitives
    // Standard primitive packing
    if sdk.OcspResponderUrl != nil {
        model.OcspResponderUrl = basetypes.NewStringValue(*sdk.OcspResponderUrl)
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "OcspResponderUrl", "value": *sdk.OcspResponderUrl})
    } else {
        model.OcspResponderUrl = basetypes.NewStringNull()
    }
    // Handling Primitives
    // Standard primitive packing
    model.SignedBy = basetypes.NewStringValue(sdk.SignedBy)
    tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "SignedBy", "value": sdk.SignedBy})
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
    if sdk.State != nil {
        model.State = basetypes.NewStringValue(*sdk.State)
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "State", "value": *sdk.State})
    } else {
        model.State = basetypes.NewStringNull()
    }
diags.Append(d...)

    obj, d := types.ObjectValueFrom(ctx, models.CertificatesPost{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
    diags.Append(d...)
    tflog.Debug(ctx, "Exiting pack helper for models.CertificatesPost", map[string]interface{}{"has_errors": diags.HasError()})
    return obj, diags

}

// --- List Unpacker for CertificatesPost ---
func unpackCertificatesPostListToSdk(ctx context.Context, list types.List) ([]identity_services.CertificatesPost, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.CertificatesPost")
	diags := diag.Diagnostics{}
	var data []models.CertificatesPost
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]identity_services.CertificatesPost, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.CertificatesPost{}.AttrTypes(), &item)
		unpacked, d := unpackCertificatesPostToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.CertificatesPost", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for CertificatesPost ---
func packCertificatesPostListFromSdk(ctx context.Context, sdks []identity_services.CertificatesPost) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.CertificatesPost")
	diags := diag.Diagnostics{}
	var data []models.CertificatesPost

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.CertificatesPost
		obj, d := packCertificatesPostFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.CertificatesPost{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.CertificatesPost", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.CertificatesPost{}.AttrType(), data)
}

// --- Unpacker for CertificatesPostAlgorithm ---
func unpackCertificatesPostAlgorithmToSdk(ctx context.Context, obj types.Object) (*identity_services.CertificatesPostAlgorithm, diag.Diagnostics) {
    tflog.Debug(ctx, "Entering unpack helper for models.CertificatesPostAlgorithm", map[string]interface{}{"tf_object": obj})
    diags := diag.Diagnostics{}
    var model models.CertificatesPostAlgorithm
    diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
    if diags.HasError() {
        tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
        return nil, diags
    }
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

    var sdk identity_services.CertificatesPostAlgorithm
    var d diag.Diagnostics
    // Handling Primitives
    if !model.EcdsaNumberOfBits.IsNull() && !model.EcdsaNumberOfBits.IsUnknown() {
        val := float32(model.EcdsaNumberOfBits.ValueFloat64())
        sdk.EcdsaNumberOfBits = &val
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "EcdsaNumberOfBits", "value": *sdk.EcdsaNumberOfBits})
    }

    // Handling Primitives
    if !model.RsaNumberOfBits.IsNull() && !model.RsaNumberOfBits.IsUnknown() {
        val := float32(model.RsaNumberOfBits.ValueFloat64())
        sdk.RsaNumberOfBits = &val
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "RsaNumberOfBits", "value": *sdk.RsaNumberOfBits})
    }

diags.Append(d...)

    tflog.Debug(ctx, "Exiting unpack helper for models.CertificatesPostAlgorithm", map[string]interface{}{"has_errors": diags.HasError()})
    return &sdk, diags

}

// --- Packer for CertificatesPostAlgorithm ---
func packCertificatesPostAlgorithmFromSdk(ctx context.Context, sdk identity_services.CertificatesPostAlgorithm) (types.Object, diag.Diagnostics) {
    tflog.Debug(ctx, "Entering pack helper for models.CertificatesPostAlgorithm", map[string]interface{}{"sdk_struct": sdk})
    diags := diag.Diagnostics{}
    var model models.CertificatesPostAlgorithm
    var d diag.Diagnostics
    // Handling Primitives
    // Standard primitive packing
    if sdk.EcdsaNumberOfBits != nil {
        model.EcdsaNumberOfBits = basetypes.NewFloat64Value(float64(*sdk.EcdsaNumberOfBits))
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "EcdsaNumberOfBits", "value": *sdk.EcdsaNumberOfBits})
    } else {
        model.EcdsaNumberOfBits = basetypes.NewFloat64Null()
    }
    // Handling Primitives
    // Standard primitive packing
    if sdk.RsaNumberOfBits != nil {
        model.RsaNumberOfBits = basetypes.NewFloat64Value(float64(*sdk.RsaNumberOfBits))
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "RsaNumberOfBits", "value": *sdk.RsaNumberOfBits})
    } else {
        model.RsaNumberOfBits = basetypes.NewFloat64Null()
    }
diags.Append(d...)

    obj, d := types.ObjectValueFrom(ctx, models.CertificatesPostAlgorithm{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
    diags.Append(d...)
    tflog.Debug(ctx, "Exiting pack helper for models.CertificatesPostAlgorithm", map[string]interface{}{"has_errors": diags.HasError()})
    return obj, diags

}

// --- List Unpacker for CertificatesPostAlgorithm ---
func unpackCertificatesPostAlgorithmListToSdk(ctx context.Context, list types.List) ([]identity_services.CertificatesPostAlgorithm, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.CertificatesPostAlgorithm")
	diags := diag.Diagnostics{}
	var data []models.CertificatesPostAlgorithm
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]identity_services.CertificatesPostAlgorithm, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.CertificatesPostAlgorithm{}.AttrTypes(), &item)
		unpacked, d := unpackCertificatesPostAlgorithmToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.CertificatesPostAlgorithm", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for CertificatesPostAlgorithm ---
func packCertificatesPostAlgorithmListFromSdk(ctx context.Context, sdks []identity_services.CertificatesPostAlgorithm) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.CertificatesPostAlgorithm")
	diags := diag.Diagnostics{}
	var data []models.CertificatesPostAlgorithm

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.CertificatesPostAlgorithm
		obj, d := packCertificatesPostAlgorithmFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.CertificatesPostAlgorithm{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.CertificatesPostAlgorithm", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.CertificatesPostAlgorithm{}.AttrType(), data)
}


*/
