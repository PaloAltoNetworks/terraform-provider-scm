package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/device_settings"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/device_settings"
)

// --- Unpacker for SessionSettings ---
func unpackSessionSettingsToSdk(ctx context.Context, obj types.Object) (*device_settings.SessionSettings, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.SessionSettings", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.SessionSettings
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk device_settings.SessionSettings
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

	// Handling Objects
	if !model.SessionSettings.IsNull() && !model.SessionSettings.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field SessionSettings")
		unpacked, d := unpackSessionSettingsSessionSettingsToSdk(ctx, model.SessionSettings)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "SessionSettings"})
		}
		if unpacked != nil {
			sdk.SessionSettings = unpacked
		}
	}

	// Handling Primitives
	if !model.Snippet.IsNull() && !model.Snippet.IsUnknown() {
		sdk.Snippet = model.Snippet.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.SessionSettings", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for SessionSettings ---
func packSessionSettingsFromSdk(ctx context.Context, sdk device_settings.SessionSettings) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.SessionSettings", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.SessionSettings
	var d diag.Diagnostics
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
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.SessionSettings != nil {
		tflog.Debug(ctx, "Packing nested object for field SessionSettings")
		packed, d := packSessionSettingsSessionSettingsFromSdk(ctx, *sdk.SessionSettings)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "SessionSettings"})
		}
		model.SessionSettings = packed
	} else {
		model.SessionSettings = basetypes.NewObjectNull(models.SessionSettingsSessionSettings{}.AttrTypes())
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

	obj, d := types.ObjectValueFrom(ctx, models.SessionSettings{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.SessionSettings", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for SessionSettings ---
func unpackSessionSettingsListToSdk(ctx context.Context, list types.List) ([]device_settings.SessionSettings, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.SessionSettings")
	diags := diag.Diagnostics{}
	var data []models.SessionSettings
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]device_settings.SessionSettings, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.SessionSettings{}.AttrTypes(), &item)
		unpacked, d := unpackSessionSettingsToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.SessionSettings", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for SessionSettings ---
func packSessionSettingsListFromSdk(ctx context.Context, sdks []device_settings.SessionSettings) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.SessionSettings")
	diags := diag.Diagnostics{}
	var data []models.SessionSettings

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.SessionSettings
		obj, d := packSessionSettingsFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.SessionSettings{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.SessionSettings", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.SessionSettings{}.AttrType(), data)
}

// --- Unpacker for SessionSettingsSessionSettings ---
func unpackSessionSettingsSessionSettingsToSdk(ctx context.Context, obj types.Object) (*device_settings.SessionSettingsSessionSettings, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.SessionSettingsSessionSettings", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.SessionSettingsSessionSettings
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk device_settings.SessionSettingsSessionSettings
	var d diag.Diagnostics
	// Handling Primitives
	if !model.AcceleratedAgingEnable.IsNull() && !model.AcceleratedAgingEnable.IsUnknown() {
		sdk.AcceleratedAgingEnable = model.AcceleratedAgingEnable.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "AcceleratedAgingEnable", "value": *sdk.AcceleratedAgingEnable})
	}

	// Handling Primitives
	if !model.AcceleratedAgingScalingFactor.IsNull() && !model.AcceleratedAgingScalingFactor.IsUnknown() {
		val := float32(model.AcceleratedAgingScalingFactor.ValueFloat64())
		sdk.AcceleratedAgingScalingFactor = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "AcceleratedAgingScalingFactor", "value": *sdk.AcceleratedAgingScalingFactor})
	}

	// Handling Primitives
	if !model.AcceleratedAgingThreshold.IsNull() && !model.AcceleratedAgingThreshold.IsUnknown() {
		val := float32(model.AcceleratedAgingThreshold.ValueFloat64())
		sdk.AcceleratedAgingThreshold = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "AcceleratedAgingThreshold", "value": *sdk.AcceleratedAgingThreshold})
	}

	// Handling Objects
	if !model.Config.IsNull() && !model.Config.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Config")
		unpacked, d := unpackSessionSettingsSessionSettingsConfigToSdk(ctx, model.Config)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Config"})
		}
		if unpacked != nil {
			sdk.Config = unpacked
		}
	}

	// Handling Primitives
	if !model.DhcpBcastSessionOn.IsNull() && !model.DhcpBcastSessionOn.IsUnknown() {
		sdk.DhcpBcastSessionOn = model.DhcpBcastSessionOn.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DhcpBcastSessionOn", "value": *sdk.DhcpBcastSessionOn})
	}

	// Handling Primitives
	if !model.Erspan.IsNull() && !model.Erspan.IsUnknown() {
		sdk.Erspan = model.Erspan.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Erspan", "value": *sdk.Erspan})
	}

	// Handling Primitives
	if !model.IcmpUnreachableRate.IsNull() && !model.IcmpUnreachableRate.IsUnknown() {
		val := float32(model.IcmpUnreachableRate.ValueFloat64())
		sdk.IcmpUnreachableRate = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "IcmpUnreachableRate", "value": *sdk.IcmpUnreachableRate})
	}

	// Handling Objects
	if !model.Icmpv6RateLimit.IsNull() && !model.Icmpv6RateLimit.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Icmpv6RateLimit")
		unpacked, d := unpackSessionSettingsSessionSettingsIcmpv6RateLimitToSdk(ctx, model.Icmpv6RateLimit)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Icmpv6RateLimit"})
		}
		if unpacked != nil {
			sdk.Icmpv6RateLimit = unpacked
		}
	}

	// Handling Primitives
	if !model.Ipv6Firewalling.IsNull() && !model.Ipv6Firewalling.IsUnknown() {
		sdk.Ipv6Firewalling = model.Ipv6Firewalling.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Ipv6Firewalling", "value": *sdk.Ipv6Firewalling})
	}

	// Handling Objects
	if !model.JumboFrame.IsNull() && !model.JumboFrame.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field JumboFrame")
		unpacked, d := unpackSessionSettingsSessionSettingsJumboFrameToSdk(ctx, model.JumboFrame)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "JumboFrame"})
		}
		if unpacked != nil {
			sdk.JumboFrame = unpacked
		}
	}

	// Handling Primitives
	if !model.MaxPendingMcastPktsPerSession.IsNull() && !model.MaxPendingMcastPktsPerSession.IsUnknown() {
		val := float32(model.MaxPendingMcastPktsPerSession.ValueFloat64())
		sdk.MaxPendingMcastPktsPerSession = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "MaxPendingMcastPktsPerSession", "value": *sdk.MaxPendingMcastPktsPerSession})
	}

	// Handling Primitives
	if !model.MulticastRouteSetupBuffering.IsNull() && !model.MulticastRouteSetupBuffering.IsUnknown() {
		sdk.MulticastRouteSetupBuffering = model.MulticastRouteSetupBuffering.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "MulticastRouteSetupBuffering", "value": *sdk.MulticastRouteSetupBuffering})
	}

	// Handling Objects
	if !model.Nat.IsNull() && !model.Nat.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Nat")
		unpacked, d := unpackSessionSettingsSessionSettingsNatToSdk(ctx, model.Nat)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Nat"})
		}
		if unpacked != nil {
			sdk.Nat = unpacked
		}
	}

	// Handling Objects
	if !model.Nat64.IsNull() && !model.Nat64.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Nat64")
		unpacked, d := unpackSessionSettingsSessionSettingsNat64ToSdk(ctx, model.Nat64)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Nat64"})
		}
		if unpacked != nil {
			sdk.Nat64 = unpacked
		}
	}

	// Handling Primitives
	if !model.PacketBufferProtectionActivate.IsNull() && !model.PacketBufferProtectionActivate.IsUnknown() {
		val := float32(model.PacketBufferProtectionActivate.ValueFloat64())
		sdk.PacketBufferProtectionActivate = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "PacketBufferProtectionActivate", "value": *sdk.PacketBufferProtectionActivate})
	}

	// Handling Primitives
	if !model.PacketBufferProtectionAlert.IsNull() && !model.PacketBufferProtectionAlert.IsUnknown() {
		val := int32(model.PacketBufferProtectionAlert.ValueInt64())
		sdk.PacketBufferProtectionAlert = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "PacketBufferProtectionAlert", "value": *sdk.PacketBufferProtectionAlert})
	}

	// Handling Primitives
	if !model.PacketBufferProtectionBlockCountdown.IsNull() && !model.PacketBufferProtectionBlockCountdown.IsUnknown() {
		val := float32(model.PacketBufferProtectionBlockCountdown.ValueFloat64())
		sdk.PacketBufferProtectionBlockCountdown = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "PacketBufferProtectionBlockCountdown", "value": *sdk.PacketBufferProtectionBlockCountdown})
	}

	// Handling Primitives
	if !model.PacketBufferProtectionBlockDurationTime.IsNull() && !model.PacketBufferProtectionBlockDurationTime.IsUnknown() {
		val := float32(model.PacketBufferProtectionBlockDurationTime.ValueFloat64())
		sdk.PacketBufferProtectionBlockDurationTime = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "PacketBufferProtectionBlockDurationTime", "value": *sdk.PacketBufferProtectionBlockDurationTime})
	}

	// Handling Primitives
	if !model.PacketBufferProtectionBlockHoldTime.IsNull() && !model.PacketBufferProtectionBlockHoldTime.IsUnknown() {
		val := float32(model.PacketBufferProtectionBlockHoldTime.ValueFloat64())
		sdk.PacketBufferProtectionBlockHoldTime = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "PacketBufferProtectionBlockHoldTime", "value": *sdk.PacketBufferProtectionBlockHoldTime})
	}

	// Handling Primitives
	if !model.PacketBufferProtectionEnable.IsNull() && !model.PacketBufferProtectionEnable.IsUnknown() {
		sdk.PacketBufferProtectionEnable = model.PacketBufferProtectionEnable.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "PacketBufferProtectionEnable", "value": *sdk.PacketBufferProtectionEnable})
	}

	// Handling Primitives
	if !model.PacketBufferProtectionLatencyActivate.IsNull() && !model.PacketBufferProtectionLatencyActivate.IsUnknown() {
		val := float32(model.PacketBufferProtectionLatencyActivate.ValueFloat64())
		sdk.PacketBufferProtectionLatencyActivate = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "PacketBufferProtectionLatencyActivate", "value": *sdk.PacketBufferProtectionLatencyActivate})
	}

	// Handling Primitives
	if !model.PacketBufferProtectionLatencyAlert.IsNull() && !model.PacketBufferProtectionLatencyAlert.IsUnknown() {
		val := float32(model.PacketBufferProtectionLatencyAlert.ValueFloat64())
		sdk.PacketBufferProtectionLatencyAlert = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "PacketBufferProtectionLatencyAlert", "value": *sdk.PacketBufferProtectionLatencyAlert})
	}

	// Handling Primitives
	if !model.PacketBufferProtectionLatencyBlockCountdown.IsNull() && !model.PacketBufferProtectionLatencyBlockCountdown.IsUnknown() {
		val := float32(model.PacketBufferProtectionLatencyBlockCountdown.ValueFloat64())
		sdk.PacketBufferProtectionLatencyBlockCountdown = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "PacketBufferProtectionLatencyBlockCountdown", "value": *sdk.PacketBufferProtectionLatencyBlockCountdown})
	}

	// Handling Primitives
	if !model.PacketBufferProtectionLatencyMaxTolerate.IsNull() && !model.PacketBufferProtectionLatencyMaxTolerate.IsUnknown() {
		val := float32(model.PacketBufferProtectionLatencyMaxTolerate.ValueFloat64())
		sdk.PacketBufferProtectionLatencyMaxTolerate = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "PacketBufferProtectionLatencyMaxTolerate", "value": *sdk.PacketBufferProtectionLatencyMaxTolerate})
	}

	// Handling Primitives
	if !model.PacketBufferProtectionMonitorOnly.IsNull() && !model.PacketBufferProtectionMonitorOnly.IsUnknown() {
		sdk.PacketBufferProtectionMonitorOnly = model.PacketBufferProtectionMonitorOnly.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "PacketBufferProtectionMonitorOnly", "value": *sdk.PacketBufferProtectionMonitorOnly})
	}

	// Handling Primitives
	if !model.PacketBufferProtectionUseLatency.IsNull() && !model.PacketBufferProtectionUseLatency.IsUnknown() {
		sdk.PacketBufferProtectionUseLatency = model.PacketBufferProtectionUseLatency.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "PacketBufferProtectionUseLatency", "value": *sdk.PacketBufferProtectionUseLatency})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.SessionSettingsSessionSettings", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for SessionSettingsSessionSettings ---
