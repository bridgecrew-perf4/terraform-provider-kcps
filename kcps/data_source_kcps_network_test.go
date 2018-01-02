package kcps

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccDataSourceKcpsNetwork(t *testing.T) {

	id := "d0f15c14-4d94-4a69-b39a-994721bfc809"
	name := "MonitoringNetwork"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceKcpsNetworkConfig_byId(id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.kcps_network.a", "id", id),
					resource.TestCheckResourceAttr("data.kcps_network.a", "name", name),
				),
			},
			{
				Config: testAccDataSourceKcpsNetworkConfig_byName(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.kcps_network.a", "id", id),
					resource.TestCheckResourceAttr("data.kcps_network.a", "network_id", id),
				),
			},
		},
	})
}

func testAccDataSourceKcpsNetworkConfig_byId(networkId string) string {
	return fmt.Sprintf(`
		data "kcps_network" "a" {
			network_id = "%s"
		}
		`, networkId)
}

func testAccDataSourceKcpsNetworkConfig_byName(name string) string {
	return fmt.Sprintf(`
		data "kcps_network" "a" {
			name = "%s"
		}
		`, name)
}
