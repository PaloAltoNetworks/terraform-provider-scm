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

	"github.com/paloaltonetworks/scm-go/generated/deployment_services"
)

// Ensure the action implementation interfaces are satisfied.
var (
	_ action.Action              = &ApplicationDefaultsAction{}
	_ action.ActionWithConfigure = &ApplicationDefaultsAction{}
)

// ApplicationDefaultsAction implements the scm_application_defaults Terraform action.
type ApplicationDefaultsAction struct {
	client *deployment_services.APIClient
}

// ApplicationDefaultsActionModel describes the action configuration data model.
type ApplicationDefaultsActionModel struct {
}

func NewApplicationDefaultsAction() action.Action {
	return &ApplicationDefaultsAction{}
}

func (a *ApplicationDefaultsAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_application_defaults"
}

func (a *ApplicationDefaultsAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Create Prisma Access application defaults. *These application defaults are normally created in the UI. This endpoint is necessary for customers that do not use the UI to create these application defaults such as certificates and configuration nodes. This endpoint will be deprecated once the UI dependencies have been eliminated.*",
		Attributes:  map[string]schema.Attribute{},
	}
}

func (a *ApplicationDefaultsAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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

	client, ok := clients["deployment_services"]
	if !ok {
		resp.Diagnostics.AddError(
			"Missing API Client",
			"The deployment_services API client was not found in provider data.",
		)
		return
	}

	a.client = client.(*deployment_services.APIClient)
}

func (a *ApplicationDefaultsAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var data ApplicationDefaultsActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Invoking action scm_application_defaults")
	resp.SendProgress(action.InvokeProgressEvent{
		Message: "Invoking action scm_application_defaults...",
	})
	// Execute the API call (no request body)
	apiReq := a.client.ApplicationDefaultsAPI.CreateApplicationDefaults(ctx)
	httpResp, err := apiReq.Execute()
	if err != nil {
		errMsg := fmt.Sprintf("Failed to execute action scm_application_defaults: %s", err)
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

	completionMsg := fmt.Sprintf("Action scm_application_defaults completed successfully (HTTP %d)", httpStatus)

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
