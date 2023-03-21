package main

import "fmt"

type Costumer struct {
	Name, Address string
	Age           int
}

func main() {
	var aldi Costumer
	aldi.Name = "Aldi"
	aldi.Address = "Jakarta"
	aldi.Age = 21
	fmt.Println(aldi)

	Hardi := Costumer{
		Name:    "Hardi",
		Address: "Jakarta",
		Age:     31,
	}
	fmt.Println(Hardi)

	Afina := Costumer{"Afina", "Jakarta", 20}
	fmt.Println(Afina)
}
