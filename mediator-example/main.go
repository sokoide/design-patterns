package main

import (
	"mediator-example/adapter"
)

func main() {
	stationManager := adapter.NewStationManager()

	passengerTrain := &adapter.PassengerTrain{
		Mediator: stationManager,
	}
	freightTrain := &adapter.FreightTrain{
		Mediator: stationManager,
	}

	passengerTrain.Arrive()
	freightTrain.Arrive()
	passengerTrain.Depart()
}