func packSessionSettingsSessionSettingsFromSdk(ctx context.Context, sdk device_settings.SessionSettingsSessionSettings) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.SessionSettingsSessionSettings", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.SessionSettingsSessionSettings
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.AcceleratedAgingEnable != nil {
		model.AcceleratedAgingEnable = basetypes.NewBoolValue(*sdk.AcceleratedAgingEnable)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "AcceleratedAgingEnable", "value": *sdk.AcceleratedAgingEnable})
	} else {
		model.AcceleratedAgingEnable = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.AcceleratedAgingScalingFactor != nil {
		model.AcceleratedAgingScalingFactor = basetypes.NewFloat64Value(float64(*sdk.AcceleratedAgingScalingFactor))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "AcceleratedAgingScalingFactor", "value": *sdk.AcceleratedAgingScalingFactor})
	} else {
		model.AcceleratedAgingScalingFactor = basetypes.NewFloat64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.AcceleratedAgingThreshold != nil {
		model.AcceleratedAgingThreshold = basetypes.NewFloat64Value(float64(*sdk.AcceleratedAgingThreshold))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "AcceleratedAgingThreshold", "value": *sdk.AcceleratedAgingThreshold})
	} else {
		model.AcceleratedAgingThreshold = basetypes.NewFloat64Null()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Config != nil {
		tflog.Debug(ctx, "Packing nested object for field Config")
		packed, d := packSessionSettingsSessionSettingsConfigFromSdk(ctx, *sdk.Config)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Config"})
		}
		model.Config = packed
	} else {
		model.Config = basetypes.NewObjectNull(models.SessionSettingsSessionSettingsConfig{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.DhcpBcastSessionOn != nil {
		model.DhcpBcastSessionOn = basetypes.NewBoolValue(*sdk.DhcpBcastSessionOn)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DhcpBcastSessionOn", "value": *sdk.DhcpBcastSessionOn})
	} else {
		model.DhcpBcastSessionOn = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Erspan != nil {
		model.Erspan = basetypes.NewBoolValue(*sdk.Erspan)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Erspan", "value": *sdk.Erspan})
	} else {
		model.Erspan = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.IcmpUnreachableRate != nil {
		model.IcmpUnreachableRate = basetypes.NewFloat64Value(float64(*sdk.IcmpUnreachableRate))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "IcmpUnreachableRate", "value": *sdk.IcmpUnreachableRate})
	} else {
		model.IcmpUnreachableRate = basetypes.NewFloat64Null()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Icmpv6RateLimit != nil {
		tflog.Debug(ctx, "Packing nested object for field Icmpv6RateLimit")
		packed, d := packSessionSettingsSessionSettingsIcmpv6RateLimitFromSdk(ctx, *sdk.Icmpv6RateLimit)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Icmpv6RateLimit"})
		}
		model.Icmpv6RateLimit = packed
	} else {
		model.Icmpv6RateLimit = basetypes.NewObjectNull(models.SessionSettingsSessionSettingsIcmpv6RateLimit{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Ipv6Firewalling != nil {
		model.Ipv6Firewalling = basetypes.NewBoolValue(*sdk.Ipv6Firewalling)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Ipv6Firewalling", "value": *sdk.Ipv6Firewalling})
	} else {
		model.Ipv6Firewalling = basetypes.NewBoolNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.JumboFrame != nil {
		tflog.Debug(ctx, "Packing nested object for field JumboFrame")
		packed, d := packSessionSettingsSessionSettingsJumboFrameFromSdk(ctx, *sdk.JumboFrame)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "JumboFrame"})
		}
		model.JumboFrame = packed
	} else {
		model.JumboFrame = basetypes.NewObjectNull(models.SessionSettingsSessionSettingsJumboFrame{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.MaxPendingMcastPktsPerSession != nil {
		model.MaxPendingMcastPktsPerSession = basetypes.NewFloat64Value(float64(*sdk.MaxPendingMcastPktsPerSession))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "MaxPendingMcastPktsPerSession", "value": *sdk.MaxPendingMcastPktsPerSession})
	} else {
		model.MaxPendingMcastPktsPerSession = basetypes.NewFloat64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.MulticastRouteSetupBuffering != nil {
		model.MulticastRouteSetupBuffering = basetypes.NewBoolValue(*sdk.MulticastRouteSetupBuffering)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "MulticastRouteSetupBuffering", "value": *sdk.MulticastRouteSetupBuffering})
	} else {
		model.MulticastRouteSetupBuffering = basetypes.NewBoolNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Nat != nil {
		tflog.Debug(ctx, "Packing nested object for field Nat")
		packed, d := packSessionSettingsSessionSettingsNatFromSdk(ctx, *sdk.Nat)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Nat"})
		}
		model.Nat = packed
	} else {
		model.Nat = basetypes.NewObjectNull(models.SessionSettingsSessionSettingsNat{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Nat64 != nil {
		tflog.Debug(ctx, "Packing nested object for field Nat64")
		packed, d := packSessionSettingsSessionSettingsNat64FromSdk(ctx, *sdk.Nat64)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Nat64"})
		}
		model.Nat64 = packed
	} else {
		model.Nat64 = basetypes.NewObjectNull(models.SessionSettingsSessionSettingsNat64{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.PacketBufferProtectionActivate != nil {
		model.PacketBufferProtectionActivate = basetypes.NewFloat64Value(float64(*sdk.PacketBufferProtectionActivate))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "PacketBufferProtectionActivate", "value": *sdk.PacketBufferProtectionActivate})
	} else {
		model.PacketBufferProtectionActivate = basetypes.NewFloat64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.PacketBufferProtectionAlert != nil {
		model.PacketBufferProtectionAlert = basetypes.NewInt64Value(int64(*sdk.PacketBufferProtectionAlert))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "PacketBufferProtectionAlert", "value": *sdk.PacketBufferProtectionAlert})
	} else {
		model.PacketBufferProtectionAlert = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.PacketBufferProtectionBlockCountdown != nil {
		model.PacketBufferProtectionBlockCountdown = basetypes.NewFloat64Value(float64(*sdk.PacketBufferProtectionBlockCountdown))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "PacketBufferProtectionBlockCountdown", "value": *sdk.PacketBufferProtectionBlockCountdown})
	} else {
		model.PacketBufferProtectionBlockCountdown = basetypes.NewFloat64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.PacketBufferProtectionBlockDurationTime != nil {
		model.PacketBufferProtectionBlockDurationTime = basetypes.NewFloat64Value(float64(*sdk.PacketBufferProtectionBlockDurationTime))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "PacketBufferProtectionBlockDurationTime", "value": *sdk.PacketBufferProtectionBlockDurationTime})
	} else {
		model.PacketBufferProtectionBlockDurationTime = basetypes.NewFloat64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.PacketBufferProtectionBlockHoldTime != nil {
		model.PacketBufferProtectionBlockHoldTime = basetypes.NewFloat64Value(float64(*sdk.PacketBufferProtectionBlockHoldTime))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "PacketBufferProtectionBlockHoldTime", "value": *sdk.PacketBufferProtectionBlockHoldTime})
	} else {
		model.PacketBufferProtectionBlockHoldTime = basetypes.NewFloat64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.PacketBufferProtectionEnable != nil {
		model.PacketBufferProtectionEnable = basetypes.NewBoolValue(*sdk.PacketBufferProtectionEnable)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "PacketBufferProtectionEnable", "value": *sdk.PacketBufferProtectionEnable})
	} else {
		model.PacketBufferProtectionEnable = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.PacketBufferProtectionLatencyActivate != nil {
		model.PacketBufferProtectionLatencyActivate = basetypes.NewFloat64Value(float64(*sdk.PacketBufferProtectionLatencyActivate))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "PacketBufferProtectionLatencyActivate", "value": *sdk.PacketBufferProtectionLatencyActivate})
	} else {
		model.PacketBufferProtectionLatencyActivate = basetypes.NewFloat64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.PacketBufferProtectionLatencyAlert != nil {
		model.PacketBufferProtectionLatencyAlert = basetypes.NewFloat64Value(float64(*sdk.PacketBufferProtectionLatencyAlert))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "PacketBufferProtectionLatencyAlert", "value": *sdk.PacketBufferProtectionLatencyAlert})
	} else {
		model.PacketBufferProtectionLatencyAlert = basetypes.NewFloat64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.PacketBufferProtectionLatencyBlockCountdown != nil {
		model.PacketBufferProtectionLatencyBlockCountdown = basetypes.NewFloat64Value(float64(*sdk.PacketBufferProtectionLatencyBlockCountdown))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "PacketBufferProtectionLatencyBlockCountdown", "value": *sdk.PacketBufferProtectionLatencyBlockCountdown})
	} else {
		model.PacketBufferProtectionLatencyBlockCountdown = basetypes.NewFloat64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.PacketBufferProtectionLatencyMaxTolerate != nil {
		model.PacketBufferProtectionLatencyMaxTolerate = basetypes.NewFloat64Value(float64(*sdk.PacketBufferProtectionLatencyMaxTolerate))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "PacketBufferProtectionLatencyMaxTolerate", "value": *sdk.PacketBufferProtectionLatencyMaxTolerate})
	} else {
		model.PacketBufferProtectionLatencyMaxTolerate = basetypes.NewFloat64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.PacketBufferProtectionMonitorOnly != nil {
		model.PacketBufferProtectionMonitorOnly = basetypes.NewBoolValue(*sdk.PacketBufferProtectionMonitorOnly)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "PacketBufferProtectionMonitorOnly", "value": *sdk.PacketBufferProtectionMonitorOnly})
	} else {
		model.PacketBufferProtectionMonitorOnly = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.PacketBufferProtectionUseLatency != nil {
		model.PacketBufferProtectionUseLatency = basetypes.NewBoolValue(*sdk.PacketBufferProtectionUseLatency)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "PacketBufferProtectionUseLatency", "value": *sdk.PacketBufferProtectionUseLatency})
	} else {
		model.PacketBufferProtectionUseLatency = basetypes.NewBoolNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.SessionSettingsSessionSettings{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.SessionSettingsSessionSettings", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for SessionSettingsSessionSettings ---
func unpackSessionSettingsSessionSettingsListToSdk(ctx context.Context, list types.List) ([]device_settings.SessionSettingsSessionSettings, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.SessionSettingsSessionSettings")
	diags := diag.Diagnostics{}
	var data []models.SessionSettingsSessionSettings
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]device_settings.SessionSettingsSessionSettings, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.SessionSettingsSessionSettings{}.AttrTypes(), &item)
		unpacked, d := unpackSessionSettingsSessionSettingsToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.SessionSettingsSessionSettings", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for SessionSettingsSessionSettings ---
func packSessionSettingsSessionSettingsListFromSdk(ctx context.Context, sdks []device_settings.SessionSettingsSessionSettings) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.SessionSettingsSessionSettings")
	diags := diag.Diagnostics{}
	var data []models.SessionSettingsSessionSettings

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.SessionSettingsSessionSettings
		obj, d := packSessionSettingsSessionSettingsFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.SessionSettingsSessionSettings{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.SessionSettingsSessionSettings", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.SessionSettingsSessionSettings{}.AttrType(), data)
}

// --- Unpacker for SessionSettingsSessionSettingsConfig ---
func unpackSessionSettingsSessionSettingsConfigToSdk(ctx context.Context, obj types.Object) (*device_settings.SessionSettingsSessionSettingsConfig, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.SessionSettingsSessionSettingsConfig", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.SessionSettingsSessionSettingsConfig
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk device_settings.SessionSettingsSessionSettingsConfig
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Rematch.IsNull() && !model.Rematch.IsUnknown() {
		sdk.Rematch = model.Rematch.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Rematch", "value": *sdk.Rematch})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.SessionSettingsSessionSettingsConfig", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for SessionSettingsSessionSettingsConfig ---
func packSessionSettingsSessionSettingsConfigFromSdk(ctx context.Context, sdk device_settings.SessionSettingsSessionSettingsConfig) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.SessionSettingsSessionSettingsConfig", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.SessionSettingsSessionSettingsConfig
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Rematch != nil {
		model.Rematch = basetypes.NewBoolValue(*sdk.Rematch)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Rematch", "value": *sdk.Rematch})
	} else {
		model.Rematch = basetypes.NewBoolNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.SessionSettingsSessionSettingsConfig{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.SessionSettingsSessionSettingsConfig", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for SessionSettingsSessionSettingsConfig ---
func unpackSessionSettingsSessionSettingsConfigListToSdk(ctx context.Context, list types.List) ([]device_settings.SessionSettingsSessionSettingsConfig, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.SessionSettingsSessionSettingsConfig")
	diags := diag.Diagnostics{}
	var data []models.SessionSettingsSessionSettingsConfig
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]device_settings.SessionSettingsSessionSettingsConfig, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.SessionSettingsSessionSettingsConfig{}.AttrTypes(), &item)
		unpacked, d := unpackSessionSettingsSessionSettingsConfigToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.SessionSettingsSessionSettingsConfig", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for SessionSettingsSessionSettingsConfig ---
func packSessionSettingsSessionSettingsConfigListFromSdk(ctx context.Context, sdks []device_settings.SessionSettingsSessionSettingsConfig) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.SessionSettingsSessionSettingsConfig")
	diags := diag.Diagnostics{}
	var data []models.SessionSettingsSessionSettingsConfig

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.SessionSettingsSessionSettingsConfig
		obj, d := packSessionSettingsSessionSettingsConfigFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.SessionSettingsSessionSettingsConfig{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.SessionSettingsSessionSettingsConfig", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.SessionSettingsSessionSettingsConfig{}.AttrType(), data)
}

// --- Unpacker for SessionSettingsSessionSettingsIcmpv6RateLimit ---
func unpackSessionSettingsSessionSettingsIcmpv6RateLimitToSdk(ctx context.Context, obj types.Object) (*device_settings.SessionSettingsSessionSettingsIcmpv6RateLimit, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.SessionSettingsSessionSettingsIcmpv6RateLimit", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.SessionSettingsSessionSettingsIcmpv6RateLimit
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk device_settings.SessionSettingsSessionSettingsIcmpv6RateLimit
	var d diag.Diagnostics
	// Handling Primitives
	if !model.BucketSize.IsNull() && !model.BucketSize.IsUnknown() {
		val := int32(model.BucketSize.ValueInt64())
		sdk.BucketSize = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "BucketSize", "value": *sdk.BucketSize})
	}

	// Handling Primitives
	if !model.PacketRate.IsNull() && !model.PacketRate.IsUnknown() {
		val := int32(model.PacketRate.ValueInt64())
		sdk.PacketRate = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "PacketRate", "value": *sdk.PacketRate})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.SessionSettingsSessionSettingsIcmpv6RateLimit", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for SessionSettingsSessionSettingsIcmpv6RateLimit ---
func packSessionSettingsSessionSettingsIcmpv6RateLimitFromSdk(ctx context.Context, sdk device_settings.SessionSettingsSessionSettingsIcmpv6RateLimit) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.SessionSettingsSessionSettingsIcmpv6RateLimit", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.SessionSettingsSessionSettingsIcmpv6RateLimit
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.BucketSize != nil {
		model.BucketSize = basetypes.NewInt64Value(int64(*sdk.BucketSize))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "BucketSize", "value": *sdk.BucketSize})
	} else {
		model.BucketSize = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.PacketRate != nil {
		model.PacketRate = basetypes.NewInt64Value(int64(*sdk.PacketRate))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "PacketRate", "value": *sdk.PacketRate})
	} else {
		model.PacketRate = basetypes.NewInt64Null()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.SessionSettingsSessionSettingsIcmpv6RateLimit{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.SessionSettingsSessionSettingsIcmpv6RateLimit", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for SessionSettingsSessionSettingsIcmpv6RateLimit ---
func unpackSessionSettingsSessionSettingsIcmpv6RateLimitListToSdk(ctx context.Context, list types.List) ([]device_settings.SessionSettingsSessionSettingsIcmpv6RateLimit, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.SessionSettingsSessionSettingsIcmpv6RateLimit")
	diags := diag.Diagnostics{}
	var data []models.SessionSettingsSessionSettingsIcmpv6RateLimit
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]device_settings.SessionSettingsSessionSettingsIcmpv6RateLimit, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.SessionSettingsSessionSettingsIcmpv6RateLimit{}.AttrTypes(), &item)
		unpacked, d := unpackSessionSettingsSessionSettingsIcmpv6RateLimitToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.SessionSettingsSessionSettingsIcmpv6RateLimit", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for SessionSettingsSessionSettingsIcmpv6RateLimit ---
func packSessionSettingsSessionSettingsIcmpv6RateLimitListFromSdk(ctx context.Context, sdks []device_settings.SessionSettingsSessionSettingsIcmpv6RateLimit) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.SessionSettingsSessionSettingsIcmpv6RateLimit")
	diags := diag.Diagnostics{}
	var data []models.SessionSettingsSessionSettingsIcmpv6RateLimit

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.SessionSettingsSessionSettingsIcmpv6RateLimit
		obj, d := packSessionSettingsSessionSettingsIcmpv6RateLimitFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.SessionSettingsSessionSettingsIcmpv6RateLimit{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.SessionSettingsSessionSettingsIcmpv6RateLimit", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.SessionSettingsSessionSettingsIcmpv6RateLimit{}.AttrType(), data)
}

// --- Unpacker for SessionSettingsSessionSettingsJumboFrame ---
func unpackSessionSettingsSessionSettingsJumboFrameToSdk(ctx context.Context, obj types.Object) (*device_settings.SessionSettingsSessionSettingsJumboFrame, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.SessionSettingsSessionSettingsJumboFrame", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.SessionSettingsSessionSettingsJumboFrame
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk device_settings.SessionSettingsSessionSettingsJumboFrame
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Mtu.IsNull() && !model.Mtu.IsUnknown() {
		val := int32(model.Mtu.ValueInt64())
		sdk.Mtu = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Mtu", "value": *sdk.Mtu})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.SessionSettingsSessionSettingsJumboFrame", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for SessionSettingsSessionSettingsJumboFrame ---
func packSessionSettingsSessionSettingsJumboFrameFromSdk(ctx context.Context, sdk device_settings.SessionSettingsSessionSettingsJumboFrame) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.SessionSettingsSessionSettingsJumboFrame", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.SessionSettingsSessionSettingsJumboFrame
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Mtu != nil {
		model.Mtu = basetypes.NewInt64Value(int64(*sdk.Mtu))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Mtu", "value": *sdk.Mtu})
	} else {
		model.Mtu = basetypes.NewInt64Null()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.SessionSettingsSessionSettingsJumboFrame{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.SessionSettingsSessionSettingsJumboFrame", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for SessionSettingsSessionSettingsJumboFrame ---
func unpackSessionSettingsSessionSettingsJumboFrameListToSdk(ctx context.Context, list types.List) ([]device_settings.SessionSettingsSessionSettingsJumboFrame, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.SessionSettingsSessionSettingsJumboFrame")
	diags := diag.Diagnostics{}
	var data []models.SessionSettingsSessionSettingsJumboFrame
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]device_settings.SessionSettingsSessionSettingsJumboFrame, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.SessionSettingsSessionSettingsJumboFrame{}.AttrTypes(), &item)
		unpacked, d := unpackSessionSettingsSessionSettingsJumboFrameToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.SessionSettingsSessionSettingsJumboFrame", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for SessionSettingsSessionSettingsJumboFrame ---
func packSessionSettingsSessionSettingsJumboFrameListFromSdk(ctx context.Context, sdks []device_settings.SessionSettingsSessionSettingsJumboFrame) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.SessionSettingsSessionSettingsJumboFrame")
	diags := diag.Diagnostics{}
	var data []models.SessionSettingsSessionSettingsJumboFrame

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.SessionSettingsSessionSettingsJumboFrame
		obj, d := packSessionSettingsSessionSettingsJumboFrameFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.SessionSettingsSessionSettingsJumboFrame{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.SessionSettingsSessionSettingsJumboFrame", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.SessionSettingsSessionSettingsJumboFrame{}.AttrType(), data)
}

// --- Unpacker for SessionSettingsSessionSettingsNat ---
func unpackSessionSettingsSessionSettingsNatToSdk(ctx context.Context, obj types.Object) (*device_settings.SessionSettingsSessionSettingsNat, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.SessionSettingsSessionSettingsNat", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.SessionSettingsSessionSettingsNat
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk device_settings.SessionSettingsSessionSettingsNat
	var d diag.Diagnostics
	// Handling Primitives
	if !model.DippOversub.IsNull() && !model.DippOversub.IsUnknown() {
		sdk.DippOversub = model.DippOversub.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DippOversub", "value": *sdk.DippOversub})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.SessionSettingsSessionSettingsNat", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for SessionSettingsSessionSettingsNat ---
func packSessionSettingsSessionSettingsNatFromSdk(ctx context.Context, sdk device_settings.SessionSettingsSessionSettingsNat) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.SessionSettingsSessionSettingsNat", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.SessionSettingsSessionSettingsNat
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.DippOversub != nil {
		model.DippOversub = basetypes.NewStringValue(*sdk.DippOversub)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DippOversub", "value": *sdk.DippOversub})
	} else {
		model.DippOversub = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.SessionSettingsSessionSettingsNat{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.SessionSettingsSessionSettingsNat", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for SessionSettingsSessionSettingsNat ---
func unpackSessionSettingsSessionSettingsNatListToSdk(ctx context.Context, list types.List) ([]device_settings.SessionSettingsSessionSettingsNat, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.SessionSettingsSessionSettingsNat")
	diags := diag.Diagnostics{}
	var data []models.SessionSettingsSessionSettingsNat
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]device_settings.SessionSettingsSessionSettingsNat, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.SessionSettingsSessionSettingsNat{}.AttrTypes(), &item)
		unpacked, d := unpackSessionSettingsSessionSettingsNatToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.SessionSettingsSessionSettingsNat", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for SessionSettingsSessionSettingsNat ---
func packSessionSettingsSessionSettingsNatListFromSdk(ctx context.Context, sdks []device_settings.SessionSettingsSessionSettingsNat) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.SessionSettingsSessionSettingsNat")
	diags := diag.Diagnostics{}
	var data []models.SessionSettingsSessionSettingsNat

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.SessionSettingsSessionSettingsNat
		obj, d := packSessionSettingsSessionSettingsNatFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.SessionSettingsSessionSettingsNat{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.SessionSettingsSessionSettingsNat", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.SessionSettingsSessionSettingsNat{}.AttrType(), data)
}

// --- Unpacker for SessionSettingsSessionSettingsNat64 ---
func unpackSessionSettingsSessionSettingsNat64ToSdk(ctx context.Context, obj types.Object) (*device_settings.SessionSettingsSessionSettingsNat64, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.SessionSettingsSessionSettingsNat64", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.SessionSettingsSessionSettingsNat64
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk device_settings.SessionSettingsSessionSettingsNat64
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Ipv6MinNetworkMtu.IsNull() && !model.Ipv6MinNetworkMtu.IsUnknown() {
		val := int32(model.Ipv6MinNetworkMtu.ValueInt64())
		sdk.Ipv6MinNetworkMtu = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Ipv6MinNetworkMtu", "value": *sdk.Ipv6MinNetworkMtu})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.SessionSettingsSessionSettingsNat64", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for SessionSettingsSessionSettingsNat64 ---
func packSessionSettingsSessionSettingsNat64FromSdk(ctx context.Context, sdk device_settings.SessionSettingsSessionSettingsNat64) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.SessionSettingsSessionSettingsNat64", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.SessionSettingsSessionSettingsNat64
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Ipv6MinNetworkMtu != nil {
		model.Ipv6MinNetworkMtu = basetypes.NewInt64Value(int64(*sdk.Ipv6MinNetworkMtu))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Ipv6MinNetworkMtu", "value": *sdk.Ipv6MinNetworkMtu})
	} else {
		model.Ipv6MinNetworkMtu = basetypes.NewInt64Null()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.SessionSettingsSessionSettingsNat64{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.SessionSettingsSessionSettingsNat64", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for SessionSettingsSessionSettingsNat64 ---
func unpackSessionSettingsSessionSettingsNat64ListToSdk(ctx context.Context, list types.List) ([]device_settings.SessionSettingsSessionSettingsNat64, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.SessionSettingsSessionSettingsNat64")
	diags := diag.Diagnostics{}
	var data []models.SessionSettingsSessionSettingsNat64
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]device_settings.SessionSettingsSessionSettingsNat64, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.SessionSettingsSessionSettingsNat64{}.AttrTypes(), &item)
		unpacked, d := unpackSessionSettingsSessionSettingsNat64ToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.SessionSettingsSessionSettingsNat64", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for SessionSettingsSessionSettingsNat64 ---
func packSessionSettingsSessionSettingsNat64ListFromSdk(ctx context.Context, sdks []device_settings.SessionSettingsSessionSettingsNat64) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.SessionSettingsSessionSettingsNat64")
	diags := diag.Diagnostics{}
	var data []models.SessionSettingsSessionSettingsNat64

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.SessionSettingsSessionSettingsNat64
		obj, d := packSessionSettingsSessionSettingsNat64FromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.SessionSettingsSessionSettingsNat64{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.SessionSettingsSessionSettingsNat64", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.SessionSettingsSessionSettingsNat64{}.AttrType(), data)
}
