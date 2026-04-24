package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/mobile_agent"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/mobile_agent"
)

// --- Unpacker for ForwardingProfiles ---
func unpackForwardingProfilesToSdk(ctx context.Context, obj types.Object) (*mobile_agent.ForwardingProfiles, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ForwardingProfiles", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ForwardingProfiles
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk mobile_agent.ForwardingProfiles
	var d diag.Diagnostics

	// Handling Primitives
	if !model.DefinitionMethod.IsNull() && !model.DefinitionMethod.IsUnknown() {
		sdk.DefinitionMethod = model.DefinitionMethod.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DefinitionMethod", "value": *sdk.DefinitionMethod})
	}

	// Handling Primitives
	if !model.Description.IsNull() && !model.Description.IsUnknown() {
		sdk.Description = model.Description.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Description", "value": *sdk.Description})
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

	// Handling Objects
	if !model.Type.IsNull() && !model.Type.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Type")
		unpacked, d := unpackForwardingProfilesTypeToSdk(ctx, model.Type)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Type"})
		}
		if unpacked != nil {
			sdk.Type = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ForwardingProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ForwardingProfiles ---
func packForwardingProfilesFromSdk(ctx context.Context, sdk mobile_agent.ForwardingProfiles) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ForwardingProfiles", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ForwardingProfiles
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.DefinitionMethod != nil {
		model.DefinitionMethod = basetypes.NewStringValue(*sdk.DefinitionMethod)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DefinitionMethod", "value": *sdk.DefinitionMethod})
	} else {
		model.DefinitionMethod = basetypes.NewStringNull()
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
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Type != nil {
		tflog.Debug(ctx, "Packing nested object for field Type")
		packed, d := packForwardingProfilesTypeFromSdk(ctx, *sdk.Type)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Type"})
		}
		model.Type = packed
	} else {
		model.Type = basetypes.NewObjectNull(models.ForwardingProfilesType{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ForwardingProfiles{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ForwardingProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ForwardingProfiles ---
func unpackForwardingProfilesListToSdk(ctx context.Context, list types.List) ([]mobile_agent.ForwardingProfiles, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ForwardingProfiles")
	diags := diag.Diagnostics{}
	var data []models.ForwardingProfiles
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]mobile_agent.ForwardingProfiles, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ForwardingProfiles{}.AttrTypes(), &item)
		unpacked, d := unpackForwardingProfilesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ForwardingProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ForwardingProfiles ---
func packForwardingProfilesListFromSdk(ctx context.Context, sdks []mobile_agent.ForwardingProfiles) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ForwardingProfiles")
	diags := diag.Diagnostics{}
	var data []models.ForwardingProfiles

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ForwardingProfiles
		obj, d := packForwardingProfilesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ForwardingProfiles{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ForwardingProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ForwardingProfiles{}.AttrType(), data)
}

// --- Unpacker for ForwardingProfilesType ---
func unpackForwardingProfilesTypeToSdk(ctx context.Context, obj types.Object) (*mobile_agent.ForwardingProfilesType, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ForwardingProfilesType", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ForwardingProfilesType
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk mobile_agent.ForwardingProfilesType
	var d diag.Diagnostics
	// Handling Objects
	if !model.GlobalProtectProxy.IsNull() && !model.GlobalProtectProxy.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field GlobalProtectProxy")
		unpacked, d := unpackForwardingProfileGlobalProtectProxyGlobalProtectProxyToSdk(ctx, model.GlobalProtectProxy)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "GlobalProtectProxy"})
		}
		if unpacked != nil {
			sdk.GlobalProtectProxy = unpacked
		}
	}

	// Handling Objects
	if !model.PacFile.IsNull() && !model.PacFile.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field PacFile")
		unpacked, d := unpackForwardingProfilePacFilePacFileToSdk(ctx, model.PacFile)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "PacFile"})
		}
		if unpacked != nil {
			sdk.PacFile = unpacked
		}
	}

	// Handling Objects
	if !model.ZtnaAgent.IsNull() && !model.ZtnaAgent.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field ZtnaAgent")
		unpacked, d := unpackForwardingProfileZtnaAgentZtnaAgentToSdk(ctx, model.ZtnaAgent)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "ZtnaAgent"})
		}
		if unpacked != nil {
			sdk.ZtnaAgent = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ForwardingProfilesType", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ForwardingProfilesType ---
func packForwardingProfilesTypeFromSdk(ctx context.Context, sdk mobile_agent.ForwardingProfilesType) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ForwardingProfilesType", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ForwardingProfilesType
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.GlobalProtectProxy != nil {
		tflog.Debug(ctx, "Packing nested object for field GlobalProtectProxy")
		packed, d := packForwardingProfileGlobalProtectProxyGlobalProtectProxyFromSdk(ctx, *sdk.GlobalProtectProxy)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "GlobalProtectProxy"})
		}
		model.GlobalProtectProxy = packed
	} else {
		model.GlobalProtectProxy = basetypes.NewObjectNull(models.ForwardingProfileGlobalProtectProxyGlobalProtectProxy{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.PacFile != nil {
		tflog.Debug(ctx, "Packing nested object for field PacFile")
		packed, d := packForwardingProfilePacFilePacFileFromSdk(ctx, *sdk.PacFile)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "PacFile"})
		}
		model.PacFile = packed
	} else {
		model.PacFile = basetypes.NewObjectNull(models.ForwardingProfilePacFilePacFile{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.ZtnaAgent != nil {
		tflog.Debug(ctx, "Packing nested object for field ZtnaAgent")
		packed, d := packForwardingProfileZtnaAgentZtnaAgentFromSdk(ctx, *sdk.ZtnaAgent)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "ZtnaAgent"})
		}
		model.ZtnaAgent = packed
	} else {
		model.ZtnaAgent = basetypes.NewObjectNull(models.ForwardingProfileZtnaAgentZtnaAgent{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ForwardingProfilesType{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ForwardingProfilesType", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ForwardingProfilesType ---
func unpackForwardingProfilesTypeListToSdk(ctx context.Context, list types.List) ([]mobile_agent.ForwardingProfilesType, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ForwardingProfilesType")
	diags := diag.Diagnostics{}
	var data []models.ForwardingProfilesType
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]mobile_agent.ForwardingProfilesType, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ForwardingProfilesType{}.AttrTypes(), &item)
		unpacked, d := unpackForwardingProfilesTypeToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ForwardingProfilesType", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ForwardingProfilesType ---
func packForwardingProfilesTypeListFromSdk(ctx context.Context, sdks []mobile_agent.ForwardingProfilesType) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ForwardingProfilesType")
	diags := diag.Diagnostics{}
	var data []models.ForwardingProfilesType

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ForwardingProfilesType
		obj, d := packForwardingProfilesTypeFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ForwardingProfilesType{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ForwardingProfilesType", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ForwardingProfilesType{}.AttrType(), data)
}

// --- Unpacker for ForwardingProfileGlobalProtectProxyGlobalProtectProxy ---
func unpackForwardingProfileGlobalProtectProxyGlobalProtectProxyToSdk(ctx context.Context, obj types.Object) (*mobile_agent.ForwardingProfileGlobalProtectProxyGlobalProtectProxy, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ForwardingProfileGlobalProtectProxyGlobalProtectProxy", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ForwardingProfileGlobalProtectProxyGlobalProtectProxy
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk mobile_agent.ForwardingProfileGlobalProtectProxyGlobalProtectProxy
	var d diag.Diagnostics
	// Handling Objects
	if !model.BlockRule.IsNull() && !model.BlockRule.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field BlockRule")
		unpacked, d := unpackBlockRuleBasicToSdk(ctx, model.BlockRule)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "BlockRule"})
		}
		if unpacked != nil {
			sdk.BlockRule = unpacked
		}
	}

	// Handling Lists
	if !model.ForwardingRules.IsNull() && !model.ForwardingRules.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field ForwardingRules")
		unpacked, d := unpackForwardingRuleBasicListToSdk(ctx, model.ForwardingRules)
		diags.Append(d...)
		sdk.ForwardingRules = unpacked
	}

	// Handling Primitives
	if !model.PacUpload.IsNull() && !model.PacUpload.IsUnknown() {
		sdk.PacUpload = model.PacUpload.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "PacUpload", "value": *sdk.PacUpload})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ForwardingProfileGlobalProtectProxyGlobalProtectProxy", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ForwardingProfileGlobalProtectProxyGlobalProtectProxy ---
