package models

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	dsschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Package: config_operations
// This file contains models for the config_operations SDK package

// JobsResponse represents the Terraform model for JobsResponse
type JobsResponse struct {
	Tfid types.String        `tfsdk:"tfid"`
	Data basetypes.ListValue `tfsdk:"data"`
}

// Jobs represents a nested structure within the JobsResponse model
type Jobs struct {
	Description basetypes.StringValue `tfsdk:"description"`
	Details     basetypes.StringValue `tfsdk:"details"`
	DeviceName  basetypes.StringValue `tfsdk:"device_name"`
	EndTs       basetypes.StringValue `tfsdk:"end_ts"`
	Id          basetypes.StringValue `tfsdk:"id"`
	JobResult   basetypes.StringValue `tfsdk:"job_result"`
	JobStatus   basetypes.StringValue `tfsdk:"job_status"`
	JobType     basetypes.StringValue `tfsdk:"job_type"`
	ParentId    basetypes.StringValue `tfsdk:"parent_id"`
	Percent     basetypes.StringValue `tfsdk:"percent"`
	ResultStr   basetypes.StringValue `tfsdk:"result_str"`
	StartTs     basetypes.StringValue `tfsdk:"start_ts"`
	StatusStr   basetypes.StringValue `tfsdk:"status_str"`
	Summary     basetypes.StringValue `tfsdk:"summary"`
	TypeStr     basetypes.StringValue `tfsdk:"type_str"`
	Uname       basetypes.StringValue `tfsdk:"uname"`
}

// AttrTypes defines the attribute types for the JobsResponse model.
func (o JobsResponse) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid": basetypes.StringType{},
		"data": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"description": basetypes.StringType{},
				"details":     basetypes.StringType{},
				"device_name": basetypes.StringType{},
				"end_ts":      basetypes.StringType{},
				"id":          basetypes.StringType{},
				"job_result":  basetypes.StringType{},
				"job_status":  basetypes.StringType{},
				"job_type":    basetypes.StringType{},
				"parent_id":   basetypes.StringType{},
				"percent":     basetypes.StringType{},
				"result_str":  basetypes.StringType{},
				"start_ts":    basetypes.StringType{},
				"status_str":  basetypes.StringType{},
				"summary":     basetypes.StringType{},
				"type_str":    basetypes.StringType{},
				"uname":       basetypes.StringType{},
			},
		}},
	}
}

// AttrType returns the attribute type for a list of JobsResponse objects.
func (o JobsResponse) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the Jobs model.
func (o Jobs) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"description": basetypes.StringType{},
		"details":     basetypes.StringType{},
		"device_name": basetypes.StringType{},
		"end_ts":      basetypes.StringType{},
		"id":          basetypes.StringType{},
		"job_result":  basetypes.StringType{},
		"job_status":  basetypes.StringType{},
		"job_type":    basetypes.StringType{},
		"parent_id":   basetypes.StringType{},
		"percent":     basetypes.StringType{},
		"result_str":  basetypes.StringType{},
		"start_ts":    basetypes.StringType{},
		"status_str":  basetypes.StringType{},
		"summary":     basetypes.StringType{},
		"type_str":    basetypes.StringType{},
		"uname":       basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of Jobs objects.
