package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var Aldi NoKTP = "3327062408999"
	var MarriedStatus Married = false

	fmt.Println(Aldi)
	fmt.Println(MarriedStatus)
}
