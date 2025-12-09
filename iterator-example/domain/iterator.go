package domain

type Iterator interface {
	HasNext() bool
	GetNext() interface{}
}

type Collection interface {
	CreateIterator() Iterator
}
