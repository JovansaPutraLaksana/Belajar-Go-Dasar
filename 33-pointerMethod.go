package main

type Man struct {
	Name string
	Age  int
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	jovan := Man{"Jovan", 30}
	jovan.Married()
	println(jovan.Name) // Output: Mr. Jovan
}
