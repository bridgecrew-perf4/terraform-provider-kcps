package kcps

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	gk "github.com/uesyn/gokcps"
)

func dataSourceKcpsSnapshotPolicy() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceKcpsSnapshotPolicyRead,

		Schema: map[string]*schema.Schema{
			"volumeid": {
				Type:     schema.TypeString,
				Required: true,
			},
			//enabled to search with the ID of SnapshotPolicy
			"snapshotpolicy_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"intervaltype": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"maxsnaps": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"schedule": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"timezone": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceKcpsSnapshotPolicyRead(d *schema.ResourceData, meta interface{}) error {
	cli := meta.(*gk.KCPSClient)

	p := cli.Snapshot.NewListSnapshotPoliciesParams()

	if volumeid, ok := d.GetOk("volumeid"); ok {
		p.SetVolumeid(volumeid.(string))
	}

	r, err := cli.Snapshot.ListSnapshotPolicies(p)

	if err != nil {
		return fmt.Errorf("Error getting Snapshot Policy list: %s", err)
	}
	if r.SnapshotPolicies == nil {
		return fmt.Errorf("Snapshot Policy not found")
	}

	var v *gk.SnapshotPolicy
	//if snapshotpolicy_id is set, searchwith it. if else return the value of SnapshotPolicies[0]
	if snapshotpolicy_id, ok := d.GetOk("snapshotpolicy_id"); ok {
		for _, s := range r.SnapshotPolicies {
			if s.Id == snapshotpolicy_id.(string) {
				v = s
			}
		}
		if v == nil {
			return fmt.Errorf("Snapshot Policy not found")
		}

	} else {
		v = r.SnapshotPolicies[0]
	}

	d.Set("snapshotpolicy_id", v.Id)
	d.Set("intervaltype", convertIntervalType(v.Intervaltype))
	d.Set("maxsnaps", v.Maxsnaps)
	d.Set("schedule", v.Schedule)
	d.Set("timezone", v.Timezone)

	d.SetId(v.Id)

	return nil
}

// API return "intervaltype id" like "2", but we need "intervaltype" like "WEEKLY"!!!!
func convertIntervalType(num int) string {
	if num == 1 {
		return "DAILY"
	} else if num == 2 {
		return "WEEKLY"
	} else if num == 3 {
		return "MONTHLY"
	}
	return "NOT SUPORTED"
}
