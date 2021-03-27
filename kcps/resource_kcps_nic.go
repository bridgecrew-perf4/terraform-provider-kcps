package kcps

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	gk "github.com/uesyn/gokcps"
)

func resourceKcpsNic() *schema.Resource {
	return &schema.Resource{
		Create: resourceKcpsNicCreate,
		Read:   resourceKcpsNicRead,
		Update: resourceKcpsNicUpdate,
		Delete: resourceKcpsNicDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"networkid": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"virtualmachineid": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"secondaryip": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
		},
	}
}

func resourceKcpsNicCreate(d *schema.ResourceData, meta interface{}) error {
	cli := meta.(*gk.KCPSClient)
	var nicId string

	networkid := d.Get("networkid").(string)
	virtualmachineid := d.Get("virtualmachineid").(string)

	mutexKV.Lock("nic-" + networkid + virtualmachineid)
	defer mutexKV.Unlock("nic-" + networkid + virtualmachineid)

	p := cli.Nic.NewAddNicToVirtualMachineParams(networkid, virtualmachineid)
	r, err := cli.Nic.AddNicToVirtualMachine(p)

	if err != nil {
		return fmt.Errorf("Error adding Nic to Virtual Machine: %s", err)
	}

	//r.Id is VirtualMachine's ID ...
	for _, v := range r.Nic {
		if v.Networkid == networkid {
			nicId = v.Id
		}
	}
	d.SetId(nicId)

	// 'addIpToNic'
	if secondaryip, ok := d.GetOk("secondaryip"); ok {
		secondaryipSet := secondaryip.(*schema.Set)
		for _, v := range secondaryipSet.List() {
			p := cli.Nic.NewAddIpToNicParams(nicId)
			p.SetIpaddress(v.(string))
			_, err := cli.Nic.AddIpToNic(p)
			if err != nil {
				return fmt.Errorf("Error adding to IP address to Nic: %s", err)
			}
		}
	}

	return resourceKcpsNicRead(d, meta)
}

func resourceKcpsNicRead(d *schema.ResourceData, meta interface{}) error {
	cli := meta.(*gk.KCPSClient)

	if d.Id() == "" {
		d.SetId("")
		return nil
	}

	p := cli.Nic.NewListNicsParams(d.Get("virtualmachineid").(string))
	p.SetNicid(d.Id())
	r, err := cli.Nic.ListNics(p)

	if err != nil {
		return fmt.Errorf("Error getting Nic list: %s", err)
	}
	if r.Nics == nil {
		d.SetId("")
		return nil
	}

	if _, ok := d.GetOk("secondaryip"); ok {
		var secondaryipList []string
		for _, v := range r.Nics[0].Secondaryip {
			secondaryipList = append(secondaryipList, v.Ipaddress)
		}
		d.Set("secondaryip", convertStringArrToInterface(secondaryipList))
	}

	return nil
}

func resourceKcpsNicUpdate(d *schema.ResourceData, meta interface{}) error {
	cli := meta.(*gk.KCPSClient)

	if d.Id() == "" {
		d.SetId("")
		return nil
	}

	p := cli.Nic.NewListNicsParams(d.Get("virtualmachineid").(string))
	p.SetNicid(d.Id())
	r, err := cli.Nic.ListNics(p)

	if err != nil {
		return fmt.Errorf("Error getting Nic list: %s", err)
	}
	if r.Nics == nil {
		d.SetId("")
		return nil
	}

	if d.HasChange("secondaryip") {
		o, n := d.GetChange("secondaryip")

		// lead old and new secondaryip
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

		// remove
		for _, v := range oldList {
			if v != "" {
				ipId, err := getIpaddressId(
					d.Get("virtualmachineid").(string),
					d.Id(),
					v,
					cli,
				)
				if err != nil {
					return err
				}

				p := cli.Nic.NewRemoveIpFromNicParams(ipId)
				_, err = cli.Nic.RemoveIpFromNic(p)
				if err != nil {
					return fmt.Errorf("Error removing IP from the Nic: %s", err)
				}
			}
		}

		// add
		for _, v := range newList {
			if v != "" {
				p := cli.Nic.NewAddIpToNicParams(d.Id())
				p.SetIpaddress(v)
				_, err := cli.Nic.AddIpToNic(p)
				if err != nil {
					return fmt.Errorf("Error adding to IP address to the Nic: %s", err)
				}
			}
		}
	}

	return resourceKcpsNicRead(d, meta)
}

func resourceKcpsNicDelete(d *schema.ResourceData, meta interface{}) error {
	cli := meta.(*gk.KCPSClient)

	networkid := d.Get("networkid").(string)
	virtualmachineid := d.Get("virtualmachineid").(string)

	mutexKV.Lock("nic-" + networkid + virtualmachineid)
	defer mutexKV.Unlock("nic-" + networkid + virtualmachineid)

	p := cli.Nic.NewRemoveNicFromVirtualMachineParams(d.Id(), d.Get("virtualmachineid").(string))
	_, err := cli.Nic.RemoveNicFromVirtualMachine(p)

	if err != nil {
		return fmt.Errorf("Error removing Nic from Virtual Machine: %s", err)
	}

	d.SetId("")
	return nil
}

func expandSecondaryIpNicSimple(nic *gk.Nic) []string {
	var result []string
	for _, s := range nic.Secondaryip {
		result = append(result, s.Ipaddress)
	}
	return result
}

func getIpaddressId(virtualmachineid string, nicid string, ipaddress string, cli *gk.KCPSClient) (string, error) {
	p := cli.Nic.NewListNicsParams(virtualmachineid)
	p.SetNicid(nicid)
	r, err := cli.Nic.ListNics(p)

	if err != nil {
		return "", fmt.Errorf("Error getting Nic list: %s", err)
	}

	for _, ip := range r.Nics[0].Secondaryip {
		if ipaddress == ip.Ipaddress {
			return ip.Id, nil
		}
	}
	return "", fmt.Errorf("Error getting Nic ID")
}
