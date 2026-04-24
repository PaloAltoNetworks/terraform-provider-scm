package models

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	dsschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/paloaltonetworks/terraform-provider-scm/internal/utils"
)

// Package: config_setup
// This file contains models for the config_setup SDK package

// Devices represents the Terraform model for Devices
type Devices struct {
	Tfid               types.String          `tfsdk:"tfid"`
	AntiVirusVersion   basetypes.StringValue `tfsdk:"anti_virus_version"`
	AppReleaseDate     basetypes.StringValue `tfsdk:"app_release_date"`
	AppVersion         basetypes.StringValue `tfsdk:"app_version"`
	AvReleaseDate      basetypes.StringValue `tfsdk:"av_release_date"`
	AvailableLicensess basetypes.ListValue   `tfsdk:"available_licensess"`
	ConnectedSince     basetypes.StringValue `tfsdk:"connected_since"`
	Description        basetypes.StringValue `tfsdk:"description"`
	DevCertDetail      basetypes.StringValue `tfsdk:"dev_cert_detail"`
	DevCertExpiryDate  basetypes.StringValue `tfsdk:"dev_cert_expiry_date"`
	DisplayName        basetypes.StringValue `tfsdk:"display_name"`
	Family             basetypes.StringValue `tfsdk:"family"`
	Folder             basetypes.StringValue `tfsdk:"folder"`
	GpClientVerion     basetypes.StringValue `tfsdk:"gp_client_verion"`
	GpDataVersion      basetypes.StringValue `tfsdk:"gp_data_version"`
	HaPeerSerial       basetypes.StringValue `tfsdk:"ha_peer_serial"`
	HaPeerState        basetypes.StringValue `tfsdk:"ha_peer_state"`
	HaState            basetypes.StringValue `tfsdk:"ha_state"`
	Hostname           basetypes.StringValue `tfsdk:"hostname"`
	Id                 basetypes.StringValue `tfsdk:"id"`
	InstalledLicenses  basetypes.ListValue   `tfsdk:"installed_licenses"`
	IotReleaseDate     basetypes.StringValue `tfsdk:"iot_release_date"`
	IotVersion         basetypes.StringValue `tfsdk:"iot_version"`
	IpAddress          basetypes.StringValue `tfsdk:"ip_address"`
	IpV6Address        basetypes.StringValue `tfsdk:"ip_v6_address"`
	IsConnected        basetypes.BoolValue   `tfsdk:"is_connected"`
	Labels             basetypes.ListValue   `tfsdk:"labels"`
	LicenseMatch       basetypes.BoolValue   `tfsdk:"license_match"`
	LogDbVersion       basetypes.StringValue `tfsdk:"log_db_version"`
	MacAddress         basetypes.StringValue `tfsdk:"mac_address"`
	Model              basetypes.StringValue `tfsdk:"model"`
	Name               basetypes.StringValue `tfsdk:"name"`
	Snippets           basetypes.ListValue   `tfsdk:"snippets"`
	SoftwareVersion    basetypes.StringValue `tfsdk:"software_version"`
	ThreatReleaseDate  basetypes.StringValue `tfsdk:"threat_release_date"`
	ThreatVersion      basetypes.StringValue `tfsdk:"threat_version"`
	Uptime             basetypes.StringValue `tfsdk:"uptime"`
	UrlDbType          basetypes.StringValue `tfsdk:"url_db_type"`
	UrlDbVer           basetypes.StringValue `tfsdk:"url_db_ver"`
	VmState            basetypes.StringValue `tfsdk:"vm_state"`
	WfReleaseDate      basetypes.StringValue `tfsdk:"wf_release_date"`
	WfVer              basetypes.StringValue `tfsdk:"wf_ver"`
}

// DevicesAvailableLicensessInner represents a nested structure within the Devices model
type DevicesAvailableLicensessInner struct {
	Authcode basetypes.StringValue `tfsdk:"authcode"`
	Expires  basetypes.StringValue `tfsdk:"expires"`
	Feature  basetypes.StringValue `tfsdk:"feature"`
	Issued   basetypes.StringValue `tfsdk:"issued"`
}

// DevicesInstalledLicensesInner represents a nested structure within the Devices model
type DevicesInstalledLicensesInner struct {
	Authcode basetypes.StringValue `tfsdk:"authcode"`
	Expired  basetypes.StringValue `tfsdk:"expired"`
	Expires  basetypes.StringValue `tfsdk:"expires"`
	Feature  basetypes.StringValue `tfsdk:"feature"`
	Issued   basetypes.StringValue `tfsdk:"issued"`
}

