package command

type Television interface {
	TurnOn()
	TurnOff()
	IsPowerOn() bool
	ChangeChannel(int)
	CurrentChannel() int
}

func NewTelevision() Television {
	return &television{
		power:   false,
		channel: 1,
	}
}

type television struct {
	power   bool
	channel int
}

func (t *television) TurnOn() {
	t.power = true
}

func (t *television) TurnOff() {
	t.power = false
}

func (t *television) IsPowerOn() bool {
	return t.power
}

func (t *television) ChangeChannel(c int) {
	if t.power {
		t.channel = c
	}
}

func (t *television) CurrentChannel() int {
	if t.power {
		return t.channel
	}
	return -1
}
