package switcher

type Switcher struct {
}

func NewSwitcher() *Switcher {
	return &Switcher{}
}

func (self *Switcher) SwitchAllToSink(sinkId int) error {
	return nil
}

func (self *Switcher) SwitchInputToSink(inputId int, sinkId int) {

}
