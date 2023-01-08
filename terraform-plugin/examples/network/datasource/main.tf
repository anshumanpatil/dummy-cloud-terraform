# cd ../../.. && make install && cd examples/network/datasource
# terraform init && terraform apply --auto-approve

terraform {
  required_providers {
    dummycloud = {
      version = "0.1"
      source  = "anshumanpatil.com/dev/dummycloud"
    }
  }
}

# Configuration-based authentication
provider "dummycloud" {
  username = "admin"
  password = "admin"
  host     = "http://localhost:8090"
}


data "dummycloud_network" "all" {}

# Returns all networks
output "all_network_available" {
  value = data.dummycloud_network.all.networks
}