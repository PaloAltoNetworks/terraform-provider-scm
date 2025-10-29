package models

import (
	"regexp"

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

// Package: security_services
// This file contains models for the security_services SDK package

// DosProtectionRules represents the Terraform model for DosProtectionRules
type DosProtectionRules struct {
	Tfid        types.String          `tfsdk:"tfid"`
	Action      basetypes.ObjectValue `tfsdk:"action"`
	Description basetypes.StringValue `tfsdk:"description"`
	Destination basetypes.ListValue   `tfsdk:"destination"`
	Device      basetypes.StringValue `tfsdk:"device"`
	Disabled    basetypes.BoolValue   `tfsdk:"disabled"`
	Folder      basetypes.StringValue `tfsdk:"folder"`
	From        basetypes.ListValue   `tfsdk:"from"`
	Id          basetypes.StringValue `tfsdk:"id"`
	LogSetting  basetypes.StringValue `tfsdk:"log_setting"`
	Name        basetypes.StringValue `tfsdk:"name"`
	Position    basetypes.StringValue `tfsdk:"position"`
	Protection  basetypes.ObjectValue `tfsdk:"protection"`
	Schedule    basetypes.StringValue `tfsdk:"schedule"`
	Service     basetypes.ListValue   `tfsdk:"service"`
	Snippet     basetypes.StringValue `tfsdk:"snippet"`
	Source      basetypes.ListValue   `tfsdk:"source"`
	SourceUser  basetypes.ListValue   `tfsdk:"source_user"`
	Tag         basetypes.ListValue   `tfsdk:"tag"`
	To          basetypes.ListValue   `tfsdk:"to"`
}

// DosProtectionRulesAction represents a nested structure within the DosProtectionRules model
type DosProtectionRulesAction struct {
	Allow   basetypes.ObjectValue `tfsdk:"allow"`
	Deny    basetypes.ObjectValue `tfsdk:"deny"`
	Protect basetypes.ObjectValue `tfsdk:"protect"`
}

// DosProtectionRulesProtection represents a nested structure within the DosProtectionRules model
type DosProtectionRulesProtection struct {
	Aggregate  basetypes.ObjectValue `tfsdk:"aggregate"`
	Classified basetypes.ObjectValue `tfsdk:"classified"`
}

// DosProtectionRulesProtectionAggregate represents a nested structure within the DosProtectionRules model
type DosProtectionRulesProtectionAggregate struct {
	Profile basetypes.StringValue `tfsdk:"profile"`
}

// DosProtectionRulesProtectionClassified represents a nested structure within the DosProtectionRules model
type DosProtectionRulesProtectionClassified struct {
	ClassificationCriteria basetypes.ObjectValue `tfsdk:"classification_criteria"`
	Profile                basetypes.StringValue `tfsdk:"profile"`
}

// DosProtectionRulesProtectionClassifiedClassificationCriteria represents a nested structure within the DosProtectionRules model
type DosProtectionRulesProtectionClassifiedClassificationCriteria struct {
	Address basetypes.StringValue `tfsdk:"address"`
}

// AttrTypes defines the attribute types for the DosProtectionRules model.
func (o DosProtectionRules) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid": basetypes.StringType{},
		"action": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"allow": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"deny": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"protect": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
			},
		},
		"description": basetypes.StringType{},
		"destination": basetypes.ListType{ElemType: basetypes.StringType{}},
		"device":      basetypes.StringType{},
		"disabled":    basetypes.BoolType{},
		"folder":      basetypes.StringType{},
		"from":        basetypes.ListType{ElemType: basetypes.StringType{}},
		"id":          basetypes.StringType{},
		"log_setting": basetypes.StringType{},
		"name":        basetypes.StringType{},
		"position":    basetypes.StringType{},
		"protection": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"aggregate": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"profile": basetypes.StringType{},
					},
				},
				"classified": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"classification_criteria": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"address": basetypes.StringType{},
							},
						},
						"profile": basetypes.StringType{},
					},
				},
			},
		},
		"schedule":    basetypes.StringType{},
		"service":     basetypes.ListType{ElemType: basetypes.StringType{}},
		"snippet":     basetypes.StringType{},
		"source":      basetypes.ListType{ElemType: basetypes.StringType{}},
		"source_user": basetypes.ListType{ElemType: basetypes.StringType{}},
		"tag":         basetypes.ListType{ElemType: basetypes.StringType{}},
		"to":          basetypes.ListType{ElemType: basetypes.StringType{}},
	}
}

// AttrType returns the attribute type for a list of DosProtectionRules objects.
func (o DosProtectionRules) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the DosProtectionRulesAction model.
func (o DosProtectionRulesAction) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"allow": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
		"deny": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
		"protect": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
	}
}

