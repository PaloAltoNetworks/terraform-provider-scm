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

	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/network_services"
	"github.com/paloaltonetworks/scm-go/generated/network_services"
)










// --- Unpacker for SslDecryptionSettings ---
func unpackSslDecryptionSettingsToSdk(ctx context.Context, obj types.Object) (*network_services.SslDecryptionSettings, diag.Diagnostics) {
    tflog.Debug(ctx, "Entering unpack helper for models.SslDecryptionSettings", map[string]interface{}{"tf_object": obj})
    diags := diag.Diagnostics{}
    var model models.SslDecryptionSettings
    diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
    if diags.HasError() {
        tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
        return nil, diags
    }
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

    var sdk network_services.SslDecryptionSettings
    var d diag.Diagnostics
    // Handling Lists
    if !model.DisabledSslExcludeCertFromPredefined.IsNull() && !model.DisabledSslExcludeCertFromPredefined.IsUnknown() {
        tflog.Debug(ctx, "Unpacking list of primitives for field DisabledSslExcludeCertFromPredefined")
        diags.Append(model.DisabledSslExcludeCertFromPredefined.ElementsAs(ctx, &sdk.DisabledSslExcludeCertFromPredefined, false)...)
    }

    // Handling Objects
    if !model.ForwardTrustCertificate.IsNull() && !model.ForwardTrustCertificate.IsUnknown() {
        tflog.Debug(ctx, "Unpacking nested object for field ForwardTrustCertificate")
        unpacked, d := unpackSslDecryptionSettingsForwardTrustCertificateToSdk(ctx, model.ForwardTrustCertificate)
        diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "ForwardTrustCertificate"})
		}
        if unpacked != nil {
            sdk.ForwardTrustCertificate = unpacked
        }
    }

    // Handling Objects
    if !model.ForwardUntrustCertificate.IsNull() && !model.ForwardUntrustCertificate.IsUnknown() {
        tflog.Debug(ctx, "Unpacking nested object for field ForwardUntrustCertificate")
        unpacked, d := unpackSslDecryptionSettingsForwardTrustCertificateToSdk(ctx, model.ForwardUntrustCertificate)
        diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "ForwardUntrustCertificate"})
		}
        if unpacked != nil {
            sdk.ForwardUntrustCertificate = unpacked
        }
    }

    // Handling Lists
    if !model.RootCaExcludeList.IsNull() && !model.RootCaExcludeList.IsUnknown() {
        tflog.Debug(ctx, "Unpacking list of primitives for field RootCaExcludeList")
        diags.Append(model.RootCaExcludeList.ElementsAs(ctx, &sdk.RootCaExcludeList, false)...)
    }

    // Handling Lists
    if !model.SslExcludeCert.IsNull() && !model.SslExcludeCert.IsUnknown() {
        tflog.Debug(ctx, "Unpacking list of objects for field SslExcludeCert")
        unpacked, d := unpackSslDecryptionSettingsSslExcludeCertInnerListToSdk(ctx, model.SslExcludeCert)
        diags.Append(d...)
        sdk.SslExcludeCert = unpacked
    }

    // Handling Lists
    if !model.TrustedRootCA.IsNull() && !model.TrustedRootCA.IsUnknown() {
        tflog.Debug(ctx, "Unpacking list of primitives for field TrustedRootCA")
        diags.Append(model.TrustedRootCA.ElementsAs(ctx, &sdk.TrustedRootCA, false)...)
    }

diags.Append(d...)

    tflog.Debug(ctx, "Exiting unpack helper for models.SslDecryptionSettings", map[string]interface{}{"has_errors": diags.HasError()})
    return &sdk, diags

}

