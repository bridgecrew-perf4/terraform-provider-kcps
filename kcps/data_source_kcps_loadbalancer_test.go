package kcps

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	gk "github.com/uesyn/gokcps"
)

func TestAccDataSourceKcpsLoadBalancer(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckKcpsLoadBalancerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceKcpsLoadBalancerConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataSourceKcpsLoadBalancer(),
				),
			},
		},
	})
}

func testAccCheckKcpsLoadBalancerDestroy(s *terraform.State) error {
	cli := testAccProvider.Meta().(*gk.KCPSClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "kcps_loadbalancer" {
			continue
		}

		if rs.Primary.ID == "" {
			continue
		}

		p := cli.LoadBalancer.NewListLoadBalancerRulesParams()
		p.SetId(rs.Primary.ID)

		_, err := cli.LoadBalancer.ListLoadBalancerRules(p)
		if err == nil {
			return fmt.Errorf("LoadBalancer still exists")
		}
	}

	return nil
}

func testAccCheckDataSourceKcpsLoadBalancer() resource.TestCheckFunc {
	return func(s *terraform.State) error {
		dsFullName := "data.kcps_loadbalancer.a"
		rsFullName := "kcps_loadbalancer.a"
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
			"algorithm",
			"name",
			"privateport",
			"publicport",
			"publicipid",
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

		//other attributes check
		dsAttrNames := []string{
			"loadbalancer_id",
			"networkid",
			"zoneid",
			"publicip",
		}
		dsCertainValues := []string{
			rsAttrs["id"],
			"7b921a2d-8c82-4016-bbd0-cf0dd6877408",
			"593697b6-c123-4025-b412-ef83822733e5",
			"27.85.233.111",
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

func testAccDataSourceKcpsLoadBalancerConfig() string {
	return fmt.Sprintf(`
		resource kcps_loadbalancer "a" {
			algorithm = "source"
			name = "testing"
			privateport = 8080
			publicport = 9080
			publicipid = "b2a3a2b6-67ff-46d2-877b-25618824b1ae"
		
			assignto = ["fa608125-4658-4ff0-aab4-33a1b428988a",
						"be4a143f-3cdd-4091-9a8c-d82c78e49ddf" ] 
		}
		data "kcps_loadbalancer" "a" {
			loadbalancer_id = "${kcps_loadbalancer.a.id}"
		}
		`)
}
