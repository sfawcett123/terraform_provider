terraform {
  required_providers {
    fawcetts = {
      version = "0.2"
      source  = "sfawcett123.github.io/github/fawcetts"
    }
  }
}

data "fawcetts_repositories" "single" {
  owner = "sfawcett123"
  name = "terraform_provider"
}

output "repos_single" {
  value = data.fawcetts_repositories.single.repositories
}

data "fawcetts_repositories" "all" {
  owner = "sfawcett123"
}

output "repos" {
  value = data.fawcetts_repositories.all.repositories
}
