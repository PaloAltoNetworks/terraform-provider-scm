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

// ldapServerProfilesSensitiveValuePatcher is an in-memory struct to temporarily store plaintext
// and encrypted values for sensitive fields during the Create/Update/Read workflows.
type ldapServerProfilesSensitiveValuePatcher struct {
	bind_password_plaintext basetypes.StringValue
	bind_password_encrypted basetypes.StringValue
}

// populatePatcherFromState populates the patcher struct from the resource's state.
func (p *ldapServerProfilesSensitiveValuePatcher) populatePatcherFromState(ctx context.Context, state models.LdapServerProfiles) diag.Diagnostics {
	var diags diag.Diagnostics
	if state.EncryptedValues.IsNull() || state.EncryptedValues.IsUnknown() {
		return diags
	}

	var ev map[string]string
	diags.Append(state.EncryptedValues.ElementsAs(ctx, &ev, false)...)
	if diags.HasError() {
		return diags
	}
	if val, ok := ev["bind_password_plaintext"]; ok {
		p.bind_password_plaintext = basetypes.NewStringValue(val)
	}
	if val, ok := ev["bind_password_encrypted"]; ok {
		p.bind_password_encrypted = basetypes.NewStringValue(val)
	}

	return diags
}

// populateEncryptedValuesMap returns a map of the patcher's values for saving to state.
func (p *ldapServerProfilesSensitiveValuePatcher) populateEncryptedValuesMap() map[string]string {
	ev := make(map[string]string)
	if !p.bind_password_plaintext.IsNull() {
		ev["bind_password_plaintext"] = p.bind_password_plaintext.ValueString()
	}
	if !p.bind_password_encrypted.IsNull() {
		ev["bind_password_encrypted"] = p.bind_password_encrypted.ValueString()
	}
	return ev
}

// --- Unpacker for LdapServerProfiles ---
func unpackLdapServerProfilesToSdk(ctx context.Context, obj types.Object) (*identity_services.LdapServerProfiles, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.LdapServerProfiles", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.LdapServerProfiles
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk identity_services.LdapServerProfiles
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Base.IsNull() && !model.Base.IsUnknown() {
		sdk.Base = model.Base.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Base", "value": *sdk.Base})
	}

	// Handling Primitives
	if !model.BindDn.IsNull() && !model.BindDn.IsUnknown() {
		sdk.BindDn = model.BindDn.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "BindDn", "value": *sdk.BindDn})
	}

	// Handling Primitives
	if !model.BindPassword.IsNull() && !model.BindPassword.IsUnknown() {
		sdk.BindPassword = model.BindPassword.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "BindPassword", "value": *sdk.BindPassword})
	}

	// Handling Primitives
	if !model.BindTimelimit.IsNull() && !model.BindTimelimit.IsUnknown() {
		sdk.BindTimelimit = model.BindTimelimit.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "BindTimelimit", "value": *sdk.BindTimelimit})
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
	if !model.Id.IsNull() && !model.Id.IsUnknown() {
		sdk.Id = model.Id.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Id", "value": sdk.Id})
	}

	// Handling Primitives
	if !model.LdapType.IsNull() && !model.LdapType.IsUnknown() {
		sdk.LdapType = model.LdapType.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "LdapType", "value": *sdk.LdapType})
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	// Handling Primitives
	if !model.RetryInterval.IsNull() && !model.RetryInterval.IsUnknown() {
		val := int32(model.RetryInterval.ValueInt64())
		sdk.RetryInterval = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "RetryInterval", "value": *sdk.RetryInterval})
	}

	// Handling Lists
	if !model.Server.IsNull() && !model.Server.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field Server")
		unpacked, d := unpackLdapServerProfilesServerInnerListToSdk(ctx, model.Server)
		diags.Append(d...)
		sdk.Server = unpacked
	}

	// Handling Primitives
	if !model.Snippet.IsNull() && !model.Snippet.IsUnknown() {
		sdk.Snippet = model.Snippet.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	}

	// Handling Primitives
	if !model.Ssl.IsNull() && !model.Ssl.IsUnknown() {
		sdk.Ssl = model.Ssl.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Ssl", "value": *sdk.Ssl})
	}

	// Handling Primitives
	if !model.Timelimit.IsNull() && !model.Timelimit.IsUnknown() {
		val := int32(model.Timelimit.ValueInt64())
		sdk.Timelimit = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Timelimit", "value": *sdk.Timelimit})
	}

	// Handling Primitives
	if !model.VerifyServerCertificate.IsNull() && !model.VerifyServerCertificate.IsUnknown() {
		sdk.VerifyServerCertificate = model.VerifyServerCertificate.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "VerifyServerCertificate", "value": *sdk.VerifyServerCertificate})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.LdapServerProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for LdapServerProfiles ---
