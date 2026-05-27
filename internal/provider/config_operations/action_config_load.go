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

	"github.com/paloaltonetworks/scm-go/generated/config_operations"
)

// Ensure the action implementation interfaces are satisfied.
var (
	_ action.Action              = &ConfigLoadAction{}
	_ action.ActionWithConfigure = &ConfigLoadAction{}
)

// ConfigLoadAction implements the scm_config_load Terraform action.
type ConfigLoadAction struct {
	client *config_operations.APIClient
}

// ConfigLoadActionModel describes the action configuration data model.
type ConfigLoadActionModel struct {
	Version types.Int64 `tfsdk:"version"`
}

func NewConfigLoadAction() action.Action {
	return &ConfigLoadAction{}
}

func (a *ConfigLoadAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_config_load"
}

func (a *ConfigLoadAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Load a specific configuration version into the candidate configuration.",
		Attributes: map[string]schema.Attribute{
			"version": schema.Int64Attribute{
				Description: "Version",
				Optional:    true,
			},
		},
	}
}

func (a *ConfigLoadAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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

	client, ok := clients["config_operations"]
	if !ok {
		resp.Diagnostics.AddError(
			"Missing API Client",
			"The config_operations API client was not found in provider data.",
		)
		return
	}

	a.client = client.(*config_operations.APIClient)
}

func (a *ConfigLoadAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var data ConfigLoadActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Invoking action scm_config_load")
	resp.SendProgress(action.InvokeProgressEvent{
		Message: "Invoking action scm_config_load...",
	})
	// Build the SDK request
	sdkReq := config_operations.NewLoadConfigWithDefaults()

	if !data.Version.IsNull() && !data.Version.IsUnknown() {
		sdkReq.SetVersion(int32(data.Version.ValueInt64()))
	}

	// Execute the API call
	apiReq := a.client.ConfigVersionsAPI.LoadConfigVersions(ctx)
	apiReq = apiReq.LoadConfig(*sdkReq)
	httpResp, err := apiReq.Execute()
	if err != nil {
		errMsg := fmt.Sprintf("Failed to execute action scm_config_load: %s", err)
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

	completionMsg := fmt.Sprintf("Action scm_config_load completed successfully (HTTP %d)", httpStatus)

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
