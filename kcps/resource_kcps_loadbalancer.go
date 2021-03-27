package kcps

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	gk "github.com/uesyn/gokcps"
)

func resourceKcpsLoadBalancer() *schema.Resource {
	return &schema.Resource{
		Create: resourceKcpsLoadBalancerCreate,
		Read:   resourceKcpsLoadBalancerRead,
		Update: resourceKcpsLoadBalancerUpdate,
		Delete: resourceKcpsLoadBalancerDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"algorithm": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateAlgorithm(),
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
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
			"publicipid": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			//'assignToLoadBalancerRule'
			"assignto": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
		},
	}
}

func resourceKcpsLoadBalancerCreate(d *schema.ResourceData, meta interface{}) error {
	cli := meta.(*gk.KCPSClient)

	algorithm := d.Get("algorithm").(string)
	name := d.Get("name").(string)
	privateport := d.Get("privateport").(int)
	publicport := d.Get("publicport").(int)
	publicipid := d.Get("publicipid").(string)

	p := cli.LoadBalancer.NewCreateLoadBalancerRuleParams(
		algorithm,
		name,
		privateport,
		publicport,
		publicipid,
	)

	r, err := cli.LoadBalancer.CreateLoadBalancerRule(p)
	if err != nil {
		return fmt.Errorf("Error creating new LoadBalancer Rule: %s", err)
	}
	d.SetId(r.Id)

	//'assignToLoadBalancerRule'
	if assignto, ok := d.GetOk("assignto"); ok {
		assigntoSet := assignto.(*schema.Set)
		assigntoList := make([]string, assigntoSet.Len())
		for i, v := range assigntoSet.List() {
			assigntoList[i] = v.(string)
		}
		p := cli.LoadBalancer.NewAssignToLoadBalancerRuleParams(r.Id, assigntoList)
		_, err := cli.LoadBalancer.AssignToLoadBalancerRule(p)
		if err != nil {
			return fmt.Errorf("Error assigning Virtual Machines to the LoadBalancer Rule: %s", err)
		}
	}

	return resourceKcpsLoadBalancerRead(d, meta)
}

func resourceKcpsLoadBalancerRead(d *schema.ResourceData, meta interface{}) error {
	cli := meta.(*gk.KCPSClient)

	if d.Id() == "" {
		d.SetId("")
		return nil
	}
	p := cli.LoadBalancer.NewListLoadBalancerRulesParams()
	p.SetId(d.Id())
	r, err := cli.LoadBalancer.ListLoadBalancerRules(p)
	if err != nil {
		return fmt.Errorf("Error getting LoadBalancer Rule list: %s", err)
	}
	if r.LoadBalancerRules == nil {
		d.SetId("")
		return nil
	}

	d.Set("algorithm", r.LoadBalancerRules[0].Algorithm)
	d.Set("name", r.LoadBalancerRules[0].Name)

	if _, ok := d.GetOk("assignto"); ok {
		p := cli.LoadBalancer.NewListLoadBalancerRuleInstancesParams(d.Id())
		r, err := cli.LoadBalancer.ListLoadBalancerRuleInstances(p)
		if err != nil {
			return fmt.Errorf("Error getting LoadBalancer Rule Instances list: %s", err)
		}

		var assigntoList []string
		for _, v := range r.LoadBalancerRuleInstances {
			assigntoList = append(assigntoList, v.Id)
		}
		d.Set("assignto", convertStringArrToInterface(assigntoList))
	}

	return nil
}

func resourceKcpsLoadBalancerUpdate(d *schema.ResourceData, meta interface{}) error {
	cli := meta.(*gk.KCPSClient)

	if d.Id() == "" {
		d.SetId("")
		return nil
	}

	p := cli.LoadBalancer.NewListLoadBalancerRulesParams()
	p.SetId(d.Id())
	r, err := cli.LoadBalancer.ListLoadBalancerRules(p)

	if err != nil {
		return fmt.Errorf("Error getting LoadBalancer Rule list: %s", err)
	}
	if r.LoadBalancerRules == nil {
		d.SetId("")
		return nil
	}

	d.Partial(true)

	//'updateLoadBalancerRule'
	if d.HasChange("algorithm") || d.HasChange("name") {
		p := cli.LoadBalancer.NewUpdateLoadBalancerRuleParams(d.Id(), d.Get("algorithm").(string))
		p.SetName(d.Get("name").(string))

		_, err := cli.LoadBalancer.UpdateLoadBalancerRule(p)
		if err != nil {
			return fmt.Errorf("Error updating the LoadBalancer Rule: %s", err)
		}

		d.SetPartial("algorithm")
		d.SetPartial("name")
	}

	//'assignToLoadBalancerRule' and 'removeFromLoadBalancerRule'
	if d.HasChange("assignto") {
		o, n := d.GetChange("assignto")

		// lead old and new assign points
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

		// remove
		oldList = remove(oldList, "")
		if oldList != nil {
			p := cli.LoadBalancer.NewRemoveFromLoadBalancerRuleParams(d.Id(), oldList)
			_, err := cli.LoadBalancer.RemoveFromLoadBalancerRule(p)
			if err != nil {
				return fmt.Errorf("Error removing Virtual Machines from the LoadBalancer Rule: %s", err)
			}
		}

		// assign
		newList = remove(newList, "")
		if newList != nil {
			p2 := cli.LoadBalancer.NewAssignToLoadBalancerRuleParams(d.Id(), newList)
			_, err = cli.LoadBalancer.AssignToLoadBalancerRule(p2)
			if err != nil {
				return fmt.Errorf("Error assigning Virtual Machines to the LoadBalancer Rule: %s", err)
			}
		}

		d.SetPartial("assignto")
	}

	d.Partial(false)

	return resourceKcpsLoadBalancerRead(d, meta)
}

func resourceKcpsLoadBalancerDelete(d *schema.ResourceData, meta interface{}) error {
	cli := meta.(*gk.KCPSClient)

	p := cli.LoadBalancer.NewDeleteLoadBalancerRuleParams(d.Id())
	_, err := cli.LoadBalancer.DeleteLoadBalancerRule(p)
	if err != nil {
		return fmt.Errorf("Error deleting LoadBalancer Rule: %s", err)
	}

	d.SetId("")
	return nil
}
