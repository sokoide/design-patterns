package adapter

import (
	"fmt"
	"visitor-example/domain"
)

type Square struct {
	Side int
}

func (s *Square) Accept(v domain.Visitor) {
	v.VisitForSquare(s)
}
func (s *Square) GetType() string { return "Square" }

type Circle struct {
	Radius int
}

func (c *Circle) Accept(v domain.Visitor) {
	v.VisitForCircle(c)
}
func (c *Circle) GetType() string { return "Circle" }

type Rectangle struct {
	L, B int
}

func (t *Rectangle) Accept(v domain.Visitor) {
	v.VisitForRectangle(t)
}
func (t *Rectangle) GetType() string { return "Rectangle" }

type AreaCalculator struct {
	area int
}

func (a *AreaCalculator) VisitForSquare(s domain.Shape) {
	sq, ok := s.(*Square)
	if !ok {
		fmt.Printf("[WARN] VisitForSquare called with %T\n", s)
		return
	}
	fmt.Printf("Calculating area for square. Side: %d\n", sq.Side)
	a.area = sq.Side * sq.Side
	fmt.Printf("Area: %d\n", a.area)
}

func (a *AreaCalculator) VisitForCircle(s domain.Shape) {
	ci, ok := s.(*Circle)
	if !ok {
		fmt.Printf("[WARN] VisitForCircle called with %T\n", s)
		return
	}
	fmt.Printf("Calculating area for circle. Radius: %d\n", ci.Radius)
	a.area = 3 * ci.Radius * ci.Radius // approx pi
	fmt.Printf("Area: %d\n", a.area)
}

func (a *AreaCalculator) VisitForRectangle(s domain.Shape) {
	r, ok := s.(*Rectangle)
	if !ok {
		fmt.Printf("[WARN] VisitForRectangle called with %T\n", s)
		return
	}
	fmt.Printf("Calculating area for rectangle.\n")
	a.area = r.L * r.B
	fmt.Printf("Area: %d\n", a.area)
}
