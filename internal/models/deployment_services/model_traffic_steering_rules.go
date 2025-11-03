package models

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	dsschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Package: deployment_services
// This file contains models for the deployment_services SDK package

// TrafficSteeringRules represents the Terraform model for TrafficSteeringRules
type TrafficSteeringRules struct {
	Tfid        types.String          `tfsdk:"tfid"`
	Action      basetypes.ObjectValue `tfsdk:"action"`
	Category    basetypes.ListValue   `tfsdk:"category"`
	Destination basetypes.ListValue   `tfsdk:"destination"`
	Folder      basetypes.StringValue `tfsdk:"folder"`
	Id          basetypes.StringValue `tfsdk:"id"`
	Name        basetypes.StringValue `tfsdk:"name"`
	Service     basetypes.ListValue   `tfsdk:"service"`
	Source      basetypes.ListValue   `tfsdk:"source"`
	SourceUser  basetypes.ListValue   `tfsdk:"source_user"`
}

// TrafficSteeringRulesAction represents a nested structure within the TrafficSteeringRules model
type TrafficSteeringRulesAction struct {
	Forward basetypes.ObjectValue `tfsdk:"forward"`
}

// TrafficSteeringRulesActionForward represents a nested structure within the TrafficSteeringRules model
type TrafficSteeringRulesActionForward struct {
	Forward basetypes.ObjectValue `tfsdk:"forward"`
	NoPbf   basetypes.ObjectValue `tfsdk:"no_pbf"`
}

// TrafficSteeringRulesActionForwardForward represents a nested structure within the TrafficSteeringRules model
type TrafficSteeringRulesActionForwardForward struct {
	Target basetypes.StringValue `tfsdk:"target"`
}

// AttrTypes defines the attribute types for the TrafficSteeringRules model.
func (o TrafficSteeringRules) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid": basetypes.StringType{},
		"action": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"forward": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"forward": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"target": basetypes.StringType{},
							},
						},
						"no_pbf": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
					},
				},
			},
		},
		"category":    basetypes.ListType{ElemType: basetypes.StringType{}},
		"destination": basetypes.ListType{ElemType: basetypes.StringType{}},
		"folder":      basetypes.StringType{},
		"id":          basetypes.StringType{},
		"name":        basetypes.StringType{},
		"service":     basetypes.ListType{ElemType: basetypes.StringType{}},
		"source":      basetypes.ListType{ElemType: basetypes.StringType{}},
		"source_user": basetypes.ListType{ElemType: basetypes.StringType{}},
	}
}

// AttrType returns the attribute type for a list of TrafficSteeringRules objects.
func (o TrafficSteeringRules) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the TrafficSteeringRulesAction model.
func (o TrafficSteeringRulesAction) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"forward": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"forward": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"target": basetypes.StringType{},
					},
				},
				"no_pbf": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
			},
		},
	}
}

// AttrType returns the attribute type for a list of TrafficSteeringRulesAction objects.
func (o TrafficSteeringRulesAction) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the TrafficSteeringRulesActionForward model.
func (o TrafficSteeringRulesActionForward) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"forward": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"target": basetypes.StringType{},
			},
		},
		"no_pbf": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
	}
}

// AttrType returns the attribute type for a list of TrafficSteeringRulesActionForward objects.
func (o TrafficSteeringRulesActionForward) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the TrafficSteeringRulesActionForwardForward model.
func (o TrafficSteeringRulesActionForwardForward) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"target": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of TrafficSteeringRulesActionForwardForward objects.
func (o TrafficSteeringRulesActionForwardForward) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// TrafficSteeringRulesResourceSchema defines the schema for TrafficSteeringRules resource
var TrafficSteeringRulesResourceSchema = schema.Schema{
	MarkdownDescription: "TrafficSteeringRule resource",
	Attributes: map[string]schema.Attribute{
		"action": schema.SingleNestedAttribute{
			MarkdownDescription: "Action",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"forward": schema.SingleNestedAttribute{
					MarkdownDescription: "Forward",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"forward": schema.SingleNestedAttribute{
							MarkdownDescription: "Forward",
							Optional:            true,
							Attributes: map[string]schema.Attribute{
								"target": schema.StringAttribute{
									MarkdownDescription: "Target",
									Optional:            true,
								},
							},
						},
						"no_pbf": schema.StringAttribute{
							MarkdownDescription: "No pbf",
							Optional:            true,
						},
					},
				},
			},
		},
		"category": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Category",
			Optional:            true,
		},
		"destination": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Destination",
			Optional:            true,
			Computed:            true,
		},
		"folder": schema.StringAttribute{
			MarkdownDescription: "The folder containing the traffic steering rule",
			Required:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"id": schema.StringAttribute{
			MarkdownDescription: "The UUID of the traffic steering rule",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"name": schema.StringAttribute{
			MarkdownDescription: "Name",
			Required:            true,
		},
		"service": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Service",
			Required:            true,
		},
		"source": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Source",
			Required:            true,
		},
		"source_user": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Source user",
			Optional:            true,
			Computed:            true,
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

// TrafficSteeringRulesDataSourceSchema defines the schema for TrafficSteeringRules data source
var TrafficSteeringRulesDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "TrafficSteeringRule data source",
	Attributes: map[string]dsschema.Attribute{
		"action": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Action",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"forward": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Forward",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"forward": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Forward",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"target": dsschema.StringAttribute{
									MarkdownDescription: "Target",
									Computed:            true,
								},
							},
						},
						"no_pbf": dsschema.StringAttribute{
							MarkdownDescription: "No pbf",
							Computed:            true,
						},
					},
				},
			},
		},
		"category": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Category",
			Computed:            true,
		},
		"destination": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Destination",
			Computed:            true,
		},
		"folder": dsschema.StringAttribute{
			MarkdownDescription: "The folder containing the traffic steering rule",
			Computed:            true,
		},
		"id": dsschema.StringAttribute{
			MarkdownDescription: "The UUID of the traffic steering rule",
			Required:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "Name",
			Optional:            true,
			Computed:            true,
		},
		"service": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Service",
			Computed:            true,
		},
		"source": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Source",
			Computed:            true,
		},
		"source_user": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Source user",
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
	},
}

// TrafficSteeringRulesListModel represents the data model for a list data source.
type TrafficSteeringRulesListModel struct {
	Tfid    types.String           `tfsdk:"tfid"`
	Data    []TrafficSteeringRules `tfsdk:"data"`
	Limit   types.Int64            `tfsdk:"limit"`
	Offset  types.Int64            `tfsdk:"offset"`
	Name    types.String           `tfsdk:"name"`
	Total   types.Int64            `tfsdk:"total"`
	Folder  types.String           `tfsdk:"folder"`
	Snippet types.String           `tfsdk:"snippet"`
	Device  types.String           `tfsdk:"device"`
}

// TrafficSteeringRulesListDataSourceSchema defines the schema for a list data source.
var TrafficSteeringRulesListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: TrafficSteeringRulesDataSourceSchema.Attributes,
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
