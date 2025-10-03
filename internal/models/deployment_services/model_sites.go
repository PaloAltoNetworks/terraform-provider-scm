package models

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	dsschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Package: deployment_services
// This file contains models for the deployment_services SDK package

// Sites represents the Terraform model for Sites
type Sites struct {
	Tfid         types.String           `tfsdk:"tfid"`
	AddressLine1 basetypes.StringValue  `tfsdk:"address_line_1"`
	AddressLine2 basetypes.StringValue  `tfsdk:"address_line_2"`
	City         basetypes.StringValue  `tfsdk:"city"`
	Country      basetypes.StringValue  `tfsdk:"country"`
	Id           basetypes.StringValue  `tfsdk:"id"`
	Latitude     basetypes.Float64Value `tfsdk:"latitude"`
	Longitude    basetypes.Float64Value `tfsdk:"longitude"`
	Members      basetypes.ListValue    `tfsdk:"members"`
	Name         basetypes.StringValue  `tfsdk:"name"`
	Qos          basetypes.ObjectValue  `tfsdk:"qos"`
	State        basetypes.StringValue  `tfsdk:"state"`
	Type         basetypes.StringValue  `tfsdk:"type"`
	ZipCode      basetypes.StringValue  `tfsdk:"zip_code"`
}

// SitesMembersInner represents a nested structure within the Sites model
type SitesMembersInner struct {
	Id            basetypes.StringValue `tfsdk:"id"`
	Mode          basetypes.StringValue `tfsdk:"mode"`
	Name          basetypes.StringValue `tfsdk:"name"`
	RemoteNetwork basetypes.StringValue `tfsdk:"remote_network"`
}

// SitesQos represents a nested structure within the Sites model
type SitesQos struct {
	BackupCir basetypes.Float64Value `tfsdk:"backup_cir"`
	Cir       basetypes.Float64Value `tfsdk:"cir"`
	Profile   basetypes.StringValue  `tfsdk:"profile"`
}

// AttrTypes defines the attribute types for the Sites model.
func (o Sites) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":           basetypes.StringType{},
		"address_line_1": basetypes.StringType{},
		"address_line_2": basetypes.StringType{},
		"city":           basetypes.StringType{},
		"country":        basetypes.StringType{},
		"id":             basetypes.StringType{},
		"latitude":       basetypes.NumberType{},
		"longitude":      basetypes.NumberType{},
		"members": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"id":             basetypes.StringType{},
				"mode":           basetypes.StringType{},
				"name":           basetypes.StringType{},
				"remote_network": basetypes.StringType{},
			},
		}},
		"name": basetypes.StringType{},
		"qos": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"backup_cir": basetypes.NumberType{},
				"cir":        basetypes.NumberType{},
				"profile":    basetypes.StringType{},
			},
		},
		"state":    basetypes.StringType{},
		"type":     basetypes.StringType{},
		"zip_code": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of Sites objects.
func (o Sites) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the SitesMembersInner model.
func (o SitesMembersInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":             basetypes.StringType{},
		"mode":           basetypes.StringType{},
		"name":           basetypes.StringType{},
		"remote_network": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of SitesMembersInner objects.
func (o SitesMembersInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the SitesQos model.
func (o SitesQos) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"backup_cir": basetypes.NumberType{},
		"cir":        basetypes.NumberType{},
		"profile":    basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of SitesQos objects.
func (o SitesQos) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// SitesResourceSchema defines the schema for Sites resource
var SitesResourceSchema = schema.Schema{
	MarkdownDescription: "Manages SCM Sites objects",
	Attributes: map[string]schema.Attribute{
		"address_line_1": schema.StringAttribute{
			MarkdownDescription: "The address in which the site exists",
			Optional:            true,
		},
		"address_line_2": schema.StringAttribute{
			MarkdownDescription: "The address in which the site exists (continued)",
			Optional:            true,
		},
		"city": schema.StringAttribute{
			MarkdownDescription: "The city in which the site exists",
			Optional:            true,
		},
		"country": schema.StringAttribute{
			MarkdownDescription: "The country in which the site exists",
			Optional:            true,
		},
		"id": schema.StringAttribute{
			MarkdownDescription: "The UUID of the site",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"latitude": schema.Float64Attribute{
			MarkdownDescription: "The latitude coordinate for the site",
			Optional:            true,
		},
		"longitude": schema.Float64Attribute{
			MarkdownDescription: "The longitude coordinate for the site",
			Optional:            true,
		},
		"members": schema.ListNestedAttribute{
			MarkdownDescription: "Members",
			Optional:            true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						MarkdownDescription: "UUID of the remote network",
						Optional:            true,
					},
					"mode": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("active", "backup"),
						},
						MarkdownDescription: "The mode of the remote network",
						Required:            true,
					},
					"name": schema.StringAttribute{
						MarkdownDescription: "The member name",
						Required:            true,
					},
					"remote_network": schema.StringAttribute{
						MarkdownDescription: "The remote network name",
						Optional:            true,
					},
				},
			},
		},
		"name": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(63),
			},
			MarkdownDescription: "The name of the site",
			Required:            true,
		},
		"qos": schema.SingleNestedAttribute{
			MarkdownDescription: "Qos",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"backup_cir": schema.Float64Attribute{
					MarkdownDescription: "The backup CIR in Mbps. This is distributed equally for all tunnels in the site.",
					Optional:            true,
				},
				"cir": schema.Float64Attribute{
					MarkdownDescription: "The CIR in Mbps. This is distributed equally for all tunnels in the site.",
					Optional:            true,
				},
				"profile": schema.StringAttribute{
					MarkdownDescription: "The name of the site QoS profile",
					Optional:            true,
				},
			},
		},
		"state": schema.StringAttribute{
			MarkdownDescription: "The state in which the site exists",
			Optional:            true,
		},
		"tfid": schema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"type": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.OneOf("prisma-sdwan", "third-party-branch", "third-party-discovered"),
			},
			MarkdownDescription: "The site type",
			Required:            true,
		},
		"zip_code": schema.StringAttribute{
			MarkdownDescription: "The postal code in which the site exists",
			Optional:            true,
		},
	},
}

