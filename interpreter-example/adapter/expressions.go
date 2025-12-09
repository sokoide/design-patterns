package adapter

import (
	"interpreter-example/domain"
	"strings"
)

type TerminalExpression struct {
	Data string
}

func (t *TerminalExpression) Interpret(context string) bool {
	return strings.Contains(context, t.Data)
}

type OrExpression struct {
	Expr1 domain.Expression
	Expr2 domain.Expression
}

func (o *OrExpression) Interpret(context string) bool {
	return o.Expr1.Interpret(context) || o.Expr2.Interpret(context)
}

type AndExpression struct {
	Expr1 domain.Expression
	Expr2 domain.Expression
}

func (a *AndExpression) Interpret(context string) bool {
	return a.Expr1.Interpret(context) && a.Expr2.Interpret(context)
}
