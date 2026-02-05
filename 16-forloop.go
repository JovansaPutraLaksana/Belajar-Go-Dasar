package main

import "fmt"

func main() {
	//perulangan for
	//for inisialisasi; kondisi; post {
	//	//blok kode yang akan diulang
	//}
	for i := 1; i <= 5; i++ {
		fmt.Println("Perulangan ke-", i)
	}

	//for dengan kondisi saja
	j := 1
	for j <= 5 {
		fmt.Println("Perulangan ke-", j)
		j++
	}

	//for range
	nama := "Golang"
	for index, char := range nama {
		fmt.Printf("Index: %d, Karakter: %c\n", index, char)
	}

	angka := []int{10, 20, 30, 40, 50}
	for index, value := range angka {
		fmt.Printf("Index: %d, Nilai: %d\n", index, value)
	}

	//for dengan break dan continue
	for k := 1; k <= 10; k++ {
		if k%2 == 0 {
			continue //melewati angka genap
		}
		if k > 7 {
			break //menghentikan perulangan jika k lebih dari 7
		}
		fmt.Println("Angka ganjil:", k)
	}

	//perulangan bersarang (nested loop)
	for a := 1; a <= 3; a++ {
		for b := 1; b <= 2; b++ {
			fmt.Printf("a: %d, b: %d\n", a, b)
		}
	}

	//for dengan short statement
	for sum := 0; sum <= 20; sum += 5 {
		fmt.Println("Sum:", sum)
	}

	//menggunakan for untuk iterasi slice
	slice := []string{"apel", "pisang", "jeruk"}
	for i := 0; i < len(slice); i++ {
		fmt.Println("Buah:", slice[i])
	}

	//menggunakan for untuk iterasi map
	mapping := map[string]int{"satu": 1, "dua": 2, "tiga": 3}
	for key, value := range mapping {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}

	//menggunakan for untuk menghitung mundur
	for countdown := 5; countdown >= 1; countdown-- {
		fmt.Println("Countdown:", countdown)
	}

	//menggunakan for untuk membuat pola bintang
	n := 5
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	//menggunakan for untuk menjumlahkan elemen dalam slice
	numbers := []int{1, 2, 3, 4, 5}
	total := 0
	for _, num := range numbers {
		total += num
	}
	fmt.Println("Total penjumlahan:", total)

	//menggunakan for untuk mencari nilai maksimum dalam slice
	nums := []int{10, 5, 8, 20, 15}
	max := nums[0]
	for _, val := range nums {
		if val > max {
			max = val
		}
	}
	fmt.Println("Nilai maksimum:", max)

	//menggunakan for untuk mengubah elemen dalam slice
	strs := []string{"go", "is", "fun"}
	for i := 0; i < len(strs); i++ {
		strs[i] = strs[i] + "!"
	}
	fmt.Println("Slice setelah diubah:", strs)

	//menggunakan for untuk menghapus elemen dari slice
	sliceToRemove := []int{1, 2, 3, 4, 5, 6}
	valueToRemove := 4
	newSlice := []int{}
	for _, v := range sliceToRemove {
		if v != valueToRemove {
			newSlice = append(newSlice, v)
		}
	}
	fmt.Println("Slice setelah penghapusan:", newSlice)

	//menggunakan for untuk menggabungkan dua slice
	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6}
	combined := []int{}
	for _, v := range slice1 {
		combined = append(combined, v)
	}
	for _, v := range slice2 {
		combined = append(combined, v)
	}
	fmt.Println("Combined slice:", combined)

	//menggunakan for untuk menghitung frekuensi elemen dalam slice
	elements := []string{"a", "b", "a", "c", "b", "a"}
	frequency := make(map[string]int)
	for _, e := range elements {
		frequency[e]++
	}
	fmt.Println("Frequency:", frequency)

}
