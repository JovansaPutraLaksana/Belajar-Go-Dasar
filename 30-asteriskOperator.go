package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Palembang", "SumSel", "Indonesia"}
	address2 := &address1

	address2.City = "Banyuasin"

	fmt.Println(address1.City)
	fmt.Println(address2.City)
	fmt.Println("========================")

	// address2 = &Address{"Jakarta", "DKI", "Indonesia"}
	// fmt.Println(address1.City)
	// fmt.Println(address2.City)
	// fmt.Println("========================")

	// using asterisk operator to change the value of address1 via address2 pointer
	*address2 = Address{"Bandung", "Jabar", "Indonesia"}

	fmt.Println(address1.City)
	fmt.Println(address2.City)

}
