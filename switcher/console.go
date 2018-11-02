package switcher

import (
	"fmt"
	"github.com/urfave/cli"
)

/// entry point for console app
func RunPactlSwitcherApp(args []string) {
	app := cli.NewApp()

	app.Name = "Pactl inputs switcher"
	app.Usage = "Simple pactl sink-inputs switcher"
	app.Flags = generateFlags()

	app.Action = func(c *cli.Context) error {
		return run(c)
	}

	err := app.Run(args)
	if err != nil {
		fmt.Errorf("An error occured during translate %v", err)
	}
}

func run(c *cli.Context) error {
	if c.GlobalIsSet("list") {
		pactl := NewPactlCommand()
		inputs, err := pactl.ListInputs()

		for _, sinkInput := range inputs {
			// TODO: add json annotation
			sinkInput.Print()
		}
		return err
	}

	if c.GlobalIsSet("sink") {
		sinkId := c.GlobalInt("sink")
		switcher := NewSwitcher()

		if c.GlobalIsSet("input") {
			inputId := c.GlobalInt("input")
			switcher.SwitchInputToSink(inputId, sinkId)
			return nil

		} else {
			err := switcher.SwitchAllToSink(sinkId)
			return err
		}
	}

	return fmt.Errorf("Unknown option")
}

func generateFlags() []cli.Flag {

	return []cli.Flag{
		cli.BoolFlag{
			Name:  "list, l",
			Usage: "if true prints input list",
		},
		cli.IntFlag{
			Name:  "sink, s",
			Usage: "pactl sink to move to",
		},
		cli.IntFlag{
			Name:  "input, i",
			Usage: "pactl input id",
		},
	}
}
