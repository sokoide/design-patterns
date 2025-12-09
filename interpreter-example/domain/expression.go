package domain

type Expression interface {
	Interpret(context string) bool
}
