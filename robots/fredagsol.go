package robots

import (
	"strings"
	"time"
)

// FredagsölBot keeps track of when to drink beer
type FredagsölBot struct {
	Location *time.Location
}

func init() {
	RegisterRobot("fredagsöl", func() (robot Robot) { return NewFredagsölBot() })
}

// NewFredagsölBot creates a new FredagsölBot
func NewFredagsölBot() *FredagsölBot {
	loc, err := time.LoadLocation("Europe/Stockholm")
	if err != nil {
		loc = time.FixedZone("UTC", 7200)
	}

	return &FredagsölBot{Location: loc}
}

// Run executes a deferred action
func (b FredagsölBot) Run(c *SlashCommand) string {
	go b.DeferredAction(c)
	return ""
}

// DeferredAction makes a incoming webhook call
func (b FredagsölBot) DeferredAction(c *SlashCommand) {
	var emoji string

	text, ok := b.isItFredagsöl(time.Now())

	if ok {
		emoji = ":beers:"
	} else {
		emoji = ":beer:"
	}

	MakeIncomingWebhookCall(&IncomingWebhook{
		Channel:   c.ChannelID,
		IconEmoji: emoji,
		Username:  "Fredagsöl",
		Text:      text,
	})
}

// Description describes the Fredagsöl bot
func (b FredagsölBot) Description() string {
	return strings.Join([]string{
		"Fredagsöl bot!",
		"Usage: /fredagsöl",
		"Expected Response: Next Fredagsöl: 48m32s",
	}, "\n\t")
}

func (b FredagsölBot) isItFredagsöl(t time.Time) (string, bool) {
	t = t.In(b.Location)

	if t.Weekday() == time.Friday && t.Hour() >= 17 {
		return "Fredagsöl is now! :beer: :beers:", true
	}

	ft := Now(t).Friday().Add(17 * time.Hour)
	return "Next Fredagsöl: " + ft.Sub(t).String(), false
}
