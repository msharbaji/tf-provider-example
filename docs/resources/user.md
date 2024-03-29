---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "example_user Resource - example"
subcategory: ""
description: |-
  User resource
---

# example_user (Resource)

User resource

## Example Usage

```terraform
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
  username = "someone_else"
  email    = "someone_else@someone.com"
}

output "user_id" {
  value = example_user.user.id
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `email` (String) The email of the user
- `username` (String) The username of the user

### Read-Only

- `created_at` (String) The created_at of the user
- `id` (String) The id of the user
- `updated_at` (String) The updated_at of the user


