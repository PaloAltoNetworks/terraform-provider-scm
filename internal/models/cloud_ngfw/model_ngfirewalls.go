package models

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	dsschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Package: cloud_ngfw
// This file contains models for the cloud_ngfw SDK package

// Ngfirewalls represents the Terraform model for Ngfirewalls
type Ngfirewalls struct {
	Tfid           types.String          `tfsdk:"tfid"`
	Response       basetypes.ObjectValue `tfsdk:"response"`
	ResponseStatus basetypes.ObjectValue `tfsdk:"response_status"`
	Region         basetypes.StringValue `tfsdk:"region"`
	Id             basetypes.StringValue `tfsdk:"id"`
}

// NgfirewallResponseBody represents a nested structure within the Ngfirewalls model
type NgfirewallResponseBody struct {
	Firewall basetypes.ObjectValue `tfsdk:"firewall"`
	Status   basetypes.ObjectValue `tfsdk:"status"`
}

// NgfirewallFirewall represents a nested structure within the Ngfirewalls model
type NgfirewallFirewall struct {
	AdvancedThreatLog           basetypes.BoolValue   `tfsdk:"advanced_threat_log"`
	AllowListAccounts           basetypes.ListValue   `tfsdk:"allow_list_accounts"`
	ChangeProtection            basetypes.ListValue   `tfsdk:"change_protection"`
	CloudwatchMetrics           basetypes.ObjectValue `tfsdk:"cloudwatch_metrics"`
	CustomerZoneIdList          basetypes.ListValue   `tfsdk:"customer_zone_id_list"`
	DeploymentUpdateToken       basetypes.StringValue `tfsdk:"deployment_update_token"`
	Description                 basetypes.StringValue `tfsdk:"description"`
	DeviceRuleStackCommitStatus basetypes.StringValue `tfsdk:"device_rule_stack_commit_status"`
	EgressNAT                   basetypes.ObjectValue `tfsdk:"egress_nat"`
	EndpointServiceName         basetypes.StringValue `tfsdk:"endpoint_service_name"`
	Endpoints                   basetypes.ListValue   `tfsdk:"endpoints"`
	FirewallId                  basetypes.StringValue `tfsdk:"firewall_id"`
	GWLB                        basetypes.ObjectValue `tfsdk:"gwlb"`
	Ipv6                        basetypes.ObjectValue `tfsdk:"ipv6"`
	LinkId                      basetypes.StringValue `tfsdk:"link_id"`
	LinkStatus                  basetypes.StringValue `tfsdk:"link_status"`
	LogConfig                   basetypes.ObjectValue `tfsdk:"log_config"`
	Notifications               basetypes.ListValue   `tfsdk:"notifications"`
	PrivateAccess               basetypes.ObjectValue `tfsdk:"private_access"`
	Region                      basetypes.StringValue `tfsdk:"region"`
	SoftwareVersion             basetypes.StringValue `tfsdk:"software_version"`
	Tags                        basetypes.ListValue   `tfsdk:"tags"`
	Tier                        basetypes.StringValue `tfsdk:"tier"`
	UpdateToken                 basetypes.StringValue `tfsdk:"update_token"`
	UserID                      basetypes.ObjectValue `tfsdk:"user_id"`
}

// NgfirewallCloudwatchMetrics represents a nested structure within the Ngfirewalls model
type NgfirewallCloudwatchMetrics struct {
	AccountId                 basetypes.StringValue `tfsdk:"account_id"`
	CloudWatchMetricNamespace basetypes.StringValue `tfsdk:"cloud_watch_metric_namespace"`
	CloudWatchMetricsFields   basetypes.ListValue   `tfsdk:"cloud_watch_metrics_fields"`
}

// NgfirewallEgressNatConfig represents a nested structure within the Ngfirewalls model
type NgfirewallEgressNatConfig struct {
	Enabled  basetypes.BoolValue   `tfsdk:"enabled"`
	Settings basetypes.ObjectValue `tfsdk:"settings"`
}

// NgfirewallEgressNatSettings represents a nested structure within the Ngfirewalls model
type NgfirewallEgressNatSettings struct {
	IPAMPoolId basetypes.StringValue `tfsdk:"ipam_pool_id"`
	IPPoolType basetypes.StringValue `tfsdk:"ip_pool_type"`
}

// NgfirewallEndpoint represents a nested structure within the Ngfirewalls model
type NgfirewallEndpoint struct {
	AccountId        basetypes.StringValue `tfsdk:"account_id"`
	EgressNATEnabled basetypes.BoolValue   `tfsdk:"egress_nat_enabled"`
	EndpointId       basetypes.StringValue `tfsdk:"endpoint_id"`
	Mode             basetypes.StringValue `tfsdk:"mode"`
	Prefixes         basetypes.ObjectValue `tfsdk:"prefixes"`
	RejectedReason   basetypes.StringValue `tfsdk:"rejected_reason"`
	Status           basetypes.StringValue `tfsdk:"status"`
	SubnetId         basetypes.StringValue `tfsdk:"subnet_id"`
	VpcId            basetypes.StringValue `tfsdk:"vpc_id"`
	ZoneId           basetypes.StringValue `tfsdk:"zone_id"`
}

// NgfirewallEndpointPrefixes represents a nested structure within the Ngfirewalls model
type NgfirewallEndpointPrefixes struct {
	PrivatePrefix basetypes.ObjectValue `tfsdk:"private_prefix"`
	PublicPrefix  basetypes.ObjectValue `tfsdk:"public_prefix"`
}

// NgfirewallPrefixList represents a nested structure within the Ngfirewalls model
type NgfirewallPrefixList struct {
	Cidrs basetypes.ListValue `tfsdk:"cidrs"`
}

// NgfirewallGwlbConfig represents a nested structure within the Ngfirewalls model
type NgfirewallGwlbConfig struct {
	DeregistrationDelay                basetypes.Int64Value `tfsdk:"deregistration_delay"`
	RejectedFlowCountAlertThreshold    basetypes.Int64Value `tfsdk:"rejected_flow_count_alert_threshold"`
	RejectedFlowCountTCPAlertThreshold basetypes.Int64Value `tfsdk:"rejected_flow_count_tcp_alert_threshold"`
	SessionRebalanceEnabled            basetypes.BoolValue  `tfsdk:"session_rebalance_enabled"`
	TCPIdleTimeout                     basetypes.Int64Value `tfsdk:"tcp_idle_timeout"`
}

// NgfirewallIpv6Config represents a nested structure within the Ngfirewalls model
type NgfirewallIpv6Config struct {
	Enabled basetypes.BoolValue `tfsdk:"enabled"`
}

// NgfirewallLogConfig represents a nested structure within the Ngfirewalls model
type NgfirewallLogConfig struct {
	AccountId            basetypes.StringValue `tfsdk:"account_id"`
	LogDestination       basetypes.StringValue `tfsdk:"log_destination"`
	LogDestinationRegion basetypes.StringValue `tfsdk:"log_destination_region"`
	LogDestinationType   basetypes.StringValue `tfsdk:"log_destination_type"`
	LogType              basetypes.ListValue   `tfsdk:"log_type"`
	RoleType             basetypes.StringValue `tfsdk:"role_type"`
}

// NgfirewallNotification represents a nested structure within the Ngfirewalls model
type NgfirewallNotification struct {
	ExpiryTime basetypes.Int64Value  `tfsdk:"expiry_time"`
	Id         basetypes.StringValue `tfsdk:"id"`
	Link       basetypes.ObjectValue `tfsdk:"link"`
	Msg        basetypes.StringValue `tfsdk:"msg"`
}

// NgfirewallNotificationLink represents a nested structure within the Ngfirewalls model
type NgfirewallNotificationLink struct {
	Target basetypes.StringValue `tfsdk:"target"`
	Text   basetypes.StringValue `tfsdk:"text"`
}

// NgfirewallPrivateAccess represents a nested structure within the Ngfirewalls model
type NgfirewallPrivateAccess struct {
	ResourceID basetypes.StringValue `tfsdk:"resource_id"`
	Type       basetypes.StringValue `tfsdk:"type"`
}

// NgfirewallTag represents a nested structure within the Ngfirewalls model
type NgfirewallTag struct {
	Key   basetypes.StringValue `tfsdk:"key"`
	Value basetypes.StringValue `tfsdk:"value"`
}

// NgfirewallUserId represents a nested structure within the Ngfirewalls model
type NgfirewallUserId struct {
	AgentName                   basetypes.StringValue `tfsdk:"agent_name"`
	CFTURLLink                  basetypes.StringValue `tfsdk:"cfturl_link"`
	CollectorName               basetypes.StringValue `tfsdk:"collector_name"`
	CustomIncludeExcludeNetwork basetypes.ListValue   `tfsdk:"custom_include_exclude_network"`
	Enabled                     basetypes.BoolValue   `tfsdk:"enabled"`
	EndpointDNS                 basetypes.StringValue `tfsdk:"endpoint_dns"`
	Port                        basetypes.Int64Value  `tfsdk:"port"`
	SecretKeyARN                basetypes.StringValue `tfsdk:"secret_key_arn"`
	UserIDStatus                basetypes.StringValue `tfsdk:"user_id_status"`
}

// NgfirewallUserIdNetwork represents a nested structure within the Ngfirewalls model
type NgfirewallUserIdNetwork struct {
	DiscoveryInclude basetypes.BoolValue   `tfsdk:"discovery_include"`
	Enabled          basetypes.BoolValue   `tfsdk:"enabled"`
	Name             basetypes.StringValue `tfsdk:"name"`
	NetworkAddress   basetypes.StringValue `tfsdk:"network_address"`
}

// NgfirewallStatus represents a nested structure within the Ngfirewalls model
type NgfirewallStatus struct {
	DeviceRuleStackCommitInfo   basetypes.ObjectValue `tfsdk:"device_rule_stack_commit_info"`
	DeviceRuleStackCommitState  basetypes.StringValue `tfsdk:"device_rule_stack_commit_state"`
	DeviceRuleStackCommitStatus basetypes.StringValue `tfsdk:"device_rule_stack_commit_status"`
	FailureReason               basetypes.StringValue `tfsdk:"failure_reason"`
	FirewallStatus              basetypes.StringValue `tfsdk:"firewall_status"`
	GlobalRuleStackCommitInfo   basetypes.ObjectValue `tfsdk:"global_rule_stack_commit_info"`
	GlobalRuleStackStatus       basetypes.StringValue `tfsdk:"global_rule_stack_status"`
	PublicIPs                   basetypes.ListValue   `tfsdk:"public_i_ps"`
	RuleStackCommitInfo         basetypes.ObjectValue `tfsdk:"rule_stack_commit_info"`
	RuleStackStatus             basetypes.StringValue `tfsdk:"rule_stack_status"`
	SCMAssocStatus              basetypes.BoolValue   `tfsdk:"scm_assoc_status"`
}

// NgfirewallCommitInfo represents a nested structure within the Ngfirewalls model
type NgfirewallCommitInfo struct {
	CommitMessages basetypes.ListValue   `tfsdk:"commit_messages"`
	CommitTS       basetypes.StringValue `tfsdk:"commit_ts"`
}

// NgfirewallPublicIp represents a nested structure within the Ngfirewalls model
type NgfirewallPublicIp struct {
	IPAddress basetypes.StringValue `tfsdk:"ip_address"`
	IPSource  basetypes.StringValue `tfsdk:"ip_source"`
	IPStatus  basetypes.StringValue `tfsdk:"ip_status"`
}

// NgfirewallResponseStatus represents a nested structure within the Ngfirewalls model
type NgfirewallResponseStatus struct {
	ErrorCode basetypes.Int64Value  `tfsdk:"error_code"`
	Reason    basetypes.StringValue `tfsdk:"reason"`
}

// NgfirewallCreateRequest represents a nested structure within the Ngfirewalls model
type NgfirewallCreateRequest struct {
	Tfid               types.String          `tfsdk:"tfid"`
	AllowListAccounts  basetypes.ListValue   `tfsdk:"allow_list_accounts"`
	ChangeProtection   basetypes.ListValue   `tfsdk:"change_protection"`
	CustomerZoneIdList basetypes.ListValue   `tfsdk:"customer_zone_id_list"`
	Description        basetypes.StringValue `tfsdk:"description"`
	EgressNAT          basetypes.ObjectValue `tfsdk:"egress_nat"`
	Endpoints          basetypes.ListValue   `tfsdk:"endpoints"`
	GwlbTcpIdleTimeout basetypes.Int64Value  `tfsdk:"gwlb_tcp_idle_timeout"`
	LinkId             basetypes.StringValue `tfsdk:"link_id"`
	Mode               basetypes.StringValue `tfsdk:"mode"`
	PrivateAccess      basetypes.ObjectValue `tfsdk:"private_access"`
	ScmLabels          basetypes.ListValue   `tfsdk:"scm_labels"`
	Source             basetypes.StringValue `tfsdk:"source"`
	Tags               basetypes.ListValue   `tfsdk:"tags"`
	Tier               basetypes.StringValue `tfsdk:"tier"`
	UserID             basetypes.ObjectValue `tfsdk:"user_id"`
	Region             basetypes.StringValue `tfsdk:"region"`
	Id                 basetypes.StringValue `tfsdk:"id"`
}

// NgfirewallsListResponse represents a nested structure within the Ngfirewalls model
type NgfirewallsListResponse struct {
	Tfid           types.String          `tfsdk:"tfid"`
	Response       basetypes.ObjectValue `tfsdk:"response"`
	ResponseStatus basetypes.ObjectValue `tfsdk:"response_status"`
}

// NgfirewallsListResponseBody represents a nested structure within the Ngfirewalls model
type NgfirewallsListResponseBody struct {
	FirewallsDescribe basetypes.ListValue `tfsdk:"firewalls_describe"`
}

// AttrTypes defines the attribute types for the Ngfirewalls model.
func (o Ngfirewalls) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid": basetypes.StringType{},
		"response": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"firewall": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"advanced_threat_log": basetypes.BoolType{},
						"allow_list_accounts": basetypes.ListType{ElemType: basetypes.StringType{}},
						"change_protection":   basetypes.ListType{ElemType: basetypes.StringType{}},
						"cloudwatch_metrics": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"account_id":                   basetypes.StringType{},
								"cloud_watch_metric_namespace": basetypes.StringType{},
								"cloud_watch_metrics_fields":   basetypes.ListType{ElemType: basetypes.StringType{}},
							},
						},
						"customer_zone_id_list":           basetypes.ListType{ElemType: basetypes.StringType{}},
						"deployment_update_token":         basetypes.StringType{},
						"description":                     basetypes.StringType{},
						"device_rule_stack_commit_status": basetypes.StringType{},
						"egress_nat": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"enabled": basetypes.BoolType{},
								"settings": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"ip_pool_type": basetypes.StringType{},
										"ipam_pool_id": basetypes.StringType{},
									},
								},
							},
						},
						"endpoint_service_name": basetypes.StringType{},
						"endpoints": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"account_id":         basetypes.StringType{},
								"egress_nat_enabled": basetypes.BoolType{},
								"endpoint_id":        basetypes.StringType{},
								"mode":               basetypes.StringType{},
								"prefixes": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"private_prefix": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"cidrs": basetypes.ListType{ElemType: basetypes.StringType{}},
											},
										},
										"public_prefix": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"cidrs": basetypes.ListType{ElemType: basetypes.StringType{}},
											},
										},
									},
								},
								"rejected_reason": basetypes.StringType{},
								"status":          basetypes.StringType{},
								"subnet_id":       basetypes.StringType{},
								"vpc_id":          basetypes.StringType{},
								"zone_id":         basetypes.StringType{},
							},
						}},
						"firewall_id": basetypes.StringType{},
						"gwlb": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"deregistration_delay":                    basetypes.Int64Type{},
								"rejected_flow_count_alert_threshold":     basetypes.Int64Type{},
								"rejected_flow_count_tcp_alert_threshold": basetypes.Int64Type{},
								"session_rebalance_enabled":               basetypes.BoolType{},
								"tcp_idle_timeout":                        basetypes.Int64Type{},
							},
						},
						"ipv6": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"enabled": basetypes.BoolType{},
							},
						},
						"link_id":     basetypes.StringType{},
						"link_status": basetypes.StringType{},
						"log_config": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"account_id":             basetypes.StringType{},
								"log_destination":        basetypes.StringType{},
								"log_destination_region": basetypes.StringType{},
								"log_destination_type":   basetypes.StringType{},
								"log_type":               basetypes.ListType{ElemType: basetypes.StringType{}},
								"role_type":              basetypes.StringType{},
							},
						},
						"notifications": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"expiry_time": basetypes.Int64Type{},
								"id":          basetypes.StringType{},
								"link": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"target": basetypes.StringType{},
										"text":   basetypes.StringType{},
									},
								},
								"msg": basetypes.StringType{},
							},
						}},
						"private_access": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"resource_id": basetypes.StringType{},
								"type":        basetypes.StringType{},
							},
						},
						"region":           basetypes.StringType{},
						"software_version": basetypes.StringType{},
						"tags": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"key":   basetypes.StringType{},
								"value": basetypes.StringType{},
							},
						}},
						"tier":         basetypes.StringType{},
						"update_token": basetypes.StringType{},
						"user_id": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"agent_name":     basetypes.StringType{},
								"cfturl_link":    basetypes.StringType{},
								"collector_name": basetypes.StringType{},
								"custom_include_exclude_network": basetypes.ListType{ElemType: basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"discovery_include": basetypes.BoolType{},
										"enabled":           basetypes.BoolType{},
										"name":              basetypes.StringType{},
										"network_address":   basetypes.StringType{},
									},
								}},
								"enabled":        basetypes.BoolType{},
								"endpoint_dns":   basetypes.StringType{},
								"port":           basetypes.Int64Type{},
								"secret_key_arn": basetypes.StringType{},
								"user_id_status": basetypes.StringType{},
							},
						},
					},
				},
				"status": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"device_rule_stack_commit_info": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"commit_messages": basetypes.ListType{ElemType: basetypes.StringType{}},
								"commit_ts":       basetypes.StringType{},
							},
						},
						"device_rule_stack_commit_state":  basetypes.StringType{},
						"device_rule_stack_commit_status": basetypes.StringType{},
						"failure_reason":                  basetypes.StringType{},
						"firewall_status":                 basetypes.StringType{},
						"global_rule_stack_commit_info": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"commit_messages": basetypes.ListType{ElemType: basetypes.StringType{}},
								"commit_ts":       basetypes.StringType{},
							},
						},
						"global_rule_stack_status": basetypes.StringType{},
						"public_i_ps": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"ip_address": basetypes.StringType{},
								"ip_source":  basetypes.StringType{},
								"ip_status":  basetypes.StringType{},
							},
						}},
						"rule_stack_commit_info": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"commit_messages": basetypes.ListType{ElemType: basetypes.StringType{}},
								"commit_ts":       basetypes.StringType{},
							},
						},
						"rule_stack_status": basetypes.StringType{},
						"scm_assoc_status":  basetypes.BoolType{},
					},
				},
			},
		},
		"response_status": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"error_code": basetypes.Int64Type{},
				"reason":     basetypes.StringType{},
			},
		},
		"region": basetypes.StringType{},
		"id":     basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of Ngfirewalls objects.
func (o Ngfirewalls) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the NgfirewallResponseBody model.
func (o NgfirewallResponseBody) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"firewall": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"advanced_threat_log": basetypes.BoolType{},
				"allow_list_accounts": basetypes.ListType{ElemType: basetypes.StringType{}},
				"change_protection":   basetypes.ListType{ElemType: basetypes.StringType{}},
				"cloudwatch_metrics": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"account_id":                   basetypes.StringType{},
						"cloud_watch_metric_namespace": basetypes.StringType{},
						"cloud_watch_metrics_fields":   basetypes.ListType{ElemType: basetypes.StringType{}},
					},
				},
				"customer_zone_id_list":           basetypes.ListType{ElemType: basetypes.StringType{}},
				"deployment_update_token":         basetypes.StringType{},
				"description":                     basetypes.StringType{},
				"device_rule_stack_commit_status": basetypes.StringType{},
				"egress_nat": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"enabled": basetypes.BoolType{},
						"settings": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"ip_pool_type": basetypes.StringType{},
								"ipam_pool_id": basetypes.StringType{},
							},
						},
					},
				},
				"endpoint_service_name": basetypes.StringType{},
				"endpoints": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"account_id":         basetypes.StringType{},
						"egress_nat_enabled": basetypes.BoolType{},
						"endpoint_id":        basetypes.StringType{},
						"mode":               basetypes.StringType{},
						"prefixes": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"private_prefix": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"cidrs": basetypes.ListType{ElemType: basetypes.StringType{}},
									},
								},
								"public_prefix": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"cidrs": basetypes.ListType{ElemType: basetypes.StringType{}},
									},
								},
							},
						},
						"rejected_reason": basetypes.StringType{},
						"status":          basetypes.StringType{},
						"subnet_id":       basetypes.StringType{},
						"vpc_id":          basetypes.StringType{},
						"zone_id":         basetypes.StringType{},
					},
				}},
				"firewall_id": basetypes.StringType{},
				"gwlb": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"deregistration_delay":                    basetypes.Int64Type{},
						"rejected_flow_count_alert_threshold":     basetypes.Int64Type{},
						"rejected_flow_count_tcp_alert_threshold": basetypes.Int64Type{},
						"session_rebalance_enabled":               basetypes.BoolType{},
						"tcp_idle_timeout":                        basetypes.Int64Type{},
					},
				},
				"ipv6": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"enabled": basetypes.BoolType{},
					},
				},
				"link_id":     basetypes.StringType{},
				"link_status": basetypes.StringType{},
				"log_config": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"account_id":             basetypes.StringType{},
						"log_destination":        basetypes.StringType{},
						"log_destination_region": basetypes.StringType{},
						"log_destination_type":   basetypes.StringType{},
						"log_type":               basetypes.ListType{ElemType: basetypes.StringType{}},
						"role_type":              basetypes.StringType{},
					},
				},
				"notifications": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"expiry_time": basetypes.Int64Type{},
						"id":          basetypes.StringType{},
						"link": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"target": basetypes.StringType{},
								"text":   basetypes.StringType{},
							},
						},
						"msg": basetypes.StringType{},
					},
				}},
				"private_access": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"resource_id": basetypes.StringType{},
						"type":        basetypes.StringType{},
					},
				},
				"region":           basetypes.StringType{},
				"software_version": basetypes.StringType{},
				"tags": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"key":   basetypes.StringType{},
						"value": basetypes.StringType{},
					},
				}},
				"tier":         basetypes.StringType{},
				"update_token": basetypes.StringType{},
				"user_id": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"agent_name":     basetypes.StringType{},
						"cfturl_link":    basetypes.StringType{},
						"collector_name": basetypes.StringType{},
						"custom_include_exclude_network": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"discovery_include": basetypes.BoolType{},
								"enabled":           basetypes.BoolType{},
								"name":              basetypes.StringType{},
								"network_address":   basetypes.StringType{},
							},
						}},
						"enabled":        basetypes.BoolType{},
						"endpoint_dns":   basetypes.StringType{},
						"port":           basetypes.Int64Type{},
						"secret_key_arn": basetypes.StringType{},
						"user_id_status": basetypes.StringType{},
					},
				},
			},
		},
		"status": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"device_rule_stack_commit_info": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"commit_messages": basetypes.ListType{ElemType: basetypes.StringType{}},
						"commit_ts":       basetypes.StringType{},
					},
				},
				"device_rule_stack_commit_state":  basetypes.StringType{},
				"device_rule_stack_commit_status": basetypes.StringType{},
				"failure_reason":                  basetypes.StringType{},
				"firewall_status":                 basetypes.StringType{},
				"global_rule_stack_commit_info": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"commit_messages": basetypes.ListType{ElemType: basetypes.StringType{}},
						"commit_ts":       basetypes.StringType{},
					},
				},
				"global_rule_stack_status": basetypes.StringType{},
				"public_i_ps": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"ip_address": basetypes.StringType{},
						"ip_source":  basetypes.StringType{},
						"ip_status":  basetypes.StringType{},
					},
				}},
				"rule_stack_commit_info": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"commit_messages": basetypes.ListType{ElemType: basetypes.StringType{}},
						"commit_ts":       basetypes.StringType{},
					},
				},
				"rule_stack_status": basetypes.StringType{},
				"scm_assoc_status":  basetypes.BoolType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of NgfirewallResponseBody objects.
