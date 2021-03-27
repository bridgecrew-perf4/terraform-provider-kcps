package kcps

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	gk "github.com/uesyn/gokcps"
)

func dataSourceKcpsFirewall() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceKcpsFirewallRead,

		Schema: map[string]*schema.Schema{
			"firewall_id": {
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

			"cidrlist": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"startport": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"endport": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"icmpcode": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"icmptype": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ipaddress": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"protocol": {
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
		},
	}
}

func dataSourceKcpsFirewallRead(d *schema.ResourceData, meta interface{}) error {
	cli := meta.(*gk.KCPSClient)

	p := cli.Firewall.NewListFirewallRulesParams()
	if firewall_id, ok := d.GetOk("firewall_id"); ok {
		p.SetId(firewall_id.(string))
	}
	if ipaddressid, ok := d.GetOk("ipaddressid"); ok {
		p.SetIpaddressid(ipaddressid.(string))
	}
	if keyword, ok := d.GetOk("keyword"); ok {
		p.SetId(keyword.(string))
	}
	if networkid, ok := d.GetOk("networkid"); ok {
		p.SetNetworkid(networkid.(string))
	}

	r, err := cli.Firewall.ListFirewallRules(p)
	if err != nil {
		return fmt.Errorf("Error getting Firewall Rule list: %s", err)
	}

	if r.FirewallRules == nil {
		return fmt.Errorf("Firewall Rule not found")
	}
	v := r.FirewallRules[0]

	d.Set("firewall_id", v.Id)
	d.Set("ipaddressid", v.Ipaddressid)
	d.Set("networkid", v.Networkid)

	d.Set("cidrlist", expandCidrFirewall(v))
	d.Set("endport", v.Endport)
	d.Set("icmpcode", v.Icmpcode)
	d.Set("icmptype", v.Icmptype)
	d.Set("ipaddress", v.Ipaddress)
	d.Set("protocol", v.Protocol)
	d.Set("startport", v.Startport)
	d.Set("tags", flattenTags(v.Tags))

	d.SetId(v.Id)

	return nil
}

//'cidrlist' of the respponse is string. not []string.
func expandCidrFirewall(fw *gk.FirewallRule) []string {
	var result []string
	clist := strings.Split(fw.Cidrlist, ",")
	for _, c := range clist {
		result = append(result, c)
	}
	return result
}
