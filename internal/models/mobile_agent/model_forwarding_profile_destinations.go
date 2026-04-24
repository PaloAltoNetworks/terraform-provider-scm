package models

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
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

// ForwardingProfileDestinations represents the Terraform model for ForwardingProfileDestinations
type ForwardingProfileDestinations struct {
	Tfid        types.String          `tfsdk:"tfid"`
	Description basetypes.StringValue `tfsdk:"description"`
	Fqdn        basetypes.ListValue   `tfsdk:"fqdn"`
	Id          basetypes.StringValue `tfsdk:"id"`
	IpAddresses basetypes.ListValue   `tfsdk:"ip_addresses"`
	Name        basetypes.StringValue `tfsdk:"name"`
	Folder      basetypes.StringValue `tfsdk:"folder"`
}

// ForwardingProfileDestinationFqdnEntry represents a nested structure within the ForwardingProfileDestinations model
type ForwardingProfileDestinationFqdnEntry struct {
	Name basetypes.StringValue `tfsdk:"name"`
	Port basetypes.Int64Value  `tfsdk:"port"`
}

// ForwardingProfileDestinationIpEntry represents a nested structure within the ForwardingProfileDestinations model
type ForwardingProfileDestinationIpEntry struct {
	Name basetypes.StringValue `tfsdk:"name"`
	Port basetypes.Int64Value  `tfsdk:"port"`
}

// AttrTypes defines the attribute types for the ForwardingProfileDestinations model.
func (o ForwardingProfileDestinations) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":        basetypes.StringType{},
		"description": basetypes.StringType{},
		"fqdn": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"name": basetypes.StringType{},
				"port": basetypes.Int64Type{},
			},
		}},
		"id": basetypes.StringType{},
		"ip_addresses": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"name": basetypes.StringType{},
				"port": basetypes.Int64Type{},
			},
		}},
		"name":   basetypes.StringType{},
		"folder": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of ForwardingProfileDestinations objects.
func (o ForwardingProfileDestinations) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ForwardingProfileDestinationFqdnEntry model.
func (o ForwardingProfileDestinationFqdnEntry) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name": basetypes.StringType{},
		"port": basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of ForwardingProfileDestinationFqdnEntry objects.
func (o ForwardingProfileDestinationFqdnEntry) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ForwardingProfileDestinationIpEntry model.
func (o ForwardingProfileDestinationIpEntry) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name": basetypes.StringType{},
		"port": basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of ForwardingProfileDestinationIpEntry objects.
func (o ForwardingProfileDestinationIpEntry) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// ForwardingProfileDestinationsResourceSchema defines the schema for ForwardingProfileDestinations resource
var ForwardingProfileDestinationsResourceSchema = schema.Schema{
	MarkdownDescription: "ForwardingProfileDestination resource",
	Attributes: map[string]schema.Attribute{
		"description": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(1023),
			},
			MarkdownDescription: "description of the destination",
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
		"fqdn": schema.ListNestedAttribute{
			MarkdownDescription: "List of FQDN based destination entries",
			Optional:            true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"name": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.LengthAtMost(255),
							stringvalidator.RegexMatches(regexp.MustCompile("^([\\*a-zA-Z0-9._-])+\\$?$"), "pattern must match "+"^([\\*a-zA-Z0-9._-])+\\$?$"),
						},
						MarkdownDescription: "alphanumeric string [*0-9a-zA-Z._-] and at most one $ by the end",
						Required:            true,
					},
					"port": schema.Int64Attribute{
						Validators: []validator.Int64{
							int64validator.Between(1, 65535),
						},
						MarkdownDescription: "Port number for fqdn based destination",
						Optional:            true,
					},
				},
			},
		},
		"id": schema.StringAttribute{
			MarkdownDescription: "The UUID of the destination",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"ip_addresses": schema.ListNestedAttribute{
			MarkdownDescription: "List of IP address based destination entries",
			Optional:            true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"name": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile("^(\\*|25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.(\\*|25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.(\\*|25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.(((\\*|(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?))|((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\/(3[0-2]|[1-2]?[0-9]))))$"), "pattern must match "+"^(\\*|25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.(\\*|25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.(\\*|25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.(((\\*|(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?))|((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\/(3[0-2]|[1-2]?[0-9]))))$"),
						},
						MarkdownDescription: "IP address with wildcards and CIDR notation support",
						Required:            true,
					},
					"port": schema.Int64Attribute{
						Validators: []validator.Int64{
							int64validator.Between(1, 65535),
						},
						MarkdownDescription: "Port number for IP address based destination",
						Optional:            true,
					},
				},
			},
		},
		"name": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(64),
				stringvalidator.RegexMatches(regexp.MustCompile("^[0-9a-zA-Z._ -]+$"), "pattern must match "+"^[0-9a-zA-Z._ -]+$"),
			},
			MarkdownDescription: "alphanumeric string [ 0-9a-zA-Z._ -]",
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

// ForwardingProfileDestinationsDataSourceSchema defines the schema for ForwardingProfileDestinations data source
var ForwardingProfileDestinationsDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "ForwardingProfileDestination data source",
	Attributes: map[string]dsschema.Attribute{
		"description": dsschema.StringAttribute{
			MarkdownDescription: "description of the destination",
			Computed:            true,
		},
		"folder": dsschema.StringAttribute{
			MarkdownDescription: "The folder in which the resource is defined\n",
			Optional:            true,
			Computed:            true,
		},
		"fqdn": dsschema.ListNestedAttribute{
			MarkdownDescription: "List of FQDN based destination entries",
			Computed:            true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: map[string]dsschema.Attribute{
					"name": dsschema.StringAttribute{
						MarkdownDescription: "alphanumeric string [*0-9a-zA-Z._-] and at most one $ by the end",
						Computed:            true,
					},
					"port": dsschema.Int64Attribute{
						MarkdownDescription: "Port number for fqdn based destination",
						Computed:            true,
					},
				},
			},
		},
		"id": dsschema.StringAttribute{
			MarkdownDescription: "The UUID of the destination",
			Required:            true,
		},
		"ip_addresses": dsschema.ListNestedAttribute{
			MarkdownDescription: "List of IP address based destination entries",
			Computed:            true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: map[string]dsschema.Attribute{
					"name": dsschema.StringAttribute{
						MarkdownDescription: "IP address with wildcards and CIDR notation support",
						Computed:            true,
					},
					"port": dsschema.Int64Attribute{
						MarkdownDescription: "Port number for IP address based destination",
						Computed:            true,
					},
				},
			},
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "alphanumeric string [ 0-9a-zA-Z._ -]",
			Optional:            true,
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
	},
}

// ForwardingProfileDestinationsListModel represents the data model for a list data source.
type ForwardingProfileDestinationsListModel struct {
	Tfid    types.String                    `tfsdk:"tfid"`
	Data    []ForwardingProfileDestinations `tfsdk:"data"`
	Limit   types.Int64                     `tfsdk:"limit"`
	Offset  types.Int64                     `tfsdk:"offset"`
	Name    types.String                    `tfsdk:"name"`
	Total   types.Int64                     `tfsdk:"total"`
	Folder  types.String                    `tfsdk:"folder"`
	Snippet types.String                    `tfsdk:"snippet"`
	Device  types.String                    `tfsdk:"device"`
}

// ForwardingProfileDestinationsListDataSourceSchema defines the schema for a list data source.
var ForwardingProfileDestinationsListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: ForwardingProfileDestinationsDataSourceSchema.Attributes,
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
