package command

type TVPowerButton struct {
	curPower   bool
	turnOnCmd  SingleCommand
	turnOffCmd SingleCommand
}

func (b *TVPowerButton) Press() {
	if b.curPower {
		b.turnOffCmd.Execute()
	} else {
		b.turnOnCmd.Execute()
	}
	b.curPower = !b.curPower
}

type TVChannelButton struct {
	cmd NumCommand
}

func (b *TVChannelButton) ChangeChannel(ch int) {
	b.cmd.Execute(ch)
}
