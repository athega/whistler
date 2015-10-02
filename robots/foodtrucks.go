package robots

import (
	"strings"
	"time"

	"github.com/peterhellberg/stockholmfoodtrucks"
)

// FoodTrucksBot is a bot that returns a list of foodtrucks in Stockholm
type FoodTrucksBot struct {
	// Client is a client for stockholmfoodtrucks.nu
	Client *stockholmfoodtrucks.Client
}

func init() {
	RegisterRobot("foodtrucks", func() (robot Robot) {
		return &FoodTrucksBot{Client: stockholmfoodtrucks.NewClient()}
	})
}

// Run executes a deferred action
func (b FoodTrucksBot) Run(c *SlashCommand) string {
	go b.DeferredAction(c)
	return ""
}

// DeferredAction makes a incoming webhook call
func (b FoodTrucksBot) DeferredAction(c *SlashCommand) {
	foodTrucks, err := b.Client.All()
	if err != nil {
		b.respond(c, "No food trucks found", []Attachment{})
		return
	}

	attachments := []Attachment{}

	for _, truck := range foodTrucks {
		sunday := Now(time.Now()).BeginningOfWeek()

		if truck.Time.After(sunday) {
			attachments = append(attachments, Attachment{
				Text:  truck.TimeText,
				Color: truck.Hex,
				Fields: []AttachmentField{
					AttachmentField{
						Title: truck.Name,
						Value: truck.Text,
					},
				},
			})
		}
	}

	b.respond(c, "", attachments)
}

// Description describes what the robot does
func (b FoodTrucksBot) Description() string {
	return strings.Join([]string{
		"Food Trucks Bot!",
		"Usage: /whistler foodtrucks",
		"Expected Response: Info about all food trucks",
	}, "\n\t")
}

func (b FoodTrucksBot) respond(c *SlashCommand, text string, attachments []Attachment) {
	MakeIncomingWebhookCall(&IncomingWebhook{
		Channel:     "@" + c.Username,
		Username:    "Food Trucks Bot",
		IconEmoji:   ":minibus:",
		Text:        text,
		Attachments: attachments,
	})
}
