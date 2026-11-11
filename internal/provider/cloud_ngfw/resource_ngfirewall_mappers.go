package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/cloud_ngfw"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/cloud_ngfw"
)

// --- Unpacker for Ngfirewalls ---
func unpackNgfirewallsToSdk(ctx context.Context, obj types.Object) (*cloud_ngfw.Ngfirewalls, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.Ngfirewalls", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.Ngfirewalls
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk cloud_ngfw.Ngfirewalls
	var d diag.Diagnostics

	// Handling Objects
	if !model.Response.IsNull() && !model.Response.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Response")
		unpacked, d := unpackNgfirewallResponseBodyToSdk(ctx, model.Response)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Response"})
		}
		if unpacked != nil {
			sdk.Response = unpacked
		}
	}

	// Handling Objects
	if !model.ResponseStatus.IsNull() && !model.ResponseStatus.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field ResponseStatus")
		unpacked, d := unpackNgfirewallResponseStatusToSdk(ctx, model.ResponseStatus)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "ResponseStatus"})
		}
		if unpacked != nil {
			sdk.ResponseStatus = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.Ngfirewalls", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for Ngfirewalls ---
func packNgfirewallsFromSdk(ctx context.Context, sdk cloud_ngfw.Ngfirewalls) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.Ngfirewalls", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.Ngfirewalls
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Response != nil {
		tflog.Debug(ctx, "Packing nested object for field Response")
		packed, d := packNgfirewallResponseBodyFromSdk(ctx, *sdk.Response)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Response"})
		}
		model.Response = packed
	} else {
		model.Response = basetypes.NewObjectNull(models.NgfirewallResponseBody{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.ResponseStatus != nil {
		tflog.Debug(ctx, "Packing nested object for field ResponseStatus")
		packed, d := packNgfirewallResponseStatusFromSdk(ctx, *sdk.ResponseStatus)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "ResponseStatus"})
		}
		model.ResponseStatus = packed
	} else {
		model.ResponseStatus = basetypes.NewObjectNull(models.NgfirewallResponseStatus{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.Ngfirewalls{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.Ngfirewalls", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for Ngfirewalls ---
func unpackNgfirewallsListToSdk(ctx context.Context, list types.List) ([]cloud_ngfw.Ngfirewalls, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.Ngfirewalls")
	diags := diag.Diagnostics{}
	var data []models.Ngfirewalls
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]cloud_ngfw.Ngfirewalls, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.Ngfirewalls{}.AttrTypes(), &item)
		unpacked, d := unpackNgfirewallsToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.Ngfirewalls", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for Ngfirewalls ---
func packNgfirewallsListFromSdk(ctx context.Context, sdks []cloud_ngfw.Ngfirewalls) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.Ngfirewalls")
	diags := diag.Diagnostics{}
	var data []models.Ngfirewalls

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.Ngfirewalls
		obj, d := packNgfirewallsFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.Ngfirewalls{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.Ngfirewalls", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.Ngfirewalls{}.AttrType(), data)
}

// --- Unpacker for NgfirewallResponseBody ---
func unpackNgfirewallResponseBodyToSdk(ctx context.Context, obj types.Object) (*cloud_ngfw.NgfirewallResponseBody, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.NgfirewallResponseBody", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.NgfirewallResponseBody
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk cloud_ngfw.NgfirewallResponseBody
	var d diag.Diagnostics
	// Handling Objects
	if !model.Firewall.IsNull() && !model.Firewall.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Firewall")
		unpacked, d := unpackNgfirewallFirewallToSdk(ctx, model.Firewall)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Firewall"})
		}
		if unpacked != nil {
			sdk.Firewall = unpacked
		}
	}

	// Handling Objects
	if !model.Status.IsNull() && !model.Status.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Status")
		unpacked, d := unpackNgfirewallStatusToSdk(ctx, model.Status)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Status"})
		}
		if unpacked != nil {
			sdk.Status = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.NgfirewallResponseBody", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for NgfirewallResponseBody ---
func packNgfirewallResponseBodyFromSdk(ctx context.Context, sdk cloud_ngfw.NgfirewallResponseBody) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.NgfirewallResponseBody", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.NgfirewallResponseBody
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Firewall != nil {
		tflog.Debug(ctx, "Packing nested object for field Firewall")
		packed, d := packNgfirewallFirewallFromSdk(ctx, *sdk.Firewall)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Firewall"})
		}
		model.Firewall = packed
	} else {
		model.Firewall = basetypes.NewObjectNull(models.NgfirewallFirewall{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Status != nil {
		tflog.Debug(ctx, "Packing nested object for field Status")
		packed, d := packNgfirewallStatusFromSdk(ctx, *sdk.Status)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Status"})
		}
		model.Status = packed
	} else {
		model.Status = basetypes.NewObjectNull(models.NgfirewallStatus{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.NgfirewallResponseBody{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.NgfirewallResponseBody", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for NgfirewallResponseBody ---
func unpackNgfirewallResponseBodyListToSdk(ctx context.Context, list types.List) ([]cloud_ngfw.NgfirewallResponseBody, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.NgfirewallResponseBody")
	diags := diag.Diagnostics{}
	var data []models.NgfirewallResponseBody
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]cloud_ngfw.NgfirewallResponseBody, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.NgfirewallResponseBody{}.AttrTypes(), &item)
		unpacked, d := unpackNgfirewallResponseBodyToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.NgfirewallResponseBody", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for NgfirewallResponseBody ---
func packNgfirewallResponseBodyListFromSdk(ctx context.Context, sdks []cloud_ngfw.NgfirewallResponseBody) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.NgfirewallResponseBody")
	diags := diag.Diagnostics{}
	var data []models.NgfirewallResponseBody

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.NgfirewallResponseBody
		obj, d := packNgfirewallResponseBodyFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.NgfirewallResponseBody{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.NgfirewallResponseBody", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.NgfirewallResponseBody{}.AttrType(), data)
}

// --- Unpacker for NgfirewallFirewall ---
func unpackNgfirewallFirewallToSdk(ctx context.Context, obj types.Object) (*cloud_ngfw.NgfirewallFirewall, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.NgfirewallFirewall", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.NgfirewallFirewall
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk cloud_ngfw.NgfirewallFirewall
	var d diag.Diagnostics
	// Handling Primitives
	if !model.AdvancedThreatLog.IsNull() && !model.AdvancedThreatLog.IsUnknown() {
		sdk.AdvancedThreatLog = model.AdvancedThreatLog.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "AdvancedThreatLog", "value": *sdk.AdvancedThreatLog})
	}

	// Handling Lists
	if !model.AllowListAccounts.IsNull() && !model.AllowListAccounts.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field AllowListAccounts")
		diags.Append(model.AllowListAccounts.ElementsAs(ctx, &sdk.AllowListAccounts, false)...)
	}

	// Handling Lists
	if !model.ChangeProtection.IsNull() && !model.ChangeProtection.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field ChangeProtection")
		diags.Append(model.ChangeProtection.ElementsAs(ctx, &sdk.ChangeProtection, false)...)
	}

	// Handling Objects
	if !model.CloudwatchMetrics.IsNull() && !model.CloudwatchMetrics.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field CloudwatchMetrics")
		unpacked, d := unpackNgfirewallCloudwatchMetricsToSdk(ctx, model.CloudwatchMetrics)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "CloudwatchMetrics"})
		}
		if unpacked != nil {
			sdk.CloudwatchMetrics = unpacked
		}
	}

	// Handling Lists
	if !model.CustomerZoneIdList.IsNull() && !model.CustomerZoneIdList.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field CustomerZoneIdList")
		diags.Append(model.CustomerZoneIdList.ElementsAs(ctx, &sdk.CustomerZoneIdList, false)...)
	}

	// Handling Primitives
	if !model.DeploymentUpdateToken.IsNull() && !model.DeploymentUpdateToken.IsUnknown() {
		sdk.DeploymentUpdateToken = model.DeploymentUpdateToken.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DeploymentUpdateToken", "value": *sdk.DeploymentUpdateToken})
	}

	// Handling Primitives
	if !model.Description.IsNull() && !model.Description.IsUnknown() {
		sdk.Description = model.Description.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Description", "value": *sdk.Description})
	}

	// Handling Primitives
	if !model.DeviceRuleStackCommitStatus.IsNull() && !model.DeviceRuleStackCommitStatus.IsUnknown() {
		sdk.DeviceRuleStackCommitStatus = model.DeviceRuleStackCommitStatus.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DeviceRuleStackCommitStatus", "value": *sdk.DeviceRuleStackCommitStatus})
	}

	// Handling Objects
	if !model.EgressNAT.IsNull() && !model.EgressNAT.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field EgressNAT")
		unpacked, d := unpackNgfirewallEgressNatConfigToSdk(ctx, model.EgressNAT)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "EgressNAT"})
		}
		if unpacked != nil {
			sdk.EgressNAT = unpacked
		}
	}

	// Handling Primitives
	if !model.EndpointServiceName.IsNull() && !model.EndpointServiceName.IsUnknown() {
		sdk.EndpointServiceName = model.EndpointServiceName.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "EndpointServiceName", "value": *sdk.EndpointServiceName})
	}

	// Handling Lists
	if !model.Endpoints.IsNull() && !model.Endpoints.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field Endpoints")
		unpacked, d := unpackNgfirewallEndpointListToSdk(ctx, model.Endpoints)
		diags.Append(d...)
		sdk.Endpoints = unpacked
	}

	// Handling Primitives
	if !model.FirewallId.IsNull() && !model.FirewallId.IsUnknown() {
		sdk.FirewallId = model.FirewallId.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "FirewallId", "value": *sdk.FirewallId})
	}

	// Handling Objects
	if !model.GWLB.IsNull() && !model.GWLB.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field GWLB")
		unpacked, d := unpackNgfirewallGwlbConfigToSdk(ctx, model.GWLB)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "GWLB"})
		}
		if unpacked != nil {
			sdk.GWLB = unpacked
		}
	}

	// Handling Objects
	if !model.Ipv6.IsNull() && !model.Ipv6.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Ipv6")
		unpacked, d := unpackNgfirewallIpv6ConfigToSdk(ctx, model.Ipv6)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Ipv6"})
		}
		if unpacked != nil {
			sdk.Ipv6 = unpacked
		}
	}

	// Handling Primitives
	if !model.LinkId.IsNull() && !model.LinkId.IsUnknown() {
		sdk.LinkId = model.LinkId.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "LinkId", "value": *sdk.LinkId})
	}

	// Handling Primitives
	if !model.LinkStatus.IsNull() && !model.LinkStatus.IsUnknown() {
		sdk.LinkStatus = model.LinkStatus.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "LinkStatus", "value": *sdk.LinkStatus})
	}

	// Handling Objects
	if !model.LogConfig.IsNull() && !model.LogConfig.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field LogConfig")
		unpacked, d := unpackNgfirewallLogConfigToSdk(ctx, model.LogConfig)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "LogConfig"})
		}
		if unpacked != nil {
			sdk.LogConfig = unpacked
		}
	}

	// Handling Lists
	if !model.Notifications.IsNull() && !model.Notifications.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field Notifications")
		unpacked, d := unpackNgfirewallNotificationListToSdk(ctx, model.Notifications)
		diags.Append(d...)
		sdk.Notifications = unpacked
	}

	// Handling Objects
	if !model.PrivateAccess.IsNull() && !model.PrivateAccess.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field PrivateAccess")
		unpacked, d := unpackNgfirewallPrivateAccessToSdk(ctx, model.PrivateAccess)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "PrivateAccess"})
		}
		if unpacked != nil {
			sdk.PrivateAccess = unpacked
		}
	}

	// Handling Primitives
	if !model.Region.IsNull() && !model.Region.IsUnknown() {
		sdk.Region = model.Region.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Region", "value": *sdk.Region})
	}

	// Handling Primitives
	if !model.SoftwareVersion.IsNull() && !model.SoftwareVersion.IsUnknown() {
		sdk.SoftwareVersion = model.SoftwareVersion.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "SoftwareVersion", "value": *sdk.SoftwareVersion})
	}

	// Handling Lists
	if !model.Tags.IsNull() && !model.Tags.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field Tags")
		unpacked, d := unpackNgfirewallTagListToSdk(ctx, model.Tags)
		diags.Append(d...)
		sdk.Tags = unpacked
	}

	// Handling Primitives
	if !model.Tier.IsNull() && !model.Tier.IsUnknown() {
		sdk.Tier = model.Tier.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Tier", "value": *sdk.Tier})
	}

	// Handling Primitives
	if !model.UpdateToken.IsNull() && !model.UpdateToken.IsUnknown() {
		sdk.UpdateToken = model.UpdateToken.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "UpdateToken", "value": *sdk.UpdateToken})
	}

	// Handling Objects
	if !model.UserID.IsNull() && !model.UserID.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field UserID")
		unpacked, d := unpackNgfirewallUserIdToSdk(ctx, model.UserID)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "UserID"})
		}
		if unpacked != nil {
			sdk.UserID = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.NgfirewallFirewall", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for NgfirewallFirewall ---
func packNgfirewallFirewallFromSdk(ctx context.Context, sdk cloud_ngfw.NgfirewallFirewall) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.NgfirewallFirewall", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.NgfirewallFirewall
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.AdvancedThreatLog != nil {
		model.AdvancedThreatLog = basetypes.NewBoolValue(*sdk.AdvancedThreatLog)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "AdvancedThreatLog", "value": *sdk.AdvancedThreatLog})
	} else {
		model.AdvancedThreatLog = basetypes.NewBoolNull()
	}
	// Handling Lists
	if sdk.AllowListAccounts != nil {
		tflog.Debug(ctx, "Packing list of primitives for field AllowListAccounts")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.AllowListAccounts, d = basetypes.NewListValueFrom(ctx, elemType, sdk.AllowListAccounts)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.AllowListAccounts = basetypes.NewListNull(elemType)
	}
	// Handling Lists
	if sdk.ChangeProtection != nil {
		tflog.Debug(ctx, "Packing list of primitives for field ChangeProtection")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.ChangeProtection, d = basetypes.NewListValueFrom(ctx, elemType, sdk.ChangeProtection)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.ChangeProtection = basetypes.NewListNull(elemType)
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.CloudwatchMetrics != nil {
		tflog.Debug(ctx, "Packing nested object for field CloudwatchMetrics")
		packed, d := packNgfirewallCloudwatchMetricsFromSdk(ctx, *sdk.CloudwatchMetrics)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "CloudwatchMetrics"})
		}
		model.CloudwatchMetrics = packed
	} else {
		model.CloudwatchMetrics = basetypes.NewObjectNull(models.NgfirewallCloudwatchMetrics{}.AttrTypes())
	}
	// Handling Lists
	if sdk.CustomerZoneIdList != nil {
		tflog.Debug(ctx, "Packing list of primitives for field CustomerZoneIdList")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.CustomerZoneIdList, d = basetypes.NewListValueFrom(ctx, elemType, sdk.CustomerZoneIdList)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.CustomerZoneIdList = basetypes.NewListNull(elemType)
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.DeploymentUpdateToken != nil {
		model.DeploymentUpdateToken = basetypes.NewStringValue(*sdk.DeploymentUpdateToken)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DeploymentUpdateToken", "value": *sdk.DeploymentUpdateToken})
	} else {
		model.DeploymentUpdateToken = basetypes.NewStringNull()
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
	if sdk.DeviceRuleStackCommitStatus != nil {
		model.DeviceRuleStackCommitStatus = basetypes.NewStringValue(*sdk.DeviceRuleStackCommitStatus)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DeviceRuleStackCommitStatus", "value": *sdk.DeviceRuleStackCommitStatus})
	} else {
		model.DeviceRuleStackCommitStatus = basetypes.NewStringNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.EgressNAT != nil {
		tflog.Debug(ctx, "Packing nested object for field EgressNAT")
		packed, d := packNgfirewallEgressNatConfigFromSdk(ctx, *sdk.EgressNAT)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "EgressNAT"})
		}
		model.EgressNAT = packed
	} else {
		model.EgressNAT = basetypes.NewObjectNull(models.NgfirewallEgressNatConfig{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.EndpointServiceName != nil {
		model.EndpointServiceName = basetypes.NewStringValue(*sdk.EndpointServiceName)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "EndpointServiceName", "value": *sdk.EndpointServiceName})
	} else {
		model.EndpointServiceName = basetypes.NewStringNull()
	}
	// Handling Lists
	if sdk.Endpoints != nil {
		tflog.Debug(ctx, "Packing list of objects for field Endpoints")
		packed, d := packNgfirewallEndpointListFromSdk(ctx, sdk.Endpoints)
		diags.Append(d...)
		model.Endpoints = packed
	} else {
		model.Endpoints = basetypes.NewListNull(models.NgfirewallEndpoint{}.AttrType())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.FirewallId != nil {
		model.FirewallId = basetypes.NewStringValue(*sdk.FirewallId)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "FirewallId", "value": *sdk.FirewallId})
	} else {
		model.FirewallId = basetypes.NewStringNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.GWLB != nil {
		tflog.Debug(ctx, "Packing nested object for field GWLB")
		packed, d := packNgfirewallGwlbConfigFromSdk(ctx, *sdk.GWLB)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "GWLB"})
		}
		model.GWLB = packed
	} else {
		model.GWLB = basetypes.NewObjectNull(models.NgfirewallGwlbConfig{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Ipv6 != nil {
		tflog.Debug(ctx, "Packing nested object for field Ipv6")
		packed, d := packNgfirewallIpv6ConfigFromSdk(ctx, *sdk.Ipv6)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Ipv6"})
		}
		model.Ipv6 = packed
	} else {
		model.Ipv6 = basetypes.NewObjectNull(models.NgfirewallIpv6Config{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.LinkId != nil {
		model.LinkId = basetypes.NewStringValue(*sdk.LinkId)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "LinkId", "value": *sdk.LinkId})
	} else {
		model.LinkId = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.LinkStatus != nil {
		model.LinkStatus = basetypes.NewStringValue(*sdk.LinkStatus)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "LinkStatus", "value": *sdk.LinkStatus})
	} else {
		model.LinkStatus = basetypes.NewStringNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.LogConfig != nil {
		tflog.Debug(ctx, "Packing nested object for field LogConfig")
		packed, d := packNgfirewallLogConfigFromSdk(ctx, *sdk.LogConfig)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "LogConfig"})
		}
		model.LogConfig = packed
	} else {
		model.LogConfig = basetypes.NewObjectNull(models.NgfirewallLogConfig{}.AttrTypes())
	}
	// Handling Lists
	if sdk.Notifications != nil {
		tflog.Debug(ctx, "Packing list of objects for field Notifications")
		packed, d := packNgfirewallNotificationListFromSdk(ctx, sdk.Notifications)
		diags.Append(d...)
		model.Notifications = packed
	} else {
		model.Notifications = basetypes.NewListNull(models.NgfirewallNotification{}.AttrType())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.PrivateAccess != nil {
		tflog.Debug(ctx, "Packing nested object for field PrivateAccess")
		packed, d := packNgfirewallPrivateAccessFromSdk(ctx, *sdk.PrivateAccess)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "PrivateAccess"})
		}
		model.PrivateAccess = packed
	} else {
		model.PrivateAccess = basetypes.NewObjectNull(models.NgfirewallPrivateAccess{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Region != nil {
		model.Region = basetypes.NewStringValue(*sdk.Region)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Region", "value": *sdk.Region})
	} else {
		model.Region = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.SoftwareVersion != nil {
		model.SoftwareVersion = basetypes.NewStringValue(*sdk.SoftwareVersion)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "SoftwareVersion", "value": *sdk.SoftwareVersion})
	} else {
		model.SoftwareVersion = basetypes.NewStringNull()
	}
	// Handling Lists
	if sdk.Tags != nil {
		tflog.Debug(ctx, "Packing list of objects for field Tags")
		packed, d := packNgfirewallTagListFromSdk(ctx, sdk.Tags)
		diags.Append(d...)
		model.Tags = packed
	} else {
		model.Tags = basetypes.NewListNull(models.NgfirewallTag{}.AttrType())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Tier != nil {
		model.Tier = basetypes.NewStringValue(*sdk.Tier)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Tier", "value": *sdk.Tier})
	} else {
		model.Tier = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.UpdateToken != nil {
		model.UpdateToken = basetypes.NewStringValue(*sdk.UpdateToken)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "UpdateToken", "value": *sdk.UpdateToken})
	} else {
		model.UpdateToken = basetypes.NewStringNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.UserID != nil {
		tflog.Debug(ctx, "Packing nested object for field UserID")
		packed, d := packNgfirewallUserIdFromSdk(ctx, *sdk.UserID)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "UserID"})
		}
		model.UserID = packed
	} else {
		model.UserID = basetypes.NewObjectNull(models.NgfirewallUserId{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.NgfirewallFirewall{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.NgfirewallFirewall", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for NgfirewallFirewall ---
func unpackNgfirewallFirewallListToSdk(ctx context.Context, list types.List) ([]cloud_ngfw.NgfirewallFirewall, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.NgfirewallFirewall")
	diags := diag.Diagnostics{}
	var data []models.NgfirewallFirewall
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]cloud_ngfw.NgfirewallFirewall, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.NgfirewallFirewall{}.AttrTypes(), &item)
		unpacked, d := unpackNgfirewallFirewallToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.NgfirewallFirewall", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for NgfirewallFirewall ---
func packNgfirewallFirewallListFromSdk(ctx context.Context, sdks []cloud_ngfw.NgfirewallFirewall) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.NgfirewallFirewall")
	diags := diag.Diagnostics{}
	var data []models.NgfirewallFirewall

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.NgfirewallFirewall
		obj, d := packNgfirewallFirewallFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.NgfirewallFirewall{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.NgfirewallFirewall", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.NgfirewallFirewall{}.AttrType(), data)
}

// --- Unpacker for NgfirewallCloudwatchMetrics ---
func unpackNgfirewallCloudwatchMetricsToSdk(ctx context.Context, obj types.Object) (*cloud_ngfw.NgfirewallCloudwatchMetrics, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.NgfirewallCloudwatchMetrics", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.NgfirewallCloudwatchMetrics
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk cloud_ngfw.NgfirewallCloudwatchMetrics
	var d diag.Diagnostics
	// Handling Primitives
	if !model.AccountId.IsNull() && !model.AccountId.IsUnknown() {
		sdk.AccountId = model.AccountId.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "AccountId", "value": *sdk.AccountId})
	}

	// Handling Primitives
	if !model.CloudWatchMetricNamespace.IsNull() && !model.CloudWatchMetricNamespace.IsUnknown() {
		sdk.CloudWatchMetricNamespace = model.CloudWatchMetricNamespace.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "CloudWatchMetricNamespace", "value": *sdk.CloudWatchMetricNamespace})
	}

	// Handling Lists
	if !model.CloudWatchMetricsFields.IsNull() && !model.CloudWatchMetricsFields.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field CloudWatchMetricsFields")
		diags.Append(model.CloudWatchMetricsFields.ElementsAs(ctx, &sdk.CloudWatchMetricsFields, false)...)
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.NgfirewallCloudwatchMetrics", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for NgfirewallCloudwatchMetrics ---
func packNgfirewallCloudwatchMetricsFromSdk(ctx context.Context, sdk cloud_ngfw.NgfirewallCloudwatchMetrics) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.NgfirewallCloudwatchMetrics", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.NgfirewallCloudwatchMetrics
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.AccountId != nil {
		model.AccountId = basetypes.NewStringValue(*sdk.AccountId)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "AccountId", "value": *sdk.AccountId})
	} else {
		model.AccountId = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.CloudWatchMetricNamespace != nil {
		model.CloudWatchMetricNamespace = basetypes.NewStringValue(*sdk.CloudWatchMetricNamespace)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "CloudWatchMetricNamespace", "value": *sdk.CloudWatchMetricNamespace})
	} else {
		model.CloudWatchMetricNamespace = basetypes.NewStringNull()
	}
	// Handling Lists
	if sdk.CloudWatchMetricsFields != nil {
		tflog.Debug(ctx, "Packing list of primitives for field CloudWatchMetricsFields")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.CloudWatchMetricsFields, d = basetypes.NewListValueFrom(ctx, elemType, sdk.CloudWatchMetricsFields)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.CloudWatchMetricsFields = basetypes.NewListNull(elemType)
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.NgfirewallCloudwatchMetrics{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.NgfirewallCloudwatchMetrics", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for NgfirewallCloudwatchMetrics ---
func unpackNgfirewallCloudwatchMetricsListToSdk(ctx context.Context, list types.List) ([]cloud_ngfw.NgfirewallCloudwatchMetrics, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.NgfirewallCloudwatchMetrics")
	diags := diag.Diagnostics{}
	var data []models.NgfirewallCloudwatchMetrics
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]cloud_ngfw.NgfirewallCloudwatchMetrics, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.NgfirewallCloudwatchMetrics{}.AttrTypes(), &item)
		unpacked, d := unpackNgfirewallCloudwatchMetricsToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.NgfirewallCloudwatchMetrics", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for NgfirewallCloudwatchMetrics ---
func packNgfirewallCloudwatchMetricsListFromSdk(ctx context.Context, sdks []cloud_ngfw.NgfirewallCloudwatchMetrics) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.NgfirewallCloudwatchMetrics")
	diags := diag.Diagnostics{}
	var data []models.NgfirewallCloudwatchMetrics

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.NgfirewallCloudwatchMetrics
		obj, d := packNgfirewallCloudwatchMetricsFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.NgfirewallCloudwatchMetrics{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.NgfirewallCloudwatchMetrics", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.NgfirewallCloudwatchMetrics{}.AttrType(), data)
}

// --- Unpacker for NgfirewallEgressNatConfig ---
func unpackNgfirewallEgressNatConfigToSdk(ctx context.Context, obj types.Object) (*cloud_ngfw.NgfirewallEgressNatConfig, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.NgfirewallEgressNatConfig", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.NgfirewallEgressNatConfig
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk cloud_ngfw.NgfirewallEgressNatConfig
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Enabled.IsNull() && !model.Enabled.IsUnknown() {
		sdk.Enabled = model.Enabled.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Enabled", "value": *sdk.Enabled})
	}

	// Handling Objects
	if !model.Settings.IsNull() && !model.Settings.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Settings")
		unpacked, d := unpackNgfirewallEgressNatSettingsToSdk(ctx, model.Settings)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Settings"})
		}
		if unpacked != nil {
			sdk.Settings = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.NgfirewallEgressNatConfig", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for NgfirewallEgressNatConfig ---
func packNgfirewallEgressNatConfigFromSdk(ctx context.Context, sdk cloud_ngfw.NgfirewallEgressNatConfig) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.NgfirewallEgressNatConfig", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.NgfirewallEgressNatConfig
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Enabled != nil {
		model.Enabled = basetypes.NewBoolValue(*sdk.Enabled)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Enabled", "value": *sdk.Enabled})
	} else {
		model.Enabled = basetypes.NewBoolNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Settings != nil {
		tflog.Debug(ctx, "Packing nested object for field Settings")
		packed, d := packNgfirewallEgressNatSettingsFromSdk(ctx, *sdk.Settings)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Settings"})
		}
		model.Settings = packed
	} else {
		model.Settings = basetypes.NewObjectNull(models.NgfirewallEgressNatSettings{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.NgfirewallEgressNatConfig{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.NgfirewallEgressNatConfig", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for NgfirewallEgressNatConfig ---
func unpackNgfirewallEgressNatConfigListToSdk(ctx context.Context, list types.List) ([]cloud_ngfw.NgfirewallEgressNatConfig, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.NgfirewallEgressNatConfig")
	diags := diag.Diagnostics{}
	var data []models.NgfirewallEgressNatConfig
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]cloud_ngfw.NgfirewallEgressNatConfig, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.NgfirewallEgressNatConfig{}.AttrTypes(), &item)
		unpacked, d := unpackNgfirewallEgressNatConfigToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.NgfirewallEgressNatConfig", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for NgfirewallEgressNatConfig ---
func packNgfirewallEgressNatConfigListFromSdk(ctx context.Context, sdks []cloud_ngfw.NgfirewallEgressNatConfig) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.NgfirewallEgressNatConfig")
	diags := diag.Diagnostics{}
	var data []models.NgfirewallEgressNatConfig

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.NgfirewallEgressNatConfig
		obj, d := packNgfirewallEgressNatConfigFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.NgfirewallEgressNatConfig{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.NgfirewallEgressNatConfig", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.NgfirewallEgressNatConfig{}.AttrType(), data)
}

// --- Unpacker for NgfirewallEgressNatSettings ---
func unpackNgfirewallEgressNatSettingsToSdk(ctx context.Context, obj types.Object) (*cloud_ngfw.NgfirewallEgressNatSettings, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.NgfirewallEgressNatSettings", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.NgfirewallEgressNatSettings
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk cloud_ngfw.NgfirewallEgressNatSettings
	var d diag.Diagnostics
	// Handling Primitives
	if !model.IPAMPoolId.IsNull() && !model.IPAMPoolId.IsUnknown() {
		sdk.IPAMPoolId = model.IPAMPoolId.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "IPAMPoolId", "value": *sdk.IPAMPoolId})
	}

	// Handling Primitives
	if !model.IPPoolType.IsNull() && !model.IPPoolType.IsUnknown() {
		sdk.IPPoolType = model.IPPoolType.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "IPPoolType", "value": *sdk.IPPoolType})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.NgfirewallEgressNatSettings", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for NgfirewallEgressNatSettings ---
func packNgfirewallEgressNatSettingsFromSdk(ctx context.Context, sdk cloud_ngfw.NgfirewallEgressNatSettings) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.NgfirewallEgressNatSettings", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.NgfirewallEgressNatSettings
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.IPAMPoolId != nil {
		model.IPAMPoolId = basetypes.NewStringValue(*sdk.IPAMPoolId)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "IPAMPoolId", "value": *sdk.IPAMPoolId})
	} else {
		model.IPAMPoolId = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.IPPoolType != nil {
		model.IPPoolType = basetypes.NewStringValue(*sdk.IPPoolType)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "IPPoolType", "value": *sdk.IPPoolType})
	} else {
		model.IPPoolType = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.NgfirewallEgressNatSettings{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.NgfirewallEgressNatSettings", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for NgfirewallEgressNatSettings ---
func unpackNgfirewallEgressNatSettingsListToSdk(ctx context.Context, list types.List) ([]cloud_ngfw.NgfirewallEgressNatSettings, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.NgfirewallEgressNatSettings")
	diags := diag.Diagnostics{}
	var data []models.NgfirewallEgressNatSettings
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]cloud_ngfw.NgfirewallEgressNatSettings, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.NgfirewallEgressNatSettings{}.AttrTypes(), &item)
		unpacked, d := unpackNgfirewallEgressNatSettingsToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.NgfirewallEgressNatSettings", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for NgfirewallEgressNatSettings ---
func packNgfirewallEgressNatSettingsListFromSdk(ctx context.Context, sdks []cloud_ngfw.NgfirewallEgressNatSettings) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.NgfirewallEgressNatSettings")
	diags := diag.Diagnostics{}
	var data []models.NgfirewallEgressNatSettings

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.NgfirewallEgressNatSettings
		obj, d := packNgfirewallEgressNatSettingsFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.NgfirewallEgressNatSettings{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.NgfirewallEgressNatSettings", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.NgfirewallEgressNatSettings{}.AttrType(), data)
}

// --- Unpacker for NgfirewallEndpoint ---
func unpackNgfirewallEndpointToSdk(ctx context.Context, obj types.Object) (*cloud_ngfw.NgfirewallEndpoint, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.NgfirewallEndpoint", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.NgfirewallEndpoint
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk cloud_ngfw.NgfirewallEndpoint
	var d diag.Diagnostics
	// Handling Primitives
	if !model.AccountId.IsNull() && !model.AccountId.IsUnknown() {
		sdk.AccountId = model.AccountId.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "AccountId", "value": *sdk.AccountId})
	}

	// Handling Primitives
	if !model.EgressNATEnabled.IsNull() && !model.EgressNATEnabled.IsUnknown() {
		sdk.EgressNATEnabled = model.EgressNATEnabled.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "EgressNATEnabled", "value": *sdk.EgressNATEnabled})
	}

	// Handling Primitives
	if !model.EndpointId.IsNull() && !model.EndpointId.IsUnknown() {
		sdk.EndpointId = model.EndpointId.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "EndpointId", "value": *sdk.EndpointId})
	}

	// Handling Primitives
	if !model.Mode.IsNull() && !model.Mode.IsUnknown() {
		sdk.Mode = model.Mode.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Mode", "value": *sdk.Mode})
	}

	// Handling Objects
	if !model.Prefixes.IsNull() && !model.Prefixes.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Prefixes")
		unpacked, d := unpackNgfirewallEndpointPrefixesToSdk(ctx, model.Prefixes)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Prefixes"})
		}
		if unpacked != nil {
			sdk.Prefixes = unpacked
		}
	}

	// Handling Primitives
	if !model.RejectedReason.IsNull() && !model.RejectedReason.IsUnknown() {
		sdk.RejectedReason = model.RejectedReason.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "RejectedReason", "value": *sdk.RejectedReason})
	}

	// Handling Primitives
	if !model.Status.IsNull() && !model.Status.IsUnknown() {
		sdk.Status = model.Status.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Status", "value": *sdk.Status})
	}

	// Handling Primitives
	if !model.SubnetId.IsNull() && !model.SubnetId.IsUnknown() {
		sdk.SubnetId = model.SubnetId.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "SubnetId", "value": *sdk.SubnetId})
	}

	// Handling Primitives
	if !model.VpcId.IsNull() && !model.VpcId.IsUnknown() {
		sdk.VpcId = model.VpcId.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "VpcId", "value": *sdk.VpcId})
	}

	// Handling Primitives
	if !model.ZoneId.IsNull() && !model.ZoneId.IsUnknown() {
		sdk.ZoneId = model.ZoneId.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "ZoneId", "value": *sdk.ZoneId})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.NgfirewallEndpoint", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for NgfirewallEndpoint ---
