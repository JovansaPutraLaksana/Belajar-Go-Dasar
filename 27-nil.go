package main

import "fmt"

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
