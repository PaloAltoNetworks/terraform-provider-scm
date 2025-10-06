package models

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	dsschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Package: security_services
// This file contains models for the security_services SDK package

// SecurityRules represents the Terraform model for SecurityRules
type SecurityRules struct {
	Tfid                   types.String          `tfsdk:"tfid"`
	Action                 basetypes.StringValue `tfsdk:"action"`
	AllowUrlCategory       basetypes.ListValue   `tfsdk:"allow_url_category"`
	AllowWebApplication    basetypes.ListValue   `tfsdk:"allow_web_application"`
	Application            basetypes.ListValue   `tfsdk:"application"`
	BlockUrlCategory       basetypes.ListValue   `tfsdk:"block_url_category"`
	BlockWebApplication    basetypes.ListValue   `tfsdk:"block_web_application"`
	Category               basetypes.ListValue   `tfsdk:"category"`
	DefaultProfileSettings basetypes.ObjectValue `tfsdk:"default_profile_settings"`
	Description            basetypes.StringValue `tfsdk:"description"`
	Destination            basetypes.ListValue   `tfsdk:"destination"`
	DestinationHip         basetypes.ListValue   `tfsdk:"destination_hip"`
	Device                 basetypes.StringValue `tfsdk:"device"`
	Devices                basetypes.ListValue   `tfsdk:"devices"`
	Disabled               basetypes.BoolValue   `tfsdk:"disabled"`
	Folder                 basetypes.StringValue `tfsdk:"folder"`
	From                   basetypes.ListValue   `tfsdk:"from"`
	Id                     basetypes.StringValue `tfsdk:"id"`
	LogEnd                 basetypes.BoolValue   `tfsdk:"log_end"`
	LogSetting             basetypes.StringValue `tfsdk:"log_setting"`
	LogSettings            basetypes.ObjectValue `tfsdk:"log_settings"`
	LogStart               basetypes.BoolValue   `tfsdk:"log_start"`
	Name                   basetypes.StringValue `tfsdk:"name"`
	NegateDestination      basetypes.BoolValue   `tfsdk:"negate_destination"`
	NegateSource           basetypes.BoolValue   `tfsdk:"negate_source"`
	NegateUser             basetypes.BoolValue   `tfsdk:"negate_user"`
	PolicyType             basetypes.StringValue `tfsdk:"policy_type"`
	ProfileSetting         basetypes.ObjectValue `tfsdk:"profile_setting"`
	Schedule               basetypes.StringValue `tfsdk:"schedule"`
	SecuritySettings       basetypes.ObjectValue `tfsdk:"security_settings"`
	Service                basetypes.ListValue   `tfsdk:"service"`
	Snippet                basetypes.StringValue `tfsdk:"snippet"`
	Source                 basetypes.ListValue   `tfsdk:"source"`
	SourceHip              basetypes.ListValue   `tfsdk:"source_hip"`
	SourceUser             basetypes.ListValue   `tfsdk:"source_user"`
	Tag                    basetypes.ListValue   `tfsdk:"tag"`
	TenantRestrictions     basetypes.ListValue   `tfsdk:"tenant_restrictions"`
	To                     basetypes.ListValue   `tfsdk:"to"`
	Position               basetypes.StringValue `tfsdk:"position"`
}

// InternetRuleTypeAllowUrlCategoryInner represents a nested structure within the SecurityRules model
type InternetRuleTypeAllowUrlCategoryInner struct {
	AdditionalAction      basetypes.StringValue `tfsdk:"additional_action"`
	CredentialEnforcement basetypes.StringValue `tfsdk:"credential_enforcement"`
	Decryption            basetypes.StringValue `tfsdk:"decryption"`
	Dlp                   basetypes.StringValue `tfsdk:"dlp"`
	FileControl           basetypes.ObjectValue `tfsdk:"file_control"`
	IsolationProfiles     basetypes.StringValue `tfsdk:"isolation_profiles"`
	Name                  basetypes.StringValue `tfsdk:"name"`
}

// InternetRuleTypeAllowUrlCategoryInnerFileControl represents a nested structure within the SecurityRules model
type InternetRuleTypeAllowUrlCategoryInnerFileControl struct {
	Download basetypes.StringValue `tfsdk:"download"`
	Upload   basetypes.StringValue `tfsdk:"upload"`
}

// InternetRuleTypeAllowWebApplicationInner represents a nested structure within the SecurityRules model
type InternetRuleTypeAllowWebApplicationInner struct {
	ApplicationFunction   basetypes.ListValue   `tfsdk:"application_function"`
	Dlp                   basetypes.StringValue `tfsdk:"dlp"`
	FileControl           basetypes.ObjectValue `tfsdk:"file_control"`
	Name                  basetypes.StringValue `tfsdk:"name"`
	SaasEnterpriseControl basetypes.ObjectValue `tfsdk:"saas_enterprise_control"`
	SaasTenantList        basetypes.ListValue   `tfsdk:"saas_tenant_list"`
	SaasUserList          basetypes.ListValue   `tfsdk:"saas_user_list"`
	TenantControl         basetypes.ObjectValue `tfsdk:"tenant_control"`
	Type                  basetypes.StringValue `tfsdk:"type"`
}

// InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControl represents a nested structure within the SecurityRules model
type InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControl struct {
	ConsumerAccess   basetypes.ObjectValue `tfsdk:"consumer_access"`
	EnterpriseAccess basetypes.ObjectValue `tfsdk:"enterprise_access"`
}

// InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlConsumerAccess represents a nested structure within the SecurityRules model
type InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlConsumerAccess struct {
	Enable basetypes.StringValue `tfsdk:"enable"`
}

// InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlEnterpriseAccess represents a nested structure within the SecurityRules model
type InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlEnterpriseAccess struct {
	Enable             basetypes.StringValue `tfsdk:"enable"`
	TenantRestrictions basetypes.ListValue   `tfsdk:"tenant_restrictions"`
}

// InternetRuleTypeAllowWebApplicationInnerTenantControl represents a nested structure within the SecurityRules model
type InternetRuleTypeAllowWebApplicationInnerTenantControl struct {
	AllowedActivities basetypes.ListValue   `tfsdk:"allowed_activities"`
	BlockedActivities basetypes.ListValue   `tfsdk:"blocked_activities"`
	ParentApplication basetypes.StringValue `tfsdk:"parent_application"`
	Tenants           basetypes.ListValue   `tfsdk:"tenants"`
}

// InternetRuleTypeDefaultProfileSettings represents a nested structure within the SecurityRules model
type InternetRuleTypeDefaultProfileSettings struct {
	Dlp         basetypes.StringValue `tfsdk:"dlp"`
	FileControl basetypes.ObjectValue `tfsdk:"file_control"`
}

// InternetRuleTypeLogSettings represents a nested structure within the SecurityRules model
type InternetRuleTypeLogSettings struct {
	LogSessions basetypes.BoolValue `tfsdk:"log_sessions"`
}

// SecurityRuleTypeProfileSetting represents a nested structure within the SecurityRules model
type SecurityRuleTypeProfileSetting struct {
	Group basetypes.ListValue `tfsdk:"group"`
}

// InternetRuleTypeSecuritySettings represents a nested structure within the SecurityRules model
type InternetRuleTypeSecuritySettings struct {
	AntiSpyware              basetypes.StringValue `tfsdk:"anti_spyware"`
	VirusAndWildfireAnalysis basetypes.StringValue `tfsdk:"virus_and_wildfire_analysis"`
	Vulnerability            basetypes.StringValue `tfsdk:"vulnerability"`
}

// AttrTypes defines the attribute types for the SecurityRules model.
func (o SecurityRules) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":   basetypes.StringType{},
		"action": basetypes.StringType{},
		"allow_url_category": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"additional_action":      basetypes.StringType{},
				"credential_enforcement": basetypes.StringType{},
				"decryption":             basetypes.StringType{},
				"dlp":                    basetypes.StringType{},
				"file_control": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"download": basetypes.StringType{},
						"upload":   basetypes.StringType{},
					},
				},
				"isolation_profiles": basetypes.StringType{},
				"name":               basetypes.StringType{},
			},
		}},
		"allow_web_application": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"application_function": basetypes.ListType{ElemType: basetypes.StringType{}},
				"dlp":                  basetypes.StringType{},
				"file_control": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"download": basetypes.StringType{},
						"upload":   basetypes.StringType{},
					},
				},
				"name": basetypes.StringType{},
				"saas_enterprise_control": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"consumer_access": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"enable": basetypes.StringType{},
							},
						},
						"enterprise_access": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"enable":              basetypes.StringType{},
								"tenant_restrictions": basetypes.ListType{ElemType: basetypes.StringType{}},
							},
						},
					},
				},
				"saas_tenant_list": basetypes.ListType{ElemType: basetypes.StringType{}},
				"saas_user_list":   basetypes.ListType{ElemType: basetypes.StringType{}},
				"tenant_control": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"allowed_activities": basetypes.ListType{ElemType: basetypes.StringType{}},
						"blocked_activities": basetypes.ListType{ElemType: basetypes.StringType{}},
						"parent_application": basetypes.StringType{},
						"tenants":            basetypes.ListType{ElemType: basetypes.StringType{}},
					},
				},
				"type": basetypes.StringType{},
			},
		}},
		"application":           basetypes.ListType{ElemType: basetypes.StringType{}},
		"block_url_category":    basetypes.ListType{ElemType: basetypes.StringType{}},
		"block_web_application": basetypes.ListType{ElemType: basetypes.StringType{}},
		"category":              basetypes.ListType{ElemType: basetypes.StringType{}},
		"default_profile_settings": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"dlp": basetypes.StringType{},
				"file_control": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"download": basetypes.StringType{},
						"upload":   basetypes.StringType{},
					},
				},
			},
		},
		"description":     basetypes.StringType{},
		"destination":     basetypes.ListType{ElemType: basetypes.StringType{}},
		"destination_hip": basetypes.ListType{ElemType: basetypes.StringType{}},
		"device":          basetypes.StringType{},
		"devices":         basetypes.ListType{ElemType: basetypes.StringType{}},
		"disabled":        basetypes.BoolType{},
		"folder":          basetypes.StringType{},
		"from":            basetypes.ListType{ElemType: basetypes.StringType{}},
		"id":              basetypes.StringType{},
		"log_end":         basetypes.BoolType{},
		"log_setting":     basetypes.StringType{},
		"log_settings": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"log_sessions": basetypes.BoolType{},
			},
		},
		"log_start":          basetypes.BoolType{},
		"name":               basetypes.StringType{},
		"negate_destination": basetypes.BoolType{},
		"negate_source":      basetypes.BoolType{},
		"negate_user":        basetypes.BoolType{},
		"policy_type":        basetypes.StringType{},
		"profile_setting": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"group": basetypes.ListType{ElemType: basetypes.StringType{}},
			},
		},
		"schedule": basetypes.StringType{},
		"security_settings": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"anti_spyware":                basetypes.StringType{},
				"virus_and_wildfire_analysis": basetypes.StringType{},
				"vulnerability":               basetypes.StringType{},
			},
		},
		"service":             basetypes.ListType{ElemType: basetypes.StringType{}},
		"snippet":             basetypes.StringType{},
		"source":              basetypes.ListType{ElemType: basetypes.StringType{}},
		"source_hip":          basetypes.ListType{ElemType: basetypes.StringType{}},
		"source_user":         basetypes.ListType{ElemType: basetypes.StringType{}},
		"tag":                 basetypes.ListType{ElemType: basetypes.StringType{}},
		"tenant_restrictions": basetypes.ListType{ElemType: basetypes.StringType{}},
		"to":                  basetypes.ListType{ElemType: basetypes.StringType{}},
		"position":            basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of SecurityRules objects.
