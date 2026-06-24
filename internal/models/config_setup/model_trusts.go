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

// TenantTrustInfo represents the Terraform model for TenantTrustInfo
type TenantTrustInfo struct {
	Tfid                 types.String          `tfsdk:"tfid"`
	Created              basetypes.StringValue `tfsdk:"created"`
	CreatedBy            basetypes.StringValue `tfsdk:"created_by"`
	CurrentStatus        basetypes.StringValue `tfsdk:"current_status"`
	DonorCluster         basetypes.StringValue `tfsdk:"donor_cluster"`
	DonorMsgUuid         basetypes.StringValue `tfsdk:"donor_msg_uuid"`
	DonorProject         basetypes.StringValue `tfsdk:"donor_project"`
	DonorRegion          basetypes.StringValue `tfsdk:"donor_region"`
	DonorTenantId        basetypes.StringValue `tfsdk:"donor_tenant_id"`
	DonorTenantName      basetypes.StringValue `tfsdk:"donor_tenant_name"`
	DonorTrustInfoId     basetypes.Int64Value  `tfsdk:"donor_trust_info_id"`
	DonorTsg             basetypes.StringValue `tfsdk:"donor_tsg"`
	ErrorDetails         basetypes.StringValue `tfsdk:"error_details"`
	LastUpdated          basetypes.StringValue `tfsdk:"last_updated"`
	Psk                  basetypes.StringValue `tfsdk:"psk"`
	RecipientCluster     basetypes.StringValue `tfsdk:"recipient_cluster"`
	RecipientMsgUuid     basetypes.StringValue `tfsdk:"recipient_msg_uuid"`
	RecipientProject     basetypes.StringValue `tfsdk:"recipient_project"`
	RecipientRegion      basetypes.StringValue `tfsdk:"recipient_region"`
	RecipientTenantId    basetypes.StringValue `tfsdk:"recipient_tenant_id"`
	RecipientTenantName  basetypes.StringValue `tfsdk:"recipient_tenant_name"`
	RecipientTrustInfoId basetypes.Int64Value  `tfsdk:"recipient_trust_info_id"`
	RecipientTsg         basetypes.StringValue `tfsdk:"recipient_tsg"`
	TrustId              basetypes.Int64Value  `tfsdk:"trust_id"`
	UpdatedBy            basetypes.StringValue `tfsdk:"updated_by"`
}

// Trusts represents a nested structure within the TenantTrustInfo model
type Trusts struct {
	Tfid                types.String          `tfsdk:"tfid"`
	DonorTenantName     basetypes.StringValue `tfsdk:"donor_tenant_name"`
	Psk                 basetypes.StringValue `tfsdk:"psk"`
	RecipientTenantName basetypes.StringValue `tfsdk:"recipient_tenant_name"`
	TrustId             basetypes.Int64Value  `tfsdk:"trust_id"`
	Tsg                 basetypes.StringValue `tfsdk:"tsg"`
}

// AttrTypes defines the attribute types for the TenantTrustInfo model.
func (o TenantTrustInfo) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":                    basetypes.StringType{},
		"created":                 basetypes.StringType{},
		"created_by":              basetypes.StringType{},
		"current_status":          basetypes.StringType{},
		"donor_cluster":           basetypes.StringType{},
		"donor_msg_uuid":          basetypes.StringType{},
		"donor_project":           basetypes.StringType{},
		"donor_region":            basetypes.StringType{},
		"donor_tenant_id":         basetypes.StringType{},
		"donor_tenant_name":       basetypes.StringType{},
		"donor_trust_info_id":     basetypes.Int64Type{},
		"donor_tsg":               basetypes.StringType{},
		"error_details":           basetypes.StringType{},
		"last_updated":            basetypes.StringType{},
		"psk":                     basetypes.StringType{},
		"recipient_cluster":       basetypes.StringType{},
		"recipient_msg_uuid":      basetypes.StringType{},
		"recipient_project":       basetypes.StringType{},
		"recipient_region":        basetypes.StringType{},
		"recipient_tenant_id":     basetypes.StringType{},
		"recipient_tenant_name":   basetypes.StringType{},
		"recipient_trust_info_id": basetypes.Int64Type{},
		"recipient_tsg":           basetypes.StringType{},
		"trust_id":                basetypes.Int64Type{},
		"updated_by":              basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of TenantTrustInfo objects.
