package models

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	dsschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/paloaltonetworks/terraform-provider-scm/internal/utils"
)

// Package: mobile_agent
// This file contains models for the mobile_agent SDK package

// ForwardingProfileRegionalAndCustomProxies represents the Terraform model for ForwardingProfileRegionalAndCustomProxies
type ForwardingProfileRegionalAndCustomProxies struct {
	Tfid                   types.String          `tfsdk:"tfid"`
	ConnectivityPreference basetypes.ListValue   `tfsdk:"connectivity_preference"`
	Description            basetypes.StringValue `tfsdk:"description"`
	FallbackOption         basetypes.StringValue `tfsdk:"fallback_option"`
	Id                     basetypes.StringValue `tfsdk:"id"`
	LocationPreference     basetypes.StringValue `tfsdk:"location_preference"`
	Name                   basetypes.StringValue `tfsdk:"name"`
	PrismaAccessLocations  basetypes.ListValue   `tfsdk:"prisma_access_locations"`
	Proxy1                 basetypes.ObjectValue `tfsdk:"proxy_1"`
	Proxy2                 basetypes.ObjectValue `tfsdk:"proxy_2"`
	Type                   basetypes.StringValue `tfsdk:"type"`
	Folder                 basetypes.StringValue `tfsdk:"folder"`
}

// ForwardingProfileRegionalAndCustomProxiesConnectivityPreferenceInner represents a nested structure within the ForwardingProfileRegionalAndCustomProxies model
type ForwardingProfileRegionalAndCustomProxiesConnectivityPreferenceInner struct {
	Enabled basetypes.BoolValue   `tfsdk:"enabled"`
	Name    basetypes.StringValue `tfsdk:"name"`
}

// ForwardingProfileRegionalAndCustomProxiesPrismaAccessLocationsInner represents a nested structure within the ForwardingProfileRegionalAndCustomProxies model
type ForwardingProfileRegionalAndCustomProxiesPrismaAccessLocationsInner struct {
	Locations basetypes.ListValue   `tfsdk:"locations"`
	Name      basetypes.StringValue `tfsdk:"name"`
}

// ForwardingProfileRegionalAndCustomProxiesProxy1 represents a nested structure within the ForwardingProfileRegionalAndCustomProxies model
type ForwardingProfileRegionalAndCustomProxiesProxy1 struct {
	Fqdn     basetypes.StringValue `tfsdk:"fqdn"`
	Location basetypes.StringValue `tfsdk:"location"`
	Port     basetypes.Int64Value  `tfsdk:"port"`
}

// ForwardingProfileRegionalAndCustomProxiesProxy2 represents a nested structure within the ForwardingProfileRegionalAndCustomProxies model
type ForwardingProfileRegionalAndCustomProxiesProxy2 struct {
	Fqdn     basetypes.StringValue `tfsdk:"fqdn"`
	Location basetypes.StringValue `tfsdk:"location"`
	Port     basetypes.Int64Value  `tfsdk:"port"`
}

// AttrTypes defines the attribute types for the ForwardingProfileRegionalAndCustomProxies model.
func (o ForwardingProfileRegionalAndCustomProxies) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid": basetypes.StringType{},
		"connectivity_preference": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"enabled": basetypes.BoolType{},
				"name":    basetypes.StringType{},
			},
		}},
		"description":         basetypes.StringType{},
		"fallback_option":     basetypes.StringType{},
		"id":                  basetypes.StringType{},
		"location_preference": basetypes.StringType{},
		"name":                basetypes.StringType{},
		"prisma_access_locations": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"locations": basetypes.ListType{ElemType: basetypes.StringType{}},
				"name":      basetypes.StringType{},
			},
		}},
		"proxy_1": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"fqdn":     basetypes.StringType{},
				"location": basetypes.StringType{},
				"port":     basetypes.Int64Type{},
			},
		},
		"proxy_2": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"fqdn":     basetypes.StringType{},
				"location": basetypes.StringType{},
				"port":     basetypes.Int64Type{},
			},
		},
		"type":   basetypes.StringType{},
		"folder": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of ForwardingProfileRegionalAndCustomProxies objects.
