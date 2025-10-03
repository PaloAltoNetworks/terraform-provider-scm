package provider

import (
	"context"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/identity_services"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/identity_services"
)

// --- Unpacker for RadiusServerProfiles ---
func unpackRadiusServerProfilesToSdk(ctx context.Context, obj types.Object) (*identity_services.RadiusServerProfiles, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.RadiusServerProfiles", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.RadiusServerProfiles
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk identity_services.RadiusServerProfiles
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

	// Handling Objects
	if !model.Protocol.IsNull() && !model.Protocol.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Protocol")
		unpacked, d := unpackRadiusServerProfilesProtocolToSdk(ctx, model.Protocol)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Protocol"})
		}
		if unpacked != nil {
			sdk.Protocol = *unpacked
		}
	}

	// Handling Primitives
	if !model.Retries.IsNull() && !model.Retries.IsUnknown() {
		val := int32(model.Retries.ValueInt64())
		sdk.Retries = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Retries", "value": *sdk.Retries})
	}

	// Handling Lists
	if !model.Server.IsNull() && !model.Server.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field Server")
		unpacked, d := unpackRadiusServerProfilesServerInnerListToSdk(ctx, model.Server)
		diags.Append(d...)
		sdk.Server = unpacked
	}

	// Handling Primitives
	if !model.Snippet.IsNull() && !model.Snippet.IsUnknown() {
		sdk.Snippet = model.Snippet.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	}

	// Handling Primitives
	if !model.Timeout.IsNull() && !model.Timeout.IsUnknown() {
		val := int32(model.Timeout.ValueInt64())
		sdk.Timeout = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Timeout", "value": *sdk.Timeout})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.RadiusServerProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for RadiusServerProfiles ---
func packRadiusServerProfilesFromSdk(ctx context.Context, sdk identity_services.RadiusServerProfiles) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.RadiusServerProfiles", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.RadiusServerProfiles
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
	// Handling Objects
	// This is a regular nested object that has its own packer.
	// Logic for non-pointer / value-type nested objects
	if !reflect.ValueOf(sdk.Protocol).IsZero() {
		tflog.Debug(ctx, "Packing nested object for field Protocol")
		packed, d := packRadiusServerProfilesProtocolFromSdk(ctx, sdk.Protocol)
		diags.Append(d...)
		model.Protocol = packed
	} else {
		model.Protocol = basetypes.NewObjectNull(models.RadiusServerProfilesProtocol{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Retries != nil {
		model.Retries = basetypes.NewInt64Value(int64(*sdk.Retries))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Retries", "value": *sdk.Retries})
	} else {
		model.Retries = basetypes.NewInt64Null()
	}
	// Handling Lists
	if sdk.Server != nil {
		tflog.Debug(ctx, "Packing list of objects for field Server")
		packed, d := packRadiusServerProfilesServerInnerListFromSdk(ctx, sdk.Server)
		diags.Append(d...)
		model.Server = packed
	} else {
		model.Server = basetypes.NewListNull(models.RadiusServerProfilesServerInner{}.AttrType())
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
	if sdk.Timeout != nil {
		model.Timeout = basetypes.NewInt64Value(int64(*sdk.Timeout))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Timeout", "value": *sdk.Timeout})
	} else {
		model.Timeout = basetypes.NewInt64Null()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.RadiusServerProfiles{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.RadiusServerProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for RadiusServerProfiles ---
func unpackRadiusServerProfilesListToSdk(ctx context.Context, list types.List) ([]identity_services.RadiusServerProfiles, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.RadiusServerProfiles")
	diags := diag.Diagnostics{}
	var data []models.RadiusServerProfiles
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]identity_services.RadiusServerProfiles, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.RadiusServerProfiles{}.AttrTypes(), &item)
		unpacked, d := unpackRadiusServerProfilesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.RadiusServerProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for RadiusServerProfiles ---
func packRadiusServerProfilesListFromSdk(ctx context.Context, sdks []identity_services.RadiusServerProfiles) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.RadiusServerProfiles")
	diags := diag.Diagnostics{}
	var data []models.RadiusServerProfiles

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.RadiusServerProfiles
		obj, d := packRadiusServerProfilesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.RadiusServerProfiles{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.RadiusServerProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.RadiusServerProfiles{}.AttrType(), data)
}

// --- Unpacker for RadiusServerProfilesProtocol ---
func unpackRadiusServerProfilesProtocolToSdk(ctx context.Context, obj types.Object) (*identity_services.RadiusServerProfilesProtocol, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.RadiusServerProfilesProtocol", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.RadiusServerProfilesProtocol
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk identity_services.RadiusServerProfilesProtocol
	var d diag.Diagnostics
	// Handling Typeless Objects
	if !model.CHAP.IsNull() && !model.CHAP.IsUnknown() {
		tflog.Debug(ctx, "Unpacking typeless object for field CHAP")
		sdk.CHAP = make(map[string]interface{})
	}

	// Handling Objects
	if !model.EAPTTLSWithPAP.IsNull() && !model.EAPTTLSWithPAP.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field EAPTTLSWithPAP")
		unpacked, d := unpackRadiusServerProfilesProtocolEAPTTLSWithPAPToSdk(ctx, model.EAPTTLSWithPAP)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "EAPTTLSWithPAP"})
		}
		if unpacked != nil {
			sdk.EAPTTLSWithPAP = unpacked
		}
	}

	// Handling Typeless Objects
	if !model.PAP.IsNull() && !model.PAP.IsUnknown() {
		tflog.Debug(ctx, "Unpacking typeless object for field PAP")
		sdk.PAP = make(map[string]interface{})
	}

	// Handling Objects
	if !model.PEAPMSCHAPv2.IsNull() && !model.PEAPMSCHAPv2.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field PEAPMSCHAPv2")
		unpacked, d := unpackRadiusServerProfilesProtocolPEAPMSCHAPv2ToSdk(ctx, model.PEAPMSCHAPv2)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "PEAPMSCHAPv2"})
		}
		if unpacked != nil {
			sdk.PEAPMSCHAPv2 = unpacked
		}
	}

	// Handling Objects
	if !model.PEAPWithGTC.IsNull() && !model.PEAPWithGTC.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field PEAPWithGTC")
		unpacked, d := unpackRadiusServerProfilesProtocolEAPTTLSWithPAPToSdk(ctx, model.PEAPWithGTC)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "PEAPWithGTC"})
		}
		if unpacked != nil {
			sdk.PEAPWithGTC = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.RadiusServerProfilesProtocol", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for RadiusServerProfilesProtocol ---
func packRadiusServerProfilesProtocolFromSdk(ctx context.Context, sdk identity_services.RadiusServerProfilesProtocol) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.RadiusServerProfilesProtocol", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.RadiusServerProfilesProtocol
	var d diag.Diagnostics
	// Handling Objects
	// This is a marker object (e.g. CHAP: {}). We just need to create an empty, non-null object.
	if sdk.CHAP != nil && !reflect.ValueOf(sdk.CHAP).IsNil() {
		tflog.Debug(ctx, "Packing typeless object for field CHAP")
		var d diag.Diagnostics
		// Create an empty object with no attributes, which signifies its presence.
		model.CHAP, d = basetypes.NewObjectValue(map[string]attr.Type{}, map[string]attr.Value{})
		diags.Append(d...)
	} else {
		// Since this field is part of a oneOf, being nil means it's not selected.
		// We make the object null with an empty attribute map.
		model.CHAP = basetypes.NewObjectNull(map[string]attr.Type{})
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.EAPTTLSWithPAP != nil {
		tflog.Debug(ctx, "Packing nested object for field EAPTTLSWithPAP")
		packed, d := packRadiusServerProfilesProtocolEAPTTLSWithPAPFromSdk(ctx, *sdk.EAPTTLSWithPAP)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "EAPTTLSWithPAP"})
		}
		model.EAPTTLSWithPAP = packed
	} else {
		model.EAPTTLSWithPAP = basetypes.NewObjectNull(models.RadiusServerProfilesProtocolEAPTTLSWithPAP{}.AttrTypes())
	}
	// Handling Objects
	// This is a marker object (e.g. CHAP: {}). We just need to create an empty, non-null object.
	if sdk.PAP != nil && !reflect.ValueOf(sdk.PAP).IsNil() {
		tflog.Debug(ctx, "Packing typeless object for field PAP")
		var d diag.Diagnostics
		// Create an empty object with no attributes, which signifies its presence.
		model.PAP, d = basetypes.NewObjectValue(map[string]attr.Type{}, map[string]attr.Value{})
		diags.Append(d...)
	} else {
		// Since this field is part of a oneOf, being nil means it's not selected.
		// We make the object null with an empty attribute map.
		model.PAP = basetypes.NewObjectNull(map[string]attr.Type{})
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.PEAPMSCHAPv2 != nil {
		tflog.Debug(ctx, "Packing nested object for field PEAPMSCHAPv2")
		packed, d := packRadiusServerProfilesProtocolPEAPMSCHAPv2FromSdk(ctx, *sdk.PEAPMSCHAPv2)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "PEAPMSCHAPv2"})
		}
		model.PEAPMSCHAPv2 = packed
	} else {
		model.PEAPMSCHAPv2 = basetypes.NewObjectNull(models.RadiusServerProfilesProtocolPEAPMSCHAPv2{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.PEAPWithGTC != nil {
		tflog.Debug(ctx, "Packing nested object for field PEAPWithGTC")
		packed, d := packRadiusServerProfilesProtocolEAPTTLSWithPAPFromSdk(ctx, *sdk.PEAPWithGTC)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "PEAPWithGTC"})
		}
		model.PEAPWithGTC = packed
	} else {
		model.PEAPWithGTC = basetypes.NewObjectNull(models.RadiusServerProfilesProtocolEAPTTLSWithPAP{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.RadiusServerProfilesProtocol{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.RadiusServerProfilesProtocol", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for RadiusServerProfilesProtocol ---
func unpackRadiusServerProfilesProtocolListToSdk(ctx context.Context, list types.List) ([]identity_services.RadiusServerProfilesProtocol, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.RadiusServerProfilesProtocol")
	diags := diag.Diagnostics{}
	var data []models.RadiusServerProfilesProtocol
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]identity_services.RadiusServerProfilesProtocol, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.RadiusServerProfilesProtocol{}.AttrTypes(), &item)
		unpacked, d := unpackRadiusServerProfilesProtocolToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.RadiusServerProfilesProtocol", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for RadiusServerProfilesProtocol ---
func packRadiusServerProfilesProtocolListFromSdk(ctx context.Context, sdks []identity_services.RadiusServerProfilesProtocol) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.RadiusServerProfilesProtocol")
	diags := diag.Diagnostics{}
	var data []models.RadiusServerProfilesProtocol

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.RadiusServerProfilesProtocol
		obj, d := packRadiusServerProfilesProtocolFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.RadiusServerProfilesProtocol{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.RadiusServerProfilesProtocol", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.RadiusServerProfilesProtocol{}.AttrType(), data)
}

// --- Unpacker for RadiusServerProfilesProtocolEAPTTLSWithPAP ---
func unpackRadiusServerProfilesProtocolEAPTTLSWithPAPToSdk(ctx context.Context, obj types.Object) (*identity_services.RadiusServerProfilesProtocolEAPTTLSWithPAP, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.RadiusServerProfilesProtocolEAPTTLSWithPAP", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.RadiusServerProfilesProtocolEAPTTLSWithPAP
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk identity_services.RadiusServerProfilesProtocolEAPTTLSWithPAP
	var d diag.Diagnostics
	// Handling Primitives
	if !model.AnonOuterId.IsNull() && !model.AnonOuterId.IsUnknown() {
		sdk.AnonOuterId = model.AnonOuterId.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "AnonOuterId", "value": *sdk.AnonOuterId})
	}

	// Handling Primitives
	if !model.RadiusCertProfile.IsNull() && !model.RadiusCertProfile.IsUnknown() {
		sdk.RadiusCertProfile = model.RadiusCertProfile.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "RadiusCertProfile", "value": *sdk.RadiusCertProfile})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.RadiusServerProfilesProtocolEAPTTLSWithPAP", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for RadiusServerProfilesProtocolEAPTTLSWithPAP ---
func packRadiusServerProfilesProtocolEAPTTLSWithPAPFromSdk(ctx context.Context, sdk identity_services.RadiusServerProfilesProtocolEAPTTLSWithPAP) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.RadiusServerProfilesProtocolEAPTTLSWithPAP", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.RadiusServerProfilesProtocolEAPTTLSWithPAP
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.AnonOuterId != nil {
		model.AnonOuterId = basetypes.NewBoolValue(*sdk.AnonOuterId)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "AnonOuterId", "value": *sdk.AnonOuterId})
	} else {
		model.AnonOuterId = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.RadiusCertProfile != nil {
		model.RadiusCertProfile = basetypes.NewStringValue(*sdk.RadiusCertProfile)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "RadiusCertProfile", "value": *sdk.RadiusCertProfile})
	} else {
		model.RadiusCertProfile = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.RadiusServerProfilesProtocolEAPTTLSWithPAP{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.RadiusServerProfilesProtocolEAPTTLSWithPAP", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for RadiusServerProfilesProtocolEAPTTLSWithPAP ---
func unpackRadiusServerProfilesProtocolEAPTTLSWithPAPListToSdk(ctx context.Context, list types.List) ([]identity_services.RadiusServerProfilesProtocolEAPTTLSWithPAP, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.RadiusServerProfilesProtocolEAPTTLSWithPAP")
	diags := diag.Diagnostics{}
	var data []models.RadiusServerProfilesProtocolEAPTTLSWithPAP
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]identity_services.RadiusServerProfilesProtocolEAPTTLSWithPAP, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.RadiusServerProfilesProtocolEAPTTLSWithPAP{}.AttrTypes(), &item)
		unpacked, d := unpackRadiusServerProfilesProtocolEAPTTLSWithPAPToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.RadiusServerProfilesProtocolEAPTTLSWithPAP", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for RadiusServerProfilesProtocolEAPTTLSWithPAP ---
func packRadiusServerProfilesProtocolEAPTTLSWithPAPListFromSdk(ctx context.Context, sdks []identity_services.RadiusServerProfilesProtocolEAPTTLSWithPAP) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.RadiusServerProfilesProtocolEAPTTLSWithPAP")
	diags := diag.Diagnostics{}
	var data []models.RadiusServerProfilesProtocolEAPTTLSWithPAP

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.RadiusServerProfilesProtocolEAPTTLSWithPAP
		obj, d := packRadiusServerProfilesProtocolEAPTTLSWithPAPFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.RadiusServerProfilesProtocolEAPTTLSWithPAP{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.RadiusServerProfilesProtocolEAPTTLSWithPAP", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.RadiusServerProfilesProtocolEAPTTLSWithPAP{}.AttrType(), data)
}

// --- Unpacker for RadiusServerProfilesProtocolPEAPMSCHAPv2 ---
func unpackRadiusServerProfilesProtocolPEAPMSCHAPv2ToSdk(ctx context.Context, obj types.Object) (*identity_services.RadiusServerProfilesProtocolPEAPMSCHAPv2, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.RadiusServerProfilesProtocolPEAPMSCHAPv2", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.RadiusServerProfilesProtocolPEAPMSCHAPv2
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk identity_services.RadiusServerProfilesProtocolPEAPMSCHAPv2
	var d diag.Diagnostics
	// Handling Primitives
	if !model.AllowPwdChange.IsNull() && !model.AllowPwdChange.IsUnknown() {
		sdk.AllowPwdChange = model.AllowPwdChange.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "AllowPwdChange", "value": *sdk.AllowPwdChange})
	}

	// Handling Primitives
	if !model.AnonOuterId.IsNull() && !model.AnonOuterId.IsUnknown() {
		sdk.AnonOuterId = model.AnonOuterId.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "AnonOuterId", "value": *sdk.AnonOuterId})
	}

	// Handling Primitives
	if !model.RadiusCertProfile.IsNull() && !model.RadiusCertProfile.IsUnknown() {
		sdk.RadiusCertProfile = model.RadiusCertProfile.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "RadiusCertProfile", "value": *sdk.RadiusCertProfile})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.RadiusServerProfilesProtocolPEAPMSCHAPv2", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for RadiusServerProfilesProtocolPEAPMSCHAPv2 ---
