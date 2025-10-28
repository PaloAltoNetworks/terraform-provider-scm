package models

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/boolvalidator"
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

// ApplicationFilters represents the Terraform model for ApplicationFilters
type ApplicationFilters struct {
	Tfid                    types.String          `tfsdk:"tfid"`
	Category                basetypes.ListValue   `tfsdk:"category"`
	Device                  basetypes.StringValue `tfsdk:"device"`
	Evasive                 basetypes.BoolValue   `tfsdk:"evasive"`
	ExcessiveBandwidthUse   basetypes.BoolValue   `tfsdk:"excessive_bandwidth_use"`
	Exclude                 basetypes.ListValue   `tfsdk:"exclude"`
	Folder                  basetypes.StringValue `tfsdk:"folder"`
	HasKnownVulnerabilities basetypes.BoolValue   `tfsdk:"has_known_vulnerabilities"`
	Id                      basetypes.StringValue `tfsdk:"id"`
	IsSaas                  basetypes.BoolValue   `tfsdk:"is_saas"`
	Name                    basetypes.StringValue `tfsdk:"name"`
	NewAppid                basetypes.BoolValue   `tfsdk:"new_appid"`
	Pervasive               basetypes.BoolValue   `tfsdk:"pervasive"`
	ProneToMisuse           basetypes.BoolValue   `tfsdk:"prone_to_misuse"`
	Risk                    basetypes.ListValue   `tfsdk:"risk"`
	SaasCertifications      basetypes.ListValue   `tfsdk:"saas_certifications"`
	SaasRisk                basetypes.ListValue   `tfsdk:"saas_risk"`
	Snippet                 basetypes.StringValue `tfsdk:"snippet"`
	Subcategory             basetypes.ListValue   `tfsdk:"subcategory"`
	Tagging                 basetypes.ObjectValue `tfsdk:"tagging"`
	Technology              basetypes.ListValue   `tfsdk:"technology"`
	TransfersFiles          basetypes.BoolValue   `tfsdk:"transfers_files"`
	TunnelsOtherApps        basetypes.BoolValue   `tfsdk:"tunnels_other_apps"`
	UsedByMalware           basetypes.BoolValue   `tfsdk:"used_by_malware"`
}

// ApplicationFiltersTagging represents a nested structure within the ApplicationFilters model
type ApplicationFiltersTagging struct {
	NoTag basetypes.BoolValue `tfsdk:"no_tag"`
	Tag   basetypes.ListValue `tfsdk:"tag"`
}

// AttrTypes defines the attribute types for the ApplicationFilters model.
func (o ApplicationFilters) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":                      basetypes.StringType{},
		"category":                  basetypes.ListType{ElemType: basetypes.StringType{}},
		"device":                    basetypes.StringType{},
		"evasive":                   basetypes.BoolType{},
		"excessive_bandwidth_use":   basetypes.BoolType{},
		"exclude":                   basetypes.ListType{ElemType: basetypes.StringType{}},
		"folder":                    basetypes.StringType{},
		"has_known_vulnerabilities": basetypes.BoolType{},
		"id":                        basetypes.StringType{},
		"is_saas":                   basetypes.BoolType{},
		"name":                      basetypes.StringType{},
		"new_appid":                 basetypes.BoolType{},
		"pervasive":                 basetypes.BoolType{},
		"prone_to_misuse":           basetypes.BoolType{},
		"risk":                      basetypes.ListType{ElemType: basetypes.Int64Type{}},
		"saas_certifications":       basetypes.ListType{ElemType: basetypes.StringType{}},
		"saas_risk":                 basetypes.ListType{ElemType: basetypes.StringType{}},
		"snippet":                   basetypes.StringType{},
		"subcategory":               basetypes.ListType{ElemType: basetypes.StringType{}},
		"tagging": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"no_tag": basetypes.BoolType{},
				"tag":    basetypes.ListType{ElemType: basetypes.StringType{}},
			},
		},
		"technology":         basetypes.ListType{ElemType: basetypes.StringType{}},
		"transfers_files":    basetypes.BoolType{},
		"tunnels_other_apps": basetypes.BoolType{},
		"used_by_malware":    basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of ApplicationFilters objects.
