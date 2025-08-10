terraform {
  required_providers {
    hashicups = {
      source = "bakwas-cloud"
    }
  }
}

provider "hashicups" {}

data "hashicups_coffees" "example" {}
