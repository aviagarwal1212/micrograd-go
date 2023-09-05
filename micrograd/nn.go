package micrograd

import (
	"fmt"
	"math"
)

type Value struct {
	Data     float64
	grad     float64
	label    string
	previous []Value
	op       string
}

func NewValue(x float64, label string) *Value {
	return &Value{Data: x, grad: 0.0, label: label, previous: []Value{}, op: ""}
}

func (value Value) String() string {
	return fmt.Sprintf("Value '%s'( Data: %.3f )\n", value.label, value.Data)
}

func (value Value) Describe() {
	printout := value.String()
	printout += fmt.Sprintf("Grad: %.3f\n", value.grad)
	printout += "Previous:\n"
	for _, child := range value.previous {
		printout += child.String()
	}
	printout += fmt.Sprintf("Op: %s\n", value.op)
	fmt.Printf("%s", printout)
}

func (value Value) Add(other *Value, label string) *Value {
	out := value.Data + other.Data
	return &Value{Data: out, label: label, previous: []Value{value, *other}, op: "+"}
}

func (value Value) Mul(other *Value, label string) *Value {
	out := value.Data * other.Data
	return &Value{Data: out, label: label, previous: []Value{value, *other}, op: "*"}
}

func (value Value) Tanh(label string) *Value {
	n := value.Data
	t := (math.Exp(2*n) - 1) / (math.Exp(2*n) + 1)
	return &Value{
		Data:     t,
		label:    label,
		previous: []Value{value},
		op:       "tanh",
	}
}
