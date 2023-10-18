package linkedlist

type SinglyNode[T any] struct {
	value T
	next  *SinglyNode[T]
}

type Singly[T any] struct {
	head *SinglyNode[T]
	tail *SinglyNode[T]
	size uint
}

func NewSinglyNode[T any](value T) *SinglyNode[T] {
	return &SinglyNode[T]{value: value, next: nil}
}

// NewSingly creates a new singly linked list
func NewSingly[T any]() *Singly[T] {
	return &Singly[T]{head: nil, tail: nil, size: 0}
}

func (list *Singly[T]) Size() uint {
	return list.size
}

func (list *Singly[T]) IsEmpty() bool {
	return list.size == 0
}

func (list *Singly[T]) Append(value T) {
	newNode := NewSinglyNode[T](value)
	if list.IsEmpty() {
		list.head = newNode
	} else {
		list.tail.next = newNode
	}
	list.tail = newNode
	list.size++
}

func (list *Singly[T]) Prepend(value T) {
	newNode := NewSinglyNode[T](value)
	if list.IsEmpty() {
		list.tail = newNode
	} else {
		newNode.next = list.head
	}
	list.head = newNode
	list.size++
}

func (list *Singly[T]) Destroy() {
	list.head = nil
	list.tail = nil
	list.size = 0
}

func (list *Singly[T]) DeleteFirst() any {
	if list.IsEmpty() {
		return nil
	}
	if list.Size() == 1 {
		list.tail = nil
	}
	value := list.head.value
	list.head = list.head.next
	list.size--
	return value
}

func (list *Singly[T]) DeleteLast() any {
	if list.IsEmpty() {
		return nil
	}
	if list.Size() == 1 {
		return list.DeleteFirst()
	}
	value := list.tail.value
	current := list.head
	for current.next.next != nil {
		current = current.next
	}
	current.next = nil
	list.tail = current
	list.size--
	return value
}
