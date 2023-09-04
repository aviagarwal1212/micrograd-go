package main

import (
	"fmt"
	"micrograd-go/micrograd"
)

func main() {
	fmt.Println("Welcome to micrograd!")

	a := micrograd.NewValue(2.0)
	b := micrograd.NewValue(-3.0)
	c := micrograd.NewValue(10.0)
	d := a.Mul(b).Add(c)

	fmt.Printf("%s", d.String())
	d.Previous()
}