func packNgfirewallEndpointFromSdk(ctx context.Context, sdk cloud_ngfw.NgfirewallEndpoint) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.NgfirewallEndpoint", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.NgfirewallEndpoint
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.AccountId != nil {
		model.AccountId = basetypes.NewStringValue(*sdk.AccountId)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "AccountId", "value": *sdk.AccountId})
	} else {
		model.AccountId = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.EgressNATEnabled != nil {
		model.EgressNATEnabled = basetypes.NewBoolValue(*sdk.EgressNATEnabled)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "EgressNATEnabled", "value": *sdk.EgressNATEnabled})
	} else {
		model.EgressNATEnabled = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.EndpointId != nil {
		model.EndpointId = basetypes.NewStringValue(*sdk.EndpointId)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "EndpointId", "value": *sdk.EndpointId})
	} else {
		model.EndpointId = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Mode != nil {
		model.Mode = basetypes.NewStringValue(*sdk.Mode)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Mode", "value": *sdk.Mode})
	} else {
		model.Mode = basetypes.NewStringNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Prefixes != nil {
		tflog.Debug(ctx, "Packing nested object for field Prefixes")
		packed, d := packNgfirewallEndpointPrefixesFromSdk(ctx, *sdk.Prefixes)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Prefixes"})
		}
		model.Prefixes = packed
	} else {
		model.Prefixes = basetypes.NewObjectNull(models.NgfirewallEndpointPrefixes{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.RejectedReason != nil {
		model.RejectedReason = basetypes.NewStringValue(*sdk.RejectedReason)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "RejectedReason", "value": *sdk.RejectedReason})
	} else {
		model.RejectedReason = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Status != nil {
		model.Status = basetypes.NewStringValue(*sdk.Status)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Status", "value": *sdk.Status})
	} else {
		model.Status = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.SubnetId != nil {
		model.SubnetId = basetypes.NewStringValue(*sdk.SubnetId)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "SubnetId", "value": *sdk.SubnetId})
	} else {
		model.SubnetId = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.VpcId != nil {
		model.VpcId = basetypes.NewStringValue(*sdk.VpcId)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "VpcId", "value": *sdk.VpcId})
	} else {
		model.VpcId = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.ZoneId != nil {
		model.ZoneId = basetypes.NewStringValue(*sdk.ZoneId)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "ZoneId", "value": *sdk.ZoneId})
	} else {
		model.ZoneId = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.NgfirewallEndpoint{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.NgfirewallEndpoint", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for NgfirewallEndpoint ---
func unpackNgfirewallEndpointListToSdk(ctx context.Context, list types.List) ([]cloud_ngfw.NgfirewallEndpoint, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.NgfirewallEndpoint")
	diags := diag.Diagnostics{}
	var data []models.NgfirewallEndpoint
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]cloud_ngfw.NgfirewallEndpoint, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.NgfirewallEndpoint{}.AttrTypes(), &item)
		unpacked, d := unpackNgfirewallEndpointToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.NgfirewallEndpoint", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for NgfirewallEndpoint ---
func packNgfirewallEndpointListFromSdk(ctx context.Context, sdks []cloud_ngfw.NgfirewallEndpoint) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.NgfirewallEndpoint")
	diags := diag.Diagnostics{}
	var data []models.NgfirewallEndpoint

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.NgfirewallEndpoint
		obj, d := packNgfirewallEndpointFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.NgfirewallEndpoint{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.NgfirewallEndpoint", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.NgfirewallEndpoint{}.AttrType(), data)
}

// --- Unpacker for NgfirewallEndpointPrefixes ---
func unpackNgfirewallEndpointPrefixesToSdk(ctx context.Context, obj types.Object) (*cloud_ngfw.NgfirewallEndpointPrefixes, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.NgfirewallEndpointPrefixes", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.NgfirewallEndpointPrefixes
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk cloud_ngfw.NgfirewallEndpointPrefixes
	var d diag.Diagnostics
	// Handling Objects
	if !model.PrivatePrefix.IsNull() && !model.PrivatePrefix.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field PrivatePrefix")
		unpacked, d := unpackNgfirewallPrefixListToSdk(ctx, model.PrivatePrefix)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "PrivatePrefix"})
		}
		if unpacked != nil {
			sdk.PrivatePrefix = unpacked
		}
	}

	// Handling Objects
	if !model.PublicPrefix.IsNull() && !model.PublicPrefix.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field PublicPrefix")
		unpacked, d := unpackNgfirewallPrefixListToSdk(ctx, model.PublicPrefix)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "PublicPrefix"})
		}
		if unpacked != nil {
			sdk.PublicPrefix = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.NgfirewallEndpointPrefixes", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for NgfirewallEndpointPrefixes ---
func packNgfirewallEndpointPrefixesFromSdk(ctx context.Context, sdk cloud_ngfw.NgfirewallEndpointPrefixes) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.NgfirewallEndpointPrefixes", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.NgfirewallEndpointPrefixes
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.PrivatePrefix != nil {
		tflog.Debug(ctx, "Packing nested object for field PrivatePrefix")
		packed, d := packNgfirewallPrefixListFromSdk(ctx, *sdk.PrivatePrefix)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "PrivatePrefix"})
		}
		model.PrivatePrefix = packed
	} else {
		model.PrivatePrefix = basetypes.NewObjectNull(models.NgfirewallPrefixList{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.PublicPrefix != nil {
		tflog.Debug(ctx, "Packing nested object for field PublicPrefix")
		packed, d := packNgfirewallPrefixListFromSdk(ctx, *sdk.PublicPrefix)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "PublicPrefix"})
		}
		model.PublicPrefix = packed
	} else {
		model.PublicPrefix = basetypes.NewObjectNull(models.NgfirewallPrefixList{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.NgfirewallEndpointPrefixes{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.NgfirewallEndpointPrefixes", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for NgfirewallEndpointPrefixes ---
func unpackNgfirewallEndpointPrefixesListToSdk(ctx context.Context, list types.List) ([]cloud_ngfw.NgfirewallEndpointPrefixes, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.NgfirewallEndpointPrefixes")
	diags := diag.Diagnostics{}
	var data []models.NgfirewallEndpointPrefixes
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]cloud_ngfw.NgfirewallEndpointPrefixes, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.NgfirewallEndpointPrefixes{}.AttrTypes(), &item)
		unpacked, d := unpackNgfirewallEndpointPrefixesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.NgfirewallEndpointPrefixes", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for NgfirewallEndpointPrefixes ---
func packNgfirewallEndpointPrefixesListFromSdk(ctx context.Context, sdks []cloud_ngfw.NgfirewallEndpointPrefixes) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.NgfirewallEndpointPrefixes")
	diags := diag.Diagnostics{}
	var data []models.NgfirewallEndpointPrefixes

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.NgfirewallEndpointPrefixes
		obj, d := packNgfirewallEndpointPrefixesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.NgfirewallEndpointPrefixes{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.NgfirewallEndpointPrefixes", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.NgfirewallEndpointPrefixes{}.AttrType(), data)
}

// --- Unpacker for NgfirewallPrefixList ---
func unpackNgfirewallPrefixListToSdk(ctx context.Context, obj types.Object) (*cloud_ngfw.NgfirewallPrefixList, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.NgfirewallPrefixList", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.NgfirewallPrefixList
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk cloud_ngfw.NgfirewallPrefixList
	var d diag.Diagnostics
	// Handling Lists
	if !model.Cidrs.IsNull() && !model.Cidrs.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Cidrs")
		diags.Append(model.Cidrs.ElementsAs(ctx, &sdk.Cidrs, false)...)
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.NgfirewallPrefixList", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for NgfirewallPrefixList ---
func packNgfirewallPrefixListFromSdk(ctx context.Context, sdk cloud_ngfw.NgfirewallPrefixList) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.NgfirewallPrefixList", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.NgfirewallPrefixList
	var d diag.Diagnostics
	// Handling Lists
	if sdk.Cidrs != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Cidrs")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Cidrs, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Cidrs)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Cidrs = basetypes.NewListNull(elemType)
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.NgfirewallPrefixList{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.NgfirewallPrefixList", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for NgfirewallPrefixList ---
func unpackNgfirewallPrefixListListToSdk(ctx context.Context, list types.List) ([]cloud_ngfw.NgfirewallPrefixList, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.NgfirewallPrefixList")
	diags := diag.Diagnostics{}
	var data []models.NgfirewallPrefixList
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]cloud_ngfw.NgfirewallPrefixList, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.NgfirewallPrefixList{}.AttrTypes(), &item)
		unpacked, d := unpackNgfirewallPrefixListToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.NgfirewallPrefixList", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for NgfirewallPrefixList ---
func packNgfirewallPrefixListListFromSdk(ctx context.Context, sdks []cloud_ngfw.NgfirewallPrefixList) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.NgfirewallPrefixList")
	diags := diag.Diagnostics{}
	var data []models.NgfirewallPrefixList

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.NgfirewallPrefixList
		obj, d := packNgfirewallPrefixListFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.NgfirewallPrefixList{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.NgfirewallPrefixList", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.NgfirewallPrefixList{}.AttrType(), data)
}

// --- Unpacker for NgfirewallGwlbConfig ---
func unpackNgfirewallGwlbConfigToSdk(ctx context.Context, obj types.Object) (*cloud_ngfw.NgfirewallGwlbConfig, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.NgfirewallGwlbConfig", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.NgfirewallGwlbConfig
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk cloud_ngfw.NgfirewallGwlbConfig
	var d diag.Diagnostics
	// Handling Primitives
	if !model.DeregistrationDelay.IsNull() && !model.DeregistrationDelay.IsUnknown() {
		val := int32(model.DeregistrationDelay.ValueInt64())
		sdk.DeregistrationDelay = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DeregistrationDelay", "value": *sdk.DeregistrationDelay})
	}

	// Handling Primitives
	if !model.RejectedFlowCountAlertThreshold.IsNull() && !model.RejectedFlowCountAlertThreshold.IsUnknown() {
		val := int32(model.RejectedFlowCountAlertThreshold.ValueInt64())
		sdk.RejectedFlowCountAlertThreshold = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "RejectedFlowCountAlertThreshold", "value": *sdk.RejectedFlowCountAlertThreshold})
	}

	// Handling Primitives
	if !model.RejectedFlowCountTCPAlertThreshold.IsNull() && !model.RejectedFlowCountTCPAlertThreshold.IsUnknown() {
		val := int32(model.RejectedFlowCountTCPAlertThreshold.ValueInt64())
		sdk.RejectedFlowCountTCPAlertThreshold = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "RejectedFlowCountTCPAlertThreshold", "value": *sdk.RejectedFlowCountTCPAlertThreshold})
	}

	// Handling Primitives
	if !model.SessionRebalanceEnabled.IsNull() && !model.SessionRebalanceEnabled.IsUnknown() {
		sdk.SessionRebalanceEnabled = model.SessionRebalanceEnabled.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "SessionRebalanceEnabled", "value": *sdk.SessionRebalanceEnabled})
	}

	// Handling Primitives
	if !model.TCPIdleTimeout.IsNull() && !model.TCPIdleTimeout.IsUnknown() {
		val := int32(model.TCPIdleTimeout.ValueInt64())
		sdk.TCPIdleTimeout = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "TCPIdleTimeout", "value": *sdk.TCPIdleTimeout})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.NgfirewallGwlbConfig", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for NgfirewallGwlbConfig ---
