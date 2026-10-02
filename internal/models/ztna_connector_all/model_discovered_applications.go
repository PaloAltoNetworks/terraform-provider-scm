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

// Package: ztna_connector_all
// This file contains models for the ztna_connector_all SDK package

// DiscoveredApplications represents the Terraform model for DiscoveredApplications
type DiscoveredApplications struct {
	Tfid         types.String           `tfsdk:"tfid"`
	Applications basetypes.ListValue    `tfsdk:"applications"`
	CieTenantId  basetypes.StringValue  `tfsdk:"cie_tenant_id"`
	Count        basetypes.Float64Value `tfsdk:"count"`
	TenantId     basetypes.StringValue  `tfsdk:"tenant_id"`
}

// DiscoveredApplicationsApplicationsInner represents a nested structure within the DiscoveredApplications model
type DiscoveredApplicationsApplicationsInner struct {
	AppSpec  basetypes.ListValue   `tfsdk:"app_spec"`
	Fqdn     basetypes.StringValue `tfsdk:"fqdn"`
	Id       basetypes.StringValue `tfsdk:"id"`
	Name     basetypes.StringValue `tfsdk:"name"`
	Port     basetypes.StringValue `tfsdk:"port"`
	Protocol basetypes.StringValue `tfsdk:"protocol"`
	Provider basetypes.StringValue `tfsdk:"provider"`
}

// DiscoveredApplicationsApplicationsInnerAppSpecInner represents a nested structure within the DiscoveredApplications model
type DiscoveredApplicationsApplicationsInnerAppSpecInner struct {
	Fqdn     basetypes.StringValue `tfsdk:"fqdn"`
	Port     basetypes.StringValue `tfsdk:"port"`
	Protocol basetypes.StringValue `tfsdk:"protocol"`
}

// AttrTypes defines the attribute types for the DiscoveredApplications model.
func (o DiscoveredApplications) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid": basetypes.StringType{},
		"applications": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"app_spec": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"fqdn":     basetypes.StringType{},
						"port":     basetypes.StringType{},
						"protocol": basetypes.StringType{},
					},
				}},
				"fqdn":     basetypes.StringType{},
				"id":       basetypes.StringType{},
				"name":     basetypes.StringType{},
				"port":     basetypes.StringType{},
				"protocol": basetypes.StringType{},
				"provider": basetypes.StringType{},
			},
		}},
		"cie_tenant_id": basetypes.StringType{},
		"count":         basetypes.Float64Type{},
		"tenant_id":     basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of DiscoveredApplications objects.
func (o DiscoveredApplications) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the DiscoveredApplicationsApplicationsInner model.
func (o DiscoveredApplicationsApplicationsInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"app_spec": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"fqdn":     basetypes.StringType{},
				"port":     basetypes.StringType{},
				"protocol": basetypes.StringType{},
			},
		}},
		"fqdn":     basetypes.StringType{},
		"id":       basetypes.StringType{},
		"name":     basetypes.StringType{},
		"port":     basetypes.StringType{},
		"protocol": basetypes.StringType{},
		"provider": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of DiscoveredApplicationsApplicationsInner objects.
func (o DiscoveredApplicationsApplicationsInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the DiscoveredApplicationsApplicationsInnerAppSpecInner model.
func (o DiscoveredApplicationsApplicationsInnerAppSpecInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"fqdn":     basetypes.StringType{},
		"port":     basetypes.StringType{},
		"protocol": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of DiscoveredApplicationsApplicationsInnerAppSpecInner objects.
func (o DiscoveredApplicationsApplicationsInnerAppSpecInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// DiscoveredApplicationsResourceSchema defines the schema for DiscoveredApplications resource
var DiscoveredApplicationsResourceSchema = schema.Schema{
	MarkdownDescription: "DiscoveredApplication resource",
	Attributes: map[string]schema.Attribute{
		"applications": schema.ListNestedAttribute{
			MarkdownDescription: "Applications",
			Optional:            true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"app_spec": schema.ListNestedAttribute{
						MarkdownDescription: "App spec",
						Optional:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"fqdn": schema.StringAttribute{
									MarkdownDescription: "Fqdn",
									Optional:            true,
									Computed:            true,
								},
								"port": schema.StringAttribute{
									MarkdownDescription: "Port",
									Optional:            true,
									Computed:            true,
								},
								"protocol": schema.StringAttribute{
									MarkdownDescription: "Protocol",
									Optional:            true,
									Computed:            true,
								},
							},
						},
					},
					"fqdn": schema.StringAttribute{
						MarkdownDescription: "Fqdn",
						Optional:            true,
						Computed:            true,
					},
					"id": schema.StringAttribute{
						MarkdownDescription: "Id",
						Optional:            true,
						Computed:            true,
					},
					"name": schema.StringAttribute{
						MarkdownDescription: "Name",
						Optional:            true,
						Computed:            true,
					},
					"port": schema.StringAttribute{
						MarkdownDescription: "Port",
						Optional:            true,
						Computed:            true,
					},
					"protocol": schema.StringAttribute{
						MarkdownDescription: "Protocol",
						Optional:            true,
						Computed:            true,
					},
					"provider": schema.StringAttribute{
						MarkdownDescription: "Provider",
						Optional:            true,
						Computed:            true,
					},
				},
			},
		},
		"cie_tenant_id": schema.StringAttribute{
			MarkdownDescription: "Cie tenant id",
			Optional:            true,
			Computed:            true,
		},
		"count": schema.Float64Attribute{
			MarkdownDescription: "Count",
			Optional:            true,
			Computed:            true,
		},
		"tenant_id": schema.StringAttribute{
			MarkdownDescription: "Tenant id",
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

// DiscoveredApplicationsDataSourceSchema defines the schema for DiscoveredApplications data source
var DiscoveredApplicationsDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "DiscoveredApplication data source",
	Attributes: map[string]dsschema.Attribute{
		"applications": dsschema.ListNestedAttribute{
			MarkdownDescription: "Applications",
			Computed:            true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: map[string]dsschema.Attribute{
					"app_spec": dsschema.ListNestedAttribute{
						MarkdownDescription: "App spec",
						Computed:            true,
						NestedObject: dsschema.NestedAttributeObject{
							Attributes: map[string]dsschema.Attribute{
								"fqdn": dsschema.StringAttribute{
									MarkdownDescription: "Fqdn",
									Computed:            true,
								},
								"port": dsschema.StringAttribute{
									MarkdownDescription: "Port",
									Computed:            true,
								},
								"protocol": dsschema.StringAttribute{
									MarkdownDescription: "Protocol",
									Computed:            true,
								},
							},
						},
					},
					"fqdn": dsschema.StringAttribute{
						MarkdownDescription: "Fqdn",
						Computed:            true,
					},
					"id": dsschema.StringAttribute{
						MarkdownDescription: "Id",
						Computed:            true,
					},
					"name": dsschema.StringAttribute{
						MarkdownDescription: "Name",
						Computed:            true,
					},
					"port": dsschema.StringAttribute{
						MarkdownDescription: "Port",
						Computed:            true,
					},
					"protocol": dsschema.StringAttribute{
						MarkdownDescription: "Protocol",
						Computed:            true,
					},
					"provider": dsschema.StringAttribute{
						MarkdownDescription: "Provider",
						Computed:            true,
					},
				},
			},
		},
		"cie_tenant_id": dsschema.StringAttribute{
			MarkdownDescription: "Cie tenant id",
			Computed:            true,
		},
		"count": dsschema.Float64Attribute{
			MarkdownDescription: "Count",
			Computed:            true,
		},
		"tenant_id": dsschema.StringAttribute{
			MarkdownDescription: "Tenant id",
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
	},
}

// DiscoveredApplicationsListModel represents the data model for a list data source.
type DiscoveredApplicationsListModel struct {
	Tfid    types.String             `tfsdk:"tfid"`
	Data    []DiscoveredApplications `tfsdk:"data"`
	Limit   types.Int64              `tfsdk:"limit"`
	Offset  types.Int64              `tfsdk:"offset"`
	Name    types.String             `tfsdk:"name"`
	Total   types.Int64              `tfsdk:"total"`
	Folder  types.String             `tfsdk:"folder"`
	Snippet types.String             `tfsdk:"snippet"`
	Device  types.String             `tfsdk:"device"`
	Sort    basetypes.StringValue    `tfsdk:"sort"`
	Search  basetypes.StringValue    `tfsdk:"search"`
	Filters basetypes.StringValue    `tfsdk:"filters"`
}

// DiscoveredApplicationsListDataSourceSchema defines the schema for a list data source.
var DiscoveredApplicationsListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: DiscoveredApplicationsDataSourceSchema.Attributes,
			},
		},
		"limit":   dsschema.Int64Attribute{Description: "The max number of items to return. Default: 200.", Optional: true},
		"offset":  dsschema.Int64Attribute{Description: "The offset of the first item to return.", Optional: true},
		"name":    dsschema.StringAttribute{Description: "The name of the item.", Optional: true},
		"total":   dsschema.Int64Attribute{Description: "The total number of items.", Computed: true},
		"folder":  dsschema.StringAttribute{Description: "The folder of the item. Default: Shared.", Optional: true},
		"snippet": dsschema.StringAttribute{Description: "The snippet of the item.", Optional: true},
		"device":  dsschema.StringAttribute{Description: "The device of the item.", Optional: true},
		"sort": dsschema.StringAttribute{
			Description: "List of fields from item response to sort by.",
			Optional:    true,
		},
		"search": dsschema.StringAttribute{
			Description: "String to filter list results by. Is searched over multiple fields in each object. Multiple searches can be specified.",
			Optional:    true,
		},
		"filters": dsschema.StringAttribute{
			Description: "String to filter list results by searching one specified field of an object. Multiple filters can be specified.",
			Optional:    true,
		},
	},
}
