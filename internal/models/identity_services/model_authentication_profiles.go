package models

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/objectvalidator"
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

// AuthenticationProfiles represents the Terraform model for AuthenticationProfiles
type AuthenticationProfiles struct {
	Tfid             types.String          `tfsdk:"tfid"`
	AllowList        basetypes.ListValue   `tfsdk:"allow_list"`
	Device           basetypes.StringValue `tfsdk:"device"`
	Folder           basetypes.StringValue `tfsdk:"folder"`
	Id               basetypes.StringValue `tfsdk:"id"`
	Lockout          basetypes.ObjectValue `tfsdk:"lockout"`
	Method           basetypes.ObjectValue `tfsdk:"method"`
	MultiFactorAuth  basetypes.ObjectValue `tfsdk:"multi_factor_auth"`
	Name             basetypes.StringValue `tfsdk:"name"`
	SingleSignOn     basetypes.ObjectValue `tfsdk:"single_sign_on"`
	Snippet          basetypes.StringValue `tfsdk:"snippet"`
	UserDomain       basetypes.StringValue `tfsdk:"user_domain"`
	UsernameModifier basetypes.StringValue `tfsdk:"username_modifier"`
}

// AuthenticationProfilesLockout represents a nested structure within the AuthenticationProfiles model
type AuthenticationProfilesLockout struct {
	FailedAttempts basetypes.Int64Value `tfsdk:"failed_attempts"`
	LockoutTime    basetypes.Int64Value `tfsdk:"lockout_time"`
}

// AuthenticationProfilesMethod represents a nested structure within the AuthenticationProfiles model
type AuthenticationProfilesMethod struct {
	Cloud         basetypes.ObjectValue `tfsdk:"cloud"`
	Kerberos      basetypes.ObjectValue `tfsdk:"kerberos"`
	Ldap          basetypes.ObjectValue `tfsdk:"ldap"`
	LocalDatabase basetypes.ObjectValue `tfsdk:"local_database"`
	Radius        basetypes.ObjectValue `tfsdk:"radius"`
	SamlIdp       basetypes.ObjectValue `tfsdk:"saml_idp"`
	Tacplus       basetypes.ObjectValue `tfsdk:"tacplus"`
}

// AuthenticationProfilesMethodCloud represents a nested structure within the AuthenticationProfiles model
type AuthenticationProfilesMethodCloud struct {
	ProfileName basetypes.StringValue `tfsdk:"profile_name"`
}

// AuthenticationProfilesMethodKerberos represents a nested structure within the AuthenticationProfiles model
type AuthenticationProfilesMethodKerberos struct {
	Realm         basetypes.StringValue `tfsdk:"realm"`
	ServerProfile basetypes.StringValue `tfsdk:"server_profile"`
}

// AuthenticationProfilesMethodLdap represents a nested structure within the AuthenticationProfiles model
type AuthenticationProfilesMethodLdap struct {
	LoginAttribute basetypes.StringValue `tfsdk:"login_attribute"`
	PasswdExpDays  basetypes.Int64Value  `tfsdk:"passwd_exp_days"`
	ServerProfile  basetypes.StringValue `tfsdk:"server_profile"`
}

// AuthenticationProfilesMethodRadius represents a nested structure within the AuthenticationProfiles model
type AuthenticationProfilesMethodRadius struct {
	Checkgroup    basetypes.BoolValue   `tfsdk:"checkgroup"`
	ServerProfile basetypes.StringValue `tfsdk:"server_profile"`
}

// AuthenticationProfilesMethodSamlIdp represents a nested structure within the AuthenticationProfiles model
type AuthenticationProfilesMethodSamlIdp struct {
	AttributeNameUsergroup    basetypes.StringValue `tfsdk:"attribute_name_usergroup"`
	AttributeNameUsername     basetypes.StringValue `tfsdk:"attribute_name_username"`
	CertificateProfile        basetypes.StringValue `tfsdk:"certificate_profile"`
	EnableSingleLogout        basetypes.BoolValue   `tfsdk:"enable_single_logout"`
	RequestSigningCertificate basetypes.StringValue `tfsdk:"request_signing_certificate"`
	ServerProfile             basetypes.StringValue `tfsdk:"server_profile"`
}

