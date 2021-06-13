package node

import (
	"github.com/RosettaFlow/Carrier-Go/carrier"
	"github.com/RosettaFlow/Carrier-Go/common/flags"
	"github.com/urfave/cli/v2"
)

const (
	clientIdentifier = "carrier" // Client identifier to advertise over the network
)

type carrierConfig struct {
	Carrier   carrier.Config
	Node      Config
}

func makeConfig(ctx *cli.Context) carrierConfig {
	// Load defaults.
	cfg := carrierConfig{
		Carrier:   carrier.DefaultConfig,
		Node:      defaultNodeConfig(),
	}
	// Load config file.
	// todo: file conf load for config.

	// Apply flags.
	flags.SetNodeConfig(ctx, &cfg.Node)
	flags.SetCarrierConfig(ctx, &cfg.Carrier)
	return cfg
}

func defaultNodeConfig() Config {
	cfg := DefaultConfig
	return cfg
}