func (o NgfirewallResponseBody) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the NgfirewallFirewall model.
func (o NgfirewallFirewall) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"advanced_threat_log": basetypes.BoolType{},
		"allow_list_accounts": basetypes.ListType{ElemType: basetypes.StringType{}},
		"change_protection":   basetypes.ListType{ElemType: basetypes.StringType{}},
		"cloudwatch_metrics": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"account_id":                   basetypes.StringType{},
				"cloud_watch_metric_namespace": basetypes.StringType{},
				"cloud_watch_metrics_fields":   basetypes.ListType{ElemType: basetypes.StringType{}},
			},
		},
		"customer_zone_id_list":           basetypes.ListType{ElemType: basetypes.StringType{}},
		"deployment_update_token":         basetypes.StringType{},
		"description":                     basetypes.StringType{},
		"device_rule_stack_commit_status": basetypes.StringType{},
		"egress_nat": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"enabled": basetypes.BoolType{},
				"settings": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"ip_pool_type": basetypes.StringType{},
						"ipam_pool_id": basetypes.StringType{},
					},
				},
			},
		},
		"endpoint_service_name": basetypes.StringType{},
		"endpoints": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"account_id":         basetypes.StringType{},
				"egress_nat_enabled": basetypes.BoolType{},
				"endpoint_id":        basetypes.StringType{},
				"mode":               basetypes.StringType{},
				"prefixes": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"private_prefix": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"cidrs": basetypes.ListType{ElemType: basetypes.StringType{}},
							},
						},
						"public_prefix": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"cidrs": basetypes.ListType{ElemType: basetypes.StringType{}},
							},
						},
					},
				},
				"rejected_reason": basetypes.StringType{},
				"status":          basetypes.StringType{},
				"subnet_id":       basetypes.StringType{},
				"vpc_id":          basetypes.StringType{},
				"zone_id":         basetypes.StringType{},
			},
		}},
		"firewall_id": basetypes.StringType{},
		"gwlb": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"deregistration_delay":                    basetypes.Int64Type{},
				"rejected_flow_count_alert_threshold":     basetypes.Int64Type{},
				"rejected_flow_count_tcp_alert_threshold": basetypes.Int64Type{},
				"session_rebalance_enabled":               basetypes.BoolType{},
				"tcp_idle_timeout":                        basetypes.Int64Type{},
			},
		},
		"ipv6": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"enabled": basetypes.BoolType{},
			},
		},
		"link_id":     basetypes.StringType{},
		"link_status": basetypes.StringType{},
		"log_config": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"account_id":             basetypes.StringType{},
				"log_destination":        basetypes.StringType{},
				"log_destination_region": basetypes.StringType{},
				"log_destination_type":   basetypes.StringType{},
				"log_type":               basetypes.ListType{ElemType: basetypes.StringType{}},
				"role_type":              basetypes.StringType{},
			},
		},
		"notifications": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"expiry_time": basetypes.Int64Type{},
				"id":          basetypes.StringType{},
				"link": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"target": basetypes.StringType{},
						"text":   basetypes.StringType{},
					},
				},
				"msg": basetypes.StringType{},
			},
		}},
		"private_access": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"resource_id": basetypes.StringType{},
				"type":        basetypes.StringType{},
			},
		},
		"region":           basetypes.StringType{},
		"software_version": basetypes.StringType{},
		"tags": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"key":   basetypes.StringType{},
				"value": basetypes.StringType{},
			},
		}},
		"tier":         basetypes.StringType{},
		"update_token": basetypes.StringType{},
		"user_id": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"agent_name":     basetypes.StringType{},
				"cfturl_link":    basetypes.StringType{},
				"collector_name": basetypes.StringType{},
				"custom_include_exclude_network": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"discovery_include": basetypes.BoolType{},
						"enabled":           basetypes.BoolType{},
						"name":              basetypes.StringType{},
						"network_address":   basetypes.StringType{},
					},
				}},
				"enabled":        basetypes.BoolType{},
				"endpoint_dns":   basetypes.StringType{},
				"port":           basetypes.Int64Type{},
				"secret_key_arn": basetypes.StringType{},
				"user_id_status": basetypes.StringType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of NgfirewallFirewall objects.
func (o NgfirewallFirewall) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the NgfirewallCloudwatchMetrics model.
func (o NgfirewallCloudwatchMetrics) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"account_id":                   basetypes.StringType{},
		"cloud_watch_metric_namespace": basetypes.StringType{},
		"cloud_watch_metrics_fields":   basetypes.ListType{ElemType: basetypes.StringType{}},
	}
}

// AttrType returns the attribute type for a list of NgfirewallCloudwatchMetrics objects.
func (o NgfirewallCloudwatchMetrics) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the NgfirewallEgressNatConfig model.
func (o NgfirewallEgressNatConfig) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"enabled": basetypes.BoolType{},
		"settings": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"ip_pool_type": basetypes.StringType{},
				"ipam_pool_id": basetypes.StringType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of NgfirewallEgressNatConfig objects.
func (o NgfirewallEgressNatConfig) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the NgfirewallEgressNatSettings model.
func (o NgfirewallEgressNatSettings) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"ipam_pool_id": basetypes.StringType{},
		"ip_pool_type": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of NgfirewallEgressNatSettings objects.
func (o NgfirewallEgressNatSettings) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the NgfirewallEndpoint model.
func (o NgfirewallEndpoint) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"account_id":         basetypes.StringType{},
		"egress_nat_enabled": basetypes.BoolType{},
		"endpoint_id":        basetypes.StringType{},
		"mode":               basetypes.StringType{},
		"prefixes": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"private_prefix": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"cidrs": basetypes.ListType{ElemType: basetypes.StringType{}},
					},
				},
				"public_prefix": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"cidrs": basetypes.ListType{ElemType: basetypes.StringType{}},
					},
				},
			},
		},
		"rejected_reason": basetypes.StringType{},
		"status":          basetypes.StringType{},
		"subnet_id":       basetypes.StringType{},
		"vpc_id":          basetypes.StringType{},
		"zone_id":         basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of NgfirewallEndpoint objects.
func (o NgfirewallEndpoint) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the NgfirewallEndpointPrefixes model.
func (o NgfirewallEndpointPrefixes) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"private_prefix": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"cidrs": basetypes.ListType{ElemType: basetypes.StringType{}},
			},
		},
		"public_prefix": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"cidrs": basetypes.ListType{ElemType: basetypes.StringType{}},
			},
		},
	}
}

// AttrType returns the attribute type for a list of NgfirewallEndpointPrefixes objects.
func (o NgfirewallEndpointPrefixes) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the NgfirewallPrefixList model.
func (o NgfirewallPrefixList) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"cidrs": basetypes.ListType{ElemType: basetypes.StringType{}},
	}
}

// AttrType returns the attribute type for a list of NgfirewallPrefixList objects.
func (o NgfirewallPrefixList) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the NgfirewallGwlbConfig model.
func (o NgfirewallGwlbConfig) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"deregistration_delay":                    basetypes.Int64Type{},
		"rejected_flow_count_alert_threshold":     basetypes.Int64Type{},
		"rejected_flow_count_tcp_alert_threshold": basetypes.Int64Type{},
		"session_rebalance_enabled":               basetypes.BoolType{},
		"tcp_idle_timeout":                        basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of NgfirewallGwlbConfig objects.
func (o NgfirewallGwlbConfig) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the NgfirewallIpv6Config model.
func (o NgfirewallIpv6Config) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"enabled": basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of NgfirewallIpv6Config objects.
func (o NgfirewallIpv6Config) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the NgfirewallLogConfig model.
func (o NgfirewallLogConfig) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"account_id":             basetypes.StringType{},
		"log_destination":        basetypes.StringType{},
		"log_destination_region": basetypes.StringType{},
		"log_destination_type":   basetypes.StringType{},
		"log_type":               basetypes.ListType{ElemType: basetypes.StringType{}},
		"role_type":              basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of NgfirewallLogConfig objects.
func (o NgfirewallLogConfig) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the NgfirewallNotification model.
func (o NgfirewallNotification) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"expiry_time": basetypes.Int64Type{},
		"id":          basetypes.StringType{},
		"link": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"target": basetypes.StringType{},
				"text":   basetypes.StringType{},
			},
		},
		"msg": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of NgfirewallNotification objects.
func (o NgfirewallNotification) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the NgfirewallNotificationLink model.
func (o NgfirewallNotificationLink) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"target": basetypes.StringType{},
		"text":   basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of NgfirewallNotificationLink objects.
func (o NgfirewallNotificationLink) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the NgfirewallPrivateAccess model.
func (o NgfirewallPrivateAccess) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"resource_id": basetypes.StringType{},
		"type":        basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of NgfirewallPrivateAccess objects.
func (o NgfirewallPrivateAccess) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the NgfirewallTag model.
func (o NgfirewallTag) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"key":   basetypes.StringType{},
		"value": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of NgfirewallTag objects.
func (o NgfirewallTag) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the NgfirewallUserId model.
func (o NgfirewallUserId) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"agent_name":     basetypes.StringType{},
		"cfturl_link":    basetypes.StringType{},
		"collector_name": basetypes.StringType{},
		"custom_include_exclude_network": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"discovery_include": basetypes.BoolType{},
				"enabled":           basetypes.BoolType{},
				"name":              basetypes.StringType{},
				"network_address":   basetypes.StringType{},
			},
		}},
		"enabled":        basetypes.BoolType{},
		"endpoint_dns":   basetypes.StringType{},
		"port":           basetypes.Int64Type{},
		"secret_key_arn": basetypes.StringType{},
		"user_id_status": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of NgfirewallUserId objects.
func (o NgfirewallUserId) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the NgfirewallUserIdNetwork model.
func (o NgfirewallUserIdNetwork) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"discovery_include": basetypes.BoolType{},
		"enabled":           basetypes.BoolType{},
		"name":              basetypes.StringType{},
		"network_address":   basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of NgfirewallUserIdNetwork objects.
func (o NgfirewallUserIdNetwork) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the NgfirewallStatus model.
func (o NgfirewallStatus) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"device_rule_stack_commit_info": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"commit_messages": basetypes.ListType{ElemType: basetypes.StringType{}},
				"commit_ts":       basetypes.StringType{},
			},
		},
		"device_rule_stack_commit_state":  basetypes.StringType{},
		"device_rule_stack_commit_status": basetypes.StringType{},
		"failure_reason":                  basetypes.StringType{},
		"firewall_status":                 basetypes.StringType{},
		"global_rule_stack_commit_info": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"commit_messages": basetypes.ListType{ElemType: basetypes.StringType{}},
				"commit_ts":       basetypes.StringType{},
			},
		},
		"global_rule_stack_status": basetypes.StringType{},
		"public_i_ps": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"ip_address": basetypes.StringType{},
				"ip_source":  basetypes.StringType{},
				"ip_status":  basetypes.StringType{},
			},
		}},
		"rule_stack_commit_info": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"commit_messages": basetypes.ListType{ElemType: basetypes.StringType{}},
				"commit_ts":       basetypes.StringType{},
			},
		},
		"rule_stack_status": basetypes.StringType{},
		"scm_assoc_status":  basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of NgfirewallStatus objects.
func (o NgfirewallStatus) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the NgfirewallCommitInfo model.
func (o NgfirewallCommitInfo) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"commit_messages": basetypes.ListType{ElemType: basetypes.StringType{}},
		"commit_ts":       basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of NgfirewallCommitInfo objects.
func (o NgfirewallCommitInfo) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the NgfirewallPublicIp model.
func (o NgfirewallPublicIp) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"ip_address": basetypes.StringType{},
		"ip_source":  basetypes.StringType{},
		"ip_status":  basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of NgfirewallPublicIp objects.
func (o NgfirewallPublicIp) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the NgfirewallResponseStatus model.
func (o NgfirewallResponseStatus) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"error_code": basetypes.Int64Type{},
		"reason":     basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of NgfirewallResponseStatus objects.
func (o NgfirewallResponseStatus) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the NgfirewallCreateRequest model.
func (o NgfirewallCreateRequest) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":                  basetypes.StringType{},
		"allow_list_accounts":   basetypes.ListType{ElemType: basetypes.StringType{}},
		"change_protection":     basetypes.ListType{ElemType: basetypes.StringType{}},
		"customer_zone_id_list": basetypes.ListType{ElemType: basetypes.StringType{}},
		"description":           basetypes.StringType{},
		"egress_nat": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"enabled": basetypes.BoolType{},
				"settings": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"ip_pool_type": basetypes.StringType{},
						"ipam_pool_id": basetypes.StringType{},
					},
				},
			},
		},
		"endpoints": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"account_id":         basetypes.StringType{},
				"egress_nat_enabled": basetypes.BoolType{},
				"endpoint_id":        basetypes.StringType{},
				"mode":               basetypes.StringType{},
				"prefixes": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"private_prefix": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"cidrs": basetypes.ListType{ElemType: basetypes.StringType{}},
							},
						},
						"public_prefix": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"cidrs": basetypes.ListType{ElemType: basetypes.StringType{}},
							},
						},
					},
				},
				"rejected_reason": basetypes.StringType{},
				"status":          basetypes.StringType{},
				"subnet_id":       basetypes.StringType{},
				"vpc_id":          basetypes.StringType{},
				"zone_id":         basetypes.StringType{},
			},
		}},
		"gwlb_tcp_idle_timeout": basetypes.Int64Type{},
		"link_id":               basetypes.StringType{},
		"mode":                  basetypes.StringType{},
		"private_access": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"resource_id": basetypes.StringType{},
				"type":        basetypes.StringType{},
			},
		},
		"scm_labels": basetypes.ListType{ElemType: basetypes.StringType{}},
		"source":     basetypes.StringType{},
		"tags": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"key":   basetypes.StringType{},
				"value": basetypes.StringType{},
			},
		}},
		"tier": basetypes.StringType{},
		"user_id": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"agent_name":     basetypes.StringType{},
				"cfturl_link":    basetypes.StringType{},
				"collector_name": basetypes.StringType{},
				"custom_include_exclude_network": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"discovery_include": basetypes.BoolType{},
						"enabled":           basetypes.BoolType{},
						"name":              basetypes.StringType{},
						"network_address":   basetypes.StringType{},
					},
				}},
				"enabled":        basetypes.BoolType{},
				"endpoint_dns":   basetypes.StringType{},
				"port":           basetypes.Int64Type{},
				"secret_key_arn": basetypes.StringType{},
				"user_id_status": basetypes.StringType{},
			},
		},
		"region": basetypes.StringType{},
		"id":     basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of NgfirewallCreateRequest objects.
func (o NgfirewallCreateRequest) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the NgfirewallsListResponse model.
func (o NgfirewallsListResponse) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid": basetypes.StringType{},
		"response": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"firewalls_describe": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"firewall": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"advanced_threat_log": basetypes.BoolType{},
								"allow_list_accounts": basetypes.ListType{ElemType: basetypes.StringType{}},
								"change_protection":   basetypes.ListType{ElemType: basetypes.StringType{}},
								"cloudwatch_metrics": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"account_id":                   basetypes.StringType{},
										"cloud_watch_metric_namespace": basetypes.StringType{},
										"cloud_watch_metrics_fields":   basetypes.ListType{ElemType: basetypes.StringType{}},
									},
								},
								"customer_zone_id_list":           basetypes.ListType{ElemType: basetypes.StringType{}},
								"deployment_update_token":         basetypes.StringType{},
								"description":                     basetypes.StringType{},
								"device_rule_stack_commit_status": basetypes.StringType{},
								"egress_nat": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"enabled": basetypes.BoolType{},
										"settings": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"ip_pool_type": basetypes.StringType{},
												"ipam_pool_id": basetypes.StringType{},
											},
										},
									},
								},
								"endpoint_service_name": basetypes.StringType{},
								"endpoints": basetypes.ListType{ElemType: basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"account_id":         basetypes.StringType{},
										"egress_nat_enabled": basetypes.BoolType{},
										"endpoint_id":        basetypes.StringType{},
										"mode":               basetypes.StringType{},
										"prefixes": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"private_prefix": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"cidrs": basetypes.ListType{ElemType: basetypes.StringType{}},
													},
												},
												"public_prefix": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"cidrs": basetypes.ListType{ElemType: basetypes.StringType{}},
													},
												},
											},
										},
										"rejected_reason": basetypes.StringType{},
										"status":          basetypes.StringType{},
										"subnet_id":       basetypes.StringType{},
										"vpc_id":          basetypes.StringType{},
										"zone_id":         basetypes.StringType{},
									},
								}},
								"firewall_id": basetypes.StringType{},
								"gwlb": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"deregistration_delay":                    basetypes.Int64Type{},
										"rejected_flow_count_alert_threshold":     basetypes.Int64Type{},
										"rejected_flow_count_tcp_alert_threshold": basetypes.Int64Type{},
										"session_rebalance_enabled":               basetypes.BoolType{},
										"tcp_idle_timeout":                        basetypes.Int64Type{},
									},
								},
								"ipv6": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"enabled": basetypes.BoolType{},
									},
								},
								"link_id":     basetypes.StringType{},
								"link_status": basetypes.StringType{},
								"log_config": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"account_id":             basetypes.StringType{},
										"log_destination":        basetypes.StringType{},
										"log_destination_region": basetypes.StringType{},
										"log_destination_type":   basetypes.StringType{},
										"log_type":               basetypes.ListType{ElemType: basetypes.StringType{}},
										"role_type":              basetypes.StringType{},
									},
								},
								"notifications": basetypes.ListType{ElemType: basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"expiry_time": basetypes.Int64Type{},
										"id":          basetypes.StringType{},
										"link": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"target": basetypes.StringType{},
												"text":   basetypes.StringType{},
											},
										},
										"msg": basetypes.StringType{},
									},
								}},
								"private_access": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"resource_id": basetypes.StringType{},
										"type":        basetypes.StringType{},
									},
								},
								"region":           basetypes.StringType{},
								"software_version": basetypes.StringType{},
								"tags": basetypes.ListType{ElemType: basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"key":   basetypes.StringType{},
										"value": basetypes.StringType{},
									},
								}},
								"tier":         basetypes.StringType{},
								"update_token": basetypes.StringType{},
								"user_id": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"agent_name":     basetypes.StringType{},
										"cfturl_link":    basetypes.StringType{},
										"collector_name": basetypes.StringType{},
										"custom_include_exclude_network": basetypes.ListType{ElemType: basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"discovery_include": basetypes.BoolType{},
												"enabled":           basetypes.BoolType{},
												"name":              basetypes.StringType{},
												"network_address":   basetypes.StringType{},
											},
										}},
										"enabled":        basetypes.BoolType{},
										"endpoint_dns":   basetypes.StringType{},
										"port":           basetypes.Int64Type{},
										"secret_key_arn": basetypes.StringType{},
										"user_id_status": basetypes.StringType{},
									},
								},
							},
						},
						"status": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"device_rule_stack_commit_info": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"commit_messages": basetypes.ListType{ElemType: basetypes.StringType{}},
										"commit_ts":       basetypes.StringType{},
									},
								},
								"device_rule_stack_commit_state":  basetypes.StringType{},
								"device_rule_stack_commit_status": basetypes.StringType{},
								"failure_reason":                  basetypes.StringType{},
								"firewall_status":                 basetypes.StringType{},
								"global_rule_stack_commit_info": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"commit_messages": basetypes.ListType{ElemType: basetypes.StringType{}},
										"commit_ts":       basetypes.StringType{},
									},
								},
								"global_rule_stack_status": basetypes.StringType{},
								"public_i_ps": basetypes.ListType{ElemType: basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"ip_address": basetypes.StringType{},
										"ip_source":  basetypes.StringType{},
										"ip_status":  basetypes.StringType{},
									},
								}},
								"rule_stack_commit_info": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"commit_messages": basetypes.ListType{ElemType: basetypes.StringType{}},
										"commit_ts":       basetypes.StringType{},
									},
								},
								"rule_stack_status": basetypes.StringType{},
								"scm_assoc_status":  basetypes.BoolType{},
							},
						},
					},
				}},
			},
		},
		"response_status": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"error_code": basetypes.Int64Type{},
				"reason":     basetypes.StringType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of NgfirewallsListResponse objects.
func (o NgfirewallsListResponse) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the NgfirewallsListResponseBody model.
func (o NgfirewallsListResponseBody) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"firewalls_describe": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"firewall": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"advanced_threat_log": basetypes.BoolType{},
						"allow_list_accounts": basetypes.ListType{ElemType: basetypes.StringType{}},
						"change_protection":   basetypes.ListType{ElemType: basetypes.StringType{}},
						"cloudwatch_metrics": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"account_id":                   basetypes.StringType{},
								"cloud_watch_metric_namespace": basetypes.StringType{},
								"cloud_watch_metrics_fields":   basetypes.ListType{ElemType: basetypes.StringType{}},
							},
						},
						"customer_zone_id_list":           basetypes.ListType{ElemType: basetypes.StringType{}},
						"deployment_update_token":         basetypes.StringType{},
						"description":                     basetypes.StringType{},
						"device_rule_stack_commit_status": basetypes.StringType{},
						"egress_nat": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"enabled": basetypes.BoolType{},
								"settings": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"ip_pool_type": basetypes.StringType{},
										"ipam_pool_id": basetypes.StringType{},
									},
								},
							},
						},
						"endpoint_service_name": basetypes.StringType{},
						"endpoints": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"account_id":         basetypes.StringType{},
								"egress_nat_enabled": basetypes.BoolType{},
								"endpoint_id":        basetypes.StringType{},
								"mode":               basetypes.StringType{},
								"prefixes": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"private_prefix": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"cidrs": basetypes.ListType{ElemType: basetypes.StringType{}},
											},
										},
										"public_prefix": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"cidrs": basetypes.ListType{ElemType: basetypes.StringType{}},
											},
										},
									},
								},
								"rejected_reason": basetypes.StringType{},
								"status":          basetypes.StringType{},
								"subnet_id":       basetypes.StringType{},
								"vpc_id":          basetypes.StringType{},
								"zone_id":         basetypes.StringType{},
							},
						}},
						"firewall_id": basetypes.StringType{},
						"gwlb": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"deregistration_delay":                    basetypes.Int64Type{},
								"rejected_flow_count_alert_threshold":     basetypes.Int64Type{},
								"rejected_flow_count_tcp_alert_threshold": basetypes.Int64Type{},
								"session_rebalance_enabled":               basetypes.BoolType{},
								"tcp_idle_timeout":                        basetypes.Int64Type{},
							},
						},
						"ipv6": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"enabled": basetypes.BoolType{},
							},
						},
						"link_id":     basetypes.StringType{},
						"link_status": basetypes.StringType{},
						"log_config": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"account_id":             basetypes.StringType{},
								"log_destination":        basetypes.StringType{},
								"log_destination_region": basetypes.StringType{},
								"log_destination_type":   basetypes.StringType{},
								"log_type":               basetypes.ListType{ElemType: basetypes.StringType{}},
								"role_type":              basetypes.StringType{},
							},
						},
						"notifications": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"expiry_time": basetypes.Int64Type{},
								"id":          basetypes.StringType{},
								"link": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"target": basetypes.StringType{},
										"text":   basetypes.StringType{},
									},
								},
								"msg": basetypes.StringType{},
							},
						}},
						"private_access": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"resource_id": basetypes.StringType{},
								"type":        basetypes.StringType{},
							},
						},
						"region":           basetypes.StringType{},
						"software_version": basetypes.StringType{},
						"tags": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"key":   basetypes.StringType{},
								"value": basetypes.StringType{},
							},
						}},
						"tier":         basetypes.StringType{},
						"update_token": basetypes.StringType{},
						"user_id": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"agent_name":     basetypes.StringType{},
								"cfturl_link":    basetypes.StringType{},
								"collector_name": basetypes.StringType{},
								"custom_include_exclude_network": basetypes.ListType{ElemType: basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"discovery_include": basetypes.BoolType{},
										"enabled":           basetypes.BoolType{},
										"name":              basetypes.StringType{},
										"network_address":   basetypes.StringType{},
									},
								}},
								"enabled":        basetypes.BoolType{},
								"endpoint_dns":   basetypes.StringType{},
								"port":           basetypes.Int64Type{},
								"secret_key_arn": basetypes.StringType{},
								"user_id_status": basetypes.StringType{},
							},
						},
					},
				},
				"status": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"device_rule_stack_commit_info": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"commit_messages": basetypes.ListType{ElemType: basetypes.StringType{}},
								"commit_ts":       basetypes.StringType{},
							},
						},
						"device_rule_stack_commit_state":  basetypes.StringType{},
						"device_rule_stack_commit_status": basetypes.StringType{},
						"failure_reason":                  basetypes.StringType{},
						"firewall_status":                 basetypes.StringType{},
						"global_rule_stack_commit_info": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"commit_messages": basetypes.ListType{ElemType: basetypes.StringType{}},
								"commit_ts":       basetypes.StringType{},
							},
						},
						"global_rule_stack_status": basetypes.StringType{},
						"public_i_ps": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"ip_address": basetypes.StringType{},
								"ip_source":  basetypes.StringType{},
								"ip_status":  basetypes.StringType{},
							},
						}},
						"rule_stack_commit_info": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"commit_messages": basetypes.ListType{ElemType: basetypes.StringType{}},
								"commit_ts":       basetypes.StringType{},
							},
						},
						"rule_stack_status": basetypes.StringType{},
						"scm_assoc_status":  basetypes.BoolType{},
					},
				},
			},
		}},
	}
}