// AuthenticationProfilesMultiFactorAuth represents a nested structure within the AuthenticationProfiles model
type AuthenticationProfilesMultiFactorAuth struct {
	Factors   basetypes.ListValue `tfsdk:"factors"`
	MfaEnable basetypes.BoolValue `tfsdk:"mfa_enable"`
}

// AuthenticationProfilesSingleSignOn represents a nested structure within the AuthenticationProfiles model
type AuthenticationProfilesSingleSignOn struct {
	KerberosKeytab basetypes.StringValue `tfsdk:"kerberos_keytab"`
	Realm          basetypes.StringValue `tfsdk:"realm"`
}

// AttrTypes defines the attribute types for the AuthenticationProfiles model.
func (o AuthenticationProfiles) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":       basetypes.StringType{},
		"allow_list": basetypes.ListType{ElemType: basetypes.StringType{}},
		"device":     basetypes.StringType{},
		"folder":     basetypes.StringType{},
		"id":         basetypes.StringType{},
		"lockout": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"failed_attempts": basetypes.Int64Type{},
				"lockout_time":    basetypes.Int64Type{},
			},
		},
		"method": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"cloud": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"profile_name": basetypes.StringType{},
					},
				},
				"kerberos": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"realm":          basetypes.StringType{},
						"server_profile": basetypes.StringType{},
					},
				},
				"ldap": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"login_attribute": basetypes.StringType{},
						"passwd_exp_days": basetypes.Int64Type{},
						"server_profile":  basetypes.StringType{},
					},
				},
				"local_database": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"radius": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"checkgroup":     basetypes.BoolType{},
						"server_profile": basetypes.StringType{},
					},
				},
				"saml_idp": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"attribute_name_usergroup":    basetypes.StringType{},
						"attribute_name_username":     basetypes.StringType{},
						"certificate_profile":         basetypes.StringType{},
						"enable_single_logout":        basetypes.BoolType{},
						"request_signing_certificate": basetypes.StringType{},
						"server_profile":              basetypes.StringType{},
					},
				},
				"tacplus": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"checkgroup":     basetypes.BoolType{},
						"server_profile": basetypes.StringType{},
					},
				},
			},
		},
		"multi_factor_auth": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"factors":    basetypes.ListType{ElemType: basetypes.StringType{}},
				"mfa_enable": basetypes.BoolType{},
			},
		},
		"name": basetypes.StringType{},
		"single_sign_on": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"kerberos_keytab": basetypes.StringType{},
				"realm":           basetypes.StringType{},
			},
		},
		"snippet":           basetypes.StringType{},
		"user_domain":       basetypes.StringType{},
		"username_modifier": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of AuthenticationProfiles objects.
func (o AuthenticationProfiles) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the AuthenticationProfilesLockout model.
func (o AuthenticationProfilesLockout) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"failed_attempts": basetypes.Int64Type{},
		"lockout_time":    basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of AuthenticationProfilesLockout objects.
func (o AuthenticationProfilesLockout) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the AuthenticationProfilesMethod model.
func (o AuthenticationProfilesMethod) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"cloud": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"profile_name": basetypes.StringType{},
			},
		},
		"kerberos": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"realm":          basetypes.StringType{},
				"server_profile": basetypes.StringType{},
			},
		},
		"ldap": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"login_attribute": basetypes.StringType{},
				"passwd_exp_days": basetypes.Int64Type{},
				"server_profile":  basetypes.StringType{},
			},
		},
		"local_database": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
		"radius": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"checkgroup":     basetypes.BoolType{},
				"server_profile": basetypes.StringType{},
			},
		},
		"saml_idp": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"attribute_name_usergroup":    basetypes.StringType{},
				"attribute_name_username":     basetypes.StringType{},
				"certificate_profile":         basetypes.StringType{},
				"enable_single_logout":        basetypes.BoolType{},
				"request_signing_certificate": basetypes.StringType{},
				"server_profile":              basetypes.StringType{},
			},
		},
		"tacplus": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"checkgroup":     basetypes.BoolType{},
				"server_profile": basetypes.StringType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of AuthenticationProfilesMethod objects.
