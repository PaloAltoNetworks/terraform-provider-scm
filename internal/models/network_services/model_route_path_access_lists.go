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

// Package: network_services
// This file contains models for the network_services SDK package

// RoutePathAccessLists represents the Terraform model for RoutePathAccessLists
type RoutePathAccessLists struct {
	Tfid        types.String          `tfsdk:"tfid"`
	AspathEntry basetypes.ListValue   `tfsdk:"aspath_entry"`
	Description basetypes.StringValue `tfsdk:"description"`
	Device      basetypes.StringValue `tfsdk:"device"`
	Folder      basetypes.StringValue `tfsdk:"folder"`
	Id          basetypes.StringValue `tfsdk:"id"`
	Name        basetypes.StringValue `tfsdk:"name"`
	Snippet     basetypes.StringValue `tfsdk:"snippet"`
}

// RoutePathAccessListsAspathEntryInner represents a nested structure within the RoutePathAccessLists model
type RoutePathAccessListsAspathEntryInner struct {
	Action      basetypes.StringValue `tfsdk:"action"`
	AspathRegex basetypes.StringValue `tfsdk:"aspath_regex"`
	Name        basetypes.Int64Value  `tfsdk:"name"`
}

// AttrTypes defines the attribute types for the RoutePathAccessLists model.
func (o RoutePathAccessLists) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid": basetypes.StringType{},
		"aspath_entry": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"action":       basetypes.StringType{},
				"aspath_regex": basetypes.StringType{},
				"name":         basetypes.Int64Type{},
			},
		}},
		"description": basetypes.StringType{},
		"device":      basetypes.StringType{},
		"folder":      basetypes.StringType{},
		"id":          basetypes.StringType{},
		"name":        basetypes.StringType{},
		"snippet":     basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of RoutePathAccessLists objects.
func (o RoutePathAccessLists) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the RoutePathAccessListsAspathEntryInner model.
func (o RoutePathAccessListsAspathEntryInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"action":       basetypes.StringType{},
		"aspath_regex": basetypes.StringType{},
		"name":         basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of RoutePathAccessListsAspathEntryInner objects.
func (o RoutePathAccessListsAspathEntryInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// RoutePathAccessListsResourceSchema defines the schema for RoutePathAccessLists resource
var RoutePathAccessListsResourceSchema = schema.Schema{
	MarkdownDescription: "RoutePathAccessList resource",
	Attributes: map[string]schema.Attribute{
		"aspath_entry": schema.ListNestedAttribute{
			MarkdownDescription: "AS paths",
			Optional:            true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"action": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("deny", "permit"),
						},
						MarkdownDescription: "Action",
						Optional:            true,
					},
					"aspath_regex": schema.StringAttribute{
						MarkdownDescription: "AS path regular expression",
						Optional:            true,
					},
					"name": schema.Int64Attribute{
						Validators: []validator.Int64{
							int64validator.Between(1, 65535),
						},
						MarkdownDescription: "Sequence number",
						Optional:            true,
					},
				},
			},
		},
		"description": schema.StringAttribute{
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
			MarkdownDescription: "UUID of the resource",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"name": schema.StringAttribute{
			MarkdownDescription: "AS path access list name",
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

// RoutePathAccessListsDataSourceSchema defines the schema for RoutePathAccessLists data source
var RoutePathAccessListsDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "RoutePathAccessList data source",
	Attributes: map[string]dsschema.Attribute{
		"aspath_entry": dsschema.ListNestedAttribute{
			MarkdownDescription: "AS paths",
			Computed:            true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: map[string]dsschema.Attribute{
					"action": dsschema.StringAttribute{
						MarkdownDescription: "Action",
						Computed:            true,
					},
					"aspath_regex": dsschema.StringAttribute{
						MarkdownDescription: "AS path regular expression",
						Computed:            true,
					},
					"name": dsschema.Int64Attribute{
						MarkdownDescription: "Sequence number",
						Computed:            true,
					},
				},
			},
		},
		"description": dsschema.StringAttribute{
			MarkdownDescription: "Description",
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
			MarkdownDescription: "UUID of the resource",
			Required:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "AS path access list name",
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

// RoutePathAccessListsListModel represents the data model for a list data source.
type RoutePathAccessListsListModel struct {
	Tfid    types.String           `tfsdk:"tfid"`
	Data    []RoutePathAccessLists `tfsdk:"data"`
	Limit   types.Int64            `tfsdk:"limit"`
	Offset  types.Int64            `tfsdk:"offset"`
	Name    types.String           `tfsdk:"name"`
	Total   types.Int64            `tfsdk:"total"`
	Folder  types.String           `tfsdk:"folder"`
	Snippet types.String           `tfsdk:"snippet"`
	Device  types.String           `tfsdk:"device"`
}

// RoutePathAccessListsListDataSourceSchema defines the schema for a list data source.
var RoutePathAccessListsListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: RoutePathAccessListsDataSourceSchema.Attributes,
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
