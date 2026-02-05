package main

import "fmt"

func main() {
	//slice merupakan potongan dari array
	//slice mirip dengan array namun ukurannya dapat berubah
	//slice dan array saling terkoneksi, dimana slice adalah data yang mengakses sebagian atau seluruh data array
	//tipe data slice memiliki 3 data, yaitu pointer,length, dan capasity
	//pointer penunjuk data pertama di array para slice
	//length panjang dari slice
	//capasity adalah kapasitas dari slice, dimana length tidak boleh lebih dari capasity

	//membuat slice dari array
	//array[low:high] =>membuat array dari index low sampai sebelum high
	//array[low:]     =>membuat slice dimulai dari low sampai akhir array
	//array[:high]    =>membuat slice dimulai dari index 0 sampai index sebelum high
	//array[:]        =>membuat slice dari 0 sampai akhir array

	//cara membuat slice dari array yang telah ada
	names := [...]string{"A", "B", "C", "D", "E", "F", "G"}
	slice1 := names[4:6]
	slice2 := names[:3]
	slice3 := names[3:]
	slice4 := names[:]
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)

	//cara membuat slice dari awal
	var slice5 []string

	//function dari slice
	// len(slice) (melihat panjang slice),
	// cap(slice) (melihat kapasitas),
	// append(slice,data) (menambah data ke slice dengan cara membuat array baru sebagai dasar dari slice),
	// make([]typedata,length,capacity) (membuat slice baru),
	// copy(destination,source) (menyalin slice dari source ke destination)
	fmt.Println(slice5)
	fmt.Println(len(slice5))
	fmt.Println(cap(slice5))
	slice5 = append(slice5, "10")
	fmt.Println(slice5)
	fmt.Println(len(slice5))
	fmt.Println(cap(slice5))
	fmt.Println("===================================================")
	//latihan
	days := [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}
	daySlice1 := days[5:]
	fmt.Println(days)
	fmt.Println(daySlice1)
	daySlice1[0] = "sabtu baru"
	daySlice1[1] = "minggu baru"
	fmt.Println(days)
	fmt.Println(daySlice1)

	daySlice2 := append(daySlice1, "Libur") // apabila length slice melebihi capasitas, maka slice akan membuat sebuah array baru, sehingga perubahan dalam slice tidak akan mengubah data dalam array
	daySlice2[0] = "sabtu lama"
	fmt.Println(daySlice2)
	fmt.Println(days)
	fmt.Println("===================================================")

	newSlices := make([]string, 2, 5)
	newSlices[0] = "J"
	newSlices[1] = "o"
	//newSlices[2] = "v"  tidak bisa bertambah karena saat membuat slice length-nya 2, sehingga perlu append untuk menambah
	fmt.Println(newSlices)
	fmt.Println(len(newSlices))
	fmt.Println(cap(newSlices))
	newSlices = append(newSlices, "v")
	fmt.Println(newSlices)
	fmt.Println(len(newSlices))
	fmt.Println(cap(newSlices))
	fmt.Println("===================================================")

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	copy(toSlice, fromSlice)
	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	//hati-hati saat membuat slice terdapat perbedaan pembuatan
	//slice []typedata{}
	//array [...]typedata{}

}
