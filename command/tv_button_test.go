package command

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type tvButtonTestSuite struct {
	suite.Suite
}

func TestTvButtonTestSuite(t *testing.T) {
	suite.Run(t, new(tvButtonTestSuite))
}

func (s tvButtonTestSuite) TestPress_PressFirstTimeToOnTV_TVOn() {
	// arrange
	tv := NewTelevision()
	tvPowerButton := &TVPowerButton{
		curPower: false,
		turnOnCmd: &PowerOnCommand{
			tv: tv,
		},
		turnOffCmd: &PowerOffCommand{
			tv: tv,
		},
	}

	// act
	tvPowerButton.Press()

	// assert
	s.True(tv.IsPowerOn())
}

func (s tvButtonTestSuite) TestPress_PressSecondTimeToTurnOffTV_TVOff() {
	// arrange
	tv := NewTelevision()
	tvPowerButton := &TVPowerButton{
		curPower: false,
		turnOnCmd: &PowerOnCommand{
			tv: tv,
		},
		turnOffCmd: &PowerOffCommand{
			tv: tv,
		},
	}

	// act
	tvPowerButton.Press()
	tvPowerButton.Press()

	// assert
	s.False(tv.IsPowerOn())
}

func (s tvButtonTestSuite) TestChangeChannel_TurnOnAndChangeChannelTo7_TVChannelIs7() {
	// arrange
	tv := NewTelevision()
	tvPowerButton := &TVPowerButton{
		curPower: false,
		turnOnCmd: &PowerOnCommand{
			tv: tv,
		},
		turnOffCmd: &PowerOffCommand{
			tv: tv,
		},
	}

	tvChannelButton := &TVChannelButton{
		cmd: &ChangeChannelCommand{
			tv: tv,
		},
	}

	// act
	tvPowerButton.Press()
	tvChannelButton.ChangeChannel(7)

	// assert
	s.Equal(7, tv.CurrentChannel())
}

func (s tvButtonTestSuite) TestChangeChannel_TurnOffAndChangeChannelTo7_TVChannelIsDefault() {
	// arrange
	tv := NewTelevision()

	tvChannelButton := &TVChannelButton{
		cmd: &ChangeChannelCommand{
			tv: tv,
		},
	}

	// act
	tvChannelButton.ChangeChannel(7)

	// assert
	s.Equal(-1, tv.CurrentChannel())
}
