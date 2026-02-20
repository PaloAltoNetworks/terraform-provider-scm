# Regression Test Plans — Null List Normalization Fix

**Branch:** `fix/normalize-null-lists-tag-issue`
**Issue:** [#92](https://github.com/PaloAltoNetworks/terraform-provider-scm/issues/92)
**Date:** 2026-02-20
**Last updated:** 2026-02-20

---

## Summary of Changes

| Change | Files | Description |
|--------|-------|-------------|
| **New utility function** | `internal/utils/normalize.go` | `NormalizeNullLists()` — recursively normalizes null lists to empty lists when the plan had `[]` |
| **New unit tests** | `internal/utils/normalize_test.go` | 10 unit tests covering all normalization scenarios |
| **Acceptance test suite** | 4 test files (see below) | 15 acceptance tests covering the original bug, multi-list resources, cross-category smoke, CRUD lifecycle, and parallel creation |
| **Test infrastructure** | `internal/provider/provider_test.go` | Shared provider factories and `testAccPreCheck` helper |
| **Resource Create/Update insertion** | 121 resource files | 7-line normalization block inserted in both `Create` and `Update` functions |

**Total committed:** 122 files changed, 1,771 lines inserted, 0 lines deleted
**Total with tests:** 128 files changed (122 committed + 6 new test/doc files)

### Acceptance Test Files

| File | Tests | Covers Plans |
|------|-------|--------------|
| `internal/provider/address_nulllist_test.go` | 4 tests | 2A, 2B, 2C, 5A, 7 |
| `internal/provider/address_group_service_nulllist_test.go` | 3 tests | 3B, 5B |
| `internal/provider/security_rule_nulllist_test.go` | 3 tests | 3A, 5, 7 (security rule variant) |
| `internal/provider/cross_category_nulllist_test.go` | 5 tests | 6D, 6E, 6F, 6G, 10 |
| `internal/provider/provider_test.go` | — | Shared test infrastructure |

---

## Test Plan 1: Unit Tests — `NormalizeNullLists` Function

**File:** `internal/utils/normalize_test.go`
**Run:** `go test ./internal/utils/ -v`
**Status:** ✅ **IMPLEMENTED — 10 tests**

| # | Test Name | Scenario | Expected Result | Status |
|---|-----------|----------|-----------------|--------|
| 1 | `TestNormalizeNullLists_NullListToEmptyList` | API returns `null` for tag, plan had `[]` | Tag normalized to `[]` | ✅ PASS |
| 2 | `TestNormalizeNullLists_NullListStaysNull` | API returns `null`, plan also had `null` (omitted) | Tag stays `null` | ✅ PASS |
| 3 | `TestNormalizeNullLists_PopulatedListPassthrough` | API returns `["web","prod"]`, plan had same | Tag passes through unchanged | ✅ PASS |
| 4 | `TestNormalizeNullLists_NullListWithPopulatedReference` | API returns `null`, plan had `["web"]` | Tag stays `null` (don't invent data) | ✅ PASS |
| 5 | `TestNormalizeNullLists_NullPackedObject` | Entire packed object is `null` | Returns `null` object without error | ✅ PASS |
| 6 | `TestNormalizeNullLists_NullReferenceObject` | Reference (plan) object is `null` | Returns packed object as-is | ✅ PASS |
| 7 | `TestNormalizeNullLists_MultipleListFields` | 3 list fields: one null→normalize, one populated, one null→stay null | Each field handled independently | ✅ PASS |
| 8 | `TestNormalizeNullLists_NestedObjects` | Object contains nested object with null list | Nested list normalized recursively | ✅ PASS |
| 9 | `TestNormalizeNullLists_NoChangeReturnsOriginal` | Nothing needs normalizing | Returns original object (Equal check) | ✅ PASS |
| 10 | `TestNormalizeNullLists_StringFieldsUnchanged` | Non-list attributes present | String fields pass through untouched | ✅ PASS |

---

## Test Plan 2: Terraform Acceptance — The Original Bug (`tag = []`)

**Purpose:** Reproduce and verify the fix for issue #92 — the exact user-reported bug.
**File:** `internal/provider/address_nulllist_test.go`
**Status:** ✅ **IMPLEMENTED — 4 tests**

### Test 2A: Address Object with `tag = []` (Create)

**Test function:** `TestAccAddress_EmptyTag_Create`

```hcl
resource "scm_address" "test" {
  folder     = "All"
  name       = "tf-test-<random>"
  ip_netmask = "10.0.0.0/24"
  tag        = []
}
```

| Step | Action | Expected |
|------|--------|----------|
| 1 | Create with `tag = []` | Applies successfully — **no "inconsistent result" error** |
| 2 | Verify state | `tag.# = 0` (empty list, not null) |
| 3 | PlanOnly (idempotency) | Shows "No changes" — empty plan |

### Test 2B: Address Object with `tag = []` (Update)

**Test function:** `TestAccAddress_EmptyTag_Update`

```hcl
# Step 1: Create with populated tag
tag = ["test-tag"]

# Step 2: Update to empty tag
tag = []
```

| Step | Action | Expected |
|------|--------|----------|
| 1 | Create with `tag = ["test-tag"]` | Creates with 1 tag element |
| 2 | Update to `tag = []` | Applies successfully — **no "inconsistent result" error** |
| 3 | PlanOnly (idempotency) | Shows "No changes" |

### Test 2C: Address Object with tag omitted (no regression)

**Test function:** `TestAccAddress_TagOmitted`

```hcl
resource "scm_address" "test" {
  folder     = "All"
  name       = "tf-test-<random>"
  ip_netmask = "10.0.0.0/24"
  # tag intentionally omitted
}
```

| Step | Action | Expected |
|------|--------|----------|
| 1 | Create without tag attribute | Applies successfully |
| 2 | Verify state | `tag.#` not present (null/omitted) |
| 3 | PlanOnly (idempotency) | Shows "No changes" |

---

## Test Plan 3: Terraform Acceptance — Multiple List Fields

**Purpose:** Verify normalization works correctly for resources with multiple list-type attributes.

### Test 3A: Security Rule (many list fields)

**File:** `internal/provider/security_rule_nulllist_test.go`
**Status:** ✅ **IMPLEMENTED — `TestAccSecurityRule_EmptyTagAndLists`**

```hcl
resource "scm_security_rule" "test" {
  folder      = "All"
  position    = "pre"
  name        = "tf-test-<random>"
  from        = ["any"]
  to          = ["any"]
  source      = ["any"]
  destination = ["any"]
  application = ["any"]
  service     = ["any"]
  action      = "allow"
  tag         = []
  category    = []
}
```

| Step | Action | Expected |
|------|--------|----------|
| 1 | Create with mixed populated/empty lists | Creates successfully |
| 2 | Verify state | `tag.# = 0`, `category.# = 0`, populated lists have correct counts |
| 3 | PlanOnly (idempotency) | Shows "No changes" |

### Test 3B: Address Group (static and dynamic list variants)

**File:** `internal/provider/address_group_service_nulllist_test.go`
**Status:** ✅ **IMPLEMENTED — `TestAccAddressGroup_EmptyTag`**

```hcl
resource "scm_address" "dep" {
  folder     = "All"
  name       = "tf-test-addr-<random>"
  ip_netmask = "192.168.1.0/24"
}

resource "scm_address_group" "test" {
  folder = "All"
  name   = "tf-test-<random>"
  tag    = []
  static = [scm_address.dep.name]
}
```

| Step | Action | Expected |
|------|--------|----------|
| 1 | Create with dependency chain | Creates successfully |
| 2 | Verify state | `tag.# = 0`, `static.# = 1` |
| 3 | PlanOnly (idempotency) | Shows "No changes" |

---

## Test Plan 4: Terraform Acceptance — Nested Object Lists

**Purpose:** Verify recursive normalization handles nested objects containing lists.
**Status:** ⏳ **NOT YET IMPLEMENTED** — covered at the unit-test level by `TestNormalizeNullLists_NestedObjects`; acceptance tests would require deep-nesting resources like anti-spyware or DNS security profiles.

### Test 4A: Anti-Spyware Profile (deeply nested structure)

```hcl
resource "scm_anti_spyware_profile" "test_nested" {
  folder      = "All"
  name        = "regression-test-nested"
  description = "Testing nested list normalization"
  tag         = []

  rules = [{
    name     = "test-rule"
    severity = ["critical", "high"]
    category = "spyware"
    action = {
      alert = true
    }
  }]
}
```

| Step | Action | Expected |
|------|--------|----------|
| 1 | `terraform apply` | Creates successfully |
| 2 | `terraform plan` | Shows "No changes" |
| 3 | Clean up | Destroy |

### Test 4B: DNS Security Profile (nested lists at multiple levels)

```hcl
resource "scm_dns_security_profile" "test_nested" {
  folder      = "All"
  name        = "regression-test-dns-nested"
  description = "Testing DNS security nested normalization"
  tag         = []
}
```

| Step | Action | Expected |
|------|--------|----------|
| 1 | `terraform apply` | Creates successfully |
| 2 | `terraform plan` | Shows "No changes" |
| 3 | Clean up | Destroy |

---

## Test Plan 5: Terraform Acceptance — No Normalization Needed (Negative Tests)

**Purpose:** Confirm the fix doesn't alter behavior when normalization is unnecessary.
**Status:** ✅ **IMPLEMENTED**

### Test 5A: Populated tags remain populated

**File:** `internal/provider/address_nulllist_test.go`
**Test function:** `TestAccAddress_PopulatedTag`

```hcl
resource "scm_address" "test" {
  folder     = "All"
  name       = "tf-test-<random>"
  ip_netmask = "10.0.0.0/24"
  tag        = ["web-server", "production"]
}
```

| Step | Action | Expected |
|------|--------|----------|
| 1 | Create with populated tags | Creates with `tag.# = 2` |
| 2 | Verify state | `tag.0 = "web-server"`, `tag.1 = "production"` |
| 3 | PlanOnly (idempotency) | Shows "No changes" |

### Test 5B: Omitted optional lists stay null

**File:** `internal/provider/address_group_service_nulllist_test.go`
**Test functions:** `TestAccService_TagOmitted`, `TestAccService_EmptyTag`

```hcl
resource "scm_service" "test" {
  folder = "All"
  name   = "tf-test-<random>"
  protocol = {
    tcp = {
      port = "8080"
    }
  }
  # tag intentionally omitted
}
```

| Step | Action | Expected |
|------|--------|----------|
| 1 | Create without tag attribute | Creates successfully |
| 2 | Verify state | `tag.#` not present (null) |
| 3 | PlanOnly (idempotency) | Shows "No changes" |

**Bonus:** `TestAccService_EmptyTag` also verifies `tag = []` works for Service resources.

---

## Test Plan 6: Terraform Acceptance — Cross-Category Resource Coverage

**Purpose:** Smoke-test one resource from each provider category to confirm the identical insertion works across all 7 directories.
**File:** `internal/provider/cross_category_nulllist_test.go`
**Status:** ✅ **PARTIALLY IMPLEMENTED** — 4 of 7 categories covered with coded tests; remaining categories can be validated manually

| # | Category | Resource | Test Function | Status |
|---|----------|----------|---------------|--------|
| 6A | `config_setup` | `scm_folder` | — | ⏳ Manual |
| 6B | `deployment_services` | `scm_remote_network` | — | ⏳ Manual |
| 6C | `device_settings` | `scm_service_route` | — | ⏳ Manual |
| 6D | `identity_services` | `scm_authentication_profile` | `TestAccAuthenticationProfile_EmptyTag` | ✅ Implemented |
| 6E | `network_services` | `scm_zone` | `TestAccZone_EmptyTag` | ✅ Implemented |
| 6F | `objects` | `scm_application_group` | `TestAccApplicationGroup_EmptyTag` | ✅ Implemented |
| 6G | `security_services` | `scm_url_category` | `TestAccUrlCategory_EmptyTag` | ✅ Implemented |

**For each resource:**

| Step | Action | Expected |
|------|--------|----------|
| 1 | Create with `tag = []` (or equivalent empty list) | Creates successfully — no inconsistent result error |
| 2 | PlanOnly (idempotency) | Shows "No changes" |

---

## Test Plan 7: Terraform Acceptance — Create/Read/Update/Delete Lifecycle

**Purpose:** Full CRUD lifecycle to verify normalization doesn't break any operation.
**Status:** ✅ **IMPLEMENTED**

### Test 7A: Address CRUD Lifecycle

**File:** `internal/provider/address_nulllist_test.go`
**Test function:** `TestAccAddress_Lifecycle`

```hcl
# Phase 1: Create with tag = []
# Phase 2: Update to tag = ["web"]
# Phase 3: Update back to tag = []
# Phase 4: Update description only (tag = [] unchanged)
# Phase 5: PlanOnly idempotency check
```

| Step | Action | Expected |
|------|--------|----------|
| 1 | Create with `tag = []` | Creates successfully — no error |
| 2 | Update — add `tag = ["web"]` | Tag changes to `["web"]` |
| 3 | Update — back to `tag = []` | Tag reverts to empty list — no error |
| 4 | Update — change `description` only | Description updates, `tag = []` unchanged |
| 5 | PlanOnly (idempotency) | Shows "No changes" |

### Test 7B: Security Rule CRUD Lifecycle

**File:** `internal/provider/security_rule_nulllist_test.go`
**Test function:** `TestAccSecurityRule_TagLifecycle`

| Step | Action | Expected |
|------|--------|----------|
| 1 | Create with `tag = []` | Creates successfully |
| 2 | Update — add `tag = ["test-tag"]` | Tag changes to populated |
| 3 | Update — back to `tag = []` | Tag reverts to empty list — no error |

---

## Test Plan 8: Import Consistency

**Purpose:** Verify that `terraform import` + `terraform plan` is not broken by the normalization. (Note: Read functions were NOT modified — normalization only applies to Create/Update.)
**Status:** ⏳ **NOT YET IMPLEMENTED** — manual validation recommended
```hcl
resource "scm_address" "imported" {
  folder     = "All"
  name       = "existing-resource-name"
  ip_netmask = "10.0.0.0/24"
  tag        = []
}
```

| Step | Action | Expected |
|------|--------|----------|
| 1 | Create resource manually or via separate apply | Resource exists in SCM |
| 2 | `terraform import scm_address.imported <id>` | Import succeeds |
| 3 | `terraform plan` | May show drift for `tag` (null in Read vs `[]` in config) — this is **expected pre-existing behavior** since Read is not modified |
| 4 | `terraform apply` | Resolves any drift — normalization kicks in on Update |
| 5 | `terraform plan` | Shows "No changes" |

---

## Test Plan 9: Error Path Verification

**Purpose:** Verify the normalization block doesn't mask or alter error handling.
**Status:** ⏳ **NOT YET IMPLEMENTED** — manual or mock-based validation

| # | Scenario | How to Trigger | Expected |
|---|----------|----------------|----------|
| 9A | API error on Create | Use invalid folder name or duplicate resource name | Error returned to user — normalization code is never reached |
| 9B | API error on Update | Try to update a deleted resource | Error returned — normalization not reached |
| 9C | Normalization returns diagnostics | (Edge case — only if types.ObjectValue fails internally) | Diagnostics propagated to user, operation halts |

---

## Test Plan 10: Concurrency / Parallel Resource Creation

**Purpose:** Verify normalization works when Terraform creates multiple resources in parallel.
**File:** `internal/provider/cross_category_nulllist_test.go`
**Status:** ✅ **IMPLEMENTED — `TestAccAddress_ParallelCreate_EmptyTag`**

```hcl
resource "scm_address" "parallel_1" {
  folder     = "All"
  name       = "tf-test-<random>"
  ip_netmask = "10.0.0.1/32"
  tag        = []
}

resource "scm_address" "parallel_2" {
  folder     = "All"
  name       = "tf-test-<random>"
  ip_netmask = "10.0.0.2/32"
  tag        = []
}

resource "scm_address" "parallel_3" {
  folder     = "All"
  name       = "tf-test-<random>"
  ip_netmask = "10.0.0.3/32"
  tag        = ["shared"]
}
```

| Step | Action | Expected |
|------|--------|----------|
| 1 | Create all 3 resources (parallel) | All 3 create successfully |
| 2 | Verify state | `parallel_1.tag.# = 0`, `parallel_2.tag.# = 0`, `parallel_3.tag.# = 1` |
| 3 | PlanOnly (idempotency) | Shows "No changes" for all 3 |

---

## Execution Summary

| Plan | Type | Tests | Implemented? | Priority |
|------|------|-------|-------------|----------|
| 1  | Unit test | 10 | ✅ Yes (`go test ./internal/utils/ -v`) | **P0 — All passing** |
| 2  | Acceptance | 4 | ✅ Yes (`address_nulllist_test.go`) | **P0 — Core bug fix** |
| 3  | Acceptance | 4 | ✅ Yes (`security_rule_nulllist_test.go`, `address_group_service_nulllist_test.go`) | P1 |
| 4  | Acceptance | 0 | ⏳ Covered by unit test #8 | P1 |
| 5  | Acceptance | 3 | ✅ Yes (`address_nulllist_test.go`, `address_group_service_nulllist_test.go`) | P1 |
| 6  | Acceptance | 4 | ✅ Partial (4 of 7 categories in `cross_category_nulllist_test.go`) | P2 |
| 7  | Acceptance | 2 | ✅ Yes (`address_nulllist_test.go`, `security_rule_nulllist_test.go`) | P1 |
| 8  | Acceptance | 0 | ⏳ Manual validation | P2 |
| 9  | Acceptance | 0 | ⏳ Manual validation | P2 |
| 10 | Acceptance | 1 | ✅ Yes (`cross_category_nulllist_test.go`) | P2 |

### Test Count Summary

| Type | Count | Status |
|------|-------|--------|
| Unit tests (`normalize_test.go`) | 10 | ✅ All implemented |
| Acceptance tests (4 files) | 15 | ✅ All implemented |
| **Total automated tests** | **25** | ✅ |

### How to Run

```bash
# Unit tests only (no API required)
go test ./internal/utils/ -v

# All acceptance tests (requires SCM_CLIENT_ID, SCM_CLIENT_SECRET, SCM_SCOPE env vars)
go test ./internal/provider/ -v -run "NullList|EmptyTag|TagOmitted|PopulatedTag|Lifecycle|ParallelCreate"

# Specific test plans
go test ./internal/provider/ -v -run "TestAccAddress_EmptyTag"          # Plan 2
go test ./internal/provider/ -v -run "TestAccSecurityRule"              # Plan 3A + 7B
go test ./internal/provider/ -v -run "TestAccAddressGroup|TestAccService" # Plan 3B + 5B
go test ./internal/provider/ -v -run "TestAccZone|TestAccAuth|TestAccApp|TestAccUrl" # Plan 6
go test ./internal/provider/ -v -run "TestAccAddress_Lifecycle"         # Plan 7A
go test ./internal/provider/ -v -run "TestAccAddress_ParallelCreate"    # Plan 10
```

**Minimum for PR validation:** Plans 1 + 2 (unit tests + original bug reproduction)
**Recommended for full confidence:** Plans 1–5, 7
**Complete regression:** All 10 plans
