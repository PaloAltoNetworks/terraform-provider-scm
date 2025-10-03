package models

import (
	"regexp"

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

// InterfaceManagementProfiles represents the Terraform model for InterfaceManagementProfiles
type InterfaceManagementProfiles struct {
	Tfid                    types.String          `tfsdk:"tfid"`
	Device                  basetypes.StringValue `tfsdk:"device"`
	Folder                  basetypes.StringValue `tfsdk:"folder"`
	Http                    basetypes.BoolValue   `tfsdk:"http"`
	HttpOcsp                basetypes.BoolValue   `tfsdk:"http_ocsp"`
	Https                   basetypes.BoolValue   `tfsdk:"https"`
	Id                      basetypes.StringValue `tfsdk:"id"`
	Name                    basetypes.StringValue `tfsdk:"name"`
	PermittedIp             basetypes.ListValue   `tfsdk:"permitted_ip"`
	Ping                    basetypes.BoolValue   `tfsdk:"ping"`
	ResponsePages           basetypes.StringValue `tfsdk:"response_pages"`
	Snippet                 basetypes.StringValue `tfsdk:"snippet"`
	Ssh                     basetypes.BoolValue   `tfsdk:"ssh"`
	Telnet                  basetypes.BoolValue   `tfsdk:"telnet"`
	UseridService           basetypes.BoolValue   `tfsdk:"userid_service"`
	UseridSyslogListenerSsl basetypes.BoolValue   `tfsdk:"userid_syslog_listener_ssl"`
	UseridSyslogListenerUdp basetypes.BoolValue   `tfsdk:"userid_syslog_listener_udp"`
}

// AttrTypes defines the attribute types for the InterfaceManagementProfiles model.
func (o InterfaceManagementProfiles) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":                       basetypes.StringType{},
		"device":                     basetypes.StringType{},
		"folder":                     basetypes.StringType{},
		"http":                       basetypes.BoolType{},
		"http_ocsp":                  basetypes.BoolType{},
		"https":                      basetypes.BoolType{},
		"id":                         basetypes.StringType{},
		"name":                       basetypes.StringType{},
		"permitted_ip":               basetypes.ListType{ElemType: basetypes.StringType{}},
		"ping":                       basetypes.BoolType{},
		"response_pages":             basetypes.StringType{},
		"snippet":                    basetypes.StringType{},
		"ssh":                        basetypes.BoolType{},
		"telnet":                     basetypes.BoolType{},
		"userid_service":             basetypes.BoolType{},
		"userid_syslog_listener_ssl": basetypes.BoolType{},
		"userid_syslog_listener_udp": basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of InterfaceManagementProfiles objects.
func (o InterfaceManagementProfiles) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// InterfaceManagementProfilesResourceSchema defines the schema for InterfaceManagementProfiles resource
var InterfaceManagementProfilesResourceSchema = schema.Schema{
	MarkdownDescription: "Manages SCM InterfaceManagementProfiles objects",
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
			MarkdownDescription: "The device in which the resource is defined",
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"folder": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.ExactlyOneOf(
					path.MatchRelative().AtParent().AtName("snippet"),
					path.MatchRelative().AtParent().AtName("device"),
				),
				stringvalidator.LengthAtMost(64),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z\\d-_\\. ]+$"), "pattern must match "+"^[a-zA-Z\\d-_\\. ]+$"),
			},
			MarkdownDescription: "The folder in which the resource is defined",
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"http": schema.BoolAttribute{
			MarkdownDescription: "Allow HTTP?",
			Optional:            true,
		},
		"http_ocsp": schema.BoolAttribute{
			MarkdownDescription: "Http ocsp",
			Optional:            true,
		},
		"https": schema.BoolAttribute{
			MarkdownDescription: "Allow HTTPS?",
			Optional:            true,
		},
		"id": schema.StringAttribute{
			MarkdownDescription: "UUID of the resource",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"name": schema.StringAttribute{
			MarkdownDescription: "Name",
			Required:            true,
		},
		"permitted_ip": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Permitted ip",
			Optional:            true,
		},
		"ping": schema.BoolAttribute{
			MarkdownDescription: "Allow ping?",
			Optional:            true,
		},
		"response_pages": schema.StringAttribute{
			MarkdownDescription: "Response pages",
			Optional:            true,
		},
		"snippet": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.ExactlyOneOf(
					path.MatchRelative().AtParent().AtName("folder"),
					path.MatchRelative().AtParent().AtName("device"),
				),
				stringvalidator.LengthAtMost(64),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z\\d-_\\. ]+$"), "pattern must match "+"^[a-zA-Z\\d-_\\. ]+$"),
			},
			MarkdownDescription: "The snippet in which the resource is defined",
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"ssh": schema.BoolAttribute{
			MarkdownDescription: "Allow SSH?",
			Optional:            true,
		},
		"telnet": schema.BoolAttribute{
			MarkdownDescription: "Allow telnet? Seriously, why would you do this?!?",
			Optional:            true,
		},
		"tfid": schema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"userid_service": schema.BoolAttribute{
			MarkdownDescription: "Userid service",
			Optional:            true,
		},
		"userid_syslog_listener_ssl": schema.BoolAttribute{
			MarkdownDescription: "Userid syslog listener ssl",
			Optional:            true,
		},
		"userid_syslog_listener_udp": schema.BoolAttribute{
			MarkdownDescription: "Userid syslog listener udp",
			Optional:            true,
		},
	},
}

