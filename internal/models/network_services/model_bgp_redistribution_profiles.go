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

// BgpRedistributionProfiles represents the Terraform model for BgpRedistributionProfiles
type BgpRedistributionProfiles struct {
	Tfid    types.String          `tfsdk:"tfid"`
	Device  basetypes.StringValue `tfsdk:"device"`
	Folder  basetypes.StringValue `tfsdk:"folder"`
	Id      basetypes.StringValue `tfsdk:"id"`
	Ipv4    basetypes.ObjectValue `tfsdk:"ipv4"`
	Name    basetypes.StringValue `tfsdk:"name"`
	Snippet basetypes.StringValue `tfsdk:"snippet"`
}

// BgpRedistributionProfilesIpv4 represents a nested structure within the BgpRedistributionProfiles model
type BgpRedistributionProfilesIpv4 struct {
	Unicast basetypes.ObjectValue `tfsdk:"unicast"`
}

// BgpRedistributionProfilesIpv4Unicast represents a nested structure within the BgpRedistributionProfiles model
type BgpRedistributionProfilesIpv4Unicast struct {
	Connected basetypes.ObjectValue `tfsdk:"connected"`
	Ospf      basetypes.ObjectValue `tfsdk:"ospf"`
	Static    basetypes.ObjectValue `tfsdk:"static"`
}

// BgpRedistributionProfilesIpv4UnicastConnected represents a nested structure within the BgpRedistributionProfiles model
type BgpRedistributionProfilesIpv4UnicastConnected struct {
	Enable   basetypes.BoolValue   `tfsdk:"enable"`
	Metric   basetypes.Int64Value  `tfsdk:"metric"`
	RouteMap basetypes.StringValue `tfsdk:"route_map"`
}

// BgpRedistributionProfilesIpv4UnicastOspf represents a nested structure within the BgpRedistributionProfiles model
type BgpRedistributionProfilesIpv4UnicastOspf struct {
	Enable   basetypes.BoolValue   `tfsdk:"enable"`
	Metric   basetypes.Int64Value  `tfsdk:"metric"`
	RouteMap basetypes.StringValue `tfsdk:"route_map"`
}

// BgpRedistributionProfilesIpv4UnicastStatic represents a nested structure within the BgpRedistributionProfiles model
type BgpRedistributionProfilesIpv4UnicastStatic struct {
	Enable   basetypes.BoolValue   `tfsdk:"enable"`
	Metric   basetypes.Int64Value  `tfsdk:"metric"`
	RouteMap basetypes.StringValue `tfsdk:"route_map"`
}

// AttrTypes defines the attribute types for the BgpRedistributionProfiles model.
func (o BgpRedistributionProfiles) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":   basetypes.StringType{},
		"device": basetypes.StringType{},
		"folder": basetypes.StringType{},
		"id":     basetypes.StringType{},
		"ipv4": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"unicast": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"connected": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"enable":    basetypes.BoolType{},
								"metric":    basetypes.Int64Type{},
								"route_map": basetypes.StringType{},
							},
						},
						"ospf": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"enable":    basetypes.BoolType{},
								"metric":    basetypes.Int64Type{},
								"route_map": basetypes.StringType{},
							},
						},
						"static": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"enable":    basetypes.BoolType{},
								"metric":    basetypes.Int64Type{},
								"route_map": basetypes.StringType{},
							},
						},
					},
				},
			},
		},
		"name":    basetypes.StringType{},
		"snippet": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of BgpRedistributionProfiles objects.
func (o BgpRedistributionProfiles) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpRedistributionProfilesIpv4 model.
func (o BgpRedistributionProfilesIpv4) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"unicast": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"connected": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"enable":    basetypes.BoolType{},
						"metric":    basetypes.Int64Type{},
						"route_map": basetypes.StringType{},
					},
				},
				"ospf": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"enable":    basetypes.BoolType{},
						"metric":    basetypes.Int64Type{},
						"route_map": basetypes.StringType{},
					},
				},
				"static": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"enable":    basetypes.BoolType{},
						"metric":    basetypes.Int64Type{},
						"route_map": basetypes.StringType{},
					},
				},
			},
		},
	}
}

// AttrType returns the attribute type for a list of BgpRedistributionProfilesIpv4 objects.
func (o BgpRedistributionProfilesIpv4) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpRedistributionProfilesIpv4Unicast model.
func (o BgpRedistributionProfilesIpv4Unicast) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"connected": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"enable":    basetypes.BoolType{},
				"metric":    basetypes.Int64Type{},
				"route_map": basetypes.StringType{},
			},
		},
		"ospf": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"enable":    basetypes.BoolType{},
				"metric":    basetypes.Int64Type{},
				"route_map": basetypes.StringType{},
			},
		},
		"static": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"enable":    basetypes.BoolType{},
				"metric":    basetypes.Int64Type{},
				"route_map": basetypes.StringType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of BgpRedistributionProfilesIpv4Unicast objects.
func (o BgpRedistributionProfilesIpv4Unicast) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpRedistributionProfilesIpv4UnicastConnected model.
func (o BgpRedistributionProfilesIpv4UnicastConnected) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"enable":    basetypes.BoolType{},
		"metric":    basetypes.Int64Type{},
		"route_map": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of BgpRedistributionProfilesIpv4UnicastConnected objects.
func (o BgpRedistributionProfilesIpv4UnicastConnected) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpRedistributionProfilesIpv4UnicastOspf model.
func (o BgpRedistributionProfilesIpv4UnicastOspf) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"enable":    basetypes.BoolType{},
		"metric":    basetypes.Int64Type{},
		"route_map": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of BgpRedistributionProfilesIpv4UnicastOspf objects.
