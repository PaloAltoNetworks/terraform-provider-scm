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
	"github.com/paloaltonetworks/terraform-provider-scm/internal/utils"
)

// Package: security_services
// This file contains models for the security_services SDK package

// DataObjects represents the Terraform model for DataObjects
type DataObjects struct {
	Tfid            types.String          `tfsdk:"tfid"`
	Description     basetypes.StringValue `tfsdk:"description"`
	Device          basetypes.StringValue `tfsdk:"device"`
	DisableOverride basetypes.StringValue `tfsdk:"disable_override"`
	Folder          basetypes.StringValue `tfsdk:"folder"`
	Id              basetypes.StringValue `tfsdk:"id"`
	Name            basetypes.StringValue `tfsdk:"name"`
	PatternType     basetypes.ObjectValue `tfsdk:"pattern_type"`
	Snippet         basetypes.StringValue `tfsdk:"snippet"`
}

// DataObjectsPatternType represents a nested structure within the DataObjects model
type DataObjectsPatternType struct {
	FileProperties basetypes.ObjectValue `tfsdk:"file_properties"`
	Predefined     basetypes.ObjectValue `tfsdk:"predefined"`
	Regex          basetypes.ObjectValue `tfsdk:"regex"`
}

// DataObjectsPatternTypeFileProperties represents a nested structure within the DataObjects model
type DataObjectsPatternTypeFileProperties struct {
	Pattern basetypes.ListValue `tfsdk:"pattern"`
}

// DataObjectsPatternTypeFilePropertiesPatternInner represents a nested structure within the DataObjects model
type DataObjectsPatternTypeFilePropertiesPatternInner struct {
	FileProperty  basetypes.StringValue `tfsdk:"file_property"`
	FileType      basetypes.StringValue `tfsdk:"file_type"`
	Name          basetypes.StringValue `tfsdk:"name"`
	PropertyValue basetypes.StringValue `tfsdk:"property_value"`
}

// DataObjectsPatternTypePredefined represents a nested structure within the DataObjects model
type DataObjectsPatternTypePredefined struct {
	Pattern basetypes.ListValue `tfsdk:"pattern"`
}

// DataObjectsPatternTypePredefinedPatternInner represents a nested structure within the DataObjects model
type DataObjectsPatternTypePredefinedPatternInner struct {
	FileType basetypes.ListValue   `tfsdk:"file_type"`
	Name     basetypes.StringValue `tfsdk:"name"`
}

// DataObjectsPatternTypeRegex represents a nested structure within the DataObjects model
type DataObjectsPatternTypeRegex struct {
	Pattern basetypes.ListValue `tfsdk:"pattern"`
}

// DataObjectsPatternTypeRegexPatternInner represents a nested structure within the DataObjects model
type DataObjectsPatternTypeRegexPatternInner struct {
	FileType basetypes.ListValue   `tfsdk:"file_type"`
	Name     basetypes.StringValue `tfsdk:"name"`
	Regex    basetypes.StringValue `tfsdk:"regex"`
}

// AttrTypes defines the attribute types for the DataObjects model.
func (o DataObjects) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":             basetypes.StringType{},
		"description":      basetypes.StringType{},
		"device":           basetypes.StringType{},
		"disable_override": basetypes.StringType{},
		"folder":           basetypes.StringType{},
		"id":               basetypes.StringType{},
		"name":             basetypes.StringType{},
		"pattern_type": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"file_properties": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"pattern": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"file_property":  basetypes.StringType{},
								"file_type":      basetypes.StringType{},
								"name":           basetypes.StringType{},
								"property_value": basetypes.StringType{},
							},
						}},
					},
				},
				"predefined": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"pattern": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"file_type": basetypes.ListType{ElemType: basetypes.StringType{}},
								"name":      basetypes.StringType{},
							},
						}},
					},
				},
				"regex": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"pattern": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"file_type": basetypes.ListType{ElemType: basetypes.StringType{}},
								"name":      basetypes.StringType{},
								"regex":     basetypes.StringType{},
							},
						}},
					},
				},
			},
		},
		"snippet": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of DataObjects objects.
func (o DataObjects) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the DataObjectsPatternType model.
func (o DataObjectsPatternType) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"file_properties": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"pattern": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"file_property":  basetypes.StringType{},
						"file_type":      basetypes.StringType{},
						"name":           basetypes.StringType{},
						"property_value": basetypes.StringType{},
					},
				}},
			},
		},
		"predefined": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"pattern": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"file_type": basetypes.ListType{ElemType: basetypes.StringType{}},
						"name":      basetypes.StringType{},
					},
				}},
			},
		},
		"regex": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"pattern": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"file_type": basetypes.ListType{ElemType: basetypes.StringType{}},
						"name":      basetypes.StringType{},
						"regex":     basetypes.StringType{},
					},
				}},
			},
		},
	}
}

