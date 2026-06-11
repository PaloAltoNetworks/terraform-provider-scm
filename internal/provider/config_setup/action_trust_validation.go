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
	_ action.Action              = &TrustValidationAction{}
	_ action.ActionWithConfigure = &TrustValidationAction{}
)

// TrustValidationAction implements the scm_trust_validation Terraform action.
type TrustValidationAction struct {
	client *config_setup.APIClient
}

// TrustValidationActionModel describes the action configuration data model.
type TrustValidationActionModel struct {
	DonorTenantName     types.String `tfsdk:"donor_tenant_name"`
	Psk                 types.String `tfsdk:"psk"`
	RecipientTenantName types.String `tfsdk:"recipient_tenant_name"`
	TrustId             types.Int64  `tfsdk:"trust_id"`
	Tsg                 types.String `tfsdk:"tsg"`
}

func NewTrustValidationAction() action.Action {
	return &TrustValidationAction{}
}

func (a *TrustValidationAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_trust_validation"
}

func (a *TrustValidationAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Validate trust.",
		Attributes: map[string]schema.Attribute{
			"donor_tenant_name": schema.StringAttribute{
				Description: "Donor tenant name",
				Required:    true,
			},
			"psk": schema.StringAttribute{
				Description: "Psk",
				Required:    true,
			},
			"recipient_tenant_name": schema.StringAttribute{
				Description: "Recipient tenant name",
				Required:    true,
			},
			"trust_id": schema.Int64Attribute{
				Description: "Trust id",
				Required:    true,
			},
			"tsg": schema.StringAttribute{
				Description: "Tsg",
				Required:    true,
			},
		},
	}
}

func (a *TrustValidationAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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

func (a *TrustValidationAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var data TrustValidationActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Invoking action scm_trust_validation")
	resp.SendProgress(action.InvokeProgressEvent{
		Message: "Invoking action scm_trust_validation...",
	})
	// Build the SDK request
	sdkReq := config_setup.NewTrustsValidationPayloadWithDefaults()

	if !data.DonorTenantName.IsNull() && !data.DonorTenantName.IsUnknown() {
		sdkReq.SetDonorTenantName(data.DonorTenantName.ValueString())
	}
	if !data.Psk.IsNull() && !data.Psk.IsUnknown() {
		sdkReq.SetPsk(data.Psk.ValueString())
	}
	if !data.RecipientTenantName.IsNull() && !data.RecipientTenantName.IsUnknown() {
		sdkReq.SetRecipientTenantName(data.RecipientTenantName.ValueString())
	}
	if !data.TrustId.IsNull() && !data.TrustId.IsUnknown() {
		sdkReq.SetTrustId(int32(data.TrustId.ValueInt64()))
	}
	if !data.Tsg.IsNull() && !data.Tsg.IsUnknown() {
		sdkReq.SetTsg(data.Tsg.ValueString())
	}

	// Execute the API call
	apiReq := a.client.TrustValidationsAPI.ValidateTrust(ctx)
	apiReq = apiReq.TrustsValidationPayload(*sdkReq)
	_, httpResp, err := apiReq.Execute()
	if err != nil {
		errMsg := fmt.Sprintf("Failed to execute action scm_trust_validation: %s", err)
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

	completionMsg := fmt.Sprintf("Action scm_trust_validation completed successfully (HTTP %d)", httpStatus)

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
