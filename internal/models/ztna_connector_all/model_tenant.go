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

// TenantStatus represents the Terraform model for TenantStatus
type TenantStatus struct {
	Tfid   types.String          `tfsdk:"tfid"`
	Status basetypes.StringValue `tfsdk:"status"`
}

// AttrTypes defines the attribute types for the TenantStatus model.
func (o TenantStatus) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":   basetypes.StringType{},
		"status": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of TenantStatus objects.
func (o TenantStatus) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// TenantStatusResourceSchema defines the schema for TenantStatus resource
var TenantStatusResourceSchema = schema.Schema{
	MarkdownDescription: "TenantStatus resource",
	Attributes: map[string]schema.Attribute{
		"status": schema.StringAttribute{
			MarkdownDescription: "Status",
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

// TenantStatusDataSourceSchema defines the schema for TenantStatus data source
var TenantStatusDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "TenantStatus data source",
	Attributes: map[string]dsschema.Attribute{
		"status": dsschema.StringAttribute{
			MarkdownDescription: "Status",
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
	},
}

// TenantStatusListModel represents the data model for a list data source.
type TenantStatusListModel struct {
	Tfid    types.String   `tfsdk:"tfid"`
	Data    []TenantStatus `tfsdk:"data"`
	Limit   types.Int64    `tfsdk:"limit"`
	Offset  types.Int64    `tfsdk:"offset"`
	Name    types.String   `tfsdk:"name"`
	Total   types.Int64    `tfsdk:"total"`
	Folder  types.String   `tfsdk:"folder"`
	Snippet types.String   `tfsdk:"snippet"`
	Device  types.String   `tfsdk:"device"`
}

// TenantStatusListDataSourceSchema defines the schema for a list data source.
var TenantStatusListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: TenantStatusDataSourceSchema.Attributes,
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
