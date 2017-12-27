package kcps

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	gk "github.com/uesyn/gokcps"
)

func dataSourceKcpsISO() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceKcpsISORead,

		Schema: map[string]*schema.Schema{
			"iso_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bootable": {
				Type:     schema.TypeBool,
				Computed: true,
				Optional: true,
			},
			"isofilter": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"isready": {
				Type:     schema.TypeBool,
				Computed: true,
				Optional: true,
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

			"displaytext": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"format": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hypervisor": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"isdynamicallyscalable": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"isextractable": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"ostypeid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"passwordenabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"ispublic": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func dataSourceKcpsISORead(d *schema.ResourceData, meta interface{}) error {
	cli := meta.(*gk.KCPSClient)

	p := cli.ISO.NewListIsosParams()
	if isoId, ok := d.GetOk("iso_id"); ok {
		p.SetId(isoId.(string))
	}
	if bootable, ok := d.GetOk("bootable"); ok {
		p.SetBootable(bootable.(bool))
	}
	if isofilter, ok := d.GetOk("isofilter"); ok {
		p.SetIsofilter(isofilter.(string))
	}
	if isready, ok := d.GetOk("isready"); ok {
		p.SetIsready(isready.(bool))
	}
	if keyword, ok := d.GetOk("keyword"); ok {
		p.SetKeyword(keyword.(string))
	}
	if name, ok := d.GetOk("name"); ok {
		p.SetName(name.(string))
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

	r, err := cli.ISO.ListIsos(p)
	if err != nil {
		return fmt.Errorf("Error getting ISO list: %s", err)
	}

	if r.Isos == nil {
		return fmt.Errorf("ISO not found")
	}

	v := r.Isos[0]

	d.Set("iso_id", v.Id)
	d.Set("bootable", v.Bootable)
	d.Set("isready", v.Isready)
	d.Set("name", v.Name)
	d.Set("tags", flattenTags(v.Tags))
	d.Set("zoneid", v.Zoneid)

	d.Set("displaytext", v.Displaytext)
	d.Set("format", v.Format)
	d.Set("hypervisor", v.Hypervisor)
	d.Set("isdynamicallyscalable", v.Isdynamicallyscalable)
	d.Set("isextractable", v.Isextractable)
	d.Set("ostypeid", v.Ostypeid)
	d.Set("passwordenabled", v.Passwordenabled)
	d.Set("ispublic", v.Ispublic)

	d.SetId(v.Id)

	return nil
}
