package kcps

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	gk "github.com/uesyn/gokcps"
)

func dataSourceKcpsValueVM() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceKcpsValueVMRead,

		Schema: map[string]*schema.Schema{
			"valuevm_id": {
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
				Optional: true,
				Computed: true,
			},
			"networkid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
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
			"templateid": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"zoneid": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},

			"diskofferingid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hypervisor": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"publicip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"publicipid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"serviceofferingid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"isoid": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceKcpsValueVMRead(d *schema.ResourceData, meta interface{}) error {
	cli := meta.(*gk.KCPSClient)

	p := cli.VirtualMachine.NewListVirtualMachinesParams()
	if valuevm_id, ok := d.GetOk("valuevm_id"); ok {
		p.SetId(valuevm_id.(string))
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
	if state, ok := d.GetOk("state"); ok {
		p.SetState(state.(string))
	}
	if tags, ok := d.GetOk("tags"); ok {
		for _, t := range tags.([]interface{}) {
			tagMap := map[string]string{}
			tMap := t.(map[string]interface{})
			tagMap[tMap["key"].(string)] = tMap["value"].(string)

			p.SetTags(tagMap)
		}
	}
	if templateid, ok := d.GetOk("templateid"); ok {
		p.SetTemplateid(templateid.(string))
	}
	if zoneid, ok := d.GetOk("zoneid"); ok {
		p.SetZoneid(zoneid.(string))
	}

	r, err := cli.VirtualMachine.ListVirtualMachines(p)
	if err != nil {
		return fmt.Errorf("Error getting Value Virtual Machines list: %s", err)
	}

	if r.VirtualMachines == nil {
		return fmt.Errorf("Viratul Machine Not found")
	}

	v := r.VirtualMachines[0]

	d.Set("valuevm_id", v.Id)
	d.Set("name", v.Name)
	d.Set("tags", flattenTags(v.Tags))
	d.Set("templateid", v.Templateid)
	d.Set("zoneid", v.Zoneid)

	d.Set("diskofferingid", v.Diskofferingid)
	d.Set("hypervisor", v.Hypervisor)
	d.Set("publicip", v.Publicip)
	d.Set("publicipid", v.Publicipid)
	d.Set("serviceofferingid", v.Serviceofferingid)
	d.Set("isoid", v.Isoid)

	d.SetId(v.Id)

	return nil

}
