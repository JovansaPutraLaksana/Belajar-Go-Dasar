package main

import "fmt"

func main() {
	var name string             //bila variable tidak dilakukan inisialisasi wajib disertakan tipe data
	var umur = 21               //bila variable dilakukan inisialisasi variable dapat otomatis untuk generate tipe data
	jenisKelamin := "Laki-laki" //var dapat diganti dengan := untuk setiap deklarasi pertama variable

	//multiple variable
	var (
		firstName  = "Jovansa"
		middleName = "Putra"
		lastName   = "Laksana"
	)

	name = "Jovansa"
	fmt.Println(name, umur, jenisKelamin)

	name = "Jovansa Putra Laksana"
	fmt.Println(name, umur, jenisKelamin)

	fmt.Println(firstName, middleName, lastName)
}