// AttrType returns the attribute type for a list of NgfirewallsListResponseBody objects.
func (o NgfirewallsListResponseBody) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// NgfirewallsResourceSchema defines the schema for Ngfirewalls resource
var NgfirewallsResourceSchema = schema.Schema{
	MarkdownDescription: "Ngfirewall resource",
	Attributes: map[string]schema.Attribute{
		"id": schema.StringAttribute{
			MarkdownDescription: "The ID of the Cloud NGFW firewall.",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"region": schema.StringAttribute{
			MarkdownDescription: "The AWS region where the Cloud NGFW firewall is deployed",
			Required:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"response": schema.SingleNestedAttribute{
			MarkdownDescription: "Inner response object containing firewall details and operational status.",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"firewall": schema.SingleNestedAttribute{
					MarkdownDescription: "Core firewall configuration and identity fields.",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"advanced_threat_log": schema.BoolAttribute{
							MarkdownDescription: "Whether advanced threat logging is enabled.",
							Optional:            true,
						},
						"allow_list_accounts": schema.ListAttribute{
							ElementType:         types.StringType,
							MarkdownDescription: "AWS account IDs permitted to access the firewall endpoint service.",
							Optional:            true,
						},
						"change_protection": schema.ListAttribute{
							ElementType:         types.StringType,
							MarkdownDescription: "List of attributes protected from change.",
							Optional:            true,
						},
						"cloudwatch_metrics": schema.SingleNestedAttribute{
							MarkdownDescription: "CloudWatch metrics configuration.",
							Optional:            true,
							Attributes: map[string]schema.Attribute{
								"account_id": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.LengthAtMost(12),
										stringvalidator.LengthAtLeast(12),
									},
									MarkdownDescription: "The AWS account ID where CloudWatch metrics are published.",
									Optional:            true,
								},
								"cloud_watch_metric_namespace": schema.StringAttribute{
									MarkdownDescription: "The CloudWatch metric namespace.",
									Optional:            true,
								},
								"cloud_watch_metrics_fields": schema.ListAttribute{
									ElementType:         types.StringType,
									MarkdownDescription: "The CloudWatch metric field names to publish.",
									Optional:            true,
								},
							},
						},
						"customer_zone_id_list": schema.ListAttribute{
							ElementType:         types.StringType,
							MarkdownDescription: "List of customer-defined security zone IDs.",
							Required:            true,
						},
						"deployment_update_token": schema.StringAttribute{
							MarkdownDescription: "Token used for deployment update sequencing.",
							Optional:            true,
						},
						"description": schema.StringAttribute{
							MarkdownDescription: "A human-readable description of the firewall.",
							Optional:            true,
						},
						"device_rule_stack_commit_status": schema.StringAttribute{
							MarkdownDescription: "Commit status of the device rulestack on the firewall.",
							Optional:            true,
						},
						"egress_nat": schema.SingleNestedAttribute{
							MarkdownDescription: "Egress NAT configuration for outbound traffic.",
							Optional:            true,
							Computed:            true,
							Attributes: map[string]schema.Attribute{
								"enabled": schema.BoolAttribute{
									MarkdownDescription: "Whether egress NAT is enabled on this firewall.",
									Optional:            true,
									Computed:            true,
								},
								"settings": schema.SingleNestedAttribute{
									MarkdownDescription: "Settings for egress NAT IP address management.",
									Optional:            true,
									Computed:            true,
									Attributes: map[string]schema.Attribute{
										"ip_pool_type": schema.StringAttribute{
											Validators: []validator.String{
												stringvalidator.OneOf("AWSService", "BYOIP"),
											},
											MarkdownDescription: "The IP pool type for egress NAT. Possible values are `AWSService` and `BYOIP`.",
											Optional:            true,
											Computed:            true,
										},
										"ipam_pool_id": schema.StringAttribute{
											MarkdownDescription: "The IPAM pool ID used for egress NAT IP allocation.",
											Optional:            true,
											Computed:            true,
										},
									},
								},
							},
						},
						"endpoint_service_name": schema.StringAttribute{
							MarkdownDescription: "The VPC endpoint service name (only populated in ServiceManaged mode).",
							Optional:            true,
						},
						"endpoints": schema.ListNestedAttribute{
							MarkdownDescription: "VPC endpoint configurations created for this firewall.",
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"account_id": schema.StringAttribute{
										Validators: []validator.String{
											stringvalidator.LengthAtMost(12),
											stringvalidator.LengthAtLeast(12),
										},
										MarkdownDescription: "The AWS account ID that owns this endpoint.",
										Optional:            true,
										Computed:            true,
									},
									"egress_nat_enabled": schema.BoolAttribute{
										MarkdownDescription: "Whether egress NAT is enabled for this endpoint.",
										Optional:            true,
										Computed:            true,
									},
									"endpoint_id": schema.StringAttribute{
										MarkdownDescription: "The VPC endpoint ID.",
										Optional:            true,
										Computed:            true,
									},
									"mode": schema.StringAttribute{
										Validators: []validator.String{
											stringvalidator.OneOf("ServiceManaged", "CustomerManaged"),
										},
										MarkdownDescription: "The endpoint management mode. Possible values are `ServiceManaged` and `CustomerManaged`.",
										Optional:            true,
										Computed:            true,
									},
									"prefixes": schema.SingleNestedAttribute{
										MarkdownDescription: "IP prefix lists associated with a firewall endpoint.",
										Optional:            true,
										Computed:            true,
										Attributes: map[string]schema.Attribute{
											"private_prefix": schema.SingleNestedAttribute{
												MarkdownDescription: "A list of CIDR prefixes.",
												Optional:            true,
												Computed:            true,
												Attributes: map[string]schema.Attribute{
													"cidrs": schema.ListAttribute{
														ElementType:         types.StringType,
														MarkdownDescription: "CIDR address blocks.",
														Optional:            true,
														Computed:            true,
													},
												},
											},
											"public_prefix": schema.SingleNestedAttribute{
												MarkdownDescription: "A list of CIDR prefixes.",
												Optional:            true,
												Computed:            true,
												Attributes: map[string]schema.Attribute{
													"cidrs": schema.ListAttribute{
														ElementType:         types.StringType,
														MarkdownDescription: "CIDR address blocks.",
														Optional:            true,
														Computed:            true,
													},
												},
											},
										},
									},
									"rejected_reason": schema.StringAttribute{
										MarkdownDescription: "Human-readable reason if the endpoint was rejected.",
										Optional:            true,
										Computed:            true,
									},
									"status": schema.StringAttribute{
										Validators: []validator.String{
											stringvalidator.OneOf("Pending", "Accepted", "Rejected", "Created", "Creating", "Deleting"),
										},
										MarkdownDescription: "The current provisioning status of the endpoint. Possible values are `Pending`, `Accepted`, `Rejected`, `Created`, `Creating` and `Deleting`.",
										Optional:            true,
										Computed:            true,
									},
									"subnet_id": schema.StringAttribute{
										MarkdownDescription: "The AWS subnet ID in which the endpoint resides.",
										Optional:            true,
										Computed:            true,
									},
									"vpc_id": schema.StringAttribute{
										MarkdownDescription: "The AWS VPC ID in which the endpoint resides.",
										Optional:            true,
										Computed:            true,
									},
									"zone_id": schema.StringAttribute{
										MarkdownDescription: "The availability zone ID for this endpoint.",
										Optional:            true,
										Computed:            true,
									},
								},
							},
						},
						"firewall_id": schema.StringAttribute{
							MarkdownDescription: "The unique identifier of the firewall.",
							Optional:            true,
						},
						"gwlb": schema.SingleNestedAttribute{
							MarkdownDescription: "Gateway Load Balancer configuration.",
							Optional:            true,
							Attributes: map[string]schema.Attribute{
								"deregistration_delay": schema.Int64Attribute{
									MarkdownDescription: "Deregistration delay in seconds.",
									Optional:            true,
									Computed:            true,
									Default:             int64default.StaticInt64(300),
								},
								"rejected_flow_count_alert_threshold": schema.Int64Attribute{
									MarkdownDescription: "Alert threshold for the total rejected flow count.",
									Optional:            true,
									Computed:            true,
									Default:             int64default.StaticInt64(10),
								},
								"rejected_flow_count_tcp_alert_threshold": schema.Int64Attribute{
									MarkdownDescription: "Alert threshold for the TCP rejected flow count.",
									Optional:            true,
									Computed:            true,
									Default:             int64default.StaticInt64(10),
								},
								"session_rebalance_enabled": schema.BoolAttribute{
									MarkdownDescription: "Whether session rebalancing is enabled.",
									Optional:            true,
									Computed:            true,
									Default:             booldefault.StaticBool(false),
								},
								"tcp_idle_timeout": schema.Int64Attribute{
									MarkdownDescription: "TCP idle timeout in seconds.",
									Optional:            true,
									Computed:            true,
									Default:             int64default.StaticInt64(350),
								},
							},
						},
						"ipv6": schema.SingleNestedAttribute{
							MarkdownDescription: "IPv6 configuration for the firewall.",
							Optional:            true,
							Attributes: map[string]schema.Attribute{
								"enabled": schema.BoolAttribute{
									MarkdownDescription: "Whether IPv6 is enabled on this firewall.",
									Optional:            true,
								},
							},
						},
						"link_id": schema.StringAttribute{
							MarkdownDescription: "The link ID associated with the firewall.",
							Optional:            true,
						},
						"link_status": schema.StringAttribute{
							MarkdownDescription: "The current link status of the firewall.",
							Optional:            true,
						},
						"log_config": schema.SingleNestedAttribute{
							MarkdownDescription: "Log destination configuration for the firewall.",
							Optional:            true,
							Attributes: map[string]schema.Attribute{
								"account_id": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.LengthAtMost(12),
										stringvalidator.LengthAtLeast(12),
									},
									MarkdownDescription: "The AWS account ID where logs are sent.",
									Optional:            true,
								},
								"log_destination": schema.StringAttribute{
									MarkdownDescription: "The ARN or name of the log destination resource.",
									Required:            true,
								},
								"log_destination_region": schema.StringAttribute{
									MarkdownDescription: "The AWS region of the log destination.",
									Optional:            true,
								},
								"log_destination_type": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.OneOf("S3", "CloudWatchLogs", "KinesisDataFirehose"),
									},
									MarkdownDescription: "The type of log destination. Possible values are `S3`, `CloudWatchLogs` and `KinesisDataFirehose`.",
									Required:            true,
								},
								"log_type": schema.ListAttribute{
									ElementType:         types.StringType,
									MarkdownDescription: "The types of logs to send to this destination. Possible values are `TRAFFIC`, `DECRYPTION`, `THREAT` and `USER_ID`.",
									Required:            true,
								},
								"role_type": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.OneOf("ResourceBased", "IamBased"),
									},
									MarkdownDescription: "The IAM role type used to publish logs. Possible values are `ResourceBased` and `IamBased`.",
									Required:            true,
								},
							},
						},
						"notifications": schema.ListNestedAttribute{
							MarkdownDescription: "Notifications associated with this firewall.",
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"expiry_time": schema.Int64Attribute{
										MarkdownDescription: "Unix timestamp when this notification expires.",
										Optional:            true,
									},
									"id": schema.StringAttribute{
										MarkdownDescription: "Unique identifier of the notification.",
										Optional:            true,
									},
									"link": schema.SingleNestedAttribute{
										MarkdownDescription: "A hyperlink associated with a notification.",
										Optional:            true,
										Attributes: map[string]schema.Attribute{
											"target": schema.StringAttribute{
												MarkdownDescription: "The URL target of the link.",
												Optional:            true,
											},
											"text": schema.StringAttribute{
												MarkdownDescription: "The display text of the link.",
												Optional:            true,
											},
										},
									},
									"msg": schema.StringAttribute{
										MarkdownDescription: "The notification message text.",
										Optional:            true,
									},
								},
							},
						},
						"private_access": schema.SingleNestedAttribute{
							MarkdownDescription: "Private access configuration for the firewall.",
							Optional:            true,
							Attributes: map[string]schema.Attribute{
								"resource_id": schema.StringAttribute{
									MarkdownDescription: "The resource ID of the private access endpoint.",
									Optional:            true,
								},
								"type": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.OneOf("NetworkLoadBalancer"),
									},
									MarkdownDescription: "The type of private access resource. Possible values are `NetworkLoadBalancer`.",
									Optional:            true,
								},
							},
						},
						"region": schema.StringAttribute{
							MarkdownDescription: "The AWS region where the firewall is deployed.",
							Optional:            true,
						},
						"software_version": schema.StringAttribute{
							MarkdownDescription: "The PAN-OS software version currently running on the firewall.",
							Optional:            true,
						},
						"tags": schema.ListNestedAttribute{
							Validators: []validator.List{
								listvalidator.SizeAtMost(50),
							},
							MarkdownDescription: "AWS resource tags applied to the firewall.",
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"key": schema.StringAttribute{
										Validators: []validator.String{
											stringvalidator.LengthAtMost(128),
											stringvalidator.LengthAtLeast(1),
										},
										MarkdownDescription: "The tag key.",
										Optional:            true,
									},
									"value": schema.StringAttribute{
										Validators: []validator.String{
											stringvalidator.LengthAtMost(256),
											stringvalidator.LengthAtLeast(1),
										},
										MarkdownDescription: "The tag value.",
										Optional:            true,
									},
								},
							},
						},
						"tier": schema.StringAttribute{
							MarkdownDescription: "The service tier of the firewall (base, standard, or premium).",
							Optional:            true,
						},
						"update_token": schema.StringAttribute{
							MarkdownDescription: "Optimistic locking token; must be supplied on update operations.",
							Optional:            true,
						},
						"user_id": schema.SingleNestedAttribute{
							MarkdownDescription: "User-ID configuration for user-based policy enforcement.",
							Optional:            true,
							Attributes: map[string]schema.Attribute{
								"agent_name": schema.StringAttribute{
									MarkdownDescription: "The name of the User-ID agent.",
									Optional:            true,
								},
								"cfturl_link": schema.StringAttribute{
									MarkdownDescription: "CloudFormation template URL link for User-ID setup.",
									Optional:            true,
								},
								"collector_name": schema.StringAttribute{
									MarkdownDescription: "The name of the log collector used by User-ID.",
									Optional:            true,
								},
								"custom_include_exclude_network": schema.ListNestedAttribute{
									MarkdownDescription: "Custom include/exclude network definitions for User-ID.",
									Optional:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"discovery_include": schema.BoolAttribute{
												MarkdownDescription: "Whether this network is included in User-ID discovery.",
												Optional:            true,
											},
											"enabled": schema.BoolAttribute{
												MarkdownDescription: "Whether this network entry is active.",
												Optional:            true,
											},
											"name": schema.StringAttribute{
												MarkdownDescription: "Name of the network entry.",
												Optional:            true,
											},
											"network_address": schema.StringAttribute{
												MarkdownDescription: "The network address (CIDR) for this entry.",
												Optional:            true,
											},
										},
									},
								},
								"enabled": schema.BoolAttribute{
									MarkdownDescription: "Whether User-ID is enabled.",
									Optional:            true,
								},
								"endpoint_dns": schema.StringAttribute{
									MarkdownDescription: "DNS name of the User-ID endpoint.",
									Optional:            true,
								},
								"port": schema.Int64Attribute{
									MarkdownDescription: "Port used by the User-ID agent.",
									Optional:            true,
								},
								"secret_key_arn": schema.StringAttribute{
									MarkdownDescription: "ARN of the AWS Secrets Manager secret for User-ID credentials.",
									Optional:            true,
								},
								"user_id_status": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.OneOf("EndpointPending", "EndpointFail", "CommitFail", "Enabled", "Disabled", "DisablePending", "UpdatePending", "UpdateFail"),
									},
									MarkdownDescription: "The operational status of the User-ID feature. Possible values are `EndpointPending`, `EndpointFail`, `CommitFail`, `Enabled`, `Disabled`, `DisablePending`, `UpdatePending` and `UpdateFail`.",
									Optional:            true,
								},
							},
						},
					},
				},
				"status": schema.SingleNestedAttribute{
					MarkdownDescription: "Operational status of the firewall.",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"device_rule_stack_commit_info": schema.SingleNestedAttribute{
							MarkdownDescription: "Commit information for a rulestack.",
							Optional:            true,
							Attributes: map[string]schema.Attribute{
								"commit_messages": schema.ListAttribute{
									ElementType:         types.StringType,
									MarkdownDescription: "Messages produced during the commit operation.",
									Optional:            true,
								},
								"commit_ts": schema.StringAttribute{
									MarkdownDescription: "Timestamp of the last commit operation.",
									Optional:            true,
								},
							},
						},
						"device_rule_stack_commit_state": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("Submitted", "None"),
							},
							MarkdownDescription: "The commit state of the device rulestack. Possible values are `Submitted` and `None`.",
							Optional:            true,
						},
						"device_rule_stack_commit_status": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("Running", "Uncommitted", "Precommit", "PrecommitDone", "PrevalidateDone", "Committing", "Failed", "Success"),
							},
							MarkdownDescription: "The commit status of the device rulestack. Possible values are `Running`, `Uncommitted`, `Precommit`, `PrecommitDone`, `PrevalidateDone`, `Committing`, `Failed` and `Success`.",
							Optional:            true,
						},
						"failure_reason": schema.StringAttribute{
							MarkdownDescription: "Human-readable reason for a firewall failure, if any.",
							Optional:            true,
						},
						"firewall_status": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("CREATING", "UPDATING", "DELETING", "CREATE_COMPLETE", "UPDATE_COMPLETE", "CREATE_FAIL", "UPDATE_FAIL", "DELETE_FAIL", "DELETE_COMPLETE"),
							},
							MarkdownDescription: "The operational lifecycle status of the firewall. Possible values are `CREATING`, `UPDATING`, `DELETING`, `CREATE_COMPLETE`, `UPDATE_COMPLETE`, `CREATE_FAIL`, `UPDATE_FAIL`, `DELETE_FAIL` and `DELETE_COMPLETE`.",
							Optional:            true,
						},
						"global_rule_stack_commit_info": schema.SingleNestedAttribute{
							MarkdownDescription: "Commit information for a rulestack.",
							Optional:            true,
							Attributes: map[string]schema.Attribute{
								"commit_messages": schema.ListAttribute{
									ElementType:         types.StringType,
									MarkdownDescription: "Messages produced during the commit operation.",
									Optional:            true,
								},
								"commit_ts": schema.StringAttribute{
									MarkdownDescription: "Timestamp of the last commit operation.",
									Optional:            true,
								},
							},
						},
						"global_rule_stack_status": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("Running", "Uncommitted", "Precommit", "PrecommitDone", "PrevalidateDone", "Committing", "Failed", "Success"),
							},
							MarkdownDescription: "The commit status of the global rulestack. Possible values are `Running`, `Uncommitted`, `Precommit`, `PrecommitDone`, `PrevalidateDone`, `Committing`, `Failed` and `Success`.",
							Optional:            true,
						},
						"public_i_ps": schema.ListNestedAttribute{
							MarkdownDescription: "Public IP addresses assigned to the firewall.",
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"ip_address": schema.StringAttribute{
										MarkdownDescription: "The public IP address.",
										Optional:            true,
									},
									"ip_source": schema.StringAttribute{
										MarkdownDescription: "The source of the IP address allocation.",
										Optional:            true,
									},
									"ip_status": schema.StringAttribute{
										MarkdownDescription: "The current status of this public IP.",
										Optional:            true,
									},
								},
							},
						},
						"rule_stack_commit_info": schema.SingleNestedAttribute{
							MarkdownDescription: "Commit information for a rulestack.",
							Optional:            true,
							Attributes: map[string]schema.Attribute{
								"commit_messages": schema.ListAttribute{
									ElementType:         types.StringType,
									MarkdownDescription: "Messages produced during the commit operation.",
									Optional:            true,
								},
								"commit_ts": schema.StringAttribute{
									MarkdownDescription: "Timestamp of the last commit operation.",
									Optional:            true,
								},
							},
						},
						"rule_stack_status": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("Running", "Uncommitted", "Precommit", "PrecommitDone", "PrevalidateDone", "Committing", "Failed", "Success"),
							},
							MarkdownDescription: "The commit status of the local rulestack. Possible values are `Running`, `Uncommitted`, `Precommit`, `PrecommitDone`, `PrevalidateDone`, `Committing`, `Failed` and `Success`.",
							Optional:            true,
						},
						"scm_assoc_status": schema.BoolAttribute{
							MarkdownDescription: "Whether the firewall is associated with Strata Cloud Manager.",
							Optional:            true,
						},
					},
				},
			},
		},
		"response_status": schema.SingleNestedAttribute{
			MarkdownDescription: "API-level response metadata at the top level of the envelope.",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"error_code": schema.Int64Attribute{
					MarkdownDescription: "Non-zero error code if the API request failed.",
					Optional:            true,
				},
				"reason": schema.StringAttribute{
					MarkdownDescription: "Human-readable reason for the error, if any.",
					Optional:            true,
				},
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

// NgfirewallCreateRequestResourceSchema defines the resource schema for the create-request model.
// Used when the POST body schema differs from the GET response schema.
var NgfirewallCreateRequestResourceSchema = schema.Schema{
	MarkdownDescription: "NgfirewallCreateRequest resource",
	Attributes: map[string]schema.Attribute{
		"allow_list_accounts": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "AWS account IDs permitted to access the firewall endpoint service.",
			Validators: []validator.List{
				listvalidator.SizeAtMost(300),
			},
			Optional: true,
			Computed: true,
		},
		"change_protection": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "List of attributes protected from change.",
			Optional:            true,
			Computed:            true,
		},
		"customer_zone_id_list": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "List of customer-defined security zone IDs. Required.",
			Required:            true,
		},
		"description": schema.StringAttribute{
			MarkdownDescription: "A human-readable description of the firewall.",
			Optional:            true,
		},
		"egress_nat": schema.SingleNestedAttribute{
			MarkdownDescription: "Egress NAT configuration for outbound traffic.",
			Optional:            true,
			Computed:            true,
			Attributes: map[string]schema.Attribute{
				"enabled": schema.BoolAttribute{
					MarkdownDescription: "Whether egress NAT is enabled on this firewall.",
					Optional:            true,
					Computed:            true,
				},
				"settings": schema.SingleNestedAttribute{
					MarkdownDescription: "Settings for egress NAT IP address management.",
					Optional:            true,
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"ip_pool_type": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("AWSService", "BYOIP"),
							},
							MarkdownDescription: "The IP pool type for egress NAT. Possible values are `AWSService` and `BYOIP`.",
							Optional:            true,
							Computed:            true,
						},
						"ipam_pool_id": schema.StringAttribute{
							MarkdownDescription: "The IPAM pool ID used for egress NAT IP allocation.",
							Optional:            true,
							Computed:            true,
						},
					},
				},
			},
		},
		"endpoints": schema.ListNestedAttribute{
			MarkdownDescription: "Firewall endpoint configurations. Assigned by the service when the firewall is created; can be changed with an update.",
			Optional:            true,
			Computed:            true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"account_id": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.LengthAtMost(12),
							stringvalidator.LengthAtLeast(12),
						},
						MarkdownDescription: "The AWS account ID that owns this endpoint.",
						Optional:            true,
						Computed:            true,
					},
					"egress_nat_enabled": schema.BoolAttribute{
						MarkdownDescription: "Whether egress NAT is enabled for this endpoint.",
						Optional:            true,
						Computed:            true,
					},
					"endpoint_id": schema.StringAttribute{
						MarkdownDescription: "The VPC endpoint ID.",
						Optional:            true,
						Computed:            true,
					},
					"mode": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("ServiceManaged", "CustomerManaged"),
						},
						MarkdownDescription: "The endpoint management mode. Possible values are `ServiceManaged` and `CustomerManaged`.",
						Optional:            true,
						Computed:            true,
					},
					"prefixes": schema.SingleNestedAttribute{
						MarkdownDescription: "IP prefix lists associated with a firewall endpoint.",
						Optional:            true,
						Computed:            true,
						Attributes: map[string]schema.Attribute{
							"private_prefix": schema.SingleNestedAttribute{
								MarkdownDescription: "A list of CIDR prefixes.",
								Optional:            true,
								Computed:            true,
								Attributes: map[string]schema.Attribute{
									"cidrs": schema.ListAttribute{
										ElementType:         types.StringType,
										MarkdownDescription: "CIDR address blocks.",
										Optional:            true,
										Computed:            true,
									},
								},
							},
							"public_prefix": schema.SingleNestedAttribute{
								MarkdownDescription: "A list of CIDR prefixes.",
								Optional:            true,
								Computed:            true,
								Attributes: map[string]schema.Attribute{
									"cidrs": schema.ListAttribute{
										ElementType:         types.StringType,
										MarkdownDescription: "CIDR address blocks.",
										Optional:            true,
										Computed:            true,
									},
								},
							},
						},
					},
					"rejected_reason": schema.StringAttribute{
						MarkdownDescription: "Human-readable reason if the endpoint was rejected.",
						Optional:            true,
						Computed:            true,
					},
					"status": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("Pending", "Accepted", "Rejected", "Created", "Creating", "Deleting"),
						},
						MarkdownDescription: "The current provisioning status of the endpoint. Possible values are `Pending`, `Accepted`, `Rejected`, `Created`, `Creating` and `Deleting`.",
						Optional:            true,
						Computed:            true,
					},
					"subnet_id": schema.StringAttribute{
						MarkdownDescription: "The AWS subnet ID in which the endpoint resides.",
						Optional:            true,
						Computed:            true,
					},
					"vpc_id": schema.StringAttribute{
						MarkdownDescription: "The AWS VPC ID in which the endpoint resides.",
						Optional:            true,
						Computed:            true,
					},
					"zone_id": schema.StringAttribute{
						MarkdownDescription: "The availability zone ID for this endpoint.",
						Optional:            true,
						Computed:            true,
					},
				},
			},
		},
		"gwlb_tcp_idle_timeout": schema.Int64Attribute{
			MarkdownDescription: "GWLB TCP idle timeout in seconds. Defaults to 350; can be changed with an update.",
			Optional:            true,
			Computed:            true,
		},
		"id": schema.StringAttribute{
			MarkdownDescription: "The ID of the Cloud NGFW firewall.",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"link_id": schema.StringAttribute{
			MarkdownDescription: "The link ID to associate with the firewall.",
			Optional:            true,
			Computed:            true,
		},
		"mode": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.OneOf("CNGFW"),
			},
			MarkdownDescription: "The operational mode of the firewall. Set when the firewall is created and cannot be changed afterward. Possible values are `CNGFW`.",
			Optional:            true,
		},
		"private_access": schema.SingleNestedAttribute{
			MarkdownDescription: "Private access configuration for the firewall.",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"resource_id": schema.StringAttribute{
					MarkdownDescription: "The resource ID of the private access endpoint.",
					Optional:            true,
				},
				"type": schema.StringAttribute{
					Validators: []validator.String{
						stringvalidator.OneOf("NetworkLoadBalancer"),
					},
					MarkdownDescription: "The type of private access resource. Possible values are `NetworkLoadBalancer`.",
					Optional:            true,
				},
			},
		},
		"region": schema.StringAttribute{
			MarkdownDescription: "The AWS region where the Cloud NGFW firewall is deployed",
			Required:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"scm_labels": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Strata Cloud Manager (SCM) labels to apply to the firewall.",
			Validators: []validator.List{
				listvalidator.ValueStringsAre(stringvalidator.LengthAtMost(256)),
			},
			Optional: true,
		},
		"source": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.OneOf("SCM-SDK", "Terraform-Provider-SCM"),
			},
			MarkdownDescription: "The source SDK or provider that created the firewall. Automatically set by the SDK or Terraform provider; do not set this field manually. Possible values are `SCM-SDK` and `Terraform-Provider-SCM`.",
			Optional:            true,
		},
		"tags": schema.ListNestedAttribute{
			Validators: []validator.List{
				listvalidator.SizeAtMost(50),
			},
			MarkdownDescription: "AWS resource tags to apply to the firewall.",
			Optional:            true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"key": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.LengthAtMost(128),
							stringvalidator.LengthAtLeast(1),
						},
						MarkdownDescription: "The tag key.",
						Optional:            true,
					},
					"value": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.LengthAtMost(256),
							stringvalidator.LengthAtLeast(1),
						},
						MarkdownDescription: "The tag value.",
						Optional:            true,
					},
				},
			},
		},
		"tfid": schema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"tier": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.OneOf("standard", "premium"),
			},
			MarkdownDescription: "The service tier for the firewall, either standard or premium. Possible values are `standard` and `premium`.",
			Optional:            true,
			Computed:            true,
		},
		"user_id": schema.SingleNestedAttribute{
			MarkdownDescription: "User-ID configuration for user-based policy enforcement.",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"agent_name": schema.StringAttribute{
					MarkdownDescription: "The name of the User-ID agent.",
					Optional:            true,
				},
				"cfturl_link": schema.StringAttribute{
					MarkdownDescription: "CloudFormation template URL link for User-ID setup.",
					Optional:            true,
				},
				"collector_name": schema.StringAttribute{
					MarkdownDescription: "The name of the log collector used by User-ID.",
					Optional:            true,
				},
				"custom_include_exclude_network": schema.ListNestedAttribute{
					MarkdownDescription: "Custom include/exclude network definitions for User-ID.",
					Optional:            true,
					NestedObject: schema.NestedAttributeObject{
						Attributes: map[string]schema.Attribute{
							"discovery_include": schema.BoolAttribute{
								MarkdownDescription: "Whether this network is included in User-ID discovery.",
								Optional:            true,
							},
							"enabled": schema.BoolAttribute{
								MarkdownDescription: "Whether this network entry is active.",
								Optional:            true,
							},
							"name": schema.StringAttribute{
								MarkdownDescription: "Name of the network entry.",
								Optional:            true,
							},
							"network_address": schema.StringAttribute{
								MarkdownDescription: "The network address (CIDR) for this entry.",
								Optional:            true,
							},
						},
					},
				},
				"enabled": schema.BoolAttribute{
					MarkdownDescription: "Whether User-ID is enabled.",
					Optional:            true,
				},
				"endpoint_dns": schema.StringAttribute{
					MarkdownDescription: "DNS name of the User-ID endpoint.",
					Optional:            true,
				},
				"port": schema.Int64Attribute{
					MarkdownDescription: "Port used by the User-ID agent.",
					Optional:            true,
				},
				"secret_key_arn": schema.StringAttribute{
					MarkdownDescription: "ARN of the AWS Secrets Manager secret for User-ID credentials.",
					Optional:            true,
				},
				"user_id_status": schema.StringAttribute{
					Validators: []validator.String{
						stringvalidator.OneOf("EndpointPending", "EndpointFail", "CommitFail", "Enabled", "Disabled", "DisablePending", "UpdatePending", "UpdateFail"),
					},
					MarkdownDescription: "The operational status of the User-ID feature. Possible values are `EndpointPending`, `EndpointFail`, `CommitFail`, `Enabled`, `Disabled`, `DisablePending`, `UpdatePending` and `UpdateFail`.",
					Optional:            true,
				},
			},
		},
	},
}