func packLdapServerProfilesFromSdk(ctx context.Context, sdk identity_services.LdapServerProfiles) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.LdapServerProfiles", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.LdapServerProfiles
	var d diag.Diagnostics
	// MEGA FIX FOR MAP TYPE MISMATCH (NOT ALL MODELS MAY HAVE EncryptedValues)
	model.EncryptedValues = basetypes.NewMapNull(basetypes.StringType{})
	// Handling Primitives
	// Standard primitive packing
	if sdk.Base != nil {
		model.Base = basetypes.NewStringValue(*sdk.Base)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Base", "value": *sdk.Base})
	} else {
		model.Base = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.BindDn != nil {
		model.BindDn = basetypes.NewStringValue(*sdk.BindDn)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "BindDn", "value": *sdk.BindDn})
	} else {
		model.BindDn = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.BindPassword != nil {
		model.BindPassword = basetypes.NewStringValue(*sdk.BindPassword)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "BindPassword", "value": *sdk.BindPassword})
	} else {
		model.BindPassword = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.BindTimelimit != nil {
		model.BindTimelimit = basetypes.NewStringValue(*sdk.BindTimelimit)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "BindTimelimit", "value": *sdk.BindTimelimit})
	} else {
		model.BindTimelimit = basetypes.NewStringNull()
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
	model.Id = basetypes.NewStringValue(sdk.Id)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Id", "value": sdk.Id})
	// Handling Primitives
	// Standard primitive packing
	if sdk.LdapType != nil {
		model.LdapType = basetypes.NewStringValue(*sdk.LdapType)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "LdapType", "value": *sdk.LdapType})
	} else {
		model.LdapType = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	// Handling Primitives
	// Standard primitive packing
	if sdk.RetryInterval != nil {
		model.RetryInterval = basetypes.NewInt64Value(int64(*sdk.RetryInterval))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "RetryInterval", "value": *sdk.RetryInterval})
	} else {
		model.RetryInterval = basetypes.NewInt64Null()
	}
	// Handling Lists
	if sdk.Server != nil {
		tflog.Debug(ctx, "Packing list of objects for field Server")
		packed, d := packLdapServerProfilesServerInnerListFromSdk(ctx, sdk.Server)
		diags.Append(d...)
		model.Server = packed
	} else {
		model.Server = basetypes.NewListNull(models.LdapServerProfilesServerInner{}.AttrType())
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
	if sdk.Ssl != nil {
		model.Ssl = basetypes.NewBoolValue(*sdk.Ssl)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Ssl", "value": *sdk.Ssl})
	} else {
		model.Ssl = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Timelimit != nil {
		model.Timelimit = basetypes.NewInt64Value(int64(*sdk.Timelimit))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Timelimit", "value": *sdk.Timelimit})
	} else {
		model.Timelimit = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.VerifyServerCertificate != nil {
		model.VerifyServerCertificate = basetypes.NewBoolValue(*sdk.VerifyServerCertificate)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "VerifyServerCertificate", "value": *sdk.VerifyServerCertificate})
	} else {
		model.VerifyServerCertificate = basetypes.NewBoolNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.LdapServerProfiles{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.LdapServerProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for LdapServerProfiles ---
func unpackLdapServerProfilesListToSdk(ctx context.Context, list types.List) ([]identity_services.LdapServerProfiles, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.LdapServerProfiles")
	diags := diag.Diagnostics{}
	var data []models.LdapServerProfiles
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]identity_services.LdapServerProfiles, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.LdapServerProfiles{}.AttrTypes(), &item)
		unpacked, d := unpackLdapServerProfilesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.LdapServerProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for LdapServerProfiles ---
func packLdapServerProfilesListFromSdk(ctx context.Context, sdks []identity_services.LdapServerProfiles) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.LdapServerProfiles")
	diags := diag.Diagnostics{}
	var data []models.LdapServerProfiles

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.LdapServerProfiles
		obj, d := packLdapServerProfilesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.LdapServerProfiles{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.LdapServerProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.LdapServerProfiles{}.AttrType(), data)
}

// --- Unpacker for LdapServerProfilesServerInner ---
func unpackLdapServerProfilesServerInnerToSdk(ctx context.Context, obj types.Object) (*identity_services.LdapServerProfilesServerInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.LdapServerProfilesServerInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.LdapServerProfilesServerInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk identity_services.LdapServerProfilesServerInner
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Address.IsNull() && !model.Address.IsUnknown() {
		sdk.Address = model.Address.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Address", "value": *sdk.Address})
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

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.LdapServerProfilesServerInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for LdapServerProfilesServerInner ---
func packLdapServerProfilesServerInnerFromSdk(ctx context.Context, sdk identity_services.LdapServerProfilesServerInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.LdapServerProfilesServerInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.LdapServerProfilesServerInner
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Address != nil {
		model.Address = basetypes.NewStringValue(*sdk.Address)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Address", "value": *sdk.Address})
	} else {
		model.Address = basetypes.NewStringNull()
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
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.LdapServerProfilesServerInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.LdapServerProfilesServerInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for LdapServerProfilesServerInner ---
func unpackLdapServerProfilesServerInnerListToSdk(ctx context.Context, list types.List) ([]identity_services.LdapServerProfilesServerInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.LdapServerProfilesServerInner")
	diags := diag.Diagnostics{}
	var data []models.LdapServerProfilesServerInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]identity_services.LdapServerProfilesServerInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.LdapServerProfilesServerInner{}.AttrTypes(), &item)
		unpacked, d := unpackLdapServerProfilesServerInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.LdapServerProfilesServerInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for LdapServerProfilesServerInner ---
func packLdapServerProfilesServerInnerListFromSdk(ctx context.Context, sdks []identity_services.LdapServerProfilesServerInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.LdapServerProfilesServerInner")
	diags := diag.Diagnostics{}
	var data []models.LdapServerProfilesServerInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.LdapServerProfilesServerInner
		obj, d := packLdapServerProfilesServerInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.LdapServerProfilesServerInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.LdapServerProfilesServerInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.LdapServerProfilesServerInner{}.AttrType(), data)
}
