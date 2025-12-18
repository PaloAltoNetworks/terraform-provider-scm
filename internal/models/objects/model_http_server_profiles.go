package models

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	dsschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Package: objects
// This file contains models for the objects SDK package

// HttpServerProfiles represents the Terraform model for HttpServerProfiles
type HttpServerProfiles struct {
	Tfid            types.String          `tfsdk:"tfid"`
	Device          basetypes.StringValue `tfsdk:"device"`
	Folder          basetypes.StringValue `tfsdk:"folder"`
	Format          basetypes.ObjectValue `tfsdk:"format"`
	Id              basetypes.StringValue `tfsdk:"id"`
	Name            basetypes.StringValue `tfsdk:"name"`
	Server          basetypes.ListValue   `tfsdk:"server"`
	Snippet         basetypes.StringValue `tfsdk:"snippet"`
	TagRegistration basetypes.BoolValue   `tfsdk:"tag_registration"`
}

// HttpServerProfilesFormat represents a nested structure within the HttpServerProfiles model
type HttpServerProfilesFormat struct {
	Auth          basetypes.ObjectValue `tfsdk:"auth"`
	Config        basetypes.ObjectValue `tfsdk:"config"`
	Correlation   basetypes.ObjectValue `tfsdk:"correlation"`
	Data          basetypes.ObjectValue `tfsdk:"data"`
	Decryption    basetypes.ObjectValue `tfsdk:"decryption"`
	Globalprotect basetypes.ObjectValue `tfsdk:"globalprotect"`
	Gtp           basetypes.ObjectValue `tfsdk:"gtp"`
	HipMatch      basetypes.ObjectValue `tfsdk:"hip_match"`
	Iptag         basetypes.ObjectValue `tfsdk:"iptag"`
	Sctp          basetypes.ObjectValue `tfsdk:"sctp"`
	System        basetypes.ObjectValue `tfsdk:"system"`
	Threat        basetypes.ObjectValue `tfsdk:"threat"`
	Traffic       basetypes.ObjectValue `tfsdk:"traffic"`
	Tunnel        basetypes.ObjectValue `tfsdk:"tunnel"`
	Url           basetypes.ObjectValue `tfsdk:"url"`
	Userid        basetypes.ObjectValue `tfsdk:"userid"`
	Wildfire      basetypes.ObjectValue `tfsdk:"wildfire"`
}

// PayloadFormat represents a nested structure within the HttpServerProfiles model
type PayloadFormat struct {
	Headers   basetypes.ListValue   `tfsdk:"headers"`
	Name      basetypes.StringValue `tfsdk:"name"`
	Params    basetypes.ListValue   `tfsdk:"params"`
	Payload   basetypes.StringValue `tfsdk:"payload"`
	UrlFormat basetypes.StringValue `tfsdk:"url_format"`
}

// PayloadFormatHeadersInner represents a nested structure within the HttpServerProfiles model
type PayloadFormatHeadersInner struct {
	Name  basetypes.StringValue `tfsdk:"name"`
	Value basetypes.StringValue `tfsdk:"value"`
}

// PayloadFormatParamsInner represents a nested structure within the HttpServerProfiles model
type PayloadFormatParamsInner struct {
	Name  basetypes.StringValue `tfsdk:"name"`
	Value basetypes.StringValue `tfsdk:"value"`
}

// HttpServerProfilesServerInner represents a nested structure within the HttpServerProfiles model
type HttpServerProfilesServerInner struct {
	Address            basetypes.StringValue `tfsdk:"address"`
	CertificateProfile basetypes.StringValue `tfsdk:"certificate_profile"`
	HttpMethod         basetypes.StringValue `tfsdk:"http_method"`
	Name               basetypes.StringValue `tfsdk:"name"`
	Port               basetypes.Int64Value  `tfsdk:"port"`
	Protocol           basetypes.StringValue `tfsdk:"protocol"`
	TlsVersion         basetypes.StringValue `tfsdk:"tls_version"`
}

// AttrTypes defines the attribute types for the HttpServerProfiles model.
func (o HttpServerProfiles) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":   basetypes.StringType{},
		"device": basetypes.StringType{},
		"folder": basetypes.StringType{},
		"format": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"auth": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"headers": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"name":  basetypes.StringType{},
								"value": basetypes.StringType{},
							},
						}},
						"name": basetypes.StringType{},
						"params": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"name":  basetypes.StringType{},
								"value": basetypes.StringType{},
							},
						}},
						"payload":    basetypes.StringType{},
						"url_format": basetypes.StringType{},
					},
				},
				"config": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"headers": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"name":  basetypes.StringType{},
								"value": basetypes.StringType{},
							},
						}},
						"name": basetypes.StringType{},
						"params": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"name":  basetypes.StringType{},
								"value": basetypes.StringType{},
							},
						}},
						"payload":    basetypes.StringType{},
						"url_format": basetypes.StringType{},
					},
				},
				"correlation": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"headers": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"name":  basetypes.StringType{},
								"value": basetypes.StringType{},
							},
						}},
						"name": basetypes.StringType{},
						"params": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"name":  basetypes.StringType{},
								"value": basetypes.StringType{},
							},
						}},
						"payload":    basetypes.StringType{},
						"url_format": basetypes.StringType{},
					},
				},
				"data": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"headers": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"name":  basetypes.StringType{},
								"value": basetypes.StringType{},
							},
						}},
						"name": basetypes.StringType{},
						"params": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"name":  basetypes.StringType{},
								"value": basetypes.StringType{},
							},
						}},
						"payload":    basetypes.StringType{},
						"url_format": basetypes.StringType{},
					},
				},
				"decryption": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"headers": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"name":  basetypes.StringType{},
								"value": basetypes.StringType{},
							},
						}},
						"name": basetypes.StringType{},
						"params": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"name":  basetypes.StringType{},
								"value": basetypes.StringType{},
							},
						}},
						"payload":    basetypes.StringType{},
						"url_format": basetypes.StringType{},
					},
				},
				"globalprotect": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"headers": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"name":  basetypes.StringType{},
								"value": basetypes.StringType{},
							},
						}},
						"name": basetypes.StringType{},
						"params": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"name":  basetypes.StringType{},
								"value": basetypes.StringType{},
							},
						}},
						"payload":    basetypes.StringType{},
						"url_format": basetypes.StringType{},
					},
				},
				"gtp": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"headers": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"name":  basetypes.StringType{},
								"value": basetypes.StringType{},
							},
						}},
						"name": basetypes.StringType{},
						"params": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"name":  basetypes.StringType{},
								"value": basetypes.StringType{},
							},
						}},
						"payload":    basetypes.StringType{},
						"url_format": basetypes.StringType{},
					},
				},
				"hip_match": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"headers": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"name":  basetypes.StringType{},
								"value": basetypes.StringType{},
							},
						}},
						"name": basetypes.StringType{},
						"params": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"name":  basetypes.StringType{},
								"value": basetypes.StringType{},
							},
						}},
						"payload":    basetypes.StringType{},
						"url_format": basetypes.StringType{},
					},
				},
				"iptag": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"headers": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"name":  basetypes.StringType{},
								"value": basetypes.StringType{},
							},
						}},
						"name": basetypes.StringType{},
						"params": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"name":  basetypes.StringType{},
								"value": basetypes.StringType{},
							},
						}},
						"payload":    basetypes.StringType{},
						"url_format": basetypes.StringType{},
					},
				},
				"sctp": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"headers": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"name":  basetypes.StringType{},
								"value": basetypes.StringType{},
							},
						}},
						"name": basetypes.StringType{},
						"params": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"name":  basetypes.StringType{},
								"value": basetypes.StringType{},
							},
						}},
						"payload":    basetypes.StringType{},
						"url_format": basetypes.StringType{},
					},
				},
				"system": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"headers": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"name":  basetypes.StringType{},
								"value": basetypes.StringType{},
							},
						}},
						"name": basetypes.StringType{},
						"params": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"name":  basetypes.StringType{},
								"value": basetypes.StringType{},
							},
						}},
						"payload":    basetypes.StringType{},
						"url_format": basetypes.StringType{},
					},
				},
				"threat": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"headers": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"name":  basetypes.StringType{},
								"value": basetypes.StringType{},
							},
						}},
						"name": basetypes.StringType{},
						"params": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"name":  basetypes.StringType{},
								"value": basetypes.StringType{},
							},
						}},
						"payload":    basetypes.StringType{},
						"url_format": basetypes.StringType{},
					},
				},
				"traffic": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"headers": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"name":  basetypes.StringType{},
								"value": basetypes.StringType{},
							},
						}},
						"name": basetypes.StringType{},
						"params": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"name":  basetypes.StringType{},
								"value": basetypes.StringType{},
							},
						}},
						"payload":    basetypes.StringType{},
						"url_format": basetypes.StringType{},
					},
				},
				"tunnel": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"headers": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"name":  basetypes.StringType{},
								"value": basetypes.StringType{},
							},
						}},
						"name": basetypes.StringType{},
						"params": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"name":  basetypes.StringType{},
								"value": basetypes.StringType{},
							},
						}},
						"payload":    basetypes.StringType{},
						"url_format": basetypes.StringType{},
					},
				},
				"url": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"headers": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"name":  basetypes.StringType{},
								"value": basetypes.StringType{},
							},
						}},
						"name": basetypes.StringType{},
						"params": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"name":  basetypes.StringType{},
								"value": basetypes.StringType{},
							},
						}},
						"payload":    basetypes.StringType{},
						"url_format": basetypes.StringType{},
					},
				},
				"userid": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"headers": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"name":  basetypes.StringType{},
								"value": basetypes.StringType{},
							},
						}},
						"name": basetypes.StringType{},
						"params": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"name":  basetypes.StringType{},
								"value": basetypes.StringType{},
							},
						}},
						"payload":    basetypes.StringType{},
						"url_format": basetypes.StringType{},
					},
				},
				"wildfire": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"headers": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"name":  basetypes.StringType{},
								"value": basetypes.StringType{},
							},
						}},
						"name": basetypes.StringType{},
						"params": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"name":  basetypes.StringType{},
								"value": basetypes.StringType{},
							},
						}},
						"payload":    basetypes.StringType{},
						"url_format": basetypes.StringType{},
					},
				},
			},
		},
		"id":   basetypes.StringType{},
		"name": basetypes.StringType{},
		"server": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"address":             basetypes.StringType{},
				"certificate_profile": basetypes.StringType{},
				"http_method":         basetypes.StringType{},
				"name":                basetypes.StringType{},
				"port":                basetypes.Int64Type{},
				"protocol":            basetypes.StringType{},
				"tls_version":         basetypes.StringType{},
			},
		}},
		"snippet":          basetypes.StringType{},
		"tag_registration": basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of HttpServerProfiles objects.
