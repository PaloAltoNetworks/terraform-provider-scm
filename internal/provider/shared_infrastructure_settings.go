package provider

// Note:  This file is automatically generated.  Manually made changes
// will be overwritten when the provider is generated.

import (
	"context"
	"strconv"
	"strings"

	"github.com/paloaltonetworks/scm-go"
	xCzsxgj "github.com/paloaltonetworks/scm-go/netsec/services/sharedinfrastructuresettings"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	dsschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Data source (listing).
var (
	_ datasource.DataSource              = &sharedInfrastructureSettingsListDataSource{}
	_ datasource.DataSourceWithConfigure = &sharedInfrastructureSettingsListDataSource{}
)

func NewSharedInfrastructureSettingsListDataSource() datasource.DataSource {
	return &sharedInfrastructureSettingsListDataSource{}
}

type sharedInfrastructureSettingsListDataSource struct {
	client *scm.Client
}

// sharedInfrastructureSettingsListDsModel is the model.
type sharedInfrastructureSettingsListDsModel struct {
	Tfid types.String `tfsdk:"tfid"`

	// Input.
	Limit  types.Int64 `tfsdk:"limit"`
	Offset types.Int64 `tfsdk:"offset"`

	// Output.
	Data []sharedInfrastructureSettingsListDsModel_eItpily_Config `tfsdk:"data"`
	// omit input: limit
	// omit input: offset
	Total types.Int64 `tfsdk:"total"`
}

type sharedInfrastructureSettingsListDsModel_eItpily_Config struct {
	ApiKey                         types.String `tfsdk:"api_key"`
	CaptivePortalRedirectIpAddress types.String `tfsdk:"captive_portal_redirect_ip_address"`
	EgressIpNotificationUrl        types.String `tfsdk:"egress_ip_notification_url"`
	InfraBgpAs                     types.String `tfsdk:"infra_bgp_as"`
	InfrastructureSubnet           types.String `tfsdk:"infrastructure_subnet"`
	InfrastructureSubnetIpv6       types.String `tfsdk:"infrastructure_subnet_ipv6"`
	Ipv6                           types.Bool   `tfsdk:"ipv6"`
	LoopbackIps                    types.List   `tfsdk:"loopback_ips"`
	TunnelMonitorIpAddress         types.String `tfsdk:"tunnel_monitor_ip_address"`
}

// Metadata returns the data source type name.
func (d *sharedInfrastructureSettingsListDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_shared_infrastructure_settings_list"
}

// Schema defines the schema for this listing data source.
func (d *sharedInfrastructureSettingsListDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = dsschema.Schema{
		Description: "Retrieves a listing of config items.",

		Attributes: map[string]dsschema.Attribute{
			// inputs:map[string]bool{"limit":true, "offset":true} outputs:map[string]bool{"data":true, "limit":true, "offset":true, "tfid":true, "total":true} forceNew:map[string]bool{"limit":true, "offset":true}
			"data": dsschema.ListNestedAttribute{
				Description: "The Data param.",
				Computed:    true,
				NestedObject: dsschema.NestedAttributeObject{
					Attributes: map[string]dsschema.Attribute{
						// inputs:map[string]bool{} outputs:map[string]bool{"api_key":true, "captive_portal_redirect_ip_address":true, "egress_ip_notification_url":true, "infra_bgp_as":true, "infrastructure_subnet":true, "infrastructure_subnet_ipv6":true, "ipv6":true, "loopback_ips":true, "tunnel_monitor_ip_address":true} forceNew:map[string]bool(nil)
						"api_key": dsschema.StringAttribute{
							Description: "The ApiKey param.",
							Computed:    true,
						},
						"captive_portal_redirect_ip_address": dsschema.StringAttribute{
							Description: "The CaptivePortalRedirectIpAddress param.",
							Computed:    true,
						},
						"egress_ip_notification_url": dsschema.StringAttribute{
							Description: "The EgressIpNotificationUrl param.",
							Computed:    true,
						},
						"infra_bgp_as": dsschema.StringAttribute{
							Description: "The InfraBgpAs param.",
							Computed:    true,
						},
						"infrastructure_subnet": dsschema.StringAttribute{
							Description: "The InfrastructureSubnet param.",
							Computed:    true,
						},
						"infrastructure_subnet_ipv6": dsschema.StringAttribute{
							Description: "The InfrastructureSubnetIpv6 param.",
							Computed:    true,
						},
						"ipv6": dsschema.BoolAttribute{
							Description: "The Ipv6 param.",
							Computed:    true,
						},
						"loopback_ips": dsschema.ListAttribute{
							Description: "The LoopbackIps param.",
							Computed:    true,
							ElementType: types.StringType,
						},
						"tunnel_monitor_ip_address": dsschema.StringAttribute{
							Description: "The TunnelMonitorIpAddress param.",
							Computed:    true,
						},
					},
				},
			},
			"limit": dsschema.Int64Attribute{
				Description: "The Limit param. A limit of -1 will return all configured items. Default: `200`.",
				Optional:    true,
				Computed:    true,
			},
			"offset": dsschema.Int64Attribute{
				Description: "The Offset param. Default: `0`.",
				Optional:    true,
				Computed:    true,
			},
			"tfid": dsschema.StringAttribute{
				Description: "The Terraform ID.",
				Computed:    true,
			},
			"total": dsschema.Int64Attribute{
				Description: "The Total param.",
				Computed:    true,
			},
		},
	}
}

// Configure prepares the struct.
func (d *sharedInfrastructureSettingsListDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*scm.Client)
}

