package provider_test

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	provider "github.com/paloaltonetworks/terraform-provider-scm/internal/provider"
)

// testAccProtoV6ProviderFactories is used to instantiate a provider during
// acceptance testing. The factory function is called for every Terraform CLI
// command executed to create a provider server to which the CLI can make calls.
var testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"scm": providerserver.NewProtocol6WithError(provider.New("test")),
}

// testAccPreCheck validates the necessary environment variables exist for
// running acceptance tests. It is called before every test function.
func testAccPreCheck(t *testing.T) {
	t.Helper()

	required := []string{"SCM_CLIENT_ID", "SCM_CLIENT_SECRET", "SCM_SCOPE"}
	for _, envVar := range required {
		if v := os.Getenv(envVar); v == "" {
			t.Fatalf("environment variable %s must be set for acceptance tests", envVar)
		}
	}
}
