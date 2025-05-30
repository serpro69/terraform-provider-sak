---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "yaml_decode function - sak"
subcategory: ""
description: |-
  Decode a YAML file containing one or multiple documents
---

# function: yaml_decode

Given a YAML text file containing one or multiple documents, will decode the file and return a tuple of object representations for each document.

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

locals {
  yaml_single = <<-YAML
  ---
  key1: value1
  key2: value2
  nested:
    key3: value3
  YAML

  yaml_multi = <<-YAML
  ---
  key1: value1
  key2: value2
  nested:
    key3: value3
  ---
  key4: value4
  key5: value5
  nested:
    key6: value6
  ---
  YAML
}

output "yaml_single" {
  value = provider::sak::yaml_decode(local.yaml_single)
}

output "yaml_multi" {
  value = provider::sak::yaml_decode(local.yaml_multi)
}
```

## Signature

<!-- signature generated by tfplugindocs -->
```text
yaml_decode(document string) dynamic
```

## Arguments

<!-- arguments generated by tfplugindocs -->
1. `document` (String) The YAML plaintext for a document

