// Source of truth for the scm_security_rule_order resource. Embedded by the
// generator (cmd/generate/terraform/convertor_generator.go) and emitted
// verbatim to
// terraform-provider-scm/internal/provider/security_services/resource_security_rule_order.go.
// The .go.tmpl extension keeps the file out of the generator's own compile
// unit while still letting //go:embed read it as raw bytes.
//
// The move planner (LIS algorithm) lives in rule_order_lis.go — this file is
// only the Terraform-plugin-framework wrapper that reads config, lists live
// SCM state, calls PlanMoves(), and dispatches the resulting :move calls via
// the SCM SDK.
package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"

	"github.com/paloaltonetworks/scm-go/generated/security_services"
	"github.com/paloaltonetworks/terraform-provider-scm/internal/utils"
)

var (
	_ resource.Resource                = &SecurityRuleOrderResource{}
	_ resource.ResourceWithConfigure   = &SecurityRuleOrderResource{}
	_ resource.ResourceWithImportState = &SecurityRuleOrderResource{}
)

func NewSecurityRuleOrderResource() resource.Resource {
	return &SecurityRuleOrderResource{}
}

type SecurityRuleOrderResource struct {
	client *security_services.APIClient
}

// SecurityRuleOrderModel is the Terraform state model.
type SecurityRuleOrderModel struct {
	Tfid            types.String `tfsdk:"tfid"`
	Folder          types.String `tfsdk:"folder"`
	Snippet         types.String `tfsdk:"snippet"`
	Device          types.String `tfsdk:"device"`
	Rulebase        types.String `tfsdk:"rulebase"`
	Destination     types.String `tfsdk:"destination"`
	DestinationRule types.String `tfsdk:"destination_rule"`
	Order           types.String `tfsdk:"order"`
	OrderedRules    types.List   `tfsdk:"ordered_rules"`
}

func (r *SecurityRuleOrderResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_rule_order"
}

