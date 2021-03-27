package kcps

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	gk "github.com/uesyn/gokcps"
)

func resourceKcpsFirewall() *schema.Resource {
	return &schema.Resource{
		Create: resourceKcpsFirewallCreate,
		Read:   resourceKcpsFirewallRead,
		Delete: resourceKcpsFirewallDelete,
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
				ValidateFunc: validateProtocol([]string{"tcp", "udp", "icmp"}),
			},

			"cidrlist": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"port": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"startport": {
							Type:     schema.TypeInt,
							Required: true,
							ForceNew: true,
						},
						"endport": {
							Type:     schema.TypeInt,
							Required: true,
							ForceNew: true,
						},
					},
				},
			},

			"icmp": {
				Type:          schema.TypeList,
				MaxItems:      1,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"port"},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"icmpcode": {
							Type:     schema.TypeInt,
							Required: true,
							ForceNew: true,
						},
						"icmptype": {
							Type:     schema.TypeInt,
							Required: true,
							ForceNew: true,
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

func resourceKcpsFirewallCreate(d *schema.ResourceData, meta interface{}) error {
	cli := meta.(*gk.KCPSClient)

	ipaddressid := d.Get("ipaddressid").(string)

	mutexKV.Lock("firewall-" + ipaddressid)
	defer mutexKV.Unlock("firewall-" + ipaddressid)

	protocol := d.Get("protocol").(string)
	var cidrlist []string
	for _, rawCidr := range d.Get("cidrlist").([]interface{}) {
		cidrlist = append(cidrlist, rawCidr.(string))
	}

	prefixPort := "port.0."
	p := cli.Firewall.NewCreateFirewallRuleParams(ipaddressid, protocol, cidrlist)
	if startport, ok := d.GetOk(prefixPort + "startport"); ok {
		p.SetStartport(startport.(int))
	}
	if endport, ok := d.GetOk(prefixPort + "endport"); ok {
		p.SetEndport(endport.(int))
	}

	prefixIcmp := "icmp.0."
	if icmptype, ok := d.GetOkExists(prefixIcmp + "icmptype"); ok {
		p.SetIcmptype(icmptype.(int))
	}
	if icmpcode, ok := d.GetOkExists(prefixIcmp + "icmpcode"); ok {
		p.SetIcmpcode(icmpcode.(int))
	}

	r, err := cli.Firewall.CreateFirewallRule(p)
	if err != nil {
		return fmt.Errorf("Error creating Firewall Rule: %s", err)
	}
	d.SetId(r.Id)

	// Set Connection Info (for remote-exec)
	sshIp := r.Ipaddress
	d.SetConnInfo(map[string]string{
		"type": "ssh",
		"host": sshIp,
	})

	d.Set("ipaddress", r.Ipaddress)
	return resourceKcpsFirewallRead(d, meta)
}

func resourceKcpsFirewallRead(d *schema.ResourceData, meta interface{}) error {
	cli := meta.(*gk.KCPSClient)

	if d.Id() == "" {
		d.SetId("")
		return nil
	}

	p := cli.Firewall.NewListFirewallRulesParams()
	p.SetId(d.Id())
	r, err := cli.Firewall.ListFirewallRules(p)

	if err != nil {
		return fmt.Errorf("Error getting Firewall Rule list: %s", err)
	}

	if r.FirewallRules == nil {
		d.SetId("")
		return nil
	}

	return nil
}

func resourceKcpsFirewallDelete(d *schema.ResourceData, meta interface{}) error {
	cli := meta.(*gk.KCPSClient)

	ipaddressid := d.Get("ipaddressid").(string)
	mutexKV.Lock("firewall-" + ipaddressid)
	defer mutexKV.Unlock("firewall-" + ipaddressid)

	p := cli.Firewall.NewDeleteFirewallRuleParams(d.Id())
	_, err := cli.Firewall.DeleteFirewallRule(p)

	if err != nil {
		return fmt.Errorf("Error deleting Firewall Rule: %s", err)
	}

	d.SetId("")
	return nil
}
