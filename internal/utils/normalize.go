package utils

/*
THIS FILE IS EMBEDDED INTO THE GENERATOR AND COPIED TO THE OUTPUT

This entire file is embedded into the terraform generator and will be written
directly to terraform-provider-scm/internal/utils/normalize.go with only
the package name changed from "terraform" to "utils".

The functions in this file are NOT used by the generator itself but are intended
for the generated Terraform provider.
*/

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// NormalizeNullLists recursively compares a planned (or prior-state)
// types.Object against the API response types.Object and reconciles the two
// representations of "no elements" — null versus empty list [] — that the SCM
// API uses interchangeably, in whichever direction is needed to match the
// plan/state. Concretely, for each list field:
//   - plan/state is empty list [], API returned null  -> normalize to []
//   - plan/state is null,          API returned []     -> normalize to null
//
// Both directions prevent "Provider produced inconsistent result after apply"
// errors caused by the SCM API returning null or [] inconsistently for empty
// list fields. In Read, the "plan" argument is prior state (savestate), since
// Read receives no plan.
//
// The function handles:
//   - Top-level list fields (e.g., tag = [])
//   - Lists nested inside objects at any depth
//   - Lists nested inside list-of-objects elements
//
// It deliberately does NOT normalize when:
//   - Both plan and response are null (field was omitted)
//   - The plan is unknown (only an explicit null plan triggers [] -> null)
//   - The plan had a populated list but API returned null (data loss)
//   - The plan is null but the API returned a populated list (real data)
func NormalizeNullLists(ctx context.Context, plan types.Object, response types.Object) (types.Object, diag.Diagnostics) {
	var diags diag.Diagnostics

	if plan.IsNull() || plan.IsUnknown() || response.IsNull() || response.IsUnknown() {
		return response, diags
	}

	planAttrs := plan.Attributes()
	responseAttrs := response.Attributes()
	attrTypes := response.AttributeTypes(ctx)

	modified := false
	newAttrs := make(map[string]attr.Value, len(responseAttrs))
	for k, v := range responseAttrs {
		newAttrs[k] = v
	}

	for key, planVal := range planAttrs {
		responseVal, exists := responseAttrs[key]
		if !exists {
			continue
		}

		switch plannedTyped := planVal.(type) {
		case basetypes.ListValue:
			responseList, ok := responseVal.(basetypes.ListValue)
			if !ok {
				continue
			}

			if plannedTyped.IsNull() || plannedTyped.IsUnknown() {
				// Reverse direction: the plan (or prior state) had null, but the
				// API returned an empty list []. Coerce the empty list back to null
				// to satisfy Terraform's post-apply consistency check. We only do
				// this for an explicit null plan (not unknown) and only when the
				// response is a genuinely empty list — a populated response is real
				// data and must be preserved.
				if plannedTyped.IsNull() &&
					!responseList.IsNull() && !responseList.IsUnknown() &&
					len(responseList.Elements()) == 0 {
					newAttrs[key] = basetypes.NewListNull(responseList.ElementType(ctx))
					modified = true
					tflog.Debug(ctx, "Normalized empty list to null to match plan", map[string]interface{}{"field": key})
				}
				continue
			}

			if len(plannedTyped.Elements()) == 0 && responseList.IsNull() {
				// Plan had empty list [], API returned null -> normalize to []
				newAttrs[key] = basetypes.NewListValueMust(plannedTyped.ElementType(ctx), []attr.Value{})
				modified = true
				tflog.Debug(ctx, "Normalized null list to empty list to match plan", map[string]interface{}{"field": key})
			} else if !responseList.IsNull() && !responseList.IsUnknown() {
				// Both have elements - recurse into object elements to normalize nested lists
				plannedElements := plannedTyped.Elements()
				responseElements := responseList.Elements()
				if len(plannedElements) == len(responseElements) && len(plannedElements) > 0 {
					listModified := false
					newElements := make([]attr.Value, len(responseElements))
					copy(newElements, responseElements)

					for i := range responseElements {
						plannedObj, pOk := plannedElements[i].(basetypes.ObjectValue)
						responseObj, rOk := responseElements[i].(basetypes.ObjectValue)
						if !pOk || !rOk {
							continue
						}
						normalizedObj, d := NormalizeNullLists(ctx, plannedObj, responseObj)
						diags.Append(d...)
						if !d.HasError() && !normalizedObj.Equal(responseObj) {
							newElements[i] = normalizedObj
							listModified = true
						}
					}

					if listModified {
						newList, d := basetypes.NewListValue(plannedTyped.ElementType(ctx), newElements)
						diags.Append(d...)
						if !d.HasError() {
							newAttrs[key] = newList
							modified = true
						}
					}
				}
			}

		case basetypes.ObjectValue:
			responseObj, ok := responseVal.(basetypes.ObjectValue)
			if !ok {
				continue
			}

			if plannedTyped.IsNull() || plannedTyped.IsUnknown() || responseObj.IsNull() || responseObj.IsUnknown() {
				continue
			}

			normalizedObj, d := NormalizeNullLists(ctx, plannedTyped, responseObj)
			diags.Append(d...)
			if !d.HasError() && !normalizedObj.Equal(responseObj) {
				newAttrs[key] = normalizedObj
				modified = true
			}
		}
	}

	if modified {
		newObj, d := basetypes.NewObjectValue(attrTypes, newAttrs)
		diags.Append(d...)
		if !d.HasError() {
			return newObj, diags
		}
	}

	return response, diags
}
