package main

import (
	"fmt"
	"micrograd-go/micrograd"
)

func main() {
	fmt.Println("Welcome to micrograd!")

	// inputs x1 and x2
	x1 := micrograd.NewValue(2.0, "x1")
	x2 := micrograd.NewValue(0.0, "x2")

	// weights w1 and w2
	w1 := micrograd.NewValue(-3.0, "w1")
	w2 := micrograd.NewValue(1.0, "w2")

	// bias
	b := micrograd.NewValue(6.8813735870195432, "b")

	// x1*w1 + x2*w2 + b
	x1w1 := x1.Mul(w1, "x1*w1")
	x2w2 := x2.Mul(w2, "x2*w2")
	x1w1x2w2 := x1w1.Add(x2w2, "x1*w1 + x2*w2")
	n := x1w1x2w2.Add(b, "n")
	o := n.Tanh("o")

	// initialize the gradient of the final node
	o.SetGradient(1.0)
	o.Describe()

	// execute the backward pass for o and check n
	o.BackwardPass()
	n.Describe()

	// execute the backward pass for n and check x1w1x2w2 and b
	n.BackwardPass()
	x1w1x2w2.Describe()
	b.Describe()

	// execute backward pass for b
	b.BackwardPass()

	// execute backward pass for x1w1x2w2 and check x1w1 and x2w2
	x1w1x2w2.BackwardPass()
	x1w1.Describe()
	x2w2.Describe()

	// execute backward pass for x1w1 and check x1 and w1
	x1w1.BackwardPass()
	x1.Describe()
	w1.Describe()

	// execute backward pass for x2w2 and check x2 and w2
	x2w2.BackwardPass()
	x2.Describe()
	w2.Describe()

	fmt.Println()

	// build a topological sorted graph
	graph := micrograd.NewExecutionGraph(o)
	fmt.Println()
	graph.Display()

}
