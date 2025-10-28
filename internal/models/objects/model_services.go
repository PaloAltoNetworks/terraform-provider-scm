package models

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/objectvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	dsschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Package: objects
// This file contains models for the objects SDK package

// Services represents the Terraform model for Services
type Services struct {
	Tfid        types.String          `tfsdk:"tfid"`
	Description basetypes.StringValue `tfsdk:"description"`
	Device      basetypes.StringValue `tfsdk:"device"`
	Folder      basetypes.StringValue `tfsdk:"folder"`
	Id          basetypes.StringValue `tfsdk:"id"`
	Name        basetypes.StringValue `tfsdk:"name"`
	Protocol    basetypes.ObjectValue `tfsdk:"protocol"`
	Snippet     basetypes.StringValue `tfsdk:"snippet"`
	Tag         basetypes.ListValue   `tfsdk:"tag"`
}

// ServicesProtocol represents a nested structure within the Services model
type ServicesProtocol struct {
	Tcp basetypes.ObjectValue `tfsdk:"tcp"`
	Udp basetypes.ObjectValue `tfsdk:"udp"`
}

// ServicesProtocolTcp represents a nested structure within the Services model
type ServicesProtocolTcp struct {
	Override   basetypes.ObjectValue `tfsdk:"override"`
	Port       basetypes.StringValue `tfsdk:"port"`
	SourcePort basetypes.StringValue `tfsdk:"source_port"`
}

// ServicesProtocolTcpOverride represents a nested structure within the Services model
type ServicesProtocolTcpOverride struct {
	HalfcloseTimeout basetypes.Int64Value `tfsdk:"halfclose_timeout"`
	Timeout          basetypes.Int64Value `tfsdk:"timeout"`
	TimewaitTimeout  basetypes.Int64Value `tfsdk:"timewait_timeout"`
}

// ServicesProtocolUdp represents a nested structure within the Services model
type ServicesProtocolUdp struct {
	Override   basetypes.ObjectValue `tfsdk:"override"`
	Port       basetypes.StringValue `tfsdk:"port"`
	SourcePort basetypes.StringValue `tfsdk:"source_port"`
}

// ServicesProtocolUdpOverride represents a nested structure within the Services model
type ServicesProtocolUdpOverride struct {
	Timeout basetypes.Int64Value `tfsdk:"timeout"`
}

// AttrTypes defines the attribute types for the Services model.
func (o Services) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":        basetypes.StringType{},
		"description": basetypes.StringType{},
		"device":      basetypes.StringType{},
		"folder":      basetypes.StringType{},
		"id":          basetypes.StringType{},
		"name":        basetypes.StringType{},
		"protocol": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"tcp": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"override": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"halfclose_timeout": basetypes.Int64Type{},
								"timeout":           basetypes.Int64Type{},
								"timewait_timeout":  basetypes.Int64Type{},
							},
						},
						"port":        basetypes.StringType{},
						"source_port": basetypes.StringType{},
					},
				},
				"udp": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"override": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"timeout": basetypes.Int64Type{},
							},
						},
						"port":        basetypes.StringType{},
						"source_port": basetypes.StringType{},
					},
				},
			},
		},
		"snippet": basetypes.StringType{},
		"tag":     basetypes.ListType{ElemType: basetypes.StringType{}},
	}
}

// AttrType returns the attribute type for a list of Services objects.
func (o Services) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ServicesProtocol model.
func (o ServicesProtocol) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tcp": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"override": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"halfclose_timeout": basetypes.Int64Type{},
						"timeout":           basetypes.Int64Type{},
						"timewait_timeout":  basetypes.Int64Type{},
					},
				},
				"port":        basetypes.StringType{},
				"source_port": basetypes.StringType{},
			},
		},
		"udp": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"override": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"timeout": basetypes.Int64Type{},
					},
				},
				"port":        basetypes.StringType{},
				"source_port": basetypes.StringType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of ServicesProtocol objects.
