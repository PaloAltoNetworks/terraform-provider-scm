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

// Package: objects
// This file contains models for the objects SDK package

// Schedules represents the Terraform model for Schedules
type Schedules struct {
	Tfid         types.String          `tfsdk:"tfid"`
	Device       basetypes.StringValue `tfsdk:"device"`
	Folder       basetypes.StringValue `tfsdk:"folder"`
	Id           basetypes.StringValue `tfsdk:"id"`
	Name         basetypes.StringValue `tfsdk:"name"`
	ScheduleType basetypes.ObjectValue `tfsdk:"schedule_type"`
	Snippet      basetypes.StringValue `tfsdk:"snippet"`
}

// SchedulesScheduleType represents a nested structure within the Schedules model
type SchedulesScheduleType struct {
	NonRecurring basetypes.ListValue   `tfsdk:"non_recurring"`
	Recurring    basetypes.ObjectValue `tfsdk:"recurring"`
}

// SchedulesScheduleTypeRecurring represents a nested structure within the Schedules model
type SchedulesScheduleTypeRecurring struct {
	Daily  basetypes.ListValue   `tfsdk:"daily"`
	Weekly basetypes.ObjectValue `tfsdk:"weekly"`
}

// SchedulesScheduleTypeRecurringWeekly represents a nested structure within the Schedules model
type SchedulesScheduleTypeRecurringWeekly struct {
	Friday    basetypes.ListValue `tfsdk:"friday"`
	Monday    basetypes.ListValue `tfsdk:"monday"`
	Saturday  basetypes.ListValue `tfsdk:"saturday"`
	Sunday    basetypes.ListValue `tfsdk:"sunday"`
	Thursday  basetypes.ListValue `tfsdk:"thursday"`
	Tuesday   basetypes.ListValue `tfsdk:"tuesday"`
	Wednesday basetypes.ListValue `tfsdk:"wednesday"`
}

// AttrTypes defines the attribute types for the Schedules model.
func (o Schedules) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":   basetypes.StringType{},
		"device": basetypes.StringType{},
		"folder": basetypes.StringType{},
		"id":     basetypes.StringType{},
		"name":   basetypes.StringType{},
		"schedule_type": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"non_recurring": basetypes.ListType{ElemType: basetypes.StringType{}},
				"recurring": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"daily": basetypes.ListType{ElemType: basetypes.StringType{}},
						"weekly": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"friday":    basetypes.ListType{ElemType: basetypes.StringType{}},
								"monday":    basetypes.ListType{ElemType: basetypes.StringType{}},
								"saturday":  basetypes.ListType{ElemType: basetypes.StringType{}},
								"sunday":    basetypes.ListType{ElemType: basetypes.StringType{}},
								"thursday":  basetypes.ListType{ElemType: basetypes.StringType{}},
								"tuesday":   basetypes.ListType{ElemType: basetypes.StringType{}},
								"wednesday": basetypes.ListType{ElemType: basetypes.StringType{}},
							},
						},
					},
				},
			},
		},
		"snippet": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of Schedules objects.
func (o Schedules) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the SchedulesScheduleType model.
func (o SchedulesScheduleType) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"non_recurring": basetypes.ListType{ElemType: basetypes.StringType{}},
		"recurring": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"daily": basetypes.ListType{ElemType: basetypes.StringType{}},
				"weekly": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"friday":    basetypes.ListType{ElemType: basetypes.StringType{}},
						"monday":    basetypes.ListType{ElemType: basetypes.StringType{}},
						"saturday":  basetypes.ListType{ElemType: basetypes.StringType{}},
						"sunday":    basetypes.ListType{ElemType: basetypes.StringType{}},
						"thursday":  basetypes.ListType{ElemType: basetypes.StringType{}},
						"tuesday":   basetypes.ListType{ElemType: basetypes.StringType{}},
						"wednesday": basetypes.ListType{ElemType: basetypes.StringType{}},
					},
				},
			},
		},
	}
}

// AttrType returns the attribute type for a list of SchedulesScheduleType objects.
func (o SchedulesScheduleType) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the SchedulesScheduleTypeRecurring model.
func (o SchedulesScheduleTypeRecurring) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"daily": basetypes.ListType{ElemType: basetypes.StringType{}},
		"weekly": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"friday":    basetypes.ListType{ElemType: basetypes.StringType{}},
				"monday":    basetypes.ListType{ElemType: basetypes.StringType{}},
				"saturday":  basetypes.ListType{ElemType: basetypes.StringType{}},
				"sunday":    basetypes.ListType{ElemType: basetypes.StringType{}},
				"thursday":  basetypes.ListType{ElemType: basetypes.StringType{}},
				"tuesday":   basetypes.ListType{ElemType: basetypes.StringType{}},
				"wednesday": basetypes.ListType{ElemType: basetypes.StringType{}},
			},
		},
	}
}

