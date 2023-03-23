terraform {
  required_providers {
    fawcetts = {
      version = "0.2"
      source  = "fawcetts.com/TEST/fawcetts"
    }
  }
}

data "fawcetts" "all" {}

output "repos" {
  value = data.fawcetts.all.items
}

