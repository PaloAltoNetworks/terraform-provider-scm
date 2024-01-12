package provider

import (
	"context"
	"fmt"

	sdk "github.com/paloaltonetworks/scm-go"
	sdkapi "github.com/paloaltonetworks/scm-go/api"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the provider implementation interface is sound.
var (
	_ provider.Provider = &ScmProvider{}
)

// ScmProvider is the provider implementation.
type ScmProvider struct {
	version string
}

// ScmProviderModel maps provider schema data to a Go type.
type ScmProviderModel struct {
	Host         types.String `tfsdk:"host"`
	ClientId     types.String `tfsdk:"client_id"`
	ClientSecret types.String `tfsdk:"client_secret"`
	Scope        types.String `tfsdk:"scope"`
	Logging      types.String `tfsdk:"logging"`
	AuthFile     types.String `tfsdk:"auth_file"`
}

// Metadata returns the provider type name.
func (p *ScmProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "scm"
	resp.Version = p.version
}

// Schema defines the provider-level schema for configuration data.
func (p *ScmProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Terraform provider to interact with Palo Alto Networks Strata Cloud Manager API.",
		Attributes: map[string]schema.Attribute{
			"host": schema.StringAttribute{
				Description: ProviderParamDescription(
					"The hostname of Strata Cloud Manager API.",
					"api.sase.paloaltonetworks.com",
					"SCM_HOST",
					"host",
				),
				Optional: true,
			},
			"client_id": schema.StringAttribute{
				Description: ProviderParamDescription(
					"The client ID for the connection.",
					"",
					"SCM_CLIENT_ID",
					"client_id",
				),
				Optional: true,
			},
			"client_secret": schema.StringAttribute{
				Description: ProviderParamDescription(
					"The client secret for the connection.",
					"",
					"SCM_CLIENT_SECRET",
					"client_secret",
				),
				Optional:  true,
				Sensitive: true,
			},
			"scope": schema.StringAttribute{
				Description: ProviderParamDescription(
					"The client scope.",
					"",
					"SCM_SCOPE",
					"scope",
				),
				Optional: true,
			},
			"logging": schema.StringAttribute{
				Description: ProviderParamDescription(
					"The logging level of the provider and the underlying communication.",
					sdkapi.LogQuiet,
					"SCM_LOGGING",
					"logging",
				),
				Optional: true,
			},
			"auth_file": schema.StringAttribute{
				Description: ProviderParamDescription(
					"The file path to the JSON file with auth creds for SCM.",
					"",
					"",
					"",
				),
				Optional: true,
			},
		},
	}
}

// Configure prepares the provider.
func (p *ScmProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	tflog.Info(ctx, "Configuring the provider client")

	var config ScmProviderModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Configure the client.
	con := &sdk.Client{
		Host:             config.Host.ValueString(),
		ClientId:         config.ClientId.ValueString(),
		ClientSecret:     config.ClientSecret.ValueString(),
		Scope:            config.Scope.ValueString(),
		AuthFile:         config.AuthFile.ValueString(),
		Logging:          config.Logging.ValueString(),
		Logger:           tflog.Debug,
		CheckEnvironment: true,
		Agent:            fmt.Sprintf("Terraform/%s Provider/%s", req.TerraformVersion, p.version),
	}

	if err := con.Setup(); err != nil {
		resp.Diagnostics.AddError("Provider parameter value error", err.Error())
		return
	}

	con.HttpClient.Transport = sdkapi.NewTransport(con.HttpClient.Transport, con)

	if err := con.RefreshJwt(ctx); err != nil {
		resp.Diagnostics.AddError("Authentication error", err.Error())
		return
	}

	resp.DataSourceData = con
	resp.ResourceData = con

	// Done.
	tflog.Info(ctx, "Configured client", map[string]any{"success": true})
}

