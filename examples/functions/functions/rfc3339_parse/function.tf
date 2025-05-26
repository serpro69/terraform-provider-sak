terraform {
  required_providers {
    sak = {
      source = "serpro69/sak"
    }
  }
}

provider "sak" {}

output "timestamp" {
  value = provider::sak::rfc3339_parse("2023-07-25T23:43:16Z")
}
