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

// Package: identity_services
// This file contains models for the identity_services SDK package

// LdapServerProfiles represents the Terraform model for LdapServerProfiles
type LdapServerProfiles struct {
	Tfid                    types.String          `tfsdk:"tfid"`
	EncryptedValues         basetypes.MapValue    `tfsdk:"encrypted_values"`
	Base                    basetypes.StringValue `tfsdk:"base"`
	BindDn                  basetypes.StringValue `tfsdk:"bind_dn"`
	BindPassword            basetypes.StringValue `tfsdk:"bind_password"`
	BindTimelimit           basetypes.StringValue `tfsdk:"bind_timelimit"`
	Device                  basetypes.StringValue `tfsdk:"device"`
	Folder                  basetypes.StringValue `tfsdk:"folder"`
	Id                      basetypes.StringValue `tfsdk:"id"`
	LdapType                basetypes.StringValue `tfsdk:"ldap_type"`
	Name                    basetypes.StringValue `tfsdk:"name"`
	RetryInterval           basetypes.Int64Value  `tfsdk:"retry_interval"`
	Server                  basetypes.ListValue   `tfsdk:"server"`
	Snippet                 basetypes.StringValue `tfsdk:"snippet"`
	Ssl                     basetypes.BoolValue   `tfsdk:"ssl"`
	Timelimit               basetypes.Int64Value  `tfsdk:"timelimit"`
	VerifyServerCertificate basetypes.BoolValue   `tfsdk:"verify_server_certificate"`
}

// LdapServerProfilesServerInner represents a nested structure within the LdapServerProfiles model
type LdapServerProfilesServerInner struct {
	Address basetypes.StringValue `tfsdk:"address"`
	Name    basetypes.StringValue `tfsdk:"name"`
	Port    basetypes.Int64Value  `tfsdk:"port"`
}

// AttrTypes defines the attribute types for the LdapServerProfiles model.
func (o LdapServerProfiles) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":             basetypes.StringType{},
		"encrypted_values": basetypes.MapType{ElemType: basetypes.StringType{}},
		"base":             basetypes.StringType{},
		"bind_dn":          basetypes.StringType{},
		"bind_password":    basetypes.StringType{},
		"bind_timelimit":   basetypes.StringType{},
		"device":           basetypes.StringType{},
		"folder":           basetypes.StringType{},
		"id":               basetypes.StringType{},
		"ldap_type":        basetypes.StringType{},
		"name":             basetypes.StringType{},
		"retry_interval":   basetypes.Int64Type{},
		"server": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"address": basetypes.StringType{},
				"name":    basetypes.StringType{},
				"port":    basetypes.Int64Type{},
			},
		}},
		"snippet":                   basetypes.StringType{},
		"ssl":                       basetypes.BoolType{},
		"timelimit":                 basetypes.Int64Type{},
		"verify_server_certificate": basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of LdapServerProfiles objects.
func (o LdapServerProfiles) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LdapServerProfilesServerInner model.
func (o LdapServerProfilesServerInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"address": basetypes.StringType{},
		"name":    basetypes.StringType{},
		"port":    basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of LdapServerProfilesServerInner objects.
func (o LdapServerProfilesServerInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// LdapServerProfilesResourceSchema defines the schema for LdapServerProfiles resource
var LdapServerProfilesResourceSchema = schema.Schema{
	MarkdownDescription: "Manages SCM LdapServerProfiles objects",
	Attributes: map[string]schema.Attribute{
		"base": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(255),
			},
			MarkdownDescription: "The base DN",
			Optional:            true,
		},
		"bind_dn": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(255),
			},
			MarkdownDescription: "The bind DN",
			Optional:            true,
		},
		"bind_password": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(121),
			},
			MarkdownDescription: "The bind password",
			Optional:            true,
			Sensitive:           true,
		},
		"bind_timelimit": schema.StringAttribute{
			MarkdownDescription: "The bind timeout (seconds)",
			Optional:            true,
		},
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
		"encrypted_values": schema.MapAttribute{
			ElementType:         basetypes.StringType{},
			MarkdownDescription: "Map of sensitive values returned from the API.",
			Computed:            true,
			Sensitive:           true,
		},
		"folder": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.ExactlyOneOf(
					path.MatchRelative().AtParent().AtName("snippet"),
					path.MatchRelative().AtParent().AtName("device"),
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
		"id": schema.StringAttribute{
			MarkdownDescription: "The UUID of the LDAP server profile",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"ldap_type": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.OneOf("active-directory", "e-directory", "sun", "other"),
			},
			MarkdownDescription: "The LDAP server time",
			Optional:            true,
		},
		"name": schema.StringAttribute{
			MarkdownDescription: "The name of the LDAP server profile",
			Required:            true,
		},
		"retry_interval": schema.Int64Attribute{
			MarkdownDescription: "The search retry interval (seconds)",
			Optional:            true,
		},
		"server": schema.ListNestedAttribute{
			MarkdownDescription: "The LDAP server configuration",
			Required:            true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"address": schema.StringAttribute{
						MarkdownDescription: "The LDAP server IP address",
						Optional:            true,
					},
					"name": schema.StringAttribute{
						MarkdownDescription: "The LDAP server name",
						Optional:            true,
					},
					"port": schema.Int64Attribute{
						Validators: []validator.Int64{
							int64validator.Between(1, 65535),
						},
						MarkdownDescription: "The LDAP server port",
						Optional:            true,
					},
				},
			},
		},
		"snippet": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.ExactlyOneOf(
					path.MatchRelative().AtParent().AtName("folder"),
					path.MatchRelative().AtParent().AtName("device"),
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
		"ssl": schema.BoolAttribute{
			MarkdownDescription: "Require SSL/TLS secured connection?",
			Optional:            true,
		},
		"tfid": schema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"timelimit": schema.Int64Attribute{
			MarkdownDescription: "The search timeout (seconds)",
			Optional:            true,
		},
		"verify_server_certificate": schema.BoolAttribute{
			MarkdownDescription: "Verify server certificate for SSL sessions?",
			Optional:            true,
		},
	},
}