// NgfirewallsDataSourceSchema defines the schema for Ngfirewalls data source
var NgfirewallsDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Ngfirewall data source",
	Attributes: map[string]dsschema.Attribute{
		"id": dsschema.StringAttribute{
			MarkdownDescription: "The ID of the Cloud NGFW firewall.",
			Required:            true,
		},
		"region": dsschema.StringAttribute{
			MarkdownDescription: "The AWS region where the Cloud NGFW firewall is deployed",
			Required:            true,
		},
		"response": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Inner response object containing firewall details and operational status.",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"firewall": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Core firewall configuration and identity fields.",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"advanced_threat_log": dsschema.BoolAttribute{
							MarkdownDescription: "Whether advanced threat logging is enabled.",
							Computed:            true,
						},
						"allow_list_accounts": dsschema.ListAttribute{
							ElementType:         types.StringType,
							MarkdownDescription: "AWS account IDs permitted to access the firewall endpoint service.",
							Computed:            true,
						},
						"change_protection": dsschema.ListAttribute{
							ElementType:         types.StringType,
							MarkdownDescription: "List of attributes protected from change.",
							Computed:            true,
						},
						"cloudwatch_metrics": dsschema.SingleNestedAttribute{
							MarkdownDescription: "CloudWatch metrics configuration.",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"account_id": dsschema.StringAttribute{
									MarkdownDescription: "The AWS account ID where CloudWatch metrics are published.",
									Computed:            true,
								},
								"cloud_watch_metric_namespace": dsschema.StringAttribute{
									MarkdownDescription: "The CloudWatch metric namespace.",
									Computed:            true,
								},
								"cloud_watch_metrics_fields": dsschema.ListAttribute{
									ElementType:         types.StringType,
									MarkdownDescription: "The CloudWatch metric field names to publish.",
									Computed:            true,
								},
							},
						},
						"customer_zone_id_list": dsschema.ListAttribute{
							ElementType:         types.StringType,
							MarkdownDescription: "List of customer-defined security zone IDs.",
							Computed:            true,
						},
						"deployment_update_token": dsschema.StringAttribute{
							MarkdownDescription: "Token used for deployment update sequencing.",
							Computed:            true,
						},
						"description": dsschema.StringAttribute{
							MarkdownDescription: "A human-readable description of the firewall.",
							Computed:            true,
						},
						"device_rule_stack_commit_status": dsschema.StringAttribute{
							MarkdownDescription: "Commit status of the device rulestack on the firewall.",
							Computed:            true,
						},
						"egress_nat": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Egress NAT configuration for outbound traffic.",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"enabled": dsschema.BoolAttribute{
									MarkdownDescription: "Whether egress NAT is enabled on this firewall.",
									Computed:            true,
								},
								"settings": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Settings for egress NAT IP address management.",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"ip_pool_type": dsschema.StringAttribute{
											MarkdownDescription: "The IP pool type for egress NAT. Possible values are `AWSService` and `BYOIP`.",
											Computed:            true,
										},
										"ipam_pool_id": dsschema.StringAttribute{
											MarkdownDescription: "The IPAM pool ID used for egress NAT IP allocation.",
											Computed:            true,
										},
									},
								},
							},
						},
						"endpoint_service_name": dsschema.StringAttribute{
							MarkdownDescription: "The VPC endpoint service name (only populated in ServiceManaged mode).",
							Computed:            true,
						},
						"endpoints": dsschema.ListNestedAttribute{
							MarkdownDescription: "VPC endpoint configurations created for this firewall.",
							Computed:            true,
							NestedObject: dsschema.NestedAttributeObject{
								Attributes: map[string]dsschema.Attribute{
									"account_id": dsschema.StringAttribute{
										MarkdownDescription: "The AWS account ID that owns this endpoint.",
										Computed:            true,
									},
									"egress_nat_enabled": dsschema.BoolAttribute{
										MarkdownDescription: "Whether egress NAT is enabled for this endpoint.",
										Computed:            true,
									},
									"endpoint_id": dsschema.StringAttribute{
										MarkdownDescription: "The VPC endpoint ID.",
										Computed:            true,
									},
									"mode": dsschema.StringAttribute{
										MarkdownDescription: "The endpoint management mode. Possible values are `ServiceManaged` and `CustomerManaged`.",
										Computed:            true,
									},
									"prefixes": dsschema.SingleNestedAttribute{
										MarkdownDescription: "IP prefix lists associated with a firewall endpoint.",
										Computed:            true,
										Attributes: map[string]dsschema.Attribute{
											"private_prefix": dsschema.SingleNestedAttribute{
												MarkdownDescription: "A list of CIDR prefixes.",
												Computed:            true,
												Attributes: map[string]dsschema.Attribute{
													"cidrs": dsschema.ListAttribute{
														ElementType:         types.StringType,
														MarkdownDescription: "CIDR address blocks.",
														Computed:            true,
													},
												},
											},
											"public_prefix": dsschema.SingleNestedAttribute{
												MarkdownDescription: "A list of CIDR prefixes.",
												Computed:            true,
												Attributes: map[string]dsschema.Attribute{
													"cidrs": dsschema.ListAttribute{
														ElementType:         types.StringType,
														MarkdownDescription: "CIDR address blocks.",
														Computed:            true,
													},
												},
											},
										},
									},
									"rejected_reason": dsschema.StringAttribute{
										MarkdownDescription: "Human-readable reason if the endpoint was rejected.",
										Computed:            true,
									},
									"status": dsschema.StringAttribute{
										MarkdownDescription: "The current provisioning status of the endpoint. Possible values are `Pending`, `Accepted`, `Rejected`, `Created`, `Creating` and `Deleting`.",
										Computed:            true,
									},
									"subnet_id": dsschema.StringAttribute{
										MarkdownDescription: "The AWS subnet ID in which the endpoint resides.",
										Computed:            true,
									},
									"vpc_id": dsschema.StringAttribute{
										MarkdownDescription: "The AWS VPC ID in which the endpoint resides.",
										Computed:            true,
									},
									"zone_id": dsschema.StringAttribute{
										MarkdownDescription: "The availability zone ID for this endpoint.",
										Computed:            true,
									},
								},
							},
						},
						"firewall_id": dsschema.StringAttribute{
							MarkdownDescription: "The unique identifier of the firewall.",
							Computed:            true,
						},
						"gwlb": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Gateway Load Balancer configuration.",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"deregistration_delay": dsschema.Int64Attribute{
									MarkdownDescription: "Deregistration delay in seconds.",
									Computed:            true,
								},
								"rejected_flow_count_alert_threshold": dsschema.Int64Attribute{
									MarkdownDescription: "Alert threshold for the total rejected flow count.",
									Computed:            true,
								},
								"rejected_flow_count_tcp_alert_threshold": dsschema.Int64Attribute{
									MarkdownDescription: "Alert threshold for the TCP rejected flow count.",
									Computed:            true,
								},
								"session_rebalance_enabled": dsschema.BoolAttribute{
									MarkdownDescription: "Whether session rebalancing is enabled.",
									Computed:            true,
								},
								"tcp_idle_timeout": dsschema.Int64Attribute{
									MarkdownDescription: "TCP idle timeout in seconds.",
									Computed:            true,
								},
							},
						},
						"ipv6": dsschema.SingleNestedAttribute{
							MarkdownDescription: "IPv6 configuration for the firewall.",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"enabled": dsschema.BoolAttribute{
									MarkdownDescription: "Whether IPv6 is enabled on this firewall.",
									Computed:            true,
								},
							},
						},
						"link_id": dsschema.StringAttribute{
							MarkdownDescription: "The link ID associated with the firewall.",
							Computed:            true,
						},
						"link_status": dsschema.StringAttribute{
							MarkdownDescription: "The current link status of the firewall.",
							Computed:            true,
						},
						"log_config": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Log destination configuration for the firewall.",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"account_id": dsschema.StringAttribute{
									MarkdownDescription: "The AWS account ID where logs are sent.",
									Computed:            true,
								},
								"log_destination": dsschema.StringAttribute{
									MarkdownDescription: "The ARN or name of the log destination resource.",
									Computed:            true,
								},
								"log_destination_region": dsschema.StringAttribute{
									MarkdownDescription: "The AWS region of the log destination.",
									Computed:            true,
								},
								"log_destination_type": dsschema.StringAttribute{
									MarkdownDescription: "The type of log destination. Possible values are `S3`, `CloudWatchLogs` and `KinesisDataFirehose`.",
									Computed:            true,
								},
								"log_type": dsschema.ListAttribute{
									ElementType:         types.StringType,
									MarkdownDescription: "The types of logs to send to this destination. Possible values are `TRAFFIC`, `DECRYPTION`, `THREAT` and `USER_ID`.",
									Computed:            true,
								},
								"role_type": dsschema.StringAttribute{
									MarkdownDescription: "The IAM role type used to publish logs. Possible values are `ResourceBased` and `IamBased`.",
									Computed:            true,
								},
							},
						},
						"notifications": dsschema.ListNestedAttribute{
							MarkdownDescription: "Notifications associated with this firewall.",
							Computed:            true,
							NestedObject: dsschema.NestedAttributeObject{
								Attributes: map[string]dsschema.Attribute{
									"expiry_time": dsschema.Int64Attribute{
										MarkdownDescription: "Unix timestamp when this notification expires.",
										Computed:            true,
									},
									"id": dsschema.StringAttribute{
										MarkdownDescription: "Unique identifier of the notification.",
										Computed:            true,
									},
									"link": dsschema.SingleNestedAttribute{
										MarkdownDescription: "A hyperlink associated with a notification.",
										Computed:            true,
										Attributes: map[string]dsschema.Attribute{
											"target": dsschema.StringAttribute{
												MarkdownDescription: "The URL target of the link.",
												Computed:            true,
											},
											"text": dsschema.StringAttribute{
												MarkdownDescription: "The display text of the link.",
												Computed:            true,
											},
										},
									},
									"msg": dsschema.StringAttribute{
										MarkdownDescription: "The notification message text.",
										Computed:            true,
									},
								},
							},
						},
						"private_access": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Private access configuration for the firewall.",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"resource_id": dsschema.StringAttribute{
									MarkdownDescription: "The resource ID of the private access endpoint.",
									Computed:            true,
								},
								"type": dsschema.StringAttribute{
									MarkdownDescription: "The type of private access resource. Possible values are `NetworkLoadBalancer`.",
									Computed:            true,
								},
							},
						},
						"region": dsschema.StringAttribute{
							MarkdownDescription: "The AWS region where the firewall is deployed.",
							Computed:            true,
						},
						"software_version": dsschema.StringAttribute{
							MarkdownDescription: "The PAN-OS software version currently running on the firewall.",
							Computed:            true,
						},
						"tags": dsschema.ListNestedAttribute{
							MarkdownDescription: "AWS resource tags applied to the firewall.",
							Computed:            true,
							NestedObject: dsschema.NestedAttributeObject{
								Attributes: map[string]dsschema.Attribute{
									"key": dsschema.StringAttribute{
										MarkdownDescription: "The tag key.",
										Computed:            true,
									},
									"value": dsschema.StringAttribute{
										MarkdownDescription: "The tag value.",
										Computed:            true,
									},
								},
							},
						},
						"tier": dsschema.StringAttribute{
							MarkdownDescription: "The service tier of the firewall (base, standard, or premium).",
							Computed:            true,
						},
						"update_token": dsschema.StringAttribute{
							MarkdownDescription: "Optimistic locking token; must be supplied on update operations.",
							Computed:            true,
						},
						"user_id": dsschema.SingleNestedAttribute{
							MarkdownDescription: "User-ID configuration for user-based policy enforcement.",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"agent_name": dsschema.StringAttribute{
									MarkdownDescription: "The name of the User-ID agent.",
									Computed:            true,
								},
								"cfturl_link": dsschema.StringAttribute{
									MarkdownDescription: "CloudFormation template URL link for User-ID setup.",
									Computed:            true,
								},
								"collector_name": dsschema.StringAttribute{
									MarkdownDescription: "The name of the log collector used by User-ID.",
									Computed:            true,
								},
								"custom_include_exclude_network": dsschema.ListNestedAttribute{
									MarkdownDescription: "Custom include/exclude network definitions for User-ID.",
									Computed:            true,
									NestedObject: dsschema.NestedAttributeObject{
										Attributes: map[string]dsschema.Attribute{
											"discovery_include": dsschema.BoolAttribute{
												MarkdownDescription: "Whether this network is included in User-ID discovery.",
												Computed:            true,
											},
											"enabled": dsschema.BoolAttribute{
												MarkdownDescription: "Whether this network entry is active.",
												Computed:            true,
											},
											"name": dsschema.StringAttribute{
												MarkdownDescription: "Name of the network entry.",
												Computed:            true,
											},
											"network_address": dsschema.StringAttribute{
												MarkdownDescription: "The network address (CIDR) for this entry.",
												Computed:            true,
											},
										},
									},
								},
								"enabled": dsschema.BoolAttribute{
									MarkdownDescription: "Whether User-ID is enabled.",
									Computed:            true,
								},
								"endpoint_dns": dsschema.StringAttribute{
									MarkdownDescription: "DNS name of the User-ID endpoint.",
									Computed:            true,
								},
								"port": dsschema.Int64Attribute{
									MarkdownDescription: "Port used by the User-ID agent.",
									Computed:            true,
								},
								"secret_key_arn": dsschema.StringAttribute{
									MarkdownDescription: "ARN of the AWS Secrets Manager secret for User-ID credentials.",
									Computed:            true,
								},
								"user_id_status": dsschema.StringAttribute{
									MarkdownDescription: "The operational status of the User-ID feature. Possible values are `EndpointPending`, `EndpointFail`, `CommitFail`, `Enabled`, `Disabled`, `DisablePending`, `UpdatePending` and `UpdateFail`.",
									Computed:            true,
								},
							},
						},
					},
				},
				"status": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Operational status of the firewall.",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"device_rule_stack_commit_info": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Commit information for a rulestack.",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"commit_messages": dsschema.ListAttribute{
									ElementType:         types.StringType,
									MarkdownDescription: "Messages produced during the commit operation.",
									Computed:            true,
								},
								"commit_ts": dsschema.StringAttribute{
									MarkdownDescription: "Timestamp of the last commit operation.",
									Computed:            true,
								},
							},
						},
						"device_rule_stack_commit_state": dsschema.StringAttribute{
							MarkdownDescription: "The commit state of the device rulestack. Possible values are `Submitted` and `None`.",
							Computed:            true,
						},
						"device_rule_stack_commit_status": dsschema.StringAttribute{
							MarkdownDescription: "The commit status of the device rulestack. Possible values are `Running`, `Uncommitted`, `Precommit`, `PrecommitDone`, `PrevalidateDone`, `Committing`, `Failed` and `Success`.",
							Computed:            true,
						},
						"failure_reason": dsschema.StringAttribute{
							MarkdownDescription: "Human-readable reason for a firewall failure, if any.",
							Computed:            true,
						},
						"firewall_status": dsschema.StringAttribute{
							MarkdownDescription: "The operational lifecycle status of the firewall. Possible values are `CREATING`, `UPDATING`, `DELETING`, `CREATE_COMPLETE`, `UPDATE_COMPLETE`, `CREATE_FAIL`, `UPDATE_FAIL`, `DELETE_FAIL` and `DELETE_COMPLETE`.",
							Computed:            true,
						},
						"global_rule_stack_commit_info": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Commit information for a rulestack.",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"commit_messages": dsschema.ListAttribute{
									ElementType:         types.StringType,
									MarkdownDescription: "Messages produced during the commit operation.",
									Computed:            true,
								},
								"commit_ts": dsschema.StringAttribute{
									MarkdownDescription: "Timestamp of the last commit operation.",
									Computed:            true,
								},
							},
						},
						"global_rule_stack_status": dsschema.StringAttribute{
							MarkdownDescription: "The commit status of the global rulestack. Possible values are `Running`, `Uncommitted`, `Precommit`, `PrecommitDone`, `PrevalidateDone`, `Committing`, `Failed` and `Success`.",
							Computed:            true,
						},
						"public_i_ps": dsschema.ListNestedAttribute{
							MarkdownDescription: "Public IP addresses assigned to the firewall.",
							Computed:            true,
							NestedObject: dsschema.NestedAttributeObject{
								Attributes: map[string]dsschema.Attribute{
									"ip_address": dsschema.StringAttribute{
										MarkdownDescription: "The public IP address.",
										Computed:            true,
									},
									"ip_source": dsschema.StringAttribute{
										MarkdownDescription: "The source of the IP address allocation.",
										Computed:            true,
									},
									"ip_status": dsschema.StringAttribute{
										MarkdownDescription: "The current status of this public IP.",
										Computed:            true,
									},
								},
							},
						},
						"rule_stack_commit_info": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Commit information for a rulestack.",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"commit_messages": dsschema.ListAttribute{
									ElementType:         types.StringType,
									MarkdownDescription: "Messages produced during the commit operation.",
									Computed:            true,
								},
								"commit_ts": dsschema.StringAttribute{
									MarkdownDescription: "Timestamp of the last commit operation.",
									Computed:            true,
								},
							},
						},
						"rule_stack_status": dsschema.StringAttribute{
							MarkdownDescription: "The commit status of the local rulestack. Possible values are `Running`, `Uncommitted`, `Precommit`, `PrecommitDone`, `PrevalidateDone`, `Committing`, `Failed` and `Success`.",
							Computed:            true,
						},
						"scm_assoc_status": dsschema.BoolAttribute{
							MarkdownDescription: "Whether the firewall is associated with Strata Cloud Manager.",
							Computed:            true,
						},
					},
				},
			},
		},
		"response_status": dsschema.SingleNestedAttribute{
			MarkdownDescription: "API-level response metadata at the top level of the envelope.",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"error_code": dsschema.Int64Attribute{
					MarkdownDescription: "Non-zero error code if the API request failed.",
					Computed:            true,
				},
				"reason": dsschema.StringAttribute{
					MarkdownDescription: "Human-readable reason for the error, if any.",
					Computed:            true,
				},
			},
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
	},
}

