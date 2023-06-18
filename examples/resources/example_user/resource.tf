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

resource "example_user" "user" {
  username = "someone_else3"
  email    = "someone_else3@someone.com"
}

output "user_id" {
  value = example_user.user.id
}
