package micrograd

import (
	"fmt"
	"math"

	"golang.org/x/exp/slices"
)

type Value struct {
	Data float64

	grad     float64
	label    string
	previous []*Value
	op       string

	backward func()
}

type ExecutionGraph struct {
	topography []*Value
	visited    []*Value
}

func NewValue(x float64, label string) *Value {
	return &Value{Data: x, grad: 0.0, label: label, previous: []*Value{}, op: "", backward: func() {}}
}

func (self *Value) String() string {
	return fmt.Sprintf("%s ( Data: %.3f )", self.label, self.Data)
}

func (self *Value) Children() string {
	output := ""
	for _, child := range self.previous {
		output += child.String() + " "
	}
	return output
}

func (self *Value) Describe() {
	printout := "\n" + self.String() + "\n"
	printout += fmt.Sprintf("Grad: %.3f\n", self.grad)
	printout += "Children: " + self.Children() + "\n"
	printout += "Op: " + self.op + "\n"
	fmt.Printf("%s", printout)
}

func (self *Value) Add(other *Value, label string) *Value {
	out := &Value{Data: (self.Data + other.Data), label: label, previous: []*Value{self, other}, op: "+"}
	out.backward = func() {
		self.grad = 1.0 * out.grad
		other.grad = 1.0 * out.grad
	}
	return out
}

func (self *Value) Mul(other *Value, label string) *Value {
	out := &Value{Data: (self.Data * other.Data), label: label, previous: []*Value{self, other}, op: "*"}
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
		previous: []*Value{self},
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

func NewExecutionGraph(finalNode *Value) *ExecutionGraph {
	graph := &ExecutionGraph{
		topography: []*Value{},
		visited:    []*Value{},
	}
	graph.BuildTopo(finalNode)
	return graph
}

func (graph *ExecutionGraph) Display() {
	fmt.Println("Execution Graph:")
	for _, node := range graph.topography {
		fmt.Println(node.String())
	}
}

func (graph *ExecutionGraph) BuildTopo(node *Value) {
	fmt.Println()
	fmt.Println("Running for ", node)
	if !slices.Contains(graph.visited, node) {
		graph.visited = append(graph.visited, node)
		graph.visited = set(graph.visited)
		fmt.Println(node, " has children ", node.Children())
		for _, child := range node.previous {
			graph.BuildTopo(child)
		}
		graph.topography = append(graph.topography, node)
		fmt.Println("Appended ", node.String(), " to Topo")
	}
}

func set(sequence []*Value) []*Value {
	setSequence := []*Value{}
	for _, node := range sequence {
		if !slices.Contains(setSequence, node) {
			setSequence = append(setSequence, node)
		}
	}
	return setSequence
}
