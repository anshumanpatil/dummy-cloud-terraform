# cd ../../.. && make install && cd examples/network/resource
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


# resource "dummycloud_network" "created" {
#   id = "b22f4546-4762-4cc8-891a-9da9a3ee049b"
#   name = "boomboom_slow"
#   size = "33gb"
#   region = "dhule"
# }
