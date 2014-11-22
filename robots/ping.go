package robots

import (
	"fmt"
)

// PingBot is a simple ping/pong bot
type PingBot struct {
}

func init() {
	RegisterRobot("ping", func() (robot Robot) { return new(PingBot) })
}

// Run executes a deferred action
func (p PingBot) Run(command *SlashCommand) (slashCommandImmediateReturn string) {
	go p.DeferredAction(command)
	return ""
}

// DeferredAction makes a incoming webhook call
func (p PingBot) DeferredAction(command *SlashCommand) {
	response := new(IncomingWebhook)
	response.Channel = command.ChannelID
	response.Username = "Ping Bot"
	response.Text = fmt.Sprintf("@%s Pong!", command.Username)
	response.IconEmoji = ":ghost:"
	response.UnfurlLinks = true
	response.Parse = "full"
	MakeIncomingWebhookCall(response)
}

// Description describes the Ping bot
func (p PingBot) Description() (description string) {
	return "Ping bot!\n\tUsage: /ping\n\tExpected Response: @user: Pong!"
}
