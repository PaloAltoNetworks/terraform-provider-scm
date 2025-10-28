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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Package: objects
// This file contains models for the objects SDK package

// HipObjects represents the Terraform model for HipObjects
type HipObjects struct {
	Tfid               types.String          `tfsdk:"tfid"`
	AntiMalware        basetypes.ObjectValue `tfsdk:"anti_malware"`
	Certificate        basetypes.ObjectValue `tfsdk:"certificate"`
	CustomChecks       basetypes.ObjectValue `tfsdk:"custom_checks"`
	DataLossPrevention basetypes.ObjectValue `tfsdk:"data_loss_prevention"`
	Description        basetypes.StringValue `tfsdk:"description"`
	Device             basetypes.StringValue `tfsdk:"device"`
	DiskBackup         basetypes.ObjectValue `tfsdk:"disk_backup"`
	DiskEncryption     basetypes.ObjectValue `tfsdk:"disk_encryption"`
	Firewall           basetypes.ObjectValue `tfsdk:"firewall"`
	Folder             basetypes.StringValue `tfsdk:"folder"`
	HostInfo           basetypes.ObjectValue `tfsdk:"host_info"`
	Id                 basetypes.StringValue `tfsdk:"id"`
	MobileDevice       basetypes.ObjectValue `tfsdk:"mobile_device"`
	Name               basetypes.StringValue `tfsdk:"name"`
	NetworkInfo        basetypes.ObjectValue `tfsdk:"network_info"`
	PatchManagement    basetypes.ObjectValue `tfsdk:"patch_management"`
	Snippet            basetypes.StringValue `tfsdk:"snippet"`
}

// HipObjectsAntiMalware represents a nested structure within the HipObjects model
type HipObjectsAntiMalware struct {
	Criteria      basetypes.ObjectValue `tfsdk:"criteria"`
	ExcludeVendor basetypes.BoolValue   `tfsdk:"exclude_vendor"`
	Vendor        basetypes.ListValue   `tfsdk:"vendor"`
}

// HipObjectsAntiMalwareCriteria represents a nested structure within the HipObjects model
type HipObjectsAntiMalwareCriteria struct {
	IsInstalled        basetypes.BoolValue   `tfsdk:"is_installed"`
	LastScanTime       basetypes.ObjectValue `tfsdk:"last_scan_time"`
	ProductVersion     basetypes.ObjectValue `tfsdk:"product_version"`
	RealTimeProtection basetypes.StringValue `tfsdk:"real_time_protection"`
	VirdefVersion      basetypes.ObjectValue `tfsdk:"virdef_version"`
}

// HipObjectsAntiMalwareCriteriaLastScanTime represents a nested structure within the HipObjects model
type HipObjectsAntiMalwareCriteriaLastScanTime struct {
	NotAvailable basetypes.ObjectValue `tfsdk:"not_available"`
	NotWithin    basetypes.ObjectValue `tfsdk:"not_within"`
	Within       basetypes.ObjectValue `tfsdk:"within"`
}

// HipObjectsAntiMalwareCriteriaLastScanTimeNotWithin represents a nested structure within the HipObjects model
type HipObjectsAntiMalwareCriteriaLastScanTimeNotWithin struct {
	Days  basetypes.Int64Value `tfsdk:"days"`
	Hours basetypes.Int64Value `tfsdk:"hours"`
}

// HipObjectsAntiMalwareCriteriaProductVersion represents a nested structure within the HipObjects model
type HipObjectsAntiMalwareCriteriaProductVersion struct {
	Contains     basetypes.StringValue `tfsdk:"contains"`
	GreaterEqual basetypes.StringValue `tfsdk:"greater_equal"`
	GreaterThan  basetypes.StringValue `tfsdk:"greater_than"`
	Is           basetypes.StringValue `tfsdk:"is"`
	IsNot        basetypes.StringValue `tfsdk:"is_not"`
	LessEqual    basetypes.StringValue `tfsdk:"less_equal"`
	LessThan     basetypes.StringValue `tfsdk:"less_than"`
	NotWithin    basetypes.ObjectValue `tfsdk:"not_within"`
	Within       basetypes.ObjectValue `tfsdk:"within"`
}

// HipObjectsAntiMalwareCriteriaProductVersionNotWithin represents a nested structure within the HipObjects model
type HipObjectsAntiMalwareCriteriaProductVersionNotWithin struct {
	Versions basetypes.Int64Value `tfsdk:"versions"`
}

// HipObjectsAntiMalwareCriteriaVirdefVersion represents a nested structure within the HipObjects model
type HipObjectsAntiMalwareCriteriaVirdefVersion struct {
	NotWithin basetypes.ObjectValue `tfsdk:"not_within"`
	Within    basetypes.ObjectValue `tfsdk:"within"`
}

// HipObjectsAntiMalwareCriteriaVirdefVersionNotWithin represents a nested structure within the HipObjects model
type HipObjectsAntiMalwareCriteriaVirdefVersionNotWithin struct {
	Days     basetypes.Int64Value `tfsdk:"days"`
	Versions basetypes.Int64Value `tfsdk:"versions"`
}

// HipObjectsAntiMalwareVendorInner represents a nested structure within the HipObjects model
type HipObjectsAntiMalwareVendorInner struct {
	Name    basetypes.StringValue `tfsdk:"name"`
	Product basetypes.ListValue   `tfsdk:"product"`
}

// HipObjectsCertificate represents a nested structure within the HipObjects model
type HipObjectsCertificate struct {
	Criteria basetypes.ObjectValue `tfsdk:"criteria"`
}

// HipObjectsCertificateCriteria represents a nested structure within the HipObjects model
type HipObjectsCertificateCriteria struct {
	CertificateAttributes basetypes.ListValue   `tfsdk:"certificate_attributes"`
	CertificateProfile    basetypes.StringValue `tfsdk:"certificate_profile"`
}

// HipObjectsCertificateCriteriaCertificateAttributesInner represents a nested structure within the HipObjects model
type HipObjectsCertificateCriteriaCertificateAttributesInner struct {
	Name  basetypes.StringValue `tfsdk:"name"`
	Value basetypes.StringValue `tfsdk:"value"`
}

// HipObjectsCustomChecks represents a nested structure within the HipObjects model
type HipObjectsCustomChecks struct {
	Criteria basetypes.ObjectValue `tfsdk:"criteria"`
}

// HipObjectsCustomChecksCriteria represents a nested structure within the HipObjects model
type HipObjectsCustomChecksCriteria struct {
	Plist       basetypes.ListValue `tfsdk:"plist"`
	ProcessList basetypes.ListValue `tfsdk:"process_list"`
	RegistryKey basetypes.ListValue `tfsdk:"registry_key"`
}

// HipObjectsCustomChecksCriteriaPlistInner represents a nested structure within the HipObjects model
type HipObjectsCustomChecksCriteriaPlistInner struct {
	Key    basetypes.ListValue   `tfsdk:"key"`
	Name   basetypes.StringValue `tfsdk:"name"`
	Negate basetypes.BoolValue   `tfsdk:"negate"`
}

// HipObjectsCustomChecksCriteriaPlistInnerKeyInner represents a nested structure within the HipObjects model
type HipObjectsCustomChecksCriteriaPlistInnerKeyInner struct {
	Name   basetypes.StringValue `tfsdk:"name"`
	Negate basetypes.BoolValue   `tfsdk:"negate"`
	Value  basetypes.StringValue `tfsdk:"value"`
}

// HipObjectsCustomChecksCriteriaProcessListInner represents a nested structure within the HipObjects model
type HipObjectsCustomChecksCriteriaProcessListInner struct {
	Name    basetypes.StringValue `tfsdk:"name"`
	Running basetypes.BoolValue   `tfsdk:"running"`
}

// HipObjectsCustomChecksCriteriaRegistryKeyInner represents a nested structure within the HipObjects model
type HipObjectsCustomChecksCriteriaRegistryKeyInner struct {
	DefaultValueData basetypes.StringValue `tfsdk:"default_value_data"`
	Name             basetypes.StringValue `tfsdk:"name"`
	Negate           basetypes.BoolValue   `tfsdk:"negate"`
	RegistryValue    basetypes.ListValue   `tfsdk:"registry_value"`
}

// HipObjectsCustomChecksCriteriaRegistryKeyInnerRegistryValueInner represents a nested structure within the HipObjects model
type HipObjectsCustomChecksCriteriaRegistryKeyInnerRegistryValueInner struct {
	Name      basetypes.StringValue `tfsdk:"name"`
	Negate    basetypes.BoolValue   `tfsdk:"negate"`
	ValueData basetypes.StringValue `tfsdk:"value_data"`
}

// HipObjectsDataLossPrevention represents a nested structure within the HipObjects model
type HipObjectsDataLossPrevention struct {
	Criteria      basetypes.ObjectValue `tfsdk:"criteria"`
	ExcludeVendor basetypes.BoolValue   `tfsdk:"exclude_vendor"`
	Vendor        basetypes.ListValue   `tfsdk:"vendor"`
}

// HipObjectsDataLossPreventionCriteria represents a nested structure within the HipObjects model
type HipObjectsDataLossPreventionCriteria struct {
	IsEnabled   basetypes.StringValue `tfsdk:"is_enabled"`
	IsInstalled basetypes.BoolValue   `tfsdk:"is_installed"`
}

// HipObjectsDataLossPreventionVendorInner represents a nested structure within the HipObjects model
type HipObjectsDataLossPreventionVendorInner struct {
	Name    basetypes.StringValue `tfsdk:"name"`
	Product basetypes.ListValue   `tfsdk:"product"`
}

// HipObjectsDiskBackup represents a nested structure within the HipObjects model
type HipObjectsDiskBackup struct {
	Criteria      basetypes.ObjectValue `tfsdk:"criteria"`
	ExcludeVendor basetypes.BoolValue   `tfsdk:"exclude_vendor"`
	Vendor        basetypes.ListValue   `tfsdk:"vendor"`
}

// HipObjectsDiskBackupCriteria represents a nested structure within the HipObjects model
type HipObjectsDiskBackupCriteria struct {
	IsInstalled    basetypes.BoolValue   `tfsdk:"is_installed"`
	LastBackupTime basetypes.ObjectValue `tfsdk:"last_backup_time"`
}

// HipObjectsDiskEncryption represents a nested structure within the HipObjects model
type HipObjectsDiskEncryption struct {
	Criteria      basetypes.ObjectValue `tfsdk:"criteria"`
	ExcludeVendor basetypes.BoolValue   `tfsdk:"exclude_vendor"`
	Vendor        basetypes.ListValue   `tfsdk:"vendor"`
}

// HipObjectsDiskEncryptionCriteria represents a nested structure within the HipObjects model
type HipObjectsDiskEncryptionCriteria struct {
	EncryptedLocations basetypes.ListValue `tfsdk:"encrypted_locations"`
	IsInstalled        basetypes.BoolValue `tfsdk:"is_installed"`
}

// HipObjectsDiskEncryptionCriteriaEncryptedLocationsInner represents a nested structure within the HipObjects model
type HipObjectsDiskEncryptionCriteriaEncryptedLocationsInner struct {
	EncryptionState basetypes.ObjectValue `tfsdk:"encryption_state"`
	Name            basetypes.StringValue `tfsdk:"name"`
}

// HipObjectsDiskEncryptionCriteriaEncryptedLocationsInnerEncryptionState represents a nested structure within the HipObjects model
type HipObjectsDiskEncryptionCriteriaEncryptedLocationsInnerEncryptionState struct {
	Is    basetypes.StringValue `tfsdk:"is"`
	IsNot basetypes.StringValue `tfsdk:"is_not"`
}

// HipObjectsFirewall represents a nested structure within the HipObjects model
type HipObjectsFirewall struct {
	Criteria      basetypes.ObjectValue `tfsdk:"criteria"`
	ExcludeVendor basetypes.BoolValue   `tfsdk:"exclude_vendor"`
	Vendor        basetypes.ListValue   `tfsdk:"vendor"`
}

// HipObjectsHostInfo represents a nested structure within the HipObjects model
type HipObjectsHostInfo struct {
	Criteria basetypes.ObjectValue `tfsdk:"criteria"`
}

// HipObjectsHostInfoCriteria represents a nested structure within the HipObjects model
type HipObjectsHostInfoCriteria struct {
	ClientVersion basetypes.ObjectValue `tfsdk:"client_version"`
	Domain        basetypes.ObjectValue `tfsdk:"domain"`
	HostId        basetypes.ObjectValue `tfsdk:"host_id"`
	HostName      basetypes.ObjectValue `tfsdk:"host_name"`
	Managed       basetypes.BoolValue   `tfsdk:"managed"`
	Os            basetypes.ObjectValue `tfsdk:"os"`
	SerialNumber  basetypes.ObjectValue `tfsdk:"serial_number"`
}

// HipObjectsHostInfoCriteriaClientVersion represents a nested structure within the HipObjects model
type HipObjectsHostInfoCriteriaClientVersion struct {
	Contains basetypes.StringValue `tfsdk:"contains"`
	Is       basetypes.StringValue `tfsdk:"is"`
	IsNot    basetypes.StringValue `tfsdk:"is_not"`
}

// HipObjectsHostInfoCriteriaOs represents a nested structure within the HipObjects model
type HipObjectsHostInfoCriteriaOs struct {
	Contains basetypes.ObjectValue `tfsdk:"contains"`
}

// HipObjectsHostInfoCriteriaOsContains represents a nested structure within the HipObjects model
type HipObjectsHostInfoCriteriaOsContains struct {
	Apple     basetypes.StringValue `tfsdk:"apple"`
	Google    basetypes.StringValue `tfsdk:"google"`
	Linux     basetypes.StringValue `tfsdk:"linux"`
	Microsoft basetypes.StringValue `tfsdk:"microsoft"`
	Other     basetypes.StringValue `tfsdk:"other"`
}

// HipObjectsMobileDevice represents a nested structure within the HipObjects model
type HipObjectsMobileDevice struct {
	Criteria basetypes.ObjectValue `tfsdk:"criteria"`
}

// HipObjectsMobileDeviceCriteria represents a nested structure within the HipObjects model
type HipObjectsMobileDeviceCriteria struct {
	Applications    basetypes.ObjectValue `tfsdk:"applications"`
	DiskEncrypted   basetypes.BoolValue   `tfsdk:"disk_encrypted"`
	Imei            basetypes.ObjectValue `tfsdk:"imei"`
	Jailbroken      basetypes.BoolValue   `tfsdk:"jailbroken"`
	LastCheckinTime basetypes.ObjectValue `tfsdk:"last_checkin_time"`
	Model           basetypes.ObjectValue `tfsdk:"model"`
	PasscodeSet     basetypes.BoolValue   `tfsdk:"passcode_set"`
	PhoneNumber     basetypes.ObjectValue `tfsdk:"phone_number"`
	Tag             basetypes.ObjectValue `tfsdk:"tag"`
}

// HipObjectsMobileDeviceCriteriaApplications represents a nested structure within the HipObjects model
type HipObjectsMobileDeviceCriteriaApplications struct {
	HasMalware      basetypes.ObjectValue `tfsdk:"has_malware"`
	HasUnmanagedApp basetypes.BoolValue   `tfsdk:"has_unmanaged_app"`
	Includes        basetypes.ListValue   `tfsdk:"includes"`
}

// HipObjectsMobileDeviceCriteriaApplicationsHasMalware represents a nested structure within the HipObjects model
type HipObjectsMobileDeviceCriteriaApplicationsHasMalware struct {
	No  basetypes.ObjectValue `tfsdk:"no"`
	Yes basetypes.ObjectValue `tfsdk:"yes"`
}

// HipObjectsMobileDeviceCriteriaApplicationsHasMalwareYes represents a nested structure within the HipObjects model
type HipObjectsMobileDeviceCriteriaApplicationsHasMalwareYes struct {
	Excludes basetypes.ListValue `tfsdk:"excludes"`
}

// HipObjectsMobileDeviceCriteriaApplicationsHasMalwareYesExcludesInner represents a nested structure within the HipObjects model
type HipObjectsMobileDeviceCriteriaApplicationsHasMalwareYesExcludesInner struct {
	Hash    basetypes.StringValue `tfsdk:"hash"`
	Name    basetypes.StringValue `tfsdk:"name"`
	Package basetypes.StringValue `tfsdk:"package"`
}

// HipObjectsMobileDeviceCriteriaLastCheckinTime represents a nested structure within the HipObjects model
type HipObjectsMobileDeviceCriteriaLastCheckinTime struct {
	NotWithin basetypes.ObjectValue `tfsdk:"not_within"`
	Within    basetypes.ObjectValue `tfsdk:"within"`
}

// HipObjectsMobileDeviceCriteriaLastCheckinTimeNotWithin represents a nested structure within the HipObjects model
type HipObjectsMobileDeviceCriteriaLastCheckinTimeNotWithin struct {
	Days basetypes.Int64Value `tfsdk:"days"`
}

// HipObjectsNetworkInfo represents a nested structure within the HipObjects model
type HipObjectsNetworkInfo struct {
	Criteria basetypes.ObjectValue `tfsdk:"criteria"`
}

// HipObjectsNetworkInfoCriteria represents a nested structure within the HipObjects model
type HipObjectsNetworkInfoCriteria struct {
	Network basetypes.ObjectValue `tfsdk:"network"`
}

