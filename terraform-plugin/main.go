package main

import (
	"context"
	dummycloud "dummy-cloud/dummycloud"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

// Provider documentation generation.

func main() {
	providerserver.Serve(context.Background(), dummycloud.New, providerserver.ServeOpts{
		// NOTE: This is not a typical Terraform Registry provider address,
		// such as registry.terraform.io/anshumanpatil.com/dummy-cloud. This specific
		// provider address is used in these tutorials in conjunction with a
		// specific Terraform CLI configuration for manual development testing
		// of this provider.
		Address: "anshumanpatil.com/dev/dummycloud",
	})
}
