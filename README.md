# Whistler

![Whistler avatar](https://raw.githubusercontent.com/athega/whistler/master/images/whistler.jpg)

[Whistler](http://starwars.wikia.com/wiki/Whistler), also called Xeno, was the astromech droid of Corran Horn.
It was the same R-series model as the legendary R2-D2.

_Whistler is based on the [slackbot](https://github.com/trinchan/slackbot) library. He's pretty cool._

## Robots

[![GoDoc](https://godoc.org/github.com/athega/whistler/robots?status.svg)](https://godoc.org/github.com/athega/whistler/robots)

### Robot skeleton

```go
package robots

// FooBot is a robot that …
type FooBot struct {
}

func init() {
	RegisterRobot("foo", func() (robot Robot) { return new(FooBot) })
}

// Run executes a deferred action
func (r FooBot) Run(command *SlashCommand) (slashCommandImmediateReturn string) {
	go r.DeferredAction(command)
	return ""
}

// Description describes what the robot does
func (r ListBot) Description() (description string) {
	return strings.Join([]string{
		"Does something!",
		"Usage: /whistler foo [args]",
		"Expected Response: Something",
	}, "\n\t")
}
```

## Development

### Running Whistler locally

```bash
SLACK_WEBHOOK_URL=https://hooks.slack.com/services/… go run main.go
```

### Restarting Whistler

You may want to get comfortable with `hk log` and `hk restart` if you're having issues.
