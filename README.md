# Dummy Cloud - Web Console, Rest Client, Terraform

[![Build Status](https://travis-ci.org/joemccann/dillinger.svg?branch=master)](https://travis-ci.org/joemccann/dillinger)

Create,Read,Update & Delete resorces and view on web console.

- Postman Support for Rest API's.
- Real-Time Web Console for resources.
- Custom Terraform Plugin for resource operations.


## Installation

Dummy Cloud requires [Golang](https://go.dev/) to compile terraform plugin locally.

Dummy Cloud requires [Golang](https://go.dev/) to run rest client locally.

Dummy Cloud requires [NodeJS](https://nodejs.org/en/) to run web console locally.

# OR

Dummy Cloud requires [Golang](https://go.dev/) to compile terraform plugin locally.

Dummy Cloud requires [Docker](https://www.docker.com/) to run on docker(Web console & Rest API) .




Run Web Console on local machine(React).

```sh
git clone https://github.com/anshumanpatil/dummy-cloud-terraform.git
cd dummy-cloud-terraform
cd cloud-console
npm i --verbose
npm start
```
Run Rest API backend on local machine(Golang-Gin).

```sh
git clone https://github.com/anshumanpatil/dummy-cloud-terraform.git
cd dummy-cloud-terraform
cd rest-client
go mod tidy
npm i --verbose
npm start 
        OR
go run .
```

Run Rest API & Web Console both on Docker.

```sh
git clone https://github.com/anshumanpatil/dummy-cloud-terraform.git
cd dummy-cloud-terraform
docker-compose up
```


Compile and Install Terraform Plugin.

```sh
git clone https://github.com/anshumanpatil/dummy-cloud-terraform.git
cd dummy-cloud-terraform
cd terraform-plugin
go mod tidy
make install
```

Run Terraform Plugin.

```sh
git clone https://github.com/anshumanpatil/dummy-cloud-terraform.git
cd dummy-cloud-terraform
cd terraform-plugin

cd examples/bucket/datasource
terraform init && terraform apply --auto-approve
cd examples/bucket/resource
terraform init && terraform apply --auto-approve

cd examples/instance/datasource
terraform init && terraform apply --auto-approve
cd examples/instance/resource
terraform init && terraform apply --auto-approve

cd examples/network/datasource
terraform init && terraform apply --auto-approve
cd examples/network/resource
terraform init && terraform apply --auto-approve

cd examples/network/datasource
terraform init && terraform apply --auto-approve
cd examples/network/resource
terraform init && terraform apply --auto-approve
```

## License

MIT
