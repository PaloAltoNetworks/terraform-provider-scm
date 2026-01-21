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

// tacacsServerProfilesSensitiveValuePatcher is an in-memory struct to temporarily store plaintext
// and encrypted values for sensitive fields during the Create/Update/Read workflows.
type tacacsServerProfilesSensitiveValuePatcher struct {
	// arrayValues stores sensitive values for fields inside arrays, with dynamic keys like "server_secret_plaintext_0"
	arrayValues map[string]basetypes.StringValue
}

// Initialize initializes the patcher (must be called before use if there are array fields)
func (p *tacacsServerProfilesSensitiveValuePatcher) init() {
	if p.arrayValues == nil {
		p.arrayValues = make(map[string]basetypes.StringValue)
	}
}

// populatePatcherFromState populates the patcher struct from the resource's state.
func (p *tacacsServerProfilesSensitiveValuePatcher) populatePatcherFromState(ctx context.Context, state models.TacacsServerProfiles) diag.Diagnostics {
	var diags diag.Diagnostics
	p.init() // Initialize the map
	if state.EncryptedValues.IsNull() || state.EncryptedValues.IsUnknown() {
		return diags
	}

	var ev map[string]string
	diags.Append(state.EncryptedValues.ElementsAs(ctx, &ev, false)...)
	if diags.HasError() {
		return diags
	}
	// Load array values (all keys that don't match the non-array patterns)
	for key, val := range ev {
		// Store any remaining keys (array values) in the map
		p.arrayValues[key] = basetypes.NewStringValue(val)
	}

	return diags
}

// populateEncryptedValuesMap returns a map of the patcher's values for saving to state.
func (p *tacacsServerProfilesSensitiveValuePatcher) populateEncryptedValuesMap() map[string]string {
	ev := make(map[string]string)
	// Add all array values
	for key, val := range p.arrayValues {
		if !val.IsNull() {
			ev[key] = val.ValueString()
		}
	}
	return ev
}

// --- Unpacker for TacacsServerProfiles ---
func unpackTacacsServerProfilesToSdk(ctx context.Context, obj types.Object) (*identity_services.TacacsServerProfiles, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.TacacsServerProfiles", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.TacacsServerProfiles
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk identity_services.TacacsServerProfiles
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

	// Handling Primitives
	if !model.Protocol.IsNull() && !model.Protocol.IsUnknown() {
		sdk.Protocol = model.Protocol.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Protocol", "value": sdk.Protocol})
	}

	// Handling Lists
	if !model.Server.IsNull() && !model.Server.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field Server")
		unpacked, d := unpackTacacsServerProfilesServerInnerListToSdk(ctx, model.Server)
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

	// Handling Primitives
	if !model.UseSingleConnection.IsNull() && !model.UseSingleConnection.IsUnknown() {
		sdk.UseSingleConnection = model.UseSingleConnection.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "UseSingleConnection", "value": *sdk.UseSingleConnection})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.TacacsServerProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for TacacsServerProfiles ---
func packTacacsServerProfilesFromSdk(ctx context.Context, sdk identity_services.TacacsServerProfiles) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.TacacsServerProfiles", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.TacacsServerProfiles
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
	model.Id = basetypes.NewStringValue(sdk.Id)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Id", "value": sdk.Id})
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	// Handling Primitives
	// Standard primitive packing
	model.Protocol = basetypes.NewStringValue(sdk.Protocol)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Protocol", "value": sdk.Protocol})
	// Handling Lists
	if sdk.Server != nil {
		tflog.Debug(ctx, "Packing list of objects for field Server")
		packed, d := packTacacsServerProfilesServerInnerListFromSdk(ctx, sdk.Server)
		diags.Append(d...)
		model.Server = packed
	} else {
		model.Server = basetypes.NewListNull(models.TacacsServerProfilesServerInner{}.AttrType())
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
	// Handling Primitives
	// Standard primitive packing
	if sdk.UseSingleConnection != nil {
		model.UseSingleConnection = basetypes.NewBoolValue(*sdk.UseSingleConnection)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "UseSingleConnection", "value": *sdk.UseSingleConnection})
	} else {
		model.UseSingleConnection = basetypes.NewBoolNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.TacacsServerProfiles{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.TacacsServerProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for TacacsServerProfiles ---
func unpackTacacsServerProfilesListToSdk(ctx context.Context, list types.List) ([]identity_services.TacacsServerProfiles, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.TacacsServerProfiles")
	diags := diag.Diagnostics{}
	var data []models.TacacsServerProfiles
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]identity_services.TacacsServerProfiles, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.TacacsServerProfiles{}.AttrTypes(), &item)
		unpacked, d := unpackTacacsServerProfilesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.TacacsServerProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for TacacsServerProfiles ---
func packTacacsServerProfilesListFromSdk(ctx context.Context, sdks []identity_services.TacacsServerProfiles) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.TacacsServerProfiles")
	diags := diag.Diagnostics{}
	var data []models.TacacsServerProfiles

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.TacacsServerProfiles
		obj, d := packTacacsServerProfilesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.TacacsServerProfiles{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.TacacsServerProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.TacacsServerProfiles{}.AttrType(), data)
}

// --- Unpacker for TacacsServerProfilesServerInner ---
func unpackTacacsServerProfilesServerInnerToSdk(ctx context.Context, obj types.Object) (*identity_services.TacacsServerProfilesServerInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.TacacsServerProfilesServerInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.TacacsServerProfilesServerInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk identity_services.TacacsServerProfilesServerInner
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

	// Handling Primitives
	if !model.Secret.IsNull() && !model.Secret.IsUnknown() {
		sdk.Secret = model.Secret.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Secret", "value": *sdk.Secret})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.TacacsServerProfilesServerInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for TacacsServerProfilesServerInner ---
func packTacacsServerProfilesServerInnerFromSdk(ctx context.Context, sdk identity_services.TacacsServerProfilesServerInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.TacacsServerProfilesServerInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.TacacsServerProfilesServerInner
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
	// Handling Primitives
	// Standard primitive packing
	if sdk.Secret != nil {
		model.Secret = basetypes.NewStringValue(*sdk.Secret)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Secret", "value": *sdk.Secret})
	} else {
		model.Secret = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.TacacsServerProfilesServerInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.TacacsServerProfilesServerInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for TacacsServerProfilesServerInner ---
func unpackTacacsServerProfilesServerInnerListToSdk(ctx context.Context, list types.List) ([]identity_services.TacacsServerProfilesServerInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.TacacsServerProfilesServerInner")
	diags := diag.Diagnostics{}
	var data []models.TacacsServerProfilesServerInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]identity_services.TacacsServerProfilesServerInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.TacacsServerProfilesServerInner{}.AttrTypes(), &item)
		unpacked, d := unpackTacacsServerProfilesServerInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.TacacsServerProfilesServerInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for TacacsServerProfilesServerInner ---
func packTacacsServerProfilesServerInnerListFromSdk(ctx context.Context, sdks []identity_services.TacacsServerProfilesServerInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.TacacsServerProfilesServerInner")
	diags := diag.Diagnostics{}
	var data []models.TacacsServerProfilesServerInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.TacacsServerProfilesServerInner
		obj, d := packTacacsServerProfilesServerInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.TacacsServerProfilesServerInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.TacacsServerProfilesServerInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.TacacsServerProfilesServerInner{}.AttrType(), data)
}