func packNgfirewallGwlbConfigFromSdk(ctx context.Context, sdk cloud_ngfw.NgfirewallGwlbConfig) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.NgfirewallGwlbConfig", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.NgfirewallGwlbConfig
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.DeregistrationDelay != nil {
		model.DeregistrationDelay = basetypes.NewInt64Value(int64(*sdk.DeregistrationDelay))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DeregistrationDelay", "value": *sdk.DeregistrationDelay})
	} else {
		model.DeregistrationDelay = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.RejectedFlowCountAlertThreshold != nil {
		model.RejectedFlowCountAlertThreshold = basetypes.NewInt64Value(int64(*sdk.RejectedFlowCountAlertThreshold))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "RejectedFlowCountAlertThreshold", "value": *sdk.RejectedFlowCountAlertThreshold})
	} else {
		model.RejectedFlowCountAlertThreshold = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.RejectedFlowCountTCPAlertThreshold != nil {
		model.RejectedFlowCountTCPAlertThreshold = basetypes.NewInt64Value(int64(*sdk.RejectedFlowCountTCPAlertThreshold))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "RejectedFlowCountTCPAlertThreshold", "value": *sdk.RejectedFlowCountTCPAlertThreshold})
	} else {
		model.RejectedFlowCountTCPAlertThreshold = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.SessionRebalanceEnabled != nil {
		model.SessionRebalanceEnabled = basetypes.NewBoolValue(*sdk.SessionRebalanceEnabled)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "SessionRebalanceEnabled", "value": *sdk.SessionRebalanceEnabled})
	} else {
		model.SessionRebalanceEnabled = basetypes.NewBoolNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.TCPIdleTimeout != nil {
		model.TCPIdleTimeout = basetypes.NewInt64Value(int64(*sdk.TCPIdleTimeout))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "TCPIdleTimeout", "value": *sdk.TCPIdleTimeout})
	} else {
		model.TCPIdleTimeout = basetypes.NewInt64Null()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.NgfirewallGwlbConfig{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.NgfirewallGwlbConfig", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for NgfirewallGwlbConfig ---
func unpackNgfirewallGwlbConfigListToSdk(ctx context.Context, list types.List) ([]cloud_ngfw.NgfirewallGwlbConfig, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.NgfirewallGwlbConfig")
	diags := diag.Diagnostics{}
	var data []models.NgfirewallGwlbConfig
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]cloud_ngfw.NgfirewallGwlbConfig, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.NgfirewallGwlbConfig{}.AttrTypes(), &item)
		unpacked, d := unpackNgfirewallGwlbConfigToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.NgfirewallGwlbConfig", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for NgfirewallGwlbConfig ---
func packNgfirewallGwlbConfigListFromSdk(ctx context.Context, sdks []cloud_ngfw.NgfirewallGwlbConfig) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.NgfirewallGwlbConfig")
	diags := diag.Diagnostics{}
	var data []models.NgfirewallGwlbConfig

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.NgfirewallGwlbConfig
		obj, d := packNgfirewallGwlbConfigFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.NgfirewallGwlbConfig{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.NgfirewallGwlbConfig", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.NgfirewallGwlbConfig{}.AttrType(), data)
}

// --- Unpacker for NgfirewallIpv6Config ---
func unpackNgfirewallIpv6ConfigToSdk(ctx context.Context, obj types.Object) (*cloud_ngfw.NgfirewallIpv6Config, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.NgfirewallIpv6Config", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.NgfirewallIpv6Config
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk cloud_ngfw.NgfirewallIpv6Config
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Enabled.IsNull() && !model.Enabled.IsUnknown() {
		sdk.Enabled = model.Enabled.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Enabled", "value": *sdk.Enabled})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.NgfirewallIpv6Config", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for NgfirewallIpv6Config ---
func packNgfirewallIpv6ConfigFromSdk(ctx context.Context, sdk cloud_ngfw.NgfirewallIpv6Config) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.NgfirewallIpv6Config", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.NgfirewallIpv6Config
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Enabled != nil {
		model.Enabled = basetypes.NewBoolValue(*sdk.Enabled)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Enabled", "value": *sdk.Enabled})
	} else {
		model.Enabled = basetypes.NewBoolNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.NgfirewallIpv6Config{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.NgfirewallIpv6Config", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for NgfirewallIpv6Config ---
func unpackNgfirewallIpv6ConfigListToSdk(ctx context.Context, list types.List) ([]cloud_ngfw.NgfirewallIpv6Config, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.NgfirewallIpv6Config")
	diags := diag.Diagnostics{}
	var data []models.NgfirewallIpv6Config
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]cloud_ngfw.NgfirewallIpv6Config, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.NgfirewallIpv6Config{}.AttrTypes(), &item)
		unpacked, d := unpackNgfirewallIpv6ConfigToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.NgfirewallIpv6Config", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for NgfirewallIpv6Config ---
func packNgfirewallIpv6ConfigListFromSdk(ctx context.Context, sdks []cloud_ngfw.NgfirewallIpv6Config) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.NgfirewallIpv6Config")
	diags := diag.Diagnostics{}
	var data []models.NgfirewallIpv6Config

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.NgfirewallIpv6Config
		obj, d := packNgfirewallIpv6ConfigFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.NgfirewallIpv6Config{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.NgfirewallIpv6Config", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.NgfirewallIpv6Config{}.AttrType(), data)
}

// --- Unpacker for NgfirewallLogConfig ---
func unpackNgfirewallLogConfigToSdk(ctx context.Context, obj types.Object) (*cloud_ngfw.NgfirewallLogConfig, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.NgfirewallLogConfig", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.NgfirewallLogConfig
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk cloud_ngfw.NgfirewallLogConfig
	var d diag.Diagnostics
	// Handling Primitives
	if !model.AccountId.IsNull() && !model.AccountId.IsUnknown() {
		sdk.AccountId = model.AccountId.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "AccountId", "value": *sdk.AccountId})
	}

	// Handling Primitives
	if !model.LogDestination.IsNull() && !model.LogDestination.IsUnknown() {
		sdk.LogDestination = model.LogDestination.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "LogDestination", "value": sdk.LogDestination})
	}

	// Handling Primitives
	if !model.LogDestinationRegion.IsNull() && !model.LogDestinationRegion.IsUnknown() {
		sdk.LogDestinationRegion = model.LogDestinationRegion.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "LogDestinationRegion", "value": *sdk.LogDestinationRegion})
	}

	// Handling Primitives
	if !model.LogDestinationType.IsNull() && !model.LogDestinationType.IsUnknown() {
		sdk.LogDestinationType = model.LogDestinationType.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "LogDestinationType", "value": sdk.LogDestinationType})
	}

	// Handling Lists
	if !model.LogType.IsNull() && !model.LogType.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field LogType")
		diags.Append(model.LogType.ElementsAs(ctx, &sdk.LogType, false)...)
	}

	// Handling Primitives
	if !model.RoleType.IsNull() && !model.RoleType.IsUnknown() {
		sdk.RoleType = model.RoleType.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "RoleType", "value": sdk.RoleType})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.NgfirewallLogConfig", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for NgfirewallLogConfig ---
