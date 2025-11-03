package provider

import (
	"context"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/network_services"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/network_services"
)

// --- Unpacker for ZoneProtectionProfiles ---
func unpackZoneProtectionProfilesToSdk(ctx context.Context, obj types.Object) (*network_services.ZoneProtectionProfiles, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ZoneProtectionProfiles", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ZoneProtectionProfiles
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.ZoneProtectionProfiles
	var d diag.Diagnostics

	// Handling Primitives
	if !model.AsymmetricPath.IsNull() && !model.AsymmetricPath.IsUnknown() {
		sdk.AsymmetricPath = model.AsymmetricPath.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "AsymmetricPath", "value": *sdk.AsymmetricPath})
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
	if !model.DiscardIcmpEmbeddedError.IsNull() && !model.DiscardIcmpEmbeddedError.IsUnknown() {
		sdk.DiscardIcmpEmbeddedError = model.DiscardIcmpEmbeddedError.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DiscardIcmpEmbeddedError", "value": *sdk.DiscardIcmpEmbeddedError})
	}

	// Handling Objects
	if !model.Flood.IsNull() && !model.Flood.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Flood")
		unpacked, d := unpackZoneProtectionProfilesFloodToSdk(ctx, model.Flood)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Flood"})
		}
		if unpacked != nil {
			sdk.Flood = unpacked
		}
	}

	// Handling Primitives
	if !model.Folder.IsNull() && !model.Folder.IsUnknown() {
		sdk.Folder = model.Folder.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Folder", "value": *sdk.Folder})
	}

	// Handling Primitives
	if !model.FragmentedTrafficDiscard.IsNull() && !model.FragmentedTrafficDiscard.IsUnknown() {
		sdk.FragmentedTrafficDiscard = model.FragmentedTrafficDiscard.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "FragmentedTrafficDiscard", "value": *sdk.FragmentedTrafficDiscard})
	}

	// Handling Primitives
	if !model.IcmpFragDiscard.IsNull() && !model.IcmpFragDiscard.IsUnknown() {
		sdk.IcmpFragDiscard = model.IcmpFragDiscard.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "IcmpFragDiscard", "value": *sdk.IcmpFragDiscard})
	}

	// Handling Primitives
	if !model.IcmpLargePacketDiscard.IsNull() && !model.IcmpLargePacketDiscard.IsUnknown() {
		sdk.IcmpLargePacketDiscard = model.IcmpLargePacketDiscard.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "IcmpLargePacketDiscard", "value": *sdk.IcmpLargePacketDiscard})
	}

	// Handling Primitives
	if !model.IcmpPingZeroIdDiscard.IsNull() && !model.IcmpPingZeroIdDiscard.IsUnknown() {
		sdk.IcmpPingZeroIdDiscard = model.IcmpPingZeroIdDiscard.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "IcmpPingZeroIdDiscard", "value": *sdk.IcmpPingZeroIdDiscard})
	}

	// Handling Primitives
	if !model.Id.IsNull() && !model.Id.IsUnknown() {
		sdk.Id = model.Id.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Id", "value": *sdk.Id})
	}

	// Handling Objects
	if !model.Ipv6.IsNull() && !model.Ipv6.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Ipv6")
		unpacked, d := unpackZoneProtectionProfilesIpv6ToSdk(ctx, model.Ipv6)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Ipv6"})
		}
		if unpacked != nil {
			sdk.Ipv6 = unpacked
		}
	}

	// Handling Objects
	if !model.L2SecGroupTagProtection.IsNull() && !model.L2SecGroupTagProtection.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field L2SecGroupTagProtection")
		unpacked, d := unpackZoneProtectionProfilesL2SecGroupTagProtectionToSdk(ctx, model.L2SecGroupTagProtection)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "L2SecGroupTagProtection"})
		}
		if unpacked != nil {
			sdk.L2SecGroupTagProtection = unpacked
		}
	}

	// Handling Primitives
	if !model.LooseSourceRoutingDiscard.IsNull() && !model.LooseSourceRoutingDiscard.IsUnknown() {
		sdk.LooseSourceRoutingDiscard = model.LooseSourceRoutingDiscard.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "LooseSourceRoutingDiscard", "value": *sdk.LooseSourceRoutingDiscard})
	}

	// Handling Primitives
	if !model.MalformedOptionDiscard.IsNull() && !model.MalformedOptionDiscard.IsUnknown() {
		sdk.MalformedOptionDiscard = model.MalformedOptionDiscard.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "MalformedOptionDiscard", "value": *sdk.MalformedOptionDiscard})
	}

	// Handling Primitives
	if !model.MismatchedOverlappingTcpSegmentDiscard.IsNull() && !model.MismatchedOverlappingTcpSegmentDiscard.IsUnknown() {
		sdk.MismatchedOverlappingTcpSegmentDiscard = model.MismatchedOverlappingTcpSegmentDiscard.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "MismatchedOverlappingTcpSegmentDiscard", "value": *sdk.MismatchedOverlappingTcpSegmentDiscard})
	}

	// Handling Primitives
	if !model.MptcpOptionStrip.IsNull() && !model.MptcpOptionStrip.IsUnknown() {
		sdk.MptcpOptionStrip = model.MptcpOptionStrip.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "MptcpOptionStrip", "value": *sdk.MptcpOptionStrip})
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	// Handling Objects
	if !model.NonIpProtocol.IsNull() && !model.NonIpProtocol.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field NonIpProtocol")
		unpacked, d := unpackZoneProtectionProfilesNonIpProtocolToSdk(ctx, model.NonIpProtocol)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "NonIpProtocol"})
		}
		if unpacked != nil {
			sdk.NonIpProtocol = unpacked
		}
	}

	// Handling Primitives
	if !model.RecordRouteDiscard.IsNull() && !model.RecordRouteDiscard.IsUnknown() {
		sdk.RecordRouteDiscard = model.RecordRouteDiscard.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "RecordRouteDiscard", "value": *sdk.RecordRouteDiscard})
	}

	// Handling Primitives
	if !model.RejectNonSynTcp.IsNull() && !model.RejectNonSynTcp.IsUnknown() {
		sdk.RejectNonSynTcp = model.RejectNonSynTcp.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "RejectNonSynTcp", "value": *sdk.RejectNonSynTcp})
	}

	// Handling Lists
	if !model.Scan.IsNull() && !model.Scan.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field Scan")
		unpacked, d := unpackZoneProtectionProfilesScanInnerListToSdk(ctx, model.Scan)
		diags.Append(d...)
		sdk.Scan = unpacked
	}

	// Handling Lists
	if !model.ScanWhiteList.IsNull() && !model.ScanWhiteList.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field ScanWhiteList")
		unpacked, d := unpackZoneProtectionProfilesScanWhiteListInnerListToSdk(ctx, model.ScanWhiteList)
		diags.Append(d...)
		sdk.ScanWhiteList = unpacked
	}

	// Handling Primitives
	if !model.SecurityDiscard.IsNull() && !model.SecurityDiscard.IsUnknown() {
		sdk.SecurityDiscard = model.SecurityDiscard.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "SecurityDiscard", "value": *sdk.SecurityDiscard})
	}

	// Handling Primitives
	if !model.Snippet.IsNull() && !model.Snippet.IsUnknown() {
		sdk.Snippet = model.Snippet.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	}

	// Handling Primitives
	if !model.SpoofedIpDiscard.IsNull() && !model.SpoofedIpDiscard.IsUnknown() {
		sdk.SpoofedIpDiscard = model.SpoofedIpDiscard.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "SpoofedIpDiscard", "value": *sdk.SpoofedIpDiscard})
	}

	// Handling Primitives
	if !model.StreamIdDiscard.IsNull() && !model.StreamIdDiscard.IsUnknown() {
		sdk.StreamIdDiscard = model.StreamIdDiscard.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "StreamIdDiscard", "value": *sdk.StreamIdDiscard})
	}

	// Handling Primitives
	if !model.StrictIpCheck.IsNull() && !model.StrictIpCheck.IsUnknown() {
		sdk.StrictIpCheck = model.StrictIpCheck.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "StrictIpCheck", "value": *sdk.StrictIpCheck})
	}

	// Handling Primitives
	if !model.StrictSourceRoutingDiscard.IsNull() && !model.StrictSourceRoutingDiscard.IsUnknown() {
		sdk.StrictSourceRoutingDiscard = model.StrictSourceRoutingDiscard.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "StrictSourceRoutingDiscard", "value": *sdk.StrictSourceRoutingDiscard})
	}

	// Handling Primitives
	if !model.SuppressIcmpNeedfrag.IsNull() && !model.SuppressIcmpNeedfrag.IsUnknown() {
		sdk.SuppressIcmpNeedfrag = model.SuppressIcmpNeedfrag.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "SuppressIcmpNeedfrag", "value": *sdk.SuppressIcmpNeedfrag})
	}

	// Handling Primitives
	if !model.SuppressIcmpTimeexceeded.IsNull() && !model.SuppressIcmpTimeexceeded.IsUnknown() {
		sdk.SuppressIcmpTimeexceeded = model.SuppressIcmpTimeexceeded.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "SuppressIcmpTimeexceeded", "value": *sdk.SuppressIcmpTimeexceeded})
	}

	// Handling Primitives
	if !model.TcpFastOpenAndDataStrip.IsNull() && !model.TcpFastOpenAndDataStrip.IsUnknown() {
		sdk.TcpFastOpenAndDataStrip = model.TcpFastOpenAndDataStrip.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "TcpFastOpenAndDataStrip", "value": *sdk.TcpFastOpenAndDataStrip})
	}

	// Handling Primitives
	if !model.TcpHandshakeDiscard.IsNull() && !model.TcpHandshakeDiscard.IsUnknown() {
		sdk.TcpHandshakeDiscard = model.TcpHandshakeDiscard.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "TcpHandshakeDiscard", "value": *sdk.TcpHandshakeDiscard})
	}

	// Handling Primitives
	if !model.TcpSynWithDataDiscard.IsNull() && !model.TcpSynWithDataDiscard.IsUnknown() {
		sdk.TcpSynWithDataDiscard = model.TcpSynWithDataDiscard.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "TcpSynWithDataDiscard", "value": *sdk.TcpSynWithDataDiscard})
	}

	// Handling Primitives
	if !model.TcpSynackWithDataDiscard.IsNull() && !model.TcpSynackWithDataDiscard.IsUnknown() {
		sdk.TcpSynackWithDataDiscard = model.TcpSynackWithDataDiscard.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "TcpSynackWithDataDiscard", "value": *sdk.TcpSynackWithDataDiscard})
	}

	// Handling Primitives
	if !model.TcpTimestampStrip.IsNull() && !model.TcpTimestampStrip.IsUnknown() {
		sdk.TcpTimestampStrip = model.TcpTimestampStrip.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "TcpTimestampStrip", "value": *sdk.TcpTimestampStrip})
	}

	// Handling Primitives
	if !model.TimestampDiscard.IsNull() && !model.TimestampDiscard.IsUnknown() {
		sdk.TimestampDiscard = model.TimestampDiscard.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "TimestampDiscard", "value": *sdk.TimestampDiscard})
	}

	// Handling Primitives
	if !model.UnknownOptionDiscard.IsNull() && !model.UnknownOptionDiscard.IsUnknown() {
		sdk.UnknownOptionDiscard = model.UnknownOptionDiscard.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "UnknownOptionDiscard", "value": *sdk.UnknownOptionDiscard})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ZoneProtectionProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ZoneProtectionProfiles ---
