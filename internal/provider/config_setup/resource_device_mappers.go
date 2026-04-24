package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/config_setup"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/config_setup"
)

// --- Unpacker for Devices ---
func unpackDevicesToSdk(ctx context.Context, obj types.Object) (*config_setup.Devices, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.Devices", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.Devices
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk config_setup.Devices
	var d diag.Diagnostics

	// Handling Primitives
	if !model.AntiVirusVersion.IsNull() && !model.AntiVirusVersion.IsUnknown() {
		sdk.AntiVirusVersion = model.AntiVirusVersion.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "AntiVirusVersion", "value": *sdk.AntiVirusVersion})
	}

	// Handling Primitives
	if !model.AppReleaseDate.IsNull() && !model.AppReleaseDate.IsUnknown() {
		sdk.AppReleaseDate = model.AppReleaseDate.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "AppReleaseDate", "value": *sdk.AppReleaseDate})
	}

	// Handling Primitives
	if !model.AppVersion.IsNull() && !model.AppVersion.IsUnknown() {
		sdk.AppVersion = model.AppVersion.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "AppVersion", "value": *sdk.AppVersion})
	}

	// Handling Primitives
	if !model.AvReleaseDate.IsNull() && !model.AvReleaseDate.IsUnknown() {
		sdk.AvReleaseDate = model.AvReleaseDate.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "AvReleaseDate", "value": *sdk.AvReleaseDate})
	}

	// Handling Lists
	if !model.AvailableLicensess.IsNull() && !model.AvailableLicensess.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field AvailableLicensess")
		unpacked, d := unpackDevicesAvailableLicensessInnerListToSdk(ctx, model.AvailableLicensess)
		diags.Append(d...)
		sdk.AvailableLicensess = unpacked
	}

	// Handling Primitives
	if !model.ConnectedSince.IsNull() && !model.ConnectedSince.IsUnknown() {
		sdk.ConnectedSince = model.ConnectedSince.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "ConnectedSince", "value": *sdk.ConnectedSince})
	}

	// Handling Primitives
	if !model.Description.IsNull() && !model.Description.IsUnknown() {
		sdk.Description = model.Description.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Description", "value": *sdk.Description})
	}

	// Handling Primitives
	if !model.DevCertDetail.IsNull() && !model.DevCertDetail.IsUnknown() {
		sdk.DevCertDetail = model.DevCertDetail.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DevCertDetail", "value": *sdk.DevCertDetail})
	}

	// Handling Primitives
	if !model.DevCertExpiryDate.IsNull() && !model.DevCertExpiryDate.IsUnknown() {
		sdk.DevCertExpiryDate = model.DevCertExpiryDate.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DevCertExpiryDate", "value": *sdk.DevCertExpiryDate})
	}

	// Handling Primitives
	if !model.DisplayName.IsNull() && !model.DisplayName.IsUnknown() {
		sdk.DisplayName = model.DisplayName.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DisplayName", "value": *sdk.DisplayName})
	}

	// Handling Primitives
	if !model.Family.IsNull() && !model.Family.IsUnknown() {
		sdk.Family = model.Family.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Family", "value": *sdk.Family})
	}

	// Handling Primitives
	if !model.Folder.IsNull() && !model.Folder.IsUnknown() {
		sdk.Folder = model.Folder.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Folder", "value": sdk.Folder})
	}

	// Handling Primitives
	if !model.GpClientVerion.IsNull() && !model.GpClientVerion.IsUnknown() {
		sdk.GpClientVerion = model.GpClientVerion.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "GpClientVerion", "value": *sdk.GpClientVerion})
	}

	// Handling Primitives
	if !model.GpDataVersion.IsNull() && !model.GpDataVersion.IsUnknown() {
		sdk.GpDataVersion = model.GpDataVersion.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "GpDataVersion", "value": *sdk.GpDataVersion})
	}

	// Handling Primitives
	if !model.HaPeerSerial.IsNull() && !model.HaPeerSerial.IsUnknown() {
		sdk.HaPeerSerial = model.HaPeerSerial.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "HaPeerSerial", "value": *sdk.HaPeerSerial})
	}

	// Handling Primitives
	if !model.HaPeerState.IsNull() && !model.HaPeerState.IsUnknown() {
		sdk.HaPeerState = model.HaPeerState.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "HaPeerState", "value": *sdk.HaPeerState})
	}

	// Handling Primitives
	if !model.HaState.IsNull() && !model.HaState.IsUnknown() {
		sdk.HaState = model.HaState.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "HaState", "value": *sdk.HaState})
	}

	// Handling Primitives
	if !model.Hostname.IsNull() && !model.Hostname.IsUnknown() {
		sdk.Hostname = model.Hostname.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Hostname", "value": *sdk.Hostname})
	}

	// Handling Primitives
	if !model.Id.IsNull() && !model.Id.IsUnknown() {
		sdk.Id = model.Id.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Id", "value": sdk.Id})
	}

	// Handling Lists
	if !model.InstalledLicenses.IsNull() && !model.InstalledLicenses.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field InstalledLicenses")
		unpacked, d := unpackDevicesInstalledLicensesInnerListToSdk(ctx, model.InstalledLicenses)
		diags.Append(d...)
		sdk.InstalledLicenses = unpacked
	}

	// Handling Primitives
	if !model.IotReleaseDate.IsNull() && !model.IotReleaseDate.IsUnknown() {
		sdk.IotReleaseDate = model.IotReleaseDate.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "IotReleaseDate", "value": *sdk.IotReleaseDate})
	}

	// Handling Primitives
	if !model.IotVersion.IsNull() && !model.IotVersion.IsUnknown() {
		sdk.IotVersion = model.IotVersion.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "IotVersion", "value": *sdk.IotVersion})
	}

	// Handling Primitives
	if !model.IpAddress.IsNull() && !model.IpAddress.IsUnknown() {
		sdk.IpAddress = model.IpAddress.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "IpAddress", "value": *sdk.IpAddress})
	}

	// Handling Primitives
	if !model.IpV6Address.IsNull() && !model.IpV6Address.IsUnknown() {
		sdk.IpV6Address = model.IpV6Address.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "IpV6Address", "value": *sdk.IpV6Address})
	}

	// Handling Primitives
	if !model.IsConnected.IsNull() && !model.IsConnected.IsUnknown() {
		sdk.IsConnected = model.IsConnected.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "IsConnected", "value": *sdk.IsConnected})
	}

	// Handling Lists
	if !model.Labels.IsNull() && !model.Labels.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Labels")
		diags.Append(model.Labels.ElementsAs(ctx, &sdk.Labels, false)...)
	}

	// Handling Primitives
	if !model.LicenseMatch.IsNull() && !model.LicenseMatch.IsUnknown() {
		sdk.LicenseMatch = model.LicenseMatch.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "LicenseMatch", "value": *sdk.LicenseMatch})
	}

	// Handling Primitives
	if !model.LogDbVersion.IsNull() && !model.LogDbVersion.IsUnknown() {
		sdk.LogDbVersion = model.LogDbVersion.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "LogDbVersion", "value": *sdk.LogDbVersion})
	}

	// Handling Primitives
	if !model.MacAddress.IsNull() && !model.MacAddress.IsUnknown() {
		sdk.MacAddress = model.MacAddress.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "MacAddress", "value": *sdk.MacAddress})
	}

	// Handling Primitives
	if !model.Model.IsNull() && !model.Model.IsUnknown() {
		sdk.Model = model.Model.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Model", "value": *sdk.Model})
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	// Handling Lists
	if !model.Snippets.IsNull() && !model.Snippets.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Snippets")
		diags.Append(model.Snippets.ElementsAs(ctx, &sdk.Snippets, false)...)
	}

	// Handling Primitives
	if !model.SoftwareVersion.IsNull() && !model.SoftwareVersion.IsUnknown() {
		sdk.SoftwareVersion = model.SoftwareVersion.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "SoftwareVersion", "value": *sdk.SoftwareVersion})
	}

	// Handling Primitives
	if !model.ThreatReleaseDate.IsNull() && !model.ThreatReleaseDate.IsUnknown() {
		sdk.ThreatReleaseDate = model.ThreatReleaseDate.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "ThreatReleaseDate", "value": *sdk.ThreatReleaseDate})
	}

	// Handling Primitives
	if !model.ThreatVersion.IsNull() && !model.ThreatVersion.IsUnknown() {
		sdk.ThreatVersion = model.ThreatVersion.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "ThreatVersion", "value": *sdk.ThreatVersion})
	}

	// Handling Primitives
	if !model.Uptime.IsNull() && !model.Uptime.IsUnknown() {
		sdk.Uptime = model.Uptime.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Uptime", "value": *sdk.Uptime})
	}

	// Handling Primitives
	if !model.UrlDbType.IsNull() && !model.UrlDbType.IsUnknown() {
		sdk.UrlDbType = model.UrlDbType.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "UrlDbType", "value": *sdk.UrlDbType})
	}

	// Handling Primitives
	if !model.UrlDbVer.IsNull() && !model.UrlDbVer.IsUnknown() {
		sdk.UrlDbVer = model.UrlDbVer.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "UrlDbVer", "value": *sdk.UrlDbVer})
	}

	// Handling Primitives
	if !model.VmState.IsNull() && !model.VmState.IsUnknown() {
		sdk.VmState = model.VmState.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "VmState", "value": *sdk.VmState})
	}

	// Handling Primitives
	if !model.WfReleaseDate.IsNull() && !model.WfReleaseDate.IsUnknown() {
		sdk.WfReleaseDate = model.WfReleaseDate.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "WfReleaseDate", "value": *sdk.WfReleaseDate})
	}

	// Handling Primitives
	if !model.WfVer.IsNull() && !model.WfVer.IsUnknown() {
		sdk.WfVer = model.WfVer.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "WfVer", "value": *sdk.WfVer})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.Devices", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for Devices ---
