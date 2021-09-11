package singleton

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type singletonConfigTestSuite struct {
	suite.Suite
}

func TestSingletonConfigTestSuite(t *testing.T) {
	suite.Run(t, new(singletonConfigTestSuite))
}

func (s *singletonConfigTestSuite) TestGetConfig_GetConfigOneTime_GetConfigDefaultValCorrectly() {
	// act
	config := GetConfig()

	// assert
	s.Equal(config.ServerUrl, "localhost")
	s.Equal(config.Port, "8080")
}

func (s *singletonConfigTestSuite) TestGetConfig_GetConfigServalTime_OnlyCreateConfigOnce() {
	// act
	config := GetConfig()
	config2 := GetConfig()

	// assert
	s.True(config == config2)
}
