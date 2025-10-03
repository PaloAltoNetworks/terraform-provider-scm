package models

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/boolvalidator"
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

// Package: network_services
// This file contains models for the network_services SDK package

// NatRules represents the Terraform model for NatRules
type NatRules struct {
	Tfid                      types.String          `tfsdk:"tfid"`
	ActiveActiveDeviceBinding basetypes.StringValue `tfsdk:"active_active_device_binding"`
	Description               basetypes.StringValue `tfsdk:"description"`
	Destination               basetypes.ListValue   `tfsdk:"destination"`
	Device                    basetypes.StringValue `tfsdk:"device"`
	Disabled                  basetypes.BoolValue   `tfsdk:"disabled"`
	Distribution              basetypes.StringValue `tfsdk:"distribution"`
	DnsRewrite                basetypes.ObjectValue `tfsdk:"dns_rewrite"`
	Folder                    basetypes.StringValue `tfsdk:"folder"`
	From                      basetypes.ListValue   `tfsdk:"from"`
	Id                        basetypes.StringValue `tfsdk:"id"`
	Name                      basetypes.StringValue `tfsdk:"name"`
	NatType                   basetypes.StringValue `tfsdk:"nat_type"`
	Service                   basetypes.StringValue `tfsdk:"service"`
	Snippet                   basetypes.StringValue `tfsdk:"snippet"`
	Source                    basetypes.ListValue   `tfsdk:"source"`
	SourceTranslation         basetypes.ObjectValue `tfsdk:"source_translation"`
	Tag                       basetypes.ListValue   `tfsdk:"tag"`
	To                        basetypes.ListValue   `tfsdk:"to"`
	ToInterface               basetypes.StringValue `tfsdk:"to_interface"`
	TranslatedAddressSingle   basetypes.StringValue `tfsdk:"translated_address_single"`
	TranslatedPort            basetypes.Int64Value  `tfsdk:"translated_port"`
	Position                  basetypes.StringValue `tfsdk:"position"`
}

// NatRulesDnsRewrite represents a nested structure within the NatRules model
type NatRulesDnsRewrite struct {
	Direction basetypes.StringValue `tfsdk:"direction"`
}

// NatRulesSourceTranslation represents a nested structure within the NatRules model
type NatRulesSourceTranslation struct {
	BiDirectional           basetypes.BoolValue   `tfsdk:"bi_directional"`
	Fallback                basetypes.ObjectValue `tfsdk:"fallback"`
	Interface               basetypes.StringValue `tfsdk:"interface"`
	TranslatedAddressArray  basetypes.ListValue   `tfsdk:"translated_address_array"`
	TranslatedAddressSingle basetypes.StringValue `tfsdk:"translated_address_single"`
}

// NatRulesSourceTranslationFallback represents a nested structure within the NatRules model
type NatRulesSourceTranslationFallback struct {
	Interface basetypes.StringValue `tfsdk:"interface"`
}

// AttrTypes defines the attribute types for the NatRules model.
func (o NatRules) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":                         basetypes.StringType{},
		"active_active_device_binding": basetypes.StringType{},
		"description":                  basetypes.StringType{},
		"destination":                  basetypes.ListType{ElemType: basetypes.StringType{}},
		"device":                       basetypes.StringType{},
		"disabled":                     basetypes.BoolType{},
		"distribution":                 basetypes.StringType{},
		"dns_rewrite": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"direction": basetypes.StringType{},
			},
		},
		"folder":   basetypes.StringType{},
		"from":     basetypes.ListType{ElemType: basetypes.StringType{}},
		"id":       basetypes.StringType{},
		"name":     basetypes.StringType{},
		"nat_type": basetypes.StringType{},
		"service":  basetypes.StringType{},
		"snippet":  basetypes.StringType{},
		"source":   basetypes.ListType{ElemType: basetypes.StringType{}},
		"source_translation": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"bi_directional": basetypes.BoolType{},
				"fallback": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"interface": basetypes.StringType{},
					},
				},
				"interface":                 basetypes.StringType{},
				"translated_address_array":  basetypes.ListType{ElemType: basetypes.StringType{}},
				"translated_address_single": basetypes.StringType{},
			},
		},
		"tag":                       basetypes.ListType{ElemType: basetypes.StringType{}},
		"to":                        basetypes.ListType{ElemType: basetypes.StringType{}},
		"to_interface":              basetypes.StringType{},
		"translated_address_single": basetypes.StringType{},
		"translated_port":           basetypes.Int64Type{},
		"position":                  basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of NatRules objects.
func (o NatRules) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the NatRulesDnsRewrite model.
func (o NatRulesDnsRewrite) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"direction": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of NatRulesDnsRewrite objects.
func (o NatRulesDnsRewrite) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the NatRulesSourceTranslation model.
func (o NatRulesSourceTranslation) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"bi_directional": basetypes.BoolType{},
		"fallback": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"interface": basetypes.StringType{},
			},
		},
		"interface":                 basetypes.StringType{},
		"translated_address_array":  basetypes.ListType{ElemType: basetypes.StringType{}},
		"translated_address_single": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of NatRulesSourceTranslation objects.
