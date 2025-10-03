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

	"github.com/paloaltonetworks/scm-go/generated/network_services"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/network_services"
)

// --- Unpacker for InterfaceManagementProfiles ---
func unpackInterfaceManagementProfilesToSdk(ctx context.Context, obj types.Object) (*network_services.InterfaceManagementProfiles, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.InterfaceManagementProfiles", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.InterfaceManagementProfiles
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk network_services.InterfaceManagementProfiles
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
	if !model.Http.IsNull() && !model.Http.IsUnknown() {
		sdk.Http = model.Http.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Http", "value": *sdk.Http})
	}

	// Handling Primitives
	if !model.HttpOcsp.IsNull() && !model.HttpOcsp.IsUnknown() {
		sdk.HttpOcsp = model.HttpOcsp.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "HttpOcsp", "value": *sdk.HttpOcsp})
	}

	// Handling Primitives
	if !model.Https.IsNull() && !model.Https.IsUnknown() {
		sdk.Https = model.Https.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Https", "value": *sdk.Https})
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

	// Handling Lists
	if !model.PermittedIp.IsNull() && !model.PermittedIp.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field PermittedIp")
		diags.Append(model.PermittedIp.ElementsAs(ctx, &sdk.PermittedIp, false)...)
	}

	// Handling Primitives
	if !model.Ping.IsNull() && !model.Ping.IsUnknown() {
		sdk.Ping = model.Ping.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Ping", "value": *sdk.Ping})
	}

	// Handling Primitives
	if !model.ResponsePages.IsNull() && !model.ResponsePages.IsUnknown() {
		sdk.ResponsePages = model.ResponsePages.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "ResponsePages", "value": sdk.ResponsePages})
	}

	// Handling Primitives
	if !model.Snippet.IsNull() && !model.Snippet.IsUnknown() {
		sdk.Snippet = model.Snippet.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	}

	// Handling Primitives
	if !model.Ssh.IsNull() && !model.Ssh.IsUnknown() {
		sdk.Ssh = model.Ssh.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Ssh", "value": *sdk.Ssh})
	}

	// Handling Primitives
	if !model.Telnet.IsNull() && !model.Telnet.IsUnknown() {
		sdk.Telnet = model.Telnet.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Telnet", "value": *sdk.Telnet})
	}

	// Handling Primitives
	if !model.UseridService.IsNull() && !model.UseridService.IsUnknown() {
		sdk.UseridService = model.UseridService.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "UseridService", "value": *sdk.UseridService})
	}

	// Handling Primitives
	if !model.UseridSyslogListenerSsl.IsNull() && !model.UseridSyslogListenerSsl.IsUnknown() {
		sdk.UseridSyslogListenerSsl = model.UseridSyslogListenerSsl.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "UseridSyslogListenerSsl", "value": *sdk.UseridSyslogListenerSsl})
	}

	// Handling Primitives
	if !model.UseridSyslogListenerUdp.IsNull() && !model.UseridSyslogListenerUdp.IsUnknown() {
		sdk.UseridSyslogListenerUdp = model.UseridSyslogListenerUdp.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "UseridSyslogListenerUdp", "value": *sdk.UseridSyslogListenerUdp})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.InterfaceManagementProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for InterfaceManagementProfiles ---
