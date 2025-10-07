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

// --- Unpacker for AntiSpywareSignatures ---
func unpackAntiSpywareSignaturesToSdk(ctx context.Context, obj types.Object) (*security_services.AntiSpywareSignatures, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.AntiSpywareSignatures", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.AntiSpywareSignatures
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.AntiSpywareSignatures
	var d diag.Diagnostics
	// Handling Lists
	if !model.Bugtraq.IsNull() && !model.Bugtraq.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Bugtraq")
		diags.Append(model.Bugtraq.ElementsAs(ctx, &sdk.Bugtraq, false)...)
	}

	// Handling Primitives
	if !model.Comment.IsNull() && !model.Comment.IsUnknown() {
		sdk.Comment = model.Comment.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Comment", "value": *sdk.Comment})
	}

	// Handling Lists
	if !model.Cve.IsNull() && !model.Cve.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Cve")
		diags.Append(model.Cve.ElementsAs(ctx, &sdk.Cve, false)...)
	}

	// Handling Objects
	if !model.DefaultAction.IsNull() && !model.DefaultAction.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field DefaultAction")
		unpacked, d := unpackAntiSpywareSignaturesDefaultActionToSdk(ctx, model.DefaultAction)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "DefaultAction"})
		}
		if unpacked != nil {
			sdk.DefaultAction = unpacked
		}
	}

	// Handling Primitives
	if !model.Device.IsNull() && !model.Device.IsUnknown() {
		sdk.Device = model.Device.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Device", "value": *sdk.Device})
	}

	// Handling Primitives
	if !model.Direction.IsNull() && !model.Direction.IsUnknown() {
		sdk.Direction = model.Direction.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Direction", "value": *sdk.Direction})
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

	// Handling Lists
	if !model.Reference.IsNull() && !model.Reference.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Reference")
		diags.Append(model.Reference.ElementsAs(ctx, &sdk.Reference, false)...)
	}

	// Handling Primitives
	if !model.Severity.IsNull() && !model.Severity.IsUnknown() {
		sdk.Severity = model.Severity.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Severity", "value": *sdk.Severity})
	}

	// Handling Objects
	if !model.Signature.IsNull() && !model.Signature.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Signature")
		unpacked, d := unpackAntiSpywareSignaturesSignatureToSdk(ctx, model.Signature)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Signature"})
		}
		if unpacked != nil {
			sdk.Signature = unpacked
		}
	}

	// Handling Primitives
	if !model.Snippet.IsNull() && !model.Snippet.IsUnknown() {
		sdk.Snippet = model.Snippet.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	}

	// Handling Primitives
	if !model.ThreatId.IsNull() && !model.ThreatId.IsUnknown() {
		sdk.ThreatId = int32(model.ThreatId.ValueInt64())
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "ThreatId", "value": sdk.ThreatId})
	}

	// Handling Primitives
	if !model.Threatname.IsNull() && !model.Threatname.IsUnknown() {
		sdk.Threatname = model.Threatname.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Threatname", "value": sdk.Threatname})
	}

	// Handling Lists
	if !model.Vendor.IsNull() && !model.Vendor.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Vendor")
		diags.Append(model.Vendor.ElementsAs(ctx, &sdk.Vendor, false)...)
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.AntiSpywareSignatures", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for AntiSpywareSignatures ---
func packAntiSpywareSignaturesFromSdk(ctx context.Context, sdk security_services.AntiSpywareSignatures) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.AntiSpywareSignatures", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.AntiSpywareSignatures
	var d diag.Diagnostics
	// Handling Lists
	if sdk.Bugtraq != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Bugtraq")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Bugtraq, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Bugtraq)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Bugtraq = basetypes.NewListNull(elemType)
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Comment != nil {
		model.Comment = basetypes.NewStringValue(*sdk.Comment)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Comment", "value": *sdk.Comment})
	} else {
		model.Comment = basetypes.NewStringNull()
	}
	// Handling Lists
	if sdk.Cve != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Cve")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Cve, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Cve)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Cve = basetypes.NewListNull(elemType)
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.DefaultAction != nil {
		tflog.Debug(ctx, "Packing nested object for field DefaultAction")
		packed, d := packAntiSpywareSignaturesDefaultActionFromSdk(ctx, *sdk.DefaultAction)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "DefaultAction"})
		}
		model.DefaultAction = packed
	} else {
		model.DefaultAction = basetypes.NewObjectNull(models.AntiSpywareSignaturesDefaultAction{}.AttrTypes())
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
	if sdk.Direction != nil {
		model.Direction = basetypes.NewStringValue(*sdk.Direction)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Direction", "value": *sdk.Direction})
	} else {
		model.Direction = basetypes.NewStringNull()
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
	// Handling Lists
	if sdk.Reference != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Reference")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Reference, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Reference)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Reference = basetypes.NewListNull(elemType)
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Severity != nil {
		model.Severity = basetypes.NewStringValue(*sdk.Severity)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Severity", "value": *sdk.Severity})
	} else {
		model.Severity = basetypes.NewStringNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Signature != nil {
		tflog.Debug(ctx, "Packing nested object for field Signature")
		packed, d := packAntiSpywareSignaturesSignatureFromSdk(ctx, *sdk.Signature)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Signature"})
		}
		model.Signature = packed
	} else {
		model.Signature = basetypes.NewObjectNull(models.AntiSpywareSignaturesSignature{}.AttrTypes())
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
	model.ThreatId = basetypes.NewInt64Value(int64(sdk.ThreatId))
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "ThreatId", "value": sdk.ThreatId})
	// Handling Primitives
	// Standard primitive packing
	model.Threatname = basetypes.NewStringValue(sdk.Threatname)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Threatname", "value": sdk.Threatname})
	// Handling Lists
	if sdk.Vendor != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Vendor")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Vendor, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Vendor)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Vendor = basetypes.NewListNull(elemType)
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.AntiSpywareSignatures{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.AntiSpywareSignatures", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for AntiSpywareSignatures ---
func unpackAntiSpywareSignaturesListToSdk(ctx context.Context, list types.List) ([]security_services.AntiSpywareSignatures, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.AntiSpywareSignatures")
	diags := diag.Diagnostics{}
	var data []models.AntiSpywareSignatures
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.AntiSpywareSignatures, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.AntiSpywareSignatures{}.AttrTypes(), &item)
		unpacked, d := unpackAntiSpywareSignaturesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.AntiSpywareSignatures", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for AntiSpywareSignatures ---
func packAntiSpywareSignaturesListFromSdk(ctx context.Context, sdks []security_services.AntiSpywareSignatures) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.AntiSpywareSignatures")
	diags := diag.Diagnostics{}
	var data []models.AntiSpywareSignatures

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.AntiSpywareSignatures
		obj, d := packAntiSpywareSignaturesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.AntiSpywareSignatures{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.AntiSpywareSignatures", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.AntiSpywareSignatures{}.AttrType(), data)
}

// --- Unpacker for AntiSpywareSignaturesDefaultAction ---
func unpackAntiSpywareSignaturesDefaultActionToSdk(ctx context.Context, obj types.Object) (*security_services.AntiSpywareSignaturesDefaultAction, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.AntiSpywareSignaturesDefaultAction", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.AntiSpywareSignaturesDefaultAction
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.AntiSpywareSignaturesDefaultAction
	var d diag.Diagnostics
	// Handling Typeless Objects
	if !model.Alert.IsNull() && !model.Alert.IsUnknown() {
		tflog.Debug(ctx, "Unpacking typeless object for field Alert")
		sdk.Alert = make(map[string]interface{})
	}

	// Handling Typeless Objects
	if !model.Allow.IsNull() && !model.Allow.IsUnknown() {
		tflog.Debug(ctx, "Unpacking typeless object for field Allow")
		sdk.Allow = make(map[string]interface{})
	}

	// Handling Objects
	if !model.BlockIp.IsNull() && !model.BlockIp.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field BlockIp")
		unpacked, d := unpackAntiSpywareSignaturesDefaultActionBlockIpToSdk(ctx, model.BlockIp)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "BlockIp"})
		}
		if unpacked != nil {
			sdk.BlockIp = unpacked
		}
	}

	// Handling Typeless Objects
	if !model.Drop.IsNull() && !model.Drop.IsUnknown() {
		tflog.Debug(ctx, "Unpacking typeless object for field Drop")
		sdk.Drop = make(map[string]interface{})
	}

	// Handling Typeless Objects
	if !model.ResetBoth.IsNull() && !model.ResetBoth.IsUnknown() {
		tflog.Debug(ctx, "Unpacking typeless object for field ResetBoth")
		sdk.ResetBoth = make(map[string]interface{})
	}

	// Handling Typeless Objects
	if !model.ResetClient.IsNull() && !model.ResetClient.IsUnknown() {
		tflog.Debug(ctx, "Unpacking typeless object for field ResetClient")
		sdk.ResetClient = make(map[string]interface{})
	}

	// Handling Typeless Objects
	if !model.ResetServer.IsNull() && !model.ResetServer.IsUnknown() {
		tflog.Debug(ctx, "Unpacking typeless object for field ResetServer")
		sdk.ResetServer = make(map[string]interface{})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.AntiSpywareSignaturesDefaultAction", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for AntiSpywareSignaturesDefaultAction ---
func packAntiSpywareSignaturesDefaultActionFromSdk(ctx context.Context, sdk security_services.AntiSpywareSignaturesDefaultAction) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.AntiSpywareSignaturesDefaultAction", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.AntiSpywareSignaturesDefaultAction
	var d diag.Diagnostics
	// Handling Objects
	// This is a marker object (e.g. CHAP: {}). We just need to create an empty, non-null object.
	if sdk.Alert != nil && !reflect.ValueOf(sdk.Alert).IsNil() {
		tflog.Debug(ctx, "Packing typeless object for field Alert")
		var d diag.Diagnostics
		// Create an empty object with no attributes, which signifies its presence.
		model.Alert, d = basetypes.NewObjectValue(map[string]attr.Type{}, map[string]attr.Value{})
		diags.Append(d...)
	} else {
		// Since this field is part of a oneOf, being nil means it's not selected.
		// We make the object null with an empty attribute map.
		model.Alert = basetypes.NewObjectNull(map[string]attr.Type{})
	}
	// Handling Objects
	// This is a marker object (e.g. CHAP: {}). We just need to create an empty, non-null object.
	if sdk.Allow != nil && !reflect.ValueOf(sdk.Allow).IsNil() {
		tflog.Debug(ctx, "Packing typeless object for field Allow")
		var d diag.Diagnostics
		// Create an empty object with no attributes, which signifies its presence.
		model.Allow, d = basetypes.NewObjectValue(map[string]attr.Type{}, map[string]attr.Value{})
		diags.Append(d...)
	} else {
		// Since this field is part of a oneOf, being nil means it's not selected.
		// We make the object null with an empty attribute map.
		model.Allow = basetypes.NewObjectNull(map[string]attr.Type{})
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.BlockIp != nil {
		tflog.Debug(ctx, "Packing nested object for field BlockIp")
		packed, d := packAntiSpywareSignaturesDefaultActionBlockIpFromSdk(ctx, *sdk.BlockIp)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "BlockIp"})
		}
		model.BlockIp = packed
	} else {
		model.BlockIp = basetypes.NewObjectNull(models.AntiSpywareSignaturesDefaultActionBlockIp{}.AttrTypes())
	}
	// Handling Objects
	// This is a marker object (e.g. CHAP: {}). We just need to create an empty, non-null object.
	if sdk.Drop != nil && !reflect.ValueOf(sdk.Drop).IsNil() {
		tflog.Debug(ctx, "Packing typeless object for field Drop")
		var d diag.Diagnostics
		// Create an empty object with no attributes, which signifies its presence.
		model.Drop, d = basetypes.NewObjectValue(map[string]attr.Type{}, map[string]attr.Value{})
		diags.Append(d...)
	} else {
		// Since this field is part of a oneOf, being nil means it's not selected.
		// We make the object null with an empty attribute map.
		model.Drop = basetypes.NewObjectNull(map[string]attr.Type{})
	}
	// Handling Objects
	// This is a marker object (e.g. CHAP: {}). We just need to create an empty, non-null object.
	if sdk.ResetBoth != nil && !reflect.ValueOf(sdk.ResetBoth).IsNil() {
		tflog.Debug(ctx, "Packing typeless object for field ResetBoth")
		var d diag.Diagnostics
		// Create an empty object with no attributes, which signifies its presence.
		model.ResetBoth, d = basetypes.NewObjectValue(map[string]attr.Type{}, map[string]attr.Value{})
		diags.Append(d...)
	} else {
		// Since this field is part of a oneOf, being nil means it's not selected.
		// We make the object null with an empty attribute map.
		model.ResetBoth = basetypes.NewObjectNull(map[string]attr.Type{})
	}
	// Handling Objects
	// This is a marker object (e.g. CHAP: {}). We just need to create an empty, non-null object.
	if sdk.ResetClient != nil && !reflect.ValueOf(sdk.ResetClient).IsNil() {
		tflog.Debug(ctx, "Packing typeless object for field ResetClient")
		var d diag.Diagnostics
		// Create an empty object with no attributes, which signifies its presence.
		model.ResetClient, d = basetypes.NewObjectValue(map[string]attr.Type{}, map[string]attr.Value{})
		diags.Append(d...)
	} else {
		// Since this field is part of a oneOf, being nil means it's not selected.
		// We make the object null with an empty attribute map.
		model.ResetClient = basetypes.NewObjectNull(map[string]attr.Type{})
	}
	// Handling Objects
	// This is a marker object (e.g. CHAP: {}). We just need to create an empty, non-null object.
	if sdk.ResetServer != nil && !reflect.ValueOf(sdk.ResetServer).IsNil() {
		tflog.Debug(ctx, "Packing typeless object for field ResetServer")
		var d diag.Diagnostics
		// Create an empty object with no attributes, which signifies its presence.
		model.ResetServer, d = basetypes.NewObjectValue(map[string]attr.Type{}, map[string]attr.Value{})
		diags.Append(d...)
	} else {
		// Since this field is part of a oneOf, being nil means it's not selected.
		// We make the object null with an empty attribute map.
		model.ResetServer = basetypes.NewObjectNull(map[string]attr.Type{})
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.AntiSpywareSignaturesDefaultAction{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.AntiSpywareSignaturesDefaultAction", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for AntiSpywareSignaturesDefaultAction ---
func unpackAntiSpywareSignaturesDefaultActionListToSdk(ctx context.Context, list types.List) ([]security_services.AntiSpywareSignaturesDefaultAction, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.AntiSpywareSignaturesDefaultAction")
	diags := diag.Diagnostics{}
	var data []models.AntiSpywareSignaturesDefaultAction
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.AntiSpywareSignaturesDefaultAction, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.AntiSpywareSignaturesDefaultAction{}.AttrTypes(), &item)
		unpacked, d := unpackAntiSpywareSignaturesDefaultActionToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.AntiSpywareSignaturesDefaultAction", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for AntiSpywareSignaturesDefaultAction ---
func packAntiSpywareSignaturesDefaultActionListFromSdk(ctx context.Context, sdks []security_services.AntiSpywareSignaturesDefaultAction) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.AntiSpywareSignaturesDefaultAction")
	diags := diag.Diagnostics{}
	var data []models.AntiSpywareSignaturesDefaultAction

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.AntiSpywareSignaturesDefaultAction
		obj, d := packAntiSpywareSignaturesDefaultActionFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.AntiSpywareSignaturesDefaultAction{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.AntiSpywareSignaturesDefaultAction", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.AntiSpywareSignaturesDefaultAction{}.AttrType(), data)
}

// --- Unpacker for AntiSpywareSignaturesDefaultActionBlockIp ---
func unpackAntiSpywareSignaturesDefaultActionBlockIpToSdk(ctx context.Context, obj types.Object) (*security_services.AntiSpywareSignaturesDefaultActionBlockIp, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.AntiSpywareSignaturesDefaultActionBlockIp", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.AntiSpywareSignaturesDefaultActionBlockIp
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.AntiSpywareSignaturesDefaultActionBlockIp
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Duration.IsNull() && !model.Duration.IsUnknown() {
		val := int32(model.Duration.ValueInt64())
		sdk.Duration = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Duration", "value": *sdk.Duration})
	}

	// Handling Primitives
	if !model.TrackBy.IsNull() && !model.TrackBy.IsUnknown() {
		sdk.TrackBy = model.TrackBy.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "TrackBy", "value": *sdk.TrackBy})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.AntiSpywareSignaturesDefaultActionBlockIp", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for AntiSpywareSignaturesDefaultActionBlockIp ---
func packAntiSpywareSignaturesDefaultActionBlockIpFromSdk(ctx context.Context, sdk security_services.AntiSpywareSignaturesDefaultActionBlockIp) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.AntiSpywareSignaturesDefaultActionBlockIp", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.AntiSpywareSignaturesDefaultActionBlockIp
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Duration != nil {
		model.Duration = basetypes.NewInt64Value(int64(*sdk.Duration))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Duration", "value": *sdk.Duration})
	} else {
		model.Duration = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.TrackBy != nil {
		model.TrackBy = basetypes.NewStringValue(*sdk.TrackBy)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "TrackBy", "value": *sdk.TrackBy})
	} else {
		model.TrackBy = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.AntiSpywareSignaturesDefaultActionBlockIp{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.AntiSpywareSignaturesDefaultActionBlockIp", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for AntiSpywareSignaturesDefaultActionBlockIp ---
func unpackAntiSpywareSignaturesDefaultActionBlockIpListToSdk(ctx context.Context, list types.List) ([]security_services.AntiSpywareSignaturesDefaultActionBlockIp, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.AntiSpywareSignaturesDefaultActionBlockIp")
	diags := diag.Diagnostics{}
	var data []models.AntiSpywareSignaturesDefaultActionBlockIp
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.AntiSpywareSignaturesDefaultActionBlockIp, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.AntiSpywareSignaturesDefaultActionBlockIp{}.AttrTypes(), &item)
		unpacked, d := unpackAntiSpywareSignaturesDefaultActionBlockIpToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.AntiSpywareSignaturesDefaultActionBlockIp", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for AntiSpywareSignaturesDefaultActionBlockIp ---
func packAntiSpywareSignaturesDefaultActionBlockIpListFromSdk(ctx context.Context, sdks []security_services.AntiSpywareSignaturesDefaultActionBlockIp) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.AntiSpywareSignaturesDefaultActionBlockIp")
	diags := diag.Diagnostics{}
	var data []models.AntiSpywareSignaturesDefaultActionBlockIp

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.AntiSpywareSignaturesDefaultActionBlockIp
		obj, d := packAntiSpywareSignaturesDefaultActionBlockIpFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.AntiSpywareSignaturesDefaultActionBlockIp{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.AntiSpywareSignaturesDefaultActionBlockIp", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.AntiSpywareSignaturesDefaultActionBlockIp{}.AttrType(), data)
}

// --- Unpacker for AntiSpywareSignaturesSignature ---
func unpackAntiSpywareSignaturesSignatureToSdk(ctx context.Context, obj types.Object) (*security_services.AntiSpywareSignaturesSignature, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.AntiSpywareSignaturesSignature", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.AntiSpywareSignaturesSignature
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.AntiSpywareSignaturesSignature
	var d diag.Diagnostics
	// Handling Objects
	if !model.Combination.IsNull() && !model.Combination.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Combination")
		unpacked, d := unpackAntiSpywareSignaturesSignatureCombinationToSdk(ctx, model.Combination)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Combination"})
		}
		if unpacked != nil {
			sdk.Combination = unpacked
		}
	}

	// Handling Lists
	if !model.Standard.IsNull() && !model.Standard.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field Standard")
		unpacked, d := unpackAntiSpywareSignaturesSignatureStandardInnerListToSdk(ctx, model.Standard)
		diags.Append(d...)
		sdk.Standard = unpacked
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.AntiSpywareSignaturesSignature", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for AntiSpywareSignaturesSignature ---
func packAntiSpywareSignaturesSignatureFromSdk(ctx context.Context, sdk security_services.AntiSpywareSignaturesSignature) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.AntiSpywareSignaturesSignature", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.AntiSpywareSignaturesSignature
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Combination != nil {
		tflog.Debug(ctx, "Packing nested object for field Combination")
		packed, d := packAntiSpywareSignaturesSignatureCombinationFromSdk(ctx, *sdk.Combination)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Combination"})
		}
		model.Combination = packed
	} else {
		model.Combination = basetypes.NewObjectNull(models.AntiSpywareSignaturesSignatureCombination{}.AttrTypes())
	}
	// Handling Lists
	if sdk.Standard != nil {
		tflog.Debug(ctx, "Packing list of objects for field Standard")
		packed, d := packAntiSpywareSignaturesSignatureStandardInnerListFromSdk(ctx, sdk.Standard)
		diags.Append(d...)
		model.Standard = packed
	} else {
		model.Standard = basetypes.NewListNull(models.AntiSpywareSignaturesSignatureStandardInner{}.AttrType())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.AntiSpywareSignaturesSignature{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.AntiSpywareSignaturesSignature", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for AntiSpywareSignaturesSignature ---
func unpackAntiSpywareSignaturesSignatureListToSdk(ctx context.Context, list types.List) ([]security_services.AntiSpywareSignaturesSignature, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.AntiSpywareSignaturesSignature")
	diags := diag.Diagnostics{}
	var data []models.AntiSpywareSignaturesSignature
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.AntiSpywareSignaturesSignature, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.AntiSpywareSignaturesSignature{}.AttrTypes(), &item)
		unpacked, d := unpackAntiSpywareSignaturesSignatureToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.AntiSpywareSignaturesSignature", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for AntiSpywareSignaturesSignature ---
func packAntiSpywareSignaturesSignatureListFromSdk(ctx context.Context, sdks []security_services.AntiSpywareSignaturesSignature) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.AntiSpywareSignaturesSignature")
	diags := diag.Diagnostics{}
	var data []models.AntiSpywareSignaturesSignature

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.AntiSpywareSignaturesSignature
		obj, d := packAntiSpywareSignaturesSignatureFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.AntiSpywareSignaturesSignature{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.AntiSpywareSignaturesSignature", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.AntiSpywareSignaturesSignature{}.AttrType(), data)
}

// --- Unpacker for AntiSpywareSignaturesSignatureCombination ---
func unpackAntiSpywareSignaturesSignatureCombinationToSdk(ctx context.Context, obj types.Object) (*security_services.AntiSpywareSignaturesSignatureCombination, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.AntiSpywareSignaturesSignatureCombination", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.AntiSpywareSignaturesSignatureCombination
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.AntiSpywareSignaturesSignatureCombination
	var d diag.Diagnostics
	// Handling Lists
	if !model.AndCondition.IsNull() && !model.AndCondition.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field AndCondition")
		unpacked, d := unpackAntiSpywareSignaturesSignatureCombinationAndConditionInnerListToSdk(ctx, model.AndCondition)
		diags.Append(d...)
		sdk.AndCondition = unpacked
	}

	// Handling Primitives
	if !model.OrderFree.IsNull() && !model.OrderFree.IsUnknown() {
		sdk.OrderFree = model.OrderFree.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "OrderFree", "value": *sdk.OrderFree})
	}

	// Handling Objects
	if !model.TimeAttribute.IsNull() && !model.TimeAttribute.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field TimeAttribute")
		unpacked, d := unpackAntiSpywareSignaturesSignatureCombinationTimeAttributeToSdk(ctx, model.TimeAttribute)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "TimeAttribute"})
		}
		if unpacked != nil {
			sdk.TimeAttribute = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.AntiSpywareSignaturesSignatureCombination", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for AntiSpywareSignaturesSignatureCombination ---
func packAntiSpywareSignaturesSignatureCombinationFromSdk(ctx context.Context, sdk security_services.AntiSpywareSignaturesSignatureCombination) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.AntiSpywareSignaturesSignatureCombination", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.AntiSpywareSignaturesSignatureCombination
	var d diag.Diagnostics
	// Handling Lists
	if sdk.AndCondition != nil {
		tflog.Debug(ctx, "Packing list of objects for field AndCondition")
		packed, d := packAntiSpywareSignaturesSignatureCombinationAndConditionInnerListFromSdk(ctx, sdk.AndCondition)
		diags.Append(d...)
		model.AndCondition = packed
	} else {
		model.AndCondition = basetypes.NewListNull(models.AntiSpywareSignaturesSignatureCombinationAndConditionInner{}.AttrType())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.OrderFree != nil {
		model.OrderFree = basetypes.NewBoolValue(*sdk.OrderFree)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "OrderFree", "value": *sdk.OrderFree})
	} else {
		model.OrderFree = basetypes.NewBoolNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.TimeAttribute != nil {
		tflog.Debug(ctx, "Packing nested object for field TimeAttribute")
		packed, d := packAntiSpywareSignaturesSignatureCombinationTimeAttributeFromSdk(ctx, *sdk.TimeAttribute)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "TimeAttribute"})
		}
		model.TimeAttribute = packed
	} else {
		model.TimeAttribute = basetypes.NewObjectNull(models.AntiSpywareSignaturesSignatureCombinationTimeAttribute{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.AntiSpywareSignaturesSignatureCombination{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.AntiSpywareSignaturesSignatureCombination", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for AntiSpywareSignaturesSignatureCombination ---
func unpackAntiSpywareSignaturesSignatureCombinationListToSdk(ctx context.Context, list types.List) ([]security_services.AntiSpywareSignaturesSignatureCombination, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.AntiSpywareSignaturesSignatureCombination")
	diags := diag.Diagnostics{}
	var data []models.AntiSpywareSignaturesSignatureCombination
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.AntiSpywareSignaturesSignatureCombination, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.AntiSpywareSignaturesSignatureCombination{}.AttrTypes(), &item)
		unpacked, d := unpackAntiSpywareSignaturesSignatureCombinationToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.AntiSpywareSignaturesSignatureCombination", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for AntiSpywareSignaturesSignatureCombination ---
func packAntiSpywareSignaturesSignatureCombinationListFromSdk(ctx context.Context, sdks []security_services.AntiSpywareSignaturesSignatureCombination) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.AntiSpywareSignaturesSignatureCombination")
	diags := diag.Diagnostics{}
	var data []models.AntiSpywareSignaturesSignatureCombination

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.AntiSpywareSignaturesSignatureCombination
		obj, d := packAntiSpywareSignaturesSignatureCombinationFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.AntiSpywareSignaturesSignatureCombination{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.AntiSpywareSignaturesSignatureCombination", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.AntiSpywareSignaturesSignatureCombination{}.AttrType(), data)
}
