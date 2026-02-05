package main

import "fmt"

func main() {
	// > | < | >= | <= | == | !=
	var (
		name1       = "jovan"
		name2       = "jovan"
		result bool = name1 != name2
	)

	fmt.Println(result)
}
