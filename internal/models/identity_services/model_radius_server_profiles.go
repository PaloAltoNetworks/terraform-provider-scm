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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Package: identity_services
// This file contains models for the identity_services SDK package

// RadiusServerProfiles represents the Terraform model for RadiusServerProfiles
type RadiusServerProfiles struct {
	Tfid            types.String          `tfsdk:"tfid"`
	EncryptedValues basetypes.MapValue    `tfsdk:"encrypted_values"`
	Device          basetypes.StringValue `tfsdk:"device"`
	Folder          basetypes.StringValue `tfsdk:"folder"`
	Id              basetypes.StringValue `tfsdk:"id"`
	Name            basetypes.StringValue `tfsdk:"name"`
	Protocol        basetypes.ObjectValue `tfsdk:"protocol"`
	Retries         basetypes.Int64Value  `tfsdk:"retries"`
	Server          basetypes.ListValue   `tfsdk:"server"`
	Snippet         basetypes.StringValue `tfsdk:"snippet"`
	Timeout         basetypes.Int64Value  `tfsdk:"timeout"`
}

// RadiusServerProfilesProtocol represents a nested structure within the RadiusServerProfiles model
type RadiusServerProfilesProtocol struct {
	CHAP           basetypes.ObjectValue `tfsdk:"chap"`
	EAPTTLSWithPAP basetypes.ObjectValue `tfsdk:"eap_ttls_with_pap"`
	PAP            basetypes.ObjectValue `tfsdk:"pap"`
	PEAPMSCHAPv2   basetypes.ObjectValue `tfsdk:"peap_mscha_pv2"`
	PEAPWithGTC    basetypes.ObjectValue `tfsdk:"peap_with_gtc"`
}

// RadiusServerProfilesProtocolEAPTTLSWithPAP represents a nested structure within the RadiusServerProfiles model
type RadiusServerProfilesProtocolEAPTTLSWithPAP struct {
	AnonOuterId       basetypes.BoolValue   `tfsdk:"anon_outer_id"`
	RadiusCertProfile basetypes.StringValue `tfsdk:"radius_cert_profile"`
}

// RadiusServerProfilesProtocolPEAPMSCHAPv2 represents a nested structure within the RadiusServerProfiles model
type RadiusServerProfilesProtocolPEAPMSCHAPv2 struct {
	AllowPwdChange    basetypes.BoolValue   `tfsdk:"allow_pwd_change"`
	AnonOuterId       basetypes.BoolValue   `tfsdk:"anon_outer_id"`
	RadiusCertProfile basetypes.StringValue `tfsdk:"radius_cert_profile"`
}

// RadiusServerProfilesServerInner represents a nested structure within the RadiusServerProfiles model
type RadiusServerProfilesServerInner struct {
	IpAddress basetypes.StringValue `tfsdk:"ip_address"`
	Name      basetypes.StringValue `tfsdk:"name"`
	Port      basetypes.Int64Value  `tfsdk:"port"`
	Secret    basetypes.StringValue `tfsdk:"secret"`
}

// AttrTypes defines the attribute types for the RadiusServerProfiles model.
func (o RadiusServerProfiles) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":             basetypes.StringType{},
		"encrypted_values": basetypes.MapType{ElemType: basetypes.StringType{}},
		"device":           basetypes.StringType{},
		"folder":           basetypes.StringType{},
		"id":               basetypes.StringType{},
		"name":             basetypes.StringType{},
		"protocol": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"chap": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"eap_ttls_with_pap": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"anon_outer_id":       basetypes.BoolType{},
						"radius_cert_profile": basetypes.StringType{},
					},
				},
				"pap": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"peap_mscha_pv2": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"allow_pwd_change":    basetypes.BoolType{},
						"anon_outer_id":       basetypes.BoolType{},
						"radius_cert_profile": basetypes.StringType{},
					},
				},
				"peap_with_gtc": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"anon_outer_id":       basetypes.BoolType{},
						"radius_cert_profile": basetypes.StringType{},
					},
				},
			},
		},
		"retries": basetypes.Int64Type{},
		"server": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"ip_address": basetypes.StringType{},
				"name":       basetypes.StringType{},
				"port":       basetypes.Int64Type{},
				"secret":     basetypes.StringType{},
			},
		}},
		"snippet": basetypes.StringType{},
		"timeout": basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of RadiusServerProfiles objects.
func (o RadiusServerProfiles) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the RadiusServerProfilesProtocol model.
func (o RadiusServerProfilesProtocol) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"chap": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
		"eap_ttls_with_pap": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"anon_outer_id":       basetypes.BoolType{},
				"radius_cert_profile": basetypes.StringType{},
			},
		},
		"pap": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
		"peap_mscha_pv2": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"allow_pwd_change":    basetypes.BoolType{},
				"anon_outer_id":       basetypes.BoolType{},
				"radius_cert_profile": basetypes.StringType{},
			},
		},
		"peap_with_gtc": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"anon_outer_id":       basetypes.BoolType{},
				"radius_cert_profile": basetypes.StringType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of RadiusServerProfilesProtocol objects.
