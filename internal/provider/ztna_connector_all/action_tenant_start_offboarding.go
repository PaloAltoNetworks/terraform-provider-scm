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

	"github.com/paloaltonetworks/scm-go/generated/ztna_connector_all"
)

// Ensure the action implementation interfaces are satisfied.
var (
	_ action.Action              = &TenantStartOffboardingAction{}
	_ action.ActionWithConfigure = &TenantStartOffboardingAction{}
)

// TenantStartOffboardingAction implements the ztna_tenant_start_offboarding Terraform action.
type TenantStartOffboardingAction struct {
	client *ztna_connector_all.APIClient
}

// TenantStartOffboardingActionModel describes the action configuration data model.
type TenantStartOffboardingActionModel struct {
}

func NewTenantStartOffboardingAction() action.Action {
	return &TenantStartOffboardingAction{}
}

func (a *TenantStartOffboardingAction) Metadata(_ context.Context, _ action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "ztna_tenant_start_offboarding"
}

func (a *TenantStartOffboardingAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Initiates tenant deletion. This marks the tenant for deletion and triggers the cleanup process.",
		Attributes:  map[string]schema.Attribute{},
	}
}

func (a *TenantStartOffboardingAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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

	client, ok := clients["ztna_connector_all"]
	if !ok {
		resp.Diagnostics.AddError(
			"Missing API Client",
			"The ztna_connector_all API client was not found in provider data.",
		)
		return
	}

	a.client = client.(*ztna_connector_all.APIClient)
}

func (a *TenantStartOffboardingAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var data TenantStartOffboardingActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Invoking action scm_tenant_start_offboarding")
	resp.SendProgress(action.InvokeProgressEvent{
		Message: "Invoking action scm_tenant_start_offboarding...",
	})
	// Execute the API call (no request body)
	apiReq := a.client.TenantAPI.StartTenantOffboarding(ctx)
	_, httpResp, err := apiReq.Execute()
	if err != nil {
		errMsg := fmt.Sprintf("Failed to execute action scm_tenant_start_offboarding: %s", err)
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

	completionMsg := fmt.Sprintf("Action scm_tenant_start_offboarding completed successfully (HTTP %d)", httpStatus)

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