// SitesDataSourceSchema defines the schema for Sites data source
var SitesDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Sites data source",
	Attributes: map[string]dsschema.Attribute{
		"address_line_1": dsschema.StringAttribute{
			MarkdownDescription: "The address in which the site exists",
			Computed:            true,
		},
		"address_line_2": dsschema.StringAttribute{
			MarkdownDescription: "The address in which the site exists (continued)",
			Computed:            true,
		},
		"city": dsschema.StringAttribute{
			MarkdownDescription: "The city in which the site exists",
			Computed:            true,
		},
		"country": dsschema.StringAttribute{
			MarkdownDescription: "The country in which the site exists",
			Computed:            true,
		},
		"id": dsschema.StringAttribute{
			MarkdownDescription: "The UUID of the site",
			Required:            true,
		},
		"latitude": dsschema.Float64Attribute{
			MarkdownDescription: "The latitude coordinate for the site",
			Computed:            true,
		},
		"longitude": dsschema.Float64Attribute{
			MarkdownDescription: "The longitude coordinate for the site",
			Computed:            true,
		},
		"members": dsschema.ListNestedAttribute{
			MarkdownDescription: "Members",
			Computed:            true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: map[string]dsschema.Attribute{
					"id": dsschema.StringAttribute{
						MarkdownDescription: "UUID of the remote network",
						Computed:            true,
					},
					"mode": dsschema.StringAttribute{
						MarkdownDescription: "The mode of the remote network",
						Computed:            true,
					},
					"name": dsschema.StringAttribute{
						MarkdownDescription: "The member name",
						Computed:            true,
					},
					"remote_network": dsschema.StringAttribute{
						MarkdownDescription: "The remote network name",
						Computed:            true,
					},
				},
			},
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "The name of the site",
			Optional:            true,
			Computed:            true,
		},
		"qos": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Qos",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"backup_cir": dsschema.Float64Attribute{
					MarkdownDescription: "The backup CIR in Mbps. This is distributed equally for all tunnels in the site.",
					Computed:            true,
				},
				"cir": dsschema.Float64Attribute{
					MarkdownDescription: "The CIR in Mbps. This is distributed equally for all tunnels in the site.",
					Computed:            true,
				},
				"profile": dsschema.StringAttribute{
					MarkdownDescription: "The name of the site QoS profile",
					Computed:            true,
				},
			},
		},
		"state": dsschema.StringAttribute{
			MarkdownDescription: "The state in which the site exists",
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
		"type": dsschema.StringAttribute{
			MarkdownDescription: "The site type",
			Computed:            true,
		},
		"zip_code": dsschema.StringAttribute{
			MarkdownDescription: "The postal code in which the site exists",
			Computed:            true,
		},
	},
}

// SitesListModel represents the data model for a list data source.
type SitesListModel struct {
	Tfid    types.String `tfsdk:"tfid"`
	Data    []Sites      `tfsdk:"data"`
	Limit   types.Int64  `tfsdk:"limit"`
	Offset  types.Int64  `tfsdk:"offset"`
	Name    types.String `tfsdk:"name"`
	Total   types.Int64  `tfsdk:"total"`
	Folder  types.String `tfsdk:"folder"`
	Snippet types.String `tfsdk:"snippet"`
	Device  types.String `tfsdk:"device"`
}

// SitesListDataSourceSchema defines the schema for a list data source.
var SitesListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: SitesDataSourceSchema.Attributes,
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