func (o ApplicationFilters) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ApplicationFiltersTagging model.
func (o ApplicationFiltersTagging) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"no_tag": basetypes.BoolType{},
		"tag":    basetypes.ListType{ElemType: basetypes.StringType{}},
	}
}

// AttrType returns the attribute type for a list of ApplicationFiltersTagging objects.
func (o ApplicationFiltersTagging) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// ApplicationFiltersResourceSchema defines the schema for ApplicationFilters resource
var ApplicationFiltersResourceSchema = schema.Schema{
	MarkdownDescription: "ApplicationFilter resource",
	Attributes: map[string]schema.Attribute{
		"category": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Category",
			Validators: []validator.List{
				listvalidator.ValueStringsAre(stringvalidator.LengthAtMost(128)),
			},
			Optional: true,
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
			MarkdownDescription: "The device in which the resource is defined",
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"evasive": schema.BoolAttribute{
			MarkdownDescription: "only True is a valid value",
			Optional:            true,
		},
		"excessive_bandwidth_use": schema.BoolAttribute{
			MarkdownDescription: "only True is a valid value",
			Optional:            true,
		},
		"exclude": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Exclude",
			Validators: []validator.List{
				listvalidator.ValueStringsAre(stringvalidator.LengthAtMost(63)),
			},
			Optional: true,
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
		"has_known_vulnerabilities": schema.BoolAttribute{
			MarkdownDescription: "only True is a valid value",
			Optional:            true,
		},
		"id": schema.StringAttribute{
			MarkdownDescription: "UUID of the resource",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"is_saas": schema.BoolAttribute{
			MarkdownDescription: "only True is a valid value",
			Optional:            true,
		},
		"name": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(31),
			},
			MarkdownDescription: "Alphanumeric string [ 0-9a-zA-Z._-]",
			Required:            true,
		},
		"new_appid": schema.BoolAttribute{
			MarkdownDescription: "only True is a valid value",
			Optional:            true,
		},
		"pervasive": schema.BoolAttribute{
			MarkdownDescription: "only True is a valid value",
			Optional:            true,
		},
		"prone_to_misuse": schema.BoolAttribute{
			MarkdownDescription: "only True is a valid value",
			Optional:            true,
		},
		"risk": schema.ListAttribute{
			ElementType:         types.Int64Type,
			MarkdownDescription: "Risk",
			Optional:            true,
		},
		"saas_certifications": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Saas certifications",
			Validators: []validator.List{
				listvalidator.ValueStringsAre(stringvalidator.LengthAtMost(32)),
			},
			Optional: true,
		},
		"saas_risk": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Saas risk",
			Validators: []validator.List{
				listvalidator.ValueStringsAre(stringvalidator.LengthAtMost(32)),
			},
			Optional: true,
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
		"subcategory": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Subcategory",
			Validators: []validator.List{
				listvalidator.ValueStringsAre(stringvalidator.LengthAtMost(128)),
			},
			Optional: true,
		},
		"tagging": schema.SingleNestedAttribute{
			MarkdownDescription: "Tagging",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"no_tag": schema.BoolAttribute{
					Validators: []validator.Bool{
						boolvalidator.ConflictsWith(
							path.MatchRelative().AtParent().AtName("tag"),
						),
					},
					MarkdownDescription: "No tag",
					Optional:            true,
				},
				"tag": schema.ListAttribute{
					ElementType:         types.StringType,
					MarkdownDescription: "Tag",
					Validators: []validator.List{
						listvalidator.ConflictsWith(
							path.MatchRelative().AtParent().AtName("no_tag"),
						),
						listvalidator.ValueStringsAre(stringvalidator.LengthAtMost(127)),
					},
					Optional: true,
				},
			},
		},
		"technology": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Technology",
			Validators: []validator.List{
				listvalidator.ValueStringsAre(stringvalidator.LengthAtMost(128)),
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
		"transfers_files": schema.BoolAttribute{
			MarkdownDescription: "only True is a valid value",
			Optional:            true,
		},
		"tunnels_other_apps": schema.BoolAttribute{
			MarkdownDescription: "only True is a valid value",
			Optional:            true,
		},
		"used_by_malware": schema.BoolAttribute{
			MarkdownDescription: "only True is a valid value",
			Optional:            true,
		},
	},
}

