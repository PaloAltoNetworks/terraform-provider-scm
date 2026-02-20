package utils

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Helper to build a types.Object with the given attribute types and values.
func mustObject(t *testing.T, attrTypes map[string]attr.Type, attrValues map[string]attr.Value) types.Object {
	t.Helper()
	obj, diags := types.ObjectValue(attrTypes, attrValues)
	if diags.HasError() {
		t.Fatalf("failed to build object: %s", diags.Errors()[0].Detail())
	}
	return obj
}

// Helper to build an empty list of strings.
func emptyStringList(t *testing.T) basetypes.ListValue {
	t.Helper()
	list, diags := basetypes.NewListValue(basetypes.StringType{}, []attr.Value{})
	if diags.HasError() {
		t.Fatalf("failed to build empty list: %s", diags.Errors()[0].Detail())
	}
	return list
}

// Helper to build a list of strings with values.
func stringList(t *testing.T, vals ...string) basetypes.ListValue {
	t.Helper()
	elems := make([]attr.Value, len(vals))
	for i, v := range vals {
		elems[i] = basetypes.NewStringValue(v)
	}
	list, diags := basetypes.NewListValue(basetypes.StringType{}, elems)
	if diags.HasError() {
		t.Fatalf("failed to build string list: %s", diags.Errors()[0].Detail())
	}
	return list
}

// Common attribute types used across tests.
var testAttrTypes = map[string]attr.Type{
	"name": basetypes.StringType{},
	"tag":  basetypes.ListType{ElemType: basetypes.StringType{}},
}

func TestNormalizeNullLists_NullListToEmptyList(t *testing.T) {
	// Scenario: API returns null for tag, plan had empty list []
	// Expected: tag should be normalized to empty list []
	ctx := context.Background()

	packed := mustObject(t, testAttrTypes, map[string]attr.Value{
		"name": basetypes.NewStringValue("test"),
		"tag":  basetypes.NewListNull(basetypes.StringType{}),
	})

	reference := mustObject(t, testAttrTypes, map[string]attr.Value{
		"name": basetypes.NewStringValue("test"),
		"tag":  emptyStringList(t),
	})

	result, diags := NormalizeNullLists(ctx, packed, reference)
	if diags.HasError() {
		t.Fatalf("unexpected diagnostics: %s", diags.Errors()[0].Detail())
	}

	tagAttr := result.Attributes()["tag"]
	tagList, ok := tagAttr.(basetypes.ListValue)
	if !ok {
		t.Fatalf("expected tag to be ListValue, got %T", tagAttr)
	}
	if tagList.IsNull() {
		t.Error("expected tag to be non-null (empty list), but got null")
	}
	if len(tagList.Elements()) != 0 {
		t.Errorf("expected tag to have 0 elements, got %d", len(tagList.Elements()))
	}
}

func TestNormalizeNullLists_NullListStaysNull(t *testing.T) {
	// Scenario: API returns null for tag, plan also had null (field omitted)
	// Expected: tag should remain null
	ctx := context.Background()

	packed := mustObject(t, testAttrTypes, map[string]attr.Value{
		"name": basetypes.NewStringValue("test"),
		"tag":  basetypes.NewListNull(basetypes.StringType{}),
	})

	reference := mustObject(t, testAttrTypes, map[string]attr.Value{
		"name": basetypes.NewStringValue("test"),
		"tag":  basetypes.NewListNull(basetypes.StringType{}),
	})

	result, diags := NormalizeNullLists(ctx, packed, reference)
	if diags.HasError() {
		t.Fatalf("unexpected diagnostics: %s", diags.Errors()[0].Detail())
	}

	tagAttr := result.Attributes()["tag"]
	tagList, ok := tagAttr.(basetypes.ListValue)
	if !ok {
		t.Fatalf("expected tag to be ListValue, got %T", tagAttr)
	}
	if !tagList.IsNull() {
		t.Error("expected tag to remain null, but got non-null")
	}
}

func TestNormalizeNullLists_PopulatedListPassthrough(t *testing.T) {
	// Scenario: API returns actual data ["web", "prod"], plan had ["web", "prod"]
	// Expected: tag should pass through unchanged
	ctx := context.Background()

	packed := mustObject(t, testAttrTypes, map[string]attr.Value{
		"name": basetypes.NewStringValue("test"),
		"tag":  stringList(t, "web", "prod"),
	})

	reference := mustObject(t, testAttrTypes, map[string]attr.Value{
		"name": basetypes.NewStringValue("test"),
		"tag":  stringList(t, "web", "prod"),
	})

	result, diags := NormalizeNullLists(ctx, packed, reference)
	if diags.HasError() {
		t.Fatalf("unexpected diagnostics: %s", diags.Errors()[0].Detail())
	}

	tagAttr := result.Attributes()["tag"]
	tagList, ok := tagAttr.(basetypes.ListValue)
	if !ok {
		t.Fatalf("expected tag to be ListValue, got %T", tagAttr)
	}
	if tagList.IsNull() {
		t.Error("expected tag to be non-null, but got null")
	}
	if len(tagList.Elements()) != 2 {
		t.Errorf("expected tag to have 2 elements, got %d", len(tagList.Elements()))
	}
}

