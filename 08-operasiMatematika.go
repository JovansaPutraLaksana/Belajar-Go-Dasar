package main

import "fmt"

func main() {
	//Operasi matematika dasar + - * / %
	var (
		a = 10
		b = 20
		c = 5
		d = a + b - c
	)
	fmt.Println(d)

	//Augmented Assignments (+=,-=,*=,/=,%=)
	var (
		i = 10
	)
	i += 10 // i=i+10
	i -= 10 //i=i-10
	i *= 2  //i=i*2
	i /= 2  //i=i/2
	i %= 10 //i=i%10

	//Unary Operator
	a++
	a--
}
