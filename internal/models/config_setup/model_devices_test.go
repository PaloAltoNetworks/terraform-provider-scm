package models

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Test_DevicesResourceSchema_RequiredAttributes verifies that the resource schema
// has the correct required vs optional vs computed attributes.
func Test_DevicesResourceSchema_RequiredAttributes(t *testing.T) {
	s := DevicesResourceSchema

	requiredFields := []string{"id", "folder"}
	for _, field := range requiredFields {
		attr, ok := s.Attributes[field]
		if !ok {
			t.Errorf("Resource schema missing required attribute: %s", field)
			continue
		}
		if sa, ok := attr.(schema.StringAttribute); ok {
			if !sa.Required {
				t.Errorf("Resource attribute %q should be Required", field)
			}
		}
	}
}

func Test_DevicesResourceSchema_OptionalAttributes(t *testing.T) {
	s := DevicesResourceSchema

	optionalFields := []string{"description", "display_name", "labels", "snippets"}
	for _, field := range optionalFields {
		attr, ok := s.Attributes[field]
		if !ok {
			t.Errorf("Resource schema missing optional attribute: %s", field)
			continue
		}
		switch a := attr.(type) {
		case schema.StringAttribute:
			if !a.Optional {
				t.Errorf("Resource attribute %q should be Optional", field)
			}
			if !a.Computed {
				t.Errorf("Resource attribute %q should also be Computed", field)
			}
		case schema.ListAttribute:
			if !a.Optional {
				t.Errorf("Resource attribute %q should be Optional", field)
			}
			if !a.Computed {
				t.Errorf("Resource attribute %q should also be Computed", field)
			}
		}
	}
}

func Test_DevicesResourceSchema_ComputedOnlyAttributes(t *testing.T) {
	s := DevicesResourceSchema

	computedOnly := []string{
		"name", "hostname", "ip_address", "ipv6_address",
		"mac_address", "family", "model", "software_version",
		"is_connected", "tfid",
	}
	for _, field := range computedOnly {
		attr, ok := s.Attributes[field]
		if !ok {
			t.Errorf("Resource schema missing computed attribute: %s", field)
			continue
		}
		switch a := attr.(type) {
		case schema.StringAttribute:
			if !a.Computed {
				t.Errorf("Resource attribute %q should be Computed", field)
			}
			if a.Required {
				t.Errorf("Resource attribute %q should NOT be Required", field)
			}
		case schema.BoolAttribute:
			if !a.Computed {
				t.Errorf("Resource attribute %q should be Computed", field)
			}
		}
	}
}

func Test_DevicesResourceSchema_IdRequiresReplace(t *testing.T) {
	s := DevicesResourceSchema

	attr, ok := s.Attributes["id"]
	if !ok {
		t.Fatal("Resource schema missing 'id' attribute")
	}
	sa, ok := attr.(schema.StringAttribute)
	if !ok {
		t.Fatal("'id' attribute is not a StringAttribute")
	}
	if len(sa.PlanModifiers) == 0 {
		t.Error("'id' attribute should have plan modifiers (RequiresReplace)")
	}
}

// Test_DevicesDataSourceSchema_Attributes verifies data source schema.
func Test_DevicesDataSourceSchema_Attributes(t *testing.T) {
	s := DevicesDataSourceSchema

	// ID is required for data source lookup
	idAttr, ok := s.Attributes["id"]
	if !ok {
		t.Fatal("Data source schema missing 'id' attribute")
	}
	if !idAttr.IsRequired() {
		t.Error("Data source 'id' should be Required")
	}

	// All other fields should be computed
	computedFields := []string{
		"name", "folder", "description", "display_name",
		"labels", "snippets", "hostname", "ip_address",
		"ipv6_address", "mac_address", "family", "model",
		"software_version", "is_connected", "tfid",
	}
	for _, field := range computedFields {
		attr, ok := s.Attributes[field]
		if !ok {
			t.Errorf("Data source schema missing attribute: %s", field)
			continue
		}
		if !attr.IsComputed() {
			t.Errorf("Data source attribute %q should be Computed", field)
		}
	}
}

// Test_DevicesListDataSourceSchema_Attributes verifies the list data source schema.
func Test_DevicesListDataSourceSchema_Attributes(t *testing.T) {
	s := DevicesListDataSourceSchema

	// Check optional filter attributes
	optionalFields := []string{"limit", "offset", "folder"}
	for _, field := range optionalFields {
		attr, ok := s.Attributes[field]
		if !ok {
			t.Errorf("List data source schema missing attribute: %s", field)
			continue
		}
		if !attr.IsOptional() {
			t.Errorf("List data source attribute %q should be Optional", field)
		}
	}

	// Check computed attributes
	computedFields := []string{"tfid", "total", "data"}
	for _, field := range computedFields {
		attr, ok := s.Attributes[field]
		if !ok {
			t.Errorf("List data source schema missing attribute: %s", field)
			continue
		}
		if !attr.IsComputed() {
			t.Errorf("List data source attribute %q should be Computed", field)
		}
	}
}

// Test_DevicesTf_AttrTypes verifies that AttrTypes returns all expected keys.
func Test_DevicesTf_AttrTypes(t *testing.T) {
	d := DevicesTf{}
	attrTypes := d.AttrTypes()

	expectedKeys := []string{
		"tfid", "id", "name", "folder", "description", "display_name",
		"labels", "snippets", "hostname", "ip_address", "ipv6_address",
		"mac_address", "family", "model", "software_version", "is_connected",
	}

	for _, key := range expectedKeys {
		if _, ok := attrTypes[key]; !ok {
			t.Errorf("AttrTypes missing key: %s", key)
		}
	}

	if len(attrTypes) != len(expectedKeys) {
		t.Errorf("AttrTypes has %d keys, expected %d", len(attrTypes), len(expectedKeys))
	}
}