// AttrTypes defines the attribute types for the Devices model.
func (o Devices) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":               basetypes.StringType{},
		"anti_virus_version": basetypes.StringType{},
		"app_release_date":   basetypes.StringType{},
		"app_version":        basetypes.StringType{},
		"av_release_date":    basetypes.StringType{},
		"available_licensess": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"authcode": basetypes.StringType{},
				"expires":  basetypes.StringType{},
				"feature":  basetypes.StringType{},
				"issued":   basetypes.StringType{},
			},
		}},
		"connected_since":      basetypes.StringType{},
		"description":          basetypes.StringType{},
		"dev_cert_detail":      basetypes.StringType{},
		"dev_cert_expiry_date": basetypes.StringType{},
		"display_name":         basetypes.StringType{},
		"family":               basetypes.StringType{},
		"folder":               basetypes.StringType{},
		"gp_client_verion":     basetypes.StringType{},
		"gp_data_version":      basetypes.StringType{},
		"ha_peer_serial":       basetypes.StringType{},
		"ha_peer_state":        basetypes.StringType{},
		"ha_state":             basetypes.StringType{},
		"hostname":             basetypes.StringType{},
		"id":                   basetypes.StringType{},
		"installed_licenses": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"authcode": basetypes.StringType{},
				"expired":  basetypes.StringType{},
				"expires":  basetypes.StringType{},
				"feature":  basetypes.StringType{},
				"issued":   basetypes.StringType{},
			},
		}},
		"iot_release_date":    basetypes.StringType{},
		"iot_version":         basetypes.StringType{},
		"ip_address":          basetypes.StringType{},
		"ip_v6_address":       basetypes.StringType{},
		"is_connected":        basetypes.BoolType{},
		"labels":              basetypes.ListType{ElemType: basetypes.StringType{}},
		"license_match":       basetypes.BoolType{},
		"log_db_version":      basetypes.StringType{},
		"mac_address":         basetypes.StringType{},
		"model":               basetypes.StringType{},
		"name":                basetypes.StringType{},
		"snippets":            basetypes.ListType{ElemType: basetypes.StringType{}},
		"software_version":    basetypes.StringType{},
		"threat_release_date": basetypes.StringType{},
		"threat_version":      basetypes.StringType{},
		"uptime":              basetypes.StringType{},
		"url_db_type":         basetypes.StringType{},
		"url_db_ver":          basetypes.StringType{},
		"vm_state":            basetypes.StringType{},
		"wf_release_date":     basetypes.StringType{},
		"wf_ver":              basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of Devices objects.
