package main

import (
	"fmt"
	"gaaaamma/vertex"
	"math/rand"
	"sort"
)

func main() {
	vtx := []vertex.Vertex{}
	vtxGen := vertex.VertexGenerator()
	for i := 0; i < 10; i++ {
		vtx = append(vtx, vtxGen(rand.Intn(10)))
	}
	fmt.Println(vtx)
	sort.Slice(vtx, func(i, j int) bool {
		return vtx[i].P < vtx[j].P
	})
	fmt.Println(vtx)
}
