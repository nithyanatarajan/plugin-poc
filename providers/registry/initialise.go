package registry

import (
	"github.com/hashicorp/go-plugin"
	"github.com/nithyanatarajan/plugin-poc/pkg/shared"
)

// serviceProviders holds all registered ServiceProviders
var serviceProviders = []*shared.ServiceProviderPlugin{}

// RegisterServiceProvider registers a ServiceProvider
func RegisterServiceProvider(p *shared.ServiceProviderPlugin) {
	serviceProviders = append(serviceProviders, p)
}

// GetAvailableProviders returns all providers present
func GetAvailableProviders() (result []string) {
	for _, provider := range serviceProviders {
		name := provider.Impl.GetName()
		result = append(result, name)
	}
	return result
}

// GetPluginSet returns all providers present as PluginSet
func GetPluginSet() plugin.PluginSet {
	pluginMap := make(plugin.PluginSet)
	for _, provider := range serviceProviders {
		pluginName := provider.Impl.GetName()
		pluginMap[pluginName] = provider
	}
	return pluginMap
}
