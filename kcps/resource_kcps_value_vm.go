package kcps

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	gk "github.com/uesyn/gokcps"
)

func resourceKcpsValueVM() *schema.Resource {
	return &schema.Resource{
		Create: resourceKcpsValueVMCreate,
		Read:   resourceKcpsValueVMRead,
		Update: resourceKcpsValueVMUpdate,
		Delete: resourceKcpsValueVMDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"serviceofferingid": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"templateid": {
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
			/* if gokcps is able to set other of VMware
			"hypervisor": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateHypervisor(),
			},
			*/
			//not support
			"iptonetworklist": {
				Type:     schema.TypeList,
				Optional: true, //Required: true (if gokcps doesn't set PublicFrontSegment as default)
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"networkid": {
							Type:     schema.TypeString,
							Required: true,
						},
						"ip": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"ipv6": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},

			"diskoffering": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"diskofferingid": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"size": {
							Type:     schema.TypeInt,
							Required: true,
							ForceNew: true,
						},
					},
				},
			},

			"publicip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"password": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceKcpsValueVMCreate(d *schema.ResourceData, meta interface{}) error {
	cli := meta.(*gk.KCPSClient)

	zoneid := d.Get("zoneid").(string)
	name := d.Get("name").(string)

	mutexKV.Lock("value-vm-" + zoneid + name)
	defer mutexKV.Unlock("value-vm-" + zoneid + name)

	p := cli.VirtualMachine.NewDeployValueVirtualMachineParams(
		d.Get("serviceofferingid").(string),
		d.Get("templateid").(string),
		zoneid,
		name,
	)

	//set 'iptonetworkslist' to param
	//If 'iptonetworklist' is not set, it is connected to PublicForwardSegment by gokcps
	var npls []gk.IptoNetworklistParams
	if v, ok := d.GetOk("iptonetworklist"); ok {
		for _, e := range v.([]interface{}) {
			conve := e.(map[string]interface{})
			networkid := conve["networkid"].(string)
			ip := conve["ip"].(string)
			ipv6 := conve["ipv6"].(string)

			np := cli.VirtualMachine.NewIptoNetworklistParams(networkid)
			np.SetIpv4(ip)
			np.SetIpv6(ipv6)
			npls = append(npls, np)
			p.SetIptoNetworklist(npls)
		}
	}

	//use diskoffering
	prefix := "diskoffering.0."
	if diskofferingid, ok := d.GetOk(prefix + "diskofferingid"); ok {
		p.SetDiskofferingid(diskofferingid.(string))
	}
	if size, ok := d.GetOk(prefix + "size"); ok {
		sizeInt := size.(int)
		p.SetSize(int64(sizeInt))
	}

	r, err := cli.VirtualMachine.DeployValueVirtualMachine(p)

	if err != nil {
		return fmt.Errorf("Error creating new Value Virtual Machine: %s", err)
	}
	d.SetId(r.Id)

	// if r.Publicip is nil, set the IP of PublicFrontSegment)
	p2 := cli.Nic.NewListPublicIpAddressesParams()
	p2.SetIssourcenat(true)
	r2, err := cli.Nic.ListPublicIpAddresses(p2)
	if err != nil {
		return fmt.Errorf("Error getting Public IP list: %s", err)
	}
	natIp := r2.PublicIpAddresses[0].Ipaddress

	pubIp := r.Publicip
	if pubIp == "" {
		pubIp = natIp
	}
	d.Set("publicip", pubIp)
	d.Set("password", r.Password)

	return resourceKcpsValueVMRead(d, meta)
}

func resourceKcpsValueVMRead(d *schema.ResourceData, meta interface{}) error {
	cli := meta.(*gk.KCPSClient)

	if d.Id() == "" {
		d.SetId("")
		return nil
	}
	p := cli.VirtualMachine.NewListVirtualMachinesParams()
	p.SetId(d.Id())
	r, err := cli.VirtualMachine.ListVirtualMachines(p)
	if err != nil {
		return fmt.Errorf("Error getting Value Virtual Machines list: %s", err)
	}
	if r.VirtualMachines == nil {
		d.SetId("")
		return nil
	}

	d.Set("iptonetworklist", flattenIptoNetworks(r.VirtualMachines[0]))

	//d.Set("hypervisor", r.VirtualMachines[0].Hypervisor)

	return nil
}

//iptonetworklistのUpdate未実装。NicのAPI叩く必要あり
func resourceKcpsValueVMUpdate(d *schema.ResourceData, meta interface{}) error {
	// Enable partial state mode
	//d.Partial(true)

	/*
		if d.HasChange("iptonetworklist") {
			lvmp := cli.VirtualMachine.NewListVirtualMachinesParams()
			idls := strings.Split(d.Id(), "/")
			lvmp.SetZoneid(idls[0])
			lvmp.SetName(idls[1])
		}
	*/
	/*
	   if d.HasChange("address") {
	     // Try updating the address
	     if err := updateAddress(d, meta); err != nil {
	         return err
	       }

	     d.SetPartial("address")
	   }
	*/

	// If we were to return here, before disabling partial mode below,
	// then only the "address" field would be saved.

	// We succeeded, disable partial mode. This causes Terraform to save
	// save all fields again.
	//d.Partial(false)

	return resourceKcpsValueVMRead(d, meta)
}

func resourceKcpsValueVMDelete(d *schema.ResourceData, meta interface{}) error {
	cli := meta.(*gk.KCPSClient)

	zoneid := d.Get("zoneid").(string)
	name := d.Get("name").(string)
	mutexKV.Lock("value-vm-" + zoneid + name)
	defer mutexKV.Unlock("value-vm-" + zoneid + name)

	p := cli.VirtualMachine.NewDestroyVirtualMachineParams(d.Id())
	_, err := cli.VirtualMachine.DestroyVirtualMachine(p)

	if err != nil {
		return fmt.Errorf("Error destroying Value Virtual Machine: %s", err)
	}

	// check completely deleted (I don't want to do this...)
	for {

		p := cli.VirtualMachine.NewListVirtualMachinesParams()
		p.SetId(d.Id())
		r, err := cli.VirtualMachine.ListVirtualMachines(p)
		if err != nil {
			return fmt.Errorf("Error getting Value Virtual Machines list: %s", err)
		}
		if r.VirtualMachines == nil {
			break
		}

		time.Sleep(7 * time.Second)
	}

	d.SetId("")
	return nil
}

func flattenIptoNetworks(vm *gk.VirtualMachine) []map[string]interface{} {
	result := make([]map[string]interface{}, 0, len(vm.Nics))
	for _, n := range vm.Nics {
		ipToMap := make(map[string]interface{})
		ipToMap["networkid"] = n.Networkid
		ipToMap["ip"] = n.Ipaddress
		ipToMap["ipv6"] = n.Ip6address

		result = append(result, ipToMap)
	}
	return result
}