// HipObjectsNetworkInfoCriteriaNetwork represents a nested structure within the HipObjects model
type HipObjectsNetworkInfoCriteriaNetwork struct {
	Is    basetypes.ObjectValue `tfsdk:"is"`
	IsNot basetypes.ObjectValue `tfsdk:"is_not"`
}

// HipObjectsNetworkInfoCriteriaNetworkIs represents a nested structure within the HipObjects model
type HipObjectsNetworkInfoCriteriaNetworkIs struct {
	Mobile  basetypes.ObjectValue `tfsdk:"mobile"`
	Unknown basetypes.ObjectValue `tfsdk:"unknown"`
	Wifi    basetypes.ObjectValue `tfsdk:"wifi"`
}

// HipObjectsNetworkInfoCriteriaNetworkIsMobile represents a nested structure within the HipObjects model
type HipObjectsNetworkInfoCriteriaNetworkIsMobile struct {
	Carrier basetypes.StringValue `tfsdk:"carrier"`
}

// HipObjectsNetworkInfoCriteriaNetworkIsWifi represents a nested structure within the HipObjects model
type HipObjectsNetworkInfoCriteriaNetworkIsWifi struct {
	Ssid basetypes.StringValue `tfsdk:"ssid"`
}

// HipObjectsNetworkInfoCriteriaNetworkIsNot represents a nested structure within the HipObjects model
type HipObjectsNetworkInfoCriteriaNetworkIsNot struct {
	Ethernet basetypes.ObjectValue `tfsdk:"ethernet"`
	Mobile   basetypes.ObjectValue `tfsdk:"mobile"`
	Unknown  basetypes.ObjectValue `tfsdk:"unknown"`
	Wifi     basetypes.ObjectValue `tfsdk:"wifi"`
}

// HipObjectsPatchManagement represents a nested structure within the HipObjects model
type HipObjectsPatchManagement struct {
	Criteria      basetypes.ObjectValue `tfsdk:"criteria"`
	ExcludeVendor basetypes.BoolValue   `tfsdk:"exclude_vendor"`
	Vendor        basetypes.ListValue   `tfsdk:"vendor"`
}

// HipObjectsPatchManagementCriteria represents a nested structure within the HipObjects model
type HipObjectsPatchManagementCriteria struct {
	IsEnabled      basetypes.StringValue `tfsdk:"is_enabled"`
	IsInstalled    basetypes.BoolValue   `tfsdk:"is_installed"`
	MissingPatches basetypes.ObjectValue `tfsdk:"missing_patches"`
}

// HipObjectsPatchManagementCriteriaMissingPatches represents a nested structure within the HipObjects model
type HipObjectsPatchManagementCriteriaMissingPatches struct {
	Check    basetypes.StringValue `tfsdk:"check"`
	Patches  basetypes.ListValue   `tfsdk:"patches"`
	Severity basetypes.ObjectValue `tfsdk:"severity"`
}

// HipObjectsPatchManagementCriteriaMissingPatchesSeverity represents a nested structure within the HipObjects model
type HipObjectsPatchManagementCriteriaMissingPatchesSeverity struct {
	GreaterEqual basetypes.Int64Value `tfsdk:"greater_equal"`
	GreaterThan  basetypes.Int64Value `tfsdk:"greater_than"`
	Is           basetypes.Int64Value `tfsdk:"is"`
	IsNot        basetypes.Int64Value `tfsdk:"is_not"`
	LessEqual    basetypes.Int64Value `tfsdk:"less_equal"`
	LessThan     basetypes.Int64Value `tfsdk:"less_than"`
}

// AttrTypes defines the attribute types for the HipObjects model.
func (o HipObjects) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid": basetypes.StringType{},
		"anti_malware": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"criteria": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"is_installed": basetypes.BoolType{},
						"last_scan_time": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"not_available": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"not_within": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"days":  basetypes.Int64Type{},
										"hours": basetypes.Int64Type{},
									},
								},
								"within": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"days":  basetypes.Int64Type{},
										"hours": basetypes.Int64Type{},
									},
								},
							},
						},
						"product_version": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"contains":      basetypes.StringType{},
								"greater_equal": basetypes.StringType{},
								"greater_than":  basetypes.StringType{},
								"is":            basetypes.StringType{},
								"is_not":        basetypes.StringType{},
								"less_equal":    basetypes.StringType{},
								"less_than":     basetypes.StringType{},
								"not_within": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"versions": basetypes.Int64Type{},
									},
								},
								"within": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"versions": basetypes.Int64Type{},
									},
								},
							},
						},
						"real_time_protection": basetypes.StringType{},
						"virdef_version": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"not_within": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"days":     basetypes.Int64Type{},
										"versions": basetypes.Int64Type{},
									},
								},
								"within": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"days":     basetypes.Int64Type{},
										"versions": basetypes.Int64Type{},
									},
								},
							},
						},
					},
				},
				"exclude_vendor": basetypes.BoolType{},
				"vendor": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"name":    basetypes.StringType{},
						"product": basetypes.ListType{ElemType: basetypes.StringType{}},
					},
				}},
			},
		},
		"certificate": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"criteria": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"certificate_attributes": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"name":  basetypes.StringType{},
								"value": basetypes.StringType{},
							},
						}},
						"certificate_profile": basetypes.StringType{},
					},
				},
			},
		},
		"custom_checks": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"criteria": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"plist": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"key": basetypes.ListType{ElemType: basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"name":   basetypes.StringType{},
										"negate": basetypes.BoolType{},
										"value":  basetypes.StringType{},
									},
								}},
								"name":   basetypes.StringType{},
								"negate": basetypes.BoolType{},
							},
						}},
						"process_list": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"name":    basetypes.StringType{},
								"running": basetypes.BoolType{},
							},
						}},
						"registry_key": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"default_value_data": basetypes.StringType{},
								"name":               basetypes.StringType{},
								"negate":             basetypes.BoolType{},
								"registry_value": basetypes.ListType{ElemType: basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"name":       basetypes.StringType{},
										"negate":     basetypes.BoolType{},
										"value_data": basetypes.StringType{},
									},
								}},
							},
						}},
					},
				},
			},
		},
		"data_loss_prevention": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"criteria": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"is_enabled":   basetypes.StringType{},
						"is_installed": basetypes.BoolType{},
					},
				},
				"exclude_vendor": basetypes.BoolType{},
				"vendor": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"name":    basetypes.StringType{},
						"product": basetypes.ListType{ElemType: basetypes.StringType{}},
					},
				}},
			},
		},
		"description": basetypes.StringType{},
		"device":      basetypes.StringType{},
		"disk_backup": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"criteria": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"is_installed": basetypes.BoolType{},
						"last_backup_time": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"not_available": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"not_within": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"days":  basetypes.Int64Type{},
										"hours": basetypes.Int64Type{},
									},
								},
								"within": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"days":  basetypes.Int64Type{},
										"hours": basetypes.Int64Type{},
									},
								},
							},
						},
					},
				},
				"exclude_vendor": basetypes.BoolType{},
				"vendor": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"name":    basetypes.StringType{},
						"product": basetypes.ListType{ElemType: basetypes.StringType{}},
					},
				}},
			},
		},
		"disk_encryption": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"criteria": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"encrypted_locations": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"encryption_state": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"is":     basetypes.StringType{},
										"is_not": basetypes.StringType{},
									},
								},
								"name": basetypes.StringType{},
							},
						}},
						"is_installed": basetypes.BoolType{},
					},
				},
				"exclude_vendor": basetypes.BoolType{},
				"vendor": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"name":    basetypes.StringType{},
						"product": basetypes.ListType{ElemType: basetypes.StringType{}},
					},
				}},
			},
		},
		"firewall": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"criteria": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"is_enabled":   basetypes.StringType{},
						"is_installed": basetypes.BoolType{},
					},
				},
				"exclude_vendor": basetypes.BoolType{},
				"vendor": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"name":    basetypes.StringType{},
						"product": basetypes.ListType{ElemType: basetypes.StringType{}},
					},
				}},
			},
		},
		"folder": basetypes.StringType{},
		"host_info": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"criteria": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"client_version": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"contains": basetypes.StringType{},
								"is":       basetypes.StringType{},
								"is_not":   basetypes.StringType{},
							},
						},
						"domain": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"contains": basetypes.StringType{},
								"is":       basetypes.StringType{},
								"is_not":   basetypes.StringType{},
							},
						},
						"host_id": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"contains": basetypes.StringType{},
								"is":       basetypes.StringType{},
								"is_not":   basetypes.StringType{},
							},
						},
						"host_name": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"contains": basetypes.StringType{},
								"is":       basetypes.StringType{},
								"is_not":   basetypes.StringType{},
							},
						},
						"managed": basetypes.BoolType{},
						"os": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"contains": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"apple":     basetypes.StringType{},
										"google":    basetypes.StringType{},
										"linux":     basetypes.StringType{},
										"microsoft": basetypes.StringType{},
										"other":     basetypes.StringType{},
									},
								},
							},
						},
						"serial_number": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"contains": basetypes.StringType{},
								"is":       basetypes.StringType{},
								"is_not":   basetypes.StringType{},
							},
						},
					},
				},
			},
		},
		"id": basetypes.StringType{},
		"mobile_device": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"criteria": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"applications": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"has_malware": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"no": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
										"yes": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"excludes": basetypes.ListType{ElemType: basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"hash":    basetypes.StringType{},
														"name":    basetypes.StringType{},
														"package": basetypes.StringType{},
													},
												}},
											},
										},
									},
								},
								"has_unmanaged_app": basetypes.BoolType{},
								"includes": basetypes.ListType{ElemType: basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"hash":    basetypes.StringType{},
										"name":    basetypes.StringType{},
										"package": basetypes.StringType{},
									},
								}},
							},
						},
						"disk_encrypted": basetypes.BoolType{},
						"imei": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"contains": basetypes.StringType{},
								"is":       basetypes.StringType{},
								"is_not":   basetypes.StringType{},
							},
						},
						"jailbroken": basetypes.BoolType{},
						"last_checkin_time": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"not_within": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"days": basetypes.Int64Type{},
									},
								},
								"within": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"days": basetypes.Int64Type{},
									},
								},
							},
						},
						"model": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"contains": basetypes.StringType{},
								"is":       basetypes.StringType{},
								"is_not":   basetypes.StringType{},
							},
						},
						"passcode_set": basetypes.BoolType{},
						"phone_number": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"contains": basetypes.StringType{},
								"is":       basetypes.StringType{},
								"is_not":   basetypes.StringType{},
							},
						},
						"tag": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"contains": basetypes.StringType{},
								"is":       basetypes.StringType{},
								"is_not":   basetypes.StringType{},
							},
						},
					},
				},
			},
		},
		"name": basetypes.StringType{},
		"network_info": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"criteria": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"network": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"is": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"mobile": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"carrier": basetypes.StringType{},
											},
										},
										"unknown": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
										"wifi": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"ssid": basetypes.StringType{},
											},
										},
									},
								},
								"is_not": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"ethernet": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
										"mobile": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"carrier": basetypes.StringType{},
											},
										},
										"unknown": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
										"wifi": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"ssid": basetypes.StringType{},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"patch_management": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"criteria": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"is_enabled":   basetypes.StringType{},
						"is_installed": basetypes.BoolType{},
						"missing_patches": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"check":   basetypes.StringType{},
								"patches": basetypes.ListType{ElemType: basetypes.StringType{}},
								"severity": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"greater_equal": basetypes.Int64Type{},
										"greater_than":  basetypes.Int64Type{},
										"is":            basetypes.Int64Type{},
										"is_not":        basetypes.Int64Type{},
										"less_equal":    basetypes.Int64Type{},
										"less_than":     basetypes.Int64Type{},
									},
								},
							},
						},
					},
				},
				"exclude_vendor": basetypes.BoolType{},
				"vendor": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"name":    basetypes.StringType{},
						"product": basetypes.ListType{ElemType: basetypes.StringType{}},
					},
				}},
			},
		},
		"snippet": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of HipObjects objects.
func (o HipObjects) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the HipObjectsAntiMalware model.
func (o HipObjectsAntiMalware) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"criteria": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"is_installed": basetypes.BoolType{},
				"last_scan_time": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"not_available": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"not_within": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"days":  basetypes.Int64Type{},
								"hours": basetypes.Int64Type{},
							},
						},
						"within": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"days":  basetypes.Int64Type{},
								"hours": basetypes.Int64Type{},
							},
						},
					},
				},
				"product_version": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"contains":      basetypes.StringType{},
						"greater_equal": basetypes.StringType{},
						"greater_than":  basetypes.StringType{},
						"is":            basetypes.StringType{},
						"is_not":        basetypes.StringType{},
						"less_equal":    basetypes.StringType{},
						"less_than":     basetypes.StringType{},
						"not_within": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"versions": basetypes.Int64Type{},
							},
						},
						"within": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"versions": basetypes.Int64Type{},
							},
						},
					},
				},
				"real_time_protection": basetypes.StringType{},
				"virdef_version": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"not_within": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"days":     basetypes.Int64Type{},
								"versions": basetypes.Int64Type{},
							},
						},
						"within": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"days":     basetypes.Int64Type{},
								"versions": basetypes.Int64Type{},
							},
						},
					},
				},
			},
		},
		"exclude_vendor": basetypes.BoolType{},
		"vendor": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"name":    basetypes.StringType{},
				"product": basetypes.ListType{ElemType: basetypes.StringType{}},
			},
		}},
	}
}

// AttrType returns the attribute type for a list of HipObjectsAntiMalware objects.
func (o HipObjectsAntiMalware) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the HipObjectsAntiMalwareCriteria model.
func (o HipObjectsAntiMalwareCriteria) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"is_installed": basetypes.BoolType{},
		"last_scan_time": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"not_available": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"not_within": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"days":  basetypes.Int64Type{},
						"hours": basetypes.Int64Type{},
					},
				},
				"within": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"days":  basetypes.Int64Type{},
						"hours": basetypes.Int64Type{},
					},
				},
			},
		},
		"product_version": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"contains":      basetypes.StringType{},
				"greater_equal": basetypes.StringType{},
				"greater_than":  basetypes.StringType{},
				"is":            basetypes.StringType{},
				"is_not":        basetypes.StringType{},
				"less_equal":    basetypes.StringType{},
				"less_than":     basetypes.StringType{},
				"not_within": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"versions": basetypes.Int64Type{},
					},
				},
				"within": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"versions": basetypes.Int64Type{},
					},
				},
			},
		},
		"real_time_protection": basetypes.StringType{},
		"virdef_version": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"not_within": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"days":     basetypes.Int64Type{},
						"versions": basetypes.Int64Type{},
					},
				},
				"within": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"days":     basetypes.Int64Type{},
						"versions": basetypes.Int64Type{},
					},
				},
			},
		},
	}
}

// AttrType returns the attribute type for a list of HipObjectsAntiMalwareCriteria objects.
func (o HipObjectsAntiMalwareCriteria) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the HipObjectsAntiMalwareCriteriaLastScanTime model.
func (o HipObjectsAntiMalwareCriteriaLastScanTime) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"not_available": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
		"not_within": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"days":  basetypes.Int64Type{},
				"hours": basetypes.Int64Type{},
			},
		},
		"within": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"days":  basetypes.Int64Type{},
				"hours": basetypes.Int64Type{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of HipObjectsAntiMalwareCriteriaLastScanTime objects.
func (o HipObjectsAntiMalwareCriteriaLastScanTime) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the HipObjectsAntiMalwareCriteriaLastScanTimeNotWithin model.
func (o HipObjectsAntiMalwareCriteriaLastScanTimeNotWithin) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"days":  basetypes.Int64Type{},
		"hours": basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of HipObjectsAntiMalwareCriteriaLastScanTimeNotWithin objects.
func (o HipObjectsAntiMalwareCriteriaLastScanTimeNotWithin) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the HipObjectsAntiMalwareCriteriaProductVersion model.
func (o HipObjectsAntiMalwareCriteriaProductVersion) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"contains":      basetypes.StringType{},
		"greater_equal": basetypes.StringType{},
		"greater_than":  basetypes.StringType{},
		"is":            basetypes.StringType{},
		"is_not":        basetypes.StringType{},
		"less_equal":    basetypes.StringType{},
		"less_than":     basetypes.StringType{},
		"not_within": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"versions": basetypes.Int64Type{},
			},
		},
		"within": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"versions": basetypes.Int64Type{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of HipObjectsAntiMalwareCriteriaProductVersion objects.
func (o HipObjectsAntiMalwareCriteriaProductVersion) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the HipObjectsAntiMalwareCriteriaProductVersionNotWithin model.
func (o HipObjectsAntiMalwareCriteriaProductVersionNotWithin) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"versions": basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of HipObjectsAntiMalwareCriteriaProductVersionNotWithin objects.
func (o HipObjectsAntiMalwareCriteriaProductVersionNotWithin) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the HipObjectsAntiMalwareCriteriaVirdefVersion model.
func (o HipObjectsAntiMalwareCriteriaVirdefVersion) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"not_within": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"days":     basetypes.Int64Type{},
				"versions": basetypes.Int64Type{},
			},
		},
		"within": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"days":     basetypes.Int64Type{},
				"versions": basetypes.Int64Type{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of HipObjectsAntiMalwareCriteriaVirdefVersion objects.
func (o HipObjectsAntiMalwareCriteriaVirdefVersion) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the HipObjectsAntiMalwareCriteriaVirdefVersionNotWithin model.
func (o HipObjectsAntiMalwareCriteriaVirdefVersionNotWithin) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"days":     basetypes.Int64Type{},
		"versions": basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of HipObjectsAntiMalwareCriteriaVirdefVersionNotWithin objects.
func (o HipObjectsAntiMalwareCriteriaVirdefVersionNotWithin) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the HipObjectsAntiMalwareVendorInner model.
func (o HipObjectsAntiMalwareVendorInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name":    basetypes.StringType{},
		"product": basetypes.ListType{ElemType: basetypes.StringType{}},
	}
}

