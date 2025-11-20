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

// SdwanPathQualityProfiles represents the Terraform model for SdwanPathQualityProfiles
type SdwanPathQualityProfiles struct {
	Tfid    types.String          `tfsdk:"tfid"`
	Device  basetypes.StringValue `tfsdk:"device"`
	Folder  basetypes.StringValue `tfsdk:"folder"`
	Id      basetypes.StringValue `tfsdk:"id"`
	Metric  basetypes.ObjectValue `tfsdk:"metric"`
	Name    basetypes.StringValue `tfsdk:"name"`
	Snippet basetypes.StringValue `tfsdk:"snippet"`
}

// SdwanPathQualityProfilesMetric represents a nested structure within the SdwanPathQualityProfiles model
type SdwanPathQualityProfilesMetric struct {
	Jitter  basetypes.ObjectValue `tfsdk:"jitter"`
	Latency basetypes.ObjectValue `tfsdk:"latency"`
	PktLoss basetypes.ObjectValue `tfsdk:"pkt_loss"`
}

// SdwanPathQualityProfilesMetricJitter represents a nested structure within the SdwanPathQualityProfiles model
type SdwanPathQualityProfilesMetricJitter struct {
	Sensitivity basetypes.StringValue `tfsdk:"sensitivity"`
	Threshold   basetypes.Int64Value  `tfsdk:"threshold"`
}

// SdwanPathQualityProfilesMetricLatency represents a nested structure within the SdwanPathQualityProfiles model
type SdwanPathQualityProfilesMetricLatency struct {
	Sensitivity basetypes.StringValue `tfsdk:"sensitivity"`
	Threshold   basetypes.Int64Value  `tfsdk:"threshold"`
}

// SdwanPathQualityProfilesMetricPktLoss represents a nested structure within the SdwanPathQualityProfiles model
type SdwanPathQualityProfilesMetricPktLoss struct {
	Sensitivity basetypes.StringValue `tfsdk:"sensitivity"`
	Threshold   basetypes.Int64Value  `tfsdk:"threshold"`
}

// AttrTypes defines the attribute types for the SdwanPathQualityProfiles model.
func (o SdwanPathQualityProfiles) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":   basetypes.StringType{},
		"device": basetypes.StringType{},
		"folder": basetypes.StringType{},
		"id":     basetypes.StringType{},
		"metric": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"jitter": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"sensitivity": basetypes.StringType{},
						"threshold":   basetypes.Int64Type{},
					},
				},
				"latency": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"sensitivity": basetypes.StringType{},
						"threshold":   basetypes.Int64Type{},
					},
				},
				"pkt_loss": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"sensitivity": basetypes.StringType{},
						"threshold":   basetypes.Int64Type{},
					},
				},
			},
		},
		"name":    basetypes.StringType{},
		"snippet": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of SdwanPathQualityProfiles objects.
