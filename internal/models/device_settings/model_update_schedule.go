package models

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/objectvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	dsschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Package: device_settings
// This file contains models for the device_settings SDK package

// UpdateSchedule represents the Terraform model for UpdateSchedule
type UpdateSchedule struct {
	Tfid           types.String          `tfsdk:"tfid"`
	Device         basetypes.StringValue `tfsdk:"device"`
	Folder         basetypes.StringValue `tfsdk:"folder"`
	Id             basetypes.StringValue `tfsdk:"id"`
	Snippet        basetypes.StringValue `tfsdk:"snippet"`
	UpdateSchedule basetypes.ObjectValue `tfsdk:"update_schedule"`
}

// UpdateScheduleUpdateSchedule represents a nested structure within the UpdateSchedule model
type UpdateScheduleUpdateSchedule struct {
	AntiVirus basetypes.ObjectValue `tfsdk:"anti_virus"`
	Threats   basetypes.ObjectValue `tfsdk:"threats"`
	Wildfire  basetypes.ObjectValue `tfsdk:"wildfire"`
}

// UpdateScheduleUpdateScheduleAntiVirus represents a nested structure within the UpdateSchedule model
type UpdateScheduleUpdateScheduleAntiVirus struct {
	Recurring basetypes.ObjectValue `tfsdk:"recurring"`
}

// UpdateScheduleUpdateScheduleAntiVirusRecurring represents a nested structure within the UpdateSchedule model
type UpdateScheduleUpdateScheduleAntiVirusRecurring struct {
	Daily      basetypes.ObjectValue `tfsdk:"daily"`
	Hourly     basetypes.ObjectValue `tfsdk:"hourly"`
	None       basetypes.ObjectValue `tfsdk:"none"`
	SyncToPeer basetypes.BoolValue   `tfsdk:"sync_to_peer"`
	Threshold  basetypes.Int64Value  `tfsdk:"threshold"`
	Weekly     basetypes.ObjectValue `tfsdk:"weekly"`
}

// UpdateScheduleUpdateScheduleAntiVirusRecurringDaily represents a nested structure within the UpdateSchedule model
type UpdateScheduleUpdateScheduleAntiVirusRecurringDaily struct {
	Action basetypes.StringValue `tfsdk:"action"`
	At     basetypes.StringValue `tfsdk:"at"`
}

// UpdateScheduleUpdateScheduleAntiVirusRecurringHourly represents a nested structure within the UpdateSchedule model
type UpdateScheduleUpdateScheduleAntiVirusRecurringHourly struct {
	Action basetypes.StringValue `tfsdk:"action"`
	At     basetypes.Int64Value  `tfsdk:"at"`
}

// UpdateScheduleUpdateScheduleAntiVirusRecurringWeekly represents a nested structure within the UpdateSchedule model
type UpdateScheduleUpdateScheduleAntiVirusRecurringWeekly struct {
	Action    basetypes.StringValue `tfsdk:"action"`
	At        basetypes.StringValue `tfsdk:"at"`
	DayOfWeek basetypes.StringValue `tfsdk:"day_of_week"`
}

// UpdateScheduleUpdateScheduleThreats represents a nested structure within the UpdateSchedule model
type UpdateScheduleUpdateScheduleThreats struct {
	Recurring basetypes.ObjectValue `tfsdk:"recurring"`
}

// UpdateScheduleUpdateScheduleThreatsRecurring represents a nested structure within the UpdateSchedule model
type UpdateScheduleUpdateScheduleThreatsRecurring struct {
	Daily           basetypes.ObjectValue `tfsdk:"daily"`
	Every30Mins     basetypes.ObjectValue `tfsdk:"every_30_mins"`
	Hourly          basetypes.ObjectValue `tfsdk:"hourly"`
	NewAppThreshold basetypes.Int64Value  `tfsdk:"new_app_threshold"`
	None            basetypes.ObjectValue `tfsdk:"none"`
	SyncToPeer      basetypes.BoolValue   `tfsdk:"sync_to_peer"`
	Threshold       basetypes.Int64Value  `tfsdk:"threshold"`
	Weekly          basetypes.ObjectValue `tfsdk:"weekly"`
}

// UpdateScheduleUpdateScheduleThreatsRecurringDaily represents a nested structure within the UpdateSchedule model
type UpdateScheduleUpdateScheduleThreatsRecurringDaily struct {
	Action            basetypes.StringValue `tfsdk:"action"`
	At                basetypes.StringValue `tfsdk:"at"`
	DisableNewContent basetypes.BoolValue   `tfsdk:"disable_new_content"`
}

// UpdateScheduleUpdateScheduleThreatsRecurringEvery30Mins represents a nested structure within the UpdateSchedule model
type UpdateScheduleUpdateScheduleThreatsRecurringEvery30Mins struct {
	Action            basetypes.StringValue `tfsdk:"action"`
	At                basetypes.Int64Value  `tfsdk:"at"`
	DisableNewContent basetypes.BoolValue   `tfsdk:"disable_new_content"`
}

// UpdateScheduleUpdateScheduleThreatsRecurringHourly represents a nested structure within the UpdateSchedule model
type UpdateScheduleUpdateScheduleThreatsRecurringHourly struct {
	Action            basetypes.StringValue  `tfsdk:"action"`
	At                basetypes.Float64Value `tfsdk:"at"`
	DisableNewContent basetypes.BoolValue    `tfsdk:"disable_new_content"`
}

// UpdateScheduleUpdateScheduleThreatsRecurringWeekly represents a nested structure within the UpdateSchedule model
type UpdateScheduleUpdateScheduleThreatsRecurringWeekly struct {
	Action            basetypes.StringValue `tfsdk:"action"`
	At                basetypes.StringValue `tfsdk:"at"`
	DayOfWeek         basetypes.StringValue `tfsdk:"day_of_week"`
	DisableNewContent basetypes.BoolValue   `tfsdk:"disable_new_content"`
}

// UpdateScheduleUpdateScheduleWildfire represents a nested structure within the UpdateSchedule model
type UpdateScheduleUpdateScheduleWildfire struct {
	Recurring basetypes.ObjectValue `tfsdk:"recurring"`
}

// UpdateScheduleUpdateScheduleWildfireRecurring represents a nested structure within the UpdateSchedule model
type UpdateScheduleUpdateScheduleWildfireRecurring struct {
	Every15Mins basetypes.ObjectValue `tfsdk:"every_15_mins"`
	Every30Mins basetypes.ObjectValue `tfsdk:"every_30_mins"`
	EveryHour   basetypes.ObjectValue `tfsdk:"every_hour"`
	EveryMin    basetypes.ObjectValue `tfsdk:"every_min"`
	None        basetypes.ObjectValue `tfsdk:"none"`
	RealTime    basetypes.ObjectValue `tfsdk:"real_time"`
}

