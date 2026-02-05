package main

import "fmt"

func main() {
	type noKTP string

	var (
		ktpJovan  noKTP  = "111111"
		contoh    string = "222222"
		contohKTP noKTP  = noKTP(contoh)
	)

	fmt.Println(ktpJovan, contohKTP)
}
