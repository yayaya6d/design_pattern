package command

type SingleCommand interface {
	Execute()
}

type PowerOnCommand struct {
	tv Television
}

func (c *PowerOnCommand) Execute() {
	c.tv.TurnOn()
}

type PowerOffCommand struct {
	tv Television
}

func (c *PowerOffCommand) Execute() {
	c.tv.TurnOff()
}

type NumCommand interface {
	Execute(int)
}

type ChangeChannelCommand struct {
	tv Television
}

func (c *ChangeChannelCommand) Execute(ch int) {
	c.tv.ChangeChannel(ch)
}
