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
