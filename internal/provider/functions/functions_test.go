// Copyright (c) serpro69
// SPDX-License-Identifier: MIT

package functions_test

import (
	"github.com/serpro69/terraform-provider-sak/internal/provider"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

// testAccProtoV6ProviderFactories is used to instantiate a provider during acceptance testing.
// The factory function is called for each Terraform CLI command to create a provider
// server that the CLI can connect to and interact with.
var testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"sak": providerserver.NewProtocol6WithError(provider.New("test")()),
}
