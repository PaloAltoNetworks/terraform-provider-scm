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

// TrustInfoWithSharedSnippets represents the Terraform model for TrustInfoWithSharedSnippets
type TrustInfoWithSharedSnippets struct {
	Tfid                          types.String          `tfsdk:"tfid"`
	Created                       basetypes.StringValue `tfsdk:"created"`
	DonorCreated                  basetypes.Int64Value  `tfsdk:"donor_created"`
	DonorSnippetFileId            basetypes.Int64Value  `tfsdk:"donor_snippet_file_id"`
	DonorSnippetVersion           basetypes.Int64Value  `tfsdk:"donor_snippet_version"`
	DonorTsg                      basetypes.StringValue `tfsdk:"donor_tsg"`
	Error                         basetypes.StringValue `tfsdk:"error"`
	Id                            basetypes.Int64Value  `tfsdk:"id"`
	LastUpdated                   basetypes.StringValue `tfsdk:"last_updated"`
	MsgUuid                       basetypes.StringValue `tfsdk:"msg_uuid"`
	RecipientPausedUpdate         basetypes.Int64Value  `tfsdk:"recipient_paused_update"`
	RecipientSnippetFileId        basetypes.Int64Value  `tfsdk:"recipient_snippet_file_id"`
	RecipientSnippetVersion       basetypes.Int64Value  `tfsdk:"recipient_snippet_version"`
	RecipientTsg                  basetypes.StringValue `tfsdk:"recipient_tsg"`
	RecipientValidateBeforeUpdate basetypes.Int64Value  `tfsdk:"recipient_validate_before_update"`
	SharedSnippets                basetypes.ListValue   `tfsdk:"shared_snippets"`
	SnippetName                   basetypes.StringValue `tfsdk:"snippet_name"`
	SnippetUuid                   basetypes.StringValue `tfsdk:"snippet_uuid"`
	Status                        basetypes.StringValue `tfsdk:"status"`
	UpdatedBy                     basetypes.StringValue `tfsdk:"updated_by"`
}

// AttrTypes defines the attribute types for the TrustInfoWithSharedSnippets model.
func (o TrustInfoWithSharedSnippets) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":                             basetypes.StringType{},
		"created":                          basetypes.StringType{},
		"donor_created":                    basetypes.Int64Type{},
		"donor_snippet_file_id":            basetypes.Int64Type{},
		"donor_snippet_version":            basetypes.Int64Type{},
		"donor_tsg":                        basetypes.StringType{},
		"error":                            basetypes.StringType{},
		"id":                               basetypes.Int64Type{},
		"last_updated":                     basetypes.StringType{},
		"msg_uuid":                         basetypes.StringType{},
		"recipient_paused_update":          basetypes.Int64Type{},
		"recipient_snippet_file_id":        basetypes.Int64Type{},
		"recipient_snippet_version":        basetypes.Int64Type{},
		"recipient_tsg":                    basetypes.StringType{},
		"recipient_validate_before_update": basetypes.Int64Type{},
		"shared_snippets": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"created":               basetypes.StringType{},
				"donor_created":         basetypes.Int64Type{},
				"donor_snippet_file_id": basetypes.Int64Type{},
				"donor_snippet_version": basetypes.Int64Type{},
				"donor_tenant_id":       basetypes.StringType{},
				"donor_tenant_name":     basetypes.StringType{},
				"donor_tsg":             basetypes.StringType{},
				"error":                 basetypes.StringType{},
				"id":                    basetypes.Int64Type{},
				"last_updated":          basetypes.StringType{},
				"msg_uuid":              basetypes.StringType{},
				"properties": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"created":          basetypes.StringType{},
						"created_by":       basetypes.StringType{},
						"donor_tenant":     basetypes.StringType{},
						"donor_tsg":        basetypes.StringType{},
						"error":            basetypes.StringType{},
						"id":               basetypes.Int64Type{},
						"msg_uuid":         basetypes.StringType{},
						"property_name":    basetypes.StringType{},
						"property_value":   basetypes.StringType{},
						"recipient_tenant": basetypes.StringType{},
						"recipient_tsg":    basetypes.StringType{},
						"snippet_name":     basetypes.StringType{},
						"snippet_uuid":     basetypes.StringType{},
						"status":           basetypes.StringType{},
						"updated":          basetypes.StringType{},
						"updated_by":       basetypes.StringType{},
					},
				}},
				"recipient_paused_update":          basetypes.BoolType{},
				"recipient_snippet_file_id":        basetypes.Int64Type{},
				"recipient_snippet_version":        basetypes.Int64Type{},
				"recipient_tenant_id":              basetypes.StringType{},
				"recipient_tenant_name":            basetypes.StringType{},
				"recipient_tsg":                    basetypes.StringType{},
				"recipient_validate_before_update": basetypes.BoolType{},
				"snippet_name":                     basetypes.StringType{},
				"snippet_uuid":                     basetypes.StringType{},
				"status":                           basetypes.StringType{},
			},
		}},
		"snippet_name": basetypes.StringType{},
		"snippet_uuid": basetypes.StringType{},
		"status":       basetypes.StringType{},
		"updated_by":   basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of TrustInfoWithSharedSnippets objects.
