package kcps

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	gk "github.com/uesyn/gokcps"
)

func resourceKcpsPublicIP() *schema.Resource {
	return &schema.Resource{
		Create: resourceKcpsPublicIPCreate,
		Read:   resourceKcpsPublicIPRead,
		Update: resourceKcpsPublicIPUpdate,
		Delete: resourceKcpsPublicIPDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"networkid": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			//enableStaticNat
			"staticnat": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"virtualmachineid": {
							Type:     schema.TypeString,
							Required: true,
						},
						"vmguestip": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},

			"ipaddress": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceKcpsPublicIPCreate(d *schema.ResourceData, meta interface{}) error {
	cli := meta.(*gk.KCPSClient)

	networkid := d.Get("networkid").(string)
	mutexKV.Lock("publicip-" + networkid)
	defer mutexKV.Unlock("publicip-" + networkid)

	p := cli.Nic.NewAssociateIpAddressParams(networkid)
	r, err := cli.Nic.AssociateIpAddress(p)

	if err != nil {
		return fmt.Errorf("Error associatingg PublicIP: %s", err)
	}

	//Get Public IP address you just created
	p2 := cli.Nic.NewListPublicIpAddressesParams()
	p2.SetId(r.Id)
	r2, err := cli.Nic.ListPublicIpAddresses(p2)
	if err != nil {
		return fmt.Errorf("Error getting Public IP list: %s", err)
	}
	d.SetId(r.Id)
	d.Set("ipaddress", r2.PublicIpAddresses[0].Ipaddress)

	//enableStaticNat
	prefix := "staticnat.0."
	if virtualmachineid, ok := d.GetOk(prefix + "virtualmachineid"); ok {
		p3 := cli.Firewall.NewEnableStaticNatParams(d.Id(), virtualmachineid.(string))

		if vmguestip, ok := d.GetOk(prefix + "vmguestip"); ok {
			p3.SetVmguestip(vmguestip.(string))
		}

		_, err := cli.Firewall.EnableStaticNat(p3)
		if err != nil {
			return fmt.Errorf("Error enabling Static NAT: %s", err)
		}
	}

	return resourceKcpsPublicIPRead(d, meta)
}

func resourceKcpsPublicIPRead(d *schema.ResourceData, meta interface{}) error {
	cli := meta.(*gk.KCPSClient)

	if d.Id() == "" {
		d.SetId("")
		return nil
	}

	p := cli.Nic.NewListPublicIpAddressesParams()
	p.SetId(d.Id())
	r, err := cli.Nic.ListPublicIpAddresses(p)

	if err != nil {
		return fmt.Errorf("Error getting Public IP list: %s", err)
	}

	if r.PublicIpAddresses == nil {
		d.SetId("")
		return nil
	}

	d.Set("staticnat_vmid", r.PublicIpAddresses[0].Virtualmachineid)

	return nil
}

func resourceKcpsPublicIPUpdate(d *schema.ResourceData, meta interface{}) error {
	cli := meta.(*gk.KCPSClient)

	if d.Id() == "" {
		d.SetId("")
		return nil
	}

	p := cli.Nic.NewListPublicIpAddressesParams()
	p.SetId(d.Id())
	r, err := cli.Nic.ListPublicIpAddresses(p)

	if err != nil {
		return fmt.Errorf("Error getting Public IP list: %s", err)
	}
	if r.PublicIpAddresses == nil {
		d.SetId("")
		return nil
	}

	//disable and enable Static NAT
	prefix := "staticnat.0."
	if d.HasChange(prefix+"virtualmachineid") || d.HasChange(prefix+"vmguestip") {
		o, n := d.GetChange(prefix + "virtualmachineid")
		_, n2 := d.GetChange(prefix + "vmguestip")

		//Does the enabled old Static NAT exist?
		if o != "" {
			//disable old Static NAT
			p2 := cli.Firewall.NewDisableStaticNatParams(d.Id())
			_, err := cli.Firewall.DisableStaticNat(p2)
			if err != nil {
				return fmt.Errorf("Error disabling Static NAT: %s", err)
			}
		}

		//Does enalbe new Static NAT ?
		if n != "" {
			//enable new Static NAT
			p2 := cli.Firewall.NewEnableStaticNatParams(d.Id(), n.(string))

			if n2 != "" {
				p2.SetVmguestip(n2.(string))
			}

			_, err := cli.Firewall.EnableStaticNat(p2)
			if err != nil {
				return fmt.Errorf("Error enabling Static NAT: %s", err)
			}
		}

	}

	return resourceKcpsPublicIPRead(d, meta)
}

func resourceKcpsPublicIPDelete(d *schema.ResourceData, meta interface{}) error {
	cli := meta.(*gk.KCPSClient)

	networkid := d.Get("networkid").(string)
	mutexKV.Lock("publicip-" + networkid)
	defer mutexKV.Unlock("publicip-" + networkid)

	p := cli.Nic.NewDisassociateIpAddressParams(d.Id())
	_, err := cli.Nic.DisassociateIpAddress(p)

	if err != nil {
		return fmt.Errorf("Error disassociating Public IP: %s", err)
	}

	d.SetId("")
	return nil
}
