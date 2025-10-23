package provider

/*

import (
	"context"
	"fmt"
	"reflect"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/network_services"
	"github.com/paloaltonetworks/scm-go/generated/network_services"
)










// --- Unpacker for AutoVpnClusters ---
func unpackAutoVpnClustersToSdk(ctx context.Context, obj types.Object) (*network_services.AutoVpnClusters, diag.Diagnostics) {
    tflog.Debug(ctx, "Entering unpack helper for models.AutoVpnClusters", map[string]interface{}{"tf_object": obj})
    diags := diag.Diagnostics{}
    var model models.AutoVpnClusters
    diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
    if diags.HasError() {
        tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
        return nil, diags
    }
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

    var sdk network_services.AutoVpnClusters
    var d diag.Diagnostics

    // Handling Lists
    if !model.Branches.IsNull() && !model.Branches.IsUnknown() {
        tflog.Debug(ctx, "Unpacking list of objects for field Branches")
        unpacked, d := unpackAutoVpnClustersBranchesInnerListToSdk(ctx, model.Branches)
        diags.Append(d...)
        sdk.Branches = unpacked
    }

    // Handling Primitives
    if !model.EnableMeshBetweenHubs.IsNull() && !model.EnableMeshBetweenHubs.IsUnknown() {
        sdk.EnableMeshBetweenHubs = model.EnableMeshBetweenHubs.ValueBoolPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "EnableMeshBetweenHubs", "value": *sdk.EnableMeshBetweenHubs})
    }

    // Handling Primitives
    if !model.EnableMeshInterconnect.IsNull() && !model.EnableMeshInterconnect.IsUnknown() {
        sdk.EnableMeshInterconnect = model.EnableMeshInterconnect.ValueBoolPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "EnableMeshInterconnect", "value": *sdk.EnableMeshInterconnect})
    }

    // Handling Primitives
    if !model.EnableSdwan.IsNull() && !model.EnableSdwan.IsUnknown() {
        sdk.EnableSdwan = model.EnableSdwan.ValueBoolPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "EnableSdwan", "value": *sdk.EnableSdwan})
    }

    // Handling Lists
    if !model.Gateways.IsNull() && !model.Gateways.IsUnknown() {
        tflog.Debug(ctx, "Unpacking list of objects for field Gateways")
        unpacked, d := unpackAutoVpnClustersGatewaysInnerListToSdk(ctx, model.Gateways)
        diags.Append(d...)
        sdk.Gateways = unpacked
    }

    // Handling Primitives
    if !model.Id.IsNull() && !model.Id.IsUnknown() {
        sdk.Id = model.Id.ValueStringPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Id", "value": *sdk.Id})
    }

    // Handling Primitives
    if !model.Name.IsNull() && !model.Name.IsUnknown() {
        sdk.Name = model.Name.ValueStringPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
    }

    // Handling Primitives
    if !model.Type.IsNull() && !model.Type.IsUnknown() {
        sdk.Type = model.Type.ValueStringPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Type", "value": *sdk.Type})
    }

diags.Append(d...)

    tflog.Debug(ctx, "Exiting unpack helper for models.AutoVpnClusters", map[string]interface{}{"has_errors": diags.HasError()})
    return &sdk, diags

}

// --- Packer for AutoVpnClusters ---
func packAutoVpnClustersFromSdk(ctx context.Context, sdk network_services.AutoVpnClusters) (types.Object, diag.Diagnostics) {
    tflog.Debug(ctx, "Entering pack helper for models.AutoVpnClusters", map[string]interface{}{"sdk_struct": sdk})
    diags := diag.Diagnostics{}
    var model models.AutoVpnClusters
    var d diag.Diagnostics
    // Handling Lists
    if sdk.Branches != nil {
        tflog.Debug(ctx, "Packing list of objects for field Branches")
        packed, d := packAutoVpnClustersBranchesInnerListFromSdk(ctx, sdk.Branches)
        diags.Append(d...)
        model.Branches = packed
    } else {
        model.Branches = basetypes.NewListNull(models.AutoVpnClustersBranchesInner{}.AttrType())
    }
    // Handling Primitives
    // Standard primitive packing
    if sdk.EnableMeshBetweenHubs != nil {
        model.EnableMeshBetweenHubs = basetypes.NewBoolValue(*sdk.EnableMeshBetweenHubs)
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "EnableMeshBetweenHubs", "value": *sdk.EnableMeshBetweenHubs})
    } else {
        model.EnableMeshBetweenHubs = basetypes.NewBoolNull()
    }
    // Handling Primitives
    // Standard primitive packing
    if sdk.EnableMeshInterconnect != nil {
        model.EnableMeshInterconnect = basetypes.NewBoolValue(*sdk.EnableMeshInterconnect)
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "EnableMeshInterconnect", "value": *sdk.EnableMeshInterconnect})
    } else {
        model.EnableMeshInterconnect = basetypes.NewBoolNull()
    }
    // Handling Primitives
    // Standard primitive packing
    if sdk.EnableSdwan != nil {
        model.EnableSdwan = basetypes.NewBoolValue(*sdk.EnableSdwan)
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "EnableSdwan", "value": *sdk.EnableSdwan})
    } else {
        model.EnableSdwan = basetypes.NewBoolNull()
    }
    // Handling Lists
    if sdk.Gateways != nil {
        tflog.Debug(ctx, "Packing list of objects for field Gateways")
        packed, d := packAutoVpnClustersGatewaysInnerListFromSdk(ctx, sdk.Gateways)
        diags.Append(d...)
        model.Gateways = packed
    } else {
        model.Gateways = basetypes.NewListNull(models.AutoVpnClustersGatewaysInner{}.AttrType())
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
    if sdk.Name != nil {
        model.Name = basetypes.NewStringValue(*sdk.Name)
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
    } else {
        model.Name = basetypes.NewStringNull()
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

    obj, d := types.ObjectValueFrom(ctx, models.AutoVpnClusters{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
    diags.Append(d...)
    tflog.Debug(ctx, "Exiting pack helper for models.AutoVpnClusters", map[string]interface{}{"has_errors": diags.HasError()})
    return obj, diags

}

// --- List Unpacker for AutoVpnClusters ---
func unpackAutoVpnClustersListToSdk(ctx context.Context, list types.List) ([]network_services.AutoVpnClusters, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.AutoVpnClusters")
	diags := diag.Diagnostics{}
	var data []models.AutoVpnClusters
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.AutoVpnClusters, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.AutoVpnClusters{}.AttrTypes(), &item)
		unpacked, d := unpackAutoVpnClustersToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.AutoVpnClusters", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for AutoVpnClusters ---
func packAutoVpnClustersListFromSdk(ctx context.Context, sdks []network_services.AutoVpnClusters) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.AutoVpnClusters")
	diags := diag.Diagnostics{}
	var data []models.AutoVpnClusters

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.AutoVpnClusters
		obj, d := packAutoVpnClustersFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.AutoVpnClusters{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.AutoVpnClusters", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.AutoVpnClusters{}.AttrType(), data)
}

// --- Unpacker for AutoVpnClustersBranchesInner ---
func unpackAutoVpnClustersBranchesInnerToSdk(ctx context.Context, obj types.Object) (*network_services.AutoVpnClustersBranchesInner, diag.Diagnostics) {
    tflog.Debug(ctx, "Entering unpack helper for models.AutoVpnClustersBranchesInner", map[string]interface{}{"tf_object": obj})
    diags := diag.Diagnostics{}
    var model models.AutoVpnClustersBranchesInner
    diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
    if diags.HasError() {
        tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
        return nil, diags
    }
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

    var sdk network_services.AutoVpnClustersBranchesInner
    var d diag.Diagnostics
    // Handling Primitives
    if !model.BgpRedistributionProfile.IsNull() && !model.BgpRedistributionProfile.IsUnknown() {
        sdk.BgpRedistributionProfile = model.BgpRedistributionProfile.ValueStringPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "BgpRedistributionProfile", "value": *sdk.BgpRedistributionProfile})
    }

    // Handling Lists
    if !model.Interfaces.IsNull() && !model.Interfaces.IsUnknown() {
        tflog.Debug(ctx, "Unpacking list of objects for field Interfaces")
        unpacked, d := unpackAutoVpnClustersBranchesInnerInterfacesInnerListToSdk(ctx, model.Interfaces)
        diags.Append(d...)
        sdk.Interfaces = unpacked
    }

    // Handling Primitives
    if !model.LogicalRouter.IsNull() && !model.LogicalRouter.IsUnknown() {
        sdk.LogicalRouter = model.LogicalRouter.ValueStringPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "LogicalRouter", "value": *sdk.LogicalRouter})
    }

    // Handling Primitives
    if !model.Name.IsNull() && !model.Name.IsUnknown() {
        sdk.Name = model.Name.ValueStringPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
    }

    // Handling Lists
    if !model.PrivateInterfaces.IsNull() && !model.PrivateInterfaces.IsUnknown() {
        tflog.Debug(ctx, "Unpacking list of objects for field PrivateInterfaces")
        unpacked, d := unpackAutoVpnClustersBranchesInnerPrivateInterfacesInnerListToSdk(ctx, model.PrivateInterfaces)
        diags.Append(d...)
        sdk.PrivateInterfaces = unpacked
    }

    // Handling Primitives
    if !model.Site.IsNull() && !model.Site.IsUnknown() {
        sdk.Site = model.Site.ValueStringPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Site", "value": *sdk.Site})
    }

diags.Append(d...)

    tflog.Debug(ctx, "Exiting unpack helper for models.AutoVpnClustersBranchesInner", map[string]interface{}{"has_errors": diags.HasError()})
    return &sdk, diags

}

// --- Packer for AutoVpnClustersBranchesInner ---
func packAutoVpnClustersBranchesInnerFromSdk(ctx context.Context, sdk network_services.AutoVpnClustersBranchesInner) (types.Object, diag.Diagnostics) {
    tflog.Debug(ctx, "Entering pack helper for models.AutoVpnClustersBranchesInner", map[string]interface{}{"sdk_struct": sdk})
    diags := diag.Diagnostics{}
    var model models.AutoVpnClustersBranchesInner
    var d diag.Diagnostics
    // Handling Primitives
    // Standard primitive packing
    if sdk.BgpRedistributionProfile != nil {
        model.BgpRedistributionProfile = basetypes.NewStringValue(*sdk.BgpRedistributionProfile)
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "BgpRedistributionProfile", "value": *sdk.BgpRedistributionProfile})
    } else {
        model.BgpRedistributionProfile = basetypes.NewStringNull()
    }
    // Handling Lists
    if sdk.Interfaces != nil {
        tflog.Debug(ctx, "Packing list of objects for field Interfaces")
        packed, d := packAutoVpnClustersBranchesInnerInterfacesInnerListFromSdk(ctx, sdk.Interfaces)
        diags.Append(d...)
        model.Interfaces = packed
    } else {
        model.Interfaces = basetypes.NewListNull(models.AutoVpnClustersBranchesInnerInterfacesInner{}.AttrType())
    }
    // Handling Primitives
    // Standard primitive packing
    if sdk.LogicalRouter != nil {
        model.LogicalRouter = basetypes.NewStringValue(*sdk.LogicalRouter)
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "LogicalRouter", "value": *sdk.LogicalRouter})
    } else {
        model.LogicalRouter = basetypes.NewStringNull()
    }
    // Handling Primitives
    // Standard primitive packing
    if sdk.Name != nil {
        model.Name = basetypes.NewStringValue(*sdk.Name)
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
    } else {
        model.Name = basetypes.NewStringNull()
    }
    // Handling Lists
    if sdk.PrivateInterfaces != nil {
        tflog.Debug(ctx, "Packing list of objects for field PrivateInterfaces")
        packed, d := packAutoVpnClustersBranchesInnerPrivateInterfacesInnerListFromSdk(ctx, sdk.PrivateInterfaces)
        diags.Append(d...)
        model.PrivateInterfaces = packed
    } else {
        model.PrivateInterfaces = basetypes.NewListNull(models.AutoVpnClustersBranchesInnerPrivateInterfacesInner{}.AttrType())
    }
    // Handling Primitives
    // Standard primitive packing
    if sdk.Site != nil {
        model.Site = basetypes.NewStringValue(*sdk.Site)
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Site", "value": *sdk.Site})
    } else {
        model.Site = basetypes.NewStringNull()
    }
diags.Append(d...)

    obj, d := types.ObjectValueFrom(ctx, models.AutoVpnClustersBranchesInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
    diags.Append(d...)
    tflog.Debug(ctx, "Exiting pack helper for models.AutoVpnClustersBranchesInner", map[string]interface{}{"has_errors": diags.HasError()})
    return obj, diags

}

// --- List Unpacker for AutoVpnClustersBranchesInner ---
func unpackAutoVpnClustersBranchesInnerListToSdk(ctx context.Context, list types.List) ([]network_services.AutoVpnClustersBranchesInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.AutoVpnClustersBranchesInner")
	diags := diag.Diagnostics{}
	var data []models.AutoVpnClustersBranchesInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.AutoVpnClustersBranchesInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.AutoVpnClustersBranchesInner{}.AttrTypes(), &item)
		unpacked, d := unpackAutoVpnClustersBranchesInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.AutoVpnClustersBranchesInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for AutoVpnClustersBranchesInner ---
func packAutoVpnClustersBranchesInnerListFromSdk(ctx context.Context, sdks []network_services.AutoVpnClustersBranchesInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.AutoVpnClustersBranchesInner")
	diags := diag.Diagnostics{}
	var data []models.AutoVpnClustersBranchesInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.AutoVpnClustersBranchesInner
		obj, d := packAutoVpnClustersBranchesInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.AutoVpnClustersBranchesInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.AutoVpnClustersBranchesInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.AutoVpnClustersBranchesInner{}.AttrType(), data)
}

// --- Unpacker for AutoVpnClustersBranchesInnerInterfacesInner ---
func unpackAutoVpnClustersBranchesInnerInterfacesInnerToSdk(ctx context.Context, obj types.Object) (*network_services.AutoVpnClustersBranchesInnerInterfacesInner, diag.Diagnostics) {
    tflog.Debug(ctx, "Entering unpack helper for models.AutoVpnClustersBranchesInnerInterfacesInner", map[string]interface{}{"tf_object": obj})
    diags := diag.Diagnostics{}
    var model models.AutoVpnClustersBranchesInnerInterfacesInner
    diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
    if diags.HasError() {
        tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
        return nil, diags
    }
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

    var sdk network_services.AutoVpnClustersBranchesInnerInterfacesInner
    var d diag.Diagnostics
    // Handling Primitives
    if !model.DhcpIp.IsNull() && !model.DhcpIp.IsUnknown() {
        sdk.DhcpIp = model.DhcpIp.ValueStringPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DhcpIp", "value": *sdk.DhcpIp})
    }

    // Handling Primitives
    if !model.Name.IsNull() && !model.Name.IsUnknown() {
        sdk.Name = model.Name.ValueStringPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
    }

    // Handling Objects
    if !model.SdwanLinkSettings.IsNull() && !model.SdwanLinkSettings.IsUnknown() {
        tflog.Debug(ctx, "Unpacking nested object for field SdwanLinkSettings")
        unpacked, d := unpackAutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsToSdk(ctx, model.SdwanLinkSettings)
        diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "SdwanLinkSettings"})
		}
        if unpacked != nil {
            sdk.SdwanLinkSettings = unpacked
        }
    }

diags.Append(d...)

    tflog.Debug(ctx, "Exiting unpack helper for models.AutoVpnClustersBranchesInnerInterfacesInner", map[string]interface{}{"has_errors": diags.HasError()})
    return &sdk, diags

}

// --- Packer for AutoVpnClustersBranchesInnerInterfacesInner ---
func packAutoVpnClustersBranchesInnerInterfacesInnerFromSdk(ctx context.Context, sdk network_services.AutoVpnClustersBranchesInnerInterfacesInner) (types.Object, diag.Diagnostics) {
    tflog.Debug(ctx, "Entering pack helper for models.AutoVpnClustersBranchesInnerInterfacesInner", map[string]interface{}{"sdk_struct": sdk})
    diags := diag.Diagnostics{}
    var model models.AutoVpnClustersBranchesInnerInterfacesInner
    var d diag.Diagnostics
    // Handling Primitives
    // Standard primitive packing
    if sdk.DhcpIp != nil {
        model.DhcpIp = basetypes.NewStringValue(*sdk.DhcpIp)
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DhcpIp", "value": *sdk.DhcpIp})
    } else {
        model.DhcpIp = basetypes.NewStringNull()
    }
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
    if sdk.SdwanLinkSettings != nil {
        tflog.Debug(ctx, "Packing nested object for field SdwanLinkSettings")
        packed, d := packAutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsFromSdk(ctx, *sdk.SdwanLinkSettings)
        diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "SdwanLinkSettings"})
		}
        model.SdwanLinkSettings = packed
    } else {
        model.SdwanLinkSettings = basetypes.NewObjectNull(models.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettings{}.AttrTypes())
    }
diags.Append(d...)

    obj, d := types.ObjectValueFrom(ctx, models.AutoVpnClustersBranchesInnerInterfacesInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
    diags.Append(d...)
    tflog.Debug(ctx, "Exiting pack helper for models.AutoVpnClustersBranchesInnerInterfacesInner", map[string]interface{}{"has_errors": diags.HasError()})
    return obj, diags

}

// --- List Unpacker for AutoVpnClustersBranchesInnerInterfacesInner ---
func unpackAutoVpnClustersBranchesInnerInterfacesInnerListToSdk(ctx context.Context, list types.List) ([]network_services.AutoVpnClustersBranchesInnerInterfacesInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.AutoVpnClustersBranchesInnerInterfacesInner")
	diags := diag.Diagnostics{}
	var data []models.AutoVpnClustersBranchesInnerInterfacesInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.AutoVpnClustersBranchesInnerInterfacesInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.AutoVpnClustersBranchesInnerInterfacesInner{}.AttrTypes(), &item)
		unpacked, d := unpackAutoVpnClustersBranchesInnerInterfacesInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.AutoVpnClustersBranchesInnerInterfacesInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for AutoVpnClustersBranchesInnerInterfacesInner ---
func packAutoVpnClustersBranchesInnerInterfacesInnerListFromSdk(ctx context.Context, sdks []network_services.AutoVpnClustersBranchesInnerInterfacesInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.AutoVpnClustersBranchesInnerInterfacesInner")
	diags := diag.Diagnostics{}
	var data []models.AutoVpnClustersBranchesInnerInterfacesInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.AutoVpnClustersBranchesInnerInterfacesInner
		obj, d := packAutoVpnClustersBranchesInnerInterfacesInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.AutoVpnClustersBranchesInnerInterfacesInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.AutoVpnClustersBranchesInnerInterfacesInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.AutoVpnClustersBranchesInnerInterfacesInner{}.AttrType(), data)
}

// --- Unpacker for AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettings ---
func unpackAutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsToSdk(ctx context.Context, obj types.Object) (*network_services.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettings, diag.Diagnostics) {
    tflog.Debug(ctx, "Entering unpack helper for models.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettings", map[string]interface{}{"tf_object": obj})
    diags := diag.Diagnostics{}
    var model models.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettings
    diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
    if diags.HasError() {
        tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
        return nil, diags
    }
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

    var sdk network_services.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettings
    var d diag.Diagnostics
    // Handling Primitives
    if !model.SdwanGateway.IsNull() && !model.SdwanGateway.IsUnknown() {
        sdk.SdwanGateway = model.SdwanGateway.ValueStringPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "SdwanGateway", "value": *sdk.SdwanGateway})
    }

    // Handling Primitives
    if !model.SdwanInterfaceProfile.IsNull() && !model.SdwanInterfaceProfile.IsUnknown() {
        sdk.SdwanInterfaceProfile = model.SdwanInterfaceProfile.ValueStringPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "SdwanInterfaceProfile", "value": *sdk.SdwanInterfaceProfile})
    }

    // Handling Objects
    if !model.UpstreamNat.IsNull() && !model.UpstreamNat.IsUnknown() {
        tflog.Debug(ctx, "Unpacking nested object for field UpstreamNat")
        unpacked, d := unpackAutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNatToSdk(ctx, model.UpstreamNat)
        diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "UpstreamNat"})
		}
        if unpacked != nil {
            sdk.UpstreamNat = unpacked
        }
    }

diags.Append(d...)

    tflog.Debug(ctx, "Exiting unpack helper for models.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettings", map[string]interface{}{"has_errors": diags.HasError()})
    return &sdk, diags

}

// --- Packer for AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettings ---
func packAutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsFromSdk(ctx context.Context, sdk network_services.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettings) (types.Object, diag.Diagnostics) {
    tflog.Debug(ctx, "Entering pack helper for models.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettings", map[string]interface{}{"sdk_struct": sdk})
    diags := diag.Diagnostics{}
    var model models.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettings
    var d diag.Diagnostics
    // Handling Primitives
    // Standard primitive packing
    if sdk.SdwanGateway != nil {
        model.SdwanGateway = basetypes.NewStringValue(*sdk.SdwanGateway)
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "SdwanGateway", "value": *sdk.SdwanGateway})
    } else {
        model.SdwanGateway = basetypes.NewStringNull()
    }
    // Handling Primitives
    // Standard primitive packing
    if sdk.SdwanInterfaceProfile != nil {
        model.SdwanInterfaceProfile = basetypes.NewStringValue(*sdk.SdwanInterfaceProfile)
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "SdwanInterfaceProfile", "value": *sdk.SdwanInterfaceProfile})
    } else {
        model.SdwanInterfaceProfile = basetypes.NewStringNull()
    }
    // Handling Objects
    // This is a regular nested object that has its own packer.
    if sdk.UpstreamNat != nil {
        tflog.Debug(ctx, "Packing nested object for field UpstreamNat")
        packed, d := packAutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNatFromSdk(ctx, *sdk.UpstreamNat)
        diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "UpstreamNat"})
		}
        model.UpstreamNat = packed
    } else {
        model.UpstreamNat = basetypes.NewObjectNull(models.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNat{}.AttrTypes())
    }
diags.Append(d...)

    obj, d := types.ObjectValueFrom(ctx, models.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettings{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
    diags.Append(d...)
    tflog.Debug(ctx, "Exiting pack helper for models.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettings", map[string]interface{}{"has_errors": diags.HasError()})
    return obj, diags

}

// --- List Unpacker for AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettings ---
func unpackAutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsListToSdk(ctx context.Context, list types.List) ([]network_services.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettings, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettings")
	diags := diag.Diagnostics{}
	var data []models.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettings
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettings, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettings{}.AttrTypes(), &item)
		unpacked, d := unpackAutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettings", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettings ---
func packAutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsListFromSdk(ctx context.Context, sdks []network_services.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettings) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettings")
	diags := diag.Diagnostics{}
	var data []models.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettings

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettings
		obj, d := packAutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettings{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettings", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettings{}.AttrType(), data)
}

// --- Unpacker for AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNat ---
func unpackAutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNatToSdk(ctx context.Context, obj types.Object) (*network_services.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNat, diag.Diagnostics) {
    tflog.Debug(ctx, "Entering unpack helper for models.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNat", map[string]interface{}{"tf_object": obj})
    diags := diag.Diagnostics{}
    var model models.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNat
    diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
    if diags.HasError() {
        tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
        return nil, diags
    }
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

    var sdk network_services.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNat
    var d diag.Diagnostics
    // Handling Primitives
    if !model.Enable.IsNull() && !model.Enable.IsUnknown() {
        sdk.Enable = model.Enable.ValueBoolPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Enable", "value": *sdk.Enable})
    }

    // Handling Objects
    if !model.StaticIp.IsNull() && !model.StaticIp.IsUnknown() {
        tflog.Debug(ctx, "Unpacking nested object for field StaticIp")
        unpacked, d := unpackAutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNatStaticIpToSdk(ctx, model.StaticIp)
        diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "StaticIp"})
		}
        if unpacked != nil {
            sdk.StaticIp = unpacked
        }
    }

diags.Append(d...)

    tflog.Debug(ctx, "Exiting unpack helper for models.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNat", map[string]interface{}{"has_errors": diags.HasError()})
    return &sdk, diags

}

// --- Packer for AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNat ---
func packAutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNatFromSdk(ctx context.Context, sdk network_services.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNat) (types.Object, diag.Diagnostics) {
    tflog.Debug(ctx, "Entering pack helper for models.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNat", map[string]interface{}{"sdk_struct": sdk})
    diags := diag.Diagnostics{}
    var model models.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNat
    var d diag.Diagnostics
    // Handling Primitives
    // Standard primitive packing
    if sdk.Enable != nil {
        model.Enable = basetypes.NewBoolValue(*sdk.Enable)
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Enable", "value": *sdk.Enable})
    } else {
        model.Enable = basetypes.NewBoolNull()
    }
    // Handling Objects
    // This is a regular nested object that has its own packer.
    if sdk.StaticIp != nil {
        tflog.Debug(ctx, "Packing nested object for field StaticIp")
        packed, d := packAutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNatStaticIpFromSdk(ctx, *sdk.StaticIp)
        diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "StaticIp"})
		}
        model.StaticIp = packed
    } else {
        model.StaticIp = basetypes.NewObjectNull(models.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNatStaticIp{}.AttrTypes())
    }
diags.Append(d...)

    obj, d := types.ObjectValueFrom(ctx, models.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNat{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
    diags.Append(d...)
    tflog.Debug(ctx, "Exiting pack helper for models.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNat", map[string]interface{}{"has_errors": diags.HasError()})
    return obj, diags

}

// --- List Unpacker for AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNat ---
func unpackAutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNatListToSdk(ctx context.Context, list types.List) ([]network_services.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNat, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNat")
	diags := diag.Diagnostics{}
	var data []models.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNat
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNat, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNat{}.AttrTypes(), &item)
		unpacked, d := unpackAutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNatToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNat", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNat ---
func packAutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNatListFromSdk(ctx context.Context, sdks []network_services.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNat) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNat")
	diags := diag.Diagnostics{}
	var data []models.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNat

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNat
		obj, d := packAutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNatFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNat{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNat", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNat{}.AttrType(), data)
}

// --- Unpacker for AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNatStaticIp ---
func unpackAutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNatStaticIpToSdk(ctx context.Context, obj types.Object) (*network_services.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNatStaticIp, diag.Diagnostics) {
    tflog.Debug(ctx, "Entering unpack helper for models.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNatStaticIp", map[string]interface{}{"tf_object": obj})
    diags := diag.Diagnostics{}
    var model models.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNatStaticIp
    diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
    if diags.HasError() {
        tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
        return nil, diags
    }
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

    var sdk network_services.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNatStaticIp
    var d diag.Diagnostics
    // Handling Primitives
    if !model.Fqdn.IsNull() && !model.Fqdn.IsUnknown() {
        sdk.Fqdn = model.Fqdn.ValueStringPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Fqdn", "value": *sdk.Fqdn})
    }

    // Handling Primitives
    if !model.IpAddress.IsNull() && !model.IpAddress.IsUnknown() {
        sdk.IpAddress = model.IpAddress.ValueStringPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "IpAddress", "value": *sdk.IpAddress})
    }

diags.Append(d...)

    tflog.Debug(ctx, "Exiting unpack helper for models.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNatStaticIp", map[string]interface{}{"has_errors": diags.HasError()})
    return &sdk, diags

}

// --- Packer for AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNatStaticIp ---
func packAutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNatStaticIpFromSdk(ctx context.Context, sdk network_services.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNatStaticIp) (types.Object, diag.Diagnostics) {
    tflog.Debug(ctx, "Entering pack helper for models.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNatStaticIp", map[string]interface{}{"sdk_struct": sdk})
    diags := diag.Diagnostics{}
    var model models.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNatStaticIp
    var d diag.Diagnostics
    // Handling Primitives
    // Standard primitive packing
    if sdk.Fqdn != nil {
        model.Fqdn = basetypes.NewStringValue(*sdk.Fqdn)
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Fqdn", "value": *sdk.Fqdn})
    } else {
        model.Fqdn = basetypes.NewStringNull()
    }
    // Handling Primitives
    // Standard primitive packing
    if sdk.IpAddress != nil {
        model.IpAddress = basetypes.NewStringValue(*sdk.IpAddress)
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "IpAddress", "value": *sdk.IpAddress})
    } else {
        model.IpAddress = basetypes.NewStringNull()
    }
diags.Append(d...)

    obj, d := types.ObjectValueFrom(ctx, models.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNatStaticIp{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
    diags.Append(d...)
    tflog.Debug(ctx, "Exiting pack helper for models.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNatStaticIp", map[string]interface{}{"has_errors": diags.HasError()})
    return obj, diags

}

// --- List Unpacker for AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNatStaticIp ---
func unpackAutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNatStaticIpListToSdk(ctx context.Context, list types.List) ([]network_services.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNatStaticIp, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNatStaticIp")
	diags := diag.Diagnostics{}
	var data []models.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNatStaticIp
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNatStaticIp, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNatStaticIp{}.AttrTypes(), &item)
		unpacked, d := unpackAutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNatStaticIpToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNatStaticIp", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNatStaticIp ---
func packAutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNatStaticIpListFromSdk(ctx context.Context, sdks []network_services.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNatStaticIp) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNatStaticIp")
	diags := diag.Diagnostics{}
	var data []models.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNatStaticIp

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNatStaticIp
		obj, d := packAutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNatStaticIpFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNatStaticIp{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNatStaticIp", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNatStaticIp{}.AttrType(), data)
}

// --- Unpacker for AutoVpnClustersBranchesInnerPrivateInterfacesInner ---
func unpackAutoVpnClustersBranchesInnerPrivateInterfacesInnerToSdk(ctx context.Context, obj types.Object) (*network_services.AutoVpnClustersBranchesInnerPrivateInterfacesInner, diag.Diagnostics) {
    tflog.Debug(ctx, "Entering unpack helper for models.AutoVpnClustersBranchesInnerPrivateInterfacesInner", map[string]interface{}{"tf_object": obj})
    diags := diag.Diagnostics{}
    var model models.AutoVpnClustersBranchesInnerPrivateInterfacesInner
    diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
    if diags.HasError() {
        tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
        return nil, diags
    }
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

    var sdk network_services.AutoVpnClustersBranchesInnerPrivateInterfacesInner
    var d diag.Diagnostics
    // Handling Primitives
    if !model.Name.IsNull() && !model.Name.IsUnknown() {
        sdk.Name = model.Name.ValueStringPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
    }

    // Handling Objects
    if !model.SdwanLinkSettings.IsNull() && !model.SdwanLinkSettings.IsUnknown() {
        tflog.Debug(ctx, "Unpacking nested object for field SdwanLinkSettings")
        unpacked, d := unpackAutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsToSdk(ctx, model.SdwanLinkSettings)
        diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "SdwanLinkSettings"})
		}
        if unpacked != nil {
            sdk.SdwanLinkSettings = unpacked
        }
    }

diags.Append(d...)

    tflog.Debug(ctx, "Exiting unpack helper for models.AutoVpnClustersBranchesInnerPrivateInterfacesInner", map[string]interface{}{"has_errors": diags.HasError()})
    return &sdk, diags

}

// --- Packer for AutoVpnClustersBranchesInnerPrivateInterfacesInner ---
func packAutoVpnClustersBranchesInnerPrivateInterfacesInnerFromSdk(ctx context.Context, sdk network_services.AutoVpnClustersBranchesInnerPrivateInterfacesInner) (types.Object, diag.Diagnostics) {
    tflog.Debug(ctx, "Entering pack helper for models.AutoVpnClustersBranchesInnerPrivateInterfacesInner", map[string]interface{}{"sdk_struct": sdk})
    diags := diag.Diagnostics{}
    var model models.AutoVpnClustersBranchesInnerPrivateInterfacesInner
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
    if sdk.SdwanLinkSettings != nil {
        tflog.Debug(ctx, "Packing nested object for field SdwanLinkSettings")
        packed, d := packAutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsFromSdk(ctx, *sdk.SdwanLinkSettings)
        diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "SdwanLinkSettings"})
		}
        model.SdwanLinkSettings = packed
    } else {
        model.SdwanLinkSettings = basetypes.NewObjectNull(models.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettings{}.AttrTypes())
    }
diags.Append(d...)

    obj, d := types.ObjectValueFrom(ctx, models.AutoVpnClustersBranchesInnerPrivateInterfacesInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
    diags.Append(d...)
    tflog.Debug(ctx, "Exiting pack helper for models.AutoVpnClustersBranchesInnerPrivateInterfacesInner", map[string]interface{}{"has_errors": diags.HasError()})
    return obj, diags

}

// --- List Unpacker for AutoVpnClustersBranchesInnerPrivateInterfacesInner ---
func unpackAutoVpnClustersBranchesInnerPrivateInterfacesInnerListToSdk(ctx context.Context, list types.List) ([]network_services.AutoVpnClustersBranchesInnerPrivateInterfacesInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.AutoVpnClustersBranchesInnerPrivateInterfacesInner")
	diags := diag.Diagnostics{}
	var data []models.AutoVpnClustersBranchesInnerPrivateInterfacesInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.AutoVpnClustersBranchesInnerPrivateInterfacesInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.AutoVpnClustersBranchesInnerPrivateInterfacesInner{}.AttrTypes(), &item)
		unpacked, d := unpackAutoVpnClustersBranchesInnerPrivateInterfacesInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.AutoVpnClustersBranchesInnerPrivateInterfacesInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for AutoVpnClustersBranchesInnerPrivateInterfacesInner ---
func packAutoVpnClustersBranchesInnerPrivateInterfacesInnerListFromSdk(ctx context.Context, sdks []network_services.AutoVpnClustersBranchesInnerPrivateInterfacesInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.AutoVpnClustersBranchesInnerPrivateInterfacesInner")
	diags := diag.Diagnostics{}
	var data []models.AutoVpnClustersBranchesInnerPrivateInterfacesInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.AutoVpnClustersBranchesInnerPrivateInterfacesInner
		obj, d := packAutoVpnClustersBranchesInnerPrivateInterfacesInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.AutoVpnClustersBranchesInnerPrivateInterfacesInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.AutoVpnClustersBranchesInnerPrivateInterfacesInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.AutoVpnClustersBranchesInnerPrivateInterfacesInner{}.AttrType(), data)
}

// --- Unpacker for AutoVpnClustersGatewaysInner ---
func unpackAutoVpnClustersGatewaysInnerToSdk(ctx context.Context, obj types.Object) (*network_services.AutoVpnClustersGatewaysInner, diag.Diagnostics) {
    tflog.Debug(ctx, "Entering unpack helper for models.AutoVpnClustersGatewaysInner", map[string]interface{}{"tf_object": obj})
    diags := diag.Diagnostics{}
    var model models.AutoVpnClustersGatewaysInner
    diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
    if diags.HasError() {
        tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
        return nil, diags
    }
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

    var sdk network_services.AutoVpnClustersGatewaysInner
    var d diag.Diagnostics
    // Handling Primitives
    if !model.AllowDiaVpnFailover.IsNull() && !model.AllowDiaVpnFailover.IsUnknown() {
        sdk.AllowDiaVpnFailover = model.AllowDiaVpnFailover.ValueBoolPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "AllowDiaVpnFailover", "value": *sdk.AllowDiaVpnFailover})
    }

    // Handling Primitives
    if !model.BgpRedistributionProfile.IsNull() && !model.BgpRedistributionProfile.IsUnknown() {
        sdk.BgpRedistributionProfile = model.BgpRedistributionProfile.ValueStringPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "BgpRedistributionProfile", "value": *sdk.BgpRedistributionProfile})
    }

    // Handling Lists
    if !model.Interfaces.IsNull() && !model.Interfaces.IsUnknown() {
        tflog.Debug(ctx, "Unpacking list of objects for field Interfaces")
        unpacked, d := unpackAutoVpnClustersGatewaysInnerInterfacesInnerListToSdk(ctx, model.Interfaces)
        diags.Append(d...)
        sdk.Interfaces = unpacked
    }

    // Handling Primitives
    if !model.LogicalRouter.IsNull() && !model.LogicalRouter.IsUnknown() {
        sdk.LogicalRouter = model.LogicalRouter.ValueStringPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "LogicalRouter", "value": *sdk.LogicalRouter})
    }

    // Handling Primitives
    if !model.Name.IsNull() && !model.Name.IsUnknown() {
        sdk.Name = model.Name.ValueStringPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
    }

    // Handling Primitives
    if !model.Priority.IsNull() && !model.Priority.IsUnknown() {
        sdk.Priority = model.Priority.ValueStringPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Priority", "value": *sdk.Priority})
    }

    // Handling Lists
    if !model.PrivateInterfaces.IsNull() && !model.PrivateInterfaces.IsUnknown() {
        tflog.Debug(ctx, "Unpacking list of objects for field PrivateInterfaces")
        unpacked, d := unpackAutoVpnClustersGatewaysInnerPrivateInterfacesInnerListToSdk(ctx, model.PrivateInterfaces)
        diags.Append(d...)
        sdk.PrivateInterfaces = unpacked
    }

    // Handling Primitives
    if !model.Site.IsNull() && !model.Site.IsUnknown() {
        sdk.Site = model.Site.ValueStringPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Site", "value": *sdk.Site})
    }

diags.Append(d...)

    tflog.Debug(ctx, "Exiting unpack helper for models.AutoVpnClustersGatewaysInner", map[string]interface{}{"has_errors": diags.HasError()})
    return &sdk, diags

}

// --- Packer for AutoVpnClustersGatewaysInner ---
func packAutoVpnClustersGatewaysInnerFromSdk(ctx context.Context, sdk network_services.AutoVpnClustersGatewaysInner) (types.Object, diag.Diagnostics) {
    tflog.Debug(ctx, "Entering pack helper for models.AutoVpnClustersGatewaysInner", map[string]interface{}{"sdk_struct": sdk})
    diags := diag.Diagnostics{}
    var model models.AutoVpnClustersGatewaysInner
    var d diag.Diagnostics
    // Handling Primitives
    // Standard primitive packing
    if sdk.AllowDiaVpnFailover != nil {
        model.AllowDiaVpnFailover = basetypes.NewBoolValue(*sdk.AllowDiaVpnFailover)
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "AllowDiaVpnFailover", "value": *sdk.AllowDiaVpnFailover})
    } else {
        model.AllowDiaVpnFailover = basetypes.NewBoolNull()
    }
    // Handling Primitives
    // Standard primitive packing
    if sdk.BgpRedistributionProfile != nil {
        model.BgpRedistributionProfile = basetypes.NewStringValue(*sdk.BgpRedistributionProfile)
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "BgpRedistributionProfile", "value": *sdk.BgpRedistributionProfile})
    } else {
        model.BgpRedistributionProfile = basetypes.NewStringNull()
    }
    // Handling Lists
    if sdk.Interfaces != nil {
        tflog.Debug(ctx, "Packing list of objects for field Interfaces")
        packed, d := packAutoVpnClustersGatewaysInnerInterfacesInnerListFromSdk(ctx, sdk.Interfaces)
        diags.Append(d...)
        model.Interfaces = packed
    } else {
        model.Interfaces = basetypes.NewListNull(models.AutoVpnClustersGatewaysInnerInterfacesInner{}.AttrType())
    }
    // Handling Primitives
    // Standard primitive packing
    if sdk.LogicalRouter != nil {
        model.LogicalRouter = basetypes.NewStringValue(*sdk.LogicalRouter)
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "LogicalRouter", "value": *sdk.LogicalRouter})
    } else {
        model.LogicalRouter = basetypes.NewStringNull()
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
    if sdk.Priority != nil {
        model.Priority = basetypes.NewStringValue(*sdk.Priority)
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Priority", "value": *sdk.Priority})
    } else {
        model.Priority = basetypes.NewStringNull()
    }
    // Handling Lists
    if sdk.PrivateInterfaces != nil {
        tflog.Debug(ctx, "Packing list of objects for field PrivateInterfaces")
        packed, d := packAutoVpnClustersGatewaysInnerPrivateInterfacesInnerListFromSdk(ctx, sdk.PrivateInterfaces)
        diags.Append(d...)
        model.PrivateInterfaces = packed
    } else {
        model.PrivateInterfaces = basetypes.NewListNull(models.AutoVpnClustersGatewaysInnerPrivateInterfacesInner{}.AttrType())
    }
    // Handling Primitives
    // Standard primitive packing
    if sdk.Site != nil {
        model.Site = basetypes.NewStringValue(*sdk.Site)
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Site", "value": *sdk.Site})
    } else {
        model.Site = basetypes.NewStringNull()
    }
diags.Append(d...)

    obj, d := types.ObjectValueFrom(ctx, models.AutoVpnClustersGatewaysInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
    diags.Append(d...)
    tflog.Debug(ctx, "Exiting pack helper for models.AutoVpnClustersGatewaysInner", map[string]interface{}{"has_errors": diags.HasError()})
    return obj, diags

}

// --- List Unpacker for AutoVpnClustersGatewaysInner ---
func unpackAutoVpnClustersGatewaysInnerListToSdk(ctx context.Context, list types.List) ([]network_services.AutoVpnClustersGatewaysInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.AutoVpnClustersGatewaysInner")
	diags := diag.Diagnostics{}
	var data []models.AutoVpnClustersGatewaysInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.AutoVpnClustersGatewaysInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.AutoVpnClustersGatewaysInner{}.AttrTypes(), &item)
		unpacked, d := unpackAutoVpnClustersGatewaysInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.AutoVpnClustersGatewaysInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for AutoVpnClustersGatewaysInner ---
func packAutoVpnClustersGatewaysInnerListFromSdk(ctx context.Context, sdks []network_services.AutoVpnClustersGatewaysInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.AutoVpnClustersGatewaysInner")
	diags := diag.Diagnostics{}
	var data []models.AutoVpnClustersGatewaysInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.AutoVpnClustersGatewaysInner
		obj, d := packAutoVpnClustersGatewaysInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.AutoVpnClustersGatewaysInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.AutoVpnClustersGatewaysInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.AutoVpnClustersGatewaysInner{}.AttrType(), data)
}

// --- Unpacker for AutoVpnClustersGatewaysInnerInterfacesInner ---
func unpackAutoVpnClustersGatewaysInnerInterfacesInnerToSdk(ctx context.Context, obj types.Object) (*network_services.AutoVpnClustersGatewaysInnerInterfacesInner, diag.Diagnostics) {
    tflog.Debug(ctx, "Entering unpack helper for models.AutoVpnClustersGatewaysInnerInterfacesInner", map[string]interface{}{"tf_object": obj})
    diags := diag.Diagnostics{}
    var model models.AutoVpnClustersGatewaysInnerInterfacesInner
    diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
    if diags.HasError() {
        tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
        return nil, diags
    }
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

    var sdk network_services.AutoVpnClustersGatewaysInnerInterfacesInner
    var d diag.Diagnostics
    // Handling Primitives
    if !model.DhcpIp.IsNull() && !model.DhcpIp.IsUnknown() {
        sdk.DhcpIp = model.DhcpIp.ValueStringPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "DhcpIp", "value": *sdk.DhcpIp})
    }

    // Handling Primitives
    if !model.Name.IsNull() && !model.Name.IsUnknown() {
        sdk.Name = model.Name.ValueStringPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
    }

    // Handling Objects
    if !model.SdwanLinkSettings.IsNull() && !model.SdwanLinkSettings.IsUnknown() {
        tflog.Debug(ctx, "Unpacking nested object for field SdwanLinkSettings")
        unpacked, d := unpackAutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettingsToSdk(ctx, model.SdwanLinkSettings)
        diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "SdwanLinkSettings"})
		}
        if unpacked != nil {
            sdk.SdwanLinkSettings = unpacked
        }
    }

diags.Append(d...)

    tflog.Debug(ctx, "Exiting unpack helper for models.AutoVpnClustersGatewaysInnerInterfacesInner", map[string]interface{}{"has_errors": diags.HasError()})
    return &sdk, diags

}

// --- Packer for AutoVpnClustersGatewaysInnerInterfacesInner ---
func packAutoVpnClustersGatewaysInnerInterfacesInnerFromSdk(ctx context.Context, sdk network_services.AutoVpnClustersGatewaysInnerInterfacesInner) (types.Object, diag.Diagnostics) {
    tflog.Debug(ctx, "Entering pack helper for models.AutoVpnClustersGatewaysInnerInterfacesInner", map[string]interface{}{"sdk_struct": sdk})
    diags := diag.Diagnostics{}
    var model models.AutoVpnClustersGatewaysInnerInterfacesInner
    var d diag.Diagnostics
    // Handling Primitives
    // Standard primitive packing
    if sdk.DhcpIp != nil {
        model.DhcpIp = basetypes.NewStringValue(*sdk.DhcpIp)
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "DhcpIp", "value": *sdk.DhcpIp})
    } else {
        model.DhcpIp = basetypes.NewStringNull()
    }
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
    if sdk.SdwanLinkSettings != nil {
        tflog.Debug(ctx, "Packing nested object for field SdwanLinkSettings")
        packed, d := packAutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettingsFromSdk(ctx, *sdk.SdwanLinkSettings)
        diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "SdwanLinkSettings"})
		}
        model.SdwanLinkSettings = packed
    } else {
        model.SdwanLinkSettings = basetypes.NewObjectNull(models.AutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettings{}.AttrTypes())
    }
diags.Append(d...)

    obj, d := types.ObjectValueFrom(ctx, models.AutoVpnClustersGatewaysInnerInterfacesInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
    diags.Append(d...)
    tflog.Debug(ctx, "Exiting pack helper for models.AutoVpnClustersGatewaysInnerInterfacesInner", map[string]interface{}{"has_errors": diags.HasError()})
    return obj, diags

}

// --- List Unpacker for AutoVpnClustersGatewaysInnerInterfacesInner ---
func unpackAutoVpnClustersGatewaysInnerInterfacesInnerListToSdk(ctx context.Context, list types.List) ([]network_services.AutoVpnClustersGatewaysInnerInterfacesInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.AutoVpnClustersGatewaysInnerInterfacesInner")
	diags := diag.Diagnostics{}
	var data []models.AutoVpnClustersGatewaysInnerInterfacesInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.AutoVpnClustersGatewaysInnerInterfacesInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.AutoVpnClustersGatewaysInnerInterfacesInner{}.AttrTypes(), &item)
		unpacked, d := unpackAutoVpnClustersGatewaysInnerInterfacesInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.AutoVpnClustersGatewaysInnerInterfacesInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for AutoVpnClustersGatewaysInnerInterfacesInner ---
func packAutoVpnClustersGatewaysInnerInterfacesInnerListFromSdk(ctx context.Context, sdks []network_services.AutoVpnClustersGatewaysInnerInterfacesInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.AutoVpnClustersGatewaysInnerInterfacesInner")
	diags := diag.Diagnostics{}
	var data []models.AutoVpnClustersGatewaysInnerInterfacesInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.AutoVpnClustersGatewaysInnerInterfacesInner
		obj, d := packAutoVpnClustersGatewaysInnerInterfacesInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.AutoVpnClustersGatewaysInnerInterfacesInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.AutoVpnClustersGatewaysInnerInterfacesInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.AutoVpnClustersGatewaysInnerInterfacesInner{}.AttrType(), data)
}

// --- Unpacker for AutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettings ---
func unpackAutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettingsToSdk(ctx context.Context, obj types.Object) (*network_services.AutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettings, diag.Diagnostics) {
    tflog.Debug(ctx, "Entering unpack helper for models.AutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettings", map[string]interface{}{"tf_object": obj})
    diags := diag.Diagnostics{}
    var model models.AutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettings
    diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
    if diags.HasError() {
        tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
        return nil, diags
    }
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

    var sdk network_services.AutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettings
    var d diag.Diagnostics
    // Handling Primitives
    if !model.SdwanGateway.IsNull() && !model.SdwanGateway.IsUnknown() {
        sdk.SdwanGateway = model.SdwanGateway.ValueStringPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "SdwanGateway", "value": *sdk.SdwanGateway})
    }

    // Handling Primitives
    if !model.SdwanInterfaceProfile.IsNull() && !model.SdwanInterfaceProfile.IsUnknown() {
        sdk.SdwanInterfaceProfile = model.SdwanInterfaceProfile.ValueStringPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "SdwanInterfaceProfile", "value": *sdk.SdwanInterfaceProfile})
    }

    // Handling Objects
    if !model.UpstreamNat.IsNull() && !model.UpstreamNat.IsUnknown() {
        tflog.Debug(ctx, "Unpacking nested object for field UpstreamNat")
        unpacked, d := unpackAutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettingsUpstreamNatToSdk(ctx, model.UpstreamNat)
        diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "UpstreamNat"})
		}
        if unpacked != nil {
            sdk.UpstreamNat = unpacked
        }
    }

diags.Append(d...)

    tflog.Debug(ctx, "Exiting unpack helper for models.AutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettings", map[string]interface{}{"has_errors": diags.HasError()})
    return &sdk, diags

}

// --- Packer for AutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettings ---
func packAutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettingsFromSdk(ctx context.Context, sdk network_services.AutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettings) (types.Object, diag.Diagnostics) {
    tflog.Debug(ctx, "Entering pack helper for models.AutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettings", map[string]interface{}{"sdk_struct": sdk})
    diags := diag.Diagnostics{}
    var model models.AutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettings
    var d diag.Diagnostics
    // Handling Primitives
    // Standard primitive packing
    if sdk.SdwanGateway != nil {
        model.SdwanGateway = basetypes.NewStringValue(*sdk.SdwanGateway)
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "SdwanGateway", "value": *sdk.SdwanGateway})
    } else {
        model.SdwanGateway = basetypes.NewStringNull()
    }
    // Handling Primitives
    // Standard primitive packing
    if sdk.SdwanInterfaceProfile != nil {
        model.SdwanInterfaceProfile = basetypes.NewStringValue(*sdk.SdwanInterfaceProfile)
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "SdwanInterfaceProfile", "value": *sdk.SdwanInterfaceProfile})
    } else {
        model.SdwanInterfaceProfile = basetypes.NewStringNull()
    }
    // Handling Objects
    // This is a regular nested object that has its own packer.
    if sdk.UpstreamNat != nil {
        tflog.Debug(ctx, "Packing nested object for field UpstreamNat")
        packed, d := packAutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettingsUpstreamNatFromSdk(ctx, *sdk.UpstreamNat)
        diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "UpstreamNat"})
		}
        model.UpstreamNat = packed
    } else {
        model.UpstreamNat = basetypes.NewObjectNull(models.AutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettingsUpstreamNat{}.AttrTypes())
    }
diags.Append(d...)

    obj, d := types.ObjectValueFrom(ctx, models.AutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettings{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
    diags.Append(d...)
    tflog.Debug(ctx, "Exiting pack helper for models.AutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettings", map[string]interface{}{"has_errors": diags.HasError()})
    return obj, diags

}

// --- List Unpacker for AutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettings ---
func unpackAutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettingsListToSdk(ctx context.Context, list types.List) ([]network_services.AutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettings, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.AutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettings")
	diags := diag.Diagnostics{}
	var data []models.AutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettings
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.AutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettings, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.AutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettings{}.AttrTypes(), &item)
		unpacked, d := unpackAutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettingsToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.AutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettings", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for AutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettings ---
func packAutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettingsListFromSdk(ctx context.Context, sdks []network_services.AutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettings) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.AutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettings")
	diags := diag.Diagnostics{}
	var data []models.AutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettings

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.AutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettings
		obj, d := packAutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettingsFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.AutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettings{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.AutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettings", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.AutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettings{}.AttrType(), data)
}

// --- Unpacker for AutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettingsUpstreamNat ---
func unpackAutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettingsUpstreamNatToSdk(ctx context.Context, obj types.Object) (*network_services.AutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettingsUpstreamNat, diag.Diagnostics) {
    tflog.Debug(ctx, "Entering unpack helper for models.AutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettingsUpstreamNat", map[string]interface{}{"tf_object": obj})
    diags := diag.Diagnostics{}
    var model models.AutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettingsUpstreamNat
    diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
    if diags.HasError() {
        tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
        return nil, diags
    }
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

    var sdk network_services.AutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettingsUpstreamNat
    var d diag.Diagnostics
    // Handling Primitives
    if !model.Enable.IsNull() && !model.Enable.IsUnknown() {
        sdk.Enable = model.Enable.ValueBoolPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Enable", "value": *sdk.Enable})
    }

    // Handling Objects
    if !model.StaticIp.IsNull() && !model.StaticIp.IsUnknown() {
        tflog.Debug(ctx, "Unpacking nested object for field StaticIp")
        unpacked, d := unpackAutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNatStaticIpToSdk(ctx, model.StaticIp)
        diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "StaticIp"})
		}
        if unpacked != nil {
            sdk.StaticIp = unpacked
        }
    }

diags.Append(d...)

    tflog.Debug(ctx, "Exiting unpack helper for models.AutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettingsUpstreamNat", map[string]interface{}{"has_errors": diags.HasError()})
    return &sdk, diags

}

// --- Packer for AutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettingsUpstreamNat ---
func packAutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettingsUpstreamNatFromSdk(ctx context.Context, sdk network_services.AutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettingsUpstreamNat) (types.Object, diag.Diagnostics) {
    tflog.Debug(ctx, "Entering pack helper for models.AutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettingsUpstreamNat", map[string]interface{}{"sdk_struct": sdk})
    diags := diag.Diagnostics{}
    var model models.AutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettingsUpstreamNat
    var d diag.Diagnostics
    // Handling Primitives
    // Standard primitive packing
    if sdk.Enable != nil {
        model.Enable = basetypes.NewBoolValue(*sdk.Enable)
        tflog.Debug(ctx, "Packed primitive pointer", map[string]interface{}{"field": "Enable", "value": *sdk.Enable})
    } else {
        model.Enable = basetypes.NewBoolNull()
    }
    // Handling Objects
    // This is a regular nested object that has its own packer.
    if sdk.StaticIp != nil {
        tflog.Debug(ctx, "Packing nested object for field StaticIp")
        packed, d := packAutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNatStaticIpFromSdk(ctx, *sdk.StaticIp)
        diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "StaticIp"})
		}
        model.StaticIp = packed
    } else {
        model.StaticIp = basetypes.NewObjectNull(models.AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNatStaticIp{}.AttrTypes())
    }
diags.Append(d...)

    obj, d := types.ObjectValueFrom(ctx, models.AutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettingsUpstreamNat{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
    diags.Append(d...)
    tflog.Debug(ctx, "Exiting pack helper for models.AutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettingsUpstreamNat", map[string]interface{}{"has_errors": diags.HasError()})
    return obj, diags

}

// --- List Unpacker for AutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettingsUpstreamNat ---
func unpackAutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettingsUpstreamNatListToSdk(ctx context.Context, list types.List) ([]network_services.AutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettingsUpstreamNat, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.AutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettingsUpstreamNat")
	diags := diag.Diagnostics{}
	var data []models.AutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettingsUpstreamNat
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.AutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettingsUpstreamNat, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.AutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettingsUpstreamNat{}.AttrTypes(), &item)
		unpacked, d := unpackAutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettingsUpstreamNatToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.AutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettingsUpstreamNat", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for AutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettingsUpstreamNat ---
func packAutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettingsUpstreamNatListFromSdk(ctx context.Context, sdks []network_services.AutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettingsUpstreamNat) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.AutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettingsUpstreamNat")
	diags := diag.Diagnostics{}
	var data []models.AutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettingsUpstreamNat

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.AutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettingsUpstreamNat
		obj, d := packAutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettingsUpstreamNatFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.AutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettingsUpstreamNat{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.AutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettingsUpstreamNat", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.AutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettingsUpstreamNat{}.AttrType(), data)
}

// --- Unpacker for AutoVpnClustersGatewaysInnerPrivateInterfacesInner ---
func unpackAutoVpnClustersGatewaysInnerPrivateInterfacesInnerToSdk(ctx context.Context, obj types.Object) (*network_services.AutoVpnClustersGatewaysInnerPrivateInterfacesInner, diag.Diagnostics) {
    tflog.Debug(ctx, "Entering unpack helper for models.AutoVpnClustersGatewaysInnerPrivateInterfacesInner", map[string]interface{}{"tf_object": obj})
    diags := diag.Diagnostics{}
    var model models.AutoVpnClustersGatewaysInnerPrivateInterfacesInner
    diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
    if diags.HasError() {
        tflog.Error(ctx, "Error converting Terraform object to Go model", map[string]interface{}{"diags": diags})
        return nil, diags
    }
	tflog.Debug(ctx, "Successfully converted Terraform object to Go model")

    var sdk network_services.AutoVpnClustersGatewaysInnerPrivateInterfacesInner
    var d diag.Diagnostics
    // Handling Primitives
    if !model.Name.IsNull() && !model.Name.IsUnknown() {
        sdk.Name = model.Name.ValueStringPointer()
        tflog.Debug(ctx, "Unpacked primitive pointer", map[string]interface{}{"field": "Name", "value": *sdk.Name})
    }

    // Handling Objects
    if !model.SdwanLinkSettings.IsNull() && !model.SdwanLinkSettings.IsUnknown() {
        tflog.Debug(ctx, "Unpacking nested object for field SdwanLinkSettings")
        unpacked, d := unpackAutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettingsToSdk(ctx, model.SdwanLinkSettings)
        diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error unpacking nested object", map[string]interface{}{"field": "SdwanLinkSettings"})
		}
        if unpacked != nil {
            sdk.SdwanLinkSettings = unpacked
        }
    }

diags.Append(d...)

    tflog.Debug(ctx, "Exiting unpack helper for models.AutoVpnClustersGatewaysInnerPrivateInterfacesInner", map[string]interface{}{"has_errors": diags.HasError()})
    return &sdk, diags

}

// --- Packer for AutoVpnClustersGatewaysInnerPrivateInterfacesInner ---
func packAutoVpnClustersGatewaysInnerPrivateInterfacesInnerFromSdk(ctx context.Context, sdk network_services.AutoVpnClustersGatewaysInnerPrivateInterfacesInner) (types.Object, diag.Diagnostics) {
    tflog.Debug(ctx, "Entering pack helper for models.AutoVpnClustersGatewaysInnerPrivateInterfacesInner", map[string]interface{}{"sdk_struct": sdk})
    diags := diag.Diagnostics{}
    var model models.AutoVpnClustersGatewaysInnerPrivateInterfacesInner
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
    if sdk.SdwanLinkSettings != nil {
        tflog.Debug(ctx, "Packing nested object for field SdwanLinkSettings")
        packed, d := packAutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettingsFromSdk(ctx, *sdk.SdwanLinkSettings)
        diags.Append(d...)
		if d.HasError() {
			tflog.Error(ctx, "Error packing nested object", map[string]interface{}{"field": "SdwanLinkSettings"})
		}
        model.SdwanLinkSettings = packed
    } else {
        model.SdwanLinkSettings = basetypes.NewObjectNull(models.AutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettings{}.AttrTypes())
    }
diags.Append(d...)

    obj, d := types.ObjectValueFrom(ctx, models.AutoVpnClustersGatewaysInnerPrivateInterfacesInner{}.AttrTypes(), &model)
	tflog.Debug(ctx, "Final object to be returned from pack helper", map[string]interface{}{"object": obj})
    diags.Append(d...)
    tflog.Debug(ctx, "Exiting pack helper for models.AutoVpnClustersGatewaysInnerPrivateInterfacesInner", map[string]interface{}{"has_errors": diags.HasError()})
    return obj, diags

}

// --- List Unpacker for AutoVpnClustersGatewaysInnerPrivateInterfacesInner ---
func unpackAutoVpnClustersGatewaysInnerPrivateInterfacesInnerListToSdk(ctx context.Context, list types.List) ([]network_services.AutoVpnClustersGatewaysInnerPrivateInterfacesInner, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list unpack helper for models.AutoVpnClustersGatewaysInnerPrivateInterfacesInner")
	diags := diag.Diagnostics{}
	var data []models.AutoVpnClustersGatewaysInnerPrivateInterfacesInner
	diags.Append(list.ElementsAs(ctx, &data, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Error converting list elements to Go models", map[string]interface{}{"diags": diags})
		return nil, diags
	}

	ans := make([]network_services.AutoVpnClustersGatewaysInnerPrivateInterfacesInner, 0, len(data))
	for i, item := range data {
		tflog.Debug(ctx, "Unpacking item from list", map[string]interface{}{"index": i})
		obj, _ := types.ObjectValueFrom(ctx, models.AutoVpnClustersGatewaysInnerPrivateInterfacesInner{}.AttrTypes(), &item)
		unpacked, d := unpackAutoVpnClustersGatewaysInnerPrivateInterfacesInnerToSdk(ctx, obj)
		diags.Append(d...)
		if unpacked != nil {
			ans = append(ans, *unpacked)
		}
	}
	tflog.Debug(ctx, "Exiting list unpack helper for models.AutoVpnClustersGatewaysInnerPrivateInterfacesInner", map[string]interface{}{"has_errors": diags.HasError()})
	return ans, diags
}

// --- List Packer for AutoVpnClustersGatewaysInnerPrivateInterfacesInner ---
func packAutoVpnClustersGatewaysInnerPrivateInterfacesInnerListFromSdk(ctx context.Context, sdks []network_services.AutoVpnClustersGatewaysInnerPrivateInterfacesInner) (types.List, diag.Diagnostics) {
	tflog.Debug(ctx, "Entering list pack helper for models.AutoVpnClustersGatewaysInnerPrivateInterfacesInner")
	diags := diag.Diagnostics{}
	var data []models.AutoVpnClustersGatewaysInnerPrivateInterfacesInner

	for i, sdk := range sdks {
		tflog.Debug(ctx, "Packing item to list", map[string]interface{}{"index": i})
		var model models.AutoVpnClustersGatewaysInnerPrivateInterfacesInner
		obj, d := packAutoVpnClustersGatewaysInnerPrivateInterfacesInnerFromSdk(ctx, sdk)
		diags.Append(d...)
		if diags.HasError() {
			return basetypes.NewListNull(models.AutoVpnClustersGatewaysInnerPrivateInterfacesInner{}.AttrType()), diags
		}
		diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
		data = append(data, model)
	}
	tflog.Debug(ctx, "Exiting list pack helper for models.AutoVpnClustersGatewaysInnerPrivateInterfacesInner", map[string]interface{}{"has_errors": diags.HasError()})
	return basetypes.NewListValueFrom(ctx, models.AutoVpnClustersGatewaysInnerPrivateInterfacesInner{}.AttrType(), data)
}


*/