// AttrType returns the attribute type for a list of DataObjectsPatternType objects.
func (o DataObjectsPatternType) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the DataObjectsPatternTypeFileProperties model.
func (o DataObjectsPatternTypeFileProperties) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"pattern": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"file_property":  basetypes.StringType{},
				"file_type":      basetypes.StringType{},
				"name":           basetypes.StringType{},
				"property_value": basetypes.StringType{},
			},
		}},
	}
}

// AttrType returns the attribute type for a list of DataObjectsPatternTypeFileProperties objects.
func (o DataObjectsPatternTypeFileProperties) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the DataObjectsPatternTypeFilePropertiesPatternInner model.
func (o DataObjectsPatternTypeFilePropertiesPatternInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"file_property":  basetypes.StringType{},
		"file_type":      basetypes.StringType{},
		"name":           basetypes.StringType{},
		"property_value": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of DataObjectsPatternTypeFilePropertiesPatternInner objects.
func (o DataObjectsPatternTypeFilePropertiesPatternInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the DataObjectsPatternTypePredefined model.
func (o DataObjectsPatternTypePredefined) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"pattern": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"file_type": basetypes.ListType{ElemType: basetypes.StringType{}},
				"name":      basetypes.StringType{},
			},
		}},
	}
}

// AttrType returns the attribute type for a list of DataObjectsPatternTypePredefined objects.
func (o DataObjectsPatternTypePredefined) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the DataObjectsPatternTypePredefinedPatternInner model.
func (o DataObjectsPatternTypePredefinedPatternInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"file_type": basetypes.ListType{ElemType: basetypes.StringType{}},
		"name":      basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of DataObjectsPatternTypePredefinedPatternInner objects.
func (o DataObjectsPatternTypePredefinedPatternInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the DataObjectsPatternTypeRegex model.
func (o DataObjectsPatternTypeRegex) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"pattern": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"file_type": basetypes.ListType{ElemType: basetypes.StringType{}},
				"name":      basetypes.StringType{},
				"regex":     basetypes.StringType{},
			},
		}},
	}
}

// AttrType returns the attribute type for a list of DataObjectsPatternTypeRegex objects.
func (o DataObjectsPatternTypeRegex) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the DataObjectsPatternTypeRegexPatternInner model.
func (o DataObjectsPatternTypeRegexPatternInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"file_type": basetypes.ListType{ElemType: basetypes.StringType{}},
		"name":      basetypes.StringType{},
		"regex":     basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of DataObjectsPatternTypeRegexPatternInner objects.
func (o DataObjectsPatternTypeRegexPatternInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// DataObjectsResourceSchema defines the schema for DataObjects resource
var DataObjectsResourceSchema = schema.Schema{
	MarkdownDescription: "DataObject resource",
	Attributes: map[string]schema.Attribute{
		"description": schema.StringAttribute{
			MarkdownDescription: "The description of the data object",
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
		"disable_override": schema.StringAttribute{
			MarkdownDescription: "Disable override",
			Optional:            true,
		},
		"folder": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.ExactlyOneOf(
					path.MatchRelative().AtParent().AtName("device"),
					path.MatchRelative().AtParent().AtName("snippet"),
				),
				stringvalidator.LengthAtMost(64),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z\\d\\-_\\. ]+$"), "pattern must match "+"^[a-zA-Z\\d\\-_\\. ]+$"),
				utils.FolderValidator(),
			},
			MarkdownDescription: "The folder in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"id": schema.StringAttribute{
			MarkdownDescription: "The UUID of the data object",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"name": schema.StringAttribute{
			MarkdownDescription: "The name of the data object",
			Optional:            true,
		},
		"pattern_type": schema.SingleNestedAttribute{
			MarkdownDescription: "Pattern type",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"file_properties": schema.SingleNestedAttribute{
					MarkdownDescription: "File properties",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"pattern": schema.ListNestedAttribute{
							MarkdownDescription: "Pattern",
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"file_property": schema.StringAttribute{
										MarkdownDescription: "File property",
										Optional:            true,
									},
									"file_type": schema.StringAttribute{
										MarkdownDescription: "File type",
										Optional:            true,
									},
									"name": schema.StringAttribute{
										MarkdownDescription: "Name",
										Optional:            true,
									},
									"property_value": schema.StringAttribute{
										MarkdownDescription: "Property value",
										Optional:            true,
									},
								},
							},
						},
					},
				},
				"predefined": schema.SingleNestedAttribute{
					MarkdownDescription: "Predefined",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"pattern": schema.ListNestedAttribute{
							MarkdownDescription: "Pattern",
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"file_type": schema.ListAttribute{
										ElementType:         types.StringType,
										MarkdownDescription: "File type",
										Optional:            true,
									},
									"name": schema.StringAttribute{
										MarkdownDescription: "Name",
										Optional:            true,
									},
								},
							},
						},
					},
				},
				"regex": schema.SingleNestedAttribute{
					MarkdownDescription: "Regex",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"pattern": schema.ListNestedAttribute{
							MarkdownDescription: "Pattern",
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"file_type": schema.ListAttribute{
										ElementType:         types.StringType,
										MarkdownDescription: "File type",
										Optional:            true,
									},
									"name": schema.StringAttribute{
										MarkdownDescription: "Name",
										Optional:            true,
									},
									"regex": schema.StringAttribute{
										MarkdownDescription: "Regex",
										Optional:            true,
									},
								},
							},
						},
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

