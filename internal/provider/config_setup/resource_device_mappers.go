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

// unpackDevicesToSdkPut converts the Terraform model into the SDK DevicesPut struct for update operations.
func unpackDevicesToSdkPut(ctx context.Context, model *models.DevicesTf) (*config_setup.DevicesPut, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for DevicesPut")
	diags := diag.Diagnostics{}

	var sdk config_setup.DevicesPut

	if !model.Folder.IsNull() && !model.Folder.IsUnknown() {
		sdk.Folder = model.Folder.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Folder", "value": *sdk.Folder})
	}

	if !model.Description.IsNull() && !model.Description.IsUnknown() {
		sdk.Description = model.Description.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Description", "value": *sdk.Description})
	}

	if !model.DisplayName.IsNull() && !model.DisplayName.IsUnknown() {
		sdk.DisplayName = model.DisplayName.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DisplayName", "value": *sdk.DisplayName})
	}

	if !model.Labels.IsNull() && !model.Labels.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Labels")
		diags.Append(model.Labels.ElementsAs(ctx, &sdk.Labels, false)...)
	}

	if !model.Snippets.IsNull() && !model.Snippets.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Snippets")
		diags.Append(model.Snippets.ElementsAs(ctx, &sdk.Snippets, false)...)
	}

	tflog.Debug(ctx, "Exiting unpack helper for DevicesPut", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags
}

// packDevicesFromSdk converts the SDK Devices struct into the Terraform model.
func packDevicesFromSdk(ctx context.Context, sdk config_setup.Devices) (*models.DevicesTf, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for Devices", map[string]interface{}{"id": sdk.Id})
	diags := diag.Diagnostics{}
	var model models.DevicesTf

	// Required fields
	model.Id = basetypes.NewStringValue(sdk.Id)
	model.Name = basetypes.NewStringValue(sdk.Name)
	model.Folder = basetypes.NewStringValue(sdk.Folder)

	// Optional string pointer fields
	if sdk.Description != nil {
		model.Description = basetypes.NewStringValue(*sdk.Description)
	} else {
		model.Description = basetypes.NewStringNull()
	}

	if sdk.DisplayName != nil {
		model.DisplayName = basetypes.NewStringValue(*sdk.DisplayName)
	} else {
		model.DisplayName = basetypes.NewStringNull()
	}

	if sdk.Hostname != nil {
		model.Hostname = basetypes.NewStringValue(*sdk.Hostname)
	} else {
		model.Hostname = basetypes.NewStringNull()
	}

	if sdk.IpAddress != nil {
		model.IpAddress = basetypes.NewStringValue(*sdk.IpAddress)
	} else {
		model.IpAddress = basetypes.NewStringNull()
	}

	if sdk.IpV6Address != nil {
		model.IpV6Address = basetypes.NewStringValue(*sdk.IpV6Address)
	} else {
		model.IpV6Address = basetypes.NewStringNull()
	}

	if sdk.MacAddress != nil {
		model.MacAddress = basetypes.NewStringValue(*sdk.MacAddress)
	} else {
		model.MacAddress = basetypes.NewStringNull()
	}

	if sdk.Family != nil {
		model.Family = basetypes.NewStringValue(*sdk.Family)
	} else {
		model.Family = basetypes.NewStringNull()
	}

	if sdk.Model != nil {
		model.Model = basetypes.NewStringValue(*sdk.Model)
	} else {
		model.Model = basetypes.NewStringNull()
	}

	if sdk.SoftwareVersion != nil {
		model.SoftwareVersion = basetypes.NewStringValue(*sdk.SoftwareVersion)
	} else {
		model.SoftwareVersion = basetypes.NewStringNull()
	}

	if sdk.IsConnected != nil {
		model.IsConnected = basetypes.NewBoolValue(*sdk.IsConnected)
	} else {
		model.IsConnected = basetypes.NewBoolNull()
	}

	// List fields
	var elemType attr.Type = basetypes.StringType{}
	if sdk.Labels != nil {
		var d diag.Diagnostics
		model.Labels, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Labels)
		diags.Append(d...)
	} else {
		model.Labels = basetypes.NewListNull(elemType)
	}

	if sdk.Snippets != nil {
		var d diag.Diagnostics
		model.Snippets, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Snippets)
		diags.Append(d...)
	} else {
		model.Snippets = basetypes.NewListNull(elemType)
	}

	tflog.Debug(ctx, "Exiting pack helper for Devices", map[string]interface{}{"has_errors": diags.HasError()})
	return &model, diags
}

// packDevicesFromSdkToObject converts the SDK Devices struct into a Terraform object value.
func packDevicesFromSdkToObject(ctx context.Context, sdk config_setup.Devices) (types.Object, diag.Diagnostics) {
	model, diags := packDevicesFromSdk(ctx, sdk)
	if diags.HasError() {
		return basetypes.NewObjectNull(models.DevicesTf{}.AttrTypes()), diags
	}
	obj, d := types.ObjectValueFrom(ctx, models.DevicesTf{}.AttrTypes(), model)
	diags.Append(d...)
	return obj, diags
}

// packDevicesListFromSdk converts a list of SDK Devices structs into a Terraform list.
func packDevicesListFromSdk(ctx context.Context, sdks []config_setup.Devices) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for Devices")
	diags := diag.Diagnostics{}
	var data []models.DevicesTf

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		model, d := packDevicesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.DevicesTf{}.AttrType()), diags
		}
		data = append(data, *model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for Devices", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.DevicesTf{}.AttrType(), data)
}
