package models

import (
	"regexp"

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

// Tags represents the Terraform model for Tags
type Tags struct {
	Tfid     types.String          `tfsdk:"tfid"`
	Color    basetypes.StringValue `tfsdk:"color"`
	Comments basetypes.StringValue `tfsdk:"comments"`
	Device   basetypes.StringValue `tfsdk:"device"`
	Folder   basetypes.StringValue `tfsdk:"folder"`
	Id       basetypes.StringValue `tfsdk:"id"`
	Name     basetypes.StringValue `tfsdk:"name"`
	Snippet  basetypes.StringValue `tfsdk:"snippet"`
}

// AttrTypes defines the attribute types for the Tags model.
func (o Tags) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":     basetypes.StringType{},
		"color":    basetypes.StringType{},
		"comments": basetypes.StringType{},
		"device":   basetypes.StringType{},
		"folder":   basetypes.StringType{},
		"id":       basetypes.StringType{},
		"name":     basetypes.StringType{},
		"snippet":  basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of Tags objects.
func (o Tags) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// TagsResourceSchema defines the schema for Tags resource
var TagsResourceSchema = schema.Schema{
	MarkdownDescription: "Tag resource",
	Attributes: map[string]schema.Attribute{
		"color": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.OneOf("Red", "Green", "Blue", "Yellow", "Copper", "Orange", "Purple", "Gray", "Light Green", "Cyan", "Light Gray", "Blue Gray", "Lime", "Black", "Gold", "Brown", "Olive", "Maroon", "Red-Orange", "Yellow-Orange", "Forest Green", "Turquoise Blue", "Azure Blue", "Cerulean Blue", "Midnight Blue", "Medium Blue", "Cobalt Blue", "Violet Blue", "Blue Violet", "Medium Violet", "Medium Rose", "Lavender", "Orchid", "Thistle", "Peach", "Salmon", "Magenta", "Red Violet", "Mahogany", "Burnt Sienna", "Chestnut"),
			},
			MarkdownDescription: "The color of the tag",
			Optional:            true,
		},
		"comments": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(1023),
			},
			MarkdownDescription: "The description of the tag",
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
		"id": schema.StringAttribute{
			MarkdownDescription: "The UUID of the tag",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"name": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(127),
			},
			MarkdownDescription: "The name of the tag",
			Required:            true,
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

// TagsDataSourceSchema defines the schema for Tags data source
var TagsDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Tag data source",
	Attributes: map[string]dsschema.Attribute{
		"color": dsschema.StringAttribute{
			MarkdownDescription: "The color of the tag",
			Computed:            true,
		},
		"comments": dsschema.StringAttribute{
			MarkdownDescription: "The description of the tag",
			Computed:            true,
		},
		"device": dsschema.StringAttribute{
			MarkdownDescription: "The device in which the resource is defined",
			Computed:            true,
		},
		"folder": dsschema.StringAttribute{
			MarkdownDescription: "The folder in which the resource is defined",
			Computed:            true,
		},
		"id": dsschema.StringAttribute{
			MarkdownDescription: "The UUID of the tag",
			Required:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "The name of the tag",
			Optional:            true,
			Computed:            true,
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

// TagsListModel represents the data model for a list data source.
type TagsListModel struct {
	Tfid    types.String `tfsdk:"tfid"`
	Data    []Tags       `tfsdk:"data"`
	Limit   types.Int64  `tfsdk:"limit"`
	Offset  types.Int64  `tfsdk:"offset"`
	Name    types.String `tfsdk:"name"`
	Total   types.Int64  `tfsdk:"total"`
	Folder  types.String `tfsdk:"folder"`
	Snippet types.String `tfsdk:"snippet"`
	Device  types.String `tfsdk:"device"`
}

// TagsListDataSourceSchema defines the schema for a list data source.
var TagsListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: TagsDataSourceSchema.Attributes,
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
