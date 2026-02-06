package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Palembang", "SumSel", "Indonesia"}
	address2 := address1  // pass by value
	address3 := &address1 // pass by reference (using pointer)

	address2.City = "Banyuasin"
	address3.City = "Prabumulih"

	fmt.Println(address1.City)
	fmt.Println(address2.City)
	fmt.Println(address3.City)
}
