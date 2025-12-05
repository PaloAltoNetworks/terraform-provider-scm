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

// --- Unpacker for ServiceRoute ---
func unpackServiceRouteToSdk(ctx context.Context, obj types.Object) (*device_settings.ServiceRoute, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ServiceRoute", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ServiceRoute
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk device_settings.ServiceRoute
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
	if !model.Route.IsNull() && !model.Route.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Route")
		unpacked, d := unpackServiceRouteRouteToSdk(ctx, model.Route)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Route"})
		}
		if unpacked != nil {
			sdk.Route = unpacked
		}
	}

	// Handling Primitives
	if !model.Snippet.IsNull() && !model.Snippet.IsUnknown() {
		sdk.Snippet = model.Snippet.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Snippet", "value": *sdk.Snippet})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ServiceRoute", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ServiceRoute ---
func packServiceRouteFromSdk(ctx context.Context, sdk device_settings.ServiceRoute) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ServiceRoute", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ServiceRoute
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
	if sdk.Route != nil {
		tflog.Debug(ctx, "Packing nested object for field Route")
		packed, d := packServiceRouteRouteFromSdk(ctx, *sdk.Route)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Route"})
		}
		model.Route = packed
	} else {
		model.Route = basetypes.NewObjectNull(models.ServiceRouteRoute{}.AttrTypes())
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

	obj, d := types.ObjectValueFrom(ctx, models.ServiceRoute{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ServiceRoute", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ServiceRoute ---
func unpackServiceRouteListToSdk(ctx context.Context, list types.List) ([]device_settings.ServiceRoute, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ServiceRoute")
	diags := diag.Diagnostics{}
	var data []models.ServiceRoute
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]device_settings.ServiceRoute, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ServiceRoute{}.AttrTypes(), &item)
		unpacked, d := unpackServiceRouteToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ServiceRoute", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ServiceRoute ---
func packServiceRouteListFromSdk(ctx context.Context, sdks []device_settings.ServiceRoute) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ServiceRoute")
	diags := diag.Diagnostics{}
	var data []models.ServiceRoute

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ServiceRoute
		obj, d := packServiceRouteFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ServiceRoute{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ServiceRoute", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ServiceRoute{}.AttrType(), data)
}

// --- Unpacker for ServiceRouteRoute ---
func unpackServiceRouteRouteToSdk(ctx context.Context, obj types.Object) (*device_settings.ServiceRouteRoute, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ServiceRouteRoute", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ServiceRouteRoute
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk device_settings.ServiceRouteRoute
	var d diag.Diagnostics
	// Handling Lists
	if !model.Destination.IsNull() && !model.Destination.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field Destination")
		unpacked, d := unpackServiceRouteRouteDestinationInnerListToSdk(ctx, model.Destination)
		diags.Append(d...)
		sdk.Destination = unpacked
	}

	// Handling Lists
	if !model.Service.IsNull() && !model.Service.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field Service")
		unpacked, d := unpackServiceRouteRouteServiceInnerListToSdk(ctx, model.Service)
		diags.Append(d...)
		sdk.Service = unpacked
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ServiceRouteRoute", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ServiceRouteRoute ---
func packServiceRouteRouteFromSdk(ctx context.Context, sdk device_settings.ServiceRouteRoute) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ServiceRouteRoute", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ServiceRouteRoute
	var d diag.Diagnostics
	// Handling Lists
	if sdk.Destination != nil {
		tflog.Debug(ctx, "Packing list of objects for field Destination")
		packed, d := packServiceRouteRouteDestinationInnerListFromSdk(ctx, sdk.Destination)
		diags.Append(d...)
		model.Destination = packed
	} else {
		model.Destination = basetypes.NewListNull(models.ServiceRouteRouteDestinationInner{}.AttrType())
	}
	// Handling Lists
	if sdk.Service != nil {
		tflog.Debug(ctx, "Packing list of objects for field Service")
		packed, d := packServiceRouteRouteServiceInnerListFromSdk(ctx, sdk.Service)
		diags.Append(d...)
		model.Service = packed
	} else {
		model.Service = basetypes.NewListNull(models.ServiceRouteRouteServiceInner{}.AttrType())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ServiceRouteRoute{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ServiceRouteRoute", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ServiceRouteRoute ---
func unpackServiceRouteRouteListToSdk(ctx context.Context, list types.List) ([]device_settings.ServiceRouteRoute, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ServiceRouteRoute")
	diags := diag.Diagnostics{}
	var data []models.ServiceRouteRoute
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]device_settings.ServiceRouteRoute, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ServiceRouteRoute{}.AttrTypes(), &item)
		unpacked, d := unpackServiceRouteRouteToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ServiceRouteRoute", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ServiceRouteRoute ---
func packServiceRouteRouteListFromSdk(ctx context.Context, sdks []device_settings.ServiceRouteRoute) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ServiceRouteRoute")
	diags := diag.Diagnostics{}
	var data []models.ServiceRouteRoute

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ServiceRouteRoute
		obj, d := packServiceRouteRouteFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ServiceRouteRoute{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ServiceRouteRoute", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ServiceRouteRoute{}.AttrType(), data)
}

// --- Unpacker for ServiceRouteRouteDestinationInner ---
func unpackServiceRouteRouteDestinationInnerToSdk(ctx context.Context, obj types.Object) (*device_settings.ServiceRouteRouteDestinationInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ServiceRouteRouteDestinationInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ServiceRouteRouteDestinationInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk device_settings.ServiceRouteRouteDestinationInner
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
	}

	// Handling Objects
	if !model.Source.IsNull() && !model.Source.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Source")
		unpacked, d := unpackServiceRouteRouteDestinationInnerSourceToSdk(ctx, model.Source)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Source"})
		}
		if unpacked != nil {
			sdk.Source = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ServiceRouteRouteDestinationInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ServiceRouteRouteDestinationInner ---
func packServiceRouteRouteDestinationInnerFromSdk(ctx context.Context, sdk device_settings.ServiceRouteRouteDestinationInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ServiceRouteRouteDestinationInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ServiceRouteRouteDestinationInner
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Name != nil {
		model.Name = basetypes.NewStringValue(*sdk.Name)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
	} else {
		model.Name = basetypes.NewStringNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Source != nil {
		tflog.Debug(ctx, "Packing nested object for field Source")
		packed, d := packServiceRouteRouteDestinationInnerSourceFromSdk(ctx, *sdk.Source)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Source"})
		}
		model.Source = packed
	} else {
		model.Source = basetypes.NewObjectNull(models.ServiceRouteRouteDestinationInnerSource{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ServiceRouteRouteDestinationInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ServiceRouteRouteDestinationInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ServiceRouteRouteDestinationInner ---
func unpackServiceRouteRouteDestinationInnerListToSdk(ctx context.Context, list types.List) ([]device_settings.ServiceRouteRouteDestinationInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ServiceRouteRouteDestinationInner")
	diags := diag.Diagnostics{}
	var data []models.ServiceRouteRouteDestinationInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]device_settings.ServiceRouteRouteDestinationInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ServiceRouteRouteDestinationInner{}.AttrTypes(), &item)
		unpacked, d := unpackServiceRouteRouteDestinationInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ServiceRouteRouteDestinationInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ServiceRouteRouteDestinationInner ---
func packServiceRouteRouteDestinationInnerListFromSdk(ctx context.Context, sdks []device_settings.ServiceRouteRouteDestinationInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ServiceRouteRouteDestinationInner")
	diags := diag.Diagnostics{}
	var data []models.ServiceRouteRouteDestinationInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ServiceRouteRouteDestinationInner
		obj, d := packServiceRouteRouteDestinationInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ServiceRouteRouteDestinationInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ServiceRouteRouteDestinationInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ServiceRouteRouteDestinationInner{}.AttrType(), data)
}

// --- Unpacker for ServiceRouteRouteDestinationInnerSource ---
func unpackServiceRouteRouteDestinationInnerSourceToSdk(ctx context.Context, obj types.Object) (*device_settings.ServiceRouteRouteDestinationInnerSource, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ServiceRouteRouteDestinationInnerSource", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ServiceRouteRouteDestinationInnerSource
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk device_settings.ServiceRouteRouteDestinationInnerSource
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Address.IsNull() && !model.Address.IsUnknown() {
		sdk.Address = model.Address.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Address", "value": *sdk.Address})
	}

	// Handling Primitives
	if !model.Interface.IsNull() && !model.Interface.IsUnknown() {
		sdk.Interface = model.Interface.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Interface", "value": *sdk.Interface})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ServiceRouteRouteDestinationInnerSource", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ServiceRouteRouteDestinationInnerSource ---
func packServiceRouteRouteDestinationInnerSourceFromSdk(ctx context.Context, sdk device_settings.ServiceRouteRouteDestinationInnerSource) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ServiceRouteRouteDestinationInnerSource", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ServiceRouteRouteDestinationInnerSource
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Address != nil {
		model.Address = basetypes.NewStringValue(*sdk.Address)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Address", "value": *sdk.Address})
	} else {
		model.Address = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Interface != nil {
		model.Interface = basetypes.NewStringValue(*sdk.Interface)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Interface", "value": *sdk.Interface})
	} else {
		model.Interface = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ServiceRouteRouteDestinationInnerSource{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ServiceRouteRouteDestinationInnerSource", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ServiceRouteRouteDestinationInnerSource ---
func unpackServiceRouteRouteDestinationInnerSourceListToSdk(ctx context.Context, list types.List) ([]device_settings.ServiceRouteRouteDestinationInnerSource, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ServiceRouteRouteDestinationInnerSource")
	diags := diag.Diagnostics{}
	var data []models.ServiceRouteRouteDestinationInnerSource
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]device_settings.ServiceRouteRouteDestinationInnerSource, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ServiceRouteRouteDestinationInnerSource{}.AttrTypes(), &item)
		unpacked, d := unpackServiceRouteRouteDestinationInnerSourceToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ServiceRouteRouteDestinationInnerSource", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ServiceRouteRouteDestinationInnerSource ---
func packServiceRouteRouteDestinationInnerSourceListFromSdk(ctx context.Context, sdks []device_settings.ServiceRouteRouteDestinationInnerSource) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ServiceRouteRouteDestinationInnerSource")
	diags := diag.Diagnostics{}
	var data []models.ServiceRouteRouteDestinationInnerSource

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ServiceRouteRouteDestinationInnerSource
		obj, d := packServiceRouteRouteDestinationInnerSourceFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ServiceRouteRouteDestinationInnerSource{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ServiceRouteRouteDestinationInnerSource", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ServiceRouteRouteDestinationInnerSource{}.AttrType(), data)
}

// --- Unpacker for ServiceRouteRouteServiceInner ---
func unpackServiceRouteRouteServiceInnerToSdk(ctx context.Context, obj types.Object) (*device_settings.ServiceRouteRouteServiceInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ServiceRouteRouteServiceInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ServiceRouteRouteServiceInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk device_settings.ServiceRouteRouteServiceInner
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
	}

	// Handling Objects
	if !model.Source.IsNull() && !model.Source.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Source")
		unpacked, d := unpackServiceRouteRouteServiceInnerSourceToSdk(ctx, model.Source)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Source"})
		}
		if unpacked != nil {
			sdk.Source = unpacked
		}
	}

	// Handling Objects
	if !model.SourceV6.IsNull() && !model.SourceV6.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field SourceV6")
		unpacked, d := unpackServiceRouteRouteServiceInnerSourceV6ToSdk(ctx, model.SourceV6)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "SourceV6"})
		}
		if unpacked != nil {
			sdk.SourceV6 = unpacked
		}
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ServiceRouteRouteServiceInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ServiceRouteRouteServiceInner ---
func packServiceRouteRouteServiceInnerFromSdk(ctx context.Context, sdk device_settings.ServiceRouteRouteServiceInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ServiceRouteRouteServiceInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ServiceRouteRouteServiceInner
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Name != nil {
		model.Name = basetypes.NewStringValue(*sdk.Name)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
	} else {
		model.Name = basetypes.NewStringNull()
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Source != nil {
		tflog.Debug(ctx, "Packing nested object for field Source")
		packed, d := packServiceRouteRouteServiceInnerSourceFromSdk(ctx, *sdk.Source)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Source"})
		}
		model.Source = packed
	} else {
		model.Source = basetypes.NewObjectNull(models.ServiceRouteRouteServiceInnerSource{}.AttrTypes())
	}
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.SourceV6 != nil {
		tflog.Debug(ctx, "Packing nested object for field SourceV6")
		packed, d := packServiceRouteRouteServiceInnerSourceV6FromSdk(ctx, *sdk.SourceV6)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "SourceV6"})
		}
		model.SourceV6 = packed
	} else {
		model.SourceV6 = basetypes.NewObjectNull(models.ServiceRouteRouteServiceInnerSourceV6{}.AttrTypes())
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ServiceRouteRouteServiceInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ServiceRouteRouteServiceInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ServiceRouteRouteServiceInner ---
func unpackServiceRouteRouteServiceInnerListToSdk(ctx context.Context, list types.List) ([]device_settings.ServiceRouteRouteServiceInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ServiceRouteRouteServiceInner")
	diags := diag.Diagnostics{}
	var data []models.ServiceRouteRouteServiceInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]device_settings.ServiceRouteRouteServiceInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ServiceRouteRouteServiceInner{}.AttrTypes(), &item)
		unpacked, d := unpackServiceRouteRouteServiceInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ServiceRouteRouteServiceInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ServiceRouteRouteServiceInner ---
func packServiceRouteRouteServiceInnerListFromSdk(ctx context.Context, sdks []device_settings.ServiceRouteRouteServiceInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ServiceRouteRouteServiceInner")
	diags := diag.Diagnostics{}
	var data []models.ServiceRouteRouteServiceInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ServiceRouteRouteServiceInner
		obj, d := packServiceRouteRouteServiceInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ServiceRouteRouteServiceInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ServiceRouteRouteServiceInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ServiceRouteRouteServiceInner{}.AttrType(), data)
}

// --- Unpacker for ServiceRouteRouteServiceInnerSource ---
func unpackServiceRouteRouteServiceInnerSourceToSdk(ctx context.Context, obj types.Object) (*device_settings.ServiceRouteRouteServiceInnerSource, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ServiceRouteRouteServiceInnerSource", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ServiceRouteRouteServiceInnerSource
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk device_settings.ServiceRouteRouteServiceInnerSource
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Address.IsNull() && !model.Address.IsUnknown() {
		sdk.Address = model.Address.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Address", "value": *sdk.Address})
	}

	// Handling Primitives
	if !model.Interface.IsNull() && !model.Interface.IsUnknown() {
		sdk.Interface = model.Interface.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Interface", "value": *sdk.Interface})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ServiceRouteRouteServiceInnerSource", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ServiceRouteRouteServiceInnerSource ---
func packServiceRouteRouteServiceInnerSourceFromSdk(ctx context.Context, sdk device_settings.ServiceRouteRouteServiceInnerSource) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ServiceRouteRouteServiceInnerSource", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ServiceRouteRouteServiceInnerSource
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Address != nil {
		model.Address = basetypes.NewStringValue(*sdk.Address)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Address", "value": *sdk.Address})
	} else {
		model.Address = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Interface != nil {
		model.Interface = basetypes.NewStringValue(*sdk.Interface)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Interface", "value": *sdk.Interface})
	} else {
		model.Interface = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ServiceRouteRouteServiceInnerSource{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ServiceRouteRouteServiceInnerSource", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ServiceRouteRouteServiceInnerSource ---
func unpackServiceRouteRouteServiceInnerSourceListToSdk(ctx context.Context, list types.List) ([]device_settings.ServiceRouteRouteServiceInnerSource, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ServiceRouteRouteServiceInnerSource")
	diags := diag.Diagnostics{}
	var data []models.ServiceRouteRouteServiceInnerSource
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]device_settings.ServiceRouteRouteServiceInnerSource, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ServiceRouteRouteServiceInnerSource{}.AttrTypes(), &item)
		unpacked, d := unpackServiceRouteRouteServiceInnerSourceToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ServiceRouteRouteServiceInnerSource", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ServiceRouteRouteServiceInnerSource ---
func packServiceRouteRouteServiceInnerSourceListFromSdk(ctx context.Context, sdks []device_settings.ServiceRouteRouteServiceInnerSource) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ServiceRouteRouteServiceInnerSource")
	diags := diag.Diagnostics{}
	var data []models.ServiceRouteRouteServiceInnerSource

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ServiceRouteRouteServiceInnerSource
		obj, d := packServiceRouteRouteServiceInnerSourceFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ServiceRouteRouteServiceInnerSource{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ServiceRouteRouteServiceInnerSource", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ServiceRouteRouteServiceInnerSource{}.AttrType(), data)
}

// --- Unpacker for ServiceRouteRouteServiceInnerSourceV6 ---
func unpackServiceRouteRouteServiceInnerSourceV6ToSdk(ctx context.Context, obj types.Object) (*device_settings.ServiceRouteRouteServiceInnerSourceV6, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.ServiceRouteRouteServiceInnerSourceV6", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.ServiceRouteRouteServiceInnerSourceV6
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk device_settings.ServiceRouteRouteServiceInnerSourceV6
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Address.IsNull() && !model.Address.IsUnknown() {
		sdk.Address = model.Address.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Address", "value": *sdk.Address})
	}

	// Handling Primitives
	if !model.Interface.IsNull() && !model.Interface.IsUnknown() {
		sdk.Interface = model.Interface.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Interface", "value": *sdk.Interface})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.ServiceRouteRouteServiceInnerSourceV6", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for ServiceRouteRouteServiceInnerSourceV6 ---
func packServiceRouteRouteServiceInnerSourceV6FromSdk(ctx context.Context, sdk device_settings.ServiceRouteRouteServiceInnerSourceV6) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.ServiceRouteRouteServiceInnerSourceV6", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.ServiceRouteRouteServiceInnerSourceV6
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.Address != nil {
		model.Address = basetypes.NewStringValue(*sdk.Address)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Address", "value": *sdk.Address})
	} else {
		model.Address = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Interface != nil {
		model.Interface = basetypes.NewStringValue(*sdk.Interface)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Interface", "value": *sdk.Interface})
	} else {
		model.Interface = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.ServiceRouteRouteServiceInnerSourceV6{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.ServiceRouteRouteServiceInnerSourceV6", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for ServiceRouteRouteServiceInnerSourceV6 ---
func unpackServiceRouteRouteServiceInnerSourceV6ListToSdk(ctx context.Context, list types.List) ([]device_settings.ServiceRouteRouteServiceInnerSourceV6, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.ServiceRouteRouteServiceInnerSourceV6")
	diags := diag.Diagnostics{}
	var data []models.ServiceRouteRouteServiceInnerSourceV6
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]device_settings.ServiceRouteRouteServiceInnerSourceV6, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.ServiceRouteRouteServiceInnerSourceV6{}.AttrTypes(), &item)
		unpacked, d := unpackServiceRouteRouteServiceInnerSourceV6ToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.ServiceRouteRouteServiceInnerSourceV6", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for ServiceRouteRouteServiceInnerSourceV6 ---
func packServiceRouteRouteServiceInnerSourceV6ListFromSdk(ctx context.Context, sdks []device_settings.ServiceRouteRouteServiceInnerSourceV6) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.ServiceRouteRouteServiceInnerSourceV6")
	diags := diag.Diagnostics{}
	var data []models.ServiceRouteRouteServiceInnerSourceV6

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.ServiceRouteRouteServiceInnerSourceV6
		obj, d := packServiceRouteRouteServiceInnerSourceV6FromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.ServiceRouteRouteServiceInnerSourceV6{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.ServiceRouteRouteServiceInnerSourceV6", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.ServiceRouteRouteServiceInnerSourceV6{}.AttrType(), data)
}