func (o Jobs) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// JobsResponseResourceSchema defines the schema for JobsResponse resource
var JobsResponseResourceSchema = schema.Schema{
	MarkdownDescription: "JobsResponse resource",
	Attributes: map[string]schema.Attribute{
		"data": schema.ListNestedAttribute{
			MarkdownDescription: "Data",
			Optional:            true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"description": schema.StringAttribute{
						MarkdownDescription: "A description provided by the administrator or service account",
						Optional:            true,
						Computed:            true,
					},
					"details": schema.StringAttribute{
						MarkdownDescription: "JSON string with detailed errors or info",
						Optional:            true,
						Computed:            true,
					},
					"device_name": schema.StringAttribute{
						MarkdownDescription: "The name of the device",
						Required:            true,
					},
					"end_ts": schema.StringAttribute{
						MarkdownDescription: "The timestamp indicating when the job was finished",
						Required:            true,
					},
					"id": schema.StringAttribute{
						MarkdownDescription: "The job ID",
						Required:            true,
					},
					"job_result": schema.StringAttribute{
						MarkdownDescription: "The job result",
						Required:            true,
					},
					"job_status": schema.StringAttribute{
						MarkdownDescription: "The current status of the job",
						Required:            true,
					},
					"job_type": schema.StringAttribute{
						MarkdownDescription: "The job type",
						Required:            true,
					},
					"parent_id": schema.StringAttribute{
						MarkdownDescription: "The parent job ID",
						Required:            true,
					},
					"percent": schema.StringAttribute{
						MarkdownDescription: "Job completion percentage",
						Required:            true,
					},
					"result_str": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("OK", "FAIL", "PEND", "WAIT", "CANCELLED", "TIMEOUT"),
						},
						MarkdownDescription: "The result of the job",
						Required:            true,
					},
					"start_ts": schema.StringAttribute{
						MarkdownDescription: "The timestamp indicating when the job was created",
						Required:            true,
					},
					"status_str": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("ACT", "FIN", "PEND", "PUSHSENT", "PUSHFAIL", "PUSHABORT", "PUSHTIMEOUT"),
						},
						MarkdownDescription: "The current status of the job",
						Required:            true,
					},
					"summary": schema.StringAttribute{
						MarkdownDescription: "The completion summary of the job",
						Required:            true,
					},
					"type_str": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("CommitAll", "CommitAndPush", "NGFW-Bootstrap-Push", "Validate"),
						},
						MarkdownDescription: "The job type",
						Required:            true,
					},
					"uname": schema.StringAttribute{
						MarkdownDescription: "The administrator or service account that created the job",
						Required:            true,
					},
				},
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

// JobsResponseDataSourceSchema defines the schema for JobsResponse data source
var JobsResponseDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "JobsResponse data source",
	Attributes: map[string]dsschema.Attribute{
		"data": dsschema.ListNestedAttribute{
			MarkdownDescription: "Data",
			Computed:            true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: map[string]dsschema.Attribute{
					"description": dsschema.StringAttribute{
						MarkdownDescription: "A description provided by the administrator or service account",
						Computed:            true,
					},
					"details": dsschema.StringAttribute{
						MarkdownDescription: "JSON string with detailed errors or info",
						Computed:            true,
					},
					"device_name": dsschema.StringAttribute{
						MarkdownDescription: "The name of the device",
						Computed:            true,
					},
					"end_ts": dsschema.StringAttribute{
						MarkdownDescription: "The timestamp indicating when the job was finished",
						Computed:            true,
					},
					"id": dsschema.StringAttribute{
						MarkdownDescription: "The job ID",
						Computed:            true,
					},
					"job_result": dsschema.StringAttribute{
						MarkdownDescription: "The job result",
						Computed:            true,
					},
					"job_status": dsschema.StringAttribute{
						MarkdownDescription: "The current status of the job",
						Computed:            true,
					},
					"job_type": dsschema.StringAttribute{
						MarkdownDescription: "The job type",
						Computed:            true,
					},
					"parent_id": dsschema.StringAttribute{
						MarkdownDescription: "The parent job ID",
						Computed:            true,
					},
					"percent": dsschema.StringAttribute{
						MarkdownDescription: "Job completion percentage",
						Computed:            true,
					},
					"result_str": dsschema.StringAttribute{
						MarkdownDescription: "The result of the job",
						Computed:            true,
					},
					"start_ts": dsschema.StringAttribute{
						MarkdownDescription: "The timestamp indicating when the job was created",
						Computed:            true,
					},
					"status_str": dsschema.StringAttribute{
						MarkdownDescription: "The current status of the job",
						Computed:            true,
					},
					"summary": dsschema.StringAttribute{
						MarkdownDescription: "The completion summary of the job",
						Computed:            true,
					},
					"type_str": dsschema.StringAttribute{
						MarkdownDescription: "The job type",
						Computed:            true,
					},
					"uname": dsschema.StringAttribute{
						MarkdownDescription: "The administrator or service account that created the job",
						Computed:            true,
					},
				},
			},
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
	},
}

// JobsResponseListModel represents the data model for a list data source.
type JobsResponseListModel struct {
	Tfid    types.String   `tfsdk:"tfid"`
	Data    []JobsResponse `tfsdk:"data"`
	Limit   types.Int64    `tfsdk:"limit"`
	Offset  types.Int64    `tfsdk:"offset"`
	Name    types.String   `tfsdk:"name"`
	Total   types.Int64    `tfsdk:"total"`
	Folder  types.String   `tfsdk:"folder"`
	Snippet types.String   `tfsdk:"snippet"`
	Device  types.String   `tfsdk:"device"`
}

// JobsResponseListDataSourceSchema defines the schema for a list data source.
var JobsResponseListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: JobsResponseDataSourceSchema.Attributes,
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
