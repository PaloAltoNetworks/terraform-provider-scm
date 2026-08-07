package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

// LocalErrorDetailCauseInfo mirrors the generated ErrorDetailCauseInfo in the SDK
type LocalErrorDetailCauseInfo struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details"` // interface{} is required for dynamic JSON details
}

// LocalGenericError mirrors the generated GenericError struct in the SDK
type LocalGenericError struct {
	Errors    []LocalErrorDetailCauseInfo `json:"_errors,omitempty"`
	RequestId *string                     `json:"_request_id,omitempty"`
	// Note: We omit AdditionalProperties and anything else unnecessary for simplicity.
}

// SCMErrorAccessor defines the methods we expect the generated SDK error (GenericOpenAPIError) to have.
// We must use this interface because we cannot directly reference the generated type.
type SCMErrorAccessor interface {
	Error() string
	Model() interface{} // Accesses the underlying decoded error struct (GenericError)
	Body() []byte       // Accesses the raw HTTP response body for debugging
}

// formatNonZeroRefsDetails produces a human-readable message from a NON_ZERO_REFS details map.
// The API returns something like:
//
//	errors: [{extra: "container/[folder]/address-group/[group]/static/[addr]", type: "NON_ZERO_REFS"}]
//
// We extract each blocked resource and its referencing container into plain English.
func formatNonZeroRefsDetails(d map[string]interface{}) string {
	raw, ok := d["errors"]
	if !ok {
		return ""
	}
	// The value is []interface{}, each element a map[string]interface{}
	entries, ok := raw.([]interface{})
	if !ok || len(entries) == 0 {
		return ""
	}

	var lines []string
	for _, entry := range entries {
		m, ok := entry.(map[string]interface{})
		if !ok {
			continue
		}
		// extra can be a string or a []interface{} depending on the API version.
		var extra string
		switch v := m["extra"].(type) {
		case string:
			extra = v
		case []interface{}:
			if len(v) > 0 {
				extra, _ = v[0].(string)
			}
		}
		if extra == "" {
			continue
		}
		if t, _ := m["type"].(string); t != "NON_ZERO_REFS" {
			continue
		}
		// extra looks like: container/[ngfw-shared]/address-group/[scm_address_group_1]/static/[tf_address_2]
		// Split on "/" and extract bracketed segments.
		parts := strings.Split(extra, "/")
		var segments []string
		for _, p := range parts {
			if strings.HasPrefix(p, "[") && strings.HasSuffix(p, "]") {
				segments = append(segments, p[1:len(p)-1])
			}
		}
		// We expect at least [folder, referencing-resource-name, object-name]
		if len(segments) < 3 {
			lines = append(lines, fmt.Sprintf("Still referenced via: %s", extra))
			continue
		}
		objectName := segments[len(segments)-1]
		refName := segments[len(segments)-2]
		// Derive a human-readable resource type from the path segment before the name.
		refType := "resource"
		for i, p := range parts {
			if strings.HasPrefix(p, "[") && p[1:len(p)-1] == refName && i > 0 {
				refType = parts[i-1]
				break
			}
		}
		lines = append(lines, fmt.Sprintf(
			"'%s' is still referenced by %s '%s'. Remove the reference from that resource first, then re-run terraform apply.",
			objectName, refType, refName,
		))
	}

	return strings.Join(lines, "\n      ")
}

// PrintScmError safely extracts, unmarshals, and formats the detailed API error body.
func PrintScmError(err error) string {
	scmErr, ok := err.(SCMErrorAccessor)
	if !ok {
		// Fallback 1: Error is not the expected SDK wrapper type (e.g., network error)
		return fmt.Sprintf("API Request Failed: %s", err.Error())
	}

	// Access the underlying decoded model (the generated GenericError struct)
	model := scmErr.Model()
	modelJSON, _ := json.Marshal(model)

	var localGenErr LocalGenericError
	unmarshalFailed := json.Unmarshal(modelJSON, &localGenErr) != nil

	// --- Check 1: Detailed Formatting (Only if Unmarshal Succeeded AND we have errors) ---
	if !unmarshalFailed && len(localGenErr.Errors) > 0 {
		var output strings.Builder

		// High-level header
		reqID := "N/A"
		if localGenErr.RequestId != nil {
			reqID = *localGenErr.RequestId
		}
		output.WriteString(fmt.Sprintf("--- Detailed API Error (Request ID: %s, %d Errors) ---\n", reqID, len(localGenErr.Errors)))

		// Iterate over the Errors array
		for i, e := range localGenErr.Errors {
			output.WriteString(fmt.Sprintf("\n  Error %d: Code=[%s], Message=\"%s\"\n", i+1, e.Code, e.Message))

			// Structured Details Logic
			if e.Details == nil {
				output.WriteString("    Details: (none)\n")
				continue
			}

			output.WriteString("    Details:\n")
			switch d := e.Details.(type) {
			case string:
				output.WriteString(fmt.Sprintf("      %s\n", d))
			case map[string]interface{}:
				if e.Code == "API_I00013" {
					if msg := formatNonZeroRefsDetails(d); msg != "" {
						output.WriteString(fmt.Sprintf("      %s\n", msg))
						continue
					}
				}
				for key, val := range d {
					output.WriteString(fmt.Sprintf("        - %s: %v\n", key, val))
				}
			default:
				// Fallback: Dump the details object (useful for debugging)
				output.WriteString("      (Error):\n")
				// Marshal the specific details field for clarity
				detailDump, _ := json.MarshalIndent(e.Details, "        ", "  ")
				output.WriteString(string(detailDump))
				output.WriteString("\n")
			}
		}
		output.WriteString("--------------------------------------------------")

		return output.String()
	}

	// --- Check 2: Raw Debug Dump (If structured error was absent or malformed) ---

	rawBody := scmErr.Body()

	var dump strings.Builder
	dump.WriteString("--- API ERROR DEBUG DUMP ---\n")

	// Add the simple error message provided by the SDK (e.g., "400 Bad Request")
	dump.WriteString(fmt.Sprintf("SDK Status/Summary: %s\n", scmErr.Error()))
	dump.WriteString("------------------------------------------------------------\n")

	if len(rawBody) > 0 {
		dump.WriteString("Raw Response Body:\n")
		// Pretty-print the raw JSON body if it exists and is valid
		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, rawBody, "", "  "); err == nil {
			dump.Write(prettyJSON.Bytes())
		} else {
			// Body exists but isn't JSON; dump as raw text
			dump.Write(rawBody)
		}
		dump.WriteString("\n")
	} else {
		dump.WriteString("Raw API response body was empty or inaccessible.\n")
	}
	dump.WriteString("------------------------------------------------------------")

	return dump.String()
}
