package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// =============================================================================
// Test Plan 2: Original Bug — tag = [] (Issue #92)
// =============================================================================

// Test 2A: Create an address with tag = [] — the exact bug from issue #92.
// Before the fix, this would fail with:
//
//	"Provider produced inconsistent result after apply"
//	tag: was null, but was cty.ListValEmpty(cty.String)
func TestAccAddress_EmptyTag_Create(t *testing.T) {
	rName := fmt.Sprintf("tf-test-%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create with tag = []
			{
				Config: testAccAddressConfig_emptyTag(rName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("scm_address.test", "name", rName),
					resource.TestCheckResourceAttr("scm_address.test", "ip_netmask", "10.0.0.0/24"),
					resource.TestCheckResourceAttr("scm_address.test", "tag.#", "0"),
				),
			},
			// Step 2: Plan should show no changes (idempotency)
			{
				Config:   testAccAddressConfig_emptyTag(rName),
				PlanOnly: true,
			},
		},
	})
}

// Test 2B: Update from populated tag to tag = []
func TestAccAddress_EmptyTag_Update(t *testing.T) {
	rName := fmt.Sprintf("tf-test-%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create with populated tag
			{
				Config: testAccAddressConfig_withTag(rName, `["test-tag"]`),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("scm_address.test", "tag.#", "1"),
					resource.TestCheckResourceAttr("scm_address.test", "tag.0", "test-tag"),
				),
			},
			// Step 2: Update to tag = []
			{
				Config: testAccAddressConfig_emptyTag(rName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("scm_address.test", "tag.#", "0"),
				),
			},
			// Step 3: Plan should show no changes
			{
				Config:   testAccAddressConfig_emptyTag(rName),
				PlanOnly: true,
			},
		},
	})
}

// Test 2C: Tag omitted entirely (should remain null — no regression)
func TestAccAddress_TagOmitted(t *testing.T) {
	rName := fmt.Sprintf("tf-test-%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAddressConfig_noTag(rName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("scm_address.test", "name", rName),
					resource.TestCheckResourceAttr("scm_address.test", "ip_netmask", "10.0.0.0/24"),
					resource.TestCheckNoResourceAttr("scm_address.test", "tag.#"),
				),
			},
			// Idempotency
			{
				Config:   testAccAddressConfig_noTag(rName),
				PlanOnly: true,
			},
		},
	})
}

// =============================================================================
// Test Plan 5: Negative — Populated tags remain populated
// =============================================================================

func TestAccAddress_PopulatedTag(t *testing.T) {
	rName := fmt.Sprintf("tf-test-%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAddressConfig_withTag(rName, `["web-server", "production"]`),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("scm_address.test", "tag.#", "2"),
					resource.TestCheckResourceAttr("scm_address.test", "tag.0", "web-server"),
					resource.TestCheckResourceAttr("scm_address.test", "tag.1", "production"),
				),
			},
			// Idempotency
			{
				Config:   testAccAddressConfig_withTag(rName, `["web-server", "production"]`),
				PlanOnly: true,
			},
		},
	})
}

// =============================================================================
// Test Plan 7: Full CRUD Lifecycle
// =============================================================================

func TestAccAddress_Lifecycle(t *testing.T) {
	rName := fmt.Sprintf("tf-test-%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Phase 1: Create with tag = []
			{
				Config: testAccAddressConfig_emptyTag(rName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("scm_address.test", "tag.#", "0"),
				),
			},
			// Phase 2: Update — add tags
			{
				Config: testAccAddressConfig_withTag(rName, `["web"]`),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("scm_address.test", "tag.#", "1"),
					resource.TestCheckResourceAttr("scm_address.test", "tag.0", "web"),
				),
			},
			// Phase 3: Update — back to empty tags
			{
				Config: testAccAddressConfig_emptyTag(rName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("scm_address.test", "tag.#", "0"),
				),
			},
			// Phase 4: Update — change non-tag field, tag stays []
			{
				Config: testAccAddressConfig_emptyTagWithDescription(rName, "updated-description"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("scm_address.test", "tag.#", "0"),
					resource.TestCheckResourceAttr("scm_address.test", "description", "updated-description"),
				),
			},
			// Phase 5: Idempotency
			{
				Config:   testAccAddressConfig_emptyTagWithDescription(rName, "updated-description"),
				PlanOnly: true,
			},
		},
	})
}

// =============================================================================
// Config Helpers — Address
// =============================================================================

func testAccAddressConfig_emptyTag(name string) string {
	return fmt.Sprintf(`
resource "scm_address" "test" {
  folder     = "All"
  name       = %q
  ip_netmask = "10.0.0.0/24"
  tag        = []
}
`, name)
}

func testAccAddressConfig_emptyTagWithDescription(name, description string) string {
	return fmt.Sprintf(`
resource "scm_address" "test" {
  folder      = "All"
  name        = %q
  ip_netmask  = "10.0.0.0/24"
  description = %q
  tag         = []
}
`, name, description)
}

func testAccAddressConfig_withTag(name, tagValue string) string {
	return fmt.Sprintf(`
resource "scm_address" "test" {
  folder     = "All"
  name       = %q
  ip_netmask = "10.0.0.0/24"
  tag        = %s
}
`, name, tagValue)
}

func testAccAddressConfig_noTag(name string) string {
	return fmt.Sprintf(`
resource "scm_address" "test" {
  folder     = "All"
  name       = %q
  ip_netmask = "10.0.0.0/24"
}
`, name)
}
