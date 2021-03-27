package kcps

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	gk "github.com/uesyn/gokcps"
)

func resourceKcpsISO() *schema.Resource {
	return &schema.Resource{
		Create: resourceKcpsISOCreate,
		Read:   resourceKcpsISORead,
		Update: resourceKcpsISOUpdate,
		Delete: resourceKcpsISODelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{

			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"displaytext": {
				Type:     schema.TypeString,
				Required: true,
			},
			"url": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"zoneid": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateZoneId(),
			},
			"ostypeid": {
				Type:     schema.TypeString,
				Required: true,
			},

			//'attachIso'
			"attachto": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},

			//'updateIsoPermissions'
			/*
				"ispublic": {
					Type:     schema.TypeBool,
					Optional: true,
				},
			*/
		},
	}
}

func resourceKcpsISOCreate(d *schema.ResourceData, meta interface{}) error {
	cli := meta.(*gk.KCPSClient)

	name := d.Get("name").(string)
	displaytext := d.Get("displaytext").(string)
	url := d.Get("url").(string)
	zoneid := d.Get("zoneid").(string)
	ostypeid := d.Get("ostypeid").(string)

	p := cli.ISO.NewRegisterIsoParams(
		displaytext,
		name,
		url,
		zoneid,
		ostypeid,
	)

	/*******************DELETE*************/
	p.SetOstypeid(ostypeid)
	/**************************************/

	r, err := cli.ISO.RegisterIso(p)
	if err != nil {
		return fmt.Errorf("Error registering new Template: %s", err)
	}

	log.Println("Hello: ", r)
	log.Println("Hello id: ", r.Id)
	d.SetId(r.Id)

	//'attachIso'
	if attachto, ok := d.GetOk("attachto"); ok {
		attachtoSet := attachto.(*schema.Set)
		for _, v := range attachtoSet.List() {
			p2 := cli.ISO.NewAttachIsoParams(r.Id, v.(string))
			_, err := cli.ISO.AttachIso(p2)
			if err != nil {
				return fmt.Errorf("Error attaching the ISO to the Virtual Machine: %s", err)
			}
		}
	}

	return resourceKcpsISORead(d, meta)
}

func resourceKcpsISORead(d *schema.ResourceData, meta interface{}) error {
	cli := meta.(*gk.KCPSClient)

	if d.Id() == "" {
		d.SetId("")
		return nil
	}
	p := cli.ISO.NewListIsosParams()
	p.SetId(d.Id())

	r, err := cli.ISO.ListIsos(p)
	if err != nil {
		return fmt.Errorf("Error getting ISO list: %s", err)
	}
	if r.Isos == nil {
		d.SetId("")
		return nil
	}

	d.Set("name", r.Isos[0].Name)
	d.Set("displaytext", r.Isos[0].Displaytext)
	d.Set("ostypeid", r.Isos[0].Ostypeid)

	// 'listIsos' can't get attaching VirtuakMachineID, so use 'listVirtualMachines'
	if _, ok := d.GetOk("attachto"); ok {
		p := cli.VirtualMachine.NewListVirtualMachinesParams()
		r, err := cli.VirtualMachine.ListVirtualMachines(p)
		if err != nil {
			return fmt.Errorf("Error getting Value Virtual Machines list: %s", err)
		}

		// get all attaching virtualmachineid
		var attachtoList []string
		for _, v := range r.VirtualMachines {
			if v.Isoid == d.Id() {
				attachtoList = append(attachtoList, v.Id)
			}
		}
		d.Set("attachto", convertStringArrToInterface(attachtoList))
	}

	return nil
}

func resourceKcpsISOUpdate(d *schema.ResourceData, meta interface{}) error {
	cli := meta.(*gk.KCPSClient)

	if d.Id() == "" {
		d.SetId("")
		return nil
	}

	p := cli.ISO.NewListIsosParams()
	p.SetId(d.Id())
	r, err := cli.ISO.ListIsos(p)

	if err != nil {
		return fmt.Errorf("Error getting ISO list: %s", err)
	}
	if r.Isos == nil {
		d.SetId("")
		return nil
	}

	d.Partial(true)

	//'updateIso'
	if d.HasChange("name") || d.HasChange("displaytext") || d.HasChange("ostypeid") {

		p := cli.ISO.NewUpdateIsoParams(d.Id())

		if name, ok := d.GetOk("name"); ok {
			p.SetName(name.(string))
		}
		if displaytext, ok := d.GetOk("displaytext"); ok {
			p.SetDisplaytext(displaytext.(string))
		}
		if ostypeid, ok := d.GetOk("ostypeid"); ok {
			p.SetOstypeid(ostypeid.(string))
		}

		_, err := cli.ISO.UpdateIso(p)
		if err != nil {
			return fmt.Errorf("Error updating the ISO: %s", err)
		}

		if _, ok := d.GetOk("name"); ok {
			d.SetPartial("name")
		}
		if _, ok := d.GetOk("displaytext"); ok {
			d.SetPartial("displaytext")
		}
		if _, ok := d.GetOk("ostypeid"); ok {
			d.SetPartial("ostypeid")
		}
	}

	//'attchIso' and 'detachIso'
	if d.HasChange("attachto") {
		o, n := d.GetChange("attachto")

		// lead old and new attach points
		// nothing to do for VM exists new and old
		olds, news := o.(*schema.Set), n.(*schema.Set)
		oldList, newList := make([]string, olds.Len()), make([]string, news.Len())
		for i, v := range olds.List() {
			oldList[i] = v.(string)
		}
		for i, v := range news.List() {
			newList[i] = v.(string)
		}
		for i, v := range oldList {
			for j, w := range newList {
				if v == w {
					oldList[i] = ""
					newList[j] = ""
				}
			}
		}

		// detach ISO
		for _, v := range oldList {
			if v != "" {
				p := cli.ISO.NewDetachIsoParams(v)
				_, err := cli.ISO.DetachIso(p)
				if err != nil {
					return fmt.Errorf("Error detaching the ISO from the Virtual Machine: %s", err)
				}
			}
		}

		// attach ISO
		for _, v := range newList {
			if v != "" {
				p := cli.ISO.NewAttachIsoParams(d.Id(), v)
				_, err := cli.ISO.AttachIso(p)
				if err != nil {
					return fmt.Errorf("Error attaching the ISO to the Virtual Machine: %s", err)
				}
			}
		}

		d.SetPartial("attachto")
	}

	d.Partial(false)

	return resourceKcpsISORead(d, meta)
}

func resourceKcpsISODelete(d *schema.ResourceData, meta interface{}) error {
	cli := meta.(*gk.KCPSClient)

	//if the ISO attached to VMs, detach them before delete
	if attachto, ok := d.GetOk("attachto"); ok {
		attachtoSet := attachto.(*schema.Set)
		for _, v := range attachtoSet.List() {
			p := cli.ISO.NewDetachIsoParams(v.(string))
			_, err := cli.ISO.DetachIso(p)
			if err != nil {
				return fmt.Errorf("Error detaching the ISO from the Virtual Machine: %s", err)
			}
		}
	}

	p := cli.ISO.NewDeleteIsoParams(d.Id())
	_, err := cli.ISO.DeleteIso(p)
	if err != nil {
		return fmt.Errorf("Error deleting ISO: %s", err)
	}

	d.SetId("")
	return nil
}