// AttrType returns the attribute type for a list of HipObjectsAntiMalwareVendorInner objects.
func (o HipObjectsAntiMalwareVendorInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the HipObjectsCertificate model.
func (o HipObjectsCertificate) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"criteria": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"certificate_attributes": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"name":  basetypes.StringType{},
						"value": basetypes.StringType{},
					},
				}},
				"certificate_profile": basetypes.StringType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of HipObjectsCertificate objects.
func (o HipObjectsCertificate) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the HipObjectsCertificateCriteria model.
func (o HipObjectsCertificateCriteria) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"certificate_attributes": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"name":  basetypes.StringType{},
				"value": basetypes.StringType{},
			},
		}},
		"certificate_profile": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of HipObjectsCertificateCriteria objects.
func (o HipObjectsCertificateCriteria) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the HipObjectsCertificateCriteriaCertificateAttributesInner model.
func (o HipObjectsCertificateCriteriaCertificateAttributesInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name":  basetypes.StringType{},
		"value": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of HipObjectsCertificateCriteriaCertificateAttributesInner objects.
func (o HipObjectsCertificateCriteriaCertificateAttributesInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the HipObjectsCustomChecks model.
func (o HipObjectsCustomChecks) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"criteria": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"plist": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"key": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"name":   basetypes.StringType{},
								"negate": basetypes.BoolType{},
								"value":  basetypes.StringType{},
							},
						}},
						"name":   basetypes.StringType{},
						"negate": basetypes.BoolType{},
					},
				}},
				"process_list": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"name":    basetypes.StringType{},
						"running": basetypes.BoolType{},
					},
				}},
				"registry_key": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"default_value_data": basetypes.StringType{},
						"name":               basetypes.StringType{},
						"negate":             basetypes.BoolType{},
						"registry_value": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"name":       basetypes.StringType{},
								"negate":     basetypes.BoolType{},
								"value_data": basetypes.StringType{},
							},
						}},
					},
				}},
			},
		},
	}
}

// AttrType returns the attribute type for a list of HipObjectsCustomChecks objects.
func (o HipObjectsCustomChecks) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the HipObjectsCustomChecksCriteria model.
func (o HipObjectsCustomChecksCriteria) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"plist": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"key": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"name":   basetypes.StringType{},
						"negate": basetypes.BoolType{},
						"value":  basetypes.StringType{},
					},
				}},
				"name":   basetypes.StringType{},
				"negate": basetypes.BoolType{},
			},
		}},
		"process_list": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"name":    basetypes.StringType{},
				"running": basetypes.BoolType{},
			},
		}},
		"registry_key": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"default_value_data": basetypes.StringType{},
				"name":               basetypes.StringType{},
				"negate":             basetypes.BoolType{},
				"registry_value": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"name":       basetypes.StringType{},
						"negate":     basetypes.BoolType{},
						"value_data": basetypes.StringType{},
					},
				}},
			},
		}},
	}
}

// AttrType returns the attribute type for a list of HipObjectsCustomChecksCriteria objects.
func (o HipObjectsCustomChecksCriteria) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the HipObjectsCustomChecksCriteriaPlistInner model.
func (o HipObjectsCustomChecksCriteriaPlistInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"key": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"name":   basetypes.StringType{},
				"negate": basetypes.BoolType{},
				"value":  basetypes.StringType{},
			},
		}},
		"name":   basetypes.StringType{},
		"negate": basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of HipObjectsCustomChecksCriteriaPlistInner objects.
func (o HipObjectsCustomChecksCriteriaPlistInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the HipObjectsCustomChecksCriteriaPlistInnerKeyInner model.
func (o HipObjectsCustomChecksCriteriaPlistInnerKeyInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name":   basetypes.StringType{},
		"negate": basetypes.BoolType{},
		"value":  basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of HipObjectsCustomChecksCriteriaPlistInnerKeyInner objects.
func (o HipObjectsCustomChecksCriteriaPlistInnerKeyInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the HipObjectsCustomChecksCriteriaProcessListInner model.
func (o HipObjectsCustomChecksCriteriaProcessListInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name":    basetypes.StringType{},
		"running": basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of HipObjectsCustomChecksCriteriaProcessListInner objects.
func (o HipObjectsCustomChecksCriteriaProcessListInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the HipObjectsCustomChecksCriteriaRegistryKeyInner model.
func (o HipObjectsCustomChecksCriteriaRegistryKeyInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"default_value_data": basetypes.StringType{},
		"name":               basetypes.StringType{},
		"negate":             basetypes.BoolType{},
		"registry_value": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"name":       basetypes.StringType{},
				"negate":     basetypes.BoolType{},
				"value_data": basetypes.StringType{},
			},
		}},
	}
}

// AttrType returns the attribute type for a list of HipObjectsCustomChecksCriteriaRegistryKeyInner objects.
func (o HipObjectsCustomChecksCriteriaRegistryKeyInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the HipObjectsCustomChecksCriteriaRegistryKeyInnerRegistryValueInner model.
func (o HipObjectsCustomChecksCriteriaRegistryKeyInnerRegistryValueInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name":       basetypes.StringType{},
		"negate":     basetypes.BoolType{},
		"value_data": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of HipObjectsCustomChecksCriteriaRegistryKeyInnerRegistryValueInner objects.
func (o HipObjectsCustomChecksCriteriaRegistryKeyInnerRegistryValueInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the HipObjectsDataLossPrevention model.
func (o HipObjectsDataLossPrevention) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"criteria": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"is_enabled":   basetypes.StringType{},
				"is_installed": basetypes.BoolType{},
			},
		},
		"exclude_vendor": basetypes.BoolType{},
		"vendor": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"name":    basetypes.StringType{},
				"product": basetypes.ListType{ElemType: basetypes.StringType{}},
			},
		}},
	}
}

// AttrType returns the attribute type for a list of HipObjectsDataLossPrevention objects.
func (o HipObjectsDataLossPrevention) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the HipObjectsDataLossPreventionCriteria model.
func (o HipObjectsDataLossPreventionCriteria) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"is_enabled":   basetypes.StringType{},
		"is_installed": basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of HipObjectsDataLossPreventionCriteria objects.
func (o HipObjectsDataLossPreventionCriteria) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the HipObjectsDataLossPreventionVendorInner model.
func (o HipObjectsDataLossPreventionVendorInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name":    basetypes.StringType{},
		"product": basetypes.ListType{ElemType: basetypes.StringType{}},
	}
}

// AttrType returns the attribute type for a list of HipObjectsDataLossPreventionVendorInner objects.
func (o HipObjectsDataLossPreventionVendorInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the HipObjectsDiskBackup model.
func (o HipObjectsDiskBackup) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"criteria": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"is_installed": basetypes.BoolType{},
				"last_backup_time": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"not_available": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"not_within": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"days":  basetypes.Int64Type{},
								"hours": basetypes.Int64Type{},
							},
						},
						"within": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"days":  basetypes.Int64Type{},
								"hours": basetypes.Int64Type{},
							},
						},
					},
				},
			},
		},
		"exclude_vendor": basetypes.BoolType{},
		"vendor": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"name":    basetypes.StringType{},
				"product": basetypes.ListType{ElemType: basetypes.StringType{}},
			},
		}},
	}
}

// AttrType returns the attribute type for a list of HipObjectsDiskBackup objects.
func (o HipObjectsDiskBackup) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the HipObjectsDiskBackupCriteria model.
func (o HipObjectsDiskBackupCriteria) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"is_installed": basetypes.BoolType{},
		"last_backup_time": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"not_available": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"not_within": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"days":  basetypes.Int64Type{},
						"hours": basetypes.Int64Type{},
					},
				},
				"within": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"days":  basetypes.Int64Type{},
						"hours": basetypes.Int64Type{},
					},
				},
			},
		},
	}
}

// AttrType returns the attribute type for a list of HipObjectsDiskBackupCriteria objects.
func (o HipObjectsDiskBackupCriteria) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the HipObjectsDiskEncryption model.
func (o HipObjectsDiskEncryption) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"criteria": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"encrypted_locations": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"encryption_state": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"is":     basetypes.StringType{},
								"is_not": basetypes.StringType{},
							},
						},
						"name": basetypes.StringType{},
					},
				}},
				"is_installed": basetypes.BoolType{},
			},
		},
		"exclude_vendor": basetypes.BoolType{},
		"vendor": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"name":    basetypes.StringType{},
				"product": basetypes.ListType{ElemType: basetypes.StringType{}},
			},
		}},
	}
}

// AttrType returns the attribute type for a list of HipObjectsDiskEncryption objects.
func (o HipObjectsDiskEncryption) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the HipObjectsDiskEncryptionCriteria model.
func (o HipObjectsDiskEncryptionCriteria) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"encrypted_locations": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"encryption_state": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"is":     basetypes.StringType{},
						"is_not": basetypes.StringType{},
					},
				},
				"name": basetypes.StringType{},
			},
		}},
		"is_installed": basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of HipObjectsDiskEncryptionCriteria objects.
func (o HipObjectsDiskEncryptionCriteria) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the HipObjectsDiskEncryptionCriteriaEncryptedLocationsInner model.
func (o HipObjectsDiskEncryptionCriteriaEncryptedLocationsInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"encryption_state": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"is":     basetypes.StringType{},
				"is_not": basetypes.StringType{},
			},
		},
		"name": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of HipObjectsDiskEncryptionCriteriaEncryptedLocationsInner objects.
func (o HipObjectsDiskEncryptionCriteriaEncryptedLocationsInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the HipObjectsDiskEncryptionCriteriaEncryptedLocationsInnerEncryptionState model.
func (o HipObjectsDiskEncryptionCriteriaEncryptedLocationsInnerEncryptionState) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"is":     basetypes.StringType{},
		"is_not": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of HipObjectsDiskEncryptionCriteriaEncryptedLocationsInnerEncryptionState objects.
func (o HipObjectsDiskEncryptionCriteriaEncryptedLocationsInnerEncryptionState) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the HipObjectsFirewall model.
func (o HipObjectsFirewall) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"criteria": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"is_enabled":   basetypes.StringType{},
				"is_installed": basetypes.BoolType{},
			},
		},
		"exclude_vendor": basetypes.BoolType{},
		"vendor": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"name":    basetypes.StringType{},
				"product": basetypes.ListType{ElemType: basetypes.StringType{}},
			},
		}},
	}
}

// AttrType returns the attribute type for a list of HipObjectsFirewall objects.
func (o HipObjectsFirewall) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the HipObjectsHostInfo model.
func (o HipObjectsHostInfo) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"criteria": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"client_version": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"contains": basetypes.StringType{},
						"is":       basetypes.StringType{},
						"is_not":   basetypes.StringType{},
					},
				},
				"domain": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"contains": basetypes.StringType{},
						"is":       basetypes.StringType{},
						"is_not":   basetypes.StringType{},
					},
				},
				"host_id": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"contains": basetypes.StringType{},
						"is":       basetypes.StringType{},
						"is_not":   basetypes.StringType{},
					},
				},
				"host_name": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"contains": basetypes.StringType{},
						"is":       basetypes.StringType{},
						"is_not":   basetypes.StringType{},
					},
				},
				"managed": basetypes.BoolType{},
				"os": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"contains": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"apple":     basetypes.StringType{},
								"google":    basetypes.StringType{},
								"linux":     basetypes.StringType{},
								"microsoft": basetypes.StringType{},
								"other":     basetypes.StringType{},
							},
						},
					},
				},
				"serial_number": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"contains": basetypes.StringType{},
						"is":       basetypes.StringType{},
						"is_not":   basetypes.StringType{},
					},
				},
			},
		},
	}
}

// AttrType returns the attribute type for a list of HipObjectsHostInfo objects.
func (o HipObjectsHostInfo) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the HipObjectsHostInfoCriteria model.
func (o HipObjectsHostInfoCriteria) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"client_version": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"contains": basetypes.StringType{},
				"is":       basetypes.StringType{},
				"is_not":   basetypes.StringType{},
			},
		},
		"domain": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"contains": basetypes.StringType{},
				"is":       basetypes.StringType{},
				"is_not":   basetypes.StringType{},
			},
		},
		"host_id": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"contains": basetypes.StringType{},
				"is":       basetypes.StringType{},
				"is_not":   basetypes.StringType{},
			},
		},
		"host_name": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"contains": basetypes.StringType{},
				"is":       basetypes.StringType{},
				"is_not":   basetypes.StringType{},
			},
		},
		"managed": basetypes.BoolType{},
		"os": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"contains": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"apple":     basetypes.StringType{},
						"google":    basetypes.StringType{},
						"linux":     basetypes.StringType{},
						"microsoft": basetypes.StringType{},
						"other":     basetypes.StringType{},
					},
				},
			},
		},
		"serial_number": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"contains": basetypes.StringType{},
				"is":       basetypes.StringType{},
				"is_not":   basetypes.StringType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of HipObjectsHostInfoCriteria objects.
func (o HipObjectsHostInfoCriteria) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the HipObjectsHostInfoCriteriaClientVersion model.
func (o HipObjectsHostInfoCriteriaClientVersion) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"contains": basetypes.StringType{},
		"is":       basetypes.StringType{},
		"is_not":   basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of HipObjectsHostInfoCriteriaClientVersion objects.
func (o HipObjectsHostInfoCriteriaClientVersion) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the HipObjectsHostInfoCriteriaOs model.
func (o HipObjectsHostInfoCriteriaOs) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"contains": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"apple":     basetypes.StringType{},
				"google":    basetypes.StringType{},
				"linux":     basetypes.StringType{},
				"microsoft": basetypes.StringType{},
				"other":     basetypes.StringType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of HipObjectsHostInfoCriteriaOs objects.
func (o HipObjectsHostInfoCriteriaOs) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the HipObjectsHostInfoCriteriaOsContains model.
func (o HipObjectsHostInfoCriteriaOsContains) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"apple":     basetypes.StringType{},
		"google":    basetypes.StringType{},
		"linux":     basetypes.StringType{},
		"microsoft": basetypes.StringType{},
		"other":     basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of HipObjectsHostInfoCriteriaOsContains objects.
func (o HipObjectsHostInfoCriteriaOsContains) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the HipObjectsMobileDevice model.
func (o HipObjectsMobileDevice) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"criteria": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"applications": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"has_malware": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"no": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"yes": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"excludes": basetypes.ListType{ElemType: basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"hash":    basetypes.StringType{},
												"name":    basetypes.StringType{},
												"package": basetypes.StringType{},
											},
										}},
									},
								},
							},
						},
						"has_unmanaged_app": basetypes.BoolType{},
						"includes": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"hash":    basetypes.StringType{},
								"name":    basetypes.StringType{},
								"package": basetypes.StringType{},
							},
						}},
					},
				},
				"disk_encrypted": basetypes.BoolType{},
				"imei": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"contains": basetypes.StringType{},
						"is":       basetypes.StringType{},
						"is_not":   basetypes.StringType{},
					},
				},
				"jailbroken": basetypes.BoolType{},
				"last_checkin_time": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"not_within": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"days": basetypes.Int64Type{},
							},
						},
						"within": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"days": basetypes.Int64Type{},
							},
						},
					},
				},
				"model": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"contains": basetypes.StringType{},
						"is":       basetypes.StringType{},
						"is_not":   basetypes.StringType{},
					},
				},
				"passcode_set": basetypes.BoolType{},
				"phone_number": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"contains": basetypes.StringType{},
						"is":       basetypes.StringType{},
						"is_not":   basetypes.StringType{},
					},
				},
				"tag": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"contains": basetypes.StringType{},
						"is":       basetypes.StringType{},
						"is_not":   basetypes.StringType{},
					},
				},
			},
		},
	}
}

// AttrType returns the attribute type for a list of HipObjectsMobileDevice objects.
func (o HipObjectsMobileDevice) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the HipObjectsMobileDeviceCriteria model.
func (o HipObjectsMobileDeviceCriteria) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"applications": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"has_malware": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"no": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"yes": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"excludes": basetypes.ListType{ElemType: basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"hash":    basetypes.StringType{},
										"name":    basetypes.StringType{},
										"package": basetypes.StringType{},
									},
								}},
							},
						},
					},
				},
				"has_unmanaged_app": basetypes.BoolType{},
				"includes": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"hash":    basetypes.StringType{},
						"name":    basetypes.StringType{},
						"package": basetypes.StringType{},
					},
				}},
			},
		},
		"disk_encrypted": basetypes.BoolType{},
		"imei": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"contains": basetypes.StringType{},
				"is":       basetypes.StringType{},
				"is_not":   basetypes.StringType{},
			},
		},
		"jailbroken": basetypes.BoolType{},
		"last_checkin_time": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"not_within": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"days": basetypes.Int64Type{},
					},
				},
				"within": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"days": basetypes.Int64Type{},
					},
				},
			},
		},
		"model": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"contains": basetypes.StringType{},
				"is":       basetypes.StringType{},
				"is_not":   basetypes.StringType{},
			},
		},
		"passcode_set": basetypes.BoolType{},
		"phone_number": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"contains": basetypes.StringType{},
				"is":       basetypes.StringType{},
				"is_not":   basetypes.StringType{},
			},
		},
		"tag": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"contains": basetypes.StringType{},
				"is":       basetypes.StringType{},
				"is_not":   basetypes.StringType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of HipObjectsMobileDeviceCriteria objects.