func (o ServicesProtocol) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ServicesProtocolTcp model.
func (o ServicesProtocolTcp) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"override": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"halfclose_timeout": basetypes.Int64Type{},
				"timeout":           basetypes.Int64Type{},
				"timewait_timeout":  basetypes.Int64Type{},
			},
		},
		"port":        basetypes.StringType{},
		"source_port": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of ServicesProtocolTcp objects.
func (o ServicesProtocolTcp) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ServicesProtocolTcpOverride model.
func (o ServicesProtocolTcpOverride) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"halfclose_timeout": basetypes.Int64Type{},
		"timeout":           basetypes.Int64Type{},
		"timewait_timeout":  basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of ServicesProtocolTcpOverride objects.
func (o ServicesProtocolTcpOverride) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ServicesProtocolUdp model.
func (o ServicesProtocolUdp) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"override": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"timeout": basetypes.Int64Type{},
			},
		},
		"port":        basetypes.StringType{},
		"source_port": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of ServicesProtocolUdp objects.
func (o ServicesProtocolUdp) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ServicesProtocolUdpOverride model.
func (o ServicesProtocolUdpOverride) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"timeout": basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of ServicesProtocolUdpOverride objects.
func (o ServicesProtocolUdpOverride) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// ServicesResourceSchema defines the schema for Services resource
var ServicesResourceSchema = schema.Schema{
	MarkdownDescription: "Service resource",
	Attributes: map[string]schema.Attribute{
		"description": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(1023),
			},
			MarkdownDescription: "Description",
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
		"id": schema.StringAttribute{
			MarkdownDescription: "The UUID of the service",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"name": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(63),
				stringvalidator.RegexMatches(regexp.MustCompile("^[ a-zA-Z\\d.\\-_]+$"), "pattern must match "+"^[ a-zA-Z\\d.\\-_]+$"),
			},
			MarkdownDescription: "The name of the service",
			Required:            true,
		},
		"protocol": schema.SingleNestedAttribute{
			MarkdownDescription: "Protocol",
			Optional:            true,
			Computed:            true,
			Attributes: map[string]schema.Attribute{
				"tcp": schema.SingleNestedAttribute{
					Validators: []validator.Object{
						objectvalidator.ConflictsWith(
							path.MatchRelative().AtParent().AtName("udp"),
						),
					},
					MarkdownDescription: "Tcp",
					Optional:            true,
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"override": schema.SingleNestedAttribute{
							MarkdownDescription: "Override",
							Optional:            true,
							Computed:            true,
							Attributes: map[string]schema.Attribute{
								"halfclose_timeout": schema.Int64Attribute{
									Validators: []validator.Int64{
										int64validator.Between(1, 604800),
									},
									MarkdownDescription: "tcp session half-close timeout value (in second)",
									Optional:            true,
									Computed:            true,
									Default:             int64default.StaticInt64(120),
								},
								"timeout": schema.Int64Attribute{
									Validators: []validator.Int64{
										int64validator.Between(1, 604800),
									},
									MarkdownDescription: "tcp session timeout value (in second)",
									Optional:            true,
									Computed:            true,
									Default:             int64default.StaticInt64(3600),
								},
								"timewait_timeout": schema.Int64Attribute{
									Validators: []validator.Int64{
										int64validator.Between(1, 600),
									},
									MarkdownDescription: "tcp session time-wait timeout value (in second)",
									Optional:            true,
									Computed:            true,
									Default:             int64default.StaticInt64(15),
								},
							},
						},
						"port": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.LengthAtMost(1023),
								stringvalidator.LengthAtLeast(1),
							},
							MarkdownDescription: "Port",
							Required:            true,
						},
						"source_port": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.LengthAtMost(1023),
								stringvalidator.LengthAtLeast(1),
							},
							MarkdownDescription: "Source port",
							Optional:            true,
							Computed:            true,
						},
					},
				},
				"udp": schema.SingleNestedAttribute{
					Validators: []validator.Object{
						objectvalidator.ConflictsWith(
							path.MatchRelative().AtParent().AtName("tcp"),
						),
					},
					MarkdownDescription: "Udp",
					Optional:            true,
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"override": schema.SingleNestedAttribute{
							MarkdownDescription: "Override",
							Optional:            true,
							Computed:            true,
							Attributes: map[string]schema.Attribute{
								"timeout": schema.Int64Attribute{
									Validators: []validator.Int64{
										int64validator.Between(1, 604800),
									},
									MarkdownDescription: "udp session timeout value (in second)",
									Optional:            true,
									Computed:            true,
									Default:             int64default.StaticInt64(30),
								},
							},
						},
						"port": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.LengthAtMost(1023),
								stringvalidator.LengthAtLeast(1),
							},
							MarkdownDescription: "Port",
							Required:            true,
						},
						"source_port": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.LengthAtMost(1023),
								stringvalidator.LengthAtLeast(1),
							},
							MarkdownDescription: "Source port",
							Optional:            true,
							Computed:            true,
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
			MarkdownDescription: "The snippet in which the resource is defined",
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"tag": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Tags for service object",
			Validators: []validator.List{
				listvalidator.SizeAtMost(64),
				listvalidator.ValueStringsAre(stringvalidator.LengthAtMost(127)),
			},
			Optional: true,
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

