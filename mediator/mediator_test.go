package mediator

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type trainStationTestSuite struct {
	suite.Suite

	sut Mediator
}

func TestTrainStationTestSuite(t *testing.T) {
	suite.Run(t, new(trainStationTestSuite))
}

func (s *trainStationTestSuite) SetupTest() {
	s.sut = NewTrainStation()
}

func (s *trainStationTestSuite) TestCanArrive_GivenOneTrainArrive_ReturnFalse() {
	// arrange
	train1 := &PassengerTrain{
		mediator: s.sut,
	}
	train1.Arrive()

	train2 := &FreightTrain{
		mediator: s.sut,
	}

	// act
	actualResult := s.sut.CanArrive(train2)

	// assert
	s.False(actualResult)
}

func (s *trainStationTestSuite) TestCanArrive_GivenOneTrainArriveAndLeave_ReturnTrue() {
	// arrange
	train1 := &PassengerTrain{
		mediator: s.sut,
	}
	train1.Arrive()
	train1.Depart()

	// act
	actualResult := s.sut.CanArrive(train1)

	// assert
	s.True(actualResult)
}

func (s *trainStationTestSuite) TestWaitionTrainCount_GivenThreeTrainArrive_ReturnExpectedResult() {
	// arrange
	train1 := &PassengerTrain{
		mediator: s.sut,
	}
	train1.Arrive()

	train2 := &FreightTrain{
		mediator: s.sut,
	}
	train2.Arrive()

	train3 := &FreightTrain{
		mediator: s.sut,
	}
	train3.Arrive()

	// act
	actualResult := s.sut.WaitingTrainCount()

	// assert
	s.Equal(2, actualResult)
}
