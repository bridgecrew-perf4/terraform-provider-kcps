package kcps

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	gk "github.com/uesyn/gokcps"
)

func dataSourceKcpsTemplate() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceKcpsTemplateRead,

		Schema: map[string]*schema.Schema{
			"templatefilter": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateTemplateFilter(),
			},
			"template_id": {
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
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateZoneId(),
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
			"ispublic": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"passwordenabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"ostypeid": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceKcpsTemplateRead(d *schema.ResourceData, meta interface{}) error {
	cli := meta.(*gk.KCPSClient)

	p := cli.Template.NewListTemplatesParams(d.Get("templatefilter").(string))
	if template_id, ok := d.GetOk("template_id"); ok {
		p.SetId(template_id.(string))
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

	p.SetName(d.Get("name").(string))
	p.SetZoneid(d.Get("zoneid").(string))
	r, err := cli.Template.ListTemplates(p)

	if err != nil {
		return fmt.Errorf("Error getting Zone list: %s", err)
	}

	if r.Templates == nil {
		return fmt.Errorf("Template not found")
	}

	v := r.Templates[0]

	d.Set("name", v.Name)
	d.Set("zoneid", v.Zoneid)

	d.Set("template_id", v.Id)
	d.Set("displaytext", v.Displaytext)
	d.Set("format", v.Format)
	d.Set("hypervisor", v.Hypervisor)
	d.Set("isdynamicallyscalable", v.Isdynamicallyscalable)
	d.Set("ispublic", v.Ispublic)
	d.Set("ostypeid", v.Ostypeid)
	d.Set("passwordenabled", v.Passwordenabled)
	d.Set("tags", flattenTags(v.Tags))

	d.SetId(v.Id)

	return nil
}
