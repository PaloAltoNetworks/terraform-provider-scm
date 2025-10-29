package provider

import (
	"context"
	"reflect"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/objects"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/objects"
)

// --- Unpacker for Applications ---
func unpackApplicationsToSdk(ctx context.Context, obj types.Object) (*objects.Applications, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.Applications", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.Applications
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.Applications
	var d diag.Diagnostics

	// Handling Primitives
	if !model.AbleToTransferFile.IsNull() && !model.AbleToTransferFile.IsUnknown() {
		sdk.AbleToTransferFile = model.AbleToTransferFile.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "AbleToTransferFile", "value": *sdk.AbleToTransferFile})
	}

	// Handling Primitives
	if !model.AlgDisableCapability.IsNull() && !model.AlgDisableCapability.IsUnknown() {
		sdk.AlgDisableCapability = model.AlgDisableCapability.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "AlgDisableCapability", "value": *sdk.AlgDisableCapability})
	}

	// Handling Primitives
	if !model.Category.IsNull() && !model.Category.IsUnknown() {
		sdk.Category = model.Category.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Category", "value": sdk.Category})
	}

	// Handling Primitives
	if !model.ConsumeBigBandwidth.IsNull() && !model.ConsumeBigBandwidth.IsUnknown() {
		sdk.ConsumeBigBandwidth = model.ConsumeBigBandwidth.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "ConsumeBigBandwidth", "value": *sdk.ConsumeBigBandwidth})
	}

	// Handling Primitives
	if !model.DataIdent.IsNull() && !model.DataIdent.IsUnknown() {
		sdk.DataIdent = model.DataIdent.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DataIdent", "value": *sdk.DataIdent})
	}

	// Handling Objects
	if !model.Default.IsNull() && !model.Default.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Default")
		unpacked, d := unpackApplicationsDefaultToSdk(ctx, model.Default)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Default"})
		}
		if unpacked != nil {
			sdk.Default = unpacked
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
	if !model.EvasiveBehavior.IsNull() && !model.EvasiveBehavior.IsUnknown() {
		sdk.EvasiveBehavior = model.EvasiveBehavior.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "EvasiveBehavior", "value": *sdk.EvasiveBehavior})
	}

	// Handling Primitives
	if !model.FileTypeIdent.IsNull() && !model.FileTypeIdent.IsUnknown() {
		sdk.FileTypeIdent = model.FileTypeIdent.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "FileTypeIdent", "value": *sdk.FileTypeIdent})
	}

	// Handling Primitives
	if !model.Folder.IsNull() && !model.Folder.IsUnknown() {
		sdk.Folder = model.Folder.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Folder", "value": *sdk.Folder})
	}

	// Handling Primitives
	if !model.HasKnownVulnerability.IsNull() && !model.HasKnownVulnerability.IsUnknown() {
		sdk.HasKnownVulnerability = model.HasKnownVulnerability.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "HasKnownVulnerability", "value": *sdk.HasKnownVulnerability})
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
	if !model.NoAppidCaching.IsNull() && !model.NoAppidCaching.IsUnknown() {
		sdk.NoAppidCaching = model.NoAppidCaching.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "NoAppidCaching", "value": *sdk.NoAppidCaching})
	}

	// Handling Primitives
	if !model.ParentApp.IsNull() && !model.ParentApp.IsUnknown() {
		sdk.ParentApp = model.ParentApp.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "ParentApp", "value": *sdk.ParentApp})
	}

	// Handling Primitives
	if !model.PervasiveUse.IsNull() && !model.PervasiveUse.IsUnknown() {
		sdk.PervasiveUse = model.PervasiveUse.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "PervasiveUse", "value": *sdk.PervasiveUse})
	}

	// Handling Primitives
	if !model.ProneToMisuse.IsNull() && !model.ProneToMisuse.IsUnknown() {
		sdk.ProneToMisuse = model.ProneToMisuse.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "ProneToMisuse", "value": *sdk.ProneToMisuse})
	}

	// Handling Primitives
	if !model.Risk.IsNull() && !model.Risk.IsUnknown() {
		sdk.Risk = model.Risk.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Risk", "value": sdk.Risk})
	}

	// Handling Lists
	if !model.Signature.IsNull() && !model.Signature.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field Signature")
		unpacked, d := unpackApplicationsSignatureInnerListToSdk(ctx, model.Signature)
		diags.Append(d...)
		sdk.Signature = unpacked
	}

	// Handling Primitives
	if !model.Snippet.IsNull() && !model.Snippet.IsUnknown() {
		sdk.Snippet = model.Snippet.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	}

	// Handling Primitives
	if !model.Subcategory.IsNull() && !model.Subcategory.IsUnknown() {
		sdk.Subcategory = model.Subcategory.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Subcategory", "value": *sdk.Subcategory})
	}

	// Handling Primitives
	if !model.TcpHalfClosedTimeout.IsNull() && !model.TcpHalfClosedTimeout.IsUnknown() {
		val := int32(model.TcpHalfClosedTimeout.ValueInt64())
		sdk.TcpHalfClosedTimeout = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "TcpHalfClosedTimeout", "value": *sdk.TcpHalfClosedTimeout})
	}

	// Handling Primitives
	if !model.TcpTimeWaitTimeout.IsNull() && !model.TcpTimeWaitTimeout.IsUnknown() {
		val := int32(model.TcpTimeWaitTimeout.ValueInt64())
		sdk.TcpTimeWaitTimeout = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "TcpTimeWaitTimeout", "value": *sdk.TcpTimeWaitTimeout})
	}

	// Handling Primitives
	if !model.TcpTimeout.IsNull() && !model.TcpTimeout.IsUnknown() {
		val := int32(model.TcpTimeout.ValueInt64())
		sdk.TcpTimeout = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "TcpTimeout", "value": *sdk.TcpTimeout})
	}

	// Handling Primitives
	if !model.Technology.IsNull() && !model.Technology.IsUnknown() {
		sdk.Technology = model.Technology.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Technology", "value": *sdk.Technology})
	}

	// Handling Primitives
	if !model.Timeout.IsNull() && !model.Timeout.IsUnknown() {
		val := int32(model.Timeout.ValueInt64())
		sdk.Timeout = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Timeout", "value": *sdk.Timeout})
	}

	// Handling Primitives
	if !model.TunnelApplications.IsNull() && !model.TunnelApplications.IsUnknown() {
		sdk.TunnelApplications = model.TunnelApplications.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "TunnelApplications", "value": *sdk.TunnelApplications})
	}

	// Handling Primitives
	if !model.TunnelOtherApplication.IsNull() && !model.TunnelOtherApplication.IsUnknown() {
		sdk.TunnelOtherApplication = model.TunnelOtherApplication.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "TunnelOtherApplication", "value": *sdk.TunnelOtherApplication})
	}

	// Handling Primitives
	if !model.UdpTimeout.IsNull() && !model.UdpTimeout.IsUnknown() {
		val := int32(model.UdpTimeout.ValueInt64())
		sdk.UdpTimeout = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "UdpTimeout", "value": *sdk.UdpTimeout})
	}

	// Handling Primitives
	if !model.UsedByMalware.IsNull() && !model.UsedByMalware.IsUnknown() {
		sdk.UsedByMalware = model.UsedByMalware.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "UsedByMalware", "value": *sdk.UsedByMalware})
	}

	// Handling Primitives
	if !model.VirusIdent.IsNull() && !model.VirusIdent.IsUnknown() {
		sdk.VirusIdent = model.VirusIdent.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "VirusIdent", "value": *sdk.VirusIdent})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.Applications", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for Applications ---
