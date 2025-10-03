package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/network_services"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/network_services"
)

// ospfAuthProfilesSensitiveValuePatcher is an in-memory struct to temporarily store plaintext
// and encrypted values for sensitive fields during the Create/Update/Read workflows.
type ospfAuthProfilesSensitiveValuePatcher struct {
	password_plaintext basetypes.StringValue
	password_encrypted basetypes.StringValue
}

// populatePatcherFromState populates the patcher struct from the resource's state.
func (p *ospfAuthProfilesSensitiveValuePatcher) populatePatcherFromState(ctx context.Context, state models.OspfAuthProfiles) diag.Diagnostics {
	var diags diag.Diagnostics
	if state.EncryptedValues.IsNull() || state.EncryptedValues.IsUnknown() {
		return diags
	}

	var ev map[string]string
	diags.Append(state.EncryptedValues.ElementsAs(ctx, &ev, false)...)
	if diags.HasError() {
		return diags
	}
	if val, ok := ev["password_plaintext"]; ok {
		p.password_plaintext = basetypes.NewStringValue(val)
	}
	if val, ok := ev["password_encrypted"]; ok {
		p.password_encrypted = basetypes.NewStringValue(val)
	}

	return diags
}

// populateEncryptedValuesMap returns a map of the patcher's values for saving to state.
func (p *ospfAuthProfilesSensitiveValuePatcher) populateEncryptedValuesMap() map[string]string {
	ev := make(map[string]string)
	if !p.password_plaintext.IsNull() {
		ev["password_plaintext"] = p.password_plaintext.ValueString()
	}
	if !p.password_encrypted.IsNull() {
		ev["password_encrypted"] = p.password_encrypted.ValueString()
	}
	return ev
}

// --- Unpacker for OspfAuthProfiles ---
func unpackOspfAuthProfilesToSdk(ctx context.Context, obj types.Object) (*network_services.OspfAuthProfiles, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.OspfAuthProfiles", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.OspfAuthProfiles
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.OspfAuthProfiles
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

	// Handling Lists
	if !model.Md5.IsNull() && !model.Md5.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field Md5")
		unpacked, d := unpackOspfAuthProfilesMd5InnerListToSdk(ctx, model.Md5)
		diags.Append(d...)
		sdk.Md5 = unpacked
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	// Handling Primitives
	if !model.Password.IsNull() && !model.Password.IsUnknown() {
		sdk.Password = model.Password.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Password", "value": *sdk.Password})
	}

	// Handling Primitives
	if !model.Snippet.IsNull() && !model.Snippet.IsUnknown() {
		sdk.Snippet = model.Snippet.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.OspfAuthProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for OspfAuthProfiles ---
func packOspfAuthProfilesFromSdk(ctx context.Context, sdk network_services.OspfAuthProfiles) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.OspfAuthProfiles", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.OspfAuthProfiles
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
	// Handling Lists
	if sdk.Md5 != nil {
		tflog.Debug(ctx, "Packing list of objects for field Md5")
		packed, d := packOspfAuthProfilesMd5InnerListFromSdk(ctx, sdk.Md5)
		diags.Append(d...)
		model.Md5 = packed
	} else {
		model.Md5 = basetypes.NewListNull(models.OspfAuthProfilesMd5Inner{}.AttrType())
	}
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	// Handling Primitives
	// Standard primitive packing
	if sdk.Password != nil {
		model.Password = basetypes.NewStringValue(*sdk.Password)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Password", "value": *sdk.Password})
	} else {
		model.Password = basetypes.NewStringNull()
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

	obj, d := types.ObjectValueFrom(ctx, models.OspfAuthProfiles{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.OspfAuthProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for OspfAuthProfiles ---
func unpackOspfAuthProfilesListToSdk(ctx context.Context, list types.List) ([]network_services.OspfAuthProfiles, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.OspfAuthProfiles")
	diags := diag.Diagnostics{}
	var data []models.OspfAuthProfiles
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.OspfAuthProfiles, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.OspfAuthProfiles{}.AttrTypes(), &item)
		unpacked, d := unpackOspfAuthProfilesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.OspfAuthProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for OspfAuthProfiles ---
func packOspfAuthProfilesListFromSdk(ctx context.Context, sdks []network_services.OspfAuthProfiles) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.OspfAuthProfiles")
	diags := diag.Diagnostics{}
	var data []models.OspfAuthProfiles

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.OspfAuthProfiles
		obj, d := packOspfAuthProfilesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.OspfAuthProfiles{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.OspfAuthProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.OspfAuthProfiles{}.AttrType(), data)
}

// --- Unpacker for OspfAuthProfilesMd5Inner ---
func unpackOspfAuthProfilesMd5InnerToSdk(ctx context.Context, obj types.Object) (*network_services.OspfAuthProfilesMd5Inner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.OspfAuthProfilesMd5Inner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.OspfAuthProfilesMd5Inner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.OspfAuthProfilesMd5Inner
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Key.IsNull() && !model.Key.IsUnknown() {
		sdk.Key = model.Key.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Key", "value": *sdk.Key})
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		val := int32(model.Name.ValueInt64())
		sdk.Name = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
	}

	// Handling Primitives
	if !model.Preferred.IsNull() && !model.Preferred.IsUnknown() {
		sdk.Preferred = model.Preferred.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Preferred", "value": *sdk.Preferred})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.OspfAuthProfilesMd5Inner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for OspfAuthProfilesMd5Inner ---
func packOspfAuthProfilesMd5InnerFromSdk(ctx context.Context, sdk network_services.OspfAuthProfilesMd5Inner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.OspfAuthProfilesMd5Inner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.OspfAuthProfilesMd5Inner
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Key != nil {
		model.Key = basetypes.NewStringValue(*sdk.Key)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Key", "value": *sdk.Key})
	} else {
		model.Key = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Name != nil {
		model.Name = basetypes.NewInt64Value(int64(*sdk.Name))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
	} else {
		model.Name = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Preferred != nil {
		model.Preferred = basetypes.NewBoolValue(*sdk.Preferred)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Preferred", "value": *sdk.Preferred})
	} else {
		model.Preferred = basetypes.NewBoolNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.OspfAuthProfilesMd5Inner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.OspfAuthProfilesMd5Inner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for OspfAuthProfilesMd5Inner ---
func unpackOspfAuthProfilesMd5InnerListToSdk(ctx context.Context, list types.List) ([]network_services.OspfAuthProfilesMd5Inner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.OspfAuthProfilesMd5Inner")
	diags := diag.Diagnostics{}
	var data []models.OspfAuthProfilesMd5Inner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.OspfAuthProfilesMd5Inner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.OspfAuthProfilesMd5Inner{}.AttrTypes(), &item)
		unpacked, d := unpackOspfAuthProfilesMd5InnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.OspfAuthProfilesMd5Inner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for OspfAuthProfilesMd5Inner ---
func packOspfAuthProfilesMd5InnerListFromSdk(ctx context.Context, sdks []network_services.OspfAuthProfilesMd5Inner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.OspfAuthProfilesMd5Inner")
	diags := diag.Diagnostics{}
	var data []models.OspfAuthProfilesMd5Inner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.OspfAuthProfilesMd5Inner
		obj, d := packOspfAuthProfilesMd5InnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.OspfAuthProfilesMd5Inner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.OspfAuthProfilesMd5Inner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.OspfAuthProfilesMd5Inner{}.AttrType(), data)
}
