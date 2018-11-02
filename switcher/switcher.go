package switcher

import (
	"os/exec"
	"strconv"
)

type Switcher struct {
}

func NewSwitcher() *Switcher {
	return &Switcher{}
}

func (self *Switcher) SwitchAllToSink(sinkId int) error {
	inputs, err := self.getAllInputs()
	if err != nil {
		return err
	}

	var batchErr error
	for _, inputId := range inputs {
		err := self.SwitchInputToSink(inputId, sinkId)
		if err != nil {
			batchErr = err
		}
	}
	return batchErr
}

func (self *Switcher) SwitchInputToSink(inputId int, sinkId int) error {

	args := []string{"move-sink-input", strconv.Itoa(inputId), strconv.Itoa(sinkId)}

	cmd := exec.Command("pactl", args...)
	err := cmd.Run()
	return err
}

func (self *Switcher) getAllInputs() ([]int, error) {

	return nil, nil
}
