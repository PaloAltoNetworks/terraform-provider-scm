package models

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/objectvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	dsschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Package: objects
// This file contains models for the objects SDK package

// ExternalDynamicLists represents the Terraform model for ExternalDynamicLists
type ExternalDynamicLists struct {
	Tfid            types.String          `tfsdk:"tfid"`
	EncryptedValues basetypes.MapValue    `tfsdk:"encrypted_values"`
	Device          basetypes.StringValue `tfsdk:"device"`
	Folder          basetypes.StringValue `tfsdk:"folder"`
	Id              basetypes.StringValue `tfsdk:"id"`
	Name            basetypes.StringValue `tfsdk:"name"`
	Snippet         basetypes.StringValue `tfsdk:"snippet"`
	Type            basetypes.ObjectValue `tfsdk:"type"`
}

// ExternalDynamicListsType represents a nested structure within the ExternalDynamicLists model
type ExternalDynamicListsType struct {
	Domain        basetypes.ObjectValue `tfsdk:"domain"`
	Imei          basetypes.ObjectValue `tfsdk:"imei"`
	Imsi          basetypes.ObjectValue `tfsdk:"imsi"`
	Ip            basetypes.ObjectValue `tfsdk:"ip"`
	PredefinedIp  basetypes.ObjectValue `tfsdk:"predefined_ip"`
	PredefinedUrl basetypes.ObjectValue `tfsdk:"predefined_url"`
	Url           basetypes.ObjectValue `tfsdk:"url"`
}

// ExternalDynamicListsTypeDomain represents a nested structure within the ExternalDynamicLists model
type ExternalDynamicListsTypeDomain struct {
	Auth               basetypes.ObjectValue `tfsdk:"auth"`
	CertificateProfile basetypes.StringValue `tfsdk:"certificate_profile"`
	Description        basetypes.StringValue `tfsdk:"description"`
	ExceptionList      basetypes.ListValue   `tfsdk:"exception_list"`
	ExpandDomain       basetypes.BoolValue   `tfsdk:"expand_domain"`
	Recurring          basetypes.ObjectValue `tfsdk:"recurring"`
	Url                basetypes.StringValue `tfsdk:"url"`
}

// ExternalDynamicListsTypeDomainAuth represents a nested structure within the ExternalDynamicLists model
type ExternalDynamicListsTypeDomainAuth struct {
	Password basetypes.StringValue `tfsdk:"password"`
	Username basetypes.StringValue `tfsdk:"username"`
}

// ExternalDynamicListsTypeDomainRecurring represents a nested structure within the ExternalDynamicLists model
type ExternalDynamicListsTypeDomainRecurring struct {
	Daily      basetypes.ObjectValue `tfsdk:"daily"`
	FiveMinute basetypes.ObjectValue `tfsdk:"five_minute"`
	Hourly     basetypes.ObjectValue `tfsdk:"hourly"`
	Monthly    basetypes.ObjectValue `tfsdk:"monthly"`
	Weekly     basetypes.ObjectValue `tfsdk:"weekly"`
}

// ExternalDynamicListsTypeDomainRecurringDaily represents a nested structure within the ExternalDynamicLists model
type ExternalDynamicListsTypeDomainRecurringDaily struct {
	At basetypes.StringValue `tfsdk:"at"`
}

// ExternalDynamicListsTypeDomainRecurringMonthly represents a nested structure within the ExternalDynamicLists model
type ExternalDynamicListsTypeDomainRecurringMonthly struct {
	At         basetypes.StringValue `tfsdk:"at"`
	DayOfMonth basetypes.Int64Value  `tfsdk:"day_of_month"`
}

// ExternalDynamicListsTypeDomainRecurringWeekly represents a nested structure within the ExternalDynamicLists model
type ExternalDynamicListsTypeDomainRecurringWeekly struct {
	At        basetypes.StringValue `tfsdk:"at"`
	DayOfWeek basetypes.StringValue `tfsdk:"day_of_week"`
}

// ExternalDynamicListsTypeImei represents a nested structure within the ExternalDynamicLists model
type ExternalDynamicListsTypeImei struct {
	Auth               basetypes.ObjectValue `tfsdk:"auth"`
	CertificateProfile basetypes.StringValue `tfsdk:"certificate_profile"`
	Description        basetypes.StringValue `tfsdk:"description"`
	ExceptionList      basetypes.ListValue   `tfsdk:"exception_list"`
	Recurring          basetypes.ObjectValue `tfsdk:"recurring"`
	Url                basetypes.StringValue `tfsdk:"url"`
}

// ExternalDynamicListsTypeImeiAuth represents a nested structure within the ExternalDynamicLists model
type ExternalDynamicListsTypeImeiAuth struct {
	Password basetypes.StringValue `tfsdk:"password"`
	Username basetypes.StringValue `tfsdk:"username"`
}

// ExternalDynamicListsTypeImeiRecurring represents a nested structure within the ExternalDynamicLists model
type ExternalDynamicListsTypeImeiRecurring struct {
	Daily      basetypes.ObjectValue `tfsdk:"daily"`
	FiveMinute basetypes.ObjectValue `tfsdk:"five_minute"`
	Hourly     basetypes.ObjectValue `tfsdk:"hourly"`
	Monthly    basetypes.ObjectValue `tfsdk:"monthly"`
	Weekly     basetypes.ObjectValue `tfsdk:"weekly"`
}

// ExternalDynamicListsTypeImeiRecurringDaily represents a nested structure within the ExternalDynamicLists model
type ExternalDynamicListsTypeImeiRecurringDaily struct {
	At basetypes.StringValue `tfsdk:"at"`
}

// ExternalDynamicListsTypeImeiRecurringMonthly represents a nested structure within the ExternalDynamicLists model
type ExternalDynamicListsTypeImeiRecurringMonthly struct {
	At         basetypes.StringValue `tfsdk:"at"`
	DayOfMonth basetypes.Int64Value  `tfsdk:"day_of_month"`
}

// ExternalDynamicListsTypeImeiRecurringWeekly represents a nested structure within the ExternalDynamicLists model
type ExternalDynamicListsTypeImeiRecurringWeekly struct {
	At        basetypes.StringValue `tfsdk:"at"`
	DayOfWeek basetypes.StringValue `tfsdk:"day_of_week"`
}

// ExternalDynamicListsTypeImsi represents a nested structure within the ExternalDynamicLists model
type ExternalDynamicListsTypeImsi struct {
	Auth               basetypes.ObjectValue `tfsdk:"auth"`
	CertificateProfile basetypes.StringValue `tfsdk:"certificate_profile"`
	Description        basetypes.StringValue `tfsdk:"description"`
	ExceptionList      basetypes.ListValue   `tfsdk:"exception_list"`
	Recurring          basetypes.ObjectValue `tfsdk:"recurring"`
	Url                basetypes.StringValue `tfsdk:"url"`
}

// ExternalDynamicListsTypeImsiAuth represents a nested structure within the ExternalDynamicLists model
type ExternalDynamicListsTypeImsiAuth struct {
	Password basetypes.StringValue `tfsdk:"password"`
	Username basetypes.StringValue `tfsdk:"username"`
}

// ExternalDynamicListsTypeImsiRecurring represents a nested structure within the ExternalDynamicLists model
type ExternalDynamicListsTypeImsiRecurring struct {
	Daily      basetypes.ObjectValue `tfsdk:"daily"`
	FiveMinute basetypes.ObjectValue `tfsdk:"five_minute"`
	Hourly     basetypes.ObjectValue `tfsdk:"hourly"`
	Monthly    basetypes.ObjectValue `tfsdk:"monthly"`
	Weekly     basetypes.ObjectValue `tfsdk:"weekly"`
}

// ExternalDynamicListsTypeImsiRecurringDaily represents a nested structure within the ExternalDynamicLists model
type ExternalDynamicListsTypeImsiRecurringDaily struct {
	At basetypes.StringValue `tfsdk:"at"`
}

// ExternalDynamicListsTypeImsiRecurringMonthly represents a nested structure within the ExternalDynamicLists model
type ExternalDynamicListsTypeImsiRecurringMonthly struct {
	At         basetypes.StringValue `tfsdk:"at"`
	DayOfMonth basetypes.Int64Value  `tfsdk:"day_of_month"`
}

// ExternalDynamicListsTypeImsiRecurringWeekly represents a nested structure within the ExternalDynamicLists model
type ExternalDynamicListsTypeImsiRecurringWeekly struct {
	At        basetypes.StringValue `tfsdk:"at"`
	DayOfWeek basetypes.StringValue `tfsdk:"day_of_week"`
}

// ExternalDynamicListsTypeIp represents a nested structure within the ExternalDynamicLists model
type ExternalDynamicListsTypeIp struct {
	Auth               basetypes.ObjectValue `tfsdk:"auth"`
	CertificateProfile basetypes.StringValue `tfsdk:"certificate_profile"`
	Description        basetypes.StringValue `tfsdk:"description"`
	ExceptionList      basetypes.ListValue   `tfsdk:"exception_list"`
	Recurring          basetypes.ObjectValue `tfsdk:"recurring"`
	Url                basetypes.StringValue `tfsdk:"url"`
}

// ExternalDynamicListsTypeIpAuth represents a nested structure within the ExternalDynamicLists model
type ExternalDynamicListsTypeIpAuth struct {
	Password basetypes.StringValue `tfsdk:"password"`
	Username basetypes.StringValue `tfsdk:"username"`
}

// ExternalDynamicListsTypeIpRecurring represents a nested structure within the ExternalDynamicLists model
type ExternalDynamicListsTypeIpRecurring struct {
	Daily      basetypes.ObjectValue `tfsdk:"daily"`
	FiveMinute basetypes.ObjectValue `tfsdk:"five_minute"`
	Hourly     basetypes.ObjectValue `tfsdk:"hourly"`
	Monthly    basetypes.ObjectValue `tfsdk:"monthly"`
	Weekly     basetypes.ObjectValue `tfsdk:"weekly"`
}

// ExternalDynamicListsTypeIpRecurringDaily represents a nested structure within the ExternalDynamicLists model
type ExternalDynamicListsTypeIpRecurringDaily struct {
	At basetypes.StringValue `tfsdk:"at"`
}

// ExternalDynamicListsTypeIpRecurringMonthly represents a nested structure within the ExternalDynamicLists model
type ExternalDynamicListsTypeIpRecurringMonthly struct {
	At         basetypes.StringValue `tfsdk:"at"`
	DayOfMonth basetypes.Int64Value  `tfsdk:"day_of_month"`
}

// ExternalDynamicListsTypeIpRecurringWeekly represents a nested structure within the ExternalDynamicLists model
type ExternalDynamicListsTypeIpRecurringWeekly struct {
	At        basetypes.StringValue `tfsdk:"at"`
	DayOfWeek basetypes.StringValue `tfsdk:"day_of_week"`
}

// ExternalDynamicListsTypePredefinedIp represents a nested structure within the ExternalDynamicLists model
type ExternalDynamicListsTypePredefinedIp struct {
	Description   basetypes.StringValue `tfsdk:"description"`
	ExceptionList basetypes.ListValue   `tfsdk:"exception_list"`
	Url           basetypes.StringValue `tfsdk:"url"`
}

// ExternalDynamicListsTypePredefinedUrl represents a nested structure within the ExternalDynamicLists model
type ExternalDynamicListsTypePredefinedUrl struct {
	Description   basetypes.StringValue `tfsdk:"description"`
	ExceptionList basetypes.ListValue   `tfsdk:"exception_list"`
	Url           basetypes.StringValue `tfsdk:"url"`
}

// ExternalDynamicListsTypeUrl represents a nested structure within the ExternalDynamicLists model
type ExternalDynamicListsTypeUrl struct {
	Auth               basetypes.ObjectValue `tfsdk:"auth"`
	CertificateProfile basetypes.StringValue `tfsdk:"certificate_profile"`
	Description        basetypes.StringValue `tfsdk:"description"`
	ExceptionList      basetypes.ListValue   `tfsdk:"exception_list"`
	Recurring          basetypes.ObjectValue `tfsdk:"recurring"`
	Url                basetypes.StringValue `tfsdk:"url"`
}