// UpdateScheduleUpdateScheduleWildfireRecurringEvery15Mins represents a nested structure within the UpdateSchedule model
type UpdateScheduleUpdateScheduleWildfireRecurringEvery15Mins struct {
	Action     basetypes.StringValue `tfsdk:"action"`
	At         basetypes.Int64Value  `tfsdk:"at"`
	SyncToPeer basetypes.BoolValue   `tfsdk:"sync_to_peer"`
}

// UpdateScheduleUpdateScheduleWildfireRecurringEvery30Mins represents a nested structure within the UpdateSchedule model
type UpdateScheduleUpdateScheduleWildfireRecurringEvery30Mins struct {
	Action     basetypes.StringValue `tfsdk:"action"`
	At         basetypes.Int64Value  `tfsdk:"at"`
	SyncToPeer basetypes.BoolValue   `tfsdk:"sync_to_peer"`
}

// UpdateScheduleUpdateScheduleWildfireRecurringEveryHour represents a nested structure within the UpdateSchedule model
type UpdateScheduleUpdateScheduleWildfireRecurringEveryHour struct {
	Action     basetypes.StringValue `tfsdk:"action"`
	At         basetypes.Int64Value  `tfsdk:"at"`
	SyncToPeer basetypes.BoolValue   `tfsdk:"sync_to_peer"`
}

// UpdateScheduleUpdateScheduleWildfireRecurringEveryMin represents a nested structure within the UpdateSchedule model
type UpdateScheduleUpdateScheduleWildfireRecurringEveryMin struct {
	Action     basetypes.StringValue `tfsdk:"action"`
	SyncToPeer basetypes.BoolValue   `tfsdk:"sync_to_peer"`
}

// AttrTypes defines the attribute types for the UpdateSchedule model.
func (o UpdateSchedule) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":    basetypes.StringType{},
		"device":  basetypes.StringType{},
		"folder":  basetypes.StringType{},
		"id":      basetypes.StringType{},
		"snippet": basetypes.StringType{},
		"update_schedule": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"anti_virus": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"recurring": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"daily": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"action": basetypes.StringType{},
										"at":     basetypes.StringType{},
									},
								},
								"hourly": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"action": basetypes.StringType{},
										"at":     basetypes.Int64Type{},
									},
								},
								"none": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"sync_to_peer": basetypes.BoolType{},
								"threshold":    basetypes.Int64Type{},
								"weekly": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"action":      basetypes.StringType{},
										"at":          basetypes.StringType{},
										"day_of_week": basetypes.StringType{},
									},
								},
							},
						},
					},
				},
				"threats": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"recurring": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"daily": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"action":              basetypes.StringType{},
										"at":                  basetypes.StringType{},
										"disable_new_content": basetypes.BoolType{},
									},
								},
								"every_30_mins": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"action":              basetypes.StringType{},
										"at":                  basetypes.Int64Type{},
										"disable_new_content": basetypes.BoolType{},
									},
								},
								"hourly": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"action":              basetypes.StringType{},
										"at":                  basetypes.Float64Type{},
										"disable_new_content": basetypes.BoolType{},
									},
								},
								"new_app_threshold": basetypes.Int64Type{},
								"none": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"sync_to_peer": basetypes.BoolType{},
								"threshold":    basetypes.Int64Type{},
								"weekly": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"action":              basetypes.StringType{},
										"at":                  basetypes.StringType{},
										"day_of_week":         basetypes.StringType{},
										"disable_new_content": basetypes.BoolType{},
									},
								},
							},
						},
					},
				},
				"wildfire": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"recurring": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"every_15_mins": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"action":       basetypes.StringType{},
										"at":           basetypes.Int64Type{},
										"sync_to_peer": basetypes.BoolType{},
									},
								},
								"every_30_mins": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"action":       basetypes.StringType{},
										"at":           basetypes.Int64Type{},
										"sync_to_peer": basetypes.BoolType{},
									},
								},
								"every_hour": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"action":       basetypes.StringType{},
										"at":           basetypes.Int64Type{},
										"sync_to_peer": basetypes.BoolType{},
									},
								},
								"every_min": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"action":       basetypes.StringType{},
										"sync_to_peer": basetypes.BoolType{},
									},
								},
								"none": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"real_time": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
							},
						},
					},
				},
			},
		},
	}
}

// AttrType returns the attribute type for a list of UpdateSchedule objects.
func (o UpdateSchedule) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the UpdateScheduleUpdateSchedule model.
func (o UpdateScheduleUpdateSchedule) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"anti_virus": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"recurring": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"daily": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"action": basetypes.StringType{},
								"at":     basetypes.StringType{},
							},
						},
						"hourly": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"action": basetypes.StringType{},
								"at":     basetypes.Int64Type{},
							},
						},
						"none": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"sync_to_peer": basetypes.BoolType{},
						"threshold":    basetypes.Int64Type{},
						"weekly": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"action":      basetypes.StringType{},
								"at":          basetypes.StringType{},
								"day_of_week": basetypes.StringType{},
							},
						},
					},
				},
			},
		},
		"threats": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"recurring": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"daily": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"action":              basetypes.StringType{},
								"at":                  basetypes.StringType{},
								"disable_new_content": basetypes.BoolType{},
							},
						},
						"every_30_mins": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"action":              basetypes.StringType{},
								"at":                  basetypes.Int64Type{},
								"disable_new_content": basetypes.BoolType{},
							},
						},
						"hourly": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"action":              basetypes.StringType{},
								"at":                  basetypes.Float64Type{},
								"disable_new_content": basetypes.BoolType{},
							},
						},
						"new_app_threshold": basetypes.Int64Type{},
						"none": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"sync_to_peer": basetypes.BoolType{},
						"threshold":    basetypes.Int64Type{},
						"weekly": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"action":              basetypes.StringType{},
								"at":                  basetypes.StringType{},
								"day_of_week":         basetypes.StringType{},
								"disable_new_content": basetypes.BoolType{},
							},
						},
					},
				},
			},
		},
		"wildfire": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"recurring": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"every_15_mins": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"action":       basetypes.StringType{},
								"at":           basetypes.Int64Type{},
								"sync_to_peer": basetypes.BoolType{},
							},
						},
						"every_30_mins": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"action":       basetypes.StringType{},
								"at":           basetypes.Int64Type{},
								"sync_to_peer": basetypes.BoolType{},
							},
						},
						"every_hour": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"action":       basetypes.StringType{},
								"at":           basetypes.Int64Type{},
								"sync_to_peer": basetypes.BoolType{},
							},
						},
						"every_min": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"action":       basetypes.StringType{},
								"sync_to_peer": basetypes.BoolType{},
							},
						},
						"none": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"real_time": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
					},
				},
			},
		},
	}
}

