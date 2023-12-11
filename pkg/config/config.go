package config

import "github.com/hashicorp/go-plugin"

var PluginName = "SERVICE_PROVIDER_PLUGIN"
var PluginPath = "./out/providers"

var handshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "SERVICE_PROVIDER_PLUGIN",
	MagicCookieValue: "This is a handshake config",
}

func HankShakeConfig() plugin.HandshakeConfig {
	return handshakeConfig
}
