package main

import "fmt"

func getFullName2() (firstName, middleName, lastName string) {
	firstName = "Aldi"
	middleName = "Resaldi"
	lastName = "Maulana"
	return
}
func main() {
	firstName, middleName, lastName := getFullName2()
	fmt.Println(firstName, middleName, lastName)
}