func (r *SecurityRuleOrderResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	containerOneOf := []validator.String{
		stringvalidator.ExactlyOneOf(
			path.MatchRoot("device"),
			path.MatchRoot("folder"),
			path.MatchRoot("snippet"),
		),
	}
	resp.Schema = schema.Schema{
		MarkdownDescription: "Enforces the order of a named set of security rules within a single (container, rulebase). " +
			"Rules present in SCM but not listed in `ordered_rules` are left in place under the default `consecutive` mode; " +
			"they may be moved under `strictly_consecutive` mode to keep the block uninterrupted. " +
			"Uses the `/security-rules/{id}:move` API. Manage the rules themselves with `scm_security_rule` resources.",
		Attributes: map[string]schema.Attribute{
			"tfid": schema.StringAttribute{
				Description: "The Terraform ID: `<folder>:<snippet>:<device>:<rulebase>` (unused container slots empty), the same shape as `scm_security_rule`.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"folder": schema.StringAttribute{
				MarkdownDescription: "The folder the rules belong to. Exactly one of `folder`, `snippet`, or `device` must be set.",
				Optional:            true,
				Validators:          append([]validator.String{utils.FolderValidator()}, containerOneOf...),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"snippet": schema.StringAttribute{
				MarkdownDescription: "The snippet the rules belong to. Exactly one of `folder`, `snippet`, or `device` must be set.",
				Optional:            true,
				Validators:          containerOneOf,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"device": schema.StringAttribute{
				MarkdownDescription: "The device the rules belong to. Exactly one of `folder`, `snippet`, or `device` must be set.",
				Optional:            true,
				Validators:          containerOneOf,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"rulebase": schema.StringAttribute{
				MarkdownDescription: "Rulebase to reorder. One of `pre` or `post`.",
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("pre", "post"),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"destination": schema.StringAttribute{
				MarkdownDescription: "Where the block of `ordered_rules` should anchor. One of `null`, `top`, `bottom`, `before`, `after`. " +
					"`null` (default) preserves relative order without pinning an absolute position. " +
					"`top` / `bottom` pin the first / last rule to the top / bottom of the rulebase. " +
					"`before` / `after` pin the first / last rule immediately adjacent to `destination_rule`.",
				Optional: true,
				Computed: true,
				Validators: []validator.String{
					stringvalidator.OneOf("null", "top", "bottom", "before", "after"),
				},
			},
			"destination_rule": schema.StringAttribute{
				MarkdownDescription: "Name of the anchor rule; required when `destination` is `before` or `after`. Must be a rule that exists in SCM but is NOT listed in `ordered_rules`.",
				Optional:            true,
			},
			"order": schema.StringAttribute{
				MarkdownDescription: "Ordering mode. `consecutive` (default) preserves the relative order of `ordered_rules` while allowing manual or third-party rules to be interleaved. " +
					"`strictly_consecutive` enforces a contiguous block with no interleaving; the algorithm may issue `:move` calls on unmanaged rules to keep total move count down.",
				Optional: true,
				Computed: true,
				Validators: []validator.String{
					stringvalidator.OneOf("consecutive", "strictly_consecutive"),
				},
			},
			"ordered_rules": schema.ListAttribute{
				MarkdownDescription: "Ordered list of security rule names, top-to-bottom. Rules will be arranged in this left-to-right order.",
				Required:            true,
				ElementType:         types.StringType,
			},
		},
	}
}

func (r *SecurityRuleOrderResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	clients, ok := req.ProviderData.(map[string]interface{})
	if !ok {
		resp.Diagnostics.AddError("Unexpected Resource Configure Type", fmt.Sprintf("Expected map[string]interface{}, got: %T.", req.ProviderData))
		return
	}
	client, ok := clients["security_services"].(*security_services.APIClient)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Client Type", "Expected *security_services.APIClient for 'security_services' client.")
		return
	}
	r.client = client
}

func (r *SecurityRuleOrderResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SecurityRuleOrderModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	applyDefaults(&data)
	r.applyOrder(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	data.Tfid = types.StringValue(buildRuleOrderTfid(data))
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SecurityRuleOrderResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SecurityRuleOrderModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	applyDefaults(&data)

	desired, diags := extractStringList(ctx, data.OrderedRules)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	currentOrder, _, listDiags := r.listAllRuleNames(ctx, &data)
	resp.Diagnostics.Append(listDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Drift detection differs by mode.
	//
	// Consecutive mode cares only about the relative order of target rules,
	// so the observed value is the projection of currentOrder onto desired.
	// If the projection differs from desired (order drift), Terraform shows
	// the diff and triggers Update.
	//
	// Strict Consecutive mode also cares about the ABSENCE of interleavers
	// inside the block window AND the block's absolute position (for
	// anchored destinations). The projection would miss both of those:
	// projecting [A, u, B, C] onto [A, B, C] still yields [A, B, C], hiding
	// the interleaver `u`. So for strict mode we return the raw slice of
	// currentOrder covering the whole enforced block window — target rules
	// plus any interleavers/prefix/suffix that must be there per the
	// destination. If any non-target rule shows up in that slice (or the
	// block is positioned wrong), the observed list differs from desired
	// and Terraform triggers Update.
	destination := data.Destination.ValueString()
	anchor := ""
	if !data.DestinationRule.IsNull() && !data.DestinationRule.IsUnknown() {
		anchor = data.DestinationRule.ValueString()
	}
	mode := data.Order.ValueString()

	observed := observeCurrentBlock(currentOrder, desired, destination, anchor, mode)

	tflog.Debug(ctx, "rule_order: Read observed layout for drift detection", map[string]interface{}{
		"mode":        mode,
		"destination": destination,
		"desired":     desired,
		"observed":    observed,
		"drift":       !stringSlicesEqual(observed, desired),
	})

	newList, diags := stringListValue(ctx, observed)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	data.OrderedRules = newList
	data.Tfid = types.StringValue(buildRuleOrderTfid(data))
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SecurityRuleOrderResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data SecurityRuleOrderModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	applyDefaults(&data)
	r.applyOrder(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	data.Tfid = types.StringValue(buildRuleOrderTfid(data))
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SecurityRuleOrderResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// No-op: the resource only asserts an ordering. Deleting it stops
	// enforcing the order but does not delete any rules.
}

// ImportState accepts an ID shaped like the scm_security_rule import ID
// (`folder:snippet:device:<id>`), where <id> is the rulebase optionally followed
// by the managed rule names:
//
//	<folder>:::<pre|post>:<rule1>,<rule2>,...
//	:<snippet>::<pre|post>:<rule1>,<rule2>,...
//	::<device>:<pre|post>:<rule1>,<rule2>,...
//
// The rule names identify which rules this resource manages (the container may
// also hold unmanaged rules). They seed `ordered_rules`; the Read that follows
// import replaces them with the live relative order of those rules, so the next
// plan shows exactly how the live rulebase differs from the configuration.
// The form without rule names is still accepted and imports an empty list.
func (r *SecurityRuleOrderResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	kind, value, rulebase, rules, err := ParseRuleOrderImportID(req.ID)
	if err != nil {
		resp.Diagnostics.AddError("Invalid Import ID", err.Error())
		return
	}
	// Same deprecated-folder check as the schema validator and scm_security_rule import.
	if kind == "folder" {
		var vr validator.StringResponse
		utils.FolderValidator().ValidateString(ctx, validator.StringRequest{
			Path:        path.Root("folder"),
			ConfigValue: types.StringValue(value),
		}, &vr)
		if vr.Diagnostics.HasError() {
			resp.Diagnostics.Append(vr.Diagnostics...)
			return
		}
	}

	data := SecurityRuleOrderModel{
		Folder:          types.StringNull(),
		Snippet:         types.StringNull(),
		Device:          types.StringNull(),
		Rulebase:        types.StringValue(rulebase),
		Destination:     types.StringValue("null"),
		DestinationRule: types.StringNull(),
		Order:           types.StringValue(OrderModeConsecutive),
	}
	switch kind {
	case "folder":
		data.Folder = types.StringValue(value)
	case "snippet":
		data.Snippet = types.StringValue(value)
	case "device":
		data.Device = types.StringValue(value)
	}

	// Every named rule must exist in (and be owned by) the container.
	if len(rules) > 0 {
		current, _, listDiags := r.listAllRuleNames(ctx, &data)
		resp.Diagnostics.Append(listDiags...)
		if resp.Diagnostics.HasError() {
			return
		}
		for _, name := range rules {
			if !containsString(current, name) {
				resp.Diagnostics.AddError(
					"Rule not found for import",
					fmt.Sprintf("Rule %q was not found in the %s %q (%s rulebase). Only rules owned by that container can be imported.", name, kind, value, rulebase),
				)
			}
		}
		if resp.Diagnostics.HasError() {
			return
		}
	}

	list, diags := stringListValue(ctx, rules)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	data.OrderedRules = list
	data.Tfid = types.StringValue(buildRuleOrderTfid(data))
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// applyDefaults fills in missing optional-with-default fields.
func applyDefaults(data *SecurityRuleOrderModel) {
	if data.Destination.IsNull() || data.Destination.IsUnknown() || data.Destination.ValueString() == "" {
		data.Destination = types.StringValue("null")
	}
	if data.Order.IsNull() || data.Order.IsUnknown() || data.Order.ValueString() == "" {
		data.Order = types.StringValue(OrderModeConsecutive)
	}
}

// applyOrder is the shared Create/Update path: extract config, read current
// SCM state, delegate to the LIS planner, dispatch :move API calls, verify.
//
// Debug troubleshooting: run Terraform with
//
//	TF_LOG=DEBUG terraform apply
//	# or, to see only the SCM provider's logs:
//	TF_LOG_PROVIDER=DEBUG terraform apply
//
// The provider emits structured tflog entries with the "rule_order" subsystem
// tag on every phase (config, live state, planner decision, per-move dispatch,
// verification). Grep for `subsystem=rule_order` in the log to filter noise
// from unrelated resources.
func (r *SecurityRuleOrderResource) applyOrder(ctx context.Context, data *SecurityRuleOrderModel, diags *diag.Diagnostics) {
	// Attach the rule_order subsystem tag to this apply's context so every
	// tflog call inherits it. Users can then filter with
	//   TF_LOG=DEBUG terraform apply 2>&1 | grep 'subsystem=rule_order'
	ctx = tflog.SetField(ctx, "subsystem", "rule_order")
	ctx = tflog.SetField(ctx, "rulebase", data.Rulebase.ValueString())
	ctx = tflog.SetField(ctx, "container", containerLabel(data))

	desired, d := extractStringList(ctx, data.OrderedRules)
	diags.Append(d...)
	if diags.HasError() {
		return
	}

	destination := data.Destination.ValueString()
	anchor := ""
	if !data.DestinationRule.IsNull() && !data.DestinationRule.IsUnknown() {
		anchor = data.DestinationRule.ValueString()
	}
	mode := data.Order.ValueString()

	tflog.Debug(ctx, "rule_order: applying", map[string]interface{}{
		"destination":      destination,
		"destination_rule": anchor,
		"order_mode":       mode,
		"target_count":     len(desired),
		"target_rules":     desired,
	})

	// A single target rule still needs planning: destination may pin it.
	if len(desired) == 0 {
		tflog.Debug(ctx, "rule_order: no target rules; nothing to order")
		return
	}

	currentOrder, nameToID, listDiags := r.listAllRuleNames(ctx, data)
	diags.Append(listDiags...)
	if diags.HasError() {
		return
	}
	tflog.Debug(ctx, "rule_order: fetched current SCM state", map[string]interface{}{
		"rule_count":    len(currentOrder),
		"current_order": currentOrder,
	})

	moves, err := PlanMoves(desired, currentOrder, destination, anchor, mode)
	if err != nil {
		tflog.Error(ctx, "rule_order: PlanMoves rejected configuration", map[string]interface{}{
			"error": err.Error(),
		})
		diags.AddError("Invalid rule ordering configuration", err.Error())
		return
	}
	tflog.Debug(ctx, "rule_order: planner produced move plan", map[string]interface{}{
		"move_count":  len(moves),
		"plan":        renderMovePlan(moves),
		"description": describePlan(desired, currentOrder, moves),
	})
	if len(moves) == 0 {
		tflog.Debug(ctx, "rule_order: already satisfies the plan; no API calls needed")
		return
	}

	rulebase := data.Rulebase.ValueString()
	for i, m := range moves {
		id, ok := nameToID[m.Rule]
		if !ok {
			tflog.Error(ctx, "rule_order: planner referenced missing rule (should not happen)", map[string]interface{}{
				"rule": m.Rule,
			})
			diags.AddError(
				fmt.Sprintf("Rule %q not found in SCM", m.Rule),
				"The planner emitted a move for a rule not present in the live rulebase. This should not happen; please file a bug.",
			)
			return
		}
		payload := security_services.NewRuleBasedMove(m.Position, rulebase)
		if m.DestinationRule != "" {
			destID, ok := nameToID[m.DestinationRule]
			if !ok {
				tflog.Error(ctx, "rule_order: planner referenced missing anchor rule", map[string]interface{}{
					"anchor": m.DestinationRule,
				})
				diags.AddError(
					fmt.Sprintf("Destination rule %q not found in SCM", m.DestinationRule),
					"The planner referenced an anchor rule that is missing from the live rulebase.",
				)
				return
			}
			payload.SetDestinationRule(destID)
		}
		tflog.Debug(ctx, "rule_order: issuing :move API call", map[string]interface{}{
			"step":             fmt.Sprintf("%d/%d", i+1, len(moves)),
			"rule":             m.Rule,
			"rule_id":          id,
			"position":         m.Position,
			"destination_rule": m.DestinationRule,
			"api_summary":      renderMove(m),
		})
		httpResp, err := r.client.SecurityRulesAPI.
			MoveSecurityRulesByID(ctx, id).
			RuleBasedMove(*payload).
			Execute()
		if err != nil {
			tflog.Error(ctx, "rule_order: :move API call failed", map[string]interface{}{
				"step":  fmt.Sprintf("%d/%d", i+1, len(moves)),
				"rule":  m.Rule,
				"error": err.Error(),
			})
			diags.AddError(
				fmt.Sprintf("Failed to move rule %q", m.Rule),
				utils.PrintScmError(err),
			)
			return
		}
		if httpResp != nil && httpResp.Body != nil {
			_ = httpResp.Body.Close()
		}
	}
	tflog.Debug(ctx, "rule_order: all :move calls succeeded", map[string]interface{}{
		"moves_issued": len(moves),
	})

	// Verify: re-list and confirm the post-apply state matches expectations.
	// The check uses observeCurrentBlock — the same function Read uses for
	// drift detection — so an apply that succeeds here is guaranteed to
	// produce a no-drift Read on the very next plan.
	//
	//   Consecutive mode          → observed = projection; must equal desired
	//   Strict Consecutive mode   → observed = enforced block window slice;
	//                               must equal desired (no interleavers, no
	//                               position drift)
	verifiedOrder, _, verifyDiags := r.listAllRuleNames(ctx, data)
	diags.Append(verifyDiags...)
	if diags.HasError() {
		return
	}
	observed := observeCurrentBlock(verifiedOrder, desired, destination, anchor, mode)
	match := stringSlicesEqual(observed, desired)
	tflog.Debug(ctx, "rule_order: post-apply verification", map[string]interface{}{
		"mode":     mode,
		"observed": observed,
		"desired":  desired,
		"match":    match,
	})
	if !match {
		tflog.Error(ctx, "rule_order: verification MISMATCH — post-apply state does not match desired", map[string]interface{}{
			"mode":     mode,
			"observed": observed,
			"desired":  desired,
		})
		diags.AddError(
			"Rule order verification failed",
			fmt.Sprintf(
				"After all :move calls, the post-apply state in %s mode did not match the desired order.\n"+
					"  desired  : %v\n"+
					"  observed : %v\n"+
					"This usually indicates a concurrent writer reordering rules, or a bug in the planner. "+
					"Re-run `terraform apply` — the resource is idempotent and will reconcile.",
				mode, desired, observed,
			),
		)
	}
}

// -----------------------------------------------------------------------------
// Debug output helpers.
// -----------------------------------------------------------------------------

// renderMove returns a human-readable summary of a single Move, matching the
// SCM API call it will trigger.
func renderMove(m Move) string {
	if m.DestinationRule != "" {
		return fmt.Sprintf(":move %s %s %s", m.Rule, m.Position, m.DestinationRule)
	}
	return fmt.Sprintf(":move %s %s", m.Rule, m.Position)
}

// renderMovePlan returns a numbered list of moves suitable for tflog output.
func renderMovePlan(moves []Move) []string {
	if len(moves) == 0 {
		return nil
	}
	out := make([]string, 0, len(moves))
	for i, m := range moves {
		out = append(out, fmt.Sprintf("[%d/%d] %s", i+1, len(moves), renderMove(m)))
	}
	return out
}

// describePlan returns a short natural-language summary of what will change.
// Kept intentionally terse — it complements, not replaces, the structured
// tflog fields.
func describePlan(desired, current []string, moves []Move) string {
	if len(moves) == 0 {
		return "no-op: rules already in the correct relative order"
	}
	movedSet := make(map[string]bool, len(moves))
	for _, m := range moves {
		movedSet[m.Rule] = true
	}
	var stayed []string
	for _, r := range desired {
		if !movedSet[r] {
			stayed = append(stayed, r)
		}
	}
	return fmt.Sprintf("%d :move call(s); %d target rule(s) stay put (%v)", len(moves), len(stayed), stayed)
}

// containerLabel returns a short "kind:value" tag for logging context.
func containerLabel(data *SecurityRuleOrderModel) string {
	switch {
	case !data.Folder.IsNull() && data.Folder.ValueString() != "":
		return "folder:" + data.Folder.ValueString()
	case !data.Snippet.IsNull() && data.Snippet.ValueString() != "":
		return "snippet:" + data.Snippet.ValueString()
	case !data.Device.IsNull() && data.Device.ValueString() != "":
		return "device:" + data.Device.ValueString()
	}
	return "unknown"
}

// listAllRuleNames paginates ListRules for the given container+rulebase and
// returns (movableNamesInOrder, nameToID, diags).
//
// The SCM API returns rules from the target container AND rules inherited
// from ancestor folders/snippets/devices. Inherited rules appear in the
// response so the caller can see the full evaluation order, but they cannot
// be `:move`d — attempting to move a rule that isn't owned by the target
// container yields a 400/403 from SCM.
//
// This function filters the response down to just the rules whose owner
// (Folder/Snippet/Device field on the SDK model) matches the resource's own
// container. The dropped inherited rules are logged at DEBUG so users can
// see them for troubleshooting.
func (r *SecurityRuleOrderResource) listAllRuleNames(ctx context.Context, data *SecurityRuleOrderModel) ([]string, map[string]string, diag.Diagnostics) {
	var diags diag.Diagnostics
	rulebase := data.Rulebase.ValueString()

	// Determine which owner field must match to consider a rule "movable".
	ownerKind, ownerValue := containerKindAndValue(data)

	const pageSize = int32(200)
	var offset int32
	var namesInOrder []string
	nameToID := map[string]string{}
	var inheritedDropped []string

	for {
		listReq := r.client.SecurityRulesAPI.ListRules(ctx).
			Position(rulebase).
			Limit(pageSize).
			Offset(offset)
		if !data.Folder.IsNull() && data.Folder.ValueString() != "" {
			listReq = listReq.Folder(data.Folder.ValueString())
		}
		if !data.Snippet.IsNull() && data.Snippet.ValueString() != "" {
			listReq = listReq.Snippet(data.Snippet.ValueString())
		}
		if !data.Device.IsNull() && data.Device.ValueString() != "" {
			listReq = listReq.Device(data.Device.ValueString())
		}

		listResponse, httpResp, err := listReq.Execute()
		if err != nil {
			diags.AddError(
				"Failed to list security rules",
				utils.PrintScmError(err),
			)
			return nil, nil, diags
		}
		if httpResp != nil && httpResp.Body != nil {
			_ = httpResp.Body.Close()
		}
		if listResponse == nil {
			break
		}
		page := listResponse.GetData()
		if len(page) == 0 {
			break
		}
		for _, rule := range page {
			name := ""
			if rule.Name != nil {
				name = *rule.Name
			}
			id := ""
			if rule.Id != nil {
				id = *rule.Id
			}
			if name == "" || id == "" {
				continue
			}
			// Filter: only rules owned by the target container are movable.
			if !ruleOwnedBy(rule, ownerKind, ownerValue) {
				inheritedDropped = append(inheritedDropped, name)
				continue
			}
			namesInOrder = append(namesInOrder, name)
			nameToID[name] = id
		}
		if int32(len(page)) < pageSize {
			break
		}
		offset += int32(len(page))
	}

	if len(inheritedDropped) > 0 {
		tflog.Debug(ctx, "rule_order: filtered out inherited rules from ancestor containers", map[string]interface{}{
			"dropped_count": len(inheritedDropped),
			"dropped_rules": inheritedDropped,
			"kept_count":    len(namesInOrder),
			"note":          "these rules exist in the rulebase view but are owned by an ancestor container and cannot be :move'd from here",
		})
	}
	return namesInOrder, nameToID, diags
}

// containerKindAndValue returns the ("folder"|"snippet"|"device", value) for
// the container the resource manages.
func containerKindAndValue(data *SecurityRuleOrderModel) (string, string) {
	switch {
	case !data.Folder.IsNull() && data.Folder.ValueString() != "":
		return "folder", data.Folder.ValueString()
	case !data.Snippet.IsNull() && data.Snippet.ValueString() != "":
		return "snippet", data.Snippet.ValueString()
	case !data.Device.IsNull() && data.Device.ValueString() != "":
		return "device", data.Device.ValueString()
	}
	return "", ""
}

// ruleOwnedBy returns true iff the rule's owner field matches the target
// container. Rules inherited from ancestor containers have a different
// owner value and are filtered out.
func ruleOwnedBy(rule security_services.SecurityRules, ownerKind, ownerValue string) bool {
	switch ownerKind {
	case "folder":
		return rule.Folder != nil && *rule.Folder == ownerValue
	case "snippet":
		return rule.Snippet != nil && *rule.Snippet == ownerValue
	case "device":
		return rule.Device != nil && *rule.Device == ownerValue
	}
	// If we can't determine ownership, keep the rule to preserve prior behavior.
	return true
}

// buildRuleOrderTfid returns the deterministic tfid for state/import.
func buildRuleOrderTfid(data SecurityRuleOrderModel) string {
	kind, value := "folder", ""
	switch {
	case !data.Folder.IsNull() && data.Folder.ValueString() != "":
		kind, value = "folder", data.Folder.ValueString()
	case !data.Snippet.IsNull() && data.Snippet.ValueString() != "":
		kind, value = "snippet", data.Snippet.ValueString()
	case !data.Device.IsNull() && data.Device.ValueString() != "":
		kind, value = "device", data.Device.ValueString()
	}
	// Same slot layout as the scm_security_rule tfid: folder:snippet:device:<id>.
	var folder, snippet, device string
	switch kind {
	case "folder":
		folder = value
	case "snippet":
		snippet = value
	case "device":
		device = value
	}
	return fmt.Sprintf("%s:%s:%s:%s", folder, snippet, device, data.Rulebase.ValueString())
}

func extractStringList(ctx context.Context, l types.List) ([]string, diag.Diagnostics) {
	var diags diag.Diagnostics
	if l.IsNull() || l.IsUnknown() {
		return nil, diags
	}
	out := make([]string, 0, len(l.Elements()))
	diags.Append(l.ElementsAs(ctx, &out, false)...)
	return out, diags
}

func stringListValue(ctx context.Context, in []string) (types.List, diag.Diagnostics) {
	elems := make([]attr.Value, 0, len(in))
	for _, s := range in {
		elems = append(elems, types.StringValue(s))
	}
	return types.ListValue(types.StringType, elems)
}
