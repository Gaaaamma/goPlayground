package main

import (
	"container/heap"
	"fmt"
	"gaaaamma/vertex"
)

type PQ_Vertex []vertex.Vertex

func (pq PQ_Vertex) Len() int {
	return len(pq)
}

func (pq PQ_Vertex) Less(i, j int) bool {
	return pq[i].P < pq[j].P
}

func (pq PQ_Vertex) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PQ_Vertex) Push(x any) {
	(*pq) = append((*pq), x.(vertex.Vertex))
}

func (pq *PQ_Vertex) Pop() interface{} {
	n := len(*pq)
	item := (*pq)[n-1]
	(*pq) = (*pq)[:n-1]
	return item
}

func main() {
	vtxGen := vertex.VertexGenerator()
	vtxPq := PQ_Vertex{}
	heap.Init(&vtxPq) // Only *PQ_Vertex implement all method of heap.interface Ref: https://blog.csdn.net/timemachine119/article/details/54927121
	fmt.Println("Init priority queue:", vtxPq)
	for i, counter := 5, 0; i > 0; i-- {
		heap.Push(&vtxPq, vtxGen(i*2))
		counter++
		fmt.Println("Round", counter, ":", vtxPq)
	}
	fmt.Println()

	// Get element from priority queue
	for len(vtxPq) > 0 {
		fmt.Println("Get element:", heap.Pop(&vtxPq).(vertex.Vertex))
		fmt.Println("After that:", vtxPq)
		fmt.Println()
	}
}
