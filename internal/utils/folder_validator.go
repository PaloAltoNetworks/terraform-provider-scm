package utils

/*
THIS FILE IS EMBEDDED INTO THE GENERATOR AND COPIED TO THE OUTPUT

This entire file is embedded into the terraform generator and will be written
directly to terraform-provider-scm/internal/utils/folder_validator.go with only
the package name changed from "terraform" to "utils".

The functions in this file are NOT used by the generator itself but are intended
for the generated Terraform provider to validate folder field values.
*/

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// FolderValueValidator validates that deprecated folder names are not used
type FolderValueValidator struct{}

// Description returns a description of the validator
func (v FolderValueValidator) Description(ctx context.Context) string {
	return "validates that deprecated folder values are not used"
}

// MarkdownDescription returns a markdown description of the validator
func (v FolderValueValidator) MarkdownDescription(ctx context.Context) string {
	return "validates that deprecated folder values are not used"
}

// ValidateString performs the validation
func (v FolderValueValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	value := req.ConfigValue.ValueString()

	// Check for deprecated values and provide helpful error messages
	// Note: Both "Shared" and "Prisma Access" are accepted as valid values
	switch value {
	case "All Firewalls":
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid Folder Value",
			fmt.Sprintf("The folder value 'All Firewalls' is not allowed. Please use 'ngfw-shared' instead."),
		)
	case "Global":
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid Folder Value",
			fmt.Sprintf("The folder value 'Global' is not allowed. Please use 'All' instead."),
		)
	case "Explicit Proxy":
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid Folder Value",
			fmt.Sprintf("The folder value 'Explicit Proxy' is not allowed. Please use 'Mobile Users Explicit Proxy' instead."),
		)
	case "Access Agent":
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid Folder Value",
			fmt.Sprintf("The folder value 'Access Agent' is not allowed. Please use 'Mobile Users' instead."),
		)
	}
}

// FolderValidator returns a new FolderValueValidator
func FolderValidator() validator.String {
	return FolderValueValidator{}
}

// FolderAPIToState translates folder values from API responses to state values.
// This handles deprecated folder values that should be normalized.
// Note: "Shared" and "Prisma Access" are both valid and will be preserved as-is.
// The user's configured value is preserved through parameter restoration in CRUD operations.
func FolderAPIToState(apiValue string) string {
	// Standard translations for deprecated values
	translationMap := map[string]string{
		"All Firewalls":               "ngfw-shared",                 // Translate to correct value
		"Global":                      "All",                         // Translate to correct value
		"ngfw-shared":                 "ngfw-shared",                 // ngfw-shared (passthrough)
		"All":                         "All",                         // All (passthrough)
		"Explicit Proxy":              "Mobile Users Explicit Proxy", // Translate to correct value
		"Mobile Users Explicit Proxy": "Mobile Users Explicit Proxy", // Mobile Users Explicit Proxy (passthrough)
		"Access Agent":                "Mobile Users",                // Translate to correct value
		"Mobile Users":                "Mobile Users",                // Mobile Users (passthrough)
		"Shared":                      "Shared",                      // Shared (passthrough)
		"Prisma Access":               "Prisma Access",               // Prisma Access (passthrough)
	}

	if stateValue, ok := translationMap[apiValue]; ok {
		return stateValue
	}

	// If not in map, return as-is (custom folder names pass through)
	return apiValue
}
