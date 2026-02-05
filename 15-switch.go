package main

import "fmt"

func main() {
	//percabangan switch
	name := "Budi"
	switch name {
	case "Andi":
		fmt.Println("Halo Andi!")
	case "Budi":
		fmt.Println("Halo Budi!")
	case "Citra":
		fmt.Println("Halo Citra!")
	default:
		fmt.Println("Halo, siapa kamu?")
	}

	//switch dengan ekspresi
	nilai := 85
	switch {
	case nilai >= 85:
		fmt.Println("Nilai Anda A")
	case nilai >= 70:
		fmt.Println("Nilai Anda B")
	case nilai >= 60:
		fmt.Println("Nilai Anda C")
	default:
		fmt.Println("Nilai Anda D")
	}

	//switch dengan multiple case
	hari := "Sabtu"
	switch hari {
	case "Sabtu", "Minggu":
		fmt.Println("Hari ini adalah akhir pekan.")
	default:
		fmt.Println("Hari ini adalah hari kerja.")
	}

	//switch dengan short statement
	switch absen := 80; {
	case absen >= 75:
		fmt.Println("Anda lulus absen.")
	default:
		fmt.Println("Anda tidak lulus absen.")
	}
}
