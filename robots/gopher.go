package robots

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// GopherBot is a robot who only cares about gophers
type GopherBot struct {
	Gophers map[string]string
}

func init() {
	RegisterRobot("gopher", func() (robot Robot) {
		return &GopherBot{Gophers: gophers}
	})
}

// Run executes a deferred action
func (b GopherBot) Run(c *SlashCommand) string {
	go b.DeferredAction(c)
	return ""
}

// DeferredAction makes a incoming webhook call
func (b GopherBot) DeferredAction(c *SlashCommand) {
	args := strings.Split(strings.ToLower(c.Text), " ")

	switch args[0] {
	default:
		b.gopher(c, args)
	case "random", "r":
		b.random(c, args[1:])
	case "sticker", "s":
		b.sticker(c, args[1:])
	}
}

// Description describes what the robot does
func (b GopherBot) Description() string {
	return strings.Join([]string{
		"Gophers!",
		"Usage: /whistler gopher [args]",
		"Commands: sticker, random",
		"Expected Response: An URL to a Gopher!",
	}, "\n\t")
}

func (b GopherBot) sticker(c *SlashCommand, args []string) {
	if len(args) == 0 {
		b.respond(c, b.randomStickerURL())
		return
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		b.respond(c, b.randomStickerURL())
		return
	}

	if url := b.stickerURL(id); url != "" {
		b.respond(c, url)
	}
}

func (b GopherBot) gopher(c *SlashCommand, args []string) {
	if len(args) == 0 {
		b.respond(c, b.randomGopherURL())
		return
	}

	if url := b.Gophers[args[0]]; url != "" {
		b.respond(c, url)
	} else {
		b.respond(c, b.randomGopherURL())
	}
}

func (b GopherBot) random(c *SlashCommand, args []string) {
	b.respond(c, b.randomGopherURL())
}

func (b GopherBot) randomGopherURL() string {
	rand.Seed(time.Now().UTC().UnixNano())

	v := make([]string, 0, len(b.Gophers))

	for _, value := range b.Gophers {
		v = append(v, value)
	}

	return v[rand.Intn(len(v))]
}

func (b GopherBot) randomStickerURL() string {
	return b.stickerURL(b.randomStickerID())
}

func (b GopherBot) randomStickerID() int {
	rand.Seed(time.Now().UTC().UnixNano())

	return rand.Intn(28) + 1
}

func (b GopherBot) stickerURL(id int) string {
	return b.Gophers[fmt.Sprintf("sticker-%02d", id)]
}

func (b GopherBot) respond(c *SlashCommand, text string) {
	MakeIncomingWebhookCall(&IncomingWebhook{
		Channel:     c.ChannelID,
		IconEmoji:   ":gopher:",
		Username:    "Gopher",
		Text:        text,
		UnfurlLinks: false,
		Parse:       "full",
	})
}

var gophers = map[string]string{
	"side":       "https://raw.githubusercontent.com/golang-samples/gopher-vector/master/gopher-side_color.png",
	"front":      "https://raw.githubusercontent.com/golang-samples/gopher-vector/master/gopher-front.png",
	"pink":       "https://qiita-image-store.s3.amazonaws.com/0/14952/db10cc6f-0727-90ce-f3c6-30e6743952cc.png",
	"sticker-01": "https://raw.githubusercontent.com/tenntenn/gopher-stickers/master/png/01.png",
	"sticker-02": "https://raw.githubusercontent.com/tenntenn/gopher-stickers/master/png/02.png",
	"sticker-03": "https://raw.githubusercontent.com/tenntenn/gopher-stickers/master/png/03.png",
	"sticker-04": "https://raw.githubusercontent.com/tenntenn/gopher-stickers/master/png/04.png",
	"sticker-05": "https://raw.githubusercontent.com/tenntenn/gopher-stickers/master/png/05.png",
	"sticker-06": "https://raw.githubusercontent.com/tenntenn/gopher-stickers/master/png/06.png",
	"sticker-07": "https://raw.githubusercontent.com/tenntenn/gopher-stickers/master/png/07.png",
	"sticker-08": "https://raw.githubusercontent.com/tenntenn/gopher-stickers/master/png/08.png",
	"sticker-09": "https://raw.githubusercontent.com/tenntenn/gopher-stickers/master/png/09.png",
	"sticker-10": "https://raw.githubusercontent.com/tenntenn/gopher-stickers/master/png/10.png",
	"sticker-11": "https://raw.githubusercontent.com/tenntenn/gopher-stickers/master/png/11.png",
	"sticker-12": "https://raw.githubusercontent.com/tenntenn/gopher-stickers/master/png/12.png",
	"sticker-13": "https://raw.githubusercontent.com/tenntenn/gopher-stickers/master/png/13.png",
	"sticker-14": "https://raw.githubusercontent.com/tenntenn/gopher-stickers/master/png/14.png",
	"sticker-15": "https://raw.githubusercontent.com/tenntenn/gopher-stickers/master/png/15.png",
	"sticker-16": "https://raw.githubusercontent.com/tenntenn/gopher-stickers/master/png/16.png",
	"sticker-17": "https://raw.githubusercontent.com/tenntenn/gopher-stickers/master/png/17.png",
	"sticker-18": "https://raw.githubusercontent.com/tenntenn/gopher-stickers/master/png/18.png",
	"sticker-19": "https://raw.githubusercontent.com/tenntenn/gopher-stickers/master/png/19.png",
	"sticker-20": "https://raw.githubusercontent.com/tenntenn/gopher-stickers/master/png/20.png",
	"sticker-21": "https://raw.githubusercontent.com/tenntenn/gopher-stickers/master/png/21.png",
	"sticker-22": "https://raw.githubusercontent.com/tenntenn/gopher-stickers/master/png/22.png",
	"sticker-23": "https://raw.githubusercontent.com/tenntenn/gopher-stickers/master/png/23.png",
	"sticker-24": "https://raw.githubusercontent.com/tenntenn/gopher-stickers/master/png/24.png",
	"sticker-25": "https://raw.githubusercontent.com/tenntenn/gopher-stickers/master/png/25.png",
	"sticker-26": "https://raw.githubusercontent.com/tenntenn/gopher-stickers/master/png/26.png",
	"sticker-27": "https://raw.githubusercontent.com/tenntenn/gopher-stickers/master/png/27.png",
	"sticker-28": "https://raw.githubusercontent.com/tenntenn/gopher-stickers/master/png/28.png",
}