func packForwardingProfileGlobalProtectProxyGlobalProtectProxyFromSdk(ctx context.Context, sdk mobile_agent.ForwardingProfileGlobalProtectProxyGlobalProtectProxy) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ForwardingProfileGlobalProtectProxyGlobalProtectProxy", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ForwardingProfileGlobalProtectProxyGlobalProtectProxy
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.BlockRule != nil {
		tflog.Debug(ctx, "Packing nested object for field BlockRule")
		packed, d := packBlockRuleBasicFromSdk(ctx, *sdk.BlockRule)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "BlockRule"})
		}
		model.BlockRule = packed
	} else {
		model.BlockRule = basetypes.NewObjectNull(models.BlockRuleBasic{}.AttrTypes())
	}
	// Handling Lists
	if sdk.ForwardingRules != nil {
		tflog.Debug(ctx, "Packing list of objects for field ForwardingRules")
		packed, d := packForwardingRuleBasicListFromSdk(ctx, sdk.ForwardingRules)
		diags.Append(d...)
		model.ForwardingRules = packed
	} else {
		model.ForwardingRules = basetypes.NewListNull(models.ForwardingRuleBasic{}.AttrType())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.PacUpload != nil {
		model.PacUpload = basetypes.NewBoolValue(*sdk.PacUpload)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "PacUpload", "value": *sdk.PacUpload})
	} else {
		model.PacUpload = basetypes.NewBoolNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ForwardingProfileGlobalProtectProxyGlobalProtectProxy{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ForwardingProfileGlobalProtectProxyGlobalProtectProxy", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ForwardingProfileGlobalProtectProxyGlobalProtectProxy ---
func unpackForwardingProfileGlobalProtectProxyGlobalProtectProxyListToSdk(ctx context.Context, list types.List) ([]mobile_agent.ForwardingProfileGlobalProtectProxyGlobalProtectProxy, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ForwardingProfileGlobalProtectProxyGlobalProtectProxy")
	diags := diag.Diagnostics{}
	var data []models.ForwardingProfileGlobalProtectProxyGlobalProtectProxy
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]mobile_agent.ForwardingProfileGlobalProtectProxyGlobalProtectProxy, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ForwardingProfileGlobalProtectProxyGlobalProtectProxy{}.AttrTypes(), &item)
		unpacked, d := unpackForwardingProfileGlobalProtectProxyGlobalProtectProxyToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ForwardingProfileGlobalProtectProxyGlobalProtectProxy", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ForwardingProfileGlobalProtectProxyGlobalProtectProxy ---
func packForwardingProfileGlobalProtectProxyGlobalProtectProxyListFromSdk(ctx context.Context, sdks []mobile_agent.ForwardingProfileGlobalProtectProxyGlobalProtectProxy) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ForwardingProfileGlobalProtectProxyGlobalProtectProxy")
	diags := diag.Diagnostics{}
	var data []models.ForwardingProfileGlobalProtectProxyGlobalProtectProxy

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ForwardingProfileGlobalProtectProxyGlobalProtectProxy
		obj, d := packForwardingProfileGlobalProtectProxyGlobalProtectProxyFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ForwardingProfileGlobalProtectProxyGlobalProtectProxy{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ForwardingProfileGlobalProtectProxyGlobalProtectProxy", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ForwardingProfileGlobalProtectProxyGlobalProtectProxy{}.AttrType(), data)
}

// --- Unpacker for BlockRuleBasic ---
func unpackBlockRuleBasicToSdk(ctx context.Context, obj types.Object) (*mobile_agent.BlockRuleBasic, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BlockRuleBasic", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BlockRuleBasic
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk mobile_agent.BlockRuleBasic
	var d diag.Diagnostics
	// Handling Objects
	if !model.AllowTcp.IsNull() && !model.AllowTcp.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field AllowTcp")
		unpacked, d := unpackBlockRuleBasicAllowTcpToSdk(ctx, model.AllowTcp)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "AllowTcp"})
		}
		if unpacked != nil {
			sdk.AllowTcp = unpacked
		}
	}

	// Handling Objects
	if !model.AllowUdp.IsNull() && !model.AllowUdp.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field AllowUdp")
		unpacked, d := unpackBlockRuleBasicAllowUdpToSdk(ctx, model.AllowUdp)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "AllowUdp"})
		}
		if unpacked != nil {
			sdk.AllowUdp = unpacked
		}
	}

	// Handling Primitives
	if !model.Enable.IsNull() && !model.Enable.IsUnknown() {
		sdk.Enable = model.Enable.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Enable", "value": *sdk.Enable})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BlockRuleBasic", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BlockRuleBasic ---
func packBlockRuleBasicFromSdk(ctx context.Context, sdk mobile_agent.BlockRuleBasic) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BlockRuleBasic", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BlockRuleBasic
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.AllowTcp != nil {
		tflog.Debug(ctx, "Packing nested object for field AllowTcp")
		packed, d := packBlockRuleBasicAllowTcpFromSdk(ctx, *sdk.AllowTcp)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "AllowTcp"})
		}
		model.AllowTcp = packed
	} else {
		model.AllowTcp = basetypes.NewObjectNull(models.BlockRuleBasicAllowTcp{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.AllowUdp != nil {
		tflog.Debug(ctx, "Packing nested object for field AllowUdp")
		packed, d := packBlockRuleBasicAllowUdpFromSdk(ctx, *sdk.AllowUdp)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "AllowUdp"})
		}
		model.AllowUdp = packed
	} else {
		model.AllowUdp = basetypes.NewObjectNull(models.BlockRuleBasicAllowUdp{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Enable != nil {
		model.Enable = basetypes.NewBoolValue(*sdk.Enable)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Enable", "value": *sdk.Enable})
	} else {
		model.Enable = basetypes.NewBoolNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BlockRuleBasic{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BlockRuleBasic", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BlockRuleBasic ---
func unpackBlockRuleBasicListToSdk(ctx context.Context, list types.List) ([]mobile_agent.BlockRuleBasic, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BlockRuleBasic")
	diags := diag.Diagnostics{}
	var data []models.BlockRuleBasic
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]mobile_agent.BlockRuleBasic, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BlockRuleBasic{}.AttrTypes(), &item)
		unpacked, d := unpackBlockRuleBasicToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BlockRuleBasic", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BlockRuleBasic ---
func packBlockRuleBasicListFromSdk(ctx context.Context, sdks []mobile_agent.BlockRuleBasic) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BlockRuleBasic")
	diags := diag.Diagnostics{}
	var data []models.BlockRuleBasic

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BlockRuleBasic
		obj, d := packBlockRuleBasicFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BlockRuleBasic{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BlockRuleBasic", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BlockRuleBasic{}.AttrType(), data)
}

// --- Unpacker for BlockRuleBasicAllowTcp ---
func unpackBlockRuleBasicAllowTcpToSdk(ctx context.Context, obj types.Object) (*mobile_agent.BlockRuleBasicAllowTcp, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BlockRuleBasicAllowTcp", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BlockRuleBasicAllowTcp
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk mobile_agent.BlockRuleBasicAllowTcp
	var d diag.Diagnostics
	// Handling Primitives
	if !model.EnableLocations.IsNull() && !model.EnableLocations.IsUnknown() {
		sdk.EnableLocations = model.EnableLocations.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "EnableLocations", "value": *sdk.EnableLocations})
	}

	// Handling Lists
	if !model.Locations.IsNull() && !model.Locations.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Locations")
		diags.Append(model.Locations.ElementsAs(ctx, &sdk.Locations, false)...)
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BlockRuleBasicAllowTcp", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BlockRuleBasicAllowTcp ---
func packBlockRuleBasicAllowTcpFromSdk(ctx context.Context, sdk mobile_agent.BlockRuleBasicAllowTcp) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BlockRuleBasicAllowTcp", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BlockRuleBasicAllowTcp
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.EnableLocations != nil {
		model.EnableLocations = basetypes.NewBoolValue(*sdk.EnableLocations)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "EnableLocations", "value": *sdk.EnableLocations})
	} else {
		model.EnableLocations = basetypes.NewBoolNull()
	}
	// Handling Lists
	if sdk.Locations != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Locations")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Locations, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Locations)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Locations = basetypes.NewListNull(elemType)
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BlockRuleBasicAllowTcp{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BlockRuleBasicAllowTcp", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BlockRuleBasicAllowTcp ---
func unpackBlockRuleBasicAllowTcpListToSdk(ctx context.Context, list types.List) ([]mobile_agent.BlockRuleBasicAllowTcp, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BlockRuleBasicAllowTcp")
	diags := diag.Diagnostics{}
	var data []models.BlockRuleBasicAllowTcp
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]mobile_agent.BlockRuleBasicAllowTcp, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BlockRuleBasicAllowTcp{}.AttrTypes(), &item)
		unpacked, d := unpackBlockRuleBasicAllowTcpToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BlockRuleBasicAllowTcp", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BlockRuleBasicAllowTcp ---
func packBlockRuleBasicAllowTcpListFromSdk(ctx context.Context, sdks []mobile_agent.BlockRuleBasicAllowTcp) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BlockRuleBasicAllowTcp")
	diags := diag.Diagnostics{}
	var data []models.BlockRuleBasicAllowTcp

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BlockRuleBasicAllowTcp
		obj, d := packBlockRuleBasicAllowTcpFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BlockRuleBasicAllowTcp{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BlockRuleBasicAllowTcp", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BlockRuleBasicAllowTcp{}.AttrType(), data)
}

// --- Unpacker for BlockRuleBasicAllowUdp ---
func unpackBlockRuleBasicAllowUdpToSdk(ctx context.Context, obj types.Object) (*mobile_agent.BlockRuleBasicAllowUdp, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BlockRuleBasicAllowUdp", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BlockRuleBasicAllowUdp
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk mobile_agent.BlockRuleBasicAllowUdp
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Destinations.IsNull() && !model.Destinations.IsUnknown() {
		sdk.Destinations = model.Destinations.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Destinations", "value": *sdk.Destinations})
	}

	// Handling Primitives
	if !model.EnableDestinations.IsNull() && !model.EnableDestinations.IsUnknown() {
		sdk.EnableDestinations = model.EnableDestinations.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "EnableDestinations", "value": *sdk.EnableDestinations})
	}

	// Handling Primitives
	if !model.EnableLocations.IsNull() && !model.EnableLocations.IsUnknown() {
		sdk.EnableLocations = model.EnableLocations.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "EnableLocations", "value": *sdk.EnableLocations})
	}

	// Handling Lists
	if !model.Locations.IsNull() && !model.Locations.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Locations")
		diags.Append(model.Locations.ElementsAs(ctx, &sdk.Locations, false)...)
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BlockRuleBasicAllowUdp", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BlockRuleBasicAllowUdp ---
func packBlockRuleBasicAllowUdpFromSdk(ctx context.Context, sdk mobile_agent.BlockRuleBasicAllowUdp) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BlockRuleBasicAllowUdp", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BlockRuleBasicAllowUdp
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Destinations != nil {
		model.Destinations = basetypes.NewStringValue(*sdk.Destinations)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Destinations", "value": *sdk.Destinations})
	} else {
		model.Destinations = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.EnableDestinations != nil {
		model.EnableDestinations = basetypes.NewBoolValue(*sdk.EnableDestinations)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "EnableDestinations", "value": *sdk.EnableDestinations})
	} else {
		model.EnableDestinations = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.EnableLocations != nil {
		model.EnableLocations = basetypes.NewBoolValue(*sdk.EnableLocations)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "EnableLocations", "value": *sdk.EnableLocations})
	} else {
		model.EnableLocations = basetypes.NewBoolNull()
	}
	// Handling Lists
	if sdk.Locations != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Locations")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Locations, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Locations)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Locations = basetypes.NewListNull(elemType)
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BlockRuleBasicAllowUdp{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BlockRuleBasicAllowUdp", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BlockRuleBasicAllowUdp ---
func unpackBlockRuleBasicAllowUdpListToSdk(ctx context.Context, list types.List) ([]mobile_agent.BlockRuleBasicAllowUdp, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BlockRuleBasicAllowUdp")
	diags := diag.Diagnostics{}
	var data []models.BlockRuleBasicAllowUdp
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]mobile_agent.BlockRuleBasicAllowUdp, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BlockRuleBasicAllowUdp{}.AttrTypes(), &item)
		unpacked, d := unpackBlockRuleBasicAllowUdpToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BlockRuleBasicAllowUdp", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BlockRuleBasicAllowUdp ---