func packZoneProtectionProfilesFromSdk(ctx context.Context, sdk network_services.ZoneProtectionProfiles) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ZoneProtectionProfiles", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ZoneProtectionProfiles
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.AsymmetricPath != nil {
		model.AsymmetricPath = basetypes.NewStringValue(*sdk.AsymmetricPath)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "AsymmetricPath", "value": *sdk.AsymmetricPath})
	} else {
		model.AsymmetricPath = basetypes.NewStringNull()
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
	if sdk.DiscardIcmpEmbeddedError != nil {
		model.DiscardIcmpEmbeddedError = basetypes.NewBoolValue(*sdk.DiscardIcmpEmbeddedError)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DiscardIcmpEmbeddedError", "value": *sdk.DiscardIcmpEmbeddedError})
	} else {
		model.DiscardIcmpEmbeddedError = basetypes.NewBoolNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Flood != nil {
		tflog.Debug(ctx, "Packing nested object for field Flood")
		packed, d := packZoneProtectionProfilesFloodFromSdk(ctx, *sdk.Flood)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Flood"})
		}
		model.Flood = packed
	} else {
		model.Flood = basetypes.NewObjectNull(models.ZoneProtectionProfilesFlood{}.AttrTypes())
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
	if sdk.FragmentedTrafficDiscard != nil {
		model.FragmentedTrafficDiscard = basetypes.NewBoolValue(*sdk.FragmentedTrafficDiscard)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "FragmentedTrafficDiscard", "value": *sdk.FragmentedTrafficDiscard})
	} else {
		model.FragmentedTrafficDiscard = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.IcmpFragDiscard != nil {
		model.IcmpFragDiscard = basetypes.NewBoolValue(*sdk.IcmpFragDiscard)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "IcmpFragDiscard", "value": *sdk.IcmpFragDiscard})
	} else {
		model.IcmpFragDiscard = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.IcmpLargePacketDiscard != nil {
		model.IcmpLargePacketDiscard = basetypes.NewBoolValue(*sdk.IcmpLargePacketDiscard)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "IcmpLargePacketDiscard", "value": *sdk.IcmpLargePacketDiscard})
	} else {
		model.IcmpLargePacketDiscard = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.IcmpPingZeroIdDiscard != nil {
		model.IcmpPingZeroIdDiscard = basetypes.NewBoolValue(*sdk.IcmpPingZeroIdDiscard)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "IcmpPingZeroIdDiscard", "value": *sdk.IcmpPingZeroIdDiscard})
	} else {
		model.IcmpPingZeroIdDiscard = basetypes.NewBoolNull()
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
	if sdk.Ipv6 != nil {
		tflog.Debug(ctx, "Packing nested object for field Ipv6")
		packed, d := packZoneProtectionProfilesIpv6FromSdk(ctx, *sdk.Ipv6)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Ipv6"})
		}
		model.Ipv6 = packed
	} else {
		model.Ipv6 = basetypes.NewObjectNull(models.ZoneProtectionProfilesIpv6{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.L2SecGroupTagProtection != nil {
		tflog.Debug(ctx, "Packing nested object for field L2SecGroupTagProtection")
		packed, d := packZoneProtectionProfilesL2SecGroupTagProtectionFromSdk(ctx, *sdk.L2SecGroupTagProtection)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "L2SecGroupTagProtection"})
		}
		model.L2SecGroupTagProtection = packed
	} else {
		model.L2SecGroupTagProtection = basetypes.NewObjectNull(models.ZoneProtectionProfilesL2SecGroupTagProtection{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.LooseSourceRoutingDiscard != nil {
		model.LooseSourceRoutingDiscard = basetypes.NewBoolValue(*sdk.LooseSourceRoutingDiscard)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "LooseSourceRoutingDiscard", "value": *sdk.LooseSourceRoutingDiscard})
	} else {
		model.LooseSourceRoutingDiscard = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.MalformedOptionDiscard != nil {
		model.MalformedOptionDiscard = basetypes.NewBoolValue(*sdk.MalformedOptionDiscard)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "MalformedOptionDiscard", "value": *sdk.MalformedOptionDiscard})
	} else {
		model.MalformedOptionDiscard = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.MismatchedOverlappingTcpSegmentDiscard != nil {
		model.MismatchedOverlappingTcpSegmentDiscard = basetypes.NewBoolValue(*sdk.MismatchedOverlappingTcpSegmentDiscard)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "MismatchedOverlappingTcpSegmentDiscard", "value": *sdk.MismatchedOverlappingTcpSegmentDiscard})
	} else {
		model.MismatchedOverlappingTcpSegmentDiscard = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.MptcpOptionStrip != nil {
		model.MptcpOptionStrip = basetypes.NewStringValue(*sdk.MptcpOptionStrip)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "MptcpOptionStrip", "value": *sdk.MptcpOptionStrip})
	} else {
		model.MptcpOptionStrip = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.NonIpProtocol != nil {
		tflog.Debug(ctx, "Packing nested object for field NonIpProtocol")
		packed, d := packZoneProtectionProfilesNonIpProtocolFromSdk(ctx, *sdk.NonIpProtocol)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "NonIpProtocol"})
		}
		model.NonIpProtocol = packed
	} else {
		model.NonIpProtocol = basetypes.NewObjectNull(models.ZoneProtectionProfilesNonIpProtocol{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.RecordRouteDiscard != nil {
		model.RecordRouteDiscard = basetypes.NewBoolValue(*sdk.RecordRouteDiscard)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "RecordRouteDiscard", "value": *sdk.RecordRouteDiscard})
	} else {
		model.RecordRouteDiscard = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.RejectNonSynTcp != nil {
		model.RejectNonSynTcp = basetypes.NewStringValue(*sdk.RejectNonSynTcp)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "RejectNonSynTcp", "value": *sdk.RejectNonSynTcp})
	} else {
		model.RejectNonSynTcp = basetypes.NewStringNull()
	}
	// Handling Lists
	if sdk.Scan != nil {
		tflog.Debug(ctx, "Packing list of objects for field Scan")
		packed, d := packZoneProtectionProfilesScanInnerListFromSdk(ctx, sdk.Scan)
		diags.Append(d...)
		model.Scan = packed
	} else {
		model.Scan = basetypes.NewListNull(models.ZoneProtectionProfilesScanInner{}.AttrType())
	}
	// Handling Lists
	if sdk.ScanWhiteList != nil {
		tflog.Debug(ctx, "Packing list of objects for field ScanWhiteList")
		packed, d := packZoneProtectionProfilesScanWhiteListInnerListFromSdk(ctx, sdk.ScanWhiteList)
		diags.Append(d...)
		model.ScanWhiteList = packed
	} else {
		model.ScanWhiteList = basetypes.NewListNull(models.ZoneProtectionProfilesScanWhiteListInner{}.AttrType())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.SecurityDiscard != nil {
		model.SecurityDiscard = basetypes.NewBoolValue(*sdk.SecurityDiscard)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "SecurityDiscard", "value": *sdk.SecurityDiscard})
	} else {
		model.SecurityDiscard = basetypes.NewBoolNull()
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
	if sdk.SpoofedIpDiscard != nil {
		model.SpoofedIpDiscard = basetypes.NewBoolValue(*sdk.SpoofedIpDiscard)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "SpoofedIpDiscard", "value": *sdk.SpoofedIpDiscard})
	} else {
		model.SpoofedIpDiscard = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.StreamIdDiscard != nil {
		model.StreamIdDiscard = basetypes.NewBoolValue(*sdk.StreamIdDiscard)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "StreamIdDiscard", "value": *sdk.StreamIdDiscard})
	} else {
		model.StreamIdDiscard = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.StrictIpCheck != nil {
		model.StrictIpCheck = basetypes.NewBoolValue(*sdk.StrictIpCheck)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "StrictIpCheck", "value": *sdk.StrictIpCheck})
	} else {
		model.StrictIpCheck = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.StrictSourceRoutingDiscard != nil {
		model.StrictSourceRoutingDiscard = basetypes.NewBoolValue(*sdk.StrictSourceRoutingDiscard)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "StrictSourceRoutingDiscard", "value": *sdk.StrictSourceRoutingDiscard})
	} else {
		model.StrictSourceRoutingDiscard = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.SuppressIcmpNeedfrag != nil {
		model.SuppressIcmpNeedfrag = basetypes.NewBoolValue(*sdk.SuppressIcmpNeedfrag)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "SuppressIcmpNeedfrag", "value": *sdk.SuppressIcmpNeedfrag})
	} else {
		model.SuppressIcmpNeedfrag = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.SuppressIcmpTimeexceeded != nil {
		model.SuppressIcmpTimeexceeded = basetypes.NewBoolValue(*sdk.SuppressIcmpTimeexceeded)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "SuppressIcmpTimeexceeded", "value": *sdk.SuppressIcmpTimeexceeded})
	} else {
		model.SuppressIcmpTimeexceeded = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.TcpFastOpenAndDataStrip != nil {
		model.TcpFastOpenAndDataStrip = basetypes.NewBoolValue(*sdk.TcpFastOpenAndDataStrip)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "TcpFastOpenAndDataStrip", "value": *sdk.TcpFastOpenAndDataStrip})
	} else {
		model.TcpFastOpenAndDataStrip = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.TcpHandshakeDiscard != nil {
		model.TcpHandshakeDiscard = basetypes.NewBoolValue(*sdk.TcpHandshakeDiscard)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "TcpHandshakeDiscard", "value": *sdk.TcpHandshakeDiscard})
	} else {
		model.TcpHandshakeDiscard = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.TcpSynWithDataDiscard != nil {
		model.TcpSynWithDataDiscard = basetypes.NewBoolValue(*sdk.TcpSynWithDataDiscard)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "TcpSynWithDataDiscard", "value": *sdk.TcpSynWithDataDiscard})
	} else {
		model.TcpSynWithDataDiscard = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.TcpSynackWithDataDiscard != nil {
		model.TcpSynackWithDataDiscard = basetypes.NewBoolValue(*sdk.TcpSynackWithDataDiscard)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "TcpSynackWithDataDiscard", "value": *sdk.TcpSynackWithDataDiscard})
	} else {
		model.TcpSynackWithDataDiscard = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.TcpTimestampStrip != nil {
		model.TcpTimestampStrip = basetypes.NewBoolValue(*sdk.TcpTimestampStrip)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "TcpTimestampStrip", "value": *sdk.TcpTimestampStrip})
	} else {
		model.TcpTimestampStrip = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.TimestampDiscard != nil {
		model.TimestampDiscard = basetypes.NewBoolValue(*sdk.TimestampDiscard)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "TimestampDiscard", "value": *sdk.TimestampDiscard})
	} else {
		model.TimestampDiscard = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.UnknownOptionDiscard != nil {
		model.UnknownOptionDiscard = basetypes.NewBoolValue(*sdk.UnknownOptionDiscard)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "UnknownOptionDiscard", "value": *sdk.UnknownOptionDiscard})
	} else {
		model.UnknownOptionDiscard = basetypes.NewBoolNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ZoneProtectionProfiles{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ZoneProtectionProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ZoneProtectionProfiles ---
func unpackZoneProtectionProfilesListToSdk(ctx context.Context, list types.List) ([]network_services.ZoneProtectionProfiles, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ZoneProtectionProfiles")
	diags := diag.Diagnostics{}
	var data []models.ZoneProtectionProfiles
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.ZoneProtectionProfiles, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ZoneProtectionProfiles{}.AttrTypes(), &item)
		unpacked, d := unpackZoneProtectionProfilesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ZoneProtectionProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ZoneProtectionProfiles ---
func packZoneProtectionProfilesListFromSdk(ctx context.Context, sdks []network_services.ZoneProtectionProfiles) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ZoneProtectionProfiles")
	diags := diag.Diagnostics{}
	var data []models.ZoneProtectionProfiles

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ZoneProtectionProfiles
		obj, d := packZoneProtectionProfilesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ZoneProtectionProfiles{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ZoneProtectionProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ZoneProtectionProfiles{}.AttrType(), data)
}

// --- Unpacker for ZoneProtectionProfilesFlood ---
func unpackZoneProtectionProfilesFloodToSdk(ctx context.Context, obj types.Object) (*network_services.ZoneProtectionProfilesFlood, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ZoneProtectionProfilesFlood", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ZoneProtectionProfilesFlood
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.ZoneProtectionProfilesFlood
	var d diag.Diagnostics
	// Handling Objects
	if !model.Icmp.IsNull() && !model.Icmp.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Icmp")
		unpacked, d := unpackZoneProtectionProfilesFloodIcmpToSdk(ctx, model.Icmp)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Icmp"})
		}
		if unpacked != nil {
			sdk.Icmp = unpacked
		}
	}

	// Handling Objects
	if !model.Icmpv6.IsNull() && !model.Icmpv6.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Icmpv6")
		unpacked, d := unpackZoneProtectionProfilesFloodIcmpv6ToSdk(ctx, model.Icmpv6)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Icmpv6"})
		}
		if unpacked != nil {
			sdk.Icmpv6 = unpacked
		}
	}

	// Handling Objects
	if !model.OtherIp.IsNull() && !model.OtherIp.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field OtherIp")
		unpacked, d := unpackZoneProtectionProfilesFloodOtherIpToSdk(ctx, model.OtherIp)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "OtherIp"})
		}
		if unpacked != nil {
			sdk.OtherIp = unpacked
		}
	}

	// Handling Objects
	if !model.SctpInit.IsNull() && !model.SctpInit.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field SctpInit")
		unpacked, d := unpackZoneProtectionProfilesFloodSctpInitToSdk(ctx, model.SctpInit)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "SctpInit"})
		}
		if unpacked != nil {
			sdk.SctpInit = unpacked
		}
	}

	// Handling Objects
	if !model.TcpSyn.IsNull() && !model.TcpSyn.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field TcpSyn")
		unpacked, d := unpackZoneProtectionProfilesFloodTcpSynToSdk(ctx, model.TcpSyn)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "TcpSyn"})
		}
		if unpacked != nil {
			sdk.TcpSyn = unpacked
		}
	}

	// Handling Objects
	if !model.Udp.IsNull() && !model.Udp.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Udp")
		unpacked, d := unpackZoneProtectionProfilesFloodUdpToSdk(ctx, model.Udp)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Udp"})
		}
		if unpacked != nil {
			sdk.Udp = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ZoneProtectionProfilesFlood", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ZoneProtectionProfilesFlood ---
func packZoneProtectionProfilesFloodFromSdk(ctx context.Context, sdk network_services.ZoneProtectionProfilesFlood) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ZoneProtectionProfilesFlood", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ZoneProtectionProfilesFlood
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Icmp != nil {
		tflog.Debug(ctx, "Packing nested object for field Icmp")
		packed, d := packZoneProtectionProfilesFloodIcmpFromSdk(ctx, *sdk.Icmp)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Icmp"})
		}
		model.Icmp = packed
	} else {
		model.Icmp = basetypes.NewObjectNull(models.ZoneProtectionProfilesFloodIcmp{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Icmpv6 != nil {
		tflog.Debug(ctx, "Packing nested object for field Icmpv6")
		packed, d := packZoneProtectionProfilesFloodIcmpv6FromSdk(ctx, *sdk.Icmpv6)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Icmpv6"})
		}
		model.Icmpv6 = packed
	} else {
		model.Icmpv6 = basetypes.NewObjectNull(models.ZoneProtectionProfilesFloodIcmpv6{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.OtherIp != nil {
		tflog.Debug(ctx, "Packing nested object for field OtherIp")
		packed, d := packZoneProtectionProfilesFloodOtherIpFromSdk(ctx, *sdk.OtherIp)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "OtherIp"})
		}
		model.OtherIp = packed
	} else {
		model.OtherIp = basetypes.NewObjectNull(models.ZoneProtectionProfilesFloodOtherIp{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.SctpInit != nil {
		tflog.Debug(ctx, "Packing nested object for field SctpInit")
		packed, d := packZoneProtectionProfilesFloodSctpInitFromSdk(ctx, *sdk.SctpInit)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "SctpInit"})
		}
		model.SctpInit = packed
	} else {
		model.SctpInit = basetypes.NewObjectNull(models.ZoneProtectionProfilesFloodSctpInit{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.TcpSyn != nil {
		tflog.Debug(ctx, "Packing nested object for field TcpSyn")
		packed, d := packZoneProtectionProfilesFloodTcpSynFromSdk(ctx, *sdk.TcpSyn)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "TcpSyn"})
		}
		model.TcpSyn = packed
	} else {
		model.TcpSyn = basetypes.NewObjectNull(models.ZoneProtectionProfilesFloodTcpSyn{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Udp != nil {
		tflog.Debug(ctx, "Packing nested object for field Udp")
		packed, d := packZoneProtectionProfilesFloodUdpFromSdk(ctx, *sdk.Udp)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Udp"})
		}
		model.Udp = packed
	} else {
		model.Udp = basetypes.NewObjectNull(models.ZoneProtectionProfilesFloodUdp{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ZoneProtectionProfilesFlood{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ZoneProtectionProfilesFlood", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ZoneProtectionProfilesFlood ---
func unpackZoneProtectionProfilesFloodListToSdk(ctx context.Context, list types.List) ([]network_services.ZoneProtectionProfilesFlood, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ZoneProtectionProfilesFlood")
	diags := diag.Diagnostics{}
	var data []models.ZoneProtectionProfilesFlood
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.ZoneProtectionProfilesFlood, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ZoneProtectionProfilesFlood{}.AttrTypes(), &item)
		unpacked, d := unpackZoneProtectionProfilesFloodToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ZoneProtectionProfilesFlood", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ZoneProtectionProfilesFlood ---
func packZoneProtectionProfilesFloodListFromSdk(ctx context.Context, sdks []network_services.ZoneProtectionProfilesFlood) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ZoneProtectionProfilesFlood")
	diags := diag.Diagnostics{}
	var data []models.ZoneProtectionProfilesFlood

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ZoneProtectionProfilesFlood
		obj, d := packZoneProtectionProfilesFloodFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ZoneProtectionProfilesFlood{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ZoneProtectionProfilesFlood", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ZoneProtectionProfilesFlood{}.AttrType(), data)
}

// --- Unpacker for ZoneProtectionProfilesFloodIcmp ---
func unpackZoneProtectionProfilesFloodIcmpToSdk(ctx context.Context, obj types.Object) (*network_services.ZoneProtectionProfilesFloodIcmp, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ZoneProtectionProfilesFloodIcmp", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ZoneProtectionProfilesFloodIcmp
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.ZoneProtectionProfilesFloodIcmp
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Enable.IsNull() && !model.Enable.IsUnknown() {
		sdk.Enable = model.Enable.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Enable", "value": *sdk.Enable})
	}

	// Handling Objects
	if !model.Red.IsNull() && !model.Red.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Red")
		unpacked, d := unpackZoneProtectionProfilesFloodIcmpRedToSdk(ctx, model.Red)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Red"})
		}
		if unpacked != nil {
			sdk.Red = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ZoneProtectionProfilesFloodIcmp", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ZoneProtectionProfilesFloodIcmp ---
func packZoneProtectionProfilesFloodIcmpFromSdk(ctx context.Context, sdk network_services.ZoneProtectionProfilesFloodIcmp) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ZoneProtectionProfilesFloodIcmp", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ZoneProtectionProfilesFloodIcmp
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Enable != nil {
		model.Enable = basetypes.NewBoolValue(*sdk.Enable)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Enable", "value": *sdk.Enable})
	} else {
		model.Enable = basetypes.NewBoolNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Red != nil {
		tflog.Debug(ctx, "Packing nested object for field Red")
		packed, d := packZoneProtectionProfilesFloodIcmpRedFromSdk(ctx, *sdk.Red)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Red"})
		}
		model.Red = packed
	} else {
		model.Red = basetypes.NewObjectNull(models.ZoneProtectionProfilesFloodIcmpRed{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ZoneProtectionProfilesFloodIcmp{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ZoneProtectionProfilesFloodIcmp", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ZoneProtectionProfilesFloodIcmp ---
func unpackZoneProtectionProfilesFloodIcmpListToSdk(ctx context.Context, list types.List) ([]network_services.ZoneProtectionProfilesFloodIcmp, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ZoneProtectionProfilesFloodIcmp")
	diags := diag.Diagnostics{}
	var data []models.ZoneProtectionProfilesFloodIcmp
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.ZoneProtectionProfilesFloodIcmp, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ZoneProtectionProfilesFloodIcmp{}.AttrTypes(), &item)
		unpacked, d := unpackZoneProtectionProfilesFloodIcmpToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ZoneProtectionProfilesFloodIcmp", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ZoneProtectionProfilesFloodIcmp ---
func packZoneProtectionProfilesFloodIcmpListFromSdk(ctx context.Context, sdks []network_services.ZoneProtectionProfilesFloodIcmp) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ZoneProtectionProfilesFloodIcmp")
	diags := diag.Diagnostics{}
	var data []models.ZoneProtectionProfilesFloodIcmp

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ZoneProtectionProfilesFloodIcmp
		obj, d := packZoneProtectionProfilesFloodIcmpFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ZoneProtectionProfilesFloodIcmp{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ZoneProtectionProfilesFloodIcmp", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ZoneProtectionProfilesFloodIcmp{}.AttrType(), data)
}

// --- Unpacker for ZoneProtectionProfilesFloodIcmpRed ---
func unpackZoneProtectionProfilesFloodIcmpRedToSdk(ctx context.Context, obj types.Object) (*network_services.ZoneProtectionProfilesFloodIcmpRed, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ZoneProtectionProfilesFloodIcmpRed", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ZoneProtectionProfilesFloodIcmpRed
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.ZoneProtectionProfilesFloodIcmpRed
	var d diag.Diagnostics
	// Handling Primitives
	if !model.ActivateRate.IsNull() && !model.ActivateRate.IsUnknown() {
		sdk.ActivateRate = int32(model.ActivateRate.ValueInt64())
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "ActivateRate", "value": sdk.ActivateRate})
	}

	// Handling Primitives
	if !model.AlarmRate.IsNull() && !model.AlarmRate.IsUnknown() {
		sdk.AlarmRate = int32(model.AlarmRate.ValueInt64())
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "AlarmRate", "value": sdk.AlarmRate})
	}

	// Handling Primitives
	if !model.MaximalRate.IsNull() && !model.MaximalRate.IsUnknown() {
		sdk.MaximalRate = int32(model.MaximalRate.ValueInt64())
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "MaximalRate", "value": sdk.MaximalRate})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ZoneProtectionProfilesFloodIcmpRed", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ZoneProtectionProfilesFloodIcmpRed ---
func packZoneProtectionProfilesFloodIcmpRedFromSdk(ctx context.Context, sdk network_services.ZoneProtectionProfilesFloodIcmpRed) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ZoneProtectionProfilesFloodIcmpRed", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ZoneProtectionProfilesFloodIcmpRed
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.ActivateRate = basetypes.NewInt64Value(int64(sdk.ActivateRate))
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "ActivateRate", "value": sdk.ActivateRate})
	// Handling Primitives
	// Standard primitive packing
	model.AlarmRate = basetypes.NewInt64Value(int64(sdk.AlarmRate))
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "AlarmRate", "value": sdk.AlarmRate})
	// Handling Primitives
	// Standard primitive packing
	model.MaximalRate = basetypes.NewInt64Value(int64(sdk.MaximalRate))
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "MaximalRate", "value": sdk.MaximalRate})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ZoneProtectionProfilesFloodIcmpRed{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ZoneProtectionProfilesFloodIcmpRed", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ZoneProtectionProfilesFloodIcmpRed ---
func unpackZoneProtectionProfilesFloodIcmpRedListToSdk(ctx context.Context, list types.List) ([]network_services.ZoneProtectionProfilesFloodIcmpRed, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ZoneProtectionProfilesFloodIcmpRed")
	diags := diag.Diagnostics{}
	var data []models.ZoneProtectionProfilesFloodIcmpRed
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.ZoneProtectionProfilesFloodIcmpRed, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ZoneProtectionProfilesFloodIcmpRed{}.AttrTypes(), &item)
		unpacked, d := unpackZoneProtectionProfilesFloodIcmpRedToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ZoneProtectionProfilesFloodIcmpRed", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ZoneProtectionProfilesFloodIcmpRed ---
func packZoneProtectionProfilesFloodIcmpRedListFromSdk(ctx context.Context, sdks []network_services.ZoneProtectionProfilesFloodIcmpRed) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ZoneProtectionProfilesFloodIcmpRed")
	diags := diag.Diagnostics{}
	var data []models.ZoneProtectionProfilesFloodIcmpRed

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ZoneProtectionProfilesFloodIcmpRed
		obj, d := packZoneProtectionProfilesFloodIcmpRedFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ZoneProtectionProfilesFloodIcmpRed{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ZoneProtectionProfilesFloodIcmpRed", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ZoneProtectionProfilesFloodIcmpRed{}.AttrType(), data)
}

// --- Unpacker for ZoneProtectionProfilesFloodIcmpv6 ---
func unpackZoneProtectionProfilesFloodIcmpv6ToSdk(ctx context.Context, obj types.Object) (*network_services.ZoneProtectionProfilesFloodIcmpv6, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ZoneProtectionProfilesFloodIcmpv6", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ZoneProtectionProfilesFloodIcmpv6
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.ZoneProtectionProfilesFloodIcmpv6
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Enable.IsNull() && !model.Enable.IsUnknown() {
		sdk.Enable = model.Enable.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Enable", "value": *sdk.Enable})
	}

	// Handling Objects
	if !model.Red.IsNull() && !model.Red.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Red")
		unpacked, d := unpackZoneProtectionProfilesFloodIcmpv6RedToSdk(ctx, model.Red)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Red"})
		}
		if unpacked != nil {
			sdk.Red = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ZoneProtectionProfilesFloodIcmpv6", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ZoneProtectionProfilesFloodIcmpv6 ---
func packZoneProtectionProfilesFloodIcmpv6FromSdk(ctx context.Context, sdk network_services.ZoneProtectionProfilesFloodIcmpv6) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ZoneProtectionProfilesFloodIcmpv6", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ZoneProtectionProfilesFloodIcmpv6
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Enable != nil {
		model.Enable = basetypes.NewBoolValue(*sdk.Enable)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Enable", "value": *sdk.Enable})
	} else {
		model.Enable = basetypes.NewBoolNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Red != nil {
		tflog.Debug(ctx, "Packing nested object for field Red")
		packed, d := packZoneProtectionProfilesFloodIcmpv6RedFromSdk(ctx, *sdk.Red)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Red"})
		}
		model.Red = packed
	} else {
		model.Red = basetypes.NewObjectNull(models.ZoneProtectionProfilesFloodIcmpv6Red{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ZoneProtectionProfilesFloodIcmpv6{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ZoneProtectionProfilesFloodIcmpv6", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ZoneProtectionProfilesFloodIcmpv6 ---
func unpackZoneProtectionProfilesFloodIcmpv6ListToSdk(ctx context.Context, list types.List) ([]network_services.ZoneProtectionProfilesFloodIcmpv6, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ZoneProtectionProfilesFloodIcmpv6")
	diags := diag.Diagnostics{}
	var data []models.ZoneProtectionProfilesFloodIcmpv6
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.ZoneProtectionProfilesFloodIcmpv6, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ZoneProtectionProfilesFloodIcmpv6{}.AttrTypes(), &item)
		unpacked, d := unpackZoneProtectionProfilesFloodIcmpv6ToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ZoneProtectionProfilesFloodIcmpv6", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ZoneProtectionProfilesFloodIcmpv6 ---
func packZoneProtectionProfilesFloodIcmpv6ListFromSdk(ctx context.Context, sdks []network_services.ZoneProtectionProfilesFloodIcmpv6) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ZoneProtectionProfilesFloodIcmpv6")
	diags := diag.Diagnostics{}
	var data []models.ZoneProtectionProfilesFloodIcmpv6

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ZoneProtectionProfilesFloodIcmpv6
		obj, d := packZoneProtectionProfilesFloodIcmpv6FromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ZoneProtectionProfilesFloodIcmpv6{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ZoneProtectionProfilesFloodIcmpv6", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ZoneProtectionProfilesFloodIcmpv6{}.AttrType(), data)
}

// --- Unpacker for ZoneProtectionProfilesFloodIcmpv6Red ---
func unpackZoneProtectionProfilesFloodIcmpv6RedToSdk(ctx context.Context, obj types.Object) (*network_services.ZoneProtectionProfilesFloodIcmpv6Red, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ZoneProtectionProfilesFloodIcmpv6Red", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ZoneProtectionProfilesFloodIcmpv6Red
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.ZoneProtectionProfilesFloodIcmpv6Red
	var d diag.Diagnostics
	// Handling Primitives
	if !model.ActivateRate.IsNull() && !model.ActivateRate.IsUnknown() {
		sdk.ActivateRate = int32(model.ActivateRate.ValueInt64())
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "ActivateRate", "value": sdk.ActivateRate})
	}

	// Handling Primitives
	if !model.AlarmRate.IsNull() && !model.AlarmRate.IsUnknown() {
		sdk.AlarmRate = int32(model.AlarmRate.ValueInt64())
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "AlarmRate", "value": sdk.AlarmRate})
	}

	// Handling Primitives
	if !model.MaximalRate.IsNull() && !model.MaximalRate.IsUnknown() {
		sdk.MaximalRate = int32(model.MaximalRate.ValueInt64())
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "MaximalRate", "value": sdk.MaximalRate})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ZoneProtectionProfilesFloodIcmpv6Red", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ZoneProtectionProfilesFloodIcmpv6Red ---
func packZoneProtectionProfilesFloodIcmpv6RedFromSdk(ctx context.Context, sdk network_services.ZoneProtectionProfilesFloodIcmpv6Red) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ZoneProtectionProfilesFloodIcmpv6Red", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ZoneProtectionProfilesFloodIcmpv6Red
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.ActivateRate = basetypes.NewInt64Value(int64(sdk.ActivateRate))
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "ActivateRate", "value": sdk.ActivateRate})
	// Handling Primitives
	// Standard primitive packing
	model.AlarmRate = basetypes.NewInt64Value(int64(sdk.AlarmRate))
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "AlarmRate", "value": sdk.AlarmRate})
	// Handling Primitives
	// Standard primitive packing
	model.MaximalRate = basetypes.NewInt64Value(int64(sdk.MaximalRate))
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "MaximalRate", "value": sdk.MaximalRate})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ZoneProtectionProfilesFloodIcmpv6Red{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ZoneProtectionProfilesFloodIcmpv6Red", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ZoneProtectionProfilesFloodIcmpv6Red ---
func unpackZoneProtectionProfilesFloodIcmpv6RedListToSdk(ctx context.Context, list types.List) ([]network_services.ZoneProtectionProfilesFloodIcmpv6Red, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ZoneProtectionProfilesFloodIcmpv6Red")
	diags := diag.Diagnostics{}
	var data []models.ZoneProtectionProfilesFloodIcmpv6Red
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.ZoneProtectionProfilesFloodIcmpv6Red, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ZoneProtectionProfilesFloodIcmpv6Red{}.AttrTypes(), &item)
		unpacked, d := unpackZoneProtectionProfilesFloodIcmpv6RedToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ZoneProtectionProfilesFloodIcmpv6Red", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ZoneProtectionProfilesFloodIcmpv6Red ---
func packZoneProtectionProfilesFloodIcmpv6RedListFromSdk(ctx context.Context, sdks []network_services.ZoneProtectionProfilesFloodIcmpv6Red) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ZoneProtectionProfilesFloodIcmpv6Red")
	diags := diag.Diagnostics{}
	var data []models.ZoneProtectionProfilesFloodIcmpv6Red

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ZoneProtectionProfilesFloodIcmpv6Red
		obj, d := packZoneProtectionProfilesFloodIcmpv6RedFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ZoneProtectionProfilesFloodIcmpv6Red{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ZoneProtectionProfilesFloodIcmpv6Red", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ZoneProtectionProfilesFloodIcmpv6Red{}.AttrType(), data)
}

// --- Unpacker for ZoneProtectionProfilesFloodOtherIp ---
func unpackZoneProtectionProfilesFloodOtherIpToSdk(ctx context.Context, obj types.Object) (*network_services.ZoneProtectionProfilesFloodOtherIp, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ZoneProtectionProfilesFloodOtherIp", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ZoneProtectionProfilesFloodOtherIp
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.ZoneProtectionProfilesFloodOtherIp
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Enable.IsNull() && !model.Enable.IsUnknown() {
		sdk.Enable = model.Enable.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Enable", "value": *sdk.Enable})
	}

	// Handling Objects
	if !model.Red.IsNull() && !model.Red.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Red")
		unpacked, d := unpackZoneProtectionProfilesFloodOtherIpRedToSdk(ctx, model.Red)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Red"})
		}
		if unpacked != nil {
			sdk.Red = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ZoneProtectionProfilesFloodOtherIp", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ZoneProtectionProfilesFloodOtherIp ---
func packZoneProtectionProfilesFloodOtherIpFromSdk(ctx context.Context, sdk network_services.ZoneProtectionProfilesFloodOtherIp) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ZoneProtectionProfilesFloodOtherIp", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ZoneProtectionProfilesFloodOtherIp
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Enable != nil {
		model.Enable = basetypes.NewBoolValue(*sdk.Enable)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Enable", "value": *sdk.Enable})
	} else {
		model.Enable = basetypes.NewBoolNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Red != nil {
		tflog.Debug(ctx, "Packing nested object for field Red")
		packed, d := packZoneProtectionProfilesFloodOtherIpRedFromSdk(ctx, *sdk.Red)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Red"})
		}
		model.Red = packed
	} else {
		model.Red = basetypes.NewObjectNull(models.ZoneProtectionProfilesFloodOtherIpRed{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ZoneProtectionProfilesFloodOtherIp{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ZoneProtectionProfilesFloodOtherIp", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ZoneProtectionProfilesFloodOtherIp ---
func unpackZoneProtectionProfilesFloodOtherIpListToSdk(ctx context.Context, list types.List) ([]network_services.ZoneProtectionProfilesFloodOtherIp, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ZoneProtectionProfilesFloodOtherIp")
	diags := diag.Diagnostics{}
	var data []models.ZoneProtectionProfilesFloodOtherIp
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.ZoneProtectionProfilesFloodOtherIp, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ZoneProtectionProfilesFloodOtherIp{}.AttrTypes(), &item)
		unpacked, d := unpackZoneProtectionProfilesFloodOtherIpToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ZoneProtectionProfilesFloodOtherIp", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ZoneProtectionProfilesFloodOtherIp ---
func packZoneProtectionProfilesFloodOtherIpListFromSdk(ctx context.Context, sdks []network_services.ZoneProtectionProfilesFloodOtherIp) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ZoneProtectionProfilesFloodOtherIp")
	diags := diag.Diagnostics{}
	var data []models.ZoneProtectionProfilesFloodOtherIp

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ZoneProtectionProfilesFloodOtherIp
		obj, d := packZoneProtectionProfilesFloodOtherIpFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ZoneProtectionProfilesFloodOtherIp{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ZoneProtectionProfilesFloodOtherIp", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ZoneProtectionProfilesFloodOtherIp{}.AttrType(), data)
}

// --- Unpacker for ZoneProtectionProfilesFloodOtherIpRed ---
func unpackZoneProtectionProfilesFloodOtherIpRedToSdk(ctx context.Context, obj types.Object) (*network_services.ZoneProtectionProfilesFloodOtherIpRed, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ZoneProtectionProfilesFloodOtherIpRed", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ZoneProtectionProfilesFloodOtherIpRed
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.ZoneProtectionProfilesFloodOtherIpRed
	var d diag.Diagnostics
	// Handling Primitives
	if !model.ActivateRate.IsNull() && !model.ActivateRate.IsUnknown() {
		sdk.ActivateRate = int32(model.ActivateRate.ValueInt64())
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "ActivateRate", "value": sdk.ActivateRate})
	}

	// Handling Primitives
	if !model.AlarmRate.IsNull() && !model.AlarmRate.IsUnknown() {
		sdk.AlarmRate = int32(model.AlarmRate.ValueInt64())
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "AlarmRate", "value": sdk.AlarmRate})
	}

	// Handling Primitives
	if !model.MaximalRate.IsNull() && !model.MaximalRate.IsUnknown() {
		sdk.MaximalRate = int32(model.MaximalRate.ValueInt64())
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "MaximalRate", "value": sdk.MaximalRate})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ZoneProtectionProfilesFloodOtherIpRed", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ZoneProtectionProfilesFloodOtherIpRed ---
func packZoneProtectionProfilesFloodOtherIpRedFromSdk(ctx context.Context, sdk network_services.ZoneProtectionProfilesFloodOtherIpRed) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ZoneProtectionProfilesFloodOtherIpRed", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ZoneProtectionProfilesFloodOtherIpRed
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.ActivateRate = basetypes.NewInt64Value(int64(sdk.ActivateRate))
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "ActivateRate", "value": sdk.ActivateRate})
	// Handling Primitives
	// Standard primitive packing
	model.AlarmRate = basetypes.NewInt64Value(int64(sdk.AlarmRate))
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "AlarmRate", "value": sdk.AlarmRate})
	// Handling Primitives
	// Standard primitive packing
	model.MaximalRate = basetypes.NewInt64Value(int64(sdk.MaximalRate))
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "MaximalRate", "value": sdk.MaximalRate})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ZoneProtectionProfilesFloodOtherIpRed{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ZoneProtectionProfilesFloodOtherIpRed", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ZoneProtectionProfilesFloodOtherIpRed ---
func unpackZoneProtectionProfilesFloodOtherIpRedListToSdk(ctx context.Context, list types.List) ([]network_services.ZoneProtectionProfilesFloodOtherIpRed, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ZoneProtectionProfilesFloodOtherIpRed")
	diags := diag.Diagnostics{}
	var data []models.ZoneProtectionProfilesFloodOtherIpRed
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.ZoneProtectionProfilesFloodOtherIpRed, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ZoneProtectionProfilesFloodOtherIpRed{}.AttrTypes(), &item)
		unpacked, d := unpackZoneProtectionProfilesFloodOtherIpRedToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ZoneProtectionProfilesFloodOtherIpRed", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ZoneProtectionProfilesFloodOtherIpRed ---
func packZoneProtectionProfilesFloodOtherIpRedListFromSdk(ctx context.Context, sdks []network_services.ZoneProtectionProfilesFloodOtherIpRed) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ZoneProtectionProfilesFloodOtherIpRed")
	diags := diag.Diagnostics{}
	var data []models.ZoneProtectionProfilesFloodOtherIpRed

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ZoneProtectionProfilesFloodOtherIpRed
		obj, d := packZoneProtectionProfilesFloodOtherIpRedFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ZoneProtectionProfilesFloodOtherIpRed{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ZoneProtectionProfilesFloodOtherIpRed", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ZoneProtectionProfilesFloodOtherIpRed{}.AttrType(), data)
}

// --- Unpacker for ZoneProtectionProfilesFloodSctpInit ---
func unpackZoneProtectionProfilesFloodSctpInitToSdk(ctx context.Context, obj types.Object) (*network_services.ZoneProtectionProfilesFloodSctpInit, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ZoneProtectionProfilesFloodSctpInit", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ZoneProtectionProfilesFloodSctpInit
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.ZoneProtectionProfilesFloodSctpInit
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Enable.IsNull() && !model.Enable.IsUnknown() {
		sdk.Enable = model.Enable.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Enable", "value": *sdk.Enable})
	}

	// Handling Objects
	if !model.Red.IsNull() && !model.Red.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Red")
		unpacked, d := unpackZoneProtectionProfilesFloodSctpInitRedToSdk(ctx, model.Red)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Red"})
		}
		if unpacked != nil {
			sdk.Red = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ZoneProtectionProfilesFloodSctpInit", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ZoneProtectionProfilesFloodSctpInit ---
func packZoneProtectionProfilesFloodSctpInitFromSdk(ctx context.Context, sdk network_services.ZoneProtectionProfilesFloodSctpInit) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ZoneProtectionProfilesFloodSctpInit", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ZoneProtectionProfilesFloodSctpInit
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Enable != nil {
		model.Enable = basetypes.NewBoolValue(*sdk.Enable)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Enable", "value": *sdk.Enable})
	} else {
		model.Enable = basetypes.NewBoolNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Red != nil {
		tflog.Debug(ctx, "Packing nested object for field Red")
		packed, d := packZoneProtectionProfilesFloodSctpInitRedFromSdk(ctx, *sdk.Red)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Red"})
		}
		model.Red = packed
	} else {
		model.Red = basetypes.NewObjectNull(models.ZoneProtectionProfilesFloodSctpInitRed{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ZoneProtectionProfilesFloodSctpInit{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ZoneProtectionProfilesFloodSctpInit", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ZoneProtectionProfilesFloodSctpInit ---
func unpackZoneProtectionProfilesFloodSctpInitListToSdk(ctx context.Context, list types.List) ([]network_services.ZoneProtectionProfilesFloodSctpInit, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ZoneProtectionProfilesFloodSctpInit")
	diags := diag.Diagnostics{}
	var data []models.ZoneProtectionProfilesFloodSctpInit
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.ZoneProtectionProfilesFloodSctpInit, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ZoneProtectionProfilesFloodSctpInit{}.AttrTypes(), &item)
		unpacked, d := unpackZoneProtectionProfilesFloodSctpInitToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ZoneProtectionProfilesFloodSctpInit", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ZoneProtectionProfilesFloodSctpInit ---
func packZoneProtectionProfilesFloodSctpInitListFromSdk(ctx context.Context, sdks []network_services.ZoneProtectionProfilesFloodSctpInit) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ZoneProtectionProfilesFloodSctpInit")
	diags := diag.Diagnostics{}
	var data []models.ZoneProtectionProfilesFloodSctpInit

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ZoneProtectionProfilesFloodSctpInit
		obj, d := packZoneProtectionProfilesFloodSctpInitFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ZoneProtectionProfilesFloodSctpInit{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ZoneProtectionProfilesFloodSctpInit", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ZoneProtectionProfilesFloodSctpInit{}.AttrType(), data)
}

// --- Unpacker for ZoneProtectionProfilesFloodSctpInitRed ---
func unpackZoneProtectionProfilesFloodSctpInitRedToSdk(ctx context.Context, obj types.Object) (*network_services.ZoneProtectionProfilesFloodSctpInitRed, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ZoneProtectionProfilesFloodSctpInitRed", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ZoneProtectionProfilesFloodSctpInitRed
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.ZoneProtectionProfilesFloodSctpInitRed
	var d diag.Diagnostics
	// Handling Primitives
	if !model.ActivateRate.IsNull() && !model.ActivateRate.IsUnknown() {
		sdk.ActivateRate = int32(model.ActivateRate.ValueInt64())
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "ActivateRate", "value": sdk.ActivateRate})
	}

	// Handling Primitives
	if !model.AlarmRate.IsNull() && !model.AlarmRate.IsUnknown() {
		sdk.AlarmRate = int32(model.AlarmRate.ValueInt64())
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "AlarmRate", "value": sdk.AlarmRate})
	}

	// Handling Primitives
	if !model.MaximalRate.IsNull() && !model.MaximalRate.IsUnknown() {
		sdk.MaximalRate = int32(model.MaximalRate.ValueInt64())
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "MaximalRate", "value": sdk.MaximalRate})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ZoneProtectionProfilesFloodSctpInitRed", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ZoneProtectionProfilesFloodSctpInitRed ---
func packZoneProtectionProfilesFloodSctpInitRedFromSdk(ctx context.Context, sdk network_services.ZoneProtectionProfilesFloodSctpInitRed) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ZoneProtectionProfilesFloodSctpInitRed", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ZoneProtectionProfilesFloodSctpInitRed
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.ActivateRate = basetypes.NewInt64Value(int64(sdk.ActivateRate))
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "ActivateRate", "value": sdk.ActivateRate})
	// Handling Primitives
	// Standard primitive packing
	model.AlarmRate = basetypes.NewInt64Value(int64(sdk.AlarmRate))
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "AlarmRate", "value": sdk.AlarmRate})
	// Handling Primitives
	// Standard primitive packing
	model.MaximalRate = basetypes.NewInt64Value(int64(sdk.MaximalRate))
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "MaximalRate", "value": sdk.MaximalRate})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ZoneProtectionProfilesFloodSctpInitRed{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ZoneProtectionProfilesFloodSctpInitRed", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ZoneProtectionProfilesFloodSctpInitRed ---
func unpackZoneProtectionProfilesFloodSctpInitRedListToSdk(ctx context.Context, list types.List) ([]network_services.ZoneProtectionProfilesFloodSctpInitRed, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ZoneProtectionProfilesFloodSctpInitRed")
	diags := diag.Diagnostics{}
	var data []models.ZoneProtectionProfilesFloodSctpInitRed
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.ZoneProtectionProfilesFloodSctpInitRed, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ZoneProtectionProfilesFloodSctpInitRed{}.AttrTypes(), &item)
		unpacked, d := unpackZoneProtectionProfilesFloodSctpInitRedToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ZoneProtectionProfilesFloodSctpInitRed", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ZoneProtectionProfilesFloodSctpInitRed ---
func packZoneProtectionProfilesFloodSctpInitRedListFromSdk(ctx context.Context, sdks []network_services.ZoneProtectionProfilesFloodSctpInitRed) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ZoneProtectionProfilesFloodSctpInitRed")
	diags := diag.Diagnostics{}
	var data []models.ZoneProtectionProfilesFloodSctpInitRed

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ZoneProtectionProfilesFloodSctpInitRed
		obj, d := packZoneProtectionProfilesFloodSctpInitRedFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ZoneProtectionProfilesFloodSctpInitRed{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ZoneProtectionProfilesFloodSctpInitRed", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ZoneProtectionProfilesFloodSctpInitRed{}.AttrType(), data)
}

// --- Unpacker for ZoneProtectionProfilesFloodTcpSyn ---
func unpackZoneProtectionProfilesFloodTcpSynToSdk(ctx context.Context, obj types.Object) (*network_services.ZoneProtectionProfilesFloodTcpSyn, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ZoneProtectionProfilesFloodTcpSyn", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ZoneProtectionProfilesFloodTcpSyn
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.ZoneProtectionProfilesFloodTcpSyn
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Enable.IsNull() && !model.Enable.IsUnknown() {
		sdk.Enable = model.Enable.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Enable", "value": *sdk.Enable})
	}

	// Handling Objects
	if !model.Red.IsNull() && !model.Red.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Red")
		unpacked, d := unpackZoneProtectionProfilesFloodTcpSynRedToSdk(ctx, model.Red)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Red"})
		}
		if unpacked != nil {
			sdk.Red = unpacked
		}
	}

	// Handling Objects
	if !model.SynCookies.IsNull() && !model.SynCookies.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field SynCookies")
		unpacked, d := unpackZoneProtectionProfilesFloodTcpSynSynCookiesToSdk(ctx, model.SynCookies)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "SynCookies"})
		}
		if unpacked != nil {
			sdk.SynCookies = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ZoneProtectionProfilesFloodTcpSyn", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ZoneProtectionProfilesFloodTcpSyn ---
func packZoneProtectionProfilesFloodTcpSynFromSdk(ctx context.Context, sdk network_services.ZoneProtectionProfilesFloodTcpSyn) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ZoneProtectionProfilesFloodTcpSyn", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ZoneProtectionProfilesFloodTcpSyn
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Enable != nil {
		model.Enable = basetypes.NewBoolValue(*sdk.Enable)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Enable", "value": *sdk.Enable})
	} else {
		model.Enable = basetypes.NewBoolNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Red != nil {
		tflog.Debug(ctx, "Packing nested object for field Red")
		packed, d := packZoneProtectionProfilesFloodTcpSynRedFromSdk(ctx, *sdk.Red)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Red"})
		}
		model.Red = packed
	} else {
		model.Red = basetypes.NewObjectNull(models.ZoneProtectionProfilesFloodTcpSynRed{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.SynCookies != nil {
		tflog.Debug(ctx, "Packing nested object for field SynCookies")
		packed, d := packZoneProtectionProfilesFloodTcpSynSynCookiesFromSdk(ctx, *sdk.SynCookies)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "SynCookies"})
		}
		model.SynCookies = packed
	} else {
		model.SynCookies = basetypes.NewObjectNull(models.ZoneProtectionProfilesFloodTcpSynSynCookies{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ZoneProtectionProfilesFloodTcpSyn{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ZoneProtectionProfilesFloodTcpSyn", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ZoneProtectionProfilesFloodTcpSyn ---
func unpackZoneProtectionProfilesFloodTcpSynListToSdk(ctx context.Context, list types.List) ([]network_services.ZoneProtectionProfilesFloodTcpSyn, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ZoneProtectionProfilesFloodTcpSyn")
	diags := diag.Diagnostics{}
	var data []models.ZoneProtectionProfilesFloodTcpSyn
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.ZoneProtectionProfilesFloodTcpSyn, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ZoneProtectionProfilesFloodTcpSyn{}.AttrTypes(), &item)
		unpacked, d := unpackZoneProtectionProfilesFloodTcpSynToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ZoneProtectionProfilesFloodTcpSyn", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ZoneProtectionProfilesFloodTcpSyn ---
func packZoneProtectionProfilesFloodTcpSynListFromSdk(ctx context.Context, sdks []network_services.ZoneProtectionProfilesFloodTcpSyn) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ZoneProtectionProfilesFloodTcpSyn")
	diags := diag.Diagnostics{}
	var data []models.ZoneProtectionProfilesFloodTcpSyn

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ZoneProtectionProfilesFloodTcpSyn
		obj, d := packZoneProtectionProfilesFloodTcpSynFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ZoneProtectionProfilesFloodTcpSyn{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ZoneProtectionProfilesFloodTcpSyn", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ZoneProtectionProfilesFloodTcpSyn{}.AttrType(), data)
}

// --- Unpacker for ZoneProtectionProfilesFloodTcpSynRed ---
func unpackZoneProtectionProfilesFloodTcpSynRedToSdk(ctx context.Context, obj types.Object) (*network_services.ZoneProtectionProfilesFloodTcpSynRed, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ZoneProtectionProfilesFloodTcpSynRed", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ZoneProtectionProfilesFloodTcpSynRed
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.ZoneProtectionProfilesFloodTcpSynRed
	var d diag.Diagnostics
	// Handling Primitives
	if !model.ActivateRate.IsNull() && !model.ActivateRate.IsUnknown() {
		sdk.ActivateRate = int32(model.ActivateRate.ValueInt64())
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "ActivateRate", "value": sdk.ActivateRate})
	}

	// Handling Primitives
	if !model.AlarmRate.IsNull() && !model.AlarmRate.IsUnknown() {
		sdk.AlarmRate = int32(model.AlarmRate.ValueInt64())
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "AlarmRate", "value": sdk.AlarmRate})
	}

	// Handling Primitives
	if !model.MaximalRate.IsNull() && !model.MaximalRate.IsUnknown() {
		sdk.MaximalRate = int32(model.MaximalRate.ValueInt64())
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "MaximalRate", "value": sdk.MaximalRate})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ZoneProtectionProfilesFloodTcpSynRed", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ZoneProtectionProfilesFloodTcpSynRed ---
func packZoneProtectionProfilesFloodTcpSynRedFromSdk(ctx context.Context, sdk network_services.ZoneProtectionProfilesFloodTcpSynRed) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ZoneProtectionProfilesFloodTcpSynRed", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ZoneProtectionProfilesFloodTcpSynRed
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.ActivateRate = basetypes.NewInt64Value(int64(sdk.ActivateRate))
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "ActivateRate", "value": sdk.ActivateRate})
	// Handling Primitives
	// Standard primitive packing
	model.AlarmRate = basetypes.NewInt64Value(int64(sdk.AlarmRate))
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "AlarmRate", "value": sdk.AlarmRate})
	// Handling Primitives
	// Standard primitive packing
	model.MaximalRate = basetypes.NewInt64Value(int64(sdk.MaximalRate))
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "MaximalRate", "value": sdk.MaximalRate})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ZoneProtectionProfilesFloodTcpSynRed{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ZoneProtectionProfilesFloodTcpSynRed", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ZoneProtectionProfilesFloodTcpSynRed ---
func unpackZoneProtectionProfilesFloodTcpSynRedListToSdk(ctx context.Context, list types.List) ([]network_services.ZoneProtectionProfilesFloodTcpSynRed, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ZoneProtectionProfilesFloodTcpSynRed")
	diags := diag.Diagnostics{}
	var data []models.ZoneProtectionProfilesFloodTcpSynRed
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.ZoneProtectionProfilesFloodTcpSynRed, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ZoneProtectionProfilesFloodTcpSynRed{}.AttrTypes(), &item)
		unpacked, d := unpackZoneProtectionProfilesFloodTcpSynRedToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ZoneProtectionProfilesFloodTcpSynRed", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ZoneProtectionProfilesFloodTcpSynRed ---
func packZoneProtectionProfilesFloodTcpSynRedListFromSdk(ctx context.Context, sdks []network_services.ZoneProtectionProfilesFloodTcpSynRed) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ZoneProtectionProfilesFloodTcpSynRed")
	diags := diag.Diagnostics{}
	var data []models.ZoneProtectionProfilesFloodTcpSynRed

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ZoneProtectionProfilesFloodTcpSynRed
		obj, d := packZoneProtectionProfilesFloodTcpSynRedFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ZoneProtectionProfilesFloodTcpSynRed{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ZoneProtectionProfilesFloodTcpSynRed", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ZoneProtectionProfilesFloodTcpSynRed{}.AttrType(), data)
}

// --- Unpacker for ZoneProtectionProfilesFloodTcpSynSynCookies ---
func unpackZoneProtectionProfilesFloodTcpSynSynCookiesToSdk(ctx context.Context, obj types.Object) (*network_services.ZoneProtectionProfilesFloodTcpSynSynCookies, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ZoneProtectionProfilesFloodTcpSynSynCookies", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ZoneProtectionProfilesFloodTcpSynSynCookies
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.ZoneProtectionProfilesFloodTcpSynSynCookies
	var d diag.Diagnostics
	// Handling Primitives
	if !model.ActivateRate.IsNull() && !model.ActivateRate.IsUnknown() {
		sdk.ActivateRate = int32(model.ActivateRate.ValueInt64())
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "ActivateRate", "value": sdk.ActivateRate})
	}

	// Handling Primitives
	if !model.AlarmRate.IsNull() && !model.AlarmRate.IsUnknown() {
		sdk.AlarmRate = int32(model.AlarmRate.ValueInt64())
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "AlarmRate", "value": sdk.AlarmRate})
	}

	// Handling Primitives
	if !model.MaximalRate.IsNull() && !model.MaximalRate.IsUnknown() {
		sdk.MaximalRate = int32(model.MaximalRate.ValueInt64())
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "MaximalRate", "value": sdk.MaximalRate})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ZoneProtectionProfilesFloodTcpSynSynCookies", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ZoneProtectionProfilesFloodTcpSynSynCookies ---
func packZoneProtectionProfilesFloodTcpSynSynCookiesFromSdk(ctx context.Context, sdk network_services.ZoneProtectionProfilesFloodTcpSynSynCookies) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ZoneProtectionProfilesFloodTcpSynSynCookies", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ZoneProtectionProfilesFloodTcpSynSynCookies
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.ActivateRate = basetypes.NewInt64Value(int64(sdk.ActivateRate))
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "ActivateRate", "value": sdk.ActivateRate})
	// Handling Primitives
	// Standard primitive packing
	model.AlarmRate = basetypes.NewInt64Value(int64(sdk.AlarmRate))
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "AlarmRate", "value": sdk.AlarmRate})
	// Handling Primitives
	// Standard primitive packing
	model.MaximalRate = basetypes.NewInt64Value(int64(sdk.MaximalRate))
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "MaximalRate", "value": sdk.MaximalRate})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ZoneProtectionProfilesFloodTcpSynSynCookies{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ZoneProtectionProfilesFloodTcpSynSynCookies", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ZoneProtectionProfilesFloodTcpSynSynCookies ---
func unpackZoneProtectionProfilesFloodTcpSynSynCookiesListToSdk(ctx context.Context, list types.List) ([]network_services.ZoneProtectionProfilesFloodTcpSynSynCookies, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ZoneProtectionProfilesFloodTcpSynSynCookies")
	diags := diag.Diagnostics{}
	var data []models.ZoneProtectionProfilesFloodTcpSynSynCookies
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.ZoneProtectionProfilesFloodTcpSynSynCookies, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ZoneProtectionProfilesFloodTcpSynSynCookies{}.AttrTypes(), &item)
		unpacked, d := unpackZoneProtectionProfilesFloodTcpSynSynCookiesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ZoneProtectionProfilesFloodTcpSynSynCookies", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ZoneProtectionProfilesFloodTcpSynSynCookies ---
func packZoneProtectionProfilesFloodTcpSynSynCookiesListFromSdk(ctx context.Context, sdks []network_services.ZoneProtectionProfilesFloodTcpSynSynCookies) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ZoneProtectionProfilesFloodTcpSynSynCookies")
	diags := diag.Diagnostics{}
	var data []models.ZoneProtectionProfilesFloodTcpSynSynCookies

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ZoneProtectionProfilesFloodTcpSynSynCookies
		obj, d := packZoneProtectionProfilesFloodTcpSynSynCookiesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ZoneProtectionProfilesFloodTcpSynSynCookies{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ZoneProtectionProfilesFloodTcpSynSynCookies", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ZoneProtectionProfilesFloodTcpSynSynCookies{}.AttrType(), data)
}

// --- Unpacker for ZoneProtectionProfilesFloodUdp ---
func unpackZoneProtectionProfilesFloodUdpToSdk(ctx context.Context, obj types.Object) (*network_services.ZoneProtectionProfilesFloodUdp, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ZoneProtectionProfilesFloodUdp", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ZoneProtectionProfilesFloodUdp
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.ZoneProtectionProfilesFloodUdp
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Enable.IsNull() && !model.Enable.IsUnknown() {
		sdk.Enable = model.Enable.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Enable", "value": *sdk.Enable})
	}

	// Handling Objects
	if !model.Red.IsNull() && !model.Red.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Red")
		unpacked, d := unpackZoneProtectionProfilesFloodUdpRedToSdk(ctx, model.Red)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Red"})
		}
		if unpacked != nil {
			sdk.Red = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ZoneProtectionProfilesFloodUdp", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ZoneProtectionProfilesFloodUdp ---
func packZoneProtectionProfilesFloodUdpFromSdk(ctx context.Context, sdk network_services.ZoneProtectionProfilesFloodUdp) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ZoneProtectionProfilesFloodUdp", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ZoneProtectionProfilesFloodUdp
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Enable != nil {
		model.Enable = basetypes.NewBoolValue(*sdk.Enable)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Enable", "value": *sdk.Enable})
	} else {
		model.Enable = basetypes.NewBoolNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Red != nil {
		tflog.Debug(ctx, "Packing nested object for field Red")
		packed, d := packZoneProtectionProfilesFloodUdpRedFromSdk(ctx, *sdk.Red)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Red"})
		}
		model.Red = packed
	} else {
		model.Red = basetypes.NewObjectNull(models.ZoneProtectionProfilesFloodUdpRed{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ZoneProtectionProfilesFloodUdp{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ZoneProtectionProfilesFloodUdp", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ZoneProtectionProfilesFloodUdp ---
func unpackZoneProtectionProfilesFloodUdpListToSdk(ctx context.Context, list types.List) ([]network_services.ZoneProtectionProfilesFloodUdp, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ZoneProtectionProfilesFloodUdp")
	diags := diag.Diagnostics{}
	var data []models.ZoneProtectionProfilesFloodUdp
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.ZoneProtectionProfilesFloodUdp, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ZoneProtectionProfilesFloodUdp{}.AttrTypes(), &item)
		unpacked, d := unpackZoneProtectionProfilesFloodUdpToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ZoneProtectionProfilesFloodUdp", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ZoneProtectionProfilesFloodUdp ---
func packZoneProtectionProfilesFloodUdpListFromSdk(ctx context.Context, sdks []network_services.ZoneProtectionProfilesFloodUdp) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ZoneProtectionProfilesFloodUdp")
	diags := diag.Diagnostics{}
	var data []models.ZoneProtectionProfilesFloodUdp

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ZoneProtectionProfilesFloodUdp
		obj, d := packZoneProtectionProfilesFloodUdpFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ZoneProtectionProfilesFloodUdp{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ZoneProtectionProfilesFloodUdp", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ZoneProtectionProfilesFloodUdp{}.AttrType(), data)
}

// --- Unpacker for ZoneProtectionProfilesFloodUdpRed ---
func unpackZoneProtectionProfilesFloodUdpRedToSdk(ctx context.Context, obj types.Object) (*network_services.ZoneProtectionProfilesFloodUdpRed, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ZoneProtectionProfilesFloodUdpRed", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ZoneProtectionProfilesFloodUdpRed
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.ZoneProtectionProfilesFloodUdpRed
	var d diag.Diagnostics
	// Handling Primitives
	if !model.ActivateRate.IsNull() && !model.ActivateRate.IsUnknown() {
		sdk.ActivateRate = int32(model.ActivateRate.ValueInt64())
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "ActivateRate", "value": sdk.ActivateRate})
	}

	// Handling Primitives
	if !model.AlarmRate.IsNull() && !model.AlarmRate.IsUnknown() {
		sdk.AlarmRate = int32(model.AlarmRate.ValueInt64())
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "AlarmRate", "value": sdk.AlarmRate})
	}

	// Handling Primitives
	if !model.MaximalRate.IsNull() && !model.MaximalRate.IsUnknown() {
		sdk.MaximalRate = int32(model.MaximalRate.ValueInt64())
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "MaximalRate", "value": sdk.MaximalRate})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ZoneProtectionProfilesFloodUdpRed", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ZoneProtectionProfilesFloodUdpRed ---
func packZoneProtectionProfilesFloodUdpRedFromSdk(ctx context.Context, sdk network_services.ZoneProtectionProfilesFloodUdpRed) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ZoneProtectionProfilesFloodUdpRed", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ZoneProtectionProfilesFloodUdpRed
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.ActivateRate = basetypes.NewInt64Value(int64(sdk.ActivateRate))
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "ActivateRate", "value": sdk.ActivateRate})
	// Handling Primitives
	// Standard primitive packing
	model.AlarmRate = basetypes.NewInt64Value(int64(sdk.AlarmRate))
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "AlarmRate", "value": sdk.AlarmRate})
	// Handling Primitives
	// Standard primitive packing
	model.MaximalRate = basetypes.NewInt64Value(int64(sdk.MaximalRate))
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "MaximalRate", "value": sdk.MaximalRate})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ZoneProtectionProfilesFloodUdpRed{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ZoneProtectionProfilesFloodUdpRed", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ZoneProtectionProfilesFloodUdpRed ---
func unpackZoneProtectionProfilesFloodUdpRedListToSdk(ctx context.Context, list types.List) ([]network_services.ZoneProtectionProfilesFloodUdpRed, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ZoneProtectionProfilesFloodUdpRed")
	diags := diag.Diagnostics{}
	var data []models.ZoneProtectionProfilesFloodUdpRed
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.ZoneProtectionProfilesFloodUdpRed, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ZoneProtectionProfilesFloodUdpRed{}.AttrTypes(), &item)
		unpacked, d := unpackZoneProtectionProfilesFloodUdpRedToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ZoneProtectionProfilesFloodUdpRed", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ZoneProtectionProfilesFloodUdpRed ---
func packZoneProtectionProfilesFloodUdpRedListFromSdk(ctx context.Context, sdks []network_services.ZoneProtectionProfilesFloodUdpRed) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ZoneProtectionProfilesFloodUdpRed")
	diags := diag.Diagnostics{}
	var data []models.ZoneProtectionProfilesFloodUdpRed

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ZoneProtectionProfilesFloodUdpRed
		obj, d := packZoneProtectionProfilesFloodUdpRedFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ZoneProtectionProfilesFloodUdpRed{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ZoneProtectionProfilesFloodUdpRed", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ZoneProtectionProfilesFloodUdpRed{}.AttrType(), data)
}

// --- Unpacker for ZoneProtectionProfilesIpv6 ---
func unpackZoneProtectionProfilesIpv6ToSdk(ctx context.Context, obj types.Object) (*network_services.ZoneProtectionProfilesIpv6, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ZoneProtectionProfilesIpv6", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ZoneProtectionProfilesIpv6
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.ZoneProtectionProfilesIpv6
	var d diag.Diagnostics
	// Handling Primitives
	if !model.AnycastSource.IsNull() && !model.AnycastSource.IsUnknown() {
		sdk.AnycastSource = model.AnycastSource.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "AnycastSource", "value": *sdk.AnycastSource})
	}

	// Handling Objects
	if !model.FilterExtHdr.IsNull() && !model.FilterExtHdr.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field FilterExtHdr")
		unpacked, d := unpackZoneProtectionProfilesIpv6FilterExtHdrToSdk(ctx, model.FilterExtHdr)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "FilterExtHdr"})
		}
		if unpacked != nil {
			sdk.FilterExtHdr = unpacked
		}
	}

	// Handling Primitives
	if !model.Icmpv6TooBigSmallMtuDiscard.IsNull() && !model.Icmpv6TooBigSmallMtuDiscard.IsUnknown() {
		sdk.Icmpv6TooBigSmallMtuDiscard = model.Icmpv6TooBigSmallMtuDiscard.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Icmpv6TooBigSmallMtuDiscard", "value": *sdk.Icmpv6TooBigSmallMtuDiscard})
	}

	// Handling Objects
	if !model.IgnoreInvPkt.IsNull() && !model.IgnoreInvPkt.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field IgnoreInvPkt")
		unpacked, d := unpackZoneProtectionProfilesIpv6IgnoreInvPktToSdk(ctx, model.IgnoreInvPkt)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "IgnoreInvPkt"})
		}
		if unpacked != nil {
			sdk.IgnoreInvPkt = unpacked
		}
	}

	// Handling Primitives
	if !model.Ipv4CompatibleAddress.IsNull() && !model.Ipv4CompatibleAddress.IsUnknown() {
		sdk.Ipv4CompatibleAddress = model.Ipv4CompatibleAddress.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Ipv4CompatibleAddress", "value": *sdk.Ipv4CompatibleAddress})
	}

	// Handling Primitives
	if !model.NeedlessFragmentHdr.IsNull() && !model.NeedlessFragmentHdr.IsUnknown() {
		sdk.NeedlessFragmentHdr = model.NeedlessFragmentHdr.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "NeedlessFragmentHdr", "value": *sdk.NeedlessFragmentHdr})
	}

	// Handling Primitives
	if !model.OptionsInvalidIpv6Discard.IsNull() && !model.OptionsInvalidIpv6Discard.IsUnknown() {
		sdk.OptionsInvalidIpv6Discard = model.OptionsInvalidIpv6Discard.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "OptionsInvalidIpv6Discard", "value": *sdk.OptionsInvalidIpv6Discard})
	}

	// Handling Primitives
	if !model.ReservedFieldSetDiscard.IsNull() && !model.ReservedFieldSetDiscard.IsUnknown() {
		sdk.ReservedFieldSetDiscard = model.ReservedFieldSetDiscard.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "ReservedFieldSetDiscard", "value": *sdk.ReservedFieldSetDiscard})
	}

	// Handling Primitives
	if !model.RoutingHeader0.IsNull() && !model.RoutingHeader0.IsUnknown() {
		sdk.RoutingHeader0 = model.RoutingHeader0.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "RoutingHeader0", "value": *sdk.RoutingHeader0})
	}

	// Handling Primitives
	if !model.RoutingHeader1.IsNull() && !model.RoutingHeader1.IsUnknown() {
		sdk.RoutingHeader1 = model.RoutingHeader1.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "RoutingHeader1", "value": *sdk.RoutingHeader1})
	}

	// Handling Primitives
	if !model.RoutingHeader253.IsNull() && !model.RoutingHeader253.IsUnknown() {
		sdk.RoutingHeader253 = model.RoutingHeader253.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "RoutingHeader253", "value": *sdk.RoutingHeader253})
	}

	// Handling Primitives
	if !model.RoutingHeader254.IsNull() && !model.RoutingHeader254.IsUnknown() {
		sdk.RoutingHeader254 = model.RoutingHeader254.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "RoutingHeader254", "value": *sdk.RoutingHeader254})
	}

	// Handling Primitives
	if !model.RoutingHeader255.IsNull() && !model.RoutingHeader255.IsUnknown() {
		sdk.RoutingHeader255 = model.RoutingHeader255.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "RoutingHeader255", "value": *sdk.RoutingHeader255})
	}

	// Handling Primitives
	if !model.RoutingHeader3.IsNull() && !model.RoutingHeader3.IsUnknown() {
		sdk.RoutingHeader3 = model.RoutingHeader3.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "RoutingHeader3", "value": *sdk.RoutingHeader3})
	}

	// Handling Primitives
	if !model.RoutingHeader4252.IsNull() && !model.RoutingHeader4252.IsUnknown() {
		sdk.RoutingHeader4252 = model.RoutingHeader4252.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "RoutingHeader4252", "value": *sdk.RoutingHeader4252})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ZoneProtectionProfilesIpv6", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ZoneProtectionProfilesIpv6 ---
