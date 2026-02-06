package main

import "fmt"

func endApplication() {
	fmt.Println("Aplikasi selesai")
}

func runApp(error bool) {
	defer endApplication()
	if error {
		panic("Aplikasi error")
	}
}

func main() {
	runApp(true)
}