func packApplicationsFromSdk(ctx context.Context, sdk objects.Applications) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.Applications", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.Applications
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.AbleToTransferFile != nil {
		model.AbleToTransferFile = basetypes.NewBoolValue(*sdk.AbleToTransferFile)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "AbleToTransferFile", "value": *sdk.AbleToTransferFile})
	} else {
		model.AbleToTransferFile = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.AlgDisableCapability != nil {
		model.AlgDisableCapability = basetypes.NewStringValue(*sdk.AlgDisableCapability)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "AlgDisableCapability", "value": *sdk.AlgDisableCapability})
	} else {
		model.AlgDisableCapability = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	model.Category = basetypes.NewStringValue(sdk.Category)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Category", "value": sdk.Category})
	// Handling Primitives
	// Standard primitive packing
	if sdk.ConsumeBigBandwidth != nil {
		model.ConsumeBigBandwidth = basetypes.NewBoolValue(*sdk.ConsumeBigBandwidth)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "ConsumeBigBandwidth", "value": *sdk.ConsumeBigBandwidth})
	} else {
		model.ConsumeBigBandwidth = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.DataIdent != nil {
		model.DataIdent = basetypes.NewBoolValue(*sdk.DataIdent)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DataIdent", "value": *sdk.DataIdent})
	} else {
		model.DataIdent = basetypes.NewBoolNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Default != nil {
		tflog.Debug(ctx, "Packing nested object for field Default")
		packed, d := packApplicationsDefaultFromSdk(ctx, *sdk.Default)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Default"})
		}
		model.Default = packed
	} else {
		model.Default = basetypes.NewObjectNull(models.ApplicationsDefault{}.AttrTypes())
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
	if sdk.EvasiveBehavior != nil {
		model.EvasiveBehavior = basetypes.NewBoolValue(*sdk.EvasiveBehavior)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "EvasiveBehavior", "value": *sdk.EvasiveBehavior})
	} else {
		model.EvasiveBehavior = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.FileTypeIdent != nil {
		model.FileTypeIdent = basetypes.NewBoolValue(*sdk.FileTypeIdent)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "FileTypeIdent", "value": *sdk.FileTypeIdent})
	} else {
		model.FileTypeIdent = basetypes.NewBoolNull()
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
	if sdk.HasKnownVulnerability != nil {
		model.HasKnownVulnerability = basetypes.NewBoolValue(*sdk.HasKnownVulnerability)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "HasKnownVulnerability", "value": *sdk.HasKnownVulnerability})
	} else {
		model.HasKnownVulnerability = basetypes.NewBoolNull()
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
	if sdk.NoAppidCaching != nil {
		model.NoAppidCaching = basetypes.NewBoolValue(*sdk.NoAppidCaching)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "NoAppidCaching", "value": *sdk.NoAppidCaching})
	} else {
		model.NoAppidCaching = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.ParentApp != nil {
		model.ParentApp = basetypes.NewStringValue(*sdk.ParentApp)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "ParentApp", "value": *sdk.ParentApp})
	} else {
		model.ParentApp = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.PervasiveUse != nil {
		model.PervasiveUse = basetypes.NewBoolValue(*sdk.PervasiveUse)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "PervasiveUse", "value": *sdk.PervasiveUse})
	} else {
		model.PervasiveUse = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.ProneToMisuse != nil {
		model.ProneToMisuse = basetypes.NewBoolValue(*sdk.ProneToMisuse)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "ProneToMisuse", "value": *sdk.ProneToMisuse})
	} else {
		model.ProneToMisuse = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Universal packer for interface{} types that are mapped to a StringValue in the model.
	// All underlying primitive types will be converted to their string representation.
	if sdk.Risk != nil {
		tflog.Debug(ctx, "Packing interface value for field Risk")
		switch v := sdk.Risk.(type) {
		case string:
			model.Risk = basetypes.NewStringValue(v)
		case int:
			model.Risk = basetypes.NewStringValue(strconv.Itoa(v))
		case int64:
			model.Risk = basetypes.NewStringValue(strconv.FormatInt(v, 10))
		case bool:
			model.Risk = basetypes.NewStringValue(strconv.FormatBool(v))
		case float64:
			model.Risk = basetypes.NewStringValue(strconv.FormatFloat(v, 'f', -1, 64))
		default:
			tflog.Warn(ctx, "Unexpected type for interface field", map[string]interface{}{"field": "Risk", "type": reflect.TypeOf(v).String()})
			model.Risk = basetypes.NewStringNull()
		}
	} else {
		model.Risk = basetypes.NewStringNull()
	}
	// Handling Lists
	if sdk.Signature != nil {
		tflog.Debug(ctx, "Packing list of objects for field Signature")
		packed, d := packApplicationsSignatureInnerListFromSdk(ctx, sdk.Signature)
		diags.Append(d...)
		model.Signature = packed
	} else {
		model.Signature = basetypes.NewListNull(models.ApplicationsSignatureInner{}.AttrType())
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
	if sdk.Subcategory != nil {
		model.Subcategory = basetypes.NewStringValue(*sdk.Subcategory)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Subcategory", "value": *sdk.Subcategory})
	} else {
		model.Subcategory = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.TcpHalfClosedTimeout != nil {
		model.TcpHalfClosedTimeout = basetypes.NewInt64Value(int64(*sdk.TcpHalfClosedTimeout))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "TcpHalfClosedTimeout", "value": *sdk.TcpHalfClosedTimeout})
	} else {
		model.TcpHalfClosedTimeout = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.TcpTimeWaitTimeout != nil {
		model.TcpTimeWaitTimeout = basetypes.NewInt64Value(int64(*sdk.TcpTimeWaitTimeout))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "TcpTimeWaitTimeout", "value": *sdk.TcpTimeWaitTimeout})
	} else {
		model.TcpTimeWaitTimeout = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.TcpTimeout != nil {
		model.TcpTimeout = basetypes.NewInt64Value(int64(*sdk.TcpTimeout))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "TcpTimeout", "value": *sdk.TcpTimeout})
	} else {
		model.TcpTimeout = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Technology != nil {
		model.Technology = basetypes.NewStringValue(*sdk.Technology)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Technology", "value": *sdk.Technology})
	} else {
		model.Technology = basetypes.NewStringNull()
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
	if sdk.TunnelApplications != nil {
		model.TunnelApplications = basetypes.NewBoolValue(*sdk.TunnelApplications)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "TunnelApplications", "value": *sdk.TunnelApplications})
	} else {
		model.TunnelApplications = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.TunnelOtherApplication != nil {
		model.TunnelOtherApplication = basetypes.NewBoolValue(*sdk.TunnelOtherApplication)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "TunnelOtherApplication", "value": *sdk.TunnelOtherApplication})
	} else {
		model.TunnelOtherApplication = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.UdpTimeout != nil {
		model.UdpTimeout = basetypes.NewInt64Value(int64(*sdk.UdpTimeout))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "UdpTimeout", "value": *sdk.UdpTimeout})
	} else {
		model.UdpTimeout = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.UsedByMalware != nil {
		model.UsedByMalware = basetypes.NewBoolValue(*sdk.UsedByMalware)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "UsedByMalware", "value": *sdk.UsedByMalware})
	} else {
		model.UsedByMalware = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.VirusIdent != nil {
		model.VirusIdent = basetypes.NewBoolValue(*sdk.VirusIdent)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "VirusIdent", "value": *sdk.VirusIdent})
	} else {
		model.VirusIdent = basetypes.NewBoolNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.Applications{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.Applications", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for Applications ---
func unpackApplicationsListToSdk(ctx context.Context, list types.List) ([]objects.Applications, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.Applications")
	diags := diag.Diagnostics{}
	var data []models.Applications
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.Applications, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.Applications{}.AttrTypes(), &item)
		unpacked, d := unpackApplicationsToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.Applications", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for Applications ---
func packApplicationsListFromSdk(ctx context.Context, sdks []objects.Applications) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.Applications")
	diags := diag.Diagnostics{}
	var data []models.Applications

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.Applications
		obj, d := packApplicationsFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.Applications{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.Applications", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.Applications{}.AttrType(), data)
}

// --- Unpacker for ApplicationsDefault ---
func unpackApplicationsDefaultToSdk(ctx context.Context, obj types.Object) (*objects.ApplicationsDefault, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ApplicationsDefault", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ApplicationsDefault
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.ApplicationsDefault
	var d diag.Diagnostics
	// Handling Objects
	if !model.IdentByIcmp6Type.IsNull() && !model.IdentByIcmp6Type.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field IdentByIcmp6Type")
		unpacked, d := unpackApplicationsDefaultIdentByIcmp6TypeToSdk(ctx, model.IdentByIcmp6Type)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "IdentByIcmp6Type"})
		}
		if unpacked != nil {
			sdk.IdentByIcmp6Type = unpacked
		}
	}

	// Handling Objects
	if !model.IdentByIcmpType.IsNull() && !model.IdentByIcmpType.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field IdentByIcmpType")
		unpacked, d := unpackApplicationsDefaultIdentByIcmp6TypeToSdk(ctx, model.IdentByIcmpType)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "IdentByIcmpType"})
		}
		if unpacked != nil {
			sdk.IdentByIcmpType = unpacked
		}
	}

	// Handling Primitives
	if !model.IdentByIpProtocol.IsNull() && !model.IdentByIpProtocol.IsUnknown() {
		sdk.IdentByIpProtocol = model.IdentByIpProtocol.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "IdentByIpProtocol", "value": *sdk.IdentByIpProtocol})
	}

	// Handling Lists
	if !model.Port.IsNull() && !model.Port.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Port")
		diags.Append(model.Port.ElementsAs(ctx, &sdk.Port, false)...)
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ApplicationsDefault", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ApplicationsDefault ---
func packApplicationsDefaultFromSdk(ctx context.Context, sdk objects.ApplicationsDefault) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ApplicationsDefault", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ApplicationsDefault
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.IdentByIcmp6Type != nil {
		tflog.Debug(ctx, "Packing nested object for field IdentByIcmp6Type")
		packed, d := packApplicationsDefaultIdentByIcmp6TypeFromSdk(ctx, *sdk.IdentByIcmp6Type)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "IdentByIcmp6Type"})
		}
		model.IdentByIcmp6Type = packed
	} else {
		model.IdentByIcmp6Type = basetypes.NewObjectNull(models.ApplicationsDefaultIdentByIcmp6Type{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.IdentByIcmpType != nil {
		tflog.Debug(ctx, "Packing nested object for field IdentByIcmpType")
		packed, d := packApplicationsDefaultIdentByIcmp6TypeFromSdk(ctx, *sdk.IdentByIcmpType)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "IdentByIcmpType"})
		}
		model.IdentByIcmpType = packed
	} else {
		model.IdentByIcmpType = basetypes.NewObjectNull(models.ApplicationsDefaultIdentByIcmp6Type{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.IdentByIpProtocol != nil {
		model.IdentByIpProtocol = basetypes.NewStringValue(*sdk.IdentByIpProtocol)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "IdentByIpProtocol", "value": *sdk.IdentByIpProtocol})
	} else {
		model.IdentByIpProtocol = basetypes.NewStringNull()
	}
	// Handling Lists
	if sdk.Port != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Port")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Port, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Port)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Port = basetypes.NewListNull(elemType)
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ApplicationsDefault{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ApplicationsDefault", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ApplicationsDefault ---
func unpackApplicationsDefaultListToSdk(ctx context.Context, list types.List) ([]objects.ApplicationsDefault, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ApplicationsDefault")
	diags := diag.Diagnostics{}
	var data []models.ApplicationsDefault
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.ApplicationsDefault, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ApplicationsDefault{}.AttrTypes(), &item)
		unpacked, d := unpackApplicationsDefaultToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ApplicationsDefault", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ApplicationsDefault ---
func packApplicationsDefaultListFromSdk(ctx context.Context, sdks []objects.ApplicationsDefault) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ApplicationsDefault")
	diags := diag.Diagnostics{}
	var data []models.ApplicationsDefault

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ApplicationsDefault
		obj, d := packApplicationsDefaultFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ApplicationsDefault{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ApplicationsDefault", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ApplicationsDefault{}.AttrType(), data)
}

// --- Unpacker for ApplicationsDefaultIdentByIcmp6Type ---
func unpackApplicationsDefaultIdentByIcmp6TypeToSdk(ctx context.Context, obj types.Object) (*objects.ApplicationsDefaultIdentByIcmp6Type, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ApplicationsDefaultIdentByIcmp6Type", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ApplicationsDefaultIdentByIcmp6Type
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.ApplicationsDefaultIdentByIcmp6Type
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Code.IsNull() && !model.Code.IsUnknown() {
		sdk.Code = model.Code.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Code", "value": *sdk.Code})
	}

	// Handling Primitives
	if !model.Type.IsNull() && !model.Type.IsUnknown() {
		sdk.Type = model.Type.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Type", "value": sdk.Type})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ApplicationsDefaultIdentByIcmp6Type", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ApplicationsDefaultIdentByIcmp6Type ---
func packApplicationsDefaultIdentByIcmp6TypeFromSdk(ctx context.Context, sdk objects.ApplicationsDefaultIdentByIcmp6Type) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ApplicationsDefaultIdentByIcmp6Type", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ApplicationsDefaultIdentByIcmp6Type
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Code != nil {
		model.Code = basetypes.NewStringValue(*sdk.Code)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Code", "value": *sdk.Code})
	} else {
		model.Code = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	model.Type = basetypes.NewStringValue(sdk.Type)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Type", "value": sdk.Type})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ApplicationsDefaultIdentByIcmp6Type{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ApplicationsDefaultIdentByIcmp6Type", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ApplicationsDefaultIdentByIcmp6Type ---
func unpackApplicationsDefaultIdentByIcmp6TypeListToSdk(ctx context.Context, list types.List) ([]objects.ApplicationsDefaultIdentByIcmp6Type, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ApplicationsDefaultIdentByIcmp6Type")
	diags := diag.Diagnostics{}
	var data []models.ApplicationsDefaultIdentByIcmp6Type
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.ApplicationsDefaultIdentByIcmp6Type, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ApplicationsDefaultIdentByIcmp6Type{}.AttrTypes(), &item)
		unpacked, d := unpackApplicationsDefaultIdentByIcmp6TypeToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ApplicationsDefaultIdentByIcmp6Type", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ApplicationsDefaultIdentByIcmp6Type ---
func packApplicationsDefaultIdentByIcmp6TypeListFromSdk(ctx context.Context, sdks []objects.ApplicationsDefaultIdentByIcmp6Type) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ApplicationsDefaultIdentByIcmp6Type")
	diags := diag.Diagnostics{}
	var data []models.ApplicationsDefaultIdentByIcmp6Type

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ApplicationsDefaultIdentByIcmp6Type
		obj, d := packApplicationsDefaultIdentByIcmp6TypeFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ApplicationsDefaultIdentByIcmp6Type{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ApplicationsDefaultIdentByIcmp6Type", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ApplicationsDefaultIdentByIcmp6Type{}.AttrType(), data)
}

// --- Unpacker for ApplicationsSignatureInner ---
func unpackApplicationsSignatureInnerToSdk(ctx context.Context, obj types.Object) (*objects.ApplicationsSignatureInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ApplicationsSignatureInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ApplicationsSignatureInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.ApplicationsSignatureInner
	var d diag.Diagnostics
	// Handling Lists
	if !model.AndCondition.IsNull() && !model.AndCondition.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field AndCondition")
		unpacked, d := unpackApplicationsSignatureInnerAndConditionInnerListToSdk(ctx, model.AndCondition)
		diags.Append(d...)
		sdk.AndCondition = unpacked
	}

	// Handling Primitives
	if !model.Comment.IsNull() && !model.Comment.IsUnknown() {
		sdk.Comment = model.Comment.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Comment", "value": *sdk.Comment})
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	// Handling Primitives
	if !model.OrderFree.IsNull() && !model.OrderFree.IsUnknown() {
		sdk.OrderFree = model.OrderFree.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "OrderFree", "value": *sdk.OrderFree})
	}

	// Handling Primitives
	if !model.Scope.IsNull() && !model.Scope.IsUnknown() {
		sdk.Scope = model.Scope.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Scope", "value": *sdk.Scope})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ApplicationsSignatureInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ApplicationsSignatureInner ---
func packApplicationsSignatureInnerFromSdk(ctx context.Context, sdk objects.ApplicationsSignatureInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ApplicationsSignatureInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ApplicationsSignatureInner
	var d diag.Diagnostics
	// Handling Lists
	if sdk.AndCondition != nil {
		tflog.Debug(ctx, "Packing list of objects for field AndCondition")
		packed, d := packApplicationsSignatureInnerAndConditionInnerListFromSdk(ctx, sdk.AndCondition)
		diags.Append(d...)
		model.AndCondition = packed
	} else {
		model.AndCondition = basetypes.NewListNull(models.ApplicationsSignatureInnerAndConditionInner{}.AttrType())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Comment != nil {
		model.Comment = basetypes.NewStringValue(*sdk.Comment)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Comment", "value": *sdk.Comment})
	} else {
		model.Comment = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	// Handling Primitives
	// Standard primitive packing
	if sdk.OrderFree != nil {
		model.OrderFree = basetypes.NewBoolValue(*sdk.OrderFree)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "OrderFree", "value": *sdk.OrderFree})
	} else {
		model.OrderFree = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Scope != nil {
		model.Scope = basetypes.NewStringValue(*sdk.Scope)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Scope", "value": *sdk.Scope})
	} else {
		model.Scope = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ApplicationsSignatureInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ApplicationsSignatureInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ApplicationsSignatureInner ---
func unpackApplicationsSignatureInnerListToSdk(ctx context.Context, list types.List) ([]objects.ApplicationsSignatureInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ApplicationsSignatureInner")
	diags := diag.Diagnostics{}
	var data []models.ApplicationsSignatureInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.ApplicationsSignatureInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ApplicationsSignatureInner{}.AttrTypes(), &item)
		unpacked, d := unpackApplicationsSignatureInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ApplicationsSignatureInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ApplicationsSignatureInner ---
func packApplicationsSignatureInnerListFromSdk(ctx context.Context, sdks []objects.ApplicationsSignatureInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ApplicationsSignatureInner")
	diags := diag.Diagnostics{}
	var data []models.ApplicationsSignatureInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ApplicationsSignatureInner
		obj, d := packApplicationsSignatureInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ApplicationsSignatureInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ApplicationsSignatureInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ApplicationsSignatureInner{}.AttrType(), data)
}

// --- Unpacker for ApplicationsSignatureInnerAndConditionInner ---
func unpackApplicationsSignatureInnerAndConditionInnerToSdk(ctx context.Context, obj types.Object) (*objects.ApplicationsSignatureInnerAndConditionInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ApplicationsSignatureInnerAndConditionInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ApplicationsSignatureInnerAndConditionInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.ApplicationsSignatureInnerAndConditionInner
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	// Handling Lists
	if !model.OrCondition.IsNull() && !model.OrCondition.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field OrCondition")
		unpacked, d := unpackApplicationsSignatureInnerAndConditionInnerOrConditionInnerListToSdk(ctx, model.OrCondition)
		diags.Append(d...)
		sdk.OrCondition = unpacked
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ApplicationsSignatureInnerAndConditionInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ApplicationsSignatureInnerAndConditionInner ---
func packApplicationsSignatureInnerAndConditionInnerFromSdk(ctx context.Context, sdk objects.ApplicationsSignatureInnerAndConditionInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ApplicationsSignatureInnerAndConditionInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ApplicationsSignatureInnerAndConditionInner
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	// Handling Lists
	if sdk.OrCondition != nil {
		tflog.Debug(ctx, "Packing list of objects for field OrCondition")
		packed, d := packApplicationsSignatureInnerAndConditionInnerOrConditionInnerListFromSdk(ctx, sdk.OrCondition)
		diags.Append(d...)
		model.OrCondition = packed
	} else {
		model.OrCondition = basetypes.NewListNull(models.ApplicationsSignatureInnerAndConditionInnerOrConditionInner{}.AttrType())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ApplicationsSignatureInnerAndConditionInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ApplicationsSignatureInnerAndConditionInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ApplicationsSignatureInnerAndConditionInner ---
func unpackApplicationsSignatureInnerAndConditionInnerListToSdk(ctx context.Context, list types.List) ([]objects.ApplicationsSignatureInnerAndConditionInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ApplicationsSignatureInnerAndConditionInner")
	diags := diag.Diagnostics{}
	var data []models.ApplicationsSignatureInnerAndConditionInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.ApplicationsSignatureInnerAndConditionInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ApplicationsSignatureInnerAndConditionInner{}.AttrTypes(), &item)
		unpacked, d := unpackApplicationsSignatureInnerAndConditionInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ApplicationsSignatureInnerAndConditionInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ApplicationsSignatureInnerAndConditionInner ---
func packApplicationsSignatureInnerAndConditionInnerListFromSdk(ctx context.Context, sdks []objects.ApplicationsSignatureInnerAndConditionInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ApplicationsSignatureInnerAndConditionInner")
	diags := diag.Diagnostics{}
	var data []models.ApplicationsSignatureInnerAndConditionInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ApplicationsSignatureInnerAndConditionInner
		obj, d := packApplicationsSignatureInnerAndConditionInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ApplicationsSignatureInnerAndConditionInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ApplicationsSignatureInnerAndConditionInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ApplicationsSignatureInnerAndConditionInner{}.AttrType(), data)
}

// --- Unpacker for ApplicationsSignatureInnerAndConditionInnerOrConditionInner ---
func unpackApplicationsSignatureInnerAndConditionInnerOrConditionInnerToSdk(ctx context.Context, obj types.Object) (*objects.ApplicationsSignatureInnerAndConditionInnerOrConditionInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ApplicationsSignatureInnerAndConditionInnerOrConditionInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ApplicationsSignatureInnerAndConditionInnerOrConditionInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.ApplicationsSignatureInnerAndConditionInnerOrConditionInner
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	// Handling Objects
	if !model.Operator.IsNull() && !model.Operator.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Operator")
		unpacked, d := unpackApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorToSdk(ctx, model.Operator)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Operator"})
		}
		if unpacked != nil {
			sdk.Operator = *unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ApplicationsSignatureInnerAndConditionInnerOrConditionInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ApplicationsSignatureInnerAndConditionInnerOrConditionInner ---
func packApplicationsSignatureInnerAndConditionInnerOrConditionInnerFromSdk(ctx context.Context, sdk objects.ApplicationsSignatureInnerAndConditionInnerOrConditionInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ApplicationsSignatureInnerAndConditionInnerOrConditionInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ApplicationsSignatureInnerAndConditionInnerOrConditionInner
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	// Handling Objects
	// This is a regular nested object that has its own packer.
	// Logic for non-pointer / value-type nested objects
	if !reflect.ValueOf(sdk.Operator).IsZero() {
		tflog.Debug(ctx, "Packing nested object for field Operator")
		packed, d := packApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorFromSdk(ctx, sdk.Operator)
		diags.Append(d...)
		model.Operator = packed
	} else {
		model.Operator = basetypes.NewObjectNull(models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperator{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ApplicationsSignatureInnerAndConditionInnerOrConditionInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ApplicationsSignatureInnerAndConditionInnerOrConditionInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ApplicationsSignatureInnerAndConditionInnerOrConditionInner ---
func unpackApplicationsSignatureInnerAndConditionInnerOrConditionInnerListToSdk(ctx context.Context, list types.List) ([]objects.ApplicationsSignatureInnerAndConditionInnerOrConditionInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ApplicationsSignatureInnerAndConditionInnerOrConditionInner")
	diags := diag.Diagnostics{}
	var data []models.ApplicationsSignatureInnerAndConditionInnerOrConditionInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.ApplicationsSignatureInnerAndConditionInnerOrConditionInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ApplicationsSignatureInnerAndConditionInnerOrConditionInner{}.AttrTypes(), &item)
		unpacked, d := unpackApplicationsSignatureInnerAndConditionInnerOrConditionInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ApplicationsSignatureInnerAndConditionInnerOrConditionInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ApplicationsSignatureInnerAndConditionInnerOrConditionInner ---
func packApplicationsSignatureInnerAndConditionInnerOrConditionInnerListFromSdk(ctx context.Context, sdks []objects.ApplicationsSignatureInnerAndConditionInnerOrConditionInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ApplicationsSignatureInnerAndConditionInnerOrConditionInner")
	diags := diag.Diagnostics{}
	var data []models.ApplicationsSignatureInnerAndConditionInnerOrConditionInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ApplicationsSignatureInnerAndConditionInnerOrConditionInner
		obj, d := packApplicationsSignatureInnerAndConditionInnerOrConditionInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ApplicationsSignatureInnerAndConditionInnerOrConditionInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ApplicationsSignatureInnerAndConditionInnerOrConditionInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ApplicationsSignatureInnerAndConditionInnerOrConditionInner{}.AttrType(), data)
}

// --- Unpacker for ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperator ---
func unpackApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorToSdk(ctx context.Context, obj types.Object) (*objects.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperator, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperator", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperator
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperator
	var d diag.Diagnostics
	// Handling Objects
	if !model.EqualTo.IsNull() && !model.EqualTo.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field EqualTo")
		unpacked, d := unpackApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorEqualToToSdk(ctx, model.EqualTo)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "EqualTo"})
		}
		if unpacked != nil {
			sdk.EqualTo = unpacked
		}
	}

	// Handling Objects
	if !model.GreaterThan.IsNull() && !model.GreaterThan.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field GreaterThan")
		unpacked, d := unpackApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThanToSdk(ctx, model.GreaterThan)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "GreaterThan"})
		}
		if unpacked != nil {
			sdk.GreaterThan = unpacked
		}
	}

	// Handling Objects
	if !model.LessThan.IsNull() && !model.LessThan.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field LessThan")
		unpacked, d := unpackApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThanToSdk(ctx, model.LessThan)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "LessThan"})
		}
		if unpacked != nil {
			sdk.LessThan = unpacked
		}
	}

	// Handling Objects
	if !model.PatternMatch.IsNull() && !model.PatternMatch.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field PatternMatch")
		unpacked, d := unpackApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorPatternMatchToSdk(ctx, model.PatternMatch)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "PatternMatch"})
		}
		if unpacked != nil {
			sdk.PatternMatch = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperator", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperator ---
func packApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorFromSdk(ctx context.Context, sdk objects.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperator) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperator", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperator
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.EqualTo != nil {
		tflog.Debug(ctx, "Packing nested object for field EqualTo")
		packed, d := packApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorEqualToFromSdk(ctx, *sdk.EqualTo)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "EqualTo"})
		}
		model.EqualTo = packed
	} else {
		model.EqualTo = basetypes.NewObjectNull(models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorEqualTo{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.GreaterThan != nil {
		tflog.Debug(ctx, "Packing nested object for field GreaterThan")
		packed, d := packApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThanFromSdk(ctx, *sdk.GreaterThan)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "GreaterThan"})
		}
		model.GreaterThan = packed
	} else {
		model.GreaterThan = basetypes.NewObjectNull(models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThan{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.LessThan != nil {
		tflog.Debug(ctx, "Packing nested object for field LessThan")
		packed, d := packApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThanFromSdk(ctx, *sdk.LessThan)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "LessThan"})
		}
		model.LessThan = packed
	} else {
		model.LessThan = basetypes.NewObjectNull(models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThan{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.PatternMatch != nil {
		tflog.Debug(ctx, "Packing nested object for field PatternMatch")
		packed, d := packApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorPatternMatchFromSdk(ctx, *sdk.PatternMatch)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "PatternMatch"})
		}
		model.PatternMatch = packed
	} else {
		model.PatternMatch = basetypes.NewObjectNull(models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorPatternMatch{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperator{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperator", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperator ---
func unpackApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorListToSdk(ctx context.Context, list types.List) ([]objects.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperator, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperator")
	diags := diag.Diagnostics{}
	var data []models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperator
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperator, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperator{}.AttrTypes(), &item)
		unpacked, d := unpackApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperator", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperator ---
func packApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorListFromSdk(ctx context.Context, sdks []objects.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperator) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperator")
	diags := diag.Diagnostics{}
	var data []models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperator

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperator
		obj, d := packApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperator{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperator", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperator{}.AttrType(), data)
}

// --- Unpacker for ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorEqualTo ---
func unpackApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorEqualToToSdk(ctx context.Context, obj types.Object) (*objects.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorEqualTo, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorEqualTo", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorEqualTo
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorEqualTo
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Context.IsNull() && !model.Context.IsUnknown() {
		sdk.Context = model.Context.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Context", "value": sdk.Context})
	}

	// Handling Primitives
	if !model.Mask.IsNull() && !model.Mask.IsUnknown() {
		sdk.Mask = model.Mask.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Mask", "value": *sdk.Mask})
	}

	// Handling Primitives
	if !model.Position.IsNull() && !model.Position.IsUnknown() {
		sdk.Position = model.Position.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Position", "value": *sdk.Position})
	}

	// Handling Primitives
	if !model.Value.IsNull() && !model.Value.IsUnknown() {
		sdk.Value = model.Value.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Value", "value": sdk.Value})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorEqualTo", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorEqualTo ---
func packApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorEqualToFromSdk(ctx context.Context, sdk objects.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorEqualTo) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorEqualTo", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorEqualTo
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.Context = basetypes.NewStringValue(sdk.Context)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Context", "value": sdk.Context})
	// Handling Primitives
	// Standard primitive packing
	if sdk.Mask != nil {
		model.Mask = basetypes.NewStringValue(*sdk.Mask)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Mask", "value": *sdk.Mask})
	} else {
		model.Mask = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Position != nil {
		model.Position = basetypes.NewStringValue(*sdk.Position)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Position", "value": *sdk.Position})
	} else {
		model.Position = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	model.Value = basetypes.NewStringValue(sdk.Value)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Value", "value": sdk.Value})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorEqualTo{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorEqualTo", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorEqualTo ---
func unpackApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorEqualToListToSdk(ctx context.Context, list types.List) ([]objects.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorEqualTo, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorEqualTo")
	diags := diag.Diagnostics{}
	var data []models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorEqualTo
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorEqualTo, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorEqualTo{}.AttrTypes(), &item)
		unpacked, d := unpackApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorEqualToToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorEqualTo", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorEqualTo ---
func packApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorEqualToListFromSdk(ctx context.Context, sdks []objects.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorEqualTo) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorEqualTo")
	diags := diag.Diagnostics{}
	var data []models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorEqualTo

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorEqualTo
		obj, d := packApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorEqualToFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorEqualTo{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorEqualTo", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorEqualTo{}.AttrType(), data)
}

// --- Unpacker for ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThan ---
func unpackApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThanToSdk(ctx context.Context, obj types.Object) (*objects.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThan, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThan", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThan
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThan
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Context.IsNull() && !model.Context.IsUnknown() {
		sdk.Context = model.Context.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Context", "value": sdk.Context})
	}

	// Handling Lists
	if !model.Qualifier.IsNull() && !model.Qualifier.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field Qualifier")
		unpacked, d := unpackApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThanQualifierInnerListToSdk(ctx, model.Qualifier)
		diags.Append(d...)
		sdk.Qualifier = unpacked
	}

	// Handling Primitives
	if !model.Value.IsNull() && !model.Value.IsUnknown() {
		sdk.Value = int32(model.Value.ValueInt64())
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Value", "value": sdk.Value})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThan", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThan ---
func packApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThanFromSdk(ctx context.Context, sdk objects.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThan) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThan", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThan
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.Context = basetypes.NewStringValue(sdk.Context)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Context", "value": sdk.Context})
	// Handling Lists
	if sdk.Qualifier != nil {
		tflog.Debug(ctx, "Packing list of objects for field Qualifier")
		packed, d := packApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThanQualifierInnerListFromSdk(ctx, sdk.Qualifier)
		diags.Append(d...)
		model.Qualifier = packed
	} else {
		model.Qualifier = basetypes.NewListNull(models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThanQualifierInner{}.AttrType())
	}
	// Handling Primitives
	// Standard primitive packing
	model.Value = basetypes.NewInt64Value(int64(sdk.Value))
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Value", "value": sdk.Value})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThan{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThan", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThan ---
func unpackApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThanListToSdk(ctx context.Context, list types.List) ([]objects.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThan, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThan")
	diags := diag.Diagnostics{}
	var data []models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThan
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThan, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThan{}.AttrTypes(), &item)
		unpacked, d := unpackApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThanToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThan", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThan ---
func packApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThanListFromSdk(ctx context.Context, sdks []objects.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThan) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThan")
	diags := diag.Diagnostics{}
	var data []models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThan

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThan
		obj, d := packApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThanFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThan{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThan", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThan{}.AttrType(), data)
}

// --- Unpacker for ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThanQualifierInner ---
func unpackApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThanQualifierInnerToSdk(ctx context.Context, obj types.Object) (*objects.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThanQualifierInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThanQualifierInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThanQualifierInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThanQualifierInner
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	// Handling Primitives
	if !model.Value.IsNull() && !model.Value.IsUnknown() {
		sdk.Value = model.Value.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Value", "value": sdk.Value})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThanQualifierInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThanQualifierInner ---
func packApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThanQualifierInnerFromSdk(ctx context.Context, sdk objects.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThanQualifierInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThanQualifierInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThanQualifierInner
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	// Handling Primitives
	// Standard primitive packing
	model.Value = basetypes.NewStringValue(sdk.Value)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Value", "value": sdk.Value})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThanQualifierInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThanQualifierInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThanQualifierInner ---
func unpackApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThanQualifierInnerListToSdk(ctx context.Context, list types.List) ([]objects.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThanQualifierInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThanQualifierInner")
	diags := diag.Diagnostics{}
	var data []models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThanQualifierInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThanQualifierInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThanQualifierInner{}.AttrTypes(), &item)
		unpacked, d := unpackApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThanQualifierInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThanQualifierInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThanQualifierInner ---
func packApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThanQualifierInnerListFromSdk(ctx context.Context, sdks []objects.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThanQualifierInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThanQualifierInner")
	diags := diag.Diagnostics{}
	var data []models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThanQualifierInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThanQualifierInner
		obj, d := packApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThanQualifierInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThanQualifierInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThanQualifierInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThanQualifierInner{}.AttrType(), data)
}

// --- Unpacker for ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorPatternMatch ---
func unpackApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorPatternMatchToSdk(ctx context.Context, obj types.Object) (*objects.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorPatternMatch, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorPatternMatch", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorPatternMatch
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk objects.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorPatternMatch
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Context.IsNull() && !model.Context.IsUnknown() {
		sdk.Context = model.Context.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Context", "value": sdk.Context})
	}

	// Handling Primitives
	if !model.Pattern.IsNull() && !model.Pattern.IsUnknown() {
		sdk.Pattern = model.Pattern.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Pattern", "value": sdk.Pattern})
	}

	// Handling Lists
	if !model.Qualifier.IsNull() && !model.Qualifier.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field Qualifier")
		unpacked, d := unpackApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThanQualifierInnerListToSdk(ctx, model.Qualifier)
		diags.Append(d...)
		sdk.Qualifier = unpacked
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorPatternMatch", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorPatternMatch ---
func packApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorPatternMatchFromSdk(ctx context.Context, sdk objects.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorPatternMatch) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorPatternMatch", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorPatternMatch
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.Context = basetypes.NewStringValue(sdk.Context)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Context", "value": sdk.Context})
	// Handling Primitives
	// Standard primitive packing
	model.Pattern = basetypes.NewStringValue(sdk.Pattern)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Pattern", "value": sdk.Pattern})
	// Handling Lists
	if sdk.Qualifier != nil {
		tflog.Debug(ctx, "Packing list of objects for field Qualifier")
		packed, d := packApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThanQualifierInnerListFromSdk(ctx, sdk.Qualifier)
		diags.Append(d...)
		model.Qualifier = packed
	} else {
		model.Qualifier = basetypes.NewListNull(models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThanQualifierInner{}.AttrType())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorPatternMatch{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorPatternMatch", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorPatternMatch ---
func unpackApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorPatternMatchListToSdk(ctx context.Context, list types.List) ([]objects.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorPatternMatch, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorPatternMatch")
	diags := diag.Diagnostics{}
	var data []models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorPatternMatch
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]objects.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorPatternMatch, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorPatternMatch{}.AttrTypes(), &item)
		unpacked, d := unpackApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorPatternMatchToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorPatternMatch", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorPatternMatch ---
func packApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorPatternMatchListFromSdk(ctx context.Context, sdks []objects.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorPatternMatch) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorPatternMatch")
	diags := diag.Diagnostics{}
	var data []models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorPatternMatch

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorPatternMatch
		obj, d := packApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorPatternMatchFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorPatternMatch{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorPatternMatch", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorPatternMatch{}.AttrType(), data)
}
