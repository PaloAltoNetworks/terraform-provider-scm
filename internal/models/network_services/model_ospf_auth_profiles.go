package models

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
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

// Package: network_services
// This file contains models for the network_services SDK package

// OspfAuthProfiles represents the Terraform model for OspfAuthProfiles
type OspfAuthProfiles struct {
	Tfid            types.String          `tfsdk:"tfid"`
	EncryptedValues basetypes.MapValue    `tfsdk:"encrypted_values"`
	Device          basetypes.StringValue `tfsdk:"device"`
	Folder          basetypes.StringValue `tfsdk:"folder"`
	Id              basetypes.StringValue `tfsdk:"id"`
	Md5             basetypes.ListValue   `tfsdk:"md5"`
	Name            basetypes.StringValue `tfsdk:"name"`
	Password        basetypes.StringValue `tfsdk:"password"`
	Snippet         basetypes.StringValue `tfsdk:"snippet"`
}

// OspfAuthProfilesMd5Inner represents a nested structure within the OspfAuthProfiles model
type OspfAuthProfilesMd5Inner struct {
	Key       basetypes.StringValue `tfsdk:"key"`
	Name      basetypes.Int64Value  `tfsdk:"name"`
	Preferred basetypes.BoolValue   `tfsdk:"preferred"`
}

// AttrTypes defines the attribute types for the OspfAuthProfiles model.
func (o OspfAuthProfiles) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":             basetypes.StringType{},
		"encrypted_values": basetypes.MapType{ElemType: basetypes.StringType{}},
		"device":           basetypes.StringType{},
		"folder":           basetypes.StringType{},
		"id":               basetypes.StringType{},
		"md5": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"key":       basetypes.StringType{},
				"name":      basetypes.Int64Type{},
				"preferred": basetypes.BoolType{},
			},
		}},
		"name":     basetypes.StringType{},
		"password": basetypes.StringType{},
		"snippet":  basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of OspfAuthProfiles objects.
func (o OspfAuthProfiles) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the OspfAuthProfilesMd5Inner model.
func (o OspfAuthProfilesMd5Inner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"key":       basetypes.StringType{},
		"name":      basetypes.Int64Type{},
		"preferred": basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of OspfAuthProfilesMd5Inner objects.
func (o OspfAuthProfilesMd5Inner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// OspfAuthProfilesResourceSchema defines the schema for OspfAuthProfiles resource
var OspfAuthProfilesResourceSchema = schema.Schema{
	MarkdownDescription: "OspfAuthProfile resource",
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
					path.MatchRelative().AtParent().AtName("device"),
					path.MatchRelative().AtParent().AtName("snippet"),
				),
				stringvalidator.LengthAtMost(64),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z\\d\\-_\\. ]+$"), "pattern must match "+"^[a-zA-Z\\d\\-_\\. ]+$"),
			},
			MarkdownDescription: "The folder in which the resource is defined",
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"id": schema.StringAttribute{
			MarkdownDescription: "UUID of the resource",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"md5": schema.ListNestedAttribute{
			Validators: []validator.List{
				listvalidator.ExactlyOneOf(
					path.MatchRelative().AtParent().AtName("password"),
				),
			},
			MarkdownDescription: "MD5s",
			Optional:            true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"key": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.LengthAtMost(16),
						},
						MarkdownDescription: "MD5 hash",
						Optional:            true,
						Sensitive:           true,
					},
					"name": schema.Int64Attribute{
						Validators: []validator.Int64{
							int64validator.Between(1, 255),
						},
						MarkdownDescription: "Key ID",
						Optional:            true,
					},
					"preferred": schema.BoolAttribute{
						MarkdownDescription: "Preferred?",
						Optional:            true,
					},
				},
			},
		},
		"name": schema.StringAttribute{
			MarkdownDescription: "Profile name",
			Required:            true,
		},
		"password": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.ExactlyOneOf(
					path.MatchRelative().AtParent().AtName("md5"),
				),
			},
			MarkdownDescription: "Password",
			Optional:            true,
			Sensitive:           true,
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

// OspfAuthProfilesDataSourceSchema defines the schema for OspfAuthProfiles data source
var OspfAuthProfilesDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "OspfAuthProfile data source",
	Attributes: map[string]dsschema.Attribute{
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
			MarkdownDescription: "UUID of the resource",
			Required:            true,
		},
		"md5": dsschema.ListNestedAttribute{
			MarkdownDescription: "MD5s",
			Computed:            true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: map[string]dsschema.Attribute{
					"key": dsschema.StringAttribute{
						MarkdownDescription: "MD5 hash",
						Computed:            true,
						Sensitive:           true,
					},
					"name": dsschema.Int64Attribute{
						MarkdownDescription: "Key ID",
						Computed:            true,
					},
					"preferred": dsschema.BoolAttribute{
						MarkdownDescription: "Preferred?",
						Computed:            true,
					},
				},
			},
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "Profile name",
			Optional:            true,
			Computed:            true,
		},
		"password": dsschema.StringAttribute{
			MarkdownDescription: "Password",
			Computed:            true,
			Sensitive:           true,
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

// OspfAuthProfilesListModel represents the data model for a list data source.
type OspfAuthProfilesListModel struct {
	Tfid    types.String       `tfsdk:"tfid"`
	Data    []OspfAuthProfiles `tfsdk:"data"`
	Limit   types.Int64        `tfsdk:"limit"`
	Offset  types.Int64        `tfsdk:"offset"`
	Name    types.String       `tfsdk:"name"`
	Total   types.Int64        `tfsdk:"total"`
	Folder  types.String       `tfsdk:"folder"`
	Snippet types.String       `tfsdk:"snippet"`
	Device  types.String       `tfsdk:"device"`
}

// OspfAuthProfilesListDataSourceSchema defines the schema for a list data source.
var OspfAuthProfilesListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: OspfAuthProfilesDataSourceSchema.Attributes,
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
