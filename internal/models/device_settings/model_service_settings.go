package models

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/objectvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	dsschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Package: device_settings
// This file contains models for the device_settings SDK package

// ServiceSettings represents the Terraform model for ServiceSettings
type ServiceSettings struct {
	Tfid            types.String          `tfsdk:"tfid"`
	EncryptedValues basetypes.MapValue    `tfsdk:"encrypted_values"`
	Device          basetypes.StringValue `tfsdk:"device"`
	Folder          basetypes.StringValue `tfsdk:"folder"`
	Id              basetypes.StringValue `tfsdk:"id"`
	Services        basetypes.ObjectValue `tfsdk:"services"`
	Snippet         basetypes.StringValue `tfsdk:"snippet"`
}

// ServiceSettingsServices represents a nested structure within the ServiceSettings model
type ServiceSettingsServices struct {
	DnsSetting            basetypes.ObjectValue  `tfsdk:"dns_setting"`
	FqdnRefreshTime       basetypes.Float64Value `tfsdk:"fqdn_refresh_time"`
	FqdnStaleEntryTimeout basetypes.Float64Value `tfsdk:"fqdn_stale_entry_timeout"`
	InlineCloudProxy      basetypes.BoolValue    `tfsdk:"inline_cloud_proxy"`
	LcaasUseProxy         basetypes.BoolValue    `tfsdk:"lcaas_use_proxy"`
	NtpServers            basetypes.ObjectValue  `tfsdk:"ntp_servers"`
	SecureProxyPassword   basetypes.StringValue  `tfsdk:"secure_proxy_password"`
	SecureProxyPort       basetypes.Float64Value `tfsdk:"secure_proxy_port"`
	SecureProxyServer     basetypes.StringValue  `tfsdk:"secure_proxy_server"`
	SecureProxyUser       basetypes.StringValue  `tfsdk:"secure_proxy_user"`
	ServerVerification    basetypes.BoolValue    `tfsdk:"server_verification"`
	UpdateServer          basetypes.StringValue  `tfsdk:"update_server"`
}

// ServiceSettingsServicesDnsSetting represents a nested structure within the ServiceSettings model
type ServiceSettingsServicesDnsSetting struct {
	DnsProxyObject basetypes.StringValue `tfsdk:"dns_proxy_object"`
	Servers        basetypes.ObjectValue `tfsdk:"servers"`
}

// ServiceSettingsServicesDnsSettingServers represents a nested structure within the ServiceSettings model
type ServiceSettingsServicesDnsSettingServers struct {
	Primary   basetypes.StringValue `tfsdk:"primary"`
	Secondary basetypes.StringValue `tfsdk:"secondary"`
}

// ServiceSettingsServicesNtpServers represents a nested structure within the ServiceSettings model
type ServiceSettingsServicesNtpServers struct {
	PrimaryNtpServer   basetypes.ObjectValue `tfsdk:"primary_ntp_server"`
	SecondaryNtpServer basetypes.ObjectValue `tfsdk:"secondary_ntp_server"`
}

// ServiceSettingsServicesNtpServersPrimaryNtpServer represents a nested structure within the ServiceSettings model
type ServiceSettingsServicesNtpServersPrimaryNtpServer struct {
	AuthenticationType basetypes.ObjectValue `tfsdk:"authentication_type"`
	NtpServerAddress   basetypes.StringValue `tfsdk:"ntp_server_address"`
}

// ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationType represents a nested structure within the ServiceSettings model
type ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationType struct {
	Autokey      basetypes.ObjectValue `tfsdk:"autokey"`
	None         basetypes.ObjectValue `tfsdk:"none"`
	SymmetricKey basetypes.ObjectValue `tfsdk:"symmetric_key"`
}

// ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey represents a nested structure within the ServiceSettings model
type ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey struct {
	Algorithm basetypes.ObjectValue  `tfsdk:"algorithm"`
	KeyId     basetypes.Float64Value `tfsdk:"key_id"`
}

// ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm represents a nested structure within the ServiceSettings model
type ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm struct {
	Md5  basetypes.ObjectValue `tfsdk:"md5"`
	Sha1 basetypes.ObjectValue `tfsdk:"sha1"`
}

// ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5 represents a nested structure within the ServiceSettings model
type ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5 struct {
	AuthenticationKey basetypes.StringValue `tfsdk:"authentication_key"`
}

// AttrTypes defines the attribute types for the ServiceSettings model.
func (o ServiceSettings) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":             basetypes.StringType{},
		"encrypted_values": basetypes.MapType{ElemType: basetypes.StringType{}},
		"device":           basetypes.StringType{},
		"folder":           basetypes.StringType{},
		"id":               basetypes.StringType{},
		"services": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"dns_setting": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"dns_proxy_object": basetypes.StringType{},
						"servers": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"primary":   basetypes.StringType{},
								"secondary": basetypes.StringType{},
							},
						},
					},
				},
				"fqdn_refresh_time":        basetypes.Float64Type{},
				"fqdn_stale_entry_timeout": basetypes.Float64Type{},
				"inline_cloud_proxy":       basetypes.BoolType{},
				"lcaas_use_proxy":          basetypes.BoolType{},
				"ntp_servers": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"primary_ntp_server": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"authentication_type": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"autokey": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
										"none": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
										"symmetric_key": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"algorithm": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"md5": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{
																"authentication_key": basetypes.StringType{},
															},
														},
														"sha1": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{
																"authentication_key": basetypes.StringType{},
															},
														},
													},
												},
												"key_id": basetypes.Float64Type{},
											},
										},
									},
								},
								"ntp_server_address": basetypes.StringType{},
							},
						},
						"secondary_ntp_server": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"authentication_type": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"autokey": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
										"none": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
										"symmetric_key": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"algorithm": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"md5": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{
																"authentication_key": basetypes.StringType{},
															},
														},
														"sha1": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{
																"authentication_key": basetypes.StringType{},
															},
														},
													},
												},
												"key_id": basetypes.Float64Type{},
											},
										},
									},
								},
								"ntp_server_address": basetypes.StringType{},
							},
						},
					},
				},
				"secure_proxy_password": basetypes.StringType{},
				"secure_proxy_port":     basetypes.Float64Type{},
				"secure_proxy_server":   basetypes.StringType{},
				"secure_proxy_user":     basetypes.StringType{},
				"server_verification":   basetypes.BoolType{},
				"update_server":         basetypes.StringType{},
			},
		},
		"snippet": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of ServiceSettings objects.
func (o ServiceSettings) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ServiceSettingsServices model.
func (o ServiceSettingsServices) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"dns_setting": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"dns_proxy_object": basetypes.StringType{},
				"servers": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"primary":   basetypes.StringType{},
						"secondary": basetypes.StringType{},
					},
				},
			},
		},
		"fqdn_refresh_time":        basetypes.Float64Type{},
		"fqdn_stale_entry_timeout": basetypes.Float64Type{},
		"inline_cloud_proxy":       basetypes.BoolType{},
		"lcaas_use_proxy":          basetypes.BoolType{},
		"ntp_servers": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"primary_ntp_server": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"authentication_type": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"autokey": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"none": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"symmetric_key": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"algorithm": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"md5": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"authentication_key": basetypes.StringType{},
													},
												},
												"sha1": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"authentication_key": basetypes.StringType{},
													},
												},
											},
										},
										"key_id": basetypes.Float64Type{},
									},
								},
							},
						},
						"ntp_server_address": basetypes.StringType{},
					},
				},
				"secondary_ntp_server": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"authentication_type": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"autokey": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"none": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"symmetric_key": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"algorithm": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"md5": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"authentication_key": basetypes.StringType{},
													},
												},
												"sha1": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"authentication_key": basetypes.StringType{},
													},
												},
											},
										},
										"key_id": basetypes.Float64Type{},
									},
								},
							},
						},
						"ntp_server_address": basetypes.StringType{},
					},
				},
			},
		},
		"secure_proxy_password": basetypes.StringType{},
		"secure_proxy_port":     basetypes.Float64Type{},
		"secure_proxy_server":   basetypes.StringType{},
		"secure_proxy_user":     basetypes.StringType{},
		"server_verification":   basetypes.BoolType{},
		"update_server":         basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of ServiceSettingsServices objects.
