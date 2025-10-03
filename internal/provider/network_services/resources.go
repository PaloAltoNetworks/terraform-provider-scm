package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// GetResources returns all resources for the network_services package
func GetResources() []func() resource.Resource {
	return []func() resource.Resource{
		NewDhcpInterfaceResource,
		NewOspfAuthProfileResource,
		NewQosProfileResource,
		NewLayer3SubinterfaceResource,
		NewRoutePrefixListResource,
		NewRouteCommunityListResource,
		NewIpsecCryptoProfileResource,
		NewSdwanErrorCorrectionProfileResource,
		// 		NewSslDecryptionSettingResource,
		NewLogicalRouterResource,
		NewBgpAuthProfileResource,
		NewLinkTagResource,
		NewSdwanTrafficDistributionProfileResource,
		NewRoutePathAccessListResource,
		NewNatRuleResource,
		// 		NewAutoVpnPushResponseResource,
		// 		NewAutoVpnClusterResource,
		NewBgpRouteMapRedistributionResource,
		NewEthernetInterfaceResource,
		NewAggregateEthernetInterfaceResource,
		NewBgpFilteringProfileResource,
		NewVlanInterfaceResource,
		NewZoneProtectionProfileResource,
		NewBgpAddressFamilyProfileResource,
		NewBgpRedistributionProfileResource,
		NewTunnelInterfaceResource,
		NewQosPolicyRuleResource,
		NewZoneResource,
		NewBgpRouteMapResource,
		NewLoopbackInterfaceResource,
		NewSdwanSaasQualityProfileResource,
		NewDnsProxyResource,
		// 		NewLicenseResultResource,
		NewInterfaceManagementProfileResource,
		NewSdwanPathQualityProfileResource,
		NewRouteAccessListResource,
		NewSdwanRuleResource,
		NewIkeGatewayResource,
		NewPbfRuleResource,
		NewIkeCryptoProfileResource,
		NewLayer2SubinterfaceResource,
		NewIpsecTunnelResource,
	}
}