func packNgfirewallLogConfigFromSdk(ctx context.Context, sdk cloud_ngfw.NgfirewallLogConfig) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.NgfirewallLogConfig", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.NgfirewallLogConfig
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.AccountId != nil {
		model.AccountId = basetypes.NewStringValue(*sdk.AccountId)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "AccountId", "value": *sdk.AccountId})
	} else {
		model.AccountId = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	model.LogDestination = basetypes.NewStringValue(sdk.LogDestination)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "LogDestination", "value": sdk.LogDestination})
	// Handling Primitives
	// Standard primitive packing
	if sdk.LogDestinationRegion != nil {
		model.LogDestinationRegion = basetypes.NewStringValue(*sdk.LogDestinationRegion)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "LogDestinationRegion", "value": *sdk.LogDestinationRegion})
	} else {
		model.LogDestinationRegion = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	model.LogDestinationType = basetypes.NewStringValue(sdk.LogDestinationType)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "LogDestinationType", "value": sdk.LogDestinationType})
	// Handling Lists
	if sdk.LogType != nil {
		tflog.Debug(ctx, "Packing list of primitives for field LogType")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.LogType, d = basetypes.NewListValueFrom(ctx, elemType, sdk.LogType)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.LogType = basetypes.NewListNull(elemType)
	}
	// Handling Primitives
	// Standard primitive packing
	model.RoleType = basetypes.NewStringValue(sdk.RoleType)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "RoleType", "value": sdk.RoleType})
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.NgfirewallLogConfig{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.NgfirewallLogConfig", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for NgfirewallLogConfig ---
func unpackNgfirewallLogConfigListToSdk(ctx context.Context, list types.List) ([]cloud_ngfw.NgfirewallLogConfig, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.NgfirewallLogConfig")
	diags := diag.Diagnostics{}
	var data []models.NgfirewallLogConfig
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]cloud_ngfw.NgfirewallLogConfig, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.NgfirewallLogConfig{}.AttrTypes(), &item)
		unpacked, d := unpackNgfirewallLogConfigToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.NgfirewallLogConfig", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for NgfirewallLogConfig ---
