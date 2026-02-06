package main

import "fmt"

type HasName interface {
	GetName() string // Method
}

type Person struct {
	Name string
}

func (p Person) GetName() string { // Implementing interface method untuk struct Person
	return p.Name
}

func PrintName(entity HasName) {
	fmt.Println("Name:", entity.GetName())
}

func main() {
	PrintName(Person{Name: "Jovan"})
}
