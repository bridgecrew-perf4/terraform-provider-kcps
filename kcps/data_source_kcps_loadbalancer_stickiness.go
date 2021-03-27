package kcps

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceKcpsLoadBalancerStickiness() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceKcpsLoadBalancerStickinessRead,

		Schema: map[string]*schema.Schema{
			"loadbalancer_stickiness_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"keyword": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"networkid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"publicipid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"virtualmachineid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"zoneid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"algorithm": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"publicip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tags": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"value": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"privateport": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"publicport": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceKcpsLoadBalancerStickinessRead(d *schema.ResourceData, meta interface{}) error {

	return nil
}
