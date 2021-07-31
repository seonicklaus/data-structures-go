package main

import (
	"fmt"
	"hash/fnv"
	"strings"
)

const (
	defaultCapacity   = 3
	defaultLoadFactor = 0.75
)

type HashTable struct {
	maxLoadFactor             float64
	capacity, threshold, size uint
	table                     []*bucket
}

type bucket struct {
	head *bucketNode
}

type bucketNode struct {
	data *entry
	next *bucketNode
}

type entry struct {
	key   interface{}
	value interface{}
	hash  uint64
}

// Initialize Hash Table function
func Init(capacity uint, maxLoadFactor float64) *HashTable {

	result := &HashTable{
		capacity:      max(defaultCapacity, capacity),
		maxLoadFactor: maxLoadFactor,
		threshold:     uint(float64(capacity) * maxLoadFactor),
		table:         make([]*bucket, capacity),
		size:          0,
	}

	return result
}

// Get size of Hash Table
func (ht *HashTable) Size() uint {
	return ht.size
}

// Check if Hash Table is empty
func (ht *HashTable) IsEmpty() bool {
	return ht.size == 0
}

func (ht *HashTable) Clear() {
	ht.table = make([]*bucket, defaultCapacity)
	ht.capacity = defaultCapacity
	ht.maxLoadFactor = defaultLoadFactor
	ht.size = 0
	ht.threshold = uint(float64(ht.capacity) * ht.maxLoadFactor)
}

func (ht *HashTable) Get(key interface{}) interface{} {

	if key == nil {
		return nil
	}

	bucketIndex := ht.normalizeIndex(ht.hashCode(key))
	entry := ht.bucketSeekEntry(bucketIndex, key)

	if entry != nil {
		return entry
	}

	return nil
}

func (ht *HashTable) normalizeIndex(keyHash uint64) int {
	return int(uint64(ht.capacity-1) & keyHash)
}

func (ht *HashTable) bucketSeekEntry(index int, key interface{}) interface{} {

	if key == nil {
		return nil
	}

	bucket := ht.table[index]

	if bucket == nil {
		return nil
	}

	bucketNode := bucket.head

	for ; bucketNode.next != nil; bucketNode = bucketNode.next {
		entry := bucketNode.data

		if entry.key == key {
			return entry
		}
	}

	return nil
}

func (ht *HashTable) hashCode(key interface{}) uint64 {
	h := fnv.New64a()
	_, _ = h.Write([]byte(fmt.Sprintf("%v", key)))

	hashValue := h.Sum64()
	return (hashValue ^ (hashValue >> 16))
}

func (entry *entry) equals(other *entry) bool {
	if entry.hash != other.hash {
		return false
	}

	return entry.key == other.key
}

func (entry *entry) String() string {
	sb := strings.Builder{}

	sb.WriteString(fmt.Sprintf("%v -> %v", entry.key, entry.value))

	return sb.String()
}

func max(x uint, y uint) uint {
	if x > y {
		return x
	} else {
		return y
	}
}

func main() {
}
