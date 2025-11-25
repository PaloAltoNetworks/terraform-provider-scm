package models

import (
	"regexp"

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

// Package: objects
// This file contains models for the objects SDK package

// AutoTagActions represents the Terraform model for AutoTagActions
type AutoTagActions struct {
	Tfid           types.String          `tfsdk:"tfid"`
	Actions        basetypes.ListValue   `tfsdk:"actions"`
	Description    basetypes.StringValue `tfsdk:"description"`
	Device         basetypes.StringValue `tfsdk:"device"`
	Filter         basetypes.StringValue `tfsdk:"filter"`
	Folder         basetypes.StringValue `tfsdk:"folder"`
	LogType        basetypes.StringValue `tfsdk:"log_type"`
	Name           basetypes.StringValue `tfsdk:"name"`
	Quarantine     basetypes.BoolValue   `tfsdk:"quarantine"`
	SendToPanorama basetypes.BoolValue   `tfsdk:"send_to_panorama"`
	Snippet        basetypes.StringValue `tfsdk:"snippet"`
}

// AutoTagActionsActionsInner represents a nested structure within the AutoTagActions model
type AutoTagActionsActionsInner struct {
	Name basetypes.StringValue `tfsdk:"name"`
	Type basetypes.ObjectValue `tfsdk:"type"`
}

// AutoTagActionsActionsInnerType represents a nested structure within the AutoTagActions model
type AutoTagActionsActionsInnerType struct {
	Tagging basetypes.ObjectValue `tfsdk:"tagging"`
}

// AutoTagActionsActionsInnerTypeTagging represents a nested structure within the AutoTagActions model
type AutoTagActionsActionsInnerTypeTagging struct {
	Action  basetypes.StringValue `tfsdk:"action"`
	Tags    basetypes.ListValue   `tfsdk:"tags"`
	Target  basetypes.StringValue `tfsdk:"target"`
	Timeout basetypes.Int64Value  `tfsdk:"timeout"`
}

// AttrTypes defines the attribute types for the AutoTagActions model.
func (o AutoTagActions) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid": basetypes.StringType{},
		"actions": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"name": basetypes.StringType{},
				"type": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"tagging": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"action":  basetypes.StringType{},
								"tags":    basetypes.ListType{ElemType: basetypes.StringType{}},
								"target":  basetypes.StringType{},
								"timeout": basetypes.Int64Type{},
							},
						},
					},
				},
			},
		}},
		"description":      basetypes.StringType{},
		"device":           basetypes.StringType{},
		"filter":           basetypes.StringType{},
		"folder":           basetypes.StringType{},
		"log_type":         basetypes.StringType{},
		"name":             basetypes.StringType{},
		"quarantine":       basetypes.BoolType{},
		"send_to_panorama": basetypes.BoolType{},
		"snippet":          basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of AutoTagActions objects.
func (o AutoTagActions) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the AutoTagActionsActionsInner model.
func (o AutoTagActionsActionsInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name": basetypes.StringType{},
		"type": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"tagging": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"action":  basetypes.StringType{},
						"tags":    basetypes.ListType{ElemType: basetypes.StringType{}},
						"target":  basetypes.StringType{},
						"timeout": basetypes.Int64Type{},
					},
				},
			},
		},
	}
}

