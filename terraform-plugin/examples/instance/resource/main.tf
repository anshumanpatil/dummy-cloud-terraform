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


resource "dummycloud_instance" "created" {
  # id = "63490a53-3457-4ee9-bb68-c720fa97cd8d"
  name = "boomboom_poo"
  size = "33gb"
  region = "dhule"
  ram = "3gb"
  os = "win"
}
