# cd ../../.. && make install && cd examples/instance/datasource
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


data "dummycloud_instance" "all" {}

# Returns all instances
output "all_instance_available" {
  value = data.dummycloud_instance.all.instances
}