// NgfirewallResponseBodyDataSourceSchema defines the schema for NgfirewallResponseBody.
var NgfirewallResponseBodyDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "NgfirewallResponseBody data source",
	Attributes: map[string]dsschema.Attribute{
		"firewall": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Core firewall configuration and identity fields.",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"advanced_threat_log": dsschema.BoolAttribute{
					MarkdownDescription: "Whether advanced threat logging is enabled.",
					Computed:            true,
				},
				"allow_list_accounts": dsschema.ListAttribute{
					ElementType:         types.StringType,
					MarkdownDescription: "AWS account IDs permitted to access the firewall endpoint service.",
					Computed:            true,
				},
				"change_protection": dsschema.ListAttribute{
					ElementType:         types.StringType,
					MarkdownDescription: "List of attributes protected from change.",
					Computed:            true,
				},
				"cloudwatch_metrics": dsschema.SingleNestedAttribute{
					MarkdownDescription: "CloudWatch metrics configuration.",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"account_id": dsschema.StringAttribute{
							MarkdownDescription: "The AWS account ID where CloudWatch metrics are published.",
							Computed:            true,
						},
						"cloud_watch_metric_namespace": dsschema.StringAttribute{
							MarkdownDescription: "The CloudWatch metric namespace.",
							Computed:            true,
						},
						"cloud_watch_metrics_fields": dsschema.ListAttribute{
							ElementType:         types.StringType,
							MarkdownDescription: "The CloudWatch metric field names to publish.",
							Computed:            true,
						},
					},
				},
				"customer_zone_id_list": dsschema.ListAttribute{
					ElementType:         types.StringType,
					MarkdownDescription: "List of customer-defined security zone IDs.",
					Computed:            true,
				},
				"deployment_update_token": dsschema.StringAttribute{
					MarkdownDescription: "Token used for deployment update sequencing.",
					Computed:            true,
				},
				"description": dsschema.StringAttribute{
					MarkdownDescription: "A human-readable description of the firewall.",
					Computed:            true,
				},
				"device_rule_stack_commit_status": dsschema.StringAttribute{
					MarkdownDescription: "Commit status of the device rulestack on the firewall.",
					Computed:            true,
				},
				"egress_nat": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Egress NAT configuration for outbound traffic.",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"enabled": dsschema.BoolAttribute{
							MarkdownDescription: "Whether egress NAT is enabled on this firewall.",
							Computed:            true,
						},
						"settings": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Settings for egress NAT IP address management.",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"ip_pool_type": dsschema.StringAttribute{
									MarkdownDescription: "The IP pool type for egress NAT. Possible values are `AWSService` and `BYOIP`.",
									Computed:            true,
								},
								"ipam_pool_id": dsschema.StringAttribute{
									MarkdownDescription: "The IPAM pool ID used for egress NAT IP allocation.",
									Computed:            true,
								},
							},
						},
					},
				},
				"endpoint_service_name": dsschema.StringAttribute{
					MarkdownDescription: "The VPC endpoint service name (only populated in ServiceManaged mode).",
					Computed:            true,
				},
				"endpoints": dsschema.ListNestedAttribute{
					MarkdownDescription: "VPC endpoint configurations created for this firewall.",
					Computed:            true,
					NestedObject: dsschema.NestedAttributeObject{
						Attributes: map[string]dsschema.Attribute{
							"account_id": dsschema.StringAttribute{
								MarkdownDescription: "The AWS account ID that owns this endpoint.",
								Computed:            true,
							},
							"egress_nat_enabled": dsschema.BoolAttribute{
								MarkdownDescription: "Whether egress NAT is enabled for this endpoint.",
								Computed:            true,
							},
							"endpoint_id": dsschema.StringAttribute{
								MarkdownDescription: "The VPC endpoint ID.",
								Computed:            true,
							},
							"mode": dsschema.StringAttribute{
								MarkdownDescription: "The endpoint management mode. Possible values are `ServiceManaged` and `CustomerManaged`.",
								Computed:            true,
							},
							"prefixes": dsschema.SingleNestedAttribute{
								MarkdownDescription: "IP prefix lists associated with a firewall endpoint.",
								Computed:            true,
								Attributes: map[string]dsschema.Attribute{
									"private_prefix": dsschema.SingleNestedAttribute{
										MarkdownDescription: "A list of CIDR prefixes.",
										Computed:            true,
										Attributes: map[string]dsschema.Attribute{
											"cidrs": dsschema.ListAttribute{
												ElementType:         types.StringType,
												MarkdownDescription: "CIDR address blocks.",
												Computed:            true,
											},
										},
									},
									"public_prefix": dsschema.SingleNestedAttribute{
										MarkdownDescription: "A list of CIDR prefixes.",
										Computed:            true,
										Attributes: map[string]dsschema.Attribute{
											"cidrs": dsschema.ListAttribute{
												ElementType:         types.StringType,
												MarkdownDescription: "CIDR address blocks.",
												Computed:            true,
											},
										},
									},
								},
							},
							"rejected_reason": dsschema.StringAttribute{
								MarkdownDescription: "Human-readable reason if the endpoint was rejected.",
								Computed:            true,
							},
							"status": dsschema.StringAttribute{
								MarkdownDescription: "The current provisioning status of the endpoint. Possible values are `Pending`, `Accepted`, `Rejected`, `Created`, `Creating` and `Deleting`.",
								Computed:            true,
							},
							"subnet_id": dsschema.StringAttribute{
								MarkdownDescription: "The AWS subnet ID in which the endpoint resides.",
								Computed:            true,
							},
							"vpc_id": dsschema.StringAttribute{
								MarkdownDescription: "The AWS VPC ID in which the endpoint resides.",
								Computed:            true,
							},
							"zone_id": dsschema.StringAttribute{
								MarkdownDescription: "The availability zone ID for this endpoint.",
								Computed:            true,
							},
						},
					},
				},
				"firewall_id": dsschema.StringAttribute{
					MarkdownDescription: "The unique identifier of the firewall.",
					Computed:            true,
				},
				"gwlb": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Gateway Load Balancer configuration.",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"deregistration_delay": dsschema.Int64Attribute{
							MarkdownDescription: "Deregistration delay in seconds.",
							Computed:            true,
						},
						"rejected_flow_count_alert_threshold": dsschema.Int64Attribute{
							MarkdownDescription: "Alert threshold for the total rejected flow count.",
							Computed:            true,
						},
						"rejected_flow_count_tcp_alert_threshold": dsschema.Int64Attribute{
							MarkdownDescription: "Alert threshold for the TCP rejected flow count.",
							Computed:            true,
						},
						"session_rebalance_enabled": dsschema.BoolAttribute{
							MarkdownDescription: "Whether session rebalancing is enabled.",
							Computed:            true,
						},
						"tcp_idle_timeout": dsschema.Int64Attribute{
							MarkdownDescription: "TCP idle timeout in seconds.",
							Computed:            true,
						},
					},
				},
				"ipv6": dsschema.SingleNestedAttribute{
					MarkdownDescription: "IPv6 configuration for the firewall.",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"enabled": dsschema.BoolAttribute{
							MarkdownDescription: "Whether IPv6 is enabled on this firewall.",
							Computed:            true,
						},
					},
				},
				"link_id": dsschema.StringAttribute{
					MarkdownDescription: "The link ID associated with the firewall.",
					Computed:            true,
				},
				"link_status": dsschema.StringAttribute{
					MarkdownDescription: "The current link status of the firewall.",
					Computed:            true,
				},
				"log_config": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Log destination configuration for the firewall.",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"account_id": dsschema.StringAttribute{
							MarkdownDescription: "The AWS account ID where logs are sent.",
							Computed:            true,
						},
						"log_destination": dsschema.StringAttribute{
							MarkdownDescription: "The ARN or name of the log destination resource.",
							Computed:            true,
						},
						"log_destination_region": dsschema.StringAttribute{
							MarkdownDescription: "The AWS region of the log destination.",
							Computed:            true,
						},
						"log_destination_type": dsschema.StringAttribute{
							MarkdownDescription: "The type of log destination. Possible values are `S3`, `CloudWatchLogs` and `KinesisDataFirehose`.",
							Computed:            true,
						},
						"log_type": dsschema.ListAttribute{
							ElementType:         types.StringType,
							MarkdownDescription: "The types of logs to send to this destination. Possible values are `TRAFFIC`, `DECRYPTION`, `THREAT` and `USER_ID`.",
							Computed:            true,
						},
						"role_type": dsschema.StringAttribute{
							MarkdownDescription: "The IAM role type used to publish logs. Possible values are `ResourceBased` and `IamBased`.",
							Computed:            true,
						},
					},
				},
				"notifications": dsschema.ListNestedAttribute{
					MarkdownDescription: "Notifications associated with this firewall.",
					Computed:            true,
					NestedObject: dsschema.NestedAttributeObject{
						Attributes: map[string]dsschema.Attribute{
							"expiry_time": dsschema.Int64Attribute{
								MarkdownDescription: "Unix timestamp when this notification expires.",
								Computed:            true,
							},
							"id": dsschema.StringAttribute{
								MarkdownDescription: "Unique identifier of the notification.",
								Computed:            true,
							},
							"link": dsschema.SingleNestedAttribute{
								MarkdownDescription: "A hyperlink associated with a notification.",
								Computed:            true,
								Attributes: map[string]dsschema.Attribute{
									"target": dsschema.StringAttribute{
										MarkdownDescription: "The URL target of the link.",
										Computed:            true,
									},
									"text": dsschema.StringAttribute{
										MarkdownDescription: "The display text of the link.",
										Computed:            true,
									},
								},
							},
							"msg": dsschema.StringAttribute{
								MarkdownDescription: "The notification message text.",
								Computed:            true,
							},
						},
					},
				},
				"private_access": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Private access configuration for the firewall.",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"resource_id": dsschema.StringAttribute{
							MarkdownDescription: "The resource ID of the private access endpoint.",
							Computed:            true,
						},
						"type": dsschema.StringAttribute{
							MarkdownDescription: "The type of private access resource. Possible values are `NetworkLoadBalancer`.",
							Computed:            true,
						},
					},
				},
				"region": dsschema.StringAttribute{
					MarkdownDescription: "The AWS region where the firewall is deployed.",
					Computed:            true,
				},
				"software_version": dsschema.StringAttribute{
					MarkdownDescription: "The PAN-OS software version currently running on the firewall.",
					Computed:            true,
				},
				"tags": dsschema.ListNestedAttribute{
					MarkdownDescription: "AWS resource tags applied to the firewall.",
					Computed:            true,
					NestedObject: dsschema.NestedAttributeObject{
						Attributes: map[string]dsschema.Attribute{
							"key": dsschema.StringAttribute{
								MarkdownDescription: "The tag key.",
								Computed:            true,
							},
							"value": dsschema.StringAttribute{
								MarkdownDescription: "The tag value.",
								Computed:            true,
							},
						},
					},
				},
				"tier": dsschema.StringAttribute{
					MarkdownDescription: "The service tier of the firewall (base, standard, or premium).",
					Computed:            true,
				},
				"update_token": dsschema.StringAttribute{
					MarkdownDescription: "Optimistic locking token; must be supplied on update operations.",
					Computed:            true,
				},
				"user_id": dsschema.SingleNestedAttribute{
					MarkdownDescription: "User-ID configuration for user-based policy enforcement.",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"agent_name": dsschema.StringAttribute{
							MarkdownDescription: "The name of the User-ID agent.",
							Computed:            true,
						},
						"cfturl_link": dsschema.StringAttribute{
							MarkdownDescription: "CloudFormation template URL link for User-ID setup.",
							Computed:            true,
						},
						"collector_name": dsschema.StringAttribute{
							MarkdownDescription: "The name of the log collector used by User-ID.",
							Computed:            true,
						},
						"custom_include_exclude_network": dsschema.ListNestedAttribute{
							MarkdownDescription: "Custom include/exclude network definitions for User-ID.",
							Computed:            true,
							NestedObject: dsschema.NestedAttributeObject{
								Attributes: map[string]dsschema.Attribute{
									"discovery_include": dsschema.BoolAttribute{
										MarkdownDescription: "Whether this network is included in User-ID discovery.",
										Computed:            true,
									},
									"enabled": dsschema.BoolAttribute{
										MarkdownDescription: "Whether this network entry is active.",
										Computed:            true,
									},
									"name": dsschema.StringAttribute{
										MarkdownDescription: "Name of the network entry.",
										Computed:            true,
									},
									"network_address": dsschema.StringAttribute{
										MarkdownDescription: "The network address (CIDR) for this entry.",
										Computed:            true,
									},
								},
							},
						},
						"enabled": dsschema.BoolAttribute{
							MarkdownDescription: "Whether User-ID is enabled.",
							Computed:            true,
						},
						"endpoint_dns": dsschema.StringAttribute{
							MarkdownDescription: "DNS name of the User-ID endpoint.",
							Computed:            true,
						},
						"port": dsschema.Int64Attribute{
							MarkdownDescription: "Port used by the User-ID agent.",
							Computed:            true,
						},
						"secret_key_arn": dsschema.StringAttribute{
							MarkdownDescription: "ARN of the AWS Secrets Manager secret for User-ID credentials.",
							Computed:            true,
						},
						"user_id_status": dsschema.StringAttribute{
							MarkdownDescription: "The operational status of the User-ID feature. Possible values are `EndpointPending`, `EndpointFail`, `CommitFail`, `Enabled`, `Disabled`, `DisablePending`, `UpdatePending` and `UpdateFail`.",
							Computed:            true,
						},
					},
				},
			},
		},
		"status": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Operational status of the firewall.",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"device_rule_stack_commit_info": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Commit information for a rulestack.",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"commit_messages": dsschema.ListAttribute{
							ElementType:         types.StringType,
							MarkdownDescription: "Messages produced during the commit operation.",
							Computed:            true,
						},
						"commit_ts": dsschema.StringAttribute{
							MarkdownDescription: "Timestamp of the last commit operation.",
							Computed:            true,
						},
					},
				},
				"device_rule_stack_commit_state": dsschema.StringAttribute{
					MarkdownDescription: "The commit state of the device rulestack. Possible values are `Submitted` and `None`.",
					Computed:            true,
				},
				"device_rule_stack_commit_status": dsschema.StringAttribute{
					MarkdownDescription: "The commit status of the device rulestack. Possible values are `Running`, `Uncommitted`, `Precommit`, `PrecommitDone`, `PrevalidateDone`, `Committing`, `Failed` and `Success`.",
					Computed:            true,
				},
				"failure_reason": dsschema.StringAttribute{
					MarkdownDescription: "Human-readable reason for a firewall failure, if any.",
					Computed:            true,
				},
				"firewall_status": dsschema.StringAttribute{
					MarkdownDescription: "The operational lifecycle status of the firewall. Possible values are `CREATING`, `UPDATING`, `DELETING`, `CREATE_COMPLETE`, `UPDATE_COMPLETE`, `CREATE_FAIL`, `UPDATE_FAIL`, `DELETE_FAIL` and `DELETE_COMPLETE`.",
					Computed:            true,
				},
				"global_rule_stack_commit_info": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Commit information for a rulestack.",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"commit_messages": dsschema.ListAttribute{
							ElementType:         types.StringType,
							MarkdownDescription: "Messages produced during the commit operation.",
							Computed:            true,
						},
						"commit_ts": dsschema.StringAttribute{
							MarkdownDescription: "Timestamp of the last commit operation.",
							Computed:            true,
						},
					},
				},
				"global_rule_stack_status": dsschema.StringAttribute{
					MarkdownDescription: "The commit status of the global rulestack. Possible values are `Running`, `Uncommitted`, `Precommit`, `PrecommitDone`, `PrevalidateDone`, `Committing`, `Failed` and `Success`.",
					Computed:            true,
				},
				"public_i_ps": dsschema.ListNestedAttribute{
					MarkdownDescription: "Public IP addresses assigned to the firewall.",
					Computed:            true,
					NestedObject: dsschema.NestedAttributeObject{
						Attributes: map[string]dsschema.Attribute{
							"ip_address": dsschema.StringAttribute{
								MarkdownDescription: "The public IP address.",
								Computed:            true,
							},
							"ip_source": dsschema.StringAttribute{
								MarkdownDescription: "The source of the IP address allocation.",
								Computed:            true,
							},
							"ip_status": dsschema.StringAttribute{
								MarkdownDescription: "The current status of this public IP.",
								Computed:            true,
							},
						},
					},
				},
				"rule_stack_commit_info": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Commit information for a rulestack.",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"commit_messages": dsschema.ListAttribute{
							ElementType:         types.StringType,
							MarkdownDescription: "Messages produced during the commit operation.",
							Computed:            true,
						},
						"commit_ts": dsschema.StringAttribute{
							MarkdownDescription: "Timestamp of the last commit operation.",
							Computed:            true,
						},
					},
				},
				"rule_stack_status": dsschema.StringAttribute{
					MarkdownDescription: "The commit status of the local rulestack. Possible values are `Running`, `Uncommitted`, `Precommit`, `PrecommitDone`, `PrevalidateDone`, `Committing`, `Failed` and `Success`.",
					Computed:            true,
				},
				"scm_assoc_status": dsschema.BoolAttribute{
					MarkdownDescription: "Whether the firewall is associated with Strata Cloud Manager.",
					Computed:            true,
				},
			},
		},
	},
}

// NgfirewallFirewallDataSourceSchema defines the schema for NgfirewallFirewall.
var NgfirewallFirewallDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "NgfirewallFirewall data source",
	Attributes: map[string]dsschema.Attribute{
		"advanced_threat_log": dsschema.BoolAttribute{
			MarkdownDescription: "Whether advanced threat logging is enabled.",
			Computed:            true,
		},
		"allow_list_accounts": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "AWS account IDs permitted to access the firewall endpoint service.",
			Computed:            true,
		},
		"change_protection": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "List of attributes protected from change.",
			Computed:            true,
		},
		"cloudwatch_metrics": dsschema.SingleNestedAttribute{
			MarkdownDescription: "CloudWatch metrics configuration.",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"account_id": dsschema.StringAttribute{
					MarkdownDescription: "The AWS account ID where CloudWatch metrics are published.",
					Computed:            true,
				},
				"cloud_watch_metric_namespace": dsschema.StringAttribute{
					MarkdownDescription: "The CloudWatch metric namespace.",
					Computed:            true,
				},
				"cloud_watch_metrics_fields": dsschema.ListAttribute{
					ElementType:         types.StringType,
					MarkdownDescription: "The CloudWatch metric field names to publish.",
					Computed:            true,
				},
			},
		},
		"customer_zone_id_list": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "List of customer-defined security zone IDs.",
			Computed:            true,
		},
		"deployment_update_token": dsschema.StringAttribute{
			MarkdownDescription: "Token used for deployment update sequencing.",
			Computed:            true,
		},
		"description": dsschema.StringAttribute{
			MarkdownDescription: "A human-readable description of the firewall.",
			Computed:            true,
		},
		"device_rule_stack_commit_status": dsschema.StringAttribute{
			MarkdownDescription: "Commit status of the device rulestack on the firewall.",
			Computed:            true,
		},
		"egress_nat": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Egress NAT configuration for outbound traffic.",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"enabled": dsschema.BoolAttribute{
					MarkdownDescription: "Whether egress NAT is enabled on this firewall.",
					Computed:            true,
				},
				"settings": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Settings for egress NAT IP address management.",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"ip_pool_type": dsschema.StringAttribute{
							MarkdownDescription: "The IP pool type for egress NAT. Possible values are `AWSService` and `BYOIP`.",
							Computed:            true,
						},
						"ipam_pool_id": dsschema.StringAttribute{
							MarkdownDescription: "The IPAM pool ID used for egress NAT IP allocation.",
							Computed:            true,
						},
					},
				},
			},
		},
		"endpoint_service_name": dsschema.StringAttribute{
			MarkdownDescription: "The VPC endpoint service name (only populated in ServiceManaged mode).",
			Computed:            true,
		},
		"endpoints": dsschema.ListNestedAttribute{
			MarkdownDescription: "VPC endpoint configurations created for this firewall.",
			Computed:            true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: map[string]dsschema.Attribute{
					"account_id": dsschema.StringAttribute{
						MarkdownDescription: "The AWS account ID that owns this endpoint.",
						Computed:            true,
					},
					"egress_nat_enabled": dsschema.BoolAttribute{
						MarkdownDescription: "Whether egress NAT is enabled for this endpoint.",
						Computed:            true,
					},
					"endpoint_id": dsschema.StringAttribute{
						MarkdownDescription: "The VPC endpoint ID.",
						Computed:            true,
					},
					"mode": dsschema.StringAttribute{
						MarkdownDescription: "The endpoint management mode. Possible values are `ServiceManaged` and `CustomerManaged`.",
						Computed:            true,
					},
					"prefixes": dsschema.SingleNestedAttribute{
						MarkdownDescription: "IP prefix lists associated with a firewall endpoint.",
						Computed:            true,
						Attributes: map[string]dsschema.Attribute{
							"private_prefix": dsschema.SingleNestedAttribute{
								MarkdownDescription: "A list of CIDR prefixes.",
								Computed:            true,
								Attributes: map[string]dsschema.Attribute{
									"cidrs": dsschema.ListAttribute{
										ElementType:         types.StringType,
										MarkdownDescription: "CIDR address blocks.",
										Computed:            true,
									},
								},
							},
							"public_prefix": dsschema.SingleNestedAttribute{
								MarkdownDescription: "A list of CIDR prefixes.",
								Computed:            true,
								Attributes: map[string]dsschema.Attribute{
									"cidrs": dsschema.ListAttribute{
										ElementType:         types.StringType,
										MarkdownDescription: "CIDR address blocks.",
										Computed:            true,
									},
								},
							},
						},
					},
					"rejected_reason": dsschema.StringAttribute{
						MarkdownDescription: "Human-readable reason if the endpoint was rejected.",
						Computed:            true,
					},
					"status": dsschema.StringAttribute{
						MarkdownDescription: "The current provisioning status of the endpoint. Possible values are `Pending`, `Accepted`, `Rejected`, `Created`, `Creating` and `Deleting`.",
						Computed:            true,
					},
					"subnet_id": dsschema.StringAttribute{
						MarkdownDescription: "The AWS subnet ID in which the endpoint resides.",
						Computed:            true,
					},
					"vpc_id": dsschema.StringAttribute{
						MarkdownDescription: "The AWS VPC ID in which the endpoint resides.",
						Computed:            true,
					},
					"zone_id": dsschema.StringAttribute{
						MarkdownDescription: "The availability zone ID for this endpoint.",
						Computed:            true,
					},
				},
			},
		},
		"firewall_id": dsschema.StringAttribute{
			MarkdownDescription: "The unique identifier of the firewall.",
			Computed:            true,
		},
		"gwlb": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Gateway Load Balancer configuration.",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"deregistration_delay": dsschema.Int64Attribute{
					MarkdownDescription: "Deregistration delay in seconds.",
					Computed:            true,
				},
				"rejected_flow_count_alert_threshold": dsschema.Int64Attribute{
					MarkdownDescription: "Alert threshold for the total rejected flow count.",
					Computed:            true,
				},
				"rejected_flow_count_tcp_alert_threshold": dsschema.Int64Attribute{
					MarkdownDescription: "Alert threshold for the TCP rejected flow count.",
					Computed:            true,
				},
				"session_rebalance_enabled": dsschema.BoolAttribute{
					MarkdownDescription: "Whether session rebalancing is enabled.",
					Computed:            true,
				},
				"tcp_idle_timeout": dsschema.Int64Attribute{
					MarkdownDescription: "TCP idle timeout in seconds.",
					Computed:            true,
				},
			},
		},
		"ipv6": dsschema.SingleNestedAttribute{
			MarkdownDescription: "IPv6 configuration for the firewall.",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"enabled": dsschema.BoolAttribute{
					MarkdownDescription: "Whether IPv6 is enabled on this firewall.",
					Computed:            true,
				},
			},
		},
		"link_id": dsschema.StringAttribute{
			MarkdownDescription: "The link ID associated with the firewall.",
			Computed:            true,
		},
		"link_status": dsschema.StringAttribute{
			MarkdownDescription: "The current link status of the firewall.",
			Computed:            true,
		},
		"log_config": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Log destination configuration for the firewall.",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"account_id": dsschema.StringAttribute{
					MarkdownDescription: "The AWS account ID where logs are sent.",
					Computed:            true,
				},
				"log_destination": dsschema.StringAttribute{
					MarkdownDescription: "The ARN or name of the log destination resource.",
					Computed:            true,
				},
				"log_destination_region": dsschema.StringAttribute{
					MarkdownDescription: "The AWS region of the log destination.",
					Computed:            true,
				},
				"log_destination_type": dsschema.StringAttribute{
					MarkdownDescription: "The type of log destination. Possible values are `S3`, `CloudWatchLogs` and `KinesisDataFirehose`.",
					Computed:            true,
				},
				"log_type": dsschema.ListAttribute{
					ElementType:         types.StringType,
					MarkdownDescription: "The types of logs to send to this destination. Possible values are `TRAFFIC`, `DECRYPTION`, `THREAT` and `USER_ID`.",
					Computed:            true,
				},
				"role_type": dsschema.StringAttribute{
					MarkdownDescription: "The IAM role type used to publish logs. Possible values are `ResourceBased` and `IamBased`.",
					Computed:            true,
				},
			},
		},
		"notifications": dsschema.ListNestedAttribute{
			MarkdownDescription: "Notifications associated with this firewall.",
			Computed:            true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: map[string]dsschema.Attribute{
					"expiry_time": dsschema.Int64Attribute{
						MarkdownDescription: "Unix timestamp when this notification expires.",
						Computed:            true,
					},
					"id": dsschema.StringAttribute{
						MarkdownDescription: "Unique identifier of the notification.",
						Computed:            true,
					},
					"link": dsschema.SingleNestedAttribute{
						MarkdownDescription: "A hyperlink associated with a notification.",
						Computed:            true,
						Attributes: map[string]dsschema.Attribute{
							"target": dsschema.StringAttribute{
								MarkdownDescription: "The URL target of the link.",
								Computed:            true,
							},
							"text": dsschema.StringAttribute{
								MarkdownDescription: "The display text of the link.",
								Computed:            true,
							},
						},
					},
					"msg": dsschema.StringAttribute{
						MarkdownDescription: "The notification message text.",
						Computed:            true,
					},
				},
			},
		},
		"private_access": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Private access configuration for the firewall.",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"resource_id": dsschema.StringAttribute{
					MarkdownDescription: "The resource ID of the private access endpoint.",
					Computed:            true,
				},
				"type": dsschema.StringAttribute{
					MarkdownDescription: "The type of private access resource. Possible values are `NetworkLoadBalancer`.",
					Computed:            true,
				},
			},
		},
		"region": dsschema.StringAttribute{
			MarkdownDescription: "The AWS region where the firewall is deployed.",
			Computed:            true,
		},
		"software_version": dsschema.StringAttribute{
			MarkdownDescription: "The PAN-OS software version currently running on the firewall.",
			Computed:            true,
		},
		"tags": dsschema.ListNestedAttribute{
			MarkdownDescription: "AWS resource tags applied to the firewall.",
			Computed:            true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: map[string]dsschema.Attribute{
					"key": dsschema.StringAttribute{
						MarkdownDescription: "The tag key.",
						Computed:            true,
					},
					"value": dsschema.StringAttribute{
						MarkdownDescription: "The tag value.",
						Computed:            true,
					},
				},
			},
		},
		"tier": dsschema.StringAttribute{
			MarkdownDescription: "The service tier of the firewall (base, standard, or premium).",
			Computed:            true,
		},
		"update_token": dsschema.StringAttribute{
			MarkdownDescription: "Optimistic locking token; must be supplied on update operations.",
			Computed:            true,
		},
		"user_id": dsschema.SingleNestedAttribute{
			MarkdownDescription: "User-ID configuration for user-based policy enforcement.",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"agent_name": dsschema.StringAttribute{
					MarkdownDescription: "The name of the User-ID agent.",
					Computed:            true,
				},
				"cfturl_link": dsschema.StringAttribute{
					MarkdownDescription: "CloudFormation template URL link for User-ID setup.",
					Computed:            true,
				},
				"collector_name": dsschema.StringAttribute{
					MarkdownDescription: "The name of the log collector used by User-ID.",
					Computed:            true,
				},
				"custom_include_exclude_network": dsschema.ListNestedAttribute{
					MarkdownDescription: "Custom include/exclude network definitions for User-ID.",
					Computed:            true,
					NestedObject: dsschema.NestedAttributeObject{
						Attributes: map[string]dsschema.Attribute{
							"discovery_include": dsschema.BoolAttribute{
								MarkdownDescription: "Whether this network is included in User-ID discovery.",
								Computed:            true,
							},
							"enabled": dsschema.BoolAttribute{
								MarkdownDescription: "Whether this network entry is active.",
								Computed:            true,
							},
							"name": dsschema.StringAttribute{
								MarkdownDescription: "Name of the network entry.",
								Computed:            true,
							},
							"network_address": dsschema.StringAttribute{
								MarkdownDescription: "The network address (CIDR) for this entry.",
								Computed:            true,
							},
						},
					},
				},
				"enabled": dsschema.BoolAttribute{
					MarkdownDescription: "Whether User-ID is enabled.",
					Computed:            true,
				},
				"endpoint_dns": dsschema.StringAttribute{
					MarkdownDescription: "DNS name of the User-ID endpoint.",
					Computed:            true,
				},
				"port": dsschema.Int64Attribute{
					MarkdownDescription: "Port used by the User-ID agent.",
					Computed:            true,
				},
				"secret_key_arn": dsschema.StringAttribute{
					MarkdownDescription: "ARN of the AWS Secrets Manager secret for User-ID credentials.",
					Computed:            true,
				},
				"user_id_status": dsschema.StringAttribute{
					MarkdownDescription: "The operational status of the User-ID feature. Possible values are `EndpointPending`, `EndpointFail`, `CommitFail`, `Enabled`, `Disabled`, `DisablePending`, `UpdatePending` and `UpdateFail`.",
					Computed:            true,
				},
			},
		},
	},
}

