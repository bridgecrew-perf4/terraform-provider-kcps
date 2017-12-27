package kcps

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	gk "github.com/uesyn/gokcps"
)

func dataSourceKcpsServiceOffering() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceKcpsServiceOfferingRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"serviceoffering_id": {
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

func dataSourceKcpsServiceOfferingRead(d *schema.ResourceData, meta interface{}) error {
	cli := meta.(*gk.KCPSClient)

	p := cli.AccountDomain.NewListServiceOfferingsParams()
	if name, ok := d.GetOk("name"); ok {
		p.SetName(name.(string))
	}
	if serviceoffering_id, ok := d.GetOk("serviceoffering_id"); ok {
		p.SetId(serviceoffering_id.(string))
	}
	if keyword, ok := d.GetOk("keyword"); ok {
		p.SetKeyword(keyword.(string))
	}
	r, err := cli.AccountDomain.ListServiceOfferings(p)

	if err != nil {
		return fmt.Errorf("Error getting Service Offering list: %s", err)
	}

	if r.ServiceOfferings == nil {
		return fmt.Errorf("Service Offering not found")
	}

	v := r.ServiceOfferings[0]

	d.Set("name", v.Name)
	d.Set("serviceoffering_id", v.Id)
	d.SetId(v.Id)

	return nil
}
