#!/usr/bin/env bash
set -euo pipefail

# =============================================================================
# Test Runner for SCM Provider
# =============================================================================
# Flow per resource:
#   1. terraform init (skipped when dev_overrides active)
#   2. terraform test (apply + assert + auto-destroy)
#
# Usage:
#   ./tests/run_tests.sh scm_address        # single resource
#   ./tests/run_tests.sh --all              # all resources with .tftest.hcl
#
# Local: populate examples/provider/provider.tf with credentials
# CI/CD: credentials come from env vars set in the repo
#
# Requires: Terraform >= 1.6
# =============================================================================

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
BASE_DIR="$(cd "$SCRIPT_DIR/.." && pwd)"

# Overall tracking
PASSED=()
FAILED=()
FAILED_ERRORS=()
FAILED_ALREADY_EXISTS=()
SKIPPED=()

# Per-phase tracking
CREATE_PASSED=()
CREATE_FAILED=()
CREATE_FAILED_ERRORS=()
UPDATE_PASSED=()
UPDATE_FAILED=()
UPDATE_FAILED_ERRORS=()
UPDATE_SKIPPED=()
DELETE_PASSED=()
DELETE_FAILED=()
DELETE_FAILED_ERRORS=()

# Resources to skip (require special setup, not testable in isolation)
SKIP_LIST="scm_auto_vpn_setting scm_auto_vpn_cluster"

# -----------------------------------------------------------------------------
# Detect dev_overrides — when active, terraform init must be skipped.
# -----------------------------------------------------------------------------
has_dev_overrides() {
  local rc_file="${TF_CLI_CONFIG_FILE:-$HOME/.terraformrc}"
  [ -f "$rc_file" ] && grep -q "dev_overrides" "$rc_file" 2>/dev/null
}

DEV_OVERRIDES=false
if has_dev_overrides; then
  DEV_OVERRIDES=true
fi

# -----------------------------------------------------------------------------
# Parse terraform test output for per-phase results
# -----------------------------------------------------------------------------
parse_phases() {
  local resource_name="$1"
  local test_output="$2"

  # Create phase
  if grep -q 'run "create"... pass' "$test_output" 2>/dev/null; then
    CREATE_PASSED+=("$resource_name")
  elif grep -q 'run "create"... fail' "$test_output" 2>/dev/null; then
    local err
    err=$(grep -m1 "^Error:" "$test_output" 2>/dev/null || echo "Unknown error")
    CREATE_FAILED+=("$resource_name")
    CREATE_FAILED_ERRORS+=("$err")
  fi

  # Update phase
  if grep -q 'run "update"... pass' "$test_output" 2>/dev/null; then
    UPDATE_PASSED+=("$resource_name")
  elif grep -q 'run "update"... fail' "$test_output" 2>/dev/null; then
    local err
    err=$(grep -m1 "^Error:" "$test_output" 2>/dev/null || echo "Unknown error")
    UPDATE_FAILED+=("$resource_name")
    UPDATE_FAILED_ERRORS+=("$err")
  elif ! grep -q 'run "update"' "$test_output" 2>/dev/null; then
    UPDATE_SKIPPED+=("$resource_name")
  fi

  # Delete phase (tearing down)
  if grep -q 'tearing down' "$test_output" 2>/dev/null; then
    # If the overall test passed, teardown succeeded
    if grep -qE '\.tftest\.hcl\.\.\. pass$' "$test_output" 2>/dev/null; then
      DELETE_PASSED+=("$resource_name")
    else
      # Check if teardown itself had errors (vs create/update failure)
      # If create or update failed, the teardown still runs to clean up
      # We count delete as failed only if there's an error during tearing down
      local teardown_err
      teardown_err=$(sed -n '/tearing down/,$ p' "$test_output" | grep -m1 "^Error:" 2>/dev/null || echo "")
      if [ -n "$teardown_err" ]; then
        DELETE_FAILED+=("$resource_name")
        DELETE_FAILED_ERRORS+=("$teardown_err")
      else
        DELETE_PASSED+=("$resource_name")
      fi
    fi
  fi
}

