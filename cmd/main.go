package main

import (
	"github.com/hashicorp/sentinel-sdk/rpc"
	env "github.com/terraworx/plugin-sentinel-format/plugin"
)

func main() {
	rpc.Serve(&rpc.ServeOpts{
		PluginFunc: env.New,
	})
}
