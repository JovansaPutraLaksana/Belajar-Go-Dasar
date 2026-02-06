package main

import "fmt"

// Nil hanya dapat digunakan untuk tipe data referensi seperti pointer, slice, map, channel, fungsi, dan interface.
// Nil menunjukkan bahwa variabel tersebut tidak menunjuk ke nilai atau objek apa pun.

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	}
	return map[string]string{
		"name": name,
	}
}

func main() {
	m := NewMap("")
	if m == nil {
		fmt.Println("Map is nil")
	}
	m2 := NewMap("Gopher")
	if m2 != nil {
		fmt.Println("Map is not nil:", m2)
	}
}