// AttrType returns the attribute type for a list of SchedulesScheduleTypeRecurring objects.
func (o SchedulesScheduleTypeRecurring) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the SchedulesScheduleTypeRecurringWeekly model.
func (o SchedulesScheduleTypeRecurringWeekly) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"friday":    basetypes.ListType{ElemType: basetypes.StringType{}},
		"monday":    basetypes.ListType{ElemType: basetypes.StringType{}},
		"saturday":  basetypes.ListType{ElemType: basetypes.StringType{}},
		"sunday":    basetypes.ListType{ElemType: basetypes.StringType{}},
		"thursday":  basetypes.ListType{ElemType: basetypes.StringType{}},
		"tuesday":   basetypes.ListType{ElemType: basetypes.StringType{}},
		"wednesday": basetypes.ListType{ElemType: basetypes.StringType{}},
	}
}

// AttrType returns the attribute type for a list of SchedulesScheduleTypeRecurringWeekly objects.
func (o SchedulesScheduleTypeRecurringWeekly) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// SchedulesResourceSchema defines the schema for Schedules resource
var SchedulesResourceSchema = schema.Schema{
	MarkdownDescription: "Schedule resource",
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
			MarkdownDescription: "The UUID of the schedule",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"name": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(31),
				stringvalidator.RegexMatches(regexp.MustCompile("^[ a-zA-Z\\d._-]+$"), "pattern must match "+"^[ a-zA-Z\\d._-]+$"),
			},
			MarkdownDescription: "The name of the schedule",
			Required:            true,
		},
		"schedule_type": schema.SingleNestedAttribute{
			MarkdownDescription: "Schedule type",
			Required:            true,
			Attributes: map[string]schema.Attribute{
				"non_recurring": schema.ListAttribute{
					ElementType:         types.StringType,
					MarkdownDescription: "Non recurring\n> ℹ️ **Note:** You must specify exactly one of `non_recurring` and `recurring`.",
					Validators: []validator.List{
						listvalidator.ExactlyOneOf(
							path.MatchRelative().AtParent().AtName("recurring"),
						),
						listvalidator.ValueStringsAre(stringvalidator.LengthAtMost(33)),
					},
					Optional: true,
				},
				"recurring": schema.SingleNestedAttribute{
					Validators: []validator.Object{
						objectvalidator.ExactlyOneOf(
							path.MatchRelative().AtParent().AtName("non_recurring"),
						),
					},
					MarkdownDescription: "Recurring\n> ℹ️ **Note:** You must specify exactly one of `non_recurring` and `recurring`.",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"daily": schema.ListAttribute{
							ElementType:         types.StringType,
							MarkdownDescription: "Daily\n> ℹ️ **Note:** You must specify exactly one of `daily` and `weekly`.",
							Validators: []validator.List{
								listvalidator.ConflictsWith(
									path.MatchRelative().AtParent().AtName("weekly"),
								),
								listvalidator.ValueStringsAre(stringvalidator.LengthAtMost(11)),
							},
							Optional: true,
						},
						"weekly": schema.SingleNestedAttribute{
							Validators: []validator.Object{
								objectvalidator.ConflictsWith(
									path.MatchRelative().AtParent().AtName("daily"),
								),
							},
							MarkdownDescription: "Weekly\n> ℹ️ **Note:** You must specify exactly one of `daily` and `weekly`.",
							Optional:            true,
							Attributes: map[string]schema.Attribute{
								"friday": schema.ListAttribute{
									ElementType:         types.StringType,
									MarkdownDescription: "Friday",
									Validators: []validator.List{
										listvalidator.ValueStringsAre(stringvalidator.LengthAtMost(11)),
									},
									Optional: true,
								},
								"monday": schema.ListAttribute{
									ElementType:         types.StringType,
									MarkdownDescription: "Monday",
									Validators: []validator.List{
										listvalidator.ValueStringsAre(stringvalidator.LengthAtMost(11)),
									},
									Optional: true,
								},
								"saturday": schema.ListAttribute{
									ElementType:         types.StringType,
									MarkdownDescription: "Saturday",
									Validators: []validator.List{
										listvalidator.ValueStringsAre(stringvalidator.LengthAtMost(11)),
									},
									Optional: true,
								},
								"sunday": schema.ListAttribute{
									ElementType:         types.StringType,
									MarkdownDescription: "Sunday",
									Validators: []validator.List{
										listvalidator.ValueStringsAre(stringvalidator.LengthAtMost(11)),
									},
									Optional: true,
								},
								"thursday": schema.ListAttribute{
									ElementType:         types.StringType,
									MarkdownDescription: "Thursday",
									Validators: []validator.List{
										listvalidator.ValueStringsAre(stringvalidator.LengthAtMost(11)),
									},
									Optional: true,
								},
								"tuesday": schema.ListAttribute{
									ElementType:         types.StringType,
									MarkdownDescription: "Tuesday",
									Validators: []validator.List{
										listvalidator.ValueStringsAre(stringvalidator.LengthAtMost(11)),
									},
									Optional: true,
								},
								"wednesday": schema.ListAttribute{
									ElementType:         types.StringType,
									MarkdownDescription: "Wednesday",
									Validators: []validator.List{
										listvalidator.ValueStringsAre(stringvalidator.LengthAtMost(11)),
									},
									Optional: true,
								},
							},
						},
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
	},
}

// SchedulesDataSourceSchema defines the schema for Schedules data source
var SchedulesDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Schedule data source",
	Attributes: map[string]dsschema.Attribute{
		"device": dsschema.StringAttribute{
			MarkdownDescription: "The device in which the resource is defined\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			Computed:            true,
		},
		"folder": dsschema.StringAttribute{
			MarkdownDescription: "The folder in which the resource is defined\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			Computed:            true,
		},
		"id": dsschema.StringAttribute{
			MarkdownDescription: "The UUID of the schedule",
			Required:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "The name of the schedule",
			Optional:            true,
			Computed:            true,
		},
		"schedule_type": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Schedule type",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"non_recurring": dsschema.ListAttribute{
					ElementType:         types.StringType,
					MarkdownDescription: "Non recurring\n> ℹ️ **Note:** You must specify exactly one of `non_recurring` and `recurring`.",
					Computed:            true,
				},
				"recurring": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Recurring\n> ℹ️ **Note:** You must specify exactly one of `non_recurring` and `recurring`.",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"daily": dsschema.ListAttribute{
							ElementType:         types.StringType,
							MarkdownDescription: "Daily\n> ℹ️ **Note:** You must specify exactly one of `daily` and `weekly`.",
							Computed:            true,
						},
						"weekly": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Weekly\n> ℹ️ **Note:** You must specify exactly one of `daily` and `weekly`.",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"friday": dsschema.ListAttribute{
									ElementType:         types.StringType,
									MarkdownDescription: "Friday",
									Computed:            true,
								},
								"monday": dsschema.ListAttribute{
									ElementType:         types.StringType,
									MarkdownDescription: "Monday",
									Computed:            true,
								},
								"saturday": dsschema.ListAttribute{
									ElementType:         types.StringType,
									MarkdownDescription: "Saturday",
									Computed:            true,
								},
								"sunday": dsschema.ListAttribute{
									ElementType:         types.StringType,
									MarkdownDescription: "Sunday",
									Computed:            true,
								},
								"thursday": dsschema.ListAttribute{
									ElementType:         types.StringType,
									MarkdownDescription: "Thursday",
									Computed:            true,
								},
								"tuesday": dsschema.ListAttribute{
									ElementType:         types.StringType,
									MarkdownDescription: "Tuesday",
									Computed:            true,
								},
								"wednesday": dsschema.ListAttribute{
									ElementType:         types.StringType,
									MarkdownDescription: "Wednesday",
									Computed:            true,
								},
							},
						},
					},
				},
			},
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
	},
}

// SchedulesListModel represents the data model for a list data source.
type SchedulesListModel struct {
	Tfid    types.String `tfsdk:"tfid"`
	Data    []Schedules  `tfsdk:"data"`
	Limit   types.Int64  `tfsdk:"limit"`
	Offset  types.Int64  `tfsdk:"offset"`
	Name    types.String `tfsdk:"name"`
	Total   types.Int64  `tfsdk:"total"`
	Folder  types.String `tfsdk:"folder"`
	Snippet types.String `tfsdk:"snippet"`
	Device  types.String `tfsdk:"device"`
}

// SchedulesListDataSourceSchema defines the schema for a list data source.
var SchedulesListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: SchedulesDataSourceSchema.Attributes,
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
