# Terraform Provider Fawcetts

## Purpose
This provider is a simple prototype that connects to an REST API and initially gets some data.

### Further Reading

CRUD - Create **Read** Update Delete 

[Perform CRUD Operations with Providers](https://developer.hashicorp.com/terraform/tutorials/providers/provider-use)

 
## Provider Delivery

### Build
Run the following command to build the provider

```shell
go build -o terraform-provider-fawcetts
```

## Install

First, build and install the provider.

```shell
make install
```

## Terraform Usage

Then, run the following command to initialize the workspace and apply the sample configuration.

```shell
cd examples
terraform init && terraform apply
```
### CRUD

#### Read

```yaml
terraform {
  required_providers {
    fawcetts = {
      version = "0.2"
      source  = "sfawcett123.github.io/TEST/fawcetts"
    }
  }
}

data "fawcetts" "all" {}

output "repos" {
  value = data.fawcetts.all.items
}```


