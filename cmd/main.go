package main

import (
	"fmt"
	"github.com/hashicorp/go-hclog"
	"github.com/nithyanatarajan/plugin-poc/pkg/config"
	"github.com/nithyanatarajan/plugin-poc/pkg/shared"
	"log"
	"os"
	"os/exec"

	"github.com/hashicorp/go-plugin"
)

func main() {
	logger := hclog.New(&hclog.LoggerOptions{
		Name:   "MAIN.LOGGER",
		Level:  hclog.Debug,
		Output: os.Stdout,
	})

	pluginMap := map[string]plugin.Plugin{
		config.PluginName: &shared.ServiceProviderPlugin{},
	}

	// We're a host! Start by launching the plugin process.
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: config.HankShakeConfig(),
		Plugins:         pluginMap,
		Logger:          logger,
		Cmd:             exec.Command(config.PluginPath),
	})
	defer client.Kill()

	// Connect via RPC
	rpcClient, err := client.Client()
	if err != nil {
		log.Fatal(err)
	}

	// Request the plugin
	raw, err := rpcClient.Dispense(config.PluginName)
	if err != nil {
		log.Fatal(err)
	}
	psp := raw.(shared.ServiceProvider)
	fmt.Println()
	fmt.Println("Calling GetTransactionID")
	fmt.Printf("\n=======================%s=======================\n", psp.GetTransactionID())
	fmt.Println("Called GetTransactionID")
	fmt.Println()
}
