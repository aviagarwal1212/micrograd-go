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
	backward func()
}

func NewValue(x float64, label string) *Value {
	return &Value{Data: x, grad: 0.0, label: label, previous: []Value{}, op: "", backward: func() {}}
}

func (self *Value) String() string {
	return fmt.Sprintf("Value '%s'( Data: %.3f )\n", self.label, self.Data)
}

func (self *Value) Describe() {
	printout := "\n" + self.String()
	printout += fmt.Sprintf("Grad: %.3f\n", self.grad)
	printout += "Previous:\n"
	for _, child := range self.previous {
		printout += child.String()
	}
	printout += fmt.Sprintf("Op: %s\n", self.op)
	fmt.Printf("%s", printout)
}

func (self *Value) Add(other *Value, label string) *Value {
	out := &Value{Data: (self.Data + other.Data), label: label, previous: []Value{*self, *other}, op: "+"}
	out.backward = func() {
		self.grad = 1.0 * out.grad
		other.grad = 1.0 * out.grad
	}
	return out
}

func (self *Value) Mul(other *Value, label string) *Value {
	out := &Value{Data: (self.Data * other.Data), label: label, previous: []Value{*self, *other}, op: "*"}
	out.backward = func() {
		self.grad = out.grad * other.Data
		other.grad = out.grad * self.Data
	}
	return out
}

func (self *Value) Tanh(label string) *Value {
	n := self.Data
	t := (math.Exp(2*n) - 1) / (math.Exp(2*n) + 1)
	out := &Value{
		Data:     t,
		label:    label,
		previous: []Value{*self},
		op:       "tanh",
	}
	out.backward = func() {
		self.grad = (1 - math.Pow(t, 2.0)) * out.grad
	}
	return out
}

func (self *Value) SetGradient(grad float64) {
	self.grad = grad
}

func (self *Value) BackwardPass() {
	self.backward()
}
