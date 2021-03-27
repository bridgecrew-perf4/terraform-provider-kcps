package main

import (
	"github.com/ezoiwana/terraform-provider-kcps/kcps"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: kcps.Provider,
	})

}
