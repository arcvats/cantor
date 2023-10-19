package main

import (
	"cantor-go/pkg/structure/heap"
	"fmt"
)

func main() {
	newHeap := heap.NewBinaryHeap[string](func(a, b string) bool {
		return len(a) < len(b)
	})
	newHeap.Insert("hello")
	newHeap.Insert("helloworld")
	newHeap.Insert("arc")
	newHeap.Insert("ac")
	newHeap.Insert("universal")
	ele, er := newHeap.Peek()

	fmt.Printf("%v", newHeap)
	if er != nil {
		fmt.Printf("%v", er)
		return
	}
	fmt.Printf("%v\n", *ele)
	h := newHeap.Sort()
	//newHeap.Delete()
	//ele, er = newHeap.Peek()
	//fmt.Printf("%v\n", *ele)
	//
	//newHeap.Delete()
	//ele, er = newHeap.Peek()
	//fmt.Printf("%v\n", *ele)
	//newHeap.Delete()
	//ele, er = newHeap.Peek()
	//fmt.Printf("%v\n", *ele)
	//newHeap.Delete()
	//ele, er = newHeap.Peek()
	//fmt.Printf("%v\n", *ele)

	fmt.Printf("%v\n%v", h, newHeap)
}
