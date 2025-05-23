# Terraform Provider Swiss Army Knife (SAK)

This repository contains the [Terraform](https://www.terraform.io) provider "Swiss Army Knife" (SAK). This provider aims to enhance the Terraform development experience by offering a collection of utilities and helpers. Currently it's a function-only provider.

## Requirements

- [Terraform](https://developer.hashicorp.com/terraform/downloads) >= 1.0
- _NB! Terraform provider functions are only available in Terraform >= 1.8_
- [Go](https://golang.org/doc/install) >= 1.24

## Building The Provider

1. Clone the repository
1. Enter the repository directory
1. Build the provider using the Go `install` command:

```shell
go install
```

## Adding Dependencies

This provider uses [Go modules](https://github.com/golang/go/wiki/Modules).
Please see the Go documentation for the most up to date information about using Go modules.

To add a new dependency `github.com/author/dependency` to your Terraform provider:

```shell
go get github.com/author/dependency
go mod tidy
```

Then commit the changes to `go.mod` and `go.sum`.

## Using the provider

This provider offers a collection of utilities to enhance your Terraform workflow.

Documentation for each function can be found in the `docs/` directory and on the [Terraform Registry page](https_link_to_registry_page_once_published).

Example usage:

```terraform
terraform {
  required_providers {
    sak = {
      source  = "serpro69/sak"
      version = "0.1.0"
    }
  }
}

provider "sak" {
  # The provider currently does not have any configuration options
}

output "timestamp" {
  value = provider::sak::rfc3339_parse("2023-07-25T23:43:16Z")
}
```

## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (see [Requirements](#requirements) above).

To compile the provider, run `go install`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

Set `dev_overrides` for the provider installation in `~/.terraformrc`:

```hcl
provider_installation {

  dev_overrides {
      "serpro69/sak" = "/Users/sergio/go/bin"
  }

  # For all other providers, install them directly from their origin provider
  # registries as normal. If you omit this, Terraform will _only_ use
  # the dev_overrides block, and so no other providers will be available.
  direct {}
}
```

To generate or update documentation, run `make generate`.

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```shell
make testacc
```
