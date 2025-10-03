package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/network_services"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/network_services"
)

// --- Unpacker for IkeCryptoProfiles ---
func unpackIkeCryptoProfilesToSdk(ctx context.Context, obj types.Object) (*network_services.IkeCryptoProfiles, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.IkeCryptoProfiles", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.IkeCryptoProfiles
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.IkeCryptoProfiles
	var d diag.Diagnostics
	// Handling Primitives
	if !model.AuthenticationMultiple.IsNull() && !model.AuthenticationMultiple.IsUnknown() {
		val := int32(model.AuthenticationMultiple.ValueInt64())
		sdk.AuthenticationMultiple = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "AuthenticationMultiple", "value": *sdk.AuthenticationMultiple})
	}

	// Handling Primitives
	if !model.Device.IsNull() && !model.Device.IsUnknown() {
		sdk.Device = model.Device.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Device", "value": *sdk.Device})
	}

	// Handling Lists
	if !model.DhGroup.IsNull() && !model.DhGroup.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field DhGroup")
		diags.Append(model.DhGroup.ElementsAs(ctx, &sdk.DhGroup, false)...)
	}

	// Handling Lists
	if !model.Encryption.IsNull() && !model.Encryption.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Encryption")
		diags.Append(model.Encryption.ElementsAs(ctx, &sdk.Encryption, false)...)
	}

	// Handling Primitives
	if !model.Folder.IsNull() && !model.Folder.IsUnknown() {
		sdk.Folder = model.Folder.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Folder", "value": *sdk.Folder})
	}

	// Handling Lists
	if !model.Hash.IsNull() && !model.Hash.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Hash")
		diags.Append(model.Hash.ElementsAs(ctx, &sdk.Hash, false)...)
	}

	// Handling Primitives
	if !model.Id.IsNull() && !model.Id.IsUnknown() {
		sdk.Id = model.Id.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Id", "value": *sdk.Id})
	}

	// Handling Objects
	if !model.Lifetime.IsNull() && !model.Lifetime.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Lifetime")
		unpacked, d := unpackIkeCryptoProfilesLifetimeToSdk(ctx, model.Lifetime)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Lifetime"})
		}
		if unpacked != nil {
			sdk.Lifetime = unpacked
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

	tflog.Debug(ctx, "Exiting unpack helper for models.IkeCryptoProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for IkeCryptoProfiles ---
func packIkeCryptoProfilesFromSdk(ctx context.Context, sdk network_services.IkeCryptoProfiles) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.IkeCryptoProfiles", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.IkeCryptoProfiles
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.AuthenticationMultiple != nil {
		model.AuthenticationMultiple = basetypes.NewInt64Value(int64(*sdk.AuthenticationMultiple))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "AuthenticationMultiple", "value": *sdk.AuthenticationMultiple})
	} else {
		model.AuthenticationMultiple = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Device != nil {
		model.Device = basetypes.NewStringValue(*sdk.Device)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Device", "value": *sdk.Device})
	} else {
		model.Device = basetypes.NewStringNull()
	}
	// Handling Lists
	if sdk.DhGroup != nil {
		tflog.Debug(ctx, "Packing list of primitives for field DhGroup")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.DhGroup, d = basetypes.NewListValueFrom(ctx, elemType, sdk.DhGroup)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.DhGroup = basetypes.NewListNull(elemType)
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
	// Handling Primitives
	// Standard primitive packing
	if sdk.Folder != nil {
		model.Folder = basetypes.NewStringValue(*sdk.Folder)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Folder", "value": *sdk.Folder})
	} else {
		model.Folder = basetypes.NewStringNull()
	}
	// Handling Lists
	if sdk.Hash != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Hash")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Hash, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Hash)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Hash = basetypes.NewListNull(elemType)
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
	if sdk.Lifetime != nil {
		tflog.Debug(ctx, "Packing nested object for field Lifetime")
		packed, d := packIkeCryptoProfilesLifetimeFromSdk(ctx, *sdk.Lifetime)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Lifetime"})
		}
		model.Lifetime = packed
	} else {
		model.Lifetime = basetypes.NewObjectNull(models.IkeCryptoProfilesLifetime{}.AttrTypes())
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

	obj, d := types.ObjectValueFrom(ctx, models.IkeCryptoProfiles{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.IkeCryptoProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for IkeCryptoProfiles ---
func unpackIkeCryptoProfilesListToSdk(ctx context.Context, list types.List) ([]network_services.IkeCryptoProfiles, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.IkeCryptoProfiles")
	diags := diag.Diagnostics{}
	var data []models.IkeCryptoProfiles
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.IkeCryptoProfiles, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.IkeCryptoProfiles{}.AttrTypes(), &item)
		unpacked, d := unpackIkeCryptoProfilesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.IkeCryptoProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for IkeCryptoProfiles ---
func packIkeCryptoProfilesListFromSdk(ctx context.Context, sdks []network_services.IkeCryptoProfiles) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.IkeCryptoProfiles")
	diags := diag.Diagnostics{}
	var data []models.IkeCryptoProfiles

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.IkeCryptoProfiles
		obj, d := packIkeCryptoProfilesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.IkeCryptoProfiles{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.IkeCryptoProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.IkeCryptoProfiles{}.AttrType(), data)
}

// --- Unpacker for IkeCryptoProfilesLifetime ---
func unpackIkeCryptoProfilesLifetimeToSdk(ctx context.Context, obj types.Object) (*network_services.IkeCryptoProfilesLifetime, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.IkeCryptoProfilesLifetime", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.IkeCryptoProfilesLifetime
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.IkeCryptoProfilesLifetime
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

	tflog.Debug(ctx, "Exiting unpack helper for models.IkeCryptoProfilesLifetime", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for IkeCryptoProfilesLifetime ---
func packIkeCryptoProfilesLifetimeFromSdk(ctx context.Context, sdk network_services.IkeCryptoProfilesLifetime) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.IkeCryptoProfilesLifetime", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.IkeCryptoProfilesLifetime
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

	obj, d := types.ObjectValueFrom(ctx, models.IkeCryptoProfilesLifetime{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.IkeCryptoProfilesLifetime", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for IkeCryptoProfilesLifetime ---
func unpackIkeCryptoProfilesLifetimeListToSdk(ctx context.Context, list types.List) ([]network_services.IkeCryptoProfilesLifetime, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.IkeCryptoProfilesLifetime")
	diags := diag.Diagnostics{}
	var data []models.IkeCryptoProfilesLifetime
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.IkeCryptoProfilesLifetime, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.IkeCryptoProfilesLifetime{}.AttrTypes(), &item)
		unpacked, d := unpackIkeCryptoProfilesLifetimeToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.IkeCryptoProfilesLifetime", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for IkeCryptoProfilesLifetime ---
func packIkeCryptoProfilesLifetimeListFromSdk(ctx context.Context, sdks []network_services.IkeCryptoProfilesLifetime) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.IkeCryptoProfilesLifetime")
	diags := diag.Diagnostics{}
	var data []models.IkeCryptoProfilesLifetime

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.IkeCryptoProfilesLifetime
		obj, d := packIkeCryptoProfilesLifetimeFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.IkeCryptoProfilesLifetime{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.IkeCryptoProfilesLifetime", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.IkeCryptoProfilesLifetime{}.AttrType(), data)
}
