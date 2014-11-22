package robots

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type DecideBot struct {
}

func init() {
	RegisterRobot("decide", func() (robot Robot) { return new(DecideBot) })
}

func (d DecideBot) Run(command *SlashCommand) (slashCommandImmediateReturn string) {
	go d.DeferredAction(command)

	if strings.TrimSpace(command.Text) == "" {
		return "I need something to decide on!"
	} else {
		return ""
	}
}

func (d DecideBot) DeferredAction(command *SlashCommand) {
	response := new(IncomingWebhook)

	response.Channel = "#" + command.ChannelName
	response.Username = "Whistler"
	response.IconEmoji = ":whistler:"
	response.UnfurlLinks = true
	response.Parse = "full"

	text := strings.TrimSpace(command.Text)
	if text != "" {
		split := strings.Split(text, ",")
		response.Text = fmt.Sprintf("@%s: Deciding between: (%s)", command.Username, strings.Join(split, ", "))
		MakeIncomingWebhookCall(response)
		response.Text = fmt.Sprintf("@%s: Decided on: %s", command.Username, Decide(split))
		MakeIncomingWebhookCall(response)
	}
}

func (d DecideBot) Description() (description string) {
	return strings.Join([]string{
		"Decides your fate!",
		"Usage: /decide Life Death ...",
		"Expected Response: Deciding on (Life, Death, ...)",
		"Decided on Life!",
	}, "\n\t")
}

func Decide(Fates []string) (result string) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := len(Fates)

	if n > 0 {
		return strings.TrimSpace(Fates[r.Intn(n)])
	} else {
		return fmt.Sprintf("Error")
	}
}
