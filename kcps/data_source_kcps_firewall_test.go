package kcps

import (
	"errors"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	gk "github.com/uesyn/gokcps"
)

func TestAccDataSourceKcpsFirewall(t *testing.T) {

	icmpcode := "0"
	icmptype := "3"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckKcpsFirewallDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceKcpsFirewallConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataSourceKcpsFirewall(),
				),
			},

			{
				Config: testAccDataSourceKcpsFirewallConfig_ICMP(icmpcode, icmptype),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.kcps_firewall.a", "icmpcode", icmpcode),
					resource.TestCheckResourceAttr("data.kcps_firewall.a", "icmptype", icmptype),
				),
			},
		},
	})
}

func testAccCheckKcpsFirewallDestroy(s *terraform.State) error {
	cli := testAccProvider.Meta().(*gk.KCPSClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "kcps_firewall" {
			continue
		}

		if rs.Primary.ID == "" {
			continue
		}

		p := cli.Firewall.NewListFirewallRulesParams()
		p.SetId(rs.Primary.ID)

		_, err := cli.Firewall.ListFirewallRules(p)
		if err == nil {
			return fmt.Errorf("Firewall still exists")
		}
	}

	return nil
}

func testAccCheckDataSourceKcpsFirewall() resource.TestCheckFunc {
	return func(s *terraform.State) error {
		dsFullName := "data.kcps_firewall.a"
		rsFullName := "kcps_firewall.a"
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
			"ipaddressid",
			"ipaddress",
			"protocol",
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

		//cidrlist check
		dsCidrlistCount, ok := dsAttrs["cidrlist.#"]
		if !ok {
			return errors.New("can't find 'cidrlist' attribute in data source")
		}

		dsNoOfCidrlist, err := strconv.Atoi(dsCidrlistCount)
		if err != nil {
			return errors.New("failed to read number of cidrlist in data source")
		}

		rsCidrlistCount, ok := rsAttrs["cidrlist.#"]
		if !ok {
			return errors.New("can't find 'cidrlist' attribute in resource")
		}

		rsNoOfCidrlist, err := strconv.Atoi(rsCidrlistCount)
		if err != nil {
			return errors.New("failed to read number of cidrlist in resource")
		}

		if dsNoOfCidrlist != rsNoOfCidrlist {
			return fmt.Errorf(
				"expected %d number of cidrlist, got %d",
				rsNoOfCidrlist,
				dsNoOfCidrlist,
			)
		}

		//other attributes check
		dsAttrNames := []string{
			"firewall_id",
			"startport",
			"endport",
		}
		dsCertainValues := []string{
			rsAttrs["id"],
			rsAttrs["port.0.startport"],
			rsAttrs["port.0.endport"],
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

func testAccDataSourceKcpsFirewallConfig() string {
	return fmt.Sprintf(`
		resource kcps_firewall "a" {
			ipaddressid = "4c9e186e-9227-42a4-b14c-38ed7f32b012"
			protocol = "tcp"
			cidrlist = ["1.1.1.1/32", "2.2.2.2/32"]
			port {
				startport = 1020
				endport = 1023
			}
		}
		data "kcps_firewall" "a" {
			firewall_id = "${kcps_firewall.a.id}"
		}
		`)
}

func testAccDataSourceKcpsFirewallConfig_ICMP(icmpcode, icmptype string) string {
	return fmt.Sprintf(`
		resource kcps_firewall "a" {
			ipaddressid = "4c9e186e-9227-42a4-b14c-38ed7f32b012"
			protocol = "icmp"
			cidrlist = ["3.3.3.3/32", "4.4.4.4/32"]
			icmp {
				icmpcode = %s
				icmptype = %s
			}
		}
		data "kcps_firewall" "a" {
			firewall_id = "${kcps_firewall.a.id}"
		}
		`, icmpcode, icmptype)
}