func (o HttpServerProfiles) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the HttpServerProfilesFormat model.
func (o HttpServerProfilesFormat) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"auth": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"headers": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"name":  basetypes.StringType{},
						"value": basetypes.StringType{},
					},
				}},
				"name": basetypes.StringType{},
				"params": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"name":  basetypes.StringType{},
						"value": basetypes.StringType{},
					},
				}},
				"payload":    basetypes.StringType{},
				"url_format": basetypes.StringType{},
			},
		},
		"config": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"headers": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"name":  basetypes.StringType{},
						"value": basetypes.StringType{},
					},
				}},
				"name": basetypes.StringType{},
				"params": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"name":  basetypes.StringType{},
						"value": basetypes.StringType{},
					},
				}},
				"payload":    basetypes.StringType{},
				"url_format": basetypes.StringType{},
			},
		},
		"correlation": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"headers": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"name":  basetypes.StringType{},
						"value": basetypes.StringType{},
					},
				}},
				"name": basetypes.StringType{},
				"params": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"name":  basetypes.StringType{},
						"value": basetypes.StringType{},
					},
				}},
				"payload":    basetypes.StringType{},
				"url_format": basetypes.StringType{},
			},
		},
		"data": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"headers": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"name":  basetypes.StringType{},
						"value": basetypes.StringType{},
					},
				}},
				"name": basetypes.StringType{},
				"params": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"name":  basetypes.StringType{},
						"value": basetypes.StringType{},
					},
				}},
				"payload":    basetypes.StringType{},
				"url_format": basetypes.StringType{},
			},
		},
		"decryption": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"headers": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"name":  basetypes.StringType{},
						"value": basetypes.StringType{},
					},
				}},
				"name": basetypes.StringType{},
				"params": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"name":  basetypes.StringType{},
						"value": basetypes.StringType{},
					},
				}},
				"payload":    basetypes.StringType{},
				"url_format": basetypes.StringType{},
			},
		},
		"globalprotect": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"headers": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"name":  basetypes.StringType{},
						"value": basetypes.StringType{},
					},
				}},
				"name": basetypes.StringType{},
				"params": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"name":  basetypes.StringType{},
						"value": basetypes.StringType{},
					},
				}},
				"payload":    basetypes.StringType{},
				"url_format": basetypes.StringType{},
			},
		},
		"gtp": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"headers": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"name":  basetypes.StringType{},
						"value": basetypes.StringType{},
					},
				}},
				"name": basetypes.StringType{},
				"params": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"name":  basetypes.StringType{},
						"value": basetypes.StringType{},
					},
				}},
				"payload":    basetypes.StringType{},
				"url_format": basetypes.StringType{},
			},
		},
		"hip_match": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"headers": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"name":  basetypes.StringType{},
						"value": basetypes.StringType{},
					},
				}},
				"name": basetypes.StringType{},
				"params": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"name":  basetypes.StringType{},
						"value": basetypes.StringType{},
					},
				}},
				"payload":    basetypes.StringType{},
				"url_format": basetypes.StringType{},
			},
		},
		"iptag": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"headers": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"name":  basetypes.StringType{},
						"value": basetypes.StringType{},
					},
				}},
				"name": basetypes.StringType{},
				"params": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"name":  basetypes.StringType{},
						"value": basetypes.StringType{},
					},
				}},
				"payload":    basetypes.StringType{},
				"url_format": basetypes.StringType{},
			},
		},
		"sctp": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"headers": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"name":  basetypes.StringType{},
						"value": basetypes.StringType{},
					},
				}},
				"name": basetypes.StringType{},
				"params": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"name":  basetypes.StringType{},
						"value": basetypes.StringType{},
					},
				}},
				"payload":    basetypes.StringType{},
				"url_format": basetypes.StringType{},
			},
		},
		"system": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"headers": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"name":  basetypes.StringType{},
						"value": basetypes.StringType{},
					},
				}},
				"name": basetypes.StringType{},
				"params": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"name":  basetypes.StringType{},
						"value": basetypes.StringType{},
					},
				}},
				"payload":    basetypes.StringType{},
				"url_format": basetypes.StringType{},
			},
		},
		"threat": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"headers": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"name":  basetypes.StringType{},
						"value": basetypes.StringType{},
					},
				}},
				"name": basetypes.StringType{},
				"params": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"name":  basetypes.StringType{},
						"value": basetypes.StringType{},
					},
				}},
				"payload":    basetypes.StringType{},
				"url_format": basetypes.StringType{},
			},
		},
		"traffic": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"headers": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"name":  basetypes.StringType{},
						"value": basetypes.StringType{},
					},
				}},
				"name": basetypes.StringType{},
				"params": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"name":  basetypes.StringType{},
						"value": basetypes.StringType{},
					},
				}},
				"payload":    basetypes.StringType{},
				"url_format": basetypes.StringType{},
			},
		},
		"tunnel": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"headers": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"name":  basetypes.StringType{},
						"value": basetypes.StringType{},
					},
				}},
				"name": basetypes.StringType{},
				"params": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"name":  basetypes.StringType{},
						"value": basetypes.StringType{},
					},
				}},
				"payload":    basetypes.StringType{},
				"url_format": basetypes.StringType{},
			},
		},
		"url": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"headers": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"name":  basetypes.StringType{},
						"value": basetypes.StringType{},
					},
				}},
				"name": basetypes.StringType{},
				"params": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"name":  basetypes.StringType{},
						"value": basetypes.StringType{},
					},
				}},
				"payload":    basetypes.StringType{},
				"url_format": basetypes.StringType{},
			},
		},
		"userid": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"headers": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"name":  basetypes.StringType{},
						"value": basetypes.StringType{},
					},
				}},
				"name": basetypes.StringType{},
				"params": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"name":  basetypes.StringType{},
						"value": basetypes.StringType{},
					},
				}},
				"payload":    basetypes.StringType{},
				"url_format": basetypes.StringType{},
			},
		},
		"wildfire": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"headers": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"name":  basetypes.StringType{},
						"value": basetypes.StringType{},
					},
				}},
				"name": basetypes.StringType{},
				"params": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"name":  basetypes.StringType{},
						"value": basetypes.StringType{},
					},
				}},
				"payload":    basetypes.StringType{},
				"url_format": basetypes.StringType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of HttpServerProfilesFormat objects.
func (o HttpServerProfilesFormat) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the PayloadFormat model.
func (o PayloadFormat) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"headers": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"name":  basetypes.StringType{},
				"value": basetypes.StringType{},
			},
		}},
		"name": basetypes.StringType{},
		"params": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"name":  basetypes.StringType{},
				"value": basetypes.StringType{},
			},
		}},
		"payload":    basetypes.StringType{},
		"url_format": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of PayloadFormat objects.
