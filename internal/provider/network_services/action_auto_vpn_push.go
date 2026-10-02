package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/hashicorp/terraform-plugin-framework/action"
	"github.com/hashicorp/terraform-plugin-framework/action/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/paloaltonetworks/scm-go/generated/network_services"
)

// Ensure the action implementation interfaces are satisfied.
var (
	_ action.Action              = &AutoVpnPushAction{}
	_ action.ActionWithConfigure = &AutoVpnPushAction{}
)

// Nested object types used by the AutoVpnPushAction model. These are declared
// in this file (rather than imported) because the action's request body
// types are not generated as resource models.
type AutoVpnPushAutoVpnDevicesItem struct {
	Name       types.String `tfsdk:"name"`
	RefreshPsk types.Bool   `tfsdk:"refresh_psk"`
}

// AutoVpnPushAction implements the scm_auto_vpn_push Terraform action.
type AutoVpnPushAction struct {
	client *network_services.APIClient
}

// AutoVpnPushActionModel describes the action configuration data model.
type AutoVpnPushActionModel struct {
	AutoVpnDevices []AutoVpnPushAutoVpnDevicesItem `tfsdk:"auto_vpn_devices"`
}

func NewAutoVpnPushAction() action.Action {
	return &AutoVpnPushAction{}
}

func (a *AutoVpnPushAction) Metadata(_ context.Context, _ action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "scm_auto_vpn_push"
}

func (a *AutoVpnPushAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Push Auto VPN configs.",
		Attributes: map[string]schema.Attribute{
			"auto_vpn_devices": schema.ListNestedAttribute{
				Description: "VPN clusters",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Description: "VPN cluster to push to",
							Optional:    true,
						},
						"refresh_psk": schema.BoolAttribute{
							Description: "Refresh psk",
							Optional:    true,
						},
					},
				},
			},
		},
	}
}

func (a *AutoVpnPushAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	clients, ok := req.ProviderData.(map[string]interface{})
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Provider Data",
			"Expected map[string]interface{} from provider, got unexpected type.",
		)
		return
	}

	client, ok := clients["network_services"]
	if !ok {
		resp.Diagnostics.AddError(
			"Missing API Client",
			"The network_services API client was not found in provider data.",
		)
		return
	}

	a.client = client.(*network_services.APIClient)
}

func (a *AutoVpnPushAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var data AutoVpnPushActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Invoking action scm_auto_vpn_push")
	resp.SendProgress(action.InvokeProgressEvent{
		Message: "Invoking action scm_auto_vpn_push...",
	})
	// Build the SDK request
	sdkReq := network_services.NewAutoVpnPushConfigWithDefaults()

	if data.AutoVpnDevices != nil {
		items := make([]network_services.AutoVpnPushConfigAutoVpnDevicesInner, 0, len(data.AutoVpnDevices))
		for _, item := range data.AutoVpnDevices {
			sdkItem := network_services.AutoVpnPushConfigAutoVpnDevicesInner{}
			sdkItem.Name = item.Name.ValueStringPointer()
			sdkItem.RefreshPsk = item.RefreshPsk.ValueBoolPointer()
			items = append(items, sdkItem)
		}
		sdkReq.SetAutoVpnDevices(items)
	}

	// Execute the API call
	apiReq := a.client.AutoVPNConfigPushAPI.CreateAutoVPNPushConfigs(ctx)
	apiReq = apiReq.AutoVpnPushConfig(*sdkReq)
	_, httpResp, err := apiReq.Execute()
	if err != nil {
		errMsg := fmt.Sprintf("Failed to execute action scm_auto_vpn_push: %s", err)
		if httpResp != nil && httpResp.Body != nil {
			body, readErr := io.ReadAll(httpResp.Body)
			if readErr == nil && len(body) > 0 {
				errMsg = fmt.Sprintf("%s\nResponse body: %s", errMsg, string(body))
			}
		}
		resp.Diagnostics.AddError("Action Failed", errMsg)
		return
	}

	httpStatus := 0
	if httpResp != nil {
		httpStatus = httpResp.StatusCode
	}

	completionMsg := fmt.Sprintf("Action scm_auto_vpn_push completed successfully (HTTP %d)", httpStatus)

	if httpResp != nil && httpResp.Body != nil {
		body, readErr := io.ReadAll(httpResp.Body)
		if readErr == nil && len(body) > 0 {
			var pretty bytes.Buffer
			if indentErr := json.Indent(&pretty, body, "", "  "); indentErr == nil {
				completionMsg = fmt.Sprintf("%s\nResponse:\n%s", completionMsg, pretty.String())
			} else {
				completionMsg = fmt.Sprintf("%s\nResponse: %s", completionMsg, string(body))
			}
		}
	}

	tflog.Info(ctx, completionMsg)
	resp.SendProgress(action.InvokeProgressEvent{
		Message: completionMsg,
	})

	// Suppress unused-import if no scalar fields needed types.*
	_ = types.StringNull
}
