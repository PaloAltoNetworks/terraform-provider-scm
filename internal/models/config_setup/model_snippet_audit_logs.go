package models

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	dsschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Package: config_setup
// This file contains models for the config_setup SDK package

// SnippetAuditHistory represents the Terraform model for SnippetAuditHistory
type SnippetAuditHistory struct {
	Tfid                types.String          `tfsdk:"tfid"`
	Action              basetypes.StringValue `tfsdk:"action"`
	Created             basetypes.StringValue `tfsdk:"created"`
	Deleted             basetypes.Int64Value  `tfsdk:"deleted"`
	Details             basetypes.StringValue `tfsdk:"details"`
	Display             basetypes.Int64Value  `tfsdk:"display"`
	DonorCreated        basetypes.Int64Value  `tfsdk:"donor_created"`
	DonorTenantName     basetypes.StringValue `tfsdk:"donor_tenant_name"`
	DonorTsg            basetypes.StringValue `tfsdk:"donor_tsg"`
	Id                  basetypes.Int64Value  `tfsdk:"id"`
	RecipientTenantName basetypes.StringValue `tfsdk:"recipient_tenant_name"`
	RecipientTsg        basetypes.StringValue `tfsdk:"recipient_tsg"`
	SnippetUuid         basetypes.StringValue `tfsdk:"snippet_uuid"`
	User                basetypes.StringValue `tfsdk:"user"`
	Version             basetypes.StringValue `tfsdk:"version"`
}

// AttrTypes defines the attribute types for the SnippetAuditHistory model.
func (o SnippetAuditHistory) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":                  basetypes.StringType{},
		"action":                basetypes.StringType{},
		"created":               basetypes.StringType{},
		"deleted":               basetypes.Int64Type{},
		"details":               basetypes.StringType{},
		"display":               basetypes.Int64Type{},
		"donor_created":         basetypes.Int64Type{},
		"donor_tenant_name":     basetypes.StringType{},
		"donor_tsg":             basetypes.StringType{},
		"id":                    basetypes.Int64Type{},
		"recipient_tenant_name": basetypes.StringType{},
		"recipient_tsg":         basetypes.StringType{},
		"snippet_uuid":          basetypes.StringType{},
		"user":                  basetypes.StringType{},
		"version":               basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of SnippetAuditHistory objects.
func (o SnippetAuditHistory) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// SnippetAuditHistoryResourceSchema defines the schema for SnippetAuditHistory resource
var SnippetAuditHistoryResourceSchema = schema.Schema{
	MarkdownDescription: "SnippetAuditHistory resource",
	Attributes: map[string]schema.Attribute{
		"action": schema.StringAttribute{
			MarkdownDescription: "Action",
			Computed:            true,
		},
		"created": schema.StringAttribute{
			MarkdownDescription: "Created",
			Computed:            true,
		},
		"deleted": schema.Int64Attribute{
			MarkdownDescription: "Deleted",
			Computed:            true,
		},
		"details": schema.StringAttribute{
			MarkdownDescription: "Details",
			Computed:            true,
		},
		"display": schema.Int64Attribute{
			MarkdownDescription: "Display",
			Computed:            true,
		},
		"donor_created": schema.Int64Attribute{
			MarkdownDescription: "Donor created",
			Computed:            true,
		},
		"donor_tenant_name": schema.StringAttribute{
			MarkdownDescription: "Donor tenant name",
			Computed:            true,
		},
		"donor_tsg": schema.StringAttribute{
			MarkdownDescription: "Donor tsg",
			Computed:            true,
		},
		"id": schema.Int64Attribute{
			MarkdownDescription: "Id",
			Computed:            true,
			PlanModifiers: []planmodifier.Int64{
				int64planmodifier.UseStateForUnknown(),
			},
		},
		"recipient_tenant_name": schema.StringAttribute{
			MarkdownDescription: "Recipient tenant name",
			Computed:            true,
		},
		"recipient_tsg": schema.StringAttribute{
			MarkdownDescription: "Recipient tsg",
			Computed:            true,
		},
		"snippet_uuid": schema.StringAttribute{
			MarkdownDescription: "Snippet uuid",
			Computed:            true,
		},
		"tfid": schema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"user": schema.StringAttribute{
			MarkdownDescription: "User",
			Computed:            true,
		},
		"version": schema.StringAttribute{
			MarkdownDescription: "Version",
			Computed:            true,
		},
	},
}

// SnippetAuditHistoryDataSourceSchema defines the schema for SnippetAuditHistory data source
var SnippetAuditHistoryDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "SnippetAuditHistory data source",
	Attributes: map[string]dsschema.Attribute{
		"action": dsschema.StringAttribute{
			MarkdownDescription: "Action",
			Computed:            true,
		},
		"created": dsschema.StringAttribute{
			MarkdownDescription: "Created",
			Computed:            true,
		},
		"deleted": dsschema.Int64Attribute{
			MarkdownDescription: "Deleted",
			Computed:            true,
		},
		"details": dsschema.StringAttribute{
			MarkdownDescription: "Details",
			Computed:            true,
		},
		"display": dsschema.Int64Attribute{
			MarkdownDescription: "Display",
			Computed:            true,
		},
		"donor_created": dsschema.Int64Attribute{
			MarkdownDescription: "Donor created",
			Computed:            true,
		},
		"donor_tenant_name": dsschema.StringAttribute{
			MarkdownDescription: "Donor tenant name",
			Computed:            true,
		},
		"donor_tsg": dsschema.StringAttribute{
			MarkdownDescription: "Donor tsg",
			Computed:            true,
		},
		"id": dsschema.Int64Attribute{
			MarkdownDescription: "Id",
			Required:            true,
		},
		"recipient_tenant_name": dsschema.StringAttribute{
			MarkdownDescription: "Recipient tenant name",
			Computed:            true,
		},
		"recipient_tsg": dsschema.StringAttribute{
			MarkdownDescription: "Recipient tsg",
			Computed:            true,
		},
		"snippet_uuid": dsschema.StringAttribute{
			MarkdownDescription: "Snippet uuid",
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
		"user": dsschema.StringAttribute{
			MarkdownDescription: "User",
			Computed:            true,
		},
		"version": dsschema.StringAttribute{
			MarkdownDescription: "Version",
			Computed:            true,
		},
	},
}

// SnippetAuditHistoryListModel represents the data model for a list data source.
type SnippetAuditHistoryListModel struct {
	Tfid    types.String          `tfsdk:"tfid"`
	Data    []SnippetAuditHistory `tfsdk:"data"`
	Limit   types.Int64           `tfsdk:"limit"`
	Offset  types.Int64           `tfsdk:"offset"`
	Name    types.String          `tfsdk:"name"`
	Total   types.Int64           `tfsdk:"total"`
	Folder  types.String          `tfsdk:"folder"`
	Snippet types.String          `tfsdk:"snippet"`
	Device  types.String          `tfsdk:"device"`
}

// SnippetAuditHistoryListDataSourceSchema defines the schema for a list data source.
var SnippetAuditHistoryListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: SnippetAuditHistoryDataSourceSchema.Attributes,
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