func packZoneProtectionProfilesIpv6FromSdk(ctx context.Context, sdk network_services.ZoneProtectionProfilesIpv6) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ZoneProtectionProfilesIpv6", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ZoneProtectionProfilesIpv6
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.AnycastSource != nil {
		model.AnycastSource = basetypes.NewBoolValue(*sdk.AnycastSource)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "AnycastSource", "value": *sdk.AnycastSource})
	} else {
		model.AnycastSource = basetypes.NewBoolNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.FilterExtHdr != nil {
		tflog.Debug(ctx, "Packing nested object for field FilterExtHdr")
		packed, d := packZoneProtectionProfilesIpv6FilterExtHdrFromSdk(ctx, *sdk.FilterExtHdr)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "FilterExtHdr"})
		}
		model.FilterExtHdr = packed
	} else {
		model.FilterExtHdr = basetypes.NewObjectNull(models.ZoneProtectionProfilesIpv6FilterExtHdr{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Icmpv6TooBigSmallMtuDiscard != nil {
		model.Icmpv6TooBigSmallMtuDiscard = basetypes.NewBoolValue(*sdk.Icmpv6TooBigSmallMtuDiscard)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Icmpv6TooBigSmallMtuDiscard", "value": *sdk.Icmpv6TooBigSmallMtuDiscard})
	} else {
		model.Icmpv6TooBigSmallMtuDiscard = basetypes.NewBoolNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.IgnoreInvPkt != nil {
		tflog.Debug(ctx, "Packing nested object for field IgnoreInvPkt")
		packed, d := packZoneProtectionProfilesIpv6IgnoreInvPktFromSdk(ctx, *sdk.IgnoreInvPkt)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "IgnoreInvPkt"})
		}
		model.IgnoreInvPkt = packed
	} else {
		model.IgnoreInvPkt = basetypes.NewObjectNull(models.ZoneProtectionProfilesIpv6IgnoreInvPkt{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Ipv4CompatibleAddress != nil {
		model.Ipv4CompatibleAddress = basetypes.NewBoolValue(*sdk.Ipv4CompatibleAddress)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Ipv4CompatibleAddress", "value": *sdk.Ipv4CompatibleAddress})
	} else {
		model.Ipv4CompatibleAddress = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.NeedlessFragmentHdr != nil {
		model.NeedlessFragmentHdr = basetypes.NewBoolValue(*sdk.NeedlessFragmentHdr)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "NeedlessFragmentHdr", "value": *sdk.NeedlessFragmentHdr})
	} else {
		model.NeedlessFragmentHdr = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.OptionsInvalidIpv6Discard != nil {
		model.OptionsInvalidIpv6Discard = basetypes.NewBoolValue(*sdk.OptionsInvalidIpv6Discard)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "OptionsInvalidIpv6Discard", "value": *sdk.OptionsInvalidIpv6Discard})
	} else {
		model.OptionsInvalidIpv6Discard = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.ReservedFieldSetDiscard != nil {
		model.ReservedFieldSetDiscard = basetypes.NewBoolValue(*sdk.ReservedFieldSetDiscard)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "ReservedFieldSetDiscard", "value": *sdk.ReservedFieldSetDiscard})
	} else {
		model.ReservedFieldSetDiscard = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.RoutingHeader0 != nil {
		model.RoutingHeader0 = basetypes.NewBoolValue(*sdk.RoutingHeader0)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "RoutingHeader0", "value": *sdk.RoutingHeader0})
	} else {
		model.RoutingHeader0 = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.RoutingHeader1 != nil {
		model.RoutingHeader1 = basetypes.NewBoolValue(*sdk.RoutingHeader1)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "RoutingHeader1", "value": *sdk.RoutingHeader1})
	} else {
		model.RoutingHeader1 = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.RoutingHeader253 != nil {
		model.RoutingHeader253 = basetypes.NewBoolValue(*sdk.RoutingHeader253)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "RoutingHeader253", "value": *sdk.RoutingHeader253})
	} else {
		model.RoutingHeader253 = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.RoutingHeader254 != nil {
		model.RoutingHeader254 = basetypes.NewBoolValue(*sdk.RoutingHeader254)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "RoutingHeader254", "value": *sdk.RoutingHeader254})
	} else {
		model.RoutingHeader254 = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.RoutingHeader255 != nil {
		model.RoutingHeader255 = basetypes.NewBoolValue(*sdk.RoutingHeader255)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "RoutingHeader255", "value": *sdk.RoutingHeader255})
	} else {
		model.RoutingHeader255 = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.RoutingHeader3 != nil {
		model.RoutingHeader3 = basetypes.NewBoolValue(*sdk.RoutingHeader3)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "RoutingHeader3", "value": *sdk.RoutingHeader3})
	} else {
		model.RoutingHeader3 = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.RoutingHeader4252 != nil {
		model.RoutingHeader4252 = basetypes.NewBoolValue(*sdk.RoutingHeader4252)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "RoutingHeader4252", "value": *sdk.RoutingHeader4252})
	} else {
		model.RoutingHeader4252 = basetypes.NewBoolNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ZoneProtectionProfilesIpv6{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ZoneProtectionProfilesIpv6", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ZoneProtectionProfilesIpv6 ---
func unpackZoneProtectionProfilesIpv6ListToSdk(ctx context.Context, list types.List) ([]network_services.ZoneProtectionProfilesIpv6, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ZoneProtectionProfilesIpv6")
	diags := diag.Diagnostics{}
	var data []models.ZoneProtectionProfilesIpv6
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.ZoneProtectionProfilesIpv6, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ZoneProtectionProfilesIpv6{}.AttrTypes(), &item)
		unpacked, d := unpackZoneProtectionProfilesIpv6ToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ZoneProtectionProfilesIpv6", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ZoneProtectionProfilesIpv6 ---
func packZoneProtectionProfilesIpv6ListFromSdk(ctx context.Context, sdks []network_services.ZoneProtectionProfilesIpv6) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ZoneProtectionProfilesIpv6")
	diags := diag.Diagnostics{}
	var data []models.ZoneProtectionProfilesIpv6

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ZoneProtectionProfilesIpv6
		obj, d := packZoneProtectionProfilesIpv6FromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ZoneProtectionProfilesIpv6{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ZoneProtectionProfilesIpv6", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ZoneProtectionProfilesIpv6{}.AttrType(), data)
}

// --- Unpacker for ZoneProtectionProfilesIpv6FilterExtHdr ---
func unpackZoneProtectionProfilesIpv6FilterExtHdrToSdk(ctx context.Context, obj types.Object) (*network_services.ZoneProtectionProfilesIpv6FilterExtHdr, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ZoneProtectionProfilesIpv6FilterExtHdr", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ZoneProtectionProfilesIpv6FilterExtHdr
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.ZoneProtectionProfilesIpv6FilterExtHdr
	var d diag.Diagnostics
	// Handling Primitives
	if !model.DestOptionHdr.IsNull() && !model.DestOptionHdr.IsUnknown() {
		sdk.DestOptionHdr = model.DestOptionHdr.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DestOptionHdr", "value": *sdk.DestOptionHdr})
	}

	// Handling Primitives
	if !model.HopByHopHdr.IsNull() && !model.HopByHopHdr.IsUnknown() {
		sdk.HopByHopHdr = model.HopByHopHdr.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "HopByHopHdr", "value": *sdk.HopByHopHdr})
	}

	// Handling Primitives
	if !model.RoutingHdr.IsNull() && !model.RoutingHdr.IsUnknown() {
		sdk.RoutingHdr = model.RoutingHdr.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "RoutingHdr", "value": *sdk.RoutingHdr})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ZoneProtectionProfilesIpv6FilterExtHdr", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ZoneProtectionProfilesIpv6FilterExtHdr ---
