---
page_title: "Provider: SAK (Swiss Army Knife)"
subcategory: "DevEx"
description: |-
  The SAK provider can be used to improve development experience when working with terraform and provides useful extensions for managing your infrastructure with terraform.
---

# SAK Provider

This SAK (Swiss Army Knife) provider aims to enhance the Terraform development experience by offering a collection of utilities and helpers. 
Currently it is a function-only provider and does not contain any managed resources or data sources.

## Example Usage

```terraform
terraform {
  required_providers {
    sak = {
      source = "serpro69/sak"
    }
  }
}

provider "sak" {}
```
