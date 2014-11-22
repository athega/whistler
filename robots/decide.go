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
func (b DecideBot) Run(c *SlashCommand) string {
	go b.DeferredAction(c)

	if strings.TrimSpace(c.Text) == "" {
		return "I need something to decide on!"
	}

	return ""
}

// DeferredAction makes two incoming webhook calls
func (b DecideBot) DeferredAction(c *SlashCommand) {
	response := &IncomingWebhook{
		Channel:     c.ChannelID,
		Username:    "Whistler",
		IconEmoji:   ":whistler:",
		UnfurlLinks: true,
		Parse:       "full",
	}

	text := strings.TrimSpace(c.Text)
	if text != "" {
		split := strings.Split(text, ",")

		response.Text = fmt.Sprintf("@%s: Deciding between: (%s)", c.Username, strings.Join(split, ", "))
		MakeIncomingWebhookCall(response)

		response.Text = fmt.Sprintf("@%s: Decided on: %s", c.Username, Decide(split))
		MakeIncomingWebhookCall(response)
	}
}

// Description describes what the robot does
func (b DecideBot) Description() string {
	return strings.Join([]string{
		"Decides your fate!",
		"Usage: /whistler decide Life, Death, ...",
		"Expected Response: Deciding on (Life, Death, ...)",
		"Decided on Life",
	}, "\n\t")
}

// Decide is the implementation of the decision logic
func Decide(Fates []string) string {
	r, n := rand.New(rand.NewSource(time.Now().UnixNano())), len(Fates)

	if n > 0 {
		return strings.TrimSpace(Fates[r.Intn(n)])
	}

	return fmt.Sprintf("Error")
}
