package kcps

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceKcpsHost() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceKcpsHostRead,

		Schema: map[string]*schema.Schema{
			"zoneid": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"hypervisor": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"number": {
				Type:     schema.TypeInt,
				Computed: true,
			},

			"distributiongroup": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceKcpsHostRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}