func packInterfaceManagementProfilesFromSdk(ctx context.Context, sdk network_services.InterfaceManagementProfiles) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.InterfaceManagementProfiles", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.InterfaceManagementProfiles
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
	if sdk.Http != nil {
		model.Http = basetypes.NewBoolValue(*sdk.Http)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Http", "value": *sdk.Http})
	} else {
		model.Http = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.HttpOcsp != nil {
		model.HttpOcsp = basetypes.NewBoolValue(*sdk.HttpOcsp)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "HttpOcsp", "value": *sdk.HttpOcsp})
	} else {
		model.HttpOcsp = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Https != nil {
		model.Https = basetypes.NewBoolValue(*sdk.Https)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Https", "value": *sdk.Https})
	} else {
		model.Https = basetypes.NewBoolNull()
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
	// Handling Lists
	if sdk.PermittedIp != nil {
		tflog.Debug(ctx, "Packing list of primitives for field PermittedIp")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.PermittedIp, d = basetypes.NewListValueFrom(ctx, elemType, sdk.PermittedIp)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.PermittedIp = basetypes.NewListNull(elemType)
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Ping != nil {
		model.Ping = basetypes.NewBoolValue(*sdk.Ping)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Ping", "value": *sdk.Ping})
	} else {
		model.Ping = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Universal packer for interface{} types that are mapped to a StringValue in the model.
	// All underlying primitive types will be converted to their string representation.
	if sdk.ResponsePages != nil {
		tflog.Debug(ctx, "Packing interface value for field ResponsePages")
		switch v := sdk.ResponsePages.(type) {
		case string:
			model.ResponsePages = basetypes.NewStringValue(v)
		case int:
			model.ResponsePages = basetypes.NewStringValue(strconv.Itoa(v))
		case int64:
			model.ResponsePages = basetypes.NewStringValue(strconv.FormatInt(v, 10))
		case bool:
			model.ResponsePages = basetypes.NewStringValue(strconv.FormatBool(v))
		case float64:
			model.ResponsePages = basetypes.NewStringValue(strconv.FormatFloat(v, 'f', -1, 64))
		default:
			tflog.Warn(ctx, "Unexpected type for interface field", map[string]interface{}{"field": "ResponsePages", "type": reflect.TypeOf(v).String()})
			model.ResponsePages = basetypes.NewStringNull()
		}
	} else {
		model.ResponsePages = basetypes.NewStringNull()
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
	if sdk.Ssh != nil {
		model.Ssh = basetypes.NewBoolValue(*sdk.Ssh)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Ssh", "value": *sdk.Ssh})
	} else {
		model.Ssh = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Telnet != nil {
		model.Telnet = basetypes.NewBoolValue(*sdk.Telnet)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Telnet", "value": *sdk.Telnet})
	} else {
		model.Telnet = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.UseridService != nil {
		model.UseridService = basetypes.NewBoolValue(*sdk.UseridService)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "UseridService", "value": *sdk.UseridService})
	} else {
		model.UseridService = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.UseridSyslogListenerSsl != nil {
		model.UseridSyslogListenerSsl = basetypes.NewBoolValue(*sdk.UseridSyslogListenerSsl)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "UseridSyslogListenerSsl", "value": *sdk.UseridSyslogListenerSsl})
	} else {
		model.UseridSyslogListenerSsl = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.UseridSyslogListenerUdp != nil {
		model.UseridSyslogListenerUdp = basetypes.NewBoolValue(*sdk.UseridSyslogListenerUdp)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "UseridSyslogListenerUdp", "value": *sdk.UseridSyslogListenerUdp})
	} else {
		model.UseridSyslogListenerUdp = basetypes.NewBoolNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.InterfaceManagementProfiles{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.InterfaceManagementProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for InterfaceManagementProfiles ---
func unpackInterfaceManagementProfilesListToSdk(ctx context.Context, list types.List) ([]network_services.InterfaceManagementProfiles, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.InterfaceManagementProfiles")
	diags := diag.Diagnostics{}
	var data []models.InterfaceManagementProfiles
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.InterfaceManagementProfiles, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.InterfaceManagementProfiles{}.AttrTypes(), &item)
		unpacked, d := unpackInterfaceManagementProfilesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.InterfaceManagementProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for InterfaceManagementProfiles ---
func packInterfaceManagementProfilesListFromSdk(ctx context.Context, sdks []network_services.InterfaceManagementProfiles) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.InterfaceManagementProfiles")
	diags := diag.Diagnostics{}
	var data []models.InterfaceManagementProfiles

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.InterfaceManagementProfiles
		obj, d := packInterfaceManagementProfilesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.InterfaceManagementProfiles{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.InterfaceManagementProfiles", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.InterfaceManagementProfiles{}.AttrType(), data)
}
