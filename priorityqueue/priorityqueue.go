package priorityqueue

import (
	"errors"
)

const (
	capacity = 10
)

// PriorityQueue represents Priority Queue data structure that holds a heap
// and a map that holds values as key and list of indexes as values
type PriorityQueue struct {
	heap         []int
	hashMap      map[int][]int
	heapSize     int
	heapCapacity int
}

// Initialize Priority Queue, and heapify so it satisfy Heap Invariant
func Init(items []int) *PriorityQueue {

	result := &PriorityQueue{heap: make([]int, 0, max(len(items), capacity)), hashMap: make(map[int][]int)}
	result.heapSize = len(items)
	result.heapCapacity = cap(result.heap)

	for i, v := range items {
		result.heap = append(result.heap, v)
		result.mapAdd(v, i)
	}

	// Heapify Process, O(n)
	for i := max(0, len(result.heap)); i >= 0; i-- {
		result.bubbleDown(i)
	}

	return result
}

func (pq *PriorityQueue) Size() int {
	return pq.heapSize
}

func (pq *PriorityQueue) Clear() {

	pq.heap = make([]int, 0, capacity)
	pq.hashMap = make(map[int][]int)
	pq.heapSize = 0
	pq.heapCapacity = capacity
}

func (pq *PriorityQueue) IsEmpty() bool {
	return pq.heapSize == 0
}

// Check if element is in Heap, O(1)
func (pq *PriorityQueue) Contain(element int) (bool, error) {
	if pq.IsEmpty() {
		return false, errors.New("priority queue is empty")
	}

	_, ok := pq.hashMap[element]

	return ok, nil
}

func (pq *PriorityQueue) Peek() (int, error) {
	if pq.IsEmpty() {
		return 0, errors.New("priority queue is empty")
	}

	return pq.heap[0], nil
}

func (pq *PriorityQueue) Dequeue() (int, error) {
	return pq.RemoveAt(0)
}

// Add element into Heap, O(log n), O(n) if resizing occurs
func (pq *PriorityQueue) Enqueue(element int) {

	if pq.heapSize < cap(pq.heap) {
		pq.heap = append(pq.heap, element)
	} else {
		pq.heap = append(pq.heap, element)
		temp := make([]int, 0, pq.heapCapacity+10)
		temp = append(temp, pq.heap...)
		pq.heap = temp
		pq.heapCapacity = cap(pq.heap)
	}

	pq.mapAdd(element, pq.heapSize)
	pq.floatUp(pq.heapSize)
	pq.heapSize++
}

// RemoveAt method, removes element based on index, O(log n), O(n) if resizing occurs
func (pq *PriorityQueue) RemoveAt(index int) (int, error) {
	if pq.IsEmpty() {
		return 0, errors.New("priority queue is empty")
	}

	pq.heapSize--
	removedData := pq.heap[index]
	pq.swap(index, pq.heapSize)
	pq.heap = pq.heap[:pq.heapSize]

	if pq.heapSize < pq.heapCapacity-10 {
		temp := make([]int, 0, pq.heapCapacity-10)
		temp = append(temp, pq.heap...)
		pq.heap = temp
		pq.heapCapacity = cap(pq.heap)
	}

	pq.mapRemove(removedData, pq.heapSize)

	if index == pq.heapSize {
		return removedData, nil
	}

	element := pq.heap[index]
	pq.bubbleDown(index)

	if pq.heap[index] == element {
		pq.floatUp(index)
	}

	return removedData, nil
}

// Remove method, removes specified element in Heap, O(log n)
func (pq *PriorityQueue) Remove(element int) (int, error) {

	index, err := pq.mapGet(element)
	if err == nil {
		return pq.RemoveAt(index)
	} else {
		return 0, err
	}
}

// Check if Heap Invariant is satisfied
func (pq *PriorityQueue) IsMinHeap(index int) bool {

	if index >= pq.heapSize {
		return true
	}

	leftChild := index*2 + 1
	rightChild := index*2 + 1

	if leftChild < pq.heapSize && !pq.less(index, leftChild) {
		return false
	}
	if rightChild < pq.heapSize && !pq.less(index, rightChild) {
		return false
	}

	return pq.IsMinHeap(leftChild) && pq.IsMinHeap(rightChild)
}

// Add element into Hash Map, if duplicate is found, ignore
func (pq *PriorityQueue) mapAdd(key int, index int) {

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

func (pq *PriorityQueue) mapGet(element int) (int, error) {

	if _, ok := pq.hashMap[element]; !ok {
		return 0, errors.New("element not found")
	}

	index := pq.hashMap[element][len(pq.hashMap[element])-1]
	return index, nil
}

func (pq *PriorityQueue) floatUp(index int) {

	if index == 0 {
		return
	}

	parent := (index - 1) / 2

	if pq.less(index, parent) {
		pq.swap(index, parent)
		pq.floatUp(parent)
	}
}

// bubbleDown method, element bubble down to satisfy Heap Invariant
func (pq *PriorityQueue) bubbleDown(index int) {

	leftChild := index*2 + 1
	rightChild := index*2 + 2
	smallest := index

	if leftChild < len(pq.heap) && pq.less(leftChild, smallest) {
		smallest = leftChild
	}
	if rightChild < len(pq.heap) && pq.less(rightChild, smallest) {
		smallest = rightChild
	}

	if smallest != index {
		pq.swap(smallest, index)
		pq.bubbleDown(smallest)
	}
}

// swap method, swap places of two elements in Heap
func (pq *PriorityQueue) swap(i int, j int) {
	iElement := pq.heap[i]
	jElement := pq.heap[j]

	pq.heap[i], pq.heap[j] = jElement, iElement
	pq.mapSwap(iElement, jElement, i, j)
}

func (pq *PriorityQueue) mapRemove(element int, idx int) {

	for index, value := range pq.hashMap[element] {
		if value == idx {
			pq.hashMap[element] = append(pq.hashMap[element][:index], pq.hashMap[element][index+1:]...)
		}
	}

	if len(pq.hashMap[element]) == 0 {
		delete(pq.hashMap, element)
	}
}

// mapSwap method, swap indexes (value) mapped to values (key) when swapping occurs
func (pq *PriorityQueue) mapSwap(element1 int, element2 int, index1 int, index2 int) {

	for idx, index := range pq.hashMap[element1] {
		if index == index1 {
			pq.hashMap[element1] = append(pq.hashMap[element1][:idx], pq.hashMap[element1][idx+1:]...)
		}
	}

	for idx, index := range pq.hashMap[element2] {
		if index == index2 {
			pq.hashMap[element2] = append(pq.hashMap[element2][:idx], pq.hashMap[element2][idx+1:]...)
		}
	}

	pq.hashMap[element1] = append(pq.hashMap[element1], index2)
	pq.hashMap[element2] = append(pq.hashMap[element2], index1)
}

func (pq *PriorityQueue) less(i int, j int) bool {
	node1 := pq.heap[i]
	node2 := pq.heap[j]
	return compareTo(node1, node2) <= 0
}

func max(x int, y int) int {
	if x <= y {
		return y
	} else {
		return x
	}
}

func compareTo(i int, j int) int {

	if i < j {
		return -1
	} else if i > j {
		return 1
	} else {
		return 0
	}

}
