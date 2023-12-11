package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/hashicorp/go-hclog"
	"github.com/nithyanatarajan/plugin-poc/pkg/config"
	"github.com/nithyanatarajan/plugin-poc/pkg/shared"
	"github.com/nithyanatarajan/plugin-poc/providers/registry"

	"github.com/hashicorp/go-plugin"
)

func main() {
	logger := hclog.New(&hclog.LoggerOptions{
		Name:   "MAIN.LOGGER",
		Level:  hclog.Error,
		Output: os.Stdout,
	})

	for _, provider := range registry.GetAvailableProviders() {
		pluginMap := map[string]plugin.Plugin{
			provider: &shared.ServiceProviderPlugin{},
		}

		// We're a host! Start by launching the plugin process.
		client := plugin.NewClient(&plugin.ClientConfig{
			HandshakeConfig: config.HankShakeConfig(),
			Plugins:         pluginMap,
			Logger:          logger,
			Cmd:             exec.Command(config.PluginPath),
		})
		defer client.Kill() //nolint:revive

		// Connect via RPC
		rpcClient, err := client.Client()
		if err != nil {
			log.Fatal(err)
		}

		// Request the plugin
		raw, err := rpcClient.Dispense(provider)
		if err != nil {
			log.Fatal(err)
		}
		psp, ok := raw.(shared.ServiceProvider)
		if !ok {
			log.Fatal(errors.New("unable to create ServiceProvider"))
		}

		fmt.Println()
		fmt.Printf("Calling GetTransactionID of %s", provider)
		fmt.Printf("\n=========%s=========\n", psp.GetTransactionID())
		fmt.Printf("Called GetTransactionID of %s", provider)
		fmt.Println()
	}
}