func (o ForwardingProfileRegionalAndCustomProxies) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ForwardingProfileRegionalAndCustomProxiesConnectivityPreferenceInner model.
func (o ForwardingProfileRegionalAndCustomProxiesConnectivityPreferenceInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"enabled": basetypes.BoolType{},
		"name":    basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of ForwardingProfileRegionalAndCustomProxiesConnectivityPreferenceInner objects.
func (o ForwardingProfileRegionalAndCustomProxiesConnectivityPreferenceInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ForwardingProfileRegionalAndCustomProxiesPrismaAccessLocationsInner model.
func (o ForwardingProfileRegionalAndCustomProxiesPrismaAccessLocationsInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"locations": basetypes.ListType{ElemType: basetypes.StringType{}},
		"name":      basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of ForwardingProfileRegionalAndCustomProxiesPrismaAccessLocationsInner objects.
func (o ForwardingProfileRegionalAndCustomProxiesPrismaAccessLocationsInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ForwardingProfileRegionalAndCustomProxiesProxy1 model.
func (o ForwardingProfileRegionalAndCustomProxiesProxy1) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"fqdn":     basetypes.StringType{},
		"location": basetypes.StringType{},
		"port":     basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of ForwardingProfileRegionalAndCustomProxiesProxy1 objects.
func (o ForwardingProfileRegionalAndCustomProxiesProxy1) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ForwardingProfileRegionalAndCustomProxiesProxy2 model.
func (o ForwardingProfileRegionalAndCustomProxiesProxy2) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"fqdn":     basetypes.StringType{},
		"location": basetypes.StringType{},
		"port":     basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of ForwardingProfileRegionalAndCustomProxiesProxy2 objects.
func (o ForwardingProfileRegionalAndCustomProxiesProxy2) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// ForwardingProfileRegionalAndCustomProxiesResourceSchema defines the schema for ForwardingProfileRegionalAndCustomProxies resource
var ForwardingProfileRegionalAndCustomProxiesResourceSchema = schema.Schema{
	MarkdownDescription: "ForwardingProfileRegionalAndCustomProxy resource",
	Attributes: map[string]schema.Attribute{
		"connectivity_preference": schema.ListNestedAttribute{
			MarkdownDescription: "List of connectivity methods and their enablement status for establishing proxy connections",
			Optional:            true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"enabled": schema.BoolAttribute{
						MarkdownDescription: "Indicates whether this connectivity method is enabled for use in the proxy configuration",
						Optional:            true,
						Computed:            true,
						Default:             booldefault.StaticBool(false),
					},
					"name": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("tunnel", "proxy", "adns", "masque"),
						},
						MarkdownDescription: "Connectivity method type - 'tunnel' for VPN tunnels, 'proxy' for HTTP/HTTPS proxies, 'adns' for authenticated DNS, 'masque' for MASQUE protocol",
						Required:            true,
					},
				},
			},
		},
		"description": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(1023),
			},
			MarkdownDescription: "regional and custom proxy configuration description",
			Optional:            true,
		},
		"fallback_option": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.OneOf("fail-open", "fail-safe"),
			},
			MarkdownDescription: "Behavior when proxy connection fails - 'fail-open' allows direct internet access, 'fail-safe' blocks traffic until proxy is restored",
			Optional:            true,
		},
		"folder": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.OneOf("Mobile Users"),
				utils.FolderValidator(),
			},
			MarkdownDescription: "The folder in which the resource is defined\n",
			Optional:            true,
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			},
		},
		"id": schema.StringAttribute{
			MarkdownDescription: "The UUID of the regional and custom proxy",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"location_preference": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.OneOf("best-available-pa-location", "specific-pa-location"),
			},
			MarkdownDescription: "Strategy for selecting Prisma Access location - 'best-available-pa-location' automatically selects optimal location, 'specific-pa-location' uses predefined locations",
			Optional:            true,
		},
		"name": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(64),
				stringvalidator.RegexMatches(regexp.MustCompile("^[0-9a-zA-Z ._-]+$"), "pattern must match "+"^[0-9a-zA-Z ._-]+$"),
			},
			MarkdownDescription: "alphanumeric string [ 0-9a-zA-Z ._-]",
			Required:            true,
		},
		"prisma_access_locations": schema.ListNestedAttribute{
			MarkdownDescription: "Select Prisma Access location Americas, Europe and Asia-Pacific.",
			Optional:            true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"locations": schema.ListAttribute{
						ElementType:         types.StringType,
						MarkdownDescription: "Add list of locations separated by space, in that region",
						Validators: []validator.List{
							listvalidator.ValueStringsAre(stringvalidator.LengthAtMost(63)),
						},
						Optional: true,
					},
					"name": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z]+$"), "pattern must match "+"^[a-zA-Z]+$"),
						},
						MarkdownDescription: "One of the region from 'americas', 'europe', 'apac'",
						Required:            true,
					},
				},
			},
		},
		"proxy_1": schema.SingleNestedAttribute{
			MarkdownDescription: "primary regional and custom proxy",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"fqdn": schema.StringAttribute{
					Validators: []validator.String{
						stringvalidator.LengthAtMost(255),
						stringvalidator.RegexMatches(regexp.MustCompile("^([\\*a-zA-Z0-9._-])+$"), "pattern must match "+"^([\\*a-zA-Z0-9._-])+$"),
					},
					MarkdownDescription: "fqdn of the primary proxy server (supports wildcards and alphanumeric characters with dots, hyphens, and underscores)",
					Optional:            true,
				},
				"location": schema.StringAttribute{
					MarkdownDescription: "Geographic or network location identifier for the primary proxy server",
					Optional:            true,
				},
				"port": schema.Int64Attribute{
					Validators: []validator.Int64{
						int64validator.Between(1, 65535),
					},
					MarkdownDescription: "port number for primary proxy",
					Optional:            true,
				},
			},
		},
		"proxy_2": schema.SingleNestedAttribute{
			MarkdownDescription: "secondary regional and custom proxy",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"fqdn": schema.StringAttribute{
					Validators: []validator.String{
						stringvalidator.LengthAtMost(255),
						stringvalidator.RegexMatches(regexp.MustCompile("^([\\*a-zA-Z0-9._-])+$"), "pattern must match "+"^([\\*a-zA-Z0-9._-])+$"),
					},
					MarkdownDescription: "Fqdn of the secondary (backup) proxy server used for failover scenarios",
					Optional:            true,
				},
				"location": schema.StringAttribute{
					MarkdownDescription: "Geographic or network location identifier for the secondary proxy server",
					Optional:            true,
				},
				"port": schema.Int64Attribute{
					Validators: []validator.Int64{
						int64validator.Between(1, 65535),
					},
					MarkdownDescription: "port number for secondary proxy",
					Optional:            true,
				},
			},
		},
		"tfid": schema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"type": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.OneOf("gp-and-pac", "ztna-agent"),
			},
			MarkdownDescription: "Proxy configuration type - 'gp-and-pac' for GlobalProtect and PAC file forwarding, 'ztna-agent' for ZTNA agent forwarding",
			Optional:            true,
			Computed:            true,
			Default:             stringdefault.StaticString("gp-and-pac"),
		},
	},
}

