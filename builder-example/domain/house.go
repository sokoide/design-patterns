package domain

type House struct {
	WindowType string
	DoorType   string
	Floor      int
}

type HouseBuilder interface {
	SetWindowType()
	SetDoorType()
	SetNumFloor()
	GetHouse() House
}
