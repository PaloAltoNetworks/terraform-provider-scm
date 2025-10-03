package models

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/objectvalidator"
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

// Package: security_services
// This file contains models for the security_services SDK package

// DnsSecurityProfiles represents the Terraform model for DnsSecurityProfiles
type DnsSecurityProfiles struct {
	Tfid          types.String          `tfsdk:"tfid"`
	BotnetDomains basetypes.ObjectValue `tfsdk:"botnet_domains"`
	Description   basetypes.StringValue `tfsdk:"description"`
	Device        basetypes.StringValue `tfsdk:"device"`
	Folder        basetypes.StringValue `tfsdk:"folder"`
	Id            basetypes.StringValue `tfsdk:"id"`
	Name          basetypes.StringValue `tfsdk:"name"`
	Snippet       basetypes.StringValue `tfsdk:"snippet"`
}

// DnsSecurityProfilesBotnetDomains represents a nested structure within the DnsSecurityProfiles model
type DnsSecurityProfilesBotnetDomains struct {
	DnsSecurityCategories basetypes.ListValue   `tfsdk:"dns_security_categories"`
	Lists                 basetypes.ListValue   `tfsdk:"lists"`
	Sinkhole              basetypes.ObjectValue `tfsdk:"sinkhole"`
	Whitelist             basetypes.ListValue   `tfsdk:"whitelist"`
}

// DnsSecurityProfilesBotnetDomainsDnsSecurityCategoriesInner represents a nested structure within the DnsSecurityProfiles model
type DnsSecurityProfilesBotnetDomainsDnsSecurityCategoriesInner struct {
	Action        basetypes.StringValue `tfsdk:"action"`
	LogLevel      basetypes.StringValue `tfsdk:"log_level"`
	Name          basetypes.StringValue `tfsdk:"name"`
	PacketCapture basetypes.StringValue `tfsdk:"packet_capture"`
}

// DnsSecurityProfilesBotnetDomainsListsInner represents a nested structure within the DnsSecurityProfiles model
type DnsSecurityProfilesBotnetDomainsListsInner struct {
	Action        basetypes.ObjectValue `tfsdk:"action"`
	Name          basetypes.StringValue `tfsdk:"name"`
	PacketCapture basetypes.StringValue `tfsdk:"packet_capture"`
}

// DnsSecurityProfilesBotnetDomainsListsInnerAction represents a nested structure within the DnsSecurityProfiles model
type DnsSecurityProfilesBotnetDomainsListsInnerAction struct {
	Alert    basetypes.ObjectValue `tfsdk:"alert"`
	Allow    basetypes.ObjectValue `tfsdk:"allow"`
	Block    basetypes.ObjectValue `tfsdk:"block"`
	Sinkhole basetypes.ObjectValue `tfsdk:"sinkhole"`
}

// DnsSecurityProfilesBotnetDomainsSinkhole represents a nested structure within the DnsSecurityProfiles model
type DnsSecurityProfilesBotnetDomainsSinkhole struct {
	Ipv4Address basetypes.StringValue `tfsdk:"ipv4_address"`
	Ipv6Address basetypes.StringValue `tfsdk:"ipv6_address"`
}

// DnsSecurityProfilesBotnetDomainsWhitelistInner represents a nested structure within the DnsSecurityProfiles model
type DnsSecurityProfilesBotnetDomainsWhitelistInner struct {
	Description basetypes.StringValue `tfsdk:"description"`
	Name        basetypes.StringValue `tfsdk:"name"`
}

// AttrTypes defines the attribute types for the DnsSecurityProfiles model.
func (o DnsSecurityProfiles) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid": basetypes.StringType{},
		"botnet_domains": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"dns_security_categories": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"action":         basetypes.StringType{},
						"log_level":      basetypes.StringType{},
						"name":           basetypes.StringType{},
						"packet_capture": basetypes.StringType{},
					},
				}},
				"lists": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"action": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"alert": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"allow": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"block": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"sinkhole": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
							},
						},
						"name":           basetypes.StringType{},
						"packet_capture": basetypes.StringType{},
					},
				}},
				"sinkhole": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"ipv4_address": basetypes.StringType{},
						"ipv6_address": basetypes.StringType{},
					},
				},
				"whitelist": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"description": basetypes.StringType{},
						"name":        basetypes.StringType{},
					},
				}},
			},
		},
		"description": basetypes.StringType{},
		"device":      basetypes.StringType{},
		"folder":      basetypes.StringType{},
		"id":          basetypes.StringType{},
		"name":        basetypes.StringType{},
		"snippet":     basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of DnsSecurityProfiles objects.
