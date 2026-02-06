package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication() {
	//defer berfungsi untuk menjalankan perintah func tersebut diakhir walaupun terjadi error tetap akan dijalankan
	defer logging()
	fmt.Println("Run application")
}

func main() {
	runApplication()
}
