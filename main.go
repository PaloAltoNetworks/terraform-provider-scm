package main

import (
	"context"
	"flag"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"

	"github.com/paloaltonetworks/terraform-provider-scm/internal/provider"
)

// Run "go generate" to format example terraform files and generate the docs for the registry/website

// If you do not have terraform installed, you can remove the formatting command, but it's suggested to
// ensure the documentation is formatted properly.
//go:generate terraform fmt -recursive ./examples/

// Run the docs generation tool. Because this provider registers resources under two prefixes
// (scm and ztna), generate-docs.sh runs tfplugindocs once per prefix so each resource gets
// the correct page title. See generate-docs.sh for details.
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