func (o RadiusServerProfilesProtocol) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the RadiusServerProfilesProtocolEAPTTLSWithPAP model.
func (o RadiusServerProfilesProtocolEAPTTLSWithPAP) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"anon_outer_id":       basetypes.BoolType{},
		"radius_cert_profile": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of RadiusServerProfilesProtocolEAPTTLSWithPAP objects.
func (o RadiusServerProfilesProtocolEAPTTLSWithPAP) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the RadiusServerProfilesProtocolPEAPMSCHAPv2 model.
func (o RadiusServerProfilesProtocolPEAPMSCHAPv2) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"allow_pwd_change":    basetypes.BoolType{},
		"anon_outer_id":       basetypes.BoolType{},
		"radius_cert_profile": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of RadiusServerProfilesProtocolPEAPMSCHAPv2 objects.
func (o RadiusServerProfilesProtocolPEAPMSCHAPv2) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the RadiusServerProfilesServerInner model.
func (o RadiusServerProfilesServerInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"ip_address": basetypes.StringType{},
		"name":       basetypes.StringType{},
		"port":       basetypes.Int64Type{},
		"secret":     basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of RadiusServerProfilesServerInner objects.
func (o RadiusServerProfilesServerInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// RadiusServerProfilesResourceSchema defines the schema for RadiusServerProfiles resource
var RadiusServerProfilesResourceSchema = schema.Schema{
	MarkdownDescription: "RadiusServerProfile resource",
	Attributes: map[string]schema.Attribute{
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
		"encrypted_values": schema.MapAttribute{
			ElementType:         basetypes.StringType{},
			MarkdownDescription: "Map of sensitive values returned from the API.",
			Computed:            true,
			Sensitive:           true,
		},
		"folder": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.ExactlyOneOf(
					path.MatchRelative().AtParent().AtName("device"),
					path.MatchRelative().AtParent().AtName("snippet"),
				),
				stringvalidator.LengthAtMost(64),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z\\d\\-_\\. ]+$"), "pattern must match "+"^[a-zA-Z\\d\\-_\\. ]+$"),
			},
			MarkdownDescription: "The folder in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"id": schema.StringAttribute{
			MarkdownDescription: "The UUID of the RADIUS server profile",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"name": schema.StringAttribute{
			MarkdownDescription: "The name of the RADIUS server profile",
			Required:            true,
		},
		"protocol": schema.SingleNestedAttribute{
			MarkdownDescription: "The RADIUS authentication protocol",
			Required:            true,
			Attributes: map[string]schema.Attribute{
				"chap": schema.SingleNestedAttribute{
					MarkdownDescription: "C h a p",
					Optional:            true,
					Attributes:          map[string]schema.Attribute{},
				},
				"eap_ttls_with_pap": schema.SingleNestedAttribute{
					MarkdownDescription: "E a p t t l s with p a p",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"anon_outer_id": schema.BoolAttribute{
							MarkdownDescription: "Anon outer id",
							Optional:            true,
						},
						"radius_cert_profile": schema.StringAttribute{
							MarkdownDescription: "Radius cert profile",
							Optional:            true,
						},
					},
				},
				"pap": schema.SingleNestedAttribute{
					MarkdownDescription: "P a p",
					Optional:            true,
					Attributes:          map[string]schema.Attribute{},
				},
				"peap_mscha_pv2": schema.SingleNestedAttribute{
					MarkdownDescription: "P e a p m s c h a pv2",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"allow_pwd_change": schema.BoolAttribute{
							MarkdownDescription: "Allow pwd change",
							Optional:            true,
						},
						"anon_outer_id": schema.BoolAttribute{
							MarkdownDescription: "Anon outer id",
							Optional:            true,
						},
						"radius_cert_profile": schema.StringAttribute{
							MarkdownDescription: "Radius cert profile",
							Optional:            true,
						},
					},
				},
				"peap_with_gtc": schema.SingleNestedAttribute{
					MarkdownDescription: "P e a p with g t c",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"anon_outer_id": schema.BoolAttribute{
							MarkdownDescription: "Anon outer id",
							Optional:            true,
						},
						"radius_cert_profile": schema.StringAttribute{
							MarkdownDescription: "Radius cert profile",
							Optional:            true,
						},
					},
				},
			},
		},
		"retries": schema.Int64Attribute{
			Validators: []validator.Int64{
				int64validator.Between(1, 5),
			},
			MarkdownDescription: "The number of RADIUS server retries",
			Optional:            true,
		},
		"server": schema.ListNestedAttribute{
			MarkdownDescription: "Server",
			Required:            true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"ip_address": schema.StringAttribute{
						MarkdownDescription: "The IP address of the RADIUS server",
						Optional:            true,
					},
					"name": schema.StringAttribute{
						MarkdownDescription: "The name of the RADIUS server",
						Optional:            true,
					},
					"port": schema.Int64Attribute{
						Validators: []validator.Int64{
							int64validator.Between(1, 65535),
						},
						MarkdownDescription: "The RADIUS server port",
						Optional:            true,
					},
					"secret": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.LengthAtMost(128),
						},
						MarkdownDescription: "The RADIUS secret",
						Optional:            true,
						Sensitive:           true,
					},
				},
			},
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
		"timeout": schema.Int64Attribute{
			Validators: []validator.Int64{
				int64validator.Between(1, 120),
			},
			MarkdownDescription: "The RADIUS server authentication timeout (seconds)",
			Optional:            true,
		},
	},
}

