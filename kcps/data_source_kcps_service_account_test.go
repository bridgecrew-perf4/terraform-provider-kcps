package kcps

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccDataSourceKcpsAccount(t *testing.T) {
	domain := "M99991415"
	account := "nos1415"

	resource.Test(t, resource.TestCase{
		//PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceKcpsAccountConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.kcps_service_account.a", "account", account),
					resource.TestCheckResourceAttr("data.kcps_service_account.a", "domain", domain),
				),
			},
		},
	})
}

func testAccDataSourceKcpsAccountConfig() string {
	return fmt.Sprintf(`
		data "kcps_service_account" "a" {}
	`)
}
