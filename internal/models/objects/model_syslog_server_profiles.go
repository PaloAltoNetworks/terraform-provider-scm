package models

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	dsschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Package: objects
// This file contains models for the objects SDK package

// SyslogServerProfiles represents the Terraform model for SyslogServerProfiles
type SyslogServerProfiles struct {
	Tfid    types.String          `tfsdk:"tfid"`
	Device  basetypes.StringValue `tfsdk:"device"`
	Folder  basetypes.StringValue `tfsdk:"folder"`
	Format  basetypes.ObjectValue `tfsdk:"format"`
	Id      basetypes.StringValue `tfsdk:"id"`
	Name    basetypes.StringValue `tfsdk:"name"`
	Servers basetypes.ObjectValue `tfsdk:"servers"`
	Snippet basetypes.StringValue `tfsdk:"snippet"`
}

// SyslogServerProfilesFormat represents a nested structure within the SyslogServerProfiles model
type SyslogServerProfilesFormat struct {
	Auth          basetypes.StringValue `tfsdk:"auth"`
	Config        basetypes.StringValue `tfsdk:"config"`
	Correlation   basetypes.StringValue `tfsdk:"correlation"`
	Data          basetypes.StringValue `tfsdk:"data"`
	Decryption    basetypes.StringValue `tfsdk:"decryption"`
	Escaping      basetypes.ObjectValue `tfsdk:"escaping"`
	Globalprotect basetypes.StringValue `tfsdk:"globalprotect"`
	Gtp           basetypes.StringValue `tfsdk:"gtp"`
	HipMatch      basetypes.StringValue `tfsdk:"hip_match"`
	Iptag         basetypes.StringValue `tfsdk:"iptag"`
	Sctp          basetypes.StringValue `tfsdk:"sctp"`
	System        basetypes.StringValue `tfsdk:"system"`
	Threat        basetypes.StringValue `tfsdk:"threat"`
	Traffic       basetypes.StringValue `tfsdk:"traffic"`
	Tunnel        basetypes.StringValue `tfsdk:"tunnel"`
	Url           basetypes.StringValue `tfsdk:"url"`
	Userid        basetypes.StringValue `tfsdk:"userid"`
	Wildfire      basetypes.StringValue `tfsdk:"wildfire"`
}

// SyslogServerProfilesFormatEscaping represents a nested structure within the SyslogServerProfiles model
type SyslogServerProfilesFormatEscaping struct {
	EscapeCharacter   basetypes.StringValue `tfsdk:"escape_character"`
	EscapedCharacters basetypes.StringValue `tfsdk:"escaped_characters"`
}

// SyslogServerProfilesServers represents a nested structure within the SyslogServerProfiles model
type SyslogServerProfilesServers struct {
	Facility  basetypes.StringValue `tfsdk:"facility"`
	Format    basetypes.StringValue `tfsdk:"format"`
	Name      basetypes.StringValue `tfsdk:"name"`
	Port      basetypes.Int64Value  `tfsdk:"port"`
	Server    basetypes.StringValue `tfsdk:"server"`
	Transport basetypes.StringValue `tfsdk:"transport"`
}

// AttrTypes defines the attribute types for the SyslogServerProfiles model.
func (o SyslogServerProfiles) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":   basetypes.StringType{},
		"device": basetypes.StringType{},
		"folder": basetypes.StringType{},
		"format": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"auth":        basetypes.StringType{},
				"config":      basetypes.StringType{},
				"correlation": basetypes.StringType{},
				"data":        basetypes.StringType{},
				"decryption":  basetypes.StringType{},
				"escaping": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"escape_character":   basetypes.StringType{},
						"escaped_characters": basetypes.StringType{},
					},
				},
				"globalprotect": basetypes.StringType{},
				"gtp":           basetypes.StringType{},
				"hip_match":     basetypes.StringType{},
				"iptag":         basetypes.StringType{},
				"sctp":          basetypes.StringType{},
				"system":        basetypes.StringType{},
				"threat":        basetypes.StringType{},
				"traffic":       basetypes.StringType{},
				"tunnel":        basetypes.StringType{},
				"url":           basetypes.StringType{},
				"userid":        basetypes.StringType{},
				"wildfire":      basetypes.StringType{},
			},
		},
		"id":   basetypes.StringType{},
		"name": basetypes.StringType{},
		"servers": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"facility":  basetypes.StringType{},
				"format":    basetypes.StringType{},
				"name":      basetypes.StringType{},
				"port":      basetypes.Int64Type{},
				"server":    basetypes.StringType{},
				"transport": basetypes.StringType{},
			},
		},
		"snippet": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of SyslogServerProfiles objects.
