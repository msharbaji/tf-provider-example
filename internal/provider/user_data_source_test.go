package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testAccUserDataSourceConfig = `
provider "example" {
  endpoint = "localhost:50051"
  key_id = "my-key-id"
  secret_key = "my-secret-key"
}

data "example_user" "user" {
	username = "someone_else"
}
`

func TestAccUserDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUserDataSourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.example_user.user", "id", "2"),
					resource.TestCheckResourceAttr("data.example_user.user", "username", "someone_else"),
				),
			},
		},
	})
}
