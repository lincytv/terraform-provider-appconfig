package appconfig

import (
	"github.com/lincytv/terraform-provider-appconfig/appconfig"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: appconfig.Provider,
	})
}