package main

import "fmt"

func main() {
	name := "Resaldi"

	switch name {
	case "Aldi":
		fmt.Println("Hello Aldi")
	case "Resaldi":
		fmt.Println("Hello Resaldi")
	default:
		fmt.Println("Hello World")
	}
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Terlalu Panjang")
	case false:
		fmt.Println("Nama Sudah Benar")
	}

}
