#!/usr/bin/env bash
set -eo pipefail

echo "Building and installing example provider"

cd cmd

go build -o terraform-provider-example

os_arch=$(go env GOOS)_$(go env GOARCH)

mkdir -p "$HOME/.terraform.d/plugins/malsharbaji.com/providers/example/1.0.0/${os_arch}/"

mv terraform-provider-example "$HOME/.terraform.d/plugins/malsharbaji.com/providers/example/1.0.0/${os_arch}/"

cd ..

export TF_VAR_key_id="my-key-id"

export TF_VAR_secret_key="my-secret-key"

#cd examples/provider
#
#echo "Running terraform"
#
#
#rm -rf .terraform .terraform.lock.hcl plan.out terraform.tfstate terraform.tfstate.backup
#
#export TF_LOG_PROVIDER=INFO
#
#echo "terraform init"
#terraform init
#
#echo "terraform valid"
#terraform validate
#
#echo "terraform plan"
#terraform plan -out=plan.out
#
#echo "terraform apply"
#terraform apply plan.out
#
#cd ../..
#
#cd examples/data-sources/example_user
#
#rm -rf .terraform .terraform.lock.hcl plan.out terraform.tfstate terraform.tfstate.backup
#
#export TF_LOG_PROVIDER=INFO
#
#echo "terraform init"
#terraform init
#
#echo "terraform valid"
#terraform validate
#
#echo "terraform plan"
#terraform plan -out=plan.out
#
#echo "terraform apply"
#terraform apply plan.out
#
#cd ../../..
#
#cd examples/data-sources/example_users
#
#rm -rf .terraform .terraform.lock.hcl plan.out terraform.tfstate terraform.tfstate.backup
#
#export TF_LOG_PROVIDER=INFO
#
#echo "terraform init"
#terraform init
#
#echo "terraform valid"
#terraform validate
#
#echo "terraform plan"
#terraform plan -out=plan.out
#
#echo "terraform apply"
#terraform apply plan.out
#
#cd ../../..

cd examples/resources/example_user

rm -rf .terraform .terraform.lock.hcl plan.out terraform.tfstate terraform.tfstate.backup

export TF_LOG_PROVIDER=INFO

echo "terraform init"
terraform init

echo "terraform valid"
terraform validate

echo "terraform plan"
terraform plan -out=plan.out

echo "terraform apply"
terraform apply plan.out

cd ../../..




