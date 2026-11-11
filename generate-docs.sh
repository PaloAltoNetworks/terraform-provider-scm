#!/usr/bin/env bash
# Generates the provider docs. Every type is registered as scm_<...>, so a single
# tfplugindocs run with --provider-name scm is enough. The "docsgen" build tag
# leaves out the deprecated unprefixed names kept for existing users (they would
# produce pages named like their scm_ counterparts).
set -euo pipefail

export GOFLAGS="${GOFLAGS:-} -tags=docsgen"

go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs generate --provider-name scm

echo "Documentation generation complete."