// AttrType returns the attribute type for a list of UpdateScheduleUpdateSchedule objects.
func (o UpdateScheduleUpdateSchedule) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the UpdateScheduleUpdateScheduleAntiVirus model.
func (o UpdateScheduleUpdateScheduleAntiVirus) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"recurring": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"daily": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"action": basetypes.StringType{},
						"at":     basetypes.StringType{},
					},
				},
				"hourly": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"action": basetypes.StringType{},
						"at":     basetypes.Int64Type{},
					},
				},
				"none": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"sync_to_peer": basetypes.BoolType{},
				"threshold":    basetypes.Int64Type{},
				"weekly": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"action":      basetypes.StringType{},
						"at":          basetypes.StringType{},
						"day_of_week": basetypes.StringType{},
					},
				},
			},
		},
	}
}

// AttrType returns the attribute type for a list of UpdateScheduleUpdateScheduleAntiVirus objects.
func (o UpdateScheduleUpdateScheduleAntiVirus) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the UpdateScheduleUpdateScheduleAntiVirusRecurring model.
func (o UpdateScheduleUpdateScheduleAntiVirusRecurring) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"daily": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"action": basetypes.StringType{},
				"at":     basetypes.StringType{},
			},
		},
		"hourly": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"action": basetypes.StringType{},
				"at":     basetypes.Int64Type{},
			},
		},
		"none": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
		"sync_to_peer": basetypes.BoolType{},
		"threshold":    basetypes.Int64Type{},
		"weekly": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"action":      basetypes.StringType{},
				"at":          basetypes.StringType{},
				"day_of_week": basetypes.StringType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of UpdateScheduleUpdateScheduleAntiVirusRecurring objects.
func (o UpdateScheduleUpdateScheduleAntiVirusRecurring) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the UpdateScheduleUpdateScheduleAntiVirusRecurringDaily model.
func (o UpdateScheduleUpdateScheduleAntiVirusRecurringDaily) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"action": basetypes.StringType{},
		"at":     basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of UpdateScheduleUpdateScheduleAntiVirusRecurringDaily objects.
func (o UpdateScheduleUpdateScheduleAntiVirusRecurringDaily) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the UpdateScheduleUpdateScheduleAntiVirusRecurringHourly model.
func (o UpdateScheduleUpdateScheduleAntiVirusRecurringHourly) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"action": basetypes.StringType{},
		"at":     basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of UpdateScheduleUpdateScheduleAntiVirusRecurringHourly objects.
func (o UpdateScheduleUpdateScheduleAntiVirusRecurringHourly) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the UpdateScheduleUpdateScheduleAntiVirusRecurringWeekly model.
func (o UpdateScheduleUpdateScheduleAntiVirusRecurringWeekly) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"action":      basetypes.StringType{},
		"at":          basetypes.StringType{},
		"day_of_week": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of UpdateScheduleUpdateScheduleAntiVirusRecurringWeekly objects.
func (o UpdateScheduleUpdateScheduleAntiVirusRecurringWeekly) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the UpdateScheduleUpdateScheduleThreats model.
func (o UpdateScheduleUpdateScheduleThreats) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"recurring": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"daily": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"action":              basetypes.StringType{},
						"at":                  basetypes.StringType{},
						"disable_new_content": basetypes.BoolType{},
					},
				},
				"every_30_mins": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"action":              basetypes.StringType{},
						"at":                  basetypes.Int64Type{},
						"disable_new_content": basetypes.BoolType{},
					},
				},
				"hourly": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"action":              basetypes.StringType{},
						"at":                  basetypes.Float64Type{},
						"disable_new_content": basetypes.BoolType{},
					},
				},
				"new_app_threshold": basetypes.Int64Type{},
				"none": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"sync_to_peer": basetypes.BoolType{},
				"threshold":    basetypes.Int64Type{},
				"weekly": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"action":              basetypes.StringType{},
						"at":                  basetypes.StringType{},
						"day_of_week":         basetypes.StringType{},
						"disable_new_content": basetypes.BoolType{},
					},
				},
			},
		},
	}
}

// AttrType returns the attribute type for a list of UpdateScheduleUpdateScheduleThreats objects.
func (o UpdateScheduleUpdateScheduleThreats) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the UpdateScheduleUpdateScheduleThreatsRecurring model.
func (o UpdateScheduleUpdateScheduleThreatsRecurring) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"daily": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"action":              basetypes.StringType{},
				"at":                  basetypes.StringType{},
				"disable_new_content": basetypes.BoolType{},
			},
		},
		"every_30_mins": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"action":              basetypes.StringType{},
				"at":                  basetypes.Int64Type{},
				"disable_new_content": basetypes.BoolType{},
			},
		},
		"hourly": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"action":              basetypes.StringType{},
				"at":                  basetypes.Float64Type{},
				"disable_new_content": basetypes.BoolType{},
			},
		},
		"new_app_threshold": basetypes.Int64Type{},
		"none": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
		"sync_to_peer": basetypes.BoolType{},
		"threshold":    basetypes.Int64Type{},
		"weekly": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"action":              basetypes.StringType{},
				"at":                  basetypes.StringType{},
				"day_of_week":         basetypes.StringType{},
				"disable_new_content": basetypes.BoolType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of UpdateScheduleUpdateScheduleThreatsRecurring objects.
func (o UpdateScheduleUpdateScheduleThreatsRecurring) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the UpdateScheduleUpdateScheduleThreatsRecurringDaily model.
func (o UpdateScheduleUpdateScheduleThreatsRecurringDaily) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"action":              basetypes.StringType{},
		"at":                  basetypes.StringType{},
		"disable_new_content": basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of UpdateScheduleUpdateScheduleThreatsRecurringDaily objects.
func (o UpdateScheduleUpdateScheduleThreatsRecurringDaily) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the UpdateScheduleUpdateScheduleThreatsRecurringEvery30Mins model.
func (o UpdateScheduleUpdateScheduleThreatsRecurringEvery30Mins) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"action":              basetypes.StringType{},
		"at":                  basetypes.Int64Type{},
		"disable_new_content": basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of UpdateScheduleUpdateScheduleThreatsRecurringEvery30Mins objects.
