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
#   id = "324cfef8-1450-4921-a7a6-90e30f34919b"
#   name = "boomboom_fast"
#   isactive = true
#   iplist = ["33gb"]
#   instancelist = [
#     {
#       name = "x"
#       region = "y"
#     }
#   ]
# }
