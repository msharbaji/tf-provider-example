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
				),
			},
		},
	})
}
