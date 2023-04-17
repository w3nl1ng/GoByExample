package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Alice", age: 30})

	fmt.Println(person{name: "jack"})

	fmt.Println(&person{name: "Ann", age:40})

	fmt.Println(newPerson("jon"))

	s := person{name:"tom", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(s.age)

	sp.age = 51
	fmt.Println(sp.age)

}