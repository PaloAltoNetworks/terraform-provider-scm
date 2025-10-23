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

// --- Unpacker for AuthenticationProfiles ---
func unpackAuthenticationProfilesToSdk(ctx context.Context, obj types.Object) (*identity_services.AuthenticationProfiles, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.AuthenticationProfiles", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.AuthenticationProfiles
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk identity_services.AuthenticationProfiles
	var d diag.Diagnostics

	// Handling Lists
	if !model.AllowList.IsNull() && !model.AllowList.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field AllowList")
		diags.Append(model.AllowList.ElementsAs(ctx, &sdk.AllowList, false)...)
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

	// Handling Objects
	if !model.Lockout.IsNull() && !model.Lockout.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Lockout")
		unpacked, d := unpackAuthenticationProfilesLockoutToSdk(ctx, model.Lockout)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Lockout"})
		}
		if unpacked != nil {
			sdk.Lockout = unpacked
		}
	}

	// Handling Objects
	if !model.Method.IsNull() && !model.Method.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Method")
		unpacked, d := unpackAuthenticationProfilesMethodToSdk(ctx, model.Method)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Method"})
		}
		if unpacked != nil {
			sdk.Method = unpacked
		}
	}

	// Handling Objects
	if !model.MultiFactorAuth.IsNull() && !model.MultiFactorAuth.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field MultiFactorAuth")
		unpacked, d := unpackAuthenticationProfilesMultiFactorAuthToSdk(ctx, model.MultiFactorAuth)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "MultiFactorAuth"})
		}
		if unpacked != nil {
			sdk.MultiFactorAuth = unpacked
		}
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	// Handling Objects
	if !model.SingleSignOn.IsNull() && !model.SingleSignOn.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field SingleSignOn")
		unpacked, d := unpackAuthenticationProfilesSingleSignOnToSdk(ctx, model.SingleSignOn)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "SingleSignOn"})
		}
		if unpacked != nil {
			sdk.SingleSignOn = unpacked
		}
	}

	// Handling Primitives
	if !model.Snippet.IsNull() && !model.Snippet.IsUnknown() {
		sdk.Snippet = model.Snippet.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	}

	// Handling Primitives
	if !model.UserDomain.IsNull() && !model.UserDomain.IsUnknown() {
		sdk.UserDomain = model.UserDomain.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "UserDomain", "value": *sdk.UserDomain})
	}

	// Handling Primitives
	if !model.UsernameModifier.IsNull() && !model.UsernameModifier.IsUnknown() {
		sdk.UsernameModifier = model.UsernameModifier.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "UsernameModifier", "value": *sdk.UsernameModifier})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.AuthenticationProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for AuthenticationProfiles ---
func packAuthenticationProfilesFromSdk(ctx context.Context, sdk identity_services.AuthenticationProfiles) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.AuthenticationProfiles", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.AuthenticationProfiles
	var d diag.Diagnostics
	// Handling Lists
	if sdk.AllowList != nil {
		tflog.Debug(ctx, "Packing list of primitives for field AllowList")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.AllowList, d = basetypes.NewListValueFrom(ctx, elemType, sdk.AllowList)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.AllowList = basetypes.NewListNull(elemType)
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
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Lockout != nil {
		tflog.Debug(ctx, "Packing nested object for field Lockout")
		packed, d := packAuthenticationProfilesLockoutFromSdk(ctx, *sdk.Lockout)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Lockout"})
		}
		model.Lockout = packed
	} else {
		model.Lockout = basetypes.NewObjectNull(models.AuthenticationProfilesLockout{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Method != nil {
		tflog.Debug(ctx, "Packing nested object for field Method")
		packed, d := packAuthenticationProfilesMethodFromSdk(ctx, *sdk.Method)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Method"})
		}
		model.Method = packed
	} else {
		model.Method = basetypes.NewObjectNull(models.AuthenticationProfilesMethod{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.MultiFactorAuth != nil {
		tflog.Debug(ctx, "Packing nested object for field MultiFactorAuth")
		packed, d := packAuthenticationProfilesMultiFactorAuthFromSdk(ctx, *sdk.MultiFactorAuth)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "MultiFactorAuth"})
		}
		model.MultiFactorAuth = packed
	} else {
		model.MultiFactorAuth = basetypes.NewObjectNull(models.AuthenticationProfilesMultiFactorAuth{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.SingleSignOn != nil {
		tflog.Debug(ctx, "Packing nested object for field SingleSignOn")
		packed, d := packAuthenticationProfilesSingleSignOnFromSdk(ctx, *sdk.SingleSignOn)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "SingleSignOn"})
		}
		model.SingleSignOn = packed
	} else {
		model.SingleSignOn = basetypes.NewObjectNull(models.AuthenticationProfilesSingleSignOn{}.AttrTypes())
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
	if sdk.UserDomain != nil {
		model.UserDomain = basetypes.NewStringValue(*sdk.UserDomain)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "UserDomain", "value": *sdk.UserDomain})
	} else {
		model.UserDomain = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.UsernameModifier != nil {
		model.UsernameModifier = basetypes.NewStringValue(*sdk.UsernameModifier)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "UsernameModifier", "value": *sdk.UsernameModifier})
	} else {
		model.UsernameModifier = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.AuthenticationProfiles{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.AuthenticationProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for AuthenticationProfiles ---
func unpackAuthenticationProfilesListToSdk(ctx context.Context, list types.List) ([]identity_services.AuthenticationProfiles, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.AuthenticationProfiles")
	diags := diag.Diagnostics{}
	var data []models.AuthenticationProfiles
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]identity_services.AuthenticationProfiles, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.AuthenticationProfiles{}.AttrTypes(), &item)
		unpacked, d := unpackAuthenticationProfilesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.AuthenticationProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for AuthenticationProfiles ---
func packAuthenticationProfilesListFromSdk(ctx context.Context, sdks []identity_services.AuthenticationProfiles) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.AuthenticationProfiles")
	diags := diag.Diagnostics{}
	var data []models.AuthenticationProfiles

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.AuthenticationProfiles
		obj, d := packAuthenticationProfilesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.AuthenticationProfiles{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.AuthenticationProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.AuthenticationProfiles{}.AttrType(), data)
}