// ExternalDynamicListsTypeUrlAuth represents a nested structure within the ExternalDynamicLists model
type ExternalDynamicListsTypeUrlAuth struct {
	Password basetypes.StringValue `tfsdk:"password"`
	Username basetypes.StringValue `tfsdk:"username"`
}

// ExternalDynamicListsTypeUrlRecurring represents a nested structure within the ExternalDynamicLists model
type ExternalDynamicListsTypeUrlRecurring struct {
	Daily      basetypes.ObjectValue `tfsdk:"daily"`
	FiveMinute basetypes.ObjectValue `tfsdk:"five_minute"`
	Hourly     basetypes.ObjectValue `tfsdk:"hourly"`
	Monthly    basetypes.ObjectValue `tfsdk:"monthly"`
	Weekly     basetypes.ObjectValue `tfsdk:"weekly"`
}

// ExternalDynamicListsTypeUrlRecurringDaily represents a nested structure within the ExternalDynamicLists model
type ExternalDynamicListsTypeUrlRecurringDaily struct {
	At basetypes.StringValue `tfsdk:"at"`
}

// ExternalDynamicListsTypeUrlRecurringMonthly represents a nested structure within the ExternalDynamicLists model
type ExternalDynamicListsTypeUrlRecurringMonthly struct {
	At         basetypes.StringValue `tfsdk:"at"`
	DayOfMonth basetypes.Int64Value  `tfsdk:"day_of_month"`
}

// ExternalDynamicListsTypeUrlRecurringWeekly represents a nested structure within the ExternalDynamicLists model
type ExternalDynamicListsTypeUrlRecurringWeekly struct {
	At        basetypes.StringValue `tfsdk:"at"`
	DayOfWeek basetypes.StringValue `tfsdk:"day_of_week"`
}

// AttrTypes defines the attribute types for the ExternalDynamicLists model.
func (o ExternalDynamicLists) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":             basetypes.StringType{},
		"encrypted_values": basetypes.MapType{ElemType: basetypes.StringType{}},
		"device":           basetypes.StringType{},
		"folder":           basetypes.StringType{},
		"id":               basetypes.StringType{},
		"name":             basetypes.StringType{},
		"snippet":          basetypes.StringType{},
		"type": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"domain": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"auth": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"password": basetypes.StringType{},
								"username": basetypes.StringType{},
							},
						},
						"certificate_profile": basetypes.StringType{},
						"description":         basetypes.StringType{},
						"exception_list":      basetypes.ListType{ElemType: basetypes.StringType{}},
						"expand_domain":       basetypes.BoolType{},
						"recurring": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"daily": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"at": basetypes.StringType{},
									},
								},
								"five_minute": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"hourly": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"monthly": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"at":           basetypes.StringType{},
										"day_of_month": basetypes.Int64Type{},
									},
								},
								"weekly": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"at":          basetypes.StringType{},
										"day_of_week": basetypes.StringType{},
									},
								},
							},
						},
						"url": basetypes.StringType{},
					},
				},
				"imei": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"auth": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"password": basetypes.StringType{},
								"username": basetypes.StringType{},
							},
						},
						"certificate_profile": basetypes.StringType{},
						"description":         basetypes.StringType{},
						"exception_list":      basetypes.ListType{ElemType: basetypes.StringType{}},
						"recurring": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"daily": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"at": basetypes.StringType{},
									},
								},
								"five_minute": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"hourly": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"monthly": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"at":           basetypes.StringType{},
										"day_of_month": basetypes.Int64Type{},
									},
								},
								"weekly": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"at":          basetypes.StringType{},
										"day_of_week": basetypes.StringType{},
									},
								},
							},
						},
						"url": basetypes.StringType{},
					},
				},
				"imsi": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"auth": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"password": basetypes.StringType{},
								"username": basetypes.StringType{},
							},
						},
						"certificate_profile": basetypes.StringType{},
						"description":         basetypes.StringType{},
						"exception_list":      basetypes.ListType{ElemType: basetypes.StringType{}},
						"recurring": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"daily": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"at": basetypes.StringType{},
									},
								},
								"five_minute": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"hourly": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"monthly": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"at":           basetypes.StringType{},
										"day_of_month": basetypes.Int64Type{},
									},
								},
								"weekly": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"at":          basetypes.StringType{},
										"day_of_week": basetypes.StringType{},
									},
								},
							},
						},
						"url": basetypes.StringType{},
					},
				},
				"ip": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"auth": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"password": basetypes.StringType{},
								"username": basetypes.StringType{},
							},
						},
						"certificate_profile": basetypes.StringType{},
						"description":         basetypes.StringType{},
						"exception_list":      basetypes.ListType{ElemType: basetypes.StringType{}},
						"recurring": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"daily": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"at": basetypes.StringType{},
									},
								},
								"five_minute": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"hourly": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"monthly": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"at":           basetypes.StringType{},
										"day_of_month": basetypes.Int64Type{},
									},
								},
								"weekly": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"at":          basetypes.StringType{},
										"day_of_week": basetypes.StringType{},
									},
								},
							},
						},
						"url": basetypes.StringType{},
					},
				},
				"predefined_ip": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"description":    basetypes.StringType{},
						"exception_list": basetypes.ListType{ElemType: basetypes.StringType{}},
						"url":            basetypes.StringType{},
					},
				},
				"predefined_url": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"description":    basetypes.StringType{},
						"exception_list": basetypes.ListType{ElemType: basetypes.StringType{}},
						"url":            basetypes.StringType{},
					},
				},
				"url": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"auth": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"password": basetypes.StringType{},
								"username": basetypes.StringType{},
							},
						},
						"certificate_profile": basetypes.StringType{},
						"description":         basetypes.StringType{},
						"exception_list":      basetypes.ListType{ElemType: basetypes.StringType{}},
						"recurring": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"daily": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"at": basetypes.StringType{},
									},
								},
								"five_minute": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"hourly": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"monthly": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"at":           basetypes.StringType{},
										"day_of_month": basetypes.Int64Type{},
									},
								},
								"weekly": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"at":          basetypes.StringType{},
										"day_of_week": basetypes.StringType{},
									},
								},
							},
						},
						"url": basetypes.StringType{},
					},
				},
			},
		},
	}
}

// AttrType returns the attribute type for a list of ExternalDynamicLists objects.
func (o ExternalDynamicLists) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ExternalDynamicListsType model.
func (o ExternalDynamicListsType) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"domain": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"auth": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"password": basetypes.StringType{},
						"username": basetypes.StringType{},
					},
				},
				"certificate_profile": basetypes.StringType{},
				"description":         basetypes.StringType{},
				"exception_list":      basetypes.ListType{ElemType: basetypes.StringType{}},
				"expand_domain":       basetypes.BoolType{},
				"recurring": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"daily": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"at": basetypes.StringType{},
							},
						},
						"five_minute": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"hourly": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"monthly": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"at":           basetypes.StringType{},
								"day_of_month": basetypes.Int64Type{},
							},
						},
						"weekly": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"at":          basetypes.StringType{},
								"day_of_week": basetypes.StringType{},
							},
						},
					},
				},
				"url": basetypes.StringType{},
			},
		},
		"imei": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"auth": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"password": basetypes.StringType{},
						"username": basetypes.StringType{},
					},
				},
				"certificate_profile": basetypes.StringType{},
				"description":         basetypes.StringType{},
				"exception_list":      basetypes.ListType{ElemType: basetypes.StringType{}},
				"recurring": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"daily": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"at": basetypes.StringType{},
							},
						},
						"five_minute": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"hourly": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"monthly": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"at":           basetypes.StringType{},
								"day_of_month": basetypes.Int64Type{},
							},
						},
						"weekly": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"at":          basetypes.StringType{},
								"day_of_week": basetypes.StringType{},
							},
						},
					},
				},
				"url": basetypes.StringType{},
			},
		},
		"imsi": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"auth": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"password": basetypes.StringType{},
						"username": basetypes.StringType{},
					},
				},
				"certificate_profile": basetypes.StringType{},
				"description":         basetypes.StringType{},
				"exception_list":      basetypes.ListType{ElemType: basetypes.StringType{}},
				"recurring": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"daily": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"at": basetypes.StringType{},
							},
						},
						"five_minute": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"hourly": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"monthly": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"at":           basetypes.StringType{},
								"day_of_month": basetypes.Int64Type{},
							},
						},
						"weekly": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"at":          basetypes.StringType{},
								"day_of_week": basetypes.StringType{},
							},
						},
					},
				},
				"url": basetypes.StringType{},
			},
		},
		"ip": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"auth": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"password": basetypes.StringType{},
						"username": basetypes.StringType{},
					},
				},
				"certificate_profile": basetypes.StringType{},
				"description":         basetypes.StringType{},
				"exception_list":      basetypes.ListType{ElemType: basetypes.StringType{}},
				"recurring": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"daily": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"at": basetypes.StringType{},
							},
						},
						"five_minute": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"hourly": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"monthly": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"at":           basetypes.StringType{},
								"day_of_month": basetypes.Int64Type{},
							},
						},
						"weekly": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"at":          basetypes.StringType{},
								"day_of_week": basetypes.StringType{},
							},
						},
					},
				},
				"url": basetypes.StringType{},
			},
		},
		"predefined_ip": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"description":    basetypes.StringType{},
				"exception_list": basetypes.ListType{ElemType: basetypes.StringType{}},
				"url":            basetypes.StringType{},
			},
		},
		"predefined_url": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"description":    basetypes.StringType{},
				"exception_list": basetypes.ListType{ElemType: basetypes.StringType{}},
				"url":            basetypes.StringType{},
			},
		},
		"url": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"auth": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"password": basetypes.StringType{},
						"username": basetypes.StringType{},
					},
				},
				"certificate_profile": basetypes.StringType{},
				"description":         basetypes.StringType{},
				"exception_list":      basetypes.ListType{ElemType: basetypes.StringType{}},
				"recurring": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"daily": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"at": basetypes.StringType{},
							},
						},
						"five_minute": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"hourly": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"monthly": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"at":           basetypes.StringType{},
								"day_of_month": basetypes.Int64Type{},
							},
						},
						"weekly": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"at":          basetypes.StringType{},
								"day_of_week": basetypes.StringType{},
							},
						},
					},
				},
				"url": basetypes.StringType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of ExternalDynamicListsType objects.
func (o ExternalDynamicListsType) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ExternalDynamicListsTypeDomain model.
func (o ExternalDynamicListsTypeDomain) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"auth": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"password": basetypes.StringType{},
				"username": basetypes.StringType{},
			},
		},
		"certificate_profile": basetypes.StringType{},
		"description":         basetypes.StringType{},
		"exception_list":      basetypes.ListType{ElemType: basetypes.StringType{}},
		"expand_domain":       basetypes.BoolType{},
		"recurring": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"daily": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"at": basetypes.StringType{},
					},
				},
				"five_minute": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"hourly": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"monthly": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"at":           basetypes.StringType{},
						"day_of_month": basetypes.Int64Type{},
					},
				},
				"weekly": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"at":          basetypes.StringType{},
						"day_of_week": basetypes.StringType{},
					},
				},
			},
		},
		"url": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of ExternalDynamicListsTypeDomain objects.
func (o ExternalDynamicListsTypeDomain) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ExternalDynamicListsTypeDomainAuth model.
func (o ExternalDynamicListsTypeDomainAuth) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"password": basetypes.StringType{},
		"username": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of ExternalDynamicListsTypeDomainAuth objects.