func (o UpdateScheduleUpdateScheduleThreatsRecurringEvery30Mins) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the UpdateScheduleUpdateScheduleThreatsRecurringHourly model.
func (o UpdateScheduleUpdateScheduleThreatsRecurringHourly) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"action":              basetypes.StringType{},
		"at":                  basetypes.Float64Type{},
		"disable_new_content": basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of UpdateScheduleUpdateScheduleThreatsRecurringHourly objects.
func (o UpdateScheduleUpdateScheduleThreatsRecurringHourly) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the UpdateScheduleUpdateScheduleThreatsRecurringWeekly model.
func (o UpdateScheduleUpdateScheduleThreatsRecurringWeekly) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"action":              basetypes.StringType{},
		"at":                  basetypes.StringType{},
		"day_of_week":         basetypes.StringType{},
		"disable_new_content": basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of UpdateScheduleUpdateScheduleThreatsRecurringWeekly objects.
func (o UpdateScheduleUpdateScheduleThreatsRecurringWeekly) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the UpdateScheduleUpdateScheduleWildfire model.
func (o UpdateScheduleUpdateScheduleWildfire) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"recurring": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"every_15_mins": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"action":       basetypes.StringType{},
						"at":           basetypes.Int64Type{},
						"sync_to_peer": basetypes.BoolType{},
					},
				},
				"every_30_mins": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"action":       basetypes.StringType{},
						"at":           basetypes.Int64Type{},
						"sync_to_peer": basetypes.BoolType{},
					},
				},
				"every_hour": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"action":       basetypes.StringType{},
						"at":           basetypes.Int64Type{},
						"sync_to_peer": basetypes.BoolType{},
					},
				},
				"every_min": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"action":       basetypes.StringType{},
						"sync_to_peer": basetypes.BoolType{},
					},
				},
				"none": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"real_time": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
			},
		},
	}
}

// AttrType returns the attribute type for a list of UpdateScheduleUpdateScheduleWildfire objects.
func (o UpdateScheduleUpdateScheduleWildfire) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the UpdateScheduleUpdateScheduleWildfireRecurring model.
func (o UpdateScheduleUpdateScheduleWildfireRecurring) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"every_15_mins": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"action":       basetypes.StringType{},
				"at":           basetypes.Int64Type{},
				"sync_to_peer": basetypes.BoolType{},
			},
		},
		"every_30_mins": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"action":       basetypes.StringType{},
				"at":           basetypes.Int64Type{},
				"sync_to_peer": basetypes.BoolType{},
			},
		},
		"every_hour": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"action":       basetypes.StringType{},
				"at":           basetypes.Int64Type{},
				"sync_to_peer": basetypes.BoolType{},
			},
		},
		"every_min": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"action":       basetypes.StringType{},
				"sync_to_peer": basetypes.BoolType{},
			},
		},
		"none": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
		"real_time": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
	}
}

// AttrType returns the attribute type for a list of UpdateScheduleUpdateScheduleWildfireRecurring objects.
func (o UpdateScheduleUpdateScheduleWildfireRecurring) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the UpdateScheduleUpdateScheduleWildfireRecurringEvery15Mins model.
func (o UpdateScheduleUpdateScheduleWildfireRecurringEvery15Mins) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"action":       basetypes.StringType{},
		"at":           basetypes.Int64Type{},
		"sync_to_peer": basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of UpdateScheduleUpdateScheduleWildfireRecurringEvery15Mins objects.
func (o UpdateScheduleUpdateScheduleWildfireRecurringEvery15Mins) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the UpdateScheduleUpdateScheduleWildfireRecurringEvery30Mins model.
func (o UpdateScheduleUpdateScheduleWildfireRecurringEvery30Mins) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"action":       basetypes.StringType{},
		"at":           basetypes.Int64Type{},
		"sync_to_peer": basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of UpdateScheduleUpdateScheduleWildfireRecurringEvery30Mins objects.
func (o UpdateScheduleUpdateScheduleWildfireRecurringEvery30Mins) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the UpdateScheduleUpdateScheduleWildfireRecurringEveryHour model.
func (o UpdateScheduleUpdateScheduleWildfireRecurringEveryHour) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"action":       basetypes.StringType{},
		"at":           basetypes.Int64Type{},
		"sync_to_peer": basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of UpdateScheduleUpdateScheduleWildfireRecurringEveryHour objects.
func (o UpdateScheduleUpdateScheduleWildfireRecurringEveryHour) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the UpdateScheduleUpdateScheduleWildfireRecurringEveryMin model.
func (o UpdateScheduleUpdateScheduleWildfireRecurringEveryMin) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"action":       basetypes.StringType{},
		"sync_to_peer": basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of UpdateScheduleUpdateScheduleWildfireRecurringEveryMin objects.