func (o SyslogServerProfiles) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the SyslogServerProfilesFormat model.
func (o SyslogServerProfilesFormat) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"auth":        basetypes.StringType{},
		"config":      basetypes.StringType{},
		"correlation": basetypes.StringType{},
		"data":        basetypes.StringType{},
		"decryption":  basetypes.StringType{},
		"escaping": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"escape_character":   basetypes.StringType{},
				"escaped_characters": basetypes.StringType{},
			},
		},
		"globalprotect": basetypes.StringType{},
		"gtp":           basetypes.StringType{},
		"hip_match":     basetypes.StringType{},
		"iptag":         basetypes.StringType{},
		"sctp":          basetypes.StringType{},
		"system":        basetypes.StringType{},
		"threat":        basetypes.StringType{},
		"traffic":       basetypes.StringType{},
		"tunnel":        basetypes.StringType{},
		"url":           basetypes.StringType{},
		"userid":        basetypes.StringType{},
		"wildfire":      basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of SyslogServerProfilesFormat objects.
func (o SyslogServerProfilesFormat) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the SyslogServerProfilesFormatEscaping model.
func (o SyslogServerProfilesFormatEscaping) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"escape_character":   basetypes.StringType{},
		"escaped_characters": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of SyslogServerProfilesFormatEscaping objects.
func (o SyslogServerProfilesFormatEscaping) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the SyslogServerProfilesServers model.
func (o SyslogServerProfilesServers) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"facility":  basetypes.StringType{},
		"format":    basetypes.StringType{},
		"name":      basetypes.StringType{},
		"port":      basetypes.Int64Type{},
		"server":    basetypes.StringType{},
		"transport": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of SyslogServerProfilesServers objects.