// DataObjectsDataSourceSchema defines the schema for DataObjects data source
var DataObjectsDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "DataObject data source",
	Attributes: map[string]dsschema.Attribute{
		"description": dsschema.StringAttribute{
			MarkdownDescription: "The description of the data object",
			Computed:            true,
		},
		"device": dsschema.StringAttribute{
			MarkdownDescription: "The device in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			Computed:            true,
		},
		"disable_override": dsschema.StringAttribute{
			MarkdownDescription: "Disable override",
			Computed:            true,
		},
		"folder": dsschema.StringAttribute{
			MarkdownDescription: "The folder in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			Computed:            true,
		},
		"id": dsschema.StringAttribute{
			MarkdownDescription: "The UUID of the data object",
			Required:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "The name of the data object",
			Optional:            true,
			Computed:            true,
		},
		"pattern_type": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Pattern type",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"file_properties": dsschema.SingleNestedAttribute{
					MarkdownDescription: "File properties",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"pattern": dsschema.ListNestedAttribute{
							MarkdownDescription: "Pattern",
							Computed:            true,
							NestedObject: dsschema.NestedAttributeObject{
								Attributes: map[string]dsschema.Attribute{
									"file_property": dsschema.StringAttribute{
										MarkdownDescription: "File property",
										Computed:            true,
									},
									"file_type": dsschema.StringAttribute{
										MarkdownDescription: "File type",
										Computed:            true,
									},
									"name": dsschema.StringAttribute{
										MarkdownDescription: "Name",
										Computed:            true,
									},
									"property_value": dsschema.StringAttribute{
										MarkdownDescription: "Property value",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"predefined": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Predefined",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"pattern": dsschema.ListNestedAttribute{
							MarkdownDescription: "Pattern",
							Computed:            true,
							NestedObject: dsschema.NestedAttributeObject{
								Attributes: map[string]dsschema.Attribute{
									"file_type": dsschema.ListAttribute{
										ElementType:         types.StringType,
										MarkdownDescription: "File type",
										Computed:            true,
									},
									"name": dsschema.StringAttribute{
										MarkdownDescription: "Name",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"regex": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Regex",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"pattern": dsschema.ListNestedAttribute{
							MarkdownDescription: "Pattern",
							Computed:            true,
							NestedObject: dsschema.NestedAttributeObject{
								Attributes: map[string]dsschema.Attribute{
									"file_type": dsschema.ListAttribute{
										ElementType:         types.StringType,
										MarkdownDescription: "File type",
										Computed:            true,
									},
									"name": dsschema.StringAttribute{
										MarkdownDescription: "Name",
										Computed:            true,
									},
									"regex": dsschema.StringAttribute{
										MarkdownDescription: "Regex",
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
		},
		"snippet": dsschema.StringAttribute{
			MarkdownDescription: "The snippet in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
	},
}

// DataObjectsListModel represents the data model for a list data source.
type DataObjectsListModel struct {
	Tfid    types.String  `tfsdk:"tfid"`
	Data    []DataObjects `tfsdk:"data"`
	Limit   types.Int64   `tfsdk:"limit"`
	Offset  types.Int64   `tfsdk:"offset"`
	Name    types.String  `tfsdk:"name"`
	Total   types.Int64   `tfsdk:"total"`
	Folder  types.String  `tfsdk:"folder"`
	Snippet types.String  `tfsdk:"snippet"`
	Device  types.String  `tfsdk:"device"`
}

// DataObjectsListDataSourceSchema defines the schema for a list data source.
var DataObjectsListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: DataObjectsDataSourceSchema.Attributes,
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
