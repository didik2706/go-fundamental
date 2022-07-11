package main

import (
	"fmt"
)

type Person struct {
	name string
	age int
}

func (m Person) getInfo() {
	fmt.Println(m)
}

func (m *Person) changeName(newName string) {
	m.name = newName
}

func main() {
	person1 := Person{ "Didik", 22 }
	person1.getInfo()
	person1.changeName("Didik Nur Hidayat")
	person1.getInfo()
}