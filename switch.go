package main

import (
	"fmt"
)

func main() {
	name := "aldi"

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
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama Terlalu Panjang")
	case length > 5:
		fmt.Println("Nama Lumayan Panjang")
	case length > 10:
		fmt.Println("Nama Sudah Benar")
	}
}
