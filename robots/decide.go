package robots

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// DecideBot decides, deal with it
type DecideBot struct {
}

func init() {
	RegisterRobot("decide", func() (robot Robot) { return new(DecideBot) })
}

// Run executes a deferred action
func (b DecideBot) Run(command *SlashCommand) (slashCommandImmediateReturn string) {
	go b.DeferredAction(command)

	if strings.TrimSpace(command.Text) == "" {
		return "I need something to decide on!"
	}

	return ""
}

// DeferredAction makes two incoming webhook calls
func (b DecideBot) DeferredAction(command *SlashCommand) {
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

// Description describes what the robot does
func (b DecideBot) Description() (description string) {
	return strings.Join([]string{
		"Decides your fate!",
		"Usage: /whistler decide Life, Death, ...",
		"Expected Response: Deciding on (Life, Death, ...)",
		"Decided on Life",
	}, "\n\t")
}

// Decide is the implementation of the decision logic
func Decide(Fates []string) (result string) {
	r, n := rand.New(rand.NewSource(time.Now().UnixNano())), len(Fates)

	if n > 0 {
		return strings.TrimSpace(Fates[r.Intn(n)])
	}

	return fmt.Sprintf("Error")
}
