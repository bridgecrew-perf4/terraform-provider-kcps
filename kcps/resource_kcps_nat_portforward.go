package kcps

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	gk "github.com/uesyn/gokcps"
)

func resourceKcpsNatPortForward() *schema.Resource {
	return &schema.Resource{
		Create: resourceKcpsNatPortForwardCreate,
		Read:   resourceKcpsNatPortForwardRead,
		Delete: resourceKcpsNatPortForwardDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"ipaddressid": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"protocol": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateProtocol([]string{"tcp", "udp"}),
			},

			"port": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Required: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"privateport": {
							Type:     schema.TypeInt,
							Required: true,
							ForceNew: true,
						},
						"publicport": {
							Type:     schema.TypeInt,
							Required: true,
							ForceNew: true,
						},
						"privateendport": {
							Type:     schema.TypeInt,
							Optional: true,
							ForceNew: true,
						},
						"publicendport": {
							Type:     schema.TypeInt,
							Optional: true,
							ForceNew: true,
						},
					},
				},
			},

			"virtualmachineid": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"vmguestip": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"ipaddress": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceKcpsNatPortForwardCreate(d *schema.ResourceData, meta interface{}) error {
	cli := meta.(*gk.KCPSClient)

	prefix := "port.0."

	ipaddressid := d.Get("ipaddressid").(string)
	privateport := d.Get(prefix + "privateport").(int)
	protocol := d.Get("protocol").(string)
	publicport := d.Get(prefix + "publicport").(int)
	virtualmachineid := d.Get("virtualmachineid").(string)

	mutexKV.Lock("nat-portforward-" + ipaddressid + virtualmachineid)
	defer mutexKV.Unlock("nat-portforward-" + ipaddressid + virtualmachineid)

	p := cli.NatPortForward.NewCreatePortForwardingRuleParams(ipaddressid, privateport, protocol, publicport, virtualmachineid)

	if privateendport, ok := d.GetOk(prefix + "privateendport"); ok {
		p.SetPrivateendport(privateendport.(int))
	}
	if publicendport, ok := d.GetOk(prefix + "publicendport"); ok {
		p.SetPublicendport(publicendport.(int))
	}
	if vmguestip, ok := d.GetOk("vmguestip"); ok {
		p.SetVmguestip(vmguestip.(string))
	}

	p.SetOpenfirewall(false) //disable automatic creating Firewall Rule

	r, err := cli.NatPortForward.CreatePortForwardingRule(p)
	if err != nil {
		return fmt.Errorf("Error creationg Port Forwarding rule: %s", err)
	}
	d.SetId(r.Id)

	d.Set("ipaddress", r.Ipaddress)
	return resourceKcpsNatPortForwardRead(d, meta)
}

func resourceKcpsNatPortForwardRead(d *schema.ResourceData, meta interface{}) error {
	cli := meta.(*gk.KCPSClient)

	if d.Id() == "" {
		d.SetId("")
		return nil
	}

	p := cli.NatPortForward.NewListPortForwardingRulesParams()
	p.SetId(d.Id())
	r, err := cli.NatPortForward.ListPortForwardingRules(p)

	if err != nil {
		return fmt.Errorf("Error getting Port Forwarding list: %s", err)
	}

	if r.PortForwardingRules == nil {
		d.SetId("")
		return nil
	}

	return nil
}

func resourceKcpsNatPortForwardDelete(d *schema.ResourceData, meta interface{}) error {
	cli := meta.(*gk.KCPSClient)

	ipaddressid := d.Get("ipaddressid").(string)
	virtualmachineid := d.Get("virtualmachineid").(string)

	mutexKV.Lock("nat-portforward-" + ipaddressid + virtualmachineid)
	defer mutexKV.Unlock("nat-portforward-" + ipaddressid + virtualmachineid)

	p := cli.NatPortForward.NewDeletePortForwardingRuleParams(d.Id())
	_, err := cli.NatPortForward.DeletePortForwardingRule(p)

	if err != nil {
		return fmt.Errorf("Error deleting Port Forwarding rule: %s", err)
	}

	d.SetId("")
	return nil
}
