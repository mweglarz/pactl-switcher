package switcher

import "fmt"

type Switcher struct {
	pactlCommand *PactlCommand
}

func NewSwitcher() *Switcher {
	return &Switcher{NewPactlCommand()}
}

func (self *Switcher) SwitchAllToSink(sinkId int) error {
	inputs, err := self.pactlCommand.ListInputs()
	if err != nil {
		return err
	}

	var batchErr error
	for _, inputId := range inputs {
		err := self.SwitchInputToSink(inputId.Id, sinkId)
		if err != nil {
			fmt.Println("switching error", err)
			batchErr = err
		}
	}
	return batchErr
}

func (self *Switcher) SwitchInputToSink(inputId int, sinkId int) error {
	return self.pactlCommand.MoveInput(inputId, sinkId)
}
