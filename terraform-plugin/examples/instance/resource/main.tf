# cd ../../.. && make install && cd examples/instance/resource
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
  username = "education"
  password = "test123"
  host     = "http://localhost:8090"
}


# resource "dummycloud_instance" "created" {
#   id = "d581956a-823e-4464-97f6-ce6ab9a364e3"
#   name = "boomboom_fast"
#   size = "33gb"
#   region = "dhule"
#   ram = "3gb"
#   os = "win"
# }
