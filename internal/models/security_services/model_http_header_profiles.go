package models

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	dsschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Package: security_services
// This file contains models for the security_services SDK package

// HttpHeaderProfiles represents the Terraform model for HttpHeaderProfiles
type HttpHeaderProfiles struct {
	Tfid                types.String          `tfsdk:"tfid"`
	Description         basetypes.StringValue `tfsdk:"description"`
	Device              basetypes.StringValue `tfsdk:"device"`
	Folder              basetypes.StringValue `tfsdk:"folder"`
	HttpHeaderInsertion basetypes.ListValue   `tfsdk:"http_header_insertion"`
	Id                  basetypes.StringValue `tfsdk:"id"`
	Name                basetypes.StringValue `tfsdk:"name"`
	Snippet             basetypes.StringValue `tfsdk:"snippet"`
}

// HttpHeaderProfilesHttpHeaderInsertionInner represents a nested structure within the HttpHeaderProfiles model
type HttpHeaderProfilesHttpHeaderInsertionInner struct {
	Name basetypes.StringValue `tfsdk:"name"`
	Type basetypes.ListValue   `tfsdk:"type"`
}

// HttpHeaderProfilesHttpHeaderInsertionInnerTypeInner represents a nested structure within the HttpHeaderProfiles model
type HttpHeaderProfilesHttpHeaderInsertionInnerTypeInner struct {
	Domains basetypes.ListValue   `tfsdk:"domains"`
	Headers basetypes.ListValue   `tfsdk:"headers"`
	Name    basetypes.StringValue `tfsdk:"name"`
}

// HttpHeaderProfilesHttpHeaderInsertionInnerTypeInnerHeadersInner represents a nested structure within the HttpHeaderProfiles model
type HttpHeaderProfilesHttpHeaderInsertionInnerTypeInnerHeadersInner struct {
	Header basetypes.StringValue `tfsdk:"header"`
	Log    basetypes.BoolValue   `tfsdk:"log"`
	Name   basetypes.StringValue `tfsdk:"name"`
	Value  basetypes.StringValue `tfsdk:"value"`
}

// AttrTypes defines the attribute types for the HttpHeaderProfiles model.
func (o HttpHeaderProfiles) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":        basetypes.StringType{},
		"description": basetypes.StringType{},
		"device":      basetypes.StringType{},
		"folder":      basetypes.StringType{},
		"http_header_insertion": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"name": basetypes.StringType{},
				"type": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"domains": basetypes.ListType{ElemType: basetypes.StringType{}},
						"headers": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"header": basetypes.StringType{},
								"log":    basetypes.BoolType{},
								"name":   basetypes.StringType{},
								"value":  basetypes.StringType{},
							},
						}},
						"name": basetypes.StringType{},
					},
				}},
			},
		}},
		"id":      basetypes.StringType{},
		"name":    basetypes.StringType{},
		"snippet": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of HttpHeaderProfiles objects.
func (o HttpHeaderProfiles) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the HttpHeaderProfilesHttpHeaderInsertionInner model.
func (o HttpHeaderProfilesHttpHeaderInsertionInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name": basetypes.StringType{},
		"type": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"domains": basetypes.ListType{ElemType: basetypes.StringType{}},
				"headers": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"header": basetypes.StringType{},
						"log":    basetypes.BoolType{},
						"name":   basetypes.StringType{},
						"value":  basetypes.StringType{},
					},
				}},
				"name": basetypes.StringType{},
			},
		}},
	}
}

// AttrType returns the attribute type for a list of HttpHeaderProfilesHttpHeaderInsertionInner objects.
func (o HttpHeaderProfilesHttpHeaderInsertionInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the HttpHeaderProfilesHttpHeaderInsertionInnerTypeInner model.
func (o HttpHeaderProfilesHttpHeaderInsertionInnerTypeInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"domains": basetypes.ListType{ElemType: basetypes.StringType{}},
		"headers": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"header": basetypes.StringType{},
				"log":    basetypes.BoolType{},
				"name":   basetypes.StringType{},
				"value":  basetypes.StringType{},
			},
		}},
		"name": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of HttpHeaderProfilesHttpHeaderInsertionInnerTypeInner objects.
