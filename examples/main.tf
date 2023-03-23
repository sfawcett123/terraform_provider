terraform {
  required_providers {
    fawcetts = {
      version = "0.2"
      source  = "fawcetts.com/TEST/fawcetts"
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