// LdapServerProfilesDataSourceSchema defines the schema for LdapServerProfiles data source
var LdapServerProfilesDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "LdapServerProfiles data source",
	Attributes: map[string]dsschema.Attribute{
		"base": dsschema.StringAttribute{
			MarkdownDescription: "The base DN",
			Computed:            true,
		},
		"bind_dn": dsschema.StringAttribute{
			MarkdownDescription: "The bind DN",
			Computed:            true,
		},
		"bind_password": dsschema.StringAttribute{
			MarkdownDescription: "The bind password",
			Computed:            true,
			Sensitive:           true,
		},
		"bind_timelimit": dsschema.StringAttribute{
			MarkdownDescription: "The bind timeout (seconds)",
			Computed:            true,
		},
		"device": dsschema.StringAttribute{
			MarkdownDescription: "The device in which the resource is defined",
			Computed:            true,
		},
		"encrypted_values": dsschema.MapAttribute{
			ElementType:         basetypes.StringType{},
			MarkdownDescription: "Map of sensitive values returned from the API.",
			Computed:            true,
			Sensitive:           true,
		},
		"folder": dsschema.StringAttribute{
			MarkdownDescription: "The folder in which the resource is defined",
			Computed:            true,
		},
		"id": dsschema.StringAttribute{
			MarkdownDescription: "The UUID of the LDAP server profile",
			Required:            true,
		},
		"ldap_type": dsschema.StringAttribute{
			MarkdownDescription: "The LDAP server time",
			Computed:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "The name of the LDAP server profile",
			Optional:            true,
			Computed:            true,
		},
		"retry_interval": dsschema.Int64Attribute{
			MarkdownDescription: "The search retry interval (seconds)",
			Computed:            true,
		},
		"server": dsschema.ListNestedAttribute{
			MarkdownDescription: "The LDAP server configuration",
			Computed:            true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: map[string]dsschema.Attribute{
					"address": dsschema.StringAttribute{
						MarkdownDescription: "The LDAP server IP address",
						Computed:            true,
					},
					"name": dsschema.StringAttribute{
						MarkdownDescription: "The LDAP server name",
						Computed:            true,
					},
					"port": dsschema.Int64Attribute{
						MarkdownDescription: "The LDAP server port",
						Computed:            true,
					},
				},
			},
		},
		"snippet": dsschema.StringAttribute{
			MarkdownDescription: "The snippet in which the resource is defined",
			Computed:            true,
		},
		"ssl": dsschema.BoolAttribute{
			MarkdownDescription: "Require SSL/TLS secured connection?",
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
		"timelimit": dsschema.Int64Attribute{
			MarkdownDescription: "The search timeout (seconds)",
			Computed:            true,
		},
		"verify_server_certificate": dsschema.BoolAttribute{
			MarkdownDescription: "Verify server certificate for SSL sessions?",
			Computed:            true,
		},
	},
}

// LdapServerProfilesListModel represents the data model for a list data source.
type LdapServerProfilesListModel struct {
	Tfid    types.String         `tfsdk:"tfid"`
	Data    []LdapServerProfiles `tfsdk:"data"`
	Limit   types.Int64          `tfsdk:"limit"`
	Offset  types.Int64          `tfsdk:"offset"`
	Name    types.String         `tfsdk:"name"`
	Total   types.Int64          `tfsdk:"total"`
	Folder  types.String         `tfsdk:"folder"`
	Snippet types.String         `tfsdk:"snippet"`
	Device  types.String         `tfsdk:"device"`
}

// LdapServerProfilesListDataSourceSchema defines the schema for a list data source.
var LdapServerProfilesListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: LdapServerProfilesDataSourceSchema.Attributes,
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
