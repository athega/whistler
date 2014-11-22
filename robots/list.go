package robots

import "strings"

// ListBot is a robot that list all the robots
type ListBot struct {
}

func init() {
	RegisterRobot("list", func() (robot Robot) { return new(ListBot) })
}

// Run returns a list of all robots
func (b ListBot) Run(c *SlashCommand) string {
	output := ""

	for name, fn := range Robots {
		output = output + "\n" + name + " - " + fn().Description() + "\n"
	}

	return output
}

// Description describes what the robot does
func (b ListBot) Description() string {
	return strings.Join([]string{
		"Lists commands!" +
			"Usage: You already know how to use this!" +
			"Expected Response: This message!",
	}, "\n\t")
}
