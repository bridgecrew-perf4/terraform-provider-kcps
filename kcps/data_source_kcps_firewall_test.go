package kcps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
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
					testAccCheckDataSourceKCPSFirewall(),
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

func testAccCheckDataSourceKCPSFirewall() resource.TestCheckFunc {
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

		attrsToTest := []string{
			"id",
			"ipaddressid",
			"ipaddress",
		}

		for _, attrToTest := range attrsToTest {
			if dsAttrs[attrToTest] != rsAttrs[attrToTest] {
				return fmt.Errorf("expected %s, but received %s", attrToTest, dsAttrs[attrToTest], rsAttrs[attrToTest])
			}
		}

		dsCidrlistCount, ok := dsAttrs["cidrlist.#"]
		if !ok {
			return errors.New("can't find 'cidrlist' attribute in data source")
		}

		dsNoOfCidrlist, err := strconv.Atoi(dsCidrlistCount)
		if err != nil {
			return errors.New("failed to read number of cidr in data source")
		}

		rsCidrlistCount, ok := rsAttrs["cidrlist.#"]
		if !ok {
			return errors.New("can't find 'cidrlist' attribute in resource")
		}

		rsNoOfCidrlist, err := strconv.Atoi(rsCidrlistCount)
		if err != nil {
			return errors.New("failed to read number of cidr in resource")
		}

		if dsNoOfCidrlist != rsNoOfCidrlist {
			return fmt.Errorf(
				"expected %d number of cidr, but received %d",
				rsNoOfCidrlist,
				dsNoOfCidrlist,
			)
		}

		//id check
		dsFirewallId, ok := dsAttrs["firewall_id"]
		if !ok {
			return errors.New("can't find 'firewall_id' attribute in data source")
		}
		if dsFirewallId != rsAttrs["id"] {
			return fmt.Errorf(
				"expected %d , but received %d",
				rsAttrs["id"],
				dsFirewallId,
			)
		}

		//protocol check
		dsProtocol, ok := dsAttrs["protocol"]
		if !ok {
			return errors.New("can't find 'protocol' attribute in data source")
		}
		if strings.ToUpper(dsProtocol) != rsAttrs["protocol"] {
			return fmt.Errorf(
				"expected %d , but received %d",
				rsAttrs["protocol"],
				strings.ToUpper(dsProtocol),
			)
		}

		return nil
	}
}

func testAccDataSourceKcpsFirewallConfig() string {
	return fmt.Sprintf(`
		resource kcps_firewall "a" {
			ipaddressid = "4c9e186e-9227-42a4-b14c-38ed7f32b012"
			protocol = "TCP"
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
			protocol = "ICMP"
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