func (o SdwanPathQualityProfiles) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the SdwanPathQualityProfilesMetric model.
func (o SdwanPathQualityProfilesMetric) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"jitter": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"sensitivity": basetypes.StringType{},
				"threshold":   basetypes.Int64Type{},
			},
		},
		"latency": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"sensitivity": basetypes.StringType{},
				"threshold":   basetypes.Int64Type{},
			},
		},
		"pkt_loss": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"sensitivity": basetypes.StringType{},
				"threshold":   basetypes.Int64Type{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of SdwanPathQualityProfilesMetric objects.
func (o SdwanPathQualityProfilesMetric) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the SdwanPathQualityProfilesMetricJitter model.
func (o SdwanPathQualityProfilesMetricJitter) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"sensitivity": basetypes.StringType{},
		"threshold":   basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of SdwanPathQualityProfilesMetricJitter objects.
func (o SdwanPathQualityProfilesMetricJitter) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the SdwanPathQualityProfilesMetricLatency model.
func (o SdwanPathQualityProfilesMetricLatency) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"sensitivity": basetypes.StringType{},
		"threshold":   basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of SdwanPathQualityProfilesMetricLatency objects.
func (o SdwanPathQualityProfilesMetricLatency) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the SdwanPathQualityProfilesMetricPktLoss model.
func (o SdwanPathQualityProfilesMetricPktLoss) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"sensitivity": basetypes.StringType{},
		"threshold":   basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of SdwanPathQualityProfilesMetricPktLoss objects.
func (o SdwanPathQualityProfilesMetricPktLoss) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// SdwanPathQualityProfilesResourceSchema defines the schema for SdwanPathQualityProfiles resource
var SdwanPathQualityProfilesResourceSchema = schema.Schema{
	MarkdownDescription: "SdwanPathQualityProfile resource",
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
		"id": schema.StringAttribute{
			MarkdownDescription: "UUID of the resource",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"metric": schema.SingleNestedAttribute{
			MarkdownDescription: "Metric",
			Required:            true,
			Attributes: map[string]schema.Attribute{
				"jitter": schema.SingleNestedAttribute{
					MarkdownDescription: "Jitter",
					Required:            true,
					Attributes: map[string]schema.Attribute{
						"sensitivity": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("low", "medium", "high"),
							},
							MarkdownDescription: "Jitter sensitivity",
							Required:            true,
						},
						"threshold": schema.Int64Attribute{
							Validators: []validator.Int64{
								int64validator.Between(10, 2000),
							},
							MarkdownDescription: "Jitter threshold (ms)",
							Required:            true,
						},
					},
				},
				"latency": schema.SingleNestedAttribute{
					MarkdownDescription: "Latency",
					Required:            true,
					Attributes: map[string]schema.Attribute{
						"sensitivity": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("low", "medium", "high"),
							},
							MarkdownDescription: "Latency sensitivity",
							Required:            true,
						},
						"threshold": schema.Int64Attribute{
							Validators: []validator.Int64{
								int64validator.Between(10, 3000),
							},
							MarkdownDescription: "Latency threshold (ms)",
							Required:            true,
						},
					},
				},
				"pkt_loss": schema.SingleNestedAttribute{
					MarkdownDescription: "Pkt loss",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"sensitivity": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("low", "medium", "high"),
							},
							MarkdownDescription: "Packet loss sensitivity",
							Required:            true,
						},
						"threshold": schema.Int64Attribute{
							Validators: []validator.Int64{
								int64validator.Between(1, 100),
							},
							MarkdownDescription: "Packet loss threshold (percentage)",
							Required:            true,
						},
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

// SdwanPathQualityProfilesDataSourceSchema defines the schema for SdwanPathQualityProfiles data source
var SdwanPathQualityProfilesDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "SdwanPathQualityProfile data source",
	Attributes: map[string]dsschema.Attribute{
		"device": dsschema.StringAttribute{
			MarkdownDescription: "The device in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Computed:            true,
		},
		"folder": dsschema.StringAttribute{
			MarkdownDescription: "The folder in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Computed:            true,
		},
		"id": dsschema.StringAttribute{
			MarkdownDescription: "UUID of the resource",
			Required:            true,
		},
		"metric": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Metric",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"jitter": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Jitter",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"sensitivity": dsschema.StringAttribute{
							MarkdownDescription: "Jitter sensitivity",
							Computed:            true,
						},
						"threshold": dsschema.Int64Attribute{
							MarkdownDescription: "Jitter threshold (ms)",
							Computed:            true,
						},
					},
				},
				"latency": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Latency",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"sensitivity": dsschema.StringAttribute{
							MarkdownDescription: "Latency sensitivity",
							Computed:            true,
						},
						"threshold": dsschema.Int64Attribute{
							MarkdownDescription: "Latency threshold (ms)",
							Computed:            true,
						},
					},
				},
				"pkt_loss": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Pkt loss",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"sensitivity": dsschema.StringAttribute{
							MarkdownDescription: "Packet loss sensitivity",
							Computed:            true,
						},
						"threshold": dsschema.Int64Attribute{
							MarkdownDescription: "Packet loss threshold (percentage)",
							Computed:            true,
						},
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
			MarkdownDescription: "The snippet in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
	},
}

// SdwanPathQualityProfilesListModel represents the data model for a list data source.
type SdwanPathQualityProfilesListModel struct {
	Tfid    types.String               `tfsdk:"tfid"`
	Data    []SdwanPathQualityProfiles `tfsdk:"data"`
	Limit   types.Int64                `tfsdk:"limit"`
	Offset  types.Int64                `tfsdk:"offset"`
	Name    types.String               `tfsdk:"name"`
	Total   types.Int64                `tfsdk:"total"`
	Folder  types.String               `tfsdk:"folder"`
	Snippet types.String               `tfsdk:"snippet"`
	Device  types.String               `tfsdk:"device"`
}

// SdwanPathQualityProfilesListDataSourceSchema defines the schema for a list data source.
var SdwanPathQualityProfilesListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: SdwanPathQualityProfilesDataSourceSchema.Attributes,
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
