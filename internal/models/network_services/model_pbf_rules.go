package models

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
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

// Package: network_services
// This file contains models for the network_services SDK package

// PbfRules represents the Terraform model for PbfRules
type PbfRules struct {
	Tfid                   types.String          `tfsdk:"tfid"`
	Action                 basetypes.ObjectValue `tfsdk:"action"`
	Application            basetypes.ListValue   `tfsdk:"application"`
	Description            basetypes.StringValue `tfsdk:"description"`
	Destination            basetypes.ListValue   `tfsdk:"destination"`
	Device                 basetypes.StringValue `tfsdk:"device"`
	EnforceSymmetricReturn basetypes.ObjectValue `tfsdk:"enforce_symmetric_return"`
	Folder                 basetypes.StringValue `tfsdk:"folder"`
	From                   basetypes.ObjectValue `tfsdk:"from"`
	Id                     basetypes.StringValue `tfsdk:"id"`
	Name                   basetypes.StringValue `tfsdk:"name"`
	Schedule               basetypes.StringValue `tfsdk:"schedule"`
	Service                basetypes.ListValue   `tfsdk:"service"`
	Snippet                basetypes.StringValue `tfsdk:"snippet"`
	Source                 basetypes.ListValue   `tfsdk:"source"`
	SourceUser             basetypes.ListValue   `tfsdk:"source_user"`
	Tag                    basetypes.ListValue   `tfsdk:"tag"`
}

// PbfRulesAction represents a nested structure within the PbfRules model
type PbfRulesAction struct {
	Discard basetypes.ObjectValue `tfsdk:"discard"`
	Forward basetypes.ObjectValue `tfsdk:"forward"`
	NoPbf   basetypes.ObjectValue `tfsdk:"no_pbf"`
}

// PbfRulesActionForward represents a nested structure within the PbfRules model
type PbfRulesActionForward struct {
	EgressInterface basetypes.StringValue `tfsdk:"egress_interface"`
	Monitor         basetypes.ObjectValue `tfsdk:"monitor"`
	Nexthop         basetypes.ObjectValue `tfsdk:"nexthop"`
}

// PbfRulesActionForwardMonitor represents a nested structure within the PbfRules model
type PbfRulesActionForwardMonitor struct {
	DisableIfUnreachable basetypes.BoolValue   `tfsdk:"disable_if_unreachable"`
	IpAddress            basetypes.StringValue `tfsdk:"ip_address"`
	Profile              basetypes.StringValue `tfsdk:"profile"`
}

// PbfRulesActionForwardNexthop represents a nested structure within the PbfRules model
type PbfRulesActionForwardNexthop struct {
	Fqdn      basetypes.StringValue `tfsdk:"fqdn"`
	IpAddress basetypes.StringValue `tfsdk:"ip_address"`
}

// PbfRulesEnforceSymmetricReturn represents a nested structure within the PbfRules model
type PbfRulesEnforceSymmetricReturn struct {
	Enabled            basetypes.BoolValue `tfsdk:"enabled"`
	NexthopAddressList basetypes.ListValue `tfsdk:"nexthop_address_list"`
}

// PbfRulesEnforceSymmetricReturnNexthopAddressListInner represents a nested structure within the PbfRules model
type PbfRulesEnforceSymmetricReturnNexthopAddressListInner struct {
	Name basetypes.StringValue `tfsdk:"name"`
}

// PbfRulesFrom represents a nested structure within the PbfRules model
type PbfRulesFrom struct {
	Interface basetypes.ListValue `tfsdk:"interface"`
	Zone      basetypes.ListValue `tfsdk:"zone"`
}