// --- Unpacker for AuthenticationProfilesLockout ---
func unpackAuthenticationProfilesLockoutToSdk(ctx context.Context, obj types.Object) (*identity_services.AuthenticationProfilesLockout, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.AuthenticationProfilesLockout", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.AuthenticationProfilesLockout
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk identity_services.AuthenticationProfilesLockout
	var d diag.Diagnostics
	// Handling Primitives
	if !model.FailedAttempts.IsNull() && !model.FailedAttempts.IsUnknown() {
		val := int32(model.FailedAttempts.ValueInt64())
		sdk.FailedAttempts = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "FailedAttempts", "value": *sdk.FailedAttempts})
	}

	// Handling Primitives
	if !model.LockoutTime.IsNull() && !model.LockoutTime.IsUnknown() {
		val := int32(model.LockoutTime.ValueInt64())
		sdk.LockoutTime = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "LockoutTime", "value": *sdk.LockoutTime})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.AuthenticationProfilesLockout", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for AuthenticationProfilesLockout ---
func packAuthenticationProfilesLockoutFromSdk(ctx context.Context, sdk identity_services.AuthenticationProfilesLockout) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.AuthenticationProfilesLockout", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.AuthenticationProfilesLockout
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.FailedAttempts != nil {
		model.FailedAttempts = basetypes.NewInt64Value(int64(*sdk.FailedAttempts))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "FailedAttempts", "value": *sdk.FailedAttempts})
	} else {
		model.FailedAttempts = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.LockoutTime != nil {
		model.LockoutTime = basetypes.NewInt64Value(int64(*sdk.LockoutTime))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "LockoutTime", "value": *sdk.LockoutTime})
	} else {
		model.LockoutTime = basetypes.NewInt64Null()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.AuthenticationProfilesLockout{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.AuthenticationProfilesLockout", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for AuthenticationProfilesLockout ---
func unpackAuthenticationProfilesLockoutListToSdk(ctx context.Context, list types.List) ([]identity_services.AuthenticationProfilesLockout, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.AuthenticationProfilesLockout")
	diags := diag.Diagnostics{}
	var data []models.AuthenticationProfilesLockout
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]identity_services.AuthenticationProfilesLockout, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.AuthenticationProfilesLockout{}.AttrTypes(), &item)
		unpacked, d := unpackAuthenticationProfilesLockoutToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.AuthenticationProfilesLockout", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for AuthenticationProfilesLockout ---
func packAuthenticationProfilesLockoutListFromSdk(ctx context.Context, sdks []identity_services.AuthenticationProfilesLockout) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.AuthenticationProfilesLockout")
	diags := diag.Diagnostics{}
	var data []models.AuthenticationProfilesLockout

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.AuthenticationProfilesLockout
		obj, d := packAuthenticationProfilesLockoutFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.AuthenticationProfilesLockout{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.AuthenticationProfilesLockout", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.AuthenticationProfilesLockout{}.AttrType(), data)
}

// --- Unpacker for AuthenticationProfilesMethod ---
func unpackAuthenticationProfilesMethodToSdk(ctx context.Context, obj types.Object) (*identity_services.AuthenticationProfilesMethod, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.AuthenticationProfilesMethod", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.AuthenticationProfilesMethod
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk identity_services.AuthenticationProfilesMethod
	var d diag.Diagnostics
	// Handling Objects
	if !model.Cloud.IsNull() && !model.Cloud.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Cloud")
		unpacked, d := unpackAuthenticationProfilesMethodCloudToSdk(ctx, model.Cloud)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Cloud"})
		}
		if unpacked != nil {
			sdk.Cloud = unpacked
		}
	}

	// Handling Objects
	if !model.Kerberos.IsNull() && !model.Kerberos.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Kerberos")
		unpacked, d := unpackAuthenticationProfilesMethodKerberosToSdk(ctx, model.Kerberos)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Kerberos"})
		}
		if unpacked != nil {
			sdk.Kerberos = unpacked
		}
	}

	// Handling Objects
	if !model.Ldap.IsNull() && !model.Ldap.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Ldap")
		unpacked, d := unpackAuthenticationProfilesMethodLdapToSdk(ctx, model.Ldap)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Ldap"})
		}
		if unpacked != nil {
			sdk.Ldap = unpacked
		}
	}

	// Handling Typeless Objects
	if !model.LocalDatabase.IsNull() && !model.LocalDatabase.IsUnknown() {
		tflog.Debug(ctx, "Unpacking typeless object for field LocalDatabase")
		sdk.LocalDatabase = make(map[string]interface{})
	}

	// Handling Objects
	if !model.Radius.IsNull() && !model.Radius.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Radius")
		unpacked, d := unpackAuthenticationProfilesMethodRadiusToSdk(ctx, model.Radius)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Radius"})
		}
		if unpacked != nil {
			sdk.Radius = unpacked
		}
	}

	// Handling Objects
	if !model.SamlIdp.IsNull() && !model.SamlIdp.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field SamlIdp")
		unpacked, d := unpackAuthenticationProfilesMethodSamlIdpToSdk(ctx, model.SamlIdp)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "SamlIdp"})
		}
		if unpacked != nil {
			sdk.SamlIdp = unpacked
		}
	}

	// Handling Objects
	if !model.Tacplus.IsNull() && !model.Tacplus.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Tacplus")
		unpacked, d := unpackAuthenticationProfilesMethodRadiusToSdk(ctx, model.Tacplus)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Tacplus"})
		}
		if unpacked != nil {
			sdk.Tacplus = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.AuthenticationProfilesMethod", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for AuthenticationProfilesMethod ---
func packAuthenticationProfilesMethodFromSdk(ctx context.Context, sdk identity_services.AuthenticationProfilesMethod) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.AuthenticationProfilesMethod", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.AuthenticationProfilesMethod
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Cloud != nil {
		tflog.Debug(ctx, "Packing nested object for field Cloud")
		packed, d := packAuthenticationProfilesMethodCloudFromSdk(ctx, *sdk.Cloud)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Cloud"})
		}
		model.Cloud = packed
	} else {
		model.Cloud = basetypes.NewObjectNull(models.AuthenticationProfilesMethodCloud{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Kerberos != nil {
		tflog.Debug(ctx, "Packing nested object for field Kerberos")
		packed, d := packAuthenticationProfilesMethodKerberosFromSdk(ctx, *sdk.Kerberos)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Kerberos"})
		}
		model.Kerberos = packed
	} else {
		model.Kerberos = basetypes.NewObjectNull(models.AuthenticationProfilesMethodKerberos{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Ldap != nil {
		tflog.Debug(ctx, "Packing nested object for field Ldap")
		packed, d := packAuthenticationProfilesMethodLdapFromSdk(ctx, *sdk.Ldap)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Ldap"})
		}
		model.Ldap = packed
	} else {
		model.Ldap = basetypes.NewObjectNull(models.AuthenticationProfilesMethodLdap{}.AttrTypes())
	}
	// Handling Objects
	// This is a marker object (e.g. CHAP: {}). We just need to create an empty, non-null object.
	if sdk.LocalDatabase != nil && !reflect.ValueOf(sdk.LocalDatabase).IsNil() {
		tflog.Debug(ctx, "Packing typeless object for field LocalDatabase")
		var d diag.Diagnostics
		// Create an empty object with no attributes, which signifies its presence.
		model.LocalDatabase, d = basetypes.NewObjectValue(map[string]attr.Type{}, map[string]attr.Value{})
		diags.Append(d...)
	} else {
		// Since this field is part of a oneOf, being nil means it's not selected.
		// We make the object null with an empty attribute map.
		model.LocalDatabase = basetypes.NewObjectNull(map[string]attr.Type{})
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Radius != nil {
		tflog.Debug(ctx, "Packing nested object for field Radius")
		packed, d := packAuthenticationProfilesMethodRadiusFromSdk(ctx, *sdk.Radius)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Radius"})
		}
		model.Radius = packed
	} else {
		model.Radius = basetypes.NewObjectNull(models.AuthenticationProfilesMethodRadius{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.SamlIdp != nil {
		tflog.Debug(ctx, "Packing nested object for field SamlIdp")
		packed, d := packAuthenticationProfilesMethodSamlIdpFromSdk(ctx, *sdk.SamlIdp)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "SamlIdp"})
		}
		model.SamlIdp = packed
	} else {
		model.SamlIdp = basetypes.NewObjectNull(models.AuthenticationProfilesMethodSamlIdp{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Tacplus != nil {
		tflog.Debug(ctx, "Packing nested object for field Tacplus")
		packed, d := packAuthenticationProfilesMethodRadiusFromSdk(ctx, *sdk.Tacplus)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Tacplus"})
		}
		model.Tacplus = packed
	} else {
		model.Tacplus = basetypes.NewObjectNull(models.AuthenticationProfilesMethodRadius{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.AuthenticationProfilesMethod{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.AuthenticationProfilesMethod", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for AuthenticationProfilesMethod ---
func unpackAuthenticationProfilesMethodListToSdk(ctx context.Context, list types.List) ([]identity_services.AuthenticationProfilesMethod, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.AuthenticationProfilesMethod")
	diags := diag.Diagnostics{}
	var data []models.AuthenticationProfilesMethod
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]identity_services.AuthenticationProfilesMethod, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.AuthenticationProfilesMethod{}.AttrTypes(), &item)
		unpacked, d := unpackAuthenticationProfilesMethodToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.AuthenticationProfilesMethod", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for AuthenticationProfilesMethod ---
func packAuthenticationProfilesMethodListFromSdk(ctx context.Context, sdks []identity_services.AuthenticationProfilesMethod) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.AuthenticationProfilesMethod")
	diags := diag.Diagnostics{}
	var data []models.AuthenticationProfilesMethod

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.AuthenticationProfilesMethod
		obj, d := packAuthenticationProfilesMethodFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.AuthenticationProfilesMethod{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.AuthenticationProfilesMethod", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.AuthenticationProfilesMethod{}.AttrType(), data)
}

// --- Unpacker for AuthenticationProfilesMethodCloud ---
func unpackAuthenticationProfilesMethodCloudToSdk(ctx context.Context, obj types.Object) (*identity_services.AuthenticationProfilesMethodCloud, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.AuthenticationProfilesMethodCloud", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.AuthenticationProfilesMethodCloud
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk identity_services.AuthenticationProfilesMethodCloud
	var d diag.Diagnostics
	// Handling Primitives
	if !model.ProfileName.IsNull() && !model.ProfileName.IsUnknown() {
		sdk.ProfileName = model.ProfileName.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "ProfileName", "value": *sdk.ProfileName})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.AuthenticationProfilesMethodCloud", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for AuthenticationProfilesMethodCloud ---
func packAuthenticationProfilesMethodCloudFromSdk(ctx context.Context, sdk identity_services.AuthenticationProfilesMethodCloud) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.AuthenticationProfilesMethodCloud", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.AuthenticationProfilesMethodCloud
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.ProfileName != nil {
		model.ProfileName = basetypes.NewStringValue(*sdk.ProfileName)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "ProfileName", "value": *sdk.ProfileName})
	} else {
		model.ProfileName = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.AuthenticationProfilesMethodCloud{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.AuthenticationProfilesMethodCloud", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for AuthenticationProfilesMethodCloud ---
func unpackAuthenticationProfilesMethodCloudListToSdk(ctx context.Context, list types.List) ([]identity_services.AuthenticationProfilesMethodCloud, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.AuthenticationProfilesMethodCloud")
	diags := diag.Diagnostics{}
	var data []models.AuthenticationProfilesMethodCloud
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]identity_services.AuthenticationProfilesMethodCloud, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.AuthenticationProfilesMethodCloud{}.AttrTypes(), &item)
		unpacked, d := unpackAuthenticationProfilesMethodCloudToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.AuthenticationProfilesMethodCloud", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for AuthenticationProfilesMethodCloud ---
func packAuthenticationProfilesMethodCloudListFromSdk(ctx context.Context, sdks []identity_services.AuthenticationProfilesMethodCloud) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.AuthenticationProfilesMethodCloud")
	diags := diag.Diagnostics{}
	var data []models.AuthenticationProfilesMethodCloud

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.AuthenticationProfilesMethodCloud
		obj, d := packAuthenticationProfilesMethodCloudFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.AuthenticationProfilesMethodCloud{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.AuthenticationProfilesMethodCloud", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.AuthenticationProfilesMethodCloud{}.AttrType(), data)
}

// --- Unpacker for AuthenticationProfilesMethodKerberos ---
func unpackAuthenticationProfilesMethodKerberosToSdk(ctx context.Context, obj types.Object) (*identity_services.AuthenticationProfilesMethodKerberos, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.AuthenticationProfilesMethodKerberos", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.AuthenticationProfilesMethodKerberos
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk identity_services.AuthenticationProfilesMethodKerberos
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Realm.IsNull() && !model.Realm.IsUnknown() {
		sdk.Realm = model.Realm.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Realm", "value": *sdk.Realm})
	}

	// Handling Primitives
	if !model.ServerProfile.IsNull() && !model.ServerProfile.IsUnknown() {
		sdk.ServerProfile = model.ServerProfile.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "ServerProfile", "value": *sdk.ServerProfile})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.AuthenticationProfilesMethodKerberos", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for AuthenticationProfilesMethodKerberos ---
func packAuthenticationProfilesMethodKerberosFromSdk(ctx context.Context, sdk identity_services.AuthenticationProfilesMethodKerberos) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.AuthenticationProfilesMethodKerberos", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.AuthenticationProfilesMethodKerberos
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Realm != nil {
		model.Realm = basetypes.NewStringValue(*sdk.Realm)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Realm", "value": *sdk.Realm})
	} else {
		model.Realm = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.ServerProfile != nil {
		model.ServerProfile = basetypes.NewStringValue(*sdk.ServerProfile)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "ServerProfile", "value": *sdk.ServerProfile})
	} else {
		model.ServerProfile = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.AuthenticationProfilesMethodKerberos{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.AuthenticationProfilesMethodKerberos", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for AuthenticationProfilesMethodKerberos ---
func unpackAuthenticationProfilesMethodKerberosListToSdk(ctx context.Context, list types.List) ([]identity_services.AuthenticationProfilesMethodKerberos, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.AuthenticationProfilesMethodKerberos")
	diags := diag.Diagnostics{}
	var data []models.AuthenticationProfilesMethodKerberos
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]identity_services.AuthenticationProfilesMethodKerberos, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.AuthenticationProfilesMethodKerberos{}.AttrTypes(), &item)
		unpacked, d := unpackAuthenticationProfilesMethodKerberosToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.AuthenticationProfilesMethodKerberos", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for AuthenticationProfilesMethodKerberos ---
func packAuthenticationProfilesMethodKerberosListFromSdk(ctx context.Context, sdks []identity_services.AuthenticationProfilesMethodKerberos) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.AuthenticationProfilesMethodKerberos")
	diags := diag.Diagnostics{}
	var data []models.AuthenticationProfilesMethodKerberos

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.AuthenticationProfilesMethodKerberos
		obj, d := packAuthenticationProfilesMethodKerberosFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.AuthenticationProfilesMethodKerberos{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.AuthenticationProfilesMethodKerberos", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.AuthenticationProfilesMethodKerberos{}.AttrType(), data)
}

// --- Unpacker for AuthenticationProfilesMethodLdap ---
func unpackAuthenticationProfilesMethodLdapToSdk(ctx context.Context, obj types.Object) (*identity_services.AuthenticationProfilesMethodLdap, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.AuthenticationProfilesMethodLdap", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.AuthenticationProfilesMethodLdap
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk identity_services.AuthenticationProfilesMethodLdap
	var d diag.Diagnostics
	// Handling Primitives
	if !model.LoginAttribute.IsNull() && !model.LoginAttribute.IsUnknown() {
		sdk.LoginAttribute = model.LoginAttribute.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "LoginAttribute", "value": *sdk.LoginAttribute})
	}

	// Handling Primitives
	if !model.PasswdExpDays.IsNull() && !model.PasswdExpDays.IsUnknown() {
		val := int32(model.PasswdExpDays.ValueInt64())
		sdk.PasswdExpDays = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "PasswdExpDays", "value": *sdk.PasswdExpDays})
	}

	// Handling Primitives
	if !model.ServerProfile.IsNull() && !model.ServerProfile.IsUnknown() {
		sdk.ServerProfile = model.ServerProfile.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "ServerProfile", "value": *sdk.ServerProfile})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.AuthenticationProfilesMethodLdap", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for AuthenticationProfilesMethodLdap ---
func packAuthenticationProfilesMethodLdapFromSdk(ctx context.Context, sdk identity_services.AuthenticationProfilesMethodLdap) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.AuthenticationProfilesMethodLdap", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.AuthenticationProfilesMethodLdap
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.LoginAttribute != nil {
		model.LoginAttribute = basetypes.NewStringValue(*sdk.LoginAttribute)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "LoginAttribute", "value": *sdk.LoginAttribute})
	} else {
		model.LoginAttribute = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.PasswdExpDays != nil {
		model.PasswdExpDays = basetypes.NewInt64Value(int64(*sdk.PasswdExpDays))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "PasswdExpDays", "value": *sdk.PasswdExpDays})
	} else {
		model.PasswdExpDays = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.ServerProfile != nil {
		model.ServerProfile = basetypes.NewStringValue(*sdk.ServerProfile)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "ServerProfile", "value": *sdk.ServerProfile})
	} else {
		model.ServerProfile = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.AuthenticationProfilesMethodLdap{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.AuthenticationProfilesMethodLdap", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for AuthenticationProfilesMethodLdap ---
func unpackAuthenticationProfilesMethodLdapListToSdk(ctx context.Context, list types.List) ([]identity_services.AuthenticationProfilesMethodLdap, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.AuthenticationProfilesMethodLdap")
	diags := diag.Diagnostics{}
	var data []models.AuthenticationProfilesMethodLdap
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]identity_services.AuthenticationProfilesMethodLdap, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.AuthenticationProfilesMethodLdap{}.AttrTypes(), &item)
		unpacked, d := unpackAuthenticationProfilesMethodLdapToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.AuthenticationProfilesMethodLdap", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for AuthenticationProfilesMethodLdap ---
func packAuthenticationProfilesMethodLdapListFromSdk(ctx context.Context, sdks []identity_services.AuthenticationProfilesMethodLdap) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.AuthenticationProfilesMethodLdap")
	diags := diag.Diagnostics{}
	var data []models.AuthenticationProfilesMethodLdap

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.AuthenticationProfilesMethodLdap
		obj, d := packAuthenticationProfilesMethodLdapFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.AuthenticationProfilesMethodLdap{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.AuthenticationProfilesMethodLdap", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.AuthenticationProfilesMethodLdap{}.AttrType(), data)
}

// --- Unpacker for AuthenticationProfilesMethodRadius ---
func unpackAuthenticationProfilesMethodRadiusToSdk(ctx context.Context, obj types.Object) (*identity_services.AuthenticationProfilesMethodRadius, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.AuthenticationProfilesMethodRadius", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.AuthenticationProfilesMethodRadius
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk identity_services.AuthenticationProfilesMethodRadius
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Checkgroup.IsNull() && !model.Checkgroup.IsUnknown() {
		sdk.Checkgroup = model.Checkgroup.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Checkgroup", "value": *sdk.Checkgroup})
	}

	// Handling Primitives
	if !model.ServerProfile.IsNull() && !model.ServerProfile.IsUnknown() {
		sdk.ServerProfile = model.ServerProfile.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "ServerProfile", "value": *sdk.ServerProfile})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.AuthenticationProfilesMethodRadius", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for AuthenticationProfilesMethodRadius ---
func packAuthenticationProfilesMethodRadiusFromSdk(ctx context.Context, sdk identity_services.AuthenticationProfilesMethodRadius) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.AuthenticationProfilesMethodRadius", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.AuthenticationProfilesMethodRadius
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Checkgroup != nil {
		model.Checkgroup = basetypes.NewBoolValue(*sdk.Checkgroup)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Checkgroup", "value": *sdk.Checkgroup})
	} else {
		model.Checkgroup = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.ServerProfile != nil {
		model.ServerProfile = basetypes.NewStringValue(*sdk.ServerProfile)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "ServerProfile", "value": *sdk.ServerProfile})
	} else {
		model.ServerProfile = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.AuthenticationProfilesMethodRadius{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.AuthenticationProfilesMethodRadius", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for AuthenticationProfilesMethodRadius ---
func unpackAuthenticationProfilesMethodRadiusListToSdk(ctx context.Context, list types.List) ([]identity_services.AuthenticationProfilesMethodRadius, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.AuthenticationProfilesMethodRadius")
	diags := diag.Diagnostics{}
	var data []models.AuthenticationProfilesMethodRadius
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]identity_services.AuthenticationProfilesMethodRadius, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.AuthenticationProfilesMethodRadius{}.AttrTypes(), &item)
		unpacked, d := unpackAuthenticationProfilesMethodRadiusToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.AuthenticationProfilesMethodRadius", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for AuthenticationProfilesMethodRadius ---
func packAuthenticationProfilesMethodRadiusListFromSdk(ctx context.Context, sdks []identity_services.AuthenticationProfilesMethodRadius) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.AuthenticationProfilesMethodRadius")
	diags := diag.Diagnostics{}
	var data []models.AuthenticationProfilesMethodRadius

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.AuthenticationProfilesMethodRadius
		obj, d := packAuthenticationProfilesMethodRadiusFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.AuthenticationProfilesMethodRadius{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.AuthenticationProfilesMethodRadius", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.AuthenticationProfilesMethodRadius{}.AttrType(), data)
}

// --- Unpacker for AuthenticationProfilesMethodSamlIdp ---
func unpackAuthenticationProfilesMethodSamlIdpToSdk(ctx context.Context, obj types.Object) (*identity_services.AuthenticationProfilesMethodSamlIdp, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.AuthenticationProfilesMethodSamlIdp", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.AuthenticationProfilesMethodSamlIdp
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk identity_services.AuthenticationProfilesMethodSamlIdp
	var d diag.Diagnostics
	// Handling Primitives
	if !model.AttributeNameUsergroup.IsNull() && !model.AttributeNameUsergroup.IsUnknown() {
		sdk.AttributeNameUsergroup = model.AttributeNameUsergroup.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "AttributeNameUsergroup", "value": *sdk.AttributeNameUsergroup})
	}

	// Handling Primitives
	if !model.AttributeNameUsername.IsNull() && !model.AttributeNameUsername.IsUnknown() {
		sdk.AttributeNameUsername = model.AttributeNameUsername.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "AttributeNameUsername", "value": *sdk.AttributeNameUsername})
	}

	// Handling Primitives
	if !model.CertificateProfile.IsNull() && !model.CertificateProfile.IsUnknown() {
		sdk.CertificateProfile = model.CertificateProfile.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "CertificateProfile", "value": *sdk.CertificateProfile})
	}

	// Handling Primitives
	if !model.EnableSingleLogout.IsNull() && !model.EnableSingleLogout.IsUnknown() {
		sdk.EnableSingleLogout = model.EnableSingleLogout.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "EnableSingleLogout", "value": *sdk.EnableSingleLogout})
	}

	// Handling Primitives
	if !model.RequestSigningCertificate.IsNull() && !model.RequestSigningCertificate.IsUnknown() {
		sdk.RequestSigningCertificate = model.RequestSigningCertificate.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "RequestSigningCertificate", "value": *sdk.RequestSigningCertificate})
	}

	// Handling Primitives
	if !model.ServerProfile.IsNull() && !model.ServerProfile.IsUnknown() {
		sdk.ServerProfile = model.ServerProfile.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "ServerProfile", "value": *sdk.ServerProfile})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.AuthenticationProfilesMethodSamlIdp", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for AuthenticationProfilesMethodSamlIdp ---
