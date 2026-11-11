package provider

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	setup "github.com/paloaltonetworks/scm-go"
	"github.com/paloaltonetworks/scm-go/generated/cloud_ngfw"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/cloud_ngfw"
	"github.com/paloaltonetworks/terraform-provider-scm/internal/utils"
)

const (
	cngfwCreateDefaultTimeout = 60 * time.Minute
	// Measured firewall deletes take 12-19 min (CloudFormation stack teardown is 11-16 min).
	cngfwDeleteDefaultTimeout = 30 * time.Minute
)

// RESOURCE for SCM Ngfirewall (Package: cloud_ngfw)
var (
	_ resource.Resource                = &NgfirewallResource{}
	_ resource.ResourceWithConfigure   = &NgfirewallResource{}
	_ resource.ResourceWithImportState = &NgfirewallResource{}
)

func NewNgfirewallResource() resource.Resource {
	return &NgfirewallResource{}
}

// NgfirewallResource defines the resource implementation.
type NgfirewallResource struct {
	client      *cloud_ngfw.APIClient
	setupClient *setup.Client
}

func (r *NgfirewallResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "scm_cngfw_ngfirewall"
}

func (r *NgfirewallResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = models.NgfirewallCreateRequestResourceSchema
}

func (r *NgfirewallResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	clients, ok := req.ProviderData.(map[string]interface{})
	if !ok {
		resp.Diagnostics.AddError("Unexpected Resource Configure Type", fmt.Sprintf("Expected map[string]interface{}, got: %T.", req.ProviderData))
		return
	}
	client, ok := clients["cloud_ngfw"].(*cloud_ngfw.APIClient)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Client Type", fmt.Sprintf("Expected *cloud_ngfw.APIClient for 'cloud_ngfw' client."))
		return
	}
	r.client = client

	sc, ok := clients["setup"].(*setup.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Setup Client Type", "Expected *setup.Client for 'setup' client.")
		return
	}
	r.setupClient = sc
}

