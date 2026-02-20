package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// =============================================================================
// Test Plan 6: Cross-Category Resource Coverage (Smoke Tests)
// One resource from each provider category to confirm normalization works.
// =============================================================================

// 6E: network_services — Zone with empty tag
func TestAccZone_EmptyTag(t *testing.T) {
	rName := fmt.Sprintf("tf-test-%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccZoneConfig_emptyTag(rName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("scm_zone.test", "name", rName),
				),
			},
			// Idempotency
			{
				Config:   testAccZoneConfig_emptyTag(rName),
				PlanOnly: true,
			},
		},
	})
}

// 6D: identity_services — Authentication Profile with empty tag
func TestAccAuthenticationProfile_EmptyTag(t *testing.T) {
	rName := fmt.Sprintf("tf-test-%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationProfileConfig_emptyTag(rName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("scm_authentication_profile.test", "name", rName),
				),
			},
			// Idempotency
			{
				Config:   testAccAuthenticationProfileConfig_emptyTag(rName),
				PlanOnly: true,
			},
		},
	})
}

// 6F: objects — Application Group with members + empty tag
func TestAccApplicationGroup_EmptyTag(t *testing.T) {
	rName := fmt.Sprintf("tf-test-%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccApplicationGroupConfig_emptyTag(rName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("scm_application_group.test", "name", rName),
					resource.TestCheckResourceAttr("scm_application_group.test", "tag.#", "0"),
				),
			},
			// Idempotency
			{
				Config:   testAccApplicationGroupConfig_emptyTag(rName),
				PlanOnly: true,
			},
		},
	})
}

// 6G: security_services — URL Category with empty tag
func TestAccUrlCategory_EmptyTag(t *testing.T) {
	rName := fmt.Sprintf("tf-test-%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUrlCategoryConfig_emptyTag(rName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("scm_url_category.test", "name", rName),
					resource.TestCheckResourceAttr("scm_url_category.test", "tag.#", "0"),
				),
			},
			// Idempotency
			{
				Config:   testAccUrlCategoryConfig_emptyTag(rName),
				PlanOnly: true,
			},
		},
	})
}

// =============================================================================
// Test Plan 10: Parallel Resource Creation
// =============================================================================

func TestAccAddress_ParallelCreate_EmptyTag(t *testing.T) {
	rName1 := fmt.Sprintf("tf-test-%s", acctest.RandString(6))
	rName2 := fmt.Sprintf("tf-test-%s", acctest.RandString(6))
	rName3 := fmt.Sprintf("tf-test-%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAddressConfig_parallel(rName1, rName2, rName3),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("scm_address.parallel_1", "tag.#", "0"),
					resource.TestCheckResourceAttr("scm_address.parallel_2", "tag.#", "0"),
					resource.TestCheckResourceAttr("scm_address.parallel_3", "tag.#", "1"),
					resource.TestCheckResourceAttr("scm_address.parallel_3", "tag.0", "shared"),
				),
			},
			// Idempotency
			{
				Config:   testAccAddressConfig_parallel(rName1, rName2, rName3),
				PlanOnly: true,
			},
		},
	})
}

// =============================================================================
// Config Helpers
// =============================================================================

func testAccZoneConfig_emptyTag(name string) string {
	return fmt.Sprintf(`
resource "scm_zone" "test" {
  folder = "All"
  name   = %q
  network = {
    layer3 = []
  }
}
`, name)
}

func testAccAuthenticationProfileConfig_emptyTag(name string) string {
	return fmt.Sprintf(`
resource "scm_authentication_profile" "test" {
  folder = "All"
  name   = %q
  method = {}
}
`, name)
}

func testAccApplicationGroupConfig_emptyTag(name string) string {
	return fmt.Sprintf(`
resource "scm_application_group" "test" {
  folder  = "All"
  name    = %q
  members = ["ssl"]
  tag     = []
}
`, name)
}

func testAccUrlCategoryConfig_emptyTag(name string) string {
	return fmt.Sprintf(`
resource "scm_url_category" "test" {
  folder = "All"
  name   = %q
  type   = "URL List"
  list   = ["example.com"]
  tag    = []
}
`, name)
}

func testAccAddressConfig_parallel(name1, name2, name3 string) string {
	return fmt.Sprintf(`
resource "scm_address" "parallel_1" {
  folder     = "All"
  name       = %q
  ip_netmask = "10.0.0.1/32"
  tag        = []
}

resource "scm_address" "parallel_2" {
  folder     = "All"
  name       = %q
  ip_netmask = "10.0.0.2/32"
  tag        = []
}

resource "scm_address" "parallel_3" {
  folder     = "All"
  name       = %q
  ip_netmask = "10.0.0.3/32"
  tag        = ["shared"]
}
`, name1, name2, name3)
}
