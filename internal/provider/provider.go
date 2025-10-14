package provider

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	setup "github.com/paloaltonetworks/scm-go"

	"github.com/paloaltonetworks/scm-go/generated/config_setup"
	"github.com/paloaltonetworks/scm-go/generated/deployment_services"
	"github.com/paloaltonetworks/scm-go/generated/identity_services"
	"github.com/paloaltonetworks/scm-go/generated/network_services"
	"github.com/paloaltonetworks/scm-go/generated/objects"
	"github.com/paloaltonetworks/scm-go/generated/security_services"

	tfProviderConfig_setup "github.com/paloaltonetworks/terraform-provider-scm/internal/provider/config_setup"
	tfProviderDeployment_services "github.com/paloaltonetworks/terraform-provider-scm/internal/provider/deployment_services"
	tfProviderIdentity_services "github.com/paloaltonetworks/terraform-provider-scm/internal/provider/identity_services"
	tfProviderNetwork_services "github.com/paloaltonetworks/terraform-provider-scm/internal/provider/network_services"
	tfProviderObjects "github.com/paloaltonetworks/terraform-provider-scm/internal/provider/objects"
	tfProviderSecurity_services "github.com/paloaltonetworks/terraform-provider-scm/internal/provider/security_services"
	"github.com/paloaltonetworks/terraform-provider-scm/internal/utils"
)

// Ensure the provider implementation interface is sound.
var (
	_ provider.Provider = &ScmProvider{}
)

func New(version string) provider.Provider {
	return &ScmProvider{
		version: version,
	}
}

// NewFactory returns a function that creates a new provider instance with the given version
func NewFactory(version string) func() provider.Provider {
	return func() provider.Provider {
		return New(version)
	}
}

// ScmProvider is the provider implementation.
type ScmProvider struct {
	version string
}

