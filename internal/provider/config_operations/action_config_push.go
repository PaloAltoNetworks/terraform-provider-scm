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
	_ action.Action              = &ConfigPushAction{}
	_ action.ActionWithConfigure = &ConfigPushAction{}
)

// ConfigPushAction implements the scm_config_push Terraform action.
type ConfigPushAction struct {
	client *config_operations.APIClient
}

// ConfigPushActionModel describes the action configuration data model.
type ConfigPushActionModel struct {
	Admin       types.List   `tfsdk:"admin"`
	Description types.String `tfsdk:"description"`
	Devices     types.List   `tfsdk:"devices"`
	Folder      types.List   `tfsdk:"folder"`
}

func NewConfigPushAction() action.Action {
	return &ConfigPushAction{}
}

func (a *ConfigPushAction) Metadata(_ context.Context, _ action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "scm_config_push"
}

func (a *ConfigPushAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Push the candidate configuration.",
		Attributes: map[string]schema.Attribute{
			"admin": schema.ListAttribute{
				Description: "Admin",
				Optional:    true,
				ElementType: types.StringType,
			},
			"description": schema.StringAttribute{
				Description: "Description",
				Optional:    true,
			},
			"devices": schema.ListAttribute{
				Description: "Devices",
				Optional:    true,
				ElementType: types.StringType,
			},
			"folder": schema.ListAttribute{
				Description: "Folder",
				Optional:    true,
				ElementType: types.StringType,
			},
		},
	}
}

func (a *ConfigPushAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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

func (a *ConfigPushAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var data ConfigPushActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Invoking action scm_config_push")
	resp.SendProgress(action.InvokeProgressEvent{
		Message: "Invoking action scm_config_push...",
	})
	// Build the SDK request
	sdkReq := config_operations.NewPushCandidateConfigVersionsRequestWithDefaults()

	if !data.Admin.IsNull() && !data.Admin.IsUnknown() {
		var vals []string
		// ElementsAs failures already include attribute path information.
		_ = data.Admin.ElementsAs(ctx, &vals, false)
		sdkReq.SetAdmin(vals)
	}
	if !data.Description.IsNull() && !data.Description.IsUnknown() {
		sdkReq.SetDescription(data.Description.ValueString())
	}
	if !data.Devices.IsNull() && !data.Devices.IsUnknown() {
		var vals []string
		// ElementsAs failures already include attribute path information.
		_ = data.Devices.ElementsAs(ctx, &vals, false)
		sdkReq.SetDevices(vals)
	}
	if !data.Folder.IsNull() && !data.Folder.IsUnknown() {
		var vals []string
		// ElementsAs failures already include attribute path information.
		_ = data.Folder.ElementsAs(ctx, &vals, false)
		sdkReq.SetFolder(vals)
	}

	// Execute the API call
	apiReq := a.client.ConfigVersionsAPI.PushCandidateConfigVersions(ctx)
	apiReq = apiReq.PushCandidateConfigVersionsRequest(*sdkReq)
	httpResp, err := apiReq.Execute()
	if err != nil {
		errMsg := fmt.Sprintf("Failed to execute action scm_config_push: %s", err)
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

	completionMsg := fmt.Sprintf("Action scm_config_push completed successfully (HTTP %d)", httpStatus)

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