func (o ExternalDynamicListsTypeDomainAuth) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ExternalDynamicListsTypeDomainRecurring model.
func (o ExternalDynamicListsTypeDomainRecurring) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"daily": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"at": basetypes.StringType{},
			},
		},
		"five_minute": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
		"hourly": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
		"monthly": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"at":           basetypes.StringType{},
				"day_of_month": basetypes.Int64Type{},
			},
		},
		"weekly": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"at":          basetypes.StringType{},
				"day_of_week": basetypes.StringType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of ExternalDynamicListsTypeDomainRecurring objects.
func (o ExternalDynamicListsTypeDomainRecurring) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ExternalDynamicListsTypeDomainRecurringDaily model.
func (o ExternalDynamicListsTypeDomainRecurringDaily) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"at": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of ExternalDynamicListsTypeDomainRecurringDaily objects.
func (o ExternalDynamicListsTypeDomainRecurringDaily) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ExternalDynamicListsTypeDomainRecurringMonthly model.
func (o ExternalDynamicListsTypeDomainRecurringMonthly) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"at":           basetypes.StringType{},
		"day_of_month": basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of ExternalDynamicListsTypeDomainRecurringMonthly objects.
func (o ExternalDynamicListsTypeDomainRecurringMonthly) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ExternalDynamicListsTypeDomainRecurringWeekly model.
func (o ExternalDynamicListsTypeDomainRecurringWeekly) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"at":          basetypes.StringType{},
		"day_of_week": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of ExternalDynamicListsTypeDomainRecurringWeekly objects.
func (o ExternalDynamicListsTypeDomainRecurringWeekly) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ExternalDynamicListsTypeImei model.
func (o ExternalDynamicListsTypeImei) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"auth": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"password": basetypes.StringType{},
				"username": basetypes.StringType{},
			},
		},
		"certificate_profile": basetypes.StringType{},
		"description":         basetypes.StringType{},
		"exception_list":      basetypes.ListType{ElemType: basetypes.StringType{}},
		"recurring": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"daily": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"at": basetypes.StringType{},
					},
				},
				"five_minute": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"hourly": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"monthly": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"at":           basetypes.StringType{},
						"day_of_month": basetypes.Int64Type{},
					},
				},
				"weekly": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"at":          basetypes.StringType{},
						"day_of_week": basetypes.StringType{},
					},
				},
			},
		},
		"url": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of ExternalDynamicListsTypeImei objects.
func (o ExternalDynamicListsTypeImei) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ExternalDynamicListsTypeImeiAuth model.
func (o ExternalDynamicListsTypeImeiAuth) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"password": basetypes.StringType{},
		"username": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of ExternalDynamicListsTypeImeiAuth objects.
func (o ExternalDynamicListsTypeImeiAuth) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ExternalDynamicListsTypeImeiRecurring model.
func (o ExternalDynamicListsTypeImeiRecurring) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"daily": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"at": basetypes.StringType{},
			},
		},
		"five_minute": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
		"hourly": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
		"monthly": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"at":           basetypes.StringType{},
				"day_of_month": basetypes.Int64Type{},
			},
		},
		"weekly": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"at":          basetypes.StringType{},
				"day_of_week": basetypes.StringType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of ExternalDynamicListsTypeImeiRecurring objects.
func (o ExternalDynamicListsTypeImeiRecurring) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ExternalDynamicListsTypeImeiRecurringDaily model.
func (o ExternalDynamicListsTypeImeiRecurringDaily) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"at": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of ExternalDynamicListsTypeImeiRecurringDaily objects.
func (o ExternalDynamicListsTypeImeiRecurringDaily) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ExternalDynamicListsTypeImeiRecurringMonthly model.
func (o ExternalDynamicListsTypeImeiRecurringMonthly) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"at":           basetypes.StringType{},
		"day_of_month": basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of ExternalDynamicListsTypeImeiRecurringMonthly objects.
func (o ExternalDynamicListsTypeImeiRecurringMonthly) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ExternalDynamicListsTypeImeiRecurringWeekly model.
func (o ExternalDynamicListsTypeImeiRecurringWeekly) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"at":          basetypes.StringType{},
		"day_of_week": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of ExternalDynamicListsTypeImeiRecurringWeekly objects.
func (o ExternalDynamicListsTypeImeiRecurringWeekly) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ExternalDynamicListsTypeImsi model.
func (o ExternalDynamicListsTypeImsi) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"auth": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"password": basetypes.StringType{},
				"username": basetypes.StringType{},
			},
		},
		"certificate_profile": basetypes.StringType{},
		"description":         basetypes.StringType{},
		"exception_list":      basetypes.ListType{ElemType: basetypes.StringType{}},
		"recurring": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"daily": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"at": basetypes.StringType{},
					},
				},
				"five_minute": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"hourly": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"monthly": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"at":           basetypes.StringType{},
						"day_of_month": basetypes.Int64Type{},
					},
				},
				"weekly": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"at":          basetypes.StringType{},
						"day_of_week": basetypes.StringType{},
					},
				},
			},
		},
		"url": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of ExternalDynamicListsTypeImsi objects.
func (o ExternalDynamicListsTypeImsi) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ExternalDynamicListsTypeImsiAuth model.
func (o ExternalDynamicListsTypeImsiAuth) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"password": basetypes.StringType{},
		"username": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of ExternalDynamicListsTypeImsiAuth objects.
func (o ExternalDynamicListsTypeImsiAuth) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ExternalDynamicListsTypeImsiRecurring model.
func (o ExternalDynamicListsTypeImsiRecurring) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"daily": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"at": basetypes.StringType{},
			},
		},
		"five_minute": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
		"hourly": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
		"monthly": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"at":           basetypes.StringType{},
				"day_of_month": basetypes.Int64Type{},
			},
		},
		"weekly": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"at":          basetypes.StringType{},
				"day_of_week": basetypes.StringType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of ExternalDynamicListsTypeImsiRecurring objects.
func (o ExternalDynamicListsTypeImsiRecurring) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ExternalDynamicListsTypeImsiRecurringDaily model.
func (o ExternalDynamicListsTypeImsiRecurringDaily) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"at": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of ExternalDynamicListsTypeImsiRecurringDaily objects.
func (o ExternalDynamicListsTypeImsiRecurringDaily) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ExternalDynamicListsTypeImsiRecurringMonthly model.
func (o ExternalDynamicListsTypeImsiRecurringMonthly) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"at":           basetypes.StringType{},
		"day_of_month": basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of ExternalDynamicListsTypeImsiRecurringMonthly objects.
func (o ExternalDynamicListsTypeImsiRecurringMonthly) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ExternalDynamicListsTypeImsiRecurringWeekly model.
func (o ExternalDynamicListsTypeImsiRecurringWeekly) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"at":          basetypes.StringType{},
		"day_of_week": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of ExternalDynamicListsTypeImsiRecurringWeekly objects.
func (o ExternalDynamicListsTypeImsiRecurringWeekly) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ExternalDynamicListsTypeIp model.
func (o ExternalDynamicListsTypeIp) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"auth": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"password": basetypes.StringType{},
				"username": basetypes.StringType{},
			},
		},
		"certificate_profile": basetypes.StringType{},
		"description":         basetypes.StringType{},
		"exception_list":      basetypes.ListType{ElemType: basetypes.StringType{}},
		"recurring": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"daily": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"at": basetypes.StringType{},
					},
				},
				"five_minute": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"hourly": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"monthly": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"at":           basetypes.StringType{},
						"day_of_month": basetypes.Int64Type{},
					},
				},
				"weekly": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"at":          basetypes.StringType{},
						"day_of_week": basetypes.StringType{},
					},
				},
			},
		},
		"url": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of ExternalDynamicListsTypeIp objects.
func (o ExternalDynamicListsTypeIp) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ExternalDynamicListsTypeIpAuth model.
func (o ExternalDynamicListsTypeIpAuth) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"password": basetypes.StringType{},
		"username": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of ExternalDynamicListsTypeIpAuth objects.
func (o ExternalDynamicListsTypeIpAuth) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ExternalDynamicListsTypeIpRecurring model.
func (o ExternalDynamicListsTypeIpRecurring) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"daily": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"at": basetypes.StringType{},
			},
		},
		"five_minute": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
		"hourly": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
		"monthly": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"at":           basetypes.StringType{},
				"day_of_month": basetypes.Int64Type{},
			},
		},
		"weekly": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"at":          basetypes.StringType{},
				"day_of_week": basetypes.StringType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of ExternalDynamicListsTypeIpRecurring objects.
func (o ExternalDynamicListsTypeIpRecurring) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ExternalDynamicListsTypeIpRecurringDaily model.
func (o ExternalDynamicListsTypeIpRecurringDaily) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"at": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of ExternalDynamicListsTypeIpRecurringDaily objects.
func (o ExternalDynamicListsTypeIpRecurringDaily) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ExternalDynamicListsTypeIpRecurringMonthly model.
func (o ExternalDynamicListsTypeIpRecurringMonthly) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"at":           basetypes.StringType{},
		"day_of_month": basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of ExternalDynamicListsTypeIpRecurringMonthly objects.
func (o ExternalDynamicListsTypeIpRecurringMonthly) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ExternalDynamicListsTypeIpRecurringWeekly model.
func (o ExternalDynamicListsTypeIpRecurringWeekly) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"at":          basetypes.StringType{},
		"day_of_week": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of ExternalDynamicListsTypeIpRecurringWeekly objects.
func (o ExternalDynamicListsTypeIpRecurringWeekly) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ExternalDynamicListsTypePredefinedIp model.
func (o ExternalDynamicListsTypePredefinedIp) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"description":    basetypes.StringType{},
		"exception_list": basetypes.ListType{ElemType: basetypes.StringType{}},
		"url":            basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of ExternalDynamicListsTypePredefinedIp objects.
func (o ExternalDynamicListsTypePredefinedIp) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ExternalDynamicListsTypePredefinedUrl model.
func (o ExternalDynamicListsTypePredefinedUrl) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"description":    basetypes.StringType{},
		"exception_list": basetypes.ListType{ElemType: basetypes.StringType{}},
		"url":            basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of ExternalDynamicListsTypePredefinedUrl objects.
func (o ExternalDynamicListsTypePredefinedUrl) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ExternalDynamicListsTypeUrl model.
func (o ExternalDynamicListsTypeUrl) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"auth": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"password": basetypes.StringType{},
				"username": basetypes.StringType{},
			},
		},
		"certificate_profile": basetypes.StringType{},
		"description":         basetypes.StringType{},
		"exception_list":      basetypes.ListType{ElemType: basetypes.StringType{}},
		"recurring": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"daily": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"at": basetypes.StringType{},
					},
				},
				"five_minute": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"hourly": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"monthly": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"at":           basetypes.StringType{},
						"day_of_month": basetypes.Int64Type{},
					},
				},
				"weekly": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"at":          basetypes.StringType{},
						"day_of_week": basetypes.StringType{},
					},
				},
			},
		},
		"url": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of ExternalDynamicListsTypeUrl objects.
func (o ExternalDynamicListsTypeUrl) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ExternalDynamicListsTypeUrlAuth model.
func (o ExternalDynamicListsTypeUrlAuth) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"password": basetypes.StringType{},
		"username": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of ExternalDynamicListsTypeUrlAuth objects.
func (o ExternalDynamicListsTypeUrlAuth) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ExternalDynamicListsTypeUrlRecurring model.
func (o ExternalDynamicListsTypeUrlRecurring) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"daily": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"at": basetypes.StringType{},
			},
		},
		"five_minute": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
		"hourly": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
		"monthly": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"at":           basetypes.StringType{},
				"day_of_month": basetypes.Int64Type{},
			},
		},
		"weekly": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"at":          basetypes.StringType{},
				"day_of_week": basetypes.StringType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of ExternalDynamicListsTypeUrlRecurring objects.
func (o ExternalDynamicListsTypeUrlRecurring) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ExternalDynamicListsTypeUrlRecurringDaily model.
func (o ExternalDynamicListsTypeUrlRecurringDaily) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"at": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of ExternalDynamicListsTypeUrlRecurringDaily objects.
func (o ExternalDynamicListsTypeUrlRecurringDaily) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ExternalDynamicListsTypeUrlRecurringMonthly model.
func (o ExternalDynamicListsTypeUrlRecurringMonthly) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"at":           basetypes.StringType{},
		"day_of_month": basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of ExternalDynamicListsTypeUrlRecurringMonthly objects.