// AttrTypes defines the attribute types for the PbfRules model.
func (o PbfRules) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid": basetypes.StringType{},
		"action": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"discard": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"forward": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"egress_interface": basetypes.StringType{},
						"monitor": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"disable_if_unreachable": basetypes.BoolType{},
								"ip_address":             basetypes.StringType{},
								"profile":                basetypes.StringType{},
							},
						},
						"nexthop": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"fqdn":       basetypes.StringType{},
								"ip_address": basetypes.StringType{},
							},
						},
					},
				},
				"no_pbf": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
			},
		},
		"application": basetypes.ListType{ElemType: basetypes.StringType{}},
		"description": basetypes.StringType{},
		"destination": basetypes.ListType{ElemType: basetypes.StringType{}},
		"device":      basetypes.StringType{},
		"enforce_symmetric_return": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"enabled": basetypes.BoolType{},
				"nexthop_address_list": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"name": basetypes.StringType{},
					},
				}},
			},
		},
		"folder": basetypes.StringType{},
		"from": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"interface": basetypes.ListType{ElemType: basetypes.StringType{}},
				"zone":      basetypes.ListType{ElemType: basetypes.StringType{}},
			},
		},
		"id":          basetypes.StringType{},
		"name":        basetypes.StringType{},
		"schedule":    basetypes.StringType{},
		"service":     basetypes.ListType{ElemType: basetypes.StringType{}},
		"snippet":     basetypes.StringType{},
		"source":      basetypes.ListType{ElemType: basetypes.StringType{}},
		"source_user": basetypes.ListType{ElemType: basetypes.StringType{}},
		"tag":         basetypes.ListType{ElemType: basetypes.StringType{}},
	}
}

// AttrType returns the attribute type for a list of PbfRules objects.
func (o PbfRules) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the PbfRulesAction model.
func (o PbfRulesAction) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"discard": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
		"forward": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"egress_interface": basetypes.StringType{},
				"monitor": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"disable_if_unreachable": basetypes.BoolType{},
						"ip_address":             basetypes.StringType{},
						"profile":                basetypes.StringType{},
					},
				},
				"nexthop": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"fqdn":       basetypes.StringType{},
						"ip_address": basetypes.StringType{},
					},
				},
			},
		},
		"no_pbf": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
	}
}

// AttrType returns the attribute type for a list of PbfRulesAction objects.
func (o PbfRulesAction) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the PbfRulesActionForward model.
func (o PbfRulesActionForward) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"egress_interface": basetypes.StringType{},
		"monitor": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"disable_if_unreachable": basetypes.BoolType{},
				"ip_address":             basetypes.StringType{},
				"profile":                basetypes.StringType{},
			},
		},
		"nexthop": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"fqdn":       basetypes.StringType{},
				"ip_address": basetypes.StringType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of PbfRulesActionForward objects.
func (o PbfRulesActionForward) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the PbfRulesActionForwardMonitor model.
func (o PbfRulesActionForwardMonitor) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"disable_if_unreachable": basetypes.BoolType{},
		"ip_address":             basetypes.StringType{},
		"profile":                basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of PbfRulesActionForwardMonitor objects.
func (o PbfRulesActionForwardMonitor) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the PbfRulesActionForwardNexthop model.
func (o PbfRulesActionForwardNexthop) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"fqdn":       basetypes.StringType{},
		"ip_address": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of PbfRulesActionForwardNexthop objects.
func (o PbfRulesActionForwardNexthop) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the PbfRulesEnforceSymmetricReturn model.
func (o PbfRulesEnforceSymmetricReturn) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"enabled": basetypes.BoolType{},
		"nexthop_address_list": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"name": basetypes.StringType{},
			},
		}},
	}
}

// AttrType returns the attribute type for a list of PbfRulesEnforceSymmetricReturn objects.
func (o PbfRulesEnforceSymmetricReturn) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the PbfRulesEnforceSymmetricReturnNexthopAddressListInner model.
func (o PbfRulesEnforceSymmetricReturnNexthopAddressListInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of PbfRulesEnforceSymmetricReturnNexthopAddressListInner objects.
func (o PbfRulesEnforceSymmetricReturnNexthopAddressListInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the PbfRulesFrom model.
func (o PbfRulesFrom) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"interface": basetypes.ListType{ElemType: basetypes.StringType{}},
		"zone":      basetypes.ListType{ElemType: basetypes.StringType{}},
	}
}

