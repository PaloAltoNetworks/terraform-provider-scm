package models

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
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

// Package: identity_services
// This file contains models for the identity_services SDK package

// AuthenticationPortals represents the Terraform model for AuthenticationPortals
type AuthenticationPortals struct {
	Tfid                  types.String          `tfsdk:"tfid"`
	AuthenticationProfile basetypes.StringValue `tfsdk:"authentication_profile"`
	CertificateProfile    basetypes.StringValue `tfsdk:"certificate_profile"`
	Device                basetypes.StringValue `tfsdk:"device"`
	Folder                basetypes.StringValue `tfsdk:"folder"`
	GpUdpPort             basetypes.Int64Value  `tfsdk:"gp_udp_port"`
	Id                    basetypes.StringValue `tfsdk:"id"`
	IdleTimer             basetypes.Int64Value  `tfsdk:"idle_timer"`
	RedirectHost          basetypes.StringValue `tfsdk:"redirect_host"`
	Snippet               basetypes.StringValue `tfsdk:"snippet"`
	Timer                 basetypes.Int64Value  `tfsdk:"timer"`
	TlsServiceProfile     basetypes.StringValue `tfsdk:"tls_service_profile"`
}

// AttrTypes defines the attribute types for the AuthenticationPortals model.
func (o AuthenticationPortals) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":                   basetypes.StringType{},
		"authentication_profile": basetypes.StringType{},
		"certificate_profile":    basetypes.StringType{},
		"device":                 basetypes.StringType{},
		"folder":                 basetypes.StringType{},
		"gp_udp_port":            basetypes.Int64Type{},
		"id":                     basetypes.StringType{},
		"idle_timer":             basetypes.Int64Type{},
		"redirect_host":          basetypes.StringType{},
		"snippet":                basetypes.StringType{},
		"timer":                  basetypes.Int64Type{},
		"tls_service_profile":    basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of AuthenticationPortals objects.
func (o AuthenticationPortals) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AuthenticationPortalsResourceSchema defines the schema for AuthenticationPortals resource
var AuthenticationPortalsResourceSchema = schema.Schema{
	MarkdownDescription: "AuthenticationPortal resource",
	Attributes: map[string]schema.Attribute{
		"authentication_profile": schema.StringAttribute{
			MarkdownDescription: "The authentication profile",
			Optional:            true,
		},
		"certificate_profile": schema.StringAttribute{
			MarkdownDescription: "The certificate profile",
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
		"gp_udp_port": schema.Int64Attribute{
			Validators: []validator.Int64{
				int64validator.Between(1, 65535),
			},
			MarkdownDescription: "The UDP port for inbound authentication prompts",
			Optional:            true,
		},
		"id": schema.StringAttribute{
			MarkdownDescription: "The UUID of the authentication portal",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"idle_timer": schema.Int64Attribute{
			Validators: []validator.Int64{
				int64validator.Between(1, 1440),
			},
			MarkdownDescription: "The idle timeout value (minutes)",
			Optional:            true,
		},
		"redirect_host": schema.StringAttribute{
			MarkdownDescription: "The authentication portal IP address or hostname",
			Required:            true,
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
		"timer": schema.Int64Attribute{
			Validators: []validator.Int64{
				int64validator.Between(1, 1440),
			},
			MarkdownDescription: "Timer",
			Optional:            true,
		},
		"tls_service_profile": schema.StringAttribute{
			MarkdownDescription: "The SSL/TLS service profile",
			Optional:            true,
		},
	},
}

// AuthenticationPortalsDataSourceSchema defines the schema for AuthenticationPortals data source
var AuthenticationPortalsDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "AuthenticationPortal data source",
	Attributes: map[string]dsschema.Attribute{
		"authentication_profile": dsschema.StringAttribute{
			MarkdownDescription: "The authentication profile",
			Computed:            true,
		},
		"certificate_profile": dsschema.StringAttribute{
			MarkdownDescription: "The certificate profile",
			Computed:            true,
		},
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
		"gp_udp_port": dsschema.Int64Attribute{
			MarkdownDescription: "The UDP port for inbound authentication prompts",
			Computed:            true,
		},
		"id": dsschema.StringAttribute{
			MarkdownDescription: "The UUID of the authentication portal",
			Required:            true,
		},
		"idle_timer": dsschema.Int64Attribute{
			MarkdownDescription: "The idle timeout value (minutes)",
			Computed:            true,
		},
		"redirect_host": dsschema.StringAttribute{
			MarkdownDescription: "The authentication portal IP address or hostname",
			Computed:            true,
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
		"timer": dsschema.Int64Attribute{
			MarkdownDescription: "Timer",
			Computed:            true,
		},
		"tls_service_profile": dsschema.StringAttribute{
			MarkdownDescription: "The SSL/TLS service profile",
			Computed:            true,
		},
	},
}

// AuthenticationPortalsListModel represents the data model for a list data source.
type AuthenticationPortalsListModel struct {
	Tfid    types.String            `tfsdk:"tfid"`
	Data    []AuthenticationPortals `tfsdk:"data"`
	Limit   types.Int64             `tfsdk:"limit"`
	Offset  types.Int64             `tfsdk:"offset"`
	Name    types.String            `tfsdk:"name"`
	Total   types.Int64             `tfsdk:"total"`
	Folder  types.String            `tfsdk:"folder"`
	Snippet types.String            `tfsdk:"snippet"`
	Device  types.String            `tfsdk:"device"`
}

// AuthenticationPortalsListDataSourceSchema defines the schema for a list data source.
var AuthenticationPortalsListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: AuthenticationPortalsDataSourceSchema.Attributes,
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
