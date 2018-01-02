package kcps

import (
	"errors"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
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
			"protocol",
			"virtualmachineid",
			"vmguestip",
		}

		for _, attrToTest := range attrsToTest {
			if dsAttrs[attrToTest] != rsAttrs[attrToTest] {
				return fmt.Errorf("'%s': expected %s, got %s", attrToTest, rsAttrs[attrToTest], dsAttrs[attrToTest])
			}
		}

		//id check
		dsNatPortForwardId, ok := dsAttrs["natportforward_id"]
		if !ok {
			return errors.New("can't find 'natportforward_id' attribute in data source")
		}
		if dsNatPortForwardId != rsAttrs["id"] {
			return fmt.Errorf(
				"'natportforward_id': expected %s , but received %s",
				rsAttrs["id"],
				dsNatPortForwardId,
			)
		}

		//port check
		rsPorts := []string{
			rsAttrs["port.0.privateport"],
			rsAttrs["port.0.privateendport"],
			rsAttrs["port.0.publicport"],
			rsAttrs["port.0.publicendport"],
		}
		dsPorts := []string{
			dsAttrs["privateport"],
			dsAttrs["privateendport"],
			dsAttrs["publicport"],
			dsAttrs["publicendport"],
		}
		for i, _ := range dsPorts {
			if dsPorts[i] != rsPorts[i] {
				return fmt.Errorf(
					"expected port %s , but received %s", rsPorts[i], dsPorts[i],
				)
			}
		}

		//check data source only attributes
		trueNetworkId := "7b921a2d-8c82-4016-bbd0-cf0dd6877408"
		if dsAttrs["networkid"] != trueNetworkId {
			return fmt.Errorf(
				"'networkid': expected %s , but received %s", trueNetworkId, dsAttrs["networkid"],
			)
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
