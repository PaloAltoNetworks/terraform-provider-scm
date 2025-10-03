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










// --- Unpacker for AuthenticationPortals ---
func unpackAuthenticationPortalsToSdk(ctx context.Context, obj types.Object) (*identity_services.AuthenticationPortals, diag.Diagnostics) {
    tflog.Debug(ctx, "Entering unpack helper for models.AuthenticationPortals", map[string]interface{}{"tf_object": obj})
    diags := diag.Diagnostics{}
    var model models.AuthenticationPortals
    diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
    if diags.HasError() {
        tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
        return nil, diags
    }
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

    var sdk identity_services.AuthenticationPortals
    var d diag.Diagnostics
    // Handling Primitives
    if !model.AuthenticationProfile.IsNull() && !model.AuthenticationProfile.IsUnknown() {
        sdk.AuthenticationProfile = model.AuthenticationProfile.ValueStringPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "AuthenticationProfile", "value": *sdk.AuthenticationProfile})
    }

    // Handling Primitives
    if !model.CertificateProfile.IsNull() && !model.CertificateProfile.IsUnknown() {
        sdk.CertificateProfile = model.CertificateProfile.ValueStringPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "CertificateProfile", "value": *sdk.CertificateProfile})
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
    if !model.GpUdpPort.IsNull() && !model.GpUdpPort.IsUnknown() {
        val := int32(model.GpUdpPort.ValueInt64())
        sdk.GpUdpPort = &val
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "GpUdpPort", "value": *sdk.GpUdpPort})
    }

    // Handling Primitives
    if !model.Id.IsNull() && !model.Id.IsUnknown() {
        sdk.Id = model.Id.ValueString()
        tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Id", "value": sdk.Id})
    }

    // Handling Primitives
    if !model.IdleTimer.IsNull() && !model.IdleTimer.IsUnknown() {
        val := int32(model.IdleTimer.ValueInt64())
        sdk.IdleTimer = &val
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "IdleTimer", "value": *sdk.IdleTimer})
    }

    // Handling Primitives
    if !model.RedirectHost.IsNull() && !model.RedirectHost.IsUnknown() {
        sdk.RedirectHost = model.RedirectHost.ValueString()
        tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "RedirectHost", "value": sdk.RedirectHost})
    }

    // Handling Primitives
    if !model.Snippet.IsNull() && !model.Snippet.IsUnknown() {
        sdk.Snippet = model.Snippet.ValueStringPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
    }

    // Handling Primitives
    if !model.Timer.IsNull() && !model.Timer.IsUnknown() {
        val := int32(model.Timer.ValueInt64())
        sdk.Timer = &val
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Timer", "value": *sdk.Timer})
    }

    // Handling Primitives
    if !model.TlsServiceProfile.IsNull() && !model.TlsServiceProfile.IsUnknown() {
        sdk.TlsServiceProfile = model.TlsServiceProfile.ValueStringPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "TlsServiceProfile", "value": *sdk.TlsServiceProfile})
    }

diags.Append(d...)

    tflog.Debug(ctx, "Exiting unpack helper for models.AuthenticationPortals", map[string]interface{}{"has_errors": diags.HasError()})
    return &sdk, diags

}

// --- Packer for AuthenticationPortals ---
func packAuthenticationPortalsFromSdk(ctx context.Context, sdk identity_services.AuthenticationPortals) (types.Object, diag.Diagnostics) {
    tflog.Debug(ctx, "Entering pack helper for models.AuthenticationPortals", map[string]interface{}{"sdk_struct": sdk})
    diags := diag.Diagnostics{}
    var model models.AuthenticationPortals
    var d diag.Diagnostics
    // Handling Primitives
    // Standard primitive packing
    if sdk.AuthenticationProfile != nil {
        model.AuthenticationProfile = basetypes.NewStringValue(*sdk.AuthenticationProfile)
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "AuthenticationProfile", "value": *sdk.AuthenticationProfile})
    } else {
        model.AuthenticationProfile = basetypes.NewStringNull()
    }
    // Handling Primitives
    // Standard primitive packing
    if sdk.CertificateProfile != nil {
        model.CertificateProfile = basetypes.NewStringValue(*sdk.CertificateProfile)
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "CertificateProfile", "value": *sdk.CertificateProfile})
    } else {
        model.CertificateProfile = basetypes.NewStringNull()
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
    if sdk.GpUdpPort != nil {
        model.GpUdpPort = basetypes.NewInt64Value(int64(*sdk.GpUdpPort))
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "GpUdpPort", "value": *sdk.GpUdpPort})
    } else {
        model.GpUdpPort = basetypes.NewInt64Null()
    }
    // Handling Primitives
    // Standard primitive packing
    model.Id = basetypes.NewStringValue(sdk.Id)
    tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Id", "value": sdk.Id})
    // Handling Primitives
    // Standard primitive packing
    if sdk.IdleTimer != nil {
        model.IdleTimer = basetypes.NewInt64Value(int64(*sdk.IdleTimer))
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "IdleTimer", "value": *sdk.IdleTimer})
    } else {
        model.IdleTimer = basetypes.NewInt64Null()
    }
    // Handling Primitives
    // Standard primitive packing
    model.RedirectHost = basetypes.NewStringValue(sdk.RedirectHost)
    tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "RedirectHost", "value": sdk.RedirectHost})
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
    if sdk.Timer != nil {
        model.Timer = basetypes.NewInt64Value(int64(*sdk.Timer))
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Timer", "value": *sdk.Timer})
    } else {
        model.Timer = basetypes.NewInt64Null()
    }
    // Handling Primitives
    // Standard primitive packing
    if sdk.TlsServiceProfile != nil {
        model.TlsServiceProfile = basetypes.NewStringValue(*sdk.TlsServiceProfile)
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "TlsServiceProfile", "value": *sdk.TlsServiceProfile})
    } else {
        model.TlsServiceProfile = basetypes.NewStringNull()
    }
diags.Append(d...)

    obj, d := types.ObjectValueFrom(ctx, models.AuthenticationPortals{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
    diags.Append(d...)
    tflog.Debug(ctx, "Exiting pack helper for models.AuthenticationPortals", map[string]interface{}{"has_errors": diags.HasError()})
    return obj, diags

}

// --- List Unpacker for AuthenticationPortals ---
func unpackAuthenticationPortalsListToSdk(ctx context.Context, list types.List) ([]identity_services.AuthenticationPortals, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.AuthenticationPortals")
	diags := diag.Diagnostics{}
	var data []models.AuthenticationPortals
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]identity_services.AuthenticationPortals, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.AuthenticationPortals{}.AttrTypes(), &item)
		unpacked, d := unpackAuthenticationPortalsToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.AuthenticationPortals", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for AuthenticationPortals ---
func packAuthenticationPortalsListFromSdk(ctx context.Context, sdks []identity_services.AuthenticationPortals) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.AuthenticationPortals")
	diags := diag.Diagnostics{}
	var data []models.AuthenticationPortals

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.AuthenticationPortals
		obj, d := packAuthenticationPortalsFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.AuthenticationPortals{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.AuthenticationPortals", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.AuthenticationPortals{}.AttrType(), data)
}


*/
