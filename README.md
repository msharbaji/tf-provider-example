# tf-provider-example

## Overview
This repository provides an example of a Terraform provider using [terraform-plugin-framework](https://github.com/hashicorp/terraform-plugin-framework). The framework is designed for constructing Terraform providers.

The provider in this example utilizes a gRPC client to establish communication with a gRPC server implemented in  [grpc-go-example](https://github.com/msharbaji/grpc-go-example).

## Usage

### Build the provider
```shell
go build -o terraform-provider-example
```

### Run the example
In this example we will move the compiled provider binary to corresponding directory based on the operated system and architecture. This allows terraform to locate and use 
provider when executing the example.

In this example the provider `malsharbaji.com/providers/example`

```shell
os_arch=$(go env GOOS)_$(go env GOARCH)
mkdir -p "$HOME/.terraform.d/plugins/malsharbaji.com/providers/example/1.0.0/${os_arch}/"
mv terraform-provider-example "$HOME/.terraform.d/plugins/malsharbaji.com/providers/example/1.0.0/${os_arch}/"

export TF_VAR_key_id="<my-key-id>"
export TF_VAR_secret_key="<my-secret-key>"

cd examples

# You change the directory to run the example for data source
cd examples/resources/example_user

rm -rf .terraform .terraform.lock.hcl plan.out terraform.tfstate terraform.tfstate.backup

export TF_LOG_PROVIDER=INFO

terraform init

terraform validate

terraform plan -out=plan.out

terraform apply plan.out

cd ../../..
```

### Test the provider
```shell
TF_ACC=1 go test -v  ./internal/provider/
```

### Provider Usage Documentation
you can check it under docs, It uses the [terraform-plugin-docs](https://github.com/hashicorp/terraform-plugin-docs) to generate the documentation for the provider.

```shell
go generate ./...
```

### References
- [terraform-plugin-framework](https://github.com/hashicorp/terraform-plugin-framework)
- [terraform-plugin-docs](https://github.com/hashicorp/terraform-plugin-docs)
- [grpc-go-example](https://github.com/msharbaji/grpc-go-example)