func (r *NgfirewallResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Starting Create function for Ngfirewall")
	var data models.NgfirewallCreateRequest

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	planObject, diags := types.ObjectValueFrom(ctx, models.NgfirewallCreateRequest{}.AttrTypes(), &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	_ = planObject
	unpackedScmObject, diags := unpackNgfirewallCreateRequestToSdk(ctx, planObject)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Always tag creates from the Terraform provider. The SDK wrapper checks for this
	// value and will not override it with CngfwSourceSDK.
	unpackedScmObject.SetSource(cloud_ngfw.CngfwSourceTerraformProvider)

	tflog.Debug(ctx, "Creating ngfirewalls on SCM API")

	// CreateCloudNGFirewall auto-fetches LinkId from the tenant SCM API when
	// the caller has not supplied one; otherwise it uses the provided value as-is.
	createdObject, _, err := r.client.NgfirewallsAPI.CreateCloudNGFirewall(ctx, *unpackedScmObject, data.Region.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error creating ngfirewalls", err.Error())
		detailedMessage := utils.PrintScmError(err)
		resp.Diagnostics.AddError("SCM Resource Creation Failed: API Request Failed", detailedMessage)
		return
	}
	// Extract the firewall ID from the nested create response.
	if createdObject.Response != nil && createdObject.Response.FirewallId != nil {
		data.Id = types.StringValue(*createdObject.Response.FirewallId)
	}

	// Restore operation parameters (e.g. region) from the plan — not returned in the response.
	_ = req.Plan.GetAttribute(ctx, path.Root("region"), &data.Region)

	// Build TFID: region:firewallId
	var idBuilder strings.Builder
	idBuilder.WriteString(data.Region.ValueString())
	idBuilder.WriteString(":")
	idBuilder.WriteString(data.Id.ValueString())
	data.Tfid = types.StringValue(idBuilder.String())

	// Wait for provisioning to complete before saving state.
	createTimeout := cngfwCreateDefaultTimeout
	createStart := time.Now()
	tflog.Debug(ctx, "Waiting for CNGFW firewall to reach CREATE_COMPLETE", map[string]interface{}{
		"firewall_id": data.Id.ValueString(),
		"region":      data.Region.ValueString(),
		"timeout":     createTimeout.String(),
	})
	if waitErr := waitForCngfwStatus(ctx, r.client, r.setupClient, data.Id.ValueString(), data.Region.ValueString(), createTimeout, "CREATE_COMPLETE", "UPDATE_COMPLETE"); waitErr != nil {
		resp.Diagnostics.AddError("CNGFW Create: provisioning failed", waitErr.Error())
		return
	}
	tflog.Info(ctx, "CNGFW firewall reached CREATE_COMPLETE", map[string]interface{}{
		"firewall_id": data.Id.ValueString(),
		"mode":        data.Mode.ValueString(),
		"elapsed":     time.Since(createStart).Round(time.Second).String(),
	})

	// Re-read the firewall so every Computed attribute has a known value before
	// state is written. Terraform rejects an apply that leaves a Computed
	// attribute unknown, and the create response carries only the firewall ID --
	// server-assigned values (Tier, LinkId, ChangeProtection, EgressNAT,
	// Endpoints, GwlbTcpIdleTimeout) are only available from a GET.
	getRes, _, getErr := r.client.NgfirewallsAPI.GetNgfirewallsByID(ctx, data.Id.ValueString()).Region(data.Region.ValueString()).Execute()
	if getErr != nil {
		resp.Diagnostics.AddError(
			"CNGFW Create: post-create read failed",
			"The firewall was created but reading it back failed, so Computed attributes cannot be resolved. "+
				"Re-run terraform plan to reconcile state. Error: "+getErr.Error(),
		)
		return
	}

	// Anything the GET does not return falls back to the planned values.
	savestate := data

	// Populate state from the GET response. Generated from the NgfirewallCreateRequest and NgfirewallFirewall models:
	// fields the GET does not return keep their planned/prior value (null on fresh import).
	if getRes != nil && getRes.Response != nil && getRes.Response.Firewall != nil {
		fwObj, fwDiags := packNgfirewallFirewallFromSdk(ctx, *getRes.Response.Firewall)
		resp.Diagnostics.Append(fwDiags...)
		var fw models.NgfirewallFirewall
		resp.Diagnostics.Append(fwObj.As(ctx, &fw, basetypes.ObjectAsOptions{})...)
		if resp.Diagnostics.HasError() {
			return
		}
		data.AllowListAccounts = fw.AllowListAccounts
		data.ChangeProtection = fw.ChangeProtection
		data.CustomerZoneIdList = fw.CustomerZoneIdList
		data.Description = fw.Description
		data.EgressNAT = fw.EgressNAT
		data.Endpoints = fw.Endpoints
		{
			var nested models.NgfirewallGwlbConfig
			if !fw.GWLB.IsNull() && !fw.GWLB.IsUnknown() {
				resp.Diagnostics.Append(fw.GWLB.As(ctx, &nested, basetypes.ObjectAsOptions{})...)
			}
			data.GwlbTcpIdleTimeout = nested.TCPIdleTimeout
		}
		data.LinkId = fw.LinkId
		data.PrivateAccess = fw.PrivateAccess
		data.Tags = fw.Tags
		data.Tier = fw.Tier
		data.UserID = fw.UserID
		if !fw.FirewallId.IsNull() {
			data.Id = fw.FirewallId
		}
	} else {
		// GET succeeded but contained no Firewall object - unexpected; restore from prior state.
		data = savestate
	}

	// Fields not returned by GET, and identity fields established above.
	data.Mode = savestate.Mode
	data.ScmLabels = savestate.ScmLabels
	data.Region = savestate.Region
	data.Tfid = savestate.Tfid

	tflog.Debug(ctx, "Created ngfirewalls", map[string]interface{}{"tfid": data.Tfid.ValueString()})
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NgfirewallResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	tflog.Debug(ctx, "Starting Read function for Ngfirewall")
	var data, savestate models.NgfirewallCreateRequest

	resp.Diagnostics.Append(req.State.Get(ctx, &savestate)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tokens := strings.Split(savestate.Tfid.ValueString(), ":")
	if len(tokens) != 2 {
		resp.Diagnostics.AddError("Error parsing TFID", fmt.Sprintf("Expected a TFID with 2 parts separated by ':', but got %d parts for TFID %s", len(tokens), savestate.Tfid.ValueString()))
		return
	}
	objectId := tokens[1]

	tflog.Debug(ctx, "Reading ngfirewalls from SCM API", map[string]interface{}{"id": objectId})
	getReq := r.client.NgfirewallsAPI.GetNgfirewallsByID(ctx, objectId)
	if tokens[0] != "" {
		getReq = getReq.Region(tokens[0])
	}
	getRes, httpErr, err := getReq.Execute()
	if err != nil {
		// The CNGFW API returns 400 (not 404) for an unknown object ID, with a
		// non-standard body: {"ResponseStatus":{"ErrorCode":N,"Reason":"... does not
		// exist ..."}}. Treat that as "gone" so an object deleted outside Terraform
		// gets recreated instead of failing the plan. Deliberately narrow: if the
		// wording changes this stops matching and we fall back to surfacing the
		// error, which is the prior behavior -- it fails closed and never drops
		// live state.
		notFound := httpErr != nil && httpErr.StatusCode == http.StatusNotFound

		if !notFound && httpErr != nil && httpErr.StatusCode == http.StatusBadRequest {
			// The reason text lives in the raw response body, not err.Error().
			detail := err.Error()
			if scmErr, ok := err.(utils.SCMErrorAccessor); ok {
				detail = string(scmErr.Body())
			}
			notFound = strings.Contains(strings.ToLower(detail), "does not exist")
		}

		if notFound {
			tflog.Debug(ctx, "Got no ngfirewalls on read SCM API. Remove from state to let terraform create", map[string]interface{}{"id": objectId})
			resp.State.RemoveResource(ctx)
		} else {
			tflog.Debug(ctx, "Got an exception on read SCM API. ", map[string]interface{}{"id": objectId})
			resp.Diagnostics.AddError("Error reading ngfirewalls", err.Error())
			detailedMessage := utils.PrintScmError(err)
			resp.Diagnostics.AddError("SCM Resource Read Failed: API Request Failed", detailedMessage)
		}
		return
	}

	// Populate state from the GET response. Generated from the NgfirewallCreateRequest and NgfirewallFirewall models:
	// fields the GET does not return keep their planned/prior value (null on fresh import).
	if getRes != nil && getRes.Response != nil && getRes.Response.Firewall != nil {
		fwObj, fwDiags := packNgfirewallFirewallFromSdk(ctx, *getRes.Response.Firewall)
		resp.Diagnostics.Append(fwDiags...)
		var fw models.NgfirewallFirewall
		resp.Diagnostics.Append(fwObj.As(ctx, &fw, basetypes.ObjectAsOptions{})...)
		if resp.Diagnostics.HasError() {
			return
		}
		data.AllowListAccounts = fw.AllowListAccounts
		data.ChangeProtection = fw.ChangeProtection
		data.CustomerZoneIdList = fw.CustomerZoneIdList
		data.Description = fw.Description
		data.EgressNAT = fw.EgressNAT
		data.Endpoints = fw.Endpoints
		{
			var nested models.NgfirewallGwlbConfig
			if !fw.GWLB.IsNull() && !fw.GWLB.IsUnknown() {
				resp.Diagnostics.Append(fw.GWLB.As(ctx, &nested, basetypes.ObjectAsOptions{})...)
			}
			data.GwlbTcpIdleTimeout = nested.TCPIdleTimeout
		}
		data.LinkId = fw.LinkId
		data.PrivateAccess = fw.PrivateAccess
		data.Tags = fw.Tags
		data.Tier = fw.Tier
		data.UserID = fw.UserID
		if !fw.FirewallId.IsNull() {
			data.Id = fw.FirewallId
		}
	} else {
		// GET succeeded but contained no Firewall object - unexpected; restore from prior state.
		data = savestate
	}

	// Fields not returned by GET — preserve from savestate (will be null on fresh import).
	// Source is set by the provider and is not a user-configurable attribute.
	data.Mode = savestate.Mode
	data.ScmLabels = savestate.ScmLabels

	// Restore region from TFID token[0].
	if tokens[0] != "" {
		data.Region = basetypes.NewStringValue(tokens[0])
	}

	// Carry over tfid from state back into data.
	data.Tfid = savestate.Tfid

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NgfirewallResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Starting Update function for Ngfirewall")
	var plan, state models.NgfirewallCreateRequest

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	planObject, diags := types.ObjectValueFrom(ctx, models.NgfirewallCreateRequest{}.AttrTypes(), &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	_ = planObject

	tokens := strings.Split(state.Tfid.ValueString(), ":")
	if len(tokens) != 2 {
		resp.Diagnostics.AddError("Error parsing TFID", fmt.Sprintf("Expected a TFID with 2 parts separated by ':', but got %d parts for TFID %s", len(tokens), state.Tfid.ValueString()))
		return
	}
	objectId := tokens[1]
	regionVal := tokens[0]

	tflog.Debug(ctx, "Updating ngfirewalls on SCM API", map[string]interface{}{"id": objectId})

	// CNGFW requires UpdateToken and DeploymentUpdateToken for optimistic locking.
	// Fetch current state to get the tokens.
	tflog.Debug(ctx, "Fetching current firewall state to obtain update tokens", map[string]interface{}{"id": objectId})
	currentState, _, getErr := r.client.NgfirewallsAPI.GetNgfirewallsByID(ctx, objectId).Region(regionVal).Execute()
	if getErr != nil {
		resp.Diagnostics.AddError("Error fetching current state for update", getErr.Error())
		return
	}
	var updateToken, deploymentUpdateToken string
	if currentState != nil && currentState.Response != nil && currentState.Response.Firewall != nil {
		if currentState.Response.Firewall.UpdateToken != nil {
			updateToken = *currentState.Response.Firewall.UpdateToken
		}
		if currentState.Response.Firewall.DeploymentUpdateToken != nil {
			deploymentUpdateToken = *currentState.Response.Firewall.DeploymentUpdateToken
		}
	}
	if updateToken == "" || deploymentUpdateToken == "" {
		resp.Diagnostics.AddError("Missing update tokens", "UpdateToken or DeploymentUpdateToken not found in current firewall state")
		return
	}
	tflog.Debug(ctx, "Obtained update tokens", map[string]interface{}{"UpdateToken": updateToken, "DeploymentUpdateToken": deploymentUpdateToken})

	// Build the update request with updatable fields plus required tokens.
	updateBody := cloud_ngfw.NgfirewallUpdateRequest{}
	updateBody.UpdateToken = updateToken
	updateBody.DeploymentUpdateToken = deploymentUpdateToken
	// Generated from the NgfirewallUpdateRequest model: every NgfirewallUpdateRequest field is taken from the planned values,
	// so a field NgfirewallUpdateRequest does not define (e.g. Mode, Source) can never reach the PATCH body.
	updateSrc, updateSrcDiags := unpackNgfirewallCreateRequestToSdk(ctx, planObject)
	resp.Diagnostics.Append(updateSrcDiags...)
	if resp.Diagnostics.HasError() {
		return
	}
	updateBody.AllowListAccounts = updateSrc.AllowListAccounts
	updateBody.ChangeProtection = updateSrc.ChangeProtection
	updateBody.CustomerZoneIdList = updateSrc.CustomerZoneIdList
	updateBody.Description = updateSrc.Description
	updateBody.EgressNAT = updateSrc.EgressNAT
	updateBody.Endpoints = updateSrc.Endpoints
	updateBody.GwlbTcpIdleTimeout = updateSrc.GwlbTcpIdleTimeout
	updateBody.LinkId = updateSrc.LinkId
	updateBody.PrivateAccess = updateSrc.PrivateAccess
	updateBody.ScmLabels = updateSrc.ScmLabels
	updateBody.Tags = updateSrc.Tags
	updateBody.Tier = updateSrc.Tier
	updateBody.UserID = updateSrc.UserID

	updateReq := r.client.NgfirewallsAPI.UpdateNgfirewallsByID(ctx, objectId).NgfirewallUpdateRequest(updateBody)
	if !plan.Region.IsNull() {
		updateReq = updateReq.Region(plan.Region.ValueString())
	}

	_, httpErr, err := updateReq.Execute()
	if err != nil {
		if httpErr != nil && httpErr.StatusCode == http.StatusNotFound {
			tflog.Debug(ctx, "Got no ngfirewalls on update SCM API. Remove from state to let terraform create", map[string]interface{}{"id": objectId})
			resp.State.RemoveResource(ctx)
		} else {
			tflog.Debug(ctx, "Got an exception on update SCM API. ", map[string]interface{}{"id": objectId})
			resp.Diagnostics.AddError("Error updating ngfirewalls", err.Error())
			detailedMessage := utils.PrintScmError(err)
			resp.Diagnostics.AddError("SCM Resource Update Failed: API Request Failed", detailedMessage)
		}
		return
	}

	// CNGFW update is asynchronous — poll until the firewall reaches UPDATE_COMPLETE.
	updateTimeout := cngfwCreateDefaultTimeout
	tflog.Debug(ctx, "Waiting for CNGFW firewall to reach UPDATE_COMPLETE", map[string]interface{}{
		"firewall_id": objectId,
		"region":      tokens[0],
		"timeout":     updateTimeout.String(),
	})
	if waitErr := waitForCngfwStatus(ctx, r.client, r.setupClient, objectId, tokens[0], updateTimeout, "UPDATE_COMPLETE"); waitErr != nil {
		resp.Diagnostics.AddError("CNGFW Update: update failed", waitErr.Error())
		return
	}
	tflog.Debug(ctx, "CNGFW firewall reached UPDATE_COMPLETE", map[string]interface{}{"firewall_id": objectId})

	// CNGFW: plan already carries the correct values; no pack from SDK needed.
	_ = req.Plan.GetAttribute(ctx, path.Root("region"), &plan.Region)

	plan.Tfid = state.Tfid
	_ = req.Plan.GetAttribute(ctx, path.Root("id"), &plan.Id)
	_ = req.Plan.GetAttribute(ctx, path.Root("region"), &plan.Region)

	// Re-read the firewall so every Computed attribute has a known value before
	// state is written. Same contract as Create: Terraform rejects an apply that
	// leaves a Computed attribute unknown, and the PATCH response does not carry
	// the server-assigned values (Tier, LinkId, ChangeProtection, EgressNAT,
	// Endpoints, GwlbTcpIdleTimeout).
	getRes, _, postReadErr := r.client.NgfirewallsAPI.GetNgfirewallsByID(ctx, objectId).Region(regionVal).Execute()
	if postReadErr != nil {
		resp.Diagnostics.AddError(
			"CNGFW Update: post-update read failed",
			"The firewall was updated but reading it back failed, so Computed attributes cannot be resolved. "+
				"Re-run terraform plan to reconcile state. Error: "+postReadErr.Error(),
		)
		return
	}

	// Anything the GET does not return falls back to the planned values.
	data := plan
	savestate := plan

	// Populate state from the GET response. Generated from the NgfirewallCreateRequest and NgfirewallFirewall models:
	// fields the GET does not return keep their planned/prior value (null on fresh import).
	if getRes != nil && getRes.Response != nil && getRes.Response.Firewall != nil {
		fwObj, fwDiags := packNgfirewallFirewallFromSdk(ctx, *getRes.Response.Firewall)
		resp.Diagnostics.Append(fwDiags...)
		var fw models.NgfirewallFirewall
		resp.Diagnostics.Append(fwObj.As(ctx, &fw, basetypes.ObjectAsOptions{})...)
		if resp.Diagnostics.HasError() {
			return
		}
		data.AllowListAccounts = fw.AllowListAccounts
		data.ChangeProtection = fw.ChangeProtection
		data.CustomerZoneIdList = fw.CustomerZoneIdList
		data.Description = fw.Description
		data.EgressNAT = fw.EgressNAT
		data.Endpoints = fw.Endpoints
		{
			var nested models.NgfirewallGwlbConfig
			if !fw.GWLB.IsNull() && !fw.GWLB.IsUnknown() {
				resp.Diagnostics.Append(fw.GWLB.As(ctx, &nested, basetypes.ObjectAsOptions{})...)
			}
			data.GwlbTcpIdleTimeout = nested.TCPIdleTimeout
		}
		data.LinkId = fw.LinkId
		data.PrivateAccess = fw.PrivateAccess
		data.Tags = fw.Tags
		data.Tier = fw.Tier
		data.UserID = fw.UserID
		if !fw.FirewallId.IsNull() {
			data.Id = fw.FirewallId
		}
	} else {
		// GET succeeded but contained no Firewall object - unexpected; restore from prior state.
		data = savestate
	}

	// Fields not returned by GET, and identity fields established above.
	data.Mode = savestate.Mode
	data.ScmLabels = savestate.ScmLabels
	data.Region = savestate.Region
	data.Tfid = savestate.Tfid

	plan = data

	tflog.Debug(ctx, "Updated ngfirewalls", map[string]interface{}{"tfid": plan.Tfid.ValueString()})
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *NgfirewallResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data models.NgfirewallCreateRequest
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tokens := strings.Split(data.Tfid.ValueString(), ":")
	if len(tokens) != 2 {
		resp.Diagnostics.AddError("Error parsing TFID", fmt.Sprintf("Expected a TFID with 2 parts separated by ':', but got %d parts for TFID %s", len(tokens), data.Tfid.ValueString()))
		return
	}
	objectId := tokens[1]

	tflog.Debug(ctx, "Deleting ngfirewalls", map[string]interface{}{"id": objectId})
	_, err := r.client.NgfirewallsAPI.DeleteCloudNGFirewall(ctx, objectId, cloud_ngfw.CngfwSourceTerraformProvider, data.Region.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error deleting ngfirewalls", err.Error())
		detailedMessage := utils.PrintScmError(err)
		resp.Diagnostics.AddError("SCM Resource Deletion Failed: API Request Failed", detailedMessage)
		return
	}

	// CNGFW deletion is asynchronous — poll until the firewall is gone.
	deleteTimeout := cngfwDeleteDefaultTimeout
	if waitErr := waitForCngfwStatus(ctx, r.client, r.setupClient, objectId, data.Region.ValueString(), deleteTimeout, "DELETE_COMPLETE"); waitErr != nil {
		resp.Diagnostics.AddError("CNGFW Delete: deletion failed", waitErr.Error())
		return
	}
	tflog.Debug(ctx, "CNGFW firewall successfully deleted", map[string]interface{}{"firewall_id": objectId})
}

func (r *NgfirewallResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Set tfid so the Read function can parse it.
	resource.ImportStatePassthroughID(ctx, path.Root("tfid"), req, resp)

	// Parse "region:firewallId" and pre-populate the Required "region" attribute so
	// that Terraform can validate the plan and -generate-config-out can emit a complete
	// resource block without the user having to supply region manually.
	tokens := strings.Split(req.ID, ":")
	if len(tokens) == 2 && tokens[0] != "" {
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("region"), tokens[0])...)
	}

}

/*
 * waitForCngfwStatus polls GetNgfirewallsByID every 30 seconds until the firewall
 * reaches one of the caller-supplied terminal OK statuses, a terminal failure status
 * (CREATE_FAIL, UPDATE_FAIL, DELETE_FAIL), context cancellation, or timeout.
 *
 * For delete polling, a 404 response is treated as success: the firewall no longer
 * exists in the API, which is the desired end-state after deletion.
 *
 * Pass one or more acceptable terminal states; the first match returns nil.
 * Example: waitForCngfwStatus(ctx, client, setupClient, id, region, timeout, "CREATE_COMPLETE", "UPDATE_COMPLETE")
 *          waitForCngfwStatus(ctx, client, setupClient, id, region, timeout, "DELETE_COMPLETE")
 */
func waitForCngfwStatus(ctx context.Context, client *cloud_ngfw.APIClient, setupClient *setup.Client, firewallID string, region string, timeout time.Duration, terminalOK ...string) error {
	okSet := make(map[string]struct{}, len(terminalOK))
	for _, s := range terminalOK {
		okSet[s] = struct{}{}
	}
	_, isDeleteWait := okSet["DELETE_COMPLETE"]

	deadline := time.Now().Add(timeout)
	pollInterval := 30 * time.Second
	for time.Now().Before(deadline) {
		refreshJwtIfNeeded(ctx, setupClient)

		res, httpResp, err := client.NgfirewallsAPI.GetNgfirewallsByID(ctx, firewallID).Region(region).Execute()
		if err != nil {
			// The CNGFW backend returns 400 (unknown ID) or 404 (not found) once a
			// firewall is gone. Only treat those codes as "deletion confirmed"; other
			// errors (network failures, 5xx) should retry so we don't falsely report
			// success on a transient fault.
			if isDeleteWait && httpResp != nil && (httpResp.StatusCode == http.StatusNotFound || httpResp.StatusCode == http.StatusBadRequest) {
				tflog.Debug(ctx, "waitForCngfwStatus: firewall not found — deletion confirmed", map[string]interface{}{"status_code": httpResp.StatusCode})
				return nil
			}
			tflog.Warn(ctx, "waitForCngfwStatus: poll error, will retry", map[string]interface{}{"error": err.Error()})
		} else if res != nil && res.Response != nil {
			if status, ok := res.Response.GetStatusOk(); ok && status != nil {
				current := status.GetFirewallStatus()
				tflog.Debug(ctx, "waitForCngfwStatus: polled status", map[string]interface{}{"status": current})
				if _, done := okSet[current]; done {
					// The caller reads the firewall right after a successful wait; the token may have
					// been reset during the last poll, so make sure it is usable before returning.
					refreshJwtIfNeeded(ctx, setupClient)
					return nil
				}
				switch current {
				case "CREATE_FAIL", "UPDATE_FAIL", "DELETE_FAIL":
					return fmt.Errorf("CNGFW firewall %s reached terminal failure status: %s", firewallID, current)
				}
			}
		}

		select {
		case <-ctx.Done():
			return fmt.Errorf("context cancelled while waiting for CNGFW firewall %s", firewallID)
		case <-time.After(pollInterval):
		}
	}
	return fmt.Errorf("timed out after %v waiting for CNGFW firewall %s to reach %v", timeout, firewallID, terminalOK)
}

// refreshJwtIfNeeded refreshes the SDK token when it is missing or past its expiry.
// CNGFW operations run long enough (a create takes roughly 10-15 minutes) to reach the last
// minutes of a token's life, where the provider's token transport can reset the token and its
// expiry; call this before API calls that follow a long wait.
func refreshJwtIfNeeded(ctx context.Context, setupClient *setup.Client) {
	if setupClient.Jwt != "" && time.Now().Before(setupClient.JwtExpiresAt) {
		return
	}
	tflog.Debug(ctx, "refreshJwtIfNeeded: token missing or expired, refreshing")
	if refreshErr := setupClient.RefreshJwt(ctx); refreshErr != nil {
		tflog.Error(ctx, "refreshJwtIfNeeded: token refresh failed", map[string]interface{}{"error": refreshErr.Error()})
	}
}