func (o HipObjectsMobileDeviceCriteria) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the HipObjectsMobileDeviceCriteriaApplications model.
func (o HipObjectsMobileDeviceCriteriaApplications) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"has_malware": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"no": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"yes": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"excludes": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"hash":    basetypes.StringType{},
								"name":    basetypes.StringType{},
								"package": basetypes.StringType{},
							},
						}},
					},
				},
			},
		},
		"has_unmanaged_app": basetypes.BoolType{},
		"includes": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"hash":    basetypes.StringType{},
				"name":    basetypes.StringType{},
				"package": basetypes.StringType{},
			},
		}},
	}
}

// AttrType returns the attribute type for a list of HipObjectsMobileDeviceCriteriaApplications objects.
func (o HipObjectsMobileDeviceCriteriaApplications) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the HipObjectsMobileDeviceCriteriaApplicationsHasMalware model.
func (o HipObjectsMobileDeviceCriteriaApplicationsHasMalware) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"no": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
		"yes": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"excludes": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"hash":    basetypes.StringType{},
						"name":    basetypes.StringType{},
						"package": basetypes.StringType{},
					},
				}},
			},
		},
	}
}

// AttrType returns the attribute type for a list of HipObjectsMobileDeviceCriteriaApplicationsHasMalware objects.
func (o HipObjectsMobileDeviceCriteriaApplicationsHasMalware) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the HipObjectsMobileDeviceCriteriaApplicationsHasMalwareYes model.
func (o HipObjectsMobileDeviceCriteriaApplicationsHasMalwareYes) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"excludes": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"hash":    basetypes.StringType{},
				"name":    basetypes.StringType{},
				"package": basetypes.StringType{},
			},
		}},
	}
}

// AttrType returns the attribute type for a list of HipObjectsMobileDeviceCriteriaApplicationsHasMalwareYes objects.
func (o HipObjectsMobileDeviceCriteriaApplicationsHasMalwareYes) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the HipObjectsMobileDeviceCriteriaApplicationsHasMalwareYesExcludesInner model.
func (o HipObjectsMobileDeviceCriteriaApplicationsHasMalwareYesExcludesInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"hash":    basetypes.StringType{},
		"name":    basetypes.StringType{},
		"package": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of HipObjectsMobileDeviceCriteriaApplicationsHasMalwareYesExcludesInner objects.
func (o HipObjectsMobileDeviceCriteriaApplicationsHasMalwareYesExcludesInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the HipObjectsMobileDeviceCriteriaLastCheckinTime model.
func (o HipObjectsMobileDeviceCriteriaLastCheckinTime) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"not_within": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"days": basetypes.Int64Type{},
			},
		},
		"within": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"days": basetypes.Int64Type{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of HipObjectsMobileDeviceCriteriaLastCheckinTime objects.
func (o HipObjectsMobileDeviceCriteriaLastCheckinTime) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the HipObjectsMobileDeviceCriteriaLastCheckinTimeNotWithin model.
func (o HipObjectsMobileDeviceCriteriaLastCheckinTimeNotWithin) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"days": basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of HipObjectsMobileDeviceCriteriaLastCheckinTimeNotWithin objects.
func (o HipObjectsMobileDeviceCriteriaLastCheckinTimeNotWithin) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the HipObjectsNetworkInfo model.
func (o HipObjectsNetworkInfo) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"criteria": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"network": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"is": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"mobile": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"carrier": basetypes.StringType{},
									},
								},
								"unknown": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"wifi": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"ssid": basetypes.StringType{},
									},
								},
							},
						},
						"is_not": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"ethernet": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"mobile": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"carrier": basetypes.StringType{},
									},
								},
								"unknown": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"wifi": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"ssid": basetypes.StringType{},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

// AttrType returns the attribute type for a list of HipObjectsNetworkInfo objects.
func (o HipObjectsNetworkInfo) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the HipObjectsNetworkInfoCriteria model.
func (o HipObjectsNetworkInfoCriteria) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"network": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"is": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"mobile": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"carrier": basetypes.StringType{},
							},
						},
						"unknown": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"wifi": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"ssid": basetypes.StringType{},
							},
						},
					},
				},
				"is_not": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"ethernet": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"mobile": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"carrier": basetypes.StringType{},
							},
						},
						"unknown": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"wifi": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"ssid": basetypes.StringType{},
							},
						},
					},
				},
			},
		},
	}
}

// AttrType returns the attribute type for a list of HipObjectsNetworkInfoCriteria objects.
func (o HipObjectsNetworkInfoCriteria) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the HipObjectsNetworkInfoCriteriaNetwork model.
func (o HipObjectsNetworkInfoCriteriaNetwork) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"is": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"mobile": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"carrier": basetypes.StringType{},
					},
				},
				"unknown": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"wifi": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"ssid": basetypes.StringType{},
					},
				},
			},
		},
		"is_not": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"ethernet": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"mobile": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"carrier": basetypes.StringType{},
					},
				},
				"unknown": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"wifi": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"ssid": basetypes.StringType{},
					},
				},
			},
		},
	}
}

// AttrType returns the attribute type for a list of HipObjectsNetworkInfoCriteriaNetwork objects.
func (o HipObjectsNetworkInfoCriteriaNetwork) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the HipObjectsNetworkInfoCriteriaNetworkIs model.
func (o HipObjectsNetworkInfoCriteriaNetworkIs) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"mobile": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"carrier": basetypes.StringType{},
			},
		},
		"unknown": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
		"wifi": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"ssid": basetypes.StringType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of HipObjectsNetworkInfoCriteriaNetworkIs objects.
func (o HipObjectsNetworkInfoCriteriaNetworkIs) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the HipObjectsNetworkInfoCriteriaNetworkIsMobile model.
func (o HipObjectsNetworkInfoCriteriaNetworkIsMobile) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"carrier": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of HipObjectsNetworkInfoCriteriaNetworkIsMobile objects.
func (o HipObjectsNetworkInfoCriteriaNetworkIsMobile) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the HipObjectsNetworkInfoCriteriaNetworkIsWifi model.
func (o HipObjectsNetworkInfoCriteriaNetworkIsWifi) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"ssid": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of HipObjectsNetworkInfoCriteriaNetworkIsWifi objects.
func (o HipObjectsNetworkInfoCriteriaNetworkIsWifi) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the HipObjectsNetworkInfoCriteriaNetworkIsNot model.
func (o HipObjectsNetworkInfoCriteriaNetworkIsNot) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"ethernet": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
		"mobile": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"carrier": basetypes.StringType{},
			},
		},
		"unknown": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
		"wifi": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"ssid": basetypes.StringType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of HipObjectsNetworkInfoCriteriaNetworkIsNot objects.
func (o HipObjectsNetworkInfoCriteriaNetworkIsNot) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the HipObjectsPatchManagement model.
func (o HipObjectsPatchManagement) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"criteria": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"is_enabled":   basetypes.StringType{},
				"is_installed": basetypes.BoolType{},
				"missing_patches": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"check":   basetypes.StringType{},
						"patches": basetypes.ListType{ElemType: basetypes.StringType{}},
						"severity": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"greater_equal": basetypes.Int64Type{},
								"greater_than":  basetypes.Int64Type{},
								"is":            basetypes.Int64Type{},
								"is_not":        basetypes.Int64Type{},
								"less_equal":    basetypes.Int64Type{},
								"less_than":     basetypes.Int64Type{},
							},
						},
					},
				},
			},
		},
		"exclude_vendor": basetypes.BoolType{},
		"vendor": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"name":    basetypes.StringType{},
				"product": basetypes.ListType{ElemType: basetypes.StringType{}},
			},
		}},
	}
}

// AttrType returns the attribute type for a list of HipObjectsPatchManagement objects.
func (o HipObjectsPatchManagement) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the HipObjectsPatchManagementCriteria model.
func (o HipObjectsPatchManagementCriteria) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"is_enabled":   basetypes.StringType{},
		"is_installed": basetypes.BoolType{},
		"missing_patches": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"check":   basetypes.StringType{},
				"patches": basetypes.ListType{ElemType: basetypes.StringType{}},
				"severity": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"greater_equal": basetypes.Int64Type{},
						"greater_than":  basetypes.Int64Type{},
						"is":            basetypes.Int64Type{},
						"is_not":        basetypes.Int64Type{},
						"less_equal":    basetypes.Int64Type{},
						"less_than":     basetypes.Int64Type{},
					},
				},
			},
		},
	}
}

