package main

import (
	"fmt"
	"micrograd-go/micrograd"
)

func main() {
	fmt.Println("Welcome to micrograd!")

	a := micrograd.NewValue(2.0, "a")
	b := micrograd.NewValue(-3.0, "b")
	c := micrograd.NewValue(10.0, "c")
	e := a.Mul(b, "e")
	d := e.Add(c, "d")
	f := micrograd.NewValue(-2.0, "f")
	L := d.Mul(f, "L")

	L.Describe()

	// inputs x1 and x2
	x1 := micrograd.NewValue(2.0, "x1")
	x2 := micrograd.NewValue(0.0, "x2")

	// weights w1 and w2
	w1 := micrograd.NewValue(-3.0, "w1")
	w2 := micrograd.NewValue(1.0, "w2")

	// bias
	b = micrograd.NewValue(6.7, "b")

	// x1*w1 + x2*w2 + b
	x1w1 := x1.Mul(w1, "x1*w1")
	x2w2 := x2.Mul(w2, "x2*w2")
	x1w1x2w2 := x1w1.Add(x2w2, "x1*w1 + x2*w2")
	n := x1w1x2w2.Add(b, "n")
	o := n.Tanh("o")
	o.Describe()

}
