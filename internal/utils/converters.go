/*
THIS FILE IS EMBEDED INTO THE GENERATOR AND COPIED TO THE OUTPUT

This entire file is embedded into the terraform generator and will be written
directly to terraform-provider-scm/internal/utils/converters.go with only
the package name changed from "terraform" to "utils".

The functions in this file are NOT used by the generator itself but are intended
to be used by the generated Terraform provider code at runtime to convert between
SDK models and Terraform framework types.

Any changes made to this file will be reflected in the generated provider.
*/

package utils

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

// Look for logs like "Terraform Provider HTTP REQUEST" and "Terraform Provider HTTP RESPONSE" if you want to check the TF_LOGS

var (
	// This one handles acronyms at the start: "EAPTTLS" -> "EAP_TTLS"
	matchAcronym = regexp.MustCompile("([A-Z]+)([A-Z][a-z])")
	// This one handles caps in the middle: "WithPAP" -> "With_PAP"
	matchCap = regexp.MustCompile("([a-z\\d])([A-Z])")
)

// toSnakeCase converts a string to snake_case.
func toSnakeCase(str string) string {
	// First, check if the string is all-caps, digits, and underscores.
	// This handles tags like "EAP_TTLS_with_PAP" or "CHAP".
	isAllUpperAndUnderscore := true
	for _, r := range str {
		if !(unicode.IsUpper(r) || unicode.IsDigit(r) || r == '_') {
			isAllUpperAndUnderscore = false
			break
		}
	}
	if isAllUpperAndUnderscore {
		// If so, just lowercase it.
		// "EAP_TTLS_with_PAP" -> "eap_ttls_with_pap"
		// "CHAP" -> "chap"
		return strings.ToLower(str)
	}

	// If not, it's likely camelCase (e.g., "EAPTTLSWithPAP" or "RadiusCertProfile").
	// Apply standard camelCase to snake_case logic.
	snake := matchAcronym.ReplaceAllString(str, "${1}_${2}")
	snake = matchCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(strings.ReplaceAll(snake, "-", "_"))
}

// getPrimitivePacker returns the correct Terraform framework function call for "packing" a
// non-nil primitive value from the SDK into a framework type.
func getPrimitivePacker(goType string) string {
	switch goType {
	case "basetypes.StringValue":
		return "basetypes.NewStringValue"
	case "basetypes.BoolValue":
		return "basetypes.NewBoolValue"
	case "basetypes.Int64Value":
		return "basetypes.NewInt64Value"
	// Add other primitive types as needed
	default:
		return "basetypes.NewStringValue" // Default case
	}
}

// getPrimitiveNuller returns the correct Terraform framework function call for creating a
// "null" framework type, used when the SDK value is nil.
func getPrimitiveNuller(goType string) string {
	switch goType {
	case "basetypes.StringValue":
		return "basetypes.NewStringNull()"
	case "basetypes.BoolValue":
		return "basetypes.NewBoolNull()"
	case "basetypes.Int64Value":
		return "basetypes.NewInt64Null()"
	// Add other primitive types as needed
	default:
		return "basetypes.NewStringNull()" // Default case
	}
}

// dict creates a map from a list of key-value pairs.
// This is a common utility function for Go templates to allow passing multiple
// arguments to a sub-template.
func dict(values ...interface{}) (map[string]interface{}, error) {
	if len(values)%2 != 0 {
		return nil, fmt.Errorf("invalid dict call: number of arguments must be even")
	}
	result := make(map[string]interface{}, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].(string)
		if !ok {
			return nil, fmt.Errorf("invalid dict call: key is not a string: %v", values[i])
		}
		result[key] = values[i+1]
	}
	return result, nil
}
