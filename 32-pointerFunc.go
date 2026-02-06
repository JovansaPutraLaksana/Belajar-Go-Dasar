package main

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	var alamat1 Address = Address{}
	alamat1.City = "Jakarta"
	alamat1.Province = "DKI Jakarta"
	alamat1.Country = "USA"
	ChangeCountryToIndonesia(&alamat1)

	println(alamat1.City)     // Output: Jakarta
	println(alamat1.Province) // Output: DKI Jakarta
	println(alamat1.Country)  // Output: Indonesia
}