// --- Packer for SslDecryptionSettings ---
func packSslDecryptionSettingsFromSdk(ctx context.Context, sdk network_services.SslDecryptionSettings) (types.Object, diag.Diagnostics) {
    tflog.Debug(ctx, "Entering pack helper for models.SslDecryptionSettings", map[string]interface{}{"sdk_struct": sdk})
    diags := diag.Diagnostics{}
    var model models.SslDecryptionSettings
    var d diag.Diagnostics
    // Handling Lists
    if sdk.DisabledSslExcludeCertFromPredefined != nil {
        tflog.Debug(ctx, "Packing list of primitives for field DisabledSslExcludeCertFromPredefined")
        var d diag.Diagnostics
        // This logic now dynamically determines the element type based on the SDK's Go type.
        var elemType attr.Type = basetypes.StringType{} // Default to string
        model.DisabledSslExcludeCertFromPredefined, d = basetypes.NewListValueFrom(ctx, elemType, sdk.DisabledSslExcludeCertFromPredefined)
        diags.Append(d...)
    } else {
        // This logic now creates a correctly typed null list.
        var elemType attr.Type = basetypes.StringType{} // Default to string
        model.DisabledSslExcludeCertFromPredefined = basetypes.NewListNull(elemType)
    }
    // Handling Objects
    // This is a regular nested object that has its own packer.
    if sdk.ForwardTrustCertificate != nil {
        tflog.Debug(ctx, "Packing nested object for field ForwardTrustCertificate")
        packed, d := packSslDecryptionSettingsForwardTrustCertificateFromSdk(ctx, *sdk.ForwardTrustCertificate)
        diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "ForwardTrustCertificate"})
		}
        model.ForwardTrustCertificate = packed
    } else {
        model.ForwardTrustCertificate = basetypes.NewObjectNull(models.SslDecryptionSettingsForwardTrustCertificate{}.AttrTypes())
    }
    // Handling Objects
    // This is a regular nested object that has its own packer.
    if sdk.ForwardUntrustCertificate != nil {
        tflog.Debug(ctx, "Packing nested object for field ForwardUntrustCertificate")
        packed, d := packSslDecryptionSettingsForwardTrustCertificateFromSdk(ctx, *sdk.ForwardUntrustCertificate)
        diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "ForwardUntrustCertificate"})
		}
        model.ForwardUntrustCertificate = packed
    } else {
        model.ForwardUntrustCertificate = basetypes.NewObjectNull(models.SslDecryptionSettingsForwardTrustCertificate{}.AttrTypes())
    }
    // Handling Lists
    if sdk.RootCaExcludeList != nil {
        tflog.Debug(ctx, "Packing list of primitives for field RootCaExcludeList")
        var d diag.Diagnostics
        // This logic now dynamically determines the element type based on the SDK's Go type.
        var elemType attr.Type = basetypes.StringType{} // Default to string
        model.RootCaExcludeList, d = basetypes.NewListValueFrom(ctx, elemType, sdk.RootCaExcludeList)
        diags.Append(d...)
    } else {
        // This logic now creates a correctly typed null list.
        var elemType attr.Type = basetypes.StringType{} // Default to string
        model.RootCaExcludeList = basetypes.NewListNull(elemType)
    }
    // Handling Lists
    if sdk.SslExcludeCert != nil {
        tflog.Debug(ctx, "Packing list of objects for field SslExcludeCert")
        packed, d := packSslDecryptionSettingsSslExcludeCertInnerListFromSdk(ctx, sdk.SslExcludeCert)
        diags.Append(d...)
        model.SslExcludeCert = packed
    } else {
        model.SslExcludeCert = basetypes.NewListNull(models.SslDecryptionSettingsSslExcludeCertInner{}.AttrType())
    }
    // Handling Lists
    if sdk.TrustedRootCA != nil {
        tflog.Debug(ctx, "Packing list of primitives for field TrustedRootCA")
        var d diag.Diagnostics
        // This logic now dynamically determines the element type based on the SDK's Go type.
        var elemType attr.Type = basetypes.StringType{} // Default to string
        model.TrustedRootCA, d = basetypes.NewListValueFrom(ctx, elemType, sdk.TrustedRootCA)
        diags.Append(d...)
    } else {
        // This logic now creates a correctly typed null list.
        var elemType attr.Type = basetypes.StringType{} // Default to string
        model.TrustedRootCA = basetypes.NewListNull(elemType)
    }