func packZoneProtectionProfilesIpv6FilterExtHdrFromSdk(ctx context.Context, sdk network_services.ZoneProtectionProfilesIpv6FilterExtHdr) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ZoneProtectionProfilesIpv6FilterExtHdr", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ZoneProtectionProfilesIpv6FilterExtHdr
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.DestOptionHdr != nil {
		model.DestOptionHdr = basetypes.NewBoolValue(*sdk.DestOptionHdr)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DestOptionHdr", "value": *sdk.DestOptionHdr})
	} else {
		model.DestOptionHdr = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.HopByHopHdr != nil {
		model.HopByHopHdr = basetypes.NewBoolValue(*sdk.HopByHopHdr)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "HopByHopHdr", "value": *sdk.HopByHopHdr})
	} else {
		model.HopByHopHdr = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.RoutingHdr != nil {
		model.RoutingHdr = basetypes.NewBoolValue(*sdk.RoutingHdr)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "RoutingHdr", "value": *sdk.RoutingHdr})
	} else {
		model.RoutingHdr = basetypes.NewBoolNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ZoneProtectionProfilesIpv6FilterExtHdr{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ZoneProtectionProfilesIpv6FilterExtHdr", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ZoneProtectionProfilesIpv6FilterExtHdr ---
func unpackZoneProtectionProfilesIpv6FilterExtHdrListToSdk(ctx context.Context, list types.List) ([]network_services.ZoneProtectionProfilesIpv6FilterExtHdr, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ZoneProtectionProfilesIpv6FilterExtHdr")
	diags := diag.Diagnostics{}
	var data []models.ZoneProtectionProfilesIpv6FilterExtHdr
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.ZoneProtectionProfilesIpv6FilterExtHdr, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ZoneProtectionProfilesIpv6FilterExtHdr{}.AttrTypes(), &item)
		unpacked, d := unpackZoneProtectionProfilesIpv6FilterExtHdrToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ZoneProtectionProfilesIpv6FilterExtHdr", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ZoneProtectionProfilesIpv6FilterExtHdr ---
func packZoneProtectionProfilesIpv6FilterExtHdrListFromSdk(ctx context.Context, sdks []network_services.ZoneProtectionProfilesIpv6FilterExtHdr) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ZoneProtectionProfilesIpv6FilterExtHdr")
	diags := diag.Diagnostics{}
	var data []models.ZoneProtectionProfilesIpv6FilterExtHdr

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ZoneProtectionProfilesIpv6FilterExtHdr
		obj, d := packZoneProtectionProfilesIpv6FilterExtHdrFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ZoneProtectionProfilesIpv6FilterExtHdr{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ZoneProtectionProfilesIpv6FilterExtHdr", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ZoneProtectionProfilesIpv6FilterExtHdr{}.AttrType(), data)
}

// --- Unpacker for ZoneProtectionProfilesIpv6IgnoreInvPkt ---
func unpackZoneProtectionProfilesIpv6IgnoreInvPktToSdk(ctx context.Context, obj types.Object) (*network_services.ZoneProtectionProfilesIpv6IgnoreInvPkt, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ZoneProtectionProfilesIpv6IgnoreInvPkt", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ZoneProtectionProfilesIpv6IgnoreInvPkt
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.ZoneProtectionProfilesIpv6IgnoreInvPkt
	var d diag.Diagnostics
	// Handling Primitives
	if !model.DestUnreach.IsNull() && !model.DestUnreach.IsUnknown() {
		sdk.DestUnreach = model.DestUnreach.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DestUnreach", "value": *sdk.DestUnreach})
	}

	// Handling Primitives
	if !model.ParamProblem.IsNull() && !model.ParamProblem.IsUnknown() {
		sdk.ParamProblem = model.ParamProblem.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "ParamProblem", "value": *sdk.ParamProblem})
	}

	// Handling Primitives
	if !model.PktTooBig.IsNull() && !model.PktTooBig.IsUnknown() {
		sdk.PktTooBig = model.PktTooBig.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "PktTooBig", "value": *sdk.PktTooBig})
	}

	// Handling Primitives
	if !model.Redirect.IsNull() && !model.Redirect.IsUnknown() {
		sdk.Redirect = model.Redirect.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Redirect", "value": *sdk.Redirect})
	}

	// Handling Primitives
	if !model.TimeExceeded.IsNull() && !model.TimeExceeded.IsUnknown() {
		sdk.TimeExceeded = model.TimeExceeded.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "TimeExceeded", "value": *sdk.TimeExceeded})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ZoneProtectionProfilesIpv6IgnoreInvPkt", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ZoneProtectionProfilesIpv6IgnoreInvPkt ---
