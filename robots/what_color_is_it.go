package robots

import (
	"fmt"
	"strings"
	"time"
)

// WhatColorIsItBot is a “What color is it?” bot
type WhatColorIsItBot struct {
	Location *time.Location
}

func init() {
	RegisterRobot("color", func() (robot Robot) { return NewWhatColorIsItBot() })
}

// NewWhatColorIsItBot creates a new WhatColorIsItBot
func NewWhatColorIsItBot() *WhatColorIsItBot {
	loc, err := time.LoadLocation("Europe/Stockholm")
	if err != nil {
		loc = time.FixedZone("UTC", 7200)
	}

	return &WhatColorIsItBot{Location: loc}
}

// Run executes a deferred action
func (b WhatColorIsItBot) Run(c *SlashCommand) string {
	go b.DeferredAction(c)
	return ""
}

// DeferredAction makes a incoming webhook call
func (b WhatColorIsItBot) DeferredAction(c *SlashCommand) {
	colorHex := b.colorHexNow()
	colorURL := fmt.Sprintf("http://www.colourlovers.com/img/%s/200/200/%s.png", colorHex, colorHex)

	MakeIncomingWebhookCall(&IncomingWebhook{
		Channel:     c.ChannelID,
		Username:    "What color is it?",
		Text:        fmt.Sprintf("The color right now is #%s\n\n%s", colorHex, colorURL),
		IconEmoji:   ":art:",
		UnfurlLinks: true,
		Parse:       "full",
	})
}

// Description describes what the robot does
func (b WhatColorIsItBot) Description() string {
	colorHex := b.colorHexNow()

	return strings.Join([]string{
		"What color is it?",
		"Usage: /whistler color",
		"Expected Response: The color right now is #" + colorHex,
		"http://www.colourlovers.com/img/" + colorHex + "/200/200/" + colorHex + ".png",
	}, "\n\t")
}

func (b WhatColorIsItBot) colorHexNow() string {
	return b.colorHex(time.Now())
}

func (b WhatColorIsItBot) colorHex(t time.Time) string {
	t = t.In(b.Location)

	return fmt.Sprintf("%02d%02d%02d", t.Hour(), t.Minute(), t.Second())
}
