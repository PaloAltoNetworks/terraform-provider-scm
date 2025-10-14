package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// GetResources returns all resources for the network_services package
func GetResources() []func() resource.Resource {
	return []func() resource.Resource{
		NewAggregateEthernetInterfaceResource,
		// 		NewAutoVpnClusterResource,
		// 		NewAutoVpnPushResponseResource,
		NewBgpAddressFamilyProfileResource,
		NewBgpAuthProfileResource,
		NewBgpFilteringProfileResource,
		NewBgpRedistributionProfileResource,
		NewBgpRouteMapResource,
		NewBgpRouteMapRedistributionResource,
		NewDhcpInterfaceResource,
		NewDnsProxyResource,
		NewEthernetInterfaceResource,
		NewIkeCryptoProfileResource,
		NewIkeGatewayResource,
		NewInterfaceManagementProfileResource,
		NewIpsecCryptoProfileResource,
		NewIpsecTunnelResource,
		NewLayer2SubinterfaceResource,
		NewLayer3SubinterfaceResource,
		// 		NewLicenseResultResource,
		NewLinkTagResource,
		NewLogicalRouterResource,
		NewLoopbackInterfaceResource,
		NewNatRuleResource,
		NewOspfAuthProfileResource,
		NewPbfRuleResource,
		NewQosPolicyRuleResource,
		NewQosProfileResource,
		NewRouteAccessListResource,
		NewRouteCommunityListResource,
		NewRoutePathAccessListResource,
		NewRoutePrefixListResource,
		NewSdwanErrorCorrectionProfileResource,
		NewSdwanPathQualityProfileResource,
		NewSdwanRuleResource,
		NewSdwanSaasQualityProfileResource,
		NewSdwanTrafficDistributionProfileResource,
		// 		NewSslDecryptionSettingResource,
		NewTunnelInterfaceResource,
		NewVlanInterfaceResource,
		NewZoneResource,
		NewZoneProtectionProfileResource,
	}
}
