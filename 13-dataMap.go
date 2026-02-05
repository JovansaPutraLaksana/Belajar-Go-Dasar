package main

import "fmt"

func main() {
	//deklarasi map[keyType]valueType
	var person map[string]string
	//menambahkan data ke map
	person["nama"] = "Jovansa"
	person["pekerjaan"] = "Programmer"
	person["alamat"] = "Jl. Merdeka No. 123"
	//mengakses data dari map
	fmt.Println(person["nama"])
	fmt.Println(person["pekerjaan"])
	fmt.Println(person["alamat"])
	fmt.Println(person)

	//menghapus data dari map
	delete(person, "alamat")
	fmt.Println(person)

	//menggunakan literal map
	country := map[string]string{
		"nama":     "Indonesia",
		"ibukota":  "Jakarta",
		"penduduk": "270 juta",
		"bahasa":   "Bahasa Indonesia",
	}
	fmt.Println(country)

	//iterasi map
	for key, value := range country {
		fmt.Println(key, ":", value)
	}

	//cek keberadaan key dalam map
	value, exists := country["nama"]
	fmt.Println("Value:", value, "Exists:", exists)

	//menghitung jumlah elemen dalam map
	fmt.Println("Jumlah elemen dalam map country:", len(country))

	//map dengan tipe data campuran
	mixedMap := map[string]interface{}{
		"nama":    "Alice",
		"umur":    30,
		"menikah": true,
	}
	fmt.Println(mixedMap)

	//nested map
	nestedMap := map[string]map[string]string{
		"person1": {
			"nama":   "Bob",
			"alamat": "Jl. Merdeka No. 456",
		},
		"person2": {
			"nama":   "Alice",
			"alamat": "Jl. Sudirman No. 789",
		},
	}
	fmt.Println(nestedMap)

	//menghapus data dalam nested map
	delete(nestedMap["person2"], "alamat")
	fmt.Println(nestedMap)

	//map dengan slice sebagai value
	sliceMap := map[string][]string{
		"fruits":  {"apple", "banana", "cherry"},
		"veggies": {"carrot", "broccoli", "spinach"},
	}
	fmt.Println(sliceMap)

	//mengakses elemen dalam slice yang ada di map
	fmt.Println(sliceMap["fruits"][1]) // Output: banana

	//menambahkan elemen ke slice dalam map
	sliceMap["fruits"] = append(sliceMap["fruits"], "date")
	fmt.Println(sliceMap)

	//iterasi map dengan slice sebagai value
	for key, values := range sliceMap {
		fmt.Println(key, ":")
		for _, value := range values {
			fmt.Println(" -", value)
		}
	}

	//menghitung jumlah elemen dalam map dengan slice sebagai value
	fmt.Println("Jumlah elemen dalam sliceMap:", len(sliceMap))

	//cek keberadaan key dalam map dengan slice sebagai value
	values, exists := sliceMap["veggies"]
	fmt.Println("Values:", values, "Exists:", exists)

	//menghapus elemen dari slice dalam map
	sliceMap["veggies"] = sliceMap["veggies"][:1]
	fmt.Println(sliceMap)

	//map dengan struct sebagai value
	type Person struct {
		Nama   string
		Umur   int
		Alamat string
	}

	//map dengan struct sebagai value
	personMap := map[string]Person{
		"person1": {"Charlie", 28, "Jl. Thamrin No. 101"},
		"person2": {"Diana", 32, "Jl. Gatot Subroto No. 202"},
	}
	fmt.Println(personMap)

	//mengakses data dalam struct yang ada di map
	fmt.Println(personMap["person1"].Nama)

	//mengupdate data dalam struct yang ada di map
	person1 := personMap["person1"]
	person1.Umur = 29
	personMap["person1"] = person1
	fmt.Println(personMap)

	//iterasi map dengan struct sebagai value
	for key, person := range personMap {
		fmt.Println(key, ":", person.Nama, person.Umur, person.Alamat)
	}

	//menghitung jumlah elemen dalam map dengan struct sebagai value
	fmt.Println("Jumlah elemen dalam personMap:", len(personMap))

	//cek keberadaan key dalam map dengan struct sebagai value
	personData, exists := personMap["person2"]
	fmt.Println("Person Data:", personData, "Exists:", exists)

	//menghapus data dalam map dengan struct sebagai value
	delete(personMap, "person2")
	fmt.Println(personMap)
}
