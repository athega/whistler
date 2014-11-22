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

	Config.Domain = env.String("SLACK_DOMAIN", "athega")
	if Config.Domain == "" {
		log.Fatal("SLACK_DOMAIN not set")
	}

	Config.Port = env.String("PORT", "5454")
	if Config.Port == "" {
		log.Fatal("PORT not set")
	}

	Config.Token = env.String("SLACK_WEBHOOK_URL", "")
	if Config.Token == "" {
		log.Fatal("SLACK_WEBHOOK_URL not set")
	}
}

// RegisterRobot registers a command and init function for a robot
func RegisterRobot(command string, RobotInitFunction func() Robot) {
	if _, ok := Robots[command]; ok {
		log.Printf("There are two robots mapped to %s!", command)
	} else {
		log.Printf("Registered: %s", command)
		Robots[command] = RobotInitFunction
	}
}

// MakeIncomingWebhookCall makes an incomming web hook call to Slack
func MakeIncomingWebhookCall(payload *IncomingWebhook) error {
	webhook, err := url.Parse(Config.WebhookURL)
	if err != nil {
		return err
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	postData := url.Values{}
	postData.Set("payload", string(jsonPayload))
	postData.Set("token", Config.Token)

	webhook.RawQuery = postData.Encode()
	resp, err := http.PostForm(webhook.String(), postData)

	if resp.StatusCode != 200 {
		message := fmt.Sprintf("ERROR: Non-200 Response from Slack Incoming Webhook API: %s", resp.Status)
		log.Println(message)
	}

	return err
}