func (o AuthenticationProfilesMethod) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the AuthenticationProfilesMethodCloud model.
func (o AuthenticationProfilesMethodCloud) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"profile_name": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of AuthenticationProfilesMethodCloud objects.
func (o AuthenticationProfilesMethodCloud) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the AuthenticationProfilesMethodKerberos model.
func (o AuthenticationProfilesMethodKerberos) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"realm":          basetypes.StringType{},
		"server_profile": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of AuthenticationProfilesMethodKerberos objects.
func (o AuthenticationProfilesMethodKerberos) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the AuthenticationProfilesMethodLdap model.
func (o AuthenticationProfilesMethodLdap) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"login_attribute": basetypes.StringType{},
		"passwd_exp_days": basetypes.Int64Type{},
		"server_profile":  basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of AuthenticationProfilesMethodLdap objects.
func (o AuthenticationProfilesMethodLdap) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the AuthenticationProfilesMethodRadius model.
func (o AuthenticationProfilesMethodRadius) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"checkgroup":     basetypes.BoolType{},
		"server_profile": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of AuthenticationProfilesMethodRadius objects.
func (o AuthenticationProfilesMethodRadius) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the AuthenticationProfilesMethodSamlIdp model.
func (o AuthenticationProfilesMethodSamlIdp) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"attribute_name_usergroup":    basetypes.StringType{},
		"attribute_name_username":     basetypes.StringType{},
		"certificate_profile":         basetypes.StringType{},
		"enable_single_logout":        basetypes.BoolType{},
		"request_signing_certificate": basetypes.StringType{},
		"server_profile":              basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of AuthenticationProfilesMethodSamlIdp objects.
func (o AuthenticationProfilesMethodSamlIdp) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the AuthenticationProfilesMultiFactorAuth model.
func (o AuthenticationProfilesMultiFactorAuth) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"factors":    basetypes.ListType{ElemType: basetypes.StringType{}},
		"mfa_enable": basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of AuthenticationProfilesMultiFactorAuth objects.
func (o AuthenticationProfilesMultiFactorAuth) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the AuthenticationProfilesSingleSignOn model.
func (o AuthenticationProfilesSingleSignOn) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"kerberos_keytab": basetypes.StringType{},
		"realm":           basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of AuthenticationProfilesSingleSignOn objects.
