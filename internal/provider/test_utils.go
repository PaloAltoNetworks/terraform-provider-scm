package provider

import (
	"context"
)

// Constants and helper functions for testing
const (
	// TestIdSeparator is the separator used in Terraform IDs for tests
	TestIdSeparator = "|"
)

// TestIsObjectNotFound returns true if the error indicates the object was not found (test implementation)
func TestIsObjectNotFound(err error) bool {
	return err != nil && err.Error() == "object not found"
}

// SecurityRuleMoveTestClient is a mock client for testing security rule move operations
type SecurityRuleMoveTestClient struct {
	// Tracks calls to the Move method
	MoveInputs []*SecurityRuleMoveInput
}

// SecurityRuleMoveInput tracks inputs for the Move method calls
type SecurityRuleMoveInput struct {
	RuleID       string
	Position     string
	TargetRuleID string
}

// NewSecurityRuleMoveTestClient creates a new mock client for testing security rule move operations
func NewSecurityRuleMoveTestClient() *SecurityRuleMoveTestClient {
	return &SecurityRuleMoveTestClient{
		MoveInputs: make([]*SecurityRuleMoveInput, 0),
	}
}

// MoveSecurityRule is a mock implementation of the Move method
func (c *SecurityRuleMoveTestClient) MoveSecurityRule(ctx context.Context, ruleID string, position string, targetRuleID string) error {
	// Record the inputs
	c.MoveInputs = append(c.MoveInputs, &SecurityRuleMoveInput{
		RuleID:       ruleID,
		Position:     position,
		TargetRuleID: targetRuleID,
	})
	return nil
}
