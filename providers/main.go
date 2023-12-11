package main

import (
	"github.com/hashicorp/go-plugin"
	"github.com/nithyanatarajan/plugin-poc/pkg/config"
	"github.com/nithyanatarajan/plugin-poc/providers/logging"
	"github.com/nithyanatarajan/plugin-poc/providers/registry"
)

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: config.HankShakeConfig(),
		Plugins:         registry.GetPluginSet(),
		Logger:          logging.GetLogger(),
	})
}
