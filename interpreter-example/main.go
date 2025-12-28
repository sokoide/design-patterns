package main

import (
	"fmt"
	"interpreter-example/adapter"
)

func main() {
	fmt.Println("=== Interpreter Pattern ===")

	// Rule: Robert and John are male
	robert := &adapter.TerminalExpression{Data: "Robert"}
	john := &adapter.TerminalExpression{Data: "John"}
	isMale := &adapter.OrExpression{Expr1: robert, Expr2: john}

	// Rule: Julie is a married woman
	julie := &adapter.TerminalExpression{Data: "Julie"}
	married := &adapter.TerminalExpression{Data: "Married"}
	isMarriedWoman := &adapter.AndExpression{Expr1: julie, Expr2: married}

	fmt.Printf("John is male? %v\n", isMale.Interpret("John"))
	fmt.Printf("Julie is a married woman? %v\n", isMarriedWoman.Interpret("Married Julie"))
}