func (o ServiceSettingsServices) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ServiceSettingsServicesDnsSetting model.
func (o ServiceSettingsServicesDnsSetting) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"dns_proxy_object": basetypes.StringType{},
		"servers": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"primary":   basetypes.StringType{},
				"secondary": basetypes.StringType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of ServiceSettingsServicesDnsSetting objects.
func (o ServiceSettingsServicesDnsSetting) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ServiceSettingsServicesDnsSettingServers model.
func (o ServiceSettingsServicesDnsSettingServers) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"primary":   basetypes.StringType{},
		"secondary": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of ServiceSettingsServicesDnsSettingServers objects.
func (o ServiceSettingsServicesDnsSettingServers) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ServiceSettingsServicesNtpServers model.
func (o ServiceSettingsServicesNtpServers) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"primary_ntp_server": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"authentication_type": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"autokey": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"none": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"symmetric_key": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"algorithm": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"md5": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"authentication_key": basetypes.StringType{},
											},
										},
										"sha1": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"authentication_key": basetypes.StringType{},
											},
										},
									},
								},
								"key_id": basetypes.Float64Type{},
							},
						},
					},
				},
				"ntp_server_address": basetypes.StringType{},
			},
		},
		"secondary_ntp_server": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"authentication_type": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"autokey": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"none": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"symmetric_key": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"algorithm": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"md5": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"authentication_key": basetypes.StringType{},
											},
										},
										"sha1": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"authentication_key": basetypes.StringType{},
											},
										},
									},
								},
								"key_id": basetypes.Float64Type{},
							},
						},
					},
				},
				"ntp_server_address": basetypes.StringType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of ServiceSettingsServicesNtpServers objects.
func (o ServiceSettingsServicesNtpServers) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ServiceSettingsServicesNtpServersPrimaryNtpServer model.
func (o ServiceSettingsServicesNtpServersPrimaryNtpServer) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"authentication_type": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"autokey": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"none": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"symmetric_key": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"algorithm": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"md5": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"authentication_key": basetypes.StringType{},
									},
								},
								"sha1": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"authentication_key": basetypes.StringType{},
									},
								},
							},
						},
						"key_id": basetypes.Float64Type{},
					},
				},
			},
		},
		"ntp_server_address": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of ServiceSettingsServicesNtpServersPrimaryNtpServer objects.
func (o ServiceSettingsServicesNtpServersPrimaryNtpServer) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationType model.
func (o ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationType) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"autokey": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
		"none": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
		"symmetric_key": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"algorithm": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"md5": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"authentication_key": basetypes.StringType{},
							},
						},
						"sha1": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"authentication_key": basetypes.StringType{},
							},
						},
					},
				},
				"key_id": basetypes.Float64Type{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationType objects.
func (o ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationType) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey model.
func (o ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"algorithm": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"md5": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"authentication_key": basetypes.StringType{},
					},
				},
				"sha1": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"authentication_key": basetypes.StringType{},
					},
				},
			},
		},
		"key_id": basetypes.Float64Type{},
	}
}

// AttrType returns the attribute type for a list of ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey objects.
func (o ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm model.
func (o ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"md5": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"authentication_key": basetypes.StringType{},
			},
		},
		"sha1": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"authentication_key": basetypes.StringType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm objects.