func (o AuthenticationProfilesSingleSignOn) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AuthenticationProfilesResourceSchema defines the schema for AuthenticationProfiles resource
var AuthenticationProfilesResourceSchema = schema.Schema{
	MarkdownDescription: "AuthenticationProfile resource",
	Attributes: map[string]schema.Attribute{
		"allow_list": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Allow list",
			Optional:            true,
			Computed:            true,
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
			MarkdownDescription: "The UUID of the authentication profile",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"lockout": schema.SingleNestedAttribute{
			MarkdownDescription: "Lockout",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"failed_attempts": schema.Int64Attribute{
					Validators: []validator.Int64{
						int64validator.Between(0, 10),
					},
					MarkdownDescription: "Failed attempts",
					Optional:            true,
				},
				"lockout_time": schema.Int64Attribute{
					Validators: []validator.Int64{
						int64validator.Between(0, 60),
					},
					MarkdownDescription: "Lockout time",
					Optional:            true,
				},
			},
		},
		"method": schema.SingleNestedAttribute{
			MarkdownDescription: "Method",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"cloud": schema.SingleNestedAttribute{
					Validators: []validator.Object{
						objectvalidator.ExactlyOneOf(
							path.MatchRelative().AtParent().AtName("local_database"),
							path.MatchRelative().AtParent().AtName("saml_idp"),
							path.MatchRelative().AtParent().AtName("ldap"),
							path.MatchRelative().AtParent().AtName("radius"),
							path.MatchRelative().AtParent().AtName("tacplus"),
							path.MatchRelative().AtParent().AtName("kerberos"),
						),
					},
					MarkdownDescription: "Cloud",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"profile_name": schema.StringAttribute{
							MarkdownDescription: "The tenant profile name",
							Optional:            true,
						},
					},
				},
				"kerberos": schema.SingleNestedAttribute{
					Validators: []validator.Object{
						objectvalidator.ExactlyOneOf(
							path.MatchRelative().AtParent().AtName("local_database"),
							path.MatchRelative().AtParent().AtName("saml_idp"),
							path.MatchRelative().AtParent().AtName("ldap"),
							path.MatchRelative().AtParent().AtName("radius"),
							path.MatchRelative().AtParent().AtName("tacplus"),
							path.MatchRelative().AtParent().AtName("cloud"),
						),
					},
					MarkdownDescription: "Kerberos",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"realm": schema.StringAttribute{
							MarkdownDescription: "Realm",
							Optional:            true,
						},
						"server_profile": schema.StringAttribute{
							MarkdownDescription: "Server profile",
							Optional:            true,
						},
					},
				},
				"ldap": schema.SingleNestedAttribute{
					Validators: []validator.Object{
						objectvalidator.ExactlyOneOf(
							path.MatchRelative().AtParent().AtName("local_database"),
							path.MatchRelative().AtParent().AtName("saml_idp"),
							path.MatchRelative().AtParent().AtName("radius"),
							path.MatchRelative().AtParent().AtName("tacplus"),
							path.MatchRelative().AtParent().AtName("kerberos"),
							path.MatchRelative().AtParent().AtName("cloud"),
						),
					},
					MarkdownDescription: "Ldap",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"login_attribute": schema.StringAttribute{
							MarkdownDescription: "Login attribute",
							Optional:            true,
						},
						"passwd_exp_days": schema.Int64Attribute{
							MarkdownDescription: "Passwd exp days",
							Optional:            true,
						},
						"server_profile": schema.StringAttribute{
							MarkdownDescription: "Server profile",
							Optional:            true,
						},
					},
				},
				"local_database": schema.SingleNestedAttribute{
					Validators: []validator.Object{
						objectvalidator.ExactlyOneOf(
							path.MatchRelative().AtParent().AtName("saml_idp"),
							path.MatchRelative().AtParent().AtName("ldap"),
							path.MatchRelative().AtParent().AtName("radius"),
							path.MatchRelative().AtParent().AtName("tacplus"),
							path.MatchRelative().AtParent().AtName("kerberos"),
							path.MatchRelative().AtParent().AtName("cloud"),
						),
					},
					MarkdownDescription: "Local database",
					Optional:            true,
					Attributes:          map[string]schema.Attribute{},
				},
				"radius": schema.SingleNestedAttribute{
					Validators: []validator.Object{
						objectvalidator.ExactlyOneOf(
							path.MatchRelative().AtParent().AtName("local_database"),
							path.MatchRelative().AtParent().AtName("saml_idp"),
							path.MatchRelative().AtParent().AtName("ldap"),
							path.MatchRelative().AtParent().AtName("tacplus"),
							path.MatchRelative().AtParent().AtName("kerberos"),
							path.MatchRelative().AtParent().AtName("cloud"),
						),
					},
					MarkdownDescription: "Radius",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"checkgroup": schema.BoolAttribute{
							MarkdownDescription: "Checkgroup",
							Optional:            true,
						},
						"server_profile": schema.StringAttribute{
							MarkdownDescription: "Server profile",
							Optional:            true,
						},
					},
				},
				"saml_idp": schema.SingleNestedAttribute{
					Validators: []validator.Object{
						objectvalidator.ExactlyOneOf(
							path.MatchRelative().AtParent().AtName("local_database"),
							path.MatchRelative().AtParent().AtName("ldap"),
							path.MatchRelative().AtParent().AtName("radius"),
							path.MatchRelative().AtParent().AtName("tacplus"),
							path.MatchRelative().AtParent().AtName("kerberos"),
							path.MatchRelative().AtParent().AtName("cloud"),
						),
					},
					MarkdownDescription: "Saml idp",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"attribute_name_usergroup": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.LengthAtMost(63),
								stringvalidator.LengthAtLeast(1),
							},
							MarkdownDescription: "Attribute name usergroup",
							Optional:            true,
						},
						"attribute_name_username": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.LengthAtMost(63),
								stringvalidator.LengthAtLeast(1),
							},
							MarkdownDescription: "Attribute name username",
							Optional:            true,
						},
						"certificate_profile": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.LengthAtMost(31),
							},
							MarkdownDescription: "Certificate profile",
							Optional:            true,
						},
						"enable_single_logout": schema.BoolAttribute{
							MarkdownDescription: "Enable single logout",
							Optional:            true,
						},
						"request_signing_certificate": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.LengthAtMost(64),
							},
							MarkdownDescription: "Request signing certificate",
							Optional:            true,
						},
						"server_profile": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.LengthAtMost(63),
							},
							MarkdownDescription: "Server profile",
							Optional:            true,
						},
					},
				},
				"tacplus": schema.SingleNestedAttribute{
					Validators: []validator.Object{
						objectvalidator.ExactlyOneOf(
							path.MatchRelative().AtParent().AtName("local_database"),
							path.MatchRelative().AtParent().AtName("saml_idp"),
							path.MatchRelative().AtParent().AtName("ldap"),
							path.MatchRelative().AtParent().AtName("radius"),
							path.MatchRelative().AtParent().AtName("kerberos"),
							path.MatchRelative().AtParent().AtName("cloud"),
						),
					},
					MarkdownDescription: "Tacplus",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"checkgroup": schema.BoolAttribute{
							MarkdownDescription: "Checkgroup",
							Optional:            true,
						},
						"server_profile": schema.StringAttribute{
							MarkdownDescription: "Server profile",
							Optional:            true,
						},
					},
				},
			},
		},
		"multi_factor_auth": schema.SingleNestedAttribute{
			MarkdownDescription: "Multi factor auth",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"factors": schema.ListAttribute{
					ElementType:         types.StringType,
					MarkdownDescription: "Factors",
					Optional:            true,
				},
				"mfa_enable": schema.BoolAttribute{
					MarkdownDescription: "Mfa enable",
					Optional:            true,
				},
			},
		},
		"name": schema.StringAttribute{
			MarkdownDescription: "The name of the authentication profile",
			Required:            true,
		},
		"single_sign_on": schema.SingleNestedAttribute{
			MarkdownDescription: "Single sign on",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"kerberos_keytab": schema.StringAttribute{
					Validators: []validator.String{
						stringvalidator.LengthAtMost(8192),
					},
					MarkdownDescription: "Kerberos keytab",
					Optional:            true,
				},
				"realm": schema.StringAttribute{
					Validators: []validator.String{
						stringvalidator.LengthAtMost(127),
					},
					MarkdownDescription: "Realm",
					Optional:            true,
				},
			},
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
		"user_domain": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(63),
			},
			MarkdownDescription: "User domain",
			Optional:            true,
		},
		"username_modifier": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.OneOf("%USERINPUT%", "%USERINPUT%@%USERDOMAIN%", "%USERDOMAIN%\\\\%USERINPUT%"),
			},
			MarkdownDescription: "Username modifier",
			Optional:            true,
		},
	},
}