func (o TenantTrustInfo) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the Trusts model.
func (o Trusts) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":                  basetypes.StringType{},
		"donor_tenant_name":     basetypes.StringType{},
		"psk":                   basetypes.StringType{},
		"recipient_tenant_name": basetypes.StringType{},
		"trust_id":              basetypes.Int64Type{},
		"tsg":                   basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of Trusts objects.
func (o Trusts) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// TenantTrustInfoResourceSchema defines the schema for TenantTrustInfo resource
var TenantTrustInfoResourceSchema = schema.Schema{
	MarkdownDescription: "TenantTrustInfo resource",
	Attributes: map[string]schema.Attribute{
		"created": schema.StringAttribute{
			MarkdownDescription: "Created",
			Computed:            true,
		},
		"created_by": schema.StringAttribute{
			MarkdownDescription: "Created by",
			Computed:            true,
		},
		"current_status": schema.StringAttribute{
			MarkdownDescription: "Current status",
			Computed:            true,
		},
		"donor_cluster": schema.StringAttribute{
			MarkdownDescription: "Donor cluster",
			Computed:            true,
		},
		"donor_msg_uuid": schema.StringAttribute{
			MarkdownDescription: "Donor msg uuid",
			Computed:            true,
		},
		"donor_project": schema.StringAttribute{
			MarkdownDescription: "Donor project",
			Computed:            true,
		},
		"donor_region": schema.StringAttribute{
			MarkdownDescription: "Donor region",
			Computed:            true,
		},
		"donor_tenant_id": schema.StringAttribute{
			MarkdownDescription: "Donor tenant id",
			Optional:            true,
		},
		"donor_tenant_name": schema.StringAttribute{
			MarkdownDescription: "Donor tenant name",
			Optional:            true,
		},
		"donor_trust_info_id": schema.Int64Attribute{
			MarkdownDescription: "Donor trust info id",
			Computed:            true,
		},
		"donor_tsg": schema.StringAttribute{
			MarkdownDescription: "Donor tsg",
			Computed:            true,
		},
		"error_details": schema.StringAttribute{
			MarkdownDescription: "Error details",
			Computed:            true,
		},
		"last_updated": schema.StringAttribute{
			MarkdownDescription: "Last updated",
			Computed:            true,
		},
		"psk": schema.StringAttribute{
			MarkdownDescription: "Psk",
			Optional:            true,
		},
		"recipient_cluster": schema.StringAttribute{
			MarkdownDescription: "Recipient cluster",
			Computed:            true,
		},
		"recipient_msg_uuid": schema.StringAttribute{
			MarkdownDescription: "Recipient msg uuid",
			Computed:            true,
		},
		"recipient_project": schema.StringAttribute{
			MarkdownDescription: "Recipient project",
			Computed:            true,
		},
		"recipient_region": schema.StringAttribute{
			MarkdownDescription: "Recipient region",
			Computed:            true,
		},
		"recipient_tenant_id": schema.StringAttribute{
			MarkdownDescription: "Recipient tenant id",
			Computed:            true,
		},
		"recipient_tenant_name": schema.StringAttribute{
			MarkdownDescription: "Recipient tenant name",
			Optional:            true,
		},
		"recipient_trust_info_id": schema.Int64Attribute{
			MarkdownDescription: "Recipient trust info id",
			Computed:            true,
		},
		"recipient_tsg": schema.StringAttribute{
			MarkdownDescription: "Recipient tsg",
			Computed:            true,
		},
		"tfid": schema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"trust_id": schema.Int64Attribute{
			MarkdownDescription: "Trust id",
			Optional:            true,
		},
		"updated_by": schema.StringAttribute{
			MarkdownDescription: "Updated by",
			Computed:            true,
		},
	},
}