func packBlockRuleBasicAllowUdpListFromSdk(ctx context.Context, sdks []mobile_agent.BlockRuleBasicAllowUdp) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BlockRuleBasicAllowUdp")
	diags := diag.Diagnostics{}
	var data []models.BlockRuleBasicAllowUdp

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BlockRuleBasicAllowUdp
		obj, d := packBlockRuleBasicAllowUdpFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BlockRuleBasicAllowUdp{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BlockRuleBasicAllowUdp", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BlockRuleBasicAllowUdp{}.AttrType(), data)
}

// --- Unpacker for ForwardingRuleBasic ---
func unpackForwardingRuleBasicToSdk(ctx context.Context, obj types.Object) (*mobile_agent.ForwardingRuleBasic, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ForwardingRuleBasic", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ForwardingRuleBasic
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk mobile_agent.ForwardingRuleBasic
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Connectivity.IsNull() && !model.Connectivity.IsUnknown() {
		sdk.Connectivity = model.Connectivity.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Connectivity", "value": *sdk.Connectivity})
	}

	// Handling Primitives
	if !model.Destinations.IsNull() && !model.Destinations.IsUnknown() {
		sdk.Destinations = model.Destinations.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Destinations", "value": *sdk.Destinations})
	}

	// Handling Primitives
	if !model.Enabled.IsNull() && !model.Enabled.IsUnknown() {
		sdk.Enabled = model.Enabled.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Enabled", "value": *sdk.Enabled})
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	// Handling Primitives
	if !model.UserLocations.IsNull() && !model.UserLocations.IsUnknown() {
		sdk.UserLocations = model.UserLocations.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "UserLocations", "value": *sdk.UserLocations})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ForwardingRuleBasic", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ForwardingRuleBasic ---
func packForwardingRuleBasicFromSdk(ctx context.Context, sdk mobile_agent.ForwardingRuleBasic) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ForwardingRuleBasic", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ForwardingRuleBasic
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Connectivity != nil {
		model.Connectivity = basetypes.NewStringValue(*sdk.Connectivity)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Connectivity", "value": *sdk.Connectivity})
	} else {
		model.Connectivity = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Destinations != nil {
		model.Destinations = basetypes.NewStringValue(*sdk.Destinations)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Destinations", "value": *sdk.Destinations})
	} else {
		model.Destinations = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Enabled != nil {
		model.Enabled = basetypes.NewBoolValue(*sdk.Enabled)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Enabled", "value": *sdk.Enabled})
	} else {
		model.Enabled = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	// Handling Primitives
	// Standard primitive packing
	if sdk.UserLocations != nil {
		model.UserLocations = basetypes.NewStringValue(*sdk.UserLocations)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "UserLocations", "value": *sdk.UserLocations})
	} else {
		model.UserLocations = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ForwardingRuleBasic{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ForwardingRuleBasic", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ForwardingRuleBasic ---
func unpackForwardingRuleBasicListToSdk(ctx context.Context, list types.List) ([]mobile_agent.ForwardingRuleBasic, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ForwardingRuleBasic")
	diags := diag.Diagnostics{}
	var data []models.ForwardingRuleBasic
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]mobile_agent.ForwardingRuleBasic, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ForwardingRuleBasic{}.AttrTypes(), &item)
		unpacked, d := unpackForwardingRuleBasicToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ForwardingRuleBasic", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ForwardingRuleBasic ---
func packForwardingRuleBasicListFromSdk(ctx context.Context, sdks []mobile_agent.ForwardingRuleBasic) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ForwardingRuleBasic")
	diags := diag.Diagnostics{}
	var data []models.ForwardingRuleBasic

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ForwardingRuleBasic
		obj, d := packForwardingRuleBasicFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ForwardingRuleBasic{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ForwardingRuleBasic", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ForwardingRuleBasic{}.AttrType(), data)
}

// --- Unpacker for ForwardingProfilePacFilePacFile ---
func unpackForwardingProfilePacFilePacFileToSdk(ctx context.Context, obj types.Object) (*mobile_agent.ForwardingProfilePacFilePacFile, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ForwardingProfilePacFilePacFile", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ForwardingProfilePacFilePacFile
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk mobile_agent.ForwardingProfilePacFilePacFile
	var d diag.Diagnostics
	// Handling Objects
	if !model.BlockRule.IsNull() && !model.BlockRule.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field BlockRule")
		unpacked, d := unpackBlockRuleBasicToSdk(ctx, model.BlockRule)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "BlockRule"})
		}
		if unpacked != nil {
			sdk.BlockRule = unpacked
		}
	}

	// Handling Lists
	if !model.ForwardingRules.IsNull() && !model.ForwardingRules.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field ForwardingRules")
		unpacked, d := unpackForwardingRuleBasicListToSdk(ctx, model.ForwardingRules)
		diags.Append(d...)
		sdk.ForwardingRules = unpacked
	}

	// Handling Primitives
	if !model.PacUpload.IsNull() && !model.PacUpload.IsUnknown() {
		sdk.PacUpload = model.PacUpload.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "PacUpload", "value": *sdk.PacUpload})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ForwardingProfilePacFilePacFile", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ForwardingProfilePacFilePacFile ---