func (o SecurityRules) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the InternetRuleTypeAllowUrlCategoryInner model.
func (o InternetRuleTypeAllowUrlCategoryInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"additional_action":      basetypes.StringType{},
		"credential_enforcement": basetypes.StringType{},
		"decryption":             basetypes.StringType{},
		"dlp":                    basetypes.StringType{},
		"file_control": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"download": basetypes.StringType{},
				"upload":   basetypes.StringType{},
			},
		},
		"isolation_profiles": basetypes.StringType{},
		"name":               basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of InternetRuleTypeAllowUrlCategoryInner objects.
func (o InternetRuleTypeAllowUrlCategoryInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the InternetRuleTypeAllowUrlCategoryInnerFileControl model.
func (o InternetRuleTypeAllowUrlCategoryInnerFileControl) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"download": basetypes.StringType{},
		"upload":   basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of InternetRuleTypeAllowUrlCategoryInnerFileControl objects.
func (o InternetRuleTypeAllowUrlCategoryInnerFileControl) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the InternetRuleTypeAllowWebApplicationInner model.
func (o InternetRuleTypeAllowWebApplicationInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"application_function": basetypes.ListType{ElemType: basetypes.StringType{}},
		"dlp":                  basetypes.StringType{},
		"file_control": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"download": basetypes.StringType{},
				"upload":   basetypes.StringType{},
			},
		},
		"name": basetypes.StringType{},
		"saas_enterprise_control": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"consumer_access": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"enable": basetypes.StringType{},
					},
				},
				"enterprise_access": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"enable":              basetypes.StringType{},
						"tenant_restrictions": basetypes.ListType{ElemType: basetypes.StringType{}},
					},
				},
			},
		},
		"saas_tenant_list": basetypes.ListType{ElemType: basetypes.StringType{}},
		"saas_user_list":   basetypes.ListType{ElemType: basetypes.StringType{}},
		"tenant_control": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"allowed_activities": basetypes.ListType{ElemType: basetypes.StringType{}},
				"blocked_activities": basetypes.ListType{ElemType: basetypes.StringType{}},
				"parent_application": basetypes.StringType{},
				"tenants":            basetypes.ListType{ElemType: basetypes.StringType{}},
			},
		},
		"type": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of InternetRuleTypeAllowWebApplicationInner objects.
func (o InternetRuleTypeAllowWebApplicationInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControl model.
func (o InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControl) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"consumer_access": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"enable": basetypes.StringType{},
			},
		},
		"enterprise_access": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"enable":              basetypes.StringType{},
				"tenant_restrictions": basetypes.ListType{ElemType: basetypes.StringType{}},
			},
		},
	}
}

