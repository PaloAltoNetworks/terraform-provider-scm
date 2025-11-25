package provider

/*

import (
	"context"
	"fmt"
	"reflect"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/deployment_services"
	"github.com/paloaltonetworks/scm-go/generated/deployment_services"
)










// --- Unpacker for SharedInfrastructureSettings ---
func unpackSharedInfrastructureSettingsToSdk(ctx context.Context, obj types.Object) (*deployment_services.SharedInfrastructureSettings, diag.Diagnostics) {
    tflog.Debug(ctx, "Entering unpack helper for models.SharedInfrastructureSettings", map[string]interface{}{"tf_object": obj})
    diags := diag.Diagnostics{}
    var model models.SharedInfrastructureSettings
    diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
    if diags.HasError() {
        tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
        return nil, diags
    }
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

    var sdk deployment_services.SharedInfrastructureSettings
    var d diag.Diagnostics

    // Handling Primitives
    if !model.ApiKey.IsNull() && !model.ApiKey.IsUnknown() {
        sdk.ApiKey = model.ApiKey.ValueStringPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "ApiKey", "value": *sdk.ApiKey})
    }

    // Handling Primitives
    if !model.CaptivePortalRedirectIpAddress.IsNull() && !model.CaptivePortalRedirectIpAddress.IsUnknown() {
        sdk.CaptivePortalRedirectIpAddress = model.CaptivePortalRedirectIpAddress.ValueStringPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "CaptivePortalRedirectIpAddress", "value": *sdk.CaptivePortalRedirectIpAddress})
    }

    // Handling Objects
    if !model.ConnectorApplicationBlocks.IsNull() && !model.ConnectorApplicationBlocks.IsUnknown() {
        tflog.Debug(ctx, "Unpacking nested object for field ConnectorApplicationBlocks")
        unpacked, d := unpackEditSharedInfrastructureSettingsConnectorApplicationBlocksToSdk(ctx, model.ConnectorApplicationBlocks)
        diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "ConnectorApplicationBlocks"})
		}
        if unpacked != nil {
            sdk.ConnectorApplicationBlocks = unpacked
        }
    }

    // Handling Objects
    if !model.ConnectorConnectorBlocks.IsNull() && !model.ConnectorConnectorBlocks.IsUnknown() {
        tflog.Debug(ctx, "Unpacking nested object for field ConnectorConnectorBlocks")
        unpacked, d := unpackEditSharedInfrastructureSettingsConnectorConnectorBlocksToSdk(ctx, model.ConnectorConnectorBlocks)
        diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "ConnectorConnectorBlocks"})
		}
        if unpacked != nil {
            sdk.ConnectorConnectorBlocks = unpacked
        }
    }

    // Handling Primitives
    if !model.EgressIpNotificationUrl.IsNull() && !model.EgressIpNotificationUrl.IsUnknown() {
        sdk.EgressIpNotificationUrl = model.EgressIpNotificationUrl.ValueStringPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "EgressIpNotificationUrl", "value": *sdk.EgressIpNotificationUrl})
    }

    // Handling Primitives
    if !model.Folder.IsNull() && !model.Folder.IsUnknown() {
        sdk.Folder = model.Folder.ValueStringPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Folder", "value": *sdk.Folder})
    }

    // Handling Primitives
    if !model.InfraBgpAs.IsNull() && !model.InfraBgpAs.IsUnknown() {
        sdk.InfraBgpAs = model.InfraBgpAs.ValueStringPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "InfraBgpAs", "value": *sdk.InfraBgpAs})
    }

    // Handling Primitives
    if !model.InfrastructureSubnet.IsNull() && !model.InfrastructureSubnet.IsUnknown() {
        sdk.InfrastructureSubnet = model.InfrastructureSubnet.ValueStringPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "InfrastructureSubnet", "value": *sdk.InfrastructureSubnet})
    }

    // Handling Primitives
    if !model.InfrastructureSubnetIpv6.IsNull() && !model.InfrastructureSubnetIpv6.IsUnknown() {
        sdk.InfrastructureSubnetIpv6 = model.InfrastructureSubnetIpv6.ValueStringPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "InfrastructureSubnetIpv6", "value": *sdk.InfrastructureSubnetIpv6})
    }

    // Handling Primitives
    if !model.Ipv6.IsNull() && !model.Ipv6.IsUnknown() {
        sdk.Ipv6 = model.Ipv6.ValueBoolPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Ipv6", "value": *sdk.Ipv6})
    }

    // Handling Lists
    if !model.LoopbackIps.IsNull() && !model.LoopbackIps.IsUnknown() {
        tflog.Debug(ctx, "Unpacking list of primitives for field LoopbackIps")
        diags.Append(model.LoopbackIps.ElementsAs(ctx, &sdk.LoopbackIps, false)...)
    }

    // Handling Primitives
    if !model.TunnelMonitorIpAddress.IsNull() && !model.TunnelMonitorIpAddress.IsUnknown() {
        sdk.TunnelMonitorIpAddress = model.TunnelMonitorIpAddress.ValueStringPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "TunnelMonitorIpAddress", "value": *sdk.TunnelMonitorIpAddress})
    }

diags.Append(d...)

    tflog.Debug(ctx, "Exiting unpack helper for models.SharedInfrastructureSettings", map[string]interface{}{"has_errors": diags.HasError()})
    return &sdk, diags

}

// --- Packer for SharedInfrastructureSettings ---
func packSharedInfrastructureSettingsFromSdk(ctx context.Context, sdk deployment_services.SharedInfrastructureSettings) (types.Object, diag.Diagnostics) {
    tflog.Debug(ctx, "Entering pack helper for models.SharedInfrastructureSettings", map[string]interface{}{"sdk_struct": sdk})
    diags := diag.Diagnostics{}
    var model models.SharedInfrastructureSettings
    var d diag.Diagnostics
    // Handling Primitives
    // Standard primitive packing
    if sdk.ApiKey != nil {
        model.ApiKey = basetypes.NewStringValue(*sdk.ApiKey)
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "ApiKey", "value": *sdk.ApiKey})
    } else {
        model.ApiKey = basetypes.NewStringNull()
    }
    // Handling Primitives
    // Standard primitive packing
    if sdk.CaptivePortalRedirectIpAddress != nil {
        model.CaptivePortalRedirectIpAddress = basetypes.NewStringValue(*sdk.CaptivePortalRedirectIpAddress)
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "CaptivePortalRedirectIpAddress", "value": *sdk.CaptivePortalRedirectIpAddress})
    } else {
        model.CaptivePortalRedirectIpAddress = basetypes.NewStringNull()
    }
    // Handling Objects
    // This is a regular nested object that has its own packer.
    if sdk.ConnectorApplicationBlocks != nil {
        tflog.Debug(ctx, "Packing nested object for field ConnectorApplicationBlocks")
        packed, d := packEditSharedInfrastructureSettingsConnectorApplicationBlocksFromSdk(ctx, *sdk.ConnectorApplicationBlocks)
        diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "ConnectorApplicationBlocks"})
		}
        model.ConnectorApplicationBlocks = packed
    } else {
        model.ConnectorApplicationBlocks = basetypes.NewObjectNull(models.EditSharedInfrastructureSettingsConnectorApplicationBlocks{}.AttrTypes())
    }
    // Handling Objects
    // This is a regular nested object that has its own packer.
    if sdk.ConnectorConnectorBlocks != nil {
        tflog.Debug(ctx, "Packing nested object for field ConnectorConnectorBlocks")
        packed, d := packEditSharedInfrastructureSettingsConnectorConnectorBlocksFromSdk(ctx, *sdk.ConnectorConnectorBlocks)
        diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "ConnectorConnectorBlocks"})
		}
        model.ConnectorConnectorBlocks = packed
    } else {
        model.ConnectorConnectorBlocks = basetypes.NewObjectNull(models.EditSharedInfrastructureSettingsConnectorConnectorBlocks{}.AttrTypes())
    }
    // Handling Primitives
    // Standard primitive packing
    if sdk.EgressIpNotificationUrl != nil {
        model.EgressIpNotificationUrl = basetypes.NewStringValue(*sdk.EgressIpNotificationUrl)
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "EgressIpNotificationUrl", "value": *sdk.EgressIpNotificationUrl})
    } else {
        model.EgressIpNotificationUrl = basetypes.NewStringNull()
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
    if sdk.InfraBgpAs != nil {
        model.InfraBgpAs = basetypes.NewStringValue(*sdk.InfraBgpAs)
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "InfraBgpAs", "value": *sdk.InfraBgpAs})
    } else {
        model.InfraBgpAs = basetypes.NewStringNull()
    }
    // Handling Primitives
    // Standard primitive packing
    if sdk.InfrastructureSubnet != nil {
        model.InfrastructureSubnet = basetypes.NewStringValue(*sdk.InfrastructureSubnet)
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "InfrastructureSubnet", "value": *sdk.InfrastructureSubnet})
    } else {
        model.InfrastructureSubnet = basetypes.NewStringNull()
    }
    // Handling Primitives
    // Standard primitive packing
    if sdk.InfrastructureSubnetIpv6 != nil {
        model.InfrastructureSubnetIpv6 = basetypes.NewStringValue(*sdk.InfrastructureSubnetIpv6)
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "InfrastructureSubnetIpv6", "value": *sdk.InfrastructureSubnetIpv6})
    } else {
        model.InfrastructureSubnetIpv6 = basetypes.NewStringNull()
    }
    // Handling Primitives
    // Standard primitive packing
    if sdk.Ipv6 != nil {
        model.Ipv6 = basetypes.NewBoolValue(*sdk.Ipv6)
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Ipv6", "value": *sdk.Ipv6})
    } else {
        model.Ipv6 = basetypes.NewBoolNull()
    }
    // Handling Lists
    if sdk.LoopbackIps != nil {
        tflog.Debug(ctx, "Packing list of primitives for field LoopbackIps")
        var d diag.Diagnostics
        // This logic now dynamically determines the element type based on the SDK's Go type.
        var elemType attr.Type = basetypes.StringType{} // Default to string
        model.LoopbackIps, d = basetypes.NewListValueFrom(ctx, elemType, sdk.LoopbackIps)
        diags.Append(d...)
    } else {
        // This logic now creates a correctly typed null list.
        var elemType attr.Type = basetypes.StringType{} // Default to string
        model.LoopbackIps = basetypes.NewListNull(elemType)
    }
    // Handling Primitives
    // Standard primitive packing
    if sdk.TunnelMonitorIpAddress != nil {
        model.TunnelMonitorIpAddress = basetypes.NewStringValue(*sdk.TunnelMonitorIpAddress)
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "TunnelMonitorIpAddress", "value": *sdk.TunnelMonitorIpAddress})
    } else {
        model.TunnelMonitorIpAddress = basetypes.NewStringNull()
    }
diags.Append(d...)

    obj, d := types.ObjectValueFrom(ctx, models.SharedInfrastructureSettings{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
    diags.Append(d...)
    tflog.Debug(ctx, "Exiting pack helper for models.SharedInfrastructureSettings", map[string]interface{}{"has_errors": diags.HasError()})
    return obj, diags

}

// --- List Unpacker for SharedInfrastructureSettings ---
func unpackSharedInfrastructureSettingsListToSdk(ctx context.Context, list types.List) ([]deployment_services.SharedInfrastructureSettings, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.SharedInfrastructureSettings")
	diags := diag.Diagnostics{}
	var data []models.SharedInfrastructureSettings
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]deployment_services.SharedInfrastructureSettings, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.SharedInfrastructureSettings{}.AttrTypes(), &item)
		unpacked, d := unpackSharedInfrastructureSettingsToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.SharedInfrastructureSettings", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for SharedInfrastructureSettings ---
func packSharedInfrastructureSettingsListFromSdk(ctx context.Context, sdks []deployment_services.SharedInfrastructureSettings) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.SharedInfrastructureSettings")
	diags := diag.Diagnostics{}
	var data []models.SharedInfrastructureSettings

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.SharedInfrastructureSettings
		obj, d := packSharedInfrastructureSettingsFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.SharedInfrastructureSettings{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.SharedInfrastructureSettings", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.SharedInfrastructureSettings{}.AttrType(), data)
}

// --- Unpacker for EditSharedInfrastructureSettingsConnectorApplicationBlocks ---
func unpackEditSharedInfrastructureSettingsConnectorApplicationBlocksToSdk(ctx context.Context, obj types.Object) (*deployment_services.EditSharedInfrastructureSettingsConnectorApplicationBlocks, diag.Diagnostics) {
    tflog.Debug(ctx, "Entering unpack helper for models.EditSharedInfrastructureSettingsConnectorApplicationBlocks", map[string]interface{}{"tf_object": obj})
    diags := diag.Diagnostics{}
    var model models.EditSharedInfrastructureSettingsConnectorApplicationBlocks
    diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
    if diags.HasError() {
        tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
        return nil, diags
    }
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

    var sdk deployment_services.EditSharedInfrastructureSettingsConnectorApplicationBlocks
    var d diag.Diagnostics
    // Handling Lists
    if !model.Member.IsNull() && !model.Member.IsUnknown() {
        tflog.Debug(ctx, "Unpacking list of primitives for field Member")
        diags.Append(model.Member.ElementsAs(ctx, &sdk.Member, false)...)
    }

diags.Append(d...)

    tflog.Debug(ctx, "Exiting unpack helper for models.EditSharedInfrastructureSettingsConnectorApplicationBlocks", map[string]interface{}{"has_errors": diags.HasError()})
    return &sdk, diags

}

// --- Packer for EditSharedInfrastructureSettingsConnectorApplicationBlocks ---
func packEditSharedInfrastructureSettingsConnectorApplicationBlocksFromSdk(ctx context.Context, sdk deployment_services.EditSharedInfrastructureSettingsConnectorApplicationBlocks) (types.Object, diag.Diagnostics) {
    tflog.Debug(ctx, "Entering pack helper for models.EditSharedInfrastructureSettingsConnectorApplicationBlocks", map[string]interface{}{"sdk_struct": sdk})
    diags := diag.Diagnostics{}
    var model models.EditSharedInfrastructureSettingsConnectorApplicationBlocks
    var d diag.Diagnostics
    // Handling Lists
    if sdk.Member != nil {
        tflog.Debug(ctx, "Packing list of primitives for field Member")
        var d diag.Diagnostics
        // This logic now dynamically determines the element type based on the SDK's Go type.
        var elemType attr.Type = basetypes.StringType{} // Default to string
        model.Member, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Member)
        diags.Append(d...)
    } else {
        // This logic now creates a correctly typed null list.
        var elemType attr.Type = basetypes.StringType{} // Default to string
        model.Member = basetypes.NewListNull(elemType)
    }
diags.Append(d...)

    obj, d := types.ObjectValueFrom(ctx, models.EditSharedInfrastructureSettingsConnectorApplicationBlocks{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
    diags.Append(d...)
    tflog.Debug(ctx, "Exiting pack helper for models.EditSharedInfrastructureSettingsConnectorApplicationBlocks", map[string]interface{}{"has_errors": diags.HasError()})
    return obj, diags

}

// --- List Unpacker for EditSharedInfrastructureSettingsConnectorApplicationBlocks ---
func unpackEditSharedInfrastructureSettingsConnectorApplicationBlocksListToSdk(ctx context.Context, list types.List) ([]deployment_services.EditSharedInfrastructureSettingsConnectorApplicationBlocks, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.EditSharedInfrastructureSettingsConnectorApplicationBlocks")
	diags := diag.Diagnostics{}
	var data []models.EditSharedInfrastructureSettingsConnectorApplicationBlocks
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]deployment_services.EditSharedInfrastructureSettingsConnectorApplicationBlocks, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.EditSharedInfrastructureSettingsConnectorApplicationBlocks{}.AttrTypes(), &item)
		unpacked, d := unpackEditSharedInfrastructureSettingsConnectorApplicationBlocksToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.EditSharedInfrastructureSettingsConnectorApplicationBlocks", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for EditSharedInfrastructureSettingsConnectorApplicationBlocks ---
func packEditSharedInfrastructureSettingsConnectorApplicationBlocksListFromSdk(ctx context.Context, sdks []deployment_services.EditSharedInfrastructureSettingsConnectorApplicationBlocks) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.EditSharedInfrastructureSettingsConnectorApplicationBlocks")
	diags := diag.Diagnostics{}
	var data []models.EditSharedInfrastructureSettingsConnectorApplicationBlocks

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.EditSharedInfrastructureSettingsConnectorApplicationBlocks
		obj, d := packEditSharedInfrastructureSettingsConnectorApplicationBlocksFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.EditSharedInfrastructureSettingsConnectorApplicationBlocks{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.EditSharedInfrastructureSettingsConnectorApplicationBlocks", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.EditSharedInfrastructureSettingsConnectorApplicationBlocks{}.AttrType(), data)
}

// --- Unpacker for EditSharedInfrastructureSettingsConnectorConnectorBlocks ---
func unpackEditSharedInfrastructureSettingsConnectorConnectorBlocksToSdk(ctx context.Context, obj types.Object) (*deployment_services.EditSharedInfrastructureSettingsConnectorConnectorBlocks, diag.Diagnostics) {
    tflog.Debug(ctx, "Entering unpack helper for models.EditSharedInfrastructureSettingsConnectorConnectorBlocks", map[string]interface{}{"tf_object": obj})
    diags := diag.Diagnostics{}
    var model models.EditSharedInfrastructureSettingsConnectorConnectorBlocks
    diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
    if diags.HasError() {
        tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
        return nil, diags
    }
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

    var sdk deployment_services.EditSharedInfrastructureSettingsConnectorConnectorBlocks
    var d diag.Diagnostics
    // Handling Lists
    if !model.Member.IsNull() && !model.Member.IsUnknown() {
        tflog.Debug(ctx, "Unpacking list of primitives for field Member")
        diags.Append(model.Member.ElementsAs(ctx, &sdk.Member, false)...)
    }

diags.Append(d...)

    tflog.Debug(ctx, "Exiting unpack helper for models.EditSharedInfrastructureSettingsConnectorConnectorBlocks", map[string]interface{}{"has_errors": diags.HasError()})
    return &sdk, diags

}

// --- Packer for EditSharedInfrastructureSettingsConnectorConnectorBlocks ---
func packEditSharedInfrastructureSettingsConnectorConnectorBlocksFromSdk(ctx context.Context, sdk deployment_services.EditSharedInfrastructureSettingsConnectorConnectorBlocks) (types.Object, diag.Diagnostics) {
    tflog.Debug(ctx, "Entering pack helper for models.EditSharedInfrastructureSettingsConnectorConnectorBlocks", map[string]interface{}{"sdk_struct": sdk})
    diags := diag.Diagnostics{}
    var model models.EditSharedInfrastructureSettingsConnectorConnectorBlocks
    var d diag.Diagnostics
    // Handling Lists
    if sdk.Member != nil {
        tflog.Debug(ctx, "Packing list of primitives for field Member")
        var d diag.Diagnostics
        // This logic now dynamically determines the element type based on the SDK's Go type.
        var elemType attr.Type = basetypes.StringType{} // Default to string
        model.Member, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Member)
        diags.Append(d...)
    } else {
        // This logic now creates a correctly typed null list.
        var elemType attr.Type = basetypes.StringType{} // Default to string
        model.Member = basetypes.NewListNull(elemType)
    }
diags.Append(d...)

    obj, d := types.ObjectValueFrom(ctx, models.EditSharedInfrastructureSettingsConnectorConnectorBlocks{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
    diags.Append(d...)
    tflog.Debug(ctx, "Exiting pack helper for models.EditSharedInfrastructureSettingsConnectorConnectorBlocks", map[string]interface{}{"has_errors": diags.HasError()})
    return obj, diags

}

// --- List Unpacker for EditSharedInfrastructureSettingsConnectorConnectorBlocks ---
func unpackEditSharedInfrastructureSettingsConnectorConnectorBlocksListToSdk(ctx context.Context, list types.List) ([]deployment_services.EditSharedInfrastructureSettingsConnectorConnectorBlocks, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.EditSharedInfrastructureSettingsConnectorConnectorBlocks")
	diags := diag.Diagnostics{}
	var data []models.EditSharedInfrastructureSettingsConnectorConnectorBlocks
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]deployment_services.EditSharedInfrastructureSettingsConnectorConnectorBlocks, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.EditSharedInfrastructureSettingsConnectorConnectorBlocks{}.AttrTypes(), &item)
		unpacked, d := unpackEditSharedInfrastructureSettingsConnectorConnectorBlocksToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.EditSharedInfrastructureSettingsConnectorConnectorBlocks", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for EditSharedInfrastructureSettingsConnectorConnectorBlocks ---
func packEditSharedInfrastructureSettingsConnectorConnectorBlocksListFromSdk(ctx context.Context, sdks []deployment_services.EditSharedInfrastructureSettingsConnectorConnectorBlocks) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.EditSharedInfrastructureSettingsConnectorConnectorBlocks")
	diags := diag.Diagnostics{}
	var data []models.EditSharedInfrastructureSettingsConnectorConnectorBlocks

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.EditSharedInfrastructureSettingsConnectorConnectorBlocks
		obj, d := packEditSharedInfrastructureSettingsConnectorConnectorBlocksFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.EditSharedInfrastructureSettingsConnectorConnectorBlocks{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.EditSharedInfrastructureSettingsConnectorConnectorBlocks", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.EditSharedInfrastructureSettingsConnectorConnectorBlocks{}.AttrType(), data)
}


*/
