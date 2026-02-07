package main

import (
	"errors"
	"fmt"
)

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("pembagi tidak boleh nol")
	}
	return nilai / pembagi, nil
}

func main() {
	hasil, err := Pembagian(10, 0)
	if err == nil {
		fmt.Println("Pembagian berhasil:", hasil)
	} else {
		fmt.Println("Terjadi kesalahan:", err.Error())
	}
}
