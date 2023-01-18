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
  username = "admin"
  password = "admin"
  host     = "http://localhost:8090"
}


# resource "dummycloud_instance" "created" {
#   id = "2bba74f6-de4a-4d22-860e-ab3619546785"
#   name = "boomboom_slow"
#   size = "33gb"
#   region = "dhule"
#   ram = "3gb"
#   os = "win"
# }
