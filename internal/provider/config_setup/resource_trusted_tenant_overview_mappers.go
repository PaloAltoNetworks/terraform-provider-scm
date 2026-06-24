package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/config_setup"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/config_setup"
)

// --- Unpacker for TrustedTenantOverview ---
func unpackTrustedTenantOverviewToSdk(ctx context.Context, obj types.Object) (*config_setup.TrustedTenantOverview, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.TrustedTenantOverview", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.TrustedTenantOverview
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk config_setup.TrustedTenantOverview
	var d diag.Diagnostics

	// Handling Objects
	if !model.Publisher.IsNull() && !model.Publisher.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Publisher")
		unpacked, d := unpackTrustedTenantOverviewPublisherToSdk(ctx, model.Publisher)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Publisher"})
		}
		if unpacked != nil {
			sdk.Publisher = unpacked
		}
	}

	// Handling Objects
	if !model.Subscriber.IsNull() && !model.Subscriber.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Subscriber")
		unpacked, d := unpackTrustedTenantOverviewPublisherToSdk(ctx, model.Subscriber)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Subscriber"})
		}
		if unpacked != nil {
			sdk.Subscriber = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.TrustedTenantOverview", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for TrustedTenantOverview ---
func packTrustedTenantOverviewFromSdk(ctx context.Context, sdk config_setup.TrustedTenantOverview) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.TrustedTenantOverview", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.TrustedTenantOverview
	var d diag.Diagnostics
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Publisher != nil {
		tflog.Debug(ctx, "Packing nested object for field Publisher")
		packed, d := packTrustedTenantOverviewPublisherFromSdk(ctx, *sdk.Publisher)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Publisher"})
		}
		model.Publisher = packed
	} else {
		model.Publisher = basetypes.NewObjectNull(models.TrustedTenantOverviewPublisher{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Subscriber != nil {
		tflog.Debug(ctx, "Packing nested object for field Subscriber")
		packed, d := packTrustedTenantOverviewPublisherFromSdk(ctx, *sdk.Subscriber)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Subscriber"})
		}
		model.Subscriber = packed
	} else {
		model.Subscriber = basetypes.NewObjectNull(models.TrustedTenantOverviewPublisher{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.TrustedTenantOverview{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.TrustedTenantOverview", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for TrustedTenantOverview ---
func unpackTrustedTenantOverviewListToSdk(ctx context.Context, list types.List) ([]config_setup.TrustedTenantOverview, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.TrustedTenantOverview")
	diags := diag.Diagnostics{}
	var data []models.TrustedTenantOverview
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]config_setup.TrustedTenantOverview, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.TrustedTenantOverview{}.AttrTypes(), &item)
		unpacked, d := unpackTrustedTenantOverviewToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.TrustedTenantOverview", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for TrustedTenantOverview ---
func packTrustedTenantOverviewListFromSdk(ctx context.Context, sdks []config_setup.TrustedTenantOverview) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.TrustedTenantOverview")
	diags := diag.Diagnostics{}
	var data []models.TrustedTenantOverview

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.TrustedTenantOverview
		obj, d := packTrustedTenantOverviewFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.TrustedTenantOverview{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.TrustedTenantOverview", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.TrustedTenantOverview{}.AttrType(), data)
}

// --- Unpacker for TrustedTenantOverviewPublisher ---
func unpackTrustedTenantOverviewPublisherToSdk(ctx context.Context, obj types.Object) (*config_setup.TrustedTenantOverviewPublisher, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.TrustedTenantOverviewPublisher", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.TrustedTenantOverviewPublisher
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk config_setup.TrustedTenantOverviewPublisher
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Pending.IsNull() && !model.Pending.IsUnknown() {
		val := int32(model.Pending.ValueInt64())
		sdk.Pending = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Pending", "value": *sdk.Pending})
	}

	// Handling Primitives
	if !model.Total.IsNull() && !model.Total.IsUnknown() {
		val := int32(model.Total.ValueInt64())
		sdk.Total = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Total", "value": *sdk.Total})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.TrustedTenantOverviewPublisher", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for TrustedTenantOverviewPublisher ---
func packTrustedTenantOverviewPublisherFromSdk(ctx context.Context, sdk config_setup.TrustedTenantOverviewPublisher) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.TrustedTenantOverviewPublisher", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.TrustedTenantOverviewPublisher
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Pending != nil {
		model.Pending = basetypes.NewInt64Value(int64(*sdk.Pending))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Pending", "value": *sdk.Pending})
	} else {
		model.Pending = basetypes.NewInt64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Total != nil {
		model.Total = basetypes.NewInt64Value(int64(*sdk.Total))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Total", "value": *sdk.Total})
	} else {
		model.Total = basetypes.NewInt64Null()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.TrustedTenantOverviewPublisher{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.TrustedTenantOverviewPublisher", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for TrustedTenantOverviewPublisher ---
func unpackTrustedTenantOverviewPublisherListToSdk(ctx context.Context, list types.List) ([]config_setup.TrustedTenantOverviewPublisher, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.TrustedTenantOverviewPublisher")
	diags := diag.Diagnostics{}
	var data []models.TrustedTenantOverviewPublisher
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]config_setup.TrustedTenantOverviewPublisher, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.TrustedTenantOverviewPublisher{}.AttrTypes(), &item)
		unpacked, d := unpackTrustedTenantOverviewPublisherToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.TrustedTenantOverviewPublisher", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for TrustedTenantOverviewPublisher ---
func packTrustedTenantOverviewPublisherListFromSdk(ctx context.Context, sdks []config_setup.TrustedTenantOverviewPublisher) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.TrustedTenantOverviewPublisher")
	diags := diag.Diagnostics{}
	var data []models.TrustedTenantOverviewPublisher

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.TrustedTenantOverviewPublisher
		obj, d := packTrustedTenantOverviewPublisherFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.TrustedTenantOverviewPublisher{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.TrustedTenantOverviewPublisher", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.TrustedTenantOverviewPublisher{}.AttrType(), data)
}