// RadiusServerProfilesDataSourceSchema defines the schema for RadiusServerProfiles data source
var RadiusServerProfilesDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "RadiusServerProfile data source",
	Attributes: map[string]dsschema.Attribute{
		"device": dsschema.StringAttribute{
			MarkdownDescription: "The device in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			Computed:            true,
		},
		"encrypted_values": dsschema.MapAttribute{
			ElementType:         basetypes.StringType{},
			MarkdownDescription: "Map of sensitive values returned from the API.",
			Computed:            true,
			Sensitive:           true,
		},
		"folder": dsschema.StringAttribute{
			MarkdownDescription: "The folder in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			Computed:            true,
		},
		"id": dsschema.StringAttribute{
			MarkdownDescription: "The UUID of the RADIUS server profile",
			Required:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "The name of the RADIUS server profile",
			Optional:            true,
			Computed:            true,
		},
		"protocol": dsschema.SingleNestedAttribute{
			MarkdownDescription: "The RADIUS authentication protocol",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"chap": dsschema.SingleNestedAttribute{
					MarkdownDescription: "C h a p",
					Computed:            true,
					Attributes:          map[string]dsschema.Attribute{},
				},
				"eap_ttls_with_pap": dsschema.SingleNestedAttribute{
					MarkdownDescription: "E a p t t l s with p a p",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"anon_outer_id": dsschema.BoolAttribute{
							MarkdownDescription: "Anon outer id",
							Computed:            true,
						},
						"radius_cert_profile": dsschema.StringAttribute{
							MarkdownDescription: "Radius cert profile",
							Computed:            true,
						},
					},
				},
				"pap": dsschema.SingleNestedAttribute{
					MarkdownDescription: "P a p",
					Computed:            true,
					Attributes:          map[string]dsschema.Attribute{},
				},
				"peap_mscha_pv2": dsschema.SingleNestedAttribute{
					MarkdownDescription: "P e a p m s c h a pv2",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"allow_pwd_change": dsschema.BoolAttribute{
							MarkdownDescription: "Allow pwd change",
							Computed:            true,
						},
						"anon_outer_id": dsschema.BoolAttribute{
							MarkdownDescription: "Anon outer id",
							Computed:            true,
						},
						"radius_cert_profile": dsschema.StringAttribute{
							MarkdownDescription: "Radius cert profile",
							Computed:            true,
						},
					},
				},
				"peap_with_gtc": dsschema.SingleNestedAttribute{
					MarkdownDescription: "P e a p with g t c",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"anon_outer_id": dsschema.BoolAttribute{
							MarkdownDescription: "Anon outer id",
							Computed:            true,
						},
						"radius_cert_profile": dsschema.StringAttribute{
							MarkdownDescription: "Radius cert profile",
							Computed:            true,
						},
					},
				},
			},
		},
		"retries": dsschema.Int64Attribute{
			MarkdownDescription: "The number of RADIUS server retries",
			Computed:            true,
		},
		"server": dsschema.ListNestedAttribute{
			MarkdownDescription: "Server",
			Computed:            true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: map[string]dsschema.Attribute{
					"ip_address": dsschema.StringAttribute{
						MarkdownDescription: "The IP address of the RADIUS server",
						Computed:            true,
					},
					"name": dsschema.StringAttribute{
						MarkdownDescription: "The name of the RADIUS server",
						Computed:            true,
					},
					"port": dsschema.Int64Attribute{
						MarkdownDescription: "The RADIUS server port",
						Computed:            true,
					},
					"secret": dsschema.StringAttribute{
						MarkdownDescription: "The RADIUS secret",
						Computed:            true,
						Sensitive:           true,
					},
				},
			},
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
		"timeout": dsschema.Int64Attribute{
			MarkdownDescription: "The RADIUS server authentication timeout (seconds)",
			Computed:            true,
		},
	},
}

// RadiusServerProfilesListModel represents the data model for a list data source.
type RadiusServerProfilesListModel struct {
	Tfid    types.String           `tfsdk:"tfid"`
	Data    []RadiusServerProfiles `tfsdk:"data"`
	Limit   types.Int64            `tfsdk:"limit"`
	Offset  types.Int64            `tfsdk:"offset"`
	Name    types.String           `tfsdk:"name"`
	Total   types.Int64            `tfsdk:"total"`
	Folder  types.String           `tfsdk:"folder"`
	Snippet types.String           `tfsdk:"snippet"`
	Device  types.String           `tfsdk:"device"`
}

// RadiusServerProfilesListDataSourceSchema defines the schema for a list data source.
var RadiusServerProfilesListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: RadiusServerProfilesDataSourceSchema.Attributes,
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