func (o SyslogServerProfilesServers) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// SyslogServerProfilesResourceSchema defines the schema for SyslogServerProfiles resource
var SyslogServerProfilesResourceSchema = schema.Schema{
	MarkdownDescription: "SyslogServerProfile resource",
	Attributes: map[string]schema.Attribute{
		"device": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.ExactlyOneOf(
					path.MatchRelative().AtParent().AtName("folder"),
					path.MatchRelative().AtParent().AtName("snippet"),
				),
				stringvalidator.LengthAtMost(64),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z\\d-_\\. ]+$"), "pattern must match "+"^[a-zA-Z\\d-_\\. ]+$"),
			},
			MarkdownDescription: "The device in which the resource is defined",
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"folder": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.ExactlyOneOf(
					path.MatchRelative().AtParent().AtName("device"),
					path.MatchRelative().AtParent().AtName("snippet"),
				),
				stringvalidator.LengthAtMost(64),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z\\d-_\\. ]+$"), "pattern must match "+"^[a-zA-Z\\d-_\\. ]+$"),
			},
			MarkdownDescription: "The folder in which the resource is defined",
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"format": schema.SingleNestedAttribute{
			MarkdownDescription: "Format",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"auth": schema.StringAttribute{
					MarkdownDescription: "Auth",
					Optional:            true,
				},
				"config": schema.StringAttribute{
					MarkdownDescription: "Config",
					Optional:            true,
				},
				"correlation": schema.StringAttribute{
					MarkdownDescription: "Correlation",
					Optional:            true,
				},
				"data": schema.StringAttribute{
					MarkdownDescription: "Data",
					Optional:            true,
				},
				"decryption": schema.StringAttribute{
					MarkdownDescription: "Decryption",
					Optional:            true,
				},
				"escaping": schema.SingleNestedAttribute{
					MarkdownDescription: "Escaping",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"escape_character": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.LengthAtMost(1),
							},
							MarkdownDescription: "Escape sequence delimiter",
							Optional:            true,
						},
						"escaped_characters": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.LengthAtMost(255),
							},
							MarkdownDescription: "A list of all the characters to be escaped (without spaces).",
							Optional:            true,
						},
					},
				},
				"globalprotect": schema.StringAttribute{
					MarkdownDescription: "Globalprotect",
					Optional:            true,
				},
				"gtp": schema.StringAttribute{
					MarkdownDescription: "Gtp",
					Optional:            true,
				},
				"hip_match": schema.StringAttribute{
					MarkdownDescription: "Hip match",
					Optional:            true,
				},
				"iptag": schema.StringAttribute{
					MarkdownDescription: "Iptag",
					Optional:            true,
				},
				"sctp": schema.StringAttribute{
					MarkdownDescription: "Sctp",
					Optional:            true,
				},
				"system": schema.StringAttribute{
					MarkdownDescription: "System",
					Optional:            true,
				},
				"threat": schema.StringAttribute{
					MarkdownDescription: "Threat",
					Optional:            true,
				},
				"traffic": schema.StringAttribute{
					MarkdownDescription: "Traffic",
					Optional:            true,
				},
				"tunnel": schema.StringAttribute{
					MarkdownDescription: "Tunnel",
					Optional:            true,
				},
				"url": schema.StringAttribute{
					MarkdownDescription: "Url",
					Optional:            true,
				},
				"userid": schema.StringAttribute{
					MarkdownDescription: "Userid",
					Optional:            true,
				},
				"wildfire": schema.StringAttribute{
					MarkdownDescription: "Wildfire",
					Optional:            true,
				},
			},
		},
		"id": schema.StringAttribute{
			MarkdownDescription: "The UUID of the syslog server profile",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"name": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(31),
			},
			MarkdownDescription: "The name of the syslog server profile",
			Required:            true,
		},
		"servers": schema.SingleNestedAttribute{
			MarkdownDescription: "Servers",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"facility": schema.StringAttribute{
					Validators: []validator.String{
						stringvalidator.OneOf("LOG_USER", "LOG_LOCAL0", "LOG_LOCAL1", "LOG_LOCAL2", "LOG_LOCAL3", "LOG_LOCAL4", "LOG_LOCAL5", "LOG_LOCAL6", "LOG_LOCAL7"),
					},
					MarkdownDescription: "Syslog facility",
					Optional:            true,
				},
				"format": schema.StringAttribute{
					Validators: []validator.String{
						stringvalidator.OneOf("BSD", "IETF"),
					},
					MarkdownDescription: "Syslog format",
					Optional:            true,
				},
				"name": schema.StringAttribute{
					MarkdownDescription: "Syslog server name",
					Optional:            true,
				},
				"port": schema.Int64Attribute{
					Validators: []validator.Int64{
						int64validator.Between(1, 65535),
					},
					MarkdownDescription: "Syslog server port",
					Optional:            true,
				},
				"server": schema.StringAttribute{
					MarkdownDescription: "Syslog server address",
					Optional:            true,
				},
				"transport": schema.StringAttribute{
					Validators: []validator.String{
						stringvalidator.OneOf("UDP", "TCP"),
					},
					MarkdownDescription: "Transport protocol",
					Optional:            true,
				},
			},
		},
		"snippet": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.ExactlyOneOf(
					path.MatchRelative().AtParent().AtName("device"),
					path.MatchRelative().AtParent().AtName("folder"),
				),
				stringvalidator.LengthAtMost(64),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z\\d-_\\. ]+$"), "pattern must match "+"^[a-zA-Z\\d-_\\. ]+$"),
			},
			MarkdownDescription: "The snippet in which the resource is defined",
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"tfid": schema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
	},
}

