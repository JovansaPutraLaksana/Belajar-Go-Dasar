package main

import "fmt"

func main() {
	var (
		nilaiAkhir      = 90
		absensi         = 80
		lulusNilaiAkhir = nilaiAkhir > 80
		lulusAbsensi    = absensi > 80
		lulus           = lulusNilaiAkhir && lulusAbsensi
	)
	fmt.Println(lulus)
}