func TestNormalizeNullLists_NullListWithPopulatedReference(t *testing.T) {
	// Scenario: API returns null for tag, but plan had ["web"] (non-empty list)
	// Expected: tag should remain null (only normalize when reference is empty list)
	ctx := context.Background()

	packed := mustObject(t, testAttrTypes, map[string]attr.Value{
		"name": basetypes.NewStringValue("test"),
		"tag":  basetypes.NewListNull(basetypes.StringType{}),
	})

	reference := mustObject(t, testAttrTypes, map[string]attr.Value{
		"name": basetypes.NewStringValue("test"),
		"tag":  stringList(t, "web"),
	})

	result, diags := NormalizeNullLists(ctx, packed, reference)
	if diags.HasError() {
		t.Fatalf("unexpected diagnostics: %s", diags.Errors()[0].Detail())
	}

	tagAttr := result.Attributes()["tag"]
	tagList, ok := tagAttr.(basetypes.ListValue)
	if !ok {
		t.Fatalf("expected tag to be ListValue, got %T", tagAttr)
	}
	if !tagList.IsNull() {
		t.Error("expected tag to remain null when reference has non-empty list, but got non-null")
	}
}

func TestNormalizeNullLists_NullPackedObject(t *testing.T) {
	// Scenario: The entire packed object is null
	// Expected: return the null object as-is without error
	ctx := context.Background()

	packed := types.ObjectNull(testAttrTypes)
	reference := mustObject(t, testAttrTypes, map[string]attr.Value{
		"name": basetypes.NewStringValue("test"),
		"tag":  emptyStringList(t),
	})

	result, diags := NormalizeNullLists(ctx, packed, reference)
	if diags.HasError() {
		t.Fatalf("unexpected diagnostics: %s", diags.Errors()[0].Detail())
	}
	if !result.IsNull() {
		t.Error("expected null object to pass through as null")
	}
}

func TestNormalizeNullLists_NullReferenceObject(t *testing.T) {
	// Scenario: The reference object is null
	// Expected: return packed object as-is without error
	ctx := context.Background()

	packed := mustObject(t, testAttrTypes, map[string]attr.Value{
		"name": basetypes.NewStringValue("test"),
		"tag":  basetypes.NewListNull(basetypes.StringType{}),
	})
	reference := types.ObjectNull(testAttrTypes)

	result, diags := NormalizeNullLists(ctx, packed, reference)
	if diags.HasError() {
		t.Fatalf("unexpected diagnostics: %s", diags.Errors()[0].Detail())
	}

	tagAttr := result.Attributes()["tag"]
	tagList, ok := tagAttr.(basetypes.ListValue)
	if !ok {
		t.Fatalf("expected tag to be ListValue, got %T", tagAttr)
	}
	if !tagList.IsNull() {
		t.Error("expected tag to remain null when reference object is null")
	}
}

func TestNormalizeNullLists_MultipleListFields(t *testing.T) {
	// Scenario: Object has multiple list fields, some null and some populated
	// Expected: Only the null list with empty reference gets normalized
	ctx := context.Background()

	multiAttrTypes := map[string]attr.Type{
		"name":   basetypes.StringType{},
		"tag":    basetypes.ListType{ElemType: basetypes.StringType{}},
		"source": basetypes.ListType{ElemType: basetypes.StringType{}},
		"dest":   basetypes.ListType{ElemType: basetypes.StringType{}},
	}

	packed := mustObject(t, multiAttrTypes, map[string]attr.Value{
		"name":   basetypes.NewStringValue("test"),
		"tag":    basetypes.NewListNull(basetypes.StringType{}),     // null - should normalize
		"source": stringList(t, "10.0.0.0/8"),                      // populated - passthrough
		"dest":   basetypes.NewListNull(basetypes.StringType{}),     // null - stays null (ref is also null)
	})

	reference := mustObject(t, multiAttrTypes, map[string]attr.Value{
		"name":   basetypes.NewStringValue("test"),
		"tag":    emptyStringList(t),                                // empty list -> triggers normalize
		"source": stringList(t, "10.0.0.0/8"),                      // populated
		"dest":   basetypes.NewListNull(basetypes.StringType{}),     // null -> stays null
	})

	result, diags := NormalizeNullLists(ctx, packed, reference)
	if diags.HasError() {
		t.Fatalf("unexpected diagnostics: %s", diags.Errors()[0].Detail())
	}

	attrs := result.Attributes()

	// tag: should be normalized to empty list
	tagList := attrs["tag"].(basetypes.ListValue)
	if tagList.IsNull() {
		t.Error("expected tag to be normalized to empty list, got null")
	}
	if len(tagList.Elements()) != 0 {
		t.Errorf("expected tag to have 0 elements, got %d", len(tagList.Elements()))
	}

	// source: should pass through with value
	sourceList := attrs["source"].(basetypes.ListValue)
	if sourceList.IsNull() {
		t.Error("expected source to remain non-null")
	}
	if len(sourceList.Elements()) != 1 {
		t.Errorf("expected source to have 1 element, got %d", len(sourceList.Elements()))
	}

	// dest: should remain null
	destList := attrs["dest"].(basetypes.ListValue)
	if !destList.IsNull() {
		t.Error("expected dest to remain null")
	}
}

