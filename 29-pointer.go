package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	//var address1 Address = Address{"Palembang", "SumSel", "Indonesia"}
	address1 := Address{"Palembang", "SumSel", "Indonesia"}
	address2 := address1 // pass by value
	//var address3 *Address = &address1 // pass by reference (using pointer)
	address3 := &address1 // pass by reference (using pointer)

	address2.City = "Banyuasin"
	address3.City = "Prabumulih"

	fmt.Println(address1.City)
	fmt.Println(address2.City)
	fmt.Println(address3.City)
}