// AttrType returns the attribute type for a list of HipObjectsPatchManagementCriteria objects.
func (o HipObjectsPatchManagementCriteria) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the HipObjectsPatchManagementCriteriaMissingPatches model.
func (o HipObjectsPatchManagementCriteriaMissingPatches) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"check":   basetypes.StringType{},
		"patches": basetypes.ListType{ElemType: basetypes.StringType{}},
		"severity": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"greater_equal": basetypes.Int64Type{},
				"greater_than":  basetypes.Int64Type{},
				"is":            basetypes.Int64Type{},
				"is_not":        basetypes.Int64Type{},
				"less_equal":    basetypes.Int64Type{},
				"less_than":     basetypes.Int64Type{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of HipObjectsPatchManagementCriteriaMissingPatches objects.
func (o HipObjectsPatchManagementCriteriaMissingPatches) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the HipObjectsPatchManagementCriteriaMissingPatchesSeverity model.
func (o HipObjectsPatchManagementCriteriaMissingPatchesSeverity) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"greater_equal": basetypes.Int64Type{},
		"greater_than":  basetypes.Int64Type{},
		"is":            basetypes.Int64Type{},
		"is_not":        basetypes.Int64Type{},
		"less_equal":    basetypes.Int64Type{},
		"less_than":     basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of HipObjectsPatchManagementCriteriaMissingPatchesSeverity objects.
func (o HipObjectsPatchManagementCriteriaMissingPatchesSeverity) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// HipObjectsResourceSchema defines the schema for HipObjects resource
var HipObjectsResourceSchema = schema.Schema{
	MarkdownDescription: "HipObject resource",
	Attributes: map[string]schema.Attribute{
		"anti_malware": schema.SingleNestedAttribute{
			MarkdownDescription: "Anti malware",
			Optional:            true,
			Computed:            true,
			Attributes: map[string]schema.Attribute{
				"criteria": schema.SingleNestedAttribute{
					MarkdownDescription: "Criteria",
					Optional:            true,
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"is_installed": schema.BoolAttribute{
							MarkdownDescription: "Is Installed",
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(true),
						},
						"last_scan_time": schema.SingleNestedAttribute{
							MarkdownDescription: "Last scan time",
							Optional:            true,
							Computed:            true,
							Attributes: map[string]schema.Attribute{
								"not_available": schema.SingleNestedAttribute{
									Validators: []validator.Object{
										objectvalidator.ConflictsWith(
											path.MatchRelative().AtParent().AtName("not_within"),
											path.MatchRelative().AtParent().AtName("within"),
										),
									},
									MarkdownDescription: "Not available",
									Optional:            true,
									Computed:            true,
									Attributes:          map[string]schema.Attribute{},
								},
								"not_within": schema.SingleNestedAttribute{
									Validators: []validator.Object{
										objectvalidator.ConflictsWith(
											path.MatchRelative().AtParent().AtName("not_available"),
											path.MatchRelative().AtParent().AtName("within"),
										),
									},
									MarkdownDescription: "Not within",
									Optional:            true,
									Computed:            true,
									Attributes: map[string]schema.Attribute{
										"days": schema.Int64Attribute{
											Validators: []validator.Int64{
												int64validator.ConflictsWith(
													path.MatchRelative().AtParent().AtName("hours"),
												),
												int64validator.Between(1, 65535),
											},
											MarkdownDescription: "specify time in days",
											Optional:            true,
											Computed:            true,
										},
										"hours": schema.Int64Attribute{
											Validators: []validator.Int64{
												int64validator.ConflictsWith(
													path.MatchRelative().AtParent().AtName("days"),
												),
												int64validator.Between(1, 65535),
											},
											MarkdownDescription: "specify time in hours",
											Optional:            true,
											Computed:            true,
										},
									},
								},
								"within": schema.SingleNestedAttribute{
									Validators: []validator.Object{
										objectvalidator.ConflictsWith(
											path.MatchRelative().AtParent().AtName("not_available"),
											path.MatchRelative().AtParent().AtName("not_within"),
										),
									},
									MarkdownDescription: "Within",
									Optional:            true,
									Computed:            true,
									Attributes: map[string]schema.Attribute{
										"days": schema.Int64Attribute{
											Validators: []validator.Int64{
												int64validator.ConflictsWith(
													path.MatchRelative().AtParent().AtName("hours"),
												),
												int64validator.Between(1, 65535),
											},
											MarkdownDescription: "specify time in days",
											Optional:            true,
											Computed:            true,
										},
										"hours": schema.Int64Attribute{
											Validators: []validator.Int64{
												int64validator.ConflictsWith(
													path.MatchRelative().AtParent().AtName("days"),
												),
												int64validator.Between(1, 65535),
											},
											MarkdownDescription: "specify time in hours",
											Optional:            true,
											Computed:            true,
										},
									},
								},
							},
						},
						"product_version": schema.SingleNestedAttribute{
							MarkdownDescription: "Product version",
							Optional:            true,
							Computed:            true,
							Attributes: map[string]schema.Attribute{
								"contains": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.ConflictsWith(
											path.MatchRelative().AtParent().AtName("greater_equal"),
											path.MatchRelative().AtParent().AtName("greater_than"),
											path.MatchRelative().AtParent().AtName("is"),
											path.MatchRelative().AtParent().AtName("is_not"),
											path.MatchRelative().AtParent().AtName("less_equal"),
											path.MatchRelative().AtParent().AtName("less_than"),
											path.MatchRelative().AtParent().AtName("not_within"),
											path.MatchRelative().AtParent().AtName("within"),
										),
										stringvalidator.LengthAtMost(255),
									},
									MarkdownDescription: "Contains",
									Optional:            true,
									Computed:            true,
								},
								"greater_equal": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.ConflictsWith(
											path.MatchRelative().AtParent().AtName("contains"),
											path.MatchRelative().AtParent().AtName("greater_than"),
											path.MatchRelative().AtParent().AtName("is"),
											path.MatchRelative().AtParent().AtName("is_not"),
											path.MatchRelative().AtParent().AtName("less_equal"),
											path.MatchRelative().AtParent().AtName("less_than"),
											path.MatchRelative().AtParent().AtName("not_within"),
											path.MatchRelative().AtParent().AtName("within"),
										),
										stringvalidator.LengthAtMost(255),
									},
									MarkdownDescription: "Greater equal",
									Optional:            true,
									Computed:            true,
								},
								"greater_than": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.ConflictsWith(
											path.MatchRelative().AtParent().AtName("contains"),
											path.MatchRelative().AtParent().AtName("greater_equal"),
											path.MatchRelative().AtParent().AtName("is"),
											path.MatchRelative().AtParent().AtName("is_not"),
											path.MatchRelative().AtParent().AtName("less_equal"),
											path.MatchRelative().AtParent().AtName("less_than"),
											path.MatchRelative().AtParent().AtName("not_within"),
											path.MatchRelative().AtParent().AtName("within"),
										),
										stringvalidator.LengthAtMost(255),
									},
									MarkdownDescription: "Greater than",
									Optional:            true,
									Computed:            true,
								},
								"is": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.ConflictsWith(
											path.MatchRelative().AtParent().AtName("contains"),
											path.MatchRelative().AtParent().AtName("greater_equal"),
											path.MatchRelative().AtParent().AtName("greater_than"),
											path.MatchRelative().AtParent().AtName("is_not"),
											path.MatchRelative().AtParent().AtName("less_equal"),
											path.MatchRelative().AtParent().AtName("less_than"),
											path.MatchRelative().AtParent().AtName("not_within"),
											path.MatchRelative().AtParent().AtName("within"),
										),
										stringvalidator.LengthAtMost(255),
									},
									MarkdownDescription: "Is",
									Optional:            true,
									Computed:            true,
								},
								"is_not": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.ConflictsWith(
											path.MatchRelative().AtParent().AtName("contains"),
											path.MatchRelative().AtParent().AtName("greater_equal"),
											path.MatchRelative().AtParent().AtName("greater_than"),
											path.MatchRelative().AtParent().AtName("is"),
											path.MatchRelative().AtParent().AtName("less_equal"),
											path.MatchRelative().AtParent().AtName("less_than"),
											path.MatchRelative().AtParent().AtName("not_within"),
											path.MatchRelative().AtParent().AtName("within"),
										),
										stringvalidator.LengthAtMost(255),
									},
									MarkdownDescription: "Is not",
									Optional:            true,
									Computed:            true,
								},
								"less_equal": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.ConflictsWith(
											path.MatchRelative().AtParent().AtName("contains"),
											path.MatchRelative().AtParent().AtName("greater_equal"),
											path.MatchRelative().AtParent().AtName("greater_than"),
											path.MatchRelative().AtParent().AtName("is"),
											path.MatchRelative().AtParent().AtName("is_not"),
											path.MatchRelative().AtParent().AtName("less_than"),
											path.MatchRelative().AtParent().AtName("not_within"),
											path.MatchRelative().AtParent().AtName("within"),
										),
										stringvalidator.LengthAtMost(255),
									},
									MarkdownDescription: "Less equal",
									Optional:            true,
									Computed:            true,
								},
								"less_than": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.ConflictsWith(
											path.MatchRelative().AtParent().AtName("contains"),
											path.MatchRelative().AtParent().AtName("greater_equal"),
											path.MatchRelative().AtParent().AtName("greater_than"),
											path.MatchRelative().AtParent().AtName("is"),
											path.MatchRelative().AtParent().AtName("is_not"),
											path.MatchRelative().AtParent().AtName("less_equal"),
											path.MatchRelative().AtParent().AtName("not_within"),
											path.MatchRelative().AtParent().AtName("within"),
										),
										stringvalidator.LengthAtMost(255),
									},
									MarkdownDescription: "Less than",
									Optional:            true,
									Computed:            true,
								},
								"not_within": schema.SingleNestedAttribute{
									Validators: []validator.Object{
										objectvalidator.ConflictsWith(
											path.MatchRelative().AtParent().AtName("contains"),
											path.MatchRelative().AtParent().AtName("greater_equal"),
											path.MatchRelative().AtParent().AtName("greater_than"),
											path.MatchRelative().AtParent().AtName("is"),
											path.MatchRelative().AtParent().AtName("is_not"),
											path.MatchRelative().AtParent().AtName("less_equal"),
											path.MatchRelative().AtParent().AtName("less_than"),
											path.MatchRelative().AtParent().AtName("within"),
										),
									},
									MarkdownDescription: "Not within",
									Optional:            true,
									Computed:            true,
									Attributes: map[string]schema.Attribute{
										"versions": schema.Int64Attribute{
											Validators: []validator.Int64{
												int64validator.Between(1, 65535),
											},
											MarkdownDescription: "versions range",
											Required:            true,
										},
									},
								},
								"within": schema.SingleNestedAttribute{
									Validators: []validator.Object{
										objectvalidator.ConflictsWith(
											path.MatchRelative().AtParent().AtName("contains"),
											path.MatchRelative().AtParent().AtName("greater_equal"),
											path.MatchRelative().AtParent().AtName("greater_than"),
											path.MatchRelative().AtParent().AtName("is"),
											path.MatchRelative().AtParent().AtName("is_not"),
											path.MatchRelative().AtParent().AtName("less_equal"),
											path.MatchRelative().AtParent().AtName("less_than"),
											path.MatchRelative().AtParent().AtName("not_within"),
										),
									},
									MarkdownDescription: "Within",
									Optional:            true,
									Computed:            true,
									Attributes: map[string]schema.Attribute{
										"versions": schema.Int64Attribute{
											Validators: []validator.Int64{
												int64validator.Between(1, 65535),
											},
											MarkdownDescription: "versions range",
											Required:            true,
										},
									},
								},
							},
						},
						"real_time_protection": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("no", "yes", "not-available"),
							},
							MarkdownDescription: "real time protection",
							Optional:            true,
							Computed:            true,
						},
						"virdef_version": schema.SingleNestedAttribute{
							MarkdownDescription: "Virdef version",
							Optional:            true,
							Computed:            true,
							Attributes: map[string]schema.Attribute{
								"not_within": schema.SingleNestedAttribute{
									Validators: []validator.Object{
										objectvalidator.ConflictsWith(
											path.MatchRelative().AtParent().AtName("within"),
										),
									},
									MarkdownDescription: "Not within",
									Optional:            true,
									Computed:            true,
									Attributes: map[string]schema.Attribute{
										"days": schema.Int64Attribute{
											Validators: []validator.Int64{
												int64validator.ConflictsWith(
													path.MatchRelative().AtParent().AtName("versions"),
												),
												int64validator.Between(1, 65535),
											},
											MarkdownDescription: "specify time in days",
											Optional:            true,
											Computed:            true,
										},
										"versions": schema.Int64Attribute{
											Validators: []validator.Int64{
												int64validator.ConflictsWith(
													path.MatchRelative().AtParent().AtName("days"),
												),
												int64validator.Between(1, 65535),
											},
											MarkdownDescription: "specify versions range",
											Optional:            true,
											Computed:            true,
										},
									},
								},
								"within": schema.SingleNestedAttribute{
									Validators: []validator.Object{
										objectvalidator.ConflictsWith(
											path.MatchRelative().AtParent().AtName("not_within"),
										),
									},
									MarkdownDescription: "Within",
									Optional:            true,
									Computed:            true,
									Attributes: map[string]schema.Attribute{
										"days": schema.Int64Attribute{
											Validators: []validator.Int64{
												int64validator.ConflictsWith(
													path.MatchRelative().AtParent().AtName("versions"),
												),
												int64validator.Between(1, 65535),
											},
											MarkdownDescription: "specify time in days",
											Optional:            true,
											Computed:            true,
										},
										"versions": schema.Int64Attribute{
											Validators: []validator.Int64{
												int64validator.ConflictsWith(
													path.MatchRelative().AtParent().AtName("days"),
												),
												int64validator.Between(1, 65535),
											},
											MarkdownDescription: "specify versions range",
											Optional:            true,
											Computed:            true,
										},
									},
								},
							},
						},
					},
				},
				"exclude_vendor": schema.BoolAttribute{
					MarkdownDescription: "Exclude vendor",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(false),
				},
				"vendor": schema.ListNestedAttribute{
					MarkdownDescription: "Vendor name",
					Optional:            true,
					Computed:            true,
					NestedObject: schema.NestedAttributeObject{
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Validators: []validator.String{
									stringvalidator.LengthAtMost(103),
								},
								MarkdownDescription: "Name",
								Required:            true,
							},
							"product": schema.ListAttribute{
								ElementType:         types.StringType,
								MarkdownDescription: "Product",
								Validators: []validator.List{
									listvalidator.ValueStringsAre(stringvalidator.LengthAtMost(1023)),
								},
								Optional: true,
							},
						},
					},
				},
			},
		},
		"certificate": schema.SingleNestedAttribute{
			MarkdownDescription: "Certificate",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"criteria": schema.SingleNestedAttribute{
					MarkdownDescription: "Criteria",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"certificate_attributes": schema.ListNestedAttribute{
							MarkdownDescription: "Certificate attributes",
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										MarkdownDescription: "Attribute Name",
										Required:            true,
									},
									"value": schema.StringAttribute{
										Validators: []validator.String{
											stringvalidator.LengthAtMost(1024),
											stringvalidator.RegexMatches(regexp.MustCompile(".*"), "pattern must match "+".*"),
										},
										MarkdownDescription: "Key value",
										Optional:            true,
									},
								},
							},
						},
						"certificate_profile": schema.StringAttribute{
							MarkdownDescription: "Profile for authenticating client certificates",
							Optional:            true,
						},
					},
				},
			},
		},
		"custom_checks": schema.SingleNestedAttribute{
			MarkdownDescription: "Custom checks",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"criteria": schema.SingleNestedAttribute{
					MarkdownDescription: "Criteria",
					Required:            true,
					Attributes: map[string]schema.Attribute{
						"plist": schema.ListNestedAttribute{
							MarkdownDescription: "Plist",
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"key": schema.ListNestedAttribute{
										MarkdownDescription: "Key",
										Optional:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Validators: []validator.String{
														stringvalidator.LengthAtMost(1023),
													},
													MarkdownDescription: "Key name",
													Required:            true,
												},
												"negate": schema.BoolAttribute{
													MarkdownDescription: "Value does not exist or match specified value data",
													Optional:            true,
													Computed:            true,
													Default:             booldefault.StaticBool(false),
												},
												"value": schema.StringAttribute{
													Validators: []validator.String{
														stringvalidator.LengthAtMost(1024),
														stringvalidator.RegexMatches(regexp.MustCompile(".*"), "pattern must match "+".*"),
													},
													MarkdownDescription: "Key value",
													Optional:            true,
												},
											},
										},
									},
									"name": schema.StringAttribute{
										Validators: []validator.String{
											stringvalidator.LengthAtMost(1023),
										},
										MarkdownDescription: "Preference list",
										Required:            true,
									},
									"negate": schema.BoolAttribute{
										MarkdownDescription: "Plist does not exist",
										Optional:            true,
										Computed:            true,
										Default:             booldefault.StaticBool(false),
									},
								},
							},
						},
						"process_list": schema.ListNestedAttribute{
							MarkdownDescription: "Process list",
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Validators: []validator.String{
											stringvalidator.LengthAtMost(1023),
										},
										MarkdownDescription: "Process Name",
										Required:            true,
									},
									"running": schema.BoolAttribute{
										MarkdownDescription: "Running",
										Optional:            true,
										Computed:            true,
										Default:             booldefault.StaticBool(true),
									},
								},
							},
						},
						"registry_key": schema.ListNestedAttribute{
							MarkdownDescription: "Registry key",
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"default_value_data": schema.StringAttribute{
										Validators: []validator.String{
											stringvalidator.LengthAtMost(1024),
											stringvalidator.RegexMatches(regexp.MustCompile(".*"), "pattern must match "+".*"),
										},
										MarkdownDescription: "Registry key default value data",
										Optional:            true,
									},
									"name": schema.StringAttribute{
										Validators: []validator.String{
											stringvalidator.LengthAtMost(1023),
										},
										MarkdownDescription: "Registry key",
										Required:            true,
									},
									"negate": schema.BoolAttribute{
										MarkdownDescription: "Key does not exist or match specified value data",
										Optional:            true,
										Computed:            true,
										Default:             booldefault.StaticBool(false),
									},
									"registry_value": schema.ListNestedAttribute{
										MarkdownDescription: "Registry value",
										Optional:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Validators: []validator.String{
														stringvalidator.LengthAtMost(1023),
													},
													MarkdownDescription: "Registry value name",
													Required:            true,
												},
												"negate": schema.BoolAttribute{
													MarkdownDescription: "Value does not exist or match specified value data",
													Optional:            true,
													Computed:            true,
													Default:             booldefault.StaticBool(false),
												},
												"value_data": schema.StringAttribute{
													Validators: []validator.String{
														stringvalidator.LengthAtMost(1024),
														stringvalidator.RegexMatches(regexp.MustCompile(".*"), "pattern must match "+".*"),
													},
													MarkdownDescription: "Registry value data",
													Optional:            true,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"data_loss_prevention": schema.SingleNestedAttribute{
			MarkdownDescription: "Data loss prevention",
			Optional:            true,
			Computed:            true,
			Attributes: map[string]schema.Attribute{
				"criteria": schema.SingleNestedAttribute{
					MarkdownDescription: "Criteria",
					Optional:            true,
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"is_enabled": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("no", "yes", "not-available"),
							},
							MarkdownDescription: "is enabled",
							Optional:            true,
							Computed:            true,
						},
						"is_installed": schema.BoolAttribute{
							MarkdownDescription: "Is Installed",
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(true),
						},
					},
				},
				"exclude_vendor": schema.BoolAttribute{
					MarkdownDescription: "Exclude vendor",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(false),
				},
				"vendor": schema.ListNestedAttribute{
					MarkdownDescription: "Vendor name",
					Optional:            true,
					Computed:            true,
					NestedObject: schema.NestedAttributeObject{
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Validators: []validator.String{
									stringvalidator.LengthAtMost(103),
								},
								MarkdownDescription: "Name",
								Required:            true,
							},
							"product": schema.ListAttribute{
								ElementType:         types.StringType,
								MarkdownDescription: "Product name",
								Validators: []validator.List{
									listvalidator.ValueStringsAre(stringvalidator.LengthAtMost(1023)),
								},
								Optional: true,
							},
						},
					},
				},
			},
		},
		"description": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(255),
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
		"disk_backup": schema.SingleNestedAttribute{
			MarkdownDescription: "Disk backup",
			Optional:            true,
			Computed:            true,
			Attributes: map[string]schema.Attribute{
				"criteria": schema.SingleNestedAttribute{
					MarkdownDescription: "Criteria",
					Optional:            true,
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"is_installed": schema.BoolAttribute{
							MarkdownDescription: "Is Installed",
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(true),
						},
						"last_backup_time": schema.SingleNestedAttribute{
							MarkdownDescription: "Last backup time",
							Optional:            true,
							Computed:            true,
							Attributes: map[string]schema.Attribute{
								"not_available": schema.SingleNestedAttribute{
									Validators: []validator.Object{
										objectvalidator.ConflictsWith(
											path.MatchRelative().AtParent().AtName("not_within"),
											path.MatchRelative().AtParent().AtName("within"),
										),
									},
									MarkdownDescription: "Not available",
									Optional:            true,
									Computed:            true,
									Attributes:          map[string]schema.Attribute{},
								},
								"not_within": schema.SingleNestedAttribute{
									Validators: []validator.Object{
										objectvalidator.ConflictsWith(
											path.MatchRelative().AtParent().AtName("not_available"),
											path.MatchRelative().AtParent().AtName("within"),
										),
									},
									MarkdownDescription: "Not within",
									Optional:            true,
									Computed:            true,
									Attributes: map[string]schema.Attribute{
										"days": schema.Int64Attribute{
											Validators: []validator.Int64{
												int64validator.ConflictsWith(
													path.MatchRelative().AtParent().AtName("hours"),
												),
												int64validator.Between(1, 65535),
											},
											MarkdownDescription: "specify time in days",
											Optional:            true,
											Computed:            true,
										},
										"hours": schema.Int64Attribute{
											Validators: []validator.Int64{
												int64validator.ConflictsWith(
													path.MatchRelative().AtParent().AtName("days"),
												),
												int64validator.Between(1, 65535),
											},
											MarkdownDescription: "specify time in hours",
											Optional:            true,
											Computed:            true,
										},
									},
								},
								"within": schema.SingleNestedAttribute{
									Validators: []validator.Object{
										objectvalidator.ConflictsWith(
											path.MatchRelative().AtParent().AtName("not_available"),
											path.MatchRelative().AtParent().AtName("not_within"),
										),
									},
									MarkdownDescription: "Within",
									Optional:            true,
									Computed:            true,
									Attributes: map[string]schema.Attribute{
										"days": schema.Int64Attribute{
											Validators: []validator.Int64{
												int64validator.ConflictsWith(
													path.MatchRelative().AtParent().AtName("hours"),
												),
												int64validator.Between(1, 65535),
											},
											MarkdownDescription: "specify time in days",
											Optional:            true,
											Computed:            true,
										},
										"hours": schema.Int64Attribute{
											Validators: []validator.Int64{
												int64validator.ConflictsWith(
													path.MatchRelative().AtParent().AtName("days"),
												),
												int64validator.Between(1, 65535),
											},
											MarkdownDescription: "specify time in hours",
											Optional:            true,
											Computed:            true,
										},
									},
								},
							},
						},
					},
				},
				"exclude_vendor": schema.BoolAttribute{
					MarkdownDescription: "Exclude vendor",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(false),
				},
				"vendor": schema.ListNestedAttribute{
					MarkdownDescription: "Vendor name",
					Optional:            true,
					Computed:            true,
					NestedObject: schema.NestedAttributeObject{
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Validators: []validator.String{
									stringvalidator.LengthAtMost(103),
								},
								MarkdownDescription: "Name",
								Required:            true,
							},
							"product": schema.ListAttribute{
								ElementType:         types.StringType,
								MarkdownDescription: "Product",
								Validators: []validator.List{
									listvalidator.ValueStringsAre(stringvalidator.LengthAtMost(1023)),
								},
								Optional: true,
							},
						},
					},
				},
			},
		},
		"disk_encryption": schema.SingleNestedAttribute{
			MarkdownDescription: "Disk encryption",
			Optional:            true,
			Computed:            true,
			Attributes: map[string]schema.Attribute{
				"criteria": schema.SingleNestedAttribute{
					MarkdownDescription: "Encryption locations",
					Optional:            true,
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"encrypted_locations": schema.ListNestedAttribute{
							MarkdownDescription: "Encrypted locations",
							Optional:            true,
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"encryption_state": schema.SingleNestedAttribute{
										MarkdownDescription: "Encryption state",
										Optional:            true,
										Attributes: map[string]schema.Attribute{
											"is": schema.StringAttribute{
												Validators: []validator.String{
													stringvalidator.ConflictsWith(
														path.MatchRelative().AtParent().AtName("is_not"),
													),
													stringvalidator.OneOf("encrypted", "unencrypted", "partial", "unknown"),
												},
												MarkdownDescription: "Is",
												Optional:            true,
											},
											"is_not": schema.StringAttribute{
												Validators: []validator.String{
													stringvalidator.ConflictsWith(
														path.MatchRelative().AtParent().AtName("is"),
													),
													stringvalidator.OneOf("encrypted", "unencrypted", "partial", "unknown"),
												},
												MarkdownDescription: "Is not",
												Optional:            true,
											},
										},
									},
									"name": schema.StringAttribute{
										Validators: []validator.String{
											stringvalidator.LengthAtMost(1023),
										},
										MarkdownDescription: "Encryption location",
										Required:            true,
									},
								},
							},
						},
						"is_installed": schema.BoolAttribute{
							MarkdownDescription: "Is Installed",
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(true),
						},
					},
				},
				"exclude_vendor": schema.BoolAttribute{
					MarkdownDescription: "Exclude vendor",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(false),
				},
				"vendor": schema.ListNestedAttribute{
					MarkdownDescription: "Vendor name",
					Optional:            true,
					Computed:            true,
					NestedObject: schema.NestedAttributeObject{
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Validators: []validator.String{
									stringvalidator.LengthAtMost(103),
								},
								MarkdownDescription: "Name",
								Required:            true,
							},
							"product": schema.ListAttribute{
								ElementType:         types.StringType,
								MarkdownDescription: "Product",
								Validators: []validator.List{
									listvalidator.ValueStringsAre(stringvalidator.LengthAtMost(1023)),
								},
								Optional: true,
							},
						},
					},
				},
			},
		},
		"firewall": schema.SingleNestedAttribute{
			MarkdownDescription: "Firewall",
			Optional:            true,
			Computed:            true,
			Attributes: map[string]schema.Attribute{
				"criteria": schema.SingleNestedAttribute{
					MarkdownDescription: "Criteria",
					Optional:            true,
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"is_enabled": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("no", "yes", "not-available"),
							},
							MarkdownDescription: "is enabled",
							Optional:            true,
							Computed:            true,
						},
						"is_installed": schema.BoolAttribute{
							MarkdownDescription: "Is Installed",
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(true),
						},
					},
				},
				"exclude_vendor": schema.BoolAttribute{
					MarkdownDescription: "Exclude vendor",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(false),
				},
				"vendor": schema.ListNestedAttribute{
					MarkdownDescription: "Vendor name",
					Optional:            true,
					Computed:            true,
					NestedObject: schema.NestedAttributeObject{
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Validators: []validator.String{
									stringvalidator.LengthAtMost(103),
								},
								MarkdownDescription: "Name",
								Required:            true,
							},
							"product": schema.ListAttribute{
								ElementType:         types.StringType,
								MarkdownDescription: "Product",
								Validators: []validator.List{
									listvalidator.ValueStringsAre(stringvalidator.LengthAtMost(1023)),
								},
								Optional: true,
							},
						},
					},
				},
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
		"host_info": schema.SingleNestedAttribute{
			MarkdownDescription: "Host info",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"criteria": schema.SingleNestedAttribute{
					MarkdownDescription: "Criteria",
					Required:            true,
					Attributes: map[string]schema.Attribute{
						"client_version": schema.SingleNestedAttribute{
							MarkdownDescription: "Client version",
							Optional:            true,
							Attributes: map[string]schema.Attribute{
								"contains": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.ConflictsWith(
											path.MatchRelative().AtParent().AtName("is"),
											path.MatchRelative().AtParent().AtName("is_not"),
										),
										stringvalidator.LengthAtMost(255),
									},
									MarkdownDescription: "Contains",
									Optional:            true,
								},
								"is": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.ConflictsWith(
											path.MatchRelative().AtParent().AtName("contains"),
											path.MatchRelative().AtParent().AtName("is_not"),
										),
										stringvalidator.LengthAtMost(255),
									},
									MarkdownDescription: "Is",
									Optional:            true,
								},
								"is_not": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.ConflictsWith(
											path.MatchRelative().AtParent().AtName("contains"),
											path.MatchRelative().AtParent().AtName("is"),
										),
										stringvalidator.LengthAtMost(255),
									},
									MarkdownDescription: "Is not",
									Optional:            true,
								},
							},
						},
						"domain": schema.SingleNestedAttribute{
							MarkdownDescription: "Domain",
							Optional:            true,
							Attributes: map[string]schema.Attribute{
								"contains": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.ConflictsWith(
											path.MatchRelative().AtParent().AtName("is"),
											path.MatchRelative().AtParent().AtName("is_not"),
										),
										stringvalidator.LengthAtMost(255),
									},
									MarkdownDescription: "Contains",
									Optional:            true,
								},
								"is": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.ConflictsWith(
											path.MatchRelative().AtParent().AtName("contains"),
											path.MatchRelative().AtParent().AtName("is_not"),
										),
										stringvalidator.LengthAtMost(255),
									},
									MarkdownDescription: "Is",
									Optional:            true,
								},
								"is_not": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.ConflictsWith(
											path.MatchRelative().AtParent().AtName("contains"),
											path.MatchRelative().AtParent().AtName("is"),
										),
										stringvalidator.LengthAtMost(255),
									},
									MarkdownDescription: "Is not",
									Optional:            true,
								},
							},
						},
						"host_id": schema.SingleNestedAttribute{
							MarkdownDescription: "Host id",
							Optional:            true,
							Attributes: map[string]schema.Attribute{
								"contains": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.ConflictsWith(
											path.MatchRelative().AtParent().AtName("is"),
											path.MatchRelative().AtParent().AtName("is_not"),
										),
										stringvalidator.LengthAtMost(255),
									},
									MarkdownDescription: "Contains",
									Optional:            true,
								},
								"is": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.ConflictsWith(
											path.MatchRelative().AtParent().AtName("contains"),
											path.MatchRelative().AtParent().AtName("is_not"),
										),
										stringvalidator.LengthAtMost(255),
									},
									MarkdownDescription: "Is",
									Optional:            true,
								},
								"is_not": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.ConflictsWith(
											path.MatchRelative().AtParent().AtName("contains"),
											path.MatchRelative().AtParent().AtName("is"),
										),
										stringvalidator.LengthAtMost(255),
									},
									MarkdownDescription: "Is not",
									Optional:            true,
								},
							},
						},
						"host_name": schema.SingleNestedAttribute{
							MarkdownDescription: "Host name",
							Optional:            true,
							Attributes: map[string]schema.Attribute{
								"contains": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.ConflictsWith(
											path.MatchRelative().AtParent().AtName("is"),
											path.MatchRelative().AtParent().AtName("is_not"),
										),
										stringvalidator.LengthAtMost(255),
									},
									MarkdownDescription: "Contains",
									Optional:            true,
								},
								"is": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.ConflictsWith(
											path.MatchRelative().AtParent().AtName("contains"),
											path.MatchRelative().AtParent().AtName("is_not"),
										),
										stringvalidator.LengthAtMost(255),
									},
									MarkdownDescription: "Is",
									Optional:            true,
								},
								"is_not": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.ConflictsWith(
											path.MatchRelative().AtParent().AtName("contains"),
											path.MatchRelative().AtParent().AtName("is"),
										),
										stringvalidator.LengthAtMost(255),
									},
									MarkdownDescription: "Is not",
									Optional:            true,
								},
							},
						},
						"managed": schema.BoolAttribute{
							MarkdownDescription: "If device is managed",
							Optional:            true,
						},
						"os": schema.SingleNestedAttribute{
							MarkdownDescription: "Os",
							Optional:            true,
							Attributes: map[string]schema.Attribute{
								"contains": schema.SingleNestedAttribute{
									MarkdownDescription: "Contains",
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"apple": schema.StringAttribute{
											MarkdownDescription: "Apple",
											Optional:            true,
										},
										"google": schema.StringAttribute{
											MarkdownDescription: "Google",
											Optional:            true,
										},
										"linux": schema.StringAttribute{
											MarkdownDescription: "Linux",
											Optional:            true,
										},
										"microsoft": schema.StringAttribute{
											MarkdownDescription: "Microsoft",
											Optional:            true,
										},
										"other": schema.StringAttribute{
											MarkdownDescription: "Other",
											Optional:            true,
										},
									},
								},
							},
						},
						"serial_number": schema.SingleNestedAttribute{
							MarkdownDescription: "Serial number",
							Optional:            true,
							Attributes: map[string]schema.Attribute{
								"contains": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.ConflictsWith(
											path.MatchRelative().AtParent().AtName("is"),
											path.MatchRelative().AtParent().AtName("is_not"),
										),
										stringvalidator.LengthAtMost(255),
									},
									MarkdownDescription: "Contains",
									Optional:            true,
								},
								"is": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.ConflictsWith(
											path.MatchRelative().AtParent().AtName("contains"),
											path.MatchRelative().AtParent().AtName("is_not"),
										),
										stringvalidator.LengthAtMost(255),
									},
									MarkdownDescription: "Is",
									Optional:            true,
								},
								"is_not": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.ConflictsWith(
											path.MatchRelative().AtParent().AtName("contains"),
											path.MatchRelative().AtParent().AtName("is"),
										),
										stringvalidator.LengthAtMost(255),
									},
									MarkdownDescription: "Is not",
									Optional:            true,
								},
							},
						},
					},
				},
			},
		},
		"id": schema.StringAttribute{
			MarkdownDescription: "UUID of the resource",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"mobile_device": schema.SingleNestedAttribute{
			MarkdownDescription: "Mobile device",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"criteria": schema.SingleNestedAttribute{
					MarkdownDescription: "Criteria",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"applications": schema.SingleNestedAttribute{
							MarkdownDescription: "Applications",
							Optional:            true,
							Attributes: map[string]schema.Attribute{
								"has_malware": schema.SingleNestedAttribute{
									MarkdownDescription: "Has malware",
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"no": schema.SingleNestedAttribute{
											Validators: []validator.Object{
												objectvalidator.ConflictsWith(
													path.MatchRelative().AtParent().AtName("yes"),
												),
											},
											MarkdownDescription: "No",
											Optional:            true,
											Attributes:          map[string]schema.Attribute{},
										},
										"yes": schema.SingleNestedAttribute{
											Validators: []validator.Object{
												objectvalidator.ConflictsWith(
													path.MatchRelative().AtParent().AtName("no"),
												),
											},
											MarkdownDescription: "Yes",
											Optional:            true,
											Attributes: map[string]schema.Attribute{
												"excludes": schema.ListNestedAttribute{
													MarkdownDescription: "Excludes",
													Optional:            true,
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"hash": schema.StringAttribute{
																Validators: []validator.String{
																	stringvalidator.LengthAtMost(1024),
																	stringvalidator.RegexMatches(regexp.MustCompile(".*"), "pattern must match "+".*"),
																},
																MarkdownDescription: "application hash",
																Optional:            true,
															},
															"name": schema.StringAttribute{
																Validators: []validator.String{
																	stringvalidator.LengthAtMost(31),
																},
																MarkdownDescription: "Name",
																Required:            true,
															},
															"package": schema.StringAttribute{
																Validators: []validator.String{
																	stringvalidator.LengthAtMost(1024),
																	stringvalidator.RegexMatches(regexp.MustCompile(".*"), "pattern must match "+".*"),
																},
																MarkdownDescription: "application package name",
																Optional:            true,
															},
														},
													},
												},
											},
										},
									},
								},
								"has_unmanaged_app": schema.BoolAttribute{
									MarkdownDescription: "Has apps that are not managed",
									Optional:            true,
								},
								"includes": schema.ListNestedAttribute{
									MarkdownDescription: "Includes",
									Optional:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"hash": schema.StringAttribute{
												Validators: []validator.String{
													stringvalidator.LengthAtMost(1024),
													stringvalidator.RegexMatches(regexp.MustCompile(".*"), "pattern must match "+".*"),
												},
												MarkdownDescription: "application hash",
												Optional:            true,
											},
											"name": schema.StringAttribute{
												Validators: []validator.String{
													stringvalidator.LengthAtMost(31),
												},
												MarkdownDescription: "Name",
												Required:            true,
											},
											"package": schema.StringAttribute{
												Validators: []validator.String{
													stringvalidator.LengthAtMost(1024),
													stringvalidator.RegexMatches(regexp.MustCompile(".*"), "pattern must match "+".*"),
												},
												MarkdownDescription: "application package name",
												Optional:            true,
											},
										},
									},
								},
							},
						},
						"disk_encrypted": schema.BoolAttribute{
							MarkdownDescription: "If device's disk is encrypted",
							Optional:            true,
						},
						"imei": schema.SingleNestedAttribute{
							MarkdownDescription: "Imei",
							Optional:            true,
							Attributes: map[string]schema.Attribute{
								"contains": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.ConflictsWith(
											path.MatchRelative().AtParent().AtName("is"),
											path.MatchRelative().AtParent().AtName("is_not"),
										),
										stringvalidator.LengthAtMost(255),
									},
									MarkdownDescription: "Contains",
									Optional:            true,
								},
								"is": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.ConflictsWith(
											path.MatchRelative().AtParent().AtName("contains"),
											path.MatchRelative().AtParent().AtName("is_not"),
										),
										stringvalidator.LengthAtMost(255),
									},
									MarkdownDescription: "Is",
									Optional:            true,
								},
								"is_not": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.ConflictsWith(
											path.MatchRelative().AtParent().AtName("contains"),
											path.MatchRelative().AtParent().AtName("is"),
										),
										stringvalidator.LengthAtMost(255),
									},
									MarkdownDescription: "Is not",
									Optional:            true,
								},
							},
						},
						"jailbroken": schema.BoolAttribute{
							MarkdownDescription: "If device is by rooted/jailbroken",
							Optional:            true,
						},
						"last_checkin_time": schema.SingleNestedAttribute{
							MarkdownDescription: "Last checkin time",
							Optional:            true,
							Attributes: map[string]schema.Attribute{
								"not_within": schema.SingleNestedAttribute{
									Validators: []validator.Object{
										objectvalidator.ConflictsWith(
											path.MatchRelative().AtParent().AtName("within"),
										),
									},
									MarkdownDescription: "Not within",
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"days": schema.Int64Attribute{
											Validators: []validator.Int64{
												int64validator.Between(1, 365),
											},
											MarkdownDescription: "specify time in days",
											Required:            true,
										},
									},
								},
								"within": schema.SingleNestedAttribute{
									Validators: []validator.Object{
										objectvalidator.ConflictsWith(
											path.MatchRelative().AtParent().AtName("not_within"),
										),
									},
									MarkdownDescription: "Within",
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"days": schema.Int64Attribute{
											Validators: []validator.Int64{
												int64validator.Between(1, 365),
											},
											MarkdownDescription: "specify time in days",
											Required:            true,
										},
									},
								},
							},
						},
						"model": schema.SingleNestedAttribute{
							MarkdownDescription: "Model",
							Optional:            true,
							Attributes: map[string]schema.Attribute{
								"contains": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.ConflictsWith(
											path.MatchRelative().AtParent().AtName("is"),
											path.MatchRelative().AtParent().AtName("is_not"),
										),
										stringvalidator.LengthAtMost(255),
									},
									MarkdownDescription: "Contains",
									Optional:            true,
								},
								"is": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.ConflictsWith(
											path.MatchRelative().AtParent().AtName("contains"),
											path.MatchRelative().AtParent().AtName("is_not"),
										),
										stringvalidator.LengthAtMost(255),
									},
									MarkdownDescription: "Is",
									Optional:            true,
								},
								"is_not": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.ConflictsWith(
											path.MatchRelative().AtParent().AtName("contains"),
											path.MatchRelative().AtParent().AtName("is"),
										),
										stringvalidator.LengthAtMost(255),
									},
									MarkdownDescription: "Is not",
									Optional:            true,
								},
							},
						},
						"passcode_set": schema.BoolAttribute{
							MarkdownDescription: "If device's passcode is present",
							Optional:            true,
						},
						"phone_number": schema.SingleNestedAttribute{
							MarkdownDescription: "Phone number",
							Optional:            true,
							Attributes: map[string]schema.Attribute{
								"contains": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.ConflictsWith(
											path.MatchRelative().AtParent().AtName("is"),
											path.MatchRelative().AtParent().AtName("is_not"),
										),
										stringvalidator.LengthAtMost(255),
									},
									MarkdownDescription: "Contains",
									Optional:            true,
								},
								"is": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.ConflictsWith(
											path.MatchRelative().AtParent().AtName("contains"),
											path.MatchRelative().AtParent().AtName("is_not"),
										),
										stringvalidator.LengthAtMost(255),
									},
									MarkdownDescription: "Is",
									Optional:            true,
								},
								"is_not": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.ConflictsWith(
											path.MatchRelative().AtParent().AtName("contains"),
											path.MatchRelative().AtParent().AtName("is"),
										),
										stringvalidator.LengthAtMost(255),
									},
									MarkdownDescription: "Is not",
									Optional:            true,
								},
							},
						},
						"tag": schema.SingleNestedAttribute{
							MarkdownDescription: "Tag",
							Optional:            true,
							Attributes: map[string]schema.Attribute{
								"contains": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.ConflictsWith(
											path.MatchRelative().AtParent().AtName("is"),
											path.MatchRelative().AtParent().AtName("is_not"),
										),
										stringvalidator.LengthAtMost(255),
									},
									MarkdownDescription: "Contains",
									Optional:            true,
								},
								"is": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.ConflictsWith(
											path.MatchRelative().AtParent().AtName("contains"),
											path.MatchRelative().AtParent().AtName("is_not"),
										),
										stringvalidator.LengthAtMost(255),
									},
									MarkdownDescription: "Is",
									Optional:            true,
								},
								"is_not": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.ConflictsWith(
											path.MatchRelative().AtParent().AtName("contains"),
											path.MatchRelative().AtParent().AtName("is"),
										),
										stringvalidator.LengthAtMost(255),
									},
									MarkdownDescription: "Is not",
									Optional:            true,
								},
							},
						},
					},
				},
			},
		},
		"name": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(31),
				stringvalidator.RegexMatches(regexp.MustCompile("^[ a-zA-Z\\d.\\-_]+$"), "pattern must match "+"^[ a-zA-Z\\d.\\-_]+$"),
			},
			MarkdownDescription: "The name of the HIP object",
			Required:            true,
		},
		"network_info": schema.SingleNestedAttribute{
			MarkdownDescription: "Network info",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"criteria": schema.SingleNestedAttribute{
					MarkdownDescription: "Criteria",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"network": schema.SingleNestedAttribute{
							MarkdownDescription: "Network",
							Optional:            true,
							Attributes: map[string]schema.Attribute{
								"is": schema.SingleNestedAttribute{
									Validators: []validator.Object{
										objectvalidator.ConflictsWith(
											path.MatchRelative().AtParent().AtName("is_not"),
										),
									},
									MarkdownDescription: "Is",
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"mobile": schema.SingleNestedAttribute{
											Validators: []validator.Object{
												objectvalidator.ConflictsWith(
													path.MatchRelative().AtParent().AtName("unknown"),
													path.MatchRelative().AtParent().AtName("wifi"),
												),
											},
											MarkdownDescription: "Mobile",
											Optional:            true,
											Attributes: map[string]schema.Attribute{
												"carrier": schema.StringAttribute{
													Validators: []validator.String{
														stringvalidator.LengthAtMost(1023),
														stringvalidator.RegexMatches(regexp.MustCompile(".*"), "pattern must match "+".*"),
													},
													MarkdownDescription: "Carrier",
													Optional:            true,
												},
											},
										},
										"unknown": schema.SingleNestedAttribute{
											Validators: []validator.Object{
												objectvalidator.ConflictsWith(
													path.MatchRelative().AtParent().AtName("mobile"),
													path.MatchRelative().AtParent().AtName("wifi"),
												),
											},
											MarkdownDescription: "Unknown",
											Optional:            true,
											Attributes:          map[string]schema.Attribute{},
										},
										"wifi": schema.SingleNestedAttribute{
											Validators: []validator.Object{
												objectvalidator.ConflictsWith(
													path.MatchRelative().AtParent().AtName("mobile"),
													path.MatchRelative().AtParent().AtName("unknown"),
												),
											},
											MarkdownDescription: "Wifi",
											Optional:            true,
											Attributes: map[string]schema.Attribute{
												"ssid": schema.StringAttribute{
													Validators: []validator.String{
														stringvalidator.LengthAtMost(1023),
														stringvalidator.RegexMatches(regexp.MustCompile(".*"), "pattern must match "+".*"),
													},
													MarkdownDescription: "SSID",
													Optional:            true,
												},
											},
										},
									},
								},
								"is_not": schema.SingleNestedAttribute{
									Validators: []validator.Object{
										objectvalidator.ConflictsWith(
											path.MatchRelative().AtParent().AtName("is"),
										),
									},
									MarkdownDescription: "Is not",
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"ethernet": schema.SingleNestedAttribute{
											Validators: []validator.Object{
												objectvalidator.ConflictsWith(
													path.MatchRelative().AtParent().AtName("mobile"),
													path.MatchRelative().AtParent().AtName("unknown"),
													path.MatchRelative().AtParent().AtName("wifi"),
												),
											},
											MarkdownDescription: "Ethernet",
											Optional:            true,
											Attributes:          map[string]schema.Attribute{},
										},
										"mobile": schema.SingleNestedAttribute{
											Validators: []validator.Object{
												objectvalidator.ConflictsWith(
													path.MatchRelative().AtParent().AtName("ethernet"),
													path.MatchRelative().AtParent().AtName("unknown"),
													path.MatchRelative().AtParent().AtName("wifi"),
												),
											},
											MarkdownDescription: "Mobile",
											Optional:            true,
											Attributes: map[string]schema.Attribute{
												"carrier": schema.StringAttribute{
													Validators: []validator.String{
														stringvalidator.LengthAtMost(1023),
														stringvalidator.RegexMatches(regexp.MustCompile(".*"), "pattern must match "+".*"),
													},
													MarkdownDescription: "Carrier",
													Optional:            true,
												},
											},
										},
										"unknown": schema.SingleNestedAttribute{
											Validators: []validator.Object{
												objectvalidator.ConflictsWith(
													path.MatchRelative().AtParent().AtName("ethernet"),
													path.MatchRelative().AtParent().AtName("mobile"),
													path.MatchRelative().AtParent().AtName("wifi"),
												),
											},
											MarkdownDescription: "Unknown",
											Optional:            true,
											Attributes:          map[string]schema.Attribute{},
										},
										"wifi": schema.SingleNestedAttribute{
											Validators: []validator.Object{
												objectvalidator.ConflictsWith(
													path.MatchRelative().AtParent().AtName("ethernet"),
													path.MatchRelative().AtParent().AtName("mobile"),
													path.MatchRelative().AtParent().AtName("unknown"),
												),
											},
											MarkdownDescription: "Wifi",
											Optional:            true,
											Attributes: map[string]schema.Attribute{
												"ssid": schema.StringAttribute{
													Validators: []validator.String{
														stringvalidator.LengthAtMost(1023),
														stringvalidator.RegexMatches(regexp.MustCompile(".*"), "pattern must match "+".*"),
													},
													MarkdownDescription: "SSID",
													Optional:            true,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"patch_management": schema.SingleNestedAttribute{
			MarkdownDescription: "Patch management",
			Optional:            true,
			Computed:            true,
			Attributes: map[string]schema.Attribute{
				"criteria": schema.SingleNestedAttribute{
					MarkdownDescription: "Criteria",
					Optional:            true,
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"is_enabled": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("no", "yes", "not-available"),
							},
							MarkdownDescription: "is enabled",
							Optional:            true,
							Computed:            true,
						},
						"is_installed": schema.BoolAttribute{
							MarkdownDescription: "Is Installed",
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(true),
						},
						"missing_patches": schema.SingleNestedAttribute{
							MarkdownDescription: "Missing patches",
							Optional:            true,
							Computed:            true,
							Attributes: map[string]schema.Attribute{
								"check": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.OneOf("has-any", "has-none", "has-all"),
									},
									MarkdownDescription: "Check",
									Required:            true,
								},
								"patches": schema.ListAttribute{
									ElementType:         types.StringType,
									MarkdownDescription: "Patches",
									Validators: []validator.List{
										listvalidator.ValueStringsAre(stringvalidator.LengthAtMost(1023)),
									},
									Optional: true,
									Computed: true,
								},
								"severity": schema.SingleNestedAttribute{
									MarkdownDescription: "Severity",
									Optional:            true,
									Computed:            true,
									Attributes: map[string]schema.Attribute{
										"greater_equal": schema.Int64Attribute{
											Validators: []validator.Int64{
												int64validator.ConflictsWith(
													path.MatchRelative().AtParent().AtName("greater_than"),
													path.MatchRelative().AtParent().AtName("is"),
													path.MatchRelative().AtParent().AtName("is_not"),
													path.MatchRelative().AtParent().AtName("less_equal"),
													path.MatchRelative().AtParent().AtName("less_than"),
												),
												int64validator.Between(0, 100000),
											},
											MarkdownDescription: "Greater equal",
											Optional:            true,
											Computed:            true,
										},
										"greater_than": schema.Int64Attribute{
											Validators: []validator.Int64{
												int64validator.ConflictsWith(
													path.MatchRelative().AtParent().AtName("greater_equal"),
													path.MatchRelative().AtParent().AtName("is"),
													path.MatchRelative().AtParent().AtName("is_not"),
													path.MatchRelative().AtParent().AtName("less_equal"),
													path.MatchRelative().AtParent().AtName("less_than"),
												),
												int64validator.Between(0, 100000),
											},
											MarkdownDescription: "Greater than",
											Optional:            true,
											Computed:            true,
										},
										"is": schema.Int64Attribute{
											Validators: []validator.Int64{
												int64validator.ConflictsWith(
													path.MatchRelative().AtParent().AtName("greater_equal"),
													path.MatchRelative().AtParent().AtName("greater_than"),
													path.MatchRelative().AtParent().AtName("is_not"),
													path.MatchRelative().AtParent().AtName("less_equal"),
													path.MatchRelative().AtParent().AtName("less_than"),
												),
												int64validator.Between(0, 100000),
											},
											MarkdownDescription: "Is",
											Optional:            true,
											Computed:            true,
										},
										"is_not": schema.Int64Attribute{
											Validators: []validator.Int64{
												int64validator.ConflictsWith(
													path.MatchRelative().AtParent().AtName("greater_equal"),
													path.MatchRelative().AtParent().AtName("greater_than"),
													path.MatchRelative().AtParent().AtName("is"),
													path.MatchRelative().AtParent().AtName("less_equal"),
													path.MatchRelative().AtParent().AtName("less_than"),
												),
												int64validator.Between(0, 100000),
											},
											MarkdownDescription: "Is not",
											Optional:            true,
											Computed:            true,
										},
										"less_equal": schema.Int64Attribute{
											Validators: []validator.Int64{
												int64validator.ConflictsWith(
													path.MatchRelative().AtParent().AtName("greater_equal"),
													path.MatchRelative().AtParent().AtName("greater_than"),
													path.MatchRelative().AtParent().AtName("is"),
													path.MatchRelative().AtParent().AtName("is_not"),
													path.MatchRelative().AtParent().AtName("less_than"),
												),
												int64validator.Between(0, 100000),
											},
											MarkdownDescription: "Less equal",
											Optional:            true,
											Computed:            true,
										},
										"less_than": schema.Int64Attribute{
											Validators: []validator.Int64{
												int64validator.ConflictsWith(
													path.MatchRelative().AtParent().AtName("greater_equal"),
													path.MatchRelative().AtParent().AtName("greater_than"),
													path.MatchRelative().AtParent().AtName("is"),
													path.MatchRelative().AtParent().AtName("is_not"),
													path.MatchRelative().AtParent().AtName("less_equal"),
												),
												int64validator.Between(0, 100000),
											},
											MarkdownDescription: "Less than",
											Optional:            true,
											Computed:            true,
										},
									},
								},
							},
						},
					},
				},
				"exclude_vendor": schema.BoolAttribute{
					MarkdownDescription: "Exclude vendor",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(false),
				},
				"vendor": schema.ListNestedAttribute{
					MarkdownDescription: "Vendor name",
					Optional:            true,
					Computed:            true,
					NestedObject: schema.NestedAttributeObject{
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Validators: []validator.String{
									stringvalidator.LengthAtMost(103),
								},
								MarkdownDescription: "Name",
								Required:            true,
							},
							"product": schema.ListAttribute{
								ElementType:         types.StringType,
								MarkdownDescription: "Product name",
								Validators: []validator.List{
									listvalidator.ValueStringsAre(stringvalidator.LengthAtMost(1023)),
								},
								Optional: true,
							},
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
		"tfid": schema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
	},
}

// HipObjectsDataSourceSchema defines the schema for HipObjects data source
var HipObjectsDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "HipObject data source",
	Attributes: map[string]dsschema.Attribute{
		"anti_malware": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Anti malware",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"criteria": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Criteria",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"is_installed": dsschema.BoolAttribute{
							MarkdownDescription: "Is Installed",
							Computed:            true,
						},
						"last_scan_time": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Last scan time",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"not_available": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Not available",
									Computed:            true,
									Attributes:          map[string]dsschema.Attribute{},
								},
								"not_within": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Not within",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"days": dsschema.Int64Attribute{
											MarkdownDescription: "specify time in days",
											Computed:            true,
										},
										"hours": dsschema.Int64Attribute{
											MarkdownDescription: "specify time in hours",
											Computed:            true,
										},
									},
								},
								"within": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Within",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"days": dsschema.Int64Attribute{
											MarkdownDescription: "specify time in days",
											Computed:            true,
										},
										"hours": dsschema.Int64Attribute{
											MarkdownDescription: "specify time in hours",
											Computed:            true,
										},
									},
								},
							},
						},
						"product_version": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Product version",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"contains": dsschema.StringAttribute{
									MarkdownDescription: "Contains",
									Computed:            true,
								},
								"greater_equal": dsschema.StringAttribute{
									MarkdownDescription: "Greater equal",
									Computed:            true,
								},
								"greater_than": dsschema.StringAttribute{
									MarkdownDescription: "Greater than",
									Computed:            true,
								},
								"is": dsschema.StringAttribute{
									MarkdownDescription: "Is",
									Computed:            true,
								},
								"is_not": dsschema.StringAttribute{
									MarkdownDescription: "Is not",
									Computed:            true,
								},
								"less_equal": dsschema.StringAttribute{
									MarkdownDescription: "Less equal",
									Computed:            true,
								},
								"less_than": dsschema.StringAttribute{
									MarkdownDescription: "Less than",
									Computed:            true,
								},
								"not_within": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Not within",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"versions": dsschema.Int64Attribute{
											MarkdownDescription: "versions range",
											Computed:            true,
										},
									},
								},
								"within": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Within",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"versions": dsschema.Int64Attribute{
											MarkdownDescription: "versions range",
											Computed:            true,
										},
									},
								},
							},
						},
						"real_time_protection": dsschema.StringAttribute{
							MarkdownDescription: "real time protection",
							Computed:            true,
						},
						"virdef_version": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Virdef version",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"not_within": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Not within",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"days": dsschema.Int64Attribute{
											MarkdownDescription: "specify time in days",
											Computed:            true,
										},
										"versions": dsschema.Int64Attribute{
											MarkdownDescription: "specify versions range",
											Computed:            true,
										},
									},
								},
								"within": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Within",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"days": dsschema.Int64Attribute{
											MarkdownDescription: "specify time in days",
											Computed:            true,
										},
										"versions": dsschema.Int64Attribute{
											MarkdownDescription: "specify versions range",
											Computed:            true,
										},
									},
								},
							},
						},
					},
				},
				"exclude_vendor": dsschema.BoolAttribute{
					MarkdownDescription: "Exclude vendor",
					Computed:            true,
				},
				"vendor": dsschema.ListNestedAttribute{
					MarkdownDescription: "Vendor name",
					Computed:            true,
					NestedObject: dsschema.NestedAttributeObject{
						Attributes: map[string]dsschema.Attribute{
							"name": dsschema.StringAttribute{
								MarkdownDescription: "Name",
								Computed:            true,
							},
							"product": dsschema.ListAttribute{
								ElementType:         types.StringType,
								MarkdownDescription: "Product",
								Computed:            true,
							},
						},
					},
				},
			},
		},
		"certificate": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Certificate",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"criteria": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Criteria",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"certificate_attributes": dsschema.ListNestedAttribute{
							MarkdownDescription: "Certificate attributes",
							Computed:            true,
							NestedObject: dsschema.NestedAttributeObject{
								Attributes: map[string]dsschema.Attribute{
									"name": dsschema.StringAttribute{
										MarkdownDescription: "Attribute Name",
										Computed:            true,
									},
									"value": dsschema.StringAttribute{
										MarkdownDescription: "Key value",
										Computed:            true,
									},
								},
							},
						},
						"certificate_profile": dsschema.StringAttribute{
							MarkdownDescription: "Profile for authenticating client certificates",
							Computed:            true,
						},
					},
				},
			},
		},
		"custom_checks": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Custom checks",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"criteria": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Criteria",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"plist": dsschema.ListNestedAttribute{
							MarkdownDescription: "Plist",
							Computed:            true,
							NestedObject: dsschema.NestedAttributeObject{
								Attributes: map[string]dsschema.Attribute{
									"key": dsschema.ListNestedAttribute{
										MarkdownDescription: "Key",
										Computed:            true,
										NestedObject: dsschema.NestedAttributeObject{
											Attributes: map[string]dsschema.Attribute{
												"name": dsschema.StringAttribute{
													MarkdownDescription: "Key name",
													Computed:            true,
												},
												"negate": dsschema.BoolAttribute{
													MarkdownDescription: "Value does not exist or match specified value data",
													Computed:            true,
												},
												"value": dsschema.StringAttribute{
													MarkdownDescription: "Key value",
													Computed:            true,
												},
											},
										},
									},
									"name": dsschema.StringAttribute{
										MarkdownDescription: "Preference list",
										Computed:            true,
									},
									"negate": dsschema.BoolAttribute{
										MarkdownDescription: "Plist does not exist",
										Computed:            true,
									},
								},
							},
						},
						"process_list": dsschema.ListNestedAttribute{
							MarkdownDescription: "Process list",
							Computed:            true,
							NestedObject: dsschema.NestedAttributeObject{
								Attributes: map[string]dsschema.Attribute{
									"name": dsschema.StringAttribute{
										MarkdownDescription: "Process Name",
										Computed:            true,
									},
									"running": dsschema.BoolAttribute{
										MarkdownDescription: "Running",
										Computed:            true,
									},
								},
							},
						},
						"registry_key": dsschema.ListNestedAttribute{
							MarkdownDescription: "Registry key",
							Computed:            true,
							NestedObject: dsschema.NestedAttributeObject{
								Attributes: map[string]dsschema.Attribute{
									"default_value_data": dsschema.StringAttribute{
										MarkdownDescription: "Registry key default value data",
										Computed:            true,
									},
									"name": dsschema.StringAttribute{
										MarkdownDescription: "Registry key",
										Computed:            true,
									},
									"negate": dsschema.BoolAttribute{
										MarkdownDescription: "Key does not exist or match specified value data",
										Computed:            true,
									},
									"registry_value": dsschema.ListNestedAttribute{
										MarkdownDescription: "Registry value",
										Computed:            true,
										NestedObject: dsschema.NestedAttributeObject{
											Attributes: map[string]dsschema.Attribute{
												"name": dsschema.StringAttribute{
													MarkdownDescription: "Registry value name",
													Computed:            true,
												},
												"negate": dsschema.BoolAttribute{
													MarkdownDescription: "Value does not exist or match specified value data",
													Computed:            true,
												},
												"value_data": dsschema.StringAttribute{
													MarkdownDescription: "Registry value data",
													Computed:            true,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"data_loss_prevention": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Data loss prevention",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"criteria": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Criteria",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"is_enabled": dsschema.StringAttribute{
							MarkdownDescription: "is enabled",
							Computed:            true,
						},
						"is_installed": dsschema.BoolAttribute{
							MarkdownDescription: "Is Installed",
							Computed:            true,
						},
					},
				},
				"exclude_vendor": dsschema.BoolAttribute{
					MarkdownDescription: "Exclude vendor",
					Computed:            true,
				},
				"vendor": dsschema.ListNestedAttribute{
					MarkdownDescription: "Vendor name",
					Computed:            true,
					NestedObject: dsschema.NestedAttributeObject{
						Attributes: map[string]dsschema.Attribute{
							"name": dsschema.StringAttribute{
								MarkdownDescription: "Name",
								Computed:            true,
							},
							"product": dsschema.ListAttribute{
								ElementType:         types.StringType,
								MarkdownDescription: "Product name",
								Computed:            true,
							},
						},
					},
				},
			},
		},
		"description": dsschema.StringAttribute{
			MarkdownDescription: "Description",
			Computed:            true,
		},
		"device": dsschema.StringAttribute{
			MarkdownDescription: "The device in which the resource is defined",
			Computed:            true,
		},
		"disk_backup": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Disk backup",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"criteria": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Criteria",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"is_installed": dsschema.BoolAttribute{
							MarkdownDescription: "Is Installed",
							Computed:            true,
						},
						"last_backup_time": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Last backup time",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"not_available": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Not available",
									Computed:            true,
									Attributes:          map[string]dsschema.Attribute{},
								},
								"not_within": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Not within",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"days": dsschema.Int64Attribute{
											MarkdownDescription: "specify time in days",
											Computed:            true,
										},
										"hours": dsschema.Int64Attribute{
											MarkdownDescription: "specify time in hours",
											Computed:            true,
										},
									},
								},
								"within": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Within",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"days": dsschema.Int64Attribute{
											MarkdownDescription: "specify time in days",
											Computed:            true,
										},
										"hours": dsschema.Int64Attribute{
											MarkdownDescription: "specify time in hours",
											Computed:            true,
										},
									},
								},
							},
						},
					},
				},
				"exclude_vendor": dsschema.BoolAttribute{
					MarkdownDescription: "Exclude vendor",
					Computed:            true,
				},
				"vendor": dsschema.ListNestedAttribute{
					MarkdownDescription: "Vendor name",
					Computed:            true,
					NestedObject: dsschema.NestedAttributeObject{
						Attributes: map[string]dsschema.Attribute{
							"name": dsschema.StringAttribute{
								MarkdownDescription: "Name",
								Computed:            true,
							},
							"product": dsschema.ListAttribute{
								ElementType:         types.StringType,
								MarkdownDescription: "Product",
								Computed:            true,
							},
						},
					},
				},
			},
		},
		"disk_encryption": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Disk encryption",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"criteria": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Encryption locations",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"encrypted_locations": dsschema.ListNestedAttribute{
							MarkdownDescription: "Encrypted locations",
							Computed:            true,
							NestedObject: dsschema.NestedAttributeObject{
								Attributes: map[string]dsschema.Attribute{
									"encryption_state": dsschema.SingleNestedAttribute{
										MarkdownDescription: "Encryption state",
										Computed:            true,
										Attributes: map[string]dsschema.Attribute{
											"is": dsschema.StringAttribute{
												MarkdownDescription: "Is",
												Computed:            true,
											},
											"is_not": dsschema.StringAttribute{
												MarkdownDescription: "Is not",
												Computed:            true,
											},
										},
									},
									"name": dsschema.StringAttribute{
										MarkdownDescription: "Encryption location",
										Computed:            true,
									},
								},
							},
						},
						"is_installed": dsschema.BoolAttribute{
							MarkdownDescription: "Is Installed",
							Computed:            true,
						},
					},
				},
				"exclude_vendor": dsschema.BoolAttribute{
					MarkdownDescription: "Exclude vendor",
					Computed:            true,
				},
				"vendor": dsschema.ListNestedAttribute{
					MarkdownDescription: "Vendor name",
					Computed:            true,
					NestedObject: dsschema.NestedAttributeObject{
						Attributes: map[string]dsschema.Attribute{
							"name": dsschema.StringAttribute{
								MarkdownDescription: "Name",
								Computed:            true,
							},
							"product": dsschema.ListAttribute{
								ElementType:         types.StringType,
								MarkdownDescription: "Product",
								Computed:            true,
							},
						},
					},
				},
			},
		},
		"firewall": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Firewall",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"criteria": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Criteria",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"is_enabled": dsschema.StringAttribute{
							MarkdownDescription: "is enabled",
							Computed:            true,
						},
						"is_installed": dsschema.BoolAttribute{
							MarkdownDescription: "Is Installed",
							Computed:            true,
						},
					},
				},
				"exclude_vendor": dsschema.BoolAttribute{
					MarkdownDescription: "Exclude vendor",
					Computed:            true,
				},
				"vendor": dsschema.ListNestedAttribute{
					MarkdownDescription: "Vendor name",
					Computed:            true,
					NestedObject: dsschema.NestedAttributeObject{
						Attributes: map[string]dsschema.Attribute{
							"name": dsschema.StringAttribute{
								MarkdownDescription: "Name",
								Computed:            true,
							},
							"product": dsschema.ListAttribute{
								ElementType:         types.StringType,
								MarkdownDescription: "Product",
								Computed:            true,
							},
						},
					},
				},
			},
		},
		"folder": dsschema.StringAttribute{
			MarkdownDescription: "The folder in which the resource is defined",
			Computed:            true,
		},
		"host_info": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Host info",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"criteria": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Criteria",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"client_version": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Client version",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"contains": dsschema.StringAttribute{
									MarkdownDescription: "Contains",
									Computed:            true,
								},
								"is": dsschema.StringAttribute{
									MarkdownDescription: "Is",
									Computed:            true,
								},
								"is_not": dsschema.StringAttribute{
									MarkdownDescription: "Is not",
									Computed:            true,
								},
							},
						},
						"domain": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Domain",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"contains": dsschema.StringAttribute{
									MarkdownDescription: "Contains",
									Computed:            true,
								},
								"is": dsschema.StringAttribute{
									MarkdownDescription: "Is",
									Computed:            true,
								},
								"is_not": dsschema.StringAttribute{
									MarkdownDescription: "Is not",
									Computed:            true,
								},
							},
						},
						"host_id": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Host id",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"contains": dsschema.StringAttribute{
									MarkdownDescription: "Contains",
									Computed:            true,
								},
								"is": dsschema.StringAttribute{
									MarkdownDescription: "Is",
									Computed:            true,
								},
								"is_not": dsschema.StringAttribute{
									MarkdownDescription: "Is not",
									Computed:            true,
								},
							},
						},
						"host_name": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Host name",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"contains": dsschema.StringAttribute{
									MarkdownDescription: "Contains",
									Computed:            true,
								},
								"is": dsschema.StringAttribute{
									MarkdownDescription: "Is",
									Computed:            true,
								},
								"is_not": dsschema.StringAttribute{
									MarkdownDescription: "Is not",
									Computed:            true,
								},
							},
						},
						"managed": dsschema.BoolAttribute{
							MarkdownDescription: "If device is managed",
							Computed:            true,
						},
						"os": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Os",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"contains": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Contains",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"apple": dsschema.StringAttribute{
											MarkdownDescription: "Apple",
											Computed:            true,
										},
										"google": dsschema.StringAttribute{
											MarkdownDescription: "Google",
											Computed:            true,
										},
										"linux": dsschema.StringAttribute{
											MarkdownDescription: "Linux",
											Computed:            true,
										},
										"microsoft": dsschema.StringAttribute{
											MarkdownDescription: "Microsoft",
											Computed:            true,
										},
										"other": dsschema.StringAttribute{
											MarkdownDescription: "Other",
											Computed:            true,
										},
									},
								},
							},
						},
						"serial_number": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Serial number",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"contains": dsschema.StringAttribute{
									MarkdownDescription: "Contains",
									Computed:            true,
								},
								"is": dsschema.StringAttribute{
									MarkdownDescription: "Is",
									Computed:            true,
								},
								"is_not": dsschema.StringAttribute{
									MarkdownDescription: "Is not",
									Computed:            true,
								},
							},
						},
					},
				},
			},
		},
		"id": dsschema.StringAttribute{
			MarkdownDescription: "UUID of the resource",
			Required:            true,
		},
		"mobile_device": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Mobile device",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"criteria": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Criteria",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"applications": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Applications",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"has_malware": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Has malware",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"no": dsschema.SingleNestedAttribute{
											MarkdownDescription: "No",
											Computed:            true,
											Attributes:          map[string]dsschema.Attribute{},
										},
										"yes": dsschema.SingleNestedAttribute{
											MarkdownDescription: "Yes",
											Computed:            true,
											Attributes: map[string]dsschema.Attribute{
												"excludes": dsschema.ListNestedAttribute{
													MarkdownDescription: "Excludes",
													Computed:            true,
													NestedObject: dsschema.NestedAttributeObject{
														Attributes: map[string]dsschema.Attribute{
															"hash": dsschema.StringAttribute{
																MarkdownDescription: "application hash",
																Computed:            true,
															},
															"name": dsschema.StringAttribute{
																MarkdownDescription: "Name",
																Computed:            true,
															},
															"package": dsschema.StringAttribute{
																MarkdownDescription: "application package name",
																Computed:            true,
															},
														},
													},
												},
											},
										},
									},
								},
								"has_unmanaged_app": dsschema.BoolAttribute{
									MarkdownDescription: "Has apps that are not managed",
									Computed:            true,
								},
								"includes": dsschema.ListNestedAttribute{
									MarkdownDescription: "Includes",
									Computed:            true,
									NestedObject: dsschema.NestedAttributeObject{
										Attributes: map[string]dsschema.Attribute{
											"hash": dsschema.StringAttribute{
												MarkdownDescription: "application hash",
												Computed:            true,
											},
											"name": dsschema.StringAttribute{
												MarkdownDescription: "Name",
												Computed:            true,
											},
											"package": dsschema.StringAttribute{
												MarkdownDescription: "application package name",
												Computed:            true,
											},
										},
									},
								},
							},
						},
						"disk_encrypted": dsschema.BoolAttribute{
							MarkdownDescription: "If device's disk is encrypted",
							Computed:            true,
						},
						"imei": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Imei",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"contains": dsschema.StringAttribute{
									MarkdownDescription: "Contains",
									Computed:            true,
								},
								"is": dsschema.StringAttribute{
									MarkdownDescription: "Is",
									Computed:            true,
								},
								"is_not": dsschema.StringAttribute{
									MarkdownDescription: "Is not",
									Computed:            true,
								},
							},
						},
						"jailbroken": dsschema.BoolAttribute{
							MarkdownDescription: "If device is by rooted/jailbroken",
							Computed:            true,
						},
						"last_checkin_time": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Last checkin time",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"not_within": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Not within",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"days": dsschema.Int64Attribute{
											MarkdownDescription: "specify time in days",
											Computed:            true,
										},
									},
								},
								"within": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Within",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"days": dsschema.Int64Attribute{
											MarkdownDescription: "specify time in days",
											Computed:            true,
										},
									},
								},
							},
						},
						"model": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Model",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"contains": dsschema.StringAttribute{
									MarkdownDescription: "Contains",
									Computed:            true,
								},
								"is": dsschema.StringAttribute{
									MarkdownDescription: "Is",
									Computed:            true,
								},
								"is_not": dsschema.StringAttribute{
									MarkdownDescription: "Is not",
									Computed:            true,
								},
							},
						},
						"passcode_set": dsschema.BoolAttribute{
							MarkdownDescription: "If device's passcode is present",
							Computed:            true,
						},
						"phone_number": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Phone number",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"contains": dsschema.StringAttribute{
									MarkdownDescription: "Contains",
									Computed:            true,
								},
								"is": dsschema.StringAttribute{
									MarkdownDescription: "Is",
									Computed:            true,
								},
								"is_not": dsschema.StringAttribute{
									MarkdownDescription: "Is not",
									Computed:            true,
								},
							},
						},
						"tag": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Tag",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"contains": dsschema.StringAttribute{
									MarkdownDescription: "Contains",
									Computed:            true,
								},
								"is": dsschema.StringAttribute{
									MarkdownDescription: "Is",
									Computed:            true,
								},
								"is_not": dsschema.StringAttribute{
									MarkdownDescription: "Is not",
									Computed:            true,
								},
							},
						},
					},
				},
			},
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "The name of the HIP object",
			Optional:            true,
			Computed:            true,
		},
		"network_info": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Network info",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"criteria": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Criteria",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"network": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Network",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"is": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Is",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"mobile": dsschema.SingleNestedAttribute{
											MarkdownDescription: "Mobile",
											Computed:            true,
											Attributes: map[string]dsschema.Attribute{
												"carrier": dsschema.StringAttribute{
													MarkdownDescription: "Carrier",
													Computed:            true,
												},
											},
										},
										"unknown": dsschema.SingleNestedAttribute{
											MarkdownDescription: "Unknown",
											Computed:            true,
											Attributes:          map[string]dsschema.Attribute{},
										},
										"wifi": dsschema.SingleNestedAttribute{
											MarkdownDescription: "Wifi",
											Computed:            true,
											Attributes: map[string]dsschema.Attribute{
												"ssid": dsschema.StringAttribute{
													MarkdownDescription: "SSID",
													Computed:            true,
												},
											},
										},
									},
								},
								"is_not": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Is not",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"ethernet": dsschema.SingleNestedAttribute{
											MarkdownDescription: "Ethernet",
											Computed:            true,
											Attributes:          map[string]dsschema.Attribute{},
										},
										"mobile": dsschema.SingleNestedAttribute{
											MarkdownDescription: "Mobile",
											Computed:            true,
											Attributes: map[string]dsschema.Attribute{
												"carrier": dsschema.StringAttribute{
													MarkdownDescription: "Carrier",
													Computed:            true,
												},
											},
										},
										"unknown": dsschema.SingleNestedAttribute{
											MarkdownDescription: "Unknown",
											Computed:            true,
											Attributes:          map[string]dsschema.Attribute{},
										},
										"wifi": dsschema.SingleNestedAttribute{
											MarkdownDescription: "Wifi",
											Computed:            true,
											Attributes: map[string]dsschema.Attribute{
												"ssid": dsschema.StringAttribute{
													MarkdownDescription: "SSID",
													Computed:            true,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"patch_management": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Patch management",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"criteria": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Criteria",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"is_enabled": dsschema.StringAttribute{
							MarkdownDescription: "is enabled",
							Computed:            true,
						},
						"is_installed": dsschema.BoolAttribute{
							MarkdownDescription: "Is Installed",
							Computed:            true,
						},
						"missing_patches": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Missing patches",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"check": dsschema.StringAttribute{
									MarkdownDescription: "Check",
									Computed:            true,
								},
								"patches": dsschema.ListAttribute{
									ElementType:         types.StringType,
									MarkdownDescription: "Patches",
									Computed:            true,
								},
								"severity": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Severity",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"greater_equal": dsschema.Int64Attribute{
											MarkdownDescription: "Greater equal",
											Computed:            true,
										},
										"greater_than": dsschema.Int64Attribute{
											MarkdownDescription: "Greater than",
											Computed:            true,
										},
										"is": dsschema.Int64Attribute{
											MarkdownDescription: "Is",
											Computed:            true,
										},
										"is_not": dsschema.Int64Attribute{
											MarkdownDescription: "Is not",
											Computed:            true,
										},
										"less_equal": dsschema.Int64Attribute{
											MarkdownDescription: "Less equal",
											Computed:            true,
										},
										"less_than": dsschema.Int64Attribute{
											MarkdownDescription: "Less than",
											Computed:            true,
										},
									},
								},
							},
						},
					},
				},
				"exclude_vendor": dsschema.BoolAttribute{
					MarkdownDescription: "Exclude vendor",
					Computed:            true,
				},
				"vendor": dsschema.ListNestedAttribute{
					MarkdownDescription: "Vendor name",
					Computed:            true,
					NestedObject: dsschema.NestedAttributeObject{
						Attributes: map[string]dsschema.Attribute{
							"name": dsschema.StringAttribute{
								MarkdownDescription: "Name",
								Computed:            true,
							},
							"product": dsschema.ListAttribute{
								ElementType:         types.StringType,
								MarkdownDescription: "Product name",
								Computed:            true,
							},
						},
					},
				},
			},
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

// HipObjectsListModel represents the data model for a list data source.
type HipObjectsListModel struct {
	Tfid    types.String `tfsdk:"tfid"`
	Data    []HipObjects `tfsdk:"data"`
	Limit   types.Int64  `tfsdk:"limit"`
	Offset  types.Int64  `tfsdk:"offset"`
	Name    types.String `tfsdk:"name"`
	Total   types.Int64  `tfsdk:"total"`
	Folder  types.String `tfsdk:"folder"`
	Snippet types.String `tfsdk:"snippet"`
	Device  types.String `tfsdk:"device"`
}

// HipObjectsListDataSourceSchema defines the schema for a list data source.
var HipObjectsListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: HipObjectsDataSourceSchema.Attributes,
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