func packNgfirewallLogConfigListFromSdk(ctx context.Context, sdks []cloud_ngfw.NgfirewallLogConfig) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.NgfirewallLogConfig")
	diags := diag.Diagnostics{}
	var data []models.NgfirewallLogConfig

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.NgfirewallLogConfig
		obj, d := packNgfirewallLogConfigFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.NgfirewallLogConfig{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.NgfirewallLogConfig", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.NgfirewallLogConfig{}.AttrType(), data)
}

// --- Unpacker for NgfirewallNotification ---
func unpackNgfirewallNotificationToSdk(ctx context.Context, obj types.Object) (*cloud_ngfw.NgfirewallNotification, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.NgfirewallNotification", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.NgfirewallNotification
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk cloud_ngfw.NgfirewallNotification
	var d diag.Diagnostics
	// Handling Primitives
	if !model.ExpiryTime.IsNull() && !model.ExpiryTime.IsUnknown() {
		val := int32(model.ExpiryTime.ValueInt64())
		sdk.ExpiryTime = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "ExpiryTime", "value": *sdk.ExpiryTime})
	}

	// Handling Primitives
	if !model.Id.IsNull() && !model.Id.IsUnknown() {
		sdk.Id = model.Id.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Id", "value": *sdk.Id})
	}

	// Handling Objects
	if !model.Link.IsNull() && !model.Link.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Link")
		unpacked, d := unpackNgfirewallNotificationLinkToSdk(ctx, model.Link)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Link"})
		}
		if unpacked != nil {
			sdk.Link = unpacked
		}
	}

	// Handling Primitives
	if !model.Msg.IsNull() && !model.Msg.IsUnknown() {
		sdk.Msg = model.Msg.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Msg", "value": *sdk.Msg})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.NgfirewallNotification", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for NgfirewallNotification ---
func packNgfirewallNotificationFromSdk(ctx context.Context, sdk cloud_ngfw.NgfirewallNotification) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.NgfirewallNotification", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.NgfirewallNotification
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.ExpiryTime != nil {
		model.ExpiryTime = basetypes.NewInt64Value(int64(*sdk.ExpiryTime))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "ExpiryTime", "value": *sdk.ExpiryTime})
	} else {
		model.ExpiryTime = basetypes.NewInt64Null()
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
	if sdk.Link != nil {
		tflog.Debug(ctx, "Packing nested object for field Link")
		packed, d := packNgfirewallNotificationLinkFromSdk(ctx, *sdk.Link)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Link"})
		}
		model.Link = packed
	} else {
		model.Link = basetypes.NewObjectNull(models.NgfirewallNotificationLink{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Msg != nil {
		model.Msg = basetypes.NewStringValue(*sdk.Msg)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Msg", "value": *sdk.Msg})
	} else {
		model.Msg = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.NgfirewallNotification{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.NgfirewallNotification", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for NgfirewallNotification ---
func unpackNgfirewallNotificationListToSdk(ctx context.Context, list types.List) ([]cloud_ngfw.NgfirewallNotification, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.NgfirewallNotification")
	diags := diag.Diagnostics{}
	var data []models.NgfirewallNotification
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]cloud_ngfw.NgfirewallNotification, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.NgfirewallNotification{}.AttrTypes(), &item)
		unpacked, d := unpackNgfirewallNotificationToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.NgfirewallNotification", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for NgfirewallNotification ---
func packNgfirewallNotificationListFromSdk(ctx context.Context, sdks []cloud_ngfw.NgfirewallNotification) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.NgfirewallNotification")
	diags := diag.Diagnostics{}
	var data []models.NgfirewallNotification

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.NgfirewallNotification
		obj, d := packNgfirewallNotificationFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.NgfirewallNotification{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.NgfirewallNotification", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.NgfirewallNotification{}.AttrType(), data)
}

// --- Unpacker for NgfirewallNotificationLink ---
func unpackNgfirewallNotificationLinkToSdk(ctx context.Context, obj types.Object) (*cloud_ngfw.NgfirewallNotificationLink, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.NgfirewallNotificationLink", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.NgfirewallNotificationLink
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk cloud_ngfw.NgfirewallNotificationLink
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Target.IsNull() && !model.Target.IsUnknown() {
		sdk.Target = model.Target.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Target", "value": *sdk.Target})
	}

	// Handling Primitives
	if !model.Text.IsNull() && !model.Text.IsUnknown() {
		sdk.Text = model.Text.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Text", "value": *sdk.Text})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.NgfirewallNotificationLink", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for NgfirewallNotificationLink ---
func packNgfirewallNotificationLinkFromSdk(ctx context.Context, sdk cloud_ngfw.NgfirewallNotificationLink) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.NgfirewallNotificationLink", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.NgfirewallNotificationLink
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Target != nil {
		model.Target = basetypes.NewStringValue(*sdk.Target)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Target", "value": *sdk.Target})
	} else {
		model.Target = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Text != nil {
		model.Text = basetypes.NewStringValue(*sdk.Text)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Text", "value": *sdk.Text})
	} else {
		model.Text = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.NgfirewallNotificationLink{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.NgfirewallNotificationLink", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for NgfirewallNotificationLink ---
func unpackNgfirewallNotificationLinkListToSdk(ctx context.Context, list types.List) ([]cloud_ngfw.NgfirewallNotificationLink, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.NgfirewallNotificationLink")
	diags := diag.Diagnostics{}
	var data []models.NgfirewallNotificationLink
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]cloud_ngfw.NgfirewallNotificationLink, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.NgfirewallNotificationLink{}.AttrTypes(), &item)
		unpacked, d := unpackNgfirewallNotificationLinkToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.NgfirewallNotificationLink", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for NgfirewallNotificationLink ---
func packNgfirewallNotificationLinkListFromSdk(ctx context.Context, sdks []cloud_ngfw.NgfirewallNotificationLink) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.NgfirewallNotificationLink")
	diags := diag.Diagnostics{}
	var data []models.NgfirewallNotificationLink

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.NgfirewallNotificationLink
		obj, d := packNgfirewallNotificationLinkFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.NgfirewallNotificationLink{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.NgfirewallNotificationLink", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.NgfirewallNotificationLink{}.AttrType(), data)
}

// --- Unpacker for NgfirewallPrivateAccess ---
func unpackNgfirewallPrivateAccessToSdk(ctx context.Context, obj types.Object) (*cloud_ngfw.NgfirewallPrivateAccess, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.NgfirewallPrivateAccess", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.NgfirewallPrivateAccess
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk cloud_ngfw.NgfirewallPrivateAccess
	var d diag.Diagnostics
	// Handling Primitives
	if !model.ResourceID.IsNull() && !model.ResourceID.IsUnknown() {
		sdk.ResourceID = model.ResourceID.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "ResourceID", "value": *sdk.ResourceID})
	}

	// Handling Primitives
	if !model.Type.IsNull() && !model.Type.IsUnknown() {
		sdk.Type = model.Type.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Type", "value": *sdk.Type})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.NgfirewallPrivateAccess", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for NgfirewallPrivateAccess ---
func packNgfirewallPrivateAccessFromSdk(ctx context.Context, sdk cloud_ngfw.NgfirewallPrivateAccess) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.NgfirewallPrivateAccess", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.NgfirewallPrivateAccess
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.ResourceID != nil {
		model.ResourceID = basetypes.NewStringValue(*sdk.ResourceID)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "ResourceID", "value": *sdk.ResourceID})
	} else {
		model.ResourceID = basetypes.NewStringNull()
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

	obj, d := types.ObjectValueFrom(ctx, models.NgfirewallPrivateAccess{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.NgfirewallPrivateAccess", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for NgfirewallPrivateAccess ---
func unpackNgfirewallPrivateAccessListToSdk(ctx context.Context, list types.List) ([]cloud_ngfw.NgfirewallPrivateAccess, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.NgfirewallPrivateAccess")
	diags := diag.Diagnostics{}
	var data []models.NgfirewallPrivateAccess
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]cloud_ngfw.NgfirewallPrivateAccess, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.NgfirewallPrivateAccess{}.AttrTypes(), &item)
		unpacked, d := unpackNgfirewallPrivateAccessToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.NgfirewallPrivateAccess", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for NgfirewallPrivateAccess ---
func packNgfirewallPrivateAccessListFromSdk(ctx context.Context, sdks []cloud_ngfw.NgfirewallPrivateAccess) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.NgfirewallPrivateAccess")
	diags := diag.Diagnostics{}
	var data []models.NgfirewallPrivateAccess

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.NgfirewallPrivateAccess
		obj, d := packNgfirewallPrivateAccessFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.NgfirewallPrivateAccess{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.NgfirewallPrivateAccess", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.NgfirewallPrivateAccess{}.AttrType(), data)
}

// --- Unpacker for NgfirewallTag ---
func unpackNgfirewallTagToSdk(ctx context.Context, obj types.Object) (*cloud_ngfw.NgfirewallTag, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.NgfirewallTag", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.NgfirewallTag
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk cloud_ngfw.NgfirewallTag
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Key.IsNull() && !model.Key.IsUnknown() {
		sdk.Key = model.Key.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Key", "value": *sdk.Key})
	}

	// Handling Primitives
	if !model.Value.IsNull() && !model.Value.IsUnknown() {
		sdk.Value = model.Value.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Value", "value": *sdk.Value})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.NgfirewallTag", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for NgfirewallTag ---
func packNgfirewallTagFromSdk(ctx context.Context, sdk cloud_ngfw.NgfirewallTag) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.NgfirewallTag", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.NgfirewallTag
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Key != nil {
		model.Key = basetypes.NewStringValue(*sdk.Key)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Key", "value": *sdk.Key})
	} else {
		model.Key = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Value != nil {
		model.Value = basetypes.NewStringValue(*sdk.Value)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Value", "value": *sdk.Value})
	} else {
		model.Value = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.NgfirewallTag{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.NgfirewallTag", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for NgfirewallTag ---
func unpackNgfirewallTagListToSdk(ctx context.Context, list types.List) ([]cloud_ngfw.NgfirewallTag, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.NgfirewallTag")
	diags := diag.Diagnostics{}
	var data []models.NgfirewallTag
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]cloud_ngfw.NgfirewallTag, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.NgfirewallTag{}.AttrTypes(), &item)
		unpacked, d := unpackNgfirewallTagToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.NgfirewallTag", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for NgfirewallTag ---
func packNgfirewallTagListFromSdk(ctx context.Context, sdks []cloud_ngfw.NgfirewallTag) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.NgfirewallTag")
	diags := diag.Diagnostics{}
	var data []models.NgfirewallTag

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.NgfirewallTag
		obj, d := packNgfirewallTagFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.NgfirewallTag{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.NgfirewallTag", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.NgfirewallTag{}.AttrType(), data)
}

// --- Unpacker for NgfirewallUserId ---
func unpackNgfirewallUserIdToSdk(ctx context.Context, obj types.Object) (*cloud_ngfw.NgfirewallUserId, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.NgfirewallUserId", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.NgfirewallUserId
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk cloud_ngfw.NgfirewallUserId
	var d diag.Diagnostics
	// Handling Primitives
	if !model.AgentName.IsNull() && !model.AgentName.IsUnknown() {
		sdk.AgentName = model.AgentName.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "AgentName", "value": *sdk.AgentName})
	}

	// Handling Primitives
	if !model.CFTURLLink.IsNull() && !model.CFTURLLink.IsUnknown() {
		sdk.CFTURLLink = model.CFTURLLink.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "CFTURLLink", "value": *sdk.CFTURLLink})
	}

	// Handling Primitives
	if !model.CollectorName.IsNull() && !model.CollectorName.IsUnknown() {
		sdk.CollectorName = model.CollectorName.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "CollectorName", "value": *sdk.CollectorName})
	}

	// Handling Lists
	if !model.CustomIncludeExcludeNetwork.IsNull() && !model.CustomIncludeExcludeNetwork.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field CustomIncludeExcludeNetwork")
		unpacked, d := unpackNgfirewallUserIdNetworkListToSdk(ctx, model.CustomIncludeExcludeNetwork)
		diags.Append(d...)
		sdk.CustomIncludeExcludeNetwork = unpacked
	}

	// Handling Primitives
	if !model.Enabled.IsNull() && !model.Enabled.IsUnknown() {
		sdk.Enabled = model.Enabled.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Enabled", "value": *sdk.Enabled})
	}

	// Handling Primitives
	if !model.EndpointDNS.IsNull() && !model.EndpointDNS.IsUnknown() {
		sdk.EndpointDNS = model.EndpointDNS.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "EndpointDNS", "value": *sdk.EndpointDNS})
	}

	// Handling Primitives
	if !model.Port.IsNull() && !model.Port.IsUnknown() {
		val := int32(model.Port.ValueInt64())
		sdk.Port = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Port", "value": *sdk.Port})
	}

	// Handling Primitives
	if !model.SecretKeyARN.IsNull() && !model.SecretKeyARN.IsUnknown() {
		sdk.SecretKeyARN = model.SecretKeyARN.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "SecretKeyARN", "value": *sdk.SecretKeyARN})
	}

	// Handling Primitives
	if !model.UserIDStatus.IsNull() && !model.UserIDStatus.IsUnknown() {
		sdk.UserIDStatus = model.UserIDStatus.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "UserIDStatus", "value": *sdk.UserIDStatus})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.NgfirewallUserId", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for NgfirewallUserId ---
func packNgfirewallUserIdFromSdk(ctx context.Context, sdk cloud_ngfw.NgfirewallUserId) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.NgfirewallUserId", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.NgfirewallUserId
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.AgentName != nil {
		model.AgentName = basetypes.NewStringValue(*sdk.AgentName)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "AgentName", "value": *sdk.AgentName})
	} else {
		model.AgentName = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.CFTURLLink != nil {
		model.CFTURLLink = basetypes.NewStringValue(*sdk.CFTURLLink)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "CFTURLLink", "value": *sdk.CFTURLLink})
	} else {
		model.CFTURLLink = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.CollectorName != nil {
		model.CollectorName = basetypes.NewStringValue(*sdk.CollectorName)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "CollectorName", "value": *sdk.CollectorName})
	} else {
		model.CollectorName = basetypes.NewStringNull()
	}
	// Handling Lists
	if sdk.CustomIncludeExcludeNetwork != nil {
		tflog.Debug(ctx, "Packing list of objects for field CustomIncludeExcludeNetwork")
		packed, d := packNgfirewallUserIdNetworkListFromSdk(ctx, sdk.CustomIncludeExcludeNetwork)
		diags.Append(d...)
		model.CustomIncludeExcludeNetwork = packed
	} else {
		model.CustomIncludeExcludeNetwork = basetypes.NewListNull(models.NgfirewallUserIdNetwork{}.AttrType())
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
	if sdk.EndpointDNS != nil {
		model.EndpointDNS = basetypes.NewStringValue(*sdk.EndpointDNS)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "EndpointDNS", "value": *sdk.EndpointDNS})
	} else {
		model.EndpointDNS = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Port != nil {
		model.Port = basetypes.NewInt64Value(int64(*sdk.Port))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Port", "value": *sdk.Port})
	} else {
		model.Port = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.SecretKeyARN != nil {
		model.SecretKeyARN = basetypes.NewStringValue(*sdk.SecretKeyARN)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "SecretKeyARN", "value": *sdk.SecretKeyARN})
	} else {
		model.SecretKeyARN = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.UserIDStatus != nil {
		model.UserIDStatus = basetypes.NewStringValue(*sdk.UserIDStatus)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "UserIDStatus", "value": *sdk.UserIDStatus})
	} else {
		model.UserIDStatus = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.NgfirewallUserId{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.NgfirewallUserId", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for NgfirewallUserId ---
func unpackNgfirewallUserIdListToSdk(ctx context.Context, list types.List) ([]cloud_ngfw.NgfirewallUserId, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.NgfirewallUserId")
	diags := diag.Diagnostics{}
	var data []models.NgfirewallUserId
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]cloud_ngfw.NgfirewallUserId, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.NgfirewallUserId{}.AttrTypes(), &item)
		unpacked, d := unpackNgfirewallUserIdToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.NgfirewallUserId", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for NgfirewallUserId ---
func packNgfirewallUserIdListFromSdk(ctx context.Context, sdks []cloud_ngfw.NgfirewallUserId) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.NgfirewallUserId")
	diags := diag.Diagnostics{}
	var data []models.NgfirewallUserId

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.NgfirewallUserId
		obj, d := packNgfirewallUserIdFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.NgfirewallUserId{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.NgfirewallUserId", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.NgfirewallUserId{}.AttrType(), data)
}

// --- Unpacker for NgfirewallUserIdNetwork ---
func unpackNgfirewallUserIdNetworkToSdk(ctx context.Context, obj types.Object) (*cloud_ngfw.NgfirewallUserIdNetwork, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.NgfirewallUserIdNetwork", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.NgfirewallUserIdNetwork
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk cloud_ngfw.NgfirewallUserIdNetwork
	var d diag.Diagnostics
	// Handling Primitives
	if !model.DiscoveryInclude.IsNull() && !model.DiscoveryInclude.IsUnknown() {
		sdk.DiscoveryInclude = model.DiscoveryInclude.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DiscoveryInclude", "value": *sdk.DiscoveryInclude})
	}

	// Handling Primitives
	if !model.Enabled.IsNull() && !model.Enabled.IsUnknown() {
		sdk.Enabled = model.Enabled.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Enabled", "value": *sdk.Enabled})
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
	}

	// Handling Primitives
	if !model.NetworkAddress.IsNull() && !model.NetworkAddress.IsUnknown() {
		sdk.NetworkAddress = model.NetworkAddress.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "NetworkAddress", "value": *sdk.NetworkAddress})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.NgfirewallUserIdNetwork", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for NgfirewallUserIdNetwork ---
func packNgfirewallUserIdNetworkFromSdk(ctx context.Context, sdk cloud_ngfw.NgfirewallUserIdNetwork) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.NgfirewallUserIdNetwork", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.NgfirewallUserIdNetwork
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.DiscoveryInclude != nil {
		model.DiscoveryInclude = basetypes.NewBoolValue(*sdk.DiscoveryInclude)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DiscoveryInclude", "value": *sdk.DiscoveryInclude})
	} else {
		model.DiscoveryInclude = basetypes.NewBoolNull()
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
	if sdk.Name != nil {
		model.Name = basetypes.NewStringValue(*sdk.Name)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
	} else {
		model.Name = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.NetworkAddress != nil {
		model.NetworkAddress = basetypes.NewStringValue(*sdk.NetworkAddress)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "NetworkAddress", "value": *sdk.NetworkAddress})
	} else {
		model.NetworkAddress = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.NgfirewallUserIdNetwork{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.NgfirewallUserIdNetwork", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for NgfirewallUserIdNetwork ---
func unpackNgfirewallUserIdNetworkListToSdk(ctx context.Context, list types.List) ([]cloud_ngfw.NgfirewallUserIdNetwork, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.NgfirewallUserIdNetwork")
	diags := diag.Diagnostics{}
	var data []models.NgfirewallUserIdNetwork
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]cloud_ngfw.NgfirewallUserIdNetwork, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.NgfirewallUserIdNetwork{}.AttrTypes(), &item)
		unpacked, d := unpackNgfirewallUserIdNetworkToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.NgfirewallUserIdNetwork", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for NgfirewallUserIdNetwork ---
func packNgfirewallUserIdNetworkListFromSdk(ctx context.Context, sdks []cloud_ngfw.NgfirewallUserIdNetwork) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.NgfirewallUserIdNetwork")
	diags := diag.Diagnostics{}
	var data []models.NgfirewallUserIdNetwork

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.NgfirewallUserIdNetwork
		obj, d := packNgfirewallUserIdNetworkFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.NgfirewallUserIdNetwork{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.NgfirewallUserIdNetwork", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.NgfirewallUserIdNetwork{}.AttrType(), data)
}

// --- Unpacker for NgfirewallStatus ---
func unpackNgfirewallStatusToSdk(ctx context.Context, obj types.Object) (*cloud_ngfw.NgfirewallStatus, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.NgfirewallStatus", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.NgfirewallStatus
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk cloud_ngfw.NgfirewallStatus
	var d diag.Diagnostics
	// Handling Objects
	if !model.DeviceRuleStackCommitInfo.IsNull() && !model.DeviceRuleStackCommitInfo.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field DeviceRuleStackCommitInfo")
		unpacked, d := unpackNgfirewallCommitInfoToSdk(ctx, model.DeviceRuleStackCommitInfo)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "DeviceRuleStackCommitInfo"})
		}
		if unpacked != nil {
			sdk.DeviceRuleStackCommitInfo = unpacked
		}
	}

	// Handling Primitives
	if !model.DeviceRuleStackCommitState.IsNull() && !model.DeviceRuleStackCommitState.IsUnknown() {
		sdk.DeviceRuleStackCommitState = model.DeviceRuleStackCommitState.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DeviceRuleStackCommitState", "value": *sdk.DeviceRuleStackCommitState})
	}

	// Handling Primitives
	if !model.DeviceRuleStackCommitStatus.IsNull() && !model.DeviceRuleStackCommitStatus.IsUnknown() {
		sdk.DeviceRuleStackCommitStatus = model.DeviceRuleStackCommitStatus.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DeviceRuleStackCommitStatus", "value": *sdk.DeviceRuleStackCommitStatus})
	}

	// Handling Primitives
	if !model.FailureReason.IsNull() && !model.FailureReason.IsUnknown() {
		sdk.FailureReason = model.FailureReason.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "FailureReason", "value": *sdk.FailureReason})
	}

	// Handling Primitives
	if !model.FirewallStatus.IsNull() && !model.FirewallStatus.IsUnknown() {
		sdk.FirewallStatus = model.FirewallStatus.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "FirewallStatus", "value": *sdk.FirewallStatus})
	}

	// Handling Objects
	if !model.GlobalRuleStackCommitInfo.IsNull() && !model.GlobalRuleStackCommitInfo.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field GlobalRuleStackCommitInfo")
		unpacked, d := unpackNgfirewallCommitInfoToSdk(ctx, model.GlobalRuleStackCommitInfo)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "GlobalRuleStackCommitInfo"})
		}
		if unpacked != nil {
			sdk.GlobalRuleStackCommitInfo = unpacked
		}
	}

	// Handling Primitives
	if !model.GlobalRuleStackStatus.IsNull() && !model.GlobalRuleStackStatus.IsUnknown() {
		sdk.GlobalRuleStackStatus = model.GlobalRuleStackStatus.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "GlobalRuleStackStatus", "value": *sdk.GlobalRuleStackStatus})
	}

	// Handling Lists
	if !model.PublicIPs.IsNull() && !model.PublicIPs.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field PublicIPs")
		unpacked, d := unpackNgfirewallPublicIpListToSdk(ctx, model.PublicIPs)
		diags.Append(d...)
		sdk.PublicIPs = unpacked
	}

	// Handling Objects
	if !model.RuleStackCommitInfo.IsNull() && !model.RuleStackCommitInfo.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field RuleStackCommitInfo")
		unpacked, d := unpackNgfirewallCommitInfoToSdk(ctx, model.RuleStackCommitInfo)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "RuleStackCommitInfo"})
		}
		if unpacked != nil {
			sdk.RuleStackCommitInfo = unpacked
		}
	}

	// Handling Primitives
	if !model.RuleStackStatus.IsNull() && !model.RuleStackStatus.IsUnknown() {
		sdk.RuleStackStatus = model.RuleStackStatus.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "RuleStackStatus", "value": *sdk.RuleStackStatus})
	}

	// Handling Primitives
	if !model.SCMAssocStatus.IsNull() && !model.SCMAssocStatus.IsUnknown() {
		sdk.SCMAssocStatus = model.SCMAssocStatus.ValueBoolPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "SCMAssocStatus", "value": *sdk.SCMAssocStatus})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.NgfirewallStatus", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for NgfirewallStatus ---