// ScmProviderModel maps provider schema data to a Go type.
type ScmProviderModel struct {
	AuthUrl      types.String `tfsdk:"auth_url"`
	Protocol     types.String `tfsdk:"protocol"`
	Host         types.String `tfsdk:"host"`
	Port         types.Int64  `tfsdk:"port"`
	Headers      types.Map    `tfsdk:"headers"`
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
			"auth_url": schema.StringAttribute{
				Description: "The URL to send auth credentials to which will return a JWT. Default: `https://auth.apps.paloaltonetworks.com/auth/v1/oauth2/access_token`. Environment variable: `SCM_AUTH_URL`. JSON config file variable: `auth_url`.",
				Optional:    true,
			},
			"protocol": schema.StringAttribute{
				Description: "The protocol to use for SCM. This should be 'http' or 'https'. Default: `https`. Environment variable: `SCM_PROTOCOL`. JSON config file variable: `protocol`.",
				Optional:    true,
			},
			"host": schema.StringAttribute{
				Description: "The hostname of Strata Cloud Manager API. Default: `api.sase.paloaltonetworks.com`. Environment variable: `SCM_HOST`. JSON config file variable: `host`.",
				Optional:    true,
			},
			"port": schema.Int64Attribute{
				Description: "The port number to use for API commands, if non-standard for the given protocol. Environment variable: `SCM_PORT`. JSON config file variable: `port`.",
				Optional:    true,
			},
			"headers": schema.MapAttribute{
				Description: "Custom HTTP headers to be sent with all API commands. Environment variable: `SCM_HEADERS`. JSON config file variable: `headers`.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"client_id": schema.StringAttribute{
				Description: "The client ID for the connection. Environment variable: `SCM_CLIENT_ID`. JSON config file variable: `client_id`.",
				Optional:    true,
			},
			"client_secret": schema.StringAttribute{
				Description: "The client secret for the connection. Environment variable: `SCM_CLIENT_SECRET`. JSON config file variable: `client_secret`.",
				Optional:    true,
				Sensitive:   true,
			},
			"scope": schema.StringAttribute{
				Description: "The client scope. Environment variable: `SCM_SCOPE`. JSON config file variable: `scope`.",
				Optional:    true,
			},
			"logging": schema.StringAttribute{
				Description: "The logging level of the provider and the underlying communication. Default: `quiet`. Environment variable: `SCM_LOGGING`. JSON config file variable: `logging`.",
				Optional:    true,
			},
			"auth_file": schema.StringAttribute{
				Description: "The file path to the JSON file with auth creds for SCM.",
				Optional:    true,
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
	ht := make(map[string]types.String, len(config.Headers.Elements()))
	resp.Diagnostics.Append(config.Headers.ElementsAs(ctx, &ht, false).Errors()...)
	if resp.Diagnostics.HasError() {
		return
	}

	headers := make(map[string]string, len(ht))
	for hkey, hval := range ht {
		headers[hkey] = hval.ValueString()
	}

	setupClient := &setup.Client{
		AuthUrl:          config.AuthUrl.ValueString(),
		Protocol:         config.Protocol.ValueString(),
		Host:             config.Host.ValueString(),
		Port:             int(config.Port.ValueInt64()),
		Headers:          headers,
		ClientId:         config.ClientId.ValueString(),
		ClientSecret:     config.ClientSecret.ValueString(),
		Scope:            config.Scope.ValueString(),
		AuthFile:         config.AuthFile.ValueString(),
		Logging:          config.Logging.ValueString(),
		Logger:           tflog.Debug,
		CheckEnvironment: true,
		Agent:            fmt.Sprintf("Terraform/%s Provider/scm Version/%s", req.TerraformVersion, p.version),
	}

	if err := setupClient.Setup(); err != nil {
		resp.Diagnostics.AddError("Provider parameter value error", err.Error())
		return
	}

	if config.Logging.ValueString() == "debug" {
		if setupClient.HttpClient == nil {
			setupClient.HttpClient = &http.Client{}
		}
		if setupClient.HttpClient.Transport == nil {
			setupClient.HttpClient.Transport = http.DefaultTransport
		}
		// FIX: Pass the context (ctx) to the LoggingRoundTripper.
		setupClient.HttpClient.Transport = &utils.LoggingRoundTripper{
			Wrapped: setupClient.HttpClient.Transport,
			Ctx:     ctx,
		}
	}

	// Refresh JWT token
	if setupClient.Jwt == "" {
		// Only refresh if no JWT exists
		if err := setupClient.RefreshJwt(ctx); err != nil {
			resp.Diagnostics.AddError("Refresh JWT Authentication error", err.Error())
			return
		}
	}

	// Create a client container with all service-specific clients
	// Create clients map
	clients := map[string]interface{}{
		"config_setup":        p.getConfig_setupAPIClient(setupClient),
		"deployment_services": p.getDeployment_servicesAPIClient(setupClient),
		"identity_services":   p.getIdentity_servicesAPIClient(setupClient),
		"network_services":    p.getNetwork_servicesAPIClient(setupClient),
		"objects":             p.getObjectsAPIClient(setupClient),
		"security_services":   p.getSecurity_servicesAPIClient(setupClient),
	}

	resp.DataSourceData = clients
	resp.ResourceData = clients
	tflog.Info(ctx, "Configured client", map[string]any{"success": true})
}

func (p *ScmProvider) getConfig_setupAPIClient(setupClient *setup.Client) *config_setup.APIClient {
	// Create the config_setup API client
	config_setupConfig := config_setup.NewConfiguration()
	config_setupConfig.Host = setupClient.GetHost()
	config_setupConfig.Scheme = "https"
	config_setupConfig.HTTPClient = setupClient.HttpClient
	// Set up the default header with JWT
	config_setupConfig.DefaultHeader = make(map[string]string)
	config_setupConfig.DefaultHeader["Authorization"] = "Bearer " + setupClient.Jwt
	config_setupConfig.DefaultHeader["x-auth-jwt"] = setupClient.Jwt
	config_setupApiClient := config_setup.NewAPIClient(config_setupConfig)
	return config_setupApiClient
}

func (p *ScmProvider) getDeployment_servicesAPIClient(setupClient *setup.Client) *deployment_services.APIClient {
	// Create the deployment_services API client
	deployment_servicesConfig := deployment_services.NewConfiguration()
	deployment_servicesConfig.Host = setupClient.GetHost()
	deployment_servicesConfig.Scheme = "https"
	deployment_servicesConfig.HTTPClient = setupClient.HttpClient
	// Set up the default header with JWT
	deployment_servicesConfig.DefaultHeader = make(map[string]string)
	deployment_servicesConfig.DefaultHeader["Authorization"] = "Bearer " + setupClient.Jwt
	deployment_servicesConfig.DefaultHeader["x-auth-jwt"] = setupClient.Jwt
	deployment_servicesApiClient := deployment_services.NewAPIClient(deployment_servicesConfig)
	return deployment_servicesApiClient
}

func (p *ScmProvider) getIdentity_servicesAPIClient(setupClient *setup.Client) *identity_services.APIClient {
	// Create the identity_services API client
	identity_servicesConfig := identity_services.NewConfiguration()
	identity_servicesConfig.Host = setupClient.GetHost()
	identity_servicesConfig.Scheme = "https"
	identity_servicesConfig.HTTPClient = setupClient.HttpClient
	// Set up the default header with JWT
	identity_servicesConfig.DefaultHeader = make(map[string]string)
	identity_servicesConfig.DefaultHeader["Authorization"] = "Bearer " + setupClient.Jwt
	identity_servicesConfig.DefaultHeader["x-auth-jwt"] = setupClient.Jwt
	identity_servicesApiClient := identity_services.NewAPIClient(identity_servicesConfig)
	return identity_servicesApiClient
}

func (p *ScmProvider) getNetwork_servicesAPIClient(setupClient *setup.Client) *network_services.APIClient {
	// Create the network_services API client
	network_servicesConfig := network_services.NewConfiguration()
	network_servicesConfig.Host = setupClient.GetHost()
	network_servicesConfig.Scheme = "https"
	network_servicesConfig.HTTPClient = setupClient.HttpClient
	// Set up the default header with JWT
	network_servicesConfig.DefaultHeader = make(map[string]string)
	network_servicesConfig.DefaultHeader["Authorization"] = "Bearer " + setupClient.Jwt
	network_servicesConfig.DefaultHeader["x-auth-jwt"] = setupClient.Jwt
	network_servicesApiClient := network_services.NewAPIClient(network_servicesConfig)
	return network_servicesApiClient
}

func (p *ScmProvider) getObjectsAPIClient(setupClient *setup.Client) *objects.APIClient {
	// Create the objects API client
	objectsConfig := objects.NewConfiguration()
	objectsConfig.Host = setupClient.GetHost()
	objectsConfig.Scheme = "https"
	objectsConfig.HTTPClient = setupClient.HttpClient
	// Set up the default header with JWT
	objectsConfig.DefaultHeader = make(map[string]string)
	objectsConfig.DefaultHeader["Authorization"] = "Bearer " + setupClient.Jwt
	objectsConfig.DefaultHeader["x-auth-jwt"] = setupClient.Jwt
	objectsApiClient := objects.NewAPIClient(objectsConfig)
	return objectsApiClient
}

func (p *ScmProvider) getSecurity_servicesAPIClient(setupClient *setup.Client) *security_services.APIClient {
	// Create the security_services API client
	security_servicesConfig := security_services.NewConfiguration()
	security_servicesConfig.Host = setupClient.GetHost()
	security_servicesConfig.Scheme = "https"
	security_servicesConfig.HTTPClient = setupClient.HttpClient
	// Set up the default header with JWT
	security_servicesConfig.DefaultHeader = make(map[string]string)
	security_servicesConfig.DefaultHeader["Authorization"] = "Bearer " + setupClient.Jwt
	security_servicesConfig.DefaultHeader["x-auth-jwt"] = setupClient.Jwt
	security_servicesApiClient := security_services.NewAPIClient(security_servicesConfig)
	return security_servicesApiClient
}

// DataSources defines the data sources for this provider.
func (p *ScmProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	var dataSources []func() datasource.DataSource
	// Add config_setup package data sources
	dataSources = append(dataSources, tfProviderConfig_setup.GetDataSources()...)
	// Add deployment_services package data sources
	dataSources = append(dataSources, tfProviderDeployment_services.GetDataSources()...)
	// Add identity_services package data sources
	dataSources = append(dataSources, tfProviderIdentity_services.GetDataSources()...)
	// Add network_services package data sources
	dataSources = append(dataSources, tfProviderNetwork_services.GetDataSources()...)
	// Add objects package data sources
	dataSources = append(dataSources, tfProviderObjects.GetDataSources()...)
	// Add security_services package data sources
	dataSources = append(dataSources, tfProviderSecurity_services.GetDataSources()...)
	return dataSources
}

// Resources defines the resources for this provider.
func (p *ScmProvider) Resources(ctx context.Context) []func() resource.Resource {
	var resources []func() resource.Resource
	// Add config_setup package resources
	resources = append(resources, tfProviderConfig_setup.GetResources()...)
	// Add deployment_services package resources
	resources = append(resources, tfProviderDeployment_services.GetResources()...)
	// Add identity_services package resources
	resources = append(resources, tfProviderIdentity_services.GetResources()...)
	// Add network_services package resources
	resources = append(resources, tfProviderNetwork_services.GetResources()...)
	// Add objects package resources
	resources = append(resources, tfProviderObjects.GetResources()...)
	// Add security_services package resources
	resources = append(resources, tfProviderSecurity_services.GetResources()...)

	return resources
}
