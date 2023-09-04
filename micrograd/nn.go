package micrograd

import "fmt"

type Value struct {
	Data float64
}

func NewValue(x float64) *Value {
	return &Value{Data: x}
}

func (value Value) String() string {
	return fmt.Sprintf("Value ( Data: %.3f )", value.Data)
}

func (value Value) Add(other *Value) *Value {
	out := value.Data + other.Data
	return &Value{Data: out}
}
