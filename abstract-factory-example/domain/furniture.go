package domain

type Chair interface {
	SitOn() string
}

type Sofa interface {
	LieOn() string
}

type FurnitureFactory interface {
	CreateChair() Chair
	CreateSofa() Sofa
}
