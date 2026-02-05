package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Jovansa"
	names[1] = "Putra"
	names[2] = "Laksana"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		90, 80,
	}
	var values2 = [...]int{ // ukuran array sesuai dengan isi awal declarasi
		90, 80,
	}
	fmt.Println(values)
	fmt.Println(len(values))
	fmt.Println(len(values2))
}