func packNgfirewallStatusFromSdk(ctx context.Context, sdk cloud_ngfw.NgfirewallStatus) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.NgfirewallStatus", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.NgfirewallStatus
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.DeviceRuleStackCommitInfo != nil {
		tflog.Debug(ctx, "Packing nested object for field DeviceRuleStackCommitInfo")
		packed, d := packNgfirewallCommitInfoFromSdk(ctx, *sdk.DeviceRuleStackCommitInfo)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "DeviceRuleStackCommitInfo"})
		}
		model.DeviceRuleStackCommitInfo = packed
	} else {
		model.DeviceRuleStackCommitInfo = basetypes.NewObjectNull(models.NgfirewallCommitInfo{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.DeviceRuleStackCommitState != nil {
		model.DeviceRuleStackCommitState = basetypes.NewStringValue(*sdk.DeviceRuleStackCommitState)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DeviceRuleStackCommitState", "value": *sdk.DeviceRuleStackCommitState})
	} else {
		model.DeviceRuleStackCommitState = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.DeviceRuleStackCommitStatus != nil {
		model.DeviceRuleStackCommitStatus = basetypes.NewStringValue(*sdk.DeviceRuleStackCommitStatus)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DeviceRuleStackCommitStatus", "value": *sdk.DeviceRuleStackCommitStatus})
	} else {
		model.DeviceRuleStackCommitStatus = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.FailureReason != nil {
		model.FailureReason = basetypes.NewStringValue(*sdk.FailureReason)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "FailureReason", "value": *sdk.FailureReason})
	} else {
		model.FailureReason = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.FirewallStatus != nil {
		model.FirewallStatus = basetypes.NewStringValue(*sdk.FirewallStatus)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "FirewallStatus", "value": *sdk.FirewallStatus})
	} else {
		model.FirewallStatus = basetypes.NewStringNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.GlobalRuleStackCommitInfo != nil {
		tflog.Debug(ctx, "Packing nested object for field GlobalRuleStackCommitInfo")
		packed, d := packNgfirewallCommitInfoFromSdk(ctx, *sdk.GlobalRuleStackCommitInfo)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "GlobalRuleStackCommitInfo"})
		}
		model.GlobalRuleStackCommitInfo = packed
	} else {
		model.GlobalRuleStackCommitInfo = basetypes.NewObjectNull(models.NgfirewallCommitInfo{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.GlobalRuleStackStatus != nil {
		model.GlobalRuleStackStatus = basetypes.NewStringValue(*sdk.GlobalRuleStackStatus)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "GlobalRuleStackStatus", "value": *sdk.GlobalRuleStackStatus})
	} else {
		model.GlobalRuleStackStatus = basetypes.NewStringNull()
	}
	// Handling Lists
	if sdk.PublicIPs != nil {
		tflog.Debug(ctx, "Packing list of objects for field PublicIPs")
		packed, d := packNgfirewallPublicIpListFromSdk(ctx, sdk.PublicIPs)
		diags.Append(d...)
		model.PublicIPs = packed
	} else {
		model.PublicIPs = basetypes.NewListNull(models.NgfirewallPublicIp{}.AttrType())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.RuleStackCommitInfo != nil {
		tflog.Debug(ctx, "Packing nested object for field RuleStackCommitInfo")
		packed, d := packNgfirewallCommitInfoFromSdk(ctx, *sdk.RuleStackCommitInfo)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "RuleStackCommitInfo"})
		}
		model.RuleStackCommitInfo = packed
	} else {
		model.RuleStackCommitInfo = basetypes.NewObjectNull(models.NgfirewallCommitInfo{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.RuleStackStatus != nil {
		model.RuleStackStatus = basetypes.NewStringValue(*sdk.RuleStackStatus)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "RuleStackStatus", "value": *sdk.RuleStackStatus})
	} else {
		model.RuleStackStatus = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.SCMAssocStatus != nil {
		model.SCMAssocStatus = basetypes.NewBoolValue(*sdk.SCMAssocStatus)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "SCMAssocStatus", "value": *sdk.SCMAssocStatus})
	} else {
		model.SCMAssocStatus = basetypes.NewBoolNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.NgfirewallStatus{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.NgfirewallStatus", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for NgfirewallStatus ---
func unpackNgfirewallStatusListToSdk(ctx context.Context, list types.List) ([]cloud_ngfw.NgfirewallStatus, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.NgfirewallStatus")
	diags := diag.Diagnostics{}
	var data []models.NgfirewallStatus
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]cloud_ngfw.NgfirewallStatus, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.NgfirewallStatus{}.AttrTypes(), &item)
		unpacked, d := unpackNgfirewallStatusToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.NgfirewallStatus", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for NgfirewallStatus ---
func packNgfirewallStatusListFromSdk(ctx context.Context, sdks []cloud_ngfw.NgfirewallStatus) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.NgfirewallStatus")
	diags := diag.Diagnostics{}
	var data []models.NgfirewallStatus

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.NgfirewallStatus
		obj, d := packNgfirewallStatusFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.NgfirewallStatus{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.NgfirewallStatus", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.NgfirewallStatus{}.AttrType(), data)
}

// --- Unpacker for NgfirewallCommitInfo ---
func unpackNgfirewallCommitInfoToSdk(ctx context.Context, obj types.Object) (*cloud_ngfw.NgfirewallCommitInfo, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.NgfirewallCommitInfo", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.NgfirewallCommitInfo
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk cloud_ngfw.NgfirewallCommitInfo
	var d diag.Diagnostics
	// Handling Lists
	if !model.CommitMessages.IsNull() && !model.CommitMessages.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field CommitMessages")
		diags.Append(model.CommitMessages.ElementsAs(ctx, &sdk.CommitMessages, false)...)
	}

	// Handling Primitives
	if !model.CommitTS.IsNull() && !model.CommitTS.IsUnknown() {
		sdk.CommitTS = model.CommitTS.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "CommitTS", "value": *sdk.CommitTS})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.NgfirewallCommitInfo", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for NgfirewallCommitInfo ---
