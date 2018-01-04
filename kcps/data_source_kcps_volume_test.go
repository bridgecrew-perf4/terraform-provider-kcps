package kcps

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	gk "github.com/uesyn/gokcps"
)

func TestAccDataSourceKcpsVolume(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckKcpsVolumeDestroy,
		Steps: []resource.TestStep{

			{
				Config: testAccDataSourceKcpsVolumeConfig_Snapshot(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataSourceKcpsVolume_Snapshot(),
				),
			},

			{
				Config: testAccDataSourceKcpsVolumeConfig_DiskOffering(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataSourceKcpsVolume_DiskOffering(),
				),
			},

			/*
				{
					Config: testAccDataSourceKcpsVolumeConfig_ServiceOffering(),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("data.kcps_volume.a", "serviceofferingid", "e3060950-1b4f-4adb-b050-bbe99694da19"),
					),
				},
			*/
		},
	})
}

func testAccCheckKcpsVolumeDestroy(s *terraform.State) error {
	cli := testAccProvider.Meta().(*gk.KCPSClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "kcps_volume" {
			continue
		}

		if rs.Primary.ID == "" {
			continue
		}

		p := cli.Volume.NewListVolumesParams()
		p.SetId(rs.Primary.ID)

		r, _ := cli.Volume.ListVolumes(p)
		if r.Volumes != nil {
			return fmt.Errorf("Volume still exists")
		}
	}

	return nil
}

func testAccCheckDataSourceKcpsVolume_Snapshot() resource.TestCheckFunc {
	return func(s *terraform.State) error {
		dsFullName := "data.kcps_volume.a"
		rsFullName := "kcps_volume.a"
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
			"name",
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
			"volume_id",
			"type",
			"zoneid",
			"diskofferingid",
		}
		dsCertainValues := []string{
			rsAttrs["id"],
			"DATADISK",
			"593697b6-c123-4025-b412-ef83822733e5",
			"cf5f671e-49c9-455e-a5b8-e563f9826e8b",
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

func testAccCheckDataSourceKcpsVolume_DiskOffering() resource.TestCheckFunc {
	return func(s *terraform.State) error {
		dsFullName := "data.kcps_volume.a"
		rsFullName := "kcps_volume.a"
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
			"name",
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
			"volume_id",
			"type",
			"virtualmachineid",
			"zoneid",
			"diskofferingid",
		}
		dsCertainValues := []string{
			rsAttrs["id"],
			"DATADISK",
			rsAttrs["attachto"],
			"593697b6-c123-4025-b412-ef83822733e5",
			rsAttrs["diskoffering.0.diskofferingid"],
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

func testAccDataSourceKcpsVolumeConfig_Snapshot() string {
	return fmt.Sprintf(`
		resource kcps_volume "a" {
			name = "hoge1"
			snapshot {
				snapshotid = "2d8a1ca6-6454-4e07-82c6-dcae245c01b1"
			}
		}
		data "kcps_volume" "a" {
			volume_id = "${kcps_volume.a.id}"
		}
		`)
}
func testAccDataSourceKcpsVolumeConfig_DiskOffering() string {
	return fmt.Sprintf(`
		resource kcps_volume "a" {
			name = "hoge2"
			diskoffering {
				zoneid = "593697b6-c123-4025-b412-ef83822733e5"
			    size = 100
			    diskofferingid = "bc1b5c0c-fcb3-4a7b-b8de-2c9d6952e0a5"
			}
			attachto = "fa608125-4658-4ff0-aab4-33a1b428988a"
		}
		data "kcps_volume" "a" {
			volume_id = "${kcps_volume.a.id}"
		}
		`)
}

func testAccDataSourceKcpsVolumeConfig_ServiceOffering() string {
	return fmt.Sprintf(`
		data "kcps_volume" "a" {
			volume_id = "270fe65e-9f47-4d0c-a77c-61904fd6333a"
		}
		`)
}
