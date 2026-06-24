package models

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	dsschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Package: config_setup
// This file contains models for the config_setup SDK package

// SnippetCategories represents the Terraform model for SnippetCategories
type SnippetCategories struct {
	Tfid                          types.String          `tfsdk:"tfid"`
	CreatedIn                     basetypes.StringValue `tfsdk:"created_in"`
	Description                   basetypes.StringValue `tfsdk:"description"`
	DisplayName                   basetypes.StringValue `tfsdk:"display_name"`
	DonorCreated                  basetypes.Int64Value  `tfsdk:"donor_created"`
	DonorSnippetFileId            basetypes.Int64Value  `tfsdk:"donor_snippet_file_id"`
	DonorSnippetVersion           basetypes.Int64Value  `tfsdk:"donor_snippet_version"`
	DonorTenantId                 basetypes.StringValue `tfsdk:"donor_tenant_id"`
	DonorTenantName               basetypes.StringValue `tfsdk:"donor_tenant_name"`
	DonorTsg                      basetypes.StringValue `tfsdk:"donor_tsg"`
	EnablePrefix                  basetypes.BoolValue   `tfsdk:"enable_prefix"`
	Error                         basetypes.StringValue `tfsdk:"error"`
	Folders                       basetypes.ListValue   `tfsdk:"folders"`
	Id                            basetypes.StringValue `tfsdk:"id"`
	Labels                        basetypes.ListValue   `tfsdk:"labels"`
	LastUpdate                    basetypes.StringValue `tfsdk:"last_update"`
	MsgUuid                       basetypes.StringValue `tfsdk:"msg_uuid"`
	Name                          basetypes.StringValue `tfsdk:"name"`
	Prefix                        basetypes.StringValue `tfsdk:"prefix"`
	RecipientPausedUpdate         basetypes.BoolValue   `tfsdk:"recipient_paused_update"`
	RecipientTenantId             basetypes.StringValue `tfsdk:"recipient_tenant_id"`
	RecipientTenantName           basetypes.StringValue `tfsdk:"recipient_tenant_name"`
	RecipientTsg                  basetypes.StringValue `tfsdk:"recipient_tsg"`
	RecipientValidateBeforeUpdate basetypes.BoolValue   `tfsdk:"recipient_validate_before_update"`
	SharedIn                      basetypes.StringValue `tfsdk:"shared_in"`
	SnippetUuid                   basetypes.StringValue `tfsdk:"snippet_uuid"`
	Status                        basetypes.StringValue `tfsdk:"status"`
	Type                          basetypes.StringValue `tfsdk:"type"`
	Version                       basetypes.Int64Value  `tfsdk:"version"`
}

// UsedFolders represents a nested structure within the SnippetCategories model
type UsedFolders struct {
	Id   basetypes.StringValue `tfsdk:"id"`
	Name basetypes.StringValue `tfsdk:"name"`
}

// AttrTypes defines the attribute types for the SnippetCategories model.
func (o SnippetCategories) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":                  basetypes.StringType{},
		"created_in":            basetypes.StringType{},
		"description":           basetypes.StringType{},
		"display_name":          basetypes.StringType{},
		"donor_created":         basetypes.Int64Type{},
		"donor_snippet_file_id": basetypes.Int64Type{},
		"donor_snippet_version": basetypes.Int64Type{},
		"donor_tenant_id":       basetypes.StringType{},
		"donor_tenant_name":     basetypes.StringType{},
		"donor_tsg":             basetypes.StringType{},
		"enable_prefix":         basetypes.BoolType{},
		"error":                 basetypes.StringType{},
		"folders": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"id":   basetypes.StringType{},
				"name": basetypes.StringType{},
			},
		}},
		"id":                               basetypes.StringType{},
		"labels":                           basetypes.ListType{ElemType: basetypes.StringType{}},
		"last_update":                      basetypes.StringType{},
		"msg_uuid":                         basetypes.StringType{},
		"name":                             basetypes.StringType{},
		"prefix":                           basetypes.StringType{},
		"recipient_paused_update":          basetypes.BoolType{},
		"recipient_tenant_id":              basetypes.StringType{},
		"recipient_tenant_name":            basetypes.StringType{},
		"recipient_tsg":                    basetypes.StringType{},
		"recipient_validate_before_update": basetypes.BoolType{},
		"shared_in":                        basetypes.StringType{},
		"snippet_uuid":                     basetypes.StringType{},
		"status":                           basetypes.StringType{},
		"type":                             basetypes.StringType{},
		"version":                          basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of SnippetCategories objects.
