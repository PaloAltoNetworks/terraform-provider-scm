package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// GetResources returns all resources for the network_services package
func GetResources() []func() resource.Resource {
	return []func() resource.Resource{
		NewIpsecTunnelResource,
		NewIkeCryptoProfileResource,
		NewBgpAddressFamilyProfileResource,
		NewLogicalRouterResource,
		NewBgpRouteMapResource,
		NewOspfAuthProfileResource,
		NewPbfRuleResource,
		NewZoneProtectionProfileResource,
		NewIpsecCryptoProfileResource,
		NewBgpFilteringProfileResource,
		NewSdwanPathQualityProfileResource,
		NewSdwanTrafficDistributionProfileResource,
		// 		NewAutoVpnClusterResource,
		NewLinkTagResource,
		NewAggregateEthernetInterfaceResource,
		NewRoutePathAccessListResource,
		NewQosProfileResource,
		NewBgpRedistributionProfileResource,
		// 		NewLicenseResultResource,
		NewSdwanErrorCorrectionProfileResource,
		NewRouteCommunityListResource,
		NewSdwanSaasQualityProfileResource,
		NewQosPolicyRuleResource,
		// 		NewSslDecryptionSettingResource,
		NewBgpRouteMapRedistributionResource,
		NewLayer2SubinterfaceResource,
		NewDnsProxyResource,
		NewIkeGatewayResource,
		NewRouteAccessListResource,
		NewVlanInterfaceResource,
		NewDhcpInterfaceResource,
		NewRoutePrefixListResource,
		NewInterfaceManagementProfileResource,
		NewBgpAuthProfileResource,
		NewLoopbackInterfaceResource,
		NewSdwanRuleResource,
		NewLayer3SubinterfaceResource,
		// 		NewAutoVpnPushResponseResource,
		NewZoneResource,
		NewNatRuleResource,
		NewTunnelInterfaceResource,
		NewEthernetInterfaceResource,
	}
}
