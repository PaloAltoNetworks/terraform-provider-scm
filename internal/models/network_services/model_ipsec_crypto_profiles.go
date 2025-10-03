package models

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	dsschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Package: network_services
// This file contains models for the network_services SDK package

// IpsecCryptoProfiles represents the Terraform model for IpsecCryptoProfiles
type IpsecCryptoProfiles struct {
	Tfid     types.String          `tfsdk:"tfid"`
	Ah       basetypes.ObjectValue `tfsdk:"ah"`
	Device   basetypes.StringValue `tfsdk:"device"`
	DhGroup  basetypes.StringValue `tfsdk:"dh_group"`
	Esp      basetypes.ObjectValue `tfsdk:"esp"`
	Folder   basetypes.StringValue `tfsdk:"folder"`
	Id       basetypes.StringValue `tfsdk:"id"`
	Lifesize basetypes.ObjectValue `tfsdk:"lifesize"`
	Lifetime basetypes.ObjectValue `tfsdk:"lifetime"`
	Name     basetypes.StringValue `tfsdk:"name"`
	Snippet  basetypes.StringValue `tfsdk:"snippet"`
}

// IpsecCryptoProfilesAh represents a nested structure within the IpsecCryptoProfiles model
type IpsecCryptoProfilesAh struct {
	Authentication basetypes.ListValue `tfsdk:"authentication"`
}

// IpsecCryptoProfilesEsp represents a nested structure within the IpsecCryptoProfiles model
type IpsecCryptoProfilesEsp struct {
	Authentication basetypes.ListValue `tfsdk:"authentication"`
	Encryption     basetypes.ListValue `tfsdk:"encryption"`
}

// IpsecCryptoProfilesLifesize represents a nested structure within the IpsecCryptoProfiles model
type IpsecCryptoProfilesLifesize struct {
	Gb basetypes.Int64Value `tfsdk:"gb"`
	Kb basetypes.Int64Value `tfsdk:"kb"`
	Mb basetypes.Int64Value `tfsdk:"mb"`
	Tb basetypes.Int64Value `tfsdk:"tb"`
}

// IpsecCryptoProfilesLifetime represents a nested structure within the IpsecCryptoProfiles model
type IpsecCryptoProfilesLifetime struct {
	Days    basetypes.Int64Value `tfsdk:"days"`
	Hours   basetypes.Int64Value `tfsdk:"hours"`
	Minutes basetypes.Int64Value `tfsdk:"minutes"`
	Seconds basetypes.Int64Value `tfsdk:"seconds"`
}

// AttrTypes defines the attribute types for the IpsecCryptoProfiles model.
func (o IpsecCryptoProfiles) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid": basetypes.StringType{},
		"ah": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"authentication": basetypes.ListType{ElemType: basetypes.StringType{}},
			},
		},
		"device":   basetypes.StringType{},
		"dh_group": basetypes.StringType{},
		"esp": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"authentication": basetypes.ListType{ElemType: basetypes.StringType{}},
				"encryption":     basetypes.ListType{ElemType: basetypes.StringType{}},
			},
		},
		"folder": basetypes.StringType{},
		"id":     basetypes.StringType{},
		"lifesize": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"gb": basetypes.Int64Type{},
				"kb": basetypes.Int64Type{},
				"mb": basetypes.Int64Type{},
				"tb": basetypes.Int64Type{},
			},
		},
		"lifetime": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"days":    basetypes.Int64Type{},
				"hours":   basetypes.Int64Type{},
				"minutes": basetypes.Int64Type{},
				"seconds": basetypes.Int64Type{},
			},
		},
		"name":    basetypes.StringType{},
		"snippet": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of IpsecCryptoProfiles objects.
func (o IpsecCryptoProfiles) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the IpsecCryptoProfilesAh model.
func (o IpsecCryptoProfilesAh) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"authentication": basetypes.ListType{ElemType: basetypes.StringType{}},
	}
}

// AttrType returns the attribute type for a list of IpsecCryptoProfilesAh objects.
func (o IpsecCryptoProfilesAh) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the IpsecCryptoProfilesEsp model.
func (o IpsecCryptoProfilesEsp) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"authentication": basetypes.ListType{ElemType: basetypes.StringType{}},
		"encryption":     basetypes.ListType{ElemType: basetypes.StringType{}},
	}
}

// AttrType returns the attribute type for a list of IpsecCryptoProfilesEsp objects.
func (o IpsecCryptoProfilesEsp) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the IpsecCryptoProfilesLifesize model.
func (o IpsecCryptoProfilesLifesize) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"gb": basetypes.Int64Type{},
		"kb": basetypes.Int64Type{},
		"mb": basetypes.Int64Type{},
		"tb": basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of IpsecCryptoProfilesLifesize objects.