func (o ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5 model.
func (o ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"authentication_key": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5 objects.
func (o ServiceSettingsServicesNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// ServiceSettingsResourceSchema defines the schema for ServiceSettings resource
var ServiceSettingsResourceSchema = schema.Schema{
	MarkdownDescription: "ServiceSetting resource",
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
		"encrypted_values": schema.MapAttribute{
			ElementType:         basetypes.StringType{},
			MarkdownDescription: "Map of sensitive values returned from the API.",
			Computed:            true,
			Sensitive:           true,
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
		"services": schema.SingleNestedAttribute{
			MarkdownDescription: "Services",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"dns_setting": schema.SingleNestedAttribute{
					MarkdownDescription: "Dns setting",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"dns_proxy_object": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.ConflictsWith(
									path.MatchRelative().AtParent().AtName("servers"),
								),
							},
							MarkdownDescription: "Dns proxy object\n> ℹ️ **Note:** You must specify exactly one of `dns_proxy_object` and `servers`.",
							Optional:            true,
						},
						"servers": schema.SingleNestedAttribute{
							Validators: []validator.Object{
								objectvalidator.ConflictsWith(
									path.MatchRelative().AtParent().AtName("dns_proxy_object"),
								),
							},
							MarkdownDescription: "Servers\n> ℹ️ **Note:** You must specify exactly one of `dns_proxy_object` and `servers`.",
							Optional:            true,
							Attributes: map[string]schema.Attribute{
								"primary": schema.StringAttribute{
									MarkdownDescription: "Primary",
									Optional:            true,
								},
								"secondary": schema.StringAttribute{
									MarkdownDescription: "Secondary",
									Optional:            true,
								},
							},
						},
					},
				},
				"fqdn_refresh_time": schema.Float64Attribute{
					MarkdownDescription: "Fqdn refresh time",
					Optional:            true,
					Computed:            true,
					Default:             float64default.StaticFloat64(15.000000),
				},
				"fqdn_stale_entry_timeout": schema.Float64Attribute{
					MarkdownDescription: "Fqdn stale entry timeout",
					Optional:            true,
					Computed:            true,
					Default:             float64default.StaticFloat64(1440.000000),
				},
				"inline_cloud_proxy": schema.BoolAttribute{
					MarkdownDescription: "Inline cloud proxy",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(false),
				},
				"lcaas_use_proxy": schema.BoolAttribute{
					MarkdownDescription: "Lcaas use proxy",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(false),
				},
				"ntp_servers": schema.SingleNestedAttribute{
					MarkdownDescription: "Ntp servers",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"primary_ntp_server": schema.SingleNestedAttribute{
							MarkdownDescription: "Primary ntp server",
							Optional:            true,
							Attributes: map[string]schema.Attribute{
								"authentication_type": schema.SingleNestedAttribute{
									MarkdownDescription: "Authentication type",
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"autokey": schema.SingleNestedAttribute{
											Validators: []validator.Object{
												objectvalidator.ConflictsWith(
													path.MatchRelative().AtParent().AtName("none"),
													path.MatchRelative().AtParent().AtName("symmetric_key"),
												),
											},
											MarkdownDescription: "Autokey\n> ℹ️ **Note:** You must specify exactly one of `autokey`, `none`, and `symmetric_key`.",
											Optional:            true,
											Attributes:          map[string]schema.Attribute{},
										},
										"none": schema.SingleNestedAttribute{
											Validators: []validator.Object{
												objectvalidator.ConflictsWith(
													path.MatchRelative().AtParent().AtName("autokey"),
													path.MatchRelative().AtParent().AtName("symmetric_key"),
												),
											},
											MarkdownDescription: "None\n> ℹ️ **Note:** You must specify exactly one of `autokey`, `none`, and `symmetric_key`.",
											Optional:            true,
											Attributes:          map[string]schema.Attribute{},
										},
										"symmetric_key": schema.SingleNestedAttribute{
											Validators: []validator.Object{
												objectvalidator.ConflictsWith(
													path.MatchRelative().AtParent().AtName("autokey"),
													path.MatchRelative().AtParent().AtName("none"),
												),
											},
											MarkdownDescription: "Symmetric key\n> ℹ️ **Note:** You must specify exactly one of `autokey`, `none`, and `symmetric_key`.",
											Optional:            true,
											Attributes: map[string]schema.Attribute{
												"algorithm": schema.SingleNestedAttribute{
													MarkdownDescription: "Algorithm",
													Optional:            true,
													Attributes: map[string]schema.Attribute{
														"md5": schema.SingleNestedAttribute{
															MarkdownDescription: "Md5",
															Optional:            true,
															Attributes: map[string]schema.Attribute{
																"authentication_key": schema.StringAttribute{
																	MarkdownDescription: "Authentication key",
																	Optional:            true,
																},
															},
														},
														"sha1": schema.SingleNestedAttribute{
															MarkdownDescription: "Sha1",
															Optional:            true,
															Attributes: map[string]schema.Attribute{
																"authentication_key": schema.StringAttribute{
																	MarkdownDescription: "Authentication key",
																	Optional:            true,
																},
															},
														},
													},
												},
												"key_id": schema.Float64Attribute{
													MarkdownDescription: "Key id",
													Optional:            true,
												},
											},
										},
									},
								},
								"ntp_server_address": schema.StringAttribute{
									MarkdownDescription: "Ntp server address",
									Optional:            true,
								},
							},
						},
						"secondary_ntp_server": schema.SingleNestedAttribute{
							MarkdownDescription: "Secondary ntp server",
							Optional:            true,
							Attributes: map[string]schema.Attribute{
								"authentication_type": schema.SingleNestedAttribute{
									MarkdownDescription: "Authentication type",
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"autokey": schema.SingleNestedAttribute{
											Validators: []validator.Object{
												objectvalidator.ConflictsWith(
													path.MatchRelative().AtParent().AtName("none"),
													path.MatchRelative().AtParent().AtName("symmetric_key"),
												),
											},
											MarkdownDescription: "Autokey\n> ℹ️ **Note:** You must specify exactly one of `autokey`, `none`, and `symmetric_key`.",
											Optional:            true,
											Attributes:          map[string]schema.Attribute{},
										},
										"none": schema.SingleNestedAttribute{
											Validators: []validator.Object{
												objectvalidator.ConflictsWith(
													path.MatchRelative().AtParent().AtName("autokey"),
													path.MatchRelative().AtParent().AtName("symmetric_key"),
												),
											},
											MarkdownDescription: "None\n> ℹ️ **Note:** You must specify exactly one of `autokey`, `none`, and `symmetric_key`.",
											Optional:            true,
											Attributes:          map[string]schema.Attribute{},
										},
										"symmetric_key": schema.SingleNestedAttribute{
											Validators: []validator.Object{
												objectvalidator.ConflictsWith(
													path.MatchRelative().AtParent().AtName("autokey"),
													path.MatchRelative().AtParent().AtName("none"),
												),
											},
											MarkdownDescription: "Symmetric key\n> ℹ️ **Note:** You must specify exactly one of `autokey`, `none`, and `symmetric_key`.",
											Optional:            true,
											Attributes: map[string]schema.Attribute{
												"algorithm": schema.SingleNestedAttribute{
													MarkdownDescription: "Algorithm",
													Optional:            true,
													Attributes: map[string]schema.Attribute{
														"md5": schema.SingleNestedAttribute{
															MarkdownDescription: "Md5",
															Optional:            true,
															Attributes: map[string]schema.Attribute{
																"authentication_key": schema.StringAttribute{
																	MarkdownDescription: "Authentication key",
																	Optional:            true,
																},
															},
														},
														"sha1": schema.SingleNestedAttribute{
															MarkdownDescription: "Sha1",
															Optional:            true,
															Attributes: map[string]schema.Attribute{
																"authentication_key": schema.StringAttribute{
																	MarkdownDescription: "Authentication key",
																	Optional:            true,
																},
															},
														},
													},
												},
												"key_id": schema.Float64Attribute{
													MarkdownDescription: "Key id",
													Optional:            true,
												},
											},
										},
									},
								},
								"ntp_server_address": schema.StringAttribute{
									MarkdownDescription: "Ntp server address",
									Optional:            true,
								},
							},
						},
					},
				},
				"secure_proxy_password": schema.StringAttribute{
					MarkdownDescription: "Secure proxy password",
					Optional:            true,
					Sensitive:           true,
				},
				"secure_proxy_port": schema.Float64Attribute{
					MarkdownDescription: "Secure proxy port",
					Optional:            true,
				},
				"secure_proxy_server": schema.StringAttribute{
					MarkdownDescription: "Secure proxy server",
					Optional:            true,
				},
				"secure_proxy_user": schema.StringAttribute{
					MarkdownDescription: "Secure proxy user",
					Optional:            true,
				},
				"server_verification": schema.BoolAttribute{
					MarkdownDescription: "Server verification",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(true),
				},
				"update_server": schema.StringAttribute{
					MarkdownDescription: "Update server",
					Optional:            true,
					Computed:            true,
					Default:             stringdefault.StaticString("updates.paloaltonetworks.com"),
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
	},
}

