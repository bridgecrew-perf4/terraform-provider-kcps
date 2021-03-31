package kcps

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	gk "github.com/uesyn/gokcps"
)

func TestAccDataSourceKcpsVMSnapshot(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckKcpsVMSnapshotDestroy,
		Steps: []resource.TestStep{

			{
				Config: testAccDataSourceKcpsVMSnapshotConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataSourceKcpsVMSnapshot(),
				),
			},
		},
	})
}

func testAccCheckKcpsVMSnapshotDestroy(s *terraform.State) error {
	cli := testAccProvider.Meta().(*gk.KCPSClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "kcps_vmsnapshot" {
			continue
		}

		if rs.Primary.ID == "" {
			continue
		}

		p := cli.Snapshot.NewListVMSnapshotParams()
		p.SetVmsnapshotid(rs.Primary.ID)

		r, _ := cli.Snapshot.ListVMSnapshot(p)
		if r.VMSnapshot != nil {
			return fmt.Errorf("VMSnapshot still exists")
		}
	}

	return nil
}

func testAccCheckDataSourceKcpsVMSnapshot() resource.TestCheckFunc {
	return func(s *terraform.State) error {
		dsFullName := "data.kcps_vmsnapshot.a"
		rsFullName := "kcps_vmsnapshot.a"
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
			"virtualmachineid",
			"displayname",
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

		//id check
		dsVMSnapshotId, _ := dsAttrs["vmsnapshot_id"]
		if dsVMSnapshotId != rsAttrs["id"] {
			return fmt.Errorf(
				"expected %s , but received %s", rsAttrs["id"], dsVMSnapshotId,
			)
		}

		//check data source only attributes
		dsAttrNames := []string{
			"state",
		}
		dsCertainValues := []string{
			"Ready",
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

		//check exeistence only
		dsAttrNames = []string{
			"displayname",
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

func testAccDataSourceKcpsVMSnapshotConfig() string {
	return fmt.Sprintf(`
		resource kcps_vmsnapshot "a" {
			virtualmachineid = "b28d0d22-853b-496d-8e3a-31a1bdbf2eff"
			displayname = "example-0"
		}
		data "kcps_vmsnapshot" "a" {
			vmsnapshot_id = "${kcps_vmsnapshot.a.id}"
		}
		`)
}