// ServicesDataSourceSchema defines the schema for Services data source
var ServicesDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Service data source",
	Attributes: map[string]dsschema.Attribute{
		"description": dsschema.StringAttribute{
			MarkdownDescription: "Description",
			Computed:            true,
		},
		"device": dsschema.StringAttribute{
			MarkdownDescription: "The device in which the resource is defined",
			Computed:            true,
		},
		"folder": dsschema.StringAttribute{
			MarkdownDescription: "The folder in which the resource is defined",
			Computed:            true,
		},
		"id": dsschema.StringAttribute{
			MarkdownDescription: "The UUID of the service",
			Required:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "The name of the service",
			Optional:            true,
			Computed:            true,
		},
		"protocol": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Protocol",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"tcp": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Tcp",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"override": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Override",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"halfclose_timeout": dsschema.Int64Attribute{
									MarkdownDescription: "tcp session half-close timeout value (in second)",
									Computed:            true,
								},
								"timeout": dsschema.Int64Attribute{
									MarkdownDescription: "tcp session timeout value (in second)",
									Computed:            true,
								},
								"timewait_timeout": dsschema.Int64Attribute{
									MarkdownDescription: "tcp session time-wait timeout value (in second)",
									Computed:            true,
								},
							},
						},
						"port": dsschema.StringAttribute{
							MarkdownDescription: "Port",
							Computed:            true,
						},
						"source_port": dsschema.StringAttribute{
							MarkdownDescription: "Source port",
							Computed:            true,
						},
					},
				},
				"udp": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Udp",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"override": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Override",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"timeout": dsschema.Int64Attribute{
									MarkdownDescription: "udp session timeout value (in second)",
									Computed:            true,
								},
							},
						},
						"port": dsschema.StringAttribute{
							MarkdownDescription: "Port",
							Computed:            true,
						},
						"source_port": dsschema.StringAttribute{
							MarkdownDescription: "Source port",
							Computed:            true,
						},
					},
				},
			},
		},
		"snippet": dsschema.StringAttribute{
			MarkdownDescription: "The snippet in which the resource is defined",
			Computed:            true,
		},
		"tag": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Tags for service object",
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
	},
}

// ServicesListModel represents the data model for a list data source.
type ServicesListModel struct {
	Tfid    types.String `tfsdk:"tfid"`
	Data    []Services   `tfsdk:"data"`
	Limit   types.Int64  `tfsdk:"limit"`
	Offset  types.Int64  `tfsdk:"offset"`
	Name    types.String `tfsdk:"name"`
	Total   types.Int64  `tfsdk:"total"`
	Folder  types.String `tfsdk:"folder"`
	Snippet types.String `tfsdk:"snippet"`
	Device  types.String `tfsdk:"device"`
}

// ServicesListDataSourceSchema defines the schema for a list data source.
var ServicesListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: ServicesDataSourceSchema.Attributes,
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
