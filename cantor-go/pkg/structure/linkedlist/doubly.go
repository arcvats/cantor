package linkedlist

type DoublyNode[T any] struct {
	value    T
	next     *DoublyNode[T]
	previous *DoublyNode[T]
}

type Doubly[T any] struct {
	head *DoublyNode[T]
	tail *DoublyNode[T]
	size uint
}

func NewDoublyNode[T any](value T) *DoublyNode[T] {
	return &DoublyNode[T]{value: value, next: nil, previous: nil}
}

func NewDoubly[T any]() *Doubly[T] {
	return &Doubly[T]{head: nil, tail: nil, size: 0}
}

func (list *Doubly[T]) Size() uint {
	return list.size
}

func (list *Doubly[T]) IsEmpty() bool {
	return list.size == 0
}

func (list *Doubly[T]) Append(value T) {
	newNode := NewDoublyNode[T](value)
	if list.IsEmpty() {
		list.head = newNode
	} else {
		newNode.previous = list.tail
		list.tail.next = newNode
	}
	list.tail = newNode
	list.size++
}

func (list *Doubly[T]) Prepend(value T) {
	newNode := NewDoublyNode[T](value)
	if list.IsEmpty() {
		list.tail = newNode
	} else {
		list.head.previous = newNode
		newNode.next = list.head
	}
	list.head = newNode
	list.size++
}

func (list *Doubly[T]) Destroy() {
	list.head = nil
	list.tail = nil
	list.size = 0
}

func (list *Doubly[T]) DeleteFirst() any {
	if list.IsEmpty() {
		return nil
	}
	value := list.head.value
	if list.Size() == 1 {
		list.Destroy()
		return value
	}
	list.head = list.head.next
	list.head.previous = nil
	list.size--
	return value
}

func (list *Doubly[T]) DeleteLast() any {
	if list.IsEmpty() {
		return nil
	}
	value := list.tail.value
	if list.Size() == 1 {
		return list.DeleteFirst()
	}
	list.tail = list.tail.previous
	list.tail.next = nil
	list.size--
	return value
}
