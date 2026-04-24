package models

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	dsschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/paloaltonetworks/terraform-provider-scm/internal/utils"
)

// Package: mobile_agent
// This file contains models for the mobile_agent SDK package

// ForwardingProfileSourceApplications represents the Terraform model for ForwardingProfileSourceApplications
type ForwardingProfileSourceApplications struct {
	Tfid         types.String          `tfsdk:"tfid"`
	Applications basetypes.ListValue   `tfsdk:"applications"`
	Description  basetypes.StringValue `tfsdk:"description"`
	Id           basetypes.StringValue `tfsdk:"id"`
	Name         basetypes.StringValue `tfsdk:"name"`
	Folder       basetypes.StringValue `tfsdk:"folder"`
}

// AttrTypes defines the attribute types for the ForwardingProfileSourceApplications model.
func (o ForwardingProfileSourceApplications) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":         basetypes.StringType{},
		"applications": basetypes.ListType{ElemType: basetypes.StringType{}},
		"description":  basetypes.StringType{},
		"id":           basetypes.StringType{},
		"name":         basetypes.StringType{},
		"folder":       basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of ForwardingProfileSourceApplications objects.
func (o ForwardingProfileSourceApplications) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// ForwardingProfileSourceApplicationsResourceSchema defines the schema for ForwardingProfileSourceApplications resource
var ForwardingProfileSourceApplicationsResourceSchema = schema.Schema{
	MarkdownDescription: "ForwardingProfileSourceApplication resource",
	Attributes: map[string]schema.Attribute{
		"applications": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "List of application names to be included in this source application profile",
			Validators: []validator.List{
				listvalidator.ValueStringsAre(stringvalidator.LengthAtMost(511)),
			},
			Required: true,
		},
		"description": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(1023),
			},
			MarkdownDescription: "fowarding profile source application description",
			Optional:            true,
		},
		"folder": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.OneOf("Mobile Users"),
				utils.FolderValidator(),
			},
			MarkdownDescription: "The folder in which the resource is defined\n",
			Optional:            true,
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			},
		},
		"id": schema.StringAttribute{
			MarkdownDescription: "The id of the source application",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"name": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(64),
				stringvalidator.RegexMatches(regexp.MustCompile("^[0-9a-zA-Z._-]+$"), "pattern must match "+"^[0-9a-zA-Z._-]+$"),
			},
			MarkdownDescription: "The unique name identifying the source application. Must be alphanumeric with allowed characters [0-9a-zA-Z._-]",
			Required:            true,
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

// ForwardingProfileSourceApplicationsDataSourceSchema defines the schema for ForwardingProfileSourceApplications data source
var ForwardingProfileSourceApplicationsDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "ForwardingProfileSourceApplication data source",
	Attributes: map[string]dsschema.Attribute{
		"applications": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "List of application names to be included in this source application profile",
			Computed:            true,
		},
		"description": dsschema.StringAttribute{
			MarkdownDescription: "fowarding profile source application description",
			Computed:            true,
		},
		"folder": dsschema.StringAttribute{
			MarkdownDescription: "The folder in which the resource is defined\n",
			Optional:            true,
			Computed:            true,
		},
		"id": dsschema.StringAttribute{
			MarkdownDescription: "The id of the source application",
			Required:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "The unique name identifying the source application. Must be alphanumeric with allowed characters [0-9a-zA-Z._-]",
			Optional:            true,
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
	},
}

// ForwardingProfileSourceApplicationsListModel represents the data model for a list data source.
type ForwardingProfileSourceApplicationsListModel struct {
	Tfid    types.String                          `tfsdk:"tfid"`
	Data    []ForwardingProfileSourceApplications `tfsdk:"data"`
	Limit   types.Int64                           `tfsdk:"limit"`
	Offset  types.Int64                           `tfsdk:"offset"`
	Name    types.String                          `tfsdk:"name"`
	Total   types.Int64                           `tfsdk:"total"`
	Folder  types.String                          `tfsdk:"folder"`
	Snippet types.String                          `tfsdk:"snippet"`
	Device  types.String                          `tfsdk:"device"`
}

// ForwardingProfileSourceApplicationsListDataSourceSchema defines the schema for a list data source.
var ForwardingProfileSourceApplicationsListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: ForwardingProfileSourceApplicationsDataSourceSchema.Attributes,
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
