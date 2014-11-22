package robots

import (
	"fmt"
	"strings"
)

// PingBot is a simple ping/pong bot
type PingBot struct {
}

func init() {
	RegisterRobot("ping", func() (robot Robot) { return new(PingBot) })
}

// Run executes a deferred action
func (b PingBot) Run(command *SlashCommand) (slashCommandImmediateReturn string) {
	go b.DeferredAction(command)
	return ""
}

// DeferredAction makes a incoming webhook call
func (b PingBot) DeferredAction(command *SlashCommand) {
	MakeIncomingWebhookCall(&IncomingWebhook{
		Channel:     command.ChannelID,
		Username:    "Ping Bot",
		Text:        fmt.Sprintf("@%s Pong!", command.Username),
		IconEmoji:   ":ghost:",
		UnfurlLinks: true,
		Parse:       "full",
	})
}

// Description describes the Ping bot
func (b PingBot) Description() (description string) {
	return strings.Join([]string{
		"Ping bot!",
		"Usage: /whistler ping",
		"Expected Response: @user: Pong!",
	}, "\n\t")
}
