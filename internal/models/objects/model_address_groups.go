package models

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/objectvalidator"
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

// AddressGroups represents the Terraform model for AddressGroups
type AddressGroups struct {
	Tfid        types.String          `tfsdk:"tfid"`
	Description basetypes.StringValue `tfsdk:"description"`
	Device      basetypes.StringValue `tfsdk:"device"`
	Dynamic     basetypes.ObjectValue `tfsdk:"dynamic"`
	Folder      basetypes.StringValue `tfsdk:"folder"`
	Id          basetypes.StringValue `tfsdk:"id"`
	Name        basetypes.StringValue `tfsdk:"name"`
	Snippet     basetypes.StringValue `tfsdk:"snippet"`
	Static      basetypes.ListValue   `tfsdk:"static"`
	Tag         basetypes.ListValue   `tfsdk:"tag"`
}

// AddressGroupsDynamic represents a nested structure within the AddressGroups model
type AddressGroupsDynamic struct {
	Filter basetypes.StringValue `tfsdk:"filter"`
}

// AttrTypes defines the attribute types for the AddressGroups model.
func (o AddressGroups) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":        basetypes.StringType{},
		"description": basetypes.StringType{},
		"device":      basetypes.StringType{},
		"dynamic": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"filter": basetypes.StringType{},
			},
		},
		"folder":  basetypes.StringType{},
		"id":      basetypes.StringType{},
		"name":    basetypes.StringType{},
		"snippet": basetypes.StringType{},
		"static":  basetypes.ListType{ElemType: basetypes.StringType{}},
		"tag":     basetypes.ListType{ElemType: basetypes.StringType{}},
	}
}

// AttrType returns the attribute type for a list of AddressGroups objects.
func (o AddressGroups) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the AddressGroupsDynamic model.
func (o AddressGroupsDynamic) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"filter": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of AddressGroupsDynamic objects.
func (o AddressGroupsDynamic) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AddressGroupsResourceSchema defines the schema for AddressGroups resource
var AddressGroupsResourceSchema = schema.Schema{
	MarkdownDescription: "AddressGroup resource",
	Attributes: map[string]schema.Attribute{
		"description": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(1023),
			},
			MarkdownDescription: "Description",
			Optional:            true,
		},
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
		"dynamic": schema.SingleNestedAttribute{
			Validators: []validator.Object{
				objectvalidator.ExactlyOneOf(
					path.MatchRelative().AtParent().AtName("static"),
				),
			},
			MarkdownDescription: "Dynamic\n\n> ℹ️ **Note:** You must specify exactly one of `dynamic` and `static`.",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"filter": schema.StringAttribute{
					Validators: []validator.String{
						stringvalidator.LengthAtMost(2047),
					},
					MarkdownDescription: "Tag based filter defining group membership",
					Required:            true,
				},
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
			MarkdownDescription: "The UUID of the address group",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"name": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(63),
				stringvalidator.RegexMatches(regexp.MustCompile("^[ a-zA-Z\\d._-]+$"), "pattern must match "+"^[ a-zA-Z\\d._-]+$"),
			},
			MarkdownDescription: "The name of the address group",
			Required:            true,
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
		"static": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Static\n\n> ℹ️ **Note:** You must specify exactly one of `dynamic` and `static`.",
			Validators: []validator.List{
				listvalidator.ExactlyOneOf(
					path.MatchRelative().AtParent().AtName("dynamic"),
				),
				listvalidator.ValueStringsAre(stringvalidator.LengthAtMost(63)),
			},
			Optional: true,
		},
		"tag": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Tags for address group object",
			Validators: []validator.List{
				listvalidator.SizeAtMost(64),
				listvalidator.ValueStringsAre(stringvalidator.LengthAtMost(127)),
			},
			Optional: true,
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

// AddressGroupsDataSourceSchema defines the schema for AddressGroups data source
var AddressGroupsDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "AddressGroup data source",
	Attributes: map[string]dsschema.Attribute{
		"description": dsschema.StringAttribute{
			MarkdownDescription: "Description",
			Computed:            true,
		},
		"device": dsschema.StringAttribute{
			MarkdownDescription: "The device in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			Computed:            true,
		},
		"dynamic": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Dynamic\n\n> ℹ️ **Note:** You must specify exactly one of `dynamic` and `static`.",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"filter": dsschema.StringAttribute{
					MarkdownDescription: "Tag based filter defining group membership",
					Computed:            true,
				},
			},
		},
		"folder": dsschema.StringAttribute{
			MarkdownDescription: "The folder in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			Computed:            true,
		},
		"id": dsschema.StringAttribute{
			MarkdownDescription: "The UUID of the address group",
			Required:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "The name of the address group",
			Optional:            true,
			Computed:            true,
		},
		"snippet": dsschema.StringAttribute{
			MarkdownDescription: "The snippet in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			Computed:            true,
		},
		"static": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Static\n\n> ℹ️ **Note:** You must specify exactly one of `dynamic` and `static`.",
			Computed:            true,
		},
		"tag": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Tags for address group object",
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
	},
}

// AddressGroupsListModel represents the data model for a list data source.
type AddressGroupsListModel struct {
	Tfid    types.String    `tfsdk:"tfid"`
	Data    []AddressGroups `tfsdk:"data"`
	Limit   types.Int64     `tfsdk:"limit"`
	Offset  types.Int64     `tfsdk:"offset"`
	Name    types.String    `tfsdk:"name"`
	Total   types.Int64     `tfsdk:"total"`
	Folder  types.String    `tfsdk:"folder"`
	Snippet types.String    `tfsdk:"snippet"`
	Device  types.String    `tfsdk:"device"`
}

// AddressGroupsListDataSourceSchema defines the schema for a list data source.
var AddressGroupsListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: AddressGroupsDataSourceSchema.Attributes,
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
