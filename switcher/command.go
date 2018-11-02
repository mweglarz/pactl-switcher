package switcher

import (
	"os/exec"
	"strconv"
	"strings"
)

type PactlCommand struct{}

func NewPactlCommand() *PactlCommand {
	return &PactlCommand{}
}

func (self *PactlCommand) RawInputs() (string, error) {
	args := []string{"list", "sink-inputs"}
	cmd := self.command(args)
	out, err := cmd.Output()
	return string(out), err
}

func (self *PactlCommand) ListSinks() ([]SinkInput, error) {
	rawInput, err := self.RawInputs()
	if err != nil {
		return nil, err
	}

	reader := strings.NewReader(rawInput)
	parser := NewParser()

	return parser.Parse(reader)
}

func (self *PactlCommand) ListInputs() ([]int, error) {

	return nil, nil
}

func (self *PactlCommand) MoveInput(inputId, sinkId int) error {
	args := []string{"move-sink-input", strconv.Itoa(inputId), strconv.Itoa(sinkId)}

	cmd := self.command(args)
	err := cmd.Run()
	return err
}

func (self *PactlCommand) command(args []string) *exec.Cmd {
	cmd := exec.Command("pactl", args...)
	return cmd
}

func (self *PactlCommand) getAndParseRawSinks() ([]int, error) {
	return nil, nil
}
