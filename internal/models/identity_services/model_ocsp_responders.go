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

// Package: identity_services
// This file contains models for the identity_services SDK package

// OcspResponders represents the Terraform model for OcspResponders
type OcspResponders struct {
	Tfid     types.String          `tfsdk:"tfid"`
	Device   basetypes.StringValue `tfsdk:"device"`
	Folder   basetypes.StringValue `tfsdk:"folder"`
	HostName basetypes.StringValue `tfsdk:"host_name"`
	Id       basetypes.StringValue `tfsdk:"id"`
	Name     basetypes.StringValue `tfsdk:"name"`
	Snippet  basetypes.StringValue `tfsdk:"snippet"`
}

// AttrTypes defines the attribute types for the OcspResponders model.
func (o OcspResponders) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":      basetypes.StringType{},
		"device":    basetypes.StringType{},
		"folder":    basetypes.StringType{},
		"host_name": basetypes.StringType{},
		"id":        basetypes.StringType{},
		"name":      basetypes.StringType{},
		"snippet":   basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of OcspResponders objects.
func (o OcspResponders) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// OcspRespondersResourceSchema defines the schema for OcspResponders resource
var OcspRespondersResourceSchema = schema.Schema{
	MarkdownDescription: "OcspResponder resource",
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
		"host_name": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(255),
				stringvalidator.LengthAtLeast(1),
			},
			MarkdownDescription: "The hostname or IP address of the OCSP server",
			Required:            true,
		},
		"id": schema.StringAttribute{
			MarkdownDescription: "The UUID of the OCSP responder profile",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"name": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(63),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9._-]+$"), "pattern must match "+"^[a-zA-Z0-9._-]+$"),
			},
			MarkdownDescription: "The name of the OCSP responder profile",
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
	},
}

// OcspRespondersDataSourceSchema defines the schema for OcspResponders data source
var OcspRespondersDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "OcspResponder data source",
	Attributes: map[string]dsschema.Attribute{
		"device": dsschema.StringAttribute{
			MarkdownDescription: "The device in which the resource is defined",
			Computed:            true,
		},
		"folder": dsschema.StringAttribute{
			MarkdownDescription: "The folder in which the resource is defined",
			Computed:            true,
		},
		"host_name": dsschema.StringAttribute{
			MarkdownDescription: "The hostname or IP address of the OCSP server",
			Computed:            true,
		},
		"id": dsschema.StringAttribute{
			MarkdownDescription: "The UUID of the OCSP responder profile",
			Required:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "The name of the OCSP responder profile",
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
	},
}

// OcspRespondersListModel represents the data model for a list data source.
type OcspRespondersListModel struct {
	Tfid    types.String     `tfsdk:"tfid"`
	Data    []OcspResponders `tfsdk:"data"`
	Limit   types.Int64      `tfsdk:"limit"`
	Offset  types.Int64      `tfsdk:"offset"`
	Name    types.String     `tfsdk:"name"`
	Total   types.Int64      `tfsdk:"total"`
	Folder  types.String     `tfsdk:"folder"`
	Snippet types.String     `tfsdk:"snippet"`
	Device  types.String     `tfsdk:"device"`
}

// OcspRespondersListDataSourceSchema defines the schema for a list data source.
var OcspRespondersListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: OcspRespondersDataSourceSchema.Attributes,
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
