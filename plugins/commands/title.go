package commands

import (
	"github.com/InvalidJokerDE/TestProxy/util"
	"go.minekube.com/brigodier"
	c "go.minekube.com/common/minecraft/component"
	"go.minekube.com/gate/pkg/command"
	"go.minekube.com/gate/pkg/edition/java/proxy"
	"go.minekube.com/gate/pkg/edition/java/title"
)

func titleCommand() brigodier.LiteralNodeBuilder {
	showTitle := command.Command(func(ctx *command.Context) error {
		player, ok := ctx.Source.(proxy.Player)
		if !ok {
			return ctx.Source.SendMessage(&c.Text{Content: "You must be a player to run this command."})
		}

		return title.ShowTitle(player, &title.Options{
			Title:    util.Text(ctx.String("title")),
			Subtitle: util.Text(ctx.String("subtitle")), // empty if arg not provided
		})
	})

	return brigodier.Literal("title").
		Then(brigodier.Argument("title", brigodier.String).Executes(showTitle).
			Then(brigodier.Argument("subtitle", brigodier.StringPhrase).Executes(showTitle)))
}
