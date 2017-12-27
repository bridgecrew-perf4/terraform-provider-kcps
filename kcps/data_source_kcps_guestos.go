package kcps

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	gk "github.com/uesyn/gokcps"
)

func dataSourceKcpsGuestOS() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceKcpsGuestOSRead,

		Schema: map[string]*schema.Schema{
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"guestos_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"keyword": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"oscategoryid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

// enabled us to search Newtwork using 'name' (this isn't KCPS API's function)
func dataSourceKcpsGuestOSRead(d *schema.ResourceData, meta interface{}) error {
	cli := meta.(*gk.KCPSClient)

	p := cli.GuestOS.NewListOsTypesParams()
	if description, ok := d.GetOk("description"); ok {
		p.SetDescription(description.(string))
	}
	if guestos_id, ok := d.GetOk("guestos_id"); ok {
		p.SetId(guestos_id.(string))
	}
	if keyword, ok := d.GetOk("keyword"); ok {
		p.SetKeyword(keyword.(string))
	}
	if oscategoryid, ok := d.GetOk("oscategoryid"); ok {
		p.SetOscategoryid(oscategoryid.(string))
	}

	r, err := cli.GuestOS.ListOsTypes(p)

	if err != nil {
		return fmt.Errorf("Error getting Guest OS list: %s", err)
	}
	if r.OsTypes == nil {
		return fmt.Errorf("Guest OS not found")
	}

	v := r.OsTypes[0]

	d.Set("description", v.Description)
	d.Set("guestos_id", v.Id)
	d.Set("oscategoryid", v.Oscategoryid)
	d.SetId(v.Id)

	return nil
}
