package main

import (
	"github.com/hashicorp/terraform/builtin/providers/appconfig"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: appconfig.Provider,
	})
}