// ForwardingProfileRegionalAndCustomProxiesDataSourceSchema defines the schema for ForwardingProfileRegionalAndCustomProxies data source
var ForwardingProfileRegionalAndCustomProxiesDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "ForwardingProfileRegionalAndCustomProxy data source",
	Attributes: map[string]dsschema.Attribute{
		"connectivity_preference": dsschema.ListNestedAttribute{
			MarkdownDescription: "List of connectivity methods and their enablement status for establishing proxy connections",
			Computed:            true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: map[string]dsschema.Attribute{
					"enabled": dsschema.BoolAttribute{
						MarkdownDescription: "Indicates whether this connectivity method is enabled for use in the proxy configuration",
						Computed:            true,
					},
					"name": dsschema.StringAttribute{
						MarkdownDescription: "Connectivity method type - 'tunnel' for VPN tunnels, 'proxy' for HTTP/HTTPS proxies, 'adns' for authenticated DNS, 'masque' for MASQUE protocol",
						Computed:            true,
					},
				},
			},
		},
		"description": dsschema.StringAttribute{
			MarkdownDescription: "regional and custom proxy configuration description",
			Computed:            true,
		},
		"fallback_option": dsschema.StringAttribute{
			MarkdownDescription: "Behavior when proxy connection fails - 'fail-open' allows direct internet access, 'fail-safe' blocks traffic until proxy is restored",
			Computed:            true,
		},
		"folder": dsschema.StringAttribute{
			MarkdownDescription: "The folder in which the resource is defined\n",
			Optional:            true,
			Computed:            true,
		},
		"id": dsschema.StringAttribute{
			MarkdownDescription: "The UUID of the regional and custom proxy",
			Required:            true,
		},
		"location_preference": dsschema.StringAttribute{
			MarkdownDescription: "Strategy for selecting Prisma Access location - 'best-available-pa-location' automatically selects optimal location, 'specific-pa-location' uses predefined locations",
			Computed:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "alphanumeric string [ 0-9a-zA-Z ._-]",
			Optional:            true,
			Computed:            true,
		},
		"prisma_access_locations": dsschema.ListNestedAttribute{
			MarkdownDescription: "Select Prisma Access location Americas, Europe and Asia-Pacific.",
			Computed:            true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: map[string]dsschema.Attribute{
					"locations": dsschema.ListAttribute{
						ElementType:         types.StringType,
						MarkdownDescription: "Add list of locations separated by space, in that region",
						Computed:            true,
					},
					"name": dsschema.StringAttribute{
						MarkdownDescription: "One of the region from 'americas', 'europe', 'apac'",
						Computed:            true,
					},
				},
			},
		},
		"proxy_1": dsschema.SingleNestedAttribute{
			MarkdownDescription: "primary regional and custom proxy",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"fqdn": dsschema.StringAttribute{
					MarkdownDescription: "fqdn of the primary proxy server (supports wildcards and alphanumeric characters with dots, hyphens, and underscores)",
					Computed:            true,
				},
				"location": dsschema.StringAttribute{
					MarkdownDescription: "Geographic or network location identifier for the primary proxy server",
					Computed:            true,
				},
				"port": dsschema.Int64Attribute{
					MarkdownDescription: "port number for primary proxy",
					Computed:            true,
				},
			},
		},
		"proxy_2": dsschema.SingleNestedAttribute{
			MarkdownDescription: "secondary regional and custom proxy",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"fqdn": dsschema.StringAttribute{
					MarkdownDescription: "Fqdn of the secondary (backup) proxy server used for failover scenarios",
					Computed:            true,
				},
				"location": dsschema.StringAttribute{
					MarkdownDescription: "Geographic or network location identifier for the secondary proxy server",
					Computed:            true,
				},
				"port": dsschema.Int64Attribute{
					MarkdownDescription: "port number for secondary proxy",
					Computed:            true,
				},
			},
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
		"type": dsschema.StringAttribute{
			MarkdownDescription: "Proxy configuration type - 'gp-and-pac' for GlobalProtect and PAC file forwarding, 'ztna-agent' for ZTNA agent forwarding",
			Computed:            true,
		},
	},
}