// AuthenticationProfilesDataSourceSchema defines the schema for AuthenticationProfiles data source
var AuthenticationProfilesDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "AuthenticationProfile data source",
	Attributes: map[string]dsschema.Attribute{
		"allow_list": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Allow list",
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
			MarkdownDescription: "The UUID of the authentication profile",
			Required:            true,
		},
		"lockout": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Lockout",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"failed_attempts": dsschema.Int64Attribute{
					MarkdownDescription: "Failed attempts",
					Computed:            true,
				},
				"lockout_time": dsschema.Int64Attribute{
					MarkdownDescription: "Lockout time",
					Computed:            true,
				},
			},
		},
		"method": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Method",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"cloud": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Cloud",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"profile_name": dsschema.StringAttribute{
							MarkdownDescription: "The tenant profile name",
							Computed:            true,
						},
					},
				},
				"kerberos": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Kerberos",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"realm": dsschema.StringAttribute{
							MarkdownDescription: "Realm",
							Computed:            true,
						},
						"server_profile": dsschema.StringAttribute{
							MarkdownDescription: "Server profile",
							Computed:            true,
						},
					},
				},
				"ldap": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Ldap",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"login_attribute": dsschema.StringAttribute{
							MarkdownDescription: "Login attribute",
							Computed:            true,
						},
						"passwd_exp_days": dsschema.Int64Attribute{
							MarkdownDescription: "Passwd exp days",
							Computed:            true,
						},
						"server_profile": dsschema.StringAttribute{
							MarkdownDescription: "Server profile",
							Computed:            true,
						},
					},
				},
				"local_database": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Local database",
					Computed:            true,
					Attributes:          map[string]dsschema.Attribute{},
				},
				"radius": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Radius",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"checkgroup": dsschema.BoolAttribute{
							MarkdownDescription: "Checkgroup",
							Computed:            true,
						},
						"server_profile": dsschema.StringAttribute{
							MarkdownDescription: "Server profile",
							Computed:            true,
						},
					},
				},
				"saml_idp": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Saml idp",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"attribute_name_usergroup": dsschema.StringAttribute{
							MarkdownDescription: "Attribute name usergroup",
							Computed:            true,
						},
						"attribute_name_username": dsschema.StringAttribute{
							MarkdownDescription: "Attribute name username",
							Computed:            true,
						},
						"certificate_profile": dsschema.StringAttribute{
							MarkdownDescription: "Certificate profile",
							Computed:            true,
						},
						"enable_single_logout": dsschema.BoolAttribute{
							MarkdownDescription: "Enable single logout",
							Computed:            true,
						},
						"request_signing_certificate": dsschema.StringAttribute{
							MarkdownDescription: "Request signing certificate",
							Computed:            true,
						},
						"server_profile": dsschema.StringAttribute{
							MarkdownDescription: "Server profile",
							Computed:            true,
						},
					},
				},
				"tacplus": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Tacplus",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"checkgroup": dsschema.BoolAttribute{
							MarkdownDescription: "Checkgroup",
							Computed:            true,
						},
						"server_profile": dsschema.StringAttribute{
							MarkdownDescription: "Server profile",
							Computed:            true,
						},
					},
				},
			},
		},
		"multi_factor_auth": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Multi factor auth",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"factors": dsschema.ListAttribute{
					ElementType:         types.StringType,
					MarkdownDescription: "Factors",
					Computed:            true,
				},
				"mfa_enable": dsschema.BoolAttribute{
					MarkdownDescription: "Mfa enable",
					Computed:            true,
				},
			},
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "The name of the authentication profile",
			Optional:            true,
			Computed:            true,
		},
		"single_sign_on": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Single sign on",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"kerberos_keytab": dsschema.StringAttribute{
					MarkdownDescription: "Kerberos keytab",
					Computed:            true,
				},
				"realm": dsschema.StringAttribute{
					MarkdownDescription: "Realm",
					Computed:            true,
				},
			},
		},
		"snippet": dsschema.StringAttribute{
			MarkdownDescription: "The snippet in which the resource is defined",
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
		"user_domain": dsschema.StringAttribute{
			MarkdownDescription: "User domain",
			Computed:            true,
		},
		"username_modifier": dsschema.StringAttribute{
			MarkdownDescription: "Username modifier",
			Computed:            true,
		},
	},
}

// AuthenticationProfilesListModel represents the data model for a list data source.
type AuthenticationProfilesListModel struct {
	Tfid    types.String             `tfsdk:"tfid"`
	Data    []AuthenticationProfiles `tfsdk:"data"`
	Limit   types.Int64              `tfsdk:"limit"`
	Offset  types.Int64              `tfsdk:"offset"`
	Name    types.String             `tfsdk:"name"`
	Total   types.Int64              `tfsdk:"total"`
	Folder  types.String             `tfsdk:"folder"`
	Snippet types.String             `tfsdk:"snippet"`
	Device  types.String             `tfsdk:"device"`
}

// AuthenticationProfilesListDataSourceSchema defines the schema for a list data source.
var AuthenticationProfilesListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: AuthenticationProfilesDataSourceSchema.Attributes,
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
