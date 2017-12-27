package kcps

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceKcpsHost() *schema.Resource {
	return &schema.Resource{
		Create: resourceKcpsHostCreate,
		Read:   resourceKcpsHostRead,
		Update: resourceKcpsHostUpdate,
		Delete: resourceKcpsHostDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"zoneid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateZoneId(),
			},

			/* gokcpsでVMware以外のhypervisorを設定できるようになった場合
			"hypervisor": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateHypervisor(),
			},
			*/

			"number": {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validateNumber(),
			},

			"distributiongroup": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceKcpsHostCreate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceKcpsHostRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceKcpsHostUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceKcpsHostDelete(d *schema.ResourceData, meta interface{}) error {
	d.SetId("")
	return nil
}