func (o DnsSecurityProfiles) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the DnsSecurityProfilesBotnetDomains model.
func (o DnsSecurityProfilesBotnetDomains) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"dns_security_categories": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"action":         basetypes.StringType{},
				"log_level":      basetypes.StringType{},
				"name":           basetypes.StringType{},
				"packet_capture": basetypes.StringType{},
			},
		}},
		"lists": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"action": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"alert": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"allow": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"block": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"sinkhole": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
					},
				},
				"name":           basetypes.StringType{},
				"packet_capture": basetypes.StringType{},
			},
		}},
		"sinkhole": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"ipv4_address": basetypes.StringType{},
				"ipv6_address": basetypes.StringType{},
			},
		},
		"whitelist": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"description": basetypes.StringType{},
				"name":        basetypes.StringType{},
			},
		}},
	}
}

// AttrType returns the attribute type for a list of DnsSecurityProfilesBotnetDomains objects.
func (o DnsSecurityProfilesBotnetDomains) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the DnsSecurityProfilesBotnetDomainsDnsSecurityCategoriesInner model.
func (o DnsSecurityProfilesBotnetDomainsDnsSecurityCategoriesInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"action":         basetypes.StringType{},
		"log_level":      basetypes.StringType{},
		"name":           basetypes.StringType{},
		"packet_capture": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of DnsSecurityProfilesBotnetDomainsDnsSecurityCategoriesInner objects.
func (o DnsSecurityProfilesBotnetDomainsDnsSecurityCategoriesInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the DnsSecurityProfilesBotnetDomainsListsInner model.
func (o DnsSecurityProfilesBotnetDomainsListsInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"action": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"alert": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"allow": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"block": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"sinkhole": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
			},
		},
		"name":           basetypes.StringType{},
		"packet_capture": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of DnsSecurityProfilesBotnetDomainsListsInner objects.
func (o DnsSecurityProfilesBotnetDomainsListsInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the DnsSecurityProfilesBotnetDomainsListsInnerAction model.
func (o DnsSecurityProfilesBotnetDomainsListsInnerAction) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"alert": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
		"allow": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
		"block": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
		"sinkhole": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
	}
}

// AttrType returns the attribute type for a list of DnsSecurityProfilesBotnetDomainsListsInnerAction objects.
func (o DnsSecurityProfilesBotnetDomainsListsInnerAction) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the DnsSecurityProfilesBotnetDomainsSinkhole model.
func (o DnsSecurityProfilesBotnetDomainsSinkhole) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"ipv4_address": basetypes.StringType{},
		"ipv6_address": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of DnsSecurityProfilesBotnetDomainsSinkhole objects.