// AttrType returns the attribute type for a list of PbfRulesFrom objects.
func (o PbfRulesFrom) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// PbfRulesResourceSchema defines the schema for PbfRules resource
var PbfRulesResourceSchema = schema.Schema{
	MarkdownDescription: "PbfRule resource",
	Attributes: map[string]schema.Attribute{
		"action": schema.SingleNestedAttribute{
			MarkdownDescription: "Action",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"discard": schema.SingleNestedAttribute{
					Validators: []validator.Object{
						objectvalidator.ConflictsWith(
							path.MatchRelative().AtParent().AtName("forward"),
							path.MatchRelative().AtParent().AtName("no_pbf"),
						),
					},
					MarkdownDescription: "Discard",
					Optional:            true,
					Attributes:          map[string]schema.Attribute{},
				},
				"forward": schema.SingleNestedAttribute{
					Validators: []validator.Object{
						objectvalidator.ConflictsWith(
							path.MatchRelative().AtParent().AtName("discard"),
							path.MatchRelative().AtParent().AtName("no_pbf"),
						),
					},
					MarkdownDescription: "Forward",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"egress_interface": schema.StringAttribute{
							MarkdownDescription: "Egress interface",
							Optional:            true,
						},
						"monitor": schema.SingleNestedAttribute{
							MarkdownDescription: "Monitor",
							Optional:            true,
							Attributes: map[string]schema.Attribute{
								"disable_if_unreachable": schema.BoolAttribute{
									MarkdownDescription: "Disable this rule if nexthop/monitor ip is unreachable?",
									Optional:            true,
								},
								"ip_address": schema.StringAttribute{
									MarkdownDescription: "Monitor IP address",
									Optional:            true,
								},
								"profile": schema.StringAttribute{
									MarkdownDescription: "Monitoring profile",
									Optional:            true,
								},
							},
						},
						"nexthop": schema.SingleNestedAttribute{
							MarkdownDescription: "Nexthop",
							Optional:            true,
							Attributes: map[string]schema.Attribute{
								"fqdn": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.ConflictsWith(
											path.MatchRelative().AtParent().AtName("ip_address"),
										),
									},
									MarkdownDescription: "Next hop FQDN",
									Optional:            true,
								},
								"ip_address": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.ConflictsWith(
											path.MatchRelative().AtParent().AtName("fqdn"),
										),
									},
									MarkdownDescription: "Next hop IP address",
									Optional:            true,
								},
							},
						},
					},
				},
				"no_pbf": schema.SingleNestedAttribute{
					Validators: []validator.Object{
						objectvalidator.ConflictsWith(
							path.MatchRelative().AtParent().AtName("discard"),
							path.MatchRelative().AtParent().AtName("forward"),
						),
					},
					MarkdownDescription: "No pbf",
					Optional:            true,
					Attributes:          map[string]schema.Attribute{},
				},
			},
		},
		"application": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Applications",
			Optional:            true,
		},
		"description": schema.StringAttribute{
			MarkdownDescription: "Description",
			Optional:            true,
		},
		"destination": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Destination addresses",
			Optional:            true,
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
			MarkdownDescription: "The device in which the resource is defined",
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"enforce_symmetric_return": schema.SingleNestedAttribute{
			MarkdownDescription: "Enforce symmetric return",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"enabled": schema.BoolAttribute{
					MarkdownDescription: "Enforce symmetric return?",
					Optional:            true,
				},
				"nexthop_address_list": schema.ListNestedAttribute{
					MarkdownDescription: "Next hop IP addresses",
					Optional:            true,
					NestedObject: schema.NestedAttributeObject{
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								MarkdownDescription: "Next hop IP address",
								Optional:            true,
							},
						},
					},
				},
			},
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
			MarkdownDescription: "The folder in which the resource is defined",
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"from": schema.SingleNestedAttribute{
			MarkdownDescription: "From",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"interface": schema.ListAttribute{
					ElementType:         types.StringType,
					MarkdownDescription: "Source interfaces",
					Validators: []validator.List{
						listvalidator.ConflictsWith(
							path.MatchRelative().AtParent().AtName("zone"),
						),
					},
					Optional: true,
				},
				"zone": schema.ListAttribute{
					ElementType:         types.StringType,
					MarkdownDescription: "Source zones",
					Validators: []validator.List{
						listvalidator.ConflictsWith(
							path.MatchRelative().AtParent().AtName("interface"),
						),
					},
					Optional: true,
				},
			},
		},
		"id": schema.StringAttribute{
			MarkdownDescription: "UUID of the resource",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"name": schema.StringAttribute{
			MarkdownDescription: "PBF rule name",
			Optional:            true,
		},
		"schedule": schema.StringAttribute{
			MarkdownDescription: "Schedule",
			Optional:            true,
		},
		"service": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Services",
			Optional:            true,
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
			MarkdownDescription: "The snippet in which the resource is defined",
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"source": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Source addresses",
			Optional:            true,
		},
		"source_user": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Source users",
			Optional:            true,
		},
		"tag": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Tags",
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

