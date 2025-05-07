package provider

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
)

func TestMoveSecurityRule_Validation(t *testing.T) {
	// Test cases for the moveSecurityRule validation logic
	tests := []struct {
		name        string
		ruleID      string
		relPos      string
		targetRule  string
		expectError bool
	}{
		{
			name:        "empty relative position",
			ruleID:      "test-rule-id",
			relPos:      "",
			targetRule:  "",
			expectError: false, // Should return early without error
		},
		{
			name:        "before position missing target",
			ruleID:      "test-rule-id",
			relPos:      "before",
			targetRule:  "",
			expectError: true, // Should error with missing target
		},
		{
			name:        "after position missing target",
			ruleID:      "test-rule-id",
			relPos:      "after",
			targetRule:  "",
			expectError: true, // Should error with missing target
		},
		{
			name:        "top position",
			ruleID:      "test-rule-id",
			relPos:      "top",
			targetRule:  "",
			expectError: false,
		},
		{
			name:        "bottom position",
			ruleID:      "test-rule-id",
			relPos:      "bottom",
			targetRule:  "",
			expectError: false,
		},
		{
			name:        "before position with target",
			ruleID:      "test-rule-id",
			relPos:      "before",
			targetRule:  "target-rule-id",
			expectError: false,
		},
		{
			name:        "after position with target",
			ruleID:      "test-rule-id",
			relPos:      "after",
			targetRule:  "target-rule-id",
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a resource with nil client to test only validation logic
			resource := &securityRuleResource{}
			err := resource.moveSecurityRule(context.Background(), tt.ruleID, tt.relPos, tt.targetRule)
			
			if tt.expectError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), "target_rule must be specified")
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestSecurityRuleResourceState_WithPositioning(t *testing.T) {
	// Test that the securityRuleRsModel correctly handles positioning attributes
	tests := []struct {
		name        string
		relativePos string
		targetRule  string
	}{
		{
			name:        "no positioning",
			relativePos: "",
			targetRule:  "",
		},
		{
			name:        "position top",
			relativePos: "top",
			targetRule:  "",
		},
		{
			name:        "position bottom",
			relativePos: "bottom",
			targetRule:  "",
		},
		{
			name:        "position before with target",
			relativePos: "before",
			targetRule:  "target-rule",
		},
		{
			name:        "position after with target",
			relativePos: "after",
			targetRule:  "target-rule",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a sample rule state
			state := securityRuleRsModel{
				Name:             types.StringValue("Test Rule"),
				Action:           types.StringValue("allow"),
				RelativePosition: types.StringNull(),
				TargetRule:       types.StringNull(),
			}

			// Set positioning values if provided
			if tt.relativePos != "" {
				state.RelativePosition = types.StringValue(tt.relativePos)
				assert.Equal(t, tt.relativePos, state.RelativePosition.ValueString())
			}
			if tt.targetRule != "" {
				state.TargetRule = types.StringValue(tt.targetRule)
				assert.Equal(t, tt.targetRule, state.TargetRule.ValueString())
			}
		})
	}
}