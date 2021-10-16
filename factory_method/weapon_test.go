package factory_method

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type FactoryMethodTestSutie struct {
	suite.Suite
}

func TestFactoryMethodTestSuite(t *testing.T) {
	suite.Run(t, new(FactoryMethodTestSutie))
}

func (s *FactoryMethodTestSutie) TestGetWeapon_GivenDifferentString_ReturnCorrespondingWeapon() {
	// arrange
	testCases := []struct {
		input    string
		expected string
	}{
		{"sword", "sword"},
		{"bow", "bow"},
		{"_", "empty-handed"},
		{"777", "empty-handed"},
	}

	// act and assert
	for _, testCase := range testCases {
		s.T().Run(fmt.Sprintf("input=%s", testCase.input), func(t *testing.T) {
			weapon := GetWeapon(testCase.input)
			s.Equal(testCase.expected, weapon.GetName())
		})
	}
}