func packForwardingProfilePacFilePacFileFromSdk(ctx context.Context, sdk mobile_agent.ForwardingProfilePacFilePacFile) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ForwardingProfilePacFilePacFile", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ForwardingProfilePacFilePacFile
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.BlockRule != nil {
		tflog.Debug(ctx, "Packing nested object for field BlockRule")
		packed, d := packBlockRuleBasicFromSdk(ctx, *sdk.BlockRule)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "BlockRule"})
		}
		model.BlockRule = packed
	} else {
		model.BlockRule = basetypes.NewObjectNull(models.BlockRuleBasic{}.AttrTypes())
	}
	// Handling Lists
	if sdk.ForwardingRules != nil {
		tflog.Debug(ctx, "Packing list of objects for field ForwardingRules")
		packed, d := packForwardingRuleBasicListFromSdk(ctx, sdk.ForwardingRules)
		diags.Append(d...)
		model.ForwardingRules = packed
	} else {
		model.ForwardingRules = basetypes.NewListNull(models.ForwardingRuleBasic{}.AttrType())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.PacUpload != nil {
		model.PacUpload = basetypes.NewBoolValue(*sdk.PacUpload)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "PacUpload", "value": *sdk.PacUpload})
	} else {
		model.PacUpload = basetypes.NewBoolNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ForwardingProfilePacFilePacFile{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ForwardingProfilePacFilePacFile", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ForwardingProfilePacFilePacFile ---
func unpackForwardingProfilePacFilePacFileListToSdk(ctx context.Context, list types.List) ([]mobile_agent.ForwardingProfilePacFilePacFile, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ForwardingProfilePacFilePacFile")
	diags := diag.Diagnostics{}
	var data []models.ForwardingProfilePacFilePacFile
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]mobile_agent.ForwardingProfilePacFilePacFile, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ForwardingProfilePacFilePacFile{}.AttrTypes(), &item)
		unpacked, d := unpackForwardingProfilePacFilePacFileToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ForwardingProfilePacFilePacFile", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ForwardingProfilePacFilePacFile ---
func packForwardingProfilePacFilePacFileListFromSdk(ctx context.Context, sdks []mobile_agent.ForwardingProfilePacFilePacFile) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ForwardingProfilePacFilePacFile")
	diags := diag.Diagnostics{}
	var data []models.ForwardingProfilePacFilePacFile

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ForwardingProfilePacFilePacFile
		obj, d := packForwardingProfilePacFilePacFileFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ForwardingProfilePacFilePacFile{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ForwardingProfilePacFilePacFile", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ForwardingProfilePacFilePacFile{}.AttrType(), data)
}

// --- Unpacker for ForwardingProfileZtnaAgentZtnaAgent ---
func unpackForwardingProfileZtnaAgentZtnaAgentToSdk(ctx context.Context, obj types.Object) (*mobile_agent.ForwardingProfileZtnaAgentZtnaAgent, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ForwardingProfileZtnaAgentZtnaAgent", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ForwardingProfileZtnaAgentZtnaAgent
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk mobile_agent.ForwardingProfileZtnaAgentZtnaAgent
	var d diag.Diagnostics
	// Handling Objects
	if !model.BlockRule.IsNull() && !model.BlockRule.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field BlockRule")
		unpacked, d := unpackBlockRuleZtnaToSdk(ctx, model.BlockRule)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "BlockRule"})
		}
		if unpacked != nil {
			sdk.BlockRule = unpacked
		}
	}

	// Handling Lists
	if !model.ForwardingRules.IsNull() && !model.ForwardingRules.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field ForwardingRules")
		unpacked, d := unpackForwardingRuleZtnaListToSdk(ctx, model.ForwardingRules)
		diags.Append(d...)
		sdk.ForwardingRules = unpacked
	}

	// Handling Primitives
	if !model.PacUpload.IsNull() && !model.PacUpload.IsUnknown() {
		sdk.PacUpload = model.PacUpload.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "PacUpload", "value": *sdk.PacUpload})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ForwardingProfileZtnaAgentZtnaAgent", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ForwardingProfileZtnaAgentZtnaAgent ---
