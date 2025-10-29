package models

import (
	"regexp"

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

// Package: network_services
// This file contains models for the network_services SDK package

// SdwanErrorCorrectionProfiles represents the Terraform model for SdwanErrorCorrectionProfiles
type SdwanErrorCorrectionProfiles struct {
	Tfid                types.String          `tfsdk:"tfid"`
	ActivationThreshold basetypes.Int64Value  `tfsdk:"activation_threshold"`
	Device              basetypes.StringValue `tfsdk:"device"`
	Folder              basetypes.StringValue `tfsdk:"folder"`
	Id                  basetypes.StringValue `tfsdk:"id"`
	Mode                basetypes.ObjectValue `tfsdk:"mode"`
	Name                basetypes.StringValue `tfsdk:"name"`
	Snippet             basetypes.StringValue `tfsdk:"snippet"`
}

// SdwanErrorCorrectionProfilesMode represents a nested structure within the SdwanErrorCorrectionProfiles model
type SdwanErrorCorrectionProfilesMode struct {
	ForwardErrorCorrection basetypes.ObjectValue `tfsdk:"forward_error_correction"`
	PacketDuplication      basetypes.ObjectValue `tfsdk:"packet_duplication"`
}

// SdwanErrorCorrectionProfilesModeForwardErrorCorrection represents a nested structure within the SdwanErrorCorrectionProfiles model
type SdwanErrorCorrectionProfilesModeForwardErrorCorrection struct {
	Ratio            basetypes.StringValue `tfsdk:"ratio"`
	RecoveryDuration basetypes.Int64Value  `tfsdk:"recovery_duration"`
}

// SdwanErrorCorrectionProfilesModePacketDuplication represents a nested structure within the SdwanErrorCorrectionProfiles model
type SdwanErrorCorrectionProfilesModePacketDuplication struct {
	RecoveryDurationPd basetypes.Int64Value `tfsdk:"recovery_duration_pd"`
}

// AttrTypes defines the attribute types for the SdwanErrorCorrectionProfiles model.
func (o SdwanErrorCorrectionProfiles) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":                 basetypes.StringType{},
		"activation_threshold": basetypes.Int64Type{},
		"device":               basetypes.StringType{},
		"folder":               basetypes.StringType{},
		"id":                   basetypes.StringType{},
		"mode": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"forward_error_correction": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"ratio":             basetypes.StringType{},
						"recovery_duration": basetypes.Int64Type{},
					},
				},
				"packet_duplication": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"recovery_duration_pd": basetypes.Int64Type{},
					},
				},
			},
		},
		"name":    basetypes.StringType{},
		"snippet": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of SdwanErrorCorrectionProfiles objects.