// TenantTrustInfoDataSourceSchema defines the schema for TenantTrustInfo data source
var TenantTrustInfoDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "TenantTrustInfo data source",
	Attributes: map[string]dsschema.Attribute{
		"created": dsschema.StringAttribute{
			MarkdownDescription: "Created",
			Computed:            true,
		},
		"created_by": dsschema.StringAttribute{
			MarkdownDescription: "Created by",
			Computed:            true,
		},
		"current_status": dsschema.StringAttribute{
			MarkdownDescription: "Current status",
			Computed:            true,
		},
		"donor_cluster": dsschema.StringAttribute{
			MarkdownDescription: "Donor cluster",
			Computed:            true,
		},
		"donor_msg_uuid": dsschema.StringAttribute{
			MarkdownDescription: "Donor msg uuid",
			Computed:            true,
		},
		"donor_project": dsschema.StringAttribute{
			MarkdownDescription: "Donor project",
			Computed:            true,
		},
		"donor_region": dsschema.StringAttribute{
			MarkdownDescription: "Donor region",
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
		"donor_trust_info_id": dsschema.Int64Attribute{
			MarkdownDescription: "Donor trust info id",
			Computed:            true,
		},
		"donor_tsg": dsschema.StringAttribute{
			MarkdownDescription: "Donor tsg",
			Computed:            true,
		},
		"error_details": dsschema.StringAttribute{
			MarkdownDescription: "Error details",
			Computed:            true,
		},
		"last_updated": dsschema.StringAttribute{
			MarkdownDescription: "Last updated",
			Computed:            true,
		},
		"psk": dsschema.StringAttribute{
			MarkdownDescription: "Psk",
			Computed:            true,
		},
		"recipient_cluster": dsschema.StringAttribute{
			MarkdownDescription: "Recipient cluster",
			Computed:            true,
		},
		"recipient_msg_uuid": dsschema.StringAttribute{
			MarkdownDescription: "Recipient msg uuid",
			Computed:            true,
		},
		"recipient_project": dsschema.StringAttribute{
			MarkdownDescription: "Recipient project",
			Computed:            true,
		},
		"recipient_region": dsschema.StringAttribute{
			MarkdownDescription: "Recipient region",
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
		"recipient_trust_info_id": dsschema.Int64Attribute{
			MarkdownDescription: "Recipient trust info id",
			Computed:            true,
		},
		"recipient_tsg": dsschema.StringAttribute{
			MarkdownDescription: "Recipient tsg",
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
		"trust_id": dsschema.Int64Attribute{
			MarkdownDescription: "Trust id",
			Computed:            true,
		},
		"updated_by": dsschema.StringAttribute{
			MarkdownDescription: "Updated by",
			Computed:            true,
		},
	},
}

// TenantTrustInfoListModel represents the data model for a list data source.
type TenantTrustInfoListModel struct {
	Tfid    types.String      `tfsdk:"tfid"`
	Data    []TenantTrustInfo `tfsdk:"data"`
	Limit   types.Int64       `tfsdk:"limit"`
	Offset  types.Int64       `tfsdk:"offset"`
	Name    types.String      `tfsdk:"name"`
	Total   types.Int64       `tfsdk:"total"`
	Folder  types.String      `tfsdk:"folder"`
	Snippet types.String      `tfsdk:"snippet"`
	Device  types.String      `tfsdk:"device"`
}

// TenantTrustInfoListDataSourceSchema defines the schema for a list data source.
var TenantTrustInfoListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: TenantTrustInfoDataSourceSchema.Attributes,
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
