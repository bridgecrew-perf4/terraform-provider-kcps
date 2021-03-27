package kcps

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccDataSourceKcpsServiceOffering(t *testing.T) {

	id := "2355f1ca-4912-457e-8c88-6e97ab177ea4"
	name := "Premium_10vCPU_Mem24GB"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceKcpsServiceOfferingConfig_byId(id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.kcps_service_offering.a", "id", id),
					resource.TestCheckResourceAttr("data.kcps_service_offering.a", "name", name),
				),
			},
			{
				Config: testAccDataSourceKcpsServiceOfferingConfig_byName(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.kcps_service_offering.a", "id", id),
					resource.TestCheckResourceAttr("data.kcps_service_offering.a", "serviceoffering_id", id),
				),
			},
		},
	})
}

func testAccDataSourceKcpsServiceOfferingConfig_byId(serviceOfferingId string) string {
	return fmt.Sprintf(`
		data "kcps_service_offering" "a" {
			serviceoffering_id = "%s"
		}
		`, serviceOfferingId)
}

func testAccDataSourceKcpsServiceOfferingConfig_byName(name string) string {
	return fmt.Sprintf(`
		data "kcps_service_offering" "a" {
			name = "%s"
		}
		`, name)
}
