package robots

// FredagsölBot keeps track of when to drink beer
type FredagsölBot struct {
}

func init() {
	RegisterRobot("fredagsöl", func() (robot Robot) { return new(FredagsölBot) })
}

// Run executes a deferred action
func (p PingBot) Run(command *SlashCommand) (slashCommandImmediateReturn string) {
	go p.DeferredAction(command)
	return ""
}

// DeferredAction makes a incoming webhook call
func (p PingBot) DeferredAction(command *SlashCommand) {
	response := new(IncomingWebhook)
	response.Channel = command.ChannelID
	response.Username = "Fredagsöl"
	response.Text = "Vet inte."
	response.IconEmoji = ":beer:"
	response.UnfurlLinks = true
	response.Parse = "full"
	MakeIncomingWebhookCall(response)
}
