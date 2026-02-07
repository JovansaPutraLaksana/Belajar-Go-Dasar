package main

import (
	"belajar-go-dasar/helper"
	"fmt"
)

// akses modifier pada golang
// jika nama diawali dengan huruf besar, maka dapat diakses dari package lain (public)
// jika nama diawali dengan huruf kecil, maka hanya dapat diakses dari package yang sama (private)
func main() {
	fmt.Println(helper.AppName)
	//fmt.Println(helper.version) //akan error karena version tidak dapat diakses dari package lain
	//fmt.Println(helper.getVersion()) //akan error karena getVersion tidak dapat diakses dari package lain
	fmt.Println(helper.SayHello("Eko"))
}