// AttrType returns the attribute type for a list of InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControl objects.
func (o InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControl) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlConsumerAccess model.
func (o InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlConsumerAccess) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"enable": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlConsumerAccess objects.
func (o InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlConsumerAccess) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlEnterpriseAccess model.
func (o InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlEnterpriseAccess) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"enable":              basetypes.StringType{},
		"tenant_restrictions": basetypes.ListType{ElemType: basetypes.StringType{}},
	}
}

// AttrType returns the attribute type for a list of InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlEnterpriseAccess objects.
func (o InternetRuleTypeAllowWebApplicationInnerSaasEnterpriseControlEnterpriseAccess) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the InternetRuleTypeAllowWebApplicationInnerTenantControl model.
func (o InternetRuleTypeAllowWebApplicationInnerTenantControl) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"allowed_activities": basetypes.ListType{ElemType: basetypes.StringType{}},
		"blocked_activities": basetypes.ListType{ElemType: basetypes.StringType{}},
		"parent_application": basetypes.StringType{},
		"tenants":            basetypes.ListType{ElemType: basetypes.StringType{}},
	}
}

// AttrType returns the attribute type for a list of InternetRuleTypeAllowWebApplicationInnerTenantControl objects.
func (o InternetRuleTypeAllowWebApplicationInnerTenantControl) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the InternetRuleTypeDefaultProfileSettings model.
func (o InternetRuleTypeDefaultProfileSettings) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"dlp": basetypes.StringType{},
		"file_control": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"download": basetypes.StringType{},
				"upload":   basetypes.StringType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of InternetRuleTypeDefaultProfileSettings objects.
func (o InternetRuleTypeDefaultProfileSettings) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the InternetRuleTypeLogSettings model.
func (o InternetRuleTypeLogSettings) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"log_sessions": basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of InternetRuleTypeLogSettings objects.
func (o InternetRuleTypeLogSettings) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the SecurityRuleTypeProfileSetting model.
func (o SecurityRuleTypeProfileSetting) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"group": basetypes.ListType{ElemType: basetypes.StringType{}},
	}
}

// AttrType returns the attribute type for a list of SecurityRuleTypeProfileSetting objects.
func (o SecurityRuleTypeProfileSetting) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the InternetRuleTypeSecuritySettings model.
func (o InternetRuleTypeSecuritySettings) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"anti_spyware":                basetypes.StringType{},
		"virus_and_wildfire_analysis": basetypes.StringType{},
		"vulnerability":               basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of InternetRuleTypeSecuritySettings objects.
