module "repositories" {
  source = "./github"
}

output "repositories" {
  value = module.repositories
}