func packNgfirewallCommitInfoFromSdk(ctx context.Context, sdk cloud_ngfw.NgfirewallCommitInfo) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.NgfirewallCommitInfo", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.NgfirewallCommitInfo
	var d diag.Diagnostics
	// Handling Lists
	if sdk.CommitMessages != nil {
		tflog.Debug(ctx, "Packing list of primitives for field CommitMessages")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.CommitMessages, d = basetypes.NewListValueFrom(ctx, elemType, sdk.CommitMessages)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.CommitMessages = basetypes.NewListNull(elemType)
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.CommitTS != nil {
		model.CommitTS = basetypes.NewStringValue(*sdk.CommitTS)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "CommitTS", "value": *sdk.CommitTS})
	} else {
		model.CommitTS = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.NgfirewallCommitInfo{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.NgfirewallCommitInfo", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for NgfirewallCommitInfo ---
func unpackNgfirewallCommitInfoListToSdk(ctx context.Context, list types.List) ([]cloud_ngfw.NgfirewallCommitInfo, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.NgfirewallCommitInfo")
	diags := diag.Diagnostics{}
	var data []models.NgfirewallCommitInfo
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]cloud_ngfw.NgfirewallCommitInfo, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.NgfirewallCommitInfo{}.AttrTypes(), &item)
		unpacked, d := unpackNgfirewallCommitInfoToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.NgfirewallCommitInfo", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for NgfirewallCommitInfo ---
func packNgfirewallCommitInfoListFromSdk(ctx context.Context, sdks []cloud_ngfw.NgfirewallCommitInfo) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.NgfirewallCommitInfo")
	diags := diag.Diagnostics{}
	var data []models.NgfirewallCommitInfo

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.NgfirewallCommitInfo
		obj, d := packNgfirewallCommitInfoFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.NgfirewallCommitInfo{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.NgfirewallCommitInfo", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.NgfirewallCommitInfo{}.AttrType(), data)
}

// --- Unpacker for NgfirewallPublicIp ---
func unpackNgfirewallPublicIpToSdk(ctx context.Context, obj types.Object) (*cloud_ngfw.NgfirewallPublicIp, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.NgfirewallPublicIp", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.NgfirewallPublicIp
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk cloud_ngfw.NgfirewallPublicIp
	var d diag.Diagnostics
	// Handling Primitives
	if !model.IPAddress.IsNull() && !model.IPAddress.IsUnknown() {
		sdk.IPAddress = model.IPAddress.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "IPAddress", "value": *sdk.IPAddress})
	}

	// Handling Primitives
	if !model.IPSource.IsNull() && !model.IPSource.IsUnknown() {
		sdk.IPSource = model.IPSource.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "IPSource", "value": *sdk.IPSource})
	}

	// Handling Primitives
	if !model.IPStatus.IsNull() && !model.IPStatus.IsUnknown() {
		sdk.IPStatus = model.IPStatus.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "IPStatus", "value": *sdk.IPStatus})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.NgfirewallPublicIp", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for NgfirewallPublicIp ---
