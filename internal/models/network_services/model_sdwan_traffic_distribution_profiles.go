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

// Package: network_services
// This file contains models for the network_services SDK package

// SdwanTrafficDistributionProfiles represents the Terraform model for SdwanTrafficDistributionProfiles
type SdwanTrafficDistributionProfiles struct {
	Tfid                types.String          `tfsdk:"tfid"`
	Device              basetypes.StringValue `tfsdk:"device"`
	Folder              basetypes.StringValue `tfsdk:"folder"`
	Id                  basetypes.StringValue `tfsdk:"id"`
	LinkTags            basetypes.ListValue   `tfsdk:"link_tags"`
	Name                basetypes.StringValue `tfsdk:"name"`
	Snippet             basetypes.StringValue `tfsdk:"snippet"`
	TrafficDistribution basetypes.StringValue `tfsdk:"traffic_distribution"`
}

// SdwanTrafficDistributionProfilesLinkTagsInner represents a nested structure within the SdwanTrafficDistributionProfiles model
type SdwanTrafficDistributionProfilesLinkTagsInner struct {
	Name   basetypes.StringValue `tfsdk:"name"`
	Weight basetypes.Int64Value  `tfsdk:"weight"`
}

// AttrTypes defines the attribute types for the SdwanTrafficDistributionProfiles model.
func (o SdwanTrafficDistributionProfiles) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":   basetypes.StringType{},
		"device": basetypes.StringType{},
		"folder": basetypes.StringType{},
		"id":     basetypes.StringType{},
		"link_tags": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"name":   basetypes.StringType{},
				"weight": basetypes.Int64Type{},
			},
		}},
		"name":                 basetypes.StringType{},
		"snippet":              basetypes.StringType{},
		"traffic_distribution": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of SdwanTrafficDistributionProfiles objects.
func (o SdwanTrafficDistributionProfiles) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the SdwanTrafficDistributionProfilesLinkTagsInner model.
func (o SdwanTrafficDistributionProfilesLinkTagsInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name":   basetypes.StringType{},
		"weight": basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of SdwanTrafficDistributionProfilesLinkTagsInner objects.
func (o SdwanTrafficDistributionProfilesLinkTagsInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// SdwanTrafficDistributionProfilesResourceSchema defines the schema for SdwanTrafficDistributionProfiles resource
var SdwanTrafficDistributionProfilesResourceSchema = schema.Schema{
	MarkdownDescription: "SdwanTrafficDistributionProfile resource",
	Attributes: map[string]schema.Attribute{
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
					path.MatchRelative().AtParent().AtName("snippet"),
					path.MatchRelative().AtParent().AtName("device"),
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
		"link_tags": schema.ListNestedAttribute{
			MarkdownDescription: "Link tags",
			Optional:            true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"name": schema.StringAttribute{
						MarkdownDescription: "Name",
						Optional:            true,
					},
					"weight": schema.Int64Attribute{
						MarkdownDescription: "Weight",
						Optional:            true,
					},
				},
			},
		},
		"name": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(31),
			},
			MarkdownDescription: "Profile name",
			Required:            true,
		},
		"snippet": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.ExactlyOneOf(
					path.MatchRelative().AtParent().AtName("folder"),
					path.MatchRelative().AtParent().AtName("device"),
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
		"traffic_distribution": schema.StringAttribute{
			MarkdownDescription: "Traffic distribution",
			Optional:            true,
		},
	},
}

// SdwanTrafficDistributionProfilesDataSourceSchema defines the schema for SdwanTrafficDistributionProfiles data source
var SdwanTrafficDistributionProfilesDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "SdwanTrafficDistributionProfile data source",
	Attributes: map[string]dsschema.Attribute{
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
		"link_tags": dsschema.ListNestedAttribute{
			MarkdownDescription: "Link tags",
			Computed:            true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: map[string]dsschema.Attribute{
					"name": dsschema.StringAttribute{
						MarkdownDescription: "Name",
						Computed:            true,
					},
					"weight": dsschema.Int64Attribute{
						MarkdownDescription: "Weight",
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
		"snippet": dsschema.StringAttribute{
			MarkdownDescription: "The snippet in which the resource is defined",
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
		"traffic_distribution": dsschema.StringAttribute{
			MarkdownDescription: "Traffic distribution",
			Computed:            true,
		},
	},
}

// SdwanTrafficDistributionProfilesListModel represents the data model for a list data source.
type SdwanTrafficDistributionProfilesListModel struct {
	Tfid    types.String                       `tfsdk:"tfid"`
	Data    []SdwanTrafficDistributionProfiles `tfsdk:"data"`
	Limit   types.Int64                        `tfsdk:"limit"`
	Offset  types.Int64                        `tfsdk:"offset"`
	Name    types.String                       `tfsdk:"name"`
	Total   types.Int64                        `tfsdk:"total"`
	Folder  types.String                       `tfsdk:"folder"`
	Snippet types.String                       `tfsdk:"snippet"`
	Device  types.String                       `tfsdk:"device"`
}

// SdwanTrafficDistributionProfilesListDataSourceSchema defines the schema for a list data source.
var SdwanTrafficDistributionProfilesListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: SdwanTrafficDistributionProfilesDataSourceSchema.Attributes,
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