func (o DnsSecurityProfilesBotnetDomainsSinkhole) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the DnsSecurityProfilesBotnetDomainsWhitelistInner model.
func (o DnsSecurityProfilesBotnetDomainsWhitelistInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"description": basetypes.StringType{},
		"name":        basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of DnsSecurityProfilesBotnetDomainsWhitelistInner objects.
func (o DnsSecurityProfilesBotnetDomainsWhitelistInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// DnsSecurityProfilesResourceSchema defines the schema for DnsSecurityProfiles resource
var DnsSecurityProfilesResourceSchema = schema.Schema{
	MarkdownDescription: "DnsSecurityProfile resource",
	Attributes: map[string]schema.Attribute{
		"botnet_domains": schema.SingleNestedAttribute{
			MarkdownDescription: "Botnet domains",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"dns_security_categories": schema.ListNestedAttribute{
					MarkdownDescription: "DNS categories",
					Optional:            true,
					NestedObject: schema.NestedAttributeObject{
						Attributes: map[string]schema.Attribute{
							"action": schema.StringAttribute{
								Validators: []validator.String{
									stringvalidator.OneOf("default", "allow", "block", "sinkhole"),
								},
								MarkdownDescription: "Action",
								Optional:            true,
								Computed:            true,
								Default:             stringdefault.StaticString("default"),
							},
							"log_level": schema.StringAttribute{
								Validators: []validator.String{
									stringvalidator.OneOf("default", "none", "low", "informational", "medium", "high", "critical"),
								},
								MarkdownDescription: "Log level",
								Optional:            true,
								Computed:            true,
								Default:             stringdefault.StaticString("default"),
							},
							"name": schema.StringAttribute{
								MarkdownDescription: "Name",
								Optional:            true,
							},
							"packet_capture": schema.StringAttribute{
								Validators: []validator.String{
									stringvalidator.OneOf("disable", "single-packet", "extended-capture"),
								},
								MarkdownDescription: "Packet capture",
								Optional:            true,
							},
						},
					},
				},
				"lists": schema.ListNestedAttribute{
					MarkdownDescription: "Dynamic lists of DNS domains",
					Optional:            true,
					NestedObject: schema.NestedAttributeObject{
						Attributes: map[string]schema.Attribute{
							"action": schema.SingleNestedAttribute{
								MarkdownDescription: "Action",
								Optional:            true,
								Attributes: map[string]schema.Attribute{
									"alert": schema.SingleNestedAttribute{
										Validators: []validator.Object{
											objectvalidator.ExactlyOneOf(
												path.MatchRelative().AtParent().AtName("allow"),
												path.MatchRelative().AtParent().AtName("block"),
												path.MatchRelative().AtParent().AtName("sinkhole"),
											),
										},
										MarkdownDescription: "Alert",
										Optional:            true,
										Attributes:          map[string]schema.Attribute{},
									},
									"allow": schema.SingleNestedAttribute{
										Validators: []validator.Object{
											objectvalidator.ExactlyOneOf(
												path.MatchRelative().AtParent().AtName("alert"),
												path.MatchRelative().AtParent().AtName("block"),
												path.MatchRelative().AtParent().AtName("sinkhole"),
											),
										},
										MarkdownDescription: "Allow",
										Optional:            true,
										Attributes:          map[string]schema.Attribute{},
									},
									"block": schema.SingleNestedAttribute{
										Validators: []validator.Object{
											objectvalidator.ExactlyOneOf(
												path.MatchRelative().AtParent().AtName("alert"),
												path.MatchRelative().AtParent().AtName("allow"),
												path.MatchRelative().AtParent().AtName("sinkhole"),
											),
										},
										MarkdownDescription: "Block",
										Optional:            true,
										Attributes:          map[string]schema.Attribute{},
									},
									"sinkhole": schema.SingleNestedAttribute{
										Validators: []validator.Object{
											objectvalidator.ExactlyOneOf(
												path.MatchRelative().AtParent().AtName("alert"),
												path.MatchRelative().AtParent().AtName("allow"),
												path.MatchRelative().AtParent().AtName("block"),
											),
										},
										MarkdownDescription: "Sinkhole",
										Optional:            true,
										Attributes:          map[string]schema.Attribute{},
									},
								},
							},
							"name": schema.StringAttribute{
								MarkdownDescription: "Name",
								Required:            true,
							},
							"packet_capture": schema.StringAttribute{
								Validators: []validator.String{
									stringvalidator.OneOf("disable", "single-packet", "extended-capture"),
								},
								MarkdownDescription: "Packet capture",
								Optional:            true,
							},
						},
					},
				},
				"sinkhole": schema.SingleNestedAttribute{
					MarkdownDescription: "DNS sinkhole settings",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"ipv4_address": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("127.0.0.1", "pan-sinkhole-default-ip"),
							},
							MarkdownDescription: "Ipv4 address",
							Optional:            true,
						},
						"ipv6_address": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("::1"),
							},
							MarkdownDescription: "Ipv6 address",
							Optional:            true,
						},
					},
				},
				"whitelist": schema.ListNestedAttribute{
					MarkdownDescription: "DNS security overrides",
					Optional:            true,
					NestedObject: schema.NestedAttributeObject{
						Attributes: map[string]schema.Attribute{
							"description": schema.StringAttribute{
								MarkdownDescription: "Description",
								Optional:            true,
							},
							"name": schema.StringAttribute{
								MarkdownDescription: "DNS domain or FQDN to be whitelisted",
								Required:            true,
							},
						},
					},
				},
			},
		},
		"description": schema.StringAttribute{
			MarkdownDescription: "The description of the DNS security profile",
			Optional:            true,
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
			MarkdownDescription: "The UUID of the DNS security profile",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"name": schema.StringAttribute{
			MarkdownDescription: "The name of the DNS security profile",
			Optional:            true,
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
	},
}

