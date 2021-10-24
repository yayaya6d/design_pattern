package builder

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type builderTestSuite struct {
	suite.Suite
}

func TestBuilderTestSuite(t *testing.T) {
	suite.Run(t, new(builderTestSuite))
}

func (s *builderTestSuite) TestGetCarBuilder_GivenBenzBuilder_ReturnExpectedCar() {
	// arrange
	director := NewDirector(&HondaCarBuilder{})

	// assert
	s.Equal(&HondaCarBuilder{}, director.GetCurCarBuilder())
}

func (s *builderTestSuite) TestSetCarBuilder_GivenBenzBuilder_ReturnExpectedCar() {
	// arrange
	director := NewDirector(&HondaCarBuilder{})

	// act
	director.SetCarBuilder(&BenzCarBuilder{})

	// assert
	s.Equal(&BenzCarBuilder{}, director.GetCurCarBuilder())
}

func (s *builderTestSuite) TestMakeSportCar_MakeSportCarWithBenzBuilder_ReturnExpectedCar() {
	// arrange
	director := NewDirector(&BenzCarBuilder{})

	// act
	acturalREsult := director.MakeSportCar()

	// assert
	expectedResult := Car{
		SeatCount: 2,
		Engine: Engine{
			name:  "Benz 777",
			power: 10,
		},
		Skin: Skin{
			color:    White,
			hardness: 15,
		},
	}
	s.Equal(expectedResult, acturalREsult)
}

func (s *builderTestSuite) TestMakeSUVCar_MakeSUVCarWithBenzBuilder_ReturnExpectedCar() {
	// arrange
	director := NewDirector(&BenzCarBuilder{})

	// act
	acturalREsult := director.MakeSUVCar()

	// assert
	expectedResult := Car{
		SeatCount: 6,
		Engine: Engine{
			name:  "Benz 777",
			power: 10,
		},
		Skin: Skin{
			color:    White,
			hardness: 15,
		},
	}
	s.Equal(expectedResult, acturalREsult)
}

func (s *builderTestSuite) TestMakeSportCar_MakeSportCarWithHondaBuilder_ReturnExpectedCar() {
	// arrange
	director := NewDirector(&HondaCarBuilder{})

	// act
	acturalREsult := director.MakeSportCar()

	// assert
	expectedResult := Car{
		SeatCount: 2,
		Engine: Engine{
			name:  "Honda 666",
			power: 8,
		},
		Skin: Skin{
			color:    Red,
			hardness: 20,
		},
	}
	s.Equal(expectedResult, acturalREsult)
}

func (s *builderTestSuite) TestMakeSUVCar_MakeSUVCarWithHondaBuilder_ReturnExpectedCar() {
	// arrange
	director := NewDirector(&HondaCarBuilder{})

	// act
	acturalREsult := director.MakeSUVCar()

	// assert
	expectedResult := Car{
		SeatCount: 6,
		Engine: Engine{
			name:  "Honda 666",
			power: 8,
		},
		Skin: Skin{
			color:    Red,
			hardness: 20,
		},
	}
	s.Equal(expectedResult, acturalREsult)
}
