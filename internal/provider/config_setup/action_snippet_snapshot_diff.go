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
	_ action.Action              = &SnippetSnapshotDiffAction{}
	_ action.ActionWithConfigure = &SnippetSnapshotDiffAction{}
)

// SnippetSnapshotDiffAction implements the scm_snippet_snapshot_diff Terraform action.
type SnippetSnapshotDiffAction struct {
	client *config_setup.APIClient
}

// SnippetSnapshotDiffActionModel describes the action configuration data model.
type SnippetSnapshotDiffActionModel struct {
	ComparingVersion types.Int64  `tfsdk:"comparing_version"`
	ObjectId         types.String `tfsdk:"object_id"`
	SnippetId        types.String `tfsdk:"snippet_id"`
	Version          types.Int64  `tfsdk:"version"`
}

func NewSnippetSnapshotDiffAction() action.Action {
	return &SnippetSnapshotDiffAction{}
}

func (a *SnippetSnapshotDiffAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_snippet_snapshot_diff"
}

func (a *SnippetSnapshotDiffAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Diff Snippet Snapshots.",
		Attributes: map[string]schema.Attribute{
			"comparing_version": schema.Int64Attribute{
				Description: "Comparing version",
				Optional:    true,
			},
			"object_id": schema.StringAttribute{
				Description: "Object id",
				Required:    true,
			},
			"snippet_id": schema.StringAttribute{
				Description: "Snippet id",
				Required:    true,
			},
			"version": schema.Int64Attribute{
				Description: "Version",
				Required:    true,
			},
		},
	}
}

func (a *SnippetSnapshotDiffAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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

func (a *SnippetSnapshotDiffAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var data SnippetSnapshotDiffActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Invoking action scm_snippet_snapshot_diff")
	resp.SendProgress(action.InvokeProgressEvent{
		Message: "Invoking action scm_snippet_snapshot_diff...",
	})
	// Build the SDK request
	sdkReq := config_setup.NewCompareTloPayloadWithDefaults()

	if !data.ComparingVersion.IsNull() && !data.ComparingVersion.IsUnknown() {
		sdkReq.SetComparingVersion(int32(data.ComparingVersion.ValueInt64()))
	}
	if !data.ObjectId.IsNull() && !data.ObjectId.IsUnknown() {
		sdkReq.SetObjectId(data.ObjectId.ValueString())
	}
	if !data.SnippetId.IsNull() && !data.SnippetId.IsUnknown() {
		sdkReq.SetSnippetId(data.SnippetId.ValueString())
	}
	if !data.Version.IsNull() && !data.Version.IsUnknown() {
		sdkReq.SetVersion(int32(data.Version.ValueInt64()))
	}

	// Execute the API call
	apiReq := a.client.SnippetSnapshotsAPI.DiffSnippetSnapshot(ctx)
	apiReq = apiReq.CompareTloPayload(*sdkReq)
	_, httpResp, err := apiReq.Execute()
	if err != nil {
		errMsg := fmt.Sprintf("Failed to execute action scm_snippet_snapshot_diff: %s", err)
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

	completionMsg := fmt.Sprintf("Action scm_snippet_snapshot_diff completed successfully (HTTP %d)", httpStatus)

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
