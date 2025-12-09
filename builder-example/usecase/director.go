package usecase

import "builder-example/domain"

type Director struct {
	builder domain.HouseBuilder
}

func NewDirector(b domain.HouseBuilder) *Director {
	return &Director{builder: b}
}

func (d *Director) BuildHouse() domain.House {
	d.builder.SetWindowType()
	d.builder.SetDoorType()
	d.builder.SetNumFloor()
	return d.builder.GetHouse()
}
