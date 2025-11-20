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

// TacacsServerProfiles represents the Terraform model for TacacsServerProfiles
type TacacsServerProfiles struct {
	Tfid                types.String          `tfsdk:"tfid"`
	Device              basetypes.StringValue `tfsdk:"device"`
	Folder              basetypes.StringValue `tfsdk:"folder"`
	Id                  basetypes.StringValue `tfsdk:"id"`
	Name                basetypes.StringValue `tfsdk:"name"`
	Protocol            basetypes.StringValue `tfsdk:"protocol"`
	Server              basetypes.ListValue   `tfsdk:"server"`
	Snippet             basetypes.StringValue `tfsdk:"snippet"`
	Timeout             basetypes.Int64Value  `tfsdk:"timeout"`
	UseSingleConnection basetypes.BoolValue   `tfsdk:"use_single_connection"`
}

// TacacsServerProfilesServerInner represents a nested structure within the TacacsServerProfiles model
type TacacsServerProfilesServerInner struct {
	Address basetypes.StringValue `tfsdk:"address"`
	Name    basetypes.StringValue `tfsdk:"name"`
	Port    basetypes.Int64Value  `tfsdk:"port"`
	Secret  basetypes.StringValue `tfsdk:"secret"`
}

// AttrTypes defines the attribute types for the TacacsServerProfiles model.
func (o TacacsServerProfiles) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":     basetypes.StringType{},
		"device":   basetypes.StringType{},
		"folder":   basetypes.StringType{},
		"id":       basetypes.StringType{},
		"name":     basetypes.StringType{},
		"protocol": basetypes.StringType{},
		"server": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"address": basetypes.StringType{},
				"name":    basetypes.StringType{},
				"port":    basetypes.Int64Type{},
				"secret":  basetypes.StringType{},
			},
		}},
		"snippet":               basetypes.StringType{},
		"timeout":               basetypes.Int64Type{},
		"use_single_connection": basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of TacacsServerProfiles objects.
func (o TacacsServerProfiles) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the TacacsServerProfilesServerInner model.
func (o TacacsServerProfilesServerInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"address": basetypes.StringType{},
		"name":    basetypes.StringType{},
		"port":    basetypes.Int64Type{},
		"secret":  basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of TacacsServerProfilesServerInner objects.
func (o TacacsServerProfilesServerInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// TacacsServerProfilesResourceSchema defines the schema for TacacsServerProfiles resource
var TacacsServerProfilesResourceSchema = schema.Schema{
	MarkdownDescription: "TacacsServerProfile resource",
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
			MarkdownDescription: "The UUID of the TACACS+ server profile",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"name": schema.StringAttribute{
			MarkdownDescription: "The name of the TACACS+ server profile",
			Required:            true,
		},
		"protocol": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.OneOf("CHAP", "PAP"),
			},
			MarkdownDescription: "The TACACS+ authentication protocol",
			Required:            true,
		},
		"server": schema.ListNestedAttribute{
			MarkdownDescription: "The TACACS+ server configuration",
			Required:            true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"address": schema.StringAttribute{
						MarkdownDescription: "The IP address of the TACACS+ server",
						Optional:            true,
					},
					"name": schema.StringAttribute{
						MarkdownDescription: "The name of the TACACS+ server",
						Optional:            true,
					},
					"port": schema.Int64Attribute{
						Validators: []validator.Int64{
							int64validator.Between(1, 65535),
						},
						MarkdownDescription: "The TACACS+ server port",
						Optional:            true,
					},
					"secret": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.LengthAtMost(64),
						},
						MarkdownDescription: "The TACACS+ secret",
						Optional:            true,
						Sensitive:           true,
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
		"timeout": schema.Int64Attribute{
			Validators: []validator.Int64{
				int64validator.Between(1, 30),
			},
			MarkdownDescription: "The TACACS+ timeout (seconds)",
			Optional:            true,
		},
		"use_single_connection": schema.BoolAttribute{
			MarkdownDescription: "Use a single TACACS+ connection?",
			Optional:            true,
		},
	},
}

// TacacsServerProfilesDataSourceSchema defines the schema for TacacsServerProfiles data source
var TacacsServerProfilesDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "TacacsServerProfile data source",
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
			MarkdownDescription: "The UUID of the TACACS+ server profile",
			Required:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "The name of the TACACS+ server profile",
			Optional:            true,
			Computed:            true,
		},
		"protocol": dsschema.StringAttribute{
			MarkdownDescription: "The TACACS+ authentication protocol",
			Computed:            true,
		},
		"server": dsschema.ListNestedAttribute{
			MarkdownDescription: "The TACACS+ server configuration",
			Computed:            true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: map[string]dsschema.Attribute{
					"address": dsschema.StringAttribute{
						MarkdownDescription: "The IP address of the TACACS+ server",
						Computed:            true,
					},
					"name": dsschema.StringAttribute{
						MarkdownDescription: "The name of the TACACS+ server",
						Computed:            true,
					},
					"port": dsschema.Int64Attribute{
						MarkdownDescription: "The TACACS+ server port",
						Computed:            true,
					},
					"secret": dsschema.StringAttribute{
						MarkdownDescription: "The TACACS+ secret",
						Computed:            true,
						Sensitive:           true,
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
		"timeout": dsschema.Int64Attribute{
			MarkdownDescription: "The TACACS+ timeout (seconds)",
			Computed:            true,
		},
		"use_single_connection": dsschema.BoolAttribute{
			MarkdownDescription: "Use a single TACACS+ connection?",
			Computed:            true,
		},
	},
}

// TacacsServerProfilesListModel represents the data model for a list data source.
type TacacsServerProfilesListModel struct {
	Tfid    types.String           `tfsdk:"tfid"`
	Data    []TacacsServerProfiles `tfsdk:"data"`
	Limit   types.Int64            `tfsdk:"limit"`
	Offset  types.Int64            `tfsdk:"offset"`
	Name    types.String           `tfsdk:"name"`
	Total   types.Int64            `tfsdk:"total"`
	Folder  types.String           `tfsdk:"folder"`
	Snippet types.String           `tfsdk:"snippet"`
	Device  types.String           `tfsdk:"device"`
}

// TacacsServerProfilesListDataSourceSchema defines the schema for a list data source.
var TacacsServerProfilesListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: TacacsServerProfilesDataSourceSchema.Attributes,
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
