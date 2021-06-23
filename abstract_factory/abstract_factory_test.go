package abstract_factory

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type AbstrcatFactoryTestSuite struct {
	suite.Suite
}

func (s *AbstrcatFactoryTestSuite) TestWinGUIFactory_UseWinGUIFactoryCreateObject_ObjectPlatformIsWin() {
	winGUIFactory := NewGUIFactory(Windows)
	winButton := winGUIFactory.CreateNewButton("")
	winCheckBox := winGUIFactory.CreateNewCheckBox()

	s.Equal(Windows, winButton.Platform())
	s.Equal(Windows, winCheckBox.Platform())
}

func (s *AbstrcatFactoryTestSuite) TestMacGUIFactory_UseMacGUIFactoryCreateObject_ObjectPlatformIsMac() {
	macGUIFactory := NewGUIFactory(Mac)
	macButton := macGUIFactory.CreateNewButton("")
	macCheckBox := macGUIFactory.CreateNewCheckBox()

	s.Equal(Mac, macButton.Platform())
	s.Equal(Mac, macCheckBox.Platform())
}

func TestAbstrcatFactoryTestSuitee(t *testing.T) {
	suite.Run(t, new(AbstrcatFactoryTestSuite))
}