func (o PayloadFormat) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the PayloadFormatHeadersInner model.
func (o PayloadFormatHeadersInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name":  basetypes.StringType{},
		"value": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of PayloadFormatHeadersInner objects.
func (o PayloadFormatHeadersInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the PayloadFormatParamsInner model.
func (o PayloadFormatParamsInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name":  basetypes.StringType{},
		"value": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of PayloadFormatParamsInner objects.
func (o PayloadFormatParamsInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the HttpServerProfilesServerInner model.
func (o HttpServerProfilesServerInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"address":             basetypes.StringType{},
		"certificate_profile": basetypes.StringType{},
		"http_method":         basetypes.StringType{},
		"name":                basetypes.StringType{},
		"port":                basetypes.Int64Type{},
		"protocol":            basetypes.StringType{},
		"tls_version":         basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of HttpServerProfilesServerInner objects.
func (o HttpServerProfilesServerInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// HttpServerProfilesResourceSchema defines the schema for HttpServerProfiles resource
var HttpServerProfilesResourceSchema = schema.Schema{
	MarkdownDescription: "HttpServerProfile resource",
	Attributes: map[string]schema.Attribute{
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
		"format": schema.SingleNestedAttribute{
			MarkdownDescription: "Format",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"auth": schema.SingleNestedAttribute{
					MarkdownDescription: "Auth",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"headers": schema.ListNestedAttribute{
							MarkdownDescription: "Headers",
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										MarkdownDescription: "Header name",
										Optional:            true,
									},
									"value": schema.StringAttribute{
										MarkdownDescription: "Header value",
										Optional:            true,
									},
								},
							},
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "The name of the payload format",
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("Default"),
						},
						"params": schema.ListNestedAttribute{
							MarkdownDescription: "Params",
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										MarkdownDescription: "Parameter name",
										Optional:            true,
									},
									"value": schema.StringAttribute{
										MarkdownDescription: "Parameter value",
										Optional:            true,
									},
								},
							},
						},
						"payload": schema.StringAttribute{
							MarkdownDescription: "The log payload format.  The accepted log field values are as follows.\n* `receive_time`\n* `serial`\n* `seqno`\n* `actionflags`\n* `type`\n* `subtype`\n* `time_generated`\n* `high_res_timestamp`\n* `dg_hier_level_1`\n* `dg_hier_level_2`\n* `dg_hier_level_3`\n* `dg_hier_level_4`\n* `vsys_name`\n* `device_name`\n* `vsys_id`\n* `host`\n* `vsys`\n* `cmd`\n* `admin`\n* `client`\n* `result`\n* `path`\n* `dg_id`\n* `comment`\n* `tpl_id`\n* `sender_sw_version`\n* `cef-formatted-receive_time`\n* `cef-formatted-time_generated`\n* `before-change-detail`\n* `after-change-detail`\n",
							Optional:            true,
						},
						"url_format": schema.StringAttribute{
							MarkdownDescription: "The URL path of the HTTP server",
							Optional:            true,
						},
					},
				},
				"config": schema.SingleNestedAttribute{
					MarkdownDescription: "Config",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"headers": schema.ListNestedAttribute{
							MarkdownDescription: "Headers",
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										MarkdownDescription: "Header name",
										Optional:            true,
									},
									"value": schema.StringAttribute{
										MarkdownDescription: "Header value",
										Optional:            true,
									},
								},
							},
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "The name of the payload format",
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("Default"),
						},
						"params": schema.ListNestedAttribute{
							MarkdownDescription: "Params",
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										MarkdownDescription: "Parameter name",
										Optional:            true,
									},
									"value": schema.StringAttribute{
										MarkdownDescription: "Parameter value",
										Optional:            true,
									},
								},
							},
						},
						"payload": schema.StringAttribute{
							MarkdownDescription: "The log payload format.  The accepted log field values are as follows.\n* `receive_time`\n* `serial`\n* `seqno`\n* `actionflags`\n* `type`\n* `subtype`\n* `time_generated`\n* `high_res_timestamp`\n* `dg_hier_level_1`\n* `dg_hier_level_2`\n* `dg_hier_level_3`\n* `dg_hier_level_4`\n* `vsys_name`\n* `device_name`\n* `vsys_id`\n* `host`\n* `vsys`\n* `cmd`\n* `admin`\n* `client`\n* `result`\n* `path`\n* `dg_id`\n* `comment`\n* `tpl_id`\n* `sender_sw_version`\n* `cef-formatted-receive_time`\n* `cef-formatted-time_generated`\n* `before-change-detail`\n* `after-change-detail`\n",
							Optional:            true,
						},
						"url_format": schema.StringAttribute{
							MarkdownDescription: "The URL path of the HTTP server",
							Optional:            true,
						},
					},
				},
				"correlation": schema.SingleNestedAttribute{
					MarkdownDescription: "Correlation",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"headers": schema.ListNestedAttribute{
							MarkdownDescription: "Headers",
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										MarkdownDescription: "Header name",
										Optional:            true,
									},
									"value": schema.StringAttribute{
										MarkdownDescription: "Header value",
										Optional:            true,
									},
								},
							},
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "The name of the payload format",
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("Default"),
						},
						"params": schema.ListNestedAttribute{
							MarkdownDescription: "Params",
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										MarkdownDescription: "Parameter name",
										Optional:            true,
									},
									"value": schema.StringAttribute{
										MarkdownDescription: "Parameter value",
										Optional:            true,
									},
								},
							},
						},
						"payload": schema.StringAttribute{
							MarkdownDescription: "The log payload format.  The accepted log field values are as follows.\n* `receive_time`\n* `serial`\n* `seqno`\n* `actionflags`\n* `type`\n* `subtype`\n* `time_generated`\n* `high_res_timestamp`\n* `dg_hier_level_1`\n* `dg_hier_level_2`\n* `dg_hier_level_3`\n* `dg_hier_level_4`\n* `vsys_name`\n* `device_name`\n* `vsys_id`\n* `host`\n* `vsys`\n* `cmd`\n* `admin`\n* `client`\n* `result`\n* `path`\n* `dg_id`\n* `comment`\n* `tpl_id`\n* `sender_sw_version`\n* `cef-formatted-receive_time`\n* `cef-formatted-time_generated`\n* `before-change-detail`\n* `after-change-detail`\n",
							Optional:            true,
						},
						"url_format": schema.StringAttribute{
							MarkdownDescription: "The URL path of the HTTP server",
							Optional:            true,
						},
					},
				},
				"data": schema.SingleNestedAttribute{
					MarkdownDescription: "Data",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"headers": schema.ListNestedAttribute{
							MarkdownDescription: "Headers",
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										MarkdownDescription: "Header name",
										Optional:            true,
									},
									"value": schema.StringAttribute{
										MarkdownDescription: "Header value",
										Optional:            true,
									},
								},
							},
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "The name of the payload format",
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("Default"),
						},
						"params": schema.ListNestedAttribute{
							MarkdownDescription: "Params",
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										MarkdownDescription: "Parameter name",
										Optional:            true,
									},
									"value": schema.StringAttribute{
										MarkdownDescription: "Parameter value",
										Optional:            true,
									},
								},
							},
						},
						"payload": schema.StringAttribute{
							MarkdownDescription: "The log payload format.  The accepted log field values are as follows.\n* `receive_time`\n* `serial`\n* `seqno`\n* `actionflags`\n* `type`\n* `subtype`\n* `time_generated`\n* `high_res_timestamp`\n* `dg_hier_level_1`\n* `dg_hier_level_2`\n* `dg_hier_level_3`\n* `dg_hier_level_4`\n* `vsys_name`\n* `device_name`\n* `vsys_id`\n* `host`\n* `vsys`\n* `cmd`\n* `admin`\n* `client`\n* `result`\n* `path`\n* `dg_id`\n* `comment`\n* `tpl_id`\n* `sender_sw_version`\n* `cef-formatted-receive_time`\n* `cef-formatted-time_generated`\n* `before-change-detail`\n* `after-change-detail`\n",
							Optional:            true,
						},
						"url_format": schema.StringAttribute{
							MarkdownDescription: "The URL path of the HTTP server",
							Optional:            true,
						},
					},
				},
				"decryption": schema.SingleNestedAttribute{
					MarkdownDescription: "Decryption",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"headers": schema.ListNestedAttribute{
							MarkdownDescription: "Headers",
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										MarkdownDescription: "Header name",
										Optional:            true,
									},
									"value": schema.StringAttribute{
										MarkdownDescription: "Header value",
										Optional:            true,
									},
								},
							},
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "The name of the payload format",
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("Default"),
						},
						"params": schema.ListNestedAttribute{
							MarkdownDescription: "Params",
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										MarkdownDescription: "Parameter name",
										Optional:            true,
									},
									"value": schema.StringAttribute{
										MarkdownDescription: "Parameter value",
										Optional:            true,
									},
								},
							},
						},
						"payload": schema.StringAttribute{
							MarkdownDescription: "The log payload format.  The accepted log field values are as follows.\n* `receive_time`\n* `serial`\n* `seqno`\n* `actionflags`\n* `type`\n* `subtype`\n* `time_generated`\n* `high_res_timestamp`\n* `dg_hier_level_1`\n* `dg_hier_level_2`\n* `dg_hier_level_3`\n* `dg_hier_level_4`\n* `vsys_name`\n* `device_name`\n* `vsys_id`\n* `host`\n* `vsys`\n* `cmd`\n* `admin`\n* `client`\n* `result`\n* `path`\n* `dg_id`\n* `comment`\n* `tpl_id`\n* `sender_sw_version`\n* `cef-formatted-receive_time`\n* `cef-formatted-time_generated`\n* `before-change-detail`\n* `after-change-detail`\n",
							Optional:            true,
						},
						"url_format": schema.StringAttribute{
							MarkdownDescription: "The URL path of the HTTP server",
							Optional:            true,
						},
					},
				},
				"globalprotect": schema.SingleNestedAttribute{
					MarkdownDescription: "Globalprotect",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"headers": schema.ListNestedAttribute{
							MarkdownDescription: "Headers",
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										MarkdownDescription: "Header name",
										Optional:            true,
									},
									"value": schema.StringAttribute{
										MarkdownDescription: "Header value",
										Optional:            true,
									},
								},
							},
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "The name of the payload format",
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("Default"),
						},
						"params": schema.ListNestedAttribute{
							MarkdownDescription: "Params",
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										MarkdownDescription: "Parameter name",
										Optional:            true,
									},
									"value": schema.StringAttribute{
										MarkdownDescription: "Parameter value",
										Optional:            true,
									},
								},
							},
						},
						"payload": schema.StringAttribute{
							MarkdownDescription: "The log payload format.  The accepted log field values are as follows.\n* `receive_time`\n* `serial`\n* `seqno`\n* `actionflags`\n* `type`\n* `subtype`\n* `time_generated`\n* `high_res_timestamp`\n* `dg_hier_level_1`\n* `dg_hier_level_2`\n* `dg_hier_level_3`\n* `dg_hier_level_4`\n* `vsys_name`\n* `device_name`\n* `vsys_id`\n* `host`\n* `vsys`\n* `cmd`\n* `admin`\n* `client`\n* `result`\n* `path`\n* `dg_id`\n* `comment`\n* `tpl_id`\n* `sender_sw_version`\n* `cef-formatted-receive_time`\n* `cef-formatted-time_generated`\n* `before-change-detail`\n* `after-change-detail`\n",
							Optional:            true,
						},
						"url_format": schema.StringAttribute{
							MarkdownDescription: "The URL path of the HTTP server",
							Optional:            true,
						},
					},
				},
				"gtp": schema.SingleNestedAttribute{
					MarkdownDescription: "Gtp",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"headers": schema.ListNestedAttribute{
							MarkdownDescription: "Headers",
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										MarkdownDescription: "Header name",
										Optional:            true,
									},
									"value": schema.StringAttribute{
										MarkdownDescription: "Header value",
										Optional:            true,
									},
								},
							},
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "The name of the payload format",
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("Default"),
						},
						"params": schema.ListNestedAttribute{
							MarkdownDescription: "Params",
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										MarkdownDescription: "Parameter name",
										Optional:            true,
									},
									"value": schema.StringAttribute{
										MarkdownDescription: "Parameter value",
										Optional:            true,
									},
								},
							},
						},
						"payload": schema.StringAttribute{
							MarkdownDescription: "The log payload format.  The accepted log field values are as follows.\n* `receive_time`\n* `serial`\n* `seqno`\n* `actionflags`\n* `type`\n* `subtype`\n* `time_generated`\n* `high_res_timestamp`\n* `dg_hier_level_1`\n* `dg_hier_level_2`\n* `dg_hier_level_3`\n* `dg_hier_level_4`\n* `vsys_name`\n* `device_name`\n* `vsys_id`\n* `host`\n* `vsys`\n* `cmd`\n* `admin`\n* `client`\n* `result`\n* `path`\n* `dg_id`\n* `comment`\n* `tpl_id`\n* `sender_sw_version`\n* `cef-formatted-receive_time`\n* `cef-formatted-time_generated`\n* `before-change-detail`\n* `after-change-detail`\n",
							Optional:            true,
						},
						"url_format": schema.StringAttribute{
							MarkdownDescription: "The URL path of the HTTP server",
							Optional:            true,
						},
					},
				},
				"hip_match": schema.SingleNestedAttribute{
					MarkdownDescription: "Hip match",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"headers": schema.ListNestedAttribute{
							MarkdownDescription: "Headers",
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										MarkdownDescription: "Header name",
										Optional:            true,
									},
									"value": schema.StringAttribute{
										MarkdownDescription: "Header value",
										Optional:            true,
									},
								},
							},
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "The name of the payload format",
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("Default"),
						},
						"params": schema.ListNestedAttribute{
							MarkdownDescription: "Params",
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										MarkdownDescription: "Parameter name",
										Optional:            true,
									},
									"value": schema.StringAttribute{
										MarkdownDescription: "Parameter value",
										Optional:            true,
									},
								},
							},
						},
						"payload": schema.StringAttribute{
							MarkdownDescription: "The log payload format.  The accepted log field values are as follows.\n* `receive_time`\n* `serial`\n* `seqno`\n* `actionflags`\n* `type`\n* `subtype`\n* `time_generated`\n* `high_res_timestamp`\n* `dg_hier_level_1`\n* `dg_hier_level_2`\n* `dg_hier_level_3`\n* `dg_hier_level_4`\n* `vsys_name`\n* `device_name`\n* `vsys_id`\n* `host`\n* `vsys`\n* `cmd`\n* `admin`\n* `client`\n* `result`\n* `path`\n* `dg_id`\n* `comment`\n* `tpl_id`\n* `sender_sw_version`\n* `cef-formatted-receive_time`\n* `cef-formatted-time_generated`\n* `before-change-detail`\n* `after-change-detail`\n",
							Optional:            true,
						},
						"url_format": schema.StringAttribute{
							MarkdownDescription: "The URL path of the HTTP server",
							Optional:            true,
						},
					},
				},
				"iptag": schema.SingleNestedAttribute{
					MarkdownDescription: "Iptag",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"headers": schema.ListNestedAttribute{
							MarkdownDescription: "Headers",
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										MarkdownDescription: "Header name",
										Optional:            true,
									},
									"value": schema.StringAttribute{
										MarkdownDescription: "Header value",
										Optional:            true,
									},
								},
							},
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "The name of the payload format",
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("Default"),
						},
						"params": schema.ListNestedAttribute{
							MarkdownDescription: "Params",
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										MarkdownDescription: "Parameter name",
										Optional:            true,
									},
									"value": schema.StringAttribute{
										MarkdownDescription: "Parameter value",
										Optional:            true,
									},
								},
							},
						},
						"payload": schema.StringAttribute{
							MarkdownDescription: "The log payload format.  The accepted log field values are as follows.\n* `receive_time`\n* `serial`\n* `seqno`\n* `actionflags`\n* `type`\n* `subtype`\n* `time_generated`\n* `high_res_timestamp`\n* `dg_hier_level_1`\n* `dg_hier_level_2`\n* `dg_hier_level_3`\n* `dg_hier_level_4`\n* `vsys_name`\n* `device_name`\n* `vsys_id`\n* `host`\n* `vsys`\n* `cmd`\n* `admin`\n* `client`\n* `result`\n* `path`\n* `dg_id`\n* `comment`\n* `tpl_id`\n* `sender_sw_version`\n* `cef-formatted-receive_time`\n* `cef-formatted-time_generated`\n* `before-change-detail`\n* `after-change-detail`\n",
							Optional:            true,
						},
						"url_format": schema.StringAttribute{
							MarkdownDescription: "The URL path of the HTTP server",
							Optional:            true,
						},
					},
				},
				"sctp": schema.SingleNestedAttribute{
					MarkdownDescription: "Sctp",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"headers": schema.ListNestedAttribute{
							MarkdownDescription: "Headers",
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										MarkdownDescription: "Header name",
										Optional:            true,
									},
									"value": schema.StringAttribute{
										MarkdownDescription: "Header value",
										Optional:            true,
									},
								},
							},
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "The name of the payload format",
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("Default"),
						},
						"params": schema.ListNestedAttribute{
							MarkdownDescription: "Params",
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										MarkdownDescription: "Parameter name",
										Optional:            true,
									},
									"value": schema.StringAttribute{
										MarkdownDescription: "Parameter value",
										Optional:            true,
									},
								},
							},
						},
						"payload": schema.StringAttribute{
							MarkdownDescription: "The log payload format.  The accepted log field values are as follows.\n* `receive_time`\n* `serial`\n* `seqno`\n* `actionflags`\n* `type`\n* `subtype`\n* `time_generated`\n* `high_res_timestamp`\n* `dg_hier_level_1`\n* `dg_hier_level_2`\n* `dg_hier_level_3`\n* `dg_hier_level_4`\n* `vsys_name`\n* `device_name`\n* `vsys_id`\n* `host`\n* `vsys`\n* `cmd`\n* `admin`\n* `client`\n* `result`\n* `path`\n* `dg_id`\n* `comment`\n* `tpl_id`\n* `sender_sw_version`\n* `cef-formatted-receive_time`\n* `cef-formatted-time_generated`\n* `before-change-detail`\n* `after-change-detail`\n",
							Optional:            true,
						},
						"url_format": schema.StringAttribute{
							MarkdownDescription: "The URL path of the HTTP server",
							Optional:            true,
						},
					},
				},
				"system": schema.SingleNestedAttribute{
					MarkdownDescription: "System",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"headers": schema.ListNestedAttribute{
							MarkdownDescription: "Headers",
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										MarkdownDescription: "Header name",
										Optional:            true,
									},
									"value": schema.StringAttribute{
										MarkdownDescription: "Header value",
										Optional:            true,
									},
								},
							},
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "The name of the payload format",
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("Default"),
						},
						"params": schema.ListNestedAttribute{
							MarkdownDescription: "Params",
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										MarkdownDescription: "Parameter name",
										Optional:            true,
									},
									"value": schema.StringAttribute{
										MarkdownDescription: "Parameter value",
										Optional:            true,
									},
								},
							},
						},
						"payload": schema.StringAttribute{
							MarkdownDescription: "The log payload format.  The accepted log field values are as follows.\n* `receive_time`\n* `serial`\n* `seqno`\n* `actionflags`\n* `type`\n* `subtype`\n* `time_generated`\n* `high_res_timestamp`\n* `dg_hier_level_1`\n* `dg_hier_level_2`\n* `dg_hier_level_3`\n* `dg_hier_level_4`\n* `vsys_name`\n* `device_name`\n* `vsys_id`\n* `host`\n* `vsys`\n* `cmd`\n* `admin`\n* `client`\n* `result`\n* `path`\n* `dg_id`\n* `comment`\n* `tpl_id`\n* `sender_sw_version`\n* `cef-formatted-receive_time`\n* `cef-formatted-time_generated`\n* `before-change-detail`\n* `after-change-detail`\n",
							Optional:            true,
						},
						"url_format": schema.StringAttribute{
							MarkdownDescription: "The URL path of the HTTP server",
							Optional:            true,
						},
					},
				},
				"threat": schema.SingleNestedAttribute{
					MarkdownDescription: "Threat",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"headers": schema.ListNestedAttribute{
							MarkdownDescription: "Headers",
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										MarkdownDescription: "Header name",
										Optional:            true,
									},
									"value": schema.StringAttribute{
										MarkdownDescription: "Header value",
										Optional:            true,
									},
								},
							},
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "The name of the payload format",
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("Default"),
						},
						"params": schema.ListNestedAttribute{
							MarkdownDescription: "Params",
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										MarkdownDescription: "Parameter name",
										Optional:            true,
									},
									"value": schema.StringAttribute{
										MarkdownDescription: "Parameter value",
										Optional:            true,
									},
								},
							},
						},
						"payload": schema.StringAttribute{
							MarkdownDescription: "The log payload format.  The accepted log field values are as follows.\n* `receive_time`\n* `serial`\n* `seqno`\n* `actionflags`\n* `type`\n* `subtype`\n* `time_generated`\n* `high_res_timestamp`\n* `dg_hier_level_1`\n* `dg_hier_level_2`\n* `dg_hier_level_3`\n* `dg_hier_level_4`\n* `vsys_name`\n* `device_name`\n* `vsys_id`\n* `host`\n* `vsys`\n* `cmd`\n* `admin`\n* `client`\n* `result`\n* `path`\n* `dg_id`\n* `comment`\n* `tpl_id`\n* `sender_sw_version`\n* `cef-formatted-receive_time`\n* `cef-formatted-time_generated`\n* `before-change-detail`\n* `after-change-detail`\n",
							Optional:            true,
						},
						"url_format": schema.StringAttribute{
							MarkdownDescription: "The URL path of the HTTP server",
							Optional:            true,
						},
					},
				},
				"traffic": schema.SingleNestedAttribute{
					MarkdownDescription: "Traffic",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"headers": schema.ListNestedAttribute{
							MarkdownDescription: "Headers",
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										MarkdownDescription: "Header name",
										Optional:            true,
									},
									"value": schema.StringAttribute{
										MarkdownDescription: "Header value",
										Optional:            true,
									},
								},
							},
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "The name of the payload format",
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("Default"),
						},
						"params": schema.ListNestedAttribute{
							MarkdownDescription: "Params",
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										MarkdownDescription: "Parameter name",
										Optional:            true,
									},
									"value": schema.StringAttribute{
										MarkdownDescription: "Parameter value",
										Optional:            true,
									},
								},
							},
						},
						"payload": schema.StringAttribute{
							MarkdownDescription: "The log payload format.  The accepted log field values are as follows.\n* `receive_time`\n* `serial`\n* `seqno`\n* `actionflags`\n* `type`\n* `subtype`\n* `time_generated`\n* `high_res_timestamp`\n* `dg_hier_level_1`\n* `dg_hier_level_2`\n* `dg_hier_level_3`\n* `dg_hier_level_4`\n* `vsys_name`\n* `device_name`\n* `vsys_id`\n* `host`\n* `vsys`\n* `cmd`\n* `admin`\n* `client`\n* `result`\n* `path`\n* `dg_id`\n* `comment`\n* `tpl_id`\n* `sender_sw_version`\n* `cef-formatted-receive_time`\n* `cef-formatted-time_generated`\n* `before-change-detail`\n* `after-change-detail`\n",
							Optional:            true,
						},
						"url_format": schema.StringAttribute{
							MarkdownDescription: "The URL path of the HTTP server",
							Optional:            true,
						},
					},
				},
				"tunnel": schema.SingleNestedAttribute{
					MarkdownDescription: "Tunnel",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"headers": schema.ListNestedAttribute{
							MarkdownDescription: "Headers",
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										MarkdownDescription: "Header name",
										Optional:            true,
									},
									"value": schema.StringAttribute{
										MarkdownDescription: "Header value",
										Optional:            true,
									},
								},
							},
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "The name of the payload format",
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("Default"),
						},
						"params": schema.ListNestedAttribute{
							MarkdownDescription: "Params",
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										MarkdownDescription: "Parameter name",
										Optional:            true,
									},
									"value": schema.StringAttribute{
										MarkdownDescription: "Parameter value",
										Optional:            true,
									},
								},
							},
						},
						"payload": schema.StringAttribute{
							MarkdownDescription: "The log payload format.  The accepted log field values are as follows.\n* `receive_time`\n* `serial`\n* `seqno`\n* `actionflags`\n* `type`\n* `subtype`\n* `time_generated`\n* `high_res_timestamp`\n* `dg_hier_level_1`\n* `dg_hier_level_2`\n* `dg_hier_level_3`\n* `dg_hier_level_4`\n* `vsys_name`\n* `device_name`\n* `vsys_id`\n* `host`\n* `vsys`\n* `cmd`\n* `admin`\n* `client`\n* `result`\n* `path`\n* `dg_id`\n* `comment`\n* `tpl_id`\n* `sender_sw_version`\n* `cef-formatted-receive_time`\n* `cef-formatted-time_generated`\n* `before-change-detail`\n* `after-change-detail`\n",
							Optional:            true,
						},
						"url_format": schema.StringAttribute{
							MarkdownDescription: "The URL path of the HTTP server",
							Optional:            true,
						},
					},
				},
				"url": schema.SingleNestedAttribute{
					MarkdownDescription: "Url",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"headers": schema.ListNestedAttribute{
							MarkdownDescription: "Headers",
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										MarkdownDescription: "Header name",
										Optional:            true,
									},
									"value": schema.StringAttribute{
										MarkdownDescription: "Header value",
										Optional:            true,
									},
								},
							},
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "The name of the payload format",
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("Default"),
						},
						"params": schema.ListNestedAttribute{
							MarkdownDescription: "Params",
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										MarkdownDescription: "Parameter name",
										Optional:            true,
									},
									"value": schema.StringAttribute{
										MarkdownDescription: "Parameter value",
										Optional:            true,
									},
								},
							},
						},
						"payload": schema.StringAttribute{
							MarkdownDescription: "The log payload format.  The accepted log field values are as follows.\n* `receive_time`\n* `serial`\n* `seqno`\n* `actionflags`\n* `type`\n* `subtype`\n* `time_generated`\n* `high_res_timestamp`\n* `dg_hier_level_1`\n* `dg_hier_level_2`\n* `dg_hier_level_3`\n* `dg_hier_level_4`\n* `vsys_name`\n* `device_name`\n* `vsys_id`\n* `host`\n* `vsys`\n* `cmd`\n* `admin`\n* `client`\n* `result`\n* `path`\n* `dg_id`\n* `comment`\n* `tpl_id`\n* `sender_sw_version`\n* `cef-formatted-receive_time`\n* `cef-formatted-time_generated`\n* `before-change-detail`\n* `after-change-detail`\n",
							Optional:            true,
						},
						"url_format": schema.StringAttribute{
							MarkdownDescription: "The URL path of the HTTP server",
							Optional:            true,
						},
					},
				},
				"userid": schema.SingleNestedAttribute{
					MarkdownDescription: "Userid",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"headers": schema.ListNestedAttribute{
							MarkdownDescription: "Headers",
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										MarkdownDescription: "Header name",
										Optional:            true,
									},
									"value": schema.StringAttribute{
										MarkdownDescription: "Header value",
										Optional:            true,
									},
								},
							},
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "The name of the payload format",
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("Default"),
						},
						"params": schema.ListNestedAttribute{
							MarkdownDescription: "Params",
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										MarkdownDescription: "Parameter name",
										Optional:            true,
									},
									"value": schema.StringAttribute{
										MarkdownDescription: "Parameter value",
										Optional:            true,
									},
								},
							},
						},
						"payload": schema.StringAttribute{
							MarkdownDescription: "The log payload format.  The accepted log field values are as follows.\n* `receive_time`\n* `serial`\n* `seqno`\n* `actionflags`\n* `type`\n* `subtype`\n* `time_generated`\n* `high_res_timestamp`\n* `dg_hier_level_1`\n* `dg_hier_level_2`\n* `dg_hier_level_3`\n* `dg_hier_level_4`\n* `vsys_name`\n* `device_name`\n* `vsys_id`\n* `host`\n* `vsys`\n* `cmd`\n* `admin`\n* `client`\n* `result`\n* `path`\n* `dg_id`\n* `comment`\n* `tpl_id`\n* `sender_sw_version`\n* `cef-formatted-receive_time`\n* `cef-formatted-time_generated`\n* `before-change-detail`\n* `after-change-detail`\n",
							Optional:            true,
						},
						"url_format": schema.StringAttribute{
							MarkdownDescription: "The URL path of the HTTP server",
							Optional:            true,
						},
					},
				},
				"wildfire": schema.SingleNestedAttribute{
					MarkdownDescription: "Wildfire",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"headers": schema.ListNestedAttribute{
							MarkdownDescription: "Headers",
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										MarkdownDescription: "Header name",
										Optional:            true,
									},
									"value": schema.StringAttribute{
										MarkdownDescription: "Header value",
										Optional:            true,
									},
								},
							},
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "The name of the payload format",
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("Default"),
						},
						"params": schema.ListNestedAttribute{
							MarkdownDescription: "Params",
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										MarkdownDescription: "Parameter name",
										Optional:            true,
									},
									"value": schema.StringAttribute{
										MarkdownDescription: "Parameter value",
										Optional:            true,
									},
								},
							},
						},
						"payload": schema.StringAttribute{
							MarkdownDescription: "The log payload format.  The accepted log field values are as follows.\n* `receive_time`\n* `serial`\n* `seqno`\n* `actionflags`\n* `type`\n* `subtype`\n* `time_generated`\n* `high_res_timestamp`\n* `dg_hier_level_1`\n* `dg_hier_level_2`\n* `dg_hier_level_3`\n* `dg_hier_level_4`\n* `vsys_name`\n* `device_name`\n* `vsys_id`\n* `host`\n* `vsys`\n* `cmd`\n* `admin`\n* `client`\n* `result`\n* `path`\n* `dg_id`\n* `comment`\n* `tpl_id`\n* `sender_sw_version`\n* `cef-formatted-receive_time`\n* `cef-formatted-time_generated`\n* `before-change-detail`\n* `after-change-detail`\n",
							Optional:            true,
						},
						"url_format": schema.StringAttribute{
							MarkdownDescription: "The URL path of the HTTP server",
							Optional:            true,
						},
					},
				},
			},
		},
		"id": schema.StringAttribute{
			MarkdownDescription: "The UUID of the HTTP server profile",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"name": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(63),
			},
			MarkdownDescription: "The name of the profile",
			Required:            true,
		},
		"server": schema.ListNestedAttribute{
			MarkdownDescription: "Server",
			Optional:            true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"address": schema.StringAttribute{
						MarkdownDescription: "HTTP server address",
						Optional:            true,
					},
					"certificate_profile": schema.StringAttribute{
						MarkdownDescription: "HTTP server certificate profile",
						Optional:            true,
						Computed:            true,
						Default:             stringdefault.StaticString("None"),
					},
					"http_method": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("GET", "POST", "PUT", "DELETE"),
						},
						MarkdownDescription: "HTTP operation to perform",
						Optional:            true,
					},
					"name": schema.StringAttribute{
						MarkdownDescription: "HTTP server name",
						Optional:            true,
					},
					"port": schema.Int64Attribute{
						MarkdownDescription: "HTTP server port",
						Optional:            true,
					},
					"protocol": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("HTTP", "HTTPS"),
						},
						MarkdownDescription: "HTTP server protocol",
						Optional:            true,
					},
					"tls_version": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("1.0", "1.1", "1.2", "1.3"),
						},
						MarkdownDescription: "HTTP server TLS version",
						Optional:            true,
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
			MarkdownDescription: "The snippet in which the resource is defined\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"tag_registration": schema.BoolAttribute{
			MarkdownDescription: "Register tags on match",
			Optional:            true,
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

