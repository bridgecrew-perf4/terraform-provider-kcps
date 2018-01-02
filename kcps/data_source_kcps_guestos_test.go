package kcps

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccDataSourceKcpsGuestOS(t *testing.T) {

	id := "81765f44-7aae-11e4-b5b5-c45444131635"
	description := "Ubuntu 10.04 (64-bit)"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceKcpsGuestOSConfig_byId(id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.kcps_guestos.a", "id", id),
					resource.TestCheckResourceAttr("data.kcps_guestos.a", "description", description),
					resource.TestCheckResourceAttr("data.kcps_guestos.a", "oscategoryid", "816c54cc-7aae-11e4-b5b5-c45444131635"),
				),
			},
			{
				Config: testAccDataSourceKcpsGuestOSConfig_byDescription(description),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.kcps_guestos.a", "id", id),
					resource.TestCheckResourceAttr("data.kcps_guestos.a", "guestos_id", id),
					resource.TestCheckResourceAttr("data.kcps_guestos.a", "oscategoryid", "816c54cc-7aae-11e4-b5b5-c45444131635"),
				),
			},
		},
	})
}

func testAccDataSourceKcpsGuestOSConfig_byId(guestosId string) string {
	return fmt.Sprintf(`
		data "kcps_guestos" "a" {
			guestos_id = "%s"
		}
		`, guestosId)
}

func testAccDataSourceKcpsGuestOSConfig_byDescription(description string) string {
	return fmt.Sprintf(`
		data "kcps_guestos" "a" {
			description = "%s"
		}
		`, description)
}