// DnsSecurityProfilesDataSourceSchema defines the schema for DnsSecurityProfiles data source
var DnsSecurityProfilesDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "DnsSecurityProfile data source",
	Attributes: map[string]dsschema.Attribute{
		"botnet_domains": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Botnet domains",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"dns_security_categories": dsschema.ListNestedAttribute{
					MarkdownDescription: "DNS categories",
					Computed:            true,
					NestedObject: dsschema.NestedAttributeObject{
						Attributes: map[string]dsschema.Attribute{
							"action": dsschema.StringAttribute{
								MarkdownDescription: "Action",
								Computed:            true,
							},
							"log_level": dsschema.StringAttribute{
								MarkdownDescription: "Log level",
								Computed:            true,
							},
							"name": dsschema.StringAttribute{
								MarkdownDescription: "Name",
								Computed:            true,
							},
							"packet_capture": dsschema.StringAttribute{
								MarkdownDescription: "Packet capture",
								Computed:            true,
							},
						},
					},
				},
				"lists": dsschema.ListNestedAttribute{
					MarkdownDescription: "Dynamic lists of DNS domains",
					Computed:            true,
					NestedObject: dsschema.NestedAttributeObject{
						Attributes: map[string]dsschema.Attribute{
							"action": dsschema.SingleNestedAttribute{
								MarkdownDescription: "Action",
								Computed:            true,
								Attributes: map[string]dsschema.Attribute{
									"alert": dsschema.SingleNestedAttribute{
										MarkdownDescription: "Alert",
										Computed:            true,
										Attributes:          map[string]dsschema.Attribute{},
									},
									"allow": dsschema.SingleNestedAttribute{
										MarkdownDescription: "Allow",
										Computed:            true,
										Attributes:          map[string]dsschema.Attribute{},
									},
									"block": dsschema.SingleNestedAttribute{
										MarkdownDescription: "Block",
										Computed:            true,
										Attributes:          map[string]dsschema.Attribute{},
									},
									"sinkhole": dsschema.SingleNestedAttribute{
										MarkdownDescription: "Sinkhole",
										Computed:            true,
										Attributes:          map[string]dsschema.Attribute{},
									},
								},
							},
							"name": dsschema.StringAttribute{
								MarkdownDescription: "Name",
								Computed:            true,
							},
							"packet_capture": dsschema.StringAttribute{
								MarkdownDescription: "Packet capture",
								Computed:            true,
							},
						},
					},
				},
				"sinkhole": dsschema.SingleNestedAttribute{
					MarkdownDescription: "DNS sinkhole settings",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"ipv4_address": dsschema.StringAttribute{
							MarkdownDescription: "Ipv4 address",
							Computed:            true,
						},
						"ipv6_address": dsschema.StringAttribute{
							MarkdownDescription: "Ipv6 address",
							Computed:            true,
						},
					},
				},
				"whitelist": dsschema.ListNestedAttribute{
					MarkdownDescription: "DNS security overrides",
					Computed:            true,
					NestedObject: dsschema.NestedAttributeObject{
						Attributes: map[string]dsschema.Attribute{
							"description": dsschema.StringAttribute{
								MarkdownDescription: "Description",
								Computed:            true,
							},
							"name": dsschema.StringAttribute{
								MarkdownDescription: "DNS domain or FQDN to be whitelisted",
								Computed:            true,
							},
						},
					},
				},
			},
		},
		"description": dsschema.StringAttribute{
			MarkdownDescription: "The description of the DNS security profile",
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
			MarkdownDescription: "The UUID of the DNS security profile",
			Required:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "The name of the DNS security profile",
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

// DnsSecurityProfilesListModel represents the data model for a list data source.
type DnsSecurityProfilesListModel struct {
	Tfid    types.String          `tfsdk:"tfid"`
	Data    []DnsSecurityProfiles `tfsdk:"data"`
	Limit   types.Int64           `tfsdk:"limit"`
	Offset  types.Int64           `tfsdk:"offset"`
	Name    types.String          `tfsdk:"name"`
	Total   types.Int64           `tfsdk:"total"`
	Folder  types.String          `tfsdk:"folder"`
	Snippet types.String          `tfsdk:"snippet"`
	Device  types.String          `tfsdk:"device"`
}

// DnsSecurityProfilesListDataSourceSchema defines the schema for a list data source.
var DnsSecurityProfilesListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: DnsSecurityProfilesDataSourceSchema.Attributes,
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
