package sumpack

type expressionSum struct {
	ValueOne float64
	ValueTwo float64
	Result   float64
}

func SumOfValues(valOne, valTwo float64) float64 {
	result := valOne + valTwo
	return result
}

func ResultOfValues(valOne, valTwo float64) expressionSum {
	expr := expressionSum{
		ValueOne: valOne,
		ValueTwo: valTwo,
		Result:   .0,
	}
	expr.Result = SumOfValues(expr.ValueOne, expr.ValueTwo)
	return expr
}
