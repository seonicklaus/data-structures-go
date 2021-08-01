package hashtable

import (
	"errors"
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

// Get size of Hash Map
func (ht *HashTable) Size() uint {
	return ht.size
}

// Check if Hash Map is empty
func (ht *HashTable) IsEmpty() bool {
	return ht.size == 0
}

// Clear Hash Map data
func (ht *HashTable) Clear() {
	ht.table = make([]*bucket, ht.capacity)
	ht.size = 0
}

// Check if key is present in Hash Map
func (ht *HashTable) ContainsKey(key interface{}) bool {
	return ht.hasKey(key)
}

// Returns a value when a key is passed in, returns nil otherwise
func (ht *HashTable) Get(key interface{}) (interface{}, error) {
	if key == nil {
		return nil, errors.New("key is nil")
	}

	index := ht.normalizeIndex(ht.hashCode(key))
	entry := ht.bucketSeekEntry(index, key)

	if entry != nil {
		return entry.value, nil
	}

	return nil, errors.New("key not found in hash map")
}

// Public method to add key value pair to Hash Map
func (ht *HashTable) Add(key, value interface{}) (interface{}, error) {
	return ht.insert(key, value)
}

// Public method to remove key value pair, returns value when suceed, nil and error otherwise
func (ht *HashTable) Remove(key interface{}) (interface{}, error) {
	if key == nil {
		return nil, errors.New("key is nil")
	}

	if ht.IsEmpty() {
		return nil, errors.New("hash map is empty")
	}

	bucketIndex := ht.normalizeIndex(ht.hashCode(key))
	return ht.bucketRemoveEntry(bucketIndex, key), nil
}

// Returns an array of keys in Hash Map
func (ht *HashTable) Keys() []interface{} {
	var keys []interface{}

	for _, bucket := range ht.table {
		if bucket != nil {
			node := bucket.head

			for ; node != nil; node = node.next {
				keys = append(keys, node.data.key)
			}
		}
	}

	return keys
}

func (ht *HashTable) Values() []interface{} {
	var values []interface{}

	for _, bucket := range ht.table {
		if bucket != nil {
			node := bucket.head

			for ; node != nil; node = node.next {
				values = append(values, node.data.value)
			}
		}
	}

	return values
}

// Insert key and value to Hash Map entry
func (ht *HashTable) insert(key, value interface{}) (interface{}, error) {
	if key == nil {
		return nil, errors.New("key is nil")
	}
	newEntry := &entry{
		key:   key,
		value: value,
		hash:  ht.hashCode(key),
	}
	index := ht.normalizeIndex(newEntry.hash)
	return ht.bucketInsertEntry(index, newEntry), nil
}

// Inserts an entry to a bucket in an index, returns nil when inserting, returns value for modification
func (ht *HashTable) bucketInsertEntry(index int, entry *entry) interface{} {

	if ht.table[index] == nil {
		ht.table[index] = &bucket{}
	}

	existentEntry := ht.bucketSeekEntry(index, entry.key)

	if existentEntry == nil {
		ht.table[index].add(entry)
		ht.size++

		if ht.size > ht.threshold {
			ht.resizeTable()
		}

		return nil

	} else {
		oldValue := existentEntry.value
		existentEntry.value = entry.value

		return oldValue
	}
}

// Returns removed entry's value from Hash Map, return nil otherwise
func (ht *HashTable) bucketRemoveEntry(index int, key interface{}) interface{} {
	entry := ht.bucketSeekEntry(index, key)

	if entry != nil {
		removedData := entry.value
		ht.table[index].remove(entry)

		if ht.table[index].head == nil {
			ht.table[index] = nil
		}

		ht.size--
		return removedData
	}

	return nil
}

// Returns an entry from a bucket when index and key is passed in
func (ht *HashTable) bucketSeekEntry(index int, key interface{}) *entry {

	if key == nil {
		return nil
	}

	if ht.table[index] == nil || ht.table[index].head == nil {
		return nil
	}

	node := ht.table[index].head

	for ; node != nil; node = node.next {
		entry := node.data

		if entry.key == key {
			return entry
		}
	}

	return nil
}

// Resize table when capacity exceed threshold
func (ht *HashTable) resizeTable() {
	if ht.capacity < (1 << 5) {
		ht.capacity *= 2
	} else {
		ht.capacity = uint(float64(ht.capacity) * 1.5)
	}

	ht.threshold = uint(float64(ht.capacity) * ht.maxLoadFactor)

	newTable := make([]*bucket, ht.capacity)

	for _, indexBucket := range ht.table {

		if indexBucket != nil {
			node := indexBucket.head

			for ; node != nil; node = node.next {
				index := ht.normalizeIndex(node.data.hash)

				if newTable[index] == nil {
					newTable[index] = &bucket{}
				}

				newTable[index].add(node.data)
			}

			indexBucket = nil
		}
	}

	ht.table = newTable
}

func (ht *HashTable) hasKey(key interface{}) bool {
	bucketIndex := ht.normalizeIndex(ht.hashCode(key))
	return ht.bucketSeekEntry(bucketIndex, key) != nil
}

// Add entry to bucket
func (b *bucket) add(entry *entry) {
	newNode := &bucketNode{}
	newNode.data = entry
	newNode.next = b.head
	b.head = newNode
}

func (b *bucket) remove(entry *entry) {

	for node := b.head; node != nil; node = node.next {
		existingEntry := node.data

		if existingEntry.key == entry.key {
			b.head = node.next
			node.data = nil
		}
	}

	if b.head == nil {
		b = nil
	}
}

// Returns an index from generated hash code
func (ht *HashTable) normalizeIndex(hashCode uint64) int {
	return int(hashCode % uint64(ht.capacity))
}

// Hashing function
func (ht *HashTable) hashCode(key interface{}) uint64 {
	h := fnv.New64a()
	h.Write([]byte(fmt.Sprintf("%v", key)))
	hashValue := h.Sum64()

	return hashValue ^ (hashValue >> 16)
}

func (entry *entry) String() string {
	sb := strings.Builder{}

	sb.WriteString(fmt.Sprintf("%v: %v", entry.key, entry.value))

	return sb.String()
}

func (ht *HashTable) String() string {
	sb := strings.Builder{}

	sb.WriteString("{")

	for _, bucket := range ht.table {
		if bucket != nil {
			node := bucket.head

			for ; node != nil; node = node.next {
				sb.WriteString(fmt.Sprintf("%s, ", node.data))
			}
		}
	}

	sb.WriteString("}")

	return sb.String()
}

// Initialize Hash Map
func Init(capacity uint, loadFactor float64) *HashTable {
	result := &HashTable{
		maxLoadFactor: maxFloat(loadFactor, defaultLoadFactor),
		capacity:      maxUint(capacity, defaultCapacity),
		size:          0,
		table:         make([]*bucket, capacity),
	}

	result.threshold = uint(float64(result.capacity) * result.maxLoadFactor)

	return result
}

func maxFloat(x, y float64) float64 {
	if x > y {
		return x
	} else {
		return y
	}
}

func maxUint(x, y uint) uint {
	if x > y {
		return x
	} else {
		return y
	}
}
