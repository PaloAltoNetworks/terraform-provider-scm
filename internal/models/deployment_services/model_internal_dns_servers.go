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

// InternalDnsServers represents the Terraform model for InternalDnsServers
type InternalDnsServers struct {
	Tfid       types.String          `tfsdk:"tfid"`
	DomainName basetypes.ListValue   `tfsdk:"domain_name"`
	Id         basetypes.StringValue `tfsdk:"id"`
	Name       basetypes.StringValue `tfsdk:"name"`
	Primary    basetypes.StringValue `tfsdk:"primary"`
	Secondary  basetypes.StringValue `tfsdk:"secondary"`
}

// AttrTypes defines the attribute types for the InternalDnsServers model.
func (o InternalDnsServers) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":        basetypes.StringType{},
		"domain_name": basetypes.ListType{ElemType: basetypes.StringType{}},
		"id":          basetypes.StringType{},
		"name":        basetypes.StringType{},
		"primary":     basetypes.StringType{},
		"secondary":   basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of InternalDnsServers objects.
func (o InternalDnsServers) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// InternalDnsServersResourceSchema defines the schema for InternalDnsServers resource
var InternalDnsServersResourceSchema = schema.Schema{
	MarkdownDescription: "Manages SCM InternalDnsServers objects",
	Attributes: map[string]schema.Attribute{
		"domain_name": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "The DNS domain name(s)",
			Required:            true,
		},
		"id": schema.StringAttribute{
			MarkdownDescription: "The UUID of the internet DNS server resource",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"name": schema.StringAttribute{
			MarkdownDescription: "The name of the internet DNS server resource",
			Required:            true,
		},
		"primary": schema.StringAttribute{
			MarkdownDescription: "The IP address of the primary DNS server",
			Required:            true,
		},
		"secondary": schema.StringAttribute{
			MarkdownDescription: "The IP address of the secondary DNS server",
			Optional:            true,
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

// InternalDnsServersDataSourceSchema defines the schema for InternalDnsServers data source
var InternalDnsServersDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "InternalDnsServers data source",
	Attributes: map[string]dsschema.Attribute{
		"domain_name": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "The DNS domain name(s)",
			Computed:            true,
		},
		"id": dsschema.StringAttribute{
			MarkdownDescription: "The UUID of the internet DNS server resource",
			Required:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "The name of the internet DNS server resource",
			Optional:            true,
			Computed:            true,
		},
		"primary": dsschema.StringAttribute{
			MarkdownDescription: "The IP address of the primary DNS server",
			Computed:            true,
		},
		"secondary": dsschema.StringAttribute{
			MarkdownDescription: "The IP address of the secondary DNS server",
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
	},
}

// InternalDnsServersListModel represents the data model for a list data source.
type InternalDnsServersListModel struct {
	Tfid    types.String         `tfsdk:"tfid"`
	Data    []InternalDnsServers `tfsdk:"data"`
	Limit   types.Int64          `tfsdk:"limit"`
	Offset  types.Int64          `tfsdk:"offset"`
	Name    types.String         `tfsdk:"name"`
	Total   types.Int64          `tfsdk:"total"`
	Folder  types.String         `tfsdk:"folder"`
	Snippet types.String         `tfsdk:"snippet"`
	Device  types.String         `tfsdk:"device"`
}

// InternalDnsServersListDataSourceSchema defines the schema for a list data source.
var InternalDnsServersListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: InternalDnsServersDataSourceSchema.Attributes,
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