// PbfRulesDataSourceSchema defines the schema for PbfRules data source
var PbfRulesDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "PbfRule data source",
	Attributes: map[string]dsschema.Attribute{
		"action": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Action",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"discard": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Discard",
					Computed:            true,
					Attributes:          map[string]dsschema.Attribute{},
				},
				"forward": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Forward",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"egress_interface": dsschema.StringAttribute{
							MarkdownDescription: "Egress interface",
							Computed:            true,
						},
						"monitor": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Monitor",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"disable_if_unreachable": dsschema.BoolAttribute{
									MarkdownDescription: "Disable this rule if nexthop/monitor ip is unreachable?",
									Computed:            true,
								},
								"ip_address": dsschema.StringAttribute{
									MarkdownDescription: "Monitor IP address",
									Computed:            true,
								},
								"profile": dsschema.StringAttribute{
									MarkdownDescription: "Monitoring profile",
									Computed:            true,
								},
							},
						},
						"nexthop": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Nexthop",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"fqdn": dsschema.StringAttribute{
									MarkdownDescription: "Next hop FQDN",
									Computed:            true,
								},
								"ip_address": dsschema.StringAttribute{
									MarkdownDescription: "Next hop IP address",
									Computed:            true,
								},
							},
						},
					},
				},
				"no_pbf": dsschema.SingleNestedAttribute{
					MarkdownDescription: "No pbf",
					Computed:            true,
					Attributes:          map[string]dsschema.Attribute{},
				},
			},
		},
		"application": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Applications",
			Computed:            true,
		},
		"description": dsschema.StringAttribute{
			MarkdownDescription: "Description",
			Computed:            true,
		},
		"destination": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Destination addresses",
			Computed:            true,
		},
		"device": dsschema.StringAttribute{
			MarkdownDescription: "The device in which the resource is defined",
			Computed:            true,
		},
		"enforce_symmetric_return": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Enforce symmetric return",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"enabled": dsschema.BoolAttribute{
					MarkdownDescription: "Enforce symmetric return?",
					Computed:            true,
				},
				"nexthop_address_list": dsschema.ListNestedAttribute{
					MarkdownDescription: "Next hop IP addresses",
					Computed:            true,
					NestedObject: dsschema.NestedAttributeObject{
						Attributes: map[string]dsschema.Attribute{
							"name": dsschema.StringAttribute{
								MarkdownDescription: "Next hop IP address",
								Computed:            true,
							},
						},
					},
				},
			},
		},
		"folder": dsschema.StringAttribute{
			MarkdownDescription: "The folder in which the resource is defined",
			Computed:            true,
		},
		"from": dsschema.SingleNestedAttribute{
			MarkdownDescription: "From",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"interface": dsschema.ListAttribute{
					ElementType:         types.StringType,
					MarkdownDescription: "Source interfaces",
					Computed:            true,
				},
				"zone": dsschema.ListAttribute{
					ElementType:         types.StringType,
					MarkdownDescription: "Source zones",
					Computed:            true,
				},
			},
		},
		"id": dsschema.StringAttribute{
			MarkdownDescription: "UUID of the resource",
			Required:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "PBF rule name",
			Optional:            true,
			Computed:            true,
		},
		"schedule": dsschema.StringAttribute{
			MarkdownDescription: "Schedule",
			Computed:            true,
		},
		"service": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Services",
			Computed:            true,
		},
		"snippet": dsschema.StringAttribute{
			MarkdownDescription: "The snippet in which the resource is defined",
			Computed:            true,
		},
		"source": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Source addresses",
			Computed:            true,
		},
		"source_user": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Source users",
			Computed:            true,
		},
		"tag": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Tags",
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
	},
}

// PbfRulesListModel represents the data model for a list data source.
type PbfRulesListModel struct {
	Tfid    types.String `tfsdk:"tfid"`
	Data    []PbfRules   `tfsdk:"data"`
	Limit   types.Int64  `tfsdk:"limit"`
	Offset  types.Int64  `tfsdk:"offset"`
	Name    types.String `tfsdk:"name"`
	Total   types.Int64  `tfsdk:"total"`
	Folder  types.String `tfsdk:"folder"`
	Snippet types.String `tfsdk:"snippet"`
	Device  types.String `tfsdk:"device"`
}

// PbfRulesListDataSourceSchema defines the schema for a list data source.
var PbfRulesListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: PbfRulesDataSourceSchema.Attributes,
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
