package provider

import (
	"context"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/security_services"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/security_services"
)

// --- Unpacker for UrlAccessProfiles ---
func unpackUrlAccessProfilesToSdk(ctx context.Context, obj types.Object) (*security_services.UrlAccessProfiles, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.UrlAccessProfiles", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.UrlAccessProfiles
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.UrlAccessProfiles
	var d diag.Diagnostics
	// Handling Lists
	if !model.Alert.IsNull() && !model.Alert.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Alert")
		diags.Append(model.Alert.ElementsAs(ctx, &sdk.Alert, false)...)
	}

	// Handling Lists
	if !model.Allow.IsNull() && !model.Allow.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Allow")
		diags.Append(model.Allow.ElementsAs(ctx, &sdk.Allow, false)...)
	}

	// Handling Lists
	if !model.Block.IsNull() && !model.Block.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Block")
		diags.Append(model.Block.ElementsAs(ctx, &sdk.Block, false)...)
	}

	// Handling Primitives
	if !model.CloudInlineCat.IsNull() && !model.CloudInlineCat.IsUnknown() {
		sdk.CloudInlineCat = model.CloudInlineCat.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "CloudInlineCat", "value": *sdk.CloudInlineCat})
	}

	// Handling Lists
	if !model.Continue.IsNull() && !model.Continue.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Continue")
		diags.Append(model.Continue.ElementsAs(ctx, &sdk.Continue, false)...)
	}

	// Handling Objects
	if !model.CredentialEnforcement.IsNull() && !model.CredentialEnforcement.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field CredentialEnforcement")
		unpacked, d := unpackUrlAccessProfilesCredentialEnforcementToSdk(ctx, model.CredentialEnforcement)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "CredentialEnforcement"})
		}
		if unpacked != nil {
			sdk.CredentialEnforcement = unpacked
		}
	}

	// Handling Primitives
	if !model.Description.IsNull() && !model.Description.IsUnknown() {
		sdk.Description = model.Description.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Description", "value": *sdk.Description})
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
		sdk.Id = model.Id.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Id", "value": *sdk.Id})
	}

	// Handling Primitives
	if !model.LocalInlineCat.IsNull() && !model.LocalInlineCat.IsUnknown() {
		sdk.LocalInlineCat = model.LocalInlineCat.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "LocalInlineCat", "value": *sdk.LocalInlineCat})
	}

	// Handling Primitives
	if !model.LogContainerPageOnly.IsNull() && !model.LogContainerPageOnly.IsUnknown() {
		sdk.LogContainerPageOnly = model.LogContainerPageOnly.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "LogContainerPageOnly", "value": *sdk.LogContainerPageOnly})
	}

	// Handling Primitives
	if !model.LogHttpHdrReferer.IsNull() && !model.LogHttpHdrReferer.IsUnknown() {
		sdk.LogHttpHdrReferer = model.LogHttpHdrReferer.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "LogHttpHdrReferer", "value": *sdk.LogHttpHdrReferer})
	}

	// Handling Primitives
	if !model.LogHttpHdrUserAgent.IsNull() && !model.LogHttpHdrUserAgent.IsUnknown() {
		sdk.LogHttpHdrUserAgent = model.LogHttpHdrUserAgent.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "LogHttpHdrUserAgent", "value": *sdk.LogHttpHdrUserAgent})
	}

	// Handling Primitives
	if !model.LogHttpHdrXff.IsNull() && !model.LogHttpHdrXff.IsUnknown() {
		sdk.LogHttpHdrXff = model.LogHttpHdrXff.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "LogHttpHdrXff", "value": *sdk.LogHttpHdrXff})
	}

	// Handling Lists
	if !model.MlavCategoryException.IsNull() && !model.MlavCategoryException.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field MlavCategoryException")
		diags.Append(model.MlavCategoryException.ElementsAs(ctx, &sdk.MlavCategoryException, false)...)
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	// Handling Lists
	if !model.Redirect.IsNull() && !model.Redirect.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Redirect")
		diags.Append(model.Redirect.ElementsAs(ctx, &sdk.Redirect, false)...)
	}

	// Handling Primitives
	if !model.SafeSearchEnforcement.IsNull() && !model.SafeSearchEnforcement.IsUnknown() {
		sdk.SafeSearchEnforcement = model.SafeSearchEnforcement.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "SafeSearchEnforcement", "value": *sdk.SafeSearchEnforcement})
	}

	// Handling Primitives
	if !model.Snippet.IsNull() && !model.Snippet.IsUnknown() {
		sdk.Snippet = model.Snippet.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.UrlAccessProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for UrlAccessProfiles ---
func packUrlAccessProfilesFromSdk(ctx context.Context, sdk security_services.UrlAccessProfiles) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.UrlAccessProfiles", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.UrlAccessProfiles
	var d diag.Diagnostics
	// Handling Lists
	if sdk.Alert != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Alert")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Alert, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Alert)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Alert = basetypes.NewListNull(elemType)
	}
	// Handling Lists
	if sdk.Allow != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Allow")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Allow, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Allow)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Allow = basetypes.NewListNull(elemType)
	}
	// Handling Lists
	if sdk.Block != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Block")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Block, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Block)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Block = basetypes.NewListNull(elemType)
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.CloudInlineCat != nil {
		model.CloudInlineCat = basetypes.NewBoolValue(*sdk.CloudInlineCat)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "CloudInlineCat", "value": *sdk.CloudInlineCat})
	} else {
		model.CloudInlineCat = basetypes.NewBoolNull()
	}
	// Handling Lists
	if sdk.Continue != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Continue")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Continue, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Continue)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Continue = basetypes.NewListNull(elemType)
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.CredentialEnforcement != nil {
		tflog.Debug(ctx, "Packing nested object for field CredentialEnforcement")
		packed, d := packUrlAccessProfilesCredentialEnforcementFromSdk(ctx, *sdk.CredentialEnforcement)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "CredentialEnforcement"})
		}
		model.CredentialEnforcement = packed
	} else {
		model.CredentialEnforcement = basetypes.NewObjectNull(models.UrlAccessProfilesCredentialEnforcement{}.AttrTypes())
	}
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
	if sdk.LocalInlineCat != nil {
		model.LocalInlineCat = basetypes.NewBoolValue(*sdk.LocalInlineCat)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "LocalInlineCat", "value": *sdk.LocalInlineCat})
	} else {
		model.LocalInlineCat = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.LogContainerPageOnly != nil {
		model.LogContainerPageOnly = basetypes.NewBoolValue(*sdk.LogContainerPageOnly)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "LogContainerPageOnly", "value": *sdk.LogContainerPageOnly})
	} else {
		model.LogContainerPageOnly = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.LogHttpHdrReferer != nil {
		model.LogHttpHdrReferer = basetypes.NewBoolValue(*sdk.LogHttpHdrReferer)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "LogHttpHdrReferer", "value": *sdk.LogHttpHdrReferer})
	} else {
		model.LogHttpHdrReferer = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.LogHttpHdrUserAgent != nil {
		model.LogHttpHdrUserAgent = basetypes.NewBoolValue(*sdk.LogHttpHdrUserAgent)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "LogHttpHdrUserAgent", "value": *sdk.LogHttpHdrUserAgent})
	} else {
		model.LogHttpHdrUserAgent = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.LogHttpHdrXff != nil {
		model.LogHttpHdrXff = basetypes.NewBoolValue(*sdk.LogHttpHdrXff)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "LogHttpHdrXff", "value": *sdk.LogHttpHdrXff})
	} else {
		model.LogHttpHdrXff = basetypes.NewBoolNull()
	}
	// Handling Lists
	if sdk.MlavCategoryException != nil {
		tflog.Debug(ctx, "Packing list of primitives for field MlavCategoryException")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.MlavCategoryException, d = basetypes.NewListValueFrom(ctx, elemType, sdk.MlavCategoryException)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.MlavCategoryException = basetypes.NewListNull(elemType)
	}
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	// Handling Lists
	if sdk.Redirect != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Redirect")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Redirect, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Redirect)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Redirect = basetypes.NewListNull(elemType)
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.SafeSearchEnforcement != nil {
		model.SafeSearchEnforcement = basetypes.NewBoolValue(*sdk.SafeSearchEnforcement)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "SafeSearchEnforcement", "value": *sdk.SafeSearchEnforcement})
	} else {
		model.SafeSearchEnforcement = basetypes.NewBoolNull()
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

	obj, d := types.ObjectValueFrom(ctx, models.UrlAccessProfiles{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.UrlAccessProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for UrlAccessProfiles ---
func unpackUrlAccessProfilesListToSdk(ctx context.Context, list types.List) ([]security_services.UrlAccessProfiles, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.UrlAccessProfiles")
	diags := diag.Diagnostics{}
	var data []models.UrlAccessProfiles
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.UrlAccessProfiles, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.UrlAccessProfiles{}.AttrTypes(), &item)
		unpacked, d := unpackUrlAccessProfilesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.UrlAccessProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for UrlAccessProfiles ---
func packUrlAccessProfilesListFromSdk(ctx context.Context, sdks []security_services.UrlAccessProfiles) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.UrlAccessProfiles")
	diags := diag.Diagnostics{}
	var data []models.UrlAccessProfiles

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.UrlAccessProfiles
		obj, d := packUrlAccessProfilesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.UrlAccessProfiles{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.UrlAccessProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.UrlAccessProfiles{}.AttrType(), data)
}

// --- Unpacker for UrlAccessProfilesCredentialEnforcement ---
func unpackUrlAccessProfilesCredentialEnforcementToSdk(ctx context.Context, obj types.Object) (*security_services.UrlAccessProfilesCredentialEnforcement, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.UrlAccessProfilesCredentialEnforcement", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.UrlAccessProfilesCredentialEnforcement
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.UrlAccessProfilesCredentialEnforcement
	var d diag.Diagnostics
	// Handling Lists
	if !model.Alert.IsNull() && !model.Alert.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Alert")
		diags.Append(model.Alert.ElementsAs(ctx, &sdk.Alert, false)...)
	}

	// Handling Lists
	if !model.Allow.IsNull() && !model.Allow.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Allow")
		diags.Append(model.Allow.ElementsAs(ctx, &sdk.Allow, false)...)
	}

	// Handling Lists
	if !model.Block.IsNull() && !model.Block.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Block")
		diags.Append(model.Block.ElementsAs(ctx, &sdk.Block, false)...)
	}

	// Handling Lists
	if !model.Continue.IsNull() && !model.Continue.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Continue")
		diags.Append(model.Continue.ElementsAs(ctx, &sdk.Continue, false)...)
	}

	// Handling Primitives
	if !model.LogSeverity.IsNull() && !model.LogSeverity.IsUnknown() {
		sdk.LogSeverity = model.LogSeverity.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "LogSeverity", "value": *sdk.LogSeverity})
	}

	// Handling Objects
	if !model.Mode.IsNull() && !model.Mode.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Mode")
		unpacked, d := unpackUrlAccessProfilesCredentialEnforcementModeToSdk(ctx, model.Mode)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Mode"})
		}
		if unpacked != nil {
			sdk.Mode = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.UrlAccessProfilesCredentialEnforcement", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for UrlAccessProfilesCredentialEnforcement ---
func packUrlAccessProfilesCredentialEnforcementFromSdk(ctx context.Context, sdk security_services.UrlAccessProfilesCredentialEnforcement) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.UrlAccessProfilesCredentialEnforcement", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.UrlAccessProfilesCredentialEnforcement
	var d diag.Diagnostics
	// Handling Lists
	if sdk.Alert != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Alert")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Alert, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Alert)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Alert = basetypes.NewListNull(elemType)
	}
	// Handling Lists
	if sdk.Allow != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Allow")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Allow, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Allow)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Allow = basetypes.NewListNull(elemType)
	}
	// Handling Lists
	if sdk.Block != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Block")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Block, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Block)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Block = basetypes.NewListNull(elemType)
	}
	// Handling Lists
	if sdk.Continue != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Continue")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Continue, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Continue)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Continue = basetypes.NewListNull(elemType)
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.LogSeverity != nil {
		model.LogSeverity = basetypes.NewStringValue(*sdk.LogSeverity)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "LogSeverity", "value": *sdk.LogSeverity})
	} else {
		model.LogSeverity = basetypes.NewStringNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Mode != nil {
		tflog.Debug(ctx, "Packing nested object for field Mode")
		packed, d := packUrlAccessProfilesCredentialEnforcementModeFromSdk(ctx, *sdk.Mode)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Mode"})
		}
		model.Mode = packed
	} else {
		model.Mode = basetypes.NewObjectNull(models.UrlAccessProfilesCredentialEnforcementMode{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.UrlAccessProfilesCredentialEnforcement{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.UrlAccessProfilesCredentialEnforcement", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for UrlAccessProfilesCredentialEnforcement ---
func unpackUrlAccessProfilesCredentialEnforcementListToSdk(ctx context.Context, list types.List) ([]security_services.UrlAccessProfilesCredentialEnforcement, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.UrlAccessProfilesCredentialEnforcement")
	diags := diag.Diagnostics{}
	var data []models.UrlAccessProfilesCredentialEnforcement
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.UrlAccessProfilesCredentialEnforcement, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.UrlAccessProfilesCredentialEnforcement{}.AttrTypes(), &item)
		unpacked, d := unpackUrlAccessProfilesCredentialEnforcementToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.UrlAccessProfilesCredentialEnforcement", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for UrlAccessProfilesCredentialEnforcement ---
func packUrlAccessProfilesCredentialEnforcementListFromSdk(ctx context.Context, sdks []security_services.UrlAccessProfilesCredentialEnforcement) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.UrlAccessProfilesCredentialEnforcement")
	diags := diag.Diagnostics{}
	var data []models.UrlAccessProfilesCredentialEnforcement

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.UrlAccessProfilesCredentialEnforcement
		obj, d := packUrlAccessProfilesCredentialEnforcementFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.UrlAccessProfilesCredentialEnforcement{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.UrlAccessProfilesCredentialEnforcement", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.UrlAccessProfilesCredentialEnforcement{}.AttrType(), data)
}

// --- Unpacker for UrlAccessProfilesCredentialEnforcementMode ---
func unpackUrlAccessProfilesCredentialEnforcementModeToSdk(ctx context.Context, obj types.Object) (*security_services.UrlAccessProfilesCredentialEnforcementMode, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.UrlAccessProfilesCredentialEnforcementMode", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.UrlAccessProfilesCredentialEnforcementMode
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.UrlAccessProfilesCredentialEnforcementMode
	var d diag.Diagnostics
	// Handling Typeless Objects
	if !model.Disabled.IsNull() && !model.Disabled.IsUnknown() {
		tflog.Debug(ctx, "Unpacking typeless object for field Disabled")
		sdk.Disabled = make(map[string]interface{})
	}

	// Handling Typeless Objects
	if !model.DomainCredentials.IsNull() && !model.DomainCredentials.IsUnknown() {
		tflog.Debug(ctx, "Unpacking typeless object for field DomainCredentials")
		sdk.DomainCredentials = make(map[string]interface{})
	}

	// Handling Primitives
	if !model.GroupMapping.IsNull() && !model.GroupMapping.IsUnknown() {
		sdk.GroupMapping = model.GroupMapping.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "GroupMapping", "value": *sdk.GroupMapping})
	}

	// Handling Typeless Objects
	if !model.IpUser.IsNull() && !model.IpUser.IsUnknown() {
		tflog.Debug(ctx, "Unpacking typeless object for field IpUser")
		sdk.IpUser = make(map[string]interface{})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.UrlAccessProfilesCredentialEnforcementMode", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for UrlAccessProfilesCredentialEnforcementMode ---
func packUrlAccessProfilesCredentialEnforcementModeFromSdk(ctx context.Context, sdk security_services.UrlAccessProfilesCredentialEnforcementMode) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.UrlAccessProfilesCredentialEnforcementMode", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.UrlAccessProfilesCredentialEnforcementMode
	var d diag.Diagnostics
	// Handling Objects
	// This is a marker object (e.g. CHAP: {}). We just need to create an empty, non-null object.
	if sdk.Disabled != nil && !reflect.ValueOf(sdk.Disabled).IsNil() {
		tflog.Debug(ctx, "Packing typeless object for field Disabled")
		var d diag.Diagnostics
		// Create an empty object with no attributes, which signifies its presence.
		model.Disabled, d = basetypes.NewObjectValue(map[string]attr.Type{}, map[string]attr.Value{})
		diags.Append(d...)
	} else {
		// Since this field is part of a oneOf, being nil means it's not selected.
		// We make the object null with an empty attribute map.
		model.Disabled = basetypes.NewObjectNull(map[string]attr.Type{})
	}
	// Handling Objects
	// This is a marker object (e.g. CHAP: {}). We just need to create an empty, non-null object.
	if sdk.DomainCredentials != nil && !reflect.ValueOf(sdk.DomainCredentials).IsNil() {
		tflog.Debug(ctx, "Packing typeless object for field DomainCredentials")
		var d diag.Diagnostics
		// Create an empty object with no attributes, which signifies its presence.
		model.DomainCredentials, d = basetypes.NewObjectValue(map[string]attr.Type{}, map[string]attr.Value{})
		diags.Append(d...)
	} else {
		// Since this field is part of a oneOf, being nil means it's not selected.
		// We make the object null with an empty attribute map.
		model.DomainCredentials = basetypes.NewObjectNull(map[string]attr.Type{})
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.GroupMapping != nil {
		model.GroupMapping = basetypes.NewStringValue(*sdk.GroupMapping)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "GroupMapping", "value": *sdk.GroupMapping})
	} else {
		model.GroupMapping = basetypes.NewStringNull()
	}
	// Handling Objects
	// This is a marker object (e.g. CHAP: {}). We just need to create an empty, non-null object.
	if sdk.IpUser != nil && !reflect.ValueOf(sdk.IpUser).IsNil() {
		tflog.Debug(ctx, "Packing typeless object for field IpUser")
		var d diag.Diagnostics
		// Create an empty object with no attributes, which signifies its presence.
		model.IpUser, d = basetypes.NewObjectValue(map[string]attr.Type{}, map[string]attr.Value{})
		diags.Append(d...)
	} else {
		// Since this field is part of a oneOf, being nil means it's not selected.
		// We make the object null with an empty attribute map.
		model.IpUser = basetypes.NewObjectNull(map[string]attr.Type{})
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.UrlAccessProfilesCredentialEnforcementMode{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.UrlAccessProfilesCredentialEnforcementMode", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for UrlAccessProfilesCredentialEnforcementMode ---
func unpackUrlAccessProfilesCredentialEnforcementModeListToSdk(ctx context.Context, list types.List) ([]security_services.UrlAccessProfilesCredentialEnforcementMode, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.UrlAccessProfilesCredentialEnforcementMode")
	diags := diag.Diagnostics{}
	var data []models.UrlAccessProfilesCredentialEnforcementMode
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.UrlAccessProfilesCredentialEnforcementMode, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.UrlAccessProfilesCredentialEnforcementMode{}.AttrTypes(), &item)
		unpacked, d := unpackUrlAccessProfilesCredentialEnforcementModeToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.UrlAccessProfilesCredentialEnforcementMode", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for UrlAccessProfilesCredentialEnforcementMode ---
func packUrlAccessProfilesCredentialEnforcementModeListFromSdk(ctx context.Context, sdks []security_services.UrlAccessProfilesCredentialEnforcementMode) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.UrlAccessProfilesCredentialEnforcementMode")
	diags := diag.Diagnostics{}
	var data []models.UrlAccessProfilesCredentialEnforcementMode

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.UrlAccessProfilesCredentialEnforcementMode
		obj, d := packUrlAccessProfilesCredentialEnforcementModeFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.UrlAccessProfilesCredentialEnforcementMode{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.UrlAccessProfilesCredentialEnforcementMode", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.UrlAccessProfilesCredentialEnforcementMode{}.AttrType(), data)
}
