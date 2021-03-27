package kcps

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	gk "github.com/uesyn/gokcps"
)

func dataSourceKcpsPublicIP() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceKcpsPublicIPRead,

		Schema: map[string]*schema.Schema{
			"associatednetworkid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"forloadbalancing": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"publicip_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipaddress": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"issourcenat": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"isstaticnat": {
				Type:     schema.TypeBool,
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
							Computed: true,
							Optional: true,
						},
						"value": {
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
						},
					},
				},
			},
			"zoneid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"networkid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"virtualmachineid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vmipaddress": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceKcpsPublicIPRead(d *schema.ResourceData, meta interface{}) error {
	cli := meta.(*gk.KCPSClient)

	p := cli.Nic.NewListPublicIpAddressesParams()

	if associatednetworkid, ok := d.GetOk("associatednetworkid"); ok {
		p.SetAssociatednetworkid(associatednetworkid.(string))
	}
	if forloadbalancing, ok := d.GetOk("forloadbalancing"); ok {
		p.SetForloadbalancing(forloadbalancing.(bool))
	}
	if publicipId, ok := d.GetOk("publicip_id"); ok {
		p.SetId(publicipId.(string))
	}
	if ipaddress, ok := d.GetOk("ipaddress"); ok {
		p.SetIpaddress(ipaddress.(string))
	}
	if issourcenat, ok := d.GetOk("issourcenat"); ok {
		p.SetIssourcenat(issourcenat.(bool))
	}
	if isstaticnat, ok := d.GetOk("isstaticnat"); ok {
		p.SetIsstaticnat(isstaticnat.(bool))
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
	if zoneid, ok := d.GetOk("zoneid"); ok {
		p.SetZoneid(zoneid.(string))
	}

	r, err := cli.Nic.ListPublicIpAddresses(p)

	if err != nil {
		return fmt.Errorf("Error getting Public IP list: %s", err)
	}
	if r.PublicIpAddresses == nil {
		return fmt.Errorf("Public IP not found")
	}

	v := r.PublicIpAddresses[0]

	d.Set("publicip_id", v.Id)
	d.Set("ipaddress", v.Ipaddress)
	d.Set("zoneid", v.Zoneid)
	d.Set("tags", flattenTags(v.Tags))
	d.Set("issourcenat", v.Issourcenat)
	d.Set("isstaticnat", v.Isstaticnat)

	d.Set("networkid", v.Networkid)
	d.Set("associatednetworkid", v.Associatednetworkid)
	d.Set("virtualmachineid", v.Virtualmachineid)
	d.Set("vmipaddress", v.Vmipaddress)
	d.SetId(v.Id)

	return nil
}