func (o UpdateScheduleUpdateScheduleWildfireRecurringEveryMin) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// UpdateScheduleResourceSchema defines the schema for UpdateSchedule resource
var UpdateScheduleResourceSchema = schema.Schema{
	MarkdownDescription: "UpdateSchedule resource",
	Attributes: map[string]schema.Attribute{
		"device": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.ExactlyOneOf(
					path.MatchRelative().AtParent().AtName("folder"),
					path.MatchRelative().AtParent().AtName("snippet"),
				),
				stringvalidator.LengthAtMost(64),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z\\d-_\\. ]+$"), "pattern must match "+"^[a-zA-Z\\d-_\\. ]+$"),
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
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z\\d-_\\. ]+$"), "pattern must match "+"^[a-zA-Z\\d-_\\. ]+$"),
			},
			MarkdownDescription: "The folder in which the resource is defined\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"id": schema.StringAttribute{
			MarkdownDescription: "UUID of the resource",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"snippet": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.ExactlyOneOf(
					path.MatchRelative().AtParent().AtName("device"),
					path.MatchRelative().AtParent().AtName("folder"),
				),
				stringvalidator.LengthAtMost(64),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z\\d-_\\. ]+$"), "pattern must match "+"^[a-zA-Z\\d-_\\. ]+$"),
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
		"update_schedule": schema.SingleNestedAttribute{
			MarkdownDescription: "Update schedule",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"anti_virus": schema.SingleNestedAttribute{
					MarkdownDescription: "Anti virus",
					Required:            true,
					Attributes: map[string]schema.Attribute{
						"recurring": schema.SingleNestedAttribute{
							MarkdownDescription: "Recurring",
							Required:            true,
							Attributes: map[string]schema.Attribute{
								"daily": schema.SingleNestedAttribute{
									Validators: []validator.Object{
										objectvalidator.ExactlyOneOf(
											path.MatchRelative().AtParent().AtName("hourly"),
											path.MatchRelative().AtParent().AtName("none"),
											path.MatchRelative().AtParent().AtName("weekly"),
										),
									},
									MarkdownDescription: "Daily\n> ℹ️ **Note:** You must specify exactly one of `daily`, `hourly`, `none`, and `weekly`.",
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"action": schema.StringAttribute{
											Validators: []validator.String{
												stringvalidator.OneOf("download-only", "download-and-install"),
											},
											MarkdownDescription: "Action",
											Optional:            true,
										},
										"at": schema.StringAttribute{
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile("^(0[0-9]|1[0-9]|2[0-3]):[0-5][0-9]$"), "pattern must match "+"^(0[0-9]|1[0-9]|2[0-3]):[0-5][0-9]$"),
											},
											MarkdownDescription: "At",
											Required:            true,
										},
									},
								},
								"hourly": schema.SingleNestedAttribute{
									Validators: []validator.Object{
										objectvalidator.ExactlyOneOf(
											path.MatchRelative().AtParent().AtName("daily"),
											path.MatchRelative().AtParent().AtName("none"),
											path.MatchRelative().AtParent().AtName("weekly"),
										),
									},
									MarkdownDescription: "Hourly\n> ℹ️ **Note:** You must specify exactly one of `daily`, `hourly`, `none`, and `weekly`.",
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"action": schema.StringAttribute{
											Validators: []validator.String{
												stringvalidator.OneOf("download-only", "download-and-install"),
											},
											MarkdownDescription: "Action",
											Optional:            true,
										},
										"at": schema.Int64Attribute{
											Validators: []validator.Int64{
												int64validator.Between(0, 59),
											},
											MarkdownDescription: "At",
											Required:            true,
										},
									},
								},
								"none": schema.SingleNestedAttribute{
									Validators: []validator.Object{
										objectvalidator.ExactlyOneOf(
											path.MatchRelative().AtParent().AtName("daily"),
											path.MatchRelative().AtParent().AtName("hourly"),
											path.MatchRelative().AtParent().AtName("weekly"),
										),
									},
									MarkdownDescription: "None\n> ℹ️ **Note:** You must specify exactly one of `daily`, `hourly`, `none`, and `weekly`.",
									Optional:            true,
									Attributes:          map[string]schema.Attribute{},
								},
								"sync_to_peer": schema.BoolAttribute{
									MarkdownDescription: "Sync to peer",
									Required:            true,
								},
								"threshold": schema.Int64Attribute{
									Validators: []validator.Int64{
										int64validator.Between(1, 336),
									},
									MarkdownDescription: "Threshold",
									Optional:            true,
								},
								"weekly": schema.SingleNestedAttribute{
									Validators: []validator.Object{
										objectvalidator.ExactlyOneOf(
											path.MatchRelative().AtParent().AtName("daily"),
											path.MatchRelative().AtParent().AtName("hourly"),
											path.MatchRelative().AtParent().AtName("none"),
										),
									},
									MarkdownDescription: "Weekly\n> ℹ️ **Note:** You must specify exactly one of `daily`, `hourly`, `none`, and `weekly`.",
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"action": schema.StringAttribute{
											Validators: []validator.String{
												stringvalidator.OneOf("download-only", "download-and-install"),
											},
											MarkdownDescription: "Action",
											Optional:            true,
										},
										"at": schema.StringAttribute{
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile("^(0[0-9]|1[0-9]|2[0-3]):[0-5][0-9]$"), "pattern must match "+"^(0[0-9]|1[0-9]|2[0-3]):[0-5][0-9]$"),
											},
											MarkdownDescription: "At",
											Optional:            true,
										},
										"day_of_week": schema.StringAttribute{
											Validators: []validator.String{
												stringvalidator.OneOf("sunday", "monday", "tuesday", "wednesday", "thursday", "friday", "saturday"),
											},
											MarkdownDescription: "Day of week",
											Optional:            true,
										},
									},
								},
							},
						},
					},
				},
				"threats": schema.SingleNestedAttribute{
					MarkdownDescription: "Threats",
					Required:            true,
					Attributes: map[string]schema.Attribute{
						"recurring": schema.SingleNestedAttribute{
							MarkdownDescription: "Recurring",
							Required:            true,
							Attributes: map[string]schema.Attribute{
								"daily": schema.SingleNestedAttribute{
									Validators: []validator.Object{
										objectvalidator.ExactlyOneOf(
											path.MatchRelative().AtParent().AtName("every_30_mins"),
											path.MatchRelative().AtParent().AtName("hourly"),
											path.MatchRelative().AtParent().AtName("none"),
											path.MatchRelative().AtParent().AtName("weekly"),
										),
									},
									MarkdownDescription: "Daily\n> ℹ️ **Note:** You must specify exactly one of `daily`, `every_30_mins`, `hourly`, `none`, and `weekly`.",
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"action": schema.StringAttribute{
											Validators: []validator.String{
												stringvalidator.OneOf("download-only", "download-and-install"),
											},
											MarkdownDescription: "Action",
											Optional:            true,
										},
										"at": schema.StringAttribute{
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile("^(0[0-9]|1[0-9]|2[0-3]):[0-5][0-9]$"), "pattern must match "+"^(0[0-9]|1[0-9]|2[0-3]):[0-5][0-9]$"),
											},
											MarkdownDescription: "At",
											Required:            true,
										},
										"disable_new_content": schema.BoolAttribute{
											MarkdownDescription: "Disable new content",
											Optional:            true,
											Computed:            true,
											Default:             booldefault.StaticBool(false),
										},
									},
								},
								"every_30_mins": schema.SingleNestedAttribute{
									Validators: []validator.Object{
										objectvalidator.ExactlyOneOf(
											path.MatchRelative().AtParent().AtName("daily"),
											path.MatchRelative().AtParent().AtName("hourly"),
											path.MatchRelative().AtParent().AtName("none"),
											path.MatchRelative().AtParent().AtName("weekly"),
										),
									},
									MarkdownDescription: "Every30 mins\n> ℹ️ **Note:** You must specify exactly one of `daily`, `every_30_mins`, `hourly`, `none`, and `weekly`.",
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"action": schema.StringAttribute{
											Validators: []validator.String{
												stringvalidator.OneOf("download-only", "download-and-install"),
											},
											MarkdownDescription: "Action",
											Optional:            true,
										},
										"at": schema.Int64Attribute{
											Validators: []validator.Int64{
												int64validator.Between(0, 29),
											},
											MarkdownDescription: "At",
											Optional:            true,
											Computed:            true,
											Default:             int64default.StaticInt64(0),
										},
										"disable_new_content": schema.BoolAttribute{
											MarkdownDescription: "Disable new content",
											Optional:            true,
											Computed:            true,
											Default:             booldefault.StaticBool(false),
										},
									},
								},
								"hourly": schema.SingleNestedAttribute{
									Validators: []validator.Object{
										objectvalidator.ExactlyOneOf(
											path.MatchRelative().AtParent().AtName("daily"),
											path.MatchRelative().AtParent().AtName("every_30_mins"),
											path.MatchRelative().AtParent().AtName("none"),
											path.MatchRelative().AtParent().AtName("weekly"),
										),
									},
									MarkdownDescription: "Hourly\n> ℹ️ **Note:** You must specify exactly one of `daily`, `every_30_mins`, `hourly`, `none`, and `weekly`.",
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"action": schema.StringAttribute{
											Validators: []validator.String{
												stringvalidator.OneOf("download-only", "download-and-install"),
											},
											MarkdownDescription: "Action",
											Optional:            true,
										},
										"at": schema.Float64Attribute{
											Validators: []validator.Float64{
												float64validator.Between(0.000000, 59.000000),
											},
											MarkdownDescription: "At",
											Required:            true,
										},
										"disable_new_content": schema.BoolAttribute{
											MarkdownDescription: "Disable new content",
											Optional:            true,
											Computed:            true,
											Default:             booldefault.StaticBool(false),
										},
									},
								},
								"new_app_threshold": schema.Int64Attribute{
									Validators: []validator.Int64{
										int64validator.Between(1, 336),
									},
									MarkdownDescription: "New app threshold",
									Optional:            true,
								},
								"none": schema.SingleNestedAttribute{
									Validators: []validator.Object{
										objectvalidator.ExactlyOneOf(
											path.MatchRelative().AtParent().AtName("daily"),
											path.MatchRelative().AtParent().AtName("every_30_mins"),
											path.MatchRelative().AtParent().AtName("hourly"),
											path.MatchRelative().AtParent().AtName("weekly"),
										),
									},
									MarkdownDescription: "None\n> ℹ️ **Note:** You must specify exactly one of `daily`, `every_30_mins`, `hourly`, `none`, and `weekly`.",
									Optional:            true,
									Attributes:          map[string]schema.Attribute{},
								},
								"sync_to_peer": schema.BoolAttribute{
									MarkdownDescription: "Sync to peer",
									Required:            true,
								},
								"threshold": schema.Int64Attribute{
									Validators: []validator.Int64{
										int64validator.Between(1, 336),
									},
									MarkdownDescription: "Threshold",
									Optional:            true,
								},
								"weekly": schema.SingleNestedAttribute{
									Validators: []validator.Object{
										objectvalidator.ExactlyOneOf(
											path.MatchRelative().AtParent().AtName("daily"),
											path.MatchRelative().AtParent().AtName("every_30_mins"),
											path.MatchRelative().AtParent().AtName("hourly"),
											path.MatchRelative().AtParent().AtName("none"),
										),
									},
									MarkdownDescription: "Weekly\n> ℹ️ **Note:** You must specify exactly one of `daily`, `every_30_mins`, `hourly`, `none`, and `weekly`.",
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"action": schema.StringAttribute{
											Validators: []validator.String{
												stringvalidator.OneOf("download-only", "download-and-install"),
											},
											MarkdownDescription: "Action",
											Optional:            true,
										},
										"at": schema.StringAttribute{
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile("^(0[0-9]|1[0-9]|2[0-3]):[0-5][0-9]$"), "pattern must match "+"^(0[0-9]|1[0-9]|2[0-3]):[0-5][0-9]$"),
											},
											MarkdownDescription: "At",
											Required:            true,
										},
										"day_of_week": schema.StringAttribute{
											Validators: []validator.String{
												stringvalidator.OneOf("sunday", "monday", "tuesday", "wednesday", "thursday", "friday", "saturday"),
											},
											MarkdownDescription: "Day of week",
											Required:            true,
										},
										"disable_new_content": schema.BoolAttribute{
											MarkdownDescription: "Disable new content",
											Optional:            true,
											Computed:            true,
											Default:             booldefault.StaticBool(false),
										},
									},
								},
							},
						},
					},
				},
				"wildfire": schema.SingleNestedAttribute{
					MarkdownDescription: "Wildfire",
					Required:            true,
					Attributes: map[string]schema.Attribute{
						"recurring": schema.SingleNestedAttribute{
							MarkdownDescription: "Recurring",
							Required:            true,
							Attributes: map[string]schema.Attribute{
								"every_15_mins": schema.SingleNestedAttribute{
									Validators: []validator.Object{
										objectvalidator.ExactlyOneOf(
											path.MatchRelative().AtParent().AtName("every_30_mins"),
											path.MatchRelative().AtParent().AtName("every_hour"),
											path.MatchRelative().AtParent().AtName("every_min"),
											path.MatchRelative().AtParent().AtName("none"),
											path.MatchRelative().AtParent().AtName("real_time"),
										),
									},
									MarkdownDescription: "Every15 mins\n> ℹ️ **Note:** You must specify exactly one of `every_15_mins`, `every_30_mins`, `every_hour`, `every_min`, `none`, and `real_time`.",
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"action": schema.StringAttribute{
											Validators: []validator.String{
												stringvalidator.OneOf("download-only", "download-and-install"),
											},
											MarkdownDescription: "Action",
											Optional:            true,
										},
										"at": schema.Int64Attribute{
											Validators: []validator.Int64{
												int64validator.Between(0, 14),
											},
											MarkdownDescription: "At",
											Optional:            true,
											Computed:            true,
											Default:             int64default.StaticInt64(0),
										},
										"sync_to_peer": schema.BoolAttribute{
											MarkdownDescription: "Sync to peer",
											Optional:            true,
											Computed:            true,
											Default:             booldefault.StaticBool(false),
										},
									},
								},
								"every_30_mins": schema.SingleNestedAttribute{
									Validators: []validator.Object{
										objectvalidator.ExactlyOneOf(
											path.MatchRelative().AtParent().AtName("every_15_mins"),
											path.MatchRelative().AtParent().AtName("every_hour"),
											path.MatchRelative().AtParent().AtName("every_min"),
											path.MatchRelative().AtParent().AtName("none"),
											path.MatchRelative().AtParent().AtName("real_time"),
										),
									},
									MarkdownDescription: "Every30 mins\n> ℹ️ **Note:** You must specify exactly one of `every_15_mins`, `every_30_mins`, `every_hour`, `every_min`, `none`, and `real_time`.",
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"action": schema.StringAttribute{
											Validators: []validator.String{
												stringvalidator.OneOf("download-only", "download-and-install"),
											},
											MarkdownDescription: "Action",
											Optional:            true,
										},
										"at": schema.Int64Attribute{
											Validators: []validator.Int64{
												int64validator.Between(0, 29),
											},
											MarkdownDescription: "At",
											Optional:            true,
											Computed:            true,
											Default:             int64default.StaticInt64(0),
										},
										"sync_to_peer": schema.BoolAttribute{
											MarkdownDescription: "Sync to peer",
											Optional:            true,
											Computed:            true,
											Default:             booldefault.StaticBool(false),
										},
									},
								},
								"every_hour": schema.SingleNestedAttribute{
									Validators: []validator.Object{
										objectvalidator.ExactlyOneOf(
											path.MatchRelative().AtParent().AtName("every_15_mins"),
											path.MatchRelative().AtParent().AtName("every_30_mins"),
											path.MatchRelative().AtParent().AtName("every_min"),
											path.MatchRelative().AtParent().AtName("none"),
											path.MatchRelative().AtParent().AtName("real_time"),
										),
									},
									MarkdownDescription: "Every hour\n> ℹ️ **Note:** You must specify exactly one of `every_15_mins`, `every_30_mins`, `every_hour`, `every_min`, `none`, and `real_time`.",
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"action": schema.StringAttribute{
											Validators: []validator.String{
												stringvalidator.OneOf("download-only", "download-and-install"),
											},
											MarkdownDescription: "Action",
											Optional:            true,
										},
										"at": schema.Int64Attribute{
											Validators: []validator.Int64{
												int64validator.Between(0, 59),
											},
											MarkdownDescription: "At",
											Optional:            true,
											Computed:            true,
											Default:             int64default.StaticInt64(0),
										},
										"sync_to_peer": schema.BoolAttribute{
											MarkdownDescription: "Sync to peer",
											Optional:            true,
											Computed:            true,
											Default:             booldefault.StaticBool(false),
										},
									},
								},
								"every_min": schema.SingleNestedAttribute{
									Validators: []validator.Object{
										objectvalidator.ExactlyOneOf(
											path.MatchRelative().AtParent().AtName("every_15_mins"),
											path.MatchRelative().AtParent().AtName("every_30_mins"),
											path.MatchRelative().AtParent().AtName("every_hour"),
											path.MatchRelative().AtParent().AtName("none"),
											path.MatchRelative().AtParent().AtName("real_time"),
										),
									},
									MarkdownDescription: "Every min\n> ℹ️ **Note:** You must specify exactly one of `every_15_mins`, `every_30_mins`, `every_hour`, `every_min`, `none`, and `real_time`.",
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"action": schema.StringAttribute{
											Validators: []validator.String{
												stringvalidator.OneOf("download-only", "download-and-install"),
											},
											MarkdownDescription: "Action",
											Optional:            true,
										},
										"sync_to_peer": schema.BoolAttribute{
											MarkdownDescription: "Sync to peer",
											Optional:            true,
											Computed:            true,
											Default:             booldefault.StaticBool(false),
										},
									},
								},
								"none": schema.SingleNestedAttribute{
									Validators: []validator.Object{
										objectvalidator.ExactlyOneOf(
											path.MatchRelative().AtParent().AtName("every_15_mins"),
											path.MatchRelative().AtParent().AtName("every_30_mins"),
											path.MatchRelative().AtParent().AtName("every_hour"),
											path.MatchRelative().AtParent().AtName("every_min"),
											path.MatchRelative().AtParent().AtName("real_time"),
										),
									},
									MarkdownDescription: "None\n> ℹ️ **Note:** You must specify exactly one of `every_15_mins`, `every_30_mins`, `every_hour`, `every_min`, `none`, and `real_time`.",
									Optional:            true,
									Attributes:          map[string]schema.Attribute{},
								},
								"real_time": schema.SingleNestedAttribute{
									Validators: []validator.Object{
										objectvalidator.ExactlyOneOf(
											path.MatchRelative().AtParent().AtName("every_15_mins"),
											path.MatchRelative().AtParent().AtName("every_30_mins"),
											path.MatchRelative().AtParent().AtName("every_hour"),
											path.MatchRelative().AtParent().AtName("every_min"),
											path.MatchRelative().AtParent().AtName("none"),
										),
									},
									MarkdownDescription: "Real time\n> ℹ️ **Note:** You must specify exactly one of `every_15_mins`, `every_30_mins`, `every_hour`, `every_min`, `none`, and `real_time`.",
									Optional:            true,
									Attributes:          map[string]schema.Attribute{},
								},
							},
						},
					},
				},
			},
		},
	},
}