func (o HttpHeaderProfilesHttpHeaderInsertionInnerTypeInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the HttpHeaderProfilesHttpHeaderInsertionInnerTypeInnerHeadersInner model.
func (o HttpHeaderProfilesHttpHeaderInsertionInnerTypeInnerHeadersInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"header": basetypes.StringType{},
		"log":    basetypes.BoolType{},
		"name":   basetypes.StringType{},
		"value":  basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of HttpHeaderProfilesHttpHeaderInsertionInnerTypeInnerHeadersInner objects.
func (o HttpHeaderProfilesHttpHeaderInsertionInnerTypeInnerHeadersInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// HttpHeaderProfilesResourceSchema defines the schema for HttpHeaderProfiles resource
var HttpHeaderProfilesResourceSchema = schema.Schema{
	MarkdownDescription: "HttpHeaderProfile resource",
	Attributes: map[string]schema.Attribute{
		"description": schema.StringAttribute{
			MarkdownDescription: "The description of the HTTP header profile",
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
		"http_header_insertion": schema.ListNestedAttribute{
			MarkdownDescription: "A list of HTTP header profile rules",
			Optional:            true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"name": schema.StringAttribute{
						MarkdownDescription: "The name of the HTTP header insertion rule",
						Required:            true,
					},
					"type": schema.ListNestedAttribute{
						MarkdownDescription: "A list of HTTP header insertion definitions (_This should be an object rather than an array_)",
						Required:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"domains": schema.ListAttribute{
									ElementType:         types.StringType,
									MarkdownDescription: "A list of DNS domains",
									Required:            true,
								},
								"headers": schema.ListNestedAttribute{
									MarkdownDescription: "Headers",
									Required:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"header": schema.StringAttribute{
												MarkdownDescription: "The HTTP header string",
												Required:            true,
											},
											"log": schema.BoolAttribute{
												MarkdownDescription: "Log the use of this HTTP header insertion?",
												Optional:            true,
												Computed:            true,
												Default:             booldefault.StaticBool(false),
											},
											"name": schema.StringAttribute{
												MarkdownDescription: "An auto-generated name (_This should be removed_)",
												Required:            true,
											},
											"value": schema.StringAttribute{
												MarkdownDescription: "The value associated with the HTTP header",
												Required:            true,
											},
										},
									},
								},
								"name": schema.StringAttribute{
									MarkdownDescription: "The HTTP header insertion type (_This is a predefined list in the UI_)",
									Required:            true,
								},
							},
						},
					},
				},
			},
		},
		"id": schema.StringAttribute{
			MarkdownDescription: "The UUID of the HTTP header profile",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"name": schema.StringAttribute{
			MarkdownDescription: "The name of the HTTP header profile",
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
		"tfid": schema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
	},
}

// HttpHeaderProfilesDataSourceSchema defines the schema for HttpHeaderProfiles data source
var HttpHeaderProfilesDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "HttpHeaderProfile data source",
	Attributes: map[string]dsschema.Attribute{
		"description": dsschema.StringAttribute{
			MarkdownDescription: "The description of the HTTP header profile",
			Computed:            true,
		},
		"device": dsschema.StringAttribute{
			MarkdownDescription: "The device in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Computed:            true,
		},
		"folder": dsschema.StringAttribute{
			MarkdownDescription: "The folder in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Computed:            true,
		},
		"http_header_insertion": dsschema.ListNestedAttribute{
			MarkdownDescription: "A list of HTTP header profile rules",
			Computed:            true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: map[string]dsschema.Attribute{
					"name": dsschema.StringAttribute{
						MarkdownDescription: "The name of the HTTP header insertion rule",
						Computed:            true,
					},
					"type": dsschema.ListNestedAttribute{
						MarkdownDescription: "A list of HTTP header insertion definitions (_This should be an object rather than an array_)",
						Computed:            true,
						NestedObject: dsschema.NestedAttributeObject{
							Attributes: map[string]dsschema.Attribute{
								"domains": dsschema.ListAttribute{
									ElementType:         types.StringType,
									MarkdownDescription: "A list of DNS domains",
									Computed:            true,
								},
								"headers": dsschema.ListNestedAttribute{
									MarkdownDescription: "Headers",
									Computed:            true,
									NestedObject: dsschema.NestedAttributeObject{
										Attributes: map[string]dsschema.Attribute{
											"header": dsschema.StringAttribute{
												MarkdownDescription: "The HTTP header string",
												Computed:            true,
											},
											"log": dsschema.BoolAttribute{
												MarkdownDescription: "Log the use of this HTTP header insertion?",
												Computed:            true,
											},
											"name": dsschema.StringAttribute{
												MarkdownDescription: "An auto-generated name (_This should be removed_)",
												Computed:            true,
											},
											"value": dsschema.StringAttribute{
												MarkdownDescription: "The value associated with the HTTP header",
												Computed:            true,
											},
										},
									},
								},
								"name": dsschema.StringAttribute{
									MarkdownDescription: "The HTTP header insertion type (_This is a predefined list in the UI_)",
									Computed:            true,
								},
							},
						},
					},
				},
			},
		},
		"id": dsschema.StringAttribute{
			MarkdownDescription: "The UUID of the HTTP header profile",
			Required:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "The name of the HTTP header profile",
			Optional:            true,
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

// HttpHeaderProfilesListModel represents the data model for a list data source.
type HttpHeaderProfilesListModel struct {
	Tfid    types.String         `tfsdk:"tfid"`
	Data    []HttpHeaderProfiles `tfsdk:"data"`
	Limit   types.Int64          `tfsdk:"limit"`
	Offset  types.Int64          `tfsdk:"offset"`
	Name    types.String         `tfsdk:"name"`
	Total   types.Int64          `tfsdk:"total"`
	Folder  types.String         `tfsdk:"folder"`
	Snippet types.String         `tfsdk:"snippet"`
	Device  types.String         `tfsdk:"device"`
}

// HttpHeaderProfilesListDataSourceSchema defines the schema for a list data source.
var HttpHeaderProfilesListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: HttpHeaderProfilesDataSourceSchema.Attributes,
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
