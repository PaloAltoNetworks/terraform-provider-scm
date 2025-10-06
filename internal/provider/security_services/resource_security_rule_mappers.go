package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/security_services"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/security_services"
)

// --- Unpacker for SecurityRules ---
func unpackSecurityRulesToSdk(ctx context.Context, obj types.Object) (*security_services.SecurityRules, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.SecurityRules", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.SecurityRules
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.SecurityRules
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Action.IsNull() && !model.Action.IsUnknown() {
		sdk.Action = model.Action.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Action", "value": *sdk.Action})
	}

	// Handling Lists
	if !model.AllowUrlCategory.IsNull() && !model.AllowUrlCategory.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field AllowUrlCategory")
		unpacked, d := unpackInternetRuleTypeAllowUrlCategoryInnerListToSdk(ctx, model.AllowUrlCategory)
		diags.Append(d...)
		sdk.AllowUrlCategory = unpacked
	}

	// Handling Lists
	if !model.AllowWebApplication.IsNull() && !model.AllowWebApplication.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field AllowWebApplication")
		unpacked, d := unpackInternetRuleTypeAllowWebApplicationInnerListToSdk(ctx, model.AllowWebApplication)
		diags.Append(d...)
		sdk.AllowWebApplication = unpacked
	}

	// Handling Lists
	if !model.Application.IsNull() && !model.Application.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Application")
		diags.Append(model.Application.ElementsAs(ctx, &sdk.Application, false)...)
	}

	// Handling Lists
	if !model.BlockUrlCategory.IsNull() && !model.BlockUrlCategory.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field BlockUrlCategory")
		diags.Append(model.BlockUrlCategory.ElementsAs(ctx, &sdk.BlockUrlCategory, false)...)
	}

	// Handling Lists
	if !model.BlockWebApplication.IsNull() && !model.BlockWebApplication.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field BlockWebApplication")
		diags.Append(model.BlockWebApplication.ElementsAs(ctx, &sdk.BlockWebApplication, false)...)
	}

	// Handling Lists
	if !model.Category.IsNull() && !model.Category.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Category")
		diags.Append(model.Category.ElementsAs(ctx, &sdk.Category, false)...)
	}

	// Handling Objects
	if !model.DefaultProfileSettings.IsNull() && !model.DefaultProfileSettings.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field DefaultProfileSettings")
		unpacked, d := unpackInternetRuleTypeDefaultProfileSettingsToSdk(ctx, model.DefaultProfileSettings)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "DefaultProfileSettings"})
		}
		if unpacked != nil {
			sdk.DefaultProfileSettings = unpacked
		}
	}

	// Handling Primitives
	if !model.Description.IsNull() && !model.Description.IsUnknown() {
		sdk.Description = model.Description.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Description", "value": *sdk.Description})
	}

	// Handling Lists
	if !model.Destination.IsNull() && !model.Destination.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Destination")
		diags.Append(model.Destination.ElementsAs(ctx, &sdk.Destination, false)...)
	}

	// Handling Lists
	if !model.DestinationHip.IsNull() && !model.DestinationHip.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field DestinationHip")
		diags.Append(model.DestinationHip.ElementsAs(ctx, &sdk.DestinationHip, false)...)
	}

	// Handling Primitives
	if !model.Device.IsNull() && !model.Device.IsUnknown() {
		sdk.Device = model.Device.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Device", "value": *sdk.Device})
	}

	// Handling Lists
	if !model.Devices.IsNull() && !model.Devices.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Devices")
		diags.Append(model.Devices.ElementsAs(ctx, &sdk.Devices, false)...)
	}

	// Handling Primitives
	if !model.Disabled.IsNull() && !model.Disabled.IsUnknown() {
		sdk.Disabled = model.Disabled.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Disabled", "value": *sdk.Disabled})
	}

	// Handling Primitives
	if !model.Folder.IsNull() && !model.Folder.IsUnknown() {
		sdk.Folder = model.Folder.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Folder", "value": *sdk.Folder})
	}

	// Handling Lists
	if !model.From.IsNull() && !model.From.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field From")
		diags.Append(model.From.ElementsAs(ctx, &sdk.From, false)...)
	}

	// Handling Primitives
	if !model.Id.IsNull() && !model.Id.IsUnknown() {
		sdk.Id = model.Id.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Id", "value": *sdk.Id})
	}

	// Handling Primitives
	if !model.LogEnd.IsNull() && !model.LogEnd.IsUnknown() {
		sdk.LogEnd = model.LogEnd.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "LogEnd", "value": *sdk.LogEnd})
	}

	// Handling Primitives
	if !model.LogSetting.IsNull() && !model.LogSetting.IsUnknown() {
		sdk.LogSetting = model.LogSetting.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "LogSetting", "value": *sdk.LogSetting})
	}

	// Handling Objects
	if !model.LogSettings.IsNull() && !model.LogSettings.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field LogSettings")
		unpacked, d := unpackInternetRuleTypeLogSettingsToSdk(ctx, model.LogSettings)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "LogSettings"})
		}
		if unpacked != nil {
			sdk.LogSettings = unpacked
		}
	}

	// Handling Primitives
	if !model.LogStart.IsNull() && !model.LogStart.IsUnknown() {
		sdk.LogStart = model.LogStart.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "LogStart", "value": *sdk.LogStart})
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
	}

	// Handling Primitives
	if !model.NegateDestination.IsNull() && !model.NegateDestination.IsUnknown() {
		sdk.NegateDestination = model.NegateDestination.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "NegateDestination", "value": *sdk.NegateDestination})
	}

	// Handling Primitives
	if !model.NegateSource.IsNull() && !model.NegateSource.IsUnknown() {
		sdk.NegateSource = model.NegateSource.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "NegateSource", "value": *sdk.NegateSource})
	}

	// Handling Primitives
	if !model.NegateUser.IsNull() && !model.NegateUser.IsUnknown() {
		sdk.NegateUser = model.NegateUser.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "NegateUser", "value": *sdk.NegateUser})
	}

	// Handling Primitives
	if !model.PolicyType.IsNull() && !model.PolicyType.IsUnknown() {
		sdk.PolicyType = model.PolicyType.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "PolicyType", "value": *sdk.PolicyType})
	}

	// Handling Objects
	if !model.ProfileSetting.IsNull() && !model.ProfileSetting.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field ProfileSetting")
		unpacked, d := unpackSecurityRuleTypeProfileSettingToSdk(ctx, model.ProfileSetting)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "ProfileSetting"})
		}
		if unpacked != nil {
			sdk.ProfileSetting = unpacked
		}
	}

	// Handling Primitives
	if !model.Schedule.IsNull() && !model.Schedule.IsUnknown() {
		sdk.Schedule = model.Schedule.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Schedule", "value": *sdk.Schedule})
	}

	// Handling Objects
	if !model.SecuritySettings.IsNull() && !model.SecuritySettings.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field SecuritySettings")
		unpacked, d := unpackInternetRuleTypeSecuritySettingsToSdk(ctx, model.SecuritySettings)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "SecuritySettings"})
		}
		if unpacked != nil {
			sdk.SecuritySettings = unpacked
		}
	}

	// Handling Lists
	if !model.Service.IsNull() && !model.Service.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Service")
		diags.Append(model.Service.ElementsAs(ctx, &sdk.Service, false)...)
	}

	// Handling Primitives
	if !model.Snippet.IsNull() && !model.Snippet.IsUnknown() {
		sdk.Snippet = model.Snippet.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	}

	// Handling Lists
	if !model.Source.IsNull() && !model.Source.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Source")
		diags.Append(model.Source.ElementsAs(ctx, &sdk.Source, false)...)
	}

	// Handling Lists
	if !model.SourceHip.IsNull() && !model.SourceHip.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field SourceHip")
		diags.Append(model.SourceHip.ElementsAs(ctx, &sdk.SourceHip, false)...)
	}

	// Handling Lists
	if !model.SourceUser.IsNull() && !model.SourceUser.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field SourceUser")
		diags.Append(model.SourceUser.ElementsAs(ctx, &sdk.SourceUser, false)...)
	}

	// Handling Lists
	if !model.Tag.IsNull() && !model.Tag.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Tag")
		diags.Append(model.Tag.ElementsAs(ctx, &sdk.Tag, false)...)
	}

	// Handling Lists
	if !model.TenantRestrictions.IsNull() && !model.TenantRestrictions.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field TenantRestrictions")
		diags.Append(model.TenantRestrictions.ElementsAs(ctx, &sdk.TenantRestrictions, false)...)
	}

	// Handling Lists
	if !model.To.IsNull() && !model.To.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field To")
		diags.Append(model.To.ElementsAs(ctx, &sdk.To, false)...)
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.SecurityRules", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for SecurityRules ---
func packSecurityRulesFromSdk(ctx context.Context, sdk security_services.SecurityRules) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.SecurityRules", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.SecurityRules
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Action != nil {
		model.Action = basetypes.NewStringValue(*sdk.Action)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Action", "value": *sdk.Action})
	} else {
		model.Action = basetypes.NewStringNull()
	}
	// Handling Lists
	if sdk.AllowUrlCategory != nil {
		tflog.Debug(ctx, "Packing list of objects for field AllowUrlCategory")
		packed, d := packInternetRuleTypeAllowUrlCategoryInnerListFromSdk(ctx, sdk.AllowUrlCategory)
		diags.Append(d...)
		model.AllowUrlCategory = packed
	} else {
		model.AllowUrlCategory = basetypes.NewListNull(models.InternetRuleTypeAllowUrlCategoryInner{}.AttrType())
	}
	// Handling Lists
	if sdk.AllowWebApplication != nil {
		tflog.Debug(ctx, "Packing list of objects for field AllowWebApplication")
		packed, d := packInternetRuleTypeAllowWebApplicationInnerListFromSdk(ctx, sdk.AllowWebApplication)
		diags.Append(d...)
		model.AllowWebApplication = packed
	} else {
		model.AllowWebApplication = basetypes.NewListNull(models.InternetRuleTypeAllowWebApplicationInner{}.AttrType())
	}
	// Handling Lists
	if sdk.Application != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Application")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Application, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Application)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Application = basetypes.NewListNull(elemType)
	}
	// Handling Lists
	if sdk.BlockUrlCategory != nil {
		tflog.Debug(ctx, "Packing list of primitives for field BlockUrlCategory")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.BlockUrlCategory, d = basetypes.NewListValueFrom(ctx, elemType, sdk.BlockUrlCategory)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.BlockUrlCategory = basetypes.NewListNull(elemType)
	}
	// Handling Lists
	if sdk.BlockWebApplication != nil {
		tflog.Debug(ctx, "Packing list of primitives for field BlockWebApplication")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.BlockWebApplication, d = basetypes.NewListValueFrom(ctx, elemType, sdk.BlockWebApplication)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.BlockWebApplication = basetypes.NewListNull(elemType)
	}
	// Handling Lists
	if sdk.Category != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Category")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Category, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Category)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Category = basetypes.NewListNull(elemType)
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.DefaultProfileSettings != nil {
		tflog.Debug(ctx, "Packing nested object for field DefaultProfileSettings")
		packed, d := packInternetRuleTypeDefaultProfileSettingsFromSdk(ctx, *sdk.DefaultProfileSettings)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "DefaultProfileSettings"})
		}
		model.DefaultProfileSettings = packed
	} else {
		model.DefaultProfileSettings = basetypes.NewObjectNull(models.InternetRuleTypeDefaultProfileSettings{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Description != nil {
		model.Description = basetypes.NewStringValue(*sdk.Description)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Description", "value": *sdk.Description})
	} else {
		model.Description = basetypes.NewStringNull()
	}
	// Handling Lists
	if sdk.Destination != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Destination")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Destination, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Destination)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Destination = basetypes.NewListNull(elemType)
	}
	// Handling Lists
	if sdk.DestinationHip != nil {
		tflog.Debug(ctx, "Packing list of primitives for field DestinationHip")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.DestinationHip, d = basetypes.NewListValueFrom(ctx, elemType, sdk.DestinationHip)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.DestinationHip = basetypes.NewListNull(elemType)
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
	if sdk.Devices != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Devices")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Devices, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Devices)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Devices = basetypes.NewListNull(elemType)
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Disabled != nil {
		model.Disabled = basetypes.NewBoolValue(*sdk.Disabled)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Disabled", "value": *sdk.Disabled})
	} else {
		model.Disabled = basetypes.NewBoolNull()
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
	if sdk.From != nil {
		tflog.Debug(ctx, "Packing list of primitives for field From")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.From, d = basetypes.NewListValueFrom(ctx, elemType, sdk.From)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.From = basetypes.NewListNull(elemType)
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
	if sdk.LogEnd != nil {
		model.LogEnd = basetypes.NewBoolValue(*sdk.LogEnd)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "LogEnd", "value": *sdk.LogEnd})
	} else {
		model.LogEnd = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.LogSetting != nil {
		model.LogSetting = basetypes.NewStringValue(*sdk.LogSetting)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "LogSetting", "value": *sdk.LogSetting})
	} else {
		model.LogSetting = basetypes.NewStringNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.LogSettings != nil {
		tflog.Debug(ctx, "Packing nested object for field LogSettings")
		packed, d := packInternetRuleTypeLogSettingsFromSdk(ctx, *sdk.LogSettings)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "LogSettings"})
		}
		model.LogSettings = packed
	} else {
		model.LogSettings = basetypes.NewObjectNull(models.InternetRuleTypeLogSettings{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.LogStart != nil {
		model.LogStart = basetypes.NewBoolValue(*sdk.LogStart)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "LogStart", "value": *sdk.LogStart})
	} else {
		model.LogStart = basetypes.NewBoolNull()
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
	if sdk.NegateDestination != nil {
		model.NegateDestination = basetypes.NewBoolValue(*sdk.NegateDestination)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "NegateDestination", "value": *sdk.NegateDestination})
	} else {
		model.NegateDestination = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.NegateSource != nil {
		model.NegateSource = basetypes.NewBoolValue(*sdk.NegateSource)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "NegateSource", "value": *sdk.NegateSource})
	} else {
		model.NegateSource = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.NegateUser != nil {
		model.NegateUser = basetypes.NewBoolValue(*sdk.NegateUser)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "NegateUser", "value": *sdk.NegateUser})
	} else {
		model.NegateUser = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.PolicyType != nil {
		model.PolicyType = basetypes.NewStringValue(*sdk.PolicyType)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "PolicyType", "value": *sdk.PolicyType})
	} else {
		model.PolicyType = basetypes.NewStringNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.ProfileSetting != nil {
		tflog.Debug(ctx, "Packing nested object for field ProfileSetting")
		packed, d := packSecurityRuleTypeProfileSettingFromSdk(ctx, *sdk.ProfileSetting)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "ProfileSetting"})
		}
		model.ProfileSetting = packed
	} else {
		model.ProfileSetting = basetypes.NewObjectNull(models.SecurityRuleTypeProfileSetting{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Schedule != nil {
		model.Schedule = basetypes.NewStringValue(*sdk.Schedule)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Schedule", "value": *sdk.Schedule})
	} else {
		model.Schedule = basetypes.NewStringNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.SecuritySettings != nil {
		tflog.Debug(ctx, "Packing nested object for field SecuritySettings")
		packed, d := packInternetRuleTypeSecuritySettingsFromSdk(ctx, *sdk.SecuritySettings)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "SecuritySettings"})
		}
		model.SecuritySettings = packed
	} else {
		model.SecuritySettings = basetypes.NewObjectNull(models.InternetRuleTypeSecuritySettings{}.AttrTypes())
	}
	// Handling Lists
	if sdk.Service != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Service")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Service, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Service)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Service = basetypes.NewListNull(elemType)
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Snippet != nil {
		model.Snippet = basetypes.NewStringValue(*sdk.Snippet)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	} else {
		model.Snippet = basetypes.NewStringNull()
	}
	// Handling Lists
	if sdk.Source != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Source")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Source, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Source)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Source = basetypes.NewListNull(elemType)
	}
	// Handling Lists
	if sdk.SourceHip != nil {
		tflog.Debug(ctx, "Packing list of primitives for field SourceHip")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.SourceHip, d = basetypes.NewListValueFrom(ctx, elemType, sdk.SourceHip)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.SourceHip = basetypes.NewListNull(elemType)
	}
	// Handling Lists
	if sdk.SourceUser != nil {
		tflog.Debug(ctx, "Packing list of primitives for field SourceUser")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.SourceUser, d = basetypes.NewListValueFrom(ctx, elemType, sdk.SourceUser)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.SourceUser = basetypes.NewListNull(elemType)
	}
	// Handling Lists
	if sdk.Tag != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Tag")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Tag, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Tag)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Tag = basetypes.NewListNull(elemType)
	}
	// Handling Lists
	if sdk.TenantRestrictions != nil {
		tflog.Debug(ctx, "Packing list of primitives for field TenantRestrictions")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.TenantRestrictions, d = basetypes.NewListValueFrom(ctx, elemType, sdk.TenantRestrictions)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.TenantRestrictions = basetypes.NewListNull(elemType)
	}
	// Handling Lists
	if sdk.To != nil {
		tflog.Debug(ctx, "Packing list of primitives for field To")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.To, d = basetypes.NewListValueFrom(ctx, elemType, sdk.To)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.To = basetypes.NewListNull(elemType)
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.SecurityRules{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.SecurityRules", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for SecurityRules ---
func unpackSecurityRulesListToSdk(ctx context.Context, list types.List) ([]security_services.SecurityRules, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.SecurityRules")
	diags := diag.Diagnostics{}
	var data []models.SecurityRules
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.SecurityRules, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.SecurityRules{}.AttrTypes(), &item)
		unpacked, d := unpackSecurityRulesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.SecurityRules", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for SecurityRules ---
func packSecurityRulesListFromSdk(ctx context.Context, sdks []security_services.SecurityRules) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.SecurityRules")
	diags := diag.Diagnostics{}
	var data []models.SecurityRules

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.SecurityRules
		obj, d := packSecurityRulesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.SecurityRules{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.SecurityRules", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.SecurityRules{}.AttrType(), data)
}

// --- Unpacker for InternetRuleTypeAllowUrlCategoryInner ---
func unpackInternetRuleTypeAllowUrlCategoryInnerToSdk(ctx context.Context, obj types.Object) (*security_services.InternetRuleTypeAllowUrlCategoryInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.InternetRuleTypeAllowUrlCategoryInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.InternetRuleTypeAllowUrlCategoryInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.InternetRuleTypeAllowUrlCategoryInner
	var d diag.Diagnostics
	// Handling Primitives
	if !model.AdditionalAction.IsNull() && !model.AdditionalAction.IsUnknown() {
		sdk.AdditionalAction = model.AdditionalAction.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "AdditionalAction", "value": *sdk.AdditionalAction})
	}

	// Handling Primitives
	if !model.CredentialEnforcement.IsNull() && !model.CredentialEnforcement.IsUnknown() {
		sdk.CredentialEnforcement = model.CredentialEnforcement.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "CredentialEnforcement", "value": *sdk.CredentialEnforcement})
	}

	// Handling Primitives
	if !model.Decryption.IsNull() && !model.Decryption.IsUnknown() {
		sdk.Decryption = model.Decryption.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Decryption", "value": *sdk.Decryption})
	}

	// Handling Primitives
	if !model.Dlp.IsNull() && !model.Dlp.IsUnknown() {
		sdk.Dlp = model.Dlp.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Dlp", "value": *sdk.Dlp})
	}

	// Handling Objects
	if !model.FileControl.IsNull() && !model.FileControl.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field FileControl")
		unpacked, d := unpackInternetRuleTypeAllowUrlCategoryInnerFileControlToSdk(ctx, model.FileControl)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "FileControl"})
		}
		if unpacked != nil {
			sdk.FileControl = unpacked
		}
	}

	// Handling Primitives
	if !model.IsolationProfiles.IsNull() && !model.IsolationProfiles.IsUnknown() {
		sdk.IsolationProfiles = model.IsolationProfiles.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "IsolationProfiles", "value": *sdk.IsolationProfiles})
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.InternetRuleTypeAllowUrlCategoryInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for InternetRuleTypeAllowUrlCategoryInner ---
func packInternetRuleTypeAllowUrlCategoryInnerFromSdk(ctx context.Context, sdk security_services.InternetRuleTypeAllowUrlCategoryInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.InternetRuleTypeAllowUrlCategoryInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.InternetRuleTypeAllowUrlCategoryInner
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.AdditionalAction != nil {
		model.AdditionalAction = basetypes.NewStringValue(*sdk.AdditionalAction)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "AdditionalAction", "value": *sdk.AdditionalAction})
	} else {
		model.AdditionalAction = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.CredentialEnforcement != nil {
		model.CredentialEnforcement = basetypes.NewStringValue(*sdk.CredentialEnforcement)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "CredentialEnforcement", "value": *sdk.CredentialEnforcement})
	} else {
		model.CredentialEnforcement = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Decryption != nil {
		model.Decryption = basetypes.NewStringValue(*sdk.Decryption)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Decryption", "value": *sdk.Decryption})
	} else {
		model.Decryption = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Dlp != nil {
		model.Dlp = basetypes.NewStringValue(*sdk.Dlp)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Dlp", "value": *sdk.Dlp})
	} else {
		model.Dlp = basetypes.NewStringNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.FileControl != nil {
		tflog.Debug(ctx, "Packing nested object for field FileControl")
		packed, d := packInternetRuleTypeAllowUrlCategoryInnerFileControlFromSdk(ctx, *sdk.FileControl)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "FileControl"})
		}
		model.FileControl = packed
	} else {
		model.FileControl = basetypes.NewObjectNull(models.InternetRuleTypeAllowUrlCategoryInnerFileControl{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.IsolationProfiles != nil {
		model.IsolationProfiles = basetypes.NewStringValue(*sdk.IsolationProfiles)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "IsolationProfiles", "value": *sdk.IsolationProfiles})
	} else {
		model.IsolationProfiles = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Name != nil {
		model.Name = basetypes.NewStringValue(*sdk.Name)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
	} else {
		model.Name = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.InternetRuleTypeAllowUrlCategoryInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.InternetRuleTypeAllowUrlCategoryInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for InternetRuleTypeAllowUrlCategoryInner ---
func unpackInternetRuleTypeAllowUrlCategoryInnerListToSdk(ctx context.Context, list types.List) ([]security_services.InternetRuleTypeAllowUrlCategoryInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.InternetRuleTypeAllowUrlCategoryInner")
	diags := diag.Diagnostics{}
	var data []models.InternetRuleTypeAllowUrlCategoryInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.InternetRuleTypeAllowUrlCategoryInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.InternetRuleTypeAllowUrlCategoryInner{}.AttrTypes(), &item)
		unpacked, d := unpackInternetRuleTypeAllowUrlCategoryInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.InternetRuleTypeAllowUrlCategoryInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for InternetRuleTypeAllowUrlCategoryInner ---
func packInternetRuleTypeAllowUrlCategoryInnerListFromSdk(ctx context.Context, sdks []security_services.InternetRuleTypeAllowUrlCategoryInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.InternetRuleTypeAllowUrlCategoryInner")
	diags := diag.Diagnostics{}
	var data []models.InternetRuleTypeAllowUrlCategoryInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.InternetRuleTypeAllowUrlCategoryInner
		obj, d := packInternetRuleTypeAllowUrlCategoryInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.InternetRuleTypeAllowUrlCategoryInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.InternetRuleTypeAllowUrlCategoryInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.InternetRuleTypeAllowUrlCategoryInner{}.AttrType(), data)
}

// --- Unpacker for InternetRuleTypeAllowUrlCategoryInnerFileControl ---
func unpackInternetRuleTypeAllowUrlCategoryInnerFileControlToSdk(ctx context.Context, obj types.Object) (*security_services.InternetRuleTypeAllowUrlCategoryInnerFileControl, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.InternetRuleTypeAllowUrlCategoryInnerFileControl", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.InternetRuleTypeAllowUrlCategoryInnerFileControl
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.InternetRuleTypeAllowUrlCategoryInnerFileControl
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Download.IsNull() && !model.Download.IsUnknown() {
		sdk.Download = model.Download.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Download", "value": *sdk.Download})
	}

	// Handling Primitives
	if !model.Upload.IsNull() && !model.Upload.IsUnknown() {
		sdk.Upload = model.Upload.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Upload", "value": *sdk.Upload})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.InternetRuleTypeAllowUrlCategoryInnerFileControl", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for InternetRuleTypeAllowUrlCategoryInnerFileControl ---
func packInternetRuleTypeAllowUrlCategoryInnerFileControlFromSdk(ctx context.Context, sdk security_services.InternetRuleTypeAllowUrlCategoryInnerFileControl) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.InternetRuleTypeAllowUrlCategoryInnerFileControl", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.InternetRuleTypeAllowUrlCategoryInnerFileControl
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Download != nil {
		model.Download = basetypes.NewStringValue(*sdk.Download)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Download", "value": *sdk.Download})
	} else {
		model.Download = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Upload != nil {
		model.Upload = basetypes.NewStringValue(*sdk.Upload)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Upload", "value": *sdk.Upload})
	} else {
		model.Upload = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.InternetRuleTypeAllowUrlCategoryInnerFileControl{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.InternetRuleTypeAllowUrlCategoryInnerFileControl", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for InternetRuleTypeAllowUrlCategoryInnerFileControl ---
func unpackInternetRuleTypeAllowUrlCategoryInnerFileControlListToSdk(ctx context.Context, list types.List) ([]security_services.InternetRuleTypeAllowUrlCategoryInnerFileControl, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.InternetRuleTypeAllowUrlCategoryInnerFileControl")
	diags := diag.Diagnostics{}
	var data []models.InternetRuleTypeAllowUrlCategoryInnerFileControl
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.InternetRuleTypeAllowUrlCategoryInnerFileControl, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.InternetRuleTypeAllowUrlCategoryInnerFileControl{}.AttrTypes(), &item)
		unpacked, d := unpackInternetRuleTypeAllowUrlCategoryInnerFileControlToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.InternetRuleTypeAllowUrlCategoryInnerFileControl", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for InternetRuleTypeAllowUrlCategoryInnerFileControl ---
func packInternetRuleTypeAllowUrlCategoryInnerFileControlListFromSdk(ctx context.Context, sdks []security_services.InternetRuleTypeAllowUrlCategoryInnerFileControl) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.InternetRuleTypeAllowUrlCategoryInnerFileControl")
	diags := diag.Diagnostics{}
	var data []models.InternetRuleTypeAllowUrlCategoryInnerFileControl

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.InternetRuleTypeAllowUrlCategoryInnerFileControl
		obj, d := packInternetRuleTypeAllowUrlCategoryInnerFileControlFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.InternetRuleTypeAllowUrlCategoryInnerFileControl{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.InternetRuleTypeAllowUrlCategoryInnerFileControl", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.InternetRuleTypeAllowUrlCategoryInnerFileControl{}.AttrType(), data)
}

// --- Unpacker for InternetRuleTypeAllowWebApplicationInner ---
func unpackInternetRuleTypeAllowWebApplicationInnerToSdk(ctx context.Context, obj types.Object) (*security_services.InternetRuleTypeAllowWebApplicationInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.InternetRuleTypeAllowWebApplicationInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.InternetRuleTypeAllowWebApplicationInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.InternetRuleTypeAllowWebApplicationInner
	var d diag.Diagnostics
	// Handling Lists
	if !model.ApplicationFunction.IsNull() && !model.ApplicationFunction.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field ApplicationFunction")
		diags.Append(model.ApplicationFunction.ElementsAs(ctx, &sdk.ApplicationFunction, false)...)
	}

	// Handling Primitives
	if !model.Dlp.IsNull() && !model.Dlp.IsUnknown() {
		sdk.Dlp = model.Dlp.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Dlp", "value": *sdk.Dlp})
	}

	// Handling Objects
	if !model.FileControl.IsNull() && !model.FileControl.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field FileControl")
		unpacked, d := unpackInternetRuleTypeAllowUrlCategoryInnerFileControlToSdk(ctx, model.FileControl)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "FileControl"})
		}
		if unpacked != nil {
			sdk.FileControl = unpacked
		}
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
	}

	// Handling Objects
	if !model.SaasEnterpriseControl.IsNull() && !model.SaasEnterpriseControl.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field SaasEnterpriseControl")
		unpacked, d := unpackInternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlToSdk(ctx, model.SaasEnterpriseControl)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "SaasEnterpriseControl"})
		}
		if unpacked != nil {
			sdk.SaasEnterpriseControl = unpacked
		}
	}

	// Handling Lists
	if !model.SaasTenantList.IsNull() && !model.SaasTenantList.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field SaasTenantList")
		diags.Append(model.SaasTenantList.ElementsAs(ctx, &sdk.SaasTenantList, false)...)
	}

	// Handling Lists
	if !model.SaasUserList.IsNull() && !model.SaasUserList.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field SaasUserList")
		diags.Append(model.SaasUserList.ElementsAs(ctx, &sdk.SaasUserList, false)...)
	}

	// Handling Objects
	if !model.TenantControl.IsNull() && !model.TenantControl.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field TenantControl")
		unpacked, d := unpackInternetRuleTypeAllowWebApplicationInnerTenantControlToSdk(ctx, model.TenantControl)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "TenantControl"})
		}
		if unpacked != nil {
			sdk.TenantControl = unpacked
		}
	}

	// Handling Primitives
	if !model.Type.IsNull() && !model.Type.IsUnknown() {
		sdk.Type = model.Type.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Type", "value": *sdk.Type})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.InternetRuleTypeAllowWebApplicationInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for InternetRuleTypeAllowWebApplicationInner ---
func packInternetRuleTypeAllowWebApplicationInnerFromSdk(ctx context.Context, sdk security_services.InternetRuleTypeAllowWebApplicationInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.InternetRuleTypeAllowWebApplicationInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.InternetRuleTypeAllowWebApplicationInner
	var d diag.Diagnostics
	// Handling Lists
	if sdk.ApplicationFunction != nil {
		tflog.Debug(ctx, "Packing list of primitives for field ApplicationFunction")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.ApplicationFunction, d = basetypes.NewListValueFrom(ctx, elemType, sdk.ApplicationFunction)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.ApplicationFunction = basetypes.NewListNull(elemType)
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Dlp != nil {
		model.Dlp = basetypes.NewStringValue(*sdk.Dlp)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Dlp", "value": *sdk.Dlp})
	} else {
		model.Dlp = basetypes.NewStringNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.FileControl != nil {
		tflog.Debug(ctx, "Packing nested object for field FileControl")
		packed, d := packInternetRuleTypeAllowUrlCategoryInnerFileControlFromSdk(ctx, *sdk.FileControl)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "FileControl"})
		}
		model.FileControl = packed
	} else {
		model.FileControl = basetypes.NewObjectNull(models.InternetRuleTypeAllowUrlCategoryInnerFileControl{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Name != nil {
		model.Name = basetypes.NewStringValue(*sdk.Name)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
	} else {
		model.Name = basetypes.NewStringNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.SaasEnterpriseControl != nil {
		tflog.Debug(ctx, "Packing nested object for field SaasEnterpriseControl")
		packed, d := packInternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlFromSdk(ctx, *sdk.SaasEnterpriseControl)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "SaasEnterpriseControl"})
		}
		model.SaasEnterpriseControl = packed
	} else {
		model.SaasEnterpriseControl = basetypes.NewObjectNull(models.InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControl{}.AttrTypes())
	}
	// Handling Lists
	if sdk.SaasTenantList != nil {
		tflog.Debug(ctx, "Packing list of primitives for field SaasTenantList")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.SaasTenantList, d = basetypes.NewListValueFrom(ctx, elemType, sdk.SaasTenantList)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.SaasTenantList = basetypes.NewListNull(elemType)
	}
	// Handling Lists
	if sdk.SaasUserList != nil {
		tflog.Debug(ctx, "Packing list of primitives for field SaasUserList")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.SaasUserList, d = basetypes.NewListValueFrom(ctx, elemType, sdk.SaasUserList)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.SaasUserList = basetypes.NewListNull(elemType)
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.TenantControl != nil {
		tflog.Debug(ctx, "Packing nested object for field TenantControl")
		packed, d := packInternetRuleTypeAllowWebApplicationInnerTenantControlFromSdk(ctx, *sdk.TenantControl)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "TenantControl"})
		}
		model.TenantControl = packed
	} else {
		model.TenantControl = basetypes.NewObjectNull(models.InternetRuleTypeAllowWebApplicationInnerTenantControl{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Type != nil {
		model.Type = basetypes.NewStringValue(*sdk.Type)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Type", "value": *sdk.Type})
	} else {
		model.Type = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.InternetRuleTypeAllowWebApplicationInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.InternetRuleTypeAllowWebApplicationInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for InternetRuleTypeAllowWebApplicationInner ---
func unpackInternetRuleTypeAllowWebApplicationInnerListToSdk(ctx context.Context, list types.List) ([]security_services.InternetRuleTypeAllowWebApplicationInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.InternetRuleTypeAllowWebApplicationInner")
	diags := diag.Diagnostics{}
	var data []models.InternetRuleTypeAllowWebApplicationInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.InternetRuleTypeAllowWebApplicationInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.InternetRuleTypeAllowWebApplicationInner{}.AttrTypes(), &item)
		unpacked, d := unpackInternetRuleTypeAllowWebApplicationInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.InternetRuleTypeAllowWebApplicationInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for InternetRuleTypeAllowWebApplicationInner ---
func packInternetRuleTypeAllowWebApplicationInnerListFromSdk(ctx context.Context, sdks []security_services.InternetRuleTypeAllowWebApplicationInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.InternetRuleTypeAllowWebApplicationInner")
	diags := diag.Diagnostics{}
	var data []models.InternetRuleTypeAllowWebApplicationInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.InternetRuleTypeAllowWebApplicationInner
		obj, d := packInternetRuleTypeAllowWebApplicationInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.InternetRuleTypeAllowWebApplicationInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.InternetRuleTypeAllowWebApplicationInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.InternetRuleTypeAllowWebApplicationInner{}.AttrType(), data)
}

// --- Unpacker for InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControl ---
func unpackInternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlToSdk(ctx context.Context, obj types.Object) (*security_services.InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControl, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControl", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControl
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControl
	var d diag.Diagnostics
	// Handling Objects
	if !model.ConsumerAccess.IsNull() && !model.ConsumerAccess.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field ConsumerAccess")
		unpacked, d := unpackInternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlConsumerAccessToSdk(ctx, model.ConsumerAccess)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "ConsumerAccess"})
		}
		if unpacked != nil {
			sdk.ConsumerAccess = unpacked
		}
	}

	// Handling Objects
	if !model.EnterpriseAccess.IsNull() && !model.EnterpriseAccess.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field EnterpriseAccess")
		unpacked, d := unpackInternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlEnterpriseAccessToSdk(ctx, model.EnterpriseAccess)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "EnterpriseAccess"})
		}
		if unpacked != nil {
			sdk.EnterpriseAccess = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControl", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControl ---
func packInternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlFromSdk(ctx context.Context, sdk security_services.InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControl) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControl", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControl
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.ConsumerAccess != nil {
		tflog.Debug(ctx, "Packing nested object for field ConsumerAccess")
		packed, d := packInternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlConsumerAccessFromSdk(ctx, *sdk.ConsumerAccess)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "ConsumerAccess"})
		}
		model.ConsumerAccess = packed
	} else {
		model.ConsumerAccess = basetypes.NewObjectNull(models.InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlConsumerAccess{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.EnterpriseAccess != nil {
		tflog.Debug(ctx, "Packing nested object for field EnterpriseAccess")
		packed, d := packInternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlEnterpriseAccessFromSdk(ctx, *sdk.EnterpriseAccess)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "EnterpriseAccess"})
		}
		model.EnterpriseAccess = packed
	} else {
		model.EnterpriseAccess = basetypes.NewObjectNull(models.InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlEnterpriseAccess{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControl{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControl", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControl ---
func unpackInternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlListToSdk(ctx context.Context, list types.List) ([]security_services.InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControl, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControl")
	diags := diag.Diagnostics{}
	var data []models.InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControl
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControl, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControl{}.AttrTypes(), &item)
		unpacked, d := unpackInternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControl", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControl ---
func packInternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlListFromSdk(ctx context.Context, sdks []security_services.InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControl) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControl")
	diags := diag.Diagnostics{}
	var data []models.InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControl

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControl
		obj, d := packInternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControl{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControl", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControl{}.AttrType(), data)
}

// --- Unpacker for InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlConsumerAccess ---
func unpackInternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlConsumerAccessToSdk(ctx context.Context, obj types.Object) (*security_services.InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlConsumerAccess, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlConsumerAccess", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlConsumerAccess
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlConsumerAccess
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Enable.IsNull() && !model.Enable.IsUnknown() {
		sdk.Enable = model.Enable.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Enable", "value": *sdk.Enable})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlConsumerAccess", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlConsumerAccess ---
func packInternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlConsumerAccessFromSdk(ctx context.Context, sdk security_services.InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlConsumerAccess) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlConsumerAccess", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlConsumerAccess
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Enable != nil {
		model.Enable = basetypes.NewStringValue(*sdk.Enable)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Enable", "value": *sdk.Enable})
	} else {
		model.Enable = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlConsumerAccess{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlConsumerAccess", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlConsumerAccess ---
func unpackInternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlConsumerAccessListToSdk(ctx context.Context, list types.List) ([]security_services.InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlConsumerAccess, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlConsumerAccess")
	diags := diag.Diagnostics{}
	var data []models.InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlConsumerAccess
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlConsumerAccess, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlConsumerAccess{}.AttrTypes(), &item)
		unpacked, d := unpackInternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlConsumerAccessToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlConsumerAccess", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlConsumerAccess ---
func packInternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlConsumerAccessListFromSdk(ctx context.Context, sdks []security_services.InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlConsumerAccess) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlConsumerAccess")
	diags := diag.Diagnostics{}
	var data []models.InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlConsumerAccess

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlConsumerAccess
		obj, d := packInternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlConsumerAccessFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlConsumerAccess{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlConsumerAccess", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlConsumerAccess{}.AttrType(), data)
}

// --- Unpacker for InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlEnterpriseAccess ---
func unpackInternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlEnterpriseAccessToSdk(ctx context.Context, obj types.Object) (*security_services.InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlEnterpriseAccess, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlEnterpriseAccess", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlEnterpriseAccess
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlEnterpriseAccess
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Enable.IsNull() && !model.Enable.IsUnknown() {
		sdk.Enable = model.Enable.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Enable", "value": *sdk.Enable})
	}

	// Handling Lists
	if !model.TenantRestrictions.IsNull() && !model.TenantRestrictions.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field TenantRestrictions")
		diags.Append(model.TenantRestrictions.ElementsAs(ctx, &sdk.TenantRestrictions, false)...)
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlEnterpriseAccess", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlEnterpriseAccess ---
func packInternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlEnterpriseAccessFromSdk(ctx context.Context, sdk security_services.InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlEnterpriseAccess) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlEnterpriseAccess", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlEnterpriseAccess
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Enable != nil {
		model.Enable = basetypes.NewStringValue(*sdk.Enable)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Enable", "value": *sdk.Enable})
	} else {
		model.Enable = basetypes.NewStringNull()
	}
	// Handling Lists
	if sdk.TenantRestrictions != nil {
		tflog.Debug(ctx, "Packing list of primitives for field TenantRestrictions")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.TenantRestrictions, d = basetypes.NewListValueFrom(ctx, elemType, sdk.TenantRestrictions)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.TenantRestrictions = basetypes.NewListNull(elemType)
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlEnterpriseAccess{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlEnterpriseAccess", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlEnterpriseAccess ---
func unpackInternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlEnterpriseAccessListToSdk(ctx context.Context, list types.List) ([]security_services.InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlEnterpriseAccess, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlEnterpriseAccess")
	diags := diag.Diagnostics{}
	var data []models.InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlEnterpriseAccess
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlEnterpriseAccess, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlEnterpriseAccess{}.AttrTypes(), &item)
		unpacked, d := unpackInternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlEnterpriseAccessToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlEnterpriseAccess", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlEnterpriseAccess ---
func packInternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlEnterpriseAccessListFromSdk(ctx context.Context, sdks []security_services.InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlEnterpriseAccess) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlEnterpriseAccess")
	diags := diag.Diagnostics{}
	var data []models.InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlEnterpriseAccess

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlEnterpriseAccess
		obj, d := packInternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlEnterpriseAccessFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlEnterpriseAccess{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlEnterpriseAccess", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlEnterpriseAccess{}.AttrType(), data)
}

// --- Unpacker for InternetRuleTypeAllowWebApplicationInnerTenantControl ---
func unpackInternetRuleTypeAllowWebApplicationInnerTenantControlToSdk(ctx context.Context, obj types.Object) (*security_services.InternetRuleTypeAllowWebApplicationInnerTenantControl, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.InternetRuleTypeAllowWebApplicationInnerTenantControl", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.InternetRuleTypeAllowWebApplicationInnerTenantControl
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.InternetRuleTypeAllowWebApplicationInnerTenantControl
	var d diag.Diagnostics
	// Handling Lists
	if !model.AllowedActivities.IsNull() && !model.AllowedActivities.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field AllowedActivities")
		diags.Append(model.AllowedActivities.ElementsAs(ctx, &sdk.AllowedActivities, false)...)
	}

	// Handling Lists
	if !model.BlockedActivities.IsNull() && !model.BlockedActivities.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field BlockedActivities")
		diags.Append(model.BlockedActivities.ElementsAs(ctx, &sdk.BlockedActivities, false)...)
	}

	// Handling Primitives
	if !model.ParentApplication.IsNull() && !model.ParentApplication.IsUnknown() {
		sdk.ParentApplication = model.ParentApplication.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "ParentApplication", "value": *sdk.ParentApplication})
	}

	// Handling Lists
	if !model.Tenants.IsNull() && !model.Tenants.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Tenants")
		diags.Append(model.Tenants.ElementsAs(ctx, &sdk.Tenants, false)...)
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.InternetRuleTypeAllowWebApplicationInnerTenantControl", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for InternetRuleTypeAllowWebApplicationInnerTenantControl ---
func packInternetRuleTypeAllowWebApplicationInnerTenantControlFromSdk(ctx context.Context, sdk security_services.InternetRuleTypeAllowWebApplicationInnerTenantControl) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.InternetRuleTypeAllowWebApplicationInnerTenantControl", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.InternetRuleTypeAllowWebApplicationInnerTenantControl
	var d diag.Diagnostics
	// Handling Lists
	if sdk.AllowedActivities != nil {
		tflog.Debug(ctx, "Packing list of primitives for field AllowedActivities")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.AllowedActivities, d = basetypes.NewListValueFrom(ctx, elemType, sdk.AllowedActivities)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.AllowedActivities = basetypes.NewListNull(elemType)
	}
	// Handling Lists
	if sdk.BlockedActivities != nil {
		tflog.Debug(ctx, "Packing list of primitives for field BlockedActivities")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.BlockedActivities, d = basetypes.NewListValueFrom(ctx, elemType, sdk.BlockedActivities)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.BlockedActivities = basetypes.NewListNull(elemType)
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.ParentApplication != nil {
		model.ParentApplication = basetypes.NewStringValue(*sdk.ParentApplication)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "ParentApplication", "value": *sdk.ParentApplication})
	} else {
		model.ParentApplication = basetypes.NewStringNull()
	}
	// Handling Lists
	if sdk.Tenants != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Tenants")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Tenants, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Tenants)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Tenants = basetypes.NewListNull(elemType)
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.InternetRuleTypeAllowWebApplicationInnerTenantControl{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.InternetRuleTypeAllowWebApplicationInnerTenantControl", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for InternetRuleTypeAllowWebApplicationInnerTenantControl ---
func unpackInternetRuleTypeAllowWebApplicationInnerTenantControlListToSdk(ctx context.Context, list types.List) ([]security_services.InternetRuleTypeAllowWebApplicationInnerTenantControl, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.InternetRuleTypeAllowWebApplicationInnerTenantControl")
	diags := diag.Diagnostics{}
	var data []models.InternetRuleTypeAllowWebApplicationInnerTenantControl
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.InternetRuleTypeAllowWebApplicationInnerTenantControl, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.InternetRuleTypeAllowWebApplicationInnerTenantControl{}.AttrTypes(), &item)
		unpacked, d := unpackInternetRuleTypeAllowWebApplicationInnerTenantControlToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.InternetRuleTypeAllowWebApplicationInnerTenantControl", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for InternetRuleTypeAllowWebApplicationInnerTenantControl ---
func packInternetRuleTypeAllowWebApplicationInnerTenantControlListFromSdk(ctx context.Context, sdks []security_services.InternetRuleTypeAllowWebApplicationInnerTenantControl) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.InternetRuleTypeAllowWebApplicationInnerTenantControl")
	diags := diag.Diagnostics{}
	var data []models.InternetRuleTypeAllowWebApplicationInnerTenantControl

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.InternetRuleTypeAllowWebApplicationInnerTenantControl
		obj, d := packInternetRuleTypeAllowWebApplicationInnerTenantControlFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.InternetRuleTypeAllowWebApplicationInnerTenantControl{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.InternetRuleTypeAllowWebApplicationInnerTenantControl", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.InternetRuleTypeAllowWebApplicationInnerTenantControl{}.AttrType(), data)
}

// --- Unpacker for InternetRuleTypeDefaultProfileSettings ---
func unpackInternetRuleTypeDefaultProfileSettingsToSdk(ctx context.Context, obj types.Object) (*security_services.InternetRuleTypeDefaultProfileSettings, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.InternetRuleTypeDefaultProfileSettings", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.InternetRuleTypeDefaultProfileSettings
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.InternetRuleTypeDefaultProfileSettings
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Dlp.IsNull() && !model.Dlp.IsUnknown() {
		sdk.Dlp = model.Dlp.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Dlp", "value": *sdk.Dlp})
	}

	// Handling Objects
	if !model.FileControl.IsNull() && !model.FileControl.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field FileControl")
		unpacked, d := unpackInternetRuleTypeAllowUrlCategoryInnerFileControlToSdk(ctx, model.FileControl)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "FileControl"})
		}
		if unpacked != nil {
			sdk.FileControl = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.InternetRuleTypeDefaultProfileSettings", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for InternetRuleTypeDefaultProfileSettings ---
func packInternetRuleTypeDefaultProfileSettingsFromSdk(ctx context.Context, sdk security_services.InternetRuleTypeDefaultProfileSettings) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.InternetRuleTypeDefaultProfileSettings", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.InternetRuleTypeDefaultProfileSettings
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Dlp != nil {
		model.Dlp = basetypes.NewStringValue(*sdk.Dlp)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Dlp", "value": *sdk.Dlp})
	} else {
		model.Dlp = basetypes.NewStringNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.FileControl != nil {
		tflog.Debug(ctx, "Packing nested object for field FileControl")
		packed, d := packInternetRuleTypeAllowUrlCategoryInnerFileControlFromSdk(ctx, *sdk.FileControl)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "FileControl"})
		}
		model.FileControl = packed
	} else {
		model.FileControl = basetypes.NewObjectNull(models.InternetRuleTypeAllowUrlCategoryInnerFileControl{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.InternetRuleTypeDefaultProfileSettings{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.InternetRuleTypeDefaultProfileSettings", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for InternetRuleTypeDefaultProfileSettings ---
func unpackInternetRuleTypeDefaultProfileSettingsListToSdk(ctx context.Context, list types.List) ([]security_services.InternetRuleTypeDefaultProfileSettings, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.InternetRuleTypeDefaultProfileSettings")
	diags := diag.Diagnostics{}
	var data []models.InternetRuleTypeDefaultProfileSettings
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.InternetRuleTypeDefaultProfileSettings, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.InternetRuleTypeDefaultProfileSettings{}.AttrTypes(), &item)
		unpacked, d := unpackInternetRuleTypeDefaultProfileSettingsToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.InternetRuleTypeDefaultProfileSettings", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for InternetRuleTypeDefaultProfileSettings ---
func packInternetRuleTypeDefaultProfileSettingsListFromSdk(ctx context.Context, sdks []security_services.InternetRuleTypeDefaultProfileSettings) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.InternetRuleTypeDefaultProfileSettings")
	diags := diag.Diagnostics{}
	var data []models.InternetRuleTypeDefaultProfileSettings

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.InternetRuleTypeDefaultProfileSettings
		obj, d := packInternetRuleTypeDefaultProfileSettingsFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.InternetRuleTypeDefaultProfileSettings{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.InternetRuleTypeDefaultProfileSettings", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.InternetRuleTypeDefaultProfileSettings{}.AttrType(), data)
}

// --- Unpacker for InternetRuleTypeLogSettings ---
func unpackInternetRuleTypeLogSettingsToSdk(ctx context.Context, obj types.Object) (*security_services.InternetRuleTypeLogSettings, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.InternetRuleTypeLogSettings", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.InternetRuleTypeLogSettings
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.InternetRuleTypeLogSettings
	var d diag.Diagnostics
	// Handling Primitives
	if !model.LogSessions.IsNull() && !model.LogSessions.IsUnknown() {
		sdk.LogSessions = model.LogSessions.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "LogSessions", "value": *sdk.LogSessions})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.InternetRuleTypeLogSettings", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for InternetRuleTypeLogSettings ---
func packInternetRuleTypeLogSettingsFromSdk(ctx context.Context, sdk security_services.InternetRuleTypeLogSettings) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.InternetRuleTypeLogSettings", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.InternetRuleTypeLogSettings
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.LogSessions != nil {
		model.LogSessions = basetypes.NewBoolValue(*sdk.LogSessions)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "LogSessions", "value": *sdk.LogSessions})
	} else {
		model.LogSessions = basetypes.NewBoolNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.InternetRuleTypeLogSettings{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.InternetRuleTypeLogSettings", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for InternetRuleTypeLogSettings ---
func unpackInternetRuleTypeLogSettingsListToSdk(ctx context.Context, list types.List) ([]security_services.InternetRuleTypeLogSettings, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.InternetRuleTypeLogSettings")
	diags := diag.Diagnostics{}
	var data []models.InternetRuleTypeLogSettings
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.InternetRuleTypeLogSettings, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.InternetRuleTypeLogSettings{}.AttrTypes(), &item)
		unpacked, d := unpackInternetRuleTypeLogSettingsToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.InternetRuleTypeLogSettings", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for InternetRuleTypeLogSettings ---
func packInternetRuleTypeLogSettingsListFromSdk(ctx context.Context, sdks []security_services.InternetRuleTypeLogSettings) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.InternetRuleTypeLogSettings")
	diags := diag.Diagnostics{}
	var data []models.InternetRuleTypeLogSettings

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.InternetRuleTypeLogSettings
		obj, d := packInternetRuleTypeLogSettingsFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.InternetRuleTypeLogSettings{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.InternetRuleTypeLogSettings", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.InternetRuleTypeLogSettings{}.AttrType(), data)
}

// --- Unpacker for SecurityRuleTypeProfileSetting ---
func unpackSecurityRuleTypeProfileSettingToSdk(ctx context.Context, obj types.Object) (*security_services.SecurityRuleTypeProfileSetting, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.SecurityRuleTypeProfileSetting", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.SecurityRuleTypeProfileSetting
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.SecurityRuleTypeProfileSetting
	var d diag.Diagnostics
	// Handling Lists
	if !model.Group.IsNull() && !model.Group.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Group")
		diags.Append(model.Group.ElementsAs(ctx, &sdk.Group, false)...)
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.SecurityRuleTypeProfileSetting", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for SecurityRuleTypeProfileSetting ---
func packSecurityRuleTypeProfileSettingFromSdk(ctx context.Context, sdk security_services.SecurityRuleTypeProfileSetting) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.SecurityRuleTypeProfileSetting", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.SecurityRuleTypeProfileSetting
	var d diag.Diagnostics
	// Handling Lists
	if sdk.Group != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Group")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Group, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Group)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Group = basetypes.NewListNull(elemType)
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.SecurityRuleTypeProfileSetting{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.SecurityRuleTypeProfileSetting", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for SecurityRuleTypeProfileSetting ---
func unpackSecurityRuleTypeProfileSettingListToSdk(ctx context.Context, list types.List) ([]security_services.SecurityRuleTypeProfileSetting, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.SecurityRuleTypeProfileSetting")
	diags := diag.Diagnostics{}
	var data []models.SecurityRuleTypeProfileSetting
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.SecurityRuleTypeProfileSetting, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.SecurityRuleTypeProfileSetting{}.AttrTypes(), &item)
		unpacked, d := unpackSecurityRuleTypeProfileSettingToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.SecurityRuleTypeProfileSetting", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for SecurityRuleTypeProfileSetting ---
func packSecurityRuleTypeProfileSettingListFromSdk(ctx context.Context, sdks []security_services.SecurityRuleTypeProfileSetting) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.SecurityRuleTypeProfileSetting")
	diags := diag.Diagnostics{}
	var data []models.SecurityRuleTypeProfileSetting

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.SecurityRuleTypeProfileSetting
		obj, d := packSecurityRuleTypeProfileSettingFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.SecurityRuleTypeProfileSetting{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.SecurityRuleTypeProfileSetting", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.SecurityRuleTypeProfileSetting{}.AttrType(), data)
}

// --- Unpacker for InternetRuleTypeSecuritySettings ---
func unpackInternetRuleTypeSecuritySettingsToSdk(ctx context.Context, obj types.Object) (*security_services.InternetRuleTypeSecuritySettings, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.InternetRuleTypeSecuritySettings", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.InternetRuleTypeSecuritySettings
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.InternetRuleTypeSecuritySettings
	var d diag.Diagnostics
	// Handling Primitives
	if !model.AntiSpyware.IsNull() && !model.AntiSpyware.IsUnknown() {
		sdk.AntiSpyware = model.AntiSpyware.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "AntiSpyware", "value": *sdk.AntiSpyware})
	}

	// Handling Primitives
	if !model.VirusAndWildfireAnalysis.IsNull() && !model.VirusAndWildfireAnalysis.IsUnknown() {
		sdk.VirusAndWildfireAnalysis = model.VirusAndWildfireAnalysis.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "VirusAndWildfireAnalysis", "value": *sdk.VirusAndWildfireAnalysis})
	}

	// Handling Primitives
	if !model.Vulnerability.IsNull() && !model.Vulnerability.IsUnknown() {
		sdk.Vulnerability = model.Vulnerability.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Vulnerability", "value": *sdk.Vulnerability})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.InternetRuleTypeSecuritySettings", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for InternetRuleTypeSecuritySettings ---
func packInternetRuleTypeSecuritySettingsFromSdk(ctx context.Context, sdk security_services.InternetRuleTypeSecuritySettings) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.InternetRuleTypeSecuritySettings", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.InternetRuleTypeSecuritySettings
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.AntiSpyware != nil {
		model.AntiSpyware = basetypes.NewStringValue(*sdk.AntiSpyware)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "AntiSpyware", "value": *sdk.AntiSpyware})
	} else {
		model.AntiSpyware = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.VirusAndWildfireAnalysis != nil {
		model.VirusAndWildfireAnalysis = basetypes.NewStringValue(*sdk.VirusAndWildfireAnalysis)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "VirusAndWildfireAnalysis", "value": *sdk.VirusAndWildfireAnalysis})
	} else {
		model.VirusAndWildfireAnalysis = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Vulnerability != nil {
		model.Vulnerability = basetypes.NewStringValue(*sdk.Vulnerability)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Vulnerability", "value": *sdk.Vulnerability})
	} else {
		model.Vulnerability = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.InternetRuleTypeSecuritySettings{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.InternetRuleTypeSecuritySettings", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for InternetRuleTypeSecuritySettings ---
func unpackInternetRuleTypeSecuritySettingsListToSdk(ctx context.Context, list types.List) ([]security_services.InternetRuleTypeSecuritySettings, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.InternetRuleTypeSecuritySettings")
	diags := diag.Diagnostics{}
	var data []models.InternetRuleTypeSecuritySettings
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.InternetRuleTypeSecuritySettings, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.InternetRuleTypeSecuritySettings{}.AttrTypes(), &item)
		unpacked, d := unpackInternetRuleTypeSecuritySettingsToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.InternetRuleTypeSecuritySettings", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for InternetRuleTypeSecuritySettings ---
func packInternetRuleTypeSecuritySettingsListFromSdk(ctx context.Context, sdks []security_services.InternetRuleTypeSecuritySettings) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.InternetRuleTypeSecuritySettings")
	diags := diag.Diagnostics{}
	var data []models.InternetRuleTypeSecuritySettings

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.InternetRuleTypeSecuritySettings
		obj, d := packInternetRuleTypeSecuritySettingsFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.InternetRuleTypeSecuritySettings{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.InternetRuleTypeSecuritySettings", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.InternetRuleTypeSecuritySettings{}.AttrType(), data)
}
