package kcps

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	gk "github.com/uesyn/gokcps"
)

func TestAccDataSourceKcpsSnapshot(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckKcpsSnapshotDestroy,
		Steps: []resource.TestStep{

			{
				Config: testAccDataSourceKcpsSnapshotConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataSourceKcpsSnapshot(),
				),
			},
		},
	})
}

func testAccCheckKcpsSnapshotDestroy(s *terraform.State) error {
	cli := testAccProvider.Meta().(*gk.KCPSClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "kcps_snapshot" {
			continue
		}

		if rs.Primary.ID == "" {
			continue
		}

		p := cli.Snapshot.NewListSnapshotsParams()
		p.SetId(rs.Primary.ID)

		r, _ := cli.Snapshot.ListSnapshots(p)
		if r.Snapshots != nil {
			return fmt.Errorf("Snapshot still exists")
		}
	}

	return nil
}

func testAccCheckDataSourceKcpsSnapshot() resource.TestCheckFunc {
	return func(s *terraform.State) error {
		dsFullName := "data.kcps_snapshot.a"
		rsFullName := "kcps_snapshot.a"
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
			"volumeid",
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
			"snapshot_id",
			"intervaltype",
		}
		dsCertainValues := []string{
			rsAttrs["id"],
			"MANUAL",
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
			"name",
		}
		for _, v := range dsAttrNames {
			if _, ok := dsAttrs[v]; !ok {
				return fmt.Errorf("can't find '%s' attribute in data source", v)
			}
		}

		return nil
	}
}

func testAccDataSourceKcpsSnapshotConfig() string {
	return fmt.Sprintf(`
		resource kcps_snapshot "a" {
			"volumeid"= "ffa4fba4-7468-42f6-a86c-85fe274fb8c1"
		}
		data "kcps_snapshot" "a" {
			snapshot_id = "${kcps_snapshot.a.id}"
		}
		`)
}
