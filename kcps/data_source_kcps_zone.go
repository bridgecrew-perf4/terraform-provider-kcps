package kcps

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	gk "github.com/uesyn/gokcps"
)

func dataSourceKcpsZone() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceKcpsZoneRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"zone_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

// enabled us to search Newtwork using 'name' (this isn't KCPS API's function)
func dataSourceKcpsZoneRead(d *schema.ResourceData, meta interface{}) error {
	cli := meta.(*gk.KCPSClient)

	p := cli.AccountDomain.NewListZonesParams()
	r, err := cli.AccountDomain.ListZones(p)

	if err != nil {
		return fmt.Errorf("Error getting Zone list: %s", err)
	}

	var v *gk.Zone

	if name, ok := d.GetOk("name"); ok {
		for _, z := range r.Zones {
			if name.(string) == z.Name {
				v = z
				break
			}
		}
		if v == nil {
			return fmt.Errorf("Zone not found")
		}

	} else {
		v = r.Zones[0]
	}
	d.Set("zone_id", v.Id)
	d.Set("name", v.Name)
	d.SetId(v.Id)

	return nil
}
