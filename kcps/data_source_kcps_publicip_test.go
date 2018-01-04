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

		attrNames := []string{
			"id",
			"ipaddress",
		}

		for _, attrName := range attrNames {
			dsAttr, ok := dsAttrs[attrName]
			if !ok {
				return fmt.Errorf("can't find '%s' attribute in data source", attrName)
			}
			rsAttr, ok := rsAttrs[attrName]
			if !ok {
				return fmt.Errorf("can't find '%s' attribute in data source", attrName)
			}

			if dsAttr != rsAttr {
				return fmt.Errorf("'%s': expected %s, got %s", attrName, rsAttr, dsAttr)
			}
		}

		//check other attributes
		dsAttrNames := []string{
			"publicip_id",
			"issourcenat",
			"isstaticnat",
			"zoneid",
			"networkid",
			"associatednetworkid",
		}
		dsCertainValues := []string{
			rsAttrs["id"],
			"false",
			"true",
			"593697b6-c123-4025-b412-ef83822733e5",
			"8594fb24-2cdf-41ed-9290-60f32b6ee0db",
			rsAttrs["networkid"],
		}

		for i, _ := range dsAttrNames {
			dsAttr, ok := dsAttrs[dsAttrNames[i]]
			if !ok {
				return fmt.Errorf("can't find '%s' attribute in data source", dsAttrNames[i])
			}

			if dsAttr != dsCertainValues[i] {
				return fmt.Errorf(
					"'%s': expected %s , but received %s",
					dsAttrNames[i],
					dsCertainValues[i],
					dsAttr,
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
