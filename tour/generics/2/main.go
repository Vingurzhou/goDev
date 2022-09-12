package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

// create nil list of type T
func NewList[T any]() *List[T] {
	return nil
}

// insert val before list
func (list *List[T]) Insert(val T) *List[T] {
	return &List[T]{
		next: list,
		val:  val,
	}
}

// insert val after list
func (list *List[T]) Append(val T) *List[T] {
	if list == nil {
		return list.Insert(val)
	}

	node := &List[T]{
		next: list.next,
		val:  val,
	}
	list.next = node
	return list
}

// return new list with elements reversed
func (list *List[T]) Reverse() *List[T] {
	var reversed *List[T]
	for cursor := list; cursor != nil; cursor = cursor.next {
		reversed = reversed.Insert(cursor.val)
	}
	return reversed
}

// traverse list one by one
func (list *List[T]) Traverse(f func(val T, l *List[T])) {
	for cursor := list; cursor != nil; cursor = cursor.next {
		f(cursor.val, cursor)
	}
}

// print list
func (list *List[T]) Print() {
	list.Traverse(func(val T, _ *List[T]) {
		fmt.Printf("%s ", val)
	})
}

// print list with line break
func (list *List[T]) Println() {
	list.Print()
	fmt.Println()
}

func main() {
	// initialize a new list
	l := NewList[string]().Insert("apple").Insert("banana").Append("cat")
	// print it
	l.Println()

	// reverse it
	l = l.Reverse()
	// print it again
	l.Println()
}
