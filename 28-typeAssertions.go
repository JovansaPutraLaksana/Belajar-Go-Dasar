package main

import "fmt"

func random() any {
	return "OK"
}

func main() {
	var result any = random()
	var resultStr string = result.(string)
	fmt.Println("String value:", resultStr)

	// Type assertion dengan pengecekan ok
	//result := random()
	str, ok := result.(string)
	if ok {
		fmt.Println("String value:", str)
	} else {
		fmt.Println("result is not a string")
	}
	// Contoh lain: type assertion yang gagal
	// num, ok := result.(int)
	// if ok {
	// 	fmt.Println("Integer value:", num)
	// } else {
	// 	fmt.Println("result is not an integer")
	// }
	// Type assertion tanpa pengecekan ok, akan menyebabkan panic jika gagal
	// Uncomment baris berikut untuk melihat panic
	// num2 := result.(int)
	// fmt.Println("Integer value:", num2)

	// Type switch
	switch value := result.(type) {
	case string:
		fmt.Println("result is a string:", value)
	case int:
		fmt.Println("result is an integer:", value)
	default:
		fmt.Println("Unknown type")
	}
}