// AttrType returns the attribute type for a list of AutoTagActionsActionsInner objects.
func (o AutoTagActionsActionsInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the AutoTagActionsActionsInnerType model.
func (o AutoTagActionsActionsInnerType) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tagging": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"action":  basetypes.StringType{},
				"tags":    basetypes.ListType{ElemType: basetypes.StringType{}},
				"target":  basetypes.StringType{},
				"timeout": basetypes.Int64Type{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of AutoTagActionsActionsInnerType objects.
func (o AutoTagActionsActionsInnerType) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the AutoTagActionsActionsInnerTypeTagging model.
func (o AutoTagActionsActionsInnerTypeTagging) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"action":  basetypes.StringType{},
		"tags":    basetypes.ListType{ElemType: basetypes.StringType{}},
		"target":  basetypes.StringType{},
		"timeout": basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of AutoTagActionsActionsInnerTypeTagging objects.
func (o AutoTagActionsActionsInnerTypeTagging) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AutoTagActionsResourceSchema defines the schema for AutoTagActions resource
var AutoTagActionsResourceSchema = schema.Schema{
	MarkdownDescription: "AutoTagAction resource",
	Attributes: map[string]schema.Attribute{
		"actions": schema.ListNestedAttribute{
			MarkdownDescription: "Actions",
			Optional:            true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"name": schema.StringAttribute{
						MarkdownDescription: "Name",
						Required:            true,
					},
					"type": schema.SingleNestedAttribute{
						MarkdownDescription: "Type",
						Required:            true,
						Attributes: map[string]schema.Attribute{
							"tagging": schema.SingleNestedAttribute{
								MarkdownDescription: "Tagging",
								Required:            true,
								Attributes: map[string]schema.Attribute{
									"action": schema.StringAttribute{
										Validators: []validator.String{
											stringvalidator.OneOf("add-tag", "remove-tag"),
										},
										MarkdownDescription: "Add or Remove tag option",
										Required:            true,
									},
									"tags": schema.ListAttribute{
										ElementType:         types.StringType,
										MarkdownDescription: "Tags for address object",
										Validators: []validator.List{
											listvalidator.SizeAtMost(64),
											listvalidator.ValueStringsAre(stringvalidator.LengthAtMost(127)),
										},
										Optional: true,
									},
									"target": schema.StringAttribute{
										MarkdownDescription: "Source or Destination Address, User, X-Forwarded-For Address",
										Required:            true,
									},
									"timeout": schema.Int64Attribute{
										MarkdownDescription: "Timeout",
										Optional:            true,
									},
								},
							},
						},
					},
				},
			},
		},
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
		"filter": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(2047),
			},
			MarkdownDescription: "Tag based filter defining group membership e.g. `tag1 AND tag2 OR tag3`",
			Required:            true,
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
		"log_type": schema.StringAttribute{
			MarkdownDescription: "Log type",
			Computed:            true,
		},
		"name": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(63),
			},
			MarkdownDescription: "Alphanumeric string [ 0-9a-zA-Z._-]",
			Required:            true,
		},
		"quarantine": schema.BoolAttribute{
			MarkdownDescription: "Quarantine",
			Optional:            true,
		},
		"send_to_panorama": schema.BoolAttribute{
			MarkdownDescription: "Send to panorama",
			Optional:            true,
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

// AutoTagActionsDataSourceSchema defines the schema for AutoTagActions data source
var AutoTagActionsDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "AutoTagAction data source",
	Attributes: map[string]dsschema.Attribute{
		"actions": dsschema.ListNestedAttribute{
			MarkdownDescription: "Actions",
			Computed:            true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: map[string]dsschema.Attribute{
					"name": dsschema.StringAttribute{
						MarkdownDescription: "Name",
						Computed:            true,
					},
					"type": dsschema.SingleNestedAttribute{
						MarkdownDescription: "Type",
						Computed:            true,
						Attributes: map[string]dsschema.Attribute{
							"tagging": dsschema.SingleNestedAttribute{
								MarkdownDescription: "Tagging",
								Computed:            true,
								Attributes: map[string]dsschema.Attribute{
									"action": dsschema.StringAttribute{
										MarkdownDescription: "Add or Remove tag option",
										Computed:            true,
									},
									"tags": dsschema.ListAttribute{
										ElementType:         types.StringType,
										MarkdownDescription: "Tags for address object",
										Computed:            true,
									},
									"target": dsschema.StringAttribute{
										MarkdownDescription: "Source or Destination Address, User, X-Forwarded-For Address",
										Computed:            true,
									},
									"timeout": dsschema.Int64Attribute{
										MarkdownDescription: "Timeout",
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
		},
		"description": dsschema.StringAttribute{
			MarkdownDescription: "Description",
			Computed:            true,
		},
		"device": dsschema.StringAttribute{
			MarkdownDescription: "The device in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Computed:            true,
		},
		"filter": dsschema.StringAttribute{
			MarkdownDescription: "Tag based filter defining group membership e.g. `tag1 AND tag2 OR tag3`",
			Computed:            true,
		},
		"folder": dsschema.StringAttribute{
			MarkdownDescription: "The folder in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Computed:            true,
		},
		"log_type": dsschema.StringAttribute{
			MarkdownDescription: "Log type",
			Computed:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "Alphanumeric string [ 0-9a-zA-Z._-]",
			Optional:            true,
			Computed:            true,
		},
		"quarantine": dsschema.BoolAttribute{
			MarkdownDescription: "Quarantine",
			Computed:            true,
		},
		"send_to_panorama": dsschema.BoolAttribute{
			MarkdownDescription: "Send to panorama",
			Computed:            true,
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

// AutoTagActionsListModel represents the data model for a list data source.
type AutoTagActionsListModel struct {
	Tfid    types.String     `tfsdk:"tfid"`
	Data    []AutoTagActions `tfsdk:"data"`
	Limit   types.Int64      `tfsdk:"limit"`
	Offset  types.Int64      `tfsdk:"offset"`
	Name    types.String     `tfsdk:"name"`
	Total   types.Int64      `tfsdk:"total"`
	Folder  types.String     `tfsdk:"folder"`
	Snippet types.String     `tfsdk:"snippet"`
	Device  types.String     `tfsdk:"device"`
}

// AutoTagActionsListDataSourceSchema defines the schema for a list data source.
var AutoTagActionsListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: AutoTagActionsDataSourceSchema.Attributes,
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