func packDevicesFromSdk(ctx context.Context, sdk config_setup.Devices) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.Devices", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.Devices
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.AntiVirusVersion != nil {
		model.AntiVirusVersion = basetypes.NewStringValue(*sdk.AntiVirusVersion)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "AntiVirusVersion", "value": *sdk.AntiVirusVersion})
	} else {
		model.AntiVirusVersion = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.AppReleaseDate != nil {
		model.AppReleaseDate = basetypes.NewStringValue(*sdk.AppReleaseDate)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "AppReleaseDate", "value": *sdk.AppReleaseDate})
	} else {
		model.AppReleaseDate = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.AppVersion != nil {
		model.AppVersion = basetypes.NewStringValue(*sdk.AppVersion)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "AppVersion", "value": *sdk.AppVersion})
	} else {
		model.AppVersion = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.AvReleaseDate != nil {
		model.AvReleaseDate = basetypes.NewStringValue(*sdk.AvReleaseDate)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "AvReleaseDate", "value": *sdk.AvReleaseDate})
	} else {
		model.AvReleaseDate = basetypes.NewStringNull()
	}
	// Handling Lists
	if sdk.AvailableLicensess != nil {
		tflog.Debug(ctx, "Packing list of objects for field AvailableLicensess")
		packed, d := packDevicesAvailableLicensessInnerListFromSdk(ctx, sdk.AvailableLicensess)
		diags.Append(d...)
		model.AvailableLicensess = packed
	} else {
		model.AvailableLicensess = basetypes.NewListNull(models.DevicesAvailableLicensessInner{}.AttrType())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.ConnectedSince != nil {
		model.ConnectedSince = basetypes.NewStringValue(*sdk.ConnectedSince)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "ConnectedSince", "value": *sdk.ConnectedSince})
	} else {
		model.ConnectedSince = basetypes.NewStringNull()
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
	if sdk.DevCertDetail != nil {
		model.DevCertDetail = basetypes.NewStringValue(*sdk.DevCertDetail)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DevCertDetail", "value": *sdk.DevCertDetail})
	} else {
		model.DevCertDetail = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.DevCertExpiryDate != nil {
		model.DevCertExpiryDate = basetypes.NewStringValue(*sdk.DevCertExpiryDate)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DevCertExpiryDate", "value": *sdk.DevCertExpiryDate})
	} else {
		model.DevCertExpiryDate = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.DisplayName != nil {
		model.DisplayName = basetypes.NewStringValue(*sdk.DisplayName)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DisplayName", "value": *sdk.DisplayName})
	} else {
		model.DisplayName = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Family != nil {
		model.Family = basetypes.NewStringValue(*sdk.Family)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Family", "value": *sdk.Family})
	} else {
		model.Family = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	model.Folder = basetypes.NewStringValue(sdk.Folder)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Folder", "value": sdk.Folder})
	// Handling Primitives
	// Standard primitive packing
	if sdk.GpClientVerion != nil {
		model.GpClientVerion = basetypes.NewStringValue(*sdk.GpClientVerion)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "GpClientVerion", "value": *sdk.GpClientVerion})
	} else {
		model.GpClientVerion = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.GpDataVersion != nil {
		model.GpDataVersion = basetypes.NewStringValue(*sdk.GpDataVersion)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "GpDataVersion", "value": *sdk.GpDataVersion})
	} else {
		model.GpDataVersion = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.HaPeerSerial != nil {
		model.HaPeerSerial = basetypes.NewStringValue(*sdk.HaPeerSerial)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "HaPeerSerial", "value": *sdk.HaPeerSerial})
	} else {
		model.HaPeerSerial = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.HaPeerState != nil {
		model.HaPeerState = basetypes.NewStringValue(*sdk.HaPeerState)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "HaPeerState", "value": *sdk.HaPeerState})
	} else {
		model.HaPeerState = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.HaState != nil {
		model.HaState = basetypes.NewStringValue(*sdk.HaState)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "HaState", "value": *sdk.HaState})
	} else {
		model.HaState = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Hostname != nil {
		model.Hostname = basetypes.NewStringValue(*sdk.Hostname)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Hostname", "value": *sdk.Hostname})
	} else {
		model.Hostname = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	model.Id = basetypes.NewStringValue(sdk.Id)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Id", "value": sdk.Id})
	// Handling Lists
	if sdk.InstalledLicenses != nil {
		tflog.Debug(ctx, "Packing list of objects for field InstalledLicenses")
		packed, d := packDevicesInstalledLicensesInnerListFromSdk(ctx, sdk.InstalledLicenses)
		diags.Append(d...)
		model.InstalledLicenses = packed
	} else {
		model.InstalledLicenses = basetypes.NewListNull(models.DevicesInstalledLicensesInner{}.AttrType())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.IotReleaseDate != nil {
		model.IotReleaseDate = basetypes.NewStringValue(*sdk.IotReleaseDate)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "IotReleaseDate", "value": *sdk.IotReleaseDate})
	} else {
		model.IotReleaseDate = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.IotVersion != nil {
		model.IotVersion = basetypes.NewStringValue(*sdk.IotVersion)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "IotVersion", "value": *sdk.IotVersion})
	} else {
		model.IotVersion = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.IpAddress != nil {
		model.IpAddress = basetypes.NewStringValue(*sdk.IpAddress)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "IpAddress", "value": *sdk.IpAddress})
	} else {
		model.IpAddress = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.IpV6Address != nil {
		model.IpV6Address = basetypes.NewStringValue(*sdk.IpV6Address)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "IpV6Address", "value": *sdk.IpV6Address})
	} else {
		model.IpV6Address = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.IsConnected != nil {
		model.IsConnected = basetypes.NewBoolValue(*sdk.IsConnected)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "IsConnected", "value": *sdk.IsConnected})
	} else {
		model.IsConnected = basetypes.NewBoolNull()
	}
	// Handling Lists
	if sdk.Labels != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Labels")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Labels, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Labels)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Labels = basetypes.NewListNull(elemType)
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.LicenseMatch != nil {
		model.LicenseMatch = basetypes.NewBoolValue(*sdk.LicenseMatch)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "LicenseMatch", "value": *sdk.LicenseMatch})
	} else {
		model.LicenseMatch = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.LogDbVersion != nil {
		model.LogDbVersion = basetypes.NewStringValue(*sdk.LogDbVersion)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "LogDbVersion", "value": *sdk.LogDbVersion})
	} else {
		model.LogDbVersion = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.MacAddress != nil {
		model.MacAddress = basetypes.NewStringValue(*sdk.MacAddress)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "MacAddress", "value": *sdk.MacAddress})
	} else {
		model.MacAddress = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Model != nil {
		model.Model = basetypes.NewStringValue(*sdk.Model)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Model", "value": *sdk.Model})
	} else {
		model.Model = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	// Handling Lists
	if sdk.Snippets != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Snippets")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Snippets, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Snippets)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Snippets = basetypes.NewListNull(elemType)
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.SoftwareVersion != nil {
		model.SoftwareVersion = basetypes.NewStringValue(*sdk.SoftwareVersion)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "SoftwareVersion", "value": *sdk.SoftwareVersion})
	} else {
		model.SoftwareVersion = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.ThreatReleaseDate != nil {
		model.ThreatReleaseDate = basetypes.NewStringValue(*sdk.ThreatReleaseDate)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "ThreatReleaseDate", "value": *sdk.ThreatReleaseDate})
	} else {
		model.ThreatReleaseDate = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.ThreatVersion != nil {
		model.ThreatVersion = basetypes.NewStringValue(*sdk.ThreatVersion)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "ThreatVersion", "value": *sdk.ThreatVersion})
	} else {
		model.ThreatVersion = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Uptime != nil {
		model.Uptime = basetypes.NewStringValue(*sdk.Uptime)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Uptime", "value": *sdk.Uptime})
	} else {
		model.Uptime = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.UrlDbType != nil {
		model.UrlDbType = basetypes.NewStringValue(*sdk.UrlDbType)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "UrlDbType", "value": *sdk.UrlDbType})
	} else {
		model.UrlDbType = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.UrlDbVer != nil {
		model.UrlDbVer = basetypes.NewStringValue(*sdk.UrlDbVer)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "UrlDbVer", "value": *sdk.UrlDbVer})
	} else {
		model.UrlDbVer = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.VmState != nil {
		model.VmState = basetypes.NewStringValue(*sdk.VmState)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "VmState", "value": *sdk.VmState})
	} else {
		model.VmState = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.WfReleaseDate != nil {
		model.WfReleaseDate = basetypes.NewStringValue(*sdk.WfReleaseDate)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "WfReleaseDate", "value": *sdk.WfReleaseDate})
	} else {
		model.WfReleaseDate = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.WfVer != nil {
		model.WfVer = basetypes.NewStringValue(*sdk.WfVer)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "WfVer", "value": *sdk.WfVer})
	} else {
		model.WfVer = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.Devices{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.Devices", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for Devices ---
func unpackDevicesListToSdk(ctx context.Context, list types.List) ([]config_setup.Devices, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.Devices")
	diags := diag.Diagnostics{}
	var data []models.Devices
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]config_setup.Devices, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.Devices{}.AttrTypes(), &item)
		unpacked, d := unpackDevicesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.Devices", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for Devices ---
func packDevicesListFromSdk(ctx context.Context, sdks []config_setup.Devices) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.Devices")
	diags := diag.Diagnostics{}
	var data []models.Devices

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.Devices
		obj, d := packDevicesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.Devices{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.Devices", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.Devices{}.AttrType(), data)
}

// --- Unpacker for DevicesAvailableLicensessInner ---
func unpackDevicesAvailableLicensessInnerToSdk(ctx context.Context, obj types.Object) (*config_setup.DevicesAvailableLicensessInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.DevicesAvailableLicensessInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.DevicesAvailableLicensessInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk config_setup.DevicesAvailableLicensessInner
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Authcode.IsNull() && !model.Authcode.IsUnknown() {
		sdk.Authcode = model.Authcode.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Authcode", "value": *sdk.Authcode})
	}

	// Handling Primitives
	if !model.Expires.IsNull() && !model.Expires.IsUnknown() {
		sdk.Expires = model.Expires.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Expires", "value": *sdk.Expires})
	}

	// Handling Primitives
	if !model.Feature.IsNull() && !model.Feature.IsUnknown() {
		sdk.Feature = model.Feature.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Feature", "value": *sdk.Feature})
	}

	// Handling Primitives
	if !model.Issued.IsNull() && !model.Issued.IsUnknown() {
		sdk.Issued = model.Issued.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Issued", "value": *sdk.Issued})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.DevicesAvailableLicensessInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for DevicesAvailableLicensessInner ---
func packDevicesAvailableLicensessInnerFromSdk(ctx context.Context, sdk config_setup.DevicesAvailableLicensessInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.DevicesAvailableLicensessInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.DevicesAvailableLicensessInner
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Authcode != nil {
		model.Authcode = basetypes.NewStringValue(*sdk.Authcode)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Authcode", "value": *sdk.Authcode})
	} else {
		model.Authcode = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Expires != nil {
		model.Expires = basetypes.NewStringValue(*sdk.Expires)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Expires", "value": *sdk.Expires})
	} else {
		model.Expires = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Feature != nil {
		model.Feature = basetypes.NewStringValue(*sdk.Feature)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Feature", "value": *sdk.Feature})
	} else {
		model.Feature = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Issued != nil {
		model.Issued = basetypes.NewStringValue(*sdk.Issued)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Issued", "value": *sdk.Issued})
	} else {
		model.Issued = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.DevicesAvailableLicensessInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.DevicesAvailableLicensessInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for DevicesAvailableLicensessInner ---
func unpackDevicesAvailableLicensessInnerListToSdk(ctx context.Context, list types.List) ([]config_setup.DevicesAvailableLicensessInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.DevicesAvailableLicensessInner")
	diags := diag.Diagnostics{}
	var data []models.DevicesAvailableLicensessInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]config_setup.DevicesAvailableLicensessInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.DevicesAvailableLicensessInner{}.AttrTypes(), &item)
		unpacked, d := unpackDevicesAvailableLicensessInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.DevicesAvailableLicensessInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for DevicesAvailableLicensessInner ---
func packDevicesAvailableLicensessInnerListFromSdk(ctx context.Context, sdks []config_setup.DevicesAvailableLicensessInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.DevicesAvailableLicensessInner")
	diags := diag.Diagnostics{}
	var data []models.DevicesAvailableLicensessInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.DevicesAvailableLicensessInner
		obj, d := packDevicesAvailableLicensessInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.DevicesAvailableLicensessInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.DevicesAvailableLicensessInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.DevicesAvailableLicensessInner{}.AttrType(), data)
}

// --- Unpacker for DevicesInstalledLicensesInner ---
func unpackDevicesInstalledLicensesInnerToSdk(ctx context.Context, obj types.Object) (*config_setup.DevicesInstalledLicensesInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.DevicesInstalledLicensesInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.DevicesInstalledLicensesInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk config_setup.DevicesInstalledLicensesInner
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Authcode.IsNull() && !model.Authcode.IsUnknown() {
		sdk.Authcode = model.Authcode.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Authcode", "value": *sdk.Authcode})
	}

	// Handling Primitives
	if !model.Expired.IsNull() && !model.Expired.IsUnknown() {
		sdk.Expired = model.Expired.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Expired", "value": *sdk.Expired})
	}

	// Handling Primitives
	if !model.Expires.IsNull() && !model.Expires.IsUnknown() {
		sdk.Expires = model.Expires.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Expires", "value": *sdk.Expires})
	}

	// Handling Primitives
	if !model.Feature.IsNull() && !model.Feature.IsUnknown() {
		sdk.Feature = model.Feature.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Feature", "value": *sdk.Feature})
	}

	// Handling Primitives
	if !model.Issued.IsNull() && !model.Issued.IsUnknown() {
		sdk.Issued = model.Issued.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Issued", "value": *sdk.Issued})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.DevicesInstalledLicensesInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for DevicesInstalledLicensesInner ---
func packDevicesInstalledLicensesInnerFromSdk(ctx context.Context, sdk config_setup.DevicesInstalledLicensesInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.DevicesInstalledLicensesInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.DevicesInstalledLicensesInner
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Authcode != nil {
		model.Authcode = basetypes.NewStringValue(*sdk.Authcode)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Authcode", "value": *sdk.Authcode})
	} else {
		model.Authcode = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Expired != nil {
		model.Expired = basetypes.NewStringValue(*sdk.Expired)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Expired", "value": *sdk.Expired})
	} else {
		model.Expired = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Expires != nil {
		model.Expires = basetypes.NewStringValue(*sdk.Expires)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Expires", "value": *sdk.Expires})
	} else {
		model.Expires = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Feature != nil {
		model.Feature = basetypes.NewStringValue(*sdk.Feature)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Feature", "value": *sdk.Feature})
	} else {
		model.Feature = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Issued != nil {
		model.Issued = basetypes.NewStringValue(*sdk.Issued)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Issued", "value": *sdk.Issued})
	} else {
		model.Issued = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.DevicesInstalledLicensesInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.DevicesInstalledLicensesInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for DevicesInstalledLicensesInner ---
func unpackDevicesInstalledLicensesInnerListToSdk(ctx context.Context, list types.List) ([]config_setup.DevicesInstalledLicensesInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.DevicesInstalledLicensesInner")
	diags := diag.Diagnostics{}
	var data []models.DevicesInstalledLicensesInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]config_setup.DevicesInstalledLicensesInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.DevicesInstalledLicensesInner{}.AttrTypes(), &item)
		unpacked, d := unpackDevicesInstalledLicensesInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.DevicesInstalledLicensesInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for DevicesInstalledLicensesInner ---
func packDevicesInstalledLicensesInnerListFromSdk(ctx context.Context, sdks []config_setup.DevicesInstalledLicensesInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.DevicesInstalledLicensesInner")
	diags := diag.Diagnostics{}
	var data []models.DevicesInstalledLicensesInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.DevicesInstalledLicensesInner
		obj, d := packDevicesInstalledLicensesInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.DevicesInstalledLicensesInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.DevicesInstalledLicensesInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.DevicesInstalledLicensesInner{}.AttrType(), data)
}
