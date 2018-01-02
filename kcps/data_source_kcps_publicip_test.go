package kcps

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	gk "github.com/uesyn/gokcps"
)

func TestAccDataSourceKcpsPublicIP(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckKcpsPublicIPDestroy,
		Steps: []resource.TestStep{

			{
				Config: testAccDataSourceKcpsPublicIPConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.kcps_publicip.a", "issourcenat", "false"),
					resource.TestCheckResourceAttr("data.kcps_publicip.a", "isstaticnat", "false"),
				),
			},

			{
				Config: testAccDataSourceKcpsPublicIPConfig_StaticNAT(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataSourceKcpsPublicIP_StaticNAT(),
				),
			},
		},
	})
}

func testAccCheckKcpsPublicIPDestroy(s *terraform.State) error {
	cli := testAccProvider.Meta().(*gk.KCPSClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "kcps_publicip" {
			continue
		}

		if rs.Primary.ID == "" {
			continue
		}

		p := cli.Nic.NewListPublicIpAddressesParams()
		p.SetId(rs.Primary.ID)

		r, _ := cli.Nic.ListPublicIpAddresses(p)
		if r.PublicIpAddresses != nil {
			return fmt.Errorf("Public IP still exists")
		}
	}

	return nil
}

func testAccCheckDataSourceKcpsPublicIP_StaticNAT() resource.TestCheckFunc {
	return func(s *terraform.State) error {
		dsFullName := "data.kcps_publicip.a"
		rsFullName := "kcps_publicip.a"
		ds, ok := s.RootModule().Resources[dsFullName]
		if !ok {
			return fmt.Errorf("cant' find resource called %s in state", dsFullName)
		}

		rs, ok := s.RootModule().Resources[rsFullName]
		if !ok {
			return fmt.Errorf("can't find data source called %s in state", rsFullName)
		}

		dsAttrs := ds.Primary.Attributes
		rsAttrs := rs.Primary.Attributes

		attrsToTest := []string{
			"id",
			"ipaddress",
		}

		for _, attrToTest := range attrsToTest {
			if dsAttrs[attrToTest] != rsAttrs[attrToTest] {
				return fmt.Errorf("'%s': expected %s, got %s", attrToTest, rsAttrs[attrToTest], dsAttrs[attrToTest])
			}
		}

		//id check
		dsPublicIpId, _ := dsAttrs["publicip_id"]
		if dsPublicIpId != rsAttrs["id"] {
			return fmt.Errorf(
				"expected %d , but received %d", rsAttrs["id"], dsPublicIpId,
			)
		}

		//networkid check ('networkid' of Data Source is different from networkid' of Resource)
		dsAssociatedNetworkId, _ := dsAttrs["associatednetworkid"]
		if dsAssociatedNetworkId != rsAttrs["networkid"] {
			return fmt.Errorf(
				"expected %d , but received %d", rsAttrs["networkid"], dsAssociatedNetworkId,
			)
		}

		//check data source only attributes
		attrName := []string{
			"issourcenat",
			"isstaticnat",
			"zoneid",
			"networkid",
		}
		dsCertainValues := []string{
			"false",
			"true",
			"593697b6-c123-4025-b412-ef83822733e5",
			"8594fb24-2cdf-41ed-9290-60f32b6ee0db",
		}

		for i, _ := range attrName {
			if dsAttrs[attrName[i]] != dsCertainValues[i] {
				return fmt.Errorf(
					"'%s': expected %s , but received %s",
					attrName[i],
					dsCertainValues[i],
					dsAttrs[attrName[i]],
				)
			}
		}

		return nil
	}
}

func testAccDataSourceKcpsPublicIPConfig() string {
	return fmt.Sprintf(`
		resource kcps_publicip "a" {
			networkid = "7b921a2d-8c82-4016-bbd0-cf0dd6877408"
		}
		data "kcps_publicip" "a" {
			publicip_id = "${kcps_publicip.a.id}"
		}
		`)
}

func testAccDataSourceKcpsPublicIPConfig_StaticNAT() string {
	return fmt.Sprintf(`
		resource kcps_publicip "a" {
			networkid = "7b921a2d-8c82-4016-bbd0-cf0dd6877408"
    
			staticnat {
				virtualmachineid = "b107347a-3447-4e9f-8778-ed49941b5821"
				vmguestip = "10.1.1.13"
			}
		}
		data "kcps_publicip" "a" {
			publicip_id = "${kcps_publicip.a.id}"
		}
		`)
}
