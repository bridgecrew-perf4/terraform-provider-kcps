package kcps

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	gk "github.com/uesyn/gokcps"
)

func resourceKcpsVMSnapshot() *schema.Resource {
	return &schema.Resource{
		Create: resourceKcpsVMSnapshotCreate,
		Read:   resourceKcpsVMSnapshotRead,
		Delete: resourceKcpsVMSnapshotDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"virtualmachineid": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			/*
				"quiescevm": {
					Type:     schema.TypeBool,
					Optional: true,
					ForceNew: true,
				},
			*/
		},
	}
}

func resourceKcpsVMSnapshotCreate(d *schema.ResourceData, meta interface{}) error {
	cli := meta.(*gk.KCPSClient)

	virtualmachineid := d.Get("virtualmachineid").(string)
	p := cli.Snapshot.NewCreateVMSnapshotParams(virtualmachineid)

	if name, ok := d.GetOk("name"); ok {
		p.SetName(name.(string))
	}

	r, err := cli.Snapshot.CreateVMSnapshot(p)
	if err != nil {
		return fmt.Errorf("Error creating new VMSnapshot: %s", err)
	}

	d.SetId(r.Id)

	return resourceKcpsVMSnapshotRead(d, meta)
}

func resourceKcpsVMSnapshotRead(d *schema.ResourceData, meta interface{}) error {
	cli := meta.(*gk.KCPSClient)

	if d.Id() == "" {
		d.SetId("")
		return nil
	}
	p := cli.Snapshot.NewListVMSnapshotParams()
	p.SetVmsnapshotid(d.Id())

	r, err := cli.Snapshot.ListVMSnapshot(p)
	if err != nil {
		return fmt.Errorf("Error getting VMSnapshot list: %s", err)
	}
	if r.VMSnapshot == nil {
		d.SetId("")
		return nil
	}

	return nil
}

func resourceKcpsVMSnapshotDelete(d *schema.ResourceData, meta interface{}) error {
	cli := meta.(*gk.KCPSClient)

	p := cli.Snapshot.NewDeleteVMSnapshotParams(d.Id())
	_, err := cli.Snapshot.DeleteVMSnapshot(p)
	if err != nil {
		return fmt.Errorf("Error deleting VMSnapshot: %s", err)
	}

	d.SetId("")
	return nil
}
