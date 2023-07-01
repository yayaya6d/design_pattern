package chain_of_responsibility

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type hosipitalTestSuite struct {
	suite.Suite
}

func TestHosipitalTestSuite(t *testing.T) {
	suite.Run(t, new(hosipitalTestSuite))
}

func (s *hosipitalTestSuite) TestExecute_GivenNewPatient_ReturnExpectedStatus() {
	// arrange
	hosipital := BuildHositital()
	patient := Patient{
		Name:              "patient 1",
		RegisterDone:      false,
		MedicineDone:      false,
		DoctorCheckupDone: false,
	}

	expectedPatientStatus := Patient{
		Name:              "patient 1",
		RegisterDone:      true,
		MedicineDone:      true,
		DoctorCheckupDone: true,
	}

	// act
	hosipital.Execute(&patient)

	// assert
	s.Equal(expectedPatientStatus, patient)
	s.Equal(1, hosipital.ExecuteCount())
}
