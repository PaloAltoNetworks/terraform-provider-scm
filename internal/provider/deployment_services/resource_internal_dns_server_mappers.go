package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/deployment_services"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/deployment_services"
)

// --- Unpacker for InternalDnsServers ---
func unpackInternalDnsServersToSdk(ctx context.Context, obj types.Object) (*deployment_services.InternalDnsServers, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.InternalDnsServers", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.InternalDnsServers
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk deployment_services.InternalDnsServers
	var d diag.Diagnostics
	// Handling Lists
	if !model.DomainName.IsNull() && !model.DomainName.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field DomainName")
		diags.Append(model.DomainName.ElementsAs(ctx, &sdk.DomainName, false)...)
	}

	// Handling Primitives
	if !model.Id.IsNull() && !model.Id.IsUnknown() {
		sdk.Id = model.Id.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Id", "value": sdk.Id})
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	// Handling Primitives
	if !model.Primary.IsNull() && !model.Primary.IsUnknown() {
		sdk.Primary = model.Primary.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Primary", "value": sdk.Primary})
	}

	// Handling Primitives
	if !model.Secondary.IsNull() && !model.Secondary.IsUnknown() {
		sdk.Secondary = model.Secondary.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Secondary", "value": *sdk.Secondary})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.InternalDnsServers", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for InternalDnsServers ---
func packInternalDnsServersFromSdk(ctx context.Context, sdk deployment_services.InternalDnsServers) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.InternalDnsServers", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.InternalDnsServers
	var d diag.Diagnostics
	// Handling Lists
	if sdk.DomainName != nil {
		tflog.Debug(ctx, "Packing list of primitives for field DomainName")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.DomainName, d = basetypes.NewListValueFrom(ctx, elemType, sdk.DomainName)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.DomainName = basetypes.NewListNull(elemType)
	}
	// Handling Primitives
	// Standard primitive packing
	model.Id = basetypes.NewStringValue(sdk.Id)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Id", "value": sdk.Id})
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	// Handling Primitives
	// Standard primitive packing
	model.Primary = basetypes.NewStringValue(sdk.Primary)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Primary", "value": sdk.Primary})
	// Handling Primitives
	// Standard primitive packing
	if sdk.Secondary != nil {
		model.Secondary = basetypes.NewStringValue(*sdk.Secondary)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Secondary", "value": *sdk.Secondary})
	} else {
		model.Secondary = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.InternalDnsServers{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.InternalDnsServers", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for InternalDnsServers ---
func unpackInternalDnsServersListToSdk(ctx context.Context, list types.List) ([]deployment_services.InternalDnsServers, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.InternalDnsServers")
	diags := diag.Diagnostics{}
	var data []models.InternalDnsServers
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]deployment_services.InternalDnsServers, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.InternalDnsServers{}.AttrTypes(), &item)
		unpacked, d := unpackInternalDnsServersToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.InternalDnsServers", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for InternalDnsServers ---
func packInternalDnsServersListFromSdk(ctx context.Context, sdks []deployment_services.InternalDnsServers) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.InternalDnsServers")
	diags := diag.Diagnostics{}
	var data []models.InternalDnsServers

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.InternalDnsServers
		obj, d := packInternalDnsServersFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.InternalDnsServers{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.InternalDnsServers", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.InternalDnsServers{}.AttrType(), data)
}
