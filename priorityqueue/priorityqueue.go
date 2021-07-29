package main

import "fmt"

type PriorityQueue struct {
	heap    []int
	hashMap map[int][]int
}

// Initialize Priority Queue
func Init(items []int) *PriorityQueue {

	result := &PriorityQueue{heap: make([]int, len(items)), hashMap: make(map[int][]int)}

	for i, v := range items {
		result.heap[i] = v
		result.MapAdd(v, i)
	}

	return result
}

// Add element into Hash Map, if duplicate is found, ignore
func (pq *PriorityQueue) MapAdd(key int, index int) {

	if pq.hashMap[key] == nil {
		pq.hashMap[key] = append(pq.hashMap[key], index)
	} else {
		duplicate := false

		for _, i := range pq.hashMap[key] {
			if i == index {
				duplicate = true
			}
		}

		if !duplicate {
			pq.hashMap[key] = append(pq.hashMap[key], index)
		}
	}
}

func (pq *PriorityQueue) bubbleDown(index int) {

	left := index*2 + 1
	right := index*2 + 2
	smallest := index

	if left < len(pq.heap) && left < smallest {
		smallest = left
	}
	if right < len(pq.heap) && right < smallest {
		smallest = right
	}
}

func main() {

	list := []int{5, 4, 3, 2, 3, 6, 7}

	myPQ := Init(list)

	fmt.Println(myPQ)

}
