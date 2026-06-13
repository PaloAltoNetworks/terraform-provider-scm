package models

import (
	"regexp"

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

// ForwardingProfileUserLocations represents the Terraform model for ForwardingProfileUserLocations
type ForwardingProfileUserLocations struct {
	Tfid                  types.String          `tfsdk:"tfid"`
	Description           basetypes.StringValue `tfsdk:"description"`
	Id                    basetypes.StringValue `tfsdk:"id"`
	InternalHostDetection basetypes.ObjectValue `tfsdk:"internal_host_detection"`
	IpAddresses           basetypes.ListValue   `tfsdk:"ip_addresses"`
	Name                  basetypes.StringValue `tfsdk:"name"`
	Folder                basetypes.StringValue `tfsdk:"folder"`
}

// ForwardingProfileUserLocationsInternalHostDetection represents a nested structure within the ForwardingProfileUserLocations model
type ForwardingProfileUserLocationsInternalHostDetection struct {
	Fqdn      basetypes.StringValue `tfsdk:"fqdn"`
	IpAddress basetypes.StringValue `tfsdk:"ip_address"`
}

// AttrTypes defines the attribute types for the ForwardingProfileUserLocations model.
func (o ForwardingProfileUserLocations) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":        basetypes.StringType{},
		"description": basetypes.StringType{},
		"id":          basetypes.StringType{},
		"internal_host_detection": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"fqdn":       basetypes.StringType{},
				"ip_address": basetypes.StringType{},
			},
		},
		"ip_addresses": basetypes.ListType{ElemType: basetypes.StringType{}},
		"name":         basetypes.StringType{},
		"folder":       basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of ForwardingProfileUserLocations objects.
func (o ForwardingProfileUserLocations) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ForwardingProfileUserLocationsInternalHostDetection model.
func (o ForwardingProfileUserLocationsInternalHostDetection) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"fqdn":       basetypes.StringType{},
		"ip_address": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of ForwardingProfileUserLocationsInternalHostDetection objects.
func (o ForwardingProfileUserLocationsInternalHostDetection) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// ForwardingProfileUserLocationsResourceSchema defines the schema for ForwardingProfileUserLocations resource
var ForwardingProfileUserLocationsResourceSchema = schema.Schema{
	MarkdownDescription: "ForwardingProfileUserLocation resource",
	Attributes: map[string]schema.Attribute{
		"description": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(1023),
			},
			MarkdownDescription: "Description of the user location",
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
			MarkdownDescription: "The UUID of the user location",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"internal_host_detection": schema.SingleNestedAttribute{
			MarkdownDescription: "Configuration for detecting internal hosts using IP address and FQDN",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"fqdn": schema.StringAttribute{
					Validators: []validator.String{
						stringvalidator.LengthAtMost(255),
						stringvalidator.RegexMatches(regexp.MustCompile("^([\\*a-zA-Z0-9._-])+$"), "pattern must match "+"^([\\*a-zA-Z0-9._-])+$"),
					},
					MarkdownDescription: "user location fqdn",
					Required:            true,
				},
				"ip_address": schema.StringAttribute{
					Validators: []validator.String{
						stringvalidator.RegexMatches(regexp.MustCompile("^(\\*|25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.(\\*|25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.(\\*|25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.((\\*|(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?))|((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\/(3[0-2]|[1-2]?[0-9])))$"), "pattern must match "+"^(\\*|25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.(\\*|25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.(\\*|25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.((\\*|(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?))|((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\/(3[0-2]|[1-2]?[0-9])))$"),
					},
					MarkdownDescription: "user location ip address",
					Required:            true,
				},
			},
		},
		"ip_addresses": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "List of IP addresses that define the user location",
			Optional:            true,
		},
		"name": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(64),
				stringvalidator.RegexMatches(regexp.MustCompile("^[0-9a-zA-Z._-]+$"), "pattern must match "+"^[0-9a-zA-Z._-]+$"),
			},
			MarkdownDescription: "alphanumeric string [ 0-9a-zA-Z._-]",
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

// ForwardingProfileUserLocationsDataSourceSchema defines the schema for ForwardingProfileUserLocations data source
var ForwardingProfileUserLocationsDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "ForwardingProfileUserLocation data source",
	Attributes: map[string]dsschema.Attribute{
		"description": dsschema.StringAttribute{
			MarkdownDescription: "Description of the user location",
			Computed:            true,
		},
		"folder": dsschema.StringAttribute{
			MarkdownDescription: "The folder in which the resource is defined\n",
			Optional:            true,
			Computed:            true,
		},
		"id": dsschema.StringAttribute{
			MarkdownDescription: "The UUID of the user location",
			Required:            true,
		},
		"internal_host_detection": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Configuration for detecting internal hosts using IP address and FQDN",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"fqdn": dsschema.StringAttribute{
					MarkdownDescription: "user location fqdn",
					Computed:            true,
				},
				"ip_address": dsschema.StringAttribute{
					MarkdownDescription: "user location ip address",
					Computed:            true,
				},
			},
		},
		"ip_addresses": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "List of IP addresses that define the user location",
			Computed:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "alphanumeric string [ 0-9a-zA-Z._-]",
			Optional:            true,
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
	},
}

// ForwardingProfileUserLocationsListModel represents the data model for a list data source.
type ForwardingProfileUserLocationsListModel struct {
	Tfid    types.String                     `tfsdk:"tfid"`
	Data    []ForwardingProfileUserLocations `tfsdk:"data"`
	Limit   types.Int64                      `tfsdk:"limit"`
	Offset  types.Int64                      `tfsdk:"offset"`
	Name    types.String                     `tfsdk:"name"`
	Total   types.Int64                      `tfsdk:"total"`
	Folder  types.String                     `tfsdk:"folder"`
	Snippet types.String                     `tfsdk:"snippet"`
	Device  types.String                     `tfsdk:"device"`
}

// ForwardingProfileUserLocationsListDataSourceSchema defines the schema for a list data source.
var ForwardingProfileUserLocationsListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: ForwardingProfileUserLocationsDataSourceSchema.Attributes,
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
