package main

import "fmt"

func main() {
	var Name = "Aldi"

	if Name == "Aldi" {
		fmt.Println("Hello Aldi")
	} else if Name == "Resaldi" {
		fmt.Println("Hello Resaldi")
	} else if Name == "Lana" {
		fmt.Println("Hello Lana")
	} else {
		fmt.Println("Ini Salah")
	}
	if length := len(Name); length > 5 {
		fmt.Println("Terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}

}
