package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// GetResources returns all resources for the network_services package
func GetResources() []func() resource.Resource {
	return []func() resource.Resource{
		// 		NewSslDecryptionSettingResource,
		NewSdwanSaasQualityProfileResource,
		NewQosPolicyRuleResource,
		// 		NewAutoVpnClusterResource,
		NewSdwanTrafficDistributionProfileResource,
		NewBgpFilteringProfileResource,
		NewRouteAccessListResource,
		NewIkeGatewayResource,
		NewPbfRuleResource,
		NewRoutePathAccessListResource,
		NewVlanInterfaceResource,
		NewRoutePrefixListResource,
		NewQosProfileResource,
		NewEthernetInterfaceResource,
		NewBgpRouteMapResource,
		NewIpsecTunnelResource,
		NewBgpRouteMapRedistributionResource,
		NewOspfAuthProfileResource,
		NewBgpAuthProfileResource,
		NewNatRuleResource,
		NewSdwanPathQualityProfileResource,
		NewBgpAddressFamilyProfileResource,
		NewLogicalRouterResource,
		NewLinkTagResource,
		NewIpsecCryptoProfileResource,
		NewInterfaceManagementProfileResource,
		NewAggregateEthernetInterfaceResource,
		NewLoopbackInterfaceResource,
		NewSdwanErrorCorrectionProfileResource,
		NewSdwanRuleResource,
		NewRouteCommunityListResource,
		NewTunnelInterfaceResource,
		NewLayer3SubinterfaceResource,
		NewZoneResource,
		NewBgpRedistributionProfileResource,
		NewDhcpInterfaceResource,
		NewIkeCryptoProfileResource,
		NewZoneProtectionProfileResource,
		// 		NewLicenseResultResource,
		NewLayer2SubinterfaceResource,
		// 		NewAutoVpnPushResponseResource,
		NewDnsProxyResource,
	}
}