// ApplicationFiltersDataSourceSchema defines the schema for ApplicationFilters data source
var ApplicationFiltersDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "ApplicationFilter data source",
	Attributes: map[string]dsschema.Attribute{
		"category": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Category",
			Computed:            true,
		},
		"device": dsschema.StringAttribute{
			MarkdownDescription: "The device in which the resource is defined",
			Computed:            true,
		},
		"evasive": dsschema.BoolAttribute{
			MarkdownDescription: "only True is a valid value",
			Computed:            true,
		},
		"excessive_bandwidth_use": dsschema.BoolAttribute{
			MarkdownDescription: "only True is a valid value",
			Computed:            true,
		},
		"exclude": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Exclude",
			Computed:            true,
		},
		"folder": dsschema.StringAttribute{
			MarkdownDescription: "The folder in which the resource is defined",
			Computed:            true,
		},
		"has_known_vulnerabilities": dsschema.BoolAttribute{
			MarkdownDescription: "only True is a valid value",
			Computed:            true,
		},
		"id": dsschema.StringAttribute{
			MarkdownDescription: "UUID of the resource",
			Required:            true,
		},
		"is_saas": dsschema.BoolAttribute{
			MarkdownDescription: "only True is a valid value",
			Computed:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "Alphanumeric string [ 0-9a-zA-Z._-]",
			Optional:            true,
			Computed:            true,
		},
		"new_appid": dsschema.BoolAttribute{
			MarkdownDescription: "only True is a valid value",
			Computed:            true,
		},
		"pervasive": dsschema.BoolAttribute{
			MarkdownDescription: "only True is a valid value",
			Computed:            true,
		},
		"prone_to_misuse": dsschema.BoolAttribute{
			MarkdownDescription: "only True is a valid value",
			Computed:            true,
		},
		"risk": dsschema.ListAttribute{
			ElementType:         types.Int64Type,
			MarkdownDescription: "Risk",
			Computed:            true,
		},
		"saas_certifications": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Saas certifications",
			Computed:            true,
		},
		"saas_risk": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Saas risk",
			Computed:            true,
		},
		"snippet": dsschema.StringAttribute{
			MarkdownDescription: "The snippet in which the resource is defined",
			Computed:            true,
		},
		"subcategory": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Subcategory",
			Computed:            true,
		},
		"tagging": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Tagging",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"no_tag": dsschema.BoolAttribute{
					MarkdownDescription: "No tag",
					Computed:            true,
				},
				"tag": dsschema.ListAttribute{
					ElementType:         types.StringType,
					MarkdownDescription: "Tag",
					Computed:            true,
				},
			},
		},
		"technology": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Technology",
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
		"transfers_files": dsschema.BoolAttribute{
			MarkdownDescription: "only True is a valid value",
			Computed:            true,
		},
		"tunnels_other_apps": dsschema.BoolAttribute{
			MarkdownDescription: "only True is a valid value",
			Computed:            true,
		},
		"used_by_malware": dsschema.BoolAttribute{
			MarkdownDescription: "only True is a valid value",
			Computed:            true,
		},
	},
}

// ApplicationFiltersListModel represents the data model for a list data source.
type ApplicationFiltersListModel struct {
	Tfid    types.String         `tfsdk:"tfid"`
	Data    []ApplicationFilters `tfsdk:"data"`
	Limit   types.Int64          `tfsdk:"limit"`
	Offset  types.Int64          `tfsdk:"offset"`
	Name    types.String         `tfsdk:"name"`
	Total   types.Int64          `tfsdk:"total"`
	Folder  types.String         `tfsdk:"folder"`
	Snippet types.String         `tfsdk:"snippet"`
	Device  types.String         `tfsdk:"device"`
}

// ApplicationFiltersListDataSourceSchema defines the schema for a list data source.
var ApplicationFiltersListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: ApplicationFiltersDataSourceSchema.Attributes,
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
