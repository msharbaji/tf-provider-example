package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testAccUsersDataSourceConfig = `
provider "example" {
  endpoint = "localhost:50051"
  key_id = "my-key-id"
  secret_key = "my-secret-key"
}

data "example_users" "users" {
}
`

func TestAccUsersDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUsersDataSourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.example_users.users", "users.#", "3"),
					resource.TestCheckResourceAttr("data.example_users.users", "users.0.id", "1"),
					resource.TestCheckResourceAttr("data.example_users.users", "users.0.username", "someone"),
					resource.TestCheckResourceAttr("data.example_users.users", "users.0.email", "someone@someone.com"),

					resource.TestCheckResourceAttr("data.example_users.users", "users.1.id", "2"),
					resource.TestCheckResourceAttr("data.example_users.users", "users.1.username", "someone_else"),
					resource.TestCheckResourceAttr("data.example_users.users", "users.1.email", "someonce2@someone.com"),
				),
			},
		},
	})
}
