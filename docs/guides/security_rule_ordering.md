---
page_title: "Security Rule Ordering"
subcategory: "Guides"
description: |-
  Guide for ordering security rules in Palo Alto Networks Strata Cloud Manager
---

# Security Rule Ordering

This guide covers how to maintain a specific order of security rules in SCM using Terraform.

## Ordering Options

The SCM Terraform provider now supports explicit ordering of security rules, allowing you to:

1. Place rules at the top or bottom of the ruleset
2. Place rules before or after specific existing rules
3. Maintain consistent ordering when modifying rules

## New Attributes

The `scm_security_rule` resource supports two new attributes for rule positioning:

| Attribute | Description |
|-----------|-------------|
| `relative_position` | Where to position the rule: `"before"`, `"after"`, `"top"`, or `"bottom"`. If not specified, rules are added at the bottom of the ruleset. |
| `target_rule` | The name or ID of the target rule to position this rule relative to. Required when `relative_position` is `"before"` or `"after"`. |

## Example Usage

### Creating rules in a specific order

```terraform
# Create a rule at the top of the ruleset
resource "scm_security_rule" "rule1" {
  name             = "Rule-1"
  # ... other attributes ...
  
  position         = "pre"
  relative_position = "top"
}

# Create a rule after rule1
resource "scm_security_rule" "rule2" {
  name             = "Rule-2"
  # ... other attributes ...
  
  position         = "pre"
  relative_position = "after"
  target_rule      = scm_security_rule.rule1.name
}

# Create a rule before rule1
resource "scm_security_rule" "rule0" {
  name             = "Rule-0"
  # ... other attributes ...
  
  position         = "pre"
  relative_position = "before"
  target_rule      = scm_security_rule.rule1.name
}
```

### Inserting a new rule between existing rules

```terraform
# Insert a new rule between Rule-1 and Rule-2
resource "scm_security_rule" "rule_inserted" {
  name             = "Rule-Inserted"
  # ... other attributes ...
  
  position         = "pre"
  relative_position = "after"
  target_rule      = scm_security_rule.rule1.name
}

# Update Rule-2 to be after the inserted rule
resource "scm_security_rule" "rule2" {
  name             = "Rule-2"
  # ... other attributes ...
  
  position         = "pre"
  relative_position = "after"
  target_rule      = scm_security_rule.rule_inserted.name
}
```

## Notes

- If positioning fails (e.g., target rule doesn't exist), the rule is still created but positioned at the default location, and a warning is generated.
- When updating a rule, you can change its position by modifying the `relative_position` and `target_rule` attributes.
- Rule ordering is processed after the rule is created or updated.