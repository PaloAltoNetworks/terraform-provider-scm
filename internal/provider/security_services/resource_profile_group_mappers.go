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

// --- Unpacker for ProfileGroups ---
func unpackProfileGroupsToSdk(ctx context.Context, obj types.Object) (*security_services.ProfileGroups, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ProfileGroups", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ProfileGroups
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk security_services.ProfileGroups
	var d diag.Diagnostics
	// Handling Lists
	if !model.AiSecurity.IsNull() && !model.AiSecurity.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field AiSecurity")
		diags.Append(model.AiSecurity.ElementsAs(ctx, &sdk.AiSecurity, false)...)
	}

	// Handling Lists
	if !model.DataFiltering.IsNull() && !model.DataFiltering.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field DataFiltering")
		diags.Append(model.DataFiltering.ElementsAs(ctx, &sdk.DataFiltering, false)...)
	}

	// Handling Primitives
	if !model.Device.IsNull() && !model.Device.IsUnknown() {
		sdk.Device = model.Device.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Device", "value": *sdk.Device})
	}

	// Handling Lists
	if !model.DnsSecurity.IsNull() && !model.DnsSecurity.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field DnsSecurity")
		diags.Append(model.DnsSecurity.ElementsAs(ctx, &sdk.DnsSecurity, false)...)
	}

	// Handling Lists
	if !model.FileBlocking.IsNull() && !model.FileBlocking.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field FileBlocking")
		diags.Append(model.FileBlocking.ElementsAs(ctx, &sdk.FileBlocking, false)...)
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

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	// Handling Lists
	if !model.SaasSecurity.IsNull() && !model.SaasSecurity.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field SaasSecurity")
		diags.Append(model.SaasSecurity.ElementsAs(ctx, &sdk.SaasSecurity, false)...)
	}

	// Handling Primitives
	if !model.Snippet.IsNull() && !model.Snippet.IsUnknown() {
		sdk.Snippet = model.Snippet.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	}

	// Handling Lists
	if !model.Spyware.IsNull() && !model.Spyware.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Spyware")
		diags.Append(model.Spyware.ElementsAs(ctx, &sdk.Spyware, false)...)
	}

	// Handling Lists
	if !model.UrlFiltering.IsNull() && !model.UrlFiltering.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field UrlFiltering")
		diags.Append(model.UrlFiltering.ElementsAs(ctx, &sdk.UrlFiltering, false)...)
	}

	// Handling Lists
	if !model.VirusAndWildfireAnalysis.IsNull() && !model.VirusAndWildfireAnalysis.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field VirusAndWildfireAnalysis")
		diags.Append(model.VirusAndWildfireAnalysis.ElementsAs(ctx, &sdk.VirusAndWildfireAnalysis, false)...)
	}

	// Handling Lists
	if !model.Vulnerability.IsNull() && !model.Vulnerability.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of primitives for field Vulnerability")
		diags.Append(model.Vulnerability.ElementsAs(ctx, &sdk.Vulnerability, false)...)
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ProfileGroups", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ProfileGroups ---
func packProfileGroupsFromSdk(ctx context.Context, sdk security_services.ProfileGroups) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ProfileGroups", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ProfileGroups
	var d diag.Diagnostics
	// Handling Lists
	if sdk.AiSecurity != nil {
		tflog.Debug(ctx, "Packing list of primitives for field AiSecurity")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.AiSecurity, d = basetypes.NewListValueFrom(ctx, elemType, sdk.AiSecurity)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.AiSecurity = basetypes.NewListNull(elemType)
	}
	// Handling Lists
	if sdk.DataFiltering != nil {
		tflog.Debug(ctx, "Packing list of primitives for field DataFiltering")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.DataFiltering, d = basetypes.NewListValueFrom(ctx, elemType, sdk.DataFiltering)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.DataFiltering = basetypes.NewListNull(elemType)
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
	if sdk.DnsSecurity != nil {
		tflog.Debug(ctx, "Packing list of primitives for field DnsSecurity")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.DnsSecurity, d = basetypes.NewListValueFrom(ctx, elemType, sdk.DnsSecurity)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.DnsSecurity = basetypes.NewListNull(elemType)
	}
	// Handling Lists
	if sdk.FileBlocking != nil {
		tflog.Debug(ctx, "Packing list of primitives for field FileBlocking")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.FileBlocking, d = basetypes.NewListValueFrom(ctx, elemType, sdk.FileBlocking)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.FileBlocking = basetypes.NewListNull(elemType)
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
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	// Handling Lists
	if sdk.SaasSecurity != nil {
		tflog.Debug(ctx, "Packing list of primitives for field SaasSecurity")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.SaasSecurity, d = basetypes.NewListValueFrom(ctx, elemType, sdk.SaasSecurity)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.SaasSecurity = basetypes.NewListNull(elemType)
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
	if sdk.Spyware != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Spyware")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Spyware, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Spyware)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Spyware = basetypes.NewListNull(elemType)
	}
	// Handling Lists
	if sdk.UrlFiltering != nil {
		tflog.Debug(ctx, "Packing list of primitives for field UrlFiltering")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.UrlFiltering, d = basetypes.NewListValueFrom(ctx, elemType, sdk.UrlFiltering)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.UrlFiltering = basetypes.NewListNull(elemType)
	}
	// Handling Lists
	if sdk.VirusAndWildfireAnalysis != nil {
		tflog.Debug(ctx, "Packing list of primitives for field VirusAndWildfireAnalysis")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.VirusAndWildfireAnalysis, d = basetypes.NewListValueFrom(ctx, elemType, sdk.VirusAndWildfireAnalysis)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.VirusAndWildfireAnalysis = basetypes.NewListNull(elemType)
	}
	// Handling Lists
	if sdk.Vulnerability != nil {
		tflog.Debug(ctx, "Packing list of primitives for field Vulnerability")
		var d diag.Diagnostics
		// This logic now dynamically determines the element type based on the SDK's Go type.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Vulnerability, d = basetypes.NewListValueFrom(ctx, elemType, sdk.Vulnerability)
		diags.Append(d...)
	} else {
		// This logic now creates a correctly typed null list.
		var elemType attr.Type = basetypes.StringType{} // Default to string
		model.Vulnerability = basetypes.NewListNull(elemType)
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ProfileGroups{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ProfileGroups", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ProfileGroups ---
func unpackProfileGroupsListToSdk(ctx context.Context, list types.List) ([]security_services.ProfileGroups, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ProfileGroups")
	diags := diag.Diagnostics{}
	var data []models.ProfileGroups
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]security_services.ProfileGroups, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ProfileGroups{}.AttrTypes(), &item)
		unpacked, d := unpackProfileGroupsToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ProfileGroups", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ProfileGroups ---
func packProfileGroupsListFromSdk(ctx context.Context, sdks []security_services.ProfileGroups) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ProfileGroups")
	diags := diag.Diagnostics{}
	var data []models.ProfileGroups

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ProfileGroups
		obj, d := packProfileGroupsFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ProfileGroups{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ProfileGroups", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ProfileGroups{}.AttrType(), data)
}