// HttpServerProfilesDataSourceSchema defines the schema for HttpServerProfiles data source
var HttpServerProfilesDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "HttpServerProfile data source",
	Attributes: map[string]dsschema.Attribute{
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
		"format": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Format",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"auth": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Auth",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"headers": dsschema.ListNestedAttribute{
							MarkdownDescription: "Headers",
							Computed:            true,
							NestedObject: dsschema.NestedAttributeObject{
								Attributes: map[string]dsschema.Attribute{
									"name": dsschema.StringAttribute{
										MarkdownDescription: "Header name",
										Computed:            true,
									},
									"value": dsschema.StringAttribute{
										MarkdownDescription: "Header value",
										Computed:            true,
									},
								},
							},
						},
						"name": dsschema.StringAttribute{
							MarkdownDescription: "The name of the payload format",
							Computed:            true,
						},
						"params": dsschema.ListNestedAttribute{
							MarkdownDescription: "Params",
							Computed:            true,
							NestedObject: dsschema.NestedAttributeObject{
								Attributes: map[string]dsschema.Attribute{
									"name": dsschema.StringAttribute{
										MarkdownDescription: "Parameter name",
										Computed:            true,
									},
									"value": dsschema.StringAttribute{
										MarkdownDescription: "Parameter value",
										Computed:            true,
									},
								},
							},
						},
						"payload": dsschema.StringAttribute{
							MarkdownDescription: "The log payload format.  The accepted log field values are as follows.\n* `receive_time`\n* `serial`\n* `seqno`\n* `actionflags`\n* `type`\n* `subtype`\n* `time_generated`\n* `high_res_timestamp`\n* `dg_hier_level_1`\n* `dg_hier_level_2`\n* `dg_hier_level_3`\n* `dg_hier_level_4`\n* `vsys_name`\n* `device_name`\n* `vsys_id`\n* `host`\n* `vsys`\n* `cmd`\n* `admin`\n* `client`\n* `result`\n* `path`\n* `dg_id`\n* `comment`\n* `tpl_id`\n* `sender_sw_version`\n* `cef-formatted-receive_time`\n* `cef-formatted-time_generated`\n* `before-change-detail`\n* `after-change-detail`\n",
							Computed:            true,
						},
						"url_format": dsschema.StringAttribute{
							MarkdownDescription: "The URL path of the HTTP server",
							Computed:            true,
						},
					},
				},
				"config": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Config",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"headers": dsschema.ListNestedAttribute{
							MarkdownDescription: "Headers",
							Computed:            true,
							NestedObject: dsschema.NestedAttributeObject{
								Attributes: map[string]dsschema.Attribute{
									"name": dsschema.StringAttribute{
										MarkdownDescription: "Header name",
										Computed:            true,
									},
									"value": dsschema.StringAttribute{
										MarkdownDescription: "Header value",
										Computed:            true,
									},
								},
							},
						},
						"name": dsschema.StringAttribute{
							MarkdownDescription: "The name of the payload format",
							Computed:            true,
						},
						"params": dsschema.ListNestedAttribute{
							MarkdownDescription: "Params",
							Computed:            true,
							NestedObject: dsschema.NestedAttributeObject{
								Attributes: map[string]dsschema.Attribute{
									"name": dsschema.StringAttribute{
										MarkdownDescription: "Parameter name",
										Computed:            true,
									},
									"value": dsschema.StringAttribute{
										MarkdownDescription: "Parameter value",
										Computed:            true,
									},
								},
							},
						},
						"payload": dsschema.StringAttribute{
							MarkdownDescription: "The log payload format.  The accepted log field values are as follows.\n* `receive_time`\n* `serial`\n* `seqno`\n* `actionflags`\n* `type`\n* `subtype`\n* `time_generated`\n* `high_res_timestamp`\n* `dg_hier_level_1`\n* `dg_hier_level_2`\n* `dg_hier_level_3`\n* `dg_hier_level_4`\n* `vsys_name`\n* `device_name`\n* `vsys_id`\n* `host`\n* `vsys`\n* `cmd`\n* `admin`\n* `client`\n* `result`\n* `path`\n* `dg_id`\n* `comment`\n* `tpl_id`\n* `sender_sw_version`\n* `cef-formatted-receive_time`\n* `cef-formatted-time_generated`\n* `before-change-detail`\n* `after-change-detail`\n",
							Computed:            true,
						},
						"url_format": dsschema.StringAttribute{
							MarkdownDescription: "The URL path of the HTTP server",
							Computed:            true,
						},
					},
				},
				"correlation": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Correlation",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"headers": dsschema.ListNestedAttribute{
							MarkdownDescription: "Headers",
							Computed:            true,
							NestedObject: dsschema.NestedAttributeObject{
								Attributes: map[string]dsschema.Attribute{
									"name": dsschema.StringAttribute{
										MarkdownDescription: "Header name",
										Computed:            true,
									},
									"value": dsschema.StringAttribute{
										MarkdownDescription: "Header value",
										Computed:            true,
									},
								},
							},
						},
						"name": dsschema.StringAttribute{
							MarkdownDescription: "The name of the payload format",
							Computed:            true,
						},
						"params": dsschema.ListNestedAttribute{
							MarkdownDescription: "Params",
							Computed:            true,
							NestedObject: dsschema.NestedAttributeObject{
								Attributes: map[string]dsschema.Attribute{
									"name": dsschema.StringAttribute{
										MarkdownDescription: "Parameter name",
										Computed:            true,
									},
									"value": dsschema.StringAttribute{
										MarkdownDescription: "Parameter value",
										Computed:            true,
									},
								},
							},
						},
						"payload": dsschema.StringAttribute{
							MarkdownDescription: "The log payload format.  The accepted log field values are as follows.\n* `receive_time`\n* `serial`\n* `seqno`\n* `actionflags`\n* `type`\n* `subtype`\n* `time_generated`\n* `high_res_timestamp`\n* `dg_hier_level_1`\n* `dg_hier_level_2`\n* `dg_hier_level_3`\n* `dg_hier_level_4`\n* `vsys_name`\n* `device_name`\n* `vsys_id`\n* `host`\n* `vsys`\n* `cmd`\n* `admin`\n* `client`\n* `result`\n* `path`\n* `dg_id`\n* `comment`\n* `tpl_id`\n* `sender_sw_version`\n* `cef-formatted-receive_time`\n* `cef-formatted-time_generated`\n* `before-change-detail`\n* `after-change-detail`\n",
							Computed:            true,
						},
						"url_format": dsschema.StringAttribute{
							MarkdownDescription: "The URL path of the HTTP server",
							Computed:            true,
						},
					},
				},
				"data": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Data",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"headers": dsschema.ListNestedAttribute{
							MarkdownDescription: "Headers",
							Computed:            true,
							NestedObject: dsschema.NestedAttributeObject{
								Attributes: map[string]dsschema.Attribute{
									"name": dsschema.StringAttribute{
										MarkdownDescription: "Header name",
										Computed:            true,
									},
									"value": dsschema.StringAttribute{
										MarkdownDescription: "Header value",
										Computed:            true,
									},
								},
							},
						},
						"name": dsschema.StringAttribute{
							MarkdownDescription: "The name of the payload format",
							Computed:            true,
						},
						"params": dsschema.ListNestedAttribute{
							MarkdownDescription: "Params",
							Computed:            true,
							NestedObject: dsschema.NestedAttributeObject{
								Attributes: map[string]dsschema.Attribute{
									"name": dsschema.StringAttribute{
										MarkdownDescription: "Parameter name",
										Computed:            true,
									},
									"value": dsschema.StringAttribute{
										MarkdownDescription: "Parameter value",
										Computed:            true,
									},
								},
							},
						},
						"payload": dsschema.StringAttribute{
							MarkdownDescription: "The log payload format.  The accepted log field values are as follows.\n* `receive_time`\n* `serial`\n* `seqno`\n* `actionflags`\n* `type`\n* `subtype`\n* `time_generated`\n* `high_res_timestamp`\n* `dg_hier_level_1`\n* `dg_hier_level_2`\n* `dg_hier_level_3`\n* `dg_hier_level_4`\n* `vsys_name`\n* `device_name`\n* `vsys_id`\n* `host`\n* `vsys`\n* `cmd`\n* `admin`\n* `client`\n* `result`\n* `path`\n* `dg_id`\n* `comment`\n* `tpl_id`\n* `sender_sw_version`\n* `cef-formatted-receive_time`\n* `cef-formatted-time_generated`\n* `before-change-detail`\n* `after-change-detail`\n",
							Computed:            true,
						},
						"url_format": dsschema.StringAttribute{
							MarkdownDescription: "The URL path of the HTTP server",
							Computed:            true,
						},
					},
				},
				"decryption": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Decryption",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"headers": dsschema.ListNestedAttribute{
							MarkdownDescription: "Headers",
							Computed:            true,
							NestedObject: dsschema.NestedAttributeObject{
								Attributes: map[string]dsschema.Attribute{
									"name": dsschema.StringAttribute{
										MarkdownDescription: "Header name",
										Computed:            true,
									},
									"value": dsschema.StringAttribute{
										MarkdownDescription: "Header value",
										Computed:            true,
									},
								},
							},
						},
						"name": dsschema.StringAttribute{
							MarkdownDescription: "The name of the payload format",
							Computed:            true,
						},
						"params": dsschema.ListNestedAttribute{
							MarkdownDescription: "Params",
							Computed:            true,
							NestedObject: dsschema.NestedAttributeObject{
								Attributes: map[string]dsschema.Attribute{
									"name": dsschema.StringAttribute{
										MarkdownDescription: "Parameter name",
										Computed:            true,
									},
									"value": dsschema.StringAttribute{
										MarkdownDescription: "Parameter value",
										Computed:            true,
									},
								},
							},
						},
						"payload": dsschema.StringAttribute{
							MarkdownDescription: "The log payload format.  The accepted log field values are as follows.\n* `receive_time`\n* `serial`\n* `seqno`\n* `actionflags`\n* `type`\n* `subtype`\n* `time_generated`\n* `high_res_timestamp`\n* `dg_hier_level_1`\n* `dg_hier_level_2`\n* `dg_hier_level_3`\n* `dg_hier_level_4`\n* `vsys_name`\n* `device_name`\n* `vsys_id`\n* `host`\n* `vsys`\n* `cmd`\n* `admin`\n* `client`\n* `result`\n* `path`\n* `dg_id`\n* `comment`\n* `tpl_id`\n* `sender_sw_version`\n* `cef-formatted-receive_time`\n* `cef-formatted-time_generated`\n* `before-change-detail`\n* `after-change-detail`\n",
							Computed:            true,
						},
						"url_format": dsschema.StringAttribute{
							MarkdownDescription: "The URL path of the HTTP server",
							Computed:            true,
						},
					},
				},
				"globalprotect": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Globalprotect",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"headers": dsschema.ListNestedAttribute{
							MarkdownDescription: "Headers",
							Computed:            true,
							NestedObject: dsschema.NestedAttributeObject{
								Attributes: map[string]dsschema.Attribute{
									"name": dsschema.StringAttribute{
										MarkdownDescription: "Header name",
										Computed:            true,
									},
									"value": dsschema.StringAttribute{
										MarkdownDescription: "Header value",
										Computed:            true,
									},
								},
							},
						},
						"name": dsschema.StringAttribute{
							MarkdownDescription: "The name of the payload format",
							Computed:            true,
						},
						"params": dsschema.ListNestedAttribute{
							MarkdownDescription: "Params",
							Computed:            true,
							NestedObject: dsschema.NestedAttributeObject{
								Attributes: map[string]dsschema.Attribute{
									"name": dsschema.StringAttribute{
										MarkdownDescription: "Parameter name",
										Computed:            true,
									},
									"value": dsschema.StringAttribute{
										MarkdownDescription: "Parameter value",
										Computed:            true,
									},
								},
							},
						},
						"payload": dsschema.StringAttribute{
							MarkdownDescription: "The log payload format.  The accepted log field values are as follows.\n* `receive_time`\n* `serial`\n* `seqno`\n* `actionflags`\n* `type`\n* `subtype`\n* `time_generated`\n* `high_res_timestamp`\n* `dg_hier_level_1`\n* `dg_hier_level_2`\n* `dg_hier_level_3`\n* `dg_hier_level_4`\n* `vsys_name`\n* `device_name`\n* `vsys_id`\n* `host`\n* `vsys`\n* `cmd`\n* `admin`\n* `client`\n* `result`\n* `path`\n* `dg_id`\n* `comment`\n* `tpl_id`\n* `sender_sw_version`\n* `cef-formatted-receive_time`\n* `cef-formatted-time_generated`\n* `before-change-detail`\n* `after-change-detail`\n",
							Computed:            true,
						},
						"url_format": dsschema.StringAttribute{
							MarkdownDescription: "The URL path of the HTTP server",
							Computed:            true,
						},
					},
				},
				"gtp": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Gtp",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"headers": dsschema.ListNestedAttribute{
							MarkdownDescription: "Headers",
							Computed:            true,
							NestedObject: dsschema.NestedAttributeObject{
								Attributes: map[string]dsschema.Attribute{
									"name": dsschema.StringAttribute{
										MarkdownDescription: "Header name",
										Computed:            true,
									},
									"value": dsschema.StringAttribute{
										MarkdownDescription: "Header value",
										Computed:            true,
									},
								},
							},
						},
						"name": dsschema.StringAttribute{
							MarkdownDescription: "The name of the payload format",
							Computed:            true,
						},
						"params": dsschema.ListNestedAttribute{
							MarkdownDescription: "Params",
							Computed:            true,
							NestedObject: dsschema.NestedAttributeObject{
								Attributes: map[string]dsschema.Attribute{
									"name": dsschema.StringAttribute{
										MarkdownDescription: "Parameter name",
										Computed:            true,
									},
									"value": dsschema.StringAttribute{
										MarkdownDescription: "Parameter value",
										Computed:            true,
									},
								},
							},
						},
						"payload": dsschema.StringAttribute{
							MarkdownDescription: "The log payload format.  The accepted log field values are as follows.\n* `receive_time`\n* `serial`\n* `seqno`\n* `actionflags`\n* `type`\n* `subtype`\n* `time_generated`\n* `high_res_timestamp`\n* `dg_hier_level_1`\n* `dg_hier_level_2`\n* `dg_hier_level_3`\n* `dg_hier_level_4`\n* `vsys_name`\n* `device_name`\n* `vsys_id`\n* `host`\n* `vsys`\n* `cmd`\n* `admin`\n* `client`\n* `result`\n* `path`\n* `dg_id`\n* `comment`\n* `tpl_id`\n* `sender_sw_version`\n* `cef-formatted-receive_time`\n* `cef-formatted-time_generated`\n* `before-change-detail`\n* `after-change-detail`\n",
							Computed:            true,
						},
						"url_format": dsschema.StringAttribute{
							MarkdownDescription: "The URL path of the HTTP server",
							Computed:            true,
						},
					},
				},
				"hip_match": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Hip match",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"headers": dsschema.ListNestedAttribute{
							MarkdownDescription: "Headers",
							Computed:            true,
							NestedObject: dsschema.NestedAttributeObject{
								Attributes: map[string]dsschema.Attribute{
									"name": dsschema.StringAttribute{
										MarkdownDescription: "Header name",
										Computed:            true,
									},
									"value": dsschema.StringAttribute{
										MarkdownDescription: "Header value",
										Computed:            true,
									},
								},
							},
						},
						"name": dsschema.StringAttribute{
							MarkdownDescription: "The name of the payload format",
							Computed:            true,
						},
						"params": dsschema.ListNestedAttribute{
							MarkdownDescription: "Params",
							Computed:            true,
							NestedObject: dsschema.NestedAttributeObject{
								Attributes: map[string]dsschema.Attribute{
									"name": dsschema.StringAttribute{
										MarkdownDescription: "Parameter name",
										Computed:            true,
									},
									"value": dsschema.StringAttribute{
										MarkdownDescription: "Parameter value",
										Computed:            true,
									},
								},
							},
						},
						"payload": dsschema.StringAttribute{
							MarkdownDescription: "The log payload format.  The accepted log field values are as follows.\n* `receive_time`\n* `serial`\n* `seqno`\n* `actionflags`\n* `type`\n* `subtype`\n* `time_generated`\n* `high_res_timestamp`\n* `dg_hier_level_1`\n* `dg_hier_level_2`\n* `dg_hier_level_3`\n* `dg_hier_level_4`\n* `vsys_name`\n* `device_name`\n* `vsys_id`\n* `host`\n* `vsys`\n* `cmd`\n* `admin`\n* `client`\n* `result`\n* `path`\n* `dg_id`\n* `comment`\n* `tpl_id`\n* `sender_sw_version`\n* `cef-formatted-receive_time`\n* `cef-formatted-time_generated`\n* `before-change-detail`\n* `after-change-detail`\n",
							Computed:            true,
						},
						"url_format": dsschema.StringAttribute{
							MarkdownDescription: "The URL path of the HTTP server",
							Computed:            true,
						},
					},
				},
				"iptag": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Iptag",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"headers": dsschema.ListNestedAttribute{
							MarkdownDescription: "Headers",
							Computed:            true,
							NestedObject: dsschema.NestedAttributeObject{
								Attributes: map[string]dsschema.Attribute{
									"name": dsschema.StringAttribute{
										MarkdownDescription: "Header name",
										Computed:            true,
									},
									"value": dsschema.StringAttribute{
										MarkdownDescription: "Header value",
										Computed:            true,
									},
								},
							},
						},
						"name": dsschema.StringAttribute{
							MarkdownDescription: "The name of the payload format",
							Computed:            true,
						},
						"params": dsschema.ListNestedAttribute{
							MarkdownDescription: "Params",
							Computed:            true,
							NestedObject: dsschema.NestedAttributeObject{
								Attributes: map[string]dsschema.Attribute{
									"name": dsschema.StringAttribute{
										MarkdownDescription: "Parameter name",
										Computed:            true,
									},
									"value": dsschema.StringAttribute{
										MarkdownDescription: "Parameter value",
										Computed:            true,
									},
								},
							},
						},
						"payload": dsschema.StringAttribute{
							MarkdownDescription: "The log payload format.  The accepted log field values are as follows.\n* `receive_time`\n* `serial`\n* `seqno`\n* `actionflags`\n* `type`\n* `subtype`\n* `time_generated`\n* `high_res_timestamp`\n* `dg_hier_level_1`\n* `dg_hier_level_2`\n* `dg_hier_level_3`\n* `dg_hier_level_4`\n* `vsys_name`\n* `device_name`\n* `vsys_id`\n* `host`\n* `vsys`\n* `cmd`\n* `admin`\n* `client`\n* `result`\n* `path`\n* `dg_id`\n* `comment`\n* `tpl_id`\n* `sender_sw_version`\n* `cef-formatted-receive_time`\n* `cef-formatted-time_generated`\n* `before-change-detail`\n* `after-change-detail`\n",
							Computed:            true,
						},
						"url_format": dsschema.StringAttribute{
							MarkdownDescription: "The URL path of the HTTP server",
							Computed:            true,
						},
					},
				},
				"sctp": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Sctp",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"headers": dsschema.ListNestedAttribute{
							MarkdownDescription: "Headers",
							Computed:            true,
							NestedObject: dsschema.NestedAttributeObject{
								Attributes: map[string]dsschema.Attribute{
									"name": dsschema.StringAttribute{
										MarkdownDescription: "Header name",
										Computed:            true,
									},
									"value": dsschema.StringAttribute{
										MarkdownDescription: "Header value",
										Computed:            true,
									},
								},
							},
						},
						"name": dsschema.StringAttribute{
							MarkdownDescription: "The name of the payload format",
							Computed:            true,
						},
						"params": dsschema.ListNestedAttribute{
							MarkdownDescription: "Params",
							Computed:            true,
							NestedObject: dsschema.NestedAttributeObject{
								Attributes: map[string]dsschema.Attribute{
									"name": dsschema.StringAttribute{
										MarkdownDescription: "Parameter name",
										Computed:            true,
									},
									"value": dsschema.StringAttribute{
										MarkdownDescription: "Parameter value",
										Computed:            true,
									},
								},
							},
						},
						"payload": dsschema.StringAttribute{
							MarkdownDescription: "The log payload format.  The accepted log field values are as follows.\n* `receive_time`\n* `serial`\n* `seqno`\n* `actionflags`\n* `type`\n* `subtype`\n* `time_generated`\n* `high_res_timestamp`\n* `dg_hier_level_1`\n* `dg_hier_level_2`\n* `dg_hier_level_3`\n* `dg_hier_level_4`\n* `vsys_name`\n* `device_name`\n* `vsys_id`\n* `host`\n* `vsys`\n* `cmd`\n* `admin`\n* `client`\n* `result`\n* `path`\n* `dg_id`\n* `comment`\n* `tpl_id`\n* `sender_sw_version`\n* `cef-formatted-receive_time`\n* `cef-formatted-time_generated`\n* `before-change-detail`\n* `after-change-detail`\n",
							Computed:            true,
						},
						"url_format": dsschema.StringAttribute{
							MarkdownDescription: "The URL path of the HTTP server",
							Computed:            true,
						},
					},
				},
				"system": dsschema.SingleNestedAttribute{
					MarkdownDescription: "System",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"headers": dsschema.ListNestedAttribute{
							MarkdownDescription: "Headers",
							Computed:            true,
							NestedObject: dsschema.NestedAttributeObject{
								Attributes: map[string]dsschema.Attribute{
									"name": dsschema.StringAttribute{
										MarkdownDescription: "Header name",
										Computed:            true,
									},
									"value": dsschema.StringAttribute{
										MarkdownDescription: "Header value",
										Computed:            true,
									},
								},
							},
						},
						"name": dsschema.StringAttribute{
							MarkdownDescription: "The name of the payload format",
							Computed:            true,
						},
						"params": dsschema.ListNestedAttribute{
							MarkdownDescription: "Params",
							Computed:            true,
							NestedObject: dsschema.NestedAttributeObject{
								Attributes: map[string]dsschema.Attribute{
									"name": dsschema.StringAttribute{
										MarkdownDescription: "Parameter name",
										Computed:            true,
									},
									"value": dsschema.StringAttribute{
										MarkdownDescription: "Parameter value",
										Computed:            true,
									},
								},
							},
						},
						"payload": dsschema.StringAttribute{
							MarkdownDescription: "The log payload format.  The accepted log field values are as follows.\n* `receive_time`\n* `serial`\n* `seqno`\n* `actionflags`\n* `type`\n* `subtype`\n* `time_generated`\n* `high_res_timestamp`\n* `dg_hier_level_1`\n* `dg_hier_level_2`\n* `dg_hier_level_3`\n* `dg_hier_level_4`\n* `vsys_name`\n* `device_name`\n* `vsys_id`\n* `host`\n* `vsys`\n* `cmd`\n* `admin`\n* `client`\n* `result`\n* `path`\n* `dg_id`\n* `comment`\n* `tpl_id`\n* `sender_sw_version`\n* `cef-formatted-receive_time`\n* `cef-formatted-time_generated`\n* `before-change-detail`\n* `after-change-detail`\n",
							Computed:            true,
						},
						"url_format": dsschema.StringAttribute{
							MarkdownDescription: "The URL path of the HTTP server",
							Computed:            true,
						},
					},
				},
				"threat": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Threat",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"headers": dsschema.ListNestedAttribute{
							MarkdownDescription: "Headers",
							Computed:            true,
							NestedObject: dsschema.NestedAttributeObject{
								Attributes: map[string]dsschema.Attribute{
									"name": dsschema.StringAttribute{
										MarkdownDescription: "Header name",
										Computed:            true,
									},
									"value": dsschema.StringAttribute{
										MarkdownDescription: "Header value",
										Computed:            true,
									},
								},
							},
						},
						"name": dsschema.StringAttribute{
							MarkdownDescription: "The name of the payload format",
							Computed:            true,
						},
						"params": dsschema.ListNestedAttribute{
							MarkdownDescription: "Params",
							Computed:            true,
							NestedObject: dsschema.NestedAttributeObject{
								Attributes: map[string]dsschema.Attribute{
									"name": dsschema.StringAttribute{
										MarkdownDescription: "Parameter name",
										Computed:            true,
									},
									"value": dsschema.StringAttribute{
										MarkdownDescription: "Parameter value",
										Computed:            true,
									},
								},
							},
						},
						"payload": dsschema.StringAttribute{
							MarkdownDescription: "The log payload format.  The accepted log field values are as follows.\n* `receive_time`\n* `serial`\n* `seqno`\n* `actionflags`\n* `type`\n* `subtype`\n* `time_generated`\n* `high_res_timestamp`\n* `dg_hier_level_1`\n* `dg_hier_level_2`\n* `dg_hier_level_3`\n* `dg_hier_level_4`\n* `vsys_name`\n* `device_name`\n* `vsys_id`\n* `host`\n* `vsys`\n* `cmd`\n* `admin`\n* `client`\n* `result`\n* `path`\n* `dg_id`\n* `comment`\n* `tpl_id`\n* `sender_sw_version`\n* `cef-formatted-receive_time`\n* `cef-formatted-time_generated`\n* `before-change-detail`\n* `after-change-detail`\n",
							Computed:            true,
						},
						"url_format": dsschema.StringAttribute{
							MarkdownDescription: "The URL path of the HTTP server",
							Computed:            true,
						},
					},
				},
				"traffic": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Traffic",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"headers": dsschema.ListNestedAttribute{
							MarkdownDescription: "Headers",
							Computed:            true,
							NestedObject: dsschema.NestedAttributeObject{
								Attributes: map[string]dsschema.Attribute{
									"name": dsschema.StringAttribute{
										MarkdownDescription: "Header name",
										Computed:            true,
									},
									"value": dsschema.StringAttribute{
										MarkdownDescription: "Header value",
										Computed:            true,
									},
								},
							},
						},
						"name": dsschema.StringAttribute{
							MarkdownDescription: "The name of the payload format",
							Computed:            true,
						},
						"params": dsschema.ListNestedAttribute{
							MarkdownDescription: "Params",
							Computed:            true,
							NestedObject: dsschema.NestedAttributeObject{
								Attributes: map[string]dsschema.Attribute{
									"name": dsschema.StringAttribute{
										MarkdownDescription: "Parameter name",
										Computed:            true,
									},
									"value": dsschema.StringAttribute{
										MarkdownDescription: "Parameter value",
										Computed:            true,
									},
								},
							},
						},
						"payload": dsschema.StringAttribute{
							MarkdownDescription: "The log payload format.  The accepted log field values are as follows.\n* `receive_time`\n* `serial`\n* `seqno`\n* `actionflags`\n* `type`\n* `subtype`\n* `time_generated`\n* `high_res_timestamp`\n* `dg_hier_level_1`\n* `dg_hier_level_2`\n* `dg_hier_level_3`\n* `dg_hier_level_4`\n* `vsys_name`\n* `device_name`\n* `vsys_id`\n* `host`\n* `vsys`\n* `cmd`\n* `admin`\n* `client`\n* `result`\n* `path`\n* `dg_id`\n* `comment`\n* `tpl_id`\n* `sender_sw_version`\n* `cef-formatted-receive_time`\n* `cef-formatted-time_generated`\n* `before-change-detail`\n* `after-change-detail`\n",
							Computed:            true,
						},
						"url_format": dsschema.StringAttribute{
							MarkdownDescription: "The URL path of the HTTP server",
							Computed:            true,
						},
					},
				},
				"tunnel": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Tunnel",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"headers": dsschema.ListNestedAttribute{
							MarkdownDescription: "Headers",
							Computed:            true,
							NestedObject: dsschema.NestedAttributeObject{
								Attributes: map[string]dsschema.Attribute{
									"name": dsschema.StringAttribute{
										MarkdownDescription: "Header name",
										Computed:            true,
									},
									"value": dsschema.StringAttribute{
										MarkdownDescription: "Header value",
										Computed:            true,
									},
								},
							},
						},
						"name": dsschema.StringAttribute{
							MarkdownDescription: "The name of the payload format",
							Computed:            true,
						},
						"params": dsschema.ListNestedAttribute{
							MarkdownDescription: "Params",
							Computed:            true,
							NestedObject: dsschema.NestedAttributeObject{
								Attributes: map[string]dsschema.Attribute{
									"name": dsschema.StringAttribute{
										MarkdownDescription: "Parameter name",
										Computed:            true,
									},
									"value": dsschema.StringAttribute{
										MarkdownDescription: "Parameter value",
										Computed:            true,
									},
								},
							},
						},
						"payload": dsschema.StringAttribute{
							MarkdownDescription: "The log payload format.  The accepted log field values are as follows.\n* `receive_time`\n* `serial`\n* `seqno`\n* `actionflags`\n* `type`\n* `subtype`\n* `time_generated`\n* `high_res_timestamp`\n* `dg_hier_level_1`\n* `dg_hier_level_2`\n* `dg_hier_level_3`\n* `dg_hier_level_4`\n* `vsys_name`\n* `device_name`\n* `vsys_id`\n* `host`\n* `vsys`\n* `cmd`\n* `admin`\n* `client`\n* `result`\n* `path`\n* `dg_id`\n* `comment`\n* `tpl_id`\n* `sender_sw_version`\n* `cef-formatted-receive_time`\n* `cef-formatted-time_generated`\n* `before-change-detail`\n* `after-change-detail`\n",
							Computed:            true,
						},
						"url_format": dsschema.StringAttribute{
							MarkdownDescription: "The URL path of the HTTP server",
							Computed:            true,
						},
					},
				},
				"url": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Url",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"headers": dsschema.ListNestedAttribute{
							MarkdownDescription: "Headers",
							Computed:            true,
							NestedObject: dsschema.NestedAttributeObject{
								Attributes: map[string]dsschema.Attribute{
									"name": dsschema.StringAttribute{
										MarkdownDescription: "Header name",
										Computed:            true,
									},
									"value": dsschema.StringAttribute{
										MarkdownDescription: "Header value",
										Computed:            true,
									},
								},
							},
						},
						"name": dsschema.StringAttribute{
							MarkdownDescription: "The name of the payload format",
							Computed:            true,
						},
						"params": dsschema.ListNestedAttribute{
							MarkdownDescription: "Params",
							Computed:            true,
							NestedObject: dsschema.NestedAttributeObject{
								Attributes: map[string]dsschema.Attribute{
									"name": dsschema.StringAttribute{
										MarkdownDescription: "Parameter name",
										Computed:            true,
									},
									"value": dsschema.StringAttribute{
										MarkdownDescription: "Parameter value",
										Computed:            true,
									},
								},
							},
						},
						"payload": dsschema.StringAttribute{
							MarkdownDescription: "The log payload format.  The accepted log field values are as follows.\n* `receive_time`\n* `serial`\n* `seqno`\n* `actionflags`\n* `type`\n* `subtype`\n* `time_generated`\n* `high_res_timestamp`\n* `dg_hier_level_1`\n* `dg_hier_level_2`\n* `dg_hier_level_3`\n* `dg_hier_level_4`\n* `vsys_name`\n* `device_name`\n* `vsys_id`\n* `host`\n* `vsys`\n* `cmd`\n* `admin`\n* `client`\n* `result`\n* `path`\n* `dg_id`\n* `comment`\n* `tpl_id`\n* `sender_sw_version`\n* `cef-formatted-receive_time`\n* `cef-formatted-time_generated`\n* `before-change-detail`\n* `after-change-detail`\n",
							Computed:            true,
						},
						"url_format": dsschema.StringAttribute{
							MarkdownDescription: "The URL path of the HTTP server",
							Computed:            true,
						},
					},
				},
				"userid": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Userid",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"headers": dsschema.ListNestedAttribute{
							MarkdownDescription: "Headers",
							Computed:            true,
							NestedObject: dsschema.NestedAttributeObject{
								Attributes: map[string]dsschema.Attribute{
									"name": dsschema.StringAttribute{
										MarkdownDescription: "Header name",
										Computed:            true,
									},
									"value": dsschema.StringAttribute{
										MarkdownDescription: "Header value",
										Computed:            true,
									},
								},
							},
						},
						"name": dsschema.StringAttribute{
							MarkdownDescription: "The name of the payload format",
							Computed:            true,
						},
						"params": dsschema.ListNestedAttribute{
							MarkdownDescription: "Params",
							Computed:            true,
							NestedObject: dsschema.NestedAttributeObject{
								Attributes: map[string]dsschema.Attribute{
									"name": dsschema.StringAttribute{
										MarkdownDescription: "Parameter name",
										Computed:            true,
									},
									"value": dsschema.StringAttribute{
										MarkdownDescription: "Parameter value",
										Computed:            true,
									},
								},
							},
						},
						"payload": dsschema.StringAttribute{
							MarkdownDescription: "The log payload format.  The accepted log field values are as follows.\n* `receive_time`\n* `serial`\n* `seqno`\n* `actionflags`\n* `type`\n* `subtype`\n* `time_generated`\n* `high_res_timestamp`\n* `dg_hier_level_1`\n* `dg_hier_level_2`\n* `dg_hier_level_3`\n* `dg_hier_level_4`\n* `vsys_name`\n* `device_name`\n* `vsys_id`\n* `host`\n* `vsys`\n* `cmd`\n* `admin`\n* `client`\n* `result`\n* `path`\n* `dg_id`\n* `comment`\n* `tpl_id`\n* `sender_sw_version`\n* `cef-formatted-receive_time`\n* `cef-formatted-time_generated`\n* `before-change-detail`\n* `after-change-detail`\n",
							Computed:            true,
						},
						"url_format": dsschema.StringAttribute{
							MarkdownDescription: "The URL path of the HTTP server",
							Computed:            true,
						},
					},
				},
				"wildfire": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Wildfire",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"headers": dsschema.ListNestedAttribute{
							MarkdownDescription: "Headers",
							Computed:            true,
							NestedObject: dsschema.NestedAttributeObject{
								Attributes: map[string]dsschema.Attribute{
									"name": dsschema.StringAttribute{
										MarkdownDescription: "Header name",
										Computed:            true,
									},
									"value": dsschema.StringAttribute{
										MarkdownDescription: "Header value",
										Computed:            true,
									},
								},
							},
						},
						"name": dsschema.StringAttribute{
							MarkdownDescription: "The name of the payload format",
							Computed:            true,
						},
						"params": dsschema.ListNestedAttribute{
							MarkdownDescription: "Params",
							Computed:            true,
							NestedObject: dsschema.NestedAttributeObject{
								Attributes: map[string]dsschema.Attribute{
									"name": dsschema.StringAttribute{
										MarkdownDescription: "Parameter name",
										Computed:            true,
									},
									"value": dsschema.StringAttribute{
										MarkdownDescription: "Parameter value",
										Computed:            true,
									},
								},
							},
						},
						"payload": dsschema.StringAttribute{
							MarkdownDescription: "The log payload format.  The accepted log field values are as follows.\n* `receive_time`\n* `serial`\n* `seqno`\n* `actionflags`\n* `type`\n* `subtype`\n* `time_generated`\n* `high_res_timestamp`\n* `dg_hier_level_1`\n* `dg_hier_level_2`\n* `dg_hier_level_3`\n* `dg_hier_level_4`\n* `vsys_name`\n* `device_name`\n* `vsys_id`\n* `host`\n* `vsys`\n* `cmd`\n* `admin`\n* `client`\n* `result`\n* `path`\n* `dg_id`\n* `comment`\n* `tpl_id`\n* `sender_sw_version`\n* `cef-formatted-receive_time`\n* `cef-formatted-time_generated`\n* `before-change-detail`\n* `after-change-detail`\n",
							Computed:            true,
						},
						"url_format": dsschema.StringAttribute{
							MarkdownDescription: "The URL path of the HTTP server",
							Computed:            true,
						},
					},
				},
			},
		},
		"id": dsschema.StringAttribute{
			MarkdownDescription: "The UUID of the HTTP server profile",
			Required:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "The name of the profile",
			Optional:            true,
			Computed:            true,
		},
		"server": dsschema.ListNestedAttribute{
			MarkdownDescription: "Server",
			Computed:            true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: map[string]dsschema.Attribute{
					"address": dsschema.StringAttribute{
						MarkdownDescription: "HTTP server address",
						Computed:            true,
					},
					"certificate_profile": dsschema.StringAttribute{
						MarkdownDescription: "HTTP server certificate profile",
						Computed:            true,
					},
					"http_method": dsschema.StringAttribute{
						MarkdownDescription: "HTTP operation to perform",
						Computed:            true,
					},
					"name": dsschema.StringAttribute{
						MarkdownDescription: "HTTP server name",
						Computed:            true,
					},
					"port": dsschema.Int64Attribute{
						MarkdownDescription: "HTTP server port",
						Computed:            true,
					},
					"protocol": dsschema.StringAttribute{
						MarkdownDescription: "HTTP server protocol",
						Computed:            true,
					},
					"tls_version": dsschema.StringAttribute{
						MarkdownDescription: "HTTP server TLS version",
						Computed:            true,
					},
				},
			},
		},
		"snippet": dsschema.StringAttribute{
			MarkdownDescription: "The snippet in which the resource is defined\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			Computed:            true,
		},
		"tag_registration": dsschema.BoolAttribute{
			MarkdownDescription: "Register tags on match",
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
	},
}

// HttpServerProfilesListModel represents the data model for a list data source.
type HttpServerProfilesListModel struct {
	Tfid    types.String         `tfsdk:"tfid"`
	Data    []HttpServerProfiles `tfsdk:"data"`
	Limit   types.Int64          `tfsdk:"limit"`
	Offset  types.Int64          `tfsdk:"offset"`
	Name    types.String         `tfsdk:"name"`
	Total   types.Int64          `tfsdk:"total"`
	Folder  types.String         `tfsdk:"folder"`
	Snippet types.String         `tfsdk:"snippet"`
	Device  types.String         `tfsdk:"device"`
}

// HttpServerProfilesListDataSourceSchema defines the schema for a list data source.
var HttpServerProfilesListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: HttpServerProfilesDataSourceSchema.Attributes,
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