// UpdateScheduleDataSourceSchema defines the schema for UpdateSchedule data source
var UpdateScheduleDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "UpdateSchedule data source",
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
			MarkdownDescription: "UUID of the resource",
			Required:            true,
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
		"update_schedule": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Update schedule",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"anti_virus": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Anti virus",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"recurring": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Recurring",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"daily": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Daily\n> ℹ️ **Note:** You must specify exactly one of `daily`, `hourly`, `none`, and `weekly`.",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"action": dsschema.StringAttribute{
											MarkdownDescription: "Action",
											Computed:            true,
										},
										"at": dsschema.StringAttribute{
											MarkdownDescription: "At",
											Computed:            true,
										},
									},
								},
								"hourly": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Hourly\n> ℹ️ **Note:** You must specify exactly one of `daily`, `hourly`, `none`, and `weekly`.",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"action": dsschema.StringAttribute{
											MarkdownDescription: "Action",
											Computed:            true,
										},
										"at": dsschema.Int64Attribute{
											MarkdownDescription: "At",
											Computed:            true,
										},
									},
								},
								"none": dsschema.SingleNestedAttribute{
									MarkdownDescription: "None\n> ℹ️ **Note:** You must specify exactly one of `daily`, `hourly`, `none`, and `weekly`.",
									Computed:            true,
									Attributes:          map[string]dsschema.Attribute{},
								},
								"sync_to_peer": dsschema.BoolAttribute{
									MarkdownDescription: "Sync to peer",
									Computed:            true,
								},
								"threshold": dsschema.Int64Attribute{
									MarkdownDescription: "Threshold",
									Computed:            true,
								},
								"weekly": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Weekly\n> ℹ️ **Note:** You must specify exactly one of `daily`, `hourly`, `none`, and `weekly`.",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"action": dsschema.StringAttribute{
											MarkdownDescription: "Action",
											Computed:            true,
										},
										"at": dsschema.StringAttribute{
											MarkdownDescription: "At",
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
					},
				},
				"threats": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Threats",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"recurring": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Recurring",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"daily": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Daily\n> ℹ️ **Note:** You must specify exactly one of `daily`, `every_30_mins`, `hourly`, `none`, and `weekly`.",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"action": dsschema.StringAttribute{
											MarkdownDescription: "Action",
											Computed:            true,
										},
										"at": dsschema.StringAttribute{
											MarkdownDescription: "At",
											Computed:            true,
										},
										"disable_new_content": dsschema.BoolAttribute{
											MarkdownDescription: "Disable new content",
											Computed:            true,
										},
									},
								},
								"every_30_mins": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Every30 mins\n> ℹ️ **Note:** You must specify exactly one of `daily`, `every_30_mins`, `hourly`, `none`, and `weekly`.",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"action": dsschema.StringAttribute{
											MarkdownDescription: "Action",
											Computed:            true,
										},
										"at": dsschema.Int64Attribute{
											MarkdownDescription: "At",
											Computed:            true,
										},
										"disable_new_content": dsschema.BoolAttribute{
											MarkdownDescription: "Disable new content",
											Computed:            true,
										},
									},
								},
								"hourly": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Hourly\n> ℹ️ **Note:** You must specify exactly one of `daily`, `every_30_mins`, `hourly`, `none`, and `weekly`.",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"action": dsschema.StringAttribute{
											MarkdownDescription: "Action",
											Computed:            true,
										},
										"at": dsschema.Float64Attribute{
											MarkdownDescription: "At",
											Computed:            true,
										},
										"disable_new_content": dsschema.BoolAttribute{
											MarkdownDescription: "Disable new content",
											Computed:            true,
										},
									},
								},
								"new_app_threshold": dsschema.Int64Attribute{
									MarkdownDescription: "New app threshold",
									Computed:            true,
								},
								"none": dsschema.SingleNestedAttribute{
									MarkdownDescription: "None\n> ℹ️ **Note:** You must specify exactly one of `daily`, `every_30_mins`, `hourly`, `none`, and `weekly`.",
									Computed:            true,
									Attributes:          map[string]dsschema.Attribute{},
								},
								"sync_to_peer": dsschema.BoolAttribute{
									MarkdownDescription: "Sync to peer",
									Computed:            true,
								},
								"threshold": dsschema.Int64Attribute{
									MarkdownDescription: "Threshold",
									Computed:            true,
								},
								"weekly": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Weekly\n> ℹ️ **Note:** You must specify exactly one of `daily`, `every_30_mins`, `hourly`, `none`, and `weekly`.",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"action": dsschema.StringAttribute{
											MarkdownDescription: "Action",
											Computed:            true,
										},
										"at": dsschema.StringAttribute{
											MarkdownDescription: "At",
											Computed:            true,
										},
										"day_of_week": dsschema.StringAttribute{
											MarkdownDescription: "Day of week",
											Computed:            true,
										},
										"disable_new_content": dsschema.BoolAttribute{
											MarkdownDescription: "Disable new content",
											Computed:            true,
										},
									},
								},
							},
						},
					},
				},
				"wildfire": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Wildfire",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"recurring": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Recurring",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"every_15_mins": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Every15 mins\n> ℹ️ **Note:** You must specify exactly one of `every_15_mins`, `every_30_mins`, `every_hour`, `every_min`, `none`, and `real_time`.",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"action": dsschema.StringAttribute{
											MarkdownDescription: "Action",
											Computed:            true,
										},
										"at": dsschema.Int64Attribute{
											MarkdownDescription: "At",
											Computed:            true,
										},
										"sync_to_peer": dsschema.BoolAttribute{
											MarkdownDescription: "Sync to peer",
											Computed:            true,
										},
									},
								},
								"every_30_mins": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Every30 mins\n> ℹ️ **Note:** You must specify exactly one of `every_15_mins`, `every_30_mins`, `every_hour`, `every_min`, `none`, and `real_time`.",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"action": dsschema.StringAttribute{
											MarkdownDescription: "Action",
											Computed:            true,
										},
										"at": dsschema.Int64Attribute{
											MarkdownDescription: "At",
											Computed:            true,
										},
										"sync_to_peer": dsschema.BoolAttribute{
											MarkdownDescription: "Sync to peer",
											Computed:            true,
										},
									},
								},
								"every_hour": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Every hour\n> ℹ️ **Note:** You must specify exactly one of `every_15_mins`, `every_30_mins`, `every_hour`, `every_min`, `none`, and `real_time`.",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"action": dsschema.StringAttribute{
											MarkdownDescription: "Action",
											Computed:            true,
										},
										"at": dsschema.Int64Attribute{
											MarkdownDescription: "At",
											Computed:            true,
										},
										"sync_to_peer": dsschema.BoolAttribute{
											MarkdownDescription: "Sync to peer",
											Computed:            true,
										},
									},
								},
								"every_min": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Every min\n> ℹ️ **Note:** You must specify exactly one of `every_15_mins`, `every_30_mins`, `every_hour`, `every_min`, `none`, and `real_time`.",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"action": dsschema.StringAttribute{
											MarkdownDescription: "Action",
											Computed:            true,
										},
										"sync_to_peer": dsschema.BoolAttribute{
											MarkdownDescription: "Sync to peer",
											Computed:            true,
										},
									},
								},
								"none": dsschema.SingleNestedAttribute{
									MarkdownDescription: "None\n> ℹ️ **Note:** You must specify exactly one of `every_15_mins`, `every_30_mins`, `every_hour`, `every_min`, `none`, and `real_time`.",
									Computed:            true,
									Attributes:          map[string]dsschema.Attribute{},
								},
								"real_time": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Real time\n> ℹ️ **Note:** You must specify exactly one of `every_15_mins`, `every_30_mins`, `every_hour`, `every_min`, `none`, and `real_time`.",
									Computed:            true,
									Attributes:          map[string]dsschema.Attribute{},
								},
							},
						},
					},
				},
			},
		},
	},
}

// UpdateScheduleListModel represents the data model for a list data source.
type UpdateScheduleListModel struct {
	Tfid    types.String     `tfsdk:"tfid"`
	Data    []UpdateSchedule `tfsdk:"data"`
	Limit   types.Int64      `tfsdk:"limit"`
	Offset  types.Int64      `tfsdk:"offset"`
	Name    types.String     `tfsdk:"name"`
	Total   types.Int64      `tfsdk:"total"`
	Folder  types.String     `tfsdk:"folder"`
	Snippet types.String     `tfsdk:"snippet"`
	Device  types.String     `tfsdk:"device"`
}

// UpdateScheduleListDataSourceSchema defines the schema for a list data source.
var UpdateScheduleListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: UpdateScheduleDataSourceSchema.Attributes,
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
