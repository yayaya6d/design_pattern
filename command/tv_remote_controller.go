package command

type TVRemoteController struct {
	turnOnCmd  SingleCommand
	turnOffCmd SingleCommand
	numCmd     NumCommand
}

func (c *TVRemoteController) TurnOnTV() {
	c.turnOnCmd.Execute()
}

func (c *TVRemoteController) TurnOffTV() {
	c.turnOffCmd.Execute()
}

func (c *TVRemoteController) ChangeChannel(ch int) {
	c.numCmd.Execute(ch)
}

func (c *TVRemoteController) TurnOnAndChangeChannel(ch int) {
	c.turnOnCmd.Execute()
	c.numCmd.Execute(ch)
}