func packZoneProtectionProfilesIpv6IgnoreInvPktFromSdk(ctx context.Context, sdk network_services.ZoneProtectionProfilesIpv6IgnoreInvPkt) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ZoneProtectionProfilesIpv6IgnoreInvPkt", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ZoneProtectionProfilesIpv6IgnoreInvPkt
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.DestUnreach != nil {
		model.DestUnreach = basetypes.NewBoolValue(*sdk.DestUnreach)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DestUnreach", "value": *sdk.DestUnreach})
	} else {
		model.DestUnreach = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.ParamProblem != nil {
		model.ParamProblem = basetypes.NewBoolValue(*sdk.ParamProblem)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "ParamProblem", "value": *sdk.ParamProblem})
	} else {
		model.ParamProblem = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.PktTooBig != nil {
		model.PktTooBig = basetypes.NewBoolValue(*sdk.PktTooBig)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "PktTooBig", "value": *sdk.PktTooBig})
	} else {
		model.PktTooBig = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Redirect != nil {
		model.Redirect = basetypes.NewBoolValue(*sdk.Redirect)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Redirect", "value": *sdk.Redirect})
	} else {
		model.Redirect = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.TimeExceeded != nil {
		model.TimeExceeded = basetypes.NewBoolValue(*sdk.TimeExceeded)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "TimeExceeded", "value": *sdk.TimeExceeded})
	} else {
		model.TimeExceeded = basetypes.NewBoolNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ZoneProtectionProfilesIpv6IgnoreInvPkt{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ZoneProtectionProfilesIpv6IgnoreInvPkt", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ZoneProtectionProfilesIpv6IgnoreInvPkt ---
func unpackZoneProtectionProfilesIpv6IgnoreInvPktListToSdk(ctx context.Context, list types.List) ([]network_services.ZoneProtectionProfilesIpv6IgnoreInvPkt, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ZoneProtectionProfilesIpv6IgnoreInvPkt")
	diags := diag.Diagnostics{}
	var data []models.ZoneProtectionProfilesIpv6IgnoreInvPkt
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.ZoneProtectionProfilesIpv6IgnoreInvPkt, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ZoneProtectionProfilesIpv6IgnoreInvPkt{}.AttrTypes(), &item)
		unpacked, d := unpackZoneProtectionProfilesIpv6IgnoreInvPktToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ZoneProtectionProfilesIpv6IgnoreInvPkt", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ZoneProtectionProfilesIpv6IgnoreInvPkt ---
func packZoneProtectionProfilesIpv6IgnoreInvPktListFromSdk(ctx context.Context, sdks []network_services.ZoneProtectionProfilesIpv6IgnoreInvPkt) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ZoneProtectionProfilesIpv6IgnoreInvPkt")
	diags := diag.Diagnostics{}
	var data []models.ZoneProtectionProfilesIpv6IgnoreInvPkt

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ZoneProtectionProfilesIpv6IgnoreInvPkt
		obj, d := packZoneProtectionProfilesIpv6IgnoreInvPktFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ZoneProtectionProfilesIpv6IgnoreInvPkt{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ZoneProtectionProfilesIpv6IgnoreInvPkt", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ZoneProtectionProfilesIpv6IgnoreInvPkt{}.AttrType(), data)
}

// --- Unpacker for ZoneProtectionProfilesL2SecGroupTagProtection ---
func unpackZoneProtectionProfilesL2SecGroupTagProtectionToSdk(ctx context.Context, obj types.Object) (*network_services.ZoneProtectionProfilesL2SecGroupTagProtection, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ZoneProtectionProfilesL2SecGroupTagProtection", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ZoneProtectionProfilesL2SecGroupTagProtection
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.ZoneProtectionProfilesL2SecGroupTagProtection
	var d diag.Diagnostics
	// Handling Lists
	if !model.Tags.IsNull() && !model.Tags.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field Tags")
		unpacked, d := unpackZoneProtectionProfilesL2SecGroupTagProtectionTagsInnerListToSdk(ctx, model.Tags)
		diags.Append(d...)
		sdk.Tags = unpacked
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ZoneProtectionProfilesL2SecGroupTagProtection", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ZoneProtectionProfilesL2SecGroupTagProtection ---
func packZoneProtectionProfilesL2SecGroupTagProtectionFromSdk(ctx context.Context, sdk network_services.ZoneProtectionProfilesL2SecGroupTagProtection) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ZoneProtectionProfilesL2SecGroupTagProtection", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ZoneProtectionProfilesL2SecGroupTagProtection
	var d diag.Diagnostics
	// Handling Lists
	if sdk.Tags != nil {
		tflog.Debug(ctx, "Packing list of objects for field Tags")
		packed, d := packZoneProtectionProfilesL2SecGroupTagProtectionTagsInnerListFromSdk(ctx, sdk.Tags)
		diags.Append(d...)
		model.Tags = packed
	} else {
		model.Tags = basetypes.NewListNull(models.ZoneProtectionProfilesL2SecGroupTagProtectionTagsInner{}.AttrType())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ZoneProtectionProfilesL2SecGroupTagProtection{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ZoneProtectionProfilesL2SecGroupTagProtection", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ZoneProtectionProfilesL2SecGroupTagProtection ---
func unpackZoneProtectionProfilesL2SecGroupTagProtectionListToSdk(ctx context.Context, list types.List) ([]network_services.ZoneProtectionProfilesL2SecGroupTagProtection, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ZoneProtectionProfilesL2SecGroupTagProtection")
	diags := diag.Diagnostics{}
	var data []models.ZoneProtectionProfilesL2SecGroupTagProtection
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.ZoneProtectionProfilesL2SecGroupTagProtection, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ZoneProtectionProfilesL2SecGroupTagProtection{}.AttrTypes(), &item)
		unpacked, d := unpackZoneProtectionProfilesL2SecGroupTagProtectionToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ZoneProtectionProfilesL2SecGroupTagProtection", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ZoneProtectionProfilesL2SecGroupTagProtection ---
func packZoneProtectionProfilesL2SecGroupTagProtectionListFromSdk(ctx context.Context, sdks []network_services.ZoneProtectionProfilesL2SecGroupTagProtection) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ZoneProtectionProfilesL2SecGroupTagProtection")
	diags := diag.Diagnostics{}
	var data []models.ZoneProtectionProfilesL2SecGroupTagProtection

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ZoneProtectionProfilesL2SecGroupTagProtection
		obj, d := packZoneProtectionProfilesL2SecGroupTagProtectionFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ZoneProtectionProfilesL2SecGroupTagProtection{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ZoneProtectionProfilesL2SecGroupTagProtection", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ZoneProtectionProfilesL2SecGroupTagProtection{}.AttrType(), data)
}

// --- Unpacker for ZoneProtectionProfilesL2SecGroupTagProtectionTagsInner ---
func unpackZoneProtectionProfilesL2SecGroupTagProtectionTagsInnerToSdk(ctx context.Context, obj types.Object) (*network_services.ZoneProtectionProfilesL2SecGroupTagProtectionTagsInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ZoneProtectionProfilesL2SecGroupTagProtectionTagsInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ZoneProtectionProfilesL2SecGroupTagProtectionTagsInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.ZoneProtectionProfilesL2SecGroupTagProtectionTagsInner
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Enable.IsNull() && !model.Enable.IsUnknown() {
		sdk.Enable = model.Enable.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Enable", "value": *sdk.Enable})
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	// Handling Primitives
	if !model.Tag.IsNull() && !model.Tag.IsUnknown() {
		sdk.Tag = model.Tag.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Tag", "value": sdk.Tag})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ZoneProtectionProfilesL2SecGroupTagProtectionTagsInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ZoneProtectionProfilesL2SecGroupTagProtectionTagsInner ---
func packZoneProtectionProfilesL2SecGroupTagProtectionTagsInnerFromSdk(ctx context.Context, sdk network_services.ZoneProtectionProfilesL2SecGroupTagProtectionTagsInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ZoneProtectionProfilesL2SecGroupTagProtectionTagsInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ZoneProtectionProfilesL2SecGroupTagProtectionTagsInner
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Enable != nil {
		model.Enable = basetypes.NewBoolValue(*sdk.Enable)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Enable", "value": *sdk.Enable})
	} else {
		model.Enable = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	// Handling Primitives
	// Standard primitive packing
	model.Tag = basetypes.NewStringValue(sdk.Tag)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Tag", "value": sdk.Tag})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ZoneProtectionProfilesL2SecGroupTagProtectionTagsInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ZoneProtectionProfilesL2SecGroupTagProtectionTagsInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ZoneProtectionProfilesL2SecGroupTagProtectionTagsInner ---
func unpackZoneProtectionProfilesL2SecGroupTagProtectionTagsInnerListToSdk(ctx context.Context, list types.List) ([]network_services.ZoneProtectionProfilesL2SecGroupTagProtectionTagsInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ZoneProtectionProfilesL2SecGroupTagProtectionTagsInner")
	diags := diag.Diagnostics{}
	var data []models.ZoneProtectionProfilesL2SecGroupTagProtectionTagsInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.ZoneProtectionProfilesL2SecGroupTagProtectionTagsInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ZoneProtectionProfilesL2SecGroupTagProtectionTagsInner{}.AttrTypes(), &item)
		unpacked, d := unpackZoneProtectionProfilesL2SecGroupTagProtectionTagsInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ZoneProtectionProfilesL2SecGroupTagProtectionTagsInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ZoneProtectionProfilesL2SecGroupTagProtectionTagsInner ---
func packZoneProtectionProfilesL2SecGroupTagProtectionTagsInnerListFromSdk(ctx context.Context, sdks []network_services.ZoneProtectionProfilesL2SecGroupTagProtectionTagsInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ZoneProtectionProfilesL2SecGroupTagProtectionTagsInner")
	diags := diag.Diagnostics{}
	var data []models.ZoneProtectionProfilesL2SecGroupTagProtectionTagsInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ZoneProtectionProfilesL2SecGroupTagProtectionTagsInner
		obj, d := packZoneProtectionProfilesL2SecGroupTagProtectionTagsInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ZoneProtectionProfilesL2SecGroupTagProtectionTagsInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ZoneProtectionProfilesL2SecGroupTagProtectionTagsInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ZoneProtectionProfilesL2SecGroupTagProtectionTagsInner{}.AttrType(), data)
}

// --- Unpacker for ZoneProtectionProfilesNonIpProtocol ---
func unpackZoneProtectionProfilesNonIpProtocolToSdk(ctx context.Context, obj types.Object) (*network_services.ZoneProtectionProfilesNonIpProtocol, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ZoneProtectionProfilesNonIpProtocol", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ZoneProtectionProfilesNonIpProtocol
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.ZoneProtectionProfilesNonIpProtocol
	var d diag.Diagnostics
	// Handling Primitives
	if !model.ListType.IsNull() && !model.ListType.IsUnknown() {
		sdk.ListType = model.ListType.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "ListType", "value": *sdk.ListType})
	}

	// Handling Lists
	if !model.Protocol.IsNull() && !model.Protocol.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field Protocol")
		unpacked, d := unpackZoneProtectionProfilesNonIpProtocolProtocolInnerListToSdk(ctx, model.Protocol)
		diags.Append(d...)
		sdk.Protocol = unpacked
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ZoneProtectionProfilesNonIpProtocol", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ZoneProtectionProfilesNonIpProtocol ---
func packZoneProtectionProfilesNonIpProtocolFromSdk(ctx context.Context, sdk network_services.ZoneProtectionProfilesNonIpProtocol) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ZoneProtectionProfilesNonIpProtocol", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ZoneProtectionProfilesNonIpProtocol
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.ListType != nil {
		model.ListType = basetypes.NewStringValue(*sdk.ListType)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "ListType", "value": *sdk.ListType})
	} else {
		model.ListType = basetypes.NewStringNull()
	}
	// Handling Lists
	if sdk.Protocol != nil {
		tflog.Debug(ctx, "Packing list of objects for field Protocol")
		packed, d := packZoneProtectionProfilesNonIpProtocolProtocolInnerListFromSdk(ctx, sdk.Protocol)
		diags.Append(d...)
		model.Protocol = packed
	} else {
		model.Protocol = basetypes.NewListNull(models.ZoneProtectionProfilesNonIpProtocolProtocolInner{}.AttrType())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ZoneProtectionProfilesNonIpProtocol{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ZoneProtectionProfilesNonIpProtocol", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ZoneProtectionProfilesNonIpProtocol ---
func unpackZoneProtectionProfilesNonIpProtocolListToSdk(ctx context.Context, list types.List) ([]network_services.ZoneProtectionProfilesNonIpProtocol, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ZoneProtectionProfilesNonIpProtocol")
	diags := diag.Diagnostics{}
	var data []models.ZoneProtectionProfilesNonIpProtocol
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.ZoneProtectionProfilesNonIpProtocol, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ZoneProtectionProfilesNonIpProtocol{}.AttrTypes(), &item)
		unpacked, d := unpackZoneProtectionProfilesNonIpProtocolToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ZoneProtectionProfilesNonIpProtocol", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ZoneProtectionProfilesNonIpProtocol ---
func packZoneProtectionProfilesNonIpProtocolListFromSdk(ctx context.Context, sdks []network_services.ZoneProtectionProfilesNonIpProtocol) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ZoneProtectionProfilesNonIpProtocol")
	diags := diag.Diagnostics{}
	var data []models.ZoneProtectionProfilesNonIpProtocol

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ZoneProtectionProfilesNonIpProtocol
		obj, d := packZoneProtectionProfilesNonIpProtocolFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ZoneProtectionProfilesNonIpProtocol{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ZoneProtectionProfilesNonIpProtocol", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ZoneProtectionProfilesNonIpProtocol{}.AttrType(), data)
}

// --- Unpacker for ZoneProtectionProfilesNonIpProtocolProtocolInner ---
func unpackZoneProtectionProfilesNonIpProtocolProtocolInnerToSdk(ctx context.Context, obj types.Object) (*network_services.ZoneProtectionProfilesNonIpProtocolProtocolInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ZoneProtectionProfilesNonIpProtocolProtocolInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ZoneProtectionProfilesNonIpProtocolProtocolInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.ZoneProtectionProfilesNonIpProtocolProtocolInner
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Enable.IsNull() && !model.Enable.IsUnknown() {
		sdk.Enable = model.Enable.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Enable", "value": *sdk.Enable})
	}

	// Handling Primitives
	if !model.EtherType.IsNull() && !model.EtherType.IsUnknown() {
		sdk.EtherType = model.EtherType.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "EtherType", "value": sdk.EtherType})
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ZoneProtectionProfilesNonIpProtocolProtocolInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ZoneProtectionProfilesNonIpProtocolProtocolInner ---
func packZoneProtectionProfilesNonIpProtocolProtocolInnerFromSdk(ctx context.Context, sdk network_services.ZoneProtectionProfilesNonIpProtocolProtocolInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ZoneProtectionProfilesNonIpProtocolProtocolInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ZoneProtectionProfilesNonIpProtocolProtocolInner
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Enable != nil {
		model.Enable = basetypes.NewBoolValue(*sdk.Enable)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Enable", "value": *sdk.Enable})
	} else {
		model.Enable = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	model.EtherType = basetypes.NewStringValue(sdk.EtherType)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "EtherType", "value": sdk.EtherType})
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ZoneProtectionProfilesNonIpProtocolProtocolInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ZoneProtectionProfilesNonIpProtocolProtocolInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ZoneProtectionProfilesNonIpProtocolProtocolInner ---
func unpackZoneProtectionProfilesNonIpProtocolProtocolInnerListToSdk(ctx context.Context, list types.List) ([]network_services.ZoneProtectionProfilesNonIpProtocolProtocolInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ZoneProtectionProfilesNonIpProtocolProtocolInner")
	diags := diag.Diagnostics{}
	var data []models.ZoneProtectionProfilesNonIpProtocolProtocolInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.ZoneProtectionProfilesNonIpProtocolProtocolInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ZoneProtectionProfilesNonIpProtocolProtocolInner{}.AttrTypes(), &item)
		unpacked, d := unpackZoneProtectionProfilesNonIpProtocolProtocolInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ZoneProtectionProfilesNonIpProtocolProtocolInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ZoneProtectionProfilesNonIpProtocolProtocolInner ---
func packZoneProtectionProfilesNonIpProtocolProtocolInnerListFromSdk(ctx context.Context, sdks []network_services.ZoneProtectionProfilesNonIpProtocolProtocolInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ZoneProtectionProfilesNonIpProtocolProtocolInner")
	diags := diag.Diagnostics{}
	var data []models.ZoneProtectionProfilesNonIpProtocolProtocolInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ZoneProtectionProfilesNonIpProtocolProtocolInner
		obj, d := packZoneProtectionProfilesNonIpProtocolProtocolInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ZoneProtectionProfilesNonIpProtocolProtocolInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ZoneProtectionProfilesNonIpProtocolProtocolInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ZoneProtectionProfilesNonIpProtocolProtocolInner{}.AttrType(), data)
}

// --- Unpacker for ZoneProtectionProfilesScanInner ---
func unpackZoneProtectionProfilesScanInnerToSdk(ctx context.Context, obj types.Object) (*network_services.ZoneProtectionProfilesScanInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ZoneProtectionProfilesScanInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ZoneProtectionProfilesScanInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.ZoneProtectionProfilesScanInner
	var d diag.Diagnostics
	// Handling Objects
	if !model.Action.IsNull() && !model.Action.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Action")
		unpacked, d := unpackZoneProtectionProfilesScanInnerActionToSdk(ctx, model.Action)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Action"})
		}
		if unpacked != nil {
			sdk.Action = unpacked
		}
	}

	// Handling Primitives
	if !model.Interval.IsNull() && !model.Interval.IsUnknown() {
		val := int32(model.Interval.ValueInt64())
		sdk.Interval = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Interval", "value": *sdk.Interval})
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	// Handling Primitives
	if !model.Threshold.IsNull() && !model.Threshold.IsUnknown() {
		val := int32(model.Threshold.ValueInt64())
		sdk.Threshold = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Threshold", "value": *sdk.Threshold})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ZoneProtectionProfilesScanInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ZoneProtectionProfilesScanInner ---
func packZoneProtectionProfilesScanInnerFromSdk(ctx context.Context, sdk network_services.ZoneProtectionProfilesScanInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ZoneProtectionProfilesScanInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ZoneProtectionProfilesScanInner
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Action != nil {
		tflog.Debug(ctx, "Packing nested object for field Action")
		packed, d := packZoneProtectionProfilesScanInnerActionFromSdk(ctx, *sdk.Action)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Action"})
		}
		model.Action = packed
	} else {
		model.Action = basetypes.NewObjectNull(models.ZoneProtectionProfilesScanInnerAction{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Interval != nil {
		model.Interval = basetypes.NewInt64Value(int64(*sdk.Interval))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Interval", "value": *sdk.Interval})
	} else {
		model.Interval = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	// Handling Primitives
	// Standard primitive packing
	if sdk.Threshold != nil {
		model.Threshold = basetypes.NewInt64Value(int64(*sdk.Threshold))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Threshold", "value": *sdk.Threshold})
	} else {
		model.Threshold = basetypes.NewInt64Null()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ZoneProtectionProfilesScanInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ZoneProtectionProfilesScanInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ZoneProtectionProfilesScanInner ---
func unpackZoneProtectionProfilesScanInnerListToSdk(ctx context.Context, list types.List) ([]network_services.ZoneProtectionProfilesScanInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ZoneProtectionProfilesScanInner")
	diags := diag.Diagnostics{}
	var data []models.ZoneProtectionProfilesScanInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.ZoneProtectionProfilesScanInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ZoneProtectionProfilesScanInner{}.AttrTypes(), &item)
		unpacked, d := unpackZoneProtectionProfilesScanInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ZoneProtectionProfilesScanInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ZoneProtectionProfilesScanInner ---
func packZoneProtectionProfilesScanInnerListFromSdk(ctx context.Context, sdks []network_services.ZoneProtectionProfilesScanInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ZoneProtectionProfilesScanInner")
	diags := diag.Diagnostics{}
	var data []models.ZoneProtectionProfilesScanInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ZoneProtectionProfilesScanInner
		obj, d := packZoneProtectionProfilesScanInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ZoneProtectionProfilesScanInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ZoneProtectionProfilesScanInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ZoneProtectionProfilesScanInner{}.AttrType(), data)
}

// --- Unpacker for ZoneProtectionProfilesScanInnerAction ---
func unpackZoneProtectionProfilesScanInnerActionToSdk(ctx context.Context, obj types.Object) (*network_services.ZoneProtectionProfilesScanInnerAction, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ZoneProtectionProfilesScanInnerAction", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ZoneProtectionProfilesScanInnerAction
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.ZoneProtectionProfilesScanInnerAction
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

	// Handling Typeless Objects
	if !model.Block.IsNull() && !model.Block.IsUnknown() {
		tflog.Debug(ctx, "Unpacking typeless object for field Block")
		sdk.Block = make(map[string]interface{})
	}

	// Handling Objects
	if !model.BlockIp.IsNull() && !model.BlockIp.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field BlockIp")
		unpacked, d := unpackZoneProtectionProfilesScanInnerActionBlockIpToSdk(ctx, model.BlockIp)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "BlockIp"})
		}
		if unpacked != nil {
			sdk.BlockIp = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ZoneProtectionProfilesScanInnerAction", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ZoneProtectionProfilesScanInnerAction ---
func packZoneProtectionProfilesScanInnerActionFromSdk(ctx context.Context, sdk network_services.ZoneProtectionProfilesScanInnerAction) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ZoneProtectionProfilesScanInnerAction", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ZoneProtectionProfilesScanInnerAction
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
	// This is a marker object (e.g. CHAP: {}). We just need to create an empty, non-null object.
	if sdk.Block != nil && !reflect.ValueOf(sdk.Block).IsNil() {
		tflog.Debug(ctx, "Packing typeless object for field Block")
		var d diag.Diagnostics
		// Create an empty object with no attributes, which signifies its presence.
		model.Block, d = basetypes.NewObjectValue(map[string]attr.Type{}, map[string]attr.Value{})
		diags.Append(d...)
	} else {
		// Since this field is part of a oneOf, being nil means it's not selected.
		// We make the object null with an empty attribute map.
		model.Block = basetypes.NewObjectNull(map[string]attr.Type{})
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.BlockIp != nil {
		tflog.Debug(ctx, "Packing nested object for field BlockIp")
		packed, d := packZoneProtectionProfilesScanInnerActionBlockIpFromSdk(ctx, *sdk.BlockIp)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "BlockIp"})
		}
		model.BlockIp = packed
	} else {
		model.BlockIp = basetypes.NewObjectNull(models.ZoneProtectionProfilesScanInnerActionBlockIp{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ZoneProtectionProfilesScanInnerAction{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ZoneProtectionProfilesScanInnerAction", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ZoneProtectionProfilesScanInnerAction ---
func unpackZoneProtectionProfilesScanInnerActionListToSdk(ctx context.Context, list types.List) ([]network_services.ZoneProtectionProfilesScanInnerAction, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ZoneProtectionProfilesScanInnerAction")
	diags := diag.Diagnostics{}
	var data []models.ZoneProtectionProfilesScanInnerAction
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.ZoneProtectionProfilesScanInnerAction, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ZoneProtectionProfilesScanInnerAction{}.AttrTypes(), &item)
		unpacked, d := unpackZoneProtectionProfilesScanInnerActionToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ZoneProtectionProfilesScanInnerAction", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ZoneProtectionProfilesScanInnerAction ---
func packZoneProtectionProfilesScanInnerActionListFromSdk(ctx context.Context, sdks []network_services.ZoneProtectionProfilesScanInnerAction) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ZoneProtectionProfilesScanInnerAction")
	diags := diag.Diagnostics{}
	var data []models.ZoneProtectionProfilesScanInnerAction

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ZoneProtectionProfilesScanInnerAction
		obj, d := packZoneProtectionProfilesScanInnerActionFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ZoneProtectionProfilesScanInnerAction{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ZoneProtectionProfilesScanInnerAction", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ZoneProtectionProfilesScanInnerAction{}.AttrType(), data)
}

// --- Unpacker for ZoneProtectionProfilesScanInnerActionBlockIp ---
func unpackZoneProtectionProfilesScanInnerActionBlockIpToSdk(ctx context.Context, obj types.Object) (*network_services.ZoneProtectionProfilesScanInnerActionBlockIp, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ZoneProtectionProfilesScanInnerActionBlockIp", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ZoneProtectionProfilesScanInnerActionBlockIp
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.ZoneProtectionProfilesScanInnerActionBlockIp
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Duration.IsNull() && !model.Duration.IsUnknown() {
		sdk.Duration = int32(model.Duration.ValueInt64())
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Duration", "value": sdk.Duration})
	}

	// Handling Primitives
	if !model.TrackBy.IsNull() && !model.TrackBy.IsUnknown() {
		sdk.TrackBy = model.TrackBy.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "TrackBy", "value": sdk.TrackBy})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ZoneProtectionProfilesScanInnerActionBlockIp", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ZoneProtectionProfilesScanInnerActionBlockIp ---
func packZoneProtectionProfilesScanInnerActionBlockIpFromSdk(ctx context.Context, sdk network_services.ZoneProtectionProfilesScanInnerActionBlockIp) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ZoneProtectionProfilesScanInnerActionBlockIp", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ZoneProtectionProfilesScanInnerActionBlockIp
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	model.Duration = basetypes.NewInt64Value(int64(sdk.Duration))
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Duration", "value": sdk.Duration})
	// Handling Primitives
	// Standard primitive packing
	model.TrackBy = basetypes.NewStringValue(sdk.TrackBy)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "TrackBy", "value": sdk.TrackBy})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ZoneProtectionProfilesScanInnerActionBlockIp{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ZoneProtectionProfilesScanInnerActionBlockIp", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ZoneProtectionProfilesScanInnerActionBlockIp ---
func unpackZoneProtectionProfilesScanInnerActionBlockIpListToSdk(ctx context.Context, list types.List) ([]network_services.ZoneProtectionProfilesScanInnerActionBlockIp, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ZoneProtectionProfilesScanInnerActionBlockIp")
	diags := diag.Diagnostics{}
	var data []models.ZoneProtectionProfilesScanInnerActionBlockIp
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.ZoneProtectionProfilesScanInnerActionBlockIp, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ZoneProtectionProfilesScanInnerActionBlockIp{}.AttrTypes(), &item)
		unpacked, d := unpackZoneProtectionProfilesScanInnerActionBlockIpToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ZoneProtectionProfilesScanInnerActionBlockIp", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ZoneProtectionProfilesScanInnerActionBlockIp ---
func packZoneProtectionProfilesScanInnerActionBlockIpListFromSdk(ctx context.Context, sdks []network_services.ZoneProtectionProfilesScanInnerActionBlockIp) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ZoneProtectionProfilesScanInnerActionBlockIp")
	diags := diag.Diagnostics{}
	var data []models.ZoneProtectionProfilesScanInnerActionBlockIp

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ZoneProtectionProfilesScanInnerActionBlockIp
		obj, d := packZoneProtectionProfilesScanInnerActionBlockIpFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ZoneProtectionProfilesScanInnerActionBlockIp{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ZoneProtectionProfilesScanInnerActionBlockIp", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ZoneProtectionProfilesScanInnerActionBlockIp{}.AttrType(), data)
}

// --- Unpacker for ZoneProtectionProfilesScanWhiteListInner ---
func unpackZoneProtectionProfilesScanWhiteListInnerToSdk(ctx context.Context, obj types.Object) (*network_services.ZoneProtectionProfilesScanWhiteListInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ZoneProtectionProfilesScanWhiteListInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ZoneProtectionProfilesScanWhiteListInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.ZoneProtectionProfilesScanWhiteListInner
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Ipv4.IsNull() && !model.Ipv4.IsUnknown() {
		sdk.Ipv4 = model.Ipv4.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Ipv4", "value": *sdk.Ipv4})
	}

	// Handling Primitives
	if !model.Ipv6.IsNull() && !model.Ipv6.IsUnknown() {
		sdk.Ipv6 = model.Ipv6.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Ipv6", "value": *sdk.Ipv6})
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ZoneProtectionProfilesScanWhiteListInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ZoneProtectionProfilesScanWhiteListInner ---
func packZoneProtectionProfilesScanWhiteListInnerFromSdk(ctx context.Context, sdk network_services.ZoneProtectionProfilesScanWhiteListInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ZoneProtectionProfilesScanWhiteListInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ZoneProtectionProfilesScanWhiteListInner
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Ipv4 != nil {
		model.Ipv4 = basetypes.NewStringValue(*sdk.Ipv4)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Ipv4", "value": *sdk.Ipv4})
	} else {
		model.Ipv4 = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Ipv6 != nil {
		model.Ipv6 = basetypes.NewStringValue(*sdk.Ipv6)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Ipv6", "value": *sdk.Ipv6})
	} else {
		model.Ipv6 = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ZoneProtectionProfilesScanWhiteListInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ZoneProtectionProfilesScanWhiteListInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ZoneProtectionProfilesScanWhiteListInner ---
func unpackZoneProtectionProfilesScanWhiteListInnerListToSdk(ctx context.Context, list types.List) ([]network_services.ZoneProtectionProfilesScanWhiteListInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ZoneProtectionProfilesScanWhiteListInner")
	diags := diag.Diagnostics{}
	var data []models.ZoneProtectionProfilesScanWhiteListInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.ZoneProtectionProfilesScanWhiteListInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ZoneProtectionProfilesScanWhiteListInner{}.AttrTypes(), &item)
		unpacked, d := unpackZoneProtectionProfilesScanWhiteListInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ZoneProtectionProfilesScanWhiteListInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ZoneProtectionProfilesScanWhiteListInner ---
func packZoneProtectionProfilesScanWhiteListInnerListFromSdk(ctx context.Context, sdks []network_services.ZoneProtectionProfilesScanWhiteListInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ZoneProtectionProfilesScanWhiteListInner")
	diags := diag.Diagnostics{}
	var data []models.ZoneProtectionProfilesScanWhiteListInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ZoneProtectionProfilesScanWhiteListInner
		obj, d := packZoneProtectionProfilesScanWhiteListInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ZoneProtectionProfilesScanWhiteListInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ZoneProtectionProfilesScanWhiteListInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ZoneProtectionProfilesScanWhiteListInner{}.AttrType(), data)
}
