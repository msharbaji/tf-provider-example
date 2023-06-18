terraform {
  required_providers {
    example = {
      version = "~> 1.0.0"
      source  = "malsharbaji.com/providers/example"
    }
  }
}

provider "example" {
  endpoint   = "localhost:50051"
  key_id     = "my-key-id"
  secret_key = "my-secret-key"
}

data "example_users" "users" {
}

output "users" {
  value = data.example_users.users
}