func packAuthenticationProfilesMethodSamlIdpFromSdk(ctx context.Context, sdk identity_services.AuthenticationProfilesMethodSamlIdp) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.AuthenticationProfilesMethodSamlIdp", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.AuthenticationProfilesMethodSamlIdp
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.AttributeNameUsergroup != nil {
		model.AttributeNameUsergroup = basetypes.NewStringValue(*sdk.AttributeNameUsergroup)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "AttributeNameUsergroup", "value": *sdk.AttributeNameUsergroup})
	} else {
		model.AttributeNameUsergroup = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.AttributeNameUsername != nil {
		model.AttributeNameUsername = basetypes.NewStringValue(*sdk.AttributeNameUsername)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "AttributeNameUsername", "value": *sdk.AttributeNameUsername})
	} else {
		model.AttributeNameUsername = basetypes.NewStringNull()
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
	if sdk.EnableSingleLogout != nil {
		model.EnableSingleLogout = basetypes.NewBoolValue(*sdk.EnableSingleLogout)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "EnableSingleLogout", "value": *sdk.EnableSingleLogout})
	} else {
		model.EnableSingleLogout = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.RequestSigningCertificate != nil {
		model.RequestSigningCertificate = basetypes.NewStringValue(*sdk.RequestSigningCertificate)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "RequestSigningCertificate", "value": *sdk.RequestSigningCertificate})
	} else {
		model.RequestSigningCertificate = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.ServerProfile != nil {
		model.ServerProfile = basetypes.NewStringValue(*sdk.ServerProfile)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "ServerProfile", "value": *sdk.ServerProfile})
	} else {
		model.ServerProfile = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.AuthenticationProfilesMethodSamlIdp{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.AuthenticationProfilesMethodSamlIdp", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for AuthenticationProfilesMethodSamlIdp ---
func unpackAuthenticationProfilesMethodSamlIdpListToSdk(ctx context.Context, list types.List) ([]identity_services.AuthenticationProfilesMethodSamlIdp, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.AuthenticationProfilesMethodSamlIdp")
	diags := diag.Diagnostics{}
	var data []models.AuthenticationProfilesMethodSamlIdp
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]identity_services.AuthenticationProfilesMethodSamlIdp, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.AuthenticationProfilesMethodSamlIdp{}.AttrTypes(), &item)
		unpacked, d := unpackAuthenticationProfilesMethodSamlIdpToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.AuthenticationProfilesMethodSamlIdp", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for AuthenticationProfilesMethodSamlIdp ---
func packAuthenticationProfilesMethodSamlIdpListFromSdk(ctx context.Context, sdks []identity_services.AuthenticationProfilesMethodSamlIdp) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.AuthenticationProfilesMethodSamlIdp")
	diags := diag.Diagnostics{}
	var data []models.AuthenticationProfilesMethodSamlIdp

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.AuthenticationProfilesMethodSamlIdp
		obj, d := packAuthenticationProfilesMethodSamlIdpFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.AuthenticationProfilesMethodSamlIdp{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.AuthenticationProfilesMethodSamlIdp", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.AuthenticationProfilesMethodSamlIdp{}.AttrType(), data)
}

// --- Unpacker for AuthenticationProfilesMultiFactorAuth ---
func unpackAuthenticationProfilesMultiFactorAuthToSdk(ctx context.Context, obj types.Object) (*identity_services.AuthenticationProfilesMultiFactorAuth, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.AuthenticationProfilesMultiFactorAuth", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.AuthenticationProfilesMultiFactorAuth
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk identity_services.AuthenticationProfilesMultiFactorAuth
	var d diag.Diagnostics
	// Handling Lists
	if !model.Factors.IsNull() && !model.Factors.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Factors")
		diags.Append(model.Factors.ElementsAs(ctx, &sdk.Factors, false)...)
	}

	// Handling Primitives
	if !model.MfaEnable.IsNull() && !model.MfaEnable.IsUnknown() {
		sdk.MfaEnable = model.MfaEnable.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "MfaEnable", "value": *sdk.MfaEnable})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.AuthenticationProfilesMultiFactorAuth", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for AuthenticationProfilesMultiFactorAuth ---
