package models

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	dsschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/paloaltonetworks/terraform-provider-scm/internal/utils"
)

// Package: network_services
// This file contains models for the network_services SDK package

// IkeCryptoProfiles represents the Terraform model for IkeCryptoProfiles
type IkeCryptoProfiles struct {
	Tfid                   types.String          `tfsdk:"tfid"`
	AuthenticationMultiple basetypes.Int64Value  `tfsdk:"authentication_multiple"`
	Device                 basetypes.StringValue `tfsdk:"device"`
	DhGroup                basetypes.ListValue   `tfsdk:"dh_group"`
	Encryption             basetypes.ListValue   `tfsdk:"encryption"`
	Folder                 basetypes.StringValue `tfsdk:"folder"`
	Hash                   basetypes.ListValue   `tfsdk:"hash"`
	Id                     basetypes.StringValue `tfsdk:"id"`
	Lifetime               basetypes.ObjectValue `tfsdk:"lifetime"`
	Name                   basetypes.StringValue `tfsdk:"name"`
	Snippet                basetypes.StringValue `tfsdk:"snippet"`
}

// IkeCryptoProfilesLifetime represents a nested structure within the IkeCryptoProfiles model
type IkeCryptoProfilesLifetime struct {
	Days    basetypes.Int64Value `tfsdk:"days"`
	Hours   basetypes.Int64Value `tfsdk:"hours"`
	Minutes basetypes.Int64Value `tfsdk:"minutes"`
	Seconds basetypes.Int64Value `tfsdk:"seconds"`
}

// AttrTypes defines the attribute types for the IkeCryptoProfiles model.
func (o IkeCryptoProfiles) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":                    basetypes.StringType{},
		"authentication_multiple": basetypes.Int64Type{},
		"device":                  basetypes.StringType{},
		"dh_group":                basetypes.ListType{ElemType: basetypes.StringType{}},
		"encryption":              basetypes.ListType{ElemType: basetypes.StringType{}},
		"folder":                  basetypes.StringType{},
		"hash":                    basetypes.ListType{ElemType: basetypes.StringType{}},
		"id":                      basetypes.StringType{},
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

// AttrType returns the attribute type for a list of IkeCryptoProfiles objects.
func (o IkeCryptoProfiles) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the IkeCryptoProfilesLifetime model.
func (o IkeCryptoProfilesLifetime) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"days":    basetypes.Int64Type{},
		"hours":   basetypes.Int64Type{},
		"minutes": basetypes.Int64Type{},
		"seconds": basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of IkeCryptoProfilesLifetime objects.
