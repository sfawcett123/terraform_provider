terraform {
  required_providers {
    fawcetts = {
      version = "0.2"
      source  = "sfawcett123.github.io/github/fawcetts"
    }
  }
}

data "fawcetts" "all" {}

output "repos" {
  value = data.fawcetts.all.items
}

