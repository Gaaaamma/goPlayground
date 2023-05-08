package vertex

import "fmt"

type Vertex struct {
	X int
	Y int
}

func (v *Vertex) Reset(X int, Y int) {
	v.X = X
	v.Y = Y
}

func (v *Vertex) Adder(val int) {
	v.X += val
	v.Y += val
}

func (v Vertex) Show() {
	fmt.Println(v.X, v.Y)
}
