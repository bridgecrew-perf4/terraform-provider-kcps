package kcps

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	gk "github.com/uesyn/gokcps"
)

func dataSourceKcpsNic() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceKcpsNicRead,

		Schema: map[string]*schema.Schema{
			"virtualmachineid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"nic_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"networkid": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},

			"ipaddress": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip6address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"macaddress": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"secondaryip": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"secondaryip_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"secondaryip_ipaddress": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

// enabled us to search Newtwork using 'networkid' (this isn't KCPS API's function)
func dataSourceKcpsNicRead(d *schema.ResourceData, meta interface{}) error {
	cli := meta.(*gk.KCPSClient)

	p := cli.Nic.NewListNicsParams(d.Get("virtualmachineid").(string))
	if nic_id, ok := d.GetOk("nic_id"); ok {
		p.SetNicid(nic_id.(string))
	}

	r, err := cli.Nic.ListNics(p)
	if err != nil {
		return fmt.Errorf("Error getting Nic list: %s", err)
	}

	if r.Nics == nil {
		return fmt.Errorf("Nic not found")
	}

	var v *gk.Nic

	if networkid, ok := d.GetOk("networkid"); ok {
		for _, n := range r.Nics {
			if n.Networkid == networkid.(string) {
				v = &n
				break
			}
		}
		if v == nil {
			return fmt.Errorf("Nic not found")
		}
	} else {
		v = &r.Nics[0]
	}

	d.Set("nic_id", v.Id)

	d.Set("ip6address", v.Ip6address)
	d.Set("ipaddress", v.Ipaddress)
	d.Set("macaddress", v.Macaddress)
	d.Set("networkid", v.Networkid)
	d.Set("secondaryip", flattenSecondaryIpNic(v))
	d.SetId(v.Id)

	return nil
}

func flattenSecondaryIpNic(nic *gk.Nic) []map[string]interface{} {
	result := make([]map[string]interface{}, 0, len(nic.Secondaryip))
	for _, s := range nic.Secondaryip {
		siMap := make(map[string]interface{})
		siMap["secondaryip_id"] = s.Id
		siMap["secondaryip_ipaddress"] = s.Ipaddress
		result = append(result, siMap)
	}
	return result
}
