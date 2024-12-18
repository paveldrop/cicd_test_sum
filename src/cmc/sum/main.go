package main

import (
	"fmt"
	sum "sum/internal"
)

func main() {
	result := sum.ResultOfValues(1, 1)
	fmt.Printf("val1=%f\nval2=%f\nresult=%f", result.ValueOne, result.ValueTwo, result.Result)
}