// SyslogServerProfilesDataSourceSchema defines the schema for SyslogServerProfiles data source
var SyslogServerProfilesDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "SyslogServerProfile data source",
	Attributes: map[string]dsschema.Attribute{
		"device": dsschema.StringAttribute{
			MarkdownDescription: "The device in which the resource is defined",
			Computed:            true,
		},
		"folder": dsschema.StringAttribute{
			MarkdownDescription: "The folder in which the resource is defined",
			Computed:            true,
		},
		"format": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Format",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"auth": dsschema.StringAttribute{
					MarkdownDescription: "Auth",
					Computed:            true,
				},
				"config": dsschema.StringAttribute{
					MarkdownDescription: "Config",
					Computed:            true,
				},
				"correlation": dsschema.StringAttribute{
					MarkdownDescription: "Correlation",
					Computed:            true,
				},
				"data": dsschema.StringAttribute{
					MarkdownDescription: "Data",
					Computed:            true,
				},
				"decryption": dsschema.StringAttribute{
					MarkdownDescription: "Decryption",
					Computed:            true,
				},
				"escaping": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Escaping",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"escape_character": dsschema.StringAttribute{
							MarkdownDescription: "Escape sequence delimiter",
							Computed:            true,
						},
						"escaped_characters": dsschema.StringAttribute{
							MarkdownDescription: "A list of all the characters to be escaped (without spaces).",
							Computed:            true,
						},
					},
				},
				"globalprotect": dsschema.StringAttribute{
					MarkdownDescription: "Globalprotect",
					Computed:            true,
				},
				"gtp": dsschema.StringAttribute{
					MarkdownDescription: "Gtp",
					Computed:            true,
				},
				"hip_match": dsschema.StringAttribute{
					MarkdownDescription: "Hip match",
					Computed:            true,
				},
				"iptag": dsschema.StringAttribute{
					MarkdownDescription: "Iptag",
					Computed:            true,
				},
				"sctp": dsschema.StringAttribute{
					MarkdownDescription: "Sctp",
					Computed:            true,
				},
				"system": dsschema.StringAttribute{
					MarkdownDescription: "System",
					Computed:            true,
				},
				"threat": dsschema.StringAttribute{
					MarkdownDescription: "Threat",
					Computed:            true,
				},
				"traffic": dsschema.StringAttribute{
					MarkdownDescription: "Traffic",
					Computed:            true,
				},
				"tunnel": dsschema.StringAttribute{
					MarkdownDescription: "Tunnel",
					Computed:            true,
				},
				"url": dsschema.StringAttribute{
					MarkdownDescription: "Url",
					Computed:            true,
				},
				"userid": dsschema.StringAttribute{
					MarkdownDescription: "Userid",
					Computed:            true,
				},
				"wildfire": dsschema.StringAttribute{
					MarkdownDescription: "Wildfire",
					Computed:            true,
				},
			},
		},
		"id": dsschema.StringAttribute{
			MarkdownDescription: "The UUID of the syslog server profile",
			Required:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "The name of the syslog server profile",
			Optional:            true,
			Computed:            true,
		},
		"servers": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Servers",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"facility": dsschema.StringAttribute{
					MarkdownDescription: "Syslog facility",
					Computed:            true,
				},
				"format": dsschema.StringAttribute{
					MarkdownDescription: "Syslog format",
					Computed:            true,
				},
				"name": dsschema.StringAttribute{
					MarkdownDescription: "Syslog server name",
					Computed:            true,
				},
				"port": dsschema.Int64Attribute{
					MarkdownDescription: "Syslog server port",
					Computed:            true,
				},
				"server": dsschema.StringAttribute{
					MarkdownDescription: "Syslog server address",
					Computed:            true,
				},
				"transport": dsschema.StringAttribute{
					MarkdownDescription: "Transport protocol",
					Computed:            true,
				},
			},
		},
		"snippet": dsschema.StringAttribute{
			MarkdownDescription: "The snippet in which the resource is defined",
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
	},
}

// SyslogServerProfilesListModel represents the data model for a list data source.
type SyslogServerProfilesListModel struct {
	Tfid    types.String           `tfsdk:"tfid"`
	Data    []SyslogServerProfiles `tfsdk:"data"`
	Limit   types.Int64            `tfsdk:"limit"`
	Offset  types.Int64            `tfsdk:"offset"`
	Name    types.String           `tfsdk:"name"`
	Total   types.Int64            `tfsdk:"total"`
	Folder  types.String           `tfsdk:"folder"`
	Snippet types.String           `tfsdk:"snippet"`
	Device  types.String           `tfsdk:"device"`
}

// SyslogServerProfilesListDataSourceSchema defines the schema for a list data source.
var SyslogServerProfilesListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: SyslogServerProfilesDataSourceSchema.Attributes,
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
