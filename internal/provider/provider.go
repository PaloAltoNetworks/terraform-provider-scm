package provider

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	setup "github.com/paloaltonetworks/scm-go"

	tfProviderConfigSetup "github.com/paloaltonetworks/terraform-provider-scm/internal/provider/config_setup"
	tfProviderDeploymentServices "github.com/paloaltonetworks/terraform-provider-scm/internal/provider/deployment_services"
	tfProviderDeviceSettings "github.com/paloaltonetworks/terraform-provider-scm/internal/provider/device_settings"
	tfProviderIdentityServices "github.com/paloaltonetworks/terraform-provider-scm/internal/provider/identity_services"
	tfProviderNetworkServices "github.com/paloaltonetworks/terraform-provider-scm/internal/provider/network_services"
	tfProviderObjects "github.com/paloaltonetworks/terraform-provider-scm/internal/provider/objects"
	tfProviderSecurityServices "github.com/paloaltonetworks/terraform-provider-scm/internal/provider/security_services"
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

	// Only wrap the HTTP transport if we're using an auth file
	// When using an auth file, we need to re-read it periodically to pick up
	// tokens refreshed by external processes (e.g., cron jobs)
	// Otherwise, leave the SDK's default transport alone (it has built-in auto-refresh)
	if setupClient.AuthFile != "" {
		// Initialize HTTP client if needed
		if setupClient.HttpClient == nil {
			setupClient.HttpClient = &http.Client{}
		}
		if setupClient.HttpClient.Transport == nil {
			setupClient.HttpClient.Transport = http.DefaultTransport
		}

		// Add token refresh transport to re-read shared auth file when token expires
		// This allows long-running operations to use tokens refreshed by external processes (e.g., cron jobs)
		setupClient.HttpClient.Transport = &utils.TokenRefreshTransport{
			Wrapped:  setupClient.HttpClient.Transport,
			Client:   setupClient,
			AuthFile: setupClient.AuthFile,
			Ctx:      ctx,
		}

	}

	// Add logging transport if debug mode is enabled (works in both auth file and non-auth-file modes)
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
	// Only refresh if no JWT exists or if it's expired/near expiry
	if setupClient.Jwt == "" || time.Now().After(setupClient.JwtExpiresAt) {
		if err := setupClient.RefreshJwt(ctx); err != nil {
			resp.Diagnostics.AddError("Refresh JWT Authentication error", err.Error())
			return
		}
	}

	// Create a client container with all service-specific clients
	// Create clients map
	clients := map[string]interface{}{
		"config_setup":        setup.GetConfigSetupAPIClient(setupClient),
		"deployment_services": setup.GetDeploymentServicesAPIClient(setupClient),
		"device_settings":     setup.GetDeviceSettingsAPIClient(setupClient),
		"identity_services":   setup.GetIdentityServicesAPIClient(setupClient),
		"network_services":    setup.GetNetworkServicesAPIClient(setupClient),
		"objects":             setup.GetObjectsAPIClient(setupClient),
		"security_services":   setup.GetSecurityServicesAPIClient(setupClient),
	}

	resp.DataSourceData = clients
	resp.ResourceData = clients
	tflog.Info(ctx, "Configured client", map[string]any{"success": true})
}

// DataSources defines the data sources for this provider.
func (p *ScmProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	var dataSources []func() datasource.DataSource
	// Add config_setup package data sources
	dataSources = append(dataSources, tfProviderConfigSetup.GetDataSources()...)
	// Add deployment_services package data sources
	dataSources = append(dataSources, tfProviderDeploymentServices.GetDataSources()...)
	// Add device_settings package data sources
	dataSources = append(dataSources, tfProviderDeviceSettings.GetDataSources()...)
	// Add identity_services package data sources
	dataSources = append(dataSources, tfProviderIdentityServices.GetDataSources()...)
	// Add network_services package data sources
	dataSources = append(dataSources, tfProviderNetworkServices.GetDataSources()...)
	// Add objects package data sources
	dataSources = append(dataSources, tfProviderObjects.GetDataSources()...)
	// Add security_services package data sources
	dataSources = append(dataSources, tfProviderSecurityServices.GetDataSources()...)
	return dataSources
}

// Resources defines the resources for this provider.
func (p *ScmProvider) Resources(ctx context.Context) []func() resource.Resource {
	var resources []func() resource.Resource
	// Add config_setup package resources
	resources = append(resources, tfProviderConfigSetup.GetResources()...)
	// Add deployment_services package resources
	resources = append(resources, tfProviderDeploymentServices.GetResources()...)
	// Add device_settings package resources
	resources = append(resources, tfProviderDeviceSettings.GetResources()...)
	// Add identity_services package resources
	resources = append(resources, tfProviderIdentityServices.GetResources()...)
	// Add network_services package resources
	resources = append(resources, tfProviderNetworkServices.GetResources()...)
	// Add objects package resources
	resources = append(resources, tfProviderObjects.GetResources()...)
	// Add security_services package resources
	resources = append(resources, tfProviderSecurityServices.GetResources()...)

	return resources
}
