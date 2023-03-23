terraform {
  required_providers {
    fawcetts = {
      version = "0.2"
      source  = "sfawcett123.github.io/github/fawcetts"
    }
  }
}

data "fawcetts" "single" {
  owner = "sfawcett123"
  name = "terraform_provider"
}

output "repos_single" {
  value = data.fawcetts.single.repositories
}

data "fawcetts" "all" {
  owner = "sfawcett123"
}

output "repos" {
  value = data.fawcetts.all.repositories
}
