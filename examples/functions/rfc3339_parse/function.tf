terraform {
  required_providers {
    exampletime = {
      source = "hashicorp.com/edu/exampletime"
    }
  }
}

provider "exampletime" {}

output "timestamp" {
  value = provider::exampletime::rfc3339_parse("2023-07-25T23:43:16Z")
}
