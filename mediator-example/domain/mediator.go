package domain

type Train interface {
	Arrive()
	Depart()
	PermitArrival()
}

type Mediator interface {
	CanArrive(Train) bool
	NotifyAboutDeparture()
}
