package domain

type Shape interface {
	GetType() string
	Accept(Visitor)
}

type Visitor interface {
	VisitForSquare(Shape)
	VisitForCircle(Shape)
	VisitForRectangle(Shape)
}
