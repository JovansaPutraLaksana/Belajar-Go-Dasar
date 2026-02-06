package main

import "fmt"

func endApp() {
	fmt.Println("Aplikasi selesai")
	massage := recover()
	if massage != nil {
		fmt.Println("Terjadi error:", massage)
	}
}

func runApps(error bool) {
	defer endApp()
	if error {
		panic("Aplikasi error")
	}
}

func main() {
	//panic berfungsi untuk menghentikan program secara paksa
	runApps(true)
}
