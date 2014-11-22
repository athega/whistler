package robots

import (
	"strings"

	"github.com/peterhellberg/giphy"
)

// GiphyBot is a robot who speak in GIFs
type GiphyBot struct {
	// Client is a client for the Giphy API
	Client *giphy.Client
}

func init() {
	RegisterRobot("giphy", func() (robot Robot) {
		return &GiphyBot{Client: giphy.DefaultClient}
	})
}

// Run executes a deferred action
func (b GiphyBot) Run(c *SlashCommand) string {
	go b.DeferredAction(c)
	return ""
}

// DeferredAction makes a incoming webhook call
func (b GiphyBot) DeferredAction(c *SlashCommand) {
	args := strings.Split(strings.ToLower(c.Text), " ")

	switch args[0] {
	default:
		b.search(c, args)
	case "search", "s":
		b.search(c, args[1:])
	case "gif", "id":
		b.gif(c, args[1:])
	case "random", "rand", "r":
		b.random(c, args[1:])
	case "translate", "trans", "t":
		b.translate(c, args[1:])
	case "trending", "trend", "tr":
		b.trending(c, args[1:])
	}
}

// Description describes what the robot does
func (b GiphyBot) Description() string {
	return strings.Join([]string{
		"Glorious GIFs! (Powered By Giphy)",
		"Usage: /whistler giphy [args]",
		"Commands: search, gif, random, translate, trending",
		"Expected Response: An URL to a GIF!",
	}, "\n\t")
}

func (b GiphyBot) search(c *SlashCommand, args []string) {
	res, err := b.Client.Search(args)
	if err != nil {
		b.respondWithUsername(c, err.Error())
		return
	}

	if len(res.Data) > 0 {
		msg := "Search result for: `" + strings.Join(args, " ") + "`"

		b.respondWithUsername(c, msg)
		b.respond(c, res.Data[0].Images.FixedHeight.URL)
	}
}

func (b GiphyBot) gif(c *SlashCommand, args []string) {
	if len(args) == 0 {
		b.respondWithUsername(c, "missing Giphy id")
		return
	}

	res, err := b.Client.GIF(args[0])
	if err != nil {
		b.respond(c, err.Error())
	}

	b.respond(c, res.Data.Images.FixedHeightDownsampled.URL)
}

func (b GiphyBot) random(c *SlashCommand, args []string) {
	res, err := b.Client.Random(args)
	if err != nil {
		b.respondWithUsername(c, err.Error())
		return
	}

	b.respond(c, res.Data.FixedHeightDownsampledURL)
}

func (b GiphyBot) translate(c *SlashCommand, args []string) {
	res, err := b.Client.Translate(args)
	if err != nil {
		b.respondWithUsername(c, err.Error())
		return
	}

	b.respond(c, res.Data.Images.FixedHeight.URL)
}

func (b GiphyBot) trending(c *SlashCommand, args []string) {
	res, err := b.Client.Trending(args)
	if err != nil {
		b.respondWithUsername(c, err.Error())
		return
	}

	b.respondWithUsername(c, "Here is a trending image")
	b.respond(c, res.Data[0].Images.FixedHeight.URL)
}

func (b GiphyBot) respond(c *SlashCommand, text string) {
	MakeIncomingWebhookCall(&IncomingWebhook{
		Channel:     c.ChannelID,
		IconEmoji:   ":giphy:",
		Username:    "Giphy",
		Text:        text,
		UnfurlLinks: false,
		Parse:       "full",
	})
}

func (b GiphyBot) respondWithUsername(c *SlashCommand, text string) {
	b.respond(c, "@"+c.Username+" "+text)
}
