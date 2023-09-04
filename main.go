package main

import (
	"fmt"
	"micrograd-go/micrograd"
)

func main() {
	fmt.Println("Welcome to micrograd!")

	a := micrograd.NewValue(2.0)
	b := micrograd.NewValue(-3.0)
	c := a.Add(b)

	fmt.Printf("%s", c.String())
}
