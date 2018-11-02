package switcher

import (
	"fmt"
)

type Switcher struct {
	pactlCommand *PactlCommand
}

func NewSwitcher() *Switcher {
	return &Switcher{NewPactlCommand()}
}

func (self *Switcher) SwitchAllToSink(sinkId int) error {
	fmt.Println("SwitchAllToSink")
	inputs, err := self.pactlCommand.ListInputs()
	if err != nil {
		return err
	}
	fmt.Printf("inputs = %+v\n", inputs)

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
	return self.pactlCommand.MoveInput(inputId, sinkId)
}