func packAuthenticationProfilesMultiFactorAuthFromSdk(ctx context.Context, sdk identity_services.AuthenticationProfilesMultiFactorAuth) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.AuthenticationProfilesMultiFactorAuth", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.AuthenticationProfilesMultiFactorAuth
	var d diag.Diagnostics
	// Handling Lists
	if sdk.Factors != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Factors")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Factors, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Factors)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Factors = basetypes.NewListNull(elemType)
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.MfaEnable != nil {
		model.MfaEnable = basetypes.NewBoolValue(*sdk.MfaEnable)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "MfaEnable", "value": *sdk.MfaEnable})
	} else {
		model.MfaEnable = basetypes.NewBoolNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.AuthenticationProfilesMultiFactorAuth{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.AuthenticationProfilesMultiFactorAuth", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for AuthenticationProfilesMultiFactorAuth ---
func unpackAuthenticationProfilesMultiFactorAuthListToSdk(ctx context.Context, list types.List) ([]identity_services.AuthenticationProfilesMultiFactorAuth, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.AuthenticationProfilesMultiFactorAuth")
	diags := diag.Diagnostics{}
	var data []models.AuthenticationProfilesMultiFactorAuth
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]identity_services.AuthenticationProfilesMultiFactorAuth, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.AuthenticationProfilesMultiFactorAuth{}.AttrTypes(), &item)
		unpacked, d := unpackAuthenticationProfilesMultiFactorAuthToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.AuthenticationProfilesMultiFactorAuth", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for AuthenticationProfilesMultiFactorAuth ---
func packAuthenticationProfilesMultiFactorAuthListFromSdk(ctx context.Context, sdks []identity_services.AuthenticationProfilesMultiFactorAuth) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.AuthenticationProfilesMultiFactorAuth")
	diags := diag.Diagnostics{}
	var data []models.AuthenticationProfilesMultiFactorAuth

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.AuthenticationProfilesMultiFactorAuth
		obj, d := packAuthenticationProfilesMultiFactorAuthFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.AuthenticationProfilesMultiFactorAuth{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.AuthenticationProfilesMultiFactorAuth", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.AuthenticationProfilesMultiFactorAuth{}.AttrType(), data)
}

// --- Unpacker for AuthenticationProfilesSingleSignOn ---
func unpackAuthenticationProfilesSingleSignOnToSdk(ctx context.Context, obj types.Object) (*identity_services.AuthenticationProfilesSingleSignOn, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.AuthenticationProfilesSingleSignOn", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.AuthenticationProfilesSingleSignOn
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk identity_services.AuthenticationProfilesSingleSignOn
	var d diag.Diagnostics
	// Handling Primitives
	if !model.KerberosKeytab.IsNull() && !model.KerberosKeytab.IsUnknown() {
		sdk.KerberosKeytab = model.KerberosKeytab.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "KerberosKeytab", "value": *sdk.KerberosKeytab})
	}

	// Handling Primitives
	if !model.Realm.IsNull() && !model.Realm.IsUnknown() {
		sdk.Realm = model.Realm.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Realm", "value": *sdk.Realm})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.AuthenticationProfilesSingleSignOn", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for AuthenticationProfilesSingleSignOn ---
func packAuthenticationProfilesSingleSignOnFromSdk(ctx context.Context, sdk identity_services.AuthenticationProfilesSingleSignOn) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.AuthenticationProfilesSingleSignOn", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.AuthenticationProfilesSingleSignOn
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.KerberosKeytab != nil {
		model.KerberosKeytab = basetypes.NewStringValue(*sdk.KerberosKeytab)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "KerberosKeytab", "value": *sdk.KerberosKeytab})
	} else {
		model.KerberosKeytab = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Realm != nil {
		model.Realm = basetypes.NewStringValue(*sdk.Realm)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Realm", "value": *sdk.Realm})
	} else {
		model.Realm = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.AuthenticationProfilesSingleSignOn{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.AuthenticationProfilesSingleSignOn", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for AuthenticationProfilesSingleSignOn ---
func unpackAuthenticationProfilesSingleSignOnListToSdk(ctx context.Context, list types.List) ([]identity_services.AuthenticationProfilesSingleSignOn, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.AuthenticationProfilesSingleSignOn")
	diags := diag.Diagnostics{}
	var data []models.AuthenticationProfilesSingleSignOn
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]identity_services.AuthenticationProfilesSingleSignOn, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.AuthenticationProfilesSingleSignOn{}.AttrTypes(), &item)
		unpacked, d := unpackAuthenticationProfilesSingleSignOnToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.AuthenticationProfilesSingleSignOn", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for AuthenticationProfilesSingleSignOn ---
func packAuthenticationProfilesSingleSignOnListFromSdk(ctx context.Context, sdks []identity_services.AuthenticationProfilesSingleSignOn) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.AuthenticationProfilesSingleSignOn")
	diags := diag.Diagnostics{}
	var data []models.AuthenticationProfilesSingleSignOn

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.AuthenticationProfilesSingleSignOn
		obj, d := packAuthenticationProfilesSingleSignOnFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.AuthenticationProfilesSingleSignOn{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.AuthenticationProfilesSingleSignOn", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.AuthenticationProfilesSingleSignOn{}.AttrType(), data)
}
