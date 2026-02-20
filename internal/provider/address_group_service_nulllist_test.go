package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// =============================================================================
// Test Plan 3B: Address Group — static list + tag = []
// =============================================================================

func TestAccAddressGroup_EmptyTag(t *testing.T) {
	rName := fmt.Sprintf("tf-test-%s", acctest.RandString(6))
	addrName := fmt.Sprintf("tf-test-addr-%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAddressGroupConfig_emptyTag(addrName, rName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("scm_address_group.test", "name", rName),
					resource.TestCheckResourceAttr("scm_address_group.test", "tag.#", "0"),
					resource.TestCheckResourceAttr("scm_address_group.test", "static.#", "1"),
				),
			},
			// Idempotency
			{
				Config:   testAccAddressGroupConfig_emptyTag(addrName, rName),
				PlanOnly: true,
			},
		},
	})
}

// =============================================================================
// Test Plan 5B: Service — tag omitted
// =============================================================================

func TestAccService_TagOmitted(t *testing.T) {
	rName := fmt.Sprintf("tf-test-%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccServiceConfig_noTag(rName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("scm_service.test", "name", rName),
					resource.TestCheckNoResourceAttr("scm_service.test", "tag.#"),
				),
			},
			// Idempotency
			{
				Config:   testAccServiceConfig_noTag(rName),
				PlanOnly: true,
			},
		},
	})
}

func TestAccService_EmptyTag(t *testing.T) {
	rName := fmt.Sprintf("tf-test-%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccServiceConfig_emptyTag(rName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("scm_service.test", "name", rName),
					resource.TestCheckResourceAttr("scm_service.test", "tag.#", "0"),
				),
			},
			// Idempotency
			{
				Config:   testAccServiceConfig_emptyTag(rName),
				PlanOnly: true,
			},
		},
	})
}

// =============================================================================
// Config Helpers — Address Group
// =============================================================================

func testAccAddressGroupConfig_emptyTag(addrName, groupName string) string {
	return fmt.Sprintf(`
resource "scm_address" "dep" {
  folder     = "All"
  name       = %q
  ip_netmask = "192.168.1.0/24"
}

resource "scm_address_group" "test" {
  folder = "All"
  name   = %q
  tag    = []
  static = [scm_address.dep.name]
}
`, addrName, groupName)
}

// =============================================================================
// Config Helpers — Service
// =============================================================================

func testAccServiceConfig_noTag(name string) string {
	return fmt.Sprintf(`
resource "scm_service" "test" {
  folder = "All"
  name   = %q
  protocol = {
    tcp = {
      port = "8080"
    }
  }
}
`, name)
}

func testAccServiceConfig_emptyTag(name string) string {
	return fmt.Sprintf(`
resource "scm_service" "test" {
  folder = "All"
  name   = %q
  tag    = []
  protocol = {
    tcp = {
      port = "8080"
    }
  }
}
`, name)
}
