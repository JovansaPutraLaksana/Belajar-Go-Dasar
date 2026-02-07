package helper

var version = "1.0.0"
var AppName = "Belajar Go"

func getVersion() string {
	return version
}

func SayHello(name string) string {
	return "Hello, " + name + "!"
}
