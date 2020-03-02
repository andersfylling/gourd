package modules

import (
	"context"
	"github.com/Salmonllama/Gourd"
	"github.com/andersfylling/disgord"
	"strings"
)

func ModuleModeration() (module *gourd.Module) {
	module = gourd.NewModule("Moderation", 0).AddCommands(kick(), ban())

	return
}

func kick() (command *gourd.Command) {
	command = gourd.NewCommand("kick", "boot", "bai")   // Aliases are mandatory
	command.SetDescription("Kicks the mentioned users") // Description is optional
	command.SetHelp("[]users, reason for the kick")     // Help is optional

	command.SetOnAction(func(ctx gourd.CommandContext) {
		if ctx.IsPrivate() {
			ctx.Reply("This command can only be used within a server!")
			return
		}

		users := ctx.Message().Mentions                      // Get any mentioned users
		reason := strings.Join(ctx.Args()[len(users):], " ") // Reason starts after the user mentions

		for _, user := range users {
			err := ctx.Client().KickMember(context.Background(), ctx.Guild().ID, user.ID, reason)
			if err != nil {
				// Handle the error
			}
		}
	})

	return
}

func ban() (command *gourd.Command) {
	command = gourd.NewCommand("ban", "banana", "hammertime")
	command.SetDescription("Bans the mentioned users")
	command.SetHelp("[]users, reason for the ban")

	command.SetOnAction(func(ctx gourd.CommandContext) {
		if ctx.IsPrivate() {
			ctx.Reply("This command can only be used within a server!")
			return
		}

		users := ctx.Message().Mentions                      // Get any mentioned users
		reason := strings.Join(ctx.Args()[len(users):], " ") // Reason starts after the user mentions

		for _, user := range users {
			banParams := &disgord.BanMemberParams{
				DeleteMessageDays: 30,
				Reason:            reason,
			}

			err := ctx.Client().BanMember(context.Background(), ctx.Guild().ID, user.ID, banParams)
			if err != nil {
				// Handle the error
			}
		}
	})

	return
}