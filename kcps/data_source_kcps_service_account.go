package kcps

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	gk "github.com/uesyn/gokcps"
)

func dataSourceKcpsServiceAccount() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceKcpsServiceAccountRead,

		Schema: map[string]*schema.Schema{
			"account": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			}, "accountid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			}, "accounttype": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			}, "apikey": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			}, "created": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			}, "domain": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			}, "domainid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			}, "email": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			}, "dirstname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			}, "id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			}, "iscallerchilddomain": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			}, "isdefault": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			}, "lastname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			}, "secretkey": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			}, "state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			}, "timezone": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			}, "username": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func dataSourceKcpsServiceAccountRead(d *schema.ResourceData, meta interface{}) error {
	cli := meta.(*gk.KCPSClient)

	p := cli.AccountDomain.NewListUsersParams()
	r, err := cli.AccountDomain.ListUsers(p)

	if err != nil {
		return fmt.Errorf("Error getting User list: %s", err)
	}

	if r.Users == nil {
		return fmt.Errorf("User List not found")
	}
	v := r.Users[0]

	d.Set("account", v.Account)
	d.Set("accountid", v.Accountid)
	d.Set("accounttype", v.Accounttype)
	d.Set("apikey", v.Apikey)
	d.Set("created", v.Created)
	d.Set("domain", v.Domain)
	d.Set("domainid", v.Domainid)
	d.Set("email", v.Email)
	d.Set("firstname", v.Firstname)
	d.Set("id", v.Id)
	d.Set("iscallerchilddomain", v.Iscallerchilddomain)
	d.Set("isdefault", v.Isdefault)
	d.Set("lastname", v.Lastname)
	d.Set("secretkey", v.Secretkey)
	d.Set("state", v.State)
	d.Set("timezone", v.Timezone)
	d.Set("username", v.Username)

	d.SetId(v.Id)

	return nil
}
