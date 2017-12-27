package kcps

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	gk "github.com/uesyn/gokcps"
)

func dataSourceKcpsDiskOffering() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceKcpsDiskOfferingRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"diskoffering_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"keyword": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func dataSourceKcpsDiskOfferingRead(d *schema.ResourceData, meta interface{}) error {
	cli := meta.(*gk.KCPSClient)

	p := cli.AccountDomain.NewListDiskOfferingsParams()
	if name, ok := d.GetOk("name"); ok {
		p.SetName(name.(string))
	}
	if diskoffering_id, ok := d.GetOk("diskoffering_id"); ok {
		p.SetId(diskoffering_id.(string))
	}
	if keyword, ok := d.GetOk("keyword"); ok {
		p.SetKeyword(keyword.(string))
	}
	r, err := cli.AccountDomain.ListDiskOfferings(p)

	if err != nil {
		return fmt.Errorf("Error getting Disk Offering list: %s", err)
	}

	if r.DiskOfferings == nil {
		return fmt.Errorf("Disk Offering not found")
	}

	v := r.DiskOfferings[0]

	d.Set("name", v.Name)
	d.Set("diskoffering_id", v.Id)
	d.SetId(v.Id)
	return nil
}
