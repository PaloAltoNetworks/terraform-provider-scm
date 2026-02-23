package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/deployment_services"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/deployment_services"
)

// --- Unpacker for Sites ---
func unpackSitesToSdk(ctx context.Context, obj types.Object) (*deployment_services.Sites, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.Sites", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.Sites
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk deployment_services.Sites
	var d diag.Diagnostics

	// Handling Primitives
	if !model.AddressLine1.IsNull() && !model.AddressLine1.IsUnknown() {
		sdk.AddressLine1 = model.AddressLine1.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "AddressLine1", "value": *sdk.AddressLine1})
	}

	// Handling Primitives
	if !model.AddressLine2.IsNull() && !model.AddressLine2.IsUnknown() {
		sdk.AddressLine2 = model.AddressLine2.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "AddressLine2", "value": *sdk.AddressLine2})
	}

	// Handling Primitives
	if !model.City.IsNull() && !model.City.IsUnknown() {
		sdk.City = model.City.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "City", "value": *sdk.City})
	}

	// Handling Primitives
	if !model.Country.IsNull() && !model.Country.IsUnknown() {
		sdk.Country = model.Country.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Country", "value": *sdk.Country})
	}

	// Handling Primitives
	if !model.Id.IsNull() && !model.Id.IsUnknown() {
		sdk.Id = model.Id.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Id", "value": *sdk.Id})
	}

	// Handling Primitives
	if !model.Latitude.IsNull() && !model.Latitude.IsUnknown() {
		sdk.Latitude = model.Latitude.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Latitude", "value": *sdk.Latitude})
	}

	// Handling Primitives
	if !model.LicenseType.IsNull() && !model.LicenseType.IsUnknown() {
		sdk.LicenseType = model.LicenseType.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "LicenseType", "value": *sdk.LicenseType})
	}

	// Handling Primitives
	if !model.Longitude.IsNull() && !model.Longitude.IsUnknown() {
		sdk.Longitude = model.Longitude.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Longitude", "value": *sdk.Longitude})
	}

	// Handling Lists
	if !model.Members.IsNull() && !model.Members.IsUnknown() {
		tflog.Debug(ctx, "Unpacking list of objects for field Members")
		unpacked, d := unpackSitesMembersInnerListToSdk(ctx, model.Members)
		diags.Append(d...)
		sdk.Members = unpacked
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	// Handling Objects
	if !model.Qos.IsNull() && !model.Qos.IsUnknown() {
		tflog.Debug(ctx, "Unpacking nested object for field Qos")
		unpacked, d := unpackSitesQosToSdk(ctx, model.Qos)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "Qos"})
		}
		if unpacked != nil {
			sdk.Qos = unpacked
		}
	}

	// Handling Primitives
	if !model.State.IsNull() && !model.State.IsUnknown() {
		sdk.State = model.State.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "State", "value": *sdk.State})
	}

	// Handling Primitives
	if !model.Type.IsNull() && !model.Type.IsUnknown() {
		sdk.Type = model.Type.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Type", "value": *sdk.Type})
	}

	// Handling Primitives
	if !model.ZipCode.IsNull() && !model.ZipCode.IsUnknown() {
		sdk.ZipCode = model.ZipCode.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "ZipCode", "value": *sdk.ZipCode})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.Sites", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for Sites ---
