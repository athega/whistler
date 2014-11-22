package robots

// Bot is a robot that list all robots
type Bot struct {
}

func init() {
	RegisterRobot("c", func() (robot Robot) { return new(Bot) })
}

// Run returns a list of all robots
func (r Bot) Run(command *SlashCommand) (slashCommandImmediateReturn string) {
	output := ""
	for command, RobotInitFunction := range Robots {
		robot := RobotInitFunction()
		output = output + "\n" + command + " - " + robot.Description() + "\n"
	}
	return output
}

// Description describes what the Robots robot does
func (r Bot) Description() (description string) {
	return "Lists commands!\n\tUsage: You already know how to use this!\n\tExpected Response: This message!"
}
