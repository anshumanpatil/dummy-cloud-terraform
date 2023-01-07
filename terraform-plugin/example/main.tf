# https://github.com/hashicorp/terraform-provider-hashicups/blob/main/hashicups/provider.go
# https://developer.hashicorp.com/terraform/tutorials/providers/provider-debug
# https://developer.hashicorp.com/terraform/tutorials/providers/provider-complex-read
# https://github.com/hashicorp/terraform-provider-hashicups/blob/main/examples/main.tf

# cd .. && make install && cd example
# terraform init && terraform apply --auto-approve
terraform {
  required_providers {
    pfxm = {
      version = "0.2"
      source  = "dell.com/dev/pfxm"
    }
  }
}

provider "pfxm" {
    username = "admin"
    password = "Scaleio!"
    host = "http://localhost:3000"
}

# Read all sdcs if id is blank, otherwise reads all sdcs
# data "pfxm_sdcs" "allsdc" {
#     sdcid = "11a0df86-0c6c-4032-845c-89e20bf4070a"
# }

# # Returns all sdcs
# output "allsdcresult" {
#   value = data.pfxm_sdcs.allsdc.sdcs
# }

resource "pfxm_sdcs" "edu" {
  sdcid = "11a0df86-0c6c-4032-845c-89e20bf4070a"
  name = "ionverynew"
}

output "edu_order" {
  value = pfxm_sdcs.edu
}