func (o SnippetCategories) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the UsedFolders model.
func (o UsedFolders) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":   basetypes.StringType{},
		"name": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of UsedFolders objects.
func (o UsedFolders) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// SnippetCategoriesResourceSchema defines the schema for SnippetCategories resource
var SnippetCategoriesResourceSchema = schema.Schema{
	MarkdownDescription: "SnippetCategory resource",
	Attributes: map[string]schema.Attribute{
		"created_in": schema.StringAttribute{
			MarkdownDescription: "Created in",
			Computed:            true,
		},
		"description": schema.StringAttribute{
			MarkdownDescription: "Description",
			Computed:            true,
		},
		"display_name": schema.StringAttribute{
			MarkdownDescription: "Display name",
			Computed:            true,
		},
		"donor_created": schema.Int64Attribute{
			MarkdownDescription: "Donor created",
			Computed:            true,
		},
		"donor_snippet_file_id": schema.Int64Attribute{
			MarkdownDescription: "Donor snippet file id",
			Computed:            true,
		},
		"donor_snippet_version": schema.Int64Attribute{
			MarkdownDescription: "Donor snippet version",
			Computed:            true,
		},
		"donor_tenant_id": schema.StringAttribute{
			MarkdownDescription: "Donor tenant id",
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
		"enable_prefix": schema.BoolAttribute{
			MarkdownDescription: "Enable prefix",
			Computed:            true,
		},
		"error": schema.StringAttribute{
			MarkdownDescription: "Error",
			Computed:            true,
		},
		"folders": schema.ListNestedAttribute{
			MarkdownDescription: "Folders",
			Optional:            true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						MarkdownDescription: "Id",
						Optional:            true,
					},
					"name": schema.StringAttribute{
						MarkdownDescription: "Name",
						Required:            true,
					},
				},
			},
		},
		"id": schema.StringAttribute{
			MarkdownDescription: "Id",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"labels": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Labels",
			Optional:            true,
		},
		"last_update": schema.StringAttribute{
			MarkdownDescription: "Last update",
			Computed:            true,
		},
		"msg_uuid": schema.StringAttribute{
			MarkdownDescription: "Msg uuid",
			Computed:            true,
		},
		"name": schema.StringAttribute{
			MarkdownDescription: "Name",
			Computed:            true,
		},
		"prefix": schema.StringAttribute{
			MarkdownDescription: "Prefix",
			Computed:            true,
		},
		"recipient_paused_update": schema.BoolAttribute{
			MarkdownDescription: "Recipient paused update",
			Computed:            true,
		},
		"recipient_tenant_id": schema.StringAttribute{
			MarkdownDescription: "Recipient tenant id",
			Computed:            true,
		},
		"recipient_tenant_name": schema.StringAttribute{
			MarkdownDescription: "Recipient tenant name",
			Computed:            true,
		},
		"recipient_tsg": schema.StringAttribute{
			MarkdownDescription: "Recipient tsg",
			Computed:            true,
		},
		"recipient_validate_before_update": schema.BoolAttribute{
			MarkdownDescription: "Recipient validate before update",
			Computed:            true,
		},
		"shared_in": schema.StringAttribute{
			MarkdownDescription: "Shared in",
			Computed:            true,
		},
		"snippet_uuid": schema.StringAttribute{
			MarkdownDescription: "Snippet uuid",
			Computed:            true,
		},
		"status": schema.StringAttribute{
			MarkdownDescription: "Status",
			Computed:            true,
		},
		"tfid": schema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"type": schema.StringAttribute{
			MarkdownDescription: "Type",
			Computed:            true,
		},
		"version": schema.Int64Attribute{
			MarkdownDescription: "Version",
			Computed:            true,
		},
	},
}

