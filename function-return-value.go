package main

import "fmt"

func getHello(name string) string {
	if name == "" {
		return "Hello World"
	} else {
		return "Hello" + name
	}
}
func main() {
	result := getHello("Aldi")
	fmt.Println(result)

	fmt.Println(getHello(""))

	getHello("Budi")
}
