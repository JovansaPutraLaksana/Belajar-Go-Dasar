package main

import "fmt"

func main() {
	//percabangan if
	//if kondisi {
	//	//blok kode yang dijalankan jika kondisi bernilai true
	//}
	nilai := 75
	if nilai >= 60 {
		fmt.Println("Selamat, Anda lulus!")
	}
	//if else
	if nilai >= 60 {
		fmt.Println("Selamat, Anda lulus!")
	} else {
		fmt.Println("Maaf, Anda tidak lulus.")
	}

	//if else if else
	if nilai >= 85 {
		fmt.Println("Nilai Anda A")
	} else if nilai >= 70 {
		fmt.Println("Nilai Anda B")
	} else if nilai >= 60 {
		fmt.Println("Nilai Anda C")
	} else {
		fmt.Println("Nilai Anda D")
	}

	//nested if
	absen := 80
	if nilai >= 60 {
		if absen >= 75 {
			fmt.Println("Selamat, Anda lulus!")
		} else {
			fmt.Println("Maaf, Anda tidak lulus karena absen kurang.")
		}
	} else {
		fmt.Println("Maaf, Anda tidak lulus karena nilai kurang.")
	}

	//if short statement
	if rataRata := (nilai + absen) / 2; rataRata >= 70 {
		fmt.Println("Selamat, Anda lulus dengan rata-rata", rataRata)
	}
}