// NgfirewallCloudwatchMetricsDataSourceSchema defines the schema for NgfirewallCloudwatchMetrics.
var NgfirewallCloudwatchMetricsDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "NgfirewallCloudwatchMetric data source",
	Attributes: map[string]dsschema.Attribute{
		"account_id": dsschema.StringAttribute{
			MarkdownDescription: "The AWS account ID where CloudWatch metrics are published.",
			Computed:            true,
		},
		"cloud_watch_metric_namespace": dsschema.StringAttribute{
			MarkdownDescription: "The CloudWatch metric namespace.",
			Computed:            true,
		},
		"cloud_watch_metrics_fields": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "The CloudWatch metric field names to publish.",
			Computed:            true,
		},
	},
}

// NgfirewallEgressNatConfigDataSourceSchema defines the schema for NgfirewallEgressNatConfig.
var NgfirewallEgressNatConfigDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "NgfirewallEgressNatConfig data source",
	Attributes: map[string]dsschema.Attribute{
		"enabled": dsschema.BoolAttribute{
			MarkdownDescription: "Whether egress NAT is enabled on this firewall.",
			Computed:            true,
		},
		"settings": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Settings for egress NAT IP address management.",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"ip_pool_type": dsschema.StringAttribute{
					MarkdownDescription: "The IP pool type for egress NAT. Possible values are `AWSService` and `BYOIP`.",
					Computed:            true,
				},
				"ipam_pool_id": dsschema.StringAttribute{
					MarkdownDescription: "The IPAM pool ID used for egress NAT IP allocation.",
					Computed:            true,
				},
			},
		},
	},
}

// NgfirewallEgressNatSettingsDataSourceSchema defines the schema for NgfirewallEgressNatSettings.
var NgfirewallEgressNatSettingsDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "NgfirewallEgressNatSetting data source",
	Attributes: map[string]dsschema.Attribute{
		"ip_pool_type": dsschema.StringAttribute{
			MarkdownDescription: "The IP pool type for egress NAT. Possible values are `AWSService` and `BYOIP`.",
			Computed:            true,
		},
		"ipam_pool_id": dsschema.StringAttribute{
			MarkdownDescription: "The IPAM pool ID used for egress NAT IP allocation.",
			Computed:            true,
		},
	},
}

// NgfirewallEndpointDataSourceSchema defines the schema for NgfirewallEndpoint.
var NgfirewallEndpointDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "NgfirewallEndpoint data source",
	Attributes: map[string]dsschema.Attribute{
		"account_id": dsschema.StringAttribute{
			MarkdownDescription: "The AWS account ID that owns this endpoint.",
			Computed:            true,
		},
		"egress_nat_enabled": dsschema.BoolAttribute{
			MarkdownDescription: "Whether egress NAT is enabled for this endpoint.",
			Computed:            true,
		},
		"endpoint_id": dsschema.StringAttribute{
			MarkdownDescription: "The VPC endpoint ID.",
			Computed:            true,
		},
		"mode": dsschema.StringAttribute{
			MarkdownDescription: "The endpoint management mode. Possible values are `ServiceManaged` and `CustomerManaged`.",
			Computed:            true,
		},
		"prefixes": dsschema.SingleNestedAttribute{
			MarkdownDescription: "IP prefix lists associated with a firewall endpoint.",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"private_prefix": dsschema.SingleNestedAttribute{
					MarkdownDescription: "A list of CIDR prefixes.",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"cidrs": dsschema.ListAttribute{
							ElementType:         types.StringType,
							MarkdownDescription: "CIDR address blocks.",
							Computed:            true,
						},
					},
				},
				"public_prefix": dsschema.SingleNestedAttribute{
					MarkdownDescription: "A list of CIDR prefixes.",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"cidrs": dsschema.ListAttribute{
							ElementType:         types.StringType,
							MarkdownDescription: "CIDR address blocks.",
							Computed:            true,
						},
					},
				},
			},
		},
		"rejected_reason": dsschema.StringAttribute{
			MarkdownDescription: "Human-readable reason if the endpoint was rejected.",
			Computed:            true,
		},
		"status": dsschema.StringAttribute{
			MarkdownDescription: "The current provisioning status of the endpoint. Possible values are `Pending`, `Accepted`, `Rejected`, `Created`, `Creating` and `Deleting`.",
			Computed:            true,
		},
		"subnet_id": dsschema.StringAttribute{
			MarkdownDescription: "The AWS subnet ID in which the endpoint resides.",
			Computed:            true,
		},
		"vpc_id": dsschema.StringAttribute{
			MarkdownDescription: "The AWS VPC ID in which the endpoint resides.",
			Computed:            true,
		},
		"zone_id": dsschema.StringAttribute{
			MarkdownDescription: "The availability zone ID for this endpoint.",
			Computed:            true,
		},
	},
}

// NgfirewallEndpointPrefixesDataSourceSchema defines the schema for NgfirewallEndpointPrefixes.
var NgfirewallEndpointPrefixesDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "NgfirewallEndpointPrefix data source",
	Attributes: map[string]dsschema.Attribute{
		"private_prefix": dsschema.SingleNestedAttribute{
			MarkdownDescription: "A list of CIDR prefixes.",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"cidrs": dsschema.ListAttribute{
					ElementType:         types.StringType,
					MarkdownDescription: "CIDR address blocks.",
					Computed:            true,
				},
			},
		},
		"public_prefix": dsschema.SingleNestedAttribute{
			MarkdownDescription: "A list of CIDR prefixes.",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"cidrs": dsschema.ListAttribute{
					ElementType:         types.StringType,
					MarkdownDescription: "CIDR address blocks.",
					Computed:            true,
				},
			},
		},
	},
}

// NgfirewallPrefixListDataSourceSchema defines the schema for NgfirewallPrefixList.
var NgfirewallPrefixListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "NgfirewallPrefixList data source",
	Attributes: map[string]dsschema.Attribute{
		"cidrs": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "CIDR address blocks.",
			Computed:            true,
		},
	},
}

// NgfirewallGwlbConfigDataSourceSchema defines the schema for NgfirewallGwlbConfig.
var NgfirewallGwlbConfigDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "NgfirewallGwlbConfig data source",
	Attributes: map[string]dsschema.Attribute{
		"deregistration_delay": dsschema.Int64Attribute{
			MarkdownDescription: "Deregistration delay in seconds.",
			Computed:            true,
		},
		"rejected_flow_count_alert_threshold": dsschema.Int64Attribute{
			MarkdownDescription: "Alert threshold for the total rejected flow count.",
			Computed:            true,
		},
		"rejected_flow_count_tcp_alert_threshold": dsschema.Int64Attribute{
			MarkdownDescription: "Alert threshold for the TCP rejected flow count.",
			Computed:            true,
		},
		"session_rebalance_enabled": dsschema.BoolAttribute{
			MarkdownDescription: "Whether session rebalancing is enabled.",
			Computed:            true,
		},
		"tcp_idle_timeout": dsschema.Int64Attribute{
			MarkdownDescription: "TCP idle timeout in seconds.",
			Computed:            true,
		},
	},
}

// NgfirewallIpv6ConfigDataSourceSchema defines the schema for NgfirewallIpv6Config.
var NgfirewallIpv6ConfigDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "NgfirewallIpv6Config data source",
	Attributes: map[string]dsschema.Attribute{
		"enabled": dsschema.BoolAttribute{
			MarkdownDescription: "Whether IPv6 is enabled on this firewall.",
			Computed:            true,
		},
	},
}

// NgfirewallLogConfigDataSourceSchema defines the schema for NgfirewallLogConfig.
var NgfirewallLogConfigDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "NgfirewallLogConfig data source",
	Attributes: map[string]dsschema.Attribute{
		"account_id": dsschema.StringAttribute{
			MarkdownDescription: "The AWS account ID where logs are sent.",
			Computed:            true,
		},
		"log_destination": dsschema.StringAttribute{
			MarkdownDescription: "The ARN or name of the log destination resource.",
			Computed:            true,
		},
		"log_destination_region": dsschema.StringAttribute{
			MarkdownDescription: "The AWS region of the log destination.",
			Computed:            true,
		},
		"log_destination_type": dsschema.StringAttribute{
			MarkdownDescription: "The type of log destination. Possible values are `S3`, `CloudWatchLogs` and `KinesisDataFirehose`.",
			Computed:            true,
		},
		"log_type": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "The types of logs to send to this destination. Possible values are `TRAFFIC`, `DECRYPTION`, `THREAT` and `USER_ID`.",
			Computed:            true,
		},
		"role_type": dsschema.StringAttribute{
			MarkdownDescription: "The IAM role type used to publish logs. Possible values are `ResourceBased` and `IamBased`.",
			Computed:            true,
		},
	},
}

// NgfirewallNotificationDataSourceSchema defines the schema for NgfirewallNotification.
var NgfirewallNotificationDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "NgfirewallNotification data source",
	Attributes: map[string]dsschema.Attribute{
		"expiry_time": dsschema.Int64Attribute{
			MarkdownDescription: "Unix timestamp when this notification expires.",
			Computed:            true,
		},
		"id": dsschema.StringAttribute{
			MarkdownDescription: "Unique identifier of the notification.",
			Required:            true,
		},
		"link": dsschema.SingleNestedAttribute{
			MarkdownDescription: "A hyperlink associated with a notification.",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"target": dsschema.StringAttribute{
					MarkdownDescription: "The URL target of the link.",
					Computed:            true,
				},
				"text": dsschema.StringAttribute{
					MarkdownDescription: "The display text of the link.",
					Computed:            true,
				},
			},
		},
		"msg": dsschema.StringAttribute{
			MarkdownDescription: "The notification message text.",
			Computed:            true,
		},
	},
}

// NgfirewallNotificationLinkDataSourceSchema defines the schema for NgfirewallNotificationLink.
var NgfirewallNotificationLinkDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "NgfirewallNotificationLink data source",
	Attributes: map[string]dsschema.Attribute{
		"target": dsschema.StringAttribute{
			MarkdownDescription: "The URL target of the link.",
			Computed:            true,
		},
		"text": dsschema.StringAttribute{
			MarkdownDescription: "The display text of the link.",
			Computed:            true,
		},
	},
}

// NgfirewallPrivateAccessDataSourceSchema defines the schema for NgfirewallPrivateAccess.
var NgfirewallPrivateAccessDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "NgfirewallPrivateAccess data source",
	Attributes: map[string]dsschema.Attribute{
		"resource_id": dsschema.StringAttribute{
			MarkdownDescription: "The resource ID of the private access endpoint.",
			Computed:            true,
		},
		"type": dsschema.StringAttribute{
			MarkdownDescription: "The type of private access resource. Possible values are `NetworkLoadBalancer`.",
			Computed:            true,
		},
	},
}

// NgfirewallTagDataSourceSchema defines the schema for NgfirewallTag.
var NgfirewallTagDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "NgfirewallTag data source",
	Attributes: map[string]dsschema.Attribute{
		"key": dsschema.StringAttribute{
			MarkdownDescription: "The tag key.",
			Computed:            true,
		},
		"value": dsschema.StringAttribute{
			MarkdownDescription: "The tag value.",
			Computed:            true,
		},
	},
}

// NgfirewallUserIdDataSourceSchema defines the schema for NgfirewallUserId.
var NgfirewallUserIdDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "NgfirewallUserId data source",
	Attributes: map[string]dsschema.Attribute{
		"agent_name": dsschema.StringAttribute{
			MarkdownDescription: "The name of the User-ID agent.",
			Computed:            true,
		},
		"cfturl_link": dsschema.StringAttribute{
			MarkdownDescription: "CloudFormation template URL link for User-ID setup.",
			Computed:            true,
		},
		"collector_name": dsschema.StringAttribute{
			MarkdownDescription: "The name of the log collector used by User-ID.",
			Computed:            true,
		},
		"custom_include_exclude_network": dsschema.ListNestedAttribute{
			MarkdownDescription: "Custom include/exclude network definitions for User-ID.",
			Computed:            true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: map[string]dsschema.Attribute{
					"discovery_include": dsschema.BoolAttribute{
						MarkdownDescription: "Whether this network is included in User-ID discovery.",
						Computed:            true,
					},
					"enabled": dsschema.BoolAttribute{
						MarkdownDescription: "Whether this network entry is active.",
						Computed:            true,
					},
					"name": dsschema.StringAttribute{
						MarkdownDescription: "Name of the network entry.",
						Computed:            true,
					},
					"network_address": dsschema.StringAttribute{
						MarkdownDescription: "The network address (CIDR) for this entry.",
						Computed:            true,
					},
				},
			},
		},
		"enabled": dsschema.BoolAttribute{
			MarkdownDescription: "Whether User-ID is enabled.",
			Computed:            true,
		},
		"endpoint_dns": dsschema.StringAttribute{
			MarkdownDescription: "DNS name of the User-ID endpoint.",
			Computed:            true,
		},
		"port": dsschema.Int64Attribute{
			MarkdownDescription: "Port used by the User-ID agent.",
			Computed:            true,
		},
		"secret_key_arn": dsschema.StringAttribute{
			MarkdownDescription: "ARN of the AWS Secrets Manager secret for User-ID credentials.",
			Computed:            true,
		},
		"user_id_status": dsschema.StringAttribute{
			MarkdownDescription: "The operational status of the User-ID feature. Possible values are `EndpointPending`, `EndpointFail`, `CommitFail`, `Enabled`, `Disabled`, `DisablePending`, `UpdatePending` and `UpdateFail`.",
			Computed:            true,
		},
	},
}

// NgfirewallUserIdNetworkDataSourceSchema defines the schema for NgfirewallUserIdNetwork.
var NgfirewallUserIdNetworkDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "NgfirewallUserIdNetwork data source",
	Attributes: map[string]dsschema.Attribute{
		"discovery_include": dsschema.BoolAttribute{
			MarkdownDescription: "Whether this network is included in User-ID discovery.",
			Computed:            true,
		},
		"enabled": dsschema.BoolAttribute{
			MarkdownDescription: "Whether this network entry is active.",
			Computed:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "Name of the network entry.",
			Optional:            true,
			Computed:            true,
		},
		"network_address": dsschema.StringAttribute{
			MarkdownDescription: "The network address (CIDR) for this entry.",
			Computed:            true,
		},
	},
}

// NgfirewallStatusDataSourceSchema defines the schema for NgfirewallStatus.
var NgfirewallStatusDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "NgfirewallStatus data source",
	Attributes: map[string]dsschema.Attribute{
		"device_rule_stack_commit_info": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Commit information for a rulestack.",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"commit_messages": dsschema.ListAttribute{
					ElementType:         types.StringType,
					MarkdownDescription: "Messages produced during the commit operation.",
					Computed:            true,
				},
				"commit_ts": dsschema.StringAttribute{
					MarkdownDescription: "Timestamp of the last commit operation.",
					Computed:            true,
				},
			},
		},
		"device_rule_stack_commit_state": dsschema.StringAttribute{
			MarkdownDescription: "The commit state of the device rulestack. Possible values are `Submitted` and `None`.",
			Computed:            true,
		},
		"device_rule_stack_commit_status": dsschema.StringAttribute{
			MarkdownDescription: "The commit status of the device rulestack. Possible values are `Running`, `Uncommitted`, `Precommit`, `PrecommitDone`, `PrevalidateDone`, `Committing`, `Failed` and `Success`.",
			Computed:            true,
		},
		"failure_reason": dsschema.StringAttribute{
			MarkdownDescription: "Human-readable reason for a firewall failure, if any.",
			Computed:            true,
		},
		"firewall_status": dsschema.StringAttribute{
			MarkdownDescription: "The operational lifecycle status of the firewall. Possible values are `CREATING`, `UPDATING`, `DELETING`, `CREATE_COMPLETE`, `UPDATE_COMPLETE`, `CREATE_FAIL`, `UPDATE_FAIL`, `DELETE_FAIL` and `DELETE_COMPLETE`.",
			Computed:            true,
		},
		"global_rule_stack_commit_info": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Commit information for a rulestack.",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"commit_messages": dsschema.ListAttribute{
					ElementType:         types.StringType,
					MarkdownDescription: "Messages produced during the commit operation.",
					Computed:            true,
				},
				"commit_ts": dsschema.StringAttribute{
					MarkdownDescription: "Timestamp of the last commit operation.",
					Computed:            true,
				},
			},
		},
		"global_rule_stack_status": dsschema.StringAttribute{
			MarkdownDescription: "The commit status of the global rulestack. Possible values are `Running`, `Uncommitted`, `Precommit`, `PrecommitDone`, `PrevalidateDone`, `Committing`, `Failed` and `Success`.",
			Computed:            true,
		},
		"public_i_ps": dsschema.ListNestedAttribute{
			MarkdownDescription: "Public IP addresses assigned to the firewall.",
			Computed:            true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: map[string]dsschema.Attribute{
					"ip_address": dsschema.StringAttribute{
						MarkdownDescription: "The public IP address.",
						Computed:            true,
					},
					"ip_source": dsschema.StringAttribute{
						MarkdownDescription: "The source of the IP address allocation.",
						Computed:            true,
					},
					"ip_status": dsschema.StringAttribute{
						MarkdownDescription: "The current status of this public IP.",
						Computed:            true,
					},
				},
			},
		},
		"rule_stack_commit_info": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Commit information for a rulestack.",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"commit_messages": dsschema.ListAttribute{
					ElementType:         types.StringType,
					MarkdownDescription: "Messages produced during the commit operation.",
					Computed:            true,
				},
				"commit_ts": dsschema.StringAttribute{
					MarkdownDescription: "Timestamp of the last commit operation.",
					Computed:            true,
				},
			},
		},
		"rule_stack_status": dsschema.StringAttribute{
			MarkdownDescription: "The commit status of the local rulestack. Possible values are `Running`, `Uncommitted`, `Precommit`, `PrecommitDone`, `PrevalidateDone`, `Committing`, `Failed` and `Success`.",
			Computed:            true,
		},
		"scm_assoc_status": dsschema.BoolAttribute{
			MarkdownDescription: "Whether the firewall is associated with Strata Cloud Manager.",
			Computed:            true,
		},
	},
}

// NgfirewallCommitInfoDataSourceSchema defines the schema for NgfirewallCommitInfo.
var NgfirewallCommitInfoDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "NgfirewallCommitInfo data source",
	Attributes: map[string]dsschema.Attribute{
		"commit_messages": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Messages produced during the commit operation.",
			Computed:            true,
		},
		"commit_ts": dsschema.StringAttribute{
			MarkdownDescription: "Timestamp of the last commit operation.",
			Computed:            true,
		},
	},
}

// NgfirewallPublicIpDataSourceSchema defines the schema for NgfirewallPublicIp.
var NgfirewallPublicIpDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "NgfirewallPublicIp data source",
	Attributes: map[string]dsschema.Attribute{
		"ip_address": dsschema.StringAttribute{
			MarkdownDescription: "The public IP address.",
			Computed:            true,
		},
		"ip_source": dsschema.StringAttribute{
			MarkdownDescription: "The source of the IP address allocation.",
			Computed:            true,
		},
		"ip_status": dsschema.StringAttribute{
			MarkdownDescription: "The current status of this public IP.",
			Computed:            true,
		},
	},
}

