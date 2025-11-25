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

// KerberosServerProfiles represents the Terraform model for KerberosServerProfiles
type KerberosServerProfiles struct {
	Tfid    types.String          `tfsdk:"tfid"`
	Device  basetypes.StringValue `tfsdk:"device"`
	Folder  basetypes.StringValue `tfsdk:"folder"`
	Id      basetypes.StringValue `tfsdk:"id"`
	Name    basetypes.StringValue `tfsdk:"name"`
	Server  basetypes.ListValue   `tfsdk:"server"`
	Snippet basetypes.StringValue `tfsdk:"snippet"`
}

// KerberosServerProfilesServerInner represents a nested structure within the KerberosServerProfiles model
type KerberosServerProfilesServerInner struct {
	Host basetypes.StringValue `tfsdk:"host"`
	Name basetypes.StringValue `tfsdk:"name"`
	Port basetypes.Int64Value  `tfsdk:"port"`
}

// AttrTypes defines the attribute types for the KerberosServerProfiles model.
func (o KerberosServerProfiles) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":   basetypes.StringType{},
		"device": basetypes.StringType{},
		"folder": basetypes.StringType{},
		"id":     basetypes.StringType{},
		"name":   basetypes.StringType{},
		"server": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"host": basetypes.StringType{},
				"name": basetypes.StringType{},
				"port": basetypes.Int64Type{},
			},
		}},
		"snippet": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of KerberosServerProfiles objects.
func (o KerberosServerProfiles) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the KerberosServerProfilesServerInner model.
func (o KerberosServerProfilesServerInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"host": basetypes.StringType{},
		"name": basetypes.StringType{},
		"port": basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of KerberosServerProfilesServerInner objects.
func (o KerberosServerProfilesServerInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// KerberosServerProfilesResourceSchema defines the schema for KerberosServerProfiles resource
var KerberosServerProfilesResourceSchema = schema.Schema{
	MarkdownDescription: "KerberosServerProfile resource",
	Attributes: map[string]schema.Attribute{
		"device": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.ExactlyOneOf(
					path.MatchRelative().AtParent().AtName("folder"),
					path.MatchRelative().AtParent().AtName("snippet"),
				),
				stringvalidator.LengthAtMost(64),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z\\d\\-_\\. ]+$"), "pattern must match "+"^[a-zA-Z\\d\\-_\\. ]+$"),
			},
			MarkdownDescription: "The device in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
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
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z\\d\\-_\\. ]+$"), "pattern must match "+"^[a-zA-Z\\d\\-_\\. ]+$"),
			},
			MarkdownDescription: "The folder in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"id": schema.StringAttribute{
			MarkdownDescription: "The UUID of the Kerberos server profile",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"name": schema.StringAttribute{
			MarkdownDescription: "The name of the Kerberos server profile",
			Required:            true,
		},
		"server": schema.ListNestedAttribute{
			MarkdownDescription: "The Kerberos server configuration",
			Required:            true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"host": schema.StringAttribute{
						MarkdownDescription: "The Kerberos server IP address",
						Optional:            true,
					},
					"name": schema.StringAttribute{
						MarkdownDescription: "The Kerberos server name",
						Optional:            true,
					},
					"port": schema.Int64Attribute{
						Validators: []validator.Int64{
							int64validator.Between(1, 65535),
						},
						MarkdownDescription: "The Kerberos server port",
						Optional:            true,
					},
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
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z\\d\\-_\\. ]+$"), "pattern must match "+"^[a-zA-Z\\d\\-_\\. ]+$"),
			},
			MarkdownDescription: "The snippet in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
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

// KerberosServerProfilesDataSourceSchema defines the schema for KerberosServerProfiles data source
var KerberosServerProfilesDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "KerberosServerProfile data source",
	Attributes: map[string]dsschema.Attribute{
		"device": dsschema.StringAttribute{
			MarkdownDescription: "The device in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Computed:            true,
		},
		"folder": dsschema.StringAttribute{
			MarkdownDescription: "The folder in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Computed:            true,
		},
		"id": dsschema.StringAttribute{
			MarkdownDescription: "The UUID of the Kerberos server profile",
			Required:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "The name of the Kerberos server profile",
			Optional:            true,
			Computed:            true,
		},
		"server": dsschema.ListNestedAttribute{
			MarkdownDescription: "The Kerberos server configuration",
			Computed:            true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: map[string]dsschema.Attribute{
					"host": dsschema.StringAttribute{
						MarkdownDescription: "The Kerberos server IP address",
						Computed:            true,
					},
					"name": dsschema.StringAttribute{
						MarkdownDescription: "The Kerberos server name",
						Computed:            true,
					},
					"port": dsschema.Int64Attribute{
						MarkdownDescription: "The Kerberos server port",
						Computed:            true,
					},
				},
			},
		},
		"snippet": dsschema.StringAttribute{
			MarkdownDescription: "The snippet in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
	},
}

// KerberosServerProfilesListModel represents the data model for a list data source.
type KerberosServerProfilesListModel struct {
	Tfid    types.String             `tfsdk:"tfid"`
	Data    []KerberosServerProfiles `tfsdk:"data"`
	Limit   types.Int64              `tfsdk:"limit"`
	Offset  types.Int64              `tfsdk:"offset"`
	Name    types.String             `tfsdk:"name"`
	Total   types.Int64              `tfsdk:"total"`
	Folder  types.String             `tfsdk:"folder"`
	Snippet types.String             `tfsdk:"snippet"`
	Device  types.String             `tfsdk:"device"`
}

// KerberosServerProfilesListDataSourceSchema defines the schema for a list data source.
var KerberosServerProfilesListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: KerberosServerProfilesDataSourceSchema.Attributes,
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
