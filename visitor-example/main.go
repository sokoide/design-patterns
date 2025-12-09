package main

import (
	"fmt"
	"visitor-example/adapter"
)

func main() {
	fmt.Println("=== Visitor Pattern ===")

	square := &adapter.Square{Side: 2}
	circle := &adapter.Circle{Radius: 3}
	rectangle := &adapter.Rectangle{L: 2, B: 3}

	areaCalculator := &adapter.AreaCalculator{}

	square.Accept(areaCalculator)
	circle.Accept(areaCalculator)
	rectangle.Accept(areaCalculator)
}
