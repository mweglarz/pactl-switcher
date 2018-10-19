package switcher

import (
	"fmt"
	"os"

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

	if c.GlobalIsSet("data") {
		fmt.Println("data set, reading from stdin")
		parser := NewParser()
		inputs, err := parser.Parse(os.Stdin)
		fmt.Println("inputs", inputs)
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
			switcher.SwitchAllToSink(sinkId)
			return nil
		}
	}

	return nil
}

func generateFlags() []cli.Flag {

	return []cli.Flag{
		cli.BoolFlag{
			Name:  "data, d",
			Usage: "if true reads pactl data from stdin",
		},
		cli.IntFlag{
			Name:  "sink, s",
			Usage: "pactl sink to move to",
			Value: -1,
		},
		cli.IntFlag{
			Name:  "input, i",
			Usage: "pactl input id",
			Value: -1,
		},
	}
}