// AttrType returns the attribute type for a list of DosProtectionRulesAction objects.
func (o DosProtectionRulesAction) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the DosProtectionRulesProtection model.
func (o DosProtectionRulesProtection) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"aggregate": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"profile": basetypes.StringType{},
			},
		},
		"classified": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"classification_criteria": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"address": basetypes.StringType{},
					},
				},
				"profile": basetypes.StringType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of DosProtectionRulesProtection objects.
func (o DosProtectionRulesProtection) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the DosProtectionRulesProtectionAggregate model.
func (o DosProtectionRulesProtectionAggregate) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"profile": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of DosProtectionRulesProtectionAggregate objects.
func (o DosProtectionRulesProtectionAggregate) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the DosProtectionRulesProtectionClassified model.
func (o DosProtectionRulesProtectionClassified) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"classification_criteria": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"address": basetypes.StringType{},
			},
		},
		"profile": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of DosProtectionRulesProtectionClassified objects.
func (o DosProtectionRulesProtectionClassified) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the DosProtectionRulesProtectionClassifiedClassificationCriteria model.
func (o DosProtectionRulesProtectionClassifiedClassificationCriteria) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"address": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of DosProtectionRulesProtectionClassifiedClassificationCriteria objects.
func (o DosProtectionRulesProtectionClassifiedClassificationCriteria) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// DosProtectionRulesResourceSchema defines the schema for DosProtectionRules resource
var DosProtectionRulesResourceSchema = schema.Schema{
	MarkdownDescription: "DosProtectionRule resource",
	Attributes: map[string]schema.Attribute{
		"action": schema.SingleNestedAttribute{
			MarkdownDescription: "The action to take on rule match",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"allow": schema.SingleNestedAttribute{
					Validators: []validator.Object{
						objectvalidator.ConflictsWith(
							path.MatchRelative().AtParent().AtName("deny"),
							path.MatchRelative().AtParent().AtName("protect"),
						),
					},
					MarkdownDescription: "Allow",
					Optional:            true,
					Attributes:          map[string]schema.Attribute{},
				},
				"deny": schema.SingleNestedAttribute{
					Validators: []validator.Object{
						objectvalidator.ConflictsWith(
							path.MatchRelative().AtParent().AtName("allow"),
							path.MatchRelative().AtParent().AtName("protect"),
						),
					},
					MarkdownDescription: "Deny",
					Optional:            true,
					Attributes:          map[string]schema.Attribute{},
				},
				"protect": schema.SingleNestedAttribute{
					Validators: []validator.Object{
						objectvalidator.ConflictsWith(
							path.MatchRelative().AtParent().AtName("allow"),
							path.MatchRelative().AtParent().AtName("deny"),
						),
					},
					MarkdownDescription: "Protect",
					Optional:            true,
					Attributes:          map[string]schema.Attribute{},
				},
			},
		},
		"description": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(255),
			},
			MarkdownDescription: "Description",
			Optional:            true,
		},
		"destination": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "List of destination addresses",
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
		"disabled": schema.BoolAttribute{
			MarkdownDescription: "Rule disabled?",
			Optional:            true,
			Computed:            true,
			Default:             booldefault.StaticBool(false),
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
		"from": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "List of source zones",
			Optional:            true,
		},
		"id": schema.StringAttribute{
			MarkdownDescription: "The UUID of the DNS security profile",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"log_setting": schema.StringAttribute{
			MarkdownDescription: "Log forwarding profile name",
			Optional:            true,
			Computed:            true,
			Default:             stringdefault.StaticString("Cortex Data Lake"),
		},
		"name": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(31),
			},
			MarkdownDescription: "Rule name",
			Required:            true,
		},
		"position": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.OneOf("pre", "post"),
			},
			MarkdownDescription: "Position relative to local device rules",
			Optional:            true,
			Computed:            true,
			Default:             stringdefault.StaticString("pre"),
		},
		"protection": schema.SingleNestedAttribute{
			MarkdownDescription: "Protection",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"aggregate": schema.SingleNestedAttribute{
					Validators: []validator.Object{
						objectvalidator.ConflictsWith(
							path.MatchRelative().AtParent().AtName("classified"),
						),
					},
					MarkdownDescription: "Aggregate",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"profile": schema.StringAttribute{
							MarkdownDescription: "Aggregate DoS protection profile",
							Required:            true,
						},
					},
				},
				"classified": schema.SingleNestedAttribute{
					Validators: []validator.Object{
						objectvalidator.ConflictsWith(
							path.MatchRelative().AtParent().AtName("aggregate"),
						),
					},
					MarkdownDescription: "Classified",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"classification_criteria": schema.SingleNestedAttribute{
							MarkdownDescription: "Classification criteria",
							Optional:            true,
							Attributes: map[string]schema.Attribute{
								"address": schema.StringAttribute{
									MarkdownDescription: "Address",
									Optional:            true,
								},
							},
						},
						"profile": schema.StringAttribute{
							MarkdownDescription: "Classified DoS protection profile",
							Required:            true,
						},
					},
				},
			},
		},
		"schedule": schema.StringAttribute{
			MarkdownDescription: "Schedule on which to enforce the rule",
			Optional:            true,
		},
		"service": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "List of services",
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
			MarkdownDescription: "List of source addresses",
			Optional:            true,
		},
		"source_user": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "List of source users and/or groups.  Reserved words include `any`, `pre-login`, `known-user`, and `unknown`.",
			Optional:            true,
		},
		"tag": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "List of tags",
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
			MarkdownDescription: "List of destination zones",
			Optional:            true,
		},
	},
}

