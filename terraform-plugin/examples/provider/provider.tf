# cd ../.. && make install && cd examples/provider
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
