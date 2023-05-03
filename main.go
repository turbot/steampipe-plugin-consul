package main

import (
	"github.com/steampipe-plugin-consul/consul"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: consul.Plugin})
}
