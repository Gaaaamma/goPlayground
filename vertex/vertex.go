package vertex

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
	P int
}

func VertexGenerator() func(p int) Vertex {
	counter := -1
	return func(p int) Vertex {
		counter++
		return Vertex{X: counter, Y: counter, P: p}
	}
}

func (v *Vertex) Reset(X int, Y int, P int) {
	v.X = X
	v.Y = Y
	v.P = P
}

func (v *Vertex) Adder(val int) {
	v.X += val
	v.Y += val
}

func (v Vertex) Show() {
	fmt.Println(v.X, v.Y, v.P)
}
