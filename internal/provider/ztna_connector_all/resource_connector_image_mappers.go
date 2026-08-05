package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/ztna_connector_all"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/ztna_connector_all"
)

// --- Unpacker for ConnectorImages ---
func unpackConnectorImagesToSdk(ctx context.Context, obj types.Object) (*ztna_connector_all.ConnectorImages, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ConnectorImages", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ConnectorImages
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk ztna_connector_all.ConnectorImages
	var d diag.Diagnostics

	// Handling Primitives
	if !model.Id.IsNull() && !model.Id.IsUnknown() {
		sdk.Id = model.Id.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Id", "value": *sdk.Id})
	}

	// Handling Primitives
	if !model.Version.IsNull() && !model.Version.IsUnknown() {
		sdk.Version = model.Version.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Version", "value": *sdk.Version})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ConnectorImages", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ConnectorImages ---
func packConnectorImagesFromSdk(ctx context.Context, sdk ztna_connector_all.ConnectorImages) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ConnectorImages", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ConnectorImages
	var d diag.Diagnostics
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
	if sdk.Version != nil {
		model.Version = basetypes.NewStringValue(*sdk.Version)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Version", "value": *sdk.Version})
	} else {
		model.Version = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ConnectorImages{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ConnectorImages", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ConnectorImages ---
func unpackConnectorImagesListToSdk(ctx context.Context, list types.List) ([]ztna_connector_all.ConnectorImages, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ConnectorImages")
	diags := diag.Diagnostics{}
	var data []models.ConnectorImages
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]ztna_connector_all.ConnectorImages, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ConnectorImages{}.AttrTypes(), &item)
		unpacked, d := unpackConnectorImagesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ConnectorImages", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ConnectorImages ---
func packConnectorImagesListFromSdk(ctx context.Context, sdks []ztna_connector_all.ConnectorImages) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ConnectorImages")
	diags := diag.Diagnostics{}
	var data []models.ConnectorImages

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ConnectorImages
		obj, d := packConnectorImagesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ConnectorImages{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ConnectorImages", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ConnectorImages{}.AttrType(), data)
}
