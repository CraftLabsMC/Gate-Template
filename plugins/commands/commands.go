package commands

import (
	"context"

	"github.com/go-logr/logr"
	"go.minekube.com/gate/pkg/edition/java/proxy"
)

// Plugin is a ping plugin that handles ping events.
var Plugin = proxy.Plugin{
	Name: "Commands",
	Init: func(ctx context.Context, p *proxy.Proxy) error {
		log := logr.FromContextOrDiscard(ctx)
		log.Info("Command System Loaded!")

		p.Command().Register(titleCommand())

		return nil
	},
}
