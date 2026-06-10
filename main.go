package main

import (
	"context"
	"flag"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"

	"github.com/paloaltonetworks/terraform-provider-scm/internal/provider"
)

// Run generate-docs.sh to format examples and generate per-prefix documentation.
// The script runs tfplugindocs once per resource prefix (scm_, ztna_, etc.) so
// each prefix gets correct, full docs without the tool crashing on cross-prefix
// resources.  See generate-docs.sh for details.
//go:generate bash generate-docs.sh

var (
	// these will be set by the goreleaser configuration
	// to appropriate values for the compiled binary.
	version string = "dev"

	// goreleaser can pass other information to the main package, such as the specific commit
	// https://goreleaser.com/cookbooks/using-main.version/
)

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := providerserver.ServeOpts{
		Address: "registry.terraform.io/paloaltonetworks/scm",
		Debug:   debug,
	}

	err := providerserver.Serve(context.Background(), provider.NewFactory(version), opts)
	if err != nil {
		log.Fatal(err.Error())
	}
}
