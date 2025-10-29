package provider

import (
	"context"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/network_services"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/network_services"
)

// --- Unpacker for IpsecCryptoProfiles ---
func unpackIpsecCryptoProfilesToSdk(ctx context.Context, obj types.Object) (*network_services.IpsecCryptoProfiles, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.IpsecCryptoProfiles", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.IpsecCryptoProfiles
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.IpsecCryptoProfiles
	var d diag.Diagnostics

	// Handling Objects
	if !model.Ah.IsNull() && !model.Ah.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Ah")
		unpacked, d := unpackIpsecCryptoProfilesAhToSdk(ctx, model.Ah)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Ah"})
		}
		if unpacked != nil {
			sdk.Ah = unpacked
		}
	}

	// Handling Primitives
	if !model.Device.IsNull() && !model.Device.IsUnknown() {
		sdk.Device = model.Device.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Device", "value": *sdk.Device})
	}

	// Handling Primitives
	if !model.DhGroup.IsNull() && !model.DhGroup.IsUnknown() {
		sdk.DhGroup = model.DhGroup.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DhGroup", "value": *sdk.DhGroup})
	}

	// Handling Objects
	if !model.Esp.IsNull() && !model.Esp.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Esp")
		unpacked, d := unpackIpsecCryptoProfilesEspToSdk(ctx, model.Esp)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Esp"})
		}
		if unpacked != nil {
			sdk.Esp = unpacked
		}
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

	// Handling Objects
	if !model.Lifesize.IsNull() && !model.Lifesize.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Lifesize")
		unpacked, d := unpackIpsecCryptoProfilesLifesizeToSdk(ctx, model.Lifesize)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Lifesize"})
		}
		if unpacked != nil {
			sdk.Lifesize = unpacked
		}
	}

	// Handling Objects
	if !model.Lifetime.IsNull() && !model.Lifetime.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Lifetime")
		unpacked, d := unpackIpsecCryptoProfilesLifetimeToSdk(ctx, model.Lifetime)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Lifetime"})
		}
		if unpacked != nil {
			sdk.Lifetime = *unpacked
		}
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

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.IpsecCryptoProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for IpsecCryptoProfiles ---
func packIpsecCryptoProfilesFromSdk(ctx context.Context, sdk network_services.IpsecCryptoProfiles) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.IpsecCryptoProfiles", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.IpsecCryptoProfiles
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Ah != nil {
		tflog.Debug(ctx, "Packing nested object for field Ah")
		packed, d := packIpsecCryptoProfilesAhFromSdk(ctx, *sdk.Ah)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Ah"})
		}
		model.Ah = packed
	} else {
		model.Ah = basetypes.NewObjectNull(models.IpsecCryptoProfilesAh{}.AttrTypes())
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
	if sdk.DhGroup != nil {
		model.DhGroup = basetypes.NewStringValue(*sdk.DhGroup)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DhGroup", "value": *sdk.DhGroup})
	} else {
		model.DhGroup = basetypes.NewStringNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Esp != nil {
		tflog.Debug(ctx, "Packing nested object for field Esp")
		packed, d := packIpsecCryptoProfilesEspFromSdk(ctx, *sdk.Esp)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Esp"})
		}
		model.Esp = packed
	} else {
		model.Esp = basetypes.NewObjectNull(models.IpsecCryptoProfilesEsp{}.AttrTypes())
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
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Lifesize != nil {
		tflog.Debug(ctx, "Packing nested object for field Lifesize")
		packed, d := packIpsecCryptoProfilesLifesizeFromSdk(ctx, *sdk.Lifesize)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Lifesize"})
		}
		model.Lifesize = packed
	} else {
		model.Lifesize = basetypes.NewObjectNull(models.IpsecCryptoProfilesLifesize{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	// Logic for non-pointer / value-type nested objects
	if !reflect.ValueOf(sdk.Lifetime).IsZero() {
		tflog.Debug(ctx, "Packing nested object for field Lifetime")
		packed, d := packIpsecCryptoProfilesLifetimeFromSdk(ctx, sdk.Lifetime)
		diags.Append(d...)
		model.Lifetime = packed
	} else {
		model.Lifetime = basetypes.NewObjectNull(models.IpsecCryptoProfilesLifetime{}.AttrTypes())
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
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.IpsecCryptoProfiles{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.IpsecCryptoProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for IpsecCryptoProfiles ---
func unpackIpsecCryptoProfilesListToSdk(ctx context.Context, list types.List) ([]network_services.IpsecCryptoProfiles, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.IpsecCryptoProfiles")
	diags := diag.Diagnostics{}
	var data []models.IpsecCryptoProfiles
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.IpsecCryptoProfiles, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.IpsecCryptoProfiles{}.AttrTypes(), &item)
		unpacked, d := unpackIpsecCryptoProfilesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.IpsecCryptoProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for IpsecCryptoProfiles ---
func packIpsecCryptoProfilesListFromSdk(ctx context.Context, sdks []network_services.IpsecCryptoProfiles) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.IpsecCryptoProfiles")
	diags := diag.Diagnostics{}
	var data []models.IpsecCryptoProfiles

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.IpsecCryptoProfiles
		obj, d := packIpsecCryptoProfilesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.IpsecCryptoProfiles{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.IpsecCryptoProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.IpsecCryptoProfiles{}.AttrType(), data)
}

// --- Unpacker for IpsecCryptoProfilesAh ---
func unpackIpsecCryptoProfilesAhToSdk(ctx context.Context, obj types.Object) (*network_services.IpsecCryptoProfilesAh, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.IpsecCryptoProfilesAh", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.IpsecCryptoProfilesAh
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.IpsecCryptoProfilesAh
	var d diag.Diagnostics
	// Handling Lists
	if !model.Authentication.IsNull() && !model.Authentication.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Authentication")
		diags.Append(model.Authentication.ElementsAs(ctx, &sdk.Authentication, false)...)
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.IpsecCryptoProfilesAh", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for IpsecCryptoProfilesAh ---
func packIpsecCryptoProfilesAhFromSdk(ctx context.Context, sdk network_services.IpsecCryptoProfilesAh) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.IpsecCryptoProfilesAh", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.IpsecCryptoProfilesAh
	var d diag.Diagnostics
	// Handling Lists
	if sdk.Authentication != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Authentication")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Authentication, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Authentication)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Authentication = basetypes.NewListNull(elemType)
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.IpsecCryptoProfilesAh{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.IpsecCryptoProfilesAh", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for IpsecCryptoProfilesAh ---
func unpackIpsecCryptoProfilesAhListToSdk(ctx context.Context, list types.List) ([]network_services.IpsecCryptoProfilesAh, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.IpsecCryptoProfilesAh")
	diags := diag.Diagnostics{}
	var data []models.IpsecCryptoProfilesAh
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.IpsecCryptoProfilesAh, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.IpsecCryptoProfilesAh{}.AttrTypes(), &item)
		unpacked, d := unpackIpsecCryptoProfilesAhToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.IpsecCryptoProfilesAh", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for IpsecCryptoProfilesAh ---
func packIpsecCryptoProfilesAhListFromSdk(ctx context.Context, sdks []network_services.IpsecCryptoProfilesAh) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.IpsecCryptoProfilesAh")
	diags := diag.Diagnostics{}
	var data []models.IpsecCryptoProfilesAh

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.IpsecCryptoProfilesAh
		obj, d := packIpsecCryptoProfilesAhFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.IpsecCryptoProfilesAh{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.IpsecCryptoProfilesAh", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.IpsecCryptoProfilesAh{}.AttrType(), data)
}

// --- Unpacker for IpsecCryptoProfilesEsp ---
func unpackIpsecCryptoProfilesEspToSdk(ctx context.Context, obj types.Object) (*network_services.IpsecCryptoProfilesEsp, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.IpsecCryptoProfilesEsp", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.IpsecCryptoProfilesEsp
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.IpsecCryptoProfilesEsp
	var d diag.Diagnostics
	// Handling Lists
	if !model.Authentication.IsNull() && !model.Authentication.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Authentication")
		diags.Append(model.Authentication.ElementsAs(ctx, &sdk.Authentication, false)...)
	}

	// Handling Lists
	if !model.Encryption.IsNull() && !model.Encryption.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Encryption")
		diags.Append(model.Encryption.ElementsAs(ctx, &sdk.Encryption, false)...)
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.IpsecCryptoProfilesEsp", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for IpsecCryptoProfilesEsp ---
func packIpsecCryptoProfilesEspFromSdk(ctx context.Context, sdk network_services.IpsecCryptoProfilesEsp) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.IpsecCryptoProfilesEsp", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.IpsecCryptoProfilesEsp
	var d diag.Diagnostics
	// Handling Lists
	if sdk.Authentication != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Authentication")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Authentication, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Authentication)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Authentication = basetypes.NewListNull(elemType)
	}
	// Handling Lists
	if sdk.Encryption != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Encryption")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Encryption, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Encryption)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Encryption = basetypes.NewListNull(elemType)
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.IpsecCryptoProfilesEsp{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.IpsecCryptoProfilesEsp", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for IpsecCryptoProfilesEsp ---
func unpackIpsecCryptoProfilesEspListToSdk(ctx context.Context, list types.List) ([]network_services.IpsecCryptoProfilesEsp, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.IpsecCryptoProfilesEsp")
	diags := diag.Diagnostics{}
	var data []models.IpsecCryptoProfilesEsp
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.IpsecCryptoProfilesEsp, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.IpsecCryptoProfilesEsp{}.AttrTypes(), &item)
		unpacked, d := unpackIpsecCryptoProfilesEspToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.IpsecCryptoProfilesEsp", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for IpsecCryptoProfilesEsp ---
func packIpsecCryptoProfilesEspListFromSdk(ctx context.Context, sdks []network_services.IpsecCryptoProfilesEsp) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.IpsecCryptoProfilesEsp")
	diags := diag.Diagnostics{}
	var data []models.IpsecCryptoProfilesEsp

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.IpsecCryptoProfilesEsp
		obj, d := packIpsecCryptoProfilesEspFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.IpsecCryptoProfilesEsp{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.IpsecCryptoProfilesEsp", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.IpsecCryptoProfilesEsp{}.AttrType(), data)
}

// --- Unpacker for IpsecCryptoProfilesLifesize ---
func unpackIpsecCryptoProfilesLifesizeToSdk(ctx context.Context, obj types.Object) (*network_services.IpsecCryptoProfilesLifesize, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.IpsecCryptoProfilesLifesize", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.IpsecCryptoProfilesLifesize
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.IpsecCryptoProfilesLifesize
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Gb.IsNull() && !model.Gb.IsUnknown() {
		val := int32(model.Gb.ValueInt64())
		sdk.Gb = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Gb", "value": *sdk.Gb})
	}

	// Handling Primitives
	if !model.Kb.IsNull() && !model.Kb.IsUnknown() {
		val := int32(model.Kb.ValueInt64())
		sdk.Kb = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Kb", "value": *sdk.Kb})
	}

	// Handling Primitives
	if !model.Mb.IsNull() && !model.Mb.IsUnknown() {
		val := int32(model.Mb.ValueInt64())
		sdk.Mb = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Mb", "value": *sdk.Mb})
	}

	// Handling Primitives
	if !model.Tb.IsNull() && !model.Tb.IsUnknown() {
		val := int32(model.Tb.ValueInt64())
		sdk.Tb = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Tb", "value": *sdk.Tb})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.IpsecCryptoProfilesLifesize", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for IpsecCryptoProfilesLifesize ---
func packIpsecCryptoProfilesLifesizeFromSdk(ctx context.Context, sdk network_services.IpsecCryptoProfilesLifesize) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.IpsecCryptoProfilesLifesize", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.IpsecCryptoProfilesLifesize
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Gb != nil {
		model.Gb = basetypes.NewInt64Value(int64(*sdk.Gb))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Gb", "value": *sdk.Gb})
	} else {
		model.Gb = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Kb != nil {
		model.Kb = basetypes.NewInt64Value(int64(*sdk.Kb))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Kb", "value": *sdk.Kb})
	} else {
		model.Kb = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Mb != nil {
		model.Mb = basetypes.NewInt64Value(int64(*sdk.Mb))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Mb", "value": *sdk.Mb})
	} else {
		model.Mb = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Tb != nil {
		model.Tb = basetypes.NewInt64Value(int64(*sdk.Tb))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Tb", "value": *sdk.Tb})
	} else {
		model.Tb = basetypes.NewInt64Null()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.IpsecCryptoProfilesLifesize{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.IpsecCryptoProfilesLifesize", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for IpsecCryptoProfilesLifesize ---
func unpackIpsecCryptoProfilesLifesizeListToSdk(ctx context.Context, list types.List) ([]network_services.IpsecCryptoProfilesLifesize, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.IpsecCryptoProfilesLifesize")
	diags := diag.Diagnostics{}
	var data []models.IpsecCryptoProfilesLifesize
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.IpsecCryptoProfilesLifesize, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.IpsecCryptoProfilesLifesize{}.AttrTypes(), &item)
		unpacked, d := unpackIpsecCryptoProfilesLifesizeToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.IpsecCryptoProfilesLifesize", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for IpsecCryptoProfilesLifesize ---
func packIpsecCryptoProfilesLifesizeListFromSdk(ctx context.Context, sdks []network_services.IpsecCryptoProfilesLifesize) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.IpsecCryptoProfilesLifesize")
	diags := diag.Diagnostics{}
	var data []models.IpsecCryptoProfilesLifesize

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.IpsecCryptoProfilesLifesize
		obj, d := packIpsecCryptoProfilesLifesizeFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.IpsecCryptoProfilesLifesize{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.IpsecCryptoProfilesLifesize", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.IpsecCryptoProfilesLifesize{}.AttrType(), data)
}

// --- Unpacker for IpsecCryptoProfilesLifetime ---
func unpackIpsecCryptoProfilesLifetimeToSdk(ctx context.Context, obj types.Object) (*network_services.IpsecCryptoProfilesLifetime, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.IpsecCryptoProfilesLifetime", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.IpsecCryptoProfilesLifetime
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.IpsecCryptoProfilesLifetime
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Days.IsNull() && !model.Days.IsUnknown() {
		val := int32(model.Days.ValueInt64())
		sdk.Days = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Days", "value": *sdk.Days})
	}

	// Handling Primitives
	if !model.Hours.IsNull() && !model.Hours.IsUnknown() {
		val := int32(model.Hours.ValueInt64())
		sdk.Hours = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Hours", "value": *sdk.Hours})
	}

	// Handling Primitives
	if !model.Minutes.IsNull() && !model.Minutes.IsUnknown() {
		val := int32(model.Minutes.ValueInt64())
		sdk.Minutes = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Minutes", "value": *sdk.Minutes})
	}

	// Handling Primitives
	if !model.Seconds.IsNull() && !model.Seconds.IsUnknown() {
		val := int32(model.Seconds.ValueInt64())
		sdk.Seconds = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Seconds", "value": *sdk.Seconds})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.IpsecCryptoProfilesLifetime", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for IpsecCryptoProfilesLifetime ---
func packIpsecCryptoProfilesLifetimeFromSdk(ctx context.Context, sdk network_services.IpsecCryptoProfilesLifetime) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.IpsecCryptoProfilesLifetime", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.IpsecCryptoProfilesLifetime
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Days != nil {
		model.Days = basetypes.NewInt64Value(int64(*sdk.Days))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Days", "value": *sdk.Days})
	} else {
		model.Days = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Hours != nil {
		model.Hours = basetypes.NewInt64Value(int64(*sdk.Hours))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Hours", "value": *sdk.Hours})
	} else {
		model.Hours = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Minutes != nil {
		model.Minutes = basetypes.NewInt64Value(int64(*sdk.Minutes))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Minutes", "value": *sdk.Minutes})
	} else {
		model.Minutes = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Seconds != nil {
		model.Seconds = basetypes.NewInt64Value(int64(*sdk.Seconds))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Seconds", "value": *sdk.Seconds})
	} else {
		model.Seconds = basetypes.NewInt64Null()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.IpsecCryptoProfilesLifetime{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.IpsecCryptoProfilesLifetime", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for IpsecCryptoProfilesLifetime ---
func unpackIpsecCryptoProfilesLifetimeListToSdk(ctx context.Context, list types.List) ([]network_services.IpsecCryptoProfilesLifetime, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.IpsecCryptoProfilesLifetime")
	diags := diag.Diagnostics{}
	var data []models.IpsecCryptoProfilesLifetime
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.IpsecCryptoProfilesLifetime, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.IpsecCryptoProfilesLifetime{}.AttrTypes(), &item)
		unpacked, d := unpackIpsecCryptoProfilesLifetimeToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.IpsecCryptoProfilesLifetime", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for IpsecCryptoProfilesLifetime ---
func packIpsecCryptoProfilesLifetimeListFromSdk(ctx context.Context, sdks []network_services.IpsecCryptoProfilesLifetime) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.IpsecCryptoProfilesLifetime")
	diags := diag.Diagnostics{}
	var data []models.IpsecCryptoProfilesLifetime

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.IpsecCryptoProfilesLifetime
		obj, d := packIpsecCryptoProfilesLifetimeFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.IpsecCryptoProfilesLifetime{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.IpsecCryptoProfilesLifetime", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.IpsecCryptoProfilesLifetime{}.AttrType(), data)
}
