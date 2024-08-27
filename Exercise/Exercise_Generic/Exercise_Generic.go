package main

import (
	"fmt"
)

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func (list *List[T]) String() string {
	if list.next != nil {
		return fmt.Sprintf("val %v, next : next: %s", list.val, list.next)
	} else {
		return fmt.Sprintf("val %v, next : next: <nil>", list.val)
	}
}

func Append[T any](xs *List[T], x T)  *List[T]{
	if xs.next == nil {
		xs.next = NewList(x)
		return xs
	} else {
		
		ys := xs
		ys = Append(ys.next, x)
		xs = ys
	}
	return xs
}

func NewList[T any](val T) *List[T] {
	xs := new(List[T])
	xs.val = val
	xs.next = nil
	return xs
} 

func main() {
	xs := NewList(0)
	fmt.Println(xs)

	for i := 1; i <= 10; i++ {
		xs = Append(xs, i)
		fmt.Println(xs)
	}
}
