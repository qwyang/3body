package examples

import (
	"testing"
)

func TestCalc(t *testing.T){
	expr := "((1+1)*(1+1))*(1+1)"
	t.Log(normalizeExpr(expr))
	t.Log(NormalizeExpr(expr))
	data,_ := NormalizeExpr(expr)
	t.Log(infixToPosfix(data))
	t.Log(Calculate(expr))
}
