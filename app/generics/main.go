package main

import "fmt"

func mapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

type element[T any] struct {
	next *element[T]
	val  T
}

type List[T any] struct {
	head, tail *element[T]
}

func (lst *List[T]) push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v, next: nil}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v, next: nil}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) getAll() []T {
	r := make([]T, 0)
	for e := lst.head; e != nil; e = e.next {
		v := e.val
		r = append(r, v)
	}
	return r
}

func main() {
	var m = map[int]string{1: "2", 2: "4", 4: "8"}

	fmt.Println(mapKeys(m))
	fmt.Println(mapKeys[int, string](m))

	lst := List[int]{}
	lst.push(20)
	lst.push(30)
	lst.push(40)

	fmt.Println(lst.getAll())
}