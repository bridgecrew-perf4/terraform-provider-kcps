package kcps

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	gk "github.com/uesyn/gokcps"
)

func dataSourceKcpsVMSnapshot() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceKcpsVMSnapshotRead,

		Schema: map[string]*schema.Schema{
			"vmsnapshot_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			//this parameter plays exactly the same role as request parameter 'name'
			//the description of the request parameter 'name' is "lists snapshot by snapshot name or display name"
			//but it correspondes 'displayname' only
			"displayname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"virtualmachineid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			//KCPS doesn't return this parameter (API reference mistake)
			/*
				"zoneid": {
					Type:     schema.TypeString,
					Computed: true,
				},
			*/
		},
	}
}

func dataSourceKcpsVMSnapshotRead(d *schema.ResourceData, meta interface{}) error {
	cli := meta.(*gk.KCPSClient)

	p := cli.Snapshot.NewListVMSnapshotParams()

	if vmsnapshotId, ok := d.GetOk("vmsnapshot_id"); ok {
		p.SetVmsnapshotid(vmsnapshotId.(string))
	}
	if name, ok := d.GetOk("name"); ok {
		p.SetName(name.(string))
	}
	if state, ok := d.GetOk("state"); ok {
		p.SetState(state.(string))
	}
	if virtualmachineid, ok := d.GetOk("virtualmachineid"); ok {
		p.SetVirtualmachineid(virtualmachineid.(string))
	}

	r, err := cli.Snapshot.ListVMSnapshot(p)
	if err != nil {
		return fmt.Errorf("Error getting VMSnapshot list: %s", err)
	}
	if r.VMSnapshot == nil {
		return fmt.Errorf("VMSnapshot not found")
	}

	v := r.VMSnapshot[0]

	d.Set("vmsnapshot_id", v.Id)
	d.Set("displayname", v.Displayname)
	d.Set("state", v.State)
	d.Set("virtualmachineid", v.Virtualmachineid)

	d.Set("zoneid", v.Zoneid)
	d.SetId(v.Id)

	return nil
}