// NgfirewallResponseStatusDataSourceSchema defines the schema for NgfirewallResponseStatus.
var NgfirewallResponseStatusDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "NgfirewallResponseStatus data source",
	Attributes: map[string]dsschema.Attribute{
		"error_code": dsschema.Int64Attribute{
			MarkdownDescription: "Non-zero error code if the API request failed.",
			Computed:            true,
		},
		"reason": dsschema.StringAttribute{
			MarkdownDescription: "Human-readable reason for the error, if any.",
			Computed:            true,
		},
	},
}

// NgfirewallCreateRequestDataSourceSchema defines the schema for NgfirewallCreateRequest.
var NgfirewallCreateRequestDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "NgfirewallCreateRequest data source",
	Attributes: map[string]dsschema.Attribute{
		"allow_list_accounts": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "AWS account IDs permitted to access the firewall endpoint service.",
			Computed:            true,
		},
		"change_protection": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "List of attributes protected from change.",
			Computed:            true,
		},
		"customer_zone_id_list": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "List of customer-defined security zone IDs. Required.",
			Computed:            true,
		},
		"description": dsschema.StringAttribute{
			MarkdownDescription: "A human-readable description of the firewall.",
			Computed:            true,
		},
		"egress_nat": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Egress NAT configuration for outbound traffic.",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"enabled": dsschema.BoolAttribute{
					MarkdownDescription: "Whether egress NAT is enabled on this firewall.",
					Computed:            true,
				},
				"settings": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Settings for egress NAT IP address management.",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"ip_pool_type": dsschema.StringAttribute{
							MarkdownDescription: "The IP pool type for egress NAT. Possible values are `AWSService` and `BYOIP`.",
							Computed:            true,
						},
						"ipam_pool_id": dsschema.StringAttribute{
							MarkdownDescription: "The IPAM pool ID used for egress NAT IP allocation.",
							Computed:            true,
						},
					},
				},
			},
		},
		"endpoints": dsschema.ListNestedAttribute{
			MarkdownDescription: "Firewall endpoint configurations. Assigned by the service when the firewall is created; can be changed with an update.",
			Computed:            true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: map[string]dsschema.Attribute{
					"account_id": dsschema.StringAttribute{
						MarkdownDescription: "The AWS account ID that owns this endpoint.",
						Computed:            true,
					},
					"egress_nat_enabled": dsschema.BoolAttribute{
						MarkdownDescription: "Whether egress NAT is enabled for this endpoint.",
						Computed:            true,
					},
					"endpoint_id": dsschema.StringAttribute{
						MarkdownDescription: "The VPC endpoint ID.",
						Computed:            true,
					},
					"mode": dsschema.StringAttribute{
						MarkdownDescription: "The endpoint management mode. Possible values are `ServiceManaged` and `CustomerManaged`.",
						Computed:            true,
					},
					"prefixes": dsschema.SingleNestedAttribute{
						MarkdownDescription: "IP prefix lists associated with a firewall endpoint.",
						Computed:            true,
						Attributes: map[string]dsschema.Attribute{
							"private_prefix": dsschema.SingleNestedAttribute{
								MarkdownDescription: "A list of CIDR prefixes.",
								Computed:            true,
								Attributes: map[string]dsschema.Attribute{
									"cidrs": dsschema.ListAttribute{
										ElementType:         types.StringType,
										MarkdownDescription: "CIDR address blocks.",
										Computed:            true,
									},
								},
							},
							"public_prefix": dsschema.SingleNestedAttribute{
								MarkdownDescription: "A list of CIDR prefixes.",
								Computed:            true,
								Attributes: map[string]dsschema.Attribute{
									"cidrs": dsschema.ListAttribute{
										ElementType:         types.StringType,
										MarkdownDescription: "CIDR address blocks.",
										Computed:            true,
									},
								},
							},
						},
					},
					"rejected_reason": dsschema.StringAttribute{
						MarkdownDescription: "Human-readable reason if the endpoint was rejected.",
						Computed:            true,
					},
					"status": dsschema.StringAttribute{
						MarkdownDescription: "The current provisioning status of the endpoint. Possible values are `Pending`, `Accepted`, `Rejected`, `Created`, `Creating` and `Deleting`.",
						Computed:            true,
					},
					"subnet_id": dsschema.StringAttribute{
						MarkdownDescription: "The AWS subnet ID in which the endpoint resides.",
						Computed:            true,
					},
					"vpc_id": dsschema.StringAttribute{
						MarkdownDescription: "The AWS VPC ID in which the endpoint resides.",
						Computed:            true,
					},
					"zone_id": dsschema.StringAttribute{
						MarkdownDescription: "The availability zone ID for this endpoint.",
						Computed:            true,
					},
				},
			},
		},
		"gwlb_tcp_idle_timeout": dsschema.Int64Attribute{
			MarkdownDescription: "GWLB TCP idle timeout in seconds. Defaults to 350; can be changed with an update.",
			Computed:            true,
		},
		"id": dsschema.StringAttribute{
			MarkdownDescription: "The ID of the Cloud NGFW firewall.",
			Required:            true,
		},
		"link_id": dsschema.StringAttribute{
			MarkdownDescription: "The link ID to associate with the firewall.",
			Computed:            true,
		},
		"mode": dsschema.StringAttribute{
			MarkdownDescription: "The operational mode of the firewall. Set when the firewall is created and cannot be changed afterward. Possible values are `CNGFW`.",
			Computed:            true,
		},
		"private_access": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Private access configuration for the firewall.",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"resource_id": dsschema.StringAttribute{
					MarkdownDescription: "The resource ID of the private access endpoint.",
					Computed:            true,
				},
				"type": dsschema.StringAttribute{
					MarkdownDescription: "The type of private access resource. Possible values are `NetworkLoadBalancer`.",
					Computed:            true,
				},
			},
		},
		"region": dsschema.StringAttribute{
			MarkdownDescription: "The AWS region where the Cloud NGFW firewall is deployed",
			Required:            true,
		},
		"scm_labels": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Strata Cloud Manager (SCM) labels to apply to the firewall.",
			Computed:            true,
		},
		"source": dsschema.StringAttribute{
			MarkdownDescription: "The source SDK or provider that created the firewall. Automatically set by the SDK or Terraform provider; do not set this field manually. Possible values are `SCM-SDK` and `Terraform-Provider-SCM`.",
			Computed:            true,
		},
		"tags": dsschema.ListNestedAttribute{
			MarkdownDescription: "AWS resource tags to apply to the firewall.",
			Computed:            true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: map[string]dsschema.Attribute{
					"key": dsschema.StringAttribute{
						MarkdownDescription: "The tag key.",
						Computed:            true,
					},
					"value": dsschema.StringAttribute{
						MarkdownDescription: "The tag value.",
						Computed:            true,
					},
				},
			},
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
		"tier": dsschema.StringAttribute{
			MarkdownDescription: "The service tier for the firewall, either standard or premium. Possible values are `standard` and `premium`.",
			Computed:            true,
		},
		"user_id": dsschema.SingleNestedAttribute{
			MarkdownDescription: "User-ID configuration for user-based policy enforcement.",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"agent_name": dsschema.StringAttribute{
					MarkdownDescription: "The name of the User-ID agent.",
					Computed:            true,
				},
				"cfturl_link": dsschema.StringAttribute{
					MarkdownDescription: "CloudFormation template URL link for User-ID setup.",
					Computed:            true,
				},
				"collector_name": dsschema.StringAttribute{
					MarkdownDescription: "The name of the log collector used by User-ID.",
					Computed:            true,
				},
				"custom_include_exclude_network": dsschema.ListNestedAttribute{
					MarkdownDescription: "Custom include/exclude network definitions for User-ID.",
					Computed:            true,
					NestedObject: dsschema.NestedAttributeObject{
						Attributes: map[string]dsschema.Attribute{
							"discovery_include": dsschema.BoolAttribute{
								MarkdownDescription: "Whether this network is included in User-ID discovery.",
								Computed:            true,
							},
							"enabled": dsschema.BoolAttribute{
								MarkdownDescription: "Whether this network entry is active.",
								Computed:            true,
							},
							"name": dsschema.StringAttribute{
								MarkdownDescription: "Name of the network entry.",
								Computed:            true,
							},
							"network_address": dsschema.StringAttribute{
								MarkdownDescription: "The network address (CIDR) for this entry.",
								Computed:            true,
							},
						},
					},
				},
				"enabled": dsschema.BoolAttribute{
					MarkdownDescription: "Whether User-ID is enabled.",
					Computed:            true,
				},
				"endpoint_dns": dsschema.StringAttribute{
					MarkdownDescription: "DNS name of the User-ID endpoint.",
					Computed:            true,
				},
				"port": dsschema.Int64Attribute{
					MarkdownDescription: "Port used by the User-ID agent.",
					Computed:            true,
				},
				"secret_key_arn": dsschema.StringAttribute{
					MarkdownDescription: "ARN of the AWS Secrets Manager secret for User-ID credentials.",
					Computed:            true,
				},
				"user_id_status": dsschema.StringAttribute{
					MarkdownDescription: "The operational status of the User-ID feature. Possible values are `EndpointPending`, `EndpointFail`, `CommitFail`, `Enabled`, `Disabled`, `DisablePending`, `UpdatePending` and `UpdateFail`.",
					Computed:            true,
				},
			},
		},
	},
}

