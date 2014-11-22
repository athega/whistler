package robots

// FredagsölBot keeps track of when to drink beer
type FredagsölBot struct {
}

func init() {
	RegisterRobot("fredagsöl", func() (robot Robot) { return new(FredagsölBot) })
}

// Run executes a deferred action
func (b FredagsölBot) Run(command *SlashCommand) (slashCommandImmediateReturn string) {
	go b.DeferredAction(command)
	return ""
}

// DeferredAction makes a incoming webhook call
func (b FredagsölBot) DeferredAction(command *SlashCommand) {
	response := new(IncomingWebhook)
	response.Channel = "#" + command.ChannelName
	response.Username = "Fredagsöl"
	response.Text = "Vet inte."
	response.IconEmoji = ":beer:"

	MakeIncomingWebhookCall(response)
}

// Description describes the Fredagsöl bot
func (b FredagsölBot) Description() (description string) {
	return "Fredagsöl bot!\n\tUsage: /fredagsöl\n\tExpected Response: Fredagsöl nu!"
}
