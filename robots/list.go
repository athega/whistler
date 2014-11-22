package robots

// ListBot is a robot that list all the robots
type ListBot struct {
}

func init() {
	RegisterRobot("list", func() (robot Robot) { return new(ListBot) })
}

// Run returns a list of all robots
func (b ListBot) Run(command *SlashCommand) (slashCommandImmediateReturn string) {
	output := ""
	for command, RobotInitFunction := range Robots {
		robot := RobotInitFunction()
		output = output + "\n" + command + " - " + robot.Description() + "\n"
	}
	return output
}

// Description describes what the robot does
func (b ListBot) Description() (description string) {
	return "Lists commands!\n\tUsage: You already know how to use this!\n\tExpected Response: This message!"
}
