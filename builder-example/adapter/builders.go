package adapter

import "builder-example/domain"

// Normal House
type NormalBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func NewNormalBuilder() *NormalBuilder { return &NormalBuilder{} }

func (b *NormalBuilder) SetWindowType() { b.windowType = "Wooden Window" }
func (b *NormalBuilder) SetDoorType()   { b.doorType = "Wooden Door" }
func (b *NormalBuilder) SetNumFloor()   { b.floor = 2 }
func (b *NormalBuilder) GetHouse() domain.House {
	return domain.House{
		WindowType: b.windowType,
		DoorType:   b.doorType,
		Floor:      b.floor,
	}
}

// Igloo
type IglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func NewIglooBuilder() *IglooBuilder { return &IglooBuilder{} }

func (b *IglooBuilder) SetWindowType() { b.windowType = "Snow Window" }
func (b *IglooBuilder) SetDoorType()   { b.doorType = "Snow Door" }
func (b *IglooBuilder) SetNumFloor()   { b.floor = 1 }
func (b *IglooBuilder) GetHouse() domain.House {
	return domain.House{
		WindowType: b.windowType,
		DoorType:   b.doorType,
		Floor:      b.floor,
	}
}
