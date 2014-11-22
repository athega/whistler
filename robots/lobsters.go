package robots

import (
	"fmt"
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
		b.respond(c, "could not find the hottest stories", []Attachment{})
		return
	}

	attachments := []Attachment{}
	for _, story := range stories {
		url := story.URL

		if url == "" {
			url = story.CommentsURL
		}

		attachments = append(attachments, Attachment{
			Text:  fmt.Sprintf("Score: %v", story.Score),
			Color: "#FF6600",
			Fields: []AttachmentField{
				AttachmentField{
					Title: story.Title,
					Value: url,
				},
			},
		})
	}

	b.respond(c, "", attachments)
}

func (b LobstersBot) respond(c *SlashCommand, text string, attachments []Attachment) {
	MakeIncomingWebhookCall(&IncomingWebhook{
		Channel:     "@" + c.Username,
		Username:    "Lobste.rs",
		IconEmoji:   ":fried_shrimp:",
		Text:        text,
		Attachments: attachments,
	})
}
