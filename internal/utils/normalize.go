package utils

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// NormalizeNullLists reconciles null lists in a packed API response with the
// plan or prior state values. When the SCM API returns null for a list field
// but the plan/state had an empty list, the packed value is normalized to an
// empty list. This prevents "Provider produced inconsistent result after apply"
// errors that occur when Terraform compares the planned empty list against a
// null value returned by the API.
//
// This function recursively handles nested objects so that lists at any
// nesting depth are normalized correctly.
func NormalizeNullLists(ctx context.Context, packed types.Object, reference types.Object) (types.Object, diag.Diagnostics) {
	var diags diag.Diagnostics

	if packed.IsNull() || packed.IsUnknown() || reference.IsNull() || reference.IsUnknown() {
		return packed, diags
	}

	packedAttrs := packed.Attributes()
	refAttrs := reference.Attributes()
	newAttrs := make(map[string]attr.Value, len(packedAttrs))
	changed := false

	for key, packedVal := range packedAttrs {
		newAttrs[key] = packedVal

		refVal, hasRef := refAttrs[key]
		if !hasRef {
			continue
		}

		switch pv := packedVal.(type) {
		case basetypes.ListValue:
			if pv.IsNull() {
				if rv, ok := refVal.(basetypes.ListValue); ok && !rv.IsNull() && len(rv.Elements()) == 0 {
					emptyList, d := basetypes.NewListValue(pv.ElementType(ctx), []attr.Value{})
					diags.Append(d...)
					if !diags.HasError() {
						newAttrs[key] = emptyList
						changed = true
						tflog.Debug(ctx, "Normalized null list to empty list", map[string]interface{}{"field": key})
					}
				}
			}

		case basetypes.ObjectValue:
			if !pv.IsNull() && !pv.IsUnknown() {
				if rv, ok := refVal.(basetypes.ObjectValue); ok && !rv.IsNull() && !rv.IsUnknown() {
					normalized, d := NormalizeNullLists(ctx, pv, rv)
					diags.Append(d...)
					if !diags.HasError() && !normalized.Equal(pv) {
						newAttrs[key] = normalized
						changed = true
					}
				}
			}
		}
	}

	if !changed {
		return packed, diags
	}

	result, d := types.ObjectValue(packed.AttributeTypes(ctx), newAttrs)
	diags.Append(d...)
	return result, diags
}