func (o NatRulesSourceTranslation) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the NatRulesSourceTranslationFallback model.
func (o NatRulesSourceTranslationFallback) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"interface": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of NatRulesSourceTranslationFallback objects.
func (o NatRulesSourceTranslationFallback) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// NatRulesResourceSchema defines the schema for NatRules resource
var NatRulesResourceSchema = schema.Schema{
	MarkdownDescription: "NatRule resource",
	Attributes: map[string]schema.Attribute{
		"active_active_device_binding": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.OneOf("primary", "both", "0", "1"),
			},
			MarkdownDescription: "Active active device binding",
			Optional:            true,
		},
		"description": schema.StringAttribute{
			MarkdownDescription: "NAT rule description",
			Optional:            true,
		},
		"destination": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Destination address(es) of the original packet",
			Required:            true,
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
		"disabled": schema.BoolAttribute{
			MarkdownDescription: "Disable NAT rule?",
			Optional:            true,
			Computed:            true,
			Default:             booldefault.StaticBool(false),
		},
		"distribution": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.OneOf("round-robin", "source-ip-hash", "ip-modulo", "ip-hash", "least-sessions"),
			},
			MarkdownDescription: "Distribution method",
			Optional:            true,
		},
		"dns_rewrite": schema.SingleNestedAttribute{
			MarkdownDescription: "DNS rewrite",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"direction": schema.StringAttribute{
					Validators: []validator.String{
						stringvalidator.OneOf("reverse", "forward"),
					},
					MarkdownDescription: "Direction",
					Optional:            true,
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
		"from": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Source zone(s) of the original packet",
			Required:            true,
		},
		"id": schema.StringAttribute{
			MarkdownDescription: "UUID of the resource",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"name": schema.StringAttribute{
			MarkdownDescription: "NAT rule name",
			Required:            true,
		},
		"nat_type": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.OneOf("ipv4", "nat64", "nptv6"),
			},
			MarkdownDescription: "NAT type",
			Optional:            true,
			Computed:            true,
			Default:             stringdefault.StaticString("ipv4"),
		},
		"position": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.OneOf("pre", "post"),
			},
			MarkdownDescription: "The relative position of the rule",
			Optional:            true,
			Computed:            true,
			Default:             stringdefault.StaticString("pre"),
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"service": schema.StringAttribute{
			MarkdownDescription: "The service of the original packet",
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
		"source": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Source address(es) of the original packet",
			Required:            true,
		},
		"source_translation": schema.SingleNestedAttribute{
			MarkdownDescription: "Source translation",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"bi_directional": schema.BoolAttribute{
					Validators: []validator.Bool{
						boolvalidator.ExactlyOneOf(
							path.MatchRelative().AtParent().AtName("interface"),
							path.MatchRelative().AtParent().AtName("fallback"),
							path.MatchRelative().AtParent().AtName("translated_address_array"),
							path.MatchRelative().AtParent().AtName("translated_address_single"),
						),
					},
					MarkdownDescription: "Bi directional",
					Optional:            true,
				},
				"fallback": schema.SingleNestedAttribute{
					Validators: []validator.Object{
						objectvalidator.ExactlyOneOf(
							path.MatchRelative().AtParent().AtName("interface"),
							path.MatchRelative().AtParent().AtName("translated_address_array"),
							path.MatchRelative().AtParent().AtName("bi_directional"),
							path.MatchRelative().AtParent().AtName("translated_address_single"),
						),
					},
					MarkdownDescription: "Fallback",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"interface": schema.StringAttribute{
							MarkdownDescription: "Interface name",
							Optional:            true,
						},
					},
				},
				"interface": schema.StringAttribute{
					Validators: []validator.String{
						stringvalidator.ExactlyOneOf(
							path.MatchRelative().AtParent().AtName("fallback"),
							path.MatchRelative().AtParent().AtName("translated_address_array"),
							path.MatchRelative().AtParent().AtName("bi_directional"),
							path.MatchRelative().AtParent().AtName("translated_address_single"),
						),
					},
					MarkdownDescription: "Interface name",
					Optional:            true,
				},
				"translated_address_array": schema.ListAttribute{
					ElementType:         types.StringType,
					MarkdownDescription: "Translated IP addresses",
					Validators: []validator.List{
						listvalidator.ExactlyOneOf(
							path.MatchRelative().AtParent().AtName("interface"),
							path.MatchRelative().AtParent().AtName("fallback"),
							path.MatchRelative().AtParent().AtName("bi_directional"),
							path.MatchRelative().AtParent().AtName("translated_address_single"),
						),
					},
					Optional: true,
				},
				"translated_address_single": schema.StringAttribute{
					Validators: []validator.String{
						stringvalidator.ExactlyOneOf(
							path.MatchRelative().AtParent().AtName("interface"),
							path.MatchRelative().AtParent().AtName("fallback"),
							path.MatchRelative().AtParent().AtName("translated_address_array"),
							path.MatchRelative().AtParent().AtName("bi_directional"),
						),
					},
					MarkdownDescription: "Translated IP address",
					Optional:            true,
				},
			},
		},
		"tag": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "NAT rule tags",
			Optional:            true,
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
			MarkdownDescription: "Destination zone of the original packet",
			Required:            true,
		},
		"to_interface": schema.StringAttribute{
			MarkdownDescription: "Destination interface of the original packet",
			Optional:            true,
		},
		"translated_address_single": schema.StringAttribute{
			MarkdownDescription: "Translated destination IP address",
			Optional:            true,
		},
		"translated_port": schema.Int64Attribute{
			Validators: []validator.Int64{
				int64validator.Between(1, 65535),
			},
			MarkdownDescription: "Translated destination port",
			Optional:            true,
		},
	},
}

