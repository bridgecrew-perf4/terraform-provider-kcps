package kcps

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
	gk "github.com/uesyn/gokcps"
)

func dataSourceKcpsNatPortForward() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceKcpsNatPortForwardRead,

		Schema: map[string]*schema.Schema{
			"natportforward_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipaddressid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"keyword": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"networkid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"ipaddress": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"privateport": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"privateendport": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"protocol": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"publicport": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"publicendport": {
				Type:     schema.TypeInt,
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
			"virtualmachineid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vmguestip": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceKcpsNatPortForwardRead(d *schema.ResourceData, meta interface{}) error {
	cli := meta.(*gk.KCPSClient)

	p := cli.NatPortForward.NewListPortForwardingRulesParams()
	if natportforwardId, ok := d.GetOk("natportforward_id"); ok {
		p.SetId(natportforwardId.(string))
	}
	if ipaddressid, ok := d.GetOk("ipaddressid"); ok {
		p.SetIpaddressid(ipaddressid.(string))
	}
	if keyword, ok := d.GetOk("keyword"); ok {
		p.SetKeyword(keyword.(string))
	}
	if networkid, ok := d.GetOk("networkid"); ok {
		p.SetNetworkid(networkid.(string))
	}

	r, err := cli.NatPortForward.ListPortForwardingRules(p)
	if err != nil {
		return fmt.Errorf("Error getting Port Forwarding Rule list: %s", err)
	}

	if r.PortForwardingRules == nil {
		return fmt.Errorf("Port Forwarding Rule not found")
	}
	v := r.PortForwardingRules[0]

	priEnd, _ := strconv.Atoi(v.Privateendport)
	pri, _ := strconv.Atoi(v.Privateport)
	pubEnd, _ := strconv.Atoi(v.Publicendport)
	pub, _ := strconv.Atoi(v.Publicport)

	d.Set("natportforward_id", v.Id)
	d.Set("ipaddressid", v.Ipaddressid)
	d.Set("networkid", v.Networkid)

	d.Set("ipaddress", v.Ipaddress)
	d.Set("privateendport", priEnd)
	d.Set("privateport", pri)
	d.Set("protocol", v.Protocol)
	d.Set("publicendport", pubEnd)
	d.Set("publicport", pub)
	d.Set("tags", flattenTags(v.Tags))
	d.Set("virtualmachineid", v.Virtualmachineid)
	d.Set("vmguestip", v.Vmguestip)

	d.SetId(v.Id)

	return nil
}