func (o TrustInfoWithSharedSnippets) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// TrustInfoWithSharedSnippetsResourceSchema defines the schema for TrustInfoWithSharedSnippets resource
var TrustInfoWithSharedSnippetsResourceSchema = schema.Schema{
	MarkdownDescription: "TrustInfoWithSharedSnippet resource",
	Attributes: map[string]schema.Attribute{
		"created": schema.StringAttribute{
			MarkdownDescription: "Created",
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
		"donor_tsg": schema.StringAttribute{
			MarkdownDescription: "Donor tsg",
			Computed:            true,
		},
		"error": schema.StringAttribute{
			MarkdownDescription: "Error",
			Computed:            true,
		},
		"id": schema.Int64Attribute{
			MarkdownDescription: "Id",
			Computed:            true,
			PlanModifiers: []planmodifier.Int64{
				int64planmodifier.UseStateForUnknown(),
			},
		},
		"last_updated": schema.StringAttribute{
			MarkdownDescription: "Last updated",
			Computed:            true,
		},
		"msg_uuid": schema.StringAttribute{
			MarkdownDescription: "Msg uuid",
			Computed:            true,
		},
		"recipient_paused_update": schema.Int64Attribute{
			MarkdownDescription: "Recipient paused update",
			Computed:            true,
		},
		"recipient_snippet_file_id": schema.Int64Attribute{
			MarkdownDescription: "Recipient snippet file id",
			Computed:            true,
		},
		"recipient_snippet_version": schema.Int64Attribute{
			MarkdownDescription: "Recipient snippet version",
			Computed:            true,
		},
		"recipient_tsg": schema.StringAttribute{
			MarkdownDescription: "Recipient tsg",
			Computed:            true,
		},
		"recipient_validate_before_update": schema.Int64Attribute{
			MarkdownDescription: "Recipient validate before update",
			Computed:            true,
		},
		"shared_snippets": schema.ListNestedAttribute{
			MarkdownDescription: "Shared snippets",
			Optional:            true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"created": schema.StringAttribute{
						MarkdownDescription: "Created",
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
					"error": schema.StringAttribute{
						MarkdownDescription: "Error",
						Computed:            true,
					},
					"id": schema.Int64Attribute{
						MarkdownDescription: "Id",
						Computed:            true,
					},
					"last_updated": schema.StringAttribute{
						MarkdownDescription: "Last updated",
						Computed:            true,
					},
					"msg_uuid": schema.StringAttribute{
						MarkdownDescription: "Msg uuid",
						Computed:            true,
					},
					"properties": schema.ListNestedAttribute{
						MarkdownDescription: "Properties",
						Optional:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"created": schema.StringAttribute{
									MarkdownDescription: "Created",
									Computed:            true,
								},
								"created_by": schema.StringAttribute{
									MarkdownDescription: "Created by",
									Computed:            true,
								},
								"donor_tenant": schema.StringAttribute{
									MarkdownDescription: "Donor tenant",
									Computed:            true,
								},
								"donor_tsg": schema.StringAttribute{
									MarkdownDescription: "Donor tsg",
									Computed:            true,
								},
								"error": schema.StringAttribute{
									MarkdownDescription: "Error",
									Computed:            true,
								},
								"id": schema.Int64Attribute{
									MarkdownDescription: "Id",
									Computed:            true,
								},
								"msg_uuid": schema.StringAttribute{
									MarkdownDescription: "Msg uuid",
									Computed:            true,
								},
								"property_name": schema.StringAttribute{
									MarkdownDescription: "Property name",
									Computed:            true,
								},
								"property_value": schema.StringAttribute{
									MarkdownDescription: "Property value",
									Computed:            true,
								},
								"recipient_tenant": schema.StringAttribute{
									MarkdownDescription: "Recipient tenant",
									Computed:            true,
								},
								"recipient_tsg": schema.StringAttribute{
									MarkdownDescription: "Recipient tsg",
									Computed:            true,
								},
								"snippet_name": schema.StringAttribute{
									MarkdownDescription: "Snippet name",
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
								"updated": schema.StringAttribute{
									MarkdownDescription: "Updated",
									Computed:            true,
								},
								"updated_by": schema.StringAttribute{
									MarkdownDescription: "Updated by",
									Computed:            true,
								},
							},
						},
					},
					"recipient_paused_update": schema.BoolAttribute{
						MarkdownDescription: "Recipient paused update",
						Computed:            true,
					},
					"recipient_snippet_file_id": schema.Int64Attribute{
						MarkdownDescription: "Recipient snippet file id",
						Computed:            true,
					},
					"recipient_snippet_version": schema.Int64Attribute{
						MarkdownDescription: "Recipient snippet version",
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
					"snippet_name": schema.StringAttribute{
						MarkdownDescription: "Snippet name",
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
				},
			},
		},
		"snippet_name": schema.StringAttribute{
			MarkdownDescription: "Snippet name",
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
		"updated_by": schema.StringAttribute{
			MarkdownDescription: "Updated by",
			Computed:            true,
		},
	},
}

// TrustInfoWithSharedSnippetsDataSourceSchema defines the schema for TrustInfoWithSharedSnippets data source
var TrustInfoWithSharedSnippetsDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "TrustInfoWithSharedSnippet data source",
	Attributes: map[string]dsschema.Attribute{
		"created": dsschema.StringAttribute{
			MarkdownDescription: "Created",
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
		"donor_tsg": dsschema.StringAttribute{
			MarkdownDescription: "Donor tsg",
			Computed:            true,
		},
		"error": dsschema.StringAttribute{
			MarkdownDescription: "Error",
			Computed:            true,
		},
		"id": dsschema.Int64Attribute{
			MarkdownDescription: "Id",
			Required:            true,
		},
		"last_updated": dsschema.StringAttribute{
			MarkdownDescription: "Last updated",
			Computed:            true,
		},
		"msg_uuid": dsschema.StringAttribute{
			MarkdownDescription: "Msg uuid",
			Computed:            true,
		},
		"recipient_paused_update": dsschema.Int64Attribute{
			MarkdownDescription: "Recipient paused update",
			Computed:            true,
		},
		"recipient_snippet_file_id": dsschema.Int64Attribute{
			MarkdownDescription: "Recipient snippet file id",
			Computed:            true,
		},
		"recipient_snippet_version": dsschema.Int64Attribute{
			MarkdownDescription: "Recipient snippet version",
			Computed:            true,
		},
		"recipient_tsg": dsschema.StringAttribute{
			MarkdownDescription: "Recipient tsg",
			Computed:            true,
		},
		"recipient_validate_before_update": dsschema.Int64Attribute{
			MarkdownDescription: "Recipient validate before update",
			Computed:            true,
		},
		"shared_snippets": dsschema.ListNestedAttribute{
			MarkdownDescription: "Shared snippets",
			Computed:            true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: map[string]dsschema.Attribute{
					"created": dsschema.StringAttribute{
						MarkdownDescription: "Created",
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
					"error": dsschema.StringAttribute{
						MarkdownDescription: "Error",
						Computed:            true,
					},
					"id": dsschema.Int64Attribute{
						MarkdownDescription: "Id",
						Computed:            true,
					},
					"last_updated": dsschema.StringAttribute{
						MarkdownDescription: "Last updated",
						Computed:            true,
					},
					"msg_uuid": dsschema.StringAttribute{
						MarkdownDescription: "Msg uuid",
						Computed:            true,
					},
					"properties": dsschema.ListNestedAttribute{
						MarkdownDescription: "Properties",
						Computed:            true,
						NestedObject: dsschema.NestedAttributeObject{
							Attributes: map[string]dsschema.Attribute{
								"created": dsschema.StringAttribute{
									MarkdownDescription: "Created",
									Computed:            true,
								},
								"created_by": dsschema.StringAttribute{
									MarkdownDescription: "Created by",
									Computed:            true,
								},
								"donor_tenant": dsschema.StringAttribute{
									MarkdownDescription: "Donor tenant",
									Computed:            true,
								},
								"donor_tsg": dsschema.StringAttribute{
									MarkdownDescription: "Donor tsg",
									Computed:            true,
								},
								"error": dsschema.StringAttribute{
									MarkdownDescription: "Error",
									Computed:            true,
								},
								"id": dsschema.Int64Attribute{
									MarkdownDescription: "Id",
									Computed:            true,
								},
								"msg_uuid": dsschema.StringAttribute{
									MarkdownDescription: "Msg uuid",
									Computed:            true,
								},
								"property_name": dsschema.StringAttribute{
									MarkdownDescription: "Property name",
									Computed:            true,
								},
								"property_value": dsschema.StringAttribute{
									MarkdownDescription: "Property value",
									Computed:            true,
								},
								"recipient_tenant": dsschema.StringAttribute{
									MarkdownDescription: "Recipient tenant",
									Computed:            true,
								},
								"recipient_tsg": dsschema.StringAttribute{
									MarkdownDescription: "Recipient tsg",
									Computed:            true,
								},
								"snippet_name": dsschema.StringAttribute{
									MarkdownDescription: "Snippet name",
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
								"updated": dsschema.StringAttribute{
									MarkdownDescription: "Updated",
									Computed:            true,
								},
								"updated_by": dsschema.StringAttribute{
									MarkdownDescription: "Updated by",
									Computed:            true,
								},
							},
						},
					},
					"recipient_paused_update": dsschema.BoolAttribute{
						MarkdownDescription: "Recipient paused update",
						Computed:            true,
					},
					"recipient_snippet_file_id": dsschema.Int64Attribute{
						MarkdownDescription: "Recipient snippet file id",
						Computed:            true,
					},
					"recipient_snippet_version": dsschema.Int64Attribute{
						MarkdownDescription: "Recipient snippet version",
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
					"snippet_name": dsschema.StringAttribute{
						MarkdownDescription: "Snippet name",
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
				},
			},
		},
		"snippet_name": dsschema.StringAttribute{
			MarkdownDescription: "Snippet name",
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
		"updated_by": dsschema.StringAttribute{
			MarkdownDescription: "Updated by",
			Computed:            true,
		},
	},
}

// TrustInfoWithSharedSnippetsListModel represents the data model for a list data source.
type TrustInfoWithSharedSnippetsListModel struct {
	Tfid    types.String                  `tfsdk:"tfid"`
	Data    []TrustInfoWithSharedSnippets `tfsdk:"data"`
	Limit   types.Int64                   `tfsdk:"limit"`
	Offset  types.Int64                   `tfsdk:"offset"`
	Name    types.String                  `tfsdk:"name"`
	Total   types.Int64                   `tfsdk:"total"`
	Folder  types.String                  `tfsdk:"folder"`
	Snippet types.String                  `tfsdk:"snippet"`
	Device  types.String                  `tfsdk:"device"`
}

// TrustInfoWithSharedSnippetsListDataSourceSchema defines the schema for a list data source.
var TrustInfoWithSharedSnippetsListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: TrustInfoWithSharedSnippetsDataSourceSchema.Attributes,
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
