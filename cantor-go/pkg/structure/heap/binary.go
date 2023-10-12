package heap

import "fmt"

type Element interface {
	Integer | UnsignedInteger | float32 | float64 | ~string
}

type Integer interface {
	int | int8 | int16 | int32 | int64
}

type UnsignedInteger interface {
	uint | uint8 | uint16 | uint32 | uint64
}

type BinaryHeap[T any] struct {
	elements   []T
	comparable func(a, b T) bool
	size       uint
}

// NewBinaryHeap creates a new binary heap (min heap by default)
func NewBinaryHeap[T Element](comparable func(a, b T) bool) *BinaryHeap[T] {
	if comparable == nil {
		comparable = func(a, b T) bool {
			return a < b
		}
	}
	return &BinaryHeap[T]{elements: make([]T, 0), comparable: comparable, size: 0}
}

func (heap *BinaryHeap[T]) Size() uint {
	return heap.size
}

func (heap *BinaryHeap[T]) IsEmpty() bool {
	return heap.size == 0
}

func (heap *BinaryHeap[T]) Peek() (*T, error) {
	if heap.IsEmpty() {
		return nil, fmt.Errorf("heap is empty")
	}
	return &heap.elements[0], nil
}

func (heap *BinaryHeap[T]) swap(i, j uint) {
	heap.elements[i], heap.elements[j] = heap.elements[j], heap.elements[i]
}

func (heap *BinaryHeap[T]) Insert(element T) {
	heap.size++
	heap.elements = append(heap.elements, element)
	index := heap.size - 1
	parentIndex := index / 2
	for index > 0 && heap.comparable(heap.elements[index], heap.elements[parentIndex]) {
		heap.swap(index, parentIndex)
		index = parentIndex
		parentIndex = index / 2
	}
}