// ForwardingProfileRegionalAndCustomProxiesListModel represents the data model for a list data source.
type ForwardingProfileRegionalAndCustomProxiesListModel struct {
	Tfid    types.String                                `tfsdk:"tfid"`
	Data    []ForwardingProfileRegionalAndCustomProxies `tfsdk:"data"`
	Limit   types.Int64                                 `tfsdk:"limit"`
	Offset  types.Int64                                 `tfsdk:"offset"`
	Name    types.String                                `tfsdk:"name"`
	Total   types.Int64                                 `tfsdk:"total"`
	Folder  types.String                                `tfsdk:"folder"`
	Snippet types.String                                `tfsdk:"snippet"`
	Device  types.String                                `tfsdk:"device"`
}

// ForwardingProfileRegionalAndCustomProxiesListDataSourceSchema defines the schema for a list data source.
var ForwardingProfileRegionalAndCustomProxiesListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: ForwardingProfileRegionalAndCustomProxiesDataSourceSchema.Attributes,
			},
		},
		"limit":   dsschema.Int64Attribute{Description: "The max number of items to return. Default: 200.", Optional: true},
		"offset":  dsschema.Int64Attribute{Description: "The offset of the first item to return.", Optional: true},
		"name":    dsschema.StringAttribute{Description: "The name of the item.", Optional: true},
		"total":   dsschema.Int64Attribute{Description: "The total number of items.", Computed: true},
		"folder":  dsschema.StringAttribute{Description: "The folder of the item. Default: Shared.", Optional: true},
		"snippet": dsschema.StringAttribute{Description: "The snippet of the item.", Optional: true},
		"device":  dsschema.StringAttribute{Description: "The device of the item.", Optional: true},
	},
}
