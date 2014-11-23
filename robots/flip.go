package robots

import (
	"strings"

	"github.com/peterhellberg/flip"
)

// FlipBot is a bot who flips text
type FlipBot struct {
}

func init() {
	RegisterRobot("flip", func() (robot Robot) { return new(FlipBot) })
}

// Run executes a deferred action
func (b FlipBot) Run(c *SlashCommand) string {
	go b.DeferredAction(c)
	return ""
}

// DeferredAction makes a incoming webhook call
func (b FlipBot) DeferredAction(c *SlashCommand) {
	args := strings.Split(c.Text, " ")

	switch args[0] {
	default:
		b.respond(c, flip.UpsideDown(strings.Join(args, " ")))
	case "table", "t":
		b.respond(c, flip.Table(strings.Join(args[1:], " ")))
	}
}

// Description describes what the robot does
func (b FlipBot) Description() string {
	return strings.Join([]string{
		"Flip!",
		"Usage: /whistler flip [table] [args]",
		"Expected Response: (╯°□°）╯︵ʇxǝʇ",
	}, "\n\t")
}

func (b FlipBot) respond(c *SlashCommand, text string) {
	MakeIncomingWebhookCall(&IncomingWebhook{
		Channel:   c.ChannelID,
		Username:  "Flip",
		IconEmoji: ":troll:",
		Text:      text,
	})
}
