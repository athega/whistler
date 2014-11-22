package robots

// SlashCommand represents the fields in a Slack slash command
type SlashCommand struct {
	Token       string  `schema:"token"`
	TeamID      string  `schema:"team_id"`
	ChannelID   string  `schema:"channel_id"`
	ChannelName string  `schema:"channel_name"`
	UserID      string  `schema:"user_id"`
	Username    string  `schema:"user_name"`
	Command     string  `schema:"command"`
	Text        string  `schema:"text,omitempty"`
	TriggerWord string  `schema:"trigger_word,omitempty"`
	TeamDomain  string  `schema:"team_domain,omitempty"`
	ServiceID   string  `schema:"service_id,omitempty"`
	Timestamp   float64 `schema:"timestamp,omitempty"`
}

// IncomingWebhook represents the fields in a incoming webhook from Slack
type IncomingWebhook struct {
	Channel     string       `json:"channel"`
	Username    string       `json:"username"`
	Text        string       `json:"text"`
	IconEmoji   string       `json:"icon_emoji,omitempty"`
	IconURL     string       `json:"icon_url,omitempty"`
	Attachments []Attachment `json:"attachments,omitempty"`
	UnfurlLinks bool         `json:"unfurl_links,omitempty"`
	Parse       string       `json:"parse,omitempty"`
	LinkNames   bool         `json:"link_names,omitempty"`
}

// Attachment represents an attachment
type Attachment struct {
	Fallback string            `json:"fallback"`
	Pretext  string            `json:"pretext,omitempty"`
	Text     string            `json:"text,omitempty"`
	Color    string            `json:"color,omitempty"`
	Fields   []AttachmentField `json:"fields,omitempty"`
}

// AttachmentField represents the values for each item in Attachment.Fields
type AttachmentField struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short,omitempty"`
}

// Configuration is the global configuration of the robots
type Configuration struct {
	Domain     string
	Port       string
	WebhookURL string
}

// Robot is the interface all robots must follow
type Robot interface {
	Run(command *SlashCommand) (slashCommandImmediateReturn string)
	Description() (description string)
}
