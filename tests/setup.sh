#!/usr/bin/env bash
set -euo pipefail

# =============================================================================
# Setup: generate test workspaces from examples/resources
# =============================================================================
# For each examples/resources/<prefix>_* (scm_*, ztna_*, ...):
#   1. Creates tests/<prefix>_<resource>/
#   2. Generates main.tf = matching provider block + resource.tf (always overwritten)
#   3. Generates a .tftest.hcl with create + update + auto-destroy
#
# Update strategy per resource:
#   - root-level description field → update description
#   - root-level comment field     → update comment
#   - neither                      → skip update test
#
# Usage:
#   ./tests/setup.sh                  # all resources (all prefixes)
#   ./tests/setup.sh scm_address      # single scm resource
#   ./tests/setup.sh ztna_connector   # single ztna resource
# =============================================================================

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
BASE_DIR="$(cd "$SCRIPT_DIR/.." && pwd)"
RESOURCES_DIR="${BASE_DIR}/examples/resources"
PROVIDER_FILE="${BASE_DIR}/examples/provider/provider.tf"

if [ ! -f "$PROVIDER_FILE" ]; then
  echo "Error: ${PROVIDER_FILE} not found."
  echo "This file must exist with at least the terraform { required_providers { ... } } block."
  exit 1
fi

# Resources that don't have an id attribute (from OpenAPI specs)
NO_ID_RESOURCES="scm_bgp_routing scm_bandwidth_allocation"

# Resources that use "oid" instead of "id" as their identifier attribute
OID_RESOURCES="ztna_connector ztna_fqdn_application ztna_subnet ztna_wildcard ztna_connector_group ztna_connector_scheduled_upgrade ztna_connector_group_scheduled_upgrade"

# -----------------------------------------------------------------------------
# Check if a field exists at root level (brace depth 1) inside any resource block.
# Uses brace-depth tracking so indentation doesn't matter.
# Returns 0 (true) if found, 1 (false) otherwise.
# -----------------------------------------------------------------------------
has_root_level_field() {
  local field="$1"
  local tf_file="$2"

  awk -v field="$field" '
    /^resource "/ { in_res = 1; depth = 0 }
    /^variable / || /^data / || /^locals / || /^output / { in_res = 0 }
    in_res && /{/ { depth++ }
    in_res && /}/ { depth--; if (depth == 0) in_res = 0 }
    in_res && depth == 1 {
      pattern = "^[[:space:]]*" field "[[:space:]]*="
      if ($0 ~ pattern) { found = 1; exit }
    }
    END { if (found) exit 0; else exit 1 }
  ' "$tf_file" 2>/dev/null
}

# -----------------------------------------------------------------------------
# Determine which field to use for update testing.
# Returns: "description", "comment", or "none"
# -----------------------------------------------------------------------------
get_update_field() {
  local resource_tf="$1"

  if has_root_level_field "description" "$resource_tf"; then
    echo "description"
  elif has_root_level_field "comment" "$resource_tf"; then
    echo "comment"
  else
    echo "none"
  fi
}

# -----------------------------------------------------------------------------
# Replace a root-level field value with the test variable.
# Only modifies lines at brace depth 1 inside resource blocks.
# Outputs the modified file content to stdout.
# -----------------------------------------------------------------------------
inject_test_variable() {
  local field="$1"
  local tf_file="$2"

  awk -v field="$field" '
    /^resource "/ { in_res = 1; depth = 0 }
    /^variable / || /^data / || /^locals / || /^output / { in_res = 0 }
    in_res && /{/ { depth++ }
    in_res && /}/ { depth--; if (depth == 0) in_res = 0 }
    in_res && depth == 1 {
      pattern = field "[[:space:]]*=[[:space:]]*\"([^\"]*)\""
      if ($0 ~ pattern) {
        sub(/"[^"]*"/, "(var.test_update_value != \"\" ? var.test_update_value : &)")
        print
        next
      }
    }
    { print }
  ' "$tf_file"
}

