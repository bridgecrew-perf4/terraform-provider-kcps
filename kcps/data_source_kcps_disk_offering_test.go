package kcps

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccDataSourceKcpsDiskOffering(t *testing.T) {

	id := "f361c03c-ed19-4228-823c-745a5569aa62"
	name := "MIDDLE_STORAGE"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceKcpsDiskOfferingConfig_byId(id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.kcps_disk_offering.a", "id", id),
					resource.TestCheckResourceAttr("data.kcps_disk_offering.a", "name", name),
				),
			},
			{
				Config: testAccDataSourceKcpsDiskOfferingConfig_byName(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.kcps_disk_offering.a", "id", id),
					resource.TestCheckResourceAttr("data.kcps_disk_offering.a", "diskoffering_id", id),
				),
			},
		},
	})
}

func testAccDataSourceKcpsDiskOfferingConfig_byId(v string) string {
	return fmt.Sprintf(`
		data "kcps_disk_offering" "a" {
			diskoffering_id = "%s"
		}
		`, v)
}

func testAccDataSourceKcpsDiskOfferingConfig_byName(v string) string {
	return fmt.Sprintf(`
		data "kcps_disk_offering" "a" {
			name = "%s"
		}
		`, v)
}
