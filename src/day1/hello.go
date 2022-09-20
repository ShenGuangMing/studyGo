package main

import (
	"basic/src/function"
	"fmt"
	"gonum.org/v1/gonum/stat"
)

func sub(a int, b int) int {
	return a - b
}

func main() {
	a := 10
	b := 20
	c := function.Add(a, b)
	fmt.Println("hello world", c)
	arr := []float64{1, 2, 3, 4, 5, 6}
	v := stat.Variance(arr, nil)
	fmt.Printf("方差=%f\n", v)
}