func (o Devices) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the DevicesAvailableLicensessInner model.
func (o DevicesAvailableLicensessInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"authcode": basetypes.StringType{},
		"expires":  basetypes.StringType{},
		"feature":  basetypes.StringType{},
		"issued":   basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of DevicesAvailableLicensessInner objects.
func (o DevicesAvailableLicensessInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the DevicesInstalledLicensesInner model.
func (o DevicesInstalledLicensesInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"authcode": basetypes.StringType{},
		"expired":  basetypes.StringType{},
		"expires":  basetypes.StringType{},
		"feature":  basetypes.StringType{},
		"issued":   basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of DevicesInstalledLicensesInner objects.
func (o DevicesInstalledLicensesInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// DevicesResourceSchema defines the schema for Devices resource
var DevicesResourceSchema = schema.Schema{
	MarkdownDescription: "Device resource",
	Attributes: map[string]schema.Attribute{
		"anti_virus_version": schema.StringAttribute{
			MarkdownDescription: "Anti virus version",
			Computed:            true,
		},
		"app_release_date": schema.StringAttribute{
			MarkdownDescription: "App release date",
			Computed:            true,
		},
		"app_version": schema.StringAttribute{
			MarkdownDescription: "App version",
			Computed:            true,
		},
		"av_release_date": schema.StringAttribute{
			MarkdownDescription: "Av release date",
			Computed:            true,
		},
		"available_licensess": schema.ListNestedAttribute{
			MarkdownDescription: "Available licensess",
			Computed:            true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"authcode": schema.StringAttribute{
						MarkdownDescription: "Authcode",
						Computed:            true,
					},
					"expires": schema.StringAttribute{
						MarkdownDescription: "Expires",
						Computed:            true,
					},
					"feature": schema.StringAttribute{
						MarkdownDescription: "Feature",
						Computed:            true,
					},
					"issued": schema.StringAttribute{
						MarkdownDescription: "Issued",
						Computed:            true,
					},
				},
			},
		},
		"connected_since": schema.StringAttribute{
			MarkdownDescription: "Connected since",
			Computed:            true,
		},
		"description": schema.StringAttribute{
			MarkdownDescription: "The description of the device",
			Optional:            true,
		},
		"dev_cert_detail": schema.StringAttribute{
			MarkdownDescription: "Dev cert detail",
			Computed:            true,
		},
		"dev_cert_expiry_date": schema.StringAttribute{
			MarkdownDescription: "Dev cert expiry date",
			Computed:            true,
		},
		"display_name": schema.StringAttribute{
			MarkdownDescription: "The display name of the device",
			Optional:            true,
		},
		"family": schema.StringAttribute{
			MarkdownDescription: "The product family of the device",
			Computed:            true,
		},
		"folder": schema.StringAttribute{
			Validators: []validator.String{
				utils.FolderValidator(),
			},
			MarkdownDescription: "The folder containing the device",
			Required:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"gp_client_verion": schema.StringAttribute{
			MarkdownDescription: "Gp client verion",
			Computed:            true,
		},
		"gp_data_version": schema.StringAttribute{
			MarkdownDescription: "Gp data version",
			Computed:            true,
		},
		"ha_peer_serial": schema.StringAttribute{
			MarkdownDescription: "Ha peer serial",
			Computed:            true,
		},
		"ha_peer_state": schema.StringAttribute{
			MarkdownDescription: "Ha peer state",
			Computed:            true,
		},
		"ha_state": schema.StringAttribute{
			MarkdownDescription: "Ha state",
			Computed:            true,
		},
		"hostname": schema.StringAttribute{
			MarkdownDescription: "The hostname of the device",
			Computed:            true,
		},
		"id": schema.StringAttribute{
			MarkdownDescription: "The UUID of the device",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"installed_licenses": schema.ListNestedAttribute{
			MarkdownDescription: "Installed licenses",
			Computed:            true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"authcode": schema.StringAttribute{
						MarkdownDescription: "Authcode",
						Computed:            true,
					},
					"expired": schema.StringAttribute{
						MarkdownDescription: "Expired",
						Computed:            true,
					},
					"expires": schema.StringAttribute{
						MarkdownDescription: "Expires",
						Computed:            true,
					},
					"feature": schema.StringAttribute{
						MarkdownDescription: "Feature",
						Computed:            true,
					},
					"issued": schema.StringAttribute{
						MarkdownDescription: "Issued",
						Computed:            true,
					},
				},
			},
		},
		"iot_release_date": schema.StringAttribute{
			MarkdownDescription: "Iot release date",
			Computed:            true,
		},
		"iot_version": schema.StringAttribute{
			MarkdownDescription: "Iot version",
			Computed:            true,
		},
		"ip_address": schema.StringAttribute{
			MarkdownDescription: "The IPv4 address of the device",
			Computed:            true,
		},
		"ip_v6_address": schema.StringAttribute{
			MarkdownDescription: "Ip v6 address",
			Optional:            true,
		},
		"is_connected": schema.BoolAttribute{
			MarkdownDescription: "Is connected",
			Computed:            true,
		},
		"labels": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Labels assigned to the device",
			Optional:            true,
		},
		"license_match": schema.BoolAttribute{
			MarkdownDescription: "License match",
			Computed:            true,
		},
		"log_db_version": schema.StringAttribute{
			MarkdownDescription: "Log db version",
			Computed:            true,
		},
		"mac_address": schema.StringAttribute{
			MarkdownDescription: "The MAC address of the device",
			Computed:            true,
		},
		"model": schema.StringAttribute{
			MarkdownDescription: "The model of the device",
			Computed:            true,
		},
		"name": schema.StringAttribute{
			MarkdownDescription: "The name of the device",
			Required:            true,
		},
		"snippets": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Snippets associated with the device",
			Optional:            true,
		},
		"software_version": schema.StringAttribute{
			MarkdownDescription: "Software version",
			Computed:            true,
		},
		"tfid": schema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"threat_release_date": schema.StringAttribute{
			MarkdownDescription: "Threat release date",
			Computed:            true,
		},
		"threat_version": schema.StringAttribute{
			MarkdownDescription: "Threat version",
			Computed:            true,
		},
		"uptime": schema.StringAttribute{
			MarkdownDescription: "Uptime",
			Computed:            true,
		},
		"url_db_type": schema.StringAttribute{
			MarkdownDescription: "Url db type",
			Computed:            true,
		},
		"url_db_ver": schema.StringAttribute{
			MarkdownDescription: "Url db ver",
			Computed:            true,
		},
		"vm_state": schema.StringAttribute{
			MarkdownDescription: "Vm state",
			Computed:            true,
		},
		"wf_release_date": schema.StringAttribute{
			MarkdownDescription: "Wf release date",
			Computed:            true,
		},
		"wf_ver": schema.StringAttribute{
			MarkdownDescription: "Wf ver",
			Computed:            true,
		},
	},
}

