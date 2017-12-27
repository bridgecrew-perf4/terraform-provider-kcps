package kcps

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	gk "github.com/uesyn/gokcps"
)

func resourceKcpsSnapshot() *schema.Resource {
	return &schema.Resource{
		Create: resourceKcpsSnapshotCreate,
		Read:   resourceKcpsSnapshotRead,
		Delete: resourceKcpsSnapshotDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"volumeid": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceKcpsSnapshotCreate(d *schema.ResourceData, meta interface{}) error {
	cli := meta.(*gk.KCPSClient)

	volumeid := d.Get("volumeid").(string)
	p := cli.Snapshot.NewCreateSnapshotParams(volumeid)

	r, err := cli.Snapshot.CreateSnapshot(p)
	if err != nil {
		return fmt.Errorf("Error creating Snapshot: %s", err)
	}
	d.SetId(r.Id)

	return resourceKcpsSnapshotRead(d, meta)
}

func resourceKcpsSnapshotRead(d *schema.ResourceData, meta interface{}) error {
	cli := meta.(*gk.KCPSClient)

	if d.Id() == "" {
		d.SetId("")
		return nil
	}

	p := cli.Snapshot.NewListSnapshotsParams()
	p.SetId(d.Id())
	r, err := cli.Snapshot.ListSnapshots(p)

	if err != nil {
		return fmt.Errorf("Error getting Snapshot list: %s", err)
	}

	if r.Snapshots == nil {
		d.SetId("")
		return nil
	}

	return nil
}

func resourceKcpsSnapshotDelete(d *schema.ResourceData, meta interface{}) error {
	cli := meta.(*gk.KCPSClient)

	p := cli.Snapshot.NewDeleteSnapshotParams(d.Id())
	_, err := cli.Snapshot.DeleteSnapshot(p)

	if err != nil {
		return fmt.Errorf("Error deleting Snapshot: %s", err)
	}

	d.SetId("")
	return nil
}