func (o InternetRuleTypeSecuritySettings) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// SecurityRulesResourceSchema defines the schema for SecurityRules resource
var SecurityRulesResourceSchema = schema.Schema{
	MarkdownDescription: "SecurityRule resource",
	Attributes: map[string]schema.Attribute{
		"action": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.OneOf("allow", "deny", "drop", "reset-client", "reset-server", "reset-both"),
			},
			MarkdownDescription: "The action to be taken when the rule is matched",
			Optional:            true,
		},
		"allow_url_category": schema.ListNestedAttribute{
			MarkdownDescription: "Allow url category",
			Optional:            true,
			Computed:            true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"additional_action": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("none", "continue", "redirect", "isolate"),
						},
						MarkdownDescription: "Additional action",
						Optional:            true,
						Computed:            true,
						Default:             stringdefault.StaticString("none"),
					},
					"credential_enforcement": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("enabled", "disabled"),
						},
						MarkdownDescription: "Credential enforcement",
						Optional:            true,
						Computed:            true,
						Default:             stringdefault.StaticString("enabled"),
					},
					"decryption": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("enabled", "disabled"),
						},
						MarkdownDescription: "Decryption",
						Optional:            true,
						Computed:            true,
						Default:             stringdefault.StaticString("enabled"),
					},
					"dlp": schema.StringAttribute{
						MarkdownDescription: "Dlp",
						Optional:            true,
					},
					"file_control": schema.SingleNestedAttribute{
						MarkdownDescription: "File control",
						Optional:            true,
						Attributes: map[string]schema.Attribute{
							"download": schema.StringAttribute{
								Validators: []validator.String{
									stringvalidator.OneOf("allow-all-file-types", "best-practice", "block-all-file-types"),
								},
								MarkdownDescription: "Download",
								Optional:            true,
							},
							"upload": schema.StringAttribute{
								Validators: []validator.String{
									stringvalidator.OneOf("allow-all-file-types", "best-practice", "block-all-file-types"),
								},
								MarkdownDescription: "Upload",
								Optional:            true,
							},
						},
					},
					"isolation_profiles": schema.StringAttribute{
						MarkdownDescription: "Isolation profiles",
						Optional:            true,
						Computed:            true,
						Default:             stringdefault.StaticString("none"),
					},
					"name": schema.StringAttribute{
						MarkdownDescription: "Name",
						Optional:            true,
					},
				},
			},
		},
		"allow_web_application": schema.ListNestedAttribute{
			MarkdownDescription: "Allow web application",
			Optional:            true,
			Computed:            true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"application_function": schema.ListAttribute{
						ElementType:         types.StringType,
						MarkdownDescription: "Application function",
						Optional:            true,
					},
					"dlp": schema.StringAttribute{
						MarkdownDescription: "Dlp",
						Optional:            true,
					},
					"file_control": schema.SingleNestedAttribute{
						MarkdownDescription: "File control",
						Optional:            true,
						Attributes: map[string]schema.Attribute{
							"download": schema.StringAttribute{
								Validators: []validator.String{
									stringvalidator.OneOf("allow-all-file-types", "best-practice", "block-all-file-types"),
								},
								MarkdownDescription: "Download",
								Optional:            true,
							},
							"upload": schema.StringAttribute{
								Validators: []validator.String{
									stringvalidator.OneOf("allow-all-file-types", "best-practice", "block-all-file-types"),
								},
								MarkdownDescription: "Upload",
								Optional:            true,
							},
						},
					},
					"name": schema.StringAttribute{
						MarkdownDescription: "Name",
						Optional:            true,
					},
					"saas_enterprise_control": schema.SingleNestedAttribute{
						MarkdownDescription: "Saas enterprise control",
						Optional:            true,
						Attributes: map[string]schema.Attribute{
							"consumer_access": schema.SingleNestedAttribute{
								MarkdownDescription: "Consumer access",
								Optional:            true,
								Attributes: map[string]schema.Attribute{
									"enable": schema.StringAttribute{
										Validators: []validator.String{
											stringvalidator.OneOf("yes", "no"),
										},
										MarkdownDescription: "Enable",
										Optional:            true,
									},
								},
							},
							"enterprise_access": schema.SingleNestedAttribute{
								MarkdownDescription: "Enterprise access",
								Optional:            true,
								Attributes: map[string]schema.Attribute{
									"enable": schema.StringAttribute{
										Validators: []validator.String{
											stringvalidator.OneOf("yes", "no"),
										},
										MarkdownDescription: "Enable",
										Optional:            true,
									},
									"tenant_restrictions": schema.ListAttribute{
										ElementType:         types.StringType,
										MarkdownDescription: "Tenant restrictions",
										Optional:            true,
									},
								},
							},
						},
					},
					"saas_tenant_list": schema.ListAttribute{
						ElementType:         types.StringType,
						MarkdownDescription: "Saas tenant list",
						Optional:            true,
					},
					"saas_user_list": schema.ListAttribute{
						ElementType:         types.StringType,
						MarkdownDescription: "Saas user list",
						Optional:            true,
					},
					"tenant_control": schema.SingleNestedAttribute{
						MarkdownDescription: "Tenant control",
						Optional:            true,
						Attributes: map[string]schema.Attribute{
							"allowed_activities": schema.ListAttribute{
								ElementType:         types.StringType,
								MarkdownDescription: "Allowed activities",
								Optional:            true,
							},
							"blocked_activities": schema.ListAttribute{
								ElementType:         types.StringType,
								MarkdownDescription: "Blocked activities",
								Optional:            true,
							},
							"parent_application": schema.StringAttribute{
								MarkdownDescription: "Parent application",
								Optional:            true,
							},
							"tenants": schema.ListAttribute{
								ElementType:         types.StringType,
								MarkdownDescription: "Tenants",
								Optional:            true,
							},
						},
					},
					"type": schema.StringAttribute{
						MarkdownDescription: "Type",
						Optional:            true,
					},
				},
			},
		},
		"application": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "The application(s) being accessed",
			Optional:            true,
			Computed:            true,
		},
		"block_url_category": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Block url category",
			Optional:            true,
			Computed:            true,
		},
		"block_web_application": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Block web application",
			Optional:            true,
			Computed:            true,
		},
		"category": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "The URL categories being accessed",
			Optional:            true,
			Computed:            true,
		},
		"default_profile_settings": schema.SingleNestedAttribute{
			MarkdownDescription: "Default profile settings",
			Optional:            true,
			Computed:            true,
			Attributes: map[string]schema.Attribute{
				"dlp": schema.StringAttribute{
					MarkdownDescription: "Dlp",
					Optional:            true,
					Computed:            true,
				},
				"file_control": schema.SingleNestedAttribute{
					MarkdownDescription: "File control",
					Optional:            true,
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"download": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("allow-all-file-types", "best-practice", "block-all-file-types"),
							},
							MarkdownDescription: "Download",
							Optional:            true,
							Computed:            true,
						},
						"upload": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("allow-all-file-types", "best-practice", "block-all-file-types"),
							},
							MarkdownDescription: "Upload",
							Optional:            true,
							Computed:            true,
						},
					},
				},
			},
		},
		"description": schema.StringAttribute{
			MarkdownDescription: "The description of the security rule",
			Optional:            true,
		},
		"destination": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "The destination address(es)",
			Optional:            true,
		},
		"destination_hip": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "The destination Host Integrity Profile(s)",
			Optional:            true,
			Computed:            true,
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
		"devices": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Devices",
			Optional:            true,
			Computed:            true,
		},
		"disabled": schema.BoolAttribute{
			MarkdownDescription: "Is the security rule disabled?",
			Optional:            true,
			Computed:            true,
			Default:             booldefault.StaticBool(false),
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
		"from": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "The source security zone(s)",
			Optional:            true,
		},
		"id": schema.StringAttribute{
			MarkdownDescription: "The UUID of the security rule",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"log_end": schema.BoolAttribute{
			MarkdownDescription: "Log at session end?",
			Optional:            true,
			Computed:            true,
		},
		"log_setting": schema.StringAttribute{
			MarkdownDescription: "The external log forwarding profile",
			Optional:            true,
			Computed:            true,
		},
		"log_settings": schema.SingleNestedAttribute{
			MarkdownDescription: "Log settings",
			Optional:            true,
			Computed:            true,
			Attributes: map[string]schema.Attribute{
				"log_sessions": schema.BoolAttribute{
					MarkdownDescription: "Log sessions",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(true),
				},
			},
		},
		"log_start": schema.BoolAttribute{
			MarkdownDescription: "Log at session start?",
			Optional:            true,
			Computed:            true,
		},
		"name": schema.StringAttribute{
			MarkdownDescription: "The name of the security rule",
			Optional:            true,
		},
		"negate_destination": schema.BoolAttribute{
			MarkdownDescription: "Negate the destination addresses(es)?",
			Optional:            true,
			Computed:            true,
		},
		"negate_source": schema.BoolAttribute{
			MarkdownDescription: "Negate the source address(es)?",
			Optional:            true,
			Computed:            true,
			Default:             booldefault.StaticBool(false),
		},
		"negate_user": schema.BoolAttribute{
			MarkdownDescription: "Negate user",
			Optional:            true,
			Computed:            true,
		},
		"policy_type": schema.StringAttribute{
			MarkdownDescription: "Policy type",
			Optional:            true,
			Computed:            true,
			Default:             stringdefault.StaticString("Security"),
		},
		"position": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.OneOf("pre", "post"),
			},
			MarkdownDescription: "The position of a security rule\n",
			Optional:            true,
			Computed:            true,
			Default:             stringdefault.StaticString("pre"),
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"profile_setting": schema.SingleNestedAttribute{
			MarkdownDescription: "The security profile object",
			Optional:            true,
			Computed:            true,
			Attributes: map[string]schema.Attribute{
				"group": schema.ListAttribute{
					ElementType:         types.StringType,
					MarkdownDescription: "The security profile group",
					Optional:            true,
					Computed:            true,
				},
			},
		},
		"schedule": schema.StringAttribute{
			MarkdownDescription: "Schedule in which this rule will be applied",
			Optional:            true,
		},
		"security_settings": schema.SingleNestedAttribute{
			MarkdownDescription: "Security settings",
			Optional:            true,
			Computed:            true,
			Attributes: map[string]schema.Attribute{
				"anti_spyware": schema.StringAttribute{
					Validators: []validator.String{
						stringvalidator.OneOf("yes", "no"),
					},
					MarkdownDescription: "Anti spyware",
					Optional:            true,
					Computed:            true,
					Default:             stringdefault.StaticString("yes"),
				},
				"virus_and_wildfire_analysis": schema.StringAttribute{
					Validators: []validator.String{
						stringvalidator.OneOf("yes", "no"),
					},
					MarkdownDescription: "Virus and wildfire analysis",
					Optional:            true,
					Computed:            true,
					Default:             stringdefault.StaticString("yes"),
				},
				"vulnerability": schema.StringAttribute{
					Validators: []validator.String{
						stringvalidator.OneOf("yes", "no"),
					},
					MarkdownDescription: "Vulnerability",
					Optional:            true,
					Computed:            true,
					Default:             stringdefault.StaticString("yes"),
				},
			},
		},
		"service": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "The service(s) being accessed",
			Optional:            true,
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
		"source": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "The source addresses(es)",
			Optional:            true,
		},
		"source_hip": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "The source Host Integrity Profile(s)",
			Optional:            true,
			Computed:            true,
		},
		"source_user": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "List of source users and/or groups.  Reserved words include `any`, `pre-login`, `known-user`, and `unknown`.",
			Optional:            true,
		},
		"tag": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "The tags associated with the security rule",
			Optional:            true,
		},
		"tenant_restrictions": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Tenant restrictions",
			Optional:            true,
			Computed:            true,
		},
		"tfid": schema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"to": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "The destination security zone(s)",
			Optional:            true,
		},
	},
}