func (o IpsecCryptoProfilesLifesize) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the IpsecCryptoProfilesLifetime model.
func (o IpsecCryptoProfilesLifetime) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"days":    basetypes.Int64Type{},
		"hours":   basetypes.Int64Type{},
		"minutes": basetypes.Int64Type{},
		"seconds": basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of IpsecCryptoProfilesLifetime objects.
func (o IpsecCryptoProfilesLifetime) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// IpsecCryptoProfilesResourceSchema defines the schema for IpsecCryptoProfiles resource
var IpsecCryptoProfilesResourceSchema = schema.Schema{
	MarkdownDescription: "IpsecCryptoProfile resource",
	Attributes: map[string]schema.Attribute{
		"ah": schema.SingleNestedAttribute{
			MarkdownDescription: "Ah",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"authentication": schema.ListAttribute{
					ElementType:         types.StringType,
					MarkdownDescription: "Authentication",
					Required:            true,
				},
			},
		},
		"device": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(64),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z\\d-_\\. ]+$"), "pattern must match "+"^[a-zA-Z\\d-_\\. ]+$"),
			},
			MarkdownDescription: "The device in which the resource is defined",
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"dh_group": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.OneOf("no-pfs", "group1", "group2", "group5", "group14", "group19", "group20"),
			},
			MarkdownDescription: "phase-2 DH group (PFS DH group)",
			Optional:            true,
			Computed:            true,
			Default:             stringdefault.StaticString("group2"),
		},
		"esp": schema.SingleNestedAttribute{
			MarkdownDescription: "Esp",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"authentication": schema.ListAttribute{
					ElementType:         types.StringType,
					MarkdownDescription: "Authentication algorithm",
					Required:            true,
				},
				"encryption": schema.ListAttribute{
					ElementType:         types.StringType,
					MarkdownDescription: "Encryption algorithm",
					Required:            true,
				},
			},
		},
		"folder": schema.StringAttribute{
			Validators: []validator.String{
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
		"lifesize": schema.SingleNestedAttribute{
			MarkdownDescription: "Lifesize",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"gb": schema.Int64Attribute{
					Validators: []validator.Int64{
						int64validator.ExactlyOneOf(
							path.MatchRelative().AtParent().AtName("kb"),
							path.MatchRelative().AtParent().AtName("mb"),
							path.MatchRelative().AtParent().AtName("tb"),
						),
						int64validator.Between(1, 65535),
					},
					MarkdownDescription: "specify lifesize in gigabytes(GB)",
					Optional:            true,
				},
				"kb": schema.Int64Attribute{
					Validators: []validator.Int64{
						int64validator.ExactlyOneOf(
							path.MatchRelative().AtParent().AtName("mb"),
							path.MatchRelative().AtParent().AtName("gb"),
							path.MatchRelative().AtParent().AtName("tb"),
						),
						int64validator.Between(1, 65535),
					},
					MarkdownDescription: "specify lifesize in kilobytes(KB)",
					Optional:            true,
				},
				"mb": schema.Int64Attribute{
					Validators: []validator.Int64{
						int64validator.ExactlyOneOf(
							path.MatchRelative().AtParent().AtName("kb"),
							path.MatchRelative().AtParent().AtName("gb"),
							path.MatchRelative().AtParent().AtName("tb"),
						),
						int64validator.Between(1, 65535),
					},
					MarkdownDescription: "specify lifesize in megabytes(MB)",
					Optional:            true,
				},
				"tb": schema.Int64Attribute{
					Validators: []validator.Int64{
						int64validator.ExactlyOneOf(
							path.MatchRelative().AtParent().AtName("kb"),
							path.MatchRelative().AtParent().AtName("mb"),
							path.MatchRelative().AtParent().AtName("gb"),
						),
						int64validator.Between(1, 65535),
					},
					MarkdownDescription: "specify lifesize in terabytes(TB)",
					Optional:            true,
				},
			},
		},
		"lifetime": schema.SingleNestedAttribute{
			MarkdownDescription: "Ipsec crypto profile lifetime",
			Required:            true,
			Attributes: map[string]schema.Attribute{
				"days": schema.Int64Attribute{
					Validators: []validator.Int64{
						int64validator.ExactlyOneOf(
							path.MatchRelative().AtParent().AtName("seconds"),
							path.MatchRelative().AtParent().AtName("minutes"),
							path.MatchRelative().AtParent().AtName("hours"),
						),
						int64validator.Between(1, 365),
					},
					MarkdownDescription: "specify lifetime in days",
					Optional:            true,
				},
				"hours": schema.Int64Attribute{
					Validators: []validator.Int64{
						int64validator.ExactlyOneOf(
							path.MatchRelative().AtParent().AtName("seconds"),
							path.MatchRelative().AtParent().AtName("minutes"),
							path.MatchRelative().AtParent().AtName("days"),
						),
						int64validator.Between(1, 65535),
					},
					MarkdownDescription: "specify lifetime in hours",
					Optional:            true,
				},
				"minutes": schema.Int64Attribute{
					Validators: []validator.Int64{
						int64validator.ExactlyOneOf(
							path.MatchRelative().AtParent().AtName("seconds"),
							path.MatchRelative().AtParent().AtName("hours"),
							path.MatchRelative().AtParent().AtName("days"),
						),
						int64validator.Between(3, 65535),
					},
					MarkdownDescription: "specify lifetime in minutes",
					Optional:            true,
				},
				"seconds": schema.Int64Attribute{
					Validators: []validator.Int64{
						int64validator.ExactlyOneOf(
							path.MatchRelative().AtParent().AtName("minutes"),
							path.MatchRelative().AtParent().AtName("hours"),
							path.MatchRelative().AtParent().AtName("days"),
						),
						int64validator.Between(180, 65535),
					},
					MarkdownDescription: "specify lifetime in seconds",
					Optional:            true,
				},
			},
		},
		"name": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(31),
			},
			MarkdownDescription: "Alphanumeric string begin with letter: [0-9a-zA-Z._-]",
			Required:            true,
		},
		"snippet": schema.StringAttribute{
			Validators: []validator.String{
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

// IpsecCryptoProfilesDataSourceSchema defines the schema for IpsecCryptoProfiles data source
var IpsecCryptoProfilesDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "IpsecCryptoProfile data source",
	Attributes: map[string]dsschema.Attribute{
		"ah": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Ah",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"authentication": dsschema.ListAttribute{
					ElementType:         types.StringType,
					MarkdownDescription: "Authentication",
					Computed:            true,
				},
			},
		},
		"device": dsschema.StringAttribute{
			MarkdownDescription: "The device in which the resource is defined",
			Computed:            true,
		},
		"dh_group": dsschema.StringAttribute{
			MarkdownDescription: "phase-2 DH group (PFS DH group)",
			Computed:            true,
		},
		"esp": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Esp",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"authentication": dsschema.ListAttribute{
					ElementType:         types.StringType,
					MarkdownDescription: "Authentication algorithm",
					Computed:            true,
				},
				"encryption": dsschema.ListAttribute{
					ElementType:         types.StringType,
					MarkdownDescription: "Encryption algorithm",
					Computed:            true,
				},
			},
		},
		"folder": dsschema.StringAttribute{
			MarkdownDescription: "The folder in which the resource is defined",
			Computed:            true,
		},
		"id": dsschema.StringAttribute{
			MarkdownDescription: "UUID of the resource",
			Required:            true,
		},
		"lifesize": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Lifesize",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"gb": dsschema.Int64Attribute{
					MarkdownDescription: "specify lifesize in gigabytes(GB)",
					Computed:            true,
				},
				"kb": dsschema.Int64Attribute{
					MarkdownDescription: "specify lifesize in kilobytes(KB)",
					Computed:            true,
				},
				"mb": dsschema.Int64Attribute{
					MarkdownDescription: "specify lifesize in megabytes(MB)",
					Computed:            true,
				},
				"tb": dsschema.Int64Attribute{
					MarkdownDescription: "specify lifesize in terabytes(TB)",
					Computed:            true,
				},
			},
		},
		"lifetime": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Ipsec crypto profile lifetime",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"days": dsschema.Int64Attribute{
					MarkdownDescription: "specify lifetime in days",
					Computed:            true,
				},
				"hours": dsschema.Int64Attribute{
					MarkdownDescription: "specify lifetime in hours",
					Computed:            true,
				},
				"minutes": dsschema.Int64Attribute{
					MarkdownDescription: "specify lifetime in minutes",
					Computed:            true,
				},
				"seconds": dsschema.Int64Attribute{
					MarkdownDescription: "specify lifetime in seconds",
					Computed:            true,
				},
			},
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "Alphanumeric string begin with letter: [0-9a-zA-Z._-]",
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

// IpsecCryptoProfilesListModel represents the data model for a list data source.
type IpsecCryptoProfilesListModel struct {
	Tfid    types.String          `tfsdk:"tfid"`
	Data    []IpsecCryptoProfiles `tfsdk:"data"`
	Limit   types.Int64           `tfsdk:"limit"`
	Offset  types.Int64           `tfsdk:"offset"`
	Name    types.String          `tfsdk:"name"`
	Total   types.Int64           `tfsdk:"total"`
	Folder  types.String          `tfsdk:"folder"`
	Snippet types.String          `tfsdk:"snippet"`
	Device  types.String          `tfsdk:"device"`
}

// IpsecCryptoProfilesListDataSourceSchema defines the schema for a list data source.
var IpsecCryptoProfilesListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: IpsecCryptoProfilesDataSourceSchema.Attributes,
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
