package decorator

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type pizzaTestSuite struct {
	suite.Suite
}

func TestPizzaTestSuite(t *testing.T) {
	suite.Run(t, new(pizzaTestSuite))
}

func (s *pizzaTestSuite) TestGetPrize_GivenCheesePizza_ReturnExpectedPrice() {
	// arrange
	sut := Cheese{
		next: &Pizza{},
	}

	// act
	actualPrice := sut.GetPrice()

	// assert
	s.Equal(13, actualPrice)
}

func (s *pizzaTestSuite) TestGetPrize_GivenTomatoPizza_ReturnExpectedPrice() {
	// arrange
	sut := Tomato{
		next: &Pizza{},
	}

	// act
	actualPrice := sut.GetPrice()

	// assert
	s.Equal(15, actualPrice)
}

func (s *pizzaTestSuite) TestGetPrize_GivenCheeseTomatoPizza_ReturnExpectedPrice() {
	// arrange
	sut := Tomato{
		next: &Cheese{
			next: &Pizza{},
		},
	}

	// act
	actualPrice := sut.GetPrice()

	// assert
	s.Equal(18, actualPrice)
}

func (s *pizzaTestSuite) TestGetName_GivenCheesePizza_ReturnExpectedString() {
	// arrange
	sut := Cheese{
		next: &Pizza{},
	}

	// act
	actualPrice := sut.GetName()

	// assert
	s.Equal("cheese pizza", actualPrice)
}

func (s *pizzaTestSuite) TestGetName_GivenTomatoPizza_ReturnExpectedString() {
	// arrange
	sut := Tomato{
		next: &Pizza{},
	}

	// act
	actualPrice := sut.GetName()

	// assert
	s.Equal("tomato pizza", actualPrice)
}

func (s *pizzaTestSuite) TestGetName_GivenCheeseTomatoPizza_ReturnExpectedString() {
	// arrange
	sut := Tomato{
		next: &Cheese{
			next: &Pizza{},
		},
	}

	// act
	actualPrice := sut.GetName()

	// assert
	s.Equal("tomato cheese pizza", actualPrice)
}