// InterfaceManagementProfilesDataSourceSchema defines the schema for InterfaceManagementProfiles data source
var InterfaceManagementProfilesDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "InterfaceManagementProfiles data source",
	Attributes: map[string]dsschema.Attribute{
		"device": dsschema.StringAttribute{
			MarkdownDescription: "The device in which the resource is defined",
			Computed:            true,
		},
		"folder": dsschema.StringAttribute{
			MarkdownDescription: "The folder in which the resource is defined",
			Computed:            true,
		},
		"http": dsschema.BoolAttribute{
			MarkdownDescription: "Allow HTTP?",
			Computed:            true,
		},
		"http_ocsp": dsschema.BoolAttribute{
			MarkdownDescription: "Http ocsp",
			Computed:            true,
		},
		"https": dsschema.BoolAttribute{
			MarkdownDescription: "Allow HTTPS?",
			Computed:            true,
		},
		"id": dsschema.StringAttribute{
			MarkdownDescription: "UUID of the resource",
			Required:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "Name",
			Optional:            true,
			Computed:            true,
		},
		"permitted_ip": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Permitted ip",
			Computed:            true,
		},
		"ping": dsschema.BoolAttribute{
			MarkdownDescription: "Allow ping?",
			Computed:            true,
		},
		"response_pages": dsschema.StringAttribute{
			MarkdownDescription: "Response pages",
			Computed:            true,
		},
		"snippet": dsschema.StringAttribute{
			MarkdownDescription: "The snippet in which the resource is defined",
			Computed:            true,
		},
		"ssh": dsschema.BoolAttribute{
			MarkdownDescription: "Allow SSH?",
			Computed:            true,
		},
		"telnet": dsschema.BoolAttribute{
			MarkdownDescription: "Allow telnet? Seriously, why would you do this?!?",
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
		"userid_service": dsschema.BoolAttribute{
			MarkdownDescription: "Userid service",
			Computed:            true,
		},
		"userid_syslog_listener_ssl": dsschema.BoolAttribute{
			MarkdownDescription: "Userid syslog listener ssl",
			Computed:            true,
		},
		"userid_syslog_listener_udp": dsschema.BoolAttribute{
			MarkdownDescription: "Userid syslog listener udp",
			Computed:            true,
		},
	},
}

// InterfaceManagementProfilesListModel represents the data model for a list data source.
type InterfaceManagementProfilesListModel struct {
	Tfid    types.String                  `tfsdk:"tfid"`
	Data    []InterfaceManagementProfiles `tfsdk:"data"`
	Limit   types.Int64                   `tfsdk:"limit"`
	Offset  types.Int64                   `tfsdk:"offset"`
	Name    types.String                  `tfsdk:"name"`
	Total   types.Int64                   `tfsdk:"total"`
	Folder  types.String                  `tfsdk:"folder"`
	Snippet types.String                  `tfsdk:"snippet"`
	Device  types.String                  `tfsdk:"device"`
}

// InterfaceManagementProfilesListDataSourceSchema defines the schema for a list data source.
var InterfaceManagementProfilesListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: InterfaceManagementProfilesDataSourceSchema.Attributes,
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
