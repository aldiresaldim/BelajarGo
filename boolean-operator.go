package main

import "fmt"

func main() {
	var NilaiAkhir = 90
	var Absensi = 80

	var LulusNilaiAkhir bool = NilaiAkhir > 80
	var LulusAbsensi bool = Absensi > 80

	var lulus bool = LulusNilaiAkhir && LulusAbsensi

	fmt.Println(lulus)
}
