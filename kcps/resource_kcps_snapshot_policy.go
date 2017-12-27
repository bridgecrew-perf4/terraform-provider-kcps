package kcps

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	gk "github.com/uesyn/gokcps"
)

func resourceKcpsSnapshotPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceKcpsSnapshotPolicyCreate,
		Read:   resourceKcpsSnapshotPolicyRead,
		Delete: resourceKcpsSnapshotPolicyDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"intervaltype": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateIntervalType(),
			},

			"maxsnaps": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},

			"schedule": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"timezone": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"volumeid": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceKcpsSnapshotPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	cli := meta.(*gk.KCPSClient)

	p := cli.Snapshot.NewCreateSnapshotPolicyParams(
		d.Get("intervaltype").(string),
		d.Get("maxsnaps").(int),
		d.Get("schedule").(string),
		d.Get("timezone").(string),
		d.Get("volumeid").(string),
	)

	r, err := cli.Snapshot.CreateSnapshotPolicy(p)

	if err != nil {
		return fmt.Errorf("Error creating new Snapshot Policy: %s", err)
	}
	d.SetId(r.Id)

	return resourceKcpsSnapshotPolicyRead(d, meta)
}

func resourceKcpsSnapshotPolicyRead(d *schema.ResourceData, meta interface{}) error {
	cli := meta.(*gk.KCPSClient)

	if d.Id() == "" {
		d.SetId("")
		return nil
	}

	p := cli.Snapshot.NewListSnapshotPoliciesParams()
	p.SetVolumeid(d.Get("volumeid").(string))
	r, err := cli.Snapshot.ListSnapshotPolicies(p)
	if err != nil {
		return fmt.Errorf("Error getting Snapshot Policy list: %s", err)
	}
	if r.SnapshotPolicies == nil {
		d.SetId("")
		return nil
	}

	//search with the ID of SnapshotShotPolicy
	find := false
	for _, s := range r.SnapshotPolicies {
		if s.Id == d.Id() {
			find = true
		}
	}
	if find == false {
		d.SetId("")
		return nil
	}

	return nil
}

func resourceKcpsSnapshotPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	cli := meta.(*gk.KCPSClient)

	p := cli.Snapshot.NewDeleteSnapshotPoliciesParams()
	p.SetId(d.Id())
	_, err := cli.Snapshot.DeleteSnapshotPolicies(p)
	if err != nil {
		return fmt.Errorf("Error deleting the Snapshot Policy: %s", err)
	}

	d.SetId("")
	return nil
}
