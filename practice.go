package main

import (
	"fmt"
	"gaaaamma/vertex"
	"sort"
)

func main() {
	vertexes := []vertex.Vertex{{X: 1, Y: 5}, {X: 2, Y: 4}, {X: 3, Y: 6}}
	fmt.Println(vertexes)
	sort.Slice(vertexes, func(i, j int) bool {
		return vertexes[i].Y < vertexes[j].Y
	})
	fmt.Println(vertexes)

	numbers := []int{2, 5, 1, 3, 3, 0, 4, 6, 7, 9, 8}
	fmt.Println(numbers)
	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j]
	})
	fmt.Println("Asc:", numbers)
	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] > numbers[j]
	})
	fmt.Println("Desc:", numbers)

}
