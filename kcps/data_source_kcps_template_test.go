package kcps

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	gk "github.com/uesyn/gokcps"
)

func TestAccDataSourceKcpsTemplate(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckKcpsTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceKcpsTemplateConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataSourceKcpsTemplate(),
				),
			},
		},
	})
}

func testAccCheckKcpsTemplateDestroy(s *terraform.State) error {
	cli := testAccProvider.Meta().(*gk.KCPSClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "kcps_template" {
			continue
		}

		if rs.Primary.ID == "" {
			continue
		}

		p := cli.Template.NewListTemplatesParams("self") //set templatefilter
		p.SetId(rs.Primary.ID)

		r, _ := cli.Template.ListTemplates(p)
		if r.Templates != nil {
			return fmt.Errorf("Template still exists")
		}
	}

	return nil
}

func testAccCheckDataSourceKcpsTemplate() resource.TestCheckFunc {
	return func(s *terraform.State) error {
		dsFullName := "data.kcps_template.a"
		rsFullName := "kcps_template.a"
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
			"name",
			"displaytext",
			"ostypeid",
			"isdynamicallyscalable",
			"ispublic",
			"passwordenabled",
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
			"template_id",
			"hypervisor",
			"format",
			"zoneid",
		}
		dsCertainValues := []string{
			rsAttrs["id"],
			"VMware",
			"OVA",
			"593697b6-c123-4025-b412-ef83822733e5",
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

func testAccDataSourceKcpsTemplateConfig() string {
	return fmt.Sprintf(`
		resource kcps_template "a" {
			name = "testa"
			displaytext = "test template"
			ostypeid = "817437d2-7aae-11e4-b5b5-c45444131635"
			snapshotid = "2d8a1ca6-6454-4e07-82c6-dcae245c01b1"
			isdynamicallyscalable = true
			passwordenabled = true
			ispublic = true
		}
		data kcps_template "a" {
			templatefilter = "self"
			template_id = "${kcps_template.a.id}"
		}
		`)
}
