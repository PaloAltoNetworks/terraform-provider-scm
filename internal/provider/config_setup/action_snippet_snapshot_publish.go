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
	_ action.Action              = &SnippetSnapshotPublishAction{}
	_ action.ActionWithConfigure = &SnippetSnapshotPublishAction{}
)

// SnippetSnapshotPublishAction implements the scm_snippet_snapshot_publish Terraform action.
type SnippetSnapshotPublishAction struct {
	client *config_setup.APIClient
}

// SnippetSnapshotPublishActionModel describes the action configuration data model.
type SnippetSnapshotPublishActionModel struct {
	Id         types.String `tfsdk:"id"`
	Name       types.String `tfsdk:"name"`
	Tsgs       types.List   `tfsdk:"tsgs"`
	Validation types.Bool   `tfsdk:"validation"`
	Version    types.Int64  `tfsdk:"version"`
}

func NewSnippetSnapshotPublishAction() action.Action {
	return &SnippetSnapshotPublishAction{}
}

func (a *SnippetSnapshotPublishAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_snippet_snapshot_publish"
}

func (a *SnippetSnapshotPublishAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Publish Snippet Snapshots.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Id",
				Optional:    true,
			},
			"name": schema.StringAttribute{
				Description: "Name",
				Optional:    true,
			},
			"tsgs": schema.ListAttribute{
				Description: "Tsgs",
				Optional:    true,
				ElementType: types.StringType,
			},
			"validation": schema.BoolAttribute{
				Description: "Validation",
				Optional:    true,
			},
			"version": schema.Int64Attribute{
				Description: "Version",
				Optional:    true,
			},
		},
	}
}

func (a *SnippetSnapshotPublishAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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

func (a *SnippetSnapshotPublishAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var data SnippetSnapshotPublishActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Invoking action scm_snippet_snapshot_publish")
	resp.SendProgress(action.InvokeProgressEvent{
		Message: "Invoking action scm_snippet_snapshot_publish...",
	})
	// Build the SDK request
	sdkReq := config_setup.NewSnippetSnapshotPublishRequestWithDefaults()

	if !data.Id.IsNull() && !data.Id.IsUnknown() {
		sdkReq.SetId(data.Id.ValueString())
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		sdkReq.SetName(data.Name.ValueString())
	}
	if !data.Tsgs.IsNull() && !data.Tsgs.IsUnknown() {
		var vals []string
		// ElementsAs failures already include attribute path information.
		_ = data.Tsgs.ElementsAs(ctx, &vals, false)
		sdkReq.SetTsgs(vals)
	}
	if !data.Validation.IsNull() && !data.Validation.IsUnknown() {
		sdkReq.SetValidation(data.Validation.ValueBool())
	}
	if !data.Version.IsNull() && !data.Version.IsUnknown() {
		sdkReq.SetVersion(int32(data.Version.ValueInt64()))
	}

	// Execute the API call
	apiReq := a.client.SnippetSnapshotsAPI.PublishSnippetSnapshot(ctx)
	apiReq = apiReq.SnippetSnapshotPublishRequest(*sdkReq)
	_, httpResp, err := apiReq.Execute()
	if err != nil {
		errMsg := fmt.Sprintf("Failed to execute action scm_snippet_snapshot_publish: %s", err)
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

	completionMsg := fmt.Sprintf("Action scm_snippet_snapshot_publish completed successfully (HTTP %d)", httpStatus)

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
