package singlylinkedlist

import (
	"errors"
	"fmt"
	"strings"
)

// Struct representation of a Node in a LinkedList
type node struct {
	value int
	next  *node
}

// Struct representation of a LinkedList
type SinglyLinkedList struct {
	head *node
	tail *node
	size int
}

// Get size of LinkedList
func (l SinglyLinkedList) GetSize() int {
	return l.size
}

// Check if LinkedList is empty
func (l SinglyLinkedList) IsEmpty() bool {
	return l.size == 0
}

// Get first element of Linked List, O(1)
func (l SinglyLinkedList) PeekFirst() int {
	return l.head.value
}

// Get last element of LinkedList, O(1)
func (l SinglyLinkedList) PeekLast() int {
	return l.tail.value
}

// Get index of an element (zero-based), O(n)
func (l SinglyLinkedList) IndexOf(value int) (int, error) {

	if l.IsEmpty() {
		return 0, errors.New("linked list is empty")
	}

	travNode := l.head
	index := 0

	for travNode != nil {

		if travNode.value == value {
			return index, nil
		} else {
			travNode = travNode.next
			index++
		}
	}

	return 0, errors.New("data not found in linked list")
}

// Get value of an element based on index (negative index reverses order), O(1) for first and last index, O(n) for in between
func (l SinglyLinkedList) Get(idx int) (int, error) {

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
		return l.head.value, nil
	} else if indexToGet == l.size-1 {
		return l.tail.value, nil
	} else if indexToGet < 0 || indexToGet >= l.size {
		return 0, errors.New("index out of range")
	}

	travNode := l.head
	index := 0

	for index < indexToGet {
		travNode = travNode.next
		index++
	}

	return travNode.value, nil
}

// Delete all element in Linked List, O(n)
func (l *SinglyLinkedList) Clear() {

	travNode := l.head

	for travNode != nil {

		temp := travNode.next
		travNode.value = 0
		travNode.next = nil
		travNode = temp
		l.size--

	}
}

// Adds element into Linked List, O(1)
func (l *SinglyLinkedList) Add(value int) {

	if l.IsEmpty() {

		l.head = &node{value: value}
		l.tail = l.head

	} else {

		l.tail.next = &node{value: value}
		l.tail = l.tail.next

	}

	l.size++
}

// Removes first element in Linked List, O(1)
func (l *SinglyLinkedList) RemoveFirst() (int, error) {

	if l.IsEmpty() {
		return 0, errors.New("linked list is empty")
	}

	tempNode := l.head
	data := tempNode.value
	l.head = tempNode.next
	tempNode.value = 0
	tempNode.next = nil
	l.size--

	return data, nil

}

// Removes last element in Linked List, O(1)
func (l *SinglyLinkedList) RemoveLast() (int, error) {

	if l.IsEmpty() {
		return 0, errors.New("linked list is empty")
	}

	travNode := l.head
	for travNode.next != l.tail {
		travNode = travNode.next
	}

	data := l.tail.value
	l.tail.value = 0
	travNode.next = nil
	l.tail = travNode

	l.size--

	return data, nil

}

// Removes element based on index (negative index reverses order), O(1) for first and last index, O(n) for in-between
func (l *SinglyLinkedList) RemoveAt(idx int) (int, error) {

	if l.IsEmpty() {
		return 0, errors.New("linked list is empty")
	}

	indexToRemove := 0

	if idx >= 0 {
		indexToRemove = idx
	} else {
		indexToRemove = l.size + idx
	}

	if indexToRemove == 0 {

		data, _ := l.RemoveFirst()
		return data, nil

	} else if indexToRemove == l.size-1 {

		data, _ := l.RemoveLast()
		return data, nil

	} else if indexToRemove < 0 || indexToRemove >= l.size {
		return 0, errors.New("index out of range")
	}

	travNode1 := l.head
	travNode2 := l.head.next
	index := 0

	for ; index < indexToRemove-1; index++ {
		travNode1 = travNode2
		travNode2 = travNode1.next
	}

	if index == indexToRemove-1 {
		travNode2 = travNode2.next
	}

	data := travNode1.next.value
	travNode1.next.value = 0
	travNode1.next.next = nil
	travNode1.next = travNode2
	l.size--

	return data, nil
}

// String representation of a Node
func (n *node) String() string {
	return fmt.Sprintf("%d", n.value)
}

// String representation of a Linked List
func (l SinglyLinkedList) String() string {

	if l.IsEmpty() {
		return "nil"
	}

	sb := strings.Builder{}

	for iterator := l.head; iterator != nil; iterator = iterator.next {
		sb.WriteString(fmt.Sprintf("%s -> ", iterator))
	}

	sb.WriteString("nil")

	return sb.String()
}
