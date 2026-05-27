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

	"github.com/paloaltonetworks/scm-go/generated/config_setup"
)

// Ensure the action implementation interfaces are satisfied.
var (
	_ action.Action              = &SnippetSnapshotUpdatesAction{}
	_ action.ActionWithConfigure = &SnippetSnapshotUpdatesAction{}
)

// SnippetSnapshotUpdatesAction implements the scm_snippet_snapshot_updates Terraform action.
type SnippetSnapshotUpdatesAction struct {
	client *config_setup.APIClient
}

// SnippetSnapshotUpdatesActionModel describes the action configuration data model.
type SnippetSnapshotUpdatesActionModel struct {
	Id       types.String `tfsdk:"id"`
	TenantId types.String `tfsdk:"tenant_id"`
}

func NewSnippetSnapshotUpdatesAction() action.Action {
	return &SnippetSnapshotUpdatesAction{}
}

func (a *SnippetSnapshotUpdatesAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_snippet_snapshot_updates"
}

func (a *SnippetSnapshotUpdatesAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Update Snippet Snapshots.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Id",
				Required:    true,
			},
			"tenant_id": schema.StringAttribute{
				Description: "Publisher Tenant ID",
				Required:    true,
			},
		},
	}
}

func (a *SnippetSnapshotUpdatesAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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

	client, ok := clients["config_setup"]
	if !ok {
		resp.Diagnostics.AddError(
			"Missing API Client",
			"The config_setup API client was not found in provider data.",
		)
		return
	}

	a.client = client.(*config_setup.APIClient)
}

func (a *SnippetSnapshotUpdatesAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var data SnippetSnapshotUpdatesActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Invoking action scm_snippet_snapshot_updates")
	resp.SendProgress(action.InvokeProgressEvent{
		Message: "Invoking action scm_snippet_snapshot_updates...",
	})
	// Build the SDK request
	sdkReq := config_setup.NewSnippetSnapshotSubscriberComparePayloadWithDefaults()

	if !data.Id.IsNull() && !data.Id.IsUnknown() {
		sdkReq.SetId(data.Id.ValueString())
	}
	if !data.TenantId.IsNull() && !data.TenantId.IsUnknown() {
		sdkReq.SetTenantId(data.TenantId.ValueString())
	}

	// Execute the API call
	apiReq := a.client.SnippetSnapshotsAPI.UpdateSnippetSnapshot(ctx)
	apiReq = apiReq.SnippetSnapshotSubscriberComparePayload(*sdkReq)
	_, httpResp, err := apiReq.Execute()
	if err != nil {
		errMsg := fmt.Sprintf("Failed to execute action scm_snippet_snapshot_updates: %s", err)
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

	completionMsg := fmt.Sprintf("Action scm_snippet_snapshot_updates completed successfully (HTTP %d)", httpStatus)

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
