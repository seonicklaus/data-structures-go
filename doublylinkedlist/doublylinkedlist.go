package doublylinkedlist

import (
	"errors"
	"fmt"
	"strings"
)

// Node represents a Node in a Doubly LinkedList that holds
// data, previous node and next node
type node struct {
	data     int
	previous *node
	next     *node
}

type DoublyLinkedList struct {
	head *node
	tail *node
	size int
}

// Get size of Doubly LinkedList
func (l *DoublyLinkedList) Size() int {
	return l.size
}

// Check if Doubly Linkedlist is empty
func (l *DoublyLinkedList) IsEmpty() bool {
	return l.size == 0
}

// Get first element data, O(1)
func (l *DoublyLinkedList) PeekFirst() int {
	return l.head.data
}

// Get last element data, O(1)
func (l *DoublyLinkedList) PeekLast() int {
	return l.tail.data
}

// Get index of element (zero-based), O(n)
func (l *DoublyLinkedList) IndexOf(data int) (int, error) {
	if l.IsEmpty() {
		return 0, errors.New("linked list is empty")
	}

	travNode := l.head
	index := 0

	for travNode != nil {
		if travNode.data == data {
			return index, nil
		}

		travNode = travNode.next
		index++
	}

	return 0, errors.New("data not found in linked list")
}

// Get element data by index (negative index reverses order), O(1) for first and last element, O(n) for the rest
func (l *DoublyLinkedList) Get(idx int) (int, error) {
	if l.IsEmpty() {
		return 0, errors.New("linked list is empty")
	}

	indexToGet := 0

	if idx >= 0 {
		indexToGet = idx
	} else {
		indexToGet = l.size + idx
	}

	if indexToGet == 0 {
		return l.head.data, nil
	} else if indexToGet == l.size-1 {
		return l.tail.data, nil
	} else if indexToGet < 0 || indexToGet >= l.size {
		return 0, errors.New("index out of range")
	}

	travNode := l.head.next

	for index := 0; index != indexToGet; index++ {
		travNode = travNode.next
	}

	return travNode.data, nil
}

// Clear linked list, O(n)
func (l *DoublyLinkedList) Clear() {
	travNode := l.head

	for travNode != nil {
		next := travNode.next
		travNode.data = 0
		travNode.previous = nil
		travNode.next = nil
		travNode = next
	}

	l.head = nil
	l.tail = nil
	l.size = 0
}

// AddFirst, initializes Node and call Prepent method
func (l *DoublyLinkedList) AddFirst(data int) {
	l.prepend(&node{data: data, next: l.head})
}

// AddLast, initializes Node and call Append method
func (l *DoublyLinkedList) AddLast(data int) {
	l.append(&node{data: data, previous: l.tail})
}

// prepent, add Node to the head of the linked list, O(1)
func (l *DoublyLinkedList) prepend(node *node) {
	if l.IsEmpty() {
		l.head = node
		l.tail = node
	} else {
		l.head.previous = node
		l.head = node
	}

	l.size++
}

// append, add Node to the tail of the Linked List, O(1)
func (l *DoublyLinkedList) append(node *node) {
	if l.IsEmpty() {
		l.head = node
		l.tail = node
	} else {
		l.tail.next = node
		l.tail = node
	}

	l.size++
}

// RemoveFirst method, removes first node from Linked List, O(1)
func (l *DoublyLinkedList) RemoveFirst() (int, error) {
	if l.IsEmpty() {
		return 0, errors.New("linked list is empty")
	}

	data := l.head.data
	l.head = l.head.next
	l.head.previous.data = 0
	l.head.previous.next = nil
	l.head.previous = nil

	l.size--

	return data, nil
}

// RemoveLast method, removes the last Node from Linked List, O(1)
func (l *DoublyLinkedList) RemoveLast() (int, error) {
	if l.IsEmpty() {
		return 0, errors.New("linked list is empty")
	}

	data := l.tail.data
	l.tail = l.tail.previous
	l.tail.next.data = 0
	l.tail.next.previous = nil
	l.tail.next = nil

	l.size--

	return data, nil
}

// RemoveAt method, removes a Node at a specific index (negative index reverses order), O(log n)
func (l *DoublyLinkedList) RemoveAt(idx int) (int, error) {
	if l.IsEmpty() {
		return 0, errors.New("linked list is empty")
	}

	indexToRemove := 0

	if idx >= 0 {
		indexToRemove = idx
	} else {
		indexToRemove = l.size + idx
	}

	if indexToRemove < 0 || indexToRemove >= l.size {
		return 0, errors.New("index out of range")
	}

	travNode := l.head

	if indexToRemove < l.size/2 {

		for index := 0; index != indexToRemove; index++ {
			travNode = travNode.next
		}

	} else {
		travNode = l.tail

		for index := l.size - 1; index != indexToRemove; index-- {
			travNode = travNode.previous
		}

	}

	data := l.RemoveNode(travNode)

	return data, nil
}

// Remove arbitrary Node from Linked List, O(1)
func (l *DoublyLinkedList) RemoveNode(node *node) int {

	if node.previous == nil {
		data, _ := l.RemoveFirst()
		return data
	}

	if node.next == nil {
		data, _ := l.RemoveLast()
		return data
	}

	data := node.data

	node.previous.next = node.next
	node.next.previous = node.previous
	node.previous = nil
	node.next = nil

	node.data = 0

	l.size--

	return data
}

func (n node) String() string {
	return fmt.Sprintf("%d", n.data)
}

func (l DoublyLinkedList) String() string {
	sb := strings.Builder{}

	for iterator := l.head; iterator != nil; iterator = iterator.next {
		sb.WriteString(fmt.Sprintf("%s -> ", iterator))
	}

	sb.WriteString("nil")

	return sb.String()
}
