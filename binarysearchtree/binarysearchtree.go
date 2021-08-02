package binarysearchtree

import (
	"errors"
)

type node struct {
	data  int
	left  *node
	right *node
}

type BinarySearchTree struct {
	root      *node
	nodeCount uint
}

type stack struct {
	items []*node
	size  int
}

type queue struct {
	items []*node
	size  int
}

func (bst *BinarySearchTree) Size() uint {
	return bst.nodeCount
}

func (bst *BinarySearchTree) IsEmpty() bool {
	return bst.nodeCount == 0
}

func (bst *BinarySearchTree) Contains(element int) bool {
	return bst.contains(bst.root, element)
}

func (bst *BinarySearchTree) Add(element int) bool {

	if bst.Contains(element) {
		return false
	} else {
		bst.root = bst.add(bst.root, element)
		bst.nodeCount++
		return true
	}
}

func (bst *BinarySearchTree) Remove(element int) bool {
	if bst.Contains(element) {
		bst.root = bst.remove(bst.root, element)
		bst.nodeCount--
		return true
	}

	return false
}

func (bst *BinarySearchTree) GetHeight() int {
	return bst.height(bst.root)
}

func (bst *BinarySearchTree) add(n *node, element int) *node {

	if n == nil {
		n = &node{data: element}
	} else {

		if compareTo(element, n.data) > 0 {
			n.right = bst.add(n.right, element)
		} else {
			n.left = bst.add(n.left, element)
		}
	}

	return n
}

func (bst *BinarySearchTree) PrintTree(order string) ([]int, error) {
	switch order {
	case "preorder":
		return bst.preorder()
	case "inorder":
		return bst.inorder()
	case "postorder":
		return bst.postorder()
	case "levelorder":
		return bst.levelorder()
	default:
		return nil, errors.New("order is invalid")
	}
}

func (bst *BinarySearchTree) remove(n *node, element int) *node {

	if n == nil {
		return nil
	}

	cmp := compareTo(element, n.data)

	if cmp < 0 {
		n.left = bst.remove(n.left, element)
	} else if cmp > 0 {
		n.right = bst.remove(n.right, element)
	} else {

		if n.left == nil {
			rightChild := n.right

			n.data = 0
			n = nil

			return rightChild
		} else if n.right == nil {
			leftChild := n.left

			n.data = 0
			n = nil

			return leftChild
		} else {
			smallestRight := bst.digLeft(n.right)
			n.data = smallestRight.data
			n.right = bst.remove(n.right, smallestRight.data)
		}
	}

	return n
}

func (bst *BinarySearchTree) digLeft(n *node) *node {
	current := n

	for current.left != nil {
		current = current.left
	}

	return current
}

// func (bst *BinarySearchTree) digRight(n *node) *node {
// 	current := n

// 	for current.right != nil {
// 		current = current.right
// 	}

// 	return current
// }

func (bst *BinarySearchTree) height(node *node) int {
	if node == nil {
		return 0
	}

	return max(bst.height(node.left), bst.height(node.right)) + 1
}

func (bst *BinarySearchTree) contains(node *node, element int) bool {

	if node == nil {
		return false
	}

	cmp := compareTo(element, node.data)

	if cmp < 0 {
		return bst.contains(node.left, element)
	}

	if cmp > 0 {
		return bst.contains(node.right, element)
	}

	return true
}

func (s *stack) push(node *node) {
	s.items = append(s.items, node)
	s.size++
}

func (q *queue) enqueue(node *node) {
	q.items = append(q.items, node)
	q.size++
}

func (s *stack) pop() *node {
	removedData := s.items[s.size-1]
	s.items = s.items[:s.size-1]
	s.size--
	return removedData
}

func (q *queue) dequeue() *node {
	removedData := q.items[0]
	q.items = q.items[1:]
	q.size--
	return removedData
}

func (s *stack) isEmpty() bool {
	return s.size == 0
}

func (q *queue) isEmpty() bool {
	return q.size == 0
}

func (bst *BinarySearchTree) preorder() ([]int, error) {
	expectedNodeCount := bst.nodeCount
	stack := stack{}
	stack.push(bst.root)
	var data []int

	for bst.root != nil && !stack.isEmpty() {

		if expectedNodeCount != bst.nodeCount {
			return nil, errors.New("modification detected during iteration")
		}

		node := stack.pop()

		if node.right != nil {
			stack.push(node.right)
		}

		if node.left != nil {
			stack.push(node.left)
		}

		data = append(data, node.data)
	}

	return data, nil
}

func (bst *BinarySearchTree) inorder() ([]int, error) {
	expectedNodeCount := bst.nodeCount
	stack := stack{}
	stack.push(bst.root)
	travNode := bst.root
	var data []int

	for bst.root != nil && !stack.isEmpty() {

		if expectedNodeCount != bst.nodeCount {
			return nil, errors.New("modification detected during iteration")
		}

		for ; travNode.left != nil; travNode = travNode.left {
			stack.push(travNode.left)
		}

		node := stack.pop()
		data = append(data, node.data)

		if node.right != nil {
			stack.push(node.right)
			travNode = node.right
		}
	}

	return data, nil
}

func (bst *BinarySearchTree) postorder() ([]int, error) {
	expectedNodeCount := bst.nodeCount
	stack1 := stack{}
	stack2 := stack{}
	stack1.push(bst.root)
	var data []int

	for !stack1.isEmpty() {
		node := stack1.pop()

		if node != nil {
			stack2.push(node)

			if node.left != nil {
				stack1.push(node.left)
			}
			if node.right != nil {
				stack1.push(node.right)
			}
		}
	}

	for bst.root != nil && !stack2.isEmpty() {

		if expectedNodeCount != bst.nodeCount {
			return nil, errors.New("modification detected during iteration")
		}

		node := stack2.pop()
		data = append(data, node.data)

	}

	return data, nil
}

func (bst *BinarySearchTree) levelorder() ([]int, error) {
	expectedNodeCount := bst.nodeCount
	queue := queue{}
	queue.enqueue(bst.root)
	var data []int

	for bst.root != nil && !queue.isEmpty() {

		if expectedNodeCount != bst.nodeCount {
			return nil, errors.New("modification detected during iteration")
		}

		node := queue.dequeue()

		if node.left != nil {
			queue.enqueue(node.left)
		}

		if node.right != nil {
			queue.enqueue(node.right)
		}

		data = append(data, node.data)
	}
	return data, nil
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func compareTo(x, y int) int {
	if x > y {
		return 1
	} else if x < y {
		return -1
	} else {
		return 0
	}
}

func NewTree() *BinarySearchTree {
	return &BinarySearchTree{}
}