// DevicesDataSourceSchema defines the schema for Devices data source
var DevicesDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Device data source",
	Attributes: map[string]dsschema.Attribute{
		"anti_virus_version": dsschema.StringAttribute{
			MarkdownDescription: "Anti virus version",
			Computed:            true,
		},
		"app_release_date": dsschema.StringAttribute{
			MarkdownDescription: "App release date",
			Computed:            true,
		},
		"app_version": dsschema.StringAttribute{
			MarkdownDescription: "App version",
			Computed:            true,
		},
		"av_release_date": dsschema.StringAttribute{
			MarkdownDescription: "Av release date",
			Computed:            true,
		},
		"available_licensess": dsschema.ListNestedAttribute{
			MarkdownDescription: "Available licensess",
			Computed:            true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: map[string]dsschema.Attribute{
					"authcode": dsschema.StringAttribute{
						MarkdownDescription: "Authcode",
						Computed:            true,
					},
					"expires": dsschema.StringAttribute{
						MarkdownDescription: "Expires",
						Computed:            true,
					},
					"feature": dsschema.StringAttribute{
						MarkdownDescription: "Feature",
						Computed:            true,
					},
					"issued": dsschema.StringAttribute{
						MarkdownDescription: "Issued",
						Computed:            true,
					},
				},
			},
		},
		"connected_since": dsschema.StringAttribute{
			MarkdownDescription: "Connected since",
			Computed:            true,
		},
		"description": dsschema.StringAttribute{
			MarkdownDescription: "The description of the device",
			Computed:            true,
		},
		"dev_cert_detail": dsschema.StringAttribute{
			MarkdownDescription: "Dev cert detail",
			Computed:            true,
		},
		"dev_cert_expiry_date": dsschema.StringAttribute{
			MarkdownDescription: "Dev cert expiry date",
			Computed:            true,
		},
		"display_name": dsschema.StringAttribute{
			MarkdownDescription: "The display name of the device",
			Computed:            true,
		},
		"family": dsschema.StringAttribute{
			MarkdownDescription: "The product family of the device",
			Computed:            true,
		},
		"folder": dsschema.StringAttribute{
			MarkdownDescription: "The folder containing the device",
			Optional:            true,
			Computed:            true,
		},
		"gp_client_verion": dsschema.StringAttribute{
			MarkdownDescription: "Gp client verion",
			Computed:            true,
		},
		"gp_data_version": dsschema.StringAttribute{
			MarkdownDescription: "Gp data version",
			Computed:            true,
		},
		"ha_peer_serial": dsschema.StringAttribute{
			MarkdownDescription: "Ha peer serial",
			Computed:            true,
		},
		"ha_peer_state": dsschema.StringAttribute{
			MarkdownDescription: "Ha peer state",
			Computed:            true,
		},
		"ha_state": dsschema.StringAttribute{
			MarkdownDescription: "Ha state",
			Computed:            true,
		},
		"hostname": dsschema.StringAttribute{
			MarkdownDescription: "The hostname of the device",
			Computed:            true,
		},
		"id": dsschema.StringAttribute{
			MarkdownDescription: "The UUID of the device",
			Required:            true,
		},
		"installed_licenses": dsschema.ListNestedAttribute{
			MarkdownDescription: "Installed licenses",
			Computed:            true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: map[string]dsschema.Attribute{
					"authcode": dsschema.StringAttribute{
						MarkdownDescription: "Authcode",
						Computed:            true,
					},
					"expired": dsschema.StringAttribute{
						MarkdownDescription: "Expired",
						Computed:            true,
					},
					"expires": dsschema.StringAttribute{
						MarkdownDescription: "Expires",
						Computed:            true,
					},
					"feature": dsschema.StringAttribute{
						MarkdownDescription: "Feature",
						Computed:            true,
					},
					"issued": dsschema.StringAttribute{
						MarkdownDescription: "Issued",
						Computed:            true,
					},
				},
			},
		},
		"iot_release_date": dsschema.StringAttribute{
			MarkdownDescription: "Iot release date",
			Computed:            true,
		},
		"iot_version": dsschema.StringAttribute{
			MarkdownDescription: "Iot version",
			Computed:            true,
		},
		"ip_address": dsschema.StringAttribute{
			MarkdownDescription: "The IPv4 address of the device",
			Computed:            true,
		},
		"ip_v6_address": dsschema.StringAttribute{
			MarkdownDescription: "Ip v6 address",
			Computed:            true,
		},
		"is_connected": dsschema.BoolAttribute{
			MarkdownDescription: "Is connected",
			Computed:            true,
		},
		"labels": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Labels assigned to the device",
			Computed:            true,
		},
		"license_match": dsschema.BoolAttribute{
			MarkdownDescription: "License match",
			Computed:            true,
		},
		"log_db_version": dsschema.StringAttribute{
			MarkdownDescription: "Log db version",
			Computed:            true,
		},
		"mac_address": dsschema.StringAttribute{
			MarkdownDescription: "The MAC address of the device",
			Computed:            true,
		},
		"model": dsschema.StringAttribute{
			MarkdownDescription: "The model of the device",
			Computed:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "The name of the device",
			Optional:            true,
			Computed:            true,
		},
		"snippets": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Snippets associated with the device",
			Computed:            true,
		},
		"software_version": dsschema.StringAttribute{
			MarkdownDescription: "Software version",
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
		"threat_release_date": dsschema.StringAttribute{
			MarkdownDescription: "Threat release date",
			Computed:            true,
		},
		"threat_version": dsschema.StringAttribute{
			MarkdownDescription: "Threat version",
			Computed:            true,
		},
		"uptime": dsschema.StringAttribute{
			MarkdownDescription: "Uptime",
			Computed:            true,
		},
		"url_db_type": dsschema.StringAttribute{
			MarkdownDescription: "Url db type",
			Computed:            true,
		},
		"url_db_ver": dsschema.StringAttribute{
			MarkdownDescription: "Url db ver",
			Computed:            true,
		},
		"vm_state": dsschema.StringAttribute{
			MarkdownDescription: "Vm state",
			Computed:            true,
		},
		"wf_release_date": dsschema.StringAttribute{
			MarkdownDescription: "Wf release date",
			Computed:            true,
		},
		"wf_ver": dsschema.StringAttribute{
			MarkdownDescription: "Wf ver",
			Computed:            true,
		},
	},
}