generate_tftest() {
  local resource_tf="$1"
  local output_file="$2"
  local update_field="$3"

  # Extract all resource "<type>" "<label>" declarations
  local assertions=""
  while IFS= read -r line; do
    local type label
    type=$(echo "$line" | sed -E 's/^resource "([^"]+)".*/\1/')
    label=$(echo "$line" | sed -E 's/^resource "[^"]+" "([^"]+)".*/\1/')
    if echo "$NO_ID_RESOURCES" | grep -qw "$type"; then
      assertions="${assertions}
  # ${type}.${label} — no id attribute, apply success is the test
"
    elif echo "$OID_RESOURCES" | grep -qw "$type"; then
      assertions="${assertions}
  assert {
    condition     = ${type}.${label}.oid != \"\"
    error_message = \"${type}.${label} was not created\"
  }
"
    else
      assertions="${assertions}
  assert {
    condition     = ${type}.${label}.id != \"\"
    error_message = \"${type}.${label} was not created\"
  }
"
    fi
  done < <(grep -E '^resource "' "$resource_tf")

  if [ -z "$assertions" ]; then
    return 1
  fi

  # If no updatable field, just generate create + destroy test
  if [ "$update_field" = "none" ]; then
    cat > "$output_file" <<EOF
# Auto-generated lifecycle test
# Verifies: create → destroy (no updatable field found)

run "create" {
  command = apply
${assertions}}

# terraform test auto-destroys all resources after the last run block
EOF
    return 0
  fi

  # Build update assertions — only for primary resources that have the field at root level
  local primary_type
  primary_type=$(basename "$(dirname "$resource_tf")")
  local update_assertions=""
  local update_value="Updated by Terraform test"
  local current_label="" current_type="" has_field=false depth=0 in_res=0

  while IFS= read -r line; do
    # Detect resource block start
    if echo "$line" | grep -qE '^resource "'; then
      # Emit assertion for the previous resource if it had the field
      if [ -n "$current_label" ] && [ "$current_type" = "$primary_type" ] && [ "$has_field" = true ]; then
        update_assertions="${update_assertions}
  assert {
    condition     = ${current_type}.${current_label}.${update_field} == \"${update_value}\"
    error_message = \"${current_type}.${current_label} ${update_field} was not updated\"
  }
"
      fi
      current_type=$(echo "$line" | sed -E 's/^resource "([^"]+)".*/\1/')
      current_label=$(echo "$line" | sed -E 's/^resource "[^"]+" "([^"]+)".*/\1/')
      has_field=false
      in_res=1
      depth=0
      continue
    fi

    if [ $in_res -eq 1 ]; then
      # Track brace depth
      case "$line" in *"{"*) depth=$((depth + 1)) ;; esac
      case "$line" in *"}"*) depth=$((depth - 1)); [ $depth -eq 0 ] && in_res=0 ;; esac

      # Check for root-level field (depth 1)
      if [ $depth -eq 1 ] && echo "$line" | grep -qE "^[[:space:]]*${update_field}[[:space:]]*="; then
        has_field=true
      fi
    fi
  done < "$resource_tf"

  # Handle the last resource
  if [ -n "$current_label" ] && [ "$current_type" = "$primary_type" ] && [ "$has_field" = true ]; then
    update_assertions="${update_assertions}
  assert {
    condition     = ${current_type}.${current_label}.${update_field} == \"${update_value}\"
    error_message = \"${current_type}.${current_label} ${update_field} was not updated\"
  }
"
  fi

  cat > "$output_file" <<EOF
# Auto-generated lifecycle test
# Verifies: create → update ${update_field} → destroy

run "create" {
  command = apply
${assertions}}

run "update" {
  command = apply

  variables {
    test_update_value = "${update_value}"
  }
${update_assertions}}

# terraform test auto-destroys all resources after the last run block
EOF
}

# =============================================================================
# Main
# =============================================================================

targets=()
if [ $# -gt 0 ]; then
  targets=("$@")
else
  for d in "${RESOURCES_DIR}"/*_*; do
    [ -d "$d" ] && targets+=("$(basename "$d")")
  done
fi

count=0
skipped=0

for resource_name in "${targets[@]}"; do
  resource_dir="${RESOURCES_DIR}/${resource_name}"

  if [ ! -f "${resource_dir}/resource.tf" ]; then
    skipped=$((skipped + 1))
    continue
  fi

  test_dir="${SCRIPT_DIR}/${resource_name}"
  mkdir -p "$test_dir"

  # Determine update strategy
  update_field=$(get_update_field "${resource_dir}/resource.tf")

  # Determine the resource prefix (everything before the first underscore).
  resource_prefix="${resource_name%%_*}"

  # Build main.tf from scratch:
  # 1. Extract only the provider "<resource_prefix>" { ... } block from provider.tf
  #    using awk brace-depth tracking.
  # 2. Append a terraform {} block declaring only this prefix's provider.
  # This avoids the "Duplicate required provider" warning that occurs when two
  # local names (scm, ztna) point to the same source binary.
  awk -v prefix="$resource_prefix" '
    /^provider "/ {
      # Extract the provider name between the first pair of quotes
      line = $0; sub(/^provider "/, "", line); sub(/".*/, "", line)
      if (line == prefix) { in_block=1; depth=0 }
    }
    in_block {
      print
      n = split($0, chars, "")
      for (i=1; i<=n; i++) {
        if (chars[i] == "{") depth++
        if (chars[i] == "}") {
          depth--
          if (depth == 0) { in_block=0 }
        }
      }
    }
  ' "$PROVIDER_FILE" > "${test_dir}/main.tf"

  cat >> "${test_dir}/main.tf" <<TFBLOCK

terraform {
  required_providers {
    ${resource_prefix} = {
      source = "paloaltonetworks-local/scm"
    }
  }
}

TFBLOCK

  if [ "$update_field" != "none" ]; then
    # Add the test variable
    cat >> "${test_dir}/main.tf" <<'VAREOF'
variable "test_update_value" {
  type    = string
  default = ""
}

VAREOF

    # Append resource.tf with root-level field replaced by the variable
    inject_test_variable "$update_field" "${resource_dir}/resource.tf" >> "${test_dir}/main.tf"
  else
    # No update field — just copy resource.tf as-is
    cat "${resource_dir}/resource.tf" >> "${test_dir}/main.tf"
  fi

  # Generate .tftest.hcl (always regenerate)
  if generate_tftest "${resource_dir}/resource.tf" "${test_dir}/${resource_name}.tftest.hcl" "$update_field"; then
    echo "  [new] ${resource_name}/ (update: ${update_field})"
  else
    echo "  [skip] ${resource_name}/ - no resource blocks found in resource.tf"
    skipped=$((skipped + 1))
    continue
  fi

  count=$((count + 1))
done

echo ""
echo "Setup complete: ${count} workspace(s) ready, ${skipped} skipped."
