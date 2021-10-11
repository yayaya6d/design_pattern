package adapter

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type adapterTestSuite struct {
	suite.Suite
}

func TestAdapterTestSuite(t *testing.T) {
	suite.Run(t, new(adapterTestSuite))
}

func (s adapterTestSuite) TestOutputPower_GivenUSSocket_ReturUSPowerValue() {
	// arrange
	// use usSocket
	usSocket := NewUSSocket()

	usPower := NewUSPower(usSocket)
	expectedOutput := USPowerVal

	// act
	actualOutput := usPower.OutputPower()

	// assert
	s.Equal(expectedOutput, actualOutput)
}

func (s adapterTestSuite) TestOutputPower_GivenEUSocketWithAdapter_ReturnUSPowerValue() {
	// arrange
	// use euSocket with adapter
	euSocket := NewEUSocket()
	adapter := NewEU2USPowerAdapter(euSocket)

	usPower := NewUSPower(adapter)
	expectedOutput := USPowerVal

	// act
	actualOutput := usPower.OutputPower()

	// assert
	s.Equal(expectedOutput, actualOutput)
}