func packSitesFromSdk(ctx context.Context, sdk deployment_services.Sites) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.Sites", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.Sites
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.AddressLine1 != nil {
		model.AddressLine1 = basetypes.NewStringValue(*sdk.AddressLine1)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "AddressLine1", "value": *sdk.AddressLine1})
	} else {
		model.AddressLine1 = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.AddressLine2 != nil {
		model.AddressLine2 = basetypes.NewStringValue(*sdk.AddressLine2)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "AddressLine2", "value": *sdk.AddressLine2})
	} else {
		model.AddressLine2 = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.City != nil {
		model.City = basetypes.NewStringValue(*sdk.City)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "City", "value": *sdk.City})
	} else {
		model.City = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Country != nil {
		model.Country = basetypes.NewStringValue(*sdk.Country)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Country", "value": *sdk.Country})
	} else {
		model.Country = basetypes.NewStringNull()
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
	if sdk.Latitude != nil {
		model.Latitude = basetypes.NewStringValue(*sdk.Latitude)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Latitude", "value": *sdk.Latitude})
	} else {
		model.Latitude = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.LicenseType != nil {
		model.LicenseType = basetypes.NewStringValue(*sdk.LicenseType)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "LicenseType", "value": *sdk.LicenseType})
	} else {
		model.LicenseType = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Longitude != nil {
		model.Longitude = basetypes.NewStringValue(*sdk.Longitude)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Longitude", "value": *sdk.Longitude})
	} else {
		model.Longitude = basetypes.NewStringNull()
	}
	// Handling Lists
	if sdk.Members != nil {
		tflog.Debug(ctx, "Packing list of objects for field Members")
		packed, d := packSitesMembersInnerListFromSdk(ctx, sdk.Members)
		diags.Append(d...)
		model.Members = packed
	} else {
		model.Members = basetypes.NewListNull(models.SitesMembersInner{}.AttrType())
	}
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	// Handling Objects
	// This is a regular nested object that has its own packer.
	if sdk.Qos != nil {
		tflog.Debug(ctx, "Packing nested object for field Qos")
		packed, d := packSitesQosFromSdk(ctx, *sdk.Qos)
		diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "Qos"})
		}
		model.Qos = packed
	} else {
		model.Qos = basetypes.NewObjectNull(models.SitesQos{}.AttrTypes())
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.State != nil {
		model.State = basetypes.NewStringValue(*sdk.State)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "State", "value": *sdk.State})
	} else {
		model.State = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Type != nil {
		model.Type = basetypes.NewStringValue(*sdk.Type)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Type", "value": *sdk.Type})
	} else {
		model.Type = basetypes.NewStringNull()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.ZipCode != nil {
		model.ZipCode = basetypes.NewStringValue(*sdk.ZipCode)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "ZipCode", "value": *sdk.ZipCode})
	} else {
		model.ZipCode = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.Sites{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.Sites", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for Sites ---
func unpackSitesListToSdk(ctx context.Context, list types.List) ([]deployment_services.Sites, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.Sites")
	diags := diag.Diagnostics{}
	var data []models.Sites
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]deployment_services.Sites, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.Sites{}.AttrTypes(), &item)
		unpacked, d := unpackSitesToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.Sites", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for Sites ---
func packSitesListFromSdk(ctx context.Context, sdks []deployment_services.Sites) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.Sites")
	diags := diag.Diagnostics{}
	var data []models.Sites

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.Sites
		obj, d := packSitesFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.Sites{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.Sites", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.Sites{}.AttrType(), data)
}

// --- Unpacker for SitesMembersInner ---
func unpackSitesMembersInnerToSdk(ctx context.Context, obj types.Object) (*deployment_services.SitesMembersInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.SitesMembersInner", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.SitesMembersInner
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk deployment_services.SitesMembersInner
	var d diag.Diagnostics
	// Handling Primitives
	if !model.Id.IsNull() && !model.Id.IsUnknown() {
		sdk.Id = model.Id.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Id", "value": *sdk.Id})
	}

	// Handling Primitives
	if !model.Mode.IsNull() && !model.Mode.IsUnknown() {
		sdk.Mode = model.Mode.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Mode", "value": sdk.Mode})
	}

	// Handling Primitives
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		sdk.Name = model.Name.ValueString()
		tflog.Debug(ctx, "Unpacked primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	}

	// Handling Primitives
	if !model.RemoteNetwork.IsNull() && !model.RemoteNetwork.IsUnknown() {
		sdk.RemoteNetwork = model.RemoteNetwork.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "RemoteNetwork", "value": *sdk.RemoteNetwork})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.SitesMembersInner", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for SitesMembersInner ---
func packSitesMembersInnerFromSdk(ctx context.Context, sdk deployment_services.SitesMembersInner) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.SitesMembersInner", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.SitesMembersInner
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
	model.Mode = basetypes.NewStringValue(sdk.Mode)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Mode", "value": sdk.Mode})
	// Handling Primitives
	// Standard primitive packing
	model.Name = basetypes.NewStringValue(sdk.Name)
	tflog.Debug(ctx, "Packed primitive value", map[string]interface{}{"field": "Name", "value": sdk.Name})
	// Handling Primitives
	// Standard primitive packing
	if sdk.RemoteNetwork != nil {
		model.RemoteNetwork = basetypes.NewStringValue(*sdk.RemoteNetwork)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "RemoteNetwork", "value": *sdk.RemoteNetwork})
	} else {
		model.RemoteNetwork = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.SitesMembersInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.SitesMembersInner", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for SitesMembersInner ---
func unpackSitesMembersInnerListToSdk(ctx context.Context, list types.List) ([]deployment_services.SitesMembersInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.SitesMembersInner")
	diags := diag.Diagnostics{}
	var data []models.SitesMembersInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]deployment_services.SitesMembersInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.SitesMembersInner{}.AttrTypes(), &item)
		unpacked, d := unpackSitesMembersInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.SitesMembersInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for SitesMembersInner ---
