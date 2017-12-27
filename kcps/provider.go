package kcps

import (
	"github.com/hashicorp/terraform/helper/mutexkv"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

//global mutexKV
var mutexKV = mutexkv.NewMutexKV()

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_url": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("KCPS_API_URL", nil),
				Description: "Endpoint URL of KCPS API. May be  " +
					"https://portal2-east.cloud-platform.kddi.ne.jp:10443/client/api",
			},
			"api_key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("KCPS_API_KEY", nil),
				Description: "Your API Key. Look at this page written about API " +
					"http://iaas.cloud-platform.kddi.ne.jp/developer/api/cloud-stack-api/use/",
			},
			"secret_key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("KCPS_SECRET_KEY", nil),
				Description: "Your Secret Key. Look at this page written about API " +
					"http://iaas.cloud-platform.kddi.ne.jp/developer/api/cloud-stack-api/use/",
			},
			"verify_ssl": {
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("KCPS_VERIFY_SSL", false),
			},
			//"timeout" (Not support)
		},

		DataSourcesMap: map[string]*schema.Resource{
			"kcps_value_vm": dataSourceKcpsValueVM(),
			//"kcps_premium_vm":      dataSourceKcpsPremiumVM(),
			"kcps_volume":                  dataSourceKcpsVolume(),
			"kcps_guestos":                 dataSourceKcpsGuestOS(),
			"kcps_template":                dataSourceKcpsTemplate(),
			"kcps_snapshot":                dataSourceKcpsSnapshot(),
			"kcps_vmsnapshot":              dataSourceKcpsVMSnapshot(),
			"kcps_snapshot_policy":         dataSourceKcpsSnapshotPolicy(),
			"kcps_nic":                     dataSourceKcpsNic(),
			"kcps_publicip":                dataSourceKcpsPublicIP(),
			"kcps_host":                    dataSourceKcpsHost(),
			"kcps_firewall":                dataSourceKcpsFirewall(),
			"kcps_nat_portforward":         dataSourceKcpsNatPortForward(),
			"kcps_loadbalancer":            dataSourceKcpsLoadBalancer(),
			"kcps_loadbalancer_stickiness": dataSourceKcpsLoadBalancerStickiness(),
			"kcps_iso":                     dataSourceKcpsISO(),
			"kcps_zone":                    dataSourceKcpsZone(),
			"kcps_network":                 dataSourceKcpsNetwork(),
			"kcps_service_offering":        dataSourceKcpsServiceOffering(),
			"kcps_disk_offering":           dataSourceKcpsDiskOffering(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"kcps_value_vm": resourceKcpsValueVM(),
			//"kcps_premium_vm":      resourceKcpsPremiumVM(),
			"kcps_volume":                  resourceKcpsVolume(),
			"kcps_template":                resourceKcpsTemplate(),
			"kcps_snapshot":                resourceKcpsSnapshot(),
			"kcps_vmsnapshot":              resourceKcpsVMSnapshot(),
			"kcps_snapshot_policy":         resourceKcpsSnapshotPolicy(),
			"kcps_nic":                     resourceKcpsNic(),
			"kcps_publicip":                resourceKcpsPublicIP(),
			"kcps_host":                    resourceKcpsHost(),
			"kcps_firewall":                resourceKcpsFirewall(),
			"kcps_nat_portforward":         resourceKcpsNatPortForward(),
			"kcps_loadbalancer":            resourceKcpsLoadBalancer(),
			"kcps_loadbalancer_stickiness": resourceKcpsLoadBalancerStickiness(),
			"kcps_iso":                     resourceKcpsISO(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		APIURL:    d.Get("api_url").(string),
		APIKey:    d.Get("api_key").(string),
		SecretKey: d.Get("secret_key").(string),
		VerifySSL: d.Get("verify_ssl").(bool),
	}
	return config.Client(), nil
}
