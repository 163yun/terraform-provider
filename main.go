package main

import (
	"github.com/163yun/terraform-provider/netease"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: netease.Provider,
	})
}