func (o ExternalDynamicListsTypeUrlRecurringMonthly) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ExternalDynamicListsTypeUrlRecurringWeekly model.
func (o ExternalDynamicListsTypeUrlRecurringWeekly) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"at":          basetypes.StringType{},
		"day_of_week": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of ExternalDynamicListsTypeUrlRecurringWeekly objects.
func (o ExternalDynamicListsTypeUrlRecurringWeekly) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// ExternalDynamicListsResourceSchema defines the schema for ExternalDynamicLists resource
var ExternalDynamicListsResourceSchema = schema.Schema{
	MarkdownDescription: "ExternalDynamicList resource",
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
			MarkdownDescription: "The device in which the resource is defined\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
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
			MarkdownDescription: "The folder in which the resource is defined\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"id": schema.StringAttribute{
			MarkdownDescription: "The UUID of the external dynamic list",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"name": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(63),
				stringvalidator.RegexMatches(regexp.MustCompile("^[ a-zA-Z\\d.\\-_]+$"), "pattern must match "+"^[ a-zA-Z\\d.\\-_]+$"),
			},
			MarkdownDescription: "The name of the external dynamic list",
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
			MarkdownDescription: "The snippet in which the resource is defined\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
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
		"type": schema.SingleNestedAttribute{
			MarkdownDescription: "Type configuration for External Dynamic List",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"domain": schema.SingleNestedAttribute{
					Validators: []validator.Object{
						objectvalidator.ConflictsWith(
							path.MatchRelative().AtParent().AtName("imei"),
							path.MatchRelative().AtParent().AtName("imsi"),
							path.MatchRelative().AtParent().AtName("ip"),
							path.MatchRelative().AtParent().AtName("predefined_ip"),
							path.MatchRelative().AtParent().AtName("predefined_url"),
							path.MatchRelative().AtParent().AtName("url"),
						),
					},
					MarkdownDescription: "Domain settings for Custom Domain type\n> ℹ️ **Note:** You must specify exactly one of `domain`, `imei`, `imsi`, `ip`, `predefined_ip`, `predefined_url`, and `url`.",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"auth": schema.SingleNestedAttribute{
							MarkdownDescription: "Authentication settings for Custom Domain type",
							Optional:            true,
							Attributes: map[string]schema.Attribute{
								"password": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.LengthAtMost(255),
									},
									MarkdownDescription: "Password for Custom Domain authentication",
									Required:            true,
									Sensitive:           true,
								},
								"username": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.LengthAtMost(255),
										stringvalidator.LengthAtLeast(1),
									},
									MarkdownDescription: "Username for Custom Domain authentication",
									Required:            true,
								},
							},
						},
						"certificate_profile": schema.StringAttribute{
							MarkdownDescription: "Profile for authenticating client certificates",
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("None"),
						},
						"description": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.LengthAtMost(255),
							},
							MarkdownDescription: "Description",
							Optional:            true,
						},
						"exception_list": schema.ListAttribute{
							ElementType:         types.StringType,
							MarkdownDescription: "Domain Exception List for Custom Domain type",
							Validators: []validator.List{
								listvalidator.ValueStringsAre(stringvalidator.LengthAtMost(255)),
							},
							Optional: true,
						},
						"expand_domain": schema.BoolAttribute{
							MarkdownDescription: "Enable/Disable expand domain",
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(false),
						},
						"recurring": schema.SingleNestedAttribute{
							MarkdownDescription: "Update Schedule for Custom Domain type",
							Required:            true,
							Attributes: map[string]schema.Attribute{
								"daily": schema.SingleNestedAttribute{
									Validators: []validator.Object{
										objectvalidator.ExactlyOneOf(
											path.MatchRelative().AtParent().AtName("five_minute"),
											path.MatchRelative().AtParent().AtName("hourly"),
											path.MatchRelative().AtParent().AtName("monthly"),
											path.MatchRelative().AtParent().AtName("weekly"),
										),
									},
									MarkdownDescription: "Daily settings for Domain recurring\n> ℹ️ **Note:** You must specify exactly one of `daily`, `five_minute`, `hourly`, `monthly`, and `weekly`.",
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"at": schema.StringAttribute{
											Validators: []validator.String{
												stringvalidator.LengthAtMost(2),
												stringvalidator.LengthAtLeast(2),
												stringvalidator.RegexMatches(regexp.MustCompile("([01][0-9]|[2][0-3])"), "pattern must match "+"([01][0-9]|[2][0-3])"),
											},
											MarkdownDescription: "Daily Time specification hh (e.g. 20) for Domain",
											Required:            true,
										},
									},
								},
								"five_minute": schema.SingleNestedAttribute{
									Validators: []validator.Object{
										objectvalidator.ExactlyOneOf(
											path.MatchRelative().AtParent().AtName("daily"),
											path.MatchRelative().AtParent().AtName("hourly"),
											path.MatchRelative().AtParent().AtName("monthly"),
											path.MatchRelative().AtParent().AtName("weekly"),
										),
									},
									MarkdownDescription: "Five minute settings for Domain recurring\n> ℹ️ **Note:** You must specify exactly one of `daily`, `five_minute`, `hourly`, `monthly`, and `weekly`.",
									Optional:            true,
									Attributes:          map[string]schema.Attribute{},
								},
								"hourly": schema.SingleNestedAttribute{
									Validators: []validator.Object{
										objectvalidator.ExactlyOneOf(
											path.MatchRelative().AtParent().AtName("daily"),
											path.MatchRelative().AtParent().AtName("five_minute"),
											path.MatchRelative().AtParent().AtName("monthly"),
											path.MatchRelative().AtParent().AtName("weekly"),
										),
									},
									MarkdownDescription: "Hourly settings for Domain recurring\n> ℹ️ **Note:** You must specify exactly one of `daily`, `five_minute`, `hourly`, `monthly`, and `weekly`.",
									Optional:            true,
									Attributes:          map[string]schema.Attribute{},
								},
								"monthly": schema.SingleNestedAttribute{
									Validators: []validator.Object{
										objectvalidator.ExactlyOneOf(
											path.MatchRelative().AtParent().AtName("daily"),
											path.MatchRelative().AtParent().AtName("five_minute"),
											path.MatchRelative().AtParent().AtName("hourly"),
											path.MatchRelative().AtParent().AtName("weekly"),
										),
									},
									MarkdownDescription: "Monthly settings for Domain recurring\n> ℹ️ **Note:** You must specify exactly one of `daily`, `five_minute`, `hourly`, `monthly`, and `weekly`.",
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"at": schema.StringAttribute{
											Validators: []validator.String{
												stringvalidator.LengthAtMost(2),
												stringvalidator.LengthAtLeast(2),
												stringvalidator.RegexMatches(regexp.MustCompile("([01][0-9]|[2][0-3])"), "pattern must match "+"([01][0-9]|[2][0-3])"),
											},
											MarkdownDescription: "Monthly Time specification hh (e.g. 20) for domain",
											Required:            true,
										},
										"day_of_month": schema.Int64Attribute{
											Validators: []validator.Int64{
												int64validator.Between(1, 31),
											},
											MarkdownDescription: "Day setting for monthly Domain updates",
											Required:            true,
										},
									},
								},
								"weekly": schema.SingleNestedAttribute{
									Validators: []validator.Object{
										objectvalidator.ExactlyOneOf(
											path.MatchRelative().AtParent().AtName("daily"),
											path.MatchRelative().AtParent().AtName("five_minute"),
											path.MatchRelative().AtParent().AtName("hourly"),
											path.MatchRelative().AtParent().AtName("monthly"),
										),
									},
									MarkdownDescription: "Weekly settings for Domain recurring\n> ℹ️ **Note:** You must specify exactly one of `daily`, `five_minute`, `hourly`, `monthly`, and `weekly`.",
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"at": schema.StringAttribute{
											Validators: []validator.String{
												stringvalidator.LengthAtMost(2),
												stringvalidator.LengthAtLeast(2),
												stringvalidator.RegexMatches(regexp.MustCompile("([01][0-9]|[2][0-3])"), "pattern must match "+"([01][0-9]|[2][0-3])"),
											},
											MarkdownDescription: "Weekly Time specification hh (e.g. 20) for Domain",
											Required:            true,
										},
										"day_of_week": schema.StringAttribute{
											Validators: []validator.String{
												stringvalidator.OneOf("sunday", "monday", "tuesday", "wednesday", "thursday", "friday", "saturday"),
											},
											MarkdownDescription: "Day of week",
											Required:            true,
										},
									},
								},
							},
						},
						"url": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.LengthAtMost(255),
							},
							MarkdownDescription: "External URL for Custom Domain type",
							Required:            true,
						},
					},
				},
				"imei": schema.SingleNestedAttribute{
					Validators: []validator.Object{
						objectvalidator.ConflictsWith(
							path.MatchRelative().AtParent().AtName("domain"),
							path.MatchRelative().AtParent().AtName("imsi"),
							path.MatchRelative().AtParent().AtName("ip"),
							path.MatchRelative().AtParent().AtName("predefined_ip"),
							path.MatchRelative().AtParent().AtName("predefined_url"),
							path.MatchRelative().AtParent().AtName("url"),
						),
					},
					MarkdownDescription: "IMEI Configuration settings\n> ℹ️ **Note:** You must specify exactly one of `domain`, `imei`, `imsi`, `ip`, `predefined_ip`, `predefined_url`, and `url`.",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"auth": schema.SingleNestedAttribute{
							MarkdownDescription: "IMEI Auth Cnfig for Custom IMEI type",
							Optional:            true,
							Attributes: map[string]schema.Attribute{
								"password": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.LengthAtMost(255),
									},
									MarkdownDescription: "IMEI Auth Password for Custom IMEI type",
									Required:            true,
									Sensitive:           true,
								},
								"username": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.LengthAtMost(255),
										stringvalidator.LengthAtLeast(1),
									},
									MarkdownDescription: "IMEI Auth username for Custom IMEI type",
									Required:            true,
								},
							},
						},
						"certificate_profile": schema.StringAttribute{
							MarkdownDescription: "IMEI Certificate Profile for Custom IMEI type",
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("None"),
						},
						"description": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.LengthAtMost(255),
							},
							MarkdownDescription: "IMEI Description for Custom IMEI type",
							Optional:            true,
						},
						"exception_list": schema.ListAttribute{
							ElementType:         types.StringType,
							MarkdownDescription: "IMEI Exception List for Custom IMEI type",
							Validators: []validator.List{
								listvalidator.ValueStringsAre(stringvalidator.LengthAtMost(32)),
							},
							Optional: true,
						},
						"recurring": schema.SingleNestedAttribute{
							MarkdownDescription: "Recurring interval for IMEI updates",
							Required:            true,
							Attributes: map[string]schema.Attribute{
								"daily": schema.SingleNestedAttribute{
									Validators: []validator.Object{
										objectvalidator.ExactlyOneOf(
											path.MatchRelative().AtParent().AtName("five_minute"),
											path.MatchRelative().AtParent().AtName("hourly"),
											path.MatchRelative().AtParent().AtName("monthly"),
											path.MatchRelative().AtParent().AtName("weekly"),
										),
									},
									MarkdownDescription: "Daily interval settings for IMEI updates\n> ℹ️ **Note:** You must specify exactly one of `daily`, `five_minute`, `hourly`, `monthly`, and `weekly`.",
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"at": schema.StringAttribute{
											Validators: []validator.String{
												stringvalidator.LengthAtMost(2),
												stringvalidator.LengthAtLeast(2),
												stringvalidator.RegexMatches(regexp.MustCompile("([01][0-9]|[2][0-3])"), "pattern must match "+"([01][0-9]|[2][0-3])"),
											},
											MarkdownDescription: "Daily Time specification hh (e.g. 20) for IMEI",
											Required:            true,
										},
									},
								},
								"five_minute": schema.SingleNestedAttribute{
									Validators: []validator.Object{
										objectvalidator.ExactlyOneOf(
											path.MatchRelative().AtParent().AtName("daily"),
											path.MatchRelative().AtParent().AtName("hourly"),
											path.MatchRelative().AtParent().AtName("monthly"),
											path.MatchRelative().AtParent().AtName("weekly"),
										),
									},
									MarkdownDescription: "Five-minute interval settings for IMEI updates\n> ℹ️ **Note:** You must specify exactly one of `daily`, `five_minute`, `hourly`, `monthly`, and `weekly`.",
									Optional:            true,
									Attributes:          map[string]schema.Attribute{},
								},
								"hourly": schema.SingleNestedAttribute{
									Validators: []validator.Object{
										objectvalidator.ExactlyOneOf(
											path.MatchRelative().AtParent().AtName("daily"),
											path.MatchRelative().AtParent().AtName("five_minute"),
											path.MatchRelative().AtParent().AtName("monthly"),
											path.MatchRelative().AtParent().AtName("weekly"),
										),
									},
									MarkdownDescription: "Hourly interval settings for IMEI updates\n> ℹ️ **Note:** You must specify exactly one of `daily`, `five_minute`, `hourly`, `monthly`, and `weekly`.",
									Optional:            true,
									Attributes:          map[string]schema.Attribute{},
								},
								"monthly": schema.SingleNestedAttribute{
									Validators: []validator.Object{
										objectvalidator.ExactlyOneOf(
											path.MatchRelative().AtParent().AtName("daily"),
											path.MatchRelative().AtParent().AtName("five_minute"),
											path.MatchRelative().AtParent().AtName("hourly"),
											path.MatchRelative().AtParent().AtName("weekly"),
										),
									},
									MarkdownDescription: "Monthly interval settings for IMEI updates\n> ℹ️ **Note:** You must specify exactly one of `daily`, `five_minute`, `hourly`, `monthly`, and `weekly`.",
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"at": schema.StringAttribute{
											Validators: []validator.String{
												stringvalidator.LengthAtMost(2),
												stringvalidator.LengthAtLeast(2),
												stringvalidator.RegexMatches(regexp.MustCompile("([01][0-9]|[2][0-3])"), "pattern must match "+"([01][0-9]|[2][0-3])"),
											},
											MarkdownDescription: "Monthly Time specification hh (e.g. 20) for IMEI",
											Required:            true,
										},
										"day_of_month": schema.Int64Attribute{
											Validators: []validator.Int64{
												int64validator.Between(1, 31),
											},
											MarkdownDescription: "Day of month for IMEI updates",
											Required:            true,
										},
									},
								},
								"weekly": schema.SingleNestedAttribute{
									Validators: []validator.Object{
										objectvalidator.ExactlyOneOf(
											path.MatchRelative().AtParent().AtName("daily"),
											path.MatchRelative().AtParent().AtName("five_minute"),
											path.MatchRelative().AtParent().AtName("hourly"),
											path.MatchRelative().AtParent().AtName("monthly"),
										),
									},
									MarkdownDescription: "Weekly interval settings for IMEI updates\n> ℹ️ **Note:** You must specify exactly one of `daily`, `five_minute`, `hourly`, `monthly`, and `weekly`.",
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"at": schema.StringAttribute{
											Validators: []validator.String{
												stringvalidator.LengthAtMost(2),
												stringvalidator.LengthAtLeast(2),
												stringvalidator.RegexMatches(regexp.MustCompile("([01][0-9]|[2][0-3])"), "pattern must match "+"([01][0-9]|[2][0-3])"),
											},
											MarkdownDescription: "Weekly Time specification hh (e.g. 20) for IMEI",
											Required:            true,
										},
										"day_of_week": schema.StringAttribute{
											Validators: []validator.String{
												stringvalidator.OneOf("sunday", "monday", "tuesday", "wednesday", "thursday", "friday", "saturday"),
											},
											MarkdownDescription: "Day of week",
											Required:            true,
										},
									},
								},
							},
						},
						"url": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.LengthAtMost(255),
							},
							MarkdownDescription: "IMEI URL for Custom IMEI type",
							Required:            true,
						},
					},
				},
				"imsi": schema.SingleNestedAttribute{
					Validators: []validator.Object{
						objectvalidator.ConflictsWith(
							path.MatchRelative().AtParent().AtName("domain"),
							path.MatchRelative().AtParent().AtName("imei"),
							path.MatchRelative().AtParent().AtName("ip"),
							path.MatchRelative().AtParent().AtName("predefined_ip"),
							path.MatchRelative().AtParent().AtName("predefined_url"),
							path.MatchRelative().AtParent().AtName("url"),
						),
					},
					MarkdownDescription: "IMSI Config for Custom IMSI type\n> ℹ️ **Note:** You must specify exactly one of `domain`, `imei`, `imsi`, `ip`, `predefined_ip`, `predefined_url`, and `url`.",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"auth": schema.SingleNestedAttribute{
							MarkdownDescription: "IMSI Auth Config for Custom IMSI type",
							Optional:            true,
							Attributes: map[string]schema.Attribute{
								"password": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.LengthAtMost(255),
									},
									MarkdownDescription: "IMSI Auth Password for Custom IMSI type",
									Required:            true,
									Sensitive:           true,
								},
								"username": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.LengthAtMost(255),
										stringvalidator.LengthAtLeast(1),
									},
									MarkdownDescription: "IMSI Auth Username for Custom IMSI type",
									Required:            true,
								},
							},
						},
						"certificate_profile": schema.StringAttribute{
							MarkdownDescription: "IMSI Certificate Profile for Custom IMSI type",
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("None"),
						},
						"description": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.LengthAtMost(255),
							},
							MarkdownDescription: "IMSI Description for Custom IMSI type",
							Optional:            true,
						},
						"exception_list": schema.ListAttribute{
							ElementType:         types.StringType,
							MarkdownDescription: "IMSI Exception List for Custom IMSI type",
							Validators: []validator.List{
								listvalidator.ValueStringsAre(stringvalidator.LengthAtMost(34)),
							},
							Optional: true,
						},
						"recurring": schema.SingleNestedAttribute{
							MarkdownDescription: "IMSI Recuring Config for Custom IMSI type",
							Required:            true,
							Attributes: map[string]schema.Attribute{
								"daily": schema.SingleNestedAttribute{
									Validators: []validator.Object{
										objectvalidator.ExactlyOneOf(
											path.MatchRelative().AtParent().AtName("five_minute"),
											path.MatchRelative().AtParent().AtName("hourly"),
											path.MatchRelative().AtParent().AtName("monthly"),
											path.MatchRelative().AtParent().AtName("weekly"),
										),
									},
									MarkdownDescription: "Daily interval settings for IMSI updates\n> ℹ️ **Note:** You must specify exactly one of `daily`, `five_minute`, `hourly`, `monthly`, and `weekly`.",
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"at": schema.StringAttribute{
											Validators: []validator.String{
												stringvalidator.LengthAtMost(2),
												stringvalidator.LengthAtLeast(2),
												stringvalidator.RegexMatches(regexp.MustCompile("([01][0-9]|[2][0-3])"), "pattern must match "+"([01][0-9]|[2][0-3])"),
											},
											MarkdownDescription: "Daily Time specification hh (e.g. 20) for IMSI",
											Required:            true,
										},
									},
								},
								"five_minute": schema.SingleNestedAttribute{
									Validators: []validator.Object{
										objectvalidator.ExactlyOneOf(
											path.MatchRelative().AtParent().AtName("daily"),
											path.MatchRelative().AtParent().AtName("hourly"),
											path.MatchRelative().AtParent().AtName("monthly"),
											path.MatchRelative().AtParent().AtName("weekly"),
										),
									},
									MarkdownDescription: "Five-minute interval settings for IMSI updates\n> ℹ️ **Note:** You must specify exactly one of `daily`, `five_minute`, `hourly`, `monthly`, and `weekly`.",
									Optional:            true,
									Attributes:          map[string]schema.Attribute{},
								},
								"hourly": schema.SingleNestedAttribute{
									Validators: []validator.Object{
										objectvalidator.ExactlyOneOf(
											path.MatchRelative().AtParent().AtName("daily"),
											path.MatchRelative().AtParent().AtName("five_minute"),
											path.MatchRelative().AtParent().AtName("monthly"),
											path.MatchRelative().AtParent().AtName("weekly"),
										),
									},
									MarkdownDescription: "Hourly interval settings for IMSI updates\n> ℹ️ **Note:** You must specify exactly one of `daily`, `five_minute`, `hourly`, `monthly`, and `weekly`.",
									Optional:            true,
									Attributes:          map[string]schema.Attribute{},
								},
								"monthly": schema.SingleNestedAttribute{
									Validators: []validator.Object{
										objectvalidator.ExactlyOneOf(
											path.MatchRelative().AtParent().AtName("daily"),
											path.MatchRelative().AtParent().AtName("five_minute"),
											path.MatchRelative().AtParent().AtName("hourly"),
											path.MatchRelative().AtParent().AtName("weekly"),
										),
									},
									MarkdownDescription: "Monthly interval settings for IMSI updates\n> ℹ️ **Note:** You must specify exactly one of `daily`, `five_minute`, `hourly`, `monthly`, and `weekly`.",
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"at": schema.StringAttribute{
											Validators: []validator.String{
												stringvalidator.LengthAtMost(2),
												stringvalidator.LengthAtLeast(2),
												stringvalidator.RegexMatches(regexp.MustCompile("([01][0-9]|[2][0-3])"), "pattern must match "+"([01][0-9]|[2][0-3])"),
											},
											MarkdownDescription: "Monthly Time specification hh (e.g. 20) for IMSI",
											Required:            true,
										},
										"day_of_month": schema.Int64Attribute{
											Validators: []validator.Int64{
												int64validator.Between(1, 31),
											},
											MarkdownDescription: "Day of the month for monthly IMSI updates",
											Required:            true,
										},
									},
								},
								"weekly": schema.SingleNestedAttribute{
									Validators: []validator.Object{
										objectvalidator.ExactlyOneOf(
											path.MatchRelative().AtParent().AtName("daily"),
											path.MatchRelative().AtParent().AtName("five_minute"),
											path.MatchRelative().AtParent().AtName("hourly"),
											path.MatchRelative().AtParent().AtName("monthly"),
										),
									},
									MarkdownDescription: "Weekly interval settings for IMSI updates\n> ℹ️ **Note:** You must specify exactly one of `daily`, `five_minute`, `hourly`, `monthly`, and `weekly`.",
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"at": schema.StringAttribute{
											Validators: []validator.String{
												stringvalidator.LengthAtMost(2),
												stringvalidator.LengthAtLeast(2),
												stringvalidator.RegexMatches(regexp.MustCompile("([01][0-9]|[2][0-3])"), "pattern must match "+"([01][0-9]|[2][0-3])"),
											},
											MarkdownDescription: "Weekly Time specification hh (e.g. 20) for IMSI",
											Required:            true,
										},
										"day_of_week": schema.StringAttribute{
											Validators: []validator.String{
												stringvalidator.OneOf("sunday", "monday", "tuesday", "wednesday", "thursday", "friday", "saturday"),
											},
											MarkdownDescription: "Day of week",
											Required:            true,
										},
									},
								},
							},
						},
						"url": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.LengthAtMost(255),
							},
							MarkdownDescription: "IMSI URL for Custom IMSI type",
							Required:            true,
						},
					},
				},
				"ip": schema.SingleNestedAttribute{
					Validators: []validator.Object{
						objectvalidator.ConflictsWith(
							path.MatchRelative().AtParent().AtName("domain"),
							path.MatchRelative().AtParent().AtName("imei"),
							path.MatchRelative().AtParent().AtName("imsi"),
							path.MatchRelative().AtParent().AtName("predefined_ip"),
							path.MatchRelative().AtParent().AtName("predefined_url"),
							path.MatchRelative().AtParent().AtName("url"),
						),
					},
					MarkdownDescription: "IP settings for Custom IP type\n> ℹ️ **Note:** You must specify exactly one of `domain`, `imei`, `imsi`, `ip`, `predefined_ip`, `predefined_url`, and `url`.",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"auth": schema.SingleNestedAttribute{
							MarkdownDescription: "Authentication settings for Custom IP type",
							Optional:            true,
							Attributes: map[string]schema.Attribute{
								"password": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.LengthAtMost(255),
									},
									MarkdownDescription: "Password for Custom IP authentication",
									Required:            true,
									Sensitive:           true,
								},
								"username": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.LengthAtMost(255),
										stringvalidator.LengthAtLeast(1),
									},
									MarkdownDescription: "Username for Custom IP authentication",
									Required:            true,
								},
							},
						},
						"certificate_profile": schema.StringAttribute{
							MarkdownDescription: "Profile for authenticating client certificates",
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("None"),
						},
						"description": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.LengthAtMost(255),
							},
							MarkdownDescription: "Description",
							Optional:            true,
						},
						"exception_list": schema.ListAttribute{
							ElementType:         types.StringType,
							MarkdownDescription: "IP Exception List for Custom IP type",
							Validators: []validator.List{
								listvalidator.ValueStringsAre(stringvalidator.LengthAtMost(255)),
							},
							Optional: true,
						},
						"recurring": schema.SingleNestedAttribute{
							MarkdownDescription: "Update Schedule for Custom IP type",
							Required:            true,
							Attributes: map[string]schema.Attribute{
								"daily": schema.SingleNestedAttribute{
									Validators: []validator.Object{
										objectvalidator.ExactlyOneOf(
											path.MatchRelative().AtParent().AtName("five_minute"),
											path.MatchRelative().AtParent().AtName("hourly"),
											path.MatchRelative().AtParent().AtName("monthly"),
											path.MatchRelative().AtParent().AtName("weekly"),
										),
									},
									MarkdownDescription: "Daily settings for IP recurring\n> ℹ️ **Note:** You must specify exactly one of `daily`, `five_minute`, `hourly`, `monthly`, and `weekly`.",
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"at": schema.StringAttribute{
											Validators: []validator.String{
												stringvalidator.LengthAtMost(2),
												stringvalidator.LengthAtLeast(2),
												stringvalidator.RegexMatches(regexp.MustCompile("([01][0-9]|[2][0-3])"), "pattern must match "+"([01][0-9]|[2][0-3])"),
											},
											MarkdownDescription: "Daily Time specification hh (e.g. 20) for IP",
											Required:            true,
										},
									},
								},
								"five_minute": schema.SingleNestedAttribute{
									Validators: []validator.Object{
										objectvalidator.ExactlyOneOf(
											path.MatchRelative().AtParent().AtName("daily"),
											path.MatchRelative().AtParent().AtName("hourly"),
											path.MatchRelative().AtParent().AtName("monthly"),
											path.MatchRelative().AtParent().AtName("weekly"),
										),
									},
									MarkdownDescription: "Five minute settings for IP recurring\n> ℹ️ **Note:** You must specify exactly one of `daily`, `five_minute`, `hourly`, `monthly`, and `weekly`.",
									Optional:            true,
									Attributes:          map[string]schema.Attribute{},
								},
								"hourly": schema.SingleNestedAttribute{
									Validators: []validator.Object{
										objectvalidator.ExactlyOneOf(
											path.MatchRelative().AtParent().AtName("daily"),
											path.MatchRelative().AtParent().AtName("five_minute"),
											path.MatchRelative().AtParent().AtName("monthly"),
											path.MatchRelative().AtParent().AtName("weekly"),
										),
									},
									MarkdownDescription: "Hourly settings for IP recurring\n> ℹ️ **Note:** You must specify exactly one of `daily`, `five_minute`, `hourly`, `monthly`, and `weekly`.",
									Optional:            true,
									Attributes:          map[string]schema.Attribute{},
								},
								"monthly": schema.SingleNestedAttribute{
									Validators: []validator.Object{
										objectvalidator.ExactlyOneOf(
											path.MatchRelative().AtParent().AtName("daily"),
											path.MatchRelative().AtParent().AtName("five_minute"),
											path.MatchRelative().AtParent().AtName("hourly"),
											path.MatchRelative().AtParent().AtName("weekly"),
										),
									},
									MarkdownDescription: "Monthly settings for IP recurring\n> ℹ️ **Note:** You must specify exactly one of `daily`, `five_minute`, `hourly`, `monthly`, and `weekly`.",
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"at": schema.StringAttribute{
											Validators: []validator.String{
												stringvalidator.LengthAtMost(2),
												stringvalidator.LengthAtLeast(2),
												stringvalidator.RegexMatches(regexp.MustCompile("([01][0-9]|[2][0-3])"), "pattern must match "+"([01][0-9]|[2][0-3])"),
											},
											MarkdownDescription: "Monthly Time specification hh (e.g. 20) for IP",
											Required:            true,
										},
										"day_of_month": schema.Int64Attribute{
											Validators: []validator.Int64{
												int64validator.Between(1, 31),
											},
											MarkdownDescription: "Day setting for monthly IP updates",
											Required:            true,
										},
									},
								},
								"weekly": schema.SingleNestedAttribute{
									Validators: []validator.Object{
										objectvalidator.ExactlyOneOf(
											path.MatchRelative().AtParent().AtName("daily"),
											path.MatchRelative().AtParent().AtName("five_minute"),
											path.MatchRelative().AtParent().AtName("hourly"),
											path.MatchRelative().AtParent().AtName("monthly"),
										),
									},
									MarkdownDescription: "Weekly settings for IP recurring\n> ℹ️ **Note:** You must specify exactly one of `daily`, `five_minute`, `hourly`, `monthly`, and `weekly`.",
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"at": schema.StringAttribute{
											Validators: []validator.String{
												stringvalidator.LengthAtMost(2),
												stringvalidator.LengthAtLeast(2),
												stringvalidator.RegexMatches(regexp.MustCompile("([01][0-9]|[2][0-3])"), "pattern must match "+"([01][0-9]|[2][0-3])"),
											},
											MarkdownDescription: "Weekly Time specification hh (e.g. 20) for IP",
											Required:            true,
										},
										"day_of_week": schema.StringAttribute{
											Validators: []validator.String{
												stringvalidator.OneOf("sunday", "monday", "tuesday", "wednesday", "thursday", "friday", "saturday"),
											},
											MarkdownDescription: "Day of week",
											Required:            true,
										},
									},
								},
							},
						},
						"url": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.LengthAtMost(255),
							},
							MarkdownDescription: "External URL for Custom IP type",
							Required:            true,
						},
					},
				},
				"predefined_ip": schema.SingleNestedAttribute{
					Validators: []validator.Object{
						objectvalidator.ConflictsWith(
							path.MatchRelative().AtParent().AtName("domain"),
							path.MatchRelative().AtParent().AtName("imei"),
							path.MatchRelative().AtParent().AtName("imsi"),
							path.MatchRelative().AtParent().AtName("ip"),
							path.MatchRelative().AtParent().AtName("predefined_url"),
							path.MatchRelative().AtParent().AtName("url"),
						),
					},
					MarkdownDescription: "Predefined IP settings for EDL type\n> ℹ️ **Note:** You must specify exactly one of `domain`, `imei`, `imsi`, `ip`, `predefined_ip`, `predefined_url`, and `url`.",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"description": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.LengthAtMost(255),
							},
							MarkdownDescription: "Description",
							Optional:            true,
						},
						"exception_list": schema.ListAttribute{
							ElementType:         types.StringType,
							MarkdownDescription: "IP Exception List for Predefined IP type",
							Validators: []validator.List{
								listvalidator.ValueStringsAre(stringvalidator.LengthAtMost(255)),
							},
							Optional: true,
						},
						"url": schema.StringAttribute{
							MarkdownDescription: "URL source for Predefined IP type",
							Required:            true,
						},
					},
				},
				"predefined_url": schema.SingleNestedAttribute{
					Validators: []validator.Object{
						objectvalidator.ConflictsWith(
							path.MatchRelative().AtParent().AtName("domain"),
							path.MatchRelative().AtParent().AtName("imei"),
							path.MatchRelative().AtParent().AtName("imsi"),
							path.MatchRelative().AtParent().AtName("ip"),
							path.MatchRelative().AtParent().AtName("predefined_ip"),
							path.MatchRelative().AtParent().AtName("url"),
						),
					},
					MarkdownDescription: "Predefined URL settings for EDL type\n> ℹ️ **Note:** You must specify exactly one of `domain`, `imei`, `imsi`, `ip`, `predefined_ip`, `predefined_url`, and `url`.",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"description": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.LengthAtMost(255),
							},
							MarkdownDescription: "Description",
							Optional:            true,
						},
						"exception_list": schema.ListAttribute{
							ElementType:         types.StringType,
							MarkdownDescription: "URL Exception List for Predefined URL type",
							Validators: []validator.List{
								listvalidator.ValueStringsAre(stringvalidator.LengthAtMost(255)),
							},
							Optional: true,
						},
						"url": schema.StringAttribute{
							MarkdownDescription: "URL source for Predefined URL type",
							Required:            true,
						},
					},
				},
				"url": schema.SingleNestedAttribute{
					Validators: []validator.Object{
						objectvalidator.ConflictsWith(
							path.MatchRelative().AtParent().AtName("domain"),
							path.MatchRelative().AtParent().AtName("imei"),
							path.MatchRelative().AtParent().AtName("imsi"),
							path.MatchRelative().AtParent().AtName("ip"),
							path.MatchRelative().AtParent().AtName("predefined_ip"),
							path.MatchRelative().AtParent().AtName("predefined_url"),
						),
					},
					MarkdownDescription: "URL settings for Custom URL type\n> ℹ️ **Note:** You must specify exactly one of `domain`, `imei`, `imsi`, `ip`, `predefined_ip`, `predefined_url`, and `url`.",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"auth": schema.SingleNestedAttribute{
							MarkdownDescription: "Authentication settings for Custom URL type",
							Optional:            true,
							Attributes: map[string]schema.Attribute{
								"password": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.LengthAtMost(255),
									},
									MarkdownDescription: "Password for Custom URL authentication",
									Required:            true,
									Sensitive:           true,
								},
								"username": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.LengthAtMost(255),
										stringvalidator.LengthAtLeast(1),
									},
									MarkdownDescription: "Username for Custom URL authentication",
									Required:            true,
								},
							},
						},
						"certificate_profile": schema.StringAttribute{
							MarkdownDescription: "Profile for authenticating client certificates",
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("None"),
						},
						"description": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.LengthAtMost(255),
							},
							MarkdownDescription: "Description",
							Optional:            true,
						},
						"exception_list": schema.ListAttribute{
							ElementType:         types.StringType,
							MarkdownDescription: "URL Exception List for Custom URL type",
							Validators: []validator.List{
								listvalidator.ValueStringsAre(stringvalidator.LengthAtMost(255)),
							},
							Optional: true,
						},
						"recurring": schema.SingleNestedAttribute{
							MarkdownDescription: "Update Schedule for Custom URL type",
							Required:            true,
							Attributes: map[string]schema.Attribute{
								"daily": schema.SingleNestedAttribute{
									Validators: []validator.Object{
										objectvalidator.ExactlyOneOf(
											path.MatchRelative().AtParent().AtName("five_minute"),
											path.MatchRelative().AtParent().AtName("hourly"),
											path.MatchRelative().AtParent().AtName("monthly"),
											path.MatchRelative().AtParent().AtName("weekly"),
										),
									},
									MarkdownDescription: "Daily settings for URL recurring\n> ℹ️ **Note:** You must specify exactly one of `daily`, `five_minute`, `hourly`, `monthly`, and `weekly`.",
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"at": schema.StringAttribute{
											Validators: []validator.String{
												stringvalidator.LengthAtMost(2),
												stringvalidator.LengthAtLeast(2),
												stringvalidator.RegexMatches(regexp.MustCompile("([01][0-9]|[2][0-3])"), "pattern must match "+"([01][0-9]|[2][0-3])"),
											},
											MarkdownDescription: "Daily Time specification hh (e.g. 20) for URL",
											Required:            true,
										},
									},
								},
								"five_minute": schema.SingleNestedAttribute{
									Validators: []validator.Object{
										objectvalidator.ExactlyOneOf(
											path.MatchRelative().AtParent().AtName("daily"),
											path.MatchRelative().AtParent().AtName("hourly"),
											path.MatchRelative().AtParent().AtName("monthly"),
											path.MatchRelative().AtParent().AtName("weekly"),
										),
									},
									MarkdownDescription: "Five minute settings for URL recurring\n> ℹ️ **Note:** You must specify exactly one of `daily`, `five_minute`, `hourly`, `monthly`, and `weekly`.",
									Optional:            true,
									Attributes:          map[string]schema.Attribute{},
								},
								"hourly": schema.SingleNestedAttribute{
									Validators: []validator.Object{
										objectvalidator.ExactlyOneOf(
											path.MatchRelative().AtParent().AtName("daily"),
											path.MatchRelative().AtParent().AtName("five_minute"),
											path.MatchRelative().AtParent().AtName("monthly"),
											path.MatchRelative().AtParent().AtName("weekly"),
										),
									},
									MarkdownDescription: "Hourly settings for URL recurring\n> ℹ️ **Note:** You must specify exactly one of `daily`, `five_minute`, `hourly`, `monthly`, and `weekly`.",
									Optional:            true,
									Attributes:          map[string]schema.Attribute{},
								},
								"monthly": schema.SingleNestedAttribute{
									Validators: []validator.Object{
										objectvalidator.ExactlyOneOf(
											path.MatchRelative().AtParent().AtName("daily"),
											path.MatchRelative().AtParent().AtName("five_minute"),
											path.MatchRelative().AtParent().AtName("hourly"),
											path.MatchRelative().AtParent().AtName("weekly"),
										),
									},
									MarkdownDescription: "Monthly settings for URL recurring\n> ℹ️ **Note:** You must specify exactly one of `daily`, `five_minute`, `hourly`, `monthly`, and `weekly`.",
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"at": schema.StringAttribute{
											Validators: []validator.String{
												stringvalidator.LengthAtMost(2),
												stringvalidator.LengthAtLeast(2),
												stringvalidator.RegexMatches(regexp.MustCompile("([01][0-9]|[2][0-3])"), "pattern must match "+"([01][0-9]|[2][0-3])"),
											},
											MarkdownDescription: "Monthly Time specification hh (e.g. 20) for URL",
											Required:            true,
										},
										"day_of_month": schema.Int64Attribute{
											Validators: []validator.Int64{
												int64validator.Between(1, 31),
											},
											MarkdownDescription: "Day setting for monthly URL updates",
											Required:            true,
										},
									},
								},
								"weekly": schema.SingleNestedAttribute{
									Validators: []validator.Object{
										objectvalidator.ExactlyOneOf(
											path.MatchRelative().AtParent().AtName("daily"),
											path.MatchRelative().AtParent().AtName("five_minute"),
											path.MatchRelative().AtParent().AtName("hourly"),
											path.MatchRelative().AtParent().AtName("monthly"),
										),
									},
									MarkdownDescription: "Weekly settings for URL recurring\n> ℹ️ **Note:** You must specify exactly one of `daily`, `five_minute`, `hourly`, `monthly`, and `weekly`.",
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"at": schema.StringAttribute{
											Validators: []validator.String{
												stringvalidator.LengthAtMost(2),
												stringvalidator.LengthAtLeast(2),
												stringvalidator.RegexMatches(regexp.MustCompile("([01][0-9]|[2][0-3])"), "pattern must match "+"([01][0-9]|[2][0-3])"),
											},
											MarkdownDescription: "Weekly Time specification hh (e.g. 20) for URL",
											Required:            true,
										},
										"day_of_week": schema.StringAttribute{
											Validators: []validator.String{
												stringvalidator.OneOf("sunday", "monday", "tuesday", "wednesday", "thursday", "friday", "saturday"),
											},
											MarkdownDescription: "Day of week",
											Required:            true,
										},
									},
								},
							},
						},
						"url": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.LengthAtMost(255),
							},
							MarkdownDescription: "External URL for Custom URL type",
							Required:            true,
						},
					},
				},
			},
		},
	},
}