# -----------------------------------------------------------------------------
# Run a single resource test
# -----------------------------------------------------------------------------
test_resource() {
  local resource_name="$1"
  local test_dir="${SCRIPT_DIR}/${resource_name}"

  echo ""
  echo "=========================================="
  echo " Testing: ${resource_name}"
  echo "=========================================="

  # Check skip list
  if echo "$SKIP_LIST" | grep -qw "$resource_name"; then
    echo "  SKIP: In skip list"
    SKIPPED+=("${resource_name}")
    return 0
  fi

  # Validate workspace
  if [ ! -d "$test_dir" ]; then
    echo "  SKIP: No test directory. Run ./tests/setup.sh first."
    SKIPPED+=("${resource_name}")
    return 0
  fi

  if [ ! -f "${test_dir}/main.tf" ]; then
    echo "  SKIP: No main.tf. Run ./tests/setup.sh first."
    SKIPPED+=("${resource_name}")
    return 0
  fi

  if ! ls "${test_dir}"/*.tftest.hcl > /dev/null 2>&1; then
    echo "  SKIP: No .tftest.hcl file found."
    SKIPPED+=("${resource_name}")
    return 0
  fi

  cd "$test_dir"

  # Clean any leftover state from previous interrupted runs
  rm -f terraform.tfstate terraform.tfstate.backup

  # --- terraform init ---
  if [ "$DEV_OVERRIDES" = true ]; then
    echo "  [init] terraform init (dev_overrides active, ignoring provider warnings)..."
    terraform init -no-color > /dev/null 2>&1 || true
  else
    echo "  [init] terraform init..."
    if ! terraform init -no-color 2>&1 | tail -5; then
      echo "  FAIL: terraform init failed"
      FAILED+=("${resource_name}")
      FAILED_ERRORS+=("init: Error: terraform init failed")
      return 1
    fi
  fi

  # --- terraform test ---
  local test_output="/tmp/tftest-${resource_name}-$$.log"
  local test_exit=0

  if [ -n "${GITHUB_ACTIONS:-}" ]; then
    terraform test -no-color > "$test_output" 2>&1 || test_exit=$?
  else
    echo "  [test] terraform test..."
    echo "  ----------------------------------------"
    terraform test -no-color 2>&1 | tee "$test_output" || test_exit=$?
    echo "  ----------------------------------------"
  fi

  if [ $test_exit -eq 0 ]; then
    echo "  PASS: ${resource_name}"
    PASSED+=("$resource_name")
    parse_phases "$resource_name" "$test_output"
  else
    if grep -qi "already exists\|OBJECT_ALREADY_EXISTS" "$test_output" 2>/dev/null; then
      echo "  FAIL: ${resource_name} (object already exists)"
      FAILED_ALREADY_EXISTS+=("${resource_name}")
    else
      local failed_step err_msg
      failed_step=$(grep -E 'run ".*"\.\.\. fail' "$test_output" 2>/dev/null | head -1 | sed -E 's/.*run "([^"]+)".*/\1/' || echo "unknown")
      err_msg=$(grep -m1 "^Error:" "$test_output" 2>/dev/null | head -1 || echo "Unknown error")
      echo "  FAIL: ${resource_name} (${failed_step})"
      echo "        ${err_msg}"
      FAILED+=("${resource_name}")
      FAILED_ERRORS+=("${failed_step}: ${err_msg}")
    fi
  fi

  rm -f "$test_output"
  return $test_exit
}

# =============================================================================
# Main
# =============================================================================

resources_to_test=()
for arg in "$@"; do
  case "$arg" in
    --all)
      for d in "${SCRIPT_DIR}"/scm_*/; do
        [ -d "$d" ] || continue
        resources_to_test+=("$(basename "$d")")
      done
      ;;
    *) resources_to_test+=("$arg") ;;
  esac
done