func (o IkeCryptoProfilesLifetime) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// IkeCryptoProfilesResourceSchema defines the schema for IkeCryptoProfiles resource
var IkeCryptoProfilesResourceSchema = schema.Schema{
	MarkdownDescription: "IkeCryptoProfile resource",
	Attributes: map[string]schema.Attribute{
		"authentication_multiple": schema.Int64Attribute{
			MarkdownDescription: "IKEv2 SA reauthentication interval equals authetication-multiple * rekey-lifetime; 0 means reauthentication disabled",
			Optional:            true,
			Computed:            true,
			Default:             int64default.StaticInt64(0),
		},
		"device": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.ExactlyOneOf(
					path.MatchRelative().AtParent().AtName("folder"),
					path.MatchRelative().AtParent().AtName("snippet"),
				),
				stringvalidator.LengthAtMost(64),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z\\d\\-_\\. ]+$"), "pattern must match "+"^[a-zA-Z\\d\\-_\\. ]+$"),
			},
			MarkdownDescription: "The device in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"dh_group": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Dh group",
			Required:            true,
		},
		"encryption": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Encryption algorithm",
			Required:            true,
		},
		"folder": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.ExactlyOneOf(
					path.MatchRelative().AtParent().AtName("device"),
					path.MatchRelative().AtParent().AtName("snippet"),
				),
				stringvalidator.LengthAtMost(64),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z\\d\\-_\\. ]+$"), "pattern must match "+"^[a-zA-Z\\d\\-_\\. ]+$"),
				utils.FolderValidator(),
			},
			MarkdownDescription: "The folder in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"hash": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Hash",
			Required:            true,
		},
		"id": schema.StringAttribute{
			MarkdownDescription: "UUID of the resource",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"lifetime": schema.SingleNestedAttribute{
			MarkdownDescription: "Ike crypto profile lifetime",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"days": schema.Int64Attribute{
					Validators: []validator.Int64{
						int64validator.ConflictsWith(
							path.MatchRelative().AtParent().AtName("hours"),
							path.MatchRelative().AtParent().AtName("minutes"),
							path.MatchRelative().AtParent().AtName("seconds"),
						),
						int64validator.Between(1, 365),
					},
					MarkdownDescription: "specify lifetime in days\n\n> ℹ️ **Note:** You must specify exactly one of `days`, `hours`, `minutes`, and `seconds`.",
					Optional:            true,
				},
				"hours": schema.Int64Attribute{
					Validators: []validator.Int64{
						int64validator.ConflictsWith(
							path.MatchRelative().AtParent().AtName("days"),
							path.MatchRelative().AtParent().AtName("minutes"),
							path.MatchRelative().AtParent().AtName("seconds"),
						),
						int64validator.Between(1, 65535),
					},
					MarkdownDescription: "specify lifetime in hours\n\n> ℹ️ **Note:** You must specify exactly one of `days`, `hours`, `minutes`, and `seconds`.",
					Optional:            true,
				},
				"minutes": schema.Int64Attribute{
					Validators: []validator.Int64{
						int64validator.ConflictsWith(
							path.MatchRelative().AtParent().AtName("days"),
							path.MatchRelative().AtParent().AtName("hours"),
							path.MatchRelative().AtParent().AtName("seconds"),
						),
						int64validator.Between(3, 65535),
					},
					MarkdownDescription: "specify lifetime in minutes\n\n> ℹ️ **Note:** You must specify exactly one of `days`, `hours`, `minutes`, and `seconds`.",
					Optional:            true,
				},
				"seconds": schema.Int64Attribute{
					Validators: []validator.Int64{
						int64validator.ConflictsWith(
							path.MatchRelative().AtParent().AtName("days"),
							path.MatchRelative().AtParent().AtName("hours"),
							path.MatchRelative().AtParent().AtName("minutes"),
						),
						int64validator.Between(180, 65535),
					},
					MarkdownDescription: "specify lifetime in seconds\n\n> ℹ️ **Note:** You must specify exactly one of `days`, `hours`, `minutes`, and `seconds`.",
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
				stringvalidator.ExactlyOneOf(
					path.MatchRelative().AtParent().AtName("device"),
					path.MatchRelative().AtParent().AtName("folder"),
				),
				stringvalidator.LengthAtMost(64),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z\\d\\-_\\. ]+$"), "pattern must match "+"^[a-zA-Z\\d\\-_\\. ]+$"),
			},
			MarkdownDescription: "The snippet in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
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

// IkeCryptoProfilesDataSourceSchema defines the schema for IkeCryptoProfiles data source
var IkeCryptoProfilesDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "IkeCryptoProfile data source",
	Attributes: map[string]dsschema.Attribute{
		"authentication_multiple": dsschema.Int64Attribute{
			MarkdownDescription: "IKEv2 SA reauthentication interval equals authetication-multiple * rekey-lifetime; 0 means reauthentication disabled",
			Computed:            true,
		},
		"device": dsschema.StringAttribute{
			MarkdownDescription: "The device in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			Computed:            true,
		},
		"dh_group": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Dh group",
			Computed:            true,
		},
		"encryption": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Encryption algorithm",
			Computed:            true,
		},
		"folder": dsschema.StringAttribute{
			MarkdownDescription: "The folder in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			Computed:            true,
		},
		"hash": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Hash",
			Computed:            true,
		},
		"id": dsschema.StringAttribute{
			MarkdownDescription: "UUID of the resource",
			Required:            true,
		},
		"lifetime": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Ike crypto profile lifetime",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"days": dsschema.Int64Attribute{
					MarkdownDescription: "specify lifetime in days\n\n> ℹ️ **Note:** You must specify exactly one of `days`, `hours`, `minutes`, and `seconds`.",
					Computed:            true,
				},
				"hours": dsschema.Int64Attribute{
					MarkdownDescription: "specify lifetime in hours\n\n> ℹ️ **Note:** You must specify exactly one of `days`, `hours`, `minutes`, and `seconds`.",
					Computed:            true,
				},
				"minutes": dsschema.Int64Attribute{
					MarkdownDescription: "specify lifetime in minutes\n\n> ℹ️ **Note:** You must specify exactly one of `days`, `hours`, `minutes`, and `seconds`.",
					Computed:            true,
				},
				"seconds": dsschema.Int64Attribute{
					MarkdownDescription: "specify lifetime in seconds\n\n> ℹ️ **Note:** You must specify exactly one of `days`, `hours`, `minutes`, and `seconds`.",
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
			MarkdownDescription: "The snippet in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
	},
}

// IkeCryptoProfilesListModel represents the data model for a list data source.
type IkeCryptoProfilesListModel struct {
	Tfid    types.String        `tfsdk:"tfid"`
	Data    []IkeCryptoProfiles `tfsdk:"data"`
	Limit   types.Int64         `tfsdk:"limit"`
	Offset  types.Int64         `tfsdk:"offset"`
	Name    types.String        `tfsdk:"name"`
	Total   types.Int64         `tfsdk:"total"`
	Folder  types.String        `tfsdk:"folder"`
	Snippet types.String        `tfsdk:"snippet"`
	Device  types.String        `tfsdk:"device"`
}

// IkeCryptoProfilesListDataSourceSchema defines the schema for a list data source.
var IkeCryptoProfilesListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: IkeCryptoProfilesDataSourceSchema.Attributes,
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
