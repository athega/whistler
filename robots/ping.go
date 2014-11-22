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
func (b PingBot) Run(c *SlashCommand) string {
	go b.DeferredAction(c)
	return ""
}

// DeferredAction makes a incoming webhook call
func (b PingBot) DeferredAction(c *SlashCommand) {
	MakeIncomingWebhookCall(&IncomingWebhook{
		Channel:     c.ChannelID,
		Username:    "Ping Bot",
		Text:        fmt.Sprintf("@%s Pong!", c.Username),
		IconEmoji:   ":ghost:",
		UnfurlLinks: true,
		Parse:       "full",
	})
}

// Description describes what the robot does
func (b PingBot) Description() string {
	return strings.Join([]string{
		"Ping bot!",
		"Usage: /whistler ping",
		"Expected Response: @user: Pong!",
	}, "\n\t")
}