func packForwardingProfileZtnaAgentZtnaAgentFromSdk(ctx context.Context, sdk mobile_agent.ForwardingProfileZtnaAgentZtnaAgent) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ForwardingProfileZtnaAgentZtnaAgent", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ForwardingProfileZtnaAgentZtnaAgent
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.BlockRule != nil {
		tflog.Debug(ctx, "Packing nested object for field BlockRule")
		packed, d := packBlockRuleZtnaFromSdk(ctx, *sdk.BlockRule)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "BlockRule"})
		}
		model.BlockRule = packed
	} else {
		model.BlockRule = basetypes.NewObjectNull(models.BlockRuleZtna{}.AttrTypes())
	}
	// Handling Lists
	if sdk.ForwardingRules != nil {
		tflog.Debug(ctx, "Packing list of objects for field ForwardingRules")
		packed, d := packForwardingRuleZtnaListFromSdk(ctx, sdk.ForwardingRules)
		diags.Append(d...)
		model.ForwardingRules = packed
	} else {
		model.ForwardingRules = basetypes.NewListNull(models.ForwardingRuleZtna{}.AttrType())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.PacUpload != nil {
		model.PacUpload = basetypes.NewBoolValue(*sdk.PacUpload)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "PacUpload", "value": *sdk.PacUpload})
	} else {
		model.PacUpload = basetypes.NewBoolNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ForwardingProfileZtnaAgentZtnaAgent{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ForwardingProfileZtnaAgentZtnaAgent", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ForwardingProfileZtnaAgentZtnaAgent ---
func unpackForwardingProfileZtnaAgentZtnaAgentListToSdk(ctx context.Context, list types.List) ([]mobile_agent.ForwardingProfileZtnaAgentZtnaAgent, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ForwardingProfileZtnaAgentZtnaAgent")
	diags := diag.Diagnostics{}
	var data []models.ForwardingProfileZtnaAgentZtnaAgent
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]mobile_agent.ForwardingProfileZtnaAgentZtnaAgent, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ForwardingProfileZtnaAgentZtnaAgent{}.AttrTypes(), &item)
		unpacked, d := unpackForwardingProfileZtnaAgentZtnaAgentToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ForwardingProfileZtnaAgentZtnaAgent", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ForwardingProfileZtnaAgentZtnaAgent ---
func packForwardingProfileZtnaAgentZtnaAgentListFromSdk(ctx context.Context, sdks []mobile_agent.ForwardingProfileZtnaAgentZtnaAgent) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ForwardingProfileZtnaAgentZtnaAgent")
	diags := diag.Diagnostics{}
	var data []models.ForwardingProfileZtnaAgentZtnaAgent

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ForwardingProfileZtnaAgentZtnaAgent
		obj, d := packForwardingProfileZtnaAgentZtnaAgentFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ForwardingProfileZtnaAgentZtnaAgent{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ForwardingProfileZtnaAgentZtnaAgent", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ForwardingProfileZtnaAgentZtnaAgent{}.AttrType(), data)
}

// --- Unpacker for BlockRuleZtna ---
func unpackBlockRuleZtnaToSdk(ctx context.Context, obj types.Object) (*mobile_agent.BlockRuleZtna, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.BlockRuleZtna", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.BlockRuleZtna
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk mobile_agent.BlockRuleZtna
	var d diag.Diagnostics
	// Handling Primitives
	if !model.AllowIcmpForTroubleshooting.IsNull() && !model.AllowIcmpForTroubleshooting.IsUnknown() {
		sdk.AllowIcmpForTroubleshooting = model.AllowIcmpForTroubleshooting.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "AllowIcmpForTroubleshooting", "value": *sdk.AllowIcmpForTroubleshooting})
	}

	// Handling Primitives
	if !model.BlockAllOtherUnmatchedOutboundConnections.IsNull() && !model.BlockAllOtherUnmatchedOutboundConnections.IsUnknown() {
		sdk.BlockAllOtherUnmatchedOutboundConnections = model.BlockAllOtherUnmatchedOutboundConnections.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "BlockAllOtherUnmatchedOutboundConnections", "value": *sdk.BlockAllOtherUnmatchedOutboundConnections})
	}

	// Handling Primitives
	if !model.BlockInboundAccessWhenConnectedToTunnel.IsNull() && !model.BlockInboundAccessWhenConnectedToTunnel.IsUnknown() {
		sdk.BlockInboundAccessWhenConnectedToTunnel = model.BlockInboundAccessWhenConnectedToTunnel.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "BlockInboundAccessWhenConnectedToTunnel", "value": *sdk.BlockInboundAccessWhenConnectedToTunnel})
	}

	// Handling Primitives
	if !model.BlockNonTcpNonUdpTrafficWhenConnectedToTunnel.IsNull() && !model.BlockNonTcpNonUdpTrafficWhenConnectedToTunnel.IsUnknown() {
		sdk.BlockNonTcpNonUdpTrafficWhenConnectedToTunnel = model.BlockNonTcpNonUdpTrafficWhenConnectedToTunnel.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "BlockNonTcpNonUdpTrafficWhenConnectedToTunnel", "value": *sdk.BlockNonTcpNonUdpTrafficWhenConnectedToTunnel})
	}

	// Handling Primitives
	if !model.BlockOutboundLanAccessWhenConnectedToTunnel.IsNull() && !model.BlockOutboundLanAccessWhenConnectedToTunnel.IsUnknown() {
		sdk.BlockOutboundLanAccessWhenConnectedToTunnel = model.BlockOutboundLanAccessWhenConnectedToTunnel.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "BlockOutboundLanAccessWhenConnectedToTunnel", "value": *sdk.BlockOutboundLanAccessWhenConnectedToTunnel})
	}

	// Handling Primitives
	if !model.EnforcerFqdnDnsResolutionViaDnsServers.IsNull() && !model.EnforcerFqdnDnsResolutionViaDnsServers.IsUnknown() {
		sdk.EnforcerFqdnDnsResolutionViaDnsServers = model.EnforcerFqdnDnsResolutionViaDnsServers.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "EnforcerFqdnDnsResolutionViaDnsServers", "value": *sdk.EnforcerFqdnDnsResolutionViaDnsServers})
	}

	// Handling Primitives
	if !model.ResolveAllFqdnsUsingDnsServersAssignedByTheTunnel.IsNull() && !model.ResolveAllFqdnsUsingDnsServersAssignedByTheTunnel.IsUnknown() {
		sdk.ResolveAllFqdnsUsingDnsServersAssignedByTheTunnel = model.ResolveAllFqdnsUsingDnsServersAssignedByTheTunnel.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "ResolveAllFqdnsUsingDnsServersAssignedByTheTunnel", "value": *sdk.ResolveAllFqdnsUsingDnsServersAssignedByTheTunnel})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.BlockRuleZtna", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for BlockRuleZtna ---
