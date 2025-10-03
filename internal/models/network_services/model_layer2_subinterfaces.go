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

// Layer2Subinterfaces represents the Terraform model for Layer2Subinterfaces
type Layer2Subinterfaces struct {
	Tfid            types.String           `tfsdk:"tfid"`
	Comment         basetypes.StringValue  `tfsdk:"comment"`
	Device          basetypes.StringValue  `tfsdk:"device"`
	Folder          basetypes.StringValue  `tfsdk:"folder"`
	Id              basetypes.StringValue  `tfsdk:"id"`
	Name            basetypes.StringValue  `tfsdk:"name"`
	ParentInterface basetypes.StringValue  `tfsdk:"parent_interface"`
	Snippet         basetypes.StringValue  `tfsdk:"snippet"`
	VlanTag         basetypes.Float64Value `tfsdk:"vlan_tag"`
}

// AttrTypes defines the attribute types for the Layer2Subinterfaces model.
func (o Layer2Subinterfaces) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":             basetypes.StringType{},
		"comment":          basetypes.StringType{},
		"device":           basetypes.StringType{},
		"folder":           basetypes.StringType{},
		"id":               basetypes.StringType{},
		"name":             basetypes.StringType{},
		"parent_interface": basetypes.StringType{},
		"snippet":          basetypes.StringType{},
		"vlan_tag":         basetypes.NumberType{},
	}
}

// AttrType returns the attribute type for a list of Layer2Subinterfaces objects.
func (o Layer2Subinterfaces) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// Layer2SubinterfacesResourceSchema defines the schema for Layer2Subinterfaces resource
var Layer2SubinterfacesResourceSchema = schema.Schema{
	MarkdownDescription: "Manages SCM Layer2Subinterfaces objects",
	Attributes: map[string]schema.Attribute{
		"comment": schema.StringAttribute{
			MarkdownDescription: "Description",
			Optional:            true,
		},
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
		"name": schema.StringAttribute{
			MarkdownDescription: "L2 sub-interface name",
			Required:            true,
		},
		"parent_interface": schema.StringAttribute{
			MarkdownDescription: "Parent interface",
			Optional:            true,
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
		"vlan_tag": schema.Float64Attribute{
			MarkdownDescription: "Vlan tag",
			Optional:            true,
		},
	},
}

// Layer2SubinterfacesDataSourceSchema defines the schema for Layer2Subinterfaces data source
var Layer2SubinterfacesDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Layer2Subinterfaces data source",
	Attributes: map[string]dsschema.Attribute{
		"comment": dsschema.StringAttribute{
			MarkdownDescription: "Description",
			Computed:            true,
		},
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
		"name": dsschema.StringAttribute{
			MarkdownDescription: "L2 sub-interface name",
			Optional:            true,
			Computed:            true,
		},
		"parent_interface": dsschema.StringAttribute{
			MarkdownDescription: "Parent interface",
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
		"vlan_tag": dsschema.Float64Attribute{
			MarkdownDescription: "Vlan tag",
			Computed:            true,
		},
	},
}

// Layer2SubinterfacesListModel represents the data model for a list data source.
type Layer2SubinterfacesListModel struct {
	Tfid    types.String          `tfsdk:"tfid"`
	Data    []Layer2Subinterfaces `tfsdk:"data"`
	Limit   types.Int64           `tfsdk:"limit"`
	Offset  types.Int64           `tfsdk:"offset"`
	Name    types.String          `tfsdk:"name"`
	Total   types.Int64           `tfsdk:"total"`
	Folder  types.String          `tfsdk:"folder"`
	Snippet types.String          `tfsdk:"snippet"`
	Device  types.String          `tfsdk:"device"`
}

// Layer2SubinterfacesListDataSourceSchema defines the schema for a list data source.
var Layer2SubinterfacesListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: Layer2SubinterfacesDataSourceSchema.Attributes,
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
