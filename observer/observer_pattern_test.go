package observer

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type observerPatternTestSuite struct {
	suite.Suite
}

func TestSingletonConfigTestSuite(t *testing.T) {
	suite.Run(t, new(observerPatternTestSuite))
}

func (s *observerPatternTestSuite) TestObserverPattern_RegisterAndNotify_NotifyRegisterObservers() {
	// arrange
	item1 := NewItem("item1", false)
	customer1 := NewCustomer("customer1")
	customer2 := NewCustomer("customer2")

	item1.Register(customer1)
	item1.Register(customer2)

	expectedOutput := "Sending email to customer customer1 for item item1, its avaliablity : false\n" +
		"Sending email to customer customer2 for item item1, its avaliablity : false\n"

	// act
	actualOutput := s.getOutput(item1.NotifyAllObservers)

	// assert
	s.Equal(expectedOutput, actualOutput)
}

func (s *observerPatternTestSuite) TestObserverPattern_RegisterAndDeRegister_DoNotNotifyDeRegisterObserver() {
	// arrange
	item1 := NewItem("item1", false)
	customer1 := NewCustomer("customer1")
	customer2 := NewCustomer("customer2")

	item1.Register(customer1)
	item1.Register(customer2)

	item1.DeRegister(customer1)

	expectedOutput := "Sending email to customer customer2 for item item1, its avaliablity : false\n"

	// act
	actualOutput := s.getOutput(item1.NotifyAllObservers)

	// assert
	s.Equal(expectedOutput, actualOutput)
}

func (s *observerPatternTestSuite) TestObserverPattern_RegisterAndUpdateAvailability_NotifyRegisterObservers() {
	// arrange
	item1 := NewItem("item1", false)
	customer1 := NewCustomer("customer1")

	item1.Register(customer1)

	expectedOutput := "Sending email to customer customer1 for item item1, its avaliablity : true\n"

	// act
	actualOutput := s.getOutput(func() {
		item1.UpdateAvailability(true)
	})

	// assert
	s.Equal(expectedOutput, string(actualOutput))
}

func (s *observerPatternTestSuite) getOutput(f func()) string {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	return string(out)
}
