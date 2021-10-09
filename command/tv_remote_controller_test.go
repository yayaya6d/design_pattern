package command

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type tvRemoteControllerTestSuite struct {
	suite.Suite
}

func TestTvRemoteControllerTestSuite(t *testing.T) {
	suite.Run(t, new(tvRemoteControllerTestSuite))
}

func (s tvRemoteControllerTestSuite) TestTurnOnTV_TurnOnTV_TVOn() {
	// arrange
	tv := NewTelevision()
	tvRemoteController := TVRemoteController{
		turnOnCmd: &PowerOnCommand{
			tv: tv,
		},
		turnOffCmd: &PowerOffCommand{
			tv: tv,
		},
		numCmd: &ChangeChannelCommand{
			tv: tv,
		},
	}

	// act
	tvRemoteController.TurnOnTV()

	// assert
	s.True(tv.IsPowerOn())
}

func (s tvRemoteControllerTestSuite) TestTurnOffTV_TurnOffTV_TVOff() {
	// arrange
	tv := NewTelevision()
	tvRemoteController := TVRemoteController{
		turnOnCmd: &PowerOnCommand{
			tv: tv,
		},
		turnOffCmd: &PowerOffCommand{
			tv: tv,
		},
		numCmd: &ChangeChannelCommand{
			tv: tv,
		},
	}

	// act
	tvRemoteController.TurnOnTV()
	tvRemoteController.TurnOffTV()

	// assert
	s.False(tv.IsPowerOn())
}

func (s tvRemoteControllerTestSuite) TestChannelCommand_TurnOnAndChangeChannelTo7_TVChannelIs7() {
	// arrange
	tv := NewTelevision()
	tvRemoteController := TVRemoteController{
		turnOnCmd: &PowerOnCommand{
			tv: tv,
		},
		turnOffCmd: &PowerOffCommand{
			tv: tv,
		},
		numCmd: &ChangeChannelCommand{
			tv: tv,
		},
	}

	// act
	tvRemoteController.TurnOnTV()
	tvRemoteController.ChangeChannel(7)

	// assert
	s.Equal(7, tv.CurrentChannel())
}

func (s tvRemoteControllerTestSuite) TestChannelCommand_TurnOffAndChangeChannelTo7_TVChannelIsDefault() {
	// arrange
	tv := NewTelevision()
	tvRemoteController := TVRemoteController{
		turnOnCmd: &PowerOnCommand{
			tv: tv,
		},
		turnOffCmd: &PowerOffCommand{
			tv: tv,
		},
		numCmd: &ChangeChannelCommand{
			tv: tv,
		},
	}

	// act
	tvRemoteController.ChangeChannel(7)

	// assert
	s.Equal(-1, tv.CurrentChannel())
}

func (s tvRemoteControllerTestSuite) TestTurnOnAndChannelCommand_TurnOnAndChangeChannelTo7_TVChannelIsDefault() {
	// arrange
	tv := NewTelevision()
	tvRemoteController := TVRemoteController{
		turnOnCmd: &PowerOnCommand{
			tv: tv,
		},
		turnOffCmd: &PowerOffCommand{
			tv: tv,
		},
		numCmd: &ChangeChannelCommand{
			tv: tv,
		},
	}

	// act
	tvRemoteController.TurnOnAndChangeChannel(7)

	// assert
	s.Equal(7, tv.CurrentChannel())
}
