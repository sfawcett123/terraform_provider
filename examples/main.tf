terraform {
  required_providers {
    fawcetts = {
      version = "0.2"
      source  = "sfawcett123.github.io/github/fawcetts"
    }
  }
}

provider "fawcetts" {}

module "psl" {
  source = "./github"
}

output "psl" {
  value = module.psl
}