func (o SdwanErrorCorrectionProfiles) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the SdwanErrorCorrectionProfilesMode model.
func (o SdwanErrorCorrectionProfilesMode) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"forward_error_correction": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"ratio":             basetypes.StringType{},
				"recovery_duration": basetypes.Int64Type{},
			},
		},
		"packet_duplication": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"recovery_duration_pd": basetypes.Int64Type{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of SdwanErrorCorrectionProfilesMode objects.
func (o SdwanErrorCorrectionProfilesMode) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the SdwanErrorCorrectionProfilesModeForwardErrorCorrection model.
func (o SdwanErrorCorrectionProfilesModeForwardErrorCorrection) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"ratio":             basetypes.StringType{},
		"recovery_duration": basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of SdwanErrorCorrectionProfilesModeForwardErrorCorrection objects.
func (o SdwanErrorCorrectionProfilesModeForwardErrorCorrection) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the SdwanErrorCorrectionProfilesModePacketDuplication model.
func (o SdwanErrorCorrectionProfilesModePacketDuplication) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"recovery_duration_pd": basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of SdwanErrorCorrectionProfilesModePacketDuplication objects.
func (o SdwanErrorCorrectionProfilesModePacketDuplication) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// SdwanErrorCorrectionProfilesResourceSchema defines the schema for SdwanErrorCorrectionProfiles resource
var SdwanErrorCorrectionProfilesResourceSchema = schema.Schema{
	MarkdownDescription: "SdwanErrorCorrectionProfile resource",
	Attributes: map[string]schema.Attribute{
		"activation_threshold": schema.Int64Attribute{
			MarkdownDescription: "Activation threshold",
			Required:            true,
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
			MarkdownDescription: "UUID of the resource",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"mode": schema.SingleNestedAttribute{
			MarkdownDescription: "Mode",
			Required:            true,
			Attributes: map[string]schema.Attribute{
				"forward_error_correction": schema.SingleNestedAttribute{
					Validators: []validator.Object{
						objectvalidator.ExactlyOneOf(
							path.MatchRelative().AtParent().AtName("packet_duplication"),
						),
					},
					MarkdownDescription: "Forward error correction",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"ratio": schema.StringAttribute{
							MarkdownDescription: "Ratio",
							Required:            true,
						},
						"recovery_duration": schema.Int64Attribute{
							MarkdownDescription: "Recovery duration",
							Required:            true,
						},
					},
				},
				"packet_duplication": schema.SingleNestedAttribute{
					Validators: []validator.Object{
						objectvalidator.ExactlyOneOf(
							path.MatchRelative().AtParent().AtName("forward_error_correction"),
						),
					},
					MarkdownDescription: "Packet duplication",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"recovery_duration_pd": schema.Int64Attribute{
							MarkdownDescription: "Recovery duration pd",
							Required:            true,
						},
					},
				},
			},
		},
		"name": schema.StringAttribute{
			MarkdownDescription: "Name",
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
			MarkdownDescription: "The snippet in which the resource is defined",
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

// SdwanErrorCorrectionProfilesDataSourceSchema defines the schema for SdwanErrorCorrectionProfiles data source
var SdwanErrorCorrectionProfilesDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "SdwanErrorCorrectionProfile data source",
	Attributes: map[string]dsschema.Attribute{
		"activation_threshold": dsschema.Int64Attribute{
			MarkdownDescription: "Activation threshold",
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
			MarkdownDescription: "UUID of the resource",
			Required:            true,
		},
		"mode": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Mode",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"forward_error_correction": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Forward error correction",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"ratio": dsschema.StringAttribute{
							MarkdownDescription: "Ratio",
							Computed:            true,
						},
						"recovery_duration": dsschema.Int64Attribute{
							MarkdownDescription: "Recovery duration",
							Computed:            true,
						},
					},
				},
				"packet_duplication": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Packet duplication",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"recovery_duration_pd": dsschema.Int64Attribute{
							MarkdownDescription: "Recovery duration pd",
							Computed:            true,
						},
					},
				},
			},
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "Name",
			Optional:            true,
			Computed:            true,
		},
		"snippet": dsschema.StringAttribute{
			MarkdownDescription: "The snippet in which the resource is defined",
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
	},
}

// SdwanErrorCorrectionProfilesListModel represents the data model for a list data source.
type SdwanErrorCorrectionProfilesListModel struct {
	Tfid    types.String                   `tfsdk:"tfid"`
	Data    []SdwanErrorCorrectionProfiles `tfsdk:"data"`
	Limit   types.Int64                    `tfsdk:"limit"`
	Offset  types.Int64                    `tfsdk:"offset"`
	Name    types.String                   `tfsdk:"name"`
	Total   types.Int64                    `tfsdk:"total"`
	Folder  types.String                   `tfsdk:"folder"`
	Snippet types.String                   `tfsdk:"snippet"`
	Device  types.String                   `tfsdk:"device"`
}

// SdwanErrorCorrectionProfilesListDataSourceSchema defines the schema for a list data source.
var SdwanErrorCorrectionProfilesListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: SdwanErrorCorrectionProfilesDataSourceSchema.Attributes,
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
