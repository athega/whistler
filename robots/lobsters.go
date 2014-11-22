package robots

import (
	"strings"

	"github.com/peterhellberg/lobsters"
)

// LobstersBot is a bot for the Lobsters API
type LobstersBot struct {
	// Client is a client for the Lobsters API
	Client *lobsters.Client
}

func init() {
	RegisterRobot("lobsters", func() (robot Robot) {
		return &LobstersBot{Client: lobsters.NewClient(nil)}
	})
}

// Run executes a deferred action
func (b LobstersBot) Run(c *SlashCommand) string {
	go b.DeferredAction(c)
	return ""
}

// Description describes what the robot does
func (b LobstersBot) Description() string {
	return strings.Join([]string{
		"Lobste.rs!",
		"Usage: /whistler lobsters",
		"Expected Response: Stories from lobste.rs",
	}, "\n\t")
}

func (b LobstersBot) DeferredAction(c *SlashCommand) {
	args := strings.Split(strings.ToLower(c.Text), " ")

	switch args[0] {
	default:
		b.hottest(c, args)
	case "hottest", "h":
		b.hottest(c, args[1:])
	}
}

func (b LobstersBot) hottest(c *SlashCommand, args []string) {
	stories, err := b.Client.Hottest()
	if err != nil {
		b.respond(c, "could not find the hottest stories")
		return
	}

	for i, story := range stories {
		if i > 1 {
			break
		}

		b.respond(c, strings.Join([]string{
			story.Title,
			story.URL,
		}, "\n"))
	}
}

func (b LobstersBot) respond(c *SlashCommand, text string) {
	MakeIncomingWebhookCall(&IncomingWebhook{
		Channel:     c.ChannelID,
		Username:    "Lobste.rs",
		Text:        text,
		IconEmoji:   ":fried_shrimp:",
		UnfurlLinks: true,
		Parse:       "full",
	})
}