func packSitesMembersInnerListFromSdk(ctx context.Context, sdks []deployment_services.SitesMembersInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.SitesMembersInner")
	diags := diag.Diagnostics{}
	var data []models.SitesMembersInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.SitesMembersInner
		obj, d := packSitesMembersInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.SitesMembersInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.SitesMembersInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.SitesMembersInner{}.AttrType(), data)
}

// --- Unpacker for SitesQos ---
func unpackSitesQosToSdk(ctx context.Context, obj types.Object) (*deployment_services.SitesQos, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering unpack helper for models.SitesQos", map[string]interface{}{"tf_object": obj})
	diags := diag.Diagnostics{}
	var model models.SitesQos
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
		return nil, diags
	}
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

	var sdk deployment_services.SitesQos
	var d diag.Diagnostics
	// Handling Primitives
	if !model.BackupCir.IsNull() && !model.BackupCir.IsUnknown() {
		val := float32(model.BackupCir.ValueFloat64())
		sdk.BackupCir = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "BackupCir", "value": *sdk.BackupCir})
	}

	// Handling Primitives
	if !model.Cir.IsNull() && !model.Cir.IsUnknown() {
		val := float32(model.Cir.ValueFloat64())
		sdk.Cir = &val
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Cir", "value": *sdk.Cir})
	}

	// Handling Primitives
	if !model.Profile.IsNull() && !model.Profile.IsUnknown() {
		sdk.Profile = model.Profile.ValueStringPointer()
		tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Profile", "value": *sdk.Profile})
	}

	diags.Append(d...)

	tflog.Debug(ctx, "Exiting unpack helper for models.SitesQos", map[string]interface{}{"has_errors": diags.HasError()})
	return &sdk, diags

}

// --- Packer for SitesQos ---
func packSitesQosFromSdk(ctx context.Context, sdk deployment_services.SitesQos) (types.Object, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering pack helper for models.SitesQos", map[string]interface{}{"sdk_struct": sdk})
	diags := diag.Diagnostics{}
	var model models.SitesQos
	var d diag.Diagnostics
	// Handling Primitives
	// Standard primitive packing
	if sdk.BackupCir != nil {
		model.BackupCir = basetypes.NewFloat64Value(float64(*sdk.BackupCir))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "BackupCir", "value": *sdk.BackupCir})
	} else {
		model.BackupCir = basetypes.NewFloat64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Cir != nil {
		model.Cir = basetypes.NewFloat64Value(float64(*sdk.Cir))
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Cir", "value": *sdk.Cir})
	} else {
		model.Cir = basetypes.NewFloat64Null()
	}
	// Handling Primitives
	// Standard primitive packing
	if sdk.Profile != nil {
		model.Profile = basetypes.NewStringValue(*sdk.Profile)
		tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Profile", "value": *sdk.Profile})
	} else {
		model.Profile = basetypes.NewStringNull()
	}
	diags.Append(d...)

	obj, d := types.ObjectValueFrom(ctx, models.SitesQos{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
	diags.Append(d...)
	tflog.Debug(ctx, "Exiting pack helper for models.SitesQos", map[string]interface{}{"has_errors": diags.HasError()})
	return obj, diags

}

// --- List Unpacker for SitesQos ---
func unpackSitesQosListToSdk(ctx context.Context, list types.List) ([]deployment_services.SitesQos, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.SitesQos")
	diags := diag.Diagnostics{}
	var data []models.SitesQos
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]deployment_services.SitesQos, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.SitesQos{}.AttrTypes(), &item)
		unpacked, d := unpackSitesQosToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.SitesQos", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for SitesQos ---
func packSitesQosListFromSdk(ctx context.Context, sdks []deployment_services.SitesQos) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.SitesQos")
	diags := diag.Diagnostics{}
	var data []models.SitesQos

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.SitesQos
		obj, d := packSitesQosFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.SitesQos{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.SitesQos", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.SitesQos{}.AttrType(), data)
}
