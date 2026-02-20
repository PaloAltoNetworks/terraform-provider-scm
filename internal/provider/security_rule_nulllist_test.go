package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// =============================================================================
// Test Plan 3A: Security Rule — Multiple list fields with tag = []
// =============================================================================

func TestAccSecurityRule_EmptyTagAndLists(t *testing.T) {
	rName := fmt.Sprintf("tf-test-%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSecurityRuleConfig_emptyLists(rName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("scm_security_rule.test", "name", rName),
					resource.TestCheckResourceAttr("scm_security_rule.test", "action", "allow"),
					resource.TestCheckResourceAttr("scm_security_rule.test", "from.#", "1"),
					resource.TestCheckResourceAttr("scm_security_rule.test", "to.#", "1"),
					resource.TestCheckResourceAttr("scm_security_rule.test", "source.#", "1"),
					resource.TestCheckResourceAttr("scm_security_rule.test", "destination.#", "1"),
					resource.TestCheckResourceAttr("scm_security_rule.test", "application.#", "1"),
					resource.TestCheckResourceAttr("scm_security_rule.test", "service.#", "1"),
					resource.TestCheckResourceAttr("scm_security_rule.test", "tag.#", "0"),
					resource.TestCheckResourceAttr("scm_security_rule.test", "category.#", "0"),
				),
			},
			// Idempotency
			{
				Config:   testAccSecurityRuleConfig_emptyLists(rName),
				PlanOnly: true,
			},
		},
	})
}

// Security rule with populated tags to verify no regression
func TestAccSecurityRule_PopulatedTag(t *testing.T) {
	rName := fmt.Sprintf("tf-test-%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSecurityRuleConfig_withTag(rName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("scm_security_rule.test", "tag.#", "1"),
					resource.TestCheckResourceAttr("scm_security_rule.test", "tag.0", "test-tag"),
				),
			},
			// Idempotency
			{
				Config:   testAccSecurityRuleConfig_withTag(rName),
				PlanOnly: true,
			},
		},
	})
}

// Security rule lifecycle: create → add tag → remove tag → destroy
func TestAccSecurityRule_TagLifecycle(t *testing.T) {
	rName := fmt.Sprintf("tf-test-%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with empty tag
			{
				Config: testAccSecurityRuleConfig_emptyLists(rName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("scm_security_rule.test", "tag.#", "0"),
				),
			},
			// Add tag
			{
				Config: testAccSecurityRuleConfig_withTag(rName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("scm_security_rule.test", "tag.#", "1"),
				),
			},
			// Remove tag
			{
				Config: testAccSecurityRuleConfig_emptyLists(rName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("scm_security_rule.test", "tag.#", "0"),
				),
			},
		},
	})
}

// =============================================================================
// Config Helpers — Security Rule
// =============================================================================

func testAccSecurityRuleConfig_emptyLists(name string) string {
	return fmt.Sprintf(`
resource "scm_security_rule" "test" {
  folder      = "All"
  position    = "pre"
  name        = %q
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
`, name)
}

func testAccSecurityRuleConfig_withTag(name string) string {
	return fmt.Sprintf(`
resource "scm_security_rule" "test" {
  folder      = "All"
  position    = "pre"
  name        = %q
  from        = ["any"]
  to          = ["any"]
  source      = ["any"]
  destination = ["any"]
  application = ["any"]
  service     = ["any"]
  action      = "allow"
  tag         = ["test-tag"]
}
`, name)
}
