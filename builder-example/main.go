package main

import (
	"builder-example/adapter"
	"builder-example/usecase"
	"fmt"
)

func main() {
	fmt.Println("=== Builder Pattern ===")

	normalBuilder := adapter.NewNormalBuilder()
	director1 := usecase.NewDirector(normalBuilder)
	house1 := director1.BuildHouse()
	fmt.Printf("Normal House: Door=%s, Window=%s, Floor=%d\n", house1.DoorType, house1.WindowType, house1.Floor)

	iglooBuilder := adapter.NewIglooBuilder()
	director2 := usecase.NewDirector(iglooBuilder)
	house2 := director2.BuildHouse()
	fmt.Printf("Igloo House: Door=%s, Window=%s, Floor=%d\n", house2.DoorType, house2.WindowType, house2.Floor)
}
