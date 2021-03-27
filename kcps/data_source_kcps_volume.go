package kcps

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	gk "github.com/uesyn/gokcps"
)

func dataSourceKcpsVolume() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceKcpsVolumeRead,

		Schema: map[string]*schema.Schema{
			"volume_id": {
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
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"virtualmachineid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"zoneid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
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

			"diskofferingid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"serviceofferingid": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceKcpsVolumeRead(d *schema.ResourceData, meta interface{}) error {
	cli := meta.(*gk.KCPSClient)

	p := cli.Volume.NewListVolumesParams()

	if volume_id, ok := d.GetOk("volume_id"); ok {
		p.SetId(volume_id.(string))
	}
	if keyword, ok := d.GetOk("keyword"); ok {
		p.SetKeyword(keyword.(string))
	}
	if name, ok := d.GetOk("name"); ok {
		p.SetName(name.(string))
	}
	//go can't use "type" on variable name
	if type_, ok := d.GetOk("type"); ok {
		p.SetType(type_.(string))
	}
	if virtualmachineid, ok := d.GetOk("virtualmachineid"); ok {
		p.SetVirtualmachineid(virtualmachineid.(string))
	}
	if zoneid, ok := d.GetOk("zoneid"); ok {
		p.SetZoneid(zoneid.(string))
	}
	if tags, ok := d.GetOk("tags"); ok {
		for _, t := range tags.([]interface{}) {
			tagMap := map[string]string{}
			tMap := t.(map[string]interface{})
			tagMap[tMap["key"].(string)] = tMap["value"].(string)

			p.SetTags(tagMap)
		}
	}

	r, err := cli.Volume.ListVolumes(p)
	if err != nil {
		return fmt.Errorf("Error getting Volume list: %s", err)
	}

	if r.Volumes == nil {
		return fmt.Errorf("Volume Not found")
	}

	v := r.Volumes[0]

	d.Set("volume_id", v.Id)
	d.Set("name", v.Name)
	d.Set("type", v.Type)
	d.Set("virtualmachineid", v.Virtualmachineid)
	d.Set("zoneid", v.Zoneid)
	d.Set("tags", flattenTags(v.Tags))

	d.Set("diskofferingid", v.Diskofferingid)
	d.Set("serviceofferingid", v.Serviceofferingid)

	d.SetId(v.Id)

	return nil

}