func packBlockRuleZtnaFromSdk(ctx context.Context, sdk mobile_agent.BlockRuleZtna) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.BlockRuleZtna", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.BlockRuleZtna
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.AllowIcmpForTroubleshooting != nil {
		model.AllowIcmpForTroubleshooting = basetypes.NewBoolValue(*sdk.AllowIcmpForTroubleshooting)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "AllowIcmpForTroubleshooting", "value": *sdk.AllowIcmpForTroubleshooting})
	} else {
		model.AllowIcmpForTroubleshooting = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.BlockAllOtherUnmatchedOutboundConnections != nil {
		model.BlockAllOtherUnmatchedOutboundConnections = basetypes.NewBoolValue(*sdk.BlockAllOtherUnmatchedOutboundConnections)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "BlockAllOtherUnmatchedOutboundConnections", "value": *sdk.BlockAllOtherUnmatchedOutboundConnections})
	} else {
		model.BlockAllOtherUnmatchedOutboundConnections = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.BlockInboundAccessWhenConnectedToTunnel != nil {
		model.BlockInboundAccessWhenConnectedToTunnel = basetypes.NewBoolValue(*sdk.BlockInboundAccessWhenConnectedToTunnel)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "BlockInboundAccessWhenConnectedToTunnel", "value": *sdk.BlockInboundAccessWhenConnectedToTunnel})
	} else {
		model.BlockInboundAccessWhenConnectedToTunnel = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.BlockNonTcpNonUdpTrafficWhenConnectedToTunnel != nil {
		model.BlockNonTcpNonUdpTrafficWhenConnectedToTunnel = basetypes.NewBoolValue(*sdk.BlockNonTcpNonUdpTrafficWhenConnectedToTunnel)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "BlockNonTcpNonUdpTrafficWhenConnectedToTunnel", "value": *sdk.BlockNonTcpNonUdpTrafficWhenConnectedToTunnel})
	} else {
		model.BlockNonTcpNonUdpTrafficWhenConnectedToTunnel = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.BlockOutboundLanAccessWhenConnectedToTunnel != nil {
		model.BlockOutboundLanAccessWhenConnectedToTunnel = basetypes.NewBoolValue(*sdk.BlockOutboundLanAccessWhenConnectedToTunnel)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "BlockOutboundLanAccessWhenConnectedToTunnel", "value": *sdk.BlockOutboundLanAccessWhenConnectedToTunnel})
	} else {
		model.BlockOutboundLanAccessWhenConnectedToTunnel = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.EnforcerFqdnDnsResolutionViaDnsServers != nil {
		model.EnforcerFqdnDnsResolutionViaDnsServers = basetypes.NewBoolValue(*sdk.EnforcerFqdnDnsResolutionViaDnsServers)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "EnforcerFqdnDnsResolutionViaDnsServers", "value": *sdk.EnforcerFqdnDnsResolutionViaDnsServers})
	} else {
		model.EnforcerFqdnDnsResolutionViaDnsServers = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.ResolveAllFqdnsUsingDnsServersAssignedByTheTunnel != nil {
		model.ResolveAllFqdnsUsingDnsServersAssignedByTheTunnel = basetypes.NewBoolValue(*sdk.ResolveAllFqdnsUsingDnsServersAssignedByTheTunnel)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "ResolveAllFqdnsUsingDnsServersAssignedByTheTunnel", "value": *sdk.ResolveAllFqdnsUsingDnsServersAssignedByTheTunnel})
	} else {
		model.ResolveAllFqdnsUsingDnsServersAssignedByTheTunnel = basetypes.NewBoolNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.BlockRuleZtna{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.BlockRuleZtna", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for BlockRuleZtna ---
func unpackBlockRuleZtnaListToSdk(ctx context.Context, list types.List) ([]mobile_agent.BlockRuleZtna, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.BlockRuleZtna")
	diags := diag.Diagnostics{}
	var data []models.BlockRuleZtna
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]mobile_agent.BlockRuleZtna, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.BlockRuleZtna{}.AttrTypes(), &item)
		unpacked, d := unpackBlockRuleZtnaToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.BlockRuleZtna", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for BlockRuleZtna ---
func packBlockRuleZtnaListFromSdk(ctx context.Context, sdks []mobile_agent.BlockRuleZtna) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.BlockRuleZtna")
	diags := diag.Diagnostics{}
	var data []models.BlockRuleZtna

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.BlockRuleZtna
		obj, d := packBlockRuleZtnaFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.BlockRuleZtna{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.BlockRuleZtna", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.BlockRuleZtna{}.AttrType(), data)
}

// --- Unpacker for ForwardingRuleZtna ---
func unpackForwardingRuleZtnaToSdk(ctx context.Context, obj types.Object) (*mobile_agent.ForwardingRuleZtna, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ForwardingRuleZtna", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ForwardingRuleZtna
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk mobile_agent.ForwardingRuleZtna
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Connectivity.IsNull() && !model.Connectivity.IsUnknown() {
		sdk.Connectivity = model.Connectivity.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Connectivity", "value": *sdk.Connectivity})
	}

	// Handling Primitives
	if !model.Destinations.IsNull() && !model.Destinations.IsUnknown() {
		sdk.Destinations = model.Destinations.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Destinations", "value": *sdk.Destinations})
	}

	// Handling Primitives
	if !model.Enabled.IsNull() && !model.Enabled.IsUnknown() {
		sdk.Enabled = model.Enabled.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Enabled", "value": *sdk.Enabled})
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	// Handling Primitives
	if !model.SourceApplications.IsNull() && !model.SourceApplications.IsUnknown() {
		sdk.SourceApplications = model.SourceApplications.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "SourceApplications", "value": *sdk.SourceApplications})
	}

	// Handling Primitives
	if !model.TrafficType.IsNull() && !model.TrafficType.IsUnknown() {
		sdk.TrafficType = model.TrafficType.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "TrafficType", "value": *sdk.TrafficType})
	}

	// Handling Primitives
	if !model.UserLocations.IsNull() && !model.UserLocations.IsUnknown() {
		sdk.UserLocations = model.UserLocations.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "UserLocations", "value": *sdk.UserLocations})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ForwardingRuleZtna", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ForwardingRuleZtna ---
