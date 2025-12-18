package provider

import (
	"context"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/objects"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/objects"
)

// externalDynamicListsSensitiveValuePatcher is an in-memory struct to temporarily store plaintext
// and encrypted values for sensitive fields during the Create/Update/Read workflows.
type externalDynamicListsSensitiveValuePatcher struct {
	type_domain_auth_password_plaintext basetypes.StringValue
	type_domain_auth_password_encrypted basetypes.StringValue
	type_imei_auth_password_plaintext   basetypes.StringValue
	type_imei_auth_password_encrypted   basetypes.StringValue
	type_imsi_auth_password_plaintext   basetypes.StringValue
	type_imsi_auth_password_encrypted   basetypes.StringValue
	type_ip_auth_password_plaintext     basetypes.StringValue
	type_ip_auth_password_encrypted     basetypes.StringValue
	type_url_auth_password_plaintext    basetypes.StringValue
	type_url_auth_password_encrypted    basetypes.StringValue
}

// populatePatcherFromState populates the patcher struct from the resource's state.
func (p *externalDynamicListsSensitiveValuePatcher) populatePatcherFromState(ctx context.Context, state models.ExternalDynamicLists) diag.Diagnostics {
	var diags diag.Diagnostics
	if state.EncryptedValues.IsNull() || state.EncryptedValues.IsUnknown() {
		return diags
	}

	var ev map[string]string
	diags.Append(state.EncryptedValues.ElementsAs(ctx, &ev, false)...)
	if diags.HasError() {
		return diags
	}
	if val, ok := ev["type_domain_auth_password_plaintext"]; ok {
		p.type_domain_auth_password_plaintext = basetypes.NewStringValue(val)
	}
	if val, ok := ev["type_domain_auth_password_encrypted"]; ok {
		p.type_domain_auth_password_encrypted = basetypes.NewStringValue(val)
	}
	if val, ok := ev["type_imei_auth_password_plaintext"]; ok {
		p.type_imei_auth_password_plaintext = basetypes.NewStringValue(val)
	}
	if val, ok := ev["type_imei_auth_password_encrypted"]; ok {
		p.type_imei_auth_password_encrypted = basetypes.NewStringValue(val)
	}
	if val, ok := ev["type_imsi_auth_password_plaintext"]; ok {
		p.type_imsi_auth_password_plaintext = basetypes.NewStringValue(val)
	}
	if val, ok := ev["type_imsi_auth_password_encrypted"]; ok {
		p.type_imsi_auth_password_encrypted = basetypes.NewStringValue(val)
	}
	if val, ok := ev["type_ip_auth_password_plaintext"]; ok {
		p.type_ip_auth_password_plaintext = basetypes.NewStringValue(val)
	}
	if val, ok := ev["type_ip_auth_password_encrypted"]; ok {
		p.type_ip_auth_password_encrypted = basetypes.NewStringValue(val)
	}
	if val, ok := ev["type_url_auth_password_plaintext"]; ok {
		p.type_url_auth_password_plaintext = basetypes.NewStringValue(val)
	}
	if val, ok := ev["type_url_auth_password_encrypted"]; ok {
		p.type_url_auth_password_encrypted = basetypes.NewStringValue(val)
	}

	return diags
}

// populateEncryptedValuesMap returns a map of the patcher's values for saving to state.
func (p *externalDynamicListsSensitiveValuePatcher) populateEncryptedValuesMap() map[string]string {
	ev := make(map[string]string)
	if !p.type_domain_auth_password_plaintext.IsNull() {
		ev["type_domain_auth_password_plaintext"] = p.type_domain_auth_password_plaintext.ValueString()
	}
	if !p.type_domain_auth_password_encrypted.IsNull() {
		ev["type_domain_auth_password_encrypted"] = p.type_domain_auth_password_encrypted.ValueString()
	}
	if !p.type_imei_auth_password_plaintext.IsNull() {
		ev["type_imei_auth_password_plaintext"] = p.type_imei_auth_password_plaintext.ValueString()
	}
	if !p.type_imei_auth_password_encrypted.IsNull() {
		ev["type_imei_auth_password_encrypted"] = p.type_imei_auth_password_encrypted.ValueString()
	}
	if !p.type_imsi_auth_password_plaintext.IsNull() {
		ev["type_imsi_auth_password_plaintext"] = p.type_imsi_auth_password_plaintext.ValueString()
	}
	if !p.type_imsi_auth_password_encrypted.IsNull() {
		ev["type_imsi_auth_password_encrypted"] = p.type_imsi_auth_password_encrypted.ValueString()
	}
	if !p.type_ip_auth_password_plaintext.IsNull() {
		ev["type_ip_auth_password_plaintext"] = p.type_ip_auth_password_plaintext.ValueString()
	}
	if !p.type_ip_auth_password_encrypted.IsNull() {
		ev["type_ip_auth_password_encrypted"] = p.type_ip_auth_password_encrypted.ValueString()
	}
	if !p.type_url_auth_password_plaintext.IsNull() {
		ev["type_url_auth_password_plaintext"] = p.type_url_auth_password_plaintext.ValueString()
	}
	if !p.type_url_auth_password_encrypted.IsNull() {
		ev["type_url_auth_password_encrypted"] = p.type_url_auth_password_encrypted.ValueString()
	}
	return ev
}