func (o BgpRedistributionProfilesIpv4UnicastOspf) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpRedistributionProfilesIpv4UnicastStatic model.
func (o BgpRedistributionProfilesIpv4UnicastStatic) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"enable":    basetypes.BoolType{},
		"metric":    basetypes.Int64Type{},
		"route_map": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of BgpRedistributionProfilesIpv4UnicastStatic objects.
func (o BgpRedistributionProfilesIpv4UnicastStatic) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// BgpRedistributionProfilesResourceSchema defines the schema for BgpRedistributionProfiles resource
var BgpRedistributionProfilesResourceSchema = schema.Schema{
	MarkdownDescription: "BgpRedistributionProfile resource",
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
		"ipv4": schema.SingleNestedAttribute{
			MarkdownDescription: "Ipv4",
			Required:            true,
			Attributes: map[string]schema.Attribute{
				"unicast": schema.SingleNestedAttribute{
					MarkdownDescription: "Unicast",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"connected": schema.SingleNestedAttribute{
							MarkdownDescription: "Connected",
							Optional:            true,
							Attributes: map[string]schema.Attribute{
								"enable": schema.BoolAttribute{
									MarkdownDescription: "Enable connected route redistribution?",
									Optional:            true,
								},
								"metric": schema.Int64Attribute{
									Validators: []validator.Int64{
										int64validator.Between(1, 65535),
									},
									MarkdownDescription: "Route metric",
									Optional:            true,
								},
								"route_map": schema.StringAttribute{
									MarkdownDescription: "Route map",
									Optional:            true,
								},
							},
						},
						"ospf": schema.SingleNestedAttribute{
							MarkdownDescription: "Ospf",
							Optional:            true,
							Attributes: map[string]schema.Attribute{
								"enable": schema.BoolAttribute{
									MarkdownDescription: "Enable OSPF route redistribution?",
									Optional:            true,
								},
								"metric": schema.Int64Attribute{
									Validators: []validator.Int64{
										int64validator.Between(1, 65535),
									},
									MarkdownDescription: "Route metric",
									Optional:            true,
								},
								"route_map": schema.StringAttribute{
									MarkdownDescription: "Route map",
									Optional:            true,
								},
							},
						},
						"static": schema.SingleNestedAttribute{
							MarkdownDescription: "Static",
							Optional:            true,
							Attributes: map[string]schema.Attribute{
								"enable": schema.BoolAttribute{
									MarkdownDescription: "Enable static route redistribution?",
									Optional:            true,
								},
								"metric": schema.Int64Attribute{
									Validators: []validator.Int64{
										int64validator.Between(1, 65535),
									},
									MarkdownDescription: "Route metric",
									Optional:            true,
								},
								"route_map": schema.StringAttribute{
									MarkdownDescription: "Route map",
									Optional:            true,
								},
							},
						},
					},
				},
			},
		},
		"name": schema.StringAttribute{
			MarkdownDescription: "Name",
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

// BgpRedistributionProfilesDataSourceSchema defines the schema for BgpRedistributionProfiles data source
var BgpRedistributionProfilesDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "BgpRedistributionProfile data source",
	Attributes: map[string]dsschema.Attribute{
		"device": dsschema.StringAttribute{
			MarkdownDescription: "The device in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			Computed:            true,
		},
		"folder": dsschema.StringAttribute{
			MarkdownDescription: "The folder in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			Computed:            true,
		},
		"id": dsschema.StringAttribute{
			MarkdownDescription: "UUID of the resource",
			Required:            true,
		},
		"ipv4": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Ipv4",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"unicast": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Unicast",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"connected": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Connected",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"enable": dsschema.BoolAttribute{
									MarkdownDescription: "Enable connected route redistribution?",
									Computed:            true,
								},
								"metric": dsschema.Int64Attribute{
									MarkdownDescription: "Route metric",
									Computed:            true,
								},
								"route_map": dsschema.StringAttribute{
									MarkdownDescription: "Route map",
									Computed:            true,
								},
							},
						},
						"ospf": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Ospf",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"enable": dsschema.BoolAttribute{
									MarkdownDescription: "Enable OSPF route redistribution?",
									Computed:            true,
								},
								"metric": dsschema.Int64Attribute{
									MarkdownDescription: "Route metric",
									Computed:            true,
								},
								"route_map": dsschema.StringAttribute{
									MarkdownDescription: "Route map",
									Computed:            true,
								},
							},
						},
						"static": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Static",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"enable": dsschema.BoolAttribute{
									MarkdownDescription: "Enable static route redistribution?",
									Computed:            true,
								},
								"metric": dsschema.Int64Attribute{
									MarkdownDescription: "Route metric",
									Computed:            true,
								},
								"route_map": dsschema.StringAttribute{
									MarkdownDescription: "Route map",
									Computed:            true,
								},
							},
						},
					},
				},
			},
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "Name",
			Optional:            true,
			Computed:            true,
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

// BgpRedistributionProfilesListModel represents the data model for a list data source.
type BgpRedistributionProfilesListModel struct {
	Tfid    types.String                `tfsdk:"tfid"`
	Data    []BgpRedistributionProfiles `tfsdk:"data"`
	Limit   types.Int64                 `tfsdk:"limit"`
	Offset  types.Int64                 `tfsdk:"offset"`
	Name    types.String                `tfsdk:"name"`
	Total   types.Int64                 `tfsdk:"total"`
	Folder  types.String                `tfsdk:"folder"`
	Snippet types.String                `tfsdk:"snippet"`
	Device  types.String                `tfsdk:"device"`
}

// BgpRedistributionProfilesListDataSourceSchema defines the schema for a list data source.
var BgpRedistributionProfilesListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: BgpRedistributionProfilesDataSourceSchema.Attributes,
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
