package provider

import (
	"context"
)

// MovePositionType represents where to move the rule in relation to the target rule
type MovePositionType string

const (
	// MovePositionBefore moves the rule before the target rule
	MovePositionBefore MovePositionType = "before"
	// MovePositionAfter moves the rule after the target rule
	MovePositionAfter MovePositionType = "after"
	// MovePositionTop moves the rule to the top of the ruleset
	MovePositionTop MovePositionType = "top"
	// MovePositionBottom moves the rule to the bottom of the ruleset
	MovePositionBottom MovePositionType = "bottom"
)

// SecurityRuleMover is an interface for moving security rules
type SecurityRuleMover interface {
	MoveSecurityRule(ctx context.Context, ruleID string, position string, targetRule string) error
}
