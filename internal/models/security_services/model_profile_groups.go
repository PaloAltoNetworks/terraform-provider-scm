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

// Package: security_services
// This file contains models for the security_services SDK package

// ProfileGroups represents the Terraform model for ProfileGroups
type ProfileGroups struct {
	Tfid                     types.String          `tfsdk:"tfid"`
	AiSecurity               basetypes.ListValue   `tfsdk:"ai_security"`
	DataFiltering            basetypes.ListValue   `tfsdk:"data_filtering"`
	Device                   basetypes.StringValue `tfsdk:"device"`
	DnsSecurity              basetypes.ListValue   `tfsdk:"dns_security"`
	FileBlocking             basetypes.ListValue   `tfsdk:"file_blocking"`
	Folder                   basetypes.StringValue `tfsdk:"folder"`
	Id                       basetypes.StringValue `tfsdk:"id"`
	Name                     basetypes.StringValue `tfsdk:"name"`
	SaasSecurity             basetypes.ListValue   `tfsdk:"saas_security"`
	Snippet                  basetypes.StringValue `tfsdk:"snippet"`
	Spyware                  basetypes.ListValue   `tfsdk:"spyware"`
	UrlFiltering             basetypes.ListValue   `tfsdk:"url_filtering"`
	VirusAndWildfireAnalysis basetypes.ListValue   `tfsdk:"virus_and_wildfire_analysis"`
	Vulnerability            basetypes.ListValue   `tfsdk:"vulnerability"`
}

// AttrTypes defines the attribute types for the ProfileGroups model.
func (o ProfileGroups) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":                        basetypes.StringType{},
		"ai_security":                 basetypes.ListType{ElemType: basetypes.StringType{}},
		"data_filtering":              basetypes.ListType{ElemType: basetypes.StringType{}},
		"device":                      basetypes.StringType{},
		"dns_security":                basetypes.ListType{ElemType: basetypes.StringType{}},
		"file_blocking":               basetypes.ListType{ElemType: basetypes.StringType{}},
		"folder":                      basetypes.StringType{},
		"id":                          basetypes.StringType{},
		"name":                        basetypes.StringType{},
		"saas_security":               basetypes.ListType{ElemType: basetypes.StringType{}},
		"snippet":                     basetypes.StringType{},
		"spyware":                     basetypes.ListType{ElemType: basetypes.StringType{}},
		"url_filtering":               basetypes.ListType{ElemType: basetypes.StringType{}},
		"virus_and_wildfire_analysis": basetypes.ListType{ElemType: basetypes.StringType{}},
		"vulnerability":               basetypes.ListType{ElemType: basetypes.StringType{}},
	}
}

// AttrType returns the attribute type for a list of ProfileGroups objects.
func (o ProfileGroups) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// ProfileGroupsResourceSchema defines the schema for ProfileGroups resource
var ProfileGroupsResourceSchema = schema.Schema{
	MarkdownDescription: "ProfileGroup resource",
	Attributes: map[string]schema.Attribute{
		"ai_security": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Ai security",
			Optional:            true,
		},
		"data_filtering": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Data filtering",
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
		"dns_security": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Dns security",
			Optional:            true,
		},
		"file_blocking": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "File blocking",
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
			},
			MarkdownDescription: "The folder in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"id": schema.StringAttribute{
			MarkdownDescription: "The UUID of the profile group",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"name": schema.StringAttribute{
			MarkdownDescription: "The name of the profile group",
			Required:            true,
		},
		"saas_security": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Saas security",
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
		"spyware": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Spyware",
			Optional:            true,
		},
		"tfid": schema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"url_filtering": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Url filtering",
			Optional:            true,
		},
		"virus_and_wildfire_analysis": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Virus and wildfire analysis",
			Optional:            true,
		},
		"vulnerability": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Vulnerability",
			Optional:            true,
		},
	},
}

// ProfileGroupsDataSourceSchema defines the schema for ProfileGroups data source
var ProfileGroupsDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "ProfileGroup data source",
	Attributes: map[string]dsschema.Attribute{
		"ai_security": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Ai security",
			Computed:            true,
		},
		"data_filtering": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Data filtering",
			Computed:            true,
		},
		"device": dsschema.StringAttribute{
			MarkdownDescription: "The device in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			Computed:            true,
		},
		"dns_security": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Dns security",
			Computed:            true,
		},
		"file_blocking": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "File blocking",
			Computed:            true,
		},
		"folder": dsschema.StringAttribute{
			MarkdownDescription: "The folder in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			Computed:            true,
		},
		"id": dsschema.StringAttribute{
			MarkdownDescription: "The UUID of the profile group",
			Required:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "The name of the profile group",
			Optional:            true,
			Computed:            true,
		},
		"saas_security": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Saas security",
			Computed:            true,
		},
		"snippet": dsschema.StringAttribute{
			MarkdownDescription: "The snippet in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			Computed:            true,
		},
		"spyware": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Spyware",
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
		"url_filtering": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Url filtering",
			Computed:            true,
		},
		"virus_and_wildfire_analysis": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Virus and wildfire analysis",
			Computed:            true,
		},
		"vulnerability": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Vulnerability",
			Computed:            true,
		},
	},
}

// ProfileGroupsListModel represents the data model for a list data source.
type ProfileGroupsListModel struct {
	Tfid    types.String    `tfsdk:"tfid"`
	Data    []ProfileGroups `tfsdk:"data"`
	Limit   types.Int64     `tfsdk:"limit"`
	Offset  types.Int64     `tfsdk:"offset"`
	Name    types.String    `tfsdk:"name"`
	Total   types.Int64     `tfsdk:"total"`
	Folder  types.String    `tfsdk:"folder"`
	Snippet types.String    `tfsdk:"snippet"`
	Device  types.String    `tfsdk:"device"`
}

// ProfileGroupsListDataSourceSchema defines the schema for a list data source.
var ProfileGroupsListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: ProfileGroupsDataSourceSchema.Attributes,
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
