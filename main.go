package main

import (
	"fmt"
	"math"
)

func f(x float64) float64 {
	return 3*math.Pow(x, 2) - 4*x + 5
}

func main() {
	fmt.Println("Welcome to micrograd!")
	h := 0.00000001
	x := -3.0
	result := (f(x+h) - f(x)) / h
	fmt.Printf("%.2f", result)
}