// --- Unpacker for ExternalDynamicLists ---
func unpackExternalDynamicListsToSdk(ctx context.Context, obj types.Object) (*objects.ExternalDynamicLists, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ExternalDynamicLists", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ExternalDynamicLists
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.ExternalDynamicLists
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
		sdk.Id = model.Id.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Id", "value": *sdk.Id})
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	// Handling Primitives
	if !model.Snippet.IsNull() && !model.Snippet.IsUnknown() {
		sdk.Snippet = model.Snippet.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	}

	// Handling Objects
	if !model.Type.IsNull() && !model.Type.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Type")
		unpacked, d := unpackExternalDynamicListsTypeToSdk(ctx, model.Type)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Type"})
		}
		if unpacked != nil {
			sdk.Type = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ExternalDynamicLists", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ExternalDynamicLists ---
func packExternalDynamicListsFromSdk(ctx context.Context, sdk objects.ExternalDynamicLists) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ExternalDynamicLists", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ExternalDynamicLists
	var d diag.Diagnostics
	// MEGA FIX FOR MAP TYPE MISMATCH (NOT ALL MODELS MAY HAVE EncryptedValues)
	model.EncryptedValues = basetypes.NewMapNull(basetypes.StringType{})
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
	if sdk.Snippet != nil {
		model.Snippet = basetypes.NewStringValue(*sdk.Snippet)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	} else {
		model.Snippet = basetypes.NewStringNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Type != nil {
		tflog.Debug(ctx, "Packing nested object for field Type")
		packed, d := packExternalDynamicListsTypeFromSdk(ctx, *sdk.Type)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Type"})
		}
		model.Type = packed
	} else {
		model.Type = basetypes.NewObjectNull(models.ExternalDynamicListsType{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ExternalDynamicLists{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ExternalDynamicLists", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ExternalDynamicLists ---
func unpackExternalDynamicListsListToSdk(ctx context.Context, list types.List) ([]objects.ExternalDynamicLists, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ExternalDynamicLists")
	diags := diag.Diagnostics{}
	var data []models.ExternalDynamicLists
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.ExternalDynamicLists, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ExternalDynamicLists{}.AttrTypes(), &item)
		unpacked, d := unpackExternalDynamicListsToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ExternalDynamicLists", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ExternalDynamicLists ---
func packExternalDynamicListsListFromSdk(ctx context.Context, sdks []objects.ExternalDynamicLists) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ExternalDynamicLists")
	diags := diag.Diagnostics{}
	var data []models.ExternalDynamicLists

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ExternalDynamicLists
		obj, d := packExternalDynamicListsFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ExternalDynamicLists{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ExternalDynamicLists", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ExternalDynamicLists{}.AttrType(), data)
}

// --- Unpacker for ExternalDynamicListsType ---
func unpackExternalDynamicListsTypeToSdk(ctx context.Context, obj types.Object) (*objects.ExternalDynamicListsType, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ExternalDynamicListsType", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ExternalDynamicListsType
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.ExternalDynamicListsType
	var d diag.Diagnostics
	// Handling Objects
	if !model.Domain.IsNull() && !model.Domain.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Domain")
		unpacked, d := unpackExternalDynamicListsTypeDomainToSdk(ctx, model.Domain)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Domain"})
		}
		if unpacked != nil {
			sdk.Domain = unpacked
		}
	}

	// Handling Objects
	if !model.Imei.IsNull() && !model.Imei.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Imei")
		unpacked, d := unpackExternalDynamicListsTypeImeiToSdk(ctx, model.Imei)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Imei"})
		}
		if unpacked != nil {
			sdk.Imei = unpacked
		}
	}

	// Handling Objects
	if !model.Imsi.IsNull() && !model.Imsi.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Imsi")
		unpacked, d := unpackExternalDynamicListsTypeImsiToSdk(ctx, model.Imsi)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Imsi"})
		}
		if unpacked != nil {
			sdk.Imsi = unpacked
		}
	}

	// Handling Objects
	if !model.Ip.IsNull() && !model.Ip.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Ip")
		unpacked, d := unpackExternalDynamicListsTypeIpToSdk(ctx, model.Ip)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Ip"})
		}
		if unpacked != nil {
			sdk.Ip = unpacked
		}
	}

	// Handling Objects
	if !model.PredefinedIp.IsNull() && !model.PredefinedIp.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field PredefinedIp")
		unpacked, d := unpackExternalDynamicListsTypePredefinedIpToSdk(ctx, model.PredefinedIp)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "PredefinedIp"})
		}
		if unpacked != nil {
			sdk.PredefinedIp = unpacked
		}
	}

	// Handling Objects
	if !model.PredefinedUrl.IsNull() && !model.PredefinedUrl.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field PredefinedUrl")
		unpacked, d := unpackExternalDynamicListsTypePredefinedUrlToSdk(ctx, model.PredefinedUrl)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "PredefinedUrl"})
		}
		if unpacked != nil {
			sdk.PredefinedUrl = unpacked
		}
	}

	// Handling Objects
	if !model.Url.IsNull() && !model.Url.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Url")
		unpacked, d := unpackExternalDynamicListsTypeUrlToSdk(ctx, model.Url)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Url"})
		}
		if unpacked != nil {
			sdk.Url = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ExternalDynamicListsType", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ExternalDynamicListsType ---
func packExternalDynamicListsTypeFromSdk(ctx context.Context, sdk objects.ExternalDynamicListsType) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ExternalDynamicListsType", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ExternalDynamicListsType
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Domain != nil {
		tflog.Debug(ctx, "Packing nested object for field Domain")
		packed, d := packExternalDynamicListsTypeDomainFromSdk(ctx, *sdk.Domain)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Domain"})
		}
		model.Domain = packed
	} else {
		model.Domain = basetypes.NewObjectNull(models.ExternalDynamicListsTypeDomain{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Imei != nil {
		tflog.Debug(ctx, "Packing nested object for field Imei")
		packed, d := packExternalDynamicListsTypeImeiFromSdk(ctx, *sdk.Imei)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Imei"})
		}
		model.Imei = packed
	} else {
		model.Imei = basetypes.NewObjectNull(models.ExternalDynamicListsTypeImei{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Imsi != nil {
		tflog.Debug(ctx, "Packing nested object for field Imsi")
		packed, d := packExternalDynamicListsTypeImsiFromSdk(ctx, *sdk.Imsi)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Imsi"})
		}
		model.Imsi = packed
	} else {
		model.Imsi = basetypes.NewObjectNull(models.ExternalDynamicListsTypeImsi{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Ip != nil {
		tflog.Debug(ctx, "Packing nested object for field Ip")
		packed, d := packExternalDynamicListsTypeIpFromSdk(ctx, *sdk.Ip)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Ip"})
		}
		model.Ip = packed
	} else {
		model.Ip = basetypes.NewObjectNull(models.ExternalDynamicListsTypeIp{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.PredefinedIp != nil {
		tflog.Debug(ctx, "Packing nested object for field PredefinedIp")
		packed, d := packExternalDynamicListsTypePredefinedIpFromSdk(ctx, *sdk.PredefinedIp)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "PredefinedIp"})
		}
		model.PredefinedIp = packed
	} else {
		model.PredefinedIp = basetypes.NewObjectNull(models.ExternalDynamicListsTypePredefinedIp{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.PredefinedUrl != nil {
		tflog.Debug(ctx, "Packing nested object for field PredefinedUrl")
		packed, d := packExternalDynamicListsTypePredefinedUrlFromSdk(ctx, *sdk.PredefinedUrl)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "PredefinedUrl"})
		}
		model.PredefinedUrl = packed
	} else {
		model.PredefinedUrl = basetypes.NewObjectNull(models.ExternalDynamicListsTypePredefinedUrl{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Url != nil {
		tflog.Debug(ctx, "Packing nested object for field Url")
		packed, d := packExternalDynamicListsTypeUrlFromSdk(ctx, *sdk.Url)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Url"})
		}
		model.Url = packed
	} else {
		model.Url = basetypes.NewObjectNull(models.ExternalDynamicListsTypeUrl{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ExternalDynamicListsType{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ExternalDynamicListsType", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ExternalDynamicListsType ---
func unpackExternalDynamicListsTypeListToSdk(ctx context.Context, list types.List) ([]objects.ExternalDynamicListsType, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ExternalDynamicListsType")
	diags := diag.Diagnostics{}
	var data []models.ExternalDynamicListsType
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.ExternalDynamicListsType, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ExternalDynamicListsType{}.AttrTypes(), &item)
		unpacked, d := unpackExternalDynamicListsTypeToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ExternalDynamicListsType", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ExternalDynamicListsType ---
func packExternalDynamicListsTypeListFromSdk(ctx context.Context, sdks []objects.ExternalDynamicListsType) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ExternalDynamicListsType")
	diags := diag.Diagnostics{}
	var data []models.ExternalDynamicListsType

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ExternalDynamicListsType
		obj, d := packExternalDynamicListsTypeFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ExternalDynamicListsType{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ExternalDynamicListsType", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ExternalDynamicListsType{}.AttrType(), data)
}

// --- Unpacker for ExternalDynamicListsTypeDomain ---
func unpackExternalDynamicListsTypeDomainToSdk(ctx context.Context, obj types.Object) (*objects.ExternalDynamicListsTypeDomain, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ExternalDynamicListsTypeDomain", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ExternalDynamicListsTypeDomain
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.ExternalDynamicListsTypeDomain
	var d diag.Diagnostics
	// Handling Objects
	if !model.Auth.IsNull() && !model.Auth.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Auth")
		unpacked, d := unpackExternalDynamicListsTypeDomainAuthToSdk(ctx, model.Auth)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Auth"})
		}
		if unpacked != nil {
			sdk.Auth = unpacked
		}
	}

	// Handling Primitives
	if !model.CertificateProfile.IsNull() && !model.CertificateProfile.IsUnknown() {
		sdk.CertificateProfile = model.CertificateProfile.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "CertificateProfile", "value": *sdk.CertificateProfile})
	}

	// Handling Primitives
	if !model.Description.IsNull() && !model.Description.IsUnknown() {
		sdk.Description = model.Description.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Description", "value": *sdk.Description})
	}

	// Handling Lists
	if !model.ExceptionList.IsNull() && !model.ExceptionList.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field ExceptionList")
		diags.Append(model.ExceptionList.ElementsAs(ctx, &sdk.ExceptionList, false)...)
	}

	// Handling Primitives
	if !model.ExpandDomain.IsNull() && !model.ExpandDomain.IsUnknown() {
		sdk.ExpandDomain = model.ExpandDomain.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "ExpandDomain", "value": *sdk.ExpandDomain})
	}

	// Handling Objects
	if !model.Recurring.IsNull() && !model.Recurring.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Recurring")
		unpacked, d := unpackExternalDynamicListsTypeDomainRecurringToSdk(ctx, model.Recurring)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Recurring"})
		}
		if unpacked != nil {
			sdk.Recurring = *unpacked
		}
	}

	// Handling Primitives
	if !model.Url.IsNull() && !model.Url.IsUnknown() {
		sdk.Url = model.Url.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Url", "value": sdk.Url})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ExternalDynamicListsTypeDomain", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ExternalDynamicListsTypeDomain ---
func packExternalDynamicListsTypeDomainFromSdk(ctx context.Context, sdk objects.ExternalDynamicListsTypeDomain) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ExternalDynamicListsTypeDomain", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ExternalDynamicListsTypeDomain
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Auth != nil {
		tflog.Debug(ctx, "Packing nested object for field Auth")
		packed, d := packExternalDynamicListsTypeDomainAuthFromSdk(ctx, *sdk.Auth)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Auth"})
		}
		model.Auth = packed
	} else {
		model.Auth = basetypes.NewObjectNull(models.ExternalDynamicListsTypeDomainAuth{}.AttrTypes())
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
	if sdk.Description != nil {
		model.Description = basetypes.NewStringValue(*sdk.Description)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Description", "value": *sdk.Description})
	} else {
		model.Description = basetypes.NewStringNull()
	}
	// Handling Lists
	if sdk.ExceptionList != nil {
		tflog.Debug(ctx, "Packing list of primitives for field ExceptionList")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.ExceptionList, d = basetypes.NewListValueFrom(ctx, elemType, sdk.ExceptionList)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.ExceptionList = basetypes.NewListNull(elemType)
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.ExpandDomain != nil {
		model.ExpandDomain = basetypes.NewBoolValue(*sdk.ExpandDomain)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "ExpandDomain", "value": *sdk.ExpandDomain})
	} else {
		model.ExpandDomain = basetypes.NewBoolNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	// Logic for non-pointer / value-type nested objects
	if !reflect.ValueOf(sdk.Recurring).IsZero() {
		tflog.Debug(ctx, "Packing nested object for field Recurring")
		packed, d := packExternalDynamicListsTypeDomainRecurringFromSdk(ctx, sdk.Recurring)
		diags.Append(d...)
		model.Recurring = packed
	} else {
		model.Recurring = basetypes.NewObjectNull(models.ExternalDynamicListsTypeDomainRecurring{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	model.Url = basetypes.NewStringValue(sdk.Url)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Url", "value": sdk.Url})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ExternalDynamicListsTypeDomain{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ExternalDynamicListsTypeDomain", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ExternalDynamicListsTypeDomain ---
func unpackExternalDynamicListsTypeDomainListToSdk(ctx context.Context, list types.List) ([]objects.ExternalDynamicListsTypeDomain, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ExternalDynamicListsTypeDomain")
	diags := diag.Diagnostics{}
	var data []models.ExternalDynamicListsTypeDomain
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.ExternalDynamicListsTypeDomain, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ExternalDynamicListsTypeDomain{}.AttrTypes(), &item)
		unpacked, d := unpackExternalDynamicListsTypeDomainToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ExternalDynamicListsTypeDomain", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ExternalDynamicListsTypeDomain ---
func packExternalDynamicListsTypeDomainListFromSdk(ctx context.Context, sdks []objects.ExternalDynamicListsTypeDomain) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ExternalDynamicListsTypeDomain")
	diags := diag.Diagnostics{}
	var data []models.ExternalDynamicListsTypeDomain

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ExternalDynamicListsTypeDomain
		obj, d := packExternalDynamicListsTypeDomainFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ExternalDynamicListsTypeDomain{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ExternalDynamicListsTypeDomain", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ExternalDynamicListsTypeDomain{}.AttrType(), data)
}

// --- Unpacker for ExternalDynamicListsTypeDomainAuth ---
func unpackExternalDynamicListsTypeDomainAuthToSdk(ctx context.Context, obj types.Object) (*objects.ExternalDynamicListsTypeDomainAuth, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ExternalDynamicListsTypeDomainAuth", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ExternalDynamicListsTypeDomainAuth
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.ExternalDynamicListsTypeDomainAuth
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Password.IsNull() && !model.Password.IsUnknown() {
		sdk.Password = model.Password.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Password", "value": sdk.Password})
	}

	// Handling Primitives
	if !model.Username.IsNull() && !model.Username.IsUnknown() {
		sdk.Username = model.Username.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Username", "value": sdk.Username})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ExternalDynamicListsTypeDomainAuth", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ExternalDynamicListsTypeDomainAuth ---
func packExternalDynamicListsTypeDomainAuthFromSdk(ctx context.Context, sdk objects.ExternalDynamicListsTypeDomainAuth) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ExternalDynamicListsTypeDomainAuth", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ExternalDynamicListsTypeDomainAuth
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.Password = basetypes.NewStringValue(sdk.Password)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Password", "value": sdk.Password})
	// Handling Primitives
	// Standard primitive packing
	model.Username = basetypes.NewStringValue(sdk.Username)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Username", "value": sdk.Username})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ExternalDynamicListsTypeDomainAuth{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ExternalDynamicListsTypeDomainAuth", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ExternalDynamicListsTypeDomainAuth ---
func unpackExternalDynamicListsTypeDomainAuthListToSdk(ctx context.Context, list types.List) ([]objects.ExternalDynamicListsTypeDomainAuth, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ExternalDynamicListsTypeDomainAuth")
	diags := diag.Diagnostics{}
	var data []models.ExternalDynamicListsTypeDomainAuth
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.ExternalDynamicListsTypeDomainAuth, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ExternalDynamicListsTypeDomainAuth{}.AttrTypes(), &item)
		unpacked, d := unpackExternalDynamicListsTypeDomainAuthToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ExternalDynamicListsTypeDomainAuth", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ExternalDynamicListsTypeDomainAuth ---
func packExternalDynamicListsTypeDomainAuthListFromSdk(ctx context.Context, sdks []objects.ExternalDynamicListsTypeDomainAuth) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ExternalDynamicListsTypeDomainAuth")
	diags := diag.Diagnostics{}
	var data []models.ExternalDynamicListsTypeDomainAuth

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ExternalDynamicListsTypeDomainAuth
		obj, d := packExternalDynamicListsTypeDomainAuthFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ExternalDynamicListsTypeDomainAuth{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ExternalDynamicListsTypeDomainAuth", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ExternalDynamicListsTypeDomainAuth{}.AttrType(), data)
}

// --- Unpacker for ExternalDynamicListsTypeDomainRecurring ---
func unpackExternalDynamicListsTypeDomainRecurringToSdk(ctx context.Context, obj types.Object) (*objects.ExternalDynamicListsTypeDomainRecurring, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ExternalDynamicListsTypeDomainRecurring", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ExternalDynamicListsTypeDomainRecurring
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.ExternalDynamicListsTypeDomainRecurring
	var d diag.Diagnostics
	// Handling Objects
	if !model.Daily.IsNull() && !model.Daily.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Daily")
		unpacked, d := unpackExternalDynamicListsTypeDomainRecurringDailyToSdk(ctx, model.Daily)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Daily"})
		}
		if unpacked != nil {
			sdk.Daily = unpacked
		}
	}

	// Handling Typeless Objects
	if !model.FiveMinute.IsNull() && !model.FiveMinute.IsUnknown() {
		tflog.Debug(ctx, "Unpacking typeless object for field FiveMinute")
		sdk.FiveMinute = make(map[string]interface{})
	}

	// Handling Typeless Objects
	if !model.Hourly.IsNull() && !model.Hourly.IsUnknown() {
		tflog.Debug(ctx, "Unpacking typeless object for field Hourly")
		sdk.Hourly = make(map[string]interface{})
	}

	// Handling Objects
	if !model.Monthly.IsNull() && !model.Monthly.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Monthly")
		unpacked, d := unpackExternalDynamicListsTypeDomainRecurringMonthlyToSdk(ctx, model.Monthly)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Monthly"})
		}
		if unpacked != nil {
			sdk.Monthly = unpacked
		}
	}

	// Handling Objects
	if !model.Weekly.IsNull() && !model.Weekly.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Weekly")
		unpacked, d := unpackExternalDynamicListsTypeDomainRecurringWeeklyToSdk(ctx, model.Weekly)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Weekly"})
		}
		if unpacked != nil {
			sdk.Weekly = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ExternalDynamicListsTypeDomainRecurring", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ExternalDynamicListsTypeDomainRecurring ---
func packExternalDynamicListsTypeDomainRecurringFromSdk(ctx context.Context, sdk objects.ExternalDynamicListsTypeDomainRecurring) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ExternalDynamicListsTypeDomainRecurring", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ExternalDynamicListsTypeDomainRecurring
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Daily != nil {
		tflog.Debug(ctx, "Packing nested object for field Daily")
		packed, d := packExternalDynamicListsTypeDomainRecurringDailyFromSdk(ctx, *sdk.Daily)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Daily"})
		}
		model.Daily = packed
	} else {
		model.Daily = basetypes.NewObjectNull(models.ExternalDynamicListsTypeDomainRecurringDaily{}.AttrTypes())
	}
	// Handling Objects
	// This is a marker object (e.g. CHAP: {}). We just need to create an empty, non-null object.
	if sdk.FiveMinute != nil && !reflect.ValueOf(sdk.FiveMinute).IsNil() {
		tflog.Debug(ctx, "Packing typeless object for field FiveMinute")
		var d diag.Diagnostics
		// Create an empty object with no attributes, which signifies its presence.
		model.FiveMinute, d = basetypes.NewObjectValue(map[string]attr.Type{}, map[string]attr.Value{})
		diags.Append(d...)
	} else {
		// Since this field is part of a oneOf, being nil means it's not selected.
		// We make the object null with an empty attribute map.
		model.FiveMinute = basetypes.NewObjectNull(map[string]attr.Type{})
	}
	// Handling Objects
	// This is a marker object (e.g. CHAP: {}). We just need to create an empty, non-null object.
	if sdk.Hourly != nil && !reflect.ValueOf(sdk.Hourly).IsNil() {
		tflog.Debug(ctx, "Packing typeless object for field Hourly")
		var d diag.Diagnostics
		// Create an empty object with no attributes, which signifies its presence.
		model.Hourly, d = basetypes.NewObjectValue(map[string]attr.Type{}, map[string]attr.Value{})
		diags.Append(d...)
	} else {
		// Since this field is part of a oneOf, being nil means it's not selected.
		// We make the object null with an empty attribute map.
		model.Hourly = basetypes.NewObjectNull(map[string]attr.Type{})
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Monthly != nil {
		tflog.Debug(ctx, "Packing nested object for field Monthly")
		packed, d := packExternalDynamicListsTypeDomainRecurringMonthlyFromSdk(ctx, *sdk.Monthly)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Monthly"})
		}
		model.Monthly = packed
	} else {
		model.Monthly = basetypes.NewObjectNull(models.ExternalDynamicListsTypeDomainRecurringMonthly{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Weekly != nil {
		tflog.Debug(ctx, "Packing nested object for field Weekly")
		packed, d := packExternalDynamicListsTypeDomainRecurringWeeklyFromSdk(ctx, *sdk.Weekly)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Weekly"})
		}
		model.Weekly = packed
	} else {
		model.Weekly = basetypes.NewObjectNull(models.ExternalDynamicListsTypeDomainRecurringWeekly{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ExternalDynamicListsTypeDomainRecurring{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ExternalDynamicListsTypeDomainRecurring", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ExternalDynamicListsTypeDomainRecurring ---
func unpackExternalDynamicListsTypeDomainRecurringListToSdk(ctx context.Context, list types.List) ([]objects.ExternalDynamicListsTypeDomainRecurring, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ExternalDynamicListsTypeDomainRecurring")
	diags := diag.Diagnostics{}
	var data []models.ExternalDynamicListsTypeDomainRecurring
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.ExternalDynamicListsTypeDomainRecurring, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ExternalDynamicListsTypeDomainRecurring{}.AttrTypes(), &item)
		unpacked, d := unpackExternalDynamicListsTypeDomainRecurringToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ExternalDynamicListsTypeDomainRecurring", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ExternalDynamicListsTypeDomainRecurring ---
func packExternalDynamicListsTypeDomainRecurringListFromSdk(ctx context.Context, sdks []objects.ExternalDynamicListsTypeDomainRecurring) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ExternalDynamicListsTypeDomainRecurring")
	diags := diag.Diagnostics{}
	var data []models.ExternalDynamicListsTypeDomainRecurring

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ExternalDynamicListsTypeDomainRecurring
		obj, d := packExternalDynamicListsTypeDomainRecurringFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ExternalDynamicListsTypeDomainRecurring{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ExternalDynamicListsTypeDomainRecurring", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ExternalDynamicListsTypeDomainRecurring{}.AttrType(), data)
}

// --- Unpacker for ExternalDynamicListsTypeDomainRecurringDaily ---
func unpackExternalDynamicListsTypeDomainRecurringDailyToSdk(ctx context.Context, obj types.Object) (*objects.ExternalDynamicListsTypeDomainRecurringDaily, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ExternalDynamicListsTypeDomainRecurringDaily", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ExternalDynamicListsTypeDomainRecurringDaily
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.ExternalDynamicListsTypeDomainRecurringDaily
	var d diag.Diagnostics
	// Handling Primitives
	if !model.At.IsNull() && !model.At.IsUnknown() {
		sdk.At = model.At.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "At", "value": sdk.At})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ExternalDynamicListsTypeDomainRecurringDaily", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ExternalDynamicListsTypeDomainRecurringDaily ---
func packExternalDynamicListsTypeDomainRecurringDailyFromSdk(ctx context.Context, sdk objects.ExternalDynamicListsTypeDomainRecurringDaily) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ExternalDynamicListsTypeDomainRecurringDaily", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ExternalDynamicListsTypeDomainRecurringDaily
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.At = basetypes.NewStringValue(sdk.At)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "At", "value": sdk.At})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ExternalDynamicListsTypeDomainRecurringDaily{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ExternalDynamicListsTypeDomainRecurringDaily", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ExternalDynamicListsTypeDomainRecurringDaily ---
func unpackExternalDynamicListsTypeDomainRecurringDailyListToSdk(ctx context.Context, list types.List) ([]objects.ExternalDynamicListsTypeDomainRecurringDaily, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ExternalDynamicListsTypeDomainRecurringDaily")
	diags := diag.Diagnostics{}
	var data []models.ExternalDynamicListsTypeDomainRecurringDaily
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.ExternalDynamicListsTypeDomainRecurringDaily, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ExternalDynamicListsTypeDomainRecurringDaily{}.AttrTypes(), &item)
		unpacked, d := unpackExternalDynamicListsTypeDomainRecurringDailyToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ExternalDynamicListsTypeDomainRecurringDaily", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ExternalDynamicListsTypeDomainRecurringDaily ---
func packExternalDynamicListsTypeDomainRecurringDailyListFromSdk(ctx context.Context, sdks []objects.ExternalDynamicListsTypeDomainRecurringDaily) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ExternalDynamicListsTypeDomainRecurringDaily")
	diags := diag.Diagnostics{}
	var data []models.ExternalDynamicListsTypeDomainRecurringDaily

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ExternalDynamicListsTypeDomainRecurringDaily
		obj, d := packExternalDynamicListsTypeDomainRecurringDailyFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ExternalDynamicListsTypeDomainRecurringDaily{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ExternalDynamicListsTypeDomainRecurringDaily", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ExternalDynamicListsTypeDomainRecurringDaily{}.AttrType(), data)
}

// --- Unpacker for ExternalDynamicListsTypeDomainRecurringMonthly ---
func unpackExternalDynamicListsTypeDomainRecurringMonthlyToSdk(ctx context.Context, obj types.Object) (*objects.ExternalDynamicListsTypeDomainRecurringMonthly, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ExternalDynamicListsTypeDomainRecurringMonthly", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ExternalDynamicListsTypeDomainRecurringMonthly
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.ExternalDynamicListsTypeDomainRecurringMonthly
	var d diag.Diagnostics
	// Handling Primitives
	if !model.At.IsNull() && !model.At.IsUnknown() {
		sdk.At = model.At.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "At", "value": sdk.At})
	}

	// Handling Primitives
	if !model.DayOfMonth.IsNull() && !model.DayOfMonth.IsUnknown() {
		sdk.DayOfMonth = int32(model.DayOfMonth.ValueInt64())
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "DayOfMonth", "value": sdk.DayOfMonth})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ExternalDynamicListsTypeDomainRecurringMonthly", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ExternalDynamicListsTypeDomainRecurringMonthly ---
func packExternalDynamicListsTypeDomainRecurringMonthlyFromSdk(ctx context.Context, sdk objects.ExternalDynamicListsTypeDomainRecurringMonthly) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ExternalDynamicListsTypeDomainRecurringMonthly", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ExternalDynamicListsTypeDomainRecurringMonthly
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.At = basetypes.NewStringValue(sdk.At)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "At", "value": sdk.At})
	// Handling Primitives
	// Standard primitive packing
	model.DayOfMonth = basetypes.NewInt64Value(int64(sdk.DayOfMonth))
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "DayOfMonth", "value": sdk.DayOfMonth})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ExternalDynamicListsTypeDomainRecurringMonthly{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ExternalDynamicListsTypeDomainRecurringMonthly", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ExternalDynamicListsTypeDomainRecurringMonthly ---
func unpackExternalDynamicListsTypeDomainRecurringMonthlyListToSdk(ctx context.Context, list types.List) ([]objects.ExternalDynamicListsTypeDomainRecurringMonthly, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ExternalDynamicListsTypeDomainRecurringMonthly")
	diags := diag.Diagnostics{}
	var data []models.ExternalDynamicListsTypeDomainRecurringMonthly
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.ExternalDynamicListsTypeDomainRecurringMonthly, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ExternalDynamicListsTypeDomainRecurringMonthly{}.AttrTypes(), &item)
		unpacked, d := unpackExternalDynamicListsTypeDomainRecurringMonthlyToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ExternalDynamicListsTypeDomainRecurringMonthly", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ExternalDynamicListsTypeDomainRecurringMonthly ---
func packExternalDynamicListsTypeDomainRecurringMonthlyListFromSdk(ctx context.Context, sdks []objects.ExternalDynamicListsTypeDomainRecurringMonthly) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ExternalDynamicListsTypeDomainRecurringMonthly")
	diags := diag.Diagnostics{}
	var data []models.ExternalDynamicListsTypeDomainRecurringMonthly

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ExternalDynamicListsTypeDomainRecurringMonthly
		obj, d := packExternalDynamicListsTypeDomainRecurringMonthlyFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ExternalDynamicListsTypeDomainRecurringMonthly{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ExternalDynamicListsTypeDomainRecurringMonthly", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ExternalDynamicListsTypeDomainRecurringMonthly{}.AttrType(), data)
}

// --- Unpacker for ExternalDynamicListsTypeDomainRecurringWeekly ---
func unpackExternalDynamicListsTypeDomainRecurringWeeklyToSdk(ctx context.Context, obj types.Object) (*objects.ExternalDynamicListsTypeDomainRecurringWeekly, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ExternalDynamicListsTypeDomainRecurringWeekly", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ExternalDynamicListsTypeDomainRecurringWeekly
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.ExternalDynamicListsTypeDomainRecurringWeekly
	var d diag.Diagnostics
	// Handling Primitives
	if !model.At.IsNull() && !model.At.IsUnknown() {
		sdk.At = model.At.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "At", "value": sdk.At})
	}

	// Handling Primitives
	if !model.DayOfWeek.IsNull() && !model.DayOfWeek.IsUnknown() {
		sdk.DayOfWeek = model.DayOfWeek.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "DayOfWeek", "value": sdk.DayOfWeek})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ExternalDynamicListsTypeDomainRecurringWeekly", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ExternalDynamicListsTypeDomainRecurringWeekly ---
func packExternalDynamicListsTypeDomainRecurringWeeklyFromSdk(ctx context.Context, sdk objects.ExternalDynamicListsTypeDomainRecurringWeekly) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ExternalDynamicListsTypeDomainRecurringWeekly", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ExternalDynamicListsTypeDomainRecurringWeekly
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.At = basetypes.NewStringValue(sdk.At)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "At", "value": sdk.At})
	// Handling Primitives
	// Standard primitive packing
	model.DayOfWeek = basetypes.NewStringValue(sdk.DayOfWeek)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "DayOfWeek", "value": sdk.DayOfWeek})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ExternalDynamicListsTypeDomainRecurringWeekly{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ExternalDynamicListsTypeDomainRecurringWeekly", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ExternalDynamicListsTypeDomainRecurringWeekly ---
func unpackExternalDynamicListsTypeDomainRecurringWeeklyListToSdk(ctx context.Context, list types.List) ([]objects.ExternalDynamicListsTypeDomainRecurringWeekly, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ExternalDynamicListsTypeDomainRecurringWeekly")
	diags := diag.Diagnostics{}
	var data []models.ExternalDynamicListsTypeDomainRecurringWeekly
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.ExternalDynamicListsTypeDomainRecurringWeekly, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ExternalDynamicListsTypeDomainRecurringWeekly{}.AttrTypes(), &item)
		unpacked, d := unpackExternalDynamicListsTypeDomainRecurringWeeklyToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ExternalDynamicListsTypeDomainRecurringWeekly", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ExternalDynamicListsTypeDomainRecurringWeekly ---
func packExternalDynamicListsTypeDomainRecurringWeeklyListFromSdk(ctx context.Context, sdks []objects.ExternalDynamicListsTypeDomainRecurringWeekly) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ExternalDynamicListsTypeDomainRecurringWeekly")
	diags := diag.Diagnostics{}
	var data []models.ExternalDynamicListsTypeDomainRecurringWeekly

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ExternalDynamicListsTypeDomainRecurringWeekly
		obj, d := packExternalDynamicListsTypeDomainRecurringWeeklyFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ExternalDynamicListsTypeDomainRecurringWeekly{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ExternalDynamicListsTypeDomainRecurringWeekly", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ExternalDynamicListsTypeDomainRecurringWeekly{}.AttrType(), data)
}

// --- Unpacker for ExternalDynamicListsTypeImei ---
func unpackExternalDynamicListsTypeImeiToSdk(ctx context.Context, obj types.Object) (*objects.ExternalDynamicListsTypeImei, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ExternalDynamicListsTypeImei", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ExternalDynamicListsTypeImei
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.ExternalDynamicListsTypeImei
	var d diag.Diagnostics
	// Handling Objects
	if !model.Auth.IsNull() && !model.Auth.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Auth")
		unpacked, d := unpackExternalDynamicListsTypeImeiAuthToSdk(ctx, model.Auth)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Auth"})
		}
		if unpacked != nil {
			sdk.Auth = unpacked
		}
	}

	// Handling Primitives
	if !model.CertificateProfile.IsNull() && !model.CertificateProfile.IsUnknown() {
		sdk.CertificateProfile = model.CertificateProfile.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "CertificateProfile", "value": *sdk.CertificateProfile})
	}

	// Handling Primitives
	if !model.Description.IsNull() && !model.Description.IsUnknown() {
		sdk.Description = model.Description.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Description", "value": *sdk.Description})
	}

	// Handling Lists
	if !model.ExceptionList.IsNull() && !model.ExceptionList.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field ExceptionList")
		diags.Append(model.ExceptionList.ElementsAs(ctx, &sdk.ExceptionList, false)...)
	}

	// Handling Objects
	if !model.Recurring.IsNull() && !model.Recurring.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Recurring")
		unpacked, d := unpackExternalDynamicListsTypeImeiRecurringToSdk(ctx, model.Recurring)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Recurring"})
		}
		if unpacked != nil {
			sdk.Recurring = *unpacked
		}
	}

	// Handling Primitives
	if !model.Url.IsNull() && !model.Url.IsUnknown() {
		sdk.Url = model.Url.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Url", "value": sdk.Url})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ExternalDynamicListsTypeImei", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ExternalDynamicListsTypeImei ---
func packExternalDynamicListsTypeImeiFromSdk(ctx context.Context, sdk objects.ExternalDynamicListsTypeImei) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ExternalDynamicListsTypeImei", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ExternalDynamicListsTypeImei
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Auth != nil {
		tflog.Debug(ctx, "Packing nested object for field Auth")
		packed, d := packExternalDynamicListsTypeImeiAuthFromSdk(ctx, *sdk.Auth)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Auth"})
		}
		model.Auth = packed
	} else {
		model.Auth = basetypes.NewObjectNull(models.ExternalDynamicListsTypeImeiAuth{}.AttrTypes())
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
	if sdk.Description != nil {
		model.Description = basetypes.NewStringValue(*sdk.Description)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Description", "value": *sdk.Description})
	} else {
		model.Description = basetypes.NewStringNull()
	}
	// Handling Lists
	if sdk.ExceptionList != nil {
		tflog.Debug(ctx, "Packing list of primitives for field ExceptionList")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.ExceptionList, d = basetypes.NewListValueFrom(ctx, elemType, sdk.ExceptionList)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.ExceptionList = basetypes.NewListNull(elemType)
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	// Logic for non-pointer / value-type nested objects
	if !reflect.ValueOf(sdk.Recurring).IsZero() {
		tflog.Debug(ctx, "Packing nested object for field Recurring")
		packed, d := packExternalDynamicListsTypeImeiRecurringFromSdk(ctx, sdk.Recurring)
		diags.Append(d...)
		model.Recurring = packed
	} else {
		model.Recurring = basetypes.NewObjectNull(models.ExternalDynamicListsTypeImeiRecurring{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	model.Url = basetypes.NewStringValue(sdk.Url)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Url", "value": sdk.Url})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ExternalDynamicListsTypeImei{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ExternalDynamicListsTypeImei", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ExternalDynamicListsTypeImei ---
func unpackExternalDynamicListsTypeImeiListToSdk(ctx context.Context, list types.List) ([]objects.ExternalDynamicListsTypeImei, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ExternalDynamicListsTypeImei")
	diags := diag.Diagnostics{}
	var data []models.ExternalDynamicListsTypeImei
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.ExternalDynamicListsTypeImei, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ExternalDynamicListsTypeImei{}.AttrTypes(), &item)
		unpacked, d := unpackExternalDynamicListsTypeImeiToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ExternalDynamicListsTypeImei", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ExternalDynamicListsTypeImei ---
func packExternalDynamicListsTypeImeiListFromSdk(ctx context.Context, sdks []objects.ExternalDynamicListsTypeImei) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ExternalDynamicListsTypeImei")
	diags := diag.Diagnostics{}
	var data []models.ExternalDynamicListsTypeImei

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ExternalDynamicListsTypeImei
		obj, d := packExternalDynamicListsTypeImeiFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ExternalDynamicListsTypeImei{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ExternalDynamicListsTypeImei", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ExternalDynamicListsTypeImei{}.AttrType(), data)
}

// --- Unpacker for ExternalDynamicListsTypeImeiAuth ---
func unpackExternalDynamicListsTypeImeiAuthToSdk(ctx context.Context, obj types.Object) (*objects.ExternalDynamicListsTypeImeiAuth, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ExternalDynamicListsTypeImeiAuth", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ExternalDynamicListsTypeImeiAuth
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.ExternalDynamicListsTypeImeiAuth
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Password.IsNull() && !model.Password.IsUnknown() {
		sdk.Password = model.Password.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Password", "value": sdk.Password})
	}

	// Handling Primitives
	if !model.Username.IsNull() && !model.Username.IsUnknown() {
		sdk.Username = model.Username.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Username", "value": sdk.Username})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ExternalDynamicListsTypeImeiAuth", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ExternalDynamicListsTypeImeiAuth ---
func packExternalDynamicListsTypeImeiAuthFromSdk(ctx context.Context, sdk objects.ExternalDynamicListsTypeImeiAuth) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ExternalDynamicListsTypeImeiAuth", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ExternalDynamicListsTypeImeiAuth
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.Password = basetypes.NewStringValue(sdk.Password)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Password", "value": sdk.Password})
	// Handling Primitives
	// Standard primitive packing
	model.Username = basetypes.NewStringValue(sdk.Username)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Username", "value": sdk.Username})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ExternalDynamicListsTypeImeiAuth{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ExternalDynamicListsTypeImeiAuth", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ExternalDynamicListsTypeImeiAuth ---
func unpackExternalDynamicListsTypeImeiAuthListToSdk(ctx context.Context, list types.List) ([]objects.ExternalDynamicListsTypeImeiAuth, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ExternalDynamicListsTypeImeiAuth")
	diags := diag.Diagnostics{}
	var data []models.ExternalDynamicListsTypeImeiAuth
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.ExternalDynamicListsTypeImeiAuth, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ExternalDynamicListsTypeImeiAuth{}.AttrTypes(), &item)
		unpacked, d := unpackExternalDynamicListsTypeImeiAuthToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ExternalDynamicListsTypeImeiAuth", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ExternalDynamicListsTypeImeiAuth ---
func packExternalDynamicListsTypeImeiAuthListFromSdk(ctx context.Context, sdks []objects.ExternalDynamicListsTypeImeiAuth) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ExternalDynamicListsTypeImeiAuth")
	diags := diag.Diagnostics{}
	var data []models.ExternalDynamicListsTypeImeiAuth

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ExternalDynamicListsTypeImeiAuth
		obj, d := packExternalDynamicListsTypeImeiAuthFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ExternalDynamicListsTypeImeiAuth{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ExternalDynamicListsTypeImeiAuth", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ExternalDynamicListsTypeImeiAuth{}.AttrType(), data)
}

// --- Unpacker for ExternalDynamicListsTypeImeiRecurring ---
func unpackExternalDynamicListsTypeImeiRecurringToSdk(ctx context.Context, obj types.Object) (*objects.ExternalDynamicListsTypeImeiRecurring, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ExternalDynamicListsTypeImeiRecurring", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ExternalDynamicListsTypeImeiRecurring
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.ExternalDynamicListsTypeImeiRecurring
	var d diag.Diagnostics
	// Handling Objects
	if !model.Daily.IsNull() && !model.Daily.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Daily")
		unpacked, d := unpackExternalDynamicListsTypeImeiRecurringDailyToSdk(ctx, model.Daily)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Daily"})
		}
		if unpacked != nil {
			sdk.Daily = unpacked
		}
	}

	// Handling Typeless Objects
	if !model.FiveMinute.IsNull() && !model.FiveMinute.IsUnknown() {
		tflog.Debug(ctx, "Unpacking typeless object for field FiveMinute")
		sdk.FiveMinute = make(map[string]interface{})
	}

	// Handling Typeless Objects
	if !model.Hourly.IsNull() && !model.Hourly.IsUnknown() {
		tflog.Debug(ctx, "Unpacking typeless object for field Hourly")
		sdk.Hourly = make(map[string]interface{})
	}

	// Handling Objects
	if !model.Monthly.IsNull() && !model.Monthly.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Monthly")
		unpacked, d := unpackExternalDynamicListsTypeImeiRecurringMonthlyToSdk(ctx, model.Monthly)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Monthly"})
		}
		if unpacked != nil {
			sdk.Monthly = unpacked
		}
	}

	// Handling Objects
	if !model.Weekly.IsNull() && !model.Weekly.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Weekly")
		unpacked, d := unpackExternalDynamicListsTypeImeiRecurringWeeklyToSdk(ctx, model.Weekly)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Weekly"})
		}
		if unpacked != nil {
			sdk.Weekly = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ExternalDynamicListsTypeImeiRecurring", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ExternalDynamicListsTypeImeiRecurring ---
func packExternalDynamicListsTypeImeiRecurringFromSdk(ctx context.Context, sdk objects.ExternalDynamicListsTypeImeiRecurring) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ExternalDynamicListsTypeImeiRecurring", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ExternalDynamicListsTypeImeiRecurring
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Daily != nil {
		tflog.Debug(ctx, "Packing nested object for field Daily")
		packed, d := packExternalDynamicListsTypeImeiRecurringDailyFromSdk(ctx, *sdk.Daily)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Daily"})
		}
		model.Daily = packed
	} else {
		model.Daily = basetypes.NewObjectNull(models.ExternalDynamicListsTypeImeiRecurringDaily{}.AttrTypes())
	}
	// Handling Objects
	// This is a marker object (e.g. CHAP: {}). We just need to create an empty, non-null object.
	if sdk.FiveMinute != nil && !reflect.ValueOf(sdk.FiveMinute).IsNil() {
		tflog.Debug(ctx, "Packing typeless object for field FiveMinute")
		var d diag.Diagnostics
		// Create an empty object with no attributes, which signifies its presence.
		model.FiveMinute, d = basetypes.NewObjectValue(map[string]attr.Type{}, map[string]attr.Value{})
		diags.Append(d...)
	} else {
		// Since this field is part of a oneOf, being nil means it's not selected.
		// We make the object null with an empty attribute map.
		model.FiveMinute = basetypes.NewObjectNull(map[string]attr.Type{})
	}
	// Handling Objects
	// This is a marker object (e.g. CHAP: {}). We just need to create an empty, non-null object.
	if sdk.Hourly != nil && !reflect.ValueOf(sdk.Hourly).IsNil() {
		tflog.Debug(ctx, "Packing typeless object for field Hourly")
		var d diag.Diagnostics
		// Create an empty object with no attributes, which signifies its presence.
		model.Hourly, d = basetypes.NewObjectValue(map[string]attr.Type{}, map[string]attr.Value{})
		diags.Append(d...)
	} else {
		// Since this field is part of a oneOf, being nil means it's not selected.
		// We make the object null with an empty attribute map.
		model.Hourly = basetypes.NewObjectNull(map[string]attr.Type{})
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Monthly != nil {
		tflog.Debug(ctx, "Packing nested object for field Monthly")
		packed, d := packExternalDynamicListsTypeImeiRecurringMonthlyFromSdk(ctx, *sdk.Monthly)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Monthly"})
		}
		model.Monthly = packed
	} else {
		model.Monthly = basetypes.NewObjectNull(models.ExternalDynamicListsTypeImeiRecurringMonthly{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Weekly != nil {
		tflog.Debug(ctx, "Packing nested object for field Weekly")
		packed, d := packExternalDynamicListsTypeImeiRecurringWeeklyFromSdk(ctx, *sdk.Weekly)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Weekly"})
		}
		model.Weekly = packed
	} else {
		model.Weekly = basetypes.NewObjectNull(models.ExternalDynamicListsTypeImeiRecurringWeekly{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ExternalDynamicListsTypeImeiRecurring{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ExternalDynamicListsTypeImeiRecurring", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ExternalDynamicListsTypeImeiRecurring ---
func unpackExternalDynamicListsTypeImeiRecurringListToSdk(ctx context.Context, list types.List) ([]objects.ExternalDynamicListsTypeImeiRecurring, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ExternalDynamicListsTypeImeiRecurring")
	diags := diag.Diagnostics{}
	var data []models.ExternalDynamicListsTypeImeiRecurring
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.ExternalDynamicListsTypeImeiRecurring, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ExternalDynamicListsTypeImeiRecurring{}.AttrTypes(), &item)
		unpacked, d := unpackExternalDynamicListsTypeImeiRecurringToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ExternalDynamicListsTypeImeiRecurring", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ExternalDynamicListsTypeImeiRecurring ---
func packExternalDynamicListsTypeImeiRecurringListFromSdk(ctx context.Context, sdks []objects.ExternalDynamicListsTypeImeiRecurring) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ExternalDynamicListsTypeImeiRecurring")
	diags := diag.Diagnostics{}
	var data []models.ExternalDynamicListsTypeImeiRecurring

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ExternalDynamicListsTypeImeiRecurring
		obj, d := packExternalDynamicListsTypeImeiRecurringFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ExternalDynamicListsTypeImeiRecurring{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ExternalDynamicListsTypeImeiRecurring", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ExternalDynamicListsTypeImeiRecurring{}.AttrType(), data)
}

// --- Unpacker for ExternalDynamicListsTypeImeiRecurringDaily ---
func unpackExternalDynamicListsTypeImeiRecurringDailyToSdk(ctx context.Context, obj types.Object) (*objects.ExternalDynamicListsTypeImeiRecurringDaily, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ExternalDynamicListsTypeImeiRecurringDaily", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ExternalDynamicListsTypeImeiRecurringDaily
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.ExternalDynamicListsTypeImeiRecurringDaily
	var d diag.Diagnostics
	// Handling Primitives
	if !model.At.IsNull() && !model.At.IsUnknown() {
		sdk.At = model.At.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "At", "value": sdk.At})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ExternalDynamicListsTypeImeiRecurringDaily", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ExternalDynamicListsTypeImeiRecurringDaily ---
func packExternalDynamicListsTypeImeiRecurringDailyFromSdk(ctx context.Context, sdk objects.ExternalDynamicListsTypeImeiRecurringDaily) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ExternalDynamicListsTypeImeiRecurringDaily", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ExternalDynamicListsTypeImeiRecurringDaily
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.At = basetypes.NewStringValue(sdk.At)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "At", "value": sdk.At})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ExternalDynamicListsTypeImeiRecurringDaily{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ExternalDynamicListsTypeImeiRecurringDaily", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ExternalDynamicListsTypeImeiRecurringDaily ---
func unpackExternalDynamicListsTypeImeiRecurringDailyListToSdk(ctx context.Context, list types.List) ([]objects.ExternalDynamicListsTypeImeiRecurringDaily, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ExternalDynamicListsTypeImeiRecurringDaily")
	diags := diag.Diagnostics{}
	var data []models.ExternalDynamicListsTypeImeiRecurringDaily
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.ExternalDynamicListsTypeImeiRecurringDaily, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ExternalDynamicListsTypeImeiRecurringDaily{}.AttrTypes(), &item)
		unpacked, d := unpackExternalDynamicListsTypeImeiRecurringDailyToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ExternalDynamicListsTypeImeiRecurringDaily", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ExternalDynamicListsTypeImeiRecurringDaily ---
func packExternalDynamicListsTypeImeiRecurringDailyListFromSdk(ctx context.Context, sdks []objects.ExternalDynamicListsTypeImeiRecurringDaily) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ExternalDynamicListsTypeImeiRecurringDaily")
	diags := diag.Diagnostics{}
	var data []models.ExternalDynamicListsTypeImeiRecurringDaily

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ExternalDynamicListsTypeImeiRecurringDaily
		obj, d := packExternalDynamicListsTypeImeiRecurringDailyFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ExternalDynamicListsTypeImeiRecurringDaily{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ExternalDynamicListsTypeImeiRecurringDaily", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ExternalDynamicListsTypeImeiRecurringDaily{}.AttrType(), data)
}

// --- Unpacker for ExternalDynamicListsTypeImeiRecurringMonthly ---
func unpackExternalDynamicListsTypeImeiRecurringMonthlyToSdk(ctx context.Context, obj types.Object) (*objects.ExternalDynamicListsTypeImeiRecurringMonthly, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ExternalDynamicListsTypeImeiRecurringMonthly", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ExternalDynamicListsTypeImeiRecurringMonthly
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.ExternalDynamicListsTypeImeiRecurringMonthly
	var d diag.Diagnostics
	// Handling Primitives
	if !model.At.IsNull() && !model.At.IsUnknown() {
		sdk.At = model.At.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "At", "value": sdk.At})
	}

	// Handling Primitives
	if !model.DayOfMonth.IsNull() && !model.DayOfMonth.IsUnknown() {
		sdk.DayOfMonth = int32(model.DayOfMonth.ValueInt64())
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "DayOfMonth", "value": sdk.DayOfMonth})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ExternalDynamicListsTypeImeiRecurringMonthly", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ExternalDynamicListsTypeImeiRecurringMonthly ---
func packExternalDynamicListsTypeImeiRecurringMonthlyFromSdk(ctx context.Context, sdk objects.ExternalDynamicListsTypeImeiRecurringMonthly) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ExternalDynamicListsTypeImeiRecurringMonthly", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ExternalDynamicListsTypeImeiRecurringMonthly
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.At = basetypes.NewStringValue(sdk.At)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "At", "value": sdk.At})
	// Handling Primitives
	// Standard primitive packing
	model.DayOfMonth = basetypes.NewInt64Value(int64(sdk.DayOfMonth))
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "DayOfMonth", "value": sdk.DayOfMonth})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ExternalDynamicListsTypeImeiRecurringMonthly{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ExternalDynamicListsTypeImeiRecurringMonthly", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ExternalDynamicListsTypeImeiRecurringMonthly ---
func unpackExternalDynamicListsTypeImeiRecurringMonthlyListToSdk(ctx context.Context, list types.List) ([]objects.ExternalDynamicListsTypeImeiRecurringMonthly, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ExternalDynamicListsTypeImeiRecurringMonthly")
	diags := diag.Diagnostics{}
	var data []models.ExternalDynamicListsTypeImeiRecurringMonthly
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.ExternalDynamicListsTypeImeiRecurringMonthly, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ExternalDynamicListsTypeImeiRecurringMonthly{}.AttrTypes(), &item)
		unpacked, d := unpackExternalDynamicListsTypeImeiRecurringMonthlyToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ExternalDynamicListsTypeImeiRecurringMonthly", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ExternalDynamicListsTypeImeiRecurringMonthly ---
func packExternalDynamicListsTypeImeiRecurringMonthlyListFromSdk(ctx context.Context, sdks []objects.ExternalDynamicListsTypeImeiRecurringMonthly) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ExternalDynamicListsTypeImeiRecurringMonthly")
	diags := diag.Diagnostics{}
	var data []models.ExternalDynamicListsTypeImeiRecurringMonthly

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ExternalDynamicListsTypeImeiRecurringMonthly
		obj, d := packExternalDynamicListsTypeImeiRecurringMonthlyFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ExternalDynamicListsTypeImeiRecurringMonthly{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ExternalDynamicListsTypeImeiRecurringMonthly", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ExternalDynamicListsTypeImeiRecurringMonthly{}.AttrType(), data)
}

// --- Unpacker for ExternalDynamicListsTypeImeiRecurringWeekly ---
func unpackExternalDynamicListsTypeImeiRecurringWeeklyToSdk(ctx context.Context, obj types.Object) (*objects.ExternalDynamicListsTypeImeiRecurringWeekly, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ExternalDynamicListsTypeImeiRecurringWeekly", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ExternalDynamicListsTypeImeiRecurringWeekly
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.ExternalDynamicListsTypeImeiRecurringWeekly
	var d diag.Diagnostics
	// Handling Primitives
	if !model.At.IsNull() && !model.At.IsUnknown() {
		sdk.At = model.At.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "At", "value": sdk.At})
	}

	// Handling Primitives
	if !model.DayOfWeek.IsNull() && !model.DayOfWeek.IsUnknown() {
		sdk.DayOfWeek = model.DayOfWeek.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "DayOfWeek", "value": sdk.DayOfWeek})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ExternalDynamicListsTypeImeiRecurringWeekly", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ExternalDynamicListsTypeImeiRecurringWeekly ---
func packExternalDynamicListsTypeImeiRecurringWeeklyFromSdk(ctx context.Context, sdk objects.ExternalDynamicListsTypeImeiRecurringWeekly) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ExternalDynamicListsTypeImeiRecurringWeekly", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ExternalDynamicListsTypeImeiRecurringWeekly
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.At = basetypes.NewStringValue(sdk.At)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "At", "value": sdk.At})
	// Handling Primitives
	// Standard primitive packing
	model.DayOfWeek = basetypes.NewStringValue(sdk.DayOfWeek)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "DayOfWeek", "value": sdk.DayOfWeek})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ExternalDynamicListsTypeImeiRecurringWeekly{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ExternalDynamicListsTypeImeiRecurringWeekly", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ExternalDynamicListsTypeImeiRecurringWeekly ---
func unpackExternalDynamicListsTypeImeiRecurringWeeklyListToSdk(ctx context.Context, list types.List) ([]objects.ExternalDynamicListsTypeImeiRecurringWeekly, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ExternalDynamicListsTypeImeiRecurringWeekly")
	diags := diag.Diagnostics{}
	var data []models.ExternalDynamicListsTypeImeiRecurringWeekly
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.ExternalDynamicListsTypeImeiRecurringWeekly, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ExternalDynamicListsTypeImeiRecurringWeekly{}.AttrTypes(), &item)
		unpacked, d := unpackExternalDynamicListsTypeImeiRecurringWeeklyToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ExternalDynamicListsTypeImeiRecurringWeekly", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ExternalDynamicListsTypeImeiRecurringWeekly ---
func packExternalDynamicListsTypeImeiRecurringWeeklyListFromSdk(ctx context.Context, sdks []objects.ExternalDynamicListsTypeImeiRecurringWeekly) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ExternalDynamicListsTypeImeiRecurringWeekly")
	diags := diag.Diagnostics{}
	var data []models.ExternalDynamicListsTypeImeiRecurringWeekly

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ExternalDynamicListsTypeImeiRecurringWeekly
		obj, d := packExternalDynamicListsTypeImeiRecurringWeeklyFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ExternalDynamicListsTypeImeiRecurringWeekly{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ExternalDynamicListsTypeImeiRecurringWeekly", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ExternalDynamicListsTypeImeiRecurringWeekly{}.AttrType(), data)
}

// --- Unpacker for ExternalDynamicListsTypeImsi ---
func unpackExternalDynamicListsTypeImsiToSdk(ctx context.Context, obj types.Object) (*objects.ExternalDynamicListsTypeImsi, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ExternalDynamicListsTypeImsi", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ExternalDynamicListsTypeImsi
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.ExternalDynamicListsTypeImsi
	var d diag.Diagnostics
	// Handling Objects
	if !model.Auth.IsNull() && !model.Auth.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Auth")
		unpacked, d := unpackExternalDynamicListsTypeImsiAuthToSdk(ctx, model.Auth)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Auth"})
		}
		if unpacked != nil {
			sdk.Auth = unpacked
		}
	}

	// Handling Primitives
	if !model.CertificateProfile.IsNull() && !model.CertificateProfile.IsUnknown() {
		sdk.CertificateProfile = model.CertificateProfile.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "CertificateProfile", "value": *sdk.CertificateProfile})
	}

	// Handling Primitives
	if !model.Description.IsNull() && !model.Description.IsUnknown() {
		sdk.Description = model.Description.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Description", "value": *sdk.Description})
	}

	// Handling Lists
	if !model.ExceptionList.IsNull() && !model.ExceptionList.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field ExceptionList")
		diags.Append(model.ExceptionList.ElementsAs(ctx, &sdk.ExceptionList, false)...)
	}

	// Handling Objects
	if !model.Recurring.IsNull() && !model.Recurring.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Recurring")
		unpacked, d := unpackExternalDynamicListsTypeImsiRecurringToSdk(ctx, model.Recurring)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Recurring"})
		}
		if unpacked != nil {
			sdk.Recurring = *unpacked
		}
	}

	// Handling Primitives
	if !model.Url.IsNull() && !model.Url.IsUnknown() {
		sdk.Url = model.Url.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Url", "value": sdk.Url})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ExternalDynamicListsTypeImsi", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ExternalDynamicListsTypeImsi ---
func packExternalDynamicListsTypeImsiFromSdk(ctx context.Context, sdk objects.ExternalDynamicListsTypeImsi) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ExternalDynamicListsTypeImsi", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ExternalDynamicListsTypeImsi
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Auth != nil {
		tflog.Debug(ctx, "Packing nested object for field Auth")
		packed, d := packExternalDynamicListsTypeImsiAuthFromSdk(ctx, *sdk.Auth)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Auth"})
		}
		model.Auth = packed
	} else {
		model.Auth = basetypes.NewObjectNull(models.ExternalDynamicListsTypeImsiAuth{}.AttrTypes())
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
	if sdk.Description != nil {
		model.Description = basetypes.NewStringValue(*sdk.Description)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Description", "value": *sdk.Description})
	} else {
		model.Description = basetypes.NewStringNull()
	}
	// Handling Lists
	if sdk.ExceptionList != nil {
		tflog.Debug(ctx, "Packing list of primitives for field ExceptionList")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.ExceptionList, d = basetypes.NewListValueFrom(ctx, elemType, sdk.ExceptionList)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.ExceptionList = basetypes.NewListNull(elemType)
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	// Logic for non-pointer / value-type nested objects
	if !reflect.ValueOf(sdk.Recurring).IsZero() {
		tflog.Debug(ctx, "Packing nested object for field Recurring")
		packed, d := packExternalDynamicListsTypeImsiRecurringFromSdk(ctx, sdk.Recurring)
		diags.Append(d...)
		model.Recurring = packed
	} else {
		model.Recurring = basetypes.NewObjectNull(models.ExternalDynamicListsTypeImsiRecurring{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	model.Url = basetypes.NewStringValue(sdk.Url)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Url", "value": sdk.Url})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ExternalDynamicListsTypeImsi{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ExternalDynamicListsTypeImsi", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ExternalDynamicListsTypeImsi ---
func unpackExternalDynamicListsTypeImsiListToSdk(ctx context.Context, list types.List) ([]objects.ExternalDynamicListsTypeImsi, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ExternalDynamicListsTypeImsi")
	diags := diag.Diagnostics{}
	var data []models.ExternalDynamicListsTypeImsi
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.ExternalDynamicListsTypeImsi, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ExternalDynamicListsTypeImsi{}.AttrTypes(), &item)
		unpacked, d := unpackExternalDynamicListsTypeImsiToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ExternalDynamicListsTypeImsi", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ExternalDynamicListsTypeImsi ---
func packExternalDynamicListsTypeImsiListFromSdk(ctx context.Context, sdks []objects.ExternalDynamicListsTypeImsi) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ExternalDynamicListsTypeImsi")
	diags := diag.Diagnostics{}
	var data []models.ExternalDynamicListsTypeImsi

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ExternalDynamicListsTypeImsi
		obj, d := packExternalDynamicListsTypeImsiFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ExternalDynamicListsTypeImsi{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ExternalDynamicListsTypeImsi", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ExternalDynamicListsTypeImsi{}.AttrType(), data)
}

// --- Unpacker for ExternalDynamicListsTypeImsiAuth ---
func unpackExternalDynamicListsTypeImsiAuthToSdk(ctx context.Context, obj types.Object) (*objects.ExternalDynamicListsTypeImsiAuth, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ExternalDynamicListsTypeImsiAuth", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ExternalDynamicListsTypeImsiAuth
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.ExternalDynamicListsTypeImsiAuth
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Password.IsNull() && !model.Password.IsUnknown() {
		sdk.Password = model.Password.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Password", "value": sdk.Password})
	}

	// Handling Primitives
	if !model.Username.IsNull() && !model.Username.IsUnknown() {
		sdk.Username = model.Username.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Username", "value": sdk.Username})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ExternalDynamicListsTypeImsiAuth", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ExternalDynamicListsTypeImsiAuth ---
func packExternalDynamicListsTypeImsiAuthFromSdk(ctx context.Context, sdk objects.ExternalDynamicListsTypeImsiAuth) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ExternalDynamicListsTypeImsiAuth", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ExternalDynamicListsTypeImsiAuth
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.Password = basetypes.NewStringValue(sdk.Password)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Password", "value": sdk.Password})
	// Handling Primitives
	// Standard primitive packing
	model.Username = basetypes.NewStringValue(sdk.Username)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Username", "value": sdk.Username})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ExternalDynamicListsTypeImsiAuth{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ExternalDynamicListsTypeImsiAuth", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ExternalDynamicListsTypeImsiAuth ---
func unpackExternalDynamicListsTypeImsiAuthListToSdk(ctx context.Context, list types.List) ([]objects.ExternalDynamicListsTypeImsiAuth, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ExternalDynamicListsTypeImsiAuth")
	diags := diag.Diagnostics{}
	var data []models.ExternalDynamicListsTypeImsiAuth
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.ExternalDynamicListsTypeImsiAuth, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ExternalDynamicListsTypeImsiAuth{}.AttrTypes(), &item)
		unpacked, d := unpackExternalDynamicListsTypeImsiAuthToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ExternalDynamicListsTypeImsiAuth", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ExternalDynamicListsTypeImsiAuth ---
func packExternalDynamicListsTypeImsiAuthListFromSdk(ctx context.Context, sdks []objects.ExternalDynamicListsTypeImsiAuth) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ExternalDynamicListsTypeImsiAuth")
	diags := diag.Diagnostics{}
	var data []models.ExternalDynamicListsTypeImsiAuth

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ExternalDynamicListsTypeImsiAuth
		obj, d := packExternalDynamicListsTypeImsiAuthFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ExternalDynamicListsTypeImsiAuth{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ExternalDynamicListsTypeImsiAuth", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ExternalDynamicListsTypeImsiAuth{}.AttrType(), data)
}

// --- Unpacker for ExternalDynamicListsTypeImsiRecurring ---
func unpackExternalDynamicListsTypeImsiRecurringToSdk(ctx context.Context, obj types.Object) (*objects.ExternalDynamicListsTypeImsiRecurring, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ExternalDynamicListsTypeImsiRecurring", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ExternalDynamicListsTypeImsiRecurring
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.ExternalDynamicListsTypeImsiRecurring
	var d diag.Diagnostics
	// Handling Objects
	if !model.Daily.IsNull() && !model.Daily.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Daily")
		unpacked, d := unpackExternalDynamicListsTypeImsiRecurringDailyToSdk(ctx, model.Daily)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Daily"})
		}
		if unpacked != nil {
			sdk.Daily = unpacked
		}
	}

	// Handling Typeless Objects
	if !model.FiveMinute.IsNull() && !model.FiveMinute.IsUnknown() {
		tflog.Debug(ctx, "Unpacking typeless object for field FiveMinute")
		sdk.FiveMinute = make(map[string]interface{})
	}

	// Handling Typeless Objects
	if !model.Hourly.IsNull() && !model.Hourly.IsUnknown() {
		tflog.Debug(ctx, "Unpacking typeless object for field Hourly")
		sdk.Hourly = make(map[string]interface{})
	}

	// Handling Objects
	if !model.Monthly.IsNull() && !model.Monthly.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Monthly")
		unpacked, d := unpackExternalDynamicListsTypeImsiRecurringMonthlyToSdk(ctx, model.Monthly)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Monthly"})
		}
		if unpacked != nil {
			sdk.Monthly = unpacked
		}
	}

	// Handling Objects
	if !model.Weekly.IsNull() && !model.Weekly.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Weekly")
		unpacked, d := unpackExternalDynamicListsTypeImsiRecurringWeeklyToSdk(ctx, model.Weekly)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Weekly"})
		}
		if unpacked != nil {
			sdk.Weekly = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ExternalDynamicListsTypeImsiRecurring", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ExternalDynamicListsTypeImsiRecurring ---
func packExternalDynamicListsTypeImsiRecurringFromSdk(ctx context.Context, sdk objects.ExternalDynamicListsTypeImsiRecurring) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ExternalDynamicListsTypeImsiRecurring", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ExternalDynamicListsTypeImsiRecurring
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Daily != nil {
		tflog.Debug(ctx, "Packing nested object for field Daily")
		packed, d := packExternalDynamicListsTypeImsiRecurringDailyFromSdk(ctx, *sdk.Daily)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Daily"})
		}
		model.Daily = packed
	} else {
		model.Daily = basetypes.NewObjectNull(models.ExternalDynamicListsTypeImsiRecurringDaily{}.AttrTypes())
	}
	// Handling Objects
	// This is a marker object (e.g. CHAP: {}). We just need to create an empty, non-null object.
	if sdk.FiveMinute != nil && !reflect.ValueOf(sdk.FiveMinute).IsNil() {
		tflog.Debug(ctx, "Packing typeless object for field FiveMinute")
		var d diag.Diagnostics
		// Create an empty object with no attributes, which signifies its presence.
		model.FiveMinute, d = basetypes.NewObjectValue(map[string]attr.Type{}, map[string]attr.Value{})
		diags.Append(d...)
	} else {
		// Since this field is part of a oneOf, being nil means it's not selected.
		// We make the object null with an empty attribute map.
		model.FiveMinute = basetypes.NewObjectNull(map[string]attr.Type{})
	}
	// Handling Objects
	// This is a marker object (e.g. CHAP: {}). We just need to create an empty, non-null object.
	if sdk.Hourly != nil && !reflect.ValueOf(sdk.Hourly).IsNil() {
		tflog.Debug(ctx, "Packing typeless object for field Hourly")
		var d diag.Diagnostics
		// Create an empty object with no attributes, which signifies its presence.
		model.Hourly, d = basetypes.NewObjectValue(map[string]attr.Type{}, map[string]attr.Value{})
		diags.Append(d...)
	} else {
		// Since this field is part of a oneOf, being nil means it's not selected.
		// We make the object null with an empty attribute map.
		model.Hourly = basetypes.NewObjectNull(map[string]attr.Type{})
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Monthly != nil {
		tflog.Debug(ctx, "Packing nested object for field Monthly")
		packed, d := packExternalDynamicListsTypeImsiRecurringMonthlyFromSdk(ctx, *sdk.Monthly)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Monthly"})
		}
		model.Monthly = packed
	} else {
		model.Monthly = basetypes.NewObjectNull(models.ExternalDynamicListsTypeImsiRecurringMonthly{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Weekly != nil {
		tflog.Debug(ctx, "Packing nested object for field Weekly")
		packed, d := packExternalDynamicListsTypeImsiRecurringWeeklyFromSdk(ctx, *sdk.Weekly)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Weekly"})
		}
		model.Weekly = packed
	} else {
		model.Weekly = basetypes.NewObjectNull(models.ExternalDynamicListsTypeImsiRecurringWeekly{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ExternalDynamicListsTypeImsiRecurring{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ExternalDynamicListsTypeImsiRecurring", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ExternalDynamicListsTypeImsiRecurring ---
func unpackExternalDynamicListsTypeImsiRecurringListToSdk(ctx context.Context, list types.List) ([]objects.ExternalDynamicListsTypeImsiRecurring, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ExternalDynamicListsTypeImsiRecurring")
	diags := diag.Diagnostics{}
	var data []models.ExternalDynamicListsTypeImsiRecurring
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.ExternalDynamicListsTypeImsiRecurring, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ExternalDynamicListsTypeImsiRecurring{}.AttrTypes(), &item)
		unpacked, d := unpackExternalDynamicListsTypeImsiRecurringToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ExternalDynamicListsTypeImsiRecurring", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ExternalDynamicListsTypeImsiRecurring ---
func packExternalDynamicListsTypeImsiRecurringListFromSdk(ctx context.Context, sdks []objects.ExternalDynamicListsTypeImsiRecurring) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ExternalDynamicListsTypeImsiRecurring")
	diags := diag.Diagnostics{}
	var data []models.ExternalDynamicListsTypeImsiRecurring

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ExternalDynamicListsTypeImsiRecurring
		obj, d := packExternalDynamicListsTypeImsiRecurringFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ExternalDynamicListsTypeImsiRecurring{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ExternalDynamicListsTypeImsiRecurring", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ExternalDynamicListsTypeImsiRecurring{}.AttrType(), data)
}

// --- Unpacker for ExternalDynamicListsTypeImsiRecurringDaily ---
func unpackExternalDynamicListsTypeImsiRecurringDailyToSdk(ctx context.Context, obj types.Object) (*objects.ExternalDynamicListsTypeImsiRecurringDaily, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ExternalDynamicListsTypeImsiRecurringDaily", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ExternalDynamicListsTypeImsiRecurringDaily
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.ExternalDynamicListsTypeImsiRecurringDaily
	var d diag.Diagnostics
	// Handling Primitives
	if !model.At.IsNull() && !model.At.IsUnknown() {
		sdk.At = model.At.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "At", "value": sdk.At})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ExternalDynamicListsTypeImsiRecurringDaily", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ExternalDynamicListsTypeImsiRecurringDaily ---
func packExternalDynamicListsTypeImsiRecurringDailyFromSdk(ctx context.Context, sdk objects.ExternalDynamicListsTypeImsiRecurringDaily) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ExternalDynamicListsTypeImsiRecurringDaily", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ExternalDynamicListsTypeImsiRecurringDaily
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.At = basetypes.NewStringValue(sdk.At)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "At", "value": sdk.At})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ExternalDynamicListsTypeImsiRecurringDaily{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ExternalDynamicListsTypeImsiRecurringDaily", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ExternalDynamicListsTypeImsiRecurringDaily ---
func unpackExternalDynamicListsTypeImsiRecurringDailyListToSdk(ctx context.Context, list types.List) ([]objects.ExternalDynamicListsTypeImsiRecurringDaily, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ExternalDynamicListsTypeImsiRecurringDaily")
	diags := diag.Diagnostics{}
	var data []models.ExternalDynamicListsTypeImsiRecurringDaily
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.ExternalDynamicListsTypeImsiRecurringDaily, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ExternalDynamicListsTypeImsiRecurringDaily{}.AttrTypes(), &item)
		unpacked, d := unpackExternalDynamicListsTypeImsiRecurringDailyToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ExternalDynamicListsTypeImsiRecurringDaily", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ExternalDynamicListsTypeImsiRecurringDaily ---
func packExternalDynamicListsTypeImsiRecurringDailyListFromSdk(ctx context.Context, sdks []objects.ExternalDynamicListsTypeImsiRecurringDaily) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ExternalDynamicListsTypeImsiRecurringDaily")
	diags := diag.Diagnostics{}
	var data []models.ExternalDynamicListsTypeImsiRecurringDaily

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ExternalDynamicListsTypeImsiRecurringDaily
		obj, d := packExternalDynamicListsTypeImsiRecurringDailyFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ExternalDynamicListsTypeImsiRecurringDaily{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ExternalDynamicListsTypeImsiRecurringDaily", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ExternalDynamicListsTypeImsiRecurringDaily{}.AttrType(), data)
}

// --- Unpacker for ExternalDynamicListsTypeImsiRecurringMonthly ---
func unpackExternalDynamicListsTypeImsiRecurringMonthlyToSdk(ctx context.Context, obj types.Object) (*objects.ExternalDynamicListsTypeImsiRecurringMonthly, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ExternalDynamicListsTypeImsiRecurringMonthly", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ExternalDynamicListsTypeImsiRecurringMonthly
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.ExternalDynamicListsTypeImsiRecurringMonthly
	var d diag.Diagnostics
	// Handling Primitives
	if !model.At.IsNull() && !model.At.IsUnknown() {
		sdk.At = model.At.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "At", "value": sdk.At})
	}

	// Handling Primitives
	if !model.DayOfMonth.IsNull() && !model.DayOfMonth.IsUnknown() {
		sdk.DayOfMonth = int32(model.DayOfMonth.ValueInt64())
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "DayOfMonth", "value": sdk.DayOfMonth})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ExternalDynamicListsTypeImsiRecurringMonthly", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ExternalDynamicListsTypeImsiRecurringMonthly ---
func packExternalDynamicListsTypeImsiRecurringMonthlyFromSdk(ctx context.Context, sdk objects.ExternalDynamicListsTypeImsiRecurringMonthly) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ExternalDynamicListsTypeImsiRecurringMonthly", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ExternalDynamicListsTypeImsiRecurringMonthly
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.At = basetypes.NewStringValue(sdk.At)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "At", "value": sdk.At})
	// Handling Primitives
	// Standard primitive packing
	model.DayOfMonth = basetypes.NewInt64Value(int64(sdk.DayOfMonth))
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "DayOfMonth", "value": sdk.DayOfMonth})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ExternalDynamicListsTypeImsiRecurringMonthly{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ExternalDynamicListsTypeImsiRecurringMonthly", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ExternalDynamicListsTypeImsiRecurringMonthly ---
func unpackExternalDynamicListsTypeImsiRecurringMonthlyListToSdk(ctx context.Context, list types.List) ([]objects.ExternalDynamicListsTypeImsiRecurringMonthly, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ExternalDynamicListsTypeImsiRecurringMonthly")
	diags := diag.Diagnostics{}
	var data []models.ExternalDynamicListsTypeImsiRecurringMonthly
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.ExternalDynamicListsTypeImsiRecurringMonthly, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ExternalDynamicListsTypeImsiRecurringMonthly{}.AttrTypes(), &item)
		unpacked, d := unpackExternalDynamicListsTypeImsiRecurringMonthlyToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ExternalDynamicListsTypeImsiRecurringMonthly", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ExternalDynamicListsTypeImsiRecurringMonthly ---
func packExternalDynamicListsTypeImsiRecurringMonthlyListFromSdk(ctx context.Context, sdks []objects.ExternalDynamicListsTypeImsiRecurringMonthly) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ExternalDynamicListsTypeImsiRecurringMonthly")
	diags := diag.Diagnostics{}
	var data []models.ExternalDynamicListsTypeImsiRecurringMonthly

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ExternalDynamicListsTypeImsiRecurringMonthly
		obj, d := packExternalDynamicListsTypeImsiRecurringMonthlyFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ExternalDynamicListsTypeImsiRecurringMonthly{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ExternalDynamicListsTypeImsiRecurringMonthly", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ExternalDynamicListsTypeImsiRecurringMonthly{}.AttrType(), data)
}

// --- Unpacker for ExternalDynamicListsTypeImsiRecurringWeekly ---
func unpackExternalDynamicListsTypeImsiRecurringWeeklyToSdk(ctx context.Context, obj types.Object) (*objects.ExternalDynamicListsTypeImsiRecurringWeekly, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ExternalDynamicListsTypeImsiRecurringWeekly", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ExternalDynamicListsTypeImsiRecurringWeekly
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.ExternalDynamicListsTypeImsiRecurringWeekly
	var d diag.Diagnostics
	// Handling Primitives
	if !model.At.IsNull() && !model.At.IsUnknown() {
		sdk.At = model.At.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "At", "value": sdk.At})
	}

	// Handling Primitives
	if !model.DayOfWeek.IsNull() && !model.DayOfWeek.IsUnknown() {
		sdk.DayOfWeek = model.DayOfWeek.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "DayOfWeek", "value": sdk.DayOfWeek})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ExternalDynamicListsTypeImsiRecurringWeekly", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ExternalDynamicListsTypeImsiRecurringWeekly ---
func packExternalDynamicListsTypeImsiRecurringWeeklyFromSdk(ctx context.Context, sdk objects.ExternalDynamicListsTypeImsiRecurringWeekly) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ExternalDynamicListsTypeImsiRecurringWeekly", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ExternalDynamicListsTypeImsiRecurringWeekly
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.At = basetypes.NewStringValue(sdk.At)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "At", "value": sdk.At})
	// Handling Primitives
	// Standard primitive packing
	model.DayOfWeek = basetypes.NewStringValue(sdk.DayOfWeek)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "DayOfWeek", "value": sdk.DayOfWeek})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ExternalDynamicListsTypeImsiRecurringWeekly{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ExternalDynamicListsTypeImsiRecurringWeekly", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ExternalDynamicListsTypeImsiRecurringWeekly ---
func unpackExternalDynamicListsTypeImsiRecurringWeeklyListToSdk(ctx context.Context, list types.List) ([]objects.ExternalDynamicListsTypeImsiRecurringWeekly, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ExternalDynamicListsTypeImsiRecurringWeekly")
	diags := diag.Diagnostics{}
	var data []models.ExternalDynamicListsTypeImsiRecurringWeekly
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.ExternalDynamicListsTypeImsiRecurringWeekly, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ExternalDynamicListsTypeImsiRecurringWeekly{}.AttrTypes(), &item)
		unpacked, d := unpackExternalDynamicListsTypeImsiRecurringWeeklyToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ExternalDynamicListsTypeImsiRecurringWeekly", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ExternalDynamicListsTypeImsiRecurringWeekly ---
func packExternalDynamicListsTypeImsiRecurringWeeklyListFromSdk(ctx context.Context, sdks []objects.ExternalDynamicListsTypeImsiRecurringWeekly) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ExternalDynamicListsTypeImsiRecurringWeekly")
	diags := diag.Diagnostics{}
	var data []models.ExternalDynamicListsTypeImsiRecurringWeekly

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ExternalDynamicListsTypeImsiRecurringWeekly
		obj, d := packExternalDynamicListsTypeImsiRecurringWeeklyFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ExternalDynamicListsTypeImsiRecurringWeekly{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ExternalDynamicListsTypeImsiRecurringWeekly", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ExternalDynamicListsTypeImsiRecurringWeekly{}.AttrType(), data)
}

// --- Unpacker for ExternalDynamicListsTypeIp ---
func unpackExternalDynamicListsTypeIpToSdk(ctx context.Context, obj types.Object) (*objects.ExternalDynamicListsTypeIp, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ExternalDynamicListsTypeIp", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ExternalDynamicListsTypeIp
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.ExternalDynamicListsTypeIp
	var d diag.Diagnostics
	// Handling Objects
	if !model.Auth.IsNull() && !model.Auth.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Auth")
		unpacked, d := unpackExternalDynamicListsTypeIpAuthToSdk(ctx, model.Auth)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Auth"})
		}
		if unpacked != nil {
			sdk.Auth = unpacked
		}
	}

	// Handling Primitives
	if !model.CertificateProfile.IsNull() && !model.CertificateProfile.IsUnknown() {
		sdk.CertificateProfile = model.CertificateProfile.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "CertificateProfile", "value": *sdk.CertificateProfile})
	}

	// Handling Primitives
	if !model.Description.IsNull() && !model.Description.IsUnknown() {
		sdk.Description = model.Description.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Description", "value": *sdk.Description})
	}

	// Handling Lists
	if !model.ExceptionList.IsNull() && !model.ExceptionList.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field ExceptionList")
		diags.Append(model.ExceptionList.ElementsAs(ctx, &sdk.ExceptionList, false)...)
	}

	// Handling Objects
	if !model.Recurring.IsNull() && !model.Recurring.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Recurring")
		unpacked, d := unpackExternalDynamicListsTypeIpRecurringToSdk(ctx, model.Recurring)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Recurring"})
		}
		if unpacked != nil {
			sdk.Recurring = *unpacked
		}
	}

	// Handling Primitives
	if !model.Url.IsNull() && !model.Url.IsUnknown() {
		sdk.Url = model.Url.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Url", "value": sdk.Url})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ExternalDynamicListsTypeIp", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ExternalDynamicListsTypeIp ---
func packExternalDynamicListsTypeIpFromSdk(ctx context.Context, sdk objects.ExternalDynamicListsTypeIp) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ExternalDynamicListsTypeIp", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ExternalDynamicListsTypeIp
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Auth != nil {
		tflog.Debug(ctx, "Packing nested object for field Auth")
		packed, d := packExternalDynamicListsTypeIpAuthFromSdk(ctx, *sdk.Auth)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Auth"})
		}
		model.Auth = packed
	} else {
		model.Auth = basetypes.NewObjectNull(models.ExternalDynamicListsTypeIpAuth{}.AttrTypes())
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
	if sdk.Description != nil {
		model.Description = basetypes.NewStringValue(*sdk.Description)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Description", "value": *sdk.Description})
	} else {
		model.Description = basetypes.NewStringNull()
	}
	// Handling Lists
	if sdk.ExceptionList != nil {
		tflog.Debug(ctx, "Packing list of primitives for field ExceptionList")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.ExceptionList, d = basetypes.NewListValueFrom(ctx, elemType, sdk.ExceptionList)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.ExceptionList = basetypes.NewListNull(elemType)
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	// Logic for non-pointer / value-type nested objects
	if !reflect.ValueOf(sdk.Recurring).IsZero() {
		tflog.Debug(ctx, "Packing nested object for field Recurring")
		packed, d := packExternalDynamicListsTypeIpRecurringFromSdk(ctx, sdk.Recurring)
		diags.Append(d...)
		model.Recurring = packed
	} else {
		model.Recurring = basetypes.NewObjectNull(models.ExternalDynamicListsTypeIpRecurring{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	model.Url = basetypes.NewStringValue(sdk.Url)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Url", "value": sdk.Url})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ExternalDynamicListsTypeIp{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ExternalDynamicListsTypeIp", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ExternalDynamicListsTypeIp ---
func unpackExternalDynamicListsTypeIpListToSdk(ctx context.Context, list types.List) ([]objects.ExternalDynamicListsTypeIp, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ExternalDynamicListsTypeIp")
	diags := diag.Diagnostics{}
	var data []models.ExternalDynamicListsTypeIp
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.ExternalDynamicListsTypeIp, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ExternalDynamicListsTypeIp{}.AttrTypes(), &item)
		unpacked, d := unpackExternalDynamicListsTypeIpToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ExternalDynamicListsTypeIp", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ExternalDynamicListsTypeIp ---
func packExternalDynamicListsTypeIpListFromSdk(ctx context.Context, sdks []objects.ExternalDynamicListsTypeIp) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ExternalDynamicListsTypeIp")
	diags := diag.Diagnostics{}
	var data []models.ExternalDynamicListsTypeIp

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ExternalDynamicListsTypeIp
		obj, d := packExternalDynamicListsTypeIpFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ExternalDynamicListsTypeIp{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ExternalDynamicListsTypeIp", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ExternalDynamicListsTypeIp{}.AttrType(), data)
}

// --- Unpacker for ExternalDynamicListsTypeIpAuth ---
func unpackExternalDynamicListsTypeIpAuthToSdk(ctx context.Context, obj types.Object) (*objects.ExternalDynamicListsTypeIpAuth, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ExternalDynamicListsTypeIpAuth", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ExternalDynamicListsTypeIpAuth
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.ExternalDynamicListsTypeIpAuth
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Password.IsNull() && !model.Password.IsUnknown() {
		sdk.Password = model.Password.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Password", "value": sdk.Password})
	}

	// Handling Primitives
	if !model.Username.IsNull() && !model.Username.IsUnknown() {
		sdk.Username = model.Username.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Username", "value": sdk.Username})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ExternalDynamicListsTypeIpAuth", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ExternalDynamicListsTypeIpAuth ---
func packExternalDynamicListsTypeIpAuthFromSdk(ctx context.Context, sdk objects.ExternalDynamicListsTypeIpAuth) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ExternalDynamicListsTypeIpAuth", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ExternalDynamicListsTypeIpAuth
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.Password = basetypes.NewStringValue(sdk.Password)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Password", "value": sdk.Password})
	// Handling Primitives
	// Standard primitive packing
	model.Username = basetypes.NewStringValue(sdk.Username)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Username", "value": sdk.Username})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ExternalDynamicListsTypeIpAuth{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ExternalDynamicListsTypeIpAuth", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ExternalDynamicListsTypeIpAuth ---
func unpackExternalDynamicListsTypeIpAuthListToSdk(ctx context.Context, list types.List) ([]objects.ExternalDynamicListsTypeIpAuth, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ExternalDynamicListsTypeIpAuth")
	diags := diag.Diagnostics{}
	var data []models.ExternalDynamicListsTypeIpAuth
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.ExternalDynamicListsTypeIpAuth, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ExternalDynamicListsTypeIpAuth{}.AttrTypes(), &item)
		unpacked, d := unpackExternalDynamicListsTypeIpAuthToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ExternalDynamicListsTypeIpAuth", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ExternalDynamicListsTypeIpAuth ---
func packExternalDynamicListsTypeIpAuthListFromSdk(ctx context.Context, sdks []objects.ExternalDynamicListsTypeIpAuth) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ExternalDynamicListsTypeIpAuth")
	diags := diag.Diagnostics{}
	var data []models.ExternalDynamicListsTypeIpAuth

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ExternalDynamicListsTypeIpAuth
		obj, d := packExternalDynamicListsTypeIpAuthFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ExternalDynamicListsTypeIpAuth{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ExternalDynamicListsTypeIpAuth", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ExternalDynamicListsTypeIpAuth{}.AttrType(), data)
}

// --- Unpacker for ExternalDynamicListsTypeIpRecurring ---
func unpackExternalDynamicListsTypeIpRecurringToSdk(ctx context.Context, obj types.Object) (*objects.ExternalDynamicListsTypeIpRecurring, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ExternalDynamicListsTypeIpRecurring", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ExternalDynamicListsTypeIpRecurring
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.ExternalDynamicListsTypeIpRecurring
	var d diag.Diagnostics
	// Handling Objects
	if !model.Daily.IsNull() && !model.Daily.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Daily")
		unpacked, d := unpackExternalDynamicListsTypeIpRecurringDailyToSdk(ctx, model.Daily)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Daily"})
		}
		if unpacked != nil {
			sdk.Daily = unpacked
		}
	}

	// Handling Typeless Objects
	if !model.FiveMinute.IsNull() && !model.FiveMinute.IsUnknown() {
		tflog.Debug(ctx, "Unpacking typeless object for field FiveMinute")
		sdk.FiveMinute = make(map[string]interface{})
	}

	// Handling Typeless Objects
	if !model.Hourly.IsNull() && !model.Hourly.IsUnknown() {
		tflog.Debug(ctx, "Unpacking typeless object for field Hourly")
		sdk.Hourly = make(map[string]interface{})
	}

	// Handling Objects
	if !model.Monthly.IsNull() && !model.Monthly.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Monthly")
		unpacked, d := unpackExternalDynamicListsTypeIpRecurringMonthlyToSdk(ctx, model.Monthly)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Monthly"})
		}
		if unpacked != nil {
			sdk.Monthly = unpacked
		}
	}

	// Handling Objects
	if !model.Weekly.IsNull() && !model.Weekly.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Weekly")
		unpacked, d := unpackExternalDynamicListsTypeIpRecurringWeeklyToSdk(ctx, model.Weekly)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Weekly"})
		}
		if unpacked != nil {
			sdk.Weekly = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ExternalDynamicListsTypeIpRecurring", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ExternalDynamicListsTypeIpRecurring ---
func packExternalDynamicListsTypeIpRecurringFromSdk(ctx context.Context, sdk objects.ExternalDynamicListsTypeIpRecurring) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ExternalDynamicListsTypeIpRecurring", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ExternalDynamicListsTypeIpRecurring
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Daily != nil {
		tflog.Debug(ctx, "Packing nested object for field Daily")
		packed, d := packExternalDynamicListsTypeIpRecurringDailyFromSdk(ctx, *sdk.Daily)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Daily"})
		}
		model.Daily = packed
	} else {
		model.Daily = basetypes.NewObjectNull(models.ExternalDynamicListsTypeIpRecurringDaily{}.AttrTypes())
	}
	// Handling Objects
	// This is a marker object (e.g. CHAP: {}). We just need to create an empty, non-null object.
	if sdk.FiveMinute != nil && !reflect.ValueOf(sdk.FiveMinute).IsNil() {
		tflog.Debug(ctx, "Packing typeless object for field FiveMinute")
		var d diag.Diagnostics
		// Create an empty object with no attributes, which signifies its presence.
		model.FiveMinute, d = basetypes.NewObjectValue(map[string]attr.Type{}, map[string]attr.Value{})
		diags.Append(d...)
	} else {
		// Since this field is part of a oneOf, being nil means it's not selected.
		// We make the object null with an empty attribute map.
		model.FiveMinute = basetypes.NewObjectNull(map[string]attr.Type{})
	}
	// Handling Objects
	// This is a marker object (e.g. CHAP: {}). We just need to create an empty, non-null object.
	if sdk.Hourly != nil && !reflect.ValueOf(sdk.Hourly).IsNil() {
		tflog.Debug(ctx, "Packing typeless object for field Hourly")
		var d diag.Diagnostics
		// Create an empty object with no attributes, which signifies its presence.
		model.Hourly, d = basetypes.NewObjectValue(map[string]attr.Type{}, map[string]attr.Value{})
		diags.Append(d...)
	} else {
		// Since this field is part of a oneOf, being nil means it's not selected.
		// We make the object null with an empty attribute map.
		model.Hourly = basetypes.NewObjectNull(map[string]attr.Type{})
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Monthly != nil {
		tflog.Debug(ctx, "Packing nested object for field Monthly")
		packed, d := packExternalDynamicListsTypeIpRecurringMonthlyFromSdk(ctx, *sdk.Monthly)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Monthly"})
		}
		model.Monthly = packed
	} else {
		model.Monthly = basetypes.NewObjectNull(models.ExternalDynamicListsTypeIpRecurringMonthly{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Weekly != nil {
		tflog.Debug(ctx, "Packing nested object for field Weekly")
		packed, d := packExternalDynamicListsTypeIpRecurringWeeklyFromSdk(ctx, *sdk.Weekly)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Weekly"})
		}
		model.Weekly = packed
	} else {
		model.Weekly = basetypes.NewObjectNull(models.ExternalDynamicListsTypeIpRecurringWeekly{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ExternalDynamicListsTypeIpRecurring{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ExternalDynamicListsTypeIpRecurring", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ExternalDynamicListsTypeIpRecurring ---
func unpackExternalDynamicListsTypeIpRecurringListToSdk(ctx context.Context, list types.List) ([]objects.ExternalDynamicListsTypeIpRecurring, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ExternalDynamicListsTypeIpRecurring")
	diags := diag.Diagnostics{}
	var data []models.ExternalDynamicListsTypeIpRecurring
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.ExternalDynamicListsTypeIpRecurring, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ExternalDynamicListsTypeIpRecurring{}.AttrTypes(), &item)
		unpacked, d := unpackExternalDynamicListsTypeIpRecurringToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ExternalDynamicListsTypeIpRecurring", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ExternalDynamicListsTypeIpRecurring ---
func packExternalDynamicListsTypeIpRecurringListFromSdk(ctx context.Context, sdks []objects.ExternalDynamicListsTypeIpRecurring) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ExternalDynamicListsTypeIpRecurring")
	diags := diag.Diagnostics{}
	var data []models.ExternalDynamicListsTypeIpRecurring

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ExternalDynamicListsTypeIpRecurring
		obj, d := packExternalDynamicListsTypeIpRecurringFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ExternalDynamicListsTypeIpRecurring{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ExternalDynamicListsTypeIpRecurring", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ExternalDynamicListsTypeIpRecurring{}.AttrType(), data)
}

// --- Unpacker for ExternalDynamicListsTypeIpRecurringDaily ---
func unpackExternalDynamicListsTypeIpRecurringDailyToSdk(ctx context.Context, obj types.Object) (*objects.ExternalDynamicListsTypeIpRecurringDaily, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ExternalDynamicListsTypeIpRecurringDaily", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ExternalDynamicListsTypeIpRecurringDaily
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.ExternalDynamicListsTypeIpRecurringDaily
	var d diag.Diagnostics
	// Handling Primitives
	if !model.At.IsNull() && !model.At.IsUnknown() {
		sdk.At = model.At.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "At", "value": sdk.At})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ExternalDynamicListsTypeIpRecurringDaily", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ExternalDynamicListsTypeIpRecurringDaily ---
func packExternalDynamicListsTypeIpRecurringDailyFromSdk(ctx context.Context, sdk objects.ExternalDynamicListsTypeIpRecurringDaily) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ExternalDynamicListsTypeIpRecurringDaily", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ExternalDynamicListsTypeIpRecurringDaily
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.At = basetypes.NewStringValue(sdk.At)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "At", "value": sdk.At})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ExternalDynamicListsTypeIpRecurringDaily{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ExternalDynamicListsTypeIpRecurringDaily", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ExternalDynamicListsTypeIpRecurringDaily ---
func unpackExternalDynamicListsTypeIpRecurringDailyListToSdk(ctx context.Context, list types.List) ([]objects.ExternalDynamicListsTypeIpRecurringDaily, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ExternalDynamicListsTypeIpRecurringDaily")
	diags := diag.Diagnostics{}
	var data []models.ExternalDynamicListsTypeIpRecurringDaily
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.ExternalDynamicListsTypeIpRecurringDaily, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ExternalDynamicListsTypeIpRecurringDaily{}.AttrTypes(), &item)
		unpacked, d := unpackExternalDynamicListsTypeIpRecurringDailyToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ExternalDynamicListsTypeIpRecurringDaily", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ExternalDynamicListsTypeIpRecurringDaily ---
func packExternalDynamicListsTypeIpRecurringDailyListFromSdk(ctx context.Context, sdks []objects.ExternalDynamicListsTypeIpRecurringDaily) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ExternalDynamicListsTypeIpRecurringDaily")
	diags := diag.Diagnostics{}
	var data []models.ExternalDynamicListsTypeIpRecurringDaily

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ExternalDynamicListsTypeIpRecurringDaily
		obj, d := packExternalDynamicListsTypeIpRecurringDailyFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ExternalDynamicListsTypeIpRecurringDaily{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ExternalDynamicListsTypeIpRecurringDaily", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ExternalDynamicListsTypeIpRecurringDaily{}.AttrType(), data)
}

// --- Unpacker for ExternalDynamicListsTypeIpRecurringMonthly ---
func unpackExternalDynamicListsTypeIpRecurringMonthlyToSdk(ctx context.Context, obj types.Object) (*objects.ExternalDynamicListsTypeIpRecurringMonthly, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ExternalDynamicListsTypeIpRecurringMonthly", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ExternalDynamicListsTypeIpRecurringMonthly
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.ExternalDynamicListsTypeIpRecurringMonthly
	var d diag.Diagnostics
	// Handling Primitives
	if !model.At.IsNull() && !model.At.IsUnknown() {
		sdk.At = model.At.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "At", "value": sdk.At})
	}

	// Handling Primitives
	if !model.DayOfMonth.IsNull() && !model.DayOfMonth.IsUnknown() {
		sdk.DayOfMonth = int32(model.DayOfMonth.ValueInt64())
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "DayOfMonth", "value": sdk.DayOfMonth})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ExternalDynamicListsTypeIpRecurringMonthly", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ExternalDynamicListsTypeIpRecurringMonthly ---
func packExternalDynamicListsTypeIpRecurringMonthlyFromSdk(ctx context.Context, sdk objects.ExternalDynamicListsTypeIpRecurringMonthly) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ExternalDynamicListsTypeIpRecurringMonthly", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ExternalDynamicListsTypeIpRecurringMonthly
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.At = basetypes.NewStringValue(sdk.At)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "At", "value": sdk.At})
	// Handling Primitives
	// Standard primitive packing
	model.DayOfMonth = basetypes.NewInt64Value(int64(sdk.DayOfMonth))
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "DayOfMonth", "value": sdk.DayOfMonth})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ExternalDynamicListsTypeIpRecurringMonthly{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ExternalDynamicListsTypeIpRecurringMonthly", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ExternalDynamicListsTypeIpRecurringMonthly ---
func unpackExternalDynamicListsTypeIpRecurringMonthlyListToSdk(ctx context.Context, list types.List) ([]objects.ExternalDynamicListsTypeIpRecurringMonthly, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ExternalDynamicListsTypeIpRecurringMonthly")
	diags := diag.Diagnostics{}
	var data []models.ExternalDynamicListsTypeIpRecurringMonthly
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.ExternalDynamicListsTypeIpRecurringMonthly, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ExternalDynamicListsTypeIpRecurringMonthly{}.AttrTypes(), &item)
		unpacked, d := unpackExternalDynamicListsTypeIpRecurringMonthlyToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ExternalDynamicListsTypeIpRecurringMonthly", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ExternalDynamicListsTypeIpRecurringMonthly ---
func packExternalDynamicListsTypeIpRecurringMonthlyListFromSdk(ctx context.Context, sdks []objects.ExternalDynamicListsTypeIpRecurringMonthly) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ExternalDynamicListsTypeIpRecurringMonthly")
	diags := diag.Diagnostics{}
	var data []models.ExternalDynamicListsTypeIpRecurringMonthly

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ExternalDynamicListsTypeIpRecurringMonthly
		obj, d := packExternalDynamicListsTypeIpRecurringMonthlyFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ExternalDynamicListsTypeIpRecurringMonthly{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ExternalDynamicListsTypeIpRecurringMonthly", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ExternalDynamicListsTypeIpRecurringMonthly{}.AttrType(), data)
}

// --- Unpacker for ExternalDynamicListsTypeIpRecurringWeekly ---
func unpackExternalDynamicListsTypeIpRecurringWeeklyToSdk(ctx context.Context, obj types.Object) (*objects.ExternalDynamicListsTypeIpRecurringWeekly, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ExternalDynamicListsTypeIpRecurringWeekly", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ExternalDynamicListsTypeIpRecurringWeekly
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.ExternalDynamicListsTypeIpRecurringWeekly
	var d diag.Diagnostics
	// Handling Primitives
	if !model.At.IsNull() && !model.At.IsUnknown() {
		sdk.At = model.At.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "At", "value": sdk.At})
	}

	// Handling Primitives
	if !model.DayOfWeek.IsNull() && !model.DayOfWeek.IsUnknown() {
		sdk.DayOfWeek = model.DayOfWeek.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "DayOfWeek", "value": sdk.DayOfWeek})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ExternalDynamicListsTypeIpRecurringWeekly", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ExternalDynamicListsTypeIpRecurringWeekly ---
func packExternalDynamicListsTypeIpRecurringWeeklyFromSdk(ctx context.Context, sdk objects.ExternalDynamicListsTypeIpRecurringWeekly) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ExternalDynamicListsTypeIpRecurringWeekly", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ExternalDynamicListsTypeIpRecurringWeekly
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.At = basetypes.NewStringValue(sdk.At)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "At", "value": sdk.At})
	// Handling Primitives
	// Standard primitive packing
	model.DayOfWeek = basetypes.NewStringValue(sdk.DayOfWeek)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "DayOfWeek", "value": sdk.DayOfWeek})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ExternalDynamicListsTypeIpRecurringWeekly{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ExternalDynamicListsTypeIpRecurringWeekly", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ExternalDynamicListsTypeIpRecurringWeekly ---
func unpackExternalDynamicListsTypeIpRecurringWeeklyListToSdk(ctx context.Context, list types.List) ([]objects.ExternalDynamicListsTypeIpRecurringWeekly, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ExternalDynamicListsTypeIpRecurringWeekly")
	diags := diag.Diagnostics{}
	var data []models.ExternalDynamicListsTypeIpRecurringWeekly
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.ExternalDynamicListsTypeIpRecurringWeekly, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ExternalDynamicListsTypeIpRecurringWeekly{}.AttrTypes(), &item)
		unpacked, d := unpackExternalDynamicListsTypeIpRecurringWeeklyToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ExternalDynamicListsTypeIpRecurringWeekly", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ExternalDynamicListsTypeIpRecurringWeekly ---
func packExternalDynamicListsTypeIpRecurringWeeklyListFromSdk(ctx context.Context, sdks []objects.ExternalDynamicListsTypeIpRecurringWeekly) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ExternalDynamicListsTypeIpRecurringWeekly")
	diags := diag.Diagnostics{}
	var data []models.ExternalDynamicListsTypeIpRecurringWeekly

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ExternalDynamicListsTypeIpRecurringWeekly
		obj, d := packExternalDynamicListsTypeIpRecurringWeeklyFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ExternalDynamicListsTypeIpRecurringWeekly{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ExternalDynamicListsTypeIpRecurringWeekly", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ExternalDynamicListsTypeIpRecurringWeekly{}.AttrType(), data)
}

// --- Unpacker for ExternalDynamicListsTypePredefinedIp ---
func unpackExternalDynamicListsTypePredefinedIpToSdk(ctx context.Context, obj types.Object) (*objects.ExternalDynamicListsTypePredefinedIp, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ExternalDynamicListsTypePredefinedIp", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ExternalDynamicListsTypePredefinedIp
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.ExternalDynamicListsTypePredefinedIp
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Description.IsNull() && !model.Description.IsUnknown() {
		sdk.Description = model.Description.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Description", "value": *sdk.Description})
	}

	// Handling Lists
	if !model.ExceptionList.IsNull() && !model.ExceptionList.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field ExceptionList")
		diags.Append(model.ExceptionList.ElementsAs(ctx, &sdk.ExceptionList, false)...)
	}

	// Handling Primitives
	if !model.Url.IsNull() && !model.Url.IsUnknown() {
		sdk.Url = model.Url.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Url", "value": sdk.Url})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ExternalDynamicListsTypePredefinedIp", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ExternalDynamicListsTypePredefinedIp ---
func packExternalDynamicListsTypePredefinedIpFromSdk(ctx context.Context, sdk objects.ExternalDynamicListsTypePredefinedIp) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ExternalDynamicListsTypePredefinedIp", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ExternalDynamicListsTypePredefinedIp
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Description != nil {
		model.Description = basetypes.NewStringValue(*sdk.Description)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Description", "value": *sdk.Description})
	} else {
		model.Description = basetypes.NewStringNull()
	}
	// Handling Lists
	if sdk.ExceptionList != nil {
		tflog.Debug(ctx, "Packing list of primitives for field ExceptionList")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.ExceptionList, d = basetypes.NewListValueFrom(ctx, elemType, sdk.ExceptionList)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.ExceptionList = basetypes.NewListNull(elemType)
	}
	// Handling Primitives
	// Standard primitive packing
	model.Url = basetypes.NewStringValue(sdk.Url)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Url", "value": sdk.Url})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ExternalDynamicListsTypePredefinedIp{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ExternalDynamicListsTypePredefinedIp", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ExternalDynamicListsTypePredefinedIp ---
func unpackExternalDynamicListsTypePredefinedIpListToSdk(ctx context.Context, list types.List) ([]objects.ExternalDynamicListsTypePredefinedIp, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ExternalDynamicListsTypePredefinedIp")
	diags := diag.Diagnostics{}
	var data []models.ExternalDynamicListsTypePredefinedIp
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.ExternalDynamicListsTypePredefinedIp, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ExternalDynamicListsTypePredefinedIp{}.AttrTypes(), &item)
		unpacked, d := unpackExternalDynamicListsTypePredefinedIpToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ExternalDynamicListsTypePredefinedIp", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ExternalDynamicListsTypePredefinedIp ---
func packExternalDynamicListsTypePredefinedIpListFromSdk(ctx context.Context, sdks []objects.ExternalDynamicListsTypePredefinedIp) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ExternalDynamicListsTypePredefinedIp")
	diags := diag.Diagnostics{}
	var data []models.ExternalDynamicListsTypePredefinedIp

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ExternalDynamicListsTypePredefinedIp
		obj, d := packExternalDynamicListsTypePredefinedIpFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ExternalDynamicListsTypePredefinedIp{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ExternalDynamicListsTypePredefinedIp", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ExternalDynamicListsTypePredefinedIp{}.AttrType(), data)
}

// --- Unpacker for ExternalDynamicListsTypePredefinedUrl ---
func unpackExternalDynamicListsTypePredefinedUrlToSdk(ctx context.Context, obj types.Object) (*objects.ExternalDynamicListsTypePredefinedUrl, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ExternalDynamicListsTypePredefinedUrl", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ExternalDynamicListsTypePredefinedUrl
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.ExternalDynamicListsTypePredefinedUrl
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Description.IsNull() && !model.Description.IsUnknown() {
		sdk.Description = model.Description.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Description", "value": *sdk.Description})
	}

	// Handling Lists
	if !model.ExceptionList.IsNull() && !model.ExceptionList.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field ExceptionList")
		diags.Append(model.ExceptionList.ElementsAs(ctx, &sdk.ExceptionList, false)...)
	}

	// Handling Primitives
	if !model.Url.IsNull() && !model.Url.IsUnknown() {
		sdk.Url = model.Url.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Url", "value": sdk.Url})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ExternalDynamicListsTypePredefinedUrl", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ExternalDynamicListsTypePredefinedUrl ---
func packExternalDynamicListsTypePredefinedUrlFromSdk(ctx context.Context, sdk objects.ExternalDynamicListsTypePredefinedUrl) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ExternalDynamicListsTypePredefinedUrl", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ExternalDynamicListsTypePredefinedUrl
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Description != nil {
		model.Description = basetypes.NewStringValue(*sdk.Description)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Description", "value": *sdk.Description})
	} else {
		model.Description = basetypes.NewStringNull()
	}
	// Handling Lists
	if sdk.ExceptionList != nil {
		tflog.Debug(ctx, "Packing list of primitives for field ExceptionList")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.ExceptionList, d = basetypes.NewListValueFrom(ctx, elemType, sdk.ExceptionList)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.ExceptionList = basetypes.NewListNull(elemType)
	}
	// Handling Primitives
	// Standard primitive packing
	model.Url = basetypes.NewStringValue(sdk.Url)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Url", "value": sdk.Url})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ExternalDynamicListsTypePredefinedUrl{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ExternalDynamicListsTypePredefinedUrl", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ExternalDynamicListsTypePredefinedUrl ---
func unpackExternalDynamicListsTypePredefinedUrlListToSdk(ctx context.Context, list types.List) ([]objects.ExternalDynamicListsTypePredefinedUrl, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ExternalDynamicListsTypePredefinedUrl")
	diags := diag.Diagnostics{}
	var data []models.ExternalDynamicListsTypePredefinedUrl
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.ExternalDynamicListsTypePredefinedUrl, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ExternalDynamicListsTypePredefinedUrl{}.AttrTypes(), &item)
		unpacked, d := unpackExternalDynamicListsTypePredefinedUrlToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ExternalDynamicListsTypePredefinedUrl", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ExternalDynamicListsTypePredefinedUrl ---
func packExternalDynamicListsTypePredefinedUrlListFromSdk(ctx context.Context, sdks []objects.ExternalDynamicListsTypePredefinedUrl) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ExternalDynamicListsTypePredefinedUrl")
	diags := diag.Diagnostics{}
	var data []models.ExternalDynamicListsTypePredefinedUrl

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ExternalDynamicListsTypePredefinedUrl
		obj, d := packExternalDynamicListsTypePredefinedUrlFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ExternalDynamicListsTypePredefinedUrl{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ExternalDynamicListsTypePredefinedUrl", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ExternalDynamicListsTypePredefinedUrl{}.AttrType(), data)
}

// --- Unpacker for ExternalDynamicListsTypeUrl ---
func unpackExternalDynamicListsTypeUrlToSdk(ctx context.Context, obj types.Object) (*objects.ExternalDynamicListsTypeUrl, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ExternalDynamicListsTypeUrl", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ExternalDynamicListsTypeUrl
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.ExternalDynamicListsTypeUrl
	var d diag.Diagnostics
	// Handling Objects
	if !model.Auth.IsNull() && !model.Auth.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Auth")
		unpacked, d := unpackExternalDynamicListsTypeUrlAuthToSdk(ctx, model.Auth)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Auth"})
		}
		if unpacked != nil {
			sdk.Auth = unpacked
		}
	}

	// Handling Primitives
	if !model.CertificateProfile.IsNull() && !model.CertificateProfile.IsUnknown() {
		sdk.CertificateProfile = model.CertificateProfile.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "CertificateProfile", "value": *sdk.CertificateProfile})
	}

	// Handling Primitives
	if !model.Description.IsNull() && !model.Description.IsUnknown() {
		sdk.Description = model.Description.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Description", "value": *sdk.Description})
	}

	// Handling Lists
	if !model.ExceptionList.IsNull() && !model.ExceptionList.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field ExceptionList")
		diags.Append(model.ExceptionList.ElementsAs(ctx, &sdk.ExceptionList, false)...)
	}

	// Handling Objects
	if !model.Recurring.IsNull() && !model.Recurring.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Recurring")
		unpacked, d := unpackExternalDynamicListsTypeUrlRecurringToSdk(ctx, model.Recurring)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Recurring"})
		}
		if unpacked != nil {
			sdk.Recurring = *unpacked
		}
	}

	// Handling Primitives
	if !model.Url.IsNull() && !model.Url.IsUnknown() {
		sdk.Url = model.Url.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Url", "value": sdk.Url})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ExternalDynamicListsTypeUrl", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ExternalDynamicListsTypeUrl ---
func packExternalDynamicListsTypeUrlFromSdk(ctx context.Context, sdk objects.ExternalDynamicListsTypeUrl) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ExternalDynamicListsTypeUrl", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ExternalDynamicListsTypeUrl
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Auth != nil {
		tflog.Debug(ctx, "Packing nested object for field Auth")
		packed, d := packExternalDynamicListsTypeUrlAuthFromSdk(ctx, *sdk.Auth)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Auth"})
		}
		model.Auth = packed
	} else {
		model.Auth = basetypes.NewObjectNull(models.ExternalDynamicListsTypeUrlAuth{}.AttrTypes())
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
	if sdk.Description != nil {
		model.Description = basetypes.NewStringValue(*sdk.Description)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Description", "value": *sdk.Description})
	} else {
		model.Description = basetypes.NewStringNull()
	}
	// Handling Lists
	if sdk.ExceptionList != nil {
		tflog.Debug(ctx, "Packing list of primitives for field ExceptionList")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.ExceptionList, d = basetypes.NewListValueFrom(ctx, elemType, sdk.ExceptionList)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.ExceptionList = basetypes.NewListNull(elemType)
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	// Logic for non-pointer / value-type nested objects
	if !reflect.ValueOf(sdk.Recurring).IsZero() {
		tflog.Debug(ctx, "Packing nested object for field Recurring")
		packed, d := packExternalDynamicListsTypeUrlRecurringFromSdk(ctx, sdk.Recurring)
		diags.Append(d...)
		model.Recurring = packed
	} else {
		model.Recurring = basetypes.NewObjectNull(models.ExternalDynamicListsTypeUrlRecurring{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	model.Url = basetypes.NewStringValue(sdk.Url)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Url", "value": sdk.Url})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ExternalDynamicListsTypeUrl{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ExternalDynamicListsTypeUrl", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ExternalDynamicListsTypeUrl ---
func unpackExternalDynamicListsTypeUrlListToSdk(ctx context.Context, list types.List) ([]objects.ExternalDynamicListsTypeUrl, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ExternalDynamicListsTypeUrl")
	diags := diag.Diagnostics{}
	var data []models.ExternalDynamicListsTypeUrl
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.ExternalDynamicListsTypeUrl, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ExternalDynamicListsTypeUrl{}.AttrTypes(), &item)
		unpacked, d := unpackExternalDynamicListsTypeUrlToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ExternalDynamicListsTypeUrl", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ExternalDynamicListsTypeUrl ---
func packExternalDynamicListsTypeUrlListFromSdk(ctx context.Context, sdks []objects.ExternalDynamicListsTypeUrl) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ExternalDynamicListsTypeUrl")
	diags := diag.Diagnostics{}
	var data []models.ExternalDynamicListsTypeUrl

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ExternalDynamicListsTypeUrl
		obj, d := packExternalDynamicListsTypeUrlFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ExternalDynamicListsTypeUrl{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ExternalDynamicListsTypeUrl", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ExternalDynamicListsTypeUrl{}.AttrType(), data)
}

// --- Unpacker for ExternalDynamicListsTypeUrlAuth ---
func unpackExternalDynamicListsTypeUrlAuthToSdk(ctx context.Context, obj types.Object) (*objects.ExternalDynamicListsTypeUrlAuth, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ExternalDynamicListsTypeUrlAuth", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ExternalDynamicListsTypeUrlAuth
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.ExternalDynamicListsTypeUrlAuth
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Password.IsNull() && !model.Password.IsUnknown() {
		sdk.Password = model.Password.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Password", "value": sdk.Password})
	}

	// Handling Primitives
	if !model.Username.IsNull() && !model.Username.IsUnknown() {
		sdk.Username = model.Username.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Username", "value": sdk.Username})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ExternalDynamicListsTypeUrlAuth", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ExternalDynamicListsTypeUrlAuth ---
func packExternalDynamicListsTypeUrlAuthFromSdk(ctx context.Context, sdk objects.ExternalDynamicListsTypeUrlAuth) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ExternalDynamicListsTypeUrlAuth", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ExternalDynamicListsTypeUrlAuth
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.Password = basetypes.NewStringValue(sdk.Password)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Password", "value": sdk.Password})
	// Handling Primitives
	// Standard primitive packing
	model.Username = basetypes.NewStringValue(sdk.Username)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Username", "value": sdk.Username})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ExternalDynamicListsTypeUrlAuth{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ExternalDynamicListsTypeUrlAuth", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ExternalDynamicListsTypeUrlAuth ---
func unpackExternalDynamicListsTypeUrlAuthListToSdk(ctx context.Context, list types.List) ([]objects.ExternalDynamicListsTypeUrlAuth, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ExternalDynamicListsTypeUrlAuth")
	diags := diag.Diagnostics{}
	var data []models.ExternalDynamicListsTypeUrlAuth
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.ExternalDynamicListsTypeUrlAuth, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ExternalDynamicListsTypeUrlAuth{}.AttrTypes(), &item)
		unpacked, d := unpackExternalDynamicListsTypeUrlAuthToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ExternalDynamicListsTypeUrlAuth", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ExternalDynamicListsTypeUrlAuth ---
func packExternalDynamicListsTypeUrlAuthListFromSdk(ctx context.Context, sdks []objects.ExternalDynamicListsTypeUrlAuth) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ExternalDynamicListsTypeUrlAuth")
	diags := diag.Diagnostics{}
	var data []models.ExternalDynamicListsTypeUrlAuth

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ExternalDynamicListsTypeUrlAuth
		obj, d := packExternalDynamicListsTypeUrlAuthFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ExternalDynamicListsTypeUrlAuth{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ExternalDynamicListsTypeUrlAuth", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ExternalDynamicListsTypeUrlAuth{}.AttrType(), data)
}

// --- Unpacker for ExternalDynamicListsTypeUrlRecurring ---
func unpackExternalDynamicListsTypeUrlRecurringToSdk(ctx context.Context, obj types.Object) (*objects.ExternalDynamicListsTypeUrlRecurring, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ExternalDynamicListsTypeUrlRecurring", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ExternalDynamicListsTypeUrlRecurring
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.ExternalDynamicListsTypeUrlRecurring
	var d diag.Diagnostics
	// Handling Objects
	if !model.Daily.IsNull() && !model.Daily.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Daily")
		unpacked, d := unpackExternalDynamicListsTypeUrlRecurringDailyToSdk(ctx, model.Daily)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Daily"})
		}
		if unpacked != nil {
			sdk.Daily = unpacked
		}
	}

	// Handling Typeless Objects
	if !model.FiveMinute.IsNull() && !model.FiveMinute.IsUnknown() {
		tflog.Debug(ctx, "Unpacking typeless object for field FiveMinute")
		sdk.FiveMinute = make(map[string]interface{})
	}

	// Handling Typeless Objects
	if !model.Hourly.IsNull() && !model.Hourly.IsUnknown() {
		tflog.Debug(ctx, "Unpacking typeless object for field Hourly")
		sdk.Hourly = make(map[string]interface{})
	}

	// Handling Objects
	if !model.Monthly.IsNull() && !model.Monthly.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Monthly")
		unpacked, d := unpackExternalDynamicListsTypeUrlRecurringMonthlyToSdk(ctx, model.Monthly)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Monthly"})
		}
		if unpacked != nil {
			sdk.Monthly = unpacked
		}
	}

	// Handling Objects
	if !model.Weekly.IsNull() && !model.Weekly.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Weekly")
		unpacked, d := unpackExternalDynamicListsTypeUrlRecurringWeeklyToSdk(ctx, model.Weekly)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Weekly"})
		}
		if unpacked != nil {
			sdk.Weekly = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ExternalDynamicListsTypeUrlRecurring", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ExternalDynamicListsTypeUrlRecurring ---
func packExternalDynamicListsTypeUrlRecurringFromSdk(ctx context.Context, sdk objects.ExternalDynamicListsTypeUrlRecurring) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ExternalDynamicListsTypeUrlRecurring", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ExternalDynamicListsTypeUrlRecurring
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Daily != nil {
		tflog.Debug(ctx, "Packing nested object for field Daily")
		packed, d := packExternalDynamicListsTypeUrlRecurringDailyFromSdk(ctx, *sdk.Daily)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Daily"})
		}
		model.Daily = packed
	} else {
		model.Daily = basetypes.NewObjectNull(models.ExternalDynamicListsTypeUrlRecurringDaily{}.AttrTypes())
	}
	// Handling Objects
	// This is a marker object (e.g. CHAP: {}). We just need to create an empty, non-null object.
	if sdk.FiveMinute != nil && !reflect.ValueOf(sdk.FiveMinute).IsNil() {
		tflog.Debug(ctx, "Packing typeless object for field FiveMinute")
		var d diag.Diagnostics
		// Create an empty object with no attributes, which signifies its presence.
		model.FiveMinute, d = basetypes.NewObjectValue(map[string]attr.Type{}, map[string]attr.Value{})
		diags.Append(d...)
	} else {
		// Since this field is part of a oneOf, being nil means it's not selected.
		// We make the object null with an empty attribute map.
		model.FiveMinute = basetypes.NewObjectNull(map[string]attr.Type{})
	}
	// Handling Objects
	// This is a marker object (e.g. CHAP: {}). We just need to create an empty, non-null object.
	if sdk.Hourly != nil && !reflect.ValueOf(sdk.Hourly).IsNil() {
		tflog.Debug(ctx, "Packing typeless object for field Hourly")
		var d diag.Diagnostics
		// Create an empty object with no attributes, which signifies its presence.
		model.Hourly, d = basetypes.NewObjectValue(map[string]attr.Type{}, map[string]attr.Value{})
		diags.Append(d...)
	} else {
		// Since this field is part of a oneOf, being nil means it's not selected.
		// We make the object null with an empty attribute map.
		model.Hourly = basetypes.NewObjectNull(map[string]attr.Type{})
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Monthly != nil {
		tflog.Debug(ctx, "Packing nested object for field Monthly")
		packed, d := packExternalDynamicListsTypeUrlRecurringMonthlyFromSdk(ctx, *sdk.Monthly)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Monthly"})
		}
		model.Monthly = packed
	} else {
		model.Monthly = basetypes.NewObjectNull(models.ExternalDynamicListsTypeUrlRecurringMonthly{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Weekly != nil {
		tflog.Debug(ctx, "Packing nested object for field Weekly")
		packed, d := packExternalDynamicListsTypeUrlRecurringWeeklyFromSdk(ctx, *sdk.Weekly)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Weekly"})
		}
		model.Weekly = packed
	} else {
		model.Weekly = basetypes.NewObjectNull(models.ExternalDynamicListsTypeUrlRecurringWeekly{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ExternalDynamicListsTypeUrlRecurring{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ExternalDynamicListsTypeUrlRecurring", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ExternalDynamicListsTypeUrlRecurring ---
func unpackExternalDynamicListsTypeUrlRecurringListToSdk(ctx context.Context, list types.List) ([]objects.ExternalDynamicListsTypeUrlRecurring, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ExternalDynamicListsTypeUrlRecurring")
	diags := diag.Diagnostics{}
	var data []models.ExternalDynamicListsTypeUrlRecurring
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.ExternalDynamicListsTypeUrlRecurring, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ExternalDynamicListsTypeUrlRecurring{}.AttrTypes(), &item)
		unpacked, d := unpackExternalDynamicListsTypeUrlRecurringToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ExternalDynamicListsTypeUrlRecurring", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ExternalDynamicListsTypeUrlRecurring ---
func packExternalDynamicListsTypeUrlRecurringListFromSdk(ctx context.Context, sdks []objects.ExternalDynamicListsTypeUrlRecurring) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ExternalDynamicListsTypeUrlRecurring")
	diags := diag.Diagnostics{}
	var data []models.ExternalDynamicListsTypeUrlRecurring

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ExternalDynamicListsTypeUrlRecurring
		obj, d := packExternalDynamicListsTypeUrlRecurringFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ExternalDynamicListsTypeUrlRecurring{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ExternalDynamicListsTypeUrlRecurring", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ExternalDynamicListsTypeUrlRecurring{}.AttrType(), data)
}

// --- Unpacker for ExternalDynamicListsTypeUrlRecurringDaily ---
func unpackExternalDynamicListsTypeUrlRecurringDailyToSdk(ctx context.Context, obj types.Object) (*objects.ExternalDynamicListsTypeUrlRecurringDaily, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ExternalDynamicListsTypeUrlRecurringDaily", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ExternalDynamicListsTypeUrlRecurringDaily
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.ExternalDynamicListsTypeUrlRecurringDaily
	var d diag.Diagnostics
	// Handling Primitives
	if !model.At.IsNull() && !model.At.IsUnknown() {
		sdk.At = model.At.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "At", "value": sdk.At})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ExternalDynamicListsTypeUrlRecurringDaily", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ExternalDynamicListsTypeUrlRecurringDaily ---
func packExternalDynamicListsTypeUrlRecurringDailyFromSdk(ctx context.Context, sdk objects.ExternalDynamicListsTypeUrlRecurringDaily) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ExternalDynamicListsTypeUrlRecurringDaily", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ExternalDynamicListsTypeUrlRecurringDaily
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.At = basetypes.NewStringValue(sdk.At)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "At", "value": sdk.At})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ExternalDynamicListsTypeUrlRecurringDaily{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ExternalDynamicListsTypeUrlRecurringDaily", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ExternalDynamicListsTypeUrlRecurringDaily ---
func unpackExternalDynamicListsTypeUrlRecurringDailyListToSdk(ctx context.Context, list types.List) ([]objects.ExternalDynamicListsTypeUrlRecurringDaily, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ExternalDynamicListsTypeUrlRecurringDaily")
	diags := diag.Diagnostics{}
	var data []models.ExternalDynamicListsTypeUrlRecurringDaily
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.ExternalDynamicListsTypeUrlRecurringDaily, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ExternalDynamicListsTypeUrlRecurringDaily{}.AttrTypes(), &item)
		unpacked, d := unpackExternalDynamicListsTypeUrlRecurringDailyToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ExternalDynamicListsTypeUrlRecurringDaily", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ExternalDynamicListsTypeUrlRecurringDaily ---
func packExternalDynamicListsTypeUrlRecurringDailyListFromSdk(ctx context.Context, sdks []objects.ExternalDynamicListsTypeUrlRecurringDaily) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ExternalDynamicListsTypeUrlRecurringDaily")
	diags := diag.Diagnostics{}
	var data []models.ExternalDynamicListsTypeUrlRecurringDaily

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ExternalDynamicListsTypeUrlRecurringDaily
		obj, d := packExternalDynamicListsTypeUrlRecurringDailyFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ExternalDynamicListsTypeUrlRecurringDaily{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ExternalDynamicListsTypeUrlRecurringDaily", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ExternalDynamicListsTypeUrlRecurringDaily{}.AttrType(), data)
}

// --- Unpacker for ExternalDynamicListsTypeUrlRecurringMonthly ---
func unpackExternalDynamicListsTypeUrlRecurringMonthlyToSdk(ctx context.Context, obj types.Object) (*objects.ExternalDynamicListsTypeUrlRecurringMonthly, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ExternalDynamicListsTypeUrlRecurringMonthly", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ExternalDynamicListsTypeUrlRecurringMonthly
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.ExternalDynamicListsTypeUrlRecurringMonthly
	var d diag.Diagnostics
	// Handling Primitives
	if !model.At.IsNull() && !model.At.IsUnknown() {
		sdk.At = model.At.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "At", "value": sdk.At})
	}

	// Handling Primitives
	if !model.DayOfMonth.IsNull() && !model.DayOfMonth.IsUnknown() {
		sdk.DayOfMonth = int32(model.DayOfMonth.ValueInt64())
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "DayOfMonth", "value": sdk.DayOfMonth})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ExternalDynamicListsTypeUrlRecurringMonthly", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ExternalDynamicListsTypeUrlRecurringMonthly ---
func packExternalDynamicListsTypeUrlRecurringMonthlyFromSdk(ctx context.Context, sdk objects.ExternalDynamicListsTypeUrlRecurringMonthly) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ExternalDynamicListsTypeUrlRecurringMonthly", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ExternalDynamicListsTypeUrlRecurringMonthly
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.At = basetypes.NewStringValue(sdk.At)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "At", "value": sdk.At})
	// Handling Primitives
	// Standard primitive packing
	model.DayOfMonth = basetypes.NewInt64Value(int64(sdk.DayOfMonth))
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "DayOfMonth", "value": sdk.DayOfMonth})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ExternalDynamicListsTypeUrlRecurringMonthly{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ExternalDynamicListsTypeUrlRecurringMonthly", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ExternalDynamicListsTypeUrlRecurringMonthly ---
func unpackExternalDynamicListsTypeUrlRecurringMonthlyListToSdk(ctx context.Context, list types.List) ([]objects.ExternalDynamicListsTypeUrlRecurringMonthly, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ExternalDynamicListsTypeUrlRecurringMonthly")
	diags := diag.Diagnostics{}
	var data []models.ExternalDynamicListsTypeUrlRecurringMonthly
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.ExternalDynamicListsTypeUrlRecurringMonthly, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ExternalDynamicListsTypeUrlRecurringMonthly{}.AttrTypes(), &item)
		unpacked, d := unpackExternalDynamicListsTypeUrlRecurringMonthlyToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ExternalDynamicListsTypeUrlRecurringMonthly", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ExternalDynamicListsTypeUrlRecurringMonthly ---
func packExternalDynamicListsTypeUrlRecurringMonthlyListFromSdk(ctx context.Context, sdks []objects.ExternalDynamicListsTypeUrlRecurringMonthly) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ExternalDynamicListsTypeUrlRecurringMonthly")
	diags := diag.Diagnostics{}
	var data []models.ExternalDynamicListsTypeUrlRecurringMonthly

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ExternalDynamicListsTypeUrlRecurringMonthly
		obj, d := packExternalDynamicListsTypeUrlRecurringMonthlyFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ExternalDynamicListsTypeUrlRecurringMonthly{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ExternalDynamicListsTypeUrlRecurringMonthly", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ExternalDynamicListsTypeUrlRecurringMonthly{}.AttrType(), data)
}

// --- Unpacker for ExternalDynamicListsTypeUrlRecurringWeekly ---
func unpackExternalDynamicListsTypeUrlRecurringWeeklyToSdk(ctx context.Context, obj types.Object) (*objects.ExternalDynamicListsTypeUrlRecurringWeekly, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ExternalDynamicListsTypeUrlRecurringWeekly", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ExternalDynamicListsTypeUrlRecurringWeekly
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.ExternalDynamicListsTypeUrlRecurringWeekly
	var d diag.Diagnostics
	// Handling Primitives
	if !model.At.IsNull() && !model.At.IsUnknown() {
		sdk.At = model.At.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "At", "value": sdk.At})
	}

	// Handling Primitives
	if !model.DayOfWeek.IsNull() && !model.DayOfWeek.IsUnknown() {
		sdk.DayOfWeek = model.DayOfWeek.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "DayOfWeek", "value": sdk.DayOfWeek})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ExternalDynamicListsTypeUrlRecurringWeekly", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ExternalDynamicListsTypeUrlRecurringWeekly ---
func packExternalDynamicListsTypeUrlRecurringWeeklyFromSdk(ctx context.Context, sdk objects.ExternalDynamicListsTypeUrlRecurringWeekly) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ExternalDynamicListsTypeUrlRecurringWeekly", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ExternalDynamicListsTypeUrlRecurringWeekly
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.At = basetypes.NewStringValue(sdk.At)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "At", "value": sdk.At})
	// Handling Primitives
	// Standard primitive packing
	model.DayOfWeek = basetypes.NewStringValue(sdk.DayOfWeek)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "DayOfWeek", "value": sdk.DayOfWeek})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ExternalDynamicListsTypeUrlRecurringWeekly{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ExternalDynamicListsTypeUrlRecurringWeekly", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ExternalDynamicListsTypeUrlRecurringWeekly ---
func unpackExternalDynamicListsTypeUrlRecurringWeeklyListToSdk(ctx context.Context, list types.List) ([]objects.ExternalDynamicListsTypeUrlRecurringWeekly, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ExternalDynamicListsTypeUrlRecurringWeekly")
	diags := diag.Diagnostics{}
	var data []models.ExternalDynamicListsTypeUrlRecurringWeekly
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.ExternalDynamicListsTypeUrlRecurringWeekly, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ExternalDynamicListsTypeUrlRecurringWeekly{}.AttrTypes(), &item)
		unpacked, d := unpackExternalDynamicListsTypeUrlRecurringWeeklyToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ExternalDynamicListsTypeUrlRecurringWeekly", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ExternalDynamicListsTypeUrlRecurringWeekly ---
func packExternalDynamicListsTypeUrlRecurringWeeklyListFromSdk(ctx context.Context, sdks []objects.ExternalDynamicListsTypeUrlRecurringWeekly) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ExternalDynamicListsTypeUrlRecurringWeekly")
	diags := diag.Diagnostics{}
	var data []models.ExternalDynamicListsTypeUrlRecurringWeekly

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ExternalDynamicListsTypeUrlRecurringWeekly
		obj, d := packExternalDynamicListsTypeUrlRecurringWeeklyFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ExternalDynamicListsTypeUrlRecurringWeekly{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ExternalDynamicListsTypeUrlRecurringWeekly", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ExternalDynamicListsTypeUrlRecurringWeekly{}.AttrType(), data)
}
