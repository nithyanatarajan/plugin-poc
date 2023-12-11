package main

import (
	"os"
	"reflect"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"github.com/nithyanatarajan/plugin-poc/pkg/config"
	"github.com/nithyanatarajan/plugin-poc/pkg/shared"
)

func main() {
	logger := hclog.New(&hclog.LoggerOptions{
		Name:       "PROVIDERS.LOGGER",
		Level:      hclog.Trace,
		Output:     os.Stderr,
		JSONFormat: true,
	})

	providerOne := ServiceProviderOne{
		logger: logger,
	}
	pluginName := reflect.TypeOf(providerOne).Name()
	logger.Debug("<<<<<<<<<<<<<<<<<<<<Log from Plugin>>>>>>>>>>>>>>>>>>>>",
		"PLUGIN_NAME", pluginName)

	pluginMap := plugin.PluginSet{
		config.PluginName: &shared.ServiceProviderPlugin{
			Impl: &providerOne,
		},
	}
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: config.HankShakeConfig(),
		Plugins:         pluginMap,
		Logger:          logger,
	})
}
