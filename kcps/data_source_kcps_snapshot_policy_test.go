package kcps

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	gk "github.com/uesyn/gokcps"
)

func TestAccDataSourceKcpsSnapshotPolicy(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckKcpsSnapshotPolicyDestroy,
		Steps: []resource.TestStep{

			{
				Config: testAccDataSourceKcpsSnapshotPolicyConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataSourceKcpsSnapshotPolicy(),
				),
			},
		},
	})
}

func testAccCheckKcpsSnapshotPolicyDestroy(s *terraform.State) error {
	cli := testAccProvider.Meta().(*gk.KCPSClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "kcps_snapshot_policy" {
			continue
		}

		if rs.Primary.ID == "" {
			continue
		}

		p := cli.Snapshot.NewListSnapshotPoliciesParams()
		p.SetVolumeid(rs.Primary.ID)

		_, err := cli.Snapshot.ListSnapshotPolicies(p)

		if err == nil {
			return fmt.Errorf("SnapshotPolicy still exists")
		}
	}

	return nil
}

func testAccCheckDataSourceKcpsSnapshotPolicy() resource.TestCheckFunc {
	return func(s *terraform.State) error {
		dsFullName := "data.kcps_snapshot_policy.a"
		rsFullName := "kcps_snapshot_policy.a"
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
			"maxsnaps",
			"schedule",
			"timezone",
			"intervaltype",
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

		//check other attributes
		dsAttrNames := []string{
			"snapshotpolicy_id",
		}
		dsCertainValues := []string{
			rsAttrs["id"],
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

func testAccDataSourceKcpsSnapshotPolicyConfig() string {
	return fmt.Sprintf(`
		resource kcps_snapshot_policy "a" {
			intervaltype = "WEEKLY"
			maxsnaps = 3
			schedule = "10:10:6"
			timezone = "JST"
			volumeid = "ffa4fba4-7468-42f6-a86c-85fe274fb8c1"
		}
		data "kcps_snapshot_policy" "a" {
			volumeid = "${kcps_snapshot_policy.a.volumeid}"
			snapshotpolicy_id = "${kcps_snapshot_policy.a.id}"
		}
		`)
}