// DevicesListModel represents the data model for a list data source.
type DevicesListModel struct {
	Tfid       types.String        `tfsdk:"tfid"`
	Data       []Devices           `tfsdk:"data"`
	Limit      types.Int64         `tfsdk:"limit"`
	Offset     types.Int64         `tfsdk:"offset"`
	Name       types.String        `tfsdk:"name"`
	Total      types.Int64         `tfsdk:"total"`
	Folder     types.String        `tfsdk:"folder"`
	Snippet    types.String        `tfsdk:"snippet"`
	Device     types.String        `tfsdk:"device"`
	Pagination basetypes.BoolValue `tfsdk:"pagination"`
}

// DevicesListDataSourceSchema defines the schema for a list data source.
var DevicesListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: DevicesDataSourceSchema.Attributes,
			},
		},
		"limit":   dsschema.Int64Attribute{Description: "The max number of items to return. Default: 200.", Optional: true},
		"offset":  dsschema.Int64Attribute{Description: "The offset of the first item to return.", Optional: true},
		"name":    dsschema.StringAttribute{Description: "The name of the item.", Optional: true},
		"total":   dsschema.Int64Attribute{Description: "The total number of items.", Computed: true},
		"folder":  dsschema.StringAttribute{Description: "The folder of the item. Default: Shared.", Optional: true},
		"snippet": dsschema.StringAttribute{Description: "The snippet of the item.", Optional: true},
		"device":  dsschema.StringAttribute{Description: "The device of the item.", Optional: true},
		"pagination": dsschema.BoolAttribute{
			Description: "The parameter to mention if the response should be paginated. By default, its set to false",
			Optional:    true,
		},
	},
}
