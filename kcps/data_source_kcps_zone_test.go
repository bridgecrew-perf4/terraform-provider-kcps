package kcps

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccDataSourceKcpsZone(t *testing.T) {

	name := "jp2-east03"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceKcpsZoneConfig(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.kcps_zone.a", "id", "593697b6-c123-4025-b412-ef83822733e5"),
					resource.TestCheckResourceAttr("data.kcps_zone.a", "zone_id", "593697b6-c123-4025-b412-ef83822733e5"),
					resource.TestCheckResourceAttr("data.kcps_zone.a", "name", name),
				),
			},
		},
	})
}

func testAccDataSourceKcpsZoneConfig(name string) string {
	return fmt.Sprintf(`
		data "kcps_zone" "a" {
			name = "%s"
		}
		`, name)
}
