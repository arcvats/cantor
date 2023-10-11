package linkedlist

type Node[T any] struct {
	value T
	next  *Node[T]
}

type Singly[T any] struct {
	head *Node[T]
	tail *Node[T]
	size uint
}

func NewNode[T any](value T) *Node[T] {
	return &Node[T]{value: value, next: nil}
}

func NewSingly[T any]() Singly[T] {
	return Singly[T]{head: nil, tail: nil, size: 0}
}

func (list *Singly[T]) Size() uint {
	return list.size
}

func (list *Singly[T]) IsEmpty() bool {
	return list.size == 0
}

func (list *Singly[T]) Append(value T) {
	newNode := NewNode[T](value)
	if list.IsEmpty() {
		list.head = newNode
		list.tail = newNode
	} else {
		list.tail.next = newNode
		list.tail = newNode
	}
	list.size++
}

func (list *Singly[T]) Prepend(value T) {
	newNode := NewNode[T](value)
	if list.IsEmpty() {
		list.head = newNode
		list.tail = newNode
	} else {
		newNode.next = list.head
		list.head = newNode
	}
	list.size++
}
