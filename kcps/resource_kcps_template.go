package kcps

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	gk "github.com/uesyn/gokcps"
)

const (
	defaultIsdynamicallyscalable = false
	defaultPasswordenabled       = false
	defaultIspublic              = false
)

func resourceKcpsTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourceKcpsTemplateCreate,
		Read:   resourceKcpsTemplateRead,
		Update: resourceKcpsTemplateUpdate,
		Delete: resourceKcpsTemplateDelete,
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
			"ostypeid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"isdynamicallyscalable": {
				Type:     schema.TypeBool,
				Default:  defaultIsdynamicallyscalable,
				Optional: true,
			},
			"passwordenabled": {
				Type:     schema.TypeBool,
				Default:  defaultPasswordenabled,
				Optional: true,
			},

			//create form snapshot
			"snapshotid": {
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},

			//create from template
			"volumeid": {
				Type:          schema.TypeString,
				ForceNew:      true,
				Optional:      true,
				ConflictsWith: []string{"snapshotid"},
			},

			/* gokcpsでVMware以外のhypervisorを設定できるようになった場合
			"hypervisor": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateHypervisor(),
			},
			*/

			"ispublic": {
				Type:     schema.TypeBool,
				Default:  defaultIspublic,
				Optional: true,
			},
		},
	}
}

func resourceKcpsTemplateCreate(d *schema.ResourceData, meta interface{}) error {
	cli := meta.(*gk.KCPSClient)

	name := d.Get("name").(string)
	displaytext := d.Get("displaytext").(string)
	ostypeid := d.Get("ostypeid").(string)

	p := cli.Template.NewCreateTemplateParams(displaytext, name, ostypeid)
	if snapshotid, ok := d.GetOk("snapshotid"); ok {
		p.SetSnapshotid(snapshotid.(string))
	}
	if volumeid, ok := d.GetOk("volumeid"); ok {
		p.SetVolumeid(volumeid.(string))
	}
	if isdynamicallyscalable, ok := d.GetOk("isdynamicallyscalable"); ok {
		p.SetIsdynamicallyscalable(isdynamicallyscalable.(bool))
	}
	if passwordenabled, ok := d.GetOk("passwordenabled"); ok {
		p.SetPasswordenabled(passwordenabled.(bool))
	}
	if ispublic, ok := d.GetOk("ispublic"); ok {
		p.SetIspublic(ispublic.(bool))
	}
	r, err := cli.Template.CreateTemplate(p)
	if err != nil {
		return fmt.Errorf("Error creating new Template: %s", err)
	}

	d.SetId(r.Id)

	return resourceKcpsTemplateRead(d, meta)
}

func resourceKcpsTemplateRead(d *schema.ResourceData, meta interface{}) error {
	cli := meta.(*gk.KCPSClient)

	if d.Id() == "" {
		d.SetId("")
		return nil
	}
	p := cli.Template.NewListTemplatesParams("self")
	p.SetId(d.Id())

	r, err := cli.Template.ListTemplates(p)
	if err != nil {
		return fmt.Errorf("Error getting Template list: %s", err)
	}
	if r.Templates == nil {
		d.SetId("")
		return nil
	}

	d.Set("name", r.Templates[0].Name)
	d.Set("displaytext", r.Templates[0].Displaytext)
	d.Set("ostypeid", r.Templates[0].Ostypeid)
	d.Set("isdynamicallyscalable", r.Templates[0].Isdynamicallyscalable)
	d.Set("passwordenabled", r.Templates[0].Passwordenabled)
	d.Set("ispublic", r.Templates[0].Ispublic)

	return nil
}

func resourceKcpsTemplateUpdate(d *schema.ResourceData, meta interface{}) error {
	cli := meta.(*gk.KCPSClient)

	if d.Id() == "" {
		d.SetId("")
		return nil
	}

	p := cli.Template.NewListTemplatesParams("self")
	p.SetId(d.Id())
	r, err := cli.Template.ListTemplates(p)

	if err != nil {
		return fmt.Errorf("Error getting Template list: %s", err)
	}
	if r.Templates == nil {
		d.SetId("")
		return nil
	}

	d.Partial(true)

	//'updateTemplate'
	if d.HasChange("name") || d.HasChange("displaytext") ||
		d.HasChange("ostypeid") || d.HasChange("isdynamicallyscalable") || d.HasChange("passwordenabled") {

		p := cli.Template.NewUpdateTemplateParams(d.Id())

		if name, ok := d.GetOk("name"); ok {
			p.SetName(name.(string))
		}
		if displaytext, ok := d.GetOk("displaytext"); ok {
			p.SetDisplaytext(displaytext.(string))
		}
		if ostypeid, ok := d.GetOk("ostypeid"); ok {
			p.SetOstypeid(ostypeid.(string))
		}
		if isdynamicallyscalable, ok := d.GetOk("isdynamicallyscalable"); ok {
			p.SetIsdynamicallyscalable(isdynamicallyscalable.(bool))
		} else {
			p.SetIsdynamicallyscalable(defaultIsdynamicallyscalable)
		}
		if passwordenabled, ok := d.GetOk("passwordenabled"); ok {
			p.SetPasswordenabled(passwordenabled.(bool))
		} else {
			p.SetPasswordenabled(defaultPasswordenabled)
		}

		_, err := cli.Template.UpdateTemplate(p)
		if err != nil {
			return fmt.Errorf("Error updating the Template: %s", err)
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
		if _, ok := d.GetOk("isdynamicallyscalable"); ok {
			d.SetPartial("isdynamicallyscalable")
		}
		if _, ok := d.GetOk("passwordenabled"); ok {
			d.SetPartial("passwordenabled")
		}
	}

	//'updateTemplatePermissions'
	if d.HasChange("ispublic") {
		p := cli.Template.NewUpdateTemplatePermissionsParams(d.Id())
		if ispublic, ok := d.GetOk("ispublic"); ok {
			p.SetIspublic(ispublic.(bool))
		} else {
			p.SetIspublic(defaultIspublic)
		}

		_, err := cli.Template.UpdateTemplatePermissions(p)
		if err != nil {
			return fmt.Errorf("Error updating permission of the Template: %s", err)
		}

		d.SetPartial("ispublic")
	}

	d.Partial(false)

	return resourceKcpsTemplateRead(d, meta)
}

func resourceKcpsTemplateDelete(d *schema.ResourceData, meta interface{}) error {
	cli := meta.(*gk.KCPSClient)

	p := cli.Template.NewDeleteTemplateParams(d.Id())
	_, err := cli.Template.DeleteTemplate(p)
	if err != nil {
		return fmt.Errorf("Error deleting Template: %s", err)
	}

	d.SetId("")
	return nil
}
