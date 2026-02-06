package main

type Address struct {
	City, Province, Country string
}

func main() {
	// new berfungsi untuk membuat data struct kosong
	var alamat1 *Address = new(Address)
	var alamat2 *Address = alamat1
	alamat2.City = "Bandung"

	println(alamat1.City) // Output: Bandung
	println(alamat2.City) // Output: Bandung
}
