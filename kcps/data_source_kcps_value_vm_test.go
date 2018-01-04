package kcps

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	gk "github.com/uesyn/gokcps"
)

func TestAccDataSourceKcpsValueVM(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		//CheckDestroy: testAccCheckKcpsValueVMDestroy,
		Steps: []resource.TestStep{

			{
				Config: testAccDataSourceKcpsValueVMConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataSourceKcpsValueVM(),
				),
			},
		},
	})
}

func testAccCheckKcpsValueVMDestroy(s *terraform.State) error {
	cli := testAccProvider.Meta().(*gk.KCPSClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "kcps_value_vm" {
			continue
		}

		if rs.Primary.ID == "" {
			continue
		}

		p := cli.VirtualMachine.NewListVirtualMachinesParams()
		p.SetId(rs.Primary.ID)

		r, _ := cli.VirtualMachine.ListVirtualMachines(p)
		if r.VirtualMachines != nil {
			return fmt.Errorf("ValueVM still exists")
		}
	}

	return nil
}

func testAccCheckDataSourceKcpsValueVM() resource.TestCheckFunc {
	return func(s *terraform.State) error {
		dsFullName := "data.kcps_value_vm.a"
		rsFullName := "kcps_value_vm.a"
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
			"serviceofferingid",
			"templateid",
			"zoneid",
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

		//other attributes check
		dsAttrNames := []string{
			"valuevm_id",
			"hypervisor",
			"name",
		}
		dsCertainValues := []string{
			rsAttrs["id"],
			"VMware",
			"v-testvma-M16503331",
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

func testAccDataSourceKcpsValueVMConfig() string {
	return fmt.Sprintf(`
		resource kcps_value_vm "a" {
			name = "testvma"
			serviceofferingid = "e3060950-1b4f-4adb-b050-bbe99694da19"
			templateid ="0eb72664-f7ad-4d36-be3e-4f4c32ffe0e5"
			zoneid = "593697b6-c123-4025-b412-ef83822733e5"

			diskoffering{
				diskofferingid = "10cc47d1-2e04-4aeb-aec9-fb08d273198e"
				size = 100
			}
		}
		data "kcps_value_vm" "a" {
			valuevm_id = "${kcps_value_vm.a.id}"
		}
		`)
}
