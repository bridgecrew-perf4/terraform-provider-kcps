package kcps

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	gk "github.com/uesyn/gokcps"
)

func dataSourceKcpsSnapshot() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceKcpsSnapshotRead,

		Schema: map[string]*schema.Schema{
			"snapshot_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"intervaltype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"keyword": {
				Type:     schema.TypeString,
				Optional: true,
			},
			//KCPS doesn't return this parameter (API reference mistake)
			"zoneid": {
				Type:     schema.TypeString,
				Optional: true,
				//Computed: true,
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
			"volumeid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func dataSourceKcpsSnapshotRead(d *schema.ResourceData, meta interface{}) error {
	cli := meta.(*gk.KCPSClient)

	p := cli.Snapshot.NewListSnapshotsParams()

	if snapshotId, ok := d.GetOk("snapshot_id"); ok {
		p.SetId(snapshotId.(string))
	}
	if intervaltype, ok := d.GetOk("intervaltype"); ok {
		p.SetIntervaltype(intervaltype.(string))
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
	if volumeid, ok := d.GetOk("volumeid"); ok {
		p.SetVolumeid(volumeid.(string))
	}
	if zoneid, ok := d.GetOk("zoneid"); ok {
		p.SetZoneid(zoneid.(string))
	}

	r, err := cli.Snapshot.ListSnapshots(p)

	if err != nil {
		return fmt.Errorf("Error getting Snapshot list: %s", err)
	}
	if r.Snapshots == nil {
		return fmt.Errorf("Snapshot not found")
	}

	v := r.Snapshots[0]

	d.Set("snapshot_id", v.Id)
	d.Set("intervaltype", v.Intervaltype)
	d.Set("name", v.Name)
	d.Set("tags", flattenTags(v.Tags))

	d.Set("volumeid", v.Volumeid)
	d.SetId(v.Id)

	return nil
}
