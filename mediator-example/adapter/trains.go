package adapter

import (
	"fmt"
	"mediator-example/domain"
)

type PassengerTrain struct {
	Mediator domain.Mediator
}

func (g *PassengerTrain) Arrive() {
	if g.Mediator == nil {
		fmt.Println("PassengerTrain: No mediator set")
		return
	}
	if !g.Mediator.CanArrive(g) {
		fmt.Println("PassengerTrain: Blocked, waiting")
		return
	}
	fmt.Println("PassengerTrain: Arrived")
}

func (g *PassengerTrain) Depart() {
	fmt.Println("PassengerTrain: Leaving")
	if g.Mediator != nil {
		g.Mediator.NotifyAboutDeparture()
	}
}

func (g *PassengerTrain) PermitArrival() {
	fmt.Println("PassengerTrain: Arrival permitted, arriving")
	g.Arrive()
}

type FreightTrain struct {
	Mediator domain.Mediator
}

func (g *FreightTrain) Arrive() {
	if g.Mediator == nil {
		fmt.Println("FreightTrain: No mediator set")
		return
	}
	if !g.Mediator.CanArrive(g) {
		fmt.Println("FreightTrain: Blocked, waiting")
		return
	}
	fmt.Println("FreightTrain: Arrived")
}

func (g *FreightTrain) Depart() {
	fmt.Println("FreightTrain: Leaving")
	if g.Mediator != nil {
		g.Mediator.NotifyAboutDeparture()
	}
}

func (g *FreightTrain) PermitArrival() {
	fmt.Println("FreightTrain: Arrival permitted, arriving")
	g.Arrive()
}
