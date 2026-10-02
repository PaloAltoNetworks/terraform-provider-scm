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
	_ action.Action              = &CertificateImportAction{}
	_ action.ActionWithConfigure = &CertificateImportAction{}
)

// CertificateImportAction implements the scm_certificate_import Terraform action.
type CertificateImportAction struct {
	client *identity_services.APIClient
}

// CertificateImportActionModel describes the action configuration data model.
type CertificateImportActionModel struct {
	CertificateFile types.String `tfsdk:"certificate_file"`
	Device          types.String `tfsdk:"device"`
	Folder          types.String `tfsdk:"folder"`
	Format          types.String `tfsdk:"format"`
	KeyFile         types.String `tfsdk:"key_file"`
	Name            types.String `tfsdk:"name"`
	Passphrase      types.String `tfsdk:"passphrase"`
	Snippet         types.String `tfsdk:"snippet"`
}

func NewCertificateImportAction() action.Action {
	return &CertificateImportAction{}
}

func (a *CertificateImportAction) Metadata(_ context.Context, _ action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "scm_certificate_import"
}

func (a *CertificateImportAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Import a certificate.",
		Attributes: map[string]schema.Attribute{
			"certificate_file": schema.StringAttribute{
				Description: "The Base64 encoded content of the certificate public key",
				Required:    true,
			},
			"device": schema.StringAttribute{
				Description: "The device in which the resource is defined",
				Optional:    true,
			},
			"folder": schema.StringAttribute{
				Description: "The folder in which the resource is defined",
				Optional:    true,
			},
			"format": schema.StringAttribute{
				Description: "Certificate format. Possible values are `pem`, `pkcs12` and `der`.",
				Required:    true,
			},
			"key_file": schema.StringAttribute{
				Description: "The Base64 encoded content of the certificate private key",
				Optional:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the certificate",
				Required:    true,
			},
			"passphrase": schema.StringAttribute{
				Description: "Passphrase to protect the certificate private key",
				Optional:    true,
			},
			"snippet": schema.StringAttribute{
				Description: "The snippet in which the resource is defined",
				Optional:    true,
			},
		},
	}
}

func (a *CertificateImportAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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

func (a *CertificateImportAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var data CertificateImportActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Invoking action scm_certificate_import")
	resp.SendProgress(action.InvokeProgressEvent{
		Message: "Invoking action scm_certificate_import...",
	})
	// Build the SDK request
	sdkReq := identity_services.NewCertificatesImportWithDefaults()

	if !data.CertificateFile.IsNull() && !data.CertificateFile.IsUnknown() {
		sdkReq.SetCertificateFile(data.CertificateFile.ValueString())
	}
	if !data.Device.IsNull() && !data.Device.IsUnknown() {
		sdkReq.SetDevice(data.Device.ValueString())
	}
	if !data.Folder.IsNull() && !data.Folder.IsUnknown() {
		sdkReq.SetFolder(data.Folder.ValueString())
	}
	if !data.Format.IsNull() && !data.Format.IsUnknown() {
		sdkReq.SetFormat(data.Format.ValueString())
	}
	if !data.KeyFile.IsNull() && !data.KeyFile.IsUnknown() {
		sdkReq.SetKeyFile(data.KeyFile.ValueString())
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		sdkReq.SetName(data.Name.ValueString())
	}
	if !data.Passphrase.IsNull() && !data.Passphrase.IsUnknown() {
		sdkReq.SetPassphrase(data.Passphrase.ValueString())
	}
	if !data.Snippet.IsNull() && !data.Snippet.IsUnknown() {
		sdkReq.SetSnippet(data.Snippet.ValueString())
	}

	// Execute the API call
	apiReq := a.client.CertificatesAPI.ImportCertificates(ctx)
	apiReq = apiReq.CertificatesImport(*sdkReq)
	_, httpResp, err := apiReq.Execute()
	if err != nil {
		errMsg := fmt.Sprintf("Failed to execute action scm_certificate_import: %s", err)
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

	completionMsg := fmt.Sprintf("Action scm_certificate_import completed successfully (HTTP %d)", httpStatus)

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
