package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Aldi",
		"address": "Jakarta",
	}
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
}
