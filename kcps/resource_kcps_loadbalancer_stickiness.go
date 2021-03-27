package kcps

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	gk "github.com/uesyn/gokcps"
)

func resourceKcpsLoadBalancerStickiness() *schema.Resource {
	return &schema.Resource{
		Create: resourceKcpsLoadBalancerStickinessCreate,
		Read:   resourceKcpsLoadBalancerStickinessRead,
		Delete: resourceKcpsLoadBalancerStickinessDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"lbruleid": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"methodname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},

			"param": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
					},
				},
			},
		},
	}
}

func resourceKcpsLoadBalancerStickinessCreate(d *schema.ResourceData, meta interface{}) error {
	cli := meta.(*gk.KCPSClient)

	lbruleid := d.Get("lbruleid").(string)
	methodname := d.Get("methodname").(string)
	name := d.Get("name").(string)

	mutexKV.Lock("loadbalancer-stckiness-" + lbruleid)
	defer mutexKV.Unlock("loadbalancer-stckiness-" + lbruleid)

	p := cli.LoadBalancer.NewCreateLBStickinessPolicyParams(
		lbruleid,
		methodname,
		name,
	)

	/*
	*
	*
	*
	*set param
	*
	*
	 */

	r, err := cli.LoadBalancer.CreateLBStickinessPolicy(p)
	if err != nil {
		return fmt.Errorf("Error creating new LoadBalancer Stickiness Policy: %s", err)
	}

	d.SetId(r.Stickinesspolicy[0].Id)

	return resourceKcpsLoadBalancerStickinessRead(d, meta)
}

func resourceKcpsLoadBalancerStickinessRead(d *schema.ResourceData, meta interface{}) error {

	return nil
}

func resourceKcpsLoadBalancerStickinessDelete(d *schema.ResourceData, meta interface{}) error {

	d.SetId("")
	return nil
}
