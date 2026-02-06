package main

import "fmt"

/*
struct adalah kumpulan dari beberapa tipe data yang berbeda
mirip dengan class pada bahasa pemrograman lain
digunakan untuk mengelompokkan beberapa data menjadi satu kesatuan
sehingga memudahkan dalam pengelolaan data
penulisan struct (secara pascal case (dimulai dengan huruf besar))
*/
type Customer struct {
	Name    string
	Address string
	Age     int
}

func main() {
	var cust Customer
	cust.Name = "Eko"
	cust.Address = "Subang"
	cust.Age = 30
	fmt.Println(cust)

	cust2 := Customer{
		Name:    "Budi",
		Address: "Jakarta",
		Age:     25,
	}
	fmt.Println(cust2)

	cust3 := Customer{"Joko", "Bandung", 35}
	fmt.Println(cust3)

	cust4 := Customer{
		Name:    "Ani",
		Address: "Surabaya",
		Age:     28,
	}
	fmt.Println(cust4)
}
