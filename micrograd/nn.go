package micrograd

import "fmt"

type Value struct {
	Data     float64
	previous []Value
}

func NewValue(x float64) *Value {
	return &Value{Data: x, previous: []Value{}}
}

func (value Value) String() string {
	return fmt.Sprintf("Value ( Data: %.3f )\n", value.Data)
}

func (value Value) Previous() {
	printout := "Previous:\n"
	for _, child := range value.previous {
		printout += child.String()
	}
	fmt.Printf("%s", printout)
}

func (value Value) Add(other *Value) *Value {
	out := value.Data + other.Data
	return &Value{Data: out, previous: []Value{value, *other}}
}

func (value Value) Mul(other *Value) *Value {
	out := value.Data * other.Data
	return &Value{Data: out, previous: []Value{value, *other}}
}
