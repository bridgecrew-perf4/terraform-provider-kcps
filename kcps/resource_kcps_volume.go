package kcps

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	gk "github.com/uesyn/gokcps"
)

func resourceKcpsVolume() *schema.Resource {
	return &schema.Resource{
		Create: resourceKcpsVolumeCreate,
		Read:   resourceKcpsVolumeRead,
		Update: resourceKcpsVolumeUpdate,
		Delete: resourceKcpsVolumeDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"diskoffering": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						//create from Disk
						"diskofferingid": {
							Type:         schema.TypeString,
							Required:     true,
							ForceNew:     true,
							ValidateFunc: validateDiskOfferingId(),
						},
						//create from Disk
						"zoneid": {
							Type:         schema.TypeString,
							Required:     true,
							ForceNew:     true,
							ValidateFunc: validateZoneId(),
						},
						//create from Disk (this attribute is used for 'resizeVolume' too)
						"size": {
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"snapshot": {
				Type:          schema.TypeList,
				MaxItems:      1,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"diskoffering"},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						//create from Snapshot
						"snapshotid": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
					},
				},
			},

			//attachVolume
			"attachto": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceKcpsVolumeCreate(d *schema.ResourceData, meta interface{}) error {
	cli := meta.(*gk.KCPSClient)

	p := cli.Volume.NewCreateVolumeParams()
	name := d.Get("name").(string)
	p.SetName(name)

	prefixDisk := "diskoffering.0."
	if diskofferingid, ok := d.GetOk(prefixDisk + "diskofferingid"); ok {
		p.SetDiskofferingid(diskofferingid.(string))
	}
	if zoneid, ok := d.GetOk(prefixDisk + "zoneid"); ok {
		p.SetZoneid(zoneid.(string))
	}
	if size, ok := d.GetOk(prefixDisk + "size"); ok {
		sizeInt := size.(int)
		p.SetSize(int64(sizeInt))
	}

	prefixSnap := "snapshot.0."
	if snapshotid, ok := d.GetOk(prefixSnap + "snapshotid"); ok {
		p.SetSnapshotid(snapshotid.(string))
	}

	r, err := cli.Volume.CreateVolume(p)
	if err != nil {
		return fmt.Errorf("Error creating new Volume: %s", err)
	}
	d.SetId(r.Id)

	//attachVolume
	if attachto, ok := d.GetOk("attachto"); ok {
		p2 := cli.Volume.NewAttachVolumeParams(r.Id, attachto.(string))
		_, err := cli.Volume.AttachVolume(p2)
		if err != nil {
			return fmt.Errorf("Error attaching the Volume to the Virtual Machine: %s", err)
		}
	}
	return resourceKcpsVolumeRead(d, meta)
}

func resourceKcpsVolumeRead(d *schema.ResourceData, meta interface{}) error {
	cli := meta.(*gk.KCPSClient)

	if d.Id() == "" {
		d.SetId("")
		return nil
	}
	p := cli.Volume.NewListVolumesParams()
	p.SetId(d.Id())
	r, err := cli.Volume.ListVolumes(p)
	if err != nil {
		return fmt.Errorf("Error getting Volume list: %s", err)
	}
	if r.Volumes == nil {
		d.SetId("")
		return nil
	}

	d.Set("attachto", r.Volumes[0].Virtualmachineid)

	return nil
}

func resourceKcpsVolumeUpdate(d *schema.ResourceData, meta interface{}) error {
	cli := meta.(*gk.KCPSClient)

	if d.Id() == "" {
		d.SetId("")
		return nil
	}

	p := cli.Volume.NewListVolumesParams()
	p.SetId(d.Id())
	r, err := cli.Volume.ListVolumes(p)

	if err != nil {
		return fmt.Errorf("Error getting Volume list: %s", err)
	}
	if r.Volumes == nil {
		d.SetId("")
		return nil
	}

	// Enable partial state mode
	d.Partial(true)

	//'attachVolume' and 'detachVolume'
	if d.HasChange("attachto") {
		o, n := d.GetChange("attachto")

		//Does the attached old Volume exist?
		if o != "" {
			//detach old Volume
			p2 := cli.Volume.NewDetachVolumeParams()
			p2.SetId(d.Id())
			_, err := cli.Volume.DetachVolume(p2)
			if err != nil {
				return fmt.Errorf("Error detaching the Volume from the Virtual Machine: %s", err)
			}
		}

		//Does the New Volume exists?
		if n != "" {
			//attach new Volume
			p2 := cli.Volume.NewAttachVolumeParams(d.Id(), n.(string))
			_, err = cli.Volume.AttachVolume(p2)
			if err != nil {
				return fmt.Errorf("Error attaching the Volume to the Virtual Machine: %s", err)
			}
		}
		d.SetPartial("attachto")
	}

	//'resizeVolume'
	prefixDisk := "diskoffering.0."
	if d.HasChange(prefixDisk + "size") {
		if size, ok := d.GetOk(prefixDisk + "size"); ok {
			sizeInt := size.(int)
			p2 := cli.Volume.NewResizeVolumeParams(d.Id(), int64(sizeInt))
			_, err := cli.Volume.ResizeVolume(p2)
			if err != nil {
				return fmt.Errorf("Error resizing the Volume: %s", err)
			}
		}
		d.SetPartial(prefixDisk + "size")
	}

	d.Partial(false)

	return resourceKcpsVolumeRead(d, meta)
}

func resourceKcpsVolumeDelete(d *schema.ResourceData, meta interface{}) error {
	cli := meta.(*gk.KCPSClient)

	//if the volume attached to a VM, detach it before delete
	//(DEBUG) if attachto ="" -> not pass through here
	if _, ok := d.GetOk("attachto"); ok {
		p := cli.Volume.NewDetachVolumeParams()
		p.SetId(d.Id())
		_, err := cli.Volume.DetachVolume(p)
		if err != nil {
			return fmt.Errorf("Error detaching the Volume from the Virtual Machine: %s", err)
		}
	}

	p2 := cli.Volume.NewDeleteVolumeParams(d.Id())
	_, err := cli.Volume.DeleteVolume(p2)

	if err != nil {
		return fmt.Errorf("Error deleting Volume: %s", err)
	}

	d.SetId("")
	return nil
}