// ServiceSettingsDataSourceSchema defines the schema for ServiceSettings data source
var ServiceSettingsDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "ServiceSetting data source",
	Attributes: map[string]dsschema.Attribute{
		"device": dsschema.StringAttribute{
			MarkdownDescription: "The device in which the resource is defined\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			Computed:            true,
		},
		"encrypted_values": dsschema.MapAttribute{
			ElementType:         basetypes.StringType{},
			MarkdownDescription: "Map of sensitive values returned from the API.",
			Computed:            true,
			Sensitive:           true,
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
		"services": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Services",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"dns_setting": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Dns setting",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"dns_proxy_object": dsschema.StringAttribute{
							MarkdownDescription: "Dns proxy object\n> ℹ️ **Note:** You must specify exactly one of `dns_proxy_object` and `servers`.",
							Computed:            true,
						},
						"servers": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Servers\n> ℹ️ **Note:** You must specify exactly one of `dns_proxy_object` and `servers`.",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"primary": dsschema.StringAttribute{
									MarkdownDescription: "Primary",
									Computed:            true,
								},
								"secondary": dsschema.StringAttribute{
									MarkdownDescription: "Secondary",
									Computed:            true,
								},
							},
						},
					},
				},
				"fqdn_refresh_time": dsschema.Float64Attribute{
					MarkdownDescription: "Fqdn refresh time",
					Computed:            true,
				},
				"fqdn_stale_entry_timeout": dsschema.Float64Attribute{
					MarkdownDescription: "Fqdn stale entry timeout",
					Computed:            true,
				},
				"inline_cloud_proxy": dsschema.BoolAttribute{
					MarkdownDescription: "Inline cloud proxy",
					Computed:            true,
				},
				"lcaas_use_proxy": dsschema.BoolAttribute{
					MarkdownDescription: "Lcaas use proxy",
					Computed:            true,
				},
				"ntp_servers": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Ntp servers",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"primary_ntp_server": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Primary ntp server",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"authentication_type": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Authentication type",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"autokey": dsschema.SingleNestedAttribute{
											MarkdownDescription: "Autokey\n> ℹ️ **Note:** You must specify exactly one of `autokey`, `none`, and `symmetric_key`.",
											Computed:            true,
											Attributes:          map[string]dsschema.Attribute{},
										},
										"none": dsschema.SingleNestedAttribute{
											MarkdownDescription: "None\n> ℹ️ **Note:** You must specify exactly one of `autokey`, `none`, and `symmetric_key`.",
											Computed:            true,
											Attributes:          map[string]dsschema.Attribute{},
										},
										"symmetric_key": dsschema.SingleNestedAttribute{
											MarkdownDescription: "Symmetric key\n> ℹ️ **Note:** You must specify exactly one of `autokey`, `none`, and `symmetric_key`.",
											Computed:            true,
											Attributes: map[string]dsschema.Attribute{
												"algorithm": dsschema.SingleNestedAttribute{
													MarkdownDescription: "Algorithm",
													Computed:            true,
													Attributes: map[string]dsschema.Attribute{
														"md5": dsschema.SingleNestedAttribute{
															MarkdownDescription: "Md5",
															Computed:            true,
															Attributes: map[string]dsschema.Attribute{
																"authentication_key": dsschema.StringAttribute{
																	MarkdownDescription: "Authentication key",
																	Computed:            true,
																},
															},
														},
														"sha1": dsschema.SingleNestedAttribute{
															MarkdownDescription: "Sha1",
															Computed:            true,
															Attributes: map[string]dsschema.Attribute{
																"authentication_key": dsschema.StringAttribute{
																	MarkdownDescription: "Authentication key",
																	Computed:            true,
																},
															},
														},
													},
												},
												"key_id": dsschema.Float64Attribute{
													MarkdownDescription: "Key id",
													Computed:            true,
												},
											},
										},
									},
								},
								"ntp_server_address": dsschema.StringAttribute{
									MarkdownDescription: "Ntp server address",
									Computed:            true,
								},
							},
						},
						"secondary_ntp_server": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Secondary ntp server",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"authentication_type": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Authentication type",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"autokey": dsschema.SingleNestedAttribute{
											MarkdownDescription: "Autokey\n> ℹ️ **Note:** You must specify exactly one of `autokey`, `none`, and `symmetric_key`.",
											Computed:            true,
											Attributes:          map[string]dsschema.Attribute{},
										},
										"none": dsschema.SingleNestedAttribute{
											MarkdownDescription: "None\n> ℹ️ **Note:** You must specify exactly one of `autokey`, `none`, and `symmetric_key`.",
											Computed:            true,
											Attributes:          map[string]dsschema.Attribute{},
										},
										"symmetric_key": dsschema.SingleNestedAttribute{
											MarkdownDescription: "Symmetric key\n> ℹ️ **Note:** You must specify exactly one of `autokey`, `none`, and `symmetric_key`.",
											Computed:            true,
											Attributes: map[string]dsschema.Attribute{
												"algorithm": dsschema.SingleNestedAttribute{
													MarkdownDescription: "Algorithm",
													Computed:            true,
													Attributes: map[string]dsschema.Attribute{
														"md5": dsschema.SingleNestedAttribute{
															MarkdownDescription: "Md5",
															Computed:            true,
															Attributes: map[string]dsschema.Attribute{
																"authentication_key": dsschema.StringAttribute{
																	MarkdownDescription: "Authentication key",
																	Computed:            true,
																},
															},
														},
														"sha1": dsschema.SingleNestedAttribute{
															MarkdownDescription: "Sha1",
															Computed:            true,
															Attributes: map[string]dsschema.Attribute{
																"authentication_key": dsschema.StringAttribute{
																	MarkdownDescription: "Authentication key",
																	Computed:            true,
																},
															},
														},
													},
												},
												"key_id": dsschema.Float64Attribute{
													MarkdownDescription: "Key id",
													Computed:            true,
												},
											},
										},
									},
								},
								"ntp_server_address": dsschema.StringAttribute{
									MarkdownDescription: "Ntp server address",
									Computed:            true,
								},
							},
						},
					},
				},
				"secure_proxy_password": dsschema.StringAttribute{
					MarkdownDescription: "Secure proxy password",
					Computed:            true,
					Sensitive:           true,
				},
				"secure_proxy_port": dsschema.Float64Attribute{
					MarkdownDescription: "Secure proxy port",
					Computed:            true,
				},
				"secure_proxy_server": dsschema.StringAttribute{
					MarkdownDescription: "Secure proxy server",
					Computed:            true,
				},
				"secure_proxy_user": dsschema.StringAttribute{
					MarkdownDescription: "Secure proxy user",
					Computed:            true,
				},
				"server_verification": dsschema.BoolAttribute{
					MarkdownDescription: "Server verification",
					Computed:            true,
				},
				"update_server": dsschema.StringAttribute{
					MarkdownDescription: "Update server",
					Computed:            true,
				},
			},
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
	},
}

// ServiceSettingsListModel represents the data model for a list data source.
type ServiceSettingsListModel struct {
	Tfid    types.String      `tfsdk:"tfid"`
	Data    []ServiceSettings `tfsdk:"data"`
	Limit   types.Int64       `tfsdk:"limit"`
	Offset  types.Int64       `tfsdk:"offset"`
	Name    types.String      `tfsdk:"name"`
	Total   types.Int64       `tfsdk:"total"`
	Folder  types.String      `tfsdk:"folder"`
	Snippet types.String      `tfsdk:"snippet"`
	Device  types.String      `tfsdk:"device"`
}

// ServiceSettingsListDataSourceSchema defines the schema for a list data source.
var ServiceSettingsListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: ServiceSettingsDataSourceSchema.Attributes,
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
