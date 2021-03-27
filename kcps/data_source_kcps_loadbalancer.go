package kcps

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	gk "github.com/uesyn/gokcps"
)

func dataSourceKcpsLoadBalancer() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceKcpsLoadBalancerRead,

		Schema: map[string]*schema.Schema{
			"loadbalancer_id": {
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

func dataSourceKcpsLoadBalancerRead(d *schema.ResourceData, meta interface{}) error {
	cli := meta.(*gk.KCPSClient)

	p := cli.LoadBalancer.NewListLoadBalancerRulesParams()

	if loadbalancerId, ok := d.GetOk("loadbalancer_id"); ok {
		p.SetId(loadbalancerId.(string))
	}
	if keyword, ok := d.GetOk("keyword"); ok {
		p.SetKeyword(keyword.(string))
	}
	if name, ok := d.GetOk("name"); ok {
		p.SetName(name.(string))
	}
	if networkid, ok := d.GetOk("networkid"); ok {
		p.SetNetworkid(networkid.(string))
	}
	if publicipid, ok := d.GetOk("publicipid"); ok {
		p.SetPublicipid(publicipid.(string))
	}
	if virtualmachineid, ok := d.GetOk("virtualmachineid"); ok {
		p.SetVirtualmachineid(virtualmachineid.(string))
	}
	if zoneid, ok := d.GetOk("zoneid"); ok {
		p.SetZoneid(zoneid.(string))
	}

	r, err := cli.LoadBalancer.ListLoadBalancerRules(p)
	if err != nil {
		return fmt.Errorf("Error getting LoadBalancer Rule list: %s", err)
	}
	if r.LoadBalancerRules == nil {
		return fmt.Errorf("LoadBalancer Rule not found")
	}
	v := r.LoadBalancerRules[0]

	pri, _ := strconv.Atoi(v.Privateport)
	pub, _ := strconv.Atoi(v.Publicport)

	d.Set("loadbalancer_id", v.Id)
	d.Set("name", v.Name)
	d.Set("networkid", v.Networkid)
	d.Set("publicipid", v.Publicipid)
	d.Set("zoneid", v.Zoneid)

	d.Set("algorithm", v.Algorithm)
	d.Set("publicip", v.Publicip)
	d.Set("tags", flattenTags(v.Tags))
	d.Set("privateport", pri)
	d.Set("publicport", pub)

	d.SetId(v.Id)

	return nil
}