func packForwardingRuleZtnaFromSdk(ctx context.Context, sdk mobile_agent.ForwardingRuleZtna) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ForwardingRuleZtna", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ForwardingRuleZtna
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Connectivity != nil {
		model.Connectivity = basetypes.NewStringValue(*sdk.Connectivity)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Connectivity", "value": *sdk.Connectivity})
	} else {
		model.Connectivity = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Destinations != nil {
		model.Destinations = basetypes.NewStringValue(*sdk.Destinations)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Destinations", "value": *sdk.Destinations})
	} else {
		model.Destinations = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Enabled != nil {
		model.Enabled = basetypes.NewBoolValue(*sdk.Enabled)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Enabled", "value": *sdk.Enabled})
	} else {
		model.Enabled = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	// Handling Primitives
	// Standard primitive packing
	if sdk.SourceApplications != nil {
		model.SourceApplications = basetypes.NewStringValue(*sdk.SourceApplications)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "SourceApplications", "value": *sdk.SourceApplications})
	} else {
		model.SourceApplications = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.TrafficType != nil {
		model.TrafficType = basetypes.NewStringValue(*sdk.TrafficType)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "TrafficType", "value": *sdk.TrafficType})
	} else {
		model.TrafficType = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.UserLocations != nil {
		model.UserLocations = basetypes.NewStringValue(*sdk.UserLocations)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "UserLocations", "value": *sdk.UserLocations})
	} else {
		model.UserLocations = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ForwardingRuleZtna{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ForwardingRuleZtna", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ForwardingRuleZtna ---
func unpackForwardingRuleZtnaListToSdk(ctx context.Context, list types.List) ([]mobile_agent.ForwardingRuleZtna, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ForwardingRuleZtna")
	diags := diag.Diagnostics{}
	var data []models.ForwardingRuleZtna
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]mobile_agent.ForwardingRuleZtna, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ForwardingRuleZtna{}.AttrTypes(), &item)
		unpacked, d := unpackForwardingRuleZtnaToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ForwardingRuleZtna", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ForwardingRuleZtna ---
func packForwardingRuleZtnaListFromSdk(ctx context.Context, sdks []mobile_agent.ForwardingRuleZtna) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ForwardingRuleZtna")
	diags := diag.Diagnostics{}
	var data []models.ForwardingRuleZtna

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ForwardingRuleZtna
		obj, d := packForwardingRuleZtnaFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ForwardingRuleZtna{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ForwardingRuleZtna", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ForwardingRuleZtna{}.AttrType(), data)
}