func packRadiusServerProfilesProtocolPEAPMSCHAPv2FromSdk(ctx context.Context, sdk identity_services.RadiusServerProfilesProtocolPEAPMSCHAPv2) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.RadiusServerProfilesProtocolPEAPMSCHAPv2", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.RadiusServerProfilesProtocolPEAPMSCHAPv2
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.AllowPwdChange != nil {
		model.AllowPwdChange = basetypes.NewBoolValue(*sdk.AllowPwdChange)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "AllowPwdChange", "value": *sdk.AllowPwdChange})
	} else {
		model.AllowPwdChange = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.AnonOuterId != nil {
		model.AnonOuterId = basetypes.NewBoolValue(*sdk.AnonOuterId)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "AnonOuterId", "value": *sdk.AnonOuterId})
	} else {
		model.AnonOuterId = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.RadiusCertProfile != nil {
		model.RadiusCertProfile = basetypes.NewStringValue(*sdk.RadiusCertProfile)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "RadiusCertProfile", "value": *sdk.RadiusCertProfile})
	} else {
		model.RadiusCertProfile = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.RadiusServerProfilesProtocolPEAPMSCHAPv2{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.RadiusServerProfilesProtocolPEAPMSCHAPv2", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for RadiusServerProfilesProtocolPEAPMSCHAPv2 ---
func unpackRadiusServerProfilesProtocolPEAPMSCHAPv2ListToSdk(ctx context.Context, list types.List) ([]identity_services.RadiusServerProfilesProtocolPEAPMSCHAPv2, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.RadiusServerProfilesProtocolPEAPMSCHAPv2")
	diags := diag.Diagnostics{}
	var data []models.RadiusServerProfilesProtocolPEAPMSCHAPv2
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]identity_services.RadiusServerProfilesProtocolPEAPMSCHAPv2, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.RadiusServerProfilesProtocolPEAPMSCHAPv2{}.AttrTypes(), &item)
		unpacked, d := unpackRadiusServerProfilesProtocolPEAPMSCHAPv2ToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.RadiusServerProfilesProtocolPEAPMSCHAPv2", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for RadiusServerProfilesProtocolPEAPMSCHAPv2 ---
func packRadiusServerProfilesProtocolPEAPMSCHAPv2ListFromSdk(ctx context.Context, sdks []identity_services.RadiusServerProfilesProtocolPEAPMSCHAPv2) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.RadiusServerProfilesProtocolPEAPMSCHAPv2")
	diags := diag.Diagnostics{}
	var data []models.RadiusServerProfilesProtocolPEAPMSCHAPv2

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.RadiusServerProfilesProtocolPEAPMSCHAPv2
		obj, d := packRadiusServerProfilesProtocolPEAPMSCHAPv2FromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.RadiusServerProfilesProtocolPEAPMSCHAPv2{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.RadiusServerProfilesProtocolPEAPMSCHAPv2", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.RadiusServerProfilesProtocolPEAPMSCHAPv2{}.AttrType(), data)
}

// --- Unpacker for RadiusServerProfilesServerInner ---
func unpackRadiusServerProfilesServerInnerToSdk(ctx context.Context, obj types.Object) (*identity_services.RadiusServerProfilesServerInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.RadiusServerProfilesServerInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.RadiusServerProfilesServerInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk identity_services.RadiusServerProfilesServerInner
	var d diag.Diagnostics
	// Handling Primitives
	if !model.IpAddress.IsNull() && !model.IpAddress.IsUnknown() {
		sdk.IpAddress = model.IpAddress.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "IpAddress", "value": *sdk.IpAddress})
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
	}

	// Handling Primitives
	if !model.Port.IsNull() && !model.Port.IsUnknown() {
		val := int32(model.Port.ValueInt64())
		sdk.Port = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Port", "value": *sdk.Port})
	}

	// Handling Primitives
	if !model.Secret.IsNull() && !model.Secret.IsUnknown() {
		sdk.Secret = model.Secret.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Secret", "value": *sdk.Secret})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.RadiusServerProfilesServerInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for RadiusServerProfilesServerInner ---
func packRadiusServerProfilesServerInnerFromSdk(ctx context.Context, sdk identity_services.RadiusServerProfilesServerInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.RadiusServerProfilesServerInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.RadiusServerProfilesServerInner
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.IpAddress != nil {
		model.IpAddress = basetypes.NewStringValue(*sdk.IpAddress)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "IpAddress", "value": *sdk.IpAddress})
	} else {
		model.IpAddress = basetypes.NewStringNull()
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
		model.Port = basetypes.NewInt64Value(int64(*sdk.Port))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Port", "value": *sdk.Port})
	} else {
		model.Port = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Secret != nil {
		model.Secret = basetypes.NewStringValue(*sdk.Secret)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Secret", "value": *sdk.Secret})
	} else {
		model.Secret = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.RadiusServerProfilesServerInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.RadiusServerProfilesServerInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for RadiusServerProfilesServerInner ---
func unpackRadiusServerProfilesServerInnerListToSdk(ctx context.Context, list types.List) ([]identity_services.RadiusServerProfilesServerInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.RadiusServerProfilesServerInner")
	diags := diag.Diagnostics{}
	var data []models.RadiusServerProfilesServerInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]identity_services.RadiusServerProfilesServerInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.RadiusServerProfilesServerInner{}.AttrTypes(), &item)
		unpacked, d := unpackRadiusServerProfilesServerInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.RadiusServerProfilesServerInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for RadiusServerProfilesServerInner ---
func packRadiusServerProfilesServerInnerListFromSdk(ctx context.Context, sdks []identity_services.RadiusServerProfilesServerInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.RadiusServerProfilesServerInner")
	diags := diag.Diagnostics{}
	var data []models.RadiusServerProfilesServerInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.RadiusServerProfilesServerInner
		obj, d := packRadiusServerProfilesServerInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.RadiusServerProfilesServerInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.RadiusServerProfilesServerInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.RadiusServerProfilesServerInner{}.AttrType(), data)
}