// DosProtectionRulesDataSourceSchema defines the schema for DosProtectionRules data source
var DosProtectionRulesDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "DosProtectionRule data source",
	Attributes: map[string]dsschema.Attribute{
		"action": dsschema.SingleNestedAttribute{
			MarkdownDescription: "The action to take on rule match",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"allow": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Allow",
					Computed:            true,
					Attributes:          map[string]dsschema.Attribute{},
				},
				"deny": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Deny",
					Computed:            true,
					Attributes:          map[string]dsschema.Attribute{},
				},
				"protect": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Protect",
					Computed:            true,
					Attributes:          map[string]dsschema.Attribute{},
				},
			},
		},
		"description": dsschema.StringAttribute{
			MarkdownDescription: "Description",
			Computed:            true,
		},
		"destination": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "List of destination addresses",
			Computed:            true,
		},
		"device": dsschema.StringAttribute{
			MarkdownDescription: "The device in which the resource is defined",
			Computed:            true,
		},
		"disabled": dsschema.BoolAttribute{
			MarkdownDescription: "Rule disabled?",
			Computed:            true,
		},
		"folder": dsschema.StringAttribute{
			MarkdownDescription: "The folder in which the resource is defined",
			Computed:            true,
		},
		"from": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "List of source zones",
			Computed:            true,
		},
		"id": dsschema.StringAttribute{
			MarkdownDescription: "The UUID of the DNS security profile",
			Required:            true,
		},
		"log_setting": dsschema.StringAttribute{
			MarkdownDescription: "Log forwarding profile name",
			Computed:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "Rule name",
			Optional:            true,
			Computed:            true,
		},
		"position": dsschema.StringAttribute{
			MarkdownDescription: "Position relative to local device rules",
			Computed:            true,
		},
		"protection": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Protection",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"aggregate": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Aggregate",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"profile": dsschema.StringAttribute{
							MarkdownDescription: "Aggregate DoS protection profile",
							Computed:            true,
						},
					},
				},
				"classified": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Classified",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"classification_criteria": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Classification criteria",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"address": dsschema.StringAttribute{
									MarkdownDescription: "Address",
									Computed:            true,
								},
							},
						},
						"profile": dsschema.StringAttribute{
							MarkdownDescription: "Classified DoS protection profile",
							Computed:            true,
						},
					},
				},
			},
		},
		"schedule": dsschema.StringAttribute{
			MarkdownDescription: "Schedule on which to enforce the rule",
			Computed:            true,
		},
		"service": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "List of services",
			Computed:            true,
		},
		"snippet": dsschema.StringAttribute{
			MarkdownDescription: "The snippet in which the resource is defined",
			Computed:            true,
		},
		"source": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "List of source addresses",
			Computed:            true,
		},
		"source_user": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "List of source users and/or groups.  Reserved words include `any`, `pre-login`, `known-user`, and `unknown`.",
			Computed:            true,
		},
		"tag": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "List of tags",
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
		"to": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "List of destination zones",
			Computed:            true,
		},
	},
}

// DosProtectionRulesListModel represents the data model for a list data source.
type DosProtectionRulesListModel struct {
	Tfid    types.String         `tfsdk:"tfid"`
	Data    []DosProtectionRules `tfsdk:"data"`
	Limit   types.Int64          `tfsdk:"limit"`
	Offset  types.Int64          `tfsdk:"offset"`
	Name    types.String         `tfsdk:"name"`
	Total   types.Int64          `tfsdk:"total"`
	Folder  types.String         `tfsdk:"folder"`
	Snippet types.String         `tfsdk:"snippet"`
	Device  types.String         `tfsdk:"device"`
}

// DosProtectionRulesListDataSourceSchema defines the schema for a list data source.
var DosProtectionRulesListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: DosProtectionRulesDataSourceSchema.Attributes,
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
