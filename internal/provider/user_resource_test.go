package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccUserResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUserResourceConfig("someone_test", "someone_test@someone.com"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("example_user.test", "username", "someone_test"),
					resource.TestCheckResourceAttr("example_user.test", "email", "someone_test@someone.com"),
				),
			},
		},
	})
}

func testAccUserResourceConfig(username, email string) string {
	return fmt.Sprintf(`
provider "example" {
  endpoint = "localhost:50051"
  key_id = "my-key-id"
  secret_key = "my-secret-key"
}

resource "example_user" "test" {
  username = %[1]q
  email = %[2]q
}
`, username, email)
}
