package main

import "fmt"

func (customer Customers) sayHello(name string) {
	fmt.Println("Hello", name, "my name is", customer.Name)
}

type Customers struct {
	Name    string
	Address string
	Age     int
}

func main() {
	cust := Customers{
		Name:    "Eko",
		Address: "Subang",
		Age:     30,
	}
	cust.sayHello("Budi")
}
