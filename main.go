package main

import (
	"github.com/fvp-mds/terraform-provider-mysql/mysql"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: mysql.Provider})
}