// NatRulesDataSourceSchema defines the schema for NatRules data source
var NatRulesDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "NatRule data source",
	Attributes: map[string]dsschema.Attribute{
		"active_active_device_binding": dsschema.StringAttribute{
			MarkdownDescription: "Active active device binding",
			Computed:            true,
		},
		"description": dsschema.StringAttribute{
			MarkdownDescription: "NAT rule description",
			Computed:            true,
		},
		"destination": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Destination address(es) of the original packet",
			Computed:            true,
		},
		"device": dsschema.StringAttribute{
			MarkdownDescription: "The device in which the resource is defined",
			Computed:            true,
		},
		"disabled": dsschema.BoolAttribute{
			MarkdownDescription: "Disable NAT rule?",
			Computed:            true,
		},
		"distribution": dsschema.StringAttribute{
			MarkdownDescription: "Distribution method",
			Computed:            true,
		},
		"dns_rewrite": dsschema.SingleNestedAttribute{
			MarkdownDescription: "DNS rewrite",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"direction": dsschema.StringAttribute{
					MarkdownDescription: "Direction",
					Computed:            true,
				},
			},
		},
		"folder": dsschema.StringAttribute{
			MarkdownDescription: "The folder in which the resource is defined",
			Computed:            true,
		},
		"from": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Source zone(s) of the original packet",
			Computed:            true,
		},
		"id": dsschema.StringAttribute{
			MarkdownDescription: "UUID of the resource",
			Required:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "NAT rule name",
			Optional:            true,
			Computed:            true,
		},
		"nat_type": dsschema.StringAttribute{
			MarkdownDescription: "NAT type",
			Computed:            true,
		},
		"service": dsschema.StringAttribute{
			MarkdownDescription: "The service of the original packet",
			Computed:            true,
		},
		"snippet": dsschema.StringAttribute{
			MarkdownDescription: "The snippet in which the resource is defined",
			Computed:            true,
		},
		"source": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Source address(es) of the original packet",
			Computed:            true,
		},
		"source_translation": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Source translation",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"bi_directional": dsschema.BoolAttribute{
					MarkdownDescription: "Bi directional",
					Computed:            true,
				},
				"fallback": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Fallback",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"interface": dsschema.StringAttribute{
							MarkdownDescription: "Interface name",
							Computed:            true,
						},
					},
				},
				"interface": dsschema.StringAttribute{
					MarkdownDescription: "Interface name",
					Computed:            true,
				},
				"translated_address_array": dsschema.ListAttribute{
					ElementType:         types.StringType,
					MarkdownDescription: "Translated IP addresses",
					Computed:            true,
				},
				"translated_address_single": dsschema.StringAttribute{
					MarkdownDescription: "Translated IP address",
					Computed:            true,
				},
			},
		},
		"tag": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "NAT rule tags",
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
		"to": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Destination zone of the original packet",
			Computed:            true,
		},
		"to_interface": dsschema.StringAttribute{
			MarkdownDescription: "Destination interface of the original packet",
			Computed:            true,
		},
		"translated_address_single": dsschema.StringAttribute{
			MarkdownDescription: "Translated destination IP address",
			Computed:            true,
		},
		"translated_port": dsschema.Int64Attribute{
			MarkdownDescription: "Translated destination port",
			Computed:            true,
		},
	},
}

// NatRulesListModel represents the data model for a list data source.
type NatRulesListModel struct {
	Tfid    types.String `tfsdk:"tfid"`
	Data    []NatRules   `tfsdk:"data"`
	Limit   types.Int64  `tfsdk:"limit"`
	Offset  types.Int64  `tfsdk:"offset"`
	Name    types.String `tfsdk:"name"`
	Total   types.Int64  `tfsdk:"total"`
	Folder  types.String `tfsdk:"folder"`
	Snippet types.String `tfsdk:"snippet"`
	Device  types.String `tfsdk:"device"`
}

// NatRulesListDataSourceSchema defines the schema for a list data source.
var NatRulesListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: NatRulesDataSourceSchema.Attributes,
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