if [ ${#resources_to_test[@]} -eq 0 ]; then
  echo "Usage:"
  echo "  $0 <resource_name>    # test one resource"
  echo "  $0 --all              # test all resources"
  echo ""
  echo "Available tests:"
  for d in "${SCRIPT_DIR}"/scm_*/; do
    [ -d "$d" ] && echo "  - $(basename "$d")"
  done
  exit 0
fi

echo "=========================================="
echo " Terraform Provider SCM - Test Runner"
echo "=========================================="
echo " Tests: ${resources_to_test[*]}"
echo ""

# Run tests
overall_exit=0
for resource in "${resources_to_test[@]}"; do
  test_resource "$resource" || overall_exit=1
done

# =============================================================================
# Terminal Summary
# =============================================================================
echo ""
echo "=========================================="
echo " SUMMARY"
echo "=========================================="
echo " Passed:          ${#PASSED[@]}"
echo " Failed:          ${#FAILED[@]}"
echo " Already Exists:  ${#FAILED_ALREADY_EXISTS[@]}"
echo " Skipped:         ${#SKIPPED[@]}"
echo ""

[ ${#PASSED[@]} -gt 0 ] && echo " Passed:" && for r in "${PASSED[@]}"; do echo "   [PASS] $r"; done
[ ${#SKIPPED[@]} -gt 0 ] && echo " Skipped:" && for r in "${SKIPPED[@]}"; do echo "   [SKIP] $r"; done

if [ ${#FAILED[@]} -gt 0 ]; then
  echo ""
  echo " Failed:"
  for i in "${!FAILED[@]}"; do
    echo "   [FAIL] ${FAILED[$i]}"
    echo "          ${FAILED_ERRORS[$i]}"
  done
fi

if [ ${#FAILED_ALREADY_EXISTS[@]} -gt 0 ]; then
  echo ""
  echo " Resources blocked by pre-existing objects (need manual cleanup):"
  for r in "${FAILED_ALREADY_EXISTS[@]}"; do echo "   $r"; done
fi

# =============================================================================
# GitHub Actions Step Summary
# =============================================================================
if [ -n "${GITHUB_STEP_SUMMARY:-}" ]; then
  {
    echo "## Terraform Resource Test Results"
    echo ""
    echo "### Overview"
    echo ""
    echo "| | Count |"
    echo "|---|---|"
    echo "| Passed | ${#PASSED[@]} |"
    echo "| Failed | ${#FAILED[@]} |"
    echo "| Already Exists | ${#FAILED_ALREADY_EXISTS[@]} |"
    echo "| Skipped | ${#SKIPPED[@]} |"
    echo ""

    # Per-phase breakdown
    echo "### Phase Breakdown"
    echo ""
    echo "| Phase | Passed | Failed | Skipped |"
    echo "|---|---|---|---|"
    echo "| Create | ${#CREATE_PASSED[@]} | ${#CREATE_FAILED[@]} | 0 |"
    echo "| Update | ${#UPDATE_PASSED[@]} | ${#UPDATE_FAILED[@]} | ${#UPDATE_SKIPPED[@]} |"
    echo "| Delete | ${#DELETE_PASSED[@]} | ${#DELETE_FAILED[@]} | 0 |"
    echo ""

    # Failed resources (actual errors, not "already exists")
    if [ ${#FAILED[@]} -gt 0 ]; then
      echo "### Failed (${#FAILED[@]})"
      echo ""
      echo "| Resource | Step | Error |"
      echo "|---|---|---|"
      for i in "${!FAILED[@]}"; do
        local_step=$(echo "${FAILED_ERRORS[$i]}" | cut -d: -f1)
        local_err=$(echo "${FAILED_ERRORS[$i]}" | cut -d: -f2-)
        echo "| ${FAILED[$i]} | ${local_step} | ${local_err} |"
      done
      echo ""
    fi

    # Create failures
    if [ ${#CREATE_FAILED[@]} -gt 0 ]; then
      echo "<details><summary>Create Failed (${#CREATE_FAILED[@]})</summary>"
      echo ""
      echo "| Resource | Error |"
      echo "|---|---|"
      for i in "${!CREATE_FAILED[@]}"; do
        echo "| ${CREATE_FAILED[$i]} | ${CREATE_FAILED_ERRORS[$i]} |"
      done
      echo ""
      echo "</details>"
      echo ""
    fi

    # Update failures
    if [ ${#UPDATE_FAILED[@]} -gt 0 ]; then
      echo "<details><summary>Update Failed (${#UPDATE_FAILED[@]})</summary>"
      echo ""
      echo "| Resource | Error |"
      echo "|---|---|"
      for i in "${!UPDATE_FAILED[@]}"; do
        echo "| ${UPDATE_FAILED[$i]} | ${UPDATE_FAILED_ERRORS[$i]} |"
      done
      echo ""
      echo "</details>"
      echo ""
    fi

    # Update skipped
    if [ ${#UPDATE_SKIPPED[@]} -gt 0 ]; then
      echo "<details><summary>Update Skipped (${#UPDATE_SKIPPED[@]})</summary>"
      echo ""
      for r in "${UPDATE_SKIPPED[@]}"; do echo "- $r"; done
      echo ""
      echo "</details>"
      echo ""
    fi

    # Delete failures
    if [ ${#DELETE_FAILED[@]} -gt 0 ]; then
      echo "<details><summary>Delete Failed (${#DELETE_FAILED[@]})</summary>"
      echo ""
      echo "| Resource | Error |"
      echo "|---|---|"
      for i in "${!DELETE_FAILED[@]}"; do
        echo "| ${DELETE_FAILED[$i]} | ${DELETE_FAILED_ERRORS[$i]} |"
      done
      echo ""
      echo "</details>"
      echo ""
    fi

    # Passed (collapsed)
    if [ ${#PASSED[@]} -gt 0 ]; then
      echo "<details><summary>Passed (${#PASSED[@]})</summary>"
      echo ""
      for r in "${PASSED[@]}"; do echo "- $r"; done
      echo ""
      echo "</details>"
      echo ""
    fi

    # Already exists (collapsed)
    if [ ${#FAILED_ALREADY_EXISTS[@]} -gt 0 ]; then
      echo "<details><summary>Already Exists - need cleanup (${#FAILED_ALREADY_EXISTS[@]})</summary>"
      echo ""
      for r in "${FAILED_ALREADY_EXISTS[@]}"; do echo "- $r"; done
      echo ""
      echo "</details>"
      echo ""
    fi

    # Skipped (collapsed)
    if [ ${#SKIPPED[@]} -gt 0 ]; then
      echo "<details><summary>Skipped (${#SKIPPED[@]})</summary>"
      echo ""
      for r in "${SKIPPED[@]}"; do echo "- $r"; done
      echo ""
      echo "</details>"
    fi
  } >> "$GITHUB_STEP_SUMMARY"
fi

exit $overall_exit
