package kcps

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	gk "github.com/uesyn/gokcps"
)

func TestAccDataSourceKcpsNatPortForward(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckKcpsNatPortForwardDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceKcpsNatPortForwardConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataSourceKcpsNatPortForward(),
				),
			},
		},
	})
}

func testAccCheckKcpsNatPortForwardDestroy(s *terraform.State) error {
	cli := testAccProvider.Meta().(*gk.KCPSClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "kcps_nat_portforward" {
			continue
		}

		if rs.Primary.ID == "" {
			continue
		}

		p := cli.NatPortForward.NewListPortForwardingRulesParams()
		p.SetId(rs.Primary.ID)

		_, err := cli.NatPortForward.ListPortForwardingRules(p)
		if err == nil {
			return fmt.Errorf("Port Forwarding Rule still exists")
		}
	}

	return nil
}

func testAccCheckDataSourceKcpsNatPortForward() resource.TestCheckFunc {
	return func(s *terraform.State) error {
		dsFullName := "data.kcps_nat_portforward.a"
		rsFullName := "kcps_nat_portforward.a"
		ds, ok := s.RootModule().Resources[dsFullName]
		if !ok {
			return fmt.Errorf("cant' find data source called %s in state", dsFullName)
		}

		rs, ok := s.RootModule().Resources[rsFullName]
		if !ok {
			return fmt.Errorf("can't find resource called %s in state", rsFullName)
		}

		dsAttrs := ds.Primary.Attributes
		rsAttrs := rs.Primary.Attributes

		attrNames := []string{
			"id",
			"ipaddressid",
			"ipaddress",
			"protocol",
			"virtualmachineid",
			"vmguestip",
		}

		for _, attrName := range attrNames {
			dsAttr, ok := dsAttrs[attrName]
			if !ok {
				return fmt.Errorf("can't find '%s' attribute in data source", attrName)
			}
			rsAttr, ok := rsAttrs[attrName]
			if !ok {
				return fmt.Errorf("can't find '%s' attribute in resource", attrName)
			}

			if dsAttr != rsAttr {
				return fmt.Errorf("'%s': expected %s, got %s", attrName, rsAttr, dsAttr)
			}
		}

		//other attributes check
		dsAttrNames := []string{
			"natportforward_id",
			"privateport",
			"privateendport",
			"publicport",
			"publicendport",
			"networkid",
		}
		dsCertainValues := []string{
			rsAttrs["id"],
			rsAttrs["port.0.privateport"],
			rsAttrs["port.0.privateendport"],
			rsAttrs["port.0.publicport"],
			rsAttrs["port.0.publicendport"],
			"7b921a2d-8c82-4016-bbd0-cf0dd6877408",
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

func testAccDataSourceKcpsNatPortForwardConfig() string {
	return fmt.Sprintf(`
		resource kcps_nat_portforward "a" {
			ipaddressid = "165bd632-738b-44e4-8449-562fcd6da509"
			protocol = "tcp"
			port {
				privateport = 1022
				privateendport = 1023
				publicport = 1022
				publicendport = 1023
			}
			
			virtualmachineid = "be4a143f-3cdd-4091-9a8c-d82c78e49ddf"
			vmguestip="10.1.1.56"
		}
		data "kcps_nat_portforward" "a" {
			natportforward_id = "${kcps_nat_portforward.a.id}"
		}
		`)
}
