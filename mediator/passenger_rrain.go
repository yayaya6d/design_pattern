package mediator

import "fmt"

type PassengerTrain struct {
	name     string
	mediator Mediator
}

func (t *PassengerTrain) Arrive() {
	if t.mediator.CanArrive(t) {
		fmt.Printf("passenger train %s arrive", t.name)
	} else {
		fmt.Printf("block passenger train %s arrive", t.name)
	}
}

func (t *PassengerTrain) Depart() {
	fmt.Printf("passenger train %s depart", t.name)
	t.mediator.NotifyAboutDeparture()
}

func (t *PassengerTrain) PermitArrival() {
	fmt.Printf("passenger train %s permit arrival", t.name)
	t.Arrive()
}