// NgfirewallsListResponseDataSourceSchema defines the schema for NgfirewallsListResponse.
var NgfirewallsListResponseDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "NgfirewallsListResponse data source",
	Attributes: map[string]dsschema.Attribute{
		"response": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Inner response body containing the firewall list.",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"firewalls_describe": dsschema.ListNestedAttribute{
					MarkdownDescription: "List of detailed firewall descriptions.",
					Computed:            true,
					NestedObject: dsschema.NestedAttributeObject{
						Attributes: map[string]dsschema.Attribute{
							"firewall": dsschema.SingleNestedAttribute{
								MarkdownDescription: "Core firewall configuration and identity fields.",
								Computed:            true,
								Attributes: map[string]dsschema.Attribute{
									"advanced_threat_log": dsschema.BoolAttribute{
										MarkdownDescription: "Whether advanced threat logging is enabled.",
										Computed:            true,
									},
									"allow_list_accounts": dsschema.ListAttribute{
										ElementType:         types.StringType,
										MarkdownDescription: "AWS account IDs permitted to access the firewall endpoint service.",
										Computed:            true,
									},
									"change_protection": dsschema.ListAttribute{
										ElementType:         types.StringType,
										MarkdownDescription: "List of attributes protected from change.",
										Computed:            true,
									},
									"cloudwatch_metrics": dsschema.SingleNestedAttribute{
										MarkdownDescription: "CloudWatch metrics configuration.",
										Computed:            true,
										Attributes: map[string]dsschema.Attribute{
											"account_id": dsschema.StringAttribute{
												MarkdownDescription: "The AWS account ID where CloudWatch metrics are published.",
												Computed:            true,
											},
											"cloud_watch_metric_namespace": dsschema.StringAttribute{
												MarkdownDescription: "The CloudWatch metric namespace.",
												Computed:            true,
											},
											"cloud_watch_metrics_fields": dsschema.ListAttribute{
												ElementType:         types.StringType,
												MarkdownDescription: "The CloudWatch metric field names to publish.",
												Computed:            true,
											},
										},
									},
									"customer_zone_id_list": dsschema.ListAttribute{
										ElementType:         types.StringType,
										MarkdownDescription: "List of customer-defined security zone IDs.",
										Computed:            true,
									},
									"deployment_update_token": dsschema.StringAttribute{
										MarkdownDescription: "Token used for deployment update sequencing.",
										Computed:            true,
									},
									"description": dsschema.StringAttribute{
										MarkdownDescription: "A human-readable description of the firewall.",
										Computed:            true,
									},
									"device_rule_stack_commit_status": dsschema.StringAttribute{
										MarkdownDescription: "Commit status of the device rulestack on the firewall.",
										Computed:            true,
									},
									"egress_nat": dsschema.SingleNestedAttribute{
										MarkdownDescription: "Egress NAT configuration for outbound traffic.",
										Computed:            true,
										Attributes: map[string]dsschema.Attribute{
											"enabled": dsschema.BoolAttribute{
												MarkdownDescription: "Whether egress NAT is enabled on this firewall.",
												Computed:            true,
											},
											"settings": dsschema.SingleNestedAttribute{
												MarkdownDescription: "Settings for egress NAT IP address management.",
												Computed:            true,
												Attributes: map[string]dsschema.Attribute{
													"ip_pool_type": dsschema.StringAttribute{
														MarkdownDescription: "The IP pool type for egress NAT. Possible values are `AWSService` and `BYOIP`.",
														Computed:            true,
													},
													"ipam_pool_id": dsschema.StringAttribute{
														MarkdownDescription: "The IPAM pool ID used for egress NAT IP allocation.",
														Computed:            true,
													},
												},
											},
										},
									},
									"endpoint_service_name": dsschema.StringAttribute{
										MarkdownDescription: "The VPC endpoint service name (only populated in ServiceManaged mode).",
										Computed:            true,
									},
									"endpoints": dsschema.ListNestedAttribute{
										MarkdownDescription: "VPC endpoint configurations created for this firewall.",
										Computed:            true,
										NestedObject: dsschema.NestedAttributeObject{
											Attributes: map[string]dsschema.Attribute{
												"account_id": dsschema.StringAttribute{
													MarkdownDescription: "The AWS account ID that owns this endpoint.",
													Computed:            true,
												},
												"egress_nat_enabled": dsschema.BoolAttribute{
													MarkdownDescription: "Whether egress NAT is enabled for this endpoint.",
													Computed:            true,
												},
												"endpoint_id": dsschema.StringAttribute{
													MarkdownDescription: "The VPC endpoint ID.",
													Computed:            true,
												},
												"mode": dsschema.StringAttribute{
													MarkdownDescription: "The endpoint management mode. Possible values are `ServiceManaged` and `CustomerManaged`.",
													Computed:            true,
												},
												"prefixes": dsschema.SingleNestedAttribute{
													MarkdownDescription: "IP prefix lists associated with a firewall endpoint.",
													Computed:            true,
													Attributes: map[string]dsschema.Attribute{
														"private_prefix": dsschema.SingleNestedAttribute{
															MarkdownDescription: "A list of CIDR prefixes.",
															Computed:            true,
															Attributes: map[string]dsschema.Attribute{
																"cidrs": dsschema.ListAttribute{
																	ElementType:         types.StringType,
																	MarkdownDescription: "CIDR address blocks.",
																	Computed:            true,
																},
															},
														},
														"public_prefix": dsschema.SingleNestedAttribute{
															MarkdownDescription: "A list of CIDR prefixes.",
															Computed:            true,
															Attributes: map[string]dsschema.Attribute{
																"cidrs": dsschema.ListAttribute{
																	ElementType:         types.StringType,
																	MarkdownDescription: "CIDR address blocks.",
																	Computed:            true,
																},
															},
														},
													},
												},
												"rejected_reason": dsschema.StringAttribute{
													MarkdownDescription: "Human-readable reason if the endpoint was rejected.",
													Computed:            true,
												},
												"status": dsschema.StringAttribute{
													MarkdownDescription: "The current provisioning status of the endpoint. Possible values are `Pending`, `Accepted`, `Rejected`, `Created`, `Creating` and `Deleting`.",
													Computed:            true,
												},
												"subnet_id": dsschema.StringAttribute{
													MarkdownDescription: "The AWS subnet ID in which the endpoint resides.",
													Computed:            true,
												},
												"vpc_id": dsschema.StringAttribute{
													MarkdownDescription: "The AWS VPC ID in which the endpoint resides.",
													Computed:            true,
												},
												"zone_id": dsschema.StringAttribute{
													MarkdownDescription: "The availability zone ID for this endpoint.",
													Computed:            true,
												},
											},
										},
									},
									"firewall_id": dsschema.StringAttribute{
										MarkdownDescription: "The unique identifier of the firewall.",
										Computed:            true,
									},
									"gwlb": dsschema.SingleNestedAttribute{
										MarkdownDescription: "Gateway Load Balancer configuration.",
										Computed:            true,
										Attributes: map[string]dsschema.Attribute{
											"deregistration_delay": dsschema.Int64Attribute{
												MarkdownDescription: "Deregistration delay in seconds.",
												Computed:            true,
											},
											"rejected_flow_count_alert_threshold": dsschema.Int64Attribute{
												MarkdownDescription: "Alert threshold for the total rejected flow count.",
												Computed:            true,
											},
											"rejected_flow_count_tcp_alert_threshold": dsschema.Int64Attribute{
												MarkdownDescription: "Alert threshold for the TCP rejected flow count.",
												Computed:            true,
											},
											"session_rebalance_enabled": dsschema.BoolAttribute{
												MarkdownDescription: "Whether session rebalancing is enabled.",
												Computed:            true,
											},
											"tcp_idle_timeout": dsschema.Int64Attribute{
												MarkdownDescription: "TCP idle timeout in seconds.",
												Computed:            true,
											},
										},
									},
									"ipv6": dsschema.SingleNestedAttribute{
										MarkdownDescription: "IPv6 configuration for the firewall.",
										Computed:            true,
										Attributes: map[string]dsschema.Attribute{
											"enabled": dsschema.BoolAttribute{
												MarkdownDescription: "Whether IPv6 is enabled on this firewall.",
												Computed:            true,
											},
										},
									},
									"link_id": dsschema.StringAttribute{
										MarkdownDescription: "The link ID associated with the firewall.",
										Computed:            true,
									},
									"link_status": dsschema.StringAttribute{
										MarkdownDescription: "The current link status of the firewall.",
										Computed:            true,
									},
									"log_config": dsschema.SingleNestedAttribute{
										MarkdownDescription: "Log destination configuration for the firewall.",
										Computed:            true,
										Attributes: map[string]dsschema.Attribute{
											"account_id": dsschema.StringAttribute{
												MarkdownDescription: "The AWS account ID where logs are sent.",
												Computed:            true,
											},
											"log_destination": dsschema.StringAttribute{
												MarkdownDescription: "The ARN or name of the log destination resource.",
												Computed:            true,
											},
											"log_destination_region": dsschema.StringAttribute{
												MarkdownDescription: "The AWS region of the log destination.",
												Computed:            true,
											},
											"log_destination_type": dsschema.StringAttribute{
												MarkdownDescription: "The type of log destination. Possible values are `S3`, `CloudWatchLogs` and `KinesisDataFirehose`.",
												Computed:            true,
											},
											"log_type": dsschema.ListAttribute{
												ElementType:         types.StringType,
												MarkdownDescription: "The types of logs to send to this destination. Possible values are `TRAFFIC`, `DECRYPTION`, `THREAT` and `USER_ID`.",
												Computed:            true,
											},
											"role_type": dsschema.StringAttribute{
												MarkdownDescription: "The IAM role type used to publish logs. Possible values are `ResourceBased` and `IamBased`.",
												Computed:            true,
											},
										},
									},
									"notifications": dsschema.ListNestedAttribute{
										MarkdownDescription: "Notifications associated with this firewall.",
										Computed:            true,
										NestedObject: dsschema.NestedAttributeObject{
											Attributes: map[string]dsschema.Attribute{
												"expiry_time": dsschema.Int64Attribute{
													MarkdownDescription: "Unix timestamp when this notification expires.",
													Computed:            true,
												},
												"id": dsschema.StringAttribute{
													MarkdownDescription: "Unique identifier of the notification.",
													Computed:            true,
												},
												"link": dsschema.SingleNestedAttribute{
													MarkdownDescription: "A hyperlink associated with a notification.",
													Computed:            true,
													Attributes: map[string]dsschema.Attribute{
														"target": dsschema.StringAttribute{
															MarkdownDescription: "The URL target of the link.",
															Computed:            true,
														},
														"text": dsschema.StringAttribute{
															MarkdownDescription: "The display text of the link.",
															Computed:            true,
														},
													},
												},
												"msg": dsschema.StringAttribute{
													MarkdownDescription: "The notification message text.",
													Computed:            true,
												},
											},
										},
									},
									"private_access": dsschema.SingleNestedAttribute{
										MarkdownDescription: "Private access configuration for the firewall.",
										Computed:            true,
										Attributes: map[string]dsschema.Attribute{
											"resource_id": dsschema.StringAttribute{
												MarkdownDescription: "The resource ID of the private access endpoint.",
												Computed:            true,
											},
											"type": dsschema.StringAttribute{
												MarkdownDescription: "The type of private access resource. Possible values are `NetworkLoadBalancer`.",
												Computed:            true,
											},
										},
									},
									"region": dsschema.StringAttribute{
										MarkdownDescription: "The AWS region where the firewall is deployed.",
										Computed:            true,
									},
									"software_version": dsschema.StringAttribute{
										MarkdownDescription: "The PAN-OS software version currently running on the firewall.",
										Computed:            true,
									},
									"tags": dsschema.ListNestedAttribute{
										MarkdownDescription: "AWS resource tags applied to the firewall.",
										Computed:            true,
										NestedObject: dsschema.NestedAttributeObject{
											Attributes: map[string]dsschema.Attribute{
												"key": dsschema.StringAttribute{
													MarkdownDescription: "The tag key.",
													Computed:            true,
												},
												"value": dsschema.StringAttribute{
													MarkdownDescription: "The tag value.",
													Computed:            true,
												},
											},
										},
									},
									"tier": dsschema.StringAttribute{
										MarkdownDescription: "The service tier of the firewall (base, standard, or premium).",
										Computed:            true,
									},
									"update_token": dsschema.StringAttribute{
										MarkdownDescription: "Optimistic locking token; must be supplied on update operations.",
										Computed:            true,
									},
									"user_id": dsschema.SingleNestedAttribute{
										MarkdownDescription: "User-ID configuration for user-based policy enforcement.",
										Computed:            true,
										Attributes: map[string]dsschema.Attribute{
											"agent_name": dsschema.StringAttribute{
												MarkdownDescription: "The name of the User-ID agent.",
												Computed:            true,
											},
											"cfturl_link": dsschema.StringAttribute{
												MarkdownDescription: "CloudFormation template URL link for User-ID setup.",
												Computed:            true,
											},
											"collector_name": dsschema.StringAttribute{
												MarkdownDescription: "The name of the log collector used by User-ID.",
												Computed:            true,
											},
											"custom_include_exclude_network": dsschema.ListNestedAttribute{
												MarkdownDescription: "Custom include/exclude network definitions for User-ID.",
												Computed:            true,
												NestedObject: dsschema.NestedAttributeObject{
													Attributes: map[string]dsschema.Attribute{
														"discovery_include": dsschema.BoolAttribute{
															MarkdownDescription: "Whether this network is included in User-ID discovery.",
															Computed:            true,
														},
														"enabled": dsschema.BoolAttribute{
															MarkdownDescription: "Whether this network entry is active.",
															Computed:            true,
														},
														"name": dsschema.StringAttribute{
															MarkdownDescription: "Name of the network entry.",
															Computed:            true,
														},
														"network_address": dsschema.StringAttribute{
															MarkdownDescription: "The network address (CIDR) for this entry.",
															Computed:            true,
														},
													},
												},
											},
											"enabled": dsschema.BoolAttribute{
												MarkdownDescription: "Whether User-ID is enabled.",
												Computed:            true,
											},
											"endpoint_dns": dsschema.StringAttribute{
												MarkdownDescription: "DNS name of the User-ID endpoint.",
												Computed:            true,
											},
											"port": dsschema.Int64Attribute{
												MarkdownDescription: "Port used by the User-ID agent.",
												Computed:            true,
											},
											"secret_key_arn": dsschema.StringAttribute{
												MarkdownDescription: "ARN of the AWS Secrets Manager secret for User-ID credentials.",
												Computed:            true,
											},
											"user_id_status": dsschema.StringAttribute{
												MarkdownDescription: "The operational status of the User-ID feature. Possible values are `EndpointPending`, `EndpointFail`, `CommitFail`, `Enabled`, `Disabled`, `DisablePending`, `UpdatePending` and `UpdateFail`.",
												Computed:            true,
											},
										},
									},
								},
							},
							"status": dsschema.SingleNestedAttribute{
								MarkdownDescription: "Operational status of the firewall.",
								Computed:            true,
								Attributes: map[string]dsschema.Attribute{
									"device_rule_stack_commit_info": dsschema.SingleNestedAttribute{
										MarkdownDescription: "Commit information for a rulestack.",
										Computed:            true,
										Attributes: map[string]dsschema.Attribute{
											"commit_messages": dsschema.ListAttribute{
												ElementType:         types.StringType,
												MarkdownDescription: "Messages produced during the commit operation.",
												Computed:            true,
											},
											"commit_ts": dsschema.StringAttribute{
												MarkdownDescription: "Timestamp of the last commit operation.",
												Computed:            true,
											},
										},
									},
									"device_rule_stack_commit_state": dsschema.StringAttribute{
										MarkdownDescription: "The commit state of the device rulestack. Possible values are `Submitted` and `None`.",
										Computed:            true,
									},
									"device_rule_stack_commit_status": dsschema.StringAttribute{
										MarkdownDescription: "The commit status of the device rulestack. Possible values are `Running`, `Uncommitted`, `Precommit`, `PrecommitDone`, `PrevalidateDone`, `Committing`, `Failed` and `Success`.",
										Computed:            true,
									},
									"failure_reason": dsschema.StringAttribute{
										MarkdownDescription: "Human-readable reason for a firewall failure, if any.",
										Computed:            true,
									},
									"firewall_status": dsschema.StringAttribute{
										MarkdownDescription: "The operational lifecycle status of the firewall. Possible values are `CREATING`, `UPDATING`, `DELETING`, `CREATE_COMPLETE`, `UPDATE_COMPLETE`, `CREATE_FAIL`, `UPDATE_FAIL`, `DELETE_FAIL` and `DELETE_COMPLETE`.",
										Computed:            true,
									},
									"global_rule_stack_commit_info": dsschema.SingleNestedAttribute{
										MarkdownDescription: "Commit information for a rulestack.",
										Computed:            true,
										Attributes: map[string]dsschema.Attribute{
											"commit_messages": dsschema.ListAttribute{
												ElementType:         types.StringType,
												MarkdownDescription: "Messages produced during the commit operation.",
												Computed:            true,
											},
											"commit_ts": dsschema.StringAttribute{
												MarkdownDescription: "Timestamp of the last commit operation.",
												Computed:            true,
											},
										},
									},
									"global_rule_stack_status": dsschema.StringAttribute{
										MarkdownDescription: "The commit status of the global rulestack. Possible values are `Running`, `Uncommitted`, `Precommit`, `PrecommitDone`, `PrevalidateDone`, `Committing`, `Failed` and `Success`.",
										Computed:            true,
									},
									"public_i_ps": dsschema.ListNestedAttribute{
										MarkdownDescription: "Public IP addresses assigned to the firewall.",
										Computed:            true,
										NestedObject: dsschema.NestedAttributeObject{
											Attributes: map[string]dsschema.Attribute{
												"ip_address": dsschema.StringAttribute{
													MarkdownDescription: "The public IP address.",
													Computed:            true,
												},
												"ip_source": dsschema.StringAttribute{
													MarkdownDescription: "The source of the IP address allocation.",
													Computed:            true,
												},
												"ip_status": dsschema.StringAttribute{
													MarkdownDescription: "The current status of this public IP.",
													Computed:            true,
												},
											},
										},
									},
									"rule_stack_commit_info": dsschema.SingleNestedAttribute{
										MarkdownDescription: "Commit information for a rulestack.",
										Computed:            true,
										Attributes: map[string]dsschema.Attribute{
											"commit_messages": dsschema.ListAttribute{
												ElementType:         types.StringType,
												MarkdownDescription: "Messages produced during the commit operation.",
												Computed:            true,
											},
											"commit_ts": dsschema.StringAttribute{
												MarkdownDescription: "Timestamp of the last commit operation.",
												Computed:            true,
											},
										},
									},
									"rule_stack_status": dsschema.StringAttribute{
										MarkdownDescription: "The commit status of the local rulestack. Possible values are `Running`, `Uncommitted`, `Precommit`, `PrecommitDone`, `PrevalidateDone`, `Committing`, `Failed` and `Success`.",
										Computed:            true,
									},
									"scm_assoc_status": dsschema.BoolAttribute{
										MarkdownDescription: "Whether the firewall is associated with Strata Cloud Manager.",
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
		},
		"response_status": dsschema.SingleNestedAttribute{
			MarkdownDescription: "API-level response metadata at the top level of the envelope.",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"error_code": dsschema.Int64Attribute{
					MarkdownDescription: "Non-zero error code if the API request failed.",
					Computed:            true,
				},
				"reason": dsschema.StringAttribute{
					MarkdownDescription: "Human-readable reason for the error, if any.",
					Computed:            true,
				},
			},
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
	},
}

// NgfirewallsListResponseBodyDataSourceSchema defines the schema for NgfirewallsListResponseBody.
var NgfirewallsListResponseBodyDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "NgfirewallsListResponseBody data source",
	Attributes: map[string]dsschema.Attribute{
		"firewalls_describe": dsschema.ListNestedAttribute{
			MarkdownDescription: "List of detailed firewall descriptions.",
			Computed:            true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: map[string]dsschema.Attribute{
					"firewall": dsschema.SingleNestedAttribute{
						MarkdownDescription: "Core firewall configuration and identity fields.",
						Computed:            true,
						Attributes: map[string]dsschema.Attribute{
							"advanced_threat_log": dsschema.BoolAttribute{
								MarkdownDescription: "Whether advanced threat logging is enabled.",
								Computed:            true,
							},
							"allow_list_accounts": dsschema.ListAttribute{
								ElementType:         types.StringType,
								MarkdownDescription: "AWS account IDs permitted to access the firewall endpoint service.",
								Computed:            true,
							},
							"change_protection": dsschema.ListAttribute{
								ElementType:         types.StringType,
								MarkdownDescription: "List of attributes protected from change.",
								Computed:            true,
							},
							"cloudwatch_metrics": dsschema.SingleNestedAttribute{
								MarkdownDescription: "CloudWatch metrics configuration.",
								Computed:            true,
								Attributes: map[string]dsschema.Attribute{
									"account_id": dsschema.StringAttribute{
										MarkdownDescription: "The AWS account ID where CloudWatch metrics are published.",
										Computed:            true,
									},
									"cloud_watch_metric_namespace": dsschema.StringAttribute{
										MarkdownDescription: "The CloudWatch metric namespace.",
										Computed:            true,
									},
									"cloud_watch_metrics_fields": dsschema.ListAttribute{
										ElementType:         types.StringType,
										MarkdownDescription: "The CloudWatch metric field names to publish.",
										Computed:            true,
									},
								},
							},
							"customer_zone_id_list": dsschema.ListAttribute{
								ElementType:         types.StringType,
								MarkdownDescription: "List of customer-defined security zone IDs.",
								Computed:            true,
							},
							"deployment_update_token": dsschema.StringAttribute{
								MarkdownDescription: "Token used for deployment update sequencing.",
								Computed:            true,
							},
							"description": dsschema.StringAttribute{
								MarkdownDescription: "A human-readable description of the firewall.",
								Computed:            true,
							},
							"device_rule_stack_commit_status": dsschema.StringAttribute{
								MarkdownDescription: "Commit status of the device rulestack on the firewall.",
								Computed:            true,
							},
							"egress_nat": dsschema.SingleNestedAttribute{
								MarkdownDescription: "Egress NAT configuration for outbound traffic.",
								Computed:            true,
								Attributes: map[string]dsschema.Attribute{
									"enabled": dsschema.BoolAttribute{
										MarkdownDescription: "Whether egress NAT is enabled on this firewall.",
										Computed:            true,
									},
									"settings": dsschema.SingleNestedAttribute{
										MarkdownDescription: "Settings for egress NAT IP address management.",
										Computed:            true,
										Attributes: map[string]dsschema.Attribute{
											"ip_pool_type": dsschema.StringAttribute{
												MarkdownDescription: "The IP pool type for egress NAT. Possible values are `AWSService` and `BYOIP`.",
												Computed:            true,
											},
											"ipam_pool_id": dsschema.StringAttribute{
												MarkdownDescription: "The IPAM pool ID used for egress NAT IP allocation.",
												Computed:            true,
											},
										},
									},
								},
							},
							"endpoint_service_name": dsschema.StringAttribute{
								MarkdownDescription: "The VPC endpoint service name (only populated in ServiceManaged mode).",
								Computed:            true,
							},
							"endpoints": dsschema.ListNestedAttribute{
								MarkdownDescription: "VPC endpoint configurations created for this firewall.",
								Computed:            true,
								NestedObject: dsschema.NestedAttributeObject{
									Attributes: map[string]dsschema.Attribute{
										"account_id": dsschema.StringAttribute{
											MarkdownDescription: "The AWS account ID that owns this endpoint.",
											Computed:            true,
										},
										"egress_nat_enabled": dsschema.BoolAttribute{
											MarkdownDescription: "Whether egress NAT is enabled for this endpoint.",
											Computed:            true,
										},
										"endpoint_id": dsschema.StringAttribute{
											MarkdownDescription: "The VPC endpoint ID.",
											Computed:            true,
										},
										"mode": dsschema.StringAttribute{
											MarkdownDescription: "The endpoint management mode. Possible values are `ServiceManaged` and `CustomerManaged`.",
											Computed:            true,
										},
										"prefixes": dsschema.SingleNestedAttribute{
											MarkdownDescription: "IP prefix lists associated with a firewall endpoint.",
											Computed:            true,
											Attributes: map[string]dsschema.Attribute{
												"private_prefix": dsschema.SingleNestedAttribute{
													MarkdownDescription: "A list of CIDR prefixes.",
													Computed:            true,
													Attributes: map[string]dsschema.Attribute{
														"cidrs": dsschema.ListAttribute{
															ElementType:         types.StringType,
															MarkdownDescription: "CIDR address blocks.",
															Computed:            true,
														},
													},
												},
												"public_prefix": dsschema.SingleNestedAttribute{
													MarkdownDescription: "A list of CIDR prefixes.",
													Computed:            true,
													Attributes: map[string]dsschema.Attribute{
														"cidrs": dsschema.ListAttribute{
															ElementType:         types.StringType,
															MarkdownDescription: "CIDR address blocks.",
															Computed:            true,
														},
													},
												},
											},
										},
										"rejected_reason": dsschema.StringAttribute{
											MarkdownDescription: "Human-readable reason if the endpoint was rejected.",
											Computed:            true,
										},
										"status": dsschema.StringAttribute{
											MarkdownDescription: "The current provisioning status of the endpoint. Possible values are `Pending`, `Accepted`, `Rejected`, `Created`, `Creating` and `Deleting`.",
											Computed:            true,
										},
										"subnet_id": dsschema.StringAttribute{
											MarkdownDescription: "The AWS subnet ID in which the endpoint resides.",
											Computed:            true,
										},
										"vpc_id": dsschema.StringAttribute{
											MarkdownDescription: "The AWS VPC ID in which the endpoint resides.",
											Computed:            true,
										},
										"zone_id": dsschema.StringAttribute{
											MarkdownDescription: "The availability zone ID for this endpoint.",
											Computed:            true,
										},
									},
								},
							},
							"firewall_id": dsschema.StringAttribute{
								MarkdownDescription: "The unique identifier of the firewall.",
								Computed:            true,
							},
							"gwlb": dsschema.SingleNestedAttribute{
								MarkdownDescription: "Gateway Load Balancer configuration.",
								Computed:            true,
								Attributes: map[string]dsschema.Attribute{
									"deregistration_delay": dsschema.Int64Attribute{
										MarkdownDescription: "Deregistration delay in seconds.",
										Computed:            true,
									},
									"rejected_flow_count_alert_threshold": dsschema.Int64Attribute{
										MarkdownDescription: "Alert threshold for the total rejected flow count.",
										Computed:            true,
									},
									"rejected_flow_count_tcp_alert_threshold": dsschema.Int64Attribute{
										MarkdownDescription: "Alert threshold for the TCP rejected flow count.",
										Computed:            true,
									},
									"session_rebalance_enabled": dsschema.BoolAttribute{
										MarkdownDescription: "Whether session rebalancing is enabled.",
										Computed:            true,
									},
									"tcp_idle_timeout": dsschema.Int64Attribute{
										MarkdownDescription: "TCP idle timeout in seconds.",
										Computed:            true,
									},
								},
							},
							"ipv6": dsschema.SingleNestedAttribute{
								MarkdownDescription: "IPv6 configuration for the firewall.",
								Computed:            true,
								Attributes: map[string]dsschema.Attribute{
									"enabled": dsschema.BoolAttribute{
										MarkdownDescription: "Whether IPv6 is enabled on this firewall.",
										Computed:            true,
									},
								},
							},
							"link_id": dsschema.StringAttribute{
								MarkdownDescription: "The link ID associated with the firewall.",
								Computed:            true,
							},
							"link_status": dsschema.StringAttribute{
								MarkdownDescription: "The current link status of the firewall.",
								Computed:            true,
							},
							"log_config": dsschema.SingleNestedAttribute{
								MarkdownDescription: "Log destination configuration for the firewall.",
								Computed:            true,
								Attributes: map[string]dsschema.Attribute{
									"account_id": dsschema.StringAttribute{
										MarkdownDescription: "The AWS account ID where logs are sent.",
										Computed:            true,
									},
									"log_destination": dsschema.StringAttribute{
										MarkdownDescription: "The ARN or name of the log destination resource.",
										Computed:            true,
									},
									"log_destination_region": dsschema.StringAttribute{
										MarkdownDescription: "The AWS region of the log destination.",
										Computed:            true,
									},
									"log_destination_type": dsschema.StringAttribute{
										MarkdownDescription: "The type of log destination. Possible values are `S3`, `CloudWatchLogs` and `KinesisDataFirehose`.",
										Computed:            true,
									},
									"log_type": dsschema.ListAttribute{
										ElementType:         types.StringType,
										MarkdownDescription: "The types of logs to send to this destination. Possible values are `TRAFFIC`, `DECRYPTION`, `THREAT` and `USER_ID`.",
										Computed:            true,
									},
									"role_type": dsschema.StringAttribute{
										MarkdownDescription: "The IAM role type used to publish logs. Possible values are `ResourceBased` and `IamBased`.",
										Computed:            true,
									},
								},
							},
							"notifications": dsschema.ListNestedAttribute{
								MarkdownDescription: "Notifications associated with this firewall.",
								Computed:            true,
								NestedObject: dsschema.NestedAttributeObject{
									Attributes: map[string]dsschema.Attribute{
										"expiry_time": dsschema.Int64Attribute{
											MarkdownDescription: "Unix timestamp when this notification expires.",
											Computed:            true,
										},
										"id": dsschema.StringAttribute{
											MarkdownDescription: "Unique identifier of the notification.",
											Computed:            true,
										},
										"link": dsschema.SingleNestedAttribute{
											MarkdownDescription: "A hyperlink associated with a notification.",
											Computed:            true,
											Attributes: map[string]dsschema.Attribute{
												"target": dsschema.StringAttribute{
													MarkdownDescription: "The URL target of the link.",
													Computed:            true,
												},
												"text": dsschema.StringAttribute{
													MarkdownDescription: "The display text of the link.",
													Computed:            true,
												},
											},
										},
										"msg": dsschema.StringAttribute{
											MarkdownDescription: "The notification message text.",
											Computed:            true,
										},
									},
								},
							},
							"private_access": dsschema.SingleNestedAttribute{
								MarkdownDescription: "Private access configuration for the firewall.",
								Computed:            true,
								Attributes: map[string]dsschema.Attribute{
									"resource_id": dsschema.StringAttribute{
										MarkdownDescription: "The resource ID of the private access endpoint.",
										Computed:            true,
									},
									"type": dsschema.StringAttribute{
										MarkdownDescription: "The type of private access resource. Possible values are `NetworkLoadBalancer`.",
										Computed:            true,
									},
								},
							},
							"region": dsschema.StringAttribute{
								MarkdownDescription: "The AWS region where the firewall is deployed.",
								Computed:            true,
							},
							"software_version": dsschema.StringAttribute{
								MarkdownDescription: "The PAN-OS software version currently running on the firewall.",
								Computed:            true,
							},
							"tags": dsschema.ListNestedAttribute{
								MarkdownDescription: "AWS resource tags applied to the firewall.",
								Computed:            true,
								NestedObject: dsschema.NestedAttributeObject{
									Attributes: map[string]dsschema.Attribute{
										"key": dsschema.StringAttribute{
											MarkdownDescription: "The tag key.",
											Computed:            true,
										},
										"value": dsschema.StringAttribute{
											MarkdownDescription: "The tag value.",
											Computed:            true,
										},
									},
								},
							},
							"tier": dsschema.StringAttribute{
								MarkdownDescription: "The service tier of the firewall (base, standard, or premium).",
								Computed:            true,
							},
							"update_token": dsschema.StringAttribute{
								MarkdownDescription: "Optimistic locking token; must be supplied on update operations.",
								Computed:            true,
							},
							"user_id": dsschema.SingleNestedAttribute{
								MarkdownDescription: "User-ID configuration for user-based policy enforcement.",
								Computed:            true,
								Attributes: map[string]dsschema.Attribute{
									"agent_name": dsschema.StringAttribute{
										MarkdownDescription: "The name of the User-ID agent.",
										Computed:            true,
									},
									"cfturl_link": dsschema.StringAttribute{
										MarkdownDescription: "CloudFormation template URL link for User-ID setup.",
										Computed:            true,
									},
									"collector_name": dsschema.StringAttribute{
										MarkdownDescription: "The name of the log collector used by User-ID.",
										Computed:            true,
									},
									"custom_include_exclude_network": dsschema.ListNestedAttribute{
										MarkdownDescription: "Custom include/exclude network definitions for User-ID.",
										Computed:            true,
										NestedObject: dsschema.NestedAttributeObject{
											Attributes: map[string]dsschema.Attribute{
												"discovery_include": dsschema.BoolAttribute{
													MarkdownDescription: "Whether this network is included in User-ID discovery.",
													Computed:            true,
												},
												"enabled": dsschema.BoolAttribute{
													MarkdownDescription: "Whether this network entry is active.",
													Computed:            true,
												},
												"name": dsschema.StringAttribute{
													MarkdownDescription: "Name of the network entry.",
													Computed:            true,
												},
												"network_address": dsschema.StringAttribute{
													MarkdownDescription: "The network address (CIDR) for this entry.",
													Computed:            true,
												},
											},
										},
									},
									"enabled": dsschema.BoolAttribute{
										MarkdownDescription: "Whether User-ID is enabled.",
										Computed:            true,
									},
									"endpoint_dns": dsschema.StringAttribute{
										MarkdownDescription: "DNS name of the User-ID endpoint.",
										Computed:            true,
									},
									"port": dsschema.Int64Attribute{
										MarkdownDescription: "Port used by the User-ID agent.",
										Computed:            true,
									},
									"secret_key_arn": dsschema.StringAttribute{
										MarkdownDescription: "ARN of the AWS Secrets Manager secret for User-ID credentials.",
										Computed:            true,
									},
									"user_id_status": dsschema.StringAttribute{
										MarkdownDescription: "The operational status of the User-ID feature. Possible values are `EndpointPending`, `EndpointFail`, `CommitFail`, `Enabled`, `Disabled`, `DisablePending`, `UpdatePending` and `UpdateFail`.",
										Computed:            true,
									},
								},
							},
						},
					},
					"status": dsschema.SingleNestedAttribute{
						MarkdownDescription: "Operational status of the firewall.",
						Computed:            true,
						Attributes: map[string]dsschema.Attribute{
							"device_rule_stack_commit_info": dsschema.SingleNestedAttribute{
								MarkdownDescription: "Commit information for a rulestack.",
								Computed:            true,
								Attributes: map[string]dsschema.Attribute{
									"commit_messages": dsschema.ListAttribute{
										ElementType:         types.StringType,
										MarkdownDescription: "Messages produced during the commit operation.",
										Computed:            true,
									},
									"commit_ts": dsschema.StringAttribute{
										MarkdownDescription: "Timestamp of the last commit operation.",
										Computed:            true,
									},
								},
							},
							"device_rule_stack_commit_state": dsschema.StringAttribute{
								MarkdownDescription: "The commit state of the device rulestack. Possible values are `Submitted` and `None`.",
								Computed:            true,
							},
							"device_rule_stack_commit_status": dsschema.StringAttribute{
								MarkdownDescription: "The commit status of the device rulestack. Possible values are `Running`, `Uncommitted`, `Precommit`, `PrecommitDone`, `PrevalidateDone`, `Committing`, `Failed` and `Success`.",
								Computed:            true,
							},
							"failure_reason": dsschema.StringAttribute{
								MarkdownDescription: "Human-readable reason for a firewall failure, if any.",
								Computed:            true,
							},
							"firewall_status": dsschema.StringAttribute{
								MarkdownDescription: "The operational lifecycle status of the firewall. Possible values are `CREATING`, `UPDATING`, `DELETING`, `CREATE_COMPLETE`, `UPDATE_COMPLETE`, `CREATE_FAIL`, `UPDATE_FAIL`, `DELETE_FAIL` and `DELETE_COMPLETE`.",
								Computed:            true,
							},
							"global_rule_stack_commit_info": dsschema.SingleNestedAttribute{
								MarkdownDescription: "Commit information for a rulestack.",
								Computed:            true,
								Attributes: map[string]dsschema.Attribute{
									"commit_messages": dsschema.ListAttribute{
										ElementType:         types.StringType,
										MarkdownDescription: "Messages produced during the commit operation.",
										Computed:            true,
									},
									"commit_ts": dsschema.StringAttribute{
										MarkdownDescription: "Timestamp of the last commit operation.",
										Computed:            true,
									},
								},
							},
							"global_rule_stack_status": dsschema.StringAttribute{
								MarkdownDescription: "The commit status of the global rulestack. Possible values are `Running`, `Uncommitted`, `Precommit`, `PrecommitDone`, `PrevalidateDone`, `Committing`, `Failed` and `Success`.",
								Computed:            true,
							},
							"public_i_ps": dsschema.ListNestedAttribute{
								MarkdownDescription: "Public IP addresses assigned to the firewall.",
								Computed:            true,
								NestedObject: dsschema.NestedAttributeObject{
									Attributes: map[string]dsschema.Attribute{
										"ip_address": dsschema.StringAttribute{
											MarkdownDescription: "The public IP address.",
											Computed:            true,
										},
										"ip_source": dsschema.StringAttribute{
											MarkdownDescription: "The source of the IP address allocation.",
											Computed:            true,
										},
										"ip_status": dsschema.StringAttribute{
											MarkdownDescription: "The current status of this public IP.",
											Computed:            true,
										},
									},
								},
							},
							"rule_stack_commit_info": dsschema.SingleNestedAttribute{
								MarkdownDescription: "Commit information for a rulestack.",
								Computed:            true,
								Attributes: map[string]dsschema.Attribute{
									"commit_messages": dsschema.ListAttribute{
										ElementType:         types.StringType,
										MarkdownDescription: "Messages produced during the commit operation.",
										Computed:            true,
									},
									"commit_ts": dsschema.StringAttribute{
										MarkdownDescription: "Timestamp of the last commit operation.",
										Computed:            true,
									},
								},
							},
							"rule_stack_status": dsschema.StringAttribute{
								MarkdownDescription: "The commit status of the local rulestack. Possible values are `Running`, `Uncommitted`, `Precommit`, `PrecommitDone`, `PrevalidateDone`, `Committing`, `Failed` and `Success`.",
								Computed:            true,
							},
							"scm_assoc_status": dsschema.BoolAttribute{
								MarkdownDescription: "Whether the firewall is associated with Strata Cloud Manager.",
								Computed:            true,
							},
						},
					},
				},
			},
		},
	},
}

// NgfirewallsListModel represents the data model for a list data source.
type NgfirewallsListModel struct {
	Tfid                types.String             `tfsdk:"tfid"`
	Total               types.Int64              `tfsdk:"total"`
	FirewallsDescribe   []NgfirewallResponseBody `tfsdk:"firewalls_describe"`
	Region              basetypes.StringValue    `tfsdk:"region"`
	Globalrulestackname basetypes.StringValue    `tfsdk:"globalrulestackname"`
}

// NgfirewallsListDataSourceSchema defines the schema for a list data source.
var NgfirewallsListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid":  dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"total": dsschema.Int64Attribute{Description: "The total number of items.", Computed: true},
		"firewalls_describe": dsschema.ListNestedAttribute{
			Description: "List of firewalls_describe items.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: NgfirewallResponseBodyDataSourceSchema.Attributes,
			},
		},
		"region": dsschema.StringAttribute{
			Description: "The AWS region where the Cloud NGFW firewall is deployed",
			Required:    true,
		},
		"globalrulestackname": dsschema.StringAttribute{
			Description: "Name of the global rulestack to filter by.",
			Optional:    true,
		},
	},
}
