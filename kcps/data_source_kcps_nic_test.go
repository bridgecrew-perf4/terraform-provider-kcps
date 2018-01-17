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

func TestAccDataSourceKcpsNic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckKcpsNicDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceKcpsNicConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataSourceKcpsNic(),
				),
			},
		},
	})
}

func testAccCheckKcpsNicDestroy(s *terraform.State) error {
	cli := testAccProvider.Meta().(*gk.KCPSClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "kcps_nic" {
			continue
		}

		if rs.Primary.ID == "" {
			continue
		}

		p := cli.Nic.NewListNicsParams("be4a143f-3cdd-4091-9a8c-d82c78e49ddf") //set virtualmachine id
		p.SetNicid(rs.Primary.ID)

		_, err := cli.Nic.ListNics(p)
		if err == nil {
			return fmt.Errorf("Nic still exists")
		}
	}

	return nil
}

func testAccCheckDataSourceKcpsNic() resource.TestCheckFunc {
	return func(s *terraform.State) error {
		dsFullName := "data.kcps_nic.a"
		rsFullName := "kcps_nic.a"
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
			"networkid",
			"virtualmachineid",
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

		//secondaryip check
		rsSecondaryIpCount, ok := rsAttrs["secondaryip.#"]
		if !ok {
			return errors.New("can't find 'secondaryip' attribute in resource")
		}
		rsNoOfSecondaryIp, err := strconv.Atoi(rsSecondaryIpCount)
		if err != nil {
			return errors.New("failed to read number of secondaryip in resource")
		}
		dsSecondaryIpCount, ok := dsAttrs["secondaryip.#"]
		if !ok {
			return errors.New("can't find 'secondaryip' attribute in data source")
		}
		dsNoOfSecondaryIp, err := strconv.Atoi(dsSecondaryIpCount)
		if err != nil {
			return errors.New("failed to read number of secondaryip in data source")
		}
		if dsNoOfSecondaryIp != rsNoOfSecondaryIp {
			return fmt.Errorf(
				"expected %d number of cidrlist, got %d",
				rsNoOfSecondaryIp,
				dsNoOfSecondaryIp,
			)
		}

		for i := 0; i < dsNoOfSecondaryIp; i++ {
			stri := strconv.Itoa(i)

			//'secondaryip_id' existence check
			_, ok = dsAttrs["secondaryip."+stri+".secondaryip_id"]
			if !ok {
				return errors.New("can't find 'secondaryip_id' attribute in resource")
			}
		}

		//id check
		dsNicId, ok := dsAttrs["nic_id"]
		if !ok {
			return errors.New("can't find 'nic_id' attribute in data source")
		}
		if dsNicId != rsAttrs["id"] {
			return fmt.Errorf(
				"'nic_id': expected %s , but received %s", rsAttrs["id"], dsNicId,
			)
		}

		//check exeistence only
		dsAttrNames := []string{
			"ipaddress",
			"macaddress",
		}
		for _, v := range dsAttrNames {
			if _, ok := dsAttrs[v]; !ok {
				return fmt.Errorf("can't find '%s' attribute in resource", v)
			}
			if dsAttrs[v] == "" {
				return fmt.Errorf("'%s' isn't set any value", v)
			}
		}

		return nil
	}
}

func testAccDataSourceKcpsNicConfig() string {
	return fmt.Sprintf(`
		resource kcps_nic "a" {
			networkid = "d0f15c14-4d94-4a69-b39a-994721bfc809"
			virtualmachineid = "be4a143f-3cdd-4091-9a8c-d82c78e49ddf"
			secondaryip = ["198.18.57.252","198.18.57.254"]
		}
		data "kcps_nic" "a" {
			virtualmachineid = "be4a143f-3cdd-4091-9a8c-d82c78e49ddf"
			nic_id = "${kcps_nic.a.id}"
		}
		`)
}