// Read performs Read for the struct.
func (d *sharedInfrastructureSettingsListDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state sharedInfrastructureSettingsListDsModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Basic logging.
	tflog.Info(ctx, "performing data source listing", map[string]any{
		"data_source_name":            "scm_shared_infrastructure_settings_list",
		"terraform_provider_function": "Read",
		"limit":                       state.Limit.ValueInt64(),
		"offset":                      state.Offset.ValueInt64(),
	})

	// Prepare to run the command.
	svc := xCzsxgj.NewClient(d.client)

	// Prepare input for the API endpoint.
	input := xCzsxgj.ListInput{}

	input.Limit = state.Limit.ValueInt64Pointer()

	input.Offset = state.Offset.ValueInt64Pointer()

	// Perform the operation.
	ans, err := svc.List(ctx, input)
	if err != nil {
		resp.Diagnostics.AddError("Error getting listing", err.Error())
		return
	}

	// Create the Terraform ID.
	var idBuilder strings.Builder
	if input.Limit != nil {
		idBuilder.WriteString(strconv.FormatInt(*input.Limit, 10))
	}

	idBuilder.WriteString(IdSeparator)
	if input.Offset != nil {
		idBuilder.WriteString(strconv.FormatInt(*input.Offset, 10))
	}

	// Store the answer to state.

	state.Tfid = types.StringValue(idBuilder.String())

	if len(ans.Data) == 0 {
		state.Data = nil
	} else {
		state.Data = make([]sharedInfrastructureSettingsListDsModel_eItpily_Config, 0, len(ans.Data))
		for _, var0 := range ans.Data {
			var1 := sharedInfrastructureSettingsListDsModel_eItpily_Config{}

			var1.ApiKey = types.StringPointerValue(var0.ApiKey)

			var1.CaptivePortalRedirectIpAddress = types.StringPointerValue(var0.CaptivePortalRedirectIpAddress)

			var1.EgressIpNotificationUrl = types.StringPointerValue(var0.EgressIpNotificationUrl)

			var1.InfraBgpAs = types.StringPointerValue(var0.InfraBgpAs)

			var1.InfrastructureSubnet = types.StringPointerValue(var0.InfrastructureSubnet)

			var1.InfrastructureSubnetIpv6 = types.StringPointerValue(var0.InfrastructureSubnetIpv6)

			var1.Ipv6 = types.BoolPointerValue(var0.Ipv6)

			var2, var3 := types.ListValueFrom(ctx, types.StringType, var0.LoopbackIps)
			var1.LoopbackIps = var2
			resp.Diagnostics.Append(var3.Errors()...)

			var1.TunnelMonitorIpAddress = types.StringPointerValue(var0.TunnelMonitorIpAddress)
			state.Data = append(state.Data, var1)
		}
	}

	state.Limit = types.Int64PointerValue(ans.Limit)

	state.Offset = types.Int64PointerValue(ans.Offset)

	state.Total = types.Int64PointerValue(ans.Total)

	// Done.
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}