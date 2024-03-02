package mediator

import "fmt"

type FreightTrain struct {
	name     string
	mediator Mediator
}

func (t *FreightTrain) Arrive() {
	if t.mediator.CanArrive(t) {
		fmt.Printf("Freight train %s arrive", t.name)
	} else {
		fmt.Printf("block Freight train %s arrive", t.name)
	}
}

func (t *FreightTrain) Depart() {
	fmt.Printf("Freight train %s depart", t.name)
	t.mediator.NotifyAboutDeparture()
}

func (t *FreightTrain) PermitArrival() {
	fmt.Printf("Freight train %s permit arrival", t.name)
	t.Arrive()
}