// SecurityRulesDataSourceSchema defines the schema for SecurityRules data source
var SecurityRulesDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "SecurityRule data source",
	Attributes: map[string]dsschema.Attribute{
		"action": dsschema.StringAttribute{
			MarkdownDescription: "The action to be taken when the rule is matched",
			Computed:            true,
		},
		"allow_url_category": dsschema.ListNestedAttribute{
			MarkdownDescription: "Allow url category",
			Computed:            true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: map[string]dsschema.Attribute{
					"additional_action": dsschema.StringAttribute{
						MarkdownDescription: "Additional action",
						Computed:            true,
					},
					"credential_enforcement": dsschema.StringAttribute{
						MarkdownDescription: "Credential enforcement",
						Computed:            true,
					},
					"decryption": dsschema.StringAttribute{
						MarkdownDescription: "Decryption",
						Computed:            true,
					},
					"dlp": dsschema.StringAttribute{
						MarkdownDescription: "Dlp",
						Computed:            true,
					},
					"file_control": dsschema.SingleNestedAttribute{
						MarkdownDescription: "File control",
						Computed:            true,
						Attributes: map[string]dsschema.Attribute{
							"download": dsschema.StringAttribute{
								MarkdownDescription: "Download",
								Computed:            true,
							},
							"upload": dsschema.StringAttribute{
								MarkdownDescription: "Upload",
								Computed:            true,
							},
						},
					},
					"isolation_profiles": dsschema.StringAttribute{
						MarkdownDescription: "Isolation profiles",
						Computed:            true,
					},
					"name": dsschema.StringAttribute{
						MarkdownDescription: "Name",
						Computed:            true,
					},
				},
			},
		},
		"allow_web_application": dsschema.ListNestedAttribute{
			MarkdownDescription: "Allow web application",
			Computed:            true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: map[string]dsschema.Attribute{
					"application_function": dsschema.ListAttribute{
						ElementType:         types.StringType,
						MarkdownDescription: "Application function",
						Computed:            true,
					},
					"dlp": dsschema.StringAttribute{
						MarkdownDescription: "Dlp",
						Computed:            true,
					},
					"file_control": dsschema.SingleNestedAttribute{
						MarkdownDescription: "File control",
						Computed:            true,
						Attributes: map[string]dsschema.Attribute{
							"download": dsschema.StringAttribute{
								MarkdownDescription: "Download",
								Computed:            true,
							},
							"upload": dsschema.StringAttribute{
								MarkdownDescription: "Upload",
								Computed:            true,
							},
						},
					},
					"name": dsschema.StringAttribute{
						MarkdownDescription: "Name",
						Computed:            true,
					},
					"saas_enterprise_control": dsschema.SingleNestedAttribute{
						MarkdownDescription: "Saas enterprise control",
						Computed:            true,
						Attributes: map[string]dsschema.Attribute{
							"consumer_access": dsschema.SingleNestedAttribute{
								MarkdownDescription: "Consumer access",
								Computed:            true,
								Attributes: map[string]dsschema.Attribute{
									"enable": dsschema.StringAttribute{
										MarkdownDescription: "Enable",
										Computed:            true,
									},
								},
							},
							"enterprise_access": dsschema.SingleNestedAttribute{
								MarkdownDescription: "Enterprise access",
								Computed:            true,
								Attributes: map[string]dsschema.Attribute{
									"enable": dsschema.StringAttribute{
										MarkdownDescription: "Enable",
										Computed:            true,
									},
									"tenant_restrictions": dsschema.ListAttribute{
										ElementType:         types.StringType,
										MarkdownDescription: "Tenant restrictions",
										Computed:            true,
									},
								},
							},
						},
					},
					"saas_tenant_list": dsschema.ListAttribute{
						ElementType:         types.StringType,
						MarkdownDescription: "Saas tenant list",
						Computed:            true,
					},
					"saas_user_list": dsschema.ListAttribute{
						ElementType:         types.StringType,
						MarkdownDescription: "Saas user list",
						Computed:            true,
					},
					"tenant_control": dsschema.SingleNestedAttribute{
						MarkdownDescription: "Tenant control",
						Computed:            true,
						Attributes: map[string]dsschema.Attribute{
							"allowed_activities": dsschema.ListAttribute{
								ElementType:         types.StringType,
								MarkdownDescription: "Allowed activities",
								Computed:            true,
							},
							"blocked_activities": dsschema.ListAttribute{
								ElementType:         types.StringType,
								MarkdownDescription: "Blocked activities",
								Computed:            true,
							},
							"parent_application": dsschema.StringAttribute{
								MarkdownDescription: "Parent application",
								Computed:            true,
							},
							"tenants": dsschema.ListAttribute{
								ElementType:         types.StringType,
								MarkdownDescription: "Tenants",
								Computed:            true,
							},
						},
					},
					"type": dsschema.StringAttribute{
						MarkdownDescription: "Type",
						Computed:            true,
					},
				},
			},
		},
		"application": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "The application(s) being accessed",
			Computed:            true,
		},
		"block_url_category": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Block url category",
			Computed:            true,
		},
		"block_web_application": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Block web application",
			Computed:            true,
		},
		"category": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "The URL categories being accessed",
			Computed:            true,
		},
		"default_profile_settings": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Default profile settings",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"dlp": dsschema.StringAttribute{
					MarkdownDescription: "Dlp",
					Computed:            true,
				},
				"file_control": dsschema.SingleNestedAttribute{
					MarkdownDescription: "File control",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"download": dsschema.StringAttribute{
							MarkdownDescription: "Download",
							Computed:            true,
						},
						"upload": dsschema.StringAttribute{
							MarkdownDescription: "Upload",
							Computed:            true,
						},
					},
				},
			},
		},
		"description": dsschema.StringAttribute{
			MarkdownDescription: "The description of the security rule",
			Computed:            true,
		},
		"destination": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "The destination address(es)",
			Computed:            true,
		},
		"destination_hip": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "The destination Host Integrity Profile(s)",
			Computed:            true,
		},
		"device": dsschema.StringAttribute{
			MarkdownDescription: "The device in which the resource is defined",
			Computed:            true,
		},
		"devices": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Devices",
			Computed:            true,
		},
		"disabled": dsschema.BoolAttribute{
			MarkdownDescription: "Is the security rule disabled?",
			Computed:            true,
		},
		"folder": dsschema.StringAttribute{
			MarkdownDescription: "The folder in which the resource is defined",
			Computed:            true,
		},
		"from": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "The source security zone(s)",
			Computed:            true,
		},
		"id": dsschema.StringAttribute{
			MarkdownDescription: "The UUID of the security rule",
			Required:            true,
		},
		"log_end": dsschema.BoolAttribute{
			MarkdownDescription: "Log at session end?",
			Computed:            true,
		},
		"log_setting": dsschema.StringAttribute{
			MarkdownDescription: "The external log forwarding profile",
			Computed:            true,
		},
		"log_settings": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Log settings",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"log_sessions": dsschema.BoolAttribute{
					MarkdownDescription: "Log sessions",
					Computed:            true,
				},
			},
		},
		"log_start": dsschema.BoolAttribute{
			MarkdownDescription: "Log at session start?",
			Computed:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "The name of the security rule",
			Optional:            true,
			Computed:            true,
		},
		"negate_destination": dsschema.BoolAttribute{
			MarkdownDescription: "Negate the destination addresses(es)?",
			Computed:            true,
		},
		"negate_source": dsschema.BoolAttribute{
			MarkdownDescription: "Negate the source address(es)?",
			Computed:            true,
		},
		"negate_user": dsschema.BoolAttribute{
			MarkdownDescription: "Negate user",
			Computed:            true,
		},
		"policy_type": dsschema.StringAttribute{
			MarkdownDescription: "Policy type",
			Computed:            true,
		},
		"profile_setting": dsschema.SingleNestedAttribute{
			MarkdownDescription: "The security profile object",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"group": dsschema.ListAttribute{
					ElementType:         types.StringType,
					MarkdownDescription: "The security profile group",
					Computed:            true,
				},
			},
		},
		"schedule": dsschema.StringAttribute{
			MarkdownDescription: "Schedule in which this rule will be applied",
			Computed:            true,
		},
		"security_settings": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Security settings",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"anti_spyware": dsschema.StringAttribute{
					MarkdownDescription: "Anti spyware",
					Computed:            true,
				},
				"virus_and_wildfire_analysis": dsschema.StringAttribute{
					MarkdownDescription: "Virus and wildfire analysis",
					Computed:            true,
				},
				"vulnerability": dsschema.StringAttribute{
					MarkdownDescription: "Vulnerability",
					Computed:            true,
				},
			},
		},
		"service": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "The service(s) being accessed",
			Computed:            true,
		},
		"snippet": dsschema.StringAttribute{
			MarkdownDescription: "The snippet in which the resource is defined",
			Computed:            true,
		},
		"source": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "The source addresses(es)",
			Computed:            true,
		},
		"source_hip": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "The source Host Integrity Profile(s)",
			Computed:            true,
		},
		"source_user": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "List of source users and/or groups.  Reserved words include `any`, `pre-login`, `known-user`, and `unknown`.",
			Computed:            true,
		},
		"tag": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "The tags associated with the security rule",
			Computed:            true,
		},
		"tenant_restrictions": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Tenant restrictions",
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
		"to": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "The destination security zone(s)",
			Computed:            true,
		},
	},
}

// SecurityRulesListModel represents the data model for a list data source.
type SecurityRulesListModel struct {
	Tfid    types.String    `tfsdk:"tfid"`
	Data    []SecurityRules `tfsdk:"data"`
	Limit   types.Int64     `tfsdk:"limit"`
	Offset  types.Int64     `tfsdk:"offset"`
	Name    types.String    `tfsdk:"name"`
	Total   types.Int64     `tfsdk:"total"`
	Folder  types.String    `tfsdk:"folder"`
	Snippet types.String    `tfsdk:"snippet"`
	Device  types.String    `tfsdk:"device"`
}

// SecurityRulesListDataSourceSchema defines the schema for a list data source.
var SecurityRulesListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: SecurityRulesDataSourceSchema.Attributes,
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