diags.Append(d...)

    obj, d := types.ObjectValueFrom(ctx, models.SslDecryptionSettings{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
    diags.Append(d...)
    tflog.Debug(ctx, "Exiting pack helper for models.SslDecryptionSettings", map[string]interface{}{"has_errors": diags.HasError()})
    return obj, diags

}

// --- List Unpacker for SslDecryptionSettings ---
func unpackSslDecryptionSettingsListToSdk(ctx context.Context, list types.List) ([]network_services.SslDecryptionSettings, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.SslDecryptionSettings")
	diags := diag.Diagnostics{}
	var data []models.SslDecryptionSettings
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.SslDecryptionSettings, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.SslDecryptionSettings{}.AttrTypes(), &item)
		unpacked, d := unpackSslDecryptionSettingsToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.SslDecryptionSettings", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for SslDecryptionSettings ---
func packSslDecryptionSettingsListFromSdk(ctx context.Context, sdks []network_services.SslDecryptionSettings) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.SslDecryptionSettings")
	diags := diag.Diagnostics{}
	var data []models.SslDecryptionSettings

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.SslDecryptionSettings
		obj, d := packSslDecryptionSettingsFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.SslDecryptionSettings{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.SslDecryptionSettings", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.SslDecryptionSettings{}.AttrType(), data)
}

// --- Unpacker for SslDecryptionSettingsForwardTrustCertificate ---
func unpackSslDecryptionSettingsForwardTrustCertificateToSdk(ctx context.Context, obj types.Object) (*network_services.SslDecryptionSettingsForwardTrustCertificate, diag.Diagnostics) {
    tflog.Debug(ctx, "Entering unpack helper for models.SslDecryptionSettingsForwardTrustCertificate", map[string]interface{}{"tf_object": obj})
    diags := diag.Diagnostics{}
    var model models.SslDecryptionSettingsForwardTrustCertificate
    diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
    if diags.HasError() {
        tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
        return nil, diags
    }
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

    var sdk network_services.SslDecryptionSettingsForwardTrustCertificate
    var d diag.Diagnostics
    // Handling Primitives
    if !model.Ecdsa.IsNull() && !model.Ecdsa.IsUnknown() {
        sdk.Ecdsa = model.Ecdsa.ValueStringPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Ecdsa", "value": *sdk.Ecdsa})
    }

    // Handling Primitives
    if !model.Rsa.IsNull() && !model.Rsa.IsUnknown() {
        sdk.Rsa = model.Rsa.ValueStringPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Rsa", "value": *sdk.Rsa})
    }

diags.Append(d...)

    tflog.Debug(ctx, "Exiting unpack helper for models.SslDecryptionSettingsForwardTrustCertificate", map[string]interface{}{"has_errors": diags.HasError()})
    return &sdk, diags

}

// --- Packer for SslDecryptionSettingsForwardTrustCertificate ---
func packSslDecryptionSettingsForwardTrustCertificateFromSdk(ctx context.Context, sdk network_services.SslDecryptionSettingsForwardTrustCertificate) (types.Object, diag.Diagnostics) {
    tflog.Debug(ctx, "Entering pack helper for models.SslDecryptionSettingsForwardTrustCertificate", map[string]interface{}{"sdk_struct": sdk})
    diags := diag.Diagnostics{}
    var model models.SslDecryptionSettingsForwardTrustCertificate
    var d diag.Diagnostics
    // Handling Primitives
    // Standard primitive packing
    if sdk.Ecdsa != nil {
        model.Ecdsa = basetypes.NewStringValue(*sdk.Ecdsa)
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Ecdsa", "value": *sdk.Ecdsa})
    } else {
        model.Ecdsa = basetypes.NewStringNull()
    }
    // Handling Primitives
    // Standard primitive packing
    if sdk.Rsa != nil {
        model.Rsa = basetypes.NewStringValue(*sdk.Rsa)
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Rsa", "value": *sdk.Rsa})
    } else {
        model.Rsa = basetypes.NewStringNull()
    }
diags.Append(d...)

    obj, d := types.ObjectValueFrom(ctx, models.SslDecryptionSettingsForwardTrustCertificate{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
    diags.Append(d...)
    tflog.Debug(ctx, "Exiting pack helper for models.SslDecryptionSettingsForwardTrustCertificate", map[string]interface{}{"has_errors": diags.HasError()})
    return obj, diags

}

// --- List Unpacker for SslDecryptionSettingsForwardTrustCertificate ---
func unpackSslDecryptionSettingsForwardTrustCertificateListToSdk(ctx context.Context, list types.List) ([]network_services.SslDecryptionSettingsForwardTrustCertificate, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.SslDecryptionSettingsForwardTrustCertificate")
	diags := diag.Diagnostics{}
	var data []models.SslDecryptionSettingsForwardTrustCertificate
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.SslDecryptionSettingsForwardTrustCertificate, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.SslDecryptionSettingsForwardTrustCertificate{}.AttrTypes(), &item)
		unpacked, d := unpackSslDecryptionSettingsForwardTrustCertificateToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.SslDecryptionSettingsForwardTrustCertificate", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for SslDecryptionSettingsForwardTrustCertificate ---
func packSslDecryptionSettingsForwardTrustCertificateListFromSdk(ctx context.Context, sdks []network_services.SslDecryptionSettingsForwardTrustCertificate) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.SslDecryptionSettingsForwardTrustCertificate")
	diags := diag.Diagnostics{}
	var data []models.SslDecryptionSettingsForwardTrustCertificate

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.SslDecryptionSettingsForwardTrustCertificate
		obj, d := packSslDecryptionSettingsForwardTrustCertificateFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.SslDecryptionSettingsForwardTrustCertificate{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.SslDecryptionSettingsForwardTrustCertificate", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.SslDecryptionSettingsForwardTrustCertificate{}.AttrType(), data)
}

// --- Unpacker for SslDecryptionSettingsSslExcludeCertInner ---
func unpackSslDecryptionSettingsSslExcludeCertInnerToSdk(ctx context.Context, obj types.Object) (*network_services.SslDecryptionSettingsSslExcludeCertInner, diag.Diagnostics) {
    tflog.Debug(ctx, "Entering unpack helper for models.SslDecryptionSettingsSslExcludeCertInner", map[string]interface{}{"tf_object": obj})
    diags := diag.Diagnostics{}
    var model models.SslDecryptionSettingsSslExcludeCertInner
    diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
    if diags.HasError() {
        tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
        return nil, diags
    }
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

    var sdk network_services.SslDecryptionSettingsSslExcludeCertInner
    var d diag.Diagnostics
    // Handling Primitives
    if !model.Description.IsNull() && !model.Description.IsUnknown() {
        sdk.Description = model.Description.ValueStringPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Description", "value": *sdk.Description})
    }

    // Handling Primitives
    if !model.Exclude.IsNull() && !model.Exclude.IsUnknown() {
        sdk.Exclude = model.Exclude.ValueBoolPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Exclude", "value": *sdk.Exclude})
    }

    // Handling Primitives
    if !model.Name.IsNull() && !model.Name.IsUnknown() {
        sdk.Name = model.Name.ValueStringPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
    }

diags.Append(d...)

    tflog.Debug(ctx, "Exiting unpack helper for models.SslDecryptionSettingsSslExcludeCertInner", map[string]interface{}{"has_errors": diags.HasError()})
    return &sdk, diags

}

// --- Packer for SslDecryptionSettingsSslExcludeCertInner ---
func packSslDecryptionSettingsSslExcludeCertInnerFromSdk(ctx context.Context, sdk network_services.SslDecryptionSettingsSslExcludeCertInner) (types.Object, diag.Diagnostics) {
    tflog.Debug(ctx, "Entering pack helper for models.SslDecryptionSettingsSslExcludeCertInner", map[string]interface{}{"sdk_struct": sdk})
    diags := diag.Diagnostics{}
    var model models.SslDecryptionSettingsSslExcludeCertInner
    var d diag.Diagnostics
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
    if sdk.Exclude != nil {
        model.Exclude = basetypes.NewBoolValue(*sdk.Exclude)
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Exclude", "value": *sdk.Exclude})
    } else {
        model.Exclude = basetypes.NewBoolNull()
    }
    // Handling Primitives
    // Standard primitive packing
    if sdk.Name != nil {
        model.Name = basetypes.NewStringValue(*sdk.Name)
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
    } else {
        model.Name = basetypes.NewStringNull()
    }
diags.Append(d...)

    obj, d := types.ObjectValueFrom(ctx, models.SslDecryptionSettingsSslExcludeCertInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
    diags.Append(d...)
    tflog.Debug(ctx, "Exiting pack helper for models.SslDecryptionSettingsSslExcludeCertInner", map[string]interface{}{"has_errors": diags.HasError()})
    return obj, diags

}

// --- List Unpacker for SslDecryptionSettingsSslExcludeCertInner ---
func unpackSslDecryptionSettingsSslExcludeCertInnerListToSdk(ctx context.Context, list types.List) ([]network_services.SslDecryptionSettingsSslExcludeCertInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.SslDecryptionSettingsSslExcludeCertInner")
	diags := diag.Diagnostics{}
	var data []models.SslDecryptionSettingsSslExcludeCertInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.SslDecryptionSettingsSslExcludeCertInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.SslDecryptionSettingsSslExcludeCertInner{}.AttrTypes(), &item)
		unpacked, d := unpackSslDecryptionSettingsSslExcludeCertInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.SslDecryptionSettingsSslExcludeCertInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for SslDecryptionSettingsSslExcludeCertInner ---
func packSslDecryptionSettingsSslExcludeCertInnerListFromSdk(ctx context.Context, sdks []network_services.SslDecryptionSettingsSslExcludeCertInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.SslDecryptionSettingsSslExcludeCertInner")
	diags := diag.Diagnostics{}
	var data []models.SslDecryptionSettingsSslExcludeCertInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.SslDecryptionSettingsSslExcludeCertInner
		obj, d := packSslDecryptionSettingsSslExcludeCertInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.SslDecryptionSettingsSslExcludeCertInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.SslDecryptionSettingsSslExcludeCertInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.SslDecryptionSettingsSslExcludeCertInner{}.AttrType(), data)
}


*/
