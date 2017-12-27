package kcps

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	gk "github.com/uesyn/gokcps"
)

func dataSourceKcpsNetwork() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceKcpsNetworkRead,

		Schema: map[string]*schema.Schema{
			//enabled to search with the name of Network
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"network_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"zoneid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"keyword": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tags": {
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"value": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceKcpsNetworkRead(d *schema.ResourceData, meta interface{}) error {
	cli := meta.(*gk.KCPSClient)

	p := cli.AccountDomain.NewListNetworksParams()
	if zoneid, ok := d.GetOk("zoneid"); ok {
		p.SetZoneid(zoneid.(string))
	}
	if network_id, ok := d.GetOk("network_id"); ok {
		p.SetId(network_id.(string))
	}
	if keyword, ok := d.GetOk("keyword"); ok {
		p.SetKeyword(keyword.(string))
	}
	if tags, ok := d.GetOk("tags"); ok {
		for _, t := range tags.([]interface{}) {
			tagMap := map[string]string{}
			tMap := t.(map[string]interface{})
			tagMap[tMap["key"].(string)] = tMap["value"].(string)

			p.SetTags(tagMap)
		}
	}

	r, err := cli.AccountDomain.ListNetworks(p)

	if err != nil {
		return fmt.Errorf("Error getting Network list: %s", err)
	}
	if r.Networks == nil {
		return fmt.Errorf("Network not found")
	}

	var v *gk.Network

	if name, ok := d.GetOk("name"); ok {
		for _, n := range r.Networks {
			if n.Name == name.(string) {
				v = n
				break
			}
		}
		if v == nil {
			return fmt.Errorf("Network not found")
		}
	} else {
		v = r.Networks[0]
	}

	d.Set("name", v.Name)
	d.Set("network_id", v.Id)
	d.Set("zoneid", v.Zoneid)
	d.Set("tags", flattenTags(v.Tags))
	d.SetId(v.Id)

	return nil
}