// DataSources defines the data sources for this provider.
func (p *ScmProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewAddressGroupDataSource,
		NewAddressGroupListDataSource,
		NewAddressObjectDataSource,
		NewAddressObjectListDataSource,
		NewAntiSpywareProfileDataSource,
		NewAntiSpywareProfileListDataSource,
		NewAntiSpywareSignatureDataSource,
		NewAntiSpywareSignatureListDataSource,
		NewAppOverrideRuleDataSource,
		NewAppOverrideRuleListDataSource,
		NewApplicationDataSource,
		NewApplicationFilterDataSource,
		NewApplicationFilterListDataSource,
		NewApplicationGroupDataSource,
		NewApplicationGroupListDataSource,
		NewApplicationListDataSource,
		NewAuthenticationPortalDataSource,
		NewAuthenticationPortalListDataSource,
		NewAuthenticationProfileDataSource,
		NewAuthenticationProfileListDataSource,
		NewAuthenticationRuleDataSource,
		NewAuthenticationRuleListDataSource,
		NewAuthenticationSequenceDataSource,
		NewAuthenticationSequenceListDataSource,
		NewAutoTagActionsListDataSource,
		NewCertificateProfileDataSource,
		NewCertificateProfileListDataSource,
		NewDecryptionExclusionDataSource,
		NewDecryptionProfileDataSource,
		NewDecryptionProfileListDataSource,
		NewDecryptionRuleDataSource,
		NewDecryptionRuleListDataSource,
		NewDeviceDataSource,
		NewDeviceListDataSource,
		NewDnsSecurityProfileDataSource,
		NewDnsSecurityProfileListDataSource,
		NewDynamicUserGroupDataSource,
		NewDynamicUserGroupListDataSource,
		NewExternalDynamicListDataSource,
		NewExternalDynamicListListDataSource,
		NewFileBlockingProfileDataSource,
		NewFileBlockingProfileListDataSource,
		NewHipObjectDataSource,
		NewHipObjectListDataSource,
		NewHipProfileDataSource,
		NewHipProfileListDataSource,
		NewHttpHeaderProfileDataSource,
		NewHttpHeaderProfileListDataSource,
		NewIkeCryptoProfileDataSource,
		NewIkeCryptoProfileListDataSource,
		NewIkeGatewayDataSource,
		NewIkeGatewayListDataSource,
		NewInternalDnsServerDataSource,
		NewInternalDnsServerListDataSource,
		NewIpsecCryptoProfileDataSource,
		NewIpsecCryptoProfileListDataSource,
		NewIpsecTunnelDataSource,
		NewIpsecTunnelListDataSource,
		NewJobsDataSource,
		NewJobsListDataSource,
		NewKerberosServerProfileDataSource,
		NewKerberosServerProfileListDataSource,
		NewLabelListDataSource,
		NewLabelsGetbyidResponseDataSource,
		NewLdapServerProfileDataSource,
		NewLdapServerProfileListDataSource,
		NewLocalUserDataSource,
		NewLocalUserGroupListDataSource,
		NewLocalUserListDataSource,
		NewMfaServerDataSource,
		NewOcspResponderDataSource,
		NewOcspResponderListDataSource,
		NewProfileGroupDataSource,
		NewProfileGroupListDataSource,
		NewQosPolicyRuleDataSource,
		NewQosPolicyRuleListDataSource,
		NewQosProfileDataSource,
		NewQosProfileListDataSource,
		NewRadiusServerProfileDataSource,
		NewRadiusServerProfileListDataSource,
		NewRegionDataSource,
		NewRegionListDataSource,
		NewRemoteNetworkDataSource,
		NewRemoteNetworkListDataSource,
		NewSamlServerProfileDataSource,
		NewSamlServerProfileListDataSource,
		NewScepProfileDataSource,
		NewScepProfileListDataSource,
		NewScheduleDataSource,
		NewScheduleListDataSource,
		NewSecurityRuleDataSource,
		NewSecurityRuleListDataSource,
		NewServiceDataSource,
		NewServiceGroupDataSource,
		NewServiceGroupListDataSource,
		NewServiceListDataSource,
		NewSharedInfrastructureSettingsListDataSource,
		NewSnippetDataSource,
		NewSnippetListDataSource,
		NewTacacsServerProfileDataSource,
		NewTacacsServerProfileListDataSource,
		NewTagDataSource,
		NewTagListDataSource,
		NewTlsServiceProfileDataSource,
		NewTlsServiceProfileListDataSource,
		NewTrustedCertificateAuthorityListDataSource,
		NewUrlAccessProfileDataSource,
		NewUrlAccessProfileListDataSource,
		NewUrlCategoryListDataSource,
		NewUrlFilteringCategoryListDataSource,
		NewVariableDataSource,
		NewVariableListDataSource,
		NewVulnerabilityProtectionProfileDataSource,
		NewVulnerabilityProtectionProfileListDataSource,
		NewVulnerabilityProtectionSignaturesDataSource,
		NewVulnerabilityProtectionSignaturesListDataSource,
		NewWildfireAntiVirusProfileDataSource,
		NewWildfireAntiVirusProfileListDataSource,
	}
}

// Resources defines the data sources for this provider.
func (p *ScmProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewAddressGroupResource,
		NewAddressObjectResource,
		NewAntiSpywareProfileResource,
		NewAntiSpywareSignatureResource,
		NewAppOverrideRuleResource,
		NewApplicationFilterResource,
		NewApplicationGroupResource,
		NewApplicationResource,
		NewAuthenticationPortalResource,
		NewAuthenticationProfileResource,
		NewAuthenticationRuleResource,
		NewAuthenticationSequenceResource,
		NewCertificateProfileResource,
		NewDecryptionExclusionResource,
		NewDecryptionProfileResource,
		NewDecryptionRuleResource,
		NewDnsSecurityProfileResource,
		NewDynamicUserGroupResource,
		NewExternalDynamicListResource,
		NewFileBlockingProfileResource,
		NewHipObjectResource,
		NewHipProfileResource,
		NewHttpHeaderProfileResource,
		NewIkeCryptoProfileResource,
		NewIkeGatewayResource,
		NewInternalDnsServerResource,
		NewIpsecCryptoProfileResource,
		NewIpsecTunnelResource,
		NewKerberosServerProfileResource,
		NewLdapServerProfileResource,
		NewLocalUserResource,
		NewMfaServerResource,
		NewOcspResponderResource,
		NewProfileGroupResource,
		NewQosPolicyRuleResource,
		NewQosProfileResource,
		NewRadiusServerProfileResource,
		NewRegionResource,
		NewRemoteNetworkResource,
		NewSamlServerProfileResource,
		NewScepProfileResource,
		NewScheduleResource,
		NewSecurityRuleResource,
		NewServiceGroupResource,
		NewServiceResource,
		NewSnippetResource,
		NewTacacsServerProfileResource,
		NewTagResource,
		NewTlsServiceProfileResource,
		NewUrlAccessProfileResource,
		NewVariableResource,
		NewVulnerabilityProtectionProfileResource,
		NewVulnerabilityProtectionSignaturesResource,
		NewWildfireAntiVirusProfileResource,
	}
}

// New is a helper function to get the provider implementation.
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &ScmProvider{
			version: version,
		}
	}
}