func TestNormalizeNullLists_NestedObjects(t *testing.T) {
	// Scenario: Object contains a nested object which itself has a null list
	// Expected: The nested list gets normalized recursively
	ctx := context.Background()

	innerAttrTypes := map[string]attr.Type{
		"members": basetypes.ListType{ElemType: basetypes.StringType{}},
	}

	outerAttrTypes := map[string]attr.Type{
		"name":  basetypes.StringType{},
		"group": types.ObjectType{AttrTypes: innerAttrTypes},
	}

	innerPacked := mustObject(t, innerAttrTypes, map[string]attr.Value{
		"members": basetypes.NewListNull(basetypes.StringType{}),
	})

	innerRef := mustObject(t, innerAttrTypes, map[string]attr.Value{
		"members": emptyStringList(t),
	})

	packed := mustObject(t, outerAttrTypes, map[string]attr.Value{
		"name":  basetypes.NewStringValue("test"),
		"group": innerPacked,
	})

	reference := mustObject(t, outerAttrTypes, map[string]attr.Value{
		"name":  basetypes.NewStringValue("test"),
		"group": innerRef,
	})

	result, diags := NormalizeNullLists(ctx, packed, reference)
	if diags.HasError() {
		t.Fatalf("unexpected diagnostics: %s", diags.Errors()[0].Detail())
	}

	groupAttr := result.Attributes()["group"].(basetypes.ObjectValue)
	membersList := groupAttr.Attributes()["members"].(basetypes.ListValue)
	if membersList.IsNull() {
		t.Error("expected nested members list to be normalized to empty list, got null")
	}
	if len(membersList.Elements()) != 0 {
		t.Errorf("expected nested members to have 0 elements, got %d", len(membersList.Elements()))
	}
}

func TestNormalizeNullLists_NoChangeReturnsOriginal(t *testing.T) {
	// Scenario: Nothing needs normalizing
	// Expected: Returns the original object (Equal check)
	ctx := context.Background()

	packed := mustObject(t, testAttrTypes, map[string]attr.Value{
		"name": basetypes.NewStringValue("test"),
		"tag":  stringList(t, "web"),
	})

	reference := mustObject(t, testAttrTypes, map[string]attr.Value{
		"name": basetypes.NewStringValue("test"),
		"tag":  stringList(t, "web"),
	})

	result, diags := NormalizeNullLists(ctx, packed, reference)
	if diags.HasError() {
		t.Fatalf("unexpected diagnostics: %s", diags.Errors()[0].Detail())
	}

	if !result.Equal(packed) {
		t.Error("expected result to equal the original packed object when no changes needed")
	}
}

func TestNormalizeNullLists_StringFieldsUnchanged(t *testing.T) {
	// Scenario: Verify that non-list attributes are not affected
	// Expected: String values pass through untouched
	ctx := context.Background()

	packed := mustObject(t, testAttrTypes, map[string]attr.Value{
		"name": basetypes.NewStringValue("my-address"),
		"tag":  basetypes.NewListNull(basetypes.StringType{}),
	})

	reference := mustObject(t, testAttrTypes, map[string]attr.Value{
		"name": basetypes.NewStringValue("my-address"),
		"tag":  emptyStringList(t),
	})

	result, diags := NormalizeNullLists(ctx, packed, reference)
	if diags.HasError() {
		t.Fatalf("unexpected diagnostics: %s", diags.Errors()[0].Detail())
	}

	nameAttr := result.Attributes()["name"].(basetypes.StringValue)
	if nameAttr.ValueString() != "my-address" {
		t.Errorf("expected name to be 'my-address', got '%s'", nameAttr.ValueString())
	}
}
