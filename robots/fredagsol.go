package robots

import "strings"

// FredagsölBot keeps track of when to drink beer
type FredagsölBot struct {
}

func init() {
	RegisterRobot("fredagsöl", func() (robot Robot) { return new(FredagsölBot) })
}

// Run executes a deferred action
func (b FredagsölBot) Run(c *SlashCommand) string {
	go b.DeferredAction(c)
	return ""
}

// DeferredAction makes a incoming webhook call
func (b FredagsölBot) DeferredAction(c *SlashCommand) {
	MakeIncomingWebhookCall(&IncomingWebhook{
		Channel:   c.ChannelID,
		IconEmoji: ":beer:",
		Username:  "Fredagsöl",
		Text:      "Vet inte.",
	})
}

// Description describes the Fredagsöl bot
func (b FredagsölBot) Description() string {
	return strings.Join([]string{
		"Fredagsöl bot!",
		"Usage: /fredagsöl",
		"Expected Response: Fredagsöl nu!",
	}, "\n\t")
}
