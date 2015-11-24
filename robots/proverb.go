package robots

import (
	"math/rand"
	"strings"
	"time"
)

// ProverbBot is a simple proverb bot
type ProverbBot struct {
	Proverbs []string
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())

	RegisterRobot("proverb", func() (robot Robot) {
		bot := ProverbBot{
			Proverbs: []string{
				"Don't communicate by sharing memory, share memory by communicating.",
				"Concurrency is not parallelism.",
				"Channels orchestrate; mutexes serialize.",
				"The bigger the interface, the weaker the abstraction.",
				"Make the zero value useful.",
				"interface{} says nothing.",
				"Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.",
				"A little copying is better than a little dependency.",
				"Syscall must always be guarded with build tags.",
				"Cgo must always be guarded with build tags.",
				"Cgo is not Go.",
				"With the unsafe package there are no guarantees.",
				"Clear is better than clever.",
				"Reflection is never clear.",
				"Errors are values.",
				"Don't just check errors, handle them gracefully.",
				"Design the architecture, name the components, document the details.",
				"Documentation is for users.",
				"Don't panic.",
			},
		}

		return bot
	})
}

// Run executes a deferred action
func (b ProverbBot) Run(c *SlashCommand) string {
	go b.DeferredAction(c)
	return ""
}

// DeferredAction makes a incoming webhook call
func (b ProverbBot) DeferredAction(c *SlashCommand) {
	MakeIncomingWebhookCall(&IncomingWebhook{
		Channel:   c.ChannelID,
		Username:  "Proverb",
		Text:      b.randomProverb(),
		IconEmoji: ":gopher:",
	})
}

func (b ProverbBot) randomProverb() string {
	return b.Proverbs[rand.Intn(len(b.Proverbs))]
}

// Description describes what the robot does
func (b ProverbBot) Description() string {
	return strings.Join([]string{
		"Proverb Bot",
		"Usage: /whistler proverb",
		"Expected Response: Errors are values.",
	}, "\n\t")
}