func packNgfirewallPublicIpFromSdk(ctx context.Context, sdk cloud_ngfw.NgfirewallPublicIp) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.NgfirewallPublicIp", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.NgfirewallPublicIp
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.IPAddress != nil {
		model.IPAddress = basetypes.NewStringValue(*sdk.IPAddress)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "IPAddress", "value": *sdk.IPAddress})
	} else {
		model.IPAddress = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.IPSource != nil {
		model.IPSource = basetypes.NewStringValue(*sdk.IPSource)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "IPSource", "value": *sdk.IPSource})
	} else {
		model.IPSource = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.IPStatus != nil {
		model.IPStatus = basetypes.NewStringValue(*sdk.IPStatus)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "IPStatus", "value": *sdk.IPStatus})
	} else {
		model.IPStatus = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.NgfirewallPublicIp{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.NgfirewallPublicIp", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for NgfirewallPublicIp ---
func unpackNgfirewallPublicIpListToSdk(ctx context.Context, list types.List) ([]cloud_ngfw.NgfirewallPublicIp, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.NgfirewallPublicIp")
	diags := diag.Diagnostics{}
	var data []models.NgfirewallPublicIp
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]cloud_ngfw.NgfirewallPublicIp, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.NgfirewallPublicIp{}.AttrTypes(), &item)
		unpacked, d := unpackNgfirewallPublicIpToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.NgfirewallPublicIp", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for NgfirewallPublicIp ---
func packNgfirewallPublicIpListFromSdk(ctx context.Context, sdks []cloud_ngfw.NgfirewallPublicIp) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.NgfirewallPublicIp")
	diags := diag.Diagnostics{}
	var data []models.NgfirewallPublicIp

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.NgfirewallPublicIp
		obj, d := packNgfirewallPublicIpFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.NgfirewallPublicIp{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.NgfirewallPublicIp", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.NgfirewallPublicIp{}.AttrType(), data)
}

// --- Unpacker for NgfirewallResponseStatus ---
func unpackNgfirewallResponseStatusToSdk(ctx context.Context, obj types.Object) (*cloud_ngfw.NgfirewallResponseStatus, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.NgfirewallResponseStatus", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.NgfirewallResponseStatus
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk cloud_ngfw.NgfirewallResponseStatus
	var d diag.Diagnostics
	// Handling Primitives
	if !model.ErrorCode.IsNull() && !model.ErrorCode.IsUnknown() {
		val := int32(model.ErrorCode.ValueInt64())
		sdk.ErrorCode = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "ErrorCode", "value": *sdk.ErrorCode})
	}

	// Handling Primitives
	if !model.Reason.IsNull() && !model.Reason.IsUnknown() {
		sdk.Reason = model.Reason.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Reason", "value": *sdk.Reason})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.NgfirewallResponseStatus", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for NgfirewallResponseStatus ---
func packNgfirewallResponseStatusFromSdk(ctx context.Context, sdk cloud_ngfw.NgfirewallResponseStatus) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.NgfirewallResponseStatus", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.NgfirewallResponseStatus
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.ErrorCode != nil {
		model.ErrorCode = basetypes.NewInt64Value(int64(*sdk.ErrorCode))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "ErrorCode", "value": *sdk.ErrorCode})
	} else {
		model.ErrorCode = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Reason != nil {
		model.Reason = basetypes.NewStringValue(*sdk.Reason)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Reason", "value": *sdk.Reason})
	} else {
		model.Reason = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.NgfirewallResponseStatus{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.NgfirewallResponseStatus", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for NgfirewallResponseStatus ---
func unpackNgfirewallResponseStatusListToSdk(ctx context.Context, list types.List) ([]cloud_ngfw.NgfirewallResponseStatus, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.NgfirewallResponseStatus")
	diags := diag.Diagnostics{}
	var data []models.NgfirewallResponseStatus
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]cloud_ngfw.NgfirewallResponseStatus, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.NgfirewallResponseStatus{}.AttrTypes(), &item)
		unpacked, d := unpackNgfirewallResponseStatusToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.NgfirewallResponseStatus", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for NgfirewallResponseStatus ---
func packNgfirewallResponseStatusListFromSdk(ctx context.Context, sdks []cloud_ngfw.NgfirewallResponseStatus) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.NgfirewallResponseStatus")
	diags := diag.Diagnostics{}
	var data []models.NgfirewallResponseStatus

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.NgfirewallResponseStatus
		obj, d := packNgfirewallResponseStatusFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.NgfirewallResponseStatus{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.NgfirewallResponseStatus", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.NgfirewallResponseStatus{}.AttrType(), data)
}

// --- Unpacker for NgfirewallCreateRequest ---
func unpackNgfirewallCreateRequestToSdk(ctx context.Context, obj types.Object) (*cloud_ngfw.NgfirewallCreateRequest, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.NgfirewallCreateRequest", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.NgfirewallCreateRequest
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk cloud_ngfw.NgfirewallCreateRequest
	var d diag.Diagnostics

	// Handling Lists
	if !model.AllowListAccounts.IsNull() && !model.AllowListAccounts.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field AllowListAccounts")
		diags.Append(model.AllowListAccounts.ElementsAs(ctx, &sdk.AllowListAccounts, false)...)
	}

	// Handling Lists
	if !model.ChangeProtection.IsNull() && !model.ChangeProtection.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field ChangeProtection")
		diags.Append(model.ChangeProtection.ElementsAs(ctx, &sdk.ChangeProtection, false)...)
	}

	// Handling Lists
	if !model.CustomerZoneIdList.IsNull() && !model.CustomerZoneIdList.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field CustomerZoneIdList")
		diags.Append(model.CustomerZoneIdList.ElementsAs(ctx, &sdk.CustomerZoneIdList, false)...)
	}

	// Handling Primitives
	if !model.Description.IsNull() && !model.Description.IsUnknown() {
		sdk.Description = model.Description.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Description", "value": *sdk.Description})
	}

	// Handling Objects
	if !model.EgressNAT.IsNull() && !model.EgressNAT.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field EgressNAT")
		unpacked, d := unpackNgfirewallEgressNatConfigToSdk(ctx, model.EgressNAT)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "EgressNAT"})
		}
		if unpacked != nil {
			sdk.EgressNAT = unpacked
		}
	}

	// Handling Lists
	if !model.Endpoints.IsNull() && !model.Endpoints.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field Endpoints")
		unpacked, d := unpackNgfirewallEndpointListToSdk(ctx, model.Endpoints)
		diags.Append(d...)
		sdk.Endpoints = unpacked
	}

	// Handling Primitives
	if !model.GwlbTcpIdleTimeout.IsNull() && !model.GwlbTcpIdleTimeout.IsUnknown() {
		val := int32(model.GwlbTcpIdleTimeout.ValueInt64())
		sdk.GwlbTcpIdleTimeout = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "GwlbTcpIdleTimeout", "value": *sdk.GwlbTcpIdleTimeout})
	}

	// Handling Primitives
	if !model.LinkId.IsNull() && !model.LinkId.IsUnknown() {
		sdk.LinkId = model.LinkId.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "LinkId", "value": *sdk.LinkId})
	}

	// Handling Primitives
	if !model.Mode.IsNull() && !model.Mode.IsUnknown() {
		sdk.Mode = model.Mode.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Mode", "value": *sdk.Mode})
	}

	// Handling Objects
	if !model.PrivateAccess.IsNull() && !model.PrivateAccess.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field PrivateAccess")
		unpacked, d := unpackNgfirewallPrivateAccessToSdk(ctx, model.PrivateAccess)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "PrivateAccess"})
		}
		if unpacked != nil {
			sdk.PrivateAccess = unpacked
		}
	}

	// Handling Lists
	if !model.ScmLabels.IsNull() && !model.ScmLabels.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field ScmLabels")
		diags.Append(model.ScmLabels.ElementsAs(ctx, &sdk.ScmLabels, false)...)
	}

	// Handling Primitives
	if !model.Source.IsNull() && !model.Source.IsUnknown() {
		sdk.Source = model.Source.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Source", "value": *sdk.Source})
	}

	// Handling Lists
	if !model.Tags.IsNull() && !model.Tags.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field Tags")
		unpacked, d := unpackNgfirewallTagListToSdk(ctx, model.Tags)
		diags.Append(d...)
		sdk.Tags = unpacked
	}

	// Handling Primitives
	if !model.Tier.IsNull() && !model.Tier.IsUnknown() {
		sdk.Tier = model.Tier.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Tier", "value": *sdk.Tier})
	}

	// Handling Objects
	if !model.UserID.IsNull() && !model.UserID.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field UserID")
		unpacked, d := unpackNgfirewallUserIdToSdk(ctx, model.UserID)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "UserID"})
		}
		if unpacked != nil {
			sdk.UserID = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.NgfirewallCreateRequest", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for NgfirewallCreateRequest ---
func packNgfirewallCreateRequestFromSdk(ctx context.Context, sdk cloud_ngfw.NgfirewallCreateRequest) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.NgfirewallCreateRequest", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.NgfirewallCreateRequest
	var d diag.Diagnostics
	// Handling Lists
	if sdk.AllowListAccounts != nil {
		tflog.Debug(ctx, "Packing list of primitives for field AllowListAccounts")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.AllowListAccounts, d = basetypes.NewListValueFrom(ctx, elemType, sdk.AllowListAccounts)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.AllowListAccounts = basetypes.NewListNull(elemType)
	}
	// Handling Lists
	if sdk.ChangeProtection != nil {
		tflog.Debug(ctx, "Packing list of primitives for field ChangeProtection")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.ChangeProtection, d = basetypes.NewListValueFrom(ctx, elemType, sdk.ChangeProtection)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.ChangeProtection = basetypes.NewListNull(elemType)
	}
	// Handling Lists
	if sdk.CustomerZoneIdList != nil {
		tflog.Debug(ctx, "Packing list of primitives for field CustomerZoneIdList")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.CustomerZoneIdList, d = basetypes.NewListValueFrom(ctx, elemType, sdk.CustomerZoneIdList)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.CustomerZoneIdList = basetypes.NewListNull(elemType)
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Description != nil {
		model.Description = basetypes.NewStringValue(*sdk.Description)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Description", "value": *sdk.Description})
	} else {
		model.Description = basetypes.NewStringNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.EgressNAT != nil {
		tflog.Debug(ctx, "Packing nested object for field EgressNAT")
		packed, d := packNgfirewallEgressNatConfigFromSdk(ctx, *sdk.EgressNAT)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "EgressNAT"})
		}
		model.EgressNAT = packed
	} else {
		model.EgressNAT = basetypes.NewObjectNull(models.NgfirewallEgressNatConfig{}.AttrTypes())
	}
	// Handling Lists
	if sdk.Endpoints != nil {
		tflog.Debug(ctx, "Packing list of objects for field Endpoints")
		packed, d := packNgfirewallEndpointListFromSdk(ctx, sdk.Endpoints)
		diags.Append(d...)
		model.Endpoints = packed
	} else {
		model.Endpoints = basetypes.NewListNull(models.NgfirewallEndpoint{}.AttrType())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.GwlbTcpIdleTimeout != nil {
		model.GwlbTcpIdleTimeout = basetypes.NewInt64Value(int64(*sdk.GwlbTcpIdleTimeout))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "GwlbTcpIdleTimeout", "value": *sdk.GwlbTcpIdleTimeout})
	} else {
		model.GwlbTcpIdleTimeout = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.LinkId != nil {
		model.LinkId = basetypes.NewStringValue(*sdk.LinkId)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "LinkId", "value": *sdk.LinkId})
	} else {
		model.LinkId = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Mode != nil {
		model.Mode = basetypes.NewStringValue(*sdk.Mode)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Mode", "value": *sdk.Mode})
	} else {
		model.Mode = basetypes.NewStringNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.PrivateAccess != nil {
		tflog.Debug(ctx, "Packing nested object for field PrivateAccess")
		packed, d := packNgfirewallPrivateAccessFromSdk(ctx, *sdk.PrivateAccess)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "PrivateAccess"})
		}
		model.PrivateAccess = packed
	} else {
		model.PrivateAccess = basetypes.NewObjectNull(models.NgfirewallPrivateAccess{}.AttrTypes())
	}
	// Handling Lists
	if sdk.ScmLabels != nil {
		tflog.Debug(ctx, "Packing list of primitives for field ScmLabels")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.ScmLabels, d = basetypes.NewListValueFrom(ctx, elemType, sdk.ScmLabels)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.ScmLabels = basetypes.NewListNull(elemType)
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Source != nil {
		model.Source = basetypes.NewStringValue(*sdk.Source)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Source", "value": *sdk.Source})
	} else {
		model.Source = basetypes.NewStringNull()
	}
	// Handling Lists
	if sdk.Tags != nil {
		tflog.Debug(ctx, "Packing list of objects for field Tags")
		packed, d := packNgfirewallTagListFromSdk(ctx, sdk.Tags)
		diags.Append(d...)
		model.Tags = packed
	} else {
		model.Tags = basetypes.NewListNull(models.NgfirewallTag{}.AttrType())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Tier != nil {
		model.Tier = basetypes.NewStringValue(*sdk.Tier)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Tier", "value": *sdk.Tier})
	} else {
		model.Tier = basetypes.NewStringNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.UserID != nil {
		tflog.Debug(ctx, "Packing nested object for field UserID")
		packed, d := packNgfirewallUserIdFromSdk(ctx, *sdk.UserID)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "UserID"})
		}
		model.UserID = packed
	} else {
		model.UserID = basetypes.NewObjectNull(models.NgfirewallUserId{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.NgfirewallCreateRequest{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.NgfirewallCreateRequest", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for NgfirewallCreateRequest ---
func unpackNgfirewallCreateRequestListToSdk(ctx context.Context, list types.List) ([]cloud_ngfw.NgfirewallCreateRequest, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.NgfirewallCreateRequest")
	diags := diag.Diagnostics{}
	var data []models.NgfirewallCreateRequest
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]cloud_ngfw.NgfirewallCreateRequest, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.NgfirewallCreateRequest{}.AttrTypes(), &item)
		unpacked, d := unpackNgfirewallCreateRequestToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.NgfirewallCreateRequest", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for NgfirewallCreateRequest ---
func packNgfirewallCreateRequestListFromSdk(ctx context.Context, sdks []cloud_ngfw.NgfirewallCreateRequest) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.NgfirewallCreateRequest")
	diags := diag.Diagnostics{}
	var data []models.NgfirewallCreateRequest

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.NgfirewallCreateRequest
		obj, d := packNgfirewallCreateRequestFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.NgfirewallCreateRequest{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.NgfirewallCreateRequest", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.NgfirewallCreateRequest{}.AttrType(), data)
}

// --- Unpacker for NgfirewallsListResponse ---
func unpackNgfirewallsListResponseToSdk(ctx context.Context, obj types.Object) (*cloud_ngfw.NgfirewallsListResponse, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.NgfirewallsListResponse", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.NgfirewallsListResponse
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk cloud_ngfw.NgfirewallsListResponse
	var d diag.Diagnostics

	// Handling Objects
	if !model.Response.IsNull() && !model.Response.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Response")
		unpacked, d := unpackNgfirewallsListResponseBodyToSdk(ctx, model.Response)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Response"})
		}
		if unpacked != nil {
			sdk.Response = unpacked
		}
	}

	// Handling Objects
	if !model.ResponseStatus.IsNull() && !model.ResponseStatus.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field ResponseStatus")
		unpacked, d := unpackNgfirewallResponseStatusToSdk(ctx, model.ResponseStatus)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "ResponseStatus"})
		}
		if unpacked != nil {
			sdk.ResponseStatus = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.NgfirewallsListResponse", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for NgfirewallsListResponse ---
func packNgfirewallsListResponseFromSdk(ctx context.Context, sdk cloud_ngfw.NgfirewallsListResponse) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.NgfirewallsListResponse", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.NgfirewallsListResponse
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Response != nil {
		tflog.Debug(ctx, "Packing nested object for field Response")
		packed, d := packNgfirewallsListResponseBodyFromSdk(ctx, *sdk.Response)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Response"})
		}
		model.Response = packed
	} else {
		model.Response = basetypes.NewObjectNull(models.NgfirewallsListResponseBody{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.ResponseStatus != nil {
		tflog.Debug(ctx, "Packing nested object for field ResponseStatus")
		packed, d := packNgfirewallResponseStatusFromSdk(ctx, *sdk.ResponseStatus)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "ResponseStatus"})
		}
		model.ResponseStatus = packed
	} else {
		model.ResponseStatus = basetypes.NewObjectNull(models.NgfirewallResponseStatus{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.NgfirewallsListResponse{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.NgfirewallsListResponse", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for NgfirewallsListResponse ---
func unpackNgfirewallsListResponseListToSdk(ctx context.Context, list types.List) ([]cloud_ngfw.NgfirewallsListResponse, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.NgfirewallsListResponse")
	diags := diag.Diagnostics{}
	var data []models.NgfirewallsListResponse
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]cloud_ngfw.NgfirewallsListResponse, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.NgfirewallsListResponse{}.AttrTypes(), &item)
		unpacked, d := unpackNgfirewallsListResponseToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.NgfirewallsListResponse", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for NgfirewallsListResponse ---
func packNgfirewallsListResponseListFromSdk(ctx context.Context, sdks []cloud_ngfw.NgfirewallsListResponse) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.NgfirewallsListResponse")
	diags := diag.Diagnostics{}
	var data []models.NgfirewallsListResponse

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.NgfirewallsListResponse
		obj, d := packNgfirewallsListResponseFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.NgfirewallsListResponse{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.NgfirewallsListResponse", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.NgfirewallsListResponse{}.AttrType(), data)
}

// --- Unpacker for NgfirewallsListResponseBody ---
func unpackNgfirewallsListResponseBodyToSdk(ctx context.Context, obj types.Object) (*cloud_ngfw.NgfirewallsListResponseBody, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.NgfirewallsListResponseBody", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.NgfirewallsListResponseBody
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk cloud_ngfw.NgfirewallsListResponseBody
	var d diag.Diagnostics
	// Handling Lists
	if !model.FirewallsDescribe.IsNull() && !model.FirewallsDescribe.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field FirewallsDescribe")
		unpacked, d := unpackNgfirewallResponseBodyListToSdk(ctx, model.FirewallsDescribe)
		diags.Append(d...)
		sdk.FirewallsDescribe = unpacked
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.NgfirewallsListResponseBody", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for NgfirewallsListResponseBody ---
func packNgfirewallsListResponseBodyFromSdk(ctx context.Context, sdk cloud_ngfw.NgfirewallsListResponseBody) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.NgfirewallsListResponseBody", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.NgfirewallsListResponseBody
	var d diag.Diagnostics
	// Handling Lists
	if sdk.FirewallsDescribe != nil {
		tflog.Debug(ctx, "Packing list of objects for field FirewallsDescribe")
		packed, d := packNgfirewallResponseBodyListFromSdk(ctx, sdk.FirewallsDescribe)
		diags.Append(d...)
		model.FirewallsDescribe = packed
	} else {
		model.FirewallsDescribe = basetypes.NewListNull(models.NgfirewallResponseBody{}.AttrType())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.NgfirewallsListResponseBody{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.NgfirewallsListResponseBody", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for NgfirewallsListResponseBody ---
func unpackNgfirewallsListResponseBodyListToSdk(ctx context.Context, list types.List) ([]cloud_ngfw.NgfirewallsListResponseBody, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.NgfirewallsListResponseBody")
	diags := diag.Diagnostics{}
	var data []models.NgfirewallsListResponseBody
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]cloud_ngfw.NgfirewallsListResponseBody, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.NgfirewallsListResponseBody{}.AttrTypes(), &item)
		unpacked, d := unpackNgfirewallsListResponseBodyToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.NgfirewallsListResponseBody", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for NgfirewallsListResponseBody ---
func packNgfirewallsListResponseBodyListFromSdk(ctx context.Context, sdks []cloud_ngfw.NgfirewallsListResponseBody) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.NgfirewallsListResponseBody")
	diags := diag.Diagnostics{}
	var data []models.NgfirewallsListResponseBody

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.NgfirewallsListResponseBody
		obj, d := packNgfirewallsListResponseBodyFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.NgfirewallsListResponseBody{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.NgfirewallsListResponseBody", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.NgfirewallsListResponseBody{}.AttrType(), data)
}