// SnippetCategoriesDataSourceSchema defines the schema for SnippetCategories data source
var SnippetCategoriesDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "SnippetCategory data source",
	Attributes: map[string]dsschema.Attribute{
		"created_in": dsschema.StringAttribute{
			MarkdownDescription: "Created in",
			Computed:            true,
		},
		"description": dsschema.StringAttribute{
			MarkdownDescription: "Description",
			Computed:            true,
		},
		"display_name": dsschema.StringAttribute{
			MarkdownDescription: "Display name",
			Computed:            true,
		},
		"donor_created": dsschema.Int64Attribute{
			MarkdownDescription: "Donor created",
			Computed:            true,
		},
		"donor_snippet_file_id": dsschema.Int64Attribute{
			MarkdownDescription: "Donor snippet file id",
			Computed:            true,
		},
		"donor_snippet_version": dsschema.Int64Attribute{
			MarkdownDescription: "Donor snippet version",
			Computed:            true,
		},
		"donor_tenant_id": dsschema.StringAttribute{
			MarkdownDescription: "Donor tenant id",
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
		"enable_prefix": dsschema.BoolAttribute{
			MarkdownDescription: "Enable prefix",
			Computed:            true,
		},
		"error": dsschema.StringAttribute{
			MarkdownDescription: "Error",
			Computed:            true,
		},
		"folders": dsschema.ListNestedAttribute{
			MarkdownDescription: "Folders",
			Computed:            true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: map[string]dsschema.Attribute{
					"id": dsschema.StringAttribute{
						MarkdownDescription: "Id",
						Computed:            true,
					},
					"name": dsschema.StringAttribute{
						MarkdownDescription: "Name",
						Computed:            true,
					},
				},
			},
		},
		"id": dsschema.StringAttribute{
			MarkdownDescription: "Id",
			Required:            true,
		},
		"labels": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Labels",
			Computed:            true,
		},
		"last_update": dsschema.StringAttribute{
			MarkdownDescription: "Last update",
			Computed:            true,
		},
		"msg_uuid": dsschema.StringAttribute{
			MarkdownDescription: "Msg uuid",
			Computed:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "Name",
			Optional:            true,
			Computed:            true,
		},
		"prefix": dsschema.StringAttribute{
			MarkdownDescription: "Prefix",
			Computed:            true,
		},
		"recipient_paused_update": dsschema.BoolAttribute{
			MarkdownDescription: "Recipient paused update",
			Computed:            true,
		},
		"recipient_tenant_id": dsschema.StringAttribute{
			MarkdownDescription: "Recipient tenant id",
			Computed:            true,
		},
		"recipient_tenant_name": dsschema.StringAttribute{
			MarkdownDescription: "Recipient tenant name",
			Computed:            true,
		},
		"recipient_tsg": dsschema.StringAttribute{
			MarkdownDescription: "Recipient tsg",
			Computed:            true,
		},
		"recipient_validate_before_update": dsschema.BoolAttribute{
			MarkdownDescription: "Recipient validate before update",
			Computed:            true,
		},
		"shared_in": dsschema.StringAttribute{
			MarkdownDescription: "Shared in",
			Computed:            true,
		},
		"snippet_uuid": dsschema.StringAttribute{
			MarkdownDescription: "Snippet uuid",
			Computed:            true,
		},
		"status": dsschema.StringAttribute{
			MarkdownDescription: "Status",
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
		"type": dsschema.StringAttribute{
			MarkdownDescription: "Type",
			Computed:            true,
		},
		"version": dsschema.Int64Attribute{
			MarkdownDescription: "Version",
			Computed:            true,
		},
	},
}

// SnippetCategoriesListModel represents the data model for a list data source.
type SnippetCategoriesListModel struct {
	Tfid    types.String        `tfsdk:"tfid"`
	Data    []SnippetCategories `tfsdk:"data"`
	Limit   types.Int64         `tfsdk:"limit"`
	Offset  types.Int64         `tfsdk:"offset"`
	Name    types.String        `tfsdk:"name"`
	Total   types.Int64         `tfsdk:"total"`
	Folder  types.String        `tfsdk:"folder"`
	Snippet types.String        `tfsdk:"snippet"`
	Device  types.String        `tfsdk:"device"`
}

// SnippetCategoriesListDataSourceSchema defines the schema for a list data source.
var SnippetCategoriesListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: SnippetCategoriesDataSourceSchema.Attributes,
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
