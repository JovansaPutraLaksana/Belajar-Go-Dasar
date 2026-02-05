package main

import "fmt"

func sayHello() {
	fmt.Println("Hello")
}

func sayName(name string) string {
	return "Hello, " + name
}

func add(a int, b int) int {
	return a + b
}

//multiple return values
func getFullName() (string, string) {
	return "John", "Doe"
}

//ignore return value
func ignoreReturnValue() {
	_, lastName := getFullName()
	fmt.Println("Last Name:", lastName)
}

//named return values
func getCoordinates() (x int, y int) {
	x = 10
	y = 20
	return
}

//variadic function
func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

//slice as argument
func sumSlice(numbers []int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

//function as parameter
func executeFunction(fn func(string) string, name string) string {
	return fn(name)
}

func main() {
	say := sayName

	sayHello()

	fmt.Println(say("Alice"))

	fmt.Println(add(5, 3))

	firstName, lastName := getFullName()
	fmt.Println("Full Name:", firstName, lastName)

	ignoreReturnValue()

	x, y := getCoordinates()
	fmt.Println("Coordinates:", x, y)

	fmt.Println("Sum:", sum(1, 2, 3, 4, 5))

	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Sum of slice:", sumSlice(numbers))
	//atau
	fmt.Println("Sum of slice (using ...):", sum(numbers...))

	result := executeFunction(sayName, "Bob")
	fmt.Println(result)
}
