package adapter

import (
	"fmt"
	"mediator-example/domain"
)

type PassengerTrain struct {
	Mediator domain.Mediator
}

func (t *PassengerTrain) Arrive() {
	if t.Mediator == nil {
		fmt.Println("PassengerTrain: No mediator set")
		return
	}
	if !t.Mediator.CanArrive(t) {
		fmt.Println("PassengerTrain: Blocked, waiting")
		return
	}
	fmt.Println("PassengerTrain: Arrived")
}

func (t *PassengerTrain) Depart() {
	fmt.Println("PassengerTrain: Leaving")
	if t.Mediator != nil {
		t.Mediator.NotifyAboutDeparture()
	}
}

func (t *PassengerTrain) PermitArrival() {
	fmt.Println("PassengerTrain: Arrival permitted, arriving")
	t.Arrive()
}

type FreightTrain struct {
	Mediator domain.Mediator
}

func (t *FreightTrain) Arrive() {
	if t.Mediator == nil {
		fmt.Println("FreightTrain: No mediator set")
		return
	}
	if !t.Mediator.CanArrive(t) {
		fmt.Println("FreightTrain: Blocked, waiting")
		return
	}
	fmt.Println("FreightTrain: Arrived")
}

func (t *FreightTrain) Depart() {
	fmt.Println("FreightTrain: Leaving")
	if t.Mediator != nil {
		t.Mediator.NotifyAboutDeparture()
	}
}

func (t *FreightTrain) PermitArrival() {
	fmt.Println("FreightTrain: Arrival permitted, arriving")
	t.Arrive()
}
