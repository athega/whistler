package robots

import (
	"fmt"
	"strings"
	"time"
)

// WhatColorIsItBot is a “What color is it?” bot
type WhatColorIsItBot struct {
}

func init() {
	RegisterRobot("what color is it?", func() (robot Robot) { return new(WhatColorIsItBot) })
}

// Run executes a deferred action
func (b WhatColorIsItBot) Run(c *SlashCommand) string {
	go b.DeferredAction(c)
	return ""
}

// DeferredAction makes a incoming webhook call
func (b WhatColorIsItBot) DeferredAction(c *SlashCommand) {
	colorHex := b.colorHexNow()

	MakeIncomingWebhookCall(&IncomingWebhook{
		Channel:     c.ChannelID,
		Username:    "What color is it?",
		Text:        fmt.Sprintf("The color right now is %s", colorHex),
		IconEmoji:   colorHex,
		UnfurlLinks: true,
		Parse:       "full",
	})
}

// Description describes what the robot does
func (b WhatColorIsItBot) Description() string {
	return strings.Join([]string{
		"What color is it?",
		"Usage: /whistler what color is it?",
		"Expected Response: The color right now is #FF6600",
	}, "\n\t")
}

func (b WhatColorIsItBot) colorHexNow() string {
	return b.colorHex(time.Now())
}

func (b WhatColorIsItBot) colorHex(t time.Time) string {
	return fmt.Printf("#%02d%02d%02d", t.Hour(), t.Minute(), t.Second())
}