// ExternalDynamicListsDataSourceSchema defines the schema for ExternalDynamicLists data source
var ExternalDynamicListsDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "ExternalDynamicList data source",
	Attributes: map[string]dsschema.Attribute{
		"device": dsschema.StringAttribute{
			MarkdownDescription: "The device in which the resource is defined\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
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
			MarkdownDescription: "The folder in which the resource is defined\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			Computed:            true,
		},
		"id": dsschema.StringAttribute{
			MarkdownDescription: "The UUID of the external dynamic list",
			Required:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "The name of the external dynamic list",
			Optional:            true,
			Computed:            true,
		},
		"snippet": dsschema.StringAttribute{
			MarkdownDescription: "The snippet in which the resource is defined\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
		"type": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Type configuration for External Dynamic List",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"domain": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Domain settings for Custom Domain type\n> ℹ️ **Note:** You must specify exactly one of `domain`, `imei`, `imsi`, `ip`, `predefined_ip`, `predefined_url`, and `url`.",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"auth": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Authentication settings for Custom Domain type",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"password": dsschema.StringAttribute{
									MarkdownDescription: "Password for Custom Domain authentication",
									Computed:            true,
									Sensitive:           true,
								},
								"username": dsschema.StringAttribute{
									MarkdownDescription: "Username for Custom Domain authentication",
									Computed:            true,
								},
							},
						},
						"certificate_profile": dsschema.StringAttribute{
							MarkdownDescription: "Profile for authenticating client certificates",
							Computed:            true,
						},
						"description": dsschema.StringAttribute{
							MarkdownDescription: "Description",
							Computed:            true,
						},
						"exception_list": dsschema.ListAttribute{
							ElementType:         types.StringType,
							MarkdownDescription: "Domain Exception List for Custom Domain type",
							Computed:            true,
						},
						"expand_domain": dsschema.BoolAttribute{
							MarkdownDescription: "Enable/Disable expand domain",
							Computed:            true,
						},
						"recurring": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Update Schedule for Custom Domain type",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"daily": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Daily settings for Domain recurring\n> ℹ️ **Note:** You must specify exactly one of `daily`, `five_minute`, `hourly`, `monthly`, and `weekly`.",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"at": dsschema.StringAttribute{
											MarkdownDescription: "Daily Time specification hh (e.g. 20) for Domain",
											Computed:            true,
										},
									},
								},
								"five_minute": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Five minute settings for Domain recurring\n> ℹ️ **Note:** You must specify exactly one of `daily`, `five_minute`, `hourly`, `monthly`, and `weekly`.",
									Computed:            true,
									Attributes:          map[string]dsschema.Attribute{},
								},
								"hourly": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Hourly settings for Domain recurring\n> ℹ️ **Note:** You must specify exactly one of `daily`, `five_minute`, `hourly`, `monthly`, and `weekly`.",
									Computed:            true,
									Attributes:          map[string]dsschema.Attribute{},
								},
								"monthly": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Monthly settings for Domain recurring\n> ℹ️ **Note:** You must specify exactly one of `daily`, `five_minute`, `hourly`, `monthly`, and `weekly`.",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"at": dsschema.StringAttribute{
											MarkdownDescription: "Monthly Time specification hh (e.g. 20) for domain",
											Computed:            true,
										},
										"day_of_month": dsschema.Int64Attribute{
											MarkdownDescription: "Day setting for monthly Domain updates",
											Computed:            true,
										},
									},
								},
								"weekly": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Weekly settings for Domain recurring\n> ℹ️ **Note:** You must specify exactly one of `daily`, `five_minute`, `hourly`, `monthly`, and `weekly`.",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"at": dsschema.StringAttribute{
											MarkdownDescription: "Weekly Time specification hh (e.g. 20) for Domain",
											Computed:            true,
										},
										"day_of_week": dsschema.StringAttribute{
											MarkdownDescription: "Day of week",
											Computed:            true,
										},
									},
								},
							},
						},
						"url": dsschema.StringAttribute{
							MarkdownDescription: "External URL for Custom Domain type",
							Computed:            true,
						},
					},
				},
				"imei": dsschema.SingleNestedAttribute{
					MarkdownDescription: "IMEI Configuration settings\n> ℹ️ **Note:** You must specify exactly one of `domain`, `imei`, `imsi`, `ip`, `predefined_ip`, `predefined_url`, and `url`.",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"auth": dsschema.SingleNestedAttribute{
							MarkdownDescription: "IMEI Auth Cnfig for Custom IMEI type",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"password": dsschema.StringAttribute{
									MarkdownDescription: "IMEI Auth Password for Custom IMEI type",
									Computed:            true,
									Sensitive:           true,
								},
								"username": dsschema.StringAttribute{
									MarkdownDescription: "IMEI Auth username for Custom IMEI type",
									Computed:            true,
								},
							},
						},
						"certificate_profile": dsschema.StringAttribute{
							MarkdownDescription: "IMEI Certificate Profile for Custom IMEI type",
							Computed:            true,
						},
						"description": dsschema.StringAttribute{
							MarkdownDescription: "IMEI Description for Custom IMEI type",
							Computed:            true,
						},
						"exception_list": dsschema.ListAttribute{
							ElementType:         types.StringType,
							MarkdownDescription: "IMEI Exception List for Custom IMEI type",
							Computed:            true,
						},
						"recurring": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Recurring interval for IMEI updates",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"daily": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Daily interval settings for IMEI updates\n> ℹ️ **Note:** You must specify exactly one of `daily`, `five_minute`, `hourly`, `monthly`, and `weekly`.",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"at": dsschema.StringAttribute{
											MarkdownDescription: "Daily Time specification hh (e.g. 20) for IMEI",
											Computed:            true,
										},
									},
								},
								"five_minute": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Five-minute interval settings for IMEI updates\n> ℹ️ **Note:** You must specify exactly one of `daily`, `five_minute`, `hourly`, `monthly`, and `weekly`.",
									Computed:            true,
									Attributes:          map[string]dsschema.Attribute{},
								},
								"hourly": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Hourly interval settings for IMEI updates\n> ℹ️ **Note:** You must specify exactly one of `daily`, `five_minute`, `hourly`, `monthly`, and `weekly`.",
									Computed:            true,
									Attributes:          map[string]dsschema.Attribute{},
								},
								"monthly": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Monthly interval settings for IMEI updates\n> ℹ️ **Note:** You must specify exactly one of `daily`, `five_minute`, `hourly`, `monthly`, and `weekly`.",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"at": dsschema.StringAttribute{
											MarkdownDescription: "Monthly Time specification hh (e.g. 20) for IMEI",
											Computed:            true,
										},
										"day_of_month": dsschema.Int64Attribute{
											MarkdownDescription: "Day of month for IMEI updates",
											Computed:            true,
										},
									},
								},
								"weekly": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Weekly interval settings for IMEI updates\n> ℹ️ **Note:** You must specify exactly one of `daily`, `five_minute`, `hourly`, `monthly`, and `weekly`.",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"at": dsschema.StringAttribute{
											MarkdownDescription: "Weekly Time specification hh (e.g. 20) for IMEI",
											Computed:            true,
										},
										"day_of_week": dsschema.StringAttribute{
											MarkdownDescription: "Day of week",
											Computed:            true,
										},
									},
								},
							},
						},
						"url": dsschema.StringAttribute{
							MarkdownDescription: "IMEI URL for Custom IMEI type",
							Computed:            true,
						},
					},
				},
				"imsi": dsschema.SingleNestedAttribute{
					MarkdownDescription: "IMSI Config for Custom IMSI type\n> ℹ️ **Note:** You must specify exactly one of `domain`, `imei`, `imsi`, `ip`, `predefined_ip`, `predefined_url`, and `url`.",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"auth": dsschema.SingleNestedAttribute{
							MarkdownDescription: "IMSI Auth Config for Custom IMSI type",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"password": dsschema.StringAttribute{
									MarkdownDescription: "IMSI Auth Password for Custom IMSI type",
									Computed:            true,
									Sensitive:           true,
								},
								"username": dsschema.StringAttribute{
									MarkdownDescription: "IMSI Auth Username for Custom IMSI type",
									Computed:            true,
								},
							},
						},
						"certificate_profile": dsschema.StringAttribute{
							MarkdownDescription: "IMSI Certificate Profile for Custom IMSI type",
							Computed:            true,
						},
						"description": dsschema.StringAttribute{
							MarkdownDescription: "IMSI Description for Custom IMSI type",
							Computed:            true,
						},
						"exception_list": dsschema.ListAttribute{
							ElementType:         types.StringType,
							MarkdownDescription: "IMSI Exception List for Custom IMSI type",
							Computed:            true,
						},
						"recurring": dsschema.SingleNestedAttribute{
							MarkdownDescription: "IMSI Recuring Config for Custom IMSI type",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"daily": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Daily interval settings for IMSI updates\n> ℹ️ **Note:** You must specify exactly one of `daily`, `five_minute`, `hourly`, `monthly`, and `weekly`.",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"at": dsschema.StringAttribute{
											MarkdownDescription: "Daily Time specification hh (e.g. 20) for IMSI",
											Computed:            true,
										},
									},
								},
								"five_minute": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Five-minute interval settings for IMSI updates\n> ℹ️ **Note:** You must specify exactly one of `daily`, `five_minute`, `hourly`, `monthly`, and `weekly`.",
									Computed:            true,
									Attributes:          map[string]dsschema.Attribute{},
								},
								"hourly": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Hourly interval settings for IMSI updates\n> ℹ️ **Note:** You must specify exactly one of `daily`, `five_minute`, `hourly`, `monthly`, and `weekly`.",
									Computed:            true,
									Attributes:          map[string]dsschema.Attribute{},
								},
								"monthly": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Monthly interval settings for IMSI updates\n> ℹ️ **Note:** You must specify exactly one of `daily`, `five_minute`, `hourly`, `monthly`, and `weekly`.",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"at": dsschema.StringAttribute{
											MarkdownDescription: "Monthly Time specification hh (e.g. 20) for IMSI",
											Computed:            true,
										},
										"day_of_month": dsschema.Int64Attribute{
											MarkdownDescription: "Day of the month for monthly IMSI updates",
											Computed:            true,
										},
									},
								},
								"weekly": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Weekly interval settings for IMSI updates\n> ℹ️ **Note:** You must specify exactly one of `daily`, `five_minute`, `hourly`, `monthly`, and `weekly`.",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"at": dsschema.StringAttribute{
											MarkdownDescription: "Weekly Time specification hh (e.g. 20) for IMSI",
											Computed:            true,
										},
										"day_of_week": dsschema.StringAttribute{
											MarkdownDescription: "Day of week",
											Computed:            true,
										},
									},
								},
							},
						},
						"url": dsschema.StringAttribute{
							MarkdownDescription: "IMSI URL for Custom IMSI type",
							Computed:            true,
						},
					},
				},
				"ip": dsschema.SingleNestedAttribute{
					MarkdownDescription: "IP settings for Custom IP type\n> ℹ️ **Note:** You must specify exactly one of `domain`, `imei`, `imsi`, `ip`, `predefined_ip`, `predefined_url`, and `url`.",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"auth": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Authentication settings for Custom IP type",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"password": dsschema.StringAttribute{
									MarkdownDescription: "Password for Custom IP authentication",
									Computed:            true,
									Sensitive:           true,
								},
								"username": dsschema.StringAttribute{
									MarkdownDescription: "Username for Custom IP authentication",
									Computed:            true,
								},
							},
						},
						"certificate_profile": dsschema.StringAttribute{
							MarkdownDescription: "Profile for authenticating client certificates",
							Computed:            true,
						},
						"description": dsschema.StringAttribute{
							MarkdownDescription: "Description",
							Computed:            true,
						},
						"exception_list": dsschema.ListAttribute{
							ElementType:         types.StringType,
							MarkdownDescription: "IP Exception List for Custom IP type",
							Computed:            true,
						},
						"recurring": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Update Schedule for Custom IP type",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"daily": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Daily settings for IP recurring\n> ℹ️ **Note:** You must specify exactly one of `daily`, `five_minute`, `hourly`, `monthly`, and `weekly`.",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"at": dsschema.StringAttribute{
											MarkdownDescription: "Daily Time specification hh (e.g. 20) for IP",
											Computed:            true,
										},
									},
								},
								"five_minute": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Five minute settings for IP recurring\n> ℹ️ **Note:** You must specify exactly one of `daily`, `five_minute`, `hourly`, `monthly`, and `weekly`.",
									Computed:            true,
									Attributes:          map[string]dsschema.Attribute{},
								},
								"hourly": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Hourly settings for IP recurring\n> ℹ️ **Note:** You must specify exactly one of `daily`, `five_minute`, `hourly`, `monthly`, and `weekly`.",
									Computed:            true,
									Attributes:          map[string]dsschema.Attribute{},
								},
								"monthly": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Monthly settings for IP recurring\n> ℹ️ **Note:** You must specify exactly one of `daily`, `five_minute`, `hourly`, `monthly`, and `weekly`.",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"at": dsschema.StringAttribute{
											MarkdownDescription: "Monthly Time specification hh (e.g. 20) for IP",
											Computed:            true,
										},
										"day_of_month": dsschema.Int64Attribute{
											MarkdownDescription: "Day setting for monthly IP updates",
											Computed:            true,
										},
									},
								},
								"weekly": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Weekly settings for IP recurring\n> ℹ️ **Note:** You must specify exactly one of `daily`, `five_minute`, `hourly`, `monthly`, and `weekly`.",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"at": dsschema.StringAttribute{
											MarkdownDescription: "Weekly Time specification hh (e.g. 20) for IP",
											Computed:            true,
										},
										"day_of_week": dsschema.StringAttribute{
											MarkdownDescription: "Day of week",
											Computed:            true,
										},
									},
								},
							},
						},
						"url": dsschema.StringAttribute{
							MarkdownDescription: "External URL for Custom IP type",
							Computed:            true,
						},
					},
				},
				"predefined_ip": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Predefined IP settings for EDL type\n> ℹ️ **Note:** You must specify exactly one of `domain`, `imei`, `imsi`, `ip`, `predefined_ip`, `predefined_url`, and `url`.",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"description": dsschema.StringAttribute{
							MarkdownDescription: "Description",
							Computed:            true,
						},
						"exception_list": dsschema.ListAttribute{
							ElementType:         types.StringType,
							MarkdownDescription: "IP Exception List for Predefined IP type",
							Computed:            true,
						},
						"url": dsschema.StringAttribute{
							MarkdownDescription: "URL source for Predefined IP type",
							Computed:            true,
						},
					},
				},
				"predefined_url": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Predefined URL settings for EDL type\n> ℹ️ **Note:** You must specify exactly one of `domain`, `imei`, `imsi`, `ip`, `predefined_ip`, `predefined_url`, and `url`.",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"description": dsschema.StringAttribute{
							MarkdownDescription: "Description",
							Computed:            true,
						},
						"exception_list": dsschema.ListAttribute{
							ElementType:         types.StringType,
							MarkdownDescription: "URL Exception List for Predefined URL type",
							Computed:            true,
						},
						"url": dsschema.StringAttribute{
							MarkdownDescription: "URL source for Predefined URL type",
							Computed:            true,
						},
					},
				},
				"url": dsschema.SingleNestedAttribute{
					MarkdownDescription: "URL settings for Custom URL type\n> ℹ️ **Note:** You must specify exactly one of `domain`, `imei`, `imsi`, `ip`, `predefined_ip`, `predefined_url`, and `url`.",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"auth": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Authentication settings for Custom URL type",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"password": dsschema.StringAttribute{
									MarkdownDescription: "Password for Custom URL authentication",
									Computed:            true,
									Sensitive:           true,
								},
								"username": dsschema.StringAttribute{
									MarkdownDescription: "Username for Custom URL authentication",
									Computed:            true,
								},
							},
						},
						"certificate_profile": dsschema.StringAttribute{
							MarkdownDescription: "Profile for authenticating client certificates",
							Computed:            true,
						},
						"description": dsschema.StringAttribute{
							MarkdownDescription: "Description",
							Computed:            true,
						},
						"exception_list": dsschema.ListAttribute{
							ElementType:         types.StringType,
							MarkdownDescription: "URL Exception List for Custom URL type",
							Computed:            true,
						},
						"recurring": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Update Schedule for Custom URL type",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"daily": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Daily settings for URL recurring\n> ℹ️ **Note:** You must specify exactly one of `daily`, `five_minute`, `hourly`, `monthly`, and `weekly`.",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"at": dsschema.StringAttribute{
											MarkdownDescription: "Daily Time specification hh (e.g. 20) for URL",
											Computed:            true,
										},
									},
								},
								"five_minute": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Five minute settings for URL recurring\n> ℹ️ **Note:** You must specify exactly one of `daily`, `five_minute`, `hourly`, `monthly`, and `weekly`.",
									Computed:            true,
									Attributes:          map[string]dsschema.Attribute{},
								},
								"hourly": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Hourly settings for URL recurring\n> ℹ️ **Note:** You must specify exactly one of `daily`, `five_minute`, `hourly`, `monthly`, and `weekly`.",
									Computed:            true,
									Attributes:          map[string]dsschema.Attribute{},
								},
								"monthly": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Monthly settings for URL recurring\n> ℹ️ **Note:** You must specify exactly one of `daily`, `five_minute`, `hourly`, `monthly`, and `weekly`.",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"at": dsschema.StringAttribute{
											MarkdownDescription: "Monthly Time specification hh (e.g. 20) for URL",
											Computed:            true,
										},
										"day_of_month": dsschema.Int64Attribute{
											MarkdownDescription: "Day setting for monthly URL updates",
											Computed:            true,
										},
									},
								},
								"weekly": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Weekly settings for URL recurring\n> ℹ️ **Note:** You must specify exactly one of `daily`, `five_minute`, `hourly`, `monthly`, and `weekly`.",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"at": dsschema.StringAttribute{
											MarkdownDescription: "Weekly Time specification hh (e.g. 20) for URL",
											Computed:            true,
										},
										"day_of_week": dsschema.StringAttribute{
											MarkdownDescription: "Day of week",
											Computed:            true,
										},
									},
								},
							},
						},
						"url": dsschema.StringAttribute{
							MarkdownDescription: "External URL for Custom URL type",
							Computed:            true,
						},
					},
				},
			},
		},
	},
}

// ExternalDynamicListsListModel represents the data model for a list data source.
type ExternalDynamicListsListModel struct {
	Tfid    types.String           `tfsdk:"tfid"`
	Data    []ExternalDynamicLists `tfsdk:"data"`
	Limit   types.Int64            `tfsdk:"limit"`
	Offset  types.Int64            `tfsdk:"offset"`
	Name    types.String           `tfsdk:"name"`
	Total   types.Int64            `tfsdk:"total"`
	Folder  types.String           `tfsdk:"folder"`
	Snippet types.String           `tfsdk:"snippet"`
	Device  types.String           `tfsdk:"device"`
}

// ExternalDynamicListsListDataSourceSchema defines the schema for a list data source.
var ExternalDynamicListsListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: ExternalDynamicListsDataSourceSchema.Attributes,
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
