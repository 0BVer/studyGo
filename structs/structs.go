package main

import "fmt"

type person struct {
	name   string
	age    int
	getAge func() int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 25

	return &p
}

func main() {

	fmt.Println(person{"Alpha", 20, func() int { return 20 }})
	fmt.Println(person{name: "Bravo", age: 22, getAge: func() int { return 22 }})
	fmt.Println(person{name: "Charlie"})
	fmt.Println(&person{name: "Delta", age: 24})
	fmt.Println(newPerson("Echo"))

	s := person{name: "Foxtrot", age: 30}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(s.age)
	fmt.Println(sp.age)
	sp.getAge = func() int { return 51 }
	fmt.Println(sp.getAge())
}
