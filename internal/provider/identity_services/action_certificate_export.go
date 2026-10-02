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

	"github.com/paloaltonetworks/scm-go/generated/identity_services"
)

// Ensure the action implementation interfaces are satisfied.
var (
	_ action.Action              = &CertificateExportAction{}
	_ action.ActionWithConfigure = &CertificateExportAction{}
)

// CertificateExportAction implements the scm_certificate_export Terraform action.
type CertificateExportAction struct {
	client *identity_services.APIClient
}

// CertificateExportActionModel describes the action configuration data model.
type CertificateExportActionModel struct {
	Format     types.String `tfsdk:"format"`
	Passphrase types.String `tfsdk:"passphrase"`
	Id         types.String `tfsdk:"id"`
}

func NewCertificateExportAction() action.Action {
	return &CertificateExportAction{}
}

func (a *CertificateExportAction) Metadata(_ context.Context, _ action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "scm_certificate_export"
}

func (a *CertificateExportAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Export a certificate.",
		Attributes: map[string]schema.Attribute{
			"format": schema.StringAttribute{
				Description: "Format. Possible values are `pkcs12`, `pem`, `der` and `pkcs10`.",
				Required:    true,
			},
			"passphrase": schema.StringAttribute{
				Description: "Passphrase",
				Optional:    true,
			},
			"id": schema.StringAttribute{
				Description: "The UUID of the configuration resource",
				Required:    true,
			},
		},
	}
}

func (a *CertificateExportAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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

	client, ok := clients["identity_services"]
	if !ok {
		resp.Diagnostics.AddError(
			"Missing API Client",
			"The identity_services API client was not found in provider data.",
		)
		return
	}

	a.client = client.(*identity_services.APIClient)
}

func (a *CertificateExportAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var data CertificateExportActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Invoking action scm_certificate_export")
	resp.SendProgress(action.InvokeProgressEvent{
		Message: "Invoking action scm_certificate_export...",
	})
	// Build the SDK request
	sdkReq := identity_services.NewExportCertificatePayloadWithDefaults()

	if !data.Format.IsNull() && !data.Format.IsUnknown() {
		sdkReq.SetFormat(data.Format.ValueString())
	}
	if !data.Passphrase.IsNull() && !data.Passphrase.IsUnknown() {
		sdkReq.SetPassphrase(data.Passphrase.ValueString())
	}

	// Execute the API call
	apiReq := a.client.CertificatesAPI.ExportCertificateByID(ctx, data.Id.ValueString())
	apiReq = apiReq.ExportCertificatePayload(*sdkReq)
	_, httpResp, err := apiReq.Execute()
	if err != nil {
		errMsg := fmt.Sprintf("Failed to execute action scm_certificate_export: %s", err)
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

	completionMsg := fmt.Sprintf("Action scm_certificate_export completed successfully (HTTP %d)", httpStatus)

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
