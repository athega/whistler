package robots

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/athega/whistler/env"
)

// Robots contains a map of robots
var Robots = make(map[string]func() Robot)

// Config contains the main configuration for the robots
var Config = new(Configuration)

func init() {
	flag.Parse()

	Config.Port = env.String("PORT", "5454")
	Config.WebhookURL = env.String("SLACK_WEBHOOK_URL", "")
}

// RegisterRobot registers a command and init function for a robot
func RegisterRobot(command string, robotInitFunction func() Robot) {
	if _, ok := Robots[command]; ok {
		log.Printf("There are two robots mapped to %s!", command)
	} else {
		Robots[command] = robotInitFunction
	}
}

// MakeIncomingWebhookCall makes an incomming web hook call to Slack
func MakeIncomingWebhookCall(payload *IncomingWebhook) error {
	_, err := url.Parse(Config.WebhookURL)
	if err != nil {
		return err
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	postData := url.Values{}
	postData.Set("payload", string(jsonPayload))

	fmt.Println("Sending payload:", payload)

	resp, err := http.PostForm(Config.WebhookURL, postData)
	if resp.StatusCode != 200 {
		message := fmt.Sprintf("ERROR: Non-200 Response from Slack Incoming Webhook API: %s", resp.Status)
		log.Println(message, resp)
	}

	return nil
}
